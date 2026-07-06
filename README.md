# Task Tracker

A solution to the [Task Tracker project on roadmap.sh](https://roadmap.sh/projects/task-tracker).

Task Tracker is a simple command-line task manager written in Go. It lets you
add, update, delete, list, and mark tasks from the terminal while storing the
data locally in a JSON file.

## Features

- Add tasks with a description.
- Update an existing task description.
- Delete tasks by ID.
- Mark tasks as `in-progress` or `done`.
- List all tasks or filter by status.
- Persist tasks in `internal/database/data.json`.

## Getting Started

Move into the project directory:

```bash
cd task-tracker
```

Run the help command:

```bash
go run . help
```

You should see the available commands and their expected arguments.

## Usage

Add a new task:

```bash
go run . add "Write project documentation"
```

Update a task description:

```bash
go run . update 1 "Write complete project documentation"
```

Delete a task:

```bash
go run . delete 1
```

Mark a task as in progress:

```bash
go run . mark-in-progress 1
```

Mark a task as done:

```bash
go run . mark-done 1
```

List all tasks:

```bash
go run . list
```

List tasks by status:

```bash
go run . list todo
go run . list in-progress
go run . list done
```

## Build

Build a local binary:

```bash
go build -o task-cli .
```

Run the compiled CLI:

```bash
./task-cli help
./task-cli add "Review pull request"
./task-cli list
```

On Windows, run the generated `task-cli.exe` instead.

## Data Storage

Tasks are stored in:

```text
task-tracker/internal/database/data.json
```

Each task contains:

- `id`: numeric task ID.
- `description`: task text.
- `status`: one of `todo`, `in-progress`, or `done`; new tasks start as
  `todo`.
- `createdAt`: creation timestamp.
- `updatedAt`: last update timestamp.

Run the CLI from the `task-tracker` directory so it can find the JSON data file
using the expected relative path.

## Project Structure

```text
task-tracker/
|-- cmd/
|   |-- add.go
|   |-- delete.go
|   |-- list.go
|   |-- mark.go
|   `-- update.go
|-- internal/
|   |-- database/
|   |   `-- data.json
|   |-- response/
|   |   `-- json.go
|   `-- validator.go
|-- go.mod
`-- main.go
```

## Notes

- Task IDs are generated from the current task list length.
- Deleting a task rewrites the task list and renumbers the remaining IDs.
- Invalid commands return a short error message and suggest using
  `./task-cli help`.
