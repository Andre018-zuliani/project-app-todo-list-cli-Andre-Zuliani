# 📝 To-Do List CLI App (Golang + Cobra)

Aplikasi **To-Do List berbasis Command Line Interface (CLI)** menggunakan **Golang** dan **Cobra**.  
Aplikasi ini memungkinkan pengguna mengelola daftar tugas secara lokal menggunakan file JSON.

---

## 📌 Fitur Utama

- ✔ Tambah tugas baru (`add`)
- ✔ Tampilkan daftar tugas (`list`)
- ✔ Tandai tugas selesai (`done`)
- ✔ Hapus tugas (`delete`)
- ✔ Pencarian tugas berdasarkan keyword (`search`)
- ✔ Penyimpanan data menggunakan JSON
- ✔ Tabel output rapi dengan border ASCII
- ✔ Warna ANSI untuk status & prioritas
- ✔ Validasi input lengkap
- ✔ Struktur folder terorganisir & idiomatik Go

---

## 📂 Struktur Project
📁 project-app-todo-list-cli-nama/
├── 📂 cmd/
│   ├── 📝 add.go
│   ├── 📝 list.go
│   ├── 📝 delete.go
│   ├── 📝 done.go
│   ├── 📝 search.go
│   └── 📝 root.go
├── 📂 data/
│   └── 📄 todo.json
├── 📂 model/
│   └── 🧩 task.go
├── 📂 service/
│   └── ⚙️ todo_service.go
├── 📂 utils/
│   └── 🔧 validation.go  (optional)
├── 📄 go.mod
├── 📄 go.sum
└── 🚀 main.go



## 🚀 Cara Install & Setup

### 1. Clone repository
```bash
git clone https://github.com/USERNAME/project-app-todo-list-cli-nama.git
cd project-app-todo-list-cli-nama```

2. Install dependency Cobra
go get -u github.com/spf13/cobra
go get -u github.com/fatih/color

3. Jalankan aplikasi
go run .

🛠 Daftar Perintah CLI
➕ Tambah tugas
go run . add --title "Belajar Go" --desc "Materi struct dan interface" --priority high


Atau versi singkat:

go run . add "Belajar Go"

# 📋 Tampilkan daftar tugas
go run . list


Contoh output:

+----+----------------------+----------------+------------+
| ID | Task                 | Status         | Priority   |
+----+----------------------+----------------+------------+
| 1  | Belajar Go           | pending        | high       |
| 2  | Makan siang          | completed      | medium     |
+----+----------------------+----------------+------------+

✔ Tandai tugas selesai
go run . done 1

❌ Hapus tugas
go run . delete 2

🔍 Cari tugas berdasarkan keyword
go run . search belajar

🧠 Teknologi yang Digunakan

🟦 Golang

⚡ Cobra CLI Framework

🗂 JSON storage

🎨 ANSI Terminal Colors (fatih/color)

📦 JSON Data Structure
{
  "id": 1,
  "title": "Belajar Go",
  "description": "Materi struct dan interface",
  "status": "pending",
  "priority": "high"
}

🛡 Validasi yang Diterapkan

Judul tugas tidak boleh kosong

Judul tidak boleh duplikat

ID harus numerik saat delete/done

Priority default: low

Status default: pending

🎯 Tujuan Project

Project ini dibuat sebagai mini project untuk memenuhi ketentuan:

Penggunaan operator dan logika

Pemakaian variabel

Minimal 3 fungsi

Slice dan array

Layout output tabel dengan fmt.Printf

Error handling

JSON processing

File handling

CLI menggunakan Cobra dan flags

Fitur search

Validasi input

Struktur folder idiomatik Go

Semua ketentuan telah dipenuhi ✔

🏷 Lisensi

Proyek ini dirilis menggunakan lisensi MIT License.

🙌 Kontribusi

Pull Request, bug report, dan saran sangat diterima.

👤 Author

Andre Zuliani
To-Do List CLI Project – Golang
