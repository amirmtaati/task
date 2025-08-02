package cli

import (
	"fmt"
	"strings"
)

func ListAction(args []string, a *App) {
	tasks := a.taskList.GetTasks()
	for _, task := range tasks {
		fmt.Printf("[%d] %s\n", task.ID, task.Todo)
	}
}

// MUCH SIMPLER - just delegate to TaskList
func AddTaskHandler(args []string, a *App) error {
	rawText := strings.Join(args, " ")
	return a.taskList.AddTaskFromRaw(rawText, a.parser)
}
