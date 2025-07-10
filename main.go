package main

import (
	"bufio"         // Untuk membaca input/output buffered, misal membaca baris dari file atau stdin
	"encoding/json" // Untuk encoding dan decoding data JSON (serialisasi/deserialisasi)
	"fmt"           // Untuk format I/O, seperti mencetak ke layar (Println, Printf, dll)
	"os"            // Untuk operasi sistem seperti membaca argumen, membuka file, keluar program
	"sort"          // Untuk mengurutkan slice, misalnya mengurutkan daftar tugas berdasarkan yang belum selesai
	"strconv"       // Untuk konversi string ke tipe numerik dan sebaliknya
)

// Task merepresentasikan satu item tugas pada to-do list.
// Title adalah judul tugas, Completed menandakan apakah tugas sudah selesai.
type Task struct {
	Title     string
	Completed bool
}

// tasks adalah slice global yang menyimpan seluruh daftar tugas.
var tasks []Task

// dataFile adalah nama file tempat data tugas disimpan dalam format JSON.
const dataFile = "tasks.json"

// main adalah fungsi utama yang menjalankan aplikasi To-Do List berbasis terminal.
// Fungsi ini akan memuat data task dari file, lalu menampilkan menu interaktif
// kepada pengguna untuk melihat, menambah, menandai selesai, atau menghapus task.
// Pengguna dapat memilih menu dengan memasukkan angka 1-5. Program akan terus
// berjalan hingga pengguna memilih untuk keluar dari aplikasi.
func main() {
	if err := loadTasksFromFile(); err != nil {
		fmt.Println("Gagal memuat tasks:", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Selamat datang di To-Do List! 📝")

	for {
		// Menampilkan menu utama
		fmt.Println("\n===== To-Do List =====")
		fmt.Println("1. Lihat semua task")
		fmt.Println("2. Tambah task")
		fmt.Println("3. Tandai task selesai")
		fmt.Println("4. Hapus task")
		fmt.Println("5. Edit task")
		fmt.Println("6. Tampilkan task yang belum selesai")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu (1-7): ")

		scanner.Scan()
		choice := scanner.Text()

		// Memproses pilihan user
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
			editTask(scanner)
		case "6":
			filterIncompleteTasks()
		case "7":
			fmt.Println("Terima kasih! 👋")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}

	}
}

func sortTasks() {
	sort.SliceStable(tasks, func(i, j int) bool {
		// false < true → belum selesai akan muncul lebih dulu
		return !tasks[i].Completed && tasks[j].Completed
	})
}

func editTask(scanner *bufio.Scanner) {
	showTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Masukkan nomor task yang ingin diedit: ")
	scanner.Scan()
	numStr := scanner.Text()
	index, err := strconv.Atoi(numStr)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Print("Masukkan judul baru: ")
	scanner.Scan()
	newTitle := scanner.Text()
	if newTitle == "" {
		fmt.Println("Judul tidak boleh kosong.")
		return
	}

	tasks[index-1].Title = newTitle
	saveTasksToFile()
	fmt.Println("Task berhasil diedit.")
}

func filterIncompleteTasks() {
	fmt.Println("Task yang belum selesai:")
	found := false
	for i, task := range tasks {
		if !task.Completed {
			found = true
			fmt.Printf("%d. [❌] %s\n", i+1, task.Title)
		}
	}
	if !found {
		fmt.Println("Semua task sudah selesai.")
	}
}

// showTasks menampilkan seluruh daftar tugas beserta statusnya (selesai/belum).
func showTasks() {
	if len(tasks) == 0 {
		fmt.Println("Belum ada task.")
		return
	}

	sortTasks()

	for i, task := range tasks {
		status := "\033[31m❌\033[0m"
		if task.Completed {
			status = "\033[32m✅\033[0m"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
}

// addTask meminta input judul tugas dari user, lalu menambahkannya ke daftar tasks.
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

	// Cek apakah semua task sudah selesai
	allCompleted := true
	for _, task := range tasks {
		if !task.Completed {
			allCompleted = false
			break
		}
	}
	if allCompleted {
		fmt.Println("Task sudah selesai semua.")
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
	if tasks[index-1].Completed {
		fmt.Println("Task ini sudah selesai.")
		return
	}
	tasks[index-1].Completed = true
	saveTasksToFile()
	fmt.Println("Task ditandai sebagai selesai.")
}

// deleteTask meminta user memilih nomor tugas, lalu menghapus tugas tersebut dari daftar.
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

// saveTasksToFile menyimpan seluruh daftar tasks ke file JSON.
func saveTasksToFile() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

// loadTasksFromFile memuat daftar tasks dari file JSON jika file tersebut ada.
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
