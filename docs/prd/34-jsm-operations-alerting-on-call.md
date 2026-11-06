# 34 — Service Management: Operations (Alerts, On-call, Escalation)

Modul ini mendefinisikan kapabilitas **Operations** di Service Management produk ini: penerimaan dan pengelolaan **alert** dari tool monitoring, API, email, dan pembuatan manual; deduplikasi, siklus hidup, aksi, dan pencarian alert; **alert policy** dan **notification policy**; **maintenance window** dan **heartbeat**; **routing rule** dan **escalation policy** per team; **on-call schedule** beserta rotation, override, zona waktu, ekspor kalender, dan forwarding; **notification rule personal**, contact method, dan aksi langsung dari notifikasi; kerangka **integrasi** incoming/outgoing (API, email, monitoring, webhook, edge connector, ChatOps, sinkron tiket); keterkaitan alert dengan **incident** dan services; pengelompokan alert berbasis AI; laporan operasional; REST API konfigurasi; serta model team, izin, pengaturan global, dan batasan plan. Kapabilitas ini adalah padanan fungsi alerting/on-call yang di ekosistem rujukan dulu berupa produk terpisah dan kini menyatu ke Service Management. Alert adalah entitas tersendiri (bukan work item) dengan state, riwayat, dan API sendiri; incident, problem, dan change tetap dimiliki modul ITSM.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas team, service project, incident, maupun infrastruktur antrean job, pengiriman SMS/voice, atau push mobile. Seluruh fitur modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/operations` dengan sub-paket `alert`, `policy`, `oncall` (schedule, rotation, override, forwarding), `notifyrule` (notification rule & contact method), `integration`, dan `opsreport`; entitas bersama (Alert, AlertActivity, OpsTeam, Schedule, Rotation, Override, Escalation, RoutingRule, Integration, Heartbeat, Maintenance, ContactMethod, NotificationRule) di `internal/shared/domain/entity`; worker latar belakang untuk eskalasi, timer heartbeat, dan pengiriman notifikasi; serta UI di `frontend/src/routes` (SPA SvelteKit) di bawah navigasi Operations. Seluruh route mengikuti konvensi repo `/api/v1/ops/...`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — tier plan, istilah, dan batas cakupan produk.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — login, API token pengguna (AUTH-012) untuk REST API konfigurasi.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — profil, zona waktu pengguna, direktori People & Teams (USR-027).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — global permission, peran, dan audit perubahan izin.
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — entitas Atlassian Teams padanan produk ini (PLAN-057, PLAN-060).
- [19 Notifikasi & Email](19-notifikasi-email.md) — infrastruktur email keluar, push mobile (NTF-038), notifikasi chat personal (NTF-039, NTF-040).
- [20 Otomasi](20-otomasi.md) — trigger/action otomasi yang membuat atau bereaksi terhadap alert.
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — navigasi global dan paritas aplikasi mobile (NAV-050).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log (ADM-014), administrasi webhook (ADM-030).
- [23 REST API & Webhook](23-api-publik-webhook.md) — konvensi REST API publik, rate limit, pengiriman webhook (API-066).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — aplikasi Slack (INT-030 s.d. INT-035) dan Microsoft Teams (INT-036).
- [16 Laporan & Insight](16-laporan-insight.md) — Analytics dan Data Lake (RPT-038 s.d. RPT-041).
- [27 NFR, Arsitektur & Platform](27-nfr-arsitektur-platform.md) — antrean job, ketersediaan, retensi data.
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) — service project tempat team ops dihubungkan.
- [32 ITSM: Incident, Problem, Change & Services](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — incident, responders, PIR, service registry.
- [33 Assets (CMDB)](33-jsm-assets-cmdb.md) — dependensi service dan objek konfigurasi.
- [30 Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — kanal email masuk Service Management yang berbagi infrastruktur dengan Email Integration (JSM-OPS-033).
- [31 Queues, SLA, Knowledge Base, CSAT & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) — laporan Service Management umum di samping laporan Operations (JSM-OPS-044).
- [35 AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md) — kapabilitas AI Service Management dan administrasi produk.

## Persona terkait

- **Manajer Layanan/ITSM** — Ops admin tingkat situs: mengelola pengaturan Operations global, global integration/policy/maintenance, laporan lintas team, dan model izin.
- **Admin Situs** — mengaktifkan produk Service Management dan Operations, mengelola lisensi agent, audit log, dan negara SMS/voice.
- **Agent Layanan** — responder on-call: menerima notifikasi, acknowledge/close alert, membuat incident, melihat siapa yang on-call, membuat override dan forwarding rule.
- **Admin Proyek/Project Lead** — Team admin operasi: mengelola schedule, rotation, escalation, routing rule, integrasi, policy, heartbeat, dan maintenance milik team.
- **Scrum Master/Team Lead** — memimpin team engineering yang memegang on-call dan memantau Team Dashboard serta laporan beban on-call.
- **Developer/Anggota Tim** — engineer on-call yang menerima alert dari sistem miliknya dan memakai aplikasi mobile.
- **Integrator/Developer API** — mengirim alert dari sistem eksternal via Alert API, mengonfigurasi webhook, edge connector, dan mengelola konfigurasi sebagai kode.
- **Stakeholder/Viewer** — melihat status alert dan on-call secara read-only tanpa bertindak.

## User story modul

1. Sebagai Integrator/Developer API, saya ingin mengirim alert dari tool monitoring dengan API key integrasi agar sinyal masalah produksi langsung masuk ke team yang tepat (JSM-OPS-007, JSM-OPS-032, JSM-OPS-034).
2. Sebagai Agent Layanan, saya ingin alert berulang dengan alias yang sama digabung menjadi satu alert dengan penghitung agar saya tidak dibanjiri notifikasi duplikat (JSM-OPS-008).
3. Sebagai Admin Proyek/Project Lead, saya ingin menyusun routing rule dan escalation policy bertahap agar alert yang tidak di-acknowledge otomatis naik ke orang berikutnya (JSM-OPS-017, JSM-OPS-018).
4. Sebagai Admin Proyek/Project Lead, saya ingin membuat schedule on-call dengan beberapa rotation dan restriction jam agar beban jaga tim terbagi adil lintas zona waktu (JSM-OPS-019, JSM-OPS-020, JSM-OPS-023).
5. Sebagai Agent Layanan, saya ingin membuat override atau forwarding rule saat cuti agar notifikasi saya diteruskan ke rekan tanpa mengubah rotation (JSM-OPS-021, JSM-OPS-025).
6. Sebagai Developer/Anggota Tim, saya ingin mengatur notification rule personal per contact method dan jam tenang agar saya hanya dibangunkan untuk alert P1/P2 di malam hari (JSM-OPS-026, JSM-OPS-027).
7. Sebagai Agent Layanan, saya ingin acknowledge atau close alert langsung dari push notification, email, atau Slack agar respons tetap cepat tanpa membuka aplikasi web (JSM-OPS-028, JSM-OPS-029, JSM-OPS-037).
8. Sebagai Manajer Layanan/ITSM, saya ingin alert policy dan notification policy global agar alert dinormalisasi dan noise ditekan secara konsisten di semua team (JSM-OPS-013, JSM-OPS-014).
9. Sebagai Admin Proyek/Project Lead, saya ingin menjadwalkan maintenance window agar deploy terencana tidak memicu alert dan panggilan ke on-call (JSM-OPS-015).
10. Sebagai Integrator/Developer API, saya ingin heartbeat untuk job terjadwal agar kegagalan senyap (job tidak jalan) tetap menghasilkan alert (JSM-OPS-016).
11. Sebagai Agent Layanan, saya ingin membuat incident dari alert dan menautkan alert terkait agar penanganan gangguan besar terpusat di satu incident (JSM-OPS-039, JSM-OPS-040).
12. Sebagai Manajer Layanan/ITSM, saya ingin laporan MTTA/MTTR dan jam on-call di luar jam kerja agar saya bisa menilai kesehatan proses on-call dan kelelahan tim (JSM-OPS-044).
13. Sebagai Scrum Master/Team Lead, saya ingin Team Dashboard yang menampilkan on-call saat ini, alert terbuka, dan incident aktif agar kondisi operasional tim terlihat sekilas (JSM-OPS-043).
14. Sebagai Integrator/Developer API, saya ingin REST API konfigurasi yang idempoten agar schedule, escalation, dan integrasi bisa dikelola sebagai infrastructure-as-code (JSM-OPS-045, JSM-OPS-046).
15. Sebagai Admin Situs, saya ingin izin Operations yang jelas antara ops admin, team admin, dan member agar konfigurasi kritis tidak diubah sembarang orang (JSM-OPS-001, JSM-OPS-002).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 23 |
| P1 | 16 |
| P2 | 4 |
| P3 | 4 |
| **Total** | **47** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 0 |
| Fase 4 | 0 |
| Fase 5 | 0 |
| Fase 6 | 0 |
| Fase 7 | 46 |
| Fase — | 1 |

Strategi urutan implementasi: seluruh modul berada di Fase 7 (Operations) karena bergantung pada team, service project, incident (modul 32), serta infrastruktur notifikasi dan antrean job; di dalam Fase 7 dibangun dalam tiga gelombang. Gelombang 7a (fondasi alerting) memuat team & izin (JSM-OPS-001, JSM-OPS-002), objek alert dan siklus hidupnya (JSM-OPS-005 s.d. JSM-OPS-012), API Integration dan kerangka integrasi (JSM-OPS-031, JSM-OPS-032), routing rule, escalation, schedule, rotation, override, who is on-call (JSM-OPS-017 s.d. JSM-OPS-022), serta notification rule dan contact method dengan email dan push (JSM-OPS-026, JSM-OPS-027); gelombang ini sudah cukup untuk siklus "alert masuk → orang on-call dibangunkan → acknowledge". Gelombang 7b menambah policy, maintenance, zona waktu, ICS, forwarding, aksi dari notifikasi, aplikasi mobile, integrasi email/monitoring/webhook/ChatOps/tiket, keterkaitan incident, dashboard, laporan, dan REST API konfigurasi. Gelombang 7c menampung fitur P2/P3 yang bergantung pada Premium atau penyedia pihak ketiga (heartbeat, service-aware alerting, AIOps, incoming call routing, edge connector, provider Terraform, batasan plan). Opsgenie Migration (JSM-OPS-047) tidak dibangun dan hanya menjadi rujukan. Modul ini menggabungkan sembilan fitur ringkas dari lampiran JSM ke fitur rinci yang setara (lihat baris **Alias inventaris** pada JSM-OPS-001, 005, 015, 016, 018, 019, 026, 030, dan 031); P0 yang ditempatkan di gelombang 7b (Monitoring Tool Integrations JSM-OPS-034, Alert to Incident Linkage JSM-OPS-039, dan Incident Responders JSM-OPS-040) tetap wajib selesai sebelum Operations dirilis ke pelanggan.

## Fitur

## Fondasi: team operasi, izin & pengaturan

### JSM-OPS-001 Operations Teams

- **Alias inventaris:** Ops teams & operations configuration
- **Deskripsi:** Team operasional berbasis entitas Teams produk ini yang memiliki alert, schedule, escalation, routing rule, integrasi, policy, heartbeat, dan maintenance sendiri. Team adalah unit kepemilikan semua resource Operations.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (membuat/menghapus team ops, mengaktifkan operations), Team admin (mengelola anggota dan resource team), Team member (melihat dan bertindak atas alert team).
- **Aturan perilaku (acceptance criteria):**
  1. Ops admin dapat membuat team ops baru (nama unik per situs, deskripsi, anggota awal) atau mengaktifkan operations pada team yang sudah ada di direktori Teams (PLAN-057); mengaktifkan operations tidak mengubah keanggotaan team asal.
  2. Setiap anggota team ops memiliki tepat satu role: `Admin` atau `User`; minimal satu Admin harus ada, dan menghapus/menurunkan Admin terakhir ditolak dengan pesan "Team harus memiliki minimal satu admin".
  3. ⚠ Team admin dapat mendefinisikan custom team role dengan hak granular (mis. kelola schedule saja, kelola integrasi saja) yang dapat diberikan ke anggota sebagai pengganti role User.
  4. Halaman team memiliki tab Members, On-call schedules, Escalations, Routing rules, Integrations, Policies, Heartbeats, Maintenance, dan Settings; tab yang tidak boleh diakses pengguna (mis. Integrations untuk member) disembunyikan.
  5. ⚠ Saat operations diaktifkan, sistem otomatis membuat resource default: schedule `<Team>_schedule`, escalation `<Team>_escalation` (notify on-call schedule pada menit 0), routing rule default menuju escalation tersebut, dan satu API integration default milik team.
  6. Menghapus team ops meminta konfirmasi dengan mengetik nama team, menampilkan jumlah resource yang ikut dihapus (schedule, escalation, integrasi, policy, heartbeat), menolak bila resource team masih dirujuk escalation/routing team lain, dan alert historis tetap tersimpan dengan nama team sebagai teks.
  7. Team ops dapat ditambahkan sebagai responder alert maupun incident; saat ditambahkan, routing rule team (JSM-OPS-017) menentukan siapa yang dinotifikasi.
  8. Seluruh anggota team dapat melihat alert yang responder-nya mencakup team dan menjalankan aksi alert (acknowledge, close, note, dll.) atasnya.
  9. Mengaktifkan operations untuk team dicatat di audit log dan menambahkan menu Operations team pada profil team serta direktori People & Teams.
  10. Halaman operations team menampilkan ringkasan jumlah resource per kategori (schedules, escalations, routing rules, integrations, policies, maintenance, heartbeats) dengan tautan ke masing-masing tab.
  11. Role team admin dapat membuat/mengubah/menghapus resource team, sedangkan role member hanya dapat melihat konfigurasi, bertindak atas alert, dan membuat override/forwarding untuk dirinya (JSM-OPS-002).
  12. Team ops dapat dihubungkan ke satu atau lebih service project; team yang terhubung muncul sebagai pilihan responder di incident project tersebut dan alert team dapat dilihat dari menu Alerts project.
  13. Pengaturan Operations global (integrasi, policy, maintenance global) dikelola di halaman pengaturan produk, bukan di halaman team (JSM-OPS-003).
  14. Migrasi data dari tool alerting eksternal (users, teams, schedules, integrations) tidak dibangun sebagai fitur native; impor dilakukan lewat REST API konfigurasi (JSM-OPS-045) dan rujukan legacy di JSM-OPS-047.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team ops tidak terikat mode project.
- **Catatan & ketergantungan:** Entitas Teams di PLAN-057/PLAN-060 ([18](18-timeline-rencana-tampilan-proyek.md)); direktori USR-027 ([02](02-pengguna-grup-organisasi.md)); keterhubungan ke service project di modul 29; responder incident di modul 32 dan JSM-OPS-040. Keputusan terbuka: apakah team ops boleh berdiri tanpa entitas Teams umum.

### JSM-OPS-002 Operations Permissions & Roles

- **Deskripsi:** Model akses fitur Operations pada level global dan team, termasuk visibilitas alert dan pembatasan konfigurasi integrasi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (admin produk Service Management), Team admin, Team member, Agent berlisensi, Stakeholder.
- **Aturan perilaku (acceptance criteria):**
  1. Hanya pengguna dengan lisensi agent Service Management yang dapat membuka menu Operations, menjadi responder, masuk rotation, dan menerima notifikasi on-call; pengguna tanpa lisensi yang ditambahkan ke rotation ditolak dengan pesan "Pengguna memerlukan lisensi agent".
  2. Admin produk (dan Admin Situs) adalah Ops admin yang mengelola pengaturan global: global integration, global policy, global maintenance, negara SMS/voice, dan laporan global.
  3. Team admin mengelola semua resource team; team member hanya membaca konfigurasi team dan bertindak atas alert team; percobaan mengubah resource tanpa hak mengembalikan HTTP 403.
  4. Setiap member dapat membuat override dan forwarding rule untuk dirinya sendiri; membuat untuk pengguna lain memerlukan Team admin (untuk schedule team) atau Ops admin (forwarding siapa pun).
  5. Alert hanya dapat dilihat oleh responder (user langsung, anggota team responder, on-call schedule/escalation responder), pengguna/team di `visibleTo`, dan Ops admin; pengguna lain mendapat 404 saat membuka tautan alert.
  6. Konfigurasi integrasi dan integration logs hanya dapat dilihat/diubah oleh Team admin pemilik integrasi dan Ops admin; API key integrasi ditampilkan tersamarkan kecuali bagi mereka.
  7. ⚠ Ops admin dapat membuat custom global role dan custom team role dengan kumpulan hak granular (mis. "kelola integrasi", "lihat laporan"), dibatasi pada plan tertentu.
  8. ⚠ Tersedia peran stakeholder read-only yang dapat melihat alert dan on-call tanpa menjalankan aksi dan tanpa memakan lisensi agent.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Global permission dan audit perubahan izin di PERM-057 ([03](03-izin-peran-keamanan.md)); lisensi agent di modul 29; peran custom team di JSM-OPS-001.

### JSM-OPS-003 Operations Global Settings

- **Deskripsi:** Pengaturan Operations tingkat situs yang berlaku untuk semua team, diakses dari pengaturan produk Service Management.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman pengaturan global memuat tab Global integrations, Global policies, dan Global maintenance; resource global tidak dimiliki team mana pun dan dievaluasi sebelum resource team.
  2. Halaman Users ops menampilkan daftar pengguna ber-lisensi agent beserta contact method (terverifikasi/belum), team, dan status on-call; Ops admin dapat memfilter pengguna tanpa contact method terverifikasi.
  3. ⚠ Ops admin dapat membatasi negara yang diizinkan untuk SMS/voice; nomor dari negara di luar daftar tidak dapat ditambahkan sebagai contact method.
  4. ⚠ Halaman Usage menampilkan jumlah SMS dan voice call terkirim per bulan, per team, dan sisa kuota plan.
  5. ⚠ Halaman Logs global menampilkan integration/request logs lintas team dengan filter yang sama seperti integration logs team (JSM-OPS-031), hanya untuk Ops admin.
  6. ⚠ Setiap perubahan konfigurasi Operations (create/update/delete schedule, escalation, routing, integrasi, policy, maintenance, role) dicatat di audit log situs dengan aktor, waktu, nilai sebelum/sesudah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Audit log ADM-014 ([22](22-administrasi-sistem.md)); kuota plan di JSM-OPS-004; administrasi produk lain di [35](35-jsm-ai-otomasi-api-admin-csm.md).

### JSM-OPS-004 Plan Limits (Free/Standard/Premium/Enterprise)

- **Deskripsi:** Pembatasan fitur Operations per tier plan Service Management, diterapkan sebagai feature flag dan kuota konfigurasi.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Billing admin/Admin Situs (melihat plan), Ops admin (melihat dampak batasan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Plan Free mendukung hingga 3 agent dengan alerting dan on-call dasar; jumlah integrasi dibatasi oleh kuota yang dikonfigurasi dan pembuatan integrasi melebihi kuota ditolak dengan pesan upgrade.
  2. Plan Standard mengaktifkan alerting, on-call, escalation, routing rule, dan integrasi tanpa batas jumlah team.
  3. ⚠ Plan Premium mengaktifkan heartbeat, global policy, analytics lanjutan, AIOps, dan incoming call routing; pada plan lebih rendah menu tersebut tampil terkunci dengan label plan yang dibutuhkan.
  4. Plan Enterprise menambah Analytics/Data Lake, data residency, dan SLA uptime; di clone ini diperlakukan sebagai flag konfigurasi deployment.
  5. ⚠ Kuota SMS/voice dihitung per pengguna per bulan dan dapat berbeda per plan dan negara; saat kuota habis, pengiriman SMS/voice dilewati, dicatat "kuota habis" di notification log, dan notifikasi dialihkan ke push/email.
  6. ⚠ Retensi riwayat alert dan integration logs ditentukan per plan (nilai dapat dikonfigurasi); data melewati retensi dihapus oleh job harian dan tidak muncul di pencarian maupun laporan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi tier umum di [00](00-visi-ruang-lingkup-glosarium-baseline.md); billing dan pricing tidak dibangun (lihat Di luar cakupan). Angka batas final adalah keputusan pemilik produk.

## Alert inti: objek, pembuatan & siklus hidup

### JSM-OPS-005 Alerts (Alert Object & Fields)

- **Alias inventaris:** Alerts (JSM Operations)
- **Deskripsi:** Entitas alert inti yang mewakili sinyal dari tool monitoring, API, email, atau pembuatan manual, dengan field, state, dan riwayat aktivitas sendiri yang terpisah dari work item. Fitur ini juga menjadi payung pengelolaan alert yang dirinci di JSM-OPS-006 s.d. JSM-OPS-012.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Integration dan API client (membuat), Responder (membaca/bertindak), Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Field `message` wajib, maksimal 130 karakter; request tanpa message ditolak HTTP 422 dengan galat "message is required".
  2. ⚠ Field `alias` (kunci deduplikasi) maksimal 512 karakter; bila kosong sistem mengisi nilai default unik (mis. UUID alert) sehingga alert tersebut tidak pernah ter-dedup dengan alert lain.
  3. Field `description` maksimal 15.000 karakter dan dirender sebagai teks biasa/markdown terbatas.
  4. Field `priority` bernilai P1-P5 dengan default P3 bila tidak dikirim.
  5. Field `tags` maksimal 20 tag; tag duplikat dalam satu alert digabung.
  6. Extra properties (`details`) berupa pasangan key-value string dengan total ukuran maksimal 8.000 karakter.
  7. Field `entity` maksimal 512 karakter; field `source` maksimal 100 karakter dengan default IP pengirim (API) atau nama integrasi.
  8. Field `responders` menerima tipe `user`, `team`, `schedule`, dan `escalation` (dirujuk via id atau nama/username), maksimal 50 entri; responder yang tidak ditemukan dicatat sebagai peringatan di activity log tanpa menggagalkan pembuatan.
  9. Field `visibleTo` (team/user yang boleh melihat tanpa dinotifikasi) maksimal 50 entri.
  10. Field `user` (aktor pembuat, untuk ditampilkan di activity log) dan `note` awal opsional dapat dikirim saat create.
  11. Field `actions` berisi custom action per alert, maksimal 10 action dengan nama maksimal 50 karakter, yang diteruskan ke integrasi outgoing (JSM-OPS-010).
  12. Field read-only `id` (UUID), `tinyId` (nomor pendek bertambah per situs), `createdAt`, `updatedAt`, `count`, `lastOccurredAt`, `owner`, dan `integration` (id/type/name) diisi sistem dan diabaikan bila dikirim klien.
  13. Field laporan read-only `ackTime`, `closeTime` (milidetik sejak create), `acknowledgedBy`, dan `closedBy` diisi saat aksi terkait terjadi.
  14. Flag status: `status` (`open`/`closed`), `acknowledged`, `isSeen`, `snoozed` beserta `snoozedUntil`.
  15. ⚠ Attachment dapat ditambahkan ke alert via UI dan API dengan batas ukuran per file yang dapat dikonfigurasi (default mengikuti batas attachment situs).
  16. ⚠ Nilai yang melebihi batas field dipotong ke panjang maksimum untuk `message`, `description`, `entity`, `source`, dan ditolak (HTTP 422, menyebut nama field) untuk `tags`, `responders`, `visibleTo`, `actions`, dan `details`; perilaku pemotongan dicatat di activity log.
  17. Halaman daftar alert menyediakan filter, pencarian, dan saved search sesuai JSM-OPS-012.
  18. Atribut minimum alert yang selalu tersedia di UI dan API adalah message, alias, description, priority P1-P5, tags, responders, extra properties, entity, dan source.
  19. Aksi Acknowledge, Unacknowledge, Close, Snooze, Escalate to next, Assign, Add note, Add responder, Add tags, Create incident, dan custom action tersedia sesuai JSM-OPS-010.
  20. Deduplikasi berdasarkan alias dengan penghitung `count` berlaku sesuai JSM-OPS-008.
  21. ⚠ Pengelompokan/korelasi alert berbasis AI tersedia hanya pada plan yang mengaktifkannya (JSM-OPS-042).
  22. Alert policy (modify, suppress, auto-close, delay) pada level global dan team dievaluasi saat alert dibuat sesuai JSM-OPS-013 dan JSM-OPS-014.
  23. Setiap alert memiliki activity log/timeline kronologis yang tidak dapat diubah (JSM-OPS-011).
  24. Ops admin melihat alert lintas semua team (tampilan global), sedangkan anggota team melihat alert yang terkait team-nya (tampilan per team); keduanya memakai daftar yang sama dengan filter team.
  25. Bulk action pada alert terpilih tersedia di daftar alert (JSM-OPS-012).
  26. Responder dinotifikasi sesuai notification rule personal masing-masing (JSM-OPS-026).
  27. Alert API mendukung create, update, dan close alert (JSM-OPS-045).
- **Team-managed vs company-managed:** Tidak ada perbedaan; alert tidak terikat project.
- **Catatan & ketergantungan:** Alert bukan work item; tautan ke work item lewat JSM-OPS-038 dan ke incident lewat JSM-OPS-039. Model data dan katalog event domain di [28](28-model-data-inti.md) perlu ditambah event `alert.*`. Repo belum punya entitas ini; tabel `ops_alerts`, `ops_alert_activities`, `ops_alert_responders` adalah migrasi baru.

### JSM-OPS-006 Alert Priority & Tags

- **Deskripsi:** Skema prioritas P1-P5 dan tag bebas untuk klasifikasi, routing, policy, filter, dan notification rule.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder (mengubah manual), Integration (mengirim), Ops admin/Team admin (memetakan severity).
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia lima level tetap: P1 Critical, P2 High, P3 Moderate, P4 Low, P5 Informational, masing-masing dengan warna dan ikon tetap; level tidak dapat ditambah atau diubah namanya.
  2. Alert yang dibuat tanpa priority mendapat P3.
  3. Pengaturan integrasi menyediakan tabel mapping severity tool monitoring (mis. `critical`, `warning`, `info`) ke P1-P5, dengan nilai fallback bila severity tidak dikenali.
  4. Priority dapat diubah manual oleh responder, oleh alert policy, atau via API; setiap perubahan tercatat di activity log dengan nilai lama dan baru.
  5. Tag bebas (maksimal 20 per alert, tanpa daftar master) dapat dipakai sebagai kondisi routing rule, policy, notification rule, dan filter pencarian.
  6. ⚠ Kondisi tag pada rule mendukung operator `contains`, `contains any`, dan `contains all`, dengan negasi.
  7. Notification rule dan routing rule dapat dibatasi ke satu atau beberapa priority (mis. hanya P1-P2).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Priority alert terpisah dari field Priority work item ([05](05-konfigurasi-work-item-skema.md)); mapping ke priority incident saat create incident di JSM-OPS-039.

### JSM-OPS-007 Alert Creation (Manual, API, Email, Integrations)

- **Deskripsi:** Seluruh jalur pembuatan alert: dialog Create alert di UI, Alert API dengan API key integrasi, email integration, integrasi monitoring, heartbeat expired, responder incident, otomasi, dan ChatOps.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder (manual), Integration, API client, Automation, System.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Create alert manual memuat message (wajib), responders, priority, description, tags, extra properties, alias, dan entity; tombol Create nonaktif sampai message terisi.
  2. Alert API `POST /api/v1/ops/alerts` diautentikasi dengan header `Authorization: <skema-key> <integration API key>`; key tidak valid atau milik integrasi disabled mengembalikan HTTP 401/403.
  3. Pemrosesan asinkron: API mengembalikan HTTP 202 dengan `requestId`, dan `GET /api/v1/ops/alerts/requests/{requestId}` mengembalikan `isSuccess`, `status`, `alertId`, serta pesan galat bila gagal.
  4. Email ke alamat email integration membuat alert dengan subject sebagai message dan body sebagai description (aturan parsing di JSM-OPS-033).
  5. Integrasi monitoring memetakan payload tool ke field alert sesuai field mapping integrasi (JSM-OPS-034).
  6. Heartbeat yang expired membuat alert dengan source sistem `Heartbeat` (JSM-OPS-016).
  7. Menambahkan responder ke incident Service Management membuat alert terkait incident untuk menotifikasi responder tersebut (JSM-OPS-040).
  8. ⚠ Otomasi menyediakan action "Create alert" (dan integrasi project Service Management) dengan pengisian field alert dari smart value.
  9. ⚠ Pengguna dapat membuat alert dari ChatOps (perintah atau shortcut pesan) di Slack/Microsoft Teams bila aplikasi terhubung.
  10. Bila `responders` kosong, responder default diambil dari team pemilik integrasi lalu diteruskan routing rule team tersebut; integrasi global tanpa responder menghasilkan alert tanpa penerima yang ditandai di activity log.
  11. Alert tidak dibuat bila integrasi disabled, integrasi sedang dalam maintenance, atau payload tidak lolos filter rule integrasi; request status melaporkan alasan (`integration disabled`, `in maintenance`, `filtered`).
  12. Alert baru yang alias-nya sama dengan alert open tidak dibuat baru tetapi di-dedup (JSM-OPS-008); request status mengembalikan `alertId` alert yang ada.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Skema header key adalah keputusan terbuka (lihat Pertanyaan terbuka); kompatibilitas path dengan API rujukan (`/v2/alerts`) opsional. Otomasi di [20](20-otomasi.md); antrean asinkron di [27](27-nfr-arsitektur-platform.md).

### JSM-OPS-008 Alert Deduplication (Alias)

- **Deskripsi:** Alert baru dengan alias sama dengan alert yang masih open tidak membuat alert baru, melainkan menaikkan `count`, untuk mencegah banjir alert dan notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Integration, System.
- **Aturan perilaku (acceptance criteria):**
  1. Deduplikasi hanya dilakukan terhadap alert berstatus `open` dengan alias identik (case-sensitive) dalam situs yang sama.
  2. Setiap duplikat menaikkan `count` sebesar 1 dan memperbarui `lastOccurredAt`; field lain alert asal tidak ditimpa kecuali oleh policy.
  3. Setiap kejadian duplikat dicatat di activity log alert asal ("Alert deduplicated, count: N") beserta sumbernya.
  4. Bila alert dengan alias tersebut sudah `closed`, kiriman berikutnya membuat alert baru dengan `count` = 1.
  5. Secara default duplikat tidak memicu notifikasi ulang; perilaku notifikasi untuk duplikat diatur lewat Deduplication action pada notification policy (JSM-OPS-014).
  6. Alias dapat diisi dari dynamic field atau ekstraksi regex di pengaturan integrasi (mis. `{{host}}-{{check}}`).
  7. Integrasi dan API dapat menutup, meng-acknowledge, atau menambah note berdasarkan alias (`identifierType=alias`); alias tanpa alert open menghasilkan request status gagal "alert not found".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Deduplikasi harus atomik terhadap request paralel (unique index parsial pada `(alias) WHERE status='open'`); lihat [27](27-nfr-arsitektur-platform.md).

### JSM-OPS-009 Alert Lifecycle & States

- **Deskripsi:** Siklus hidup alert dari open sampai closed, termasuk acknowledged, snoozed, dan seen, beserta efeknya pada eskalasi dan notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, System, Integration.
- **Aturan perilaku (acceptance criteria):**
  1. Alert hanya memiliki dua status utama, `open` dan `closed`; acknowledged, snoozed, dan seen adalah flag tambahan.
  2. Acknowledge menghentikan langkah escalation berkondisi "if not acknowledged" dan notifikasi berulang, tetapi langkah "if not closed" tetap berjalan.
  3. Unacknowledge mengembalikan alert ke open-unacked dan langkah escalation yang tersisa dapat berjalan lagi dari posisi waktu saat ini.
  4. Snooze menahan notifikasi sampai `snoozedUntil`; saat waktu habis, flag snoozed dilepas, alert kembali unacked, dan notifikasi/escalation dimulai ulang.
  5. Flag seen ditandai per pengguna saat pengguna membuka detail alert; `isSeen` bernilai true bila minimal satu responder sudah melihat.
  6. ⚠ Close menghentikan seluruh escalation dan notifikasi; alert closed hanya dapat menerima note, tag, dan attachment, sedangkan aksi lain ditolak dengan pesan "Alert sudah ditutup".
  7. Alert dapat ditutup otomatis oleh integrasi (status recovery), oleh Auto-close policy, atau oleh opsi close setelah repeat escalation selesai; aktornya dicatat sebagai integrasi/policy/sistem.
  8. Tidak ada aksi reopen; alias yang muncul lagi setelah close membuat alert baru.
  9. Owner diisi otomatis dengan pengguna yang pertama melakukan acknowledge, assign, atau take ownership; acknowledge berikutnya tidak menimpa owner.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine escalation di JSM-OPS-018; timer snooze dijalankan worker latar belakang.

### JSM-OPS-010 Alert Actions

- **Deskripsi:** Tindakan responder terhadap alert dari web, mobile, ChatOps, notifikasi, atau API, termasuk custom action yang didefinisikan lewat integrasi outgoing.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, Ops admin, Team admin, Integration, API client.
- **Aturan perilaku (acceptance criteria):**
  1. Acknowledge dan Unacknowledge tersedia sebagai tombol yang saling bertukar sesuai state.
  2. Close menutup alert dan meminta note opsional.
  3. Snooze dapat diatur sampai tanggal/jam tertentu atau durasi preset (15 menit, 30 menit, 1 jam, 4 jam, 1 hari); waktu di masa lalu ditolak.
  4. Escalate to next melompatkan alert ke langkah escalation berikutnya atau ke escalation tertentu yang dipilih.
  5. Assign menetapkan owner ke pengguna tertentu; Take ownership menetapkan pengguna saat ini sebagai owner.
  6. Add responder (user/team/schedule/escalation) menambah responder dan memicu notifikasi baru hanya kepada responder yang ditambahkan.
  7. Add note menambah catatan teks (maksimal 25.000 karakter) yang tampil di tab Notes.
  8. Add/remove tags mengubah daftar tag dengan tetap menghormati batas 20 tag.
  9. Add/remove extra properties (details) mengubah key-value tanpa melebihi batas total 8.000 karakter.
  10. Update priority, message, dan description tersedia bagi responder dengan validasi batas field yang sama seperti create.
  11. Create incident atau link ke incident yang ada dapat dilakukan dari alert (JSM-OPS-039).
  12. ⚠ Delete alert hanya dapat dilakukan Ops admin dan Team admin pemilik; alert yang dihapus hilang dari pencarian dan laporan, dan penghapusan dicatat di audit log.
  13. Setiap aksi tercatat di activity log (aktor, channel, waktu) dan dapat memicu integrasi outgoing yang berlangganan aksi tersebut.
  14. Aksi via API menerima identifier `id`, `tiny`, atau `alias` (`identifierType`) dan diproses asinkron dengan `requestId`.
  15. Custom action didefinisikan sebagai nama action pada konfigurasi integrasi outgoing (edge connector, webhook, atau integrasi lain) atau dikirim di field `actions` saat create alert.
  16. ⚠ Custom action dapat diberi filter kondisi (tag, source, priority, dll.) sehingga hanya tampil dan berlaku pada alert yang cocok.
  17. ⚠ Tombol custom action tampil di halaman detail alert dan aplikasi mobile; di daftar alert tersedia lewat menu "More".
  18. Eksekusi custom action diteruskan ke edge connector untuk menjalankan skrip di jaringan internal atau ke outgoing webhook dengan payload alert.
  19. ⚠ Eksekusi custom action dan hasilnya tercatat di activity log alert dan integration logs; kegagalan atau timeout (default 60 detik) dicatat sebagai note "Action <nama> gagal".
  20. Hanya pengguna yang dapat mengakses alert (responder, `visibleTo`, anggota team terkait) yang dapat menjalankan custom action; lainnya mendapat 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aksi dari notifikasi di JSM-OPS-028, mobile di JSM-OPS-029, ChatOps di JSM-OPS-037, edge connector di JSM-OPS-036.

### JSM-OPS-011 Alert Details Page

- **Deskripsi:** Halaman/panel detail alert yang menampilkan seluruh informasi, tombol aksi, dan riwayat satu alert.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, Ops admin (dengan hak lihat alert sesuai JSM-OPS-002).
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan `#tinyId`, message, badge priority, status (open/acked/snoozed/closed), `count`, dan tombol aksi utama (Acknowledge, Close, Snooze, Escalate, More).
  2. Bagian Details menampilkan description, tags, extra properties, entity, source, alias, dan integrasi asal; nilai kosong disembunyikan.
  3. Panel Responders dan Visible to menampilkan entri beserta tipe (user/team/schedule/escalation) dan tombol tambah sesuai izin.
  4. Tab Notes menampilkan note berurutan waktu (terbaru di atas) dengan penulis dan channel.
  5. Tab Activity log menampilkan semua perubahan state dan aksi beserta aktor dan channel (create, dedup, acknowledge, escalate, assign, aksi integrasi, policy).
  6. Tab Notification/escalation log menampilkan setiap pengiriman notifikasi per pengguna per contact method (email, SMS, voice, push) dengan status terkirim/gagal/dilewati, waktu, serta progres langkah escalation.
  7. ⚠ Log penerapan policy menampilkan alert policy dan notification policy yang cocok beserta aksi yang diterapkan (mis. "Priority diubah P3 → P1 oleh policy X").
  8. Bagian Linked incident dan Affected services menampilkan incident tertaut dan service terdampak dengan tautan.
  9. Attachment dapat diunduh; penghapusan hanya bagi pengunggah dan Team admin/Ops admin.
  10. Informasi siapa yang acknowledge/close beserta waktunya serta daftar pengguna yang sudah melihat (seen) ditampilkan.
  11. Tombol Salin tautan menyalin URL permanen alert `/ops/alerts/{id}` ke clipboard.
  12. Panel insight AI (similar alerts, past responders, related resources) tampil hanya bila AIOps aktif (JSM-OPS-042).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman diperbarui realtime (event `alert.updated`) mengikuti infrastruktur realtime di [27](27-nfr-arsitektur-platform.md).

### JSM-OPS-012 Alert List, Search Query & Saved Searches

- **Deskripsi:** Halaman daftar alert (Operations > Alerts dan konteks service project) dengan bahasa query, filter cepat, bulk action, dan saved search.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, Ops admin; hasil dibatasi hak lihat alert.
- **Aturan perilaku (acceptance criteria):**
  1. Quick filter tersedia: Open, Unacked, Not seen, Assigned to me, Alert team saya, Closed; quick filter dapat dikombinasikan dengan query.
  2. Bahasa query mendukung field `status`, `priority`, `tag`, `owner`, `teams`, `responders`, `message`, `description`, `alias`, `entity`, `source`, `count`, `acknowledged`, `isSeen`, `snoozed`, `createdAt`, `integration.name`, dan `details.<key>`.
  3. Operator yang didukung: `AND`, `OR`, `NOT`, `:`, `>`, `<`, `>=`, `<=`, wildcard `*`, dan tanda kutip untuk frasa; query tidak valid menampilkan galat dengan posisi token yang salah.
  4. Filter rentang waktu menyediakan preset (1 jam, 24 jam, 7 hari, 30 hari) dan rentang tanggal kustom.
  5. ⚠ Saved search dapat disimpan, diubah nama, dan dihapus oleh pemiliknya; berbagi saved search ke team tersedia sebagai opsi.
  6. Hasil dapat diurutkan berdasarkan created, updated, atau priority, dan pengguna dapat memilih kolom yang tampil (pilihan tersimpan per pengguna).
  7. Bulk action pada alert terpilih: acknowledge, close, snooze, add note, add tags, assign, delete; bulk delete hanya untuk Team admin/Ops admin dan setiap alert dicatat terpisah di activity log.
  8. Hasil dipaginasi (default 20, maksimal 100 per halaman) dengan total jumlah hasil.
  9. Daftar diperbarui otomatis saat alert baru masuk atau berubah tanpa memuat ulang halaman, dengan indikator "N alert baru".
  10. ⚠ Daftar alert hasil query dapat diekspor ke CSV (maksimal 10.000 baris).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bahasa query alert terpisah dari JQL ([12](12-jql-bahasa-kueri.md)) tetapi parser dapat berbagi komponen tokenizer.

## Policy, maintenance & heartbeat

### JSM-OPS-013 Alert Policies (Modify Alert)

- **Deskripsi:** Aturan yang memodifikasi alert saat dibuat berdasarkan kondisi, pada level team atau global.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Team admin (policy team), Ops admin (policy global).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Policy dapat dibuat di level team (oleh Team admin, berlaku untuk alert yang responder-nya mencakup team) dan level global (oleh Ops admin, berlaku untuk semua alert; tersedia pada plan Premium ke atas).
  2. Filter kondisi memiliki mode Match all alerts, Match all conditions, atau Match any condition.
  3. Field kondisi: message, alias, description, source, entity, tags, actions, details (key/value), priority, responders/recipients, integration, dan teams.
  4. Operator kondisi: is, contains, starts with, ends with, matches (regex), is empty, greater than, less than, masing-masing dapat dinegasikan dengan NOT; regex tidak valid ditolak saat simpan.
  5. Time restriction opsional: jam tertentu setiap hari atau kombinasi hari+jam, dievaluasi dalam time zone yang dipilih.
  6. Modifikasi yang dapat diterapkan: message, alias, description, entity, source, priority, tags, extra properties, responders, dan actions.
  7. Nilai modifikasi mendukung dynamic field seperti `{{message}}`, `{{tags}}`, `{{extraProperties.key}}` serta opsi append ke nilai asli (bukan menimpa).
  8. Opsi "Continue processing" menentukan apakah policy berikutnya tetap dievaluasi setelah policy ini cocok; default berhenti di policy pertama yang cocok.
  9. Policy dapat diurutkan ulang (drag), di-enable/disable, diduplikasi, dan dihapus; urutan global dievaluasi sebelum urutan team.
  10. Policy dapat dinonaktifkan sementara oleh maintenance window (JSM-OPS-015) dan kembali aktif otomatis setelahnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Evaluator kondisi dipakai bersama oleh notification policy, routing rule, notification rule, dan filter integrasi; dirancang sebagai satu komponen `condition` di domain.

### JSM-OPS-014 Notification Policies (Delay/Suppress, Auto-close, Auto-restart, Deduplication action)

- **Deskripsi:** Policy yang mengatur kapan notifikasi dikirim, ditunda, ditekan, atau diulang, serta penutupan alert otomatis.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Team admin, Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Delay/Suppress notifications: tunda notifikasi X menit, tunda sampai waktu tertentu, atau suppress sepenuhnya; alert tetap dibuat dan tampil di daftar meskipun notifikasi ditekan.
  2. Auto-close policy menutup alert otomatis X menit/jam setelah dibuat bila belum ditutup, dengan aktor "Auto-close policy".
  3. Auto-restart notifications memulai ulang notifikasi dan escalation bila alert belum closed setelah X menit, berulang hingga N kali (N maksimal 20).
  4. Deduplication action mendukung mode value based (kirim notifikasi saat `count` mencapai N) dan frequency based (kirim bila terjadi N kali dalam X menit).
  5. Filter kondisi dan time restriction memakai komponen yang sama dengan alert policy (JSM-OPS-013).
  6. ⚠ Notification policy tersedia di level team dan level global; level global terbatas pada plan Premium ke atas.
  7. Notification policy dapat diurutkan, di-enable/disable, dan dapat dinonaktifkan/diaktifkan oleh maintenance window.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Di UI, alert policy dan notification policy dapat tampil sebagai satu daftar Policies dengan kolom tipe; keputusan tampilan di Pertanyaan terbuka.

### JSM-OPS-015 Maintenance Windows

- **Alias inventaris:** Maintenance windows
- **Deskripsi:** Penjadwalan periode pemeliharaan yang menonaktifkan integrasi dan/atau policy sementara agar on-call tidak terganggu alert saat pekerjaan terencana.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Team admin (maintenance team), Ops admin (maintenance global).
- **Aturan perilaku (acceptance criteria):**
  1. Durasi dapat ditetapkan sebagai X menit/jam mulai sekarang atau rentang terjadwal start-end; end harus setelah start dan start terjadwal tidak boleh di masa lalu.
  2. Rules maintenance memilih integrasi dan policy beserta aksinya: integrasi di-disable, policy di-enable atau di-disable selama periode.
  3. Maintenance dapat dibuat di level team (hanya resource team) dan level global (resource global dan semua team).
  4. Maintenance memiliki deskripsi wajib yang tampil di daftar dan di Team Dashboard.
  5. Daftar maintenance dikelompokkan menjadi Ongoing, Scheduled, dan Past.
  6. Maintenance yang sedang berjalan atau terjadwal dapat dibatalkan; pembatalan mengembalikan status resource ke kondisi sebelum maintenance secara langsung.
  7. Maintenance terjadwal dapat diedit; maintenance yang sudah berjalan hanya dapat diubah waktu end-nya.
  8. API maintenance mendukung create, list, get, update, cancel, dan delete.
  9. Maintenance dapat dimulai segera ("now") atau dijadwalkan untuk waktu mendatang.
  10. Ruang lingkup maintenance dibatasi ke integrasi dan/atau policy tertentu yang dipilih; resource lain tidak terpengaruh.
  11. Aksi yang tersedia per resource adalah disable integration, enable policy, atau disable policy.
  12. Status maintenance dihitung otomatis: `scheduled` sebelum start, `active` selama periode, `past` setelah end atau dibatalkan.
  13. Maintenance dapat dihentikan lebih awal; waktu end dicatat sebagai waktu pembatalan dan status menjadi `past`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alert yang ditolak karena maintenance tercatat di integration logs (JSM-OPS-031); heartbeat dalam maintenance di JSM-OPS-016.

### JSM-OPS-016 Heartbeats

- **Alias inventaris:** Heartbeats
- **Deskripsi:** Pemantauan liveness: sistem eksternal mengirim ping berkala, dan bila ping tidak datang dalam interval, alert dibuat otomatis.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Team admin, Ops admin, API client (ping), System.
- **Aturan perilaku (acceptance criteria):**
  1. Heartbeat memiliki field name (unik per situs, tanpa spasi, dipakai di URL ping), description, interval + unit (minutes/hours/days, minimal 1 menit), dan enabled.
  2. Heartbeat dimiliki satu owner team, dan alert expired dirutekan ke team tersebut.
  3. Konfigurasi alert expired memuat message (default "HeartbeatName is expired"), tags, dan priority.
  4. Ping dilakukan via `GET`/`POST /api/v1/ops/heartbeats/{name}/ping` dengan API key integrasi; heartbeat tidak ditemukan mengembalikan 404.
  5. ⚠ Ping juga dapat dikirim via email ke alamat integrasi dengan subject berisi nama heartbeat.
  6. Status heartbeat ditampilkan sebagai active, expired, atau disabled beserta waktu ping terakhir.
  7. Saat tidak ada ping dalam satu interval sejak ping terakhir, status menjadi expired dan alert "heartbeat expired" dibuat (sumber sistem) dengan alias tetap per heartbeat agar tidak berlipat.
  8. ⚠ Ping berikutnya setelah expired mengembalikan status ke active dan menutup alert expired secara otomatis.
  9. Heartbeat dapat di-enable/disable, diedit, dan dihapus via UI dan API.
  10. ⚠ Heartbeat yang disabled atau berada dalam maintenance tidak memicu alert, dan timer interval dihitung ulang saat aktif kembali.
  11. Membuat heartbeat memerlukan nama, interval, dan tim pemilik; validasi gagal menampilkan galat per field.
  12. Endpoint API khusus ping tersedia dan mengembalikan HTTP 202 tanpa body tambahan saat sukses.
  13. Saat expired, alert dibuat otomatis dengan message dan priority yang dapat diatur di konfigurasi heartbeat.
  14. Enable/disable heartbeat tersedia sebagai toggle di daftar heartbeat.
  15. Daftar heartbeat menampilkan status dan waktu ping terakhir untuk setiap heartbeat.
  16. ⚠ Heartbeat hanya tersedia pada plan Premium dan Enterprise; pada plan lebih rendah menu tampil terkunci.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Timer dijalankan worker penjadwal ([27](27-nfr-arsitektur-platform.md)); plan di JSM-OPS-004; widget status di JSM-OPS-043.

## Routing & eskalasi

### JSM-OPS-017 Team Routing Rules

- **Deskripsi:** Aturan per team yang menentukan alert untuk team diteruskan ke escalation atau schedule mana, berdasarkan kondisi field alert dan waktu.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Team admin, Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Routing rule disusun dalam daftar berurutan dan dievaluasi dari atas; rule pertama yang cocok yang dipakai, rule lain diabaikan.
  2. Kondisi rule memiliki mode Match all alerts, Match all conditions, atau Match any condition pada field alert (komponen kondisi JSM-OPS-013).
  3. Time restriction opsional (harian atau hari+jam) dievaluasi dalam time zone milik rule.
  4. Tujuan rule adalah satu escalation, satu schedule, atau "No one" (alert tetap dibuat tanpa notifikasi, tercatat di activity log).
  5. Default routing rule selalu ada di posisi terakhir, tidak dapat dihapus maupun dipindah, dan hanya tujuannya yang dapat diubah.
  6. Team admin dapat menambah, mengedit, mengurutkan ulang (drag), dan menghapus rule non-default.
  7. Default routing rule dibuat otomatis saat operations team diaktifkan dengan tujuan escalation default team.
  8. Routing diterapkan setiap kali team ditambahkan sebagai responder alert atau incident, termasuk saat ditambahkan belakangan; activity log mencatat rule yang dipakai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Escalation di JSM-OPS-018; schedule di JSM-OPS-019; time zone di JSM-OPS-023.

### JSM-OPS-018 Escalation Policies

- **Alias inventaris:** Escalation policies & routing rules
- **Deskripsi:** Rangkaian langkah notifikasi bertahap yang berjalan bila alert belum di-acknowledge atau di-close dalam waktu tertentu, dipakai bersama routing rule team.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Team admin, Ops admin, System (engine escalation).
- **Aturan perilaku (acceptance criteria):**
  1. Escalation memiliki nama (unik per situs), deskripsi, dan owner team.
  2. Escalation terdiri dari langkah dengan delay menit sejak alert dibuat (langkah pertama boleh 0 menit); delay setiap langkah harus lebih besar atau sama dengan langkah sebelumnya.
  3. Setiap langkah memiliki kondisi "if alert is not acknowledged" atau "if alert is not closed"; langkah dilewati bila kondisi sudah tidak terpenuhi saat jatuh tempo.
  4. Penerima langkah dapat berupa user, schedule (on-call user(s), next on-call, previous on-call), atau team (all members, random member).
  5. ⚠ Repeat escalation dapat diaktifkan setiap X menit setelah langkah terakhir hingga N kali (N maksimal 20).
  6. Opsi saat repeat: kembalikan alert ke unacknowledged dan unseen sebelum siklus berikutnya.
  7. Opsi close alert otomatis setelah semua repeat selesai tanpa penutupan manual.
  8. Escalation default per team dibuat otomatis saat operations diaktifkan (notify on-call schedule team pada menit 0).
  9. Escalation dapat dipakai sebagai tujuan routing rule, responder alert, atau participant rotation.
  10. Progres langkah escalation (dijadwalkan, dijalankan, dilewati, beserta penerima) terlihat di notification/escalation log alert.
  11. Menghapus escalation yang masih dipakai routing rule, rotation, atau alert open ditolak dengan daftar pemakainya; Team admin harus mengganti rujukan terlebih dulu.
  12. Langkah escalation dapat menotifikasi schedule, user, atau team dengan delay menit masing-masing.
  13. Escalation hanya berjalan selama kondisi "belum di-ack" atau "belum di-close" terpenuhi; acknowledge/close menghentikan langkah berkondisi terkait.
  14. Repeat escalation berhenti setelah N kali atau saat alert closed, mana yang lebih dulu.
  15. Routing rule team berdasarkan kondisi alert (tags, priority, source) mengarahkan alert ke escalation atau schedule sesuai JSM-OPS-017.
  16. Routing berbasis waktu memungkinkan tujuan berbeda untuk jam kerja dan di luar jam kerja (dua rule dengan time restriction berbeda).
  17. Setiap team selalu memiliki default routing rule yang menangkap alert yang tidak cocok rule lain.
  18. Responder dapat menjalankan Escalate to next secara manual dari alert untuk langsung memicu langkah berikutnya tanpa menunggu delay.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine escalation adalah worker berbasis timer yang harus tahan restart (jadwal langkah dipersistenkan); lihat [27](27-nfr-arsitektur-platform.md).

## On-call: schedule, rotation & override

### JSM-OPS-019 On-call Schedules

- **Alias inventaris:** On-call schedules
- **Deskripsi:** Jadwal on-call per team yang menentukan siapa yang bertugas, terdiri dari satu atau lebih rotation dan override, dengan time zone sendiri.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Team admin (kelola), Responder (melihat, override diri), Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Schedule dibuat dengan nama (unik per situs), deskripsi, owner team, time zone, dan status enabled/disabled.
  2. ⚠ Satu schedule dapat memiliki banyak rotation; bila tumpang tindih, rotation dengan urutan lebih atas didahulukan pada baris final schedule (layering), dan urutan dapat diubah dengan drag.
  3. Timeline view tersedia dalam rentang hari, minggu, 2 minggu, dan bulan dengan navigasi maju/mundur dan tombol Today.
  4. Baris final schedule menampilkan hasil gabungan rotation dan override yang benar-benar dipakai untuk notifikasi.
  5. Rentang waktu tanpa on-call ditandai sebagai gap berwarna peringatan pada timeline.
  6. Schedule yang disabled tidak menghasilkan on-call: escalation/routing yang menunjuknya melewati langkah tersebut dan mencatat "schedule disabled" di log alert.
  7. ⚠ Schedule default `<Team>_schedule` dibuat otomatis saat operations team diaktifkan dengan satu rotation weekly berisi anggota team.
  8. Menghapus schedule yang dipakai escalation atau routing rule menampilkan peringatan berisi daftar pemakainya dan meminta konfirmasi; rujukan yang tersisa ditandai rusak.
  9. Profil pengguna menampilkan jadwal on-call milik pengguna (shift mendatang lintas semua schedule).
  10. Schedule dibuat per team dengan time zone yang menentukan interpretasi jam rotation.
  11. Rotation di dalam schedule mendukung tipe daily, weekly, dan custom dengan durasi shift dan urutan participant (JSM-OPS-020).
  12. Restriction pada rentang jam/hari tertentu dapat diterapkan per rotation (JSM-OPS-020).
  13. Override untuk mengganti on-call sementara dapat dibuat di atas schedule (JSM-OPS-021).
  14. Beberapa rotation berlapis dalam satu schedule didukung, misalnya rotation jam kerja dan rotation malam.
  15. Timeline (hari/minggu/bulan) dan panel "who is on-call now" tampil di halaman schedule.
  16. Schedule dapat diekspor ke kalender melalui iCal/webcal (JSM-OPS-024).
  17. Pengguna menerima notifikasi sebelum shift dimulai sesuai rule Schedule start (JSM-OPS-026).
  18. Schedule dapat di-enable dan di-disable kapan saja tanpa kehilangan konfigurasi rotation.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perhitungan final schedule adalah fungsi murni (rotation + override + restriction + time zone) yang harus punya unit test DST; dipakai ulang oleh who is on-call, escalation, ICS, dan laporan.

### JSM-OPS-020 Rotations

- **Deskripsi:** Pola pergantian on-call di dalam schedule dengan tipe, participant, waktu handoff, dan restriction.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Team admin.
- **Aturan perilaku (acceptance criteria):**
  1. Tipe rotation: daily, weekly, hourly, dan custom (panjang n jam/hari/minggu, n ≥ 1).
  2. Rotation memiliki start date/time wajib dan end date opsional; setelah end, rotation tidak menghasilkan shift.
  3. Handoff time menentukan jam pergantian; untuk weekly juga hari handoff.
  4. Participant dapat berupa user, team, escalation, atau "No one" untuk gap yang disengaja; participant tanpa lisensi agent ditolak.
  5. Urutan participant dapat diatur dengan drag dan menentukan giliran shift.
  6. Restriction mendukung mode time of day (mis. 09:00-17:00) atau weekday and time of day (mis. Senin 09:00 - Jumat 17:00), dengan beberapa interval per rotation.
  7. Rotation memiliki nama dan warna yang dipakai di timeline.
  8. Mengedit rotation hanya berlaku untuk shift mulai saat ini; shift yang sudah lewat dan laporan on-call historis tidak berubah.
  9. Rotation dapat dihapus; penghapusan tidak menghapus override yang sudah dibuat di atasnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Histori shift disimpan sebagai snapshot saat rotation diubah agar laporan JSM-OPS-044 konsisten.

### JSM-OPS-021 Schedule Overrides

- **Deskripsi:** Penggantian on-call sementara di atas rotation untuk tukar shift, cuti, atau cover darurat.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder (untuk dirinya), Team admin (untuk siapa pun di schedule team).
- **Aturan perilaku (acceptance criteria):**
  1. Override dibuat dengan user pengganti, start time, dan end time; end harus setelah start.
  2. Override dapat berlaku untuk rotation tertentu atau seluruh schedule (field `rotations` di API kosong berarti seluruh schedule).
  3. Take on-call cepat tersedia dengan durasi preset (30 menit, 1 jam, 2 jam, 4 jam, sampai akhir shift), terutama dari aplikasi mobile.
  4. Override dapat diedit dan dihapus; override yang sudah lewat tidak dapat diubah.
  5. Override tampil di baris Overrides tersendiri dan tercermin di baris final schedule.
  6. API override mendukung create, update, delete, dan list per schedule.
  7. ⚠ Team admin dapat membuat override untuk siapa pun; member hanya dapat membuat override di mana dirinya menjadi pengganti.
  8. ⚠ Pengguna yang digantikan dan pengguna pengganti menerima notifikasi saat override dibuat, diubah, atau dihapus.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Izin di JSM-OPS-002; ICS mencerminkan override (JSM-OPS-024).

### JSM-OPS-022 Who Is On-call

- **Deskripsi:** Tampilan dan API untuk mengetahui siapa yang sedang, berikutnya, atau sebelumnya on-call per schedule atau team.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, Agent Layanan, API client; dibatasi hak lihat schedule.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman On-call menampilkan on-call saat ini per team dan per schedule, termasuk contact method yang dapat dihubungi.
  2. Next on-call dan previous on-call ditampilkan beserta waktu shift.
  3. Daftar dapat ditampilkan flat (daftar user unik) atau per-rotation; API menerima parameter `flat=true`.
  4. API menerima parameter `date` (ISO 8601) untuk menanyakan on-call pada waktu tertentu.
  5. Widget on-call tampil di Team Dashboard (JSM-OPS-043).
  6. Navigasi global menampilkan indikator bila pengguna saat ini sedang on-call beserta sisa waktu shift.
  7. Halaman On-call memiliki pencarian team/schedule.
  8. Endpoint `GET /api/v1/ops/schedules/{id}/on-calls` dan `GET /api/v1/ops/schedules/{id}/next-on-calls` tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai juga oleh panel responder incident di modul 32.

### JSM-OPS-023 Time Zones (Schedules & Users)

- **Deskripsi:** Penanganan zona waktu pada schedule, rotation, restriction, rule, profil pengguna, dan tampilan.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Team admin, Responder.
- **Aturan perilaku (acceptance criteria):**
  1. Time zone schedule (IANA, mis. `Asia/Jakarta`) menentukan interpretasi handoff time dan restriction semua rotation di dalamnya.
  2. Time zone profil pengguna dipakai untuk menampilkan waktu di UI dan notifikasi.
  3. ⚠ Timeline schedule dapat dialihkan antara time zone schedule dan time zone pengguna lewat toggle; pilihan diingat per pengguna.
  4. Pergantian DST ditangani tanpa gap atau tumpang tindih: shift yang melewati perubahan DST tetap berakhir pada jam dinding handoff, sehingga durasinya bisa 23 atau 25 jam.
  5. Time restriction pada routing rule, policy, notification rule, dan forwarding rule masing-masing menyimpan time zone sendiri (default time zone team/pengguna).
  6. Ekspor ICS menuliskan waktu dalam UTC atau dengan `TZID` yang benar sehingga kalender klien menampilkan waktu lokal yang tepat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Preferensi time zone pengguna di [02](02-pengguna-grup-organisasi.md)/[21](21-navigasi-ux-preferensi.md); server menyimpan semua waktu dalam UTC.

### JSM-OPS-024 On-call Calendar Export (ICS)

- **Deskripsi:** Ekspor jadwal on-call ke kalender eksternal melalui file atau URL iCalendar.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Responder, Team admin.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat mengunduh file `.ics` atau menyalin URL `webcal://` untuk subscribe.
  2. Ekspor dapat mencakup seluruh schedule atau hanya shift milik pengguna (lintas semua schedule).
  3. Feed dapat disinkronkan berkala oleh Google Calendar, Outlook, dan Apple Calendar; feed mencakup minimal 90 hari ke depan.
  4. ⚠ URL feed berbasis token rahasia tanpa login; pengguna dapat meregenerasi token sehingga URL lama langsung tidak berlaku (HTTP 404).
  5. Override ikut tercermin dalam feed sebagai pengganti shift rotation.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Time zone pada ICS di JSM-OPS-023.

### JSM-OPS-025 Forwarding Rules

- **Deskripsi:** Aturan meneruskan semua notifikasi dari satu pengguna ke pengguna lain selama periode tertentu, mis. saat cuti.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Responder (untuk dirinya), Ops admin (untuk pengguna lain).
- **Aturan perilaku (acceptance criteria):**
  1. Forwarding rule memiliki field from user, to user, start, dan end; from dan to tidak boleh sama, dan rantai forwarding melingkar ditolak.
  2. Periode forwarding diinterpretasikan dalam time zone yang dipilih pada rule.
  3. Pengguna dapat membuat forwarding untuk dirinya; Ops admin dapat membuat untuk pengguna lain.
  4. Daftar forwarding menampilkan rule aktif dan terjadwal, dengan aksi edit dan hapus.
  5. Notifikasi yang diteruskan tercatat di log alert ("Diteruskan dari A ke B oleh forwarding rule").
  6. API forwarding rule mendukung create, list, get, update, dan delete.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Izin di JSM-OPS-002.

## Notifikasi personal & kanal responder

### JSM-OPS-026 Notification Rules (Personal) & Quiet Hours

- **Alias inventaris:** Personal notification rules & contact methods
- **Deskripsi:** Aturan per pengguna yang menentukan kapan dan lewat contact method apa pengguna dinotifikasi untuk setiap event alert dan on-call, termasuk pembatasan jam (quiet hours) dan notifikasi berbasis role team.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder (rule miliknya), Team admin (role-based notification).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tipe event rule: New alert, Acknowledged alert, Closed alert, Assigned alert, Add note, Renotified alert, Schedule start, Schedule end, dan Incoming call routing.
  2. Rule dapat diberi kondisi pada field alert (priority, tags, message, teams, dll.) dengan komponen kondisi bersama.
  3. Time restriction per rule memberi efek quiet hours (mis. rule P3-P5 hanya aktif 08:00-20:00).
  4. Rule terdiri dari langkah: kirim via contact method tertentu setelah X menit (mis. push 0 menit, SMS 5 menit, voice 10 menit).
  5. Repeat notifikasi tiap X menit hingga alert di-acknowledge atau di-close (opsional per rule).
  6. Rule Schedule start/end mengirim pengingat X menit sebelum shift dimulai/berakhir.
  7. Pengguna baru mendapat rule default otomatis (New alert: email + push 0 menit; Schedule start: email 15 menit sebelumnya).
  8. Rule dapat di-enable/disable dan diurutkan; rule pertama yang cocok per tipe event yang dipakai.
  9. Alert yang snoozed tidak memicu notifikasi sampai snooze berakhir.
  10. ⚠ Pengaturan quiet hours/Do Not Disturb terpisah di web tidak dibangun; quiet hours dicapai lewat time restriction, sedangkan aplikasi mobile menyediakan pengaturan notifikasi perangkat.
  11. ⚠ Team admin dapat mendefinisikan role-based notification per role team (admin, member, custom role) sehingga anggota menerima notifikasi tanpa menyiapkan rule personal.
  12. ⚠ Bila pengguna memiliki rule personal yang cocok, rule personal didahulukan atas role-based notification; role-based hanya berlaku bila tidak ada rule personal, dengan kanal email, SMS, voice, dan push.
  13. Contact method yang dapat dipakai rule adalah email, SMS, voice call, dan mobile push (JSM-OPS-027).
  14. Nomor telepon harus terverifikasi sebelum dapat dipakai langkah SMS/voice; langkah dengan nomor belum terverifikasi ditandai peringatan di halaman rule.
  15. Rule tersedia untuk alert baru, alert di-assign, on-call start/end, dan perubahan schedule (pengguna ditambah/dihapus dari rotation).
  16. Setiap rule mendukung delay per langkah dan pengulangan notifikasi.
  17. Quiet hours dapat memberi pengecualian per priority (mis. P1 tetap dikirim di luar jam aktif).
  18. Critical alert di mobile dapat menembus mode senyap perangkat bila diizinkan pengguna (JSM-OPS-029).
  19. Pengguna dapat mengirim tes notifikasi dari halaman rule ke setiap contact method.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Notifikasi Operations terpisah dari notification scheme work item ([19](19-notifikasi-email.md)) tetapi memakai infrastruktur pengiriman email/push yang sama.

### JSM-OPS-027 Contact Methods & Verification

- **Deskripsi:** Saluran notifikasi milik pengguna beserta verifikasi kepemilikannya.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder (miliknya), Ops admin (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Email akun pengguna menjadi contact method default yang tidak dapat dihapus; pengguna dapat menambah email lain setelah verifikasi tautan.
  2. SMS memakai nomor dalam format E.164 dengan kode negara.
  3. ⚠ Voice call membacakan message dan priority alert, dengan input keypad untuk acknowledge/close (JSM-OPS-028).
  4. Mobile push terdaftar otomatis saat pengguna login di aplikasi mobile; setiap perangkat tampil sebagai contact method terpisah.
  5. Nomor telepon wajib diverifikasi dengan kode 6 digit (berlaku 10 menit, maksimal 5 percobaan); nomor yang belum terverifikasi tidak menerima SMS/voice.
  6. Setiap contact method dapat di-enable/disable tanpa dihapus.
  7. Daftar negara yang didukung untuk SMS/voice ditampilkan; nomor dari negara yang dibatasi ditolak dengan pesan "Negara tidak didukung".
  8. ⚠ Bila SMS/voice gagal atau tidak tersedia, sistem otomatis mengirim lewat push lalu email dan mencatat fallback di notification log.
  9. Tombol "Send test notification" tersedia per contact method.
  10. ⚠ Slack/Microsoft Teams bukan contact method personal; notifikasi chat dikirim lewat ChatOps (JSM-OPS-037) ke channel atau DM.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia SMS/voice adalah layanan pihak ketiga yang dipilih saat implementasi (lihat Di luar cakupan); kuota di JSM-OPS-004.

### JSM-OPS-028 Notification Action Responses

- **Deskripsi:** Kemampuan responder bertindak langsung dari notifikasi push, email, SMS, atau voice tanpa membuka halaman alert.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Responder yang menerima notifikasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Push notification menampilkan tombol Acknowledge dan Close (serta Take ownership bila tersedia).
  2. Email notifikasi memuat tautan aksi Acknowledge, Close, dan View yang bertanda tangan, berlaku 24 jam, dan hanya untuk penerima tersebut.
  3. ⚠ Balasan SMS dengan kode (mis. `1` acknowledge, `2` close) diproses untuk negara yang mendukung SMS dua arah.
  4. ⚠ Voice call menerima input keypad untuk acknowledge (1), close (2), dan escalate to next (3).
  5. Setiap aksi dari notifikasi tercatat di activity log dengan channel asal (push/email/SMS/voice).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aksi yang diproses sama dengan JSM-OPS-010.

### JSM-OPS-029 Mobile App for On-call

- **Deskripsi:** Fungsi Operations di aplikasi mobile untuk menerima push, menangani alert, dan melihat jadwal on-call.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Responder berlisensi agent.
- **Aturan perilaku (acceptance criteria):**
  1. Push notification alert memiliki tombol aksi acknowledge dan close.
  2. Pengguna dapat mengizinkan Critical alerts di iOS dan bypass Do Not Disturb di Android untuk priority tertentu.
  3. Aplikasi menampilkan daftar alert, detail alert, add note, snooze, escalate, dan execute custom action.
  4. Aplikasi menampilkan schedule dan who is on-call, serta tombol take on-call (JSM-OPS-021).
  5. Pengguna dapat mengelola notification rule dan contact method dari aplikasi.
  6. Perangkat terdaftar otomatis sebagai contact method push saat login dan dihapus saat logout.
  7. ⚠ Pengguna dapat memilih suara notifikasi berbeda per priority.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Strategi aplikasi mobile produk ini di NAV-050 ([21](21-navigasi-ux-preferensi.md)) dan push di NTF-038 ([19](19-notifikasi-email.md)). Aplikasi native belum ada di repo; alternatif awal adalah PWA dengan Web Push.

### JSM-OPS-030 Incoming Call Routing

- **Alias inventaris:** Incoming call routing
- **Deskripsi:** Nomor telepon yang meneruskan panggilan masuk ke responder on-call dan membuat alert.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (mengaktifkan nomor), Team admin, Caller, Responder.
- **Aturan perilaku (acceptance criteria):**
  1. Ops admin dapat mengaktifkan nomor telepon per negara dari daftar negara yang didukung penyedia.
  2. Panggilan diteruskan ke on-call team/escalation yang dipilih, mengikuti urutan langkah escalation.
  3. Bila tidak diangkat, penelepon diarahkan ke voicemail dan alert dibuat berisi rekaman/transkrip voicemail.
  4. Greeting kustom (teks-ke-suara atau rekaman) dapat diatur per nomor.
  5. Notification rule tipe Incoming call routing menentukan cara responder dinotifikasi saat ada panggilan.
  6. Setiap nomor ditetapkan ke satu team pemilik.
  7. Routing panggilan dapat menuju schedule atau escalation team.
  8. Voicemail yang tidak dijawab menghasilkan alert dengan priority yang dapat diatur.
  9. Greeting dapat berbeda per nomor dan per bahasa.
  10. Log panggilan mencatat nomor penelepon, waktu, durasi, penerima yang menjawab, dan alert terkait.
  11. ⚠ Fitur hanya tersedia pada plan Premium dan Enterprise.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung penyedia telepon pihak ketiga; plan di JSM-OPS-004.

## Integrasi alert

### JSM-OPS-031 Integrations Framework

- **Alias inventaris:** Alert integrations
- **Deskripsi:** Kerangka umum integrasi incoming, outgoing, dan bidirectional untuk menerima dan mengirim event alert, termasuk integration logs untuk troubleshooting.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (integrasi global), Team admin (integrasi team), Integration.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Integrations menampilkan katalog integrasi dengan pencarian dan kategori (Monitoring, ITSM, ChatOps, Generic, Ticketing); katalog awal produk ini berisi integrasi generik dan sejumlah integrasi monitoring prioritas (JSM-OPS-034).
  2. Integrasi dibuat di level team (dimiliki team, responder default team) atau level global (Ops admin).
  3. Setiap integrasi memiliki API key unik yang dapat diregenerasi; key lama langsung tidak berlaku setelah regenerasi.
  4. Integrasi dapat di-enable/disable; request ke integrasi disabled ditolak dan dicatat di log.
  5. Integrasi dapat menetapkan responders default dan opsi suppress notifications (alert dibuat tanpa notifikasi).
  6. ⚠ Pengaturan aksi incoming memuat Create, Close, Acknowledge, Add note, dan Ignore, masing-masing dengan filter kondisi; rule dievaluasi berurutan.
  7. Field mapping mendukung dynamic field (`{{payload.path}}`) dan ekstraksi regex dari payload untuk setiap field alert.
  8. Opsi "Send alert updates back" mengirim aksi alert kembali ke tool sumber untuk integrasi bidirectional.
  9. Konfigurasi integrasi team hanya dapat diubah Team admin pemilik dan Ops admin.
  10. Integration logs mencatat payload request masuk, hasil parsing/aksi yang diambil, error mapping, serta request/respons outgoing (webhook dan integrasi bidirectional).
  11. ⚠ Integration logs dapat difilter berdasarkan rentang waktu, integrasi, team, dan level (info/error).
  12. Akses integration logs dibatasi untuk Ops admin dan Team admin pemilik integrasi.
  13. ⚠ Retensi integration logs mengikuti plan (default 30 hari, dapat dikonfigurasi).
  14. ⚠ Jumlah integrasi per situs dibatasi sesuai plan (JSM-OPS-004).
  15. Library integrasi mencakup tool monitoring umum (Datadog, AWS CloudWatch, Prometheus/Alertmanager, Grafana, New Relic, Azure Monitor, Splunk, Zabbix, Nagios, Pingdom, dsb.) sesuai JSM-OPS-034.
  16. Integrasi berbasis email tersedia sebagai tipe integrasi (JSM-OPS-033).
  17. Integrasi API (REST) dan webhook tersedia sebagai tipe integrasi generik (JSM-OPS-032, JSM-OPS-035).
  18. Payload dipetakan ke field alert dan diproses hanya bila lolos filter kondisi integrasi.
  19. Outgoing action dikirim ke tool eksternal saat alert berubah sesuai aksi yang dilanggan.
  20. Integrasi dapat dimiliki per team atau global, dengan perbedaan responder default.
  21. Integrasi dapat di-enable/disable dan setiap integrasi memiliki log sendiri.
  22. Sinkronisasi status dua arah dengan tool sumber didukung pada integrasi bertipe bidirectional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi webhook dan keamanan payload di API-066 ([23](23-api-publik-webhook.md)); tipe integrasi diimplementasikan sebagai adapter plug-in di `internal/features/operations/integration`.

### JSM-OPS-032 API Integration

- **Deskripsi:** Integrasi generik yang menyediakan API key untuk membuat dan mengelola alert dari sistem apa pun.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, Team admin, API client.
- **Aturan perilaku (acceptance criteria):**
  1. Membuat API integration menghasilkan API key integrasi yang ditampilkan penuh sekali dan dapat disalin.
  2. ⚠ Hak key dapat dipilih: read, create and update, delete, dan configuration access; request di luar hak mengembalikan HTTP 403.
  3. API integration dapat di-assign ke team untuk routing default.
  4. ⚠ Endpoint integrasi berada di `/api/v1/ops/...` pada host produk ini; deployment multi-region memakai host per region.
  5. API key integrasi juga dapat dipakai untuk heartbeat ping (JSM-OPS-016).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alert API di JSM-OPS-045; key disimpan dalam bentuk hash.

### JSM-OPS-033 Email Integration

- **Deskripsi:** Membuat dan mengelola alert dari email masuk ke alamat khusus integrasi dengan aturan parsing.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, Team admin, External system (pengirim email).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap email integration memiliki alamat unik pada domain situs (mis. `<nama>@<situs>.ops.<domain>`), yang dapat diubah prefix-nya selama unik.
  2. Default mapping: subject → message, body → description, from → source; mapping dapat diubah per field.
  3. Rule Create/Close/Acknowledge/Add note dipilih berdasarkan kondisi pada subject, body, atau from (termasuk regex).
  4. Alias dapat diambil dari subject atau hasil regex untuk deduplikasi.
  5. Filter dapat mengabaikan email tertentu (mis. auto-reply, dari pengirim tertentu).
  6. Email HTML dikonversi ke teks biasa untuk description; bila ada bagian plain text, bagian itu dipakai.
  7. ⚠ Email melebihi batas ukuran (default 1 MB tanpa attachment) ditolak dan dicatat di integration logs.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Infrastruktur inbound mail berbagi dengan kanal email Service Management ([30](30-jsm-help-center-portal-pelanggan-kanal.md)) dan NTF ([19](19-notifikasi-email.md)).

### JSM-OPS-034 Monitoring Tool Integrations

- **Deskripsi:** Integrasi native dengan tool monitoring/observability untuk membuat dan menutup alert otomatis berdasarkan payload webhook masing-masing tool.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, Team admin, Integration.
- **Aturan perilaku (acceptance criteria):**
  1. Datadog didukung secara bidirectional (alert masuk, ack/close dikirim balik).
  2. New Relic didukung dengan pemetaan incident open/close.
  3. Prometheus/Alertmanager didukung dengan alias dari `fingerprint`/`groupKey` dan status `firing`/`resolved`.
  4. Grafana alerting didukung dengan status `alerting`/`ok`.
  5. Amazon CloudWatch via Amazon SNS dan layanan AWS lain didukung, termasuk konfirmasi subscription SNS otomatis.
  6. Nagios dan Zabbix didukung, termasuk lewat edge connector untuk instalasi on-prem (JSM-OPS-036).
  7. Splunk didukung dengan pemetaan saved search alert.
  8. Sentry didukung dengan pemetaan issue alert.
  9. Pingdom didukung dengan status down/up.
  10. ⚠ Integrasi lain (Dynatrace, AppDynamics, Azure Monitor, Google Cloud Monitoring, Elastic) masuk katalog bertahap; tool tanpa integrasi khusus memakai API integration atau incoming webhook generik.
  11. Status recovery/resolved dari tool menutup alert terkait secara otomatis berdasarkan alias.
  12. Severity tool dipetakan ke priority P1-P5 sesuai tabel mapping integrasi.
  13. Ack/close di produk ini disinkronkan balik ke tool bila tool mendukung API balik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Daftar integrasi rilis pertama adalah keputusan terbuka; setiap integrasi dibangun sebagai template field mapping di atas JSM-OPS-031.

### JSM-OPS-035 Outgoing Webhook Integration

- **Deskripsi:** Mengirim event alert ke endpoint HTTP eksternal saat aksi tertentu terjadi.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, Team admin, External system (penerima).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi memuat URL tujuan (HTTPS wajib kecuali mode pengembangan) dan custom header.
  2. ⚠ Aksi pemicu dapat dipilih: create, acknowledge, unacknowledge, close, add note, assign ownership, add responder, add/remove tags, escalate, snooze, dan custom action.
  3. Setiap aksi pemicu dapat diberi filter kondisi.
  4. Opsi menyertakan alert description dan alert details di payload dapat diaktifkan/dimatikan.
  5. Payload JSON berisi objek `action`, `alert`, dan `source`, serta `integrationId` dan timestamp.
  6. Status respons atau kegagalan pengiriman terlihat di integration logs beserta kode HTTP dan potongan body.
  7. ⚠ Pengiriman gagal (5xx/timeout) diulang hingga 3 kali dengan backoff eksponensial; 4xx tidak diulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme pengiriman berbagi dengan webhook umum API-066 ([23](23-api-publik-webhook.md)).

### JSM-OPS-036 Jira Edge Connector (JEC)

- **Deskripsi:** Agen on-premise (padanan "edge connector" produk ini) yang menjembatani sistem lokal dengan Operations dan menjalankan aksi/skrip lokal sebagai respons alert.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, System admin pelanggan.
- **Aturan perilaku (acceptance criteria):**
  1. Agen disediakan sebagai binary untuk Linux dan Windows yang dapat dipasang sebagai service.
  2. Agen melakukan polling antrean aksi dari cloud via HTTPS keluar sehingga tidak perlu membuka port masuk.
  3. Agen menjalankan skrip lokal saat aksi alert terjadi, termasuk pemetaan custom action ke skrip tertentu.
  4. ⚠ Hasil eksekusi skrip (exit code, stdout terpotong) dikirim balik sebagai note alert dan integration log.
  5. Konfigurasi agen memakai file config lokal dan API key integrasi.
  6. ⚠ Skrip dapat ditulis dalam Go, Python, Groovy, atau shell (agen hanya mengeksekusi perintah yang dikonfigurasi).
  7. Agen berfungsi sebagai bridge untuk Nagios, Zabbix, dan tool on-prem lain yang tidak dapat dijangkau dari internet.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** P3; dapat diganti sementara oleh outgoing webhook ke layanan internal pelanggan.

### JSM-OPS-037 ChatOps (Slack & Microsoft Teams)

- **Deskripsi:** Integrasi chat untuk menerima notifikasi alert di channel dan melakukan aksi alert langsung dari chat.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (menghubungkan workspace), Team admin (channel team), Responder.
- **Aturan perilaku (acceptance criteria):**
  1. Ops admin dapat menghubungkan workspace Slack atau tenant Microsoft Teams melalui OAuth.
  2. Notifikasi alert dikirim ke channel terpilih dengan filter kondisi per channel.
  3. ⚠ Pesan alert memuat tombol aksi Acknowledge, Close, Take ownership, Add note, dan Snooze.
  4. ⚠ Perintah chat (mis. `/ops oncall <team>`, `/ops alert ack <tinyId>`) menampilkan on-call dan mengelola alert.
  5. Akun chat dipetakan ke pengguna produk ini lewat alur connect akun; aksi dari akun yang belum terhubung ditolak dengan tautan connect.
  6. Pesan alert diperbarui (status, owner) saat state alert berubah.
  7. Pembuatan channel incident dari incident ditangani modul ITSM (32).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aplikasi Slack/Teams dasar di INT-030 s.d. INT-036 ([25](25-integrasi-impor-ekspor-ekstensi.md)); notifikasi personal chat di NTF-039/NTF-040 ([19](19-notifikasi-email.md)).

### JSM-OPS-038 Jira Issue Sync & Ticketing Integrations

- **Deskripsi:** Sinkronisasi alert dengan work item produk ini dan sistem tiket lain secara satu atau dua arah.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, Agent Layanan, Automation.
- **Aturan perilaku (acceptance criteria):**
  1. Integrasi dapat membuat work item dari alert dengan konfigurasi project, work type, dan mapping field (message → Summary, description → Description, priority → Priority).
  2. Sinkron dua arah: alert ditutup saat work item mencapai status category Done, dan work item ditransisikan ke status yang dikonfigurasi saat alert closed.
  3. ⚠ Note alert menjadi comment work item dan sebaliknya, dengan penanda asal untuk mencegah loop.
  4. Alert menampilkan tautan ke work item dan work item menampilkan panel alert tertaut.
  5. ⚠ Integrasi ke sistem tiket eksternal (ServiceNow, Zendesk, Freshservice) tersedia lewat outgoing webhook atau konektor katalog.
  6. ⚠ Otomasi menyediakan trigger "Alert created/updated" dan action "Create alert"/"Close alert".
- **Team-managed vs company-managed:** Project tujuan dapat team-managed maupun company-managed; mapping field mengikuti konfigurasi field project masing-masing.
- **Catatan & ketergantungan:** Work item inti di [06](06-work-item-inti.md); otomasi di [20](20-otomasi.md); trigger alert dan action alert otomasi Service Management dirinci di JSM-AI-016 dan JSM-AI-019 ([35](35-jsm-ai-otomasi-api-admin-csm.md)).

## Incident, service & AIOps

### JSM-OPS-039 Alert to Incident Linkage & Post-incident Review

- **Deskripsi:** Keterkaitan alert dengan incident Service Management, termasuk pembuatan incident dari alert dan penautan ke post-incident review (PIR).
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Responder, Incident manager (Manajer Layanan/ITSM), Automation; membuat incident memerlukan izin create di service project tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Create incident dari alert membuka dialog dengan summary (dari message), priority (P1-P5 dipetakan ke priority incident), dan affected services terisi otomatis serta pilihan service project.
  2. Alert dapat ditautkan ke incident yang ada; satu incident dapat memiliki banyak alert, satu alert hanya terhubung ke satu incident.
  3. Incident menampilkan daftar alert tertaut (status, priority, count), dan alert menampilkan info incident (key, status).
  4. Alert dapat di-unlink dari incident; unlink dicatat di activity log keduanya.
  5. ⚠ Incident dapat dibuat otomatis dari alert lewat alert policy atau otomasi berdasarkan kondisi (mis. P1 dengan tag `production`).
  6. ⚠ Menutup incident tidak otomatis menutup alert tertaut; pengguna ditawari aksi "Close linked alerts".
  7. PIR dibuat dari incident dan menampilkan tautan ke alert terkait; detail PIR dirinci di modul 32.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas incident dan PIR di modul 32; services di JSM-OPS-041.

### JSM-OPS-040 Incident Responders

- **Deskripsi:** Mekanisme alerting di balik responder incident: saat user atau team ditambahkan sebagai responder, modul ini membuat alert terkait incident, merutekannya ke schedule/escalation team, dan menyinkronkan status acknowledge kembali ke incident. UI panel Responders pada incident (tambah/hapus, tampilan status, responder dari affected service, action Automation) dimiliki oleh JSM-ITSM-019 di [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md); fitur ini hanya memiliki mekanisme alert-nya.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan/incident manager (menambah responder lewat JSM-ITSM-019), Responder (menerima dan meng-acknowledge alert), Team admin (mengelola routing rule, schedule, dan escalation team yang dipakai).
- **Aturan perilaku (acceptance criteria):**
  1. Field Responders pada incident (UI di JSM-ITSM-019) menerima pengguna dan team; hanya team yang terdaftar sebagai team ops (JSM-OPS-001) yang dapat dipilih, team lain ditolak dengan pesan "Team belum dikonfigurasi untuk Operations".
  2. Menambahkan responder membuat alert terkait incident (message dari summary, alias berbasis key incident) untuk menotifikasi responder tersebut; responder user dinotifikasi sesuai notification rule miliknya (JSM-OPS-026).
  3. Alert untuk team responder dievaluasi oleh routing rule team (JSM-OPS-017) dan diteruskan ke schedule on-call atau escalation yang menjadi target rule tersebut (JSM-OPS-018, JSM-OPS-019). ⚠ Bila team tidak memiliki on-call schedule maupun target routing yang valid, alert tetap tercatat tetapi tidak ada yang dinotifikasi; sistem mengembalikan peringatan ke panel Responders (JSM-ITSM-019) dan tidak melakukan fallback ke seluruh anggota team.
  4. Status per responder diturunkan dari state alert terkait dengan label tetap Not acknowledged (alert open belum di-ack), Acknowledged (alert di-ack), dan Closed (alert ditutup); status ini diperbarui otomatis setiap kali state alert berubah dan ditampilkan di panel Responders incident (JSM-ITSM-019).
  5. Menghapus responder dari incident tidak menutup alert yang sudah dibuat; alert tetap mengikuti siklus hidupnya sendiri (JSM-OPS-005).
  6. ⚠ Saran responder berbasis AI tersedia pada plan Premium (JSM-OPS-042).
  7. Conference call dari incident dirinci di modul ITSM (32).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan: JSM-ITSM-019 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)) adalah pemilik UI responder pada incident dan aturan validasinya; JSM-OPS-040 adalah pemilik pembuatan alert, routing ke schedule/escalation, dan sinkronisasi status. Prioritas (P0) dan label status (Not acknowledged/Acknowledged/Closed) wajib sama di kedua fitur. Bergantung entitas incident (modul 32) dan keterkaitan alert-incident (JSM-OPS-039).

### JSM-OPS-041 Service-aware Alerting

- **Deskripsi:** Mengaitkan alert dan incident dengan services di service registry untuk konteks dampak.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Responder, Service owner.
- **Aturan perilaku (acceptance criteria):**
  1. Alert dan incident memiliki field affected services (multi-pilih dari service registry); integrasi dapat mengisinya lewat mapping.
  2. Halaman service menampilkan alert dan incident terbuka untuk service tersebut.
  3. Service tier ditampilkan di alert untuk menilai dampak dan dapat dipakai sebagai kondisi policy/routing.
  4. ⚠ Owner team service dapat ditambahkan otomatis sebagai responder saat alert/incident menyebut service tersebut.
  5. ⚠ Status kesehatan service (Operational/Degraded/Down) diturunkan dari alert/incident terbuka berprioritas tinggi.
  6. Dependensi service dan objek Assets ditampilkan sebagai konteks dampak hulu/hilir.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Service registry di modul 32; Assets di modul 33.

### JSM-OPS-042 AIOps Alert Grouping

- **Deskripsi:** Kemampuan berbasis AI untuk mengelompokkan alert terkait, membedakan signal dan noise, dan membantu triase.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Responder, Incident manager, Team admin (memberi feedback dan mengaktifkan).
- **Aturan perilaku (acceptance criteria):**
  1. Alert serupa (kemiripan message, entity, source, waktu) dikelompokkan otomatis menjadi alert group.
  2. Halaman alert group menampilkan ringkasan group dan daftar alert anggota.
  3. Detail alert menampilkan related alert groups.
  4. ⚠ Alert group tersedia di aplikasi mobile.
  5. Incident dapat dibuat dari alert group dan seluruh alert anggota ditautkan.
  6. ⚠ Alert diklasifikasikan sebagai signal atau noise, dengan filter/tampilan alerts by signal or noise di daftar alert.
  7. Panel similar alerts dan past responders tampil di detail alert.
  8. ⚠ Related resources (runbook, work item, dokumen) disarankan di detail alert.
  9. Ringkasan AI alert/group dan saran responder tersedia.
  10. Pengguna dapat memberi feedback akurasi grouping (benar/salah) yang dipakai untuk penyetelan.
  11. Seluruh fitur hanya tersedia bila kapabilitas AI situs diaktifkan admin; bila mati, panel AI disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kapabilitas AI umum di [35](35-jsm-ai-otomasi-api-admin-csm.md) dan ISS-045 ([06](06-work-item-inti.md)); grouping awal dapat berbasis heuristik non-AI sebelum model tersedia. Kapabilitas AIOps incident (ringkasan, suggested responders, root cause) dirinci di JSM-AI-045 ([35](35-jsm-ai-otomasi-api-admin-csm.md)); fitur ini memegang spesifikasi grouping alert.

## Visibilitas, laporan, API & IaC

### JSM-OPS-043 Team Dashboard

- **Deskripsi:** Halaman ringkasan per team untuk kondisi operasional saat ini.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Team member, Team admin.
- **Aturan perilaku (acceptance criteria):**
  1. Widget who is on-call menampilkan on-call sekarang dan berikutnya untuk setiap schedule team.
  2. Daftar alert open/unacked team ditampilkan dengan tautan ke daftar alert terfilter.
  3. Incident aktif yang melibatkan team ditampilkan.
  4. ⚠ Statistik ringkas alert (jumlah 7 hari, MTTA, MTTR) ditampilkan.
  5. Tautan cepat ke schedule, escalation, dan integrasi team tersedia.
  6. Status heartbeat (expired disorot) dan maintenance yang sedang aktif ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Gadget dashboard umum di [17](17-dashboard-gadget.md) dapat menampilkan widget yang sama.

### JSM-OPS-044 Operations Reports & Analytics

- **Deskripsi:** Laporan metrik alert, notifikasi, dan on-call per team dan global.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Ops admin (global), Team admin (team), Manager.
- **Aturan perilaku (acceptance criteria):**
  1. Alert analytics menampilkan volume alert per waktu, per priority, per integrasi/source, dan per tag.
  2. MTTA (rata-rata `ackTime`) dan MTTR (rata-rata `closeTime`) dihitung per team dan priority; alert tanpa ack dikecualikan dari MTTA.
  3. Notification analytics menampilkan jumlah notifikasi per pengguna per contact method, termasuk di luar jam kerja.
  4. On-call analytics menampilkan jam on-call per pengguna, termasuk jam di luar jam kerja team.
  5. Semua laporan dapat difilter rentang waktu, team, dan priority.
  6. ⚠ Data laporan dapat diekspor ke CSV.
  7. ⚠ Laporan global lintas team tersedia pada plan Premium.
  8. Data Operations tersedia di Analytics/Data Lake pada plan Enterprise.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Analytics/Data Lake di RPT-038 s.d. RPT-041 ([16](16-laporan-insight.md)); laporan JSM umum di [31](31-jsm-antrian-sla-knowledge-base-laporan.md).

### JSM-OPS-045 Operations REST API

- **Deskripsi:** API publik untuk alert, on-call, dan konfigurasi Operations.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Integrator/Developer API, API client.
- **Aturan perilaku (acceptance criteria):**
  1. Alert API (autentikasi API key integrasi) mendukung create, close, acknowledge, unacknowledge, snooze, escalate, assign, add note, tags/details/responders, list, get, count, request status, dan attachment.
  2. API konfigurasi berada di `/api/v1/ops/...` dengan autentikasi basic auth email + API token pengguna (AUTH-012) atau Bearer token sesi, dan izin mengikuti JSM-OPS-002.
  3. Endpoint schedules, rotations, overrides, on-calls, dan timeline tersedia.
  4. Endpoint escalations, routing rules, dan team operations tersedia.
  5. Endpoint heartbeats, maintenance, integrations, dan policies tersedia.
  6. Endpoint contacts, notification rules, dan forwarding rules tersedia.
  7. Endpoint list mendukung pagination `offset`/`limit` (maksimal 100) dan sorting.
  8. Melebihi rate limit mengembalikan HTTP 429 dengan header `Retry-After`.
  9. ⚠ Global API key (key tingkat situs dengan hak konfigurasi) tidak disediakan; akses konfigurasi selalu melalui API token pengguna atau key integrasi berhak configuration access.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi API publik di [23](23-api-publik-webhook.md); dokumentasi Swagger (`make swag`) wajib untuk semua endpoint.

### JSM-OPS-046 Atlassian Operations Terraform Provider

- **Deskripsi:** Provider Terraform untuk mengelola konfigurasi Operations produk ini sebagai infrastructure-as-code di atas REST API konfigurasi; untuk clone, fokusnya memastikan API konfigurasi mendukung CRUD idempoten.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Ops admin, DevOps/Platform engineer.
- **Aturan perilaku (acceptance criteria):**
  1. Provider dikonfigurasi dengan URL situs, email + API token pengguna, dan id situs.
  2. ⚠ Resource tersedia untuk teams, schedules, rotations, escalation policies, routing rules, integrations, alert/notification policies, contacts, dan notification rules.
  3. ⚠ Data source tersedia untuk membaca resource yang sudah ada (mis. team berdasarkan nama).
  4. Dokumentasi pemetaan resource dari provider tool alerting lama ke resource provider produk ini disediakan.
  5. Perubahan dari UI yang menyebabkan drift terdeteksi pada `terraform plan`; resource default team (routing rule default) tidak dapat dihapus dan hanya dapat diimpor/diubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung JSM-OPS-045; provider dipublikasikan terpisah dari binary utama.

### JSM-OPS-047 Opsgenie Migration (Legacy)

- **Deskripsi:** Migrasi dari produk alerting lama ekosistem rujukan ke Operations sesuai jadwal sunset-nya. Tidak dibangun di clone; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Ops admin, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan: produk lama berstatus end of sale 4 Juni 2025 dan end of support 5 April 2027; clone tidak memiliki produk lama sehingga tidak ada jadwal sunset.
  2. Rujukan: tujuan migrasi di ekosistem rujukan adalah Service Management atau produk katalog service; clone hanya menyediakan Operations di Service Management.
  3. ⚠ Data yang dipindah di ekosistem rujukan (users, teams, schedules, escalations, routing, integrations, policies, heartbeats) di clone dapat diimpor lewat REST API konfigurasi (JSM-OPS-045) tanpa wizard khusus.
  4. Perubahan endpoint API tidak relevan; clone hanya memiliki `/api/v1/ops/...`.
  5. Migrasi konfigurasi Terraform cukup didukung oleh dokumentasi pemetaan di JSM-OPS-046.
  6. Pemetaan peran lama ke peran Operations dicakup oleh model peran JSM-OPS-002.
  7. ⚠ Daftar fitur produk lama yang tidak dibawa/berganti nama tidak dipelihara di clone.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## Pertanyaan terbuka

1. Skema header autentikasi Alert API apa yang dipakai (mis. `Authorization: OpsKey <key>`), dan apakah perlu kompatibilitas path/header dengan API alerting rujukan agar tool monitoring yang sudah ada dapat dipakai tanpa modifikasi (JSM-OPS-007)?
2. Apakah team ops wajib memakai entitas Teams umum (PLAN-057), atau boleh ada team khusus Operations (JSM-OPS-001)?
3. Penyedia SMS/voice/telepon mana yang dipakai, di negara mana saja, dan berapa kuota per plan (JSM-OPS-004, JSM-OPS-027, JSM-OPS-030)?
4. Apakah aplikasi mobile native dibangun, atau Fase 7 cukup dengan PWA + Web Push (dengan konsekuensi tidak ada Critical alerts iOS) (JSM-OPS-029)?
5. Integrasi monitoring mana yang masuk rilis pertama selain API, email, dan webhook generik (JSM-OPS-034)?
6. Apakah alert policy dan notification policy ditampilkan sebagai satu daftar Policies atau dua daftar terpisah, dan apakah global policy dibatasi plan Premium (JSM-OPS-013, JSM-OPS-014)?
7. Apakah Operations boleh diaktifkan tanpa service project/incident (alerting mandiri), atau selalu terikat ke Service Management (JSM-OPS-039)?
8. Apakah AIOps grouping dimulai dengan heuristik non-AI, dan model/penyedia AI apa yang dipakai bila kapabilitas AI aktif (JSM-OPS-042)?
9. Berapa nilai default retensi alert dan integration logs per plan (JSM-OPS-004, JSM-OPS-031)?
10. Apakah edge connector (JSM-OPS-036) dan provider Terraform (JSM-OPS-046) dibangun oleh tim inti atau diserahkan ke komunitas?

## Di luar cakupan & ekosistem

- **Produk alerting standalone lama ekosistem rujukan** (Opsgenie) beserta wizard migrasinya (JSM-OPS-047, tier dc-only-or-deprecated) tidak dibangun; impor konfigurasi dilakukan lewat REST API.
- **Produk katalog service/developer portal terpisah** (Compass) sebagai tujuan migrasi alternatif tidak dibangun.
- **App Marketplace pihak ketiga** untuk on-call/alerting (mis. pager dan status page eksternal) tidak dibangun sebagai fitur native; interoperabilitas lewat outgoing webhook dan API integration.
- **Penyedia SMS, voice, telepon, dan push** (gateway telekomunikasi, APNs/FCM) adalah layanan pihak ketiga; produk ini hanya membangun adapter dan tidak menjadi operator telekomunikasi.
- **Aplikasi Slack/Microsoft Teams** dan tool monitoring (Datadog, Prometheus, Grafana, dll.) adalah produk pihak ketiga; yang dibangun hanya konektornya.
- **Atlassian Analytics/Data Lake, Atlassian Intelligence/Rovo, dan Atlassian Teams** sebagai layanan lintas produk tidak dibangun apa adanya; produk ini memakai padanan internal (modul 16, 18, 35) atau menandainya sebagai flag plan.
- **Billing, pricing, dan pengelolaan lisensi** plan (JSM-OPS-004) tidak dibangun; tier diperlakukan sebagai feature flag konfigurasi.
- **Status page publik** dan conference call bridge berada di luar modul ini (rujukan modul 32).
