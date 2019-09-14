package bot

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

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
