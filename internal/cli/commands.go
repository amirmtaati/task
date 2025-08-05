package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func ListHandler(args []string, a *App) error {
	tasks := a.taskList.GetTasks()
	for _, task := range tasks {
		if task.Raw != "" {
			fmt.Printf("[%d] %s\n", task.ID, task.Todo)
			fmt.Printf("%t", task.Done)
		}
	}
	return nil
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

func FilterHandler(args []string, a *App) error {
	filteredTasks := a.taskList.Filter(args)
	for _, task := range filteredTasks {
		if task.Raw != "" {
			fmt.Printf("[%d] %s\n", task.ID, task.Todo)
		}
	}
	return nil
}

func HelpHandler(args []string, a *App) error {
	help := `Task - A Simple CLI Task Manager for todo.txt
Usage:
  task add <task>
  task done <task_id> 
  task delete <task_id>
  task list
  task -f=/path/to/todo.txt ...
  task --file=/path/to/todo.txt ...

Options:
-f --file Set the location of your task file`

	fmt.Println(help)
	return nil
}
