# 15 — Rilis (Versi) & Komponen

Modul ini mendefinisikan dua entitas pengelompokan per proyek di produk ini: **Version** (di UI disebut "Releases") yang mewakili titik rilis software atau milestone, dan **Component** yang mewakili sub-bagian proyek (modul, area, tim) untuk mengelompokkan work item dan menentukan default assignee. Cakupannya meliputi atribut dan siklus hidup version (create, edit, reorder, Release/Unrelease, Archive/Unarchive, merge, delete dengan migrasi nilai), dialog rilis dengan penanganan work item unresolved, halaman Releases dan version detail page (progress bar, Driver, Approvers, Related work, Release warnings), pembuatan release notes (manual dan AI opsional), field Fix versions, Affects versions, dan custom field Version Picker, representasi version di backlog, board Kanban, Timeline, Plans, dan laporan, trigger/action otomasi, REST API dan webhook, aturan permission, serta seluruh pengelolaan component (field, halaman, default assignee, delete dengan migrasi, API, katalog komponen eksternal). Istilah API dan JQL tetap memakai "version" (`fixVersion`, `affectedVersion`) meskipun UI memakai "Releases". Path REST produk referensi (`/rest/api/3/...`) dicantumkan sebagai target kompatibilitas; produk ini mengeksposnya di bawah `/api/v1/` dengan bentuk payload setara (keputusan akhir di modul 23). Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga tabel `versions`, `components`, tabel relasi work item ↔ version (Fix/Affects) dan work item ↔ component, serta endpoint terkait adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/release` dan `internal/features/component`) mengikuti pola Clean Architecture repo, dengan halaman Releases, version detail, dan Components sebagai route baru di SPA `frontend/`.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Administer Projects & Browse Projects (PERM-020), Assignable User (PERM-022), Resolve Issues (PERM-023), Manage Versions (PERM-026), View Development Tools (PERM-027), penegakan issue security (PERM-044).
- [04 Proyek](04-proyek.md) — Features toggle Releases/Components (PRJ-049, PRJ-050), Versions & Components di project settings (PRJ-055, PRJ-056), delete/archive project.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — katalog system field (CFG-011), Version picker custom field (CFG-032), field development info (CFG-023), field configuration & screens.
- [06 Work Item Inti](06-work-item-inti.md) — Move Work Item Wizard (ISS-032), Bulk Change (ISS-035), Restrict Work Item (ISS-042), web/remote link (ISS-030).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — History tab (CMT-022) dan model changelog (CMT-023).
- [10 Board](10-board-kolom-kartu.md) — board filter dan sub-filter Kanban (BRD-010, BRD-012), Kanban board UI (BRD-013), quick filter dan filter toolbar.
- [11 Backlog & Sprint](11-backlog-sprint.md) — Versions Panel in Backlog (SPR-009), Release from Board (SPR-025).
- [12 JQL](12-jql-bahasa-kueri.md) — field `fixVersion`, `affectedVersion`, `component` dan fungsi `unreleasedVersions()`, `releasedVersions()`, `earliestUnreleasedVersion()`, `latestReleasedVersion()`, `componentsLeadByUser()`.
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — issue navigator yang dibuka dari tautan jumlah work item.
- [16 Laporan & Insight](16-laporan-insight.md) — Version report dan Release burndown.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget statistik per component/version.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — marker release di Timeline dan cross-project release di Plans.
- [20 Otomasi](20-otomasi.md) — trigger dan action version, smart value `{{version.*}}`.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — konvensi REST, paginasi, expand, dan pengiriman webhook.
- [24 Integrasi DevOps](24-integrasi-devops.md) — data branch, commit, PR, build, deployment untuk Release warnings dan kolom Development.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — integrasi wiki eksternal untuk Related work dan release notes, katalog komponen eksternal.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — konsistensi operasi massal (merge/delete/release), paginasi, fitur AI opsional.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Version, Component, relasi work item, dan katalog event `version.*`/`component.*`.

## Persona terkait

- **Admin Proyek/Project Lead** — membuat, mengedit, mengurutkan, merge, menghapus version; mengelola component, component lead, dan default assignee; mengatur siapa yang boleh mengelola rilis.
- **Scrum Master/Team Lead** — merencanakan isi rilis lewat versions panel backlog, merilis dari board Kanban, memantau Release burndown.
- **Developer/Anggota Tim** — mengisi Fix versions/Affects versions/Components pada work item, melihat progress dan Release warnings sebelum rilis.
- **Product Owner/PM** — menetapkan tanggal rilis, menyusun release notes, menautkan related work, mengoordinasikan rilis lintas proyek.
- **Stakeholder/Viewer** — membaca halaman Releases, version detail page, dan release notes tanpa mengubahnya.
- **Integrator/Developer API** — menyinkronkan version dan component dari sistem CI/CD atau katalog layanan lewat REST API dan webhook.
- **Admin Situs** — mengaktifkan fitur AI, mengatur setting global "Allow unassigned issues", dan permission scheme.

## User story modul

- Sebagai Admin Proyek/Project Lead, saya ingin membuat version dengan nama, tanggal mulai, tanggal rilis, dan Driver agar setiap rilis punya identitas dan penanggung jawab yang jelas (REL-001, REL-002, REL-015).
- Sebagai Admin Proyek/Project Lead, saya ingin mengurutkan version dengan drag-and-drop agar picker, versions panel, dan fungsi JQL berbasis urutan mencerminkan roadmap rilis (REL-003).
- Sebagai Product Owner/PM, saya ingin menandai version sebagai Released dan memindahkan work item yang belum selesai ke version berikutnya agar tidak ada pekerjaan yang tertinggal tanpa target (REL-004, REL-005).
- Sebagai Admin Proyek/Project Lead, saya ingin menghapus atau menggabungkan version dengan memindahkan nilai Fix/Affects versions agar data work item tetap konsisten (REL-006, REL-007).
- Sebagai Developer/Anggota Tim, saya ingin mengisi Fix versions dan Affects versions pada work item agar jelas di rilis mana work item dikirim dan di rilis mana bug ditemukan (REL-009, REL-010).
- Sebagai Stakeholder/Viewer, saya ingin melihat halaman Releases dengan progress bar dan penanda Overdue agar saya tahu rilis mana yang berisiko terlambat (REL-012, REL-013).
- Sebagai Developer/Anggota Tim, saya ingin version detail page menampilkan work item, status development, dan Release warnings agar rilis bermasalah tercegah sebelum tombol Release ditekan (REL-014, REL-018).
- Sebagai Product Owner/PM, saya ingin menghasilkan release notes dari isi version, termasuk ringkasan AI opsional, agar komunikasi rilis tidak perlu disusun manual (REL-019, REL-020).
- Sebagai Scrum Master/Team Lead, saya ingin men-drag work item ke version di backlog dan merilis kolom terakhir board Kanban agar perencanaan dan rilis berkelanjutan cepat (REL-021, REL-022).
- Sebagai Product Owner/PM, saya ingin membuat release lintas beberapa proyek di Plans agar rilis program terkoordinasi (REL-024).
- Sebagai Integrator/Developer API, saya ingin mengelola version dan menerima webhook `version_released` agar pipeline deployment dan otomasi tersinkron (REL-026, REL-027).
- Sebagai Admin Proyek/Project Lead, saya ingin membuat component dengan component lead dan default assignee agar work item baru otomatis diterima orang yang tepat (REL-028, REL-030, REL-031).
- Sebagai Developer/Anggota Tim, saya ingin memilih component pada work item dan memfilter board/JQL per component agar pekerjaan per area mudah dipantau (REL-028, REL-029).
- Sebagai Admin Proyek/Project Lead, saya ingin menghapus component sambil memindahkan work item ke component lain agar klasifikasi lama tidak hilang (REL-033).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 6 |
| P1 | 16 |
| P2 | 9 |
| P3 | 5 |
| **Total** | **36** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 22 |
| Fase 3 | 8 |
| Fase 4 | 5 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: seluruh modul dimulai di Fase 2 karena versi/komponen termasuk paket "konfigurabilitas & visibilitas tim", tetapi skema data harus dirancang sejak awal Fase 2 agar referensi work item ke version/component berbasis ID (rename tercermin otomatis) dan setiap perubahan menulis history serta event domain (modul 28). Urutan Fase 2: entitas Version dan Component beserta penegakan permission (REL-001, REL-008, REL-028), field Fix versions/Affects versions/Components dan JQL dasarnya (REL-009, REL-010, REL-029), lalu halaman Releases, dialog rilis, delete dengan migrasi, versions panel, dan release dari Kanban (REL-012, REL-005, REL-007, REL-021, REL-022), kemudian halaman Components serta default assignee (REL-031, REL-032), dan terakhir REST API + webhook version dan component (REL-027, REL-034) yang selaras dengan API-039 dan API-038 di modul 23. Fase 3 menambah operasi dan visibilitas lanjutan yang bergantung pada modul lain: merge, custom field Version Picker, Driver, Related work, release notes, Version report, otomasi, serta Timeline/filter (modul 05, 16, 18, 20). Fase 4 berisi fitur yang bergantung pada integrasi DevOps, Plans, AI, atau katalog eksternal (Release warnings, AI release notes, cross-project release, katalog komponen eksternal) ditambah Release approvers (P3) yang ditunda karena nilainya kecil sebelum alur rilis lain matang. Archive components hanya dicatat sebagai rujukan karena berasal dari Data Center.

## Fitur

## Model & siklus hidup version

### REL-001 Versions (Releases) entity

- **Deskripsi:** Entitas version/release milik satu proyek yang mewakili titik rilis software atau milestone; dipakai oleh field Fix versions dan Affects versions, halaman Releases, laporan, dan Timeline. UI menyebutnya "Releases", sedangkan API dan JQL memakai "version".
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola), release manager, Developer/Anggota Tim (memakai di work item); melihat memerlukan Browse Projects; mengelola diatur REL-008.
- **Aturan perilaku (acceptance criteria):**
  1. Version memiliki atribut: Name (wajib, maks 255 karakter, unik per proyek), Description (opsional), Start date (opsional), Release date (opsional), Driver (user, opsional), Approvers (daftar), status (Unreleased/Released/Archived), sequence (urutan dalam proyek), dan flag overdue yang dihitung otomatis.
  2. ⚠ Description dibatasi 16.000 karakter; input yang lebih panjang ditolak dengan pesan "Deskripsi version terlalu panjang" (batas final mengikuti produk referensi setelah dipastikan).
  3. Validasi menolak Release date yang lebih awal dari Start date dengan pesan "Tanggal rilis tidak boleh sebelum tanggal mulai"; Name kosong ditolak dengan pesan "Nama version wajib diisi".
  4. ⚠ Nama duplikat dalam satu proyek ditolak dengan pesan "Version dengan nama ini sudah ada di proyek"; pembandingan nama bersifat case-insensitive (mis. "v1.0" dan "V1.0" dianggap sama) dan mengabaikan spasi di awal/akhir.
  5. Version dimiliki tepat satu proyek dan tidak dapat dipindahkan ke proyek lain; release lintas proyek hanya tersedia melalui Plans (REL-024).
  6. Flag overdue bernilai benar bila Release date sudah lewat (dibandingkan dengan tanggal hari ini di zona waktu site) dan status masih Unreleased; lozenge "Overdue" tampil di daftar Releases, picker, dan version detail page, dan hilang begitu version dirilis atau Release date diubah ke masa depan.
  7. Sistem menyediakan atribut turunan: jumlah work item per status category (To Do/In Progress/Done) untuk Fix version, jumlah work item unresolved, jumlah work item yang memakai version sebagai Fix version dan sebagai Affects version, serta jumlah Release warning (REL-018).
  8. Setiap version memiliki ID numerik global yang dipakai di URL version detail page dan REST API; work item mereferensikan version lewat ID sehingga rename langsung tercermin di semua work item tanpa migrasi data.
  9. ⚠ Company-managed software project: fitur version selalu tersedia. Team-managed: halaman Releases dan field Fix versions hanya muncul setelah feature "Releases" dinyalakan di Project settings > Features (PRJ-049); mematikan toggle menyembunyikan halaman dan field tanpa menghapus data version maupun nilai di work item, dan menyalakannya kembali memulihkan tampilan.
  10. ⚠ Proyek bisnis (padanan eks Work Management) tidak menampilkan Releases secara default; produk ini menyediakan toggle Features yang sama agar admin proyek bisnis dapat mengaktifkannya (keputusan terbuka).
  11. Menghapus proyek ikut menghapus semua version-nya; mengarsipkan proyek ikut menyembunyikan version-nya dari picker dan pencarian lintas proyek, dan memulihkan proyek memulihkan version dengan status sebelumnya.
- **Team-managed vs company-managed:** Company-managed selalu memiliki version; team-managed memerlukan feature Releases aktif (butir 9).
- **Catatan & ketergantungan:** Skema tabel dan event `version.created/updated/released/unreleased/moved/merged/deleted` didefinisikan di modul 28. Fungsi JQL berbasis version ada di modul 12. Konfigurasi Versions di project settings (PRJ-055) merujuk modul ini sebagai sumber detail.

### REL-002 Create version

- **Deskripsi:** Membuat version baru dari halaman Releases, versions panel backlog, field Fix versions, dialog release Kanban, otomasi, atau REST API.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, release manager; memerlukan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Create version" di halaman Releases membuka dialog berisi Name, Start date, Release date, Description, dan Driver; tombol simpan nonaktif sampai Name terisi.
  2. Versions panel backlog (Scrum dan Kanban backlog, company-managed) menyediakan create inline dengan field yang sama (REL-021, SPR-009).
  3. ⚠ Mengetik nama yang belum ada di field Fix versions/Affects versions menampilkan opsi "Create version '<nama>'" yang langsung membuat version Unreleased; opsi ini hanya tampil bagi pengguna yang punya izin mengelola version, dan di team-managed mengikuti role Administrator.
  4. Date picker mengikuti format tanggal dan locale pengguna (modul 21); validasi Start date <= Release date dijalankan di klien dan di server.
  5. Version baru ditempatkan di akhir urutan (sequence terakhir) proyek.
  6. Server menolak pembuatan oleh pengguna tanpa Administer Projects atau Manage Versions (company-managed, REL-008) atau role Administrator (team-managed) dengan HTTP 403, terlepas dari jalur UI yang dipakai.
  7. Setiap pembuatan menerbitkan event "version created" yang dipakai webhook `jira:version_created`-kompatibel dan trigger otomasi Version created (REL-026, REL-027).
- **Team-managed vs company-managed:** Company-managed memakai Administer Projects (atau Manage Versions, PERM-026); team-managed memakai role Administrator.
- **Catatan & ketergantungan:** Pembuatan dari dialog release Kanban ada di REL-022; dari otomasi di REL-026.

### REL-003 Edit and reorder versions

- **Deskripsi:** Mengubah atribut version dan menyusun ulang urutannya; urutan memengaruhi picker, versions panel, laporan, dan fungsi JQL berbasis urutan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, release manager dengan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Edit tersedia lewat dialog Edit di menu "..." halaman Releases dan dari header version detail page (Name, Start date, Release date, Description, Driver); edit inline langsung di sel tabel Releases tidak disediakan di rilis pertama.
  2. Drag handle di halaman Releases memungkinkan reorder dengan drag-and-drop; urutan baru disimpan sebagai sequence dan berlaku untuk semua pengguna.
  3. API menyediakan reorder setara `POST /rest/api/3/version/{id}/move` dengan parameter `after` (URL version lain) atau `position` (`First`, `Last`, `Earlier`, `Later`); mengirim keduanya sekaligus ditolak dengan HTTP 400.
  4. Urutan menentukan hasil `earliestUnreleasedVersion()` dan `latestReleasedVersion()`, urutan opsi di picker, dan urutan di versions panel.
  5. Start date dan Release date boleh dikosongkan saat edit; validasi Start date <= Release date tetap berlaku bila keduanya terisi.
  6. Rename langsung tercermin di semua work item; filter tersimpan atau JQL yang memakai nama version lama dapat menjadi tidak cocok, sehingga dialog Edit menampilkan peringatan "Filter yang memakai nama lama tidak ikut diperbarui".
  7. Edit menerbitkan event webhook `jira:version_updated` dan reorder menerbitkan `jira:version_moved`.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber izin (REL-008).
- **Catatan & ketergantungan:** Penyimpanan JQL berbasis ID vs nama adalah keputusan modul 12.

### REL-004 Version statuses: Release, Unrelease, Archive, Unarchive

- **Deskripsi:** Siklus hidup version: Unreleased menjadi Released, serta Archived untuk menyembunyikannya dari pemilihan; setiap transisi dapat dibalik.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Release manager, Admin Proyek dengan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Unreleased adalah status default version baru; version Unreleased tampil di picker Fix versions/Affects versions dan di versions panel.
  2. Released ditandai dengan tanggal rilis (REL-005); version Released tetap tampil di picker dalam kelompok "Released versions".
  3. Archived disembunyikan dari picker dan versions panel; work item tetap menyimpan nilainya, version tidak dapat dipilih lagi untuk work item baru atau work item lain, tetapi nilai yang sudah ada boleh dihapus dari work item.
  4. Aksi Unrelease di menu "..." mengembalikan Released menjadi Unreleased dan mempertahankan Release date yang ada.
  5. Aksi Archive/Unarchive tersedia di menu "..."; version Archived tetap dapat dicari via JQL (`fixVersion = "X"`) dan tampil di halaman Releases bila filter status "Archived" dipilih; Unarchive mengembalikan version ke status sebelum diarsipkan (Released atau Unreleased).
  6. ⚠ Version Archived tidak dihitung dalam hasil `unreleasedVersions()` maupun `releasedVersions()`.
  7. ⚠ Release menerbitkan webhook `jira:version_released`, Unrelease menerbitkan `jira:version_unreleased`, sedangkan Archive/Unarchive dikirim sebagai `jira:version_updated` dengan atribut `archived` berubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status lozenge di UI: UNRELEASED, RELEASED, ARCHIVED, plus penanda OVERDUE (REL-001, REL-012).

### REL-005 Release version dialog (unresolved issues handling)

- **Deskripsi:** Dialog untuk menandai version sebagai Released, termasuk penanganan work item yang belum selesai di dalam version tersebut.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Release manager dengan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog berisi field Release date dengan default hari ini (atau Release date yang sudah tersimpan bila ada) dan dapat diubah.
  2. Bila version memiliki work item unresolved, dialog menampilkan jumlahnya dan dua opsi: "Ignore and proceed" atau "Move issues to" dengan dropdown version Unreleased lain di proyek yang sama; dropdown tidak memuat version yang sedang dirilis maupun version Archived.
  3. Opsi pemindahan mengganti nilai Fix versions dari version yang dirilis menjadi version target hanya untuk work item unresolved; nilai Fix versions lain pada work item dipertahankan, dan work item resolved tidak berubah.
  4. ⚠ Work item dianggap unresolved bila field Resolution kosong; kebijakan alternatif berbasis status category Done menjadi keputusan terbuka dan, bila diadopsi, berlaku seragam untuk dialog ini, progress bar, dan API `unresolvedIssueCount`.
  5. Dialog menyediakan tautan untuk membuka daftar work item unresolved di issue navigator (modul 13) sebelum rilis.
  6. Release warnings (PR open, build gagal) ditampilkan di version detail page sebelum tombol Release ditekan (REL-018); dialog menampilkan ringkasan jumlah warning bila ada, tanpa memblokir rilis.
  7. ⚠ Pemindahan massal tidak mengirim email per work item; perubahan dicatat di History tiap work item dan satu event ringkas diterbitkan untuk notifikasi/otomasi.
  8. API menyediakan rilis setara `PUT /rest/api/3/version/{id}` dengan `released=true`, `releaseDate`, dan `moveUnfixedIssuesTo`; pemindahan berjalan dalam satu transaksi sehingga kegagalan membatalkan rilis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai juga oleh SPR-025 dan REL-022 untuk penanganan unresolved; operasi massal mengikuti batas konsistensi modul 27.

### REL-006 Merge versions

- **Deskripsi:** Menggabungkan version sumber ke version target sehingga semua nilai work item dipindahkan dan version sumber dihapus.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek dengan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Merge di menu "..." halaman Releases membuka dialog untuk memilih version sumber dan version target dari proyek yang sama; sumber dan target yang sama ditolak.
  2. Semua nilai Fix versions dan Affects versions yang menunjuk sumber diganti dengan target; work item yang sudah memiliki target tidak mendapat nilai duplikat.
  3. Version sumber dihapus setelah merge dan operasi tidak dapat di-undo.
  4. Dialog konfirmasi menampilkan jumlah work item terdampak dan teks "Tindakan ini tidak dapat dibatalkan" sebelum eksekusi.
  5. API menyediakan merge setara `PUT /rest/api/3/version/{id}/mergeto/{moveIssuesTo}`.
  6. ⚠ Merge memerlukan Administer Projects (atau Manage Versions) di company-managed; di team-managed aksi Merge disediakan bagi role Administrator.
  7. Merge menerbitkan event webhook `jira:version_merged` dengan ID sumber dan target.
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** Nilai custom field Version Picker yang menunjuk sumber ikut diganti dengan target (CFG-032, REL-011).

### REL-007 Delete version

- **Deskripsi:** Menghapus version dengan opsi migrasi nilai Fix versions, Affects versions, dan custom field Version Picker pada work item terkait.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek dengan izin mengelola version (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Delete menampilkan jumlah work item yang memakai version sebagai Fix version dan sebagai Affects version.
  2. Dialog menyediakan opsi terpisah untuk Fix version dan Affects version: pindahkan ke version lain (dropdown version proyek yang sama, tidak termasuk version yang dihapus) atau hapus nilai dari work item.
  3. Setiap custom field bertipe Version Picker yang memakai version tersebut juga mendapat opsi pengganti sendiri.
  4. Bila tidak ada work item terkait, dialog hanya menampilkan konfirmasi sederhana lalu menghapus.
  5. Penghapusan bersifat permanen; dialog memperingatkan bahwa filter JQL yang mereferensikan version dapat menjadi invalid.
  6. API menyediakan penghapusan setara `POST /rest/api/3/version/{id}/removeAndSwap` (`moveFixIssuesTo`, `moveAffectedIssuesTo`, `customFieldReplacementList`); padanan `DELETE /version/{id}` gaya lama disediakan hanya sebagai kompatibilitas dan ditandai deprecated.
  7. API menyediakan penghitungan setara `GET /version/{id}/relatedIssueCounts` untuk menampilkan jumlah work item terdampak sebelum penghapusan.
  8. Penghapusan menerbitkan event webhook `jira:version_deleted`.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber izin (REL-008).
- **Catatan & ketergantungan:** Penggantian nilai dicatat di History setiap work item (CMT-023) dan berjalan transaksional (modul 27).

### REL-008 Release and component management permissions

- **Deskripsi:** Aturan permission yang mengatur siapa yang boleh mengelola version/component dan siapa yang boleh mengisi field terkait.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs (permission scheme); seluruh pengguna sebagai subjek pemeriksaan izin.
- **Aturan perilaku (acceptance criteria):**
  1. Membuat, mengedit, merilis, merge, mengarsipkan, dan menghapus version memerlukan Administer Projects (company-managed) atau role Administrator (team-managed); server menegakkan aturan ini di semua endpoint.
  2. Membuat, mengedit, dan menghapus component memerlukan Administer Projects atau peran project admin.
  3. Mengisi Fix versions memerlukan Resolve Issues dan Edit Issues (company-managed); mengisi Affects versions dan Components cukup dengan Edit Issues.
  4. Release dari board Kanban memerlukan Administer Projects di semua proyek yang tercakup filter board (REL-022).
  5. Melihat halaman Releases dan Components cukup dengan Browse Projects; aksi pengelolaan disembunyikan dari pengguna yang tidak berizin.
  6. ⚠ Produk ini menyediakan project permission terpisah "Manage Versions" (PERM-026) sebagai alternatif Administer Projects untuk seluruh aksi pada butir 1; pemegang Administer Projects tetap dapat mengelola version tanpa izin ini.
- **Team-managed vs company-managed:** Company-managed memakai permission scheme; team-managed memakai role Administrator/Member.
- **Catatan & ketergantungan:** Definisi permission ada di PERM-020, PERM-023, PERM-026 (modul 03); field-level aturan di REL-009, REL-010, REL-029.

## Field version pada work item

### REL-009 Fix versions field

- **Deskripsi:** System field multi-value yang menyatakan version tempat work item akan atau sudah dirilis (`fixVersion` di JQL dan API).
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, release manager; mengubah nilai memerlukan Resolve Issues + Edit Issues (company-managed) atau role Member ke atas (team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Field berupa multi-select version picker yang hanya menawarkan version milik proyek work item; version proyek lain ditolak oleh server dengan pesan "Version bukan milik proyek ini".
  2. Picker mengelompokkan opsi menjadi "Unreleased versions" dan "Released versions" sesuai urutan sequence; version Archived tidak ditawarkan, tetapi nilai Archived yang sudah ada tetap ditampilkan dan dapat dihapus.
  3. Field dapat diedit di issue view, create dialog, bulk edit (Add to existing / Replace all / Remove / Clear all, ISS-035), dan dengan men-drag work item ke versions panel backlog (REL-021).
  4. Company-managed: hanya pengguna dengan Resolve Issues dan Edit Issues yang dapat mengubah Fix versions; tanpa Resolve Issues field tampil read-only dan API mengembalikan HTTP 403 untuk perubahan field ini.
  5. Field dapat dijadikan required atau hidden melalui field configuration dan screens (company-managed, modul 05).
  6. Setiap penambahan/penghapusan nilai tercatat di History work item dengan nilai lama dan baru (CMT-022).
  7. Saat work item dipindahkan ke proyek lain, Move wizard (ISS-032) meminta pemetaan setiap Fix version ke version proyek tujuan atau mengosongkannya; nilai tidak pernah menunjuk version proyek asal setelah move.
  8. Issue view menampilkan status rilis setiap nilai version (Released, Unreleased, atau Overdue) sebagai lozenge di samping nama.
  9. JQL mendukung `fixVersion = X`, `fixVersion in unreleasedVersions()`, `releasedVersions()`, `earliestUnreleasedVersion()`, `latestReleasedVersion()`, dan `fixVersion is EMPTY` (modul 12).
  10. Field menjadi sumber data Version report, Release burndown, release dari Kanban, dan marker Timeline (REL-022, REL-023, REL-025).
  11. Team-managed: field tersedia di work item setelah feature Releases aktif (REL-001); sebelum itu field tidak tampil di layar maupun picker field.
- **Team-managed vs company-managed:** Company-managed memakai Resolve Issues + screens/field configuration; team-managed memerlukan feature Releases dan memakai role.
- **Catatan & ketergantungan:** Definisi di katalog system field (CFG-011); progress bar (REL-013) hanya menghitung Fix versions.

### REL-010 Affects versions field

- **Deskripsi:** System field multi-value untuk mencatat version tempat bug ditemukan (`affectedVersion` di JQL).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** QA, Developer/Anggota Tim, pelapor; mengubah nilai memerlukan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field berupa multi-select version picker yang memuat version Released maupun Unreleased milik proyek work item; version Archived tidak ditawarkan.
  2. Pada template company-managed, field ditampilkan secara default untuk work type Bug melalui screen scheme dan tidak ditampilkan untuk work type lain.
  3. Field dikonfigurasi lewat screens dan field configuration dan dapat dibuat required atau hidden.
  4. JQL mendukung `affectedVersion = X` dan `affectedVersion in releasedVersions()` beserta fungsi version lainnya.
  5. Nilai field ikut dimigrasi pada dialog Delete version (REL-007) dan Merge versions (REL-006).
  6. Bulk edit dan Move work item memetakan nilai field ini dengan aturan yang sama seperti Fix versions (REL-009).
  7. ⚠ Team-managed menyediakan Affects versions sebagai field opsional yang dapat ditambahkan ke work type setelah feature Releases aktif.
- **Team-managed vs company-managed:** Lihat butir 7; company-managed memakai screen scheme.
- **Catatan & ketergantungan:** Tidak memengaruhi progress bar (REL-013).

### REL-011 Version picker custom fields

- **Deskripsi:** Custom field bertipe Version untuk kebutuhan tambahan seperti "Target version" atau "Verified in version".
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (membuat custom field); pengisi memerlukan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia dua tipe: Version Picker (single version) dan Version Picker (multiple versions).
  2. Opsi diambil dari version proyek work item dengan pengelompokan Released/Unreleased seperti Fix versions.
  3. Nilai field ikut dipetakan saat delete version melalui `customFieldReplacementList` (REL-007) dan saat merge (REL-006).
  4. Field dapat dicari di JQL memakai nama custom field (mis. `"Target version" = "2.0"`) dan mendukung fungsi version.
  5. Tipe ini hanya tersedia di company-managed karena merupakan custom field global.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Definisi tipe field ada di CFG-032 (modul 05); modul ini hanya mengatur perilaku terhadap siklus hidup version.

## Halaman Releases & version detail

### REL-012 Releases page (version list)

- **Deskripsi:** Halaman "Releases" di sidebar proyek yang menampilkan semua version beserta status, progress, tanggal, dan aksi pengelolaan (dahulu disebut "Release hub").
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, release manager, Admin Proyek, Stakeholder/Viewer; melihat memerlukan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel menampilkan kolom Version, Status (lozenge UNRELEASED/RELEASED/ARCHIVED plus penanda OVERDUE), Progress bar, Start date, Release date, Description, dan Driver.
  2. Filter status multi-select berisi Unreleased, Released, Archived; default menampilkan Unreleased dan Released.
  3. Search box memfilter daftar berdasarkan nama version (substring, case-insensitive).
  4. Menu "..." per baris menyediakan Edit, Release/Unrelease, Archive/Unarchive, Merge, dan Delete sesuai status version.
  5. Drag handle per baris memungkinkan reorder (REL-003); drag dinonaktifkan saat search atau filter aktif agar urutan tidak ambigu.
  6. Klik nama version membuka version detail page (REL-014).
  7. Tombol "Create version" membuka dialog pembuatan (REL-002).
  8. Bila proyek belum punya version, halaman menampilkan empty state berisi penjelasan singkat fungsi Releases dan tombol "Create version" (hanya bagi pengguna berizin).
  9. Daftar dipaginasi (default 50 version per halaman) untuk proyek dengan banyak version.
  10. Hanya pengguna dengan Browse Projects yang dapat membuka halaman; bagi pengguna tanpa izin mengelola version, tombol Create, drag handle, dan menu "..." disembunyikan.
- **Team-managed vs company-managed:** Tersedia di software project company-managed dan team-managed (dengan feature Releases aktif).
- **Catatan & ketergantungan:** Halaman Versions di project settings (PRJ-055) dapat mengarah ke halaman ini.

### REL-013 Release progress bar

- **Deskripsi:** Visualisasi kemajuan version berdasarkan status category work item yang memiliki version tersebut sebagai Fix version.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Release manager, Developer/Anggota Tim; hanya work item yang boleh dilihat pengguna yang dihitung.
- **Aturan perilaku (acceptance criteria):**
  1. Progress bar terdiri atas segmen Done (hijau), In progress (biru), dan To do (abu-abu) dengan panjang proporsional terhadap jumlah work item per status category.
  2. Hover segmen menampilkan tooltip jumlah work item pada segmen tersebut (mis. "12 work item Done").
  3. Klik segmen membuka daftar work item terkait (issue navigator atau tab work item di detail page) dengan filter status category segmen tersebut.
  4. Hanya work item yang memiliki version sebagai Fix version yang dihitung; Affects versions diabaikan.
  5. Progress bar tampil di halaman Releases, versions panel backlog, dan header version detail page dengan angka yang sama.
  6. ⚠ Sub-task dihitung sebagai work item terpisah bila sub-task itu sendiri memiliki Fix version; sub-task tidak mewarisi Fix version parent untuk perhitungan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status category didefinisikan di modul 09; version tanpa work item menampilkan bar kosong dengan teks "Belum ada work item".

### REL-014 Version detail page

- **Deskripsi:** Halaman detail per version berisi ringkasan, daftar work item, status development, warnings, related work, approvers, release notes, dan aksi rilis.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Release manager, Developer/Anggota Tim, Stakeholder/Viewer; melihat memerlukan Browse Projects; aksi memerlukan izin mengelola version.
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan nama, status lozenge, Start/Release date, Description, Driver, tombol Release, dan menu "..." (Edit, Archive, Merge, Delete); tombol Release diganti Unrelease bila version sudah Released.
  2. Ringkasan menampilkan jumlah work item Done / In progress / To do dan tautan "View in Issue navigator" yang membuka JQL `fixVersion = <id>`.
  3. Daftar work item menampilkan key, summary, work type, priority, assignee, status, dan kolom Development (branch, commit, PR, build, deployment) bila integrasi DevOps tersedia dan pengguna punya View Development Tools.
  4. Daftar dapat difilter berdasarkan status category atau kategori warning (REL-018).
  5. Halaman memuat section Release warnings (REL-018), Related work (REL-017), dan Approvers (REL-016); section yang fiturnya belum aktif disembunyikan.
  6. Tombol "Release notes" membuka generator catatan rilis manual (REL-019) atau AI bila diaktifkan (REL-020).
  7. Company-managed menampilkan tautan ke Version report dan Release burndown (REL-025).
  8. Halaman memiliki deep link stabil `/projects/{key}/versions/{id}/tab/release-report-all-issues`; ID version yang tidak ada atau tidak dapat diakses menampilkan halaman 404 "Version tidak ditemukan".
- **Team-managed vs company-managed:** Butir 7 hanya company-managed.
- **Catatan & ketergantungan:** Kolom Development bergantung pada modul 24 (Fase 4); sebelum itu kolom disembunyikan.

### REL-015 Version driver

- **Deskripsi:** Field Driver pada version untuk menunjuk penanggung jawab rilis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Release manager, Admin Proyek dengan izin mengelola version.
- **Aturan perilaku (acceptance criteria):**
  1. Driver berupa user picker tunggal; avatar dan nama Driver tampil di daftar Releases dan version detail page.
  2. Driver bersifat opsional dan dapat dikosongkan kapan saja.
  3. Objek Version di REST API memuat field `driver` berupa account ID; nilai account ID yang tidak dikenal ditolak dengan HTTP 400.
  4. ⚠ Otomasi menyediakan smart value `{{version.driver}}` (nama, email, account ID) agar rule dapat menyebut atau mengirim notifikasi ke Driver.
  5. Driver bersifat informatif dan tidak memberikan permission tambahan apa pun terhadap version atau work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengguna yang dinonaktifkan tetap tampil sebagai Driver dengan label "nonaktif" (modul 02).

### REL-016 Release approvers

- **Deskripsi:** Daftar approver pada version detail page untuk mencatat persetujuan rilis.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Release manager (mengelola daftar approver, memerlukan izin mengelola version), approver (mengubah status miliknya sendiri).
- **Aturan perilaku (acceptance criteria):**
  1. Release manager dapat menambah approver melalui user picker beserta deskripsi peran/ruang lingkup persetujuan (mis. "QA sign-off"); user yang sama tidak dapat ditambahkan dua kali.
  2. Setiap approver memiliki status Pending, Approved, atau Declined; Declined dapat disertai alasan (`declineReason`).
  3. Hanya approver yang bersangkutan yang dapat mengubah status persetujuannya sendiri; pengguna lain melihat status read-only.
  4. Version detail page menampilkan ringkasan status approval (mis. "2 dari 3 menyetujui, 1 menolak").
  5. ⚠ Approval bersifat informatif dan tidak memblokir aksi Release; dialog Release menampilkan peringatan bila masih ada approver Pending atau Declined.
  6. ⚠ Objek Version REST memuat array `approvers` (accountId, description, status, declineReason) yang dapat dibaca dengan expand `approvers`; endpoint khusus untuk mengubah status approver disediakan di bawah `/api/v1/version/{id}/approvers`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dengan Approvals JSM (modul 32) dan approval proyek bisnis (modul 26); tidak memakai workflow.

### REL-017 Related work

- **Deskripsi:** Section di version detail page untuk menautkan pekerjaan pendukung rilis yang bukan work item (dokumen, desain, komunikasi, dsb.).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Release manager, Product Owner/PM; mengelola item memerlukan izin mengelola version; melihat cukup dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna menambah item dengan kategori (Communication, Design, Documentation, Software, Other), URL, dan Title; URL wajib berskema http/https.
  2. Bila integrasi wiki eksternal (padanan Confluence, modul 25) terhubung, pengguna dapat menautkan halaman wiki sebagai smart link atau membuat halaman baru dari template (mis. release notes, retrospective).
  3. ⚠ Item related work dapat ditautkan ke satu work item (field `issueId` di API) dan menampilkan key work item tersebut di kartu item.
  4. Item dapat diedit dan dihapus; penghapusan meminta konfirmasi.
  5. URL ditampilkan sebagai smart link dengan judul dan favicon yang diambil server-side; kegagalan pengambilan menampilkan URL mentah.
  6. API menyediakan operasi setara `GET/POST/PUT /rest/api/3/version/{id}/relatedwork` dan `DELETE /version/{versionId}/relatedwork/{relatedWorkId}`.
  7. Hanya pengguna yang boleh mengelola version yang dapat menambah, mengubah, atau menghapus item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Smart link mengikuti modul 07 (rendering) dan ISS-030 (remote link).

### REL-018 Release warnings

- **Deskripsi:** Peringatan otomatis di version detail page berdasarkan data development yang terhubung untuk mencegah rilis bermasalah.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Release manager, Developer/Anggota Tim; memerlukan View Development Tools (PERM-027).
- **Aturan perilaku (acceptance criteria):**
  1. Kategori "Open pull requests": work item sudah Done tetapi masih memiliki PR berstatus open.
  2. Kategori "Unreviewed code": work item memiliki commit yang belum menjadi bagian dari pull request/review.
  3. Kategori "Failing builds": work item yang build terakhirnya gagal.
  4. Kategori "Open reviews" untuk alat code review legacy hanya muncul bila integrasi semacam itu terhubung; produk ini tidak membangun integrasi khusus untuknya.
  5. Jumlah warning tampil di version detail page dan sebagai badge di daftar Releases.
  6. Klik kategori warning memfilter daftar work item di detail page ke work item yang terkena warning tersebut.
  7. Warning hanya muncul bila minimal satu development tool terhubung (modul 24) dan pengguna memiliki View Development Tools; selain itu section disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data PR/build/commit berasal dari CFG-023 dan modul 24.

## Release notes

### REL-019 Release notes generation

- **Deskripsi:** Pembuatan catatan rilis dari work item di sebuah version, melalui halaman klasik maupun generator di version detail page.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Release manager, Product Owner/PM; memerlukan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Release Notes klasik (company-managed): pengguna memilih version dan format HTML atau Text; work item dikelompokkan per work type; output dapat disalin.
  2. Halaman klasik dapat diakses lewat URL setara `/secure/ReleaseNote.jspa?projectId=...&version=...` yang dialihkan ke route SPA produk ini.
  3. Generator di version detail page memungkinkan memilih work type yang disertakan dan apakah key work item serta field tertentu (mis. summary, priority) ditampilkan.
  4. Output dapat disalin ke clipboard atau, bila integrasi wiki eksternal terhubung, dibuat sebagai halaman wiki baru.
  5. Release notes tidak disimpan sebagai entitas; setiap pembukaan selalu dibangkitkan ulang dari data saat ini.
  6. Hanya work item yang dapat dilihat pengguna yang disertakan (menghormati issue security, PERM-044).
- **Team-managed vs company-managed:** Butir 1-2 hanya company-managed; generator (butir 3) tersedia di keduanya.
- **Catatan & ketergantungan:** Konten berformat mengikuti renderer modul 07.

### REL-020 AI-generated release notes (Atlassian Intelligence / Rovo)

- **Deskripsi:** Pembuatan release notes berbasis AI yang merangkum work item di version menjadi narasi siap pakai; di produk ini disediakan sebagai fitur AI opsional, bukan layanan vendor referensi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / premium
- **Aktor & izin:** Release manager, Product Owner/PM; memerlukan fitur AI diaktifkan oleh Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Version detail page menyediakan opsi "Generate with AI" pada generator release notes.
  2. Ringkasan mengelompokkan isi menjadi fitur baru, bug fix, dan improvement berdasarkan work type dan isi work item.
  3. Hasil dapat diedit di editor sebelum disalin atau dibuat sebagai halaman wiki eksternal.
  4. Model hanya menerima data work item yang dapat diakses pengguna yang memicu (menghormati permission dan issue security).
  5. Opsi hanya tampil bila Admin Situs mengaktifkan fitur AI di pengaturan site (modul 22/27); bila nonaktif, opsi disembunyikan.
  6. Hasil menampilkan disclaimer konten AI serta tombol feedback (suka/tidak suka).
  7. ⚠ Pengguna dapat memilih tone/audiens (mis. teknis vs pelanggan) sebelum generate.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia model dan kebijakan data AI ditentukan di modul 27; layanan AI vendor referensi (Atlassian Intelligence/Rovo) tidak dibangun.

## Version di backlog, board, perencanaan & laporan

### REL-021 Versions panel (backlog)

- **Deskripsi:** Panel samping di backlog untuk melihat version, memfilter backlog, dan menetapkan Fix version dengan drag-and-drop.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Scrum Master/Team Lead, release manager; menetapkan Fix version mengikuti REL-009; membuat version mengikuti REL-008.
- **Aturan perilaku (acceptance criteria):**
  1. Backlog company-managed (Scrum dan Kanban backlog) memiliki toggle panel Versions yang dapat ditampilkan atau disembunyikan; status toggle diingat per pengguna per board.
  2. Panel menampilkan daftar version Unreleased dengan detail tanggal, jumlah work item, progress bar (REL-013), dan tautan ke version detail page.
  3. Klik version memfilter backlog dan sprint ke work item dengan Fix version tersebut; opsi "Issues without versions" menampilkan work item tanpa Fix version dan "All issues" menghapus filter.
  4. Men-drag satu atau beberapa work item ke version menambahkan version itu ke Fix versions tanpa menghapus nilai lain; pengguna tanpa Resolve Issues mendapat pesan galat dan nilai tidak berubah.
  5. Panel menyediakan create version inline (REL-002) bagi pengguna berizin.
  6. Pada board Kanban, panel hanya tampil bila Kanban backlog diaktifkan (SPR-024).
  7. ⚠ Team-managed tidak memiliki panel khusus; filter Release tersedia di toolbar backlog/board (REL-023).
- **Team-managed vs company-managed:** Lihat butir 7.
- **Catatan & ketergantungan:** Alias lintas modul: SPR-009 di modul 11 mendefinisikan sisi backlog; modul ini menjadi sumber aturan version.

### REL-022 Release from Kanban board

- **Deskripsi:** Merilis seluruh work item di kolom terakhir board Kanban menjadi sebuah version dan menghilangkannya dari board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Release manager, Scrum Master/Team Lead; memerlukan Administer Projects (atau Manage Versions) di semua proyek yang tercakup filter board.
- **Aturan perilaku (acceptance criteria):**
  1. Board Kanban company-managed menampilkan tombol "Release" di header kolom paling kanan; tombol nonaktif bila kolom kosong.
  2. ⚠ Dialog berisi Version name, Release date (default hari ini), dan Description; pengguna dapat mengetik nama baru (version dibuat otomatis) atau memilih version Unreleased yang sudah ada; nama baru yang bentrok dengan version yang sudah ada ditolak.
  3. Semua work item di kolom terakhir yang terlihat di board mendapat Fix version tersebut, dan version langsung berstatus Released.
  4. Work item yang dirilis hilang dari board karena sub-filter default board Kanban adalah `fixVersion in unreleasedVersions() OR fixVersion is EMPTY`; sub-filter dapat diubah di Board settings > General (BRD-012).
  5. Bila pengguna tidak memiliki Administer Projects (atau Manage Versions) di salah satu proyek pada filter board, tombol Release disembunyikan dan API menolak dengan HTTP 403.
  6. Alternatif tanpa rilis: setting board "Hide completed issues older than" menyembunyikan work item Done setelah periode tertentu (modul 10).
  7. ⚠ Board Kanban team-managed tidak menampilkan tombol Release; rilis dilakukan dari halaman Releases.
- **Team-managed vs company-managed:** Lihat butir 7.
- **Catatan & ketergantungan:** Alias lintas modul: SPR-025 (modul 11) dan aturan kolom Done board (modul 10). Untuk board multi-proyek, version dibuat di proyek mana adalah pertanyaan terbuka.

### REL-023 Releases in Timeline and board filters

- **Deskripsi:** Representasi version di Timeline serta penyaringan berdasarkan release di board dan backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Product Owner/PM; mengikuti akses board/Timeline.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Timeline menampilkan marker release pada Release date setiap version Unreleased dan Released; toolbar Timeline memiliki toggle untuk menampilkan/menyembunyikan marker.
  2. ⚠ Hover marker menampilkan nama version, status, dan Release date; klik marker membuka version detail page.
  3. Toolbar board dan backlog (team-managed dan company-managed) memiliki filter Release/Version multi-select yang memfilter kartu berdasarkan Fix versions.
  4. Company-managed mendukung quick filter custom dengan JQL `fixVersion` (mis. `fixVersion = earliestUnreleasedVersion()`).
- **Team-managed vs company-managed:** Butir 4 hanya company-managed.
- **Catatan & ketergantungan:** Detail Timeline di modul 18; quick filter di modul 10.

### REL-024 Cross-project releases (Plans)

- **Deskripsi:** Release di Plans yang mencakup beberapa proyek dan dipetakan ke version di tiap proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Program manager (Product Owner/PM), release manager; memerlukan izin mengelola version di setiap proyek terlibat saat menyimpan perubahan.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna membuat cross-project release di sebuah plan dan memilih proyek yang terlibat.
  2. Penyimpanan menghasilkan version bernama sama di setiap proyek terlibat; version yang sudah bernama sama di proyek tertentu dipakai ulang, bukan diduplikasi.
  3. ⚠ Start date dan Release date dapat fixed atau relatif (mis. setelah release sebelumnya atau mengikuti sprint).
  4. Status dan progress agregat lintas proyek ditampilkan di plan.
  5. Perubahan di plan bersifat draft dan baru tersimpan ke proyek setelah pengguna menjalankan "Review changes" lalu menyimpan.
  6. Aturan lengkap Plans didefinisikan di modul 18; modul ini hanya mengatur efek pada entitas version proyek.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada modul 18 (Plans, Fase 4).

### REL-025 Version report and Release burndown

- **Deskripsi:** Laporan proyek berbasis version untuk memantau kemajuan dan memprediksi tanggal rilis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Release manager, Scrum Master/Team Lead; memerlukan akses ke board/proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Version report (company-managed, board Scrum) menampilkan progres story points, jumlah work item, atau estimasi terhadap waktu sesuai estimation statistic board, beserta Predicted release date yang dihitung dari velocity harian tim serta Optimistic dan Pessimistic date dengan deviasi ±10% dari velocity tersebut (aturan identik dengan RPT-015 butir 3).
  2. Release burndown (company-managed Scrum) menampilkan pekerjaan yang diselesaikan per sprint untuk satu version dan perkiraan jumlah sprint tersisa dari velocity rata-rata (aturan identik dengan RPT-016 butir 2-3).
  3. Kedua laporan menyediakan dropdown untuk memilih version.
  4. Kedua laporan dapat dibuka dari version detail page (REL-014) dengan version terpilih otomatis.
  5. ⚠ Team-managed tidak menyediakan laporan berbasis release; tautan laporan disembunyikan di detail page team-managed.
  6. Definisi perhitungan dan grafik tidak didefinisikan ulang di sini; bila terjadi perbedaan, RPT-015 dan RPT-016 yang berlaku.
- **Team-managed vs company-managed:** Lihat butir 5.
- **Catatan & ketergantungan:** Spesifikasi laporan dimiliki RPT-015 dan RPT-016 ([16](16-laporan-insight.md)); fitur ini hanya entry point dari halaman Releases (REL-012) dan version detail page (REL-014). Membutuhkan history Fix versions per work item (CMT-023) dan data sprint (modul 11). Modul 16 perlu menambahkan rujukan balik ke REL-025 pada RPT-015 dan RPT-016.

## Otomasi & API version

### REL-026 Version automation triggers and actions

- **Deskripsi:** Trigger dan action otomasi yang terkait siklus hidup version.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, pembuat rule otomasi (modul 20).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Trigger tersedia: Version created, Version updated, Version released, Version unreleased; trigger Version archived tidak disediakan di rilis pertama (perubahan archive memicu Version updated).
  2. ⚠ Action tersedia: Create version dan Release version; action Unrelease/Archive version tidak disediakan di rilis pertama.
  3. Smart value `{{version.name}}`, `{{version.releaseDate}}`, `{{version.startDate}}`, `{{version.description}}`, `{{version.released}}`, dan `{{version.id}}` tersedia di rule yang dipicu event version.
  4. Template rule contoh disediakan: membuat version berikutnya otomatis saat sebuah version dirilis, dan mengirim release notes via email atau chat saat version dirilis.
  5. Mesin otomasi, eksekusi, dan audit log didefinisikan di modul 20.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Event berasal dari katalog modul 28.

### REL-027 Version REST API and webhooks

- **Deskripsi:** Endpoint REST dan event webhook untuk mengelola version dari luar produk.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Integrator/Developer API; izin sama dengan UI (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia daftar semua version setara `GET /rest/api/3/project/{projectIdOrKey}/versions` dan versi paginasi setara `/project/{projectIdOrKey}/version` dengan filter `status` (released, unreleased, archived), `query`, serta `orderBy` (`description`, `name`, `releaseDate`, `sequence`, `startDate`, dengan awalan `-` untuk descending).
  2. Tersedia `POST /version`, `GET /version/{id}`, dan `PUT /version/{id}` dengan `expand` `issuesstatus`, `operations`, `approvers`, `driver`.
  3. Tersedia operasi `POST /version/{id}/move`, `PUT /version/{id}/mergeto/{moveIssuesTo}`, dan `POST /version/{id}/removeAndSwap` (REL-003, REL-006, REL-007).
  4. Tersedia `GET /version/{id}/relatedIssueCounts` dan `GET /version/{id}/unresolvedIssueCount`.
  5. Tersedia endpoint related work `/version/{id}/relatedwork` (REL-017).
  6. Webhook diterbitkan untuk `jira:version_created`, `jira:version_updated`, `jira:version_released`, `jira:version_unreleased`, `jira:version_moved`, `jira:version_merged`, dan `jira:version_deleted` dengan payload objek version.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Spesifikasi endpoint kanonis ada di [API-039 Project Versions API](23-api-publik-webhook.md) (P1 / Fase 2); prioritas dan fase fitur ini disamakan dengannya agar endpoint version tersedia bersamaan dengan halaman Releases. Prefix path, autentikasi, rate limit, dan penamaan event final ditetapkan di modul 23; nama event berawalan `jira:` dipertahankan hanya sebagai alias kompatibilitas.

## Component

### REL-028 Components entity

- **Deskripsi:** Sub-bagian proyek (modul, area, tim) untuk mengelompokkan work item dan menentukan default assignee.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola), Developer/Anggota Tim (memakai); melihat memerlukan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Component memiliki atribut Name (wajib, unik per proyek, maks 255 karakter), Description (opsional), Component lead (user, opsional), dan Default assignee (assignee type, REL-031).
  2. ⚠ Keunikan nama component dibandingkan secara case-insensitive dan mengabaikan spasi di awal/akhir, sama dengan aturan nama version (REL-001).
  3. Component berlaku per proyek; tidak ada component global kecuali katalog komponen eksternal diaktifkan (REL-035).
  4. JQL mendukung `component = X`, `component is EMPTY`, dan `component in componentsLeadByUser()` (modul 12).
  5. Component dapat dipakai di board filter, quick filter, swimlane berbasis query (modul 10), dan gadget dashboard seperti statistik per component (modul 17).
  6. ⚠ Company-managed selalu menyediakan components. Team-managed juga menyediakan components dengan default assignee yang sama (PRJ-056) dan toggle "Components" di Features (PRJ-049).
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** Work item mereferensikan component lewat ID sehingga rename langsung tercermin; skema di modul 28.

### REL-029 Components field

- **Deskripsi:** System field multi-value pada work item untuk memilih component proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, pelapor; mengubah memerlukan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field berupa multi-select component picker yang hanya menawarkan component milik proyek work item.
  2. Field dapat diedit di create dialog, issue view, dan bulk edit (Add / Replace all / Remove / Clear all).
  3. Field dapat dibuat required atau hidden melalui field configuration dan screens (company-managed).
  4. Move work item ke proyek lain memetakan setiap component ke component proyek tujuan atau mengosongkannya (ISS-032).
  5. Setiap perubahan tercatat di History work item.
  6. ⚠ Membuat component baru langsung dari field tidak didukung; mengetik nama yang tidak ada menampilkan "Tidak ada component yang cocok" beserta tautan ke halaman Components bagi admin.
  7. Saat katalog komponen eksternal aktif untuk proyek (REL-035), field menampilkan component dari katalog tersebut alih-alih component proyek.
- **Team-managed vs company-managed:** Butir 3 hanya company-managed.
- **Catatan & ketergantungan:** Definisi field di CFG-011.

### REL-030 Create and edit components

- **Deskripsi:** Dialog untuk membuat dan mengedit component dari halaman Components.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek dengan Administer Projects (REL-008).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Create component" membuka dialog berisi Name, Description, Component lead, dan Default assignee.
  2. Menu "..." > Edit membuka dialog yang sama dengan nilai terisi.
  3. Name wajib dan unik per proyek; pelanggaran ditolak dengan pesan "Component dengan nama ini sudah ada".
  4. Dialog menampilkan peringatan bila default assignee yang dihasilkan (mis. Component lead) bukan assignable user di proyek (PERM-022), tanpa memblokir penyimpanan.
  5. Rename langsung tercermin di semua work item yang memakai component.
  6. API menyediakan operasi setara `POST /rest/api/3/component` dan `PUT /component/{id}`.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber izin.
- **Catatan & ketergantungan:** PRJ-056 merujuk ke fitur ini.

### REL-031 Component default assignee

- **Deskripsi:** Aturan auto-assign work item berdasarkan component saat work item dibuat tanpa assignee eksplisit.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (mengatur), pelapor (terdampak saat create).
- **Aturan perilaku (acceptance criteria):**
  1. Opsi default assignee: Project default, Component lead, Project lead, Unassigned.
  2. Opsi Unassigned hanya tersedia bila setting global "Allow unassigned issues" aktif (CFG-057).
  3. Aturan diterapkan saat work item dibuat dengan Assignee "Automatic"/kosong dan memiliki minimal satu component.
  4. ⚠ Bila work item memiliki beberapa component dengan default assignee berbeda, dipakai default dari component pertama menurut urutan nama alfabetis.
  5. Bila user target tidak assignable, sistem fallback ke project default assignee dan halaman Components menampilkan ikon peringatan pada component tersebut.
  6. Aturan tidak diterapkan ulang bila component diubah setelah work item dibuat.
  7. Objek component di API memuat `assigneeType` (`PROJECT_DEFAULT`, `COMPONENT_LEAD`, `PROJECT_LEAD`, `UNASSIGNED`), `realAssignee`, `realAssigneeType`, dan `isAssigneeTypeValid`.
- **Team-managed vs company-managed:** Tidak ada perbedaan (lihat REL-028 butir 6).
- **Catatan & ketergantungan:** Project default assignee didefinisikan di modul 04.

### REL-032 Components page

- **Deskripsi:** Halaman proyek untuk melihat dan mengelola semua component.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (mengelola), Developer/Anggota Tim (melihat dengan Browse Projects).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel menampilkan kolom Component, Description, Lead, Default assignee, dan jumlah Issues yang tertaut ke issue navigator dengan JQL `component = <id>`.
  2. Search box memfilter berdasarkan nama component.
  3. ⚠ Kolom Component, Lead, dan Issues dapat di-sort naik/turun; default urutan nama alfabetis.
  4. Menu "..." per baris menyediakan Edit dan Delete bagi pengguna berizin.
  5. Tombol "Create component" tersedia bagi pengguna berizin.
  6. Halaman diakses dari sidebar proyek "Components" (software project) atau Project settings > Components.
  7. Bila katalog komponen eksternal tersedia, halaman menyediakan toggle sumber component: component proyek vs katalog eksternal (REL-035).
  8. Bila belum ada component, halaman menampilkan empty state berisi penjelasan dan tombol Create component.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain lokasi menu.
- **Catatan & ketergantungan:** Daftar dipaginasi mengikuti konvensi modul 27.

### REL-033 Delete component with migration

- **Deskripsi:** Menghapus component dengan opsi memindahkan work item ke component lain.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek dengan Administer Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Delete menampilkan jumlah work item yang memakai component.
  2. Dialog menawarkan opsi pindahkan ke component lain (dropdown component proyek yang sama) atau hapus component dari work item.
  3. Penghapusan bersifat permanen dan tidak dapat di-undo; dialog menampilkan teks peringatan tersebut.
  4. Dialog memperingatkan bahwa filter JQL yang mereferensikan component dapat menjadi invalid.
  5. API menyediakan penghapusan setara `DELETE /rest/api/3/component/{id}?moveIssuesTo={id}`.
  6. API menyediakan `GET /component/{id}/relatedIssueCounts` untuk jumlah work item terdampak.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penggantian nilai dicatat di History dan berjalan transaksional.

### REL-034 Component REST API and webhooks

- **Deskripsi:** Endpoint REST dan event webhook untuk mengelola component dari luar produk.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Integrator/Developer API; izin sama dengan UI.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia daftar semua component setara `GET /rest/api/3/project/{projectIdOrKey}/components` dan versi paginasi setara `/project/{projectIdOrKey}/component` dengan `orderBy`, `query`, dan `componentSource` (internal/eksternal).
  2. Tersedia `POST /component`, `GET/PUT/DELETE /component/{id}`, dan `GET /component/{id}/relatedIssueCounts`.
  3. ⚠ Tersedia `GET /component` untuk mencari component lintas proyek yang dapat dilihat pemanggil, dengan paginasi dan filter `projectIdsOrKeys`.
  4. ⚠ Webhook diterbitkan untuk component created, updated, dan deleted dengan nama event `component_created`, `component_updated`, `component_deleted` (alias kompatibilitas `jira:component_*` ditetapkan di modul 23).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Spesifikasi endpoint kanonis ada di [API-038 Project Components API](23-api-publik-webhook.md) (P1 / Fase 2); prioritas dan fase fitur ini disamakan dengannya. Konvensi API di modul 23; event di modul 28.

### REL-035 Compass components in Jira

- **Deskripsi:** Memakai component dari katalog komponen eksternal (padanan Compass) sebagai pengganti component proyek; untuk produk ini diperlakukan sebagai integrasi opsional, bukan produk katalog yang dibangun sendiri.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek (mengaktifkan per proyek), platform engineer (mengelola katalog di sistem eksternal).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Components menyediakan toggle per proyek untuk beralih antara component proyek dan component katalog eksternal.
  2. Component katalog memiliki tipe (Service, Library, Application, Capability, Cloud resource, Data pipeline, Machine learning model, UI element, Website, Other), owner team, dan tautan ke katalog sumber.
  3. Field Components pada work item menampilkan component katalog yang berlaku lintas proyek di site.
  4. ⚠ Component proyek lama tetap disimpan saat beralih ke katalog dan dapat dipakai kembali bila proyek beralih balik, tanpa kehilangan nilai di work item.
  5. Katalog eksternal dapat menampilkan work item yang tertaut ke component-nya melalui API (REL-034 dengan `componentSource`).
  6. ⚠ Component lead dan default assignee tidak berlaku untuk component katalog; auto-assign REL-031 dilewati bila sumber component adalah katalog.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk katalog (Compass) di luar cakupan; integrasi lewat modul 25.

### REL-036 Archive components

- **Deskripsi:** Mengarsipkan component agar tidak dapat dipilih lagi tanpa menghapus data historis; fitur ini berasal dari Data Center dan tidak dibangun, dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. (Rujukan) Aksi Archive/Restore di halaman Components.
  2. (Rujukan) Component archived disembunyikan dari picker tetapi tetap tampil di work item lama.
  3. (Rujukan) Daftar component memiliki filter status archived.
  4. (Rujukan) Component archived tetap dapat dicari via JQL.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tidak dibangun karena ketersediaannya di produk cloud referensi belum terkonfirmasi; kebutuhan serupa dapat dipenuhi dengan Delete component with migration (REL-033). Masuk daftar Di luar cakupan.

## Pertanyaan terbuka

- Apakah definisi work item "unresolved" untuk dialog rilis, progress, dan `unresolvedIssueCount` memakai Resolution kosong atau status category Done (REL-005)?
- Apakah produk ini mengadopsi permission terpisah "Manage Versions" (PERM-026) di samping Administer Projects, dan apakah izin serupa diperlukan untuk component (REL-008)?
- Apakah Releases dapat diaktifkan di proyek bisnis, atau dibatasi ke software project (REL-001)?
- Apakah pembandingan keunikan nama version dan component bersifat case-insensitive, dan berapa batas panjang Description version (REL-001, REL-028)?
- Pada board Kanban multi-proyek, di proyek mana version baru dibuat saat Release from Kanban board, atau apakah dibuat di setiap proyek (REL-022)?
- Apakah Release approvers boleh memblokir aksi Release (mis. wajib semua Approved) sebagai opsi per proyek (REL-016)?
- Apakah team-managed mendapat versions panel backlog, tombol Release di Kanban, dan laporan berbasis release seperti company-managed (REL-021, REL-022, REL-025)?
- Bagaimana aturan default assignee bila work item memiliki beberapa component dengan default berbeda: alfabetis, urutan pemilihan, atau urutan component di proyek (REL-031)?
- Apakah fitur AI release notes masuk plan berbayar tertentu atau tersedia untuk semua instalasi yang mengonfigurasi penyedia model (REL-020)?

## Di luar cakupan & ekosistem

- Archive components (REL-036) sebagai fitur Data Center; produk ini memakai delete dengan migrasi.
- Produk katalog komponen/layanan (Compass) itu sendiri; produk ini hanya menyediakan integrasi opsional (REL-035).
- Produk wiki/dokumentasi (Confluence) beserta template release notes/retrospective-nya; produk ini hanya menautkan atau mengirim konten melalui integrasi modul 25 (REL-017, REL-019).
- Layanan AI vendor referensi (Atlassian Intelligence, Rovo); AI release notes memakai fitur AI opsional produk ini (REL-020).
- Integrasi review kode legacy (Fisheye/Crucible) untuk kategori warning "Open reviews" (REL-018).
- App Marketplace pihak ketiga untuk manajemen rilis lanjutan (release calendar, release checklist/governance, release notes generator berformat kustom, versi lintas proyek tanpa Plans, hierarki component/sub-component, bundle versi); tidak dibangun sebagai fitur native.
- Pipeline deployment, deployment gating, dan environment tracking didefinisikan di modul 24 dan 32, bukan di modul ini.
