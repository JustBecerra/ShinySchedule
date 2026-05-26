package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shinyschedule",
	Short: "ShinySchedule is a CLI tool for seeing my weekly / daily schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to ShinySchedule!")
		fmt.Println("Use 'shinyschedule help' to see available commands")
		fmt.Println("Remember the tasks reset every week except for default tasks you set.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
