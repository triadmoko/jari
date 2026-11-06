# 35 — Service Management: Virtual Service Agent, AI, Otomasi, API, Admin & Customer Service Management

Modul ini menutup lapisan lintas-fitur Service Management produk ini: **administrasi produk** (peran dan permission default, customer permissions per service project, lisensi agent, collaborator, matriks fitur per plan, pengaturan global, batas email channel, akun dan autentikasi customer, audit dan privasi, data residency, pemantauan pemakaian fitur berkuota); **ekstensi JQL** khusus Service Management (field request, fungsi SLA dan approval); **otomasi** khusus Service Management (trigger, condition, action, smart values, template, kuota, serta rujukan mesin legacy); **REST API** khusus Service Management (service desk, request, request type, queue, customer/organization, knowledge base, forms, Assets, Operations) beserta extension point app; **produktivitas agent berbasis AI** (canned responses, ringkasan, bantuan menulis, request serupa, pencarian bahasa alami, sentiment, agen AI, Request Resolver, AIOps, Feature Lab); **Virtual Service Agent** percakapan (intent, flow builder, standard flow, AI answers, kanal, hand-off, pengujian, analitik, multibahasa, kuota); **Customer Service Management** untuk dukungan pelanggan eksternal (products, entitlements, detail customer/organization, context panel, agen AI customer-facing); serta aplikasi mobile, analitik/data lake, migrasi, dan sandbox. Semua kapabilitas AI di modul ini bersifat **opsional dan pluggable**: produk ini tidak bergantung pada layanan AI vendor tertentu, dan seluruh fitur AI wajib dapat dimatikan tanpa merusak fungsi inti. Istilah merek ekosistem rujukan yang muncul di nama fitur inventaris (mis. Atlassian Intelligence, Rovo, Opsgenie, Workato, Workday) dipakai hanya sebagai rujukan padanan; nama produk ini tidak memakai merek tersebut. Prosa modul ini memakai istilah "service project" dan "work item" seperti modul 29-31, sedangkan label UI yang tampil ke pengguna (mis. "space" pada navigasi dan pengaturan) mengikuti set terminologi VISI-002 ([00](00-visi-ruang-lingkup-glosarium-baseline.md)).

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`, JWT di `pkg/jwt`, context key `"sess"`). Belum ada service project, request type, customer, organization, queue, SLA, maupun mesin otomasi atau penyedia AI. Modul ini diusulkan menjadi beberapa feature module baru: `internal/features/servicedesk/admin` (pengaturan global, customer permissions, lisensi, usage), `internal/features/servicedesk/api` (REST API publik Service Management di bawah `/api/v1/servicedeskapi/...`, lihat konvensi sub-area REST API), `internal/features/servicedesk/aiassist` (AI untuk agent), `internal/features/virtualagent` (intent, flow, percakapan, kanal), dan `internal/features/csm` (products, entitlements, detail profil). Entitas bersama (CustomerAccount, Organization, CannedResponse, VirtualAgentConfig, Intent, ConversationFlow, Conversation, Product, Entitlement, FeatureUsage) diletakkan di `internal/shared/domain/entity`; antarmuka penyedia AI (`AIProvider`) di `pkg/` agar bisa diganti; UI di SPA SvelteKit `frontend/src/routes`. Seluruh route mengikuti konvensi `/api/v1`, dengan pesan galat dwibahasa EN/ID lewat `internal/shared/constants/error.go`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — tier plan, istilah, dan batas cakupan produk.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — login, CAPTCHA (AUTH-006), API token (AUTH-012), consent OAuth (AUTH-013).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — product access (USR-018), batas pengguna (USR-020), SAML (USR-034), data residency (USR-042), API data pribadi (USR-044), audit log (USR-045, USR-046).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — product access (PERM-004), batas permission per plan (PERM-005), izin agen AI (PERM-050).
- [06 Work Item Inti](06-work-item-inti.md) — AI pada work item (ISS-045) dan agen AI sebagai assignee (ISS-046).
- [12 JQL — Bahasa Kueri](12-jql-bahasa-kueri.md) — grammar, kompatibilitas operator (JQL-013), custom field (JQL-033), bahasa alami ke JQL (JQL-059).
- [16 Laporan & Insight](16-laporan-insight.md) — Analytics dan Data Lake (RPT-038 s.d. RPT-041).
- [19 Notifikasi & Email](19-notifikasi-email.md) — push mobile (NTF-038) dan email keluar.
- [20 Otomasi](20-otomasi.md) — mesin rule, template (AUTO-010), kuota (AUTO-015, AUTO-016), smart values (AUTO-063 s.d. AUTO-065), AI di otomasi (AUTO-073).
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — aplikasi mobile (NAV-050), chat dan agen AI (NAV-051), katalog fitur AI (NAV-052).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log (ADM-014), migrasi (ADM-039), usage & plan limits (ADM-042), data residency (ADM-048), sandbox (ADM-049), pengaturan AI (ADM-052), billing (ADM-053), koleksi lisensi (ADM-056).
- [23 REST API & Webhook](23-api-publik-webhook.md) — rate limit (API-006), API token (API-008), OAuth 2.0 (API-009, API-010), Forge/Connect (API-013, API-015), entity properties (API-025), katalog event webhook (API-065, API-066).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — iPaaS (INT-003), migrasi (INT-026), Slack/Teams (INT-030 s.d. INT-036), Statuspage (INT-052), connector pengetahuan (INT-055), extension point (INT-060, INT-068 s.d. INT-080).
- [27 NFR, Arsitektur & Platform](27-nfr-arsitektur-platform.md) — batas & guardrail (NFR-004), audit keamanan (NFR-017), privasi & GDPR (NFR-018).
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) — service project, request type, forms, tipe pengguna, permission project.
- [30 Help Center, Portal Pelanggan & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — help center, akun customer, organizations, email channel, widget, chat.
- [31 Antrian, SLA, Knowledge Base & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) — queues, SLA, canned responses dasar (JSM-QSK-019), AI assist dasar (JSM-QSK-020), KB, CSAT.
- [32 ITSM: Incident, Problem, Change & Services](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — approvals, incident, PIR, change, Statuspage.
- [33 Assets (CMDB)](33-jsm-assets-cmdb.md) — AQL, objek, import, Assets di JQL.
- [34 Operations, Alerting & On-call](34-jsm-operations-alerting-on-call.md) — alert, schedule, AIOps grouping (JSM-OPS-042), REST API konfigurasi (JSM-OPS-045).

## Persona terkait

- **Admin Situs** — mengaktifkan Service Management, mengelola lisensi agent, pengaturan global, customer accounts, autentikasi customer, audit, data residency, pengaturan AI, dan pemakaian fitur berkuota.
- **Admin Proyek/Project Lead** — pengelola service project: mengatur customer permissions, rule otomasi, Virtual Service Agent (intent, flow, kanal), Feature Lab, sentiment analysis, Request Resolver, dan canned responses bersama.
- **Agent Layanan** — memakai canned responses, ringkasan dan bantuan menulis AI, request serupa, saran artikel, context panel customer, serta aplikasi mobile.
- **Manajer Layanan/ITSM** — memantau analitik Virtual Service Agent, kuota, AIOps incident, sentiment, dan laporan data lake.
- **Customer/Pelapor** — berinteraksi dengan Virtual Service Agent dan agen AI customer-facing di portal, widget, chat, dan email; login ke portal.
- **Integrator/Developer API** — memakai REST API Service Management, Forms, Assets, Operations, webhook, dan extension point app.
- **Tamu/Kolaborator Eksternal** — collaborator tanpa lisensi agent yang membantu lewat internal comment.
- **Stakeholder/Viewer** — melihat laporan dan dashboard analitik tanpa mengubah konfigurasi.

## User story modul

1. Sebagai Admin Situs, saya ingin peran dan permission scheme default Service Management dibuat otomatis agar agent, customer, dan admin langsung mendapat akses yang benar tanpa konfigurasi manual (JSM-AI-001, JSM-AI-003).
2. Sebagai Admin Proyek/Project Lead, saya ingin menentukan siapa yang boleh mengajukan dan berbagi request di service project agar portal saya tidak dibanjiri pengguna yang tidak berhak (JSM-AI-002).
3. Sebagai Agent Layanan, saya ingin mencari request memakai JQL berdasarkan request type, organization, dan sisa waktu SLA agar saya bisa memprioritaskan request yang hampir breach (JSM-AI-013, JSM-AI-014).
4. Sebagai Admin Proyek/Project Lead, saya ingin rule otomasi yang bereaksi saat SLA hampir breach atau approval selesai agar penanganan request tidak bergantung pada pengecekan manual (JSM-AI-016, JSM-AI-018, JSM-AI-021).
5. Sebagai Integrator/Developer API, saya ingin membuat request atas nama customer beserta lampiran sementara lewat REST API agar sistem internal saya bisa mengajukan tiket ke service desk (JSM-AI-024, JSM-AI-025).
6. Sebagai Agent Layanan, saya ingin menyisipkan canned response dan meminta AI meringkas thread panjang agar saya dapat membalas lebih cepat dan konsisten (JSM-AI-035, JSM-AI-037, JSM-AI-038).
7. Sebagai Agent Layanan, saya ingin melihat request serupa yang pernah diselesaikan agar saya dapat memakai ulang solusinya (JSM-AI-039).
8. Sebagai Admin Situs, saya ingin mematikan seluruh fitur AI dari satu pengaturan AI pusat dan mengatur fitur AI Service Management secara granular agar organisasi tetap patuh kebijakan data internal (JSM-AI-036, ADM-052).
9. Sebagai Admin Proyek/Project Lead, saya ingin menyusun intent dan conversation flow Virtual Service Agent tanpa kode agar pertanyaan rutin seperti reset password terselesaikan tanpa agent (JSM-AI-047, JSM-AI-048, JSM-AI-049).
10. Sebagai Customer/Pelapor, saya ingin mendapat jawaban instan dari knowledge base di chat atau portal, dan diteruskan ke manusia bila jawabannya tidak membantu, agar masalah saya cepat selesai (JSM-AI-052, JSM-AI-053, JSM-AI-054).
11. Sebagai Manajer Layanan/ITSM, saya ingin melihat berapa percakapan yang diselesaikan versus dieskalasi per intent dan memantau kuota assisted conversations agar saya bisa menilai ROI Virtual Service Agent (JSM-AI-056, JSM-AI-058, JSM-AI-012).
12. Sebagai Agent Layanan, saya ingin melihat produk, entitlement, dan riwayat interaksi customer di context panel agar saya memberi dukungan sesuai hak kontraknya (JSM-AI-060, JSM-AI-061, JSM-AI-063).
13. Sebagai Admin Situs, saya ingin menghapus akun portal-only atas permintaan customer dan mencatatnya di audit log agar hak penghapusan data terpenuhi (JSM-AI-008, JSM-AI-010).
14. Sebagai Agent Layanan, saya ingin menyetujui approval dan membalas customer dari aplikasi mobile agar request tetap bergerak saat saya tidak di depan komputer (JSM-AI-065).
15. Sebagai Integrator/Developer API, saya ingin extension point portal dan queue untuk app agar saya bisa menambah panel khusus tanpa mengubah inti produk (JSM-AI-033).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 7 |
| P1 | 19 |
| P2 | 35 |
| P3 | 6 |
| **Total** | **67** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 0 |
| Fase 4 | 0 |
| Fase 5 | 27 |
| Fase 6 | 0 |
| Fase 7 | 37 |
| Fase — | 3 |

Strategi urutan implementasi: modul ini sepenuhnya bergantung pada service project, request type, customer, dan queue dari modul 29-31, sehingga tidak ada fitur sebelum Fase 5. Gelombang 5a membangun fondasi administrasi yang wajib ada sebelum portal dibuka untuk customer: role dan permission default, customer permissions, lisensi agent, collaborator, pengaturan global, akun dan autentikasi customer, serta audit/privasi (JSM-AI-001 s.d. JSM-AI-010). Gelombang 5b menambah JQL Service Management (JSM-AI-013 s.d. JSM-AI-015), komponen otomasi Service Management (JSM-AI-016 s.d. JSM-AI-021), REST API Service Management inti dan extension point (JSM-AI-024 s.d. JSM-AI-030, JSM-AI-033), serta canned responses (JSM-AI-035). Fase 7 menampung API Assets/Operations, seluruh kapabilitas AI (dibangun di atas antarmuka `AIProvider` yang dapat dimatikan), Virtual Service Agent, Customer Service Management, mobile, analitik, dan sandbox; di dalam Fase 7 urutannya adalah administrasi AI (JSM-AI-036) → AI assist agent → Virtual Service Agent → agen AI otonom (Request Resolver, Rovo agents) karena tiap lapis memerlukan lapis sebelumnya. Tiga fitur legacy/migrasi (JSM-AI-022, JSM-AI-023, JSM-AI-034) tidak dibangun dan hanya menjadi rujukan. Inventaris menggabungkan lima belas fitur ringkas dari lampiran dan domain inti ke fitur rinci yang setara (lihat baris **Alias inventaris**).

## Fitur

## Administrasi produk, lisensi & izin

### JSM-AI-001 JSM Permissions & Role Defaults

- **Deskripsi:** Peran project dan permission scheme default yang dibuat otomatis untuk setiap service project, memisahkan agent, customer, dan administrator.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengelola permission scheme global), Admin Proyek (menetapkan anggota peran di service project).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap service project company-managed memiliki peran `Service Desk Team` (agent) yang memberi akses queues, reply to customer, internal note, transisi, dan assign.
  2. Peran `Service Desk Customers` dibuat otomatis; customer yang ditambahkan ke service project (portal, email, undangan, API) otomatis menjadi anggota peran ini.
  3. Peran `Administrators` memberi hak Administer Projects pada service project (mengelola request type, portal, SLA, otomasi, people).
  4. Permission scheme menyediakan security type `Service Project Customer - Portal Access` yang dievaluasi benar untuk pengguna yang memiliki akses portal ke service project tersebut.
  5. Permission scheme default Service Management memberi Browse Projects, Create Work Items, Add Comments, dan Create Attachments kepada `Service Project Customer - Portal Access`, sehingga customer dapat bekerja lewat portal tanpa lisensi agent; scheme default tidak dapat dihapus selama masih dipakai service project.
  6. Pengguna hanya dapat bekerja sebagai agent bila memiliki product access Service Management (lisensi agent) DAN peran `Service Desk Team`; menambahkan pengguna tanpa lisensi ke peran agent ditolak dengan pesan "Pengguna memerlukan lisensi agent".
  7. ⚠ Peran bawaan service project team-managed di produk referensi (`Administrator`, `Agent`, `Customer`) tidak dibangun karena clone hanya menyediakan service project company-managed (JSM-PRJ-012); importer memetakan ketiga peran tersebut ke `Administrators`, `Service Desk Team`, dan `Service Desk Customers` (butir 1-3).
  8. Pengguna tanpa lisensi agent yang berada di peran non-agent diperlakukan sebagai collaborator (JSM-AI-004): dapat melihat dan menambah internal comment, tetapi aksi transisi dan public comment ditolak HTTP 403 dan tombolnya disembunyikan.
  9. Internal comment dan work item dengan issue security level yang tidak mencakup customer tidak pernah tampil di portal, email customer, maupun REST API customer-facing; pengujian kebocoran wajib ada di test suite.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); service project hanya company-managed dengan permission scheme dan tiga peran project di butir 1-3.
- **Catatan & ketergantungan:** Model izin umum di [03](03-izin-peran-keamanan.md) (PERM-004, PERM-005); tipe pengguna di JSM-PRJ-045 s.d. JSM-PRJ-050 ([29](29-jsm-service-project-request-type-formulir.md)). Migrasi: tabel `project_roles` dan `permission_schemes` belum ada di repo dan harus dibuat oleh modul 03 lebih dulu.

### JSM-AI-002 Customer Permissions (Space)

- **Deskripsi:** Pengaturan per service project yang menentukan siapa yang boleh mengajukan request, dengan siapa customer dapat berbagi request, dan siapa yang dapat menambah customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengubah pengaturan), Customer/Pelapor (terkena dampak).
- **Aturan perilaku (acceptance criteria):**
  1. Opsi "Who can raise requests" memiliki tiga nilai: hanya customer yang ditambahkan ke service project, siapa pun yang memiliki akun di situs, atau siapa pun di web (anonymous); default adalah hanya customer yang ditambahkan.
  2. Opsi anonymous hanya dapat dipilih bila pengaturan customer access global (JSM-PRT-015, toggle "Allow access without login") mengizinkan; bila tidak, opsi tampil nonaktif dengan tautan ke pengaturan global, dan API menolak dengan HTTP 409 "Akses anonim dinonaktifkan di tingkat situs".
  3. Opsi "Who can customers share requests with" memiliki nilai: tidak dengan siapa pun, hanya anggota organization yang sama, atau customer/organization mana pun di service project; pilihan menentukan daftar pencarian participant di portal.
  4. ⚠ Opsi pencarian customer menentukan apakah customer dapat mencari dan menambahkan customer lain lewat nama/email saat berbagi; bila dimatikan, customer hanya dapat mengetik alamat email lengkap.
  5. ⚠ Opsi "Who can add customers" bernilai agent dan admin, atau hanya admin; agent tanpa hak yang menambah customer mendapat HTTP 403.
  6. Request type restrictions membatasi request type tertentu agar hanya terlihat dan dapat diajukan oleh user/group/organization terpilih (detail di JSM-PRJ-029); customer di luar daftar tidak melihat request type di portal dan API create mengembalikan HTTP 403.
  7. Customer yang dihapus dari service project langsung kehilangan akses portal service project tersebut (sesi berikutnya), tetapi request yang pernah dia ajukan tetap ada dengan reporter tidak berubah dan masih terlihat oleh agent.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012).
- **Catatan & ketergantungan:** Pemilik perilaku: JSM-PRT-016 Customer permissions (project level) ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); fitur ini hanya menambah butir 4-7 (pencarian customer, who can add customers, request type restrictions, efek penghapusan customer), sedangkan butir 1-3 adalah ringkasan JSM-PRT-016 dan bila berbeda JSM-PRT-016 yang berlaku. Halaman pengaturan di service project adalah JSM-PRJ-051 ([29](29-jsm-service-project-request-type-formulir.md)). Prasyarat global (public signup, akses tanpa login) dimiliki JSM-PRT-015 Customer access & signup settings ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); JSM-AI-006 hanya halaman induk yang menautkannya. Anonymous request di JSM-PRT-017.

### JSM-AI-003 Licensing & Agent Seats

- **Deskripsi:** Model lisensi Service Management berbasis jumlah agent; customer tidak memakai lisensi.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs/Billing admin (mengelola product access dan kursi).
- **Aturan perilaku (acceptance criteria):**
  1. Plan Free mengizinkan maksimal 3 agent; menambahkan agent ke-4 ditolak dengan pesan upgrade, dan batas dikonfigurasi sebagai parameter plan, bukan konstanta kode.
  2. Customer (portal-only maupun akun situs yang hanya memakai portal) tidak dihitung sebagai kursi dan jumlahnya tidak dibatasi.
  3. Product access Service Management dikelola dari konsol admin (USR-018); grup default product access otomatis menerima pengguna baru yang diberi akses, dan mencabut akses mengembalikan pengguna ke status collaborator/customer tanpa menghapus datanya.
  4. Penagihan bulanan/tahunan dan harga per agent tidak dibangun sebagai mesin billing; produk ini hanya menampilkan jumlah kursi terpakai vs batas plan dan mengekspos angka tersebut ke sistem billing eksternal (ADM-053).
  5. ⚠ Batas maksimum pengguna per plan (mis. jumlah agent maksimum di Standard/Premium) dapat dikonfigurasi; saat tercapai, undangan agent baru ditolak dan admin melihat banner peringatan pada 90% kapasitas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Batas pengguna umum di USR-020 ([02](02-pengguna-grup-organisasi.md)). Pasangan tumpang tindih: JSM-PRJ-046 Agent role and licensing ([29](29-jsm-service-project-request-type-formulir.md)). Pemilik: fitur ini memiliki penghitungan kursi agent dan batas lisensi tingkat situs; JSM-PRJ-046 memiliki perilaku peran agent di tingkat service project (role Service Desk Team, hak berkomunikasi dengan customer) dan merujuk ke sini untuk batas kursi. Keputusan terbuka: apakah clone self-hosted menerapkan batas kursi sama sekali.

### JSM-AI-004 Collaborators

- **Deskripsi:** Pengguna situs tanpa lisensi agent yang membantu agent menyelesaikan request melalui internal comment.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Tamu/Kolaborator Eksternal atau anggota tim lain (collaborator), Admin Proyek (menambahkan).
- **Aturan perilaku (acceptance criteria):**
  1. Collaborator tidak memakai kursi agent dan tidak tercantum di hitungan lisensi (JSM-AI-003).
  2. Collaborator dapat melihat request di service project tempat ia berperan, menambah internal comment, menjadi watcher, dan di-@mention.
  3. Collaborator tidak dapat menambah public comment (reply to customer), menjalankan transisi, mengubah assignee, maupun membuka halaman queues; percobaan via API mengembalikan HTTP 403.
  4. Collaborator ditambahkan dengan memberi pengguna peran project non-agent (mis. peran kustom atau peran bawaan non-agent) di halaman People; pengguna yang kemudian diberi lisensi agent otomatis naik menjadi agent tanpa langkah tambahan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pasangan tumpang tindih: JSM-PRJ-047 Collaborator role ([29](29-jsm-service-project-request-type-formulir.md)). Pemilik: JSM-PRJ-047 adalah versi lengkap peran collaborator (akses, hak, batasan); fitur ini hanya menambah aspek lisensi (butir 1, dihitung di JSM-AI-003) dan promosi otomatis menjadi agent (butir 4), dan bila terjadi perbedaan JSM-PRJ-047 yang berlaku. @mention di JSM-QSK-014.

### JSM-AI-005 Plan Feature Matrix & Service Collection

- **Deskripsi:** Matriks kapabilitas Service Management per plan (Free, Standard, Premium, Enterprise) dan paket bundel, diterapkan sebagai feature flag tingkat situs.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs/Billing admin (melihat plan), semua pengguna (melihat fitur terkunci).
- **Aturan perilaku (acceptance criteria):**
  1. Plan Free mengaktifkan request management, portal, SLA, knowledge base, queues, otomasi dengan kuota kecil, dan storage terbatas; fitur di luar daftar tampil terkunci dengan label plan minimum.
  2. Plan Standard menambah kapasitas agent dan storage, audit log, data residency, dan dukungan jam kerja (atribut terakhir hanya informatif di clone).
  3. Plan Premium menambah Virtual Service Agent, Assets, incident/change lanjutan, sandbox, release tracks, dan target uptime SLA yang lebih tinggi (informatif di clone); menonaktifkan Premium membuat menu terkait terkunci tetapi data tetap tersimpan.
  4. Plan Enterprise menambah analitik/data lake (JSM-AI-066), kontrol keamanan lanjutan (padanan Atlassian Guard), target uptime tertinggi, dan dukungan 24/7 (informatif); di clone diperlakukan sebagai flag konfigurasi deployment.
  5. ⚠ Bundel "Service Collection" (Service Management + Customer Service Management + Assets + AI) dimodelkan sebagai preset yang mengaktifkan sekumpulan flag sekaligus; komposisi bundel dapat dikonfigurasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Batas plan service project di JSM-PRJ-070; koleksi lisensi ADM-056 dan prompt upgrade NAV-024. Angka storage/uptime adalah keputusan pemilik produk.

### JSM-AI-006 JSM Global Settings

- **Alias inventaris:** Jira Service Management product settings (lampiran); Global Jira Service Management configuration
- **Deskripsi:** Halaman pengaturan produk Service Management di tingkat situs yang memuat konfigurasi global akses customer, email, help center, knowledge base, Assets, Operations, AI, dan bahasa. Heading ini menggabungkan entri lampiran administrasi sistem dan lampiran Service Management.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tab Configuration memuat opsi global seperti izin collaborator menjadi approver dan aturan berbagi request lintas service project; setiap opsi memiliki default yang terdokumentasi.
  2. Tab Customer access menautkan pengaturan public signup (on/off) untuk semua help center yang perilakunya dimiliki JSM-PRT-015; bila public signup mati, service project tidak dapat memilih "siapa pun di web" atau "siapa pun dengan akun" yang memerlukan signup (JSM-AI-002).
  3. CAPTCHA pada formulir signup dan request anonim dapat diaktifkan global; bila aktif, signup tanpa token CAPTCHA valid ditolak HTTP 400 (AUTH-006).
  4. Customer access default menentukan nilai awal "Who can raise requests" untuk service project baru dan apakah help center dapat diakses tanpa login.
  5. ⚠ Request type restrictions global menentukan apakah Admin Proyek boleh membatasi request type ke group/organization (JSM-PRJ-029).
  6. Tab Email requests memuat sakelar global on/off untuk seluruh email channel; saat off, mailbox tidak diproses dan email masuk dicatat sebagai "diabaikan: dinonaktifkan global" di global email log (JSM-AI-007).
  7. Email channel per service project tetap dikonfigurasi di JSM-PRJ-063, tetapi halaman global menampilkan daftar semua mailbox yang terhubung beserta service project-nya.
  8. Pengaturan branding/layout help center global (logo, warna, nama help center) berlaku sebagai default untuk semua help center (JSM-PRT-002, JSM-PRT-003); custom domain dikelola di JSM-PRT-009.
  9. Look & feel template email customer global (logo, warna, footer) dikelola dari halaman ini dan dipakai semua service project (JSM-PRT-038).
  10. Tab Customer accounts menautkan ke pengelolaan akun portal-only (JSM-AI-008).
  11. Pengaturan organization menentukan siapa yang dapat membuat dan mengelola organization: hanya admin situs, atau juga Admin Proyek dan agent; pembuatan organization di luar hak ditolak HTTP 403.
  12. Tab Knowledge base memuat pengaturan global KB (mis. apakah customer dapat melihat artikel tanpa login, sumber pengetahuan eksternal yang diizinkan).
  13. Tab Assets memuat konfigurasi global Assets, akses schema, dan imports (JSM-AST-003).
  14. Tab Operations memuat API key, heartbeats, maintenance, dan notifikasi Operations global (JSM-OPS-003).
  15. Tab AI & Virtual agent menampilkan status kapabilitas AI situs (JSM-AI-036) dan default Virtual Service Agent (JSM-AI-047), serta pengaturan knowledge untuk AI answers.
  16. Tab Languages mengatur bahasa yang tersedia untuk help center dan bahasa default (JSM-PRT-007).
  17. Default approvals (mis. perilaku saat approver tidak tersedia) dan template SLA default untuk service project baru dikonfigurasi dari halaman ini.
  18. Setiap perubahan pengaturan global dicatat di audit log (ADM-014) dengan nilai sebelum dan sesudah.
- **Team-managed vs company-managed:** Tidak ada perbedaan; pengaturan global berlaku untuk semua service project (hanya company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Pasangan tumpang tindih: JSM-PRJ-019 JSM global configuration (product settings) ([29](29-jsm-service-project-request-type-formulir.md)). Pemilik: fitur ini adalah induk halaman pengaturan produk Service Management tingkat situs; JSM-PRJ-019 hanya mendefinisikan subset pengaturan yang memengaruhi service project dan diimplementasikan sebagai bagian halaman ini. Perilaku customer access dan signup dimiliki JSM-PRT-015 ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); halaman ini hanya menautkannya.

### JSM-AI-007 Email Requests Global Settings & Limits

- **Deskripsi:** Pengaturan dan pembatasan email channel di tingkat situs yang melindungi sistem dari spam dan loop email.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Global email log menampilkan setiap email masuk lintas mailbox dengan waktu, pengirim, mailbox, service project, hasil (dibuat request, ditambah komentar, diabaikan, gagal) dan alasan; log dapat difilter dan disimpan sesuai retensi yang dapat dikonfigurasi.
  2. Allowlist dan blocklist pengirim (alamat atau domain) dapat dikelola global; email dari pengirim di blocklist diabaikan dan dicatat, sedangkan bila allowlist tidak kosong hanya pengirim di allowlist yang diproses.
  3. ⚠ Rate limit email per mailbox dan per pengirim (nilai dapat dikonfigurasi, mis. jumlah email per jam); email yang melebihi batas ditunda atau diabaikan dengan status "rate limited" di log dan admin menerima notifikasi.
  4. Custom mailbox dapat dihubungkan via OAuth Google, OAuth Microsoft, atau IMAP/POP dengan kredensial terenkripsi; kegagalan autentikasi menandai mailbox "perlu dihubungkan ulang" dan memberi tahu admin.
  5. Domain email keluar dapat diverifikasi (SPF/DKIM) agar notifikasi customer dikirim dari alamat domain perusahaan; domain yang belum terverifikasi memaksa pemakaian alamat default produk ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail email channel per service project di JSM-PRT-040 s.d. JSM-PRT-043 ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); infrastruktur mail di ADM-012, ADM-013 ([22](22-administrasi-sistem.md)).

### JSM-AI-008 Customer Accounts Management

- **Deskripsi:** Pengelolaan akun customer, khususnya akun portal-only, dari konsol admin situs.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs/Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. Sistem membedakan dua tipe akun customer: portal-only (hanya berlaku di help center situs ini) dan akun penuh situs; tipe ditampilkan sebagai label di daftar akun.
  2. Admin dapat mengubah nama tampilan dan email akun portal-only; perubahan email memerlukan verifikasi ke alamat baru, dan admin dapat mengirim tautan reset password.
  3. Admin dapat menonaktifkan akun portal-only (tidak dapat login, request tetap ada) atau menghapusnya permanen (data pribadi dianonimkan, lihat JSM-AI-010) dengan konfirmasi.
  4. Akun portal-only dapat dimigrasikan menjadi akun penuh situs; riwayat request, organization, dan participant dipertahankan, dan migrasi tidak dapat dibatalkan.
  5. Daftar customer accounts per situs dapat dicari berdasarkan nama/email, difilter berdasarkan status (aktif/nonaktif) dan tipe, serta dipaginasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Overlap dengan JSM-PRT-014 dan JSM-PRT-020 ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); migrasi: tabel `users` saat ini tidak membedakan tipe akun, perlu kolom `account_type` atau tabel `customer_accounts` terpisah.

### JSM-AI-009 Customer Authentication

- **Deskripsi:** Metode login customer ke help center dan portal.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengonfigurasi), Customer/Pelapor (login).
- **Aturan perilaku (acceptance criteria):**
  1. Customer portal-only login dengan email + password; kebijakan kata sandi dan proteksi brute force mengikuti AUTH-006, dan login gagal menampilkan pesan generik tanpa membocorkan apakah email terdaftar.
  2. ⚠ SSO SAML untuk customer portal-only dapat dikonfigurasi lewat identity provider terpisah dari SSO karyawan; bila aktif, login password untuk customer tersebut dimatikan.
  3. ⚠ Login sosial Google dan Microsoft untuk customer dapat diaktifkan per situs; akun yang dibuat lewat login sosial ditautkan ke email yang terverifikasi oleh penyedia.
  4. ⚠ Metode autentikasi dapat diatur per help center (JSM-PRT-008), sehingga help center berbeda dapat memakai IdP berbeda.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** SAML umum di USR-034 ([02](02-pengguna-grup-organisasi.md)); kebijakan SSO customer di JSM-PRT-018. Migrasi: auth repo saat ini berbasis JWT tunggal (`pkg/jwt`); sesi customer memerlukan audience/claim terpisah agar token customer tidak dapat memanggil API agent.

### JSM-AI-010 JSM Audit & Customer Data Privacy

- **Deskripsi:** Pencatatan audit untuk event Service Management dan kontrol privasi data customer (termasuk hak penghapusan).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs/Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. Audit log situs (ADM-014) mencatat event Service Management: perubahan customer permissions, request type, SLA, portal, email channel, pengaturan global, dan penambahan/penghapusan customer/organization.
  2. Audit log organisasi (USR-045) mencatat event akun customer: pembuatan, perubahan email, nonaktif, hapus, migrasi akun.
  3. Penghapusan akun portal-only (right to erasure) menganonimkan nama dan email di semua request, komentar, dan audit trail menjadi "Former customer" dalam batas waktu yang dapat dikonfigurasi (default 30 hari), dan proses tidak dapat dibatalkan.
  4. Personal data reporting API memungkinkan app terdaftar menanyakan akun mana yang telah dihapus/diubah agar app menghapus salinan datanya (USR-044).
  5. Help center menampilkan tautan Terms dan Privacy Policy yang dapat dikonfigurasi admin di footer dan halaman signup.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** NFR-017, NFR-018 ([27](27-nfr-arsitektur-platform.md)).

### JSM-AI-011 JSM Data Residency

- **Deskripsi:** Penempatan data Service Management pada region penyimpanan tertentu.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat memilih realm data (mis. US, EU, Jerman, Australia, Kanada, India, Jepang, Singapura, UK) dari daftar yang dikonfigurasi deployment; pada clone self-hosted, realm diterjemahkan sebagai lokasi database/object storage yang dipilih operator.
  2. ⚠ Data in-scope mencakup request, komentar, lampiran, customer, dan konfigurasi Service Management; status in-scope data Assets dan Operations ditampilkan eksplisit per kategori.
  3. Pemindahan realm dijalankan pada migration window yang dijadwalkan admin; selama window, situs dalam mode read-only dan banner pemeliharaan tampil.
  4. ⚠ Fitur ini hanya tersedia pada plan Standard ke atas; pada Free menu tampil terkunci.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data residency umum di USR-042 dan ADM-048.

### JSM-AI-012 Track Feature Usage

- **Deskripsi:** Halaman admin "Track your feature usage" yang menampilkan pemakaian fitur Service Management berkuota terhadap batas plan.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Org admin saja.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman menampilkan daftar fitur berkuota (assisted conversations Virtual Service Agent, rule runs otomasi, objek Assets, alert/SMS Operations) dengan pemakaian saat ini, batas plan, dan persentase.
  2. Indikator kuning tampil pada ≥ 80% dan merah pada ≥ 100% batas, disertai ajakan upgrade plan atau tautan pengaturan batas.
  3. Baris Virtual Service Agent menautkan ke "Manage the usage limit" dan tagihan extra assisted conversations (JSM-AI-058).
  4. Setiap fitur menampilkan periode reset pemakaian (umumnya bulanan, tanggal reset berikutnya).
  5. Halaman dan API pemakaian hanya dapat diakses Admin Situs/Org admin; pengguna lain mendapat HTTP 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Usage umum di ADM-042; kuota otomasi AUTO-015; kuota Assets JSM-AST-004.

## JQL Service Management

### JSM-AI-013 JSM JQL Fields

- **Alias inventaris:** JSM Fields & Functions (Approvals, SLA, Request, Assets); JSM JQL extensions
- **Deskripsi:** Field JQL khusus Service Management untuk mencari request berdasarkan request type, organization, participant, kanal, kepuasan, SLA, Assets, dan field ITSM. Heading ini menggabungkan lampiran field/fungsi Service Management dari domain JQL dan lampiran Service Management.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek; hasil selalu dibatasi Browse Projects dan issue security pencari.
- **Aturan perilaku (acceptance criteria):**
  1. Field `"Request Type"` dan aliasnya `"Customer Request Type"` menerima nama request type atau key internal (mis. `"Request Type" = "Get IT help"`); nama yang sama di beberapa service project cocok dengan semuanya kecuali dibatasi `project = X`.
  2. `"Request Type" IS EMPTY` mengembalikan work item di service project yang tidak memiliki request type (mis. dibuat langsung dari Jira tanpa request type).
  3. Field `Organizations` menerima nama atau id organization dan mendukung operator `=`, `!=`, `IN`, `NOT IN`, `IS EMPTY`.
  4. Field `"Request participants"` menerima akun pengguna, mendukung fungsi `currentUser()`, dan cocok bila pengguna adalah salah satu participant.
  5. Field `Satisfaction` bernilai 1-5 dan mendukung operator perbandingan (`>=`, `<` dst.); `"Satisfaction date"` adalah field tanggal yang mendukung fungsi tanggal relatif (mis. `-7d`).
  6. Field `"Request channel type"` menerima nilai `email`, `portal`, `api`, `jira`, `anonymousportal`, `widget`, dan `chat` (selaras JSM-PRJ-061 butir 3); nilai tak dikenal menghasilkan galat validasi yang menyebut daftar nilai yang valid.
  7. ⚠ Field `"Request language"` memfilter request berdasarkan bahasa yang dipakai customer saat mengajukan (kode bahasa help center).
  8. ⚠ Field `"Request last activity time"` memfilter berdasarkan waktu aktivitas terakhir customer atau agent pada request, mendukung operator tanggal.
  9. Field SLA (`"Time to first response"`, `"Time to resolution"`, dan setiap SLA kustom berdasarkan namanya) dapat dipakai di JQL hanya bersama fungsi SLA (JSM-AI-014); nama SLA yang tidak ada menghasilkan galat "Field tidak ditemukan".
  10. Field `Approvals` dapat dipakai hanya bersama fungsi approval (JSM-AI-015).
  11. ⚠ Field Assets object mendukung `"<Asset field>" in aqlFunction("<AQL>")`; nama lama `iqlFunction()` diterima sebagai alias dengan peringatan deprecated di UI (JSM-AST-028).
  12. Field ITSM `"Affected services"`, `Urgency`, `Impact`, dan `"Pending reason"` dapat dicari dengan operator standar field select/objek (JSM-ITSM-005).
  13. ⚠ Field change management `"Change type"`, `"Change risk"`, `"Change start date"`, `"Change completion date"`, dan `"Change gating type"` tersedia di JQL dengan tipe sesuai definisinya (JSM-ITSM-036).
  14. ⚠ Fungsi `organizationMembers("<nama organization>")` mengembalikan anggota organization dan dapat dipakai pada field pengguna (mis. `reporter in organizationMembers("Acme")`).
  15. ⚠ Field `Sentiment` dari Customer Sentiment Analysis (JSM-AI-042) dapat dicari dengan nilai `Positive`, `Neutral`, `Negative`.
  16. Contoh kueri yang harus didukung dan diuji: `"Time to resolution" = breached()` dan `"Time to first response" < remaining("1h")`.
  17. Fungsi SLA, fungsi approval, dan `aqlFunction()` hanya valid pada field SLA, Approvals, dan Assets yang sesuai; memakai fungsi pada field lain menghasilkan galat validasi "Fungsi <nama> tidak didukung pada field <field>" sebelum kueri dijalankan.
  18. Semua field di atas muncul di autocomplete JQL dan di daftar field pencarian dasar navigator ketika pengguna memiliki akses ke minimal satu service project.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012).
- **Catatan & ketergantungan:** Grammar dan kompatibilitas operator di JQL-013 ([12](12-jql-bahasa-kueri.md)); pasangan tumpang tindih JSM-PRJ-061 JSM JQL extensions for requests ([29](29-jsm-service-project-request-type-formulir.md)). Pemilik: fitur ini memiliki daftar lengkap field/fungsi JQL Service Management; JSM-PRJ-061 hanya mendefinisikan subset field request yang wajib tersedia saat service project pertama dirilis dan merujuk ke sini. Overlap lain: JSM-QSK-027 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)), dan JSM-ITSM-014 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)).

### JSM-AI-014 JSM JQL SLA Functions

- **Deskripsi:** Fungsi JQL yang mengevaluasi kondisi siklus SLA (breach, berjalan, jeda, sisa waktu).
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. `<SLA> = breached()` cocok bila siklus SLA berjalan atau terakhir telah melampaui goal.
  2. `<SLA> = everBreached()` cocok bila siklus SLA mana pun (termasuk yang sudah selesai) pernah breach.
  3. `<SLA> = completed()` cocok bila siklus terakhir sudah berhenti (stop condition terpenuhi).
  4. `<SLA> = running()` cocok bila ada siklus aktif yang tidak sedang dijeda.
  5. `<SLA> = paused()` cocok bila siklus aktif sedang dijeda oleh pause condition atau di luar jam kalender.
  6. `<SLA> > elapsed("<durasi>")` atau `<` membandingkan waktu yang sudah berjalan pada siklus aktif; format durasi menerima `w`, `d`, `h`, `m` (mis. `"1h 30m"`) dan format tidak valid menghasilkan galat.
  7. `<SLA> < remaining("<durasi>")` membandingkan sisa waktu menuju goal; nilai negatif berarti sudah breach.
  8. `<SLA> = withinCalendarHours()` cocok bila waktu saat ini berada di dalam jam kerja kalender SLA tersebut.
  9. Fungsi SLA dapat dikombinasikan dengan `AND`/`OR`/`NOT` dan dengan field lain, mis. `"Time to resolution" = remaining("2h") AND priority = High`.
  10. Work item yang SLA-nya tidak memiliki siklus aktif tidak pernah cocok dengan `running()` atau `paused()`, dan tidak cocok dengan `elapsed()`/`remaining()`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi SLA di JSM-QSK-022 s.d. JSM-QSK-025; fitur ini adalah spesifikasi evaluasi JQL yang dipakai JSM-QSK-027. Evaluasi `remaining()` bergantung waktu sehingga hasilnya tidak boleh di-cache lebih dari 60 detik.

### JSM-AI-015 JSM JQL Approval Functions

- **Deskripsi:** Fungsi JQL berbasis status approval pada field Approvals.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, approver.
- **Aturan perilaku (acceptance criteria):**
  1. `approval = pending()` cocok dengan work item yang memiliki approval step aktif yang belum diputuskan.
  2. `approval = pendingBy(<user>, ...)` cocok bila salah satu pengguna yang disebut masih harus memberi keputusan; argumen pengguna tidak valid menghasilkan galat.
  3. `approval = myPending()` setara `pendingBy(currentUser())` dan dipakai untuk queue "Menunggu persetujuan saya".
  4. `approval = myApproval()` cocok dengan work item tempat pengguna saat ini pernah atau sedang menjadi approver.
  5. `approval = approver(<user>, ...)` cocok dengan work item tempat pengguna yang disebut adalah approver, apa pun keputusannya.
  6. `approval = approved()` cocok bila approval step terakhir disetujui.
  7. `approval = declined()` cocok bila approval step terakhir ditolak.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Approval di JSM-ITSM-011 s.d. JSM-ITSM-014 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)).

## Otomasi Service Management

### JSM-AI-016 Jira Automation - JSM Triggers

- **Alias inventaris:** Jira Service Management triggers
- **Deskripsi:** Trigger khusus Service Management pada mesin otomasi produk ini (modul 20), mencakup SLA, approval, alert, Assets, service limit, Slack, forms, dan filter request pada trigger generik.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (rule project), Admin Situs (rule global/multi-project).
- **Aturan perilaku (acceptance criteria):**
  1. Trigger `SLA threshold breached` meminta pilihan SLA dan ambang: "akan breach dalam <durasi>", "breached", atau "breached <durasi> lalu"; trigger dipicu tepat sekali per siklus SLA per ambang.
  2. Trigger `Approval required` dipicu saat work item memasuki status dengan approval step, saat approval di-update, atau saat approver baru ditambahkan.
  3. Trigger `Approval completed` dipicu saat approval step diputuskan (approved atau declined) dan menyediakan hasil keputusan di smart value.
  4. Trigger `Alert created`, `Alert updated` (priority/summary/description), `Alert status changed`, dan `Alert note added` tersedia untuk alert Operations (JSM-OPS-005).
  5. Trigger `Manual trigger from alert` mengatur siapa yang boleh menjalankan, filter alert tempat tombol muncul, dan prompt input opsional bagi pengguna.
  6. Trigger `Assets object` (created/updated/deleted per object schema) dan `Attribute value changed` hanya tersedia di rule global; menyediakan smart value `{{attributeChangelog}}` dan `{{containerObject}}`.
  7. Trigger `Service limit breached` dipicu saat waktu proses rule harian/per jam mendekati batas (AUTO-016) dan menyediakan `{{breachedSummary}}` dan `{{breachedRules}}`.
  8. Trigger `Emoji reaction to Slack message` dipicu saat pengguna memberi reaksi emoji pada pesan yang dibuat integrasi chat Service Management (JSM-PRT-045).
  9. Trigger form `Form attached`, `Form opened for edits`, dan `Form submitted` tersedia untuk forms di work item (AUTO-029).
  10. Trigger generik mendukung filter Service Management: work item created per request type/kanal, field value changed untuk `Request Type`/`Organizations`/`Request participants`, dan work item commented.
  11. Trigger `Work item commented` memiliki filter visibilitas komentar: semua, hanya public (reply to customer), atau hanya internal.
  12. ⚠ Trigger berbasis event sistem HR (karyawan baru, perubahan status) tersedia bila koneksi HR native dikonfigurasi (JSM-AI-018).
  13. ⚠ Trigger Service Management mengekspos smart values konteks `{{approval}}`, `{{sla}}`, `{{customer}}`, dan `{{organization}}` (JSM-AI-020).
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); rule Service Management hanya berjalan di service project company-managed.
- **Catatan & ketergantungan:** Tidak ada trigger khusus request type/organization/customer; filter dilakukan via trigger generik (AUTO-018, AUTO-021). Trigger SLA di JSM-QSK-028.

### JSM-AI-017 Jira Automation - JSM Conditions

- **Deskripsi:** Kondisi otomasi untuk konteks service request, sebagian besar memakai kondisi generik dengan field Service Management.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. User condition dapat memeriksa apakah pengguna adalah customer atau agent, serta keanggotaan organization, group, atau peran project (AUTO-034).
  2. Visibilitas komentar dapat diperiksa via smart value `{{comment.internal}}` (true/false) atau kondisi comment visibility.
  3. Request type diperiksa via Work item fields condition (AUTO-032) dengan nilai request type terpilih.
  4. Organizations diperiksa via field condition (contains any/all/none).
  5. Status approval diperiksa via JQL condition dengan fungsi approval (JSM-AI-015).
  6. Kondisi SLA diperiksa via JQL condition dengan fungsi SLA (JSM-AI-014).
  7. Kondisi Assets memakai AQL condition yang menguji objek terkait (JSM-AST-024).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kondisi generik di AUTO-032 s.d. AUTO-038 ([20](20-otomasi.md)).

### JSM-AI-018 Jira Automation - JSM Actions

- **Alias inventaris:** Jira Service Management actions
- **Deskripsi:** Aksi otomasi khusus Service Management untuk request, customer, approval, email, forms, Assets, dan koneksi HR native.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs; aksi dijalankan atas nama rule actor.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi `Create work item with a request type` (nama lama "Create service request") membuat request dengan request type terpilih; hidden fields tidak didukung kecuali untuk Major Incident pada plan Premium/Enterprise.
  2. Aksi create mendukung opsi "Raise on behalf of" untuk menetapkan reporter customer, dan menolak (log error) bila customer tidak memiliki akses ke service project.
  3. Aksi `Edit request type` menolak dengan error di audit log rule bila request type tujuan tidak cocok dengan work item type work item.
  4. Aksi `Approve/Decline request` memutuskan approval pending tanpa memperhitungkan jumlah approver yang dibutuhkan, dan tidak menerima smart value sebagai input keputusan.
  5. Aksi `Add service project customer` menambah customer (opsional ke organization); customer baru mungkin baru muncul di pencarian setelah jeda indexing hingga ±15 detik, dan aksi berikutnya yang membutuhkannya harus memakai re-fetch.
  6. ⚠ Aksi `Add/remove request participants` menambah atau menghapus participant dari daftar pengguna atau smart value.
  7. Aksi `Send customized email` mendukung penerima "All customers involved" (reporter, participant, anggota organization, subscriber notifikasi).
  8. Field Organizations di aksi Create/Clone/Edit/Transition mendukung opsi "Use reporter's email domain" (memetakan domain email terkelola ke organization) dan "Create organization if it doesn't exist".
  9. Aksi `Comment on work item`, `Edit comment`, dan `Delete comment` memiliki pengaturan visibilitas public atau internal; default untuk rule service project adalah internal.
  10. Aksi forms: `Attach forms`, `Attach forms with values`, `Change form status`, dan `Copy forms` (JSM-PRJ-034).
  11. Aksi Assets: `Create object`, `Edit object` (di rule global atau branch AQL), `Edit Assets field attributes`, dan `Lookup objects` yang mengembalikan maksimal 100 objek ke `{{lookupObjects}}`.
  12. Koneksi HR native (padanan "Use Workday with Automation") dapat dihubungkan ke otomasi dari service project company-managed dengan kredensial tersimpan aman (AUTO-056).
  13. ⚠ Aksi lookup data karyawan dan aksi HR dari koneksi tersebut tersedia untuk alur onboarding/offboarding dan journeys (JSM-ITSM-050).
  14. ⚠ Aksi alert Operations (create/acknowledge/close) dispesifikasikan di JSM-AI-019.
  15. ⚠ Aksi Statuspage incident dispesifikasikan di JSM-AI-019.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); koneksi HR dan aksi lain berjalan di service project company-managed.
- **Catatan & ketergantungan:** Aksi generik di AUTO-039 s.d. AUTO-060 ([20](20-otomasi.md)). Koneksi HR native adalah integrasi first-party; iPaaS umum di INT-003 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JSM-AI-019 Jira Automation - Incident & Ops Actions

- **Deskripsi:** Aksi otomasi untuk incident management dan Operations.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Incident manager (Manajer Layanan/ITSM).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi `Create incident` membuat incident dengan request type tertentu, dengan opsi "lewati bila incident terbuka sejenis sudah ada" berdasarkan kriteria JQL.
  2. Aksi `Add stakeholders to incident` menambah pengguna/team sebagai stakeholder (JSM-ITSM-025).
  3. Aksi `Create post-incident review` membuat PIR dengan judul/deskripsi dari smart values dan menautkannya ke incident (JSM-ITSM-032).
  4. Aksi `Create meeting for incident` membuat tautan conference call melalui integrasi video first-party, dengan opsi melampirkan rekaman (JSM-ITSM-024).
  5. Aksi `Create Slack channel for incident` dan `Set channel topic for incident Slack channel` membuat dan memperbarui kanal chat incident (JSM-ITSM-023).
  6. Aksi `Send update to Statuspage` (Enterprise) membuat, memperbarui, atau me-resolve incident status page (JSM-ITSM-027).
  7. Aksi `Assign to on-call responder` menetapkan assignee ke pengguna yang sedang on-call pada schedule team terpilih; bila tidak ada yang on-call, aksi dicatat gagal di audit log rule.
  8. ⚠ Aksi alert `Create alert`, `Acknowledge alert`, dan `Close alert` tersedia untuk rule yang terhubung ke Operations (JSM-OPS-010).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Video meeting dan Statuspage adalah layanan pihak ketiga; yang dibangun adalah konektornya (INT-052).

### JSM-AI-020 Jira Automation - JSM Smart Values

- **Deskripsi:** Smart values khusus Service Management untuk template teks di otomasi.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. `{{issue.Request Type}}` dan `{{issue.Customer Request Type}}` mengembalikan objek request type dengan atribut `name` dan `description`.
  2. `{{issue.url.customer}}` mengembalikan tautan request di portal customer.
  3. `{{issue.Organizations}}` dan `{{issue.Request participants}}` mengembalikan daftar yang dapat diiterasi (mis. `{{#issue.Organizations}}{{name}}{{/}}`).
  4. Smart value SLA seperti `{{issue.Time to resolution.ongoingCycle.remainingTime.friendly}}`, serta atribut `breached`, `elapsedTime`, dan `goalDuration` tersedia untuk setiap SLA.
  5. `{{issue.Satisfaction}}` dan `{{issue.Satisfaction date}}` mengembalikan rating CSAT dan tanggalnya.
  6. `{{comment.internal}}` bernilai `true` untuk internal note dan `false` untuk reply to customer.
  7. `{{lookupObjects}}`, `{{attributeChangelog}}`, dan `{{containerObject}}` tersedia untuk rule Assets.
  8. `{{alert}}`, `{{alert.status}}`, dan `{{alert.note}}` tersedia untuk rule berbasis trigger alert.
  9. `{{breachedSummary}}` dan `{{breachedRules}}` tersedia untuk trigger Service limit breached.
  10. ⚠ Smart value approval (mis. `{{approval.decision}}`, approver) dan `{{issue.Request language}}` tersedia; smart value yang tidak dikenal dirender kosong dan dicatat sebagai peringatan di audit log rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sintaks dan debugging smart value di AUTO-063 s.d. AUTO-065.

### JSM-AI-021 Automation Templates & Usage Limits for JSM

- **Alias inventaris:** JSM automation specifics
- **Deskripsi:** Template rule Service Management siap pakai, assignment otomatis, dan batas rule runs per plan. Heading ini menggabungkan ringkasan komponen otomasi khusus Service Management dari lampiran.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (memakai template), Admin Situs (melihat kuota situs).
- **Aturan perilaku (acceptance criteria):**
  1. Pustaka template Service Management memuat minimal: auto-assign, re-open on customer comment, close resolved requests setelah X hari, SLA breach alert, dan approval reminder; template dapat diedit sebelum diaktifkan.
  2. Template "Auto-close resolved requests" menutup request berstatus resolved setelah X hari tanpa aktivitas customer (X dapat diatur, default 3).
  3. Template "Reopen on customer comment" dan "Transition when customer replies" mentransisikan request ke status tertentu saat customer berkomentar dan mengabaikan komentar agent/internal.
  4. Template incident/Operations dan Assets tersedia di kategori terpisah (JSM-ITSM-048).
  5. Aksi assignment otomatis mendukung metode round-robin, balanced workload (berdasarkan jumlah work item terbuka), dan random, dibatasi ke anggota peran/grup tertentu (AUTO-041).
  6. Trigger Service Management yang tersedia didefinisikan di JSM-AI-016 (SLA, approval, komentar dengan filter customer).
  7. Aksi Service Management yang tersedia (participants, organization, edit request type, notifikasi/email customer, komentar public/internal, create alert) didefinisikan di JSM-AI-018 dan JSM-AI-019.
  8. Kondisi Service Management (komentar public/internal, request type, komentar oleh customer) didefinisikan di JSM-AI-017.
  9. ⚠ Smart values Service Management (mis. `{{issue.Request Type}}`, `{{issue.customerRequest.link}}` sebagai alias `{{issue.url.customer}}`) didefinisikan di JSM-AI-020.
  10. ⚠ Batas rule runs per bulan ditentukan per plan Service Management dan dapat dikonfigurasi (Enterprise dapat tanpa batas); nilai default adalah keputusan pemilik produk.
  11. Hanya eksekusi rule yang menjalankan minimal satu aksi yang dihitung terhadap kuota; rule yang berhenti di kondisi tidak dihitung.
  12. Admin menerima notifikasi saat pemakaian mencapai 80% dan 100% kuota; setelah kuota habis, rule yang dihitung berhenti berjalan sampai periode reset dan audit log mencatat "Kuota habis".
  13. Dashboard pemakaian menampilkan rule runs per rule dan per project untuk periode berjalan (AUTO-015).
  14. Service limits (waktu proses per rule per jam/hari) berlaku sesuai AUTO-016 dan memicu trigger Service limit breached (JSM-AI-016).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rule default service project di JSM-PRJ-033; pustaka template umum AUTO-010.

### JSM-AI-022 Legacy JSM Automation

- **Alias inventaris:** Legacy JSM automation
- **Deskripsi:** Mesin otomasi lama khusus service desk (WHEN/IF/THEN per project) yang di ekosistem rujukan sudah dimigrasi ke mesin otomasi umum dan hanya tersisa di edisi self-managed. Tidak dibangun di produk ini; dicatat sebagai rujukan pemetaan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan: rule lama berada di Project settings > Automation (legacy) dengan format WHEN (trigger), IF (condition), THEN (action); produk ini tidak menyediakan halaman tersebut.
  2. ⚠ Rujukan pemetaan trigger WHEN lama (Issue created, Comment added, Comment edited, Status changed, Issue resolved, Participant added, Organization added, Approval required, SLA time remaining, Linked issue transitioned) ke trigger modul 20 dan JSM-AI-016.
  3. Rujukan pemetaan kondisi IF lama (comment internal/public, comment is primary action, user type customer/agent, request type, issue matches JQL, comment contains) ke JSM-AI-017.
  4. ⚠ Rujukan pemetaan aksi THEN lama (Transition issue, Add comment, Alert user, Edit request type, Webhook) ke JSM-AI-018 dan AUTO-054.
  5. Template lama (Transition on comment, Re-open on customer comment, Close resolved request, Keep on top of SLAs, Triage issues, Alert agents of high-priority issues, Update when a linked issue changes) tersedia sebagai padanan di pustaka template JSM-AI-021.
  6. Opsi "run rule as" (pengguna pemicu vs project default) dipetakan ke rule actor otomasi modul 20.
  7. Audit log dan error log per rule dipetakan ke audit log otomasi (AUTO-014).
  8. Migrasi otomatis rule legacy tidak diperlukan karena produk ini tidak pernah memiliki mesin legacy; impor dari instalasi self-managed memakai jalur migrasi JSM-AI-067.
  9. Mesin legacy tidak tersedia untuk service project baru maupun lama di produk ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

### JSM-AI-023 Legacy Assets Automation

- **Deskripsi:** Otomasi lama tingkat object schema Assets yang di ekosistem rujukan digantikan trigger/aksi Assets di mesin otomasi umum. Tidak dibangun; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Rujukan: rule lama per object schema berformat When (object created/updated/deleted, attachment added, scheduled AQL), If (AQL condition), Then (update attribute, send email, webhook); padanannya adalah trigger Assets (JSM-AI-016), AQL condition (JSM-AI-017), dan aksi Assets (JSM-AI-018).
  2. Migrasi ke rule global otomasi tidak diperlukan karena produk ini tidak memiliki mesin legacy.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## REST API Service Management & ekstensibilitas

Konvensi sub-area ini: base path REST API Service Management produk ini adalah `/api/v1/servicedeskapi` (padanan `/rest/servicedeskapi` di ekosistem rujukan); semua path endpoint di bawah ditulis relatif terhadap base tersebut kecuali disebut lain. Format galat mengikuti `internal/shared/delivery/http/response` dan konvensi API publik modul 23.

### JSM-AI-024 JSM REST API - Info, Service Desk & Temporary Files

- **Alias inventaris:** Jira Service Management REST API (Appendix); JSM REST APIs
- **Deskripsi:** Fondasi REST API Service Management: metadata, daftar service desk, lampiran sementara, pagination, endpoint experimental, autentikasi, dan rate limit; sekaligus payung seluruh keluarga API Service Management yang dirinci di JSM-AI-025 s.d. JSM-AI-032. Heading ini menggabungkan dua entri ringkas REST API Service Management dari lampiran.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API; akses mengikuti permission pengguna pemanggil (agent, customer, atau app).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /info` mengembalikan versi API, versi build, dan tautan dokumentasi tanpa memerlukan autentikasi.
  2. `GET /servicedesk` mengembalikan service desk yang dapat diakses pemanggil dengan atribut `id`, `projectId`, `projectKey`, `projectName`.
  3. `GET /servicedesk/{serviceDeskId}` menerima id service desk, dan juga `projectKey` atau `projectId` sebagai pengenal; pengenal tak dikenal atau tanpa akses mengembalikan HTTP 404.
  4. `POST /servicedesk/{serviceDeskId}/attachTemporaryFile` (multipart) mengunggah file dan mengembalikan `temporaryAttachmentId`; id tersebut dipakai saat create request atau tambah attachment, dan file sementara yang tidak dipakai dihapus otomatis setelah batas waktu yang dapat dikonfigurasi (default 1 jam).
  5. Semua endpoint daftar memakai pagination `start`/`limit` dan mengembalikan `size`, `start`, `limit`, `isLastPage`, `values`, serta `_links`; `limit` maksimum dapat dikonfigurasi (default 50) dan nilai lebih besar dipotong ke maksimum.
  6. Parameter `expand` memuat sub-objek opsional (mis. `participant`, `status`, `sla`, `requestType`); nilai expand tak dikenal diabaikan.
  7. Endpoint experimental hanya dapat dipanggil dengan header `X-ExperimentalApi: opt-in`; tanpa header, server mengembalikan HTTP 400 dengan pesan yang menyebut header wajib.
  8. Autentikasi mendukung Basic dengan API token (API-008), OAuth 2.0 3LO dengan scope Service Management (`read:servicedesk-request`, `write:servicedesk-request`, `manage:servicedesk-customer`) (API-009, API-010), dan autentikasi app (API-013, API-015); scope tidak memadai mengembalikan HTTP 403.
  9. Rate limit berlaku per pengguna/app (API-006); permintaan melewati batas mengembalikan HTTP 429 dengan header `Retry-After` dalam detik.
  10. Keluarga endpoint service desk, request type, dan request type fields dispesifikasikan di JSM-AI-026.
  11. Keluarga endpoint customer request (create/read, participants, approvals, SLA, status, komentar public/internal, lampiran sementara, feedback, notification) dispesifikasikan di JSM-AI-025.
  12. Pembuatan request atas nama customer (`raiseOnBehalfOf`) dispesifikasikan di JSM-AI-025.
  13. Keluarga endpoint customer dan organization (termasuk endpoint experimental) dispesifikasikan di JSM-AI-028.
  14. Endpoint queue dispesifikasikan di JSM-AI-027 dan pencarian knowledge base di JSM-AI-029.
  15. API Assets (objects, schemas, AQL, imports) adalah API terpisah dengan base path sendiri (JSM-AI-031).
  16. API Operations (alerts, schedules, on-call, heartbeats, integrations) adalah API terpisah (JSM-AI-032).
  17. ⚠ Event webhook untuk request Service Management (dibuat, dikomentari, status berubah, approval) dikirim lewat mekanisme webhook umum dengan atribut tambahan request type dan visibilitas komentar (JSM-AI-033, API-065).
  18. Dokumentasi Swagger untuk seluruh endpoint Service Management dihasilkan via `make swag` dan diberi tag `ServiceDesk`.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); `GET /servicedesk` hanya mengembalikan service project company-managed, dan project team-managed hasil impor sudah dikonversi menjadi company-managed.
- **Catatan & ketergantungan:** Konvensi API publik di [23](23-api-publik-webhook.md); REST API help center/customer di JSM-PRT-050. Keputusan terbuka: apakah disediakan alias path kompatibel `/rest/servicedeskapi` agar klien yang sudah ada dapat dipakai tanpa modifikasi.

### JSM-AI-025 JSM REST API - Requests

- **Deskripsi:** Endpoint pengelolaan customer request beserta komentar, lampiran, approval, feedback, notifikasi, participant, SLA, status, dan transisi customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API sebagai agent atau customer; customer hanya mengakses request miliknya atau yang dibagikan kepadanya.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /request` menerima filter `searchTerm`, `requestOwnership` (`OWNED_REQUESTS`, `PARTICIPATED_REQUESTS`, `ORGANIZATION`, `ALL_REQUESTS`), `requestStatus` (`OPEN_REQUESTS`, `CLOSED_REQUESTS`, `ALL_REQUESTS`), `approvalStatus`, `organizationId`, `serviceDeskId`, `requestTypeId`, dan `expand`.
  2. `POST /request` menerima `serviceDeskId`, `requestTypeId`, `requestFieldValues`, `requestParticipants`, `raiseOnBehalfOf`, `channel`, `form` (answers), dan `isAdfRequest`; field wajib yang hilang mengembalikan HTTP 400 dengan daftar field per nama; `raiseOnBehalfOf` hanya diizinkan untuk agent.
  3. `GET /request/{issueIdOrKey}` mengembalikan request dengan field yang terlihat customer (bukan field internal) dan `_links` ke portal dan agent view.
  4. `GET/POST /request/{id}/comment` dengan flag `public`; customer hanya dapat membuat komentar public dan hanya melihat komentar public, `GET /comment/{commentId}` mengembalikan satu komentar, dan lampiran komentar dapat dibaca via `/comment/{id}/attachment`.
  5. `GET/POST /request/{id}/attachment` menambah lampiran dari `temporaryAttachmentIds` dengan flag `public` dan komentar opsional; konten dan thumbnail dapat diunduh lewat tautan yang dihormati permission.
  6. `GET /request/{id}/approval`, `GET /approval/{approvalId}`, dan `POST /approval/{approvalId}` dengan `decision` `approve` atau `decline`; pemanggil yang bukan approver mendapat HTTP 403, dan keputusan pada approval yang sudah final mendapat HTTP 409.
  7. `GET/POST/DELETE /request/{id}/feedback` membaca, mengirim (rating 1-5 dan komentar), dan menghapus CSAT; mengirim feedback hanya diizinkan untuk reporter setelah request resolved.
  8. `GET/PUT/DELETE /request/{id}/notification` membaca, berlangganan, dan berhenti berlangganan notifikasi request untuk pemanggil.
  9. `GET/POST/DELETE /request/{id}/participant` mengelola participant sesuai aturan berbagi customer permissions (JSM-AI-002); penambahan yang dilarang aturan mengembalikan HTTP 403.
  10. `GET /request/{id}/sla` dan `GET /request/{id}/sla/{slaMetricId}` mengembalikan siklus SLA (`ongoingCycle`, `completedCycles`, `breached`, `remainingTime`, `goalDuration`) untuk agent; customer mendapat HTTP 403.
  11. `GET /request/{id}/status` mengembalikan riwayat status customer-facing dengan waktu perubahan.
  12. `GET /request/{id}/transition` mengembalikan transisi customer yang tersedia (JSM-PRT-028) dan `POST /request/{id}/transition` menjalankannya dengan komentar opsional; transisi yang tidak tersedia mengembalikan HTTP 400.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request via portal di JSM-PRT-024 s.d. JSM-PRT-032; forms di JSM-PRJ-034.

### JSM-AI-026 JSM REST API - Request Types & Fields

- **Deskripsi:** Endpoint request type, request type group, field request type, dan entity property request type.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API; create/delete memerlukan Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /requesttype` mengembalikan request type lintas service desk yang dapat diakses, dengan filter `searchQuery` (nama) dan `serviceDeskId`.
  2. `GET /servicedesk/{id}/requesttype` mendaftar request type service desk; `POST` (experimental) membuat request type baru dengan `issueTypeId`, `name`, `description`, `helpText`; `GET/DELETE /servicedesk/{id}/requesttype/{requestTypeId}` membaca dan menghapus request type.
  3. `GET /servicedesk/{id}/requesttype/{requestTypeId}/field` mengembalikan field beserta `required`, `validValues`, `defaultValues`, serta flag `canRaiseOnBehalfOf` dan `canAddRequestParticipants` untuk pemanggil.
  4. Entity property request type: `GET .../requesttype/{id}/property` mendaftar key, dan `GET/PUT/DELETE .../property/{propertyKey}` membaca, menulis, dan menghapus nilai.
  5. ⚠ Nilai property adalah JSON valid dengan batas ukuran per property yang dapat dikonfigurasi (default 32 KB); nilai melebihi batas mengembalikan HTTP 413, property tidak tampil di UI maupun portal, dan hanya Admin Proyek atau app berizin yang dapat menulis.
  6. `GET /servicedesk/{id}/requesttypegroup` mengembalikan portal group dengan id dan nama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request type di JSM-PRJ-020 s.d. JSM-PRJ-030; entity properties umum di API-025.

### JSM-AI-027 JSM REST API - Queues

- **Deskripsi:** Endpoint daftar queue dan isinya.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API sebagai agent.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /servicedesk/{id}/queue` mendaftar queue; parameter `includeCount=true` menambahkan jumlah work item per queue.
  2. `GET /servicedesk/{id}/queue/{queueId}` mengembalikan detail queue (nama, JQL, field kolom).
  3. `GET /servicedesk/{id}/queue/{queueId}/issue` mengembalikan work item di queue sesuai urutan queue dengan pagination.
  4. Seluruh endpoint queue hanya dapat diakses agent dengan akses service desk; customer dan collaborator mendapat HTTP 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Queues di JSM-QSK-001 s.d. JSM-QSK-009.

### JSM-AI-028 JSM REST API - Customers & Organizations

- **Deskripsi:** Endpoint pengelolaan customer, organization, keanggotaan, dan entity property organization.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API; pembuatan customer/organization memerlukan hak sesuai JSM-AI-002 dan JSM-AI-006.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /customer` membuat customer portal-only dengan `email` dan `displayName`; email yang sudah terdaftar mengembalikan HTTP 409.
  2. `GET/POST/DELETE /servicedesk/{id}/customer` mendaftar, menambah (`accountIds`), dan mengeluarkan customer dari service desk.
  3. `GET/POST /organization` mendaftar dan membuat organization (nama unik per situs), serta `GET/DELETE /organization/{organizationId}` membaca dan menghapusnya.
  4. `GET/POST/DELETE /organization/{id}/user` mendaftar, menambah, dan mengeluarkan anggota organization.
  5. ⚠ Entity property organization: `GET .../organization/{id}/property` mendaftar key dan `GET/PUT/DELETE .../property/{key}` mengelola nilai JSON dengan batas ukuran yang sama seperti property request type; property tidak tampil di UI/portal.
  6. `GET/POST/DELETE /servicedesk/{id}/organization` mendaftar, menghubungkan, dan memutus organization dari service desk.
  7. ⚠ Endpoint detail customer/organization (CSM) membaca dan menulis detail fields (JSM-AI-062) dan ditandai experimental.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Organizations di JSM-PRT-034; detail profil di JSM-PRT-035.

### JSM-AI-029 JSM REST API - Knowledge Base

- **Deskripsi:** Endpoint pencarian artikel knowledge base.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API sebagai agent atau customer.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /knowledgebase/article?query=<teks>&highlight=<bool>` mencari artikel lintas KB yang dapat diakses; `highlight=true` membungkus kata yang cocok dengan penanda highlight.
  2. `GET /servicedesk/{id}/knowledgebase/article` mencari artikel dalam KB service desk tertentu.
  3. Hasil difilter sesuai permission pemanggil terhadap ruang KB (internal vs eksternal, JSM-QSK-032); artikel tanpa akses tidak pernah dikembalikan.
  4. Respons hanya memuat judul, excerpt, dan URL; konten penuh diambil lewat API modul knowledge base/wiki.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Knowledge base di JSM-QSK-031 s.d. JSM-QSK-039.

### JSM-AI-030 Forms REST API

- **Deskripsi:** API untuk forms Service Management (padanan forms eks ProForma).
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API; hak sesuai permission work item.
- **Aturan perilaku (acceptance criteria):**
  1. Base path Forms API adalah `/api/v1/forms` (padanan `api.atlassian.com/jira/forms/cloud/{cloudId}`) dan memakai autentikasi yang sama dengan JSM-AI-024.
  2. Endpoint mendaftar form pada work item, melampirkan form dari template, dan menghapus form dari work item.
  3. Endpoint membaca dan menyimpan jawaban, men-submit form, membuka kembali form ter-submit (reopen), dan mengubah visibilitas internal/external; submit form yang gagal validasi mengembalikan HTTP 400 dengan daftar pertanyaan bermasalah.
  4. Endpoint mendaftar form template per project.
  5. ⚠ Endpoint ekspor form ke PDF dan XLSX mengembalikan file yang sama dengan ekspor UI (JSM-PRJ-044).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Forms di JSM-PRJ-034 s.d. JSM-PRJ-044.

### JSM-AI-031 Assets REST API

- **Deskripsi:** API CRUD objek, schema, tipe objek, atribut, dan import Assets. Entri ini hanya mendaftarkan Assets REST API dalam keluarga API Service Management; spesifikasi lengkap dimiliki JSM-AST-045 Assets REST API ([33](33-jsm-assets-cmdb.md)).
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Integrator/Developer API; hak sesuai role object schema (JSM-AST-002).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /assets/workspace` (di bawah base Service Management) mengembalikan `workspaceId` situs.
  2. Base API Assets adalah `/api/v1/assets/workspace/{workspaceId}/v1` dengan satu path kompatibel `/jsm/assets/workspace/{workspaceId}/v1` (JSM-AST-045 butir 1).
  3. `POST /object/aql` menjalankan AQL dengan pagination `startAt`/`maxResults`; AQL tidak valid mengembalikan HTTP 400 dengan posisi galat.
  4. Endpoint objek mendukung create, read, update, delete, riwayat (`/object/{id}/history`), dan `referenceinfo`.
  5. Endpoint `objectschema`, `objecttype`, dan `objecttypeattribute` mengelola struktur schema.
  6. Endpoint status types, icons, dan config global tersedia.
  7. Imports API mendukung start import, submit data chunks, dan pengecekan `executionstatus` (JSM-AST-036).
  8. Autentikasi memakai API token atau OAuth 2.0 dengan scope Assets (`read:assets`, `write:assets`, `delete:assets`) yang terpisah dari scope Service Management (JSM-AST-045 butir 2).
  9. Endpoint IQL lama (`/iql`) dilayani sebagai alias deprecated dengan header `Deprecation` (JSM-AST-045 butir 8); dokumentasi mencantumkan pemetaan IQL → AQL.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik: JSM-AST-045 Assets REST API ([33](33-jsm-assets-cmdb.md)) adalah sumber kebenaran daftar endpoint, autentikasi, dan kode galat; butir di atas adalah ringkasan, tidak menambah endpoint di luar JSM-AST-045, dan bila berbeda JSM-AST-045 yang berlaku. Detail Assets lain di [33](33-jsm-assets-cmdb.md).

### JSM-AI-032 JSM Operations REST API

- **Deskripsi:** API untuk alert, team, schedule, eskalasi, integrasi, dan heartbeat Operations.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Integrator/Developer API, Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Base API Operations adalah `/api/v1/ops` (padanan `api.atlassian.com/jsm/ops/api/{cloudId}/v1`).
  2. Endpoint alert mendukung create, acknowledge, close, escalate, assign, snooze, notes, dan tags.
  3. Endpoint schedules, rotations, overrides, dan on-calls tersedia.
  4. Endpoint escalations dan routing rules tersedia.
  5. Endpoint integrations dan integration API key, serta endpoint incoming integration, tersedia.
  6. Endpoint heartbeats, maintenance, notification rules, dan contacts tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Spesifikasi rinci (autentikasi, pagination, rate limit, kebijakan API key) di JSM-OPS-045 ([34](34-jsm-operations-alerting-on-call.md)), yang menjadi sumber kebenaran; prioritas fitur ini disamakan dengan JSM-OPS-045 (P1) karena keduanya adalah API yang sama. Fitur ini hanya memastikan API Operations terdaftar dalam keluarga API Service Management (JSM-AI-024) dan tidak menambah endpoint di luar JSM-OPS-045. Pasangan tumpang tindih dua arah: JSM-OPS-045 Operations REST API ([34](34-jsm-operations-alerting-on-call.md)) ↔ fitur ini.

### JSM-AI-033 JSM Extensibility (Forge, Connect & Webhooks)

- **Deskripsi:** Extension point platform khusus Service Management untuk app (portal, queue, organization, Assets import) dan penerapan webhook umum pada request.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / platform
- **Aktor & izin:** Integrator/Developer API (pengembang app), Admin Situs (memasang app).
- **Aturan perilaku (acceptance criteria):**
  1. Extension point portal request: `portalRequestDetail`, `portalRequestDetailPanel`, `portalRequestViewAction`, dan `portalRequestCreatePropertyPanel` merender UI app di halaman request portal.
  2. Extension point portal layout: `portalHeader`, `portalSubheader`, `portalFooter`, `portalProfilePanel`, dan `portalUserMenuAction`.
  3. Extension point agent/admin: `queuePage`, `organizationPanel`, dan `assetsImportType`.
  4. Modul legacy (padanan Connect: serviceDeskPortalHeaders/Footers/SubHeaders, serviceDeskQueues, serviceDeskReports, serviceDeskOrganizationPanels, serviceDeskPortalRequestViewActions, serviceDeskPortalProfileActions, serviceDeskPortalUserMenuActions) tidak dibangun sebagai model terpisah; semua dipetakan ke extension point runtime app tunggal (INT-068).
  5. Webhook umum (event issue dan comment) berlaku untuk request Service Management, dengan payload yang menyertakan request type dan flag visibilitas komentar.
  6. ⚠ Extension point portal menerima konteks customer (akun customer, service desk, request) melalui bridge runtime app, tanpa mengekspos token agent.
  7. App dapat menyimpan konfigurasi di entity property request type dan organization (JSM-AI-026, JSM-AI-028).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Strategi ekstensi INT-060, INT-068 s.d. INT-080 ([25](25-integrasi-impor-ekspor-ekstensi.md)). Connector pihak ketiga (mis. connector katalog layanan cloud provider) adalah produk vendor, bukan fitur native.

### JSM-AI-034 Opsgenie to JSM Operations Migration

- **Deskripsi:** Migrasi pengguna produk alerting lama ekosistem rujukan ke Operations Service Management. Produk ini tidak memiliki produk lama tersebut, sehingga fitur dicatat sebagai rujukan dan tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / jsm
- **Aktor & izin:** Org admin, Ops admin.
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan: produk lama berstatus end of sale Juni 2025 dan end of support April 2027; tidak ada jadwal sunset di produk ini.
  2. Rujukan: data yang dimigrasi di ekosistem rujukan (teams, schedules, escalations, integrations, alerts) dapat diimpor ke produk ini lewat REST API Operations (JSM-AI-032, JSM-OPS-045).
  3. Dokumentasi pemetaan endpoint API lama ke `/api/v1/ops` disediakan sebagai panduan bagi integrator.
  4. ⚠ Tidak ada halaman pengaturan migrasi di admin.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konsisten dengan JSM-OPS-047 ([34](34-jsm-operations-alerting-on-call.md)); lihat bagian Di luar cakupan.

## Produktivitas agent & kapabilitas AI

Prinsip sub-area ini: semua fitur AI memanggil antarmuka `AIProvider` yang dapat diganti (model lokal, penyedia komersial, atau tidak ada). Bila tidak ada provider terkonfigurasi, admin mematikan AI di pengaturan AI pusat (ADM-052) atau di toggle Service Management (JSM-AI-036), atau situs memakai profil regulated (VISI-003), semua tombol/panel AI disembunyikan dan endpoint AI mengembalikan HTTP 403 `ai_disabled` (sama dengan ADM-052). Prompt ke provider hanya boleh berisi konten yang dapat dilihat pengguna pemicu.

### JSM-AI-035 Canned Responses

- **Deskripsi:** Respons tersimpan yang dapat disisipkan agent ke reply to customer atau internal note, termasuk variabel dan saran berbasis AI.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (membuat personal, memakai), Admin Proyek (mengelola shared).
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat membuat, mengedit, dan menghapus canned response langsung dari menu canned responses di editor komentar; menghapus meminta konfirmasi.
  2. Saat membalas, agent dapat mencari canned response berdasarkan judul (pencocokan sebagian, case-insensitive) dan menyisipkannya ke posisi kursor; konten tetap dapat diedit sebelum dikirim.
  3. ⚠ Visibilitas canned response adalah Personal (hanya pembuat) atau Shared (semua agent di service project); hanya Admin Proyek yang dapat membuat atau mengubah ke Shared.
  4. ⚠ Konten mendukung variabel (mis. `{{reporter.name}}`, `{{issue.key}}`, `{{portal.url}}`) yang terisi saat disisipkan; variabel yang tidak dikenal dibiarkan apa adanya.
  5. ⚠ Bila AI aktif, menu canned responses menampilkan hingga 3 saran canned response yang paling relevan dengan isi request di bagian atas daftar, dengan label "Disarankan".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Spesifikasi dasar (batas jumlah, signature) di JSM-QSK-019 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)); fitur ini menambahkan aturan variabel dan saran AI. Implementasi harus memakai satu entitas `CannedResponse` yang sama dengan JSM-QSK-019; aturan batas jumlah, scope per project, dan izin Shared mengikuti JSM-QSK-019 bila ada perbedaan. Rujukan balik dari JSM-QSK-019 ke fitur ini perlu ditambahkan di modul 31.

### JSM-AI-036 Atlassian Intelligence / Rovo Administration for JSM

- **Alias inventaris:** AI features in JSM (Atlassian Intelligence / Rovo)
- **Deskripsi:** Perluasan khusus Service Management atas pengaturan AI pusat (ADM-052): toggle granular fitur AI Service Management, tata kelola, dan katalog kapabilitas AI (padanan Atlassian Intelligence/Rovo). Heading ini menggabungkan entri katalog fitur AI dari lampiran Service Management.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Org admin/Admin Situs menyalakan AI situs dan app Service Management di ADM-052 (satu-satunya sakelar utama); Admin Situs mengatur toggle granular fitur AI Service Management di halaman yang sama. Admin Proyek tidak dapat menyalakan AI yang dimatikan di tingkat situs.
- **Aturan perilaku (acceptance criteria):**
  1. Sakelar utama AI per situs dan per app (Work, Service Management, Product Discovery) hanya ada di pengaturan AI pusat (ADM-052) dan dioperasikan Org admin/Admin Situs; default untuk instalasi baru adalah nonaktif. Fitur ini tidak memiliki sakelar utama sendiri: toggle granular Service Management (butir 17) hanya dapat diubah bila AI app Service Management aktif di ADM-052, dan pada profil regulated (VISI-003) semuanya terkunci OFF dan read-only.
  2. Mematikan AI menyembunyikan chat AI dan agen AI; fitur inti non-AI (pencarian kata kunci, JQL, queue) tetap berjalan tanpa degradasi.
  3. ⚠ Ketersediaan AI tidak digating plan/tier (mengikuti ADM-052 butir 14); kapabilitas AI Service Management hanya bergantung pada konfigurasi penyedia LLM, sakelar di ADM-052, dan toggle granular fitur ini. Pembedaan AI dasar vs kapabilitas penuh (chat, agents, builder agen) per plan di ekosistem rujukan hanya dicatat sebagai rujukan.
  4. Data situs tidak dipakai untuk melatih model publik; bila provider eksternal dipakai, halaman admin menampilkan nama provider, region pemrosesan, dan kebijakan retensi yang dikonfigurasi operator.
  5. Setiap aksi AI menghormati permission pengguna: konteks yang dikirim ke model hanya memuat konten yang dapat diakses pengguna pemicu, dan hasilnya tidak pernah memuat konten dari work item yang tidak dapat dilihatnya.
  6. ⚠ Pemakaian AI Service Management diukur memakai kuota/kredit yang sama dengan ADM-052 (butir 6-11); halaman pemakaian AI di ADM-052 tetap menjadi sumber kebenaran, Track Feature Usage (JSM-AI-012) menampilkan ringkasan pemakaian app Service Management, dan saat kuota habis fitur AI Service Management dibatasi dengan pesan "Kuota AI habis".
  7. ⚠ Region pemrosesan AI dapat dibatasi agar mengikuti data residency situs (JSM-AI-011); bila provider tidak mendukung region tersebut, AI tidak dapat diaktifkan.
  8. Audit log mencatat aktivasi, penonaktifan, perubahan provider, dan perubahan toggle fitur AI dengan aktor dan waktu.
  9. Katalog fitur AI Service Management mencakup ringkasan request dan thread komentar panjang (JSM-AI-037).
  10. Katalog mencakup bantuan menulis balasan: ubah nada, perbaiki, ringkas (JSM-AI-038).
  11. ⚠ Katalog mencakup triage AI: saran request type, priority, dan field (JSM-AI-041).
  12. Katalog mencakup similar requests untuk menyelesaikan atau menautkan (JSM-AI-039).
  13. Katalog mencakup AI answers di help center/portal dan Virtual Service Agent (JSM-AI-052).
  14. Katalog mencakup bahasa alami ke JQL dan ke AQL (JSM-AI-039).
  15. Katalog mencakup AIOps: alert grouping, ringkasan incident, suggested responders, saran root cause, draf PIR (JSM-AI-045).
  16. ⚠ Katalog mencakup agen AI untuk service desk (JSM-AI-043, JSM-AI-044).
  17. Selain sakelar utama di ADM-052, Admin Situs dapat mematikan fitur AI Service Management tertentu secara granular (mis. sentiment analysis aktif, writing assistance mati); toggle granular ditampilkan sebagai bagian Service Management di halaman ADM-052, bukan halaman terpisah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Prasyarat keras: ADM-052 ([22](22-administrasi-sistem.md), P3 / Fase 4) adalah kontrol pusat satu-satunya untuk menyalakan AI; fitur ini hanya menambah toggle granular, katalog, dan tata kelola khusus Service Management, sehingga prioritasnya P2 (setara fitur AI Service Management yang digatingnya) dan ADM-052 harus dibangun lebih dulu. Daftar fitur yang digating ADM-052 perlu dilengkapi dengan fitur AI modul ini (JSM-AI-036 s.d. JSM-AI-058 dan JSM-AI-064) di modul 22. Katalog fitur AI Work di NAV-052 ([21](21-navigasi-ux-preferensi.md)); izin agen AI di PERM-050.

### JSM-AI-037 AI Work Item Summaries

- **Deskripsi:** Ringkasan otomatis request (deskripsi dan komentar) untuk agent.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dengan akses work item; AI aktif.
- **Aturan perilaku (acceptance criteria):**
  1. Request view agent menyediakan aksi "Summarize" yang tampil bila request memiliki deskripsi atau minimal satu komentar.
  2. Ringkasan memuat tiga bagian: inti masalah, langkah yang sudah dilakukan, dan status/keputusan terakhir.
  3. Ringkasan hanya memakai konten yang dapat dilihat agent (termasuk internal note bila agent berhak melihatnya).
  4. Agent dapat memberi feedback jempol naik/turun per ringkasan; feedback disimpan untuk metrik kualitas tanpa menyimpan isi ringkasan.
  5. Ringkasan tidak disimpan otomatis; tombol "Salin ke internal note" menyisipkannya ke editor internal note untuk diedit sebelum disimpan.
  6. Ringkasan alert/incident Operations memakai mekanisme yang sama (JSM-AI-045, JSM-OPS-042).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Overlap dengan JSM-QSK-020 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)) dan ISS-045 ([06](06-work-item-inti.md)).

### JSM-AI-038 AI Writing Assistance for Agents

- **Deskripsi:** Bantuan menulis balasan agent berbasis AI.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan; AI aktif.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Aksi "Suggest a response" menyusun draf balasan dari konteks request dan artikel KB terkait, dengan tautan artikel sumber bila dipakai.
  2. Aksi "Improve writing", "Fix spelling & grammar", "Make shorter", dan "Make longer" mengubah teks terpilih atau seluruh draf.
  3. ⚠ Aksi "Change tone" menyediakan pilihan nada minimal professional, empathetic, dan casual.
  4. Hasil AI hanya mengisi editor; agent tetap harus memilih secara manual "Reply to customer" atau "Add internal note" dan menekan kirim, dan AI tidak pernah mengirim balasan sendiri.
  5. Agent dapat membatalkan (undo) atau membuang hasil AI untuk kembali ke teks sebelumnya, serta memberi feedback jempol naik/turun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AI di editor EDT-030 ([07](07-editor-konten-kaya.md)).

### JSM-AI-039 AI Agent Assist & Natural Language Search

- **Deskripsi:** Bantuan konteks bagi agent: work item serupa, artikel relevan, chat AI dalam konteks request, serta pencarian bahasa alami ke JQL dan AQL.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder incident; AI aktif.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Panel "Similar work items" di issue view menampilkan hingga 5 item dengan kemiripan semantik tertinggi, masing-masing dengan ringkasan, status, dan resolusi.
  2. ⚠ Dari panel, agent dapat membuka item, menautkannya sebagai duplicate atau related, dan menyalin solusi ke editor.
  3. Untuk incident, panel menampilkan incident serupa dan alert terkait sebagai bantuan triage (JSM-AI-045).
  4. ⚠ Panel hanya muncul bila AI aktif untuk Service Management dan fitur "Find similar work items" diaktifkan di Feature Lab space (JSM-AI-046).
  5. Panel hanya menampilkan item yang dapat dilihat agent; item tanpa akses tidak dihitung maupun disebut.
  6. Agent dapat memberi feedback relevansi (jempol naik/turun) per hasil.
  7. Panel menyarankan artikel KB yang relevan dan agent dapat membagikannya ke customer sebagai tautan di reply.
  8. Chat AI dalam konteks request menjawab pertanyaan agent, mencari work item historis, artikel KB, dan pakar terkait (NAV-051).
  9. Pencarian work item dan queue menerima kueri bahasa alami yang diterjemahkan ke JQL; JQL hasil terjemahan ditampilkan dan dapat diedit sebelum dijalankan (JQL-059).
  10. Pencarian Assets menerima kueri bahasa alami yang diterjemahkan ke AQL dengan perilaku yang sama.
  11. Deteksi sentimen customer dispesifikasikan di JSM-AI-042.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Indeks embedding memerlukan komponen pencarian vektor (keputusan arsitektur di modul 27).

### JSM-AI-040 AI Knowledge Article Drafting

- **Deskripsi:** Pembuatan draf artikel knowledge base dari request yang selesai atau dari prompt.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, knowledge manager; memerlukan izin create di ruang KB.
- **Aturan perilaku (acceptance criteria):**
  1. Request view menyediakan aksi "Create article" yang menghasilkan draf AI dari deskripsi, komentar public, dan resolusi request.
  2. Agent memilih ruang KB tujuan dari daftar ruang yang terhubung ke service project.
  3. Draf disimpan sebagai draft (belum terbit) di ruang KB dan harus direview sebelum dipublikasikan; draf ditautkan ke request asal.
  4. Editor artikel KB menyediakan aksi generate artikel dari prompt bebas.
  5. Agent tanpa izin create di ruang KB tidak melihat aksi tersebut; panggilan API langsung mengembalikan HTTP 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Create article manual di JSM-QSK-037.

### JSM-AI-041 AI Request Type & Space Setup Suggestions

- **Deskripsi:** Saran AI saat membuat service project dan request type, serta klasifikasi request masuk.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (setup), Agent Layanan (triage).
- **Aturan perilaku (acceptance criteria):**
  1. Saat membuat service project, admin dapat menulis deskripsi tim dan menerima daftar request type yang disarankan.
  2. AI dapat menghasilkan request type lengkap: nama, deskripsi, dan field yang disarankan.
  3. Setiap saran dapat diterima, diubah, atau ditolak secara individual; tidak ada request type yang dibuat tanpa persetujuan admin.
  4. ⚠ Triage AI menyarankan request type, priority, dan nilai field untuk request masuk; agent menerima saran dengan satu klik, dan saran tidak pernah diterapkan otomatis kecuali lewat rule otomasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Wizard service project di JSM-PRJ-001; triage manual di JSM-QSK-016.

### JSM-AI-042 Customer Sentiment Analysis

- **Deskripsi:** AI menganalisis nada pesan customer dan mengisi field Sentiment (Positive/Neutral/Negative) agar agent dapat memprioritaskan request.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (mengaktifkan), Agent Layanan (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Toggle "Customer sentiment analysis" tersedia di Project settings > Features > Work item view; default nonaktif.
  2. Field Sentiment bernilai Positive, Neutral, atau Negative dan tampil di panel Details work item dengan ikon warna.
  3. Setelah fitur aktif, work item baru langsung dihitung sentiment-nya saat dibuat.
  4. Work item yang dibuat sebelum fitur aktif baru dihitung saat reporter/customer menambahkan komentar baru.
  5. Sentiment dihitung ulang setiap kali customer menambahkan komentar; komentar agent dan internal note tidak memicu perhitungan.
  6. Field Sentiment tampil di issue view dan dapat dipilih sebagai kolom queue (JSM-QSK-003).
  7. Fitur hanya dapat diaktifkan bila AI aktif untuk Service Management; bila AI dimatikan, nilai lama tetap tersimpan tetapi tidak diperbarui.
  8. ⚠ Fitur tidak tersedia pada deployment yang ditandai lingkungan pemerintahan/terisolasi tanpa provider AI.
  9. ⚠ Field Sentiment dapat dipakai di JQL (JSM-AI-013), filter queue, dan kondisi otomasi untuk eskalasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur project di JSM-PRJ-016.

### JSM-AI-043 Rovo Agents for Service Management

- **Deskripsi:** Agen AI siap pakai untuk alur service management (padanan Rovo agents) dan integrasinya dengan otomasi.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek; izin agen AI mengikuti PERM-050.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tersedia agen AI bawaan untuk service management (mis. pembantu triage service request dan agen operasi) yang dapat dipanggil dari issue view.
  2. Rule otomasi dapat memanggil agen AI pada work item dengan prompt dan menyimpan keluarannya ke komentar internal atau field (AUTO-073).
  3. Admin dapat membuat agen kustom memakai builder agen dengan konteks Service Management (request, KB, customer) sebagai sumber.
  4. ⚠ Agen AI self-service untuk customer dapat diaktifkan di portal dalam status beta yang dikendalikan flag.
  5. ⚠ Agen deflection portal di help center dibedakan dari Request Resolver (JSM-AI-044) dan Virtual Service Agent (JSM-AI-047): ia hanya menjawab, tidak mengeksekusi aksi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Agen AI umum di NAV-051, ISS-046, WF-043.

### JSM-AI-044 Request Resolver

- **Deskripsi:** Rekan kerja AI yang, bila diaktifkan untuk request type service request, menyusun rencana penyelesaian work item lalu menjalankannya secara otonom atau dengan supervisi agent.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (konfigurasi), Agent Layanan (supervisi).
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih request type mana yang ditangani Request Resolver; request type lain tidak terpengaruh.
  2. Execution mode diatur per request type: Supervised (plan harus disetujui agent sebelum dieksekusi) atau Autonomous; default Supervised.
  3. Admin memilih ruang KB sebagai knowledge source; admin hanya dapat memilih ruang yang dapat ia akses.
  4. Request Resolver membuat plan per work item yang tampil di panel "Work with Request Resolver"; agent dapat mengedit langkah, menyetujui, atau menolak plan.
  5. Eksekusi plan dapat mencakup aksi di work item (edit field, assign), balasan ke customer, dan transisi status; setiap aksi dicatat di riwayat work item dengan aktor "Request Resolver".
  6. Halaman konfigurasi menampilkan panduan rollout: mulai dari request type bervolume tinggi dan terdokumentasi baik, naik ke Autonomous setelah plan konsisten disetujui tanpa edit (metrik persentase plan tanpa edit ditampilkan).
  7. Request Resolver hanya dapat diaktifkan untuk request type kategori service request (bukan incident, change, problem) dan tidak melakukan deflection di help center.
  8. Halaman aktivasi menampilkan peringatan bahwa fitur tidak boleh dipakai untuk data kesehatan terlindungi (PHI) atau data yang memerlukan kepatuhan khusus.
  9. Setiap plan dan eksekusi mengonsumsi kredit AI (JSM-AI-036) dan tercatat di Track Feature Usage.
  10. Halaman "What Request Resolver can resolve" menjelaskan cakupan aksi yang didukung dan batasannya.
  11. ⚠ Request Resolver dapat dipakai dalam langkah journeys onboarding (JSM-ITSM-050).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari Virtual Service Agent (customer-facing) dan agen AI JSM-AI-043; keputusan terbuka mengenai status rilis.

### JSM-AI-045 AIOps for Incidents

- **Deskripsi:** Fitur AI pada incident dan alert management.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Responder, Incident manager (Manajer Layanan/ITSM).
- **Aturan perilaku (acceptance criteria):**
  1. Alert serupa dikelompokkan otomatis berbasis AI (spesifikasi di JSM-OPS-042).
  2. Incident view menyarankan responder dan menampilkan incident/alert terkait.
  3. ⚠ Panel incident serupa di issue view incident membantu triage dengan menampilkan resolusi sebelumnya.
  4. Ringkasan AI incident/alert tersedia di incident view dan alert detail.
  5. Draf post-incident review dapat dibuat dari timeline incident dan harus direview sebelum disimpan (JSM-ITSM-032).
  6. ⚠ AI menyarankan kemungkinan root cause dan perubahan/deployment terkait dari data change dan deployment (JSM-ITSM-028).
  7. Incident dapat dibuat dari alert group dengan seluruh alert anggota ditautkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pasangan tumpang tindih: JSM-ITSM-030 AI Assistance for Incidents ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)) dan JSM-OPS-042 ([34](34-jsm-operations-alerting-on-call.md)). Pemilik: fitur ini memiliki kapabilitas AIOps incident/alert (saran responder, incident serupa, ringkasan, draf PIR, saran root cause, incident dari alert group) di atas antarmuka `AIProvider` (JSM-AI-036); JSM-ITSM-030 memiliki aksi AI di incident view (Find related resources, Investigate with AI, asisten chat) dan wajib memakai ringkasan serta analisis root cause yang sama dari fitur ini; JSM-OPS-042 memiliki spesifikasi grouping alert (butir 1).

### JSM-AI-046 JSM Feature Lab

- **Deskripsi:** Halaman di pengaturan service project untuk mengaktifkan fitur eksperimen/opsional per service project.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Feature Lab menampilkan daftar fitur eksperimen/opsional dengan nama, deskripsi, dan status (eksperimen/beta).
  2. Setiap fitur memiliki toggle aktif/nonaktif per service project; perubahan berlaku tanpa memuat ulang aplikasi dan tercatat di audit log.
  3. Contoh fitur yang tersedia: "Find similar work items" (JSM-AI-039).
  4. Contoh fitur yang tersedia: menambahkan Services, Alerts, dan On-call ke sidebar service project.
  5. Hanya Admin Proyek yang dapat mengubah toggle; agent melihat halaman read-only atau tidak melihatnya sama sekali.
  6. Isi daftar dikendalikan oleh registry feature flag; fitur yang lulus ke GA hilang dari daftar dan aktif permanen, fitur yang dihapus hilang dan state-nya dibersihkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur project di JSM-PRJ-016.


## Virtual Service Agent

### JSM-AI-047 Virtual Service Agent - Setup & Activation

- **Alias inventaris:** Virtual service agent
- **Deskripsi:** Virtual agent percakapan per service project yang menjawab customer secara otomatis lewat intent flow dan AI answers, lalu menyerahkan ke agent manusia bila perlu. Heading ini juga menggabungkan entri ringkas "Virtual service agent" dari lampiran Service Management; butir lampiran yang bersifat detail dirujuk ke fitur VSA rinci di bawah.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek dan Admin Situs mengonfigurasi; Customer/Pelapor berinteraksi; Agent Layanan menerima eskalasi.
- **Aturan perilaku (acceptance criteria):**
  1. Menu Project settings > Virtual service agent menampilkan halaman overview (status, channel aktif, jumlah intent, pemakaian kuota) dan wizard "Siapkan virtual service agent" untuk service project yang belum pernah dikonfigurasi.
  2. ⚠ Virtual agent hanya tersedia bila plan situs Premium/Enterprise (flag plan JSM-AI-005); AI answers hanya berfungsi bila kapabilitas AI aktif di situs (JSM-AI-036); berlaku untuk semua service project (hanya company-managed, JSM-PRJ-012).
  3. Admin dapat menetapkan atau mengubah default request type yang dipakai standard flow Escalate bila flow tidak menentukan request type; default request type wajib diisi sebelum virtual agent dapat diaktifkan, dan request type yang dihapus/disembunyikan memicu peringatan di overview.
  4. Admin dapat mematikan virtual agent seluruhnya atau per channel tanpa menghapus intent, flow, maupun konfigurasi; menyalakan kembali memulihkan konfigurasi terakhir.
  5. Mode respons: pesan customer dicocokkan ke intent lebih dulu, lalu jatuh ke AI answers bila tidak ada intent yang cocok; AI answers dapat diaktifkan/dinonaktifkan terpisah dari intent flow.
  6. ⚠ Nama tampilan virtual agent dan pesan Greet dapat dikustomisasi per service project; avatar dapat diganti dengan gambar PNG/JPG maksimal 1 MB.
  7. Hanya Admin Proyek dan Admin Situs yang dapat mengubah konfigurasi (API mengembalikan HTTP 403 untuk pengguna lain); pengaturan "siapa yang dapat melihat percakapan" menentukan apakah transkrip dapat dilihat oleh admin saja atau juga oleh agent.
  8. Bila kapabilitas AI dimatikan di pengaturan organisasi, AI answers berhenti otomatis dan hanya intent flow yang berjalan; overview menampilkan banner penyebabnya.
  9. ⚠ Bila plan turun dari Premium ke Standard, virtual agent berhenti melayani di seluruh channel, tetapi intent, flow, dan riwayat percakapan tetap disimpan read-only selama masa retensi sehingga dapat dipulihkan bila plan naik kembali.
  10. ⚠ Lampiran: virtual agent melayani channel Slack, Microsoft Teams, help center/portal, widget, dan email (rinci di JSM-AI-053).
  11. Lampiran: jawaban generatif dari knowledge base dan knowledge sources (rinci di JSM-AI-052).
  12. Lampiran: intent dengan contoh frasa, intent template, dan saran intent AI (rinci di JSM-AI-048).
  13. Lampiran: conversation flow dengan langkah message, ask question, branch/condition, raise request, dan transfer to agent (rinci di JSM-AI-049).
  14. Lampiran: test mode sebelum live (rinci di JSM-AI-055).
  15. Pengaturan sapaan, nada bahasa (formal/santai), dan pesan fallback ("Maaf, saya belum dapat membantu; permintaan Anda akan diteruskan ke tim") dapat diubah per service project.
  16. Lampiran: handover ke agent manusia dengan membuat request (rinci di JSM-AI-054).
  17. Lampiran: analitik assisted, resolved, escalated, dan kinerja per intent (rinci di JSM-AI-056).
  18. ⚠ Lampiran: batas assisted conversations per bulan per tier (rinci di JSM-AI-058).
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012).
- **Catatan & ketergantungan:** Bergantung pada request type ([29](29-jsm-service-project-request-type-formulir.md)), channel help center/widget/chat ([30](30-jsm-help-center-portal-pelanggan-kanal.md)), knowledge base dan CSAT ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)), dan antarmuka `AIProvider` (JSM-AI-036). Entitas baru `VirtualAgentConfig` di `internal/shared/domain/entity`, feature module `internal/features/virtualagent`. Rebrand ekosistem rujukan ke "agen layanan" berbasis AI dicatat di JSM-AI-064.

### JSM-AI-048 Virtual Service Agent - Intents

- **Deskripsi:** Pengelolaan intent, yaitu maksud customer yang dikenali virtual agent dari training phrases dan dipetakan ke conversation flow.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar intent menampilkan nama, status aktif/nonaktif, channel tempat intent berlaku, serta pengubah dan waktu perubahan terakhir; daftar dapat dicari berdasarkan nama.
  2. Admin dapat membuat, mengedit, dan menghapus intent dengan field nama (wajib, unik per service project), deskripsi (dipakai saat konfirmasi "Apakah maksud Anda ...?"), training phrases, dan flow; hapus memerlukan konfirmasi.
  3. ⚠ Admin melatih intent dengan menambah training phrases yang bervariasi; intent memerlukan minimal 3 dan maksimal 50 frasa, masing-masing maksimal 200 karakter; menyimpan di bawah minimum ditolak dengan pesan "Tambahkan minimal 3 contoh frasa".
  4. Intent template bawaan (mis. reset password, akses software, VPN, hardware) tersedia sesuai template service project dan dapat ditambahkan lalu diedit seperti intent biasa.
  5. Saran intent berbasis AI dihasilkan dari analisis request historis service project (kelompok request bertema sama) dan hanya tampil bila AI aktif; admin dapat menerima (menjadi intent draf) atau mengabaikan saran.
  6. Admin dapat mengaktifkan atau menonaktifkan intent; intent nonaktif tidak dipakai untuk matching di channel mana pun.
  7. ⚠ Saat menyimpan, sistem memperingatkan bila training phrase sangat mirip dengan frasa intent lain, beserta nama intent yang bertabrakan, tanpa memblokir penyimpanan.
  8. Menghapus intent tidak menghapus atau mengubah request yang sudah dibuat dari intent tersebut; referensi intent di data percakapan historis tetap ditampilkan sebagai "(intent dihapus)".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `Intent` (relasi ke `ConversationFlow`); matching dilakukan oleh standard flow Match intent (JSM-AI-051); pelatihan ulang setelah perubahan di JSM-AI-055.

### JSM-AI-049 Virtual Service Agent - Conversation Flow Builder

- **Deskripsi:** Editor visual no-code untuk menyusun alur percakapan setiap intent dalam bentuk rangkaian step dan branch.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap intent memiliki tab Flow; flow terdiri dari rangkaian step yang dapat bercabang sesuai respons customer dan ditampilkan sebagai diagram vertikal.
  2. Step "Send message" mengirim pesan dengan format dasar (tebal, miring, daftar) dan hyperlink; pesan kosong tidak dapat disimpan.
  3. Step "Offer choices" menerima 1-10 pilihan dan setiap pilihan membuat branch sendiri; 1-5 pilihan dirender sebagai tombol, 6-10 sebagai dropdown; menambah pilihan ke-11 ditolak.
  4. ⚠ Step "Ask for information" mengumpulkan jawaban ke field request; field yang didukung adalah Summary, Description, dan custom field bertipe Date, Paragraph, Short text, dan Single-select; field lain tidak muncul di pemilih.
  5. Step teknis "Change request type" (tidak terlihat oleh customer) menetapkan request type dan nilai field untuk work item yang dibuat setelah titik tersebut.
  6. Step "Send web request" tersedia di builder (perilaku rinci di JSM-AI-050).
  7. Setiap branch wajib diakhiri dengan memulai standard flow Escalate atau Resolve with CSAT; flow dengan branch tanpa akhir tidak dapat dipublikasikan dan branch bermasalah disorot.
  8. Satu conversation flow dibatasi maksimal 50 step; tombol tambah step dinonaktifkan dengan tooltip batas saat limit tercapai.
  9. Menghapus step yang bercabang hanya mempertahankan branch paling kiri; dialog konfirmasi menyebutkan jumlah step yang akan ikut terhapus.
  10. Builder menampilkan panduan "menulis pesan yang efektif" (kalimat pendek, satu pertanyaan per step, hindari jargon) sebagai panel bantuan yang dapat ditutup.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `ConversationFlow` disimpan sebagai graf step berversi (draf vs terbit); pemetaan jawaban ke field request di JSM-AI-054; field request type dari [29](29-jsm-service-project-request-type-formulir.md).

### JSM-AI-050 Virtual Service Agent - Integrations in Flows

- **Deskripsi:** Step flow yang memanggil sistem eksternal atau rule otomasi untuk mengambil data atau menjalankan aksi selama percakapan.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (menyusun step); Admin Situs (mengelola koneksi dan kredensial terpusat).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Step "Send web request" mengonfigurasi method (GET/POST/PUT/PATCH/DELETE), URL HTTPS, header, dan body dengan variabel dari jawaban customer; nilai dari respons JSON dapat dipetakan ke variabel untuk step berikutnya; kredensial header disimpan terenkripsi dan tidak ditampilkan ulang.
  2. ⚠ Flow dapat menjalankan rule otomasi yang memakai trigger khusus "Virtual agent: flow step dijalankan"; rule menerima variabel percakapan dan dapat mengembalikan nilai ke flow.
  3. Koneksi iPaaS (padanan Workato recipes) dapat dipanggil dari flow melalui konektor iPaaS generik (INT-003) bila dikonfigurasi Admin Situs.
  4. Koneksi sistem HR (padanan Workday actions, mis. data karyawan atau saldo cuti) dapat dipanggil dari flow melalui konektor yang sama; tanpa konektor terpasang, opsi ini disembunyikan.
  5. ⚠ Web request memiliki timeout 10 detik; respons non-2xx atau timeout mengarahkan ke branch "Gagal" yang wajib dikonfigurasi (default: Escalate) dan tercatat di log percakapan tanpa memuat kredensial.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mesin otomasi di [20](20-otomasi.md); konektor iPaaS di [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-003). Konektor merek pihak ketiga tertentu di luar cakupan native; yang dibangun adalah antarmuka konektor generik. Wajib proteksi SSRF (blokir alamat privat/loopback) sesuai [27](27-nfr-arsitektur-platform.md).

### JSM-AI-051 Virtual Service Agent - Standard Flows

- **Deskripsi:** Lima standard flow bawaan yang mengatur bagian umum percakapan di luar intent spesifik.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (mengedit teks/opsi standard flow).
- **Aturan perilaku (acceptance criteria):**
  1. Greet adalah sapaan opsional; mengosongkan field Message mematikan sapaan, dan pesan pertama customer tetap dianalisis untuk intent.
  2. Match intent mencocokkan pesan pertama ke intent atau AI answer; bila confidence sedang, virtual agent meminta konfirmasi memakai deskripsi intent ("Apakah maksud Anda ...?"), dan bila confidence tinggi langsung menjalankan flow intent.
  3. Clarification meminta customer menjelaskan ulang bila pesannya tidak jelas atau confidence rendah, sebelum jatuh ke Escalate.
  4. Escalate membuat work item untuk agent manusia dan memberi tahu customer bahwa request telah dibuat beserta nomornya.
  5. Resolve with CSAT mengakhiri percakapan sebagai resolved dan meminta rating kepuasan (skala CSAT modul 31).
  6. Teks dan opsi sebagian standard flow dapat diedit, tetapi standard flow tidak dapat dihapus atau dinonaktifkan sepenuhnya (kecuali Greet yang dapat dikosongkan); API menolak penghapusan dengan HTTP 400.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** CSAT di [31](31-jsm-antrian-sla-knowledge-base-laporan.md); hand-off rinci di JSM-AI-054.

### JSM-AI-052 Virtual Service Agent - AI Answers

- **Deskripsi:** Jawaban generatif dari knowledge base terhubung, dengan tautan ke artikel sumber, untuk pertanyaan yang tidak cocok dengan intent.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (mengaktifkan dan memilih sumber), Customer/Pelapor (penerima jawaban), pengelola knowledge base.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat mengaktifkan atau menonaktifkan AI answers per service project dan per channel.
  2. Sumber pengetahuan adalah space knowledge base yang terhubung ke service project (modul 31); halaman konfigurasi menautkan panduan "siapkan knowledge base" bila belum ada sumber.
  3. Hanya artikel yang dapat diakses customer sesuai permission knowledge base yang dipakai untuk menyusun jawaban; artikel terbatas tidak pernah dikutip.
  4. Setiap jawaban menyertakan tautan ke artikel sumber yang dipakai.
  5. Setelah jawaban, virtual agent menanyakan apakah jawaban membantu; "Ya" melanjutkan ke Resolve with CSAT, "Tidak" melanjutkan ke Escalate.
  6. Bila tidak ada konten relevan di atas ambang relevansi, virtual agent tidak mengarang jawaban dan langsung menjalankan Escalate.
  7. AI answers juga muncul di pencarian help center sebagai kartu jawaban di atas hasil artikel.
  8. Perubahan atau artikel baru memerlukan waktu indexing; status indexing terakhir dan jumlah artikel terindeks ditampilkan di halaman konfigurasi.
  9. ⚠ Sumber selain knowledge base bawaan (konektor pengetahuan eksternal) dapat ditambahkan melalui konektor pengetahuan INT-055 bila tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada `AIProvider` dan retrieval (embedding/index) yang dapat diganti (JSM-AI-036); knowledge base di [31](31-jsm-antrian-sla-knowledge-base-laporan.md); konektor di [25](25-integrasi-impor-ekspor-ekstensi.md).

### JSM-AI-053 Virtual Service Agent - Channels

- **Deskripsi:** Penyambungan virtual agent ke kanal portal, help center, embeddable widget, Slack, Microsoft Teams, dan email.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (mengaktifkan channel per service project), Admin Situs (memasang integrasi chat), Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Di portal dan help center, virtual agent/AI answers muncul saat customer mencari atau mulai membuat request, sebelum form request type ditampilkan.
  2. Embeddable widget menyediakan opsi "Gunakan virtual agent di widget" sehingga situs eksternal yang memasang widget (modul 30) mendapat percakapan virtual agent.
  3. Channel Slack memerlukan integrasi chat Service Management untuk Slack terhubung; virtual agent menjawab di request channel yang dikonfigurasi.
  4. Channel Microsoft Teams memerlukan integrasi chat Service Management untuk Teams terhubung; perilaku sama dengan Slack.
  5. ⚠ Channel email: virtual agent membalas email masuk ke alamat service project dengan AI answer atau flow intent; balasan customer pada thread yang sama melanjutkan percakapan, dan bila tidak terselesaikan email dikonversi menjadi request biasa.
  6. Setiap channel dapat diaktifkan/dinonaktifkan sendiri, dan default request type untuk Escalate dapat di-override per channel.
  7. Render step menyesuaikan channel: pilihan tampil sebagai tombol atau dropdown di portal/widget/Slack/Teams, dan sebagai daftar bernomor yang dijawab dengan teks di email.
  8. Bila integrasi chat diputus, channel virtual agent terkait berhenti otomatis dan overview menampilkan status "Terputus" dengan tautan untuk menyambungkan ulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Email channel dan widget di [30](30-jsm-help-center-portal-pelanggan-kanal.md); integrasi Slack/Teams di [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-030 s.d. INT-036).

### JSM-AI-054 Virtual Service Agent - Hand-off to Agent

- **Deskripsi:** Eskalasi dari virtual agent ke agent manusia lewat standard flow Escalate dengan membawa konteks percakapan.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Customer/Pelapor (memicu), Agent Layanan (menerima work item), Admin Proyek (mengatur default).
- **Aturan perilaku (acceptance criteria):**
  1. Escalate membuat work item dengan default request type (per service project atau per channel) atau request type yang ditetapkan step "Change request type".
  2. Jawaban dari step "Ask for information" dipetakan ke field request yang sesuai; field wajib request type yang belum terisi ditanyakan sebelum work item dibuat.
  3. Transkrip percakapan dilampirkan ke work item sebagai internal comment (atau panel transkrip) yang dapat dibaca agent.
  4. Customer menerima pesan konfirmasi berisi nomor dan tautan request di channel yang sama.
  5. Di Slack/Teams, balasan public agent pada work item tersinkron ke thread chat dan balasan customer di thread masuk sebagai komentar.
  6. ⚠ Customer dapat meminta bantuan manusia kapan saja dengan frasa seperti "bicara dengan agent"/"talk to a human"; frasa ini diperlakukan sebagai intent sistem yang langsung menjalankan Escalate.
  7. ⚠ Bila customer tidak memiliki akses ke request type tujuan (request type restrictions), Escalate memakai default request type service project; bila itu pun tidak dapat diakses, customer diarahkan ke portal dengan pesan "Silakan ajukan permintaan melalui portal".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request type restrictions di [29](29-jsm-service-project-request-type-formulir.md); komentar internal/public di [06](06-work-item-inti.md).

### JSM-AI-055 Virtual Service Agent - Testing

- **Deskripsi:** Alat uji virtual agent sebelum dan sesudah go-live.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. "Uji virtual service agent" membuka simulator chat di halaman konfigurasi untuk mencoba intent, flow (termasuk draf yang belum dipublikasikan), dan AI answers.
  2. ⚠ Selama pengujian, panel samping menampilkan intent yang ter-match beserta tingkat confidence (tinggi/sedang/rendah) dan step flow yang sedang berjalan.
  3. ⚠ Escalate di mode uji tidak membuat request nyata; simulator menampilkan pratinjau work item (request type dan field) yang akan dibuat, dan web request dapat dijalankan nyata atau di-mock sesuai pilihan admin.
  4. ⚠ Setelah perubahan training phrases, model dilatih ulang otomatis; status "Melatih..." tampil sampai selesai (target di bawah 5 menit), dan simulator memakai model terbaru setelah status selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Percakapan uji tidak dihitung kuota (JSM-AI-058).

### JSM-AI-056 Virtual Service Agent - Conversation Data & Analytics

- **Deskripsi:** Data percakapan untuk mengukur kinerja virtual agent dan memperbaiki intent serta knowledge base.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Manajer Layanan/ITSM (service owner), Agent Layanan (transkrip bila diizinkan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Dashboard menampilkan metrik jumlah percakapan, assisted, resolved, escalated, dan unresolved beserta persentasenya.
  2. Metrik dapat dipecah per channel, per intent, dan untuk AI answers.
  3. Semua metrik dapat difilter rentang tanggal (7, 30, 90 hari, atau kustom).
  4. Skor CSAT dari Resolve with CSAT ditampilkan rata-rata dan distribusinya.
  5. ⚠ Daftar percakapan dan transkrip dapat dibuka sesuai pengaturan "siapa yang dapat melihat percakapan" (admin saja atau admin dan agent); pengguna lain mendapat HTTP 403.
  6. "Gunakan data percakapan" menampilkan frasa yang tidak ter-match, dengan aksi tambahkan sebagai training phrase intent atau tandai sebagai topik artikel baru.
  7. Halaman "Cara data set dilatih" menjelaskan data yang dipakai melatih model intent (training phrases service project, tanpa data lintas tenant).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `Conversation` (channel, intent, outcome, CSAT, transkrip) dengan retensi mengikuti kebijakan privasi (JSM-AI-010); data juga diekspor ke data lake (JSM-AI-066).

### JSM-AI-057 Virtual Service Agent - Multilingual Support

- **Deskripsi:** Dukungan beberapa bahasa untuk virtual agent dan AI answers.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (menyediakan konten per bahasa), Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Bahasa pesan customer dideteksi otomatis pada pesan pertama; bila deteksi gagal, dipakai bahasa default help center.
  2. ⚠ Konten intent (training phrases, pesan) dan teks standard flow dapat disediakan per bahasa sesuai bahasa yang diaktifkan di help center; bahasa tanpa terjemahan jatuh ke bahasa default.
  3. ⚠ AI answers disusun dalam bahasa customer walaupun artikel sumber berbahasa lain.
  4. ⚠ Daftar bahasa yang didukung minimal mencakup bahasa yang tersedia di UI produk (EN dan ID), dapat diperluas melalui konfigurasi provider AI.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Multibahasa help center di [30](30-jsm-help-center-portal-pelanggan-kanal.md); i18n EN/ID mengikuti `internal/shared/constants/error.go`.

### JSM-AI-058 Virtual Service Agent - Usage Limits & Extra Assisted Conversations

- **Deskripsi:** Model kuota dan penagihan berbasis assisted conversations untuk virtual agent.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Org admin dan admin billing (mengatur batas), Admin Proyek (melihat pemakaian space).
- **Aturan perilaku (acceptance criteria):**
  1. Assisted conversation adalah percakapan di mana virtual agent memberi bantuan (intent ter-match atau AI answer diberikan); percakapan yang langsung dieskalasi tanpa bantuan tidak dihitung.
  2. ⚠ Kuota bulanan assisted conversations termasuk dalam plan dan dapat dikonfigurasi per plan (default Premium 1.000/bulan, Enterprise tanpa batas); kuota direset pada tanggal awal siklus billing.
  3. Halaman Track Feature Usage (JSM-AI-012) menampilkan pemakaian virtual agent terhadap batas, per service project dan total situs.
  4. Admin dapat menetapkan batas pemakaian maksimum (termasuk overage) agar tidak melebihi anggaran; batas tidak boleh lebih kecil dari pemakaian bulan berjalan.
  5. ⚠ Extra assisted conversations di atas kuota dicatat per percakapan dan diteruskan ke modul billing (ADM-053) dengan tarif yang dapat dikonfigurasi operator.
  6. ⚠ Saat batas tercapai, virtual agent berhenti memberi bantuan dan percakapan baru langsung menjalankan Escalate (atau diarahkan ke portal pada channel tanpa pembuatan request); admin menerima notifikasi pada 80% dan 100% batas.
  7. ⚠ Percakapan di simulator uji (JSM-AI-055) tidak dihitung kuota.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Usage & plan limits di [22](22-administrasi-sistem.md) (ADM-042, ADM-053). Keputusan terbuka: apakah clone self-hosted menerapkan kuota sama sekali atau hanya metering.

## Customer Service Management (dukungan pelanggan eksternal)

### JSM-AI-059 Customer Service Management - Customer Service Management App

- **Alias inventaris:** Customer service management (external support)
- **Deskripsi:** App terpisah dalam koleksi Service (padanan Service Collection) untuk dukungan pelanggan eksternal, dengan konteks customer yang kaya, kanal omnichannel, live chat, dan AI. Heading ini juga menggabungkan entri ringkas "Customer service management (external support)" dari lampiran Service Management.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (mengaktifkan app dan lisensi), Admin Proyek (mengonfigurasi customer service project), Agent Layanan, Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia template "Customer service project" yang membuat request type, queue, SLA, dan field customer/organization/product/entitlement bawaan untuk dukungan eksternal.
  2. Satu situs dapat memiliki beberapa help center ber-brand (logo, warna, nama, domain kustom) yang masing-masing menautkan customer service project tertentu.
  3. Kanal yang didukung: email, help center, dan embeddable widget/chat; semua kanal membuat atau memperbarui work item di space yang sama.
  4. Admin mendapat snippet embed (script) untuk memasang widget chat/support di situs eksternal perusahaan, dengan daftar domain yang diizinkan.
  5. ⚠ Percakapan bersifat AI-first oleh agen layanan customer (JSM-AI-064), dengan hand-off ke agent manusia secara live chat dalam sesi yang sama.
  6. ⚠ Agent CSM memiliki inbox/antrian percakapan live chat (status: menunggu, aktif, selesai) yang setiap percakapannya terhubung ke work item.
  7. Percakapan dapat dikonversi menjadi work item, dan riwayat percakapan tampil di issue view work item tersebut.
  8. Saat chat, panel agent menampilkan konteks customer, organization, product, dan entitlement (JSM-AI-063).
  9. ⚠ Admin dapat mengatur jam layanan dan ketersediaan agent untuk live chat; di luar jam layanan, widget menampilkan pesan offline dan menawarkan pembuatan request via email/form.
  10. ⚠ Konteks customer lintas departemen (request di service project lain dan work item terkait) ditampilkan melalui graf relasi data internal produk ini (padanan Teamwork Graph), sebatas permission agent.
  11. ⚠ Laporan CSM bawaan mencakup volume per kanal, waktu respons pertama chat, CSAT, dan volume per produk.
  12. Lampiran: profil customer memiliki detail fields kustom (rinci di JSM-AI-062).
  13. Lampiran: profil organization memiliki detail fields (rinci di JSM-AI-062).
  14. ⚠ Lampiran: products dapat dipilih pada request (rinci di JSM-AI-060).
  15. ⚠ Lampiran: entitlements per customer/organization per produk (rinci di JSM-AI-061).
  16. Lampiran: customer context panel di issue view (rinci di JSM-AI-063).
  17. Lampiran: help center eksternal ber-brand dengan custom domain (aturan 2 di atas; domain kustom di modul 30).
  18. Lampiran: kanal email, portal, widget/chat, dan agen AI (aturan 3, 5 di atas dan JSM-AI-064).
  19. ⚠ Lampiran: CSM tersedia sebagai app terpisah dalam koleksi Service (JSM-AI-005), dengan lisensi agent terpisah dari service project internal.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); customer service project dibuat sebagai company-managed.
- **Catatan & ketergantungan:** Help center, custom domain, widget di [30](30-jsm-help-center-portal-pelanggan-kanal.md); queue, SLA, CSAT di [31](31-jsm-antrian-sla-knowledge-base-laporan.md); koleksi lisensi di [22](22-administrasi-sistem.md) (ADM-056). Kanal voice/telepon tidak dibangun (lihat Di luar cakupan). Feature module baru `internal/features/csm`.

### JSM-AI-060 Customer Service Management - Products

- **Deskripsi:** Katalog produk yang didukung tim dukungan, dipakai untuk mengklasifikasi request dan entitlement.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Admin Proyek (mengelola katalog), Agent Layanan (memilih produk di request).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat membuat, mengubah, mengarsipkan, dan menghapus produk (nama unik wajib, deskripsi, status aktif); produk yang masih dipakai request atau entitlement hanya dapat diarsipkan, tidak dihapus.
  2. ⚠ Field sistem "Product" tersedia di request dan dapat diisi customer di portal (bila ditampilkan) atau oleh agent; hanya produk aktif yang muncul di pilihan.
  3. ⚠ Produk dapat dihubungkan ke request type sehingga pilihan Product di request type tersebut dibatasi ke produk yang terhubung.
  4. ⚠ Laporan dan filter JQL mendukung pengelompokan per produk (mis. `product = "Aplikasi Mobile"`).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `Product`; field JQL baru didaftarkan ke [12](12-jql-bahasa-kueri.md).

### JSM-AI-061 Customer Service Management - Entitlements

- **Deskripsi:** Hak dukungan customer atau organization atas produk, dengan tingkat dukungan dan masa berlaku.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Admin Proyek (mengelola entitlement), Agent Layanan (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Entitlement dibuat per customer atau per organization untuk satu produk; satu customer/organization dapat memiliki banyak entitlement untuk produk berbeda.
  2. ⚠ Entitlement memiliki detail field seperti support tier (mis. Basic/Gold/Platinum), tanggal mulai, dan tanggal kedaluwarsa; entitlement kedaluwarsa ditandai "Expired" dan tidak dipakai untuk SLA.
  3. ⚠ Kondisi SLA dan otomasi dapat memakai atribut entitlement (mis. goal SLA lebih ketat untuk tier Platinum, rule peringatan bila entitlement kedaluwarsa).
  4. ⚠ Entitlement aktif yang relevan dengan produk request tampil di context panel (JSM-AI-063).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `Entitlement`; SLA di [31](31-jsm-antrian-sla-knowledge-base-laporan.md); otomasi di [20](20-otomasi.md).

### JSM-AI-062 Customer Service Management - Customer & Organization Details

- **Deskripsi:** Field detail kustom untuk profil customer dan organization.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin (mendefinisikan field), Agent Layanan (mengisi/mengedit nilai sesuai izin).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat mendefinisikan detail field bertipe text, number, select (single/multi), date, dan URL untuk customer maupun organization; nama field unik per jenis profil.
  2. Agent dan admin dapat mengedit nilai detail field di profil; admin dapat membatasi field tertentu menjadi hanya-admin, dan setiap perubahan tercatat di riwayat profil.
  3. Daftar customer dan organization dapat dicari dan difilter berdasarkan nilai detail field.
  4. Organization memiliki daftar managed email domains; customer baru dengan domain email yang cocok dapat dipetakan otomatis ke organization tersebut melalui otomasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Customers dan organizations di [30](30-jsm-help-center-portal-pelanggan-kanal.md); API baca/tulis detail ditandai experimental (JSM-AI-028).

### JSM-AI-063 Customer Service Management - Customer Profile & Context Panel

- **Deskripsi:** Profil customer terpadu dan panel konteks di request view.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan (hanya data yang boleh dilihat sesuai permission).
- **Aturan perilaku (acceptance criteria):**
  1. Profil customer menampilkan riwayat request, organizations, entitlements, dan detail fields dalam satu halaman.
  2. Context panel di request view menampilkan ringkasan profil pelapor (organization, entitlement aktif, detail penting, request terbuka lain) dan tautan ke profil lengkap.
  3. Profil menampilkan riwayat interaksi lintas kanal (email, portal, chat, widget) dalam satu timeline berurutan waktu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request yang tidak dapat diakses agent tidak muncul di riwayat; rujukan JSM-AI-061, JSM-AI-062.

### JSM-AI-064 Rovo Customer Service Agent

- **Deskripsi:** Agen AI percakapan customer-facing multi-kanal untuk Customer Service Management (padanan Rovo Customer Service Agent), evolusi dari Virtual Service Agent.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin (mengonfigurasi), Customer/Pelapor, Agent Layanan (menerima hand-off).
- **Aturan perilaku (acceptance criteria):**
  1. Agen menjawab customer di email, chat/widget, dan help center memakai knowledge yang dikonfigurasi.
  2. ⚠ Admin mengonfigurasi knowledge sources (knowledge base, konektor) dan instruksi (nada, batasan topik, kebijakan eskalasi) dalam teks bebas.
  3. ⚠ Agen mengeskalasi ke agent manusia, termasuk hand-off live chat dalam sesi yang sama bila ada agent tersedia, atau membuat work item bila tidak.
  4. ⚠ Analitik percakapan (jumlah, resolved, escalated, CSAT) tersedia seperti JSM-AI-056.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai `AIProvider` yang dapat dimatikan (JSM-AI-036), kuota JSM-AI-012; berbagi engine percakapan dengan Virtual Service Agent (JSM-AI-047). Nama fitur produk ini tidak memakai merek "Rovo".

## Mobile, analitik, migrasi & sandbox

### JSM-AI-065 JSM Mobile App

- **Alias inventaris:** JSM mobile apps
- **Deskripsi:** Kapabilitas Service Management di aplikasi mobile produk ini untuk agent, approver, dan responder on-call. Heading ini juga menggabungkan entri ringkas "JSM mobile apps" dari lampiran Service Management.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, approver, responder on-call; Customer/Pelapor lewat browser mobile.
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat melihat daftar queue dan request di dalamnya, serta membuka detail request.
  2. Agent dapat menambah komentar internal maupun public; pilihan tipe komentar wajib eksplisit sebelum kirim.
  3. Approver dapat menyetujui atau menolak approval dari mobile.
  4. Agent dapat melihat status SLA (sisa waktu, breach) pada request.
  5. Push notification dikirim untuk request yang di-assign, mention, approval, dan alert.
  6. Responder dapat acknowledge/close alert dan melihat jadwal on-call.
  7. Customer memakai portal melalui browser mobile dengan layout responsif tanpa app native.
  8. Lampiran: agent dapat melakukan transisi status request dari mobile.
  9. Lampiran: responder dapat membuat override jadwal on-call dari mobile.
  10. Lampiran: critical alert dapat menembus mode Do Not Disturb perangkat bila pengguna mengizinkannya.
  11. ⚠ Lampiran: customer dapat membuat request dari app mobile (opsional; default lewat browser mobile).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aplikasi mobile di [21](21-navigasi-ux-preferensi.md) (NAV-050); push di [19](19-notifikasi-email.md) (NTF-038); alert/on-call di [34](34-jsm-operations-alerting-on-call.md).

### JSM-AI-066 Atlassian Analytics & Data Lake for JSM

- **Deskripsi:** Analitik tingkat Enterprise dan data lake untuk data Service Management.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Org admin, analis (Stakeholder/Viewer untuk dashboard).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Data lake menyediakan tabel requests, SLA, CSAT, Assets, alerts, dan percakapan virtual agent dengan skema terdokumentasi.
  2. Tersedia template dashboard Service Management (volume, SLA, CSAT, beban agent).
  3. Admin dapat mengontrol data yang masuk data lake (pilih service project dan jenis data); data yang dikecualikan dihapus dari data lake pada sinkronisasi berikutnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Analytics & Data Lake umum di [16](16-laporan-insight.md) (RPT-038 s.d. RPT-041). Nama produk ini tidak memakai merek "Atlassian Analytics".

### JSM-AI-067 JSM Cloud Migration & Sandbox

- **Deskripsi:** Migrasi data Service Management dari instalasi lain dan lingkungan pengujian (sandbox).
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. Alat migrasi (padanan JCMA) memindahkan project Service Management beserta request type, queue, SLA, customer, dan organization; laporan pra-migrasi menampilkan item yang tidak didukung.
  2. Data Assets (schema, object type, object) dapat dimigrasi dengan alat yang sama.
  3. Sandbox tersedia untuk plan Premium/Enterprise untuk menguji konfigurasi Service Management sebelum diterapkan ke produksi.
  4. Release tracks memungkinkan admin menunda rilis perubahan fitur ke situs produksi dan mengujinya di sandbox lebih dulu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Migrasi di [22](22-administrasi-sistem.md) (ADM-039) dan [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-026); sandbox ADM-049.

## Pertanyaan terbuka

1. Apakah clone self-hosted menerapkan batas kursi agent dan kuota assisted conversations sama sekali, atau hanya melakukan metering tanpa pemblokiran (JSM-AI-003, JSM-AI-058)?
2. Apakah REST API Service Management menyediakan alias path kompatibel `/rest/servicedeskapi` agar klien yang sudah ada dapat dipakai tanpa modifikasi (JSM-AI-024)?
3. Penyedia AI default apa (model self-hosted vs API eksternal) yang dikemas bersama produk, dan apakah AI answers/Virtual Service Agent tersedia tanpa penyedia eksternal (JSM-AI-036, JSM-AI-052)?
4. Apakah Virtual Service Agent (JSM-AI-047) dan agen layanan customer AI (JSM-AI-064) dibangun sebagai satu engine dengan dua mode konfigurasi atau dua produk terpisah?
5. Apakah Customer Service Management (JSM-AI-059) dikemas sebagai app berlisensi terpisah atau sebagai template di dalam Service Management?
6. Apakah live chat CSM memerlukan infrastruktur realtime sendiri atau memakai kanal realtime umum produk (modul 27)?
7. Apakah prioritas ADM-052 (saat ini P3) perlu dinaikkan minimal ke P2 agar tidak lebih rendah daripada fitur AI Service Management yang bergantung padanya (JSM-AI-036)?
8. Bagaimana status rilis Request Resolver dan agen AI otonom: diluncurkan sebagai beta di Feature Lab (JSM-AI-046) atau langsung GA (JSM-AI-044)?

## Di luar cakupan & ekosistem

- **Fitur legacy / dc-only-or-deprecated:** Legacy JSM automation dan Legacy Assets Automation (JSM-AI-022, JSM-AI-023) serta jalur migrasi dari alat on-call eksternal (JSM-AI-034) hanya menjadi rujukan; tidak dibangun karena produk ini tidak pernah memiliki mesin legacy tersebut.
- **App Marketplace pihak ketiga:** add-on Service Management (portal tambahan, chat, telepon/voice, CRM, survei) tidak dibangun native; integrasi hanya melalui REST API, webhook, dan extension point (JSM-AI-033).
- **Konektor merek pihak ketiga:** Workato, Workday, dan iPaaS lain tidak dibangun sebagai konektor bermerek; yang dibangun adalah konektor generik (JSM-AI-050, INT-003).
- **Layanan vendor lintas produk:** Atlassian Intelligence/Rovo, Teamwork Graph, Atlassian Analytics, Opsgenie, Statuspage, dan Confluence sebagai layanan eksternal tidak di-clone; produk ini menyediakan padanan internal (antarmuka `AIProvider`, graf relasi, data lake, knowledge base bawaan) atau integrasi.
- **Kanal voice/telepon dan SMS** untuk Customer Service Management tidak dibangun.
- **Model AI milik vendor:** produk ini tidak melatih atau mengemas model AI proprietary; seluruh fitur AI bersifat opsional dan dapat dimatikan.
