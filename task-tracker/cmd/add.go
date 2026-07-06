package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/problem-setter/task-tracker/internal/response"
)

func AddTask(s string) error {
	fmt.Println("add.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	// for a, b := range task {
	// 	fmt.Println(a, "->", b)
	// }

	id := 1
	if len(task) != 0 {
		id = task[len(task)-1].ID + 1
	}

	newTask := response.Task{
		ID:          id,
		Description: s,
		Status:      response.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task = append(task, newTask)

	data, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	if err := os.WriteFile("internal/database/data.json", data, 0644); err != nil {
		return fmt.Errorf("failed to add new task: %w", err)
	}

	// fmt.Println(string(data))
	return nil
}
