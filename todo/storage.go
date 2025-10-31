package todo

import (
	"encoding/json"
	"os"
)

const filename = "tasks.json"

func LoadTask() ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var task []Task
	if err := json.Unmarshal(file, &task); err != nil {
		return nil, err
	}

	return task, nil
}

func SaveTasks(task []Task) error {
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}