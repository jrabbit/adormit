package adormit

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/gotk3/gotk3/glib"
	"os/exec"
	"reflect"
	"time"
)

type Alarm struct {
	name   string
	early  bool
	time   time.Time
	active bool
	id     string
}

type Timer struct {
	name     string
	start    time.Time
	duration time.Duration
	end      time.Time
	command  string
	args     []string
}

func (t Timer) Countdown() {
	timer1 := time.NewTimer(t.duration)
	<-timer1.C
	fmt.Println("Timer completed")
	t.end = time.Now()
	cmd := exec.Command(t.command, t.args...)
	cmd.Run()
}

func DemoTimer() {
	args := []string{"-i", "clock", "Timer over!", "adormit"}
	t := Timer{duration: time.Second * 1, command: "notify-send", args: args}
	t.Countdown()
}

func MakeAlarm() {
	alarm := Alarm{name: "my-alarm", active: true}
	SetAlarm(alarm)
}

func SetAlarm(alarm Alarm) {
	var insert []map[string]dbus.Variant
	a := make(map[string]dbus.Variant)
	a["name"] = dbus.MakeVariant(alarm.name)
	a["active"] = dbus.MakeVariant(alarm.active)
	existing_alarms := GetGnomeAlarms()
	insert = append(insert, a)
	for _, v := range existing_alarms {
		insert = append(insert, v)
	}
	// insert = append(insert, get_v().([]interface{})[0].(map[string]interface{}))
	sig, _ := dbus.ParseSignature("aa{sv}")
	existing_alarms_var := dbus.MakeVariantWithSignature(insert, sig)
	fmt.Println(existing_alarms_var)
}

func debug(ty interface{}) {
	fooType := reflect.TypeOf(ty)
	fmt.Println(fooType)
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}
}

func get_v() interface{} {
	settings := glib.SettingsNew("org.gnome.clocks")
	alarms := settings.GetValue("alarms")
	sig, _ := dbus.ParseSignature("aa{sv}")
	v, _ := dbus.ParseVariant(alarms.String(), sig)
	return v.Value()
}

func GetGnomeAlarms() []map[string]dbus.Variant {
	//gsettings get org.gnome.clocks alarms
	settings := glib.SettingsNew("org.gnome.clocks")
	alarms := settings.GetValue("alarms")
	// debug(settings)
	sig, _ := dbus.ParseSignature("aa{sv}")
	v, _ := dbus.ParseVariant(alarms.String(), sig)
	val := v.Value()
	alarmMaps := val.([]map[string]dbus.Variant)
	fmt.Println(alarmMaps)
	return alarmMaps
}
