# Go Expense Approval API

Sebuah RESTful API untuk sistem persetujuan pengeluaran (expense approval) bertahap menggunakan Go (Golang), Fiber, dan PostgreSQL.

## 📦 Fitur Utama

- ✅ Autentikasi dengan JWT (Login)
- 👥 Manajemen Pengguna (CRUD User)
- 👨‍💼 Manajemen Approver dan Approval Stage
- 💸 Pengajuan Expense dan Proses Approval Bertahap
- 📄 Validasi menggunakan DTO
- 🔐 Middleware JWT
- 🧪 Unit Testing menggunakan Testify

---

## 📁 Struktur Proyek

├── cmd/
├── config/
├── internal/
│ ├── dto/
│ ├── handlers/
│ ├── middleware/
│ ├── models/
│ ├── repository/
│ ├── routes/
│ └── services/
├── test/
│ └── mocks/
├── go.mod
├── go.sum
└── main.go

---

## 🚀 Instalasi

### 1. Clone Repo

````bash
git clone https://github.com/username/go-expense-approval-api.git
cd go-expense-approval-api

2. Jalankan PostgreSQL
Pastikan PostgreSQL berjalan dan tersedia. Buat database:

```bash
CREATE DATABASE expense_approval;

3. Konfigurasi .env
Buat file .env (opsional jika hardcoded) dan atur kredensial database:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=expense_approval
JWT_SECRET=supersecret

4. Jalankan Migrasi Otomatis & Server
```bash
go run main.go

API tersedia di http://localhost:3000/api

5. Menjalankan Testing

```bash
go test internal/handlers/user_handler_test.go

📬 API Endpoint
🔐 Auth
POST /api/login – Login dan dapatkan JWT

👤 User
GET /api/users
POST /api/users

PUT /api/users/:id

PATCH /api/users/:id

DELETE /api/users/:id

DELETE /api/users/:id/hard (hard delete)

🧑‍💼 Approver
GET /api/approvers

POST /api/approvers

PUT /api/approvers/:id

DELETE /api/approvers/:id

🪜 Approval Stage
GET /api/approval-stages

POST /api/approval-stages

PUT /api/approval-stages/:id

DELETE /api/approval-stages/:id

💸 Expense
GET /api/expenses

POST /api/expenses

PUT /api/expenses/:id

DELETE /api/expenses/:id

✅ Approval
POST /api/expenses/:id/approve

🛡️ Middleware
Middleware JWT dipasang di semua route yang membutuhkan autentikasi.

Authorization: Bearer <token> harus dikirim di setiap request yang dilindungi.


🧰 Tools & Library
Fiber – Web Framework

GORM – ORM

JWT – Autentikasi

Testify – Unit Testing
````
