# Cinema Book System
---
### Achmad Nashruddin Riskynanda
---
### Tentang
Api ini dibangun menggunakan Golang dan Gin Framework. Pengelolaan yang akan diimplementasikan meliputi manajemen user, manajemen pegawai, multi-cinema, scheduling, booking, hingga payment.
- Menerapkan Dependency Injection/Inversion
- Load Config yang terstruktur di `config/`
- Manajemen service atau framework yang digunakan dilakukan didalam `service/*`
- Manajemen database, external api, komunikasi protokol lain berada didalam `infrastructure/*`
- Manajemen middleware dilakukan didalam `internal/middleware/*`
- Manajemen utility function dilakukan didalam `internal/utils/*`
- Manajemen error dilakukan dengan membuat custom error type beserta fungsi untuk menentukan kode http status dari error tersebut `internal/utils/errors.go`

### Kelebihan
Dengan penerapan yang telah dilakukan, dapat dilihat bahwa:
- Role based engine untuk employee multi-tenant. Dapat mendukung modularitas role pada tenant yang berbeda-beda. Sehingga setiap tenant punya bentuk role mereka sendiri.
- Struktur kode yang modular dan terorganisir memudahkan proses pengembangan, debugging, dan scaling aplikasi.
- Dependency Injection memungkinkan pengujian unit (unit testing) menjadi lebih mudah karena setiap komponen dapat digantikan dengan mock sesuai kebutuhan.
- Clean Architecture menjaga batas tanggung jawab setiap layer (handler, usecase, repository), sehingga logika bisnis tidak tercampur dengan detail implementasi teknis seperti database atau HTTP framework.
- Konfigurasi fleksibel melalui file .env yang memudahkan pengaturan environment tanpa harus mengubah source code.
- Middleware terstruktur memungkinkan penerapan fitur seperti autentikasi JWT secara konsisten dan mudah dikembangkan lebih lanjut (misalnya logging, rate-limiting).
- Custom error handling memberikan kontrol lebih atas response error API agar lebih informatif dan sesuai dengan standar HTTP status code.
Dengan kelebihan tersebut tentunya akan dapat memudahkan proses maintenance dan reduce technical debt secara jangka panjang.
---

## 🧩 Soal A. System Design

### Topologi
![System Design](documents/system-design/SystemDesign.png)
---
### Booking Flow
![Book Flow](documents/system-design/Bookings.png)
---
### Penjelasan

Platform ini dirancang untuk mendukung pembelian tiket bioskop secara online dengan **jaminan eksklusivitas kursi**, skalabilitas untuk bioskop nasional, dan fleksibilitas saat terjadi perubahan jadwal atau pembatalan.

---

### 🎫 1. Pemilihan Tempat Duduk (Seat Selection System)

Untuk menghindari konflik pemesanan antar pengguna secara paralel, sistem pemilihan tempat duduk dibangun dengan pendekatan sebagai berikut:

- Saat user memilih kursi, sistem akan **melock kursi tersebut secara temporer** dengan status `LOCKED` dalam tabel `booking_seats`, dan disimpan sementara di Redis (untuk fast-access).
- Kursi yang berstatus `LOCKED` akan otomatis diubah menjadi `EXPIRED` jika **tidak dibayar dalam waktu tertentu** (misalnya 5 menit), melalui mekanisme background job (cron/kafka event).
- Saat pembayaran berhasil, status akan berubah menjadi `BOOKED`, dan tidak bisa digunakan oleh pengguna lain.
- Penggunaan Redis membuat sistem cepat dan bisa menangani ribuan user secara paralel tanpa blocking ke DB utama.

---

### 🔁 2. Restok Tiket (Releasing Booked/Locked Seats)

Tiket (kursi) yang pernah dipesan bisa direstock dalam beberapa kondisi:

- **Kursi tidak dibayar** → otomatis masuk ke status `EXPIRED` dan dapat dipilih ulang oleh user lain.
- **Refund berhasil** → sistem akan melakukan rollback pada `booking_seats` dan menghapus relasi booking-nya. Kursi akan tersedia kembali.
- Redis di-repopulate kembali oleh consumer (event listener) untuk mencerminkan kursi yang kembali tersedia.

---

### 💸 3. Refund / Pembatalan Jadwal

Jika pihak bioskop membatalkan atau menunda film, alur refund dilakukan sebagai berikut:

- `schedules.status` akan diubah menjadi `CANCELLED` atau `POSTPONED`
- Semua booking terkait akan diproses oleh **background worker**:
  - Jika `CANCELLED`, maka entri `payments` akan diupdate jadi `REFUNDED`
  - Dana akan dikembalikan via payment gateway
  - User akan mendapat notifikasi (via Notification Service)
  - Kursi dihapus dari booking dan kembali tersedia

---


## 🗂️ Teknologi yang Digunakan

| Layer | Tools |
|------|-------|
| Bahasa Pemrograman | Golang |
| Database | PostgreSQL (transaksi), MongoDB (analitik) |
| Cache & Locking | Redis |
| Observability | Prometheus, Grafana |
| Authentication | JWT |
| Container | Docker |
| API Testing | Postman |

---
## 💽 Soal B. Database Design

Desain skema relasional dengan PostgreSQL. Mendukung:

- Pemesanan berdasarkan jadwal tayang
- Lock kursi sementara
- Status transaksi dan refund
- Tipe ENUM untuk status
- Multi Tenant
- Adaptive Multi Role Tenant & Resource Based
**Tenant** disini berupa **Cinema**
---
![Database ERD](documents/erd/erd.png)
---

## Soal C. Skill Test
Buat API sederhana menggunakan Golang dengan database point test B. API tersebut adalah :
1. Buat API Login User
2. Buat API CRUD data jadwal tayang
Note: API tersebut sudah menggunakan Authorization pada point 1

## Auth Endpoint
| Method | Endpoint         | Payload                     | Deskripsi                   |
|--------|------------------|-----------------------------|-----------------------------|
| POST   | `/auth/login`    | `{"email":"val", "password":"val"}`|Login dan mendapatkan token
| POST   | `/auth/refresh`  | http cookie only | Refresh Token
| POST   | `/auth/register`  | `{"fullname": "val", "email":"val", "password":"val"}` | Refresh Token

## Schedules Endpoint (Employee)
| Method | Endpoint         | Payload                     | Deskripsi                   |
|--------|------------------|-----------------------------|-----------------------------|
| GET   | `/cinema/:cinema_id/schedules`    | - | Mendapatkan semua schedule
| GET   | `/cinema/:cinema_id/schedules/:id`  | - | Mendapatkan schedule berdasarkan ID
| POST   | `/cinema/:cinema_id/schedules`  | `{"hall_id" : "val","movie_id" : "val","show_time" : "val","price" : "val"}` | Membuat jadwal tayang
| PUT   | `/cinema/:cinema_id/schedules/:id`  | `{"show_time" : "val","price" : "val"}` | Update jadwal tayang
| POST   | `/cinema/:cinema_id/schedules/:id/cancel`  | - | Cancel jadwal tayang
| POST   | `/cinema/:cinema_id/schedules/:id/postpone`  | `{"show_time" : "val"}` | Postpone jadwal tayang

## 📚 Dokumentasi Teknis
| Komponen                          | Link                                     |
|----------------------------------|------------------------------------------|
| 🧠 ERD (Entity Relationship Diagram) | [erd](documents/erd)   |
| 🗄️ DDL (SQL Schema PostgreSQL)       | [ddl](documents/ddl)    |
| 📮 Postman Collection               | [postman](documents/postman-export)|
| 🖼️ System Design (Topology)         | [system design](documents/system-design)|
| 📝 Contoh .env (konfigurasi)        | [env.example](./env.example) |

