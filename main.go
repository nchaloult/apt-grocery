package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	// Automatically load in env vars from .env file
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	botID := os.Getenv("BOT_ID")

	reqBodyAsStr := fmt.Sprintf(`{"text": "more testing", "bot_id": "%s"}`, botID)
	reqBodyAsJSONBytes := []byte(reqBodyAsStr)

	req, err := http.NewRequest("POST", "https://api.groupme.com/v3/bots/post", bytes.NewBuffer(reqBodyAsJSONBytes))
	if err != nil {
		log.Fatal("Failed to create new POST request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to send POST request")
	}
	defer res.Body.Close()

	fmt.Printf("Response status code: %v\n", res.Status)
}
