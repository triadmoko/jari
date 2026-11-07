# Architecture Decision Records (ADR)

Keputusan teknis penting proyek dicatat di sini, satu file per keputusan, supaya alasan dan konsekuensinya bisa dirujuk kembali oleh PRD (`docs/prd/`) dan tiket Jira (project KAN).

## Aturan

- Nama file `NNNN-judul-singkat.md` dengan nomor empat digit berurutan. `0000-template.md` adalah templat.
- Bagian wajib: Status, Tanggal, Konteks, Opsi, Keputusan, Konsekuensi, Rujukan.
- Status: **Proposed** (usulan, menunggu konfirmasi) → **Accepted** (berlaku) → **Superseded oleh ADR-NNNN** (digantikan). ADR yang sudah Accepted tidak diubah isinya; keputusan baru ditulis sebagai ADR baru yang menggantikannya.
- Tautkan ADR dari modul PRD terkait dan tandai pertanyaan terbuka di `docs/prd/README.md` yang terjawab.

## Indeks

| No | Judul | Status | Tanggal | Tiket |
| --- | --- | --- | --- | --- |
| [0001](0001-library-drag-and-drop-board.md) | Library drag-and-drop untuk board | Accepted | 2026-11-07 | KAN-11 |
| [0002](0002-penyimpanan-lampiran.md) | Penyimpanan lampiran (attachment) | Accepted | 2026-11-07 | KAN-11 |
| [0003](0003-transport-email-keluar.md) | Transport email keluar | Accepted | 2026-11-07 | KAN-11 |
| [0004](0004-job-latar-dan-queue.md) | Mekanisme job latar dan queue | Accepted | 2026-11-07 | KAN-11 |

Nomor 0005 dan seterusnya disediakan untuk ADR berikutnya, termasuk ADR model data inti (KAN-12, S0-15).
