package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"apt-grocery/list"

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
	botID string
}

// NewBot returns a pointer to a new Bot struct
func NewBot(botID string) *Bot {
	return &Bot{botID: botID}
}

// SendMessage sends the provided message to a GroupMe group as the bot.
//
// Currently not configured to choose which group to send the message to;
// assumes that bot is a member of only one group.
func (b *Bot) SendMessage(msg string) {
	log.Printf("SendMessage() invoked with message: %s", msg)

	// Create POST request body to send to GroupMe's API
	reqBodyAsStr := fmt.Sprintf(`{"text": "%s", "bot_id": "%s"}`, msg, b.botID)
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
	if groupmeMessage.Name != "apt grocery" && strings.HasPrefix(groupmeMessage.Text, ".gl") {
		// Remove the ".gl" prefix from the message contents
		input := strings.TrimSpace(groupmeMessage.Text)[4:]

		if input == "view" {
			log.Print("ProcessMessage(): bot invoked with \"view\" command")

			list := list.ReadList()

			if len(list) < 1 {
				b.SendMessage("The list is empty")
			} else {
				// Send a new message for each user's own grocery list
				for user, list := range list {
					b.SendMessage(user + ": " + strings.Join(list, ", "))
				}
			}
		} else if input[:5] == "clear" {
			log.Print("ProcessMessage(): bot invoked with \"clear\" command")

			list.ClearList()
			b.SendMessage("List cleared")
		} else if input[:3] == "add" {
			log.Print("ProcessMessage(): bot invoked with \"add\" command")

			//TODO: Separate/split the input at commas
			list.WriteList(groupmeMessage.Name, []string{input[4:]})
			b.SendMessage("Added: " + input[4:])
		} else {
			log.Print("ProcessMessage(): command/input not recognized.")

			b.SendMessage(fmt.Sprintf("Command/input not recognized: %s", groupmeMessage.Text))
		}
	}
}
