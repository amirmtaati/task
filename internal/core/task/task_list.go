package task

import (
	"github.com/amirmtaati/task/internal/models"
	"strings"
)

const NewLine = "\n"

type TaskList struct {
	Tasks   []models.Task
	storage models.TaskStorage  // FIXED: Should be pointer
}

func NewTaskList(storage models.TaskStorage) *TaskList {
	return &TaskList{
		Tasks:   []models.Task{},  // FIXED: Initialize empty slice
		storage: storage,   // FIXED: Use parameter name, not variable name
	}
}

func (tl *TaskList) String() string {
	var sb strings.Builder 
	for _, task := range tl.Tasks {
		sb.WriteString(task.String())
		sb.WriteString(NewLine)
	}
	return sb.String()
}

// FIXED: This method had multiple compilation errors
func (tl *TaskList) loadTasks() error {
		lines, err := tl.storage.Load()
			if err != nil {
						return err
							}
								
								for _, line := range lines {
											task := models.Task{Raw: line}  // This creates a value (not pointer)
													tl.Tasks = append(tl.Tasks, task)  // REMOVED the * dereference
														}
															return nil
														}

func (tl *TaskList) GetTasks() []models.Task {
	return tl.Tasks
}
