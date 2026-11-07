# Jari

Jari adalah platform manajemen kerja dan proyek open-source yang meniru fitur inti Jira: proyek, work item, board, backlog dan sprint, workflow, pencarian (JQL), laporan, dan administrasi. "Jari" adalah **nama kerja** sampai nama produk final ditetapkan.

Proyek ini masih dalam tahap awal (Sprint 0, remediasi baseline). Ruang lingkup, fase delivery, dan keputusan teknis ada di dokumen berikut:

- [PRD](docs/prd/README.md): visi, ruang lingkup per fase, dan spesifikasi fitur per modul. Mulai dari [modul 00](docs/prd/00-visi-ruang-lingkup-glosarium-baseline.md).
- [ADR](docs/adr/README.md): catatan keputusan teknis beserta alasannya.
- Pelacakan pekerjaan: project **KAN** di Jira (`triadmoko.atlassian.net`).

## Stack

| Komponen | Teknologi |
|----------|-----------|
| Backend | Go 1.26, [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io) |
| Database | PostgreSQL 15+, migrasi dengan [golang-migrate](https://github.com/golang-migrate/migrate) |
| Queue / job latar | Redis 7 dengan `asynq` ([ADR-0004](docs/adr/0004-job-latar-dan-queue.md)); belum dipakai kode |
| Autentikasi | JWT ([golang-jwt](https://github.com/golang-jwt/jwt) v5), password bcrypt |
| Frontend | SvelteKit (SPA statis) di `frontend/`, di-embed ke binary Go |
| Dokumentasi API | Swagger ([swag](https://github.com/swaggo/swag)) dengan UI Scalar |
| Logging | [Logrus](https://github.com/sirupsen/logrus) |

Arsitektur backend memakai Clean Architecture dengan modul per fitur. Panduan lengkap untuk kontributor ada di [CLAUDE.md](CLAUDE.md).

## Prasyarat

- Go 1.26+
- PostgreSQL 15+
- [bun](https://bun.sh) untuk membangun frontend
- Docker dan Docker Compose (opsional, untuk menjalankan semua layanan)
- [gow](https://github.com/mitranim/gow) untuk hot-reload (opsional)
- [migrate CLI](https://github.com/golang-migrate/migrate) (hanya untuk `make migration-create`)
- [swag](https://github.com/swaggo/swag) (hanya untuk `make swag`)

## Menjalankan

### Dengan Docker

```bash
cp .env.example .env
# isi JWT_SECRET (openssl rand -hex 32) dan DB_PASS di .env
docker compose up -d
```

Compose menjalankan `postgres`, `redis`, job `migrate` satu kali (menerapkan migrasi), lalu `app`. API tersedia di `http://localhost:8080`.

### Tanpa Docker (pengembangan lokal)

```bash
cp .env.example .env     # sesuaikan koneksi database
make migration-up        # terapkan migrasi
make dev                 # backend dengan hot-reload (butuh gow)
make web-dev             # opsional: Vite dev server, proxy /api ke :8080
```

Untuk satu binary berisi backend dan frontend: `make build`, lalu jalankan `./main`.

### Migrasi

File SQL ada di `migrations/` (`NNN_nama.up.sql` / `.down.sql`) dan ter-embed di binary. Server menolak start jika skema belum termigrasi, kotor, atau lebih baru dari binary.

```bash
make migration-up                    # terapkan semua migrasi
make migration-down n=1              # mundur N migrasi (n=all untuk semuanya)
make migration-version               # versi saat ini
make migration-create name=xxx       # buat file migrasi baru
make migration-force version=N       # paksa versi (untuk memulihkan status kotor)
```

## Konfigurasi

Semua konfigurasi dibaca dari environment variable (lihat `.env.example`).

| Variabel | Keterangan | Default |
|----------|------------|---------|
| `SERVER_HOST` / `SERVER_PORT` | Alamat dan port HTTP | `0.0.0.0` / `8080` |
| `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`, `DB_SSLMODE` | Koneksi PostgreSQL | `localhost`, `5432`, `postgres`, `password`, `db_name`, `disable` |
| `DATABASE_URL` | Connection string tunggal, lebih diutamakan daripada variabel `DB_*` | kosong |
| `JWT_SECRET` | Kunci tanda tangan JWT. Wajib kecuali `ENV=development`; server menolak start jika kosong atau berupa placeholder | wajib |
| `CORS_ALLOWED_ORIGINS` | Origin browser yang boleh memanggil API, dipisah koma (`*` mengizinkan semua tanpa kredensial). Kosong menolak semua permintaan lintas origin; SPA yang di-embed same-origin tidak terpengaruh. Pada `development`, kosong mengizinkan `http://localhost:5173` | kosong |
| `ENV` | `development`, `production`, atau `test` | `development` |

## API

Semua rute berada di bawah `/api/v1`. Dokumentasi interaktif (Scalar) ada di `http://localhost:8080/swagger/index.html`, spesifikasinya di `/swagger/doc.json`. Setelah mengubah anotasi API, jalankan `make swag`.

| Method | Endpoint | Auth | Keterangan |
|--------|----------|:----:|------------|
| `POST` | `/api/v1/auth/register` | Tidak | Daftar pengguna baru |
| `POST` | `/api/v1/auth/login` | Tidak | Login, mengembalikan JWT |
| `GET` | `/api/v1/users/profile` | Ya | Profil pengguna terautentikasi |
| `PUT` | `/api/v1/users/profile` | Ya | Perbarui profil |
| `GET` | `/api/v1/users` | Admin | Daftar pengguna (paginasi) |
| `GET` | `/health` | Tidak | Health check |

Autentikasi memakai header `Authorization: Bearer <token>`.

## Struktur repo

```
├── cmd/api/              # Entry point (server dan subcommand migrate)
├── internal/
│   ├── app/              # Inisialisasi aplikasi dan routing
│   ├── core/config/      # Pemuatan dan validasi konfigurasi
│   ├── features/         # Modul fitur (auth, user)
│   └── shared/           # Domain, infrastruktur, middleware, response, webui
├── pkg/                  # Paket reusable (jwt, crypto, logger, pagination)
├── migrations/           # Migrasi SQL (ter-embed ke binary)
├── frontend/             # SvelteKit SPA
└── docs/
    ├── prd/              # Product Requirements Document
    ├── adr/              # Architecture Decision Records
    └── (swagger)         # Hasil generate swag
```

## Perintah pengembangan

| Perintah | Keterangan |
|----------|------------|
| `make dev` | Jalankan backend dengan hot-reload |
| `make test` | Jalankan semua test |
| `make test-coverage` / `make test-coverage-report` | Laporan coverage (HTML / terminal) |
| `make swag` | Generate ulang dokumentasi Swagger |
| `make mock-gen` | Generate mock dengan mockery |
| `make web-dev` / `make web-build` | Frontend: dev server / build ke dalam binary |
| `make build` | Build frontend lalu binary Go |

## Lisensi

Lisensi proyek dan nama final produk masih keputusan terbuka (lihat [PRD modul 00](docs/prd/00-visi-ruang-lingkup-glosarium-baseline.md)). Sampai ditetapkan, berkas ini belum menyertakan lisensi.
