package main

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	_ = telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String(
		"bot_token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
