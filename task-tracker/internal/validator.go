package validator

import (
	"fmt"
	"strconv"

	add "github.com/problem-setter/task-tracker/cmd"
	delete "github.com/problem-setter/task-tracker/cmd"
	list "github.com/problem-setter/task-tracker/cmd"
	mark "github.com/problem-setter/task-tracker/cmd"
	update "github.com/problem-setter/task-tracker/cmd"
	"github.com/problem-setter/task-tracker/internal/response"
)

func PrintHelp() {
	fmt.Println(`Task Tracker, a simple command-line task tracker in Go.
Programmed by https://github.com/problem-setter

usage: ./task-cli <command> [arguments]
available commands:
  add <description>              Add a new task.
  update <id> <description>      Update an existing task description.
  delete <id>                    Delete a task by ID.
  mark-in-progress <id>          Mark a task as in progress.
  mark-done <id>                 Mark a task as done.
  list                           List all tasks.
  list todo                      List all tasks with todo status.
  list in-progress               List all tasks with in-progress status.
  list done                      List all tasks with done status.
  `)
}

func IsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Validation(args []string) error {
	if len(args) == 0 || (len(args) == 1 && args[0] == "help") {
		PrintHelp()
		return nil
	}

	if len(args) > 3 {
		return fmt.Errorf("failed to execute this command.")
	}

	if len(args) == 2 && args[0] == "add" {
		return add.AddTask(args[1])
	}

	if len(args) == 3 && args[0] == "update" && IsNum(args[1]) {
		id, _ := strconv.Atoi(args[1])
		return update.UpdateTask(id, args[2])
	}

	if len(args) == 1 && args[0] == "list" {
		return list.TaskList()
	}

	if len(args) == 2 && args[0] == "delete" && IsNum(args[1]) {
		id, _ := strconv.Atoi(args[1])
		return delete.DeleteTask(id)
	}

	if len(args) == 2 && args[0] == "mark-in-progress" && IsNum(args[1]) {
		id, _ := strconv.Atoi(args[1])
		return mark.MarkTask(id, response.StatusInProgress)
	}

	if len(args) == 2 && args[0] == "mark-done" && IsNum(args[1]) {
		id, _ := strconv.Atoi(args[1])
		return mark.MarkTask(id, response.StatusDone)
	}

	if len(args) == 2 && args[0] == "list" {
		switch args[1] {
		case "todo":
			return list.TaskListWithParam(response.StatusTodo)
		case "in-progress":
			return list.TaskListWithParam(response.StatusInProgress)
		case "done":
			return list.TaskListWithParam(response.StatusDone)
		}
	}

	return fmt.Errorf("unknown command, help usage: ./task-cli help")
}
