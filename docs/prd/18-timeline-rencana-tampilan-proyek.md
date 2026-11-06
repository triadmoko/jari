# 18 — Timeline, Rencana Lintas Proyek & Tampilan Proyek

Modul ini mendefinisikan semua tampilan perencanaan dan tampilan alternatif work item di dalam proyek pada produk ini: navigasi tab view proyek, **Timeline** proyek (Gantt per proyek untuk software dan business project, berbasis Start date/Due date, Parent, Rank, Issue color, dan issue link `blocks`), **List** (tabel ala spreadsheet dengan inline edit, kolom, grouping, filter, ekspor, conditional formatting, dan view bersama), **Calendar** (penjadwalan bulanan dengan panel unscheduled), **Summary** proyek, dan **Board view** business project. Modul ini juga mencakup **Plans** tier premium (eks Advanced Roadmaps): perencanaan lintas proyek dengan issue sources, exclusion rules, view Timeline/List/Calendar/Summary/Dependencies/Program board, mekanisme draft dan Review changes, scenarios, rollups, capacity per team, cross-project releases, auto-scheduler, warnings, saved views, serta REST API; entitas **Teams** tingkat platform beserta Team field dan profil team; serta **Goals** yang dapat ditautkan ke work item. Integrasi dengan produk portofolio terpisah (Jira Align, Focus, Talent) dan Overviews legacy hanya dicatat sebagai rujukan. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, klaim JWT di context `sess`), sehingga seluruh tabel (mis. `project_views`, `timeline_view_configs`, `list_view_configs`, `list_format_rules`, `plans`, `plan_sources`, `plan_exclusion_rules`, `plan_scenarios`, `plan_changes`, `plan_views`, `plan_teams`, `teams`, `team_members`, `goals`, `goal_links`) dan endpoint di bawah `/api/v1` adalah pekerjaan baru yang dikemas sebagai feature module baru mengikuti pola Clean Architecture repo (mis. `internal/features/projectview`, `internal/features/timeline`, `internal/features/plan`, `internal/features/team`, `internal/features/goal`), dengan route SPA baru di `frontend/` mengikuti pola `/projects/{key}/{view}` (PRJ-027). Tampilan proyek (Timeline, List, Calendar) tidak menyimpan data work item sendiri; semuanya membaca dan menulis field work item yang didefinisikan di modul 05 dan 06, sedangkan Plans menambahkan lapisan draft per plan per scenario sebelum perubahan diterapkan ke work item.

Modul terkait:

- [00 Visi & Baseline](00-visi-ruang-lingkup-glosarium-baseline.md) — Plan Tiers & Feature Flags (VISI-003) untuk gating fitur premium/platform, terminologi work item/work type (VISI-002), penamaan produk (VISI-001).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects (PERM-020), Create/Edit Issues (PERM-021), Link & Schedule Issues (PERM-024), Manage Sprints (PERM-025), Manage Versions (PERM-026), Plans Permissions (PERM-052), Plan-tier Permission Limitations (PERM-005), penegakan issue security (PERM-044).
- [04 Proyek](04-proyek.md) — Project Navigation & Views (PRJ-027), Project Summary (PRJ-028), Features toggle (PRJ-049, PRJ-050), model team-managed vs company-managed (PRJ-003).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — Work type hierarchy (CFG-006), Custom hierarchy levels above Epic (CFG-007), Parent field (CFG-013), Team field (CFG-019), Issue color field (CFG-020), Target start/Target end (CFG-022), link types (CFG-056).
- [06 Work Item Inti](06-work-item-inti.md) — Issue View Modes (ISS-008), Due Date & Start Date (ISS-021), Rank/LexoRank (ISS-022), Parent Field & Epic Children (ISS-025), Issue Links (ISS-029), Bulk Change (ISS-035).
- [09 Workflow](09-workflow-status-transisi.md) — transisi saat kartu di-drag di Board view dan status category untuk progres.
- [10 Board](10-board-kolom-kartu.md) — board filter dan cross-project board (BRD-003), Board Settings - Timeline Toggle (BRD-042), Card colors (BRD-038), Board view Kanban (BRD-013).
- [11 Backlog & Sprint](11-backlog-sprint.md) — Sprint field dan siklus sprint yang dirender di Timeline dan Plans.
- [12 JQL](12-jql-bahasa-kueri.md) — filter lanjutan List, sumber Filter di Plans, field `team`, `parent`, fungsi hierarki.
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — List view hasil pencarian (SRC-023), Project Work items page (SRC-032), Share search (SRC-048), saved filter.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — estimasi waktu/story points untuk capacity dan rollup; User Workload Report (TIME-032) sebagai pembanding PLAN-067.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Releases in Timeline and board filters (REL-023), Cross-project releases (REL-024), halaman Releases (REL-012).
- [16 Laporan & Insight](16-laporan-insight.md) — metrik & widget Summary (RPT-034, pemilik), Space Insights (RPT-043), Epic report.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget kalender dan alternatif Overviews lintas proyek.
- [19 Notifikasi](19-notifikasi-email.md) — notifikasi share Timeline/plan dan update goal.
- [20 Otomasi](20-otomasi.md) — trigger perubahan tanggal, rank, dan Team field.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — sidebar (menu Plans, Teams, Goals), Starred & Recent items, preferensi per pengguna.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — Plans settings, Teams administration, Goals administration.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — konvensi REST, Plans API, Teams API.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — embed/Smart Link ke wiki eksternal, ekspor Plans/Timeline/List, konektor portofolio.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — business project type, navigasi view business, Category field, approvals.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — anggaran performa render, batas skala, realtime.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Plan, Team, Goal, view config, dan event domain.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) — Project summary page service space (JSM-PRJ-014) sebagai turunan Summary.

## Persona terkait

- **Admin Situs** — mengaktifkan tier premium/Plans, mengonfigurasi hierarki di atas Epic, mengelola akses semua plan, mengelola team terkelola (managed) dan Goals.
- **Admin Proyek/Project Lead** — mengaktifkan Timeline/view proyek, menyimpan view bersama Timeline/List/Board, mengatur tab, mengelola conditional formatting.
- **Scrum Master/Team Lead** — menjadwalkan Epic di Timeline, mengatur sprint dan capacity team di Plans, memantau warning jadwal.
- **Developer/Anggota Tim** — membuat dan menjadwalkan work item dari Timeline/List/Calendar, mengedit inline, mengisi Team field.
- **Product Owner/PM** — menyusun roadmap lintas proyek di Plans, membandingkan scenario, menyelaraskan release dan Goals.
- **Stakeholder/Viewer** — membaca Timeline, Summary, Plan Summary, dan tautan share secara read-only.
- **Pengguna Tim Bisnis** — memakai List, Board, Calendar, Timeline, dan Summary di business project.
- **Integrator/Developer API** — mengelola plan dan team lewat REST API.

## User story modul

- Sebagai Scrum Master/Team Lead, saya ingin men-drag bar Epic di Timeline untuk mengisi Start date dan Due date agar jadwal tim terlihat dalam satu layar tanpa membuka setiap work item (PLAN-002, PLAN-004).
- Sebagai Developer/Anggota Tim, saya ingin membuat Epic dan child langsung dari Timeline serta memindahkan child antar Epic agar penyusunan roadmap cepat (PLAN-005, PLAN-006).
- Sebagai Scrum Master/Team Lead, saya ingin menghubungkan dua bar dengan dependency dan melihat warning konflik agar urutan pekerjaan yang salah segera terdeteksi (PLAN-008, PLAN-009).
- Sebagai Admin Proyek/Project Lead, saya ingin menyimpan filter dan view settings Timeline atau List sebagai view bersama agar seluruh anggota proyek melihat konfigurasi yang sama (PLAN-010, PLAN-017).
- Sebagai Pengguna Tim Bisnis, saya ingin mengedit work item seperti spreadsheet di List dan mewarnai baris berisiko dengan conditional formatting agar pekerjaan penting langsung terlihat (PLAN-016, PLAN-021).
- Sebagai Pengguna Tim Bisnis, saya ingin menjadwalkan item tanpa tanggal dengan men-drag dari panel Unscheduled ke Calendar agar perencanaan kampanye mudah (PLAN-022, PLAN-023).
- Sebagai Stakeholder/Viewer, saya ingin membuka Summary proyek dan mengklik grafik untuk melihat daftar item terkait agar status proyek mudah dipahami (PLAN-025, rujukan ke PRJ-028 dan RPT-034).
- Sebagai Product Owner/PM, saya ingin membuat plan dari beberapa board, proyek, dan filter agar roadmap lintas tim terlihat dalam satu hierarki (PLAN-027, PLAN-028, PLAN-034).
- Sebagai Product Owner/PM, saya ingin mengubah jadwal di plan sebagai draft dan meninjaunya sebelum disimpan agar eksperimen perencanaan tidak langsung mengganggu tim (PLAN-045, PLAN-046).
- Sebagai Scrum Master/Team Lead, saya ingin mengatur capacity per sprint untuk setiap team dan melihat indikator over capacity agar komitmen realistis (PLAN-049, PLAN-050, PLAN-053).
- Sebagai Product Owner/PM, saya ingin membuat cross-project release dan melihat item yang dijadwalkan setelah tanggal rilis agar rilis program terkoordinasi (PLAN-051, PLAN-053).
- Sebagai Admin Situs, saya ingin mendefinisikan level Initiative di atas Epic agar portofolio dapat direncanakan di Plans dan Timeline (PLAN-032).
- Sebagai Developer/Anggota Tim, saya ingin memilih Team pada work item agar pekerjaan tercatat untuk team yang benar dan terhitung di capacity plan (PLAN-057, PLAN-058).
- Sebagai Product Owner/PM, saya ingin menautkan work item ke Goal agar kontribusi pekerjaan terhadap tujuan organisasi terlihat (PLAN-061, PLAN-062).
- Sebagai Integrator/Developer API, saya ingin membuat dan memperbarui plan serta team lewat REST API agar konfigurasi perencanaan dapat diotomasi (PLAN-056, PLAN-060).
- Sebagai Scrum Master/Team Lead, saya ingin mengalokasikan waktu setiap orang per minggu ke work item dan pekerjaan non-proyek di tab Capacity agar beban berlebih terlihat sebelum terjadi (PLAN-067).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 6 |
| P1 | 19 |
| P2 | 36 |
| P3 | 6 |
| **Total** | **67** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 16 |
| Fase 3 | 10 |
| Fase 4 | 37 |
| Fase — (dc-only/deprecated atau produk terpisah, rujukan) | 4 |

Catatan: inventaris berisi 71 fitur; 5 pasang yang tumpang tindih antara domain perencanaan dan domain business work digabung menjadi satu heading dengan baris **Alias inventaris** (Business project Timeline, Project List view, Project Calendar view, Project Summary page, Overviews), sehingga file ini memuat 66 heading fitur inventaris, ditambah 1 fitur di luar inventaris bertanda ⚠ (PLAN-067 Individual capacity planning), total 67 heading.

Strategi urutan implementasi: Fase 1 tidak membangun fitur modul ini, tetapi harus sudah menyiapkan field yang dipakai seluruh tampilan (Start date, Due date, Parent, Rank berbasis LexoRank, Issue color, issue link `blocks`) di modul 05/06 dan kerangka tab proyek (PRJ-027). Fase 2 membangun Timeline proyek software secara utuh (render, scheduling, inline create, ranking, dependency, progres, warning, view settings, marker sprint/release, batas performa), List view inti dengan inline edit dan inline create, serta Summary proyek, karena semuanya hanya membaca/menulis field yang sudah ada. Fase 3 menambah tampilan untuk business project dan konfigurabilitas tampilan (Calendar, Board view business, Timeline business, kolom/grouping/filter/ekspor List, conditional formatting, ekspor Timeline). Fase 4 membangun Plans sebagai feature module tersendiri dengan urutan: entitas plan + issue sources + permission → mesin draft (change set per scenario) dan Review changes → view Timeline/List → teams/capacity, releases, rollups, warnings → scenarios, Program board, auto-scheduler, API; Teams dan Goals dibangun lebih dulu dalam Fase 4 karena dipakai capacity dan Team field; tab Capacity individual (PLAN-067) dibangun setelah Teams karena memakai data anggota dan jam kerja yang sama. Integrasi produk portofolio terpisah dan Overviews legacy tidak dibangun.

## Fitur

## Navigasi tampilan proyek

### PLAN-001 Project views navigation (tabs)

- **Deskripsi:** Navigasi tab view di header proyek (Summary, Timeline, Backlog, Board, Calendar, List, dll.) beserta pengaturan tab mana yang tampil, urutannya, dan view default saat proyek dibuka. Modul ini merinci perilaku tab untuk view yang didefinisikan di sini; kerangka navigasi proyek ada di PRJ-027.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (memakai, perlu Browse Projects); Admin Proyek (mengaktifkan/menonaktifkan view lewat Features, mengatur default proyek).
- **Aturan perilaku (acceptance criteria):**
  1. Header proyek menampilkan tab horizontal sesuai tipe proyek: software (Summary, Timeline, Backlog, Board, Calendar, List, Forms, Reports, dll.) dan business (Summary, List, Board, Calendar, Timeline, Forms, Approvals, dll.); tab yang tidak muat dipindah ke menu "More".
  2. Tombol "+" di ujung tab membuka daftar view yang belum aktif; memilih view menambahkannya ke tab (bagi Admin Proyek: mengaktifkan feature proyek; bagi member: hanya menampilkan kembali view yang disembunyikan secara personal).
  3. ⚠ Menu "..." pada tab menyediakan "Hide from navigation"; menyembunyikan berlaku per pengguna, sedangkan menghapus view untuk semua anggota hanya bisa dilakukan Admin Proyek melalui Features (PRJ-049/PRJ-050).
  4. ⚠ Tab dapat diurutkan ulang dengan drag-and-drop; urutan personal disimpan per pengguna dan Admin Proyek dapat menetapkan urutan default proyek.
  5. ⚠ Saat proyek dibuka tanpa path view, sistem membuka view terakhir yang dikunjungi pengguna di proyek tersebut; bila belum ada, view default proyek dari template (mis. Board untuk Scrum/Kanban, List untuk business).
  6. Ketersediaan view bergantung pada tipe proyek dan feature yang diaktifkan; route view yang tidak tersedia (mis. `/projects/{key}/timeline` saat Timeline dinonaktifkan) menampilkan halaman "View ini tidak tersedia di proyek ini" beserta tautan ke Project settings bagi admin, bukan galat server.
- **Team-managed vs company-managed:** Team-managed: view dikontrol oleh Project settings > Features. Company-managed: Timeline, Backlog, dan Board melekat pada board terpilih (board switcher), view lain melekat pada proyek.
- **Catatan & ketergantungan:** Kerangka tab dan sidebar di PRJ-027 dan [21 Navigasi](21-navigasi-ux-preferensi.md); tab business project dirinci juga di [26](26-kerja-bisnis-formulir-persetujuan.md). Preferensi tab personal disimpan di tabel preferensi pengguna (modul 21), bukan `localStorage`, agar konsisten antar perangkat.

## Timeline proyek (software project)

### PLAN-002 Project Timeline (core view)

- **Deskripsi:** Tampilan Gantt per proyek yang menampilkan work item level Epic, child-nya, dan level di atas Epic (tier premium) di sepanjang sumbu waktu. Tersedia di semua tier untuk team-managed dan company-managed software project; dulu dikenal sebagai "Roadmap".
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dan Admin Proyek (mengedit sesuai PERM-021/PERM-024); Stakeholder/Viewer dengan Browse Projects (read-only).
- **Aturan perilaku (acceptance criteria):**
  1. Layout dua panel: panel kiri berisi daftar hierarki (ikon work type, key, summary) dan panel kanan berisi grid waktu dengan bar; baris kiri dan kanan selalu sejajar saat di-scroll vertikal.
  2. Panel kiri dapat di-resize dengan drag pembatas (lebar minimum 200 px, lebar disimpan per pengguna) atau dilipat penuh agar grid lebih lebar.
  3. Baris Epic dapat di-expand/collapse untuk menampilkan child work item (Story, Task, Bug, dan custom work type level standar); status expand disimpan per pengguna per proyek.
  4. Klik baris atau bar membuka issue detail panel di sisi kanan tanpa meninggalkan Timeline; panel menyediakan opsi "Open in full page" (ISS-008), dan perubahan di panel langsung tercermin pada bar.
  5. Header waktu bersifat sticky saat scroll vertikal dan menampilkan garis vertikal penanda "Today" berdasarkan zona waktu pengguna.
  6. Baris tanpa tanggal tampil sebagai baris kosong di grid; hover pada grid baris tersebut memunculkan placeholder bar untuk menjadwalkan (PLAN-004).
  7. Lozenge status dan avatar assignee tampil di baris/bar sesuai view settings (PLAN-010).
  8. Search box memfilter baris berdasarkan summary atau key (case-insensitive, minimal 1 karakter); ancestor dari item yang cocok tetap ditampilkan agar konteks hierarki tidak hilang.
  9. Deep link: URL Timeline dapat memuat work item terpilih (query `selectedIssue=<KEY>`) dan filter aktif; membuka URL tersebut men-scroll ke baris itu dan membuka detail panel.
  10. Team-managed: Timeline adalah tab bawaan proyek dengan sumber data seluruh work item proyek.
  11. ⚠ Company-managed: Timeline melekat pada board dan berbasis filter board (BRD-003); hanya work item yang lolos filter board yang tampil, dan board lintas proyek dapat menampilkan Epic dari beberapa proyek.
  12. ⚠ Tier premium: level hierarki di atas Epic (mis. Initiative) tampil sebagai baris induk di atas Epic bila hierarchy level dikonfigurasi (PLAN-032); di team-managed level atas hanya tampil bila Epic proyek memiliki parent level atas.
  13. ⚠ Subtask ditampilkan di Timeline proyek sebagai baris di bawah parent-nya (level di bawah child Epic) dan dapat di-expand/collapse; subtask dapat dijadwalkan dan diedit langsung di view, ikut tersaring oleh filter Timeline (PLAN-010), dan dapat dipilih untuk bulk update (PLAN-005 butir 10). Perilaku ini menggantikan perilaku lama yang menyembunyikan subtask.
  14. ⚠ Work item level standar tanpa parent Epic ditampilkan sebagai baris level atas di bagian bawah daftar (setelah Epic) agar tetap dapat dijadwalkan; perilaku ini dapat dimatikan lewat view settings.
  15. Pengguna yang hanya memiliki Browse Projects melihat Timeline read-only: bar tidak dapat di-drag, tombol create dan menu aksi edit disembunyikan.
  16. Empty state saat proyek belum punya Epic menampilkan ilustrasi, penjelasan singkat, dan ajakan "Create Epic" (disembunyikan bila pengguna tidak punya Create Issues).
- **Team-managed vs company-managed:** Team-managed: satu Timeline per proyek dengan sumber seluruh proyek. Company-managed: satu Timeline per board dengan sumber filter board, dapat lintas proyek.
- **Catatan & ketergantungan:** Data model memakai field Start date, Due date (ISS-021), Parent (ISS-025, CFG-013), Rank (ISS-022), Issue color (CFG-020). Tidak ada tabel jadwal terpisah; Timeline hanya menyimpan konfigurasi tampilan. Anggaran performa render di modul 27.

### PLAN-003 Timeline enablement & configuration

- **Deskripsi:** Pengaturan untuk mengaktifkan/menonaktifkan Timeline dan menentukan cara child work item dijadwalkan (berdasarkan tanggal atau sprint). Lokasi setting berbeda antara team-managed (Project settings > Features) dan company-managed (Board settings > Timeline).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed), Board administrator (company-managed), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Team-managed: toggle "Timeline" di Project settings > Features (PRJ-049) mengaktifkan atau menyembunyikan tab Timeline untuk seluruh proyek.
  2. Company-managed: Board settings > Timeline menyediakan toggle aktifkan Timeline per board (BRD-042); default aktif untuk board baru.
  3. ⚠ Company-managed: bila filter board tidak menyertakan work type Epic, Timeline menampilkan banner peringatan "Filter board tidak menyertakan Epic, sehingga Epic tidak tampil di Timeline" beserta tautan ke pengaturan filter bagi board admin.
  4. ⚠ Opsi "Child issue planning" menentukan penjadwalan child: berdasarkan Dates (Start date/Due date) atau berdasarkan Sprints; opsi Sprints hanya dapat dipilih pada board Scrum atau proyek dengan fitur Sprints aktif, dan tidak tampil di board Kanban.
  5. Mode Sprints: bar child mengikuti rentang tanggal sprint yang di-assign dan tidak dapat di-drag bebas; child tanpa sprint tampil tanpa bar.
  6. Field tanggal yang dipakai adalah Start date dan Due date; bila field tidak ada di screen edit work type terkait, bar item tersebut read-only dengan tooltip "Field Start date/Due date tidak tersedia di layar work type ini".
  7. Hanya Admin Proyek (team-managed) atau Board administrator (company-managed) yang dapat mengubah setting; pengguna lain melihat setting read-only atau tidak melihatnya sama sekali.
  8. Menonaktifkan Timeline hanya menyembunyikan tab; nilai Start date, Due date, dan issue link dependency tetap tersimpan dan muncul kembali saat Timeline diaktifkan ulang.
  9. Bila Timeline tidak tersedia (dinonaktifkan untuk board atau proyek tidak memenuhi syarat), membuka route Timeline menampilkan pesan penjelasan alasan dan, bagi admin, tautan untuk mengaktifkannya.
- **Team-managed vs company-managed:** Team-managed per proyek lewat Features; company-managed per board lewat Board settings, termasuk peringatan filter Epic.
- **Catatan & ketergantungan:** Toggle board didefinisikan juga di BRD-042 (modul 10) sebagai titik masuk UI; aturan perilaku detail menjadi milik fitur ini. Label setting persis mengikuti glosarium modul 00.

### PLAN-004 Timeline bar scheduling (drag dates)

- **Deskripsi:** Menjadwalkan work item langsung di grid Timeline dengan membuat, memindahkan, dan mengubah panjang bar yang mengisi field Start date dan Due date.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Edit Issues dan Schedule Issues (PERM-021, PERM-024).
- **Aturan perilaku (acceptance criteria):**
  1. Klik-dan-drag di grid pada baris tanpa tanggal membuat bar dan mengisi Start date (titik awal drag) dan Due date (titik akhir drag); klik tunggal tanpa drag membuat bar satu satuan zoom (mis. satu minggu pada zoom Weeks).
  2. Drag seluruh bar menggeser Start date dan Due date sekaligus dengan durasi tetap.
  3. Drag tepi kiri atau kanan bar mengubah hanya Start date atau hanya Due date.
  4. Selama drag, tooltip menampilkan rentang tanggal hasil (format tanggal sesuai locale pengguna) dan durasi dalam hari.
  5. Posisi bar selalu snap ke satuan hari; pada zoom Quarters snap tetap per hari tetapi pergerakan minimum mengikuti lebar piksel satu hari.
  6. Epic tanpa tanggal sendiri menampilkan bar rolled-up/inferred (gaya bergaris/transparan) dari Start date child terawal sampai Due date child terakhir; men-drag bar inferred mengonversinya menjadi tanggal eksplisit Epic setelah konfirmasi.
  7. Child yang dijadwalkan by sprint (PLAN-003) menampilkan bar mengikuti rentang sprint dan tidak dapat di-drag; tooltip menjelaskan "Dijadwalkan oleh sprint <nama>".
  8. Tanggal juga dapat diedit di issue detail panel; perubahan di panel memperbarui bar tanpa reload.
  9. Validasi: Due date tidak boleh sebelum Start date; drag yang menghasilkan kondisi ini dikunci pada batas, dan input manual ditolak dengan pesan "Due date tidak boleh sebelum Start date".
  10. Perubahan langsung tersimpan ke work item (tanpa mode draft seperti Plans), tercatat di history work item (CMT-022, CMT-023), memicu event domain `issue.updated`, dan tersebar realtime ke pengguna lain yang membuka Timeline yang sama.
  11. Bar read-only (kursor default, tanpa handle) bila pengguna tidak memiliki Edit Issues atau Schedule Issues pada work item tersebut; server menolak update dengan HTTP 403 bila dipaksa.
  12. Bar child yang berada di luar rentang tanggal Epic tetap tampil penuh dan dapat memicu warning (PLAN-009).
  13. ⚠ Item yang hanya memiliki Start date tampil sebagai bar satu hari dengan ujung kanan memudar; item yang hanya memiliki Due date tampil sebagai bar satu hari berakhir di Due date; men-drag tepi yang hilang mengisi field yang kosong.
- **Team-managed vs company-managed:** Company-managed: mengubah tanggal memerlukan Schedule Issues dari permission scheme. Team-managed: mengikuti peran proyek yang berhak mengedit work item.
- **Catatan & ketergantungan:** Field ISS-021; permission PERM-024; history di [08](08-komentar-lampiran-aktivitas.md); event realtime di modul 27/28. Pembaruan dikirim sebagai PATCH parsial agar tidak menimpa field lain yang diubah bersamaan.

### PLAN-005 Timeline inline create & hierarchy

- **Deskripsi:** Membuat Epic dan child work item langsung dari Timeline serta menampilkan hierarki parent-child.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Create Issues (PERM-021).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "+ Create Epic" di bawah daftar membuka baris input inline; mengetik summary lalu Enter membuat Epic, Esc membatalkan, dan input tetap terbuka untuk membuat Epic berikutnya.
  2. Tombol "+" pada baris Epic membuka baris input child di bawah Epic tersebut dengan pemilih work type child; child yang dibuat otomatis memiliki Parent = Epic itu.
  3. Pilihan work type mengikuti work type scheme proyek (Story, Task, Bug, custom level standar); work type default adalah work type default scheme.
  4. Item yang dibuat inline langsung tampil di posisi rank yang sesuai (di akhir Epic/daftar) tanpa reload dan langsung menjadi item terpilih.
  5. Bila work type memiliki field wajib selain Summary yang belum terisi, sistem membuka dialog Create lengkap dengan Summary, work type, dan Parent sudah terisi.
  6. ⚠ Hierarki standar adalah Epic > work item standar > Sub-task; Timeline proyek menampilkan Epic, child, dan subtask di bawah child-nya (PLAN-002 butir 13); subtask baru dapat dibuat dari tombol "+" pada baris child, dengan Parent otomatis = child tersebut.
  7. ⚠ Tier premium: bila level di atas Epic tampil (PLAN-002 butir 12), tersedia "+ Create <nama level>" dan "+" pada baris level atas untuk membuat Epic child.
  8. Tombol atau menu "Expand all" / "Collapse all" membuka atau menutup semua baris berhierarki sekaligus.
  9. ⚠ Team-managed dengan fitur Epic dinonaktifkan (work type Epic tidak tersedia) tidak menampilkan baris Epic maupun "+ Create Epic"; Timeline menampilkan work item level standar sebagai baris level atas.
  10. ⚠ Bulk update di dalam Timeline: pengguna memilih beberapa baris (checkbox, Shift-click untuk rentang, termasuk baris subtask) lalu bar aksi massal menyediakan perubahan Status, Assignee, Start date/Due date, Sprint, Parent, dan field editable lain; perubahan diterapkan tanpa meninggalkan view, item yang gagal (izin, validasi, transisi tidak valid) dilaporkan per item, dan pilihan di atas batas aksi cepat diarahkan ke wizard Bulk Change (ISS-035). Membutuhkan Edit Issues dan Make Bulk Changes (PERM-011).
- **Team-managed vs company-managed:** Company-managed: work type yang tersedia mengikuti issue type scheme. Team-managed: mengikuti work type proyek.
- **Catatan & ketergantungan:** Validasi create sama dengan ISS create (modul 06) dan field configuration (modul 05); Parent field CFG-013; hierarki CFG-006.

### PLAN-006 Timeline ranking & re-parenting

- **Deskripsi:** Mengurutkan ulang baris dan memindahkan child antar Epic dengan drag-and-drop di panel daftar Timeline.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Schedule Issues/Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Drag baris Epic ke posisi lain mengubah rank Epic; urutan baru tercermin di backlog, panel Epics, dan board.
  2. Drag child di dalam Epic yang sama mengubah rank child.
  3. Drag child ke Epic lain mengganti Parent child tersebut dan menempatkannya di posisi drop; perubahan Parent tercatat di history.
  4. ⚠ Menu baris menyediakan "Move" dengan opsi Move up, Move down, Move to top, Move to bottom sebagai alternatif drag yang dapat diakses keyboard.
  5. Ranking memerlukan Schedule Issues (company-managed) atau hak edit work item (team-managed); tanpa izin, handle drag tidak muncul.
  6. Rank bersifat global (field Rank berbasis LexoRank, ISS-022) sehingga perubahan di Timeline memengaruhi urutan di board dan backlog yang memakai urutan Rank.
  7. ⚠ Drag ranking dinonaktifkan dengan tooltip penjelasan saat pencarian teks aktif atau filter menyembunyikan sebagian sibling, sehingga posisi drop ambigu; re-parenting via menu tetap tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission.
- **Catatan & ketergantungan:** Rank ISS-022 dan board ranking BRD-017; Parent ISS-025. Operasi rank memakai endpoint rank yang sama dengan backlog (modul 11).

### PLAN-007 Timeline row actions & issue color

- **Deskripsi:** Menu aksi per baris/bar di Timeline (meatball menu) untuk aksi cepat seperti mengubah warna Epic, menghapus tanggal, menambah dependency, dan mengganti parent.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim; setiap aksi mengikuti permission terkait (Edit, Schedule, Link Issues).
- **Aturan perilaku (acceptance criteria):**
  1. Menu "..." tersedia pada hover baris di panel kiri dan pada klik kanan/hover bar di grid.
  2. "Change color" menampilkan palet warna tetap (14 warna) untuk mengubah field Issue color Epic (CFG-020); warna baru langsung diterapkan pada bar.
  3. Warna Epic yang sama dipakai sebagai badge Epic di kartu board dan baris backlog (BRD-038).
  4. ⚠ "Remove dates" mengosongkan Start date dan Due date item setelah konfirmasi; baris kembali menjadi baris tanpa bar.
  5. ⚠ "Add dependency" membuka submenu "Blocks" atau "Is blocked by" lalu pemilih work item untuk membuat issue link (PLAN-008).
  6. ⚠ "Change parent" membuka pemilih Epic (dengan pencarian) atau opsi "Remove parent" bagi child.
  7. Opsi "Move" (Move up/down/to top/to bottom) mengubah urutan rank sesuai PLAN-006.
  8. Item menu hanya muncul bila pengguna memiliki permission aksi tersebut; bila tidak ada aksi yang diizinkan, tombol "..." tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue color CFG-020, card colors BRD-038, link ISS-029.

### PLAN-008 Timeline dependencies

- **Deskripsi:** Membuat, menampilkan, dan menghapus dependency antar work item di Timeline; dependency disimpan sebagai issue link `blocks` / `is blocked by`.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Link Issues (PERM-024); Viewer hanya melihat garis.
- **Aturan perilaku (acceptance criteria):**
  1. Hover di ujung kiri atau kanan bar memunculkan titik konektor; drag dari konektor ke bar lain membuat dependency, dengan bar tujuan disorot selama drag.
  2. Dependency bertipe finish-to-start dan disimpan sebagai issue link tipe "Blocks" (item asal `blocks` item tujuan); tidak ada tipe dependency lain di Timeline proyek.
  3. Garis dependency digambar dari ujung kanan bar pemblok ke ujung kiri bar yang diblok; garis dapat disembunyikan lewat view settings (PLAN-010).
  4. Garis berwarna merah (warning) bila item yang diblok dijadwalkan mulai sebelum item pemblok selesai (Start date item diblok < Due date pemblok), disebut dependency conflict.
  5. Klik garis membuka popover detail dependency berisi kedua item, status konflik, tombol "Remove dependency", dan tautan untuk membuka masing-masing item.
  6. Dependency dapat dibuat antar Epic, antar child, dan antara Epic dan child (termasuk lintas Epic).
  7. Dependency yang sama juga dapat dikelola lewat bagian Linked work items di issue detail (ISS-029); perubahan di sana tercermin di Timeline.
  8. Bila salah satu item tidak memiliki tanggal, link tetap tersimpan tetapi garis tidak dirender; detail panel tetap menampilkan link.
  9. ⚠ Dependency ke item di luar Timeline (tidak lolos filter, proyek lain, atau tanpa izin lihat) tidak digambar; bar yang memilikinya menampilkan ikon kecil "dependency eksternal" dengan jumlah link tersembunyi.
  10. Membuat dependency memerlukan issue linking aktif di instance (CFG-056) dan permission Link Issues; bila issue linking dinonaktifkan, konektor tidak muncul dan aksi menu "Add dependency" disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue link ISS-029, link types CFG-056. Membuat link duplikat (pasangan dan arah sama) ditolak dengan pesan "Dependency ini sudah ada"; link yang membentuk siklus langsung (A blocks B dan B blocks A) diizinkan tetapi keduanya ditandai konflik.

### PLAN-009 Timeline progress & warnings

- **Deskripsi:** Indikator progres pada baris Epic berdasarkan status child dan peringatan masalah jadwal di Timeline proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Baris Epic menampilkan progress bar tiga segmen berdasarkan proporsi jumlah child per status category (To Do / In Progress / Done); Epic tanpa child menampilkan progress kosong.
  2. Hover progress bar menampilkan tooltip jumlah child per status category, mis. "Done 4, In Progress 2, To Do 1".
  3. Toggle "Progress" di view settings menampilkan atau menyembunyikan progress bar untuk semua baris.
  4. Warning dependency conflict ditampilkan sebagai ikon peringatan di bar dan garis merah bila item yang diblok mulai sebelum pemblok selesai.
  5. ⚠ Warning ditampilkan pada bar child yang dijadwalkan (sebagian atau seluruhnya) di luar rentang tanggal eksplisit parent Epic, dengan tooltip "Dijadwalkan di luar tanggal parent".
  6. ⚠ Pada mode planning by sprint, warning ditampilkan bila child berada di sprint yang berakhir setelah Due date Epic atau sprint sudah selesai sementara child belum Done.
  7. ⚠ Toggle "Warnings" di view settings menyembunyikan ikon dan pewarnaan warning tanpa menghapus deteksinya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status category dari modul 09. Katalog warning yang lebih lengkap untuk Plans ada di PLAN-053; perhitungan warning dilakukan di sisi klien dari data yang sudah dimuat agar tidak menambah query.

### PLAN-010 Timeline view settings & filters

- **Deskripsi:** Kontrol tampilan, overlay field, dan filter di Timeline proyek untuk mempersempit informasi yang dirender, termasuk menyimpan view settings sebagai view bersama proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (penyesuaian personal); Admin Proyek (menyimpan view bersama).
- **Aturan perilaku (acceptance criteria):**
  1. Panel "View settings" menyediakan toggle Dependencies, Progress, dan Warnings.
  2. Filter item selesai: opsi tampilkan completed dalam 1, 3, 6, atau 12 bulan terakhir, tampilkan semua, atau sembunyikan semua item ber-status category Done; default 3 bulan.
  3. View settings menyediakan tombol Expand all / Collapse all.
  4. ⚠ View settings menyediakan pilihan field overlay yang tampil di baris/bar: status, assignee, key; masing-masing dapat dinyalakan/dimatikan.
  5. ⚠ View settings menyediakan toggle tampilkan/sembunyikan baris Sprints dan release markers (PLAN-012).
  6. Toolbar filter menyediakan filter Assignee (avatar), Label, Status category, Work type, dan Epic/Parent; filter antar field berlogika AND, nilai dalam satu field berlogika OR.
  7. Company-managed: tersedia filter Version/Release dan Component, serta Quick filters dari board sumber.
  8. Pencarian teks summary/key tersedia di toolbar (sama dengan PLAN-002 butir 8).
  9. Tombol "Clear filters" muncul bila ada filter aktif, beserta indikator jumlah filter aktif.
  10. Filter dan setting personal disimpan per pengguna per proyek/board di preferensi server sehingga bertahan antar sesi dan perangkat.
  11. ⚠ Admin Proyek dapat memilih "Save view settings" > "Save view" untuk menyimpan kombinasi filter dan view settings sebagai view default bersama untuk semua anggota proyek; tombol ini tidak tampil bagi non-admin.
  12. Penyesuaian personal oleh member tidak menimpa view bersama yang tersimpan.
  13. "Reset" mengembalikan Timeline pengguna ke konfigurasi view bersama yang tersimpan (atau default sistem bila belum ada).
  14. ⚠ Indikator "Unsaved changes" tampil di tombol view settings bila konfigurasi saat ini berbeda dari view tersimpan.
  15. ⚠ Filter Timeline berlaku juga pada baris subtask (PLAN-002 butir 13); subtask yang tidak lolos filter disembunyikan, sedangkan parent-nya tetap tampil sebagai konteks bila ada child/subtask yang lolos.
  16. ⚠ Aksi "Select all" pada hasil filter memilih semua baris yang terlihat sebagai masukan bulk update di dalam view (PLAN-005 butir 10); baris yang tersembunyi oleh filter tidak ikut terpilih.
- **Team-managed vs company-managed:** Butir 7 hanya company-managed. Company-managed menyimpan view bersama per board; team-managed per proyek.
- **Catatan & ketergantungan:** View Timeline proyek berbeda dari saved view Plans (PLAN-043) dan saved filter JQL (modul 13). Penyimpanan di tabel `timeline_view_configs` (scope: personal/shared, owner, board/proyek).

### PLAN-011 Timeline zoom & navigation

- **Deskripsi:** Kontrol skala waktu dan navigasi horizontal pada grid Timeline.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Pemilih zoom menyediakan Weeks, Months, dan Quarters; mengganti zoom mempertahankan tanggal di tengah viewport.
  2. Tombol "Today" men-scroll grid sehingga hari ini berada di tengah viewport.
  3. Grid dapat di-scroll horizontal dengan trackpad, shift+scroll roda mouse, atau drag area kosong grid.
  4. ⚠ Bila bar suatu baris berada di luar viewport, tampil indikator panah di tepi kiri/kanan baris; klik panah men-scroll ke bar tersebut.
  5. Header menampilkan satuan sesuai zoom: Weeks (tanggal mulai minggu dan bulan), Months (nama bulan dan tahun), Quarters (Q1-Q4 dan tahun); awal minggu mengikuti preferensi locale pengguna.
  6. Pilihan zoom diingat per pengguna per Timeline.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Preferensi per pengguna di modul 21. Navigasi keyboard (panah kiri/kanan untuk scroll) mengikuti pedoman aksesibilitas modul 27.

### PLAN-012 Timeline sprints & release markers

- **Deskripsi:** Baris header tambahan yang menampilkan sprint dan release (version) di Timeline sebagai konteks penjadwalan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Baris "Sprints" di header (board/proyek Scrum) menampilkan sprint active dan future yang memiliki tanggal sebagai blok sesuai rentangnya; sprint closed dalam rentang tampilan ditampilkan dengan gaya pudar.
  2. Hover sprint menampilkan nama, tanggal mulai-selesai, dan status sprint (active/future/closed).
  3. ⚠ Klik sprint menyorot rentang tanggalnya sebagai pita vertikal di seluruh grid; klik lagi atau Esc menghapus sorotan.
  4. Release markers (ikon diamond dengan garis vertikal) tampil untuk version yang memiliki Release date.
  5. Hover marker menampilkan nama release, Release date, dan status (released/unreleased, serta Overdue bila lewat).
  6. Beberapa release pada tanggal sama digabung menjadi satu marker bertanda jumlah; klik membuka daftar release tersebut.
  7. Team-managed: baris Sprints hanya tampil bila fitur Sprints aktif dan markers hanya tampil bila fitur Releases aktif.
  8. Company-managed: sprint diambil dari board sumber Timeline; release diambil dari proyek-proyek yang tercakup filter board.
- **Team-managed vs company-managed:** Lihat butir 7 dan 8.
- **Catatan & ketergantungan:** Sprint dari modul 11, version dari REL-023 (modul 15). Archived version tidak ditampilkan.

### PLAN-013 Timeline export & share

- **Deskripsi:** Membagikan Timeline lewat tautan atau mengekspornya sebagai gambar.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects; penerima share harus memiliki akses proyek.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Menu "Export" > "Export as image (PNG)" menghasilkan gambar Timeline yang mencakup panel daftar dan grid sesuai filter dan view settings aktif.
  2. Dialog ekspor menyediakan pilihan rentang tanggal: sesuai tampilan saat ini, preset (bulan ini, quarter ini, 6 bulan, 12 bulan), atau custom start-end; custom dengan end sebelum start ditolak.
  3. Tombol "Share" menyediakan "Copy link" dan kirim ke pengguna/email dengan pesan opsional (notifikasi via modul 19).
  4. ⚠ Tautan share membawa filter aktif dan zoom melalui parameter URL sehingga penerima melihat tampilan yang sama.
  5. ⚠ Timeline dapat di-embed di wiki eksternal melalui Smart Link/oEmbed (modul 25) yang menampilkan Timeline read-only.
  6. Penerima tautan atau embed hanya melihat work item sesuai permission dan issue security masing-masing; pengguna tanpa akses proyek melihat halaman "Anda tidak memiliki akses".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Render PNG dilakukan di klien (canvas) agar tidak membutuhkan headless browser di server; batas gambar maksimum 16.000 px lebar dengan pesan "Rentang terlalu besar untuk diekspor, persempit rentang". Ekspor terkait di modul 25.

### PLAN-014 Timeline limits & performance

- **Deskripsi:** Batasan jumlah item dan perilaku saat batas terlampaui di Timeline proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna; sistem menegakkan batas.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Timeline merender maksimal 5.000 work item (termasuk child) per tampilan; batas disimpan sebagai konfigurasi sistem yang dapat diubah Admin Situs.
  2. ⚠ Jumlah item level atas (Epic dan level di atasnya) dibatasi 500 per tampilan; item di atas batas tidak dimuat.
  3. Saat batas terlampaui, banner peringatan tampil "Timeline menampilkan sebagian item karena batas tampilan terlampaui" beserta saran memakai filter atau menyembunyikan item selesai.
  4. Child dimuat lazy saat Epic di-expand pertama kali (endpoint terpisah per parent) dan di-cache selama sesi.
  5. Item selesai lebih lama dari 3 bulan disembunyikan secara default (PLAN-010 butir 2) untuk menjaga performa.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Anggaran performa (render awal dan interaksi drag) ditetapkan di modul 27; grid memakai virtualisasi baris dan kolom.

### PLAN-015 Business project Timeline

- **Alias inventaris:** Timeline view (business)

- **Deskripsi:** Varian Timeline (Gantt sederhana) untuk business project yang menjadwalkan semua work item berdasarkan Start date dan Due date, bukan hanya Epic.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna Tim Bisnis, manajer proyek bisnis; permission Edit/Schedule Issues untuk mengubah jadwal.
- **Aturan perilaku (acceptance criteria):**
  1. Timeline business menampilkan semua work item proyek (bukan hanya Epic) sebagai baris, termasuk hierarki parent-child bila ada.
  2. Drag bar menggeser Start date/Due date sesuai aturan PLAN-004.
  3. Work item dapat dibuat inline dari Timeline (baris "+ Create" di bawah daftar dan "+" pada baris parent).
  4. Toolbar filter menyediakan assignee, status, work type, label, dan field lain (termasuk Category).
  5. ⚠ Dependencies (issue link `blocks`) dan warnings berperilaku sama dengan Timeline software (PLAN-008, PLAN-009).
  6. ⚠ Sub-task ditampilkan di bawah parent-nya (berbeda dengan Timeline software) dan dapat dijadwalkan.
  7. Tidak ada baris Sprints karena business project tidak memakai sprint; release markers juga tidak tampil.
  8. Bar dijadwalkan berdasarkan Start date dan Due date dengan drag pada badan bar untuk memindahkan dan drag tepi untuk resize.
  9. Hierarki parent → child ditampilkan dengan expand/collapse per baris parent.
  10. ⚠ Dependency antar item ditampilkan sebagai panah dari bar pemblok ke bar yang diblok.
  11. Zoom tersedia dalam Weeks, Months, dan Quarters (PLAN-011).
  12. Item dapat dibuat inline dengan memilih work type business proyek; field wajib lain memicu dialog Create lengkap.
  13. Filter diterapkan seperti butir 4, dan warna bar mengikuti warna Category work item (abu-abu bila tanpa Category).
  14. Item tanpa tanggal ditampilkan tanpa bar dan dapat diberi tanggal dengan klik-drag di grid barisnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk business project; company-managed business project tidak memiliki board sumber sehingga sumber data selalu seluruh proyek.
- **Catatan & ketergantungan:** Category field dan business project type di [26](26-kerja-bisnis-formulir-persetujuan.md). Implementasi memakai komponen grid yang sama dengan PLAN-002 dengan konfigurasi "semua level, tanpa sprint".

## List view proyek

### PLAN-016 Project List view

- **Alias inventaris:** List view (business)

- **Deskripsi:** Tampilan tabel bergaya spreadsheet untuk melihat dan mengedit work item proyek secara inline; tersedia di business dan software project. Fitur ini menggabungkan List view inti proyek dan perilaku List business project.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dan Pengguna Tim Bisnis (Browse Projects untuk melihat, Edit Issues untuk inline edit).
- **Aturan perilaku (acceptance criteria):**
  1. List menampilkan tabel work item proyek dengan kolom field (default: Work type, Key, Summary, Status, Assignee, Due date, Priority, Created).
  2. Sel field yang dapat diedit menjadi editor inline saat diklik; Enter/klik luar menyimpan, Esc membatalkan, dan galat validasi tampil di bawah sel tanpa mengubah nilai.
  3. Data dimuat bertahap dengan infinite scroll (halaman 50 baris) dan virtualisasi baris; total jumlah item ditampilkan di toolbar.
  4. Klik key/summary membuka issue detail panel di sisi kanan tanpa meninggalkan List (opsi buka halaman penuh).
  5. ⚠ Pengaturan density baris tersedia: Comfortable dan Compact; pilihan disimpan per pengguna.
  6. Kolom dapat dikonfigurasi: tambah, hapus, dan reorder field, termasuk custom field dan Category (detail di PLAN-017).
  7. Inline edit tersedia minimal untuk summary, status (menjalankan transisi workflow yang valid), assignee, due date, priority, labels, dan category; status yang tidak memiliki transisi valid tidak muncul di pilihan.
  8. Work item dapat dibuat inline di baris paling bawah dan child dapat dibuat dari menu baris parent (detail di PLAN-019).
  9. Hierarki parent → child/subtask ditampilkan dengan tombol expand/collapse per baris parent (detail di PLAN-018).
  10. Sort per kolom, filter (assignee, status, work type, category, label, due date), dan pencarian teks tersedia di toolbar (detail di PLAN-018 dan PLAN-020).
  11. ⚠ Kriteria filter List dapat disimpan sebagai saved filter (JQL) melalui "Save as filter" yang membuka filter di modul 13.
  12. Bulk select via checkbox per baris dan "select all" pada halaman, lalu bulk action (detail di PLAN-019).
  13. ⚠ Export CSV tersedia dari List dengan kolom sesuai tampilan; bila tidak tersedia di suatu konteks, tersedia tautan "Open in search" ke navigator untuk ekspor.
  14. Header kolom menyediakan "Create field" bagi Admin Proyek untuk membuat custom field baru yang langsung ditambahkan sebagai kolom.
  15. Untuk proyek besar, data dimuat bertahap (infinite scroll) tanpa memblokir interaksi; baris yang belum dimuat ditampilkan sebagai skeleton.
  16. ⚠ List dan halaman All work proyek (Project Work items page, SRC-032) digabung menjadi satu view List per space; tautan lama ke All work me-redirect ke List dengan filter yang setara.
  17. ⚠ Saved view List dibatasi 10 per space dan hanya dapat dibuat, diubah, atau dihapus oleh admin space (Admin Proyek); anggota lain hanya dapat memilih saved view dan menyesuaikannya secara personal. Detail penyimpanan view di PLAN-017 butir 5-8.
- **Team-managed vs company-managed:** Team-managed: "Create field" menambah field ke work type proyek. Company-managed: "Create field" hanya bagi Admin Situs atau menampilkan tautan ke administrasi field; editabilitas mengikuti screen/field configuration.
- **Catatan & ketergantungan:** Berbagi komponen tabel dengan List view hasil pencarian (SRC-023) dan Project Work items page (SRC-032). Category field di modul 26. Inline edit memakai endpoint update work item yang sama dengan issue view (modul 06).

### PLAN-017 List columns & fields

- **Deskripsi:** Konfigurasi kolom di List view serta penyimpanan konfigurasi List sebagai view bersama proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (personal); Admin Proyek (menyimpan view bersama).
- **Aturan perilaku (acceptance criteria):**
  1. Menu kolom menyediakan pencarian field untuk menambah kolom dan opsi "Remove column" untuk menghapusnya; Summary tidak dapat dihapus.
  2. Kolom dapat di-reorder dengan drag header dan di-resize dengan drag tepi header (lebar minimum 80 px).
  3. Custom field (semua tipe yang dapat ditampilkan) dapat ditambahkan sebagai kolom; field yang tidak berlaku untuk work type suatu baris tampil kosong dan tidak dapat diedit.
  4. ⚠ Pengaturan kolom personal disimpan per pengguna per proyek sampai di-reset.
  5. Tombol "Save or reset view settings" > "Save list" menyimpan filter, kolom/field, group by, sort, status hierarki, dan density sebagai view bersama untuk semua anggota proyek.
  6. Hanya Admin Proyek yang dapat menyimpan view bersama; member biasa hanya menyesuaikan view secara personal tanpa menimpa versi tersimpan.
  7. "Reset" mengembalikan List pengguna ke konfigurasi view bersama tersimpan.
  8. ⚠ Proyek dapat memiliki hingga 10 saved list view bernama dengan switcher di toolbar; percobaan menyimpan view ke-11 ditolak dengan pesan "Batas jumlah view tercapai".
  9. ⚠ Field atau nilai filter yang direferensikan view tersimpan lalu dihapus diabaikan saat view dibuka, dengan pemberitahuan "Beberapa kolom/filter tidak lagi tersedia dan dilewati".
- **Team-managed vs company-managed:** Tidak ada perbedaan selain siapa yang dapat membuat field.
- **Catatan & ketergantungan:** View List berbeda dari saved filter JQL (modul 13) dan saved view Plans (PLAN-043). Disimpan di tabel `list_view_configs`.

### PLAN-018 List group by, sort & hierarchy

- **Deskripsi:** Grouping, sorting, dan tampilan hierarki di List view.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Group by tersedia untuk Status, Assignee, Priority, Parent, Sprint (software), Work type, dan Category (business); header grup menampilkan nama dan jumlah item.
  2. Klik header kolom mengurutkan asc, klik kedua desc, klik ketiga kembali ke urutan default (Rank); sort berlaku di dalam setiap grup.
  3. Child dan sub-task dapat ditampilkan di bawah parent dengan expand; saat sort aktif, child diurutkan di dalam parent-nya.
  4. Setiap grup dapat di-collapse/expand; status collapse disimpan per pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Group by sejalan dengan SRC-034 (modul 13).

### PLAN-019 List inline create & bulk actions

- **Deskripsi:** Membuat item inline dan aksi massal di List view.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Create Issues untuk create; Edit/Delete Issues dan Make Bulk Changes (PERM-011) untuk aksi massal.
- **Aturan perilaku (acceptance criteria):**
  1. Baris "+ Create" tersedia di bawah tabel dan di bawah setiap grup; item yang dibuat di grup otomatis mendapat nilai field grup (mis. Status atau Assignee grup tersebut).
  2. Baris inline create menyediakan pemilih work type; field wajib lain memicu dialog Create lengkap.
  3. Multi-select baris via checkbox, shift-click untuk rentang, dan "select all" untuk item yang sudah dimuat; jumlah terpilih tampil di bar aksi.
  4. ⚠ Bar aksi massal menyediakan edit field (assignee, status, priority, labels, due date, dan field lain yang editable) serta delete dengan konfirmasi jumlah; item yang gagal diproses dilaporkan per item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bulk Change ISS-035 (modul 06) untuk operasi di atas 50 item atau operasi move; List mengarahkan ke wizard Bulk Change bila pilihan melebihi batas aksi cepat.

### PLAN-020 List filters & export

- **Deskripsi:** Filter dan ekspor di List view.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian teks memfilter berdasarkan summary dan key (dan deskripsi bila didukung mesin pencari).
  2. Basic filter menyediakan pemilih nilai untuk field umum (assignee, status, work type, label, priority, due date, created) dengan kombinasi AND antar field.
  3. ⚠ Tombol "Switch to JQL" mengubah filter aktif menjadi kueri JQL yang dapat diedit (modul 12); kueri yang tidak dapat direpresentasikan sebagai basic filter tetap berada di mode JQL.
  4. ⚠ Pemilih "Saved filters" menerapkan filter tersimpan yang dapat diakses pengguna, dibatasi ke work item proyek ini.
  5. ⚠ "Export CSV" mengekspor item hasil filter dengan kolom tampilan saat ini, maksimal 1.000 baris per ekspor dengan pesan arahan ke navigator untuk jumlah lebih besar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JQL di modul 12, saved filter dan ekspor di modul 13.

### PLAN-021 List conditional formatting

- **Deskripsi:** Aturan format bersyarat di List view yang mewarnai baris work item yang memenuhi kriteria field tertentu sehingga item berisiko atau penting langsung terlihat tanpa filter.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek atau pengguna dengan hak edit proyek (mengelola rule); semua pengguna dengan akses List (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Menu List view menyediakan "Conditional formatting" yang membuka panel pengelolaan rule.
  2. Kondisi rule berbasis nilai field (status, priority, assignee, due date, labels, custom field) dengan operator is, is not, contains, is empty, before/after tanggal, dan relatif (mis. due date dalam N hari); satu rule dapat memiliki beberapa kondisi AND.
  3. Warna highlight baris dipilih dari palet tetap (8 warna) yang memenuhi kontras teks di light dan dark mode.
  4. Beberapa rule memiliki urutan prioritas yang dapat diubah dengan drag; bila beberapa rule cocok, rule teratas yang menang.
  5. Setiap rule dapat di-enable/disable, edit, duplicate, dan delete.
  6. ⚠ Rule tersimpan per proyek per view List dan berlaku untuk semua pengguna yang membuka view tersebut.
  7. ⚠ Hanya Admin Proyek atau pengguna dengan Edit Issues di proyek yang dapat mengelola rule; pengguna lain hanya melihat hasil pewarnaan.
  8. Rule dievaluasi ulang secara realtime saat nilai field berubah lewat inline edit atau pembaruan dari pengguna lain.
  9. Pewarnaan tetap berlaku bersama group by, sort, dan filter.
  10. ⚠ Maksimal 20 rule per view; menambah rule ke-21 ditolak dengan pesan "Batas jumlah rule tercapai".
  11. ⚠ Tersedia di business dan software project, baik team-managed maupun company-managed.
- **Team-managed vs company-managed:** Tidak ada perbedaan (lihat butir 11).
- **Catatan & ketergantungan:** Disimpan di tabel `list_format_rules`; evaluasi di klien memakai nilai yang sudah dimuat. Rule tidak memengaruhi ekspor CSV.

## Calendar proyek

### PLAN-022 Project Calendar view

- **Alias inventaris:** Calendar view (business)

- **Deskripsi:** Tampilan kalender proyek yang menempatkan work item berdasarkan Start date dan Due date; tersedia di business dan software project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dan Pengguna Tim Bisnis; mengubah tanggal memerlukan Edit/Schedule Issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Calendar menyediakan tampilan bulanan (default) dan mingguan.
  2. Item dengan Start date dan Due date tampil sebagai bar multi-hari yang melintasi sel; item yang hanya memiliki Due date tampil pada satu hari.
  3. ⚠ Drag item ke tanggal lain menggeser tanggal dengan durasi tetap; resize tepi bar mengubah rentang Start date/Due date.
  4. Klik sel tanggal kosong membuka create inline dengan Due date (dan Start date) terisi tanggal tersebut.
  5. Tombol "Today" dan navigasi bulan/minggu sebelum dan sesudah tersedia di toolbar.
  6. Klik item membuka issue detail panel.
  7. Hari dengan item melebihi tinggi sel menampilkan tautan "+N more" yang membuka popover daftar item hari tersebut.
  8. Drag, resize, dan create mengikuti permission Edit/Schedule/Create Issues; tanpa izin, item tidak dapat di-drag.
  9. Work item tampil sebagai bar dari Start date sampai Due date, atau di Due date saja bila Start date kosong (sama dengan butir 2, berlaku juga di business project).
  10. Drag item memindahkan tanggal dan resize bar mengubah durasi; validasi Due date tidak boleh sebelum Start date berlaku.
  11. Klik tanggal membuat work item dengan tanggal terisi otomatis; work type default mengikuti work type default proyek.
  12. Panel "Unscheduled" berisi item tanpa tanggal yang dapat di-drag ke kalender (detail di PLAN-023).
  13. Toolbar menyediakan pilihan tampilan bulan/minggu dan tombol Today.
  14. Filter tersedia untuk assignee, status, work type, dan category (detail di PLAN-024).
  15. Klik item membuka issue detail.
  16. ⚠ Opsi view menampilkan atau menyembunyikan item Done.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Subscribe/iCal feed ke kalender eksternal bukan fitur native (lihat Di luar cakupan). Gadget kalender dashboard di modul 17.

### PLAN-023 Calendar unscheduled panel

- **Deskripsi:** Panel samping berisi work item tanpa tanggal yang dapat di-drag ke kalender.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Schedule Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan item proyek yang Due date-nya kosong, diurutkan berdasarkan Rank, dengan jumlah total di header panel.
  2. Panel memiliki pencarian teks dan filter (assignee, work type, status).
  3. Drag item dari panel ke tanggal mengisi Due date (dan Start date sama dengan tanggal tersebut) lalu item hilang dari panel.
  4. Tombol toggle membuka/menutup panel; status disimpan per pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Panel ini dipakai bersama oleh Plan Calendar view (PLAN-036).

### PLAN-024 Calendar sprints, releases & filters

- **Deskripsi:** Overlay sprint dan release serta filter dan view settings di Calendar.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Software project: sprint ditampilkan sebagai pita di bagian atas minggu yang tercakup rentang sprint, dengan nama sprint.
  2. Release (version) dengan Release date ditampilkan sebagai penanda pada tanggalnya; hover menampilkan nama dan status.
  3. Filter tersedia untuk assignee, work type, status, label, dan field lain; indikator jumlah filter aktif dan "Clear filters".
  4. ⚠ Toggle "Show completed" menampilkan atau menyembunyikan item ber-status category Done.
  5. ⚠ Toggle "Show weekends"; bila dimatikan, kolom Sabtu-Minggu disembunyikan dan item yang hanya jatuh di akhir pekan ditampilkan di hari kerja terdekat sebelumnya dengan penanda.
- **Team-managed vs company-managed:** Team-managed: overlay sprint/release hanya bila fitur Sprints/Releases aktif.
- **Catatan & ketergantungan:** Sprint (modul 11), release (modul 15).

## Summary & Board business project

### PLAN-025 Project Summary page

- **Alias inventaris:** Summary view (business)

- **Deskripsi:** Rujukan singkat. Halaman Summary proyek (business dan software) tidak dirinci di modul ini. Pemilik view (rute, tab, toggle Features, ketersediaan per tipe proyek) adalah PRJ-028 di [04 Proyek](04-proyek.md); pemilik metrik dan perilaku widget adalah RPT-034 di [16 Laporan & Insight](16-laporan-insight.md).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects; Stakeholder/Viewer (lihat PRJ-028).
- **Aturan perilaku (acceptance criteria):**
  1. Kartu metrik 7 hari (completed, updated, created, due soon) mengikuti RPT-034 butir 1.
  2. Status overview (donut dengan total di tengah) mengikuti RPT-034 butir 2.
  3. Recent activity feed mengikuti RPT-034 butir 3.
  4. Priority breakdown (bar chart, termasuk business project) mengikuti RPT-034 butir 4.
  5. Types of work (persentase per work type) mengikuti RPT-034 butir 5.
  6. Team workload (termasuk baris Unassigned) mengikuti RPT-034 butir 6.
  7. Epic/parent progress (software, dan business bila hierarki parent aktif) mengikuti RPT-034 butir 7.
  8. Klik kartu atau segmen chart membuka List/navigator terfilter mengikuti RPT-034 butir 8 dan PRJ-028 butir 3.
  9. Varian business project memakai widget yang sama; tidak ada aturan widget terpisah di modul ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keputusan kepemilikan: view = PRJ-028 ([04](04-proyek.md)), metrik & widget = RPT-034 ([16](16-laporan-insight.md)); bila ada perbedaan, kedua sumber itu yang berlaku dan modul ini tidak menambah aturan. Turunan untuk service space: JSM-PRJ-014 di [29](29-jsm-service-project-request-type-formulir.md) memakai komponen yang sama ditambah kartu khusus request type/SLA. ⚠ Space Insights (RPT-043, [16](16-laporan-insight.md)) adalah panel insight terpisah yang ditautkan dari Summary dan dimiliki modul 16; modul ini tidak mendefinisikannya.

### PLAN-026 Board view (business)

- **Deskripsi:** Board Kanban sederhana untuk business project dengan kolom yang dipetakan ke status workflow, opsi beberapa board view tersimpan, dan cover image pada kartu.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna Tim Bisnis (Transition Issues untuk drag); Admin Proyek (kolom, board view bersama).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap business project memiliki satu board default dengan pemetaan 1 kolom = 1 status.
  2. Drag kartu antar kolom menjalankan transisi workflow; transisi yang tidak valid ditolak (kolom tujuan tampil non-droppable) dengan pesan "Transisi ini tidak tersedia".
  3. Team-managed: menambah kolom membuat status baru di workflow proyek. Company-managed: perubahan status/kolom mengikuti workflow yang dikelola Admin Situs.
  4. Kolom dapat di-rename, di-reorder, dan dihapus; menghapus kolom yang berisi item mewajibkan memilih kolom tujuan untuk memindahkan item.
  5. ⚠ Group by tersedia untuk None, Assignee, Priority, Category, dan Subtask/parent.
  6. Filter toolbar: pencarian, avatar assignee, work type, label, dan category.
  7. Quick create di bagian bawah setiap kolom membuat item dengan status kolom tersebut.
  8. Kartu menampilkan summary, key, assignee, due date, priority, dan warna category.
  9. Board business tidak memiliki sprint, backlog, maupun board filter JQL.
  10. ⚠ Item yang sedang menunggu approval tidak dapat di-drag ke status lain di luar keputusan approval; kartu menampilkan badge "Awaiting approval".
  11. Admin Proyek dapat membuat, mengganti nama, dan menghapus board view tambahan di business project (board view default tidak dapat dihapus).
  12. Setiap board view menyimpan filter aktif, group by, field kartu, workflow terpilih (bila proyek punya beberapa workflow), dan urutan kolom.
  13. Switcher board view tersedia di tab Board untuk berpindah antar view.
  14. ⚠ Hanya Admin Proyek yang dapat menyimpan board view untuk semua anggota; perubahan sementara oleh member tidak menimpa view bersama.
  15. ⚠ Maksimal 10 board view per proyek; membuat view ke-11 ditolak dengan pesan "Batas jumlah board view tercapai".
  16. Multiple board views berbeda dari company-managed software project (banyak board berbasis saved filter, BRD-003); team-managed software project tetap satu board.
  17. ⚠ Attachment gambar dapat dijadikan cover kartu dari issue view atau menu kartu ("Make cover").
  18. Cover dapat diganti atau dihapus; cover hilang otomatis bila attachment-nya dihapus.
  19. ⚠ Pengaturan view board menyediakan toggle tampil/sembunyi cover image.
  20. ⚠ Cover hanya untuk format gambar yang didukung media viewer (PNG, JPEG, GIF, WebP); gambar pertama tidak otomatis menjadi cover.
- **Team-managed vs company-managed:** Butir 3: team-managed membuat status baru, company-managed mengikuti workflow admin.
- **Catatan & ketergantungan:** Fitur Kanban penuh (WIP limit, swimlane, board settings) ada di [10 Board](10-board-kolom-kartu.md) (BRD-013); approvals di modul 26; attachment di modul 08. Multiple board views P2 dan card cover P3 dalam fitur ini dapat dirilis setelah board dasar.

## Plans: fondasi, sumber & akses

### PLAN-027 Plans directory & lifecycle

- **Deskripsi:** Halaman daftar Plans dan siklus hidup plan (create, copy, archive, trash, restore, star). Plans adalah perencanaan lintas proyek yang hanya tersedia di tier premium (Premium/Enterprise) produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Product Owner/PM dan Scrum Master (pembuat/editor plan); Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Sidebar menyediakan menu "Plans" berisi plan starred dan recent serta tautan "View all plans".
  2. Halaman "Plans" menampilkan kolom nama, lead/owner, sumber (jumlah dan tipe), status (active/archived), dan tanggal dibuat/diubah; mendukung pencarian nama dan sort per kolom.
  3. Pengguna dapat star/unstar plan; plan starred tampil paling atas di menu sidebar (modul 21).
  4. Tombol "Create plan" membuka wizard pembuatan (PLAN-028); tombol disembunyikan bila tier situs bukan premium.
  5. ⚠ "Copy plan" menduplikasi settings, issue sources, exclusion rules, teams, dan saved views ke plan baru bernama "Copy of <nama>"; scenario non-default dan perubahan draft yang belum disimpan tidak ikut disalin.
  6. "Archive" memindahkan plan ke tab "Archived plans" (read-only, tidak muncul di sidebar); "Restore" mengaktifkannya kembali.
  7. ⚠ "Move to trash" memindahkan plan ke tab Trash; plan dapat di-restore selama 60 hari, lalu dihapus permanen otomatis; "Delete permanently" segera menghapus setelah konfirmasi mengetik nama plan.
  8. Pemilik/editor dapat me-rename plan dan mengubah plan lead/owner (harus pengguna aktif berlisensi).
  9. Plans hanya dapat diakses pengguna di situs dengan tier premium dan lisensi produk aktif; akses dari tier lain menampilkan halaman upsell tanpa data plan.
  10. Plan tidak menyimpan salinan data work item; data selalu dibaca dari work item kecuali perubahan draft yang belum disimpan (PLAN-045).
- **Team-managed vs company-managed:** Tidak ada perbedaan; plan dapat mencakup kedua jenis proyek.
- **Catatan & ketergantungan:** Gating tier VISI-003 dan PERM-005; administrasi Plans di modul 22. Tabel `plans` (id, name, owner_id, status active/archived/trashed, trashed_at). Trash dibersihkan oleh job terjadwal.

### PLAN-028 Plan creation & issue sources

- **Deskripsi:** Wizard pembuatan plan untuk menentukan sumber work item (boards, projects, filters) dan pengaturan awal.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Pembuat plan (pengguna berlisensi di tier premium); sumber hanya dapat dipilih bila pengguna dapat mengaksesnya.
- **Aturan perilaku (acceptance criteria):**
  1. Langkah pertama: input nama plan (wajib, maks 255 karakter) dan access (Open / Private, PLAN-030).
  2. Pengguna menambah issue sources bertipe Board, Project, atau Filter (saved filter JQL).
  3. Satu plan dapat menggabungkan banyak sumber dari proyek berbeda (cross-project); item yang muncul di beberapa sumber hanya ditampilkan sekali.
  4. ⚠ Maksimal 100 issue sources per plan; menambah sumber ke-101 ditolak dengan pesan "Batas jumlah sumber plan tercapai".
  5. Wizard menampilkan preview jumlah work item dari sumber terpilih sebelum plan dibuat, dengan peringatan bila mendekati batas item (PLAN-055).
  6. Langkah exclusion rules saat create memungkinkan mengecualikan completed work items, work types, statuses, dan releases (PLAN-029).
  7. Sumber Board membawa sprint, team, dan capacity dari board tersebut (board Scrum memberi sprint; board dapat dihubungkan ke team).
  8. Sumber Project menyertakan semua work item proyek.
  9. Sumber Filter menyertakan item secara dinamis mengikuti JQL filter tersimpan; perubahan kueri filter langsung memengaruhi isi plan.
  10. Sumber dapat diubah kemudian via Plan settings > Issue sources (PLAN-031).
  11. Sumber yang dihapus atau tidak dapat diakses pengguna menampilkan error sumber "Sumber tidak tersedia" di Plan settings dan banner di plan, tanpa menggagalkan pemuatan sumber lain.
- **Team-managed vs company-managed:** Board team-managed dapat dipakai sebagai sumber seperti board company-managed.
- **Catatan & ketergantungan:** Saved filter di modul 13, board di modul 10. Tabel `plan_sources` (plan_id, type, ref_id, position).

### PLAN-029 Plan exclusion rules & removed items

- **Deskripsi:** Aturan untuk mengecualikan work item dari plan serta pengelolaan item yang dikeluarkan secara manual.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Rule "Exclude completed work items" mengecualikan item yang selesai lebih dari X hari lalu (X dapat diisi 0-10.000; default 30).
  2. Rule berdasarkan work type mengecualikan semua item dengan work type terpilih.
  3. Rule berdasarkan status atau status category mengecualikan item dengan status terpilih.
  4. Rule berdasarkan release/version mengecualikan item yang Fix versions-nya adalah release terpilih.
  5. Aksi "Remove from plan" per item mengeluarkan item dari plan tanpa menghapus atau mengubah work item.
  6. Plan settings > Removed work items menampilkan daftar item yang dikeluarkan manual dengan opsi "Include again".
  7. Item yang di-exclude atau di-remove tidak dihitung dalam rollup dan capacity.
  8. Perubahan exclusion rules berlaku untuk semua view dan semua scenario plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan di `plan_exclusion_rules` dan `plan_removed_items`. Exclusion dievaluasi sebelum filter view (PLAN-041).

### PLAN-030 Plan permissions & access

- **Deskripsi:** Kontrol siapa yang dapat melihat dan mengedit plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Owner plan, Editor, Viewer, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Access level plan: Open (semua pengguna berlisensi di situs dapat melihat dan mengedit) atau Private (hanya pengguna/grup yang ditambahkan).
  2. Owner/editor dapat menambah pengguna atau grup sebagai Viewer atau Editor.
  3. Owner/lead plan selalu memiliki akses edit dan tidak dapat diturunkan ke Viewer.
  4. Viewer tidak dapat membuat perubahan draft maupun menyimpan ke work item; kontrol edit disembunyikan.
  5. Editor hanya dapat menyimpan perubahan ke work item yang ia miliki permission-nya di proyek terkait; perubahan lain gagal di Review changes dengan pesan permission (PLAN-045).
  6. Item dari proyek yang tidak dapat di-browse pengguna (atau dibatasi issue security) tidak ditampilkan dan tidak dihitung dalam rollup bagi pengguna tersebut.
  7. ⚠ Admin Situs dapat melihat dan mengelola akses semua plan, termasuk plan Private, dari halaman Plans.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model permission plan dirinci di PERM-052. Global permission bergaya "Portfolio user/viewer" adalah konsep Data Center dan tidak dibangun.

### PLAN-031 Plan settings

- **Deskripsi:** Halaman pengaturan plan yang mengumpulkan semua konfigurasi plan dalam satu tempat bernavigasi tab/sidebar.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor dan owner plan (mengubah); Viewer (melihat sebagian).
- **Aturan perilaku (acceptance criteria):**
  1. Bagian Plan details: ubah nama dan owner/lead plan.
  2. Bagian Issue sources: tambah, hapus, dan ubah urutan sumber (PLAN-028).
  3. Bagian Exclusion rules (PLAN-029).
  4. Bagian Removed work items (PLAN-029).
  5. Bagian Permissions (PLAN-030).
  6. Bagian Custom fields (PLAN-040).
  7. Bagian Dependencies: pilihan issue link type yang dianggap dependency (PLAN-048).
  8. Bagian Scheduling: estimation unit, field tanggal, inferred dates (PLAN-033).
  9. Bagian Scenarios: aktifkan/nonaktifkan scenarios (PLAN-046).
  10. Bagian Teams: tautan ke halaman Teams plan (PLAN-050).
  11. Bagian Releases: tautan ke halaman Releases plan (PLAN-051).
  12. Bagian Archive/trash plan dengan konfirmasi (PLAN-027).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengaturan global Plans (mis. mengaktifkan Plans per situs) ada di modul 22. Perubahan settings bukan draft; langsung berlaku dan tercatat di audit plan.

### PLAN-032 Issue type hierarchy configuration

- **Deskripsi:** Konfigurasi level hierarki di atas Epic (mis. Initiative, Theme) yang dipakai Plans, Timeline premium, dan field Parent. Modul ini merinci perilaku yang terlihat dari sisi perencanaan; aturan admin lengkap ada di CFG-006/CFG-007.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs (Administer, PERM-007).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > Work items > Work type hierarchy tersedia untuk semua tier, tetapi penambahan level di atas Epic hanya diizinkan di tier premium.
  2. Admin dapat menambah level baru di atas Epic dan memberi nama (unik, maks 60 karakter).
  3. Satu atau lebih work type dipetakan ke setiap level.
  4. Level bawaan Epic, Story (standar), dan Sub-task tidak dapat dihapus; tombol hapus tidak tersedia.
  5. Setiap work type hanya boleh berada di satu level; memetakan work type yang sudah dipakai di level lain ditolak dengan pesan "Work type sudah dipetakan ke level <nama>".
  6. Field Parent menghubungkan item ke level tepat di atasnya (menggantikan Epic Link dan Parent Link legacy); parent dari level yang tidak tepat ditolak.
  7. Menghapus level menampilkan peringatan jumlah relasi parent terdampak dan, setelah konfirmasi, memutus relasi parent pada level tersebut.
  8. Company-managed: work type level atas harus ada di issue type scheme proyek agar dapat dibuat di proyek tersebut.
  9. ⚠ Team-managed: Epic team-managed dapat memiliki parent level atas (dari proyek company-managed) melalui Parent field dan Plans, tetapi proyek team-managed tidak dapat membuat work type level atas sendiri.
  10. Tier non-premium (free-standard): hierarki tetap 3 level (Epic > Standard > Sub-task).
- **Team-managed vs company-managed:** Lihat butir 8 dan 9.
- **Catatan & ketergantungan:** Sumber kebenaran konfigurasi di CFG-006 dan CFG-007 (modul 05); fitur ini menambahkan dampaknya terhadap Plans dan Timeline (PLAN-002, PLAN-034). Downgrade tier mengikuti CFG-007.

### PLAN-033 Target dates & scheduling fields

- **Deskripsi:** Pengaturan field tanggal dan estimasi yang dipakai plan untuk menjadwalkan item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Field Target start dan Target end tersedia sebagai field khusus perencanaan yang tersimpan di work item (CFG-022) dan dapat diedit dari plan.
  2. Plan settings > Scheduling memilih field tanggal start dan end yang dipakai bar plan: Target start/end, Start date/Due date, atau custom date field; pilihan berlaku per plan.
  3. Opsi inferred dates (on/off) menurunkan tanggal dari sprint atau release bila field tanggal kosong.
  4. Estimation unit dipilih antara story points atau time (jam/hari).
  5. ⚠ Untuk estimation time, plan memakai konversi hours per day dan days per week dari pengaturan time tracking situs (modul 14), default 8 jam dan 5 hari.
  6. Default estimate dapat diisi untuk item tanpa estimate dan dipakai oleh capacity serta auto-scheduler.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Field di CFG-022 (modul 05) dan estimasi di modul 14.

## Plans: view, field & konfigurasi tampilan

### PLAN-034 Plan Timeline view

- **Deskripsi:** Tampilan utama plan berupa daftar hierarki multi-proyek dengan kolom field dan grid timeline untuk perencanaan lintas tim.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan (mengubah sebagai draft), Viewer plan (read-only).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar hierarki menampilkan item dari level teratas (mis. Initiative) hingga Sub-task sesuai filter Hierarchy (PLAN-041).
  2. Kolom field tampil di antara daftar dan grid (PLAN-040) dan dapat disembunyikan seluruhnya.
  3. Bar dapat di-drag/resize untuk mengubah tanggal start/end sesuai field scheduling plan; perubahan menjadi draft.
  4. Bar inferred (dari sprint, release, atau rollup child) ditampilkan dengan gaya berbeda (bergaris/transparan) dan tooltip sumber inferensinya.
  5. Zoom tersedia Weeks, Months, Quarters, Years, dan custom date range.
  6. Tombol "Today" men-scroll ke hari ini.
  7. Saat group by Team, setiap grup team menampilkan baris sprint/iterasi dengan capacity bar (PLAN-050).
  8. Release markers tampil di header (PLAN-051).
  9. Dependency ditampilkan sebagai garis atau badge sesuai view settings (PLAN-042, PLAN-048).
  10. Field atau bar yang memiliki perubahan belum disimpan ditandai highlight (warna draft) sampai disimpan atau dibuang.
  11. Inline create tersedia di level mana pun (PLAN-044).
  12. Klik item membuka issue detail panel di plan dengan opsi "Open in new tab"; edit di panel juga menjadi draft.
  13. Switcher view di header menyediakan perpindahan cepat Timeline / List / Calendar / Program board / Summary dengan mempertahankan filter aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai komponen grid yang sama dengan PLAN-002, ditambah lapisan draft. Data plan dimuat via endpoint agregat plan yang menggabungkan nilai work item dengan change set scenario aktif.

### PLAN-035 Plan List view

- **Deskripsi:** Tampilan spreadsheet di plan tanpa grid timeline untuk pengeditan massal field secara inline.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan (read-only).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar hierarki ditampilkan dengan kolom field selebar layar.
  2. Semua field plan yang didukung dapat diedit inline; perubahan menjadi draft sampai disimpan.
  3. Baris dapat di-rank via drag-and-drop (draft).
  4. Filter, group by, sort, dan saved view yang sama dengan Timeline berlaku.
  5. Hierarki dapat di-expand/collapse per baris dan per level.
  6. Lebar dan urutan kolom dapat diatur dan tersimpan di saved view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ekspor CSV di PLAN-054.

### PLAN-036 Plan Calendar view

- **Deskripsi:** Tampilan kalender untuk work item di plan berdasarkan field tanggal yang dikonfigurasi di plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Kalender bulanan menampilkan item multi-hari sebagai bar sesuai field scheduling plan.
  2. Drag item mengubah tanggal dan menghasilkan perubahan draft.
  3. Panel unscheduled work berisi item plan tanpa tanggal (memakai komponen PLAN-023).
  4. Filter dan view settings plan berlaku di Calendar.
  5. Klik item membuka issue detail panel plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbagi komponen kalender dengan PLAN-022. Ketersediaan dan opsi persis Plan Calendar view di produk referensi belum terkonfirmasi; perilaku di atas adalah rancangan produk ini.

### PLAN-037 Plan Summary view

- **Deskripsi:** Halaman ringkasan plan berisi progres dan indikator kesehatan plan untuk stakeholder.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Viewer plan, Stakeholder/Viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Widget progres menampilkan progres keseluruhan plan dan per level hierarki (berdasarkan jumlah item atau estimate sesuai pengaturan).
  2. Widget breakdown status menampilkan jumlah item per status category.
  3. Widget release menampilkan release mendatang dengan release date dan target date level atas.
  4. Widget warnings menampilkan jumlah warning per jenis dengan tautan ke filter "Show only items with warnings".
  5. ⚠ Widget ringkasan per team menampilkan jumlah item dan progres per team.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Summary dihitung dari nilai scenario aktif termasuk draft, dengan label "Termasuk perubahan belum disimpan" bila ada.

### PLAN-038 Plan Dependencies view

- **Deskripsi:** Tampilan khusus (dependency report) yang memvisualisasikan jaringan dependency lintas tim dan proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Graph mode menampilkan node item dan panah dependency berarah.
  2. List mode menampilkan tabel kolom from, to, link type, status, dan team.
  3. Filter tersedia berdasarkan project, team, sprint, release, dan work type.
  4. Dependency conflict disorot merah di graph dan list.
  5. ⚠ Graph dapat dikelompokkan dalam lane per team atau per proyek.
  6. Klik node membuka issue detail item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi dependency dari PLAN-048.

### PLAN-039 Program board

- **Deskripsi:** Tampilan board di plan yang menampilkan pekerjaan per team dan per sprint/iterasi dalam satu increment, beserta dependency lintas team (gaya PI planning).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Board menampilkan kolom sprint/iterasi dan baris (swimlane) per team.
  2. Kartu menampilkan item pada level hierarki terpilih (mis. Epic/feature).
  3. Garis dependency digambar antar kartu, merah bila konflik.
  4. Drag kartu antar sprint/team mengubah Sprint dan Team sebagai draft.
  5. ⚠ Pengguna memilih increment (rentang sprint) dan filter yang diterapkan.
  6. Setiap sel team × sprint menampilkan indikator capacity (dipakai/total).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada teams/capacity (PLAN-050) dan sprint assignment (PLAN-049). Dapat dirilis setelah view lain stabil. Fitur ini relatif baru di produk referensi sehingga detailnya dapat berubah.

### PLAN-040 Plan fields & custom fields

- **Deskripsi:** Kolom field yang dapat ditampilkan dan diedit di plan, termasuk custom field yang ditambahkan lewat plan settings.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Field bawaan tersedia: Assignee, Status, Team, Sprint, Releases (Fix versions), Target start, Target end, Start date, Due date, Estimate (story points/time), Progress, Dependencies, Parent, Labels, Components, Priority, Work type, Project, Reporter.
  2. Kolom dapat di-show/hide dan di-reorder.
  3. Plan settings > Custom fields memungkinkan menambah custom field ke plan.
  4. ⚠ Tipe custom field yang didukung: single select, multi select, checkbox, radio, date picker, number, single line text, user picker, dan labels; tipe lain tidak tampil di pemilih.
  5. Kolom Progress dihitung berdasarkan jumlah item atau estimate (sesuai pengaturan).
  6. Kolom status breakdown menampilkan distribusi status child untuk item parent.
  7. Field yang tidak ada di screen/context proyek suatu item bersifat read-only pada baris tersebut dengan ikon peringatan dan tooltip penjelasan.
  8. Mengedit field menghasilkan perubahan draft sampai disimpan.
- **Team-managed vs company-managed:** Team-managed: field yang tidak ada di work type proyek read-only (butir 7).
- **Catatan & ketergantungan:** Field didefinisikan di modul 05.

### PLAN-041 Plan filters

- **Deskripsi:** Filter tampilan plan untuk mempersempit work item yang tampil tanpa mengubah sumber plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian teks berdasarkan summary/key.
  2. Hierarchy range memilih level atas dan bawah yang ditampilkan (dari level X sampai level Y).
  3. Filter Projects, Work types, Teams, Assignees, dan Reporters.
  4. Filter Releases, Sprints, Components, dan Labels.
  5. Filter Status dan Status category.
  6. Filter Dependencies menampilkan hanya item yang memiliki dependency.
  7. Filter Date range menampilkan item yang rentang tanggalnya overlap dengan rentang terpilih.
  8. Toggle "Show only items with warnings".
  9. Custom field yang ditambahkan ke plan dapat dipakai sebagai filter.
  10. Toggle "Show full hierarchy" menampilkan ancestor yang tidak lolos filter (dengan gaya pudar) agar konteks terlihat.
  11. Indikator jumlah filter aktif dan tombol "Clear filters".
  12. Filter disimpan sebagai bagian saved view (PLAN-043).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Filter berbeda dari exclusion rules (PLAN-029): filter hanya memengaruhi tampilan, item tetap dihitung di rollup.

### PLAN-042 Plan view settings

- **Deskripsi:** Opsi tampilan plan: grouping, coloring, sorting, dan penampilan elemen.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Group by: none, Assignee, Component, Label, Project, Release, Sprint, Team, atau custom field.
  2. Color by: none, Status, Team, Project, Label, Component, Work type, Priority, atau custom field; legenda warna dapat diatur per nilai.
  3. Sort by: Rank (default), tanggal, estimate, atau field lain dengan arah asc/desc.
  4. Sort selain Rank menonaktifkan drag ranking dengan tooltip penjelasan.
  5. Dependencies style: lines atau badges.
  6. Toggle rolled-up values, warnings, dan capacity.
  7. Timeline range dapat dipilih dari preset (3, 6, 12 bulan, dll.) atau custom.
  8. Tombol collapse/expand semua level.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan di saved view (PLAN-043).

### PLAN-043 Plan saved views & sharing

- **Deskripsi:** Menyimpan kombinasi filter, field, dan view settings sebagai view bernama yang dapat dibagikan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan (membuat/mengubah view); Viewer plan (memakai view).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap plan memiliki view bawaan (default view) yang tidak dapat dihapus.
  2. Editor dapat Create view baru, Save, Save as, Rename, dan Delete view (nama unik per plan).
  3. "Set as default view" menjadikan view terbuka pertama bagi semua pengguna plan.
  4. Indikator perubahan view yang belum disimpan tampil di switcher, dengan aksi Save atau Discard.
  5. View terlihat oleh semua pengguna plan sesuai permission plan.
  6. View dapat dibagikan via link yang membuka plan pada view tersebut.
  7. View menyimpan filters, kolom field, group by, color by, sort, timeline range, dan mode view (Timeline/List/Calendar).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `plan_views`. Berbeda dari view bersama Timeline proyek (PLAN-010) dan List (PLAN-017).

## Plans: pengeditan, draft & penjadwalan

### PLAN-044 Plan inline create, bulk change & ranking

- **Deskripsi:** Membuat, mengubah massal, dan mengurutkan work item di dalam plan; semua perubahan berstatus draft.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan; penyimpanan tunduk pada permission proyek terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "+ Create" di header atau pada baris membuat item baru atau child; pengguna memilih proyek (dari sumber plan) dan work type yang valid untuk level tersebut.
  2. Item baru berstatus draft (ikon "plan-only") dan belum memiliki key sampai disimpan ke work item.
  3. Multi-select dilakukan via checkbox dan shift-click untuk rentang.
  4. Bulk change tersedia untuk assignee, team, sprint, release, labels, components, tanggal, estimate, rank, dan parent.
  5. Bulk "Remove from plan" mengeluarkan item terpilih dari plan (PLAN-029).
  6. Re-parent dilakukan via drag di daftar atau dengan mengubah kolom Parent; parent harus berasal dari level tepat di atasnya.
  7. Rank dapat diubah via drag, serta aksi "Move to top" / "Move to bottom".
  8. Perubahan rank juga berstatus draft sampai disimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validasi create final terjadi saat save (PLAN-045). Rank ISS-022.

### PLAN-045 Review changes & save to Jira

- **Deskripsi:** Mekanisme draft di Plans: perubahan disimpan di plan dan baru diterapkan ke work item produk ini lewat dialog Review changes.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan; setiap perubahan dievaluasi terhadap permission pengguna di proyek work item.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Review changes" di header plan menampilkan jumlah perubahan draft pada scenario aktif; nonaktif bila tidak ada perubahan.
  2. Dialog menampilkan daftar perubahan per item dengan kolom field, nilai lama, dan nilai baru.
  3. Pengguna dapat "Select all" atau memilih sebagian perubahan per item atau per field.
  4. "Save selected changes" menerapkan perubahan terpilih ke work item secara berurutan dengan indikator progres.
  5. "Discard selected changes" membuang perubahan terpilih setelah konfirmasi.
  6. Hasil save ditampilkan per item: sukses atau gagal dengan pesan galat (permission, field wajib, transisi workflow tidak valid); perubahan yang gagal tetap tersimpan sebagai draft.
  7. ⚠ Bila item telah berubah di work item setelah diubah di plan, dialog menandai konflik dengan nilai terkini; pengguna memilih menimpa dengan nilai plan atau mempertahankan nilai terkini.
  8. Perubahan pembuatan item, release, dan team (plan-only) juga tercantum di dialog dan disimpan sebelum perubahan yang mereferensikannya.
  9. Draft tersimpan di server per plan per scenario lintas sesi sampai disimpan atau dibuang, dan terlihat oleh semua editor plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data model: change set draft `plan_changes` (plan_id, scenario_id, entity_type, entity_ref, field, old_value, new_value, base_version, author_id, updated_at). Penyimpanan memakai endpoint update work item biasa sehingga history, notifikasi, otomasi, dan webhook berjalan normal.

### PLAN-046 Plan scenarios

- **Deskripsi:** Skenario what-if di plan untuk membandingkan alternatif jadwal dan alokasi tanpa mengubah work item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Scenarios diaktifkan di Plan settings > Scenarios; saat nonaktif plan hanya memiliki satu scenario implisit.
  2. Editor dapat membuat scenario baru, kosong (tanpa draft) atau salinan scenario lain beserta draft-nya.
  3. Scenario dapat di-rename dan di-delete (scenario terakhir tidak dapat dihapus).
  4. Dropdown di header plan berpindah antar scenario; setiap scenario memiliki warna penanda.
  5. Perubahan di satu scenario tidak memengaruhi scenario lain.
  6. Save ke work item hanya dilakukan dari scenario aktif; scenario lain otomatis menampilkan nilai terbaru untuk field yang tidak memiliki draft di scenario tersebut.
  7. ⚠ Maksimal 50 scenario per plan; membuat scenario ke-51 ditolak dengan pesan "Batas jumlah scenario tercapai".
  8. Scenario menyimpan perubahan field, tanggal, team, sprint, release, dan item plan-only.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `plan_scenarios`; change set di PLAN-045 dikunci per scenario.

### PLAN-047 Rollups

- **Deskripsi:** Agregasi nilai child ke parent di plan untuk tanggal, estimate, dan progres.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Rolled-up dates: parent tanpa tanggal sendiri menampilkan rentang dari tanggal child terawal sampai terakhir.
  2. Rolled-up estimate menampilkan total story points atau waktu child (termasuk keturunan tak langsung), di samping estimate parent sendiri.
  3. Progress rollup dihitung berdasarkan jumlah item Done atau estimate Done terhadap total.
  4. Status breakdown per parent menampilkan distribusi status category seluruh keturunan.
  5. Nilai rollup ditampilkan berbeda (italic/abu-abu) dari nilai langsung, dengan tooltip "Dihitung dari child".
  6. Toggle "Rolled-up values" di view settings menampilkan atau menyembunyikan nilai rollup.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Item yang di-exclude tidak ikut (PLAN-029); rollup memakai nilai draft scenario aktif.

### PLAN-048 Plan dependencies configuration

- **Deskripsi:** Pengaturan dan pengelolaan dependency di plan berbasis issue link.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Plan settings > Dependencies memilih satu atau lebih issue link type yang dianggap dependency (default "Blocks") dan arah blocking-nya.
  2. Dependency dapat ditampilkan sebagai lines atau badges (jumlah incoming/outgoing).
  3. Dependency dibuat dengan drag antar bar atau dari badge dengan memilih item tujuan; hasilnya berstatus draft.
  4. Dependency dapat dihapus (draft) dari popover atau badge.
  5. Kolom Dependencies dan filter Dependencies tersedia (PLAN-040, PLAN-041).
  6. Warning dependency conflict ditampilkan bila urutan jadwal melanggar dependency.
  7. Dependency ke item di luar plan ditandai "external" dan tidak digambar sebagai garis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Link types CFG-056, issue links ISS-029.

### PLAN-049 Sprint assignment in plans

- **Deskripsi:** Menetapkan work item ke sprint dari board team di plan dan melihat dampaknya pada jadwal dan capacity.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan; penyimpanan memerlukan Schedule Issues/Edit Issues di proyek item.
- **Aturan perilaku (acceptance criteria):**
  1. Kolom Sprint menampilkan sprint active dan future dari board yang terhubung ke team item.
  2. Assign ke sprint berstatus draft hingga disimpan.
  3. Tanggal item di-infer dari rentang sprint bila tanggal eksplisit kosong dan inferred dates aktif.
  4. Sprint dari board di luar plan ditandai "external sprint".
  5. Warning ditampilkan bila item berada di sprint milik team lain.
  6. Warning ditampilkan bila item berada di sprint yang sudah selesai dan item belum Done.
  7. Warning ditampilkan bila rentang sprint tidak selaras dengan release atau tanggal item.
  8. Group by Sprint tersedia di view settings.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sprint field dan siklus sprint di modul 11.

### PLAN-050 Plan capacity planning & teams

- **Deskripsi:** Halaman Teams di plan untuk menambahkan team, mengatur capacity per iterasi, dan memvisualisasikan pemakaian capacity.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Teams di plan menampilkan daftar team plan dan tombol tambah team.
  2. Team dapat berupa team shared (entitas Teams platform, PLAN-057) atau plan-only team yang hanya ada di plan.
  3. Tipe scheduling team dipilih Scrum (berbasis sprint) atau Kanban (berbasis iterasi).
  4. Team Scrum dihubungkan ke issue source board untuk mengambil sprint-nya.
  5. Capacity diatur per sprint (story points atau jam sesuai estimation unit) untuk team Scrum dan per iterasi untuk team Kanban.
  6. Iteration length (dalam minggu) diatur untuk team Kanban.
  7. Velocity default diisi dari rata-rata sprint sebelumnya bila tersedia.
  8. Anggota team dan weekly hours per anggota dapat diatur untuk menghitung capacity berbasis waktu.
  9. Capacity bar per sprint di timeline menunjukkan under / at / over capacity dengan warna berbeda.
  10. Warning over capacity ditampilkan saat estimate di sprint melebihi capacity.
  11. Field Team pada work item dipakai untuk assignment team; mengubah team di plan mengisi field Team (draft).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `plan_teams`; Team field di PLAN-058 dan CFG-019.

### PLAN-051 Cross-project releases & Releases view

- **Deskripsi:** Mengelola release (version) di plan, termasuk cross-project release yang mengelompokkan version dari beberapa proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan; Admin Proyek atau Manage Versions (PERM-026) untuk menyimpan release proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Releases di plan menampilkan daftar release per proyek dan cross-project release.
  2. Create release di plan: nama, proyek, start date, release date; berstatus draft sampai disimpan.
  3. ⚠ Start date dan release date dapat fixed atau relative to previous release (dihitung dari release sebelumnya di proyek yang sama).
  4. Cross-project release dibuat dengan nama dan berisi project releases dari beberapa proyek.
  5. Release dapat di-edit, di-rename, dan di-archive dari plan.
  6. Status release ditampilkan: released, unreleased, archived.
  7. Release markers tampil di timeline plan.
  8. Item dapat di-assign ke release (Fix versions) dari kolom Releases (draft).
  9. Warning ditampilkan bila item dijadwalkan selesai setelah release date.
  10. Filter by release tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed memerlukan fitur Releases aktif.
- **Catatan & ketergantungan:** Model release di REL-024 dan REL-001 (modul 15).

### PLAN-052 Auto-scheduler

- **Deskripsi:** Menjadwalkan item otomatis ke sprint/tanggal berdasarkan rank, estimate, capacity, dependency, dan release.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / premium
- **Aktor & izin:** Editor plan.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Auto-schedule" tersedia di header plan bagi editor.
  2. ⚠ Scope dapat dipilih: semua item yang terlihat atau hanya item terpilih.
  3. Algoritma mempertimbangkan rank, estimate, team capacity, dependency, dan release date.
  4. Pengguna memilih field yang boleh ditimpa: sprint, team, release, tanggal.
  5. Default estimate dipakai untuk item tanpa estimate (PLAN-033).
  6. Preview perubahan (highlight) ditampilkan sebelum pengguna menerima atau membatalkan.
  7. Hasil yang diterima menjadi draft yang harus di-review/save (PLAN-045).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status dan opsi auto-scheduler terbaru di produk referensi belum terkonfirmasi; algoritma produk ini deterministik (hasil sama untuk input sama).

### PLAN-053 Plan warnings

- **Deskripsi:** Katalog peringatan yang mendeteksi masalah jadwal dan alokasi di plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan.
- **Aturan perilaku (acceptance criteria):**
  1. Warning: child dijadwalkan di luar tanggal parent.
  2. Warning: dependency conflict.
  3. Warning: item di sprint team lain.
  4. Warning: item di sprint yang sudah selesai.
  5. Warning: item dijadwalkan setelah release date.
  6. Warning: parent Done tetapi child belum selesai.
  7. Warning: over capacity sprint/iterasi.
  8. Warning: tanggal release tidak konsisten (start setelah release date, atau release cross-project tidak selaras).
  9. ⚠ Warning: sub-task berada di sprint berbeda dari parent.
  10. Ikon warning tampil per baris, dan filter "Show only items with warnings" tersedia.
  11. Panel warning menampilkan daftar warning item terpilih dengan penjelasan dan saran perbaikan.
  12. ⚠ Setiap jenis warning dapat dinyalakan/dimatikan per plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Warning Timeline proyek (PLAN-009) adalah subset katalog ini dan memakai mesin deteksi yang sama.

## Plans: distribusi, batas & API

### PLAN-054 Plan export, embed & share

- **Deskripsi:** Membagikan plan ke stakeholder lewat link, ekspor, atau embed di wiki eksternal.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Editor plan, Viewer plan; penerima harus memiliki akses plan.
- **Aturan perilaku (acceptance criteria):**
  1. Menu Share > "Copy link" menyalin tautan ke plan pada view dan scenario aktif.
  2. Export CSV dari view menghasilkan berkas dengan kolom sesuai view (termasuk nilai draft, dengan kolom penanda "Unsaved").
  3. Timeline plan dapat diekspor sebagai image (PNG) sesuai rentang tampilan.
  4. Plan dapat di-embed di wiki eksternal melalui Share > Smart Link/oEmbed (modul 25).
  5. Penerima tanpa akses edit melihat plan atau embed secara read-only.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ekspor Plans/Timeline/List juga dicatat di modul 25.

### PLAN-055 Plan limits & performance

- **Deskripsi:** Batasan ukuran dan performa plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Sistem; Editor plan menerima peringatan.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Batas jumlah item per plan default 5.000 dan dapat dinaikkan untuk tier enterprise melalui konfigurasi sistem.
  2. Batas issue sources 100 per plan (sama dengan PLAN-028 butir 4).
  3. Batas scenario per plan ditegakkan (PLAN-046).
  4. Saat batas item terlampaui, banner peringatan menampilkan jumlah item yang tidak dimuat dan saran menambah exclusion rules atau filter.
  5. Child dimuat lazy saat parent di-expand dan rentang hierarki default dibatasi agar render awal tetap dalam anggaran performa.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Angka batas produk referensi belum terkonfirmasi; nilai di sini adalah default produk ini. Anggaran performa di modul 27.

### PLAN-056 Plans REST API

- **Deskripsi:** Endpoint REST publik untuk mengelola plans dan team di plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Integrator/Developer API; akses mengikuti permission plan.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /api/v1/plans` (paginasi) dan `GET /api/v1/plans/{id}` mengembalikan daftar dan detail plan.
  2. `POST /api/v1/plans` membuat plan beserta sources, exclusion rules, scheduling, dan permissions; validasi sama dengan wizard (PLAN-028).
  3. `PUT /api/v1/plans/{id}` memperbarui plan dengan body JSON Patch (RFC 6902); operasi tidak valid ditolak HTTP 400 dengan detail path.
  4. Endpoint archive, trash, dan duplicate plan tersedia (`POST /api/v1/plans/{id}/archive`, `/trash`, `/duplicate`).
  5. Endpoint team di plan: tambah/hapus team platform, CRUD plan-only team.
  6. Akses mengikuti permission plan; pengguna tanpa akses mendapat HTTP 404 (tidak membocorkan keberadaan plan Private).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kontrak endpoint di API-059 ([23](23-api-publik-webhook.md)); modul 23 adalah pemilik kontrak endpoint (path, skema request/response, kode galat, paginasi), sedangkan butir di atas hanya merangkum perilaku domain Plans (PLAN-028 dan permission plan) yang harus dihormati API. Modul 23 perlu menambahkan rujukan balik ke PLAN-056. Path produk referensi `/rest/api/3/plans/...` menjadi target kompatibilitas. Anotasi Swagger wajib (`make swag`).

## Teams

### PLAN-057 Atlassian Teams

- **Deskripsi:** Entitas team tingkat platform di produk ini (padanan Atlassian Teams) yang dipakai lintas proyek dan modul, termasuk managed team yang keanggotaannya disinkronkan dari group atau identity provider.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna (membuat team biasa), anggota team, Admin Situs/org admin (managed team).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat membuat team dengan nama (wajib, maks 255 karakter), deskripsi, dan undangan anggota.
  2. Anggota dapat ditambah dan dihapus; anggota dapat leave team.
  3. Team dapat dihapus oleh anggota atau admin; menghapus team mengosongkan Team field di work item yang mereferensikannya setelah konfirmasi.
  4. Direktori team menyediakan daftar dan pencarian team berdasarkan nama dan anggota.
  5. ⚠ Pengaturan membership menentukan siapa yang boleh bergabung atau menambah anggota: siapa saja, hanya anggota, atau hanya dengan undangan.
  6. ⚠ Team types: team biasa (open/member-invite) dan managed team (keanggotaan dikendalikan admin).
  7. Managed team dibuat oleh org admin dari group yang ada atau disinkronkan dari identity provider via SCIM (USR-035).
  8. Keanggotaan managed team read-only di UI Teams; perubahan hanya lewat sinkronisasi group/IdP.
  9. ⚠ Badge "Managed" tampil di team profile dan team picker untuk managed team.
  10. ⚠ Bila group sumber dihapus atau sinkronisasi diputus, managed team dikonversi menjadi team biasa dengan anggota terakhir dipertahankan dan admin diberi notifikasi.
  11. Hanya org admin yang boleh mengubah tipe team atau memutus sinkronisasi.
  12. Perubahan anggota managed team otomatis tercermin pada Team field dan anggota team di capacity Plans.
  13. ⚠ Team hierarchy (parent team / sub-team) tidak dibangun pada rilis awal; model data menyediakan kolom `parent_team_id` opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team bersifat lintas proyek.
- **Catatan & ketergantungan:** Administrasi team di [22](22-administrasi-sistem.md), group di USR-015, direktori People & Teams di USR-027. Tabel `teams`, `team_members`. Nama "Atlassian Teams" hanya nama inventaris; produk ini menyebutnya "Teams".

### PLAN-058 Team field

- **Deskripsi:** Field Team di work item yang mereferensikan entitas Teams.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim dengan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Team dapat dipilih di issue view dan dialog Create melalui team picker dengan pencarian.
  2. Field Team dapat dipakai di JQL (`team = <id>`), filter board, dan plan.
  3. Satu work item hanya dapat memiliki satu team; memilih team lain menggantikan nilai sebelumnya.
  4. Plans memakai field Team untuk capacity dan group by Team.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi field di CFG-019 (modul 05); JQL di modul 12.

### PLAN-059 Team profile page

- **Deskripsi:** Halaman profil team yang menampilkan informasi dan pekerjaan team.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna berlisensi; anggota dapat mengedit profil.
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan nama, deskripsi, dan gambar header team (dapat diedit anggota).
  2. Daftar anggota dengan avatar dan tautan profil.
  3. Daftar links ke resource terkait (URL bebas dengan judul).
  4. Bagian pekerjaan team menampilkan work item dengan field Team = team ini, terbaru diperbarui lebih dulu, sesuai permission pengguna.
  5. ⚠ Bagian Goals terkait menampilkan goal yang mencantumkan team ini; konten wiki eksternal tampil bila integrasi modul 25 aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Direktori di USR-027 (modul 02).

### PLAN-060 Teams REST API

- **Deskripsi:** API publik untuk mengelola Teams dan anggotanya.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, org admin.
- **Aturan perilaku (acceptance criteria):**
  1. CRUD team via `GET/POST /api/v1/teams`, `GET/PATCH/DELETE /api/v1/teams/{id}`; managed team tidak dapat diubah anggotanya via API (HTTP 409).
  2. Kelola anggota team via `GET/POST/DELETE /api/v1/teams/{id}/members`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kontrak endpoint di API-060 ([23](23-api-publik-webhook.md)); modul 23 adalah pemilik kontrak endpoint (path, skema request/response, kode galat, paginasi), sedangkan butir di atas hanya merangkum perilaku domain Teams (PLAN-057) yang harus dihormati API. Prioritas diselaraskan dengan API-060 (P2). Modul 23 perlu menambahkan rujukan balik ke PLAN-060. Cakupan endpoint produk referensi belum terkonfirmasi.

## Goals

### PLAN-061 Atlassian Goals

- **Deskripsi:** Goals tingkat platform untuk melacak tujuan dan progresnya, yang dapat dihubungkan ke work item. Produk ini membangun entitas Goals native yang ringan sebagai padanan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna berlisensi (membuat), goal owner (mengelola).
- **Aturan perilaku (acceptance criteria):**
  1. Goal dibuat dengan nama (wajib), owner (wajib, default pembuat), dan target date (opsional).
  2. ⚠ Status goal: Pending, On track, At risk, Off track, Done, Paused, Cancelled; default Pending.
  3. Owner atau contributor dapat menulis status update berkala berisi status dan catatan; riwayat update ditampilkan terbaru lebih dulu.
  4. Goal dapat memiliki sub-goals dan satu parent goal; siklus ditolak.
  5. Pengguna dapat menjadi follower; owner dapat menambah contributors.
  6. Goal dapat diberi tags untuk pengelompokan dan pencarian.
  7. ⚠ Goal dapat di-archive (disembunyikan dari daftar aktif, tetap dapat dilihat lewat filter "Archived") dan di-unarchive.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Administrasi Goals di [22](22-administrasi-sistem.md); tabel `goals`, `goal_updates`, `goal_followers`. Nama "Atlassian Goals" hanya nama inventaris.

### PLAN-062 Link work items to Goals

- **Deskripsi:** Menghubungkan work item ke goal lewat field Goals dan menampilkan progres serta update goal di produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim dengan Edit Issues; Stakeholder/Viewer melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Field Goals di issue view memungkinkan menautkan satu atau lebih goal ke work item.
  2. Halaman goal menampilkan daftar linked work (work item yang menautkan goal) sesuai permission pengguna.
  3. ⚠ Progres goal dapat dihitung dari persentase linked work yang Done, sebagai informasi di samping status manual.
  4. ⚠ Panel goal di Summary proyek dan plan menampilkan goal yang ditautkan item di dalamnya.
  5. ⚠ JQL mendukung `goals = <id>` untuk memfilter work item berdasarkan goal.
  6. Follower goal menerima notifikasi saat status update baru diterbitkan (modul 19).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `goal_links` (goal_id, issue_id). Integrasi dengan layanan goals eksternal di modul 25.

## Legacy & integrasi portofolio (rujukan)

### PLAN-063 Overviews (cross-project views)

- **Alias inventaris:** Overviews (Jira Work Management, deprecated)

- **Deskripsi:** Fitur lama business project yang menggabungkan beberapa proyek ke satu Overview dengan view List, Board, Timeline, Calendar, dan Summary gabungan. Sudah dihentikan di produk referensi dan digantikan Plans; dicatat hanya sebagai rujukan dan inspirasi multi-project view, bukan target implementasi.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna Tim Bisnis, Project lead (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ (Rujukan) Membuat Overview dengan memilih beberapa proyek sumber (business; software bila didukung).
  2. (Rujukan) View List lintas proyek dengan kolom, filter, dan inline edit.
  3. (Rujukan) View Timeline lintas proyek.
  4. (Rujukan) View Calendar lintas proyek.
  5. (Rujukan) View Summary gabungan (status overview, priority breakdown, types of work).
  6. (Rujukan) Filter Overview per proyek, assignee, status, work type, label.
  7. (Rujukan) Edit nama dan daftar proyek Overview; delete Overview.
  8. (Rujukan) Overview muncul di sidebar atau menu Projects.
  9. (Rujukan) Hasil mengikuti permission: work item dari proyek yang tidak dapat di-browse tidak tampil.
  10. ⚠ (Rujukan) Batas jumlah proyek per Overview.
  11. (Rujukan) Membuat, mengedit, dan menghapus overview berisi sekumpulan business project.
  12. (Rujukan) View lintas proyek: List, Board, Calendar, Timeline, Summary.
  13. ⚠ (Rujukan) Akses dibatasi oleh lisensi/tier business project.
  14. (Rujukan) Pemberitahuan deprecation dan migrasi ke Plans (tier premium) sebelum akses berakhir.
  15. Untuk produk ini, Overviews hanya dicatat sebagai inspirasi multi-project view; kebutuhan lintas proyek dipenuhi Plans (PLAN-027), saved filter + List (modul 13), dan dashboard (modul 17).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tidak dibangun; masuk Di luar cakupan.

### PLAN-064 Jira Align integration

- **Deskripsi:** Integrasi dengan produk perencanaan enterprise terpisah (Jira Align). Dicatat sebagai rujukan; produk ini tidak membangun konektor tersebut.
- **Prioritas / Fase / Tier:** P3 / Fase — / enterprise
- **Aktor & izin:** Admin Situs (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan) Connector sinkronisasi dua arah antara produk ini dan produk portofolio enterprise.
  2. (Rujukan) Mapping proyek dan team antara kedua produk.
  3. (Rujukan) Sinkronisasi work item dan sprint.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk terpisah, di luar cakupan; konektor setara bisa dibangun pihak ketiga memakai REST API dan webhook (modul 23, PLAN-056). Lihat juga modul 25 (Jira Align connector).

### PLAN-065 Atlassian Focus integration

- **Deskripsi:** Titik integrasi dengan produk strategic portfolio planning terpisah (Focus) yang menghubungkan focus area dan pendanaan ke Goals dan pekerjaan. Dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / enterprise
- **Aktor & izin:** Portfolio manager, eksekutif, Admin Situs (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan) Focus areas hierarkis ditautkan ke Goals dan work item/plan.
  2. (Rujukan) Rollup progres dari pekerjaan ke focus area.
  3. (Rujukan) Entry point di app switcher/sidebar hanya bila produk dilisensikan.
  4. (Rujukan) Permission dan lisensi terpisah dari produk inti.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk terpisah, di luar cakupan. Inventaris menandai nama resmi dan titik integrasi persis produk ini belum terkonfirmasi (penanda ⚠ di tingkat fitur), sehingga seluruh butir di atas hanya rujukan. Kebutuhan dasar keterkaitan strategi-pekerjaan dipenuhi Goals dengan sub-goals (PLAN-061, PLAN-062).

### PLAN-066 Atlassian Talent integration

- **Deskripsi:** Titik integrasi dengan produk workforce planning terpisah (Talent) yang menghubungkan kebutuhan orang/skill dan alokasi kapasitas ke Teams dan pekerjaan. Dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / enterprise
- **Aktor & izin:** Org admin, portfolio/program manager, team lead (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan) Menampilkan alokasi orang (headcount/skill) per team yang dipakai di Team field.
  2. (Rujukan) Sinkronisasi team produk workforce dengan Teams yang dipakai di Plans.
  3. (Rujukan) Keterhubungan rencana tenaga kerja dengan produk portofolio lain.
  4. (Rujukan) Entry point dari sidebar/app switcher hanya bila produk dilisensikan.
  5. (Rujukan) Gating lisensi lewat paket produk strategi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk terpisah, di luar cakupan. Inventaris menandai nama resmi dan titik integrasi persis produk ini belum terkonfirmasi (penanda ⚠ di tingkat fitur), sehingga seluruh butir di atas hanya rujukan. Capacity per team dan weekly hours anggota (PLAN-050) adalah padanan minimal di produk ini.

## Capacity individual per space

### PLAN-067 ⚠ Individual capacity planning (Capacity tab)

- **Deskripsi:** Tab Capacity di navigasi space untuk merencanakan alokasi waktu setiap orang per minggu ke work item, epic, project, dan pekerjaan non-proyek, dengan capacity bar realtime yang menandai over-allocation. Berbeda dari capacity team per sprint/iterasi di Plans (PLAN-050) dan dari User Workload Report berbasis Remaining Estimate (TIME-032); fitur ini dicatat dari rilis produk referensi Juli 2026 dan belum ada di inventaris.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Scrum Master/Team Lead dan Admin Proyek/Project Lead (menambah tab, mengedit alokasi); Developer/Anggota Tim (melihat alokasi dirinya dan, ⚠ bila diizinkan, mengedit alokasinya sendiri); Stakeholder/Viewer dengan Browse Projects (read-only). Menambah tab memerlukan Administer Projects; mengedit alokasi memerlukan ⚠ izin khusus "Manage capacity" atau Administer Projects. Fitur hanya tersedia pada tier premium dan enterprise (VISI-003).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tab Capacity ditambahkan ke navigasi space lewat tombol "+" (PLAN-001) dan tersedia untuk software maupun business space, team-managed maupun company-managed; pada tier di bawah premium tombol menampilkan lozenge upgrade dan tidak membuat tab.
  2. ⚠ People view menampilkan baris per orang dengan sub-baris per work item yang dialokasikan, dan kolom per minggu berisi nilai alokasi; total per orang per minggu ditampilkan di baris orang.
  3. ⚠ Work view menampilkan hierarki epic, project, dan work item beserta assignee dan alokasi per minggu, dengan indikator kecukupan resource (mis. kurang, cukup, berlebih) dibandingkan kebutuhan estimasi work item.
  4. ⚠ Alokasi dapat dibuat untuk pekerjaan non-proyek dengan kategori Cuti, On-call, Training, dan Maintenance; alokasi Cuti otomatis mengurangi kapasitas tersedia orang tersebut pada minggu terkait (bukan dihitung sebagai beban kerja).
  5. ⚠ Unit alokasi dapat dipilih Persen, Jam, atau Hari; default Persen dengan basis 40 jam per minggu (100% = 40 jam, 1 hari = 8 jam), dan mengganti unit mengonversi nilai yang ada tanpa kehilangan presisi di bawah 0,5 jam.
  6. ⚠ Capacity bar per orang per minggu diperbarui realtime saat nilai diubah (tanpa reload, modul 27) dan berwarna merah bila total alokasi melebihi 100%; sistem tidak menerapkan hard cap, sehingga nilai di atas 100% tetap dapat disimpan dengan peringatan visual.
  7. ⚠ Alokasi bercakupan seluruh site: alokasi seseorang di satu space terlihat (sebagai total dan, bila pengguna dapat mem-browse space sumber, sebagai rincian) di tab Capacity space lain, sehingga over-allocation lintas space terdeteksi; rincian work item dari space yang tidak dapat di-browse ditampilkan sebagai "Pekerjaan lain" tanpa key/summary.
  8. ⚠ Grid alokasi dapat diedit gaya spreadsheet: klik sel untuk mengetik nilai, copy-paste rentang sel (termasuk dari spreadsheet eksternal), drag-to-fill untuk menyalin nilai ke minggu berikutnya, dan Undo untuk perubahan terakhir; nilai negatif atau non-numerik ditolak dengan pesan galat di sel.
  9. ⚠ Impor CSV: pengguna mengunggah file berisi orang, work item/kategori non-proyek, minggu, dan nilai; sistem menampilkan langkah pemetaan kolom ke anggota dan work item, menandai baris yang tidak cocok untuk direview, dan baru menerapkan alokasi setelah pengguna mengonfirmasi hasil review.
- **Team-managed vs company-managed:** Tidak ada perbedaan; tab tersedia di kedua model dan pada software maupun business space.
- **Catatan & ketergantungan:** Bedakan dari PLAN-050 (capacity team per sprint/iterasi di Plans, berbasis story points/jam team) dan TIME-032 ([14](14-pelacakan-waktu-estimasi.md), laporan Remaining Estimate per user tanpa alokasi). ⚠ Keputusan terbuka apakah weekly hours anggota (PLAN-050 butir 8) dan data cuti dipakai bersama. Data baru: tabel `capacity_allocations` (user, work item atau kategori non-proyek, minggu, nilai, unit) bercakupan site, dan izin baru di modul 03. Impor CSV memakai pola pemetaan modul 25. Kalender tim/cuti lengkap dan resource leveling otomatis tetap di luar cakupan (ekosistem app).

## Pertanyaan terbuka

- Apakah view bersama (Timeline, List, Board business) disimpan per proyek saja, atau setiap pengguna boleh menyimpan beberapa view personal bernama di samping view bersama (PLAN-010, PLAN-017, PLAN-026)?
- Apakah Timeline proyek software menampilkan work item standar tanpa Epic sebagai baris level atas secara default, atau disembunyikan seperti perilaku historis produk referensi (PLAN-002)?
- Apakah pengaturan urutan dan penyembunyian tab view berlaku per pengguna saja atau Admin Proyek juga dapat mengunci urutan default untuk semua anggota (PLAN-001)?
- Apakah Plans dan Teams/Goals dibatasi untuk tier premium/platform berbayar, atau semua instalasi self-hosted produk ini mendapatkannya tanpa gating (PLAN-027, PLAN-057, PLAN-061)?
- Bagaimana kebijakan penyelesaian konflik saat save draft plan: plan menang, nilai terkini menang, atau selalu meminta pilihan pengguna (PLAN-045)?
- Apakah batas default (5.000 item Timeline/plan, 100 sumber, 50 scenario, 10 view, 20 rule format) cukup untuk target pelanggan, dan apakah dapat dikonfigurasi per tier (PLAN-014, PLAN-055)?
- Apakah team hierarchy (parent/sub-team) diperlukan sejak rilis awal Teams, mengingat dampaknya pada capacity dan group by Team (PLAN-057)?
- Apakah progres goal otomatis dari linked work ditampilkan sebagai angka resmi atau hanya informasi pendamping status manual (PLAN-062)?
- Apakah tab Capacity individual (PLAN-067) dan capacity team di Plans (PLAN-050) berbagi satu sumber jam kerja dan cuti per orang, atau dikelola terpisah; dan apakah alokasi Capacity perlu digating tier premium seperti produk referensi?
- Apakah auto-scheduler (P3) tetap dibangun, atau cukup scheduling manual + warning mengingat kompleksitas algoritma (PLAN-052)?

## Di luar cakupan & ekosistem

- Overviews legacy business project (PLAN-063): fitur yang sudah dihentikan di produk referensi; kebutuhan lintas proyek dipenuhi Plans, saved filter + List, dan dashboard.
- Produk portofolio dan strategi terpisah beserta konektornya: Jira Align (PLAN-064), Focus (PLAN-065), Talent (PLAN-066); hanya dimungkinkan melalui REST API dan webhook publik produk ini.
- Layanan platform lintas produk referensi (Atlassian Home, Teamwork Graph, Confluence Team Calendars, Smart Link milik vendor); produk ini membangun Teams dan Goals native yang ringan dan embed generik via oEmbed/Smart Link (modul 25).
- Subscribe/iCal feed Calendar ke Google/Outlook/Apple Calendar; bukan fitur native, tersedia di ekosistem app.
- Global permission bergaya "Portfolio user/viewer" dan konfigurasi Portfolio for Jira Server/Data Center; Plans cloud memakai permission per plan (PLAN-030).
- App Marketplace pihak ketiga untuk perencanaan lanjutan (Gantt berbasis critical path, resource leveling, baseline/snapshot jadwal, portofolio berbasis struktur bebas, timesheet capacity, kalender tim/cuti); tidak dibangun sebagai fitur native.
- Fitur AI vendor referensi (mis. pembuatan plan atau ringkasan roadmap berbasis Atlassian Intelligence/Rovo); bila ada, mengikuti fitur AI opsional produk ini di modul 27.
