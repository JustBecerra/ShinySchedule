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

		if start >= end {
			fmt.Println("Error: start time must be before end time")
			return
		}

		defaults, _ := scheduler.LoadDefaults("defaults.json")
		state, _ := scheduler.LoadState("extratasks.json")
		existing := scheduler.GetDay(day, defaults, state)

		newTask := scheduler.Task{
			Start:    start,
			End:      end,
			Activity: activity,
		}
		if conflict, ok := scheduler.FindConflict(existing, newTask); ok {
			fmt.Printf("Error: time conflict with existing task (%s - %s %s)\n",
				conflict.Start, conflict.End, conflict.Activity)
			return
		}

		state.Extras[day] = append(state.Extras[day], newTask)
		scheduler.SaveState("extratasks.json", state)
		fmt.Printf("Added \"%s\" to %s (%s - %s)\n", activity, day, start, end)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
