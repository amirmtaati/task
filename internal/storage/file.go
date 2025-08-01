package file

import (
	"bufio"
	"os"

	"github.com/amirmtaati/task/internal/core/task"
)

type FileHandler interface {
	Load() ([]string, error)
	Save(tasks []task.Task) error
}

type Storage struct {
	path string
}

func NewFile(path string) *Storage {
	return &Storage{
		path: path,
	}
}

func (f *Storage) loadFile() (*os.File, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *Storage) Load() ([]string, error) {
	var lines []string
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, err
}

func (f *Storage) Save(tasks []task.Task) error {
	file, err := f.loadFile()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		if _, err := writer.WriteString(task.String()); err != nil {
			return err
		}
	}
	return nil
}
