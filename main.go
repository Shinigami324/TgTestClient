package main

import (
	"flag"
	"log"
	tgClient "tgBot/TG-BOT/clients/telegram"
	eventconsumer "tgBot/TG-BOT/consumer/event-consumer"
	"tgBot/TG-BOT/events/telegram"
	"tgBot/TG-BOT/storage/file"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()),
		file.New(storagePath))

	consumer := eventconsumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("sevice is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token

}
