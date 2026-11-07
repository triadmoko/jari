# PRD — Platform Pelacak Kerja, Layanan & Discovery (nama kerja `jari`)

Dokumen ini adalah indeks dan ringkasan eksekutif Product Requirements Document (PRD) untuk produk clone yang setara fungsional dengan Jira Cloud: Jira terpadu (eks Jira Software dan Jira Work Management, semua plan), Jira Service Management, dan Jira Product Discovery. Produk dibangun di repo ini: backend Go (Clean Architecture, Gin, GORM, PostgreSQL) dengan SPA SvelteKit 5 yang di-embed ke satu binary. Saat PRD ditulis, repo baru memiliki fitur auth dan user.

PRD terdiri atas 38 file modul (`00`–`37`). README ini tidak menggantikan modul. Fungsinya membantu pembaca menemukan modul yang tepat, memahami konvensi, melihat sebaran fase, dan mengumpulkan keputusan yang masih terbuka di satu tempat.

## Daftar isi README

1. [Cara membaca & konvensi](#cara-membaca--konvensi)
2. [Visi singkat & catatan penamaan](#visi-singkat--catatan-penamaan)
3. [Persona](#persona)
4. [Ruang lingkup](#ruang-lingkup)
5. [Daftar modul](#daftar-modul)
6. [Matriks modul × fase](#matriks-modul--fase)
7. [Peta 7 fase delivery](#peta-7-fase-delivery)
8. [Pertanyaan terbuka gabungan](#pertanyaan-terbuka-gabungan)
9. [Risiko](#risiko)
10. [Backlog verifikasi (⚠)](#backlog-verifikasi-)
11. [Glosarium ringkas](#glosarium-ringkas)
12. [Sumber & metodologi](#sumber--metodologi)

## Cara membaca & konvensi

**Urutan baca yang disarankan.**

1. Mulai dari [00 Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md). Modul ini memuat visi, glosarium lengkap, kondisi repo saat ini beserta remediasi Sprint 0, peta fase, dan konvensi dokumen yang mengikat semua modul lain.
2. Baca modul fondasi platform: [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) dan [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md). Hampir semua modul fitur bergantung pada keduanya.
3. Baca modul fitur sesuai fase yang sedang direncanakan (lihat [Matriks modul × fase](#matriks-modul--fase)). Untuk Fase 1, mulailah dari modul 01, 03, 04, 06, 09, 10, 11, dan 21.
4. Gunakan [Pertanyaan terbuka gabungan](#pertanyaan-terbuka-gabungan) sebagai agenda keputusan pemilik produk sebelum fase dimulai.

**Struktur setiap file modul** (urutan tetap): judul `# <id> — <judul>` beserta paragraf cakupan dan modul terkait → *Persona terkait* → *User story modul* → *Ringkasan prioritas & fase* → *Fitur* (dikelompokkan per sub-area) → *Pertanyaan terbuka* → *Di luar cakupan & ekosistem* (selalu bagian terakhir).

**Template fitur.** Setiap fitur ditulis sebagai heading `### <PREFIX>-<NNN> <nama fitur persis seperti di inventaris>` dengan isi:

- **Deskripsi:** 1-3 kalimat.
- **Prioritas / Fase / Tier:** `P0-P3 / Fase 1-7 atau Fase — / tier`.
- **Aktor & izin:** siapa yang memakai fitur, serta permission atau peran yang mengontrolnya.
- **Aturan perilaku (acceptance criteria):** daftar bernomor, minimal satu butir per sub-fitur inventaris, dan setiap butir dapat diuji.
- **Team-managed vs company-managed:** perbedaan perilaku, atau "Tidak ada perbedaan."
- **Catatan & ketergantungan:** rujukan silang ke ID atau modul lain, keputusan terbuka, dan catatan migrasi dari baseline repo.
- **Alias inventaris:** (opsional, tepat di bawah heading) nama fitur inventaris lain yang digabung ke heading ini karena tumpang tindih.

**Prioritas.**

| Kode | Arti |
| --- | --- |
| P0 | Wajib untuk fase rilisnya. Tanpanya alur utama tidak dapat dipakai, atau produk melanggar aspek keamanan/hukum. |
| P1 | Penting untuk paritas yang diharapkan pengguna Jira; boleh menyusul satu fase setelah fitur inti. |
| P2 | Nilai tambah atau diferensiasi plan yang lebih tinggi; dikerjakan setelah P0/P1 di area yang sama. |
| P3 | Opsional atau niche, dengan risiko rendah bila ditunda. |

**Tier** menunjukkan asal fitur di produk pembanding: `free-standard`, `premium`, `enterprise`, `platform` (lintas produk/infrastruktur), `jsm` (Service Management), `jpd` (Product Discovery), dan `dc-only-or-deprecated` (khusus Data Center atau sudah dihentikan; hanya dicatat sebagai rujukan dan tidak dibangun). Di clone, gating tier bersifat konfigurasi dan diatur oleh `VISI-003` ([modul 00](00-visi-ruang-lingkup-glosarium-baseline.md)).

**Fase 1-7** mengikuti [Peta 7 fase delivery](#peta-7-fase-delivery). Tulisan "Fase —" berarti fitur tidak dibangun dan hanya dicatat sebagai rujukan.

**Tanda ⚠** di awal butir aturan berarti perilaku pembanding di Jira Cloud belum terverifikasi oleh riset. Butir tersebut wajib divalidasi sebelum diimplementasikan (lihat [Backlog verifikasi](#backlog-verifikasi-)). Tanda ⚠ pada fitur yang diletakkan di akhir bagian Fitur berarti fitur tersebut ditambahkan penulis di luar inventaris.

**Prefix ID** unik per modul, misalnya `AUTH` (01), `ISS` (06), `JQL` (12), `JSM-PRJ` (29), dan `JPD-IDE` (36). Daftar lengkapnya ada di kolom *Prefix* pada [Daftar modul](#daftar-modul). Nomor ID berurutan sesuai kemunculan di file dan tidak pernah dipakai ulang. Rujukan silang ditulis sebagai ID fitur disertai tautan relatif ke file modulnya.

Detail normatif semua konvensi di atas ada di [00 › Konvensi dokumen](00-visi-ruang-lingkup-glosarium-baseline.md#konvensi-dokumen).

## Visi singkat & catatan penamaan

Produk ini adalah **pelacak kerja open-source dan self-hosted dengan cakupan setara Jira Cloud**. Produk ditujukan bagi tim di Indonesia maupun global yang ingin mengendalikan sendiri data, biaya, dan lokasi hosting-nya. Nilai utamanya:

- **Kendali data & kedaulatan.** Produk dipasang di server sendiri. Tersedia profil *regulated deployment* yang mematikan semua fitur yang membutuhkan egress, untuk instansi pemerintah, BUMN, perbankan, dan organisasi teregulasi.
- **Satu produk, tiga domain.** Pelacakan kerja tim software dan bisnis, Service Management, dan Product Discovery dirilis bertahap dalam 7 fase.
- **Bilingual sejak awal.** UI, pesan galat, dan email tersedia dalam Bahasa Indonesia dan Inggris.
- **Mudah dimigrasikan.** Tersedia bahasa kueri yang kompatibel dengan JQL, REST API dengan payload yang familier, dan importer dari Jira.
- **Sederhana dioperasikan.** Dependensi minimumnya hanya satu binary Go dan PostgreSQL. Cache, queue, object storage, dan search engine bersifat opsional.
- **Terbuka & dapat diperluas.** Tersedia REST API publik, webhook, dan titik ekstensi sendiri, tanpa keterikatan pada ekosistem Marketplace.

**Catatan penamaan.** Produk ini **belum memiliki nama final**. Seluruh PRD menyebutnya "produk ini" atau "clone", sedangkan `jari` hanya nama kerja repo. Nama dan merek Atlassian (Jira, Confluence, Opsgenie, Rovo, dan lainnya) tidak dipakai sebagai nama produk, logo, nama fitur di UI, maupun identitas visual. Merek tersebut hanya disebut secara nominatif dan deskriptif, misalnya untuk menjelaskan acuan paritas, kompatibilitas format/API, atau importer. Istilah generik metodologi (work item, Epic, Sprint, Backlog, Board, Workflow, Request type, SLA, dan sebagainya) dipakai apa adanya. Aturan rinci ada di `VISI-001 Product Naming & Trademark Compliance` dan `VISI-002 Terminology Update` di [modul 00](00-visi-ruang-lingkup-glosarium-baseline.md). Jira dan Atlassian adalah merek dagang Atlassian Pty Ltd. Produk ini tidak berafiliasi dengan Atlassian.

## Persona

Setiap modul memilih persona yang relevan dari daftar baku berikut.

### Pelacakan kerja (inti)

| Persona | Deskripsi | Modul utama |
| --- | --- | --- |
| **Admin Situs** (Rina, IT/Platform Admin) | Mengelola instalasi: pengguna, grup, peran global, pengaturan sistem, SMTP, storage, backup, keamanan, dan audit. Butuh kontrol terpusat, helper izin, dan log audit yang jelas. | 01, 02, 03, 22, 27 |
| **Admin Proyek / Project Lead** (Budi, Engineering Manager) | Membuat dan mengonfigurasi proyek: anggota dan peran, workflow, tipe work item, field, board, komponen, dan rilis. Ingin bisa mengonfigurasi tanpa harus meminta Admin Situs. | 04, 05, 09, 10, 15 |
| **Scrum Master / Team Lead** (Sari) | Mengelola backlog, merencanakan dan menjalankan sprint, memfasilitasi stand-up di board, memantau burndown/velocity, dan menangani hambatan (flag). | 10, 11, 16 |
| **Developer / Anggota Tim** (Andi) | Pengguna harian yang melihat pekerjaan yang ditugaskan, memindahkan kartu, memperbarui status, berkomentar, melampirkan file, mencatat waktu, dan menautkan branch/commit. Mengutamakan kecepatan dan pintasan keyboard. | 06, 07, 08, 14, 21, 24 |
| **Product Owner / PM** (Dewi) | Menulis epic dan story, memprioritaskan backlog, merencanakan rilis dan timeline lintas tim, serta membuat filter tersimpan dan dashboard untuk melaporkan progres. | 11, 13, 15, 17, 18 |
| **Stakeholder / Viewer** (Pak Hendra, Direktur) | Punya akses baca ke dashboard, laporan, timeline, dan status rilis, dan sesekali berkomentar. Butuh ringkasan yang mudah dipahami tanpa pelatihan. | 16, 17, 18 |
| **Pengguna Tim Bisnis** (Maya, HR/Marketing/Ops) | Memakai proyek bisnis untuk task tracking, persetujuan dokumen, atau rekrutmen lewat tampilan list, calendar, board, dan formulir. Tidak memakai konsep sprint. | 18, 26 |
| **Integrator / Developer API** (Fajar) | Membangun integrasi lewat REST API, webhook, API token/OAuth, dan importer. Butuh dokumentasi OpenAPI yang akurat, rate limit yang dapat diprediksi, dan kompatibilitas versi. | 23, 24, 25 |
| **Tamu / Kolaborator Eksternal** (klien, vendor) | Punya akses terbatas ke proyek atau work item tertentu untuk melihat dan berkomentar tanpa melihat data internal lain. Dikendalikan oleh guest access dan issue security. | 02, 03 |

### Service Management

| Persona | Deskripsi | Modul utama |
| --- | --- | --- |
| **Agent Layanan** (Rudi, Service Desk Agent) | Anggota tim layanan berlisensi. Menangani request dari queue, membalas customer atau menulis internal note, memakai canned response dan artikel knowledge base, serta menjaga SLA. | 29, 31, 35 |
| **Customer / Pelapor** (Lina, karyawan atau pelanggan eksternal) | Mengajukan dan memantau request lewat help center, portal, email, atau widget tanpa lisensi agent. Butuh formulir yang jelas, status yang transparan, dan swalayan lewat artikel. | 30 |
| **Manajer Layanan / ITSM** (Tono, IT Service Manager) | Merancang request type, SLA, approval, dan proses incident/problem/change. Memantau laporan, CSAT, on-call, dan Assets/CMDB, serta memastikan kepatuhan praktik ITSM. | 29, 31, 32, 33, 34 |

### Product Discovery

| Persona | Deskripsi | Modul utama |
| --- | --- | --- |
| **Product Manager Discovery** (Nadia) | Mengumpulkan dan memprioritaskan ide dengan field, formula (mis. RICE), dan view list/board/matrix/timeline. Menautkan insight pelanggan, membagikan roadmap, dan menghubungkan ide ke Epic delivery. | 36, 37 |
| **Kontributor / Stakeholder Discovery** (tim sales, support, eksekutif) | Mengirim ide dan insight, memberi vote, dan berkomentar tanpa mengubah ide. Melihat roadmap atau view yang dipublikasikan. | 36, 37 |

## Ruang lingkup

**Dalam cakupan** (dirilis bertahap sesuai [Peta 7 fase delivery](#peta-7-fase-delivery)):

- **Jira terpadu** untuk semua plan Free/Standard/Premium/Enterprise: proyek team-managed dan company-managed, work item dan hierarki, editor teks kaya, komentar/lampiran, workflow, board Scrum/Kanban, backlog dan sprint, JQL, navigator dan filter, time tracking, rilis dan komponen, laporan, dashboard, timeline dan tampilan proyek, Plans/Teams/Goals, otomasi, notifikasi, proyek bisnis dan formulir, administrasi, REST API dan webhook, integrasi DevOps, serta impor/ekspor dan ekstensi (modul 01–26).
- **Service Management:** service project, request type, forms, help center dan portal, customers dan organizations, kanal email/widget/chat, queues, SLA, knowledge base, approvals, canned responses, CSAT, laporan, ITSM (incident, problem, change, services, deployment gating), Assets/CMDB dan AQL, Operations (alerts, on-call, escalation), virtual service agent dan AI opsional, serta Customer Service Management (modul 29–35).
- **Product Discovery:** ideas, fields dan formula, views, sharing dan roadmap, insights, delivery tracking, roles, dan integrasi (modul 36–37).
- **Fondasi lintas produk:** visi, merek, tier dan feature flag (00), kebutuhan non-fungsional dan arsitektur (27), serta model data dan katalog event domain (28).

**Di luar cakupan** (dicatat di bagian penutup setiap modul):

- Aplikasi mobile native iOS/Android dan aplikasi desktop. Kebutuhan mobile dipenuhi web responsif/PWA, dan keputusannya masih terbuka.
- Atlassian Marketplace dan app pihak ketiga (mis. Tempo, ScriptRunner, Xray, Zephyr, BigPicture, eazyBI), juga runtime Forge/Connect. Produk ini hanya menyediakan REST API, webhook, dan titik ekstensi sendiri.
- Layanan Atlassian lintas produk sebagai produk tersendiri: Confluence, Bitbucket, Trello, Loom, Compass, Statuspage, Atlas, Guard/Access, Atlassian Analytics, Rovo sebagai layanan terkelola, dan Teamwork Graph. Yang dibangun hanya padanan fungsional minimal yang dibutuhkan, misalnya entitas Team, Goal, dan knowledge base native.
- Fitur berstatus "Fase —" (58 fitur, sebagian besar ber-tier `dc-only-or-deprecated`). Headingnya tetap ada di modul masing-masing sebagai rujukan.
- Penagihan komersial, faktur, dan manajemen langganan. Gating tier hanya berupa konfigurasi.
- Pemakaian logo, aset brand, atau identitas visual Atlassian dalam bentuk apa pun.
- Infrastruktur multi-region terkelola (data residency multi-region, Government/Isolated Cloud sebagai layanan).

## Daftar modul

Angka di tabel dihitung langsung dari isi file modul, bukan dari laporan penulis:

- **Fitur**: jumlah baris heading yang cocok dengan pola `^### [A-Z][A-Z0-9-]*-[0-9]{3} `.
- **Butir aturan**: jumlah baris yang cocok dengan `^\s*[0-9]+\. `, termasuk daftar bernomor di luar acceptance criteria, misalnya di Pertanyaan terbuka modul 01.
- **⚠**: jumlah kemunculan `⚠ ` di file.

Jalankan ulang perhitungan setelah modul direvisi (lihat [Sumber & metodologi](#sumber--metodologi)).

| No | Modul | Prefix | Fitur | Butir aturan | ⚠ |
| --- | --- | --- | ---: | ---: | ---: |
| 00 | [Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md) | `VISI` | 3 | 36 | 6 |
| 01 | [Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) | `AUTH` | 14 | 157 | 8 |
| 02 | [Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) | `USR` | 49 | 420 | 59 |
| 03 | [Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) | `PERM` | 67 | 438 | 21 |
| 04 | [Proyek: Tipe, Template, Pengaturan & Siklus Hidup](04-proyek.md) | `PRJ` | 68 | 507 | 68 |
| 05 | [Konfigurasi Work Item: Tipe, Field, Layar & Skema](05-konfigurasi-work-item-skema.md) | `CFG` | 61 | 475 | 42 |
| 06 | [Work Item Inti: Buat, Lihat, Ubah & Operasi](06-work-item-inti.md) | `ISS` | 47 | 411 | 57 |
| 07 | [Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) | `EDT` | 30 | 263 | 46 |
| 08 | [Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) | `CMT` | 32 | 264 | 47 |
| 09 | [Workflow, Status & Transisi](09-workflow-status-transisi.md) | `WF` | 45 | 417 | 41 |
| 10 | [Board: Kolom, Kartu, Swimlane & Kanban](10-board-kolom-kartu.md) | `BRD` | 54 | 577 | 94 |
| 11 | [Backlog & Sprint (Scrum)](11-backlog-sprint.md) | `SPR` | 26 | 195 | 28 |
| 12 | [Bahasa Kueri (JQL-kompatibel)](12-jql-bahasa-kueri.md) | `JQL` | 61 | 518 | 73 |
| 13 | [Pencarian, Navigator, Filter Tersimpan & Ekspor](13-pencarian-filter-ekspor.md) | `SRC` | 58 | 433 | 36 |
| 14 | [Pelacakan Waktu, Worklog & Estimasi](14-pelacakan-waktu-estimasi.md) | `TIME` | 49 | 402 | 48 |
| 15 | [Rilis (Versi) & Komponen](15-rilis-versi-komponen.md) | `REL` | 36 | 241 | 37 |
| 16 | [Laporan & Insight](16-laporan-insight.md) | `RPT` | 42 | 335 | 43 |
| 17 | [Dashboard & Gadget](17-dashboard-gadget.md) | `DSH` | 63 | 487 | 53 |
| 18 | [Timeline, Rencana Lintas Proyek & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) | `PLAN` | 66 | 553 | 97 |
| 19 | [Notifikasi In-app & Email](19-notifikasi-email.md) | `NTF` | 45 | 330 | 48 |
| 20 | [Otomasi Berbasis Aturan](20-otomasi.md) | `AUTO` | 74 | 549 | 105 |
| 21 | [Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) | `NAV` | 56 | 556 | 114 |
| 22 | [Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) | `ADM` | 58 | 479 | 110 |
| 23 | [REST API Publik & Webhook](23-api-publik-webhook.md) | `API` | 74 | 702 | 63 |
| 24 | [Integrasi Pengembangan & DevOps](24-integrasi-devops.md) | `DEV` | 32 | 207 | 30 |
| 25 | [Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) | `INT` | 80 | 605 | 130 |
| 26 | [Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) | `BIZ` | 48 | 324 | 75 |
| 27 | [Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) | `NFR` | 24 | 221 | 11 |
| 28 | [Model Data Inti & Katalog Event Domain](28-model-data-inti.md) | `DM` | 15 | 148 | 5 |
| 29 | [Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) | `JSM-PRJ` | 70 | 668 | 125 |
| 30 | [Service Management: Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) | `JSM-PRT` | 50 | 615 | 99 |
| 31 | [Service Management: Queues, SLA, Knowledge Base, CSAT & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) | `JSM-QSK` | 47 | 451 | 88 |
| 32 | [Service Management: ITSM (Incident, Problem, Change, Services, Approvals)](32-jsm-itsm-insiden-problem-perubahan-layanan.md) | `JSM-ITSM` | 52 | 564 | 76 |
| 33 | [Service Management: Assets (CMDB)](33-jsm-assets-cmdb.md) | `JSM-AST` | 48 | 526 | 123 |
| 34 | [Service Management: Operations (Alerts, On-call, Escalation)](34-jsm-operations-alerting-on-call.md) | `JSM-OPS` | 47 | 486 | 85 |
| 35 | [Service Management: Virtual Service Agent, AI, Otomasi, API, Admin & Customer Service Management](35-jsm-ai-otomasi-api-admin-csm.md) | `JSM-AI` | 67 | 533 | 116 |
| 36 | [Product Discovery: Ideas, Fields, Formula & Views](36-jpd-ideas-field-views.md) | `JPD-IDE` | 59 | 703 | 99 |
| 37 | [Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) | `JPD-INS` | 44 | 452 | 67 |
| | **Total (38 modul)** | | **1861** | **16248** | **2473** |

Ringkasan prioritas seluruh PRD: **P0 499**, **P1 623**, **P2 477**, **P3 262** (total 1.861 fitur). Sebaran tier: `free-standard` 927, `jsm` 381, `platform` 316, `jpd` 104, `premium` 57, `dc-only-or-deprecated` 54, dan `enterprise` 22. Fitur yang menyebut lebih dari satu tier dihitung di setiap tier yang disebut.

Pengelompokan modul:

- **Fondasi:** 00 (visi & konvensi), 27 (NFR & arsitektur), 28 (model data & event).
- **Identitas & akses:** 01, 02, 03.
- **Inti pelacakan kerja:** 04–11 (proyek, konfigurasi, work item, editor, komentar, workflow, board, backlog/sprint).
- **Visibilitas & perencanaan:** 12–18 (JQL, pencarian, waktu, rilis, laporan, dashboard, timeline/plans).
- **Komunikasi, otomasi & UX:** 19, 20, 21.
- **Administrasi & ekosistem:** 22–26 (admin sistem, API & webhook, DevOps, integrasi/impor/ekstensi, kerja bisnis & formulir).
- **Service Management:** 29–35.
- **Product Discovery:** 36–37.

## Matriks modul × fase

Setiap sel berisi jumlah fitur modul yang dijadwalkan pada fase tersebut, dihitung dari baris `**Prioritas / Fase / Tier:**` setiap fitur. Tanda **●** menandai fase utama modul, yaitu fase dengan jumlah fitur terbanyak di luar "Fase —". Kolom F— berisi fitur rujukan yang tidak dibangun.

| Modul | F1 | F2 | F3 | F4 | F5 | F6 | F7 | F— | Fase utama |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| [00](00-visi-ruang-lingkup-glosarium-baseline.md) | **2** ● |  | 1 |  |  |  |  |  | Fase 1 |
| [01](01-identitas-autentikasi-sesi.md) | **6** ● | 3 | 3 | 2 |  |  |  |  | Fase 1 |
| [02](02-pengguna-grup-organisasi.md) | 12 | 2 | **17** ● | **17** ● |  |  |  | 1 | Fase 3, Fase 4 |
| [03](03-izin-peran-keamanan.md) | 18 | 15 | **21** ● | 11 | 1 |  |  | 1 | Fase 3 |
| [04](04-proyek.md) | 19 | **24** ● | 15 | 8 |  |  |  | 2 | Fase 2 |
| [05](05-konfigurasi-work-item-skema.md) | 11 | **30** ● | 10 | 5 |  |  |  | 5 | Fase 2 |
| [06](06-work-item-inti.md) | **18** ● | 16 | 8 | 4 |  |  |  | 1 | Fase 1 |
| [07](07-editor-konten-kaya.md) | 9 | **13** ● | 5 | 2 |  |  |  | 1 | Fase 2 |
| [08](08-komentar-lampiran-aktivitas.md) | **10** ● | **10** ● | 6 | 4 | 1 |  |  | 1 | Fase 1, Fase 2 |
| [09](09-workflow-status-transisi.md) | 10 | **20** ● | 10 | 3 |  |  |  | 2 | Fase 2 |
| [10](10-board-kolom-kartu.md) | 17 | **27** ● | 9 |  |  |  |  | 1 | Fase 2 |
| [11](11-backlog-sprint.md) | **11** ● | 10 | 4 | 1 |  |  |  |  | Fase 1 |
| [12](12-jql-bahasa-kueri.md) | **21** ● | 19 | 13 | 7 |  |  |  | 1 | Fase 1 |
| [13](13-pencarian-filter-ekspor.md) | 12 | **22** ● | 12 | 8 |  |  |  | 4 | Fase 2 |
| [14](14-pelacakan-waktu-estimasi.md) | 3 | **25** ● | 13 | 7 |  |  |  | 1 | Fase 2 |
| [15](15-rilis-versi-komponen.md) |  | **20** ● | 10 | 5 |  |  |  | 1 | Fase 2 |
| [16](16-laporan-insight.md) | 6 | 9 | **17** ● | 8 |  |  |  | 2 | Fase 3 |
| [17](17-dashboard-gadget.md) |  |  | **37** ● | 23 |  |  |  | 3 | Fase 3 |
| [18](18-timeline-rencana-tampilan-proyek.md) |  | 16 | 10 | **36** ● |  |  |  | 4 | Fase 4 |
| [19](19-notifikasi-email.md) | 8 | **15** ● | 12 | 5 |  |  |  | 5 | Fase 2 |
| [20](20-otomasi.md) |  |  | **62** ● | 11 |  |  |  | 1 | Fase 3 |
| [21](21-navigasi-ux-preferensi.md) | **17** ● | **17** ● | 11 | 8 |  |  |  | 3 | Fase 1, Fase 2 |
| [22](22-administrasi-sistem.md) | 5 | 7 | 18 | **24** ● |  |  |  | 4 | Fase 4 |
| [23](23-api-publik-webhook.md) | 14 | 17 | **22** ● | 19 |  |  |  | 2 | Fase 3 |
| [24](24-integrasi-devops.md) |  |  |  | **30** ● |  |  | 1 | 1 | Fase 4 |
| [25](25-integrasi-impor-ekspor-ekstensi.md) |  | 1 | 19 | **51** ● |  | 1 | 3 | 5 | Fase 4 |
| [26](26-kerja-bisnis-formulir-persetujuan.md) | 5 | 4 | **32** ● | 7 |  |  |  |  | Fase 3 |
| [27](27-nfr-arsitektur-platform.md) | **10** ● | 8 | 4 | 2 |  |  |  |  | Fase 1 |
| [28](28-model-data-inti.md) | **10** ● | 2 | 1 | 2 |  |  |  |  | Fase 1 |
| [29](29-jsm-service-project-request-type-formulir.md) |  |  |  |  | **68** ● |  | 1 | 1 | Fase 5 |
| [30](30-jsm-help-center-portal-pelanggan-kanal.md) |  |  |  |  | **46** ● |  | 4 |  | Fase 5 |
| [31](31-jsm-antrian-sla-knowledge-base-laporan.md) |  |  |  |  | **42** ● |  | 5 |  | Fase 5 |
| [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md) |  |  |  |  | 8 |  | **44** ● |  | Fase 7 |
| [33](33-jsm-assets-cmdb.md) |  |  |  |  |  |  | **47** ● | 1 | Fase 7 |
| [34](34-jsm-operations-alerting-on-call.md) |  |  |  |  |  |  | **46** ● | 1 | Fase 7 |
| [35](35-jsm-ai-otomasi-api-admin-csm.md) |  |  |  |  | 27 |  | **37** ● | 3 | Fase 7 |
| [36](36-jpd-ideas-field-views.md) |  |  |  |  |  | **59** ● |  |  | Fase 6 |
| [37](37-jpd-insights-delivery-peran-integrasi.md) |  |  |  |  |  | **44** ● |  |  | Fase 6 |
| **Total** | **254** | **352** | **402** | **310** | **193** | **104** | **188** | **58** | |

Cara membaca matriks:

- Modul inti Fase 1 (01, 06, 11, 12, 27, 28) memuat P0 yang menjadi syarat MVP. Sebagian besar fitur modul 04, 05, 09, 10, dan 21 baru masuk Fase 2. Artinya Fase 1 hanya memakai konfigurasi default dari modul-modul tersebut (template proyek, workflow To Do/In Progress/Done, board dasar).
- Modul 12 (JQL) punya 21 fitur Fase 1 karena pencarian dasar dan navigator proyek membutuhkan parser serta kompilasi kueri ke SQL sejak MVP. JQL lengkap yang terlihat pengguna tetap menjadi sasaran Fase 2.
- Modul 17 (dashboard) dan 20 (otomasi) tidak memuat fitur Fase 1-2. Keduanya bergantung pada filter tersimpan (13), notifikasi/job queue (19, 27), dan model event (28).
- Modul 29–35 hanya berada di Fase 5 dan 7, dan modul 36–37 seluruhnya di Fase 6. Semuanya dibangun di atas fondasi Fase 1-3 (work item, workflow, JQL, otomasi, notifikasi, API).

## Peta 7 fase delivery

Ringkasan berikut merangkum [00 › Peta fase delivery](00-visi-ruang-lingkup-glosarium-baseline.md#peta-fase-delivery). Kolom Fitur dan sebaran prioritas dihitung dari seluruh modul.

| Fase | Nama | Tujuan | Fitur | P0 / P1 / P2 / P3 | Modul utama |
| --- | --- | --- | ---: | --- | --- |
| 1 | MVP pelacak kerja tim inti | Tim kecil mengelola pekerjaan harian dari awal sampai akhir di atas fondasi yang aman: auth & sesi, proyek, work item, workflow default, board/backlog/sprint dasar, komentar & lampiran, pencarian dasar, notifikasi in-app, navigasi, serta remediasi baseline Sprint 0. | 254 | 249 / 4 / 1 / 0 | 00, 01, 03, 04, 06, 07, 08, 09, 10, 11, 12, 19, 21, 27, 28 |
| 2 | Konfigurabilitas & visibilitas tim | Tim dapat menyesuaikan proses dan melihat kemajuan: JQL, navigator & filter, editor workflow, custom field & layar, board lanjutan, versi & komponen, time tracking, laporan agile, email, realtime (SSE), serta layanan feature flag. | 352 | 55 / 283 / 14 / 0 | 04, 05, 07, 09, 10, 13, 14, 15, 19, 21 |
| 3 | Skala organisasi, otomasi & integrasi | Organisasi multi-tim dapat mengatur akses dan mengotomasi pekerjaan: permission scheme penuh, manajemen pengguna lanjut, dashboard, otomasi, tampilan proyek, REST API publik & webhook, impor/ekspor, proyek bisnis & formulir, serta administrasi sistem. | 402 | 52 / 136 / 196 / 18 | 02, 03, 16, 17, 20, 22, 23, 25, 26 |
| 4 | Perencanaan lanjut, DevOps & ekosistem | Perencanaan lintas tim dan keterhubungan dengan alat pengembangan serta identitas perusahaan: Plans, Teams, Goals, hierarki di atas Epic, GitHub/GitLab, SSO/SCIM, OAuth 2.0, Slack/Teams, ekstensi, dan AI opsional. | 310 | 0 / 18 / 142 / 150 | 02, 18, 22, 23, 24, 25 |
| 5 | Service Management inti | Tim layanan menerima dan menyelesaikan request customer sesuai target layanan: service project, request type, forms, help center & portal, customers & organizations, kanal email & widget, queues, SLA, knowledge base, approvals, canned responses, CSAT, dan laporan. | 193 | 74 / 87 / 28 / 4 | 29, 30, 31, 35 |
| 6 | Product Discovery | Tim produk mengumpulkan, memprioritaskan, dan menghubungkan ide ke delivery: ideas, fields & formula, views, sharing, insights, delivery tracking, dan roles. | 104 | 26 / 40 / 27 / 11 | 36, 37 |
| 7 | ITSM lanjut & Operations | Organisasi TI menjalankan praktik ITSM penuh dan respons operasi: incident/problem/change, services & deployment gating, Assets/CMDB & AQL, alerts/on-call/escalation, virtual service agent & AI, Customer Service Management, dan chat. | 188 | 43 / 55 / 68 / 22 | 32, 33, 34, 35 |
| — | Rujukan (tidak dibangun) | Fitur Data Center-only atau deprecated, dicatat agar inventaris tercakup. | 58 | 0 / 0 / 1 / 57 | tersebar |

Aturan ibu jari penempatan fase:

| Kondisi fitur | Fase |
| --- | --- |
| P0 modul inti | Fase 1-2 |
| P1 | Fase 2-3 |
| P2, `premium`, atau `enterprise` | Fase 3-4 |
| `jsm` | Fase 5 (inti) atau Fase 7 (ITSM/Assets/Ops/AI) |
| `jpd` | Fase 6 |
| `dc-only-or-deprecated` | Fase — |

Penyimpangan dari aturan ini dijelaskan di bagian "Ringkasan prioritas & fase" modul terkait. Contohnya: P0 di Fase 3 untuk dashboard dan otomasi (modul 17, 20), dan sebagian P0 Service Management di Fase 7 karena ITSM/Assets/Ops memang domain Fase 7.

**Sebelum Fase 1 (Sprint 0).** Baseline repo memiliki utang teknis yang harus diselesaikan lebih dulu, misalnya migrasi `001` yang tidak valid, initdb compose yang menjalankan `down.sql`, path Makefile, `JWT_SECRET` default, CORS `*`, `GET /users` tanpa cek admin, `api.ts` usang, dan halaman frontend yang terhapus. Daftar remediasi lengkapnya ada di [00 › Baseline repo & gap](00-visi-ruang-lingkup-glosarium-baseline.md#baseline-repo--gap).

## Pertanyaan terbuka gabungan

Bagian ini menggabungkan 20 pertanyaan dari outline riset dengan 357 butir dari bagian "Pertanyaan terbuka" ke-38 modul. Pertanyaan yang sama di beberapa modul dilebur menjadi satu, lalu semuanya dikelompokkan per tema.

Cara membaca sumber: label di akhir setiap butir menunjuk asal pertanyaan. **O** berarti outline riset. **Mxx** berarti modul `xx` (lihat [Daftar modul](#daftar-modul)), dan ID fitur disebut bila ada. Rumusan lengkap dan konteksnya tetap ada di modul sumber.

Pertanyaan outline "apakah Service Management dan Product Discovery di luar cakupan" sudah **terjawab**: keduanya masuk cakupan sebagai Fase 5-7 (modul 29–37).

### A. Identitas produk, merek & terminologi

1. **Nama final produk** dan perlu tidaknya pendaftaran merek (mis. di DJKI) sebelum rilis publik. Nama ini juga dipakai untuk prefix subject email, header From, aktor otomasi, perintah slash, dan konektor chat. (O; M00; M19 NTF-018, NTF-024, NTF-037; M25)
2. **Lisensi open-source** proyek (MIT, Apache-2.0, AGPL-3.0). Pilihan ini menentukan daftar tolak lisensi dependensi. (M00 VISI-001)
3. **Nama bahasa kueri di UI**, dan apakah akronim "JQL" boleh tampil di UI. (M00)
4. **Label UI**: "Issue" atau "Work item", "Project" atau "Space" (di UI Indonesia "Proyek" atau "Ruang"). Juga set terminologi default site ("baru" atau "klasik"), apakah dapat dimatikan per site, dan apakah berlaku untuk proyek discovery. (O; M00 VISI-002; M04 PRJ-001; M26 BIZ-001; M36)
5. **Alias JQL berbasis istilah work item** (mis. `workitemkey`, `workItemType`) di samping nama klasik. (M12 JQL-001)
6. **Nilai bermerek vendor di kontrak API dan data**: `accountType: atlassian`, header `X-Atlassian-*`, scope `*:jira-*`, nama grup sistem `jira-users-<site>`, permission key `BROWSE_PROJECTS`, role `atlassian-addons-project-access`. Pilihannya dipertahankan demi kompatibilitas, atau diganti nama netral dengan nama lama sebagai alias. (M01 AUTH-001; M02; M03; M23)
7. **Nama peran discovery** Creator/Contributor: diberi nama ulang di glosarium atau dipakai apa adanya. (M37)
8. **Bahasa**: Bahasa Indonesia atau Inggris sebagai default UI, dokumentasi pengguna dwibahasa, bahasa UI lain yang menjadi target, dan kapan terjemahan data konfigurasi (nama status, work type, field, priority, resolution) tersedia. (O; M05 CFG-061; M21 NAV-033; M27 NFR-021)
9. **Telemetri penggunaan anonim**: dikirim atau tidak, dan bila ya apakah opt-in. Juga tujuan feedback produk dan permintaan akses/upgrade di instalasi self-hosted. (M00; M21 NAV-007, NAV-022, NAV-024)

### B. Tenancy, edisi, tier & batas plan

1. **Model tenancy**: satu organisasi per instance (self-hosted) atau SaaS multi-tenant (`org_id` di semua tabel atau schema per tenant). Keputusan ini menentukan cakupan multi-site/multi-workspace dalam satu instalasi (account switcher, "Select site", org switcher, pemindahan site, discovered products, data residency) dan fitur mode hosted (billing, trial, deaktivasi Free, rename subdomain). **Keputusan paling mendesak karena memengaruhi model data sejak MVP.** (O; M01 AUTH-001, AUTH-003; M02 USR-001, USR-014, USR-042; M21 NAV-007, NAV-008; M22 ADM-048, ADM-053–056; M27 NFR-014)
2. **Tier aktif default di self-hosted**: semua fitur terbuka (setara Enterprise) atau simulasi plan bertingkat, dengan gating hanya untuk layanan terkelola. (O; M00 VISI-003; M03; M21 NAV-024)
3. **Penerapan batas plan** (Free 10 pengguna/2 GB), kuota guest gratis per seat berbayar, dan kuota service account non-billable. (M00; M02)
4. **Fitur yang berbayar di produk pembanding, dibuka atau digate:**
   - paket keamanan lanjutan (SSO, SCIM, data security policy, audit lanjutan) serta kebijakan autentikasi, enforcement 2SV, dan kontrol API token;
   - issue security, custom role, dan archive;
   - Target start/end dan hierarki di atas Epic, serta Plans, Teams, dan Goals;
   - multiple help center, custom pages, dan custom domain;
   - paket ITSM Premium/Enterprise, Assets (termasuk kemungkinan versi dasar di Standard), Surveys dan Advanced Analytics, serta Multi-space work;
   - plan Product Discovery (idea types, Tree view, roadmap lintas project, public link, sandbox, kuota Creator Free);
   - AI release notes.

   (M01 AUTH-010–012; M02; M03; M05 CFG-007, CFG-022; M15 REL-020; M18 PLAN-027, PLAN-057, PLAN-061; M30 JSM-PRT-003, JSM-PRT-008, JSM-PRT-009; M31; M32; M33; M36; M37)
5. **Kuota & metering**: kursi agent, kuota assisted conversations, dan kuota rule run otomasi. Pilihannya hanya metering, pemblokiran, atau hanya berlaku di edisi hosted. (M20 AUTO-015; M28; M35 JSM-AI-003, JSM-AI-058)
6. **Sandbox**: dibangun sebagai fitur produk (salinan site) atau cukup panduan menjalankan instance kedua. Juga apakah promote konfigurasi sandbox ke produksi masuk Fase 4. (M00; M22 ADM-049)
7. **Kebutuhan enterprise lebih awal**: apakah ada calon pelanggan yang memerlukan SSO SAML/OIDC, SCIM, IP allowlist, atau data residency sebelum Fase 4. (O)
8. **Kemasan Customer Service Management**: app berlisensi terpisah atau template di Service Management, dan apakah team-managed didukung. (M35 JSM-AI-059)

### C. Autentikasi, izin & keamanan

1. **Model sesi**:
   - Model A: access token ±15 menit + refresh token dengan rotasi.
   - Model B: JWT 12 jam sliding dengan `jti` yang dicek ke tabel sesi.

   Juga cara penyimpanan di SPA (cookie httpOnly dengan CSRF, atau token di memori dengan refresh token) dan lama masa transisi sebelum token tanpa `jti`/`kid` ditolak. (M01 AUTH-007, AUTH-008; M27 NFR-015, NFR-020)
2. **Nasib `username`** dari skema repo: dihapus, dipertahankan sebagai handle tampilan non-unik, atau tetap menjadi alias login legacy. (M01; M28)
3. **Penyedia keamanan yang ramah self-host**: cek password bocor (layanan online k-anonymity atau dataset lokal), penyedia CAPTCHA, penyedia login sosial prioritas, dan fase aktivasi passkey serta login kode email. (M01 AUTH-003, AUTH-005, AUTH-006)
4. **Protokol tambahan**: OIDC generik selain SAML dan Google Workspace, dan Personal Access Token bergaya Bearer di samping Basic auth dengan API token. (M02; M23)
5. **Step-up auth admin**: apakah aktif secara default pada instalasi tanpa SSO. (M22 ADM-001)
6. **Migrasi kolom `Role` entity User** (default `user`) ke group dan global permission. (M03)
7. **Perilaku mesin izin**:
   - target propagasi perubahan izin: real time (≤ 5 detik) atau eventual (≤ 1 menit);
   - proteksi lockout saat grant Administer terakhir dihapus: ditolak keras atau konfirmasi berlapis;
   - role default permission scheme: Member/Viewer saja, atau juga Developers/Users untuk impor;
   - siapa yang boleh menghapus group filter subscription;
   - izin yang mensyaratkan pemadaman email pada bulk operation;
   - role team-managed yang boleh menyetel issue restriction, dan apakah restriction diwariskan ke child.

   (M03)
8. **Kontrol data enterprise**: apakah Guest Access, Data Classification, Data Security Policies, dan Data Redaction masuk roadmap. Untuk Data Redaction juga diputuskan apakah menjadi add-on atau bagian Enterprise, dan apakah deteksi data sensitif otomatis dibangun sendiri. (M03; M06 ISS-043)
9. **Project Private**: apakah Admin Situs dapat membaca isinya tanpa menjadi anggota, apakah status Private discovery berlaku mutlak, dan apakah Private tersedia di plan Free. (M04 PRJ-039; M37)
10. **Akses publik/anonim**:
    - public sharing dan akses anonim dashboard;
    - ekspor CSV oleh pengguna anonim di project publik;
    - form bisnis publik tanpa login;
    - public link view JPD dan siapa yang boleh mengaktifkannya;
    - apakah request anonim wajib mencantumkan email;
    - Single Email Address sebagai penerima notifikasi.

    (M13; M17; M19 NTF-006; M26 BIZ-036; M30 JSM-PRT-017; M36; M37 JPD-INS-029)
11. **Konten & egress berisiko**:
    - upload ikon SVG (risiko XSS);
    - antivirus (ClamAV) aktif secara default;
    - Send web request ke alamat jaringan privat (SSRF; diblokir dengan allowlist admin);
    - approve lewat balasan email (risiko spoofing) dan verifikasi yang diwajibkan;
    - embed iframe smart link aktif secara default.

    (M05 CFG-004; M07 EDT-021; M08 CMT-017, CMT-019; M20; M32 JSM-ITSM-013)
12. **Profil kepatuhan**: HIPAA-ready dan BYOK dibangun, atau cukup didokumentasikan sebagai panduan operator. Juga kebijakan retensi untuk UU PDP Indonesia dan GDPR. (O; M27 NFR-016, NFR-019)

### D. Arsitektur, model data & infrastruktur

1. **Tipe identitas**: `id` project bigint numerik (kompatibel API) sementara `users.id` tetap UUID `VARCHAR(36)`, atau diseragamkan. Juga apakah `account_id` terpisah atau memakai UUID yang ada. (M04 PRJ-001; M28)
2. **Penyimpanan custom field**: hybrid JSONB + tabel bertipe sebagai keputusan final, atau EAV murni di Fase 1. (M28)
3. **Format rich text**: dokumen JSON ala ADF (ProseMirror/TipTap) atau Markdown. Juga apakah impor/ekspor ADF dan wiki markup diperlukan, termasuk alias `/rest/api/2`. (O; M23)
4. **Search engine**: Postgres FTS saja atau OpenSearch/Meilisearch, dan pada ambang berapa mode eksternal wajib. Juga apakah stemming Bahasa Indonesia menjadi syarat rilis Fase 1, dan apakah konfigurasinya per situs atau per dokumen. (O; M12 JQL-012; M13 SRC-001; M27 NFR-010)
5. **Queue, cache & pub/sub**: River (PostgreSQL) dengan Redis opsional, atau Redis wajib sejak Fase 2. Juga apakah Redis di docker-compose tetap dipakai, dan apakah outbox cukup dengan Postgres (polling/LISTEN-NOTIFY) atau perlu broker sejak Fase 3. (O; M27 NFR-003, NFR-008, NFR-009; M28)

   **Terjawab sebagian ([ADR-0004](../adr/0004-job-latar-dan-queue.md)):** queue memakai Redis (`asynq`) dan Redis wajib sejak Fase 1, bukan River; Redis tetap ada di docker-compose; outbox tetap tabel PostgreSQL yang diteruskan ke Redis oleh relay (at-least-once). Yang masih terbuka: apakah broker terpisah dibutuhkan sejak Fase 3.
6. **Konkurensi edit**: optimistic locking dengan HTTP 409 atau last-write-wins per field, untuk field umum dan Description. Juga apakah co-editing realtime dibutuhkan. (O; M07 EDT-026; M27 NFR-006)
7. **Riwayat & graf**: apakah riwayat versi workflow dengan rollback dibangun di Fase 2, dan apakah graf entitas (DM-015) dibangun di Fase 4 tanpa konsumen AI yang pasti. (M28)
8. **Logger**: migrasi dari logrus ke `log/slog`, atau logrus dipertahankan dengan formatter JSON. (M27 NFR-012)
9. **Lampiran**:
   - disk lokal, S3-kompatibel, atau keduanya;
   - batas ukuran per file dan kuota default (usulan 100 MB per file, tanpa kuota di self-hosted);
   - apakah konversi dokumen Office dan transcoder video menjadi komponen bawaan atau dependency opsional.

   (O; M08 CMT-015, CMT-016, CMT-019)

   **Terjawab ([ADR-0002](../adr/0002-penyimpanan-lampiran.md)):** penyimpanan S3-kompatibel saja (tanpa disk lokal; MinIO untuk dev dan self-hosted); 100 MB per file (konfigurasi hingga 2 GB), 2.000 attachment per work item, tanpa kuota default; konversi Office dan transcoder video bukan komponen bawaan, melainkan dependency opsional di luar Fase 1.
10. **Email keluar & masuk**:
    - penyedia email keluar (SMTP sendiri, SES, Mailgun);
    - email masuk lewat domain inbound bawaan, penerima SMTP/MX di dalam binary, atau hanya IMAP/POP dan forwarding;
    - domain alamat email bawaan service project, dan apakah inbound mail wajib dikonfigurasi sebelum service project dibuat;
    - apakah balasan email notifikasi menjadi comment di project non-service, dan pada fase berapa.

    (O; M08 CMT-001; M19 NTF-023, NTF-034; M26 BIZ-047; M29 JSM-PRJ-063; M30 JSM-PRT-040)

    **Terjawab sebagian ([ADR-0003](../adr/0003-transport-email-keluar.md)):** penyedia email keluar adalah SMTP generik di balik interface `Mailer` (SES dan Mailgun dapat dipakai lewat relay SMTP), dikirim lewat queue. Tiga butir lain (email masuk, domain bawaan service project, balasan email menjadi comment) tetap terbuka.
11. **Custom domain & TLS**: dikelola produk (ACME) atau diserahkan ke reverse proxy. (M30 JSM-PRT-009)
12. **Akun customer**: disimpan di tabel `users` dengan kolom tipe akun atau di tabel terpisah. Juga apakah satu email boleh memiliki akun portal-only sekaligus akun situs, dan bagaimana deduplikasi notifikasi bagi pengguna berlisensi yang juga reporter/participant. (M29 JSM-PRJ-045, JSM-PRJ-048, JSM-PRJ-068; M30 JSM-PRT-014, JSM-PRT-036)
13. **Infrastruktur realtime live chat CSM**: kanal realtime umum (modul 27) atau infrastruktur sendiri. (M35)

### E. Kompatibilitas API, URL & migrasi

1. **Lapisan kompatibel bergaya produk pembanding** di samping `/api/v1`. Pilihannya dibangun penuh, subset populer (issue, search, comment, transition, project, user, board, sprint), atau ditunda, dan apakah sejak Fase 1 atau Fase 3. Cakupannya meliputi:
   - `/rest/api/3` dan `/rest/agile/1.0`, termasuk endpoint worklog;
   - path DevOps (`/rest/devinfo/0.10/bulk` dan sejenisnya);
   - `/rest/servicedeskapi`;
   - URL Assets workspace;
   - header Alert API.

   (M02; M04 PRJ-064; M14 TIME-040; M23; M24; M33; M34 JSM-OPS-007; M35 JSM-AI-024)
2. **Format galat**: bentuk native `{error, status, message, data, errors}` dipertahankan berbeda dari `{errorMessages, errors}` di lapisan kompatibel, atau diseragamkan. (M23)
3. **Skema URL**: `/browse`, `/jira/software/c/projects/...`, dan redirect `/secure/*.jspa` dipertahankan persis, atau disederhanakan dengan redirect. (M21)
4. **Bahasa ekspresi** untuk condition/validator workflow, ekspresi API, dan titik ekstensi app: kompatibel dengan Jira expressions, atau CEL/expr dengan sintaks berbeda. (M09 WF-027, WF-028, WF-042; M23 API-031)
5. **Kedalaman kompatibilitas JQL**:
   - sintaks persis atau subset terdokumentasi;
   - kebijakan menambah reserved word;
   - nama field terlokalisasi sebagai alias;
   - pencocokan `labels` case-sensitive atau tidak;
   - apakah `childIssuesOf()` tersedia di semua plan.

   (O; M12 JQL-004, JQL-005, JQL-022, JQL-047)
6. **Format rule otomasi**: definisi JSON dan smart values kompatibel penuh dengan ekspor rule produk pembanding, atau subset terdokumentasi. (M20)
7. **Prioritas migrasi**:
   - importer "dari Jira" (REST API sumber, termasuk history dan lampiran) dimajukan ke Fase 2/3 sebagai jalur adopsi utama;
   - exporter JSON per proyek yang round-trip;
   - daftar importer alat lain yang wajib saat rilis.

   (O; M22 ADM-038; M25 INT-016, INT-026)
8. **Service project team-managed**: dihapus sepenuhnya termasuk untuk impor, atau didukung baca-saja untuk project hasil migrasi. (M29 JSM-PRJ-012)

### F. Batas default, guardrail & retensi

Semua angka di bawah adalah **usulan** di modul sumber dan perlu ditetapkan final, termasuk apakah nilainya berbeda untuk self-hosted dan hosted, dapat dikonfigurasi Admin Situs, atau bergantung tier. Pertanyaan payungnya adalah apakah angka disalin dari produk pembanding atau ditetapkan lewat uji beban (M22 ADM-041).

1. **Proyek & konfigurasi**:
   - panjang nama, ukuran avatar, panjang description, jumlah field dan work type team-managed, ukuran project property (M04 PRJ-011, PRJ-016, PRJ-018, PRJ-051, PRJ-062);
   - work type per project team-managed, opsi per context, context per field, field per project/work type, custom field per situs, level hierarki kustom (M05; M28);
   - panjang nama status (60), status per workflow team-managed (30), status/transition per workflow company-managed (100/500), workflow per situs (2.000), work item per bulk transition (1.000) (M09 WF-001, WF-013, WF-017, WF-025).
2. **Work item**:
   - guardrail per work item dan apakah dapat diubah (M06 ISS-044);
   - batas attachment, comment, worklog, link, dan sub-task, serta `maxResults` endpoint ID-only (M27 NFR-004);
   - worklog 5.000 atau 10.000 per work item (M14 TIME-021; lihat juga butir keselarasan 8);
   - maksimal 3 pinned comment beserta pemegang izinnya (M08 CMT-007);
   - nama sprint 30 karakter dan batas panjang sprint goal (M11 SPR-015, SPR-016).
3. **Board, laporan, dashboard & rencana**:
   - batas kartu per board atau per kolom (M10 BRD-051; M28);
   - 5.000 work item per laporan dan perhitungan asinkron (M16 RPT-004);
   - ambang 20% item tanpa estimasi untuk banner (M16 RPT-003);
   - 50 gadget per dashboard dengan timeout 30 detik, dan ukuran halaman direktori 20 atau 25 baris (M17);
   - 5.000 item Timeline/plan, 100 sumber, 50 scenario, 10 view, 20 rule format (M18 PLAN-014, PLAN-055).
4. **Kueri & pencarian**:
   - panjang kueri, jumlah clause, item `IN`, timeout, `maxResults` (M12 JQL-009, JQL-052);
   - timeout dan panjang JQL/URL navigator (M13 SRC-006);
   - work item per email subscription dan lampiran CSV untuk hasil besar (M13 SRC-047).
5. **Notifikasi, otomasi & API**:
   - jendela batching 10 menit dan apakah user boleh menonaktifkannya (M19 NTF-019);
   - recipient per event, email per jam per penerima, batas email otomasi (M19 NTF-009, NTF-037);
   - rule run per plan (M20; M28);
   - burst, quota, dan biaya per endpoint rate limit (M23).
6. **Navigasi**: retensi For you 90 hari, recent 20/200 per tipe, starred 100 per tipe, retensi notifikasi 30 hari, rate limit share 50/jam. (M21)
7. **Service Management, Assets & Discovery**:
   - request type, portal group, participant, alamat email, dan form per work item (M29);
   - alamat email per project, help center per situs, custom pages, throttling email, dan retensi email log (M30);
   - jumlah project di Multi-space work (M31);
   - jumlah status page (M32);
   - object per workspace, attribute per object type, hasil AQL, ukuran file import, object per field multi, dan bulk (M33);
   - field, view, idea type, dan level hierarki discovery (M36);
   - Category, work type, field, dan form per project bisnis (M26).
8. **Retensi data**:
   - masa simpan tempat sampah sebelum purge;
   - audit log (usulan 90 hari untuk otomasi);
   - log aplikasi, log pengiriman webhook, task async, notifikasi, dan log eksekusi otomasi;
   - alert dan integration logs per plan, serta email log;
   - apakah kebijakan retensi ADM-047 (belum terkonfirmasi ada di produk pembanding) tetap dibangun sebagai nilai tambah kepatuhan.

   (O; M20; M22 ADM-047; M23; M27 NFR-017, NFR-018; M28; M30; M34 JSM-OPS-004, JSM-OPS-031)

### G. Proyek, konfigurasi work item & workflow

1. **Gaya proyek**: mendukung team-managed dan company-managed, atau satu model terpadu berbasis skema dengan preset. Bila keduanya, apakah Fase 1 membangun keduanya atau team-managed lebih dulu (company-managed memakai skema default sampai Fase 2). (O; M04 PRJ-003)
2. **Aturan proyek**:
   - apakah project type company-managed dapat diubah (PRJ-023);
   - apakah key boleh memuat underscore (PRJ-012);
   - apakah key yang dihapus permanen dapat dipakai lagi (PRJ-033);
   - siapa yang boleh mengarsipkan dan me-restore dari trash (PRJ-032, PRJ-033);
   - apakah perubahan Key/Category hanya untuk Admin Situs (PRJ-016);
   - apakah migrasi team-managed ke company-managed memetakan Story point estimate ke Story Points (PRJ-024).

   (M04)
3. **Estimasi**: satu field estimasi terpadu atau dua field berbeda ID demi impor, dan apakah keduanya dikonsolidasikan di UI sejak awal. Juga estimasi default (story point atau waktu, dipilih per board?) dan perlu tidaknya hierarki di atas Epic sebelum Fase 4. (O; M05 CFG-016; M14 TIME-008)
4. **Model field company-managed**: Field schemes yang disederhanakan saja, atau juga field configuration dan field configuration scheme klasik. (M05 CFG-043–045)
5. **Detail konfigurasi**:
   - apakah Category tersedia di semua tipe project (CFG-021; BIZ-021);
   - apakah subtask boleh langsung berada di bawah Epic (CFG-006);
   - apakah default Description per work type di company-managed dibangun native (CFG-052);
   - apakah format Percentage menyimpan 0.25 atau 25 (CFG-026).

   (M05; M26)
6. **Workflow**:
   - model draft workflow legacy (draft terpisah dan backup saat publish) atau editor sesi dengan optimistic locking (WF-022);
   - apakah transisi ke kategori Done tanpa resolution dicegah secara default (WF-009);
   - apakah simplified workflow board-managed dibangun (WF-026);
   - apakah saat migrasi status seluruh project dikunci atau hanya work item terdampak (WF-036).

   (M09)
7. **Status default template Kanban company-managed**: To Do / In Progress / Done atau Backlog / Selected for Development / In Progress / Done. (M10 BRD-001)

### H. Work item, editor, komentar & waktu

1. **Operasi work item**:
   - memilih status awal di Create dialog (ISS-002);
   - Parent lintas project di team-managed (ISS-025);
   - Bulk Edit "Find and remove" label (ISS-035);
   - memilih project tujuan langsung di Clone (ISS-031);
   - apakah Issue Collectors yang legacy tetap dibangun atau langsung diganti Forms publik dan portal (ISS-040).

   (M06)
2. **Editor**:
   - retensi draft lokal dan sinkronisasi server (EDT-001);
   - UI untuk mark atau atribut yang belum pasti ada: highlight, alignment, custom panel, caption/border gambar, decision item (EDT-004, EDT-005, EDT-011, EDT-016, EDT-019);
   - layout multi-kolom (EDT-013);
   - visibilitas media di comment ber-visibilitas terbatas (EDT-019);
   - konversi otomatis issue key menjadi inline card (EDT-020);
   - resolver smart link native (EDT-021);
   - fase custom emoji dan siapa yang boleh mengunggah (EDT-015).

   (M07)
3. **Komentar**:
   - nasib reply saat comment induk dihapus (usulan: placeholder) (CMT-003);
   - `parentId` di REST API publik (CMT-005);
   - apakah Image Annotation dibangun (CMT-020).

   (M08)
4. **Strategi dokumen/wiki**: modul Pages/dokumen internal, konektor wiki eksternal, atau sekadar daftar tautan. Keputusan ini sekaligus menentukan integrasi wiki (CMT-030, CMT-031), Linked pages sprint (SPR-026), tab Pages bisnis (BIZ-026), INT-042–046, embed tabel Assets, dan knowledge base native atau konektor wiki eksternal. (M08; M11; M25; M26; M31; M33)
5. **Waktu**:
   - apakah Total Time Spent mencakup worklog terbatas bagi pengguna tanpa akses (TIME-019, TIME-037);
   - izin mencatat worklog atas nama orang lain (TIME-021);
   - laporan timesheet native atau diserahkan ke ekstensi.

   (M14)

### I. Board, backlog, sprint, rilis, laporan & perencanaan

1. **Board**:
   - nasib board saat project location dihapus atau diarsipkan (BRD-004);
   - sprint per board atau per project pada multiple board team-managed (BRD-008, BRD-049);
   - board lintas project team-managed (BRD-003, BRD-048);
   - trash board dengan masa pemulihan (BRD-005);
   - Days in column dengan hari kalender atau hari kerja, beserta ambang warnanya (BRD-039);
   - apakah jumlah di header kolom dan WIP mengikuti filter aktif (BRD-028, BRD-034);
   - timer stand-up dan urutan giliran (BRD-036);
   - inline create di sprint aktif tanpa izin Manage Sprints (BRD-019).

   (M10)
2. **Sprint**:
   - apakah Split Issue, Parallel Sprints, Sprint Reopen, dan Versions panel tersedia di team-managed (SPR-009, SPR-010, SPR-021, SPR-023);
   - field yang diwarisi hasil split (SPR-010);
   - apakah sprint aktif boleh dihapus (SPR-022);
   - apakah sprint kosong boleh di-start (SPR-018);
   - penyimpanan preferensi View settings di server (SPR-002);
   - syarat Manage Sprints pada board multi-project (SPR-013).

   (M11)
3. **Rilis & komponen**:
   - definisi "unresolved" (Resolution kosong atau kategori Done) (REL-005);
   - permission terpisah Manage Versions dan izin serupa untuk component (REL-008);
   - Releases di proyek bisnis (REL-001);
   - keunikan nama case-insensitive (REL-001, REL-028);
   - Release from Kanban multi-project (REL-022);
   - apakah approver release dapat memblokir (REL-016);
   - fitur rilis di team-managed (REL-021, REL-022, REL-025);
   - default assignee bila ada beberapa component (REL-031).

   (M15)
4. **Laporan**:
   - Version report di Kanban (RPT-001, RPT-015);
   - ekspor CSV/PNG native (RPT-036);
   - working days per board team-managed (RPT-005);
   - apakah Kanban sub-filter membatasi populasi CFD dan Control chart (RPT-007, RPT-017);
   - kriteria "Issues for attention" dapat dikonfigurasi (RPT-033);
   - Analytics/Data Lake native atau koneksi read-only ke BI eksternal (RPT-038–041; ADM-042).

   (M16; M22)
5. **Dashboard**:
   - apakah dashboard dimajukan ke Fase 2 sebagai beranda tim;
   - keunikan nama per owner;
   - minimize gadget per pengguna atau per dashboard;
   - populasi Average Number of Times in Status;
   - status archived di REST;
   - mode kiosk/full-screen pengganti Wallboard;
   - iCal dan view mingguan di Work Items Calendar;
   - gadget yang wajib di tampilan mobile.

   (M17)
6. **Timeline & Plans**:
   - view bersama atau view personal bernama (PLAN-010, PLAN-017, PLAN-026);
   - work item tanpa Epic di Timeline (PLAN-002);
   - kunci urutan tab view oleh admin (PLAN-001);
   - penyelesaian konflik saat save plan (PLAN-045);
   - team hierarchy sejak awal (PLAN-057);
   - status progres goal otomatis (PLAN-062);
   - apakah auto-scheduler P3 tetap dibangun (PLAN-052).

   (M18)

### J. Pencarian, filter & navigasi

1. **Halaman hasil pencarian penuh** (SRC-012) atau Enter di quick search selalu membuka advanced search. (M13)
2. **Hak atas filter**: apakah Editors filter boleh mengubah daftar Viewers/Editors (SRC-046), dan bagaimana eksekusi filter milik pengguna nonaktif (izin owner terakhir, izin viewer, atau nonaktif sampai dialihkan) (SRC-049). (M13)
3. **Konektor spreadsheet** Google Sheets/Excel (SRC-055). (M13)
4. **Navigasi**: command palette global `Ctrl/Cmd+K` (NAV-042), apakah unwatch manual mencegah autowatch berikutnya, dan apakah edit field memicu autowatch (NAV-031). (M21)

### K. Notifikasi, otomasi & integrasi

1. **Notifikasi**:
   - apakah notifikasi in-app mengikuti notification scheme (NTF-016);
   - Digest email notifikasi belum dibaca, atau cukup Filter Subscriptions (NTF-017);
   - apakah Admin Proyek boleh mengubah project email address (NTF-025).

   (M19)
2. **Otomasi**:
   - version history rule dengan compare/restore di Fase 3 atau ditunda (AUTO-007);
   - action native "Delay/Wait" sebagai pembeda (AUTO-003);
   - action native Create/Start/Complete sprint (AUTO-050);
   - fungsi JQL `recurrence` untuk work item berulang, yang perlu ditambahkan ke modul 12 (AUTO-072).

   (M20)
3. **DevOps**:
   - konektor native GitHub/GitLab/Bitbucket, atau di rilis pertama hanya DevOps REST API dan webhook generik;
   - apakah Workflow Triggers (DEV-025) dibangun, atau cukup template otomasi DevOps;
   - kanal pelaporan kegagalan Smart commits;
   - masa simpan dev info setelah integrasi dicabut (dihapus segera atau masa tenggang 30 hari);
   - plugin Jenkins, orb CircleCI, task Azure Pipelines, dan ekstensi VS Code buatan sendiri;
   - fase panel Security, Feature flags, Designs, dan Operations;
   - ekstensi UI di halaman PR GitHub/Bitbucket.

   (M24)
4. **Chat & app**:
   - apakah Slack (P1) ditarik ke Fase 3 untuk Jira inti dan ke Fase 5 untuk Service Management;
   - app Slack/Teams discovery dibangun sendiri atau menunggu app chat umum.

   (M25 INT-030–034; M30 JSM-PRT-045, JSM-PRT-047; M37 JPD-INS-032, JPD-INS-033)
5. **Ekosistem ekstensi**:
   - perlu tidaknya ekosistem app, atau cukup REST + webhook + otomasi;
   - bentuk registry app (katalog lokal, katalog komunitas, atau keduanya) dan kurator badge keamanan (INT-061);
   - UI extension point berbasis iframe (INT-069–075);
   - model app: OAuth 2.0 3LO + webhook dinamis saja, atau juga runtime fungsi bergaya Forge;
   - apakah GraphQL API dan Remote MCP server masuk roadmap.

   (O; M23; M25)
6. **Integrasi discovery & ITSM eksternal**:
   - app Intercom (JPD-INS-016);
   - penghubung ke pelacak kerja self-hosted eksternal (JPD-INS-024);
   - REST API publik untuk insight dan view (JPD-INS-037);
   - integrasi status page eksternal dan katalog komponen developer (JSM-ITSM-010, JSM-ITSM-027).

   (M32; M37)

### L. Proyek bisnis & formulir

1. **Kapan builder form lanjutan dibangun**: conditional logic, section, dan template form untuk form bisnis di Fase 3, atau bersama engine forms Service Management (BIZ-030, BIZ-033). (M26)
2. **Approval**:
   - apakah langkah approval tersedia di team-managed dan project software sejak awal (BIZ-042);
   - aturan jumlah persetujuan default (satu approver atau semua approver);
   - apakah approver boleh diganti saat approval pending (BIZ-043).

   (M26)
3. **Katalog template bisnis** resmi Fase 3, dan apakah anggota project (bukan hanya admin) boleh membuat form (BIZ-034). (M26)

### M. Service Management

1. **Service project & forms**:
   - aksi Duplicate request type dan Copy form ke project lain;
   - transisi yang boleh dipicu customer, dimodelkan sebagai properti "customer can trigger";
   - lingkup dan metode autentikasi data connections Forms;
   - hak collaborator atas form internal;
   - jadwal template katalog (Design, Sales, demo space).

   (M29 JSM-PRJ-010, JSM-PRJ-011, JSM-PRJ-020, JSM-PRJ-031, JSM-PRJ-032, JSM-PRJ-034, JSM-PRJ-039, JSM-PRJ-042, JSM-PRJ-047)
2. **Portal & customer**:
   - peran khusus Help center admin (JSM-PRT-049);
   - impor customer via CSV atau hanya REST API (JSM-PRT-020, JSM-PRT-033);
   - apakah mode safe customer notifications dibangun di gelombang pertama Fase 5 (JSM-PRT-036).

   (M30)
3. **Queue, SLA, KB & CSAT**:
   - queue personal dan aksi Duplicate queue;
   - ambang warna "mendekati breach" SLA;
   - definisi resmi deflection;
   - apakah CSAT dapat diperbarui dan survei dikirim ulang saat reopen;
   - algoritma routing default Workforce Management, dan apakah jadwal shift berbagi model dengan on-call modul 34.

   (M31)
4. **ITSM**:
   - dukungan ITSM, risk assessment, dan Journeys di service space team-managed;
   - conference call bawaan atau integrasi Zoom/Teams/Meet (JSM-ITSM-024);
   - status page native lengkap (JSM-ITSM-026);
   - freeze window sebagai peringatan atau blokir keras (JSM-ITSM-041);
   - mekanisme risk assessment default: rule builder atau skor berbobot (JSM-ITSM-042);
   - penempatan Case Management dan Journeys (JSM-ITSM-050, JSM-ITSM-051);
   - waktu tunggu dan perilaku timeout deployment gating (JSM-ITSM-047);
   - sertifikasi PinkVERIFY (JSM-ITSM-052).

   (M32)
5. **Assets**:
   - field Assets di project non-service sejak rilis pertama;
   - role schema untuk project role, dan apakah Agent otomatis menjadi User;
   - data contoh di template schema;
   - fase Watch/Move object, inline edit, filter AQL tersimpan, dan comment object;
   - lingkup reference type;
   - scheduled import dari URL;
   - Assets Discovery dan Data Manager dibangun sendiri atau diserahkan ke ekosistem;
   - konektor pertama: Microsoft Entra ID atau Google Workspace.

   (M33)
6. **Operations**:
   - team ops memakai entitas Teams umum (PLAN-057) atau team khusus;
   - penyedia SMS/voice beserta negara dan kuotanya;
   - integrasi monitoring rilis pertama;
   - Policies dalam satu daftar atau dua, dan apakah global policy khusus Premium;
   - Operations tanpa service project;
   - edge connector dan provider Terraform dibangun tim inti atau komunitas.

   (M34 JSM-OPS-001, JSM-OPS-004, JSM-OPS-013, JSM-OPS-014, JSM-OPS-027, JSM-OPS-030, JSM-OPS-034, JSM-OPS-036, JSM-OPS-039, JSM-OPS-046)
7. **Agen virtual & CSM**: Virtual Service Agent dan agen layanan customer AI sebagai satu engine dengan dua mode atau dua produk (JSM-AI-047, JSM-AI-064), dan status rilis Request Resolver serta agen otonom: beta di Feature Lab atau langsung GA (JSM-AI-044, JSM-AI-046). (M35)

### N. Product Discovery

1. **Kolaborasi & penilaian**:
   - vote budget per pengguna pada field Votes;
   - perubahan konfigurasi view dengan autosave atau banner Save/Discard;
   - formula bersarang non-siklik dan aturan pembulatan.

   (M36)
2. **Cakupan operasi**: apakah Merge dan Move lintas project discovery masuk Fase 6, dan apakah field discovery (Impact, Effort, Goals, formula) dapat dicari penuh di JQL. (M36)
3. **Peran**: siapa yang boleh membuat field di space discovery (Administrator saja atau juga Creator). (M37)

### O. Kecerdasan buatan (AI)

1. **Kebijakan AI lintas produk**:
   - apakah AI masuk roadmap inti atau modul opsional;
   - penyedia model: API eksternal, model self-hosted/on-premise, atau BYO-key;
   - model kuota dan penagihan (per pengguna, per organisasi, atau kredit);
   - kebijakan pengiriman data work item/konten ke penyedia;
   - apakah AI dapat dimatikan total per site, sesuai profil regulated deployment VISI-003.

   (O; M06 ISS-045, ISS-046; M07 EDT-030; M20 AUTO-073; M21 NAV-051–053; M32; M35 JSM-AI-036, JSM-AI-052; M37 JPD-INS-041)
2. **Fitur AI spesifik yang menunggu kebijakan di atas**:
   - agent AI pada transition (WF-043);
   - bahasa alami ke JQL (JQL-059, SRC-019);
   - AI release notes (REL-020);
   - pembuatan request type/form dan ringkasan/sentiment (JSM-PRJ-021, JSM-PRJ-035, JSM-PRJ-052);
   - ICC, AI risk assessment, dan draft PIR (M32);
   - AIOps grouping yang dimulai dari heuristik non-AI (JSM-OPS-042);
   - NL ke AQL dan ringkasan CI (M33);
   - agen AI coding sendiri atau integrasi agen pihak ketiga (DEV-032).

### P. Klien mobile & aksesibilitas perangkat

1. **Aplikasi mobile native** atau web responsif/PWA. Keputusan ini menentukan Mobile Push (diganti Web Push?), kontrol akses klien mobile (USR-039), pencatatan waktu mobile (TIME-049), gadget mobile, dan Critical alerts iOS untuk on-call, yang tidak tersedia lewat PWA. (O; M02 USR-039; M14 TIME-049; M17; M19 NTF-038; M21 NAV-050; M34 JSM-OPS-029)

### Q. Baseline repo

1. **Working tree saat PRD ditulis**: apakah perubahan uncommitted di `migrations/001_create_users_table.up.sql` serta penghapusan route `(guest)` dan `+error.svelte` disengaja, sehingga remediasi S0-01 dan S0-09 dapat dikerjakan di atasnya. (M00)

### Keselarasan lintas modul yang perlu diputuskan

Butir berikut berasal dari laporan review per modul. Isinya bukan pertanyaan produk murni, melainkan **ketidakselarasan antardokumen** yang harus dibereskan pemilik modul terkait:

1. **Guest dan anonymous access** diatur di dua tempat: modul 02 (USR-019, USR-024, level org/app) dan modul 03 (PERM-046, PERM-047, level scheme). Aturan di 03 perlu diperiksa agar tidak bertentangan dengan 02.
2. **Asumsi Fase 1 modul 02**: semua akun lokal diperlakukan sebagai akun terkelola satu organisasi, dan asumsi ini menjadi dasar Deactivate/Delete sebelum ada verifikasi domain. Asumsi ini perlu dikonfirmasi bersama keputusan tenancy (B.1).
3. **Resolutions administration**: WF-010 (modul 09) tumpang tindih dengan CFG-055 (modul 05). Perlu ditetapkan modul pemilik entitas dan endpoint-nya.
4. **Kolom `sequence` sprint**: tabel sprint di DM-010 (modul 28) belum memiliki kolom `sequence` untuk urutan future sprint, padahal SPR-012 butir 9 dan SPR-017 membutuhkannya.
5. **Parallel sprints di team-managed**: ADM-026 butir 3 (modul 22) menyatakan didukung, sedangkan SPR-023 (modul 11) menandainya ⚠ dan masih terbuka.
6. **SPR-016 Sprint Goal** berstatus P1/Fase 2, padahal field goal sudah ada di dialog Edit/Start sprint Fase 1 (SPR-015, SPR-018). Keduanya tidak bertentangan, tetapi dapat disederhanakan.
7. **"Copy comment to work description"**: CFG-058 butir 7 (modul 05) menyatakan opsi ini tidak dibangun dan khusus Data Center, sedangkan TIME-002 butir 6 dan TIME-014 butir 4 (modul 14) menganggapnya fitur Cloud.
8. **Batas worklog**: TIME-021 butir 1 memakai 5.000 agar sama dengan CMT-024 butir 6. Riset menyebut batas pembanding 10.000 worklog dan 5.000 comment, jadi kedua modul perlu diselaraskan bersamaan.
9. **ADM-057** berstatus Fase — (dc-only), padahal butir 16-21 (System info, health check, retry job, sesi aktif, logging/pprof) berguna bagi self-hosted. Perlu diputuskan apakah dinaikkan menjadi fitur Fase 3 atau dipindahkan ke modul 27.
10. **Tautan modul 22**: rujukan ke modul 25, 26, 27, 28, dan 35 masih berupa teks biasa dan perlu diubah menjadi tautan relatif.
11. **JSM-AI-034 Opsgenie migration** ber-tier `jsm` tetapi ditulis Fase — agar konsisten dengan JSM-OPS-047. Tier dan fase di kedua modul perlu diseragamkan.
12. **JSM-AI-059 butir 10, JSM-AI-060, dan JSM-AI-061** (CSM Products & Entitlements) tetap bertanda ⚠ walaupun inventaris tidak menandainya, karena detail CSM masih baru.
13. **Impor CSV customer**: JSM-PRT-033 butir 13 dan JSM-PRT-020 butir 13 (modul 30) masih berupa kandidat yang saling bertentangan sampai butir M.2 diputuskan.
14. **Heading gabungan modul 10**: 11 heading gabungan memuat 567 butir dari 581 sub-fitur inventaris. Selisihnya berasal dari sub-fitur yang identik antaralias dan sudah diperiksa manual, jadi tidak perlu tindakan kecuali pemilik produk ingin butir dipecah.

## Risiko

Ringkasan risiko dari outline riset, ditambah tiga risiko yang muncul saat menyusun PRD (butir 14-16). Mitigasi rinci ada di modul yang dirujuk.

| # | Risiko | Mitigasi utama | Modul |
| ---: | --- | --- | --- |
| 1 | **Ledakan cakupan.** Lebih dari 2.000 fitur inventaris berisiko membuat delivery tertunda tanpa batas. | Disiplin P0-P3 dan Fase 1-7 per fitur; Fase 1 dibatasi ke P0 inti; "Fase —" tidak dibangun. | semua |
| 2 | **Kompleksitas model izin** (scheme, role, grup, security level, visibilitas komentar) harus konsisten di API, pencarian, notifikasi, laporan, dan ekspor. Bila satu jalur terlewat, data bocor. | Satu layanan otorisasi terpusat (deny by default) dan uji izin otomatis per endpoint. | 03, 12, 27 |
| 3 | **Kinerja JQL → SQL dan custom field JSONB** pada skala jutaan baris (join berat, operator riwayat, sort custom field). | Indeks JSONB/tabel bertipe, guardrail kueri, dan opsi search engine eksternal. | 12, 13, 28 |
| 4 | **Konkurensi**: rank LexoRank bertabrakan saat drag-and-drop bersamaan, dan edit bersamaan menimpa perubahan. | Rebalancing rank dan optimistic locking (kolom `version`). | 10, 11, 27, 28 |
| 5 | **Utang teknis baseline**: migrasi `001` tidak valid, initdb menjalankan `down.sql`, path Makefile, `JWT_SECRET` lemah, CORS `*` + credentials, `GET /users` tanpa cek admin, `api.ts` usang, halaman frontend terhapus. | Remediasi Sprint 0 sebelum fitur baru. | 00, 01, 27 |
| 6 | **Infrastruktur pendukung belum ada** (job queue, email, storage, realtime, rate limit, audit). MVP butuh pekerjaan platform yang tidak terlihat pengguna. | Rencanakan kapasitas platform eksplisit di Fase 1-2. | 27, 28 |
| 7 | **Keamanan konten**: rich text, lampiran, smart link, dan web request otomasi/webhook membuka celah XSS, malware, dan SSRF. | Sanitasi server-side, pemindaian lampiran, dan proteksi SSRF dengan allowlist. | 07, 08, 20, 23 |
| 8 | **Merek dagang & hukum**: meniru nama, logo, atau istilah eksklusif Atlassian secara berlebihan. | VISI-001: satu sumber nama produk, pemindaian CI, dan disclaimer nominatif. | 00 |
| 9 | **Frontend**: belum ada library drag-and-drop yang aksesibel; editor kaya dan virtualisasi daftar besar menambah ukuran bundle. | Diputuskan: drag-and-drop dibangun sendiri sebagai modul reusable `lib/dnd` dengan target WCAG 2.1 AA ([ADR-0001](../adr/0001-library-drag-and-drop-board.md)); code-splitting dan virtualisasi tetap berlaku. | 07, 10, 21 |
| 10 | **Laporan historis** (burndown, CFD, control chart, sprint report) bergantung pada changelog lengkap sejak awal. Data historis yang tidak tercatat di MVP tidak dapat dipulihkan. | Changelog dan event domain wajib sejak Fase 1. | 08, 16, 28 |
| 11 | **Single binary** membatasi skala horizontal SSE dan job worker bila state tidak dipisahkan. | Pub/sub terpisah dan mode worker terpisah. | 27 |
| 12 | **Keputusan produk terlambat** (tenancy, format konten, gaya proyek, search engine) memaksa migrasi data besar di fase berikutnya. | Selesaikan butir B.1, D.1-D.5, dan G.1 di [Pertanyaan terbuka](#pertanyaan-terbuka-gabungan) sebelum Fase 1. | 00, 04, 27, 28 |
| 13 | **Cakupan pengujian**: baru ada unit test, tanpa integration/e2e, sehingga regresi pada alur inti berisiko tinggi. | Integration test dengan Postgres nyata dan e2e Playwright untuk alur kritis. | 27 |
| 14 | **Backlog verifikasi besar**: 2.429 butir aturan (15%) bertanda ⚠. Bila tidak divalidasi, perilaku clone bisa menyimpang dari ekspektasi pengguna migrasi. | Verifikasi bergelombang per fase (lihat [Backlog verifikasi](#backlog-verifikasi-)). | semua |
| 15 | **Ketidakselarasan lintas modul** (14 butir) dapat menghasilkan implementasi ganda atau bertentangan. | Selesaikan [daftar keselarasan](#keselarasan-lintas-modul-yang-perlu-diputuskan) sebelum fase terkait. | 05, 09, 11, 14, 22, 28, 30, 35 |
| 16 | **Ketergantungan Service Management dan Product Discovery pada fondasi Fase 1-3** (workflow, JQL, otomasi, notifikasi, API). Keterlambatan fondasi menggeser Fase 5-7. | Desain model data (28) yang sudah menyiapkan entitas JSM/JPD sejak awal. | 28, 29–37 |

## Backlog verifikasi (⚠)

**Arti tanda ⚠.** Butir aturan berawalan ⚠ menggambarkan perilaku Jira Cloud yang menurut riset multi-agent belum pasti. Penyebabnya antara lain dokumentasi resmi yang tidak menyebutnya, sumber yang bertentangan, perilaku yang berubah sejak 2025 (mis. terminologi work item/space), angka batas yang tidak dipublikasikan, atau fitur beta/rollout bertahap. Butir ini **tidak boleh diimplementasikan apa adanya** sebelum diverifikasi, dan pemilik produk boleh mengubahnya secara sadar. Fitur yang ditambahkan penulis di luar inventaris juga ditandai ⚠.

Total: **2.473 kemunculan ⚠** di 38 modul, dengan **2.429 butir aturan** berawalan ⚠ dari 16.248 baris bernomor (±15%).

Sebaran butir ber-⚠ menurut fase fiturnya:

| Fase | Butir ber-⚠ |
| --- | ---: |
| Fase 1 | 274 |
| Fase 2 | 381 |
| Fase 3 | 493 |
| Fase 4 | 381 |
| Fase 5 | 336 |
| Fase 6 | 167 |
| Fase 7 | 368 |
| Fase — | 29 |

| Modul | ⚠ (kemunculan) | Butir aturan ber-⚠ | Butir aturan | % butir ber-⚠ | Fase yang dilayani |
| --- | ---: | ---: | ---: | ---: | --- |
| [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) | 130 | 122 | 605 | 20% | Fase 2–7 |
| [29 Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) | 125 | 125 | 668 | 19% | Fase 5–7 |
| [33 Service Management: Assets (CMDB)](33-jsm-assets-cmdb.md) | 123 | 123 | 526 | 23% | Fase 7 |
| [35 Service Management: Virtual Service Agent, AI, Otomasi, API, Admin & Customer Service Management](35-jsm-ai-otomasi-api-admin-csm.md) | 116 | 115 | 533 | 22% | Fase 5–7 |
| [21 Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) | 114 | 113 | 556 | 20% | Fase 1–4 |
| [22 Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) | 110 | 110 | 479 | 23% | Fase 1–4 |
| [20 Otomasi Berbasis Aturan](20-otomasi.md) | 105 | 105 | 549 | 19% | Fase 3–4 |
| [30 Service Management: Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) | 99 | 97 | 615 | 16% | Fase 5–7 |
| [36 Product Discovery: Ideas, Fields, Formula & Views](36-jpd-ideas-field-views.md) | 99 | 99 | 703 | 14% | Fase 6 |
| [18 Timeline, Rencana Lintas Proyek & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) | 97 | 95 | 553 | 17% | Fase 2–4 |
| [10 Board: Kolom, Kartu, Swimlane & Kanban](10-board-kolom-kartu.md) | 94 | 93 | 577 | 16% | Fase 1–3 |
| [31 Service Management: Queues, SLA, Knowledge Base, CSAT & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) | 88 | 88 | 451 | 20% | Fase 5–7 |
| [34 Service Management: Operations (Alerts, On-call, Escalation)](34-jsm-operations-alerting-on-call.md) | 85 | 85 | 486 | 17% | Fase 7 |
| [32 Service Management: ITSM (Incident, Problem, Change, Services, Approvals)](32-jsm-itsm-insiden-problem-perubahan-layanan.md) | 76 | 76 | 564 | 13% | Fase 5–7 |
| [26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) | 75 | 75 | 324 | 23% | Fase 1–4 |
| [12 Bahasa Kueri (JQL-kompatibel)](12-jql-bahasa-kueri.md) | 73 | 73 | 518 | 14% | Fase 1–4 |
| [04 Proyek: Tipe, Template, Pengaturan & Siklus Hidup](04-proyek.md) | 68 | 68 | 507 | 13% | Fase 1–4 |
| [37 Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) | 67 | 67 | 452 | 15% | Fase 6 |
| [23 REST API Publik & Webhook](23-api-publik-webhook.md) | 63 | 55 | 702 | 8% | Fase 1–4 |
| [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) | 59 | 59 | 420 | 14% | Fase 1–4 |
| [06 Work Item Inti: Buat, Lihat, Ubah & Operasi](06-work-item-inti.md) | 57 | 57 | 411 | 14% | Fase 1–4 |
| [17 Dashboard & Gadget](17-dashboard-gadget.md) | 53 | 53 | 487 | 11% | Fase 3–4 |
| [14 Pelacakan Waktu, Worklog & Estimasi](14-pelacakan-waktu-estimasi.md) | 48 | 47 | 402 | 12% | Fase 1–4 |
| [19 Notifikasi In-app & Email](19-notifikasi-email.md) | 48 | 44 | 330 | 13% | Fase 1–4 |
| [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) | 47 | 47 | 264 | 18% | Fase 1–5 |
| [07 Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) | 46 | 46 | 263 | 17% | Fase 1–4 |
| [16 Laporan & Insight](16-laporan-insight.md) | 43 | 43 | 335 | 13% | Fase 1–4 |
| [05 Konfigurasi Work Item: Tipe, Field, Layar & Skema](05-konfigurasi-work-item-skema.md) | 42 | 42 | 475 | 9% | Fase 1–4 |
| [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) | 41 | 41 | 417 | 10% | Fase 1–4 |
| [15 Rilis (Versi) & Komponen](15-rilis-versi-komponen.md) | 37 | 37 | 241 | 15% | Fase 2–4 |
| [13 Pencarian, Navigator, Filter Tersimpan & Ekspor](13-pencarian-filter-ekspor.md) | 36 | 34 | 433 | 8% | Fase 1–4 |
| [24 Integrasi Pengembangan & DevOps](24-integrasi-devops.md) | 30 | 29 | 207 | 14% | Fase 4–7 |
| [11 Backlog & Sprint (Scrum)](11-backlog-sprint.md) | 28 | 24 | 195 | 12% | Fase 1–4 |
| [03 Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) | 21 | 15 | 438 | 3% | Fase 1–5 |
| [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) | 11 | 10 | 221 | 5% | Fase 1–4 |
| [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) | 8 | 8 | 157 | 5% | Fase 1–4 |
| [00 Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md) | 6 | 4 | 36 | 11% | Fase 1–3 |
| [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) | 5 | 5 | 148 | 3% | Fase 1–4 |
| **Total** | **2473** | **2429** | **16248** | **15%** | |

**Proses verifikasi.**

1. **Jadwal bergelombang.** Butir ⚠ diverifikasi paling lambat satu sprint sebelum fitur pemiliknya masuk sprint implementasi. Gelombang pertama mencakup 274 butir milik fitur Fase 1.
2. **Sumber utama.** Butir dicocokkan dengan dokumentasi resmi Atlassian: support.atlassian.com (panduan produk Jira, Service Management, dan Product Discovery), developer.atlassian.com (REST API v3, Agile API, Service Management API, Assets API, webhook, OAuth scope), serta halaman batas dan perbandingan plan.
3. **Sumber pendukung.** Bila dokumentasi tidak menyebut perilakunya, dilakukan uji langsung di instance pembanding (site uji Free atau trial) dan dicatat dengan tangkapan layar, payload API, dan tanggal uji. Changelog/blog produk dipakai hanya untuk perubahan terbaru.
4. **Hasil.** Untuk setiap butir ada tiga kemungkinan:
   - **terverifikasi:** hapus tanda ⚠ dan tambahkan rujukan sumber di "Catatan & ketergantungan";
   - **dikoreksi:** ubah aturan sesuai temuan, lalu hapus ⚠;
   - **keputusan produk:** perilaku pembanding tidak dapat dipastikan atau sengaja tidak diikuti. Tulis aturan clone sebagai keputusan sendiri dan pindahkan pertanyaannya ke "Pertanyaan terbuka" bila belum diputuskan.
5. **Pencatatan.** Perubahan ⚠ dilakukan lewat edit tertarget di file modul (bukan menulis ulang file). Jalankan `prdtool.py coverage <id>` setelah edit, lalu perbarui angka di README ini dengan skrip pada [Sumber & metodologi](#sumber--metodologi).
6. **Prioritas dalam gelombang.** Dahulukan butir pada fitur P0, butir yang menyangkut keamanan/izin, kontrak API publik (karena memengaruhi kompatibilitas klien), dan angka batas yang menentukan desain skema data.

## Glosarium ringkas

Dua puluh istilah inti berikut dipakai di semua modul. Glosarium lengkap (139 istilah dalam delapan kelompok, beserta padanan Indonesia dan modul pemilik) ada di [00 › Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md#glosarium).

| Istilah | Padanan Indonesia | Definisi singkat | Modul |
| --- | --- | --- | --- |
| Work item (Issue) | Item kerja | Unit kerja yang dilacak (task, bug, story, request, idea). "Issue" tetap dipakai di API dan JQL. | [06](06-work-item-inti.md) |
| Work type (Issue type) | Tipe kerja | Klasifikasi work item (Epic, Story, Task, Bug, Subtask, atau kustom) yang menentukan field, layar, dan workflow. | [05](05-konfigurasi-work-item-skema.md) |
| Project (Space) | Proyek (Ruang) | Wadah work item beserta konfigurasinya, dengan project key sebagai prefix issue key (mis. `WEB-42`). | [04](04-proyek.md) |
| Team-managed / company-managed | Dikelola tim / dikelola perusahaan | Dua gaya proyek. Team-managed punya konfigurasi mandiri per proyek, sedangkan company-managed memakai scheme bersama yang dikelola admin. | [04](04-proyek.md) |
| Epic | Epic | Work item besar yang menampung story/task; level hierarki di atas tipe standar. | [05](05-konfigurasi-work-item-skema.md) |
| Workflow | Alur kerja | Kumpulan status dan transition (dengan condition, validator, dan post function) yang dilalui work item. | [09](09-workflow-status-transisi.md) |
| Status category | Kategori status | Pengelompokan status menjadi To Do, In Progress, dan Done untuk board dan laporan. | [09](09-workflow-status-transisi.md) |
| Scheme | Skema | Konfigurasi bersama (workflow, field, screen, permission, notification) yang dipasang ke banyak proyek company-managed. | [05](05-konfigurasi-work-item-skema.md) |
| Custom field | Field kustom | Field tambahan buatan admin dengan tipe tertentu dan context per proyek/tipe. | [05](05-konfigurasi-work-item-skema.md) |
| Board | Papan | Visualisasi work item dalam kolom yang dipetakan ke status (Scrum atau Kanban). | [10](10-board-kolom-kartu.md) |
| Backlog | Backlog | Daftar terurut (rank) work item yang belum dikerjakan atau belum masuk sprint. | [11](11-backlog-sprint.md) |
| Sprint | Sprint | Iterasi berjangka waktu tetap berisi komitmen work item dan sprint goal. | [11](11-backlog-sprint.md) |
| JQL | Bahasa kueri (JQL-kompatibel) | Bahasa kueri `field operator value` untuk mencari work item. Nama UI ditetapkan di VISI-001. | [12](12-jql-bahasa-kueri.md) |
| Filter (saved filter) | Filter tersimpan | Kueri JQL yang disimpan dan dibagikan, lalu dipakai board, dashboard, atau langganan email. | [13](13-pencarian-filter-ekspor.md) |
| Project role | Peran proyek | Peran per proyek (Administrator, Member, Viewer) yang dipetakan ke izin lewat permission scheme. | [03](03-izin-peran-keamanan.md) |
| Automation rule | Aturan otomasi | Definisi tanpa kode berisi trigger, condition, dan action, dengan smart value `{{issue.summary}}`. | [20](20-otomasi.md) |
| Request type | Tipe permintaan | Jenis permintaan yang dipilih customer di portal, dipetakan ke work type dan form. | [29](29-jsm-service-project-request-type-formulir.md) |
| SLA | SLA (perjanjian tingkat layanan) | Target waktu (time to first response, time to resolution) dengan kalender kerja dan penanganan breach. | [31](31-jsm-antrian-sla-knowledge-base-laporan.md) |
| Assets / CMDB | Aset / CMDB | Basis data item konfigurasi (object schema, object type, object) beserta relasinya, dikueri dengan AQL. | [33](33-jsm-assets-cmdb.md) |
| Idea | Ide | Work item di proyek discovery yang mewakili peluang produk, dinilai dengan field/formula dan didukung insight. | [36](36-jpd-ideas-field-views.md) |

## Sumber & metodologi

**Riset inventaris (2026-11-05).** Inventaris fitur disusun lewat riset multi-agent terhadap dokumentasi publik Jira Cloud, Jira Service Management, dan Jira Product Discovery, dengan beberapa putaran kritik untuk mencari fitur yang terlewat. Hasilnya disimpan sebagai JSON read-only di [`_inventory/`](_inventory/):

| Berkas | Isi | Domain | Fitur | Sub-fitur | Putaran kritik (item baru) |
| --- | --- | ---: | ---: | ---: | --- |
| `_inventory/jira-core.json` | Jira terpadu, NFR, lampiran JSM/JPD, ditambah `outline` (modul, fase, persona, NFR, model data inti, pertanyaan terbuka, risiko) dan `baseline` repo | 26 | 1.578 | 12.234 | 5 (75) |
| `_inventory/jsm-jpd.json` | Service Management rinci (7 domain) dan Product Discovery rinci (2 domain) | 9 | 478 | 3.784 | 4 (67) |
| `_inventory/modules.json` | Alokasi domain/fitur ke 38 modul (`whole` atau `pick` per domain) | — | — | — | — |

**Dari inventaris ke modul.** 2.056 fitur inventaris (16.018 sub-fitur) dialokasikan ke 38 modul tanpa sisa (`prdtool.py check`: 0 galat, 0 unassigned). Setiap modul lalu ditulis dan direview oleh agent terpisah dengan template yang sama. Bila fitur dari domain berbeda tumpang tindih (mis. lampiran JSM vs domain JSM rinci, atau Board Copy vs Copy board & delete board), fitur-fitur itu digabung ke satu heading dengan baris **Alias inventaris**. Karena itu jumlah heading PRD (**1.861**) lebih kecil dari jumlah fitur inventaris (2.056), tanpa ada sub-fitur yang hilang. Pemeriksaan `prdtool.py coverage` atas seluruh modul memberi status **OK, 0 masalah**: setiap nama fitur muncul sebagai heading atau alias, prefix ID benar, tautan relatif valid, dan bagian penutup ada.

**Alat bantu** (`_inventory/prdtool.py`, hanya membaca):

```bash
cd docs/prd/_inventory
python3 -I prdtool.py check            # validasi alokasi modules.json, hitung per modul
python3 -I prdtool.py names <id>       # daftar nama fitur jatah modul
python3 -I prdtool.py extract <id>     # Markdown detail fitur jatah modul (hasil di _extract/<id>.md)
python3 -I prdtool.py coverage [<id>]  # cek file modul: heading, prefix ID, tautan, bagian penutup
```

**Menghitung ulang angka README.** Tabel [Daftar modul](#daftar-modul), [Matriks modul × fase](#matriks-modul--fase), dan [Backlog verifikasi](#backlog-verifikasi-) dihitung dari file modul dengan aturan berikut (jalankan dari `docs/prd`):

```bash
python3 -I - <<'PY'
import glob, re
for f in sorted(glob.glob('[0-9][0-9]-*.md')):
    t = open(f, encoding='utf-8').read(); L = t.split('\n')
    feat  = sum(bool(re.match(r'^### [A-Z][A-Z0-9-]*-[0-9]{3} ', l)) for l in L)
    rules = sum(bool(re.match(r'^\s*[0-9]+\. ', l)) for l in L)
    flags = t.count('⚠ ')
    phase = {}
    for l in L:
        if '**Prioritas / Fase / Tier:**' in l and '<P0-P3>' not in l:
            p = re.search(r'Fase\s*([0-9—])', l).group(1); phase[p] = phase.get(p, 0) + 1
    print(f[:2], feat, rules, flags, dict(sorted(phase.items())))
PY
```

Baris template di modul 00 (yang memuat `<P0-P3>`) dikecualikan dari hitungan fase. Butir aturan dihitung dari semua baris bernomor sesuai konvensi, sehingga angkanya sedikit lebih tinggi dari jumlah acceptance criteria murni.

**Konteks repo.** PRD ditulis terhadap kondisi repo `jari` pada branch `main` (commit `2ea8062`): backend Go Clean Architecture (module path `app`) dengan fitur auth dan user, migrasi golang-migrate, Swagger via swag/Scalar, dan SPA SvelteKit ter-embed lewat `go:embed`. Gap dan remediasinya dirinci di [00 › Baseline repo & gap](00-visi-ruang-lingkup-glosarium-baseline.md#baseline-repo--gap) serta di modul [27](27-nfr-arsitektur-platform.md) dan [28](28-model-data-inti.md).

**Batasan.** Inventaris mencerminkan dokumentasi publik pada tanggal riset. Perilaku yang tidak terdokumentasi ditandai ⚠ dan wajib diverifikasi (lihat [Backlog verifikasi](#backlog-verifikasi-)). Penyebutan Jira, Atlassian, dan produk pihak ketiga di PRD ini bersifat nominatif, untuk menjelaskan acuan paritas dan kompatibilitas.
