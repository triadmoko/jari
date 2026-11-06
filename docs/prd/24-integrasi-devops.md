# 24 — Integrasi Pengembangan & DevOps

Modul ini mendefinisikan bagaimana produk ini terhubung dengan alat pengembangan perangkat lunak (SCM seperti GitHub, GitLab, Bitbucket Cloud; CI/CD seperti GitHub Actions, GitLab CI, Bitbucket Pipelines, Jenkins, CircleCI, Azure Pipelines; serta penyedia feature flag, security scanner, desain, dan operations) sehingga data development (repository, branch, commit, pull request, build, deployment, feature flag, vulnerability, desain, incident) tertaut ke work item melalui issue key. Cakupannya meliputi model data dan kontrak API ingest (padanan Development Information, Builds, Deployments, Feature Flags, Security, Operations, Design API), kredensial OAuth untuk tool self-hosted, Toolchain dan konektor per penyedia, deteksi issue key, permission View Development Tools, permukaan visibilitas (Development panel di issue view, field/kolom Development di pencarian, indikator di kartu board, Work suggestions), halaman proyek Code, Deployments, dan Security, panel Feature flags, Designs, dan Operations, aksi dari kode (Smart commits, Workflow Triggers, trigger/action otomasi DevOps, deployment gating dengan Service Management), metrik DevOps (cycle time dan deployment frequency), template/onboarding DevOps, serta integrasi IDE dan agen AI coding opsional. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`); seluruh modul ini adalah pekerjaan baru yang dikemas sebagai feature module baru, misalnya `internal/features/devops` (ingest, connector, devinfo query) dengan entitas di `internal/shared/domain/entity` dan repository di `internal/shared/infrastructure/repository`, endpoint di bawah `/api/v1/devops/...` (path kompatibel produk referensi seperti `/rest/devinfo/0.10/bulk` dicatat sebagai target kompatibilitas, keputusan akhir di modul 23), webhook masuk per penyedia di `/api/v1/devops/webhooks/{provider}`, dan halaman Code, Deployments, Security, serta Toolchain sebagai route baru di SPA `frontend/`. Konektor dibangun oleh produk ini sendiri (mis. GitHub App milik produk ini); app/plugin yang dibuat untuk produk referensi tidak dapat dipakai langsung, sehingga nama app pihak lain di inventaris dicatat sebagai padanan fungsional.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission View Development Tools (PERM-027) dan View Aggregated Data (PERM-028).
- [04 Proyek](04-proyek.md) — entri menu Toolchain (PRJ-058, aturan dimiliki DEV-006), halaman Development Tools di project settings (PRJ-059), Apps (PRJ-060), template proyek.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field read-only Development & deployment info (CFG-023).
- [06 Work Item Inti](06-work-item-inti.md) — issue view tempat Development panel, Feature flags, Designs, dan Security dirender; web/remote link.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar dari Smart commits dan entri History.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — service transisi tunggal yang dipakai Smart commits dan Workflow Triggers.
- [10 Board](10-board-kolom-kartu.md) — card layout dan indikator dev status di kartu.
- [11 Backlog & Sprint](11-backlog-sprint.md) — kartu backlog dan panel sprint tempat Work suggestions muncul.
- [12 JQL](12-jql-bahasa-kueri.md) — field `development[...]` (JQL-035).
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — kolom Development di navigator dan List view.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — efek worklog Smart commits `#time` (TIME-044); parser smart commit dimiliki DEV-024.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Release warnings dan status development di version detail page, Compass/DevOps components (REL-035).
- [16 Laporan & Insight](16-laporan-insight.md) — tampilan laporan Cycle time (RPT-019) dan Deployment frequency (RPT-020).
- [19 Notifikasi & Email](19-notifikasi-email.md) — email kegagalan Smart commits.
- [20 Otomasi](20-otomasi.md) — DevOps triggers (AUTO-030), Build and deployment triggers (AUTO-031), Cross-product and DevOps actions (AUTO-055), koneksi (AUTO-056).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — halaman admin Apps, OAuth credentials, DVCS legacy.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — konvensi REST, autentikasi, rate limit, paginasi.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — kerangka ekstensi (padanan Connect/Forge), Marketplace, smart link; pemilik integrasi desain Figma (INT-048, dirujuk DEV-022) dan ekstensi IDE (INT-050, dirujuk DEV-031).
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — antrean ingest, idempotensi, retensi, fitur AI opsional.
- [28 Model Data Inti](28-model-data-inti.md) — entitas inti work item dan katalog event domain `devinfo.*`/`build.*`/`deployment.*`.
- [32 Service Management ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — change request, approval, change calendar untuk deployment gating; incident untuk korelasi deploy-incident.
- [34 Service Management Operations](34-jsm-operations-alerting-on-call.md) — alert dan on-call sebagai sumber data Operations.

## Persona terkait

- **Developer/Anggota Tim** — menyebut issue key di branch/commit/PR, membuat branch dari work item, memakai Smart commits, membaca Development panel, dev status di kartu, Work suggestions, dan ekstensi IDE.
- **Admin Situs** — memasang/menghubungkan konektor GitHub, GitLab, Bitbucket, CI/CD, membuat OAuth credentials, mencabut integrasi, mengatur permission scheme.
- **Admin Proyek/Project Lead** — mengatur Toolchain proyek, memilih repository yang dikaitkan, mengaktifkan halaman Code/Deployments/Security, membuat template DevOps, mengonfigurasi Workflow Triggers dan rule otomasi.
- **Scrum Master/Team Lead** — memantau PR menunggu review, build gagal, deployment per environment, dan metrik cycle time/deployment frequency tim.
- **Product Owner/PM** — melihat status rilis fitur (deployment, feature flag rollout) tanpa membuka tool development.
- **Integrator/Developer API** — mengirim data development, build, deployment, feature flag, security, dan operations lewat DevOps REST API dari tool internal atau self-hosted.
- **Manajer Layanan/ITSM** — menyetujui change request yang dibuat otomatis oleh deployment gating dan mengorelasikan deployment dengan incident.
- **Stakeholder/Viewer** — membaca halaman Deployments dan status development tanpa akses ke repository.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin branch, commit, dan pull request yang menyebut issue key otomatis tertaut ke work item agar konteks kode terlihat tanpa pencatatan manual (DEV-002, DEV-013).
- Sebagai Admin Situs, saya ingin menghubungkan organisasi GitHub, group GitLab, dan workspace Bitbucket Cloud beserta backfill data historis agar seluruh repository tim langsung terlihat (DEV-007, DEV-008, DEV-009).
- Sebagai Integrator/Developer API, saya ingin mengirim data build dan deployment dari Jenkins atau pipeline self-hosted dengan OAuth credentials agar tool internal ikut tampil di Development panel (DEV-004, DEV-005, DEV-010).
- Sebagai Developer/Anggota Tim, saya ingin membuat branch dengan nama berisi issue key langsung dari work item agar penautan selalu benar sejak awal (DEV-014).
- Sebagai Developer/Anggota Tim, saya ingin menulis `#comment`, `#time`, dan `#<transisi>` di pesan commit agar work item ter-update tanpa meninggalkan terminal (DEV-024).
- Sebagai Admin Proyek/Project Lead, saya ingin transisi workflow berjalan otomatis saat PR dibuat atau di-merge agar status board selalu mencerminkan kondisi kode (DEV-025, DEV-026).
- Sebagai Scrum Master/Team Lead, saya ingin memfilter work item dengan PR terbuka atau build gagal dan melihat ikon status di kartu board agar hambatan terlihat saat stand-up (DEV-015, DEV-016, DEV-017).
- Sebagai Product Owner/PM, saya ingin halaman Deployments menampilkan work item apa saja yang sudah sampai ke Production agar saya tahu fitur mana yang sudah tersedia bagi pengguna (DEV-019, DEV-003).
- Sebagai Admin Proyek/Project Lead, saya ingin Toolchain proyek menampilkan tool per tahap Plan/Code/Build/Deploy/Operate agar konfigurasi integrasi proyek mudah diaudit (DEV-006, DEV-018).
- Sebagai Admin Situs, saya ingin data development hanya tampil bagi pengguna dengan View Development Tools agar informasi repository privat tidak bocor ke viewer anonim atau pelanggan (DEV-012).
- Sebagai Developer/Anggota Tim, saya ingin melihat vulnerability, feature flag, dan desain yang tertaut di work item agar keputusan kerja memakai konteks lengkap (DEV-020, DEV-021, DEV-022).
- Sebagai Manajer Layanan/ITSM, saya ingin deployment ke Production ditahan sampai change request disetujui agar perubahan berisiko terkendali (DEV-027).
- Sebagai Scrum Master/Team Lead, saya ingin metrik cycle time dan deployment frequency dihitung dari data commit dan deployment agar tim dapat memperbaiki alur pengiriman (DEV-028, DEV-029).
- Sebagai Admin Proyek/Project Lead, saya ingin template proyek DevOps yang sudah menyiapkan Toolchain dan rule otomasi agar tim baru siap bekerja dalam satu langkah (DEV-030).
- Sebagai Developer/Anggota Tim, saya ingin mengirim work item beserta konteksnya ke agen AI coding eksternal yang sudah didaftarkan admin dan meninjau PR hasilnya di Development panel agar pekerjaan rutin dapat didelegasikan tanpa keluar dari produk ini (DEV-033).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 0 |
| P1 | 7 |
| P2 | 12 |
| P3 | 14 |
| **Total** | **33** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1-3 | 0 |
| Fase 4 | 31 |
| Fase 7 | 1 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: modul ini berada di Fase 4 ("DevOps & ekosistem") karena bergantung pada work item, workflow, JQL, otomasi, dan REST API publik yang selesai di Fase 1-3; key permission View Development Tools sudah didaftarkan di Fase 2 (PERM-027) agar penegakannya langsung tersedia saat data mulai masuk. Di dalam Fase 4, urutan pertama adalah fondasi data: skema dev info dan deteksi issue key (DEV-001, DEV-002, DEV-003), API ingest beserta OAuth credentials (DEV-004, DEV-005), lalu konektor GitHub, Bitbucket Cloud, dan GitLab berbasis webhook bertanda tangan (DEV-007, DEV-009, DEV-008). Gelombang kedua membangun permukaan baca bernilai tertinggi: Development panel, Create branch, Smart commits, field `development[...]`, dan dev status di kartu (DEV-013 s.d. DEV-016, DEV-024), disusul Toolchain, halaman Code dan Deployments, Workflow Triggers, serta trigger otomasi DevOps (DEV-006, DEV-018, DEV-019, DEV-025, DEV-026). Gelombang ketiga (P3) menambah panel opsional berbasis API generik (Security, Feature flags, Designs, Operations), metrik, template DevOps, Work suggestions, IDE, agen AI bawaan (DEV-032), serta konektor generik ke agen AI coding eksternal (DEV-033). Deployment gating ditunda ke Fase 7 karena memerlukan change management Service Management (modul 32), sedangkan DVCS accounts legacy hanya dicatat sebagai rujukan.

## Fitur

## Fondasi data, penautan & API ingest

### DEV-001 Data Model: Development Information & Deployments

- **Deskripsi:** Skema penyimpanan informasi development dari tool SCM/CI (repository, branch, commit, pull request, build, deployment, feature flag) yang ditautkan ke work item lewat issue key, beserta agregat ringkasan per work item untuk panel dan JQL.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Engineering (implementasi), Integrator/Developer API (pengirim data); data hanya dapat dibaca pengguna dengan View Development Tools (DEV-012).
- **Aturan perilaku (acceptance criteria):**
  1. Entitas dev info terdiri dari `dev_provider_connection` (penyedia + instance), `dev_repository`, `dev_branch`, `dev_commit`, dan `dev_pull_request` (status `OPEN`/`MERGED`/`DECLINED`, plus `DRAFT` sebagai atribut); setiap entitas menyimpan `external_id` unik per koneksi, URL ke tool asal, dan daftar issue key yang terdeteksi dari nama branch, pesan commit, atau judul/source branch PR (DEV-002).
  2. Entitas `dev_build` (pipeline, build number, state, jumlah test) dan `dev_deployment` (pipeline, environment dengan type `unmapped`/`development`/`testing`/`staging`/`production`, state) ditautkan ke work item melalui tabel relasi `dev_entity_issue` (entity_type, entity_id, issue_id) sehingga satu entitas dapat tertaut ke banyak work item dan sebaliknya.
  3. Sistem memelihara tabel ringkasan `issue_dev_summary` per work item (jumlah branch, commit, PR per status, build per state, deployment terakhir per environment type, jumlah feature flag) yang diperbarui secara asinkron maksimal 60 detik setelah ingest, dan menjadi sumber Development panel (DEV-013), kolom/JQL `development[...]` (DEV-015, JQL-035), dan indikator kartu (DEV-016).
  4. Ingest melalui webhook GitHub/GitLab/Bitbucket wajib lolos verifikasi signature (HMAC `X-Hub-Signature-256` untuk GitHub, `X-Gitlab-Token` untuk GitLab, signature/secret webhook Bitbucket); payload dengan signature tidak valid ditolak HTTP 401 dan dicatat di log integrasi tanpa menyimpan data. Idempotensi dijamin dengan `updateSequenceNumber` atau ID delivery webhook: event dengan sequence lebih kecil atau sama dari yang tersimpan diabaikan tanpa galat.
  5. Mencabut (disconnect) integrasi menghapus seluruh entitas dev info milik koneksi tersebut beserta relasi dan ringkasannya dalam proses latar; ringkasan work item terdampak dihitung ulang, dan progres penghapusan terlihat di halaman integrasi sampai selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan; data disimpan per site dan dibaca per proyek melalui issue key.
- **Catatan & ketergantungan:** Event domain `devinfo.updated`, `build.updated`, `deployment.updated` didefinisikan di modul 28 dan dipakai otomasi (DEV-026). Entitas work item inti di modul 28; antrean ingest, partisi, dan retensi di modul 27. Migrasi: tabel baru, tidak ada data lama di repo.

### DEV-002 Issue key linking

- **Deskripsi:** Mekanisme penautan data development ke work item dengan menyebut issue key (mis. `PROJ-123`) di nama branch, pesan commit, judul PR, atau data build/deployment.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (menulis issue key), Sistem (parser ingest); tidak ada izin khusus untuk penautan, tetapi menampilkannya memerlukan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Issue key dikenali dengan pola `[A-Z][A-Z0-9_]+-[0-9]+` dengan batas kata di kedua sisi; key huruf kecil (`proj-123`) tidak ditautkan, dan key yang proyeknya tidak ada di site diabaikan tanpa galat.
  2. Branch ditautkan bila nama branch memuat issue key di posisi mana pun, termasuk setelah prefiks (mis. `feature/PROJ-123-login` dan `PROJ-123_fix`).
  3. Commit ditautkan bila pesan commit (judul maupun badan) memuat issue key.
  4. Pull request ditautkan bila issue key ada di judul, nama source branch, atau pesan salah satu commit di dalam PR; PR tetap tertaut selama salah satu sumber tersebut masih memuat key.
  5. Build dan deployment ditautkan ke issue key yang diturunkan dari commit/branch yang dibangun atau dideploy, atau dari daftar `issueKeys`/`associations` yang dikirim eksplisit melalui API (DEV-004); keduanya digabung tanpa duplikasi.
  6. Satu commit atau PR yang memuat beberapa issue key (mis. `PROJ-1 PROJ-2 fix`) ditautkan ke semua work item tersebut dan muncul di Development panel masing-masing.
  7. ⚠ Setelah project key diganti (rename), key lama tetap dikenali karena sistem menyimpan riwayat key proyek; commit baru yang memakai key lama tetap ditautkan ke work item yang sama.
  8. Saat koneksi pertama dibuat, konektor menjalankan backfill data historis (rentang dan dukungan tergantung penyedia, DEV-007 s.d. DEV-009); work item yang dipindah ke proyek lain tetap mempertahankan dev info karena relasi disimpan berdasarkan ID work item, bukan teks key.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Riwayat key proyek di modul 04 (rename project key). Move work item di modul 06 (ISS-032). Pola key harus sama dengan validasi project key di modul 04.

### DEV-003 Deployment environments

- **Deskripsi:** Model environment yang dikirim tool CI/CD dan dipetakan ke tipe environment standar, beserta state deployment, pipeline, dan aturan urutan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** DevOps engineer, Integrator/Developer API (mengirim); pengguna dengan View Development Tools (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Environment memiliki atribut `id` (unik per pipeline), `displayName`, dan `type` dengan nilai `unmapped`, `development`, `testing`, `staging`, atau `production`; nilai type di luar daftar ditolak HTTP 400 dengan pesan galat per field.
  2. State deployment yang diterima: `unknown`, `pending`, `in_progress`, `cancelled`, `failed`, `rolled_back`, `successful`; state lain ditolak HTTP 400.
  3. Pipeline memiliki `id`, `displayName`, dan `url`; deployment tanpa pipeline ditolak.
  4. `deploymentSequenceNumber` menentukan urutan deployment dalam satu pipeline+environment, sedangkan `updateSequenceNumber` menentukan versi data satu deployment; update dengan `updateSequenceNumber` yang lebih kecil diabaikan (idempotensi).
  5. ⚠ Association deployment mendukung tipe `issueKeys`, `issueIdOrKeys`, `serviceIdOrKeys` (komponen DevOps/layanan, REL-035), dan `commit`/`repository`; association commit/repository diselesaikan menjadi issue key melalui commit tersimpan (DEV-002).
  6. Deployment terbaru per environment (berdasarkan `deploymentSequenceNumber` lalu `lastUpdated`) ditampilkan di Development panel (DEV-013) dan Deployments page (DEV-019); deployment `rolled_back` menggantikan status sebelumnya di environment yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai deployment gating (DEV-027) dan metrik (DEV-028, DEV-029). Pemetaan environment Bitbucket/GitHub/GitLab ke type dijelaskan di konektor masing-masing.

### DEV-004 DevOps REST APIs (Software Cloud)

- **Deskripsi:** Kumpulan API publik untuk mengirim data development, build, deployment, feature flag, security, operations, dan desain ke produk ini dari tool apa pun, termasuk pipeline internal.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, Developer; autentikasi memakai OAuth credentials (DEV-005) atau token ekstensi; setiap kredensial dibatasi ke kategori data yang diizinkan.
- **Aturan perilaku (acceptance criteria):**
  1. Development Information API (padanan `devinfo/0.10`) menerima bulk submit repository beserta commits, branches, dan pull requests dalam satu request, serta menyediakan delete per repository, per entity (`commit`/`branch`/`pull_request` + ID), dan per properti (mis. `?accountId=...`) yang menghapus semua entity berproperti tersebut.
  2. Builds API (padanan `builds/0.1`) menerima build per pipeline dengan `buildNumber`, `state` (`pending`, `in_progress`, `successful`, `failed`, `cancelled`, `unknown`), `testInfo` (`totalNumber`, `numberPassed`, `numberFailed`, `numberSkipped`), dan `issueKeys`; build tanpa issue key valid diterima tetapi dilaporkan di `unknownIssueKeys` pada respons.
  3. Deployments API (padanan `deployments/0.1`) menerima deployment, environment, dan pipeline sesuai DEV-003 serta menyediakan endpoint gating status (DEV-027).
  4. Feature Flags API (padanan `featureflags/0.1`) menerima flag dengan key, display name, issue key, dan status per environment (DEV-021).
  5. Security Information API menerima security workspace (container) dan vulnerability dengan severity, status, identifier, dan association ke issue (DEV-020).
  6. Operations API menerima incident dan post-incident review beserta association ke issue dan service (DEV-023).
  7. Remote Links API menerima tautan generik (mis. dokumen, alert) yang ditampilkan di work item, dan DevOps Components API menerima komponen/layanan yang dapat dirujuk association deployment (REL-035).
  8. ⚠ Design API menerima desain dengan URL, nama, thumbnail, dan status (DEV-022); versi kontrak ditetapkan produk ini sebagai `designs/1.0`.
  9. Autentikasi menerima token dari OAuth credentials (client credentials, DEV-005) atau token modul ekstensi yang mendeklarasikan kapabilitas setara `jiraDevelopmentTool`, `jiraBuildInfoProvider`, `jiraDeploymentInfoProvider`, `jiraFeatureFlagInfoProvider`, `jiraSecurityInfoProvider` (modul 25); token tanpa kapabilitas untuk kategori yang dikirim ditolak HTTP 403.
  10. Setiap entitas wajib membawa `updateSequenceNumber`; data dengan nilai lebih kecil dari yang tersimpan diabaikan dan dilaporkan di respons. Batas yang ditetapkan produk ini: payload maksimal 1 MB per request, maksimal 100 entitas per bulk request, rate limit per kredensial 100 request/menit dengan respons HTTP 429 dan header `Retry-After` saat terlampaui.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi umum REST, galat, dan rate limit di modul 23. Ingest diproses asinkron melalui antrean (modul 27); respons `202 Accepted` mengembalikan `acceptedEntities`, `rejectedEntities`, dan `unknownIssueKeys`. Keputusan terbuka: apakah path kompatibel produk referensi diekspos selain `/api/v1/devops/...`.

### DEV-005 OAuth credentials for CI/CD

- **Deskripsi:** Kredensial OAuth yang dibuat admin untuk tool on-premise/self-hosted (mis. Jenkins) agar dapat mengirim build dan deployment info lewat API.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (membuat dan mencabut); tool CI/CD (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman berada di Pengaturan situs > Apps > OAuth credentials dan hanya dapat diakses Admin Situs; pengguna lain menerima HTTP 403.
  2. Form Create credential berisi nama app (wajib, maks 255 karakter), server base URL (wajib, URL valid http/https), logo URL (opsional), dan permission berupa checkbox Builds dan/atau Deployments (minimal satu wajib dipilih).
  3. Setelah dibuat, sistem menghasilkan Client ID dan Client secret; secret hanya ditampilkan sekali dengan tombol salin, disimpan dalam bentuk hash, dan tidak dapat dilihat lagi (hanya dapat dibuat ulang).
  4. Menghapus credential langsung mencabut akses: token yang sudah diterbitkan dari credential tersebut ditolak HTTP 401 pada request berikutnya, dan dialog konfirmasi menyebut nama app yang akan kehilangan akses.
  5. Credential dipakai Jenkins dan tool self-hosted lain dengan alur OAuth 2.0 client credentials; token akses berlaku 15 menit dan cakupannya terbatas ke kategori permission yang dipilih (token Builds tidak dapat mengirim deployment).
- **Team-managed vs company-managed:** Tidak ada perbedaan; credential berlaku site-wide.
- **Catatan & ketergantungan:** Server OAuth 2.0 produk ini di modul 01/23 (AUTH-013 untuk consent aplikasi terhubung). Audit log pembuatan/penghapusan di modul 22.

## Koneksi tool, Toolchain & konektor

### DEV-006 Toolchain

- **Deskripsi:** Pengaturan proyek untuk melihat dan menghubungkan tool development berdasarkan kategori alur DevOps, serta memilih repository spesifik untuk proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead (mengaitkan tool/repository ke proyek), Admin Situs (memasang konektor di level site).
- **Aturan perilaku (acceptance criteria):**
  1. Toolchain berada di Project settings > Toolchain untuk proyek software dan hanya dapat dibuka Admin Proyek (Administer Projects atau role Administrator); proyek bisnis dan service tidak menampilkan menu ini.
  2. ⚠ Tool dikelompokkan ke kategori tetap Plan, Code, Build, Deploy, dan Operate (ditambah Security dan Design sebagai kategori opsional bila penyedianya terpasang).
  3. Setiap kategori menampilkan tool yang sudah terhubung (nama, ikon, status koneksi) dan tombol "Add" yang menawarkan pilihan: menghubungkan konektor yang sudah terpasang di site atau membuka halaman pemasangan konektor baru.
  4. Kategori Code memungkinkan memilih repository spesifik dari konektor terhubung untuk dikaitkan dengan proyek; repository yang dikaitkan tampil di Code page (DEV-018) meskipun belum ada commit ber-issue key.
  5. Kategori Operate menghubungkan service project/Operations milik produk ini (modul 32, 34) atau penyedia operations eksternal sebagai sumber incident (DEV-023).
  6. Pemasangan konektor baru memerlukan Admin Situs; bila Admin Proyek non-admin situs menekan "Add" untuk konektor yang belum terpasang, tampil pesan "Minta admin situs memasang integrasi ini" beserta tombol kirim permintaan.
- **Team-managed vs company-managed:** Tidak ada perbedaan; di team-managed akses mengikuti role Administrator.
- **Catatan & ketergantungan:** DEV-006 adalah spesifikasi pemilik (source of truth) Toolchain; PRJ-058 Toolchain di [04 Proyek](04-proyek.md) hanya entri menu Project settings yang merujuk ke sini, dan bila ada perbedaan aturan, DEV-006 berlaku. Aturan putus hubungan tool dari project dan ringkasan sinkron per tool (PRJ-058 butir 3-4) berlaku sebagai bagian Toolchain ini. Halaman Development Tools di project settings (PRJ-059) menjadi bagian dari Toolchain. Template DevOps mengaktifkan Toolchain (DEV-030).

### DEV-007 GitHub for Jira

- **Deskripsi:** Konektor GitHub milik produk ini (padanan fungsional app "GitHub for Jira") untuk menghubungkan GitHub Cloud dan GitHub Enterprise Server melalui GitHub App, mencakup repository, branch, commit, PR, GitHub Actions, dan deployment.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (memasang, menghubungkan organisasi, mengelola backfill); Developer/Anggota Tim (memakai data).
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs memasang konektor GitHub dari halaman Apps lalu menghubungkan satu atau lebih organisasi GitHub melalui alur instalasi GitHub App; organisasi yang terhubung tampil di daftar beserta akun GitHub yang memasang.
  2. Akses repository dapat dipilih "All repositories" atau "Only select repositories" sesuai pengaturan instalasi GitHub App; perubahan pilihan di GitHub tersinkron otomatis melalui event `installation_repositories`.
  3. Backfill data historis dijalankan setelah koneksi dengan rentang waktu yang dapat dipilih (mis. 1 bulan, 6 bulan, 1 tahun, semua); backfill berjalan di latar dan menghormati rate limit GitHub.
  4. Sinkronisasi real-time via webhook mencakup branch (`create`/`delete`), commit (`push`), pull request (opened/edited/closed/merged/reopened, reviewer dan status review), workflow run GitHub Actions sebagai build, dan deployment/deployment_status sebagai deployment dengan pemetaan environment ke type DEV-003.
  5. GitHub Enterprise Server dihubungkan dengan GitHub App yang didaftarkan manual oleh admin: form meminta base URL server, App ID, client ID/secret, private key, dan webhook secret; koneksi divalidasi sebelum disimpan dan galat koneksi ditampilkan dengan pesan spesifik (URL tak terjangkau, kredensial salah).
  6. Create branch dari work item ke repository GitHub didukung (DEV-014) memakai token instalasi GitHub App.
  7. Smart commits diproses dari pesan commit GitHub (DEV-024).
  8. Halaman koneksi menampilkan status sinkronisasi per organisasi (Pending, In progress, Finished, Failed) dengan jumlah repository dan waktu sinkron terakhir, serta aksi "Restart backfill" dan "Retry" untuk koneksi gagal.
  9. ⚠ Alert Dependabot dan code scanning (GitHub Advanced Security) dikirim ke halaman dan panel Security (DEV-020) bila admin mengaktifkan opsi security di koneksi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App "GitHub for Jira" buatan Atlassian hanya berfungsi untuk produk referensi sehingga produk ini membangun GitHub App sendiri. Konektor GitHub DVCS lama dicatat di DEV-011. Data model di DEV-001.

### DEV-008 GitLab for Jira Cloud

- **Deskripsi:** Konektor GitLab milik produk ini (padanan fungsional app "GitLab for Jira Cloud") untuk GitLab.com dan GitLab self-managed, mencakup branch, commit, merge request, pipeline, deployment, dan feature flag.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (menghubungkan group/namespace), Developer/Anggota Tim.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs menghubungkan GitLab dengan OAuth application atau group access token dan memilih satu atau lebih group/namespace; seluruh project di namespace terpilih disinkronkan.
  2. GitLab self-managed didukung dengan mengisi URL instance (https wajib) dan kredensial; URL yang tidak terjangkau ditolak dengan pesan "Instance GitLab tidak dapat dihubungi".
  3. ⚠ Sinkronisasi via webhook group mencakup branch, commit (push), merge request (dipetakan sebagai pull request), pipeline (dipetakan sebagai build), deployment (dengan environment tier GitLab dipetakan ke type DEV-003), dan feature flag GitLab (dipetakan ke DEV-021).
  4. Create branch dari work item ke project GitLab didukung (DEV-014).
  5. ⚠ Smart commits diproses dari pesan commit GitLab dengan aturan yang sama dengan penyedia lain (DEV-024).
  6. Alternatif integrasi di sisi GitLab (fitur integrasi issue tracker bawaan GitLab yang menambah tautan dan komentar otomatis) dapat diarahkan ke REST API produk ini; produk ini mendokumentasikan endpoint yang kompatibel tetapi konektor native tetap jalur yang direkomendasikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App "GitLab for Jira Cloud" dikembangkan GitLab untuk produk referensi; produk ini membangun konektor sendiri.

### DEV-009 Bitbucket Cloud integration

- **Deskripsi:** Konektor Bitbucket Cloud untuk repository, branch, commit, pull request, Bitbucket Pipelines, dan Bitbucket Deployments.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (menghubungkan workspace), Developer/Anggota Tim.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs menghubungkan satu atau lebih Bitbucket workspace melalui OAuth consumer/app password workspace dari halaman Apps produk ini.
  2. Branch, commit, PR (termasuk reviewer dan approval), build Pipelines, dan deployment disinkronkan ke Development panel melalui webhook workspace dan backfill awal.
  3. Create branch dari work item mendukung pemilihan branch type Bitbucket (feature, bugfix, hotfix, release, custom) dan source branch (DEV-014).
  4. Smart commits diproses dari pesan commit Bitbucket (DEV-024).
  5. Tampilan work item produk ini di halaman PR/branch Bitbucket (lihat, buat, dan transisi) memerlukan ekstensi di sisi Bitbucket; produk ini menyediakan endpoint REST yang dibutuhkan, sedangkan ekstensi UI di Bitbucket dicatat sebagai keputusan terbuka.
  6. Environment Bitbucket Deployments (Test, Staging, Production) dipetakan ke type `testing`, `staging`, `production`; environment kustom dipetakan ke `unmapped` sampai admin memetakannya.
  7. Deployment gating dengan Service Management untuk Bitbucket Pipelines didukung melalui endpoint gating status (DEV-027).
  8. ⚠ Bitbucket Data Center/Server dihubungkan dengan konektor terpisah memakai personal access token dan webhook server; dukungan ini P3 dan tidak memakai application link.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bitbucket adalah produk vendor lain; produk ini hanya membangun konektornya. Pemetaan environment memakai DEV-003.

### DEV-010 Azure DevOps, Jenkins, CircleCI and other CI/CD integrations

- **Deskripsi:** Dukungan tool CI/CD lain yang mengirim build dan deployment info melalui DevOps REST API, baik lewat plugin/orb yang disediakan produk ini maupun skrip generik.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** DevOps engineer, Admin Situs (membuat OAuth credentials).
- **Aturan perilaku (acceptance criteria):**
  1. Jenkins: produk ini menyediakan plugin Jenkins sumber terbuka (padanan plugin "Atlassian Jira Software Cloud") yang memakai OAuth credentials (DEV-005) dan menyediakan step pipeline `sendBuildInfo` dan `sendDeploymentInfo` (padanan `jiraSendBuildInfo`/`jiraSendDeploymentInfo`) dengan parameter site, environment, dan issue key opsional.
  2. CircleCI: produk ini menyediakan orb yang mengirim build dan deployment info di akhir job dengan memakai OAuth credentials.
  3. ⚠ Azure Pipelines: dukungan disediakan melalui task Azure Pipelines yang memanggil Builds/Deployments API; integrasi Azure Repos (commit/PR) dipertimbangkan sebagai konektor terpisah di keputusan terbuka.
  4. Octopus Deploy, Harness, AWS CodePipeline, dan tool lain terhubung lewat Deployments API dengan contoh skrip `curl` yang didokumentasikan.
  5. Semua sumber CI/CD menghasilkan data yang tampil seragam di Development panel (DEV-013) dan Deployments page (DEV-019) dengan nama provider sesuai OAuth credential yang dipakai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App pihak ketiga (mis. Git Integration for Jira, app Azure Pipelines buatan Microsoft untuk produk referensi) termasuk ekosistem dan di luar cakupan.

### DEV-011 DVCS accounts (legacy Development tools)

- **Deskripsi:** Halaman admin legacy produk referensi untuk menghubungkan akun Bitbucket/GitHub melalui DVCS connector; tidak dibangun, dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman setara "Jira settings > Products > DVCS accounts" tidak dibangun; tautan dari PRJ-059 ke "konfigurasi akun DVCS" diarahkan ke halaman Apps/konektor (DEV-007 s.d. DEV-009).
  2. Opsi autolink repository baru digantikan pilihan "All repositories" di konektor (DEV-007 butir 2) dan sinkronisasi namespace (DEV-008).
  3. Toggle smart commits per akun dan per repository digantikan toggle per koneksi dan per repository di DEV-024.
  4. Enable/disable repository dan sinkronisasi manual digantikan aksi "Restart backfill"/"Retry" di konektor.
  5. Penghapusan akun digantikan disconnect koneksi beserta penghapusan data (DEV-001 butir 5).
  6. Konektor GitHub/GitHub Enterprise via DVCS sudah deprecated di produk referensi dan diganti app; produk ini tidak menyediakan jalur DVCS sama sekali.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alasan tidak dibangun: deprecated dan digantikan konektor modern. Lihat bagian Di luar cakupan.

## Visibilitas development di work item, pencarian & board

### DEV-012 View Development Tools permission

- **Deskripsi:** Penegakan permission proyek View Development Tools di semua permukaan yang menampilkan informasi development.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead dan Admin Situs (mengatur grant); semua pengguna (dibatasi olehnya).
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: permission "View Development Tools" di permission scheme (PERM-027) menjadi satu-satunya penentu akses dev info per proyek; default scheme memberikannya ke role Developers dan Administrators.
  2. Permission mengontrol Development panel (DEV-013), kolom dan JQL Development (DEV-015), release warnings dan status development di version detail page (REL-018), halaman Code, Deployments, dan Security (DEV-018 s.d. DEV-020), dev status di kartu board/backlog (DEV-016), serta Work suggestions (DEV-017).
  3. ⚠ Team-managed: akses diatur lewat kapabilitas role "Access development tools"; role Administrator dan Member memilikinya secara default, role Viewer tidak.
  4. Tanpa permission, panel, kolom, ikon kartu, dan menu halaman DevOps disembunyikan (bukan dinonaktifkan), endpoint baca dev info mengembalikan HTTP 403, dan klausa JQL `development[...]` tidak mencocokkan work item mana pun untuk pengguna tersebut (sama dengan DEV-015 butir 5).
  5. Grant ke "Public"/anonim dan ke customer portal ditolak (sesuai PERM-027 butir 3); definisi permission scheme lengkap ada di modul 03.
- **Team-managed vs company-managed:** Company-managed memakai permission scheme; team-managed memakai kapabilitas role (butir 3).
- **Catatan & ketergantungan:** Key permission didaftarkan di Fase 2 oleh PERM-027; penegakan di permukaan modul ini berlaku sejak Fase 4.

### DEV-013 Development panel (issue view)

- **Deskripsi:** Panel di issue view yang merangkum data development dari tool terhubung dan membuka dialog detail per jenis data.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, reviewer, release manager; memerlukan View Development Tools (DEV-012).
- **Aturan perilaku (acceptance criteria):**
  1. Ringkasan menampilkan baris terpisah: jumlah branch, jumlah commit, pull request (jumlah dan lozenge status terbaru `OPEN`/`MERGED`/`DECLINED`), build (status terbaru dan jumlah), deployment (environment terakhir yang dicapai beserta state), dan feature flag (bila ada); baris tanpa data tidak ditampilkan.
  2. Klik baris ringkasan membuka dialog Development details dengan tab: Branches (nama, repository, tautan "Create pull request"), Commits (hash pendek, pesan, author, tanggal, jumlah file, dikelompokkan per repository), Pull requests (judul, status, author, reviewer dan status approval, update terakhir, branch tujuan), Builds (pipeline, status successful/failed/in progress/cancelled/unknown, jumlah test passed/failed/skipped), dan Deployments (environment dan type, pipeline, state, waktu).
  3. Bila beberapa tool/instance terhubung, data di dialog dikelompokkan per tool/instance dengan nama dan ikon penyedia, dan setiap item memiliki tautan yang membuka item di tool asal pada tab baru.
  4. Aksi "Create branch" (DEV-014) dan "Create commit" tersedia; "Create commit" menampilkan contoh perintah git berisi issue key (mis. `git commit -m "PROJ-123 <pesan>"`) dengan tombol salin.
  5. Panel tampil bila work item punya data dev atau proyek memiliki tool terhubung; bila belum ada tool terhubung, panel menampilkan ajakan "Hubungkan tool development" yang mengarah ke Toolchain (DEV-006) untuk Admin Proyek, dan disembunyikan bagi pengguna lain.
  6. Panel hanya tampil untuk pengguna dengan View Development Tools pada proyek work item; data dimuat dari ringkasan DEV-001 dan dialog memuat detail secara lazy.
  7. Panel tersedia di proyek software company-managed dan team-managed.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission (DEV-012).
- **Catatan & ketergantungan:** Field locked "Development" di CFG-023; tata letak issue view di modul 06.

### DEV-014 Create branch from issue

- **Deskripsi:** Membuat branch di repository terhubung langsung dari work item dengan nama yang mengandung issue key.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan View Development Tools; pembuatan langsung juga memerlukan akses tulis di repository tujuan pada tool asal.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Create branch" tersedia di Development panel (DEV-013) dan menu aksi issue view.
  2. Dialog meminta pemilihan tool (GitHub, GitLab, Bitbucket) dan repository dari daftar repository terhubung, dengan pencarian nama repository.
  3. Nama branch default berformat `<ISSUE-KEY>-<slug-summary>` (huruf kecil, spasi menjadi tanda hubung, karakter non-alfanumerik dibuang, maks 100 karakter) dan dapat diedit; bila issue key dihapus dari nama, tampil peringatan "Branch tanpa issue key tidak akan tertaut".
  4. Untuk Bitbucket, dialog menyediakan pilihan branch type (feature, bugfix, hotfix, release) yang menambah prefiks (mis. `feature/`) dan pilihan source branch (default branch utama repository).
  5. Dialog menampilkan perintah git siap salin, mis. `git checkout -b PROJ-123-perbaiki-login`.
  6. Tombol "Create" membuat branch langsung melalui konektor (DEV-007 s.d. DEV-009); bila konektor tidak punya izin tulis, dialog menawarkan "Buka di <tool>" yang membuka halaman create branch tool asal dengan nama terisi.
  7. Branch yang dibuat langsung tampil di Development panel tanpa menunggu webhook, dan dapat memicu otomasi Branch created untuk auto-transition (DEV-026).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Action otomasi Create branch di AUTO-055/DEV-026.

### DEV-015 Development field in search and lists

- **Deskripsi:** Field yang dapat dicari `development[...]` dan kolom Development untuk menyaring serta melihat status development work item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, release manager; memerlukan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. JQL mendukung `development[pullrequests].all`, `development[pullrequests].open`, `development[commits].all`, `development[builds].all`, `development[builds].failing`, `development[builds].passing`, dan `development[reviews].all`/`.open` (legacy, selalu 0 karena code review terpisah tidak didukung) dengan operator pembanding numerik.
  2. ⚠ JQL untuk deployment didukung sebagai `development[deployments].all` (JQL-035 butir 5, bersama `development[branches].all`) dan, sebagai tambahan modul ini, `development[deployments].production` (jumlah deployment sukses ke environment `production`); kedua counter dibaca dari ringkasan DEV-001.
  3. Kolom Development di issue navigator dan List view menampilkan ringkasan teks (mis. "2 pull requests OPEN", "1 build failed") dengan ikon, dan klik sel membuka dialog Development details (DEV-013).
  4. Field dapat dipakai di saved filter, board filter, quick filter, dan kondisi JQL otomasi; nilai dibaca dari ringkasan DEV-001 sehingga hasil konsisten di semua tempat.
  5. ⚠ Untuk pengguna tanpa View Development Tools, kolom Development tidak tersedia di pemilih kolom dan klausa `development[...]` tidak mencocokkan work item mana pun.
  6. Grammar lengkap, operator, dan alias legacy didefinisikan di JQL-035 (modul 12).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Navigator dan List view di modul 13.

### DEV-016 Dev status on board cards

- **Deskripsi:** Indikator status development di kartu board dan backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Kartu menampilkan satu ikon dev status yang mewakili tahap terjauh: branch, commit, PR (dengan status), atau build (hijau sukses/merah gagal); kartu tanpa data dev tidak menampilkan ikon.
  2. Hover ikon menampilkan tooltip ringkasan (jumlah branch, commit, PR, build terakhir) dan klik membuka dialog Development details.
  3. ⚠ Company-managed: tampilan diatur di Board settings > Card layout dengan toggle "Development"; team-managed: diatur di Board > Card fields; default aktif.
  4. ⚠ Indikator deployment ditampilkan sebagai ikon tambahan dengan environment tertinggi yang dicapai (mis. "Production") bila ada deployment sukses.
  5. Ikon hanya dirender untuk pengguna dengan View Development Tools; pengguna lain melihat kartu tanpa ikon tanpa placeholder.
- **Team-managed vs company-managed:** Lokasi pengaturan berbeda (butir 3).
- **Catatan & ketergantungan:** Card layout di modul 10; kartu backlog di modul 11.

### DEV-017 Work suggestions

- **Deskripsi:** Panel saran kerja di board/sprint yang menyorot item yang butuh perhatian berdasarkan data development dan aktivitas work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan View Development Tools; saran hanya mencakup proyek yang dapat dilihat pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Kategori saran: PR yang menunggu review pengguna, PR milik pengguna dengan build gagal, PR terbuka lebih dari 7 hari, PR yang sudah disetujui tetapi belum di-merge, dan work item In Progress milik pengguna yang tidak diperbarui lebih dari 5 hari kerja.
  2. Saran dipersonalisasi per pengguna (berdasarkan assignee, author PR, dan reviewer) dan diurutkan dari yang paling mendesak.
  3. Klik saran membuka work item atau PR terkait (PR di tab baru tool asal).
  4. Setiap saran dapat di-dismiss; saran yang di-dismiss tidak muncul lagi kecuali kondisinya berubah (mis. build gagal lagi setelah sukses).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditampilkan di board/sprint aktif (modul 10, 11). Nama dan lokasi mengikuti produk referensi 2025-2026 yang belum dipastikan; ruang lingkup final adalah keputusan terbuka.

## Halaman proyek DevOps & panel penyedia tambahan

### DEV-018 Code page

- **Deskripsi:** Halaman "Code" di sidebar proyek software yang menampilkan repository terhubung dan aktivitas kode terkait proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Admin Proyek/Project Lead; memerlukan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar menampilkan repository yang memiliki branch/commit/PR dengan issue key proyek ini, ditambah repository yang dikaitkan eksplisit di Toolchain (DEV-006 butir 4).
  2. Setiap baris menampilkan nama repository, ikon tool asal, waktu update terakhir, jumlah work item tertaut, dan tautan ke repository di tool asal.
  3. Bila belum ada repository, halaman menampilkan onboarding "Connect your code" dengan pilihan GitHub, GitLab, Bitbucket, dan "Lainnya (API)" yang membuka dokumentasi DEV-004.
  4. Kotak pencarian menyaring repository berdasarkan nama (case-insensitive) dan filter tool asal; hasil kosong menampilkan "Tidak ada repository yang cocok".
  5. Tautan "Kelola di Toolchain" tampil bagi Admin Proyek dan membuka DEV-006.
  6. Pengguna tanpa View Development Tools tidak melihat item sidebar "Code", dan akses langsung ke URL menampilkan halaman 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Item sidebar proyek di modul 21. Diaktifkan oleh template DevOps (DEV-030).

### DEV-019 Deployments page

- **Deskripsi:** Halaman "Deployments" di sidebar proyek yang menampilkan timeline deployment per environment beserta work item yang dideploy.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, release manager, DevOps engineer, Stakeholder/Viewer dengan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Timeline dibagi menjadi lajur per environment type dengan urutan Production, Staging, Testing, Development, Unmapped; lajur tanpa deployment dalam rentang waktu tetap tampil kosong.
  2. Setiap deployment ditampilkan sebagai penanda berwarna sesuai state dengan informasi pipeline, nama environment, state, waktu, dan jumlah work item terkait.
  3. ⚠ Filter tersedia untuk environment type, issue type, rentang waktu (default 14 hari terakhir), pipeline, dan Fix version; filter tersimpan di URL agar dapat dibagikan.
  4. Klik deployment membuka panel samping berisi daftar work item yang dideploy beserta status dan assignee, dengan tautan ke work item.
  5. Tersedia mode tampilan daftar work item (tabel) yang menampilkan per work item environment terakhir yang dicapai.
  6. Bila belum ada tool CI/CD terhubung, halaman menampilkan onboarding dengan pilihan penyedia (GitHub Actions, GitLab, Bitbucket Pipelines, Jenkins, lainnya).
  7. Sumber data adalah Deployments API (DEV-004) dari semua penyedia (Bitbucket Pipelines, GitHub Actions, GitLab, Jenkins, CircleCI, Octopus Deploy, Azure Pipelines, dan lainnya) yang tampil seragam.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model environment di DEV-003; metrik di DEV-029.

### DEV-020 Security in Jira (vulnerabilities)

- **Deskripsi:** Halaman Security dan panel di work item untuk menampilkan vulnerability dari tool security terhubung dan menautkannya ke work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Security engineer, Developer/Anggota Tim; melihat memerlukan View Development Tools, membuat work item dari vulnerability memerlukan Create Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman proyek "Security" menampilkan daftar vulnerability dengan kolom severity (critical, high, medium, low, unknown), status (open, closed, ignored, unknown), container, tanggal diperkenalkan, dan work item tertaut.
  2. Filter dan pencarian tersedia berdasarkan severity, status, container, dan tool penyedia; filter dapat dikombinasikan.
  3. Aksi "Create issue" dari vulnerability membuka dialog create dengan summary dan description terisi dari judul, identifier, dan URL advisory; aksi "Link to existing" menautkan ke work item yang dipilih.
  4. Vulnerability yang tertaut ditampilkan di panel Security pada issue view dengan severity dan status terbaru.
  5. Pengaturan security container memungkinkan Admin Proyek memilih repo/project di tool security yang dikaitkan dengan proyek; hanya vulnerability dari container terpilih yang tampil.
  6. Penyedia mengirim data melalui Security Information API (DEV-004 butir 5), mis. Snyk, Mend, Checkmarx, JFrog, Sonatype, dan GitHub Advanced Security (DEV-007 butir 9); konektor penyedia selain GitHub termasuk ekosistem.
  7. Setiap vulnerability menampilkan identifier CVE, CWE, dan URL advisory bila dikirim, sebagai tautan.
  8. ⚠ Permission: melihat halaman dan panel memerlukan View Development Tools; mengubah pengaturan container memerlukan Administer Projects.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission.
- **Catatan & ketergantungan:** Trigger otomasi "Vulnerability found" dibahas di DEV-026.

### DEV-021 Feature flags panel

- **Deskripsi:** Panel di issue view yang menampilkan feature flag terkait work item beserta status rollout per environment.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Product Owner/PM dengan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Create feature flag" dan "Connect feature flag" mengarahkan pengguna ke halaman penyedia dengan issue key terisi sebagai parameter.
  2. Ringkasan panel menampilkan jumlah flag, status enabled/disabled, dan rollout di environment teratas (production bila ada).
  3. Detail per flag menampilkan key, display name, environment, status, rollout (percentage, text, atau rules), dan tautan ke penyedia.
  4. Penyedia mengirim data melalui Feature Flags API (DEV-004 butir 4), mis. LaunchDarkly, Split, Optimizely, Flagsmith, Unleash, dan GitLab feature flags (DEV-008).
  5. Panel hanya tampil bila minimal satu penyedia feature flag terdaftar di site.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konektor penyedia flag dibuat vendor (ekosistem); produk ini menyediakan panel dan API.

### DEV-022 Design integration (Figma)

- **Deskripsi:** Rujukan singkat. Penautan desain (mis. Figma) ke work item dengan preview dan status desain dispesifikasikan penuh di INT-048 Figma for Jira & design integrations ([25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md)), yang menjadi pemilik fitur ini. Modul ini hanya memiliki kontrak Design API (DEV-004 butir 8) dan penempatan panel Designs di permukaan DevOps.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Designer, Developer/Anggota Tim; menambah desain memerlukan Edit Issues (aturan akses desain privat di INT-048 butir 6).
- **Aturan perilaku (acceptance criteria):**
  1. Desain ditambahkan dengan menempel URL Figma di panel Designs issue view sesuai INT-048 butir 1; URL bukan desain yang didukung atau file tanpa akses ditolak dengan pesan "Desain tidak dapat ditautkan".
  2. Preview thumbnail, nama desain/frame, dan tautan "Buka di Figma" mengikuti INT-048 butir 2 dan 5.
  3. ⚠ Status desain (mis. "Ready for development") ditampilkan sebagai lozenge sesuai INT-048 butir 2-3.
  4. Sinkronisasi perubahan desain mengikuti INT-048 butir 3; bila penyedia tidak mengirim webhook, produk ini melakukan polling berkala maksimal 1 jam.
  5. ⚠ Halaman/tab Designs di level proyek tidak dibangun di Fase 4: sidebar proyek tidak menampilkan item "Designs", dan desain hanya terlihat dari panel Designs di issue view.
  6. Tool desain lain menyediakan data melalui Design API (DEV-004 butir 8, padanan `designInfoProvider` di INT-048 butir 7) sebagai modul ekstensi (modul 25).
  7. URL Figma di description/comment tetap dirender sebagai smart link (modul 07) sebagai fallback bila integrasi desain tidak terpasang (selaras INT-048 butir 8).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik: INT-048 di modul 25; bila ada perbedaan aturan, INT-048 berlaku kecuali untuk kontrak Design API (DEV-004) yang dimiliki modul ini. Tier diselaraskan menjadi platform mengikuti INT-048. Produk referensi memerlukan app "Figma for Jira"; produk ini membangun konektor Figma sendiri atau hanya menyediakan Design API (keputusan terbuka).

### DEV-023 Operations and incidents panel

- **Deskripsi:** Menampilkan incident dan post-incident review terkait dari tool operations pada work item dan mengorelasikannya dengan deployment.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** SRE, Developer/Anggota Tim dengan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Data incident dan post-incident review berasal dari Operations API (DEV-004 butir 6), baik dari Operations milik produk ini (modul 34) maupun penyedia eksternal (mis. PagerDuty).
  2. Incident dan post-incident review dapat ditautkan ke work item (via association atau aksi manual) dan tampil dengan judul, severity, status, dan tautan ke sumber.
  3. Halaman Deployments (DEV-019) menampilkan penanda incident pada timeline sehingga incident yang terjadi setelah deployment ke environment yang sama terlihat berdekatan (korelasi deploy-incident).
  4. ⚠ Lokasi tampilan di proyek software: panel "Operations" di issue view dan penanda di Deployments page; halaman proyek khusus incident tidak dibangun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Incident management penuh di modul 32; alert/on-call di modul 34. Kategori Operate di Toolchain (DEV-006 butir 5).

## Aksi dari kode: Smart commits, triggers & gating

### DEV-024 Smart commits

- **Alias inventaris:** Smart Commits & external transition commands
- **Deskripsi:** Perintah di pesan commit untuk menambah komentar, mencatat waktu, dan menjalankan transisi workflow pada work item atas nama pengguna yang email-nya cocok dengan committer.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (committer), Sistem (prosesor); eksekusi memakai permission pengguna yang cocok: Add Comments, Work On Issues, Transition Issues; toggle dikelola Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Sintaks umum: `<ISSUE-KEY> <teks diabaikan> #<command> <argumen opsional>`; teks di antara issue key dan perintah pertama diabaikan, dan perintah dikenali hanya bila diawali `#` setelah spasi atau awal baris.
  2. `#comment <teks>` menambah komentar berisi teks sampai perintah berikutnya atau akhir baris, atas nama pengguna yang email-nya cocok dengan email committer; komentar diberi penanda sumber "via commit <hash pendek>".
  3. `#time <n>w <n>d <n>h <n>m <komentar>` menambah worklog dengan durasi dan komentar tersebut; bila time tracking dinonaktifkan di site/proyek, perintah gagal dengan alasan "Time tracking tidak aktif" (aturan worklog lengkap di TIME-044).
  4. `#<nama-transisi>` menjalankan transisi workflow yang namanya cocok dengan spasi diganti tanda hubung (mis. `#start-progress` untuk "Start Progress"), dan dapat digabung dengan `#comment` dalam satu perintah sehingga komentar ditambahkan pada transisi yang sama.
  5. Satu pesan commit dapat memuat banyak issue key dan banyak perintah; setiap perintah diterapkan ke semua issue key yang mendahuluinya pada baris yang sama, dan perintah diproses berurutan dari kiri ke kanan.
  6. ⚠ Transisi yang memerlukan input field wajib di screen gagal dan dilaporkan; field Resolution tidak dapat diisi lewat smart commit, sehingga transisi ke status Done hanya mengisi resolution bila workflow memakai post function set resolution.
  7. Bila nama transisi cocok dengan lebih dari satu transisi yang tersedia dari status saat ini (ambigu), perintah transisi tidak dijalankan dan dilaporkan sebagai galat "Nama transisi ambigu".
  8. Pengguna dicocokkan lewat email committer terhadap email pengguna aktif; bila tidak ada kecocokan unik atau pengguna tidak punya Add Comments (untuk `#comment`), Work On Issues (untuk `#time`), atau Transition Issues (untuk transisi) pada work item, perintah terkait dilewati.
  9. Galat setiap perintah yang gagal dikirim melalui email ke alamat committer berisi hash commit, perintah, dan alasan (modul 19); bila email committer tidak cocok dengan pengguna mana pun, tidak ada email yang dikirim.
  10. Smart commits dapat diaktifkan/dinonaktifkan per koneksi integrasi dan per repository oleh Admin Situs; default aktif.
  11. Perintah hanya diproses dari pesan commit; issue key di nama branch atau judul PR hanya menautkan data (DEV-002) dan tidak menjalankan perintah.
  12. Sintaks rangkaian penuh `ISSUE-KEY #transition-name #comment <teks> #time <durasi>` dalam satu baris diterapkan sebagai satu kesatuan: komentar dan worklog tetap diterapkan meskipun transisi gagal, dan masing-masing dilaporkan terpisah.
  13. Pencocokan nama transisi bersifat case-insensitive dan memilih transisi pertama yang cocok secara unik di antara transisi yang tersedia dari status work item saat ini; prefiks nama yang unik (mis. `#start` untuk satu-satunya "Start Progress") diterima.
  14. Transisi dari smart commit selalu menghormati condition, validator, dan permission pengguna yang cocok (berbeda dengan Workflow Triggers DEV-025); pengguna yang tidak memenuhi condition tidak dapat menjalankan transisi.
  15. Transisi dijalankan tanpa menampilkan screen; validator field wajib yang tidak terpenuhi menyebabkan transisi gagal dan dicatat di log smart commit.
  16. ⚠ Kegagalan transisi dilaporkan ke committer via email dengan format yang sama dengan butir 9, dan juga tercatat di log integrasi yang dapat dibuka Admin Situs.
  17. Toggle smart commits di koneksi berlaku untuk ketiga jenis perintah sekaligus; menonaktifkan toggle tidak memengaruhi penautan data dev maupun Workflow Triggers.
- **Team-managed vs company-managed:** Tidak ada perbedaan; transisi mengikuti workflow proyek masing-masing.
- **Catatan & ketergantungan:** DEV-024 adalah pemilik parser smart commit (sintaks, pencocokan pengguna, urutan eksekusi, toggle, pelaporan galat) untuk ketiga jenis perintah; TIME-044 Smart Commits #time di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) hanya mengatur efek worklog dari perintah `#time` (pembuatan worklog dan penyesuaian Remaining Estimate). Eksekusi transisi memakai service transisi tunggal modul 09; komentar di modul 08. Proses idempoten per hash commit sehingga commit yang dikirim ulang (mis. setelah backfill) tidak menjalankan perintah dua kali.

### DEV-025 Workflow Triggers (DevOps)

- **Deskripsi:** Trigger pada transisi workflow yang otomatis menjalankan transisi ketika ada event dari development tool terhubung (GitHub, GitLab, Bitbucket).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (mengedit workflow company-managed), Admin Proyek (workflow team-managed), Sistem (eksekusi).
- **Aturan perilaku (acceptance criteria):**
  1. Event yang dapat dipilih: Branch created, Commit created, Pull request created, Pull request merged, Pull request declined, dan Pull request reopened.
  2. Event code review (Review started, approved, rejected, abandoned, closed, summarized) dari tool code review terpisah tidak tersedia; pemilih event hanya menampilkan event butir 1.
  3. Satu transisi dapat memiliki beberapa trigger; trigger mana pun yang terpenuhi menjalankan transisi.
  4. Source trigger dapat diatur "Semua development tool terhubung" atau integrasi tertentu (mis. hanya koneksi GitHub organisasi X).
  5. Transisi dari trigger mengabaikan condition, validator, dan permission pengguna; History mencatat aktor "Sistem (Workflow trigger)" beserta event sumbernya.
  6. Trigger hanya berlaku untuk work item yang key-nya disebut di branch/commit/PR terkait, dan hanya bila work item berada di status asal transisi; work item di status lain tidak berubah.
  7. Halaman diagnostik trigger di editor workflow menampilkan eksekusi terakhir dan alasan tidak berjalan (mis. "Issue key tidak ditemukan", "Status asal tidak cocok", "Integrasi tidak terhubung").
  8. Editor trigger menampilkan saran bahwa Automation (DEV-026) adalah alternatif yang lebih fleksibel, dengan tautan ke template rule DevOps.
  9. ⚠ Konfigurasi trigger tersedia di tab "Triggers" pada panel transisi editor workflow modul 09.
- **Team-managed vs company-managed:** Company-managed dikonfigurasi di editor workflow (dipakai lewat workflow scheme) oleh Admin Situs. ⚠ Di produk referensi workflow team-managed tidak menyediakan Workflow Triggers dan auto-transition dilakukan lewat otomasi (DEV-026); produk ini boleh menyediakan tab Triggers yang sama di editor workflow team-managed untuk Admin Proyek (keputusan terbuka).
- **Catatan & ketergantungan:** Tidak ada workflow trigger untuk event build/deployment; gunakan otomasi (DEV-026). Event Fisheye/Crucible bersifat Data Center sehingga tidak dibangun (lihat Di luar cakupan).

### DEV-026 DevOps automation triggers and actions

- **Deskripsi:** Trigger dan action otomasi yang bereaksi pada event development dan membuat objek di tool development.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, penulis rule otomasi; action memerlukan koneksi integrasi yang valid.
- **Aturan perilaku (acceptance criteria):**
  1. Trigger tersedia: Branch created, Commit created, Pull request created, Pull request merged, Pull request declined, Build successful, Build failed, Build status changed, Deployment successful, Deployment failed, dan Deployment status changed; definisi di AUTO-030 dan AUTO-031.
  2. ⚠ Trigger build dapat difilter berdasarkan nama build/pipeline dan branch; trigger deployment dapat difilter berdasarkan environment type dan nama environment.
  3. Action tersedia: Create branch in Bitbucket, Create branch in GitHub, dan Create branch in GitLab (AUTO-055).
  4. Smart value tersedia minimal: `{{pullRequest.title}}`, `{{pullRequest.url}}`, `{{branch.name}}`, `{{commit.hash}}`, `{{build.name}}`, `{{deployment.environment.type}}`, `{{deployment.state}}`.
  5. Galeri template menyediakan rule siap pakai: transisi ke In Progress saat branch dibuat, ke In Review saat PR dibuat, dan ke Done saat PR di-merge.
  6. ⚠ Trigger "Vulnerability found" terpicu saat vulnerability baru dengan severity minimal yang dipilih tertaut ke work item atau container proyek (DEV-020).
  7. Engine, eksekusi, audit log, dan batas otomasi mengikuti modul 20; modul ini hanya menyuplai event `devinfo.*`, `build.*`, `deployment.*`, `vulnerability.*`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template rule diaktifkan oleh template DevOps (DEV-030).

### DEV-027 Deployment gating with Jira Service Management

- **Deskripsi:** Menahan deployment ke environment tertentu sampai change request di Service Management produk ini disetujui.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change approver), DevOps engineer; konfigurasi memerlukan Admin Proyek service project.
- **Aturan perilaku (acceptance criteria):**
  1. Gating dikonfigurasi di service project dengan change management (memilih pipeline/environment yang di-gate dan request type change) dan di tool CI/CD (Bitbucket Pipelines, Jenkins, GitHub Actions, GitLab, CircleCI, Octopus).
  2. Deployment ke environment yang di-gate otomatis membuat change request berisi daftar work item, commit, dan tautan pipeline terkait.
  3. Pipeline menunggu gating status yang bernilai `allowed`, `prevented`, atau `awaiting`; status `awaiting` dipertahankan sampai keputusan atau batas waktu yang dikonfigurasi.
  4. Approver menyetujui atau menolak di service project; persetujuan mengubah status gating menjadi `allowed`, penolakan menjadi `prevented`.
  5. Endpoint `GET .../pipelines/{pipelineId}/environments/{environmentId}/deployments/{deploymentSequenceNumber}/gating-status` mengembalikan status gating beserta tautan change request; deployment tanpa gating mengembalikan HTTP 404.
  6. Risk assessment, change calendar, dan approval workflow didefinisikan di modul 32.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung DEV-003, DEV-004, dan modul 32 (Fase 7).

## Metrik DevOps & onboarding

### DEV-028 Cycle time report

- **Deskripsi:** Perhitungan metrik cycle time (waktu dari commit pertama hingga deploy ke production per work item) yang menjadi sumber data laporan Cycle time (RPT-019).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Engineering manager, Scrum Master/Team Lead; akses laporan mengikuti modul 16 dan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan tersedia di Reports proyek software yang memiliki code tool dan deployment tool terhubung; tanpa keduanya laporan menampilkan empty state dengan tautan ke Toolchain.
  2. Cycle time per work item = waktu deployment sukses pertama ke environment `production` dikurangi waktu commit pertama yang tertaut; median dan rata-rata dihitung per periode (minggu).
  3. ⚠ Filter rentang waktu tersedia (pilihan 4, 12, 26 minggu atau rentang kustom; default 12 minggu); filter tambahan repository dan issue type tersedia dan dapat dikombinasikan, dan semua filter tersimpan di URL.
  4. Daftar work item dengan cycle time tertinggi ditampilkan untuk investigasi, diurutkan menurun.
  5. Work item tanpa commit tertaut atau tanpa deployment sukses ke production dikeluarkan dari perhitungan dan jumlahnya ditampilkan sebagai "Tidak terhitung".
  6. ⚠ Breakdown per tahap (time to first PR, time in review, time to merge, time to deploy) dihitung dari stempel waktu PR dan deployment yang tersimpan (lihat RPT-019 butir 3).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tampilan chart dan akses di RPT-019 (modul 16); modul ini menetapkan rumus dan sumber data.

### DEV-029 Deployment frequency report

- **Deskripsi:** Perhitungan jumlah deployment ke production dalam periode tertentu yang menjadi sumber laporan Deployment frequency (RPT-020).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Engineering manager, DevOps engineer; akses laporan mengikuti modul 16 dan View Development Tools.
- **Aturan perilaku (acceptance criteria):**
  1. Grafik menampilkan jumlah deployment sukses per hari atau per minggu (dapat dipilih).
  2. Filter environment (default `production`) dan rentang tanggal tersedia.
  3. Rata-rata frekuensi deployment ditampilkan untuk rentang terpilih (mis. "3,2 deployment per minggu").
  4. Daftar deployment di bawah grafik menampilkan waktu, pipeline, environment, dan work item terkait.
  5. Sumber data adalah Deployments API (DEV-004); deployment dihitung unik per kombinasi pipeline, environment, dan `deploymentSequenceNumber`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tampilan di RPT-020 (modul 16).

### DEV-030 Open DevOps template and onboarding

- **Deskripsi:** Template proyek dan alur onboarding untuk tim DevOps yang menyiapkan board, halaman DevOps, Toolchain, dan rule otomasi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead (membuat proyek dari template memerlukan izin create project).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Galeri template proyek menyediakan template "DevOps" berbasis Scrum atau Kanban dengan Toolchain yang disarankan.
  2. Onboarding setelah proyek dibuat memandu menghubungkan GitHub/GitLab/Bitbucket, wiki/dokumentasi, dan Service Management/Operations milik produk ini; setiap langkah dapat dilewati.
  3. Template mengaktifkan halaman Code, Deployments, Security, dan Toolchain di sidebar proyek.
  4. Rule otomasi DevOps (DEV-026 butir 5) aktif otomatis dari template dan dapat dimatikan.
  5. Halaman DevOps tanpa data menampilkan empty state dan panduan di sidebar yang mengarah ke langkah onboarding yang belum selesai.
- **Team-managed vs company-managed:** Template tersedia untuk kedua tipe; pilihan tipe mengikuti pembuatan proyek modul 04.
- **Catatan & ketergantungan:** Template proyek di modul 04; Open DevOps di produk referensi adalah bundling fitur yang sudah ada.

## Alat pengembang & AI coding

### DEV-031 Atlassian for VS Code (IDE integration)

- **Deskripsi:** Rujukan singkat. Ekstensi IDE milik produk ini (padanan fungsional ekstensi "Atlassian for VS Code") dispesifikasikan penuh di INT-050 Atlassian for VS Code (IDE integration) ([25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md)), yang menjadi pemilik fitur ini.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim; aksi mengikuti permission pengguna di produk ini.
- **Aturan perilaku (acceptance criteria):**
  1. Melihat, mencari (JQL), membuat, dan mentransisi work item dari panel IDE mengikuti INT-050 butir 2.
  2. "Start work" membuat branch berisi issue key dan mentransisi work item ke In Progress mengikuti INT-050 butir 3; deteksi key pada branch memakai DEV-002.
  3. Integrasi PR dan pipeline untuk penyedia terhubung (GitHub, GitLab, Bitbucket) mengikuti INT-050 butir 5 dan membaca data DEV-001.
  4. Login memakai akun produk ini melalui OAuth 2.0 dengan PKCE (bukan akun Atlassian) mengikuti INT-050 butir 1 dan 6.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik: INT-050 di modul 25; bila ada perbedaan aturan, INT-050 berlaku. Modul ini hanya menyediakan data development (DEV-001) dan deteksi issue key (DEV-002) yang dibaca ekstensi.

### DEV-032 Rovo Dev coding from work items

- **Deskripsi:** Integrasi agen AI coding opsional (padanan Rovo Dev) yang memulai pekerjaan kode dari work item dan menghasilkan PR.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim; memerlukan Edit Issues, View Development Tools, dan akses tulis repository melalui konektor.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Mulai dengan agen AI" di issue view membuat sesi agen yang menghasilkan perubahan kode dan PR berdasarkan summary, description, dan acceptance criteria work item; pengguna meninjau rencana sebelum PR dibuat.
  2. PR hasil agen memuat issue key di judul dan branch sehingga tertaut kembali ke Development panel (DEV-002, DEV-013).
  3. ⚠ Fitur hanya aktif bila Admin Situs mengaktifkan AI (modul 27), lisensi/kuota agen tersedia, dan repository terhubung; bila salah satu tidak terpenuhi, aksi disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia model AI dan kebijakan data mengikuti modul 27; nama, akses, dan plan produk referensi belum dipastikan.

### DEV-033 ⚠ External AI coding agent integration

- **Deskripsi:** Konektor generik yang memungkinkan work item di-assign atau dikirim ke agen AI coding eksternal (padanan fungsional "Assign to coding agent" dan "Launch coding agent from Jira" di produk referensi sejak pertengahan 2026), lalu melacak status run dan menautkan branch/PR hasilnya ke work item. Integrasi dirancang vendor-agnostic melalui kontrak HTTP yang terdokumentasi, bukan konektor khusus per vendor.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (mendaftarkan dan mencabut koneksi coding agent); Developer/Anggota Tim (mengirim work item ke agen, memerlukan Browse Projects, Edit Issues, dan View Development Tools pada proyek, serta izin AI agent sesuai PERM-050); Admin Proyek/Project Lead (membuat rule otomasi yang memakai aksi "Send to coding agent").
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs dapat mendaftarkan koneksi coding agent di halaman admin Apps/Integrations dengan atribut: nama tampilan, endpoint HTTPS untuk memulai run, endpoint/URL callback status, metode autentikasi (bearer token, HMAC signature, atau OAuth 2.0 client credentials), dan daftar proyek yang boleh memakai koneksi; endpoint non-HTTPS ditolak dengan pesan "Endpoint coding agent harus memakai HTTPS".
  2. Secret koneksi disimpan terenkripsi melalui mekanisme koneksi dan secret otomasi (AUTO-056), tidak pernah ditampilkan ulang setelah disimpan (hanya dapat diganti), dan setiap panggilan keluar melewati SSRF & Outbound Request Guard (INT-001) sehingga endpoint yang me-resolve ke alamat privat, loopback, atau metadata cloud ditolak.
  3. Admin dapat menguji koneksi ("Test connection"); kegagalan menampilkan kode HTTP atau galat jaringan tanpa membocorkan secret, dan koneksi dapat dinonaktifkan atau dihapus kapan saja; menonaktifkan koneksi menyembunyikan aksi terkait dari issue view dan membuat rule otomasi yang memakainya gagal dengan status "Koneksi coding agent tidak aktif".
  4. Di issue view tersedia aksi "Kirim ke coding agent" (dan pilihan agen pada field assignee khusus agen) yang hanya tampil bila minimal satu koneksi aktif berlaku untuk proyek, AI diaktifkan Admin Situs (modul 27), dan pengguna memiliki izin pada butir aktor; bila tidak terpenuhi, aksi disembunyikan.
  5. Saat mengirim, produk ini menyusun paket konteks: issue key, summary, description, acceptance criteria (field atau bagian deskripsi), tipe dan prioritas, linked work items (key, summary, tipe tautan), komentar terbaru yang dipilih pengguna, serta repository target dan base branch dari Toolchain/konektor (DEV-006, DEV-007, DEV-008, DEV-009); pengguna dapat meninjau dan menambahkan instruksi sebelum mengirim.
  6. Paket konteks tidak memuat field yang tidak dapat dilihat pengguna pengirim (security level, field tersembunyi) maupun lampiran, kecuali pengguna secara eksplisit memilih lampiran; ukuran paket dibatasi (default 256 KB) dengan galat "Konteks terlalu besar untuk dikirim ke coding agent" bila terlampaui.
  7. Setiap pengiriman membuat entitas run dengan ID unik, agen, pengirim, waktu mulai, dan status (`queued`, `running`, `needs_input`, `succeeded`, `failed`, `cancelled`); agen memperbarui status melalui callback bertanda tangan atau produk ini melakukan polling bila agen tidak mendukung callback.
  8. Status run terbaru tampil di issue view (panel "Coding agent" atau bagian Development panel) dengan tautan ke sesi agen eksternal, ringkasan hasil, dan riwayat run sebelumnya; pengirim dapat membatalkan run yang belum selesai bila agen mendukung pembatalan.
  9. Callback dengan signature tidak valid, run ID tidak dikenal, atau transisi status mundur (mis. `succeeded` ke `running`) ditolak dan dicatat; callback duplikat diproses idempoten.
  10. Branch, commit, dan PR yang dibuat agen wajib memuat issue key sehingga tertaut otomatis ke Development panel (DEV-002, DEV-013) melalui konektor SCM yang sama; PR berlabel "dibuat oleh coding agent <nama>" dan tetap memerlukan review manusia sesuai aturan repository, produk ini tidak melakukan merge otomatis.
  11. Bila run selesai tetapi tidak ada PR/branch yang terdeteksi dalam batas waktu yang dapat dikonfigurasi, status run menampilkan peringatan "Tidak ada perubahan kode yang tertaut".
  12. Modul otomasi (modul 20) menyediakan aksi "Send to coding agent" dengan parameter koneksi, repository, dan instruksi tambahan (mendukung smart values), serta trigger/kondisi atas perubahan status run; aksi berjalan dengan identitas actor rule dan tetap dibatasi PERM-050.
  13. Setiap pendaftaran/perubahan koneksi, pengiriman, pembatalan, dan callback status tercatat di audit log (modul 22) dan di History work item dengan actor pengguna pengirim serta identitas agen.
  14. Admin dapat menetapkan kuota run per pengguna/proyek per hari; melebihi kuota menolak pengiriman dengan pesan yang menyebut batas yang berlaku.
- **Team-managed vs company-managed:** Tidak ada perbedaan; ketersediaan per proyek diatur oleh daftar proyek pada koneksi.
- **Catatan & ketergantungan:** Berbeda dari DEV-032 (agen AI coding bawaan, padanan Rovo Dev, yang berjalan di dalam produk ini): DEV-033 hanya mengorkestrasi agen pihak ketiga melalui kontrak generik dan tidak menjalankan model AI sendiri. Bergantung pada AUTO-056 (secret), INT-001 (SSRF guard), PERM-050 (AI Agent Permissions), AUTO-055 (aksi lintas produk/DevOps), konektor SCM DEV-007 s.d. DEV-009, dan kebijakan AI modul 27. Kontrak HTTP (payload paket konteks, skema callback, signature) didokumentasikan di modul 23. Fitur ini ditambahkan di luar inventaris berdasarkan rilis produk referensi Mei-Juli 2026; nama fitur, plan, dan batasan persisnya masih perlu dikonfirmasi. Integrasi spesifik per vendor dicatat di bagian Di luar cakupan.

## Pertanyaan terbuka

- Apakah produk ini membangun konektor native GitHub, GitLab, dan Bitbucket sendiri (GitHub App, OAuth app) atau pada rilis pertama hanya menyediakan DevOps REST API dan webhook generik?
- Apakah path kompatibel produk referensi (`/rest/devinfo/0.10/bulk`, `/rest/builds/0.1/bulk`, `/rest/deployments/0.1/bulk`) diekspos di samping `/api/v1/devops/...` agar plugin/skrip CI yang sudah ada dapat dipakai ulang dengan sedikit perubahan?
- Apakah Workflow Triggers (DEV-025) dibangun, atau cukup template otomasi DevOps (DEV-026) sebagai satu-satunya mekanisme auto-transition? Bila dibangun, apakah juga tersedia di workflow team-managed?
- Apakah kegagalan Smart commits dilaporkan hanya lewat email ke committer, atau juga lewat notifikasi in-app bila email cocok dengan pengguna?
- Berapa lama data dev info disimpan setelah integrasi dicabut: dihapus segera (DEV-001 butir 5) atau ditahan masa tenggang (mis. 30 hari) agar penyambungan ulang tidak memerlukan backfill?
- Apakah produk ini membuat plugin Jenkins, orb CircleCI, task Azure Pipelines, dan ekstensi VS Code sendiri, atau hanya dokumentasi API?
- Apakah panel Security, Feature flags, Designs, dan Operations dibangun di Fase 4 atau ditunda sampai ada penyedia/konektor nyata?
- Apakah ekstensi UI di sisi Bitbucket/GitHub (menampilkan work item di halaman PR) termasuk cakupan produk ini?
- Apakah agen AI coding (DEV-032) dibangun sendiri, atau hanya integrasi dengan agen pihak ketiga yang membuat PR berisi issue key?
- Untuk konektor coding agent eksternal (DEV-033), apakah agen dimodelkan sebagai pengguna khusus yang dapat menjadi assignee (mirip akun app) atau hanya sebagai aksi "Kirim ke coding agent" tanpa mengubah assignee? Apakah kontrak generiknya mengadopsi protokol terbuka yang sudah ada (mis. MCP/A2A) atau skema HTTP milik produk ini?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** seperti Git Integration for Jira, konektor Snyk/Mend/Checkmarx/JFrog/Sonatype, penyedia feature flag (LaunchDarkly, Split, Optimizely, Flagsmith, Unleash), app Azure Pipelines buatan Microsoft, dan app GitLab for Jira Cloud buatan GitLab tidak dibangun sebagai fitur native; produk ini hanya menyediakan API (DEV-004) yang dapat dipakai vendor.
- **App/plugin yang terikat ke produk referensi** (GitHub for Jira, plugin Jenkins "Atlassian Jira Software Cloud", Figma for Jira, Atlassian for VS Code) tidak dapat dipakai; produk ini membangun padanan sendiri bila diputuskan (DEV-007, DEV-010, DEV-022, DEV-031).
- **DVCS accounts (legacy Development tools)** (DEV-011) bertanda dc-only-or-deprecated dan tidak dibangun.
- **Fisheye/Crucible dan event code review terkait** (Review started/approved/rejected/abandoned/closed/summarized, `development[reviews]`) hanya relevan di Data Center melalui application link dan tidak dibangun.
- **Bitbucket Data Center via application link** tidak dibangun; dukungan Bitbucket Data Center, bila ada, memakai konektor token (DEV-009 butir 8).
- **Layanan lintas produk Atlassian** seperti Compass sebagai produk, Opsgenie, Rovo Dev sebagai layanan, akun Atlassian, dan Bitbucket sebagai produk tidak dibangun; produk ini hanya menyediakan padanan internal (Operations di modul 34, komponen di REL-035) atau konektor.
- **Integrasi coding agent spesifik per vendor** (mis. Cursor, Claude Code, GitHub Copilot coding agent, OpenAI Codex, Devin) tidak dibangun sebagai konektor khusus; produk ini hanya menyediakan konektor generik DEV-033 yang dapat dipenuhi vendor atau adaptor buatan pelanggan.
