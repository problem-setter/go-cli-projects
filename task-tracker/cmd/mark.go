package cmd

import (
	"fmt"
	"time"

	"github.com/problem-setter/task-tracker/internal/response"
)

func MarkTask(id int, status response.Status) error {
	// fmt.Println("mark.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	idx := SearchTask(id, task)
	if idx == -1 {
		return fmt.Errorf("no task found.")
	}
	task[idx].Status = status
	task[idx].UpdatedAt = time.Now()

	if err := response.UpdateJSON(task, response.FilePath); err != nil {
		return fmt.Errorf("failed to mark task: %w", err)
	}

	fmt.Printf("task with id: %d was successfully marked as %v.\n", id, status)
	return nil
}
