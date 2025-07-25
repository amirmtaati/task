package task

import (
	"github.com/amirmtaati/tempus/internal/storage/file"
	"strings"
)

const NewLine = "\n"

type TaskList struct {
	Tasks   []Task
	storage file.Storage
}

func NewTaskList(storage *file.Storage) *TaskList {
	var Tasks []Task
	return &TaskList{
		Tasks:   Tasks,
		storage: storage,
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

func (tl *TaskList) loadTasks() {
		for task, err := file.Storage.Load(); err == nil {
				tl.tasks = append(tl.tasks, task)
		}
}

func (tl *TaskList) GetTasks() []Task {
		return tl.Tasks
}
