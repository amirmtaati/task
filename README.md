# Task - A Simple Todo.txt CLI Tool

A lightweight command-line task management tool that follows the [todo.txt format](http://todotxt.org/).

## Installation

```bash
go install github.com/amirmtaati/task/cmd/task@latest
```

## Usage

### Basic Commands

```bash
# List all tasks
task list

# Add a new task
task add "Buy groceries"
task add "Call mom @phone +personal"

# Mark a task as complete
task done 1

# Delete a task
task delete 2
```

### Custom File Location

By default, tasks are stored in `~/.todo.txt`. You can specify a different file:

```bash
# Use a custom todo file
task -f /path/to/my/todos.txt list
task --file /path/to/work.txt add "Finish project"
```

## Todo.txt Format Support

This tool supports the standard todo.txt format including:

- **Priority**: `(A) High priority task`
- **Projects**: `+project`
- **Contexts**: `@context`
- **Tags**: `key:value`
- **Dates**: Creation and completion dates
- **Completion**: `x` prefix for completed tasks

### Examples

```bash
task add "(A) Call dentist @phone due:2024-12-15"
task add "Review code +work @computer"
task add "Buy milk @store +personal"
```

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `list` | Show all tasks | `task list` |
| `add` | Add a new task | `task add "Task description"` |
| `done` | Mark task as complete | `task done 1` |
| `delete` | Delete a task | `task delete 1` |
