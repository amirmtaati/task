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

func HelpHandler(args []string, a *App) error {
	help := `Naval Fate.
Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

	fmt.Println(help)
	return nil
}
