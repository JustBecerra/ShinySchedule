package cmd

import (
	"fmt"
	"shinyschedule/scheduler"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <day> <start> <end> <activity>",
	Short: "Add a task to a day this week",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		day, start, end, activity := strings.ToLower(args[0]), args[1], args[2], args[3]
		state, _ := scheduler.LoadState("extratasks.json")
		state.Extras[day] = append(state.Extras[day], scheduler.Task{
			Start:    start,
			End:      end,
			Activity: activity,
		})
		scheduler.SaveState("extratasks.json", state)
		fmt.Printf("Added \"%s\" to %s (%s - %s)\n", activity, day, start, end)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
