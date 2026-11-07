# ADR-0003: Transport email keluar

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-11 (S0-14)

## Konteks

Notifikasi email (modul 19) wajib ada di Fase 1. README §D.10 membuka pertanyaan penyedia email keluar: SMTP sendiri, SES, atau Mailgun. ADR ini hanya memutuskan **transport keluar**. Email masuk (IMAP/webhook), domain bawaan service project, dan balasan email menjadi comment tetap terbuka karena baru relevan pada Fase 3 (NTF-031, NTF-034, modul 29–30).

Batasan dari PRD (NTF-023):

- Produk self-hosted: konfigurasi SMTP server wajib, dengan TLS implisit dan STARTTLS, port default 587 atau 465, autentikasi, dan konfigurasi dari environment (`SMTP_*`) atau admin UI.
- Pengiriman lewat antrean dengan retry dan rate limit per domain tujuan (default 5 koneksi paralel dan 60 pesan per menit), template HTML + plaintext per bahasa dan timezone penerima, serta header `Message-ID`, `In-Reply-To`, `References`, `List-Unsubscribe`.
- Mode development memakai mail catcher (Mailpit) lewat `SMTP_HOST`/`SMTP_PORT`; toggle `MAIL_OUTGOING_DISABLED`.
- Kredensial SMTP disimpan terenkripsi dan tidak ditampilkan kembali (NFR-016).

## Opsi

### Opsi 1 — SMTP generik sebagai satu-satunya transport Fase 1

- Kelebihan: bekerja dengan penyedia mana pun, termasuk SES dan Mailgun lewat antarmuka SMTP mereka; sesuai kebutuhan self-hosted; mudah diuji lokal dengan Mailpit; tidak ada SDK vendor.
- Kekurangan: fitur khusus penyedia (webhook bounce/complaint, statistik pengiriman, template server-side) tidak tersedia; penanganan bounce (suppression list, NTF-023 butir 9) harus lewat jalur lain.

### Opsi 2 — SDK/API Amazon SES

- Kelebihan: throughput tinggi, notifikasi bounce/complaint lewat SNS.
- Kekurangan: mengikat ke AWS, memaksa kredensial dan SDK cloud pada instalasi self-hosted.

### Opsi 3 — SDK/API Mailgun

- Kelebihan: API bounce/event yang kaya, inbound routing.
- Kekurangan: mengikat ke vendor tertentu; biaya langganan; sama-sama tidak cocok sebagai satu-satunya jalur self-hosted.

## Keputusan

**Usulan: Opsi 1.** Fase 1 memakai **SMTP** sebagai satu-satunya transport. Dibuat interface `Mailer` di domain dengan satu implementasi SMTP di infrastruktur sehingga penyedia berbasis API (SES, Mailgun) dapat ditambah belakangan tanpa mengubah use case.

- Konfigurasi: `SMTP_HOST`, `SMTP_PORT`, `SMTP_USERNAME`, `SMTP_PASSWORD`, mode TLS/STARTTLS, `MAIL_FROM`, dan `MAIL_OUTGOING_DISABLED`, dimuat lewat `internal/core/config/` dengan validasi fail-fast (KAN-8). Override dari admin UI menyusul bersama modul 22.
- Pengiriman tidak sinkron di handler HTTP: use case menulis outbox dan job pengiriman email ke queue dari ADR-0004 (retry, backoff, rate limit per domain di sisi job).
- Template memakai paket `html/template` + `text/template` dengan katalog label ID/EN yang sama dengan UI (S0-13).
- Dev lokal: Mailpit ditambahkan ke compose development (tindak lanjut), produksi memakai relay SMTP pilihan operator.

## Konsekuensi

- Tidak ada ketergantungan SDK vendor di Fase 1; operator memilih relay sendiri (termasuk SES atau Mailgun lewat SMTP).
- Penanganan bounce dan suppression list tidak otomatis dari webhook penyedia; Fase 1 mencatat kegagalan SMTP sinkron saja, dan webhook bounce menjadi tindak lanjut Fase 2 atau 3 bila dibutuhkan.
- Mengandalkan ADR-0004: tanpa queue yang andal, retry email tidak dapat dipenuhi.
- Perlu dependensi pustaka SMTP/Go (dipilih saat implementasi; kandidat dinilai pada tiket implementasi, bukan di ADR ini).
- Tindak lanjut: kontrak `Mailer`, tabel outbox email dan log kirim (modul 19, 28), compose dev dengan Mailpit.

## Rujukan

- PRD: NTF-023, NTF-024, NTF-018, NTF-010, NFR-008, NFR-016.
- `docs/prd/README.md` › Pertanyaan terbuka §D.10.
