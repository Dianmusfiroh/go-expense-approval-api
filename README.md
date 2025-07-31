# go-expense-approval-api

<h1> Tiga Level User </h1>
<h2>1. Admin</h2>

# Fungsi:
   - Mengelola data user & approver
   - Menentukan urutan approval (approval stages)
   - Melihat semua pengajuan & statusnya
# Hak akses:
   - CRUD data user, approver, approval_stage
   - Melihat semua expenses
   - Tidak ikut menyetujui pengeluaran
     
<h2>2. Pemohon (Requester / User Biasa) </h2>

# Fungsi:
   - Mengajukan pengeluaran
   - Melihat status pengajuan mereka
# Hak akses:
  - Buat expense
  - Lihat expense milik sendiri
  - Tidak bisa approve

<h2>3. Approver</h2>

# Fungsi:
  - Menyetujui atau menolak pengeluaran berdasarkan urutan
  - Hanya bisa approve jika sudah waktunya (berdasarkan stage)

# Hak akses:
  - Lihat pengeluaran yang perlu dia approve
  - Kirim persetujuan/penolakan
  - Tidak bisa buat expense atau kelola user lain
