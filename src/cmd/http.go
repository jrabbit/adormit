package cmd

import (
	"adormit"
	"github.com/spf13/cobra"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
	"strconv"
	"time"
)

func init() {
	RootCmd.AddCommand(httpCmd)
}
func index() string {
	return "welcome to adormit/http version " + adormit.Version()
}
func hc() string {
	return "OK"
}
func version() string {
	return adormit.Version()
}

func newTimer(req *http.Request) string {
	name := req.FormValue("name")
	d, err := strconv.Atoi(req.FormValue("duration"))
	if err != nil {
		panic(err)
	}
	duration := time.Second * time.Duration(d)
	alert, err := strconv.ParseBool(req.FormValue("alert"))
	if err != nil {
		panic(err)
	}
	if alert {
		args := []string{"-i", "clock", "Timer over!", "adormit"}
		t := adormit.Timer{Duration: duration, Command: "notify-send", Args: args}
		go t.Countdown()

	} else {
		t := adormit.Timer{Name: name, Duration: duration}
		go t.Countdown()
	}
	return "OK"
}

func delTimer(req *http.Request) string {
	return "OK"
}

func listTimers(ctx *macaron.Context) {
	ctx.JSON(200, &adormit.CurrentTimers)
	return
}

func newAlarm(req *http.Request, log *log.Logger) string {
	//insert new alarm locally
	name := req.FormValue("name")
	early, err := strconv.ParseBool(req.FormValue("early"))
	// time, err := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", req.FormValue("time"))
	active, err := strconv.ParseBool(req.FormValue("active"))
	if err != nil {
		panic(err)
	}
	// id := adormit.MakeAlarmId()
	a := adormit.Alarm{Active: active, Early: early, Name: name}
	a.MakeId()
	g := a.SetAlarm()
	log.Printf("http/newAlarm: %s", a.Uuid)
	return g
}
func delAlarm(req *http.Request) string {
	// take id remove it from madormit.CurrentAlarms[:i]ap & gnome-clocks
	i := req.FormValue("id")
	alarm := adormit.CurrentAlarms[i]
	delete(adormit.CurrentAlarms, i)
	// adormit.CurrentAlarms = append(adormit.CurrentAlarms[:i], adormit.CurrentAlarms[i+1:]...)
	// balete from gnome-clocks
	err := alarm.UnsetAlarm()
	if err != nil {
		panic(err)
	}
	return "OK"
}
func listAlarms(ctx *macaron.Context) {
	ctx.JSON(200, &adormit.CurrentAlarms)
	return
}

func runServer() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", index)
	m.Group("/meta", func() {
		m.Get("/healthcheck", hc)
		m.Get("/version", version)
	})
	m.Group("/timers", func() {
		m.Post("/new", newTimer)
		m.Delete("/:id", delTimer)
		m.Get("/", listTimers)
	})
	m.Group("/alarms", func() {
		m.Post("/new", newAlarm)
		m.Delete("/:id", delAlarm)
		m.Get("/", listAlarms)
	})
	m.Run()
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "runs a server for networked alarms",
	Long:  `runs a server for networked alarms dabdabdab `,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}
