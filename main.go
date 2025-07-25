package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("مرحبا")
	})

	b.Start()
}
