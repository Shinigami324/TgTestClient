package eventconsumer

import (
	"log"
	"tgBot/TG-BOT/events"
	"time"
)

type Consumer struct {
	fetcher   events.Featche
	processor events.Processor
	batchSize int
}

func New(fetcher events.Featche, process events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: process,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Featch(c.batchSize)

		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}
	return nil
}
