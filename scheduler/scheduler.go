package scheduler

import (
	"encoding/json"
	"os"
	"sort"
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

func SaveDefaults(path string, d Defaults) error {
	data, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
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

// GetDay merges defaults and extras for a given weekday name, sorted by start time.
func GetDay(day string, defaults Defaults, state State) []Task {
	var tasks []Task
	tasks = append(tasks, defaults[day]...)
	tasks = append(tasks, state.Extras[day]...)
	sortTasksByStart(tasks)
	return tasks
}

func sortTasksByStart(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Start < tasks[j].Start
	})
}

func tasksEqual(a, b Task) bool {
	return a.Start == b.Start && a.End == b.End && a.Activity == b.Activity
}

// FindTaskIndex returns the index of target in tasks, or -1 if not found.
func FindTaskIndex(tasks []Task, target Task) int {
	for i, t := range tasks {
		if tasksEqual(t, target) {
			return i
		}
	}
	return -1
}

// TasksOverlap reports whether two tasks occupy overlapping time ranges.
func TasksOverlap(a, b Task) bool {
	if a.Start == "" || a.End == "" || b.Start == "" || b.End == "" {
		return false
	}
	return a.Start < b.End && b.Start < a.End
}

// FindConflict returns the first existing task that overlaps with newTask.
func FindConflict(tasks []Task, newTask Task) (Task, bool) {
	for _, t := range tasks {
		if TasksOverlap(t, newTask) {
			return t, true
		}
	}
	return Task{}, false
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
