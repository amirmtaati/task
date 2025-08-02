package task

import (
	"strings"

	"github.com/amirmtaati/task/internal/models"
)

const NewLine = "\n"

type TaskList struct {
	Tasks   []models.Task
	storage models.TaskStorage // FIXED: Should be pointer
}

func NewTaskList(storage models.TaskStorage) *TaskList {
	return &TaskList{
		Tasks:   []models.Task{}, // FIXED: Initialize empty slice
		storage: storage,         // FIXED: Use parameter name, not variable name
	}
}

func (tl *TaskList) AddTask(task *models.Task) {
	newTask := models.Task{Todo: task.Todo,
		Raw:      task.Raw,
		Priority: task.Priority,
		Projects: task.Projects,
		Contexts: task.Contexts}

	tl.Tasks = append(tl.Tasks, newTask) // REMOVED the * dereference
}

func (tl *TaskList) String() string {
	var sb strings.Builder
	for _, task := range tl.Tasks {
		sb.WriteString(task.String())
		sb.WriteString(NewLine)
	}
	return sb.String()
}

func (tl *TaskList) LoadTasks() error {
	lines, err := tl.storage.Load()
	if err != nil {
		return err
	}

	for i, line := range lines {
		task := models.Task{ID: i, Raw: line} // This creates a value (not pointer)
		tl.Tasks = append(tl.Tasks, task)     // REMOVED the * dereference
	}
	return nil
}

func (tl *TaskList) GetTasks() []models.Task {
	return tl.Tasks
}
