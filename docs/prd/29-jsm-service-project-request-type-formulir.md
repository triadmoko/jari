# 29 — Service Management: Service Project, Request Type, Forms & Agent

Modul ini mendefinisikan fondasi **Service Management** di produk ini: tipe project `service_desk` (service project/service space) beserta wizard pembuatan dan katalog templatenya (General, ITSM, HR, Facilities, Legal, Finance, Customer service, Marketing/Analytics/Design/Sales, Blank, demo space), navigasi dan halaman pengaturan khas service project, **Request type** sebagai lapisan customer-facing di atas work type (request form, hidden fields dengan preset, layout issue view agent, nama status customer-facing, pembatasan akses, pemetaan per channel), **portal groups**, workflow dan resolusi bawaan service project serta automation rule bawaan, **Forms** (form builder eks ProForma: validasi, conditional logic, linked fields, data connections, template, lifecycle, ekspor), model pengguna **agent / collaborator / customer** beserta role dan permission scheme khas service project, **pengalaman agent di issue view** (Reply to customer vs Add internal note, request details & participants, raise on behalf, panel customer, SLA, approvals, knowledge base), field sistem dan ekstensi JQL khusus request, halaman pengaturan kanal dan komunikasi customer di tingkat project (portal, email, widget, chat, language support, customer notifications, notifikasi internal, satisfaction), serta batas plan service project. Mesin umum (entitas project, work item, field, workflow engine, notifikasi, automation, JQL) dimiliki modul inti; modul ini mengatur **perilaku khusus service management** dan pemasangannya. Detail help center/portal, queue, SLA, knowledge base, ITSM, Assets, Operations, dan AI dirinci di modul 30-35.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users` dengan kolom `email`, `username`, `password`, `name`, `is_active`; migrasi `001_create_users_table`). Belum ada entitas project, work item, request type, maupun form, dan belum ada pembeda tipe akun (akun berlisensi vs customer portal-only). Seluruh fitur modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/serviceproject` (pembuatan project, template, Features, settings hub), `internal/features/requesttype` (request type, portal group, request form, preset, status customer-facing), `internal/features/jsmform` (form builder, submission, lifecycle, ekspor), dan `internal/features/agentview` (endpoint issue view agent, reply/internal note, participants), dengan entitas bersama di `internal/shared/domain/entity` dan UI agent di `frontend/src/routes` (SPA SvelteKit).

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah Work item/Work type/Space (VISI-002) dan tier plan sebagai feature flag (VISI-003).
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — login akun berlisensi dan akun customer portal-only.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — akun pengguna, grup, product access.
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission scheme (PERM-014, PERM-016), project role (PERM-018, PERM-019), product access (PERM-004), issue security (PERM-042 s.d. PERM-044), visibilitas komentar (PERM-045).
- [04 Proyek](04-proyek.md) — entitas project, template library (PRJ-005), wizard (PRJ-010), aturan key (PRJ-012), share settings (PRJ-014), details (PRJ-016), summary (PRJ-028), trash (PRJ-033), settings hub (PRJ-034), People (PRJ-037).
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — work type, custom field (CFG-025 s.d. CFG-037), field configuration (CFG-043), screen (CFG-046 s.d. CFG-048), resolution (CFG-055).
- [06 Work Item Inti](06-work-item-inti.md) — create dialog (ISS-002), issue view (ISS-007), reporter (ISS-017), resolution (ISS-019), link (ISS-029), move/clone (ISS-031, ISS-032), bulk change (ISS-035), restrict (ISS-042).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar (CMT-001 s.d. CMT-010), lampiran (CMT-011 s.d. CMT-019), history (CMT-022).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category (WF-002), resolution semantics (WF-009), workflow editor (WF-021), validator/post function (WF-028, WF-029), terjemahan status (WF-040).
- [12 JQL](12-jql-bahasa-kueri.md) — referensi custom field (JQL-033) dan fungsi (JQL-037).
- [19 Notifikasi & Email](19-notifikasi-email.md) — notification scheme (NTF-004), infrastruktur email masuk/keluar (NTF-023, NTF-033), template email (NTF-042).
- [20 Otomasi](20-otomasi.md) — engine rule (AUTO-003, AUTO-012), trigger form submitted (AUTO-029), smart values (AUTO-063).
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — sidebar (NAV-002, NAV-003), create dialog global (NAV-013), batas plan (NAV-024).
- [26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) — Forms kerja bisnis (BIZ-029 s.d. BIZ-041) yang berbagi engine form dengan modul ini.
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) dan [28 Model Data Inti](28-model-data-inti.md) — entitas dan event domain.
- [30 Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — help center, portal customer, customers & organizations, email channel, widget, chat.
- [31 Queues, SLA, Knowledge Base, CSAT & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) — queue, SLA, KB, CSAT, canned responses.
- [32 ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — incident, problem, change, services, approvals.
- [33 Assets](33-jsm-assets-cmdb.md), [34 Operations](34-jsm-operations-alerting-on-call.md), [35 Virtual Service Agent, AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md).

## Persona terkait

- **Admin Situs** — memberi product access Service Management (seat agent), mengatur konfigurasi global Service Management, menyetujui permintaan pembuatan service space, mengelola permission scheme dan notification scheme bersama.
- **Admin Proyek/Project Lead** — membuat service project dari template, mengelola request type, portal group, request form, Forms, Features, role, customer permissions, kanal, terjemahan, dan notifikasi customer.
- **Agent Layanan** — bekerja di queue, membalas customer, menulis internal note, raise request on behalf, mengelola participants, memakai panel SLA/approval/KB, mengisi dan membuka ulang form.
- **Manajer Layanan/ITSM** — memilih template ITSM, mengaktifkan work categories, merancang workflow, SLA, dan automation rule bawaan, serta memantau ringkasan project.
- **Customer/Pelapor** — mengajukan request lewat portal/email/widget/chat, mengisi form, melihat status customer-facing, menambah participant, memberi rating kepuasan.
- **Developer/Anggota Tim** — berperan sebagai collaborator (berlisensi produk inti tanpa seat agent) yang membantu agent lewat internal note dan linked work item.
- **Stakeholder/Viewer** — menerima update incident dan melihat ringkasan tanpa seat agent.
- **Integrator/Developer API** — membuat request lewat API dengan request type, membaca field request, menjalankan JQL khusus request.

## User story modul

1. Sebagai Admin Proyek/Project Lead, saya ingin membuat service project dari template (General, ITSM, HR, dll.) agar portal, request type, workflow, queue, dan SLA bawaan langsung siap dipakai (JSM-PRJ-001, JSM-PRJ-002, JSM-PRJ-003, JSM-PRJ-005).
2. Sebagai Admin Proyek/Project Lead, saya ingin mendefinisikan request type dengan nama, ikon, dan deskripsi yang ramah customer di atas work type internal agar customer menemukan jenis permintaan yang tepat tanpa memahami struktur internal (JSM-PRJ-020, JSM-PRJ-021, JSM-PRJ-022).
3. Sebagai Admin Proyek/Project Lead, saya ingin mengatur field yang terlihat customer beserta label dan help text khusus, serta field tersembunyi dengan nilai preset, agar data request lengkap dan otomatis ter-routing (JSM-PRJ-025, JSM-PRJ-026).
4. Sebagai Customer/Pelapor, saya ingin melihat nama status yang mudah dipahami seperti "Being reviewed" alih-alih status internal agar saya tahu posisi permintaan saya (JSM-PRJ-028, JSM-PRJ-031).
5. Sebagai Admin Proyek/Project Lead, saya ingin menyusun form dengan section bersyarat, validasi, dan pertanyaan yang tertaut ke field agar permintaan kompleks seperti onboarding karyawan terkumpul terstruktur (JSM-PRJ-034, JSM-PRJ-035, JSM-PRJ-036, JSM-PRJ-037, JSM-PRJ-038).
6. Sebagai Agent Layanan, saya ingin membedakan dengan jelas antara balasan ke customer dan catatan internal agar informasi internal tidak pernah bocor ke customer (JSM-PRJ-053, JSM-PRJ-068).
7. Sebagai Agent Layanan, saya ingin satu issue view yang menampilkan request type, channel asal, SLA, approval, konteks customer, dan artikel KB agar saya dapat menyelesaikan request tanpa berpindah halaman (JSM-PRJ-052, JSM-PRJ-056, JSM-PRJ-057, JSM-PRJ-058, JSM-PRJ-059).
8. Sebagai Agent Layanan, saya ingin membuat request atas nama customer yang menelepon agar request tetap tercatat dengan reporter dan notifikasi yang benar (JSM-PRJ-055).
9. Sebagai Admin Situs, saya ingin hanya pengguna dengan seat Service Management yang dapat menjadi agent dan customer tidak memakan lisensi agar biaya terkendali (JSM-PRJ-045, JSM-PRJ-046, JSM-PRJ-048, JSM-PRJ-070).
10. Sebagai Admin Proyek/Project Lead, saya ingin menentukan siapa yang boleh mengajukan request dan dengan siapa customer boleh berbagi request agar portal internal maupun eksternal aman (JSM-PRJ-051, JSM-PRJ-029).
11. Sebagai Customer/Pelapor, saya ingin portal, form, dan notifikasi tampil dalam bahasa saya agar permintaan mudah diajukan (JSM-PRJ-066, JSM-PRJ-067).
12. Sebagai Manajer Layanan/ITSM, saya ingin request type dikelompokkan ke work categories ITSM (Service requests, Incidents, Problems, Changes) agar tiap praktik punya queue dan tampilan sendiri (JSM-PRJ-003, JSM-PRJ-017).
13. Sebagai Admin Proyek/Project Lead, saya ingin menemukan work item yang tidak punya request type dan menetapkannya secara massal agar semua request terlihat di portal customer (JSM-PRJ-054, JSM-PRJ-061).
14. Sebagai Developer/Anggota Tim, saya ingin membantu agent lewat internal note tanpa membutuhkan seat agent agar eskalasi teknis cepat dan hemat lisensi (JSM-PRJ-047).
15. Sebagai Integrator/Developer API, saya ingin request yang dibuat lewat API otomatis mendapat request type default agar tidak ada request yang "hilang" dari portal (JSM-PRJ-020, JSM-PRJ-030).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 27 |
| P1 | 29 |
| P2 | 12 |
| P3 | 2 |
| **Total** | **70** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1-4 | 0 |
| Fase 5 | 68 |
| Fase 6 | 0 |
| Fase 7 | 1 |
| Fase — | 1 |

Strategi urutan implementasi: seluruh modul ini adalah pintu masuk Fase 5 sehingga dikerjakan dalam tiga gelombang di dalam fase tersebut. Gelombang 5a membangun kerangka: tipe project `service_desk`, wizard dengan template General (JSM-PRJ-001, JSM-PRJ-002), model agent/customer dan role/permission bawaan (JSM-PRJ-045 s.d. JSM-PRJ-051), request type, portal group, request form, status customer-facing (JSM-PRJ-020 s.d. JSM-PRJ-028), workflow dan resolusi bawaan (JSM-PRJ-031, JSM-PRJ-032), serta issue view agent dengan Reply to customer/Add internal note (JSM-PRJ-052 s.d. JSM-PRJ-055) — ini cukup untuk alur portal dan email minimal. Gelombang 5b menambah konfigurabilitas: hidden fields, issue view per request type, access restriction, request type per channel, Features, settings email/portal/widget, customer notifications, notifikasi internal, field sistem dan JQL khusus request, panel SLA/approval/KB, satisfaction, serta template ITSM (bagian service request) dan HR/Customer service. Gelombang 5c menambahkan Forms lengkap (JSM-PRJ-034 s.d. JSM-PRJ-044), language support, template P2, ringkasan project, demo space, dan pembedaan plan; engine Forms dibangun sekali di modul 26 (BIZ-030 Form builder di Fase 3, BIZ-033 Form conditional logic) dan gelombang ini hanya menambah binding ke request type, portal, versi form, dan instance form pada request, sehingga engine tidak dibangun dua kali. Pemilik fitur administrasi yang tumpang tindih dengan modul 35 ditetapkan per pasangan di Catatan masing-masing fitur (JSM-PRJ-019, JSM-PRJ-046, JSM-PRJ-047, JSM-PRJ-050, JSM-PRJ-061, JSM-PRJ-070). Pengaturan Chat (JSM-PRJ-065) mengikuti kanal chat di Fase 7, sedangkan bagian incident/problem/change dari template ITSM dan work categories baru aktif penuh bersamaan dengan modul 32 di Fase 7. Inventaris memuat banyak pasangan fitur tumpang tindih dari lampiran JSM dan domain JSM rinci; 15 pasangan digabung menjadi satu heading dengan baris Alias inventaris.

## Fitur

## Pembuatan service project & template

### JSM-PRJ-001 Service project creation wizard

- **Alias inventaris:** Service project creation & templates
- **Deskripsi:** Alur pembuatan service project baru mulai dari memilih template, mengisi detail, sampai project siap dipakai lengkap dengan request type, workflow, queue, SLA, portal, dan alamat email bawaan. Memakai wizard project inti (PRJ-010) dengan kategori template "Service management".
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (global permission Administer/Create project); pengguna yang diberi izin Create project oleh admin; pengguna internal non-admin hanya dapat mengajukan permintaan pembuatan service space bila setting diizinkan.
- **Aturan perilaku (acceptance criteria):**
  1. Wizard dapat dibuka dari Projects > Create project lalu kategori "Service management" di template library, dan dari tombol Create di sidebar/top nav; kategori ini hanya tampil bila site memiliki product access Service Management.
  2. Setiap kartu template menampilkan preview: deskripsi, work type yang dibuat, request type bawaan, workflow, dan queue yang akan dibuat; preview tidak membuat data apa pun.
  3. Setelah "Use template", langkah detail meminta Name (wajib), Key (dibentuk otomatis dari nama, dapat diedit, wajib unik, hanya huruf kapital/angka, diawali huruf), dan Access/project lead; tombol Create nonaktif selama ada field tidak valid.
  4. Opsi "Share settings with an existing project" memakai scheme yang sama (work type, workflow, field configuration, screen, permission, notification) dari project service lain alih-alih membuat scheme baru; UI menampilkan peringatan bahwa perubahan scheme berdampak ke kedua project.
  5. ⚠ Wizard menanyakan tipe customer (internal atau eksternal) atau help center tujuan; jawaban menentukan default Customer permissions (JSM-PRJ-051) dan help center tempat portal digabungkan.
  6. Pembuatan berjalan atomik dan otomatis membuat: issue type scheme, request types, portal groups, queues, SLA bawaan, workflow scheme, permission scheme Service Management, notification scheme, dan default automation rules; bila satu langkah gagal seluruh pembuatan di-rollback dan pesan galat ditampilkan.
  7. Satu portal dibuat per project dan langsung tergabung ke help center site (modul 30).
  8. Hanya pengguna dengan global permission Administer (atau izin Create project yang diberikan admin) yang dapat membuat project; pengguna lain mendapat galat 403 di API dan tombol Create tidak tampil di UI.
  9. ⚠ Bila setting "Allow internal accounts to request a service space" aktif (JSM-PRJ-019), pengguna internal non-admin dapat mengirim permintaan pembuatan service space berisi nama dan alasan; admin menerima notifikasi dan dapat menyetujui (membuka wizard terisi) atau menolak dengan alasan.
  10. Validasi key: key bentrok dengan project aktif, key yang pernah dipakai project lain (alias key lama) atau project di trash ditolak dengan pesan spesifik; panjang nama dan key mengikuti batas PRJ-011/PRJ-012.
  11. Setelah berhasil, pengguna diarahkan ke Queues atau halaman getting started project.
  12. ⚠ Checklist getting started menampilkan langkah: atur portal, undang tim (agent), hubungkan email, sesuaikan request type; status tiap langkah tercentang otomatis saat langkah selesai dan checklist dapat ditutup.
  13. ⚠ Tidak ada aksi "change template" otomatis setelah project dibuat; perubahan template hanya dapat dilakukan manual dengan mengubah request type, workflow, dan Features.
  14. Template ITSM membuat work category Service requests, Incidents, Problems, Changes, dan Post-incident reviews dengan workflow dan SLA ITSM bawaan (rinci di JSM-PRJ-003).
  15. Template General service management membuat service desk umum tanpa kategori ITIL (JSM-PRJ-002).
  16. Template Customer service management ditujukan untuk support pelanggan eksternal (JSM-PRJ-006).
  17. ⚠ Katalog template bisnis memuat HR, Facilities, Legal, Finance, Marketing, Analytics, Design, Sales service management, dan Blank (JSM-PRJ-004 s.d. JSM-PRJ-010); daftar template dikelola sebagai data seed agar dapat ditambah tanpa rilis kode.
  18. Pengguna memilih company-managed atau team-managed; company-managed memakai issue type scheme, workflow scheme, dan field configuration bersama, sedangkan team-managed membuat work type sendiri per request type (lihat JSM-PRJ-012 untuk status dukungan team-managed).
  19. Opsi share settings hanya tersedia untuk company-managed dan hanya menawarkan project service yang dapat dikelola pengguna.
  20. ⚠ Checkbox opsional "Include sample data" membuat beberapa request contoh beserta customer contoh yang dapat dihapus massal.
  21. Detail project yang ditetapkan saat pembuatan: project key, nama, avatar, project lead, dan default assignee (Project lead atau Unassigned).
  22. Pembuatan project otomatis menyiapkan portal, queue default, SLA default, alamat email cloud bawaan project (JSM-PRJ-063), serta role Service Desk Team dan Service Desk Customers.
  23. ⚠ Checklist onboarding untuk admin project baru tampil pertama kali admin membuka project dan dapat dibuka kembali dari menu bantuan project.
  24. Batas 3 agent pada plan Free dihitung per site (seat agent unik lintas semua service project), bukan per project; wizard tidak menolak pembuatan project baru karena batas agent.
- **Team-managed vs company-managed:** Company-managed adalah target utama clone (scheme bersama, share settings). Team-managed membuat work type per request type dan tidak mendukung share settings; status dukungannya diatur di JSM-PRJ-012.
- **Catatan & ketergantungan:** Wizard inti PRJ-010, aturan nama/key PRJ-011/PRJ-012, share settings PRJ-014, onboarding PRJ-015 ([04](04-proyek.md)). Sejak 2025 label UI "project" dapat ditampilkan sebagai "space" (VISI-002). Repo saat ini belum memiliki tabel `projects`; modul ini menambah kolom `project_type_key='service_desk'` dan tabel `service_project_settings`. Permintaan service space butuh entitas `service_space_requests` (requester, name, reason, status).

### JSM-PRJ-002 General service management template

- **Deskripsi:** Template serbaguna untuk tim layanan internal/bisnis dengan request type umum dan workflow sederhana; menjadi template default pertama yang dikirim di Fase 5.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead dan Admin Situs saat memakai wizard.
- **Aturan perilaku (acceptance criteria):**
  1. Template membuat work type: [System] Service request, [System] Service request with approvals, dan Task; work type berlabel [System] tidak dapat dihapus selama dipakai request type.
  2. ⚠ Template membuat contoh request type: Get help/Ask a question, Report a problem, Request access, Suggest an improvement, masing-masing di portal group default.
  3. Workflow sederhana memuat status Waiting for support > In progress > Waiting for customer > Resolved/Closed, dengan resolusi di-set saat transisi ke Resolved.
  4. SLA bawaan Time to first response dan Time to resolution dibuat dengan kalender default dan goal bawaan (modul 31).
  5. Queue bawaan dibuat: All open, Assigned to me, Unassigned, Recently resolved.
  6. Kapabilitas lain (approvals, KB, CSAT, dll.) dapat diaktifkan kemudian lewat halaman Features (JSM-PRJ-016) tanpa membuat ulang project.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Daftar request type contoh disimpan sebagai data seed template. SLA dan queue di modul 31.

### JSM-PRJ-003 IT service management template (ITSM)

- **Deskripsi:** Template service project untuk tim IT yang menyiapkan praktik ITIL: service request, incident, problem, change, dan post-incident review. Di Fase 5 bagian service request berfungsi penuh; bagian incident/problem/change memakai workflow dan field bawaan dan diperkaya modul 32-34 di Fase 7.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs, Admin Proyek/Project Lead, Manajer Layanan/ITSM; agent memakai hasilnya.
- **Aturan perilaku (acceptance criteria):**
  1. Template membuat work type: [System] Service request, [System] Service request with approvals, [System] Incident, [System] Problem, [System] Change, [System] Post-incident review, Task, dan Sub-task.
  2. Template mengaktifkan work categories Service requests, Incidents, Problems, Changes, Post-incident reviews, masing-masing dengan halaman/queue sendiri di sidebar (JSM-PRJ-017).
  3. ⚠ Template membuat contoh request type: Get IT help, Fix an account problem, Get a guest wifi account, Request a new account, Request admin access, Report a system problem, Request new hardware, Request new software, Request a change, Onboard new employees.
  4. ⚠ Template membuat portal group bawaan: Common requests, Logins and accounts, Computers, Applications, Servers and infrastructure.
  5. Workflow bawaan per work type: Service request fulfilment, Service request fulfilment with approvals, Incident management, Problem management, Change management.
  6. SLA bawaan dibuat: Time to first response, Time to resolution, Time to close after resolution, Time to approve normal change.
  7. Queue bawaan per kategori dibuat: All open, Assigned to me, Unassigned, Incidents, Service requests, Changes, Problems.
  8. Fitur ITSM Services, Assets, change calendar, deployment gating, Alerts/On-call (Operations), dan major incident tercatat aktif di Features; selama modul terkait belum tersedia, item menu disembunyikan dan toggle tampil sebagai "Segera hadir".
  9. Field ITSM bawaan ditambahkan ke screen terkait: Urgency, Impact, Priority, Affected services, Change type, Change risk, Change reason, Pending reason, Root cause, Workaround.
  10. Template ITSM adalah satu-satunya template yang mengaktifkan work categories ITSM lengkap; template lain hanya mendapat kategori Service requests atau tidak ada kategori.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed tidak menawarkan template ITSM lengkap (JSM-PRJ-012).
- **Catatan & ketergantungan:** Incident/problem/change/PIR di modul 32, Assets modul 33, Operations modul 34. Change risk assessment lanjutan dan deployment gating dapat bergantung plan (JSM-PRJ-070).

### JSM-PRJ-004 Blank service project template

- **Deskripsi:** Template service project kosong dengan konfigurasi minimal untuk dibangun dari nol.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Template hanya membuat satu work type dasar ([System] Service request), satu request type dasar, dan workflow sederhana (Waiting for support, In progress, Resolved).
  2. Portal tetap dibuat dan tergabung ke help center seperti template lain.
  3. Tidak ada portal group selain satu grup default, tidak ada queue selain All open, dan tidak ada SLA; admin menambahkan request type, portal group, queue, dan SLA sendiri.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Padanan Blank project inti (PRJ-008).

### JSM-PRJ-005 HR service management template

- **Deskripsi:** Template untuk tim HR menangani onboarding, offboarding, pertanyaan benefit, dan perubahan data karyawan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead; agent HR.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat contoh request type: Onboard new employee, Offboard employee, Ask a question about benefits/payroll, Update personal details, Request time off, Request employment verification.
  2. Template membuat workflow onboarding/offboarding dan workflow general HR request.
  3. ⚠ Template membuat portal group HR: Employee onboarding, Payroll and benefits, Personal details.
  4. Halaman getting started menampilkan rekomendasi untuk membatasi request type sensitif (JSM-PRJ-029) dan memakai issue security (JSM-PRJ-050) untuk data karyawan.
  5. Integrasi onboarding lintas tim (membuat sub-task/work item ke project IT/Facilities lewat automation) disediakan sebagai template rule yang nonaktif secara default.
  6. Template menyarankan menautkan knowledge base kebijakan HR untuk self-service (modul 31); bila KB belum tersambung, saran tampil di checklist.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Automation di [20](20-otomasi.md); KB di modul 31.

### JSM-PRJ-006 Customer service management template

- **Deskripsi:** Template untuk tim layanan pelanggan eksternal (external support) dengan help center yang dapat diakses customer di luar organisasi.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead; agent support.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat contoh request type: Get help with a product, Report a bug, Request a feature, Billing question, Account issue.
  2. Portal ditujukan untuk customer eksternal; public signup atau request anonim hanya aktif bila diatur di Customer permissions (JSM-PRJ-051) dan diizinkan global.
  3. Customer & organization details (custom attributes) tampil di issue view agent melalui Customer context panel (JSM-PRJ-056).
  4. Kanal email, portal, widget, dan chat disiapkan sebagai kanal utama dan ditandai di checklist getting started.
  5. SLA bawaan respons pertama dan resolusi dibuat.
  6. ⚠ Customer Service Management (CSM) sebagai aplikasi terpisah di atas platform Service Management tidak dibangun di modul ini; template ini tetap ditawarkan sebagai template service project dan CSM dirinci di modul 35.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Customers & organizations di modul 30; CSM di modul 35 (Fase 7).

### JSM-PRJ-007 Facilities service management template

- **Deskripsi:** Template untuk tim fasilitas menangani permintaan ruang, perbaikan, pemeliharaan, dan keamanan gedung.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead; agent fasilitas.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat contoh request type: Request maintenance/repairs, Report a facilities problem, Request a desk/office move, Request access card, Report a health and safety issue.
  2. Template membuat workflow fulfilment dan workflow maintenance.
  3. Template membuat portal group berdasarkan kategori fasilitas.
  4. Template membuat custom field lokasi/gedung (select list); bila Assets aktif, field dapat diganti Assets object field (modul 33).
  5. SLA bawaan respons dan resolusi dibuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Custom field di [05](05-konfigurasi-work-item-skema.md).

### JSM-PRJ-008 Legal service management template

- **Deskripsi:** Template untuk tim legal menangani review kontrak, NDA, dan konsultasi hukum.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead; agent legal.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat contoh request type: Request contract review, Request an NDA, Ask a legal question, Report a compliance issue.
  2. Template membuat workflow legal request dengan langkah review/approval.
  3. Checklist getting started menyarankan pembatasan request type (JSM-PRJ-029) dan pemakaian internal note untuk kerahasiaan.
  4. Template menyertakan form legal (JSM-PRJ-034) untuk detail kontrak: pihak, nilai, dan tanggal.
  5. SLA bawaan dibuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Approval di modul 32; Forms di JSM-PRJ-034.

### JSM-PRJ-009 Finance service management template

- **Deskripsi:** Template untuk tim keuangan menangani pengeluaran, pembelian, invoice, dan pertanyaan keuangan.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead; agent finance; approver.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Template membuat contoh request type: Submit an expense/reimbursement, Request a purchase order, Invoice query, Ask a finance question.
  2. Request type pengeluaran/pembelian memakai workflow Service request with approvals untuk persetujuan anggaran.
  3. Template membuat custom field nominal (number), mata uang (select), dan cost center (select).
  4. Approver diambil dari field user/group yang dikonfigurasi di langkah approval.
  5. SLA bawaan dibuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Approvals di modul 32 dan JSM-PRJ-058.

### JSM-PRJ-010 Marketing, Analytics, Design and Sales service management templates

- **Deskripsi:** Template tambahan untuk tim bisnis lain yang membawa request type, portal group, workflow, dan SLA khusus use case tersebut.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Template Marketing service management membuat request type untuk campaign, konten, dan dukungan event.
  2. Template Analytics service management membuat request type untuk laporan, akses data, dan dashboard.
  3. ⚠ Template Design service management membuat request type untuk aset desain dan review desain.
  4. ⚠ Template Sales service management membuat request type untuk quote dan dukungan deal.
  5. Setiap template membawa request types, portal groups, workflow, dan SLA spesifik use case-nya.
  6. Daftar template disimpan sebagai data seed berversi sehingga dapat berubah antar rilis tanpa memengaruhi project yang sudah dibuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme seed sama dengan JSM-PRJ-001.

### JSM-PRJ-011 Demo / sample service space

- **Deskripsi:** Service space contoh berisi data dummy untuk mempelajari Service Management sebagai agent dan admin, termasuk tur tugas dasar.
- **Prioritas / Fase / Tier:** P3 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs atau pengguna dengan izin Create project; agent memakai tur.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Demo space dapat dibuat dari checklist onboarding atau dari template library ("Try a demo space").
  2. Demo space berisi data contoh: request types, work items, customers contoh (akun dummy tanpa email terkirim), queues, dan SLA.
  3. ⚠ Tur terpandu memandu langkah: assign diri ke work item, membalas customer, dan resolve request; progres tur tersimpan per pengguna.
  4. Demo space dapat dihapus atau dipindahkan ke trash memakai fitur delete project inti (PRJ-033).
  5. ⚠ Demo space dihitung sebagai project biasa untuk batas jumlah project, tetapi customer contohnya tidak dihitung dan tidak menerima email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditandai P3; boleh ditunda tanpa memblokir fitur lain.

### JSM-PRJ-012 Team-managed service project (legacy)

- **Deskripsi:** Varian service project yang dikelola tim dengan konfigurasi mandiri tanpa scheme bersama. Dicatat sebagai rujukan; clone tidak menargetkan varian ini.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Proyek (pada produk referensi); di clone tidak ada aktor karena varian ini tidak dibangun.
- **Aturan perilaku (acceptance criteria):**
  1. Di produk referensi, request type, field, dan workflow dikonfigurasi langsung di project tanpa scheme global; clone tidak membangun mode ini untuk service project.
  2. ⚠ Notifikasi internal team-managed diatur lewat toggle per event tanpa notification scheme; clone memakai notification scheme (JSM-PRJ-068).
  3. Varian ini tidak mendukung sebagian fitur company-managed (fitur ITSM, scheme bersama); karena itu clone hanya menyediakan service project company-managed.
  4. ⚠ Pembuatan service project team-managed baru dibatasi/dihentikan di produk referensi; importer (modul 25) memetakan project team-managed lama menjadi company-managed dengan scheme yang dibentuk dari konfigurasinya.
- **Team-managed vs company-managed:** Seluruh fitur ini adalah varian team-managed; clone hanya membangun company-managed.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan. Importer di [25](25-integrasi-impor-ekspor-ekstensi.md).

## Navigasi, pengaturan project & konfigurasi global

### JSM-PRJ-013 Service project navigation sidebar

- **Deskripsi:** Navigasi khas service project untuk agent yang mengelompokkan queues, work categories, customer, kanal, dan pengaturan di sidebar project.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Browse project + role Service Desk Team); Admin Proyek; collaborator melihat sidebar terbatas tanpa Queues.
- **Aturan perilaku (acceptance criteria):**
  1. Sidebar menampilkan menu Queues (bagian starred dan semua queue, masing-masing dengan jumlah work item yang diperbarui berkala) serta work categories (Service requests, Incidents, Problems, Changes, Post-incident reviews) bila aktif.
  2. ⚠ Menu Summary (JSM-PRJ-014) tampil di atas Queues bila fitur Summary aktif.
  3. Menu Alerts/On-call/Operations hanya tampil bila fitur terkait aktif di Features dan modul 34 tersedia.
  4. Menu Knowledge base, Reports, Customers, Services, dan Assets tampil sesuai fitur aktif dan izin pengguna.
  5. Bagian Channels & self service menautkan pengaturan portal, email, chat, widget, dan virtual service agent.
  6. Aksi "Raise a request" / "View portal" membuka portal project sebagai agent di tab baru.
  7. Menu Project settings hanya tampil bagi pengguna dengan izin Administer project.
  8. Item menu tampil/hilang mengikuti toggle Features (JSM-PRJ-016) dan plan site (JSM-PRJ-070) tanpa reload halaman setelah toggle diubah.
  9. Sidebar dapat di-collapse/expand dengan state tersimpan per pengguna, dan pengguna dapat memberi bintang pada queue sehingga muncul di bagian starred.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sidebar umum NAV-002/NAV-003 ([21](21-navigasi-ux-preferensi.md)); queues di modul 31. Label menu mengikuti navigasi 2025 dan disimpan sebagai kunci i18n.

### JSM-PRJ-014 Project summary page (service space)

- **Deskripsi:** Halaman ringkasan service space yang menampilkan gambaran status, prioritas, work type, dan aktivitas work item terbaru bagi agent dan admin, memakai komponen Summary inti (PRJ-028).
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek, Manajer Layanan/ITSM (Browse project).
- **Aturan perilaku (acceptance criteria):**
  1. Kartu ringkasan menampilkan jumlah work item dibuat, diperbarui, selesai, dan jatuh tempo dalam 7 hari terakhir.
  2. Chart breakdown status, priority, dan work type ditampilkan untuk work item project.
  3. Recent activity feed menampilkan perubahan terbaru dengan aktor dan waktu.
  4. Ringkasan workload per assignee menampilkan jumlah work item terbuka per agent, termasuk Unassigned.
  5. Klik kartu atau segmen chart membuka daftar work item terkait dengan filter JQL yang sesuai.
  6. Halaman dapat diakses dari sidebar navigasi space (JSM-PRJ-013).
  7. ⚠ Kartu khusus Service Management menampilkan breakdown per request type dan jumlah SLA breached/near breach.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen PRJ-028 ([04](04-proyek.md)); data SLA dari modul 31.

### JSM-PRJ-015 Project settings - Details

- **Deskripsi:** Pengaturan identitas dasar service project.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project); Admin Situs untuk ubah key dan category.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat mengubah Name, Key, Category, dan Avatar/icon; mengubah key menampilkan peringatan dampak (URL, integrasi) dan key lama tetap me-redirect ke work item yang sama.
  2. Admin dapat mengatur Project lead dan Default assignee (Project lead atau Unassigned).
  3. Admin dapat mengisi URL dan description project.
  4. Nama dan deskripsi portal tidak diatur di sini melainkan di Portal settings (JSM-PRJ-062); halaman menampilkan tautan ke sana.
  5. Validasi: key unik dan sesuai format, nama tidak kosong; galat ditampilkan inline dan perubahan tidak disimpan.
  6. Pindah ke trash, arsip, dan hapus project memakai fitur inti (PRJ-032, PRJ-033).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** PRJ-016, PRJ-019 s.d. PRJ-022 ([04](04-proyek.md)).

### JSM-PRJ-016 Project settings - Features

- **Alias inventaris:** Project features toggles (service project)
- **Deskripsi:** Halaman toggle untuk mengaktifkan atau menonaktifkan kapabilitas service project per project tanpa menghapus data.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle Knowledge base tersedia untuk semua template.
  2. Toggle Incidents, Problems, Changes, dan Post-incident reviews (work categories ITSM) tersedia untuk project berbasis template ITSM.
  3. ⚠ Toggle Services dan Assets mengaitkan service registry dan object schema ke project.
  4. Toggle Alerts dan On-call (Operations) tersedia bila modul 34 aktif.
  5. ⚠ Toggle Development/Deployments mengaktifkan panel development dan deployment gating (modul 24 dan 32).
  6. ⚠ Toggle Chat dan virtual service agent hanya tampil bila kanal tersebut didukung plan.
  7. Mematikan fitur menyembunyikan menu sidebar dan halaman settings terkait tanpa menghapus konfigurasi atau data; menyalakan kembali memulihkan konfigurasi sebelumnya.
  8. Toggle yang hanya tersedia pada template ITSM atau plan tertentu ditampilkan nonaktif dengan keterangan alasan (template atau plan yang dibutuhkan).
  9. ⚠ Daftar toggle mencakup Knowledge base, Approvals, SLAs, Customer satisfaction, Change management, Incidents, Alerts/On-call, Deployments, Assets, Virtual service agent, dan Chat; daftar yang ditampilkan disaring per template.
  10. ⚠ Status nonaktif menyimpan data (mis. SLA tetap tersimpan) dan hanya menghentikan tampilan/eksekusi; perilaku per fitur didokumentasikan di halaman (mis. SLA berhenti dihitung saat dinonaktifkan).
  11. Pada company-managed, sebagian fitur (SLA, KB, automation) juga memiliki halaman project settings terpisah yang tetap dapat diakses dari halaman Features.
  12. Toggle Work categories mengaktifkan atau menonaktifkan pengelompokan per kategori (JSM-PRJ-017).
  13. Perubahan toggle langsung tercermin di sidebar project (Queues, Alerts, On-call, Change calendar, Reports) sesuai JSM-PRJ-013.
  14. Bagian Features lab menampilkan fitur eksperimental/beta (mis. Journeys) dengan label "Beta"; isi Features lab dikendalikan feature flag server dan dapat berubah antar rilis.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk clone (hanya company-managed).
- **Catatan & ketergantungan:** Mekanisme Features project inti di [04](04-proyek.md). Simpan sebagai tabel `project_features` (project_id, feature_key, enabled, updated_by).

### JSM-PRJ-017 Work categories

- **Alias inventaris:** Work categories (ITSM practices)
- **Deskripsi:** Kategori kerja ITSM yang mengelompokkan request type, work type, queue, dan tampilan ke halaman terpisah: Service requests, Incidents, Problems, Changes, Post-incident reviews.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (melihat), Admin Proyek (mengatur), Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Kategori yang tersedia tetap: Service requests, Incidents, Problems, Changes, Post-incident reviews; kategori tidak dapat ditambah atau diganti nama oleh admin.
  2. Request type dikaitkan ke kategori melalui work type-nya ([System] Incident → Incidents, dst.).
  3. Halaman kategori menampilkan queue dan view khusus kategori tersebut.
  4. Kategori diaktifkan/nonaktifkan lewat Features (JSM-PRJ-016).
  5. Kategori memengaruhi laporan bawaan dan SLA bawaan (mis. SLA Time to approve normal change hanya untuk Changes).
  6. Setiap request type dipetakan ke tepat satu work category.
  7. Sidebar menampilkan queue per kategori, dan kategori yang tidak dipakai dapat disembunyikan.
  8. Tampilan list Incidents menampilkan kolom priority, affected services, dan responders (diisi penuh saat modul 32/34 aktif).
  9. Tampilan Changes terhubung ke change calendar (modul 32).
  10. Tampilan Problems menampilkan incident tertaut untuk setiap problem.
  11. Kategori dapat diaktifkan/nonaktifkan dari project settings; menonaktifkan kategori tidak mengubah work type maupun request type yang ada.
  12. ⚠ Request type yang work type-nya tidak terpetakan ke kategori masuk grup "Other/Unassigned" di daftar request type dan queue.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk clone.
- **Catatan & ketergantungan:** Berbeda dari field Category kerja bisnis (BIZ-021). View Incidents/Changes/Problems rinci di modul 32.

### JSM-PRJ-018 Project settings - JSM module pages

- **Deskripsi:** Entry point pengaturan modul Service Management lain di project settings yang detailnya dikelola modul terpisah.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project).
- **Aturan perilaku (acceptance criteria):**
  1. Item SLAs membuka pengaturan goal, kondisi, dan kalender SLA (modul 31).
  2. Item Knowledge base membuka pengaturan sumber artikel KB dan akses customer ke artikel (modul 31).
  3. Item Automation membuka daftar rule project termasuk default rules (JSM-PRJ-033, AUTO-001).
  4. Item Services membuka pengaitan service ke project (modul 32).
  5. Item Change management membuka koneksi CI/CD, deployment gating, dan change calendar (modul 32).
  6. ⚠ Item Incident management membuka pengaturan incident, stakeholder, dan major incident (modul 32).
  7. Item Assets membuka object schema dan field Assets terkait project (modul 33).
  8. ⚠ Item Approvals membuka ringkasan langkah approval yang dikonfigurasi di status workflow (modul 32).
  9. Item Queues mengarahkan ke "Manage queues" di sidebar (modul 31).
  10. Item Operations/Alerts/On-call membuka pengaturan Operations (modul 34).
  11. Item Apps menampilkan aplikasi/ekstensi yang terhubung ke project (modul 25).
  12. Item menu hanya tampil bila Feature terkait aktif (JSM-PRJ-016) dan plan mendukung (JSM-PRJ-070).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Settings hub inti PRJ-034 ([04](04-proyek.md)).

### JSM-PRJ-019 JSM global configuration (product settings)

- **Deskripsi:** Pengaturan tingkat site untuk Service Management yang memengaruhi semua service project.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (Administer).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin mengatur public signup dan kebijakan akses help center di tingkat site; project tidak dapat mengizinkan public signup bila setting global mati (perilaku dimiliki JSM-PRT-015 Customer access & signup settings; halaman ini hanya menautkannya).
  2. Admin mengatur look and feel email notifikasi customer (template global: logo, warna, footer) yang dipakai semua project (perilaku dimiliki JSM-PRT-038).
  3. ⚠ Admin mengatur pemrosesan email requests global (mis. aktif/nonaktif seluruh email channel, aturan abaikan auto-reply).
  4. ⚠ Admin mengatur setting approvals/collaborator global (mis. apakah collaborator boleh menjadi approver).
  5. ⚠ Admin mengatur toggle "Allow internal accounts to request a service space" yang mengaktifkan alur permintaan di JSM-PRJ-001.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Administrasi sistem di [22](22-administrasi-sistem.md); help center di modul 30. Pemilik: halaman pengaturan produk Service Management tingkat situs dimiliki JSM-AI-006 JSM Global Settings ([35](35-jsm-ai-otomasi-api-admin-csm.md)), yang menjadi induk halaman; fitur ini hanya mendefinisikan subset pengaturan yang memengaruhi service project dan harus diimplementasikan sebagai bagian dari halaman JSM-AI-006, bukan halaman terpisah. Prioritas sama (P1 / Fase 5).

## Request type & portal groups

### JSM-PRJ-020 Project settings - Request types

- **Alias inventaris:** Request types
- **Deskripsi:** Request type adalah lapisan customer-facing di atas work type yang menentukan form portal, nama status yang dilihat customer, portal group, dan kanal. Halaman Project settings > Request types adalah pusat untuk melihat, membuat, mengedit, mengelompokkan, dan menghapus request type.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project) mengelola; Agent Layanan dan Customer/Pelapor memakai hasilnya.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar request type dikelompokkan per portal group (tab/section per grup) ditambah grup "Hidden from portal" untuk request type tanpa grup.
  2. Setiap baris menampilkan ikon, nama, deskripsi, work type, portal group, dan indikator restricted (JSM-PRJ-029).
  3. Drag-and-drop urutan request type dalam satu portal group menentukan urutan tampil di portal grup tersebut.
  4. Tombol Create request type menawarkan pilihan blank atau from template (JSM-PRJ-021, JSM-PRJ-023).
  5. Menu aksi per request type: Edit (tab Request form, Issue view, Workflow statuses, Access), Change portal groups, dan Delete.
  6. Daftar dapat difilter berdasarkan work category (ITSM) dan dicari berdasarkan nama.
  7. ⚠ Admin dapat menetapkan default request type per work type; work item yang dibuat di luar portal (dialog Create internal, import, API, automation) tanpa request type otomatis mendapat default ini.
  8. ⚠ Work item dengan request type tidak dikenal/kosong ditangani lewat default request type per work type saat dibuat; work item lama tidak diubah otomatis dan ditangani lewat JSM-PRJ-054.
  9. Menghapus request type meminta konfirmasi yang menyebut jumlah request terdampak; work item yang ada tetap tersimpan tetapi field Request type menjadi kosong sehingga tidak tampil di portal customer.
  10. Work type yang dipakai request type harus ada di issue type scheme project; dropdown work type hanya menampilkan work type di scheme, dengan tautan untuk menambah work type ke scheme terlebih dahulu.
  11. ⚠ Mengubah work type pada request type yang sudah ada hanya mengubah request baru; request lama tetap memakai work type lamanya dan ditandai tidak konsisten di daftar audit.
  12. ⚠ Aksi Duplicate request type tidak disediakan di rilis awal; bila nanti ditambahkan, salinan memuat request form, hidden fields, dan status customer-facing.
  13. ⚠ Jumlah request type per project dibatasi oleh konfigurasi server (default diusulkan 200); saat batas tercapai tombol Create nonaktif dengan pesan galat.
  14. Pembuatan request type mengisi nama, deskripsi, ikon (library atau upload), dan work type dasar (rinci di JSM-PRJ-021).
  15. Pemetaan work type pada company-managed memilih dari issue type scheme, sedangkan team-managed membuat work type baru bersama request type (hanya rujukan, lihat JSM-PRJ-012).
  16. Request form memilih field dan mengubah display name serta help text khusus customer tanpa mengubah nama field asli (JSM-PRJ-025).
  17. Field wajib di portal dapat berbeda dari field configuration; request form dapat mewajibkan field yang opsional di field configuration.
  18. Hidden fields dengan preset values (mis. Priority, Components, Labels) dikonfigurasi per request type (JSM-PRJ-026).
  19. Layout issue view agent dikonfigurasi per request type: field di description area dan context panel (JSM-PRJ-027).
  20. Workflow statuses memetakan nama status internal ke nama status customer-facing per request type (JSM-PRJ-028).
  21. Request type dapat di-assign ke satu atau lebih portal group; tanpa grup, request type tersembunyi dari portal dan hanya dipakai via email/agent/API.
  22. Request type dapat dipilih sebagai request type untuk email channel per alamat email (JSM-PRJ-063).
  23. Urutan request type dalam grup diatur dengan drag-and-drop dan juga dapat diubah dengan keyboard (aksesibel).
  24. Forms dapat dilampirkan per request type (JSM-PRJ-041).
  25. Nama, deskripsi, field, dan help text dapat diterjemahkan lewat Language support (JSM-PRJ-066).
  26. ⚠ Request lama dari request type yang dihapus menampilkan label "No request type" di issue view agent dan nilai kosong di JQL.
  27. Mengubah request type pada request yang ada (Change request type) memetakan field dan status customer-facing ke request type baru (JSM-PRJ-054).
  28. ⚠ Batas jumlah request type per project berlaku sama untuk UI dan API (lihat butir 13).
  29. Field khusus seperti Attachment, Request participants, dan Share with/Organizations hanya muncul di portal bila ditambahkan ke request form.
  30. Template ITSM membawa contoh request type bawaan seperti Get IT help, Report a system problem, Request a change, Request new hardware/software (JSM-PRJ-003).
- **Team-managed vs company-managed:** Clone hanya company-managed: request type selalu menunjuk work type yang ada di issue type scheme project.
- **Catatan & ketergantungan:** Usulan tabel `request_types` (id, project_id, work_type_id, name, description, icon_id, position_default, deleted_at) dan `request_type_group_members` (request_type_id, portal_group_id, position). Field `Customer Request Type` disimpan sebagai field sistem (JSM-PRJ-060). API publik request type di modul 35.

### JSM-PRJ-021 Create request type dialog

- **Deskripsi:** Dialog untuk membuat request type baru dengan identitas portal dan mapping ke work type.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project).
- **Aturan perilaku (acceptance criteria):**
  1. Name wajib, tampil di portal, maksimal 255 karakter.
  2. Description opsional, tampil di portal di bawah nama request type, maksimal 2.000 karakter teks polos.
  3. Ikon dipilih dari library ikon atau diupload custom (JSM-PRJ-022); default ikon generik bila tidak dipilih.
  4. Work type wajib dipilih dan menentukan workflow, screen, dan field yang tersedia.
  5. Portal groups dapat dipilih satu atau lebih; tanpa grup, request type tersembunyi dari portal tetapi tetap dapat dipakai via email/API/agent, dan dialog menampilkan keterangan ini.
  6. Work category (ITSM) mengikuti work type yang dipilih dan ditampilkan read-only.
  7. Dialog dapat dimulai dari template (Create from template) atau blank.
  8. ⚠ Pembuatan request type berbantuan AI (dari deskripsi bebas menjadi request type beserta usulan field) hanya tampil bila fitur AI opsional aktif dan plan mendukung; hasilnya selalu berupa draft yang ditinjau admin.
  9. Validasi: nama wajib dan unik dalam project (case-insensitive); galat "Nama request type sudah dipakai" ditampilkan inline.
  10. Setelah dibuat, admin diarahkan ke tab Request form request type tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk clone.
- **Catatan & ketergantungan:** AI opsional di modul 35.

### JSM-PRJ-022 Request type icons

- **Deskripsi:** Pengelolaan ikon visual request type yang tampil di portal, help center, widget, dan issue view.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Library ikon bawaan produk (hardware, akses, pertanyaan, dll.) tersedia untuk semua project; ikon bawaan adalah aset milik produk ini, bukan aset merek pihak lain.
  2. ⚠ Upload ikon custom menerima PNG, JPG, GIF, dan SVG ter-sanitasi dengan batas ukuran 1 MB dan dimensi minimal 48x48 px; file di luar batas ditolak dengan pesan galat.
  3. Ikon custom yang diupload disimpan per project dan dapat dipakai ulang di request type lain dalam project tersebut.
  4. Ikon dapat diganti kapan saja tanpa memengaruhi work item.
  5. Ikon tampil di portal, hasil pencarian help center, widget, dan badge request type di issue view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan file CMT-019 ([08](08-komentar-lampiran-aktivitas.md)).

### JSM-PRJ-023 Request type templates

- **Deskripsi:** Galeri template request type siap pakai yang membawa nama, deskripsi, field, dan form yang direkomendasikan.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Galeri dikelompokkan per kategori: IT, HR, Facilities, Legal, Finance, Marketing, dan lainnya.
  2. Admin dapat melihat preview field/form sebelum menambahkan template.
  3. Template dapat membawa form (JSM-PRJ-034) bawaan yang ikut dibuat di project.
  4. Saat dipakai, template dipetakan ke work type yang ada di project; admin wajib memilih work type bila tidak ada kecocokan otomatis.
  5. Request type hasil template sepenuhnya dapat diedit dan tidak lagi terkait ke template.
  6. Request type bawaan project juga berasal dari project template saat project dibuat (JSM-PRJ-001), memakai katalog yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog disimpan sebagai data seed berversi.

### JSM-PRJ-024 Request type portal groups

- **Alias inventaris:** Portal groups
- **Deskripsi:** Pengelompokan request type ke kategori (portal groups) yang tampil sebagai navigasi di portal customer agar customer mudah menemukan jenis permintaan. Di modul ini fitur ini adalah halaman pengaturan Project settings > Request types (penugasan request type ke grup, urutan, kelola grup); pemilik perilaku portal group (model data, aturan tampil di portal, help resources) adalah JSM-PRT-022 Portal groups di [30](30-jsm-help-center-portal-pelanggan-kanal.md).
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengelola; Customer/Pelapor melihat di portal.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat membuat, mengganti nama, dan menghapus portal group; nama wajib, unik per project, maksimal 100 karakter.
  2. Urutan portal group diatur dengan drag-and-drop dan menentukan urutan di portal.
  3. Satu request type dapat masuk ke beberapa portal group.
  4. Request type tanpa portal group tersembunyi dari portal tetapi tetap dapat dipakai email/API/agent.
  5. Grup yang tidak memiliki request type terlihat (setelah access restriction diterapkan untuk customer tersebut) tidak tampil di portal.
  6. ⚠ Portal tanpa grup sama sekali menampilkan daftar request type datar.
  7. Nama grup dapat diterjemahkan lewat Language support (JSM-PRJ-066; perilaku terjemahan portal di JSM-PRT-007).
  8. ⚠ Jumlah portal group per project dibatasi konfigurasi server (default diusulkan 50).
  9. Menghapus grup tidak menghapus request type di dalamnya; request type yang tidak lagi punya grup berpindah ke "Hidden from portal" dan admin diperingatkan sebelum konfirmasi.
  10. Satu request type yang berada di banyak grup tampil di setiap grup tersebut di portal.
  11. Urutan grup dan urutan request type per grup disimpan terpisah, sehingga request type yang sama dapat berada di posisi berbeda di tiap grup.
  12. Grup kosong tidak tampil di portal tetapi tetap terlihat di halaman pengaturan dengan label "Kosong".
  13. Portal menampilkan grup sebagai navigasi samping hanya bila ada lebih dari satu grup yang terlihat.
  14. Terjemahan nama grup memakai fallback ke default language bila kosong.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik perilaku: JSM-PRT-022 Portal groups ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); butir 1-6, 10, 12, 13 adalah sisi pengaturan dari aturan yang sama (JSM-PRT-022 butir 1-6) dan bila terjadi perbedaan, JSM-PRT-022 yang berlaku. Help resources sebagai item grup hanya didefinisikan di JSM-PRT-022 butir 8-11. Tabel `portal_groups` (id, project_id, name, position) dan tabel urutan item per grup (`position` per pasangan grup-item) didefinisikan sekali dan dipakai bersama oleh halaman pengaturan ini dan render portal di modul 30.

### JSM-PRJ-025 Request form configuration (fields)

- **Deskripsi:** Konfigurasi field yang dilihat customer saat mengajukan request melalui portal/widget untuk setiap request type.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Tab Request form menampilkan daftar Visible fields yang tampil di portal sesuai urutan.
  2. Admin dapat menambah field yang tersedia pada screen/field configuration work type request type tersebut.
  3. Admin dapat meng-override "Display name" field per request type tanpa mengubah nama field asli.
  4. Admin dapat mengisi help text field per request type (maks. 1.000 karakter, tampil di bawah field di portal).
  5. Toggle Required dapat diatur per request type; Summary selalu ada di form, wajib, dan tidak dapat dihapus.
  6. Urutan field diatur dengan drag-and-drop.
  7. Preview menampilkan tampilan portal persis seperti dilihat customer tanpa membuat request.
  8. ⚠ Field yang tidak didukung portal (mis. Assignee, Sprint, Reporter, field sistem internal, time tracking) tidak muncul di daftar field yang dapat ditambahkan.
  9. Field yang didukung meliputi Attachment, Components, Due date, Priority, Labels, custom field teks/select/checkbox/date/number/user/URL, dan Assets object field.
  10. Customer melihat nilai field visible pada halaman detail request di portal; field hidden dan internal tidak pernah dikirim ke klien portal.
  11. Halaman menampilkan peringatan bila ada field yang required di field configuration tetapi tidak ada di form dan tidak punya preset, karena pembuatan request akan gagal.
  12. Admin dapat menambahkan form (Forms) ke request form (JSM-PRJ-041).
  13. Tersedia tautan ke screen/field configuration untuk menambah field baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk clone.
- **Catatan & ketergantungan:** Field configuration CFG-043, screen CFG-046 ([05](05-konfigurasi-work-item-skema.md)). Tabel usulan `request_type_fields` (request_type_id, field_id, display_name, help_text, required, visible, preset_value, position).

### JSM-PRJ-026 Hidden fields with preset values

- **Deskripsi:** Field yang tidak tampil ke customer tetapi diisi otomatis dengan nilai preset saat request dibuat.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Tab Request form memiliki bagian Hidden fields terpisah dari Visible fields.
  2. Admin dapat memindahkan field dari visible ke hidden dan sebaliknya; nilai preset tetap tersimpan saat dipindah kembali.
  3. Admin dapat mengatur nilai preset (mis. Priority, Component, Labels, custom field select).
  4. Nilai preset diterapkan saat request dibuat sehingga dapat dipakai untuk routing queue, SLA, dan automation.
  5. Field required yang di-hidden wajib diberi nilai preset; simpan ditolak dengan galat "Field wajib yang disembunyikan harus memiliki nilai preset".
  6. ⚠ Tipe field yang tidak mendukung preset (mis. Attachment, field read-only/terkomputasi) tidak dapat dipindahkan ke hidden.
  7. Preset berlaku untuk semua kanal yang memakai request type tersebut (portal, email, widget, chat, API).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Preset tidak menimpa nilai yang dikirim eksplisit oleh agent/API bila field tersebut juga dikirim.

### JSM-PRJ-027 Request type issue view configuration

- **Deskripsi:** Konfigurasi layout field yang dilihat agent pada issue view untuk request type tertentu.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Agent Layanan dan collaborator melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Request type memiliki tab Issue view.
  2. Layout terdiri dari Description fields (area utama) dan Context fields (panel samping).
  3. Admin dapat menambah, menghapus, dan mengurutkan field per request type.
  4. Admin dapat mengaktifkan opsi menyembunyikan field saat nilainya kosong.
  5. Layout berbeda per request type walaupun work type sama.
  6. Layout berlaku untuk agent dan collaborator; field yang tidak ada di layout tetap dapat diakses lewat "More fields".
  7. Preview menampilkan tampilan issue view dengan data contoh.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Layout issue view inti ISS-007 ([06](06-work-item-inti.md)).

### JSM-PRJ-028 Request type workflow statuses (customer-facing)

- **Deskripsi:** Pemetaan status workflow internal ke nama status yang dilihat customer di portal dan notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Request type memiliki tab Workflow statuses.
  2. Tab menampilkan daftar status dari workflow work type terkait.
  3. Admin dapat mengubah "Status name shown to customers" per status (mis. "Waiting for support" menjadi "Being reviewed"); default sama dengan nama status internal; maks. 60 karakter.
  4. Beberapa status internal dapat dipetakan ke nama customer yang sama.
  5. Nama customer-facing dipakai di portal, notifikasi email customer, dan filter request customer.
  6. Status berkategori Done menandai request sebagai closed/resolved di portal.
  7. Tab menyediakan tautan ke workflow editor untuk mengubah workflow (WF-021).
  8. Nama status customer-facing dapat diterjemahkan per bahasa (JSM-PRJ-066).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status category WF-002 ([09](09-workflow-status-transisi.md)). Status baru yang ditambahkan ke workflow otomatis memakai nama internalnya sampai diubah.

### JSM-PRJ-029 Request type access restrictions

- **Deskripsi:** Membatasi customer/pengguna yang dapat melihat dan mengajukan request type tertentu di portal.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi akses: "Anyone who can access the portal" (default) atau "Only specific people".
  2. Pada "Only specific people", admin memilih individual customer/pengguna, groups, dan/atau organizations.
  3. Request type terbatas tidak tampil bagi yang tidak berhak di portal, pencarian help center, dan widget; akses langsung lewat URL mengembalikan halaman "tidak ditemukan".
  4. Daftar request type (JSM-PRJ-020) menampilkan indikator restricted.
  5. ⚠ Pembatasan tidak berlaku untuk agent (agent dapat raise on behalf memakai request type terbatas); request via email/API memakai request type tersebut tetap diterima bila kanal dikonfigurasi memakainya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ketersediaan per plan diatur di JSM-PRJ-070. Organizations di modul 30.

### JSM-PRJ-030 Request type per channel

- **Deskripsi:** Penentuan request type yang dipakai atau tampil pada tiap kanal (portal, email, widget, chat, virtual service agent).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Portal menampilkan request type yang berada di portal group (JSM-PRJ-024).
  2. Email memakai satu request type default per alamat email; request type itu tidak boleh memiliki field wajib selain Summary/Description/Attachment kecuali field tersebut punya preset.
  3. Widget menampilkan request type yang dipilih admin di pengaturan widget (JSM-PRJ-064).
  4. Chat (Slack/Microsoft Teams) memakai request type yang dipilih untuk request dari chat (JSM-PRJ-065).
  5. Virtual service agent memakai request type tujuan handoff/intent (modul 35).
  6. Penyimpanan konfigurasi kanal memvalidasi kompatibilitas field required; request type yang tidak kompatibel tidak dapat dipilih dan alasannya ditampilkan.
  7. Kanal asal tercatat di work item (request channel: portal, email, widget, chat, api, jira, anonymousportal) dan tidak dapat diubah setelah dibuat; nilai enum `jira` berarti dibuat lewat dialog Create internal dan dipertahankan demi kompatibilitas API/impor, sedangkan label UI-nya "Dibuat internal".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail kanal di modul 30.

## Workflow, resolusi & otomasi bawaan

### JSM-PRJ-031 Service project workflows

- **Alias inventaris:** Default service workflows & customer-facing statuses
- **Deskripsi:** Workflow khas Service Management yang dipakai request type melalui work type, beserta status, resolusi, dan pemetaan ke status yang dilihat customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek dan Admin Situs (mengedit workflow scheme); Agent Layanan (transisi); Customer/Pelapor (melihat status customer-facing).
- **Aturan perilaku (acceptance criteria):**
  1. Workflow bawaan yang disediakan sebagai seed: Service request fulfilment, Service request fulfilment with approvals, Incident management, Problem management, Change management, dan workflow general service.
  2. Status umum yang disediakan: Waiting for support, Waiting for customer, In progress, Pending, Escalated, Waiting for approval, Resolved, Closed, Canceled, dengan status category yang benar (Resolved/Closed/Canceled = Done).
  3. Langkah approval dikonfigurasi pada status workflow (field approver, jumlah approval, transisi approve/decline) dan detailnya di modul 32.
  4. Workflow per work type berlaku untuk semua request type dengan work type tersebut.
  5. Transisi otomatis berbasis komentar dijalankan oleh default automation (JSM-PRJ-033), bukan hard-coded di engine workflow.
  6. Status customer-facing dipetakan per request type (JSM-PRJ-028).
  7. Modifikasi workflow memakai workflow editor inti (WF-021) tanpa editor khusus.
  8. ⚠ Status default bervariasi per template: General memakai Waiting for support, Waiting for customer, In progress, Resolved, Closed; ITSM menambah Pending, Escalated, Canceled.
  9. Workflow ITSM terpisah untuk Service request, Incident, Problem, Change, dan Post-incident review.
  10. Resolution di-set pada transisi ke status selesai, dan nilainya memengaruhi SLA dan CSAT (JSM-PRJ-032).
  11. Nama status customer-facing dikonfigurasi per request type, bukan di workflow (JSM-PRJ-028).
  12. Komentar customer memindahkan Waiting for customer ke Waiting for support melalui automation rule bawaan.
  13. ⚠ Transisi customer dari portal (tombol Resolve/Cancel request) tersedia bila workflow memiliki transisi yang ditandai "customer can trigger" dan request berstatus terbuka.
  14. Company-managed: workflow diedit lewat workflow scheme; team-managed (rujukan, JSM-PRJ-012) mengedit workflow per request type/work type.
- **Team-managed vs company-managed:** Clone hanya company-managed (workflow scheme).
- **Catatan & ketergantungan:** Engine workflow di [09](09-workflow-status-transisi.md); approvals di modul 32. Properti transisi "customer can trigger" adalah keputusan desain clone untuk butir 13.

### JSM-PRJ-032 Request resolution handling

- **Deskripsi:** Penanganan resolution pada request yang memengaruhi SLA, portal, dan pelaporan.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (transisi), Customer/Pelapor (resolve/cancel bila diizinkan), Admin Proyek (konfigurasi).
- **Aturan perilaku (acceptance criteria):**
  1. Field Resolution di-set lewat transition screen atau post function "Set resolution" pada transisi ke status Done.
  2. Status Done tanpa resolution tetap dianggap unresolved; halaman settings workflow menampilkan peringatan bila transisi ke status Done tidak mengisi resolution.
  3. SLA berhenti berdasarkan kondisi "Resolution set" atau status (modul 31).
  4. Request resolved tampil di daftar "Closed requests" portal customer.
  5. ⚠ Customer dapat menyelesaikan atau membatalkan request sendiri dari portal bila workflow menyediakan transisi customer (JSM-PRJ-031 butir 13); resolusi yang di-set adalah "Done" atau "Cancelled" sesuai transisi.
  6. Request dibuka kembali (re-open) saat customer berkomentar setelah resolved, melalui default automation; resolution dikosongkan saat re-open.
  7. Request resolved ditutup otomatis setelah beberapa hari (default 3 hari kerja, dapat diubah di rule) melalui default automation.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semantik resolution WF-009 dan ISS-019; resolutions administration CFG-055.

### JSM-PRJ-033 Default service project automation rules

- **Deskripsi:** Rule automation bawaan yang dibuat bersama service project untuk alur komunikasi customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengelola rule project).
- **Aturan perilaku (acceptance criteria):**
  1. Rule "Transition when customer comments": customer berkomentar memindahkan status Waiting for customer ke Waiting for support.
  2. Rule "Transition when agent replies": agent membalas customer (Reply to customer) memindahkan status ke Waiting for customer; internal note tidak memicu rule ini.
  3. Rule "Re-open on customer comment": request resolved dibuka kembali saat customer berkomentar.
  4. Rule "Close resolved requests": request resolved ditutup setelah sejumlah hari yang dapat dikonfigurasi.
  5. ⚠ Rule bawaan lain (triage otomatis/auto-assign) dibuat dalam keadaan nonaktif dan dapat diaktifkan admin.
  6. Semua rule bawaan dapat dinonaktifkan dan diedit di halaman Automation project; menghapus rule bawaan meminta konfirmasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine automation di [20](20-otomasi.md) (AUTO-001, AUTO-012, AUTO-042). Trigger "Comment added" harus membedakan komentar publik vs internal dan pengarang customer vs agent.

## Forms

### JSM-PRJ-034 Forms (Jira Forms, formerly ProForma)

- **Alias inventaris:** Forms (JSM request forms)
- **Deskripsi:** Form builder native (setara Forms eks ProForma) untuk mengumpulkan data terstruktur yang kaya pada request dan work item, menempel ke request type dan request dengan logika kondisional. Engine form dipakai bersama Forms kerja bisnis (modul 26).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengelola form; Agent Layanan menambah/mengisi form di work item; Customer/Pelapor mengisi form eksternal; pengguna berlisensi mengisi form di dialog Create internal (ISS-002).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Project settings > Forms menampilkan daftar form dengan pencarian, serta info request type/work type yang memakai setiap form.
  2. Create form dapat dimulai blank atau dari template (JSM-PRJ-040).
  3. Form dapat diedit, di-copy/duplicate, dan dihapus; menghapus form yang sedang dipakai request type menampilkan peringatan berisi daftar request type terdampak.
  4. ⚠ Form yang sudah terlampir di work item tetap tersimpan (snapshot struktur dan jawaban) walaupun form sumber dihapus.
  5. ⚠ Form dapat di-copy ke service project lain yang dikelola pengguna; linked field yang tidak tersedia di project tujuan diputus dan dilaporkan.
  6. ⚠ Form settings memuat nama, request type yang memakai form di portal, work type yang direkomendasikan, dan opsi menambahkan form otomatis di dialog Create internal (ISS-002).
  7. ⚠ Form share link menghasilkan tautan langsung ke form untuk membuat work item/request tanpa menavigasi portal; work type/request type tujuan mengikuti pengaturan form; audiens link dibatasi (customer login atau pengguna berlisensi) dan link dapat dinonaktifkan atau diregenerasi (link lama langsung tidak berlaku).
  8. ⚠ Admin dapat menambahkan terjemahan form untuk bahasa yang didukung project (JSM-PRJ-066).
  9. Satu form dapat dihubungkan ke beberapa request type.
  10. Satu work item dapat memiliki beberapa form.
  11. Permission: Admin Proyek mengelola form; agent menambah/mengisi form di work item; customer hanya mengisi form berstatus external.
  12. ⚠ Batas default 20 form per work item dan 500 pertanyaan per form; melewati batas ditolak dengan pesan galat.
  13. ⚠ Form builder drag & drop mendukung tipe text, paragraph, choice (radio, checkbox, dropdown), date, time, number, user, attachment, Assets object, dan table (rinci di JSM-PRJ-035).
  14. Form mendukung sections dan conditional logic berdasarkan jawaban (JSM-PRJ-037).
  15. Linked fields menghubungkan pertanyaan ke field work item agar nilainya tersinkron (JSM-PRJ-038).
  16. Validasi mencakup wajib, min/max, regex, dan rentang tanggal (JSM-PRJ-036).
  17. Form dapat dilampirkan ke request type agar tampil di portal saat request dibuat (JSM-PRJ-041).
  18. Agent dapat menambah form ke request setelah dibuat, lalu meminta customer mengisinya dengan mengubah visibilitas menjadi external (JSM-PRJ-042).
  19. Status form: open, submitted, locked, dan reopen oleh agent (JSM-PRJ-043).
  20. Form templates bawaan tersedia (JSM-PRJ-040).
  21. Form dapat diekspor ke PDF/XLSX (JSM-PRJ-044).
  22. ⚠ Data connections ke sumber eksternal untuk opsi dropdown dinamis tersedia sesuai plan (JSM-PRJ-039).
  23. Visibilitas form per work item adalah external (terlihat di portal) atau internal-only.
  24. Linked fields hanya dapat menunjuk field yang tersedia di project; pada company-managed daftarnya mengikuti field configuration dan screen project.
- **Team-managed vs company-managed:** Pada company-managed linked fields mengikuti field configuration bersama; varian team-managed hanya rujukan (JSM-PRJ-012).
- **Catatan & ketergantungan:** Fitur ProForma Marketplace lama di luar Forms native tidak dibangun. Model data usulan: `forms` (id, project_id, name, version, settings_json), `form_versions` (struktur JSON), `form_instances` (work_item_id, form_id, form_version, status, visibility, answers_json, submitted_by, submitted_at). Engine bersama BIZ-029 s.d. BIZ-041 ([26](26-kerja-bisnis-formulir-persetujuan.md)). Keputusan: produk memakai SATU engine form (setara ProForma) bersama modul 26. Pemilik skema form (definisi pertanyaan/field, kondisi, validasi) adalah BIZ-030 Form builder dan BIZ-033 Form conditional logic ([26](26-kerja-bisnis-formulir-persetujuan.md)), yang dikirim lebih dulu di Fase 3; skema kanonis mengikuti BIZ-030 (`forms`, `form_questions`), sedangkan modul ini menambahkan `form_versions` (snapshot struktur saat Save), `form_instances` (form yang menempel di request/work item), dan tabel binding `request_type_forms` (request_type_id, form_id, posisi, opsi portal). Model data di atas adalah perluasan, bukan skema kedua.

### JSM-PRJ-035 Form builder

- **Deskripsi:** Editor drag-and-drop untuk menyusun pertanyaan, konten, dan layout form.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Elemen konten tersedia: heading, paragraf rich text, list, dan divider; image dan tabel layout disediakan bila editor konten kaya (modul 07) mendukungnya.
  2. ⚠ Field teks tersedia: short text, paragraph/long text, dan rich text.
  3. Field Number, Email, URL, Date, Time, dan Date & time tersedia.
  4. Field pilihan tersedia: Dropdown, Multi-select dropdown, Checkboxes, Radio buttons.
  5. User picker tersedia dalam mode single dan multiple.
  6. Attachment field tersedia dengan batas ukuran mengikuti setting attachment site (CMT-016).
  7. ⚠ Assets object field dengan filter AQL tersedia bila modul 33 aktif.
  8. Sections mengelompokkan pertanyaan dan menjadi dasar conditional logic.
  9. Setiap pertanyaan memiliki label (wajib), deskripsi, placeholder, dan default value.
  10. Opsi pilihan dapat diisi manual, diambil dari opsi field (bila linked), atau dari data connection.
  11. Preview menampilkan form seperti dilihat pengisi, termasuk logika kondisional.
  12. Perubahan disimpan eksplisit dengan tombol Save; meninggalkan builder dengan perubahan belum tersimpan memunculkan dialog konfirmasi.
  13. ⚠ Pembuatan form berbantuan AI (dari deskripsi menjadi draft pertanyaan) hanya tersedia bila fitur AI opsional aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Editor rich text di [07](07-editor-konten-kaya.md); setiap Save membuat `form_versions` baru sehingga instance lama tetap konsisten. Mesin form bersama dibangun di BIZ-030/BIZ-033 ([26](26-kerja-bisnis-formulir-persetujuan.md)); fitur ini menambah binding request type/portal, section, template, dan versioning form di atas editor yang sama. Pemilik skema form (field, kondisi, validasi) adalah BIZ-030; perubahan skema yang dibutuhkan Service Management diajukan ke BIZ-030, bukan diimplementasikan sebagai editor kedua.

### JSM-PRJ-036 Form validation rules

- **Deskripsi:** Aturan validasi pada pertanyaan form untuk menjamin kualitas data sebelum submit.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor dan agent terkena validasi.
- **Aturan perilaku (acceptance criteria):**
  1. Pertanyaan dapat ditandai Required.
  2. Teks dapat dibatasi min/max jumlah karakter atau kata.
  3. ⚠ Teks dapat divalidasi dengan regex/pattern beserta pesan galat custom; regex dievaluasi dengan engine aman (RE2) dan batas waktu.
  4. Number dapat dibatasi min/max nilai, dan Date dapat dibatasi rentang tanggal.
  5. Format Email dan URL divalidasi otomatis.
  6. Checkbox/multi-select dapat dibatasi jumlah minimum/maksimum pilihan.
  7. Saat submit gagal, pesan galat inline tampil di setiap pertanyaan bermasalah dan fokus berpindah ke galat pertama.
  8. Validasi diabaikan untuk pertanyaan di section yang tersembunyi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validasi dijalankan ulang di server; validasi klien hanya untuk UX.

### JSM-PRJ-037 Form conditional logic

- **Deskripsi:** Menampilkan atau menyembunyikan section berdasarkan jawaban pada pertanyaan pilihan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; pengisi form melihat hasilnya.
- **Aturan perilaku (acceptance criteria):**
  1. Section dapat dijadikan conditional section yang tampil hanya bila kondisinya terpenuhi.
  2. Kondisi berbasis jawaban pertanyaan Dropdown, Radio, atau Checkbox yang berada sebelum section tersebut.
  3. ⚠ Operator kondisi: "salah satu dari opsi tertentu" dan "bukan salah satu dari opsi tertentu".
  4. Satu section dapat memiliki beberapa kondisi yang digabung dengan AND.
  5. Pertanyaan di section tersembunyi tidak wajib diisi dan nilainya tidak disimpan saat submit.
  6. Preview dapat dipakai untuk menguji logika dengan mengubah jawaban secara interaktif.
  7. Menghapus pertanyaan yang menjadi sumber kondisi menampilkan peringatan dan, bila dikonfirmasi, menghapus kondisi terkait.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Evaluasi kondisi identik di klien dan server. Mesin form bersama dibangun di BIZ-030/BIZ-033 ([26](26-kerja-bisnis-formulir-persetujuan.md)); fitur ini menambah binding request type/portal. Keputusan urutan: engine kondisi dibangun SATU kali sebagai bagian dari BIZ-033 di paket form bersama, sehingga JSM-PRJ-037 (Fase 5) bergantung pada BIZ-033 (Fase 4) dan hanya mengaktifkan kondisi untuk form request type dan portal. Bila BIZ-033 ditunda karena ketersediaannya di form kerja bisnis belum terkonfirmasi, engine kondisi tetap dibangun di paket form bersama sebagai bagian JSM-PRJ-037 dan BIZ-033 kemudian cukup mengaktifkannya untuk form bisnis. Rekomendasi ke pemilik modul 26: naikkan BIZ-033 ke P1 agar prioritas selaras dengan fitur ini (lihat Pertanyaan terbuka).

### JSM-PRJ-038 Form linked Jira fields

- **Deskripsi:** Menghubungkan pertanyaan form dengan field work item agar jawaban tersimpan ke field tersebut.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Pertanyaan dapat di-link ke field sistem/custom (Summary, Description, Priority, Due date, Labels, custom field, dll.).
  2. Tipe pertanyaan harus kompatibel dengan tipe field; field yang tidak kompatibel tidak muncul di pilihan link.
  3. Nilai field work item diisi saat form disubmit.
  4. ⚠ Perubahan field work item tercermin kembali ke form yang belum dikunci (sinkron dua arah); pada form Locked, form menyimpan nilai saat submit.
  5. Opsi pilihan dapat diambil dari opsi field yang di-link.
  6. Field ter-link dapat dipakai di JQL, queue, SLA, dan automation; pertanyaan tidak ter-link hanya tersimpan sebagai data form.
  7. ⚠ Bila field yang sama muncul di request form standar dan sebagai pertanyaan form ter-link, field standar disembunyikan dari portal dan builder menampilkan peringatan duplikasi.
  8. ⚠ Summary request dapat diisi dari pertanyaan ter-link; jika ada, field Summary standar disembunyikan di portal.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk clone.
- **Catatan & ketergantungan:** Tipe field di CFG-025 s.d. CFG-031 ([05](05-konfigurasi-work-item-skema.md)).

### JSM-PRJ-039 Form data connections

- **Deskripsi:** Mengisi opsi pertanyaan pilihan secara dinamis dari sumber data eksternal.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (data connection project); Admin Situs (bila diputuskan global).
- **Aturan perilaku (acceptance criteria):**
  1. Data connections dikelola di pengaturan Forms (daftar, buat, edit, hapus, uji koneksi).
  2. Sumber berupa REST API/URL HTTPS dengan respons JSON; URL ke alamat jaringan privat/internal ditolak (perlindungan SSRF).
  3. Admin memetakan path JSON ke label dan value opsi, dan hasil pemetaan dipratinjau sebelum disimpan.
  4. ⚠ Autentikasi didukung lewat header custom atau bearer token yang disimpan terenkripsi dan tidak pernah ditampilkan ulang.
  5. ⚠ Data di-cache dengan TTL yang dapat diatur (default 1 jam) dan dapat di-refresh manual.
  6. Data connection dapat dipakai pada pertanyaan Dropdown, Checkbox, dan Radio.
  7. Saat sumber tidak tersedia, pertanyaan memakai cache terakhir bila ada; bila tidak ada, pertanyaan menampilkan pesan galat dan form tetap dapat disubmit kecuali pertanyaan itu wajib.
  8. Untuk data CMDB, builder menyarankan Assets object field (modul 33) sebagai alternatif native.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ruang lingkup project vs global menjadi pertanyaan terbuka.

### JSM-PRJ-040 Form templates

- **Deskripsi:** Library template form siap pakai untuk berbagai use case.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Galeri template dikelompokkan per kategori (IT, HR, Facilities, Legal, Finance, Marketing, dll.).
  2. Template dapat dipratinjau sebelum dipakai.
  3. Form yang dibuat dari template dapat diedit bebas dan tidak lagi terikat ke template.
  4. Template form ikut dibawa oleh request type template (JSM-PRJ-023).
  5. ⚠ Admin dapat menyimpan form sebagai template organisasi yang muncul di galeri semua project site.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog template sebagai data seed berversi.

### JSM-PRJ-041 Forms on portal (request type forms)

- **Deskripsi:** Melampirkan form ke request type sehingga customer mengisi form saat mengajukan request di portal.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih form pada konfigurasi request form request type (JSM-PRJ-025).
  2. Form tampil bersama field request form standar, di bawah field standar.
  3. ⚠ Summary tetap diperlukan dan dapat berasal dari field standar, linked field, atau preset; bila tidak ada satupun, konfigurasi ditolak.
  4. Form tersubmit otomatis saat request dibuat, dan validasi form dijalankan sebelum request dibuat.
  5. Customer melihat form yang sudah disubmit di halaman request portal (read-only).
  6. Customer dapat mengedit form yang di-reopen oleh agent melalui portal dan men-submit ulang.
  7. Form portal berstatus external (terlihat customer).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Render portal di modul 30.

### JSM-PRJ-042 Forms on issue view

- **Deskripsi:** Pengelolaan form langsung dari issue view oleh agent.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Edit work item); collaborator melihat form internal.
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat menambah form ke work item dengan memilih dari form project; form yang direkomendasikan untuk request type/work type tampil di atas.
  2. Panel Forms menampilkan semua form terlampir dengan status masing-masing.
  3. Agent dapat mengisi/mengedit, menyimpan draft, dan submit form.
  4. Agent dapat mengubah visibilitas form internal/external (membagikan ke customer); mengubah ke external memunculkan konfirmasi.
  5. ⚠ Admin dapat membatasi siapa yang dapat melihat form di work item berdasarkan project role atau grup.
  6. Agent dapat menghapus form dari work item dengan konfirmasi; jawaban yang sudah tersinkron ke linked field tidak ikut terhapus.
  7. Penambahan, perubahan, submit, dan penghapusan form tercatat di history/activity (CMT-022).
  8. ⚠ Form dapat ditambahkan saat membuat work item di dialog Create bila form settings mengaktifkannya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kemampuan collaborator mengedit form menjadi pertanyaan terbuka.

### JSM-PRJ-043 Form lifecycle (submit, lock, reopen)

- **Deskripsi:** Status dan perilaku siklus hidup form pada work item beserta integrasinya dengan workflow dan automation.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (reopen/unlock), Customer/Pelapor (submit), Admin Proyek (konfigurasi).
- **Aturan perilaku (acceptance criteria):**
  1. Form memiliki status Open, Submitted, dan Locked.
  2. Submit memvalidasi semua pertanyaan yang terlihat; submit gagal tidak mengubah status.
  3. Opsi lock saat submit membuat customer tidak dapat mengubah form setelah submit.
  4. Agent dapat reopen/unlock form agar dapat diedit kembali; customer menerima notifikasi bila form external di-reopen.
  5. ⚠ Automation menyediakan trigger "Forms submitted" serta action "Attach form" dan "Copy forms" (AUTO-029).
  6. ⚠ Workflow menyediakan validator/condition berbasis form (form attached / form submitted) yang memblokir transisi bila tidak terpenuhi.
  7. ⚠ Data form tersedia sebagai smart values di automation (mis. jawaban pertanyaan berdasarkan ID).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validator di WF-028 ([09](09-workflow-status-transisi.md)); smart values AUTO-063 ([20](20-otomasi.md)).

### JSM-PRJ-044 Form export

- **Deskripsi:** Mengunduh atau mengekspor isi form untuk arsip atau pelaporan.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek (Browse project dan izin melihat form).
- **Aturan perilaku (acceptance criteria):**
  1. Form dapat diunduh sebagai PDF dari issue view sesuai layout form.
  2. ⚠ Form dapat diunduh sebagai XLSX (satu baris per pertanyaan).
  3. ⚠ Data form lintas work item dapat diekspor (satu baris per work item, kolom per pertanyaan) dari daftar form.
  4. ⚠ Lampiran dalam ekspor dicantumkan sebagai nama file dan tautan, bukan disematkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ekspor umum di [13](13-pencarian-filter-ekspor.md). Ekspor menghormati visibilitas form.

## Pengguna, peran & izin service project

### JSM-PRJ-045 Agents, collaborators & customers (user types)

- **Deskripsi:** Model pengguna Service Management: agent berlisensi, collaborator dengan lisensi produk inti tanpa seat agent, customer tanpa lisensi, serta stakeholder dan approver. Fitur ini menetapkan aturan klasifikasi; detail tiap tipe di JSM-PRJ-046 s.d. JSM-PRJ-048.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (product access), Admin Proyek (role project).
- **Aturan perilaku (acceptance criteria):**
  1. Agent adalah pengguna dengan product access Service Management yang berada di role Service Desk Team; agent dapat public reply, mengubah request, dan di-assign.
  2. Collaborator adalah pengguna berlisensi produk inti tanpa seat Service Management; collaborator dapat melihat request, menambah internal comment, dan menjadi watcher, tetapi tidak bisa public comment atau di-assign.
  3. Customer tidak berlisensi dan tidak dibatasi jumlahnya; customer hanya mengakses portal/email/chat/widget.
  4. ⚠ Stakeholder (incident management) menerima update incident tanpa seat agent; tersedia sesuai plan (modul 32).
  5. Approver dapat berupa customer atau agent; keputusan approval tidak memerlukan seat agent.
  6. Admin Proyek dapat menambah agent dari halaman People/Agents project dan mengundang pengguna baru (JSM-PRJ-049).
  7. Assign request ke pengguna non-agent ditolak dengan galat "Pengguna ini tidak memiliki lisensi agent" di UI maupun API.
  8. Plan Free membatasi maksimal 3 agent per site (JSM-PRJ-070).
  9. Pengguna yang juga customer di project lain dapat memakai portal sebagai customer dan sisi agent sebagai agent dengan satu akun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Product access PERM-004 ([03](03-izin-peran-keamanan.md)). Repo saat ini: tabel `users` belum memiliki tipe akun; perlu kolom `account_type` (`licensed`, `customer`) atau tabel `customer_accounts` terpisah (keputusan di modul 30), serta tabel `product_access` per pengguna.

### JSM-PRJ-046 Agent role and licensing

- **Deskripsi:** Peran agent berlisensi Service Management yang bekerja penuh di service project dan berkomunikasi dengan customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (product access), Admin Proyek (role Service Desk Team), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Agent memerlukan seat Service Management; satu seat berlaku lintas semua service project site.
  2. Agent ditambahkan ke project role Service Desk Team.
  3. Agent dapat memakai queues, membalas customer, menulis internal note, melakukan transisi, mengubah field, mengelola KB, raise request atas nama customer, serta mengelola customer/organization.
  4. Plan Free maksimal 3 agent; plan berbayar dibatasi jumlah seat yang dibeli/dikonfigurasi.
  5. Pengguna tanpa seat Service Management tidak dapat menjadi agent; paling jauh menjadi collaborator.
  6. Admin Situs tidak otomatis menjadi agent tanpa product access Service Management.
  7. Agent ditambah dari halaman People/Access project; undangan pengguna baru memberi product access hanya bila seat tersedia dan admin mengizinkan, jika tidak undangan ditolak dengan pesan "Seat agent habis".
  8. Product access dikelola di administrasi pengguna produk ini (modul 02/22), bukan di project.
  9. Menghapus agent (mencabut product access) membebaskan seat; work item yang di-assign tetap ada dengan assignee yang ditandai "tidak aktif sebagai agent".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Billing tidak dibangun; seat diperlakukan sebagai batas konfigurasi (VISI-003). Pasangan tumpang tindih: JSM-AI-003 Licensing & Agent Seats ([35](35-jsm-ai-otomasi-api-admin-csm.md)). Pemilik: fitur ini memiliki perilaku peran agent di tingkat service project (role Service Desk Team, hak berkomunikasi dengan customer); JSM-AI-003 memiliki penghitungan kursi agent dan batas lisensi tingkat situs. Prioritas sama (P0 / Fase 5).

### JSM-PRJ-047 Collaborator role

- **Deskripsi:** Pengguna berlisensi produk inti (bukan Service Management) yang membantu agent secara internal tanpa berkomunikasi langsung dengan customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Developer/Anggota Tim sebagai collaborator; Agent Layanan; Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Collaborator mendapat akses melalui permission Browse project tanpa role Service Desk Team.
  2. Collaborator dapat melihat work item, menambah internal note, watch, dan di-mention.
  3. Collaborator tidak dapat membalas customer, tidak mengakses queues, dan hanya dapat mengubah field yang diizinkan permission scheme (default: tidak ada edit field request).
  4. Collaborator tidak memakan seat agent.
  5. ⚠ Balasan collaborator via email ke work item diperlakukan sebagai internal note.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komentar internal di JSM-PRJ-053. Pasangan tumpang tindih: JSM-AI-004 Collaborators ([35](35-jsm-ai-otomasi-api-admin-csm.md)). Pemilik: fitur ini adalah versi lengkap peran collaborator; JSM-AI-004 cukup merujuk ke sini. Prioritas sama (P1 / Fase 5).

### JSM-PRJ-048 Customer role (portal-only)

- **Deskripsi:** Customer yang mengajukan dan memantau request lewat portal/email/chat tanpa lisensi.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor; Agent Layanan dan Admin Proyek mengelola.
- **Aturan perilaku (acceptance criteria):**
  1. Customer gratis dan tidak dibatasi jumlahnya.
  2. Customer dapat memakai akun portal-only (hanya untuk help center) atau akun pengguna site.
  3. Customer berada di role Service Desk Customers yang mendapat permission lewat security type portal access.
  4. Customer dapat raise request, melihat request sendiri/yang dibagikan, menambah komentar publik, approve bila menjadi approver, dan memberi rating CSAT.
  5. Customer tidak dapat melihat internal note, queue, atau field yang tidak ditampilkan di request form; API portal tidak mengembalikan data tersebut.
  6. Customer ditambahkan via halaman Customers, email masuk, signup portal, atau organization.
  7. Akun customer dikelola di administrasi akun customer site (modul 30), terpisah dari pengguna berlisensi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Login akun portal-only di [01](01-identitas-autentikasi-sesi.md) dan modul 30.

### JSM-PRJ-049 Project settings - People and access

- **Deskripsi:** Halaman pengelolaan pengguna dan role pada service project.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer project).
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat menambah pengguna ke role Administrators, Service Desk Team (agent), atau role lain.
  2. Admin dapat mengubah atau menghapus role pengguna.
  3. Pengguna tanpa seat Service Management yang ditambahkan ke Service Desk Team ditandai indikator "Bukan agent (tanpa lisensi)" dan diperlakukan sebagai collaborator.
  4. Customer tidak dikelola di halaman ini melainkan di halaman Customers (modul 30), dengan tautan ke sana.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** People page inti PRJ-037 ([04](04-proyek.md)) dan PERM-019.

### JSM-PRJ-050 Service project roles and permissions

- **Alias inventaris:** JSM permission defaults & roles
- **Deskripsi:** Role dan permission scheme khas service project yang mengatur akses agent, collaborator, dan customer, beserta aturan visibilitas khusus customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs (permission scheme).
- **Aturan perilaku (acceptance criteria):**
  1. Role default: Administrators, Service Desk Team, Service Desk Customers.
  2. Permission scheme Service Management dibuat otomatis saat project dibuat (kecuali share settings).
  3. Visibilitas komentar internal vs publik mengikuti role: hanya pemegang Browse project non-customer yang melihat internal note.
  4. ⚠ Permission khusus "Manage customers and organizations" mengatur siapa yang dapat menambah customer/organization ke project.
  5. Halaman project settings mendeteksi masalah permission scheme Service Management (mis. customer diberi Browse project langsung, Service Desk Team kehilangan Assignable User) dan menampilkan peringatan beserta saran perbaikan.
  6. Issue security level dapat diterapkan pada request.
  7. ⚠ Stakeholder untuk update incident diatur sesuai plan (modul 32).
  8. Project role default dibuat dengan nama tetap Administrators, Service Desk Team, Service Desk Customers dan tidak dapat dihapus selama project service ada.
  9. Permission scheme default Service Management tersedia sebagai template scheme di administrasi.
  10. Security type khusus "Service Project Customer - Portal Access" dipakai untuk Browse Projects, Create issues, Add comments, Create attachments, dan sejenisnya; security type ini hanya memberi akses lewat portal.
  11. Customer hanya melihat request yang dia buat, yang dia ikuti sebagai participant, atau yang dibagikan ke organisasinya.
  12. Internal comment tidak pernah tampil ke customer, dan komentar publik ditandai "shared with customer" di issue view.
  13. ⚠ Saat permission scheme diubah sehingga akses customer/agent rusak, sistem menampilkan peringatan sebelum simpan dan di halaman Customer permissions.
  14. ⚠ Varian team-managed (rujukan, JSM-PRJ-012) memakai role Administrator, Agent, Customer dan pengaturan akses project.
  15. Issue security level dapat membatasi agent tertentu dari request sensitif (company-managed).
  16. Permission Assignable User hanya efektif untuk pemegang seat agent; pengguna lain tidak muncul di picker Assignee.
  17. Permission scheme hanya dapat diubah di project company-managed.
- **Team-managed vs company-managed:** Clone hanya company-managed (permission scheme); butir 14 adalah rujukan.
- **Catatan & ketergantungan:** PERM-014, PERM-016, PERM-018, PERM-042 s.d. PERM-045 ([03](03-izin-peran-keamanan.md)). Security type baru perlu ditambahkan ke evaluator permission inti. Pasangan tumpang tindih: JSM-AI-001 JSM Permissions & Role Defaults ([35](35-jsm-ai-otomasi-api-admin-csm.md)). Pemilik: fitur ini adalah versi lengkap role dan permission scheme service project; JSM-AI-001 cukup merujuk ke sini untuk default role dan permission. Prioritas sama (P0 / Fase 5).

### JSM-PRJ-051 Project settings - Customer permissions

- **Deskripsi:** Halaman Project settings > Customer permissions yang mengatur siapa yang dapat mengajukan request dan dengan siapa customer dapat berbagi request. Pemilik perilaku (lapisan otorisasi customer, opsi akses, sharing, akses KB) adalah JSM-PRT-016 Customer permissions (project level) di [30](30-jsm-help-center-portal-pelanggan-kanal.md); fitur ini hanya mendefinisikan halaman pengaturannya di service project.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. "Who can raise requests" memiliki opsi: hanya customer yang ditambahkan ke project, siapa saja dengan akun di site, atau siapa saja di web (public signup).
  2. ⚠ Request anonim tanpa login dapat diizinkan bila setting global "Allow access without login" (JSM-PRT-015) mengizinkannya; request anonim tercatat dengan channel `anonymousportal`.
  3. Customer sharing memiliki opsi: dengan organization, dengan customer lain di project, atau tidak boleh berbagi.
  4. Cakupan pencarian customer saat sharing dibatasi sesuai opsi (customer tidak dapat mencari pengguna di luar cakupan).
  5. Opsi public signup dan anonim nonaktif bila setting global help center melarangnya (JSM-PRT-015 Customer access & signup settings), dengan keterangan alasan.
  6. Pengaturan ini menentukan pemrosesan email dari pengirim tak dikenal: dibuat sebagai customer baru bila "siapa saja" diizinkan, ditolak bila hanya customer terdaftar (JSM-PRT-016 butir 6, JSM-PRT-041).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik perilaku: JSM-PRT-016 Customer permissions (project level) ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); butir 1-6 adalah sisi halaman pengaturan dari JSM-PRT-016 butir 1-4, 6, dan 9, dan bila terjadi perbedaan, JSM-PRT-016 yang berlaku. Setting global yang membatasi opsi di halaman ini adalah JSM-PRT-015 (bukan JSM-PRJ-019; JSM-PRJ-019 hanya menautkan ke halaman admin tersebut). Sharing rinci di JSM-PRT-030; customers & organizations di JSM-PRT-033/JSM-PRT-034; email di JSM-PRJ-063.

## Pengalaman agent di issue view

### JSM-PRJ-052 Agent issue view (service request)

- **Alias inventaris:** Agent request view (JSM issue view)
- **Deskripsi:** Issue view yang diperkaya untuk agent, menyatukan konteks customer, percakapan (balasan ke customer dan internal note), SLA, dan detail request.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (penuh), collaborator (terbatas: internal note, watch).
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan badge request type (ikon + nama) dan channel asal.
  2. Tautan "View request in portal" membuka request persis seperti dilihat customer (hanya komentar publik dan nama status customer-facing); URL portal dapat disalin dan dibagikan.
  3. Agent dapat membagikan request ke customer/participants dari issue view (Share).
  4. Status ditampilkan bersama konteks nama customer-facing (mis. "In progress — dilihat customer sebagai Being worked on"), beserta transitions, assignee, dan priority.
  5. Context panel berisi field sesuai layout request type (JSM-PRJ-027).
  6. Panel SLA, approvals, forms, KB, Assets, dan services tampil sesuai Features aktif.
  7. Activity memiliki tab All, Comments, History, Work log, dan Approvals.
  8. Indikator "Raised on behalf of" tampil bila request dibuat agent atas nama customer.
  9. ⚠ Fitur AI (ringkasan request, similar requests, sentiment) hanya tampil bila fitur AI opsional aktif dan plan mendukung (modul 35).
  10. Editor komentar memiliki dua mode, "Reply to customer" (publik) dan "Add internal note" (internal), dengan pembeda visual (JSM-PRJ-053).
  11. Panel SLA menampilkan status per metrik: on track, near breach, breached, paused, completed (JSM-PRJ-057).
  12. Info request type dan channel asal (portal, email, chat, widget, API, dialog Create internal `jira`, deployment) tampil di panel detail.
  13. Customer context panel menampilkan profil requester, organisasi, dan request lain dari customer (JSM-PRJ-056).
  14. Aksi "View customer request" membuka tampilan portal (sama dengan butir 2).
  15. Agent dapat mengelola Share with organisasi dan Request participants (JSM-PRJ-054).
  16. Panel Knowledge base memungkinkan mencari dan membagikan artikel serta membuat artikel dari request (JSM-PRJ-059).
  17. Panel similar requests, linked alerts, dan related incidents tampil bila modul terkait aktif (modul 32, 34, 35).
  18. Aksi Change request type tersedia dari issue view (JSM-PRJ-054).
  19. Panel Approvals menampilkan status, approver, dan decision (JSM-PRJ-058).
  20. Rating CSAT (bintang dan komentar) tampil setelah survei diisi customer (JSM-PRJ-069).
  21. Agent dapat membuat linked work item ke project software untuk eskalasi ke tim development, dan development panel (commit/branch/PR) dari work item tertaut ditampilkan (modul 24).
  22. ⚠ Indikator bahasa dan zona waktu customer tampil di samping nama reporter.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue view inti ISS-007/ISS-008 ([06](06-work-item-inti.md)); link ISS-029. Endpoint issue view agent mengembalikan data customer-facing dan internal dalam satu respons untuk agent saja.

### JSM-PRJ-053 Reply to customer vs Add internal note

- **Alias inventaris:** Internal Notes vs Reply to Customer (JSM)
- **Deskripsi:** Dua mode komentar di issue view untuk memisahkan komunikasi publik ke customer dan catatan internal agent.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (keduanya), collaborator (internal note saja), Customer/Pelapor (menerima balasan).
- **Aturan perilaku (acceptance criteria):**
  1. "Reply to customer" menghasilkan komentar publik yang dikirim ke reporter dan participants sebagai customer notification.
  2. "Add internal note" menghasilkan komentar internal yang hanya terlihat agent/collaborator dan tidak dikirim ke customer.
  3. Komentar internal memiliki pembeda visual (label "Internal" dan latar berwarna) di activity.
  4. Collaborator hanya dapat menulis internal note; tombol Reply to customer tidak tampil baginya dan API menolak dengan 403.
  5. Lampiran yang ditambahkan pada balasan customer terlihat customer di portal; lampiran pada internal note tidak.
  6. Mention pengguna internal (@) tersedia di kedua mode; mention di internal note tidak pernah mengirim notifikasi ke customer.
  7. Edit/hapus komentar mengikuti permission komentar inti (PERM-030); mengubah komentar publik memicu notifikasi "Public comment edited" bila aktif.
  8. ⚠ Balasan agent lewat email ke alamat project menjadi komentar publik.
  9. Canned responses dirinci di modul 31 dan dapat disisipkan di kedua mode.
  10. Editor menampilkan dua aksi terpisah, "Add internal note" dan "Reply to customer", dengan mode aktif terakhir diingat per pengguna.
  11. Internal note berlatar kuning (token warna "internal") dan tidak terlihat customer di portal.
  12. Reply to customer dikirim lewat email dan tampil di portal.
  13. Saat agent mengirim komentar publik, editor menampilkan peringatan "Komentar ini akan terlihat oleh customer" sebelum dikirim.
  14. Canned responses (template balasan) tersedia dari toolbar editor (modul 31).
  15. Entity komentar menyimpan flag boolean `is_public` yang menentukan visibilitas ke customer dan diekspos di REST API sebagai properti `jsdPublic` (hanya untuk komentar di service project).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komentar inti CMT-001 s.d. CMT-010 ([08](08-komentar-lampiran-aktivitas.md)); tampilan editor agent di JSM-QSK-012 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)). **Skema visibilitas komentar tunggal (berlaku untuk CMT-001, CMT-004, CMT-010, JSM-QSK-012, dan DM-007 [28](28-model-data-inti.md)):** tabel `comments` memiliki dua dimensi yang ortogonal: (a) `visibility_type` (`role`/`group`/NULL) + `visibility_value` untuk restriksi ke role/grup internal (CMT-004, PERM-045), berlaku di semua tipe project; (b) `is_public` BOOLEAN NOT NULL DEFAULT `false`, hanya bermakna di service project: `true` = Reply to customer (terlihat customer di portal, memicu customer notification), `false` = internal note. Aturan: di project non-service use case mengisi `is_public = true` dan nilainya diabaikan oleh semua jalur baca (tidak diekspos sebagai `jsdPublic`); komentar dengan `is_public = true` tidak boleh memiliki `visibility_type` non-NULL (kombinasi ditolak 400 karena customer tidak memiliki role/grup), sedangkan internal note boleh ditambah restriksi role/grup. Representasi turunan: `jsdPublic` di REST API = `is_public`; nilai `visibility` `public`/`internal` di JSM-QSK-012 dan event `comment.created(visibility)` adalah turunan dari `is_public`, bukan kolom terpisah; flag `is_internal` di DM-007 butir 1 harus diganti menjadi `is_public` (dengan `is_internal = NOT is_public`) agar default aman (komentar di service project tanpa penanda eksplisit diperlakukan internal). Modul 08, 28, dan 31 perlu menyelaraskan Catatan masing-masing ke skema ini.

### JSM-PRJ-054 Request details and request participants

- **Alias inventaris:** Request participants
- **Deskripsi:** Panel detail request untuk mengelola reporter, participants, organizations, dan request type, termasuk penanganan work item tanpa request type. Request participants adalah customer atau agent tambahan yang menerima notifikasi dan dapat melihat request.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Modify reporter, Edit work item), Admin Proyek, Customer/Pelapor (menambah participant dari portal).
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat mengubah Reporter (raise on behalf) bila memiliki permission Modify Reporter.
  2. Agent dapat menambah/menghapus Request participants; participant menerima notifikasi dan dapat melihat request di portal.
  3. Agent dapat menambah Organizations untuk berbagi request dengan seluruh anggota organisasi.
  4. Change request type memilih request type lain dengan work type yang sama; request type dengan work type berbeda memerlukan ubah work type/Move terlebih dahulu.
  5. Work item yang dibuat langsung (Create, import, API, automation) tanpa request type memiliki field Request type kosong sehingga tidak tampil di portal maupun daftar My requests customer.
  6. ⚠ Issue view menampilkan indikator "No request type" beserta aksi "Set request type".
  7. ⚠ Request type menjadi kosong setelah request type dihapus, work type diubah ke work type lain, atau work item dipindah/di-clone antar service project, sehingga perlu dipilih ulang.
  8. Tautan "View request in portal" tidak tersedia bila work item tidak punya request type.
  9. ⚠ Request type dapat di-set massal lewat bulk edit field Request Type atau action automation "Edit field: Request type".
  10. JQL `"Request Type" is EMPTY` dapat dipakai untuk audit work item tanpa request type.
  11. ⚠ Customer notifications tidak dikirim untuk work item tanpa request type.
  12. Panel menampilkan channel asal (portal, email, API, chat, dialog Create internal `jira`).
  13. Field Request participants bertipe multi-user ditampilkan di issue view agent.
  14. Customer dapat menambah participant dari portal sesuai pengaturan Customer sharing (JSM-PRJ-051).
  15. Penerima CC email ditambahkan sebagai participant bila opsi diaktifkan di pengaturan email (JSM-PRJ-063).
  16. Participant melihat request di filter portal "Where I'm a participant".
  17. Participant menerima customer notifications sesuai aturan notifikasi.
  18. Participant dapat dihapus oleh agent atau oleh reporter dari portal.
  19. Penambahan participant memicu customer notification "Participant added" ke participant baru.
  20. JQL `"Request participants" = user` mengembalikan request tempat pengguna menjadi participant.
  21. ⚠ Jumlah participant per request dibatasi konfigurasi server (default diusulkan 100); penambahan melebihi batas ditolak dengan pesan galat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bulk change ISS-035, move/clone ISS-031/ISS-032 ([06](06-work-item-inti.md)); default request type per work type di JSM-PRJ-020.

### JSM-PRJ-055 Raise a request on behalf of customer

- **Alias inventaris:** Raise request on behalf & agent-created requests
- **Deskripsi:** Kemampuan agent membuat request untuk customer melalui portal atau create dialog, termasuk penetapan request type pada work item yang dibuat tanpa request type.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Create work item + Modify Reporter).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Raise a request" di sidebar membuka portal sebagai agent.
  2. Portal menampilkan field "Raise this request on behalf of" hanya bagi agent.
  3. Create dialog global (ISS-002) di service project menampilkan field Request type.
  4. ⚠ Agent dapat memilih customer yang ada atau menambah customer baru dengan mengetik alamat email, bila Customer permissions mengizinkan.
  5. Notifikasi "Request created" dikirim ke customer yang diwakili, bukan ke agent.
  6. Request tampil di daftar request portal milik customer tersebut.
  7. Create dialog menyediakan pilihan request type dan field "Raise this request on behalf of".
  8. Opsi "Raise this request on behalf of" tersedia di portal (sama dengan butir 2) dan juga lewat API dengan parameter `raiseOnBehalfOf`.
  9. Reporter diisi customer, dan channel dicatat `jira` (create dialog) atau `portal` (portal agent).
  10. Customer baru hanya dibuat saat raise on behalf bila "Who can raise requests" bukan "hanya customer yang ditambahkan" atau agent memiliki permission mengelola customer.
  11. Work item tanpa request type tidak tampil di portal.
  12. Agent dapat menetapkan request type pada work item yang dibuat tanpa request type (JSM-PRJ-054).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Reporter rules ISS-017 dan PERM-024 (Modify Reporter).

### JSM-PRJ-056 Customer context panel

- **Deskripsi:** Panel di issue view yang menampilkan informasi customer dan organization pemohon. Fitur ini adalah pemilik kerangka panel dasar (Fase 5); tampilan panel di agent issue view dirinci di JSM-QSK-015 dan perluasan profil pelanggan Customer Service Management (Fase 7) di JSM-AI-063.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan detail reporter: nama, email, avatar, dan timezone (email mengikuti setting visibilitas email, NTF-044).
  2. Panel menampilkan organization customer.
  3. ⚠ Customer & organization details (custom attributes) dapat diedit agent bila plan mendukung (modul 30).
  4. Panel menampilkan daftar request terbaru customer (maks. 5, dengan tautan "Lihat semua").
  5. Panel menautkan ke profil customer.
  6. ⚠ Panel menampilkan Assets terkait customer bila modul 33 aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik perilaku panel dasar: fitur ini. JSM-QSK-015 Agent Issue View - Customer Context Panel ([31](31-jsm-antrian-sla-knowledge-base-laporan.md), P1 / Fase 5) adalah rujukan tampilan yang memakai kerangka ini dan tidak menambah aturan baru. JSM-AI-063 Customer Service Management - Customer Profile & Context Panel ([35](35-jsm-ai-otomasi-api-admin-csm.md), P2 / Fase 7) adalah perluasan khusus CSM (profil pelanggan lengkap, riwayat lintas kanal) yang dibangun di atas panel ini; perbedaan prioritas/fase disengaja: panel dasar P1 / Fase 5, perluasan CSM P2 / Fase 7. Data customer/organization dan customer details fields (butir 3) dimiliki JSM-PRT-033 s.d. JSM-PRT-035 ([30](30-jsm-help-center-portal-pelanggan-kanal.md)).

### JSM-PRJ-057 SLA panel in issue view

- **Deskripsi:** Panel yang menampilkan status dan sisa waktu SLA request untuk agent.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Browse project).
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan daftar SLA aktif (Time to first response, Time to resolution, dll.).
  2. Setiap SLA menampilkan status running, paused, breached, atau met.
  3. Setiap SLA menampilkan sisa waktu atau waktu terlampaui dalam format jam/menit kerja.
  4. Setiap SLA menampilkan goal/target waktu.
  5. Warna indikator berubah saat mendekati breach (ambang default 30 menit) dan saat breached.
  6. Riwayat siklus SLA yang sudah selesai dapat dibuka dari panel.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi dan perhitungan SLA di modul 31.

### JSM-PRJ-058 Approvals panel in issue view

- **Deskripsi:** Panel approval yang menampilkan approver dan memungkinkan keputusan approve/decline.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, approver (agent atau customer).
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan daftar approver dan status keputusan masing-masing (pending, approved, declined).
  2. Tombol Approve/Decline tampil bagi approver; customer approver juga dapat memutuskan dari portal atau email.
  3. Panel menampilkan syarat approval (semua approver atau jumlah tertentu; approver dari group field).
  4. ⚠ Agent dapat mengubah field Approvers sebelum keputusan pertama diambil.
  5. Riwayat keputusan tercatat di tab Approvals activity.
  6. Keputusan memicu transisi workflow yang dikonfigurasi (approve/decline).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine approval di modul 32.

### JSM-PRJ-059 Linked knowledge base articles in issue view

- **Deskripsi:** Integrasi knowledge base di issue view untuk menautkan dan membagikan artikel.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Panel menyarankan artikel KB berdasarkan summary dan description request.
  2. Agent dapat mencari artikel KB dari work item.
  3. Agent dapat membagikan artikel ke customer lewat balasan (menyisipkan tautan ke Reply to customer).
  4. Agent dapat membuat artikel KB dari work item dengan isi awal dari request.
  5. Panel menampilkan daftar artikel yang ditautkan ke request.
  6. ⚠ Artikel yang dibagikan lewat balasan terlihat customer di portal; artikel yang hanya ditautkan tidak ditampilkan ke customer.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Knowledge base di modul 31.

## Field sistem & JQL request

### JSM-PRJ-060 Customer request system fields

- **Alias inventaris:** Jira Service Management fields
- **Deskripsi:** Field sistem khas Service Management untuk request yang muncul di konfigurasi field ketika Service Management aktif dan dipakai di JQL, queue, SLA, dan laporan.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (mengisi), Admin Proyek dan Admin Situs (konfigurasi field).
- **Aturan perilaku (acceptance criteria):**
  1. Field Request participants bertipe multi-user.
  2. Field Organizations bertipe multi-organization.
  3. Field Request type (Customer Request Type) menyimpan ID request type project.
  4. Field Satisfaction (1-5) dan Satisfaction date diisi saat customer memberi rating.
  5. Field SLA (Time to first response, Time to resolution, dll.) dibuat per SLA dan hanya diubah oleh engine SLA.
  6. Field Approvals dan Approvers/Approver groups tersedia untuk langkah approval.
  7. Field ITSM tersedia: Urgency, Impact, Affected services, Pending reason, Change type, Change risk, Change reason, Change start date, Change end date.
  8. ⚠ Field Request language menyimpan bahasa request (JSM-PRJ-066).
  9. ⚠ Field Sentiment (AI) hanya dibuat bila fitur AI opsional aktif (modul 35).
  10. Field Request type, Request participants, Organizations, Satisfaction, Approvals, dan field SLA dikunci (locked): tidak dapat dihapus atau diganti tipe, hanya konteks/screen yang dapat diatur.
  11. Request Type (Customer Request Type), Request participants, dan Organizations muncul di daftar field administrasi saat Service Management aktif.
  12. Satisfaction dan Satisfaction date muncul di daftar field administrasi.
  13. Time to first response dan Time to resolution muncul sebagai field SLA berstatus locked.
  14. Approvals dan Approvers muncul di daftar field administrasi.
  15. Impact, Urgency, Change type, Change risk, Change reason, dan Pending reason muncul di daftar field administrasi.
  16. Affected services dan Assets objects field muncul bila modul 32/33 aktif.
  17. ⚠ Request language dan Sentiment muncul di daftar field administrasi sesuai butir 8 dan 9.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Administrasi custom field CFG-034 ([05](05-konfigurasi-work-item-skema.md)). Field ini disarankan sebagai field sistem bertipe khusus (bukan custom field biasa) agar indeks JQL efisien.

### JSM-PRJ-061 JSM JQL extensions for requests

- **Deskripsi:** Field dan fungsi JQL khusus Service Management untuk mencari request di queue, filter, dan automation.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek; hasil tetap disaring permission.
- **Aturan perilaku (acceptance criteria):**
  1. JQL mendukung field `"Request type"`, `"Request participants"`, `Organizations`, dan `Satisfaction` dengan operator yang sesuai tipe (`=`, `!=`, `in`, `is EMPTY`, `>`/`<` untuk Satisfaction).
  2. `"Request Type" is EMPTY` / `is not EMPTY` dapat dipakai untuk audit work item tanpa request type.
  3. ⚠ Field `request-channel-type` mendukung nilai `email`, `portal`, `api`, `jira`, `anonymousportal`, `widget`, `chat`.
  4. ⚠ Fungsi approval tersedia: `approver()`, `myApproval()`, `myPending()`, `pendingBy()`.
  5. Fungsi SLA (`breached()`, `remaining()`, `elapsed()`, dll.) dirinci di modul 31.
  6. Field dan fungsi ini menjadi dasar JQL queue (modul 31) dan autocomplete JQL editor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine JQL di [12](12-jql-bahasa-kueri.md) (JQL-033, JQL-037). Pasangan tumpang tindih: JSM-AI-013 JSM JQL Fields ([35](35-jsm-ai-otomasi-api-admin-csm.md)), yang memuat alias inventaris "JSM JQL extensions". Pemilik: JSM-AI-013 memiliki daftar lengkap field/fungsi JQL Service Management (termasuk SLA, Assets, dan field ITSM); fitur ini hanya mendefinisikan subset field request (request type, participant, organization, satisfaction, kanal, approval) yang wajib tersedia saat service project pertama dirilis. Prioritas diselaraskan ke P0 karena queue (modul 31) bergantung padanya.

## Kanal & komunikasi customer di tingkat project

### JSM-PRJ-062 Project settings - Portal settings

- **Deskripsi:** Pengaturan tampilan portal milik service project di dalam help center.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Admin mengatur nama portal (wajib, maks. 255 karakter) dan introduction text.
  2. Admin mengatur logo/avatar portal (gambar, maks. 1 MB).
  3. Admin mengatur portal announcement (banner pesan) yang tampil di atas portal; announcement kosong tidak menampilkan banner.
  4. Halaman menautkan ke pengaturan help center (branding, layout) yang dikelola global (modul 30).
  5. ⚠ Bila multiple help centers aktif (plan lanjut), admin dapat menghubungkan portal ke help center tertentu.
  6. Tombol pratinjau/buka portal membuka portal di tab baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Help center & portal di modul 30.

### JSM-PRJ-063 Project settings - Email requests

- **Deskripsi:** Halaman Project settings > Email requests di service project untuk mengelola alamat email yang mengubah email masuk menjadi request. Pemilik perilaku email channel adalah modul 30: koneksi akun email JSM-PRT-040, pemrosesan dan parsing JSM-PRT-041, loop prevention JSM-PRT-042, dan email logs JSM-PRT-043; fitur ini hanya mendefinisikan halaman pengaturannya.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap project memiliki alamat email cloud bawaan (mis. `<key>@<site-mail-domain>`) yang aktif saat project dibuat.
  2. Admin dapat menambah email custom: Google (Gmail via OAuth), Microsoft (Exchange Online/365 via OAuth), atau server IMAP/POP lain dengan kredensial terenkripsi.
  3. Admin memilih request type default per alamat email.
  4. Penyimpanan memvalidasi field required pada request type email (JSM-PRJ-030 butir 2).
  5. Email log menampilkan riwayat pemrosesan per pesan (diterima, dibuat request, ditambahkan sebagai komentar, ditolak) beserta alasan galat.
  6. Setiap alamat email dapat diaktifkan/dinonaktifkan.
  7. Pengirim tak dikenal diperlakukan sesuai Customer permissions: dibuat sebagai customer atau ditolak (JSM-PRJ-051).
  8. ⚠ Opsi "Add CC recipients as request participants" menambahkan penerima CC sebagai participant.
  9. ⚠ Jumlah alamat email per project dibatasi konfigurasi server (default diusulkan 10).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik perilaku: JSM-PRT-040 (butir 1-2, 6, 9), JSM-PRT-041 Email request processing & parsing (butir 3-4, 7-8: pembuatan request, threading, CC, pengirim tak dikenal), dan JSM-PRT-043 (butir 5) di [30](30-jsm-help-center-portal-pelanggan-kanal.md); bila terjadi perbedaan, aturan modul 30 yang berlaku. Engine email masuk yang sama dipakai project business lewat BIZ-047 Project email requests ([26](26-kerja-bisnis-formulir-persetujuan.md), Fase 3), sehingga Fase 5 menambah lapisan khusus service project (request type default, customer, participant) di atas engine tersebut, bukan membangun engine baru. Infrastruktur email NTF-023/NTF-033 ([19](19-notifikasi-email.md)).

### JSM-PRJ-064 Project settings - Widget

- **Deskripsi:** Pengaturan widget yang dapat disematkan di situs web untuk membuat request dan mencari knowledge base.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle mengaktifkan/menonaktifkan widget; widget nonaktif tidak merender apa pun di situs host.
  2. Admin memilih request type yang tampil di widget.
  3. Admin dapat mengaktifkan pencarian knowledge base di widget.
  4. Halaman menyediakan kode embed (script) untuk dipasang di situs, dengan tombol salin.
  5. ⚠ Hanya field yang didukung widget (teks, select, date, attachment) yang dapat dipakai; request type dengan field wajib tak didukung tidak dapat dipilih.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail widget di modul 30.

### JSM-PRJ-065 Project settings - Chat

- **Deskripsi:** Pengaturan kanal chat (Slack/Microsoft Teams) agar customer dapat membuat dan menindaklanjuti request dari chat.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek; Agent Layanan memakai triage channel.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat menghubungkan workspace Slack atau Microsoft Teams lewat OAuth aplikasi produk ini.
  2. Admin memilih request channel (tempat customer meminta bantuan) dan triage channel untuk agent.
  3. Admin memilih request type untuk request dari chat.
  4. ⚠ Request dapat dibuat dari pesan chat (mis. reaksi emoji tertentu pada pesan).
  5. Komentar di thread chat tersinkron dengan komentar publik work item dan sebaliknya.
  6. Integrasi virtual service agent di chat tersedia pada plan lanjut (modul 35).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail kanal chat di modul 30; integrasi chat umum di [25](25-integrasi-impor-ekspor-ekstensi.md).

### JSM-PRJ-066 Project settings - Language support

- **Alias inventaris:** Language support (portal & customer-facing translations)
- **Deskripsi:** Mengaktifkan banyak bahasa pada service project dan menerjemahkan teks customer-facing agar portal dan notifikasi tampil sesuai bahasa customer.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek menerjemahkan; Customer/Pelapor melihat hasilnya.
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih default language project (default: bahasa default site).
  2. Admin dapat menambah additional languages dari daftar bahasa yang didukung.
  3. Halaman menampilkan status translasi per bahasa (complete/incomplete) beserta jumlah teks yang belum diterjemahkan.
  4. Nama dan deskripsi request type dapat diterjemahkan.
  5. Display name dan help text field dapat diterjemahkan.
  6. Nama portal group serta nama dan intro portal dapat diterjemahkan.
  7. Nama status customer-facing dapat diterjemahkan.
  8. Template customer notification dan teks satisfaction survey dapat diterjemahkan.
  9. ⚠ Konten form (label, deskripsi, opsi pertanyaan) dapat diterjemahkan (JSM-PRJ-034).
  10. Bahasa customer ditentukan dari profil, lalu browser; konten tanpa translasi jatuh ke default language.
  11. ⚠ Bahasa request tercatat di field Request language saat request dibuat (JSM-PRJ-060).
  12. ⚠ Translasi dapat diekspor dan diimpor (CSV/JSON per bahasa) untuk diterjemahkan di luar produk.
  13. Pemilihan default language dan bahasa tambahan dilakukan di satu halaman yang sama.
  14. Teks yang dapat diterjemahkan mencakup nama portal, intro, request type, nama/help text field, status customer-facing, dan portal group.
  15. Customer notification, portal announcement, dan teks CSAT dapat diterjemahkan.
  16. Customer melihat konten sesuai bahasa profil atau browser.
  17. Fallback ke default language berlaku untuk setiap teks yang terjemahannya kosong, per teks (bukan per halaman).
  18. Daftar bahasa yang didukung mengikuti daftar locale UI produk ini (minimal EN dan ID sesuai `internal/shared/constants/error.go`) dan dapat diperluas tanpa perubahan skema.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Terjemahan status workflow WF-040 ([09](09-workflow-status-transisi.md)). Usulan tabel `translations` (entity_type, entity_id, locale, field, value).

### JSM-PRJ-067 Project settings - Customer notifications

- **Deskripsi:** Halaman Project settings > Customer notifications yang mengelola template dan aturan notifikasi email ke customer untuk event request, terpisah dari notification scheme internal. Pemilik perilaku (event, penerima, pengiriman, variabel template, look and feel) adalah JSM-PRT-036 s.d. JSM-PRT-039 di [30](30-jsm-help-center-portal-pelanggan-kanal.md).
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor menerima.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Notifikasi default mengikuti daftar event JSM-PRT-036: Request created, Public comment added, Customer-visible status changed, Request resolved, Request reopened, Participant added, Organization added, Approval required; "Public comment edited" tersedia tetapi default nonaktif.
  2. Setiap notifikasi dapat diaktifkan/dinonaktifkan.
  3. Subject dan body dapat diedit dengan variabel/placeholder (summary, link request, komentar, status customer-facing, nama customer); placeholder tidak dikenal ditolak saat simpan (katalog variabel dimiliki JSM-PRT-037).
  4. Template tersedia dalam format HTML dan plain text.
  5. Admin dapat mempratinjau template dengan data contoh dan me-reset ke default.
  6. Template dapat diterjemahkan per bahasa (JSM-PRJ-066).
  7. ⚠ Custom notification rules (trigger, kondisi, recipients) dapat ditambahkan selain notifikasi default.
  8. Look and feel global email (logo, warna, footer) diatur di level admin (JSM-PRT-038, ditautkan dari JSM-PRJ-019), bukan per project.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik perilaku: JSM-PRT-036 Customer notifications ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); variabel JSM-PRT-037, look and feel JSM-PRT-038, sender/reply-to JSM-PRT-039. Butir 1-8 adalah sisi halaman pengaturan dari aturan tersebut dan bila terjadi perbedaan, modul 30 yang berlaku. Template email NTF-042 ([19](19-notifikasi-email.md)). Notifikasi tidak dikirim ke work item tanpa request type (JSM-PRJ-054).

### JSM-PRJ-068 Project settings - Notifications (internal notifications)

- **Deskripsi:** Pengaturan notifikasi internal untuk agent/collaborator di service project memakai notification scheme inti, terpisah dari Customer notifications, dengan perilaku khusus untuk internal note dan reply ke customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs (notification scheme), Agent Layanan dan collaborator (penerima).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Project settings > Notifications menampilkan notification scheme yang dipakai dan memungkinkan menggantinya.
  2. Event (Work item created, Commented, Assigned, Resolved, dll.) dipetakan ke penerima internal Assignee, Reporter, Watchers, Project role, Group, Single user memakai editor notification scheme inti.
  3. Customer (reporter/participant tanpa lisensi) tidak pernah menerima email dari notification scheme; customer hanya menerima Customer notifications.
  4. "Add internal note" hanya memicu notifikasi internal (watchers/agent) dan tidak pernah dikirim ke customer.
  5. "Reply to customer" memicu notifikasi internal ke penerima scheme dan customer notification ke reporter/participants.
  6. ⚠ Pengguna berlisensi yang juga reporter/participant menerima satu email per event (notifikasi customer diutamakan untuk komentar publik, notifikasi internal untuk event lain).
  7. ⚠ Notifikasi Approval required dikirim sebagai notifikasi internal untuk approver berlisensi dan sebagai customer notification untuk approver customer.
  8. Notifikasi internal menghormati preferensi notifikasi pribadi pengguna dan batched email notifications (NTF-019, NTF-020).
- **Team-managed vs company-managed:** Clone hanya company-managed; varian toggle per event team-managed hanya rujukan (JSM-PRJ-012).
- **Catatan & ketergantungan:** Notification scheme NTF-004/NTF-005, delivery rules NTF-008 ([19](19-notifikasi-email.md)).

### JSM-PRJ-069 Project settings - Satisfaction

- **Deskripsi:** Pengaturan survei kepuasan customer (CSAT) setelah request resolved.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengonfigurasi; Customer/Pelapor (reporter) mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle "Collect satisfaction feedback" mengaktifkan survei untuk project.
  2. Survei memakai rating 5 bintang dengan komentar opsional; satu rating per request (dapat diubah sampai request dibuka kembali).
  3. ⚠ Admin dapat mengubah pertanyaan dan teks survei.
  4. Survei dikirim via email dan tersedia di portal saat request resolved.
  5. Rating disimpan di field Satisfaction dan Satisfaction date (JSM-PRJ-060).
  6. Hasil tampil di Satisfaction report (modul 31).
  7. Teks survei dapat diterjemahkan (JSM-PRJ-066).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** CSAT dan laporannya di modul 31.

## Plan & batas service project

### JSM-PRJ-070 JSM plans and limits for service projects

- **Alias inventaris:** JSM licensing & tiers
- **Deskripsi:** Batasan dan perbedaan kemampuan service project per plan Free/Standard/Premium/Enterprise serta model lisensi per agent. Di clone, plan diperlakukan sebagai feature flag dan batas konfigurasi (VISI-003), tanpa billing.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengatur plan/feature flag).
- **Aturan perilaku (acceptance criteria):**
  1. Plan Free membatasi maksimal 3 agent dengan fitur dasar (request types, portal, SLA, KB).
  2. Plan Standard mengizinkan jumlah agent sesuai konfigurasi dan fitur dasar lengkap.
  3. Plan Premium membuka kapabilitas lanjutan (virtual service agent, sandbox, multiple help centers, kapasitas automation lebih besar).
  4. Plan Enterprise membuka kapabilitas Premium plus data/analytics lanjutan dan compliance.
  5. Batas penyimpanan dan eksekusi automation per plan dikonfigurasi di server; saat batas tercapai admin menerima peringatan.
  6. Feature gating: fitur yang tidak termasuk plan disembunyikan atau ditampilkan nonaktif dengan keterangan plan yang dibutuhkan (NAV-024), dan API mengembalikan 403 dengan kode galat `plan_required`.
  7. Model lisensi dihitung per agent, sedangkan customer tidak terbatas dan gratis.
  8. Plan Free mencakup hingga 3 agent dengan fitur inti (sama dengan butir 1).
  9. Plan Standard mencakup SLA, approval, KB, report, dan fitur sejenis.
  10. ⚠ Plan Premium mencakup Assets, virtual service agent, change risk engine, incident lanjutan, heartbeat, dan multiple help centers.
  11. Plan Enterprise mencakup analytics lanjutan, data lake/ekspor data, serta batas dan keamanan tambahan.
  12. Paket bundel lintas produk (mis. bundel Service Management, CSM, Assets, dan asisten AI) tidak dibangun; setiap kapabilitas diaktifkan lewat feature flag terpisah.
  13. ⚠ Batas jumlah objek Assets per plan dikonfigurasi di server (modul 33).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tier global di VISI-003 ([00](00-visi-ruang-lingkup-glosarium-baseline.md)) dan PERM-005 ([03](03-izin-peran-keamanan.md)). Billing tidak dibangun. Pasangan tumpang tindih: JSM-AI-005 Plan Feature Matrix & Service Collection ([35](35-jsm-ai-otomasi-api-admin-csm.md)). Pemilik: JSM-AI-005 memiliki matriks kapabilitas per plan dan mekanisme feature flag tingkat situs; fitur ini hanya mendefinisikan batas yang berlaku pada service project (jumlah agent, gating fitur project, kode galat `plan_required`) dan mengikuti matriks JSM-AI-005 bila terjadi selisih. Prioritas diselaraskan ke P1.

## Pertanyaan terbuka

1. Apakah clone sepenuhnya menghapus varian team-managed service project (JSM-PRJ-012), termasuk untuk impor, atau tetap mendukung mode baca-saja untuk project hasil migrasi?
2. Apakah customer disimpan di tabel `users` yang sama dengan kolom tipe akun, atau di tabel `customer_accounts` terpisah dengan namespace email sendiri (JSM-PRJ-045, JSM-PRJ-048)?
3. Berapa batas default jumlah request type, portal group, participant, alamat email, dan form per work item yang akan dipakai (JSM-PRJ-020, JSM-PRJ-024, JSM-PRJ-054, JSM-PRJ-063, JSM-PRJ-034)?
4. Apakah aksi Duplicate request type (JSM-PRJ-020 butir 12) dan Copy form ke project lain (JSM-PRJ-034 butir 5) dibangun di rilis awal?
5. Apakah transisi yang boleh dipicu customer dari portal (Resolve/Cancel) dimodelkan sebagai properti transisi "customer can trigger" (JSM-PRJ-031, JSM-PRJ-032)?
6. Apakah data connections Forms berlingkup project atau global, dan metode autentikasi apa yang didukung (JSM-PRJ-039)?
7. Apakah collaborator boleh mengedit form internal dan field request, atau hanya menambah internal note (JSM-PRJ-042, JSM-PRJ-047)?
8. Bagaimana de-duplikasi notifikasi untuk pengguna berlisensi yang juga reporter/participant (JSM-PRJ-068 butir 6)?
9. Apakah template katalog (Design, Sales, demo space) dikirim di rilis awal Fase 5 atau ditambahkan bertahap (JSM-PRJ-010, JSM-PRJ-011)?
10. Apakah fitur AI (pembuatan request type/form berbantuan AI, ringkasan, sentiment) termasuk ke dalam cakupan atau hanya disiapkan sebagai titik integrasi opsional (JSM-PRJ-021, JSM-PRJ-035, JSM-PRJ-052)?
11. Domain email apa yang dipakai untuk alamat email cloud bawaan project, dan apakah instalasi self-hosted wajib mengonfigurasi inbound mail sebelum service project dapat dibuat (JSM-PRJ-063)?
12. Apakah BIZ-033 Form conditional logic ([26](26-kerja-bisnis-formulir-persetujuan.md)) dinaikkan ke P1 agar engine kondisi bersama dikirim di Fase 4 sebelum JSM-PRJ-037, atau engine kondisi dikirim pertama bersama JSM-PRJ-037 di Fase 5 dan form bisnis menyusul?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga**: fitur ProForma Marketplace lama di luar Forms native, app form lanjutan (pembayaran, e-signature), app CMDB pihak ketiga, dan app tambahan untuk portal tidak dibangun sebagai fitur native; kebutuhan dipenuhi lewat REST API, webhook, dan ekstensi (modul 23 dan 25).
- **Fitur dc-only-or-deprecated**: Team-managed service project (legacy) (JSM-PRJ-012) hanya dicatat sebagai rujukan; clone hanya membangun service project company-managed, dan importer memetakan project team-managed lama ke company-managed.
- **Layanan Atlassian lintas produk**: administrasi organisasi Atlassian (admin.atlassian.com), billing dan pricing per agent, bundel Service Collection, Rovo, Atlassian Intelligence, Atlassian Analytics, dan Data Lake tidak dibangun; produk ini memakai administrasi pengguna sendiri (modul 02/22), plan sebagai feature flag (JSM-PRJ-070), dan fitur AI opsional (modul 35).
- **Confluence sebagai sumber knowledge base** tidak dibangun; knowledge base native dirinci di modul 31, dan panel KB di issue view (JSM-PRJ-059) hanya memakai KB tersebut.
- **Aplikasi chat pihak ketiga** (Slack, Microsoft Teams) bukan bagian produk; JSM-PRJ-065 hanya mengatur integrasi lewat API resmi kedua platform.
- **Detail modul Service Management lain** (help center, queue, SLA, KB, ITSM, Assets, Operations, virtual service agent, Customer Service Management) tidak dirinci di modul ini dan dirujuk ke modul 30-35.
