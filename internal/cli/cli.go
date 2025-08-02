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
	commands []Command // FIXED: Field name should match usage
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

func (a *App) Init() error {
	//a.taskList.LoadTasks()
	lines, err := a.storage.Load()
	if err != nil {
		return err
	}

	for _, line := range lines {
		task, err := a.parser.Parse(line)
		if err != nil {
			return err
		}
		a.taskList.AddTask(task)
	}

	return nil
}

func (a *App) SaveTasks() error {
	return a.storage.Save(a.taskList.GetTasks())
}

func (a *App) Register(cmd Command) {
	a.commands = append(a.commands, cmd) // FIXED: Use lowercase field name
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
