# ADR-0002: Penyimpanan lampiran (attachment)

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-11 (S0-14)

## Konteks

Work item dan komentar menerima lampiran (modul 08). README §D.9 meminta keputusan: disk lokal, S3-kompatibel, atau keduanya; batas ukuran per file dan kuota; serta nasib konversi dokumen Office dan transcoder video.

Batasan dari PRD:

- CMT-016: batas per file default 100 MB (dapat dikonfigurasi hingga 2 GB), maksimal 2.000 attachment per work item, kuota storage default tanpa batas pada self-hosted.
- CMT-017 (Fase 3): file disajikan dari origin media terpisah bila dikonfigurasi, tipe berisiko disajikan sebagai `Content-Disposition: attachment`, enkripsi at-rest dan in-transit, akses selalu mengikuti permission work item.
- CMT-019: penegakan kuota dan pemindaian antivirus opsional.
- Aplikasi dapat berjalan banyak replika (NFR-011), jadi penyimpanan tidak boleh terikat ke satu proses atau satu volume.

## Opsi

### Opsi 1 — Disk lokal saja

- Kelebihan: nol dependensi tambahan, paling sederhana untuk satu node.
- Kekurangan: sulit untuk multi-replika (butuh volume bersama), tidak ada presigned URL, kuota dan enkripsi at-rest bergantung pada host.

### Opsi 2 — S3-kompatibel saja (AWS S3, MinIO, dan penyimpanan sejenis)

- Kelebihan: satu jalur kode, skala horizontal dan multi-replika tanpa volume bersama, presigned URL, enkripsi sisi server, kuota dan lifecycle di sisi storage; cocok untuk hosted maupun self-hosted.
- Kekurangan: operator self-hosted harus menyediakan endpoint S3-kompatibel (mis. MinIO) walau hanya untuk mencoba produk; menambah satu komponen di compose.

### Opsi 3 — Abstraksi `Storage` dengan dua backend (disk lokal + S3-kompatibel)

- Kelebihan: default sederhana untuk self-hosted dan jalur naik ke object storage.
- Kekurangan: dua implementasi untuk dijaga dan diuji, dengan perilaku yang harus identik.

## Keputusan

**Keputusan pengguna: Opsi 2.** Penyimpanan lampiran memakai **S3-kompatibel saja**. Tidak ada backend disk lokal.

- Interface `AttachmentStorage` didefinisikan di `internal/shared/domain/` (put, get/stream, delete, exists, presigned URL) dengan satu implementasi S3-kompatibel di `internal/shared/infrastructure/`. Interface dipertahankan untuk pengujian (mock lewat `make mock-gen`), bukan untuk menyediakan backend alternatif.
- Konfigurasi lewat environment dan dimuat di `internal/core/config/` dengan validasi fail-fast (KAN-8): endpoint, region, bucket, access key, secret key, dan opsi path-style (diperlukan MinIO). Kredensial tidak masuk ke log atau respons API (NFR-016).
- Development dan self-hosted default memakai **MinIO** yang ditambahkan ke `docker-compose.yml` (tindak lanjut), produksi dapat memakai AWS S3 atau penyedia kompatibel lain.
- Metadata (nama asli, content type, ukuran, checksum, pengunggah, work item) disimpan di PostgreSQL; byte file hanya di bucket. Object key berbasis ID/hash, bukan nama file asli. Streaming upload/download tanpa memuat file ke memori.
- Batas default: 100 MB per file (konfigurasi hingga 2 GB), 2.000 attachment per work item, kuota default tanpa batas.
- Konversi dokumen Office dan transcoder video **bukan komponen bawaan**: keduanya dependency opsional di luar Fase 1. Pemindaian antivirus opsional (CMT-019) dan origin media terpisah (CMT-017) mengikuti jadwal fasenya.

## Konsekuensi

- Domain dan use case tidak bergantung pada SDK S3; pengujian memakai mock storage.
- Instalasi minimum menjadi PostgreSQL + endpoint S3-kompatibel; `docker-compose.yml` perlu service MinIO, bucket awal, dan variabel `.env.example` baru. Backup mencakup PostgreSQL dan bucket.
- Multi-replika tidak memerlukan volume bersama; unduhan dapat memakai presigned URL tetapi tetap harus melewati pemeriksaan permission work item (CMT-017 butir 5), sehingga URL berumur pendek dan diterbitkan setelah otorisasi.
- Versi SDK S3 (mis. `aws-sdk-go-v2` atau `minio-go`) dipilih pada tiket implementasi; kandidat dinilai di sana, bukan di ADR ini.
- Bila kelak ada kebutuhan disk lokal, diperlukan ADR baru yang menggantikan ADR ini.
- Tindak lanjut: tiket compose (MinIO + bucket init), implementasi modul 08 (CMT-001 dan seterusnya), kebijakan keamanan CMT-017 pada Fase 3.

## Rujukan

- PRD: CMT-015, CMT-016, CMT-017, CMT-019 (modul 08); NFR-011, NFR-016 (modul 27).
- `docs/prd/README.md` › Pertanyaan terbuka §D.9.
