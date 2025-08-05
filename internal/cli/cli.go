package cli

import (
	"fmt"
	"os"
	

	"github.com/amirmtaati/task/internal/core/task"
	"github.com/amirmtaati/task/internal/parser"
	"github.com/amirmtaati/task/internal/storage/file"
	"path/filepath"
	"flag"
)

type CommandHandler func(args []string, a *App) error

type Command struct {
	Name   string
	Action func(args []string, a *App)
}

type App struct {
	commands []Command
	taskList *task.TaskList
	parser   *parser.Parser
	storage  *file.Storage
	path string
	args []string
}

func NewApp() *App {
	parser := parser.NewParser()
	return &App{
		parser: parser,
	}
}

func (a *App) Init() error {
	var path string
	flag.StringVar(&path, "file", "", "Path to todo.txt file")
	flag.StringVar(&path, "f", "", "Path to todo.txt file (shorthand)")
	flag.Parse()
	args := flag.Args()

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
	a.args = args
	a.path = todoPath

	return a.taskList.LoadFromStorage(a.parser)
}

func (a *App) Register(cmd Command) {
	a.commands = append(a.commands, cmd)
}

func (a *App) loadCommands() {
	commands := a.getHandlers()

	for command, handler := range commands {
		a.Register(Command{
			Name: command,
			Action: func(args []string, app *App) {
				if err := handler(args, app); err != nil {
					fmt.Printf("error while running command: %v", err)
				}
			},
		})
	}
}

func (a *App) Run() {
	if len(os.Args) < 2 {
		HelpHandler([]string{}, a)
		return
	}

	inputCmd := a.args[0]
	args := a.args[1:]

	a.loadCommands()

	for _, cmd := range a.commands {
		if cmd.Name == inputCmd {
			cmd.Action(args, a)
			return
		}
	}

	fmt.Printf("Unknown command: %s\n", inputCmd)
}

func (a *App) getHandlers() map[string]CommandHandler {
	return map[string]CommandHandler{
		"list":   ListHandler,
		"add":    AddTaskHandler,
		"done":   CompleteTaskHandler,
		"delete": DeleteTaskHandler,
		"help": HelpHandler,
	}
}
