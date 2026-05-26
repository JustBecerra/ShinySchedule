package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help for a command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available commands:")
		fmt.Println("  today - Show today's schedule")
		fmt.Println("  week - Show the full week schedule")
		fmt.Println("  day <day> - Show the schedule for a specific day")
		fmt.Println("  add - Add a task to a day this week")
		fmt.Println("  remove <day> <index> - Remove a task from a day this week")
		fmt.Println("  help - Show help for a command")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
