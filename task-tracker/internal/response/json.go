package response

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const FilePath = "internal/database/data.json"

func ReadJSON(path string) ([]Task, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer file.Close()

	var taskSlice []Task
	json.NewDecoder(file).Decode(&taskSlice)
	return taskSlice, nil
}

func UpdateJSON(task []Task, path string) error {
	data, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}
	return nil
}
