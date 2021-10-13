package main

import (
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/spf13/viper"
	"fmt"
)

func main() {
    conf := viper.New()
    conf.SetConfigName("conf") // name of config file (without extension)
    conf.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
    conf.AddConfigPath("files")   // path to look for the config file in
    conf.AddConfigPath(".")               // optionally look for config in the working directory
    err := conf.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
    	panic(fmt.Errorf("Fatal error files/conf.yaml file: %w \n", err))
    }

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