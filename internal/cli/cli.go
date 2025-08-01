package cli

import (
	"github.com/amirmtaati/task/internal/core/task"
	"github.com/amirmtaati/task/internal/parser"
	"github.com/amirmtaati/task/internal/storage/file"
	"os"
)

type Command struct {
	Name   string
	Action func(args []string)
}

type App struct {
	commands []Command
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
	a.Commands = append(a.Commands, cmd)
}

func (a *App) Run() {
	inputCmd := os.Args[1]
	args := os.Args[2:]

	for _, cmd := range a.Commands {
		if cmd.Name == inputCmd {
			cmd.Action(args)
			return
		}
	}

}
