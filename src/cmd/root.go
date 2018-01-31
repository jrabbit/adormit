package cmd

import (
	"adormit"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "adormit-cli",
	Short: "Your personal time agent",
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.Help()
		adormit.GetGnomeAlarms()
		adormit.MakeAlarm()
		adormit.DemoTimer()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	RootCmd.Execute()
}

func initConfig() {
	return
}
