package validator

import (
	"fmt"

	add "github.com/problem-setter/task-tracker/cmd"
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

func Validation(args []string) {
	if len(args) == 0 || (len(args) == 1 && args[0] == "help") {
		PrintHelp()
		return
	}

	if len(args) > 3 {
		fmt.Println("Failed to execute this command!")
		return
	}

	if len(args) == 2 && args[0] == "add" { // len = 2
		add.AddTask(args[1])
	}

	// fmt.Println(args)
}
