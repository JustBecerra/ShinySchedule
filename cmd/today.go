package cmd

import (
	"fmt"
	"shinyschedule/scheduler"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's schedule",
	Run: func(cmd *cobra.Command, args []string) {
		day := strings.ToLower(time.Now().Weekday().String())
		defaults, _ := scheduler.LoadDefaults("defaults.json")
		state, _ := scheduler.LoadState("extratasks.json")
		tasks := scheduler.GetDay(day, defaults, state)
		fmt.Println(scheduler.RenderDay(day, tasks))
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
