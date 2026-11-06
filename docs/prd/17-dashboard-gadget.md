# 17 — Dashboard & Gadget

Modul ini mendefinisikan seluruh kemampuan dashboard di produk ini: direktori dan menu navigasi dashboard, bintang (favorit) dan daftar Recent, siklus hidup dashboard (buat, ubah detail, salin, hapus, system dashboard bawaan), model berbagi Viewers/Editors beserta global permission "Share dashboards and filters", akses publik/anonim, kewenangan admin atas dashboard milik siapa pun, view mode dan edit mode dengan lima pilihan layout kolom, katalog gadget dan penyusunan gadget dengan drag-and-drop, kerangka gadget bersama (konfigurasi umum, menu aksi, refresh otomatis, rendering yang selalu menghormati izin viewer, drill-down chart, konfigurasi gadget agile), katalog gadget bawaan (daftar work item personal dan berbasis filter, statistik dan distribusi, tren waktu, sprint dan rilis, aktivitas dan informasi), REST API dashboard dan gadget, titik ekstensi gadget dari app, audit dan batasan, serta akses mobile. Logika kalkulasi chart historis (burndown, created vs resolved, time in status) dimiliki modul 16 dan dipakai ulang oleh gadget di sini; definisi global permission dimiliki modul 03. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, middleware bahasa EN/ID), sehingga seluruh modul ini adalah pekerjaan baru yang dikemas sebagai feature module `internal/features/dashboard` (handler, DTO, use case) dengan entitas `Dashboard`, `DashboardSharePermission`, `DashboardGadget`, `DashboardItemProperty`, dan `DashboardFavourite` di `internal/shared/domain/entity`, endpoint di bawah `/api/v1/dashboards`, serta halaman SPA `/dashboards` dan `/dashboards/:id` di `frontend/`.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi "work item", plan tier, feature flag.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — halaman login untuk pengguna anonim (rujukan DSH-063).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — deaktivasi pengguna (USR-007, USR-008), global permission (USR-022), akses anonim (USR-024), audit log (USR-046).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Users and Groups (PERM-008), Share Dashboards and Filters (PERM-009), Browse Projects (PERM-020), Voters & Watchers (PERM-033), Issue Security (PERM-044), Anonymous & Public Access (PERM-046), Board & Filter Permissions (PERM-051), Deactivated Users in Grants (PERM-054).
- [04 Proyek](04-proyek.md) — project category, Starred & Recent Projects (PRJ-026), Versions (PRJ-055).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — custom field dengan statistics searcher, field configuration, Labels custom field (CFG-031), Issue features voting/watching (CFG-057).
- [06 Work Item Inti](06-work-item-inti.md) — Labels (ISS-020), Due Date (ISS-021), Watchers (ISS-036), Votes (ISS-037).
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — editor teks Introduction gadget (DSH-053).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — changelog dan history (CMT-022, CMT-023) sebagai sumber Activity Stream.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category, Resolution, status history (WF-014).
- [10 Board](10-board-kolom-kartu.md) — board Scrum/Kanban, Working days (BRD-040), Flag.
- [11 Backlog & Sprint](11-backlog-sprint.md) — sprint aktif, Sprint Start (SPR-018), parallel sprints.
- [12 JQL](12-jql-bahasa-kueri.md) — `currentUser()` (JQL-039), `watchedIssues()`/`votedIssues()` (JQL-041), klausa `filter =` (JQL-030).
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — saved filter, starred filter, Issue Navigator sebagai tujuan drill-down, konfigurasi kolom (SRC-025).
- [14 Pelacakan Waktu & Estimasi](14-pelacakan-waktu-estimasi.md) — estimation statistic (TIME-022), Workload pie chart (TIME-034), gadget berbasis waktu (TIME-035).
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — version, release date, Releases page (REL-012), progress bar (REL-013).
- [16 Laporan & Insight](16-laporan-insight.md) — mesin kalkulasi historis (RPT-004), estimation & working days (RPT-005), komponen chart (RPT-008), Burndown (RPT-009).
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Timeline/Plans dan Calendar view project (pembeda Road Map dan Work Items Calendar gadget).
- [19 Notifikasi & Email](19-notifikasi-email.md) — event dashboard dibagikan.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — sidebar, Starred, Recent, halaman awal, katalog keyboard shortcut.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — General configuration (Introduction, public sharing), System dashboard, Shared items.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — kontrak REST kompatibel, rate limit.
- [24 Integrasi DevOps](24-integrasi-devops.md) — informasi development pengganti gadget FishEye/Crucible/Bamboo (DSH-063).
- [25 Integrasi & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — titik ekstensi app (modul gadget), application link.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — cache, timeout query, aksesibilitas, responsif.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Dashboard, Filter, event domain `dashboard.*`.

## Persona terkait

- **Developer/Anggota Tim** — memakai dashboard pribadi berisi Assigned to Me, Issues in Progress, Watched Issues, dan Filter Results sebagai halaman kerja harian.
- **Scrum Master/Team Lead** — menyusun dashboard tim berisi Sprint Burndown, Sprint Health, Days Remaining in Sprint, dan Created vs Resolved, lalu membagikannya ke project.
- **Product Owner/PM** — memantau distribusi dan tren (Pie Chart, Two Dimensional Filter Statistics, Road Map) untuk keputusan prioritas dan rilis.
- **Admin Proyek/Project Lead** — membagikan dashboard ke project atau project role, mengelola editor dashboard tim.
- **Stakeholder/Viewer** — membuka dashboard yang dibagikan tanpa hak edit; hanya melihat data yang boleh ia browse.
- **Pengguna Tim Bisnis** — memakai dashboard berbasis filter untuk project bisnis tanpa board/sprint.
- **Admin Situs** — mengonfigurasi system dashboard dan teks Introduction, mengatur global permission berbagi dan public sharing, mengganti owner dan membersihkan dashboard pengguna nonaktif.
- **Integrator/Developer API** — mengelola dashboard dan gadget lewat REST API, membuat gadget dari app.
- **Tamu/Kolaborator Eksternal** — membuka dashboard Public secara anonim bila diizinkan admin.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin dashboard berisi gadget Assigned to Me dan Issues in Progress yang selalu dihitung untuk diri saya agar saya langsung tahu pekerjaan saya setelah login (DSH-027, DSH-028, DSH-024).
- Sebagai Scrum Master/Team Lead, saya ingin membuat dashboard, memilih layout tiga kolom, dan menyusun gadget dengan drag-and-drop agar satu halaman merangkum kesehatan sprint (DSH-005, DSH-018, DSH-020).
- Sebagai Scrum Master/Team Lead, saya ingin gadget Sprint Burndown, Sprint Health, dan Days Remaining in Sprint yang membaca board Scrum dan working days agar standup cukup melihat dashboard (DSH-047, DSH-048, DSH-049, DSH-026).
- Sebagai Admin Proyek/Project Lead, saya ingin membagikan dashboard ke project tertentu (opsional dibatasi project role) dan menunjuk editor agar tim dapat merawat dashboard bersama (DSH-010, DSH-011).
- Sebagai Stakeholder/Viewer, saya ingin dashboard yang dibagikan kepada saya hanya menampilkan work item yang boleh saya lihat agar tidak ada data rahasia yang bocor lewat angka statistik (DSH-024).
- Sebagai Product Owner/PM, saya ingin mengklik slice Pie Chart atau sel Two Dimensional Filter Statistics untuk membuka pencarian dengan JQL yang sesuai agar saya bisa menindaklanjuti angka (DSH-025, DSH-035, DSH-034).
- Sebagai Product Owner/PM, saya ingin gadget Created vs Resolved dan Road Map agar saya tahu apakah backlog membesar dan version mana yang akan rilis (DSH-039, DSH-050).
- Sebagai Developer/Anggota Tim, saya ingin menandai dashboard dengan bintang dan membukanya dari menu Dashboards agar akses cepat tanpa mencari (DSH-003, DSH-002).
- Sebagai Developer/Anggota Tim, saya ingin menyalin dashboard rekan atau system dashboard menjadi milik saya agar saya bisa memodifikasinya tanpa mengubah aslinya (DSH-007, DSH-009).
- Sebagai Admin Situs, saya ingin mengganti owner dashboard milik pengguna yang dinonaktifkan agar dashboard tim tetap terawat (DSH-014, DSH-015).
- Sebagai Admin Situs, saya ingin mengatur siapa yang boleh berbagi dashboard dan apakah dashboard publik diizinkan agar kebijakan keamanan organisasi terpenuhi (DSH-012, DSH-013).
- Sebagai Integrator/Developer API, saya ingin membuat dashboard dan menambahkan gadget beserta konfigurasinya lewat REST API agar dashboard standar dapat disiapkan otomatis untuk setiap tim baru (DSH-056, DSH-057).
- Sebagai Integrator/Developer API, saya ingin app saya menyediakan gadget dashboard sendiri yang tampil di panel Add gadget agar metrik khusus dapat ditampilkan di dashboard (DSH-058).
- Sebagai Tamu/Kolaborator Eksternal, saya ingin membuka dashboard publik tanpa login agar saya bisa memantau progres project open-source (DSH-013, DSH-016).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 20 |
| P1 | 14 |
| P2 | 16 |
| P3 | 13 |
| **Total** | **63** |

| Fase | Jumlah fitur |
|---|---|
| Fase 2 | 1 |
| Fase 3 | 37 |
| Fase 4 | 22 |
| Fase — (dc-only/deprecated, rujukan) | 3 |

Strategi urutan implementasi: dashboard termasuk cakupan Fase 3 (skala organisasi), sehingga seluruh fitur P0 dan P1 modul ini dijadwalkan di Fase 3, setelah saved filter, JQL, dan laporan agile Fase 2 tersedia sebagai sumber data gadget. Urutan internal Fase 3 dimulai dari model data dan izin (DSH-010, DSH-011, DSH-024) karena rendering gadget yang selalu dihitung dengan izin viewer tidak bisa disusulkan, lalu kerangka view/edit/layout/drag-and-drop dan konfigurasi gadget umum (DSH-016 s.d. DSH-023), lalu siklus hidup dan direktori, kemudian gadget P0 (Assigned to Me, Filter Results, Issue Statistics, Pie Chart) dan system dashboard bawaan beserta Introduction gadget yang menjadi bagian seed-nya (DSH-053, P2 yang dimajukan ke Fase 3). Pengecualian lain yang dimajukan ke Fase 3 adalah DSH-015 (P2), DSH-059 (P3), dan DSH-013 (P2, diselaraskan dengan PERM-046/USR-024 Fase 3) dengan alasan tertulis di masing-masing fitur. DSH-012 berstatus Fase 2 karena global permission-nya (PERM-009) sudah hadir sejak Fase 2 untuk filter sharing (SRC-045); dashboard baru memakainya di Fase 3. Gadget P1 (Created vs Resolved, Sprint Burndown, Two Dimensional Filter Statistics, Activity Stream, Issues in Progress) dan REST API dashboard/gadget menutup Fase 3. Fase 4 menambah gadget P2/P3, mobile, keyboard shortcut, dan titik ekstensi gadget app. Semua gadget chart memakai ulang komponen chart dan kontrak kalkulasi modul 16, sehingga satu perbaikan perhitungan berlaku sama di laporan dan dashboard.

## Fitur

## Direktori, navigasi & favorit

### DSH-001 Dashboards Directory (View all dashboards)

- **Deskripsi:** Halaman daftar semua dashboard yang dapat diakses pengguna, lengkap dengan pencarian, filter, pengurutan, dan aksi per dashboard. Menjadi titik masuk utama untuk menemukan, membuat, dan mengelola dashboard.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna login (melihat dashboard miliknya dan yang dibagikan kepadanya), owner/editor dashboard (aksi Edit), Admin Situs dengan Administer Jira (melihat semua dashboard dan aksi Change owner).
- **Aturan perilaku (acceptance criteria):**
  1. Direktori tersedia di route SPA `/dashboards` dan dibuka dari menu Dashboards > "View all dashboards"; URL kompatibilitas `/jira/dashboards` di-redirect 301 ke `/dashboards`.
  2. Kotak pencarian mencocokkan nama dashboard secara parsial dan case-insensitive (mis. "sprint" menemukan "Sprint Tim A"); spasi di awal/akhir diabaikan dan pencarian dijalankan setelah jeda ketik 300 ms.
  3. Filter Owner (user picker), Project (dashboard yang dibagikan ke project tersebut, termasuk lewat project role), dan Group (dashboard yang dibagikan ke group tersebut) dapat dikombinasikan dan digabung dengan logika AND bersama kata kunci pencarian.
  4. Tabel menampilkan kolom Starred (ikon bintang), Name (dengan description di bawahnya, dipotong 1 baris dengan tooltip penuh), Owner (avatar + nama), Viewers (ringkasan entri share), Editors (ringkasan entri edit), dan Popularity/Starred by (jumlah pengguna yang men-star).
  5. Kolom Name, Owner, dan Popularity dapat diurutkan ascending/descending dengan klik header; urutan default Name ascending dan indikator panah menunjukkan arah aktif.
  6. ⚠ Hasil dipaginasi dengan ukuran halaman default 20 baris dan teks "Menampilkan 1-20 dari N"; navigasi halaman berikut/sebelumnya dan nomor halaman tersedia.
  7. Ikon bintang di setiap baris dapat di-toggle langsung tanpa membuka dashboard; perubahan diterapkan optimistis dan dibatalkan dengan toast galat bila server menolak.
  8. Menu "..." per baris berisi Edit (nama/deskripsi/izin, DSH-006), Copy (DSH-007), dan Delete (DSH-008); Change owner (DSH-014) hanya muncul untuk Admin Situs.
  9. Aksi Edit hanya muncul bagi owner dan editor dashboard; Delete hanya bagi owner; Admin Situs melihat semua aksi; pengguna yang hanya viewer hanya melihat Copy.
  10. Tombol "Create dashboard" di kanan atas membuka dialog pembuatan (DSH-005) dan tidak tampil untuk pengguna anonim.
  11. Empty state dibedakan: "Belum ada dashboard" dengan CTA Create bila pengguna tidak memiliki/menerima dashboard apa pun, dan "Tidak ada dashboard yang cocok" dengan tombol Reset filter bila pencarian/filter tidak memberi hasil.
  12. Non-admin hanya melihat dashboard miliknya atau yang dibagikan kepadanya sebagai Viewer atau Editor; dashboard private milik orang lain tidak pernah muncul, termasuk pada hitungan total.
  13. Admin Situs melihat semua dashboard termasuk private milik orang lain (setara `extendAdminPermissions=true`), dengan lencana "Private" pada dashboard yang tidak dibagikan kepadanya.
  14. System dashboard (DSH-009) tampil di daftar dengan lencana "Default" dan tidak memiliki aksi Delete bagi siapa pun.
  15. ⚠ State pencarian, filter, urutan, dan halaman disimpan di query string (mis. `/dashboards?q=sprint&owner=<id>&project=ABC&sort=-popularity&page=2`) sehingga reload dan tautan bersama mempertahankan tampilan.
- **Team-managed vs company-managed:** Tidak ada perbedaan; filter Project mencakup kedua jenis project.
- **Catatan & ketergantungan:** Endpoint `GET /api/v1/dashboards/search` dipakai bersama REST API (DSH-056). Di produk acuan DC halaman ini bertab Favourite/My/Popular/Search; produk ini mengikuti model satu halaman dengan filter. Teks UI di katalog pesan EN/ID (`internal/shared/constants`).

### DSH-002 Dashboards navigation menu

- **Deskripsi:** Menu Dashboards pada navigasi global/sidebar yang menampilkan dashboard berbintang dan yang baru dilihat untuk akses cepat. Modul ini hanya mengatur perilaku khusus dashboard; struktur navigasi global ada di modul 21.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / platform
- **Aktor & izin:** Semua pengguna login; isi menu dibatasi pada dashboard yang masih dapat dilihat pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Item "Dashboards" tersedia di sidebar navigasi global ([21 Navigasi](21-navigasi-ux-preferensi.md)); klik item membuka flyout/menu, bukan langsung berpindah halaman.
  2. Bagian Starred menampilkan dashboard yang di-star pengguna, diurutkan alfabetis; bila kosong tampil teks "Beri bintang dashboard agar muncul di sini".
  3. Bagian Recent menampilkan dashboard yang terakhir dibuka di view mode, diurutkan dari waktu akses terbaru; akses dicatat per pengguna saat dashboard berhasil dimuat (bukan saat galat akses).
  4. Tautan "View all dashboards" membuka direktori (DSH-001).
  5. Aksi "Create dashboard" membuka dialog pembuatan (DSH-005) di atas halaman aktif tanpa berpindah halaman.
  6. ⚠ Setiap item di menu memiliki ikon bintang yang muncul saat hover/fokus dan dapat di-toggle langsung; item yang di-unstar tetap terlihat di Recent.
  7. ⚠ Bagian Recent menampilkan maksimal 5 dashboard; dashboard yang dihapus atau tidak lagi dapat diakses dikeluarkan dari daftar.
  8. Dashboard berbintang juga muncul di bagian Starred sidebar bersama project, filter, dan board, memakai ikon dashboard sebagai pembeda.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan Recent memakai mekanisme recent items bersama (SRC-011 di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md)); tampilannya di sidebar diatur modul 21. Bintang memakai DSH-003.

### DSH-003 Star dashboard (Favourites)

- **Deskripsi:** Menandai dashboard sebagai favorit (star) agar muncul di menu Starred dan dihitung dalam popularitas.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna login yang dapat melihat dashboard; pengguna anonim tidak dapat men-star.
- **Aturan perilaku (acceptance criteria):**
  1. Ikon bintang tersedia di header dashboard (view mode) dan di setiap baris direktori; status terisi/kosong menunjukkan apakah pengguna sudah men-star.
  2. Dashboard berbintang muncul di bagian Starred menu Dashboards (DSH-002) dan di bagian Starred sidebar.
  3. Popularity sebuah dashboard sama dengan jumlah pengguna unik yang men-star dashboard tersebut dan diperbarui segera setelah star/unstar.
  4. Unstar menghapus dashboard dari daftar Starred pengguna dan mengurangi Popularity sebesar 1.
  5. Star bersifat per pengguna; star oleh satu pengguna tidak memengaruhi daftar Starred pengguna lain.
  6. Hanya dashboard yang dapat dilihat pengguna yang dapat di-star; bila akses dicabut, star tetap tersimpan tetapi dashboard disembunyikan dari daftar Starred sampai akses kembali, dan permintaan star ke dashboard tanpa akses ditolak 404.
  7. Pengurutan dashboard favorit secara manual sebagai tab (perilaku DC) tidak dibangun; daftar Starred diurutkan alfabetis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `dashboard_favourites (dashboard_id, user_id, created_at)` dengan unique key; penghapusan dashboard menghapus semua star (DSH-008).

### DSH-004 Dashboard keyboard shortcuts

- **Deskripsi:** Shortcut keyboard global untuk menuju dashboard.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna login dengan keyboard shortcut aktif di preferensi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Urutan tombol `g` lalu `d` (dalam 1 detik) membuka dashboard terakhir yang dilihat atau, bila tidak ada, system dashboard; shortcut tidak aktif saat fokus berada di input teks atau editor.
  2. Shortcut tercantum di dialog Keyboard shortcuts (`?`) pada kelompok "Global" dengan label "Buka dashboard".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog shortcut dan preferensi on/off dimiliki modul 21.

## Siklus hidup dashboard

### DSH-005 Create dashboard

- **Deskripsi:** Dialog untuk membuat dashboard baru kosong dengan nama, deskripsi, serta pengaturan Viewers dan Editors. Setelah dibuat, pengguna langsung masuk ke edit mode untuk menambah gadget.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna login; opsi berbagi selain Private memerlukan global permission Share dashboards and filters (DSH-012).
- **Aturan perilaku (acceptance criteria):**
  1. Field Name wajib diisi, maksimal 255 karakter setelah trim; nama kosong atau hanya spasi ditolak dengan pesan "Nama dashboard wajib diisi" dan nama terlalu panjang dengan "Nama maksimal 255 karakter".
  2. ⚠ Nama harus unik per owner (perbandingan case-insensitive); bila owner sudah memiliki dashboard dengan nama sama, dialog menampilkan galat inline "Anda sudah memiliki dashboard dengan nama ini" dan API mengembalikan 400.
  3. Field Description opsional berupa plain text, maksimal 1.000 karakter; markup tidak dirender.
  4. Bagian Viewers dan Editors tersedia dengan nilai default Private untuk keduanya; konfigurasinya mengikuti DSH-010 dan DSH-011.
  5. Pembuat otomatis menjadi Owner dan tidak dapat memilih owner lain di dialog ini.
  6. ⚠ Dashboard baru otomatis di-star untuk pembuatnya sehingga langsung muncul di menu Starred.
  7. Setelah tombol Create ditekan dan berhasil, pengguna diarahkan ke `/dashboards/{id}` dalam edit mode dengan panel Add gadget terbuka (DSH-017, DSH-019).
  8. Layout awal dashboard baru adalah dua kolom sama lebar (AA, DSH-018).
  9. Pengguna tanpa Share dashboards and filters hanya dapat membuat dashboard Private; kontrol Viewers/Editors tidak tampil dan permintaan API dengan share selain Private ditolak 403.
  10. Pengguna anonim tidak dapat membuat dashboard; tombol Create tidak tampil dan API mengembalikan 401.
  11. Tombol Cancel atau Esc menutup dialog tanpa menyimpan apa pun.
  12. Opsi "Start from" (blank atau salin dashboard lain) di dialog create seperti DC tidak dibangun; penyalinan dilakukan lewat aksi Copy (DSH-007).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** `POST /api/v1/dashboards`; event domain `dashboard.created` (modul 28). Batas jumlah dashboard per pengguna mengikuti DSH-059.

### DSH-006 Edit dashboard details (Rename & share)

- **Deskripsi:** Mengubah nama, deskripsi, Viewers, dan Editors dashboard yang sudah ada melalui dialog "Edit and share dashboard".
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog edit dibuka dari menu "..." di header dashboard (DSH-016) atau dari menu aksi baris direktori (DSH-001), terisi nilai saat ini.
  2. Name dapat diubah dengan validasi yang sama dengan pembuatan: wajib, maksimal 255 karakter, dan unik per owner (yang dicek adalah owner dashboard, bukan pengguna yang mengedit).
  3. Description dapat diubah atau dikosongkan.
  4. Viewers dan Editors dapat diubah mengikuti aturan DSH-010 dan DSH-011; perubahan izin berlaku segera untuk permintaan berikutnya.
  5. Hanya Owner, Editor, atau Admin Situs yang dapat membuka dialog; permintaan lain ditolak 403 dan menu tidak menampilkan aksi Edit.
  6. Perubahan nama langsung tercermin di menu Starred/Recent, direktori, dan judul tab browser tanpa reload penuh.
  7. Owner tidak dapat diubah dari dialog ini; field owner tampil read-only, dan perpindahan owner hanya lewat Change owner oleh Admin Situs (DSH-014).
  8. System dashboard tidak dapat di-rename atau dibagikan ulang oleh pengguna biasa; hanya Admin Situs yang dapat mengubah nama dan deskripsinya, sedangkan sharing-nya tetap terkunci ke semua pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** `PUT /api/v1/dashboards/{id}`; event `dashboard.updated` dicatat di audit (DSH-059).

### DSH-007 Copy dashboard

- **Deskripsi:** Menduplikasi dashboard beserta layout dan seluruh konfigurasi gadget menjadi dashboard baru milik pengguna yang menyalin.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna login yang dapat melihat dashboard sumber.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Copy tersedia di menu "..." header dashboard dan di menu baris direktori.
  2. Dialog Copy berisi Name (default "Copy of <nama asli>", dipotong agar tidak melebihi 255 karakter), Description (default salinan deskripsi asli), Viewers, dan Editors.
  3. Seluruh gadget, posisi (kolom, baris), layout, judul, warna, dan item properties (konfigurasi) disalin sebagai instance gadget baru dengan ID baru.
  4. Pengguna yang menyalin menjadi owner salinan.
  5. ⚠ Viewers dan Editors salinan default Private; sharing dashboard sumber tidak ikut tersalin, dan pengguna dapat mengaturnya di dialog Copy bila memiliki Share dashboards and filters.
  6. Pengguna dapat menyalin dashboard milik orang lain, termasuk system dashboard, selama ia dapat melihatnya.
  7. Nama salinan divalidasi unik per owner baru; bila bentrok, dialog menampilkan galat dan pengguna harus mengganti nama.
  8. Filter, board, dan project yang direferensikan gadget tidak disalin; konfigurasi salinan tetap merujuk objek asli sehingga gadget tunduk pada izin filter asli (DSH-024).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** `POST /api/v1/dashboards/{id}/copy` (DSH-056). Setelah sukses, pengguna diarahkan ke salinan dalam view mode.

### DSH-008 Delete dashboard

- **Deskripsi:** Menghapus dashboard secara permanen beserta semua gadget-nya.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner dashboard dan Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Delete tersedia di menu "..." header dashboard dan di menu baris direktori.
  2. Dialog konfirmasi menampilkan nama dashboard dan teks "Tindakan ini permanen dan tidak dapat dibatalkan"; tombol Delete berwarna destruktif dan Cancel menutup dialog.
  3. Hanya owner dan Admin Situs yang dapat menghapus; Editor dan Viewer tidak melihat aksi ini dan API mengembalikan 403.
  4. System dashboard tidak dapat dihapus oleh siapa pun; API mengembalikan 400 dengan pesan "System dashboard tidak dapat dihapus".
  5. Penghapusan menghapus star dari semua pengguna dan mengeluarkan dashboard dari daftar Recent.
  6. ⚠ Penghapusan bersifat permanen tanpa trash/restore di UI; baris dashboard, gadget, item properties, dan share permission dihapus dalam satu transaksi, dan REST tidak menyediakan status "deleted" yang dapat dipulihkan.
  7. Setelah penghapusan berhasil, pengguna diarahkan ke direktori dengan toast "Dashboard <nama> dihapus".
  8. Filter yang dipakai gadget tidak ikut terhapus.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** `DELETE /api/v1/dashboards/{id}`; event `dashboard.deleted` dicatat di audit (DSH-059). Pengguna yang membuka URL dashboard terhapus melihat pesan tidak ditemukan (DSH-016).

### DSH-009 System default dashboard

- **Deskripsi:** Dashboard bawaan sistem yang ditampilkan kepada pengguna yang belum memiliki dashboard sendiri dan dikelola oleh Admin Situs.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengonfigurasi), semua pengguna (melihat dan menyalin), pengguna anonim bila akses anonim aktif.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs mengonfigurasi system dashboard dari Settings > System > User interface > System dashboard ([22 Administrasi Sistem](22-administrasi-sistem.md)), yang membuka dashboard tersebut dalam edit mode.
  2. Instance baru membuat system dashboard dengan gadget Introduction (DSH-053), Assigned to Me (DSH-027), dan Activity Stream (DSH-052) pada layout AA; seed dijalankan lewat migrasi data awal.
  3. System dashboard dibagikan ke semua pengguna login dan tidak dapat dihapus.
  4. Owner system dashboard tidak dapat diubah, dan hanya Admin Situs yang dapat mengubah layout serta gadget-nya; pengguna biasa tidak melihat tombol Edit.
  5. Pengguna biasa dapat menyalin system dashboard (DSH-007) sebagai dasar dashboard pribadi.
  6. Perubahan pada system dashboard langsung berlaku bagi semua pengguna yang membukanya; data gadget tetap dihitung per viewer (DSH-024).
  7. ⚠ Bila akses anonim situs aktif, pengguna anonim yang membuka `/dashboards` melihat system dashboard dengan gadget yang membutuhkan login menampilkan pesan login (DSH-013).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan sebagai baris `dashboards` dengan flag `is_system=true` (tepat satu baris). Tujuan default menu Dashboards untuk pengguna tanpa riwayat adalah system dashboard (DSH-016).

## Berbagi, izin & kepemilikan

### DSH-010 Dashboard sharing (Viewers)

- **Deskripsi:** Mengatur siapa saja yang dapat melihat dashboard. Mendukung beberapa entri share sekaligus dengan berbagai tipe target.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs; opsi selain Private memerlukan Share dashboards and filters (PERM-009, DSH-012).
- **Aturan perilaku (acceptance criteria):**
  1. Tipe Private: hanya owner dan pengguna yang terdaftar sebagai Editor yang dapat melihat dashboard; memilih Private menghapus semua entri Viewers lain.
  2. Tipe "My organization": semua pengguna yang login dan memiliki akses produk di situs (padanan "Any logged-in user" lama) dapat melihat dashboard.
  3. Tipe Group: memilih satu group; anggota group saat ini (dievaluasi pada setiap akses) dapat melihat dashboard.
  4. Tipe Project: semua pengguna dengan Browse Projects pada project terpilih dapat melihat; opsional dibatasi ke satu Project role sehingga hanya anggota role itu di project tersebut yang dapat melihat.
  5. Tipe User: pengguna spesifik yang dipilih lewat user picker.
  6. Tipe Public: siapa pun di web termasuk pengguna anonim; opsi ini hanya tampil bila public sharing diizinkan Admin Situs (DSH-013).
  7. Beberapa entri share dapat ditambahkan sebagai baris terpisah (mis. Group A + Project B + User C) dan setiap baris dapat dihapus satu per satu; entri duplikat ditolak dengan pesan "Entri ini sudah ada".
  8. ⚠ Pengguna non-admin hanya dapat membagikan ke group yang ia anggotai dan ke project yang dapat ia browse; picker hanya menawarkan target tersebut dan API menolak target lain dengan 400.
  9. Opsi selain Private memerlukan global permission Share dashboards and filters; tanpa izin itu hanya Private yang tersedia.
  10. Membagikan dashboard tidak otomatis membagikan filter yang dipakai gadget; dialog menampilkan catatan "Viewer tetap memerlukan akses ke filter yang dipakai gadget" dan gadget tetap menghormati izin filter (DSH-024).
  11. ⚠ Saat memilih Public, dialog menampilkan peringatan "Siapa pun di internet dapat melihat dashboard ini" yang harus dikonfirmasi sebelum disimpan.
  12. Bila project atau group yang menjadi target share dihapus, entri share tersebut ikut dihapus dan audiens dashboard menyempit tanpa notifikasi ke owner.
  13. ⚠ Share ke project team-managed berlaku untuk semua pengguna yang memiliki akses project, tetapi role team-managed (Administrator/Member/Viewer) tidak ditawarkan sebagai pilihan Project role.
- **Team-managed vs company-managed:** Company-managed: Project dapat dibatasi ke project role dari permission scheme. Team-managed: hanya share ke seluruh project; role team-managed tidak tersedia sebagai pembatas.
- **Catatan & ketergantungan:** Tabel `dashboard_share_permissions (dashboard_id, kind[view|edit], type, project_id, role_id, group_id, user_id)` dipakai bersama tipe REST `user, group, project, projectRole, global, loggedin, authenticated, project-unknown` (DSH-056). Model share sama dengan saved filter (modul 13) dan sebaiknya diimplementasikan sebagai komponen bersama. Tipe Public (butir 6) baru dapat dipilih setelah pengaturan public sharing DSH-013 (Fase 4) tersedia; sebelum itu opsi tidak ditampilkan.

### DSH-011 Dashboard edit permissions (Editors)

- **Deskripsi:** Mengatur siapa yang dapat mengubah dashboard (layout, gadget, detail). Editor dapat ditetapkan ke user, group, project, atau project role.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner dan Editor dashboard; penetapan editor selain Private memerlukan Share dashboards and filters.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi Editors Private berarti hanya owner (dan Admin Situs) yang dapat mengedit; ini nilai default.
  2. Editor dapat berupa User spesifik.
  3. Editor dapat berupa Group; keanggotaan dievaluasi pada setiap akses.
  4. Editor dapat berupa Project, opsional dibatasi ke satu Project role.
  5. Beberapa entri editor dapat ditambahkan dan dihapus satu per satu, dengan validasi duplikat yang sama seperti Viewers.
  6. Editor dapat menambah, menghapus, dan mengonfigurasi gadget, mengubah layout, serta mengubah nama, deskripsi, dan sharing dashboard.
  7. ⚠ Editor tidak dapat menghapus dashboard; aksi Delete hanya untuk owner dan Admin Situs.
  8. Editor tidak dapat mengubah owner.
  9. Setiap editor otomatis juga memiliki akses view, meskipun tidak tercantum di Viewers.
  10. Pengguna yang hanya viewer tidak melihat tombol Edit maupun aksi konfigurasi gadget; permintaan API ubah dari viewer ditolak 403.
  11. ⚠ Opsi Public dan "My organization" tidak tersedia untuk Editors; editor selalu harus ditetapkan ke target spesifik (user, group, project, project role).
- **Team-managed vs company-managed:** Sama seperti DSH-010: project role hanya tersedia untuk project company-managed.
- **Catatan & ketergantungan:** Disimpan di tabel yang sama dengan Viewers dengan `kind=edit`. Pemeriksaan izin terpusat di use case `CanEditDashboard(user, dashboard)` agar konsisten antara UI dan API.

### DSH-012 Share dashboards and filters global permission

- **Deskripsi:** Global permission yang menentukan apakah pengguna boleh membagikan dashboard dan filter ke pengguna lain, dikelola Admin Situs. Definisi permission dimiliki PERM-009; fitur ini mengatur dampaknya pada dashboard.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (memberi/mencabut), semua pengguna (terkena dampak).
- **Aturan perilaku (acceptance criteria):**
  1. Dikelola di Settings > System > Global permissions dengan nama "Share dashboards and filters" (padanan "Create shared objects" di DC).
  2. Diberikan ke group; nilai default instance baru adalah group semua pengguna produk.
  3. Tanpa permission ini pengguna hanya dapat membuat dashboard Private, tidak melihat kontrol Viewers/Editors selain Private, dan API menolak share selain Private dengan 403.
  4. ⚠ Bila permission dicabut, dashboard yang sudah dibagikan tetap terbagikan; pengguna tersebut masih dapat menghapus entri share tetapi tidak dapat menambah entri baru.
  5. Global permission Browse users and groups (PERM-008) diperlukan agar user/group picker di dialog share dapat mencari pengguna dan group; tanpa itu picker hanya menerima input yang cocok persis dengan email/nama pengguna yang dapat browse project bersama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan utama [03 Izin](03-izin-peran-keamanan.md) PERM-009 dan PERM-008. Berlaku juga untuk saved filter (modul 13, SRC-045 Fase 2). Fase 2 diselaraskan dengan PERM-009: permission sudah ada dan ditegakkan untuk filter sharing sejak Fase 2, sedangkan penerapannya pada kontrol share dashboard (DSH-010, DSH-011) baru aktif saat modul dashboard dibangun di Fase 3.

### DSH-013 Public & anonymous dashboard access

- **Deskripsi:** Kemampuan berbagi dashboard ke publik/anonim dan pengaturan admin yang mengizinkan atau melarangnya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengatur), owner dashboard (membagikan), pengguna anonim (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Opsi share "Public" hanya tersedia bila Admin Situs mengizinkan public sharing dan akses anonim situs aktif; bila tidak, opsi tidak tampil dan API menolak tipe `global` dengan 400.
  2. ⚠ Admin Situs dapat menonaktifkan public sharing; entri share Public yang sudah ada tetap tersimpan tetapi berhenti berlaku (pengguna anonim mendapat halaman login) sampai public sharing diaktifkan kembali.
  3. Pengguna anonim hanya melihat data project yang Browse Projects-nya diberikan ke Public/Anyone (PERM-046).
  4. Gadget yang membutuhkan konteks pengguna (Assigned to Me, Issues in Progress, Watched, Voted, Favourite Filters) menampilkan pesan "Masuk untuk melihat gadget ini" beserta tautan login bagi pengguna anonim.
  5. Saat dashboard dijadikan Public, peringatan keamanan ditampilkan dan harus dikonfirmasi (sama dengan DSH-010 butir 11).
  6. ⚠ Pengaturan public sharing berada di Settings > System > General configuration (modul 22) bersama pengaturan akses anonim; pengaturan juga dapat diblokir oleh data security policy (PERM-065).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk acuan Cloud memperketat public sharing sejak 2021 (opsi Public lama diganti "My organization"); produk ini menyediakan Public sebagai opsi yang mati secara default. Endpoint dashboard anonim harus melewati rate limit khusus (modul 27). Fase 3 diselaraskan dengan PERM-046 Anonymous & Public Access ([03 Izin](03-izin-peran-keamanan.md)) dan USR-024 Anonymous & public access ([02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md)) yang menyediakan pengaturan akses anonim situs; walau P2, fitur ini dimajukan ke Fase 3 agar kontrol akses anonim dan opsi share Public dirilis bersamaan dan tidak ada celah perilaku antar-modul.

### DSH-014 Admin dashboard management (Change owner)

- **Deskripsi:** Kemampuan Admin Situs melihat, mengedit, mengubah owner, dan menghapus dashboard milik siapa pun, termasuk yang private.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin melihat semua dashboard di direktori (DSH-001) dan juga lewat halaman Settings > System > Shared items > Dashboards (modul 22) yang menampilkan dashboard semua pengguna.
  2. Aksi Change owner memindahkan kepemilikan ke pengguna aktif lain yang dipilih lewat user picker; owner lama kehilangan hak owner (kecuali tetap tercantum sebagai editor/viewer).
  3. Admin dapat mengubah Viewers dan Editors dashboard milik orang lain.
  4. Admin dapat menghapus dashboard milik siapa pun kecuali system dashboard.
  5. ⚠ Admin dapat memilih banyak dashboard di halaman Shared items lalu menjalankan bulk Change owner atau bulk ubah permission; operasi yang sama tersedia lewat `PUT /api/v1/dashboards/bulk/edit` (DSH-056) dan hasilnya dilaporkan per dashboard (sukses/gagal beserta alasan).
  6. ⚠ Bila owner baru tidak memiliki akses ke project/group yang menjadi target share, dialog menampilkan peringatan "Owner baru tidak dapat melihat sebagian target share" tetapi perubahan tetap diizinkan.
  7. Halaman admin mendukung pencarian berdasarkan nama dashboard dan owner.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validasi nama unik per owner berlaku untuk owner baru; bila bentrok, perubahan ditolak dengan pesan galat. Perubahan owner dicatat di audit (DSH-059).

### DSH-015 Dashboards of deactivated users

- **Deskripsi:** Penanganan dashboard yang owner-nya dinonaktifkan atau dihapus dari situs.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs; pengguna lain tetap melihat sesuai sharing.
- **Aturan perilaku (acceptance criteria):**
  1. Dashboard milik pengguna yang dinonaktifkan tetap ada dan dapat diakses sesuai sharing; tidak ada penghapusan otomatis.
  2. Owner ditampilkan sebagai "<nama> (nonaktif)" bila dinonaktifkan atau "Former user" bila akun dihapus, di direktori dan header dashboard.
  3. Admin Situs dapat memindahkan kepemilikan lewat Change owner (DSH-014); direktori admin menyediakan filter "Owner nonaktif".
  4. Gadget berbasis `currentUser()` tetap berfungsi karena dihitung per viewer, bukan per owner.
  5. ⚠ Filter milik pengguna nonaktif yang dipakai gadget tetap berjalan bila filter tersebut dibagikan ke viewer; filter private milik owner nonaktif membuat gadget menampilkan galat "Filter tidak tersedia" bagi viewer lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase 3 dipilih karena deaktivasi pengguna (USR-007, USR-008) sudah tersedia lebih dulu dan dashboard tim tidak boleh rusak saat itu terjadi. Lihat juga PERM-054.

## Tampilan & penyusunan dashboard

### DSH-016 Dashboard view mode

- **Deskripsi:** Tampilan dashboard saat dibuka, merender semua gadget sesuai layout dengan header berisi nama dan aksi.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat dashboard (owner, editor, viewer), pengguna anonim untuk dashboard Public.
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan nama dashboard, ikon bintang, tombol Refresh, tombol Edit (hanya bila berhak edit), dan menu "..." berisi Edit and share, Copy, Delete (owner/admin), serta Change owner (admin).
  2. Gadget dimuat secara lazy dan paralel; setiap gadget menampilkan spinner sendiri dan gadget di luar viewport dimuat saat mendekati viewport.
  3. Galat pada satu gadget (timeout, konfigurasi tidak valid, galat server) hanya ditampilkan di dalam gadget tersebut; gadget lain dan header tetap berfungsi.
  4. Tombol Refresh di header me-reload data semua gadget tanpa me-reload halaman.
  5. Deep link `/dashboards/{id}` membuka dashboard tertentu; URL legacy `/secure/Dashboard.jspa?selectPageId={id}` dan `/jira/dashboards/{id}` di-redirect ke URL baru.
  6. ⚠ Membuka `/dashboards/last-visited` atau item Dashboards tanpa id membuka dashboard terakhir yang dilihat pengguna; bila belum ada riwayat, system dashboard dibuka.
  7. Membuka dashboard tanpa izin menampilkan halaman "Anda tidak memiliki akses ke dashboard ini" dengan tautan ke direktori; respons API 404 (bukan 403) agar keberadaan dashboard private tidak bocor.
  8. Membuka dashboard yang sudah dihapus menampilkan pesan "Dashboard tidak ditemukan" dengan tautan ke direktori.
  9. Pada layar sempit (lebar < 768 px) semua kolom ditumpuk menjadi satu kolom dengan urutan kolom 1, lalu kolom 2, lalu kolom 3.
  10. Pengguna anonim dapat membuka dashboard yang dibagikan Public (DSH-013); header tidak menampilkan bintang, Edit, Copy, atau Delete.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** `GET /api/v1/dashboards/{id}` mengembalikan metadata, layout, dan daftar gadget; data tiap gadget diambil lewat endpoint gadget terpisah agar paralel. Membuka dashboard mencatat Recent (DSH-002).

### DSH-017 Dashboard edit mode

- **Deskripsi:** Mode edit untuk menyusun dashboard: menambah gadget, mengubah layout, memindahkan, dan mengonfigurasi gadget, diakhiri dengan tombol Done.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Edit mode dimasuki lewat tombol Edit di header; panel samping "Add gadget" (DSH-019) tampil dan URL berubah menjadi `/dashboards/{id}/edit`.
  2. Tombol "Change layout" tersedia di edit mode dan membuka pilihan layout (DSH-018).
  3. Setiap kolom menampilkan drop zone bergaris putus-putus tempat gadget dapat ditaruh, termasuk kolom kosong.
  4. Gadget yang baru ditambahkan langsung menampilkan form konfigurasi inline (DSH-021).
  5. Tombol Done keluar dari edit mode kembali ke view mode; form konfigurasi yang belum disimpan memicu konfirmasi "Buang perubahan konfigurasi?".
  6. Setiap aksi (tambah, pindah, hapus gadget, ubah layout, simpan konfigurasi) disimpan langsung ke server; tidak ada draft atau publish dashboard.
  7. ⚠ Dua editor dapat mengedit bersamaan tanpa locking; perubahan terakhir yang tersimpan menang per gadget/posisi, dan editor lain melihat perubahan setelah refresh atau menerima banner "Dashboard diubah oleh pengguna lain, muat ulang".
  8. Pengguna tanpa hak edit tidak melihat tombol Edit; membuka URL `/edit` langsung mengarah ke view mode dengan toast "Anda tidak memiliki izin mengedit dashboard ini".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Deteksi perubahan bersamaan memakai kolom `version` pada dashboard (optimistic concurrency untuk layout) dan event realtime bila tersedia (modul 27).

### DSH-018 Dashboard layouts

- **Deskripsi:** Pilihan tata letak kolom dashboard yang menentukan jumlah dan lebar kolom tempat gadget diletakkan.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Layout satu kolom (A): satu kolom penuh.
  2. Layout dua kolom sama lebar (AA) adalah default dashboard baru.
  3. Layout dua kolom dengan kiri sempit dan kanan lebar (BA, rasio sekitar 1:2).
  4. Layout dua kolom dengan kiri lebar dan kanan sempit (AB, rasio sekitar 2:1).
  5. Layout tiga kolom sama lebar (AAA).
  6. Saat beralih ke layout dengan kolom lebih sedikit, gadget dari kolom yang hilang dipindahkan ke bawah kolom terakhir yang tersisa dengan urutan relatif tetap; tidak ada gadget yang terhapus.
  7. Gadget menyesuaikan lebar kolom; chart di-render ulang mengikuti lebar baru tanpa distorsi.
  8. Layout disimpan per dashboard dan berlaku sama untuk semua viewer.
  9. ⚠ Kode layout disimpan dan diekspos di API produk ini sebagai enum `A`, `AA`, `BA`, `AB`, `AAA`; endpoint kompatibilitas mengikuti kontrak produk acuan yang tidak mengekspos layout secara eksplisit.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Posisi gadget disimpan sebagai `(column, row)`; validasi kolom terhadap layout ada di DSH-057.

### DSH-019 Add gadget panel (Gadget directory)

- **Deskripsi:** Katalog gadget yang tersedia untuk ditambahkan ke dashboard, mencakup gadget bawaan produk ini dan gadget dari app terpasang.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs (dalam edit mode).
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan daftar gadget dengan nama, deskripsi singkat, dan ikon/thumbnail.
  2. Pencarian gadget berdasarkan nama (parsial, case-insensitive) memfilter daftar secara langsung.
  3. ⚠ Gadget dikelompokkan dalam kategori (Daftar work item, Chart & statistik, Agile, Aktivitas & informasi, App) yang dapat dipakai sebagai filter.
  4. Gadget dapat ditambahkan dengan tombol Add pada kartunya atau dengan drag dari panel ke kolom tertentu.
  5. Gadget yang ditambahkan lewat tombol Add diletakkan di posisi teratas kolom pertama dan menggeser gadget lain ke bawah.
  6. Gadget dari app terpasang (DSH-058) muncul di panel dengan label nama app.
  7. Gadget agile (Sprint Burndown, Sprint Health, Days Remaining in Sprint) hanya tampil bila fitur board/sprint tersedia di situs.
  8. ⚠ Gadget legacy/deprecated tidak ditampilkan di panel sehingga tidak dapat ditambahkan lagi, tetapi instance lama yang sudah ada tetap dirender dengan lencana "Deprecated".
  9. Gadget yang sama dapat ditambahkan berkali-kali ke satu dashboard dengan konfigurasi berbeda.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog berasal dari registry gadget di backend (`GET /api/v1/dashboards/gadgets`, DSH-057) agar gadget bawaan dan app memakai kontrak yang sama.

### DSH-020 Gadget drag-and-drop arrangement

- **Deskripsi:** Memindahkan posisi gadget antar kolom dan mengubah urutan dalam kolom dengan drag-and-drop.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Header gadget menampilkan drag handle di edit mode; di view mode handle tidak tampil.
  2. Gadget dapat diurutkan ulang secara vertikal dalam satu kolom.
  3. Gadget dapat dipindahkan ke kolom lain.
  4. Selama drag, placeholder seukuran gadget menunjukkan posisi drop.
  5. Posisi baru `(column, row)` semua gadget yang terdampak disimpan permanen segera setelah drop; bila penyimpanan gagal, gadget kembali ke posisi semula dengan toast galat.
  6. Gadget dapat di-drop ke kolom kosong.
  7. Hanya pengguna dengan hak edit yang dapat melakukan drag; viewer tidak dapat memindahkan gadget.
  8. ⚠ Pemindahan dapat dilakukan dengan keyboard: fokus pada handle, Space untuk mengangkat, panah untuk memindahkan, Space untuk menaruh, Esc untuk membatalkan, dengan pengumuman posisi lewat ARIA live region.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Endpoint `PUT /api/v1/dashboards/{id}/gadgets/positions` menyimpan banyak posisi dalam satu transaksi. Standar aksesibilitas mengikuti modul 27.

## Kerangka gadget bersama

### DSH-021 Gadget configuration (common)

- **Deskripsi:** Form konfigurasi gadget yang menentukan sumber data, tampilan, dan perilaku refresh. Parameter umum ini dipakai bersama oleh banyak gadget sehingga dibangun sebagai komponen form yang dapat digunakan ulang.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Owner, Editor dashboard, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Picker "Project or saved filter" menyediakan pencarian nama dengan hasil dipisah ke kelompok Projects dan Filters; hanya project yang dapat di-browse dan filter yang dapat dilihat editor yang ditawarkan.
  2. ⚠ Tautan "Advanced search" pada picker membuka popup pencarian filter/project dengan kriteria tambahan (owner, project, favorit); input JQL langsung tidak disediakan di gadget, sehingga JQL harus disimpan sebagai saved filter terlebih dahulu.
  3. Field Statistic type menawarkan Assignee, Reporter, Creator, Priority, Status, Issue type, Project, Component, Fix version, Affects version, Labels, Resolution, Sprint, serta custom field yang memiliki statistics searcher (select, multi-select, user picker, version, dan sejenisnya).
  4. Field Period menawarkan Hourly, Daily, Weekly, Monthly, Quarterly, dan Yearly.
  5. ⚠ Field Days previously menerima bilangan bulat positif dengan batas per period: Hourly maksimal 10 hari, Daily maksimal 300 hari, Weekly maksimal 1.750 hari, Monthly maksimal 7.500 hari, Quarterly maksimal 22.500 hari, Yearly maksimal 36.500 hari; nilai di luar batas ditolak dengan pesan yang menyebutkan batasnya.
  6. Field Number of results menerima bilangan bulat 1-50 (default 10).
  7. Field Columns to display berupa field picker berurutan untuk gadget berbasis daftar work item, memakai katalog kolom navigator (SRC-025).
  8. Checkbox "Refresh otomatis setiap 15 menit" tersedia di setiap gadget yang mendukung auto-refresh (DSH-023).
  9. Field wajib yang kosong atau nilai tidak valid ditandai inline dengan pesan galat di bawah field dan tombol Save tidak mengirim permintaan.
  10. Form memiliki tombol Save (menyimpan dan merender gadget) dan Cancel (membatalkan perubahan; pada gadget yang belum pernah dikonfigurasi, Cancel menghapus gadget setelah konfirmasi).
  11. Gadget yang belum dikonfigurasi menampilkan form secara otomatis saat pertama ditambahkan; viewer yang melihat gadget belum dikonfigurasi melihat pesan "Gadget ini belum dikonfigurasi".
  12. Konfigurasi disimpan per instance gadget sebagai dashboard item properties berformat JSON; dua instance gadget yang sama tidak berbagi konfigurasi.
  13. Hanya pengguna dengan hak edit yang dapat mengubah konfigurasi; API menolak permintaan lain dengan 403.
  14. Hanya filter yang dapat diakses editor saat mengonfigurasi yang dapat dipilih; ID filter tanpa akses ditolak dengan 400.
- **Team-managed vs company-managed:** Tidak ada perbedaan; custom field team-managed yang memiliki searcher statistik ikut ditawarkan.
- **Catatan & ketergantungan:** Daftar statistic type diturunkan dari metadata searcher custom field (modul 05) dan dipakai bersama laporan Pie chart modul 16. Popup Advanced search memakai pencarian filter modul 13.

### DSH-022 Gadget actions menu

- **Deskripsi:** Menu aksi pada header gadget untuk mengonfigurasi, menyegarkan, meminimalkan, memaksimalkan, dan menghapus gadget.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Owner/Editor (semua aksi), viewer (Refresh, Minimize, Maximize).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Configure membuka form konfigurasi gadget (DSH-021) dan hanya tersedia bagi editor.
  2. Aksi Refresh me-reload data gadget ini saja.
  3. Aksi Minimize menciutkan isi gadget sehingga hanya header yang tampil; Expand mengembalikannya; status minimize disimpan per pengguna per gadget.
  4. Aksi Maximize membuka gadget dalam tampilan penuh satu gadget (canvas view) di URL `/dashboards/{id}/gadgets/{gadgetId}`, dengan tombol Restore untuk kembali ke dashboard.
  5. Aksi Delete menghapus gadget setelah konfirmasi "Hapus gadget <judul>?" dan hanya tersedia bagi editor.
  6. Pada gadget berbasis filter/project, judul gadget berupa tautan ke Issue Navigator dengan JQL sumber gadget.
  7. Isi menu menyesuaikan izin: viewer hanya melihat Refresh, Minimize/Expand, dan Maximize.
  8. ⚠ Judul gadget bawaan tidak dapat di-rename dari UI kecuali gadget yang memiliki field Title di konfigurasinya (mis. Activity Stream); judul dan warna dapat diatur lewat REST API (DSH-057) dan judul hasil REST ditampilkan di UI.
  9. Pemilihan warna chrome gadget tidak disediakan di UI; warna hanya dapat diatur lewat REST API dan dirender sebagai aksen garis atas header gadget.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status minimize per pengguna disimpan di tabel preferensi pengguna (modul 21), bukan di dashboard.

### DSH-023 Gadget refresh & auto-refresh

- **Deskripsi:** Penyegaran data gadget secara manual atau otomatis berkala agar dashboard tetap terbaru tanpa reload halaman.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengatur auto-refresh), semua viewer (refresh manual).
- **Aturan perilaku (acceptance criteria):**
  1. Checkbox auto-refresh setiap 15 menit di konfigurasi gadget; bila aktif, gadget memuat ulang datanya setiap 15 menit selama halaman terbuka.
  2. Pilihan interval legacy DC (Never, 15 menit, 30 menit, 1 jam, 2 jam) tidak ditampilkan di UI; nilai interval hasil impor dari DC dipetakan ke "aktif 15 menit" bila bukan Never.
  3. Refresh manual per gadget tersedia di menu gadget (DSH-022).
  4. Tombol Refresh di header dashboard memuat ulang semua gadget.
  5. Refresh satu gadget tidak memuat ulang atau mengganggu interaksi gadget lain.
  6. Selama refresh, gadget menampilkan indikator loading tanpa mengosongkan data lama, dan menampilkan waktu "Diperbarui <waktu relatif>" setelah selesai.
  7. ⚠ Auto-refresh dijeda saat tab browser tidak aktif (Page Visibility API) dan langsung dijalankan sekali saat tab kembali aktif bila jadwalnya sudah terlewati.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hasil query gadget di-cache singkat per viewer (mis. 60 detik) untuk meredam banyak refresh bersamaan (modul 27).

### DSH-024 Gadget permission-aware rendering

- **Deskripsi:** Isi gadget selalu dihitung dengan izin pengguna yang melihat, sehingga dashboard yang dibagikan tidak membocorkan data.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua viewer termasuk anonim; owner dashboard tidak dapat "meminjamkan" izinnya ke viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Query gadget dieksekusi dengan Browse Projects dan issue security level milik viewer (PERM-020, PERM-044), bukan milik owner atau editor dashboard.
  2. Bila filter sumber gadget tidak dibagikan ke viewer, gadget menampilkan galat "Filter yang dipakai gadget ini tidak tersedia untuk Anda" tanpa membuka nama atau JQL filter.
  3. Bila filter atau project sumber dihapus, gadget menampilkan galat konfigurasi "Sumber data gadget tidak ditemukan"; editor melihat tombol "Konfigurasi ulang".
  4. Bila board gadget agile tidak dapat diakses viewer, gadget menampilkan pesan "Anda tidak memiliki akses ke board ini".
  5. Pengguna anonim hanya melihat data project yang Browse Projects-nya diberikan ke Public/Anyone.
  6. Field yang disembunyikan oleh field configuration pada project terkait tidak ditampilkan di kolom gadget daftar work item (sel kosong).
  7. Jumlah, persentase, dan statistik dihitung hanya dari work item yang terlihat oleh viewer; total tidak pernah memasukkan work item tersembunyi.
  8. Gadget berbasis `currentUser()` (Assigned to Me, Watched, Voted, Issues in Progress) selalu dihitung untuk viewer yang sedang melihat.
  9. Gadget dari app yang sudah di-uninstall menampilkan pesan "Gadget ini tidak lagi tersedia" dan tetap dapat dihapus oleh editor.
- **Team-managed vs company-managed:** Tidak ada perbedaan; akses project team-managed mengikuti access level project (PERM-038).
- **Catatan & ketergantungan:** Semua endpoint data gadget wajib melewati satu lapisan eksekusi query (modul 12/13) yang menyuntikkan kondisi izin; pengujian integrasi wajib mencakup viewer dengan izin lebih sempit dari owner.

### DSH-025 Chart gadget drill-down & rendering

- **Deskripsi:** Perilaku umum gadget chart/statistik: tooltip, legend, dan klik data untuk membuka Issue Navigator dengan JQL yang sesuai.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Hover atau fokus keyboard pada bar/slice/titik menampilkan tooltip berisi label dan nilai.
  2. Legend menampilkan nilai, jumlah, dan persentase setiap seri/kategori.
  3. Klik elemen chart membuka Issue Navigator dengan JQL gabungan sumber gadget dan nilai/periode yang diklik (mis. `filter = 10010 AND priority = High` atau `project = ABC AND created >= "2026-10-01" AND created < "2026-10-08"`).
  4. ⚠ Gadget Pie Chart, Created vs Resolved, dan Recently Created menyediakan tabel data pendukung yang dapat dibuka lewat toggle "Tampilkan data" di bawah chart.
  5. Chart di-resize mengikuti lebar kolom dan ukuran jendela.
  6. Bila sumber tidak menghasilkan work item, gadget menampilkan pesan "Tidak ada data" alih-alih chart kosong.
  7. ⚠ Chart dirender di sisi klien (SVG) dari data JSON yang dikembalikan backend; rendering gambar di server seperti DC tidak dibangun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai komponen chart bersama modul 16 (RPT-008) termasuk palet yang aman untuk buta warna dan dukungan mode gelap.

### DSH-026 Agile gadgets common configuration

- **Deskripsi:** Konfigurasi bersama untuk gadget berbasis board/sprint (Sprint Burndown, Sprint Health, Days Remaining in Sprint).
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor dashboard (mengonfigurasi), viewer (melihat bila dapat mengakses board).
- **Aturan perilaku (acceptance criteria):**
  1. Picker Board hanya menawarkan board Scrum yang dapat diakses editor.
  2. ⚠ Picker Sprint menawarkan opsi "Sprint aktif (otomatis)" sebagai default, yang selalu memilih sprint aktif board saat gadget dirender, atau sprint aktif tertentu.
  3. Bila board menjalankan parallel sprints, opsi otomatis memilih sprint aktif dengan tanggal mulai paling awal dan editor dapat memilih salah satu sprint aktif secara eksplisit.
  4. Opsi "Tampilkan nama board" dan "Tampilkan nama sprint" mengatur teks di header gadget.
  5. Pesan khusus ditampilkan bila board berubah menjadi Kanban ("Gadget ini hanya untuk board Scrum"), tidak ada sprint aktif ("Tidak ada sprint aktif"), atau board dihapus ("Board tidak ditemukan").
  6. Perhitungan hari memakai working days dan non-working days dari konfigurasi board (BRD-040).
  7. ⚠ Berlaku untuk board company-managed dan board team-managed dengan fitur Sprints aktif; board team-managed tampil di picker dengan nama project.
- **Team-managed vs company-managed:** Board team-managed hanya tersedia bila fitur Sprints aktif; working days team-managed mengikuti pengaturan project.
- **Catatan & ketergantungan:** Data sprint dari modul 11; estimation statistic dari TIME-022.

## Gadget daftar work item

### DSH-027 Assigned to Me gadget

- **Deskripsi:** Menampilkan daftar work item belum selesai yang di-assign ke pengguna yang sedang melihat.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna login; editor mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Query implisit `assignee = currentUser() AND resolution = Unresolved ORDER BY priority DESC, updated DESC`; tidak dapat diubah lewat konfigurasi.
  2. Konfigurasi Number of results (1-50, default 10).
  3. Konfigurasi Columns to display (default Issue Type, Key, Summary, Priority).
  4. Klik header kolom mengurutkan hasil ascending/descending pada kolom yang dapat diurutkan.
  5. Hasil dipaginasi sesuai Number of results dengan teks "1-10 dari N".
  6. Tautan "Lihat di pencarian" membuka Issue Navigator dengan JQL gadget.
  7. Auto-refresh opsional (DSH-023).
  8. Data selalu untuk viewer meskipun dashboard dibagikan (DSH-024).
  9. Bila tidak ada work item yang di-assign, gadget menampilkan empty state "Tidak ada work item yang di-assign ke Anda".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen tabel daftar work item bersama dipakai DSH-027 s.d. DSH-031. Halaman "Your work" (modul 21) memakai query serupa.

### DSH-028 Issues in Progress gadget

- **Deskripsi:** Menampilkan work item yang di-assign ke pengguna yang melihat dan sedang berada dalam status In Progress.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Query implisit `assignee = currentUser() AND statusCategory = "In Progress"`, sehingga semua status berkategori In Progress (bukan hanya status bernama "In Progress") ikut dihitung.
  2. Konfigurasi Number of results (1-50, default 10).
  3. Konfigurasi Columns to display.
  4. Sorting lewat klik header kolom.
  5. Tautan "Lihat di pencarian" ke Issue Navigator dengan JQL gadget.
  6. Auto-refresh opsional.
  7. Data selalu dihitung per viewer.
- **Team-managed vs company-managed:** Tidak ada perbedaan; status team-managed dipetakan ke status category yang sama.
- **Catatan & ketergantungan:** Status category didefinisikan di modul 09.

### DSH-029 Watched Issues gadget

- **Deskripsi:** Daftar work item yang di-watch oleh pengguna yang sedang melihat.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Query implisit `issue in watchedIssues()` (JQL-041).
  2. ⚠ Opsi "Tampilkan jumlah watcher" menambahkan kolom Watchers berisi jumlah watcher setiap work item.
  3. Opsi "Tampilkan work item resolved"; bila tidak dicentang (default), query ditambah `resolution = Unresolved`.
  4. Konfigurasi Number of results dan Columns to display.
  5. Sorting lewat header kolom dan pagination hasil.
  6. Bila fitur watching dinonaktifkan admin (CFG-057), gadget menampilkan pesan "Fitur watching tidak aktif di situs ini".
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Watchers di ISS-036, izin View Voters and Watchers (PERM-033) menentukan apakah jumlah watcher boleh ditampilkan.

### DSH-030 Voted Issues gadget

- **Deskripsi:** Daftar work item yang di-vote oleh pengguna yang sedang melihat.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Query implisit `issue in votedIssues()` (JQL-041).
  2. Opsi "Tampilkan total vote" menambahkan kolom jumlah vote.
  3. Opsi "Tampilkan work item resolved" (default tidak dicentang).
  4. Konfigurasi Number of results dan Columns to display.
  5. Sorting lewat header kolom dan pagination hasil.
  6. Bila voting dinonaktifkan admin (CFG-057), gadget menampilkan pesan "Fitur voting tidak aktif di situs ini".
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Votes di ISS-037.

### DSH-031 Filter Results gadget

- **Deskripsi:** Menampilkan daftar work item hasil saved filter dalam bentuk tabel dengan kolom yang dapat dikonfigurasi.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan akses ke filter (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Saved filter wajib dipilih; project tidak dapat dipakai sebagai sumber.
  2. Number of results 1-50 per halaman (default 10).
  3. Columns to display dapat dikonfigurasi, termasuk opsi "Gunakan kolom default filter" yang mengikuti konfigurasi kolom saved filter.
  4. Klik header kolom mengubah urutan; urutan default mengikuti `ORDER BY` filter.
  5. Pagination menampilkan total jumlah work item hasil filter.
  6. Judul gadget menampilkan nama filter dan menaut ke Issue Navigator dengan filter tersebut.
  7. Field dirender sesuai tipenya: avatar assignee, ikon priority dan work type, lozenge status berwarna status category.
  8. Klik key work item membuka work item tersebut.
  9. Auto-refresh opsional.
  10. Bila filter dihapus atau tidak dibagikan ke viewer, gadget menampilkan galat sesuai DSH-024.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi kolom memakai katalog kolom navigator (SRC-025) dan paginasi navigator (SRC-027).

### DSH-032 Favourite Filters gadget

- **Deskripsi:** Daftar filter berbintang milik pengguna yang melihat beserta jumlah work item masing-masing.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Gadget menampilkan semua starred filter milik viewer yang masih dapat ia lihat, diurutkan alfabetis.
  2. Opsi "Tampilkan jumlah work item" menampilkan jumlah hasil setiap filter, dihitung dengan izin viewer.
  3. Klik nama filter membuka Issue Navigator dengan filter tersebut.
  4. Tautan "Kelola filter" membuka halaman Filters (modul 13).
  5. Data selalu per viewer; dashboard yang dibagikan menampilkan starred filter masing-masing viewer.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bintang filter dimiliki modul 13. Penghitungan jumlah dibatasi timeout per filter (DSH-059); filter yang timeout menampilkan "?".

## Gadget statistik & distribusi

### DSH-033 Issue Statistics gadget

- **Deskripsi:** Tabel statistik jumlah work item dari project/filter yang dikelompokkan berdasarkan satu field, dengan bar persentase.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer (melihat sesuai izin).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Statistic type (wajib, DSH-021).
  3. Sort by Natural (urutan alami nilai field, mis. urutan priority atau urutan version) atau Total (jumlah work item).
  4. Sort direction Ascending atau Descending.
  5. ⚠ Opsi "Sertakan work item resolved" (default tidak dicentang); bila tidak dicentang, statistik hanya menghitung work item `resolution = Unresolved`.
  6. Setiap baris berisi nilai field, jumlah work item, bar persentase, dan persentase dengan satu desimal.
  7. Baris "None" menghitung work item tanpa nilai pada field tersebut.
  8. Baris Total di bagian bawah menampilkan total work item; untuk field multi-nilai (Labels, Component), catatan menjelaskan bahwa jumlah baris dapat melebihi total.
  9. Klik nilai atau jumlah membuka Issue Navigator dengan JQL sumber ditambah klausa nilai (atau `IS EMPTY` untuk None).
  10. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Agregasi memakai query `GROUP BY` di atas lapisan eksekusi query berizin (DSH-024).

### DSH-034 Two Dimensional Filter Statistics gadget

- **Deskripsi:** Tabel matriks jumlah work item dari saved filter dengan dua field statistik sebagai sumbu X dan Y.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan akses ke filter.
- **Aturan perilaku (acceptance criteria):**
  1. Saved filter wajib dipilih.
  2. XAxis dan YAxis masing-masing memilih statistic type; keduanya wajib dan boleh sama.
  3. Sort by Natural atau Total untuk baris sumbu Y.
  4. Sort direction Ascending atau Descending.
  5. Opsi "Tampilkan total" menambahkan baris dan kolom total.
  6. Number of results menentukan jumlah baris sumbu Y yang ditampilkan (default 5).
  7. Setiap sel berisi jumlah yang dapat diklik ke Issue Navigator dengan JQL filter ditambah klausa nilai X dan Y.
  8. Baris dan kolom "None" menampung work item tanpa nilai pada field sumbu.
  9. Tautan "Tampilkan lebih banyak" muncul bila jumlah baris melebihi limit dan menampilkan seluruh baris di dalam gadget.
  10. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sel bernilai nol dirender kosong agar matriks mudah dibaca; kolom sumbu X yang melebihi lebar gadget dapat di-scroll horizontal di dalam gadget.

### DSH-035 Pie Chart gadget

- **Deskripsi:** Pie chart distribusi work item dari project/filter berdasarkan satu field statistik.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer (melihat sesuai izin).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Statistic type (wajib).
  3. Legend menampilkan nilai, jumlah, dan persentase setiap slice.
  4. Hover atau fokus slice menampilkan jumlah dan persentase.
  5. Klik slice membuka Issue Navigator dengan JQL sumber ditambah klausa nilai slice.
  6. ⚠ Bila terdapat lebih dari 10 nilai, nilai dengan porsi terkecil digabung menjadi slice "Other"; klik slice Other membuka pencarian dengan klausa `NOT IN` nilai yang ditampilkan.
  7. Total work item ditampilkan di bawah chart.
  8. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sama dengan laporan Pie chart modul 16; satu endpoint agregasi melayani keduanya.

### DSH-036 Heat Map gadget

- **Deskripsi:** Menampilkan distribusi nilai sebuah field pada work item project/filter dengan ukuran/intensitas sesuai proporsi jumlah work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Statistic type (wajib).
  3. Setiap nilai ditampilkan sebagai label yang ukuran huruf/intensitas warnanya sebanding dengan jumlah work item nilai tersebut, dengan minimal 5 tingkat skala.
  4. Hover menampilkan tooltip jumlah dan persentase.
  5. Klik nilai membuka Issue Navigator dengan JQL terkait.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai endpoint agregasi yang sama dengan DSH-033.

### DSH-037 Labels gadget

- **Deskripsi:** Menampilkan semua label yang dipakai di sebuah project dalam daftar alfabetis yang dapat diklik.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan Browse Projects pada project.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project wajib (saved filter tidak didukung).
  2. Konfigurasi Label field: field Labels sistem atau custom field bertipe Labels (CFG-031).
  3. Label dikelompokkan per huruf awal (A, B, ...; angka dan simbol di kelompok "#").
  4. Klik label membuka Issue Navigator dengan JQL `project = X AND labels = "<label>"` (atau field custom terkait).
  5. Hanya label dari work item yang terlihat oleh viewer yang ditampilkan.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan; project team-managed memakai field Labels yang sama.
- **Catatan & ketergantungan:** Labels di ISS-020.

### DSH-038 Workload Pie Chart gadget

- **Deskripsi:** Pie chart distribusi beban kerja berbasis time tracking (estimate atau time spent) per assignee atau field lain.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Statistic type (default Assignee).
  3. Konfigurasi Time field: Original estimate, Remaining estimate, atau Time spent.
  4. Nilai ditampilkan dalam format durasi sesuai pengaturan time tracking (mis. "3w 2d 4h" atau jam desimal).
  5. Legend menampilkan durasi dan persentase; klik slice membuka Issue Navigator dengan JQL terkait.
  6. Bila time tracking tidak aktif, gadget tidak tampil di panel Add gadget dan instance yang ada menampilkan pesan "Time tracking tidak aktif".
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Padanan laporan Workload pie chart (TIME-034) dan gadget berbasis waktu (TIME-035) di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md).

## Gadget tren & waktu

### DSH-039 Created vs Resolved Chart gadget

- **Deskripsi:** Chart perbandingan jumlah work item yang dibuat dan yang diselesaikan per periode, untuk melihat apakah backlog bertambah atau berkurang.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Period dan Days previously (default Daily, 30 hari).
  3. Opsi Cumulative totals (Yes/No); Yes menampilkan jumlah kumulatif sejak awal rentang.
  4. Opsi "Tampilkan tren unresolved" menambahkan subchart jumlah work item unresolved per periode.
  5. Opsi Display versions: All versions, Only major versions, atau None; version yang rilis dalam rentang ditandai marker vertikal berlabel nama version.
  6. Area antara dua garis berwarna merah bila created > resolved dan hijau bila resolved > created, dengan pola tambahan agar tidak bergantung pada warna saja.
  7. Tooltip per titik menampilkan jumlah created dan resolved; klik titik membuka Issue Navigator dengan JQL periode tersebut (created atau resolved sesuai seri).
  8. Ringkasan di bawah chart menampilkan total work item created dan resolved pada rentang.
  9. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan; marker version hanya bila project memakai Releases.
- **Catatan & ketergantungan:** Kalkulasi dan definisi "major version" mengikuti laporan Created vs resolved modul 16.

### DSH-040 Recently Created Chart gadget

- **Deskripsi:** Bar chart jumlah work item yang dibuat per periode, dibedakan antara yang sudah resolved dan belum.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Period dan Days previously.
  3. Setiap bar ditumpuk: bagian resolved (hijau) dan unresolved (merah), berdasarkan status resolusi saat ini.
  4. Tooltip menampilkan jumlah setiap bagian bar.
  5. Klik bar (atau bagiannya) membuka Issue Navigator dengan JQL periode dan status resolusi terkait.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sama dengan laporan Recently created modul 16.

### DSH-041 Resolution Time gadget

- **Deskripsi:** Bar chart rata-rata waktu penyelesaian work item (dari created hingga resolved) per periode.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Period dan Days previously.
  3. Nilai ditampilkan dalam hari dengan satu desimal.
  4. Hanya work item yang resolved di dalam periode bar tersebut yang dihitung.
  5. Tooltip menampilkan rata-rata dan jumlah work item; klik bar membuka Issue Navigator dengan JQL `resolved` pada periode tersebut.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sama dengan laporan Resolution time modul 16.

### DSH-042 Average Age Chart gadget

- **Deskripsi:** Bar chart rata-rata umur (hari) work item yang belum terselesaikan pada setiap periode untuk project/filter.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Period (Hourly sampai Yearly).
  3. Konfigurasi Days previously dengan batas per period (DSH-021).
  4. Auto-refresh opsional.
  5. Tooltip per bar menampilkan rata-rata umur dan jumlah work item unresolved pada akhir periode.
  6. Klik bar membuka Issue Navigator dengan JQL work item yang unresolved pada periode itu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Umur pada periode dihitung dari changelog resolusi (RPT-004), bukan hanya status saat ini.

### DSH-043 Time Since Chart gadget

- **Deskripsi:** Bar chart jumlah work item yang field tanggalnya (mis. Created, Resolved) jatuh pada setiap periode.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Date field: Created, Updated, Resolved, Due date, atau custom field date/datetime.
  3. Konfigurasi Period dan Days previously.
  4. Opsi Cumulative totals.
  5. Tooltip per bar dan klik bar membuka Issue Navigator dengan JQL field tanggal pada periode tersebut.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sama dengan laporan Time since modul 16.

### DSH-044 Average Time in Status gadget

- **Deskripsi:** Chart rata-rata lama waktu work item berada di setiap status untuk work item yang diselesaikan dalam periode.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Pemilihan status yang ditampilkan (multi-select, minimal satu) dari status yang dipakai workflow project/filter.
  3. Konfigurasi Period dan Days previously.
  4. Setiap status menjadi satu seri dengan warna dan entri legend sendiri.
  5. Satuan waktu dipilih otomatis: jam bila rata-rata maksimal < 48 jam, selain itu hari.
  6. Tooltip menampilkan rata-rata waktu per status pada periode tersebut.
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan; status team-managed per project ditampilkan dengan nama project bila namanya bentrok.
- **Catatan & ketergantungan:** Durasi dihitung dari status history (WF-014) dan changelog (CMT-023).

### DSH-045 Average Number of Times in Status gadget

- **Deskripsi:** Chart rata-rata berapa kali work item masuk ke status tertentu dalam periode waktu.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Pemilihan satu atau beberapa status.
  3. Konfigurasi Period dan Days previously.
  4. Legend dan tooltip per status menampilkan rata-rata jumlah masuk status.
  5. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi populasi (hanya work item resolved di periode, atau semua work item yang masuk status di periode) belum dipastikan terhadap produk acuan; produk ini memakai work item yang resolved di periode agar konsisten dengan DSH-044, dicatat sebagai pertanyaan terbuka.

### DSH-046 Bubble Chart gadget

- **Deskripsi:** Visualisasi work item sebagai gelembung untuk menyoroti work item yang paling ramai berdasarkan partisipan atau vote serta aktivitas komentar.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Bubble size: Participants (jumlah pengguna unik yang berkomentar, reporter, dan assignee) atau Votes.
  3. ⚠ Sumbu X adalah umur work item (hari sejak dibuat) dan sumbu Y adalah jumlah komentar.
  4. Warna bubble menunjukkan kebaruan komentar terakhir berdasarkan konfigurasi "recent comment period" dalam hari (default 7): bubble dengan komentar dalam periode itu berwarna lebih pekat.
  5. Hover bubble menampilkan key dan summary work item; klik bubble membuka work item.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Jumlah work item yang dirender dibatasi (mis. 200 teratas) agar chart tetap terbaca; batas mengikuti DSH-059.

## Gadget sprint, rilis & kalender

### DSH-047 Sprint Burndown gadget

- **Deskripsi:** Chart burndown sprint aktif dari sebuah board Scrum, setara dengan laporan Burndown chart.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan akses ke board.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Board (Scrum) dan Sprint mengikuti DSH-026.
  2. Memakai estimation statistic board (Story points, Original time estimate, atau Issue count) (TIME-022).
  3. Menampilkan garis Guideline dari total komitmen awal ke nol dan garis Remaining values yang turun saat work item selesai.
  4. Non-working days ditampilkan sebagai area berarsir dan guideline datar pada hari tersebut.
  5. Penanda scope change ditampilkan saat work item ditambahkan ke atau dihapus dari sprint setelah sprint dimulai, beserta tooltip work item terkait.
  6. Opsi tampilkan nama sprint dan nama board.
  7. Tautan "Lihat laporan lengkap" membuka laporan Burndown chart (RPT-009) untuk board dan sprint yang sama.
  8. Pesan ditampilkan bila board Kanban atau tidak ada sprint aktif (DSH-026).
  9. Auto-refresh opsional.
- **Team-managed vs company-managed:** Mengikuti DSH-026.
- **Catatan & ketergantungan:** Kalkulasi dimiliki RPT-004/RPT-009; gadget hanya menampilkan versi ringkas tanpa tabel event.

### DSH-048 Sprint Health gadget

- **Deskripsi:** Ringkasan kesehatan sprint aktif: waktu berjalan, progres pekerjaan, perubahan scope, dan blocker.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan akses ke board.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Board dan Sprint mengikuti DSH-026.
  2. Progress bar tersegmentasi To Do / In Progress / Done berdasarkan estimation statistic board dan status category.
  3. Persentase Time elapsed = hari kerja berlalu / total hari kerja sprint.
  4. Persentase Work complete = estimasi Done / total estimasi sprint saat ini.
  5. Persentase Scope change = perubahan total estimasi sejak sprint dimulai / komitmen awal, bertanda + atau -.
  6. Jumlah Blockers = work item sprint yang sedang di-flag; klik angka membuka pencarian work item flagged di sprint.
  7. Opsi tampilkan nama sprint dan nama board.
  8. Pesan "Tidak ada sprint aktif" bila board tidak memiliki sprint aktif.
  9. Auto-refresh opsional.
- **Team-managed vs company-managed:** Mengikuti DSH-026.
- **Catatan & ketergantungan:** Flag di modul 06/10; komitmen awal dari snapshot sprint start (SPR-018).

### DSH-049 Days Remaining in Sprint gadget

- **Deskripsi:** Menampilkan sisa hari kerja pada sprint aktif sebuah board Scrum.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan akses ke board.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Board (Scrum) dan Sprint mengikuti DSH-026.
  2. Tampilan berupa angka besar jumlah hari kerja tersisa beserta tanggal berakhir sprint; sprint yang melewati tanggal berakhir menampilkan "Lewat N hari".
  3. Perhitungan mengecualikan non-working days board (BRD-040).
  4. Opsi tampilkan nama board dan nama sprint.
  5. Pesan "Tidak ada sprint aktif" bila board tidak memiliki sprint aktif.
  6. Auto-refresh opsional.
- **Team-managed vs company-managed:** Mengikuti DSH-026.
- **Catatan & ketergantungan:** Zona waktu perhitungan mengikuti zona waktu board/situs (modul 21/22).

### DSH-050 Road Map gadget

- **Deskripsi:** Menampilkan version (release) yang dijadwalkan rilis dalam N hari ke depan beserta progres work item-nya.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer dengan Browse Projects pada project terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Projects (satu atau beberapa) atau project category.
  2. Konfigurasi Days ke depan (1-1.000, default 30).
  3. Konfigurasi jumlah maksimum version per project (default 10).
  4. Daftar version unreleased dengan release date dalam jangka waktu, menampilkan release date dan progress bar resolved vs unresolved (REL-013).
  5. Klik version membuka halaman detail version (REL-014).
  6. Version yang release date-nya sudah lewat tetapi belum dirilis ditandai lencana "Overdue".
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Project team-managed hanya muncul bila fitur Releases aktif.
- **Catatan & ketergantungan:** Bukan Timeline atau Plans (modul 18); gadget ini berbasis version.

### DSH-051 Jira Issues Calendar gadget

- **Deskripsi:** Kalender bulanan yang menampilkan work item berdasarkan field tanggal dan tanggal rilis version dari project/filter.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Project or saved filter (wajib).
  2. Konfigurasi Date field: Due date (default), Created, Updated, Resolved, atau custom field date/datetime.
  3. Opsi "Tampilkan version" menambahkan release date version pada kalender.
  4. ⚠ Tampilan Month view dengan navigasi bulan sebelumnya/berikutnya dan tombol "Hari ini"; view mingguan tidak dibangun pada rilis awal.
  5. Klik work item membuka detail work item.
  6. Bila satu tanggal memiliki lebih banyak work item daripada ruang yang tersedia (lebih dari 3), tampil tautan "+N lainnya" yang membuka daftar lengkap tanggal itu.
  7. ⚠ Langganan iCal tidak dibangun pada rilis awal; kebutuhan ekspor kalender dicatat sebagai pertanyaan terbuka.
  8. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama tampilan di produk ini "Work Items Calendar"; berbeda dari Calendar view project (modul 18).

## Gadget aktivitas & informasi

### DSH-052 Activity Stream gadget

- **Deskripsi:** Gadget yang menampilkan aliran aktivitas terbaru (work item dibuat, diubah, dikomentari, ditransisikan) dari produk ini dan aplikasi yang terhubung.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer (melihat aktivitas yang boleh ia lihat).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Title gadget (default "Activity Stream", maksimal 255 karakter).
  2. Global filter yang dapat dikombinasikan: Project (is/is not), Username (is/is not), Update date (before/after/within last N days), dan Issue key (is/is not).
  3. Filter per stream source: produk ini (jenis aktivitas comment, create, transition, attachment, update field, link, worklog, sprint) dan aplikasi terhubung lewat application link bila ada (modul 25).
  4. Limit jumlah item (1-100, default 10).
  5. Opsi auto-refresh.
  6. Item dikelompokkan per tanggal dengan label "Hari ini", "Kemarin", lalu tanggal lengkap.
  7. Setiap item menampilkan avatar, aktor, aksi, key dan summary work item, serta waktu relatif dengan tooltip waktu absolut.
  8. Tautan "Tampilkan lebih banyak" memuat item berikutnya sebanyak limit.
  9. Aktivitas dari work item yang tidak dapat dilihat viewer tidak pernah muncul, termasuk komentar dengan restriction visibility yang tidak mencakup viewer.
  10. ⚠ Aksi Comment inline dan feed RSS/Atom tidak dibangun pada rilis awal; item menaut ke work item untuk berkomentar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sumber data adalah changelog dan event domain (CMT-023, modul 28); query feed dibatasi rentang waktu agar efisien.

### DSH-053 Introduction gadget

- **Deskripsi:** Gadget yang menampilkan teks pengantar yang dikonfigurasi Admin Situs, biasanya di system dashboard.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengisi teks), semua pengguna (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Konten diambil dari pengaturan "Introduction" di General configuration (modul 22) dan sama untuk semua instance gadget.
  2. Bila belum dikustomisasi, gadget menampilkan teks sambutan default beserta tautan panduan memulai.
  3. Tidak ada konfigurasi per gadget.
  4. Gadget juga tampil untuk pengguna anonim pada system dashboard bila akses anonim diizinkan.
  5. ⚠ Teks Introduction ditulis dengan editor teks kaya (modul 07) dan dirender sebagai konten tersanitasi; HTML mentah dan script tidak didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Instance baru memuat gadget ini di system dashboard (DSH-009). Ditempatkan di Fase 3 walau P2 karena seed system dashboard (DSH-009, P0 Fase 3) memerlukan gadget ini; implementasinya kecil (teks admin tanpa konfigurasi per gadget).

### DSH-054 Projects gadget

- **Deskripsi:** Menampilkan ringkasan project terpilih beserta tautan cepat ke daftar work item project.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Editor (mengonfigurasi), viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi Projects: beberapa project, All projects, atau project category.
  2. ⚠ Konfigurasi View: Collapsed (nama saja), Brief (nama, key, lead, tautan dasar), atau Full (ditambah tautan per work type, version, dan component).
  3. Konfigurasi Number of columns (1-3).
  4. Setiap project menampilkan avatar, nama, key, dan project lead.
  5. Tautan filter cepat: All, Unresolved, Added recently, Resolved recently, serta per work type/version/component pada view Full; setiap tautan membuka Issue Navigator dengan JQL terkait.
  6. Hanya project yang dapat di-browse viewer yang ditampilkan.
  7. Auto-refresh opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan; tautan per version hanya untuk project yang memakai Releases.
- **Catatan & ketergantungan:** Project category di modul 04.

### DSH-055 Quick Links gadget

- **Deskripsi:** Kumpulan tautan cepat ke halaman yang sering dipakai seperti work item yang dilaporkan, di-vote, di-watch, dan pencarian.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Tautan My Reported Issues (`reporter = currentUser()`), Voted Issues, dan Watched Issues ke Issue Navigator.
  2. Tautan Create work item dan Search work items.
  3. Tautan Browse projects.
  4. Tautan administrasi (Settings) hanya tampil untuk Admin Situs.
  5. Gadget tidak memerlukan konfigurasi.
  6. Tautan disesuaikan dengan izin viewer: tautan Create tidak tampil bila viewer tidak memiliki Create Issues di project mana pun, dan tautan Voted/Watched tidak tampil bila fiturnya dinonaktifkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tautan memakai fungsi JQL modul 12.

## API, ekstensi & operasional

### DSH-056 Dashboard REST API

- **Deskripsi:** Endpoint REST untuk membuat, membaca, mencari, menyalin, mengubah, dan menghapus dashboard serta mengelola izin secara massal, dengan kontrak yang kompatibel dengan REST API v2/v3 produk acuan.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API (dengan token pengguna), Admin Situs; setiap endpoint menerapkan aturan izin yang sama dengan UI.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /rest/api/3/dashboard` mengembalikan daftar dashboard dengan parameter `filter=my|favourite`, `startAt`, dan `maxResults` (default 20, maksimal 1.000).
  2. `GET /rest/api/3/dashboard/search` mendukung parameter `dashboardName`, `accountId`, `owner`, `groupname`, `groupId`, `projectId`, `orderBy` (`name`, `description`, `owner`, `id`, `favorite_count`, `is_favorite`, dengan prefiks `-` untuk descending), `status` (`active`, `archived`, `deleted`), dan `expand` (`description`, `owner`, `viewUrl`, `favourite`, `favouritedCount`, `sharePermissions`, `editPermissions`, `isWritable`).
  3. `GET /rest/api/3/dashboard/{id}` mengembalikan satu dashboard; dashboard tanpa akses mengembalikan 404.
  4. `POST /rest/api/3/dashboard` membuat dashboard dengan `name`, `description`, `sharePermissions`, dan `editPermissions`.
  5. `PUT /rest/api/3/dashboard/{id}` mengubah detail dan izin (penggantian penuh daftar permission).
  6. `DELETE /rest/api/3/dashboard/{id}` menghapus dashboard (204) dengan aturan DSH-008.
  7. `POST /rest/api/3/dashboard/{id}/copy` menyalin dashboard dengan body detail baru (DSH-007).
  8. `PUT /rest/api/3/dashboard/bulk/edit` menerima `entityIds` dan `action` (`changeOwner`, `changePermission`, `addPermission`, `removePermission`) dan mengembalikan daftar kegagalan per dashboard.
  9. Query parameter `extendAdminPermissions=true` memungkinkan Admin Situs mengelola dashboard milik orang lain; tanpa parameter itu admin diperlakukan seperti pengguna biasa.
  10. Tipe share permission yang diterima dan dikembalikan: `user`, `group`, `project`, `projectRole`, `global`, `loggedin`, `authenticated`, `project-unknown` (project yang tidak dapat dilihat pemanggil).
  11. Galat: 400 untuk nama kosong/duplikat atau permission tidak valid (dengan pesan per field), 401 tanpa autentikasi, 403 tanpa izin, 404 bila tidak ditemukan.
  12. Respons berisi `id`, `name`, `owner`, `isFavourite`, `popularity`, `rank`, `self`, `view` (URL), `sharePermissions`, dan `editPermissions`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** API internal SPA memakai `/api/v1/dashboards` (Swagger via `make swag`); jalur `/rest/api/3/...` adalah lapisan kompatibilitas yang dimiliki [23 API Publik](23-api-publik-webhook.md). Status `archived` tidak memiliki fitur UI dan dicatat di pertanyaan terbuka.

### DSH-057 Dashboard gadgets REST API

- **Deskripsi:** Endpoint REST untuk mengelola gadget pada dashboard dan properti item dashboard (konfigurasi gadget).
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API; mengubah gadget memerlukan hak edit dashboard, membaca memerlukan hak view.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /rest/api/3/dashboard/gadgets` mengembalikan katalog gadget yang tersedia (`moduleKey` atau `uri`, `title`).
  2. `GET /rest/api/3/dashboard/{dashboardId}/gadget` mengembalikan gadget pada dashboard, dapat difilter dengan `moduleKey`, `uri`, dan `gadgetId`.
  3. `POST /rest/api/3/dashboard/{dashboardId}/gadget` menambah gadget dengan `moduleKey` atau `uri`, `color`, `position` (`column`, `row`), `title`, dan `ignoreUriAndModuleKeyValidation`; mengirim `moduleKey` dan `uri` sekaligus ditolak 400.
  4. `PUT /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId}` mengubah `title`, `color`, dan `position`.
  5. `DELETE /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId}` menghapus gadget (204).
  6. `GET /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties` mengembalikan daftar property key.
  7. `GET`/`PUT`/`DELETE /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}` membaca, menyimpan, dan menghapus konfigurasi gadget berformat JSON.
  8. Warna gadget berupa enum `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, `white`; nilai lain ditolak 400.
  9. Posisi divalidasi: `column` harus berada dalam jumlah kolom layout (0-based) dan `row` ≥ 0; posisi yang sudah terisi menggeser gadget di bawahnya.
  10. ⚠ Nilai property dibatasi 32 KB per key; nilai lebih besar ditolak 413 dengan pesan batas ukuran.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format item properties dipakai juga oleh gadget app (DSH-058). Rate limit mengikuti modul 23/27.

### DSH-058 App-provided dashboard gadgets (Forge/Connect)

- **Deskripsi:** Mekanisme ekstensi agar app dapat menyediakan gadget dashboard sendiri dengan konfigurasi dan rendering kustom. Produk ini menyediakan titik ekstensi setara; app pihak ketiga itu sendiri di luar cakupan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (membuat app), Admin Situs (memasang app), editor dashboard (menambah gadget app).
- **Aturan perilaku (acceptance criteria):**
  1. Manifest app dapat mendeklarasikan modul `dashboardGadget` (padanan `jira:dashboardGadget`) dengan resource `view` dan `edit` yang dirender dari bundle app.
  2. ⚠ Manifest dapat mendeklarasikan modul `dashboardBackgroundScript` (padanan `jira:dashboardBackgroundScript`) yang dijalankan sekali per dashboard untuk logika latar yang dipakai beberapa gadget app.
  3. Modul legacy `jiraDashboardItems` (model Connect yang deprecated) tidak didukung; importer melaporkannya sebagai tidak didukung.
  4. Gadget app tampil di panel Add gadget (DSH-019) dengan nama app.
  5. Konfigurasi gadget app disimpan di dashboard item properties (DSH-057).
  6. Gadget app dirender di iframe sandbox dengan origin terpisah dan Content Security Policy ketat; komunikasi ke host lewat bridge pesan yang terbatas.
  7. Bila app di-uninstall atau dinonaktifkan, gadget app menampilkan pesan "Gadget ini tidak lagi tersedia" (DSH-024) dan instance-nya tetap dapat dihapus.
  8. Context yang diterima gadget berisi `dashboardId`, `gadgetId`, `accountId` viewer, dan mode `view`/`edit`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kerangka ekstensi umum di [25 Integrasi & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md). Gadget app Marketplace (mis. time tracking, BI, custom chart pihak ketiga) bukan fitur native.

### DSH-059 Dashboard audit & limits

- **Deskripsi:** Pencatatan aktivitas dan batasan sistem terkait dashboard dan gadget.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / platform
- **Aktor & izin:** Admin Situs (membaca audit log).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Audit log (USR-046) mencatat create, update, delete dashboard, perubahan owner, dan perubahan permission di kategori "Dashboards" beserta aktor, waktu, dan nilai sebelum/sesudah.
  2. ⚠ Satu dashboard dibatasi maksimal 50 gadget; menambah gadget ke-51 ditolak dengan pesan "Dashboard mencapai batas 50 gadget", dan dashboard dengan lebih dari 20 gadget menampilkan saran memecah dashboard.
  3. Gadget daftar work item menampilkan maksimal 50 baris per halaman.
  4. Query gadget yang melewati timeout (default 30 detik) dihentikan dan gadget menampilkan pesan "Gadget memerlukan waktu terlalu lama; persempit filter" dengan tombol coba lagi.
  5. Rate limiting REST dashboard dan gadget mengikuti batas platform (modul 27) dengan respons 429 dan header `Retry-After`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditempatkan di Fase 3 walau P3 karena batas dan timeout harus ada sejak dashboard dirilis agar tidak membebani database.

### DSH-060 Dashboards in Jira mobile app

- **Deskripsi:** Akses lihat-saja ke dashboard melalui aplikasi atau tampilan mobile.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna login di perangkat mobile; aturan izin sama dengan web.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar dashboard berbintang dan Recent tersedia di menu mobile.
  2. Gadget yang didukung dirender dalam satu kolom mengikuti urutan kolom 1, 2, 3.
  3. Gadget yang tidak didukung di mobile menampilkan placeholder "Buka di web" dengan tautan ke dashboard versi web.
  4. Pull-to-refresh memuat ulang semua gadget.
  5. ⚠ Edit mode tidak tersedia di mobile; tombol Edit tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rilis awal cukup memakai SPA responsif (DSH-016 butir 9); aplikasi native mengikuti keputusan modul 21/27. Daftar gadget yang didukung mobile menjadi pertanyaan terbuka.

## Legacy & Data Center-only (rujukan)

### DSH-061 My Jira Home (Dashboard as homepage)

- **Deskripsi:** Pengaturan profil DC untuk menjadikan dashboard sebagai halaman awal setelah login. Tidak ada di produk acuan Cloud; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Pilihan halaman awal di profil (Dashboard, Boards, Issue Navigator) tidak dibangun sebagai fitur modul ini; halaman awal diatur oleh navigasi modul 21.
  2. Bila kelak preferensi halaman awal dibangun di modul 21 dan memilih Dashboard, pengguna mendarat di dashboard terakhir atau system dashboard (DSH-016 butir 6).
  3. ⚠ Menu Dashboards tanpa id membuka dashboard terakhir yang dilihat; tidak ada pengaturan halaman awal khusus dashboard.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alasan: perilaku DC-only; padanan Cloud adalah halaman "For you/Your work" (modul 21).

### DSH-062 Wallboard / Slideshow view

- **Deskripsi:** Mode tampilan dashboard untuk layar besar dengan rotasi otomatis antar dashboard, hanya ada di DC lewat plugin bawaan. Dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Owner dashboard, pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "View as Wallboard" dari menu Tools tidak dibangun.
  2. Tema gelap berukuran besar untuk layar TV tidak dibangun; mode gelap umum aplikasi (modul 21) dan full-screen browser menjadi pengganti.
  3. Slideshow rotasi beberapa dashboard favorit dengan interval dan efek transisi tidak dibangun.
  4. Konfigurasi slideshow (pilih dashboard, durasi per slide, efek transisi) tidak dibangun.
  5. Pembatasan "hanya gadget yang mendukung wallboard" tidak berlaku karena wallboard tidak ada.
  6. Agile Wallboard Gadget dan Wallboard Spacer Gadget tidak dibangun (lihat DSH-063).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alasan: tidak ada di produk acuan Cloud; kebutuhan dipenuhi app Marketplace atau full-screen. Mode kiosk opsional dicatat sebagai pertanyaan terbuka.

### DSH-063 Legacy DC/Server-only gadgets

- **Deskripsi:** Gadget yang hanya ada di Server/Data Center atau sudah deprecated, dicatat agar tidak salah diimplementasikan sebagai fitur Cloud.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs, pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Text gadget (HTML bebas, nonaktif default, dapat diaktifkan admin) tidak dibangun karena risiko XSS; kebutuhan teks bebas dipenuhi Introduction gadget (DSH-053).
  2. Login gadget untuk pengguna anonim di system dashboard tidak dibangun; pengguna anonim diarahkan ke halaman login standar (modul 01).
  3. Admin gadget (ringkasan dan tugas admin) tidak dibangun; ringkasan admin ada di halaman Settings (modul 22).
  4. Agile Wallboard Gadget dan Wallboard Spacer Gadget tidak dibangun.
  5. Bugzilla ID Search gadget tidak dibangun.
  6. Gadget integrasi FishEye/Crucible/Bamboo (chart, recent changesets, plan summary) tidak dibangun; informasi development ada di modul 24.
  7. ⚠ Version Report gadget dari modul agile lama tidak dibangun; laporan Version report tersedia di modul 16.
  8. Pilihan warna chrome gadget di UI tidak dibangun; warna hanya lewat REST (DSH-022, DSH-057).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Importer dari DC (modul 25) melewati gadget-gadget ini dan melaporkannya dalam ringkasan impor sebagai "gadget tidak didukung".

## Pertanyaan terbuka

- Apakah dashboard (dan gadget P0 seperti Assigned to Me dan Filter Results) perlu dimajukan ke Fase 2 sebagai halaman beranda tim, atau halaman "Your work" modul 21 sudah cukup sampai Fase 3?
- Berapa ukuran halaman default direktori dashboard (20 atau 25 baris) dan apakah pengguna dapat mengubahnya?
- Apakah keunikan nama dashboard per owner ditegakkan keras (galat) atau cukup peringatan, mengingat produk acuan Cloud tidak terdokumentasi jelas?
- Apakah opsi Public sharing dan akses anonim dashboard disediakan sama sekali di distribusi self-hosted, atau dihapus permanen demi keamanan?
- Apakah status minimize gadget disimpan per pengguna (asumsi saat ini) atau per dashboard?
- Bagaimana definisi populasi Average Number of Times in Status: hanya work item yang resolved di periode, atau semua work item yang memasuki status di periode?
- Apakah status `archived` pada pencarian dashboard REST perlu didukung dengan fitur arsip dashboard di UI, atau hanya dikembalikan kosong untuk kompatibilitas?
- Apakah perlu disediakan mode kiosk/full-screen dashboard native sebagai pengganti Wallboard, mengingat kebutuhan layar TV tim?
- Apakah Work Items Calendar perlu ekspor/langganan iCal dan view mingguan?
- Gadget mana saja yang wajib didukung pada tampilan mobile, dan apakah aplikasi mobile native termasuk roadmap?
- Berapa batas maksimal gadget per dashboard dan timeout query gadget yang dapat diterima (asumsi 50 gadget dan 30 detik)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** penyedia gadget dan dashboard (mis. gadget time tracking, BI/analitik, custom chart, wallboard pihak ketiga) — tidak dibangun sebagai fitur native; produk ini hanya menyediakan titik ekstensi gadget (DSH-058) dan REST API (DSH-056, DSH-057).
- **Fitur Data Center-only (dc-only-or-deprecated)** — dicatat sebagai rujukan saja: My Jira Home (DSH-061), Wallboard/Slideshow (DSH-062), dan gadget legacy Server/DC seperti Text, Login, Admin, Agile Wallboard, Wallboard Spacer, Bugzilla ID Search, FishEye/Crucible/Bamboo, Version Report lama, serta pilihan warna chrome gadget di UI (DSH-063); juga tab Favourite/My/Popular di halaman Manage Dashboards DC, interval auto-refresh legacy (30 menit/1 jam/2 jam), opsi "Start from" di dialog create, rendering chart sebagai gambar di server, dan modul Connect `jiraDashboardItems`.
- **Layanan lintas produk milik vendor acuan** — aliran aktivitas dari produk lain milik vendor acuan (wiki, repositori kode) di luar application link generik, analitik lintas produk (setara Atlassian Analytics/Data Lake, ditangani modul 16 bila dibangun), dan dashboard berbasis AI generatif.
- **Laporan halaman penuh dan kalkulasinya** — dimiliki [16 Laporan & Insight](16-laporan-insight.md); modul ini hanya menampilkan versi gadget.
- **Dashboard dan laporan Service Management** (gadget SLA, antrian, CSAT) — dibahas di [31 JSM Queues, SLA & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md).
