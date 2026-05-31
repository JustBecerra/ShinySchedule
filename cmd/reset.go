package cmd

import (
	"bufio"
	"fmt"
	"os"
	"shinyschedule/scheduler"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Remove all tasks from the schedule",
	Run: func(cmd *cobra.Command, args []string) {
		defaults, err := scheduler.LoadDefaults("defaults.json")
		if err != nil {
			fmt.Printf("Error loading defaults: %v\n", err)
			return
		}

		state, err := scheduler.LoadState("extratasks.json")
		if err != nil {
			fmt.Printf("Error loading extra tasks: %v\n", err)
			return
		}

		hasDefaults := scheduler.HasDefaultTasks(defaults)
		hasExtras := scheduler.HasExtras(state)
		if !hasDefaults && !hasExtras {
			fmt.Println("Nothing to reset.")
			return
		}

		if hasDefaults {
			fmt.Println("This will delete all default tasks and extra tasks.")
			fmt.Print("Are you sure? (yes/no): ")

			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error reading input: %v\n", err)
				return
			}
			if strings.TrimSpace(strings.ToLower(answer)) != "yes" {
				fmt.Println("Cancelled.")
				return
			}
		}

		clearedDefaults := scheduler.ClearDefaults(days)
		if err := scheduler.SaveDefaults("defaults.json", clearedDefaults); err != nil {
			fmt.Printf("Error saving defaults: %v\n", err)
			return
		}

		state.WeekStart = time.Now().Format("2006-01-02")
		state.Extras = make(map[string][]scheduler.Task)
		if err := scheduler.SaveState("extratasks.json", state); err != nil {
			fmt.Printf("Error saving extra tasks: %v\n", err)
			return
		}

		fmt.Println("All tasks have been removed.")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
