# To-Do List Dengan Golang

Aplikasi To-Do List interaktif berbasis command-line yang ditulis dengan Go. Kelola tugas harian Anda dengan mudah langsung dari terminal!

## Fitur

- Melihat semua task (dengan status berwarna)
- Menambah task baru
- Menandai task sebagai selesai
- Menghapus task
- Mengedit judul task
- Menampilkan hanya task yang belum selesai
- Task otomatis disimpan di file `tasks.json`

## Cara Menjalankan

1. **Clone repository ini**
  ```bash
  git clone https://github.com/tobangado69/To-Do-List-With-Golang
  cd To-Do-List-With-Golang
  ```
2. **Build aplikasi**
  ```bash
  go build -o todo
  ```
3. **Jalankan aplikasi**
  ```bash
  ./todo
  ```

## Penggunaan

Jalankan program dan ikuti menu interaktif:

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

### Penjelasan Menu

- **1. Lihat semua task**: Menampilkan semua task beserta statusnya (✅ untuk selesai, ❌ untuk belum selesai). Task yang belum selesai akan muncul di atas.
- **2. Tambah task**: Menambah task baru dengan memasukkan judulnya.
- **3. Tandai task selesai**: Menandai task sebagai selesai dengan memasukkan nomornya.
- **4. Hapus task**: Menghapus task dengan memasukkan nomornya.
- **5. Edit task**: Mengedit judul task yang sudah ada.
- **6. Tampilkan task yang belum selesai**: Menampilkan hanya task yang belum selesai.
- **7. Keluar**: Keluar dari aplikasi.

## Contoh

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
Masukkan judul task: Beli kebutuhan dapur
Task ditambahkan.

Pilih menu (1-7): 1
1. [❌] Beli kebutuhan dapur

Pilih menu (1-7): 3
Masukkan nomor task yang selesai: 1
Task ditandai sebagai selesai.

Pilih menu (1-7): 1
1. [✅] Beli kebutuhan dapur

Pilih menu (1-7): 5
Masukkan nomor task yang ingin diedit: 1
Masukkan judul baru: Beli kebutuhan dapur dan susu
Task berhasil diedit.

Pilih menu (1-7): 6
Task yang belum selesai:
Semua task sudah selesai.

Pilih menu (1-7): 4
Masukkan nomor task yang ingin dihapus: 1
Task dihapus.
```

---

**Catatan:**  
- Aplikasi ini menggunakan warna pada status task (❌ merah, ✅ hijau) jika terminal Anda mendukung ANSI colors.
- Semua data disimpan secara lokal di `tasks.json` pada direktori yang sama dengan file executable.
- Menu dan prompt menggunakan bahasa Indonesia.
- Membutuhkan Go versi 1.18 atau lebih baru.