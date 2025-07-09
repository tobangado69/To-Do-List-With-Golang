# To-Do List With Golang

A simple command-line To-Do List application written in Go. Manage your daily tasks easily from your terminal!

## Features

- View all tasks
- Add new tasks
- Mark tasks as completed
- Delete tasks
- Data is automatically saved in the `tasks.json` file

## How to Run

1. **Clone this repository**
    ```bash
    git clone https://github.com/yourusername/To-Do-List-With-Golang.git
    cd To-Do-List-With-Golang
    ```
2. **Build the application**
    ```bash
    go build -o todo
    ```
3. **Run the application**
    ```bash
    ./todo
    ```

## Usage

- To view all tasks:
  ```bash
  ./todo list
  ```
- To add a new task:
  ```bash
  ./todo add "Your new task"
  ```
- To mark a task as completed:
  ```bash
  ./todo done <task_id>
  ```
- To delete a task:
  ```bash
  ./todo delete <task_id>
  ```

## Example

```bash
$ ./todo add "Buy groceries"
Task added: Buy groceries

$ ./todo list
1. [ ] Buy groceries

$ ./todo done 1
Task 1 marked as completed.

$ ./todo list
1. [x] Buy groceries

$ ./todo delete 1
Task 1 deleted.
```