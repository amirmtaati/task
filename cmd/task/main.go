package main

import (
	"fmt"
	"os"
	"github.com/amirmtaati/task/internal/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Init(); err != nil {
		fmt.Printf("Error initializing app: %v\n", err)
		os.Exit(1)
	}
	app.Run()
}
