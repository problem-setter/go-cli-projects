package cmd

import (
	"fmt"

	"github.com/problem-setter/task-tracker/internal/response"
)

func TaskList() error {
	// fmt.Println("list.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	if len(task) == 0 {
		return fmt.Errorf("no task found.")
	}

	fmt.Printf("displaying all tasks.\n")
	for idx, task := range task {
		fmt.Printf("\n%3d. %s\n", idx+1, task.Description)
		fmt.Printf("     ID         : %d\n", task.ID)
		fmt.Printf("     Status     : %s\n", task.Status)
		fmt.Printf("     Created At : %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("     Updated At : %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
	return nil
}

func TaskListWithParam(status response.Status) error {
	// fmt.Println("list.go: ping!")

	task, err := response.ReadJSON(response.FilePath)
	if err != nil {
		return fmt.Errorf("failed to read json: %w", err)
	}

	if len(task) == 0 {
		return fmt.Errorf("no task found.")
	}

	fmt.Printf("displaying all tasks marked as %v.\n", status)
	isFound := false
	for idx, task := range task {
		if task.Status == status {
			fmt.Printf("\n%3d. %s\n", idx+1, task.Description)
			fmt.Printf("     ID         : %d\n", task.ID)
			fmt.Printf("     Status     : %s\n", task.Status)
			fmt.Printf("     Created At : %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("     Updated At : %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
			isFound = true
		}
	}

	if !isFound {
		return fmt.Errorf("no task found.")
	}
	return nil
}
