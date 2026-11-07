# ADR-0006: Penyimpanan nilai custom field

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-12 (S0-15)

## Konteks

Work item punya field bawaan dan custom field yang dibuat admin atau admin project, dengan banyak tipe (teks, angka, tanggal, select, user picker, versi, dan seterusnya), context per project/issue type, dan opsi bertingkat. Nilainya harus bisa dibaca cepat saat membuka issue dan dicari cepat lewat JQL (modul 12) dengan operator `=`, `IN`, `>`, `<`, `~`, serta sorting. README §D.2 bertanya: hybrid JSONB + tabel bertipe sebagai keputusan final, atau EAV murni di Fase 1.

Batasan dari PRD:

- DM-005 butir 4 dan 5: kolom `issues.fields` JSONB (kunci `customfield_<id>`) untuk baca cepat, ditambah tabel `issue_field_values` bertipe untuk indexing dan JQL; GIN `jsonb_path_ops` untuk `=`/`IN`, B-tree `(field_id, number_value)` dan `(field_id, date_value)` untuk range; `EXPLAIN` pada 100 ribu issue tidak boleh sequential scan.
- DM-005 butir 10: penulisan nilai ke JSONB dan tabel bertipe dalam transaksi yang sama.
- NFR-004: batas jumlah custom field per situs (default usulan 700 field aktif).
- Compiler JQL ke SQL (JQL-050, JQL-033) adalah satu-satunya jalur pencarian; semua konsumen memakai use case yang sama.

PostgreSQL: operator class `jsonb_path_ops` hanya mendukung operator containment (`@>`) dan jsonpath (`@?`, `@@`), bukan operator range, sehingga pencarian range tidak bisa mengandalkan GIN.

## Opsi

### Opsi 1 — Hybrid: `issues.fields` JSONB + `issue_field_values` bertipe

- Kelebihan: satu baris `issues` sudah memuat semua nilai sehingga issue view, daftar, dan API tidak butuh join per field; `issue_field_values` bertipe memberi index B-tree yang tepat untuk range dan sorting; JQL `=`/`IN` memanfaatkan GIN; sesuai DM-005.
- Kekurangan: penulisan ganda dalam satu transaksi; risiko dua sumber tidak sinkron sehingga perlu job verifikasi/rebuild; skema dan kode penulisan lebih banyak.

### Opsi 2 — EAV murni (`issue_field_values` saja)

- Kelebihan: satu sumber kebenaran; sederhana; sama dengan model Jira Server/Data Center.
- Kekurangan: membaca satu issue dengan N custom field butuh query/join tambahan atau agregasi; setiap klausa JQL pada custom field menjadi join atau subquery tambahan; performa daftar issue turun seiring jumlah field. JSONB baru ditambahkan belakangan saat performa menuntut, berarti migrasi data dan perubahan kode penulisan pada sistem yang sudah berisi data.

### Opsi 3 — JSONB saja

- Kelebihan: skema paling sederhana, baca cepat.
- Kekurangan: range query dan sorting pada field dinamis butuh expression index per field (jumlah field dibuat pengguna, hingga ratusan), tidak praktis; tipe nilai tidak ditegakkan database.

## Keputusan

**Opsi 1 (hybrid).** Kami memakai `issues.fields` JSONB sebagai salinan baca dan `issue_field_values` bertipe sebagai sumber untuk indexing dan JQL, sesuai DM-005.

Aturan:

- **Sumber kebenaran:** nilai tervalidasi server-side (DM-005 butir 11) ditulis ke kedua tempat dalam satu transaksi (butir 10). `issue_field_values` dianggap dapat dibangun ulang dari JSONB untuk satu field lewat job verifikasi/rebuild.
- **Pemetaan JQL:** `=`, `IN`, `IS EMPTY` pada custom field dikompilasi ke containment JSONB (`fields @> ...`) bila didukung index GIN; operator range (`>`, `>=`, `<`, `<=`) dan sorting dikompilasi ke `issue_field_values` memakai kolom bertipe (`number_value`, `date_value`); field multi-value disimpan sebagai beberapa baris. Detail kompilasi per tipe ada di JQL-033 dan JQL-050.
- **Id:** `custom_fields.id` dari sequence mulai 10000 dan tidak dipakai ulang (lihat [ADR-0005](0005-tipe-identitas.md)).
- **Bukti performa:** pengujian `EXPLAIN` pada data uji 100 ribu issue menjadi syarat selesai tiket implementasi (DM-005 butir 5).
- **Fase 1:** cakupan tipe field Fase 1 mengikuti DM-005 butir 6; tipe di luar daftar ditolak 400.

## Konsekuensi

- Penulisan nilai custom field wajib melewati satu service yang menulis JSONB dan tabel bertipe dalam satu transaksi; handler tidak boleh menulis salah satunya langsung.
- Perlu job verifikasi/rebuild `issue_field_values` dan metrik drift antara kedua sumber.
- Penyimpanan bertambah (nilai ada di dua tempat); diterima demi performa baca dan pencarian.
- Pemilihan tipe JSONB di sisi Go (GORM/pgx) dan format migrasi indeks GIN diputuskan pada tiket implementasi; `go.mod` saat ini belum memuat library JSON khusus GORM.
- Mengubah `type_key` custom field tetap tidak didukung (DM-005 butir 9), sehingga tidak ada migrasi tipe nilai.
- Tindak lanjut: tiket migrasi `custom_fields`, `custom_field_contexts`, `custom_field_options`, `issue_field_values` dan kolom `issues.fields`; tiket compiler JQL untuk custom field.

## Rujukan

- PRD: DM-005 (butir 1–11), NFR-004, NFR-005; JQL-033, JQL-050; modul 28 › Pertanyaan terbuka butir 3; CFG-025 s.d. CFG-038, CFG-040 (modul 05).
- `docs/prd/README.md` › Pertanyaan terbuka §D.2.
- Dokumentasi PostgreSQL: tipe JSONB dan operator class `jsonb_path_ops`.
