package task

import (
	"github.com/amirmtaati/task/internal/models"
)

type Filter func(task *models.Task) bool

func (tl *TaskList) Filter(inpFilters []string) []*models.Task {
	filteredTasks := make([]*models.Task, 0)
	AllFilters := getFilters()

	for _, task := range tl.tasks {
		for _, filter := range inpFilters {
			if AllFilters[filter](task) {
				filteredTasks = append(filteredTasks, task)
				break
			}
		}
	}

	return filteredTasks
}

func FilterDone(t *models.Task) bool {
	return t.Done
}

func getFilters() map[string]Filter {
	return map[string]Filter {
		"done": FilterDone,
	}
}
