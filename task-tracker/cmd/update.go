package cmd

import (
	"fmt"
	"time"

	"github.com/problem-setter/task-tracker/internal/response"
)

func SearchTask(id int, task []response.Task) int {
	idx := 0
	for _, tasks := range task {
		if tasks.ID == id {
			return idx
		}
		idx++
	}
	return -1
}

func UpdateTask(id int, s string) error {
	fmt.Println("update.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	idx := SearchTask(id, task)
	if idx == -1 {
		return fmt.Errorf("no task found.")
	}
	task[idx].Description = s
	task[idx].UpdatedAt = time.Now()

	if err := response.UpdateJSON(task, response.FilePath); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}
