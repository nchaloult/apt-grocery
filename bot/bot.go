package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"apt-grocery/storage"

	"github.com/julienschmidt/httprouter"
)

// groupmeCallbackBody represents the bodies of POST requests which GroupMe
// sends each time a new message is posted in a group that your bot is in.
type groupmeCallbackBody struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

// Bot exposes behaviors for a GroupMe bot
type Bot struct {
	id     string
	name   string
	prefix string
}

// NewBot returns a pointer to a new Bot struct
func NewBot(id, name, prefix string) *Bot {
	return &Bot{
		id:     id,
		name:   name,
		prefix: prefix,
	}
}

// SendMessage sends the provided message to a GroupMe group as the bot.
//
// Currently not configured to choose which group to send the message to;
// assumes that bot is a member of only one group.
func (b *Bot) SendMessage(msg string) {
	log.Printf("SendMessage(): invoked with message: \"%s\"", msg)

	// Create POST request body to send to GroupMe's API
	reqBodyAsStr := fmt.Sprintf(`{"text": "%s", "bot_id": "%s"}`, msg, b.id)
	reqBodyAsJSONBytes := []byte(reqBodyAsStr)

	// Turn that request body into an HTTP request
	req, err := http.NewRequest("POST", "https://api.groupme.com/v3/bots/post", bytes.NewBuffer(reqBodyAsJSONBytes))
	if err != nil {
		// TODO: better error handling
		log.Print("SendMessage(): Failed to create new POST request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Fire off the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Print("SendMessage(): Failed to send POST request")
	}
	defer res.Body.Close()

	log.Printf("SendMessage(): sent POST request. Response status code: %v\n", res.Status)
}

// ProcessMessage is called when a new message is posted in our GroupMe group.
// Decides if the bot was invoked in someone's message, and if so, how to
// handle it.
func (b *Bot) ProcessMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Unmarshal info about the GroupMe message into a groupmeCallbackBody{}
	groupmeMessage := groupmeCallbackBody{}
	err := json.NewDecoder(r.Body).Decode(&groupmeMessage)
	if err != nil {
		log.Print("ProcessMessage(): Failed to parse GroupMe request to callback")
		http.Error(w, "Failed to parse GroupMe request to callback: ", http.StatusInternalServerError)
		return
	}

	// If our bot was invoked:
	if groupmeMessage.Name != b.name && strings.HasPrefix(groupmeMessage.Text, b.prefix) {
		// Remove the bot's prefix from the message contents
		input := strings.TrimSpace(groupmeMessage.Text)[len(b.prefix)+1:]

		if input == "view" {
			log.Print("ProcessMessage(): bot invoked with \"view\" command")

			list := storage.ReadList()

			if len(list) < 1 {
				b.SendMessage("The list is empty")
			} else {
				// Send a new message for each user's own grocery list
				for user, list := range list {
					b.SendMessage(user + ": " + strings.Join(list, ", "))
				}
			}
		} else if input == "prices" {
			log.Print("ProcessMessage(): bot invoked with \"prices\" command")

			prices := storage.ReadPrices()

			if len(prices) < 1 {
				b.SendMessage("There aren't any stored prices for commonly-purchased items")
			} else {
				// Send one message with all of the prices that we store
				//
				// This one message could get very, very long as we store more and more prices.
				// Think about a better solution for displaying which prices we've stored.
				msg := ""
				for item, price := range prices {
					msg += fmt.Sprintf("%s: %v, ", item, price)
				}

				// Trim off trailing comma
				b.SendMessage(strings.TrimSuffix(msg, ", "))
			}
		} else if input[:5] == "clear" {
			log.Print("ProcessMessage(): bot invoked with \"clear\" command")

			storage.ClearList()
			b.SendMessage("List cleared")
		} else if input[:3] == "add" {
			log.Print("ProcessMessage(): bot invoked with \"add\" command")

			// Separate each new list item in user input by commas
			storage.WriteList(groupmeMessage.Name, strings.Split(input[4:], ", "))
			b.SendMessage("Added: " + input[4:])
		} else if input[:5] == "price" {
			log.Print("ProcessMessage(): bot invoked with \"price\" command")

			input := strings.Split(input[6:], " ")
			providedPrice, err := strconv.ParseFloat(input[1], 32)
			if err != nil {
				b.SendMessage("Couldn't turn your provided price into a number: " + input[1])
			} else {
				storage.WritePrice(input[0], float32(providedPrice))
				b.SendMessage("Added price for " + input[0] + ": " + input[1])
			}
		} else {
			log.Print("ProcessMessage(): command/input not recognized.")

			b.SendMessage(fmt.Sprintf("Command/input not recognized: %s", groupmeMessage.Text))
		}
	}
}
