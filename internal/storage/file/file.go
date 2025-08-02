package file

import (
	"bufio"
	"os"

	"github.com/amirmtaati/task/internal/models"
)

type Storage struct {
	path string
}

func NewFile(path string) *Storage {
	return &Storage{
		path: path,
	}
}

func (f *Storage) Load() ([]string, error) {
	var lines []string
	file, err := os.Open(f.path)
	if err != nil {
		return lines, err // Return empty slice if file doesn't exist
	}
	defer file.Close() // ADDED: Always close file

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err() // FIXED: Return scanner error, not open error
}

func (f *Storage) Save(tasks []models.Task) error {
	// FIXED: Open file for writing, not reading
	file, err := os.OpenFile(f.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // ADDED: Ensure data is written

	for _, task := range tasks {
		if _, err := writer.WriteString(task.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}
