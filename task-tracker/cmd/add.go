package cmd

import (
	"fmt"
	"time"

	"github.com/problem-setter/task-tracker/internal/response"
)

func AddTask(s string) error {
	fmt.Println("add.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	// id := 1
	// if len(task) != 0 {
	// 	id = task[len(task)-1].ID + 1
	// }

	newTask := response.Task{
		ID:          len(task) + 1,
		Description: s,
		Status:      response.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	task = append(task, newTask)

	if err := response.UpdateJSON(task, response.FilePath); err != nil {
		return fmt.Errorf("failed to add new task: %w", err)
	}
	return nil
}
