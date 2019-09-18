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

	router := httprouter.New()
	router.POST("/", bot.ProcessMessage)

	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

	fmt.Println("Yeeeeeeeeeeeeeeeeeeeeeet")
}
