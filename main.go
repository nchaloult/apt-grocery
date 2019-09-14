package main

import (
	"os"

	"github.com/nchaloult/apt-grocery/bot"

	// Automatically load in env vars from .env file
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	botID := os.Getenv("BOT_ID")

	bot := bot.NewBot(botID)
	bot.SendMessage("first message sent through Bot struct abstraction")
	bot.SendMessage("Don't mind me, just yeeting my shit through here")
}
