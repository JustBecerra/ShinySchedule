package cmd

import (
	"fmt"
	"shinyschedule/scheduler"

	"github.com/spf13/cobra"
)

var days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show the full week schedule",
	Run: func(cmd *cobra.Command, args []string) {
		defaults, _ := scheduler.LoadDefaults("defaults.json")
		state, _ := scheduler.LoadState("extratasks.json")
		tasks := scheduler.GetWeek(days, defaults, state)
		fmt.Println(scheduler.RenderWeek(days, tasks))
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)
}
