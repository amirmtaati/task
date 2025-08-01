package main

import (
	"github.com/amirmtaati/task/internal/cli"
)

func main() {
	app := cli.NewApp("~/.todo.txt")
	app.List()
}
