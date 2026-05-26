package scheduler

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RenderDay(day string, tasks []Task) string {
	var sb strings.Builder

	title := cases.Title(language.English).String(day)
	sb.WriteString(fmt.Sprintf("\n%s\n", title))
	sb.WriteString(strings.Repeat("─", 35) + "\n")

	if len(tasks) == 0 {
		sb.WriteString(" No tasks scheduled.\n")
	} else {
		for _, t := range tasks {
			sb.WriteString(fmt.Sprintf(" %s - %s   %s\n", t.Start, t.End, t.Activity))
		}
	}

	sb.WriteString(strings.Repeat("─", 35) + "\n")
	return sb.String()
}

func RenderWeek(days []string, allTasks []WeekTasks) string {
	const colWidth = 26
	const chunkSize = 3
	var sb strings.Builder

	for i := 0; i < len(days); i += chunkSize {
		// Slice the current chunk (may be smaller than chunkSize on last row)
		end := i + chunkSize
		if end > len(days) {
			end = len(days)
		}
		chunkDays := days[i:end]
		chunkTasks := allTasks[i:end]

		// Header row
		for _, day := range chunkDays {
			title := cases.Title(language.English).String(day)
			sb.WriteString(fmt.Sprintf("%-*s", colWidth, title))
		}
		sb.WriteString("\n")

		// Divider row
		for range chunkDays {
			sb.WriteString(strings.Repeat("─", colWidth-1) + " ")
		}
		sb.WriteString("\n")

		// Find max tasks in this chunk
		maxTasks := 0
		for _, wt := range chunkTasks {
			if len(wt.Tasks) > maxTasks {
				maxTasks = len(wt.Tasks)
			}
		}

		// Task rows
		if maxTasks == 0 {
			for range chunkDays {
				sb.WriteString(fmt.Sprintf("%-*s", colWidth, "No tasks scheduled."))
			}
			sb.WriteString("\n")
		} else {
			for row := 0; row < maxTasks; row++ {
				for _, wt := range chunkTasks {
					if row < len(wt.Tasks) {
						t := wt.Tasks[row]
						cell := fmt.Sprintf("%s-%s %s", t.Start, t.End, t.Activity)
						sb.WriteString(fmt.Sprintf("%-*s", colWidth, cell))
					} else {
						sb.WriteString(strings.Repeat(" ", colWidth))
					}
				}
				sb.WriteString("\n")
			}
		}

		sb.WriteString("\n") // blank line between chunk rows
	}

	return sb.String()
}
