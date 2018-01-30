package adormit

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Alarm struct {
	name  string
	early bool
	time  time.Time
}

// func MakeAlarm() {
// 	alarm := Alarm{name:"my-alarm"}

// }

func GetGnomeAlarms() {
	//gsettings get org.gnome.clocks  alarms
	cmd := exec.Command("gsettings", "get", "org.gnome.clocks", "alarms")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
