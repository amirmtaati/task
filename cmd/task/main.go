package main

import (
	"fmt"
	"github.com/amirmtaati/task/internal/cli"
)

func main() {
	app := cli.NewApp("~/.todo.txt")



	fmt.Println("calling List()")
	tasks := app.List()
	for _, task := range tasks {
		println(task.Raw)
	}

	fmt.Println("Hello World")
}
