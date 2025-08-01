package cli

import (
"github.com/amirmtaati/task/internal/models"
)

func (a *App) List() []models.Task {
	return a.taskList.GetTasks()
}
