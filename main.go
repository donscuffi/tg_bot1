package main

import (
	"flag"
	"log"
)

const (
	tgBotHost = "https://api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	//
	//fetcher = fetcher.New(tgClient)
	//
	//processor = processor.New(tgClient)
	//
	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	// bot -tg-bot-token 'my token'
	token := flag.String( // *
		"token-bot-token",
		"",
		"token to access to tg bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
