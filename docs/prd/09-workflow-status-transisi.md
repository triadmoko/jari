# 09 — Workflow, Status & Transisi

Modul ini mendefinisikan mesin alur kerja produk ini: entitas status dan tiga status category tetap, transition (initial, directed, global, looped) beserta urutan eksekusinya, field Resolution dan semantik reopen, cara pengguna menjalankan transition dari issue view, board, list view, bulk change, dan REST, serta seluruh lapisan konfigurasi workflow. Untuk project company-managed, cakupannya meliputi workflow global yang bisa dibagi, editor workflow visual, rules (conditions, validators, post functions), transition screen, workflow properties, workflow scheme, draft, dan migrasi status. Untuk project team-managed, cakupannya meliputi editor workflow sederhana per work type dan rules berbasis form. Modul ini juga mencakup izin, event, audit, terjemahan, batasan sistem, permukaan REST API domain workflow, titik ekstensi, dan agent AI pada transition. Kondisi repo saat ini baru memiliki fitur auth dan user (`internal/features/auth`, `internal/features/user`, `entity.User`); belum ada tabel status, workflow, transition, resolution, maupun scheme. Seluruh isi modul ini adalah pekerjaan baru yang dibangun sebagai feature module baru (mis. `internal/features/workflow/`) dengan entitas bersama di `internal/shared/domain/entity/` dan skema tabel yang ditetapkan di modul 28.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission Transition/Resolve/Close issues, Assignable user, Administer projects, Administer Jira, Make bulk changes, override izin per status.
- [04 Proyek](04-proyek.md) — model company-managed vs team-managed, Project settings > Workflows, extended project administration, template project.
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — work type, screen, screen scheme, field configuration, daftar nilai Resolution global.
- [06 Work Item Inti](06-work-item-inti.md) — issue view, create, move, change work type, konversi subtask.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — tab History dan activity feed.
- [10 Board](10-board-kolom-kartu.md) — kolom board, pemetaan status, drag-drop kartu.
- [11 Backlog & Sprint](11-backlog-sprint.md) — Complete sprint dan definisi selesai.
- [12 JQL](12-jql-bahasa-kueri.md) — `status`, `statusCategory`, `resolution`, operator historis `WAS`/`CHANGED`.
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — bulk change dari navigator.
- [16 Laporan & Insight](16-laporan-insight.md) — control chart, CFD, Created vs Resolved, Resolution time.
- [19 Notifikasi & Email](19-notifikasi-email.md) — notification scheme dan event.
- [20 Otomasi](20-otomasi.md) — trigger "Work item transitioned" dan action "Transition work item".
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log, Events, WebHooks.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — autentikasi, rate limit, webhook `jira:issue_updated`.
- [24 Integrasi DevOps](24-integrasi-devops.md) — workflow triggers dan smart commits.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — titik ekstensi app dan impor.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — approvals pada status workflow.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — background task, optimistic locking, realtime.
- [28 Model Data Inti](28-model-data-inti.md) — tabel status, workflow, transition, rule, property, scheme, layout, dan event domain.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) dan [32 JSM ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — workflow service project, condition customer/agent, approvals JSM.

## Persona terkait

- **Admin Situs** — mengelola status global, resolution, workflow company-managed, workflow scheme, rules, properties, migrasi status, terjemahan, dan audit.
- **Admin Proyek/Project Lead** — mengedit workflow team-managed per work type, melihat atau (lewat extended project administration) mengedit workflow company-managed yang tidak dibagi, memasang rules team-managed.
- **Scrum Master/Team Lead** — memetakan status ke kolom board, mengandalkan kategori Done dan kolom paling kanan untuk Complete sprint, membaca cycle time.
- **Developer/Anggota Tim** — menjalankan transition dari issue view, board, list view, dan bulk change; mengisi transition screen.
- **Product Owner/PM** — membaca progres epic dan rilis berdasarkan status category serta resolution.
- **Stakeholder/Viewer** — melihat status dan diagram workflow read-only tanpa bisa mengubahnya.
- **Pengguna Tim Bisnis** — memakai workflow sederhana pada project bisnis dan approvals di status.
- **Integrator/Developer API** — menjalankan transition lewat REST, mengelola workflow secara massal, membangun rule ekstensi.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin memindahkan work item ke status berikutnya langsung dari tombol status di issue view agar progres terlihat tanpa membuka halaman lain (WF-011, WF-004).
- Sebagai Scrum Master/Team Lead, saya ingin menyeret kartu antar kolom board dan melihat kolom mana yang tidak menerima drop agar tim tidak melanggar alur kerja (WF-012, WF-006).
- Sebagai Admin Proyek/Project Lead, saya ingin menambah status "In Review" dan mengatur transisinya di project team-managed tanpa bantuan Admin Situs agar workflow sesuai cara kerja tim (WF-017, WF-018).
- Sebagai Admin Situs, saya ingin mendesain workflow company-managed di editor visual lalu membagikannya ke banyak project lewat workflow scheme agar proses seragam di seluruh organisasi (WF-021, WF-033, WF-035).
- Sebagai Admin Situs, saya ingin mewajibkan field tertentu, membatasi siapa yang boleh menutup work item, dan mengisi field otomatis saat transition agar kualitas data terjaga (WF-027, WF-028, WF-029, WF-030).
- Sebagai Admin Situs, saya ingin mengganti workflow sebuah project dengan wizard pemetaan status agar tidak ada work item yang hilang statusnya (WF-036, WF-034, WF-022).
- Sebagai Product Owner/PM, saya ingin resolution terisi dan dikosongkan secara konsisten saat work item selesai atau dibuka kembali agar laporan Created vs Resolved akurat (WF-009, WF-010).
- Sebagai Stakeholder/Viewer, saya ingin melihat diagram workflow sebuah work item dan posisi status saat ini agar saya paham langkah berikutnya (WF-015).
- Sebagai Scrum Master/Team Lead, saya ingin riwayat setiap perubahan status tercatat dengan pelaku dan waktunya agar cycle time dan time-in-status bisa dihitung (WF-014, WF-016).
- Sebagai Developer/Anggota Tim, saya ingin menutup puluhan work item sekaligus dengan satu isian resolution agar pembersihan backlog cepat (WF-013).
- Sebagai Integrator/Developer API, saya ingin mengambil dan menjalankan transition lewat REST dengan ID stabil agar integrasi CI dan otomasi tidak rusak saat nama status berubah (WF-041, WF-001, WF-004).
- Sebagai Admin Situs, saya ingin setiap perubahan status, workflow, dan scheme tercatat di audit log agar perubahan konfigurasi bisa ditelusuri (WF-039, WF-003).
- Sebagai Admin Proyek/Project Lead, saya ingin mengedit workflow company-managed milik project saya sendiri dalam batas aman agar tidak selalu menunggu Admin Situs (WF-038, WF-037).
- Sebagai Admin Situs, saya ingin tim team-managed memakai status global standar organisasi agar laporan dan JQL lintas space konsisten tanpa mengambil alih kendali tim atas transition dan board (WF-046, WF-001).
- Sebagai Admin Situs, saya ingin sistem menolak workflow yang tidak valid (tanpa initial transition, status ganda, versi kedaluwarsa) agar konfigurasi tidak pernah merusak work item (WF-025).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 9 |
| P1 | 22 |
| P2 | 10 |
| P3 | 5 |
| **Total** | **46** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 10 |
| Fase 2 | 20 |
| Fase 3 | 11 |
| Fase 4 | 3 |
| Fase — (dc-only-or-deprecated) | 2 |
| **Total** | **46** |

Catatan: inventaris memuat 46 fitur; dua fitur dengan cakupan sama (kontrol status di issue view dari domain issues-core dan dropdown status dari domain workflows-statuses) digabung menjadi WF-011 dengan baris alias. WF-046 (status global di workflow team-managed) ditambahkan di luar inventaris dengan tanda ⚠ berdasarkan pengumuman produk Oktober 2026.

Strategi implementasi: Fase 1 menetapkan model data yang final sejak awal — status (dengan scope global atau project), status category tetap, workflow berversi, transition dengan himpunan status asal, tabel rule dan property (walau belum dipakai), serta workflow scheme default — karena work item, board, JQL, dan laporan semuanya merujuk ke ID status. Pada Fase 1, project team-managed memakai workflow To Do/In Progress/Done yang di-seed dari template dengan global transitions, dan project company-managed memakai workflow sistem bawaan lewat Default Workflow Scheme; eksekusi transition (issue view, board, permission Transition issues, resolution otomatis, event) sudah lengkap. Fase 2 membuka editor team-managed, editor visual company-managed, rules, transition screen, workflow scheme, draft, dan wizard migrasi status, karena semuanya bergantung pada screen dan field configuration dari modul 05. Fase 3 menambahkan fitur skala organisasi (properties, extended project administration, scheme draft, copy dan penghapusan workflow inactive, self-transition, urutan transition, simplified workflow, status global di team-managed, audit, REST publik). Fase 4 menambahkan titik ekstensi app, terjemahan, dan agent AI; dua fitur legacy Data Center hanya dicatat sebagai rujukan.

## Fitur

## Status & status category

### WF-001 Statuses (global status entity)

- **Deskripsi:** Status adalah entitas yang menunjukkan posisi work item di workflow (mis. To Do, In Progress, Done). Di company-managed, status bersifat global dan dipakai bersama lintas workflow; di team-managed, status dibuat per project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (global permission Administer Jira) mengelola status global; Admin Proyek (Administer projects) mengelola status project team-managed; semua pengguna dengan Browse projects melihat status.
- **Aturan perilaku (acceptance criteria):**
  1. Status memiliki atribut `id` (numerik, immutable), `name`, `description`, `status_category` (wajib), dan `scope` (`GLOBAL` untuk company-managed, `PROJECT` dengan `scope_project_id` untuk team-managed); membuat status tanpa kategori ditolak dengan galat validasi.
  2. Warna lozenge status selalu diturunkan dari status category: abu-abu untuk To Do, biru untuk In Progress, hijau untuk Done; API dan UI tidak menyediakan warna custom per status.
  3. ⚠ Nama status wajib diisi (setelah trim) dan dibatasi panjangnya (default 60 karakter); nama kosong atau melebihi batas ditolak dengan pesan galat yang menyebut batasnya.
  4. ⚠ Company-managed: nama status global unik tanpa membedakan huruf besar/kecil; membuat atau me-rename menjadi nama yang sudah ada ditolak dengan pesan "Status dengan nama ini sudah ada".
  5. Team-managed: status di-scope ke project; nama yang sama boleh ada di project berbeda dan masing-masing menjadi entri terpisah dengan `id` berbeda yang muncul terpisah di JQL dan filter.
  6. Seed sistem dan template menyediakan status bawaan antara lain Open, In Progress, Reopened, Resolved, Closed, To Do, Done, Backlog, Selected for Development, dan In Review, masing-masing dengan kategori yang sesuai.
  7. Field Status tidak bisa diubah lewat edit field (inline edit, edit dialog, bulk edit, atau `PUT /issue`); perubahan status hanya terjadi lewat eksekusi transition atau migrasi/move yang dijalankan sistem.
  8. Status yang masih dipakai workflow apa pun, termasuk workflow inactive, tidak bisa dihapus; API mengembalikan 409 beserta daftar workflow yang memakainya.
  9. Rename status langsung tercermin di semua workflow, board, filter, dan riwayat work item, karena semua relasi menyimpan `status_id`, bukan nama.
  10. Status tampil sebagai field sistem "Status" di issue view, list/navigator, JQL (`status`, `statusCategory`), board, dan laporan.
  11. ID status stabil dipakai di REST API, JQL (mis. `status = 10001`), dan workflow properties; ID tidak pernah dipakai ulang setelah status dihapus.
  12. Untuk kompatibilitas, respons REST status menyertakan `iconUrl` legacy (statusIconUrl) yang diisi nilai default; UI produk ini tidak memakainya (P3).
  13. ⚠ Team-managed juga dapat memakai status `GLOBAL` dari katalog Admin Situs; work item company-managed dan team-managed yang memakai status global yang sama memiliki `status_id` yang sama sehingga JQL `status = X` mencakup keduanya (WF-046).
- **Team-managed vs company-managed:** Company-managed memakai status global yang unik per site dan dikelola Admin Situs (WF-003). Team-managed memakai status yang di-scope ke project dan dikelola dari editor workflow project (WF-017), dan ⚠ dapat pula memakai status global (WF-046).
- **Catatan & ketergantungan:** Model data di modul 28: tabel `statuses` dengan `scope_project_id` nullable (null = global). Operator JQL di [12 JQL](12-jql-bahasa-kueri.md). Repo saat ini belum punya entitas status; perlu migration baru dan seed status bawaan.

### WF-002 Status Categories

- **Deskripsi:** Tiga kategori tetap (To Do, In Progress, Done) yang mengelompokkan status. Kategori menentukan warna, perilaku board, laporan, Timeline, dan JQL `statusCategory`.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem (mendefinisikan kategori); Admin Situs dan Admin Proyek memilih kategori saat membuat atau mengedit status.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia tepat tiga kategori tetap: To Do (key `new`), In Progress (`indeterminate`), dan Done (`done`); tidak ada UI atau API untuk membuat, mengubah, atau menghapus kategori.
  2. Kategori legacy "No Category" (key `undefined`) hanya ada untuk kompatibilitas data impor yang statusnya tidak berkategori (P3); kategori ini tidak bisa dipilih saat membuat status baru.
  3. Kategori menentukan warna lozenge status di seluruh UI (issue view, kartu board, list, hasil pencarian, picker).
  4. Kategori Done dipakai untuk menghitung progres epic/parent di Timeline, jumlah work item selesai di Releases, dan sprint report.
  5. JQL mendukung `statusCategory = "To Do" | "In Progress" | Done` (nama atau key) serta field `statusCategoryChangedDate`.
  6. Field sistem "Status Category Changed" (`statusCategoryChangedDate`) diperbarui hanya saat kategori status work item berubah; transition antar dua status berkategori sama tidak mengubahnya.
  7. Team-managed: memindahkan work item ke status berkategori Done otomatis mengisi resolution "Done", dan memindahkannya keluar dari kategori Done mengosongkan resolution.
  8. Mengubah kategori sebuah status langsung memengaruhi semua workflow, board, dan laporan yang memakainya; UI menampilkan peringatan jumlah workflow terdampak sebelum disimpan.
  9. Board company-managed: saat Complete sprint, work item di kolom paling kanan dianggap selesai terlepas dari kategorinya, sedangkan kategori Done tetap dipakai untuk progres Timeline/epic.
- **Team-managed vs company-managed:** Team-managed mengaitkan kategori Done dengan resolution otomatis (butir 7); company-managed tidak, resolution diatur lewat transition screen atau post function (WF-009).
- **Catatan & ketergantungan:** Complete sprint di [11 Backlog & Sprint](11-backlog-sprint.md); Timeline di modul 18; laporan di [16 Laporan](16-laporan-insight.md). Kategori di-seed lewat migration dan dirujuk dengan key tetap.

### WF-003 Status administration page (Admin > Work items > Statuses)

- **Deskripsi:** Halaman admin global untuk melihat, membuat, mengedit, menerjemahkan, dan menghapus status company-managed beserta informasi pemakaiannya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (global permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar status menampilkan nama, deskripsi, kategori (lozenge), dan workflow yang memakainya.
  2. Daftar bisa dicari berdasarkan nama (substring, tidak peka huruf besar/kecil) dan difilter berdasarkan kategori.
  3. Tombol "Add status" membuka form nama, deskripsi, dan kategori; status baru bersifat global dan langsung tersedia di semua editor workflow company-managed.
  4. Edit status mengizinkan perubahan nama, deskripsi, dan kategori, dengan validasi yang sama seperti WF-001.
  5. Hapus status hanya aktif bila status tidak dipakai workflow mana pun, termasuk workflow inactive; status yang tidak dipakai ditandai "Tidak dipakai" sebagai kandidat cleanup, dan tombol hapus pada status terpakai dinonaktifkan dengan tooltip alasan.
  6. Setiap baris menyediakan aksi "Translate" untuk menerjemahkan nama dan deskripsi per bahasa (lihat WF-040).
  7. ⚠ Panel pemakaian menampilkan workflow yang memakai status serta project dan work type terkait (setara REST `statuses/{id}/projectIssueTypeUsages` dan `workflowUsages`); jumlah work item per status ditampilkan bila dapat dihitung tanpa membebani kinerja.
  8. Status team-managed tidak muncul di halaman ini; status tersebut dikelola dari editor workflow project (WF-017).
  9. Halaman dan endpoint-nya hanya bisa diakses pengguna dengan Administer Jira; pengguna lain menerima 403.
  10. Setiap create, update, delete, dan translate status dicatat di audit log (WF-039).
- **Team-managed vs company-managed:** Hanya untuk status company-managed (global).
- **Catatan & ketergantungan:** Bergantung pada WF-001; audit log di [22 Administrasi Sistem](22-administrasi-sistem.md).

## Transition

### WF-004 Transitions

- **Deskripsi:** Transition adalah jalur satu arah yang memindahkan work item dari satu status ke status lain, dan bisa dilengkapi screen, rules, serta properties.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs dan Admin Proyek (team-managed) mendefinisikan transition; pengguna dengan permission Transition issues menjalankannya.
- **Aturan perilaku (acceptance criteria):**
  1. Transition memiliki atribut `id`, `name` (label tombol), `description`, `from` (satu atau lebih status asal, atau kosong untuk global), dan `to` (satu status tujuan).
  2. Transition bersifat satu arah; perpindahan bolak-balik antara dua status membutuhkan dua transition terpisah.
  3. Nama transition boleh berbeda dari nama status tujuan (mis. "Start progress" menuju In Progress); bila nama tidak diisi, default-nya nama status tujuan.
  4. Common/shared transition didukung: satu transition dengan beberapa status asal (multi-from) memakai satu set rules, screen, dan properties yang sama.
  5. Dua transition bernama sama dari status asal berbeda tetap entitas terpisah dengan `id` dan rules masing-masing, kecuali dibuat sebagai shared transition.
  6. ⚠ Transition bisa dihapus dari editor; bila transition dirujuk rule otomasi atau smart commit, editor menampilkan peringatan sebelum penghapusan dikonfirmasi.
  7. Tipe transition yang didukung: `initial`, `directed` (dari status tertentu), `global` (dari Any status), dan `looped`/self.
  8. Transition bisa dijalankan dari issue view, drag-drop board, inline status di list view, bulk change, REST API, action otomasi "Transition work item", smart commits, dan workflow triggers; semua jalur memakai satu service domain yang sama.
  9. ID transition unik dalam satu workflow dan stabil antar versi workflow; ID ini yang dipakai `POST /issue/{key}/transitions`.
  10. Urutan eksekusi tetap: cek conditions, tampilkan screen, jalankan validators, ubah status, jalankan post functions, fire event, lalu tulis history.
  11. Eksekusi bersifat atomik dalam satu transaksi database: bila condition atau validator gagal, pesan galat ditampilkan dan work item tetap di status asal tanpa perubahan field apa pun.
- **Team-managed vs company-managed:** Team-managed default-nya memakai global transitions sehingga setiap status bisa menuju status mana pun, dan transition manual bisa ditambahkan (WF-017). Company-managed mendukung semua tipe beserta rules penuh (WF-027 s.d. WF-031).
- **Catatan & ketergantungan:** Model `workflow_transition` (from-set, to, type) di modul 28; action otomasi di [20 Otomasi](20-otomasi.md); smart commits dan workflow triggers di [24 Integrasi DevOps](24-integrasi-devops.md).

### WF-005 Initial transition (Create)

- **Deskripsi:** Transition khusus "Create" yang dijalankan saat work item dibuat dan menentukan status awalnya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; Sistem menjalankan saat create (hak membuat diatur permission Create issues).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap workflow wajib memiliki tepat satu initial transition menuju initial status; workflow tanpa initial transition, atau dengan lebih dari satu, ditolak saat disimpan.
  2. Initial status bisa diubah di editor; perubahan hanya berlaku untuk work item yang dibuat setelah publish.
  3. Initial transition mendukung validators (mis. field wajib saat create) dan post functions (mis. assign, update field).
  4. Initial transition tidak mendukung conditions; rule picker tidak menampilkan kelompok Restrict transition untuk transition ini, karena hak membuat diatur permission Create issues.
  5. Initial transition tidak punya transition screen; field yang tampil saat create berasal dari create screen di screen scheme.
  6. Post functions wajib bawaan pada initial transition adalah "Create the issue originally", "Fire Issue Created event", dan "Re-index"; ketiganya tidak bisa dihapus atau diurutkan ulang.
  7. ⚠ Team-managed: workflow memiliki penanda status awal yang bisa dipindahkan admin project ke status mana pun di workflow tersebut.
  8. ⚠ Inline create di kolom board: di team-managed, work item yang dibuat di kolom tertentu langsung berada di status kolom itu; di company-managed, inline create hanya tersedia di kolom yang memuat initial status.
- **Team-managed vs company-managed:** Lihat butir 7 dan 8; rules initial hanya ada di company-managed.
- **Catatan & ketergantungan:** Create dialog dan inline create di [06 Work Item Inti](06-work-item-inti.md); create screen di [05 Konfigurasi](05-konfigurasi-work-item-skema.md).

### WF-006 Global transitions ('Allow all statuses to transition to this one')

- **Deskripsi:** Transition dari "Any status" ke status tertentu, sehingga work item bisa dipindahkan ke status itu dari status mana pun.
- **Prioritas / Fase / Tier:** P1 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Global transition dibuat dengan mencentang "Allow all statuses to transition to this one" pada status, atau dengan Add transition dengan `from = Any status`; kedua cara menghasilkan entitas yang sama.
  2. Diagram menandai status tujuan dengan penanda "All" (Any status).
  3. Nama default global transition sama dengan nama status tujuan dan bisa diubah.
  4. Global transition bisa diberi conditions, validators, post functions, screen, dan properties seperti transition biasa.
  5. Global transition tidak ditampilkan untuk work item yang sudah berada di status tujuannya.
  6. Workflow simplified dan workflow default team-managed memakai global transitions secara default, sehingga work item bebas berpindah antar kolom.
  7. Bila global transition dan transition biasa menuju status yang sama sama-sama lolos, keduanya tampil di dropdown status; editor menampilkan peringatan risiko duplikasi nama saat keduanya bernama sama.
- **Team-managed vs company-managed:** Team-managed mengaktifkan toggle ini per status secara default (WF-017); company-managed default-nya nonaktif kecuali di workflow simplified.
- **Catatan & ketergantungan:** Ditempatkan di Fase 1 (pengecualian dari P1) karena workflow default team-managed dan workflow simplified bergantung padanya. Lihat WF-026.

### WF-007 Self / looped transitions

- **Deskripsi:** Transition yang status asal dan tujuannya sama, dipakai untuk menjalankan aksi (screen, post functions) tanpa mengubah status.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; pengguna dengan Transition issues menjalankan.
- **Aturan perilaku (acceptance criteria):**
  1. Self-transition dapat dibuat dari status X ke status X (mis. "Add info", "Escalate") dan tampil di dropdown status work item yang berada di X.
  2. ⚠ Looped/reflexive global transition didukung: dari Any status dengan target "Itself", sehingga satu transition tersedia di setiap status dan mengembalikan work item ke status asalnya masing-masing.
  3. Self-transition tetap menjalankan conditions, screen, validators, dan post functions; status tidak berubah, tetapi perubahan field lain tercatat di history.
  4. Self-transition dengan transition screen dapat dipakai untuk mengedit field tertentu pada status yang read-only (`jira.issue.editable=false`).
  5. Diagram menampilkan looped transition sebagai badge pada status, bukan panah.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Bergantung pada WF-030 dan WF-031. `statusCategoryChangedDate` tidak berubah pada self-transition (WF-002).

### WF-008 Transition ordering & 'All transitions' view

- **Deskripsi:** Mengatur urutan transition di dropdown status dan cara pengguna melihat semua transition yang tersedia.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs mengatur urutan; semua pengguna dengan Transition issues melihat dropdown.
- **Aturan perilaku (acceptance criteria):**
  1. Tanpa konfigurasi, urutan transition di dropdown mengikuti urutan pembuatan (ID transition menaik).
  2. Property `opsbar-sequence` pada transition mengatur urutan: angka lebih kecil tampil lebih dulu; transition tanpa property ditaruh setelah yang memilikinya; nilai non-angka diabaikan.
  3. Dropdown status di issue view menampilkan semua transition yang lolos conditions dengan format "nama transition → status tujuan".
  4. Dropdown menyediakan opsi "View workflow" yang membuka diagram workflow work item (WF-015).
  5. Transition yang gagal condition tidak ditampilkan sama sekali (bukan dinonaktifkan).
  6. Daftar transition yang panjang ditampilkan sebagai list yang bisa di-scroll, dengan kotak pencarian di bagian atas.
  7. ⚠ Command palette menyediakan aksi "Transition: <nama>" untuk work item yang sedang dibuka, dan shortcut keyboard membuka dropdown status.
  8. Board: bila sebuah kartu bisa di-drop ke beberapa status atau transition di satu kolom, board menampilkan drop zone per status (lihat WF-012).
- **Team-managed vs company-managed:** `opsbar-sequence` hanya di company-managed karena team-managed tidak punya properties; team-managed memakai urutan status di workflow.
- **Catatan & ketergantungan:** Command palette di [21 Navigasi](21-navigasi-ux-preferensi.md). Properties di WF-031.

## Resolution

### WF-009 Resolution & reopen semantics

- **Deskripsi:** Field sistem Resolution menandai alasan work item selesai; resolution yang terisi berarti work item dianggap resolved di seluruh produk, terlepas dari statusnya. Fitur ini juga mengatur perilaku saat work item dibuka kembali (reopen).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Resolve issues mengisi resolution lewat transition; Admin Situs mengonfigurasi; Sistem mengisi otomatis di team-managed dan simplified workflow.
- **Aturan perilaku (acceptance criteria):**
  1. Resolution kosong (ditampilkan "Unresolved") berarti work item masih terbuka; resolution terisi (Done, Won't Do, Duplicate, Cannot Reproduce, dll.) berarti resolved.
  2. Field `resolutiondate` (Resolved) terisi otomatis dengan waktu server saat resolution diset dan dikosongkan saat resolution dihapus; mengganti satu nilai resolution ke nilai lain tidak mengubah `resolutiondate`.
  3. Work item yang resolved ditampilkan dengan key dicoret (strikethrough) di link, hasil pencarian, dan issue picker.
  4. JQL mendukung `resolution = Unresolved`, `resolution is EMPTY`, `resolved >= -7d`, dan `resolution CHANGED`.
  5. Gadget dan laporan Created vs Resolved, Resolution time, Average age, dan Recently created memakai resolution, bukan status.
  6. Company-managed: resolution hanya diset lewat transition screen atau post function Update Issue Field, dan tidak otomatis mengikuti kategori Done.
  7. Company-managed: editor workflow memberi peringatan bila sebuah transition menuju status berkategori Done tanpa mengisi resolution (work item tetap "Unresolved"), atau bila transition keluar dari kategori Done tanpa mengosongkan resolution (work item tetap "resolved").
  8. Team-managed: resolution otomatis diisi "Done" saat work item masuk status berkategori Done dan dikosongkan saat keluar dari kategori tersebut.
  9. Simplified workflow: masuk ke kolom yang diberi opsi "Set resolution" otomatis mengisi resolution (WF-026).
  10. Konfigurasi default sistem tidak memasang field Resolution di create screen maupun edit screen; Admin Situs yang menambahkannya mendapat peringatan bahwa work item bisa langsung resolved.
  11. Property transition `jira.field.resolution.include` / `jira.field.resolution.exclude` (daftar ID resolution) membatasi pilihan resolution di transition screen.
  12. Reopen: transition "Reopen Issue" di workflow sistem disertai post function yang mengosongkan resolution.
  13. Saat reopen, `statusCategoryChangedDate` diperbarui ketika work item keluar dari kategori Done.
  14. Work item yang dibuka kembali setelah sprint ditutup tidak otomatis kembali ke sprint lama; work item tetap tanpa sprint aktif sampai dimasukkan ulang.
  15. History mencatat transition reopen beserta pengguna, waktu, dan perubahan resolution.
  16. Laporan Resolution time dan Created vs Resolved menghitung ulang work item yang dibuka kembali: work item dianggap resolved menurut `resolutiondate` terbaru.
- **Team-managed vs company-managed:** Team-managed mengikat resolution ke kategori Done (butir 8); company-managed mengatur resolution secara eksplisit lewat screen atau post function (butir 6-7).
- **Catatan & ketergantungan:** Daftar nilai resolution di WF-010 dan [05 Konfigurasi](05-konfigurasi-work-item-skema.md) (CFG-055). Pola reopen otomatis via otomasi ada di [20 Otomasi](20-otomasi.md). Laporan di [16 Laporan](16-laporan-insight.md). Butir 11 (`jira.field.resolution.include/exclude`) baru aktif bersama WF-031 di Fase 3; pada Fase 1 transition screen menampilkan semua resolution global.

### WF-010 Resolutions administration

- **Deskripsi:** Halaman admin untuk mengelola daftar nilai resolution global yang dipakai semua project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Seed menyediakan resolution bawaan Done, Won't Do, Duplicate, dan Cannot Reproduce; data impor dari site lama dapat membawa Fixed, Won't Fix, dan Incomplete.
  2. "Add resolution" meminta nama (wajib, unik tanpa membedakan huruf besar/kecil) dan deskripsi.
  3. Resolution bisa diedit dan dihapus; penghapusan resolution yang dipakai work item mewajibkan pemilihan resolution pengganti, lalu work item terdampak diperbarui sebagai background task.
  4. Urutan resolution bisa diubah (move up/down) dan urutan ini dipakai di dropdown transition screen.
  5. Admin dapat menetapkan satu resolution sebagai Default yang otomatis terpilih di transition screen; menghapus default mengosongkan penanda default.
  6. Setiap resolution memiliki aksi "Translate" untuk nama dan deskripsi (WF-040).
  7. Resolution bersifat global untuk semua project; tidak ada resolution scheme, dan pembatasan per transition dilakukan lewat property `jira.field.resolution.include/exclude` (WF-031).
  8. Halaman dan endpoint pengelolaan hanya bisa diakses pengguna dengan Administer Jira.
- **Team-managed vs company-managed:** Team-managed hanya memakai resolution "Done" secara otomatis; daftar global tetap tersedia untuk company-managed.
- **Catatan & ketergantungan:** Tumpang tindih dengan CFG-055 di modul 05: modul 05 memiliki entitas dan nilai resolution sebagai konfigurasi global, modul ini memiliki halaman admin dan perilakunya di workflow. Keputusan pemilik halaman tercatat di Pertanyaan terbuka.

## Menjalankan transition

### WF-011 Transitioning from issue view (status dropdown)

- **Alias inventaris:** Status & Transition Control in Issue View
- **Deskripsi:** Cara pengguna menjalankan transition dari tampilan detail work item melalui tombol status: dropdown menampilkan transition yang tersedia dan menjalankannya, termasuk transition screen bila ada.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan permission Transition issues pada project; tanpa permission itu status tampil read-only.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol lozenge status di issue view membuka dropdown berisi transition yang tersedia untuk work item tersebut.
  2. Dropdown hanya berisi transition yang lolos condition untuk pengguna yang sedang login.
  3. Label setiap item menampilkan nama transition dan status tujuan.
  4. Transition tanpa screen langsung dijalankan dengan optimistic update; bila server menolak, status di UI di-rollback ke nilai semula dan pesan galat ditampilkan.
  5. Transition dengan screen membuka modal dialog berisi field screen (mis. Resolution, komentar, field wajib); menutup dialog tanpa submit tidak mengubah apa pun.
  6. Jika validator gagal, pesan galat tampil di dialog (atau sebagai pesan inline untuk transition tanpa screen) dan transition dibatalkan.
  7. Transition yang gagal condition tidak ditampilkan di dropdown.
  8. Hasil post function (resolution terisi, assignee berubah, field diperbarui) langsung terlihat di issue view tanpa reload halaman.
  9. Dropdown menyediakan tautan "View workflow" yang membuka diagram workflow read-only (WF-015).
  10. Dropdown menyediakan kotak pencarian ketika jumlah transition banyak (default lebih dari 8).
  11. Server memeriksa permission Transition issues pada setiap eksekusi; request tanpa izin ditolak 403 walau dikirim langsung ke API.
  12. Pengguna tanpa Transition issues melihat lozenge status sebagai teks read-only tanpa dropdown.
  13. Team-managed: secara default setiap status bisa berpindah ke status mana pun, kecuali dibatasi rules (WF-018).
  14. Warna lozenge status mengikuti status category (WF-002).
  15. Transition juga tersedia lewat inline status pada List view dan issue navigator dengan aturan condition, screen, dan validator yang sama.
  16. Setiap transition yang berhasil memicu event, notifikasi, webhook `jira:issue_updated`, dan trigger otomasi "Work item transitioned".
  17. Opsi "View workflow" tetap tersedia bagi pengguna yang hanya punya akses baca, selama mereka memiliki Browse projects.
  18. Transition yang sama tidak bisa dijalankan dua kali karena klik ganda: tombol dikunci sampai respons server diterima, dan server menolak eksekusi dari status asal yang sudah tidak berlaku dengan galat "Status work item sudah berubah, muat ulang".
- **Team-managed vs company-managed:** Team-managed default-nya memakai global transitions (semua status ke semua status); company-managed menampilkan transition sesuai desain workflow dan rules.
- **Catatan & ketergantungan:** Definisi workflow ada di bagian lain modul ini; issue view di [06 Work Item Inti](06-work-item-inti.md); event di WF-016; trigger otomasi di [20 Otomasi](20-otomasi.md). Butir 17-18 adalah perincian perilaku dari sub-fitur "View workflow" dan eksekusi atomik (WF-004). Pada Fase 1 butir 16 cukup mem-fire event domain dan notifikasi in-app; webhook (modul 23) dan trigger otomasi (modul 20) mengonsumsi event yang sama saat modul tersebut tersedia di Fase 3.

### WF-012 Board transitions & column-status mapping

- **Deskripsi:** Menjalankan transition dengan drag-drop kartu di board, serta memetakan status ke kolom board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Transition issues (drag-drop); Board administrator dan Admin Proyek (pemetaan kolom); Admin Situs (menambah status global dari board).
- **Aturan perilaku (acceptance criteria):**
  1. Men-drag kartu ke sebuah kolom menjalankan transition ke status yang dipetakan ke kolom tersebut.
  2. Bila satu kolom memetakan beberapa status, atau ada beberapa transition valid ke kolom itu, board menampilkan drop zone terpisah per status/transition saat kartu di-drag.
  3. Kolom yang tidak memiliki transition valid dari status kartu digelapkan selama drag dan tidak menerima drop.
  4. Bila transition memiliki screen, dialog muncul saat drop; error validator ditampilkan sebagai pesan dan kartu kembali ke posisi semula.
  5. Board company-managed: Board settings > Columns memetakan status ke kolom, termasuk area "Unmapped statuses"; aturan pemetaan rinci (satu kolom banyak status, satu status hanya di satu kolom per board) mengikuti BRD-025.
  6. Status yang tidak dipetakan ke kolom mana pun tidak muncul di board, dan work item di status tersebut tidak tampil sebagai kartu.
  7. Menambah status dari Board settings > Columns membuat status baru di workflow; aksi ini hanya untuk Admin Situs, atau board administrator bila project memakai simplified workflow.
  8. Team-managed: kolom board terhubung dengan status milik project dan satu kolom dapat memuat beberapa status lewat halaman Columns and statuses (BRD-025 butir 10); menambah kolom membuat status baru, sedangkan me-rename atau menghapus status mengubah workflow project (WF-017).
  9. Pada simplified workflow, opsi "Set resolution" tersedia per status di Board settings > Columns (bukan per kolom): work item yang masuk status bertanda itu otomatis diberi resolution, dan keluar dari status itu mengosongkan resolution (BRD-025 butir 6, BRD-027 butir 5, WF-026).
- **Team-managed vs company-managed:** Lihat butir 5, 7, dan 8.
- **Catatan & ketergantungan:** Detail board (kolom, swimlane, kartu) ada di [10 Board](10-board-kolom-kartu.md); modul ini hanya mengatur interaksi board dengan workflow. Pemetaan status ke kolom dirinci di BRD-025 (Column status mapping) dan BRD-026 (Unmapped statuses) modul 10; WF-012 dan BRD-025 harus diimplementasikan dengan aturan yang sama. Pemetaan kolom ke status global di team-managed lihat WF-046.

### WF-013 Bulk transition

- **Deskripsi:** Menjalankan transition untuk banyak work item sekaligus dari issue navigator lewat bulk change.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan global permission Make bulk changes dan Transition issues pada setiap project terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Bulk change > Transition issues hanya menawarkan transition yang tersedia untuk work item dengan workflow dan status asal yang sama.
  2. Work item terpilih dikelompokkan berdasarkan workflow dan transition yang tersedia; pengguna memilih satu transition per kelompok.
  3. Transition screen ditampilkan sekali, dan nilai yang diisi diterapkan ke semua work item dalam kelompok.
  4. Validators dijalankan per work item; work item yang gagal dilewati dan dilaporkan beserta pesan galatnya, sementara yang lolos tetap diproses.
  5. Opsi "Send notification" hanya bisa dimatikan oleh pengguna yang berwenang (Admin Proyek/Admin Situs); default-nya aktif.
  6. Operasi dijalankan sebagai background task dengan progress bar dan ringkasan akhir (berhasil/gagal).
  7. ⚠ Satu operasi bulk dibatasi maksimal 1000 work item; pilihan yang melebihi batas ditolak dengan pesan galat yang menyebut batasnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** UI bulk change di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md); permission Make bulk changes di [03 Izin](03-izin-peran-keamanan.md) (PERM-011); background task di [27 NFR](27-nfr-arsitektur-platform.md).

### WF-014 Status history & transition audit on issue

- **Deskripsi:** Pencatatan setiap perubahan status pada work item untuk history, laporan, dan JQL historis.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse projects melihat; Sistem mencatat.
- **Aturan perilaku (acceptance criteria):**
  1. Tab History mencatat setiap transition dengan status asal, status tujuan, pengguna, dan waktu.
  2. ⚠ Activity menyediakan tampilan "Transitions" yang menampilkan lama waktu work item berada di setiap status beserta jumlah kunjungan per status.
  3. JQL historis mendukung `status WAS`, `WAS IN`, `WAS NOT`, serta `status CHANGED FROM/TO/BY/DURING/ON/AFTER/BEFORE`.
  4. Field `resolutiondate`, `statusCategoryChangedDate`, dan `updated` diperbarui konsisten oleh setiap transition.
  5. Data riwayat status menjadi sumber control chart, cumulative flow diagram, cycle time, dan time-in-status.
  6. Perubahan status oleh otomasi atau sistem dicatat dengan pelakunya (mis. "Automation" atau "System"), bukan pengguna yang memicu rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penulisan baris changelog status sudah wajib sejak Fase 1 (bagian dari WF-004 butir 10) agar data historis tidak hilang; tampilan dan JQL historis di Fase 2. Tab History di [08 Aktivitas](08-komentar-lampiran-aktivitas.md); JQL di [12 JQL](12-jql-bahasa-kueri.md); laporan di [16 Laporan](16-laporan-insight.md).

### WF-015 Workflow viewer (read-only diagram)

- **Deskripsi:** Diagram workflow read-only yang bisa dibuka pengguna non-admin untuk memahami alur status work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse projects (dari issue view); Admin Proyek (dari Project settings).
- **Aturan perilaku (acceptance criteria):**
  1. Viewer bisa dibuka dari issue view ("View workflow") dan dari Project settings > Workflows.
  2. Bila dibuka dari issue view, status work item saat ini disorot.
  3. Viewer menampilkan semua status dan transition, termasuk global transitions dengan penanda "All".
  4. ⚠ Pengguna non-admin hanya melihat status, transition, dan nama; detail rules, screen, dan properties hanya ditampilkan untuk pengguna yang berhak mengedit workflow.
  5. Viewer mendukung zoom in/out dan fit-to-screen.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai komponen kanvas yang sama dengan editor (WF-021) dalam mode read-only. Permission View read-only workflow di [03 Izin](03-izin-peran-keamanan.md) (PERM-027).

### WF-016 Workflow-related events

- **Deskripsi:** Event yang di-fire oleh transition dan dipakai notification scheme, webhook, dan otomasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs memilih event di post function; Sistem mem-fire event.
- **Aturan perilaku (acceptance criteria):**
  1. Event sistem tersedia: Issue Created, Issue Resolved, Issue Closed, Issue Reopened, Work Started on Issue, Work Stopped on Issue, Generic Event, dan Issue Updated.
  2. Custom event yang dibuat Admin Situs (Admin > System > Events) bisa dipilih di post function Fire event (WF-029).
  3. Setiap transition mem-fire tepat satu event; bila tidak dikonfigurasi, event default adalah Generic Event.
  4. Event menentukan template email dan penerima menurut notification scheme project.
  5. Webhook `jira:issue_updated` membawa changelog status (from/to), dan trigger otomasi "Work item transitioned" bisa difilter berdasarkan status asal dan tujuan.
- **Team-managed vs company-managed:** Team-managed tidak bisa memilih event per transition; sistem mem-fire event bawaan yang sesuai.
- **Catatan & ketergantungan:** Notification scheme di [19 Notifikasi](19-notifikasi-email.md); manajemen Events di [22 Administrasi Sistem](22-administrasi-sistem.md); webhook di [23 API](23-api-publik-webhook.md). Event in-app dasar (status berubah) sudah dipakai notifikasi Fase 1.

## Workflow team-managed

### WF-017 Team-managed workflow editor (per work type)

- **Deskripsi:** Editor workflow sederhana untuk project team-managed; setiap work type punya workflow sendiri yang dikelola admin project tanpa scheme.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (Administer projects pada project team-managed) dan Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Editor dibuka dari Project settings > Work types > [work type] > Edit workflow, atau dari menu board "Manage workflow".
  2. Setiap work type, termasuk subtask, memiliki workflow sendiri yang tidak dibagi dengan work type atau project lain.
  3. Workflow default untuk work type baru berisi To Do, In Progress, dan Done, dengan semua status saling terhubung lewat global transitions.
  4. "Add status" meminta nama dan kategori; status baru di-scope ke project dan langsung bisa dipakai work type lain di project yang sama.
  5. Setiap status memiliki toggle "Allow all statuses to transition to this one" yang default-nya aktif.
  6. "Add transition" manual meminta status asal (satu, beberapa, atau Any status), status tujuan, dan nama.
  7. Status dan transition bisa di-rename dan dihapus; menghapus status yang masih berisi work item mewajibkan pemilihan status pengganti, lalu work item dipindahkan ke status itu.
  8. Kolom board terhubung dengan status: menambah kolom membuat status baru, status yang ada bisa dipetakan ke kolom, dan satu kolom dapat memuat beberapa status (selaras WF-012 butir 8 dan BRD-025 butir 10).
  9. Tidak ada conditions, validators, atau post functions klasik; pembatasan dan aksi memakai Rules (WF-018).
  10. Tidak ada properties, transition screen custom, maupun workflow triggers DevOps di editor ini.
  11. ⚠ Jumlah status per workflow dibatasi (default 30 status per work type) dan penambahan di atas batas ditolak dengan pesan galat.
  12. Hanya Admin Proyek atau Admin Situs yang bisa membuka editor dalam mode edit; pengguna lain mendapat tampilan read-only (WF-015).
  13. Editor menyediakan tombol Save/Update dan Discard changes; perubahan yang disimpan berlaku langsung tanpa draft.
  14. ⚠ Selain membuat status project, editor dapat mencari dan menambahkan status global dari katalog Admin Situs; status global tidak di-scope ke project dan tidak bisa di-rename dari editor team-managed (WF-046).
- **Team-managed vs company-managed:** Khusus team-managed. Padanan company-managed adalah WF-021 dengan scheme WF-033.
- **Catatan & ketergantungan:** Fase 2 selaras dengan modul [04 Proyek](04-proyek.md) (PRJ-052): pada Fase 1 project team-managed memakai workflow yang di-seed dari template tanpa editor. Kolom board di [10 Board](10-board-kolom-kartu.md).

### WF-018 Team-managed workflow rules

- **Deskripsi:** Rules berbasis form pada transition team-managed untuk membatasi, memvalidasi, dan menjalankan aksi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Restrict transition: rule "Restrict who can move an issue" membatasi transition ke pengguna tertentu, project role tertentu, assignee, atau reporter.
  2. ⚠ Restrict transition: rule "Restrict when all subtasks/child are in a status" hanya mengizinkan transition bila semua subtask/child berada di status yang dipilih.
  3. Validate details: rule "Validate that a field is filled" mewajibkan field tertentu terisi sebelum transition berhasil.
  4. ⚠ Validate details: rule "Validate that a field has changed" mewajibkan field tertentu diubah selama transition.
  5. Perform actions: rule "Assign an issue" meng-assign ke pengguna tertentu, reporter, current user, atau unassigned.
  6. Perform actions: rule "Update an issue field" menetapkan nilai field tertentu.
  7. Satu transition bisa memiliki beberapa rule; rule restrict digabung dengan logika AND.
  8. Rule dikonfigurasi lewat dialog sederhana tanpa properties atau ekspresi.
  9. Rule restrict menyembunyikan transition dari dropdown, dan drag-drop di board ke status terlarang ditolak dengan pesan yang menjelaskan alasannya.
  10. Global transition ("All") juga bisa diberi rules.
- **Team-managed vs company-managed:** Khusus team-managed. Padanan company-managed adalah WF-027, WF-028, dan WF-029.
- **Catatan & ketergantungan:** Implementasi sebaiknya memakai mesin rule yang sama dengan company-managed (tabel `transition_rule` dengan `kind` dan config JSON) dengan katalog rule yang lebih kecil.

## Workflow company-managed: administrasi & editor

### WF-019 Workflows administration page (Admin > Work items > Workflows)

- **Deskripsi:** Halaman daftar semua workflow company-managed yang menampilkan status aktif/inaktif, scheme dan project pemakai tiap workflow, serta aksi pengelolaan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menyediakan filter/tab Active dan Inactive (definisi inactive di WF-024).
  2. Kolom tabel: nama, deskripsi, terakhir diubah (waktu dan pengguna), scheme yang memakai, dan project/work type terkait.
  3. Workflow bisa dicari berdasarkan nama (substring, tidak peka huruf besar/kecil).
  4. Aksi per baris: View, Edit, Copy, dan Delete; Delete hanya aktif untuk workflow inactive.
  5. ⚠ Tombol "Create workflow" membuka editor dengan status awal default (To Do, In Progress, Done) atau, bila dipilih, dari salah satu template workflow bawaan.
  6. Workflow sistem read-only ("jira" dan "Builds Workflow") diberi penanda "System" dan tidak memiliki aksi Edit maupun Delete.
  7. Workflow yang memiliki draft belum dipublish (data legacy hasil impor) diberi penanda "Draft".
  8. Tabel mendukung sorting per kolom dan pagination (default 50 per halaman) untuk site dengan banyak workflow.
  9. Workflow team-managed tidak tampil di halaman ini dan hanya bisa dilihat dari project masing-masing.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Bergantung pada WF-021, WF-023, WF-024, WF-033. Aksi Copy (WF-023) dan Delete beserta tab Inactive yang lengkap (WF-024) dijadwalkan Fase 3; pada Fase 2 aksi tersebut disembunyikan, dan admin membuat workflow baru lewat Create workflow.

### WF-020 System default workflow ('jira' workflow) & Default Workflow Scheme

- **Deskripsi:** Workflow bawaan read-only bernama "jira" dan Default Workflow Scheme yang dipakai project company-managed tanpa scheme khusus.
- **Prioritas / Fase / Tier:** P1 / Fase 1 / free-standard
- **Aktor & izin:** Sistem (seed); Admin Situs (melihat dan menyalin).
- **Aturan perilaku (acceptance criteria):**
  1. Workflow "jira" berisi status Open, In Progress, Resolved, Reopened, dan Closed, dengan transition Start Progress, Stop Progress, Resolve Issue, Close Issue, dan Reopen Issue.
  2. Transition Resolve Issue dan Close Issue memakai "Resolve Issue Screen" untuk mengisi Resolution dan Fix Version.
  3. Transition Reopen Issue memakai post function yang mengosongkan resolution.
  4. Workflow ini tidak bisa diedit atau dihapus (UI dan API menolak dengan 403/409), tetapi bisa di-copy menjadi workflow baru yang bisa diedit (WF-023).
  5. Default Workflow Scheme memetakan semua work type ke workflow "jira".
  6. ⚠ Project baru yang dibuat dengan shared configuration mewarisi workflow scheme project sumber; project baru tanpa template workflow khusus memakai Default Workflow Scheme.
  7. Template project company-managed (Scrum, Kanban, business) membuat workflow dan scheme sendiri yang diberi prefix key project, mis. "ABC: Software Simplified Workflow for Project ABC".
  8. Workflow sistem "Builds Workflow" (read-only) di-seed untuk dipakai fitur DevOps/Builds.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai workflow default per work type (WF-017 butir 3).
- **Catatan & ketergantungan:** Ditempatkan di Fase 1 (pengecualian dari P1) karena [04 Proyek](04-proyek.md) menetapkan project company-managed Fase 1 memakai scheme default sistem tanpa editor. Kemampuan copy pada butir 4 tersedia bersama WF-023 di Fase 3; sebelum itu workflow "jira" hanya bisa dilihat. Shared configuration di PRJ-014; Builds di [24 Integrasi DevOps](24-integrasi-devops.md). Resolve Issue Screen di-seed oleh modul [05](05-konfigurasi-work-item-skema.md).

### WF-021 Workflow Editor (company-managed, new editor)

- **Deskripsi:** Editor visual berbasis kanvas untuk membuat dan mengubah workflow company-managed: menambah status, transition, rules, approvals, dan properties. Produk ini hanya membangun satu editor diagram modern ini; editor legacy tidak direplikasi (WF-044).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek dengan extended project administration dalam batasan WF-038; pengguna lain read-only.
- **Aturan perilaku (acceptance criteria):**
  1. Kanvas menampilkan node status dan panah transition, dengan pan, zoom in/out, dan fit to screen.
  2. Status bisa di-drag untuk memindahkan posisinya; tata letak disimpan sebagai layout (koordinat x,y per status) bersama versi workflow.
  3. "Add status" memungkinkan memilih status existing atau membuat status baru (nama + kategori); status baru bersifat global.
  4. "Add transition" meminta status asal (satu, beberapa, atau Any status), status tujuan, nama, deskripsi, dan screen.
  5. Memilih status atau transition membuka panel samping untuk mengedit nama, rules, screen, dan properties.
  6. Rule picker "Add rule" pada transition dikelompokkan menjadi Restrict transition (conditions), Validate details (validators), dan Perform actions (post functions); rule dari app ekstensi muncul di kelompok yang sesuai.
  7. Rule dikonfigurasi lewat form dengan picker user/group/role/field/status; ringkasan setiap rule tampil di panel transition dengan tombol edit dan delete.
  8. Toggle "Show transition labels" menampilkan atau menyembunyikan nama transition di kanvas.
  9. Setiap status memiliki opsi "Allow all statuses to transition to this one" untuk membuat global transition (WF-006).
  10. Initial status (status tempat work item dibuat) bisa diubah (WF-005).
  11. Header menampilkan scope pemakaian ("Used by N projects / M work types") dan peringatan bila workflow dipakai bersama (shared).
  12. Validasi sebelum menyimpan: initial transition wajib ada, status yang tidak terjangkau diberi warning, dan rule yang konfigurasinya tidak lengkap ditolak dengan sorotan pada rule terkait.
  13. Tombol "Update workflow" mempublish perubahan, dan "Discard changes" membatalkan seluruh perubahan sesi edit.
  14. Menyimpan workflow aktif yang statusnya dihapus memicu wizard status migration: work item di status yang dihapus dipetakan ke status pengganti per work type (WF-036).
  15. Pengguna tanpa hak edit membuka editor dalam mode read-only (View workflow).
  16. Nama dan deskripsi workflow bisa diedit; nama mengikuti aturan unik WF-025.
  17. Status bisa diberi approval (lihat BIZ-042 di [26 Kerja Bisnis](26-kerja-bisnis-formulir-persetujuan.md), tersedia mulai Fase 3).
  18. ⚠ Transition bisa diberi agent AI yang berjalan saat work item berpindah ke status tujuan (WF-043).
  19. ⚠ Undo/redo dalam satu sesi edit bersifat opsional (nice-to-have) dan tidak menjadi syarat rilis editor.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai WF-017.
- **Catatan & ketergantungan:** Fitur ini mencakup juga rule picker editor baru. Komponen kanvas dipakai ulang oleh viewer WF-015. Konkurensi dan versi di WF-022 dan WF-025. Frontend: komponen kanvas SvelteKit (pustaka diagram dipilih di [27 NFR](27-nfr-arsitektur-platform.md)).

### WF-022 Workflow drafts (draft / publish / discard)

- **Deskripsi:** Mekanisme untuk mengedit workflow aktif tanpa langsung memengaruhi work item, sampai perubahan dipublish atau dibuang.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Model legacy (untuk kompatibilitas impor dan API): mengedit workflow aktif membuat draft, sedangkan workflow inactive diedit langsung.
  2. Model legacy: Publish Draft menyediakan opsi "Save a backup copy of the active workflow" dengan nama backup yang bisa diatur; backup tersimpan sebagai workflow inactive.
  3. Model legacy: Discard Draft membuang semua perubahan draft.
  4. Model legacy: di draft workflow aktif, status tidak bisa dihapus dan linked status tidak bisa diubah; UI mengarahkan admin untuk meng-copy workflow lalu migrasi lewat scheme.
  5. Editor baru (default produk ini): perubahan disimpan sebagai sesi edit lokal lalu dipublish dengan "Update workflow"; penghapusan status yang masih punya work item memicu migrasi status.
  6. Hanya ada satu draft per workflow pada model legacy.
  7. Konkurensi: penyimpanan dengan versi workflow yang sudah kedaluwarsa ditolak dengan 409 dan pesan "Workflow telah diubah pengguna lain"; REST memakai `version.id`/`versionNumber`.
  8. Publish berlaku langsung ke semua project dan work type yang memakai workflow tersebut.
  9. Setiap publish dicatat di audit log (WF-039).
- **Team-managed vs company-managed:** Team-managed tidak memakai draft; perubahan berlaku langsung (WF-017 butir 13).
- **Catatan & ketergantungan:** Keputusan apakah model draft legacy dibangun penuh atau hanya editor sesi tercatat di Pertanyaan terbuka. Workflow disimpan berversi (modul 28).

### WF-023 Copy workflow

- **Deskripsi:** Menduplikasi workflow yang ada, termasuk workflow sistem, beserta status, transition, rules, screen, dan properties.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Copy tersedia dari daftar workflow (WF-019) dan dari editor.
  2. Copy meminta nama baru (wajib, unik) dan deskripsi; nama default adalah "Copy of <nama asal>".
  3. Salinan menyertakan layout diagram, rules, screen mapping, dan properties.
  4. Status global dipakai bersama oleh salinan dan tidak diduplikasi.
  5. Salinan selalu berstatus inactive sampai ditambahkan ke scheme yang aktif.
  6. Workflow sistem "jira" bisa di-copy sebagai titik awal kustomisasi.
  7. ⚠ Team-managed tidak menyediakan copy workflow antar project.
- **Team-managed vs company-managed:** Hanya company-managed (butir 7).
- **Catatan & ketergantungan:** Pola "memisahkan workflow satu project" memakai fitur ini bersama WF-033.

### WF-024 Inactive workflows & deletion

- **Deskripsi:** Pengelolaan workflow yang tidak dipakai scheme aktif mana pun, termasuk penghapusan dan pembersihan.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Workflow disebut inactive bila tidak ada di scheme mana pun, atau hanya ada di scheme yang tidak diasosiasikan ke project.
  2. Workflow inactive diedit langsung tanpa draft dan tanpa migrasi status.
  3. Delete hanya diizinkan untuk workflow inactive; workflow sistem tidak pernah bisa dihapus.
  4. ⚠ Workflow yang masih tercantum di scheme inactive harus dilepas dari scheme tersebut sebelum bisa dihapus; UI menampilkan daftar scheme yang menahannya.
  5. Status hanya bisa dihapus bila tidak ada workflow, termasuk workflow inactive, yang memakainya.
  6. Workflow team-managed terhapus otomatis bersama project atau work type pemiliknya.
  7. ⚠ Cleanup massal didukung lewat REST bulk delete; UI menghapus satu per satu dengan konfirmasi.
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** Bergantung pada WF-033 untuk status keterkaitan scheme.

### WF-025 Workflow limits & validation constraints

- **Deskripsi:** Batasan sistem dan aturan validasi yang menjaga konsistensi workflow.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Sistem; Admin Situs menerima pesan validasi.
- **Aturan perilaku (acceptance criteria):**
  1. Workflow wajib memiliki initial transition dan minimal satu status; pelanggaran ditolak saat disimpan.
  2. Nama workflow company-managed unik per site (tidak peka huruf besar/kecil) dan dibatasi panjangnya (default 255 karakter).
  3. Nama status dan transition wajib diisi dan dibatasi panjangnya (default 60 karakter untuk status, 255 untuk transition).
  4. Satu status tidak boleh muncul dua kali dalam satu workflow.
  5. ⚠ Batas jumlah status/transition per workflow dan jumlah workflow per site dapat dikonfigurasi (default 100 status, 500 transition per workflow, dan 2000 workflow per site); pelanggaran ditolak dengan pesan yang menyebut batasnya.
  6. Workflow sistem tidak bisa diedit atau dihapus.
  7. Transition yang menuju status yang tidak ada di workflow ditolak.
  8. Status tanpa transition masuk (unreachable) atau tanpa transition keluar (dead-end, selain status berkategori Done) diberi peringatan, bukan galat.
  9. Optimistic locking: setiap update workflow lewat REST atau editor wajib menyertakan versi; versi kedaluwarsa ditolak dengan 409.
- **Team-managed vs company-managed:** Batas jumlah status team-managed lebih ketat (WF-017 butir 11); aturan validasi lainnya sama.
- **Catatan & ketergantungan:** Implikasi model data (modul 28): `workflow` (berversi), `workflow_status`, `workflow_transition` (from-set, to, type initial/global/directed/looped), `transition_rule` (kind, config JSON, order), `workflow_property` (owner_type, key, value), `workflow_scheme`, `workflow_scheme_mapping` (`issue_type_id` nullable = default), dan `workflow_layout`. Angka batas default adalah usulan dan perlu keputusan pemilik produk.

### WF-026 Simplified Workflow (board-managed)

- **Deskripsi:** Mode workflow di project company-managed software yang dikelola dari board: semua status terhubung lewat global transition, tanpa screen dan rules.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Board administrator, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Template company-managed Scrum/Kanban membuat workflow "Software Simplified Workflow for Project XXX" secara otomatis dengan flag `simplified`.
  2. Opsi legacy "Use Jira Workflow Simplification" saat membuat board dari project yang sudah ada mengonversi workflow project menjadi simplified.
  3. Semua status saling terhubung lewat global transitions, tanpa conditions, validators, maupun post functions.
  4. Board administrator bisa menambah dan menghapus status dari Board settings > Columns tanpa menjadi Admin Situs.
  5. Resolution otomatis diisi saat work item masuk status yang diberi checkbox "Set resolution" (opsi per status di Board settings > Columns, bukan per kolom), dan dikosongkan saat keluar dari status tersebut (BRD-027 butir 5).
  6. Bila admin mengedit workflow secara manual di editor (WF-021), flag simplified dihapus setelah konfirmasi peringatan, dan board administrator kehilangan kemampuan mengelola status.
  7. Mode simplified hanya berlaku bila project memakai satu workflow untuk semua work type; opsi tidak ditawarkan bila tidak terpenuhi.
- **Team-managed vs company-managed:** Hanya company-managed software; team-managed sudah memakai pola serupa secara bawaan.
- **Catatan & ketergantungan:** Tidak wajib untuk MVP. Board di [10 Board](10-board-kolom-kartu.md); resolution di WF-009. Sisi board (tombol Simplify workflow, Add status, Set resolution per status, wizard migrasi) dirinci di BRD-027 modul 10; WF-026 memegang engine workflow-nya dan keduanya harus memakai aturan yang sama. Mapping kolom di WF-012 dan BRD-025.

## Rules, screen & properties company-managed

### WF-027 Transition Conditions (company-managed)

- **Deskripsi:** Rules yang menentukan apakah sebuah transition terlihat dan bisa dijalankan oleh pengguna tertentu untuk work item tertentu; transition yang gagal condition disembunyikan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; condition dievaluasi untuk setiap pengguna yang melihat atau menjalankan transition.
- **Aturan perilaku (acceptance criteria):**
  1. Di rule picker, conditions dikelompokkan sebagai "Restrict transition".
  2. Only Assignee Condition: transition hanya tersedia bagi assignee work item.
  3. Only Reporter Condition: transition hanya tersedia bagi reporter work item.
  4. Permission Condition: transition hanya tersedia bagi pengguna yang memiliki project permission tertentu (mis. Resolve issues, Close issues).
  5. User Is In Group: transition hanya tersedia bagi anggota salah satu group yang dipilih.
  6. User Is In Project Role: transition hanya tersedia bagi pengguna dengan project role tertentu di project work item.
  7. User Is In Group Custom Field / User Is In Custom Field: transition hanya tersedia bagi pengguna yang tercantum di field user picker atau anggota group di field group picker yang dipilih.
  8. Hide From User Condition: transition tidak pernah tampil di UI dan hanya bisa dipicu otomasi, workflow triggers, REST, atau post function.
  9. Sub-Task Blocking Condition ("Restrict to when all subtasks have a specific status"): transition hanya tersedia bila semua subtask berada di status yang dipilih.
  10. Separation of Duties Condition: pengguna yang pernah menjalankan transition tertentu pada work item ini tidak boleh menjalankan transition ini.
  11. Previous Status Condition: work item pernah berada di status tertentu, atau (opsi most recent) status sebelumnya adalah status tersebut.
  12. Value Field Condition: nilai field dibandingkan dengan nilai tertentu memakai operator =, !=, >, <, >=, <=.
  13. ⚠ Compare Number Custom Field Condition: nilai custom field angka dibandingkan dengan konstanta memakai operator perbandingan.
  14. ⚠ Condition berbasis ekspresi (build-your-own) bawaan: admin menulis ekspresi boolean atas data work item dan pengguna; ekspresi yang tidak valid ditolak saat disimpan.
  15. Beberapa condition dikombinasikan dengan logika ALL (AND) atau ANY (OR) di tingkat transition; grouping bersarang tidak didukung di editor baru.
  16. App ekstensi dapat menyediakan condition berbasis ekspresi sebagai titik ekstensi (WF-042).
  17. Condition dievaluasi saat dropdown dirender dan dicek ulang di server saat eksekusi; transition yang lolos di UI tetapi gagal di server ditolak dengan galat.
  18. ⚠ Workflow triggers DevOps mengabaikan conditions, sedangkan action otomasi "Transition work item" tetap menghormati conditions untuk aktor rule.
- **Team-managed vs company-managed:** Hanya company-managed; padanan team-managed adalah rule restrict di WF-018.
- **Catatan & ketergantungan:** Rule dari app Marketplace pihak ketiga (JMWE, JSU, ScriptRunner) tidak dibangun native. Condition khusus JSM (customer/agent) ada di modul 29. Bahasa ekspresi ditetapkan bersama modul [23 API](23-api-publik-webhook.md).

### WF-028 Transition Validators (company-managed)

- **Deskripsi:** Rules yang memeriksa input atau keadaan work item saat transition dijalankan; jika gagal, transition dibatalkan dan pesan galat ditampilkan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; validator dijalankan untuk semua aktor yang menjalankan transition.
- **Aturan perilaku (acceptance criteria):**
  1. Di rule picker, validators dikelompokkan sebagai "Validate details".
  2. Field Required Validator: satu atau lebih field wajib terisi, dengan opsi mengabaikan context field dan pesan galat custom.
  3. Field Changed Validator: field yang dipilih harus diubah selama transition (nilai baru berbeda dari nilai lama).
  4. Field Has Single Value Validator: field multi-value hanya boleh berisi satu nilai, dengan opsi mengabaikan nilai kosong.
  5. Regular Expression Check: nilai field harus cocok dengan regex yang dikonfigurasi; regex tidak valid ditolak saat disimpan.
  6. Date Compare Validator: membandingkan dua field tanggal dengan operator >, >=, =, <=, < dan opsi menyertakan waktu.
  7. ⚠ Date Range Validator: selisih dua field tanggal harus berada dalam jumlah hari tertentu.
  8. Permission Validator: pengguna yang menjalankan transition harus memiliki permission tertentu.
  9. User Permission Validator: pengguna di field user tertentu harus memiliki permission tertentu.
  10. Previous Status Validator: work item pernah berada di status tertentu, dengan opsi most recent only.
  11. Parent Status Validator: parent work item harus berada di status tertentu (untuk subtask).
  12. ⚠ Validator berbasis ekspresi (build-your-own) tersedia bawaan, dan app ekstensi dapat menyediakan validator sendiri (WF-042).
  13. Pesan galat custom ditampilkan di dialog transition, inline di issue view, atau toast di board sesuai jalur eksekusi.
  14. Validator juga berlaku pada initial transition (create).
  15. Validator berjalan setelah transition screen disubmit dan sebelum post functions.
  16. Validator dijalankan juga pada eksekusi lewat REST API, otomasi, dan bulk transition; work item yang gagal dilaporkan per work item.
- **Team-managed vs company-managed:** Hanya company-managed; padanan team-managed adalah rule validate di WF-018.
- **Catatan & ketergantungan:** Validator JSU/JMWE adalah app Marketplace dan tidak dibangun native. Field configuration (field wajib) di [05 Konfigurasi](05-konfigurasi-work-item-skema.md) tetap berlaku terpisah dari validator.

### WF-029 Transition Post Functions (company-managed)

- **Deskripsi:** Aksi yang dijalankan otomatis setelah transition berhasil, seperti mengubah field, assign, fire event, atau memanggil webhook.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; Sistem menjalankan.
- **Aturan perilaku (acceptance criteria):**
  1. Di rule picker, post functions dikelompokkan sebagai "Perform actions".
  2. Post functions wajib bawaan tidak bisa dihapus dan urutannya tetap: Set issue status to the linked status, Add a comment if entered, Update change history and store issue, Re-index issue, dan Fire event.
  3. Assign to current user: hanya berhasil bila pengguna memiliki permission Assignable user; jika tidak, assignee tidak diubah dan transition tetap berhasil.
  4. Assign to lead developer: assign ke component lead, atau ke project lead bila work item tidak memiliki component berlead.
  5. Assign to reporter: assign ke reporter work item.
  6. Update Issue Field: menetapkan nilai field sistem (Assignee, Resolution, Priority, Summary, Description, Environment, Original/Remaining Estimate, dll.).
  7. ⚠ Update Issue Custom Field: menetapkan nilai custom field bertipe teks, angka, tanggal, select, dan user picker.
  8. Clear Field Value: mengosongkan field yang dipilih.
  9. ⚠ Copy Value From Other Field: menyalin nilai field ke field lain di work item yang sama, dengan opsi menyalin dari parent.
  10. Trigger a Webhook: memanggil webhook yang terdaftar di Admin > System > WebHooks.
  11. Fire event: memilih event (Generic Event, Issue Resolved, Issue Closed, custom event) yang dipakai notification scheme (WF-016).
  12. Urutan post functions non-wajib bisa diubah (move up/down); editor menolak penempatan post function pengubah field setelah "store issue" dengan pesan bahwa perubahan tidak akan tersimpan.
  13. App ekstensi dapat menyediakan post function sendiri (WF-042).
  14. Pola reopen: Update Issue Field dengan Resolution = None mengosongkan resolution.
  15. Post functions juga dijalankan untuk transition yang dipicu otomasi, REST, workflow triggers, dan bulk change.
  16. Pemanggilan webhook berjalan asinkron; kegagalannya dicatat di log pengiriman webhook dan tidak membatalkan transition.
- **Team-managed vs company-managed:** Hanya company-managed; padanan team-managed adalah rule perform actions di WF-018.
- **Catatan & ketergantungan:** Post function "Set issue security level based on user's project role" dan "Create Perforce Job" hanya ada di Server/DC dan tidak dibangun. WebHooks di [23 API](23-api-publik-webhook.md).

### WF-030 Transition screens

- **Deskripsi:** Dialog screen yang tampil saat pengguna menjalankan transition untuk mengumpulkan input, misalnya Resolution, Fix Version, atau komentar.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs memilih screen; pengguna yang menjalankan transition mengisi dialog.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap transition memilih satu screen dari dropdown daftar screens, atau "No screen" sehingga transition langsung dijalankan tanpa dialog.
  2. Seed menyediakan screen bawaan "Resolve Issue Screen" dan "Workflow Screen".
  3. Dialog transition menampilkan field sesuai screen, ditambah field Comment yang selalu tersedia.
  4. Field Resolution hanya muncul di dialog bila ditambahkan ke transition screen.
  5. Field wajib dari field configuration dan validators diterapkan saat dialog disubmit; galat ditampilkan per field.
  6. Dialog juga muncul saat drag-drop di board dan saat bulk transition, dengan satu input yang diterapkan ke semua work item.
  7. REST: field yang dikirim bersama transition harus ada di transition screen; jika tidak, request ditolak dengan pesan "Field cannot be set. It is not on the appropriate screen".
  8. Tab screen ditampilkan sebagai tab di dialog transition.
  9. Satu screen per transition berlaku untuk semua work type yang memakai workflow tersebut.
  10. ⚠ Team-managed tidak memiliki transition screen custom; resolution diatur otomatis, dan rule "show a screen" tidak disediakan.
- **Team-managed vs company-managed:** Lihat butir 10.
- **Catatan & ketergantungan:** Definisi screen dan screen scheme di [05 Konfigurasi](05-konfigurasi-work-item-skema.md) (CFG-046, CFG-047). Bulk transition di WF-013.

### WF-031 Workflow properties (status & transition properties)

- **Deskripsi:** Pasangan key-value pada status atau transition yang mengubah perilaku produk, misalnya membuat work item read-only atau membatasi permission per status.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Editor menyediakan panel Properties untuk status dan transition dengan aksi add, edit, dan delete key/value.
  2. `jira.issue.editable = false` pada status membuat work item tidak bisa diedit di status itu, kecuali lewat transition screen.
  3. `jira.permission.[permission].[type]` (mis. `jira.permission.comment.user`, `jira.permission.edit.group`, `jira.permission.assign.denied`) membatasi atau memberi permission per status.
  4. `jira.permission.subtasks.[permission].[type]` menerapkan pembatasan yang sama pada subtask dari work item di status tersebut.
  5. `opsbar-sequence` pada transition mengatur urutan tampil (WF-008).
  6. `jira.field.resolution.include` / `jira.field.resolution.exclude` (daftar ID resolution dipisah koma) membatasi pilihan resolution pada transition.
  7. `jira.i18n.title` / `jira.i18n.description` menyediakan label transition yang diterjemahkan (WF-040).
  8. ⚠ `jira.i18n.submit` menyediakan label tombol submit dialog transition yang diterjemahkan.
  9. Key berawalan `jira.` dicadangkan untuk sistem; property custom dengan key lain disimpan apa adanya dan bisa dibaca app atau REST.
  10. Properties tidak tersedia di team-managed.
  11. Panel Properties menampilkan peringatan bahwa property yang salah konfigurasi bisa mengunci work item (mis. menolak semua edit), dan key `jira.*` yang tidak dikenal ditandai.
- **Team-managed vs company-managed:** Hanya company-managed (butir 10).
- **Catatan & ketergantungan:** Evaluasi `jira.permission.*` dilakukan oleh mesin izin di [03 Izin](03-izin-peran-keamanan.md) (PERM-037). Daftar lengkap key yang didukung ditetapkan saat desain.

### WF-032 Subtask workflow

- **Deskripsi:** Workflow untuk work type subtask serta interaksi status antara parent dan subtask.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; pengguna menjalankan transition parent dan subtask.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: work type subtask dipetakan ke workflow lewat workflow scheme seperti work type lain.
  2. Team-managed: subtask memiliki workflow sendiri yang diedit seperti work type lain (WF-017).
  3. Sub-Task Blocking Condition pada parent mencegah parent ditutup sebelum semua subtask berada di status tertentu (WF-027).
  4. Parent Status Validator pada subtask mewajibkan parent berada di status tertentu (WF-028).
  5. Permission per status untuk subtask diatur lewat property `jira.permission.subtasks.*` (WF-031).
  6. Board company-managed: ketika subtask terakhir dipindahkan ke kolom Done, board menawarkan dialog untuk memindahkan parent ke Done juga.
  7. Konversi work item ke subtask atau sebaliknya memicu pemetaan status bila workflow berbeda (WF-036).
  8. Auto-transition parent ketika semua subtask selesai dilakukan lewat otomasi (template rule), bukan perilaku bawaan workflow.
- **Team-managed vs company-managed:** Lihat butir 1 dan 2.
- **Catatan & ketergantungan:** Template rule otomasi di [20 Otomasi](20-otomasi.md); konversi subtask di [06 Work Item Inti](06-work-item-inti.md).

## Workflow scheme & migrasi status

### WF-033 Workflow schemes

- **Deskripsi:** Scheme yang memetakan work type ke workflow dan diasosiasikan ke satu atau lebih project company-managed; fitur ini juga mencakup analisis dampak workflow yang dibagi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengelola; Admin Proyek melihat scheme project-nya.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman admin Workflow schemes menampilkan daftar scheme active/inactive, project yang memakai, serta aksi Edit, Copy, dan Delete.
  2. "Create workflow scheme" meminta nama (wajib, unik) dan deskripsi.
  3. Admin dapat menambahkan workflow ke scheme lalu meng-assign satu atau lebih work type ke workflow tersebut.
  4. Setiap scheme memiliki default workflow untuk work type yang tidak dipetakan eksplisit ("All unassigned issue types").
  5. Satu work type hanya dipetakan ke satu workflow dalam satu scheme; meng-assign work type yang sudah dipetakan memindahkan pemetaannya.
  6. Remove workflow dari scheme mengembalikan work type terkait ke default workflow scheme.
  7. Scheme bisa di-copy beserta seluruh pemetaannya.
  8. Delete hanya diizinkan untuk scheme yang tidak diasosiasikan ke project mana pun.
  9. Satu scheme bisa dipakai banyak project, dan perubahan scheme berdampak ke semua project tersebut.
  10. Mengedit scheme yang aktif dipakai project membuat draft scheme (WF-034).
  11. Default Workflow Scheme sistem bersifat read-only dan dipakai project tanpa scheme khusus (WF-020).
  12. Analisis dampak: workflow/scheme yang dipakai lebih dari satu project diberi badge "Shared", disertai daftar project dan work type pemakai, serta peringatan sebelum publish.
  13. Alur memisahkan workflow untuk satu project didukung: copy workflow dan scheme, lalu switch scheme project tersebut (WF-023, WF-036).
- **Team-managed vs company-managed:** Hanya company-managed; team-managed tidak memakai scheme.
- **Catatan & ketergantungan:** Fitur ini mencakup juga analisis dampak berbagi workflow. Butir 13 memakai Copy workflow (WF-023, Fase 3); pada Fase 2 admin membuat workflow baru lalu switch scheme. Pola scheme selaras dengan scheme lain di [05 Konfigurasi](05-konfigurasi-work-item-skema.md) (CFG-059).

### WF-034 Workflow scheme drafts

- **Deskripsi:** Saat scheme yang aktif dipakai project diubah, sistem membuat draft yang perlu dipublish; publish bisa memicu migrasi status work item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Draft otomatis dibuat ketika admin mengedit scheme aktif (assign work type atau mengganti workflow).
  2. Scheme yang memiliki draft menampilkan banner "Draft" dengan aksi Publish dan Discard.
  3. Saat publish, wizard memeriksa work item yang statusnya tidak ada di workflow baru dan meminta pemetaan status (WF-036).
  4. Publish berjalan sebagai background task dengan progress bar.
  5. Hanya ada satu draft per scheme; edit berikutnya memperbarui draft yang sama.
  6. Discard membuang draft dan mengembalikan scheme ke versi aktif.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pada Fase 2, sebelum fitur ini tersedia di Fase 3, edit scheme aktif langsung menjalankan wizard migrasi WF-036.

### WF-035 Project-level workflow settings (company-managed)

- **Deskripsi:** Halaman Project settings > Workflows yang menampilkan workflow scheme project, workflow per work type, dan akses untuk mengganti scheme atau mengedit workflow.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (melihat; mengedit dengan extended project administration), Admin Situs (penuh).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan nama scheme aktif dan daftar workflow per work type.
  2. Setiap workflow memiliki tautan View dan Edit; Edit hanya tampil untuk Admin Situs, atau Admin Proyek dengan extended project administration (WF-038).
  3. Tombol "Switch scheme" hanya tersedia untuk Admin Situs.
  4. Workflow shared diberi penanda beserta jumlah project yang memakainya.
  5. Admin Proyek tanpa hak edit melihat halaman secara read-only.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai halaman Work types (WF-017).
- **Catatan & ketergantungan:** Halaman Project settings di [04 Proyek](04-proyek.md) (PRJ-044).

### WF-036 Associate workflow scheme & status migration

- **Deskripsi:** Proses mengganti workflow scheme project atau workflow sebuah work type, termasuk memetakan status lama ke status baru untuk work item yang sudah ada.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (switch scheme, publish); pengguna dengan Move issues (move/convert work item); Sistem menjalankan migrasi.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Workflows > Switch scheme meminta pemilihan scheme target.
  2. Wizard migrasi berjalan per work type: setiap status lama yang tidak ada di workflow baru wajib dipetakan ke status tujuan sebelum tombol lanjut aktif.
  3. Status yang ada di kedua workflow (ID sama) dipertahankan otomatis tanpa pemetaan.
  4. Sebelum konfirmasi, wizard menampilkan ringkasan jumlah work item terdampak per work type dan status.
  5. Migrasi berjalan sebagai background task dengan progress bar dan tidak bisa dibatalkan setelah dimulai.
  6. ⚠ Selama migrasi, konfigurasi workflow project dikunci; upaya transition work item terdampak ditolak sementara dengan pesan "Migrasi workflow sedang berjalan".
  7. ⚠ Perubahan status hasil migrasi dicatat di history work item sebagai entri perubahan status dengan pelaku "System (migrasi workflow)", bukan sebagai transition biasa.
  8. Migrasi juga terjadi saat publish workflow aktif yang statusnya dihapus (WF-021) dan saat publish scheme draft (WF-034).
  9. Move work item ke project atau work type lain dengan workflow berbeda meminta pemilihan status target.
  10. Change work type di tempat dan konversi work item ke subtask atau sebaliknya meminta pemetaan status bila workflow berbeda.
  11. Subtask ikut dipetakan sesuai workflow work type subtask di scheme target.
  12. Work item yang gagal dimigrasi dilaporkan di ringkasan akhir beserta alasannya.
  13. REST: switch scheme dengan status mapping berjalan sebagai async task yang bisa di-poll.
- **Team-managed vs company-managed:** Team-managed hanya memakai pemetaan saat menghapus status (WF-017 butir 7) dan saat move/convert; switch scheme hanya company-managed.
- **Catatan & ketergantungan:** Move dan change work type di [06 Work Item Inti](06-work-item-inti.md); background task di [27 NFR](27-nfr-arsitektur-platform.md); async task di WF-041.

## Izin, audit & terjemahan

### WF-037 Workflow-related permissions

- **Deskripsi:** Project permission dan global permission yang mengatur siapa yang bisa menjalankan, meresolve, dan mengonfigurasi workflow.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek, semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Transition issues adalah syarat dasar untuk menjalankan transition apa pun; tanpa permission ini semua jalur eksekusi menolak dengan 403.
  2. Resolve issues mengatur siapa yang bisa mengisi resolution dan Fix Version, dan biasa dipakai di Permission Condition.
  3. Close issues dipakai Permission Condition pada transition Close.
  4. Assignable user menjadi syarat agar post function assign berhasil untuk pengguna tujuan.
  5. Administer projects memberi hak mengedit workflow team-managed dan, bila diaktifkan, extended project administration.
  6. Administer Jira (global) memberi hak mengelola status, workflow, scheme, dan resolution.
  7. Make bulk changes (global) menjadi syarat bulk transition.
  8. Pada company-managed, permission bisa di-override per status lewat property `jira.permission.*` (WF-031).
- **Team-managed vs company-managed:** Butir 8 hanya company-managed; team-managed memakai rule restrict (WF-018).
- **Catatan & ketergantungan:** Definisi permission scheme di [03 Izin](03-izin-peran-keamanan.md) (PERM-023, PERM-029, PERM-011). Fase 1 cukup memeriksa Transition issues dan Administer projects dengan scheme default; butir lain aktif bersama fiturnya.

### WF-038 Extended project administration (workflow editing by project admins)

- **Deskripsi:** Admin Proyek company-managed bisa mengedit workflow project mereka dengan batasan, tanpa menjadi Admin Situs.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (company-managed) dengan Administer projects.
- **Aturan perilaku (acceptance criteria):**
  1. Hak edit hanya berlaku bila workflow tidak dipakai bersama project lain (tidak shared).
  2. Admin Proyek bisa menambah status existing, menambah/menghapus transition, dan mengubah tata letak diagram.
  3. ⚠ Admin Proyek tidak bisa membuat status global baru, me-rename status global, atau mengedit conditions, validators, post functions, screens, maupun properties; kontrol tersebut disembunyikan di editor.
  4. ⚠ Admin Proyek tidak bisa menghapus status yang masih berisi work item; aksi hapus dinonaktifkan dengan tooltip alasan.
  5. Fitur ini mewajibkan permission Administer projects pada project terkait.
  6. Workflow shared ditampilkan read-only dengan pesan untuk menghubungi Admin Situs.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Aturan umum extended project administration di [04 Proyek](04-proyek.md) (PRJ-036).

### WF-039 Workflow audit logging

- **Deskripsi:** Pencatatan perubahan konfigurasi workflow, status, scheme, dan migrasi di audit log.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs dan Admin Organisasi melihat; Sistem mencatat.
- **Aturan perilaku (acceptance criteria):**
  1. Audit log mencatat create/update/delete workflow, status, resolution, dan workflow scheme.
  2. Audit log mencatat publish draft, migrasi scheme, dan perubahan rules dengan ringkasan perubahan (sebelum/sesudah).
  3. Setiap entri memuat pelaku, waktu, objek, dan project terkait.
  4. Entri diakses lewat Admin > System > Audit log dan bisa difilter berdasarkan kategori "Workflows".
- **Team-managed vs company-managed:** Perubahan workflow team-managed juga dicatat dengan project sebagai konteks.
- **Catatan & ketergantungan:** Infrastruktur audit log di [22 Administrasi Sistem](22-administrasi-sistem.md) (ADM-014, Fase 3), sehingga butir audit di WF-003, WF-010, dan WF-022 (Fase 2) baru tertulis ke audit log mulai Fase 3; sebelumnya perubahan konfigurasi dicatat di log aplikasi terstruktur. Lihat juga [02 Pengguna](02-pengguna-grup-organisasi.md) (USR-046).

### WF-040 Translations for statuses, resolutions & transitions

- **Deskripsi:** Terjemahan nama dan deskripsi status, resolution, serta label transition untuk pengguna dengan bahasa berbeda.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Translate di halaman Statuses dan Resolutions tersedia untuk setiap bahasa yang didukung produk (minimal EN dan ID).
  2. Property transition `jira.i18n.title` / `jira.i18n.description` menyediakan label transition terjemahan.
  3. Bila terjemahan untuk bahasa pengguna tidak ada, UI menampilkan nama asli.
  4. ⚠ JQL hanya menerima nama asli atau ID status/resolution; nama terjemahan tidak dikenali sebagai nilai JQL.
- **Team-managed vs company-managed:** Hanya company-managed (team-managed tidak punya halaman Statuses maupun properties).
- **Catatan & ketergantungan:** Bahasa pengguna ditentukan seperti `LanguageMiddleware` di repo saat ini (`Accept-Language`, EN/ID). Selaras dengan keputusan terjemahan data konfigurasi di modul 05 (CFG-061).

## API, ekstensi & AI

### WF-041 Workflow REST API surface (workflow domain)

- **Deskripsi:** Endpoint REST untuk membaca dan mengelola workflow, status, transition, resolution, dan scheme, yang dibutuhkan untuk integrasi dan otomasi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, otomasi, Sistem; setiap endpoint menghormati permission yang sama dengan UI.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/{key}/transitions` (dengan `expand=transitions.fields`) mengembalikan transition yang tersedia bagi pemanggil, dan `POST /issue/{key}/transitions` menjalankan transition beserta `fields`, `update`, dan `comment`.
  2. Statuses API menyediakan search, bulk create/update/delete, usages, serta `GET /status` dan `GET /statuscategory`.
  3. Workflows API menyediakan search, bulk create/update (termasuk layout, rules, dan statuses), validate create/update tanpa menyimpan, dan capabilities (daftar rule yang tersedia).
  4. Workflow schemes API menyediakan CRUD, draft, pemetaan work type, dan switch scheme dengan status mapping sebagai async task.
  5. Workflow transition properties API dan workflow transition rules API tersedia untuk app ekstensi.
  6. Resolutions API menyediakan CRUD, search, dan set default.
  7. Async task polling (`GET /task/{id}`) mengembalikan status, progres, dan hasil migrasi.
- **Team-managed vs company-managed:** Endpoint scheme dan properties hanya berlaku untuk company-managed; endpoint transition dan status berlaku untuk keduanya.
- **Catatan & ketergantungan:** Nama endpoint mengacu ke pola REST v3 untuk kompatibilitas integrasi; autentikasi, rate limit, dan versi di [23 API](23-api-publik-webhook.md). Endpoint internal yang dipakai SPA sudah ada sejak Fase 1 di bawah `/api/v1` dengan anotasi Swagger (`make swag`); fitur ini adalah permukaan publik yang stabil.

### WF-042 Forge/Connect workflow extension points

- **Deskripsi:** Titik ekstensi resmi agar app dapat menambahkan condition, validator, dan post function custom ke workflow.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app (Integrator/Developer API), Admin Situs yang memasang app.
- **Aturan perilaku (acceptance criteria):**
  1. Model ekstensi menyediakan tiga jenis modul: workflow condition, workflow validator, dan workflow post function, masing-masing berbasis ekspresi atau pemanggilan fungsi app (setara `jira:workflowCondition`, `jira:workflowValidator`, `jira:workflowPostFunction`).
  2. Model kompatibilitas setara Connect (`jiraWorkflowPostFunctions` serta condition/validator berbasis ekspresi) dipetakan ke tiga jenis modul yang sama.
  3. App menyediakan halaman konfigurasi rule (create/edit/view) yang di-embed di panel rule editor.
  4. ⚠ Rule dari app yang sudah di-uninstall ditandai "tidak valid" di editor; condition dan validator tidak valid dianggap gagal (transition diblokir), sedangkan post function tidak valid dilewati dan dicatat.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Platform ekstensi di [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md). App Marketplace (JMWE, JSU, ScriptRunner) hanya contoh konsumen dan tidak dibangun.

### WF-043 Rovo agents in workflows

- **Deskripsi:** Agent AI dapat dipasang ke workflow lewat editor workflow, sehingga agent berjalan otomatis saat work item berpindah ke status tertentu (mis. merapikan deskripsi, memecah pekerjaan, menambahkan komentar analisis). Di produk ini dimodelkan sebagai post-transition action asinkron yang memanggil agent LLM.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs dan Admin Proyek yang berhak mengedit workflow; agent sebagai aktor sistem.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Agent ditambahkan dari editor workflow company-managed (WF-021) pada transition, dan dijalankan ketika transition menuju status tujuan selesai.
  2. Admin memilih agent bawaan produk atau agent custom yang didefinisikan organisasi.
  3. Agent berjalan setelah transition terjadi dengan context work item (fields, description, comments).
  4. Output agent berupa komentar atau pembaruan field di work item, dan hasilnya bisa ditinjau manusia.
  5. Jejak eksekusi agent terlihat di activity/history work item dengan agent sebagai pelaku.
  6. Hanya admin yang berhak mengedit workflow yang bisa memasang agent, dan agent tunduk pada permission aktor agent di project.
  7. ⚠ Pemakaian agent dibatasi kuota per site (mis. jumlah eksekusi per bulan) yang dapat dikonfigurasi; eksekusi di atas kuota dilewati dan dicatat.
  8. Agent tidak tersedia bila fitur AI dimatikan di level organisasi/site; konfigurasi agent yang ada dinonaktifkan tanpa dihapus.
  9. ⚠ Workflow team-managed tidak mendukung agent pada rilis awal fitur ini.
  10. ⚠ Kegagalan agent tidak membatalkan transition; galat dicatat di activity work item.
- **Team-managed vs company-managed:** Lihat butir 9.
- **Catatan & ketergantungan:** Rovo adalah layanan Atlassian lintas produk dan tidak direplikasi; produk ini menyediakan kapabilitas setara yang opsional dan provider LLM-agnostik (lihat modul AI di [35 JSM AI](35-jsm-ai-otomasi-api-admin-csm.md) dan [27 NFR](27-nfr-arsitektur-platform.md)).

## Fitur legacy (rujukan, tidak dibangun)

### WF-044 Legacy workflow editor: Diagram & Text mode

- **Deskripsi:** Editor legacy dengan toggle Diagram/Text; Text mode menampilkan tabel steps, linked status, dan outgoing transitions, beserta akses ke halaman conditions, validators, dan post functions. Tidak dibangun karena digantikan WF-021.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: toggle Diagram/Text di editor legacy.
  2. Tidak dibangun: daftar Steps dengan nama step, linked status, dan transition keluar (nama, id, tujuan); padanannya adalah daftar transition di panel WF-021.
  3. Konsep step terpisah dari status tidak dibangun: satu node di workflow produk ini selalu adalah status; data impor yang nama step-nya berbeda dari nama status memakai nama status.
  4. Tidak dibangun: Add Step dengan nama step dan linked status.
  5. Tidak dibangun: Add Transition dari step; padanannya "Add transition" di WF-021.
  6. Tidak dibangun: Edit/Delete step dan View properties step; properties status tersedia lewat WF-031.
  7. Tidak dibangun: halaman detail transition dengan tab Conditions, Validators, Post Functions, Triggers, dan Properties; padanannya panel samping WF-021.
  8. ID transition tetap ditampilkan di panel transition WF-021 karena dipakai REST API dan otomasi.
  9. Grouping conditions bersarang (AND/OR) tidak dibangun; editor produk ini hanya mendukung ALL/ANY satu tingkat (WF-027).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Editor legacy sedang dipensiunkan di Cloud dan tetap lengkap di Data Center. Lihat Di luar cakupan.

### WF-045 Import / export workflow (XML & workflow bundles)

- **Deskripsi:** Ekspor dan impor definisi workflow dalam format XML (OSWorkflow) atau bundle .jwb. Tidak dibangun di UI; alternatifnya REST bulk workflows (WF-041) dan importer modul 25.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: Export as XML (definisi OSWorkflow).
  2. Tidak dibangun di UI: Import from XML; importer data dari instance lain (modul 25) boleh membaca XML workflow sebagai sumber migrasi.
  3. Tidak dibangun: Export/Import Workflow bundle (.jwb) berisi layout, screens, dan custom fields.
  4. Tidak dibangun: impor workflow bundle dari Atlassian Marketplace (sudah dihentikan di Cloud).
  5. Bila importer modul 25 membaca workflow, mapping status meminta admin memilih status existing atau membuat status baru.
  6. Bila importer modul 25 membaca workflow, rule dari app yang tidak terpasang dihapus disertai peringatan.
  7. Alternatif yang dibangun: REST API Workflows (bulk create/update/search, WF-041) dan alat migrasi di [25 Integrasi](25-integrasi-impor-ekspor-ekstensi.md).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat Di luar cakupan.

## Penyelarasan status lintas tipe project

### WF-046 ⚠ Global statuses in team-managed workflows

- **Deskripsi:** Workflow team-managed dapat memakai status global dari katalog Admin Situs (bukan hanya status ber-scope project), sehingga tim tetap mengatur alurnya sendiri sementara organisasi mendapat status standar yang sama untuk pelaporan dan JQL lintas space. Fitur ini ditambahkan di luar inventaris berdasarkan pengumuman "Align team workflows with company-wide standards" (Oktober 2026).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (Administer projects pada project team-managed) menambahkan dan memetakan status global; Admin Situs (Administer Jira) mengelola katalog status global (WF-003) dan Space Policy; pengguna dengan Transition issues menjalankan transition.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Editor workflow team-managed (WF-017) menyediakan aksi "Add status" dengan pilihan mencari status global dari katalog Admin (pencarian nama substring, tidak peka huruf besar/kecil, menampilkan lozenge kategori); status global yang dipilih ditambahkan tanpa membuat salinan ber-scope project dan tidak bisa di-rename, diubah kategorinya, atau dihapus dari editor team-managed (hanya dilepas dari workflow).
  2. ⚠ Kolom board team-managed dapat dipetakan ke status global sama seperti status project, termasuk satu kolom memuat campuran status global dan status project (WF-012, BRD-025).
  3. ⚠ Transition dari issue view (WF-011), board, dan bulk transition ke atau dari status global berjalan sama seperti status project; riwayat dan event (WF-014, WF-016) mencatat `status_id` global.
  4. ⚠ Tim tetap mengatur transition, rules (WF-018), dan mapping kolom secara mandiri; memakai status global tidak memindahkan kepemilikan workflow ke Admin Situs dan tidak mewajibkan workflow scheme.
  5. ⚠ JQL lintas space `status = X` (nama atau ID status global) mengembalikan work item company-managed dan team-managed yang memakai status global yang sama; status project bernama sama tetap entri terpisah dan tidak ikut kecuali ditulis dengan ID-nya atau nama dicocokkan seperti perilaku WF-001 butir 5.
  6. ⚠ Admin Situs dapat menegakkan "wajib status global" lewat Space Policy (diusulkan sebagai PRJ-069 di modul 04): bila aktif untuk project team-managed, "Add status" hanya menawarkan status global, status project existing ditandai dan diberi wizard pemetaan ke status global (memakai mekanisme migrasi WF-036), dan percobaan membuat status project ditolak dengan pesan yang menyebut kebijakan tersebut.
  7. ⚠ Status global yang dipakai workflow team-managed dihitung sebagai "dipakai" di halaman Statuses (WF-003 butir 5 dan 7), sehingga tidak bisa dihapus dan panel pemakaiannya mencantumkan project team-managed terkait.
- **Team-managed vs company-managed:** Khusus team-managed; company-managed sejak awal hanya memakai status global (WF-001).
- **Catatan & ketergantungan:** Model data tidak berubah: tabel `statuses` dengan `scope_project_id` null sudah mengakomodasi (WF-001); workflow team-managed cukup boleh merujuk status `GLOBAL`. Space Policy belum ada di [04 Proyek](04-proyek.md) (ID PRJ-069 masih usulan; keputusan terbuka). JQL di [12 JQL](12-jql-bahasa-kueri.md); board di [10 Board](10-board-kolom-kartu.md). Ditambahkan di luar inventaris (added_beyond_inventory).

## Pertanyaan terbuka

- Apakah halaman admin Resolutions dimiliki modul ini (WF-010) atau modul 05 (CFG-055), dan modul mana yang memegang entitas serta endpoint-nya?
- Apakah model draft workflow legacy (draft terpisah, backup copy saat publish) dibangun untuk kompatibilitas impor, atau cukup editor berbasis sesi dengan optimistic locking (WF-022)?
- Berapa angka batas default yang dipakai: panjang nama status (60), jumlah status per workflow team-managed (30), status/transition per workflow company-managed (100/500), workflow per site (2000), dan work item per bulk transition (1000) (WF-001, WF-017, WF-025, WF-013)?
- Bahasa ekspresi apa yang dipakai untuk condition dan validator build-your-own serta titik ekstensi app: subset ekspresi kompatibel Jira expressions, CEL, atau bahasa lain (WF-027, WF-028, WF-042)?
- Apakah work item yang masuk status berkategori Done di company-managed tanpa resolution perlu dicegah secara default, atau cukup peringatan di editor seperti perilaku Jira (WF-009)?
- Apakah fitur agent AI pada transition dibangun sama sekali, dan bila ya, dengan provider LLM apa serta model kuota apa (WF-043)?
- Apakah simplified workflow (board-managed) perlu dibangun, mengingat team-managed sudah memberi pengalaman serupa (WF-026)?
- Apakah selama migrasi status project dikunci sepenuhnya atau hanya work item terdampak (WF-036)?
- Apakah Space Policy "wajib status global" dibangun sebagai fitur baru modul 04 (usulan PRJ-069), dan bagaimana nasib status project existing saat kebijakan diaktifkan: migrasi wajib atau hanya larangan membuat status baru (WF-046)?

## Di luar cakupan & ekosistem

- **Fitur Data Center-only/deprecated (hanya rujukan, tidak dibangun):** Legacy workflow editor Diagram & Text mode beserta konsep step terpisah dari status dan grouping conditions bersarang (WF-044); Import/export workflow XML (OSWorkflow) dan workflow bundle .jwb, termasuk impor bundle dari Marketplace (WF-045); post function "Set issue security level based on user's project role" dan "Create Perforce Job" (WF-029).
- **App Marketplace pihak ketiga (tidak dibangun sebagai fitur native):** rule workflow tambahan dari JMWE, JSU, ScriptRunner, dan sejenisnya (condition/validator/post function scripted, grouping lanjutan, linked transitions). Kebutuhan ini dilayani lewat condition/validator berbasis ekspresi bawaan (WF-027, WF-028), titik ekstensi app (WF-042), dan otomasi (modul 20).
- **Layanan lintas produk Atlassian:** Rovo dan Rovo Studio sebagai layanan agent (produk ini menyediakan agent AI opsional sendiri, WF-043), kuota Atlassian Intelligence/AI credits, Atlassian Marketplace, dan Jira Cloud Migration Assistant (diganti importer modul 25).
- **Didefinisikan di modul lain:** workflow triggers DevOps dan smart commits (modul 24), approvals pada status workflow (modul 26 dan 32), condition dan workflow khusus service project (modul 29 dan 32), notification scheme (modul 19), screen dan screen scheme (modul 05), board dan kolom (modul 10), action dan trigger otomasi (modul 20).
