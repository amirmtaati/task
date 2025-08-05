package models

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

type TaskStorage interface {
	Save([]*Task) error
	Load() ([]string, error)
}

func NewTask(raw string) *Task {
	return &Task{
		Raw:          raw,
		CreationDate: time.Now(),
	}
}

func GenerateID() int64 {
	return time.Now().UnixNano()
}

func (t *Task) String() string {
	if t.Done {
		return "x " + t.Raw
	}
	return t.Raw
}
