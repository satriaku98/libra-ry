Installation

go mod install

Running the application

go run cmd/main.go

Swagger initialization

swag init -g cmd/main.go

Access Swagger UI

http://localhost:{APP_PORT}/docs

1. ERD (Entity Relationship Diagram)
Entitas dan Relasi
Buku (id_buku, judul, penulis, penerbit, tahun_terbit, stok)
Anggota (id_anggota, nama, alamat, no_hp)
Peminjaman (id_peminjaman, id_anggota, tanggal_pinjam, tanggal_kembali, status_peminjaman)
Detail Peminjaman (id_detail, id_peminjaman, id_buku, jumlah)

Relasi:
Anggota bisa meminjam banyak buku → 1 anggota bisa memiliki banyak peminjaman
Setiap peminjaman bisa memiliki banyak buku → Relasi many-to-many antara Buku dan Peminjaman, direpresentasikan oleh tabel Detail Peminjaman

2. Flowchart
a. Flowchart CRUD Stok Buku
Admin memilih menu manajemen buku
Admin dapat menambahkan, mengubah, atau menghapus buku
Jika menambahkan → Masukkan data buku
Jika mengubah → Pilih buku, edit data
Jika menghapus → Pilih buku, hapus
Simpan perubahan ke database

b. Flowchart Transaksi Peminjaman & Pengembalian
Anggota memilih buku yang ingin dipinjam
Sistem mengecek ketersediaan stok
Jika stok tersedia → Lanjutkan ke transaksi
Jika stok tidak tersedia → Beri pemberitahuan
Sistem mencatat transaksi peminjaman
Anggota mengembalikan buku
Sistem mengecek status pengembalian
Jika buku dikembalikan tepat waktu → Transaksi selesai
Jika terlambat → Tampilkan denda (jika ada)