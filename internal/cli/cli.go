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
	Action func(args []string, a *App)
}

type App struct {
	commands []Command
	taskList *task.TaskList
	parser   *parser.Parser
	storage  *file.Storage
}

func NewApp(path string) *App {
	storage := file.NewFile(path)
	parser := parser.NewParser()
	taskList := task.NewTaskList(storage)

	return &App{
		taskList: taskList,
		parser:   parser,
		storage:  storage,
	}
}

func (a *App) Init() error {
	return a.taskList.LoadFromStorage(a.parser)
}

func (a *App) Register(cmd Command) {
	a.commands = append(a.commands, cmd)
}

func (a *App) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command>")
		return
	}

	inputCmd := os.Args[1]
	args := os.Args[2:]

	for _, cmd := range a.commands {
		if cmd.Name == inputCmd {
			cmd.Action(args, a)
			return
		}
	}

	fmt.Printf("Unknown command: %s\n", inputCmd)
}
