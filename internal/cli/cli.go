package cli

import (
	"fmt"
	"os"

	"path/filepath"
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

func NewApp() *App {
	parser := parser.NewParser()
	return &App{
		parser:   parser,
	}
}

func (a *App) Init(path string) error {
	var todoPath string
	if path != "" {
		todoPath = path
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting home directory: %v", err)
		}
		todoPath = filepath.Join(home, ".todo.txt")
	}

	a.storage = file.NewFile(todoPath)
	a.taskList = task.NewTaskList(a.storage)
	
	return a.taskList.LoadFromStorage(a.parser)
}

func (a *App) Register(cmd Command) {
	a.commands = append(a.commands, cmd)
}

func (a *App) Run(inpArgs []string) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command>")
		return
	}

	inputCmd := inpArgs[0]
	args := inpArgs[1:]

	for _, cmd := range a.commands {
		if cmd.Name == inputCmd {
			cmd.Action(args, a)
			return
		}
	}

	fmt.Printf("Unknown command: %s\n", inputCmd)
}
