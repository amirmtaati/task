package cli

import (
	"fmt"
	"os"
	"github.com/amirmtaati/task/internal/core/task"
	"github.com/amirmtaati/task/internal/parser"
	"github.com/amirmtaati/task/internal/storage/file"
)

type Command struct {
	Name   string
	Action func(args []string)
}

type App struct {
	commands []Command  // FIXED: Field name should match usage
	taskList *task.TaskList
	parser   *parser.Parser
	storage  *file.Storage
}

func NewApp(path string) *App {
	storage := file.NewFile(path)
	taskList := task.NewTaskList(storage)
	parser := parser.NewParser()

	return &App{
		taskList: taskList,
		parser:   parser,
		storage:  storage,
	}
}

func (a *App) Register(cmd Command) {
	a.commands = append(a.commands, cmd)  // FIXED: Use lowercase field name
}

func (a *App) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command>")
		return
	}

	inputCmd := os.Args[1]
	args := os.Args[2:]

	for _, cmd := range a.commands {  // FIXED: Use lowercase field name
		if cmd.Name == inputCmd {
			cmd.Action(args)
			return
		}
	}
	
	fmt.Printf("Unknown command: %s\n", inputCmd)
}
