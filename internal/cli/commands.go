package cli

import (
	"fmt"
)

func ListAction(a *App) {
	tasks := a.taskList.GetTasks()
	for _, task := range tasks {
		fmt.Println(task)

	}
}
