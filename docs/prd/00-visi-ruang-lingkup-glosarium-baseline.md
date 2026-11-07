# 00 — Visi, Ruang Lingkup, Glosarium & Baseline Repo

Modul ini adalah pintu masuk seluruh PRD. Isinya menetapkan visi produk, batas ruang lingkup per fase delivery, glosarium istilah Jira ↔ padanan Indonesia yang dipakai konsisten di semua modul, kondisi repo saat ini beserta gap dan remediasi Sprint 0, serta konvensi penulisan dokumen (template fitur, arti prioritas/tier/fase, prefix ID). Selain itu modul ini memuat tiga fitur lintas platform yang menjadi fondasi identitas produk: kepatuhan nama & merek, terminologi UI (work item & space), serta mekanisme tier plan & feature flag. Nama final produk clone adalah **Jari**; dokumen ini menyebutnya "produk ini" atau "clone".

Modul terkait:

- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — titik awal remediasi auth yang sudah ada.
- [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — batas jumlah pengguna per plan (VISI-003).
- [03 Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) — RBAC yang belum ada di baseline.
- [04 Proyek](04-proyek.md) — istilah Project/Space (VISI-002).
- [06 Work Item Inti](06-work-item-inti.md) — istilah Issue/Work item (VISI-002).
- [21 Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) — label UI, bahasa, prompt upgrade plan.
- [22 Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) — konsol admin untuk flag dan tier.
- [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — importer Jira (penyebutan merek secara deskriptif).
- [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) — i18n, deployment, keamanan, pengujian.
- [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) — entitas yang belum ada di baseline.

## Persona terkait

- **Admin Situs** — mengelola tier/edisi yang aktif, feature flag, kill switch, dan profil deployment teregulasi (VISI-003); memastikan branding instansi tidak melanggar merek pihak lain (VISI-001).
- **Admin Proyek/Project Lead** — terdampak oleh terminologi project/space dan fitur yang digating per tier di tingkat proyek.
- **Developer/Anggota Tim** — pengguna harian yang membaca label UI, email, dan dokumentasi dengan terminologi konsisten (VISI-002).
- **Product Owner/PM** — dalam konteks modul ini juga mewakili pemilik produk clone: memutuskan nama produk, pemetaan fitur → tier, dan jadwal rollout fitur (VISI-001, VISI-003).
- **Integrator/Developer API** — bergantung pada kompatibilitas nama field API/JQL lama (`issuetype`, `project`) dan pada endpoint konfigurasi flag (VISI-002, VISI-003).
- **Stakeholder/Viewer** — membaca dokumentasi dan email dengan istilah yang sama dengan UI.

Pemangku kepentingan internal non-pengguna: tim Engineering (implementasi flag dan i18n) dan penasihat hukum/Legal (review merek dan lisensi dependensi).

## User story modul

- Sebagai Product Owner/PM, saya ingin produk memakai nama, logo, dan palet warna sendiri di UI, email, dokumentasi, domain, dan user agent agar produk tidak dianggap berafiliasi dengan pemilik merek Jira (VISI-001).
- Sebagai Product Owner/PM, saya ingin materi pemasaran dan importer hanya menyebut "kompatibel dengan format/API Jira" secara nominatif agar calon pengguna paham kemampuan migrasi tanpa risiko pelanggaran merek (VISI-001).
- Sebagai Admin Situs, saya ingin laporan lisensi seluruh dependensi OSS tersedia di setiap rilis agar distribusi self-hosted di instansi saya patuh terhadap kebijakan lisensi (VISI-001).
- Sebagai Developer/Anggota Tim, saya ingin label sidebar, breadcrumb, dialog, dan email memakai istilah "work item" dan "space" secara konsisten agar saya tidak bingung saat berpindah layar (VISI-002).
- Sebagai Integrator/Developer API, saya ingin field API dan kueri lama seperti `issuetype` dan `project` tetap berfungsi agar skrip dan integrasi saya tidak rusak setelah perubahan istilah (VISI-002).
- Sebagai Admin Situs, saya ingin mengaktifkan atau menonaktifkan fitur per tenant, per pengguna, atau per persentase pengguna agar rollout fitur baru dapat dilakukan bertahap dan aman (VISI-003).
- Sebagai Admin Situs, saya ingin kill switch yang mematikan fitur bermasalah seketika tanpa redeploy agar dampak insiden dapat dibatasi (VISI-003).
- Sebagai Admin Situs di instansi teregulasi, saya ingin preset "regulated deployment" yang mematikan fitur yang butuh egress ke layanan eksternal agar instalasi saya memenuhi kebijakan data residency (VISI-003).
- Sebagai Product Owner/PM, saya ingin pemetaan fitur → tier (Free/Standard/Premium/Enterprise) dikonfigurasi terpusat agar edisi komersial atau simulasi plan dapat dibentuk tanpa mengubah kode fitur (VISI-003).
- Sebagai Integrator/Developer API, saya ingin SPA dan klien lain membaca daftar flag aktif lewat endpoint konfigurasi terautentikasi agar UI menyembunyikan fitur yang tidak tersedia secara konsisten dengan backend (VISI-003).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur | ID |
| --- | --- | --- |
| P0 | 1 | VISI-001 |
| P1 | 0 | — |
| P2 | 2 | VISI-002, VISI-003 |
| P3 | 0 | — |
| **Total** | **3** | |

| Fase | Jumlah fitur | ID |
| --- | --- | --- |
| Fase 1 | 2 | VISI-001, VISI-002 |
| Fase 2 | 0 | — (komponen kill switch & endpoint flag dari VISI-003 dimajukan ke sini) |
| Fase 3 | 1 | VISI-003 |
| Fase 4-7 | 0 | — |

Strategi urutan implementasi: VISI-001 dikerjakan paling awal (Sprint 0) karena nama produk, aset brand, dan user agent menyentuh setiap layar dan email; mengganti belakangan jauh lebih mahal. VISI-002 secara inventaris berprioritas P2, tetapi keputusan terminologi diambil di Fase 1 karena biayanya hampir nol bila semua label sejak awal dipusatkan di layer i18n, sedangkan migrasi istilah setelah ratusan string tersebar sangat mahal. VISI-003 ditempatkan di Fase 3 untuk mapping tier dan batas plan, namun inti layanan feature flag, kill switch, dan endpoint konfigurasi untuk SPA (butir aturan 1-3 dan 9-12) sudah dibangun di Fase 2 dengan prioritas P1 karena dibutuhkan untuk rollout fitur konfigurabilitas. Sandbox/release tracks dan profil regulated deployment dikerjakan paling akhir (Fase 3-4) setelah daftar fitur yang memerlukan egress (AI, smart link publik, integrasi SaaS) benar-benar ada.

## Visi produk

Produk ini adalah **clone Jira Cloud open-source dan self-hosted** yang dibangun di atas satu binary Go (Gin + GORM + PostgreSQL) dengan SPA SvelteKit ter-embed. Targetnya adalah tim Indonesia dan global yang membutuhkan pelacak kerja setara Jira Cloud tetapi ingin mengendalikan data, biaya, dan lokasi hostingnya sendiri.

Nilai utama:

- **Kendali data & kedaulatan** — dipasang di server sendiri (on-premise, cloud privat, atau VPS lokal), tanpa ketergantungan pada layanan SaaS eksternal; cocok untuk instansi pemerintah, BUMN, perbankan, dan organisasi teregulasi (lihat profil regulated deployment pada VISI-003).
- **Cakupan setara Jira Cloud** — satu produk mencakup pelacakan kerja tim perangkat lunak dan bisnis (eks Jira Software + Work Management), Service Management (portal, queue, SLA, ITSM, Assets, Operations), dan Product Discovery (ideas, insights, views), dirilis bertahap dalam 7 fase.
- **Bilingual sejak awal** — UI, pesan galat, dan email tersedia dalam Bahasa Indonesia dan Inggris (baseline sudah memiliki `LanguageMiddleware` dan pesan galat id/en).
- **Kompatibel untuk migrasi** — bahasa kueri JQL-kompatibel, REST API dengan bentuk payload yang familiar, dan importer dari Jira (CSV/JSON/API) agar tim dapat pindah tanpa kehilangan riwayat.
- **Sederhana dioperasikan** — satu binary + PostgreSQL sebagai dependensi minimum; komponen lain (cache, queue, object storage, mesin pencarian) bersifat opsional dan dapat ditingkatkan sesuai skala.
- **Terbuka & dapat diperluas** — kode sumber terbuka, REST API publik, webhook, dan titik ekstensi; tanpa mengikat pengguna pada ekosistem Marketplace tertentu.
- **Bebas dari merek pihak ketiga** — memakai identitas sendiri dan istilah generik (VISI-001).

Prinsip produk:

- Default yang masuk akal untuk tim kecil (team-managed) dengan jalur naik ke konfigurasi organisasi (company-managed).
- Perilaku server adalah sumber kebenaran: setiap izin, flag, dan batas plan ditegakkan di backend, UI hanya mencerminkan.
- Fitur yang bergantung pada layanan eksternal (AI, unfurl link publik, integrasi SaaS) selalu dapat dimatikan.

## Ruang lingkup

**In-scope** (dirilis bertahap sesuai [Peta fase delivery](#peta-fase-delivery)):

- Jira terpadu (eks Jira Software + Jira Work Management) untuk semua plan Free/Standard/Premium/Enterprise: proyek team-managed dan company-managed, work item, workflow, board Scrum/Kanban, backlog & sprint, JQL, filter, dashboard, laporan, time tracking, rilis & komponen, timeline, Plans, Teams, Goals, otomasi, notifikasi, proyek bisnis & formulir, administrasi, REST API & webhook, integrasi DevOps, impor/ekspor.
- Jira Service Management: service project, request type, forms, help center & portal, customers & organizations, kanal email & widget, queues, SLA, knowledge base, approvals, CSAT, laporan, ITSM (incident, problem, change, services), Assets/CMDB & AQL, Operations (alerts, on-call, escalation), virtual service agent & AI opsional, Customer Service Management.
- Jira Product Discovery: ideas, fields & formula, views (list, board, matrix, timeline), sharing & roadmap, insights, delivery tracking, roles.
- Kebutuhan non-fungsional lintas produk: keamanan, i18n ID/EN, aksesibilitas, observabilitas, backup, packaging self-hosted.

**Out-of-scope awal** (dapat ditinjau ulang setelah Fase 7):

- Aplikasi mobile native iOS/Android; pengalaman mobile dipenuhi oleh web responsif (lihat modul [21](21-navigasi-ux-preferensi.md)).
- Atlassian Marketplace dan app pihak ketiga (Tempo, ScriptRunner, Xray, BigPicture, dsb.) serta runtime Forge/Connect; produk ini hanya menyediakan API, webhook, dan titik ekstensi sendiri.
- Layanan Atlassian lintas produk: Confluence, Bitbucket, Trello, Loom, Compass, Statuspage, Atlas, Atlassian Guard/Access, Atlassian Analytics, Rovo sebagai layanan terkelola, Atlassian Home/Teamwork Graph sebagai produk; hanya padanan fungsional minimal yang dibutuhkan Jira (mis. entitas Team dan Goal) yang dibangun.
- Fitur Data Center-only atau yang sudah deprecated (tier `dc-only-or-deprecated`) — dicatat hanya sebagai rujukan.
- Penagihan (billing) komersial dan lisensi berbayar; gating tier di VISI-003 bersifat konfigurasi, bukan sistem pembayaran.

## Glosarium

Istilah Jira ditulis dalam bahasa Inggris dan dipakai apa adanya di seluruh PRD; kolom padanan menunjukkan terjemahan yang boleh dipakai di UI Bahasa Indonesia (keputusan akhir label UI mengikuti VISI-002 dan katalog i18n). Kolom "Modul" menunjuk modul yang menjadi pemilik konsep.

### Struktur kerja & proyek

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Work item (Issue) | Item kerja | Unit kerja yang dilacak (tugas, bug, story, request, idea); "issue" adalah istilah lama yang tetap dipakai di API dan JQL. | 06 |
| Work type (Issue type) | Tipe kerja | Klasifikasi work item (Epic, Story, Task, Bug, Subtask, atau tipe kustom) yang menentukan field, layar, dan workflow. | 05 |
| Issue key | Kunci work item | Pengenal unik terbaca manusia berformat `<PROJECTKEY>-<nomor>` (mis. `WEB-42`). | 28 |
| Project (Space) | Proyek (Ruang) | Wadah work item beserta konfigurasinya; "space" adalah istilah UI baru, "project" tetap dipakai di API/JQL. | 04 |
| Project key | Kunci proyek | Kode pendek huruf kapital unik per site yang menjadi prefix issue key. | 04 |
| Team-managed project | Proyek dikelola tim | Proyek dengan konfigurasi mandiri (tipe, field, workflow) yang diatur oleh admin proyek tanpa skema global. | 04 |
| Company-managed project | Proyek dikelola perusahaan | Proyek yang konfigurasinya memakai skema bersama dan dikelola Admin Situs. | 04 |
| Epic | Epic | Work item besar yang menampung beberapa story/task; level hierarki di atas tipe standar. | 05, 06 |
| Story | Story | Kebutuhan dari sudut pandang pengguna yang dapat diselesaikan dalam satu sprint. | 06 |
| Task | Tugas | Pekerjaan umum yang perlu diselesaikan. | 06 |
| Bug | Bug | Cacat atau perilaku yang tidak sesuai harapan. | 06 |
| Subtask | Sub-tugas | Pecahan work item standar; level hierarki di bawah tipe standar. | 06 |
| Hierarchy level | Level hierarki | Tingkat dalam hierarki work item (Subtask < standar < Epic < level kustom di atas Epic). | 05 |
| Parent | Induk | Work item satu level di atas dalam hierarki (mis. Epic untuk Story). | 06 |
| Issue link | Tautan work item | Relasi bertipe antarwork item (blocks, relates to, duplicates, clones). | 06 |
| Summary | Ringkasan | Judul singkat work item (wajib). | 06 |
| Description | Deskripsi | Isi rinci work item berformat teks kaya. | 07 |
| Assignee | Penerima tugas | Pengguna yang bertanggung jawab mengerjakan work item. | 06 |
| Reporter | Pelapor | Pengguna yang membuat atau melaporkan work item. | 06 |
| Priority | Prioritas | Tingkat kepentingan work item (Highest … Lowest), dapat diatur per priority scheme. | 05 |
| Label | Label | Penanda teks bebas untuk mengelompokkan work item lintas proyek. | 06 |
| Component | Komponen | Sub-bagian proyek (modul/tim) yang dapat memiliki default assignee. | 15 |
| Version (Release) | Versi (Rilis) | Titik rilis yang menghimpun work item lewat field Fix version/Affects version. | 15 |
| Fix version | Versi perbaikan | Versi tempat work item direncanakan/dirilis. | 15 |
| Affects version | Versi terdampak | Versi tempat bug ditemukan. | 15 |
| Watcher | Pemantau | Pengguna yang berlangganan notifikasi perubahan sebuah work item. | 06, 19 |
| Vote | Suara | Dukungan pengguna terhadap work item sebagai sinyal prioritas. | 06 |
| Comment | Komentar | Diskusi pada work item, mendukung mention dan pembatasan visibilitas. | 08 |
| Attachment | Lampiran | Berkas yang dilekatkan pada work item atau komentar. | 08 |
| Mention | Sebutan | Penyebutan pengguna dengan `@` yang memicu notifikasi. | 07 |
| History / Changelog | Riwayat perubahan | Catatan setiap perubahan field work item beserta pelaku dan waktunya. | 08 |
| Flag (Impediment) | Penanda hambatan | Tanda bahwa work item terhambat. | 10 |
| Clone | Duplikasi | Salinan work item beserta field terpilih. | 06 |

### Konfigurasi & workflow

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Workflow | Alur kerja | Kumpulan status dan transisi yang dilalui work item. | 09 |
| Status | Status | Posisi work item dalam workflow, termasuk kategori To Do / In Progress / Done. | 09 |
| Status category | Kategori status | Pengelompokan status menjadi To Do, In Progress, Done untuk laporan dan board. | 09 |
| Transition | Transisi | Perpindahan dari satu status ke status lain, dapat diberi condition, validator, dan post function. | 09 |
| Condition / Validator / Post function | Syarat / Validator / Aksi pasca-transisi | Aturan yang mengontrol siapa boleh bertransisi, apa yang harus valid, dan apa yang terjadi sesudahnya. | 09 |
| Resolution | Resolusi | Alasan penyelesaian work item (Done, Won't Do, Duplicate); work item tanpa resolusi dianggap belum selesai. | 09 |
| Scheme | Skema | Konfigurasi bersama yang dapat dipasang ke banyak proyek company-managed (workflow, field, screen, permission, notification). | 05 |
| Screen | Layar | Susunan field yang tampil saat create, edit, view, atau transisi. | 05 |
| Field | Field | Atribut data work item (sistem atau kustom). | 05 |
| Custom field | Field kustom | Field tambahan buatan admin dengan tipe tertentu (teks, angka, pilihan, tanggal, pengguna, dsb.). | 05 |
| Field configuration | Konfigurasi field | Aturan per field: wajib/tidak, tersembunyi, renderer, deskripsi. | 05 |
| Context (custom field) | Konteks field | Cakupan proyek/tipe tempat field kustom berlaku beserta opsi default-nya. | 05 |
| Issue layout | Tata letak work item | Susunan field di tampilan detail work item. | 05 |

### Agile: board, backlog & sprint

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Board | Papan | Visualisasi work item dalam kolom yang dipetakan ke status. | 10 |
| Kanban | Kanban | Metode alur kontinu dengan batas WIP; board tanpa sprint. | 10 |
| Scrum | Scrum | Metode iteratif berbasis sprint dengan backlog terurut. | 11 |
| Backlog | Backlog | Daftar terurut work item yang belum dikerjakan atau belum masuk sprint. | 11 |
| Sprint | Sprint | Iterasi berjangka waktu tetap berisi komitmen work item. | 11 |
| Sprint goal | Tujuan sprint | Pernyataan hasil yang ingin dicapai dalam sprint. | 11 |
| Velocity | Kecepatan tim | Rata-rata estimasi yang diselesaikan per sprint. | 16 |
| Rank | Peringkat urutan | Urutan relatif work item di backlog/board (disimpan dengan LexoRank). | 11, 28 |
| Swimlane | Lajur | Pengelompokan horizontal kartu di board (per assignee, epic, kueri, dsb.). | 10 |
| WIP limit | Batas WIP | Batas minimum/maksimum jumlah work item dalam satu kolom. | 10 |
| Quick filter | Filter cepat | Tombol filter berbasis JQL di board/backlog. | 10 |
| Card | Kartu | Representasi work item di board. | 10 |
| Estimate | Estimasi | Perkiraan usaha (story point atau waktu) sebuah work item. | 14 |
| Story point | Poin story | Satuan estimasi relatif untuk kompleksitas/usaha. | 14 |
| Burndown / Burnup | Burndown / Burnup | Grafik sisa/penyelesaian pekerjaan terhadap waktu. | 16 |
| Cumulative flow diagram | Diagram alur kumulatif | Grafik jumlah work item per status sepanjang waktu. | 16 |

### Pencarian, waktu & pelaporan

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| JQL | Bahasa kueri (JQL-kompatibel) | Bahasa kueri terstruktur `field operator value` untuk mencari work item; nama UI sendiri ditetapkan di VISI-001. | 12 |
| Filter (saved filter) | Filter tersimpan | Kueri JQL yang disimpan, dapat dibagikan dan dipakai board, dashboard, atau langganan email. | 13 |
| Issue navigator | Navigator work item | Halaman hasil pencarian berbentuk daftar/detail dengan kolom yang dapat diatur. | 13 |
| Worklog | Catatan waktu kerja | Entri waktu yang dihabiskan untuk sebuah work item. | 14 |
| Original / Remaining estimate | Estimasi awal / sisa | Perkiraan waktu awal dan sisa yang dipakai time tracking. | 14 |
| Dashboard | Dasbor | Halaman berisi gadget yang menampilkan ringkasan data. | 17 |
| Gadget | Gadget | Widget di dashboard (pie chart, filter results, sprint burndown, dsb.). | 17 |
| Report | Laporan | Analisis bawaan (velocity, sprint report, control chart, dsb.). | 16 |

### Perencanaan

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Timeline | Linimasa | Tampilan Gantt sederhana per proyek untuk Epic dan dependensi. | 18 |
| Plan (Advanced Roadmaps) | Rencana | Perencanaan lintas proyek dan tim dengan kapasitas, skenario, dan dependensi (Premium). | 18 |
| Team | Tim | Entitas tim lintas proyek yang dipakai di field Team, Plans, dan kapasitas. | 02, 18 |
| Goal | Sasaran | Tujuan organisasi yang ditautkan ke work item untuk melacak kontribusi. | 18 |
| Capacity | Kapasitas | Jumlah kerja yang sanggup diselesaikan tim per iterasi dalam Plan. | 18 |
| Dependency | Dependensi | Relasi urutan antarwork item yang divisualkan di timeline/plan. | 18 |

### Akses, keamanan & notifikasi

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Site | Situs | Satu instalasi/tenant produk beserta seluruh proyek dan penggunanya. | 22 |
| Organization (admin) | Organisasi | Wadah administratif di atas site untuk domain, kebijakan, dan pengguna terkelola. | 02 |
| Group | Grup | Kumpulan pengguna untuk pemberian izin massal. | 02 |
| Global permission | Izin global | Izin tingkat site (administer, browse users, create shared objects, dsb.). | 03 |
| Permission scheme | Skema izin | Pemetaan izin proyek (browse, create, edit, transition, dsb.) ke peran, grup, atau pengguna. | 03 |
| Project role | Peran proyek | Peran per proyek (Administrator, Member, Viewer) yang diisi pengguna/grup. | 03 |
| Issue security (level/scheme) | Keamanan work item | Pembatasan siapa yang dapat melihat work item tertentu. | 03 |
| Notification scheme | Skema notifikasi | Pemetaan event work item ke penerima notifikasi email. | 19 |
| API token | Token API | Kredensial pribadi untuk akses REST API tanpa kata sandi. | 01 |
| OAuth 2.0 | OAuth 2.0 | Protokol otorisasi untuk aplikasi pihak ketiga atas nama pengguna. | 23 |
| Webhook | Webhook | Panggilan HTTP keluar ke URL terdaftar saat event terjadi. | 23 |
| SSO / SAML / SCIM | SSO / SAML / SCIM | Login tunggal melalui IdP dan provisioning pengguna otomatis. | 01, 02 |
| Audit log | Log audit | Catatan aktivitas administratif dan keamanan. | 22, 27 |

### Otomasi & ekstensi

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Automation | Otomasi | Mesin aturan tanpa kode untuk menjalankan aksi berdasarkan event atau jadwal. | 20 |
| Rule | Aturan | Satu definisi otomasi berisi trigger, condition, dan action. | 20 |
| Trigger | Pemicu | Event yang memulai aturan (work item created, transitioned, scheduled, dsb.). | 20 |
| Condition (automation) | Syarat | Pemeriksaan yang harus terpenuhi agar aturan berlanjut. | 20 |
| Action | Aksi | Langkah yang dijalankan aturan (edit field, transisi, kirim email, webhook). | 20 |
| Smart value | Nilai pintar | Ekspresi templat `{{issue.summary}}` untuk mengakses dan memanipulasi data dalam aturan. | 20 |
| Audit log (automation) | Log eksekusi aturan | Riwayat eksekusi aturan beserta hasilnya. | 20 |
| Entity property | Properti entitas | Data kunci-nilai JSON tambahan yang dilekatkan ke work item/proyek oleh integrasi. | 28 |
| Development panel | Panel pengembangan | Informasi branch, commit, PR, build, dan deployment terkait work item. | 24 |
| Smart commit | Smart commit | Perintah dalam pesan commit untuk berkomentar, mencatat waktu, atau bertransisi. | 24 |
| Feature flag | Penanda fitur | Saklar konfigurasi untuk mengaktifkan fitur per tenant/pengguna/persentase. | 00 (VISI-003) |
| Kill switch | Saklar darurat | Feature flag khusus untuk mematikan fitur seketika saat insiden. | 00 (VISI-003) |

### Service Management

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Service project (Service space) | Proyek layanan | Proyek untuk menerima dan menangani permintaan layanan. | 29 |
| Request type | Tipe permintaan | Jenis permintaan yang dipilih customer di portal, dipetakan ke work type dan form. | 29 |
| Request (ticket) | Permintaan (tiket) | Work item di service project yang diajukan customer. | 29 |
| Form | Formulir | Formulir dinamis yang dilampirkan ke request type atau work item. | 29, 26 |
| Help center | Pusat bantuan | Halaman beranda publik/terautentikasi yang menghimpun portal dan artikel. | 30 |
| Portal (customer portal) | Portal pelanggan | Antarmuka customer untuk mengajukan dan memantau request suatu service project. | 30 |
| Customer | Pelanggan | Pengguna yang mengajukan request lewat portal/email tanpa lisensi agent. | 30 |
| Organization (JSM) | Organisasi pelanggan | Kelompok customer yang dapat berbagi request satu sama lain. | 30 |
| Request participant | Peserta permintaan | Customer lain yang diikutsertakan dalam request. | 29 |
| Agent | Agen | Anggota tim layanan berlisensi yang menangani request. | 29 |
| Queue | Antrean | Daftar request berbasis JQL yang harus ditangani agent. | 31 |
| SLA | SLA (perjanjian tingkat layanan) | Target waktu (time to first response, time to resolution) dengan kalender kerja dan penanganan pelanggaran. | 31 |
| SLA calendar | Kalender SLA | Jam kerja dan hari libur yang memengaruhi perhitungan SLA. | 31 |
| Canned response | Balasan siap pakai | Templat balasan yang dapat disisipkan agent. | 31 |
| Knowledge base | Basis pengetahuan | Kumpulan artikel bantuan untuk swalayan dan deflection. | 31 |
| CSAT | Kepuasan pelanggan (CSAT) | Survei penilaian kepuasan setelah request selesai. | 31 |
| Approval | Persetujuan | Langkah workflow yang menunggu keputusan approver. | 32, 26 |
| Incident | Insiden | Gangguan layanan tak terencana yang harus dipulihkan secepatnya. | 32 |
| Problem | Masalah | Akar penyebab dari satu atau lebih insiden. | 32 |
| Change | Perubahan | Permintaan perubahan terkontrol pada layanan/infrastruktur, dengan penilaian risiko dan persetujuan. | 32 |
| Post-incident review (PIR) | Tinjauan pasca-insiden | Analisis setelah insiden untuk pembelajaran dan tindak lanjut. | 32 |
| Service (service registry) | Layanan | Entitas layanan TI yang dapat ditautkan ke incident/change dan deployment. | 32 |
| Asset / CMDB | Aset / CMDB | Basis data item konfigurasi (objek) beserta relasinya. | 33 |
| Object schema / Object type | Skema objek / Tipe objek | Struktur data Assets: skema berisi tipe objek dengan atribut. | 33 |
| AQL | AQL (bahasa kueri aset) | Bahasa kueri untuk mencari objek Assets. | 33 |
| Alert | Peringatan | Sinyal dari alat monitoring yang perlu direspons tim operasi. | 34 |
| On-call | Jadwal siaga | Rotasi petugas yang bertanggung jawab merespons alert. | 34 |
| Escalation | Eskalasi | Aturan pemberitahuan berjenjang bila alert tidak direspons. | 34 |
| Heartbeat | Heartbeat | Sinyal berkala dari sistem; ketiadaannya memicu alert. | 34 |
| Virtual service agent | Agen layanan virtual | Asisten percakapan otomatis yang menjawab customer sebelum ke agent manusia. | 35 |

### Product Discovery

| Istilah | Padanan Indonesia | Definisi | Modul |
| --- | --- | --- | --- |
| Idea | Ide | Work item di proyek discovery yang mewakili peluang/usulan produk. | 36 |
| Insight | Wawasan | Bukti (kutipan pelanggan, data, tautan) yang dilekatkan ke idea. | 37 |
| View | Tampilan | Cara melihat kumpulan idea: list, board, matrix, atau timeline. | 36 |
| Formula field | Field formula | Field yang nilainya dihitung dari field lain (mis. skor RICE). | 36 |
| Impact / Effort | Dampak / Usaha | Field penilaian untuk prioritisasi idea. | 36 |
| Roadmap | Peta jalan | View lintas proyek yang dapat dibagikan ke stakeholder. | 36 |
| Delivery tracking | Pelacakan delivery | Tautan idea ke Epic/work item delivery beserta progresnya. | 37 |
| Contributor | Kontributor | Peran discovery yang dapat memberi insight/komentar tanpa mengubah idea. | 37 |

## Baseline repo & gap

Repo saat ini adalah template Go (module `app`, Gin, Clean Architecture, GORM/PostgreSQL, logrus, JWT) dengan SPA SvelteKit 5 yang di-embed ke binary. Baru ada fitur **auth** (register/login) dan **user** (profil, daftar pengguna). Seluruh domain Jira (proyek, work item, workflow, board/sprint, komentar, lampiran, label, notifikasi, aktivitas) belum ada, baik di backend maupun frontend.

### Kondisi saat ini

| Area | Kondisi saat ini | Gap terhadap PRD | Modul pemilik |
| --- | --- | --- | --- |
| Arsitektur backend | Entry `cmd/api/main.go` → `internal/app/app.go`; interface `Feature{Name(); RegisterRoutes(rg)}`; slice features = `[auth, user]` di `/api/v1`; pola `internal/features/<x>/{module.go, delivery/http/{handler,dto}, usecase}`; repository shared dibuat sekali di `app.go`. | Pola sudah tepat dan dipertahankan; perlu puluhan feature module baru. | 27 |
| Endpoint publik | `GET /health` (statis, tanpa cek DB); `POST /api/v1/auth/register` (email, username 3-20, password min 6, name); `POST /api/v1/auth/login` (email+password → user + JWT). | Tidak ada logout, refresh/sesi, reset kata sandi, verifikasi email, 2SV, API token. Health tidak memeriksa DB. | 01, 27 |
| Endpoint terproteksi | `GET/PUT /api/v1/users/profile` (PUT hanya `name`); `GET /api/v1/users` (paginasi + filter id/email/username/name/status/gender/role/provider). | `GET /users` dapat diakses semua pengguna login (tidak ada cek admin); `AuthMiddleware` dipasang inline per-route. | 02, 03 |
| Middleware | Global: Logger → Recovery → CORS → Language (Accept-Language `id` → id, lainnya en). Per-route: `AuthMiddleware` (Bearer JWT, `c.Set("sess", *jwt.Claims)` dengan `UserID`). | Tidak ada RBAC/role middleware, rate limiting, request-id/tracing. | 03, 23, 27 |
| Format respons & i18n | Respons seragam `{error, status, message, data, errors}`; galat validasi `map[field][]string`; pesan id/en di `internal/shared/constants`; DTO memakai `Validate(lang)` manual. | Pola dipertahankan; perlu katalog i18n terpusat juga untuk label UI dan terminologi (VISI-002). | 27, 21 |
| Entitas & repository | `entity.User` (UUID string, email/username unik, password, name, phone, status, birth_date, gender, role, provider, is_active, soft delete); `FilterUser`; `UserRepository` (Create, GetByID, GetByEmail, GetByUsername, Update, Delete, List). Domain error termasuk sisa template `ErrProductNotFound`. | Belum ada entitas project, workspace/organization, issue, issue type, status/workflow, priority, sprint, board, comment, attachment, label, watcher, audit log, notification, membership/role. | 28 |
| Paket pendukung | `pkg/jwt` (Generate/Validate, claims `UserID`), `pkg/crypto/password`, `pkg/pagination.go` (offset, order `created_at DESC`). | Pagination offset saja; perlu keyset/cursor untuk daftar besar. Klaim JWT tidak memuat role yang terverifikasi. | 01, 27 |
| Migrasi database | Satu migrasi `001_create_users_table` (golang-migrate). `id VARCHAR(36) PRIMARY KEY DEFAULT VARCHAR(36)` adalah SQL tidak valid; kolom phone, status, birth_date, gender, role, provider tidak ada padahal dipakai entity/DTO/filter. Working tree punya perubahan uncommitted pada `001...up.sql`. | Tidak ada runner migrasi otomatis; target Makefile menunjuk `./migration/` (salah ketik); docker-compose me-mount `./migrations` ke `initdb.d` sehingga `down.sql` ikut tereksekusi. | 27, 28 |
| Frontend: rute | SvelteKit (adapter-static, `ssr=false`), hanya `+layout.svelte` dan `+layout.ts` (route guard). `(guest)/+layout.svelte`, `(guest)/+page.svelte`, dan `+error.svelte` terhapus di working tree. Guard mengacu `/auth/login`, `/`, `/admin`, `/profile` yang belum ada filenya. | Saat ini tidak ada halaman sama sekali; belum ada login, register, profil, proyek, work item, board. | 01, 21 |
| Frontend: fitur | `lib/features/auth` lengkap (session store rune, TTL 12 jam sliding, localStorage `jari_session`, route guard + test); `lib/features/admin` berisi file kosong 0 byte; `lib/api/api.ts` hasil generator masih memakai `first_name/last_name` (stale). | Belum ada use case login/register, AuthRepository, halaman; tipe `SessionUser` usang. Belum ada library drag-and-drop untuk board. | 01, 10, 21 |
| Frontend: komponen | shadcn-svelte (data-table, sidebar, dialog/sheet/drawer, command, select, calendar, breadcrumb, tabs, dropdown/context menu, resizable, form, chart, attachment, message, dsb.), Tailwind v4, bits-ui, mode-watcher (dark mode), tabler icons, vitest + playwright. | Fondasi UI memadai untuk Fase 1; konvensi string UI Bahasa Indonesia sudah ada, perlu diperluas ke i18n ID/EN. | 21, 27 |
| Konfigurasi | `config.Load()` dari env (dipanggil berulang): SERVER_*, DATABASE_URL atau DB_*, JWT_SECRET (default `your-secret-key`), LOG_LEVEL; `.env.example` memakai `DB_PASS` sedangkan docker-compose memakai `DB_PASSWORD` dan `DB_NAME=yopatungan`. | Default rahasia lemah, nama variabel tidak konsisten, sisa nama proyek lain; Makefile wajib `.env`. | 27 |
| Packaging & deploy | Dockerfile multi-stage (bun build → go build dengan embed → alpine); docker-compose: postgres 15, redis 7 (tidak dipakai kode), app; container/network `yopatungan-*`; tanpa healthcheck. Webui embed dengan cache immutable untuk `_app/immutable/*`. | Binary tidak menjalankan migrasi; tidak ada CI/CD; nama warisan template. | 27 |
| API docs | swag → `docs/{docs.go,swagger.json,swagger.yaml}`; UI Scalar via CDN di `/swagger/index.html`; judul "Backend API v1.0". | Dokumen kemungkinan stale (first_name/last_name); judul generik. | 23 |
| CORS | `Allow-Origin *` + `Allow-Credentials true` (kombinasi tidak valid untuk permintaan credentialed); metode POST, OPTIONS, GET, PUT, DELETE. | Tidak ada PATCH; origin belum dibatasi lewat konfigurasi. | 27 |
| Logging & observabilitas | logrus dengan formatter teks berwarna; GORM logger mode Info (verbose); graceful shutdown SIGINT/SIGTERM. | Tidak ada log terstruktur JSON produksi, metrik, tracing, audit log. | 27, 22 |
| Infrastruktur yang belum ada | — | Mail/SMTP, penyimpanan berkas (lokal/S3), job/queue/worker/cron, cache, websocket/SSE, rate limiting, RBAC, audit log, request-id, health DB. | 08, 19, 27 |
| Pengujian | Unit test handler/usecase/repo (testify, sqlmock, mockery), jwt, crypto, webui, apidocs, route-guard, session storage. | Tidak ada integration test terhadap PostgreSQL nyata maupun e2e. | 27 |
| Dokumentasi | README template "Go Project Template"; `CLAUDE.md` backend & frontend; skill proyek di `.claude/skills/`; swagger. | Belum ada ADR/ERD; PRD ini menjadi dokumen produk pertama. | 00 |

### Remediasi Sprint 0

Pekerjaan berikut diselesaikan sebelum fitur Fase 1 dimulai, karena menjadi prasyarat kebenaran data, keamanan, dan identitas produk:

- **S0-01 Perbaiki migrasi 001** — ganti `DEFAULT VARCHAR(36)` dengan `DEFAULT gen_random_uuid()::text` (pgcrypto sudah diaktifkan) atau hapus default karena UUID diisi hook `BeforeCreate`; selesaikan perubahan uncommitted di `001...up.sql`.
- **S0-02 Sinkronkan skema users** — tambah migrasi `002` untuk kolom `phone`, `status`, `birth_date`, `gender`, `role`, `provider` (atau perbaiki 001 bila belum pernah dirilis), beserta indeks yang dipakai filter.
- **S0-03 Runner migrasi yang aman** — perbaiki path Makefile `./migration/` → `./migrations/`; hentikan mount `./migrations` ke `docker-entrypoint-initdb.d` (mencegah `down.sql` tereksekusi); sediakan perintah migrasi terkontrol (subcommand binary atau job terpisah) dan pemeriksaan versi skema saat start.
- **S0-04 Konfigurasi aman** — tolak start (fail-fast) bila `JWT_SECRET` kosong/default di luar mode development; selaraskan `DB_PASS`/`DB_PASSWORD`; baca `ENV`; muat konfigurasi sekali dan injeksikan; hapus sisa `yopatungan` dari compose, nama container, dan database.
- **S0-05 CORS & metode HTTP** — allowlist origin dari konfigurasi (bukan `*` bersama credentials); izinkan `PATCH`.
- **S0-06 Otorisasi minimum** — batasi `GET /api/v1/users` untuk admin; tambahkan middleware peran berbasis klaim terverifikasi dari database sebagai jembatan sampai modul [03](03-izin-peran-keamanan.md) selesai.
- **S0-07 Bersihkan warisan template** — hapus `ErrProductNotFound`, nasib redis sudah diputuskan: dipakai untuk queue dan tetap di compose ([ADR-0004](../adr/0004-job-latar-dan-queue.md)), ganti README template, ganti judul swagger "Backend API" dengan nama produk.
- **S0-08 Regenerasi kontrak API** — jalankan `make swag` dan regenerasi `frontend/src/lib/api/api.ts`; ganti `first_name/last_name` menjadi `name` di `SessionUser`; verifikasi interceptor Bearer dan penanganan 401 di `client.ts`.
- **S0-09 Pulihkan kerangka frontend** — putuskan apakah `(guest)` dan `+error.svelte` yang terhapus dipulihkan atau diganti; buat halaman login, register, home, profil minimal; isi atau hapus stub `lib/features/admin`.
- **S0-10 Operasional dasar** — `/health` memeriksa koneksi DB (liveness vs readiness), request-id per permintaan, healthcheck Docker, level GORM logger mengikuti env, log JSON di produksi.
- **S0-11 CI** — pipeline lint + unit test Go + test frontend + build image untuk setiap PR.
- **S0-12 Identitas produk** — terapkan nama produk "Jari" dan aset brand sendiri di UI, email, user agent, swagger (VISI-001); tambahkan laporan lisensi dependensi.
- **S0-13 Fondasi i18n & terminologi** — siapkan katalog label UI terpusat ID/EN dengan istilah "work item"/"space" (VISI-002), memperluas pola konstanta galat id/en yang sudah ada.
- **S0-14 Keputusan teknis Fase 1** — pilih library drag-and-drop untuk board, strategi penyimpanan berkas lampiran, transport SMTP, dan mekanisme job latar; dicatat sebagai ADR (rinciannya di modul [27](27-nfr-arsitektur-platform.md)). Hasil: [ADR-0001 sampai ADR-0004](../adr/README.md) berstatus Accepted — drag-and-drop dibangun sendiri (`lib/dnd`), lampiran di S3-kompatibel, email keluar lewat SMTP, dan queue di Redis (`asynq`).

## Peta fase delivery

| Fase | Nama | Tujuan | Cakupan kunci | Modul utama |
| --- | --- | --- | --- | --- |
| Fase 1 | MVP pelacak kerja tim inti | Tim kecil dapat mengelola pekerjaan harian end-to-end di satu produk self-hosted. | Auth & sesi, proyek, work item, workflow default, board/backlog/sprint dasar, komentar & lampiran, pencarian dasar, notifikasi in-app, navigasi global, identitas produk & terminologi. Umumnya fitur P0 modul inti. | [00](00-visi-ruang-lingkup-glosarium-baseline.md), [01](01-identitas-autentikasi-sesi.md), [04](04-proyek.md), [06](06-work-item-inti.md), [07](07-editor-konten-kaya.md), [08](08-komentar-lampiran-aktivitas.md), [09](09-workflow-status-transisi.md), [10](10-board-kolom-kartu.md), [11](11-backlog-sprint.md), [19](19-notifikasi-email.md), [21](21-navigasi-ux-preferensi.md), [27](27-nfr-arsitektur-platform.md), [28](28-model-data-inti.md) |
| Fase 2 | Konfigurabilitas & visibilitas tim | Tim dapat menyesuaikan proses dan melihat kemajuan. | JQL, navigator & filter tersimpan, editor workflow, custom field & layar, board lanjutan, versi & komponen, time tracking, laporan agile, email, realtime; layanan feature flag & kill switch. | [05](05-konfigurasi-work-item-skema.md), [09](09-workflow-status-transisi.md), [10](10-board-kolom-kartu.md), [12](12-jql-bahasa-kueri.md), [13](13-pencarian-filter-ekspor.md), [14](14-pelacakan-waktu-estimasi.md), [15](15-rilis-versi-komponen.md), [16](16-laporan-insight.md), [19](19-notifikasi-email.md) |
| Fase 3 | Skala organisasi, otomasi & integrasi | Organisasi dengan banyak tim dapat mengatur akses, mengotomasi, dan mengintegrasikan. | Permission scheme penuh, manajemen pengguna lanjut, dashboard, otomasi, tampilan proyek, REST API publik & webhook, impor/ekspor, proyek bisnis & formulir, administrasi sistem, mapping tier & batas plan. | [02](02-pengguna-grup-organisasi.md), [03](03-izin-peran-keamanan.md), [17](17-dashboard-gadget.md), [18](18-timeline-rencana-tampilan-proyek.md), [20](20-otomasi.md), [22](22-administrasi-sistem.md), [23](23-api-publik-webhook.md), [25](25-integrasi-impor-ekspor-ekstensi.md), [26](26-kerja-bisnis-formulir-persetujuan.md) |
| Fase 4 | Perencanaan lanjut, DevOps & ekosistem | Perencanaan lintas tim dan keterhubungan dengan alat pengembangan serta identitas perusahaan. | Plans, Teams, Goals, hierarki di atas Epic, GitHub/GitLab, enterprise SSO/SCIM, OAuth 2.0, Slack/Teams, ekstensi, AI opsional, sandbox/release tracks. | [18](18-timeline-rencana-tampilan-proyek.md), [24](24-integrasi-devops.md), [25](25-integrasi-impor-ekspor-ekstensi.md), [01](01-identitas-autentikasi-sesi.md), [02](02-pengguna-grup-organisasi.md), [23](23-api-publik-webhook.md) |
| Fase 5 | Service Management inti | Tim layanan dapat menerima dan menyelesaikan permintaan customer dengan target layanan. | Service project, request type, forms, help center & portal, customers & organizations, customer notifications, email channel & widget, queues, SLA, knowledge base, approvals, canned responses, CSAT, laporan JSM. | [29](29-jsm-service-project-request-type-formulir.md), [30](30-jsm-help-center-portal-pelanggan-kanal.md), [31](31-jsm-antrian-sla-knowledge-base-laporan.md), [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md) (approvals) |
| Fase 6 | Product Discovery | Tim produk dapat mengumpulkan, memprioritaskan, dan menghubungkan ide ke delivery. | Ideas, fields & formula, views, sharing, insights, delivery tracking, roles. | [36](36-jpd-ideas-field-views.md), [37](37-jpd-insights-delivery-peran-integrasi.md) |
| Fase 7 | ITSM lanjut & Operations | Organisasi TI dapat menjalankan praktik ITSM penuh dan respons operasi. | Incident/problem/change, services & deployment gating, Assets/CMDB & AQL, alerts/on-call/escalation, virtual service agent & AI JSM, Customer Service Management, chat. | [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md), [33](33-jsm-assets-cmdb.md), [34](34-jsm-operations-alerting-on-call.md), [35](35-jsm-ai-otomasi-api-admin-csm.md), [29](29-jsm-service-project-request-type-formulir.md) (chat) |

Aturan ibu jari penempatan fase: P0 inti → Fase 1-2; P1 → Fase 2-3; P2/premium/enterprise → Fase 3-4; tier `jsm` → Fase 5 (inti) atau Fase 7 (ITSM/Assets/Ops/AI); tier `jpd` → Fase 6; `dc-only-or-deprecated` → "Fase —" (rujukan saja). Penyimpangan dari aturan ini wajib dijelaskan di "Ringkasan prioritas & fase" modul terkait.

## Konvensi dokumen

### Struktur file modul

Setiap file modul memiliki urutan bagian tetap: judul `# <id> — <judul>` dengan paragraf cakupan dan modul terkait; "Persona terkait"; "User story modul"; "Ringkasan prioritas & fase"; "Fitur" (dikelompokkan per sub-area); "Pertanyaan terbuka"; dan "Di luar cakupan & ekosistem" sebagai bagian terakhir.

### Template fitur

Setiap fitur ditulis sebagai heading `### <PREFIX>-<NNN> <nama fitur persis seperti inventaris>` dengan isi:

- **Deskripsi:** 1-3 kalimat.
- **Prioritas / Fase / Tier:** `<P0-P3> / <Fase 1-7 atau Fase —> / <tier>`.
- **Aktor & izin:** siapa yang memakai dan permission/peran yang mengontrol.
- **Aturan perilaku (acceptance criteria):** daftar bernomor, minimal satu butir per sub-fitur inventaris, masing-masing dapat diuji.
- **Team-managed vs company-managed:** perbedaan perilaku, atau "Tidak ada perbedaan."
- **Catatan & ketergantungan:** rujukan silang, keputusan terbuka, catatan migrasi dari baseline.
- **Alias inventaris:** (opsional, ditulis tepat di bawah heading sebelum Deskripsi) nama fitur inventaris lain yang digabung ke heading ini, dipisah titik koma.

Nomor ID berurutan sesuai urutan kemunculan di file dan tidak dipakai ulang; fitur yang dihapus di revisi berikutnya ditandai "dihapus" alih-alih nomornya diberikan ke fitur lain.

### Arti prioritas

| Prioritas | Arti |
| --- | --- |
| P0 | Wajib untuk fase rilisnya; tanpa ini produk tidak dapat dipakai untuk alur utama atau melanggar keamanan/hukum. |
| P1 | Penting untuk paritas yang diharapkan pengguna Jira; dapat menyusul satu fase setelah inti. |
| P2 | Bernilai tambah atau diferensiasi plan lebih tinggi; dikerjakan setelah P0/P1 di area yang sama. |
| P3 | Opsional, niche, atau berisiko rendah bila ditunda; dikerjakan bila kapasitas tersedia. |

### Arti tier

| Tier | Arti |
| --- | --- |
| `free-standard` | Tersedia di plan Free/Standard Jira Cloud; di clone aktif secara default. |
| `premium` | Fitur plan Premium; di clone dapat digating lewat VISI-003. |
| `enterprise` | Fitur plan Enterprise (skala organisasi, kepatuhan); dapat digating lewat VISI-003. |
| `platform` | Kapabilitas lintas produk/infrastruktur, tidak terikat plan. |
| `jsm` | Bagian Jira Service Management. |
| `jpd` | Bagian Jira Product Discovery. |
| `dc-only-or-deprecated` | Hanya ada di Data Center atau sudah dihentikan di Cloud; dicatat sebagai rujukan, tidak dibangun. |

### Arti fase dan tanda ⚠

- **Fase 1-7** mengikuti [Peta fase delivery](#peta-fase-delivery); "Fase —" berarti tidak dibangun.
- **⚠** pada heading inventaris atau di awal butir aturan berarti perilaku pembanding Jira Cloud belum terverifikasi oleh riset; butir tersebut harus divalidasi (dokumentasi resmi atau uji di instance pembanding) sebelum diimplementasikan, dan boleh disesuaikan oleh pemilik produk.
- **⚠ pada fitur di akhir bagian Fitur** menandai fitur yang ditambahkan penulis di luar inventaris riset.

### Prefix ID per modul

| Modul | Prefix | File |
| --- | --- | --- |
| 00 | `VISI` | [Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md) |
| 01 | `AUTH` | [Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) |
| 02 | `USR` | [Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) |
| 03 | `PERM` | [Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) |
| 04 | `PRJ` | [Proyek: Tipe, Template, Pengaturan & Siklus Hidup](04-proyek.md) |
| 05 | `CFG` | [Konfigurasi Work Item: Tipe, Field, Layar & Skema](05-konfigurasi-work-item-skema.md) |
| 06 | `ISS` | [Work Item Inti: Buat, Lihat, Ubah & Operasi](06-work-item-inti.md) |
| 07 | `EDT` | [Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) |
| 08 | `CMT` | [Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) |
| 09 | `WF` | [Workflow, Status & Transisi](09-workflow-status-transisi.md) |
| 10 | `BRD` | [Board: Kolom, Kartu, Swimlane & Kanban](10-board-kolom-kartu.md) |
| 11 | `SPR` | [Backlog & Sprint (Scrum)](11-backlog-sprint.md) |
| 12 | `JQL` | [Bahasa Kueri (JQL-kompatibel)](12-jql-bahasa-kueri.md) |
| 13 | `SRC` | [Pencarian, Navigator, Filter Tersimpan & Ekspor](13-pencarian-filter-ekspor.md) |
| 14 | `TIME` | [Pelacakan Waktu, Worklog & Estimasi](14-pelacakan-waktu-estimasi.md) |
| 15 | `REL` | [Rilis (Versi) & Komponen](15-rilis-versi-komponen.md) |
| 16 | `RPT` | [Laporan & Insight](16-laporan-insight.md) |
| 17 | `DSH` | [Dashboard & Gadget](17-dashboard-gadget.md) |
| 18 | `PLAN` | [Timeline, Rencana Lintas Proyek & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) |
| 19 | `NTF` | [Notifikasi In-app & Email](19-notifikasi-email.md) |
| 20 | `AUTO` | [Otomasi Berbasis Aturan](20-otomasi.md) |
| 21 | `NAV` | [Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) |
| 22 | `ADM` | [Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) |
| 23 | `API` | [REST API Publik & Webhook](23-api-publik-webhook.md) |
| 24 | `DEV` | [Integrasi Pengembangan & DevOps](24-integrasi-devops.md) |
| 25 | `INT` | [Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) |
| 26 | `BIZ` | [Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) |
| 27 | `NFR` | [Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) |
| 28 | `DM` | [Model Data Inti & Katalog Event Domain](28-model-data-inti.md) |
| 29 | `JSM-PRJ` | [Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) |
| 30 | `JSM-PRT` | [Service Management: Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) |
| 31 | `JSM-QSK` | [Service Management: Queues, SLA, Knowledge Base, CSAT & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) |
| 32 | `JSM-ITSM` | [Service Management: ITSM (Incident, Problem, Change, Services, Approvals)](32-jsm-itsm-insiden-problem-perubahan-layanan.md) |
| 33 | `JSM-AST` | [Service Management: Assets (CMDB)](33-jsm-assets-cmdb.md) |
| 34 | `JSM-OPS` | [Service Management: Operations (Alerts, On-call, Escalation)](34-jsm-operations-alerting-on-call.md) |
| 35 | `JSM-AI` | [Service Management: Virtual Service Agent, AI, Otomasi, API, Admin & Customer Service Management](35-jsm-ai-otomasi-api-admin-csm.md) |
| 36 | `JPD-IDE` | [Product Discovery: Ideas, Fields, Formula & Views](36-jpd-ideas-field-views.md) |
| 37 | `JPD-INS` | [Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) |

Rujukan silang ditulis sebagai ID fitur (mis. `ISS-001`) disertai tautan relatif ke file modul. Bila ID tujuan belum pasti, rujuk dengan nama fitur inventaris dan nomor modul.

## Fitur

## Identitas produk & terminologi

### VISI-001 Product Naming & Trademark Compliance

- **Deskripsi:** Kepatuhan merek: produk clone memakai nama, logo, dan identitas visual sendiri serta tidak memakai nama, logo, atau branding Atlassian/Jira sebagai identitas produk. Merek pihak lain hanya disebut secara nominatif dan deskriptif (mis. untuk kompatibilitas dan impor).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Product Owner/PM (pemilik produk) dan Legal menetapkan nama dan pedoman merek; Engineering menerapkannya; Admin Situs dapat mengubah nama tampilan site dan logo instansi lewat pengaturan tampilan (modul 22); nama produk dasar tetap tercantum di halaman About dan berkas atribusi lisensi.
- **Aturan perilaku (acceptance criteria):**
  1. Nama produk berasal dari satu sumber konfigurasi build (nama final produk: `Jari`) dan dipakai di judul tab browser, header aplikasi, layar login, manifest/favicon, halaman About, dan judul dokumen OpenAPI; mengganti nama cukup dilakukan di satu tempat tanpa mengubah komponen.
  2. Email keluar memakai nama produk (atau nama tampilan site) pada display name pengirim, subjek, dan footer; tidak ada templat email yang memuat kata "Jira" atau "Atlassian" sebagai nama layanan.
  3. User agent untuk semua permintaan HTTP keluar (webhook, unfurl, integrasi) berformat `<NamaProduk>/<versi> (+<url-proyek>)`; domain dan hostname contoh di dokumentasi serta konfigurasi default memakai `example.com` atau domain milik proyek, tidak pernah `*.atlassian.net` atau domain yang memuat "jira".
  4. Pemeriksaan otomatis di CI memindai bundle SPA, katalog i18n, templat email, dan dokumentasi untuk kata "Jira"/"Atlassian"; kemunculan di luar daftar konteks deskriptif yang diizinkan (importer, disclaimer kompatibilitas, glosarium, PRD) menggagalkan build.
  5. Produk tidak memakai logo, ikon tipe work item, ilustrasi, font merek, atau palet warna Atlassian secara identik; ikon memakai set tabler (baseline) atau aset milik proyek, dan token warna didefinisikan sendiri di tema.
  6. Komponen Atlassian Design System hanya boleh dipakai sesuai lisensi paketnya; aset brand ADS tidak dipakai. Default baseline (shadcn-svelte, bits-ui, tabler icons) tidak bergantung pada ADS, dan setiap penambahan dependensi ADS wajib melalui review lisensi (butir 11).
  7. Istilah generik metodologi dan pelacakan kerja (issue/work item, epic, story, task, bug, sprint, board, backlog, kanban, scrum, workflow) boleh dipakai bebas di UI dan dokumentasi.
  8. ⚠ Bahasa kueri ditampilkan di UI dengan nama sendiri (mis. "Kueri lanjutan"), sedangkan dokumentasi boleh menyatakan sintaksnya kompatibel dengan JQL; penggunaan akronim "JQL" di UI diputuskan setelah status mereknya divalidasi Legal.
  9. Klaim kompatibilitas hanya berbentuk nominatif, mis. "kompatibel dengan format/API Jira", tanpa logo atau gaya visual Jira dan tanpa kalimat yang menyiratkan afiliasi, dukungan, atau sertifikasi; halaman About, README, dan dokumentasi importer memuat disclaimer "Jira dan Atlassian adalah merek dagang Atlassian Pty Ltd; produk ini tidak berafiliasi dengan Atlassian".
  10. Importer menyebut Jira hanya secara deskriptif sebagai sumber data (mis. "Impor dari Jira (CSV/JSON/REST API)"), tanpa logo Jira, dengan disclaimer pada layar pertama wizard impor.
  11. Setiap rilis menghasilkan daftar lisensi dependensi Go dan npm (SBOM) yang dikemas dalam image sebagai berkas pemberitahuan pihak ketiga dan dapat dilihat dari halaman About; build gagal bila ditemukan dependensi dengan lisensi di daftar tolak atau lisensi tidak dikenal yang belum disetujui.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Jira dan Atlassian adalah merek dagang Atlassian Pty Ltd. Pengaturan tampilan site (nama site, logo, warna) ada di modul [22](22-administrasi-sistem.md); importer di modul [25](25-integrasi-impor-ekspor-ekstensi.md); bahasa kueri di modul [12](12-jql-bahasa-kueri.md). Baseline: judul swagger "Backend API" dan nama warisan `yopatungan` harus diganti (Remediasi S0-07, S0-12). Lisensi proyek sendiri dan nama final masih keputusan terbuka.

### VISI-002 Terminology Update (Work items & Spaces)

- **Deskripsi:** Mengadopsi istilah UI Jira Cloud 2025: "issues" menjadi "work items", "issue types" menjadi "work types", dan "projects" menjadi "spaces", diterapkan konsisten di navigasi, dialog, email, dan dokumentasi, sementara kontrak API dan JQL tetap memakai nama lama demi kompatibilitas. Semua label terminologi dipusatkan di layer i18n agar mudah dikonfigurasi.
- **Prioritas / Fase / Tier:** P2 / Fase 1 / platform
- **Aktor & izin:** Semua pengguna melihat terminologi; Admin Situs dapat memilih set terminologi site (butir 9); Integrator/Developer API bergantung pada nama lama di API.
- **Aturan perilaku (acceptance criteria):**
  1. Label UI di top bar, sidebar, breadcrumb, judul halaman, dialog, tombol, empty state, flag/toast, dan pesan galat berbasis UI memakai istilah baru (EN: work item, work type, space; ID: padanan sesuai glosarium) dan diambil dari kunci terminologi terpusat di katalog i18n, bukan string literal di komponen.
  2. Lint frontend menolak string literal "issue"/"project" (dan padanan Indonesianya) di komponen Svelte di luar katalog i18n; bentuk tunggal/jamak dan huruf kapital awal kalimat ditangani oleh kunci terpisah.
  3. Templat email dan notifikasi in-app memakai kunci terminologi yang sama sehingga istilah di email identik dengan UI pada bahasa penerima.
  4. REST API mempertahankan nama sumber daya dan field lama (`issue`, `issues`, `issuetype`, `project`, `projectId`) pada path, parameter, dan payload; perubahan terminologi tidak boleh mengubah kontrak API maupun kode galat mesin.
  5. JQL tetap menerima `project`, `issuetype`, dan alias `type` persis seperti sebelumnya; kueri dan filter tersimpan lama menghasilkan hasil yang sama setelah perubahan istilah. Penambahan alias berbasis istilah baru hanya boleh bila tidak bentrok dengan nama field kustom.
  6. URL dan deep link (mis. `/browse/<KEY>`) tidak berubah ketika set terminologi berganti, sehingga tautan lama tetap valid.
  7. Dokumentasi pengguna dan help panel memakai istilah baru; glosarium dan artikel bantuan menyebut padanan lama satu kali (mis. "work item (sebelumnya issue)"), dan pencarian help menemukan artikel dengan istilah lama maupun baru.
  8. ⚠ Terminologi baru diluncurkan bertahap per site melalui feature flag (VISI-003); instalasi baru memakai istilah baru secara default, dan pergantian set terminologi berlaku tanpa restart setelah SPA dimuat ulang.
  9. ⚠ Admin Situs dapat memilih set terminologi "baru" (work item/space) atau "klasik" (issue/project) untuk site; pilihan ini hanya memengaruhi label UI, email, dan help, tidak memengaruhi API, JQL, atau data.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog i18n dan deteksi bahasa di modul [27](27-nfr-arsitektur-platform.md) (Internationalization (i18n ID/EN)) dan [21](21-navigasi-ux-preferensi.md) (Language, Region & Date/Time Display); API di modul [23](23-api-publik-webhook.md); JQL di modul [12](12-jql-bahasa-kueri.md). Di Fase 1 set terminologi dipilih lewat pengaturan site (butir 9); rollout bertahap berbasis feature flag (butir 8) baru berlaku setelah layanan flag inti VISI-003 tersedia di Fase 2. Butir 8-9 bertanda ⚠ karena keberadaan opt-out di pembanding belum terverifikasi; opsi "klasik" tetap merupakan keputusan produk clone. PRD ini menulis "work item" dan "project/proyek" secara bergantian sesuai glosarium; implementasi mengikuti kunci terminologi. Baseline: konstanta galat id/en sudah ada di `internal/shared/constants`, tetapi label UI frontend masih string Bahasa Indonesia literal; perlu dipindahkan ke katalog (Remediasi S0-13). Padanan Indonesia untuk "space" ("Ruang" vs tetap "Proyek") adalah keputusan terbuka.

## Edisi, tier & feature flag

### VISI-003 Plan Tiers & Feature Flags

- **Deskripsi:** Mekanisme gating fitur sesuai tier Free/Standard/Premium/Enterprise serta rollout bertahap melalui feature flag, termasuk kill switch saat insiden dan preset untuk deployment teregulasi. Untuk clone open-source, gating tier bersifat konfigurasi (simulasi edisi), bukan sistem penagihan.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs (izin global administer) mengelola tier aktif, override flag, kill switch, dan profil deployment; operator server dapat mengatur flag lewat berkas konfigurasi/env; Engineering mendefinisikan flag; semua pengguna terautentikasi membaca flag efektif miliknya.
- **Aturan perilaku (acceptance criteria):**
  1. Layanan feature flag membaca definisi flag dari berkas (dimuat saat start) dan override runtime dari database; setiap flag punya kunci unik, tipe (boolean atau varian string), nilai default, pemilik, dan tanggal kedaluwarsa opsional.
  2. Evaluasi flag mendukung target per tenant/site, per pengguna, dan per persentase pengguna dengan urutan prioritas: kill switch → preset regulated deployment (butir 15) → override pengguna → override tenant → persentase → default berkas; fitur yang tidak tercakup tier aktif (butir 5) tetap nonaktif apa pun hasil evaluasi flag; pengelompokan persentase deterministik (hash kunci flag + ID pengguna) sehingga pengguna yang sama selalu mendapat hasil yang sama.
  3. Flag yang tidak dikenal dievaluasi sebagai nonaktif dan menghasilkan log peringatan; perubahan override di database berlaku paling lambat 60 detik tanpa restart dan tercatat di log audit (pelaku, waktu, nilai lama, nilai baru).
  4. Pemetaan fitur → tier minimum (Free < Standard < Premium < Enterprise) disimpan sebagai konfigurasi terpusat; contoh default: Plans, sandbox, release tracks, archiving, dan IP allowlisting = Premium/Enterprise; project roles dan issue security = Standard ke atas.
  5. Site memiliki satu tier aktif yang dipilih Admin Situs; backend menegakkan gating pada setiap endpoint fitur dan mengembalikan HTTP 403 dengan kode `feature_not_in_plan` beserta pesan i18n, sedangkan UI menyembunyikan atau menonaktifkan entry point dan menampilkan penjelasan plan.
  6. Menurunkan tier tidak menghapus data fitur yang tidak lagi tercakup; data menjadi tersembunyi atau hanya-baca dan aktif kembali utuh saat tier dinaikkan.
  7. ⚠ Batas plan dapat dikonfigurasi per tier, dengan nilai pembanding Free maksimum 10 pengguna dan 2 GB penyimpanan; batas pengguna untuk Standard, Premium, dan Enterprise diisi setelah divalidasi.
  8. Saat batas pengguna tercapai, undangan, sign-up, dan aktivasi pengguna baru ditolak dengan pesan yang menyebut batas dan cara menaikkannya; saat kuota penyimpanan habis, unggahan baru ditolak dengan HTTP 413 berkode `storage_quota_exceeded` dan pesan i18n yang menyebut kuota serta pemakaian saat ini, tanpa memengaruhi berkas yang sudah ada; batas yang bernilai kosong berarti tanpa batas.
  9. Kill switch mematikan fitur seketika untuk semua pengguna dan mengalahkan semua override; endpoint fitur yang dimatikan mengembalikan HTTP 503 dengan kode `feature_disabled`, dan UI menyembunyikan entry point serta menampilkan banner pemberitahuan.
  10. Kill switch juga dapat diaktifkan lewat variabel lingkungan atau berkas konfigurasi tanpa akses database, agar tetap dapat dipakai saat database bermasalah.
  11. Endpoint konfigurasi terautentikasi (mis. `GET /api/v1/config/features`) mengembalikan flag yang sudah dievaluasi untuk pengguna saat ini, tier aktif, dan batas plan yang relevan; permintaan tanpa token mendapat 401, dan respons tidak membocorkan aturan targeting, daftar pengguna, atau tenant lain.
  12. SPA memuat flag efektif saat boot dan memperbaruinya secara berkala atau saat tab kembali fokus; flag yang dibutuhkan sebelum login (mis. sign-up diizinkan) hanya diekspos melalui daftar publik terbatas yang ditentukan eksplisit.
  13. (Opsional) Release tracks: site dapat memilih jalur rilis "berkelanjutan" (fitur baru aktif otomatis) atau "terjadwal" (fitur baru menunggu aktivasi admin atau tanggal rilis yang dijadwalkan); hanya tersedia untuk tier Premium/Enterprise.
  14. (Opsional) Sandbox: Admin Situs dapat membuat lingkungan sandbox terisolasi untuk menguji konfigurasi dan flag, dengan salinan konfigurasi dan data opsional; data dan notifikasi sandbox tidak pernah bocor ke site produksi (email keluar dari sandbox dinonaktifkan secara default).
  15. Profil "regulated deployment" adalah preset flag yang menonaktifkan semua fitur yang memerlukan layanan eksternal atau egress (AI/asisten, smart link unfurl publik, integrasi SaaS pihak ketiga, analytics/telemetri) dan menyembunyikan entry point-nya dari UI; preset tidak dapat di-override oleh pengguna atau admin proyek, dan permintaan keluar ke host di luar allowlist ditolak.
  16. Halaman admin menampilkan daftar flag beserta nilai efektif, sumber nilai (berkas, override, kill switch, preset), tier aktif, dan pemakaian terhadap batas plan.
- **Team-managed vs company-managed:** Gating tier dan flag site berlaku sama untuk keduanya. Toggle fitur per proyek team-managed (mis. backlog, sprint, timeline di halaman Features proyek) adalah mekanisme terpisah di modul [04](04-proyek.md)/[10](10-board-kolom-kartu.md) dan hanya dievaluasi bila flag/tier site mengizinkan fitur tersebut.
- **Catatan & ketergantungan:** Butir 1-3 dan 9-12 (layanan flag inti dan evaluasi, kill switch, endpoint konfigurasi) dimajukan ke Fase 2 dengan prioritas P1; butir 13-14 (release tracks, sandbox) Fase 4; butir 15 dimulai di Fase 3 ketika fitur ber-egress pertama tersedia, dan setiap fitur ber-egress baru (hingga AI di Fase 4/7) wajib terdaftar di preset ini sebelum dirilis (modul [25](25-integrasi-impor-ekspor-ekstensi.md) "SSRF & Outbound Request Guard", modul [07](07-editor-konten-kaya.md) smart link, modul AI). Prompt upgrade dan tampilan batas pengguna ada di modul [21](21-navigasi-ux-preferensi.md) "Plan Editions, User Limits & Upgrade Prompts"; konsol admin di modul [22](22-administrasi-sistem.md); penegakan batas pengguna di modul [02](02-pengguna-grup-organisasi.md); kuota penyimpanan di modul [08](08-komentar-lampiran-aktivitas.md); log audit di modul [27](27-nfr-arsitektur-platform.md). Baseline belum memiliki mekanisme flag apa pun; redis di docker-compose dapat dipakai sebagai cache evaluasi tetapi tidak wajib (database + cache memori sudah cukup). Default tier untuk instalasi self-hosted adalah keputusan terbuka.

## Pertanyaan terbuka

- Apakah nama final "Jari" perlu didaftarkan sebagai merek (mis. di DJKI) sebelum rilis publik pertama?
- Nama apa yang dipakai di UI untuk bahasa kueri JQL-kompatibel, dan apakah akronim "JQL" boleh muncul di UI?
- Apakah UI Bahasa Indonesia memakai "Ruang" sebagai padanan "space" atau tetap "Proyek", dan apakah set terminologi default site adalah "baru" atau "klasik"?
- Tier apa yang aktif secara default pada instalasi self-hosted: semua fitur terbuka (setara Enterprise) atau simulasi plan bertingkat untuk edisi komersial di masa depan?
- Apakah batas plan (mis. 10 pengguna/2 GB untuk Free) benar-benar diterapkan di clone, atau hanya tersedia sebagai opsi bagi penyelenggara hosting?
- Apakah sandbox dibangun sebagai fitur produk (salin site di dalam aplikasi) atau cukup sebagai panduan menjalankan instance kedua dengan impor konfigurasi?
- Apakah produk mengirim telemetri penggunaan anonim, dan bila ya apakah bersifat opt-in (dimatikan secara default)?
- Apakah perubahan uncommitted di `migrations/001_create_users_table.up.sql` dan penghapusan route `(guest)` serta `+error.svelte` di working tree disengaja, sehingga Remediasi S0-01 dan S0-09 dapat dikerjakan di atasnya?

## Di luar cakupan & ekosistem

- **Aplikasi mobile native** iOS/Android dan aplikasi desktop; kebutuhan mobile dipenuhi web responsif.
- **Atlassian Marketplace dan app pihak ketiga** (mis. Tempo, ScriptRunner, Xray, Zephyr, BigPicture, eazyBI) serta runtime Forge/Connect; produk ini hanya menyediakan REST API, webhook, dan titik ekstensi sendiri (modul [25](25-integrasi-impor-ekspor-ekstensi.md)).
- **Layanan Atlassian lintas produk**: Confluence, Bitbucket, Trello, Loom, Compass, Statuspage, Atlas, Atlassian Guard/Access, Atlassian Analytics, Rovo sebagai layanan terkelola, dan Teamwork Graph sebagai produk; hanya padanan fungsional minimal yang dibutuhkan pelacakan kerja yang dibangun.
- **Fitur Data Center-only dan deprecated** (tier `dc-only-or-deprecated`), dicatat di modul masing-masing sebagai rujukan dengan "Fase —".
- **Penagihan dan lisensi berbayar**: VISI-003 hanya menyediakan gating berbasis konfigurasi, tanpa pembayaran, faktur, atau manajemen langganan.
- **Pemakaian merek, logo, dan aset brand Atlassian** dalam bentuk apa pun di luar penyebutan nominatif (VISI-001).
- **Infrastruktur multi-region terkelola** ala Atlassian Cloud (data residency multi-region, Government/Isolated Cloud sebagai layanan); produk ini menyediakan profil regulated deployment untuk instalasi sendiri.
