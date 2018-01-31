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

func (alarm Alarm) SetAlarm() {
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

type Timer struct {
	Name     string
	start    time.Time
	Duration time.Duration
	end      time.Time
	Command  string
	Args     []string
}

func (t Timer) Countdown() {
	timer1 := time.NewTimer(t.Duration)
	<-timer1.C
	fmt.Println("Timer completed")
	t.end = time.Now()
	cmd := exec.Command(t.Command, t.Args...)
	cmd.Run()
}

func DemoTimer() {
	args := []string{"-i", "clock", "Timer over!", "adormit"}
	t := Timer{Duration: time.Second * 1, Command: "notify-send", Args: args}
	t.Countdown()
}

func MakeAlarm() {
	alarm := Alarm{name: "my-alarm", active: true}
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

// func get_v() interface{} {
// 	settings := glib.SettingsNew("org.gnome.clocks")
// 	alarms := settings.GetValue("alarms")
// 	sig, _ := dbus.ParseSignature("aa{sv}")
// 	v, _ := dbus.ParseVariant(alarms.String(), sig)
// 	return v.Value()
// }

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
