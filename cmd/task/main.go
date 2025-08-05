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

	app.Run(args)
}
