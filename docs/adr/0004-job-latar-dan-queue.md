# ADR-0004: Mekanisme job latar dan queue

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-11 (S0-14)

## Konteks

Notifikasi, webhook, otomasi, indexing, real-time, dan operasi bulk membutuhkan job asinkron serta transactional outbox (NFR-008). README §D.5 membuka pertanyaan: River di atas PostgreSQL dengan Redis opsional, atau Redis wajib; nasib Redis di `docker-compose.yml`; dan apakah outbox cukup dengan Postgres.

Kondisi repo:

- Redis sudah ada di `docker-compose.yml` (service `redis`, volume `redis_data`, dependensi service `app`), tetapi belum ada kode Go yang memakainya.
- Akses database memakai GORM (`gorm.io/driver/postgres`); migrasi lewat `golang-migrate` di `migrations/`.

Kebutuhan dari NFR-008: retry dengan exponential backoff dan dead-letter queue yang bisa dilihat/diulang/dibuang admin, idempotency key, antrean terpisah per kelas prioritas (interaktif vs batch), scheduler terjadwal yang hanya berjalan sekali per jadwal meski banyak replika, worker satu binary atau terpisah (`--role=worker`), dan shutdown yang rapi. Outbox harus ditulis dalam transaksi yang sama dengan perubahan domain (NFR-006 butir 3).

## Opsi

### Opsi 1 — River di PostgreSQL, Redis opsional

- Kelebihan (terverifikasi di dokumentasi River): job dikunci dengan `FOR UPDATE SKIP LOCKED`; `InsertTx` menyisipkan job atomik bersama transaksi aplikasi; retry/discarded/cancelled dan retry manual; `UniqueOpts`; beberapa queue; migrasi lewat `rivermigrate`; driver `riverdatabasesql` dan `riverpgxv5`.
- Kekurangan: skema River berjalan paralel dengan `golang-migrate`; integrasi GORM lewat `database/sql`; kapasitas queue berbagi database dengan data domain.

### Opsi 2 — Redis wajib (asynq)

- Kelebihan (terverifikasi di dokumentasi asynq): queue berprioritas dengan bobot (`Queues: {"critical": 6, "default": 3, "low": 1}`); `MaxRetry` per task dan `RetryDelayFunc` kustom; `Unique(ttl)` yang menolak duplikat dengan `ErrDuplicateTask`; task yang gagal habis retry masuk state archived dan dapat dijalankan ulang atau dihapus lewat `Inspector.RunTask` / `Inspector.DeleteTask`; dukungan Redis Sentinel dan Cluster; Redis sudah ada di compose; throughput tinggi dan tidak membebani PostgreSQL.
- Kekurangan: Redis menjadi dependensi infra wajib (instalasi minimum: PostgreSQL + Redis + S3); enqueue ke Redis tidak bisa atomik dengan transaksi PostgreSQL, sehingga outbox tetap harus di PostgreSQL dengan relay (lihat Keputusan); durabilitas bergantung konfigurasi persistensi Redis (AOF/RDB); bertentangan dengan NFR-008 butir 1 dan NFR-003 butir 2 saat ini, sehingga PRD perlu diperbarui.

### Opsi 3 — Implementasi sendiri di atas PostgreSQL (`SKIP LOCKED` + outbox)

- Kelebihan: tanpa dependensi baru.
- Kekurangan: retry/backoff, DLQ, scheduler, visibility timeout, dan API admin harus dibangun sendiri; risiko bug pada komponen kritis.

## Keputusan

**Keputusan pengguna: Opsi 2.** Queue dan job latar memakai **Redis** (library `asynq`); **Redis wajib** mulai Fase 1. River tidak dipakai.

Rancangan:

- **Outbox tetap di PostgreSQL.** Perubahan domain dan baris outbox ditulis dalam satu transaksi (NFR-006 butir 3). Sebuah *relay/dispatcher* membaca baris outbox yang belum terkirim (polling atau `LISTEN/NOTIFY`), meng-enqueue ke Redis, lalu menandai baris terkirim. Pengiriman bersifat *at-least-once*, jadi setiap task wajib idempoten: idempotency key diturunkan dari ID outbox dan dipakai sebagai `Unique`/task ID (opsi ini diverifikasi saat implementasi), dan handler memeriksa status baris outbox.
- **Queue**: minimal `interactive` (email, real-time; bobot tinggi) dan `batch` (bulk change, reindex, impor/ekspor; bobot rendah) dengan `Concurrency` terpisah untuk menjaga notifikasi dari operasi bulk (NFR-008 butir 6).
- **Retry dan DLQ**: `MaxRetry` dan `RetryDelayFunc` (exponential backoff) per jenis task; task archived adalah DLQ. `/api/v1/admin/jobs` dan halaman Background jobs memakai `asynq.Inspector` (daftar per queue/status, detail galat terakhir, `RunTask` untuk retry, `DeleteTask` untuk discard); retry/discard tercatat lewat `AuditRecorder.Record` (NFR-017).
- **Worker**: satu binary, mode all-in-one (default) atau `--role=worker`; shutdown rapi menunggu task aktif sampai batas waktu.
- **Scheduler**: job terjadwal (rebalance LexoRank, purge trash, dan seterusnya) hanya boleh berjalan satu kali per jadwal pada banyak replika (NFR-008 butir 8). Mekanismenya (peran scheduler tunggal atau advisory lock PostgreSQL di sekitar scheduler) diputuskan dan diverifikasi pada tiket implementasi; asynq tidak diasumsikan menangani ini sendiri.
- **Progres job** (NFR-008 butir 7): disimpan aplikasi (tabel PostgreSQL atau kunci Redis) dan diekspos lewat API; tidak diasumsikan tersedia otomatis dari asynq.
- **Persistensi Redis**: AOF diaktifkan di compose dan dokumentasi operasional; produksi dianjurkan Sentinel atau Cluster. Kehilangan Redis tidak menghilangkan pekerjaan karena baris outbox yang belum ditandai terkirim akan dikirim ulang oleh relay.
- **Konfigurasi**: `REDIS_ADDR`/`REDIS_PASSWORD`/`REDIS_DB` dimuat di `internal/core/config/` dengan validasi fail-fast (KAN-8); aplikasi menolak start tanpa Redis.
- **Penggunaan Redis lain** (cache izin dan counter rate limit NFR-003, pub/sub NFR-009) menjadi mungkin sejak Fase 2 tanpa infrastruktur tambahan, tetapi tetap tidak boleh membuat logika bisnis bergantung pada cache (hasil benar walau cache kosong).

## Konsekuensi

- Instalasi minimum menjadi PostgreSQL + Redis + endpoint S3 (ADR-0002); `docker-compose.yml` tetap memuat Redis, dengan AOF dan healthcheck, dan service `app` serta worker bergantung padanya.
- **PRD sudah diperbarui** agar konsisten: NFR-008 butir 1 dan 2 (Redis wajib dengan outbox PostgreSQL + relay), butir 8 dan catatan NFR-008, NFR-003 butir 2 dan catatan, serta README §D.5 ditandai terjawab sebagian (River ditolak, Redis wajib).
- Outbox + relay menambah satu komponen yang harus dibangun dan diuji (kegagalan di antara commit PostgreSQL dan enqueue Redis, deduplikasi, urutan).
- Skema database tidak mendapat tabel queue tambahan; `golang-migrate` tetap satu-satunya sumber kebenaran skema.
- Deployment dan pemantauan Redis (memori, persistensi, failover) menjadi tanggung jawab operator; metrik queue dibuka lewat NFR-012.
- Keputusan ini menjadi dasar ADR-0003 (pengiriman email lewat job) dan modul 20 (otomasi), 23 (webhook).
- Tindak lanjut: tiket worker + relay outbox, tiket compose (AOF, healthcheck), pembaruan `CLAUDE.md` setelah worker ada.

## Rujukan

- PRD: NFR-003, NFR-006, NFR-008, NFR-009, NFR-012, NFR-013, NFR-017; modul 28 (outbox/katalog event).
- `docs/prd/README.md` › Pertanyaan terbuka §D.5.
- Dokumentasi asynq: <https://github.com/hibiken/asynq>; River (opsi yang ditolak): <https://github.com/riverqueue/river>.
