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

func AddTaskHandler(args []string, a *App) {
	line := strings.Join(args, " ")
	task, err := a.parser.Parse(line)
	if err != nil {
		return
	}
	a.taskList.AddTask(task)
}
