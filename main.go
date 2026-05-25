package main

import (
	"shinyschedule/cmd"
	"shinyschedule/scheduler"
)

func main() {
	state, _ := scheduler.LoadState("extratasks.json")
	scheduler.CheckReset(&state)
	scheduler.SaveState("extratasks.json", state)
	cmd.Execute()
}
