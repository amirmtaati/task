package main

import (
	"fmt"

	"os"
	"path/filepath"

	"github.com/amirmtaati/task/internal/cli"
)

func main() {
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	todoPath := filepath.Join(home, ".todo.txt")

	app := cli.NewApp(todoPath)
	app.Init()
	//args := os.Args[2:]

	app.Register(cli.Command{
		Name: "list",
		Action: func(args []string, app *cli.App) {
			cli.ListAction(args, app)
		},
	})

	app.Register(cli.Command{
		Name: "add",
		Action: func(args []string, app *cli.App) {
			if err := cli.AddTaskHandler(args, app); err != nil {
				fmt.Errorf("error while adding task: %d", err)
			}
		},
	})

	app.Register(cli.Command{
		Name: "done",
		Action: func(args []string, app *cli.App) {
			if err := cli.CompleteTaskHandler(args, app); err != nil {
				fmt.Errorf("error while completing task: %d", err)
			}
		},
	})

	app.Register(cli.Command{
		Name: "delete",
		Action: func(args []string, app *cli.App) {
			if err := cli.DeleteTaskHandler(args, app); err != nil {
				fmt.Errorf("error while deleting task: %d", err)
			}
		},
	})

	app.Run()
}
