package cli

import "fmt"

func (a *App) List() {
	tasks := a.taskList.GetTasks()

	for _, task := range tasks {
		fmt.Println(task.Raw)
	}
}
