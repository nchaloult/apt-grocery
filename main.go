package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nchaloult/apt-grocery/bot"

	// Automatically load in env vars from .env file
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	botID := os.Getenv("BOT_ID")
	bot := bot.NewBot(botID)

	router := httprouter.New()
	router.POST("/", bot.ProcessMessage)

	log.Fatal(http.ListenAndServe(":5000", router))

	fmt.Println("Yeeeeeeeeeeeeeeeeeeeeeet")
}
