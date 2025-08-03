package task

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/amirmtaati/task/internal/models"
	"github.com/amirmtaati/task/internal/parser"
)

const NewLine = "\n"

type TaskList struct {
	tasks   map[int64]*models.Task
	storage models.TaskStorage
	nextID  int
}

func NewTaskList(storage models.TaskStorage) *TaskList {
	return &TaskList{
		tasks:   map[int64]*models.Task{},
		storage: storage,
		nextID:  1,
	}
}

func (tl *TaskList) AddTaskFromRaw(rawText string, p *parser.Parser) error {
	task := models.NewTask(rawText)
	//task.ID = tl.nextID

	if err := p.PopulateTask(task, rawText); err != nil {
		return err
	}

	//tl.tasks = append(tl.tasks, *task)
	tl.tasks[task.ID] = task
	tl.nextID++

	return tl.save()
}

func (tl *TaskList) CompleteTask(id int) error {
	if id < 1 || id > len(tl.tasks) {
		return fmt.Errorf("task %d not found", id)
	}

	tl.tasks[int64(id)].Done = true
	return tl.save()
}

func (tl *TaskList) DeleteTask(id int) error {
	if id < 1 || id > len(tl.tasks) {
		return fmt.Errorf("task %d not found", id)
	}

	delete(tl.tasks, int64(id))
	return tl.save()
}

// func (tl *TaskList) AddTask(task *models.Task) error {
// 	if task.ID == 0 {
// 		task.ID = tl.nextID
// 		tl.nextID++
// 	}
// 	//tl.tasks = append(tl.tasks, *task)
// 	tl.tasks[task.ID] = task

// 	return tl.save()
// }

func (tl *TaskList) GetTasks() []*models.Task {
	return slices.Collect(maps.Values(tl.tasks))
}

func (tl *TaskList) LoadFromStorage(p *parser.Parser) error {
	lines, err := tl.storage.Load()
	if err != nil {
		return err
	}

	tl.tasks = map[int64]*models.Task{}
	tl.nextID = 1

	for _, line := range lines {
		task := models.NewTask(line)
		//task.ID = tl.nextID

		if err := p.PopulateTask(task, line); err != nil {
			return err
		}

		tl.tasks[task.ID] = task
		tl.nextID++
	}

	return nil
}

func (tl *TaskList) save() error {
	return tl.storage.Save(tl.GetTasks())
}

func (tl *TaskList) String() string {
	var sb strings.Builder
	for _, task := range tl.tasks {
		sb.WriteString(task.String())
		sb.WriteString(NewLine)
	}
	return sb.String()
}
