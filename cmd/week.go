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
		for _, day := range days {
			tasks := scheduler.GetDay(day, defaults, state)
			fmt.Println(scheduler.RenderDay(day, tasks))
		}
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)
}
