package adormit

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/gotk3/gotk3/glib"
	"os/exec"
	"reflect"
	"time"
)

func Version() string {
	return "0.0.1a1"
}

var CurrentTimers []Timer
var CurrentAlarms map[string]Alarm

type Alarm struct {
	Name   string
	Early  bool
	Time   time.Time
	Active bool
	Id     string
}

func (alarm Alarm) SetAlarm() string {
	var insert []map[string]dbus.Variant
	a := make(map[string]dbus.Variant)
	a["name"] = dbus.MakeVariant(alarm.Name)
	a["active"] = dbus.MakeVariant(alarm.Active)
	existing_alarms := GetGnomeAlarms()
	insert = append(insert, a)
	for _, v := range existing_alarms {
		insert = append(insert, v)
	}
	sig, _ := dbus.ParseSignature("aa{sv}")
	existing_alarms_var := dbus.MakeVariantWithSignature(insert, sig)
	CurrentAlarms[alarm.Id] = alarm
	fmt.Println(existing_alarms_var)
	return alarm.Id
}

func (alarm Alarm) UnsetAlarm() {

}

type Timer struct {
	Name     string
	start    time.Time
	Duration time.Duration
	end      time.Time
	Command  string
	Args     []string
}

func (t Timer) Countdown() {
	CurrentTimers = append(CurrentTimers, t)
	t.start = time.Now()
	timer1 := time.NewTimer(t.Duration)
	<-timer1.C
	fmt.Println("Timer completed")
	t.end = time.Now()
	cmd := exec.Command(t.Command, t.Args...)
	cmd.Run()
}

func (t Timer) String() string {
	return t.Name
}

func DemoTimer() {
	args := []string{"-i", "clock", "Timer over!", "adormit"}
	t := Timer{Duration: time.Second * 1, Command: "notify-send", Args: args}
	t.Countdown()
}

func MakeAlarm() {
	alarm := Alarm{Name: "my-alarm", Active: true}
	alarm.SetAlarm()
}

func debug(ty interface{}) {
	fooType := reflect.TypeOf(ty)
	fmt.Println(fooType)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}

func GetGnomeAlarms() []map[string]dbus.Variant {
	//gsettings get org.gnome.clocks alarms
	settings := glib.SettingsNew("org.gnome.clocks")
	alarms := settings.GetValue("alarms")
	sig, _ := dbus.ParseSignature("aa{sv}")
	v, _ := dbus.ParseVariant(alarms.String(), sig)
	val := v.Value()
	alarmMaps := val.([]map[string]dbus.Variant)
	fmt.Println(alarmMaps)
	return alarmMaps
}
