package task

import (
	"strings"

	"github.com/amirmtaati/task/internal/models"
	"github.com/amirmtaati/task/internal/parser"
)

const NewLine = "\n"

type TaskList struct {
	tasks   []models.Task
	storage models.TaskStorage
	nextID  int
}

func NewTaskList(storage models.TaskStorage) *TaskList {
	return &TaskList{
		tasks:   []models.Task{},
		storage: storage,
		nextID:  1,
	}
}

// AddTaskFromRaw - This is the main public method for adding tasks
func (tl *TaskList) AddTaskFromRaw(rawText string, p *parser.Parser) error {
	// Create base task
	task := models.NewTask(rawText)
	task.ID = tl.nextID

	// Parse the raw text to populate fields
	if err := p.PopulateTask(task, rawText); err != nil {
		return err
	}

	// Add to in-memory state
	tl.tasks = append(tl.tasks, *task)
	tl.nextID++

	// Auto-save to maintain consistency
	return tl.save()
}

// AddTask - For internal use when task is already parsed
func (tl *TaskList) AddTask(task *models.Task) {
	if task.ID == 0 {
		task.ID = tl.nextID
		tl.nextID++
	}
	tl.tasks = append(tl.tasks, *task)
}

func (tl *TaskList) GetTasks() []models.Task {
	return tl.tasks
}

func (tl *TaskList) LoadFromStorage(p *parser.Parser) error {
	lines, err := tl.storage.Load()
	if err != nil {
		return err
	}

	tl.tasks = []models.Task{} // Clear existing
	tl.nextID = 1

	for _, line := range lines {
		task := models.NewTask(line)
		task.ID = tl.nextID

		if err := p.PopulateTask(task, line); err != nil {
			return err
		}

		tl.tasks = append(tl.tasks, *task)
		tl.nextID++
	}

	return nil
}

func (tl *TaskList) save() error {
	return tl.storage.Save(tl.tasks)
}

func (tl *TaskList) String() string {
	var sb strings.Builder
	for _, task := range tl.tasks {
		sb.WriteString(task.String())
		sb.WriteString(NewLine)
	}
	return sb.String()
}
