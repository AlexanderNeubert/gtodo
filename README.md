# gtodo

A simple, powerful command-line todo application built with Go. Manage your tasks efficiently with features for creating, updating, listing, and deleting todo items, all from your terminal.

---

## Features

- **Add tasks** with optional categories
- **List tasks** filtered by status (all, complete, incomplete) and/or category
- **Update tasks** (edit content, change category, mark as complete/incomplete)
- **Delete tasks** by ID
- **Persistent storage** in a JSON file in your Documents folder
- **Simple, readable table output** for your todos

---

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/AlexanderNeubert/gtodo.git
   cd gtodo
   ```

2. **Build the application:**
   ```sh
   go build
   ```
---

## Usage

### 1. Initialize Storage

Before using gtodo, initialize the storage file:

```sh
gtodo init
```

This creates a `.todos.json` file in your `~/Documents` directory.

---

### 2. Add a Task

```sh
gtodo add --task "Buy groceries" --cat "Personal"
```
- `--task` (or `-t`): The task description (required)
- `--cat` (or `-c`): The category (optional, defaults to "Uncategorized")

---

### 3. List Tasks

```sh
gtodo list
```

**Options:**
- `--status` (or `-s`): Filter by status (`0` = incomplete, `1` = complete, `2` = all, default)
- `--cat` (or `-c`): Filter by category

**Examples:**
```sh
gtodo list --status 0           # List incomplete tasks
gtodo list --status 1           # List completed tasks
gtodo list --cat "Work"         # List all tasks in "Work" category
gtodo list --status 0 --cat "Personal"
```

---

### 4. Update a Task

```sh
gtodo update --id 3 --task "Buy groceries and milk" --cat "Errands" --done
```
- `--id` (or `-i`): The ID of the task to update (required)
- `--task` (or `-t`): New task description (optional)
- `--cat` (or `-c`): New category (optional)
- `--done` (or `-d`): Mark as complete
- `--undone` (or `-u`): Mark as incomplete

**Note:** You must specify at least one of `--task`, `--cat`, `--done`, or `--undone`.

---

### 5. Delete a Task

```sh
gtodo delete --id 2
```
- `--id` (or `-i`): The ID of the task to delete (required)

---

## Data Storage

- All tasks are stored in a JSON file at:
  `~/Documents/.todos.json`

---

## Example Workflow

```sh
gtodo init
gtodo add --task "Finish Go project" --cat "Work"
gtodo add --task "Read a book"
gtodo list
gtodo update --id 1 --done
gtodo list --status 1
gtodo delete --id 2
```

---

## Development

- Requires Go 1.18 or newer.
- Dependencies are managed via Go modules.
---

## License

[MIT](https://github.com/AlexanderNeubert/gtodo/blob/main/LICENSE)

---

## Author

[Alexander Neubert](https://github.com/AlexanderNeubert)