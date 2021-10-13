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
	var rgxpTaskcc = regexp.MustCompile(`^\/{0,1}task#{0,1}$`)
	var rgxpTasknclose = regexp.MustCompile(`^.*task.*$`)
	var rgxpHelp = regexp.MustCompile(`^\/{0,1}help$`)
	var rgxpHelpclose = regexp.MustCompile(`^.*\/{0,1}.*help.*$`)
	var rgxpStart = regexp.MustCompile(`^\/{1}start$`)
	//vars end | other static params start
	r, _ := regexp.Compile("[0-9]{1,2}")
	//other static params end | viper load start
	conf := viper.New()
	conf.SetConfigName("conf")  // name of config file (without extension)
	conf.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	conf.AddConfigPath("files") // path to look for the config file in
	conf.MergeInConfig()
	conf.SetConfigName("tasks") // name of config file (without extension)
	conf.SetConfigType("yaml")  // REQUIRED if the config file does not have the extension in the name
	conf.AddConfigPath("files") // path to look for the config file in
	conf.MergeInConfig()
	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
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
		if !m.Private() {
			msgFunc("error not private msg")
			return
		}
		userMsglower := strings.ToLower(m.Text)
		userMsg := strings.TrimSpace(userMsglower)
		msgFunc(userMsg)
		switch {
		case rgxpTasks.MatchString(userMsg):
			b.Send(m.Sender, "ok u want see tasks")
		case rgxpTasksclose.MatchString(userMsg):
			b.Send(m.Sender, "I'm find \"tasks\" in your msg, if you want check list of tasks, write /tasks\n"+
				"Full list of comands u can get using /help")
		case rgxpTaskn.MatchString(userMsg):
			tasknum := r.FindString(userMsg)
			b.Send(m.Sender, "ok task with number ("+tasknum+")")
		case rgxpTaskcc.MatchString(userMsg):
			b.Send(m.Sender, "U forget specify the number. Example /task1 .")
		case rgxpTasknclose.MatchString(userMsg):
			b.Send(m.Sender, "I'm find \"task\" in your msg, if you want check tasks write /task# where # is nubmer of the task, example /task1\n"+
				"If u want check full list of tasks, write /tasks"+
				"Full list of comands u can get using /help")
		case rgxpGit.MatchString(userMsg):
			b.Send(m.Sender, "github.com/gogolevjuri/DevOps-s-course-Sep-Oct-2021")
		case rgxpGitclose.MatchString(userMsg):
			b.Send(m.Sender, "I'm find \"git\" in your msg, if you want check my git wirte /git\nFull list of comands u can get using /help")
		case rgxpHelp.MatchString(userMsg):
			b.Send(m.Sender, "Hey, there list of comands, hope u find what you've been looking for\n"+
				"/help - help comand, u are here now\n"+
				"/git - show git repository\n"+
				"/tasks - show list of all tasks\n"+
				"/task# - return info about selected task; where # is nubmer of the task, example /task1\n\n"+
				"created by Juri Gogolev")
		case rgxpHelpclose.MatchString(userMsg):
			b.Send(m.Sender, "If u searching help, write /help")
		case rgxpStart.MatchString(userMsg):
			b.Send(m.Sender, "Hi, u probably new here...this is comands what i accept\n"+
				"/help - help comand\n"+
				"/git - show git repository\n"+
				"/tasks - show list of all tasks\n"+
				"/task# - return info about selected task; where # is nubmer of the task, example /task1\n\n",
			)
		default:
			b.Send(m.Sender, "Please use comands from /help.")
		}
	})
	//main script end
	b.Start()
}
