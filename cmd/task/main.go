package main

import (
	"fmt"
	

	"os"

	"github.com/amirmtaati/task/internal/cli"
	"flag"
)

func main() {
	var todoPath string
	flag.StringVar(&todoPath, "file", "", "Path to todo.txt file")
	flag.StringVar(&todoPath, "f", "", "Path to todo.txt file (shorthand)")
	flag.Parse()
	args := flag.Args()

	app := cli.NewApp()
	if err := app.Init(todoPath); err != nil {
		fmt.Printf("Error initializing app: %v\n", err)
		os.Exit(1)
	}

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
				fmt.Printf("error while adding task: %v", err)
			}
		},
	})

	app.Register(cli.Command{
		Name: "done",
		Action: func(args []string, app *cli.App) {
			if err := cli.CompleteTaskHandler(args, app); err != nil {
				fmt.Printf("error while completing task: %v", err)
			}
		},
	})

	app.Register(cli.Command{
		Name: "delete",
		Action: func(args []string, app *cli.App) {
			if err := cli.DeleteTaskHandler(args, app); err != nil {
				fmt.Printf("error while deleting task: %v", err)
			}
		},
	})

	app.Run(args)
}
