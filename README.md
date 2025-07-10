# To-Do List With Golang

A simple interactive command-line To-Do List application written in Go. Manage your daily tasks easily from your terminal!

## Features

- View all tasks (with colored status)
- Add new tasks
- Mark tasks as completed
- Delete tasks
- Edit task titles
- Filter and display only incomplete tasks
- Tasks are automatically saved in the `tasks.json` file

## How to Run

1. **Clone this repository**
  ```bash
  git clone https://github.com/tobangado69/To-Do-List-With-Golang.git
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

Run the program and follow the interactive menu:

```
===== To-Do List =====
1. Lihat semua task
2. Tambah task
3. Tandai task selesai
4. Hapus task
5. Edit task
6. Tampilkan task yang belum selesai
7. Keluar
Pilih menu (1-7):
```

### Menu Descriptions

- **1. Lihat semua task**: Show all tasks with their status (✅ for completed, ❌ for incomplete). Tasks are sorted so incomplete tasks appear first.
- **2. Tambah task**: Add a new task by entering its title.
- **3. Tandai task selesai**: Mark a task as completed by entering its number.
- **4. Hapus task**: Delete a task by entering its number.
- **5. Edit task**: Edit the title of an existing task.
- **6. Tampilkan task yang belum selesai**: Show only tasks that are not yet completed.
- **7. Keluar**: Exit the application.

## Example

```text
===== To-Do List =====
1. Lihat semua task
2. Tambah task
3. Tandai task selesai
4. Hapus task
5. Edit task
6. Tampilkan task yang belum selesai
7. Keluar
Pilih menu (1-7): 2
Masukkan judul task: Buy groceries
Task ditambahkan.

Pilih menu (1-7): 1
1. [❌] Buy groceries

Pilih menu (1-7): 3
Masukkan nomor task yang selesai: 1
Task ditandai sebagai selesai.

Pilih menu (1-7): 1
1. [✅] Buy groceries

Pilih menu (1-7): 5
Masukkan nomor task yang ingin diedit: 1
Masukkan judul baru: Buy groceries and milk
Task berhasil diedit.

Pilih menu (1-7): 6
Task yang belum selesai:
Semua task sudah selesai.

Pilih menu (1-7): 4
Masukkan nomor task yang ingin dihapus: 1
Task dihapus.
```

---

**Note:**  
- The application uses colored output for task status (❌ in red, ✅ in green) if your terminal supports ANSI colors.
- All data is stored locally in `tasks.json` in the same directory as the executable.
- The menu and prompts are in Indonesian.
- Requires Go 1.18 or newer.