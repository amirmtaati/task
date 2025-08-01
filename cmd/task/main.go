package main

import (
	"fmt"
	"github.com/amirmtaati/task/internal/cli"
)

func main() {
	// Initialize the CLI app
	app := cli.NewApp("~/.todo.txt")

	fmt.Println("Task Manager initialized")
	fmt.Println("\nListing sample task...")

	app.Register(cli.Command{
		Name: "list",
		Action: cli.ListAction(app),
	})

	app.Run()
}
