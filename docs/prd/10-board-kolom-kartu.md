# 10 — Board: Kolom, Kartu, Swimlane & Kanban

Modul ini mendefinisikan board di produk ini: cara board dibuat (dari template project Kanban/Scrum, dari project, atau dari saved filter lintas project), lokasi, administrator, visibilitas, dan navigasi board; tampilan board Kanban dan board sprint aktif beserta kartu (card) work item, drag-and-drop transisi, Rank, panel detail, inline create, menu aksi kartu, flag, dan sub-task; konfigurasi kolom (pemetaan status, unmapped statuses, Simplified Workflow, column constraint, WIP limit, semantik kolom Done); swimlane, Group by, quick filter, filter bar, View settings, dan mode stand-up; kustomisasi kartu (card layout, card colors, days in column) serta pengaturan board yang memengaruhi perhitungan waktu (working days, estimation statistic, toggle Timeline); alur khusus Kanban (Kanplan, Epics panel, Versions panel, Kanban sub-filter, Release dari board); perilaku board team-managed dan board project bisnis; serta batas skala, REST API Agile, dan event board/sprint. Backlog Scrum dan siklus hidup sprint (buat, mulai, selesaikan) ada di modul 11; modul ini hanya mengatur bagaimana sprint aktif ditampilkan di board. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga seluruh entitas board (board, kolom, pemetaan status, swimlane, quick filter, aturan warna, card layout, administrator, working days, preferensi tampilan per pengguna) adalah pekerjaan baru yang dikemas sebagai feature module `internal/features/board` mengikuti pola Clean Architecture repo dan dirender oleh SPA SvelteKit di `frontend/`.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects, Transition Issues, Schedule Issues, Edit/Delete Issues, Manage Sprints, issue security level.
- [04 Proyek](04-proyek.md) — galeri template (Kanban/Scrum), company-managed vs team-managed, akses project Open/Limited/Private, Features toggle.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — work type, field sistem/custom, screen, field context, field Flagged dan Story points.
- [06 Work Item Inti](06-work-item-inti.md) — issue view, Create dialog, field Rank, Flag, Parent, Subtask, archive.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category, transition, condition, validator, post function, transition screen, workflow team-managed.
- [11 Backlog & Sprint](11-backlog-sprint.md) — backlog Scrum, sprint, parallel sprints, Complete sprint, Manage Sprints Permission.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — grammar JQL, saved filter, sharing filter, issue navigator.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — Original/Remaining estimate, Time spent, estimasi Story points.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — version, fixVersion, halaman Releases, Manage versions.
- [16 Laporan & Insight](16-laporan-insight.md) — Burndown, Velocity, CFD, Control chart, Insights.
- [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Timeline board, tampilan proyek bisnis.
- [19 Notifikasi](19-notifikasi-email.md) — notifikasi perubahan work item akibat aksi di board.
- [20 Otomasi](20-otomasi.md) — trigger Sprint created/started/completed, aksi otomasi pada board.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — sidebar project, Starred/Recent, keyboard shortcut global.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST API Agile dan webhook board/sprint.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — project bisnis dan tampilan board-nya.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — realtime, batas kinerja, paginasi.
- [28 Model Data Inti](28-model-data-inti.md) — tabel board, kolom, rank, dan katalog event domain.

## Persona terkait

- **Developer/Anggota Tim** — memindahkan kartu antar kolom, mengurutkan dengan Rank, membuat work item inline, memberi flag, memfilter board, dan menjalankan stand-up.
- **Scrum Master/Team Lead** — memantau sprint aktif, mengatur WIP limit, swimlane, quick filter, dan memandu stand-up harian.
- **Admin Proyek/Project Lead** — membuat board, menjadi board administrator, mengatur kolom, status, Features toggle team-managed, dan Release dari board.
- **Product Owner/PM** — mengelola Kanplan, Epics panel, dan Versions panel untuk menyiapkan antrean kerja.
- **Admin Situs** — mengonfigurasi board apa pun, mengelola workflow di balik Simplified Workflow, dan menangani board yang rusak (filter terhapus).
- **Stakeholder/Viewer** — melihat board lintas project secara read-only lewat tautan atau Starred board.
- **Pengguna Tim Bisnis** — memakai board project bisnis dengan kolom = status yang sederhana.
- **Integrator/Developer API** — membaca dan mengubah board, sprint, dan Rank lewat REST API Agile serta menerima webhook board/sprint.

## User story modul

- Sebagai Admin Proyek/Project Lead, saya ingin membuat board Kanban atau Scrum dari satu atau beberapa project maupun dari saved filter agar tim lintas project punya satu tempat kerja bersama (BRD-002, BRD-003).
- Sebagai Developer/Anggota Tim, saya ingin menyeret kartu ke kolom lain untuk menjalankan transisi workflow dengan validasi yang jelas agar status pekerjaan selalu akurat tanpa membuka work item (BRD-016).
- Sebagai Developer/Anggota Tim, saya ingin mengurutkan kartu dengan drag vertikal dan aksi Move to top agar prioritas tim terlihat dari posisi kartu (BRD-017, BRD-020).
- Sebagai Scrum Master/Team Lead, saya ingin melihat sprint aktif lengkap dengan sisa hari, sprint goal, dan jumlah kartu per kolom agar saya dapat memantau kemajuan sprint setiap hari (BRD-014).
- Sebagai Scrum Master/Team Lead, saya ingin menetapkan WIP limit minimum dan maksimum per kolom dengan penanda warna agar hambatan alur kerja segera terlihat (BRD-028, BRD-029).
- Sebagai Admin Proyek/Project Lead, saya ingin memetakan status workflow ke kolom dan menyembunyikan status tertentu lewat unmapped statuses agar board mencerminkan cara kerja tim tanpa mengubah workflow (BRD-023, BRD-025, BRD-026).
- Sebagai Scrum Master/Team Lead, saya ingin membagi board ke dalam swimlane berbasis JQL, story, assignee, atau epic agar pekerjaan prioritas tinggi (Expedite) terpisah dari yang lain (BRD-031, BRD-032).
- Sebagai Developer/Anggota Tim, saya ingin memfilter board berdasarkan teks, assignee, epic, dan quick filter tanpa mengubah konfigurasi board agar saya fokus pada pekerjaan saya sendiri (BRD-033, BRD-034).
- Sebagai Developer/Anggota Tim, saya ingin menandai work item yang terhambat dengan flag beserta komentar alasan agar tim langsung tahu kartu mana yang butuh bantuan (BRD-021).
- Sebagai Product Owner/PM, saya ingin memisahkan antrean (Kanplan) dari board kerja serta mengelompokkan item per epic dan version agar perencanaan tidak mengganggu tim yang sedang bekerja (BRD-043, BRD-044, BRD-045).
- Sebagai Admin Proyek/Project Lead, saya ingin merilis semua kartu di kolom Done sebagai sebuah version langsung dari board agar board tetap bersih dan riwayat rilis tercatat (BRD-046, BRD-047).
- Sebagai Scrum Master/Team Lead, saya ingin menjalankan mode stand-up yang memfilter board per orang secara bergiliran agar rapat harian singkat dan terarah (BRD-036).
- Sebagai Admin Proyek/Project Lead di project team-managed, saya ingin mengelola kolom langsung dari board sehingga menambah kolom sama dengan menambah status agar tim kecil tidak perlu memahami skema workflow (BRD-048, BRD-049).
- Sebagai Integrator/Developer API, saya ingin API board, sprint, dan rank yang stabil serta webhook saat konfigurasi board berubah agar alat eksternal tetap sinkron (BRD-052, BRD-053).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 18 |
| P1 | 23 |
| P2 | 12 |
| P3 | 2 |
| **Total** | **55** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 18 |
| Fase 2 | 27 |
| Fase 3 | 9 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 membangun entitas board yang selalu berbasis saved filter (termasuk filter otomatis `project = KEY ORDER BY Rank ASC` untuk board dari template), kolom dengan pemetaan status berdasarkan status category, tampilan board Kanban dan board sprint aktif, kartu dengan elemen default, drag-and-drop yang memanggil mesin transisi modul 09, Rank leksikografis global, panel detail, inline create, menu aksi kartu, filter bar dasar, serta perilaku board team-managed (kolom = status), dan endpoint internal REST API Agile `/api/v1/agile/...` yang dipakai SPA (BRD-052, selaras API-032 s.d. API-034). Sejak Fase 1, setiap aksi board harus memakai endpoint domain yang sama dengan issue view (transisi, rank, edit field) sehingga permission dan event domain (modul 28) konsisten. Fase 2 menambah konfigurabilitas: halaman Board settings lengkap (administrator, location, unmapped statuses, constraint dan WIP limit, swimlane, quick filter, card colors, card layout, working days, estimation), Group by, flag, Kanplan dengan panel Epics/Versions, Kanban sub-filter dan Release dari board, directory dan switcher board, Copy/Delete board, serta guardrail batas jumlah kartu. Beberapa fitur P2 sengaja ditempatkan di Fase 2 karena menjadi prasyarat langsung fitur Fase 1-2 lain: Board Location (dipakai wizard Create board), Copy/Delete board (siklus hidup board yang sudah dapat dibuat bebas sejak Fase 1), Working days (sisa hari sprint dan Burndown), serta infinite scroll dan guardrail kinerja board besar (BRD-051). Epics panel dan Versions panel (P1) hanya menjadi rujukan ke pemiliknya, SPR-008 dan REL-021, sedangkan Release from board merujuk REL-022, agar sprint planning tidak dikerjakan ganda. Fase 3 menambah kemampuan yang bergantung pada modul skala organisasi: multiple board di team-managed, board project bisnis, toggle Timeline, kontrol Simplified Workflow dari board (Add status, Set resolution, migrasi), View settings beserta saved views board (BRD-055), indikator Days in column, mode stand-up, lapisan kompatibel publik REST API Agile (read-only Fase 2, tulis Fase 3), dan event webhook/otomasi board. Konfigurasi board disimpan sebagai data relasional terversi agar Copy board, event `board_configuration_changed`, dan audit dapat dibangun tanpa migrasi ulang.

## Fitur

## Pembuatan, siklus hidup & navigasi board

### BRD-001 Kanban project template

- **Deskripsi:** Template project Kanban (team-managed dan company-managed) yang dalam satu langkah membuat project, workflow default, dan board Kanban siap pakai. Ini titik masuk utama bagi tim continuous flow.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan global permission Create projects (company-managed, biasanya Admin Situs) atau Create team-managed projects; Admin Proyek mengubah fitur setelah project dibuat.
- **Aturan perilaku (acceptance criteria):**
  1. Galeri template (PRJ, modul 04) menampilkan template "Kanban" di kategori Software development; wizard meminta pengguna memilih tipe project team-managed atau company-managed sebelum melanjutkan.
  2. ⚠ Fase 1: project company-managed dari template Kanban memakai workflow template statis bawaan, yaitu salinan read-only per project berstatus To Do, In Progress, Done (kategori To Do/In Progress/Done) dengan global transitions dan tanpa board-managed status editing (tombol Add status dan Set resolution tidak tampil); workflow ini tidak memakai Default Workflow Scheme sistem ("jira", WF-020). Fase 3: saat BRD-027/WF-026 tersedia, workflow ini dikonversi (diberi flag `simplified`) menjadi "Software Simplified Workflow for Project <KEY>" tanpa migrasi status karena himpunan statusnya sama; project baru sejak Fase 3 langsung memakai Simplified Workflow. Opsi konfigurasi instance untuk memakai set lama Backlog, Selected for Development, In Progress, Done tersedia bersama BRD-027.
  3. Project company-managed dari template Kanban mendapat tepat satu board Kanban yang memakai saved filter baru bernama "Filter for <KEY> board" dengan JQL `project = <KEY> ORDER BY Rank ASC`, dimiliki oleh pembuat project.
  4. ⚠ Board company-managed default dibuat dengan swimlane berbasis Queries berisi "Expedite" (`priority = Blocker`) dan "Everything Else", serta quick filter "Only My Issues" dan "Recently Updated" (BRD-031, BRD-033).
  5. Project team-managed dari template Kanban mendapat satu board otomatis dengan tiga kolom yang identik dengan status To Do, In Progress, Done.
  6. Kanban backlog (Kanplan, BRD-043) nonaktif secara default; Admin Proyek dapat mengaktifkannya kemudian lewat Board settings > Columns (company-managed) atau Project settings > Features > Backlog (team-managed).
  7. ⚠ Di team-managed, fitur Backlog, Timeline, Reports, Releases, dan Sprints dapat di-toggle lewat Project settings > Features (BRD-049); menyalakan Sprints pada project Kanban mengubah board menjadi berperilaku Scrum (board sprint aktif + backlog sprint) tanpa membuat board baru.
  8. Template membuat work type default sesuai tipe project: Epic, Story/Task, Bug, dan Subtask (company-managed memakai issue type scheme template; team-managed membuat work type lokal project).
  9. Layar preview template menampilkan deskripsi, diagram workflow, dan daftar work type sebelum pengguna menekan "Use template".
  10. Wizard meminta Name dan Key project (validasi key mengikuti modul 04); board yang dibuat otomatis bernama "<KEY> board".
  11. Project company-managed yang dibuat dari template lain (mis. Scrum) tetap dapat diberi board Kanban tambahan lewat Create board (BRD-002); project team-managed tidak dapat diberi board Kanban kedua kecuali lewat BRD-008.
  12. Tipe board company-managed (Scrum/Kanban) tidak dapat diubah in-place; UI Board settings tidak menyediakan opsi ganti tipe, dan dokumentasi bantuan mengarahkan pengguna membuat board baru dengan filter yang sama.
- **Team-managed vs company-managed:** Company-managed membuat saved filter, workflow template (statis read-only di Fase 1, Simplified Workflow sejak Fase 3), swimlane Expedite, dan quick filter default; team-managed membuat board yang terikat ke project dengan kolom = status dan fitur yang diatur via Features toggle.
- **Catatan & ketergantungan:** Bergantung pada galeri template dan pembuatan project di [04 Proyek](04-proyek.md), workflow di [09 Workflow](09-workflow-status-transisi.md), dan saved filter di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md). Keputusan fase workflow template (butir 2): workflow template statis Fase 1 adalah workflow per project yang di-seed dari definisi template bawaan, terpisah dari Default Workflow Scheme sistem (WF-020, dipakai project company-managed yang tidak dibuat dari template software); WF-020 butir 7 di modul 09 perlu diselaraskan oleh pemilik modul 09 agar menyebut workflow template statis di Fase 1 dan Simplified Workflow (WF-026) sejak Fase 3. Label UI "issue"/"work item" dan "project"/"space" sebaiknya dapat dikonfigurasi lewat kamus i18n (EN/ID) karena terminologi Jira Cloud terus berubah. Template Scrum dibahas di [11 Backlog & Sprint](11-backlog-sprint.md).

### BRD-002 Scrum Board Creation

- **Alias inventaris:** Create board (company-managed)
- **Deskripsi:** Membuat board Scrum atau Kanban baru, baik otomatis saat membuat project dari template maupun lewat wizard Create board dari project yang sudah ada atau dari saved filter. Board adalah tampilan atas sekumpulan work item yang ditentukan oleh filter JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna terautentikasi yang punya Browse Projects pada project sumber atau akses lihat ke saved filter; pembuat otomatis menjadi Board administrator; Admin Situs dapat membuat board untuk project mana pun.
- **Aturan perilaku (acceptance criteria):**
  1. Entry point "Create board" tersedia di board switcher sidebar project, tombol "Create board" di halaman View all boards (BRD-006), dan dijalankan otomatis saat project dibuat dari template Scrum atau Kanban.
  2. ⚠ Entry point tambahan "Create board" tersedia di menu global Search/Filters (bersama "View all boards"); posisi persis di navigasi mengikuti modul 21.
  3. Langkah pertama wizard meminta tipe board: Scrum atau Kanban; pilihan ini menentukan apakah board memiliki backlog sprint (Scrum) atau alur kontinu (Kanban).
  4. Langkah kedua (company-managed) menawarkan dua sumber: "Board from an existing project" dan "Board from an existing saved filter".
  5. ⚠ Opsi lama "Board with sample data" dan "Board with new project" tidak disediakan; pembuatan project baru hanya lewat galeri template modul 04.
  6. Board name wajib diisi (1-255 karakter, spasi di awal/akhir dipangkas); nama board duplikat diperbolehkan dan tidak memicu galat.
  7. Sumber "existing project" menerima satu atau beberapa project (multi-select) dan sistem otomatis membuat saved filter baru bernama "Filter for <board name>" dengan JQL `project in (<KEY1>, <KEY2>) ORDER BY Rank ASC` (untuk satu project: `project = <KEY> ORDER BY Rank ASC`) yang dimiliki pembuat board.
  8. Sumber "existing saved filter" hanya menampilkan filter yang dimiliki atau di-share ke pengguna; filter tetap dimiliki owner aslinya dan tidak disalin (BRD-003).
  9. Wizard meminta Location board: sebuah project atau profil pengguna (board personal); default adalah project pertama yang dipilih (BRD-004).
  10. Validasi: pengguna harus memiliki Browse Projects pada setiap project yang dipilih; project tanpa izin tidak muncul di pemilih, dan request API yang memuatnya ditolak dengan HTTP 403 dan pesan "Anda tidak memiliki akses ke project <KEY>".
  11. Kolom awal board dibuat dari status workflow semua project terkait dan dipetakan menurut status category: To Do → kolom "To Do", In Progress → "In Progress", Done → "Done" (BRD-025).
  12. ⚠ Konfigurasi default board baru: estimation statistic = Story Points (Scrum), card colors berdasarkan Issue Types, quick filter "Only My Issues" dan "Recently Updated", serta swimlane default "No swimlanes" untuk Scrum dan Queries Expedite untuk Kanban.
  13. Setelah dibuat, pengguna diarahkan ke Backlog (Scrum) atau Board (Kanban) yang baru, dan pembuat otomatis tercatat sebagai Board administrator (BRD-011).
  14. Tipe board tidak dapat diubah setelah dibuat di company-managed; alternatif yang ditawarkan adalah membuat board baru dengan filter yang sama.
  15. Company-managed: satu project dapat memiliki banyak board dan satu board dapat mencakup banyak project; tidak ada batas jumlah board per project di Fase 1 selain guardrail global modul 27.
  16. Filter tanpa `ORDER BY Rank` tetap menghasilkan board, tetapi ranking dinonaktifkan dan board menampilkan peringatan dengan tombol "Add Rank" (BRD-017).
  17. Work item dari project yang tidak dapat di-browse oleh pengguna yang sedang melihat tidak ditampilkan bagi pengguna tersebut, walaupun cocok dengan filter board (BRD-009).
  18. Team-managed: board dibuat otomatis bersama project dan tidak ada wizard Create board; perilaku Scrum ditentukan oleh toggle Sprints di Features (BRD-049), bukan oleh pilihan tipe di wizard.
  19. ⚠ Global permission "Share dashboards and filters" memengaruhi apakah saved filter otomatis dapat di-share; bila pengguna tidak memilikinya, filter dibuat privat dan wizard menampilkan peringatan bahwa board hanya terlihat oleh pembuatnya sampai filter di-share.
  20. Wizard dapat dibatalkan di langkah mana pun tanpa membuat filter maupun board (pembuatan filter dan board terjadi dalam satu transaksi saat tombol Create ditekan).
  21. Board baru tercatat dengan `created_by`, `created_at`, dan menerbitkan event domain `board_created` (BRD-053).
  22. Kanban board baru memakai pengaturan Kanban default: tanpa Kanplan, sub-filter `fixVersion in unreleasedVersions() OR fixVersion is EMPTY`, dan hide completed issues 2 minggu (BRD-046).
- **Team-managed vs company-managed:** Company-managed memakai wizard dengan pilihan sumber, location, dan tipe; team-managed membuat board otomatis per project dan board tambahan hanya lewat BRD-008.
- **Catatan & ketergantungan:** Saved filter dan sharing di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md); permission di [03 Izin](03-izin-peran-keamanan.md). Board yang dibuat otomatis dari template Scrum/Kanban di Fase 1 memakai workflow template statis read-only (To Do/In Progress/Done) sesuai BRD-001 butir 2; mapping kolom butir 11 berlaku apa adanya, dan kontrol board-managed status (Add status, Set resolution) baru tampil setelah konversi ke Simplified Workflow di Fase 3 (BRD-027, WF-026). Model data: tabel `boards (id, name, type, filter_id, location_type, location_id, created_by)`; repo saat ini belum memiliki entitas filter sehingga modul 13 harus tersedia minimal dalam bentuk filter internal di Fase 1.

### BRD-003 Board from Saved Filter & Cross-project Boards

- **Alias inventaris:** Cross-project Kanban boards
- **Deskripsi:** Board company-managed dapat dibuat dari saved filter JQL apa pun sehingga menampilkan work item dari beberapa project sekaligus. Isi board, ranking, dan sprint mengikuti hasil filter tersebut.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator, owner filter, Admin Situs; Manajer lintas tim sebagai pengguna utama board lintas project; izin per work item tetap mengikuti project asalnya.
- **Aturan perilaku (acceptance criteria):**
  1. Wizard Create board menampilkan pemilih saved filter berisi filter yang dimiliki pengguna atau di-share kepadanya, dengan pencarian nama filter.
  2. Filter harus berakhiran `ORDER BY Rank` agar ranking aktif; bila tidak, board menampilkan banner "Ranking dinonaktifkan karena filter tidak diurutkan berdasarkan Rank" dengan tombol "Add Rank" untuk pengguna yang boleh mengedit filter.
  3. Filter boleh memuat beberapa project, component, label, team, atau JQL kompleks apa pun yang valid; board yang dihasilkan disebut board lintas project.
  4. JQL lintas project yang umum (mis. `project in (A, B)` atau `labels = platform`) didukung tanpa batasan tambahan selain validitas JQL modul 12.
  5. Board menyimpan referensi ke filter (bukan salinan); perubahan JQL filter langsung mengubah isi board pada pemuatan berikutnya atau lewat pembaruan realtime.
  6. Board dengan filter tanpa klausa project (mis. `assignee = currentUser() ORDER BY Rank`) tetap valid; Location ditetapkan terpisah (BRD-004).
  7. Daftar "Projects in board" di Board settings > General dihitung dari project work item hasil filter, bukan dari teks JQL.
  8. Mapping kolom di board lintas project menampilkan status dari semua workflow project terkait, dikelompokkan per status dengan penanda workflow asal saat nama status sama tetapi berbeda ID.
  9. Swimlane "Projects" (BRD-031) tersedia untuk memisahkan kartu per project.
  10. Sprint yang dibuat di board ini dapat terlihat di board lain bila work item-nya cocok dengan filter board lain (modul 11).
  11. Untuk mengelola sprint di board multi-project, pengguna butuh Manage Sprints di semua project yang terlibat; aksi sprint ditolak dengan pesan yang menyebut project yang kurang izinnya.
  12. ⚠ Release dari board lintas project (BRD-047) meminta pengguna memilih satu project tujuan version; tombol Release disembunyikan bila pengguna tidak memiliki Manage versions di project mana pun.
  13. ⚠ Inline create di board lintas project memakai project location board sebagai default dan menyediakan pemilih project bila board mencakup lebih dari satu project yang mengizinkan Create Issues.
  14. Permission per project (Browse, Transition, Edit, Schedule) tetap dievaluasi per work item; kartu dari project tanpa Transition Issues tidak dapat di-drag oleh pengguna tersebut.
  15. Simplified Workflow (BRD-027) tidak tersedia untuk board yang mencakup lebih dari satu project; tombolnya disembunyikan dengan tooltip alasan.
  16. Owner filter dapat berbeda dari board administrator; mengubah JQL filter memerlukan hak edit pada filter tersebut, sedangkan board administrator non-owner hanya dapat mengganti filter board dengan filter lain yang dapat dilihatnya.
- **Team-managed vs company-managed:** Hanya company-managed. Board team-managed selalu terikat ke project-nya sendiri; ⚠ dukungan board team-managed yang menarik work item dari project lain tidak dibangun dan dicatat sebagai pertanyaan terbuka.
- **Catatan & ketergantungan:** JQL di [12 JQL](12-jql-bahasa-kueri.md); sprint lintas board di [11 Backlog & Sprint](11-backlog-sprint.md). Implementasi harus memakai eksekutor JQL yang sama dengan navigator agar hasil board dan pencarian identik.

### BRD-004 Board Location

- **Deskripsi:** Setiap board company-managed memiliki lokasi (sebuah project atau profil pengguna) yang menentukan di mana board muncul dalam navigasi, terpisah dari isi filter board.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator menetapkan dan mengubah location; pengguna lain hanya melihat dampaknya di navigasi.
- **Aturan perilaku (acceptance criteria):**
  1. Pilihan location adalah project (board muncul di board switcher/sidebar project tersebut) atau pengguna (board personal milik pengguna yang dipilih).
  2. Location ditetapkan saat Create board dan dapat diubah di Board settings > General oleh board administrator.
  3. Board yang berlokasi di sebuah project tampil di daftar board project tersebut walaupun filternya mencakup project lain.
  4. Board yang filternya mencakup project X tetapi berlokasi di project Y tidak otomatis muncul di sidebar project X (kecuali aturan switcher BRD-006 butir 14 diaktifkan).
  5. Board berlokasi pengguna tidak tampil di sidebar project mana pun; board diakses lewat View all boards, Starred, atau Recent.
  6. ⚠ Mengubah location ke project memerlukan Browse Projects pada project tujuan; pemilih hanya menampilkan project yang dapat di-browse pengguna, dan Admin Situs dapat memilih project mana pun.
  7. Breadcrumb dan konteks navigasi (Backlog, Reports, Timeline) mengikuti location board; board berlokasi pengguna menampilkan breadcrumb "Boards / <nama board>".
- **Team-managed vs company-managed:** Team-managed: location selalu project-nya sendiri dan tidak dapat diubah; field Location tidak ditampilkan.
- **Catatan & ketergantungan:** Sidebar dan breadcrumb di [21 Navigasi](21-navigasi-ux-preferensi.md). Menghapus project location (modul 04) mengubah board menjadi berlokasi di pengguna pembuatnya; perilaku ini perlu dikonfirmasi pemilik produk (lihat Pertanyaan terbuka).

### BRD-005 Copy board & delete board

- **Alias inventaris:** Board Copy; Board Delete
- **Deskripsi:** Menduplikasi konfigurasi board company-managed sebagai titik awal board baru, atau menghapus board tanpa menghapus work item, saved filter, maupun project di baliknya.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Copy: pengguna yang dapat melihat board (menjadi admin salinan); Delete: Board administrator atau Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Copy" tersedia dari menu "..." per baris di halaman View all boards.
  2. ⚠ Aksi "Copy board" juga tersedia dari menu "..." di header board bagi pengguna yang dapat melihat board.
  3. Board salinan bernama "Copy of <nama board>" dan dapat di-rename setelahnya di Board settings > General.
  4. Copy menyalin konfigurasi kolom, status mapping, column constraint, WIP limit, swimlane, quick filter, card colors, card layout, estimation, working days, Kanban sub-filter, dan hide completed issues.
  5. ⚠ Board salinan mereferensikan saved filter yang sama (filter tidak diduplikasi); bila pengguna ingin mengubah isi board salinan, ia harus memilih atau membuat filter lain.
  6. Pengguna yang menyalin menjadi satu-satunya administrator board salinan; administrator board asal tidak ikut tersalin.
  7. Sprint tidak diduplikasi; sprint tetap melekat ke work item sehingga dapat tampil di board salinan bila work item-nya cocok dengan filter.
  8. Copy tidak tersedia untuk board team-managed; menu tidak menampilkan opsi tersebut.
  9. Aksi "Delete" tersedia dari menu "..." di View all boards dan dari menu board, hanya untuk Board administrator atau Admin Situs; pengguna lain tidak melihat aksinya dan API mengembalikan HTTP 403.
  10. Dialog konfirmasi menyebut nama board dan menegaskan bahwa work item tidak ikut terhapus; tombol hapus memerlukan klik eksplisit.
  11. Saved filter board tidak ikut terhapus.
  12. ⚠ Sprint milik board tetap tercatat di field Sprint work item; sprint active/future yang kehilangan board asal tetap ada dan dapat dikelola dari board lain yang filternya mencakup work item tersebut, sedangkan sprint tanpa board mana pun hanya dapat diakses lewat API.
  13. Project company-managed boleh tidak memiliki board sama sekali setelah penghapusan.
  14. Data report board (Sprint report, Velocity, Burndown) tidak lagi dapat diakses setelah board dihapus; tautan lama menampilkan "Board tidak ditemukan".
  15. ⚠ Board yang dihapus tidak masuk trash dan tidak dapat dipulihkan; penghapusan bersifat permanen dan dicatat di audit log.
  16. Team-managed: board utama tidak dapat di-copy atau dihapus terpisah dari project; hanya fitur Backlog/Sprints yang dapat dimatikan (BRD-049) dan board tambahan dikelola lewat BRD-008.
  17. Penghapusan menerbitkan event `board_deleted` (BRD-053) dan membersihkan preferensi per pengguna (last board, collapse state, group by) yang merujuk board tersebut.
- **Team-managed vs company-managed:** Copy dan Delete hanya untuk company-managed; board team-managed mengikuti siklus hidup project.
- **Catatan & ketergantungan:** Audit log di [22 Administrasi Sistem](22-administrasi-sistem.md). Penyimpanan konfigurasi board dalam tabel relasional terpisah memudahkan copy dalam satu transaksi.

### BRD-006 Boards Directory (View all boards)

- **Alias inventaris:** Board directory & board switcher
- **Deskripsi:** Halaman daftar semua board yang dapat diakses pengguna, dengan pencarian, filter, sort, star, dan aksi admin, ditambah aturan board mana yang tampil di switcher sidebar project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna melihat board yang filternya dapat mereka lihat; Board administrator dan Admin Situs melihat aksi Copy/Delete.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman "View all boards" diakses dari top navigation (menu Projects/Boards > "View all boards") atau dari board switcher.
  2. Tabel memuat kolom Star, Name, Type (Scrum/Kanban), Admin, dan Location.
  3. Pencarian berdasarkan nama board bersifat case-insensitive dan diterapkan saat mengetik (debounce ±300 ms).
  4. ⚠ Filter berdasarkan project dan tipe board (Scrum/Kanban) dapat dikombinasikan dengan pencarian nama.
  5. Sort tersedia pada kolom Name, Type, Admin, dan Location (naik/turun); default urut nama naik.
  6. Paginasi server-side (default 25 per halaman) untuk instance dengan banyak board.
  7. Menu aksi per baris menampilkan Copy dan Delete sesuai hak (BRD-005); aksi yang tidak diizinkan tidak ditampilkan.
  8. Star/unstar board dari kolom Star membuat board muncul di bagian Starred pada navigasi global (modul 21).
  9. Hanya board yang saved filternya dapat dilihat pengguna yang ditampilkan; board tanpa akses filter tidak muncul.
  10. Tombol "Create board" di halaman ini membuka wizard BRD-002.
  11. Recent boards (maksimal 5 board terakhir dibuka) tampil di menu navigasi global.
  12. ⚠ Board yang filternya tidak di-share ke pengguna tidak tampil di directory; bila pengguna membuka URL board tersebut langsung, board tampil tanpa work item dengan pesan akses (BRD-009).
  13. Sort dan filter tersimpan di query string URL sehingga dapat dibagikan.
  14. Board switcher di sidebar project company-managed menampilkan board yang location-nya project tersebut; board lain yang filternya mencakup project tersebut ditampilkan di bagian terpisah "Board lain yang memuat project ini".
- **Team-managed vs company-managed:** Board team-managed tampil di directory dengan Location = project dan tanpa aksi Copy/Delete.
- **Catatan & ketergantungan:** Navigasi global, Starred, dan Recent di [21 Navigasi](21-navigasi-ux-preferensi.md). Endpoint daftar board dipakai bersama REST API BRD-052.

### BRD-007 Board Switcher & Board Navigation

- **Deskripsi:** Navigasi antar board milik project melalui sidebar project company-managed, beserta URL board permanen, deep link, share link, dan kompatibilitas URL lama.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat board.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown board switcher di sidebar project menampilkan semua board yang berlokasi di project tersebut, diurutkan berdasarkan nama, dengan penanda tipe Scrum/Kanban.
  2. Board terakhir yang dibuka diingat per pengguna per project; membuka project kembali langsung mengarah ke board tersebut.
  3. Item sidebar per board: Backlog (Scrum atau Kanplan aktif), Board/Active sprints, Reports, dan Timeline (bila aktif, BRD-042).
  4. Switcher menyediakan opsi "Create board" (BRD-002) dan "View all boards" (BRD-006).
  5. Recent dan Starred boards tampil di menu navigasi global dan halaman Your work.
  6. URL board permanen dan berbasis ID: `/projects/<KEY>/boards/<ID>`, `/projects/<KEY>/boards/<ID>/backlog`, `/projects/<KEY>/boards/<ID>/reports`; board berlokasi pengguna memakai `/boards/<ID>`.
  7. Deep link work item terpilih lewat parameter `?selectedIssue=<KEY>` membuka panel detail (BRD-018) saat halaman dimuat.
  8. URL legacy `RapidBoard.jspa?rapidView=<ID>` (beserta parameter `view=planning` dan `view=reporting`) di-redirect ke URL board baru untuk kompatibilitas impor dari Jira.
  9. Tombol Share/Copy link di header board dan backlog menyalin URL yang memuat filter aktif (quick filter, assignee) dan menampilkan konfirmasi "Tautan disalin".
  10. ⚠ Team-managed: switcher hanya ditampilkan bila project memiliki lebih dari satu board (BRD-008); dengan satu board, sidebar langsung menampilkan item Board/Backlog.
  11. URL board dengan ID yang tidak ada atau tidak dapat diakses menampilkan halaman "Board tidak ditemukan atau Anda tidak memiliki akses" dengan tautan ke View all boards.
- **Team-managed vs company-managed:** Company-managed memakai switcher berbasis location; team-managed memakai satu board kecuali multiple board aktif.
- **Catatan & ketergantungan:** Sidebar project dan routing SPA di [21 Navigasi](21-navigasi-ux-preferensi.md); SPA fallback `index.html` di `webui.RegisterRoutes` sudah mendukung deep link sisi klien.

### BRD-008 Multiple Boards in Team-managed Software Projects

- **Deskripsi:** Project team-managed software dapat memiliki lebih dari satu board, masing-masing dengan filter subset, kolom, dan pengaturan tampilannya sendiri, sambil tetap memakai workflow team-managed yang sama.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Administrator project team-managed membuat dan mengelola board tambahan; Member memakai board; Scrum Master/Team Lead mengatur sprint per board.
- **Aturan perilaku (acceptance criteria):**
  1. Board tambahan dibuat dari sidebar project (aksi "+ Add board") atau navigasi views project team-managed.
  2. Board tambahan dapat diberi nama, diganti nama, dan dihapus; board utama tidak dapat dihapus selama project ada.
  3. Setiap board tambahan memiliki filter subset berbasis atribut sederhana (label, work type, team, component) yang hanya dapat mempersempit work item project, tidak menarik project lain.
  4. Konfigurasi kolom per board dapat berbeda (urutan dan pengelompokan), tetapi setiap kolom tetap dipetakan ke status workflow team-managed yang sama; menambah status dari satu board membuat status yang tersedia di semua board.
  5. ⚠ Sprint dimiliki per board (origin board) tetapi work item sprint dapat tampil di board lain yang filternya cocok, mengikuti model company-managed; keputusan akhir dicatat di Pertanyaan terbuka.
  6. ⚠ Features Backlog dan Sprints berlaku per board ketika project memiliki banyak board, sedangkan Estimation, Releases, dan Reports berlaku per project.
  7. Board switcher muncul di sidebar project team-managed segera setelah board kedua dibuat (BRD-007).
  8. ⚠ Hanya peran Administrator project yang dapat membuat, mengubah, dan menghapus board; Member hanya dapat memakai board dan mengubah preferensi tampilan pribadinya.
  9. ⚠ Batas jumlah board per project team-managed adalah 20 (konfigurabel oleh Admin Situs); melebihi batas menampilkan pesan "Project ini sudah mencapai batas board".
- **Team-managed vs company-managed:** Khusus team-managed; company-managed sudah mendukung banyak board lewat BRD-002 dan BRD-003.
- **Catatan & ketergantungan:** Fitur relatif baru di Jira Cloud dan masih bergeser; implementasi sebaiknya memakai entitas board yang sama dengan company-managed (dengan `filter_kind = project_subset`) agar tidak ada dua model data. Multiple board view untuk project bisnis ada di BRD-050 dan modul 18.

## Akses, filter board & administrasi board

### BRD-009 Board Visibility & Permissions

- **Deskripsi:** Aturan siapa yang dapat membuka board dan work item apa yang tampil di dalamnya, menggabungkan sharing saved filter, permission project, dan issue security.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna (melihat), Board administrator (mengatur share filter lewat owner filter), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Visibilitas board company-managed mengikuti sharing saved filter board; pengguna tanpa akses filter yang membuka URL board melihat pesan "Board ini memakai filter yang tidak di-share kepada Anda" dan tidak ada work item yang dimuat.
  2. Board settings > General menampilkan tautan "Edit filter shares" yang membuka pengaturan share filter (modul 13); tautan hanya aktif bagi owner/editor filter.
  3. Work item hanya tampil bila pengguna memiliki Browse Projects pada project-nya dan lolos issue security level, walaupun cocok dengan filter; jumlah kartu per kolom hanya menghitung work item yang terlihat oleh pengguna tersebut.
  4. Work item yang diarsipkan (archive Premium/Enterprise, modul 06) tidak tampil di board maupun backlog.
  5. Board settings bersifat read-only untuk non-admin: semua kontrol dinonaktifkan dan banner "Hanya board administrator yang dapat mengubah pengaturan ini" ditampilkan.
  6. Team-managed: akses board mengikuti akses project (Open/Limited/Private) dan role; pengguna yang dapat membuka project dapat membuka board-nya.
  7. Evaluasi izin dilakukan di server untuk setiap endpoint board (daftar kartu, aksi kartu, konfigurasi); UI tidak pernah menjadi satu-satunya penegak izin.
- **Team-managed vs company-managed:** Company-managed: visibilitas = share filter + izin project per work item; team-managed: visibilitas = akses project.
- **Catatan & ketergantungan:** Bergantung pada [03 Izin](03-izin-peran-keamanan.md) (Browse Projects, issue security) dan sharing filter di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md). Middleware `AuthMiddleware` dan klaim `sess` yang sudah ada dipakai untuk identitas pengguna.

### BRD-010 Board filter

- **Deskripsi:** Saved filter JQL yang menentukan work item apa saja yang masuk ke board company-managed, termasuk penggantian filter, sharing, ranking, dan penanganan filter rusak.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board administrator (mengganti filter), owner/editor filter (mengubah JQL dan share).
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > General menampilkan Saved Filter (nama, tautan), Filter Query (JQL read-only), Shares, tautan Edit filter permissions, status Ranking, dan Projects in board.
  2. Tautan "Edit filter query" membuka issue navigator (modul 13) dengan JQL filter; tautan hanya tampil bagi owner atau editor filter, pengguna lain melihat JQL read-only.
  3. Board administrator dapat mengganti saved filter board dengan filter lain yang dimiliki atau di-share kepadanya; penggantian langsung mengubah isi board dan daftar Projects in board.
  4. Shares filter menentukan siapa yang dapat melihat board; bila filter tidak di-share ke siapa pun selain owner, Board settings menampilkan peringatan "Filter ini privat; hanya owner yang dapat melihat board".
  5. Filter wajib berakhiran `ORDER BY Rank` (ASC) agar ranking aktif; status ranking ditampilkan sebagai "Enabled" atau "Disabled" (BRD-017).
  6. Work item dari project tempat pengguna tidak memiliki Browse Projects tidak tampil (BRD-009).
  7. Bila filter board dihapus atau JQL-nya menjadi tidak valid (mis. field dihapus), board menampilkan galat "Filter board tidak tersedia atau tidak valid" beserta tombol "Pilih filter lain" untuk board administrator; menghapus saved filter yang dipakai board memperingatkan owner bahwa N board akan rusak.
  8. Daftar "Projects in board" dihitung dari hasil filter dan diperbarui saat filter atau data berubah.
- **Team-managed vs company-managed:** Team-managed tidak punya board filter yang dapat diubah; board selalu memuat semua work item project (kecuali board tambahan BRD-008).
- **Catatan & ketergantungan:** Grammar JQL di [12 JQL](12-jql-bahasa-kueri.md). Pembuatan dan penyimpanan saved filter mengikuti SRC-036 (Save filter) dan pemakaian saved filter sebagai sumber data board mengikuti SRC-044 (Saved filters as data source) di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md). Eksekusi filter board harus memakai cache hasil per board dengan invalidasi berbasis event work item (modul 27).

### BRD-011 Board Administrators

- **Deskripsi:** Board administrator adalah pengguna atau grup yang berhak mengubah konfigurasi board company-managed. Hak ini terpisah dari permission project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator, Admin Situs, Admin Proyek (lihat butir 3).
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > General memiliki field Administrators untuk menambah dan menghapus pengguna maupun grup.
  2. Pembuat board otomatis menjadi board administrator.
  3. ⚠ Admin Situs selalu dapat mengonfigurasi board apa pun walau bukan board administrator; Admin Proyek (Administer projects) pada project location board juga diperlakukan sebagai board administrator.
  4. Hanya board administrator yang dapat mengubah filter board, kolom, swimlane, quick filter, card colors, card layout, estimation, working days, dan toggle Timeline; API mengembalikan HTTP 403 untuk pengguna lain.
  5. Non-admin melihat Board settings dalam mode read-only (BRD-009 butir 5).
  6. Board administrator tidak otomatis mendapat Manage Sprints; manajemen sprint mengikuti permission scheme project (modul 03, 11).
  7. Kolom "Admin" di View all boards menampilkan administrator board (maksimal 3 nama lalu "+N").
  8. ⚠ Minimal satu administrator (pengguna atau grup) wajib ada; menghapus administrator terakhir ditolak dengan pesan "Board harus memiliki setidaknya satu administrator".
- **Team-managed vs company-managed:** Team-managed tidak memiliki board administrator terpisah; konfigurasi board dikendalikan role Administrator project.
- **Catatan & ketergantungan:** Grup pengguna di [02 Pengguna & Grup](02-pengguna-grup-organisasi.md). Model data: tabel `board_admins (board_id, principal_type, principal_id)`.

### BRD-012 Board settings (company-managed)

- **Alias inventaris:** Board Settings - General
- **Deskripsi:** Kumpulan halaman konfigurasi board company-managed beserta tab General yang memuat nama, administrator, location, filter, sub-filter, dan status ranking.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator dan Admin Situs mengubah; pengguna lain melihat read-only.
- **Aturan perilaku (acceptance criteria):**
  1. Board settings dibuka dari menu "..." header board > "Board settings"/"Configure board"; tab untuk board Kanban: General, Columns, Swimlanes, Quick Filters, Card colors, Card layout, Working days; board Scrum menambah tab Estimation (BRD-041).
  2. ⚠ Tab "Issue Detail View" warisan issue view lama tidak dibangun (BRD-054).
  3. Tab General memuat Board name (rename inline, wajib, 1-255 karakter).
  4. Tab General memuat Administrators (BRD-011): tambah pengguna atau grup; pembuat board otomatis admin.
  5. Tab General memuat Location (project/pengguna, BRD-004).
  6. Tab General memuat Saved filter dengan aksi pilih filter lain, "Edit filter query", daftar Shares, dan tautan "Edit filter shares" (BRD-010).
  7. Tab General menampilkan JQL filter read-only beserta owner filter.
  8. Tab General menampilkan status Ranking (Enabled/Disabled) dan tombol "Add Rank" untuk menambahkan `ORDER BY Rank ASC` ke filter bila pengguna adalah owner/editor filter.
  9. Tab General menampilkan daftar "Projects in board" hasil filter.
  10. Tab General board Kanban memuat Kanban board sub-filter dan opsi Hide completed issues (BRD-046).
  11. ⚠ Admin Situs selalu dapat mengonfigurasi board; Admin Proyek pada project location board diperlakukan setara board administrator (sama dengan BRD-011 butir 3).
  12. Board administrator hanya dapat mengedit JQL filter bila ia owner atau memiliki edit permission pada filter; bila tidak, tombol "Edit filter query" diganti teks "Hubungi owner filter <nama>".
  13. Non-admin melihat semua tab Board settings read-only.
  14. Team-managed: Board settings minimal berisi Columns and statuses dan Card layout; pengaturan lain ada di Project settings > Features (BRD-048, BRD-049).
  15. Setiap penyimpanan perubahan di tab mana pun menampilkan flag "Perubahan disimpan", tercatat di audit log, dan menerbitkan event `board_configuration_changed` (BRD-053).
  16. Perubahan konfigurasi tersimpan per field (autosave per kontrol) dan langsung berlaku bagi semua pengguna yang membuka board tanpa perlu reload manual.
- **Team-managed vs company-managed:** Company-managed memakai halaman Board settings bertab penuh; team-managed hanya Columns and statuses dan Card layout.
- **Catatan & ketergantungan:** Tab lain dirinci di BRD-023 (Columns), BRD-031 (Swimlanes), BRD-033 (Quick Filters), BRD-038 (Card colors), BRD-037 (Card layout), BRD-040 (Working days), BRD-041 (Estimation), BRD-042 (Timeline).

## Tampilan board & interaksi kartu

### BRD-013 Board view (Kanban board UI)

- **Deskripsi:** Tampilan utama board Kanban berupa kolom vertikal berisi kartu work item, dengan drag-and-drop, collapse kolom, dan pembaruan realtime.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects (melihat); Board administrator (akses konfigurasi dari header).
- **Aturan perilaku (acceptance criteria):**
  1. Header board menampilkan nama board, breadcrumb project, tombol Release (di kolom terakhir, BRD-047), tombol Insights (modul 16), ikon Automation (modul 20), dan menu "..." berisi Board settings/Configure board dan Manage workflow (bila berizin).
  2. Header kolom menampilkan nama kolom, jumlah work item, dan indikator WIP limit (BRD-029).
  3. Kartu di setiap kolom diurutkan menurut Rank (BRD-017); bila ranking nonaktif, urutan mengikuti klausa ORDER BY filter board.
  4. Board dengan banyak kolom dapat di-scroll horizontal dengan header kolom sticky saat scroll vertikal.
  5. Empty state ditampilkan untuk kolom kosong ("Tidak ada work item") dan untuk board tanpa hasil filter ("Board ini kosong" beserta tautan Board settings bagi admin atau tautan Clear filters bila filter bar aktif).
  6. Perubahan work item oleh pengguna lain (transisi, rank, edit field, create, delete) tampil tanpa refresh lewat kanal realtime (modul 27) dalam ≤ 5 detik.
  7. Board hanya menampilkan work item yang lolos board filter + Kanban sub-filter + hide completed (BRD-046) dan berstatus yang dipetakan ke kolom.
  8. Board tidak menampilkan work item terarsip.
  9. Bila Kanplan aktif (BRD-043), status di kolom Kanban backlog dipindah ke halaman Backlog dan tidak tampil di board.
  10. ⚠ Kolom dapat di-collapse/expand lewat aksi di header kolom (ikon panah atau menu kolom), di company-managed maupun team-managed; collapse adalah preferensi tampilan dan tidak mengubah konfigurasi board.
  11. Kolom terciut tampil sebagai strip sempit dengan nama kolom vertikal dan jumlah work item tetap terlihat.
  12. ⚠ Drag kartu ke kolom terciut tetap menjalankan transisi ke status kolom tersebut (BRD-016) tanpa membuka kolom.
  13. State collapse kolom disimpan per pengguna per board dan tidak memengaruhi tampilan pengguna lain.
  14. Klik kartu membuka detail work item (BRD-018) sebagai side panel atau modal.
  15. ⚠ Board dapat dilihat di aplikasi/tampilan mobile (lebar layar < 768 px): kolom tampil satu per layar dengan swipe horizontal, kartu dapat dipindah antar kolom lewat aksi "Move to", dan konfigurasi board tidak tersedia di mobile.
- **Team-managed vs company-managed:** Komponen UI sama; perbedaan konfigurasi dijelaskan di BRD-048.
- **Catatan & ketergantungan:** Realtime dan performa di [27 NFR](27-nfr-arsitektur-platform.md); Insights di [16 Laporan](16-laporan-insight.md). Komponen board di SPA (`frontend/`) harus dapat dipakai ulang oleh board sprint aktif (BRD-014) dan board bisnis (BRD-050).

### BRD-014 Active Sprint Board

- **Deskripsi:** Tampilan board untuk sprint yang sedang berjalan berupa kolom status dan kartu work item, lengkap dengan header sprint dan kontrol stand-up.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Anggota tim (melihat dan memindahkan kartu), Scrum Master/Team Lead (Complete sprint memerlukan Manage Sprints).
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan nama sprint, sisa hari (dihitung dengan working days board, BRD-040), tooltip tanggal start/end, sprint goal, tombol "Complete sprint" (hanya untuk pemegang Manage Sprints, modul 11), dan menu "...".
  2. Kolom mengikuti konfigurasi board dengan jumlah work item per kolom dan indikator column constraint (BRD-028).
  3. Kartu menampilkan field sesuai card layout mode Active sprints (BRD-037).
  4. Bila tidak ada sprint aktif, board menampilkan empty state "Tidak ada sprint aktif" dengan tautan ke Backlog untuk memulai sprint.
  5. Swimlane/Group by (BRD-031, BRD-032) tersedia dengan collapse per swimlane.
  6. Toolbar menyediakan search, quick filter, dan filter assignee/epic/type/label (BRD-034).
  7. Saat parallel sprints aktif, pemilih sprint di header memungkinkan memilih satu atau beberapa sprint aktif yang ditampilkan.
  8. Toggle Insights membuka panel insight sprint (modul 16).
  9. Tombol Stand-up memulai mode stand-up harian (BRD-036).
  10. Menu "..." berisi Edit sprint, Configure board/Board settings, dan Manage workflow (team-managed, untuk Administrator project).
  11. ⚠ Ikon Automation di header board membuka daftar aturan otomasi project dan pembuat aturan cepat (modul 20) bagi pengguna yang berwenang.
  12. Board diperbarui otomatis secara realtime saat work item berubah (sama dengan BRD-013 butir 6).
  13. Board mendukung scroll horizontal untuk banyak kolom dan tampilan responsif (sama dengan BRD-013 butir 4 dan 15).
- **Team-managed vs company-managed:** Team-managed menampilkan board sprint hanya bila toggle Sprints aktif (BRD-049); tanpa Sprints, board berperilaku Kanban (BRD-013).
- **Catatan & ketergantungan:** Siklus hidup sprint (start, complete, parallel sprints) di [11 Backlog & Sprint](11-backlog-sprint.md); sisa hari memakai working days BRD-040.

### BRD-015 Board Card Display

- **Deskripsi:** Spesifikasi informasi default yang ditampilkan pada kartu work item di board dan perilaku visualnya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat board; konten kartu mengikuti card layout (BRD-037).
- **Aturan perilaku (acceptance criteria):**
  1. Elemen default kartu: summary (maksimal 3 baris lalu dipotong dengan elipsis dan tooltip lengkap), key work item, ikon work type, ikon priority, avatar assignee (atau ikon Unassigned), dan estimate bila estimation statistic board bukan Issue Count.
  2. Lozenge parent/epic berwarna sesuai warna epic ditampilkan bila work item memiliki parent di level Epic.
  3. Indikator flag (ikon bendera dan latar berwarna, BRD-021) ditampilkan untuk work item yang di-flag.
  4. Indikator days in column (titik) ditampilkan bila diaktifkan (BRD-039).
  5. Field tambahan dari card layout ditampilkan di bawah summary, maksimal 3 field per mode.
  6. Strip warna di sisi kiri kartu mengikuti card colors board (company-managed, BRD-038).
  7. ⚠ Team-managed: kartu menampilkan labels, due date, jumlah child, dan info development sesuai toggle card layout team-managed (BRD-037 butir 9).
  8. Hover (atau fokus keyboard) pada kartu menampilkan tombol menu "..." untuk aksi kartu (BRD-020).
  9. Key work item yang sudah memiliki resolution ditampilkan dengan gaya coret, mengikuti aturan modul 06.
  10. ⚠ Board modern mendukung inline edit langsung di kartu untuk summary, assignee, dan estimate (klik field pada kartu), memakai endpoint update work item yang sama dengan issue view (ISS-010); field tanpa Edit Issues tidak dapat diedit dari kartu.
- **Team-managed vs company-managed:** Company-managed memakai card colors dan field tambahan; team-managed memakai toggle elemen kartu.
- **Catatan & ketergantungan:** Field dan ikon work type di [05 Konfigurasi](05-konfigurasi-work-item-skema.md). Endpoint kartu mengembalikan hanya field yang diperlukan card layout untuk menjaga ukuran payload.

### BRD-016 Board Drag-and-Drop Transitions

- **Alias inventaris:** Drag and drop transitions
- **Deskripsi:** Memindahkan kartu antar kolom untuk menjalankan workflow transition ke status yang dipetakan, termasuk menangani transition screen, condition, validator, dan perubahan field antar swimlane/grup.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Anggota tim dengan Transition Issues; Schedule Issues untuk mengubah rank; Edit Issues untuk perubahan field akibat drag antar swimlane.
- **Aturan perilaku (acceptance criteria):**
  1. Drag kartu ke kolom lain menjalankan workflow transition ke status yang dipetakan ke kolom tujuan memakai mesin transisi modul 09.
  2. Saat drag dimulai, kolom dengan transisi valid di-highlight dan kolom tanpa transisi valid ditampilkan nonaktif (redup, kursor not-allowed).
  3. Bila kolom tujuan memuat beberapa status, kolom terbagi menjadi drop zone per status yang dapat dicapai, masing-masing berlabel nama status.
  4. Bila ada beberapa transition ke status tujuan yang sama, muncul dialog untuk memilih nama transition.
  5. Transition yang memiliki screen menampilkan transition screen dalam dialog (mis. Resolution, Comment) sebelum perpindahan selesai; menekan Cancel mengembalikan kartu ke posisi semula.
  6. Condition yang tidak terpenuhi menyembunyikan transition (kolom/zone tampil nonaktif); validator yang gagal menolak transisi dengan pesan galat dari validator dan kartu kembali ke posisi semula.
  7. Post function transition dijalankan (mis. set resolution, assign) dan hasilnya langsung tercermin pada kartu.
  8. Drag di dalam kolom yang sama hanya mengubah rank (BRD-017) dan tidak menjalankan transisi.
  9. Drop di posisi tertentu pada kolom tujuan sekaligus mengatur rank work item relatif terhadap tetangganya (memerlukan Schedule Issues; tanpa izin itu kartu ditaruh sesuai rank yang ada).
  10. Sub-task dapat di-drag sebagai kartu sendiri maupun di dalam swimlane Stories.
  11. Saat semua sub-task sebuah parent selesai karena drag, muncul prompt "Semua sub-task selesai. Pindahkan <parent> ke <status Done>?" dengan pilihan Ya/Nanti (BRD-022).
  12. Drag memerlukan Transition Issues; pengguna tanpa izin tidak dapat mengangkat kartu (handle drag disembunyikan) dan API menolak dengan HTTP 403.
  13. Field wajib pada transition screen mengikuti permission field; field yang tidak dapat diedit pengguna tetapi wajib membuat transisi gagal dengan pesan "Anda tidak memiliki izin mengisi <field> yang wajib untuk transisi ini".
  14. ⚠ Multi-select kartu (Ctrl/Cmd+klik) lalu drag memindahkan semua kartu terpilih sekaligus; setiap kartu ditransisikan sendiri-sendiri dan kegagalan sebagian ditampilkan dalam ringkasan "N dari M berhasil".
  15. Team-managed: drag menjalankan transition; bila workflow membatasi transition ("allow all statuses" dimatikan), zone yang tidak valid ditandai nonaktif.
  16. ⚠ Team-managed: drag antar lane saat Group by Assignee/Epic mengubah assignee/parent kartu sesuai lane tujuan.
  17. Drag antar swimlane/grup Assignee atau Epic di company-managed sekaligus mengubah field assignee atau parent (memerlukan Assign Issues atau Edit Issues).
  18. ⚠ Drag antar swimlane Queries, Stories, atau Projects di company-managed tidak diizinkan; drop zone lintas lane tersebut dinonaktifkan dan hanya perpindahan kolom dalam lane yang sama yang diterima.
  19. Pembaruan optimistis: kartu langsung berpindah di UI, lalu dikembalikan ke posisi semula beserta flag galat bila server menolak transisi.
  20. Work item yang berpindah ke status yang tidak dipetakan ke kolom mana pun (mis. lewat issue view) hilang dari board, dan pengguna yang melakukannya dari board melihat flag "Work item dipindah ke status yang tidak ditampilkan di board".
  21. Bila transisi berhasil tetapi work item tidak lagi cocok dengan filter board atau quick filter aktif, kartu menghilang dengan flag penjelasan.
  22. Drag dapat dilakukan dengan keyboard (fokus kartu, Space untuk mengangkat, panah untuk memindah, Space untuk melepas, Esc untuk batal) demi aksesibilitas.
  23. Transisi dari board menulis history work item dan menerbitkan event domain yang sama dengan transisi dari issue view (modul 28).
  24. Konflik konkuren (work item sudah ditransisikan pengguna lain) menghasilkan galat "Work item telah berubah; board dimuat ulang" dan kartu disegarkan ke status terbaru.
- **Team-managed vs company-managed:** Company-managed memetakan banyak status ke satu kolom dengan drop zone; team-managed memetakan kolom ke status secara langsung dan mendukung perubahan field antar lane Group by.
- **Catatan & ketergantungan:** Engine workflow di [09 Workflow](09-workflow-status-transisi.md); izin di [03 Izin](03-izin-peran-keamanan.md). Endpoint transisi board wajib memakai use case yang sama dengan issue view.

### BRD-017 Board ranking (Rank)

- **Deskripsi:** Pengurutan prioritas kartu memakai field Rank global (string leksikografis mirip LexoRank) yang menjadi dasar urutan di board, backlog Scrum, dan Kanplan.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Anggota tim dengan Schedule Issues (company-managed); Board administrator memperbaiki filter agar ranking aktif.
- **Aturan perilaku (acceptance criteria):**
  1. Field sistem Rank menentukan urutan; board filter harus berakhiran `ORDER BY Rank ASC` agar ranking aktif.
  2. Board settings > General menampilkan status ranking dan tombol "Add Rank" yang menambahkan `ORDER BY Rank ASC` ke filter (hanya owner/editor filter).
  3. Drag vertikal dalam kolom mengubah rank; aksi kartu "Move to top" dan "Move to bottom" menempatkan kartu di atas/bawah kolom (BRD-020).
  4. Rank bersifat global: perubahan di satu board ikut memengaruhi urutan di board lain yang memakai Rank.
  5. Mengubah rank memerlukan Schedule Issues (company-managed); pengguna tanpa izin tidak dapat melakukan drag vertikal dan aksi Move to top/bottom tidak tampil.
  6. Rank sub-task dibatasi di dalam parent-nya: sub-task hanya dapat diurutkan relatif terhadap sub-task lain dengan parent yang sama.
  7. Bila filter diurutkan dengan field lain, drag vertikal dinonaktifkan dan board mengikuti urutan filter dengan tooltip penjelasan.
  8. Re-rank saat sebagian work item tersembunyi oleh quick filter memposisikan work item relatif terhadap tetangga yang terlihat (rank di antara rank kartu atas dan bawah yang terlihat).
  9. Implementasi rank memakai string leksikografis dengan bucket dan rebalancing latar belakang ketika jarak antar rank habis; rebalancing tidak mengubah urutan relatif dan tidak terlihat oleh pengguna.
- **Team-managed vs company-managed:** Team-managed memakai field Rank yang sama; izin rank mengikuti role project (Member dapat me-rank).
- **Catatan & ketergantungan:** Field Rank di [06 Work Item Inti](06-work-item-inti.md); API rank di BRD-052; model data rank di [28 Model Data](28-model-data-inti.md).

### BRD-018 Issue Detail Panel on Board and Backlog

- **Alias inventaris:** Issue detail view on board
- **Deskripsi:** Menampilkan issue view sebagai panel samping (atau modal) ketika work item dipilih di board atau backlog, sehingga pengguna dapat membaca dan mengedit tanpa meninggalkan halaman.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat work item; pengeditan mengikuti Edit Issues, Transition Issues, dan Add Comments.
- **Aturan perilaku (acceptance criteria):**
  1. Klik kartu atau baris backlog membuka panel kanan berisi issue view (modul 06); panel ditutup via tombol close atau tombol Esc.
  2. ⚠ Pengguna dapat memilih mode tampilan detail board: side panel (default) atau modal; pilihan disimpan per pengguna.
  3. Lebar panel dapat diubah dengan drag tepi kiri (minimal 360 px, maksimal 70% lebar layar) dan lebar terakhir disimpan per pengguna.
  4. Panel menyediakan tombol untuk membuka work item full page (`/browse/KEY`) dan opsi membuka di tab baru (Ctrl/Cmd+klik pada key).
  5. Tombol "Copy link" di panel menyalin URL work item.
  6. Saat panel terbuka, shortcut `j`/`k` memilih work item berikutnya/sebelumnya di board atau backlog dan panel memuat work item tersebut.
  7. Shortcut `t` menampilkan/menyembunyikan panel detail untuk work item terpilih.
  8. Semua field dapat diedit inline sesuai permission; field tanpa izin tampil read-only.
  9. Transisi status (dropdown status) dan penambahan komentar dapat dilakukan langsung dari panel.
  10. URL board memuat parameter `selectedIssue=<KEY>` selama panel terbuka sehingga dapat dijadikan deep link; menutup panel menghapus parameter tersebut.
  11. Perubahan di panel (summary, assignee, status, flag, estimate) langsung tercermin pada kartu tanpa reload; perubahan status dapat memindahkan kartu ke kolom lain.
  12. Membuka `selectedIssue` untuk work item yang tidak ada di board (mis. tersaring) tetap membuka panel dengan banner "Work item ini tidak ditampilkan di board karena filter".
  13. Panel tidak dimuat bila pengguna kehilangan akses; panel menampilkan "Anda tidak memiliki akses ke work item ini".
- **Team-managed vs company-managed:** Tidak ada perbedaan; isi panel mengikuti layout work type (team-managed) atau screen (company-managed) sesuai modul 05.
- **Catatan & ketergantungan:** Isi dan layout issue view mengikuti ISS-007 (Issue View Layout) dan mode tampilan panel/modal/full page mengikuti ISS-008 (Issue View Modes) di [06 Work Item Inti](06-work-item-inti.md); pengeditan inline mengikuti ISS-010; shortcut global di [21 Navigasi](21-navigasi-ux-preferensi.md). Komponen issue view SPA harus dapat di-embed di panel.

### BRD-019 Inline create on board

- **Alias inventaris:** Inline Create on Board
- **Deskripsi:** Membuat work item baru langsung dari kolom board dengan mengetik summary, tanpa membuka Create dialog lengkap.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues pada project board.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Company-managed: tombol "+ Create issue" tampil di bawah kolom pertama board.
  2. ⚠ Team-managed: tombol "+ Create" tampil di kolom pertama dan muncul saat hover di kolom lain, sehingga work item dapat dibuat langsung di kolom mana pun.
  3. Pengguna memilih work type (default work type project), mengetik summary, lalu menekan Enter untuk membuat; Esc membatalkan.
  4. Setelah Enter, input tetap terbuka untuk membuat work item berikutnya secara beruntun.
  5. Di board sprint aktif, work item baru otomatis masuk sprint aktif yang sedang ditampilkan (bila pengguna memiliki Manage Sprints atau izin yang setara untuk menambah ke sprint aktif, modul 11).
  6. Status work item baru mengikuti kolom tempat ia dibuat (team-managed) atau initial status workflow (company-managed).
  7. Work item baru mewarisi konteks swimlane/grup tempat ia dibuat (mis. assignee di lane Assignee, parent di lane Epic).
  8. Bila create screen memuat field wajib selain Summary, sistem membuka Create dialog lengkap dengan summary dan work type yang sudah diketik (modul 06).
  9. Bila work item baru tidak cocok dengan filter board atau quick filter aktif sehingga tidak terlihat, flag "Work item <KEY> dibuat tetapi tidak terlihat karena filter" ditampilkan dengan tautan ke work item.
  10. Tombol Create global (modul 06) tetap tersedia dan default ke project board terpilih.
  11. Board lintas project memakai project location board sebagai default (BRD-003 butir 13).
  12. Summary kosong tidak dapat disubmit; tombol Enter tidak melakukan apa pun dan input menampilkan border galat.
- **Team-managed vs company-managed:** Company-managed: inline create di kolom pertama dengan initial status; team-managed: inline create di kolom mana pun dengan status kolom tersebut.
- **Catatan & ketergantungan:** Validasi create mengikuti ISS-003 di [06 Work Item Inti](06-work-item-inti.md).

### BRD-020 Card context menu & card actions

- **Alias inventaris:** Board Card Actions Menu
- **Deskripsi:** Menu aksi cepat pada kartu (klik kanan atau tombol "...") untuk operasi work item yang umum tanpa membuka work item.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Anggota tim; setiap aksi tampil hanya bila pengguna memiliki izin terkait (Schedule, Transition, Edit, Assign, Delete, Archive).
- **Aturan perilaku (acceptance criteria):**
  1. Menu dibuka lewat tombol "..." pada kartu atau klik kanan pada kartu.
  2. Move: "Move to top" dan "Move to bottom" kolom (mengubah rank, memerlukan Schedule Issues).
  3. Move: pindah ke kolom/status lain lewat submenu daftar kolom, yang menjalankan transisi seperti drag (BRD-016).
  4. Add flag / Remove flag (BRD-021).
  5. Change parent / Add parent membuka pemilih parent yang valid sesuai hierarki work type.
  6. Assign membuka pemilih assignee; "Assign to me" langsung menugaskan ke pengguna saat ini (memerlukan Assign Issues dan Assignable User).
  7. Copy issue link menyalin URL work item; Copy issue key menyalin key.
  8. ⚠ Add label membuka input label dengan autocomplete label yang sudah ada.
  9. ⚠ Link issue membuka dialog tautan work item (modul 06) di board Scrum maupun Kanban.
  10. Split issue (company-managed, board Scrum) membuka dialog pemecahan work item menjadi beberapa work item (modul 11).
  11. ⚠ Remove from sprint memindahkan work item ke backlog (board sprint aktif, memerlukan Manage Sprints atau Schedule Issues).
  12. Archive tampil hanya di plan Premium/Enterprise dengan izin archive (modul 06).
  13. Delete tampil bagi pemegang Delete Issues dan meminta konfirmasi.
  14. Aksi yang tidak diizinkan tidak ditampilkan sama sekali (bukan sekadar dinonaktifkan).
  15. ⚠ Bila beberapa kartu terpilih (Ctrl/Cmd+klik), menu menampilkan aksi bulk yang didukung (Move, Flag, Assign, Delete) dan menerapkannya pada semua kartu terpilih.
  16. Setiap aksi menampilkan flag hasil ("Ditandai", "Ditugaskan ke ...") dan kesalahan server dikembalikan sebagai flag galat tanpa mengubah kartu.
  17. Menu dapat dioperasikan dengan keyboard (Shift+F10 atau tombol menu konteks pada kartu terfokus).
- **Team-managed vs company-managed:** Split issue hanya company-managed; aksi lainnya sama.
- **Catatan & ketergantungan:** Operasi work item di [06 Work Item Inti](06-work-item-inti.md); split dan sprint di [11 Backlog & Sprint](11-backlog-sprint.md).

### BRD-021 Flagging Issues

- **Alias inventaris:** Flagging
- **Deskripsi:** Menandai work item sebagai impediment dengan flag yang menonjol di board, backlog, issue view, dan Insights.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Anggota tim dengan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. "Add flag" tersedia dari menu kartu, menu baris backlog, issue view ("..." > Add flag), dan dialog operasi keyboard (`.`).
  2. Opsi "Add flag and comment" membuka dialog opsional untuk menulis komentar alasan; komentar disimpan sebagai komentar biasa dengan penanda flag.
  3. "Remove flag" tersedia dari lokasi yang sama, dengan opsi komentar.
  4. Flag disimpan di field sistem "Flagged" (checkbox dengan satu opsi "Impediment").
  5. Kartu/baris yang di-flag menampilkan ikon bendera dan latar berwarna (kuning/oranye) yang tetap kontras di tema terang dan gelap.
  6. Bulk flag/unflag tersedia pada multi-select di backlog dan board.
  7. Work item yang di-flag muncul di insight "Issues for attention" (modul 16).
  8. JQL `Flagged = Impediment` dan `flagged is not EMPTY` dapat dipakai di quick filter, swimlane, dan pencarian.
  9. Flag dapat dipakai sebagai kondisi atau aksi di otomasi (modul 20).
  10. Mengubah flag memerlukan Edit Issues; pengguna tanpa izin tidak melihat aksi flag.
  11. Flag terlihat di issue view (banner "Flagged") selain di board dan backlog.
  12. Penambahan/penghapusan flag tercatat di history work item dan memicu notifikasi Issue updated (modul 19).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Field Flagged di [06 Work Item Inti](06-work-item-inti.md) dan [05 Konfigurasi](05-konfigurasi-work-item-skema.md).

### BRD-022 Sub-tasks and hierarchy on board

- **Alias inventaris:** Sub-tasks on Board and Backlog
- **Deskripsi:** Aturan tampilan dan perilaku sub-task, parent, dan epic di board dan backlog, termasuk keterikatan sprint dan rank.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Anggota tim.
- **Aturan perilaku (acceptance criteria):**
  1. Sub-task tidak tampil sebagai baris tersendiri di backlog; backlog menampilkan jumlah sub-task pada baris parent.
  2. Di board, sub-task tampil sebagai kartu terpisah di kolom statusnya.
  3. Swimlane Stories (BRD-031) atau Group by Subtask (BRD-032) menampilkan parent bersama sub-task-nya dalam satu lane.
  4. Sub-task selalu mengikuti sprint parent; field Sprint sub-task tidak dapat diubah terpisah dan API menolak perubahan dengan pesan galat.
  5. Prompt transisi parent muncul ketika semua sub-task selesai (BRD-016 butir 11).
  6. Sprint tidak dapat diselesaikan bila parent berstatus done memiliki sub-task yang belum selesai; dialog Complete sprint menyebut work item yang bermasalah (modul 11).
  7. ⚠ Estimasi Story points sub-task tidak dijumlahkan ke total sprint; time tracking (Remaining Estimate) sub-task dijumlahkan ke parent sesuai pengaturan rollup modul 14.
  8. Kartu parent menampilkan indikator jumlah dan progress child (mis. "2/5").
  9. Lozenge parent/epic pada kartu memakai warna epic.
  10. Rank sub-task terbatas di dalam parent (BRD-017 butir 6).
  11. Column constraint "Issue Count excluding Sub-tasks" mengecualikan sub-task dari hitungan WIP (BRD-028).
  12. Epic tampil sebagai kartu di board Kanban kecuali Epics panel di Kanplan aktif (BRD-044).
  13. Parent yang tidak cocok dengan filter board tetapi memiliki sub-task yang cocok ditampilkan sebagai header lane Stories bertanda "di luar filter board".
  14. Sub-task tidak dapat dipindah ke lane parent lain lewat drag; perubahan parent hanya lewat aksi Change parent.
- **Team-managed vs company-managed:** Team-managed memakai Group by Subtask sebagai padanan swimlane Stories.
- **Catatan & ketergantungan:** Hierarki work type di [05 Konfigurasi](05-konfigurasi-work-item-skema.md); sprint di [11 Backlog & Sprint](11-backlog-sprint.md); rollup estimasi di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md).

## Kolom, status & batas WIP

### BRD-023 Board Settings - Columns & Constraints

- **Deskripsi:** Halaman Board settings > Columns sebagai satu tempat untuk mengatur kolom, pemetaan status, unmapped statuses, column constraint, dan batas min/max per kolom. Fitur ini adalah spesifikasi halaman; detail perilaku tiap kontrol ada di BRD-024 s.d. BRD-030.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board administrator (ubah); pengguna lain read-only.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan kolom board secara horizontal; kolom dapat ditambah ("Add column"), di-rename inline, dihapus (ikon tempat sampah), dan diurutkan ulang dengan drag header kolom.
  2. Status dipetakan ke kolom dengan drag lozenge status; satu status hanya boleh berada di satu kolom per board.
  3. Area "Unmapped statuses" di sisi kanan menampung status yang tidak dipetakan; work item berstatus unmapped tidak tampil di board maupun backlog board tersebut.
  4. Status baru yang ditambahkan ke workflow project di board company-managed otomatis masuk ke Unmapped statuses.
  5. Setiap lozenge status menampilkan jumlah work item (hasil filter board) yang berstatus tersebut.
  6. Tombol "Add status" tersedia bila project memakai Simplified Workflow (BRD-027, mulai Fase 3) atau pengguna adalah Admin Situs (membuat status global yang lalu harus ditambahkan ke workflow lewat modul 09); di Fase 1 project dari template memakai workflow template statis read-only (BRD-001 butir 2) sehingga tombol ini tidak tampil bagi Board administrator non-Admin Situs.
  7. Checkbox "Set resolution" per status hanya tampil pada Simplified Workflow (mulai Fase 3); di Fase 1 checkbox tidak tampil.
  8. Dropdown Column constraint menawarkan None, Issue Count, dan Issue Count excluding Sub-tasks; board Scrum menambahkan opsi berbasis estimasi (BRD-028).
  9. Setiap kolom memiliki input Min dan Max; header kolom di board berwarna saat jumlah di bawah minimum atau melebihi maksimum (BRD-029).
  10. Kolom paling kanan dianggap kolom Done untuk penyelesaian sprint (work item di kolom ini dianggap selesai saat Complete sprint) dan untuk backlog (BRD-030).
  11. Pada board Kanban, kolom khusus "Kanban backlog" paling kiri dapat diisi status untuk mengaktifkan Kanplan (BRD-043).
  12. Team-managed: kolom dikelola langsung dari board (tambah kolom = buat status, rename, hapus dengan memindahkan work item, set column limit, pindah kolom) dan tidak ada area unmapped statuses (BRD-024 butir 6-9).
  13. Board harus memiliki minimal satu kolom; menghapus kolom terakhir ditolak dengan pesan "Board memerlukan setidaknya satu kolom".
- **Team-managed vs company-managed:** Company-managed: kolom adalah pemetaan tampilan atas status workflow; team-managed: kolom identik dengan status.
- **Catatan & ketergantungan:** Model data: `board_columns (board_id, position, name, min, max)` dan `board_column_statuses (column_id, status_id)`; status dan category di [09 Workflow](09-workflow-status-transisi.md).

### BRD-024 Column management

- **Deskripsi:** Menambah, mengganti nama, mengurutkan ulang, dan menghapus kolom board, dengan semantik berbeda antara company-managed (pemetaan tampilan) dan team-managed (kolom = status).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board administrator (company-managed); Administrator project (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: Board settings > Columns menyediakan "Add column" (nama wajib, 1-255 karakter), rename inline, drag untuk reorder, dan hapus kolom.
  2. Menghapus kolom di company-managed memindahkan status-statusnya ke Unmapped statuses; work item tidak berubah status.
  3. Perubahan kolom company-managed hanya mengubah pemetaan tampilan, bukan workflow; workflow tidak tersentuh.
  4. Kolom paling kanan berperan sebagai kolom Done untuk sub-filter, hide completed, dan Release (BRD-030).
  5. Kolom paling kiri berganti menjadi "Kanban backlog" saat Kanplan aktif (BRD-043).
  6. Team-managed: tombol "+" di ujung kanan board menambah kolom yang sekaligus membuat status baru di workflow project (kategori default In Progress).
  7. Team-managed: rename kolom dari menu kolom mengganti nama status; menghapus kolom meminta pengguna memilih kolom tujuan untuk memindahkan semua work item sebelum status dihapus.
  8. ⚠ Team-managed: Board settings > "Columns and statuses" menampilkan status per kolom, memungkinkan beberapa status dalam satu kolom, dan menandai kolom Done (centang).
  9. Team-managed: menu kolom "..." berisi Rename, Set column limit, Delete, dan Move (kiri/kanan).
  10. ⚠ Batas jumlah kolom per board adalah 30 dan batas status per workflow mengikuti modul 09; melampaui batas menampilkan pesan "Jumlah kolom maksimum tercapai".
  11. Nama kolom duplikat dalam satu board diperbolehkan di company-managed tetapi ditolak di team-managed karena nama kolom = nama status.
- **Team-managed vs company-managed:** Company-managed tidak mengubah workflow; team-managed mengubah workflow (status) secara langsung.
- **Catatan & ketergantungan:** Pembuatan/penghapusan status team-managed memakai use case workflow modul 09 agar validasi (status dipakai transisi) konsisten.

### BRD-025 Column status mapping

- **Deskripsi:** Memetakan satu atau beberapa status workflow ke tiap kolom board company-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board administrator; Admin Situs untuk tautan Manage workflow.
- **Aturan perilaku (acceptance criteria):**
  1. Status dipetakan dengan drag-and-drop lozenge status ke kolom di Board settings > Columns.
  2. Satu kolom dapat berisi banyak status; satu status hanya dapat ada di satu kolom per board (drop status yang sudah dipetakan memindahkannya dari kolom lama).
  3. Lozenge status diberi warna sesuai status category: To Do abu-abu, In Progress biru, Done hijau.
  4. Jumlah work item per status ditampilkan pada lozenge untuk membantu pemetaan.
  5. Board multi-project/multi-workflow menampilkan status dari semua workflow terkait; status yang hanya dipakai sebagian project diberi tooltip daftar project.
  6. Checkbox "Set resolution" per status hanya tampil untuk Simplified Workflow (BRD-027, mulai Fase 3); workflow template statis Fase 1 (BRD-001 butir 2) tidak menampilkannya.
  7. Tautan "Manage workflow" ke editor workflow (modul 09) tampil untuk Admin Situs atau Admin Proyek yang berhak.
  8. Perubahan mapping langsung memengaruhi work item yang tampil di board tanpa reload.
  9. Status baru yang ditambahkan ke workflow otomatis masuk Unmapped statuses (BRD-026).
  10. ⚠ Team-managed: kolom dapat memuat beberapa status lewat halaman Columns and statuses dengan drag status antar kolom.
- **Team-managed vs company-managed:** Company-managed memakai mapping bebas; team-managed memakai halaman Columns and statuses dengan status milik project.
- **Catatan & ketergantungan:** Status category di [09 Workflow](09-workflow-status-transisi.md); hubungan pemetaan kolom-status dengan transisi board didefinisikan di WF-012 (Board transitions & column-status mapping). Modul ini pemilik UI pemetaan kolom; WF-012 pemilik aturan transisi yang dipicu oleh pemetaan tersebut.

### BRD-026 Unmapped statuses

- **Deskripsi:** Area penampung status workflow yang tidak dipetakan ke kolom mana pun, sehingga work item berstatus tersebut tidak tampil di board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Kolom "Unmapped statuses" ditampilkan di sisi kanan halaman Board settings > Columns.
  2. Status yang dikeluarkan dari kolom dan status baru di workflow otomatis masuk ke area ini.
  3. Work item berstatus unmapped tidak terlihat di board tetapi tetap muncul di pencarian dan navigator.
  4. Drag status dari unmapped ke sebuah kolom membuat work item berstatus tersebut tampil lagi.
  5. Setiap status unmapped menampilkan jumlah work item hasil filter board.
  6. Status unmapped tetap dapat dicapai lewat transisi di issue view; work item lalu hilang dari board (BRD-016 butir 20).
  7. Admin dapat sengaja memakai unmapped untuk menyembunyikan status tertentu (mis. "Won't do") dari board.
  8. ⚠ Team-managed: status yang tidak dipetakan ke kolom mana pun ditampilkan sebagai "Unassigned statuses" di halaman Columns and statuses dengan peringatan bahwa work item berstatus itu tidak tampil di board.
  9. Board settings menampilkan banner peringatan "N work item berada di status yang tidak dipetakan" bila jumlah work item unmapped > 0.
- **Team-managed vs company-managed:** Utamanya company-managed; team-managed hanya memiliki "Unassigned statuses" pada board dengan pengelompokan status manual.
- **Catatan & ketergantungan:** Terkait BRD-023 dan BRD-025.

### BRD-027 Simplified Workflow (board-managed statuses)

- **Alias inventaris:** Simplified Workflow
- **Deskripsi:** Mode company-managed di mana board mengelola workflow project sendiri, sehingga board administrator dapat menambah status dan mengatur resolution langsung dari Board settings tanpa workflow designer.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board administrator yang juga Admin Proyek; Admin Situs (mengganti workflow scheme).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Simplify workflow"/"Use Simplified Workflow" tampil di Board settings > Columns hanya bila board mencakup tepat satu project, pengguna adalah Admin Proyek project tersebut, dan project memakai workflow default.
  2. Mengaktifkannya memigrasikan project ke workflow baru dengan global transitions sehingga setiap status dapat berpindah ke status mana pun.
  3. Workflow yang dihasilkan bernama "Software Simplified Workflow for Project <KEY>"; sejak Fase 3 project company-managed baru dari template Kanban/Scrum memakainya secara default, dan workflow template statis read-only milik project yang dibuat di Fase 1 (BRD-001 butir 2) dikonversi otomatis menjadi Simplified Workflow (flag `simplified`) tanpa migrasi status.
  4. Tombol "Add status" di Board settings > Columns langsung membuat status baru (nama + status category) dan menambahkannya ke workflow serta ke kolom yang dipilih.
  5. Checkbox "Set resolution" per status: work item yang masuk status tersebut otomatis diberi resolution (default "Done"), dan keluar dari status itu mengosongkan resolution.
  6. Status yang tidak dipakai work item mana pun dapat dihapus langsung dari Board settings; status yang masih dipakai ditolak dengan pesan berisi jumlah work item.
  7. Wizard migrasi saat mengaktifkan memetakan status lama ke status baru bila ada status lama yang tidak tersedia di workflow baru; migrasi berjalan sebagai job asinkron dengan progress.
  8. Opsi tidak tersedia bila workflow dipakai project lain, atau memiliki condition/validator/post function khusus; tooltip menjelaskan alasannya.
  9. Kembali ke workflow standar dilakukan dengan mengganti workflow scheme project oleh Admin Situs; setelah itu kontrol status (Add status, Set resolution) di board hilang.
  10. Simplified Workflow tidak tersedia untuk board yang mencakup lebih dari satu project (BRD-003 butir 15).
  11. ⚠ Setelah kembali ke workflow biasa, board tidak dapat di-"simplify" ulang secara otomatis kecuali workflow scheme dikembalikan ke workflow default; Admin Situs yang melakukannya melalui modul 09.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed sudah mengelola status dari board secara bawaan (BRD-048).
- **Catatan & ketergantungan:** Pemilik semantik workflow (flag `simplified`, global transitions, aturan resolution, hilangnya mode saat workflow diedit manual) adalah WF-026 (Simplified Workflow (board-managed)) di [09 Workflow](09-workflow-status-transisi.md); fitur ini pemilik UI board (tombol Simplify workflow, Add status, Set resolution per status, wizard migrasi, hapus status) dan keduanya harus memakai aturan yang sama. Job asinkron di [27 NFR](27-nfr-arsitektur-platform.md). Fitur masih aktif di Jira Cloud dan menjadi default template, jadi bukan fitur deprecated. Fase: di Fase 1 template memakai workflow template statis read-only berstatus To Do/In Progress/Done tanpa board-managed status editing (BRD-001 butir 2), bukan Default Workflow Scheme sistem (WF-020); di Fase 3 (P2, bersama WF-026) workflow tersebut dikonversi menjadi "Software Simplified Workflow for Project <KEY>" dan kontrol status dari Board settings mulai tersedia. Sebelum Fase 3 status project tersebut tidak dapat diubah dari board; Admin Situs yang memerlukan status berbeda menyalin workflow lalu menggantinya lewat modul 09 (WF-021, WF-023), dan project tersebut kemudian tidak lagi memenuhi syarat konversi otomatis.

### BRD-028 Column constraints

- **Deskripsi:** Pengaturan jenis perhitungan yang dipakai WIP limit (min/max) di semua kolom board company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown "Column constraint" berada di Board settings > Columns.
  2. Opsi untuk board Kanban: None, Issue Count, Issue Count excluding Sub-tasks.
  3. Board Scrum menambahkan constraint berbasis estimasi (Story points/Original estimate sesuai estimation statistic board); opsi ini tidak tersedia di board Kanban.
  4. Memilih None menyembunyikan input Min/Max di semua kolom dan menghapus indikator WIP dari board (nilai tersimpan dipertahankan bila constraint diaktifkan kembali).
  5. Constraint berlaku seragam untuk semua kolom board; tidak ada constraint berbeda per kolom.
  6. ⚠ Perhitungan memakai work item hasil board filter + Kanban sub-filter + hide completed, dan tidak terpengaruh quick filter atau filter bar yang sedang aktif.
- **Team-managed vs company-managed:** Team-managed tidak memiliki dropdown constraint; column limit selalu menghitung jumlah work item (termasuk sub-task) dan hanya maksimum.
- **Catatan & ketergantungan:** Dipakai oleh BRD-029.

### BRD-029 WIP limits (min/max)

- **Deskripsi:** Batas minimum dan maksimum work in progress per kolom, dengan penanda visual saat batas dilanggar.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (company-managed) dan Administrator project (team-managed) mengatur; anggota tim melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: input Min dan Max per kolom tersedia di Board settings > Columns ketika column constraint bukan None.
  2. Header kolom di board menampilkan "Max: N" dan/atau "Min: N".
  3. Bila jumlah work item melebihi Max, latar kolom (header dan badan) menjadi merah.
  4. Bila jumlah work item di bawah Min, latar kolom menjadi kuning.
  5. WIP limit bersifat soft: tidak memblokir perpindahan kartu, hanya memberi peringatan visual.
  6. Team-managed: menu kolom "..." > "Set column limit" (hanya maksimum); header menampilkan "jumlah/limit" dan berwarna saat terlampaui.
  7. Validasi: nilai harus bilangan bulat non-negatif; Min tidak boleh lebih besar dari Max (pesan "Minimum tidak boleh melebihi maksimum"); kosong berarti tanpa limit.
  8. Jumlah dihitung sesuai column constraint (dengan atau tanpa sub-task, atau total estimasi untuk board Scrum).
  9. Warna penanda memenuhi kontras WCAG AA dan disertai ikon/teks (tidak hanya warna) agar dapat dibaca pengguna buta warna.
- **Team-managed vs company-managed:** Company-managed: Min dan Max; team-managed: hanya Max.
- **Catatan & ketergantungan:** Laporan WIP/CFD di [16 Laporan](16-laporan-insight.md).

### BRD-030 Done column behavior

- **Deskripsi:** Semantik kolom terakhir (Done) board yang memengaruhi resolution, Release, pembersihan board, penyelesaian sprint, dan laporan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (konfigurasi), anggota tim (perilaku).
- **Aturan perilaku (acceptance criteria):**
  1. Kolom paling kanan dianggap kolom selesai untuk tombol Release (BRD-047), opsi hide completed (BRD-046), dan penyelesaian sprint (modul 11).
  2. Status di kolom terakhir sebaiknya berkategori Done; status category (bukan posisi kolom) menentukan perhitungan "resolved/done" di laporan dan JQL `statusCategory`.
  3. Resolution tidak otomatis terisi saat kartu dipindah ke kolom Done kecuali lewat transition screen, post function, atau "Set resolution" di Simplified Workflow.
  4. ⚠ Board settings > Columns menampilkan peringatan konfigurasi bila status berkategori Done dipetakan ke kolom selain kolom terakhir, atau kolom terakhir berisi status non-Done.
  5. ⚠ Team-managed: kolom yang ditandai Done (centang) menentukan status category Done dan resolution otomatis diisi saat work item masuk kolom tersebut.
  6. ⚠ Kartu di kolom Done menampilkan key work item dengan gaya coret bila work item memiliki resolution (konsisten dengan BRD-015 butir 9).
- **Team-managed vs company-managed:** Company-managed: Done = posisi kolom terakhir; team-managed: Done = kolom bertanda centang.
- **Catatan & ketergantungan:** Status category dan resolution di [09 Workflow](09-workflow-status-transisi.md) dan [06 Work Item Inti](06-work-item-inti.md).

## Swimlane, pengelompokan & penyaringan

### BRD-031 Swimlanes (company-managed)

- **Alias inventaris:** Board Swimlanes; Board Settings - Swimlanes
- **Deskripsi:** Pengelompokan horizontal kartu di board berdasarkan strategi yang dikonfigurasi board administrator (Queries, Stories, Assignees, Epics, Projects), termasuk interaksi collapse/expand dan perilaku drag antar lane.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (konfigurasi di Board settings > Swimlanes); anggota tim (melihat, collapse, drag).
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > Swimlanes menyediakan pilihan "Base swimlanes on": Queries, Stories, Assignees, Epics, Projects, No swimlanes (satu strategi aktif per board).
  2. Mode Queries: board administrator menambah swimlane dengan Name (wajib), JQL (wajib), dan Description (opsional).
  3. Swimlane query dapat diedit, diurutkan ulang dengan drag, dan dihapus.
  4. JQL swimlane divalidasi saat simpan; JQL tidak valid ditolak dengan pesan galat parser JQL (modul 12) dan swimlane tidak tersimpan.
  5. Setiap work item masuk ke swimlane query pertama (dari atas) yang JQL-nya cocok; work item tidak pernah muncul di dua lane.
  6. Mode Queries selalu memiliki swimlane sisa "Everything Else" untuk work item yang tidak cocok query mana pun; lane ini tidak dapat dihapus tetapi posisinya dapat diubah.
  7. ⚠ Board Kanban baru memiliki swimlane default "Expedite" (`priority = Blocker`) di atas "Everything Else"; board Scrum baru memakai No swimlanes.
  8. Mode Stories: setiap parent yang memiliki sub-task di board menjadi satu swimlane; work item lain dikumpulkan di lane "Other issues" yang dapat ditaruh di atas atau di bawah (opsi di Board settings).
  9. Header lane Stories menampilkan key, summary, dan status parent, beserta jumlah sub-task.
  10. Swimlane Stories: saat semua sub-task selesai, muncul prompt untuk memindahkan parent ke Done (sama dengan BRD-016 butir 11), dan parent dapat ditransisikan dari menu header lane.
  11. ⚠ Mode Assignees: satu swimlane per assignee (urut nama); lane "Unassigned" dapat ditempatkan di atas atau di bawah (default di bawah) lewat opsi Board settings.
  12. Mode Epics: satu swimlane per epic berdasarkan field Parent (setelah migrasi Epic Link ke Parent); work item tanpa epic masuk lane "Issues without epics".
  13. Mode Projects: satu swimlane per project, berguna untuk board lintas project (BRD-003).
  14. Mode No swimlanes: semua kartu tampil dalam satu area tanpa header lane.
  15. Header setiap swimlane menampilkan jumlah work item di lane tersebut.
  16. Setiap swimlane dapat di-collapse/expand dengan klik header; shortcut `-` melakukan collapse semua/expand semua.
  17. State collapse disimpan per pengguna per board dan tidak memengaruhi pengguna lain.
  18. Drag kartu antar swimlane Assignees atau Epics mengubah assignee atau parent kartu (BRD-016 butir 17).
  19. ⚠ Drag kartu antar swimlane Queries, Stories, atau Projects tidak diizinkan (BRD-016 butir 18).
  20. ⚠ Team-managed tidak memiliki swimlane terkonfigurasi; padanannya Group by None, Assignee, Epic/Parent, Subtask (BRD-032).
  21. Team-managed: pilihan Group by diingat per pengguna (BRD-032 butir 3).
  22. Swimlane diterapkan di board Kanban dan board sprint aktif; backlog tidak memakai swimlane.
  23. Swimlane yang tidak berisi kartu setelah filter bar diterapkan tetap ditampilkan dengan header berjumlah 0 di mode Queries, dan disembunyikan di mode Assignees/Epics/Stories.
- **Team-managed vs company-managed:** Company-managed: strategi swimlane per board oleh admin; team-managed: Group by per pengguna.
- **Catatan & ketergantungan:** JQL di [12 JQL](12-jql-bahasa-kueri.md); field Parent di [06 Work Item Inti](06-work-item-inti.md). Model data: `board_swimlanes (board_id, strategy, position, name, jql, description)` dan preferensi per pengguna `board_user_prefs`.

### BRD-032 Group by

- **Deskripsi:** Pengelompokan kartu lewat menu "Group by" di toolbar board. Asalnya fitur team-managed, juga tersedia sebagai override per pengguna di company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Anggota tim (preferensi pribadi); perubahan field akibat drag mengikuti Edit/Assign Issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Menu "Group by" di toolbar/View settings menawarkan None, Assignee, Epic/Parent, Subtask, dan Priority; opsi Queries tidak tersedia di Group by.
  2. ⚠ Company-managed: Group by per pengguna menimpa sementara swimlane yang dikonfigurasi admin; opsi "Board default" mengembalikan swimlane admin.
  3. Pilihan grup disimpan per pengguna per board.
  4. Setiap grup dapat di-collapse/expand.
  5. Drag antar grup Assignee mengubah assignee; drag antar grup Epic/Parent mengubah parent.
  6. ⚠ Bila opsi Group by Priority dipakai (butir 1), drag antar grup Priority mengubah priority work item (memerlukan Edit Issues).
  7. Grup "Unassigned"/"No epic" menampung work item tanpa nilai pada atribut pengelompokan.
  8. Grup Subtask menampilkan parent bersama sub-task-nya (padanan swimlane Stories).
  9. Header grup menampilkan jumlah work item dan, untuk grup Assignee, avatar pengguna.
- **Team-managed vs company-managed:** Team-managed memakai Group by sebagai satu-satunya mekanisme swimlane; company-managed memakainya sebagai override per pengguna.
- **Catatan & ketergantungan:** Terkait BRD-031 dan BRD-035.

### BRD-033 Quick filters

- **Alias inventaris:** Board Settings - Quick Filters
- **Deskripsi:** Tombol filter cepat berbasis JQL di atas board dan backlog untuk mempersempit kartu yang tampil tanpa mengubah filter board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (konfigurasi di Board settings > Quick Filters); anggota tim (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Board company-managed baru memiliki quick filter default "Only My Issues" (`assignee = currentUser()`) dan "Recently Updated" (`updatedDate >= -1d`).
  2. Board administrator menambah quick filter dengan Name (wajib), JQL Query (wajib), dan Description (tooltip).
  3. Quick filter dapat diedit, dihapus, dan diurutkan ulang dengan drag.
  4. JQL divalidasi saat menyimpan; JQL tidak valid ditolak dengan pesan galat parser.
  5. Quick filter tampil sebagai tombol toggle di toolbar board dan backlog company-managed.
  6. ⚠ Bila jumlah quick filter lebih dari 5 (atau tidak muat dalam lebar toolbar), sisanya dikelompokkan ke dropdown "Quick filters" dengan penanda jumlah yang aktif.
  7. Beberapa quick filter aktif digabung dengan AND, lalu digabung dengan filter board dan filter bar.
  8. Quick filter juga tersedia di Kanplan backlog (BRD-043) dan di opsi "Refine report" CFD/Control chart (modul 16).
  9. Status quick filter yang aktif disimpan di URL (`quickFilter=<id>`) dan dipertahankan per sesi saat pengguna kembali ke board.
  10. ⚠ Team-managed: tidak ada quick filter terkonfigurasi; padanannya "Custom filters" berbasis JQL di dropdown filter board yang dibuat oleh Administrator project.
  11. Quick filter yang JQL-nya menjadi tidak valid (mis. field dihapus) tetap tampil dengan ikon galat dan tidak dapat diaktifkan sampai diperbaiki.
- **Team-managed vs company-managed:** Company-managed: quick filter per board; team-managed: custom filters.
- **Catatan & ketergantungan:** JQL di [12 JQL](12-jql-bahasa-kueri.md); refine report di [16 Laporan](16-laporan-insight.md). Model data: `board_quick_filters (board_id, position, name, jql, description)`.

### BRD-034 Board and Backlog Filters & Search

- **Alias inventaris:** Board filter bar & search
- **Deskripsi:** Toolbar filter interaktif di board dan backlog untuk mencari dan menyaring kartu berdasarkan teks, orang, dan atribut, tanpa mengubah filter board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna yang melihat board/backlog.
- **Aturan perilaku (acceptance criteria):**
  1. Kotak "Search board" mencari di summary, key, dan teks yang tampil di kartu (case-insensitive, substring).
  2. Avatar assignee di toolbar dapat diklik untuk memfilter per orang, mendukung multi-select (OR antar orang), termasuk avatar "Unassigned"; avatar yang tidak muat dikelompokkan ke "+N".
  3. Quick filter board (company-managed) tampil di toolbar dan digabung dengan AND (BRD-033).
  4. Filter Epic dan Version tersedia lewat panel (Epics/Versions panel di backlog, BRD-044, BRD-045) atau dropdown.
  5. Dropdown filter Type, Label, dan Epic/Parent tersedia di team-managed; company-managed menambah Version dan Component.
  6. ⚠ Team-managed: dropdown "Custom filters" berbasis JQL tersedia di semua plan (BRD-033 butir 10).
  7. Tombol "Clear filters" dan indikator jumlah filter aktif ditampilkan bila ada filter aktif.
  8. Filter hanya mengubah tampilan dan tidak pernah mengubah saved filter board.
  9. Filter diterapkan instan di sisi klien untuk data yang sudah dimuat (≤ 200 ms) tanpa reload halaman.
  10. ⚠ State filter tersimpan di URL (dapat dibagikan) dan dipulihkan per sesi saat kembali ke board.
  11. ⚠ Jumlah di header kolom dan indikator WIP tetap menghitung seluruh work item kolom (bukan hasil filter bar); header menampilkan "x dari y" saat filter bar aktif.
  12. Kombinasi filter: kriteria dalam satu dimensi digabung OR, antar dimensi digabung AND.
  13. Filter bar yang sama dipakai di board Kanban, board sprint aktif, backlog Scrum, dan Kanplan.
- **Team-managed vs company-managed:** Company-managed menampilkan quick filter, Version, dan Component; team-managed menampilkan Type, Label, Epic/Parent, dan custom filters.
- **Catatan & ketergantungan:** Backlog Scrum di [11 Backlog & Sprint](11-backlog-sprint.md).

### BRD-035 View settings (board display options)

- **Deskripsi:** Panel preferensi tampilan per pengguna di board/backlog untuk menyalakan/mematikan elemen visual tanpa mengubah konfigurasi board.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Anggota tim (preferensi pribadi).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "View settings" tampil di toolbar board dan backlog.
  2. Di backlog, View settings menyediakan toggle Epics panel dan Versions panel (BRD-044, BRD-045).
  3. View settings menampung pilihan Group by (BRD-032).
  4. ⚠ Toggle tampilan kartu kompak (hanya summary dan key) dan toggle "Show empty columns" tersedia di board.
  5. Semua preferensi disimpan per pengguna per board di server sehingga berlaku di perangkat lain.
  6. Tombol "Reset to default" mengembalikan semua preferensi pengguna ke default board.
  7. ⚠ Board saved views: kombinasi preferensi tampilan (filter, Group by, toggle tampilan) dapat disimpan sebagai view bernama bersama oleh board administrator; perilaku lengkap di BRD-055.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Daftar toggle di Jira Cloud terus berubah; preferensi disimpan di tabel `board_user_prefs` (JSON) agar mudah diperluas. Fase 3 (P2); Group by (BRD-032, Fase 2) sudah dapat dipilih dari toolbar sebelum tombol View settings tersedia. Preferensi per pengguna di fitur ini berbeda dari saved view bersama (BRD-055).

### BRD-036 Stand-ups (Board Stand-up Mode)

- **Deskripsi:** Mode stand-up harian di board yang memandu tim membahas pekerjaan per anggota secara bergiliran; selama giliran seseorang, board otomatis difilter ke work item miliknya.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead (memimpin), anggota tim; mode ini hanya mengubah tampilan lokal fasilitator.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Stand-up"/"Start stand-up" tampil di header board Scrum (sprint aktif) dan Kanban.
  2. Urutan giliran assignee dapat diacak atau mengikuti urutan daftar avatar; fasilitator memakai tombol Next dan Skip.
  3. Selama giliran seseorang, board otomatis memfilter work item milik assignee tersebut (setara filter avatar).
  4. Work item tanpa assignee dibahas sebagai giliran tersendiri "Unassigned".
  5. ⚠ Timer per orang (default 2 menit, dapat diubah) ditampilkan dan berubah warna saat waktu habis, tanpa memaksa pindah giliran.
  6. Mengakhiri stand-up (tombol End atau Esc) mengembalikan filter board ke kondisi sebelum stand-up dimulai.
  7. ⚠ Mode stand-up tersedia di board company-managed dan team-managed dengan perilaku sama.
  8. Mode stand-up tidak menulis data apa pun ke work item; tidak ada notifikasi yang dikirim.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur native yang relatif baru di Jira Cloud; app stand-up Marketplace pihak ketiga di luar cakupan.

## Kustomisasi kartu & pengaturan waktu board

### BRD-037 Card layout

- **Alias inventaris:** Board Settings - Card Layout
- **Deskripsi:** Pengaturan field tambahan yang ditampilkan di kartu board dan baris backlog, serta toggle elemen kartu di team-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (company-managed); Administrator project (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: Board settings > Card layout memiliki konfigurasi terpisah untuk mode Backlog dan mode board (Active sprints untuk Scrum, Kanban board untuk Kanban).
  2. Maksimal 3 field tambahan per mode; tombol Add dinonaktifkan dengan tooltip "Maksimal 3 field" saat batas tercapai.
  3. Field yang didukung: system field dan custom field bertipe teks, select, multi-select, user, number, date, labels, version, dan component.
  4. Field tambahan dapat dihapus dari layout.
  5. Urutan field tambahan dapat diatur dengan drag.
  6. Elemen default kartu selalu tampil dan tidak dihitung dalam batas 3: key, summary, work type, priority, avatar assignee, lozenge parent/epic, estimate (bila ada), dan flag.
  7. ⚠ Checkbox "Days in column" berada di Board settings > Card layout company-managed (BRD-039).
  8. Kartu parent menampilkan indikator dan progress child/sub-task.
  9. ⚠ Team-managed: Board settings > Card layout berisi toggle Work type, Key, Priority, Assignee, Due date, Labels, Parent, Development, Story points, Child issues, dan Days in column.
  10. Due date pada kartu team-managed diberi warna merah bila sudah lewat dan oranye bila jatuh tempo ≤ 2 hari.
  11. Field tambahan yang kosong pada work item tidak ditampilkan di kartu (tanpa placeholder).
  12. Field yang dihapus dari instance otomatis hilang dari card layout dan board settings menampilkan pemberitahuan sekali.
- **Team-managed vs company-managed:** Company-managed: hingga 3 field bebas per mode; team-managed: toggle elemen kartu yang sudah ditentukan.
- **Catatan & ketergantungan:** Field di [05 Konfigurasi](05-konfigurasi-work-item-skema.md); info development di [24 Integrasi DevOps](24-integrasi-devops.md).

### BRD-038 Card colors

- **Alias inventaris:** Board Settings - Card Colors
- **Deskripsi:** Pewarnaan strip kiri kartu berdasarkan metode tertentu agar work item mudah dibedakan di board dan backlog company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > Card colors menyediakan "Colors based on": Issue Types, Priorities, Assignees, Queries, None.
  2. Untuk Issue Types, Priorities, dan Assignees, setiap nilai mendapat warna yang dapat diubah lewat color picker (palet + input hex).
  3. Mode Queries: board administrator menambah baris JQL + warna; work item mendapat warna dari query pertama yang cocok, dan baris dapat diurutkan ulang dengan drag.
  4. JQL card color divalidasi saat simpan dengan pesan galat parser bila tidak valid.
  5. Mode Assignees memberi warna otomatis per assignee yang dapat diubah.
  6. Warna tampil di kartu board dan baris backlog (Scrum dan Kanplan).
  7. Mode None menghilangkan strip warna.
  8. ⚠ Team-managed: tidak ada card colors berbasis aturan; pengguna dapat memberi warna per kartu lewat menu kartu > "Card color", disimpan sebagai atribut tampilan board.
- **Team-managed vs company-managed:** Company-managed: aturan warna per board; team-managed: warna manual per kartu.
- **Catatan & ketergantungan:** Model data: `board_card_colors (board_id, strategy, key, jql, color, position)`.

### BRD-039 Days in column indicator

- **Deskripsi:** Indikator visual di kartu yang menunjukkan berapa lama work item berada di kolom saat ini.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board administrator (mengaktifkan); anggota tim (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Titik-titik (dots) di bagian bawah kartu bertambah dan berubah warna seiring lamanya work item di kolom.
  2. Tooltip titik menampilkan "N hari di kolom ini".
  3. Indikator di-toggle lewat Card layout (company-managed dan team-managed).
  4. Hitungan direset ketika work item pindah ke kolom lain (pindah status dalam kolom yang sama tidak mereset).
  5. ⚠ Hitungan memakai hari kalender, bukan working days board; keputusan ini dicatat di Pertanyaan terbuka.
  6. ⚠ Ambang yang dipakai: 1 titik per 1 hari hingga 4 titik untuk ≥ 4 hari (abu-abu), lalu titik berubah merah untuk ≥ 8 hari; ambang disimpan sebagai konstanta konfigurasi.
  7. Waktu masuk kolom dihitung dari changelog transisi status (modul 08/28); work item tanpa transisi memakai waktu create.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butuh riwayat transisi status (changelog) dari [08 Aktivitas](08-komentar-lampiran-aktivitas.md) dan [28 Model Data](28-model-data-inti.md). Fase 3 (P2); checkbox "Days in column" di Card layout (BRD-037 butir 7) baru muncul setelah fitur ini dibangun.

### BRD-040 Working days

- **Alias inventaris:** Board Settings - Working Days
- **Deskripsi:** Konfigurasi time zone, hari kerja standar, dan hari libur board, dipakai untuk sisa hari sprint, burndown, dan Control chart.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > Working days menyediakan pemilih time zone board (default time zone sistem dari modul 22).
  2. Checkbox hari kerja standar Senin sampai Minggu (default Senin-Jumat); minimal satu hari harus dipilih.
  3. Non-working days: board administrator menambah dan menghapus tanggal libur spesifik (date picker, tanggal duplikat ditolak).
  4. Working days memengaruhi Burndown chart (area non-working day diarsir) dan sisa hari sprint di header board (BRD-014).
  5. Working days dipakai oleh Control chart untuk opsi compress/hide non-working days (modul 16).
  6. Hanya board administrator yang dapat mengubah; pengguna lain melihat read-only.
  7. ⚠ Team-managed: tidak tersedia halaman working days; perhitungan memakai Senin-Jumat dan time zone sistem.
- **Team-managed vs company-managed:** Hanya company-managed yang dapat dikonfigurasi.
- **Catatan & ketergantungan:** Laporan di [16 Laporan](16-laporan-insight.md). Model data: `board_working_days (board_id, timezone, weekdays_mask)` dan `board_non_working_dates`.

### BRD-041 Board Settings - Estimation & Tracking

- **Deskripsi:** Mengatur estimation statistic dan time tracking board Scrum yang dipakai di total backlog, laporan, dan insight.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (company-managed); Administrator project (team-managed lewat Features).
- **Aturan perilaku (acceptance criteria):**
  1. Estimation statistic dapat dipilih dari Story Points, Original Time Estimate, Issue Count, atau custom field bertipe number.
  2. Time tracking dapat dipilih None atau "Remaining Estimate and Time Spent".
  3. Statistik yang dipilih dipakai di total backlog/sprint, Burndown, Velocity, Insights, dan Sprint Workload by Assignee.
  4. Bila field estimasi tidak ada di screen atau field context work type yang ada di board, halaman menampilkan peringatan berisi daftar work type yang terdampak beserta tautan konfigurasi screen.
  5. Team-managed: Estimation diatur di Project settings > Features (Story points atau Time) dan memakai field "Story point estimate".
  6. Mengubah estimation statistic saat sprint aktif tidak mengubah data historis laporan; laporan menampilkan catatan "Statistik estimasi diubah pada <tanggal>".
- **Team-managed vs company-managed:** Company-managed: tab Estimation di Board settings; team-managed: toggle Estimation di Features.
- **Catatan & ketergantungan:** Estimasi dan time tracking di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md); laporan di [16 Laporan](16-laporan-insight.md).

### BRD-042 Board Settings - Timeline Toggle

- **Deskripsi:** Mengaktifkan atau menonaktifkan Timeline untuk board company-managed beserta pengaturan perencanaan child.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board administrator (company-managed); Administrator project (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Board settings menyediakan toggle Timeline (default aktif untuk board Scrum dan Kanban baru).
  2. ⚠ Pengaturan child issue planning (berbasis tanggal atau berbasis sprint) berada di Board settings > Timeline.
  3. Item Timeline muncul di sidebar board hanya bila toggle aktif.
  4. Team-managed: toggle Timeline berada di Project settings > Features (BRD-049).
- **Team-managed vs company-managed:** Company-managed per board; team-managed per project.
- **Catatan & ketergantungan:** Timeline dirinci di [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md).

## Alur Kanban: backlog, sub-filter & rilis

### BRD-043 Kanplan (Kanban backlog)

- **Deskripsi:** Mengaktifkan halaman Backlog terpisah untuk board Kanban sehingga perencanaan antrean terpisah dari board kerja.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator (mengaktifkan, company-managed), Administrator project (team-managed), anggota tim (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: board administrator mengaktifkan Kanplan dengan men-drag status ke kolom khusus "Kanban backlog" (paling kiri) di Board settings > Columns.
  2. ⚠ Tidak ada prasyarat toggle global oleh Admin Situs; Kanplan selalu dapat diaktifkan per board.
  3. Team-managed: Kanplan diaktifkan lewat toggle "Backlog" di Project settings > Features (BRD-049).
  4. Halaman Backlog memiliki dua bagian: "Board" (work item yang sudah berada di kolom board) dan "Backlog".
  5. Drag work item dari bagian Backlog ke bagian Board, atau aksi "Move to board", memindahkan status work item ke status pertama kolom pertama board (menjalankan transisi).
  6. Work item di bagian Backlog tidak tampil di board (BRD-013 butir 9).
  7. Ranking, inline create, quick filter, search, dan multi-select (Shift+klik untuk rentang, Ctrl/Cmd+klik untuk individual) tersedia di Kanplan.
  8. Header setiap bagian menampilkan jumlah work item.
  9. Klik work item membuka panel detail di kanan (BRD-018).
  10. Menonaktifkan Kanplan (mengosongkan kolom Kanban backlog atau mematikan toggle) mengembalikan status ke kolom board biasa (kolom pertama) dan halaman Backlog hilang dari sidebar.
  11. Kanplan memakai infinite scroll/paginasi server-side dan guardrail yang sama dengan board (BRD-051) serta lazy loading untuk daftar besar.
- **Team-managed vs company-managed:** Company-managed: aktivasi lewat kolom Kanban backlog; team-managed: lewat toggle Backlog.
- **Catatan & ketergantungan:** BRD-043 adalah heading utama Kanplan; SPR-024 (Kanplan (Kanban Backlog)) di [11 Backlog & Sprint](11-backlog-sprint.md) cukup menjadi rujukan ke fitur ini dengan Prioritas/Fase/Tier yang sama (P1 / Fase 2 / free-standard). Backlog Scrum di modul 11 memakai komponen daftar yang sama. Panel Epics di Kanplan dimiliki SPR-008 (BRD-044) dan panel Versions dimiliki REL-021 (BRD-045).

### BRD-044 Epics panel

- **Deskripsi:** Panel samping epic di Kanplan backlog. Fitur ini adalah rujukan; heading utama dan aturan lengkap ada di SPR-008 (Epics Panel in Backlog) di [11 Backlog & Sprint](11-backlog-sprint.md), modul ini hanya mencatat aturan khusus board Kanban. Lihat SPR-008.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Mengikuti SPR-008 (Edit Issues untuk menetapkan parent, Create Issues untuk membuat epic); board administrator mengaktifkan panel di board Kanban company-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: toggle "Epics panel" untuk board Kanban tersedia di Board settings > Columns dan hanya aktif setelah Kanplan diaktifkan (BRD-043). ⚠ Untuk team-managed berlaku SPR-008 butir 11 (epic sebagai filter toolbar, bukan panel).
  2. Daftar epic dengan nama, warna, progress, dan jumlah work item: lihat SPR-008 butir 2.
  3. Filter per epic dan opsi "Issues without epics": lihat SPR-008 butir 3.
  4. Drag work item ke epic untuk menetapkan parent: lihat SPR-008 butir 4.
  5. Tombol "Create epic" inline: lihat SPR-008 butir 5.
  6. Expand epic untuk detail tanggal, progress, dan jumlah per status category: lihat SPR-008 butir 6.
  7. Epic selesai disembunyikan dan menu "Mark as Done": lihat SPR-008 butir 7 dan 10.
  8. Khusus board: bila panel aktif, epic tidak tampil sebagai kartu di board Kanban maupun sebagai baris di Kanplan backlog.
- **Team-managed vs company-managed:** Lihat butir 1; perilaku panel company-managed mengikuti SPR-008.
- **Catatan & ketergantungan:** Pemilik: SPR-008 (P1 / Fase 2). Prioritas fitur ini disamakan dari P2 menjadi P1 agar panel di Kanplan dan backlog Scrum rilis bersamaan.

### BRD-045 Versions panel

- **Deskripsi:** Panel samping version di Kanplan backlog. Fitur ini adalah rujukan; heading utama dan aturan version ada di REL-021 (Versions panel (backlog)) di [15 Rilis & Komponen](15-rilis-versi-komponen.md), sedangkan sisi backlog dijelaskan di SPR-009 di [11 Backlog & Sprint](11-backlog-sprint.md). Lihat REL-021.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Mengikuti REL-021 (menetapkan Fix version mengikuti REL-009, membuat version mengikuti REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle panel Versions di View settings/toolbar backlog: lihat REL-021 butir 1.
  2. Daftar version Unreleased dengan progress dan tanggal rilis: lihat REL-021 butir 2.
  3. Filter per version dan opsi "Issues without versions": lihat REL-021 butir 3.
  4. Drag work item ke version menambah fixVersion tanpa menimpa nilai lain: lihat REL-021 butir 4.
  5. "Create version" inline hanya bagi pengguna berizin (tombol disembunyikan tanpa izin): lihat REL-021 butir 5.
  6. Tautan version ke halaman Releases/version detail: lihat REL-021 butir 2.
  7. Khusus board: pada board Kanban panel hanya tersedia bila Kanplan aktif (BRD-043) dan project memakai versions. ⚠ Untuk team-managed berlaku REL-021 butir 7 (tanpa panel khusus; filter Release di toolbar).
- **Team-managed vs company-managed:** Lihat butir 7.
- **Catatan & ketergantungan:** Pemilik: REL-021 (P1 / Fase 2). Prioritas fitur ini disamakan dari P2 menjadi P1; SPR-009 di modul 11 perlu disamakan menjadi P1 oleh pemilik modul 11.

### BRD-046 Kanban board sub-filter & hide completed issues

- **Deskripsi:** Filter tambahan khusus Kanban untuk membersihkan work item selesai dari board, baik berdasarkan rilis maupun umur penyelesaian.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board administrator.
- **Aturan perilaku (acceptance criteria):**
  1. Board settings > General menyediakan field "Kanban board sub-filter" (JQL) untuk board Kanban.
  2. Sub-filter default adalah `fixVersion in unreleasedVersions() OR fixVersion is EMPTY`, sehingga work item yang version-nya sudah dirilis hilang dari board.
  3. ⚠ Sub-filter hanya diterapkan pada work item di kolom paling kanan (Done); kolom lain selalu menampilkan semua work item hasil filter board.
  4. ⚠ Opsi "Hide completed issues older than" menawarkan 1 week, 2 weeks (default), 4 weeks, dan Show all.
  5. ⚠ "Completed" berarti work item berada di kolom paling kanan; umurnya dihitung dari waktu terakhir masuk kolom tersebut (bukan resolution date).
  6. Work item yang disembunyikan tetap muncul di pencarian dan tetap dihitung di CFD/Control chart.
  7. JQL sub-filter divalidasi saat simpan; JQL tidak valid ditolak dengan pesan galat parser.
  8. ⚠ Team-managed: tidak ada sub-filter yang dapat diedit; work item Done disembunyikan otomatis setelah 14 hari atau saat di-release (BRD-048 butir 10).
- **Team-managed vs company-managed:** Company-managed dapat dikonfigurasi; team-managed memakai aturan tetap.
- **Catatan & ketergantungan:** Fungsi JQL `unreleasedVersions()` di [12 JQL](12-jql-bahasa-kueri.md); laporan di [16 Laporan](16-laporan-insight.md).

### BRD-047 Release from board

- **Deskripsi:** Tombol Release di kolom Done board Kanban. Fitur ini adalah rujukan; heading utama dan aturan lengkap ada di REL-022 (Release from Kanban board) di [15 Rilis & Komponen](15-rilis-versi-komponen.md). Modul ini hanya mencatat perilaku khusus board. Lihat REL-022.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Mengikuti REL-022 (Administer Projects atau Manage Versions di semua project dalam filter board).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Release" di header kolom paling kanan, hanya untuk pengguna berizin: lihat REL-022 butir 1 dan 5.
  2. Dialog Release membuat version baru atau memilih version Unreleased: lihat REL-022 butir 2; khusus board, dialog juga menyediakan Start date opsional.
  3. Semua work item yang terlihat di kolom Done (setelah filter bar dan quick filter) diberi fixVersion lalu version ditandai Released (REL-022 butir 3); khusus board, dialog menampilkan jumlah work item yang akan dirilis sebelum konfirmasi.
  4. Work item yang dirilis hilang dari board karena sub-filter default (BRD-046; REL-022 butir 4).
  5. ⚠ Board multi-project: dialog meminta pengguna memilih satu project untuk version, dan hanya work item project tersebut yang diberi fixVersion (keputusan terbuka di REL-022).
  6. ⚠ Team-managed: berlaku REL-022 butir 7 (tanpa tombol Release di board; rilis dari halaman Releases).
  7. Release tidak dapat di-undo dari board; unrelease dilakukan di halaman Releases (modul 15).
  8. Work item yang sudah punya fixVersion lain mendapat tambahan version baru tanpa menghapus nilai lama.
  9. Release dijalankan sebagai satu operasi bulk; kegagalan edit sebagian work item dilaporkan dalam ringkasan dan version tetap dirilis untuk work item yang berhasil.
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** Pemilik: REL-022 (P1 / Fase 2). SPR-025 di modul 11 cukup menjadi rujukan ke REL-022 dan perlu disamakan menjadi P1 oleh pemilik modul 11. Bulk edit di [06 Work Item Inti](06-work-item-inti.md).

## Team-managed & board project bisnis

### BRD-048 Team-managed Kanban board behavior

- **Deskripsi:** Perilaku khusus board di project team-managed, di mana kolom identik dengan status dan konfigurasinya disederhanakan untuk tim kecil.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Administrator project team-managed (konfigurasi), Member project (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Satu board utama per project dibuat otomatis; board tambahan hanya lewat BRD-008 (Fase 3) dan board team-managed tidak dapat menarik work item dari project lain.
  2. Board filter tidak dapat diedit; board utama selalu memuat semua work item project.
  3. Kolom = status: menambah kolom membuat status, mengganti nama kolom berarti mengganti nama status (BRD-024 butir 6-7).
  4. Workflow per work type dapat diedit di workflow editor team-managed (modul 09) dan board mencerminkan status hasil editan secara langsung.
  5. Board settings berisi Columns and statuses serta Card layout; fitur (Backlog, Releases, Reports, Sprints, Estimation, Timeline) diatur di Project settings > Features (BRD-049).
  6. ⚠ Group by menawarkan None, Assignee, Epic/Parent, Subtask, dan Priority (BRD-032).
  7. ⚠ Filter bar menyediakan search, avatar assignee, Epic/Parent, Type, Label, dan Custom filters berbasis JQL (BRD-034).
  8. WIP limit hanya maksimum per kolom (tanpa minimum, BRD-029 butir 6).
  9. Tidak ada swimlane JQL, card colors berbasis aturan, working days, maupun Kanban sub-filter klasik.
  10. ⚠ Work item di kolom Done dibersihkan lewat Release (bila fitur Releases aktif) atau disembunyikan otomatis setelah 14 hari di kolom Done.
  11. ⚠ Reports team-managed Kanban menyediakan CFD; Control chart tidak tersedia (modul 16).
  12. Status baru yang dibuat dari workflow editor otomatis ditambahkan sebagai kolom baru di board utama (tidak ada unmapped statuses).
- **Team-managed vs company-managed:** Fitur ini mendefinisikan perbedaan itu sendiri; lihat BRD-023 untuk padanan company-managed.
- **Catatan & ketergantungan:** Atlassian terus menyamakan kedua tipe project, sehingga arsitektur sebaiknya memakai satu entitas board dengan flag `management = team|company` yang mengaktifkan/menonaktifkan kemampuan, bukan dua implementasi terpisah.

### BRD-049 Team-managed Board Features Toggles

- **Deskripsi:** Halaman Features di project team-managed untuk menyalakan dan mematikan fitur board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Administrator project team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle Backlog menyalakan/mematikan halaman Backlog (Kanplan untuk Kanban, BRD-043).
  2. Toggle Sprints menjadikan board berperilaku Scrum (board sprint aktif + backlog sprint) bila aktif dan Kanban bila mati.
  3. Toggle Estimation beserta pilihan tipe (Story points atau Time) mengatur estimasi project (BRD-041 butir 5).
  4. Toggle Timeline menampilkan/menyembunyikan Timeline di sidebar (BRD-042).
  5. Toggle Reports menampilkan/menyembunyikan halaman Reports.
  6. Toggle Releases menampilkan/menyembunyikan halaman Releases dan tombol Release di board.
  7. ⚠ Mematikan Sprints saat ada sprint aktif memerlukan konfirmasi; sprint aktif harus diselesaikan terlebih dahulu (dialog menautkan ke Complete sprint), sedangkan sprint future dihapus dan work item-nya kembali ke backlog setelah konfirmasi.
  8. ⚠ Ketika project memiliki banyak board (BRD-008), toggle Backlog dan Sprints berlaku per board dan toggle lainnya per project.
  9. Perubahan toggle berlaku seketika untuk semua anggota project dan tidak menghapus data (mis. mematikan Estimation tidak menghapus nilai Story point estimate).
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah tipe board dan Board settings.
- **Catatan & ketergantungan:** Halaman Project settings di [04 Proyek](04-proyek.md); sprint di [11 Backlog & Sprint](11-backlog-sprint.md).

### BRD-050 Business project board (Kanban-style)

- **Deskripsi:** Board di project bisnis (eks Jira Work Management) yang berperilaku mirip board Kanban team-managed dengan konfigurasi yang lebih sederhana.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Anggota project bisnis (memakai), Administrator project bisnis (mengatur kolom/status).
- **Aturan perilaku (acceptance criteria):**
  1. Board view tersedia di project bisnis dengan kolom = status workflow project.
  2. Board bisnis mendukung drag-drop transisi (BRD-016), filter bar (BRD-034), dan Group by (BRD-032).
  3. Board bisnis memakai card layout sederhana (toggle elemen kartu) dan mendukung inline create (BRD-019).
  4. ⚠ WIP limit, swimlane JQL, dan laporan Kanban (CFD/Control chart) tidak tersedia di board bisnis.
  5. ⚠ Collapse/expand kolom di board bisnis berperilaku sama dengan board software (BRD-013 butir 10-13).
  6. Board bisnis memakai komponen board yang sama dengan software sehingga perbaikan bug dan peningkatan berlaku untuk keduanya.
- **Team-managed vs company-managed:** Project bisnis mengikuti model team-managed atau company-managed sesuai pembuatannya (modul 04); perilaku board mengikuti BRD-048 untuk team-managed.
- **Catatan & ketergantungan:** Tampilan lain project bisnis (List, Calendar, Timeline, Form) di [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) dan [26 Kerja Bisnis](26-kerja-bisnis-formulir-persetujuan.md).

## Skala, API & event board

### BRD-051 Board issue limit handling

- **Alias inventaris:** Board Limits & Performance
- **Deskripsi:** Perilaku board berukuran besar: board memuat kartu bertahap (infinite scroll/paginasi) tanpa batas keras jumlah work item, dengan guardrail konfigurasi dan saran untuk mengecilkan cakupan board.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna (melihat peringatan), Board administrator (memperbaiki konfigurasi), Admin Situs (mengatur guardrail).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Board tidak memiliki batas keras jumlah work item yang ditampilkan; setiap kolom memuat kartu bertahap dengan infinite scroll sehingga seluruh hasil filter board dapat dijangkau dengan menggulir.
  2. ⚠ Admin Situs dapat menetapkan guardrail konfigurasi (ambang jumlah work item per board, default nonaktif); bila hasil filter melampaui ambang, board menampilkan banner peringatan kinerja beserta jumlah total hasil filter, tanpa menyembunyikan work item.
  3. Pesan memuat saran perbaikan: mempersempit board filter, memakai Kanban sub-filter/hide completed issues, atau mengaktifkan Kanban backlog; tautan ke Board settings ditampilkan bagi board administrator.
  4. ⚠ Infinite scroll dan guardrail yang sama berlaku untuk halaman Kanplan dan backlog Scrum.
  5. Laporan (CFD/Control chart) tetap menghitung data dari riwayat dan tidak terpengaruh pemuatan bertahap board.
  6. Backlog dan Kanplan memakai lazy loading/virtualized list sehingga daftar ribuan baris tetap responsif (render awal ≤ 2 detik untuk 1.000 baris).
  7. Endpoint kartu board memakai paginasi server-side per kolom berbasis cursor dan hanya mengembalikan field yang diperlukan kartu.
  8. Header kolom menampilkan jumlah total work item kolom dari server, bukan hanya jumlah kartu yang sudah termuat; drag-and-drop dan Rank tetap benar terhadap kartu yang belum termuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Target kinerja dan guardrail di [27 NFR](27-nfr-arsitektur-platform.md). Pada Oktober 2026 board company-managed di Jira Cloud dipindah ke board modern tanpa batas tampilan 5.000 work item dan memakai infinite scroll; rilis Juli 2026 menambah swimlane (BRD-031), subtask di kartu (BRD-022), inline edit di kartu (BRD-015 butir 10), dan saved views (BRD-055). Clone mengikuti perilaku board modern sejak awal; nilai ambang guardrail adalah konstanta konfigurasi.

### BRD-052 Jira Software (Agile) REST API for Boards & Sprints

- **Deskripsi:** API untuk board, backlog, sprint, dan ranking. Endpoint internal dipakai SPA sejak Fase 1; di atasnya disediakan lapisan kompatibel publik dengan bentuk endpoint Agile REST API Jira agar alat pihak ketiga mudah dimigrasikan. Kontrak endpoint rinci dimiliki API-032 (Agile Boards API & Board Configuration), API-033 (Agile Backlog & Rank API), dan API-034 (Agile Sprints API) di [23 API Publik & Webhook](23-api-publik-webhook.md).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Integrator/Developer API, Admin Situs, SPA produk ini; setiap endpoint menegakkan permission yang sama dengan UI.
- **Aturan perilaku (acceptance criteria):**
  1. Board: list (dengan filter nama, tipe, projectKeyOrId, paginasi `startAt`/`maxResults`), get, create (nama, tipe, filterId, location), delete board, get configuration (kolom, status mapping, estimation, ranking), serta get issues/backlog/epics/versions/projects per board (kontrak rinci API-032).
  2. Sprint: create, get, update parsial (POST) maupun penuh (PUT), delete, start/close lewat perubahan `state`, get issues in sprint, move issues to sprint, dan swap sprint (kontrak rinci API-034).
  3. Backlog: endpoint move issues to backlog (maksimal 50 work item per request; kontrak rinci API-033).
  4. Rank: endpoint rank issues dengan `rankBeforeIssue` atau `rankAfterIssue` (maksimal 50 work item per request), dengan respons parsial (207) bila sebagian gagal (API-033).
  5. Board properties dan sprint properties (key-value JSON) dapat di-set, di-get, dan dihapus untuk menyimpan data integrasi.
  6. Respons sprint memuat `id`, `state`, `name`, `startDate`, `endDate`, `completeDate`, `originBoardId`, dan `goal`.
  7. Endpoint internal berada di bawah prefix versi produk (`/api/v1/agile/...`), dipakai SPA sejak Fase 1, dan didokumentasikan lewat anotasi Swagger (`make swag`) dengan `@Security BearerAuth`; lapisan kompatibel publik `/rest/agile/1.0/...` (API-032 s.d. API-034) memetakan ke handler dan use case yang sama.
  8. Galat mengikuti format respons terpadu repo (`internal/shared/delivery/http/response`) dengan pesan multi-bahasa (EN/ID).
  9. Pembagian fase ditulis eksplisit: Fase 1 (P0) = endpoint internal `/api/v1/agile/...` untuk board, configuration, issues per kolom, backlog, rank, dan sprint yang dibutuhkan SPA (butir 1-4, 6-8); Fase 2 = lapisan kompatibel publik read-only (`GET` board, configuration, issues, backlog, sprint) dengan autentikasi token API; Fase 3 = lapisan kompatibel publik untuk operasi tulis (create/delete board, sprint, move, rank), properties (butir 5), rate limit, dan event webhook (BRD-053).
- **Team-managed vs company-managed:** Board team-managed dapat dibaca lewat API yang sama; create/delete board ditolak untuk team-managed dengan HTTP 400 (API-032 butir 12).
- **Catatan & ketergantungan:** Konvensi API publik, autentikasi token, dan rate limit di [23 API Publik & Webhook](23-api-publik-webhook.md); sprint di [11 Backlog & Sprint](11-backlog-sprint.md). Prioritas disamakan menjadi P0 / Fase 1 mengikuti API-032 s.d. API-034; status Fase 1 di API-032 s.d. API-034 dibaca sebagai bentuk kontrak yang sudah dipakai endpoint internal, sedangkan publikasi sebagai kontrak publik bertahap Fase 2-3 sesuai butir 9 (perlu diselaraskan di modul 23).

### BRD-053 Board & Sprint Events (Automation & Webhooks)

- **Deskripsi:** Event board dan sprint yang dapat memicu aturan otomasi dan webhook.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs dan Admin Proyek (membuat aturan/webhook), Integrator/Developer API (menerima webhook).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Trigger otomasi tersedia untuk Sprint created, Sprint started, Sprint completed, dan Sprint mendekati end date (N hari sebelum, dapat dikonfigurasi).
  2. Webhook event sprint: `sprint_created`, `sprint_updated`, `sprint_started`, `sprint_closed`, `sprint_deleted`.
  3. Webhook event board: `board_created`, `board_updated`, `board_deleted`, `board_configuration_changed`.
  4. Payload memuat data sprint/board terkait (bentuk sama dengan respons BRD-052), `timestamp`, dan `webhookEvent`.
  5. Event diterbitkan setelah transaksi berhasil di-commit (outbox) sehingga tidak ada event untuk perubahan yang di-rollback.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mesin otomasi di [20 Otomasi](20-otomasi.md); pengiriman webhook dan retry di [23 API Publik & Webhook](23-api-publik-webhook.md); katalog event domain di DM-014 (Issue Event Catalog & Domain Events) di [28 Model Data](28-model-data-inti.md) dan katalog event webhook di API-065 (Webhook Events Catalogue). Event board/sprint di butir 2-3 wajib terdaftar di kedua katalog tersebut.

## Rujukan fitur lama (Data Center/deprecated)

### BRD-054 Board Settings - Issue Detail View

- **Deskripsi:** Konfigurasi field pada panel detail work item board company-managed versi lama. Tidak dibangun karena issue view baru memakai layout work type; dicantumkan sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Board administrator (pada produk lama).
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: tab dengan section General fields, Date fields, People, dan Links pada Board settings tidak tersedia di produk ini.
  2. Tidak dibangun: penambahan, penghapusan, dan pengurutan field per section; panel detail board (BRD-018) selalu memakai layout issue view dari modul 05/06.
  3. ⚠ Konfigurasi Issue Detail View lama yang diimpor dari Jira diabaikan dan importer mencatat peringatan; tab ini tidak memengaruhi issue view baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan; impor di [25 Integrasi & Impor](25-integrasi-impor-ekspor-ekstensi.md).

## Tambahan di luar inventaris

### BRD-055 Board saved views

- **Alias inventaris:** Board saved views (catatan: fitur board modern Jira Cloud rilis Juli 2026, belum tercatat di inventaris)
- **Deskripsi:** Board administrator menyimpan konfigurasi tampilan board (filter bar, quick filter aktif, Group by, toggle tampilan) sebagai view bernama yang dibagikan ke semua pengguna board, mirip saved view List view (PLAN-016/PLAN-017).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board administrator (company-managed) atau Admin Proyek (team-managed) membuat, mengganti nama, dan menghapus view bersama; semua pengguna board memilih view dan menyesuaikannya secara personal.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Toolbar board menyediakan switcher view; view default board selalu ada dan tidak dapat dihapus.
  2. ⚠ "Save view" menyimpan kombinasi filter bar (BRD-034), quick filter aktif (BRD-033), Group by (BRD-032), dan toggle View settings (BRD-035) sebagai view bernama bersama; hanya board administrator/Admin Proyek yang dapat menyimpan atau menimpa view bersama.
  3. ⚠ Pengguna lain dapat mengubah tampilan secara personal tanpa menimpa view bersama; tombol "Reset" mengembalikan tampilan ke konfigurasi view tersimpan.
  4. ⚠ Satu board dapat memiliki maksimal 10 saved view (selaras PLAN-017 butir 8); menyimpan view ke-11 ditolak dengan pesan "Batas jumlah view tercapai".
  5. ⚠ URL board memuat parameter `view=<id>` sehingga view dapat dibagikan sebagai deep link; view yang dihapus jatuh ke view default dengan pemberitahuan.
  6. ⚠ Filter atau field yang direferensikan view lalu dihapus diabaikan saat view dibuka dengan pemberitahuan "Beberapa filter tidak lagi tersedia dan dilewati".
  7. Saved view tidak mengubah board filter, pemetaan kolom, atau konfigurasi Board settings; view hanya mengubah penyajian.
- **Team-managed vs company-managed:** ⚠ Company-managed: dikelola board administrator; team-managed: dikelola Admin Proyek.
- **Catatan & ketergantungan:** Ditambahkan di luar inventaris berdasarkan rilis board modern Jira Cloud (Juli 2026); butir rinci belum terverifikasi. Preferensi personal di BRD-035 (`board_user_prefs`); view bersama disimpan di tabel terpisah `board_saved_views`. Berbeda dari saved filter JQL (modul 13).

## Pertanyaan terbuka

1. Apakah himpunan status default template Kanban company-managed memakai To Do / In Progress / Done (selaras team-managed) atau set lama Backlog / Selected for Development / In Progress / Done (BRD-001)?
2. Saat project yang menjadi location board dihapus atau diarsipkan, apakah board dipindah ke location pengguna pembuatnya, ikut diarsipkan, atau dihapus (BRD-004)?
3. Pada multiple board team-managed, apakah sprint dimiliki per board atau dibagi di seluruh project, dan apakah toggle Backlog/Sprints berlaku per board (BRD-008, BRD-049)?
4. Apakah board team-managed perlu mendukung work item dari project lain (board lintas project team-managed) di masa depan (BRD-003, BRD-048)?
5. Apakah board yang dihapus perlu masuk trash dengan masa pemulihan (mis. 30 hari) sebagai pembeda dari Jira Cloud (BRD-005)?
6. Apakah Days in column dihitung dengan hari kalender atau working days board, dan berapa ambang warna yang dipakai (BRD-039)?
7. Apakah jumlah di header kolom dan indikator WIP harus mengikuti filter bar yang aktif atau selalu menghitung seluruh isi kolom (BRD-034, BRD-028)?
8. Berapa ambang default guardrail jumlah work item per board (atau per kolom) yang memicu banner peringatan kinerja, dan apakah guardrail dapat diatur per board selain global (BRD-051)?
9. Apakah fitur stand-up memerlukan timer per orang dan pengaturan urutan giliran yang disimpan per board (BRD-036)?
10. Apakah inline create di board sprint aktif tetap memasukkan work item ke sprint bila pengguna tidak memiliki Manage Sprints (BRD-019)?
11. Apakah publikasi lapisan kompatibel publik REST API Agile mengikuti pembagian Fase 2 (read-only) dan Fase 3 (tulis, properties, webhook) pada BRD-052 butir 9, atau seluruhnya dipublikasikan sekaligus bersama API-032 s.d. API-034?
12. Apakah saved view board (BRD-055) juga tersedia di backlog/Kanplan, dan apakah batas 10 view berlaku per board atau per project?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk board: app stand-up/retrospective, swimlane dan card color lanjutan, board lintas instance, WIP analytics, dan Kanban metrics tambahan tidak dibangun sebagai fitur native; integrasi dimungkinkan lewat REST API (BRD-052) dan webhook (BRD-053), sedangkan platform ekstensi dibahas di modul 25.
- **Fitur Data Center/deprecated:** tab Board settings "Issue Detail View" (BRD-054), opsi wizard lama "Board with sample data"/"Board with new project", prasyarat toggle global Kanplan, dan pola key work item kustom tidak dibangun; URL legacy `RapidBoard.jspa` hanya didukung sebagai redirect (BRD-007).
- **Aplikasi mobile native** (iOS/Android): produk ini hanya menyediakan tampilan web responsif (BRD-013 butir 15); aplikasi native di luar cakupan modul ini.
- **Layanan Atlassian lintas produk**: Atlassian Intelligence/Rovo pada board, Atlassian Analytics, Compass, dan sinkronisasi board dengan Confluence/Trello tidak dibangun; AI opsional produk ini dibahas di modul 06 dan 25.
- **Backlog Scrum dan siklus hidup sprint** (Create/Start/Complete sprint, parallel sprints, sprint permission) dibahas di modul 11, bukan di modul ini.
- **Laporan agile** (Burndown, Velocity, CFD, Control chart, Insights) dibahas di modul 16; modul ini hanya menyediakan data konfigurasi board yang dipakai laporan tersebut.
