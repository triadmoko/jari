# 14 — Pelacakan Waktu, Worklog & Estimasi

Modul ini mendefinisikan seluruh kapabilitas pelacakan waktu dan estimasi di produk ini: aktivasi dan pengaturan global time tracking (provider, hours/day, days/week, format tampilan, default unit), format input durasi, system field Original Estimate / Remaining Estimate / Time Spent beserta agregat Σ dari sub-task dan Work Ratio, field estimasi relatif (Story Points dan Story point estimate), panel time tracking di work item view, dialog Log work dan aturan penyesuaian Remaining Estimate, pengelolaan worklog (lihat, edit, hapus, izin, visibilitas, riwayat, batasan data), pemakaian estimasi di board, backlog, dan sprint (estimation statistic, time tracking statistic, rollover saat sprint selesai), laporan agile dan laporan waktu/beban kerja yang bergantung pada data estimasi, field JQL waktu dan worklog, kolom navigator dan ekspor, event notifikasi, REST API dan webhook worklog, aksi otomasi, Smart Commits `#time`, impor CSV, estimasi dan kapasitas di Plans, extension point time tracking provider, serta akses mobile. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga tabel `worklogs`, kolom estimasi per work item (detik, `BIGINT`), cache agregat, konfigurasi time tracking global, dan endpoint `/api/v1/issues/{key}/worklogs` adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/timetracking`) mengikuti pola Clean Architecture repo, dengan dialog Log work dan panel Time tracking sebagai komponen baru di SPA `frontend/`.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Time Tracking Permissions (PERM-032), Comment & Worklog Visibility Restriction (PERM-045), Edit Issues (PERM-021), Administer global (PERM-007), role team-managed (PERM-039, PERM-040), issue security (PERM-044), workflow property (PERM-037).
- [04 Proyek](04-proyek.md) — Features team-managed (PRJ-049) dan company-managed (PRJ-050), Screens (PRJ-045), Notifications project (PRJ-053).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — Time tracking fields (CFG-017), Time tracking settings (CFG-058), Story Points & Story point estimate (CFG-016), Number custom field (CFG-026), custom field contexts (CFG-035), field configuration (CFG-043), screen dan screen scheme (CFG-046, CFG-047), field team-managed per work type (CFG-042).
- [06 Work Item Inti](06-work-item-inti.md) — work item view (ISS-007), clone (ISS-031), move (ISS-032), bulk change (ISS-035), realtime (ISS-014).
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — editor rich text ringkas untuk deskripsi worklog.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — Activity tabs (CMT-021), History (CMT-022, CMT-023), Work Log Tab (CMT-024).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — transition screens (WF-030), status category.
- [10 Board](10-board-kolom-kartu.md) — Board settings (BRD-012), Card display (BRD-015), Active Sprint Board (BRD-014), sub-task di board (BRD-022).
- [11 Backlog & Sprint](11-backlog-sprint.md) — Inline estimation (SPR-007), Sprint start (SPR-018), Sprint complete (SPR-020), Sprint field (SPR-012).
- [12 JQL](12-jql-bahasa-kueri.md) — field dan operator durasi, fungsi tanggal dan user.
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — kolom navigator dan ekspor CSV.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Fix Version untuk laporan versi.
- [16 Laporan & Insight](16-laporan-insight.md) — pemilik laporan Burndown (RPT-009), Sprint report (RPT-010), Velocity (RPT-012), Time tracking (RPT-029), User/Version workload (RPT-030, RPT-031), dan keputusan Workload pie chart (RPT-032).
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget yang memakai data waktu, termasuk Workload Pie Chart gadget (DSH-038).
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — Plans: estimasi, roll-up, kapasitas tim.
- [19 Notifikasi](19-notifikasi-email.md) — notification scheme dan event worklog.
- [20 Otomasi](20-otomasi.md) — aksi Log work, trigger Work logged, smart values.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — halaman pengaturan global dan audit log.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — Worklog REST API, webhook worklog.
- [24 Integrasi DevOps](24-integrasi-devops.md) — Smart Commits.
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — impor CSV, extension point provider.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — konkurensi, batas data, mobile/responsif.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Worklog dan event `worklog.*`.

## Persona terkait

- **Admin Situs** — mengaktifkan time tracking, memilih provider, mengatur hours/day, days/week, format, dan default unit; mengimpor estimasi dan worklog dari CSV.
- **Admin Proyek/Project Lead** — mengatur permission scheme worklog, screen yang memuat field time tracking dan Log Work, Features Estimation di team-managed, dan notification scheme event worklog.
- **Scrum Master/Team Lead** — memilih estimation statistic dan time tracking statistic board, memakai total estimasi saat sprint planning, membaca Burndown, Velocity, dan Sprint report.
- **Developer/Anggota Tim** — mengisi estimasi, mencatat waktu lewat Log work, transisi, atau Smart Commits, serta mengedit dan menghapus worklog miliknya.
- **Product Owner/PM** — memakai story points dan forecast release/epic untuk memproyeksikan tanggal rilis; memakai Plans untuk roll-up estimasi dan kapasitas tim.
- **Stakeholder/Viewer** — membaca progress bar time tracking, laporan workload, dan gadget dashboard tanpa mengubah data.
- **Integrator/Developer API** — menyinkronkan worklog ke sistem penggajian/penagihan lewat REST API inkremental dan webhook, atau mendaftarkan provider time tracking dari ekstensi.

## User story modul

- Sebagai Admin Situs, saya ingin mengatur jumlah jam kerja per hari dan hari kerja per minggu agar input "1d" dan "1w" sesuai kebiasaan organisasi (TIME-002, TIME-003).
- Sebagai Developer/Anggota Tim, saya ingin mengisi Original Estimate dan melihat progress bar waktu tercatat vs sisa agar saya tahu apakah pekerjaan masih sesuai perkiraan (TIME-005, TIME-010, TIME-011).
- Sebagai Developer/Anggota Tim, saya ingin mencatat waktu kerja beserta tanggal mulai dan deskripsi dengan Remaining Estimate yang menyesuaikan otomatis agar sisa pekerjaan selalu akurat (TIME-012, TIME-013).
- Sebagai Developer/Anggota Tim, saya ingin mencatat waktu sekaligus saat memindahkan work item ke Done agar tidak perlu membuka dialog terpisah (TIME-014).
- Sebagai Admin Proyek/Project Lead, saya ingin mengatur siapa yang boleh mengedit dan menghapus worklog milik orang lain agar data waktu yang dipakai penagihan tidak dimanipulasi (TIME-016, TIME-017, TIME-018).
- Sebagai Developer/Anggota Tim, saya ingin membatasi visibilitas worklog tertentu ke role atau group agar catatan internal tidak terlihat oleh semua orang (TIME-019).
- Sebagai Scrum Master/Team Lead, saya ingin memilih statistik estimasi board (Story Points, waktu, atau jumlah work item) agar velocity dan burndown mencerminkan cara tim memperkirakan (TIME-022, TIME-023).
- Sebagai Scrum Master/Team Lead, saya ingin melihat total estimasi per sprint dan membandingkannya dengan rata-rata velocity agar komitmen sprint realistis (TIME-025, TIME-028/RPT-012).
- Sebagai Scrum Master/Team Lead, saya ingin estimasi work item yang belum selesai terbawa utuh ke sprint berikutnya agar tidak ada pekerjaan yang hilang dari perhitungan (TIME-026, TIME-029).
- Sebagai Product Owner/PM, saya ingin forecast jumlah sprint tersisa untuk sebuah versi atau Epic agar dapat mengomunikasikan perkiraan tanggal rilis (TIME-030).
- Sebagai Admin Proyek/Project Lead, saya ingin laporan Time Tracking dan workload per user/versi agar dapat melihat akurasi estimasi dan beban kerja tim (TIME-031/RPT-029, TIME-032/RPT-030, TIME-033/RPT-031).
- Sebagai Scrum Master/Team Lead, saya ingin mencari work item yang saya log minggu ini atau yang melebihi estimasi lewat JQL agar mudah menyusun timesheet dan mendeteksi overrun (TIME-007, TIME-037, TIME-038).
- Sebagai Integrator/Developer API, saya ingin mengambil worklog yang berubah atau terhapus sejak waktu tertentu agar sistem penagihan eksternal tetap sinkron tanpa memindai ulang semua data (TIME-040, TIME-042).
- Sebagai Developer/Anggota Tim, saya ingin mencatat waktu langsung dari pesan commit agar pencatatan waktu tidak memutus alur kerja coding (TIME-044).
- Sebagai Product Owner/PM, saya ingin estimasi child ter-roll-up ke Epic dan level di atasnya serta melihat kapasitas tim per sprint agar rencana lintas tim tidak overbooked (TIME-047, TIME-048).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 15 |
| P1 | 14 |
| P2 | 12 |
| P3 | 8 |
| **Total** | **49** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 5 |
| Fase 2 | 23 |
| Fase 3 | 13 |
| Fase 4 | 7 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 hanya membangun estimasi relatif yang dibutuhkan board, backlog, dan sprint dasar, yaitu field Story Points / Story point estimate (TIME-008), estimation statistic board (TIME-022), total estimasi di backlog serta dialog Start sprint (TIME-025), serta aturan estimasi untuk Burndown dan Sprint report Fase 1 milik modul 16 (TIME-027/RPT-009, TIME-029/RPT-010); sejak Fase 1 snapshot estimasi saat sprint dimulai dan setiap perubahan estimasi wajib dicatat di history agar laporan agile Fase 2 dapat dihitung. Fase 2 membangun time tracking lengkap sebagai satu paket: pengaturan global dan parser durasi bersama di backend dan frontend (TIME-001 s.d. TIME-003), system field dalam detik beserta cache agregat (TIME-005, TIME-006), panel dan dialog Log work, penyesuaian Remaining Estimate, CRUD worklog dengan permission (TIME-010 s.d. TIME-021), statistik board berbasis waktu, rollover sprint, varian statistik waktu (Original Time Estimate/Remaining) pada Burndown dan Sprint report serta Velocity, kolom navigator, dan field JQL waktu/worklog. Parser durasi dan konversi detik harus menjadi satu paket bersama (`pkg/duration`) yang dipakai handler, laporan, ekspor, dan JQL agar hasil konsisten. Fase 3 menambah visibilitas worklog terbatas, aspek waktu untuk laporan Time tracking dan User/Version workload yang dimiliki modul 16 (RPT-029 s.d. RPT-031), gadget, Work Ratio, event notifikasi, REST API publik, webhook, otomasi, dan impor CSV. Fase 4 menambah aspek waktu Workload pie chart (TIME-034/RPT-032), Plans (roll-up dan kapasitas), Smart Commits, operasi bulk worklog via API, extension point provider, dan pengalaman mobile.

## Fitur

## Konfigurasi global time tracking

### TIME-001 Time Tracking Provider Activation (Global)

- **Deskripsi:** Pengaturan global untuk mengaktifkan atau menonaktifkan time tracking bawaan produk ini, atau memilih provider yang didaftarkan ekstensi. Pengaturan ini menentukan apakah field estimasi waktu dan worklog muncul di seluruh site.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs dengan global permission Administer (PERM-007); pengguna lain hanya terdampak oleh hasilnya.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > Work items > Time tracking menampilkan status saat ini ("Aktif"/"Nonaktif") beserta tombol Activate atau Deactivate yang sesuai; status default pada instalasi baru adalah aktif dengan provider bawaan.
  2. Dropdown "Time tracking provider" berisi provider bawaan produk ini ditambah setiap provider yang didaftarkan ekstensi terpasang lewat modul provider (TIME-046); provider ekstensi yang dicopot hilang dari daftar dan, bila sedang dipilih, site kembali otomatis ke provider bawaan.
  3. Saat provider ekstensi dipilih, aksi Log work di work item view membuka UI milik ekstensi; field `timetracking` bawaan beserta datanya tetap tersimpan dan tetap dapat dibaca/diubah lewat API.
  4. Saat dinonaktifkan, field Original Estimate, Remaining Estimate, Time Spent, field Log Work, panel Time tracking, dan tab Work log hilang dari seluruh UI; data worklog dan estimasi tidak dihapus dan muncul kembali utuh saat time tracking diaktifkan ulang.
  5. ⚠ Saat dinonaktifkan, laporan Time Tracking, User Workload, Version Workload, dan Workload Pie Chart tidak tersedia di menu Reports, dan field JQL `workratio` ditolak dengan pesan "Field workratio tidak tersedia karena time tracking nonaktif"; field durasi JQL lain (`originalEstimate`, `remainingEstimate`, `timeSpent`) tetap dapat dicari.
  6. Hanya pengguna dengan global permission Administer yang dapat mengubah status dan provider; permintaan lain ditolak HTTP 403.
  7. Pengaturan berlaku global untuk semua project company-managed dan team-managed; tidak ada override per project selain pilihan Estimation team-managed (TIME-009).
  8. API: `GET /api/v1/configuration/timetracking` mengembalikan provider aktif (`key`, `name`, atau kosong bila nonaktif), `PUT` yang sama mengganti provider/menonaktifkan (body kosong = nonaktif), dan `GET /api/v1/configuration/timetracking/list` mengembalikan daftar provider yang tersedia.
  9. ⚠ Setiap aktivasi, deaktivasi, dan pergantian provider tercatat di audit log site (modul 22) dengan kategori "Time tracking", aktor, nilai lama, dan nilai baru.
- **Team-managed vs company-managed:** Berlaku sama untuk keduanya; team-managed tetap memerlukan Features > Estimation = Time agar field waktu muncul di layout (TIME-009).
- **Catatan & ketergantungan:** Melengkapi CFG-058 (modul 05). Extension point provider di TIME-046; aplikasi timesheet pihak ketiga yang memakai extension point tersebut berada di luar cakupan.

### TIME-002 Time Tracking Settings (hours/day, days/week, format, default unit)

- **Deskripsi:** Konfigurasi global untuk konversi satuan waktu, format tampilan durasi, dan unit default input, yang dipakai di seluruh estimasi, worklog, laporan, dan ekspor.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs dengan global permission Administer (PERM-007).
- **Aturan perilaku (acceptance criteria):**
  1. Field "Working hours per day" default 8 dan menerima desimal satu atau dua angka di belakang koma (mis. 7.5).
  2. Field "Working days per week" default 5 dan menerima desimal (mis. 4.5).
  3. Konversi: 1m = 60 detik, 1h = 3600 detik, 1d = hours/day x 3600 detik, 1w = days/week x hours/day x 3600 detik; semua nilai disimpan sebagai bilangan bulat detik.
  4. "Time display format" memiliki tiga opsi: pretty (mis. `4w 3d 2h`), days (mis. `4.5d`), dan hours (mis. `34h`); default pretty.
  5. "Default unit" (minute, hour, day, week; default minute) menentukan arti angka tanpa satuan, mis. input `3` dibaca `3h` bila default unit hour.
  6. Opsi "Copy comment to work description" (default aktif): bila komentar diisi bersamaan dengan Log Work (mis. pada transition screen, TIME-014) dan deskripsi worklog kosong, isi komentar disalin menjadi deskripsi worklog.
  7. Mengubah hours/day atau days/week tidak mengubah detik yang sudah tersimpan, hanya konversi input baru dan tampilannya; contoh: worklog `1d` yang tersimpan 28800 detik ditampilkan `1d 2h` setelah hours/day diubah menjadi 6.
  8. ⚠ Validasi: hours/day harus dalam rentang 1-24 dan days/week dalam rentang 1-7 (desimal diperbolehkan); nilai di luar rentang ditolak dengan pesan "Jam kerja per hari harus antara 1 dan 24" atau "Hari kerja per minggu harus antara 1 dan 7" (selaras CFG-058 butir 2).
  9. Laporan, gadget, ekspor CSV, notifikasi, dan API memakai pengaturan yang sama; API mengembalikan string terformat (mis. `timeSpent`) dan nilai detik (`timeSpentSeconds`) berdampingan.
  10. API: `GET`/`PUT /api/v1/configuration/timetracking/options` dengan field `workingHoursPerDay`, `workingDaysPerWeek`, `timeFormat` (`pretty`|`days`|`hours`), dan `defaultUnit` (`minute`|`hour`|`day`|`week`).
  11. Hanya Admin Situs yang dapat mengubah; pengguna lain dapat membaca pengaturan lewat API (diperlukan untuk parser durasi di klien) tetapi `PUT` ditolak HTTP 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan; pengaturan global berlaku juga untuk team-managed.
- **Catatan & ketergantungan:** Plans memiliki hours/day sendiri (TIME-047). Konflik dengan CFG-058 butir 7 (opsi "Copy comment" dinyatakan tidak dibangun) dicatat di Pertanyaan terbuka.

### TIME-003 Duration Input Format & Parsing

- **Deskripsi:** Format input durasi untuk estimasi dan worklog memakai satuan w/d/h/m, beserta aturan parsing, normalisasi, dan validasi yang sama di UI, API, impor, dan JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang mengisi estimasi atau worklog.
- **Aturan perilaku (acceptance criteria):**
  1. Input gabungan beberapa unit dipisahkan spasi diterima, mis. `2w 4d 6h 45m`; urutan unit bebas tetapi unit yang sama tidak boleh muncul dua kali (`1h 2h` ditolak).
  2. Satuan yang dikenali: `w` (minggu), `d` (hari), `h` (jam), `m` (menit), tidak peka huruf besar/kecil.
  3. Angka tanpa satuan dibaca sesuai default unit global (TIME-002 butir 5).
  4. ⚠ Nilai desimal seperti `1.5h` dan `0.5d` diterima di semua konteks (dialog, API string, impor, JQL) dan dibulatkan ke menit terdekat.
  5. Konversi ke detik mengikuti hours/day dan days/week global yang berlaku saat input disimpan.
  6. ⚠ Input tidak valid (unit tak dikenal, angka negatif, karakter lain) menampilkan galat inline "Format durasi tidak valid. Gunakan format seperti 2w 4d 6h 45m" dan tombol Simpan dinonaktifkan; API mengembalikan HTTP 400 dengan pesan yang sama.
  7. ⚠ Granularitas minimum di UI adalah 1 menit; API menerima `timeSpentSeconds` dalam detik tetapi worklog di bawah 60 detik ditolak HTTP 400 "Time spent minimal 1 menit".
  8. Remaining Estimate dan Original Estimate boleh bernilai 0; Time Spent pada worklog harus lebih dari 0.
  9. Dialog Time tracking dan Log work menampilkan helper text contoh format (`2w 4d 6h 45m`) beserta keterangan satuan `w = minggu, d = hari, h = jam, m = menit`.
  10. Tampilan dinormalisasi sesuai format global, mis. input `90m` ditampilkan `1h 30m` pada format pretty dan `1.5h` pada format hours.
  11. ⚠ Label satuan pada tampilan dilokalkan mengikuti bahasa pengguna (EN/ID lewat layer i18n), sedangkan parser tetap menerima satuan baku `w d h m` di semua bahasa.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Diimplementasikan sebagai paket bersama backend (`pkg/duration`) dan util frontend dengan tabel uji yang sama. Dipakai TIME-005, TIME-012, TIME-037, TIME-040, TIME-045.

### TIME-004 Legacy Mode Time Tracking

- **Deskripsi:** Mode warisan time tracking ketika Original Estimate hanya bisa diubah sebelum ada worklog. Dicatat sebagai rujukan; tidak dibangun karena hanya ada di edisi Data Center.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (di produk referensi).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Time tracking settings produk ini tidak menampilkan checkbox "Legacy mode".
  2. Perilaku legacy (Original Estimate terkunci setelah worklog pertama) tidak diterapkan; Original Estimate tetap dapat diedit setelah ada worklog.
  3. Form edit tidak pernah menyembunyikan salah satu field estimasi berdasarkan ada/tidaknya worklog; Original dan Remaining selalu tampil bersama.
  4. Perilaku standar (setara Legacy OFF) berlaku: Original dan Remaining Estimate dapat diedit secara independen kapan saja oleh pemilik Edit Issues.
  5. Karena tidak ada pergantian mode, data estimasi tersimpan tidak pernah berubah akibat pengaturan mode; importer yang membawa flag legacy mode dari sumber Data Center mengabaikan flag tersebut tanpa mengubah nilai Original, Remaining, maupun worklog yang diimpor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dimasukkan ke bagian Di luar cakupan sebagai rujukan. ⚠ Ketiadaan opsi ini di edisi Cloud produk referensi belum terkonfirmasi (endpoint opsi time tracking Cloud tidak memuat flag legacy); selaras dengan CFG-058 butir 7.

## Field estimasi & nilai waktu

### TIME-005 Time Tracking System Fields (Original Estimate, Remaining Estimate, Time Spent)

- **Deskripsi:** Tiga system field inti time tracking per work item yang disimpan dalam detik dan ditampilkan mengikuti format global.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (Edit Issues untuk estimasi, Work On Issues untuk worklog), Admin Proyek (screen/field configuration), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Original Estimate (`timeoriginalestimate`) menyimpan perkiraan awal total usaha dalam detik.
  2. Remaining Estimate (`timeestimate`) menyimpan sisa usaha yang diperkirakan dalam detik.
  3. Time Spent (`timespent`) adalah jumlah durasi semua worklog work item tersebut dan bersifat read-only.
  4. Ketiga field disimpan sebagai bilangan bulat detik (`BIGINT`) dan ditampilkan sebagai string sesuai format dan unit global (TIME-002).
  5. Field komposit "Time tracking" pada screen menampilkan input Original dan Remaining dalam satu kontrol.
  6. Nilai kosong (`NULL`) dibedakan dari 0: kosong ditampilkan "-" / "Tidak ada estimasi", sedangkan 0 ditampilkan `0m`, dan JQL `IS EMPTY` hanya cocok dengan nilai kosong.
  7. Remaining Estimate tidak pernah negatif; bila logging dengan mode auto melebihi sisa, nilainya menjadi 0.
  8. Mengisi Original Estimate saat Remaining masih kosong dan belum ada worklog otomatis menyetel Remaining sama dengan Original.
  9. Time Spent hanya berubah lewat create/edit/delete worklog; upaya mengubahnya langsung (UI atau API) ditolak HTTP 400 "Time Spent dihitung dari worklog".
  10. Setiap perubahan Original, Remaining, dan Time Spent tercatat di History work item dengan nilai lama dan baru (TIME-020).
  11. Company-managed: field dapat ditambahkan ke screen Create/Edit/View dan transition screen lewat screen scheme (CFG-046, CFG-047).
  12. ⚠ Company-managed: field configuration dapat menyembunyikan field time tracking per project/work type dan dapat menandai field "Time tracking" sebagai required, yang mewajibkan Original Estimate diisi saat create.
  13. Team-managed: field "Original estimate" dan "Time tracking" diatur lewat layout per work type (CFG-042).
  14. API: field `timetracking` berisi `originalEstimate`, `remainingEstimate`, `timeSpent` (string terformat) dan `originalEstimateSeconds`, `remainingEstimateSeconds`, `timeSpentSeconds`; field individual `timeoriginalestimate`, `timeestimate`, `timespent` juga tersedia dalam detik.
  15. Estimasi diubah lewat API create/edit work item memakai field `timetracking` dengan operasi `set` (mengganti) atau `update` (mengubah sebagian), dengan validasi yang sama dengan UI.
- **Team-managed vs company-managed:** Company-managed memakai screen dan field configuration; team-managed memakai layout per work type dan Features > Estimation.
- **Catatan & ketergantungan:** Definisi field di CFG-017. Data model: kolom detik per work item ditambah cache agregat (TIME-006) di modul 28.

### TIME-006 Aggregate Time Tracking (Σ fields from sub-tasks)

- **Deskripsi:** Nilai agregat yang menjumlahkan estimasi dan time spent sebuah work item beserta semua sub-task langsungnya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dan semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Σ Original Estimate (`aggregatetimeoriginalestimate`) = Original Estimate parent + jumlah Original Estimate sub-task.
  2. Σ Remaining Estimate (`aggregatetimeestimate`) = Remaining parent + jumlah Remaining sub-task.
  3. Σ Time Spent (`aggregatetimespent`) = Time Spent parent + jumlah Time Spent sub-task.
  4. Progress = time spent / (time spent + remaining) x 100 (dibulatkan ke bilangan bulat), dan Σ Progress memakai nilai Σ; bila penyebut 0 nilainya kosong.
  5. Hanya sub-task (hierarki level -1) yang diagregasi; child dari Epic atau level di atasnya tidak ikut (roll-up lintas hierarki hanya di Plans, TIME-047).
  6. Original Estimate parent boleh berbeda dari jumlah estimasi sub-task; field Σ selalu menampilkan total, tanpa memaksa konsistensi.
  7. Kolom Σ tersedia di navigator/list view, ekspor CSV, dan API.
  8. ⚠ Panel Time tracking parent menyediakan toggle "Termasuk sub-task" yang mengganti nilai progress bar ke nilai Σ; pilihan diingat per pengguna.
  9. Nilai Σ dihitung ulang otomatis dalam transaksi yang sama ketika worklog/estimasi sub-task berubah, serta saat sub-task dibuat, dihapus, dipindah ke parent lain, atau dikonversi menjadi/dari standard work item.
  10. Field Σ tidak dapat diedit langsung; API menolak penulisan dengan HTTP 400.
  11. Field Σ tidak tersedia di JQL; kueri yang memakainya ditolak dengan pesan "Field tidak didukung dalam pencarian".
  12. Nilai sub-task tetap dijumlahkan walaupun field time tracking disembunyikan di screen/field configuration sub-task.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed memakai level sub-task bawaan.
- **Catatan & ketergantungan:** Cache agregat disimpan di tabel work item dan diperbarui oleh use case worklog/estimasi; rekalkulasi penuh tersedia sebagai job admin untuk perbaikan data (modul 27).

### TIME-007 Work Ratio

- **Deskripsi:** Field turunan berupa persentase Time Spent terhadap Original Estimate, untuk mendeteksi pekerjaan yang melebihi anggaran.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Scrum Master, manajer; dibaca oleh siapa pun dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Rumus: Work Ratio = (Time Spent / Original Estimate) x 100, dibulatkan ke bawah ke bilangan bulat persen.
  2. Nilai kosong bila Original Estimate kosong atau 0.
  3. Nilai dapat lebih dari 100% (mis. 150 berarti waktu tercatat 1,5 kali estimasi).
  4. Tersedia sebagai kolom "Work Ratio" di navigator dan ekspor CSV.
  5. JQL `workratio` (nama field case-insensitive, JQL-028) mendukung operator `=`, `!=`, `>`, `>=`, `<`, `<=`, `IS`, `IS NOT`, `IN`, `NOT IN` dengan nilai angka persen (mis. `workratio > 100`).
  6. Field dan JQL hanya tersedia bila time tracking global aktif (TIME-001 butir 5).
  7. Contoh penggunaan didokumentasikan di help: `workratio > 100 AND resolution IS EMPTY` untuk mencari work item yang melebihi anggaran.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JQL di modul 12 dan TIME-037; kolom navigator di TIME-036.

### TIME-008 Story Points vs Story Point Estimate Fields

- **Deskripsi:** Field numerik untuk estimasi relatif: "Story Points" (company-managed) dan "Story point estimate" (team-managed), masing-masing dengan konteks yang berbeda.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (Edit Issues), Admin Proyek, Admin Situs (field context).
- **Aturan perilaku (acceptance criteria):**
  1. "Story point estimate" adalah field bawaan team-managed yang terkunci: tidak dapat dihapus, diganti nama, atau diubah tipenya.
  2. "Story Points" adalah custom field numerik bawaan untuk company-managed yang memiliki field context per project/work type (CFG-035).
  3. ⚠ Tipe nilai desimal non-negatif (maks 2 angka di belakang koma); nilai negatif ditolak dengan pesan "Estimasi tidak boleh negatif".
  4. Field context menentukan work type yang memiliki Story Points; work type di luar context tidak menampilkan field tersebut dan edit lewat API ditolak.
  5. Board company-managed memakai field yang dipilih di estimation statistic (TIME-022); bila field yang dipilih tidak ada di context work item, badge estimasi tidak tampil dan Board settings menampilkan peringatan "Field estimasi tidak tersedia untuk sebagian work type di board ini".
  6. JQL mendukung `"Story Points"`, `"Story point estimate"`, atau `cf[ID]` dengan operator numerik (`=`, `!=`, `>`, `>=`, `<`, `<=`, `IS EMPTY`, `IS NOT EMPTY`).
  7. Plans dapat memakai field ini sebagai unit estimasi (TIME-047).
  8. Field dapat diedit inline di card board, baris backlog, dan work item view; Enter/blur menyimpan, Esc membatalkan.
  9. ⚠ Produk ini menyimpan kedua field di kolom/entitas yang sama secara internal dengan label berbeda per jenis project, sehingga konsolidasi di masa depan tidak memerlukan migrasi data; API tetap mengekspos kedua nama.
- **Team-managed vs company-managed:** Company-managed memakai Story Points dengan field context; team-managed memakai Story point estimate yang terkunci.
- **Catatan & ketergantungan:** Definisi field di CFG-016; inline edit backlog di SPR-007.

### TIME-009 Time Tracking in Team-managed Projects

- **Deskripsi:** Perilaku estimasi dan time tracking khusus project team-managed yang dikonfigurasi lewat Features project dan layout per work type.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects) untuk Features dan layout; Developer/Anggota Tim untuk mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Features > Estimation dapat diaktifkan/dimatikan; bila aktif, admin memilih "Story points" atau "Time (original estimate)".
  2. Field "Story point estimate", "Original estimate", dan "Time tracking" ditambahkan atau dihapus dari layout per work type (CFG-042).
  3. Estimation = Time: card backlog/board menampilkan Original Estimate dan header sprint menampilkan total Original Estimate terformat.
  4. Estimation = Story points: card dan header memakai field Story point estimate.
  5. Pengaturan global hours/day dan days/week (TIME-002) tetap berlaku.
  6. Tidak tersedia field configuration, screen scheme, maupun transition screen kustom; field Log Work hanya dapat diakses dari work item view.
  7. Laporan klasik Time Tracking, User Workload, dan Version Workload tidak tersedia; laporan yang tersedia adalah Burnup, Burndown, Sprint report, Velocity, dan Cumulative flow.
  8. Burndown dan Velocity memakai statistik sesuai pilihan Estimation project; mengganti pilihan berlaku retroaktif pada laporan.
- **Team-managed vs company-managed:** Fitur ini khusus team-managed; padanan company-managed adalah Board Estimation Statistic (TIME-022) dan field configuration.
- **Catatan & ketergantungan:** Features toggle di PRJ-049; laporan di modul 16.

## Work item view: panel, estimasi & Log work

### TIME-010 Time Tracking Panel & Progress Bar (Issue View)

- **Deskripsi:** Field/panel di work item view yang membandingkan estimasi dengan waktu yang sudah dicatat lewat progress bar, sekaligus menjadi pintu masuk dialog Time tracking.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Stakeholder/Viewer (baca dengan Browse Projects); interaksi tulis memerlukan Work On Issues atau Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field "Time tracking" tampil di panel Details work item view; klik field membuka dialog Time tracking yang memuat aksi Log work dan edit estimasi (TIME-011, TIME-012).
  2. Progress bar memakai tiga warna token tema: biru untuk waktu tercatat, abu-abu untuk sisa, dan oranye untuk bagian yang melebihi Original Estimate; warna memenuhi kontras WCAG AA dan disertai teks, bukan warna saja.
  3. Label teks di bawah bar menampilkan "Xh tercatat" dan "Yh tersisa"; Original Estimate ditampilkan sebagai referensi ("Estimasi awal: Zh").
  4. Bila belum ada worklog, panel menampilkan "Belum ada waktu tercatat" dan bar hanya berisi sisa/estimasi.
  5. Produk ini tidak menyediakan tampilan tiga bar terpisah (Estimated/Remaining/Logged) gaya legacy; satu bar gabungan dipakai di semua mode work item view (ISS-008).
  6. Panel hanya tampil bila time tracking global aktif (TIME-001) dan field Time tracking ada di screen (company-managed) atau layout (team-managed).
  7. Panel parent dapat menampilkan nilai agregat dengan sub-task lewat toggle "Termasuk sub-task" (TIME-006 butir 8).
  8. Bila pengguna tidak memiliki Work On Issues maupun Edit Issues, panel tampil read-only: klik tidak membuka dialog dan kursor tidak berubah menjadi pointer.
  9. Saat Time Spent melebihi Original Estimate, panel menampilkan indikator over-estimate berupa segmen oranye dan teks "Xh melebihi estimasi".
- **Team-managed vs company-managed:** Keberadaan panel mengikuti screen (company-managed) atau layout/Features Estimation (team-managed).
- **Catatan & ketergantungan:** Work item view di ISS-007; nilai diperbarui realtime setelah worklog disimpan oleh pengguna lain (ISS-014).

### TIME-011 Edit Estimates (Original / Remaining) Inline

- **Deskripsi:** Mengubah Original Estimate dan Remaining Estimate langsung dari work item view, dialog create, edit screen, atau badge di backlog/board.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Edit Issues (PERM-021).
- **Aturan perilaku (acceptance criteria):**
  1. Field Original estimate tersedia di dialog create (bila ada di screen/layout) dan di panel Details work item view.
  2. Dialog Time tracking di work item view menyediakan input Original estimate dan Remaining estimate yang dapat diubah bersamaan atau sendiri-sendiri.
  3. Mengubah estimasi memerlukan Edit Issues; tanpa izin tersebut input dinonaktifkan dan API mengembalikan HTTP 403.
  4. Mengisi Original saat Remaining kosong dan belum ada worklog otomatis menyetel Remaining sama dengan Original (TIME-005 butir 8).
  5. Estimasi dapat dikosongkan (clear) sehingga nilainya kembali `NULL`, berbeda dari menyetel 0.
  6. Setiap perubahan tercatat di History dan memicu event `issue.updated` (modul 28).
  7. Bila estimation statistic board = Original Time Estimate, badge estimasi di backlog/board dapat diklik untuk inline edit memakai parser durasi (TIME-003).
  8. ⚠ Bulk change (ISS-035) mendukung operasi "Ubah Original Estimate" dan "Ubah Remaining Estimate" untuk banyak work item sekaligus, dengan validasi format dan izin per work item.
- **Team-managed vs company-managed:** Team-managed hanya dapat mengedit field yang ditambahkan ke layout work type; company-managed mengikuti screen.
- **Catatan & ketergantungan:** Inline edit backlog dirinci di SPR-007; perubahan estimasi di sprint aktif dicatat sebagai scope change (SPR-019, TIME-029).

### TIME-012 Log Work Dialog

- **Deskripsi:** Dialog untuk mencatat worklog pada work item, berisi durasi, tanggal mulai, sisa estimasi, deskripsi, dan visibility.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Work On Issues (PERM-032).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog dapat dibuka dari field Time tracking, menu "•••" > Log work, tombol "Log time" di tab Work log, atau command palette dengan shortcut `.` lalu "Log work".
  2. Field "Time spent" wajib diisi dengan format durasi `w d h m` (TIME-003).
  3. Field "Time remaining" terisi otomatis (remaining lama dikurangi time spent, minimal 0) dan diperbarui saat Time spent diketik; pengguna dapat menimpanya manual.
  4. Field "Date started" berupa date-time picker dengan default waktu sekarang dalam timezone profil pengguna.
  5. ⚠ Field "Work description" memakai editor rich text ringkas (modul 07) yang mendukung mention dan emoji, disimpan dalam format dokumen yang sama dengan komentar.
  6. Kontrol visibility memungkinkan membatasi worklog ke project role/group (TIME-019).
  7. Original Estimate ditampilkan read-only di dialog sebagai referensi.
  8. Validasi: time spent > 0, format durasi valid, dan Date started berupa tanggal valid; galat tampil inline per field dan tombol Simpan nonaktif selama ada galat.
  9. Tanpa Work On Issues, aksi Log work tidak tampil di menu mana pun dan API mengembalikan HTTP 403.
  10. Work item di status dengan workflow property `issue.editable=false` (PERM-037) tidak dapat diberi worklog lewat UI (aksi tidak tampil); hanya API dengan `overrideEditableFlag=true` oleh admin atau ekstensi yang diizinkan.
  11. Setelah disimpan: Time Spent bertambah, Remaining disetel sesuai field Time remaining, Σ parent diperbarui, entri baru muncul di tab Work log, dan event `worklog.created` ("Work Logged On Issue") diterbitkan.
  12. Produk ini tidak menampilkan radio button opsi penyesuaian gaya legacy (Adjust automatically / Leave / Set to / Reduce by) di dialog; opsi tersebut tersedia lewat API (TIME-013 butir 8) dan field Time remaining menggantikannya di UI.
  13. Cancel, Esc, atau klik di luar dialog menutup tanpa menyimpan apa pun; bila ada isian, muncul konfirmasi "Buang perubahan?".
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission (permission scheme vs role team-managed).
- **Catatan & ketergantungan:** Endpoint `POST /api/v1/issues/{key}/worklogs` (TIME-040). Event di modul 28; notifikasi di TIME-039.

### TIME-013 Remaining Estimate Adjustment Options

- **Deskripsi:** Aturan penyesuaian Remaining Estimate saat worklog dibuat, diubah, atau dihapus, yang sama di UI, API, otomasi, dan Smart Commits.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Integrator/Developer API; mengikuti izin create/edit/delete worklog.
- **Aturan perilaku (acceptance criteria):**
  1. Mode `auto` (Adjust automatically) saat create: Remaining = max(0, Remaining - time spent).
  2. Mode `leave` (Leave estimate unchanged): Remaining tidak berubah.
  3. Mode `new` (Set to): Remaining disetel ke nilai `newEstimate`; nilai wajib ada dan valid, bila tidak HTTP 400.
  4. Mode `manual` saat create (Reduce by): Remaining = max(0, Remaining - `reduceBy`); `reduceBy` wajib ada.
  5. Mode `manual` saat delete (Increase by): Remaining = Remaining + `increaseBy`; `increaseBy` wajib ada.
  6. Delete dengan mode `auto`: Remaining ditambah kembali sebesar time spent worklog yang dihapus.
  7. Edit dengan mode `auto`: Remaining disesuaikan sebesar selisih time spent lama dan baru (bertambah bila durasi dikurangi, berkurang bila durasi ditambah, batas bawah 0).
  8. API menerima `adjustEstimate=new|leave|manual|auto` (default `auto`) dengan parameter `newEstimate`, `reduceBy` (hanya create), dan `increaseBy` (hanya delete); parameter yang tidak cocok dengan operasi ditolak HTTP 400.
  9. Bila Remaining kosong dan Original ada, mode `auto` menghitung dari Original (Remaining = max(0, Original - time spent)); bila keduanya kosong, Remaining tetap kosong.
  10. Bila logging melebihi Remaining, Remaining menjadi 0 dan tidak pernah negatif.
  11. UI menyederhanakan opsi menjadi field Time remaining yang terisi otomatis dengan hasil mode `auto` dan dapat diedit; nilai yang ditimpa pengguna dikirim sebagai mode `new`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai TIME-012, TIME-016, TIME-017, TIME-041, TIME-043, TIME-044. Logika diletakkan di use case domain tunggal agar semua kanal konsisten.

### TIME-014 Log Work on Create and Transition Screens

- **Deskripsi:** Field "Log Work" yang dapat ditaruh di screen company-managed agar pengguna mencatat waktu saat create, edit, atau transisi workflow.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (Work On Issues); Admin Proyek/Admin Situs untuk konfigurasi screen dan workflow.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: field "Log Work" dapat ditambahkan ke screen Create, Edit, dan transition screen lewat screen scheme dan workflow (CFG-046, WF-030).
  2. Transition screen yang memuat Log Work menampilkan Time spent, Date started, Time remaining, dan Work description seperti dialog Log work.
  3. Worklog dibuat atomik dalam transaksi yang sama dengan transisi; bila validasi worklog gagal (mis. format durasi), transisi dibatalkan dan galat tampil pada field terkait.
  4. Bila opsi global "Copy comment to work description" aktif (TIME-002 butir 6) dan Work description kosong, komentar transisi dipakai sebagai deskripsi worklog.
  5. Field "Time tracking" (Original/Remaining) juga dapat ditaruh di screen create/transition dan disimpan bersama transisi.
  6. ⚠ Admin dapat mewajibkan Log Work saat transisi lewat validator workflow "Field required" pada field Log Work; transisi tanpa worklog ditolak dengan pesan "Log Work wajib diisi untuk transisi ini".
  7. Team-managed tidak memiliki transition screen kustom; log work hanya dari work item view.
  8. Permission Work On Issues tetap dicek saat transisi; pengguna tanpa izin tersebut tidak melihat field Log Work pada screen, tetapi transisi tetap dapat dijalankan.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai dialog Log work dari work item view.
- **Catatan & ketergantungan:** Bergantung pada transition screens (WF-030) dan validator workflow (modul 09).

## Worklog: kelola, izin, visibilitas & riwayat

### TIME-015 Work Log Activity Tab

- **Deskripsi:** Tab di bagian Activity work item view yang menampilkan semua worklog work item secara kronologis beserta aksi edit dan hapus.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Stakeholder/Viewer (Browse Projects untuk melihat; izin edit/delete sesuai TIME-018).
- **Aturan perilaku (acceptance criteria):**
  1. Bagian Activity memiliki tab All, Comments, History, dan Work log (CMT-021); tab Work log hanya tampil bila time tracking aktif.
  2. Setiap entri menampilkan avatar dan nama author, durasi "tercatat" sesuai format global, tanggal/waktu started dalam timezone viewer, deskripsi rich text, dan label visibility bila dibatasi.
  3. Urutan entri mengikuti preferensi sort activity pengguna (terbaru dulu atau terlama dulu) berdasarkan waktu started.
  4. Tab All menggabungkan worklog dengan komentar dan history dalam satu linimasa berurutan waktu.
  5. Aksi Edit dan Delete inline tampil per entri sesuai permission own/all (TIME-016, TIME-017).
  6. Daftar dimuat bertahap 20 entri per halaman dengan tombol "Muat lebih banyak" untuk work item dengan banyak worklog.
  7. Bila belum ada worklog, tab menampilkan pesan kosong "Belum ada waktu tercatat" beserta tombol "Log time" bila pengguna memiliki Work On Issues.
  8. Hanya worklog milik work item ini yang tampil; worklog sub-task tidak ikut.
  9. ⚠ Setiap entri memiliki permalink `?focusedWorklogId={id}` yang membuka work item, memilih tab Work log, dan menggulir serta menyorot entri tersebut; bila entri tidak terlihat oleh viewer, halaman tetap terbuka tanpa sorotan.
  10. Tidak ada tampilan tab Work Log legacy terpisah; satu tampilan tab dipakai di semua mode work item view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Selaras dengan CMT-024 (modul 08) yang mendefinisikan tampilan tab; modul ini mendefinisikan aturan worklog.

### TIME-016 Edit Worklog

- **Deskripsi:** Mengubah worklog yang sudah ada (durasi, tanggal mulai, deskripsi, visibility, sisa estimasi) sesuai permission own/all.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Author dengan Edit Own Worklogs; Admin Proyek atau pemegang Edit All Worklogs untuk worklog orang lain.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Edit tersedia pada entri di tab Work log dan membuka dialog yang sama dengan Log work, terisi nilai saat ini.
  2. Field yang dapat diubah: time spent, date started, work description, visibility, dan time remaining.
  3. Penyesuaian Remaining memakai mode `auto` (selisih), `leave`, atau `new` (TIME-013 butir 7); UI default `auto`.
  4. Penyimpanan mencatat `updateAuthor` dan timestamp `updated`; `author` dan `created` tidak berubah, dan entri menampilkan label "(diedit)".
  5. Event `worklog.updated` ("Issue Worklog Updated") diterbitkan untuk notifikasi dan webhook `worklog_updated`.
  6. Tombol Edit disembunyikan bila pengguna tidak memiliki Edit Own Worklogs (untuk worklog miliknya) atau Edit All Worklogs; API mengembalikan HTTP 403.
  7. UI tidak menyediakan pemindahan worklog ke work item lain; pemindahan hanya lewat API bulk move (TIME-041).
  8. Author tidak dapat diubah ke pengguna lain: UI tidak menampilkan kontrol author, dan payload edit API yang memuat `author` berbeda dari author asli ditolak HTTP 400 "Author tidak dapat diubah" (konsisten dengan TIME-040 butir 6).
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission.
- **Catatan & ketergantungan:** Perubahan Time Spent tercatat di History (TIME-020).

### TIME-017 Delete Worklog

- **Deskripsi:** Menghapus worklog dengan opsi penyesuaian Remaining Estimate.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Author dengan Delete Own Worklogs; pemegang Delete All Worklogs untuk worklog orang lain.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Delete pada entri worklog membuka dialog konfirmasi "Hapus worklog ini?" yang menampilkan durasi dan tanggal worklog.
  2. Dialog menawarkan opsi Remaining: Adjust automatically (ditambah kembali, default), Leave unchanged, Set to, dan Increase by (TIME-013).
  3. Setelah dihapus, Time Spent work item dan Σ parent berkurang sebesar durasi worklog.
  4. Event `worklog.deleted` ("Issue Worklog Deleted") diterbitkan untuk notifikasi dan webhook `worklog_deleted`.
  5. ID worklog terhapus beserta waktu penghapusan disimpan (tombstone) dan tersedia di endpoint `/api/v1/worklogs/deleted?since=` untuk sinkronisasi (TIME-040).
  6. Aksi mengikuti Delete Own Worklogs / Delete All Worklogs; tanpa izin, tombol tidak tampil dan API mengembalikan HTTP 403.
  7. Menghapus work item ikut menghapus semua worklognya dan mencatat tombstone masing-masing.
  8. Tidak ada recycle bin atau undo untuk worklog; konfirmasi menyebutkan bahwa penghapusan permanen.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission.
- **Catatan & ketergantungan:** History mencatat entri penghapusan (TIME-020).

### TIME-018 Worklog Permissions

- **Deskripsi:** Permission yang mengatur siapa yang dapat membuat, mengedit, dan menghapus worklog serta mengubah estimasi.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (permission scheme), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Work On Issues mengizinkan membuat worklog (UI, API, otomasi, Smart Commits).
  2. Edit Own Worklogs mengizinkan mengedit worklog milik sendiri; Edit All Worklogs mengizinkan mengedit worklog siapa pun.
  3. Delete Own Worklogs mengizinkan menghapus worklog milik sendiri; Delete All Worklogs mengizinkan menghapus worklog siapa pun.
  4. Mengedit Original/Remaining Estimate secara langsung (bukan efek samping worklog) memerlukan Edit Issues; penyesuaian Remaining saat log work hanya memerlukan izin worklog yang bersangkutan.
  5. Company-managed: izin diberikan lewat permission scheme ke project role, group, user, Reporter, Current Assignee, dan grantee lain (PERM-015).
  6. ⚠ Team-managed: role bawaan Administrator dan Member dapat log work, Viewer tidak; custom role memiliki izin granular "Log work on any work item", "Edit their own work log entry", "Edit any work log entry", "Delete their own work log entry", "Delete any work log entry" (PERM-040).
  7. Issue security level membatasi akses ke work item sekaligus seluruh worklognya (PERM-044).
  8. Browse Projects diperlukan untuk melihat worklog apa pun.
  9. Parameter API `overrideEditableFlag` hanya dihormati untuk Admin Situs dan ekstensi terpasang saat work item berada di status non-editable; pemanggil lain mendapat HTTP 403.
- **Team-managed vs company-managed:** Company-managed lewat permission scheme; team-managed lewat role bawaan/custom.
- **Catatan & ketergantungan:** Definisi permission di PERM-032; evaluasi permission umum di PERM-001.

### TIME-019 Worklog Visibility Restriction

- **Deskripsi:** Pembatasan siapa yang dapat melihat sebuah worklog berdasarkan project role atau group.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (penulis worklog), Admin Proyek; Admin Situs untuk setting visibility global.
- **Aturan perilaku (acceptance criteria):**
  1. Pilihan visibility: semua pengguna yang dapat melihat work item (default), project role tertentu, atau group tertentu.
  2. Dropdown hanya menampilkan role/group tempat pengguna menjadi anggota; API menolak target lain dengan HTTP 400.
  3. Worklog terbatas tidak terlihat oleh pengguna di luar role/group target di tab Work log, API, ekspor, webhook, maupun hasil JQL `worklogComment`/`worklogAuthor`/`worklogDate`.
  4. ⚠ Total Time Spent work item tetap mencakup durasi worklog terbatas untuk semua viewer, tanpa menampilkan author, deskripsi, atau tanggal worklog tersebut.
  5. API memakai objek `visibility` `{type: "group"|"role", value, identifier}`; `identifier` (ID group/role) diutamakan bila ada.
  6. Entri worklog terbatas menampilkan ikon gembok dan label "Dibatasi untuk <role/group>".
  7. ⚠ Kontrol visibility tersedia di dialog Log work dan Edit worklog pada semua mode work item view, serta di API.
  8. ⚠ Setting global Comment visibility (role saja, atau group dan role; PERM-045) juga menentukan daftar opsi visibility worklog.
- **Team-managed vs company-managed:** Team-managed membatasi pilihan ke role project (PERM-045).
- **Catatan & ketergantungan:** Mekanisme bersama dengan komentar (CMT-004, PERM-045).

### TIME-020 Time Tracking History & Audit

- **Deskripsi:** Pencatatan perubahan estimasi dan worklog di History work item serta metadata audit worklog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (melihat History dengan Browse Projects), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. History mencatat perubahan Original Estimate, Remaining Estimate, dan Time Spent dengan nilai lama → baru, terformat sesuai setting dan disimpan dalam detik.
  2. History mencatat entri "Worklog Id" saat worklog dibuat atau dihapus (nilai baru atau lama berisi ID worklog).
  3. Setiap worklog menyimpan metadata `author`, `created`, `updateAuthor`, dan `updated`.
  4. Changelog tersedia lewat `GET /api/v1/issues/{key}/changelog` dengan item field `timeoriginalestimate`, `timeestimate`, `timespent`, dan `WorklogId`.
  5. Perubahan oleh rule otomasi tercatat atas nama actor rule (default akun sistem "Automation"), bukan pengguna pemicu rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model changelog di CMT-023; event domain di modul 28.

### TIME-021 Worklog Data Constraints & Edge Cases

- **Deskripsi:** Batasan data dan perilaku edge case pada worklog dan estimasi yang harus ditegakkan konsisten di semua kanal.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Developer/Anggota Tim, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Satu work item menampung maksimal 5.000 worklog (konstanta platform yang sama dengan CMT-024); worklog berikutnya ditolak dengan pesan "Batas worklog per work item telah tercapai".
  2. ⚠ Deskripsi worklog dibatasi 32.767 karakter teks; lebih dari itu ditolak dengan pesan "Deskripsi worklog terlalu panjang".
  3. Worklog dengan Date started di masa depan diperbolehkan.
  4. Worklog milik pengguna yang dinonaktifkan atau dihapus tetap ada; author ditampilkan "Mantan pengguna" (Former user) untuk pengguna terhapus dan nama dengan label "Nonaktif" untuk pengguna nonaktif (USR-008).
  5. Move work item ke project lain (ISS-032) memindahkan seluruh worklog; nilai time tracking tetap tersimpan walaupun field tersembunyi di project tujuan.
  6. ⚠ Clone work item (ISS-031) menyalin Original Estimate dan menyetel Remaining = Original, tetapi tidak menyalin worklog maupun Time Spent.
  7. Work item yang sudah resolved tetap dapat diberi worklog, kecuali statusnya non-editable (TIME-012 butir 10).
  8. Date started disimpan sebagai timestamp UTC beserta offset input, dan dikonversi ke timezone viewer saat ditampilkan.
  9. Log work atas nama pengguna lain tidak didukung di UI maupun API; author selalu pengguna terautentikasi (pengecualian hanya importer, TIME-045).
  10. Time Spent dan Σ tetap konsisten saat dua pengguna log work bersamaan: pembaruan memakai increment atomik di database (bukan read-modify-write) sehingga tidak ada worklog yang hilang dari total.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konkurensi dan batas data di modul 27; penanganan pengguna nonaktif di USR-008.

## Estimasi di board, backlog & sprint

### TIME-022 Board Estimation Statistic

- **Deskripsi:** Pengaturan board Scrum yang menentukan metrik estimasi work item untuk velocity, burndown, dan komitmen sprint.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Board admin (BRD-011), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Board Scrum company-managed memiliki halaman Board settings > Estimation (BRD-012).
  2. Pilihan Estimation statistic: Story Points (default board baru), Original Time Estimate, Issue Count, atau custom field numerik lain yang tersedia di project board.
  3. Nilai estimasi tampil sebagai badge di card backlog dan board.
  4. Total estimasi per sprint dan per bagian Backlog ditampilkan di header masing-masing.
  5. Statistic yang dipilih dipakai Velocity chart, Burndown chart, Sprint report, serta Epic/Release burndown (TIME-027 s.d. TIME-030).
  6. Mengubah statistic berlaku retroaktif: laporan menghitung ulang sprint lama dengan statistic baru berdasarkan history nilai field; halaman settings menampilkan peringatan tersebut sebelum disimpan.
  7. Sub-task tidak dihitung dengan estimation statistic; hanya standard work item yang masuk commitment, velocity, dan burndown.
  8. Story points sub-task tidak ikut velocity maupun total sprint, walaupun terisi.
  9. Board Kanban company-managed tidak memiliki halaman Estimation.
  10. Hanya board admin yang dapat mengubah; pengguna lain melihat pengaturan read-only, dan API mengembalikan HTTP 403.
  11. Team-managed: statistic diatur di Project settings > Features > Estimation (TIME-009), bukan di board settings.
- **Team-managed vs company-managed:** Company-managed per board; team-managed per project lewat Features.
- **Catatan & ketergantungan:** Board settings di BRD-012; history nilai estimasi diperlukan sejak Fase 1 agar perhitungan retroaktif mungkin.

### TIME-023 Board Time Tracking Statistic

- **Deskripsi:** Pengaturan board Scrum untuk melacak sisa pekerjaan berdasarkan Remaining Estimate dan Time Spent selama sprint.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board admin.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi "Time Tracking" di Board settings > Estimation: None (default) atau Remaining Estimate and Time Spent.
  2. ⚠ Opsi hanya muncul bila time tracking global aktif, dan hanya dapat dipilih bila estimation statistic = Original Time Estimate; bila statistic lain dipilih, opsi dinonaktifkan dengan keterangan "Hanya tersedia untuk Original Time Estimate".
  3. Bila aktif, Burndown turun setiap kali Remaining Estimate berkurang (logging atau edit), bukan hanya saat work item selesai.
  4. Remaining estimate sub-task dijumlahkan ke parent saat ditampilkan di board dan dihitung burndown.
  5. Card menampilkan Remaining Estimate (bukan Original) selama sprint aktif.
  6. Bila None, Burndown hanya turun saat work item masuk kolom paling kanan board.
  7. Tidak tersedia di team-managed.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Dipakai TIME-027 butir 1 dan 8.

### TIME-024 Estimate Display on Board Cards

- **Deskripsi:** Penampilan nilai estimasi dan field time tracking pada card board dan backlog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board admin (card layout), Developer/Anggota Tim (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Card Scrum menampilkan badge estimasi (story points atau Original Estimate terformat sesuai estimation statistic); work item tanpa estimasi tidak menampilkan badge di board dan "-" di backlog.
  2. Company-managed: Card layout dapat menambah hingga 3 field ekstra per mode (Backlog dan Active sprint) di board Scrum, atau di board Kanban, termasuk Remaining Estimate, Time Spent, dan Original Estimate (BRD-015).
  3. Board Kanban company-managed tidak menampilkan badge estimation; field estimasi hanya dapat ditampilkan lewat card layout.
  4. Team-managed: estimasi tampil di card saat Features > Estimation aktif.
  5. Durasi pada card diformat sesuai pengaturan global (TIME-002).
- **Team-managed vs company-managed:** Company-managed memakai card layout; team-managed mengikuti Features Estimation.
- **Catatan & ketergantungan:** Card display di BRD-015.

### TIME-025 Estimation in Backlog & Sprint Planning

- **Deskripsi:** Tampilan dan inline edit estimasi di backlog untuk sprint planning, termasuk total per sprint dan breakdown per status category.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (Edit Issues untuk mengubah), Scrum Master/Team Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap baris work item di backlog menampilkan badge estimasi sesuai estimation statistic.
  2. Klik badge membuka inline edit estimasi (SPR-007) bila pengguna memiliki Edit Issues.
  3. Header sprint menampilkan total estimasi dengan breakdown To Do / In Progress / Done dalam tiga badge berwarna berdasarkan status category.
  4. Header bagian Backlog menampilkan total estimasi dan jumlah work item.
  5. ⚠ Panel Epic dan panel Version di backlog menampilkan progress dan total estimasi child (SPR-008, SPR-009).
  6. Dialog Start sprint menampilkan jumlah work item dan total estimasi yang dikomitmenkan, dan nilai tersebut menjadi snapshot commitment (SPR-018).
  7. Backlog insights menampilkan sprint commitment dibandingkan rata-rata velocity tiga sprint terakhir (modul 16).
  8. Total diperbarui realtime tanpa reload saat work item di-drag ke/dari sprint atau estimasinya berubah.
  9. ⚠ Work item tanpa estimasi ditandai dengan badge "-" dan tooltip total sprint menyebut "N work item belum diestimasi".
- **Team-managed vs company-managed:** Statistic company-managed dari board settings, team-managed dari Features Estimation.
- **Catatan & ketergantungan:** Halaman backlog di SPR-001; inline edit di SPR-007.

### TIME-026 Remaining Estimate Rollover on Sprint Completion

- **Deskripsi:** Perilaku estimasi untuk work item yang belum selesai ketika sprint ditutup dan dipindahkan ke sprint berikutnya atau backlog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead dengan Manage Sprints (PERM-025).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Complete sprint memindahkan work item belum selesai ke sprint berikutnya, sprint baru, atau backlog sesuai pilihan (SPR-020).
  2. Story points/estimasi dibawa penuh tanpa partial credit untuk sprint yang ditutup.
  3. Remaining Estimate dan Time Spent tetap apa adanya, tidak di-reset.
  4. Work item belum selesai tidak dihitung sebagai completed di velocity sprint yang ditutup.
  5. Parent yang belum selesai dipindah beserta semua sub-task-nya, termasuk sub-task yang sudah selesai.
  6. Work item yang melintasi beberapa sprint tercatat di field Sprint multi-value (sprint closed + sprint baru, SPR-012).
  7. Sprint report sprint lama mencatat work item tersebut di bagian "Issues not completed" beserta estimasinya.
  8. Estimasi work item di sprint berikutnya menjadi bagian commitment baru saat sprint tersebut dimulai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sprint complete di SPR-020; laporan di TIME-028 (RPT-012) dan TIME-029.

## Laporan agile berbasis estimasi

### TIME-027 Burndown Chart (Estimation-based)

- **Deskripsi:** Rujukan ke RPT-009 Burndown chart di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (chart sisa pekerjaan sprint terhadap guideline ideal). Modul ini hanya menetapkan sumber data estimasi dan time tracking yang dipakai laporan tersebut; varian statistik berbasis waktu (Original time estimate, Remaining estimate, garis Time Spent) baru aktif setelah time tracking Fase 2 (TIME-001, TIME-005, TIME-023) tersedia.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer/Anggota Tim; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown statistic berisi Story points, Original time estimate, Issue count, dan Remaining estimate (hanya bila time tracking statistic aktif, TIME-023), mengikuti RPT-009 butir 2; default mengikuti estimation statistic board (TIME-022). Di Fase 1 hanya Story points dan Issue count yang tersedia; opsi berbasis waktu muncul di Fase 2.
  2. Guideline ideal ditarik dari total commitment saat sprint dimulai (snapshot SPR-018) ke 0 di tanggal akhir sprint, mengikuti RPT-009 butir 3.
  3. Garis Remaining Values turun saat work item masuk kolom paling kanan atau, dalam mode time tracking, saat Remaining berkurang karena worklog (TIME-012), mengikuti RPT-009 butir 4 dan 10.
  4. Penambahan work item ke sprint atau kenaikan estimasi (history estimasi TIME-020) ditampilkan sebagai lonjakan naik (scope change), mengikuti RPT-009 butir 5 dan 7.
  5. Hari non-working board diarsir dan tidak menurunkan guideline (RPT-009 butir 8).
  6. Tabel event memuat perubahan estimasi dengan notasi `old → new` dan nilai remaining setelah event; struktur kolom mengikuti RPT-009 butir 6.
  7. Sprint dipilih dari dropdown berisi sprint aktif dan closed board (RPT-009 butir 1).
  8. Garis/kolom Time Spent tampil bila time tracking statistic aktif (Fase 2), dengan nilai dari field TIME-005 dan format durasi TIME-002 (RPT-009 butir 10).
  9. Tersedia di company-managed dan team-managed (RPT-009 butir 14).
- **Team-managed vs company-managed:** Team-managed memakai statistic dari Features Estimation (TIME-009) dan tidak memiliki opsi Remaining estimate; selebihnya mengikuti RPT-009.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-009 (modul 16), yang menetapkan P0 / Fase 1 / free-standard; nilai di heading ini disamakan. Modul ini hanya menambah varian statistik Original Time Estimate/Remaining/Time Spent (Fase 2) di atas laporan Fase 1. Jangan menambah aturan tampilan di sini; ubah RPT-009. Snapshot commitment dari SPR-018; gadget Sprint Burndown di TIME-035.

### TIME-028 Velocity Chart

- **Deskripsi:** Rujukan ke RPT-012 Velocity chart di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (chart Commitment vs Completed per sprint). Modul ini hanya menetapkan aturan sumber estimasi yang dipakai laporan tersebut.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Product Owner/PM; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Bar Commitment dan Completed per sprint, default 7 sprint closed terakhir, mengikuti RPT-012 butir 1-2.
  2. Nilai dihitung memakai estimation statistic board (TIME-022), sesuai RPT-012 butir 3.
  3. Tabel Sprint/Commitment/Completed mengikuti RPT-012 butir 4.
  4. Commitment memakai snapshot estimasi saat sprint dimulai (SPR-018) ditambah work item yang ditambahkan selama sprint dengan estimasinya saat ditambahkan (RPT-004 butir 2).
  5. Rata-rata velocity pada chart adalah rata-rata Completed dari sprint yang ditampilkan (RPT-012 butir 5). Konsumen lain memakai definisi masing-masing: forecast Release/Epic burndown memakai rata-rata tiga sprint closed terakhir (TIME-030, RPT-014), dan default kapasitas Plans memakai nilai velocity yang diekspos RPT-012 butir 9 (TIME-048).
  6. Mengubah estimation statistic board menghitung ulang chart secara retroaktif dari history estimasi (TIME-020).
  7. Tersedia di company-managed dan team-managed (RPT-012 butir 8).
- **Team-managed vs company-managed:** Team-managed memakai statistic dari Features Estimation (TIME-009); selebihnya mengikuti RPT-012.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-012 (modul 16). Jangan menambah aturan tampilan di sini; ubah RPT-012.

### TIME-029 Sprint Report (Estimation & Scope Change)

- **Deskripsi:** Rujukan ke RPT-010 Sprint report di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (work item completed/not completed/removed beserta burndown mini). Modul ini hanya menetapkan aturan commitment, scope change, dan sumber estimasi yang dipakai laporan tersebut.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer/Anggota Tim; akses mengikuti RPT-006.
- **Aturan perilaku (acceptance criteria):**
  1. Estimasi setiap work item saat sprint dimulai dicatat sebagai commitment (snapshot SPR-018), menjadi dasar nilai awal di RPT-010.
  2. Perubahan estimasi setelah sprint dimulai dicatat sebagai scope change dari history estimasi (TIME-020) dan ditampilkan dengan notasi `old → new` pada kolom estimasi (RPT-010 butir 9).
  3. Mengisi estimasi work item yang sebelumnya kosong selama sprint dihitung sebagai scope change (perubahan dari kosong ke nilai).
  4. Work item yang ditambahkan setelah sprint dimulai diberi tanda asterisk (*) dengan legenda "Ditambahkan setelah sprint dimulai" (RPT-010 butir 8).
  5. Bagian Completed Issues, Issues Not Completed, Issues completed outside of this sprint, dan Issues Removed From Sprint mengikuti RPT-010 butir 4-7.
  6. Ringkasan total estimasi completed vs not completed dihitung memakai estimation statistic board (TIME-022), mengikuti RPT-010 butir 10.
  7. Burndown ringkas di bagian atas laporan memakai aturan TIME-027 (RPT-010 butir 3).
  8. Tersedia di company-managed dan team-managed (RPT-010 butir 13).
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber statistic (TIME-009); bagian laporan per jenis project mengikuti RPT-010.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-010 (modul 16), yang menetapkan P0 / Fase 1 / free-standard; nilai di heading ini disamakan. Statistik berbasis waktu (Original Time Estimate/Remaining) baru tersedia di Fase 2 bersama time tracking. Scope change di SPR-019.

### TIME-030 Release / Epic Burndown & Forecast

- **Deskripsi:** Laporan lintas sprint yang memproyeksikan penyelesaian versi atau Epic berdasarkan velocity dan estimasi.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, Scrum Master/Team Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Release burndown menampilkan per sprint: work completed, work remaining, dan work added untuk versi terpilih.
  2. Epic burndown menampilkan data yang sama untuk Epic terpilih.
  3. Forecast jumlah sprint tersisa = sisa estimasi dibagi rata-rata velocity tiga sprint closed terakhir (sumber velocity RPT-012/TIME-028; jendela tiga sprint khusus forecast, berbeda dari rata-rata sprint yang ditampilkan di Velocity chart), dibulatkan ke atas; bila belum ada sprint closed (velocity tidak tersedia), forecast tidak ditampilkan dan diganti keterangan "Velocity belum tersedia".
  4. Bila ada work item tanpa estimasi, laporan menampilkan peringatan "N work item belum diestimasi; forecast mungkin tidak akurat".
  5. Version report menampilkan progress berbasis estimasi dengan prediksi tanggal rilis optimistic dan pessimistic.
  6. Epic report menampilkan jumlah work item estimated vs unestimated beserta persentasenya.
  7. Burnup chart menampilkan scope total vs completed work per sprint, tersedia di team-managed dan company-managed.
  8. Release/Epic burndown, Version report, dan Epic report hanya tersedia di board Scrum company-managed.
- **Team-managed vs company-managed:** Team-managed hanya mendapat Burnup chart.
- **Catatan & ketergantungan:** Versi di modul 15; laporan di modul 16.

## Laporan waktu & beban kerja

### TIME-031 Time Tracking Report

- **Deskripsi:** Rujukan ke RPT-029 Time tracking report di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (Original Estimate, Remaining, Time Spent, dan Accuracy per work item untuk sebuah versi). Modul ini menetapkan sumber data waktu dan rumus akurasinya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Product Owner/PM, Scrum Master/Team Lead (Browse Projects); nilai worklog mengikuti TIME-018 dan TIME-019.
- **Aturan perilaku (acceptance criteria):**
  1. Pilihan Fix Version (version tertentu, semua versi, tanpa versi) mengikuti RPT-029 butir 1.
  2. Urutan Least completed first / Most completed first mengikuti RPT-029 butir 1.
  3. Opsi sub-task inclusion mengikuti RPT-029 butir 1; nilai sub-task diambil dari field agregat TIME-006.
  4. Kolom Original Estimate, Estimated Time Remaining, Time Spent, Accuracy, dan progress bar mengikuti RPT-029 butir 2-4, dengan format durasi dari TIME-002.
  5. Accuracy = Original - (Spent + Remaining); nilai positif berarti lebih cepat dari estimasi, negatif berarti melebihi estimasi (RPT-029 butir 4).
  6. Baris total versi mengikuti RPT-029 butir 5.
  7. Hanya work item yang memiliki data time tracking (minimal satu dari Original, Remaining, atau Time Spent terisi) yang muncul.
  8. Bila time tracking nonaktif (TIME-001), laporan menampilkan empty state RPT-029 butir 7.
  9. Tidak tersedia di team-managed.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-029 (modul 16); butir 7 di atas perlu tercermin di RPT-029. Versi di modul 15.

### TIME-032 User Workload Report

- **Deskripsi:** Rujukan ke RPT-030 User workload report di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (total Remaining Estimate work item unresolved milik seorang pengguna, per project).
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Scrum Master/Team Lead; viewer hanya melihat project yang dapat diaksesnya.
- **Aturan perilaku (acceptance criteria):**
  1. Pemilihan pengguna dengan user picker mengikuti RPT-030 butir 1.
  2. Opsi sub-task inclusion mengikuti RPT-030 butir 2.
  3. Tabel per project (jumlah unresolved dan total Remaining Estimate) mengikuti RPT-030 butir 3; Remaining diambil dari field TIME-005.
  4. Baris total workload mengikuti RPT-030 butir 4.
  5. Hanya work item yang dapat dilihat pembuka laporan yang dihitung (RPT-030 butir 8).
  6. Diakses dari Reports project company-managed dan memerlukan time tracking aktif (RPT-030 butir 7).
- **Team-managed vs company-managed:** Mengikuti RPT-030.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-030 (modul 16). Prioritas disamakan ke P3 sesuai RPT-030.

### TIME-033 Version Workload Report

- **Deskripsi:** Rujukan ke RPT-031 Version workload report di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (sisa Remaining Estimate per pengguna dan per work item untuk version unreleased).
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner/PM, Admin Proyek, release manager; akses mengikuti RPT-021.
- **Aturan perilaku (acceptance criteria):**
  1. Pemilihan version unreleased mengikuti RPT-031 butir 1.
  2. Opsi sub-task inclusion mengikuti RPT-031 butir 1.
  3. Ringkasan per assignee per work type termasuk baris Unassigned mengikuti RPT-031 butir 2-3.
  4. Daftar work item per pengguna beserta Remaining Estimate mengikuti RPT-031 butir 4.
  5. Total Remaining version mengikuti RPT-031 butir 5.
  6. Hanya work item unresolved yang dihitung (RPT-031 butir 6).
  7. Hanya di company-managed dan memerlukan time tracking aktif (RPT-031 butir 7).
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pemilik laporan: RPT-031 (modul 16); versi di modul 15. Prioritas disamakan ke P3 sesuai RPT-031.

### TIME-034 Workload Pie Chart Report

- **Deskripsi:** Rujukan ke RPT-032 Workload pie chart report di [16 Laporan & Insight](16-laporan-insight.md), yang menjadi pemilik spesifikasi laporan (pie chart proporsi beban kerja berbasis waktu per assignee atau field lain dari project atau saved filter). Laporan ini dibangun di Fase 4; modul ini hanya mengatur statistik waktu/estimasi yang dipakainya.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, Scrum Master/Team Lead; Browse Projects, nilai waktu mengikuti TIME-018 dan TIME-019.
- **Aturan perilaku (acceptance criteria):**
  1. Sumber data project atau saved filter mengikuti RPT-032 butir 1.
  2. Statistic type (default Assignee) mengikuti RPT-032 butir 1.
  3. Time field Current Estimate (Remaining), Original Estimate, atau Time Spent diambil dari field TIME-005 (bukan agregat Σ kecuali sub-task ikut dihitung sesuai TIME-006).
  4. Persentase dan nilai durasi per irisan diformat dengan format durasi global TIME-002 (RPT-032 butir 2).
  5. ⚠ Klik irisan membuka navigator dengan JQL sumber ditambah kondisi irisan (RPT-032 butir 4).
  6. Laporan hanya muncul di Reports project company-managed bila time tracking aktif (TIME-001 butir 5, RPT-032 butir 3).
  7. Gadget Workload Pie Chart (DSH-038) memakai endpoint dan aturan statistik waktu yang sama; gadget Pie Chart standar dan Pie chart report (RPT-023) tetap memakai jumlah work item, bukan waktu.
- **Team-managed vs company-managed:** Hanya company-managed (RPT-032).
- **Catatan & ketergantungan:** Pemilik: RPT-032 ([16](16-laporan-insight.md)); modul ini hanya mengatur statistik waktu/estimasi. Gadget: DSH-038 ([17](17-dashboard-gadget.md)), juga Fase 4.

### TIME-035 Dashboard Gadgets with Time/Estimation Data

- **Deskripsi:** Gadget dashboard yang menampilkan data estimasi atau time tracking. Detail kerangka gadget ada di modul 17.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, manajer, Stakeholder/Viewer (sesuai izin dashboard dan filter).
- **Aturan perilaku (acceptance criteria):**
  1. Gadget Filter Results dapat menampilkan kolom Original/Remaining/Time Spent, Σ Original/Remaining/Time Spent, dan Work Ratio.
  2. Gadget Sprint Burndown menampilkan burndown sprint aktif board terpilih berdasarkan estimation statistic board.
  3. Gadget Sprint Health menampilkan persentase estimasi done/in progress/to do, persentase scope change, dan jumlah blocker (flagged).
  4. ⚠ Gadget Days Remaining in Sprint menampilkan jumlah hari kerja tersisa sprint aktif board terpilih.
  5. ⚠ Gadget Agile Wallboard menampilkan board terpilih dalam mode layar besar read-only dengan badge estimasi.
  6. ⚠ Gadget Workload Pie Chart dispesifikasikan di DSH-038 (Fase 4) dan memakai endpoint yang sama dengan laporan TIME-034/RPT-032 (juga Fase 4).
- **Team-managed vs company-managed:** Gadget berbasis board mendukung board dari kedua jenis project.
- **Catatan & ketergantungan:** Dashboard dan gadget di modul 17.

## Pencarian, JQL & ekspor

### TIME-036 Issue Navigator Time Columns & Export

- **Deskripsi:** Kolom waktu dan estimasi di navigator/list view beserta ekspornya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, manajer (Browse Projects).
- **Aturan perilaku (acceptance criteria):**
  1. Kolom yang dapat dipilih: Original Estimate, Remaining Estimate, Time Spent, Σ Original Estimate, Σ Remaining Estimate, Σ Time Spent, Progress, Σ Progress, dan Work Ratio; nilai ditampilkan sesuai format global.
  2. Kolom Story Points dan Story point estimate dapat dipilih.
  3. Hasil dapat diurutkan berdasarkan kolom waktu (kecuali kolom Σ dan Progress), dengan nilai kosong di akhir untuk urutan naik.
  4. ⚠ Kolom Original Estimate, Remaining Estimate, dan Story Points dapat diedit inline di list view oleh pemilik Edit Issues, memakai parser durasi yang sama.
  5. Ekspor CSV (all fields maupun current fields) menyertakan nilai waktu dalam detik.
  6. ⚠ Ekspor CSV all fields menyertakan kolom "Log Work" berulang, satu per worklog yang terlihat oleh pengekspor, dengan format `comment;started;authorAccountId;timeSpentSeconds`, sehingga dapat diimpor kembali (TIME-045).
  7. Navigator tidak menampilkan baris total/agregasi native; total tersedia lewat laporan dan gadget.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Navigator dan ekspor di modul 13.

### TIME-037 JQL Time Tracking Fields

- **Deskripsi:** Rujukan ke JQL-028 Time Tracking & Worklog Fields di [12 JQL](12-jql-bahasa-kueri.md), yang menjadi pemilik sintaks, nama field, dan operator. Grammar dan operator di JQL-028; modul ini hanya menetapkan aturan khusus nilai waktu dan estimasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Field `originalEstimate` (alias `timeOriginalEstimate`) mengikuti JQL-028 butir 1; nilainya dari field TIME-005.
  2. Field `remainingEstimate` (alias `timeEstimate`) mengikuti JQL-028 butir 2.
  3. Field `timeSpent` mengikuti JQL-028 butir 3.
  4. Field `workratio` (nama kanonik; pencocokan nama field case-insensitive dicatat di JQL-028) mengikuti JQL-028 butir 4 dengan rumus TIME-007; work item tanpa original estimate tidak cocok dengan perbandingan apa pun.
  5. Penulisan nilai durasi (`2h`, `"1d 4h"`, kutip wajib untuk nilai berspasi) mengikuti JQL-028 butir 5; konversi memakai parser bersama TIME-003 dan pengaturan TIME-002.
  6. Operator yang didukung mengikuti JQL-028 butir 1-4.
  7. Field Σ/agregat tidak tersedia di JQL (TIME-006 butir 11, JQL-028 butir 11).
  8. `"Story Points"` dan `"Story point estimate"` dapat dicari sebagai field numerik (TIME-008).
  9. `ORDER BY` didukung untuk semua field di atas sesuai aturan pengurutan modul 12.
  10. ⚠ `timeSpent` dibandingkan dengan total Time Spent work item, termasuk durasi worklog yang dibatasi visibility-nya (konsisten dengan TIME-019 butir 4).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik sintaks/operator: JQL-028 (modul 12). Butir khusus modul ini adalah butir 8 dan 10; butir lain hanya rujukan. Konversi durasi memakai paket bersama TIME-003.

### TIME-038 JQL Worklog Fields

- **Deskripsi:** Field JQL untuk mencari work item berdasarkan atribut worklog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, manajer.
- **Aturan perilaku (acceptance criteria):**
  1. `worklogAuthor` mendukung `=`, `!=`, `IN`, `NOT IN`, `IS EMPTY`, `IS NOT EMPTY`, serta fungsi `currentUser()` dan `membersOf()`.
  2. `worklogDate` mendukung `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, tanggal relatif (mis. `-1w`), dan fungsi tanggal seperti `startOfWeek()` dan `endOfMonth()`.
  3. `worklogComment` mendukung pencarian teks `~` dan `!~` pada deskripsi worklog.
  4. `worklogAuthor` dan `worklogDate` dievaluasi per work item, bukan harus pada worklog yang sama; dokumentasi dan tooltip kueri menjelaskan bahwa `worklogAuthor = currentUser() AND worklogDate >= startOfWeek()` juga cocok bila worklog milik pengguna lebih lama dan worklog minggu ini milik orang lain.
  5. ⚠ Hanya worklog yang terlihat oleh pengguna pencari (TIME-019) yang dipertimbangkan dalam evaluasi ketiga field.
  6. Kasus umum `worklogAuthor = currentUser() AND worklogDate >= startOfWeek()` tersedia sebagai contoh kueri di help dan memberikan hasil dalam batas performa pencarian standar (modul 27).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Indeks pencarian worklog di modul 12/27.

## Notifikasi, API, webhook, otomasi & integrasi

### TIME-039 Work Logged Notification Events

- **Deskripsi:** System event worklog yang dipakai notification scheme untuk mengirim email dan notifikasi in-app.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek (notification scheme).
- **Aturan perilaku (acceptance criteria):**
  1. Event "Work Logged On Issue" diterbitkan saat worklog dibuat.
  2. Event "Issue Worklog Updated" diterbitkan saat worklog diubah.
  3. Event "Issue Worklog Deleted" diterbitkan saat worklog dihapus.
  4. Company-managed: penerima tiap event diatur di notification scheme (Current Assignee, Reporter, Watchers, project role, group, user, dll).
  5. ⚠ Team-managed: event worklog dapat diaktifkan/dimatikan dan dipilih penerimanya di pengaturan notifikasi project, seperti event lain.
  6. Email dan notifikasi in-app menampilkan durasi worklog, Date started, dan deskripsi worklog; worklog dengan visibility terbatas hanya dikirim ke anggota target.
  7. Parameter API `notifyUsers=false` menekan notifikasi; hanya diizinkan untuk pemegang Administer global atau Administer Projects, selain itu HTTP 403.
  8. Preferensi personal "Beri tahu saya tentang perubahan saya sendiri" berlaku untuk event worklog.
- **Team-managed vs company-managed:** Company-managed lewat notification scheme; team-managed lewat pengaturan notifikasi project.
- **Catatan & ketergantungan:** Notifikasi di modul 19; event domain di modul 28.

### TIME-040 Worklog REST API

- **Deskripsi:** Endpoint REST untuk CRUD worklog dan sinkronisasi worklog inkremental ke sistem eksternal. Detail autentikasi dan konvensi API umum ada di modul 23.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, ekstensi.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /api/v1/issues/{issueIdOrKey}/worklogs` mendukung `startAt`, `maxResults` (default 50, maks 5.000), `startedAfter`, `startedBefore` (epoch ms), dan `expand=properties`.
  2. `GET /api/v1/issues/{key}/worklogs/{id}` mengembalikan satu worklog; ID yang tidak ada atau tidak terlihat mengembalikan HTTP 404.
  3. `POST /api/v1/issues/{key}/worklogs` menerima `timeSpent` (string durasi) atau `timeSpentSeconds`, `started` (ISO 8601 dengan offset), `comment` (format dokumen rich text; string plain text diterima dan dikonversi), `visibility`, dan `properties`.
  4. `PUT /api/v1/issues/{key}/worklogs/{id}` dan `DELETE /api/v1/issues/{key}/worklogs/{id}` menerima `adjustEstimate`, `newEstimate`, `reduceBy`, dan `increaseBy` sesuai TIME-013.
  5. Parameter `notifyUsers` (TIME-039 butir 7) dan `overrideEditableFlag` (TIME-018 butir 9) didukung pada create, update, dan delete.
  6. Author worklog selalu pengguna pemanggil; field author pada payload ditolak HTTP 400 "Author tidak dapat ditentukan".
  7. `POST /api/v1/worklogs/list` mengembalikan worklog untuk daftar ID (maks 1.000 ID per permintaan), hanya yang terlihat oleh pemanggil.
  8. `GET /api/v1/worklogs/updated?since=` dan `/api/v1/worklogs/deleted?since=` mengembalikan ID dan waktu perubahan sejak timestamp, dengan paginasi `nextPage`/`lastPage`.
  9. Worklog properties: `GET`/`PUT`/`DELETE /api/v1/issues/{key}/worklogs/{id}/properties/{propertyKey}` untuk menyimpan JSON milik ekstensi (maks 32 KB per property).
  10. Field `worklog` pada `GET /api/v1/issues/{key}` hanya menyertakan 20 worklog pertama beserta `total`; sisanya diambil lewat endpoint worklog.
  11. Field `timetracking` pada create/edit work item menerima `originalEstimate` dan `remainingEstimate` (TIME-005 butir 15).
  12. Permission sama dengan UI (TIME-018); format durasi tidak valid menghasilkan HTTP 400 dengan pesan TIME-003.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lapisan kompatibilitas path bergaya `/rest/api/3/...` (bila diputuskan) didefinisikan di modul 23.

### TIME-041 Bulk Move / Bulk Delete Worklogs

- **Deskripsi:** Endpoint REST untuk memindahkan atau menghapus banyak worklog sekaligus pada sebuah work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, Admin Situs; memerlukan Delete All Worklogs pada work item sumber dan Work On Issues pada tujuan (move).
- **Aturan perilaku (acceptance criteria):**
  1. `DELETE /api/v1/issues/{issueIdOrKey}/worklogs` dengan body daftar ID worklog dan `adjustEstimate` (`leave` atau `auto`) menghapus semuanya dalam satu transaksi.
  2. `POST /api/v1/issues/{issueIdOrKey}/worklogs/move` memindahkan daftar worklog ke work item tujuan dengan `adjustEstimate` (`leave` atau `auto`) yang diterapkan pada sumber dan tujuan; author, started, dan deskripsi tidak berubah.
  3. ⚠ Maksimal 5.000 ID worklog per permintaan; lebih dari itu ditolak HTTP 400.
  4. ⚠ Operasi bulk tidak menerbitkan notifikasi maupun webhook per worklog; satu entri history dan satu event ringkasan per work item yang terdampak dicatat.
  5. Tidak ada padanan UI untuk operasi ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Modul 23.

### TIME-042 Worklog Webhooks

- **Deskripsi:** Webhook untuk event worklog bagi integrasi eksternal.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, Admin Situs (pendaftaran webhook).
- **Aturan perilaku (acceptance criteria):**
  1. Event `worklog_created` dikirim saat worklog dibuat.
  2. Event `worklog_updated` dikirim saat worklog diubah.
  3. Event `worklog_deleted` dikirim saat worklog dihapus.
  4. Payload berisi objek worklog (`id`, `issueId`, `author`, `updateAuthor`, `timeSpentSeconds`, `started`, `created`, `updated`, `visibility`).
  5. ⚠ Filter JQL pada webhook berlaku untuk event worklog berdasarkan work item induknya.
  6. Event `issue_updated` juga dikirim dengan changelog `timespent`, `timeestimate`, dan `WorklogId`.
  7. ⚠ Event yang sama tersedia untuk ekstensi platform sebagai event `worklog.created`, `worklog.updated`, dan `worklog.deleted` (modul 25).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Webhook di modul 23; katalog event di modul 28.

### TIME-043 Automation: Log Work Action & Time Smart Values

- **Deskripsi:** Dukungan otomasi berbasis aturan untuk mencatat waktu, mengubah estimasi, dan bereaksi terhadap worklog. Detail engine ada di modul 20.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs (pembuat rule); rule berjalan sebagai actor rule (PERM-049).
- **Aturan perilaku (acceptance criteria):**
  1. Action "Log work" mengisi time spent, date started, comment, dan penyesuaian Remaining (TIME-013).
  2. ⚠ Trigger "Work logged" tersedia untuk worklog dibuat, dengan opsi juga untuk worklog diubah dan dihapus.
  3. Action Edit issue dapat menyetel Original estimate dan Remaining estimate (`timetracking`).
  4. Smart values `{{issue.timetracking.originalEstimate}}`, `{{issue.timetracking.remainingEstimate}}`, `{{issue.timetracking.timeSpent}}` beserta varian `*Seconds`.
  5. Smart values `{{worklog.timeSpent}}`, `{{worklog.timeSpentSeconds}}`, `{{worklog.author}}`, `{{worklog.comment}}`, `{{worklog.started}}`.
  6. ⚠ Lookup issues mendukung penjumlahan seperti `{{lookupIssues.Story Points.sum}}` dan `{{lookupIssues.timetracking.timeSpentSeconds.sum}}`.
  7. Template rule bawaan "Jumlahkan story points/estimasi sub-task ke parent" memakai math expression.
  8. Condition dapat memeriksa field estimasi (mis. Story Points kosong, Remaining > 0).
  9. Setiap eksekusi rule yang mengubah worklog/estimasi tercatat di audit log rule beserta work item terdampak.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Modul 20; actor rule di TIME-020 butir 5.

### TIME-044 Smart Commits #time

- **Deskripsi:** Mencatat waktu dari perintah `#time` di pesan commit yang diterima lewat integrasi development tool.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Work On Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Sintaks `ISSUE-123 #time 1w 2d 4h 30m Komentar` membuat worklog dengan durasi dan deskripsi tersebut.
  2. Email committer harus cocok dengan email pengguna aktif yang memiliki Work On Issues pada work item; bila tidak, perintah `#time` tidak dijalankan (tidak ada worklog dibuat) dan kegagalan dilaporkan sesuai butir 5.
  3. `#time` dapat digabung dengan `#comment` dan `#<transisi>` dalam satu pesan.
  4. Remaining Estimate disesuaikan dengan mode `auto`.
  5. ⚠ Kegagalan (format durasi salah, pengguna tidak cocok, izin kurang) diberitahukan lewat email ke committer dengan alasan per perintah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Integrasi DVCS di modul 24.

### TIME-045 CSV/External Import of Estimates & Worklogs

- **Deskripsi:** Impor nilai estimasi dan worklog dari CSV.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Kolom CSV dapat dipetakan ke Original Estimate, Remaining Estimate, dan Time Spent dalam detik.
  2. Kolom Work log berformat `comment;date;author;seconds`, atau hanya `seconds`.
  3. Satu work item dapat memiliki banyak kolom Work log; setiap kolom menjadi satu worklog.
  4. Mapping user menentukan author worklog dari nilai CSV; ini satu-satunya cara worklog tercatat atas nama pengguna lain.
  5. Kolom Story Points dapat diimpor ke field numerik estimasi.
  6. Tanggal worklog divalidasi sesuai format tanggal yang diatur di wizard impor; baris gagal dilaporkan per baris tanpa menggagalkan seluruh impor.
  7. Hanya Admin Situs yang dapat memakai impor eksternal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Wizard impor di modul 25.

### TIME-046 Time Tracking Provider Module (Forge/Connect)

- **Deskripsi:** Extension point platform yang memungkinkan ekstensi mendaftarkan time tracking provider sendiri untuk menggantikan UI log work bawaan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pembuat ekstensi), Admin Situs (memilih provider).
- **Aturan perilaku (acceptance criteria):**
  1. Manifest ekstensi produk ini menyediakan modul `timeTrackingProvider` (padanan `jira:timeTrackingProvider` / `jiraTimeTrackingProviders`) berisi `key`, `name`, dan URL/entry UI Log work.
  2. Provider terdaftar muncul di dropdown Time tracking provider (TIME-001 butir 2).
  3. Saat provider ekstensi aktif, aksi Log work di work item view memanggil UI milik ekstensi dalam iframe/panel terisolasi dengan konteks work item.
  4. Field bawaan dan Worklog REST API tetap tersedia bagi ekstensi untuk menyimpan hasil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kerangka ekstensi di modul 25. Aplikasi timesheet pihak ketiga yang memakai modul ini di luar cakupan.

## Perencanaan lanjut & mobile

### TIME-047 Plans: Estimation Unit & Roll-up

- **Deskripsi:** Pengaturan estimasi di Plans (perencanaan lintas proyek) dan roll-up estimasi lintas hierarki work item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Product Owner/PM sebagai planner (PERM-052), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Plan settings menyediakan estimation unit Story points, Days, atau Hours.
  2. ⚠ Plan memiliki pengaturan hours per day sendiri (default mengikuti setting global TIME-002) yang hanya dipakai untuk konversi days/hours di plan tersebut.
  3. Estimasi child di-roll-up ke parent (Epic, initiative, dan level hierarki kustom) dan ditampilkan sebagai total di samping estimasi milik parent.
  4. Kolom Estimate di tabel plan dapat diedit inline; perubahan berstatus draft sampai pengguna menjalankan Review changes > Save.
  5. ⚠ Estimasi berbasis waktu memakai Original Estimate saat belum ada worklog dan Remaining Estimate setelah ada worklog; edit dari plan mengubah field yang sama.
  6. Kolom Progress (estimate) ditampilkan berdampingan dengan Progress (issue count).
  7. ⚠ Work item tanpa estimasi dijadwalkan dengan default estimate per plan (dapat diatur, default 0) dan ditandai di tabel.
  8. Field story points diambil dari estimation field board sumber plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan; plan dapat memuat keduanya.
- **Catatan & ketergantungan:** Plans di modul 18.

### TIME-048 Plans: Team Capacity

- **Deskripsi:** Pengelolaan kapasitas tim di Plans berdasarkan velocity (Scrum) atau jam per minggu (Kanban), beserta visualisasi overbooking.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Product Owner/PM sebagai planner, Scrum Master/Team Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Tim Scrum memiliki velocity per sprint dalam story points atau jam/hari sesuai unit plan; default diisi dari nilai velocity board yang diekspos RPT-012 butir 9 (TIME-028).
  2. Tim Kanban memiliki iteration length dan weekly hours.
  3. Kapasitas per sprint/iterasi dapat diubah satu per satu (mis. karena libur).
  4. Capacity bar per sprint di timeline berwarna hijau bila dalam kapasitas dan merah bila overbooked.
  5. Tooltip capacity bar menampilkan total estimasi terjadwal vs kapasitas.
  6. ⚠ Plans menyediakan tampilan "Capacity" untuk mengelola kapasitas sprint per tim.
  7. Auto-scheduler memakai kapasitas dan estimasi untuk menjadwalkan work item.
  8. Kapasitas dimiliki entitas Team bersama (modul 18) atau plan-only team.
  9. Work item tanpa estimasi tidak mengonsumsi kapasitas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Teams dan Plans di modul 18.

### TIME-049 Mobile Time Tracking

- **Deskripsi:** Log work dan edit estimasi dari perangkat mobile, lewat SPA responsif (dan aplikasi mobile bila kelak dibangun).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, dengan izin yang sama seperti web.
- **Aturan perilaku (acceptance criteria):**
  1. Log work dapat dilakukan dari work item view mobile dengan input durasi dan date picker yang ramah sentuh.
  2. Original dan Remaining Estimate dapat diedit dari mobile.
  3. Worklog dapat dilihat di activity work item mobile.
  4. ⚠ Worklog dapat diedit dan dihapus dari mobile sesuai permission own/all.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Strategi mobile/responsif di modul 27.

## Pertanyaan terbuka

- Apakah opsi "Copy comment to work description" (TIME-002 butir 6, TIME-014 butir 4) dibangun? Inventaris memasukkannya sebagai fitur Cloud, sedangkan CFG-058 butir 7 menyatakannya tidak dibangun; salah satu dokumen perlu diselaraskan.
- Berapa batas jumlah worklog per work item (TIME-021 butir 1): 5.000 seperti CMT-024 atau 10.000 seperti produk referensi?
- Apakah Total Time Spent yang ditampilkan kepada pengguna tanpa akses ke worklog terbatas tetap mencakup durasi worklog tersebut (TIME-019 butir 4, TIME-037 butir 10), atau dikurangi demi kerahasiaan?
- Apakah produk ini menyediakan izin khusus atau peran admin untuk mencatat worklog atas nama pengguna lain (mis. untuk koreksi timesheet), atau tetap hanya lewat importer (TIME-021 butir 9)?
- Apakah Story Points dan Story point estimate dikonsolidasikan menjadi satu field di UI sejak awal (TIME-008 butir 9)?
- Apakah dibutuhkan laporan timesheet native (worklog per pengguna per hari) atau kebutuhan tersebut diserahkan ke ekosistem ekstensi?
- Apakah aplikasi mobile native termasuk peta jalan produk, atau pengalaman mobile cukup lewat SPA responsif (TIME-049)?
- Apakah lapisan kompatibilitas path API bergaya `/rest/api/3/...` disediakan untuk endpoint worklog agar integrasi yang ada dapat dipakai tanpa perubahan (TIME-040)?

## Di luar cakupan & ekosistem

- Legacy Mode Time Tracking (TIME-004) hanya ada di edisi Data Center dan tidak dibangun; perilaku standar (Original dan Remaining dapat diedit independen) berlaku.
- Aplikasi timesheet, penagihan, dan pelaporan waktu pihak ketiga dari Marketplace (mis. Tempo Timesheets, Clockwork, Timesheet Reports, Everhour); produk ini hanya menyediakan extension point provider (TIME-046), Worklog REST API (TIME-040), dan webhook (TIME-042) yang dapat dipakai ekstensi semacam itu.
- App Marketplace untuk planning poker, estimasi kolaboratif, dan kapasitas sprint lanjutan di luar Plans.
- Fitur tagihan/biaya (rate per jam, cost tracking, invoicing) dan persetujuan timesheet yang tidak ada di produk referensi secara native.
- Layanan lintas produk vendor referensi (Atlassian Teams sebagai layanan platform, Atlassian Analytics, Rovo/AI) yang memakai data waktu; produk ini memakai entitas Team sendiri di modul 18.
- Kerangka laporan, dashboard, otomasi, webhook, dan integrasi DVCS secara umum didefinisikan di modul 16, 17, 20, 23, dan 24; modul ini hanya mendefinisikan aspek waktu dan estimasinya.
