package cmd

import (
	"fmt"

	"github.com/problem-setter/task-tracker/internal/response"
)

func DeleteTask(id int) error {
	fmt.Println("delete.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	newId := 1
	var taskTemp []response.Task

	isFound := false
	for _, tasks := range task {
		if tasks.ID != id {
			taskTemp = append(taskTemp, response.Task{
				ID:          newId,
				Description: tasks.Description,
				Status:      tasks.Status,
				CreatedAt:   tasks.CreatedAt,
				UpdatedAt:   tasks.UpdatedAt,
			})
			newId++
		} else {
			isFound = true
		}
	}

	if !isFound {
		return fmt.Errorf("no task found.")
	}

	if err := response.UpdateJSON(taskTemp, response.FilePath); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
