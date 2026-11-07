# ADR-0005: Tipe identitas entitas inti

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-12 (S0-15)

## Konteks

Migrasi entitas domain pertama (project, work item, konfigurasi) harus memilih tipe primary key. README §D.1 membuka tiga pertanyaan: tipe `id` project/work item (bigint numerik atau UUID), nasib `users.id` UUID `VARCHAR(36)` (dipertahankan atau diseragamkan), dan apakah `account_id` berupa kolom terpisah.

Kondisi repo (terverifikasi):

- `migrations/001_create_users_table.up.sql`: `users.id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text`.
- `internal/shared/domain/entity/user.go`: `ID string` dengan tag `gorm:"type:varchar(36);primaryKey"`, diisi `uuid.New().String()`.
- Klaim JWT dan context `"sess"` membawa id pengguna sebagai string.

Batasan dari PRD:

- DM-003 butir 1 dan 7: `projects.id` dan `issues.id` bertipe `BIGINT`; issue dirujuk lewat key `PROJ-123` dan id numerik (`/rest/api/3/issue/{issueIdOrKey}`).
- DM-005 butir 1: id custom field berupa sequence numerik mulai 10000 (`customfield_10000`) dan tidak dipakai ulang.
- DM-001 butir 1: `users` punya `id` internal dan `account_id` opaque, unik, immutable, dipakai di semua API/payload publik.
- Tujuan kompatibilitas bentuk API dengan vendor pembanding (README bagian E), tanpa mereplikasi id internal vendor.

Kriteria penilaian: kompatibilitas API, biaya migrasi terhadap pekerjaan yang sudah selesai (S0-01/S0-02), integritas FK, dan ketahanan kontrak API terhadap perubahan internal.

## Opsi

### Opsi 1 — BIGINT untuk entitas domain, `users.id` tetap `VARCHAR(36)`, `account_id` kolom terpisah

- Kelebihan: sejalan dengan DM-001 dan DM-003; tidak ada migrasi primary key `users` (S0-01/S0-02 tidak tersentuh); id numerik kompak untuk index, FK, dan join pada tabel besar (`issues`, `issue_field_values`, changelog); kontrak API publik (`account_id`) terlepas dari id internal.
- Kekurangan: dua gaya id di satu skema; satu kolom tambahan di `users` beserta backfill.

### Opsi 2 — BIGINT untuk entitas domain, `users.id` tetap, `account_id` = `users.id`

- Kelebihan: tanpa kolom baru, migrasi paling sederhana.
- Kekurangan: id internal terikat ke kontrak API publik sehingga tidak bisa diubah tanpa memutus klien; menyimpang dari DM-001 butir 1 sehingga PRD harus diubah.

### Opsi 3 — BIGINT untuk entitas domain, `users.id` diubah ke tipe `uuid` native

- Kelebihan: tipe lebih ketat dan lebih hemat ruang (16 byte).
- Kekurangan: migrasi primary key `users` dan semua FK yang kelak menunjuknya; menyentuh S0-01/S0-02 yang sudah selesai; kode memakai `string` di entity, JWT, dan repository sehingga perubahan menjalar tanpa manfaat fungsional.

### Opsi 4 — UUID untuk semua entitas

- Kelebihan: satu gaya id, tanpa sequence.
- Kekurangan: tidak kompatibel dengan bentuk API yang diacu (id numerik, `customfield_10000`); index lebih besar dan UUIDv4 acak memperburuk lokalitas index pada tabel besar; menyimpang dari DM-003 dan DM-005.

## Keputusan

**Opsi 1.** Kami memakai:

- `BIGINT` (identity/sequence) untuk primary key `projects`, `issues`, dan entitas konfigurasi/domain lain; id custom field dari sequence yang dimulai 10000 (DM-005).
- `users.id` tetap `VARCHAR(36)` berisi UUID teks, tanpa migrasi primary key.
- Kolom `users.account_id` terpisah: string opaque, unik, immutable, dibuat saat user dibuat (migrasi DM-001 mengisi baris yang sudah ada). Hanya `account_id` yang muncul di API/payload publik, token, dan webhook; FK internal antar tabel memakai `users.id`.

Aturan turunan:

- API menserialisasi id numerik sebagai **string** di JSON (mis. `"id": "10042"`), agar aman dari batas integer JavaScript dan setara dengan bentuk API pembanding.
- Id numerik tidak pernah dipakai ulang setelah baris dihapus (sequence, bukan `MAX(id)+1`).
- Id adalah identitas global per instalasi. Keputusan tenancy (README §B.1) belum diambil; bila `org_id` ditambahkan kelak, ia menjadi kolom tambahan dan tidak mengubah tipe id yang diputuskan di sini.

## Konsekuensi

- Skema `users` mendapat migrasi baru (bukan perubahan `001`): tambah `account_id` + unique index + backfill. Ini dikerjakan pada tiket migrasi DM-001, bukan di sini.
- Tidak ada perubahan pada S0-01/S0-02; entity `User.ID` tetap `string`.
- Kode perlu dua tipe id: `string` untuk user dan `int64` untuk entitas domain; handler harus membedakan `account_id` (publik) dari `users.id` (internal) pada DTO.
- Kolom PRD yang bernama `*_account_id` (mis. `projects.lead_account_id` pada DM-003) perlu dipetakan: disarankan menyimpan FK ke `users.id` demi integritas referensial dan mengubahnya ke `account_id` pada lapisan API. Penamaan kolom final ditetapkan pada tiket migrasi tabel `projects`.
- Tindak lanjut: tiket migrasi `account_id` (DM-001), tiket migrasi `projects`/`issues` dengan identity column, dan konvensi serialisasi id string di paket respons API.

## Rujukan

- PRD: DM-001 butir 1, DM-003 butir 1 dan 7, DM-005 butir 1; modul 28 › Pertanyaan terbuka butir 1; modul 04 (PRJ-001).
- `docs/prd/README.md` › Pertanyaan terbuka §D.1 (juga §B.1 tenancy dan bagian E kompatibilitas API).
- Repo: `migrations/001_create_users_table.up.sql`, `internal/shared/domain/entity/user.go`.
