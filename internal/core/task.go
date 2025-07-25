package task

import (
	"time"
)

type Task struct {
	ID             int
	Raw            string
	Done           bool
	Priority       string
	CompletionDate time.Time
	CreationDate   time.Time
	DueDate        time.Time
	Todo           string
	Projects       []string
	Contexts       []string
	Tags           map[string]string
}

func NewTask() *Task {
	return &Task{
		CreationDate: time.Now(),
	}
}

func (t *Task) String() string {
		return t.Raw()
}



