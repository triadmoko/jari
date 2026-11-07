# Laporan lisensi dependensi

Dihasilkan untuk KAN-18 (S0-12). Berkas di folder ini adalah keluaran tool, jangan diedit tangan.

| Berkas | Isi | Sumber |
|--------|-----|--------|
| `go.csv` | Dependensi Go yang masuk ke binary `cmd/api` (paket repo sendiri dikecualikan): `paket,URL lisensi,jenis lisensi` | `go-licenses` v1.6.0 |
| `frontend.csv` | Semua dependensi `frontend/` (termasuk devDependencies, karena Svelte, bits-ui, dan sejenisnya ikut ter-bundle ke SPA): `modul,lisensi,repositori` | `license-checker-rseidelsohn` 4.4.2 |

## Perintah

```bash
make licenses         # regenerasi go.csv dan frontend.csv
make licenses-check   # gagal bila ada lisensi di luar daftar yang diizinkan
```

## Kebijakan

Proyek ini berlisensi **Apache 2.0** (berkas `LICENSE` di root). Kebijakan: dependensi hanya boleh berlisensi permisif atau copyleft lemah tingkat berkas. Lisensi copyleft kuat (GPL, AGPL, LGPL, SSPL, dan sejenisnya) ditolak oleh `make licenses-check`; khusus GPLv2-only memang tidak kompatibel dengan Apache 2.0.

Saat mendistribusikan binary atau image, sertakan teks lisensi dan berkas `NOTICE` dari dependensi Apache 2.0 yang memilikinya.

Hasil terakhir (2026-11-07):

- **Go (50 paket):** MIT 30, BSD-3-Clause 9, Apache-2.0 11. Tidak ada pelanggaran.
- **Frontend (508 modul):** didominasi MIT, ISC, Apache-2.0, dan BSD. Lisensi yang perlu dicatat:
  - `MPL-2.0` (lightningcss, alat build, tidak ikut ter-bundle) dan `(MPL-2.0 OR Apache-2.0)` (dompurify, dapat dipakai di bawah Apache-2.0). Copyleft lemah tingkat berkas, tidak menular ke kode proyek.
  - `OFL-1.1` (`@fontsource-variable/inter` dan `lora`): lisensi font, boleh didistribusikan bersama aplikasi.
  - `Python-2.0` (argparse), `0BSD`, `BlueOak-1.0.0`, `Unlicense`: permisif.
  - `MIT*` pada runed dan svelte-toolbelt: tool tidak menemukan berkas lisensi baku, tetapi `package.json` menyatakan MIT.

Bila menambah dependensi dengan lisensi baru, tinjau dulu kecocokannya dengan lisensi final proyek, lalu tambahkan ke `WEB_ALLOWED_LICENSES` di `Makefile` dan catat di sini.

## Bila lisensi proyek berubah

Samakan `LICENSE`, anotasi `@license` di `cmd/api/main.go` (lalu `make swag api-gen`), bagian Lisensi di `README.md`, komentar di `Makefile`, dan tinjau ulang daftar lisensi yang diizinkan di atas.
