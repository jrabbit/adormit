package cmd

import (
	"adormit"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version of adormit",
	Long:  `Output the version of adormit `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(adormit.Version())
	},
}
