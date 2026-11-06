# 05 — Konfigurasi Work Item: Tipe, Field, Layar & Skema

Modul ini mendefinisikan seluruh lapisan konfigurasi yang menentukan "bentuk" sebuah work item di produk ini: work type (issue type) dan hierarkinya (Subtask, Standard, Epic, serta level kustom di atas Epic), katalog system field dan field yang dikelola sistem (Parent, Sprint, Rank, Story Points, Team, dst.), tipe custom field beserta context, opsi, default value, dan searcher, serta mesin skema company-managed (issue type scheme, field configuration (scheme), screen, screen scheme, issue type screen scheme, priority scheme) dan padanannya yang disederhanakan di team-managed project (field library dan layout per work type). Modul ini juga mencakup nilai global yang dipakai semua project (priority, resolution, link type, toggle fitur work item, pengaturan time tracking). Perilaku runtime work item (membuat, mengubah, melihat) ada di modul 06; modul ini hanya mengatur aturan konfigurasinya. Kondisi repo saat ini baru memiliki auth dan user, sehingga seluruh entitas di modul ini (tabel work type, field, context, option, screen, scheme, priority, resolution, link type) adalah pekerjaan baru yang harus dirancang bersama model data di modul 28.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission Administer Jira/Administer projects, field user/group sebagai sumber grant, issue security.
- [04 Proyek](04-proyek.md) — tipe project (company-managed vs team-managed), template, Project settings, Features toggle.
- [06 Work Item Inti](06-work-item-inti.md) — create/view/edit, Create dialog, issue view, konversi subtask, clone, link.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — workflow scheme, transition screen, post function resolution.
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — Rank, Sprint, Flagged, estimation statistic.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — searcher, operator, fungsi JQL field.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — worklog, estimasi, dan detail Time Tracking Settings (TIME-001, TIME-002).
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Fix/Affects versions, Components, Version picker.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget yang menampilkan nilai field, termasuk hasil Formula custom field (CFG-062).
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — hierarki di atas Epic, Target start/end, Team, Issue color.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — General configuration, audit log.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — endpoint issuetype, field, screen, priority, resolution.
- [24 Integrasi DevOps](24-integrasi-devops.md) — field Development.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — tipe custom field dari ekstensi.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — Category field, view bisnis.
- [28 Model Data Inti](28-model-data-inti.md) — skema tabel dan event domain konfigurasi.
- [07 Editor Konten Kaya](07-editor-konten-kaya.md), [08 Komentar & Lampiran](08-komentar-lampiran-aktivitas.md), [16 Laporan](16-laporan-insight.md), [20 Otomasi](20-otomasi.md), [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — format rich text, penyimpanan berkas, laporan estimasi, alternatif otomasi, guardrail dan job background.
- [29 Service Management: Request Type](29-jsm-service-project-request-type-formulir.md) dan [36 Product Discovery: Ideas & Fields](36-jpd-ideas-field-views.md) — work type dan field khusus produk tersebut.

## Persona terkait

- **Admin Situs** — mengelola work type global, custom field, context, field configuration, screen, scheme, priority, resolution, link type, dan toggle global.
- **Admin Proyek/Project Lead** — mengelola work type, field, dan layout di team-managed project; melihat konfigurasi dan mengatur layout/screen yang tidak dibagikan di company-managed project (extended project administration).
- **Scrum Master/Team Lead** — bergantung pada Sprint, Rank, Story Points, Flagged, dan estimation statistic.
- **Developer/Anggota Tim** — mengisi field saat create/edit, memakai Parent, Flagged, priority, dan link.
- **Product Owner/PM** — merancang hierarki (Epic dan level di atasnya), Team, Target start/end.
- **Pengguna Tim Bisnis** — memakai Category field dan field sederhana di project bisnis.
- **Integrator/Developer API** — bergantung pada ID stabil work type/field/option, searcher JQL, dan tipe field dari ekstensi.

## User story modul

- Sebagai Admin Situs, saya ingin membuat work type custom dengan ikon sendiri dan memasukkannya ke issue type scheme agar setiap project company-managed hanya melihat tipe yang relevan (CFG-001, CFG-003, CFG-004, CFG-008).
- Sebagai Admin Proyek/Project Lead, saya ingin menambah, mengganti nama, dan menghapus work type di team-managed project saya tanpa bantuan Admin Situs agar tim bisa beradaptasi cepat (CFG-002, CFG-041, CFG-042).
- Sebagai Product Owner/PM, saya ingin menambah level Initiative di atas Epic agar roadmap lintas tim bisa digulung dari Story sampai inisiatif strategis (CFG-006, CFG-007, CFG-013).
- Sebagai Admin Situs, saya ingin memindahkan work item secara aman saat sebuah work type dihapus atau scheme diganti agar tidak ada data status atau field yang hilang tanpa peringatan (CFG-009).
- Sebagai Admin Situs, saya ingin membuat custom field dengan tipe yang tepat (teks, angka, tanggal, pilihan, user) dan membatasinya lewat context agar field hanya muncul di project dan work type yang membutuhkannya (CFG-025 s.d. CFG-031, CFG-034, CFG-035).
- Sebagai Admin Situs, saya ingin mengatur field mana yang wajib, tersembunyi, dan tampil di create/edit/view lewat field configuration dan screen agar kualitas data terjaga (CFG-043, CFG-044, CFG-046, CFG-047, CFG-048).
- Sebagai Admin Proyek/Project Lead, saya ingin menyusun tata letak issue view (Description fields, Context fields, Hide when empty) agar informasi penting langsung terlihat (CFG-042, CFG-049).
- Sebagai Developer/Anggota Tim, saya ingin tahu mengapa sebuah field tidak muncul di work item agar saya bisa meminta perbaikan yang tepat ke admin (CFG-051).
- Sebagai Scrum Master/Team Lead, saya ingin field Sprint, Rank, Story Points, dan Flagged berperilaku konsisten agar backlog, board, dan laporan velocity akurat (CFG-014, CFG-015, CFG-016, CFG-018).
- Sebagai Admin Situs, saya ingin mengelola priority, priority scheme, resolution, dan link type secara terpusat agar istilah dan urutan seragam di seluruh site (CFG-053, CFG-054, CFG-055, CFG-056).
- Sebagai Admin Situs, saya ingin menyalakan atau mematikan voting, watching, unassigned, subtask, linking, dan time tracking secara global agar produk sesuai kebijakan organisasi (CFG-057, CFG-058).
- Sebagai Admin Situs, saya ingin melihat pemakaian field dan memindahkan field tak terpakai ke trash agar site tidak melampaui guardrail kinerja (CFG-034, CFG-039).
- Sebagai Integrator/Developer API, saya ingin ID work type, field, dan option tetap stabil saat di-rename agar JQL, otomasi, dan integrasi tidak rusak (CFG-003, CFG-034, CFG-037).
- Sebagai Admin Situs, saya ingin membuat custom field Formula yang menghitung nilai dari field lain agar tim tidak perlu mengekspor data ke spreadsheet hanya untuk menghitung sisa anggaran atau umur work item (CFG-062).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 19 |
| P1 | 24 |
| P2 | 12 |
| P3 | 7 |
| **Total** | **62** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 11 |
| Fase 2 | 30 |
| Fase 3 | 11 |
| Fase 4 | 5 |
| Fase — (dc-only/deprecated, rujukan) | 5 |

Catatan: inventaris berisi 67 fitur; 6 pasang yang tumpang tindih digabung menjadi satu heading dengan baris **Alias inventaris**, sehingga file ini memuat 61 heading fitur dari inventaris, ditambah 1 fitur native di luar inventaris (CFG-062 Formula custom field, bertanda ⚠) sehingga total 62 heading.

Strategi urutan implementasi: Fase 1 membangun model data konfigurasi minimal yang tetap kompatibel dengan model penuh, yaitu work type bawaan per template, hierarki tiga level (Subtask/Standard/Epic), katalog system field, field yang dikelola sistem untuk agile (Parent, Sprint, Rank, Story point estimate), priority bawaan, dan link type bawaan; semua dikonfigurasi lewat seed dan default scheme tanpa UI admin yang lengkap. Fase 2 membuka konfigurabilitas: custom field berikut context/opsi/default, field library dan layout team-managed, lalu mesin skema company-managed (issue type scheme, field configuration, screen, screen scheme, issue type screen scheme, issue layout) beserta wizard migrasi work type. Sejak Fase 1 setiap project company-managed sudah harus menunjuk ke default scheme secara eksplisit di basis data agar Fase 2 tidak memerlukan migrasi data besar. Fase 3 menambah alat skala organisasi (priority scheme, field scheme sederhana, guardrail, searcher, Find your field, berbagi scheme) serta Formula custom field (CFG-062) yang bergantung pada searcher dan reindex background, sedangkan Fase 4 menambah hierarki Premium di atas Epic, Team, Target start/end, field Development, dan tipe field dari ekstensi. Fitur dc-only/deprecated hanya dicatat sebagai rujukan dan alias kompatibilitas.

## Fitur

## Work type & hierarki

### CFG-001 Work Types (Issue Types) – built-in & custom

- **Deskripsi:** Setiap work item memiliki tepat satu work type yang menentukan ikon, field, workflow, dan level hierarkinya. Produk ini menyediakan tipe bawaan sesuai template project dan mengizinkan admin membuat tipe custom.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) untuk tipe global company-managed; Admin Proyek (Administer projects) untuk tipe team-managed; semua pengguna dengan Browse projects melihat tipe.
- **Aturan perilaku (acceptance criteria):**
  1. Template software menghasilkan tipe bawaan Epic, Story, Task, Bug, dan Subtask; template business menghasilkan Task dan Subtask, dengan Epic dan tipe lain sebagai tambahan opsional. Tipe khusus Service Management dan Product Discovery didefinisikan di modul 29 dan 36, bukan di sini.
  2. Di company-managed, work type bersifat global dan dibagikan ke project lewat issue type scheme (CFG-008); setiap project hanya melihat subset tipe di scheme-nya dan memiliki default type sendiri.
  3. Di team-managed, work type bersifat lokal per project; Admin Proyek dapat membuat, mengubah, dan menghapusnya, dan perubahan tersebut tidak pernah memengaruhi project lain.
  4. Atribut work type: name (wajib, unik case-insensitive dalam cakupannya, yaitu global untuk company-managed atau per project untuk team-managed), description, avatar/icon (bawaan atau upload), kategori Standard atau Sub-task, dan hierarchy level.
  5. Default work type project otomatis terpilih di Create dialog; bila default tidak diatur, tipe pertama sesuai urutan scheme yang dipilih.
  6. Menghapus work type yang masih dipakai work item: di company-managed admin wajib memilih tipe pengganti atau menjalankan migrasi lewat scheme (CFG-009); di team-managed sistem membuka wizard pemindahan tipe. Penghapusan tanpa pengganti ditolak dengan pesan "Work type masih dipakai oleh N work item".
  7. Kategori Standard dan Sub-task tidak dapat ditukar setelah tipe dibuat; permintaan API yang mengubah kategori ditolak (HTTP 400) dengan pesan galat yang menjelaskan bahwa kategori bersifat tetap.
  8. Ikon work type tampil konsisten di card board, list/backlog, hasil pencarian, link work item, breadcrumb, dan notifikasi (in-app dan email).
  9. Rename tipe bawaan diizinkan; rename Epic (mis. menjadi "Feature") tidak mengubah level hierarkinya dan tidak memutus relasi Parent yang ada.
  10. ⚠ Nama work type dapat diterjemahkan per bahasa antarmuka; fitur ini mengikuti keputusan CFG-061 dan bersifat opsional di produk ini.
  11. ⚠ Jumlah work type per team-managed project dibatasi oleh batas yang dapat dikonfigurasi di tingkat sistem; saat batas tercapai tombol Add dinonaktifkan dan API mengembalikan galat "Batas jumlah work type tercapai".
  12. Subtask dapat dinonaktifkan secara global untuk company-managed lewat Settings > Work items > Sub-tasks (CFG-005); saat nonaktif, tipe Sub-task disembunyikan dari Create dialog.
- **Team-managed vs company-managed:** Company-managed memakai tipe global yang dibagikan lewat scheme; team-managed memakai tipe lokal per project yang dikelola Admin Proyek (detail di CFG-002).
- **Catatan & ketergantungan:** Fondasi bagi CFG-003, CFG-006, CFG-008, workflow scheme di modul 09, dan Create dialog di modul 06. UI memakai istilah "work type"; API dan JQL tetap memakai `issuetype` demi kompatibilitas. Tabel `work_types` baru dibuat di modul 28 dengan kolom `scope` (global/project) dan `project_id` nullable.

### CFG-002 Team-managed project work types

- **Deskripsi:** Pengelolaan work type per project untuk team-managed project, independen dari scheme global, dari halaman Project settings > Work types.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (permission Administer projects pada project tersebut); Admin Situs dapat bertindak atas nama Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Project settings > Work types menampilkan daftar work type project (Epic, Story, Task, Bug, Subtask sesuai template) lengkap dengan ikon dan jumlah work item.
  2. ⚠ Add work type meminta nama, deskripsi, dan ikon; jumlah work type per project dibatasi oleh batas sistem yang dapat dikonfigurasi, dan pesan galat ditampilkan saat batas tercapai.
  3. Edit nama, deskripsi, dan ikon hanya berlaku untuk project tersebut; tipe bernama sama di project lain tidak berubah.
  4. Delete work type yang masih dipakai mewajibkan pemilihan work type pengganti untuk semua work item yang ada sebelum penghapusan diproses.
  5. Toggle di Project settings > Features (Epics/Subtasks, Estimation) memengaruhi ketersediaan work type dan field terkait; mematikan Subtasks menyembunyikan tipe Subtask dari Create dialog tanpa menghapus data.
  6. Setiap work type memiliki layout field sendiri (CFG-042) dan dapat memiliki workflow sendiri (modul 09).
  7. Work type team-managed tidak dapat dibagikan ke project lain; tipe bernama sama di dua project adalah entitas berbeda dengan ID berbeda.
  8. Pengelolaan memerlukan Administer projects pada project itu, bukan Administer Jira; pengguna tanpa izin menerima HTTP 403.
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah CFG-003 dan CFG-008.
- **Catatan & ketergantungan:** Bergantung pada tipe project di modul 04 dan Features toggle project. Memakai tabel `work_types` yang sama dengan CFG-001 (scope = project).

### CFG-003 Work types (Issue types) global administration

- **Deskripsi:** Halaman admin global (Settings > Work items > Work types) untuk mengelola semua work type Standard dan Sub-task yang dipakai company-managed project; menjadi fondasi issue type scheme, screen scheme, field configuration, dan workflow scheme.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (global permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar menampilkan setiap work type dengan ikon, nama, deskripsi, tipe (Standard/Sub-task), dan hierarchy level.
  2. Daftar dapat dicari berdasarkan nama dan difilter berdasarkan tipe (Standard/Sub-task).
  3. Dialog Add work type berisi Name (wajib, unik case-insensitive termasuk terhadap nama sub-task type), Description (opsional), dan Type (Standard issue type / Sub-task issue type); nama duplikat ditolak dengan pesan "Nama work type sudah dipakai".
  4. Nama dibatasi maksimum 60 karakter; input lebih panjang ditolak di UI dan API.
  5. ⚠ Tipe Standard/Sub-task tidak dapat diubah setelah dibuat; field Type ditampilkan read-only di dialog Edit.
  6. Work type baru otomatis masuk Default Issue Type Scheme dan Base level hierarki, tetapi tidak otomatis tersedia di project yang memakai scheme lain; UI menampilkan petunjuk untuk menambahkannya ke scheme terkait.
  7. Edit work type mengubah nama, deskripsi, dan avatar; rename berlaku global di semua company-managed project dan ID tetap sama sehingga JQL berbasis ID dan aturan otomasi tidak rusak.
  8. Rename tipe Epic (mis. menjadi "Feature") lewat Edit membuat label di board, backlog, dan timeline ikut berubah.
  9. ⚠ Delete work type yang masih dipakai wajib memilih work type pengganti dan memicu migrasi (CFG-009); work type terakhir di level hierarkinya tidak dapat dihapus.
  10. ⚠ Work type bawaan adalah Epic, Story, Task, Bug, dan Sub-task (bervariasi per template); Epic adalah satu-satunya work type bawaan di Epic level.
  11. ⚠ Sebelum menghapus, admin dapat melihat daftar scheme yang mereferensikan work type tersebut (issue type scheme, workflow scheme, screen scheme, field configuration scheme).
  12. ID work type stabil dan dipakai di REST API, JQL (mis. `issuetype = 10001`), dan otomasi.
  13. Work type team-managed tidak muncul di halaman global ini karena cakupannya per project.
  14. Seluruh aksi di halaman ini memerlukan global permission Administer Jira; pengguna lain menerima HTTP 403.
- **Team-managed vs company-managed:** Hanya company-managed; tipe team-managed dikelola di CFG-002.
- **Catatan & ketergantungan:** Istilah UI "work item/work type/space" vs istilah API "issue/issuetype/project" diputuskan di glosarium modul 00. Diturunkan ke P1 karena MVP dapat berjalan dengan tipe bawaan dan model team-managed.

### CFG-004 Issue type icons (avatars)

- **Deskripsi:** Pengelolaan ikon/avatar untuk setiap work type, termasuk pustaka ikon bawaan dan upload ikon kustom.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs untuk tipe global; Admin Proyek untuk tipe team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat memilih ikon dari pustaka avatar bawaan produk ini (berbagai warna dan simbol) yang di-seed saat instalasi.
  2. ⚠ Admin dapat meng-upload ikon kustom PNG/JPG/GIF (dukungan SVG ditentukan saat implementasi) dengan crop ke persegi; ukuran tampilan 16-24px dan sumber yang direkomendasikan minimal 48x48px; file di luar format atau melebihi batas ukuran ditolak dengan pesan galat.
  3. Avatar kustom yang di-upload dapat dihapus; avatar sistem tidak dapat dihapus (tombol hapus tidak tersedia, API mengembalikan 403).
  4. Ikon ditampilkan di board card, backlog, issue view, hasil search/list, picker, breadcrumb, dan email notifikasi.
  5. Di team-managed, ikon work type diatur per project dari Project settings > Work types.
  6. Avatar disimpan dengan ID dan diakses lewat API avatar universal dengan tipe `issuetype`.
  7. Bila avatar yang dipakai terhapus, work type otomatis memakai ikon default tanpa error tampilan.
- **Team-managed vs company-managed:** Mekanisme sama; lokasi pengaturan berbeda (global vs Project settings).
- **Catatan & ketergantungan:** Memakai penyimpanan berkas yang sama dengan lampiran (modul 08) atau storage avatar pengguna (modul 02).

### CFG-005 Sub-task issue types & sub-task settings

- **Deskripsi:** Konfigurasi work type level Sub-task dan pengaktifan fitur subtask secara global.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (toggle global, tipe global); Admin Proyek (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle global enable/disable sub-tasks tersedia di Settings > Work items > Sub-tasks; saat nonaktif, subtask tidak dapat dibuat lewat UI maupun API (galat "Sub-task dinonaktifkan").
  2. ⚠ Sub-tasks tidak dapat dinonaktifkan selama masih ada subtask di instance; sistem menampilkan jumlah subtask yang masih ada.
  3. Admin dapat membuat beberapa sub-task work type (mis. Sub-task, Sub-bug, QA Sub-task).
  4. Subtask wajib memiliki parent dari standard work type (Base level) di project yang sama; subtask tidak dapat memiliki subtask.
  5. Subtask mewarisi project, issue security level, dan sprint dari parent; perubahan nilai tersebut di parent ikut diterapkan ke subtask.
  6. Sub-task work type harus ada di issue type scheme project agar dapat dibuat di project tersebut.
  7. Konversi standard ke subtask (Convert to sub-task) dan sebaliknya (Convert to issue) hanya ditawarkan bila ada sub-task/standard type yang tersedia di scheme project.
  8. ⚠ Team-managed menyediakan work type "Subtask" default yang dapat di-toggle lewat Features; penambahan sub-task type lain di team-managed ditentukan saat implementasi.
- **Team-managed vs company-managed:** Company-managed memakai toggle global dan sub-task type global; team-managed memakai toggle Features per project.
- **Catatan & ketergantungan:** Operasi konversi dan pembuatan subtask ada di modul 06; pewarisan sprint di modul 11; security level di modul 03.

### CFG-006 Work type hierarchy (Issue type hierarchy)

- **Alias inventaris:** Issue Type Hierarchy (levels)
- **Deskripsi:** Work type disusun ke dalam level hierarki Sub-task (-1), Base/Standard (0: Story/Task/Bug), Epic (1), dan level kustom di atas Epic untuk tier Premium, dengan halaman Settings > Work items > Work type hierarchy untuk menampilkan dan mengaturnya. Level menentukan relasi parent-child yang sah.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (mengatur halaman hierarki); semua pengguna (relasi parent-child saat create/edit).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Level -1 (Subtask) wajib memiliki parent dari level 0; di company-managed subtask langsung di bawah Epic tidak diizinkan secara default dan di team-managed tidak didukung.
  2. Level 0 (Standard) boleh tanpa parent; bila memiliki parent, parent harus dari level 1 (Epic).
  3. Level 1 (Epic) hanya dapat memiliki parent bila level kustom di atas Epic (CFG-007) aktif.
  4. Validasi relasi: subtask tidak boleh memiliki subtask, parent tidak boleh dari level yang sama, dan relasi melingkar ditolak dengan galat "Relasi parent membentuk siklus".
  5. ⚠ Di company-managed, parent dan child boleh lintas project (kecuali subtask); di team-managed parent lintas project dibatasi ke project yang dapat dilihat pengguna dan dikonfigurasi per project.
  6. Hierarki dipakai oleh Parent field (CFG-013), panel child work items, Timeline, Plans, dan JQL (`parent`, `childIssuesOf`).
  7. Parent menampilkan rollup progres (jumlah child selesai dibanding total, berdasarkan status category Done).
  8. Level bawaan di semua tier: Sub-task level (-1), Base level (0) berisi semua standard work type, dan Epic level (1).
  9. Standard work type baru default masuk Base level; sub-task type otomatis masuk Sub-task level.
  10. ⚠ Nama level dapat di-rename (mis. level "Epic" menjadi "Feature") terpisah dari nama work type; level Sub-task dan Base tidak dapat di-rename.
  11. Child hanya dapat memiliki parent tepat satu level di atasnya; pilihan Parent di UI difilter sesuai aturan ini.
  12. Hierarki dipakai oleh Parent field, Timeline, Plans, serta JQL `parent` dan `childIssuesOf`; `portfolioChildIssuesOf` diterima sebagai alias legacy.
  13. Pada tier free/standard halaman hierarki bersifat read-only untuk level di atas Epic (tombol tambah level tidak tersedia, disertai info upgrade).
  14. Team-managed project mengikuti level global Epic/Story/Subtask.
- **Team-managed vs company-managed:** Team-managed memakai tiga level global tanpa konfigurasi per project dan tanpa subtask di bawah Epic; company-managed mengikuti konfigurasi hierarki site.
- **Catatan & ketergantungan:** Butir 1-7 dari "Issue Type Hierarchy (levels)" (issues-core, P0), butir 8-14 dari "Work type hierarchy (Issue type hierarchy)" (admin-issue-config, P1); prioritas gabungan memakai yang lebih tinggi. Inventaris menyebut subtask langsung di bawah Epic "mungkin" di company-managed; produk ini memilih tidak mengizinkannya secara default agar konsisten dengan CFG-005 butir 4 dan CFG-013 butir 1 (lihat Pertanyaan terbuka). Dulu dikonfigurasi di Advanced Roadmaps > Hierarchy configuration; di produk ini halaman berada di Settings > Work items. Rollup juga dipakai modul 18.

### CFG-007 Custom hierarchy levels above Epic

- **Alias inventaris:** Custom Hierarchy Levels above Epic
- **Deskripsi:** Kemampuan tier Premium/Enterprise untuk menambah level hierarki di atas Epic (mis. Initiative, Theme) dan memetakan work type ke level tersebut; dipakai di Plans, Timeline, dan Parent field.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs (Administer Jira) pada site dengan tier Premium.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > Work items > Work type hierarchy memungkinkan menambah, mengurutkan, dan menghapus level 2 ke atas.
  2. Satu level dapat dipetakan ke satu atau lebih work type.
  3. Parent field mendukung parent dari level tepat di atasnya; data Parent Link lama dimigrasikan ke Parent (lihat CFG-024).
  4. Work type level atas biasanya dibuat di company-managed project khusus; UI menyarankan hal ini tetapi tidak memaksanya.
  5. ⚠ Menghapus level membuat work type terkait kembali ke level standar dan relasi parent level atas dilepas setelah konfirmasi yang menyebut jumlah relasi terdampak.
  6. Level kustom dipakai di Plans/Timeline dan JQL `childIssuesOf`; `portfolioChildIssuesOf` hanya alias legacy.
  7. Downgrade ke tier Standard membuat level tambahan berhenti berfungsi (tidak tampil di Parent field dan Plans).
  8. Admin dapat menambah level 2, 3, dst. dengan nama kustom.
  9. Admin membuat standard work type baru (mis. Initiative) lalu memetakannya ke level kustom; satu work type hanya boleh berada di satu level.
  10. Satu level kustom dapat berisi lebih dari satu work type.
  11. Level kustom dapat diurutkan ulang (reorder) dengan drag; urutan menentukan relasi parent yang sah.
  12. ⚠ Menghapus level kustom mengembalikan work type ke Base level dan melepaskan relasi parent di atasnya; tindakan ini tercatat di audit log.
  13. Work item di level kustom dapat menjadi parent Epic lintas project.
  14. Level kustom tampil di Plans dan dapat dipakai oleh work item team-managed lewat Plans/Parent.
  15. ⚠ Setelah downgrade dari Premium, level kustom tidak dapat dikelola; data relasi tetap disimpan dan aktif kembali bila tier di-upgrade.
  16. ⚠ Jumlah level kustom dibatasi oleh batas sistem yang dapat dikonfigurasi; melebihi batas ditolak dengan pesan galat.
- **Team-managed vs company-managed:** Level dikonfigurasi global; work type level atas dibuat di company-managed, sedangkan Epic team-managed dapat memakai parent level atas lewat Plans/Parent.
- **Catatan & ketergantungan:** Dua entri inventaris yang namanya hanya berbeda kapitalisasi ("Custom Hierarchy Levels above Epic" dari issues-core dan "Custom hierarchy levels above Epic" dari admin-issue-config) digabung di heading ini; butir 1-7 dari issues-core, 8-16 dari admin-issue-config. Bergantung pada gating tier dan perilaku downgrade (ADM-054 di modul 22) serta Plans (modul 18); sisi perencanaan dari konfigurasi hierarki dirinci di PLAN-032, sedangkan aturan admin lengkap tetap di heading ini dan CFG-006.

### CFG-008 Issue type schemes

- **Deskripsi:** Kumpulan work type yang tersedia untuk satu atau lebih company-managed project, lengkap dengan default work type dan urutan tampilan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira); Admin Proyek hanya melihat scheme yang dipakai project-nya.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar issue type scheme menampilkan nama, deskripsi, work type yang termuat, dan project yang memakai.
  2. Default Issue Type Scheme berisi semua work type, dipakai project tanpa scheme khusus, tidak dapat dihapus, dan isinya tidak dapat diatur manual (selalu sinkron dengan semua work type).
  3. Add issue type scheme meminta nama dan deskripsi, lalu admin menarik work type dari "Available" ke "Issue types for current scheme".
  4. Admin dapat mengatur urutan work type; urutan ini dipakai di dropdown Create.
  5. Default work type scheme terpilih otomatis di Create dialog dan harus termasuk dalam scheme.
  6. Edit scheme yang menghapus work type yang masih memiliki work item di project terkait memicu wizard migrasi (CFG-009).
  7. Scheme dapat disalin (Copy) menjadi scheme baru yang independen.
  8. Delete scheme membuat project terkait kembali ke Default Issue Type Scheme (dengan migrasi bila perlu).
  9. Scheme dapat diasosiasikan ke project dari halaman admin atau dari Project settings > Work types > Actions > Use a different scheme; asosiasi memicu migrasi bila ada work type yang tidak lagi tersedia.
  10. Satu project hanya memiliki satu issue type scheme; satu scheme dapat dipakai banyak project.
  11. Project settings menampilkan indikator "dipakai bersama oleh N project".
  12. Daftar scheme dapat dicari berdasarkan nama dan difilter berdasarkan project.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed tidak memakai scheme.
- **Catatan & ketergantungan:** Terkait CFG-059 (scheme default dan berbagi). Fase 1 memakai Default Issue Type Scheme hasil seed tanpa UI.

### CFG-009 Issue type migration wizard

- **Deskripsi:** Wizard untuk memindahkan work item dari work type yang dihapus atau dikeluarkan dari scheme ke work type lain, termasuk mapping status dan field.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (company-managed); Admin Proyek (team-managed, versi sederhana saat menghapus work type).
- **Aturan perilaku (acceptance criteria):**
  1. Wizard dipicu saat menghapus work type, mengganti issue type scheme project, atau menghapus work type yang masih dipakai dari scheme.
  2. Langkah 1: admin memilih work type tujuan untuk setiap work type terdampak per project.
  3. Langkah 2: admin memetakan setiap status lama ke status di workflow tujuan bila workflow berbeda; langkah dilewati bila workflow sama.
  4. Langkah 3: admin mengisi nilai untuk field wajib yang belum terisi menurut field configuration tujuan.
  5. Langkah 4: admin mengonfirmasi ringkasan, lalu migrasi diproses di background dengan progress bar dan dapat dipantau setelah halaman ditutup.
  6. Subtask hanya dapat dimigrasi ke sub-task type dan standard hanya ke standard; pilihan yang tidak sah tidak ditampilkan.
  7. Nilai field yang tidak berlaku di konfigurasi tujuan dihapus setelah peringatan yang mencantumkan nama field terdampak.
  8. Setiap perubahan work type, status, dan field tercatat di history work item.
- **Team-managed vs company-managed:** Team-managed memakai versi sederhana (pilih tipe pengganti dan mapping status) saat menghapus work type.
- **Catatan & ketergantungan:** Memakai job background (modul 27) dan mapping status dari modul 09. Logika yang sama dipakai operasi Move work item di modul 06.

### CFG-010 Configure terminology (Sprint/Epic rename)

- **Deskripsi:** Fitur Data Center untuk mengganti istilah global "Sprint" dan "Epic" di seluruh UI; dicatat sebagai rujukan dan tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: rename istilah Sprint dan Epic (singular dan plural) di seluruh UI tidak tersedia sebagai pengaturan global.
  2. Bila kelak dibangun, rename istilah tidak boleh mengubah keyword JQL maupun REST API.
  3. Pengganti di produk ini: hanya nama work type (mis. Epic, CFG-003) dan nama hierarchy level (CFG-006) yang dapat di-rename; Sprint dan Story Points tidak dapat di-rename.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## System field & field yang dikelola sistem

### CFG-011 Jira-managed system fields catalog

- **Deskripsi:** Daftar system field bawaan yang wajib tersedia di produk ini beserta perilaku konfigurasinya; system field selalu ada di setiap work item dan tidak dapat dihapus.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (hide/required lewat field configuration); sistem (field turunan).
- **Aturan perilaku (acceptance criteria):**
  1. Summary wajib diisi, maksimum 255 karakter, tidak dapat di-hide maupun dibuat opsional; input kosong ditolak dengan galat "Summary wajib diisi".
  2. Description adalah field rich text (format dokumen modul 07).
  3. Tersedia system field Issue Type, Project, Key, Reporter, Creator, Assignee, Priority, Resolution, Status, Status Category, dan Status Category Changed.
  4. Tersedia system field Labels, Components, Affects versions, dan Fix versions (perilaku versi/komponen di modul 15).
  5. Tersedia system field Due date, Start date, dan Environment (paragraph).
  6. Tersedia system field Attachment, Linked Issues, Comment, Sub-tasks, dan Parent.
  7. Tersedia system field time tracking: Original estimate, Remaining estimate, Time spent, Log Work, Progress, Σ Progress, dan Work Ratio (CFG-017).
  8. Tersedia system field Security Level, Watchers, Votes, Created, Updated, Resolved, dan Last Viewed.
  9. Field Development bersifat read-only (CFG-023); Request participants dan field Service Management lain hanya aktif bila modul Service Management aktif.
  10. System field tidak dapat dihapus; sebagian (mis. Environment, Due date, Labels) dapat di-hide lewat field configuration (CFG-043).
  11. Field turunan (Created, Updated, Resolved, Votes, Σ) bersifat read-only; upaya mengubahnya lewat API ditolak.
- **Team-managed vs company-managed:** Field yang sama tersedia; team-managed tidak memakai field configuration sehingga visibilitasnya diatur lewat layout (CFG-042) dan Features toggle.
- **Catatan & ketergantungan:** Katalog ini menjadi kontrak ID field sistem (mis. `summary`, `description`, `issuetype`, `duedate`) untuk REST API (modul 23) dan JQL (modul 12). Kolom inti disimpan langsung di tabel work item (modul 28), bukan sebagai custom field.

### CFG-012 Locked fields

- **Deskripsi:** Field yang dikelola sistem produk ini atau oleh ekstensi dan tidak dapat diubah konfigurasinya secara bebas.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Sistem (pemilik field); Admin Situs (hanya pengaturan yang diizinkan).
- **Aturan perilaku (acceptance criteria):**
  1. Locked field ditandai ikon gembok di halaman Fields (CFG-034).
  2. Locked field tidak dapat dihapus, di-rename, atau diubah tipenya; aksi tersebut tidak tampil di menu.
  3. Sebagian locked field mengizinkan pengaturan context/work type (mis. field estimasi); daftar pengaturan yang diizinkan didefinisikan per field.
  4. Contoh locked field: Sprint, Rank, Team, Development, Target start/end, dan Story point estimate.
  5. Field yang didaftarkan ekstensi (tipe `customField` dari ekstensi) juga bersifat locked.
  6. Saat admin mencoba mengedit, ditampilkan pesan penjelasan "Field ini dikelola oleh sistem/ekstensi".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel field memiliki kolom `managed_by` (system/app/admin) dan daftar kapabilitas yang diizinkan. Terkait CFG-013 s.d. CFG-023 dan CFG-033.

### CFG-013 Parent field

- **Deskripsi:** System field relasi parent-child sesuai work type hierarchy, menggantikan Epic Link dan Parent Link.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues (mengubah parent); Admin Situs (konfigurasi).
- **Aturan perilaku (acceptance criteria):**
  1. Satu field "Parent" dipakai untuk semua level: subtask ke base, base ke Epic, dan Epic ke level kustom di atasnya.
  2. Parent dapat diedit di issue view dan dipilih di Create dialog; picker hanya menampilkan kandidat yang sah.
  3. Validasi: parent harus tepat satu level di atas child; nilai tidak sah ditolak dengan galat "Parent harus berasal dari level hierarki di atasnya".
  4. Parent lintas project diizinkan untuk base ke Epic dan level atas; subtask wajib berada di project yang sama dengan parent.
  5. JQL mendukung `parent = KEY` dan `parent in (...)`.
  6. Parent ditampilkan sebagai breadcrumb di issue view.
  7. ⚠ Parent tidak dapat dihapus; Parent tidak dapat di-hide lewat field configuration agar relasi hierarki tetap konsisten.
- **Team-managed vs company-managed:** Tidak ada perbedaan perilaku; team-managed membatasi subtask di bawah Epic (CFG-006).
- **Catatan & ketergantungan:** Disimpan sebagai kolom `parent_id` di work item (modul 28). Alias JQL `"Epic Link"` dan `"Parent Link"` di CFG-024.

### CFG-014 Sprint field

- **Deskripsi:** Locked field yang menyimpan sprint tempat work item berada, termasuk riwayat sprint sebelumnya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem; pengguna dengan Schedule issues/Edit issues untuk mengubah nilai; Admin Situs tidak dapat mengubah konfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Sprint field locked: tidak dapat di-rename, dihapus, atau diubah context-nya.
  2. Field menyimpan riwayat beberapa sprint (work item yang carry-over antar sprint tetap mencatat sprint lama yang sudah ditutup).
  3. Editor hanya menampilkan sprint aktif dan future yang dapat dipilih pada board yang relevan; sprint closed tidak dapat dipilih.
  4. Field tampil di issue view hanya bila ada di screen/layout.
  5. JQL mendukung `sprint in openSprints()`, `closedSprints()`, `futureSprints()`, serta `sprint = <ID atau nama>`.
  6. Subtask mewarisi sprint dari parent dan field Sprint pada subtask tidak dapat diubah terpisah.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed menampilkan Sprint bila Sprints diaktifkan di Features.
- **Catatan & ketergantungan:** Siklus hidup sprint ada di modul 11.

### CFG-015 Rank field

- **Deskripsi:** Locked field untuk urutan global work item berbasis algoritme urutan leksikografis (LexoRank-kompatibel) yang dipakai backlog dan board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem; pengguna dengan Schedule issues mengubah urutan lewat drag-and-drop.
- **Aturan perilaku (acceptance criteria):**
  1. Rank bersifat global lintas project dan disimpan sebagai string terurut leksikografis.
  2. Rank tidak terlihat di issue view; dipakai untuk `ORDER BY Rank` di JQL.
  3. Drag-and-drop ranking di board/backlog hanya aktif bila filter board diurutkan `ORDER BY Rank`; selain itu UI menampilkan peringatan.
  4. Rebalancing rank berjalan otomatis di background tanpa UI admin; halaman manajemen rank tidak dibangun.
  5. Rank locked: tidak dapat dihapus atau diubah context-nya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Algoritme dan kebutuhan konkurensi di modul 27; pemakaian di modul 10 dan 11.

### CFG-016 Story Points & Story point estimate fields

- **Deskripsi:** Field estimasi numerik: "Story Points" (company-managed) dan "Story point estimate" (team-managed).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues mengisi nilai; Admin Situs mengatur context; Admin Proyek mengatur mode Estimation team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed memakai field estimasi bertipe number dengan context yang daftar work type-nya dapat diatur.
  2. Team-managed memakai field "Story point estimate"; mode Estimation (story points vs time) di-toggle lewat Project settings > Features.
  3. Kedua field memiliki ID berbeda sehingga JQL dan laporan harus mempertimbangkan keduanya; produk ini dapat memilih satu field terpadu dengan dua nama tampilan (keputusan terbuka).
  4. Estimation statistic board company-managed dapat berupa Story Points, Original Time Estimate, Issue Count, atau number field lain.
  5. Context field estimasi harus mencakup work type terkait agar field tampil di backlog dan issue view.
  6. Nilai desimal diterima (mis. 0.5) dan disimpan sebagai tipe Number (CFG-026); input non-numerik ditolak dengan galat validasi.
  7. Nilai dipakai di backlog, sprint report, velocity, dan burndown.
- **Team-managed vs company-managed:** Company-managed memakai Story Points dengan context; team-managed memakai Story point estimate dengan toggle Estimation.
- **Catatan & ketergantungan:** Estimation statistic di modul 10/11; laporan di modul 16.

### CFG-017 Time tracking fields

- **Deskripsi:** Field bawaan untuk estimasi dan pencatatan waktu beserta agregat (Σ) yang mencakup subtask.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues/Work on issues; Admin Situs (visibilitas lewat field configuration).
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia Original Estimate (`timeoriginalestimate`), Remaining Estimate (`timeestimate`), dan Time Spent (`timespent`), disimpan dalam detik.
  2. Field agregat Σ Original Estimate, Σ Remaining Estimate, dan Σ Time Spent menjumlahkan nilai work item dan semua subtask-nya.
  3. Field "Time tracking" di screen menampilkan estimasi gabungan; "Log Work" di transition screen memungkinkan worklog saat transisi.
  4. Format input `1w 2d 3h 30m` dikonversi mengikuti setting hours per day dan days per week (CFG-058); format tidak valid ditolak.
  5. Field time tracking dapat disembunyikan lewat field configuration.
  6. Di team-managed, field tersedia bila Estimation diatur ke Time di Project settings > Features.
- **Team-managed vs company-managed:** Company-managed memakai field configuration; team-managed memakai pengaturan Estimation.
- **Catatan & ketergantungan:** Pemilik: modul ini (CFG-017) memiliki definisi field (ID, satuan detik, agregat Σ, visibilitas lewat field configuration). Pencarian JQL atas field ini dimiliki [12 JQL](12-jql-bahasa-kueri.md) JQL-028; TIME-037 di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) adalah rujukan ke JQL-028. Konversi satuan mengikuti TIME-002 (rujukan dari CFG-058); permission Log work dan pengelolaan worklog di PERM-032 ([03 Izin](03-izin-peran-keamanan.md)). Worklog dan perilaku estimasi rinci di modul 14.

### CFG-018 Flagged field

- **Deskripsi:** Field Checkboxes "Flagged" dengan opsi tunggal Impediment yang menandai work item terhambat di board dan backlog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues; Admin Situs (konfigurasi field).
- **Aturan perilaku (acceptance criteria):**
  1. Flagged adalah custom field Checkboxes dengan satu opsi "Impediment" yang di-seed saat instalasi.
  2. Work item yang di-flag diberi highlight dan ikon bendera di board dan backlog.
  3. Aksi Add flag / Remove flag tersedia dengan komentar opsional yang tersimpan sebagai komentar work item.
  4. JQL mendukung `Flagged = Impediment` dan `Flagged IS NOT EMPTY`.
  5. Team-managed memakai mekanisme flag yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tampilan di modul 10/11; komentar di modul 08.

### CFG-019 Team field (Atlassian Teams)

- **Deskripsi:** Locked field "Team" yang terhubung ke entitas Teams lintas produk di produk ini untuk menetapkan tim pada work item (padanan Atlassian Teams).
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues; Admin Situs dan Admin Proyek (menambahkan ke screen/layout).
- **Aturan perilaku (acceptance criteria):**
  1. Team adalah locked field bawaan yang dipakai di Plans, board/filter, dan JQL (`team = <team ID>`).
  2. Picker Team menampilkan tim dari entitas Teams produk ini (modul 18, PLAN-057) yang dibagikan ke seluruh modul; tim yang dihapus atau diarsipkan tidak dapat dipilih lagi tetapi nilainya tetap tampil di work item lama.
  3. Field harus ditambahkan ke screen/layout agar tampil di work item.
  4. Di team-managed, field Team dapat ditambahkan ke layout work type (CFG-042).
  5. Plans memakai Team untuk capacity planning (modul 18).
  6. ⚠ Tim lama khusus-plan (plan-only team) dan shared team dimigrasikan ke entitas Teams terpadu saat fitur Plans dibangun.
- **Team-managed vs company-managed:** Tersedia di keduanya; cara menampilkannya berbeda (screen vs layout).
- **Catatan & ketergantungan:** Bergantung pada entitas Teams (PLAN-057 di [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md)) yang dijadwalkan Fase 4, sehingga field ini ikut Fase 4 meski prioritas P1. Modul ini memegang definisi field (locked, screen/layout, JQL); perilaku picker, satu tim per work item, dan pemakaian capacity dirinci di PLAN-058. Direktori orang/tim di navigasi ada di modul 02 (USR-027).

### CFG-020 Issue color field

- **Deskripsi:** Field warna work item (pengganti Epic Colour) yang dipakai untuk warna bar di Timeline/Plans dan label epic di board/backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Edit issues; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih warna dari palet tetap (mis. 14 warna); nilai di luar palet ditolak API.
  2. Warna dipakai untuk label epic di backlog, card board, dan bar Timeline.
  3. Warna dapat diatur dari issue view atau dari Timeline.
  4. ⚠ Field bersifat locked/dikelola sistem dan tersedia untuk semua work type.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai modul 10, 11, 18. Data Epic Colour legacy dipetakan ke field ini (CFG-024).

### CFG-021 Category field

- **Deskripsi:** Field kategori berwarna per project yang dipakai di project bisnis untuk mengelompokkan work item di List, Board, dan Calendar.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola kategori); pengguna dengan Edit issues (memilih kategori).
- **Aturan perilaku (acceptance criteria):**
  1. Admin Proyek mengelola kategori (nama dan warna) per project; nama kategori unik dalam project.
  2. Satu work item hanya memiliki satu kategori.
  3. Kategori dipakai untuk filter dan grouping di view bisnis (List, Board, Calendar).
  4. ⚠ Ketersediaan di project software ditentukan oleh pemilik produk; default produk ini: tersedia di project bisnis dan dapat diaktifkan di project lain lewat Features.
- **Team-managed vs company-managed:** Tidak ada perbedaan; kategori selalu per project.
- **Catatan & ketergantungan:** View bisnis di modul 26 dan 18. Jangan rancukan dengan Project category (modul 04).

### CFG-022 Target start / Target end & date fields for planning

- **Deskripsi:** Field tanggal yang dipakai Plans/Timeline untuk penjadwalan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs (konfigurasi); pengguna dengan Edit issues/Schedule issues mengisi tanggal.
- **Aturan perilaku (acceptance criteria):**
  1. Target start dan Target end adalah locked field yang dibuat oleh fitur Plans.
  2. Start date dan Due date menjadi default tanggal Timeline bila tidak dikonfigurasi lain.
  3. Plans dapat dikonfigurasi memakai date field kustom sebagai start/end.
  4. Calendar view dapat memakai date field kustom (Date picker) sebagai sumber tanggal.
  5. Field tidak muncul di issue view sampai ditambahkan ke screen/layout.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ketersediaan Target start/end di tier Standard masih diragukan di sumber dan menjadi keputusan terbuka (lihat Pertanyaan terbuka). Pemakaian di modul 18.

### CFG-023 Development & deployment info fields

- **Deskripsi:** Field read-only yang merangkum informasi pengembangan (branch, commit, pull request, build, deployment).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Sistem (diisi integrasi); pengguna dengan View development tools melihat panel.
- **Aturan perilaku (acceptance criteria):**
  1. Field "Development" bersifat locked dan ditampilkan sebagai panel di issue view.
  2. JQL mendukung `development[pullrequests].open`, `development[commits].all`, dan atribut sejenis.
  3. Field tidak dapat diedit manual; nilainya diisi oleh integrasi.
  4. Panel hanya tampil bila ada integrasi pengembangan yang terhubung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail integrasi di modul 24.

### CFG-024 Epic Link, Epic Name, Epic Status, Epic Colour, Parent Link (legacy)

- **Deskripsi:** Field legacy relasi Epic dan Advanced Roadmaps yang sudah digantikan Parent field dan Issue color; tidak dibangun sebagai field, hanya sebagai alias kompatibilitas dan jalur migrasi data.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (impor data).
- **Aturan perilaku (acceptance criteria):**
  1. Epic Link tidak dibangun; data Epic Link saat impor dikonversi ke Parent (CFG-013).
  2. Epic Name tidak dibangun; Summary dipakai sebagai nama Epic, dan nilai Epic Name hasil impor disimpan di Summary bila Summary kosong.
  3. Epic Status tidak dibangun; status selesai Epic ditentukan oleh status category Done.
  4. Epic Colour tidak dibangun; nilainya dipetakan ke Issue color (CFG-020).
  5. Parent Link tidak dibangun; datanya dipetakan ke Parent.
  6. ⚠ JQL menerima `"Epic Link"` dan `"Parent Link"` sebagai alias dari `parent` untuk kompatibilitas filter lama.
  7. Produk ini cukup menyediakan alias dan migrasi data, bukan UI penuh untuk field legacy.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Impor di modul 25; alias JQL di modul 12. Lihat Di luar cakupan.

## Tipe custom field

### CFG-025 Text custom fields (single line, paragraph, read-only)

- **Deskripsi:** Tipe field teks: Short text / Text Field (single line), Paragraph / Text Field multi-line (rich text), dan Text Field (read only) yang hanya diisi lewat API atau otomasi.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (membuat field global); Admin Proyek (team-managed); pengguna dengan Edit issues mengisi nilai.
- **Aturan perilaku (acceptance criteria):**
  1. Short text menerima maksimum 255 karakter tanpa format; input lebih panjang ditolak dengan galat validasi.
  2. Paragraph menyimpan rich text (format dokumen modul 07) dengan editor penuh; batas 32.767 karakter.
  3. Text Field (read only) tidak dapat diedit di UI dan hanya diisi lewat REST API atau otomasi.
  4. Default value dapat diatur per context (CFG-036).
  5. Searcher dapat dipilih antara Free Text Searcher (operator `~` dan `!~`) atau Exact Text Searcher (operator `=`/`!=`).
  6. Renderer diatur di field configuration: Default text renderer atau Wiki style renderer (CFG-043).
  7. Team-managed menyediakan tipe "Short text" dan "Paragraph".
- **Team-managed vs company-managed:** Team-managed hanya Short text dan Paragraph, tanpa read-only dan tanpa pemilihan searcher/renderer.
- **Catatan & ketergantungan:** Indeks pencarian teks di modul 12/27.

### CFG-026 Number custom field

- **Deskripsi:** Field numerik desimal untuk estimasi, nilai bisnis, biaya, dan sejenisnya, dengan opsi format tampilan.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed); pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field menyimpan bilangan desimal presisi ganda; nilai negatif diizinkan; input non-numerik ditolak.
  2. Default value dapat diatur per context.
  3. Searcher Number Searcher / Number range searcher mendukung operator `=`, `!=`, `>`, `>=`, `<`, `<=`.
  4. Field dapat dipilih sebagai estimation statistic board company-managed.
  5. Nilai dipakai untuk sum/roll-up di Plans dan gadget tertentu.
  6. Team-managed menyediakan tipe "Number".
  7. Format tampilan angka (pemisah ribuan dan desimal) mengikuti locale pengguna.
  8. ⚠ Dialog create/edit field Number menyediakan opsi Format: Number (#), Percentage (%), dan Currency ($).
  9. ⚠ Format Currency mengizinkan pemilihan kode/simbol mata uang dari daftar ISO 4217.
  10. Format hanya memengaruhi tampilan; data tetap numerik sehingga JQL, sorting, dan agregasi tidak berubah, dan tidak ada konversi kurs.
  11. ⚠ Percentage menyimpan nilai apa adanya dan menampilkannya dengan sufiks % (25 ditampilkan "25%"); aturan ini dikunci saat implementasi.
  12. ⚠ Tampilan terformat dipakai di issue view, list/navigator, dan board card; export CSV mengeluarkan nilai numerik mentah.
  13. ⚠ Opsi format juga tersedia untuk Number di team-managed.
- **Team-managed vs company-managed:** Team-managed menyediakan tipe Number tanpa pemilihan searcher; ketersediaan opsi format mengikuti butir 13.
- **Catatan & ketergantungan:** Format field number pihak ketiga (field pack Marketplace) di luar cakupan.

### CFG-027 Date picker & Date time picker custom fields

- **Deskripsi:** Field tanggal (tanpa waktu) dan tanggal-waktu dengan date picker di UI.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed); pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Date Picker menyimpan tanggal saja (tanpa zona waktu); format tampilan mengikuti setting global/locale.
  2. Date Time Picker menyimpan tanggal dan jam dalam UTC dan ditampilkan sesuai zona waktu pengguna.
  3. Default value dapat berupa none, current date (Date Picker) / current date and time (Date Time Picker), atau tanggal tetap.
  4. Searcher Date range / Date time range mendukung tanggal relatif JQL (mis. `-1w`) dan fungsi seperti `startOfDay()` dan `endOfWeek()`.
  5. Field dapat dipakai sebagai sumber tanggal di Calendar, Timeline, dan Plans.
  6. Team-managed menyediakan tipe "Date" dan "Time stamp".
- **Team-managed vs company-managed:** Nama tipe berbeda (Date/Time stamp); perilaku sama.
- **Catatan & ketergantungan:** Fungsi tanggal JQL di modul 12.

### CFG-028 Select list, radio, and checkbox custom fields

- **Deskripsi:** Field pilihan Select List (single choice), Select List (multiple choices), Radio Buttons, dan Checkboxes dengan opsi yang dikelola per context.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed); pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Select List (single choice) menampilkan dropdown satu nilai.
  2. Select List (multiple choices) menerima beberapa nilai.
  3. Radio Buttons menerima satu nilai dan ditampilkan sebagai radio.
  4. Checkboxes menerima beberapa nilai dan ditampilkan sebagai checkbox.
  5. Opsi didefinisikan per context; context berbeda dapat memiliki daftar opsi berbeda.
  6. Default value berupa satu opsi (single/radio) atau beberapa opsi (multi/checkbox).
  7. Searcher Multi Select Searcher mendukung JQL `=`, `in`, `not in`, dan `IS EMPTY`.
  8. Opsi yang dinonaktifkan tetap tersimpan dan tampil pada work item lama tetapi tidak dapat dipilih lagi.
  9. Team-managed menyediakan "Dropdown" (single select) dan "Checkbox" (multi) dengan opsi per project.
  10. ⚠ Jumlah opsi per context dibatasi oleh batas sistem yang dapat dikonfigurasi; melebihi batas ditolak dengan pesan galat.
- **Team-managed vs company-managed:** Team-managed hanya Dropdown dan Checkbox, opsi per project tanpa context.
- **Catatan & ketergantungan:** Pengelolaan opsi di CFG-037.

### CFG-029 Select List (cascading) custom field

- **Deskripsi:** Field dropdown bertingkat dua level: parent option dan child option.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field memiliki tepat dua level (parent dan child); level ketiga tidak didukung.
  2. Admin mengelola parent option lalu child options per parent.
  3. Default value berupa parent dan (opsional) child.
  4. Child boleh kosong (hanya parent yang dipilih).
  5. Searcher Cascading Select mendukung fungsi JQL `cascadeOption(parent, child)` dan `cascadeOption(parent, none)`.
  6. Opsi dapat diurutkan ulang, dinonaktifkan, dan dihapus per level.
  7. ⚠ Tipe ini tidak tersedia di field library team-managed.
- **Team-managed vs company-managed:** Hanya company-managed (lihat butir 7).
- **Catatan & ketergantungan:** CFG-037 butir 9.

### CFG-030 User picker & Group picker custom fields

- **Deskripsi:** Field untuk memilih pengguna atau grup, single maupun multiple, termasuk pemanfaatannya sebagai sumber permission dan notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed); pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia User Picker (single user) dan User Picker (multiple users).
  2. Tersedia Group Picker (single group) dan Group Picker (multiple groups).
  3. User filter per context membatasi pilihan pengguna ke grup dan/atau project role tertentu; pengguna di luar filter tidak muncul di picker dan ditolak API.
  4. ⚠ Default value user/group dapat diatur per context untuk tipe yang mendukungnya (minimal User Picker single dan Group Picker single).
  5. User picker dapat dipakai sebagai "User custom field value" di permission scheme, notification scheme, dan issue security level.
  6. Group picker dapat dipakai sebagai "Group custom field value" di permission/notification scheme dan issue security level.
  7. JQL mendukung `currentUser()`, `membersOf()`, dan `EMPTY` untuk field ini.
  8. Team-managed menyediakan tipe "People" (single atau multiple).
  9. Tampilan nama/avatar mengikuti pengaturan privasi profil pengguna (modul 02).
  10. Produk ini tidak memiliki field-level permission native: field terlihat oleh semua pengguna yang dapat melihat work item.
- **Team-managed vs company-managed:** Team-managed hanya tipe People tanpa user filter dan tanpa Group picker.
- **Catatan & ketergantungan:** Pemakaian di scheme pada modul 03 dan 19. Field-level security adalah domain app Marketplace (di luar cakupan).

### CFG-031 URL field & Labels custom field

- **Deskripsi:** Field URL tervalidasi dan field label kustom yang terpisah dari system field Labels.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed); pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. URL Field memvalidasi format URL (skema http/https/mailto) dan ditampilkan sebagai link yang dapat diklik; URL tidak valid ditolak.
  2. Default value URL dapat diatur per context.
  3. Labels (custom) menyimpan kumpulan label bebas yang terpisah dari system Labels, dengan autosuggest dari nilai yang sudah ada di field tersebut.
  4. Label tidak boleh mengandung spasi dan maksimum 255 karakter per label; pelanggaran ditolak dengan galat validasi.
  5. Searcher Label mendukung JQL `=`, `in`, `not in`, dan `IS EMPTY`.
  6. Team-managed menyediakan tipe "Labels".
- **Team-managed vs company-managed:** Team-managed hanya tipe Labels; URL field hanya company-managed.
- **Catatan & ketergantungan:** System Labels di CFG-011.

### CFG-032 Project picker & Version picker custom fields

- **Deskripsi:** Field untuk memilih project atau version dari project work item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; pengguna dengan Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Project Picker (single project) hanya menampilkan project yang dapat dilihat pengguna.
  2. Tersedia Version Picker (single version) dan Version Picker (multiple versions) yang menampilkan versi dari project work item.
  3. Version picker memisahkan versi released dan unreleased dalam grup terpisah.
  4. Searcher Project/Version mendukung fungsi JQL `releasedVersions()` dan `unreleasedVersions()`.
  5. Nilai versi ikut berubah saat versi di-rename atau di-merge karena disimpan berbasis ID.
  6. Tipe ini tidak tersedia di field library team-managed.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Versi di modul 15.

### CFG-033 Forge/Connect custom field types

- **Deskripsi:** Extension point platform yang memungkinkan ekstensi mendefinisikan custom field atau tipe custom field baru (padanan Forge `jira:customField`/`jira:customFieldType`).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pembuat ekstensi); Admin Situs (memasang ekstensi dan membuat field).
- **Aturan perilaku (acceptance criteria):**
  1. Ekstensi dapat mendaftarkan custom field (`customField`) atau custom field type (`customFieldType`) lewat manifest.
  2. Tipe field dari ekstensi muncul di wizard Create custom field dengan label nama ekstensi.
  3. Nilai disimpan oleh produk ini; rendering dan validasi dilakukan oleh ekstensi.
  4. Field yang didefinisikan ekstensi sebagai `customField` bersifat locked (CFG-012).
  5. Field kalkulasi/read-only (computed) yang tidak dapat dinyatakan dengan tipe Formula custom field native (CFG-062), mis. kalkulasi yang memerlukan data eksternal, dapat dibuat lewat ekstensi sebagai pengganti tipe field Jira Toolkit DC.
  6. ⚠ Saat ekstensi di-uninstall, field disembunyikan dan tidak berfungsi, tetapi nilainya tetap disimpan dan kembali aktif bila ekstensi dipasang ulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Runtime ekstensi di modul 25. Kalkulasi berbasis field produk ini dilayani secara native oleh CFG-062; ekstensi hanya untuk kebutuhan di luarnya. Tipe field Jira Toolkit (Date of First Response, Participants of an issue, Message custom field) bukan core dan di luar cakupan.

## Administrasi custom field (company-managed)

### CFG-034 Custom fields administration page

- **Deskripsi:** Halaman admin global (Settings > Work items > Fields) untuk melihat, mencari, membuat, mengedit, dan menghapus custom field company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar field menampilkan nama, tipe, jumlah context, jumlah screen, Last value update, dan deskripsi.
  2. Daftar dapat difilter berdasarkan tipe field, project, dan screen, termasuk filter "tidak ada di screen mana pun".
  3. Daftar dapat dicari berdasarkan nama.
  4. Create custom field berupa wizard: pilih tipe, isi nama dan deskripsi, (opsional) isi opsi, lalu asosiasikan ke screen.
  5. Tipe field tidak dapat diubah setelah field dibuat.
  6. Edit details mengubah nama, deskripsi, dan search template (searcher, CFG-038).
  7. Menu aksi per field: Edit details, Contexts and default value, Screens, Move to trash.
  8. Field yang dihapus masuk trash selama 60 hari dan dapat di-restore (beserta context dan nilai) atau dihapus permanen; setelah 60 hari dihapus permanen otomatis.
  9. Tab/filter "Trashed" menampilkan field di trash beserta tanggal penghapusan permanen otomatis.
  10. ⚠ Field di trash tidak tampil di screen dan create/edit, dan tidak dapat dipakai di JQL; filter/board yang mereferensikannya menampilkan galat "Field tidak tersedia" sampai field di-restore.
  11. Move to trash, restore, dan delete permanen tercatat di audit log (modul 22).
  12. System field dan locked field tidak dapat dipindahkan ke trash.
  13. Locked field tidak dapat diedit atau dihapus dan ditandai ikon gembok.
  14. Nama field boleh duplikat; JQL memakai `cf[ID]` untuk disambiguasi dan autocomplete menampilkan ID saat nama ganda.
  15. Halaman tidak menampilkan field team-managed (cakupan per project).
  16. Seluruh aksi memerlukan Administer Jira; pengguna lain menerima HTTP 403.
- **Team-managed vs company-managed:** Hanya company-managed; field team-managed dikelola di CFG-041.
- **Catatan & ketergantungan:** ID field (mis. `customfield_10010`) stabil untuk API (modul 23). Tabel `custom_fields` dengan kolom `deleted_at` untuk trash.

### CFG-035 Custom field contexts

- **Deskripsi:** Context menentukan project dan work type tempat custom field berlaku, beserta opsi dan default value per context.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Context dapat berupa global context (semua project) atau project-specific context.
  2. Context memilih work type yang berlaku (semua atau subset).
  3. Satu project hanya boleh berada di satu context per field; menambahkan project yang sudah ada di context lain ditolak dengan galat.
  4. Add new context meminta nama, deskripsi, work type, dan project.
  5. Context dapat diedit dan dihapus; field tanpa context tidak berlaku di mana pun (dengan peringatan saat menghapus context terakhir).
  6. Field yang tidak berlaku untuk project/work type tidak tampil meskipun ada di screen.
  7. Field select dapat memiliki opsi dan default value berbeda per context.
  8. Context dapat dicari dan daftar project terasosiasi ditampilkan.
  9. ⚠ Jumlah context per field dibatasi oleh batas sistem yang dapat dikonfigurasi.
- **Team-managed vs company-managed:** Hanya company-managed; field team-managed selalu berlaku untuk project-nya saja.
- **Catatan & ketergantungan:** Pemilik domain context adalah CFG-035; endpoint REST context dan default value dimiliki API-042 di [23 API Publik & Webhook](23-api-publik-webhook.md), yang merujuk balik ke modul ini. Resolusi context dipakai saat render form (modul 06) dan evaluasi JQL (modul 12). Scope Formula custom field memakai context yang sama (CFG-062).

### CFG-036 Custom field default values

- **Deskripsi:** Nilai default yang diisi otomatis saat work item dibuat, diatur per context.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Default dapat diatur per context untuk tipe text, paragraph, number, date, date time, select, multi-select, cascading, URL, dan user/group yang mendukung.
  2. Default Date dapat berupa current date atau tanggal tetap.
  3. Default diisikan di create screen dan dapat diubah pengguna sebelum menyimpan.
  4. ⚠ Default diterapkan saat create lewat API bila field tidak dikirim dalam payload; field yang dikirim eksplisit bernilai null tetap kosong.
  5. Default tidak diterapkan secara retroaktif ke work item lama.
  6. Admin dapat menghapus (clear) default value.
- **Team-managed vs company-managed:** Team-managed mengatur default per work type lewat layout (CFG-042).
- **Catatan & ketergantungan:** Terkait CFG-035 dan CFG-037.

### CFG-037 Custom field options management

- **Deskripsi:** Pengelolaan opsi untuk field select, radio, checkbox, dan cascading.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat menambah opsi satu per satu atau beberapa sekaligus (satu per baris).
  2. Rename opsi membuat nilai di work item ikut berubah karena disimpan berbasis ID opsi.
  3. Opsi dapat diurutkan ulang lewat drag atau sort alfabetis.
  4. Opsi yang dinonaktifkan tidak dapat dipilih lagi tetapi tetap tampil di work item lama.
  5. Menghapus opsi menghapus nilainya dari work item yang memakainya setelah peringatan yang menyebut jumlah work item terdampak.
  6. Opsi dapat diatur sebagai default.
  7. Opsi berlaku per context dan tidak tersalin otomatis ke context lain.
  8. Nama opsi duplikat dalam context yang sama ditolak (case-insensitive).
  9. Untuk cascading, admin mengelola child options per parent option.
  10. Di team-managed, opsi Dropdown/Checkbox dikelola per project.
- **Team-managed vs company-managed:** Company-managed per context; team-managed per project.
- **Catatan & ketergantungan:** Pemilik domain opsi adalah CFG-037; endpoint REST opsi (termasuk cascading) dimiliki API-042 di [23 API Publik & Webhook](23-api-publik-webhook.md). ID opsi stabil untuk API dan JQL.

### CFG-038 Custom field searchers & indexing

- **Deskripsi:** Search template yang menentukan bagaimana field diindeks dan dicari di JQL dan basic search.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Searcher dipilih di Edit details (mis. Free Text Searcher, Exact Text Searcher, Number range, Multi Select) dari daftar yang kompatibel dengan tipe field.
  2. Searcher "None" membuat field tidak dapat dicari dan tidak muncul di autocomplete JQL.
  3. Perubahan searcher diikuti reindex otomatis di background tanpa langkah manual admin.
  4. Searcher menentukan operator JQL yang tersedia untuk field.
  5. Searcher menentukan apakah field tampil di filter basic search.
- **Team-managed vs company-managed:** Team-managed memakai searcher default per tipe tanpa opsi pemilihan.
- **Catatan & ketergantungan:** Mesin indeks di modul 12/27. Sebelum Fase 3 setiap tipe memakai searcher default.

### CFG-039 Field usage & field guardrails (Cloud)

- **Deskripsi:** Informasi pemakaian custom field dan batas jumlah field untuk menjaga kinerja.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Kolom "Last value update" di daftar Fields membantu menemukan field yang tidak terpakai.
  2. Filter field yang tidak ada di screen mana pun tersedia.
  3. Field tidak terpakai dapat dipindahkan ke trash langsung dari hasil filter.
  4. ⚠ Sistem membatasi jumlah custom field per site dan jumlah field per kombinasi project + work type dengan angka yang dapat dikonfigurasi, menampilkan peringatan saat mencapai 80% batas, dan menolak pembuatan saat batas tercapai.
- **Team-managed vs company-managed:** Batas juga berlaku untuk field team-managed (dihitung per project).
- **Catatan & ketergantungan:** Angka batas ditetapkan bersama NFR (modul 27).

### CFG-040 Custom field optimizer

- **Deskripsi:** Alat Data Center yang memindai custom field global berkonteks lebar dan merekomendasikan pembatasan context; tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: pemindaian field global yang hanya dipakai sedikit project; padanan parsial adalah data Last value update (CFG-039).
  2. Tidak dibangun: rekomendasi dan penerapan otomatis pembatasan context.
  3. Tidak dibangun: laporan jumlah work item/project per field.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat Di luar cakupan.

## Field di team-managed project

### CFG-041 Team-managed field library & field types

- **Deskripsi:** Field yang dibuat per project di team-managed project dengan himpunan tipe terbatas.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (Administer projects).
- **Aturan perilaku (acceptance criteria):**
  1. Tipe yang tersedia: Short text, Paragraph, Date, Time stamp, Number, Dropdown, Checkbox, Labels, dan People.
  2. Field bawaan yang tersedia: Start date, Due date, Sprint, Team, Story point estimate, Original estimate/Time tracking, Labels, dan Components (bila diaktifkan).
  3. Field yang dibuat hanya berlaku di project tersebut; field bernama sama di project lain adalah field berbeda.
  4. Satu field dapat dipakai ulang di beberapa work type dalam project yang sama.
  5. ⚠ Jumlah field per work type/project dibatasi oleh batas sistem yang dapat dikonfigurasi.
  6. Menghapus field dari project menghapus nilainya di semua work item setelah konfirmasi eksplisit (mengetik nama field).
  7. Field dapat di-rename dan deskripsinya (help text) diedit.
  8. Untuk Dropdown/Checkbox, Admin Proyek mengelola opsi, urutan, dan default.
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah CFG-034 s.d. CFG-037.
- **Catatan & ketergantungan:** Disimpan di tabel custom field yang sama dengan kolom `project_id` (scope project).

### CFG-042 Team-managed field configuration per work type

- **Deskripsi:** Editor drag-and-drop di Project settings > Work types > [tipe] untuk menyusun field per work type di team-managed project.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Panel kanan berisi field library project dan tipe field baru yang dapat ditarik ke layout.
  2. Layout memiliki section "Description fields" (area utama), "Context fields" (sidebar), dan "Hide when empty".
  3. ⚠ Field dapat ditandai Required dan diberi default value untuk tipe yang mendukung (minimal teks, angka, tanggal, dropdown).
  4. Field dapat diurutkan ulang dengan drag dan dipindah antar section.
  5. Remove field dari work type tidak menghapus field dari library project.
  6. Perubahan hanya berlaku setelah Save changes; Discard membatalkan semua perubahan sejak penyimpanan terakhir.
  7. Satu layout berlaku untuk create, view, dan edit (tidak ada screen terpisah).
  8. Posisi Summary dan Description terkunci.
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah CFG-043, CFG-046, dan CFG-049.
- **Catatan & ketergantungan:** Fase 1 memakai layout default per template tanpa editor. Layout disimpan sebagai dokumen JSON per work type.

## Field configuration, screen & layout (company-managed)

### CFG-043 Field configurations

- **Deskripsi:** Pengaturan perilaku field (hide, required, deskripsi, renderer) untuk company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia Default Field Configuration bawaan; admin dapat membuat, menyalin, mengedit, dan menghapus field configuration baru.
  2. Per field, admin dapat mengedit deskripsi (help text yang tampil di bawah field).
  3. Hide/Show: field yang di-hide tidak tampil di semua screen untuk work type yang memakai konfigurasi ini.
  4. Required/Optional: field wajib divalidasi saat create/edit di UI maupun REST API dengan galat "<Nama field> wajib diisi".
  5. Renderer dapat dipilih: Default Text Renderer atau Wiki Style Renderer untuk text/paragraph; Autocomplete atau Select list renderer untuk Components/Versions.
  6. Dari dalam field configuration admin dapat melihat dan mengatur screen yang memuat field.
  7. Summary dan Issue Type tidak dapat di-hide atau dibuat opsional.
  8. Field required yang tidak ada di create screen membuat create gagal dengan galat yang menyebut nama field.
  9. Daftar item di dalam field configuration dapat dicari.
  10. Field configuration yang dipakai field configuration scheme tidak dapat dihapus.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed mengatur Required di layout (CFG-042).
- **Catatan & ketergantungan:** Validasi required dijalankan di use case create/edit (modul 06) dan transisi (modul 09).

### CFG-044 Field configuration schemes

- **Deskripsi:** Memetakan field configuration ke work type dan diasosiasikan ke company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Add field configuration scheme meminta nama dan deskripsi.
  2. Admin memetakan work type ke field configuration; entri "Default" berlaku untuk work type tanpa mapping.
  3. Mapping dapat diedit/dihapus dan scheme dapat disalin.
  4. Scheme diasosiasikan ke project lewat Project settings > Fields > Use a different scheme.
  5. Project tanpa scheme memakai System Default Field Configuration.
  6. Scheme hanya dapat dihapus bila tidak dipakai project mana pun.
  7. Daftar scheme menampilkan indikator project yang memakai scheme.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Dapat digantikan Field schemes (CFG-045).

### CFG-045 Field schemes (simplified field configuration)

- **Deskripsi:** Model konfigurasi yang lebih sederhana yang menyatukan field configuration dan field configuration scheme menjadi satu scheme untuk mengatur field, required, dan deskripsi per work type.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Satu field scheme mendefinisikan field yang tersedia per work type beserta status required dan deskripsinya.
  2. Tersedia migrasi otomatis dari field configuration scheme lama ke field scheme dengan pratinjau hasil.
  3. Kedua model dapat berdampingan selama masa transisi; satu project hanya memakai salah satu model pada satu waktu.
  4. Field scheme diasosiasikan ke company-managed project.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Keputusan terbuka: apakah produk ini langsung memakai model field scheme saja (lebih sederhana) dan menghilangkan field configuration klasik, atau mendukung keduanya demi kompatibilitas impor.

### CFG-046 Screens

- **Deskripsi:** Screen mendefinisikan field yang tampil dan urutannya saat create/edit/view work item dan saat transisi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek dengan extended project administration untuk screen yang hanya dipakai project-nya (CFG-050).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar screen menampilkan nama, deskripsi, serta screen scheme dan workflow yang memakai.
  2. Screen dapat ditambah, diubah nama/deskripsinya, disalin, dan dihapus (hanya bila tidak dipakai).
  3. Configure screen: tambah field lewat pencarian, hapus field, dan drag untuk reorder.
  4. Screen mendukung tab: tambah, rename, reorder, hapus, serta pindahkan field antar tab.
  5. Screen bawaan: Default Screen, Resolve Issue Screen, Workflow Screen, ditambah screen per template project.
  6. Transition screen dipasang di workflow transition (detail di modul 09).
  7. Field yang di-hide di field configuration atau tanpa context tidak tampil walau ada di screen.
  8. Screen tidak menentukan required; required diatur field configuration.
  9. Editor screen menampilkan peringatan bila field required tidak ada di create screen.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai satu layout (CFG-042).
- **Catatan & ketergantungan:** Fase 1 memakai screen default hasil seed.

### CFG-047 Screen schemes

- **Deskripsi:** Memetakan operasi work item (Create, Edit, View) ke screen.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia Default Screen Scheme bawaan.
  2. Add screen scheme meminta nama, deskripsi, dan default screen.
  3. Operasi Create issue, Edit issue, dan View issue dapat dipetakan ke screen berbeda.
  4. Operasi tanpa mapping memakai default screen.
  5. Screen scheme dapat diedit, disalin, dan dihapus (hanya bila tidak dipakai).
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Dipakai oleh CFG-048.

### CFG-048 Issue type screen schemes

- **Deskripsi:** Memetakan work type ke screen scheme dan diasosiasikan ke company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia Default Issue Type Screen Scheme bawaan.
  2. Add meminta nama, deskripsi, dan default screen scheme.
  3. Admin memetakan work type ke screen scheme; entri Default berlaku untuk work type lain.
  4. Mapping dapat diedit/dihapus; scheme dapat disalin dan dihapus (hanya bila tidak dipakai).
  5. Scheme diasosiasikan ke project lewat Project settings > Screens > Use a different scheme.
  6. Satu project memiliki tepat satu issue type screen scheme.
  7. Daftar menampilkan indikator project yang memakai scheme.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Resolusi screen: project → issue type screen scheme → screen scheme → screen per operasi.

### CFG-049 Issue layout (company-managed)

- **Alias inventaris:** Issue Layout Configuration
- **Deskripsi:** Konfigurasi tata letak issue view per work type dengan section Description fields, Context fields, dan Hide when empty; di company-managed berbasis screen, di team-managed lewat editor work type.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek dan Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: Project settings > Issue layout diatur per issue type screen scheme; field di-drag ke Description fields, Context fields, atau Hide when empty, dan field harus ada di view screen.
  2. Company-managed: perubahan berlaku untuk semua project yang memakai screen yang sama, dengan peringatan "konfigurasi dipakai bersama".
  3. Team-managed: Project settings > Work types > [tipe] untuk drag field ke Description/Context/Hide when empty dan membuat field baru langsung (CFG-042).
  4. Team-managed: field dapat ditandai required dan diberi default value per work type.
  5. Tombol Configure di issue view memberi akses cepat ke layout bagi pengguna yang berwenang.
  6. Posisi field sistem tertentu (Summary, Status, Description) tetap dan tidak dapat dipindahkan.
  7. Layout diakses dari Project settings > Issue layout (pilih work type) atau dari Configure di issue view.
  8. Section yang tersedia: Description fields, Context fields, Hide when empty.
  9. Context fields dapat dikelompokkan ke grup collapsible (mis. "More fields").
  10. ⚠ Field berasal dari screen View/Edit; menambah field dari layout ikut mengubah screen, dan bila screen dipakai bersama project lain penambahan hanya diizinkan untuk Admin Situs.
  11. Field dapat di-drag untuk reorder dan dipindah antar section.
  12. Pinned fields pribadi pengguna terpisah dari layout admin dan tidak memengaruhi pengguna lain.
  13. Layout berlaku per work type per project, dengan peringatan bila screen dipakai project lain.
  14. Perubahan disimpan lewat Save atau dibatalkan lewat Discard.
- **Team-managed vs company-managed:** Company-managed berbasis screen dan dapat dipakai bersama; team-managed berbasis layout per work type yang juga mengatur required/default.
- **Catatan & ketergantungan:** Pemilik: CFG-049 adalah versi lengkap konfigurasi Issue layout; PRJ-048 Issue Layout (company-managed) di [04 Proyek](04-proyek.md) adalah rujukan titik masuk dari Project settings dan tidak mendefinisikan aturan sendiri. Butir 1-6 dari domain issues-core, 7-14 dari admin-issue-config. Render issue view di modul 06. Hasil Formula custom field (CFG-062) ditata seperti field lain di layout ini.

### CFG-050 Project-level work type configuration view (company-managed)

- **Deskripsi:** Halaman Project settings untuk melihat dan mengubah konfigurasi work type, field, dan screen pada company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (melihat, extended project administration); Admin Situs (mengganti/mengedit scheme).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Work types menampilkan work type dari scheme beserta workflow, field configuration, dan screen per work type.
  2. Per work type, admin melihat field (required/hidden) dan screen terkait.
  3. Project settings > Fields menampilkan field configuration scheme yang dipakai.
  4. Project settings > Screens menampilkan issue type screen scheme yang dipakai.
  5. Aksi "Use a different scheme" dan "Edit scheme" hanya tersedia untuk Admin Situs.
  6. Extended project administration: Admin Proyek dapat menambah, menghapus, dan mengurutkan field di screen yang hanya dipakai project-nya, tanpa membuat custom field baru.
  7. Admin Proyek company-managed tidak dapat membuat custom field atau mengubah scheme (HTTP 403).
  8. Halaman menampilkan indikator "dipakai bersama dengan N project lain".
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai CFG-002/CFG-042.
- **Catatan & ketergantungan:** Halaman Project settings di modul 04; permission di modul 03.

### CFG-051 Find your field (field visibility helper)

- **Deskripsi:** Alat bantu di issue view yang menjelaskan mengapa sebuah field tidak terlihat pada work item tertentu.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs dan Admin Proyek (dengan tautan perbaikan); pengguna lain (penjelasan saja).
- **Aturan perilaku (acceptance criteria):**
  1. Alat diakses dari issue view lewat menu Configure > Find your field.
  2. Pengguna memilih field, lalu sistem menjelaskan penyebab: tidak ada di screen/layout, di-hide di field configuration, tidak ada context untuk project/work type, atau kosong di section Hide when empty.
  3. Bagi admin tersedia tautan langsung ke konfigurasi terkait untuk memperbaiki.
  4. ⚠ Non-admin hanya melihat penjelasan tanpa tautan konfigurasi.
- **Team-managed vs company-managed:** Di team-managed penyebab yang mungkin hanya: tidak ada di layout atau kosong di Hide when empty.
- **Catatan & ketergantungan:** Membutuhkan fungsi resolusi konfigurasi yang sama dengan render issue view.

### CFG-052 Description templates per work type

- **Deskripsi:** Template konten default untuk field Description per work type.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed); Admin Situs (company-managed).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Team-managed: default value Description diatur per work type di layout editor.
  2. ⚠ Company-managed: produk ini menyediakan default Description per work type di layout/field configuration sebagai fitur native (bukan hanya lewat otomasi).
  3. Template diterapkan saat create hanya bila Description kosong.
  4. Template mendukung rich text (heading, checklist, tabel).
- **Team-managed vs company-managed:** Lokasi pengaturan berbeda (layout vs konfigurasi company-managed).
- **Catatan & ketergantungan:** Butir 2 adalah keputusan produk: sumber menyebut company-managed tidak memiliki default Description native (biasanya dibuat lewat otomasi), sedangkan produk ini memilih menyediakannya secara native agar setara dengan team-managed (lihat Pertanyaan terbuka). Alternatif otomasi di modul 20; editor di modul 07.

## Nilai global: priority, resolution, link & fitur work item

### CFG-053 Priorities

- **Deskripsi:** Pengelolaan nilai priority global: nama, deskripsi, ikon, warna, dan urutan.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Priority bawaan di-seed: Highest, High, Medium, Low, Lowest (default Medium).
  2. Add priority meminta nama (wajib, unik case-insensitive), deskripsi, ikon dari pustaka, dan warna (hex `#RRGGBB`); format warna tidak valid ditolak.
  3. Priority dapat diedit.
  4. Delete priority mewajibkan pemilihan priority pengganti untuk work item yang memakai; priority terakhir tidak dapat dihapus.
  5. Priority dapat diurutkan ulang; urutan menentukan hasil `ORDER BY priority`.
  6. ID priority dipakai JQL dan API dan tidak berubah saat rename.
  7. Team-managed memakai Default priority scheme.
- **Team-managed vs company-managed:** Team-managed selalu memakai Default priority scheme; company-managed dapat memakai priority scheme (CFG-054).
- **Catatan & ketergantungan:** Fase 1 cukup seed + daftar read-only; UI tambah/ubah dapat menyusul di Fase 2 tanpa perubahan skema.

### CFG-054 Priority schemes

- **Alias inventaris:** Priority Schemes
- **Deskripsi:** Kumpulan priority yang tersedia untuk company-managed project tertentu beserta default priority, sehingga setiap project dapat memiliki subset priority dan default sendiri.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Admin membuat priority scheme dengan memilih priority yang tersedia dan default-nya.
  2. Priority scheme diasosiasikan ke project.
  3. Saat project berganti scheme dan ada nilai priority yang tidak tersedia, admin wajib memetakan nilai lama ke priority di scheme baru.
  4. Default priority scheme dipakai project yang belum diatur.
  5. Priority scheme hanya berlaku untuk company-managed project.
  6. Default priority scheme berisi semua priority dan tidak dapat dihapus.
  7. Add priority scheme meminta nama, deskripsi, subset priority, dan default priority (default wajib termasuk subset).
  8. Satu scheme dapat diasosiasikan ke satu atau banyak company-managed project.
  9. Mengganti scheme project menjalankan wizard mapping untuk priority yang tidak tersedia, diproses di background untuk jumlah besar.
  10. Scheme dapat diedit, disalin, dan dihapus (project terkait kembali ke default scheme).
  11. Default priority per scheme dipakai saat create bila pengguna tidak memilih priority.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai default scheme.
- **Catatan & ketergantungan:** Dua entri inventaris yang namanya hanya berbeda kapitalisasi ("Priority Schemes" dari issues-core: butir 1-5; "Priority schemes" dari admin-issue-config: butir 6-11) digabung di heading ini.

### CFG-055 Resolutions

- **Deskripsi:** Pengelolaan nilai resolution global yang menandai work item selesai.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Resolution bawaan di-seed: Done, Won't Do, Duplicate, Cannot Reproduce (bervariasi per template).
  2. Add resolution meminta nama (unik) dan deskripsi.
  3. Resolution dapat diedit, dihapus (dengan memilih pengganti), dan diurutkan ulang.
  4. Admin dapat menetapkan default resolution.
  5. Work item dianggap resolved bila resolution terisi, terlepas dari statusnya; field Resolved terisi saat resolution diisi dan dikosongkan saat resolution dihapus.
  6. ⚠ Resolution di-set lewat transition screen atau post function; di team-managed resolution diisi otomatis (Done) saat work item masuk status berkategori Done dan dikosongkan saat keluar.
  7. Pembatasan resolution per status didukung lewat workflow property `jira.field.resolution.include`/`exclude` (WF-009 butir 11, modul 09).
  8. JQL mendukung `resolution = Unresolved` dan `resolution IS EMPTY`.
- **Team-managed vs company-managed:** Company-managed mengisi lewat transition screen/post function; team-managed otomatis mengikuti status category.
- **Catatan & ketergantungan:** Pemilik: CFG-055 memiliki entitas resolution, nilai seed, default resolution, dan halaman admin daftar nilai (Settings > Work items > Resolutions, sejajar dengan Priorities); WF-010 Resolutions administration di [09 Workflow](09-workflow-status-transisi.md) diharapkan menjadi rujukan ke sini. Semantik runtime (pengisian/pengosongan saat transisi, reopen, include/exclude per status) dimiliki WF-009; butir 5-7 di sini merangkum aturan tersebut untuk konsistensi. Seed resolution Done diperlukan sejak Fase 1 oleh workflow default (modul 09) walau UI admin di Fase 2.

### CFG-056 Issue linking settings & link types

- **Alias inventaris:** Issue Link Types Administration
- **Deskripsi:** Aktivasi issue linking global dan pengelolaan link type dengan deskripsi outward/inward yang berlaku untuk semua project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (pengaturan); pengguna dengan Link issues (membuat link, modul 06).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle global Issue linking ON/OFF tersedia di Settings > Work items > Work item linking.
  2. Link type bawaan: Blocks, Cloners, Duplicate, Relates; modul lain dapat menambahkan link type sendiri (mis. "causes" dari Service Management dan link Product Discovery).
  3. Tambah link type meminta nama, deskripsi outward, dan deskripsi inward.
  4. Nama dan deskripsi link type dapat diedit.
  5. Menghapus link type yang masih dipakai mewajibkan pilihan: pindahkan link ke link type pengganti atau hapus semua link-nya.
  6. Link type berlaku global untuk semua project.
  7. Link type "Cloners" dipakai otomatis oleh fitur Clone (modul 06).
  8. Toggle global aktif/nonaktif: saat nonaktif, bagian Linked Issues dan aksi Link disembunyikan dan API create link ditolak; link yang sudah ada tetap tersimpan.
  9. Deskripsi bawaan: Blocks (blocks / is blocked by), Cloners (clones / is cloned by), Duplicate (duplicates / is duplicated by), Relates (relates to / relates to).
  10. Add link type memvalidasi nama unik dan kedua deskripsi wajib diisi.
  11. Rename link type tidak mengubah link yang sudah ada (berbasis ID).
  12. Dialog delete menampilkan jumlah link terdampak sebelum admin memilih swap ke tipe lain atau hapus semua link.
  13. Link type tambahan dari produk lain (mis. Problem/Incident dan causes dari Service Management; "is implemented by" dari Product Discovery) didaftarkan saat modul tersebut aktif dan ditandai sebagai milik modulnya.
  14. Link dapat lintas project; remote link (web link, halaman wiki eksternal) dikelola terpisah di modul 06.
  15. JQL mendukung `issueLinkType` dan fungsi `linkedIssues()`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 1-7 dari issues-core, 8-15 dari admin-issue-config. Fase 1 cukup seed link type bawaan + toggle; UI CRUD link type dapat menyusul di Fase 2.

### CFG-057 Issue features (voting, watching, unassigned)

- **Alias inventaris:** Global Work Item Settings
- **Deskripsi:** Toggle tingkat site di Settings > System > General configuration dan pengaturan terkait yang memengaruhi perilaku work item di semua project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Voting on/off: bila off, field Votes dan tombol vote tidak tampil dan API vote ditolak.
  2. Watching on/off: bila off, watcher tidak tersedia dan notifikasi berbasis watcher tidak dikirim.
  3. Allow unassigned issues: bila off, setiap work item wajib memiliki assignee dan default assignee project tidak boleh Unassigned.
  4. Sub-tasks enable/disable mengikuti CFG-005.
  5. Issue linking enable/disable mengikuti CFG-056.
  6. Permission terkait (View voters and watchers, Manage watchers) diatur di permission scheme (modul 03).
  7. Perubahan berlaku global di semua project.
  8. Pengaturan Allow unassigned work items tersedia di General configuration.
  9. Toggle Voting ON/OFF tersedia di halaman yang sama.
  10. Toggle Watching ON/OFF tersedia di halaman yang sama.
  11. Toggle Sub-tasks enable/disable tersedia di Settings > Work items > Sub-tasks.
  12. Toggle Work item linking ON/OFF tersedia di Settings > Work items > Work item linking.
  13. Toggle Time tracking enable/disable tersedia (detail di CFG-058 dan modul 14).
  14. Perubahan setting langsung memengaruhi UI (tombol disembunyikan tanpa reload penuh) dan validasi API (galat "Fitur dinonaktifkan oleh administrator").
- **Team-managed vs company-managed:** Tidak ada perbedaan; berlaku untuk semua project.
- **Catatan & ketergantungan:** Butir 1-7 dari admin-issue-config, 8-14 dari issues-core. Disimpan sebagai setting site (modul 22).

### CFG-058 Time tracking settings

- **Deskripsi:** Konfigurasi global provider time tracking dan format waktu.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Provider time tracking dapat berupa time tracking bawaan produk ini atau provider dari ekstensi.
  2. Working hours per day (default 8, boleh desimal) dan working days per week (default 5) dapat diatur; nilai di luar 1-24 jam/1-7 hari ditolak.
  3. Format tampilan waktu: pretty (`1w 2d`), days, atau hours.
  4. Default unit untuk input tanpa satuan: minute, hour, day, atau week.
  5. Disable time tracking menyembunyikan field time tracking dan Log work di semua project; data worklog tetap tersimpan.
  6. Team-managed mengaktifkan time tracking lewat pengaturan Estimation project.
  7. ⚠ Legacy mode tidak dibangun (hanya Data Center); opsi "Copy comment to work description" mengikuti pemilik di TIME-002 butir 6, karena inventaris menandainya Data Center-only sementara modul 14 memperlakukannya sebagai fitur Cloud (lihat Pertanyaan terbuka).
- **Team-managed vs company-managed:** Team-managed memakai pengaturan Estimation per project di atas setting global.
- **Catatan & ketergantungan:** Pemilik: TIME-002 Time Tracking Settings di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) memiliki versi lengkap (rumus konversi, validasi dan pesan galat, API `configuration/timetracking/options`); aktivasi provider dimiliki TIME-001. CFG-058 adalah titik masuk halaman admin global di modul ini dan harus selaras dengan TIME-002. Permission Work on issues serta edit/delete worklog di PERM-032 ([03 Izin](03-izin-peran-keamanan.md)). Field yang terpengaruh di CFG-017.

## Skema default, berbagi & alat skema

### CFG-059 Default configuration & scheme sharing

- **Deskripsi:** Mekanisme scheme default untuk project baru dan berbagi konfigurasi antar company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; pembuat project (permission Create projects/Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Scheme default yang tersedia: Default Issue Type Scheme, Default Field Configuration (Scheme), Default Screen Scheme, Default Issue Type Screen Scheme, dan Default priority scheme.
  2. Pembuatan company-managed project menyediakan opsi "Share settings with an existing project" yang memakai scheme yang sama dengan project acuan.
  3. Project dari template mendapat scheme baru berprefix key project (mis. "ABC: Scrum Issue Type Scheme").
  4. Admin dapat menyalin scheme untuk dimodifikasi tanpa memengaruhi project lain.
  5. Admin dapat mengidentifikasi dan menghapus scheme yang tidak terpakai (filter "tidak dipakai project mana pun").
  6. Saat mengedit scheme yang dipakai banyak project, ditampilkan peringatan yang menyebut jumlah project terdampak.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pembuatan project di modul 04. Fase 1-2 project company-managed memakai scheme default; prefix per project aktif sejak template punya konfigurasi khusus.

### CFG-060 Scheme tools (Scheme comparison, Group to project role mapping, Bulk delete schemes)

- **Deskripsi:** Kumpulan alat admin Data Center untuk membandingkan scheme, mengganti grant berbasis group menjadi project role secara massal, dan menghapus scheme tak terpakai sekaligus; tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs / System administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: Scheme Comparison Tool (matriks perbedaan grant antar permission/notification scheme).
  2. Tidak dibangun: wizard Group to Project Role Mapping untuk permission/notification/issue security scheme.
  3. Tidak dibangun: opsi menambahkan anggota group ke project role saat mapping.
  4. Tidak dibangun: Bulk Delete Schemes; padanan di produk ini adalah filter scheme tak terpakai per halaman scheme (CFG-059 butir 5).
  5. Tidak dibangun: ringkasan hasil (jumlah scheme diubah, role membership ditambahkan).
  6. Bila kelak dibangun, akses dibatasi Admin Situs/System administrator.
  7. Tidak ada padanan langsung di Cloud; pembersihan dilakukan per halaman scheme.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bila diperlukan kelak, lebih tepat ditempatkan di modul 03. Lihat Di luar cakupan.

### CFG-061 Field and issue type translations

- **Deskripsi:** Terjemahan nama/deskripsi custom field, work type, priority, resolution, dan status per bahasa; fitur Data Center yang tidak dibangun di fase mana pun saat ini.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: terjemahan nama dan deskripsi custom field per bahasa.
  2. Tidak dibangun: terjemahan work type, priority, resolution, dan status.
  3. Bila kelak dibangun, pengguna melihat label sesuai bahasa profil (EN/ID, selaras `LanguageMiddleware`) dengan fallback ke nama asli.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk ini sudah mendukung pesan galat EN/ID (`internal/shared/constants/error.go`); terjemahan data konfigurasi buatan admin adalah keputusan terbuka. Lihat Di luar cakupan.

## Field kalkulasi (tambahan di luar inventaris)

### CFG-062 ⚠ Formula custom field (calculated fields)

- **Deskripsi:** Tipe custom field native yang nilainya dihitung otomatis dari system field dan custom field lain lewat ekspresi formula, sehingga tim tidak perlu mengekspor data ke spreadsheet untuk menghitung (padanan fitur produk acuan yang dirilis 27 Juli 2026 dengan judul "Stop exporting and start calculating"). Fitur ini berbeda dari Formula Fields di Product Discovery (JPD-IDE-036, modul 36) yang khusus untuk skor dan agregasi idea di project discovery.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) membuat dan mengubah formula global atau per context; Admin Proyek/Project Lead (Administer projects) membuat formula di team-managed project; semua pengguna dengan Browse projects melihat hasilnya; tidak ada pengguna yang dapat mengisi nilainya secara manual.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Wizard Create custom field menyediakan tipe "Formula" dengan tipe keluaran wajib dipilih saat pembuatan: Number, Text, Date, atau Duration; tipe keluaran tidak dapat diubah setelah field dibuat (selaras CFG-034 butir 5).
  2. Editor ekspresi menyediakan katalog fungsi (aritmetika, teks, tanggal, kondisional) dengan autocomplete nama field dan fungsi, serta validasi sintaks langsung yang menunjuk posisi galat sebelum field dapat disimpan.
  3. ⚠ Editor menyediakan bantuan bahasa alami lewat AI opsional untuk menyusun ekspresi dari deskripsi teks; bantuan ini mengikuti pengaturan AI site (ADM-052, modul 22), dimatikan sepenuhnya pada profil regulated deployment (VISI-003), dan ekspresi hasil AI selalu ditampilkan untuk ditinjau sebelum disimpan.
  4. Input formula dapat merujuk system field (mis. `created`, `updated`, `duedate`, `priority`, time tracking CFG-017) dan custom field lain (mis. Number, Date, Select) yang berlaku pada context yang sama; rujukan ke field yang tidak tersedia ditolak dengan galat "Field tidak tersedia di scope formula ini".
  5. ⚠ Nilai hasil disimpan dan diindeks sehingga dapat dipakai di JQL dengan operator sesuai tipe keluaran dan di `ORDER BY`, mis. `"Budget Remaining" < 0 ORDER BY "Budget Remaining" ASC` (searcher, CFG-038; grammar di modul 12).
  6. Hasil formula tampil di issue view (ditata lewat layout CFG-049/CFG-042), sebagai kolom di issue navigator dan List view (modul 13 dan 18), serta di gadget dashboard yang mendukung field (modul 17).
  7. ⚠ Perubahan nilai hasil formula tercatat di history work item dan dapat memicu trigger otomasi Field value changed (AUTO-021, modul 20); perubahan yang dipicu perhitungan ulang massal karena formula diedit tidak memicu notifikasi email per work item.
  8. Scope diatur lewat field context (CFG-035): global atau per space/project dan work type tertentu; Formula custom field tersedia di company-managed lewat halaman Fields dan di team-managed lewat field library project (CFG-041).
  9. Nilai dihitung ulang setiap kali field sumber berubah pada work item yang sama; mengedit ekspresi memicu perhitungan ulang semua work item terdampak sebagai job background dengan indikator progres (modul 27).
  10. Referensi melingkar (formula A merujuk B yang merujuk A, langsung atau tidak langsung) ditolak saat menyimpan dengan galat "Formula membentuk referensi melingkar"; galat evaluasi per work item (mis. pembagian dengan nol, field sumber kosong) menghasilkan nilai kosong dan tidak menggagalkan penyimpanan work item.
  11. Pembuatan, perubahan ekspresi, perubahan tipe keluaran yang ditolak, dan penghapusan Formula custom field tercatat di audit log (modul 22) beserta ekspresi sebelum dan sesudahnya.
  12. Field bersifat read-only di Create dan Edit screen, Bulk edit, CSV import, dan REST API; upaya menulis nilai lewat API ditolak HTTP 400 dengan pesan bahwa field dihitung otomatis.
- **Team-managed vs company-managed:** Fungsi formula dan perilaku sama; company-managed mengatur scope lewat context dan field configuration, team-managed lewat field library dan layout per work type (CFG-041, CFG-042).
- **Catatan & ketergantungan:** Ditambahkan di luar inventaris (added_beyond_inventory); seluruh rincian bertanda ⚠ perlu dikonfirmasi terhadap dokumentasi produk acuan, termasuk tier ketersediaannya. Rujukan silang: JQL dan `ORDER BY` di [12 JQL](12-jql-bahasa-kueri.md), gadget di [17 Dashboard & Gadget](17-dashboard-gadget.md), trigger AUTO-021 di [20 Otomasi](20-otomasi.md), Formula Fields discovery JPD-IDE-036 di [36 Product Discovery: Ideas & Fields](36-jpd-ideas-field-views.md) (entitas dan editor terpisah; keduanya sebaiknya berbagi pustaka evaluator ekspresi). Kebutuhan kalkulasi dengan data eksternal tetap lewat ekstensi (CFG-033). Model data: tabel `custom_fields` menambah kolom `formula_expression` dan `output_type`, nilai hasil disimpan seperti nilai custom field biasa agar dapat diindeks (modul 28).

## Pertanyaan terbuka

- Apakah produk ini memakai satu field estimasi terpadu (Story point estimate) untuk company-managed dan team-managed, atau mempertahankan dua field berbeda ID demi kompatibilitas impor (CFG-016)?
- Apakah company-managed langsung memakai model Field schemes yang disederhanakan saja, atau mendukung juga field configuration + field configuration scheme klasik (CFG-043, CFG-044, CFG-045)?
- Berapa angka batas sistem default untuk jumlah work type per team-managed project, opsi per context, context per field, field per project/work type, custom field per site, dan level hierarki kustom (CFG-001, CFG-002, CFG-007, CFG-028, CFG-035, CFG-039, CFG-041)?
- Apakah Target start/Target end dan level hierarki di atas Epic hanya tersedia di tier Premium, atau produk ini (sebagai open source) membukanya untuk semua instalasi (CFG-007, CFG-022)?
- Apakah Category field tersedia di semua tipe project atau hanya project bisnis (CFG-021)?
- Apakah dibutuhkan terjemahan data konfigurasi (nama work type, field, priority, resolution, status) dalam EN/ID, mengingat produk sudah mendukung pesan galat dwibahasa (CFG-061)?
- Apakah upload ikon work type mendukung SVG, mengingat risiko XSS pada SVG yang diunggah (CFG-004)?
- Apakah subtask boleh langsung berada di bawah Epic pada company-managed project (sumber menyebut mungkin, tetapi aturan Parent di modul ini mewajibkan parent dari Base level) (CFG-006, CFG-005, CFG-013)?
- Apakah default Description per work type di company-managed disediakan sebagai fitur native (seperti di team-managed), atau cukup dilayani aturan otomasi seperti di sumber (CFG-052)?
- Apakah opsi "Copy comment to work description" dibangun? Inventaris menandainya Data Center-only (CFG-058 butir 7), sedangkan TIME-002 butir 6 memperlakukannya sebagai fitur Cloud; pemilik keputusan adalah modul 14.
- Apakah Formula custom field (CFG-062) tersedia di semua tier atau dibatasi tier berbayar, dan berapa batas jumlah Formula field per site serta kedalaman rantai formula yang saling merujuk?
- Bagaimana perilaku format Percentage pada Number field: menyimpan 0.25 dan menampilkan 25%, atau menyimpan 25 dan menampilkan 25% (CFG-026)?

## Di luar cakupan & ekosistem

- **Fitur Data Center-only/deprecated (hanya rujukan, tidak dibangun):** Configure terminology untuk rename Sprint/Epic (CFG-010); field legacy Epic Link, Epic Name, Epic Status, Epic Colour, dan Parent Link sebagai field tersendiri (CFG-024, hanya alias JQL dan migrasi data); Custom field optimizer (CFG-040); Scheme tools berupa Scheme Comparison, Group to Project Role Mapping, dan Bulk Delete Schemes (CFG-060); Field and issue type translations (CFG-061); halaman admin LexoRank management (CFG-015); legacy mode time tracking (CFG-058); status "Copy comment to work description" ditentukan di TIME-002.
- **App Marketplace pihak ketiga (tidak dibangun sebagai fitur native):** field pack/tipe field tambahan (mis. kalkulasi yang memerlukan data eksternal di luar Formula custom field native CFG-062, field bersyarat, dynamic forms), field-level security/permission per field, checklist sebagai custom field, behaviours/scripted fields, dan tipe field Jira Toolkit DC (Date of First Response, Participants of an issue, Message custom field). Kebutuhan ini dilayani lewat extension point tipe custom field (CFG-033) dan otomasi (modul 20).
- **Layanan lintas produk Atlassian:** direktori Atlassian Teams sebagai layanan eksternal (produk ini memakai entitas Teams internal, CFG-019), pustaka avatar Atlassian (diganti pustaka ikon milik produk ini, CFG-004), dan integrasi Confluence untuk remote link (rujukan ke modul 25).
- **Didefinisikan di modul lain:** work type dan field khusus Service Management (modul 29) dan Product Discovery (modul 36), workflow scheme dan transition screen (modul 09), permission/notification/issue security scheme (modul 03 dan 19), serta perilaku runtime create/edit/view work item (modul 06).
