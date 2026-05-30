package cmd

import (
	"bufio"
	"fmt"
	"os"
	"shinyschedule/scheduler"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <day> <index>",
	Short: "remove a task from a day this week",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		day, index := strings.ToLower(args[0]), args[1]
		indexInt, err := strconv.Atoi(index)
		if err != nil {
			fmt.Printf("Invalid index: %s\n", index)
			return
		}
		defaults, _ := scheduler.LoadDefaults("defaults.json")
		state, _ := scheduler.LoadState("extratasks.json")
		merged := scheduler.GetDay(day, defaults, state)

		if indexInt < 1 || indexInt > len(merged) {
			fmt.Printf("Index %d out of range, %s has %d tasks\n", indexInt, day, len(merged))
			return
		}

		task := merged[indexInt-1]

		if defaultIdx := scheduler.FindTaskIndex(defaults[day], task); defaultIdx >= 0 {
			fmt.Printf("Task %d is a default: %s - %s %s\n", indexInt, task.Start, task.End, task.Activity)
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

			defaults[day] = append(defaults[day][:defaultIdx], defaults[day][defaultIdx+1:]...)
			if err := scheduler.SaveDefaults("defaults.json", defaults); err != nil {
				fmt.Printf("Error saving defaults: %v\n", err)
				return
			}
			fmt.Printf("Removed default task %d from %s\n", indexInt, day)
			return
		}

		extrasIdx := scheduler.FindTaskIndex(state.Extras[day], task)
		if extrasIdx < 0 {
			fmt.Println("Task not found.")
			return
		}
		tasks := state.Extras[day]
		state.Extras[day] = append(tasks[:extrasIdx], tasks[extrasIdx+1:]...)
		scheduler.SaveState("extratasks.json", state)
		fmt.Printf("Removed task %d from %s\n", indexInt, day)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
