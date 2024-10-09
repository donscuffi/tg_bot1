package main

import (
	"context"
	"flag"
	"log"
	tgClient "tg_bot1/clients/telegram"
	event_consumer "tg_bot1/consumer/event-consumer"
	"tg_bot1/events/telegram"
	"tg_bot1/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//s:=files.New(sqliteStoragePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
		"tg-bot-token",
		"",
		"token to access to tg bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
