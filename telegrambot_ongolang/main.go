package main

/*
Написать функционирующего телеграмм бота на языке GO
имеющего минимум 3 команды
- Git возвращает адрес вашего репозитория /git
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
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"), "|", lnmsg)
}

func main() {
	//vars start
	var configdata Data
	var rgxpGit = regexp.MustCompile(`^\/{0,1}git$`)
	var rgxpGitclose = regexp.MustCompile(`^.*git.*$`)
	var rgxpTasks = regexp.MustCompile(`^\/{0,1}tasks$`)
	var rgxpTasksclose = regexp.MustCompile(`^.*tasks.*$`)
	var rgxpTaskn = regexp.MustCompile(`^\/{0,1}task\s*\d{1,2}$`)
	var rgxpTasknclose = regexp.MustCompile(`^.*task.*$`)
	var rgxpHelp = regexp.MustCompile(`^\/{0,1}help$`)
	var rgxpHelpclose = regexp.MustCompile(`^.*\/{0,1}.*help.*$`)
	//vars end | other static params start
	r, _ := regexp.Compile("[0-9]{1,2}")
	//other static params end | viper load start
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
	//viper end load | bot load start
	APITOKEN := configdata.APITOKEN

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
	//bot load end | main script start
	b.Handle(tb.OnText, func(m *tb.Message) {
		userMsglower := strings.ToLower(m.Text)
		userMsg := strings.TrimSpace(userMsglower)
		msgFunc(userMsg)
		switch {
		case rgxpTasks.MatchString(userMsg):
			b.Send(m.Sender, "ok u want see tasks")
		case rgxpTasksclose.MatchString(userMsg):
			b.Send(m.Sender, "ok it looks like tasks now but i'm not sure")
		case rgxpTaskn.MatchString(userMsg):
			b.Send(m.Sender, "ok task with number")
			tasknum := r.FindString(userMsg)
			b.Send(m.Sender, "ok task with number ("+tasknum+")")
		case rgxpTasknclose.MatchString(userMsg):
			b.Send(m.Sender, "ok just task")
		case rgxpGit.MatchString(userMsg):
			b.Send(m.Sender, "github.com/gogolevjuri/DevOps-s-course-Sep-Oct-2021")
		case rgxpGit.MatchString(userMsg):
			b.Send(m.Sender, "github.com/gogolevjuri/DevOps-s-course-Sep-Oct-2021")
		case rgxpHelp.MatchString(userMsg):
			b.Send(m.Sender, "its help \n sss")
		case rgxpHelpclose.MatchString(userMsg):
			b.Send(m.Sender, "If u searching help, write /help")
		case rgxpGitclose.MatchString(userMsg):
			b.Send(m.Sender, "I'm find \"git\" in your msg, if you want check my git wirte /git\nFull list of comands u can get using /help")
		default:
			b.Send(m.Sender, "I  make help info from this, but not now")
		}
	})
	//main script end
	b.Start()
}
