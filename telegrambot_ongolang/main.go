package main

/*
Написать функционирующего телеграмм бота на языке GO
имеющего минимум 3 команды
- Git возвращает адрес вашего репозитория
- Tasks Возвращает нумерованный списов ваших выполенных заданий
- Task# где # номер задания, возвращает ссылку на папку в вашем репозитории с выполенной задачей
*/
import (
	"fmt"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strings"
	"time"
)

type Data struct {
	APITOKEN string
}

func msgFunc(lnmsg string) {
	fmt.Println(time.Now(), lnmsg)
}
func main() {
	var configdata Data

	conf := viper.New()
	conf.SetConfigName("conf")  // name of config file (without extension)
	conf.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	conf.AddConfigPath("files") // path to look for the config file in
	conf.AddConfigPath(".")     // optionally look for config in the working directory
	err := conf.ReadInConfig()  // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error files/conf.yaml file: %w \n", err))
	}
	err = conf.Unmarshal(&configdata)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	APITOKEN := configdata.APITOKEN
	//APITOKEN := strings.ToLower("sss")
	teststring := " Tasks asd"
	match, _ := regexp.MatchString("tasks", strings.ToLower(teststring))
	fmt.Println(match)
	msgFunc("Starting bot")
	b, err := tb.NewBot(tb.Settings{
		Token:  APITOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		msgFunc("Bot starting error")
		return
	} else {
		msgFunc("Bot started")
	}
	b.Handle("/test", func(m *tb.Message) {

		b.Send(m.Sender, "test")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		msgFunc(m.Text)
		b.Send(m.Sender, "I  make help info from this, but not now")
	})

	b.Start()
}
