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
