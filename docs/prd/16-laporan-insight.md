# 16 — Laporan & Insight

Modul ini mendefinisikan seluruh kemampuan pelaporan dan insight di produk ini: halaman pusat Reports per project/board (company-managed dan team-managed), mesin kalkulasi historis yang merekonstruksi status, sprint, estimasi, dan assignment work item pada titik waktu tertentu dari change history, input kalkulasi dari konfigurasi board (estimation statistic, working days, kolom, swimlane, quick filter), aturan akses dan empty state laporan, laporan agile Scrum (Burndown, Burnup, Sprint report, Velocity), laporan epic dan rilis (Epic report, Epic burndown, Version report, Release burndown), laporan alur (Cumulative flow diagram, Control chart), laporan DevOps (Cycle time, Deployment frequency), laporan Issue analysis dan Forecast & management berbasis project atau saved filter, panel Insights kontekstual di board/backlog, halaman Summary project, berbagi/ekspor/cetak laporan, titik ekstensi report dari app, analitik lanjutan setara paket Enterprise (query builder, dashboard analitik, data lake), serta insight berbantuan AI yang opsional. Gadget dashboard yang memakai kalkulasi yang sama didefinisikan di modul 17; modul ini menjadi pemilik logika kalkulasinya. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, middleware bahasa EN/ID), sehingga seluruh laporan adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/report` untuk endpoint laporan dan `internal/features/insight` untuk panel Insights) mengikuti pola Clean Architecture repo, dengan halaman `/projects/:key/reports/...` di SPA `frontend/`. Prasyarat yang tidak bisa disusulkan belakangan: changelog lengkap setiap perubahan field work item dan event sprint wajib disimpan sejak Fase 1 (modul 08 dan 28), karena tanpa riwayat itu laporan historis tidak dapat direkonstruksi.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi "work item", plan tier, feature flag.
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects (PERM-020), team-managed access levels (PERM-038), Issue Security Enforcement (PERM-044), Anonymous & Public Access (PERM-046), Board & Filter Permissions (PERM-051).
- [04 Proyek](04-proyek.md) — navigasi project (PRJ-027), Project Summary sebagai view (PRJ-028), Features team-managed (PRJ-049) dan company-managed (PRJ-050).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field Story Points, Parent, custom field number/select/date, screen.
- [06 Work Item Inti](06-work-item-inti.md) — Flag (ISS-023), Issue Links (ISS-029), Move (ISS-032), Delete (ISS-033), Archive (ISS-034), Agents in Jira (ISS-046).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — history/changelog sebagai sumber data laporan.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category, Resolution.
- [10 Board](10-board-kolom-kartu.md) — board filter (BRD-010), column status mapping (BRD-025), Done column (BRD-030), swimlane (BRD-031), quick filter (BRD-033), board settings (BRD-012), Working days (BRD-040), Estimation & Tracking (BRD-041), Kanban sub-filter (BRD-046).
- [11 Backlog & Sprint](11-backlog-sprint.md) — Sprint Start (SPR-018), Sprint Scope Change (SPR-019), Sprint Complete (SPR-020), Parallel Sprints (SPR-023), Sprint Linked Pages (SPR-026).
- [12 JQL](12-jql-bahasa-kueri.md) — klausa JQL yang dihasilkan drill-down laporan.
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — Issue Navigator, saved filter, ekspor CSV/Excel.
- [14 Pelacakan Waktu & Estimasi](14-pelacakan-waktu-estimasi.md) — aktivasi dan setting time tracking (TIME-001, TIME-002), Board Estimation Statistic (TIME-022), Board Time Tracking Statistic (TIME-023), sumber data estimasi untuk laporan; TIME-027 s.d. TIME-034 adalah rujukan ke laporan yang dimiliki modul ini (RPT-009 s.d. RPT-016, RPT-029 s.d. RPT-032).
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — version, Fix version, halaman Releases.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget yang memakai kalkulasi laporan.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — view Summary/List project, Plans yang memakai velocity.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — sidebar project, zona waktu dan locale profil.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — pengaturan global time tracking dan tampilan.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — endpoint data laporan.
- [24 Integrasi DevOps](24-integrasi-devops.md) — data development dan deployment untuk laporan DevOps.
- [25 Integrasi & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — titik ekstensi app, embed ke wiki.
- [26 Kerja Bisnis](26-kerja-bisnis-formulir-persetujuan.md) — project bisnis yang memakai Summary.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — background job, cache, batas ukuran query, aksesibilitas.
- [28 Model Data Inti](28-model-data-inti.md) — changelog, event domain `issue.*`/`sprint.*`.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) — Project summary page service space (JSM-PRJ-014) yang memakai Space Insights.
- [31 JSM Queues, SLA & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) — laporan Service Management memakai komponen chart yang sama.

## Persona terkait

- **Scrum Master/Team Lead** — memantau sprint lewat Burndown, Sprint report, Velocity, dan panel Insights; memakai Control chart dan CFD untuk perbaikan proses.
- **Developer/Anggota Tim** — melihat progres sprint dan Issues for attention di board, membuka drill-down ke work item.
- **Product Owner/PM** — memprediksi penyelesaian epic dan version (Epic report, Version report, Release burndown), memakai velocity untuk perencanaan kapasitas.
- **Admin Proyek/Project Lead** — mengaktifkan fitur Reports (team-managed), mengonfigurasi estimation statistic dan working days board, membaca laporan workload.
- **Stakeholder/Viewer** — membuka tautan laporan dan halaman Summary yang dibagikan tanpa bisa mengubah konfigurasi.
- **Pengguna Tim Bisnis** — memakai halaman Summary project tanpa board/sprint.
- **Admin Situs** — mengatur time tracking global, akses anonim, analitik lanjutan, data lake, dan kontrol AI.
- **Integrator/Developer API** — mengambil data laporan lewat API, membangun report app lewat titik ekstensi.
- **Tamu/Kolaborator Eksternal** — melihat laporan project publik bila anonymous browse diizinkan.

## User story modul

- Sebagai Scrum Master/Team Lead, saya ingin melihat Burndown sprint aktif dengan guideline dan tabel event agar saya tahu sejak dini apakah sprint akan selesai tepat waktu (RPT-009, RPT-005).
- Sebagai Scrum Master/Team Lead, saya ingin Sprint report yang memisahkan work item selesai, tidak selesai, dan dihapus serta menandai penambahan scope agar sprint review dan retrospective berbasis fakta (RPT-010).
- Sebagai Product Owner/PM, saya ingin Velocity chart Commitment vs Completed tujuh sprint terakhir agar komitmen sprint berikutnya realistis (RPT-012, RPT-033).
- Sebagai Product Owner/PM, saya ingin Version report dengan predicted, optimistic, dan pessimistic release date agar saya dapat mengomunikasikan tanggal rilis ke stakeholder (RPT-015, RPT-016).
- Sebagai Product Owner/PM, saya ingin Epic report dan Epic burndown agar saya tahu berapa sprint lagi sebuah epic selesai (RPT-013, RPT-014).
- Sebagai Scrum Master/Team Lead, saya ingin Cumulative flow diagram dan Control chart yang dapat disaring per kolom, swimlane, dan quick filter agar bottleneck dan variasi cycle time terlihat (RPT-017, RPT-018, RPT-007).
- Sebagai Developer/Anggota Tim, saya ingin panel Insights di board yang menyorot work item flagged, blocked, dan prioritas tinggi agar saya tahu apa yang perlu segera ditangani (RPT-033).
- Sebagai Admin Proyek/Project Lead, saya ingin Time tracking report dan User workload report agar saya bisa menilai akurasi estimasi dan beban per orang (RPT-029, RPT-030, RPT-031).
- Sebagai Stakeholder/Viewer, saya ingin membuka halaman Summary project dan tautan laporan yang dibagikan agar saya memahami kesehatan project tanpa konfigurasi (RPT-034, RPT-036).
- Sebagai Admin Proyek/Project Lead, saya ingin laporan Created vs resolved dan Pie chart dari saved filter agar saya bisa melihat apakah backlog bertambah dan bagaimana sebarannya (RPT-021, RPT-022, RPT-023).
- Sebagai Stakeholder/Viewer, saya ingin laporan tidak pernah menampilkan work item yang dilindungi issue security agar data rahasia tidak bocor lewat angka total (RPT-006, RPT-004).
- Sebagai Admin Proyek/Project Lead, saya ingin pesan yang menjelaskan kenapa laporan kosong beserta tombol perbaikannya agar tim tidak bingung saat belum ada sprint atau version (RPT-003).
- Sebagai Integrator/Developer API, saya ingin menambahkan report kustom dari app ke halaman Reports agar kebutuhan metrik khusus dapat dipenuhi tanpa mengubah inti produk (RPT-037).
- Sebagai Admin Proyek/Project Lead, saya ingin menyusun layout Space Insights berisi chart, list, dan activity stream agar anggota space melihat metrik yang relevan tanpa membuat dashboard terpisah (RPT-043).
- Sebagai Admin Situs, saya ingin melihat Agent usage dashboard berisi pengguna aktif, jumlah sesi, dan cycle time work item yang dibantu agent agar dampak AI dapat dinilai (RPT-044).
- Sebagai Admin Situs, saya ingin membangun dashboard analitik lintas project dengan query builder di atas data lake yang menghormati permission agar eksekutif mendapat metrik organisasi (RPT-038, RPT-039, RPT-040, RPT-041).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 6 |
| P1 | 9 |
| P2 | 11 |
| P3 | 18 |
| **Total** | **44** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 6 |
| Fase 2 | 9 |
| Fase 3 | 17 |
| Fase 4 | 11 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 membangun fondasi yang tidak bisa diganti belakangan, yaitu mesin kalkulasi historis di atas changelog dan snapshot komitmen sprint (RPT-004), input kalkulasi dari board (RPT-005), dan filter permission di dalam query laporan (RPT-006), lalu di atasnya halaman Reports hub serta dua laporan Scrum paling kritis, Burndown dan Sprint report (RPT-001, RPT-009, RPT-010). Fase 2 melengkapi laporan agile dan visibilitas tim: Burnup, Velocity, CFD, Control chart, panel Refine report, empty state, perilaku chart yang aksesibel, halaman Reports team-managed, dan halaman Summary project. Fase 3 menambah laporan berbasis epic/version, kerangka Issue analysis berbasis project atau saved filter beserta tujuh laporannya, laporan Forecast & management berbasis time tracking, panel Insights, serta berbagi/cetak laporan. Fase 4 berisi laporan DevOps (bergantung integrasi modul 24), Workload pie chart report sebagai varian berbasis waktu dari Pie chart report, titik ekstensi report app, analitik lanjutan setara Enterprise, Space Insights yang dapat dikustomisasi, Agent usage dashboard, dan insight AI opsional. Hanya report JWM lama yang dicatat sebagai rujukan dan tidak dibangun. Modul ini adalah pemilik tunggal laporan agile dan laporan berbasis time tracking (RPT-009 s.d. RPT-016, RPT-029 s.d. RPT-032); modul 14 hanya mendefinisikan sumber data estimasi dan time tracking (TIME-027 s.d. TIME-034 merujuk ke sini), sehingga prioritas, fase, dan tier yang berlaku adalah yang tertulis di modul ini. Setiap laporan memakai satu komponen chart dan satu kontrak API kalkulasi yang sama, sehingga gadget dashboard (modul 17) dan laporan JSM (modul 31) cukup menggunakan ulang lapisan ini.

## Fitur

## Pusat laporan & navigasi

### RPT-001 Reports hub (project Reports page)

- **Alias inventaris:** Kanban board reports list
- **Deskripsi:** Halaman pusat di sidebar project yang menampilkan semua laporan untuk project/board aktif, dikelompokkan per kategori (Agile, Issue analysis, Forecast & management, DevOps, Other). Menjadi titik masuk utama ke seluruh laporan, termasuk daftar laporan yang tersedia dari board Kanban.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Anggota project, Scrum Master/Team Lead, Admin Proyek; membutuhkan Browse Projects pada project dan akses ke board terpilih (RPT-006).
- **Aturan perilaku (acceptance criteria):**
  1. Item sidebar "Reports" tampil di navigasi project (PRJ-027); pada company-managed, daftar laporan agile bergantung pada board terpilih karena laporan agile dihitung per board, bukan per project.
  2. ⚠ Kategori company-managed: "Agile" (Burndown, Burnup, Sprint report, Velocity, Cumulative flow, Control chart, Epic report, Epic burndown, Version report, Release burndown), "Issue analysis" (Average age, Created vs resolved, Pie chart, Recently created, Resolution time, Single level group by, Time since), "Forecast & management" (Time tracking, User workload, Version workload), "DevOps" (Cycle time, Deployment frequency), dan "Other" untuk laporan dari app (RPT-037); nama kategori dapat dilokalkan EN/ID.
  3. ⚠ Daftar laporan agile difilter menurut tipe board: board Scrum menampilkan Burndown, Burnup, Sprint report, Velocity, Epic report, Epic burndown, Version report, Release burndown, Control chart, dan CFD; board Kanban menampilkan Cumulative flow dan Control chart; Version report tidak ditampilkan di board Kanban sampai keputusan terbuka diputuskan.
  4. Kartu laporan berisi nama, thumbnail ilustrasi, dan deskripsi singkat kegunaan; seluruh kartu dapat diklik dan dapat difokuskan dengan keyboard.
  5. Header setiap laporan memiliki dropdown "Switch report" untuk berpindah ke laporan lain pada board/project yang sama tanpa kembali ke hub; parameter yang relevan (sprint, rentang waktu) dipertahankan bila laporan tujuan mendukungnya.
  6. Bila project memiliki lebih dari satu board, board selector tampil di header; laporan agile dihitung ulang mengikuti board terpilih dan pilihan board terakhir diingat per pengguna per project.
  7. Laporan Issue analysis dan Forecast & management membuka form konfigurasi parameter terlebih dahulu (RPT-021) sebelum menampilkan hasil.
  8. Setiap laporan memiliki tautan "Cara membaca laporan ini" yang membuka panduan bantuan internal produk ini.
  9. Bila board belum punya sprint, version, atau fitur yang dibutuhkan belum aktif, kartu laporan tetap tampil tetapi laporan membuka empty state yang menjelaskan prasyarat (RPT-003).
  10. URL laporan dapat di-bookmark dan dibagikan, memuat board ID atau project key, jenis laporan, dan parameter di query string (mis. `/projects/ABC/reports/burndown?board=12&sprint=34`); membuka URL dengan parameter tidak valid menampilkan laporan dengan parameter default dan pesan "Parameter tidak valid, menampilkan default".
  11. Pada board Kanban, kategori Agile hanya berisi Cumulative flow diagram dan Control chart.
  12. Pada board Kanban, kategori Issue analysis berisi Average age, Created vs resolved, Pie chart, Recently created, Resolution time, Single level group by, dan Time since issues, memakai board filter sebagai sumber data default.
  13. ⚠ Pada board Kanban, kategori Forecast & management berisi Time tracking, User workload, dan Version workload, tampil hanya bila time tracking aktif.
  14. Laporan khusus Scrum (Burndown, Burnup, Sprint report, Velocity, Epic burndown, Release burndown) tidak tampil di board Kanban; membuka URL-nya langsung menampilkan pesan "Laporan ini hanya tersedia untuk board Scrum".
  15. ⚠ Version report di board Kanban: tidak ditampilkan pada rilis awal; bila keputusan terbuka memutuskan menampilkannya, perhitungan memakai populasi board filter Kanban.
- **Team-managed vs company-managed:** Company-managed: laporan terikat board dan board selector. Team-managed: satu board per project, halaman Reports diaktifkan lewat Features dan berisi daftar laporan yang lebih sedikit (RPT-002).
- **Catatan & ketergantungan:** Menggunakan ulang komponen chart bersama (RPT-008). Laporan app di kategori "Other" bergantung pada RPT-037. Route SPA `/projects/:key/reports` dan endpoint `GET /api/v1/projects/:key/reports` (katalog laporan beserta ketersediaannya).

### RPT-002 Team-managed Reports page & Overview

- **Deskripsi:** Halaman Reports untuk project team-managed yang hanya muncul bila fitur Reports aktif, dengan halaman Overview berisi kartu ringkas laporan. Daftar laporannya lebih sedikit dibanding company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (mengaktifkan fitur), anggota project dan Viewer (melihat); akses mengikuti access level project team-managed (PERM-038).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle "Reports" di Project settings > Features (PRJ-049); bila dimatikan, item sidebar Reports hilang dan URL laporan mengembalikan halaman "Fitur Reports tidak aktif di project ini" dengan tombol aktivasi untuk Admin Proyek.
  2. Burndown, Burnup, Sprint report, dan Velocity hanya muncul bila fitur "Sprints" aktif; Cumulative flow tersedia dengan atau tanpa sprint.
  3. Halaman Overview menampilkan grid kartu laporan dengan preview mini chart (data terbaru) dan tautan membuka laporan lengkap.
  4. Laporan yang tersedia: Burnup, Burndown, Sprint report, Velocity report, Cumulative flow diagram, serta Cycle time dan Deployment frequency bila Development/Deployments terhubung (modul 24).
  5. ⚠ Laporan Issue analysis, Forecast & management, Control chart, Epic report, Epic burndown, Version report, dan Release burndown tidak tersedia di team-managed; URL-nya mengembalikan pesan "Laporan tidak tersedia untuk project team-managed".
  6. Estimasi mengikuti setting Estimation di Features (Story points atau Time/Original estimate); Issue count tersedia sebagai alternatif di pemilih statistik laporan.
  7. ⚠ Filter laporan berupa sprint picker, rentang tanggal, dan filter work item (assignee, work type, label, epic/parent); filter tidak mengubah board.
  8. Status mapping diambil langsung dari kolom board team-managed; kolom paling kanan dianggap selesai.
  9. Semua pengguna yang memiliki akses ke project (role Viewer ke atas) dapat melihat laporan; tidak ada izin khusus laporan.
- **Team-managed vs company-managed:** Fitur ini khusus team-managed; tidak ada board filter JQL kustom, quick filter, atau swimlane sebagai parameter laporan.
- **Catatan & ketergantungan:** Bergantung pada Features team-managed (PRJ-049) dan RPT-007 untuk kontrol filter versi team-managed.

### RPT-003 Report empty states & prerequisites

- **Deskripsi:** Penanganan kondisi ketika laporan belum bisa ditampilkan karena data atau konfigurasi kurang, beserta panduan perbaikan yang dapat ditindaklanjuti.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Anggota project (melihat pesan), Admin Proyek/board admin (melihat CTA perbaikan); CTA yang mengubah konfigurasi hanya tampil bagi pengguna dengan izin yang sesuai.
- **Aturan perilaku (acceptance criteria):**
  1. Velocity chart dan Sprint report menampilkan pesan "Belum ada sprint yang ditutup di board ini" bila tidak ada sprint closed; Sprint report tetap bisa dibuka untuk sprint aktif.
  2. Burndown menampilkan pesan "Tidak ada sprint aktif" bila tidak ada sprint aktif dan tidak ada sprint terpilih, serta pesan "Work item di sprint ini belum diestimasi" bila total estimasi nol dengan statistik selain Issue count.
  3. Version report dan Release burndown menampilkan pesan "Project ini belum memiliki version" beserta CTA "Buat version" untuk pengguna dengan Administer Projects.
  4. Epic report dan Epic burndown menampilkan pesan "Tidak ada epic dalam board ini" beserta tautan membuat epic.
  5. Time tracking report, User workload, dan Version workload menampilkan pesan "Time tracking tidak aktif" dengan CTA ke pengaturan time tracking untuk Admin Situs.
  6. Laporan DevOps menampilkan CTA "Hubungkan" yang menuju halaman integrasi development tool atau pipeline (modul 24) bila belum ada data.
  7. Team-managed menampilkan CTA mengaktifkan fitur Sprints atau Reports bagi Admin Proyek; pengguna lain melihat pesan "Minta admin project mengaktifkan fitur ini".
  8. Laporan berbasis estimasi menampilkan banner peringatan bila lebih dari ambang (default 20%) work item dalam populasi belum diestimasi, lengkap dengan jumlah dan tautan ke daftar work item tersebut.
- **Team-managed vs company-managed:** Butir 7 hanya team-managed; selebihnya berlaku keduanya.
- **Catatan & ketergantungan:** Teks pesan disimpan di katalog pesan EN/ID (`internal/shared/constants`). Ambang peringatan estimasi dapat menjadi keputusan terbuka. Empty state minimal untuk Burndown dan Sprint report (butir 1–2) ikut dibangun di Fase 1 bersama RPT-009/RPT-010; sisanya di Fase 2.

## Fondasi kalkulasi & kontrol laporan

### RPT-004 Report historical calculation engine

- **Deskripsi:** Kemampuan inti untuk merekonstruksi status, sprint, estimasi, version, parent, dan assignment work item pada titik waktu tertentu dari change history, sebagai dasar semua laporan berbasis waktu. Ini kebutuhan implementasi, bukan halaman UI.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Sistem; dipanggil oleh semua endpoint laporan, gadget, dan panel Insights dengan konteks pengguna yang meminta.
- **Aturan perilaku (acceptance criteria):**
  1. Mesin membaca changelog untuk field status, resolution, sprint, story points (dan custom number field estimasi), original estimate, remaining estimate, fix version, parent, dan assignee, beserta timestamp dan aktor setiap perubahan.
  2. Saat sprint dimulai, snapshot komitmen sprint (daftar work item, estimasi, status) disimpan secara atomik bersama start sprint (SPR-018) dan tidak berubah oleh edit berikutnya; untuk data impor tanpa snapshot, komitmen direkonstruksi dari changelog pada `startDate`.
  3. Mapping status → kolom board → status category dihitung memakai konfigurasi kolom board saat ini untuk seluruh periode historis (retroaktif); work item dengan status tidak terpetakan diabaikan dari laporan berbasis kolom.
  4. Work item yang dipindah antar project (key berubah) tetap dihitung sebagai satu entitas berdasarkan ID internal dan tampil dengan key terbaru; work item yang keluar dari board filter dihitung hanya selama berada di dalam populasi board pada tanggal tersebut.
  5. ⚠ Work item yang dihapus hilang dari seluruh laporan termasuk data historis; work item yang diarsipkan tetap dihitung di laporan historis tetapi ditandai "Archived" pada daftar dan tidak dapat dibuka oleh pengguna tanpa izin melihat arsip.
  6. Reopen, perubahan resolution, dan transisi mundur ke kolom sebelumnya diperhitungkan: work item yang keluar dari kolom paling kanan dicatat sebagai "reopened" dan menambah sisa pekerjaan.
  7. Agregasi per periode (jam, hari, minggu, bulan, kuartal, tahun) memakai zona waktu board untuk laporan agile dan zona waktu profil pengguna untuk laporan Issue analysis.
  8. Setiap kueri laporan memfilter Browse Projects dan issue security level di dalam query (bukan pasca-filter di aplikasi) sehingga total, rata-rata, dan jumlah tidak memasukkan work item tersembunyi.
  9. ⚠ Laporan dibatasi maksimal 5.000 work item per populasi (konfigurabel di modul 27); melebihi batas menampilkan pesan "Data terlalu besar, persempit filter atau rentang waktu" tanpa hasil parsial yang menyesatkan.
  10. Hasil kalkulasi di-cache per (laporan, board/filter, parameter, versi permission pengguna) dengan TTL default 5 menit dan invalidasi oleh event `issue.*`/`sprint.*` yang relevan; laporan menampilkan "Diperbarui <waktu>" dan tombol "Refresh".
- **Team-managed vs company-managed:** Tidak ada perbedaan pada mesin; sumber mapping kolom berbeda (konfigurasi board vs kolom team-managed).
- **Catatan & ketergantungan:** Bergantung pada changelog dan event domain modul 28 serta event scope change SPR-019. Background job dan target latensi di modul 27. Catatan migrasi: tabel changelog belum ada di repo; harus dibuat bersama migrasi work item pertama agar data historis tidak hilang.

### RPT-005 Estimation statistic & working days for reports

- **Deskripsi:** Pengaturan board yang memengaruhi perhitungan laporan: estimation statistic, time tracking, working days, dan time zone. Modul 10 memiliki UI pengaturannya; modul ini menetapkan bagaimana nilainya dipakai dalam kalkulasi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board admin dan Admin Proyek mengubah pengaturan (BRD-011, BRD-012, BRD-040, BRD-041); semua pengguna laporan terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. Board configuration > Estimation menyediakan Story points, Original time estimate, Issue count, atau satu custom number field; laporan agile memakai pilihan ini sebagai default statistik.
  2. Opsi time tracking board: None atau "Remaining estimate and time spent"; bila yang kedua aktif, garis Burndown mengikuti Remaining estimate dan worklog menurunkan sisa pekerjaan.
  3. Board configuration > Working days menyimpan time zone board, hari kerja mingguan (default Senin–Jumat), dan daftar non-working days spesifik (tanggal libur).
  4. Non-working days dibayangi abu-abu di Burndown dan Burnup, dan garis guideline datar pada hari tersebut.
  5. Perubahan nilai estimasi pada work item setelah sprint dimulai dicatat sebagai scope change (SPR-019) dan tampil sebagai event "Estimate changed" di laporan.
  6. ⚠ Team-managed: estimasi diatur di Features (Story points atau Time); working days memakai default site (Senin–Jumat, zona waktu site) dan tidak dapat diubah per board.
  7. Field estimasi harus ada di screen work type agar bisa diisi; bila field tidak ada di screen, halaman pengaturan Estimation menampilkan peringatan "Field estimasi belum ada di screen untuk work type: <daftar>".
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** UI pengaturan di modul 10 (BRD-040 Working days, BRD-041 Estimation & Tracking); aturan estimation statistic di modul 14 (TIME-022, TIME-023); screen di modul 05.

### RPT-006 Report permissions & visibility

- **Deskripsi:** Aturan akses yang menentukan siapa dapat melihat laporan dan data apa yang tampil di dalamnya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek, anggota project; dikontrol oleh Browse Projects, share permission board filter, issue security level, dan access level team-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Melihat laporan project membutuhkan Browse Projects pada project terkait (PERM-020); tanpa izin, halaman menampilkan pesan akses ditolak dan API mengembalikan 403.
  2. Laporan board membutuhkan akses ke board, mengikuti share permission board filter (PERM-051); pengguna yang tidak dapat melihat board filter mendapat pesan "Anda tidak memiliki akses ke board ini".
  3. Work item dengan issue security level yang tidak boleh dilihat pengguna disembunyikan dari chart, daftar, dan total (PERM-044); jumlah tersembunyi tidak diungkapkan.
  4. Saved filter yang tidak dibagikan ke pengguna tidak dapat dipakai sebagai sumber laporan; membuka laporan dengan filter tersebut menampilkan "Filter tidak ditemukan atau tidak dibagikan kepada Anda".
  5. Bila anonymous browse diizinkan pada project (PERM-046), laporan dapat dilihat tanpa login dengan data yang dibatasi sama seperti issue view anonim.
  6. Team-managed: access level project (Open/Limited/Private) menentukan akses laporan (PERM-038); Limited dan Open memberi akses lihat laporan ke semua pengguna yang dapat melihat project.
  7. ⚠ Tidak ada permission khusus "View reports"; akses laporan sepenuhnya diturunkan dari Browse Projects dan akses board.
- **Team-managed vs company-managed:** Butir 6 untuk team-managed; butir 2 hanya relevan untuk board company-managed.
- **Catatan & ketergantungan:** Penegakan dilakukan di RPT-004 (filter di dalam query). Perilaku UI akses-aware di PERM-056.

### RPT-007 Report refinement controls (board-based)

- **Deskripsi:** Kontrol penyaringan laporan agile company-managed yang bersumber dari konfigurasi board: kolom, swimlane, quick filter, dan rentang waktu.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, board admin, anggota project; memilih kontrol tidak membutuhkan izin edit board.
- **Aturan perilaku (acceptance criteria):**
  1. Panel "Refine report" menyediakan checkbox Columns, Swimlanes, dan Quick filters; setidaknya satu kolom harus tercentang, bila tidak tombol terapkan nonaktif dengan pesan "Pilih minimal satu kolom".
  2. Rentang waktu preset (Past week, 2 weeks, month, 3 months, 6 months, All time) dan custom date range; tanggal akhir tidak boleh sebelum tanggal awal.
  3. Daftar quick filter dan swimlane identik dengan definisi di Board configuration (BRD-031, BRD-033); query-nya digabung dengan AND ke board filter.
  4. Board filter (JQL) menentukan populasi work item seluruh laporan board (BRD-010).
  5. ⚠ Kanban sub-filter (BRD-046) memengaruhi laporan board Kanban: work item yang dikeluarkan sub-filter (mis. sudah dirilis) tidak dihitung pada tanggal setelah keluar dari sub-filter; sebaliknya opsi tampilan "hide completed", pencarian board, dan quick filter yang sedang aktif di board view tidak memengaruhi laporan.
  6. Status mapping ke kolom menentukan kapan work item dianggap selesai, yaitu saat masuk kolom paling kanan (BRD-025, BRD-030).
  7. Perubahan konfigurasi board (kolom, mapping, filter) berlaku retroaktif pada seluruh laporan historis.
  8. ⚠ Pilihan refine disimpan di query string URL sehingga dapat dibagikan, dan pilihan terakhir diingat per pengguna per board selama sesi.
- **Team-managed vs company-managed:** Team-managed tidak memakai panel ini; filter yang tersedia adalah assignee, work type, label, epic/parent, dan rentang tanggal (RPT-002 butir 7).
- **Catatan & ketergantungan:** Dipakai oleh CFD (RPT-017) dan Control chart (RPT-018).

### RPT-008 Chart interaction & accessibility

- **Deskripsi:** Perilaku interaktif umum pada seluruh chart laporan agar mudah dibaca dan dapat diakses, diimplementasikan sebagai satu komponen chart bersama di SPA.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / platform
- **Aktor & izin:** Semua pengguna laporan; tidak ada izin khusus.
- **Aturan perilaku (acceptance criteria):**
  1. Hover atau fokus keyboard pada titik/batang/area menampilkan tooltip berisi nilai, satuan, dan tanggal.
  2. Control chart dan CFD mendukung zoom/timeframe dengan seleksi rentang (drag) dan tombol reset zoom.
  3. Setiap chart memiliki tabel data di bawahnya atau toggle "Tampilkan sebagai tabel" sebagai alternatif teks.
  4. Chart mendukung tema terang dan gelap; palet warna memenuhi kontras WCAG 2.1 AA dan seri tidak dibedakan hanya dengan warna (pola atau label langsung).
  5. Format tanggal dan angka mengikuti locale profil pengguna (EN/ID) dan zona waktu yang relevan.
  6. ⚠ Elemen chart dapat dinavigasi dengan keyboard (Tab/panah), memiliki label ARIA yang mendeskripsikan seri dan nilai, dan memiliki ringkasan teks untuk pembaca layar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen chart dipakai juga oleh gadget (modul 17) dan laporan JSM (modul 31). Target aksesibilitas umum di modul 27. Tooltip dasar (butir 1) dan tabel data (butir 3) dibangun bersama Burndown dan Sprint report di Fase 1 karena RPT-009/RPT-010 bergantung padanya; butir lain menyusul di Fase 2.

## Laporan sprint (Scrum)

### RPT-009 Burndown chart

- **Deskripsi:** Laporan Scrum yang melacak sisa pekerjaan dalam sprint dibanding guideline ideal untuk memprediksi apakah sprint selesai tepat waktu, memakai estimation statistic board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer/Anggota Tim, Product Owner/PM; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown sprint menampilkan sprint aktif dan sprint closed board, diurutkan dari yang terbaru; default adalah sprint aktif (atau sprint closed terakhir bila tidak ada yang aktif); sprint future tidak dapat dipilih.
  2. Pemilih statistik menyediakan Story points, Original time estimate, Issue count, atau custom number field board; bila time tracking "Remaining estimate and time spent" aktif, garis mengikuti Remaining estimate (RPT-005).
  3. Garis "Guideline" ditarik dari total komitmen awal sprint pada tanggal mulai ke nol pada tanggal akhir sprint, datar pada non-working days.
  4. Garis "Remaining values" (merah) turun saat work item masuk kolom paling kanan board dan naik saat scope ditambah atau work item di-reopen.
  5. Garis digambar sebagai step chart dengan satu langkah per event: sprint start, work item added, work item removed, estimate changed, work item completed, work item reopened, sprint end.
  6. Tabel event di bawah chart berisi kolom Date, Issue, Event type, Event detail, Inc., Dec., dan Remaining; key work item dapat diklik membuka issue view.
  7. Work item yang ditambahkan setelah sprint dimulai ditandai indikator scope change pada chart dan pada tabel.
  8. Non-working days dibayangi abu-abu berdasarkan Working days board, termasuk time zone board dan tanggal libur.
  9. Hover setiap titik/event menampilkan tooltip berisi tanggal, event, dan nilai sisa (RPT-008).
  10. Mode time tracking: worklog yang mengurangi Remaining estimate menurunkan garis, dan tabel menambah kolom Time spent.
  11. Story points hanya dihitung pada parent; untuk statistik berbasis waktu, estimasi sub-task ikut dihitung dalam total.
  12. Penyelesaian ditentukan oleh kolom paling kanan board (company-managed), bukan hanya oleh field Resolution; work item dengan Resolution terisi tetapi belum di kolom paling kanan tetap dihitung sisa.
  13. Tautan "View in Issue Navigator" membuka navigator dengan JQL `sprint = <id>` digabung board filter.
  14. Team-managed: Burndown memakai Story points, Issue count, atau Original estimate sesuai setting Estimation di Features.
- **Team-managed vs company-managed:** Lihat butir 12 dan 14; pada team-managed penyelesaian ditentukan kolom paling kanan board project.
- **Catatan & ketergantungan:** Bergantung pada RPT-004 (changelog status, sprint, estimasi) dan snapshot SPR-018; TIME-027 (modul 14) hanya rujukan ke fitur ini dan mendefinisikan sumber data estimasinya; prioritas/fase/tier yang berlaku adalah P0 / Fase 1 / free-standard di sini. Versi gadget Sprint burndown ada di modul 17. Endpoint `GET /api/v1/boards/:id/reports/burndown?sprint=`.

### RPT-010 Sprint report

- **Deskripsi:** Ringkasan satu sprint berisi work item selesai, tidak selesai, selesai di luar sprint, dan dihapus, dengan burndown mini; dipakai untuk sprint review dan retrospective.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer/Anggota Tim, Product Owner/PM; akses mengikuti RPT-006; menautkan halaman membutuhkan Manage Sprints atau Edit Issues sesuai SPR-026.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown sprint memuat sprint aktif dan closed dari board; setelah Complete sprint, pengguna diarahkan ke Sprint report sprint tersebut (SPR-020).
  2. Header menampilkan status sprint, tanggal mulai, tanggal selesai terencana, tanggal ditutup (bila closed), dan sprint goal.
  3. Chart burndown mini memakai estimation statistic board ditampilkan di atas daftar.
  4. Bagian "Completed issues" mencantumkan work item yang berada di kolom paling kanan saat sprint ditutup (atau saat ini untuk sprint aktif) beserta nilai estimasinya.
  5. Bagian "Issues not completed" mencantumkan work item yang tidak berada di kolom Done saat sprint ditutup.
  6. Bagian "Issues completed outside of this sprint" mencantumkan work item yang sudah selesai sebelum ditambahkan ke sprint atau diselesaikan di luar board.
  7. Bagian "Issues removed from sprint" mencantumkan work item yang dikeluarkan dari sprint selama sprint berjalan.
  8. Work item yang ditambahkan setelah sprint dimulai diberi tanda asterisk (`*`) dengan legenda "Ditambahkan setelah sprint dimulai".
  9. Perubahan estimasi selama sprint ditampilkan sebagai "nilai awal → nilai akhir".
  10. Setiap bagian menampilkan total estimasi sesuai statistik terpilih.
  11. Setiap bagian memiliki tautan "View in Issue Navigator" dengan JQL yang mereproduksi daftar tersebut.
  12. Bagian "Linked pages" memungkinkan menautkan atau membuat halaman dokumentasi (mis. retrospective) ke sprint melalui integrasi wiki (SPR-026, modul 25).
  13. Team-managed: Sprint report menyediakan pemilihan sprint dan daftar work item completed, not completed, dan removed.
  14. ⚠ Work item yang tidak selesai saat sprint ditutup menampilkan tujuan pemindahannya ("Dipindah ke Backlog" atau "Dipindah ke <nama sprint>") pada kolom keterangan.
- **Team-managed vs company-managed:** Lihat butir 13; team-managed tidak memiliki bagian "Issues completed outside of this sprint".
- **Catatan & ketergantungan:** Snapshot estimasi saat sprint dimulai harus disimpan (SPR-018) atau direkonstruksi dari history (RPT-004). Tautan halaman bergantung modul 25; tanpa integrasi wiki, bagian Linked pages menerima tautan URL biasa. TIME-029 (modul 14) hanya rujukan ke fitur ini untuk aturan commitment dan scope change; prioritas/fase/tier yang berlaku adalah P0 / Fase 1 / free-standard di sini.

### RPT-011 Burnup chart

- **Deskripsi:** Laporan Scrum yang menampilkan pekerjaan selesai dibanding total scope sprint sehingga perubahan scope terlihat terpisah dari progres.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Product Owner/PM, Stakeholder/Viewer; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih sprint dan statistik (Story points, Original time estimate, Issue count); default sama dengan Burndown.
  2. Garis "Work scope" naik/turun saat work item ditambah/dihapus atau estimasi berubah.
  3. Garis "Completed work" naik saat work item selesai dan turun saat work item di-reopen.
  4. Garis "Guideline" ideal ditarik dari nol ke total scope awal sprint pada tanggal akhir.
  5. Non-working days dibayangi sesuai Working days board dan guideline datar pada hari tersebut.
  6. Tabel event di bawah chart menampilkan perubahan scope dan completed per tanggal.
  7. Hover menampilkan tooltip nilai scope dan completed pada tanggal tersebut.
  8. Laporan tersedia di project company-managed dan team-managed (team-managed hanya bila fitur Sprints aktif).
- **Team-managed vs company-managed:** Tidak ada perbedaan selain prasyarat fitur Sprints pada team-managed.
- **Catatan & ketergantungan:** Burnup bersifat per sprint; progres lintas rilis ditangani Version report (RPT-015) dan Release burndown (RPT-016).

### RPT-012 Velocity chart

- **Deskripsi:** Laporan Scrum yang membandingkan Commitment vs Completed per sprint untuk sprint-sprint terakhir, sebagai dasar perencanaan kapasitas.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Product Owner/PM; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Chart menampilkan bar ganda per sprint: Commitment (abu-abu, total estimasi saat sprint dimulai dari snapshot) dan Completed (hijau, total estimasi work item selesai saat sprint ditutup).
  2. ⚠ Default menampilkan 7 sprint closed terakhir; pengguna dapat memilih jumlah sprint (3–15) atau rentang tanggal.
  3. Pemilih statistik menyediakan Story points, Original time estimate, Issue count, dan custom number field board.
  4. Tabel di bawah chart berisi kolom Sprint, Commitment, Completed; nama sprint menautkan ke Sprint report sprint tersebut.
  5. ⚠ Rata-rata velocity (Completed) dari sprint yang ditampilkan ditampilkan sebagai garis/angka ringkasan.
  6. Hanya sprint closed yang dihitung; sprint aktif tidak ditampilkan.
  7. Velocity dihitung per board; work item dari project lain yang masuk board filter ikut dihitung.
  8. Team-managed: Velocity report tersedia dengan pilihan estimasi sesuai Features.
  9. Nilai velocity diekspos lewat API internal untuk dipakai Plans (modul 18) dan kartu Sprint commitment di panel Insights (RPT-033).
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Bergantung pada snapshot komitmen (RPT-004 butir 2). Gadget Velocity di modul 17. TIME-028 (modul 14) hanya rujukan; nilai yang berlaku P1 / Fase 2 / free-standard.

## Laporan epic & rilis

### RPT-013 Epic report

- **Deskripsi:** Laporan company-managed yang menampilkan progres satu epic dari waktu ke waktu beserta daftar work item di dalamnya dan perkiraan jumlah sprint tersisa.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, Scrum Master/Team Lead; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown epic menampilkan epic dari project yang termasuk board filter, dengan pencarian nama/key.
  2. Chart menampilkan total scope epic, work completed, dan jumlah work item tanpa estimasi sepanjang waktu.
  3. Prediksi jumlah sprint tersisa dihitung dari sisa estimasi dibagi velocity rata-rata board (RPT-012); bila velocity belum ada, ditampilkan "Belum dapat diprediksi".
  4. Daftar work item dikelompokkan menjadi Completed, Incomplete, dan Unestimated, masing-masing dengan estimasi.
  5. Persentase work item tanpa estimasi ditampilkan beserta peringatan bila melewati ambang RPT-003.
  6. Tautan "View in Issue Navigator" membuka JQL `parent = <epic>` digabung board filter.
  7. Statistik estimasi mengikuti estimation statistic board.
  8. Laporan hanya tersedia di board Scrum company-managed.
  9. Keanggotaan epic mengikuti relasi Parent (pengganti Epic Link); work item pada level hierarki di atas epic tidak dihitung.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Parent field di modul 05; hierarki di atas epic di modul 18.

### RPT-014 Epic burndown

- **Deskripsi:** Laporan yang menunjukkan laju penyelesaian epic per sprint dan memprediksi berapa sprint lagi untuk menyelesaikannya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, Scrum Master/Team Lead; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown memilih epic dari project dalam board filter.
  2. Bar bertumpuk per sprint menampilkan Work completed, Work remaining, Work added, dan scope yang dikurangi.
  3. ⚠ Predicted sprints dihitung dari rata-rata velocity epic pada tiga sprint terakhir dan ditampilkan sebagai bar proyeksi bergaris.
  4. Tabel per sprint berisi work added, completed, dan remaining.
  5. Daftar work item dikelompokkan Completed, Incomplete, dan Unestimated.
  6. Hover pada bar menampilkan tooltip nilai setiap segmen.
  7. Laporan hanya tersedia di board Scrum company-managed.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Bergantung RPT-004 untuk menentukan sprint di mana setiap perubahan scope epic terjadi.

### RPT-015 Version report

- **Deskripsi:** Laporan progres version terhadap tanggal rilis dengan proyeksi tanggal selesai berdasarkan laju tim.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, release manager, Admin Proyek; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih version (unreleased dan released) dari project yang termasuk board.
  2. Chart menampilkan garis total scope dan completed sepanjang waktu.
  3. Predicted release date dihitung dari velocity harian tim; Optimistic dan Pessimistic date memakai deviasi ±10% dari velocity tersebut.
  4. Penanda Start date dan Release date version ditampilkan sebagai garis vertikal berlabel.
  5. Persentase work item tanpa estimasi ditampilkan dengan peringatan bila melewati ambang.
  6. Daftar work item dikelompokkan Completed, Incomplete, dan Unestimated.
  7. Statistik estimasi mengikuti estimation statistic board.
  8. Laporan dapat dibuka dari halaman Releases/detail version (modul 15) dengan version terpilih.
  9. ⚠ Laporan hanya tersedia di company-managed pada board Scrum; ketersediaan di board Kanban mengikuti keputusan terbuka RPT-001.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Version dan Fix version di modul 15. Entry point dari halaman Releases dan version detail page adalah REL-025 di [15 Rilis, Versi & Komponen](15-rilis-versi-komponen.md); spesifikasi laporan (kalkulasi, chart, ketersediaan) dimiliki fitur ini dan REL-025 tidak mendefinisikannya ulang.

### RPT-016 Release burndown

- **Deskripsi:** Laporan Scrum yang menampilkan sisa pekerjaan sebuah version per sprint dan memprediksi berapa sprint lagi hingga rilis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, release manager, Scrum Master/Team Lead; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown memilih version dari project dalam board.
  2. Bar per sprint menampilkan Work completed, Work remaining, Work added, dan scope reduction.
  3. Prediksi jumlah sprint tersisa dihitung dari velocity rata-rata; bila belum ada sprint closed, prediksi tidak ditampilkan.
  4. Tabel per sprint dan daftar work item Completed/Incomplete/Unestimated ditampilkan di bawah chart.
  5. Statistik estimasi mengikuti estimation statistic board.
  6. Populasi adalah work item dengan Fix version = version terpilih; perubahan Fix version dihitung sebagai work added/removed.
  7. Laporan hanya tersedia di board Scrum company-managed.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Changelog Fix version wajib dicatat (RPT-004 butir 1). Pemilik spesifikasi laporan adalah modul ini (RPT-016); REL-025 di [15 Rilis, Versi & Komponen](15-rilis-versi-komponen.md) hanya menyediakan entry point dari halaman Releases (REL-012) dan version detail page (REL-014) dengan version terpilih otomatis, serta merujuk butir 2-3 fitur ini. Bila terjadi perbedaan, RPT-016 yang berlaku.

## Laporan alur (flow)

### RPT-017 Cumulative flow diagram

- **Deskripsi:** Diagram area bertumpuk yang menunjukkan jumlah work item di tiap kolom board dari waktu ke waktu untuk mendeteksi bottleneck dan WIP; tersedia untuk board Kanban dan Scrum.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, tim Kanban, Product Owner/PM; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan dibuka dari Reports > Cumulative flow diagram pada board, termasuk board Kanban.
  2. Satu area bertumpuk per kolom board (bukan per status) dengan warna berbeda dan urutan sesuai urutan kolom.
  3. Rentang waktu preset Past week, 2 weeks, month, 3 months, 6 months, All time, dan Custom date range.
  4. Panel "Refine report" company-managed menyediakan pilihan Columns, Swimlanes, dan Quick filters yang dihitung (RPT-007).
  5. Hover menampilkan tooltip per tanggal berisi jumlah work item di setiap kolom.
  6. ⚠ Opsi tambahan "Include non-working days" (default aktif) dan "Start at zero" (memulai sumbu Y dari jumlah work item pada awal rentang = 0) tersedia di panel opsi.
  7. ⚠ Overlay opsional menampilkan garis WIP limit kolom bila kolom memiliki WIP limit (BRD-029); overlay Little's Law tidak dibangun pada rilis awal.
  8. ⚠ Data dihitung dari riwayat status (RPT-004) dan tidak terpengaruh oleh filter tampilan sementara di board view (pencarian, quick filter yang aktif di board, "hide completed"); Kanban sub-filter tetap berlaku sesuai RPT-007 butir 5.
  9. Klik area kolom pada tanggal tertentu membuka "View in Issue Navigator" dengan work item yang berada di kolom tersebut pada tanggal itu.
  10. Perubahan konfigurasi kolom diterapkan retroaktif terhadap status historis.
  11. Laporan tersedia di board Scrum dan Kanban.
  12. Team-managed: CFD tersedia dengan opsi lebih sederhana, yaitu filter tanggal, filter work item (assignee, type, label, epic), dan pilihan kolom.
  13. Panduan "Cara membaca chart ini" tersedia dari header laporan.
- **Team-managed vs company-managed:** Lihat butir 12.
- **Catatan & ketergantungan:** Kerangka refine di RPT-007; komponen chart di RPT-008.

### RPT-018 Control chart

- **Deskripsi:** Scatter plot cycle time/lead time per work item untuk kolom board terpilih, dengan rolling average dan standard deviation, untuk menganalisis prediktabilitas alur kerja.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, tim Kanban, process improvement lead; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan dibuka dari Reports > Control chart pada board company-managed (Scrum maupun Kanban).
  2. Setiap titik mewakili satu work item atau kluster work item; sumbu X adalah tanggal keluar dari kolom terpilih, sumbu Y adalah waktu yang dihabiskan di kolom terpilih.
  3. Garis Average, garis Rolling average, dan area Standard deviation ditampilkan dengan legenda.
  4. Statistik ringkas menampilkan Issue count, Average, Median, Min time, dan Max time.
  5. Rentang waktu preset dan Custom dapat dipilih, dan pengguna dapat zoom ke rentang waktu tertentu dengan seleksi drag.
  6. Panel "Refine report" menyediakan Columns (kolom awal–akhir yang dihitung sebagai cycle time), Swimlanes, dan Quick filters; kolom yang dipilih harus berurutan, bila tidak tampil pesan "Pilih kolom yang berurutan".
  7. ⚠ Toggle menampilkan/menyembunyikan rolling average dan standard deviation, serta opsi "Exclude non-working days" yang memadatkan durasi sesuai Working days board.
  8. Klik titik atau kluster menampilkan daftar work item terkait (key, summary, waktu di kolom) dengan tautan ke issue view.
  9. Tautan "View in Issue Navigator" membuka work item yang terpilih/terlihat.
  10. Work item yang bolak-balik antar kolom dihitung dengan menjumlahkan seluruh waktu di kolom terpilih.
  11. Tidak ada ekspor native dari chart; data mentah diperoleh lewat Issue Navigator (RPT-036).
  12. ⚠ Team-managed: Control chart tidak tersedia; URL-nya menampilkan pesan "Laporan tidak tersedia untuk project team-managed".
- **Team-managed vs company-managed:** Hanya company-managed (lihat butir 12).
- **Catatan & ketergantungan:** Toggle "legacy control chart" tidak dibangun. Perhitungan waktu di kolom memakai RPT-004 butir 3.

## Laporan DevOps

### RPT-019 Cycle time report (DevOps)

- **Deskripsi:** Laporan DevOps yang mengukur waktu dari commit pertama hingga deploy ke production, selaras metrik DORA lead time for changes; membutuhkan integrasi development tool dan deployment.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Engineering manager, Scrum Master/Team Lead, Developer/Anggota Tim; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan hanya menampilkan data bila development tool (GitHub, GitLab, Bitbucket, Azure DevOps) dan deployment pipeline terhubung ke project (modul 24).
  2. Chart menampilkan median cycle time per periode (minggu) dalam rentang waktu terpilih.
  3. ⚠ Breakdown fase menampilkan median durasi time to first PR, time in review, time to merge, dan time to deploy.
  4. Filter rentang tanggal tersedia (default 12 minggu terakhir).
  5. Daftar work item dengan cycle time terpanjang ditampilkan untuk investigasi, dengan tautan ke issue view dan development panel.
  6. Laporan tersedia di company-managed dan team-managed.
  7. Bila belum ada data, tampil empty state dengan panduan menghubungkan tool (RPT-003).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung data commit, PR, dan deployment dari modul 24. Pembagian kepemilikan: modul ini (RPT-019) memiliki laporan, yaitu tampilan chart, akses, filter, dan empty state; DEV-028 di [24](24-integrasi-devops.md) memiliki sumber data, yaitu rumus cycle time per work item, stempel waktu PR/deployment, dan aturan pengecualian work item tanpa commit atau deployment. Prioritas diselaraskan dengan DEV-028 (P3 / Fase 4).

### RPT-020 Deployment frequency report

- **Deskripsi:** Laporan DevOps yang menampilkan seberapa sering tim deploy ke production dari data deployment pipeline yang terhubung, selaras metrik DORA deployment frequency.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Engineering manager, DevOps engineer; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan membutuhkan CI/CD (mis. GitHub Actions, GitLab CI, Jenkins, Bitbucket Pipelines, CircleCI, Azure Pipelines, Octopus) yang mengirim deployment info lewat integrasi modul 24.
  2. Bar chart menampilkan jumlah deployment per periode.
  3. Filter environment type tersedia dengan default production.
  4. Filter rentang tanggal tersedia.
  5. Rata-rata deployment per periode ditampilkan sebagai angka ringkasan.
  6. Laporan tersedia di company-managed dan team-managed.
  7. Bila belum ada data deployment, tampil empty state dengan instruksi menghubungkan pipeline.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung entitas Deployment dan environment type di modul 24. Pembagian kepemilikan: modul ini (RPT-020) memiliki laporan, yaitu tampilan chart, akses, filter, dan empty state; DEV-029 di [24](24-integrasi-devops.md) memiliki sumber data, yaitu Deployments API (DEV-004), aturan keunikan deployment, dan daftar deployment yang menjadi dasar hitungan. Prioritas sama dengan DEV-029 (P3 / Fase 4).

## Issue analysis (berbasis project atau saved filter)

### RPT-021 Issue analysis report configuration (common)

- **Deskripsi:** Form konfigurasi standar untuk laporan Issue analysis dan Forecast & management: sumber data (project atau saved filter), periode, dan parameter grafik. Semua laporan RPT-022 s.d. RPT-031 memakai kerangka ini.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Anggota project, Product Owner/PM, Admin Proyek; sumber data mengikuti Browse Projects dan share permission saved filter.
- **Aturan perilaku (acceptance criteria):**
  1. Field "Project or saved filter" menyediakan pencarian project dan saved filter yang dapat diakses pengguna (filter milik sendiri, favorit, dan yang dibagikan), dengan default project aktif.
  2. Field Period menyediakan Hourly, Daily, Weekly, Monthly, Quarterly, dan Yearly (default Daily).
  3. ⚠ Field "Days previously" (default 30) dibatasi maksimum per period: Hourly 10, Daily 300, Weekly 1.750, Monthly 7.500, Quarterly 22.500, Yearly 36.500 hari.
  4. Input divalidasi sebelum generate; parameter tidak valid (kosong, bukan angka, melewati batas) menampilkan pesan galat inline per field, mis. "Days previously untuk period Daily maksimal 300".
  5. Tombol "Next" menghasilkan laporan, dan tombol "Configure" pada halaman hasil kembali ke form dengan nilai sebelumnya terisi.
  6. Seluruh parameter tersimpan di query string URL sehingga hasil dapat di-bookmark dan dibagikan.
  7. Saved filter yang tidak dibagikan ke pengguna lain tidak dapat dibuka oleh mereka lewat URL laporan; mereka melihat pesan dari RPT-006 butir 4.
  8. Kueri menghormati filter sharing permission dan issue security (RPT-004 butir 8).
  9. Batas periode (awal hari/minggu/bulan) dihitung dalam zona waktu profil pengguna yang menjalankan laporan.
- **Team-managed vs company-managed:** Tersedia untuk company-managed; team-managed tidak menampilkan kategori Issue analysis (RPT-002 butir 5), tetapi saved filter yang mencakup work item team-managed tetap dapat dipakai dari project company-managed.
- **Catatan & ketergantungan:** Saved filter di modul 13; JQL di modul 12. Endpoint generik `GET /api/v1/reports/:type?source=project:ABC|filter:123&period=&daysPrevious=`.

### RPT-022 Created vs resolved issues report

- **Deskripsi:** Laporan yang membandingkan jumlah work item dibuat vs diselesaikan per periode untuk melihat apakah backlog bertambah atau berkurang.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Product Owner/PM, Scrum Master/Team Lead; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi memakai Project or saved filter, Period, dan Days previously (RPT-021).
  2. Opsi "Count" menyediakan Cumulative totals atau Count per periode.
  3. Opsi "Display the trend of unresolved" menambahkan garis jumlah work item unresolved per periode.
  4. Opsi "Display versions" menyediakan None, Only major versions, dan All versions; version yang dirilis dalam rentang ditandai sebagai penanda vertikal pada chart.
  5. Area merah menandai periode created > resolved, dan area hijau menandai resolved > created.
  6. Tabel data per periode berisi jumlah created dan resolved.
  7. Perhitungan memakai tanggal field Created dan Resolved; work item yang di-unresolve lalu di-resolve ulang dihitung pada tanggal Resolved terakhir.
  8. Klik titik chart membuka Issue Navigator dengan JQL periode tersebut (mis. `created >= ... AND created < ...`).
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Gadget Created vs resolved di modul 17 memakai kalkulasi yang sama.

### RPT-023 Pie chart report

- **Deskripsi:** Laporan pie chart yang mengelompokkan work item dari project atau saved filter berdasarkan satu field (statistic type).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Scrum Master/Team Lead, Stakeholder/Viewer; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Project or saved filter dan Statistic type: Assignee, Reporter, Priority, Status, Issue type, Resolution, Component, Fix version, Affects version, Labels, Project, serta custom field bertipe select, multi-select, user picker, dan version picker.
  2. Pie chart menampilkan legenda dengan nilai dan persentase setiap slice.
  3. Tabel data berisi nilai, jumlah work item, dan persentase, diurutkan dari jumlah terbesar.
  4. Klik slice atau baris tabel membuka Issue Navigator dengan JQL yang mereproduksi slice tersebut.
  5. Nilai kosong ditampilkan sebagai "None", dan untuk field user sebagai "Unassigned".
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Untuk field multi-nilai (Labels, Component), satu work item dihitung di setiap nilai sehingga total persentase dapat melebihi jumlah work item; tabel mencantumkan catatan ini.

### RPT-024 Average age report

- **Deskripsi:** Laporan rata-rata umur (hari) work item unresolved per interval dalam periode tertentu untuk melihat backlog yang menumpuk.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, support lead; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi memakai Project or saved filter, Period, dan Days previously.
  2. Bar chart menampilkan rata-rata umur work item unresolved pada setiap periode.
  3. Tabel data berisi Period, Issues unresolved, Total age, dan Avg. age.
  4. Work item tanpa Resolution pada akhir periode dianggap unresolved; umur dihitung dari Created sampai akhir periode.
  5. ⚠ Klik bar membuka Issue Navigator dengan work item unresolved pada periode tersebut.
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Membutuhkan rekonstruksi Resolution historis (RPT-004).

### RPT-025 Recently created issues report

- **Deskripsi:** Laporan jumlah work item yang dibuat per periode beserta berapa yang sudah resolved, untuk melihat laju pembuatan work item.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, support lead; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi memakai Project or saved filter, Period, dan Days previously.
  2. Bar bertumpuk per periode pembuatan menampilkan bagian resolved (hijau) dan unresolved (merah) berdasarkan status Resolution saat ini.
  3. Tabel data per periode berisi jumlah created, resolved, dan unresolved.
  4. Klik segmen bar membuka Issue Navigator dengan work item terkait.
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Tidak ada.

### RPT-026 Resolution time report

- **Deskripsi:** Laporan rata-rata waktu penyelesaian work item (Created hingga Resolved) per periode untuk melihat tren kecepatan resolusi.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, support lead; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi memakai Project or saved filter, Period, dan Days previously.
  2. Bar chart menampilkan rata-rata hari resolusi per periode berdasarkan tanggal Resolved.
  3. Tabel data berisi Period, Issues resolved, Total resolution time, dan Avg. resolution time.
  4. Hanya work item dengan Resolution terisi yang dihitung.
  5. Klik bar membuka Issue Navigator dengan work item resolved pada periode tersebut.
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Laporan waktu resolusi berbasis SLA untuk JSM ada di modul 31.

### RPT-027 Single level group by report

- **Deskripsi:** Laporan yang menampilkan hasil saved filter dikelompokkan per nilai satu field, dengan progress resolved/unresolved per kelompok.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Scrum Master/Team Lead; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Filter (saved filter wajib; project tidak dapat dipilih) dan Statistic type sebagai field pengelompokan; tanpa filter, tombol Next nonaktif dengan pesan "Pilih saved filter".
  2. Setiap grup menampilkan nilai, jumlah work item, dan progress bar resolved vs unresolved.
  3. Setiap grup dapat diperluas untuk menampilkan daftar work item dengan kolom Type, Key, Summary, Status, Priority, dan Assignee.
  4. ⚠ Bila time tracking aktif, progress bar dapat beralih ke basis waktu (Time spent vs Remaining estimate).
  5. Setiap grup memiliki tautan ke Issue Navigator dengan JQL filter ditambah nilai grup.
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Saved filter di modul 13.

### RPT-028 Time since issues report

- **Deskripsi:** Laporan jumlah work item berdasarkan field tanggal (Created, Updated, Resolved, Due, custom date) per periode.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Product Owner/PM; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Project or saved filter, Date field (Created, Updated, Resolved, Due date, atau custom field bertipe date/datetime), Period, dan Days previously.
  2. Opsi Cumulative totals menampilkan jumlah kumulatif.
  3. Bar chart menampilkan jumlah work item per periode berdasarkan field tanggal terpilih; work item tanpa nilai tanggal tidak dihitung.
  4. Tabel data per periode ditampilkan di bawah chart.
  5. Klik bar membuka Issue Navigator dengan JQL rentang tanggal periode tersebut.
- **Team-managed vs company-managed:** Mengikuti RPT-021.
- **Catatan & ketergantungan:** Custom field date di modul 05.

## Forecast & management (berbasis time tracking)

### RPT-029 Time tracking report

- **Deskripsi:** Laporan company-managed yang membandingkan Original estimate, Remaining estimate, dan Time spent per work item dalam satu version, beserta akurasi estimasi.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Product Owner/PM, Scrum Master/Team Lead; Browse Projects pada project; nilai worklog mengikuti izin lihat worklog (modul 14).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Konfigurasi berisi Fix version (version tertentu, All versions, atau No fix version), Sorting (Least completed / Most completed), Issues (All / Incomplete issues only), dan Sub-task inclusion (Only including sub-tasks with the selected version / Also including sub-tasks without a version set / Including all sub-tasks).
  2. Tabel per work item berisi kolom Type, Key, Priority, Summary, Original estimate, Estimated time remaining, Time spent, dan Accuracy.
  3. Setiap baris menampilkan bar progres visual (time spent vs remaining terhadap original estimate).
  4. Kolom Accuracy menampilkan selisih over/under estimate dengan warna (merah untuk over, hijau untuk under/tepat).
  5. Baris total di atas tabel menjumlahkan seluruh kolom waktu dan menghitung akurasi keseluruhan.
  6. Format waktu mengikuti pengaturan global time tracking (hours per day, days per week, format tampilan) dari modul 14/22.
  7. Laporan hanya tersedia bila time tracking aktif; bila tidak, tampil empty state RPT-003 butir 5.
  8. Tersedia print view yang ramah cetak browser (tanpa navigasi, tabel penuh).
  9. Laporan menghormati Browse Projects dan issue security; work item tersembunyi tidak ikut total.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Detail worklog dan estimasi di modul 14. TIME-031 (modul 14) hanya rujukan ke fitur ini; nilai yang berlaku P2 / Fase 3 / free-standard.

### RPT-030 User workload report

- **Deskripsi:** Laporan yang menampilkan total Remaining estimate semua work item unresolved yang di-assign ke seorang pengguna, per project.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Scrum Master/Team Lead; viewer hanya melihat project yang dapat diaksesnya.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi User picker (default pengguna saat ini).
  2. ⚠ Opsi Sub-task inclusion menentukan apakah sub-task yang di-assign ke pengguna ikut dihitung (Only include sub-tasks... / Include all sub-tasks).
  3. Tabel per project berisi jumlah work item unresolved dan total remaining workload.
  4. Baris total keseluruhan ditampilkan di bawah tabel.
  5. Setiap baris project menautkan ke Issue Navigator dengan JQL `assignee = <user> AND resolution IS EMPTY AND project = <key>`.
  6. Hanya work item unresolved yang dihitung.
  7. Laporan bergantung pada time tracking aktif.
  8. Hanya project yang dapat diakses viewer yang ditampilkan; workload di project lain tidak dijumlahkan.
- **Team-managed vs company-managed:** Laporan dibuka dari project company-managed; work item team-managed yang dapat diakses ikut dihitung.
- **Catatan & ketergantungan:** User picker memakai API modul 02. TIME-032 (modul 14) hanya rujukan ke fitur ini; nilai yang berlaku P3 / Fase 3 / free-standard.

### RPT-031 Version workload report

- **Deskripsi:** Laporan yang menampilkan sisa pekerjaan per pengguna dan per work item untuk sebuah version unreleased.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, release manager; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Version (hanya unreleased) dan opsi Sub-task inclusion.
  2. Ringkasan per pengguna menampilkan remaining workload dipecah per issue type.
  3. Baris "Unassigned" menampilkan work item tanpa assignee.
  4. Detail per pengguna menampilkan daftar work item dengan Remaining estimate.
  5. Total workload version ditampilkan di bagian atas.
  6. Hanya work item unresolved yang dihitung.
  7. Laporan bergantung pada time tracking aktif.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Version di modul 15. TIME-033 (modul 14) hanya rujukan ke fitur ini; nilai yang berlaku P3 / Fase 3 / free-standard.

### RPT-032 Workload pie chart report

- **Alias inventaris:** Workload Pie Chart Report
- **Deskripsi:** Laporan pie chart yang menampilkan proporsi beban kerja berbasis waktu (Remaining, Original estimate, atau Time spent) per assignee atau field lain dari sebuah project atau saved filter. Dibangun sebagai varian time tracking dari Pie chart report (RPT-023) yang memakai komponen chart yang sama.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, Scrum Master/Team Lead; Browse Projects pada work item yang dihitung, nilai waktu mengikuti izin lihat worklog (modul 14).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Project or saved filter, Statistic type (default Assignee; opsi lain Reporter, Component, Priority, work type), dan statistik waktu Current estimate (Remaining) / Original estimate / Time spent.
  2. Pie chart per nilai statistic type menampilkan persentase dan nilai durasi terformat sesuai pengaturan global time tracking; irisan tanpa nilai (mis. Unassigned) tetap ditampilkan.
  3. Laporan hanya tersedia bila time tracking aktif; bila nonaktif, laporan tidak muncul di menu Reports dan URL langsung menampilkan empty state RPT-003 butir 5.
  4. ⚠ Klik irisan membuka Issue Navigator dengan JQL sumber ditambah kondisi irisan tersebut.
  5. ⚠ Laporan muncul di kategori "Forecast & management" Reports project company-managed; ketersediaan laporan ini di Jira Cloud saat ini belum dipastikan, sehingga dijadwalkan paling akhir di Fase 4.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** TIME-034 (modul 14) hanya rujukan ke fitur ini; nilai yang berlaku P3 / Fase 4 / free-standard. Versi gadget (Workload Pie Chart gadget) ada di modul 17 dan memakai endpoint yang sama. Bila pemilik produk memutuskan tidak membangunnya, kebutuhan sebaran beban tetap dipenuhi RPT-023 (Statistic type Assignee) dan RPT-030.

## Insight kontekstual & ringkasan project

### RPT-033 Board & backlog Insights panel

- **Alias inventaris:** Insights Panel (Backlog & Board); Board insights (Kanban)
- **Deskripsi:** Panel "Insights" di sisi kanan board dan backlog yang menampilkan kartu metrik kontekstual real-time untuk keputusan sprint dan flow, termasuk Sprint commitment, Sprint progress, dan Issues for attention, untuk board Scrum maupun Kanban.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer/Anggota Tim, Product Owner/PM; akses mengikuti akses board (RPT-006).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Insights" di toolbar board dan backlog membuka/menutup panel kanan; status buka/tutup diingat per pengguna per board.
  2. Backlog Scrum menampilkan kartu Sprint commitment yang membandingkan estimasi sprint yang direncanakan (future sprint teratas) dengan rata-rata completed beberapa sprint terakhir (RPT-012), lengkap dengan indikator under/on/over commitment.
  3. ⚠ Backlog Scrum menampilkan kartu Issue type breakdown dan Epic progress (progres epic berdasarkan status category anak).
  4. Board Scrum (active sprint) menampilkan kartu Sprint progress (To Do / In Progress / Done) dan Sprint burndown mini.
  5. Board Scrum menampilkan kartu Issues for attention dan Issue type breakdown.
  6. ⚠ Issues for attention menyorot work item yang flagged (ISS-023), blocked (memiliki link "is blocked by" ke work item yang belum selesai, ISS-029), berprioritas tinggi (Highest/High), overdue, atau lama tidak bergerak (tidak ada perubahan status ≥ 5 hari kerja).
  7. Kartu memakai estimation statistic board (Story points, waktu, atau Issue count).
  8. Bila data belum cukup (mis. belum ada sprint yang selesai), kartu menampilkan pesan "Belum cukup data, selesaikan minimal satu sprint" alih-alih angka.
  9. Kartu memiliki tautan "View report" ke laporan terkait (Velocity, Sprint report, Burndown).
  10. Saat Parallel Sprints aktif (SPR-023), kartu mengikuti sprint yang terpilih di board/backlog.
  11. Klik bagian kartu (mis. segmen status atau kategori attention) memfilter board/backlog ke work item terkait; klik work item membuka issue view.
  12. Setiap kartu memiliki info "Bagaimana ini dihitung?" yang menjelaskan rumus dan sumber data.
  13. Panel tersedia di company-managed dan team-managed.
  14. ⚠ Board admin dapat mematikan Insights untuk board di Board settings > General; pengguna dapat menyembunyikan panel untuk dirinya sendiri.
  15. Board Kanban: tombol Insights di toolbar board/backlog Kanban membuka panel kanan yang sama.
  16. Board Kanban menampilkan kartu Issues for attention dengan kategori yang sama dengan butir 6, termasuk work item yang lama tidak bergerak di kolom.
  17. ⚠ Board Kanban menampilkan kartu flow tambahan: Issue type breakdown dan rata-rata waktu di kolom (time in column) untuk 14 hari terakhir.
  18. Kartu insight dapat disembunyikan per pengguna dan diberi feedback (berguna/tidak berguna) yang dicatat sebagai event analitik produk.
  19. ⚠ Ketersediaan panel per tipe project/tier: tersedia di semua tier untuk project software; project bisnis tanpa board tidak memiliki panel ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan isi kartu; team-managed memakai setting Estimation di Features dan pengaturan matikan Insights ada di Features.
- **Catatan & ketergantungan:** Kalkulasi memakai RPT-004 dan velocity RPT-012. Daftar kartu dapat bertambah; arsitektur kartu dibuat sebagai registry di `internal/features/insight` agar kartu baru tidak mengubah API panel.

### RPT-034 Project Summary page

- **Deskripsi:** Halaman Summary project berisi kartu ringkasan kesehatan project: aktivitas 7 hari terakhir, status, prioritas, tipe kerja, beban tim, dan progres epic. Fitur ini adalah pemilik tunggal isi kartu (daftar kartu, metrik, kalkulasi, dan perilaku klik); PRJ-028 di [04 Proyek](04-proyek.md) memiliki entri halaman (rute, tab view, toggle Features), sedangkan PLAN-025 di [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) dan JSM-PRJ-014 di [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) merujuk ke fitur ini.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Anggota project, Product Owner/PM, Stakeholder/Viewer, Pengguna Tim Bisnis; membutuhkan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan statistik 7 hari terakhir: Completed, Updated, Created, dan Due soon (jatuh tempo dalam 7 hari ke depan); setiap angka dapat diklik membuka daftar work item terkait.
  2. Kartu Status overview menampilkan donut chart per status dengan total work item di tengah.
  3. Kartu Recent activity menampilkan feed perubahan terbaru (pembuatan, transisi, komentar) dengan pagination "Muat lagi".
  4. Kartu Priority breakdown menampilkan bar chart jumlah work item per priority.
  5. Kartu Types of work menampilkan distribusi per work type dengan persentase.
  6. Kartu Team workload menampilkan distribusi work item per assignee, termasuk baris Unassigned.
  7. Kartu Epic progress menampilkan progres setiap epic berdasarkan status category anaknya (To Do / In Progress / Done).
  8. Klik elemen chart membuka view List/search project yang sudah terfilter sesuai elemen tersebut.
  9. Summary tersedia sebagai view di project team-managed dan company-managed, software maupun bisnis.
  10. Tab Summary dapat disembunyikan/ditampilkan lewat pengaturan view project (PRJ-027).
  11. Data ditampilkan real-time (tanpa cache lebih dari 60 detik) dan tidak memerlukan konfigurasi.
  12. Semua angka menghormati Browse Projects dan issue security.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keputusan kepemilikan Summary (empat heading: PRJ-028, RPT-034, PLAN-025, JSM-PRJ-014):
  - Isi kartu dan metrik: RPT-034 (fitur ini). Daftar kartu baku adalah header statistik 7 hari, Status overview, Recent activity, Priority breakdown, Types of work, Team workload, dan Epic progress (butir 1-7). Modul lain wajib merujuk butir-butir ini dan tidak menulis ulang definisinya.
  - Entri halaman (rute, tab, toggle Features, ketersediaan per tipe project): PRJ-028 di [04 Proyek](04-proyek.md).
  - Penempatan Summary di antara view project (business dan software): PLAN-025 di [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md), yang merujuk butir 1-8 fitur ini satu per satu tanpa aturan widget tambahan.
  - Varian service space: JSM-PRJ-014 di [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) memakai kartu yang sama dari fitur ini, ditambah kartu khusus request type/SLA yang didefinisikan di modul 29.
  - Bila terjadi perbedaan definisi kartu, RPT-034 yang berlaku. Prioritas Summary software/business adalah P1 / Fase 2; varian service space mengikuti prioritas dan fase modul 29.
  - Fitur ini menggantikan laporan bisnis lama (RPT-035).

### RPT-035 Business project reports (legacy JWM)

- **Deskripsi:** Laporan sederhana lama pada project bisnis yang kini digantikan halaman Summary (RPT-034); dicatat sebagai rujukan dan tidak dibangun terpisah.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna Tim Bisnis, Admin Proyek (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Laporan beban kerja per assignee dan progres per status dipenuhi oleh kartu Team workload dan Status overview di Summary, bukan laporan terpisah.
  2. Tab Reports terpisah pada project bisnis tidak dibangun; project bisnis memakai view Summary.
  3. Kebutuhan tanpa ketergantungan board/sprint dipenuhi karena Summary tidak memerlukan board.
  4. Klik chart di Summary menuju view List terfilter (RPT-034 butir 8).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## Berbagi, ekspor & ekstensibilitas

### RPT-036 Report sharing, export & printing

- **Deskripsi:** Cara membagikan dan mengeluarkan data laporan: tautan URL, cetak, drill-down ke Issue Navigator untuk ekspor, gadget dashboard, dan embed ke halaman wiki.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, Scrum Master/Team Lead, Stakeholder/Viewer; penerima tautan tetap dievaluasi dengan izinnya sendiri (RPT-006).
- **Aturan perilaku (acceptance criteria):**
  1. URL laporan stabil dan memuat board ID atau project key, jenis laporan, sprint/version/epic terpilih, dan parameter refine; tombol "Salin tautan" menyalin URL lengkap ke clipboard.
  2. Penerima tanpa akses board filter atau project mendapat pesan galat permission (RPT-006 butir 1–2) tanpa bocoran nama sprint, version, atau jumlah work item.
  3. Laporan Issue analysis dan Time tracking memiliki print view (CSS cetak) yang menampilkan chart dan tabel penuh tanpa navigasi.
  4. Drill-down "View in Issue Navigator" membuka daftar work item yang dapat diekspor ke CSV/Excel melalui fitur ekspor modul 13.
  5. ⚠ Laporan agile tidak memiliki ekspor gambar/PDF/CSV native pada rilis awal; tabel data (RPT-008 butir 3) dapat disalin, dan ekspor CSV tabel laporan menjadi keputusan terbuka.
  6. Versi gadget di dashboard (Sprint burndown, Sprint health, Velocity, Created vs resolved, Pie chart, Average age, Resolution time, Recently created, Time since, dll) memakai endpoint kalkulasi laporan yang sama; detail gadget di modul 17.
  7. Chart Pie, Created vs resolved, dan Two dimensional dapat di-embed ke halaman wiki melalui makro chart pada integrasi modul 25, dengan data dievaluasi memakai izin pembaca.
  8. Sprint report menampilkan "Linked pages" wiki (RPT-010 butir 12).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kebutuhan ekspor/chart kustom lanjutan biasanya dipenuhi app Marketplace pihak ketiga (di luar cakupan).

### RPT-037 Report extensibility (app report modules)

- **Deskripsi:** Titik ekstensi agar app menambahkan laporan kustom ke halaman Reports project; isi laporan app bukan tanggung jawab produk ini.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (membuat app), Admin Situs (memasang app); laporan app tampil bagi pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Manifest app dapat mendeklarasikan modul `projectReport` dengan atribut kunci, nama (EN/ID), deskripsi, kategori, thumbnail, dan URL/entry point laporan.
  2. Laporan dari app tampil di kategori "Other" pada Reports hub (RPT-001) kecuali app mendeklarasikan kategori lain yang valid.
  3. Saat dibuka, konteks project (ID, key) dan board (ID, tipe) serta locale pengguna diteruskan ke app melalui parameter yang ditandatangani.
  4. ⚠ Model ekstensi mengikuti mekanisme app modul 25 (setara modul halaman project); tidak ada model ganda lama/baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail ekosistem ekstensi di modul 25; API di modul 23.

## Analitik lanjutan (Enterprise)

### RPT-038 Atlassian Analytics

- **Deskripsi:** Padanan produk ini untuk analitik lintas produk paket Enterprise: ruang analitik terpisah untuk membangun chart dan dashboard dari data lake internal (RPT-041) dan sumber data eksternal. Nama inventaris dipertahankan sebagai rujukan; produk ini tidak memakai merek tersebut.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Situs/org admin (mengelola akses dan koneksi), analis data, eksekutif; membutuhkan product access Analytics.
- **Aturan perilaku (acceptance criteria):**
  1. Analytics hanya aktif pada plan Enterprise dan diakses lewat app switcher; mencakup data work item, Service Management, dan Assets yang dimiliki produk ini.
  2. Sumber data meliputi koneksi data lake internal (work item, sprint, worklog, changelog, Service Management, Assets) dan database eksternal (PostgreSQL, MySQL, SQL Server, Snowflake, BigQuery, Redshift, Databricks) dengan kredensial terenkripsi.
  3. Admin mengelola product access Analytics, group pengguna, dan permission per data source (Can use / Can manage).
  4. ⚠ Data diperbarui periodik (default setiap 3 jam) dan tidak real-time; waktu sinkronisasi terakhir ditampilkan di setiap chart.
  5. Plan Free, Standard, dan Premium tidak menampilkan Analytics; URL-nya menampilkan halaman upgrade.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Layanan analitik milik Atlassian sendiri di luar cakupan; yang dibangun adalah padanan native. Feature flag plan di modul 00.

### RPT-039 Atlassian Analytics charts & query builder

- **Deskripsi:** Pembuat chart Analytics dengan mode Visual SQL tanpa kode dan editor SQL penuh.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Analis data, Product Owner/PM dengan product access Analytics dan izin Can use pada data source.
- **Aturan perilaku (acceptance criteria):**
  1. Mode Visual SQL memungkinkan memilih tabel/kolom, agregasi, filter, group by, dan sort tanpa menulis SQL; SQL yang dihasilkan dapat dilihat.
  2. Mode SQL memungkinkan query manual dengan dialek sumber data; query bersifat read-only dan perintah selain SELECT ditolak dengan pesan "Hanya query baca yang diizinkan".
  3. Visual SQL steps memungkinkan transformasi hasil (pivot, formula kolom, filter, merge) secara berurutan dan dapat dihapus per langkah.
  4. Hasil beberapa query dari sumber berbeda dapat di-merge berdasarkan kolom kunci.
  5. Tipe chart meliputi bar, line, area, pie, donut, scatter, table, pivot table, single value, funnel, map, heatmap, dan gauge.
  6. Format chart mencakup warna, label, axis, dan conditional formatting.
  7. ⚠ Pembuatan query SQL dari bahasa alami dengan AI tersedia hanya bila AI diaktifkan admin (RPT-042 butir 4).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Query dijalankan dengan timeout dan batas baris (modul 27).

### RPT-040 Atlassian Analytics dashboards & sharing

- **Deskripsi:** Dashboard Analytics yang menggabungkan chart dengan kontrol interaktif, template, dan opsi berbagi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Analis data, eksekutif, Product Owner/PM; izin view/edit per dashboard.
- **Aturan perilaku (acceptance criteria):**
  1. Dashboard menyusun chart, blok teks, dan gambar dalam layout grid yang dapat diubah ukuran dan posisinya.
  2. Controls (dropdown, date range, text input, toggle) memfilter chart melalui variabel yang dirujuk di query.
  3. Template dashboard siap pakai tersedia untuk metrik work item dan Service Management.
  4. Dashboard dapat dibagikan per pengguna/group dengan level view atau edit.
  5. Chart Analytics dapat di-embed ke halaman wiki melalui integrasi modul 25.
  6. ⚠ Dashboard dapat diekspor (PDF/PNG), dijadwalkan refresh, dan dikirim sebagai email subscription terjadwal (modul 19).
  7. Dashboard dapat diduplikasi dan dikelompokkan dalam folder.
  8. Data yang terlihat di dashboard mengikuti permission koneksi data lake pembaca (RPT-041 butir 2).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari dashboard operasional modul 17.

### RPT-041 Atlassian Data Lake & data shares

- **Deskripsi:** Lapisan data terkelola yang memuat data produk ini untuk analitik, dipakai Analytics atau dibagikan ke data warehouse eksternal.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Situs/org admin, data engineer.
- **Aturan perilaku (acceptance criteria):**
  1. Admin membuat data lake connection dengan memilih produk (work item, Service Management, Assets) dan project yang disertakan.
  2. Opsi "Enforce permissions" memastikan data yang dibaca pengguna mengikuti Browse Projects dan issue security; bila dimatikan, koneksi hanya boleh diberikan ke group admin dan aksi tercatat di audit log.
  3. Dokumentasi schema tabel (issue, sprint, worklog, changelog, field values, dll) tersedia di dalam produk dengan deskripsi kolom.
  4. ⚠ Data share ke platform eksternal (Snowflake/Databricks atau ekspor Parquet ke object storage) tersedia sebagai opsi lanjutan.
  5. Latensi sinkronisasi (default ≤ 3 jam) dan retensi data (mengikuti retensi produk) ditampilkan di halaman koneksi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Implementasi dapat berupa read replica PostgreSQL dengan view tersanitasi; keputusan arsitektur di modul 27.

## Insight berbantuan AI

### RPT-042 AI-assisted report insights

- **Deskripsi:** Fitur AI opsional yang membantu merangkum atau menafsirkan data project/sprint terkait laporan; tidak masuk MVP.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Product Owner/PM, Scrum Master/Team Lead; Admin Situs mengontrol aktivasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tombol "Ringkas" pada Sprint report dan Summary menghasilkan ringkasan progres dalam bahasa alami (EN/ID) dari data yang dapat dilihat pengguna, dengan label "Dibuat oleh AI".
  2. Drill-down dari ringkasan dapat memakai natural language ke JQL (modul 12) untuk membuka Issue Navigator.
  3. ⚠ Asisten chat dapat menjawab pertanyaan metrik project (mis. "berapa velocity rata-rata 3 sprint terakhir?") dengan mengutip angka dari endpoint laporan, bukan menghitung sendiri.
  4. Admin Situs dapat mengaktifkan/menonaktifkan seluruh fitur AI di tingkat site; bila nonaktif, tombol AI tidak tampil dan API mengembalikan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia model dan kebijakan data AI di modul 27/22; konteks yang dikirim ke model hanya data yang lolos permission (RPT-004 butir 8).

## Insight space & penggunaan agent

### RPT-043 ⚠ Space Insights (dashboard space yang dapat dikustomisasi)

- **Deskripsi:** Dashboard per space (project) yang tampil di tab Reports/Summary dan dapat disusun admin space dari komponen chart, list, smart link, dan activity stream, sehingga halaman Summary tidak lagi berupa kartu tetap. ⚠ Mengikuti rilis Jira Cloud April/Juli 2026; detail perilaku perlu dicocokkan ulang dengan dokumentasi resmi sebelum implementasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard (komponen AI butir 6-7: platform)
- **Aktor & izin:** Admin Proyek/Project Lead (Administer Projects, mengatur layout); anggota space, Stakeholder/Viewer, Pengguna Tim Bisnis (Browse Projects, hanya melihat dan menerapkan filter sementara); Admin Situs mengontrol aktivasi AI.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Setiap space memiliki satu dashboard Space Insights yang dibuka dari tab Reports/Summary; layout default berisi kartu Summary standar (RPT-034) sehingga space lama tidak kehilangan tampilan.
  2. ⚠ Admin space dapat menambah, menghapus, mengubah ukuran, dan menyusun ulang komponen bertipe chart (bar, donut/pie, line/trend), list work item, smart link (pratinjau URL), dan activity stream; perubahan layout tersimpan per space dan berlaku untuk semua viewer.
  3. Pengguna tanpa Administer Projects tidak melihat kontrol edit layout; upaya menyimpan layout lewat API ditolak dengan 403 dan pesan "Hanya admin space yang dapat mengubah layout".
  4. Anggota dapat menerapkan filter sementara (mis. assignee, work type, rentang tanggal) yang hanya berlaku untuk sesi tampilannya dan tidak mengubah layout yang tersimpan; tombol "Reset filter" mengembalikan tampilan default.
  5. Seluruh komponen hanya membaca data dari satu space tersebut; pemilih sumber data tidak menawarkan space lain, dan tautan "Butuh data lintas space? Buat dashboard" mengarahkan ke dashboard modul 17.
  6. ⚠ Bila AI aktif (VISI-003, RPT-042 butir 4), admin dapat membuat chart dari bahasa alami (mis. "jumlah bug per prioritas bulan ini"); sistem menampilkan pratinjau chart dan konfigurasi/JQL yang dihasilkan sebelum disimpan, dan admin dapat mengoreksinya.
  7. ⚠ Bila AI aktif, chart dapat menampilkan callout insight berbahasa biasa di sampingnya (mis. "Bug prioritas tinggi naik 30% dibanding minggu lalu") berlabel "Dibuat oleh AI", hanya dari data yang lolos permission viewer; bila AI nonaktif, callout dan prompt bahasa alami tidak tampil.
  8. Semua angka menghormati Browse Projects dan issue security (RPT-006); klik elemen chart membuka view List terfilter seperti RPT-034 butir 8.
- **Team-managed vs company-managed:** Tidak ada perbedaan; berlaku juga untuk space bisnis dan service space (JSM-PRJ-014 kini dapat dikustomisasi dengan komponen yang sama).
- **Catatan & ketergantungan:** Memperluas RPT-034 (metrik kartu tetap menjadi komponen default); view Summary di PRJ-028 (modul 04) dan PLAN-025 (modul 18) wajib merujuk fitur ini sebagai sumber layout. Komponen chart dan list sebaiknya memakai registry widget yang sama dengan gadget modul 17 agar tidak ada dua engine. AI bergantung RPT-042 dan VISI-003.

### RPT-044 ⚠ Agent usage dashboard

- **Deskripsi:** Laporan di tab Reports yang mengukur pemakaian dan dampak agent AI (ISS-046) di sebuah space: pengguna aktif agent, jumlah sesi, dan cycle time work item yang dikerjakan atau dibantu agent. ⚠ Mengikuti rilis Jira Cloud Oktober 2026; detail perlu dicocokkan ulang dengan dokumentasi resmi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek/Project Lead, Product Owner/PM, Admin Situs; Browse Projects pada space; tampil hanya bila fitur AI aktif.
- **Aturan perilaku (acceptance criteria):**
  1. Laporan "Agent usage" muncul di Reports space dan memuat tiga chart: pengguna aktif agent per periode, jumlah sesi agent, dan cycle time work item yang dikerjakan atau dibantu agent dibanding work item tanpa agent.
  2. ⚠ Work item dihitung "dikerjakan agent" bila agent pernah menjadi assignee, dan "dibantu agent" bila agent di-@mention atau menjalankan aksi pada work item tersebut (sumber event dari ISS-046).
  3. Filter periode (7 hari, 30 hari, 90 hari, rentang kustom) dan filter per agent (multi-select) mengubah seluruh chart sekaligus.
  4. ⚠ Pengguna dapat menyimpan kombinasi filter sebagai saved report; maksimal 5 saved report per space, dan percobaan menyimpan yang keenam ditolak dengan pesan "Batas 5 laporan tersimpan per space tercapai".
  5. Laporan hanya tampil bila fitur AI aktif di tingkat site (VISI-003); bila nonaktif, laporan tidak muncul di Reports dan endpoint mengembalikan 403.
  6. Bila belum ada aktivitas agent pada periode terpilih, laporan menampilkan empty state "Belum ada aktivitas agent pada periode ini" alih-alih chart kosong.
  7. Angka menghormati Browse Projects dan issue security; work item tersembunyi tidak ikut menghitung cycle time.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung ISS-046 (modul 06) untuk event agent dan PRJ-070 (modul 04, fitur baru terkait agent di space) untuk pengaturan agent per space. Cycle time memakai kalkulasi RPT-004/RPT-019. Pencatatan sesi agent diatur di modul 27.

## Pertanyaan terbuka

- Apakah Version report ditampilkan di board Kanban, dan jika ya dengan populasi board filter Kanban (RPT-001, RPT-015)?
- Berapa ambang persentase work item tanpa estimasi yang memicu banner peringatan (default usulan 20%) dan apakah dapat dikonfigurasi per board (RPT-003)?
- Berapa batas maksimum work item per laporan sebelum laporan menolak dihitung (usulan 5.000), dan apakah laporan besar dihitung asinkron sebagai background job (RPT-004)?
- Apakah laporan agile mendapat ekspor CSV/PNG native pada Fase 3, atau cukup lewat drill-down ke Issue Navigator (RPT-036)?
- Apakah team-managed perlu working days per board atau cukup default site (RPT-005)?
- Apakah Analytics dan Data Lake (RPT-038 s.d. RPT-041) dibangun native atau diganti dengan koneksi read-only ke BI eksternal (mis. Metabase/Superset) sebagai jalur yang lebih murah?
- Apakah Kanban sub-filter (BRD-046) ikut membatasi populasi CFD dan Control chart (sehingga work item yang sudah dirilis keluar dari laporan), atau laporan hanya memakai board filter? Inventaris memuat dua klaim yang bertentangan; PRD memilih sub-filter berlaku (RPT-007 butir 5, RPT-017 butir 8).
- Apakah Workload pie chart report (RPT-032) tetap dibangun di Fase 4 atau dibatalkan bila terbukti tidak tersedia lagi di Jira Cloud (kebutuhannya tetap tertutup RPT-023 dan RPT-030)?
- Apakah komponen Space Insights (RPT-043) dan dashboard lintas space (modul 17) memakai satu engine widget yang sama, dan apakah layout Space Insights dapat disalin menjadi dashboard?
- Apakah kriteria "lama tidak bergerak" di Issues for attention (usulan ≥ 5 hari kerja) dapat dikonfigurasi board admin (RPT-033)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk pelaporan lanjutan (mis. eazyBI, Custom Charts, Time in Status, ekspor laporan PDF/Excel kustom, app stand-up/retrospective) tidak dibangun sebagai fitur native; produk ini hanya menyediakan titik ekstensi report (RPT-037) dan API data (modul 23).
- **Business project reports (legacy JWM)** (RPT-035): laporan lama project bisnis; digantikan halaman Summary (RPT-034).
- **Layanan Atlassian lintas produk** (Atlassian Analytics, Atlassian Data Lake, Atlassian Intelligence/Rovo, Rovo agents, makro chart Confluence) tidak diintegrasikan; produk ini membangun padanan native opsional (RPT-038 s.d. RPT-042) dan integrasi wiki generik di modul 25.
- **Toggle legacy control chart** dan overlay Little's Law pada CFD tidak dibangun pada rilis awal.
- **Laporan Service Management** (SLA, CSAT, workload agent) didefinisikan di modul 31, dan **gadget dashboard** di modul 17; modul ini hanya menyediakan kalkulasi dan komponen chart bersama.
