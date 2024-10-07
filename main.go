package main

import (
	"flag"
	tgClient "github.com/donscuffi/tg_bot1/clients/telegram"
	event_consumer "github.com/donscuffi/tg_bot1/consumer/event-consumer"
	"github.com/donscuffi/tg_bot1/events/telegram"
	"github.com/donscuffi/tg_bot1/storage/files"
	"log"
)

const (
	tgBotHost   = "https://api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Println("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
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
