package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title     string
	Completed bool
}

var tasks []Task

const dataFile = "tasks.json"

func main() {
	if err := loadTasksFromFile(); err != nil {
		fmt.Println("Gagal memuat tasks:", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Selamat datang di To-Do List! 📝")

	for {
		fmt.Println("\n===== To-Do List =====")
		fmt.Println("1. Lihat semua task")
		fmt.Println("2. Tambah task")
		fmt.Println("3. Tandai task selesai")
		fmt.Println("4. Hapus task")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showTasks()
		case "2":
			addTask(scanner)
		case "3":
			completeTask(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Terima kasih! 👋")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada task.")
		return
	}
	for i, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan judul task: ")
	scanner.Scan()
	title := scanner.Text()

	if title == "" {
		fmt.Println("Task tidak boleh kosong.")
		return
	}

	tasks = append(tasks, Task{Title: title})
	saveTasksToFile()
	fmt.Println("Task ditambahkan.")
}

func completeTask(scanner *bufio.Scanner) {
	showTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Masukkan nomor task yang selesai: ")
	scanner.Scan()
	numStr := scanner.Text()
	index, err := strconv.Atoi(numStr)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	tasks[index-1].Completed = true
	saveTasksToFile()
	fmt.Println("Task ditandai sebagai selesai.")
}

func deleteTask(scanner *bufio.Scanner) {
	showTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Masukkan nomor task yang ingin dihapus: ")
	scanner.Scan()
	numStr := scanner.Text()
	index, err := strconv.Atoi(numStr)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	saveTasksToFile()
	fmt.Println("Task dihapus.")
}

func saveTasksToFile() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func loadTasksFromFile() error {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&tasks)
}
