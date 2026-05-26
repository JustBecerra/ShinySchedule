package scheduler

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Activity string `json:"activity"`
}

type WeekTasks struct {
	Tasks []Task `json:"tasks"`
	Day   string `json:"day"`
}

// Defaults is keyed by weekday name, e.g. "monday"
type Defaults map[string][]Task

// State holds the extra tasks and when the current week started
type State struct {
	WeekStart string            `json:"week_start"`
	Extras    map[string][]Task `json:"extras"`
}

func LoadDefaults(path string) (Defaults, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var d Defaults
	return d, json.Unmarshal(data, &d)
}

func LoadState(path string) (State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// If the file doesn't exist yet, return a fresh state
		if os.IsNotExist(err) {
			return State{
				WeekStart: time.Now().Format("2006-01-02"),
				Extras:    make(map[string][]Task),
			}, nil
		}
		return State{}, err
	}
	var s State
	return s, json.Unmarshal(data, &s)
}

func SaveState(path string, s State) error {
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// CheckReset wipes Extras if 7 days have passed since WeekStart
func CheckReset(s *State) {
	weekStart, err := time.Parse("2006-01-02", s.WeekStart)
	if err != nil || time.Since(weekStart) >= 7*24*time.Hour {
		s.WeekStart = time.Now().Format("2006-01-02")
		s.Extras = make(map[string][]Task)
	}
}

// GetDay merges defaults and extras for a given weekday name
func GetDay(day string, defaults Defaults, state State) []Task {
	var tasks []Task
	tasks = append(tasks, defaults[day]...)
	tasks = append(tasks, state.Extras[day]...)
	return tasks
}

// GetWeek merges defaults and extras for all weekdays
func GetWeek(days []string, defaults Defaults, state State) []WeekTasks {
	var tasks []WeekTasks
	for _, day := range days {
		tasks = append(tasks, WeekTasks{
			Tasks: GetDay(day, defaults, state),
			Day:   day,
		})
	}
	return tasks
}
