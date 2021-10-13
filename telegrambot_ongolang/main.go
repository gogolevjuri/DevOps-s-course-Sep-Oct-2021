package main

import (
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/spf13/viper"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "https://pkg.go.dev/gopkg.in/tucnak/telebot.v2",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}