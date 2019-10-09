package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"apt-grocery/bot"

	// Automatically load in env vars from .env file
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	botID := os.Getenv("BOT_ID")
	bot := bot.NewBot(botID)

	// GroupMe sends a POST request to a callback URL, which you specify, when a
	// new message is posted in a group that your bot is in. Listen for those
	// messages here.
	router := httprouter.New()
	router.POST("/", bot.ProcessMessage)

	log.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
