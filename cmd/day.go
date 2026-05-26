package cmd

import (
	"fmt"
	"shinyschedule/scheduler"

	"github.com/spf13/cobra"
)

var dayCmd = &cobra.Command{
	Use:   "day <weekday>",
	Short: "Show schedule for a specific day",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day := args[0]
		defaults, _ := scheduler.LoadDefaults("defaults.json")
		state, _ := scheduler.LoadState("extratasks.json")
		tasks := scheduler.GetDay(day, defaults, state)
		fmt.Println(scheduler.RenderDay(day, tasks))
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)
}
