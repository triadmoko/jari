# ADR-0007: Format penyimpanan rich text

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-12 (S0-15)

## Konteks

Description work item, komentar, dan custom field paragraph memuat teks kaya: heading, daftar, tabel, kode, mention, panel, status, emoji, dan lampiran inline (modul 07). README §D.3 bertanya: dokumen JSON ala ADF (ProseMirror/TipTap) atau Markdown; serta apakah impor/ekspor ADF dan wiki markup, termasuk alias `/rest/api/2`, diperlukan.

Batasan dari PRD:

- EDT-002 butir 1 dan 2: format penyimpanan utama adalah ADF JSON; server memvalidasi skema ADF pada setiap penyimpanan; HTML hanya dihasilkan oleh renderer aman.
- EDT-002 butir 4 dan 5: wiki markup dari REST v2 dan field renderer wiki ditangani sebagai konversi; markdown shortcut dan paste dikonversi ke ADF di klien lalu divalidasi server.
- EDT-001 butir 19: node ADF yang tidak didukung klien tampil sebagai blok fallback.
- Konversi keluar dari ADF: ke HTML untuk email notifikasi (dengan mention), ke wiki markup/plain text untuk REST v2 dan ekspor CSV, ke plain text untuk indeks pencarian dan preview (modul 07, bagian konversi).
- DM-003 butir 7: kolom `issues.description` berupa ADF JSON; DM-005 butir 6: custom field paragraph memakai ADF dengan batas sekitar 32.767 karakter.
- Modul 23: bentuk v3 (ADF) kanonik, v2 sebagai alias.

Kondisi repo: `frontend/package.json` belum memuat TipTap, ProseMirror, maupun library Markdown (terverifikasi dengan pencarian), jadi pilihan ini tidak terikat dependensi yang sudah ada.

## Opsi

### Opsi 1 — Dokumen JSON ala ADF, editor TipTap/ProseMirror

- Kelebihan: struktur terdefinisi sehingga mention, panel, status, tabel, dan lampiran inline terwakili setia; validasi skema di server; konversi ke HTML aman (tanpa menyimpan HTML mentah), ke plain text untuk indeks, dan ke wiki markup di tepi API; kompatibel dengan bentuk REST v3; ProseMirror/TipTap berbasis dokumen JSON.
- Kekurangan: perlu menambah TipTap di frontend dan menulis validator serta konverter ADF di backend (Go); dokumen JSON tidak nyaman diedit manual dan lebih besar dari teks biasa.

### Opsi 2 — Markdown

- Kelebihan: ringan, mudah dibaca dan diedit manual, ekosistem parser matang.
- Kekurangan: tidak dapat mewakili mention, panel, status, kolom layout, atau tabel kompleks dengan setia; dialek berbeda-beda; tidak kompatibel dengan REST v3 sehingga seluruh PRD modul 07 dan 23 perlu ditulis ulang; sanitasi HTML hasil render tetap wajib.

### Opsi 3 — HTML tersanitasi

- Kelebihan: ekspresif, langsung dirender.
- Kekurangan: sanitasi menjadi satu-satunya pertahanan XSS pada data tersimpan; sulit divalidasi terhadap skema; tidak sesuai EDT-002 butir 1.

## Keputusan

**Opsi 1.** Kami menyimpan rich text sebagai dokumen JSON ala ADF (`{"version": 1, "type": "doc", "content": [...]}`) di kolom `JSONB`, dan memakai TipTap/ProseMirror di frontend.

Aturan:

- **Skema:** Fase 1 mendukung subset node/mark ADF yang dirinci di EDT-001; server menolak node tidak dikenal atau struktur tidak valid dengan 400 (EDT-002 butir 2). Versi skema disimpan di dokumen agar bisa dimigrasi.
- **Konversi di tepi, bukan di penyimpanan:** ADF → HTML aman (render dan email), ADF → plain text (indeks pencarian dan preview), ADF → wiki markup (hanya bila alias REST v2 dibangun). Markdown shortcut dan paste dikonversi ke ADF di klien.
- **REST v2, impor/ekspor wiki markup:** tidak diwajibkan oleh ADR ini dan bukan syarat Fase 1; fase dan cakupannya ditetapkan di modul 23 (API) dan 25 (impor/ekspor). Karena penyimpanan tetap ADF, keputusan itu hanya menambah konverter dan tidak mengubah skema database.
- **Batas ukuran:** mengikuti PRD (paragraph sekitar 32.767 karakter pada custom field); batas description/komentar mengikuti modul 07 dan guardrail NFR-004.

## Konsekuensi

- Frontend menambah dependensi TipTap/ProseMirror (komponen berat seperti rich text editor dipecah per route, lihat code-splitting di modul 27); ukuran bundle perlu dipantau.
- Backend menambah paket `adf` (validator, walker, konverter HTML/plain text); ini komponen keamanan sehingga perlu uji unit dan fuzz.
- Kolom description/komentar bertipe `JSONB`; pencarian teks memakai kolom plain text hasil ekstraksi (diisi saat simpan), bukan JSON mentah.
- Mention dan referensi lampiran disimpan sebagai node dengan id (`account_id` per [ADR-0005](0005-tipe-identitas.md), id lampiran), sehingga rename user atau file tidak merusak dokumen.
- Impor dari sumber Markdown/HTML eksternal membutuhkan konverter ke ADF; dikerjakan di modul 25.
- Tindak lanjut: tiket skema ADF + validator, tiket integrasi TipTap di frontend, tiket ekstraksi plain text untuk indeks.

## Rujukan

- PRD: EDT-001, EDT-002 (butir 1–5), konversi keluar di modul 07; DM-003 butir 7; DM-005 butir 6; modul 23 (API v3/v2), modul 25 (impor/ekspor).
- `docs/prd/README.md` › Pertanyaan terbuka §D.3.
- Dokumentasi: spesifikasi ADF Atlassian dan TipTap/ProseMirror (diverifikasi saat implementasi).
