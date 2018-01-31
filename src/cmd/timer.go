package cmd

import (
	"adormit"
	"github.com/spf13/cobra"
	"time"
)

var Length int

func init() {
	timerCmd.PersistentFlags().IntVarP(&Length, "length", "l", 60, "length of time in seconds")
	RootCmd.AddCommand(timerCmd)
}

var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "run a simple timer",
	Long:  `runs a timer optionally with `,
	Run: func(cmd *cobra.Command, args []string) {
		alarmArgs := []string{"-i", "clock", "Timer over!", "adormit"}
		t := adormit.Timer{Duration: time.Second * time.Duration(Length), Command: "notify-send", Args: alarmArgs}
		t.Countdown()
	},
}
