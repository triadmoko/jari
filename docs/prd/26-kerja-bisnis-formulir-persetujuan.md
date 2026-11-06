# 26 — Proyek Bisnis, Formulir & Persetujuan

Modul ini mendefinisikan kapabilitas **kerja bisnis** di produk ini: tipe project `business` untuk tim non-teknis (marketing, HR, finance, legal, operations, sales, desain), katalog template bisnis, konfigurasi default (work type, workflow, field, Category), navigasi tab dan view pendukung khas project bisnis (Shortcuts, Attachments, Pages, Reports, background, Get started checklist), **Forms** sebagai saluran intake yang mengubah jawaban formulir menjadi work item, **Approvals** sebagai langkah persetujuan di dalam workflow (non-JSM), serta intake lewat email dan template otomasi bisnis. Setelah penggabungan produk Jira Software dan Jira Work Management menjadi satu produk, kapabilitas ini tidak lagi menjadi produk terpisah: tipe project `business` hanya membedakan template, tab default, dan view default, sementara view bisnis (List, Calendar, Timeline, Forms, Summary) juga dapat dipakai di project software. Mekanik umum (entitas project, workflow engine, field, notifikasi, otomasi, view List/Calendar/Timeline/Summary) dimiliki modul masing-masing; modul ini mengatur **perilaku khusus bisnis** dan pemasangannya.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas project, work item, workflow, form, maupun approval. Seluruh fitur di modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/businessproject` (template, navigasi, Category, shortcut, background, checklist), `internal/features/form` (form builder, sharing, submission), dan `internal/features/approval` (langkah approval, keputusan, riwayat), dengan entitas bersama di `internal/shared/domain/entity` dan UI di `frontend/src/routes` (SPA SvelteKit).

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah Work item/Work type/Space dan tier plan (VISI-002, VISI-003).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — access level dan role team-managed, permission scheme, issue security.
- [04 Proyek](04-proyek.md) — entitas project, tipe project, template library, Features, Shortcuts, Pages, archive/trash.
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — work type, field, screen, Category field (CFG-021).
- [06 Work Item Inti](06-work-item-inti.md) — create, issue view, watcher, move, issue security level.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — attachment, activity/history, tab Approvals di activity (CMT-025).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — workflow editor, status category, global transition, validator.
- [10 Board](10-board-kolom-kartu.md) — board bisnis gaya Kanban (BRD-050).
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) dan [12 JQL](12-jql-bahasa-kueri.md) — filter, List view hasil pencarian, ekspor CSV.
- [16 Laporan & Insight](16-laporan-insight.md) dan [17 Dashboard & Gadget](17-dashboard-gadget.md) — laporan bisnis dan gadget.
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — view List, Calendar, Timeline, Summary, Board bisnis, Plans.
- [19 Notifikasi & Email](19-notifikasi-email.md) — notifikasi approval, email masuk menjadi work item (NTF-034).
- [20 Otomasi](20-otomasi.md) — trigger "Work item created"/"Form submitted", template rule bisnis.
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — sidebar, navigasi project, onboarding.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log, setting attachment, data lifecycle.
- [23 REST API & Webhook](23-api-publik-webhook.md), [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — `projectTypeKey`, migrasi dari server/DC.
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) dan [32 ITSM: Insiden, Problem, Perubahan & Layanan](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — JSM forms (JSM-PRJ-034 s.d. JSM-PRJ-037) dan JSM approvals (JSM-ITSM-011 s.d. JSM-ITSM-015) yang lebih lengkap (dicatat terpisah).

## Persona terkait

- **Admin Situs** — membuat project business company-managed, mengubah tipe project, mengelola scheme, field configuration, dan setting email/attachment global.
- **Admin Proyek/Project Lead** — memilih template, mengatur Features, role anggota, Category, shortcut, background, form, email requests, dan langkah approval.
- **Pengguna Tim Bisnis** — memakai view List/Board/Calendar/Timeline, mengisi Category, menjadi approver, membuat dan mengisi form.
- **Product Owner/PM** — memakai template Project management/Go-to-market, Timeline, dan laporan bisnis.
- **Stakeholder/Viewer** — melihat progres project bisnis tanpa mengubah work item (role Viewer).
- **Developer/Anggota Tim** — memakai Forms di project software untuk intake bug report dan feature request.
- **Tamu/Kolaborator Eksternal** — submitter form publik (bila diaktifkan) atau pengirim email ke alamat project.
- **Integrator/Developer API** — membaca `projectTypeKey=business`, membuat work item dari integrasi, dan memakai template otomasi.

## User story modul

1. Sebagai Admin Proyek/Project Lead, saya ingin membuat project bisnis dari galeri template per kategori tim agar proses marketing atau HR saya langsung punya workflow, field, dan view yang relevan (BIZ-007, BIZ-008, BIZ-016).
2. Sebagai Pengguna Tim Bisnis, saya ingin project bisnis yang sederhana tanpa sprint, backlog, dan tab Development agar tim non-teknis tidak bingung dengan konsep software (BIZ-002, BIZ-022).
3. Sebagai Admin Proyek/Project Lead, saya ingin menyalakan atau mematikan view seperti Calendar, Timeline, atau Forms tanpa kehilangan data agar navigasi project sesuai cara kerja tim (BIZ-004, BIZ-022).
4. Sebagai Pengguna Tim Bisnis, saya ingin memberi Category berwarna pada work item agar pekerjaan per workstream mudah dibedakan di Board, Calendar, dan Timeline (BIZ-021).
5. Sebagai Admin Proyek/Project Lead, saya ingin menyusun form intake dengan drag-drop yang setiap pertanyaannya dipetakan ke field agar permintaan masuk sudah terstruktur (BIZ-030, BIZ-031, BIZ-032).
6. Sebagai Pengguna Tim Bisnis, saya ingin mengirim permintaan lewat tautan form dan langsung mendapat tautan ke work item yang tercipta agar saya bisa memantau tindak lanjutnya (BIZ-035, BIZ-037, BIZ-038).
7. Sebagai Admin Proyek/Project Lead, saya ingin menambahkan langkah approval pada status review agar work item hanya maju setelah approver yang ditunjuk memutuskan (BIZ-042, BIZ-043).
8. Sebagai Pengguna Tim Bisnis yang ditunjuk sebagai approver, saya ingin menerima notifikasi dan menekan Approve/Decline di issue view agar persetujuan cepat dan terekam (BIZ-044, BIZ-045).
9. Sebagai Admin Situs, saya ingin jejak keputusan approval yang tidak bisa diedit agar audit proses pengadaan dan dokumen dapat dipertanggungjawabkan (BIZ-046, BIZ-013).
10. Sebagai Tamu/Kolaborator Eksternal, saya ingin mengirim email ke alamat project agar permintaan saya tercatat sebagai work item tanpa harus memahami tool ini (BIZ-047).
11. Sebagai Developer/Anggota Tim, saya ingin memakai Forms di project software agar bug report dari tim lain masuk dengan field yang lengkap (BIZ-041).
12. Sebagai Admin Proyek/Project Lead, saya ingin Get started checklist yang menyesuaikan template agar anggota baru tahu langkah awal memakai project (BIZ-023).
13. Sebagai Product Owner/PM, saya ingin laporan dan Summary project bisnis serta ekspor jawaban form agar saya bisa menganalisis volume permintaan (BIZ-027, BIZ-040).
14. Sebagai Admin Proyek/Project Lead, saya ingin memakai template otomasi bisnis seperti pengingat due date dan pengisian approver otomatis agar pekerjaan rutin tidak dilakukan manual (BIZ-048).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 7 |
| P1 | 21 |
| P2 | 15 |
| P3 | 5 |
| **Total** | **48** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 5 |
| Fase 2 | 4 |
| Fase 3 | 32 |
| Fase 4 | 7 |
| Fase — | 0 |

Strategi urutan implementasi: Fase 1 cukup menyediakan tipe project `business` (BIZ-002) dengan template Task tracking (BIZ-008) serta work type, workflow, dan field default (BIZ-018 s.d. BIZ-020), sehingga tim bisnis sudah bisa memakai mesin project, work item, dan board inti tanpa fitur khusus. Fase 2 menambahkan galeri template (BIZ-007), navigasi tab bisnis (BIZ-022), label terminologi yang dapat dikonfigurasi (BIZ-001), dan role/akses (BIZ-003), bersamaan dengan view List/Calendar/Timeline dari modul 18. Fase 3 adalah inti modul ini: Forms (BIZ-029 s.d. BIZ-041 kecuali conditional logic dan sharing publik), Approvals (BIZ-042 s.d. BIZ-046), Category, Features toggle, email requests, template bisnis P1/P2, dan view pendukung; Forms dibangun lebih dulu karena Approvals sering bergantung pada field approver yang diisi lewat form. Fase 4 menampung fitur yang masih perlu konfirmasi atau bergantung pada integrasi luar (conditional logic, form publik, Pages, ubah tipe project, background, template P3, pembedaan tier). Dua entri inventaris approvals (domain workflow dan domain kerja bisnis) digabung di modul ini menjadi satu fitur BIZ-042.

## Fitur

## Fondasi produk & tipe project bisnis

### BIZ-001 Jira Work Management retirement & unified Jira naming

- **Deskripsi:** Pensiunnya Jira Work Management sebagai produk terpisah dan penyatuan menjadi satu produk dengan tipe project `business`. Untuk clone, fitur ini berarti satu lisensi produk, tipe project `business` yang tetap ada di API, view bisnis yang juga tersedia di project software, dan label terminologi yang dapat dikonfigurasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Admin Situs (lisensi, product access, label terminologi tingkat site); Admin Proyek; semua pengguna (melihat label).
- **Aturan perilaku (acceptance criteria):**
  1. Produk hanya memiliki satu product access "produk inti" per plan (Free/Standard/Premium/Enterprise) yang mencakup project software dan business; tidak ada seat atau application role terpisah untuk kerja bisnis, dan data lisensi lama (bila diimpor) dipetakan otomatis ke product access tunggal ini.
  2. Nilai `projectTypeKey` `business` tetap dipakai di REST API (mis. `GET /rest/api/3/project/{key}` mengembalikan `"projectTypeKey": "business"`) dan ditampilkan sebagai nilai kolom Type di project directory.
  3. View bisnis (List, Calendar, Timeline, Forms, Summary) dapat ditambahkan ke project software; yang membedakan tipe project hanyalah template dan set tab/view default saat project dibuat.
  4. UI menampilkan "Work item" untuk issue dan "Work type" untuk issue type, sedangkan API, JQL (`issue`, `issuetype`), nama kolom DB, dan nama event tetap memakai istilah lama.
  5. ⚠ UI dapat menampilkan "Space" sebagai label untuk project (mis. "Space settings", "Create space") lewat flag konfigurasi; cakupan rename (menu, URL, email) mengikuti keputusan di VISI-002 dan dapat dimatikan per site.
  6. Semua label terminologi (project/space, issue/work item, issue type/work type) berasal dari kunci i18n (EN/ID) yang dapat dikonfigurasi, bukan teks hard-coded; pengujian memastikan tidak ada string literal "issue"/"project" di komponen UI yang tampil ke pengguna.
  7. Route SPA view project bisnis memakai pola `/projects/{KEY}/{view}` (padanan pola `/jira/core/projects/{KEY}/{view}` di Jira Cloud) dan tidak di-redirect ke route khusus software; URL view yang dibagikan tetap stabil setelah tipe project berubah.
  8. Dokumentasi dan teks bantuan produk menyebut kapabilitas ini sebagai "kerja bisnis"/"project untuk tim bisnis", bukan sebagai produk terpisah.
  9. Impor dari instance server/Data Center berbasis Jira Core atau Jira Work Management memetakan project-nya menjadi project `business` (lihat modul 25).
  10. Fitur khusus produk lama yang tidak dibawa ke produk terpadu (mis. Overviews lintas project) tidak dibangun; kebutuhan lintas project diarahkan ke Plans ([18](18-timeline-rencana-tampilan-proyek.md)).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Terminologi dasar di [VISI-002](00-visi-ruang-lingkup-glosarium-baseline.md) dan entitas project di PRJ-001/PRJ-002 ([04](04-proyek.md)). Impor/migrasi di modul 25. Repo saat ini belum punya infrastruktur i18n di frontend (hanya pesan galat EN/ID di `internal/shared/constants/error.go`); kamus label UI perlu ditambahkan di SPA.

### BIZ-002 Business project type (company-managed & team-managed)

- **Deskripsi:** Tipe project untuk tim non-teknis dengan view yang lebih sederhana, tersedia dalam mode company-managed dan team-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (membuat company-managed; global permission Administer Jira), pengguna dengan global permission "Create team-managed projects", Admin Proyek, anggota project.
- **Aturan perilaku (acceptance criteria):**
  1. Wizard pembuatan project business meminta pilihan jenis pengelolaan: company-managed (memakai scheme bersama) atau team-managed (konfigurasi mandiri); pilihan disimpan sebagai `style` project dan tidak dapat diubah setelah dibuat.
  2. Project business company-managed memakai workflow scheme, issue type scheme, field configuration (scheme), screen scheme, permission scheme, notification scheme, dan issue security scheme; setiap scheme dapat diganti oleh Admin Situs di Project settings.
  3. Project business team-managed menyimpan work type, field, workflow, dan role secara project-scoped; Admin Proyek mengubahnya langsung tanpa scheme dan perubahan tidak memengaruhi project lain.
  4. Project business team-managed memiliki access level Open, Limited, atau Private dengan perilaku sesuai PERM-038 (Open: semua pengguna berlisensi bisa melihat dan membuat; Limited: semua bisa melihat, hanya anggota yang mengubah; Private: hanya anggota).
  5. Secara default project business tidak memiliki sprint, backlog, konfigurasi JQL board, maupun tab Development/Code; field Sprint dan Story points tidak ada di layar default.
  6. Hanya Admin Situs yang dapat membuat project business company-managed; pengguna dengan global permission "Create team-managed projects" dapat membuat team-managed; pengguna lain tidak melihat opsi tersebut dan API mengembalikan 403.
  7. Halaman Details memuat Name, Key, avatar, Project lead, Default assignee, Project category (company-managed saja), dan URL, dengan validasi yang sama seperti project lain (PRJ-012).
  8. Archive, trash (pemulihan 60 hari), restore, dan hapus permanen project business mengikuti lifecycle project umum (PRJ-032, PRJ-033) tanpa aturan khusus.
  9. Project team-managed tidak dapat dikonversi ke company-managed atau sebaliknya; UI menampilkan panduan untuk membuat project baru lalu memindahkan work item secara bulk move (ISS-032), dan API konversi tidak tersedia.
- **Team-managed vs company-managed:** Kedua mode didukung; company-managed mengambil konfigurasi dari scheme global, team-managed dari konfigurasi project-scoped dan access level.
- **Catatan & ketergantungan:** Entitas dan tipe project di PRJ-002/PRJ-003 ([04](04-proyek.md)); scheme di [03](03-izin-peran-keamanan.md), [05](05-konfigurasi-work-item-skema.md), [09](09-workflow-status-transisi.md). Fase 1 cukup mendukung team-managed dengan template Task tracking; dukungan company-managed penuh mengikuti ketersediaan scheme di Fase 2-3.

### BIZ-003 Business project roles & access

- **Deskripsi:** Role dan akses bawaan project business yang mengatur siapa yang mengelola, berkontribusi, atau hanya melihat.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (role Administrator atau permission Administer projects), Admin Situs, Admin Situs organisasi (product access).
- **Aturan perilaku (acceptance criteria):**
  1. Project business team-managed menyediakan role bawaan Administrator, Member, dan Viewer dengan hak sesuai PERM-039.
  2. ⚠ Project business team-managed mendukung custom role dengan permission granular (mis. tanpa hak delete, tanpa hak comment) pada plan Standard ke atas; pada plan Free menu custom role tampil dengan badge upgrade.
  3. Project business company-managed memakai project role global (termasuk Administrators) dan permission scheme yang ditetapkan Admin Situs.
  4. ⚠ Role Viewer dapat browse project, membuka work item, dan melihat semua view, tetapi tidak dapat membuat atau mengedit work item; hak menambah komentar untuk Viewer mengikuti matriks PERM-039 dan dapat dikonfigurasi.
  5. Dialog "Add people" menerima nama, email, atau grup, menyediakan dropdown role (default Member), dan memberi pesan galat bila pengguna sudah menjadi anggota.
  6. Pengguna tanpa product access produk inti tidak dapat masuk ke project meskipun ditambahkan sebagai anggota; dialog Add people menawarkan untuk mengundang atau memberi akses bila pengundang adalah admin.
- **Team-managed vs company-managed:** Team-managed memakai role Administrator/Member/Viewer plus custom role; company-managed memakai project role dan permission scheme.
- **Catatan & ketergantungan:** Detail permission di PERM-018, PERM-038, PERM-039, PERM-040 ([03](03-izin-peran-keamanan.md)) dan PRJ-038/PRJ-039 ([04](04-proyek.md)). Repo saat ini belum memiliki konsep role/permission selain autentikasi JWT; claims `sess` perlu diperluas atau permission dihitung per request.

### BIZ-004 Team-managed business features toggles

- **Deskripsi:** Halaman Features di Project settings project business team-managed untuk menyalakan atau mematikan view dan kapabilitas project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (role Administrator team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman Features menampilkan toggle untuk Summary, List, Calendar, Timeline, Forms, Pages, Shortcuts, dan Attachments, masing-masing dengan nama, deskripsi singkat, dan status on/off; daftar final ditetapkan di konfigurasi produk agar mudah ditambah.
  2. Mematikan fitur hanya menyembunyikan tab/view dan entry point-nya; data terkait (form, shortcut, Category, nilai field) tetap tersimpan dan muncul kembali utuh saat fitur dinyalakan lagi.
  3. Hanya Administrator project yang dapat mengubah toggle; Member dan Viewer melihat halaman (bila dibuka via URL) dalam mode baca atau mendapat 403 dari API.
  4. Perubahan toggle langsung berlaku untuk semua anggota tanpa reload manual (navigasi diperbarui pada pemuatan berikutnya atau lewat event realtime bila tersedia), dan akses ke URL view yang dimatikan menampilkan pesan "Fitur ini dinonaktifkan untuk project ini".
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah PRJ-050 (Features company-managed) dan pengaturan tab di BIZ-022.
- **Catatan & ketergantungan:** Mekanisme Features umum di PRJ-049/PRJ-050 ([04](04-proyek.md)) dan toggle board di BRD-049 ([10](10-board-kolom-kartu.md)). Toggle Approvals dan Goals sengaja tidak dimasukkan karena belum terkonfirmasi.

### BIZ-005 Change project type (business ↔ software)

- **Deskripsi:** Admin Situs dapat mengubah tipe project company-managed antara business dan software sehingga navigasi dan fitur default ikut berubah, tanpa mengubah data.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (global permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Opsi "Project type" tersedia di Project settings > Details hanya untuk project company-managed dan hanya untuk Admin Situs; Admin Proyek melihat nilainya tanpa bisa mengubah.
  2. Perubahan tipe tidak mengubah work item, workflow, maupun scheme; yang berubah adalah tab navigasi default dan fitur khusus tipe (mis. board Scrum, Development, Releases untuk software; Forms, Calendar, Category untuk business).
  3. ⚠ Board Scrum/Kanban yang sudah ada tidak dihapus saat tipe diubah ke business; board tetap dapat diakses via URL dan daftar board, walau tidak tampil di navigasi default.
  4. Project team-managed tidak menampilkan opsi ubah tipe; API mengembalikan 400 dengan pesan "Tipe project team-managed tidak dapat diubah".
  5. Sebelum menyimpan, dialog konfirmasi menjelaskan dampak pada navigasi dan fitur; pembatalan tidak mengubah apa pun.
  6. Setiap perubahan tipe tercatat di audit log (aktor, project, tipe lama, tipe baru, waktu) sesuai ADM-014.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed tidak bisa berganti tipe.
- **Catatan & ketergantungan:** Audit log di ADM-014 ([22](22-administrasi-sistem.md)); tipe project di PRJ-002 ([04](04-proyek.md)). Karena view bisnis juga tersedia di project software (BIZ-001), nilai fitur ini kecil dan ditempatkan di Fase 4.

### BIZ-006 Business plans & tier differences

- **Deskripsi:** Perbedaan kemampuan fitur bisnis antar plan (Free, Standard, Premium, Enterprise) di produk terpadu.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs, admin billing.
- **Aturan perilaku (acceptance criteria):**
  1. Plan Free (maksimum 10 pengguna) mencakup project business, List, Board, Calendar, Timeline, Forms, dan Summary; menambah pengguna ke-11 ditolak dengan pesan ajakan upgrade.
  2. Plan Standard menaikkan batas pengguna, membuka role/permission lengkap (custom role, permission scheme penuh), serta batas storage dan eksekusi automation yang lebih besar.
  3. Plan Premium membuka Plans lintas project termasuk project business (pengganti Overviews), sandbox, dan batas automation lebih tinggi.
  4. Plan Enterprise membuka multiple instance, analytics, dan kontrol keamanan lanjutan.
  5. Fitur yang tidak termasuk plan aktif tampil dengan badge tier (mis. "Premium") dan tombol ajakan upgrade; aksi pada fitur tersebut diblokir di server dengan galat 402/403 yang konsisten, bukan hanya disembunyikan di UI.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Matriks tier dan feature flag di VISI-003 ([00](00-visi-ruang-lingkup-glosarium-baseline.md)). Billing nyata dan halaman pricing berada di luar cakupan fungsional modul ini; clone open source dapat memperlakukan tier sebagai feature flag konfigurasi.

## Galeri & katalog template bisnis

### BIZ-007 Business project template gallery

- **Deskripsi:** Galeri template untuk membuat project business, dikelompokkan per kategori tim, dengan preview work type, workflow, dan view sebelum project dibuat.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (template company-managed), pengguna dengan global permission "Create team-managed projects" (template team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Galeri menampilkan kategori Work management, Marketing, Human resources, Finance, Design, Personal, Operations, Legal, dan Sales; label kategori berasal dari konfigurasi katalog (i18n) sehingga dapat diubah tanpa rilis kode.
  2. Setiap kartu template menampilkan nama, deskripsi singkat, ilustrasi, dan badge jenis pengelolaan yang didukung (team-managed, company-managed, atau keduanya).
  3. Halaman preview menampilkan daftar work type, diagram atau daftar status workflow, view yang direkomendasikan, dan field kunci template sebelum pengguna berkomitmen.
  4. Tombol "Use template" membuka langkah "Choose a project type" (team-managed/company-managed) hanya bila template mendukung keduanya dan pengguna berhak atas keduanya; jika hanya satu yang mungkin, langkah ini dilewati.
  5. Form detail project meminta Name, Key (dibuat otomatis dari nama, dapat diedit, divalidasi unik dan sesuai format dengan pesan galat inline), Access (team-managed: Open/Limited/Private), dan opsi "Share settings with an existing project" (company-managed: memakai scheme project terpilih, yang mengabaikan konfigurasi bawaan template).
  6. Galeri menampilkan rekomendasi template berdasarkan tim/role yang dipilih pengguna saat onboarding (mis. pengguna marketing melihat template Marketing di urutan teratas); bila data onboarding tidak ada, urutan default dipakai.
  7. Galeri yang sama juga memuat template software dan service (bila produk aktif), dengan filter per kategori dan pencarian nama template.
  8. Setelah dibuat, project membuka view default template (Summary, List, atau Board) dan menampilkan Get started checklist (BIZ-023).
  9. Field, work type, workflow, form, dan Category bawaan template dibuat otomatis dalam satu transaksi saat project dibuat; bila salah satu gagal, pembuatan project dibatalkan seluruhnya dan galat ditampilkan.
- **Team-managed vs company-managed:** Template team-managed membuat konfigurasi project-scoped; template company-managed membuat (atau memakai ulang) scheme global bernama sesuai key project, kecuali dipilih "Share settings".
- **Catatan & ketergantungan:** Template library umum di PRJ-005/PRJ-007 ([04](04-proyek.md)); onboarding di NAV-023 ([21](21-navigasi-ux-preferensi.md)). Daftar template disimpan sebagai definisi deklaratif (mis. JSON/YAML ter-embed di binary) agar template baru bisa ditambah tanpa migrasi skema.

### BIZ-008 Template: Task tracking

- **Deskripsi:** Template paling sederhana untuk melacak daftar tugas tim dan menjadi baseline project business di MVP.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (membuat), anggota project (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Template membuat work type Task (level 0) dan Subtask (level -1).
  2. Template membuat workflow To Do → In Progress → Done dengan status category To Do, In Progress, dan Done.
  3. View default adalah List dan Board; Calendar tersedia sebagai view opsional yang dapat ditambahkan lewat tombol "+" (BIZ-022).
  4. Layar default memuat field minimal Summary, Assignee, Due date, Priority, dan Labels.
  5. Template ini adalah template project business yang wajib tersedia di Fase 1, bahkan sebelum galeri lengkap (BIZ-007) ada; wizard Fase 1 dapat langsung memakainya sebagai default untuk tipe `business`.
- **Team-managed vs company-managed:** Tersedia di kedua mode; di Fase 1 minimal team-managed.
- **Catatan & ketergantungan:** Workflow default di BIZ-019, field di BIZ-020; board bisnis di BRD-050 ([10](10-board-kolom-kartu.md)).

### BIZ-009 Template: Blank project (business)

- **Deskripsi:** Template kosong project business dengan konfigurasi minimum untuk tim yang ingin menyusun sendiri.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Template hanya membuat work type Task dan Subtask.
  2. Template membuat workflow To Do → In Progress → Done.
  3. View default yang aktif adalah Summary, List, Board, Calendar, dan Timeline.
  4. Template tidak membuat field custom, Category, maupun form; layar hanya memuat system field default (BIZ-020).
- **Team-managed vs company-managed:** Tersedia di kedua mode.
- **Catatan & ketergantungan:** Ketersediaan template ini di kategori bisnis Jira Cloud saat ini belum terkonfirmasi; untuk clone, template ini tetap disediakan karena murah dan berguna. Padanan software di PRJ-008 ([04](04-proyek.md)).

### BIZ-010 Template: Project management

- **Deskripsi:** Template bisnis serbaguna untuk merencanakan dan melacak project tim dengan view List, Board, Calendar, dan Timeline.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (membuat), anggota project (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat work type Task dan Subtask, dengan Epic sebagai parent level 1 bila hierarki diaktifkan di site.
  2. Template membuat workflow To Do → In Progress → Done.
  3. View yang direkomendasikan dan aktif secara default: Summary, Timeline, List, Board, Calendar, dan Forms.
  4. Layar memuat field kunci Summary, Description, Assignee, Reporter, Priority, Start date, Due date, Labels, Category, dan Attachment.
  5. Template tersedia dalam mode team-managed dan company-managed, sehingga langkah "Choose a project type" muncul di galeri.
- **Team-managed vs company-managed:** Tersedia di kedua mode dengan konfigurasi setara.
- **Catatan & ketergantungan:** Timeline bisnis di PLAN-015, Calendar di PLAN-022 ([18](18-timeline-rencana-tampilan-proyek.md)); hierarki Epic di [05](05-konfigurasi-work-item-skema.md).

### BIZ-011 Template: Process control

- **Deskripsi:** Template untuk proses berulang yang terstandar dengan langkah review atau persetujuan.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, approver, anggota project.
- **Aturan perilaku (acceptance criteria):**
  1. Template membuat work type Task dan Subtask.
  2. ⚠ Template membuat workflow multi-langkah dengan status review/approval: Open, In Progress, Under Review, Approved, Rejected, Done.
  3. Pada mode company-managed, transisi ke status akhir (Approved/Rejected/Done) mengisi field Resolution (mis. Done atau Won't do) lewat post function, dan transisi keluar dari status akhir mengosongkannya.
  4. Status Under Review dapat langsung diberi langkah approval (BIZ-042) dari template; bila diaktifkan, hasil Approved diarahkan ke status Approved dan hasil Declined ke status Rejected secara otomatis.
- **Team-managed vs company-managed:** Team-managed mengisi Resolution=Done secara otomatis saat work item masuk status berkategori Done dan mengosongkannya saat keluar; company-managed mengisi Resolution hanya lewat transition screen atau post function (butir 3), sesuai WF-009.
- **Catatan & ketergantungan:** Berasal dari template Jira Core; dukungan team-managed belum terkonfirmasi. Semantik Resolution dan reopen di WF-009 ([09](09-workflow-status-transisi.md)).

### BIZ-012 Template: Document approval

- **Deskripsi:** Template untuk mengelola siklus review dan persetujuan dokumen.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, pemilik dokumen (reporter), approver.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat workflow Draft/To Do → In Review → Approved/Rejected → Done.
  2. Status In Review dikonfigurasi dengan langkah approval (BIZ-042) yang memakai field Approvers sebagai sumber approver.
  3. Layar memuat field Approvers (multi user picker), Due date, dan Attachment; Approvers dibuat sebagai custom field template bila belum ada.
  4. Work item dapat ditautkan ke halaman dokumentasi lewat tab Pages (BIZ-026) atau tautan halaman di issue view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Approver di BIZ-043; tautan halaman di CMT-030 ([08](08-komentar-lampiran-aktivitas.md)).

### BIZ-013 Template: Procurement

- **Deskripsi:** Template untuk melacak permintaan pembelian dari pengajuan dan persetujuan sampai barang diterima.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Requester (Pengguna Tim Bisnis), approver (manager/finance), tim procurement.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat workflow Requested → In Review → Approved/Rejected → Ordered → Received/Done.
  2. ⚠ Template membuat field bisnis Vendor (text), Cost (number dengan dua desimal), dan Quantity (number bulat ≥ 1).
  3. ⚠ Template membuat form intake "Purchase request" yang memetakan pertanyaan ke Summary, Vendor, Cost, Quantity, Description, dan Attachment.
  4. Status In Review memiliki langkah approval oleh manager/finance (BIZ-042) dengan hasil Approved menuju status Approved dan hasil Declined menuju status Rejected; transisi Approved → Ordered dilakukan manual oleh tim procurement.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Form intake di BIZ-030; template otomasi pengisian approver di BIZ-048.

### BIZ-014 Template: Recruitment

- **Deskripsi:** Template untuk melacak kandidat di sepanjang pipeline rekrutmen.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Recruiter, hiring manager.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat work type Candidate sebagai work type standar project (level 0).
  2. ⚠ Template membuat workflow pipeline Applied → Screening → Interview → Offer → Hired/Rejected.
  3. Board menampilkan setiap status sebagai kolom pipeline kandidat dan kartu dapat di-drag antar tahap (transisi dijalankan sesuai workflow).
  4. Layar Candidate memuat field Attachment untuk CV, dengan tipe file mengikuti setting attachment site.
  5. Saat memilih template ini, wizard menyarankan access Private atau Limited (default Private) dan menampilkan peringatan bahwa data kandidat bersifat sensitif.
- **Team-managed vs company-managed:** Saran access level hanya berlaku di team-managed; di company-managed peringatan menyarankan permission scheme terbatas.
- **Catatan & ketergantungan:** Board bisnis di BRD-050 ([10](10-board-kolom-kartu.md)); access level di PERM-038 ([03](03-izin-peran-keamanan.md)).

### BIZ-015 Template: Lead tracking & Sales pipeline

- **Deskripsi:** Template untuk melacak prospek penjualan dari lead sampai closing.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sales rep, sales manager.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat workflow Lead → Contacted → Qualified → Proposal → Won/Lost.
  2. ⚠ Template membuat field Company (text), Contact (text), dan Deal value (number/currency).
  3. Board berfungsi sebagai sales pipeline dengan kolom per status dan kartu yang menampilkan Deal value.
  4. Template membuat Category contoh untuk segmen atau sumber lead (mis. Inbound, Outbound, Referral).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Apakah "Sales pipeline" merupakan template terpisah dari "Lead tracking" belum terkonfirmasi; clone menyediakan satu template gabungan. CRM penuh berada di luar cakupan.

### BIZ-016 Template: Marketing templates (Content management, Campaign management, Email campaign, Event planning, Go-to-market)

- **Deskripsi:** Kumpulan template marketing untuk konten, kampanye, email, event, dan peluncuran produk.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Tim marketing (Pengguna Tim Bisnis), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Content management: workflow Idea → Draft → Review → Approved → Published.
  2. Campaign management: work type tugas kampanye dengan view Timeline sebagai jadwal kampanye (Start date dan Due date wajib tampil di layar).
  3. ⚠ Email campaign: workflow dengan tahap Design, Review, dan Scheduled/Sent (jadwal kirim).
  4. Event planning: tugas event dengan view Calendar sebagai view utama dan Due date di layar create.
  5. Go-to-market: koordinasi peluncuran lintas tim dengan view Timeline dan Category per workstream (mis. Product, Marketing, Sales, Support) yang dibuat otomatis.
  6. Semua template marketing menjadikan Calendar dan Timeline sebagai view yang tampil di urutan awal tab.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Calendar dan Timeline di [18](18-timeline-rencana-tampilan-proyek.md); Category di BIZ-021. Pengiriman email kampanye ke pelanggan berada di luar cakupan.

### BIZ-017 Template: Other business templates (HR, Finance, Legal, Operations, Personal, Design)

- **Deskripsi:** Template bisnis lain untuk onboarding karyawan, penilaian kinerja, anggaran, kebijakan, aset, tugas pribadi, desain, dan kontrak.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek, tim bisnis (HR, finance, legal, operations, desain).
- **Aturan perilaku (acceptance criteria):**
  1. Employee onboarding: work item per karyawan baru dengan Subtask per langkah checklist yang dibuat otomatis dari template.
  2. ⚠ Performance review: work item per karyawan per periode review dengan workflow review.
  3. ⚠ Budget creation: workflow penyusunan anggaran dengan langkah approval (BIZ-042).
  4. ⚠ Policy management: workflow draft, review, dan publikasi kebijakan.
  5. ⚠ Asset tracking: pelacakan aset sederhana sebagai work item (bukan CMDB).
  6. Personal task planner: dibuat sebagai team-managed dengan access Private secara default.
  7. ⚠ Design/creative request: intake permintaan desain lewat form bawaan template (BIZ-030).
  8. ⚠ Legal/contract request tracking: intake dan pelacakan permintaan kontrak dengan status review.
- **Team-managed vs company-managed:** Personal task planner khusus team-managed; template lain tersedia di kedua mode.
- **Catatan & ketergantungan:** Keberadaan tiap template di Jira Cloud saat ini belum terkonfirmasi; clone dapat menambahkannya bertahap lewat definisi template deklaratif. Asset tracking yang sesungguhnya (CMDB) ada di modul 33 (Assets).

## Konfigurasi default project bisnis

### BIZ-018 Default business work types

- **Deskripsi:** Work type bawaan project business beserta aturan hierarki dan pengelolaannya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed), Admin Situs (company-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Project business memiliki work type standar Task (level 0) dan Subtask (level -1); Epic/parent level 1 tersedia bila hierarki diaktifkan, dan Subtask hanya dapat dibuat di bawah work item level 0.
  2. ⚠ Template dapat menambahkan work type khusus (mis. Candidate, Lead) yang berlaku sebagai level 0 dengan ikon dan deskripsi sendiri.
  3. Team-managed: Admin Proyek dapat menambah, mengubah (nama, ikon, deskripsi), dan menghapus work type di Project settings > Work types; nama work type unik dalam project (tanpa membedakan huruf besar/kecil).
  4. Company-managed: work type diatur lewat issue type scheme; work type global hanya dibuat oleh Admin Situs dan Admin Proyek tidak dapat menambah work type baru.
  5. Menghapus work type yang masih dipakai mewajibkan pemilihan work type pengganti; semua work item terkait dipindahkan ke work type tersebut sebelum penghapusan selesai, dan penghapusan work type terakhir di level 0 ditolak.
  6. ⚠ Jumlah work type per project team-managed dibatasi oleh nilai konfigurasi (default 30); melewati batas menampilkan pesan galat yang menyebut batasnya.
- **Team-managed vs company-managed:** Team-managed mengelola work type per project; company-managed lewat issue type scheme global.
- **Catatan & ketergantungan:** Konfigurasi mendalam work type dan hierarki di [05](05-konfigurasi-work-item-skema.md).

### BIZ-019 Default business workflows & statuses

- **Deskripsi:** Workflow default project business dan cara mengeditnya di team-managed dan company-managed.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed), Admin Situs (company-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Workflow dasar To Do → In Progress → Done dengan status category To Do (abu-abu), In Progress (biru), dan Done (hijau); semua status saling dapat ditransisikan secara default.
  2. Team-managed: workflow diedit per work type di editor sederhana; menambah kolom pada Board project membuat status baru di workflow work type yang ditampilkan board tersebut.
  3. Company-managed: workflow diatur lewat workflow scheme dan workflow editor baru, yang mendukung penambahan langkah Approval pada status (BIZ-042).
  4. Status category menentukan warna lozenge status di semua view dan menentukan perhitungan "completed" pada Summary (work item berkategori Done dihitung selesai).
  5. Opsi "Allow all statuses to transition to this one" (global transition) tersedia per status dan diaktifkan secara default di workflow bawaan template bisnis.
  6. ⚠ Template bisnis company-managed menyertakan post function set/clear resolution pada transisi ke dan dari status Done (WF-009); team-managed mengisi Resolution=Done otomatis saat masuk kategori Done dan mengosongkannya saat keluar.
- **Team-managed vs company-managed:** Team-managed memakai editor workflow per work type tanpa scheme; company-managed memakai workflow scheme dan editor workflow baru.
- **Catatan & ketergantungan:** Mekanik workflow umum di WF-002, WF-006, WF-017, WF-021 ([09](09-workflow-status-transisi.md)). Workflow default Fase 1 disimpan sebagai seed per project saat project dibuat. Semantik Resolution di WF-009 ([09](09-workflow-status-transisi.md)).

### BIZ-020 Default business fields & screens

- **Deskripsi:** Set field bawaan project business dan penempatannya di layar create dan issue view.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed), Admin Situs (company-managed).
- **Aturan perilaku (acceptance criteria):**
  1. System field yang tersedia: Summary, Description, Assignee, Reporter, Priority, Labels, Due date, Start date, Attachment, Parent, Linked work items, Subtasks, dan Category.
  2. Start date dan Due date menjadi sumber penempatan work item di Calendar dan Timeline; work item tanpa kedua tanggal muncul di panel "unscheduled".
  3. Team-managed: Admin Proyek mengatur field per work type dengan drag-drop ke bagian Description fields atau Context fields; field yang dihapus dari layout tidak menghapus nilainya.
  4. Company-managed: field diatur lewat field configuration dan screen scheme oleh Admin Situs.
  5. Field custom bawaan template dibuat otomatis saat project dibuat dan langsung dipasang di layar work type terkait.
  6. Field yang dibuat dari List ("Create field") atau dari form builder (BIZ-030) otomatis ditambahkan ke konfigurasi work type terkait (team-managed) atau ke field configuration project (company-managed, bila pengguna berhak).
  7. ⚠ Jumlah field custom per project team-managed dibatasi oleh nilai konfigurasi (default 50); melewati batas menampilkan pesan galat.
- **Team-managed vs company-managed:** Team-managed: layout per work type project-scoped; company-managed: field configuration dan screen scheme global.
- **Catatan & ketergantungan:** Field, screen, dan layout di [05](05-konfigurasi-work-item-skema.md); panel unscheduled di PLAN-023 ([18](18-timeline-rencana-tampilan-proyek.md)).

### BIZ-021 Work categories (Category field)

- **Deskripsi:** Category adalah field label berwarna di level project untuk mengelompokkan pekerjaan dan memberi kode warna di semua view bisnis.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Anggota project (memilih/membuat Category saat mengedit work item, butuh Edit work items), Admin Proyek (mengelola).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap project memiliki satu field single-select "Category" dengan nilai berupa nama (unik per project, maks. 255 karakter) dan warna.
  2. Category baru dapat dibuat langsung dari issue view, List, atau Board dengan mengetik nama yang belum ada lalu memilih "Create"; warna diberikan otomatis dari palet dan dapat diubah.
  3. Halaman kelola Category memungkinkan rename, ganti warna (palet terbatas, mis. 10-12 warna), dan hapus.
  4. ⚠ Menghapus Category mengosongkan nilai Category di semua work item terkait setelah dialog konfirmasi yang menyebut jumlah work item terdampak; operasi tercatat di history work item.
  5. Warna Category tampil di kartu Board, bar Calendar, dan bar Timeline (sebagai warna bar atau penanda) secara konsisten.
  6. Filter berdasarkan Category tersedia di List, Board, Calendar, dan Timeline, termasuk nilai "None".
  7. ⚠ Category dapat dicari via JQL (mis. `category = "Marketing"`) dan muncul di autocomplete JQL.
  8. Category project berbeda dari Project category global yang dikelola Admin Situs untuk mengelompokkan project; UI dan API memakai nama yang berbeda agar tidak tertukar.
  9. ⚠ Jumlah Category per project dibatasi oleh nilai konfigurasi; melewati batas menampilkan pesan galat yang menyebut batasnya.
- **Team-managed vs company-managed:** Ketersediaan di team-managed belum terkonfirmasi di Jira Cloud; clone menyediakan Category di kedua mode dengan perilaku sama.
- **Catatan & ketergantungan:** Definisi field di CFG-021 ([05](05-konfigurasi-work-item-skema.md)); JQL di JQL-034 ([12](12-jql-bahasa-kueri.md)); Project category global di PRJ-019 ([04](04-proyek.md)).

## Navigasi & view pendukung project bisnis

### BIZ-022 Business project navigation & view tabs

- **Deskripsi:** Navigasi tab project business untuk berpindah antar view dan mengatur view mana yang tampil.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Anggota project (memakai tab), Admin Proyek (menambah, mengurutkan, menyembunyikan tab).
- **Aturan perilaku (acceptance criteria):**
  1. Tab yang tersedia: Summary, List, Board, Calendar, Timeline, Forms, Pages, Attachments, Reports (company-managed saja), dan Shortcuts; hanya tab yang aktif dan didukung plan yang ditampilkan.
  2. Tombol "+" menampilkan daftar view yang tersembunyi atau belum ditambahkan dan menambahkannya ke deret tab.
  3. Tab dapat diurutkan ulang dengan drag-drop; tab yang tidak muat lebar layar dipindahkan ke menu "More" dengan urutan yang sama.
  4. ⚠ Tab dapat disembunyikan; di team-managed penyembunyian dilakukan lewat Features (BIZ-004) dan berlaku per project, sedangkan urutan tab yang diubah lewat drag-drop disimpan per project oleh Admin Proyek.
  5. Header project menampilkan nama dan avatar project, tombol Share (salin tautan view), Star, dan menu "..." yang memuat Project settings serta aksi lain sesuai izin.
  6. Setiap view memiliki deep link: `/projects/{KEY}/summary`, `/list`, `/board`, `/calendar`, `/timeline`, dan `/form`; membuka deep link view yang dinonaktifkan menampilkan pesan dan tautan ke view default.
  7. ⚠ Sidebar navigasi menampilkan project yang di-star/recent beserta daftar view-nya sebagai sub-item yang dapat diklik.
- **Team-managed vs company-managed:** Reports hanya di company-managed; penyembunyian tab team-managed lewat Features, company-managed lewat PRJ-050.
- **Catatan & ketergantungan:** Navigasi global dan sidebar di NAV-002/NAV-003 ([21](21-navigasi-ux-preferensi.md)); view List/Calendar/Timeline/Summary/Board di PLAN-015, PLAN-016, PLAN-022, PLAN-025, PLAN-026 ([18](18-timeline-rencana-tampilan-proyek.md)). Tab "Approvals" tersendiri tidak dibangun karena keberadaannya tidak terkonfirmasi. SPA perlu route bertingkat di `frontend/src/routes`; fallback `index.html` dari `webui` sudah mendukung deep link.

### BIZ-023 Get started checklist

- **Deskripsi:** Checklist onboarding yang memandu tim baru menyiapkan dan memakai project business.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, pengguna baru di project.
- **Aturan perilaku (acceptance criteria):**
  1. Checklist muncul otomatis sebagai panel setelah project baru (atau site baru) dibuat, bagi pembuat project dan anggota yang pertama kali membuka project.
  2. ⚠ Item contoh: Create work item, Invite team, Try a view (Calendar/Timeline), Create a form, Customize workflow/fields.
  3. Panel menampilkan progress bar; langkah tercentang otomatis saat aksi terkait berhasil dilakukan (dideteksi dari event domain, mis. work item pertama dibuat).
  4. Setiap item memiliki tutorial singkat dan tautan "Learn more" ke dokumentasi.
  5. Panel dapat di-minimize, di-dismiss, dan dibuka kembali dari menu Help.
  6. Isi checklist ditentukan oleh template project (mis. template dengan form menyertakan "Share your form").
  7. Status checklist (item selesai, dismissed, minimized) disimpan per pengguna per project di server, bukan di browser storage.
- **Team-managed vs company-managed:** Tidak ada perbedaan; item yang membutuhkan izin admin hanya ditampilkan ke admin.
- **Catatan & ketergantungan:** Onboarding umum di PRJ-015 ([04](04-proyek.md)) dan NAV-023 ([21](21-navigasi-ux-preferensi.md)).

### BIZ-024 Shortcuts

- **Deskripsi:** Tautan pintasan di navigasi project ke resource internal atau eksternal.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola), anggota project (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog "Add shortcut" meminta URL (wajib, divalidasi skema `http`/`https`, pesan galat "URL tidak valid") dan nama tampilan (wajib, maks. 255 karakter).
  2. Shortcut terlihat oleh semua pengguna yang dapat browse project, di tab Shortcuts dan/atau di sidebar project.
  3. ⚠ Shortcut dapat diedit dan dihapus oleh Admin Proyek; anggota tanpa hak admin hanya dapat membuka tautan.
  4. Ikon/favicon situs tujuan ditampilkan otomatis bila dapat diambil; bila gagal, ikon tautan generik dipakai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur umum shortcut di PRJ-029 ([04](04-proyek.md)); pengambilan favicon dilakukan server-side dengan timeout dan proteksi SSRF (lihat modul 27).

### BIZ-025 Attachments view

- **Deskripsi:** View level project yang menampilkan semua attachment dari work item di project.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Anggota project dengan Browse projects; hapus butuh permission Delete attachments.
- **Aturan perilaku (acceptance criteria):**
  1. View menyediakan tampilan grid/thumbnail dan list yang dapat dialihkan, dengan pilihan tersimpan per pengguna.
  2. Setiap attachment menampilkan nama file, tipe, ukuran, uploader, tanggal upload, dan work item asal (key + summary).
  3. ⚠ Filter berdasarkan tipe file (gambar, dokumen, PDF, lainnya) dan sort berdasarkan tanggal atau nama tersedia.
  4. Aksi per attachment: preview di media viewer, download, dan buka work item asal.
  5. Hapus attachment hanya tampil bagi pengguna dengan permission Delete attachments (own/all) dan tercatat di history work item asal.
  6. Attachment dari work item yang tidak dapat dilihat pengguna karena issue security level tidak ditampilkan dan tidak dihitung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Attachment dan media viewer di CMT-011 s.d. CMT-018 ([08](08-komentar-lampiran-aktivitas.md)); permission di PERM-031 ([03](03-izin-peran-keamanan.md)).

### BIZ-026 Pages (Confluence) tab

- **Deskripsi:** Tab Pages menghubungkan project dengan space dokumentasi (di Jira asli: Confluence) untuk dokumentasi project.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek (menghubungkan/memutus), anggota project (melihat dan membuat halaman sesuai izin di sistem dokumentasi).
- **Aturan perilaku (acceptance criteria):**
  1. Admin Proyek dapat menghubungkan space dokumentasi yang sudah ada atau membuat space baru dari tab Pages.
  2. Page tree space yang terhubung ditampilkan di dalam tab Pages dengan navigasi hierarki.
  3. Pengguna dapat membuat halaman baru dari template dokumentasi (mis. meeting notes, project plan).
  4. Pengguna dapat menautkan halaman tertentu (bukan seluruh space) ke project.
  5. Bila belum ada sistem dokumentasi terhubung, tab menampilkan prompt untuk menambahkan atau mengonfigurasi integrasi.
  6. Izin melihat dan mengedit halaman mengikuti sistem dokumentasi; halaman yang tidak dapat diakses pengguna tidak ditampilkan.
  7. Admin Proyek dapat mengganti space yang terhubung atau memutus koneksi tanpa menghapus halaman di sistem dokumentasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Membutuhkan produk dokumentasi eksternal; clone dapat memakai modul dokumen internal di masa depan atau sekadar daftar tautan eksternal. Fitur Pages umum di PRJ-030 ([04](04-proyek.md)) dan CMT-030 ([08](08-komentar-lampiran-aktivitas.md)). Integrasi Confluence sebagai produk Atlassian berada di luar cakupan.

### BIZ-027 Reports (business project)

- **Deskripsi:** Laporan dasar yang tersedia di project business.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, manager, anggota project dengan Browse projects.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Company-managed menyediakan laporan standar Created vs resolved, Pie chart, Average age, Resolution time, Recently created, Time since, dan User workload.
  2. Laporan agile (burndown, velocity, sprint report) tidak ditampilkan karena project business tidak memiliki sprint.
  3. Team-managed tidak memiliki tab Reports; Summary dipakai sebagai pengganti laporan.
  4. Halaman Reports menyediakan tautan ke Dashboard untuk laporan kustom (gadget).
- **Team-managed vs company-managed:** Reports hanya di company-managed; team-managed memakai Summary.
- **Catatan & ketergantungan:** Detail laporan di RPT-034/RPT-035 ([16](16-laporan-insight.md)); dashboard di [17](17-dashboard-gadget.md).

### BIZ-028 Project background/theme

- **Deskripsi:** Personalisasi visual project (space) dengan warna solid, gradient, atau gambar latar; berlaku untuk semua jenis space, baik business maupun software.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek (mengubah), anggota project (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Pemilih background menyediakan warna solid, gradient, atau gambar dari galeri bawaan.
  2. ⚠ Background tampil di view Board, List, Calendar, dan Timeline project; kontras teks tetap memenuhi WCAG 2.1 AA (overlay otomatis bila perlu).
  3. Background diatur Admin Proyek dan berlaku untuk semua anggota project.
  4. Opsi "Reset to default" mengembalikan tampilan tanpa background.
  5. Fitur tersedia untuk project business dan project software (team-managed maupun company-managed); pemilih background dibuka dari pengaturan Details project di samping Project Avatar (PRJ-018).
  6. ⚠ Admin Proyek dapat mengunggah gambar kustom sebagai background; hanya tipe PNG, JPEG, GIF, atau WebP yang diterima dan ukuran berkas dibatasi nilai konfigurasi (default 10 MB). Berkas dengan tipe atau ukuran tidak valid ditolak dengan pesan galat yang menyebut tipe yang diizinkan atau batas ukuran, dan background lama tetap dipakai.
  7. Gambar kustom disimpan lewat layanan penyimpanan berkas yang sama dengan lampiran dan avatar (CMT-019); mengganti atau me-reset background menghapus referensi gambar lama dari project.
  8. ⚠ Background tetap terbaca pada tema light, dark, dan system (NAV-034); overlay penyesuaian kontras mengikuti tema aktif pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dikonfirmasi oleh rilis Jira Cloud Februari 2026 "Customize your Jira spaces" untuk semua space, dengan gambar kustom, warna solid, dan gradient. Project Avatar di PRJ-018 ([04](04-proyek.md)); penyimpanan berkas di CMT-019 ([08](08-komentar-lampiran-aktivitas.md)); tema di NAV-034 dan aksesibilitas di NAV-036 ([21](21-navigasi-ux-preferensi.md)).

## Formulir (Forms)

### Pembuatan & konfigurasi form

### BIZ-029 Forms tab & forms directory

- **Deskripsi:** Tab Forms menampilkan semua form milik project dan menjadi titik masuk untuk membuat, mengedit, dan membagikan form.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola form), anggota project (melihat daftar dan membuka form untuk diisi).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar form menampilkan kolom nama form, work type tujuan, pembuat, dan tanggal update terakhir, dapat diurutkan per kolom.
  2. Tombol "Create form" tampil hanya bagi pengguna yang berhak membuat form (BIZ-034) dan membuka form builder dengan form kosong.
  3. ⚠ Aksi per form: Open (mengisi), Edit, Copy link, Duplicate, dan Delete; aksi yang tidak diizinkan disembunyikan.
  4. Bila belum ada form, tab menampilkan empty state dengan penjelasan singkat dan ajakan "Create your first form".
  5. Delete form meminta konfirmasi yang menyatakan bahwa work item yang sudah dibuat dari form tidak ikut terhapus; setelah dihapus, tautan form menampilkan "Form tidak ditemukan".
  6. Tab Forms dapat disembunyikan lewat Features di team-managed (BIZ-004) tanpa menghapus form.
  7. ⚠ Jumlah form per project dibatasi oleh nilai konfigurasi; tombol "Create form" dinonaktifkan dengan tooltip yang menyebut batas saat batas tercapai.
- **Team-managed vs company-managed:** Tidak ada perbedaan; penyembunyian tab company-managed mengikuti BIZ-022.
- **Catatan & ketergantungan:** Forms awalnya hanya di project business, kini juga di project software (BIZ-041). JSM forms (request type forms) dirinci di JSM-PRJ-034 ([29](29-jsm-service-project-request-type-formulir.md)) dan bukan bagian dari fitur ini.

### BIZ-030 Form builder

- **Deskripsi:** Editor visual untuk menyusun form intake (judul, deskripsi, dan pertanyaan) yang dipetakan ke field work item.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek atau pembuat form yang berhak (BIZ-034).
- **Aturan perilaku (acceptance criteria):**
  1. Pembuat form dapat mengedit judul form (wajib, maks. 255 karakter) dan deskripsi (rich text sederhana).
  2. Pembuat form wajib memilih satu work type tujuan; panel field hanya menampilkan field yang tersedia di konfigurasi work type tersebut.
  3. Panel field memungkinkan menambah field project ke form dengan klik atau drag-drop; field yang sudah dipakai ditandai dan tidak dapat ditambahkan dua kali.
  4. Urutan pertanyaan dapat diubah dengan drag-drop dan dengan keyboard (aksesibel).
  5. Menghapus pertanyaan dari form tidak menghapus field dari work type maupun nilainya di work item lama.
  6. Pertanyaan Summary selalu ada di form, tidak dapat dihapus, dan ditandai wajib.
  7. ⚠ "Create new field" dari builder mendukung tipe text, paragraph, number, date, dropdown, checkbox, user, dan labels; field baru otomatis ditambahkan ke work type tujuan (BIZ-020).
  8. ⚠ Perubahan form disimpan otomatis (autosave) sebagai draft dan baru berlaku untuk submitter setelah pembuat menekan Save/Publish; indikator status simpan ditampilkan.
  9. Mode Preview menampilkan form persis seperti dilihat submitter tanpa membuat work item.
  10. ⚠ Field yang tidak didukung form (mis. Sprint, time tracking, Parent bila tidak relevan) tidak muncul di panel field.
  11. ⚠ Pembuat form dapat menambahkan section/heading untuk mengelompokkan pertanyaan; section tidak dipetakan ke field.
  12. ⚠ Pembuat form dapat memulai dari template form siap pakai (mis. "Bug report", "Design request", "Purchase request") yang mengisi pertanyaan awal.
- **Team-managed vs company-managed:** Team-managed: "Create new field" langsung menambah field project-scoped; company-managed: hanya Admin Situs yang dapat membuat custom field global, sehingga opsi ini disembunyikan bagi Admin Proyek.
- **Catatan & ketergantungan:** Section dan template form pasti ada di JSM forms (JSM-PRJ-035, [29](29-jsm-service-project-request-type-formulir.md)); untuk kerja bisnis belum terkonfirmasi sehingga ditandai. Tampilan/tema form tidak dijadikan fitur terpisah. Model data form disarankan: `forms` (project_id, work_type_id, title, description, status, version) dan `form_questions` (form_id, field_id, label, help_text, required, position, section, condition).

### BIZ-031 Form field mapping to Jira fields

- **Deskripsi:** Setiap pertanyaan form dipetakan 1:1 ke field work item, sehingga jawaban menjadi nilai field pada work item yang dibuat.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pembuat form, Admin Situs (field configuration company-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap pertanyaan dipetakan ke tepat satu system field atau custom field, dan satu field hanya dapat dipetakan ke satu pertanyaan per form.
  2. ⚠ System field yang dapat dipakai: Summary, Description, Attachment, Due date, Start date, Priority, Labels, Assignee, dan Category.
  3. ⚠ Tipe custom field yang dapat dipakai: text single-line dan multi-line, number, date, date-time, select single/multi, checkbox, radio, user picker, dan URL.
  4. Opsi pada pertanyaan select/checkbox/radio selalu mengikuti option field (termasuk urutan dan opsi nonaktif yang disembunyikan); perubahan option field langsung tercermin di form.
  5. Bila pertanyaan tidak diisi, default value field berlaku pada work item yang dibuat.
  6. Company-managed: field hanya dapat dipakai bila ada di screen create dan tidak hidden di field configuration work type tujuan; field yang hidden tidak muncul di panel field.
  7. Bila field dihapus dari project/work type setelah dipakai di form, builder menampilkan peringatan "Pertanyaan ini tidak valid" pada pertanyaan terkait dan form tidak dapat dipublikasikan ulang sebelum diperbaiki; form yang sudah terbit mengabaikan pertanyaan tidak valid saat submit.
  8. Mengganti work type tujuan melepas pertanyaan yang field-nya tidak ada di work type baru, setelah dialog konfirmasi yang mendaftar pertanyaan yang akan dilepas.
- **Team-managed vs company-managed:** Company-managed tunduk pada screen dan field configuration; team-managed mengikuti layout field work type.
- **Catatan & ketergantungan:** Field dan field configuration di [05](05-konfigurasi-work-item-skema.md).

### BIZ-032 Form field properties (label, help text, required)

- **Deskripsi:** Pengaturan tampilan dan validasi per pertanyaan, terpisah dari konfigurasi field.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pembuat form.
- **Aturan perilaku (acceptance criteria):**
  1. Label pertanyaan dapat berbeda dari nama field (mis. field "Due date" berlabel "Kapan Anda membutuhkannya?"); label kosong memakai nama field.
  2. Setiap pertanyaan dapat memiliki help text/deskripsi yang tampil di bawah label.
  3. Toggle Required hanya berlaku di form dan tidak mengubah field configuration maupun layar work item.
  4. Field yang required di field configuration selalu wajib di form dan toggle Required-nya terkunci aktif dengan tooltip penjelasan.
  5. Pertanyaan Summary selalu required.
  6. Validasi tipe dilakukan di sisi klien (number, date, URL) dan diulang di server; input tidak valid menampilkan pesan per field.
  7. Batas tipe dan ukuran attachment pada pertanyaan Attachment mengikuti setting attachment site (ADM-011); file yang melanggar ditolak dengan pesan yang menyebut batasnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber aturan required (field configuration di company-managed, layout field di team-managed).
- **Catatan & ketergantungan:** Setting attachment di ADM-011 ([22](22-administrasi-sistem.md)) dan CMT-016 ([08](08-komentar-lampiran-aktivitas.md)).

### BIZ-033 Form conditional logic

- **Deskripsi:** Aturan untuk menampilkan atau menyembunyikan pertanyaan berdasarkan jawaban pertanyaan lain.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Pembuat form.
- **Aturan perilaku (acceptance criteria):**
  1. Pembuat form dapat menambahkan aturan "show when [field] is [value]" pada pertanyaan, dengan sumber kondisi berupa pertanyaan bertipe select, checkbox, atau radio yang berada di atasnya.
  2. Pertanyaan yang sedang tersembunyi tidak divalidasi required, baik di klien maupun di server.
  3. Nilai pertanyaan yang tersembunyi tidak dikirim dan tidak ditulis ke work item, walau sebelumnya sempat diisi.
  4. Builder menampilkan indikator (ikon dan ringkasan kondisi) pada pertanyaan yang memiliki kondisi, dan mencegah kondisi melingkar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Conditional logic pasti ada di JSM forms (JSM-PRJ-037, [29](29-jsm-service-project-request-type-formulir.md)), tetapi ketersediaannya di form kerja bisnis belum terkonfirmasi; jangan dijadikan target MVP sebelum keputusan produk (lihat Pertanyaan terbuka). Implementasi sebaiknya berbagi engine kondisi dengan JSM forms.

### BIZ-034 Form access & permissions

- **Deskripsi:** Aturan siapa yang boleh membuat, mengedit, dan mengisi form.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs, submitter form.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Membuat, mengedit, dan menghapus form hanya untuk Admin Proyek (team-managed: role Administrator; company-managed: permission Administer projects); opsi agar role Member juga dapat mengelola form disediakan sebagai setting project yang default-nya nonaktif.
  2. Submit form membutuhkan permission Browse projects dan Create work items pada project; server memeriksa ulang izin saat submit.
  3. Role Viewer di team-managed tidak dapat submit karena tidak memiliki izin create; form menampilkan pesan "Anda tidak memiliki izin untuk mengirim form ini".
  4. Akses anonim (tanpa login) ke form internal tidak didukung; pengguna yang belum login diarahkan ke halaman login lalu kembali ke form.
- **Team-managed vs company-managed:** Team-managed memakai role Administrator/Member/Viewer; company-managed memakai permission scheme.
- **Catatan & ketergantungan:** Permission di [03](03-izin-peran-keamanan.md); form publik di BIZ-036.

### Berbagi & pengisian form

### BIZ-035 Form sharing (internal link & embed)

- **Deskripsi:** Membagikan form lewat tautan kepada pengguna site agar mereka dapat mengajukan permintaan.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pembuat form (membagikan), submitter form (membuka tautan).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Copy link" menghasilkan URL unik dan stabil berpola `/projects/{KEY}/form/{id}` dan menampilkan toast "Link copied".
  2. Penerima tautan harus login dan memiliki permission Browse projects serta Create work items untuk mengisi form.
  3. Pada project team-managed Open, semua pengguna dengan product access dapat submit; pada Limited/Private, kemampuan submit mengikuti role anggota.
  4. ⚠ Pengguna tanpa akses melihat halaman "no access" dengan tombol "Request access" yang mengirim permintaan ke Admin Proyek (USR-013).
  5. Tautan form tampil sebagai smart link (judul form dan nama project) saat ditempel di editor produk ini dan di aplikasi yang membaca metadata Open Graph/oEmbed.
  6. ⚠ Tidak ada kode embed iframe resmi untuk situs eksternal; embed terbatas pada smart link/embed card di dalam produk ini.
- **Team-managed vs company-managed:** Team-managed mengikuti access level; company-managed mengikuti permission scheme.
- **Catatan & ketergantungan:** Smart link di EDT-021 ([07](07-editor-konten-kaya.md)) dan NAV-017 ([21](21-navigasi-ux-preferensi.md)); access request di USR-013 ([02](02-pengguna-grup-organisasi.md)). Unfurl di Slack/Microsoft Teams bergantung pada integrasi modul 25.

### BIZ-036 Public/external form sharing

- **Deskripsi:** Membagikan form kepada orang di luar site (tanpa akun) melalui tautan publik.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Pembuat form, submitter eksternal (Tamu/Kolaborator Eksternal), Admin Situs (mengizinkan fitur di tingkat site).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Setting akses form memiliki dua pilihan: "Only people with access to this project" (default) dan "Anyone with the link"; opsi kedua hanya tampil bila Admin Situs mengizinkan form publik.
  2. ⚠ Untuk submitter eksternal, Reporter diisi dengan akun default yang dikonfigurasi per form (atau pemilik form bila tidak diatur), dan email submitter (bila diminta) disimpan sebagai nilai field.
  3. Admin Situs dapat mengaktifkan/menonaktifkan form publik di tingkat site, dan pembuat form dapat mencabut tautan publik (membuat token baru) sehingga tautan lama langsung tidak berlaku.
  4. ⚠ Submission publik dilindungi dari spam dengan rate limit per IP dan tantangan anti-bot; submission yang gagal verifikasi ditolak dengan pesan umum.
  5. ⚠ Pertanyaan bertipe user picker dan field yang mengekspos data internal tidak dapat dipakai pada form publik; builder menampilkan peringatan saat form diubah menjadi publik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Secara historis form kerja bisnis hanya untuk pengguna yang login; untuk intake eksternal, saluran utama adalah portal dan forms Service Management ([29](29-jsm-service-project-request-type-formulir.md) dan [30](30-jsm-help-center-portal-pelanggan-kanal.md)). Jangan diimplementasikan sebelum keputusan produk (lihat Pertanyaan terbuka). Akses anonim umum di PERM-046 ([03](03-izin-peran-keamanan.md)).

### BIZ-037 Form submission experience

- **Deskripsi:** Pengalaman submitter saat mengisi dan mengirim form.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Submitter form (Browse projects + Create work items).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman form adalah halaman fokus (tanpa navigasi project penuh) dengan header berisi judul form, deskripsi, dan nama project.
  2. Validasi required dan format dijalankan sebelum submit; pesan galat ditampilkan per field dan fokus berpindah ke field pertama yang galat.
  3. Pertanyaan Attachment mendukung upload dengan drag-drop dan pemilih file, menampilkan progres upload, dan tombol submit dinonaktifkan selama upload berlangsung.
  4. Setelah submit berhasil, halaman menampilkan pesan sukses, tautan ke work item yang dibuat (hanya bila submitter dapat melihatnya), dan tombol "Submit another response".
  5. ⚠ Isian form tidak disimpan sebagai draft; menutup halaman dengan isian belum terkirim memunculkan konfirmasi browser "perubahan belum disimpan".
  6. Bila form atau field berubah ketika submitter sedang mengisi (versi form berbeda saat submit), server menolak dengan pesan "Form telah diperbarui, muat ulang halaman" dan isian yang masih valid dipertahankan setelah muat ulang bila memungkinkan.
  7. Layout responsif untuk desktop dan mobile (lebar minimum 320px) tanpa scroll horizontal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Upload memakai layanan attachment CMT-011/CMT-019 ([08](08-komentar-lampiran-aktivitas.md)); NFR responsif dan aksesibilitas di modul 27 dan NAV-036 ([21](21-navigasi-ux-preferensi.md)).

### BIZ-038 Form submissions → work items

- **Deskripsi:** Setiap submission form langsung membuat work item baru di project dengan nilai field dari jawaban.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Submitter form, anggota project (menindaklanjuti).
- **Aturan perilaku (acceptance criteria):**
  1. Submission membuat satu work item dengan work type tujuan form, di status awal workflow work type tersebut, dalam satu transaksi bersama upload attachment yang sudah selesai.
  2. Reporter diisi submitter; Assignee diisi default assignee project bila pertanyaan Assignee tidak ada atau dikosongkan.
  3. Submitter otomatis ditambahkan sebagai watcher dan menerima notifikasi sesuai notification scheme (event Work item created).
  4. Trigger automation "Work item created" (dan "Form submitted" bila didukung) berjalan untuk work item tersebut.
  5. Pada company-managed, validator create transition dan field required tetap berlaku; bila gagal, submission ditolak dengan pesan galat yang dapat dipahami submitter dan tidak ada work item parsial yang tercipta.
  6. Default issue security level project diterapkan pada work item yang dibuat.
  7. ⚠ Activity/history work item mencatat sumber pembuatan "Created via form: <nama form>" dan menyimpan `form_id` sebagai properti work item.
  8. ⚠ Anggota project dapat melihat semua submission per form lewat filter (mis. tautan "View submissions" yang membuka List dengan JQL berbasis properti form).
- **Team-managed vs company-managed:** Validator create transition hanya ada di company-managed; team-managed hanya menerapkan field required layout.
- **Catatan & ketergantungan:** Create work item di [06](06-work-item-inti.md); watcher di ISS-036 dan NTF-012 ([19](19-notifikasi-email.md)); trigger di AUTO-029 ([20](20-otomasi.md)); validator di WF-028 ([09](09-workflow-status-transisi.md)); event domain di modul 28.

### Siklus hidup & pelaporan form

### BIZ-039 Forms management & lifecycle

- **Deskripsi:** Pengelolaan form setelah dibuat (edit, duplikat, tutup, hapus) dan dampaknya ke submission.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, pembuat form yang berhak (BIZ-034).
- **Aturan perilaku (acceptance criteria):**
  1. Form yang sudah dibagikan dapat diedit; setelah disimpan/dipublikasikan, perubahan langsung berlaku di tautan yang sama tanpa membuat tautan baru.
  2. ⚠ Duplicate form membuat salinan dengan nama "Copy of <nama>", pertanyaan dan properti yang sama, serta tautan baru.
  3. Delete form menghapus form beserta tautannya; work item yang sudah dibuat dari form tetap ada dan penanda sumbernya tetap terbaca.
  4. Work type tujuan dapat diganti dengan perilaku pelepasan pertanyaan sesuai BIZ-031.
  5. ⚠ Form dapat ditutup ("Close form") agar berhenti menerima submission; tautan menampilkan pesan "Form ini tidak lagi menerima respons" dan form dapat dibuka kembali.
  6. Form ikut diarsipkan saat project diarsipkan, ikut masuk trash saat project dihapus, dan ikut dipulihkan saat project direstore.
  7. Saat work item dipindah ke project lain, keterkaitan dengan form asal tidak ikut berpindah (form tetap milik project asal) tetapi catatan sumber di history tetap ada.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lifecycle project di PRJ-032/PRJ-033 ([04](04-proyek.md)); move work item di ISS-032 ([06](06-work-item-inti.md)).

### BIZ-040 Form responses export & reporting

- **Deskripsi:** Jawaban form disimpan sebagai nilai field work item, sehingga ekspor dan analisis memakai mekanisme pencarian dan ekspor standar.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, manager, anggota project dengan Browse projects.
- **Aturan perilaku (acceptance criteria):**
  1. Ekspor jawaban dilakukan lewat issue navigator/List dengan opsi CSV (all fields) atau CSV (current fields).
  2. Submission dapat difilter dengan JQL pada work type, field, atau properti form (BIZ-038).
  3. Ringkasan jawaban tersedia lewat Summary project atau gadget Dashboard (mis. Pie chart per nilai field).
  4. Tidak tersedia "Export form data" XLSX khusus per form seperti di JSM forms; UI tidak menampilkan opsi tersebut di project bisnis.
  5. Ekspor menghormati permission project dan issue security; work item yang tidak dapat dilihat pengguna tidak ikut diekspor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ekspor di [13](13-pencarian-filter-ekspor.md); gadget di [17](17-dashboard-gadget.md); ekspor XLSX JSM forms di JSM-PRJ-034 ([29](29-jsm-service-project-request-type-formulir.md)).

### BIZ-041 Forms in software projects

- **Deskripsi:** Setelah penggabungan produk, forms juga dapat dipakai di project software.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek project software, developer, submitter form.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tab Forms dapat ditambahkan di project software team-managed dan company-managed lewat tombol "+" atau Features.
  2. Perilaku form builder, sharing, submission, dan lifecycle di project software identik dengan project business (BIZ-029 s.d. BIZ-039).
  3. Use case utama di project software adalah intake bug report dan feature request; bila template form (BIZ-030) tersedia, template "Bug report" dan "Feature request" ditawarkan sebagai titik awal.
  4. Field yang tersedia di panel field mengikuti work type software tujuan (mis. Bug, Story), termasuk field custom-nya, dengan pengecualian field yang tidak didukung (BIZ-030).
- **Team-managed vs company-managed:** Tidak ada perbedaan selain mekanisme menampilkan tab (Features team-managed vs PRJ-050 company-managed).
- **Catatan & ketergantungan:** Cakupan per tier belum terkonfirmasi. Fitur Issue collector software (bila ada) dibahas di [04](04-proyek.md).

## Persetujuan (Approvals) di workflow

### BIZ-042 Approvals in business workflows

- **Alias inventaris:** Approvals in workflow (Jira, non-JSM)
- **Deskripsi:** Langkah persetujuan bawaan pada status workflow (non-JSM): work item yang masuk ke status tersebut menunggu keputusan approver dan baru berpindah ke status berikutnya setelah disetujui atau ditolak.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (workflow company-managed), Admin Proyek (workflow team-managed bila didukung), approver, reporter/pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Langkah approval dikonfigurasi di workflow editor baru dengan memilih status lalu "Add approval".
  2. Setiap status hanya dapat memiliki satu langkah approval; work item di status itu berada dalam keadaan "menunggu keputusan" sampai keputusan diambil.
  3. Konfigurasi langkah wajib menentukan transition/status tujuan untuk hasil Approved dan untuk hasil Declined.
  4. Editor menolak menyimpan langkah approval bila belum ada transition keluar dari status approval ke status tujuan Approved dan Declined, dengan pesan galat yang menyebut transition yang kurang.
  5. ⚠ Selama approval berjalan, transition lain dari status approval dibatasi sehingga work item hanya dapat keluar lewat hasil keputusan.
  6. ⚠ Mengubah atau menghapus langkah approval saat ada work item pending menampilkan jumlah work item terdampak; setelah penghapusan, work item pending dilepas dari keadaan menunggu dan transition normal kembali tersedia, sedangkan perubahan approver/aturan hanya berlaku untuk langkah approval berikutnya.
  7. ⚠ Langkah approval didukung pertama untuk project business company-managed; dukungan untuk team-managed dan project software diaktifkan lewat feature flag setelah keputusan produk.
  8. Pada diagram workflow, status yang memiliki langkah approval ditandai badge "Approval" dan panel status menampilkan ringkasan konfigurasi (field approver, jumlah persetujuan, tujuan Approved/Declined).
  9. Approver ditentukan dari field user picker (single atau multi) pada work item; konfigurasi detail diatur di BIZ-043.
  10. ⚠ Langkah approval menyimpan aturan jumlah persetujuan yang dibutuhkan: satu approver saja atau semua approver.
  11. Transition hasil Approved dan Declined dijalankan oleh sistem atas nama approver; transition tersebut tidak tampil sebagai tombol transisi manual bagi pengguna.
  12. Issue view menampilkan panel approval dengan tombol Approve/Decline bagi approver (BIZ-044).
  13. Setiap keputusan approval tercatat di activity/history work item (BIZ-046).
  14. Saat work item masuk status approval, approver menerima notifikasi (BIZ-045).
  15. ⚠ Selama approval masih pending, tombol transisi manual lain dari status approval dinonaktifkan dengan tooltip "Menunggu persetujuan", dan API transition menolak permintaan tersebut dengan HTTP 409.
  16. Approvals Service Management (approver group, jumlah approval per request type, auto-decline, approve via email, JQL `approver()`/`pendingBy()`) tidak termasuk fitur ini dan dirinci terpisah di JSM-ITSM-011 s.d. JSM-ITSM-015 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)).
- **Team-managed vs company-managed:** Utama untuk company-managed (workflow editor baru); team-managed belum terkonfirmasi dan berada di balik feature flag.
- **Catatan & ketergantungan:** Dua entri inventaris digabung: domain workflow (prioritas P2, tier belum pasti) dan domain kerja bisnis (P1); prioritas gabungan memakai yang lebih tinggi. Workflow editor di WF-021 ([09](09-workflow-status-transisi.md)); tab Approvals di activity view di CMT-025 ([08](08-komentar-lampiran-aktivitas.md)); JSM approvals di JSM-ITSM-011 s.d. JSM-ITSM-015 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)) dan JSM-PRT-031 ([30](30-jsm-help-center-portal-pelanggan-kanal.md)). Model data disarankan: `workflow_status_approvals` (workflow_id, status_id, approver_field_id, rule `any|all`, approved_transition_id, declined_transition_id) dan `work_item_approvals` (work_item_id, step_id, status, created_at, completed_at) serta `work_item_approval_decisions` (approval_id, approver_id, decision, decided_at).

### BIZ-043 Approvers configuration

- **Deskripsi:** Cara menentukan siapa yang harus menyetujui pada langkah approval.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek/Admin Situs (konfigurasi), reporter (mengisi approver), approver.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi langkah approval memilih satu field user picker (single atau multi) pada work item, mis. field "Approvers", sebagai sumber daftar approver; hanya field user picker yang ada di work type terkait yang dapat dipilih.
  2. ⚠ Aturan jumlah persetujuan dapat dipilih: "Any one approver" (cukup satu persetujuan) atau "All approvers" (semua harus setuju); dengan "All", satu Decline langsung menghasilkan Declined.
  3. Approver dapat diisi saat create (termasuk lewat form atau automation) atau kapan pun sebelum work item masuk status approval; daftar approver di-snapshot saat work item masuk status approval.
  4. ⚠ Bila field approver kosong saat work item masuk status approval, panel approval menampilkan peringatan "Belum ada approver" dan approval tidak dapat diproses sampai field diisi.
  5. ⚠ Approver dapat diganti selama approval pending oleh pengguna yang berhak mengedit field approver; perubahan memperbarui daftar approver aktif, keputusan yang sudah diberikan approver yang tetap ada dipertahankan, dan perubahan tercatat di history.
  6. Approver harus memiliki permission Browse projects; pengguna tanpa izin tersebut tidak dapat dipilih (atau ditandai "Tidak dapat menyetujui: tidak memiliki akses") dan tidak dihitung dalam aturan "All".
- **Team-managed vs company-managed:** Tidak ada perbedaan pada aturan; field approver team-managed bersifat project-scoped.
- **Catatan & ketergantungan:** Field user picker di [05](05-konfigurasi-work-item-skema.md); pengisian approver otomatis lewat automation di BIZ-048 dan [20](20-otomasi.md).

### BIZ-044 Approval decision (issue view)

- **Deskripsi:** Panel approval di issue view untuk menyetujui atau menolak work item.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Approver yang ditunjuk (Approve/Decline), anggota project (melihat status).
- **Aturan perilaku (acceptance criteria):**
  1. Panel approval menampilkan setiap approver dengan avatar dan status Pending, Approved, atau Declined, serta aturan jumlah persetujuan.
  2. Tombol Approve dan Decline hanya muncul bagi approver yang ditunjuk dan belum memberi keputusan; server menolak keputusan dari non-approver dengan HTTP 403.
  3. Setelah aturan jumlah persetujuan terpenuhi (atau satu Decline), work item otomatis ditransisikan ke status tujuan Approved/Declined; bila transition gagal (mis. validator), keputusan tetap tercatat dan panel menampilkan galat untuk ditindaklanjuti admin.
  4. Non-approver hanya melihat status "Waiting for approval" beserta daftar approver, tanpa tombol aksi.
  5. ⚠ Keputusan bersifat final per langkah dan tidak dapat dibatalkan (tanpa undo); bila work item kembali ke status approval, langkah approval baru dimulai.
  6. ⚠ Board dan List menampilkan indikator "Awaiting approval" (ikon pada kartu/baris) untuk work item yang sedang menunggu approval.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue view di [06](06-work-item-inti.md); kartu board di [10](10-board-kolom-kartu.md); List di PLAN-016 ([18](18-timeline-rencana-tampilan-proyek.md)).

### BIZ-045 Approval notifications & email

- **Deskripsi:** Notifikasi kepada approver dan pihak terkait selama proses approval.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Approver, reporter, watcher.
- **Aturan perilaku (acceptance criteria):**
  1. Saat work item masuk status approval, setiap approver menerima email dan notifikasi in-app "Approval needed" berisi key, summary, dan pemohon.
  2. Tautan di notifikasi membuka issue view dengan panel approval tersorot.
  3. Saat keputusan final diambil, reporter dan watcher menerima notifikasi berisi hasil (Approved/Declined) dan nama pemberi keputusan.
  4. ⚠ Approve/Decline langsung dari email tidak tersedia untuk project business; email hanya berisi tautan ke issue view.
  5. Pengiriman notifikasi approval mengikuti notification scheme project (event approval needed/decided) dan preferensi notifikasi pengguna.
- **Team-managed vs company-managed:** Company-managed mengikuti notification scheme; team-managed mengikuti pengaturan notifikasi project.
- **Catatan & ketergantungan:** Mekanik notifikasi di [19](19-notifikasi-email.md); event domain approval perlu ditambahkan ke katalog event modul 28.

### BIZ-046 Approval history & tracking

- **Deskripsi:** Jejak keputusan approval pada work item beserta cara melacaknya lewat pencarian dan automation.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Anggota project, auditor, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Activity/History mencatat setiap keputusan: approver, keputusan (Approved/Declined), waktu, dan langkah/status terkait.
  2. Riwayat langkah approval sebelumnya tetap tersimpan dan ditampilkan bila work item kembali ke status approval.
  3. Entri riwayat approval tidak dapat diedit atau dihapus oleh pengguna mana pun lewat UI maupun API.
  4. Work item yang menunggu approval dapat dilacak lewat filter status di List/Search (mis. `status = "Under Review"`).
  5. ⚠ Automation dapat mengisi field approver saat create (mis. berdasarkan Category) dan menyediakan trigger khusus "Approval decided" bila keputusan produk menyetujui.
  6. Fungsi JQL `approver()` dan `pendingBy()` tidak tersedia untuk approvals project business (hanya Service Management); memakainya pada project business mengembalikan hasil kosong atau galat validasi JQL yang menjelaskan batasan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** History di WF-014 ([09](09-workflow-status-transisi.md)) dan CMT-025 ([08](08-komentar-lampiran-aktivitas.md)); automation di [20](20-otomasi.md); JQL di [12](12-jql-bahasa-kueri.md).

## Intake email & otomasi bisnis

### BIZ-047 Project email requests

- **Deskripsi:** Project business dapat memiliki alamat email; email yang masuk ke alamat itu otomatis menjadi work item.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (mengatur), pengirim email.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Pengaturan tersedia di Project settings > Email requests untuk project business company-managed; dukungan team-managed diaktifkan setelah keputusan produk.
  2. ⚠ Alamat email dibuat otomatis (mis. `{key}-{acak}@{domain-inbound-site}`), ditampilkan dengan tombol salin, dan dapat dibuat ulang (alamat lama berhenti berlaku).
  3. Admin Proyek memilih work type untuk work item hasil email; work type harus mendukung field Summary dan Description.
  4. Penerimaan email dapat di-enable/disable; email yang masuk saat nonaktif ditolak (bounce) atau diabaikan sesuai konfigurasi, tanpa membuat work item.
  5. Pemetaan: subject → Summary (dipotong 255 karakter), body → Description (HTML dikonversi ke rich text), attachment email → Attachment work item (mengikuti batas attachment site).
  6. ⚠ Pengirim yang memiliki akun dengan alamat email sama menjadi Reporter; email dari pengirim tanpa akun ditolak secara default atau, bila diizinkan, dibuat dengan Reporter akun default dan alamat pengirim dicatat di work item.
  7. Proteksi loop: email auto-reply (header `Auto-Submitted`, `X-Autoreply`, `Precedence: bulk`), email dari alamat sistem sendiri, dan lonjakan email dari pengirim yang sama dalam jangka pendek diabaikan dan dicatat di log.
  8. Custom email domain tidak didukung untuk project business (berbeda dengan Service Management); hanya domain inbound bawaan site yang dipakai.
- **Team-managed vs company-managed:** Utama untuk company-managed; team-managed belum terkonfirmasi.
- **Catatan & ketergantungan:** Mekanik email masuk/keluar di NTF-034 ([19](19-notifikasi-email.md)); channel email Service Management di modul 30. Clone membutuhkan infrastruktur inbound mail (mis. webhook dari mail gateway atau poller IMAP) yang belum ada di repo; konfigurasinya masuk `internal/core/config`.

### BIZ-048 Business automation templates

- **Deskripsi:** Template automation untuk kasus bisnis yang ditawarkan di project business.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (membuat rule project-scoped).
- **Aturan perilaku (acceptance criteria):**
  1. Galeri menyediakan template contoh: auto-assign saat work item dibuat, reminder sebelum due date, isi approver otomatis (mis. berdasarkan Category), dan buat subtask onboarding.
  2. Template diakses dari Project settings > Automation atau dari ikon automation di view project; memilih template membuka editor rule yang sudah terisi dan dapat disesuaikan sebelum diaktifkan.
  3. Batas eksekusi rule mengikuti tier plan site; saat batas tercapai, rule bisnis berhenti dijalankan dan admin menerima peringatan.
  4. Submission form memicu trigger "Work item created", sehingga template dapat dipakai untuk memproses permintaan dari form.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine dan editor automation di [20](20-otomasi.md) (AUTO-013 loop protection, AUTO-029 trigger form); tier di VISI-003 ([00](00-visi-ruang-lingkup-glosarium-baseline.md)).

## Pertanyaan terbuka

1. Apakah form kerja bisnis perlu mendukung pengiriman publik tanpa login (BIZ-036), atau intake eksternal sepenuhnya diarahkan ke portal Service Management?
2. Apakah conditional logic (BIZ-033), section, dan template form (BIZ-030) dibangun untuk form bisnis di Fase 3, atau ditunda dan dibangun bersama engine JSM forms?
3. Apakah langkah approval (BIZ-042) juga diaktifkan untuk project team-managed dan project software sejak awal, atau hanya untuk project business company-managed?
4. Aturan jumlah persetujuan mana yang menjadi default (satu approver atau semua approver), dan apakah approver boleh diganti saat approval pending (BIZ-043)?
5. Apakah UI berbahasa Indonesia memakai istilah "Project" atau "Space" (dan padanan Indonesianya) sebagai default, dan apakah rename ini dapat dimatikan per site (BIZ-001)?
6. Untuk tab Pages (BIZ-026), apakah produk ini akan membangun modul dokumen internal, mengintegrasikan wiki pihak ketiga, atau hanya menyediakan daftar tautan eksternal?
7. Bagaimana format alamat email project dan perlakuan email dari pengirim tanpa akun (BIZ-047), serta infrastruktur inbound mail apa yang dipakai?
8. Apakah Category (BIZ-021) tersedia di project team-managed dan apa batas jumlah Category, work type, field, dan form per project?
9. Template bisnis mana saja (di luar Task tracking dan Project management) yang menjadi katalog resmi rilis Fase 3?
10. Apakah anggota project (bukan hanya Admin Proyek) boleh membuat dan mengedit form (BIZ-034)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk form lanjutan (mis. pembuat form eksternal dengan logika kompleks, pembayaran, e-signature), approval multi-level kompleks, checklist, dan CRM penuh tidak dibangun sebagai fitur native; kebutuhan tersebut dipenuhi lewat API publik dan webhook (modul 23) atau ekstensi (modul 25).
- **Produk dokumentasi Atlassian (Confluence)** sebagai layanan lintas produk tidak dibangun; tab Pages (BIZ-026) hanya menyediakan titik integrasi.
- **Overviews** lintas project khas Jira Work Management lama tidak dibangun; kebutuhan lintas project diarahkan ke Plans (modul 18).
- **Approvals Service Management** (approver group, auto-decline, approve via email, JQL `approver()`/`pendingBy()`) dan **JSM forms** (ex ProForma, ekspor XLSX) dirinci di modul Service Management ([29](29-jsm-service-project-request-type-formulir.md) dan [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)), bukan di modul ini.
- **Billing, pricing, dan pengelolaan lisensi Atlassian** (BIZ-006) tidak dibangun; tier diperlakukan sebagai feature flag konfigurasi.
- **Unfurl smart link di Slack/Microsoft Teams** bergantung pada integrasi pihak ketiga di modul 25.
- **Fitur Data Center-only** terkait Jira Core server (mis. migrasi produk Server/DC) hanya dirujuk lewat impor di modul 25; tidak ada fitur bertier dc-only-or-deprecated di jatah modul ini.
- **CMDB/Assets** untuk template Asset tracking berada di modul Service Management Assets ([33](33-jsm-assets-cmdb.md)); template di modul ini hanya pelacakan sederhana berbasis work item.
