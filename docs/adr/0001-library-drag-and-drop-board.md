# ADR-0001: Drag-and-drop untuk board (dibangun sendiri, reusable)

- **Status:** Accepted
- **Tanggal:** 2026-11-07
- **Tiket:** KAN-11 (S0-14)

## Konteks

Board Kanban (modul 10) adalah antarmuka utama Fase 1. Kartu harus bisa dipindah antarkolom (menjalankan transisi workflow, BRD-016), diurutkan dalam kolom (rank, BRD-017), dan dipindah antarswimlane. Drag-and-drop juga dibutuhkan di luar board: backlog dan sprint (modul 11), urutan kolom dan status di konfigurasi board/workflow, urutan field/layar (modul 05), dan lampiran. Frontend memakai SvelteKit dengan Svelte 5 (`svelte ^5.57`, runes) sebagai SPA statis (NFR-020). `frontend/package.json` belum memuat library drag-and-drop. README › Risiko #9 menyebut belum ada library DnD yang aksesibel dan meminta evaluasi di Sprint 0.

Kriteria penilaian:

1. **Svelte 5 / runes** tanpa lapisan kompatibilitas.
2. **Aksesibilitas keyboard** (BRD-016 butir 22, NFR-022 / WCAG 2.1 AA): fokus kartu, Space mengangkat, panah memindah, Space melepas, Esc membatalkan; pengumuman untuk screen reader.
3. **Multi-list**: pindah antarkolom dan antarswimlane, serta drop zone yang dapat dinonaktifkan (BRD-016 butir 2 dan 18).
4. **Integrasi dengan optimistic update + rollback** (NFR-020 butir 3) dan **virtualisasi daftar besar** (NFR-020 butir 4).
5. Ukuran bundle (NFR-001 butir 9), lisensi, tingkat pemeliharaan, dan kemampuan dipakai ulang di banyak layar.

Evaluasi library berbasis dokumentasi resmi masing-masing (via Context7, 2026-11-07); tidak ada prototipe.

## Opsi

### Opsi 1 — `svelte-dnd-action`

- Kelebihan: sintaks event Svelte 5 (`onconsider` / `onfinalize`) didukung; aksesibilitas bawaan (drag keyboard dengan Space/Enter, ARIA otomatis); multi-list lewat `type`, serta `dropFromOthersDisabled` / `dragDisabled` per zona; drag handle dan sentuhan.
- Kekurangan: `items` harus sama urutannya dengan anak DOM (`#each`), sehingga virtualisasi kolom perlu penanganan khusus; model action + event `consider`/`finalize` perlu dibungkus agar cocok dengan runes dan optimistic update; perilaku keyboard dan pengumuman ditentukan library.

### Opsi 2 — `@dnd-kit-svelte/svelte` (port komunitas dnd-kit)

- Kelebihan: ditulis untuk Svelte 5 (`$state`, `{@attach}`); `useSortable` dengan `group`/`type`/`accept`; `KeyboardSensor` dan `PointerSensor`; overlay terpisah dari data.
- Kekurangan: port komunitas, dokumentasi bersumber dari cabang `experimental`; risiko stabilitas API dan pemeliharaan.

### Opsi 3 — `@atlaskit/pragmatic-drag-and-drop`

- Kelebihan: framework-agnostic, sangat ringan (adapter elemen ±4,7 kB gzip menurut dokumentasinya), memakai drag-and-drop bawaan browser.
- Kekurangan: paket inti tidak menyediakan kontrol aksesibel otomatis (dokumentasinya meminta alternatif tombol/menu); tidak ada binding Svelte resmi.

### Opsi 4 — `@horuse/svelte-dnd`

- Kelebihan: khusus Svelte 5, tanpa dependensi lain, sensor pointer/sentuh/keyboard dan pengumuman screen reader dapat dikustomisasi, pemindahan antarkontainer bergaya kanban.
- Kekurangan: proyek lebih muda dengan reputasi sumber lebih rendah (Medium); kematangan belum teruji.

### Opsi 5 — Bangun sendiri sebagai modul reusable

- Kelebihan: kontrol penuh atas perilaku keyboard, pengumuman, dan aturan drop (izin Transition/Schedule, drop zone nonaktif) persis sesuai PRD; satu API untuk semua layar (board, backlog, sprint, urutan kolom/field); tidak ada dependensi runtime baru dan tidak terikat jadwal pemeliharaan pihak lain; dapat dirancang ramah virtualisasi sejak awal (tidak mensyaratkan urutan DOM = urutan data).
- Kekurangan: biaya awal dan pemeliharaan paling besar; kasus tepi yang sudah ditangani library matang (auto-scroll, sentuhan, fokus setelah drop, layar kecil, screen reader) harus dibangun dan diuji sendiri; risiko aksesibilitas ada pada kita.

## Keputusan

**Keputusan pengguna: Opsi 5.** Kita membangun drag-and-drop sendiri sebagai modul **reusable** yang dapat dipakai di fitur mana pun, bukan library eksternal.

Arah rancangan (detail final diputuskan pada tiket implementasi):

- **Lokasi**: `frontend/src/lib/dnd/`, tidak bergantung pada fitur board; fitur memakainya lewat API publik modul ini. Tanpa dependensi runtime baru.
- **Headless**: modul hanya menyediakan perilaku (state, gerakan, aksesibilitas), bukan tampilan. Tampilan (kartu, placeholder, overlay) disediakan pemakai dengan snippet/komponen Svelte; token desain mengikuti NFR-020 butir 8.
- **API Svelte 5**: primitif berbasis runes dan attachment (`{@attach}`), misalnya penyedia konteks drag, item yang dapat diangkat, dan zona drop, dengan opsi `type`/`accept` untuk membatasi sumber ke tujuan, `disabled` per item/zona, dan *drop validator* yang dapat menolak drop (mis. kolom tanpa transisi valid, BRD-016 butir 2 dan 18).
- **Input**: Pointer Events untuk mouse, sentuh, dan pena (mendukung drag handle dan ambang jarak sebelum drag dimulai); **keyboard** sesuai BRD-016 butir 22: fokus kartu, Space angkat, panah pindah, Space lepas, Esc batal.
- **Aksesibilitas**: region `aria-live` untuk pengumuman (angkat, pindah ke kolom/posisi, lepas, batal) dengan teks dari katalog i18n ID/EN (S0-13); instruksi tersedia untuk screen reader; fokus dipulihkan ke item setelah drop; dukungan `prefers-reduced-motion`. Alternatif non-drag (menu aksi "Pindahkan ke…", "Move to top/bottom", BRD-020) disediakan untuk pengguna yang tidak dapat melakukan drag.
- **Hasil drop sebagai event, bukan mutasi**: modul tidak mengubah data; ia melaporkan `{item, dari, ke, indeks}` kepada pemakai, yang mengirim transisi/rank ke API dengan optimistic update dan rollback saat ditolak (NFR-020 butir 3).
- **Virtualisasi**: perhitungan posisi berbasis data dan geometri item yang terukur, bukan urutan DOM, agar kolom tervirtualisasi (NFR-020 butir 4) tetap dapat menjadi tujuan drop; auto-scroll kontainer di tepi.
- **Multi-select** (BRD-016 butir 14, ⚠) dan layout horizontal/grid tidak termasuk cakupan awal, tetapi API tidak boleh menghalanginya.

## Konsekuensi

- Tidak menambah dependensi runtime; tidak ada risiko lisensi/pemeliharaan pihak ketiga untuk DnD.
- Efek ke jadwal: modul ini harus ada sebelum BRD-016/BRD-017 dan mendahului backlog (modul 11); estimasinya perlu dimasukkan ke perencanaan Fase 1. Risiko #9 README berubah dari "belum ada library" menjadi "bangun sendiri dengan target WCAG 2.1 AA".
- **Pengujian wajib**: unit test untuk logika posisi/drop (`vitest`), pengujian browser untuk pointer dan keyboard memakai `@vitest/browser-playwright` yang sudah ada di devDependencies, serta pemeriksaan manual dengan screen reader sebelum rilis board.
- **Syarat selesai modul** (diusulkan): drag pointer dan sentuh, drag keyboard lengkap, pengumuman ARIA ID/EN, drop validator, multi-list, auto-scroll, dan kompatibel dengan daftar tervirtualisasi, sebelum dipakai di fitur pertama.
- Bila biaya atau cacat aksesibilitas melebihi toleransi, fallback adalah mengadopsi `svelte-dnd-action` di balik API modul yang sama; karena fitur hanya memakai API `lib/dnd`, penggantian tidak mengubah pemakai. Pertimbangan ini dicatat sebagai ADR baru bila terjadi.
- Tindak lanjut: tiket desain API modul `lib/dnd` (S0-09 / Fase 1), tiket pengujian aksesibilitas, dan pembaruan Risiko #9 di README.

## Rujukan

- PRD: BRD-016 (Board Drag-and-Drop Transitions), BRD-017, BRD-020, NFR-020, NFR-022, NFR-001.
- `docs/prd/README.md` › Risiko #9.
- Dokumentasi pembanding: <https://github.com/isaacHagoel/svelte-dnd-action>, <https://github.com/hanielu/dnd-kit-svelte>, <https://github.com/atlassian/pragmatic-drag-and-drop>, <https://github.com/horuse/svelte-dnd>.
