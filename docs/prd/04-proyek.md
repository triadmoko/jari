# 04 — Proyek: Tipe, Template, Pengaturan & Siklus Hidup

Modul ini mendefinisikan **project** (di UI baru disebut *Space*) sebagai wadah utama work item di produk ini: identitas dan atribut project, tipe project (software, business, service, product discovery), dua model manajemen (team-managed dan company-managed), katalog template dan wizard pembuatan, aturan nama dan key, halaman Project settings beserta seluruh sub-halamannya, direktori dan navigasi project, serta siklus hidup project (aktif, arsip, trash, hapus permanen). Modul ini menjadi poros konfigurasi: hampir semua modul lain (work item, workflow, board, permission, notifikasi, otomasi, API) merujuk ke entitas project yang didefinisikan di sini. Detail mendalam tiap konfigurasi (permission scheme, workflow editor, custom field, dsb.) tetap dimiliki modul masing-masing; modul ini hanya mengatur **pemasangannya pada level project**.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`); belum ada entitas project, role, maupun scheme. Seluruh fitur di modul ini adalah pekerjaan baru dan akan menjadi feature module `internal/features/project` beserta entitas di `internal/shared/domain/entity`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah Project/Space, Work item, tier plan.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — product access, undangan, access request tingkat site.
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission scheme, project role, issue security, global permission.
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — work type, field, screen, field configuration, priority scheme, issue layout.
- [06 Work Item Inti](06-work-item-inti.md) — issue key, bulk move, archive per work item.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — workflow scheme dan editor workflow.
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — board per project, Scrum/Kanban.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — field/fungsi JQL project.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md), [16 Laporan](16-laporan-insight.md), [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md).
- [19 Notifikasi](19-notifikasi-email.md), [20 Otomasi](20-otomasi.md), [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md), [22 Administrasi Sistem](22-administrasi-sistem.md).
- [23 REST API & Webhook](23-api-publik-webhook.md), [24 Integrasi DevOps](24-integrasi-devops.md), [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md).
- [26 Kerja Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md), [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md), [28 Model Data Inti](28-model-data-inti.md).
- [29 Service Project](29-jsm-service-project-request-type-formulir.md) dan [36 Product Discovery](36-jpd-ideas-field-views.md) — tipe project service dan product discovery.

## Persona terkait

- **Admin Situs** — membuat project company-managed, mengelola kategori, role global, scheme, arsip, dan trash.
- **Admin Proyek/Project Lead** — mengubah detail project, mengelola anggota/akses, features, work type team-managed, versions, components, shortcut.
- **Scrum Master/Team Lead** — memilih template Scrum/Kanban, mengaktifkan Sprints/Backlog/Estimation.
- **Developer/Anggota Tim** — membuka project lewat directory/starred/recent, memakai navigasi dan view project.
- **Product Owner/PM** — memakai Summary, Timeline, Releases, dan membuat project dari template.
- **Stakeholder/Viewer** — melihat project Limited/Open atau sebagai role Viewer.
- **Pengguna Tim Bisnis** — memakai template business dan view List/Calendar/Forms.
- **Integrator/Developer API** — memakai REST API project, project properties, webhook event project.
- **Tamu/Kolaborator Eksternal** — mengirim feedback lewat Issue collector; meminta akses ke project Private.

## User story modul

1. Sebagai Admin Proyek/Project Lead, saya ingin membuat project dari template Scrum atau Kanban dalam satu wizard agar tim bisa langsung bekerja tanpa konfigurasi manual (PRJ-006, PRJ-010).
2. Sebagai Admin Situs, saya ingin memilih antara model team-managed dan company-managed saat membuat project agar tim otonom dan tim yang distandardisasi sama-sama terlayani (PRJ-003, PRJ-013).
3. Sebagai Developer/Anggota Tim, saya ingin key project yang pendek dan unik menjadi prefiks setiap work item agar saya bisa merujuk work item dengan mudah (PRJ-001, PRJ-012).
4. Sebagai Admin Proyek/Project Lead, saya ingin mengganti key project tanpa memutus tautan lama agar rebranding tim tidak merusak bookmark dan JQL (PRJ-022).
5. Sebagai Admin Proyek/Project Lead, saya ingin mengatur access level Open/Limited/Private dan role anggota agar hanya orang yang tepat yang dapat melihat dan mengubah work item (PRJ-038, PRJ-039).
6. Sebagai Admin Situs, saya ingin project baru berbagi scheme dengan project yang sudah ada agar konfigurasi lintas tim konsisten (PRJ-014).
7. Sebagai Developer/Anggota Tim, saya ingin menandai bintang project dan melihat daftar recent agar saya cepat berpindah ke project yang sering saya pakai (PRJ-025, PRJ-026).
8. Sebagai Admin Situs, saya ingin memindahkan project ke trash dengan masa pemulihan 60 hari agar penghapusan tidak sengaja dapat dibatalkan (PRJ-031, PRJ-033).
9. Sebagai Admin Situs, saya ingin mengarsipkan project yang tidak aktif agar pencarian dan board tetap bersih tanpa kehilangan data historis (PRJ-032).
10. Sebagai Scrum Master/Team Lead, saya ingin menyalakan atau mematikan fitur Sprints, Backlog, dan Estimation pada project team-managed agar gaya kerja project sesuai praktik tim (PRJ-049, PRJ-023).
11. Sebagai Admin Proyek/Project Lead, saya ingin menambah work type, field, dan status sendiri di project team-managed agar tidak perlu menunggu Admin Situs (PRJ-051, PRJ-052).
12. Sebagai Product Owner/PM, saya ingin halaman Summary project agar saya melihat kemajuan dan beban kerja tim dalam sekali pandang (PRJ-028).
13. Sebagai Pengguna Tim Bisnis, saya ingin template business seperti Marketing atau HR agar proses non-software saya langsung punya workflow yang relevan (PRJ-007).
14. Sebagai Integrator/Developer API, saya ingin membuat, mencari, dan mengubah project lewat REST API dengan validasi key/nama agar provisioning project bisa diotomasi (PRJ-064, PRJ-062).
15. Sebagai Tamu/Kolaborator Eksternal, saya ingin meminta akses ke project Private yang tautannya saya terima agar admin dapat menambahkan saya tanpa proses manual di luar sistem (PRJ-038, PRJ-004).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 26 |
| P1 | 24 |
| P2 | 13 |
| P3 | 7 |
| **Total** | **70** |

| Fase | Jumlah fitur | Isi utama |
|---|---|---|
| Fase 1 | 21 | Entitas project, tipe, dua model, template Scrum/Kanban, wizard, aturan nama/key, izin create, project lead, detail, directory, navigasi, settings hub, People/Access, role global, Features team-managed, REST API inti project |
| Fase 2 | 24 | Licensing gate, template library & blank, default assignee, avatar, kategori, ganti key, trash & lifecycle, starred/recent, Summary, halaman konfigurasi scheme & issue layout, Features company-managed, work type/workflow team-managed, notifikasi, versions/components, JQL project |
| Fase 3 | 15 | Template business, custom template (paid plan), share settings, shortcuts, ganti gaya/migrasi model, archive, extended project administration, issue security, priorities, email sender, otomasi, project properties, administrasi project global, Space Policies |
| Fase 4 | 8 | Onboarding, Pages, toolchain/development tools, apps, issue collectors, migrasi antar-instance, halaman Agents |
| Fase — | 2 | Fitur khas Data Center (rujukan saja) |

Strategi implementasi: Fase 1 harus menetapkan **model data project yang final** — ID numerik immutable, key + tabel alias key, `project_type_key`, `style` (team-managed/company-managed), dan `status` lifecycle — karena hampir semua tabel lain (work item, board, workflow, scheme) akan berelasi ke sini dan sulit diubah belakangan. Kedua model manajemen didukung sejak awal di level skema data, tetapi pada Fase 1 project company-managed cukup memakai scheme default sistem (tanpa editor), sedangkan team-managed memakai konfigurasi project-scoped yang di-seed dari template. Fase 2 membuka halaman konfigurasi scheme per project dan editor team-managed setelah modul [05](05-konfigurasi-work-item-skema.md) dan [09](09-workflow-status-transisi.md) siap. Endpoint REST inti project (search, CRUD, validasi key/nama) dan definisi role global ikut Fase 1 karena dipakai SPA dan halaman People (PRJ-064, PRJ-040). Fase 3 menambahkan kemampuan skala organisasi (archive, share settings, extended administration, endpoint API lanjutan, custom template paid plan, Space Policies). Fitur yang bergantung pada ekosistem eksternal (Pages, toolchain, apps) ditunda ke Fase 4.

## Fitur

## Model & tipe project

### PRJ-001 Project (Space) Entity & Terminology

- **Deskripsi:** Project adalah wadah utama work item dengan identitas unik (ID, key, name) dan konfigurasi sendiri. UI memakai label "Space" untuk project dan "Work item" untuk issue, sementara data, filter, JQL, dan API tetap memakai istilah `project`.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna berlisensi (melihat sesuai Browse Projects/access level); Admin Proyek mengubah atribut; Admin Situs mengelola semua project.
- **Aturan perilaku (acceptance criteria):**
  1. Entitas project menyimpan atribut inti: `id` numerik (immutable, dipakai internal dan API), `key`, `name`, `description`, `lead`, `default_assignee_type`, `avatar`, `url`, `category_id`, `project_type_key` (`software`/`business`/`service_desk`/`product_discovery`), `style` (`classic` = company-managed, `next-gen` = team-managed), flag `simplified`, serta `archived_at/archived_by` dan `deleted_at/deleted_by`; `id` tidak pernah berubah walau key atau nama diganti.
  2. Setiap work item milik tepat satu project; issue key berformat `PROJECTKEY-<n>` dengan `n` bertambah per project secara atomik (counter per project), dan nomor tidak pernah dipakai ulang walau work item dihapus atau dipindah ke project lain.
  3. Label UI "Projects" ditampilkan sebagai "Spaces" (menu Spaces, Space settings, Create space) melalui kunci i18n yang dapat dikonfigurasi; REST API dan JQL tetap memakai kata `project`, dan path URL lama (`/projects/...`) tetap berfungsi berdampingan dengan path `/spaces/...`.
  4. Label "Issue types" tampil sebagai "Work types" dan "Issues" sebagai "Work items" di seluruh UI melalui kamus i18n yang sama (EN/ID), tanpa mengubah nama kolom data atau parameter API.
  5. Project memiliki salah satu status lifecycle `active`, `archived` (khusus plan Premium/Enterprise), `trashed`, atau `deleted` (permanen); status ini disimpan eksplisit dan dipakai oleh pemeriksaan izin, indeks pencarian, dan job latar (lihat PRJ-031).
  6. Project dapat dirujuk oleh board (melalui filter), filter tersimpan, gadget dashboard, plan, scope automation rule, webhook, dan JQL (`project = KEY`, `project = 10001`, `project = "Nama"`); rujukan internal disimpan memakai `id`, bukan key, agar tahan terhadap ganti key.
  7. Satu site menampung banyak project lintas produk (software/business, service, product discovery) yang semuanya tampil di satu project directory (PRJ-025), dengan filter per produk.
  8. Setelah key diganti, key lama disimpan sebagai alias (tabel `project_key_aliases`) yang dipakai untuk redirect URL dan resolusi JQL/issue key lama (lihat PRJ-022).
- **Team-managed vs company-managed:** Atribut entitas sama; nilai `style` membedakan apakah konfigurasi diambil dari scheme global atau dari konfigurasi project-scoped.
- **Catatan & ketergantungan:** Definisi entitas final dan event domain ada di [28 Model Data Inti](28-model-data-inti.md); kamus istilah di [00](00-visi-ruang-lingkup-glosarium-baseline.md). Repo saat ini memakai UUID `varchar(36)` untuk `users.id`; untuk project disarankan `bigint` identity sebagai `id` publik (kompatibel dengan API bergaya Jira) — keputusan konsistensi tipe ID dicatat di Pertanyaan terbuka. Format path URL final (`/projects/...` vs `/spaces/...`) pada referensi belum terkonfirmasi; clone mendukung keduanya sebagai alias route (butir 3).

### PRJ-002 Project Types (Software vs Business)

- **Deskripsi:** Produk terpadu memiliki tipe project `software` dan `business`; tipe `service_desk` dan `product_discovery` tersedia bila produk terkait dilisensikan. Tipe menentukan fitur default, view, field default, dan katalog template.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs dan pembuat project memilih tipe lewat template; semua pengguna melihat tipe di directory.
- **Aturan perilaku (acceptance criteria):**
  1. Project `software` menyediakan backlog, sprint, board Scrum/Kanban, development panel, Releases, Deployments, Code, Security, On-call, dan laporan agile (masing-masing dapat dinonaktifkan via Features, PRJ-049/PRJ-050).
  2. Project `business` menyediakan view List, Board, Calendar, Timeline, Summary, Forms, Approvals, Goals, Attachments, dan Pages; sprint dan backlog Scrum **tidak** aktif secara default.
  3. Semua pengguna dengan product access produk inti dapat membuat dan memakai project `software` maupun `business` tanpa lisensi tambahan.
  4. Project directory menampilkan kolom "Type" dan menyediakan filter tipe/produk (PRJ-025).
  5. Tipe project yang dipilih menentukan template yang tampil di create wizard (template software hanya muncul di kategori Software development, dst.).
  6. Project `service_desk` dan `product_discovery` hanya dapat dibuka oleh pengguna dengan product access produk tersebut; pengguna lain menerima halaman "no access" (PRJ-004).
  7. `projectTypeKey` tersedia di REST API: daftar tipe (`GET /project/type`), detail tipe (`/project/type/{key}`), dan tipe yang dapat diakses caller (`/project/type/{key}/accessible`).
  8. Field default berbeda per tipe: project `business` tidak memiliki field Sprint dan Story points pada layar default; project `software` memilikinya.
  9. ⚠ JQL menyediakan field bawaan `projectType` (mis. `projectType = business`) untuk memfilter work item berdasarkan tipe project.
- **Team-managed vs company-managed:** Kedua tipe tersedia di kedua model; perbedaan model dibahas di PRJ-003.
- **Catatan & ketergantungan:** Tipe service dirinci di [29](29-jsm-service-project-request-type-formulir.md), tipe product discovery di [36](36-jpd-ideas-field-views.md). View business di [26](26-kerja-bisnis-formulir-persetujuan.md) dan [18](18-timeline-rencana-tampilan-proyek.md).

### PRJ-003 Team-managed vs Company-managed Projects

- **Deskripsi:** Dua model manajemen project: team-managed dikonfigurasi Administrator project secara terisolasi per project; company-managed memakai scheme global yang dikelola Admin Situs dan dapat dibagi antar project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (company-managed), Administrator project (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Team-managed: work type, custom field, workflow, dan status bersifat project-scoped (memiliki `project_id`) dan tidak dapat dipilih atau dibagi ke project lain.
  2. Team-managed: akses diatur lewat access level Open/Limited/Private dan role Administrator/Member/Viewer, ditambah custom role pada plan Standard ke atas (PRJ-038, PRJ-039).
  3. Team-managed: satu project memiliki tepat satu board; feature Sprints menentukan board berperilaku Scrum (aktif) atau Kanban (nonaktif).
  4. Team-managed: menyediakan editor workflow sederhana per work type (PRJ-052) dan field editor drag-and-drop (PRJ-051).
  5. Company-managed: konfigurasi diambil dari permission scheme, notification scheme, issue security scheme, workflow scheme, issue type scheme, issue type screen scheme, field configuration scheme, dan priority scheme yang dipasang ke project.
  6. Company-managed: dapat memiliki banyak board, termasuk board lintas project berbasis filter.
  7. Company-managed: hanya Admin Situs (global permission Administer) yang dapat membuat project dan mengganti scheme; Admin Proyek terbatas pada hal di PRJ-036 kecuali memiliki Extended project administration.
  8. Fitur khusus company-managed: Issue collectors, issue security scheme, field configuration, screens, banyak board, dan sharing konfigurasi; menu ini tidak tampil di project team-managed.
  9. Field estimasi berbeda: team-managed memakai "Story point estimate", company-managed memakai "Story Points"; custom field team-managed terpisah per project walau namanya sama dengan field di project lain.
  10. Hierarki: team-managed Epic > work type standar > Subtask; company-managed mengikuti hierarki global (level di atas Epic lewat Plans, lihat [05](05-konfigurasi-work-item-skema.md)).
  11. Model tidak dapat dikonversi langsung; sistem tidak menyediakan aksi "convert", dan migrasi dilakukan dengan memindahkan work item (PRJ-024).
  12. Bagian bawah sidebar project menampilkan label "You're in a team-managed project" / "You're in a company-managed project" dengan tautan penjelasan; directory menampilkan indikator model.
  13. Create wizard menampilkan tabel perbandingan kedua model (siapa yang mengonfigurasi, sharing, fitur khusus) sebelum pengguna memilih.
- **Team-managed vs company-managed:** Fitur ini sendiri adalah definisi perbedaannya.
- **Catatan & ketergantungan:** Ini keputusan arsitektur terbesar. Model data harus mendukung dua sumber konfigurasi: tabel scheme global dan tabel konfigurasi project-scoped. Resolver konfigurasi ("config resolver") tunggal di domain layer disarankan agar use case work item tidak bercabang. Lihat [03](03-izin-peran-keamanan.md), [05](05-konfigurasi-work-item-skema.md), [09](09-workflow-status-transisi.md), [10](10-board-kolom-kartu.md).

### PRJ-004 Product Licensing Gate for Projects

- **Deskripsi:** Kontrol akses project berdasarkan product access pengguna dan plan site, termasuk halaman "no access" yang konsisten.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Admin Organisasi/Admin Situs mengatur product access; pengguna terdampak saat membuka project.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna tanpa product access produk inti tidak dapat membuka project software/business (UI menampilkan halaman no-access; API mengembalikan 403/404 sesuai kebijakan anti-enumerasi).
  2. Project service dan product discovery memerlukan product access produk masing-masing; pemeriksaan dilakukan di server untuk setiap request project-scoped.
  3. Fitur Premium/Enterprise (archive project, custom template, archived work items) disembunyikan atau ditampilkan dengan ajakan upgrade sesuai feature flag plan; endpoint terkait mengembalikan galat "fitur tidak tersedia di plan ini".
  4. Plan Free: maksimal 10 pengguna berlisensi; permission dan role project tidak dapat disunting; issue security tidak tersedia.
  5. Project yang tidak dapat diakses menampilkan halaman "You don't have access to this project/space"; bila berlaku, halaman memuat tombol "Request access" (PRJ-038) atau ajakan meminta product access ke admin (lihat [02](02-pengguna-grup-organisasi.md)).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme tier/feature flag ditetapkan di [00](00-visi-ruang-lingkup-glosarium-baseline.md) dan [27](27-nfr-arsitektur-platform.md); product access di [02](02-pengguna-grup-organisasi.md).

## Template project

### PRJ-005 Project Template Library

- **Deskripsi:** Galeri template yang dipakai saat membuat project, dikelompokkan per kategori kerja, dengan halaman preview yang menjelaskan work type, workflow, dan view yang akan dibuat.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan izin create project (PRJ-013); Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Sidebar galeri menampilkan kategori: Software development, Work management/Project management, Marketing, Design, HR, Finance, Operations, Legal, Sales, Personal; ditambah Service management dan Product management bila produk terkait dilisensikan.
  2. Halaman detail template memuat deskripsi, ilustrasi, daftar work type beserta ikon, preview workflow (status dan transisi), view/fitur yang aktif, dan model manajemen yang didukung.
  3. Tombol "Use template" melanjutkan ke langkah pilih model dan detail project (PRJ-010).
  4. Template yang hanya mendukung satu model (mis. Bug tracking hanya company-managed) menampilkan keterangan tersebut dan wizard melewati langkah pilih model.
  5. Template menentukan board type, work type, workflow, screen, dan field default yang dibuat otomatis saat project dibuat; definisi template disimpan sebagai data deklaratif (seed) yang diberi versi.
  6. Template untuk produk yang tidak dilisensikan tampil dengan ajakan "Try" dan tombol "Use template" dinonaktifkan.
  7. Template kustom organisasi (PRJ-009) tampil di galeri dengan label organisasi bila fitur tersebut tersedia.
  8. ⚠ Galeri menampilkan badge "Recommended" dan/atau daftar template yang terakhir dipakai pengguna.
- **Team-managed vs company-managed:** Template mendeklarasikan model yang didukung; lihat butir 4.
- **Catatan & ketergantungan:** Isi galeri sebaiknya dapat diperluas tanpa rilis kode (template sebagai berkas JSON/YAML di-embed).

### PRJ-006 Software Project Templates

- **Deskripsi:** Template untuk tim software: Scrum, Kanban, Bug tracking, DevOps, dan Cross-team/Top-level planning, yang menyiapkan board, work type, dan workflow sesuai praktik agile.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs; pengguna dengan global permission Create team-managed projects.
- **Aturan perilaku (acceptance criteria):**
  1. Template Scrum membuat backlog, sprint, board Scrum, work type Epic/Story/Task/Bug/Subtask, dan laporan sprint (burndown, velocity); tersedia untuk team-managed dan company-managed.
  2. Template Kanban membuat board kontinu dengan kolom To Do/In Progress/Done, dukungan WIP limit, dan backlog Kanban opsional (nonaktif default); tersedia untuk team-managed dan company-managed.
  3. Template Bug tracking hanya company-managed, tanpa board agile secara default, dengan work type Bug/Task/Subtask.
  4. ⚠ Template DevOps adalah company-managed berbasis Kanban dengan integrasi development, deployment, dan operasi (incident/on-call) yang disiapkan di Toolchain.
  5. ⚠ Template Cross-team planning/Top-level planning membuat project untuk work level atas lintas tim (mis. Initiative) yang dipakai bersama Plans; ketersediaannya mengikuti plan Premium.
  6. ⚠ Workflow default Scrum/Kanban adalah To Do → In Progress → Done (semua transisi diizinkan); workflow Bug tracking memakai status default yang didefinisikan di seed template.
  7. Setelah memilih template software, wizard menawarkan langkah opsional "Connect your tools" (GitHub, GitLab, Bitbucket, dll.) yang dapat dilewati.
  8. Untuk template yang mendukung kedua model, wizard menampilkan pilihan Team-managed atau Company-managed.
- **Team-managed vs company-managed:** Scrum/Kanban tersedia di keduanya; Bug tracking dan DevOps hanya company-managed.
- **Catatan & ketergantungan:** MVP Fase 1 minimal Scrum dan Kanban untuk kedua model; Bug tracking, DevOps, dan Top-level planning menyusul (Fase 2-4). Board di [10](10-board-kolom-kartu.md), sprint di [11](11-backlog-sprint.md), integrasi di [24](24-integrasi-devops.md), Plans di [18](18-timeline-rencana-tampilan-proyek.md).

### PRJ-007 Business Project Templates

- **Deskripsi:** Template untuk tim non-software yang menyiapkan work type, workflow, field contoh, dan view sesuai proses bisnis.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan izin create project; Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Template Project management membuat work type Task/Subtask dengan view List/Board/Calendar/Timeline aktif.
  2. Template Task tracking membuat workflow sederhana To Do/In Progress/Done.
  3. Template Process control membuat workflow dengan tahapan proses formal (status berurutan, transisi linear).
  4. Template Content management membuat work type konten dengan alur draft → review → published.
  5. Template Document approval membuat alur persetujuan dokumen yang terintegrasi dengan Approvals ([26](26-kerja-bisnis-formulir-persetujuan.md)).
  6. Kategori Marketing menyediakan template Campaign management, Email campaign, Event planning, dan Go-to-market.
  7. Kategori Sales menyediakan template Lead tracking dan Sales pipeline.
  8. Kategori HR menyediakan template Recruitment, Employee/Personnel onboarding, Performance review, dan Policy management.
  9. ⚠ Kategori Finance menyediakan template Procurement, Budget creation, dan Month-end close.
  10. ⚠ Kategori Legal/Operations/Design masing-masing menyediakan template permintaan atau proses khas kategori tersebut.
  11. ⚠ Sebagian besar template business tersedia untuk team-managed dan company-managed; model yang didukung dideklarasikan per template.
  12. Template dapat menyertakan form default, custom field contoh, dan contoh work item; contoh work item dapat dihapus massal oleh admin.
- **Team-managed vs company-managed:** Lihat butir 11.
- **Catatan & ketergantungan:** Forms/Approvals di [26](26-kerja-bisnis-formulir-persetujuan.md). Daftar template di atas adalah target katalog; nama dan isi final ditetapkan pemilik produk.

### PRJ-008 Blank Project Template

- **Deskripsi:** Template kosong/minimal untuk memulai project tanpa konfigurasi khusus.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan izin create project; Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Blank project business membuat work type Task (dan Subtask) dengan workflow To Do/In Progress/Done.
  2. ⚠ Tersedia varian minimal untuk software team-managed (board Kanban tanpa sprint, work type Task/Subtask).
  3. Setelah dibuat, pengguna berizin dapat menambah work type, field, dan workflow seperti project lain.
  4. Template blank tidak membuat contoh work item maupun form.
  5. ⚠ Model manajemen yang tersedia untuk template blank dideklarasikan di seed template (default: team-managed dan company-managed).
- **Team-managed vs company-managed:** Lihat butir 5.
- **Catatan & ketergantungan:** Berguna sebagai template uji otomatis dan dasar API create tanpa `projectTemplateKey`.

### PRJ-009 Custom Project Templates (Enterprise)

- **Deskripsi:** Admin dapat menyimpan project yang ada sebagai template kustom organisasi (custom space template), lalu memakainya sendiri atau memberi pengguna/tim tertentu akses self-serve untuk membuat project dari template yang disetujui. Nama heading mengikuti inventaris; ⚠ sejak rilis 26 Oktober 2026 fitur ini tidak lagi eksklusif plan Enterprise, melainkan tersedia di semua paid plan.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard (khusus paid plan Standard ke atas, juga Premium dan Enterprise; tidak tersedia di Free; gating lewat VISI-003)
- **Aktor & izin:** Admin Situs (membuat, mengelola, dan memberi akses template); pengguna, grup, atau tim non-admin yang diberi akses self-serve ke template tertentu.
- **Aturan perilaku (acceptance criteria):**
  1. Menu More actions (…) project menyediakan "Save as template" yang menyimpan konfigurasi project (work type, workflow, field, layar, features, board) tanpa work item.
  2. Template kustom dapat dibuat dari project software maupun business, team-managed maupun company-managed.
  3. Template kustom tampil di galeri create project dengan label organisasi (PRJ-005).
  4. Delegated space creation: admin dapat mengizinkan pengguna/grup non-admin membuat project dari template kustom tertentu tanpa global permission Administer.
  5. ⚠ Untuk company-managed, project yang dibuat dari template kustom menghasilkan salinan scheme baru (bukan berbagi scheme dengan project sumber).
  6. Admin dapat mengedit metadata, menghapus template, dan mengatur siapa yang dapat memakainya.
  7. ⚠ Tersedia API "Create custom project" untuk membuat project beserta konfigurasi penuh dalam satu panggilan.
  8. ⚠ Ketersediaan plan: template kustom tersedia di semua paid plan (Standard, Premium, Enterprise). Di plan Free, menu "Save as template" disembunyikan, galeri tidak menampilkan template kustom, dan pengguna melihat keterangan plan minimum (VISI-003).
  9. ⚠ Akses self-serve: admin dapat menandai template sebagai disetujui (approved) dan memberi tim/grup tertentu akses self-serve. Hanya template approved yang dapat dipakai lewat akses ini, dan pengguna di luar daftar akses tidak melihat template tersebut di galeri.
  10. ⚠ Template dapat ditautkan ke satu Space policy (PRJ-069) melalui project sumbernya. Project team-managed yang dibuat dari template mewarisi policy tersebut, sedangkan template tanpa policy memakai policy default site.
- **Team-managed vs company-managed:** Lihat butir 2 dan 5; policy pada butir 10 hanya berlaku untuk project team-managed.
- **Catatan & ketergantungan:** Inventaris masih menandai tier enterprise. Perubahan ke paid plan (pengumuman rilis 26 Oktober 2026) dan akses self-serve tim (April 2026) dicatat dengan ⚠, sehingga prioritas dinaikkan ke P2 / Fase 3. Guardrail template melalui Space Policies di PRJ-069. Clone project lengkap beserta data (seperti app Marketplace "Deep Clone") di luar cakupan. Delegasi izin di [03](03-izin-peran-keamanan.md); gating plan di [00](00-visi-ruang-lingkup-glosarium-baseline.md) (VISI-003); endpoint `POST /project-template` di API-036 ([23](23-api-publik-webhook.md)).

## Pembuatan project & validasi

### PRJ-010 Create Project Wizard

- **Deskripsi:** Alur multi-langkah untuk membuat project: pilih template, pilih model manajemen, isi detail, lalu langkah opsional seperti menghubungkan tools dan mengundang tim.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (company-managed dan team-managed); pengguna dengan global permission Create team-managed projects (team-managed saja).
- **Aturan perilaku (acceptance criteria):**
  1. Wizard dapat dibuka dari tombol "Create project/space" di project directory, menu Projects/Spaces di sidebar, tombol "+" di navigasi global, dan halaman detail template; entry point disembunyikan bagi pengguna tanpa izin create.
  2. Langkah 1: pengguna memilih template dari galeri beserta preview (PRJ-005).
  3. Langkah 2: pengguna memilih Team-managed atau Company-managed dengan tabel perbandingan; opsi yang tidak diizinkan untuk pengguna (mis. company-managed bagi non-Admin Situs) dinonaktifkan dengan tooltip alasan.
  4. Langkah 3: Name (wajib), Key (auto-generate, dapat diedit), Access level (hanya team-managed, default Open), dan checkbox "Share settings with an existing project" (hanya company-managed, PRJ-014).
  5. Dari langkah detail, tautan "Change template" mengembalikan pengguna ke pemilihan template/model tanpa kehilangan nilai Name/Key yang sudah diisi.
  6. Validasi nama dan key berjalan inline secara real-time (debounce ≤ 500 ms) dan menampilkan pesan spesifik untuk duplikat, format salah, panjang salah, atau key yang sudah terpakai sebagai alias; tombol Create nonaktif selama ada galat.
  7. Langkah opsional setelah detail: "Connect your tools" (hanya template software) dan "Add teammates/Invite people"; keduanya dapat dilewati.
  8. Project lead default adalah pembuat project.
  9. Setelah berhasil, pengguna diarahkan ke view default template (Board, Backlog, Summary, atau List) dan panduan onboarding ditampilkan (PRJ-015).
  10. Pembuatan bersifat atomik dalam satu transaksi: project, board, scheme/konfigurasi project-scoped hasil template, dan membership pembuat sebagai role Administrator; bila salah satu gagal, tidak ada data parsial yang tersisa dan pengguna melihat galat umum yang dapat dicoba ulang.
  11. Project dapat dibuat lewat REST API `POST /project` dengan field `key`, `name`, `projectTypeKey`, `projectTemplateKey`, `leadAccountId`, `assigneeType`, dan ID scheme opsional; validasi sama dengan wizard.
- **Team-managed vs company-managed:** Access level hanya untuk team-managed; share settings hanya untuk company-managed.
- **Catatan & ketergantungan:** Undangan pengguna memakai alur di [02](02-pengguna-grup-organisasi.md); event `project_created` di [23](23-api-publik-webhook.md) dan [28](28-model-data-inti.md).

### PRJ-011 Project Name Rules

- **Deskripsi:** Aturan validasi nama project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pembuat project, Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Nama wajib diisi (setelah trim spasi) dan unik di site secara case-insensitive; pesan galat: "Nama project sudah dipakai".
  2. ⚠ Panjang maksimum nama 80 karakter; input lebih panjang ditolak dengan pesan yang menyebut batasnya.
  3. ⚠ Panjang minimum nama 2 karakter.
  4. Nama boleh berisi spasi, angka, tanda baca, dan karakter Unicode (termasuk huruf non-Latin dan emoji); karakter kontrol ditolak.
  5. Nama dapat diubah kapan saja oleh Admin Proyek tanpa mengubah key atau issue key.
  6. Nama dipakai sebagai dasar auto-generate key (PRJ-012).
  7. Nama dapat dipakai di JQL (`project = "Nama Project"`); perubahan nama langsung berlaku di resolusi JQL.
  8. ⚠ Nama yang sudah dipakai project lain, termasuk project di trash atau archive, ditolak.
  9. Endpoint `GET /projectvalidate/validProjectName?name=` mengembalikan nama yang diminta bila valid atau saran nama valid terdekat (mis. menambahkan sufiks angka).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi panjang/pola tidak tersedia (lihat PRJ-068, khas Data Center).

### PRJ-012 Project Key Rules & Auto-generation

- **Deskripsi:** Key project adalah identifier pendek yang menjadi prefiks issue key, dengan aturan format dan keunikan yang ketat.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pembuat project, Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Format key: diawali huruf kapital Latin `A-Z`, diikuti huruf kapital atau angka (regex `^[A-Z][A-Z0-9]+$`); dukungan underscore diputuskan pemilik produk (default: tidak diizinkan).
  2. Panjang key 2 sampai 10 karakter.
  3. Key di-auto-generate dari inisial kata nama project (mis. "Mobile App" → `MA`) atau potongan nama untuk nama satu kata (mis. "Phoenix" → `PHOE`), dan dapat diedit sebelum project dibuat; bila hasil sudah terpakai, sistem menambah huruf/angka hingga unik.
  4. Input key otomatis diubah ke huruf kapital saat diketik; karakter di luar format dibuang dari hasil auto-generate.
  5. Key unik di site; tidak boleh sama dengan key aktif project lain maupun key lama (alias) project lain.
  6. Key milik project di trash atau archive tetap dicadangkan dan tidak dapat dipakai project baru.
  7. Pesan validasi spesifik per pelanggaran: format ("Key harus diawali huruf dan hanya berisi huruf kapital/angka"), panjang ("Key harus 2–10 karakter"), dan sudah dipakai ("Key sudah dipakai oleh project lain").
  8. Endpoint `GET /projectvalidate/key?key=` mengembalikan daftar galat validasi, dan `GET /projectvalidate/validProjectKey?key=` mengembalikan saran key valid.
  9. Key dapat diubah setelah project dibuat melalui alur PRJ-022.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kata cadangan (CON, NUL, AUX, dll.) dan pola key yang dapat dikonfigurasi adalah perilaku Data Center (PRJ-068) dan tidak dibangun. Penomoran issue key di [06](06-work-item-inti.md).

### PRJ-013 Create Project Permissions

- **Deskripsi:** Global permission yang menentukan siapa yang boleh membuat project untuk tiap model.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs, Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Project company-managed hanya dapat dibuat oleh pemegang global permission "Administer" (Admin Situs); request dari pengguna lain ditolak 403.
  2. Project team-managed dapat dibuat oleh pemegang global permission "Create team-managed projects", yang secara default diberikan ke grup pengguna produk.
  3. Admin Situs dapat mencabut "Create team-managed projects" dari grup default untuk membatasi pertumbuhan project; perubahan berlaku segera di wizard (entry point hilang).
  4. Pembuat project team-managed otomatis menjadi role Administrator di project tersebut.
  5. Plan Enterprise: pembuatan dari template kustom dapat didelegasikan ke pengguna non-admin (PRJ-009).
  6. Permission ini diatur di Settings > System > Global permissions.
- **Team-managed vs company-managed:** Lihat butir 1-2.
- **Catatan & ketergantungan:** Definisi global permission di [03](03-izin-peran-keamanan.md). Migrasi: field `Role` pada entitas `User` saat ini (`varchar(50)`, default `user`; belum tercantum di migrasi `001_create_users_table`) dapat dipetakan sementara ke global permission "Administer" untuk nilai admin sampai modul 03 tersedia.

### PRJ-014 Share Settings with an Existing Project

- **Deskripsi:** Saat membuat project company-managed, Admin Situs dapat memakai scheme project lain. Konfigurasi dibagi (bukan disalin) sehingga perubahan berlaku di semua project yang berbagi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Create wizard company-managed menampilkan checkbox "Share settings with an existing project".
  2. Bila dicentang, muncul dropdown project sumber yang hanya berisi project company-managed aktif yang dapat diakses admin.
  3. Project baru memakai scheme yang sama dengan sumber: issue type, workflow, issue type screen, field configuration, permission, notification, issue security, dan priority scheme (referensi ke scheme yang sama, bukan salinan).
  4. ⚠ Work item, versions, components, board, shortcut, dan role membership tidak disalin dari project sumber.
  5. ⚠ Project lead project baru adalah pembuat, bukan lead project sumber.
  6. Perubahan scheme setelahnya berdampak ke semua project yang berbagi; halaman settings menampilkan indikator "Shared with N projects" beserta daftar project.
  7. Opsi ini tidak tersedia untuk team-managed (checkbox tidak dirender).
  8. Tipe project baru mengikuti template yang dipilih dan harus kompatibel dengan tipe project sumber; sumber dengan tipe tidak kompatibel tidak tampil di dropdown.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Copy project lengkap termasuk data bukan fitur native (kecuali PRJ-009). Scheme di [03](03-izin-peran-keamanan.md), [05](05-konfigurasi-work-item-skema.md), [09](09-workflow-status-transisi.md).

### PRJ-015 Project Onboarding & Setup Guidance

- **Deskripsi:** Panduan awal setelah project dibuat untuk membantu tim memulai.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Pengguna dan Admin Proyek yang baru membuat/membuka project.
- **Aturan perilaku (acceptance criteria):**
  1. Checklist "Get started" menampilkan langkah buat work item, undang tim, mulai sprint (Scrum), dan hubungkan tools; langkah ditandai selesai otomatis saat aksi dilakukan.
  2. Tooltip/spotlight fitur tampil pada kunjungan pertama board dan backlog, satu kali per pengguna.
  3. Pengguna dapat menutup panduan; status tertutup disimpan per pengguna per project dan tidak muncul lagi.
  4. ⚠ Sebagian template membuat contoh work item yang dapat dihapus massal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Preferensi personal di [21](21-navigasi-ux-preferensi.md).

## Detail & identitas project

### PRJ-016 Project Details Settings (company-managed)

- **Deskripsi:** Halaman Project settings > Details untuk mengubah atribut dasar project company-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pemegang project permission Administer Projects; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman memuat field Name, Key, URL, Project type, Project category, Avatar, Description, Project lead, dan Default assignee; field Project type ditampilkan read-only kecuali pemilik produk memutuskan tipe dapat diubah (PRJ-023).
  2. URL bersifat opsional dan harus berupa URL valid dengan skema `http` atau `https`; nilai lain ditolak dengan pesan "URL tidak valid".
  3. ⚠ Description bersifat opsional, berupa teks biasa dengan batas panjang yang ditetapkan (usulan 1.000 karakter).
  4. Tombol Save menjalankan validasi yang sama seperti saat create (PRJ-011, PRJ-012) dan menampilkan galat per field tanpa menyimpan perubahan parsial.
  5. Mengubah Key menampilkan peringatan dampak sebelum disimpan (lihat PRJ-022).
  6. Pilihan Default assignee hanya menampilkan "Unassigned" bila global setting "Allow unassigned work items" aktif.
  7. Project lead harus pengguna aktif dengan product access; pemilihan pengguna nonaktif ditolak.
  8. Menu "…" di header halaman menyediakan Move to trash, Archive (Premium/Enterprise), dan Save as template (Enterprise), masing-masing hanya tampil bila pengguna berizin dan plan mendukung.
  9. ⚠ Halaman hanya dapat diakses pemegang Administer Projects dan Admin Situs; perubahan Key dan Category dibatasi ke Admin Situs bila pemilik produk memutuskan demikian.
- **Team-managed vs company-managed:** Versi team-managed di PRJ-017.
- **Catatan & ketergantungan:** Perubahan dicatat di audit log ([22](22-administrasi-sistem.md)) dan memicu event `project_updated`.

### PRJ-017 Project Details Settings (team-managed)

- **Deskripsi:** Halaman Details untuk project team-managed dengan field yang lebih sederhana.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Role Administrator project team-managed; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman memuat field Name, Key, Avatar/icon, Project category, Description, Project lead, dan Default assignee (pilihan Project lead/Unassigned).
  2. Project lead harus member project; memilih non-member ditolak dengan pesan "Project lead harus anggota project".
  3. Menu aksi menyediakan Move to trash dan Archive (Premium/Enterprise).
  4. Perubahan disimpan dengan tombol Save dengan validasi yang sama seperti saat create.
  5. Halaman dapat diakses role Administrator project dan Admin Situs; role lain menerima 403.
- **Team-managed vs company-managed:** Versi company-managed di PRJ-016.
- **Catatan & ketergantungan:** Role team-managed di PRJ-038.

### PRJ-018 Project Avatar

- **Deskripsi:** Gambar atau ikon yang mengidentifikasi project di directory, sidebar, tampilan work item, dan picker.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat memilih avatar dari set avatar bawaan sistem.
  2. Pengguna dapat mengunggah gambar kustom PNG/JPG/GIF melalui dialog crop persegi; format lain ditolak.
  3. ⚠ Ukuran berkas unggahan dibatasi (usulan maksimum 1 MB); berkas melebihi batas ditolak dengan pesan yang menyebut batasnya.
  4. Avatar kustom yang pernah diunggah dapat dihapus; bila avatar aktif dihapus, project kembali ke avatar default.
  5. Avatar disajikan dalam ukuran 16, 24, 32, dan 48 px (`avatarUrls`) yang dihasilkan server saat unggah.
  6. Avatar default ditetapkan otomatis saat project dibuat (dipilih dari set bawaan).
  7. REST API: `GET /project/{id}/avatars`, `POST /project/{id}/avatar2` (unggah), `PUT /project/{id}/avatar` (set aktif), `DELETE /project/{id}/avatar/{avatarId}`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan berkas memakai storage lampiran ([08](08-komentar-lampiran-aktivitas.md), [27](27-nfr-arsitektur-platform.md)).

### PRJ-019 Project Categories

- **Deskripsi:** Kategori untuk mengelompokkan project, dibuat Admin Situs dan dipakai untuk filter, JQL, gadget, dan scope.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengelola kategori; Admin Proyek memilih kategori project.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs mengelola kategori di Settings > Projects > Project categories: Add (name, description), Edit, Delete.
  2. Satu project memiliki paling banyak satu kategori (atau none).
  3. Kategori dipasang ke project melalui halaman Details.
  4. Menghapus kategori membuat semua project terkait menjadi tanpa kategori; dialog konfirmasi menyebut jumlah project terdampak.
  5. Project directory dapat difilter berdasarkan kategori.
  6. JQL mendukung `category = "Nama"`.
  7. ⚠ Kategori dapat dipilih di gadget dashboard (pilihan project/category) dan sebagai scope automation rule.
  8. Nama kategori unik (case-insensitive).
  9. REST API `/projectCategory` menyediakan CRUD (GET daftar, GET by id, POST, PUT, DELETE).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik aturan: ADM-022 di [22](22-administrasi-sistem.md) (pengelolaan kategori di Settings > Projects); modul ini memiliki pemasangan kategori ke project dan filter directory. Bila ada perbedaan, ADM-022 berlaku untuk CRUD kategori. Gadget di [17](17-dashboard-gadget.md), scope otomasi di [20](20-otomasi.md), JQL di [12](12-jql-bahasa-kueri.md).

### PRJ-020 Project Lead

- **Deskripsi:** Pengguna yang menjadi penanggung jawab project dan dapat menjadi default assignee.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Project lead diset saat create (default pembuat) dan dapat diubah di Details.
  2. Project lead harus pengguna aktif dengan product access; pada team-managed juga harus member project.
  3. Bila Default assignee = "Project lead", work item baru tanpa assignee di-assign ke lead.
  4. Project lead menjadi fallback default assignee component bila component memilih "Project lead" atau "Project default" yang bernilai lead.
  5. Project directory menampilkan kolom "Lead" dengan avatar dan nama.
  6. JQL function `projectsLeadByUser()` (default pengguna saat ini) mengembalikan project yang dipimpin pengguna.
  7. ⚠ Bila lead dinonaktifkan atau dihapus, project tetap ada, field lead ditandai "pengguna nonaktif", dan Admin Proyek/Admin Situs mendapat peringatan di Details untuk mengganti lead; default assignee "Project lead" diperlakukan sebagai Unassigned sampai diganti.
  8. Project Lead dapat dipilih sebagai penerima di notification scheme dan sebagai grantee di permission scheme.
- **Team-managed vs company-managed:** Pada team-managed lead harus member (butir 2).
- **Catatan & ketergantungan:** Penonaktifan pengguna di [02](02-pengguna-grup-organisasi.md); grantee di [03](03-izin-peran-keamanan.md).

### PRJ-021 Default Assignee

- **Deskripsi:** Aturan assignee otomatis untuk work item baru bila pembuat tidak mengisi assignee.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi default assignee project: Project lead atau Unassigned.
  2. Opsi Unassigned hanya muncul bila global setting "Allow unassigned work items" aktif.
  3. Default assignee component (Project default, Project lead, Component lead, Unassigned) meng-override default project bila work item memiliki component; bila work item memiliki beberapa component dengan default assignee berbeda, dipakai default dari component pertama menurut urutan nama alfabetis (REL-031 butir 4).
  4. ⚠ Assignee untuk work item baru tanpa assignee eksplisit ("Automatic"/kosong) ditentukan dengan urutan tunggal berikut, berhenti pada langkah pertama yang menghasilkan pengguna assignable: (a) default assignee component (butir 3, REL-031); (b) default assignee project (Project lead); (c) Unassigned, hanya bila global setting "Allow unassigned work items" aktif (CFG-057); (d) bila semua langkah gagal, pembuatan work item ditolak dengan galat "Assignee default tidak dapat ditugaskan". Pengguna yang tidak memiliki permission Assignable User atau nonaktif dianggap gagal dan sistem lanjut ke langkah berikutnya.
  5. Aturan berlaku sama untuk pembuatan lewat UI, email (mail handler), API, impor, dan automation.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik urutan fallback default assignee: fitur ini (butir 4). REL-031 butir 5 di [15](15-rilis-versi-komponen.md) dan ISS-016 butir 3 di [06](06-work-item-inti.md) perlu diselaraskan dengan rujukan ke PRJ-021 butir 4 (termasuk aturan beberapa component, REL-031 butir 4). Global setting CFG-057 di [05](05-konfigurasi-work-item-skema.md).

### PRJ-022 Change Project Key

- **Deskripsi:** Mengubah key project yang sudah ada beserta dampaknya pada semua work item dan referensi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Key diubah di Project settings > Details oleh Admin Proyek/Admin Situs, baik di company-managed maupun team-managed.
  2. Semua issue key diperbarui (mis. `ABC-12` menjadi `XYZ-12`) dengan nomor urut tetap.
  3. Key lama disimpan sebagai alias: URL lama redirect (HTTP 301) ke key baru, pencarian `ABC-12` menemukan `XYZ-12`, dan JQL yang memakai key lama tetap berfungsi.
  4. Key lama tidak dapat dipakai project lain selama alias tercatat.
  5. Project dapat kembali memakai key lamanya sendiri; alias tersebut kemudian menjadi key aktif.
  6. ⚠ Proses berjalan sebagai background task; selama berjalan project ditandai "sedang mengganti key" dan operasi tulis dapat ditolak sementara dengan pesan yang jelas.
  7. ⚠ Commit, branch, dan PR yang memakai key lama tetap tertaut ke work item melalui resolusi alias.
  8. Filter tersimpan, board, dan automation tetap berfungsi karena merujuk `id` project, bukan key.
  9. Dialog konfirmasi menampilkan peringatan dampak (bookmark, integrasi eksternal, balasan email ke notifikasi lama).
  10. Perubahan key tercatat di audit log dengan key lama, key baru, dan pelaku.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Indeks pencarian perlu re-index issue key ([13](13-pencarian-filter-ekspor.md)); integrasi DVCS di [24](24-integrasi-devops.md). Implementasi disarankan menyimpan `issue_number` terpisah dari key agar update key cukup di tabel project.

### PRJ-023 Change Project Type / Board Style

- **Deskripsi:** Mengubah tipe atau gaya kerja project setelah project dibuat.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Team-managed software beralih Kanban ↔ Scrum dengan toggle feature Sprints (dan Backlog) di Features (PRJ-049).
  2. Company-managed: tipe board yang ada tidak diubah; untuk beralih, pengguna membuat board Scrum/Kanban baru untuk project.
  3. ⚠ Company-managed: perubahan project type software/business/service hanya diizinkan bila pemilik produk memutuskan mendukungnya; bila didukung, hanya Admin Situs yang dapat melakukannya dan fitur khas tipe lama disembunyikan.
  4. Team-managed tidak dapat diubah antar tipe software dan business.
  5. Tidak ada konversi model team-managed ↔ company-managed secara langsung (lihat PRJ-024).
- **Team-managed vs company-managed:** Lihat butir 1-4.
- **Catatan & ketergantungan:** Board di [10](10-board-kolom-kartu.md).

### PRJ-024 Project Migration between Team-managed and Company-managed

- **Deskripsi:** Memindahkan work item antara project team-managed dan company-managed sebagai cara "konversi" model.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek (dengan izin bulk move di kedua project).
- **Aturan perilaku (acceptance criteria):**
  1. Alur migrasi: buat project tujuan dengan model lain, lalu bulk move semua work item dari project sumber.
  2. Wizard bulk move mewajibkan mapping work type dan status sumber ke tujuan sebelum eksekusi.
  3. ⚠ Wizard menampilkan daftar data yang berisiko hilang (custom field project-scoped, Story point estimate ke Story Points, rank, relasi parent/epic, sprint, detail board), sedangkan komentar, lampiran, dan history tetap dipertahankan.
  4. Issue key berubah mengikuti project tujuan; key lama di-redirect ke key baru.
  5. Produk menyediakan checklist migrasi bawaan (dokumentasi in-app) yang mengikuti urutan langkah di atas.
  6. Project sumber dapat dipindahkan ke trash setelah migrasi selesai.
- **Team-managed vs company-managed:** Fitur ini sendiri adalah jalur perpindahan model.
- **Catatan & ketergantungan:** Bulk move di [06](06-work-item-inti.md). Pemetaan otomatis Story point estimate → Story Points disarankan sebagai keputusan terbuka.

## Direktori, navigasi & halaman project

### PRJ-025 Project Directory

- **Deskripsi:** Halaman "Projects/Spaces" yang menampilkan semua project yang dapat diakses pengguna, beserta pencarian, filter, dan aksi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna berlisensi; aksi baris sesuai izin Admin Proyek/Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tabel menampilkan kolom Star, Name (dengan avatar), Key, Type, Lead, dan Last work item update; kolom Category dan URL ditampilkan opsional.
  2. Kotak pencarian mencocokkan nama atau key (case-insensitive, substring).
  3. Filter tersedia untuk produk (software/business, service, product discovery), project type, dan category; filter dapat dikombinasikan.
  4. Tabel dapat diurutkan berdasarkan name, key, type, dan lead (asc/desc).
  5. Daftar dipaginasi (default 25 baris per halaman) dengan total jumlah project.
  6. Menu aksi "…" per baris hanya menampilkan aksi yang diizinkan: Project settings, Move to trash, Archive.
  7. Halaman menampilkan tombol "Create project/space" (bila berizin, PRJ-013) dan tautan ke galeri template.
  8. Hanya project dengan Browse Projects (company-managed) atau akses sesuai access level (team-managed) yang tampil; project Private tidak terlihat oleh non-member, termasuk dalam hitungan total.
  9. Admin Situs melihat tautan ke daftar Archived dan Trash (PRJ-063).
  10. Bila belum ada project yang dapat diakses, halaman menampilkan empty state dengan ajakan membuat project (atau meminta akses bila tidak berizin create).
- **Team-managed vs company-managed:** Logika visibilitas berbeda (butir 8); tampilan sama.
- **Catatan & ketergantungan:** Endpoint `GET /project/search` (PRJ-064). Navigasi global di [21](21-navigasi-ux-preferensi.md).

### PRJ-026 Starred & Recent Projects

- **Deskripsi:** Akses cepat ke project yang sering dipakai lewat tanda bintang dan daftar recent.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Setiap pengguna untuk dirinya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat star/unstar project dari directory, header project, atau menu Projects/Spaces di sidebar; status bintang berubah tanpa reload.
  2. Menu/sidebar Projects menampilkan section "Starred" dan "Recent".
  3. Recent dihitung dari project yang terakhir dikunjungi pengguna (diurutkan waktu kunjungan terbaru).
  4. Starred bersifat personal per pengguna dan disimpan di server sehingga sinkron antar perangkat.
  5. ⚠ Jumlah item yang tampil di menu dibatasi (usulan 5 starred dan 5 recent) dengan tautan "View all projects".
  6. Project recent menjadi pilihan default project di dialog Create work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Project yang kehilangan akses, di-trash, atau diarsipkan otomatis tersembunyi dari Starred/Recent. Lihat [21](21-navigasi-ux-preferensi.md).

### PRJ-027 Project Navigation & Views

- **Deskripsi:** Navigasi di dalam project berupa sidebar dan tab horizontal yang menampilkan view dan halaman yang aktif.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna berakses; Admin Proyek mengatur Features.
- **Aturan perilaku (acceptance criteria):**
  1. Project software dapat menampilkan: Summary, Timeline, Backlog, Board/Active sprints, Calendar, List, Reports, Forms, Work items (navigator), Code, Releases, Deployments, Security, On-call, Components, Pages, Shortcuts, dan Archived work items (Premium).
  2. Project business dapat menampilkan: Summary, List, Board, Calendar, Timeline, Approvals, Forms, Goals, Pages, Attachments, Reports, dan Shortcuts.
  3. View ditampilkan sebagai tab horizontal di header project; tab yang tidak muat lebar layar dipindahkan ke menu "More".
  4. ⚠ Pengguna dapat mengatur urutan dan menyembunyikan tab; pengaturan berlaku per pengguna (Admin Proyek dapat menetapkan urutan default project).
  5. View yang tampil bergantung pada Features toggle (PRJ-049/PRJ-050) dan permission pengguna; view yang tidak tersedia tidak dirender dan route-nya mengembalikan 404 dalam aplikasi.
  6. Company-managed dengan banyak board menampilkan board switcher di header board.
  7. Sidebar memuat tautan "Project settings" (hanya bagi admin) dan label model manajemen (PRJ-003).
  8. Header project memuat avatar, nama, tombol star, menu "…", dan tombol share/copy link.
  9. ⚠ Setiap project memiliki default landing view dari template; pengguna dapat menimpanya dengan landing view personal.
- **Team-managed vs company-managed:** Butir 6 hanya company-managed.
- **Catatan & ketergantungan:** Detail view ada di [10](10-board-kolom-kartu.md), [11](11-backlog-sprint.md), [16](16-laporan-insight.md), [18](18-timeline-rencana-tampilan-proyek.md), [26](26-kerja-bisnis-formulir-persetujuan.md). Definisi tab view: PLAN-001 di [18](18-timeline-rencana-tampilan-proyek.md) (tab view project software) dan BIZ-022 di [26](26-kerja-bisnis-formulir-persetujuan.md) (tab view project business); modul ini memiliki kerangka navigasi project (sidebar, header, Features toggle), sedangkan aturan tab mengikuti PLAN-001/BIZ-022 bila ada perbedaan. Routing SPA di `frontend/` mengikuti pola `/projects/{key}/{view}` dan `/spaces/{key}/{view}`.

### PRJ-028 Project Summary

- **Deskripsi:** Halaman (view) Summary project yang menampilkan statistik dan aktivitas terbaru. Kepemilikan dibagi tiga: modul ini memiliki entri halaman (rute, tab view, toggle Features, ketersediaan per tipe project), RPT-034 di [16](16-laporan-insight.md) memiliki metrik dan perilaku widget, dan PLAN-025 di [18](18-timeline-rencana-tampilan-proyek.md) hanya menjadi rujukan ke RPT-034.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan akses Browse ke project.
- **Aturan perilaku (acceptance criteria):**
  1. Kartu ringkas menampilkan jumlah work item Completed, Updated, Created, dan Due soon dalam 7 hari, dengan definisi perhitungan mengikuti RPT-034 butir 1.
  2. Widget Status overview (donut), Priority breakdown, Types of work, Team workload, Recent activity, dan Epic progress ditampilkan dengan perilaku sesuai RPT-034 butir 2-7.
  3. Setiap kartu/segmen menautkan ke daftar work item terkait (navigator/List dengan JQL yang sesuai), sesuai RPT-034 butir 8.
  4. Summary tersedia untuk project business dan software, team-managed dan company-managed.
  5. Data dihitung mengikuti permission pengguna; work item dengan issue security yang tidak dapat dilihat tidak dihitung (RPT-034 butir 12).
  6. Summary dapat diaktifkan/dinonaktifkan lewat Features (PRJ-049, PRJ-050) atau pengaturan view project (PRJ-027).
  7. ⚠ Tersedia filter periode (mis. 7/14/30 hari) untuk kartu ringkas; bila diadopsi, definisi periode ditambahkan ke RPT-034.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik metrik & widget: RPT-034 ([16](16-laporan-insight.md)). PLAN-025 ([18](18-timeline-rencana-tampilan-proyek.md)) adalah duplikat rincian widget yang harus dibaca sebagai rujukan ke RPT-034. Bila ada perbedaan, RPT-034 berlaku untuk widget dan PRJ-028 berlaku untuk entri halaman. Modul 16 dan 18 perlu menyelaraskan rujukan silangnya ke keputusan ini.

### PRJ-029 Project Shortcuts

- **Deskripsi:** Tautan cepat di navigasi project, mis. ke dokumentasi, repository, atau tool lain.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengelola; semua pengguna berakses melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Add shortcut" membuka form dengan field Web address dan Name (keduanya wajib).
  2. Shortcut tampil di section Shortcuts pada sidebar/tab project sesuai urutan dibuat.
  3. Shortcut dapat diedit dan dihapus.
  4. Aksi "Add repository" membuat shortcut repository dengan ikon khusus.
  5. Hanya Admin Proyek yang dapat menambah/mengedit/menghapus; semua pengguna dengan akses project dapat melihat.
  6. Shortcut dibuka di tab browser baru (`target="_blank"` dengan `rel="noopener noreferrer"`).
  7. Web address harus URL valid `http`/`https`; skema lain (mis. `javascript:`) ditolak.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik aturan shortcut di project business: BIZ-024 di [26](26-kerja-bisnis-formulir-persetujuan.md); fitur ini berlaku untuk project software dan sebagai perilaku dasar bersama. Dapat dinonaktifkan via Features team-managed (PRJ-049).

### PRJ-030 Project Pages

- **Deskripsi:** Tab Pages di project untuk dokumentasi yang terhubung ke ruang dokumentasi (space) atau halaman eksternal.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek menghubungkan; pengguna membaca sesuai izin sistem dokumentasi.
- **Aturan perilaku (acceptance criteria):**
  1. Tab Pages menampilkan page tree dari ruang dokumentasi yang terhubung.
  2. Admin Proyek dapat menghubungkan ke ruang dokumentasi yang ada atau membuat yang baru.
  3. Pengguna dapat membuat halaman baru (dengan template) dari dalam project.
  4. Bila site belum memiliki sistem dokumentasi terhubung, tab menampilkan ajakan menghubungkan penyedia dokumentasi atau menautkan halaman eksternal.
  5. Admin Proyek dapat memutus atau mengganti ruang dokumentasi yang terhubung.
  6. Permission halaman tetap mengikuti sistem dokumentasi sumber; produk ini tidak menampilkan konten yang tidak dapat dibuka pengguna.
  7. Tab Pages dapat diaktifkan/dinonaktifkan lewat Features.
  8. Company-managed legacy: menu "Project pages" menautkan ruang/halaman dokumentasi sebagai tautan sederhana.
- **Team-managed vs company-managed:** Butir 8 hanya company-managed.
- **Catatan & ketergantungan:** Referensi Jira memakai Confluence (layanan lintas produk Atlassian, di luar cakupan); clone menyediakan antarmuka konektor generik di [25](25-integrasi-impor-ekspor-ekstensi.md). Halaman bawaan business project di [26](26-kerja-bisnis-formulir-persetujuan.md).

## Siklus hidup project

### PRJ-031 Project Lifecycle States & Events

- **Deskripsi:** Model state lifecycle project (active, archived, trashed, deleted) beserta transisi dan event yang dipicu.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Sistem; Admin Situs memicu transisi.
- **Aturan perilaku (acceptance criteria):**
  1. Transisi yang diizinkan: `active ↔ archived`; `active → trashed → active` (restore) atau `trashed → deleted`; transisi lain ditolak dengan galat 409.
  2. ⚠ Transisi langsung `archived → trashed` diizinkan hanya bila pemilik produk memutuskan demikian; default: project harus di-restore dulu.
  3. Setiap state memiliki efek terdefinisi pada indeks pencarian (archived/trashed dikeluarkan), pemeriksaan izin (trashed = tidak dapat diakses, archived = read-only), notifikasi dan automation (berhenti untuk archived/trashed), serta webhook.
  4. Job terjadwal harian menghapus permanen project yang berada di trash lebih dari 60 hari.
  5. Event webhook dipancarkan: `project_created`, `project_updated`, `project_deleted`, `project_soft_deleted`, `project_restored_deleted`, `project_archived`, `project_restored_archived`.
  6. Setiap perubahan state dan perubahan detail project dicatat di audit log (pelaku, waktu, nilai lama/baru).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Webhook di [23](23-api-publik-webhook.md); katalog event di [28](28-model-data-inti.md); audit log di [22](22-administrasi-sistem.md).

### PRJ-032 Project Archive

- **Deskripsi:** Mengarsipkan project yang tidak aktif agar read-only dan tersembunyi dari pencarian, tanpa menghapus data.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Fitur hanya tersedia di plan Premium dan Enterprise; di plan lain aksi Archive tidak tampil dan API mengembalikan galat plan.
  2. ⚠ Archive dilakukan Admin Situs dari directory, Settings > Projects, atau Details; Administrator project team-managed tidak dapat mengarsipkan kecuali pemilik produk memutuskan sebaliknya.
  3. Work item project archived bersifat read-only dan tidak muncul di search/JQL, board, dashboard, atau filter.
  4. ⚠ Work item project archived masih dapat dibuka lewat direct link oleh pengguna berizin, dengan banner penanda "Archived".
  5. Tidak dapat membuat work item baru maupun melakukan transisi di project archived; upaya lewat API ditolak 409.
  6. Key project archived tetap dicadangkan.
  7. Daftar project archived tersedia di Settings > Projects > Archived.
  8. Restore mengembalikan project ke status aktif penuh (indeks, board, notifikasi kembali normal).
  9. Notifikasi dan automation untuk project archived berhenti dijalankan.
  10. Archive mendukung project team-managed dan company-managed.
  11. REST API: `POST /project/{key}/archive` dan `POST /project/{key}/restore`.
- **Team-managed vs company-managed:** Tidak ada perbedaan (butir 10).
- **Catatan & ketergantungan:** Archive per work item di [06](06-work-item-inti.md).

### PRJ-033 Project Trash, Restore & Permanent Delete

- **Deskripsi:** Project yang dihapus masuk trash selama 60 hari dan dapat di-restore dalam periode itu; setelahnya, atau bila dihapus manual dari trash, project terhapus permanen.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Administrator project team-managed untuk project sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Move to trash" tersedia di directory atau Details dan memerlukan dialog konfirmasi.
  2. Project di trash disimpan 60 hari, lalu dihapus permanen otomatis oleh job terjadwal.
  3. Selama di trash, work item tidak dapat diakses atau dicari; board/filter terkait tidak menampilkan datanya.
  4. Key project tetap dicadangkan selama di trash.
  5. Halaman Trash di Settings > Projects menampilkan tanggal penghapusan, pelaku, dan sisa hari sebelum hapus permanen.
  6. Restore mengembalikan project beserta work item, konfigurasi, dan board ke status active.
  7. "Delete permanently" dari trash sebelum 60 hari memerlukan konfirmasi (mengetik key project) dan tidak dapat dibatalkan.
  8. Delete permanen menghapus work item, lampiran, komentar, worklog, versions, components, dan issue link.
  9. Team-managed: konfigurasi project-scoped (work type, field, workflow, board) ikut terhapus permanen.
  10. Company-managed: scheme tidak dihapus (menjadi tidak terpakai); board yang filternya hanya project itu menjadi kosong.
  11. ⚠ Izin: Admin Situs dapat trash/restore/delete semua project; Administrator project team-managed dapat memindahkan project sendiri ke trash, sedangkan restore dibatasi ke Admin Situs kecuali pemilik produk memutuskan lain.
  12. Fitur trash berlaku di semua plan.
  13. REST API: `DELETE /project/{key}?enableUndo=true` (ke trash) atau `false` (permanen); `POST /project/{key}/restore`; penghapusan asinkron lewat `POST /project/{key}/delete` yang mengembalikan task ID.
  14. ⚠ Setelah delete permanen, key dan alias project dilepas dan dapat dipakai lagi oleh project baru.
  15. Setiap aksi trash, restore, dan delete permanen tercatat di audit log.
- **Team-managed vs company-managed:** Lihat butir 9-11.
- **Catatan & ketergantungan:** Kebijakan trash terpadu dimiliki ADM-025 di [22](22-administrasi-sistem.md) (Fase 3). Karena fitur ini berada di Fase 2, job purge terjadwal dan trash project (ADM-025 butir 1 dan butir 6) dibangun di Fase 2 sebagai bagian awal ADM-025; Fase 3 memperluas mekanisme yang sama ke archive, trash custom field, laporan storage, dan riwayat purge. Modul 22 perlu mencatat pemecahan fase ini pada ADM-025. Penghapusan permanen harus berupa background job bertahap (batch) untuk project besar ([27](27-nfr-arsitektur-platform.md)).

## Project settings, anggota & akses

### PRJ-034 Project Settings Hub (company-managed)

- **Deskripsi:** Navigasi Project settings company-managed yang mengelompokkan semua konfigurasi project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pemegang Administer Projects; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Menu settings memuat: Summary, Details, People, Automation, Notifications, Permissions, Issue security, Work types/Issue types (dengan Issue layout), Workflows, Screens, Fields, Priorities, Versions/Releases, Components, Features, Toolchain, Development tools, Issue collectors, dan Apps; project business menambahkan Forms dan Approvals.
  2. Halaman Summary settings meringkas scheme yang dipakai (Work types, Workflows, Screens, Fields, Permissions, Notifications, Issue security, Priorities) dengan tautan ke halaman masing-masing.
  3. Tautan untuk mengedit scheme global hanya tampil bagi Admin Situs.
  4. Admin Proyek melihat item konfigurasi yang hanya dapat diubah Admin Situs dalam mode read-only (tanpa tombol aksi).
  5. Menu yang tampil bergantung pada tipe project dan app yang terpasang; menu fitur yang tidak berlaku tidak dirender.
  6. Tombol "Back to project" kembali ke landing view project.
- **Team-managed vs company-managed:** Versi team-managed di PRJ-035.
- **Catatan & ketergantungan:** Fase 1 cukup Summary, Details, People, Features, dan halaman scheme read-only; halaman lain menyusul sesuai fase fiturnya.

### PRJ-035 Project Settings Hub (team-managed)

- **Deskripsi:** Navigasi Project settings team-managed yang lebih sederhana dan sepenuhnya dikelola Administrator project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Role Administrator project team-managed; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Menu settings memuat: Details, Access, Notifications, Automation, Features, Work types (sub-menu per type berisi fields dan workflow), Apps, Toolchain/Development tools, serta Forms/Approvals untuk project business.
  2. Semua item dapat diubah role Administrator project tanpa keterlibatan Admin Situs.
  3. Daftar work type tampil sebagai sub-menu dengan tombol tambah type.
  4. Tombol "Back to project" kembali ke landing view project.
- **Team-managed vs company-managed:** Versi company-managed di PRJ-034.
- **Catatan & ketergantungan:** Lihat PRJ-051, PRJ-052.

### PRJ-036 Project Administration Rights & Extended Project Administration

- **Deskripsi:** Batasan apa yang boleh dilakukan Admin Proyek dibanding Admin Situs di project company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects / Extended project administration), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Permission "Administer Projects" memberi akses ke Details, People, Versions, Components, Issue layout, Automation, Features, Shortcuts, dan Issue collectors.
  2. Hanya Admin Situs yang dapat mengganti scheme (permission, notification, workflow, issue type, screen, field configuration, priority, issue security).
  3. Permission "Extended project administration" (diberikan via permission scheme) mengizinkan Admin Proyek mengedit workflow dan screen yang hanya dipakai project-nya.
  4. ⚠ Batasan Extended project administration: tidak dapat membuat status baru (hanya memakai status yang ada), tidak dapat mengedit condition/validator/post function, dan tidak berlaku bila workflow/screen dibagi dengan project lain.
  5. Admin Proyek tidak dapat membuat custom field global.
  6. Admin Situs selalu dapat mengelola semua project, termasuk project Private team-managed.
  7. Team-managed: role Administrator project dapat mengubah semua konfigurasi project-scoped.
- **Team-managed vs company-managed:** Lihat butir 7.
- **Catatan & ketergantungan:** Definisi permission di [03](03-izin-peran-keamanan.md); editor workflow di [09](09-workflow-status-transisi.md); screen di [05](05-konfigurasi-work-item-skema.md).

### PRJ-037 People (company-managed)

- **Deskripsi:** Halaman untuk mengelola keanggotaan pengguna dan grup ke project role di project company-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pemegang Administer Projects; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel anggota menampilkan pengguna/grup dengan kolom Name, Email, dan Role.
  2. Tombol "Add people" memungkinkan memilih satu atau lebih pengguna/grup dan satu atau lebih role sekaligus.
  3. Role anggota dapat diedit inline (multi-select) dan anggota dapat dihapus dari project.
  4. Daftar dapat difilter berdasarkan role dan dicari berdasarkan nama.
  5. ⚠ Role default yang tersedia: Administrator, Member, Viewer (ditambah role sistem untuk akses app); site migrasi dapat memiliki Administrators/Developers/Users.
  6. Membership role dipakai oleh permission scheme, notification scheme, issue security scheme, dan condition workflow; perubahan membership berlaku segera pada pemeriksaan izin.
  7. Sistem menampilkan peringatan bila pengguna yang ditambahkan tidak memiliki product access.
  8. Admin dapat mengundang pengguna baru dari dialog Add people bila kebijakan site mengizinkan undangan.
  9. ⚠ Plan Free: role dan permission tidak dapat disunting; halaman People tetap tersedia dalam mode read-only yang menampilkan anggota default.
  10. ⚠ Pengguna tanpa Browse Projects dapat mengirim "Request access"; admin menyetujui dengan menambahkan pengguna ke role pilihan atau menolak.
- **Team-managed vs company-managed:** Versi team-managed di PRJ-038.
- **Catatan & ketergantungan:** Role global di PRJ-040; undangan di [02](02-pengguna-grup-organisasi.md).

### PRJ-038 Access (team-managed)

- **Deskripsi:** Halaman pengelolaan anggota, role, dan access level pada project team-managed, termasuk alur permintaan akses.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Role Administrator project team-managed; Admin Situs; pengguna pemohon akses.
- **Aturan perilaku (acceptance criteria):**
  1. Administrator memilih access level Open/Limited/Private (PRJ-039).
  2. Administrator menambah pengguna atau grup dengan role Administrator, Member, Viewer, atau custom role.
  3. Role anggota dapat diubah dan anggota dapat dihapus.
  4. Custom role dapat dibuat dengan kumpulan permission terpilih; tersedia di plan Standard dan Premium/Enterprise, tidak di Free.
  5. Halaman menampilkan daftar permission per role (read-only untuk role bawaan).
  6. Project harus memiliki minimal satu Administrator; menghapus atau menurunkan Administrator terakhir ditolak dengan pesan "Project harus memiliki minimal satu Administrator".
  7. Administrator dapat mengundang pengguna baru via email bila kebijakan site mengizinkan.
  8. Plan Free: role/permission tidak dapat disunting.
  9. ⚠ Pengguna non-member yang membuka tautan project Private melihat halaman "You don't have access" dengan tombol "Request access" dan kolom pesan opsional.
  10. Permintaan akses dikirim ke Administrator project (email dan in-app); status "Request pending" mencegah permintaan ganda dari pengguna yang sama.
  11. ⚠ Administrator menyetujui (pengguna ditambahkan sebagai Member atau role yang dipilih) atau menolak permintaan; pemohon menerima notifikasi hasilnya.
  12. Bila pemohon belum memiliki product access, alur beralih ke product access request di level site/organisasi.
  13. ⚠ Penambahan anggota melalui persetujuan permintaan akses tercatat di audit log.
- **Team-managed vs company-managed:** Versi company-managed di PRJ-037.
- **Catatan & ketergantungan:** Berbeda dari access request tingkat produk di [02](02-pengguna-grup-organisasi.md), yang hanya mengatur product access. Custom role di [03](03-izin-peran-keamanan.md); notifikasi di [19](19-notifikasi-email.md).

### PRJ-039 Project Access Levels (team-managed)

- **Deskripsi:** Pengaturan visibilitas dan hak akses umum untuk project team-managed, tanpa permission scheme. Modul ini memiliki sisi UI (pilihan di create wizard, halaman Access, indikator gembok), sedangkan semantik dan enforcement access level dimiliki PERM-038 di [03](03-izin-peran-keamanan.md).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Role Administrator project team-managed; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Open: semua pengguna berlisensi di site dapat melihat, membuat, mengedit, dan berkomentar pada work item project (PERM-038 butir 1).
  2. Limited: semua pengguna berlisensi dapat melihat dan berkomentar; hanya anggota yang dapat membuat dan mengedit (PERM-038 butir 2).
  3. Private: hanya anggota yang ditambahkan yang dapat melihat project; Admin Situs tetap dapat mengelola project (PERM-038 butir 3 dan 6).
  4. Access level dipilih di create wizard dan dapat diubah di Project settings > Access.
  5. Mengubah ke Private langsung mencabut akses non-member; filter, board, dan dashboard milik non-member tidak lagi menampilkan work item project tersebut (PERM-038 butir 3).
  6. Directory dan settings menampilkan ikon gembok/indikator access level.
  7. ⚠ Admin Situs melihat project Private di Settings > Projects dan dapat mengelola aksesnya; Admin Situs tidak dapat melihat isi work item tanpa menambahkan dirinya sebagai anggota (aksi tercatat di audit log); selaras PERM-038 butir 6.
  8. ⚠ Plan Free: permission dan role project tidak dapat disunting; opsi Private tidak tersedia di Free kecuali pemilik produk memutuskan sebaliknya (PERM-038 butir 8).
  9. Akses anonim/publik tidak tersedia untuk project team-managed (PERM-038 butir 5).
- **Team-managed vs company-managed:** Company-managed mengatur visibilitas lewat permission Browse Projects di permission scheme (PRJ-041).
- **Catatan & ketergantungan:** Pemilik semantik & enforcement: PERM-038 di [03](03-izin-peran-keamanan.md) (Authorizer PERM-002). Bila ada perbedaan, PERM-038 berlaku. Pencarian harus menerapkan filter akses yang sama ([13](13-pencarian-filter-ekspor.md)).

### PRJ-040 Project Roles Administration (global)

- **Deskripsi:** Definisi project role global oleh Admin Situs, yang dipakai project company-managed untuk membership dan scheme. Perilaku lengkap dimiliki PERM-018 Project Roles (Global Definition) di [03](03-izin-peran-keamanan.md); fitur ini merangkum keterkaitannya dengan project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs mengelola role di Settings > System > Project roles: Add (name, description), Edit, Delete; nama role unik secara case-insensitive (PERM-018 butir 1-2).
  2. Sebelum menghapus role, sistem menampilkan pemakaian role di permission, notification, dan issue security scheme serta workflow; penghapusan memerlukan konfirmasi (PERM-018 butir 3-4).
  3. Role tersedia di semua project company-managed; membership diatur per project di People (PRJ-037). Fase 1 minimal memuat role bawaan (PERM-018 butir 6) agar People dapat dipakai.
  4. (Rujukan, tidak dibangun) "Default members" per role, yaitu anggota yang otomatis ditambahkan ke role pada project baru, adalah perilaku Data Center (PERM-067). Project company-managed baru tidak mewarisi anggota dari definisi role; anggota awal ditentukan wizard create project (PRJ-010) dan template.
  5. JQL function `projectsWhereUserHasRole("role")` mengembalikan project tempat pengguna saat ini memiliki role tersebut (PRJ-065 butir 5).
  6. REST API `/role` (CRUD role global) dan `/project/{key}/role/{id}` (membership per project) mengikuti PERM-061, dengan fase sendiri (P2 / Fase 3); UI Fase 1 memakai endpoint internal SPA.
- **Team-managed vs company-managed:** Role team-managed (Administrator/Member/Viewer/custom) terpisah dan diatur di Access (PRJ-038).
- **Catatan & ketergantungan:** Pemilik definisi role: PERM-018. REST: PERM-061. Default members: PERM-067 (di luar cakupan). Fase diselaraskan ke Fase 1 karena People (PRJ-037, Fase 1) memerlukan role. Permission scheme di [03](03-izin-peran-keamanan.md).

### PRJ-041 Permissions (project settings)

- **Deskripsi:** Entri Project settings > Permissions yang menampilkan permission scheme project company-managed beserta grant-nya. Aturan halaman dimiliki PERM-017 di [03](03-izin-peran-keamanan.md).
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (ubah); Admin Proyek (lihat).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan nama scheme dan daftar permission (Browse Projects, Create, Edit, Assign, Transition, Delete, Administer Projects, dll.) beserta penerima tiap grant (PERM-017 butir 1).
  2. Admin Situs memiliki aksi "Use a different scheme" dan "Edit permissions" (PERM-017 butir 2-3).
  3. Admin Proyek melihat halaman dalam mode read-only (PERM-017 butir 1).
  4. Halaman menautkan ke Permission helper (PERM-055) dengan konteks project terisi, untuk memeriksa mengapa pengguna tertentu memiliki atau tidak memiliki izin pada work item/project (PERM-017 butir 5).
  5. ⚠ Grant Browse Projects ke "Anyone on the web" membuat project publik hanya bila akses publik diizinkan di level site; bila tidak diizinkan, opsi tersebut tidak tersedia.
  6. Project team-managed tidak memiliki halaman ini (memakai Access/roles, PRJ-038).
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pemilik aturan halaman: PERM-017 di [03](03-izin-peran-keamanan.md); bila ada perbedaan, PERM-017 berlaku. Permission helper di PERM-055. Fase 1 memakai scheme default sistem tanpa halaman ini.

### PRJ-042 Issue Security (project settings)

- **Deskripsi:** Pengaturan issue security scheme untuk membatasi visibilitas work item tertentu di project company-managed.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan scheme yang dipakai serta security level dan anggotanya.
  2. Admin Situs dapat memilih scheme lain melalui wizard yang memetakan level lama ke level baru untuk work item yang sudah memiliki level.
  3. Default security level scheme ditampilkan dan diterapkan ke work item baru.
  4. Fitur tidak tersedia di plan Free (mulai Standard).
  5. ⚠ Project team-managed tidak mendukung security level kecuali pemilik produk memutuskan menambahkannya.
- **Team-managed vs company-managed:** Lihat butir 5.
- **Catatan & ketergantungan:** Detail issue security di [03](03-izin-peran-keamanan.md).

## Konfigurasi work item project (company-managed)

### PRJ-043 Issue Types (company-managed project settings)

- **Deskripsi:** Halaman yang menampilkan issue type scheme project dan konfigurasi per work type.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (ubah scheme); Admin Proyek (lihat, akses Issue layout).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan daftar work type dalam scheme beserta default type.
  2. Admin Situs dapat mengganti issue type scheme melalui wizard migrasi yang mewajibkan pemetaan work item yang type-nya tidak ada di scheme baru sebelum perubahan diterapkan.
  3. Per work type ditampilkan tautan ke workflow, screen, dan field configuration yang berlaku.
  4. Per work type tersedia akses ke Issue layout (PRJ-048).
  5. Hierarki work type (Epic, standar, subtask) mengikuti konfigurasi global dan ditampilkan sebagai label level.
- **Team-managed vs company-managed:** Versi team-managed di PRJ-051.
- **Catatan & ketergantungan:** Issue type scheme di [05](05-konfigurasi-work-item-skema.md).

### PRJ-044 Workflows (company-managed project settings)

- **Deskripsi:** Halaman yang menampilkan workflow scheme dan pemetaan workflow ke work type.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek dengan Extended project administration.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan daftar workflow beserta work type yang memakainya, dengan preview diagram atau teks (status dan transisi).
  2. Admin Situs memiliki aksi switch scheme, edit workflow, dan add workflow.
  3. Mengganti workflow scheme memicu wizard migrasi status yang memetakan status lama ke status workflow baru untuk semua work item terdampak.
  4. Admin Proyek dengan Extended project administration dapat mengedit workflow yang tidak dibagi (PRJ-036).
  5. Workflow yang dibagi dengan project lain diberi indikator "Shared with N projects".
- **Team-managed vs company-managed:** Versi team-managed di PRJ-052.
- **Catatan & ketergantungan:** Editor dan migrasi status di [09](09-workflow-status-transisi.md).

### PRJ-045 Screens (project settings)

- **Deskripsi:** Halaman yang menampilkan issue type screen scheme dan screen untuk tiap operasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek dengan Extended project administration.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan screen scheme per work type dan screen untuk operasi Create/Edit/View.
  2. Admin Situs dapat mengganti issue type screen scheme.
  3. Admin Proyek dengan Extended project administration dapat menambah/menghapus field pada screen yang tidak dibagi.
  4. Screen yang dibagi dengan project lain diberi indikator "Shared with N projects".
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pemilik aturan: CFG-046 (Screens), CFG-047 (Screen schemes), dan CFG-048 (Issue type screen schemes) di [05](05-konfigurasi-work-item-skema.md); halaman ini hanya entri project settings dan bila ada perbedaan, aturan modul 05 berlaku.

### PRJ-046 Fields (project settings)

- **Deskripsi:** Halaman yang menampilkan field configuration scheme project company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (ubah); Admin Proyek (lihat).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan field configuration yang berlaku per work type.
  2. Per field ditampilkan status required, hidden, renderer, dan screen tempat field dipakai.
  3. Admin Situs memiliki aksi ganti scheme dan edit field configuration.
  4. Admin Proyek melihat halaman dalam mode read-only.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed mengatur field di PRJ-051.
- **Catatan & ketergantungan:** Field configuration di [05](05-konfigurasi-work-item-skema.md).

### PRJ-047 Priorities (project settings)

- **Deskripsi:** Halaman untuk memilih priority scheme project company-managed.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan priority scheme beserta daftar prioritas dan prioritas default.
  2. Admin Situs dapat mengganti scheme; work item dengan prioritas yang tidak tersedia di scheme baru wajib dipetakan ke prioritas pengganti sebelum perubahan diterapkan.
  3. Project team-managed memakai priority scheme default global dan tidak menampilkan halaman ini.
- **Team-managed vs company-managed:** Lihat butir 3.
- **Catatan & ketergantungan:** Pemilik aturan: CFG-053 (Priorities) dan CFG-054 (Priority schemes) di [05](05-konfigurasi-work-item-skema.md). Perbedaan prioritas disengaja: CFG-053 (P0/Fase 1) adalah daftar prioritas global yang wajib ada sejak MVP, sedangkan fitur ini (P2/Fase 3) adalah pemilihan priority scheme per project yang baru diperlukan setelah CFG-054 tersedia; sebelum itu semua project memakai scheme default global.

### PRJ-048 Issue Layout (company-managed)

- **Deskripsi:** Entri Project settings > Issue layout untuk mengatur tata letak field tampilan work item per work type di project company-managed. Aturan lengkap dimiliki CFG-049 di [05](05-konfigurasi-work-item-skema.md).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pemegang Administer Projects; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih work type (dan screen terkait) yang layout-nya diatur.
  2. Layout terdiri dari section Description fields, Context fields (panel Details), Hidden when empty, dan Hidden fields.
  3. Field dapat di-drag-and-drop antar section dan diurutkan.
  4. Admin dapat menambah field yang sudah ada, dan field yang ditambahkan otomatis masuk ke screen terkait. ⚠ Bila screen dipakai bersama project lain, penambahan hanya diizinkan untuk Admin Situs (CFG-049 butir 10).
  5. Perubahan berlaku ke semua project yang berbagi screen yang sama; editor menampilkan peringatan berisi jumlah project terdampak sebelum disimpan.
  6. ⚠ Admin dapat mengatur tampil/sembunyi panel bawaan (Development, Releases, Automation) di tampilan work item.
  7. Tombol "Save changes" menyimpan dan "Discard" membatalkan semua perubahan yang belum disimpan.
- **Team-managed vs company-managed:** Team-managed memakai field editor di PRJ-051.
- **Catatan & ketergantungan:** Pemilik aturan: CFG-049 di [05](05-konfigurasi-work-item-skema.md); bila ada perbedaan, CFG-049 berlaku. Tampilan work item di [06](06-work-item-inti.md).

## Features & konfigurasi team-managed

### PRJ-049 Features (team-managed)

- **Deskripsi:** Halaman toggle fitur project team-managed yang mengontrol perilaku board, perencanaan, dan integrasi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Role Administrator project team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle Backlog: bila aktif, project memiliki backlog terpisah dari board.
  2. Toggle Sprints: bila aktif, project berperilaku Scrum (sprint); bila nonaktif, berperilaku Kanban.
  3. Toggle Estimation: pilihan Story points atau Time/original estimate sebagai satuan estimasi.
  4. Toggle Timeline (Roadmap) beserta pengaturan child level dan dependencies.
  5. Toggle Reports.
  6. Toggle Releases.
  7. Toggle Code dan Deployments.
  8. Toggle Feature flags.
  9. Toggle Security.
  10. Toggle On-call.
  11. ⚠ Toggle Pages, Goals, Components, dan Shortcuts tersedia sebagai toggle tersendiri.
  12. Project business team-managed memiliki toggle List, Calendar, Timeline, Board, Approvals, Forms, Summary, dan Attachments.
  13. Setiap toggle menampilkan deskripsi singkat dan ilustrasi.
  14. ⚠ Menonaktifkan Sprints saat masih ada sprint aktif ditolak dengan pesan "Selesaikan sprint aktif terlebih dahulu".
- **Team-managed vs company-managed:** Versi company-managed di PRJ-050.
- **Catatan & ketergantungan:** Menonaktifkan toggle hanya menyembunyikan view/fitur, tidak menghapus data. API `GET/PUT /project/{key}/features/{featureKey}` (PRJ-064). Board/sprint di [10](10-board-kolom-kartu.md), [11](11-backlog-sprint.md); estimasi di [14](14-pelacakan-waktu-estimasi.md).

### PRJ-050 Features (company-managed)

- **Deskripsi:** Halaman untuk mengaktifkan/menonaktifkan fitur tertentu di project company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pemegang Administer Projects.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Grup Planning: toggle List, Calendar, Summary, Timeline, dan Goals.
  2. Grup Development: toggle Code, Deployments, Releases, Feature flags, dan Security.
  3. Grup Operations: toggle On-call.
  4. Grup Knowledge: toggle Pages.
  5. Toggle menyembunyikan view dari navigasi tanpa menghapus data; mengaktifkan kembali memulihkan view dengan data utuh.
  6. Fitur board (backlog Kanban, estimation, sprint) tidak diatur di halaman ini melainkan di board settings; halaman menampilkan tautan ke board settings.
- **Team-managed vs company-managed:** Versi team-managed di PRJ-049.
- **Catatan & ketergantungan:** Board settings di [10](10-board-kolom-kartu.md).

### PRJ-051 Team-managed Issue Types & Fields

- **Deskripsi:** Pengelolaan work type dan field per work type milik project team-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Role Administrator project team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan daftar work type; Administrator dapat menambah type baru (nama, deskripsi, ikon).
  2. Hierarki work type: Epic, type standar, Subtask; type baru ditempatkan pada level standar atau subtask.
  3. Nama/ikon dapat diedit; menghapus type yang masih dipakai mewajibkan pemindahan work item ke type lain terlebih dahulu.
  4. Field editor drag-and-drop per work type dengan section Description fields, Context fields, dan Hidden fields (termasuk opsi Hide when empty).
  5. Tipe field yang tersedia: Short text, Paragraph, Date, Number, Time stamp, Dropdown, Checkbox, Labels, People, Start date, Story point estimate, dan lainnya sesuai katalog di modul 05.
  6. Field yang dibuat dapat dipakai ulang di work type lain dalam project yang sama, tetapi tidak di project lain.
  7. ⚠ Jumlah field per work type dan jumlah work type per project dibatasi (usulan 50 field per type dan 30 type per project); melewati batas ditolak dengan pesan yang menyebut batasnya.
  8. ⚠ Field dapat ditandai required dan diberi default value.
  9. ⚠ Di JQL, custom field team-managed ditampilkan terpisah per project (mis. "Nama field (Nama Project)") agar tidak bentrok dengan field lain bernama sama.
- **Team-managed vs company-managed:** Versi company-managed di PRJ-043, PRJ-046, PRJ-048.
- **Catatan & ketergantungan:** Katalog tipe field di [05](05-konfigurasi-work-item-skema.md) (CFG-041); JQL di [12](12-jql-bahasa-kueri.md). Pembuatan field lokal dapat dikunci Space policy "Global fields only" (PRJ-069).

### PRJ-052 Team-managed Workflows per Issue Type

- **Deskripsi:** Editor workflow sederhana per work type di project team-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Role Administrator project team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Workflow dibuka dari halaman work type (tombol "Edit workflow").
  2. Administrator dapat menambah/menghapus status dengan kategori To do / In progress / Done; status yang masih dipakai work item wajib dipetakan ke status lain sebelum dihapus.
  3. Administrator dapat membuat transisi antar status dan mengaktifkan opsi "Allow all statuses to transition to this one".
  4. ⚠ Rules transisi yang tersedia: restrict who can move, validate field, auto assign, dan update field.
  5. Status hanya berlaku di project itu; perubahan status tersambung ke kolom board (status baru otomatis mendapat kolom atau dipetakan ke kolom yang ada).
  6. ⚠ Satu workflow dapat dibagi ke beberapa work type dalam project yang sama.
- **Team-managed vs company-managed:** Versi company-managed di PRJ-044.
- **Catatan & ketergantungan:** Engine workflow di [09](09-workflow-status-transisi.md); kolom board di [10](10-board-kolom-kartu.md). Workflow default Fase 1 (To Do/In Progress/Done) di-seed tanpa editor. Pembuatan status lokal dapat dikunci Space policy "Global statuses only" (PRJ-069, WF-017).

## Notifikasi, rilis, komponen & otomasi project

### PRJ-053 Notifications (project settings)

- **Deskripsi:** Pengaturan notifikasi email project: notification scheme di company-managed dan pengaturan per event di team-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: halaman menampilkan notification scheme yang dipakai beserta event dan penerimanya (NTF-005 butir 1).
  2. Company-managed: menu Actions menyediakan "Use a different scheme" dan edit scheme, keduanya khusus Admin Situs (NTF-005 butir 2-4).
  3. Team-managed: halaman menampilkan daftar event (Work item created, updated, assigned, commented, transitioned, deleted, dll.) beserta penerimanya (Assignee, Reporter, Watchers, All members, role, pengguna/grup), sesuai NTF-007 butir 1-3.
  4. Team-managed: Administrator dapat menambah/menghapus penerima per event (NTF-007 butir 3).
  5. Halaman menyediakan pengaturan alamat pengirim email project (PRJ-054, NTF-025).
  6. Preferensi personal pengguna tetap berlaku di atas pengaturan project (mis. tidak menerima notifikasi untuk perubahan sendiri).
- **Team-managed vs company-managed:** Lihat butir 1-4.
- **Catatan & ketergantungan:** Pemilik aturan: NTF-004 (Notification Scheme company-managed), NTF-005 (asosiasi scheme company-managed), dan NTF-007 (notifikasi team-managed) di [19](19-notifikasi-email.md); bila ada perbedaan, aturan modul 19 berlaku. Event dan template email juga di [19](19-notifikasi-email.md).

### PRJ-054 Project Email Sender

- **Deskripsi:** Entri pengaturan alamat "From" untuk email notifikasi yang dikirim dari project. Aturan lengkap dimiliki NTF-025 Project Email Address (Custom Sender) di [19](19-notifikasi-email.md).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; ⚠ Admin Proyek (Administer Projects) terbatas pada domain terverifikasi, sesuai NTF-025.
- **Aturan perilaku (acceptance criteria):**
  1. Default pengirim adalah from address site (NTF-024), mis. `noreply@<domain-instance>` yang dikonfigurasi lewat environment (NTF-025 butir 3).
  2. Alamat diubah di Project settings > Notifications > Actions > Edit project email address (NTF-025 butir 1).
  3. Alamat harus berasal dari domain terverifikasi (NTF-026, dengan SPF/DKIM) atau domain default instalasi; alamat lain ditolak dengan pesan galat (NTF-025 butir 2).
  4. Format nama pengirim (mis. nama pengguna pemicu "via <produk>") mengikuti From header format site (NTF-025 butir 4).
  5. ⚠ Pengaturan pengirim juga tersedia di project team-managed (NTF-025 butir 5).
  6. Alamat email divalidasi formatnya (RFC 5322 sederhana); format salah ditolak.
- **Team-managed vs company-managed:** Lihat butir 5.
- **Catatan & ketergantungan:** Pemilik aturan: NTF-025 di [19](19-notifikasi-email.md); bila ada perbedaan, NTF-025 berlaku. Verifikasi domain di NTF-026 dan [02](02-pengguna-grup-organisasi.md).

### PRJ-055 Versions (project settings)

- **Deskripsi:** Pengelolaan versions/releases project dari settings atau halaman Releases.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek; pengguna dengan permission pengelolaan versi.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan daftar versions dengan status Unreleased/Released/Archived, start date, release date, dan description.
  2. Pengguna berizin dapat create, edit, release, unrelease, archive, merge, dan delete version; delete menawarkan opsi memindahkan fix version/affects version work item ke versi lain.
  3. Urutan versions dapat diubah dengan drag.
  4. Versions tersedia di team-managed (bila feature Releases aktif) dan company-managed.
  5. ⚠ Akses dikontrol permission Administer Projects atau permission khusus "Manage versions" sesuai keputusan modul 03.
- **Team-managed vs company-managed:** Lihat butir 4.
- **Catatan & ketergantungan:** Detail di [15](15-rilis-versi-komponen.md).

### PRJ-056 Components (project settings)

- **Deskripsi:** Pengelolaan components project untuk mengelompokkan work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Create component dengan field name (wajib, unik per project), description, component lead, dan default assignee.
  2. Pilihan default assignee: Project default, Project lead, Component lead, Unassigned.
  3. Component dapat diedit dan dihapus (dengan opsi memindahkan work item ke component lain); daftar dapat dicari/difilter.
  4. ⚠ Project team-managed juga mendukung components dengan pilihan default assignee yang sama.
  5. Company-managed dapat beralih ke katalog komponen eksternal (setara Compass components) melalui integrasi bila tersedia.
- **Team-managed vs company-managed:** Lihat butir 4-5.
- **Catatan & ketergantungan:** Detail di [15](15-rilis-versi-komponen.md). Katalog komponen eksternal adalah layanan lintas produk (di luar cakupan).

### PRJ-057 Project Automation

- **Deskripsi:** Akses ke rule automation ber-scope project dari Project settings.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (rule single-project); Admin Situs (rule global/multi-project).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan rule single-project serta rule global/multi-project yang berlaku pada project ini (yang terakhir read-only bagi Admin Proyek).
  2. Rule dapat dibuat dari template atau dari awal.
  3. Admin Proyek dapat membuat dan mengelola rule single-project.
  4. Audit log eksekusi rule ditampilkan per project.
  5. Batas eksekusi bulanan mengikuti plan; halaman menampilkan pemakaian terhadap batas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik aturan: AUTO-001 (Project automation rule list) dan AUTO-009 (Team-managed project automation) di [20](20-otomasi.md); halaman ini hanya entri project settings dan bila ada perbedaan, aturan modul 20 berlaku.

## Integrasi & ekstensi project

### PRJ-058 Toolchain

- **Deskripsi:** Entri Project settings > Toolchain yang menampilkan tool terhubung ke project per kategori. Aturan lengkap dimiliki DEV-006 Toolchain di [24](24-integrasi-devops.md).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek (mengaitkan tool ke project); Admin Situs (memasang konektor di level site). Menu hanya ada di project software (DEV-006 butir 1).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tool dikelompokkan per kategori tetap sesuai DEV-006 butir 2 (Plan, Code, Build, Deploy, Operate, ditambah Security dan Design bila penyedianya terpasang). Kategori di inventaris (Development, Deployment, Feature flags, Security, Operations/On-call, Design, Documentation) dipetakan ke kategori tersebut.
  2. Admin dapat menambah tool dari konektor yang terpasang di site; konektor yang belum terpasang memerlukan Admin Situs (DEV-006 butir 3 dan 6).
  3. Admin dapat memutus hubungan tool dari project tanpa menghapus integrasi di level site.
  4. Setiap tool menampilkan status koneksi dan ringkasan data yang disinkronkan (mis. jumlah repository, waktu sinkron terakhir).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik aturan: DEV-006 di [24](24-integrasi-devops.md); bila ada perbedaan, DEV-006 berlaku. Development Tools (PRJ-059) menjadi bagian Toolchain.

### PRJ-059 Development Tools

- **Deskripsi:** Halaman yang menampilkan tool development yang terhubung (code, branch, PR, build, deployment).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek; Admin Situs; pengguna dengan View development tools.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan daftar provider terhubung beserta status sinkronisasinya.
  2. Panel Development pada work item hanya tampil bagi pengguna dengan permission "View development tools".
  3. Halaman memuat petunjuk menyertakan issue key di pesan commit, nama branch, dan judul PR.
  4. Halaman menautkan ke konfigurasi integrasi/akun DVCS di level site.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail di [24](24-integrasi-devops.md); permission di [03](03-izin-peran-keamanan.md).

### PRJ-060 Apps (project settings)

- **Deskripsi:** Bagian settings yang menampung halaman konfigurasi ekstensi yang terpasang di project.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek; Admin Situs (instalasi di level site).
- **Aturan perilaku (acceptance criteria):**
  1. Menu Apps menampilkan halaman settings yang didaftarkan ekstensi (padanan modul `projectSettingsPage`).
  2. Halaman project yang didaftarkan ekstensi (padanan modul `projectPage`) muncul di navigasi project.
  3. Team-managed: halaman Apps memungkinkan mengaktifkan/menonaktifkan ekstensi per project.
  4. Ketersediaan bergantung pada permission pengguna dan instalasi ekstensi di level site; ekstensi yang belum dipasang tidak tampil.
- **Team-managed vs company-managed:** Lihat butir 3.
- **Catatan & ketergantungan:** Kerangka ekstensi di [25](25-integrasi-impor-ekspor-ekstensi.md); app Marketplace pihak ketiga di luar cakupan.

### PRJ-061 Issue Collectors

- **Deskripsi:** Entri Project settings > Issue collectors untuk widget pengumpul feedback yang membuat work item di project company-managed. Aturan perilaku lengkap dimiliki INT-006 Issue collector di [25](25-integrasi-impor-ekspor-ekstensi.md) (widget) dan ADM-033 di [22](22-administrasi-sistem.md) (daftar global); ISS-040 di [06](06-work-item-inti.md) juga hanya rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs (konfigurasi); pengunjung website (mengirim).
- **Aturan perilaku (acceptance criteria):**
  1. Issue collector hanya tersedia untuk project company-managed; menu tidak tampil di team-managed (INT-006 butir 1).
  2. Konfigurasi memuat name, description, work type, dan reporter (INT-006 butir 2 dan 8).
  3. Template form mengikuti INT-006 butir 3.
  4. Trigger mengikuti INT-006 butir 4.
  5. Informasi browser dan kontak pengirim dikumpulkan sesuai INT-006 butir 5.
  6. Kode embed dihasilkan sesuai INT-006 butir 6.
  7. Collector dapat di-enable/disable dan dihapus; collector nonaktif menolak kiriman dengan pesan di widget (INT-006 butir 7, ADM-033 butir 2).
  8. Statistik pemakaian collector ditampilkan sesuai INT-006 butir 7.
  9. Field yang dikumpulkan harus ada di create screen work type terpilih (INT-006 butir 3).
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pemilik aturan: INT-006 (widget) dan ADM-033 (daftar global); bila ada perbedaan, kedua fitur itu berlaku. Fitur lama; Forms ([26](26-kerja-bisnis-formulir-persetujuan.md)) adalah alternatif modern untuk intake.

### PRJ-062 Project Properties

- **Deskripsi:** Penyimpanan key-value JSON (entity properties) per project yang dipakai ekstensi dan integrasi.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API; ekstensi.
- **Aturan perilaku (acceptance criteria):**
  1. REST API `GET/PUT/DELETE /project/{key}/properties/{propertyKey}` dan `GET /project/{key}/properties` (daftar key).
  2. ⚠ Nilai berupa JSON valid dengan batas ukuran (usulan 32 KB per property); nilai melebihi batas ditolak 400.
  3. Baca memerlukan Browse Projects; tulis dan hapus memerlukan Administer Projects.
  4. Ekstensi memakai properties untuk menyimpan konfigurasi per project; properties ikut terhapus saat project dihapus permanen.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi entity properties di [23](23-api-publik-webhook.md).

## Administrasi global, API & migrasi project

### PRJ-063 Projects Administration (Jira settings)

- **Deskripsi:** Halaman administrasi global semua project di Settings > Projects.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Manage projects menampilkan semua project di site, termasuk yang tidak dapat di-browse admin (mis. team-managed Private), dengan pencarian dan filter.
  2. Sub-halaman Project categories menyediakan create/edit/delete kategori (PRJ-019).
  3. Sub-halaman Archived projects (Premium/Enterprise) menampilkan daftar project archived dengan aksi Restore.
  4. Sub-halaman Trash menampilkan project di trash dengan aksi Restore dan Delete permanently (PRJ-033).
  5. Setiap baris memiliki akses cepat ke Project settings project tersebut.
  6. Seluruh halaman khusus Admin Situs; pengguna lain menerima 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Struktur menu Settings di [22](22-administrasi-sistem.md).

### PRJ-064 Project REST API & Validation

- **Deskripsi:** Aturan domain dan validasi untuk endpoint REST siklus hidup dan metadata project, dipadankan dengan API bergaya Jira Cloud. Kontrak endpoint publik dimiliki API-036 Projects API di [23](23-api-publik-webhook.md).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform (butir 1-3, 5, dan 10 di Fase 1 karena dipakai SPA; butir 4 bagian archive serta butir 6-9 di Fase 3)
- **Aktor & izin:** Integrator/Developer API; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /project/search` mendukung paginasi (`startAt`, `maxResults`) dan filter `query`, `typeKey`, `categoryId`, `action` (`view`/`browse`/`edit`), `status` (`live`/`archived`/`deleted`), serta `expand` (API-036 butir 1).
  2. `GET`/`PUT`/`DELETE /project/{projectIdOrKey}` dan `POST /project` tersedia; `projectIdOrKey` menerima ID, key aktif, atau key alias (API-036 butir 2-3 dan 6).
  3. `GET /project/recent` mengembalikan project yang terakhir dibuka caller.
  4. `DELETE /project/{key}?enableUndo=true` memindahkan project ke trash, dan `POST /project/{key}/delete` menghapus permanen secara asinkron dengan mengembalikan task yang dapat dipantau. Archive/restore lewat `POST /project/{key}/archive` dan `POST /project/{key}/restore` menyusul di Fase 3 (API-036 butir 7-8).
  5. Validasi lewat `/projectvalidate/key`, `/projectvalidate/validProjectKey`, dan `/projectvalidate/validProjectName` (API-036 butir 9).
  6. (Fase 3) Project types lewat `/project/type` dan `/project/type/{key}/accessible`.
  7. (Fase 3) Project features (team-managed) lewat `GET/PUT /project/{key}/features/{featureKey}`.
  8. (Fase 3) Roles lewat `/project/{key}/role` dan `/project/{key}/roledetails` (PERM-061).
  9. (Fase 3) Tersedia endpoint project categories, avatars, components, versions, statuses (`/project/{key}/statuses`), hierarchy, notification scheme, dan issue security level per project.
  10. Respons memperhatikan permission caller: project yang tidak dapat di-browse tidak dikembalikan di search dan menghasilkan 404 pada akses langsung.
- **Team-managed vs company-managed:** Butir 7 hanya team-managed; field `style` di respons membedakan kedua model (API-036 butir 10).
- **Catatan & ketergantungan:** Pemilik kontrak endpoint: API-036 di [23](23-api-publik-webhook.md); modul ini memiliki aturan domain dan validasi (nama/key PRJ-011, PRJ-012; trash PRJ-033). Fase diselaraskan dengan API-036 (P0 / Fase 1) untuk endpoint inti. Endpoint native di bawah `/api/v1` sesuai konvensi repo; lapisan kompatibilitas path `/rest/api/3/...` diputuskan di [23](23-api-publik-webhook.md). Swagger wajib diperbarui (`make swag`).

### PRJ-065 Project JQL Functions & Fields

- **Deskripsi:** Field dan fungsi JQL terkait project untuk mencari work item berdasarkan atribut project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang memakai pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `project = KEY`, `project = ID`, `project = "Name"`, dan `project in (...)` didukung; key lama (alias) tetap dikenali.
  2. `category = "Nama category"` memfilter work item berdasarkan kategori project.
  3. `projectsLeadByUser([user])` mengembalikan project yang dipimpin pengguna (default pengguna saat ini).
  4. `projectsWhereUserHasPermission("permission")` mengembalikan project tempat pengguna saat ini memiliki permission tersebut.
  5. `projectsWhereUserHasRole("role")` mengembalikan project tempat pengguna saat ini memiliki role tersebut.
  6. Autocomplete project di editor JQL hanya menampilkan project yang dapat di-browse pengguna.
  7. Work item dari project archived atau trash tidak masuk hasil pencarian.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Parser dan evaluator di [12](12-jql-bahasa-kueri.md).

### PRJ-066 Cloud-to-Cloud Project Migration

- **Deskripsi:** Alat migrasi untuk memindahkan project terpilih beserta datanya dari satu instance/site ke instance lain.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih project yang dimigrasi (bukan seluruh site).
  2. Migrasi membawa work item, konfigurasi, pengguna, dan grup terkait.
  3. Pengecekan pra-migrasi menghasilkan laporan konflik (key/nama project, field) sebelum eksekusi.
  4. Alat hanya tersedia untuk Admin Organisasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Prioritas rendah (Fase 4); dibangun sebagai alat ekspor/impor project antar-instance native di atas infrastruktur impor/ekspor [25](25-integrasi-impor-ekspor-ekstensi.md), bukan memakai tool migrasi layanan admin pihak referensi. Cakupan data final ditetapkan saat desain modul 25.

### PRJ-067 Project Import from Backup (Data Center)

- **Deskripsi:** Import satu project dari file backup XML ke instance yang sudah ada; fitur khas Server/Data Center yang tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan, tidak dibangun) Admin memilih file backup dan project yang akan di-restore.
  2. (Rujukan, tidak dibangun) Sistem memvalidasi kecocokan versi, custom field, dan scheme sebelum import.
  3. (Rujukan, tidak dibangun) Project tujuan harus kosong atau belum ada.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pasangan rujukan: INT-027 di [25](25-integrasi-impor-ekspor-ekstensi.md) (juga dc-only-or-deprecated). Tidak tersedia di Jira Cloud; kebutuhan serupa dipenuhi PRJ-066 dan impor di [25](25-integrasi-impor-ekspor-ekstensi.md). Lihat "Di luar cakupan".

### PRJ-068 Project Name/Key Format Configuration (Data Center)

- **Deskripsi:** Konfigurasi panjang maksimum nama/key dan pola regex key lewat advanced settings; fitur khas Data Center yang tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan, tidak dibangun) Pengaturan `jira.projectkey.pattern` untuk pola regex key.
  2. (Rujukan, tidak dibangun) Pengaturan `jira.projectkey.maxlength` dan `jira.projectname.maxlength`.
  3. (Rujukan, tidak dibangun) Daftar kata cadangan untuk key (CON, PRN, AUX, NUL, COM1, LPT1, dll.).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Clone memakai aturan tetap PRJ-011 dan PRJ-012. Lihat "Di luar cakupan".

## Guardrail & agent pada space (tambahan di luar inventaris)

### PRJ-069 ⚠ Space Policies (guardrail team-managed)

- **Deskripsi:** Kebijakan tingkat site yang membatasi apa yang boleh dibuat secara lokal oleh project (space) team-managed, sehingga tim tetap otonom tetapi memakai field dan status global yang distandardisasi. Sumbernya pengumuman rilis 26 Oktober 2026, dan fitur ini belum ada di inventaris.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Situs (membuat, mengedit, menghapus, dan menautkan policy); role Administrator project team-managed (terkena batasan, tetap mengatur hal di luar policy).
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs dapat membuat, mengedit, dan menghapus policy di Settings (Jira settings) > Space policies. Policy yang masih ditautkan ke space atau template menampilkan daftar pemakaiannya dan meminta konfirmasi sebelum dihapus.
  2. Setiap policy memuat opsi "Global fields only" (tim tidak bisa membuat custom field lokal), "Global statuses only" (tim tidak bisa membuat status lokal), atau keduanya. Policy tanpa satu pun opsi aktif ditolak saat disimpan.
  3. Satu policy dapat dijadikan default site dan otomatis diterapkan ke setiap space team-managed baru. Hanya satu policy yang boleh menjadi default pada satu waktu.
  4. Policy dapat ditautkan ke space sumber custom template (PRJ-009), sehingga space yang dibuat dari template tersebut mewarisi policy-nya. Template tanpa policy memakai policy default site, dan bila tidak ada default, space dibuat tanpa policy.
  5. Satu space team-managed hanya boleh memiliki satu policy. Menerapkan policy baru ke space yang sudah punya policy menggantikan policy lama setelah konfirmasi.
  6. Di dalam batas policy, tim tetap dapat mengatur views, kolom board, layout work item, dan pengaturan tim (access, features, notifikasi) tanpa keterlibatan Admin Situs.
  7. Saat policy aktif, UI field editor team-managed (CFG-041) dan workflow editor team-managed (WF-017) menyembunyikan aksi "create field" atau "create status" lokal dan menampilkan alasannya, mis. "Space ini memakai policy <nama>: hanya field global yang diizinkan". Pembuatan lewat API ditolak 403 dengan pesan serupa. ⚠ Field atau status lokal yang sudah ada sebelum policy diterapkan tetap berfungsi, tetapi tidak dapat ditambah.
  8. Setiap pembuatan, perubahan, penghapusan, penetapan default, dan penerapan policy ke space dicatat di audit log (ADM-014) beserta aktor, policy, space, dan nilai sebelum/sesudah.
- **Team-managed vs company-managed:** Hanya team-managed; project company-managed sudah memakai field dan status global lewat scheme.
- **Catatan & ketergantungan:** ⚠ Fitur di luar inventaris (added beyond inventory) berdasarkan pengumuman rilis Oktober 2026, sehingga ketersediaan per plan belum dipastikan (usulan Premium, gating VISI-003). Terkait CFG-041 di [05](05-konfigurasi-work-item-skema.md), WF-017 dan WF-001 (status global) di [09](09-workflow-status-transisi.md), template PRJ-009, serta work type dan workflow team-managed di PRJ-051 dan PRJ-052.

### PRJ-070 ⚠ Agents in space settings (halaman Agents)

- **Deskripsi:** Halaman Space settings > Agents tempat admin space menelusuri, mengaktifkan, dan mengatur agent AI bawaan maupun kustom yang bekerja di project tersebut. Sumbernya rilis Oktober 2026, dan fitur ini belum ada di inventaris. Fitur ini opsional dan nonaktif pada profil regulated deployment (VISI-003).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek/Project Lead (Administer Projects atau role Administrator team-managed); Admin Situs (mengaktifkan AI di level site dan mendefinisikan agent kustom). Permission agent mengikuti PERM-050.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Space settings > Agents menampilkan katalog agent bawaan dan agent kustom yang tersedia di site, dan admin space dapat mengaktifkan atau menonaktifkannya per project. Halaman tidak tampil bila AI dimatikan di level site atau profil regulated deployment aktif.
  2. Setiap agent menampilkan trigger yang menjalankannya: transisi workflow (WF-043), rule otomasi (AUTO-073), atau event work item seperti assign dan @mention (ISS-046).
  3. Admin space dapat mengisi konteks per space berupa instruksi, aturan bisnis, dan glosarium yang diberikan ke agent saat berjalan di project ini. ⚠ Konteks memiliki batas panjang (usulan 10.000 karakter) dan ditolak bila melebihi batas.
  4. Setiap agent memiliki level otonomi per space, minimal "Hanya saran" (tidak mengubah data), "Perlu review" (aksi menunggu persetujuan manusia), dan "Otomatis" (aksi langsung dijalankan dalam batas aturan auto-approval).
  5. Aturan auto-approval berisi kriteria aksi yang boleh dijalankan tanpa review, misalnya hanya mengubah field tertentu, hanya menambah komentar, atau hanya work type tertentu. Aksi di luar kriteria selalu masuk antrean review.
  6. Semua aksi agent (aktivasi, perubahan konfigurasi, aksi yang dijalankan atau disetujui) tercatat di audit log project dan site, dengan aktor agent dan manusia penyetuju dibedakan.
- **Team-managed vs company-managed:** Tidak ada perbedaan; di company-managed izin mengikuti permission scheme (PERM-050), di team-managed mengikuti role Administrator.
- **Catatan & ketergantungan:** ⚠ Fitur di luar inventaris (added beyond inventory) dan opsional, nonaktif secara default pada profil regulated deployment (VISI-003). Terkait PERM-050 di [03](03-izin-peran-keamanan.md), ISS-046 di [06](06-work-item-inti.md), WF-043 di [09](09-workflow-status-transisi.md), dan AUTO-073 di [20](20-otomasi.md). Agent pihak ketiga dari Marketplace di luar cakupan.

## Pertanyaan terbuka

1. Apakah `id` project memakai `bigint` numerik (kompatibel API bergaya Jira) sementara `users.id` tetap UUID `varchar(36)`, atau seluruh entitas diseragamkan? (PRJ-001)
2. Apakah label UI default memakai "Space"/"Work item" (terminologi terbaru) atau "Project"/"Issue", dan apakah admin dapat memilih? (PRJ-001)
3. Apakah key project boleh memuat underscore (`A_B`)? (PRJ-012)
4. Apakah Fase 1 membangun kedua model manajemen sekaligus, atau team-managed lebih dulu dengan company-managed memakai scheme default saja sampai Fase 2? (PRJ-003)
5. Apakah project type company-managed dapat diubah (software ↔ business ↔ service) setelah dibuat? (PRJ-023)
6. Siapa yang boleh mengarsipkan dan me-restore project dari trash: hanya Admin Situs, atau juga Administrator project team-managed? (PRJ-032, PRJ-033)
7. Apakah Admin Situs dapat membaca isi work item project Private tanpa menjadi anggota? (PRJ-039)
8. Apakah opsi Private tersedia di plan Free? (PRJ-039)
9. Apakah migrasi team-managed → company-managed memetakan otomatis "Story point estimate" ke "Story Points"? (PRJ-024)
10. Apakah perubahan Key/Category di Details company-managed dibatasi ke Admin Situs saja? (PRJ-016)
11. Berapa nilai batas final: panjang nama, ukuran avatar, panjang description, jumlah field/work type team-managed, ukuran project property? (PRJ-011, PRJ-016, PRJ-018, PRJ-051, PRJ-062)
12. Apakah key project yang dihapus permanen dapat dipakai lagi? (PRJ-033)
13. Apakah lapisan kompatibilitas REST `/rest/api/3/...` dibangun selain API native `/api/v1`? (PRJ-064)
14. Apakah custom project template (PRJ-009) dibuka untuk semua paid plan seperti referensi terbaru, atau tetap dibatasi Premium/Enterprise di clone? (PRJ-009)
15. Plan mana yang mendapat Space Policies, dan apakah field/status lokal yang sudah ada dibekukan atau wajib dimigrasi ke global saat policy diterapkan? (PRJ-069)
16. Apakah halaman Agents dibangun sama sekali, dan level otonomi apa yang diizinkan secara default? (PRJ-070)

## Di luar cakupan & ekosistem

- **Fitur Data Center-only (dc-only-or-deprecated), rujukan saja:** Project Import from Backup (PRJ-067) dan Project Name/Key Format Configuration termasuk kata cadangan key (PRJ-068).
- **App Marketplace pihak ketiga:** kloning project lengkap beserta data (mis. app sejenis "Deep Clone"), template pack pihak ketiga, dan app yang hanya menambah halaman project; produk ini hanya menyediakan kerangka ekstensi (PRJ-060, [25](25-integrasi-impor-ekspor-ekstensi.md)).
- **Layanan lintas produk Atlassian:** Confluence (Project Pages dibangun lewat konektor dokumentasi generik, PRJ-030), Compass components (PRJ-056), Atlassian Goals/Home, Bitbucket Cloud sebagai layanan bawaan, serta tool migrasi di admin.atlassian.com (PRJ-066 dibangun sebagai alat migrasi antar-instance native di Fase 4, bukan integrasi ke tool tersebut).
- **Infrastruktur email Atlassian:** domain pengirim `atlassian.net`; produk ini memakai domain pengirim milik instance (PRJ-054).
- **Agent AI pihak ketiga dan Default members role:** agent dari Marketplace/penyedia eksternal pada halaman Agents (PRJ-070), serta "Default members" per project role yang merupakan perilaku Data Center (PRJ-040, PERM-067).
- **Konversi langsung model team-managed ↔ company-managed:** sengaja tidak dibangun; jalur resmi adalah migrasi lewat bulk move (PRJ-024).
