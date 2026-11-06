# 20 — Otomasi Berbasis Aturan

Modul ini mendefinisikan mesin otomasi tanpa kode (no-code automation) produk ini: rule yang terdiri dari tepat satu trigger lalu rangkaian condition, action, branch, dan blok If/else, disusun di rule builder visual dan dijalankan secara asinkron oleh engine eksekusi atas nama rule actor. Cakupannya meliputi pengelolaan rule per project dan global (daftar rule, scope, detail, actor, siklus hidup, izin, template, ekspor/impor), semantik eksekusi beserta perlindungan loop, audit log, kuota rule run dan service limit, katalog trigger (event work item, field, jadwal, manual, webhook masuk, sprint, version, worklog, form, DevOps, build/deployment), katalog condition dan action, branching, bahasa template smart values beserta fungsi tanggal, teks, matematika, list, dan JSON, integrasi automation di issue view, work item berulang (Set to recur), connections/secrets/REST API automation, serta AI opsional untuk menyusun rule. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, middleware bahasa EN/ID) tanpa event bus maupun background worker, sehingga seluruh modul ini adalah pekerjaan baru: feature module `internal/features/automation` (handler, DTO, use case pengelolaan rule) ditambah worker eksekusi terpisah yang mengonsumsi event domain dari outbox (modul 28 dan 27), entitas `automation_rules`, `automation_rule_components`, `automation_executions`, `automation_execution_logs`, dan `automation_usage` di `internal/shared/domain/entity`, serta halaman SPA `/projects/:key/settings/automation` dan `/admin/automation` di `frontend/`. Prasyarat yang tidak bisa disusulkan: setiap perubahan work item, sprint, version, worklog, komentar, dan link wajib memancarkan event domain berisi changelog sejak Fase 1, karena trigger automation dibangun di atas katalog event tersebut.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi "work item", plan tier (Free/Standard/Premium/Enterprise), feature flag.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — grup untuk Manual trigger, user nonaktif (USR-008), Jira audit log (USR-046).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Automation Actor Permissions (PERM-049), Create/Edit/Delete (PERM-021), Assignable user (PERM-022), Transition (PERM-023), Manage Versions (PERM-026), Watchers (PERM-033), visibilitas komentar/worklog (PERM-045).
- [04 Proyek](04-proyek.md) — Project settings > Automation (PRJ-057), siklus hidup project (PRJ-031, PRJ-033), Versions (PRJ-055).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — custom field context (CFG-035), screens (CFG-046), link types (CFG-056).
- [06 Work Item Inti](06-work-item-inti.md) — Due date (ISS-021), Issue Links (ISS-029), Clone (ISS-031), Delete (ISS-033), Bulk Change (ISS-035), Watchers (ISS-036).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — changelog (CMT-023), History (CMT-022), panel Automation di issue view (CMT-026), hapus lampiran (CMT-013).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — transisi (WF-004), event workflow (WF-016), conditions/validators/post functions (WF-027, WF-028, WF-029).
- [10 Board](10-board-kolom-kartu.md) — event board & sprint untuk automation (BRD-053).
- [11 Backlog & Sprint](11-backlog-sprint.md) — Sprint Create/Start/Complete (SPR-014, SPR-018, SPR-020).
- [12 JQL](12-jql-bahasa-kueri.md) — validasi JQL, fungsi, dan batas eksekusi (JQL-037, JQL-052), entity property search (JQL-036).
- [14 Pelacakan Waktu & Estimasi](14-pelacakan-waktu-estimasi.md) — Log Work action & time smart values (TIME-043), worklog (TIME-012, TIME-018).
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — version entity dan Release/Unrelease (REL-001, REL-002, REL-004).
- [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Calendar view yang menampilkan ikon recurrence.
- [19 Notifikasi & Email](19-notifikasi-email.md) — pengiriman email, notifikasi error rule, opsi "Send notifications".
- [22 Administrasi Sistem](22-administrasi-sistem.md) — halaman admin global, pengaturan site.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST API yang dipanggil Send web request, kontrak autentikasi API token.
- [24 Integrasi DevOps](24-integrasi-devops.md) — data branch/commit/PR/build/deployment untuk trigger DevOps.
- [25 Integrasi & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — Slack/Teams, sandbox, impor CSV.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — Forms yang memicu Form submitted trigger.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — antrean job, worker, rate limiting, observabilitas.
- [28 Model Data Inti](28-model-data-inti.md) — katalog event domain `issue.*`, `sprint.*`, `version.*`, `worklog.*`.
- [35 JSM AI, Otomasi & Admin](35-jsm-ai-otomasi-api-admin-csm.md) — trigger/condition/action khusus Service Management (SLA, approvals, customer).

## Persona terkait

- **Admin Situs** — mengelola rule global/multi-project, kuota, service limit, performance insights, ekspor/impor, connections dan secrets.
- **Admin Proyek/Project Lead** — membuat dan memelihara rule single-project, memasang template, membaca audit log project.
- **Scrum Master/Team Lead** — membuat rule sprint (ringkasan sprint, pembersihan backlog), auto-assign, dan menjaga konsistensi status parent/sub-task.
- **Developer/Anggota Tim** — menjalankan Manual trigger dari issue view, mendapat transisi otomatis dari PR merged/deployment, mengatur work item berulang.
- **Product Owner/PM** — memakai template notifikasi dan release, memantau rule yang mengubah backlog.
- **Pengguna Tim Bisnis** — memakai work item berulang, rule dari Form submitted, dan template bisnis.
- **Integrator/Developer API** — memicu rule lewat Incoming webhook, memanggil sistem eksternal lewat Send web request, mengelola rule lewat Automation REST API.
- **Agent Layanan** — memakai rule Manual trigger dan komentar otomatis di service project (detail JSM di modul 35).
- **Manajer Layanan/ITSM** — menyusun rule lintas project bertipe Service management dan memantau kuota.

## User story modul

- Sebagai Admin Proyek/Project Lead, saya ingin membuat rule dari trigger, condition, dan action di builder visual agar pekerjaan berulang tim berjalan otomatis tanpa menulis kode (AUTO-003, AUTO-018, AUTO-039).
- Sebagai Admin Proyek/Project Lead, saya ingin memasang template populer seperti "tutup parent bila semua sub-task Done" lewat wizard singkat agar automation pertama aktif dalam hitungan menit (AUTO-010, AUTO-035).
- Sebagai Admin Situs, saya ingin membuat rule ber-scope global atau multi-project yang tidak bisa diubah project admin agar kebijakan organisasi konsisten di semua project (AUTO-002, AUTO-005, AUTO-008).
- Sebagai Admin Proyek/Project Lead, saya ingin audit log per eksekusi dengan log per komponen dan pesan error spesifik agar saya dapat memperbaiki rule yang gagal (AUTO-014, AUTO-057).
- Sebagai Admin Situs, saya ingin loop protection dan throttling bawaan agar rule yang salah konfigurasi tidak membanjiri sistem (AUTO-013, AUTO-016).
- Sebagai Admin Situs, saya ingin halaman Usage berisi konsumsi rule run per project dan peringatan ambang agar kuota plan tidak habis tanpa disadari (AUTO-015).
- Sebagai Scrum Master/Team Lead, saya ingin rule yang terpicu saat sprint selesai dan mengirim email ringkasan berisi daftar work item agar stakeholder mendapat laporan tanpa kerja manual (AUTO-026, AUTO-052, AUTO-061).
- Sebagai Developer/Anggota Tim, saya ingin work item berpindah ke Done otomatis saat pull request di-merge agar status board selalu sesuai kondisi kode (AUTO-030, AUTO-040).
- Sebagai Developer/Anggota Tim, saya ingin menjalankan rule manual dengan form input dari issue view agar tindakan rutin yang kompleks cukup satu klik (AUTO-024, AUTO-071).
- Sebagai Integrator/Developer API, saya ingin memicu rule lewat webhook masuk dan memanggil API eksternal dengan respons yang bisa dipakai komponen berikutnya agar produk ini terhubung dengan sistem lain (AUTO-025, AUTO-054).
- Sebagai Admin Proyek/Project Lead, saya ingin smart values dengan fungsi tanggal, teks, dan matematika agar isi komentar, email, dan field yang dihasilkan rule dinamis (AUTO-063, AUTO-066, AUTO-067, AUTO-068).
- Sebagai Pengguna Tim Bisnis, saya ingin menjadikan work item berulang langsung dari field Due date agar tugas rutin mingguan dibuat otomatis tanpa memahami rule builder (AUTO-072).
- Sebagai Scrum Master/Team Lead, saya ingin auto-assign dengan Balanced workload atau Round-robin dari sebuah grup agar beban kerja tim merata (AUTO-041).
- Sebagai Admin Situs, saya ingin mengekspor rule dari sandbox lalu mengimpornya ke production dalam status disabled agar perubahan automation dapat ditinjau sebelum aktif (AUTO-011).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 22 |
| P1 | 29 |
| P2 | 19 |
| P3 | 4 |
| **Total** | **74** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 62 |
| Fase 4 | 11 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: automation ditempatkan di Fase 3 sesuai pemetaan fase produk, tetapi prasyaratnya (event domain dengan changelog, outbox, dan background worker dari modul 27/28) harus sudah ada sejak Fase 1, sehingga di Fase 1-2 tidak ada fitur automation yang dirilis. Fase 3 dimulai dengan inti mesin: model eksekusi asinkron, loop protection, rule actor, permission, audit log, dan rule builder dengan trigger event work item (created, transitioned, updated, field changed), condition dasar (Issue fields, JQL), action dasar (Edit, Transition, Assign, Create, Comment), serta issue smart values; inilah 22 fitur P0. Gelombang kedua Fase 3 menambah fitur P1 yang paling sering dipakai: Scheduled dan Manual trigger, branch, If/else, lookup/variabel, Send email/web request, fungsi smart value tanggal dan teks, template, global administration, usage dan service limits, dan panel automation di issue view, disusul fitur P2 yang murah karena memakai engine yang sama (webhook masuk, worklog, form, delete, fungsi math/list/JSON). Fase 4 berisi fitur yang bergantung pada integrasi atau ekosistem: trigger/action DevOps dan CI/CD, pesan Slack/Teams/SMS, cross-product, ekspor/impor untuk sandbox, performance insights, entity property, connections/secrets/REST API, work item berulang, condition khusus JSM/Assets, dan AI penyusun rule yang opsional. Konfigurasi automation versi Data Center hanya dicatat sebagai rujukan.

## Fitur

## Pengelolaan rule

### AUTO-001 Project automation rule list

- **Deskripsi:** Halaman Project settings > Automation yang menampilkan semua rule yang berlaku untuk project (single-project, multi-project, project-type, dan global), tempat rule dibuat, dicari, difilter, dan dikelola. Halaman ini juga menjadi titik masuk automation dari board, backlog, list, dan issue view.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects di company-managed, role Administrator di team-managed) mengelola rule single-project; Admin Situs (Jira admin) mengelola semua rule; pengguna lain tidak melihat halaman ini (AUTO-008).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman dapat diakses lewat Project settings > Automation (PRJ-057) dan lewat tombol Automation (ikon petir) di top bar board, backlog, list, serta di issue view; tombol di board/backlog/list membuka panel Automation (butir 13).
  2. Tabel rule menampilkan kolom Name, Scope/Projects, Owner, Status (Enabled/Disabled), Last modified, dan Labels; rule global/multi-project/project-type yang berlaku ke project ini ikut ditampilkan dengan badge scope.
  3. Tombol "Create rule" membuka rule builder kosong dengan scope default single-project, dan tombol "Templates" membuka pustaka template (AUTO-010).
  4. Kotak pencarian memfilter rule berdasarkan nama (case-insensitive, substring) secara langsung saat mengetik dengan debounce 300 ms.
  5. ⚠ Filter tersedia untuk status, label, owner, scope, dan tipe trigger/komponen yang dipakai; filter dapat dikombinasikan dan tersimpan di query string URL.
  6. Panel Labels di samping menampilkan semua label beserta jumlah rule per label; label dapat dibuat, diubah nama, diberi warna dari palet tetap, dan dihapus (menghapus label tidak menghapus rule).
  7. Toggle Enable/Disable di setiap baris langsung mengubah status rule, menampilkan toast konfirmasi, dan dicatat sebagai CONFIG CHANGE (AUTO-014).
  8. Menu aksi per rule berisi Edit, Copy/Duplicate, Enable/Disable, Delete, Add/Remove label, dan Audit log; aksi yang tidak diizinkan untuk pengguna disembunyikan.
  9. Rule global, multi-project, dan project-type tampil read-only bagi Admin Proyek (ikon gembok, tooltip "Hanya admin situs yang dapat mengubah rule ini"); hanya Admin Situs yang dapat mengeditnya.
  10. Tautan "Audit log" di header membuka audit log tingkat project yang memuat eksekusi semua rule yang berlaku di project ini.
  11. Kolom Name dan Last modified dapat disortir naik/turun; default urutan Last modified terbaru di atas.
  12. Bila project belum punya rule, empty state menampilkan penjelasan singkat automation dan 3-6 rekomendasi template populer sesuai tipe project.
  13. Panel Automation di board/backlog/list menampilkan daftar rule project, template yang disarankan, dan tautan "Manage rules" ke halaman ini; pengguna non-admin hanya melihat template yang disarankan berlabel "Minta admin untuk mengaktifkan" atau panel tidak ditampilkan bila tidak ada konten relevan.
  14. ⚠ Rule dengan Scheduled trigger memiliki aksi "Run rule" di menu aksi daftar dan di builder untuk eksekusi seketika (AUTO-023).
- **Team-managed vs company-managed:** Halaman sama untuk kedua tipe; di team-managed hanya rule single-project yang dapat dibuat oleh Admin Proyek (AUTO-009).
- **Catatan & ketergantungan:** Endpoint `GET /api/v1/projects/:key/automation/rules` dengan parameter filter, sort, dan pagination. Panel Automation di issue view didefinisikan di AUTO-071 dan CMT-026.

### AUTO-002 Global automation administration

- **Deskripsi:** Halaman admin (Settings > System > Automation) untuk mengelola seluruh rule di site lintas project, termasuk audit log global, usage, performance insights, templates, ekspor/impor, dan pengaturan global automation.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Jira admin, Administer global permission); org admin melihat usage lintas produk di administrasi organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar menampilkan semua rule di site (global, multi-project, project type, single-project) dengan kolom yang sama seperti AUTO-001 ditambah kolom Project, dan dapat difilter per project.
  2. ⚠ Halaman memiliki tab Rules, Audit log, Usage, Performance insights, dan Templates; tab yang fiturnya belum dirilis disembunyikan lewat feature flag.
  3. Tab Audit log menampilkan audit log global lintas rule dengan filter status, rule, project, dan rentang tanggal (AUTO-014).
  4. Tombol "Create rule" di halaman ini memungkinkan memilih scope Global, Multiple projects, atau Project type selain Single project (AUTO-005).
  5. Menu "..." menyediakan Export (semua atau rule terpilih) ke JSON dan Import dari JSON (AUTO-011).
  6. ⚠ Pengaturan global automation memungkinkan Admin Situs membatasi kemampuan Admin Proyek: menonaktifkan pembuatan rule project, membatasi pilihan actor, dan membatasi action tertentu (mis. Send web request) ke Admin Situs saja.
  7. Menu Connections dan Secrets menuju pengelolaan kredensial integrasi yang dipakai rule (AUTO-056).
  8. ⚠ Bulk actions atas rule terpilih: Enable, Disable, Delete, Add label, Remove label, dan Change owner; hasil ditampilkan sebagai ringkasan berhasil/gagal per rule.
  9. Org admin dapat melihat usage automation lintas produk (Software, Service Management, Product Discovery) di halaman administrasi organisasi (modul 02).
  10. Halaman hanya dapat diakses Admin Situs; pengguna lain yang membuka URL mendapat 403 "Anda tidak memiliki akses ke administrasi automation".
  11. Admin Situs dapat membuka dan mengedit rule single-project milik project mana pun dari halaman ini tanpa harus menjadi Admin Proyek project tersebut.
- **Team-managed vs company-managed:** Rule single-project team-managed ikut tampil dan dapat diedit Admin Situs.
- **Catatan & ketergantungan:** Route SPA `/admin/automation`; endpoint `GET /api/v1/admin/automation/rules`. Usage dan performance di AUTO-015 dan AUTO-017.

### AUTO-003 Rule builder

- **Deskripsi:** Editor visual berbentuk alur vertikal untuk menyusun rule dari komponen: tepat satu trigger, lalu kombinasi condition, action, branch, dan blok If/else. Setiap komponen dikonfigurasi lewat panel samping dengan validasi.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (rule single-project), Admin Situs (semua scope).
- **Aturan perilaku (acceptance criteria):**
  1. Rule wajib memiliki tepat satu trigger di posisi pertama; trigger tidak dapat dihapus tanpa diganti, dan komponen lain berjalan berurutan di bawahnya.
  2. Tombol "+" (Add component) di antara dua komponen dan di akhir alur menyisipkan Condition, Action, Branch, atau If/else di posisi tersebut.
  3. Dialog pemilihan komponen memiliki pencarian dan pengelompokan kategori (Work item actions, Notifications, Integrations, DevOps, Variables & lookups, Advanced); komponen yang tidak tersedia untuk scope/tipe project disembunyikan.
  4. Panel konfigurasi memvalidasi field wajib; tombol simpan komponen nonaktif dan field yang kurang ditandai merah dengan pesan "Field ini wajib diisi" sampai konfigurasi lengkap.
  5. Komponen dapat diubah urutannya dengan drag-and-drop, termasuk dipindah masuk atau keluar branch/If/else; drop yang melanggar aturan (mis. branch ke dalam branch) ditolak dengan indikator visual.
  6. Setiap komponen memiliki aksi Duplicate dan Delete; menghapus branch atau If/else meminta konfirmasi karena seluruh komponen di dalamnya ikut terhapus.
  7. Panel Rule details (AUTO-004) dapat dibuka dari builder tanpa meninggalkan halaman.
  8. Komponen yang tidak valid (mis. field/status yang tidak ada di semua project scope, user nonaktif) diberi ikon peringatan dan tooltip penyebab; rule dengan komponen invalid tidak dapat diaktifkan.
  9. ⚠ Field teks yang mendukung smart values menampilkan picker smart value dengan autocomplete untuk objek konteks (issue, triggerIssue, fieldChange, dll.) dan tautan ke dokumentasi internal (AUTO-065).
  10. Komponen berbasis JQL dan Send web request memiliki tombol "Validate" yang menjalankan validasi/uji coba tanpa mengubah data dan menampilkan hasilnya inline.
  11. Kartu setiap komponen menampilkan ringkasan bahasa alami, mis. "When: Work item transitioned from To Do to Done", yang dilokalkan EN/ID.
  12. Tombol "Turn on rule" menyimpan dan mengaktifkan; "Save" menyimpan tanpa mengubah status (rule disabled tetap disabled); "Discard changes" membuang perubahan setelah konfirmasi.
  13. Meninggalkan builder dengan perubahan belum tersimpan memunculkan dialog konfirmasi "Perubahan belum disimpan. Tinggalkan halaman?".
  14. ⚠ Satu rule maksimal berisi 65 komponen (termasuk komponen di dalam branch); menambah komponen ke-66 ditolak dengan pesan "Batas 65 komponen per rule tercapai"; angka batas dapat dikonfigurasi admin.
  15. Branch tidak dapat di-nest di dalam branch lain; If/else boleh diletakkan di dalam branch.
  16. Builder memiliki tab Audit log yang menampilkan eksekusi rule ini saja (AUTO-014).
  17. Tidak ada action "Delay/Wait"; penundaan dimodelkan dengan Scheduled trigger + JQL, dan builder menampilkan petunjuk ini saat pengguna mencari "delay" atau "wait".
- **Team-managed vs company-managed:** Builder sama; daftar field/status menyesuaikan konfigurasi project (AUTO-009).
- **Catatan & ketergantungan:** Definisi rule disimpan sebagai pohon komponen versi-skema (JSON) yang juga dipakai ekspor/impor (AUTO-011). Endpoint `POST/PUT /api/v1/automation/rules/:id`.

### AUTO-004 Rule details

- **Deskripsi:** Pengaturan metadata dan perilaku rule: nama, deskripsi, scope, owner, actor, notifikasi error, izin dipicu oleh rule lain, dan label.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule owner, Admin Proyek (rule single-project), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Name wajib diisi (1-255 karakter) dan Description opsional (maks. 2.000 karakter); nama kosong menampilkan "Nama rule wajib diisi".
  2. Scope ditampilkan dan diubah sesuai aturan AUTO-005.
  3. Owner adalah user penanggung jawab, default pembuat rule, menerima email error, dan dapat diubah ke user aktif lain.
  4. Actor menentukan identitas pelaksana action sesuai AUTO-006.
  5. Notify on error memiliki tiga opsi: "E-mail rule owner once when rule starts failing after success" (default), "E-mail rule owner every time rule fails", dan "Don't notify".
  6. Checkbox Allow rule trigger ("Check to allow other rule actions to trigger this rule") default tidak dicentang untuk mencegah loop (AUTO-013).
  7. Rule dapat memiliki satu atau lebih label dari panel Labels project/global.
  8. Created by, created date, last modified by, dan last modified date tampil read-only.
  9. Rule ID (UUID) tampil read-only dengan tombol salin, dipakai untuk referensi, REST API, dan ekspor.
  10. Status rule (Enabled/Disabled) tampil di panel dan dapat diubah dari sana.
  11. Setiap perubahan rule details tercatat di audit log sebagai CONFIG CHANGE beserta field yang berubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Email error dikirim lewat modul 19.

### AUTO-005 Rule scope

- **Deskripsi:** Menentukan di mana rule berlaku: satu project, beberapa project tertentu, semua project bertipe tertentu, atau seluruh site (global).
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (semua scope), Admin Proyek (hanya single-project).
- **Aturan perilaku (acceptance criteria):**
  1. Single project adalah scope default saat rule dibuat dari Project settings dan dapat dikelola Admin Proyek project tersebut.
  2. Multiple projects berisi daftar project spesifik (minimal dua) dan hanya dapat dipilih Admin Situs.
  3. ⚠ Project type mencakup semua project bertipe Software, Business, Service management, atau Product Discovery, termasuk project yang dibuat setelah rule disimpan; hanya Admin Situs.
  4. Global mencakup semua project di site dan hanya dapat dipilih Admin Situs.
  5. Scope menentukan work item mana yang dapat memicu trigger dan field/status/work type mana yang muncul di pilihan builder (gabungan dari semua project scope).
  6. Mengubah scope menampilkan peringatan bila ada komponen yang merujuk field/status/work type yang tidak ada di semua project scope baru, dengan daftar komponen terdampak.
  7. Semua scope dihitung sama untuk kuota usage (AUTO-015).
  8. Perubahan scope tercatat di audit log sebagai CONFIG CHANGE.
  9. ⚠ Bila project dihapus ke trash atau diarsipkan, rule single-project-nya otomatis dinonaktifkan (bukan dihapus) dan diaktifkan kembali secara manual setelah restore; bila project dihapus permanen, rule single-project ikut dihapus dan project dikeluarkan dari daftar rule multi-project.
  10. ⚠ JQL pada Lookup issues, branch JQL, dan Scheduled trigger secara default dibatasi ke project dalam scope rule; checkbox "Include work items outside this rule's scope" (hanya Admin Situs) menghapus pembatasan itu.
- **Team-managed vs company-managed:** Admin Proyek team-managed hanya membuat rule single-project; Admin Situs dapat memasukkan project team-managed ke rule multi-project/global.
- **Catatan & ketergantungan:** Siklus hidup project di PRJ-031 dan PRJ-033.

### AUTO-006 Rule actor

- **Deskripsi:** Identitas user yang mengeksekusi action rule. Actor menentukan permission yang dipakai dan nama yang tampil di history, komentar, dan notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek; pemeriksaan izin actor di PERM-049.
- **Aturan perilaku (acceptance criteria):**
  1. Actor default adalah akun aplikasi khusus "Automation" (bot internal produk ini) yang tidak memakan lisensi pengguna dan tidak dapat login.
  2. Actor dapat diganti menjadi user tertentu; seluruh perubahan, komentar, dan notifikasi kemudian tampil atas nama user itu.
  3. ⚠ Opsi "User who triggered the event" (initiator) tersedia untuk Manual trigger dan trigger event work item; untuk Scheduled dan Incoming webhook opsi ini tidak tersedia karena tidak ada initiator.
  4. Actor wajib punya permission yang dibutuhkan setiap action (Browse projects, Edit, Transition, Assign, Add comments, Delete, dst.); bila tidak, action gagal dan audit log mencatat "Actor <nama> tidak memiliki izin <permission>".
  5. Workflow conditions/validators, issue security level, dan visibilitas komentar dievaluasi terhadap actor.
  6. ⚠ Admin Proyek non-admin situs hanya dapat memilih dirinya sendiri, akun Automation, atau initiator; memilih user lain hanya dapat dilakukan Admin Situs.
  7. Smart value `{{rule.actor}}` mengembalikan user actor.
  8. Bila actor dinonaktifkan atau dihapus dari site, eksekusi rule gagal dengan status FAILURE dan pesan "Rule actor tidak aktif", dan owner diberi tahu sesuai Notify on error.
  9. Email dari rule memakai nama pengirim default site kecuali "From name" diatur di Send email (AUTO-052).
- **Team-managed vs company-managed:** Tidak ada perbedaan; permission dievaluasi lewat role team-managed atau permission scheme company-managed.
- **Catatan & ketergantungan:** PERM-049, USR-008. Akun bot dibuat oleh migrasi seed sebagai user bertipe `app`.

### AUTO-007 Rule lifecycle and bulk management

- **Deskripsi:** Siklus hidup rule: membuat, menyimpan, mengaktifkan/menonaktifkan, menduplikasi, mengubah scope, dan menghapus, termasuk operasi massal dan penonaktifan otomatis oleh sistem.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs, Rule owner (sejauh ia juga admin pada scope tersebut).
- **Aturan perilaku (acceptance criteria):**
  1. Rule memiliki status Enabled dan Disabled; rule baru dapat disimpan dalam status Disabled.
  2. Turn on/Turn off mengubah status tanpa menghapus konfigurasi maupun riwayat audit log.
  3. Copy/Duplicate menghasilkan rule baru bernama "Copy of <nama>" dengan status Disabled dan owner = pengguna yang menduplikasi.
  4. Delete rule meminta konfirmasi dengan mengetik/menyetujui nama rule; penghapusan permanen dan tidak dapat dikembalikan.
  5. Admin Situs dapat mengubah scope atau memindahkan rule single-project ke project lain.
  6. Perubahan pada rule aktif baru berlaku setelah disimpan; eksekusi yang sedang berjalan memakai versi sebelum perubahan.
  7. ⚠ Sistem menonaktifkan rule otomatis bila rule melanggar service limit lebih dari 3 kali dalam 24 jam atau konfigurasinya rusak (mis. project/field yang dirujuk terhapus), dengan alasan tercatat di audit log.
  8. Owner menerima email saat rule dinonaktifkan otomatis, berisi alasan dan tautan ke rule.
  9. ⚠ Daftar rule mendukung multi-select dengan aksi massal Enable, Disable, Delete, Add label, Remove label, dan Change owner.
  10. Setiap perubahan siklus hidup tercatat sebagai CONFIG CHANGE di audit log.
  11. ⚠ Setiap penyimpanan membuat versi rule; panel Version history menampilkan daftar versi (siapa, kapan), dapat membandingkan dua versi, dan memulihkan versi lama sebagai perubahan baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Version history adalah keputusan produk clone (lihat Pertanyaan terbuka).

### AUTO-008 Automation permissions

- **Deskripsi:** Aturan siapa yang dapat melihat, membuat, mengedit, dan menjalankan rule pada setiap scope.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek, pengguna biasa.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs mengelola semua rule di semua scope, global automation, serta ekspor/impor.
  2. Admin Proyek company-managed (permission Administer Projects) mengelola rule single-project di project yang ia administer.
  3. Admin Proyek team-managed (role Administrator) mengelola rule single-project di project-nya.
  4. Rule multi-project/project-type/global terlihat di daftar project tetapi tidak dapat diedit oleh Admin Proyek; API mengembalikan 403 untuk upaya edit.
  5. Pengguna biasa tidak dapat membuat atau mengedit rule; ia hanya dapat menjalankan rule Manual trigger yang diizinkan.
  6. Manual trigger dapat dibatasi ke grup tertentu, dan pengguna tetap wajib punya Browse permission pada work item.
  7. Action dieksekusi dengan permission rule actor, bukan pembuat rule atau initiator, kecuali actor = initiator.
  8. ⚠ Non-admin situs tidak dapat mengganti actor ke user lain selain dirinya, akun Automation, atau initiator (AUTO-006).
  9. Incoming webhook diautentikasi dengan secret token per rule, bukan permission user.
  10. Audit log sebuah rule dapat dilihat oleh siapa pun yang dapat mengelola rule tersebut.
  11. ⚠ Pengaturan global "Izinkan admin proyek mengelola automation" (default aktif) dapat dimatikan Admin Situs sehingga hanya Admin Situs yang dapat membuat/mengedit rule.
  12. ⚠ Pengguna non-admin dengan permission Create issues dapat mengatur recurrence lewat "Set to recur" tanpa akses ke rule builder (AUTO-072).
- **Team-managed vs company-managed:** Company-managed memakai permission scheme; team-managed memakai role Administrator.
- **Catatan & ketergantungan:** PERM-049; middleware `AuthMiddleware` dan pengecekan izin di use case automation.

### AUTO-009 Team-managed project automation

- **Deskripsi:** Automation di project team-managed: rule single-project yang dikelola Admin Proyek dengan builder yang sama, tetapi dengan perbedaan field, status, dan hierarki.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek team-managed, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman diakses lewat Project settings > Automation di project team-managed.
  2. Admin Proyek hanya dapat membuat rule ber-scope single-project.
  3. Custom field team-managed bersifat lokal project; smart value merujuknya lewat nama atau ID unik project tersebut, dan nama ganda antar project tidak saling bertabrakan.
  4. Hierarki memakai field Parent (bukan Epic Link).
  5. Status dan transisi berasal dari workflow team-managed project.
  6. ⚠ Komponen yang bergantung pada fitur yang dimatikan di project (Sprints, Releases, Components) disembunyikan dari dialog komponen dan rule yang sudah ada menandainya invalid.
  7. Tidak ada notification scheme/custom event, sehingga opsi custom event di Multiple issue events tidak ditampilkan.
  8. Rule run team-managed dihitung ke usage dengan cara yang sama dengan company-managed.
- **Team-managed vs company-managed:** Fitur ini khusus team-managed; perbedaan di butir 3-7.
- **Catatan & ketergantungan:** Workflow rules team-managed (rule pada transisi) termasuk modul 09, bukan automation.

### AUTO-010 Automation templates library

- **Deskripsi:** Pustaka template rule siap pakai yang dikategorikan per use case, dipasang lewat wizard singkat lalu dapat disesuaikan penuh di rule builder.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Galeri template memiliki kategori Popular, DevOps, Agile, Work item management, Notifications, Integrations, Service management, dan Business.
  2. Pencarian template berdasarkan kata kunci pada nama dan deskripsi.
  3. Halaman detail template menampilkan deskripsi dan pratinjau trigger/condition/action.
  4. Wizard "Use template" meminta input minimal (mis. status tujuan, channel, user), lalu membuat rule dengan status Enabled setelah konfirmasi atau Disabled bila pengguna memilih "Simpan tanpa mengaktifkan".
  5. Setelah template diterapkan, rule dapat disesuaikan penuh di rule builder seperti rule biasa.
  6. Template bawaan minimal mencakup: tutup parent bila semua sub-task Done, sinkron status epic/parent, auto-assign saat create, komentar saat transisi, email ringkasan sprint, dan transisi saat PR merged.
  7. Rekomendasi template kontekstual tampil di board, backlog, issue view, dan halaman automation project.
  8. Set template yang ditampilkan berbeda per tipe project (software, business, service).
  9. ⚠ Admin Situs dapat menyimpan rule sebagai template kustom tingkat organisasi yang muncul di kategori "Organisasi Anda".
- **Team-managed vs company-managed:** Template yang memakai fitur tidak aktif di team-managed (mis. Components) disembunyikan.
- **Catatan & ketergantungan:** Template disimpan sebagai definisi rule JSON berparameter (format sama dengan AUTO-011).

### AUTO-011 Rule export and import

- **Deskripsi:** Ekspor definisi rule ke file JSON dan impor ke site yang sama atau site lain, mis. dari sandbox ke production.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Export dapat dilakukan untuk satu rule, beberapa rule terpilih, atau semua rule, menghasilkan satu file JSON.
  2. Import dilakukan dari halaman Global automation (AUTO-002) dengan mengunggah file JSON.
  3. Wizard import menampilkan isi file dan meminta pengguna memilih rule mana yang akan diimpor.
  4. Wizard memetakan ulang project, owner, dan actor yang tidak ditemukan di site tujuan; yang tidak dipetakan diganti dengan nilai default (pengimpor sebagai owner, akun Automation sebagai actor).
  5. Rule hasil impor selalu dibuat dalam status Disabled untuk ditinjau.
  6. ID custom field/status/link type yang berbeda antar site tidak dipetakan otomatis; komponen dengan rujukan tidak valid ditandai di builder.
  7. Hanya Admin Situs yang dapat melakukan export dan import; endpoint mengembalikan 403 untuk pengguna lain.
  8. JSON memuat versi skema, trigger, komponen, condition, branch, scope, label, dan rule details; secret dan token webhook tidak pernah diekspor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Promosi sandbox ke production di modul 25.

## Mesin eksekusi, audit & batas

### AUTO-012 Rule execution model

- **Deskripsi:** Semantik eksekusi rule: asinkron, berurutan per komponen, berhenti saat condition gagal, branch dieksekusi terpisah, dan data work item dimuat saat trigger.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / platform
- **Aktor & izin:** Sistem (worker automation); tidak ada interaksi pengguna langsung.
- **Aturan perilaku (acceptance criteria):**
  1. Rule dieksekusi asinkron setelah event di-commit (target p95 jeda < 10 detik), di luar transaksi pengguna; kegagalan rule tidak pernah membatalkan aksi pengguna.
  2. Komponen dijalankan berurutan; condition yang gagal menghentikan semua komponen berikutnya pada jalur yang sama, dan eksekusi berstatus NO ACTIONS PERFORMED bila belum ada action yang berjalan.
  3. Action yang error menghasilkan status SOME ERRORS bila sebagian action berhasil, atau FAILURE bila action pertama gagal dan rule terhenti.
  4. Branch dieksekusi sebagai unit terpisah dan tidak dijamin selesai sebelum komponen setelah branch; perubahan yang dibuat di dalam branch tidak terlihat oleh jalur utama.
  5. Data work item dimuat sekali saat trigger; perubahan dari action sebelumnya baru terlihat setelah action Re-fetch issue data (AUTO-058).
  6. Tidak ada jaminan urutan antara rule berbeda yang dipicu event yang sama; rule dijadwalkan independen.
  7. Smart values dievaluasi saat komponen yang memakainya dijalankan, bukan saat trigger.
  8. Rule chaining hanya terjadi bila rule target mengaktifkan Allow rule trigger (AUTO-004, AUTO-013).
  9. Saat beban tinggi, eksekusi masuk antrean dengan status QUEUED FOR PROCESSING dan diproses FIFO per rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konsumsi event domain dari outbox (modul 28) dan antrean job (modul 27); idempotensi per (rule_id, event_id) agar event yang terkirim ulang tidak dieksekusi dua kali.

### AUTO-013 Loop protection and throttling

- **Deskripsi:** Mekanisme yang mencegah rule saling memicu tanpa henti dan membatasi beban berlebih.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / platform
- **Aktor & izin:** Sistem; Admin Situs memantau dan mengonfigurasi ambang.
- **Aturan perilaku (acceptance criteria):**
  1. Secara default, perubahan yang dibuat action sebuah rule tidak memicu rule lain, kecuali rule target mengaktifkan Allow rule trigger.
  2. ⚠ Setiap event membawa kedalaman rantai (chain depth); rantai pemicuan beruntun yang melebihi 10 level dihentikan dengan status LOOP.
  3. Rule yang memicu dirinya sendiri berulang kali pada work item yang sama dalam jendela 1 menit terdeteksi dan dihentikan dengan status LOOP.
  4. Bila sebuah rule menerima event melebihi ambang per menit, eksekusi tambahan ditunda atau dibuang dengan status THROTTLED.
  5. Owner menerima email untuk status LOOP dan THROTTLED sesuai pengaturan Notify on error.
  6. Builder menampilkan petunjuk praktik terbaik saat Edit issue memodifikasi field yang sama dengan field yang dipantau trigger: "Tambahkan condition untuk mengecek nilai sebelum mengubahnya agar tidak terjadi loop".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Angka ambang dapat dikonfigurasi lewat konfigurasi server (modul 27).

### AUTO-014 Audit log

- **Deskripsi:** Riwayat setiap eksekusi rule dan perubahan konfigurasi, lengkap dengan status, work item terkait, durasi, dan log per komponen.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs, Rule owner yang dapat mengelola rule.
- **Aturan perilaku (acceptance criteria):**
  1. Audit log tersedia di tiga tingkat: per rule (tab di builder), per project (semua rule yang berlaku di project), dan global (AUTO-002).
  2. ⚠ Status eksekusi: SUCCESS, NO ACTIONS PERFORMED, SOME ERRORS, FAILURE, ABORTED, LOOP, THROTTLED, CONFIG CHANGE, IN PROGRESS, dan QUEUED FOR PROCESSING, masing-masing dengan warna badge berbeda.
  3. Tabel menampilkan kolom Date, Rule, Status, Duration, dan Associated items (work item key).
  4. Baris dapat diekspansi untuk melihat log per komponen: trigger, condition lolos/gagal, action sukses/error, dan work item yang diubah.
  5. Pesan dari Log action (AUTO-057) tampil di detail eksekusi sesuai urutan.
  6. Pesan error spesifik ditampilkan, mis. "Field X tidak dapat diisi karena tidak ada di screen", "Permission denied", "JQL tidak valid", "Transisi tidak tersedia dari status saat ini".
  7. Entri CONFIG CHANGE memuat siapa yang mengubah rule, kapan, dan ringkasan perubahan.
  8. Filter tersedia berdasarkan status, rentang tanggal, rule, dan associated item.
  9. Riwayat panjang dipaginasi (default 25 baris per halaman).
  10. Associated items adalah tautan ke work item terkait; work item yang tidak dapat diakses pengguna ditampilkan sebagai key tanpa tautan.
  11. ⚠ Retensi audit log default 90 hari, dapat dikonfigurasi admin; entri lebih tua dihapus oleh job harian.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Terpisah dari Jira audit log site (USR-046), tetapi CONFIG CHANGE rule global juga diteruskan ke audit log site.

### AUTO-015 Automation usage limits

- **Deskripsi:** Penghitungan rule run bulanan dan kuota per plan, beserta halaman Usage untuk memantau konsumsi dan peringatan saat mendekati limit.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Admin Situs, org admin, billing admin.
- **Aturan perilaku (acceptance criteria):**
  1. Plan Free: 100 rule run per bulan (nilai default, dapat dikonfigurasi admin instalasi).
  2. Plan Standard: 1.700 rule run per bulan (nilai default, dapat dikonfigurasi).
  3. Plan Premium: 1.000 rule run x jumlah user berlisensi per bulan, dipool untuk site (nilai default, dapat dikonfigurasi).
  4. Plan Enterprise: tanpa batas (fair use) yang tetap tunduk pada service limit (AUTO-016).
  5. Rule run hanya dihitung bila rule benar-benar menjalankan minimal satu action; eksekusi NO ACTIONS PERFORMED tidak dihitung.
  6. Semua scope (single, multi, project type, global) dihitung sama.
  7. Halaman Usage menampilkan total terpakai vs limit, tanggal reset, grafik harian, rule paling konsumtif, dan breakdown per project.
  8. ⚠ Email dikirim ke Admin Situs saat pemakaian mencapai 80% dan 100% limit.
  9. Saat limit tercapai, rule berhenti dieksekusi (status ABORTED dengan pesan "Kuota automation bulan ini habis") sampai reset atau upgrade plan.
  10. ⚠ Kuota direset pada tanggal 1 setiap bulan pukul 00:00 UTC.
  11. ⚠ Kuota dihitung terpisah per produk (pelacak kerja, Service Management, Product Discovery).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Untuk instalasi self-hosted clone, kuota dapat dimatikan sepenuhnya lewat konfigurasi (lihat Pertanyaan terbuka). Plan tier di modul 00.

### AUTO-016 Automation service limits

- **Deskripsi:** Batasan teknis platform untuk melindungi performa: waktu proses, antrean, jumlah work item yang diproses, dan ukuran data, termasuk trigger Service limit breached.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Sistem; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Total waktu proses per rule dibatasi 60 menit dalam jendela 12 jam; pelanggaran menghasilkan THROTTLED.
  2. ⚠ Antrean per rule maksimal 25.000 item; event berikutnya ditolak dengan status ABORTED.
  3. ⚠ Scheduled trigger dengan JQL memproses maksimal 1.000 work item per eksekusi.
  4. ⚠ Lookup issues mengembalikan maksimal 100 work item.
  5. ⚠ Branch JQL/related issues memproses maksimal 1.000 work item per branch.
  6. ⚠ Satu rule maksimal 65 komponen (AUTO-003).
  7. ⚠ Send email maksimal 100 penerima per action, Send web request timeout 30 detik dan respons maksimal 1 MB.
  8. Pelanggaran limit menghasilkan status THROTTLED atau ABORTED di audit log dengan nama limit yang dilanggar.
  9. Rule yang berulang kali melanggar limit dinonaktifkan otomatis (AUTO-007).
  10. ⚠ Trigger "Service limit breached" (hanya Admin Situs, scope global) memicu rule saat rule lain melanggar atau mendekati 80% limit, dengan smart values `{{breachedRule.name}}`, `{{breachedRule.url}}`, dan `{{breachedLimit}}`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semua angka dikonfigurasi terpusat (modul 27).

### AUTO-017 Performance insights

- **Deskripsi:** Halaman analitik performa automation: jumlah eksekusi, waktu proses, dan rule yang lambat atau mendekati service limit.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Grafik menampilkan jumlah eksekusi dan total waktu proses per hari/minggu.
  2. Daftar rule paling banyak dieksekusi dan paling lama diproses, masing-masing 10 teratas.
  3. Filter per project, rule, dan rentang waktu (7, 30, 90 hari).
  4. Rule yang mendekati service limit atau sering THROTTLED ditandai dengan badge peringatan.
  5. Setiap baris memiliki tautan ke rule dan audit log terkait.
  6. ⚠ Tersedia untuk semua plan; ketersediaan per plan diatur lewat feature flag.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dibedakan dari AUTO-015 (performa vs kuota); data diagregasi dari `automation_executions`.

## Trigger

### AUTO-018 Issue created trigger

- **Deskripsi:** Memicu rule saat work item baru dibuat di scope rule, baik lewat UI, API, email, impor, clone, maupun rule lain (bila diizinkan).
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Sistem; dipicu oleh pengguna mana pun yang membuat work item.
- **Aturan perilaku (acceptance criteria):**
  1. Berlaku untuk semua work type termasuk sub-task; penyaringan dilakukan lewat condition.
  2. ⚠ Terpicu untuk work item hasil clone, REST API, email, dan impor CSV; impor massal dapat menonaktifkan trigger lewat opsi impor "Jangan jalankan automation".
  3. Smart values `{{issue}}`/`{{triggerIssue}}` berisi work item baru dan `{{initiator}}` berisi pembuat.
  4. Builder menampilkan saran menambahkan Re-fetch issue data bila rule membaca field yang dapat diisi setelah create (mis. oleh post function).
  5. Tidak terpicu oleh work item yang dibuat rule lain, kecuali Allow rule trigger aktif.
  6. Bulk create memicu satu eksekusi per work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Event `issue.created` (modul 28).

### AUTO-019 Issue transitioned trigger

- **Deskripsi:** Memicu rule saat work item berpindah status lewat transisi workflow, dengan filter status asal dan/atau tujuan.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Sistem; dipicu pengguna yang melakukan transisi.
- **Aturan perilaku (acceptance criteria):**
  1. From status menerima satu atau lebih status, atau kosong yang berarti status apa pun.
  2. To status menerima satu atau lebih status, atau kosong yang berarti status apa pun.
  3. Smart values `{{changelog.status.fromString}}` dan `{{changelog.status.toString}}` tersedia.
  4. Terpicu oleh transisi dari UI, board drag-and-drop, API, dan rule lain (bila diizinkan).
  5. Daftar status pada pilihan mengikuti scope: status global di company-managed dan status project di team-managed.
  6. `{{issue.status.name}}` berisi status setelah transisi.
- **Team-managed vs company-managed:** Sumber daftar status berbeda (butir 5).
- **Catatan & ketergantungan:** WF-004, WF-016; event `issue.transitioned`.

### AUTO-020 Issue event triggers

- **Deskripsi:** Trigger berbasis event work item lainnya: updated, assigned, commented, comment edited, deleted, moved, linked, dan link deleted.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Work item updated terpicu untuk setiap perubahan field (bukan transisi, komentar, atau worklog) dengan smart value `{{changelog}}`.
  2. Work item assigned terpicu saat assignee berubah, termasuk menjadi Unassigned.
  3. Work item commented terpicu saat komentar ditambahkan dengan smart value `{{comment}}`; di service project tersedia filter Public/Internal comment.
  4. ⚠ Comment edited terpicu saat komentar diubah dengan smart value `{{comment}}` berisi versi baru.
  5. Work item deleted terpicu saat work item dihapus; hanya field dasar (key, summary, project, work type) yang tersedia sebagai smart value.
  6. ⚠ Work item moved terpicu saat dipindah antar project, dengan opsi filter From project dan From work type.
  7. Work item linked memiliki opsi link type dan smart values `{{destinationIssue}}`, `{{sourceIssue}}`, `{{linkType}}`.
  8. Work item link deleted memiliki opsi link type.
  9. Event dari Bulk Change memicu satu eksekusi per work item.
  10. Nama trigger di UI memakai istilah "Work item" (bukan "Issue").
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog event `issue.*`, `comment.*`, `issuelink.*` (modul 28); ISS-035.

### AUTO-021 Field value changed trigger

- **Deskripsi:** Memicu rule saat field tertentu berubah, dengan pilihan field yang dipantau, jenis perubahan, dan operasi asal perubahan.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Fields to monitor menerima satu atau lebih field system/custom, atau regex nama field.
  2. ⚠ Change type: "Any changes to the field value", "Value added", dan "Value deleted" (dua terakhir hanya untuk field multi-value).
  3. ⚠ For: operasi yang dipantau, yaitu All work item operations, Edit work item, Create work item, dan Transition work item.
  4. Smart values `{{fieldChange.field}}`, `{{fieldChange.from}}`, `{{fieldChange.fromString}}`, `{{fieldChange.to}}`, `{{fieldChange.toString}}` tersedia.
  5. Satu eksekusi dihasilkan per perubahan field yang cocok; satu edit yang mengubah tiga field yang dipantau menghasilkan tiga eksekusi.
  6. Field Status tidak dapat dipilih; builder menyarankan Issue transitioned trigger.
  7. Mendukung field Sprint, Fix versions, Labels, Story points, Assignee, Priority, Due date, dan semua custom field.
- **Team-managed vs company-managed:** Custom field team-managed hanya terlihat untuk project tersebut.
- **Catatan & ketergantungan:** Berbasis changelog CMT-023.

### AUTO-022 Multiple issue events trigger

- **Deskripsi:** Satu trigger yang mendengarkan beberapa event work item sekaligus, termasuk custom event dari workflow.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih satu atau lebih event: Work item created, updated, assigned, resolved, closed, reopened, commented, Work logged, Generic event, dan event lain dari katalog.
  2. Custom event yang dipancarkan post function "Fire event" (company-managed) dapat dipilih.
  3. ⚠ Smart value `{{eventType}}` berisi nama event yang terjadi agar condition/If-else dapat membedakan jalur.
  4. Trigger ini ditempatkan di kategori trigger "Advanced" pada dialog komponen.
- **Team-managed vs company-managed:** Custom event hanya ada di company-managed (notification scheme + post function); di team-managed opsi itu disembunyikan.
- **Catatan & ketergantungan:** WF-029 untuk post function Fire event.

### AUTO-023 Scheduled trigger

- **Deskripsi:** Menjalankan rule berdasarkan jadwal (interval tetap atau cron), opsional untuk setiap work item hasil JQL.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Sistem (scheduler); Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Fixed rate interval: setiap N menit/jam/hari/minggu dengan tanggal-waktu mulai.
  2. Cron expression format Quartz untuk jadwal kompleks; ekspresi tidak valid ditolak dengan pesan "Cron expression tidak valid".
  3. Pengaturan time zone (default zona waktu site) dan tampilan "Next run" berisi lima jadwal berikutnya.
  4. Checkbox "Run a JQL search and execute actions for each work item in the query" mengaktifkan mode per work item.
  5. Field JQL memiliki tombol Validate dan mendukung smart values.
  6. Opsi "Only include work items that have changed since the last time this rule executed" membatasi hasil ke work item yang `updated` setelah eksekusi sukses terakhir.
  7. Tanpa JQL, rule berjalan tanpa work item konteks (mis. untuk membuat work item berulang).
  8. ⚠ Smart value `{{issues}}` memberi akses ke seluruh hasil JQL dalam satu action, mis. untuk email ringkasan.
  9. ⚠ Hasil JQL maksimal 1.000 work item per eksekusi dan interval minimum 5 menit (AUTO-016).
  10. Tombol "Run rule" menjalankan rule seketika tanpa menunggu jadwal, dicatat di audit log dengan initiator.
  11. Rule disabled tidak berjalan dan jadwal yang terlewat saat disabled atau downtime tidak menumpuk (tidak ada catch-up).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Scheduler terdistribusi dengan lock agar satu jadwal hanya dieksekusi satu worker (modul 27).

### AUTO-024 Manual trigger

- **Deskripsi:** Rule yang dijalankan pengguna secara manual dari issue view, dengan opsi form input dan pembatasan grup.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Browse permission (dan anggota grup bila dibatasi); Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Rule Manual trigger muncul di panel Automation pada issue view (AUTO-071).
  2. Restrict to groups membatasi rule hanya dapat dijalankan anggota grup tertentu; pengguna lain tidak melihat rule tersebut.
  3. Prompt for user input menampilkan form saat rule dijalankan.
  4. ⚠ Tipe input: Short text, Paragraph, Number, Dropdown, Date, dan User picker.
  5. Setiap input memiliki label, nama variabel, status wajib/opsional, nilai default, dan daftar opsi (untuk Dropdown); form tidak dapat disubmit bila input wajib kosong.
  6. Nilai input diakses lewat `{{userInputs.<namaVariabel>}}`.
  7. `{{initiator}}` berisi pengguna yang menjalankan rule.
  8. ⚠ Rule manual dapat dijalankan untuk beberapa work item sekaligus dari bulk/list view; satu eksekusi dibuat per work item dan form input diisi sekali.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Endpoint `POST /api/v1/issues/:key/automation/rules/:id/run`.

### AUTO-025 Incoming webhook trigger

- **Deskripsi:** Rule dipicu oleh HTTP POST dari sistem eksternal ke URL unik; payload-nya dapat diakses lewat smart values.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sistem eksternal (pemegang secret), Admin Proyek/Admin Situs mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap rule memiliki URL webhook unik yang dibuat saat trigger disimpan.
  2. Secret token dikirim via header `X-Automation-Webhook-Token`; request tanpa token atau token salah ditolak 401, dan secret dapat di-regenerate (secret lama langsung tidak berlaku).
  3. Mode eksekusi: "No work items from the webhook", "Work items provided in the webhook HTTP POST body", dan "Work items from a JQL search".
  4. Format body untuk mode kedua adalah `{"issues":["KEY-1"]}` atau berisi ID work item; key di luar scope diabaikan dan dicatat.
  5. Smart value `{{webhookData}}` memberi akses ke payload JSON.
  6. Panel trigger menampilkan contoh perintah curl lengkap dengan URL dan header.
  7. Endpoint langsung merespons 202 Accepted (asinkron) tanpa mengembalikan hasil rule.
  8. ⚠ Webhook dibatasi 100 request per menit per rule; kelebihan dijawab 429 dan dicatat THROTTLED.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Endpoint publik `POST /api/v1/automation/webhooks/:ruleId` di luar `AuthMiddleware` JWT; webhook keluar umum di modul 23.

### AUTO-026 Sprint triggers

- **Deskripsi:** Trigger yang dipicu peristiwa sprint pada board Scrum.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Sprint created terpicu saat sprint baru dibuat (SPR-014).
  2. Sprint started terpicu saat sprint dimulai (SPR-018).
  3. Sprint completed terpicu saat sprint diselesaikan (SPR-020).
  4. Pengguna memilih satu atau lebih board yang dipantau; board di luar scope rule tidak dapat dipilih.
  5. Smart values `{{sprint.name}}`, `{{sprint.goal}}`, `{{sprint.startDate}}`, `{{sprint.endDate}}`, `{{sprint.completeDate}}`, `{{sprint.state}}`, `{{sprint.id}}` tersedia.
  6. Branch "Work items in the sprint" memproses work item yang ada di sprint terkait.
  7. Rule berjalan tanpa work item utama (konteks sprint); action yang membutuhkan work item harus berada di dalam branch.
- **Team-managed vs company-managed:** Di team-managed hanya tersedia bila fitur Sprints aktif.
- **Catatan & ketergantungan:** BRD-053, event `sprint.*`.

### AUTO-027 Version triggers

- **Deskripsi:** Trigger yang dipicu peristiwa version/release di project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Version created terpicu saat version baru dibuat.
  2. Version updated terpicu saat nama, deskripsi, atau tanggal version berubah.
  3. Version released terpicu saat version di-release.
  4. ⚠ Version unreleased terpicu saat version di-unrelease.
  5. Smart values `{{version.name}}`, `{{version.description}}`, `{{version.releaseDate}}`, `{{version.startDate}}`, `{{version.released}}`, `{{version.project}}` tersedia.
  6. Branch "Work items fixed in version" memproses work item dengan Fix version terkait.
  7. Project yang dipantau mengikuti scope rule.
- **Team-managed vs company-managed:** Di team-managed hanya tersedia bila fitur Releases aktif.
- **Catatan & ketergantungan:** REL-001, REL-004; event `version.*`.

### AUTO-028 Work logged trigger

- **Deskripsi:** Memicu rule saat worklog dicatat pada work item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Terpicu saat worklog dibuat; opsi tambahan memungkinkan terpicu saat worklog diubah atau dihapus.
  2. Smart values `{{worklog.timeSpent}}`, `{{worklog.timeSpentSeconds}}`, `{{worklog.author}}`, `{{worklog.comment}}`, `{{worklog.started}}` tersedia.
  3. Dapat dikombinasikan dengan condition pada Original/Remaining estimate.
  4. Terpicu juga untuk worklog yang dibuat via REST API (TIME-040).
- **Team-managed vs company-managed:** Tersedia hanya bila time tracking aktif di project.
- **Catatan & ketergantungan:** TIME-012, TIME-043.

### AUTO-029 Form submitted trigger

- **Deskripsi:** Memicu rule saat form yang terlampir pada work item disubmit.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sistem; dipicu pengguna yang mengirim form.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih form spesifik atau semua form.
  2. ⚠ Jawaban form diakses lewat smart value `{{forms.<formId>.<fieldKey>}}` dan `{{forms.last.<fieldKey>}}`.
  3. Trigger hanya tersedia di project yang mengaktifkan Forms.
  4. Template bawaan mengombinasikan trigger ini dengan Edit issue untuk menyalin jawaban form ke field work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Forms di modul 26 (dan modul 29 untuk Service Management).

### AUTO-030 DevOps triggers

- **Deskripsi:** Trigger dari aktivitas development tool yang terhubung (mis. GitHub, GitLab, Bitbucket, Azure DevOps), berdasarkan work item key di nama branch, commit, atau pull request.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Sistem; Developer memicu lewat aktivitas repositori.
- **Aturan perilaku (acceptance criteria):**
  1. Branch created terpicu saat branch yang memuat work item key dibuat.
  2. Commit created terpicu saat commit dengan work item key di pesan di-push.
  3. Pull request created terpicu saat PR yang merujuk work item dibuat.
  4. Pull request merged terpicu saat PR di-merge.
  5. Pull request declined terpicu saat PR ditolak/ditutup tanpa merge.
  6. Smart values `{{branch.name}}`, `{{branch.url}}`, `{{commit.message}}`, `{{commit.hash}}`, `{{pullRequest.title}}`, `{{pullRequest.url}}`, `{{pullRequest.sourceBranch}}`, `{{pullRequest.destinationBranch}}`, `{{pullRequest.state}}`, `{{repository.name}}` tersedia.
  7. Trigger hanya terpicu bila work item key terdeteksi di nama branch, pesan commit, atau judul/branch PR; satu eksekusi per work item yang terdeteksi.
  8. Trigger hanya tersedia bila integrasi dev tool terpasang dan terhubung ke project; bila tidak, dialog menampilkan tautan "Hubungkan dev tool".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Integrasi dan event dev di modul 24.

### AUTO-031 Build and deployment triggers

- **Deskripsi:** Trigger dari tool CI/CD terhubung untuk status build dan deployment yang terkait work item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Build successful terpicu saat build terkait work item sukses.
  2. Build failed terpicu saat build gagal.
  3. Build status changed terpicu untuk setiap perubahan status build.
  4. Deployment successful terpicu saat deployment sukses.
  5. Deployment failed terpicu saat deployment gagal.
  6. Deployment status changed terpicu untuk setiap perubahan status deployment.
  7. Filter environment type: Production, Staging, Testing, Development, Unmapped.
  8. ⚠ Filter nama environment atau pipeline berupa teks/regex.
  9. Smart values `{{build.name}}`, `{{build.url}}`, `{{build.state}}`, `{{deployment.name}}`, `{{deployment.environment.name}}`, `{{deployment.environment.type}}`, `{{deployment.state}}`, `{{deployment.url}}` tersedia.
  10. Trigger memerlukan data build/deployment yang dikirim lewat API Builds/Deployments produk ini (modul 24) oleh tool CI/CD seperti GitHub Actions, GitLab CI, atau Jenkins.
  11. ⚠ Trigger "Vulnerability found" terpicu saat integrasi keamanan melaporkan kerentanan terkait work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Modul 24; deployment gating Service Management di modul 32.

## Condition & percabangan logika

### AUTO-032 Issue fields condition

- **Deskripsi:** Condition sederhana yang membandingkan nilai field work item dengan nilai tertentu atau dengan field work item lain.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi; dievaluasi sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih satu field system atau custom.
  2. Operator: equals, does not equal, is one of, is not one of, contains any of, contains all of, contains none of, is empty, is not empty; daftar operator menyesuaikan tipe field.
  3. Nilai pembanding dapat berupa nilai spesifik, smart value, atau salinan dari trigger issue/parent/epic.
  4. Bila condition gagal, komponen berikutnya di jalur itu tidak dijalankan dan audit log mencatat "Condition tidak terpenuhi".
  5. Work item yang dievaluasi mengikuti branch tempat condition berada (`{{issue}}` cabang).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Evaluasi di memori tanpa query database (lebih cepat dari JQL condition).

### AUTO-033 JQL condition

- **Deskripsi:** Condition yang mengecek apakah work item saat ini cocok dengan query JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Input JQL memiliki tombol Validate yang memeriksa sintaks dan menampilkan galat.
  2. Smart values didukung di dalam JQL dan di-resolve sebelum query dijalankan.
  3. Evaluasi hanya terhadap work item saat ini (JQL dibungkus `key = <current> AND (...)`).
  4. Semua fungsi JQL yang didukung modul 12 dapat dipakai.
  5. JQL tidak valid saat eksekusi menghasilkan FAILURE dengan pesan galat parser di audit log.
  6. Panel condition menampilkan catatan performa "Gunakan Issue fields condition bila memungkinkan karena lebih cepat".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JQL-037, JQL-052.

### AUTO-034 User condition

- **Deskripsi:** Condition yang mengecek user (initiator, assignee, reporter, field user, atau smart value) terhadap kriteria identitas, grup, atau role.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Subjek: User who triggered the event, Assignee, Reporter, Creator, field user kustom, atau smart value.
  2. Check "exists" dan "does not exist".
  3. Check "is one of" dan "is not one of" terhadap user spesifik.
  4. Check "is in group" dan "is not in group".
  5. Check "is in project role" dan "is not in project role" pada project work item.
  6. ⚠ Beberapa kriteria dapat dikombinasikan dengan pilihan All (semua cocok) atau Any (salah satu cocok).
- **Team-managed vs company-managed:** Di team-managed, role yang tersedia adalah role team-managed (Administrator, Member, Viewer, role kustom).
- **Catatan & ketergantungan:** PERM-018, USR-015.

### AUTO-035 Related issues condition

- **Deskripsi:** Condition yang mengecek keberadaan atau keadaan work item terkait (linked, sub-task, child, parent).
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Jenis relasi: Linked work items (dengan pilihan link type), Sub-tasks, Children (work item di bawah Epic/parent), dan Parent.
  2. Condition "Are present" dan "Are not present".
  3. Condition "All match specified JQL", "Some match specified JQL", dan "None match specified JQL".
  4. Template "tutup parent bila semua sub-task Done" memakai condition ini dengan "All match specified JQL: statusCategory = Done".
- **Team-managed vs company-managed:** Team-managed memakai field Parent untuk relasi child.
- **Catatan & ketergantungan:** ISS-029, CFG-056.

### AUTO-036 Advanced compare condition

- **Deskripsi:** Condition yang membandingkan dua nilai bebas (smart values) dengan operator teks, numerik, atau regex.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. First value dan Second value adalah teks bebas yang mendukung smart values.
  2. ⚠ Operator: equals, does not equal, starts with, ends with, contains, does not contain, contains regular expression, does not contain regular expression, greater than, greater than or equal, less than, less than or equal.
  3. Bila kedua nilai dapat diparse sebagai angka, operator perbandingan memakai perbandingan numerik; selain itu perbandingan teks.
  4. Contoh penggunaan di bantuan: membandingkan `{{fieldChange.from}}` vs `{{fieldChange.to}}` dan mengecek `{{webhookResponse.status}}`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Regex memakai sintaks RE2 untuk mencegah backtracking berlebih.

### AUTO-037 If/else block

- **Deskripsi:** Blok percabangan yang menjalankan action berbeda berdasarkan kumpulan condition, lalu rule berlanjut setelah blok.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Blok If berisi satu atau lebih condition.
  2. ⚠ Setiap cabang memiliki pilihan logika "All conditions match" (AND) atau "Any condition matches" (OR).
  3. Pengguna dapat menambah cabang Else-if tambahan.
  4. Cabang Else dijalankan bila tidak ada cabang yang cocok.
  5. Setiap cabang dapat berisi action dan condition sendiri.
  6. Setelah satu cabang selesai, rule berlanjut ke komponen setelah blok.
  7. ⚠ If/else dapat ditempatkan di dalam branch dan di-nest maksimal 2 level (If/else di dalam If/else).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hanya cabang pertama yang cocok yang dijalankan.

### AUTO-038 Other conditions

- **Deskripsi:** Condition spesifik lainnya, sebagian khusus Service Management dan Assets.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Issue attachments condition mengecek ada/tidak ada lampiran, opsional dengan filter nama file (regex).
  2. Affected services condition (khusus service project) mengecek layanan yang terdampak; dirilis bersama modul Services di Fase 7.
  3. ⚠ Assets/object condition mengecek atribut objek Assets yang tertaut; dirilis bersama modul 33 di Fase 7.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Condition JSM rinci didefinisikan di [35 JSM AI, Otomasi & Admin](35-jsm-ai-otomasi-api-admin-csm.md); hanya butir 1 yang masuk Fase 4.

## Action work item

### AUTO-039 Edit issue action

- **Deskripsi:** Mengubah field work item saat ini, termasuk mode lanjutan JSON untuk operasi kompleks.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih satu atau lebih field untuk diubah.
  2. Nilai dapat berupa nilai spesifik, smart value, dikosongkan (clear), atau disalin dari trigger/parent/epic/destination/source issue.
  3. Field multi-value (Labels, Components, Fix versions, multi-select) memiliki mode Add, Remove, dan Set.
  4. More options > Additional fields menerima JSON dengan blok `fields` dan `update` (operasi set/add/remove/edit); JSON tidak valid ditolak saat simpan.
  5. Checkbox "Send notifications" (default aktif) menentukan apakah notifikasi perubahan dikirim (modul 19).
  6. ⚠ Field yang tidak ada di Edit screen work type terkait tetap dapat diisi oleh automation (screen tidak membatasi automation), tetapi field yang tidak dikenal menghasilkan error "Field X tidak dapat diisi: tidak dikenal".
  7. Field yang tidak ada di field context work type/project menghasilkan error "Field X tidak tersedia untuk work type ini" (CFG-035).
  8. ⚠ Action terpisah "Add value to the field" menambahkan option baru ke field select company-managed (hanya Admin Situs sebagai actor).
  9. Mendukung field Sprint, Parent, Fix versions, Labels, Components, Story points, Due date, dan custom field.
- **Team-managed vs company-managed:** Butir 8 hanya company-managed.
- **Catatan & ketergantungan:** Perubahan tercatat di History atas nama actor (CMT-022).

### AUTO-040 Transition issue action

- **Deskripsi:** Memindahkan work item ke status lain melalui transisi workflow.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Transition issues.
- **Aturan perilaku (acceptance criteria):**
  1. Destination status dapat berupa status spesifik.
  2. ⚠ Opsi destination tambahan: "Copy from trigger issue" dan "Same status (loop)" untuk menjalankan self-transition.
  3. Bila ada beberapa transisi ke status yang sama, pengguna dapat memilih transisi spesifik.
  4. Field (mis. Resolution) dan komentar dapat diisi saat transisi.
  5. Workflow conditions/validators dievaluasi terhadap actor; action gagal dengan pesan validator bila tidak terpenuhi.
  6. Additional fields JSON tersedia untuk transisi.
  7. Bila transisi tidak tersedia dari status saat ini, action gagal dengan pesan "Transisi ke <status> tidak tersedia dari <status saat ini>".
- **Team-managed vs company-managed:** Team-managed tidak memiliki validator kustom; workflow rules team-managed tetap dievaluasi.
- **Catatan & ketergantungan:** WF-004, WF-027, WF-028.

### AUTO-041 Assign issue action

- **Deskripsi:** Menetapkan assignee dengan berbagai metode, termasuk distribusi beban kerja.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Assign issues.
- **Aturan perilaku (acceptance criteria):**
  1. Assign ke user spesifik.
  2. Assign ke user dari field atau smart value (reporter, initiator, dll.).
  3. Salin assignee dari work item lain (trigger/parent).
  4. Previous assignee mengambil assignee sebelumnya dari history.
  5. Unassigned mengosongkan assignee (bila project mengizinkan unassigned).
  6. ⚠ Automatic assignee memakai default assignee project/component.
  7. Assign dari daftar user, grup, atau project role dengan metode Balanced workload, Round-robin, atau Random.
  8. Balanced workload dapat dibatasi dengan JQL untuk menentukan work item yang dihitung sebagai beban (default: unresolved di project).
  9. Target assignee wajib punya permission Assignable user di project; bila tidak, action gagal dengan pesan "User <nama> tidak dapat di-assign".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** PERM-022; status Round-robin disimpan per rule.

### AUTO-042 Comment on issue action

- **Deskripsi:** Menambahkan komentar ke work item dengan isi yang mendukung smart values.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Add comments.
- **Aturan perilaku (acceptance criteria):**
  1. Isi komentar mendukung markup rich text produk ini dan smart values.
  2. Visibility (company-managed) dapat diatur ke semua pengguna, project role, atau grup.
  3. Di service project tersedia pilihan komentar internal vs share with customer.
  4. Opsi "Prevent duplicates by only adding this comment once to a particular work item" mencegah komentar yang sama ditambahkan berulang.
  5. Mention user didukung lewat sintaks mention atau smart value akun user.
  6. Komentar dibuat atas nama actor.
- **Team-managed vs company-managed:** Butir 2 hanya company-managed.
- **Catatan & ketergantungan:** Modul 07 (rendering), PERM-045.

### AUTO-043 Create issue and create sub-tasks actions

- **Deskripsi:** Membuat work item baru di project mana pun, atau sub-task pada work item saat ini.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Create issues di project tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Create work item: pilih project (Same project atau project spesifik), work type, dan field.
  2. Parent dapat di-set untuk membuat child work item.
  3. Field dapat disalin dari trigger/current work item.
  4. Additional fields JSON tersedia.
  5. Smart values `{{createdIssue}}` (terakhir) dan `{{createdIssues}}` (semua dalam eksekusi) tersedia.
  6. Create sub-tasks membuat beberapa sub-task sekaligus, masing-masing dengan summary.
  7. Mode lanjutan Create sub-task mengizinkan field lengkap per sub-task.
  8. Branch "Most recently created work item" dapat dipakai untuk mengoperasikan work item baru.
  9. Bila field wajib tidak diisi, action gagal dengan pesan "Field wajib <nama> belum diisi".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** PERM-021.

### AUTO-044 Clone issue action

- **Deskripsi:** Menduplikasi work item saat ini ke project/work type tertentu sambil mengubah field.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Create issues di project tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih project dan work type tujuan.
  2. Summary dapat dikustomisasi (default awalan "CLONE - ").
  3. Pengguna memilih field yang disalin atau diubah.
  4. Additional fields JSON tersedia.
  5. `{{createdIssue}}` berisi work item hasil clone.
  6. ⚠ Link ke work item asli tidak dibuat otomatis; dibuat lewat action Link issues terpisah (AUTO-045).
  7. ⚠ Lampiran, komentar, dan sub-task tidak disalin secara default; checkbox opsional menyalin lampiran.
  8. Action ini dipakai di balik layar oleh Recurring work items (AUTO-072).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Clone manual pengguna di ISS-031.

### AUTO-045 Link issues and delete issue links actions

- **Deskripsi:** Membuat link antar work item atau menghapus link yang sudah ada.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Link issues.
- **Aturan perilaku (acceptance criteria):**
  1. Link issues: pilih link type dan arah (inward/outward).
  2. Target link dapat berupa trigger issue, most recently created issue, destination/source issue, work item key, atau smart value.
  3. Delete issue links: pilih link type dan target (semua atau work item tertentu).
  4. Link dapat dibuat ke work item hasil Lookup issues atau dari dalam branch.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ISS-029, CFG-056, PERM-024.

### AUTO-046 Manage watchers action

- **Deskripsi:** Menambah atau menghapus watcher work item.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Manage watchers.
- **Aturan perilaku (acceptance criteria):**
  1. Add watchers dari user spesifik, field user, atau smart value.
  2. Remove watchers menghapus user dari daftar watcher dengan sumber pilihan yang sama (user spesifik, field user, atau smart value); user yang bukan watcher dilewati tanpa error.
  3. User yang ditambahkan wajib punya Browse permission; bila tidak, user dilewati dan dicatat di audit log, dan action butuh permission Manage watchers pada actor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ISS-036, PERM-033.

### AUTO-047 Log work action

- **Deskripsi:** Mencatat worklog pada work item secara otomatis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Work on issues.
- **Aturan perilaku (acceptance criteria):**
  1. Time spent wajib, mendukung format durasi (mis. `1h 30m`) dan smart values.
  2. Date started default waktu eksekusi dan dapat diisi smart value tanggal.
  3. Work description/komentar worklog opsional.
  4. ⚠ Pengaturan remaining estimate: Adjust automatically (default), Leave unchanged, Set to, Reduce by.
  5. Action memerlukan time tracking aktif dan permission Work on issues; bila tidak, action gagal dengan pesan yang sesuai.
- **Team-managed vs company-managed:** Di team-managed memerlukan fitur time tracking aktif.
- **Catatan & ketergantungan:** TIME-043, TIME-018.

### AUTO-048 Delete actions

- **Deskripsi:** Action destruktif untuk menghapus work item, lampiran, atau komentar.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Delete issues/Delete all attachments/Delete all comments.
- **Aturan perilaku (acceptance criteria):**
  1. Delete work item menghapus work item saat ini termasuk sub-task-nya.
  2. Delete attachments dengan filter nama file regex (kosong = semua lampiran).
  3. ⚠ Delete comment menghapus komentar pemicu atau komentar yang cocok dengan filter isi.
  4. Actor wajib punya permission Delete issues/Delete all attachments/Delete all comments sesuai action.
  5. Penghapusan tidak dapat dibatalkan; builder menampilkan peringatan merah dan menyarankan condition ketat sebelum action ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ISS-033, CMT-013, CMT-003.

### AUTO-049 Version actions

- **Deskripsi:** Mengelola version project dari automation.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Manage versions.
- **Aturan perilaku (acceptance criteria):**
  1. Create version: nama, deskripsi, start date, dan release date (mendukung smart values); nama ganda di project ditolak.
  2. Release version: version spesifik atau next unreleased version, dengan release date.
  3. Unrelease version mengembalikan version ke status unreleased.
  4. ⚠ Smart value `{{version}}` berisi version terkait, termasuk version yang baru dibuat oleh Create version.
  5. Action hanya tersedia untuk project yang fitur Releases-nya aktif.
- **Team-managed vs company-managed:** Di team-managed memerlukan fitur Releases aktif.
- **Catatan & ketergantungan:** REL-002, REL-004, PERM-026.

### AUTO-050 Sprint management via automation

- **Deskripsi:** Mengelola keanggotaan sprint work item lewat automation, terutama memindahkan work item antar sprint.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan permission Edit issues dan Schedule issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Menambah/memindahkan work item ke sprint via Edit issue (field Sprint) dengan opsi sprint spesifik, "Active sprint", atau "Next sprint" pada board terpilih.
  2. Mengeluarkan work item dari sprint dengan mengosongkan field Sprint.
  3. ⚠ Tidak ada action Create/Start/Complete sprint pada rilis awal; kebutuhan itu dipenuhi lewat Send web request ke REST API sprint produk ini.
- **Team-managed vs company-managed:** Di team-managed memerlukan fitur Sprints aktif.
- **Catatan & ketergantungan:** SPR-019 (perubahan scope sprint tercatat untuk laporan).

### AUTO-051 Set entity property action

- **Deskripsi:** Menyimpan data tersembunyi (entity property) pada work item untuk dipakai automation lain atau app.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Rule actor dengan permission Edit issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Entity yang didukung: work item (project dan user tidak didukung pada rilis awal).
  2. Property key wajib dan value berupa teks atau JSON yang mendukung smart values; value JSON tidak valid ditolak.
  3. Property dibaca lewat `{{issue.properties.<key>}}`.
  4. ⚠ Property dapat dicari via JQL hanya bila diindeks oleh deklarasi indeks app/ekstensi (JQL-036).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entity property API di modul 23.

## Action notifikasi & integrasi

### AUTO-052 Send email action

- **Deskripsi:** Mengirim email dengan konten kustom ke penerima dinamis.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor; penerima tidak perlu permission untuk menerima, tetapi alamat user mengikuti pengaturan privasi.
- **Aturan perilaku (acceptance criteria):**
  1. To dapat berisi Assignee, Reporter, Watchers, Voters, user dalam field, grup, project role, alamat email spesifik, atau smart value.
  2. CC dan BCC tersedia dengan pilihan penerima yang sama.
  3. Subject dan Content mendukung smart values; subject kosong ditolak.
  4. Content type HTML atau plain text, dengan opsi "Convert line breaks to HTML line breaks".
  5. From name dan Reply-to dapat diatur.
  6. ⚠ Maksimal 100 penerima unik per action; kelebihan penerima dipotong dan dicatat sebagai peringatan di audit log.
  7. Email dikirim dari alamat pengirim site yang dikonfigurasi di modul 19.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengiriman lewat antrean email modul 19; bounce dan kegagalan SMTP dicatat di audit log.

### AUTO-053 Messaging integrations actions

- **Deskripsi:** Mengirim pesan ke tool chat atau SMS eksternal lewat webhook/API penyedia.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Rule actor; konfigurasi oleh Admin Proyek/Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Send Slack message via incoming webhook URL dengan channel/user, pesan, dan username.
  2. Send Microsoft Teams message via webhook (Workflows) dengan judul, pesan, dan opsi "Include work item summary".
  3. Send Twilio notification (SMS) dengan Account SID, auth token (disimpan sebagai secret), nomor pengirim/penerima, dan pesan.
  4. Isi pesan mendukung smart values.
  5. Error jaringan atau respons non-2xx dari webhook tercatat di audit log dengan kode status.
  6. Tool chat lain dihubungkan lewat Send web request (AUTO-054).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konektor webhook Teams mengikuti mekanisme Workflows terbaru dari Microsoft. Integrasi Slack/Teams dua arah di modul 25.

### AUTO-054 Send web request action

- **Deskripsi:** Mengirim HTTP request ke URL eksternal, dengan opsi menunggu dan memakai responsnya di komponen berikutnya.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor; Admin Situs dapat membatasi action ini (AUTO-002).
- **Aturan perilaku (acceptance criteria):**
  1. URL mendukung smart values; skema selain http/https ditolak, dan alamat jaringan privat/loopback diblokir secara default (perlindungan SSRF) kecuali di-allowlist admin.
  2. Method: GET, POST, PUT, PATCH, DELETE.
  3. Headers kustom dapat ditambah, dengan opsi menyembunyikan nilai (hidden) agar tidak tampil di UI dan audit log.
  4. Web request body: Empty, Issue data (Automation format), Issue data (format REST API produk ini), atau Custom data.
  5. Opsi "Delay execution of subsequent rule actions until we've received a response" membuat action sinkron terhadap komponen berikutnya.
  6. Smart values `{{webhookResponse.status}}`, `{{webhookResponse.body}}`, `{{webhookResponse.headers}}`, dan `{{webhookResponses}}` tersedia bila opsi butir 5 aktif.
  7. Opsi "Continue running the rule even if the request response is not successful" untuk respons non-2xx; tanpa opsi ini rule berhenti dengan FAILURE.
  8. Tombol Validate mengirim request uji dan menampilkan status, header, dan body respons.
  9. ⚠ Timeout 30 detik dan respons maksimal 1 MB; pelanggaran menghasilkan error di audit log.
  10. Action dapat memanggil REST API produk ini sendiri dengan autentikasi Basic/API token di header (modul 23).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Secret header dirujuk dari AUTO-056.

### AUTO-055 Cross-product and DevOps actions

- **Deskripsi:** Action yang membuat objek di sistem lain, seperti branch repositori dan halaman wiki/dokumentasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Rule actor dengan koneksi integrasi yang valid.
- **Aturan perilaku (acceptance criteria):**
  1. Create branch di GitHub, GitLab, dan Bitbucket: pilih repositori, nama branch (smart values, mis. `{{issue.key}}-{{issue.summary.replace(" ","-").lowerCase()}}`), dan base branch.
  2. ⚠ Create page di sistem wiki/dokumentasi terhubung: pilih site/space, parent page, template, dan judul.
  3. Action hanya tersedia bila integrasi/koneksi produk terkait sudah dikonfigurasi (AUTO-056); bila koneksi kedaluwarsa, action gagal dengan pesan "Koneksi <nama> tidak valid".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Modul 24 (koneksi repositori). Produk wiki milik vendor lain di luar cakupan sebagai produk; hanya integrasinya.

### AUTO-056 Automation connections, secrets and REST API

- **Deskripsi:** Pengelolaan koneksi integrasi, penyimpanan secret, dan API untuk mengelola rule secara programatik.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs; Integrator/Developer API memakai REST API dengan token admin.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Connections menyimpan kredensial integrasi (GitHub, GitLab, Slack, Teams, Twilio, HTTP generik) yang dipakai action, dengan status Connected/Expired.
  2. ⚠ Secrets disimpan terenkripsi, dirujuk di header/body Send web request lewat `{{secrets.<nama>}}`, tidak pernah ditampilkan kembali setelah disimpan, dan disamarkan di audit log.
  3. ⚠ Automation REST API menyediakan endpoint list, get, create, update, enable, dan disable rule di `/api/v1/automation/rules`.
  4. Hanya Admin Situs yang dapat mengelola connections dan secrets.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Enkripsi at-rest dan rotasi kunci di modul 27; kontrak REST API publik di modul 23.

## Action data, variabel & branching

### AUTO-057 Log action

- **Deskripsi:** Menulis pesan ke audit log untuk debugging rule.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Pesan teks mendukung smart values (maks. 2.000 karakter setelah evaluasi; selebihnya dipotong).
  2. Pesan tampil di detail eksekusi audit log sesuai urutan komponen.
  3. Dapat digabung dengan `{{#debug}}...{{/}}` untuk mencetak nilai mentah.
  4. ⚠ Log action tidak dihitung sebagai action untuk kuota; eksekusi yang hanya menjalankan Log action tercatat NO ACTIONS PERFORMED.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-014, AUTO-065.

### AUTO-058 Re-fetch issue data action

- **Deskripsi:** Memuat ulang data work item terbaru dari database agar komponen berikutnya melihat nilai terkini.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Action tidak memiliki konfigurasi.
  2. Dipakai setelah Issue created (bila field diisi menyusul) atau setelah Edit issue/action sebelumnya.
  3. Panel menampilkan catatan "Memperlambat eksekusi; gunakan seperlunya" dan durasi re-fetch dicatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-012 butir 5.

### AUTO-059 Create variable and lookup table actions

- **Deskripsi:** Menyimpan nilai sementara selama eksekusi rule, sebagai variabel atau tabel key-value.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Create variable: nama (huruf, angka, underscore; unik dalam rule) dan nilai (smart values).
  2. Variabel diakses lewat `{{namaVariabel}}` di komponen berikutnya.
  3. Nilai variabel selalu teks; konversi lewat fungsi seperti `asNumber`.
  4. ⚠ Variabel yang dibuat di dalam branch hanya berlaku di branch itu dan tidak terlihat di jalur utama.
  5. Create lookup table: nama tabel dan baris key-value.
  6. Nilai diakses lewat `{{namaTabel.get(key)}}`; key tidak ditemukan menghasilkan string kosong.
  7. ⚠ Lookup table maksimal 200 baris.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-065.

### AUTO-060 Lookup issues action

- **Deskripsi:** Menjalankan JQL dan menyimpan hasilnya sebagai daftar work item di smart value tanpa membuat branch.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor dengan Browse permission (hasil difilter sesuai izin actor).
- **Aturan perilaku (acceptance criteria):**
  1. JQL mendukung smart values dan tombol Validate.
  2. Hasil tersedia di `{{lookupIssues}}`, mis. `{{lookupIssues.size}}` dan `{{#lookupIssues}}{{key}}{{/}}`.
  3. ⚠ Hasil maksimal 100 work item; work item berikutnya diabaikan dan audit log mencatat peringatan.
  4. List functions (size, join, sum, average) dapat dipakai pada hasil.
  5. Lookup bersifat read-only; untuk mengubah work item hasil query gunakan branch JQL (AUTO-061).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-005 butir 10, AUTO-069.

### AUTO-061 Branch rule / related issues

- **Deskripsi:** Menjalankan komponen pada work item terkait yang berbeda dari work item utama.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Rule actor; Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Target "Current work item".
  2. Target "Sub-tasks".
  3. Target "Parent".
  4. Target "Children / stories in epic".
  5. Target "Linked work items" dengan pilihan link type.
  6. Target "JQL" dengan smart values.
  7. Target "Most recently created work item".
  8. ⚠ Target "Destination work item" dan "Source work item" untuk trigger Work item linked.
  9. Target "Work items in the sprint" (hanya dengan sprint trigger).
  10. Target "Work items fixed in version" (hanya dengan version trigger).
  11. ⚠ Branch memproses maksimal 1.000 work item (AUTO-016).
  12. Branch tidak dapat di-nest; setiap cabang dieksekusi terpisah dan tidak mengembalikan data ke jalur utama.
  13. Di dalam branch, `{{issue}}` merujuk work item cabang, sedangkan `{{triggerIssue}}` tetap work item pemicu.
- **Team-managed vs company-managed:** Target Children memakai field Parent di team-managed.
- **Catatan & ketergantungan:** AUTO-012 butir 4.

### AUTO-062 Advanced branching

- **Deskripsi:** Branch "For each" yang beriterasi atas daftar dari smart value.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengonfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berisi Variable name dan smart value yang menghasilkan list.
  2. Satu iterasi dijalankan per item list.
  3. Item diakses lewat `{{namaVariabel}}` di dalam branch.
  4. ⚠ Urutan dan paralelisme iterasi tidak dijamin; iterasi dapat berjalan paralel.
  5. Tidak dapat di-nest dengan branch lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-069.

## Smart values & fungsi

### AUTO-063 Issue smart values

- **Deskripsi:** Smart values untuk membaca field work item dan atribut user di dalam komponen rule.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template; nilai di-resolve dengan permission actor.
- **Aturan perilaku (acceptance criteria):**
  1. `{{issue.key}}`, `{{issue.id}}`, `{{issue.summary}}`, `{{issue.description}}` tersedia.
  2. `{{issue.status.name}}`, `{{issue.priority.name}}`, `{{issue.issueType.name}}`, `{{issue.resolution}}`, `{{issue.project.key}}` tersedia.
  3. `{{issue.assignee.displayName}}`, `{{issue.assignee.accountId}}`, `{{issue.reporter.emailAddress}}`, `{{issue.creator}}` tersedia.
  4. `{{issue.labels}}`, `{{issue.components.name}}`, `{{issue.fixVersions}}`, `{{issue.versions}}` tersedia.
  5. `{{issue.sprint.name}}`, `{{issue.parent.key}}`, `{{issue.parent.summary}}`, `{{issue.subtasks}}` tersedia.
  6. `{{issue.issuelinks}}`, `{{issue.comments.last.body}}`, `{{issue.attachment}}`, `{{issue.watchers}}` tersedia.
  7. `{{issue.created}}`, `{{issue.updated}}`, `{{issue.duedate}}`, `{{issue.resolutiondate}}` tersedia sebagai objek tanggal.
  8. `{{issue.timetracking.originalEstimate}}`, `{{issue.timetracking.remainingEstimate}}`, `{{issue.Story Points}}` tersedia.
  9. `{{issue.url}}` berisi URL work item; di service project `{{issue.url.customer}}` berisi URL portal pelanggan.
  10. Custom field diakses lewat nama (`{{issue.Nama Field}}`) atau ID (`{{issue.customfield_10010}}`).
  11. Akses alternatif lewat `{{issue.fields.<field>}}` dan property lewat `{{issue.properties.<key>}}`.
  12. Atribut user: displayName, accountId, emailAddress (kosong bila disembunyikan oleh pengaturan privasi user), dan timeZone.
- **Team-managed vs company-managed:** Custom field team-managed hanya dapat dirujuk dalam rule project tersebut.
- **Catatan & ketergantungan:** Format ID custom field mengikuti modul 05; field yang dilindungi issue security tidak di-resolve bila actor tidak berhak.

### AUTO-064 Trigger and context smart values

- **Deskripsi:** Smart values dari konteks eksekusi rule dan trigger.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. `{{triggerIssue}}`, `{{initiator}}`, `{{rule.name}}`, `{{rule.actor}}`, `{{project.key}}`, `{{project.name}}`, `{{baseUrl}}` tersedia di semua rule.
  2. `{{changelog.<field>.fromString/toString}}` dan `{{fieldChange}}` tersedia untuk trigger berbasis perubahan.
  3. `{{webhookData}}` (Incoming webhook) dan `{{webhookResponse}}` (Send web request) tersedia.
  4. `{{userInputs}}`, `{{createdIssue}}`, `{{createdIssues}}`, `{{lookupIssues}}` tersedia setelah komponen terkait.
  5. `{{sprint}}`, `{{version}}`, `{{worklog}}`, `{{comment}}`, `{{attachment}}` tersedia sesuai trigger.
  6. `{{branch}}`, `{{commit}}`, `{{pullRequest}}`, `{{repository}}`, `{{build}}`, `{{deployment}}` tersedia untuk trigger DevOps.
  7. `{{now}}`, `{{issues}}`, `{{destinationIssue}}`, `{{sourceIssue}}`, `{{linkType}}` tersedia sesuai konteks.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Picker smart value di builder (AUTO-003) hanya menampilkan objek yang relevan untuk trigger rule.

### AUTO-065 Smart value syntax and debugging

- **Deskripsi:** Aturan sintaks smart values: notasi titik, iterasi, nilai default, case-sensitivity, encoding, dan debugging.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. Sintaks `{{objek.properti}}` bergaya Mustache dan dapat di-chain dengan fungsi, mis. `{{issue.summary.upperCase()}}`.
  2. Iterasi list memakai `{{#list}}...{{/}}` dengan akses properti item di dalamnya.
  3. Nilai default lewat pipe, mis. `{{issue.description|Tidak ada deskripsi}}`.
  4. Smart value yang tidak ter-resolve menghasilkan string kosong tanpa error.
  5. `{{#debug}}...{{/}}` mencetak nilai mentah ke audit log.
  6. Custom field dapat dirujuk via nama (case-insensitive) atau ID customfield.
  7. Encoding output: `.htmlEncode`, `.jsonEncode`, `.urlEncode`, `.xmlEncode`, `.text`, `.html`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine template diimplementasikan sebagai interpreter sandbox tanpa akses eksekusi kode, dengan batas waktu evaluasi per komponen (modul 27).

### AUTO-066 Date and time functions

- **Deskripsi:** Fungsi untuk memanipulasi dan memformat tanggal di smart values.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. Fungsi aritmetika: plusDays/minusDays, plusHours, plusMinutes, plusWeeks, plusMonths, plusYears, plusBusinessDays/minusBusinessDays.
  2. Fungsi penyesuaian: withDayOfMonth, withNextDayOfWeek, startOfMonth, endOfMonth, firstBusinessDayOfMonth, lastBusinessDayOfMonth, firstOfTheMonth, lastOfTheMonth, toBusinessDay, toBusinessDayBackwards, isBusinessDay.
  3. `diff(date)` dengan `.days`, `.businessDays`, `.hours`, `.weeks`, `.prettyPrint`.
  4. Perbandingan: isAfter, isBefore, isEquals, compareTo.
  5. Format: jiraDate, jiraDateTime, jqlDate, jqlDateTime, shortDate, mediumDate, longDate, fullDate, shortTime, longDateTime, dan `format("pattern")`.
  6. Zona waktu dan locale: convertToTimeZone, setTimeZone, locale (mis. `locale("id")`).
  7. Parsing teks menjadi tanggal: toDate dan toDateTime; teks tidak valid menghasilkan nilai kosong.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hari kerja memakai Senin-Jumat default; nama format (jiraDate) dipertahankan untuk kompatibilitas migrasi rule.

### AUTO-067 Text functions

- **Deskripsi:** Fungsi untuk memanipulasi teks di smart values.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. abbreviate, capitalize, concat, charAt.
  2. contains, startsWith, endsWith, equals, equalsIgnoreCase, indexOf, lastIndexOf.
  3. isAlpha, isAlphanumeric, isNumeric, isEmpty, isNotEmpty.
  4. left, right, leftPad, rightPad, length, lowerCase, upperCase.
  5. match (regex), replace, replaceAll (regex), remove, reverse.
  6. split, substring, substringAfter, substringAfterLast, substringBefore, substringBeforeLast, substringBetween, trim.
  7. htmlEncode, jsonEncode, urlEncode, xmlEncode, quote, asNumber.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Regex memakai RE2.

### AUTO-068 Math expressions

- **Deskripsi:** Ekspresi matematika dan fungsi angka di smart values.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. Blok `{{#=}}...{{/}}` mengevaluasi ekspresi dengan operator `+`, `-`, `*`, `/`, `%`, `^`; pembagian dengan nol menghasilkan nilai kosong dan peringatan di audit log.
  2. ⚠ Fungsi dalam ekspresi: round, floor, ceil, abs, max, min, sqrt.
  3. Fungsi angka: `.plus()`, `.minus()`, `.multiply()`, `.divide()`, `.round()`, `.abs()`, `.ceil()`, `.floor()`.
  4. Perbandingan angka: `.gt()`, `.gte()`, `.lt()`, `.lte()`, `.eq()`.
  5. ⚠ Format: `.format("pattern")`, `.asCurrency()`, `.asPercentage()`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aritmetika desimal presisi tetap untuk menghindari galat floating point pada mata uang.

### AUTO-069 List functions

- **Deskripsi:** Fungsi untuk mengolah daftar di smart values.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. first, last, get(index), getFromEnd(index), size.
  2. ⚠ join(separator), distinct, flatten.
  3. sum, average, max, min untuk list angka.
  4. ⚠ Iterasi `{{#list}}...{{/}}` menyediakan `{{index}}` (mulai 0) dan penanda `{{^last}}...{{/}}` untuk item selain terakhir.
  5. Akses properti tiap item, mis. `{{issue.subtasks.key}}`, menghasilkan daftar key.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AUTO-060, AUTO-062.

### AUTO-070 Conditional logic and JSON functions

- **Deskripsi:** Logika kondisional dan fungsi JSON di smart values.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek menulis template.
- **Aturan perilaku (acceptance criteria):**
  1. Blok `{{#if(kondisi)}}...{{/}}` dan inline `if(kondisi, nilaiYa, nilaiTidak)`.
  2. Fungsi logika: and, or, not, exists, equals, eq, gt, gte, lt, lte.
  3. Nilai default via pipe `|`.
  4. asJsonString, asJsonStringArray, asJsonArray, asJsonObject, asJsonObjectArray untuk menyusun body JSON.
  5. ⚠ jsonStringToObject mem-parse teks JSON menjadi objek yang propertinya dapat diakses; teks tidak valid menghasilkan nilai kosong.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai bersama Send web request (AUTO-054).

## Automation di work item & recurrence

### AUTO-071 Automation in issue view

- **Deskripsi:** Integrasi automation di issue view: menjalankan rule manual dan melihat riwayat eksekusi rule untuk work item tersebut.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Browse permission (menjalankan rule manual yang diizinkan), Admin Proyek (mengelola).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Automation (ikon petir) di header issue view membuka panel Automation.
  2. Panel menampilkan daftar rule Manual trigger yang tersedia untuk work item sesuai scope dan grup yang diizinkan (AUTO-024); bila tidak ada, tampil "Belum ada rule manual untuk work item ini".
  3. Klik rule membuka form user input bila dikonfigurasi, lalu rule dieksekusi.
  4. ⚠ Bagian "Rule executions" menampilkan 10 eksekusi terbaru yang terkait work item ini beserta status dan tautan ke audit log (bagi pengguna yang berhak).
  5. Tautan "Manage automation" dan "Create automation" hanya tampil bagi Admin Proyek/Admin Situs.
  6. Setelah rule manual dijalankan, flag "Rule sedang berjalan" tampil dan berganti "Rule selesai" atau "Rule gagal" saat eksekusi selesai (via realtime atau polling).
  7. Perubahan oleh rule tampil di tab History dan Activity work item atas nama actor.
  8. Menu Automation memuat entry point "Set to recur" untuk membuat work item berulang (AUTO-072).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Panel di issue view didefinisikan secara visual di CMT-026.

### AUTO-072 Recurring work items (Set to recur)

- **Alias inventaris:** Recurring Work Items
- **Deskripsi:** Opsi native untuk menjadikan work item berulang langsung dari issue view (field Due date atau menu Automation) tanpa menulis rule sendiri. Di belakangnya sistem membuat automation tersembunyi yang meng-clone work item sesuai jadwal atau saat work item selesai.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Anggota project dengan permission Create issues (dan Edit issues pada work item sumber); Admin Proyek; Sistem menjalankan clone.
- **Aturan perilaku (acceptance criteria):**
  1. Entry point "Set to recur" tersedia di menu field Due date dan di menu Automation (ikon petir) pada issue view.
  2. ⚠ Aksi "Make recurring" juga tersedia di menu ••• work item.
  3. Form Recurring berisi frekuensi/jadwal: harian, mingguan, bulanan, dan kustom.
  4. ⚠ Pola recurrence: harian, mingguan (pilih satu atau lebih hari), bulanan (tanggal atau hari ke-N), tahunan, dan interval kustom (setiap N hari/minggu/bulan).
  5. ⚠ Basis recurrence: "Berdasarkan jadwal" atau "Saat work item selesai" (completion-based, dihitung sejak work item masuk status category Done).
  6. ⚠ Pada basis jadwal, item berikutnya dibuat oleh scheduler pada tanggal kemunculan; pada basis penyelesaian, item berikutnya dibuat saat item sebelumnya selesai.
  7. ⚠ Recurrence dapat diakhiri pada tanggal akhir atau setelah sejumlah kemunculan.
  8. ⚠ Form memiliki tanggal mulai dan kondisi berhenti: Tidak pernah, Setelah N kali, atau Pada tanggal tertentu; tanggal akhir sebelum tanggal mulai ditolak.
  9. ⚠ Item baru menyalin summary, description, assignee, labels, work type, dan priority dari work item sumber, dengan due date disesuaikan otomatis.
  10. Due date salinan dihitung relatif dari tanggal pembuatan (mis. jadwal mingguan: due date = created + 7 hari).
  11. Salinan mewarisi pengaturan recurrence, yang otomatis aktif pada work item baru, dan recurrence di work item lama dinonaktifkan agar hanya satu rantai aktif.
  12. Ikon recur tampil di samping field Due date dan pada Calendar view (modul 18).
  13. Indikator recurrence tampil di work item view (tooltip berisi ringkasan pola, mis. "Berulang setiap Senin") dan di calendar view.
  14. Form Recurring menyediakan Edit recurrence dan Delete recurrence (menghentikan pengulangan tanpa menghapus work item yang sudah dibuat).
  15. Pola dapat diubah atau recurrence dihentikan langsung dari work item; perubahan berlaku untuk kemunculan berikutnya.
  16. Setiap pembuatan salinan dihitung sebagai satu rule run pada kuota automation (AUTO-015).
  17. ⚠ Konsumsi kuota tercatat per project di halaman Usage; rilis awal tidak menyediakan halaman terpusat daftar semua recurrence di site, tetapi Admin Situs dapat mencarinya lewat JQL `recurrence is not EMPTY`.
  18. Validasi: user/actor wajib dapat meng-clone work item (permission Create issues di project); bila tidak, opsi "Set to recur" disembunyikan.
  19. ⚠ Bila work item sumber dihapus, recurrence-nya ikut dihapus; bila dipindah ke project lain, recurrence dinonaktifkan dan pengguna diberi peringatan di dialog Move.
  20. ⚠ Recurrence tidak tampil sebagai rule biasa di daftar rule project, tetapi setiap eksekusinya tercatat di audit log project dengan nama "Recurrence: <key>".
  21. Alternatif lama tetap didukung: rule manual dengan Scheduled trigger + Clone issue action (AUTO-023, AUTO-044).
  22. ⚠ Tersedia di project software dan bisnis, team-managed dan company-managed; di service project disembunyikan pada rilis awal.
- **Team-managed vs company-managed:** Tersedia di keduanya (butir 22).
- **Catatan & ketergantungan:** Bergantung pada Due date (ISS-021), Clone (AUTO-044, ISS-031), dan Calendar view (modul 18). Fungsi JQL `recurrence` adalah usulan yang perlu didaftarkan di modul 12. App recurring tasks pihak ketiga hanya ekosistem.

## AI & varian deployment

### AUTO-073 Atlassian Intelligence and Rovo in automation

- **Deskripsi:** Fitur AI opsional untuk menyusun rule dari deskripsi bahasa alami, serta action yang memanggil agent AI di tengah eksekusi rule. Di produk ini fitur tersebut memakai penyedia AI yang dikonfigurasi admin, bukan layanan AI vendor tertentu.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek, Admin Situs; hanya bila AI diaktifkan admin organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. "Create rule with AI": pengguna menulis deskripsi bahasa alami (EN/ID), lalu AI menyusun draft trigger/condition/action di builder.
  2. Draft AI selalu dibuka di builder dalam status belum tersimpan dan wajib ditinjau serta disimpan manual oleh pengguna.
  3. Bila permintaan tidak dapat dipetakan ke komponen yang ada, AI menampilkan pesan "Permintaan belum dapat diubah menjadi rule; coba jelaskan trigger dan action-nya" tanpa membuat draft parsial yang invalid.
  4. ⚠ Action "Use AI agent" menjalankan agent dengan prompt berisi smart values, dan outputnya tersedia di komponen berikutnya lewat `{{agentResponse}}`.
  5. Fitur hanya tampil bila AI diaktifkan admin organisasi dan plan mendukung (feature flag).
  6. ⚠ Pemakaian AI dihitung dengan kuota kredit AI terpisah dari rule run.
  7. ⚠ Workspace terpadu untuk membangun agent, rule, dan app low-code (setara Atlassian Studio) tidak dibangun; rule yang dibuat lewat AI tetap dijalankan engine automation yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kebijakan data dan penyedia AI di modul 27 dan modul AI terkait (35); lihat Pertanyaan terbuka.

### AUTO-074 Automation for Jira Data Center configuration

- **Deskripsi:** Pengaturan khusus automation versi Data Center (konfigurasi global service limits oleh admin, tanpa kuota rule run). Dicatat sebagai rujukan pembeda; tidak dibangun sebagai fitur terpisah.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. Pengaturan service limit oleh admin (antrean, waktu proses, batas work item) dipenuhi lewat konfigurasi server AUTO-016, bukan halaman UI terpisah.
  2. Pengaturan izin project admin dan pembatasan rule actor dipenuhi oleh AUTO-002 butir 6 dan AUTO-008 butir 11.
  3. Mode tanpa kuota rule run bulanan dipenuhi oleh opsi mematikan kuota di AUTO-015 untuk instalasi self-hosted.
  4. ⚠ Pengaturan thread/eksekusi worker tidak diekspos di UI; dikonfigurasi lewat variabel lingkungan (modul 27).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## Pertanyaan terbuka

- Apakah kuota rule run (AUTO-015) diberlakukan di produk ini, atau hanya pada edisi hosted dengan opsi dimatikan sepenuhnya untuk instalasi self-hosted?
- Apakah format definisi rule JSON dan sintaks smart values dibuat kompatibel penuh dengan ekspor rule Jira Cloud agar pelanggan dapat memigrasikan rule, atau cukup subset yang didokumentasikan?
- Apakah version history rule dengan compare/restore (AUTO-007 butir 11) dibangun pada Fase 3 atau ditunda, dengan riwayat sementara hanya dari audit log CONFIG CHANGE?
- Apakah produk ini menambahkan action "Delay/Wait" native sebagai pembeda, padahal referensi tidak memilikinya (AUTO-003 butir 17)?
- Apakah action native Create/Start/Complete sprint (AUTO-050 butir 3) perlu disediakan, mengingat alternatifnya Send web request ke REST API sendiri?
- Berapa retensi audit log default (usulan 90 hari) dan apakah berbeda per plan?
- Apakah Send web request boleh menjangkau alamat jaringan privat secara default pada instalasi self-hosted, atau tetap diblokir dengan allowlist admin?
- Penyedia AI mana yang didukung untuk AUTO-073 (API eksternal vs model self-hosted) dan bagaimana kebijakan pengiriman data work item ke penyedia tersebut?
- Apakah fungsi JQL `recurrence` untuk mencari work item berulang (AUTO-072 butir 17) disetujui dan ditambahkan ke modul 12?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga:** ekstensi automation/scripting (mis. post function dan scripting Groovy, workflow extension, recurring tasks app seperti Recurring Tasks/Cadence) tidak dibangun sebagai fitur native; kebutuhan serupa dipenuhi lewat Send web request, Incoming webhook, dan titik ekstensi modul 25.
- **Fitur Data Center:** Automation for Jira Data Center configuration (AUTO-074) hanya rujukan; pengaturan setaranya ada di konfigurasi server.
- **Layanan lintas produk Atlassian:** Atlassian Studio, platform pembangun agent Rovo, automation untuk Confluence/Bitbucket, dan halaman usage di administrasi organisasi Atlassian tidak dibangun; produk ini hanya menyediakan engine automation, AI opsional, dan halaman usage internalnya sendiri.
- **Layanan pihak ketiga:** Slack, Microsoft Teams, Twilio, GitHub, GitLab, Bitbucket, dan sistem wiki eksternal adalah layanan milik vendor lain; yang dibangun hanya konektor action/trigger (AUTO-030, AUTO-031, AUTO-053, AUTO-055).
- **Automation khusus Service Management:** trigger/condition/action SLA, approval, customer, Assets, dan Operations didefinisikan di modul 35, 33, dan 34, bukan di modul ini.
- **Eksekusi kode arbitrer:** tidak ada action "Run script" yang mengeksekusi kode pengguna di server produk ini.
