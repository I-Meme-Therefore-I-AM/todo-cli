# To-Do CLI Application

A simple Command Line Interface (CLI) application for managing tasks directly from your terminal.

## Features
- Perform CRUD operations on tasks via the terminal.
- Manage tasks with the following commands:
  - Add new tasks.
  - List uncompleted tasks or all tasks (including completed ones).
  - Mark tasks as complete.
  - Delete tasks.

## Requirements
- Go installed on your system.
- A JSON or plain text file as the underlying data store for tasks.

---

## Installation
Clone the repository and build the application:

```bash
$ git clone <repository_url>
$ cd todo-cli
$ go build -o tasks
$ sudo mv tasks /usr/local/bin
```

---

## Usage
The CLI provides the following commands to manage your tasks:

### 1. **Add a Task**
The `add` command is used to create new tasks in the underlying data store.

#### Syntax:
```bash
$ tasks add <description>
```

#### Example:
```bash
$ tasks add "Tidy my desk"
```
This adds a new task with the description "Tidy my desk".

---

### 2. **List Tasks**
The `list` command returns a list of uncompleted tasks by default. Use the `-a` or `--all` flag to list all tasks, including completed ones.

#### Syntax:
```bash
$ tasks list
```

#### Example:
```bash
$ tasks list
```
Output:
```
ID    Task                                                Created
1     Tidy up my desk                                     a minute ago
3     Change my keyboard mapping to use escape/control    a few seconds ago
```

#### Show All Tasks:
```bash
$ tasks list -a
```
Output:
```
ID    Task                                                Created          Done
1     Tidy up my desk                                     2 minutes ago    false
2     Write up documentation for new project feature      a minute ago     true
3     Change my keyboard mapping to use escape/control    a minute ago     false
```

---

### 3. **Mark a Task as Complete**
The `complete` command is used to mark a task as done.

#### Syntax:
```bash
$ tasks complete <taskid>
```

#### Example:
```bash
$ tasks complete 1
```
This marks the task with ID 1 as completed.

---

### 4. **Delete a Task**
The `delete` command is used to remove a task from the data store.

#### Syntax:
```bash
$ tasks delete <taskid>
```

#### Example:
```bash
$ tasks delete 3
```
This deletes the task with ID 3.

---

## Example Workflow
```bash
$ tasks add "Buy groceries"
$ tasks add "Read Go documentation"
$ tasks list
ID    Task                          Created
1     Buy groceries                 a few seconds ago
2     Read Go documentation         a few seconds ago

$ tasks complete 1
$ tasks list -a
ID    Task                          Created             Done
1     Buy groceries                 a minute ago        true
2     Read Go documentation         a minute ago        false

$ tasks delete 1
$ tasks list
ID    Task                          Created
2     Read Go documentation         a minute ago
```

---

## Notes
- The tasks are stored in a file in JSON format (or a similar data store).
- Task IDs are automatically assigned and incremented for each new task.
- The timestamps for `Created` are dynamically generated and formatted for readability.

---
