package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func ListAction(args []string, a *App) {
	tasks := a.taskList.GetTasks()
	for _, task := range tasks {
		if task.Raw != "" {
			fmt.Printf("[%d] %s\n", task.ID, task.Todo)
		}
	}
}

func AddTaskHandler(args []string, a *App) error {
	rawText := strings.Join(args, " ")
	return a.taskList.AddTaskFromRaw(rawText, a.parser)
}

func CompleteTaskHandler(args []string, a *App) error {
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	a.taskList.CompleteTask(taskId)
	return nil
}

func DeleteTaskHandler(args []string, a *App) error {
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	a.taskList.DeleteTask(taskId)
	return nil
}


