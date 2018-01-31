package adormit

import (
	"fmt"
	"github.com/godbus/dbus"
	"github.com/gotk3/gotk3/glib"
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
}

func MakeAlarm() {
	alarm := Alarm{name: "my-alarm", active: true}
	SetAlarm(alarm)
}

func SetAlarm(alarm Alarm) {
	var insert []map[string]interface{}
	// insert := make(map[string]dbus.Variant)
	a := make(map[string]interface{})
	a["name"] = alarm.name
	a["active"] = alarm.active
	// existing_alarms := GetGnomeAlarms()
	// total_update := make([]map[string]dbus.Variant, 10)
	// total_update = append(total_update, insert)
	// total_update = append(total_update, existing_alarms)
	// total_update[2] = insert
	insert = append(insert, a)
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
