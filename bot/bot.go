package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	if groupmeMessage.Name != "apt grocery" {
		b.SendMessage(fmt.Sprintf("Repeating what you said: %s", groupmeMessage.Text))
	}
}
