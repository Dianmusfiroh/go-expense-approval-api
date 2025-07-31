# go-expense-approval-api

Tiga Level User
1. Admin
# Fungsi:
   -Mengelola data user & approver
   - Menentukan urutan approval (approval stages)
   - Melihat semua pengajuan & statusnya
# Hak akses:
   - CRUD data user, approver, approval_stage
   - Melihat semua expenses
   - Tidak ikut menyetujui pengeluaran
     
2. Pemohon (Requester / User Biasa)
# Fungsi:
  - Mengajukan pengeluaran
  - Melihat status pengajuan mereka
# Hak akses:
  - Buat expense
  - Lihat expense milik sendiri
  - Tidak bisa approve

3. Approver
# Fungsi:
  - Menyetujui atau menolak pengeluaran berdasarkan urutan
  - Hanya bisa approve jika sudah waktunya (berdasarkan stage)

# Hak akses:
  - Lihat pengeluaran yang perlu dia approve
  - Kirim persetujuan/penolakan
  - Tidak bisa buat expense atau kelola user lain
