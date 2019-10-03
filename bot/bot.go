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

type groupmeCallbackBody struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Bot struct {
	botID string
}

func NewBot(botID string) *Bot {
	return &Bot{botID: botID}
}

func (b *Bot) SendMessage(msg string) {
	reqBodyAsStr := fmt.Sprintf(`{"text": "%s", "bot_id": "%s"}`, msg, b.botID)
	reqBodyAsJSONBytes := []byte(reqBodyAsStr)

	req, err := http.NewRequest("POST", "https://api.groupme.com/v3/bots/post", bytes.NewBuffer(reqBodyAsJSONBytes))
	if err != nil {
		// TODO: better error handling
		log.Fatal("SendMessage(): Failed to create new POST request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("SendMessage(): Failed to send POST request")
	}
	defer res.Body.Close()

	log.Printf("Response status code: %v\n", res.Status)
}

// ProcessMessage is called when a new message is posted in our GroupMe group
func (b *Bot) ProcessMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupmeMessage := groupmeCallbackBody{}
	err := json.NewDecoder(r.Body).Decode(&groupmeMessage)
	if err != nil {
		http.Error(w, "Failed to parse GroupMe request to callback: ", http.StatusInternalServerError)
		return
	}

	if groupmeMessage.Name != "apt grocery" && strings.HasPrefix(groupmeMessage.Text, ".gl") {
		input := strings.TrimSpace(groupmeMessage.Text)[4:]

		if input == "view" {
			list := list.ReadList()

			if len(list) < 1 {
				b.SendMessage("The list is empty")
			} else {
				for user, list := range list {
					b.SendMessage(user + ": " + strings.Join(list, ", "))
				}
			}
		} else if input[:5] == "clear" {
			list.ClearList()
			b.SendMessage("List cleared")
		} else if input[:3] == "add" {
			//TODO: Separate/split the input at commas
			list.WriteList(groupmeMessage.Name, []string{input[4:]})
			b.SendMessage("Added: " + input[4:])
		} else {
			b.SendMessage(fmt.Sprintf("Repeating what you said: %s", groupmeMessage.Text))
		}
	}
}
