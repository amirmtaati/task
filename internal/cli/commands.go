package cli

import (
	"fmt"
	"strings"
)

func ListAction(args []string, a *App) {
	tasks := a.taskList.GetTasks()
	for _, task := range tasks {
		fmt.Println(task.Todo)
	}
}

func AddTaskHandler(args []string, a *App) error {
	task, err := a.parser.Parse(strings.Join(args, " "))
	if err != nil {
		return err
	}

	a.taskList.AddTask(task)

	if err := a.SaveTasks(); err != nil {
		return err
	}

	return nil
}
