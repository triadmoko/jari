# 19 — Notifikasi In-app & Email

Modul ini mendefinisikan seluruh lapisan notifikasi dan email di produk ini: katalog event sistem yang dipicu perubahan work item dan event kustom yang ditembakkan transition workflow, notification scheme company-managed beserta asosiasinya ke project dan jenis recipient, konfigurasi notifikasi per project team-managed, pipeline pengiriman terpadu (resolver penerima, filter permission dan issue security, preferensi user, deduplikasi, batching, antrean asinkron dengan retry), mekanisme watch, autowatch, dan mention, pusat notifikasi in-app (ikon lonceng) beserta preferensinya, isi dan format email notifikasi work item, pengaturan notifikasi personal dan default, alat bantu diagnosis Notification Helper, infrastruktur email keluar (SMTP, from address, sender per project, custom sending domain), email transaksional akun, email dari aksi pengguna (bulk change, Share, filter subscription), email masuk (incoming mail server, mail handler, aturan parsing, alamat email project bisnis), REST API notifikasi, interaksi dengan Automation, kanal personal eksternal (push mobile, Slack, Microsoft Teams), lokalisasi, batas volume, dan audit. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, `middleware.LanguageMiddleware()` untuk EN/ID) tanpa kemampuan mengirim email maupun menyimpan notifikasi, sehingga seluruh modul ini adalah pekerjaan baru yang dikemas sebagai feature module baru mengikuti pola Clean Architecture repo, misalnya `internal/features/notification` (event consumer, scheme, preferensi, hub in-app) dan `internal/features/mail` (outbox, SMTP sender, incoming mail, mail handler), dengan entity dan repository interface di `internal/shared/domain/`, konfigurasi `SMTP_*`/`MAIL_*` melalui `internal/core/config/`, serta halaman SPA `frontend/` untuk panel lonceng, `/settings/personal/notifications`, dan halaman admin. Notifikasi Service Management untuk customer (customer notifications, email channel JSM) dibangun di atas pipeline modul ini tetapi didefinisikan di modul 30.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi "work item", plan tier, feature flag.
- [01 Identitas & Autentikasi](01-identitas-autentikasi-sesi.md) — verifikasi email, reset password, email utama akun.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — undangan user, permintaan akses produk, status deactivated, group, Teams.
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects (PERM-020), Voters & Watchers Permissions (PERM-033), Issue Security Enforcement (PERM-044), Manage Group Filter Subscriptions (PERM-010), Make Bulk Changes (PERM-011), visibility comment/worklog.
- [04 Proyek](04-proyek.md) — project settings, share settings with an existing project, project lead, business project.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field configuration (field tersembunyi), user/group picker custom field, component lead.
- [06 Work Item Inti](06-work-item-inti.md) — Bulk Change (ISS-035), Watchers (ISS-036), Votes (ISS-037), Share Work Item (ISS-038), Move, Delete.
- [07 Editor Teks Kaya](07-editor-konten-kaya.md) — Mentions (EDT-014), render ADF ke HTML email.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — event comment, attachment, worklog, changelog.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — Workflow-related events (WF-016), Transition Post Functions (WF-029), Bulk transition (WF-013).
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — Filter subscriptions (SRC-047), Share search (SRC-048), bulk change dari hasil pencarian (SRC-031).
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — Share dashboard.
- [20 Otomasi](20-otomasi.md) — aksi Send email, Send Slack/Microsoft Teams message, opsi Send notifications.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — top navigation (ikon lonceng), personal settings, bahasa dan timezone profil, keyboard shortcut.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — Settings > System (General configuration, Events, Incoming mail), audit log.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — konvensi REST, paginasi, rate limit, parameter `notifyUsers`.
- [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — importer CSV, integrasi Slack/Microsoft Teams, aplikasi mobile.
- [26 Kerja Bisnis](26-kerja-bisnis-formulir-persetujuan.md) — business project yang memakai alamat email project.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — background job, antrean, realtime (WebSocket/SSE), rate limit, observabilitas.
- [28 Model Data Inti](28-model-data-inti.md) — katalog event domain `issue.*`, `comment.*`, `worklog.*` yang dikonsumsi modul ini.
- [30 JSM Help Center & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — customer notifications dan email channel Service Management yang memakai ulang infrastruktur email modul ini.

## Persona terkait

- **Developer/Anggota Tim** — menerima notifikasi in-app dan email untuk work item yang ia kerjakan, watch, atau tempat ia di-mention; mengatur preferensi personal dan autowatch.
- **Scrum Master/Team Lead** — memantau perubahan pada work item tim lewat watch dan filter subscription, menghindari banjir email saat bulk change.
- **Product Owner/PM** — menerima ringkasan berkala lewat filter subscription, membagikan work item/filter/dashboard ke stakeholder via Share.
- **Admin Proyek/Project Lead** — mengatur notifikasi team-managed project, melihat scheme aktif di company-managed project, mengatur alamat email project.
- **Admin Situs** — mengelola event, notification scheme, outgoing/incoming mail, mail handler, default preferensi user, Notification Helper, dan audit pengiriman.
- **Stakeholder/Viewer** — menerima email Share dan filter subscription tanpa harus aktif di produk setiap hari.
- **Pengguna Tim Bisnis** — membuat work item di business project dengan mengirim email ke alamat project.
- **Integrator/Developer API** — mengirim notifikasi kustom lewat endpoint `/notify`, menekan notifikasi saat update massal lewat API, mengelola notification scheme lewat REST.
- **Tamu/Kolaborator Eksternal** — menerima email Share atau undangan dan membalas email untuk menambah comment bila mail handler dikonfigurasi.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin menerima notifikasi in-app dan email saat work item yang saya kerjakan atau watch berubah agar saya tidak perlu memeriksa board terus-menerus (NTF-001, NTF-012, NTF-015, NTF-018).
- Sebagai Developer/Anggota Tim, saya ingin langsung diberi tahu ketika di-@mention agar pertanyaan yang ditujukan kepada saya cepat terjawab (NTF-013).
- Sebagai Developer/Anggota Tim, saya ingin perubahan beruntun pada satu work item digabung menjadi satu email agar inbox saya tidak banjir (NTF-019).
- Sebagai Developer/Anggota Tim, saya ingin mematikan kategori email tertentu dan memilih format HTML atau Text agar notifikasi sesuai kebiasaan kerja saya (NTF-020, NTF-016).
- Sebagai Admin Situs, saya ingin memetakan setiap event ke recipient dalam notification scheme yang dipakai bersama banyak project agar kebijakan notifikasi konsisten di organisasi (NTF-004, NTF-005, NTF-006).
- Sebagai Admin Proyek/Project Lead, saya ingin mengatur siapa yang dinotifikasi untuk tiap event di team-managed project saya tanpa bantuan admin situs agar tim saya mandiri (NTF-007).
- Sebagai Admin Situs, saya ingin notifikasi tidak pernah terkirim ke user yang tidak berhak melihat work item atau comment terbatas agar data rahasia tidak bocor lewat email (NTF-008).
- Sebagai Admin Situs, saya ingin mengetahui kenapa user tertentu tidak menerima notifikasi untuk suatu work item agar keluhan "email tidak datang" cepat terselesaikan (NTF-022, NTF-011).
- Sebagai Admin Situs, saya ingin mengonfigurasi SMTP server sendiri, from address, dan menguji kiriman agar instalasi self-hosted dapat mengirim email dengan andal (NTF-023, NTF-024).
- Sebagai Scrum Master/Team Lead, saya ingin mematikan notifikasi saat bulk edit ratusan work item agar tim tidak menerima ratusan email yang tidak berguna (NTF-028).
- Sebagai Product Owner/PM, saya ingin berlangganan hasil saved filter setiap Senin pagi agar saya mendapat ringkasan pekerjaan tanpa membuka aplikasi (NTF-030).
- Sebagai Pengguna Tim Bisnis, saya ingin email yang saya kirim ke alamat project otomatis menjadi work item agar permintaan dari luar tercatat tanpa input manual (NTF-034, NTF-031, NTF-032, NTF-033).
- Sebagai Tamu/Kolaborator Eksternal, saya ingin membalas email notifikasi dan balasan saya menjadi comment agar saya dapat berkolaborasi tanpa login (NTF-025, NTF-032).
- Sebagai Integrator/Developer API, saya ingin mengirim notifikasi kustom terkait work item dan menekan notifikasi saat sinkronisasi massal agar integrasi tidak mengganggu pengguna (NTF-035, NTF-036).
- Sebagai Developer/Anggota Tim, saya ingin menerima notifikasi personal di Slack atau Microsoft Teams dan bertindak langsung dari pesan agar saya tetap responsif di alat chat tim (NTF-039, NTF-040).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 9 |
| P1 | 15 |
| P2 | 15 |
| P3 | 6 |
| **Total** | **45** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 8 |
| Fase 2 | 15 |
| Fase 3 | 12 |
| Fase 4 | 5 |
| Fase — (dc-only/deprecated, rujukan) | 5 |

Strategi urutan implementasi: Fase 1 membangun fondasi yang dipakai semua kanal, yaitu konsumen katalog event sistem (NTF-001) di atas event domain modul 28, pipeline pengiriman asinkron dengan filter permission dan issue security (NTF-008), watchers dan mention (NTF-012, NTF-013), pusat notifikasi in-app dengan satu default scheme tetap (All Watchers, Current Assignee, Reporter), serta infrastruktur SMTP, outgoing mail settings, dan email transaksional akun (NTF-023, NTF-024, NTF-027) agar undangan dan reset password sudah berfungsi. Fase 2 menyalakan email notifikasi work item (NTF-018, NTF-019, NTF-020), menjadikan notification scheme dan recipient type dapat dikonfigurasi (NTF-004 s.d. NTF-007), menghubungkan event ke post function workflow (NTF-002), serta menambah email dari aksi pengguna (bulk change, Share, filter subscription), lokalisasi, dan batas volume. Fase 3 menambah konfigurabilitas tingkat organisasi: custom event, default preferensi, Notification Helper, audit, sender per project, email masuk lengkap (incoming mail server, mail handler, alamat email project bisnis), REST API notifikasi, dan integrasi Automation. Fase 4 berisi custom sending domain, digest, push mobile, dan notifikasi personal Slack/Microsoft Teams. Lima fitur khas Data Center hanya dicatat sebagai rujukan; kebutuhan operasionalnya (antrean internal, visibilitas email) sudah diserap NTF-008, NTF-011, dan NTF-023.

## Fitur

## Event & pemicu notifikasi

### NTF-001 Notification Events (System Events)

- **Deskripsi:** Katalog event sistem yang dipicu saat work item berubah dan menjadi dasar pengiriman notifikasi email dan in-app. Setiap event punya ID numerik tetap (kompatibel dengan ID event produk acuan) yang dipakai notification scheme, post function workflow, dan REST API `/rest/api/3/events`.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem menerbitkan event; Admin Situs melihat daftar event (global permission Administer Jira). Tidak ada aksi user yang langsung mengelola event sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Issue Created (ID 1) dipicu saat work item dibuat lewat UI, REST API, mail handler, clone, atau automation; email memuat semua field yang terisi saat pembuatan.
  2. Issue Updated (ID 2) dipicu untuk perubahan field yang tidak punya event khusus (summary, description, priority, labels, attachment, issue link, due date, custom field, dan sejenisnya); satu operasi edit yang mengubah beberapa field menghasilkan satu event dengan daftar perubahan.
  3. Issue Assigned (ID 3) dipicu saat assignee berubah, termasuk unassign dan auto-assign (default assignee project atau component lead) pada saat create maupun edit.
  4. Issue Resolved (ID 4) hanya dipicu oleh transition yang dikonfigurasi menembakkan event ini (default workflow: transition Resolve Issue), bukan oleh perubahan field Resolution secara langsung.
  5. Issue Closed (ID 5) hanya dipicu oleh transition yang dikonfigurasi menembakkan event ini (default workflow: transition Close Issue).
  6. Issue Commented (ID 6) dipicu saat comment ditambahkan di luar transition; comment yang ditambahkan pada transition screen ikut di event transition.
  7. Issue Comment Edited (ID 14) dipicu saat comment diubah dan Issue Comment Deleted (ID 17) dipicu saat comment dihapus; payload menyimpan visibility comment untuk filter penerima.
  8. Issue Reopened (ID 7) dipicu oleh transition Reopen pada default workflow atau transition lain yang dikonfigurasi menembakkan event ini.
  9. Issue Deleted (ID 8): konten email dirender (snapshot key, summary, project, field utama) sebelum work item dihapus sehingga penerima tetap mendapat ringkasan walaupun tautan tidak lagi dapat dibuka.
  10. Issue Moved (ID 9) dipicu saat work item dipindah antar project atau issue type; email menampilkan key lama dan key baru.
  11. Work Logged On Issue (ID 10), Issue Worklog Updated (ID 15), dan Issue Worklog Deleted (ID 16) dipicu oleh operasi worklog; payload menyimpan visibility worklog.
  12. Work Started On Issue (ID 11) dan Work Stopped On Issue (ID 12) tersedia sebagai event legacy untuk transition Start/Stop Progress pada default workflow.
  13. Generic Event (ID 13) adalah event default untuk transition yang tidak memakai event spesifik.
  14. Satu aksi user memicu satu event utama; misalnya transition dengan comment memicu event transition saja dan comment ikut ditampilkan di email yang sama, bukan dua email terpisah.
  15. Pembuatan sub-task memicu Issue Created pada sub-task, bukan Issue Updated pada parent.
  16. ⚠ Perubahan Rank (drag di board/backlog) tidak mengirim notifikasi email maupun in-app; perubahan field Sprint lewat drag dari backlog juga tidak mengirim notifikasi, tetapi tetap tercatat di history.
  17. ⚠ Import (CSV atau importer) tidak mengirim notifikasi untuk work item yang dibuat atau diubah oleh proses import.
  18. Event dari bulk change hanya menghasilkan notifikasi bila opsi send notification aktif pada langkah konfirmasi (lihat NTF-028).
  19. Event dari REST API dengan query param `notifyUsers=false` tidak menghasilkan email maupun notifikasi in-app event (event tetap diterbitkan untuk webhook dan automation).
- **Team-managed vs company-managed:** Company-managed memakai daftar event ber-ID di atas yang dapat dipetakan di notification scheme. Team-managed memakai daftar event sendiri tanpa ID publik yang dapat dikonfigurasi (NTF-007), tetapi secara internal dipetakan ke event ber-ID yang sama.
- **Catatan & ketergantungan:** Event domain sumber didefinisikan di modul 28; changelog di [08](08-komentar-lampiran-aktivitas.md) CMT-023. ID event dipakai post function (NTF-002, [09](09-workflow-status-transisi.md) WF-016) dan REST API (NTF-036). Konsumen event yang sama juga dipakai webhook (modul 23) dan Automation (modul 20), sehingga keputusan "tidak menotifikasi" hanya menekan kanal notifikasi, bukan event domain.

### NTF-002 Workflow Transition Event Firing

- **Deskripsi:** Konfigurasi post function di company-managed workflow yang menentukan event mana yang ditembakkan sebuah transition, sehingga menentukan notifikasi yang dikirim.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) saat mengedit workflow company-managed.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap transition di company-managed workflow memiliki post function bawaan "Fire a Generic Event that can be processed by the listeners" saat dibuat.
  2. Admin dapat mengedit post function tersebut untuk memilih event lain dari daftar event aktif (Issue Resolved, Issue Closed, Issue Reopened, Work Started, Work Stopped, custom event, dan seterusnya).
  3. Hanya satu event yang ditembakkan per transition; editor workflow menolak penambahan post function fire event kedua dengan pesan galat yang jelas.
  4. Transition Create (initial transition) selalu menembakkan Issue Created dan pilihan event-nya tidak dapat diubah.
  5. ⚠ Post function fire event tidak dapat dihapus dari transition, hanya diganti event-nya; tombol hapus tidak ditampilkan untuk post function ini.
  6. Team-managed workflow tidak menampilkan pilihan event; setiap transition dinotifikasi sebagai perubahan status (event status/transition di NTF-007).
- **Team-managed vs company-managed:** Hanya company-managed yang dapat memilih event per transition; team-managed memakai event perubahan status tetap.
- **Catatan & ketergantungan:** Editor post function di [09](09-workflow-status-transisi.md) WF-029 dan WF-016; daftar event dari NTF-001 dan NTF-003. Workflow default produk ini wajib dikirim dengan pemetaan Resolve/Close/Reopen/Start/Stop Progress ke event yang sesuai.

### NTF-003 Custom Events (Event administration)

- **Deskripsi:** Halaman admin untuk membuat event kustom yang ditembakkan dari post function workflow lalu dipetakan ke recipient di notification scheme.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > System > Events menampilkan daftar event dengan kolom nama, deskripsi, template email terkait, status Active/Inactive, serta notification scheme dan workflow yang memakainya.
  2. Add event meminta nama (wajib, unik tanpa memperhatikan huruf besar/kecil; duplikat ditolak dengan pesan galat), deskripsi opsional, dan pilihan template dari template bawaan (misalnya Generic Event, Issue Created, Issue Commented).
  3. Event berstatus Active bila dipakai minimal satu notification scheme (memiliki recipient) atau satu post function workflow; selain itu Inactive.
  4. Hanya custom event berstatus Inactive yang dapat dihapus; system event tidak dapat dihapus dan tombol hapus tidak ditampilkan.
  5. Admin dapat mengedit nama dan deskripsi custom event; perubahan nama langsung tampil di scheme dan editor workflow.
  6. Custom event dapat dipilih di post function Fire Event pada transition (NTF-002).
  7. Custom event baru otomatis muncul di setiap notification scheme dengan daftar recipient kosong.
- **Team-managed vs company-managed:** Hanya untuk company-managed project; team-managed tidak mendukung custom event.
- **Catatan & ketergantungan:** Membuat template email baru untuk custom event tidak didukung (lihat NTF-042); custom event selalu memakai salah satu template bawaan. Halaman admin berada di [22](22-administrasi-sistem.md).

## Skema & penerima notifikasi

### NTF-004 Notification Scheme (company-managed)

- **Deskripsi:** Objek konfigurasi global yang memetakan setiap event ke daftar recipient; satu scheme dapat dipakai banyak company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) membuat dan mengubah scheme.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman daftar Settings > Issues > Notification schemes menampilkan nama, deskripsi, dan project yang memakai setiap scheme.
  2. Add notification scheme meminta nama unik dan deskripsi opsional; scheme baru tidak memiliki recipient di event mana pun.
  3. Copy scheme membuat salinan lengkap dengan nama "Copy of <nama>" dan tidak terasosiasi ke project mana pun.
  4. Admin dapat mengedit nama dan deskripsi scheme; nama tetap harus unik.
  5. Delete scheme hanya diizinkan bila scheme tidak terasosiasi ke project; bila masih terasosiasi, tombol delete dinonaktifkan dengan keterangan daftar project.
  6. Halaman detail menampilkan tabel semua event (system dan custom) beserta recipient per event.
  7. Add notification memungkinkan memilih satu atau lebih event sekaligus lalu satu recipient type beserta parameternya (misalnya user tertentu, nama group, project role, custom field).
  8. Admin dapat menghapus recipient dari sebuah event tanpa memengaruhi event lain.
  9. Recipient yang sama (tipe dan parameter identik) tidak dapat ditambahkan dua kali pada event yang sama; percobaan ditolak dengan pesan galat.
  10. Default Notification Scheme bawaan tersedia sejak instalasi dan memetakan event utama ke All Watchers, Current Assignee, dan Reporter; scheme ini dipakai project company-managed baru secara default.
  11. Perubahan scheme langsung berlaku untuk semua project terkait pada event berikutnya tanpa perlu re-asosiasi.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai konfigurasi per project (NTF-007).
- **Catatan & ketergantungan:** Recipient type di NTF-006, asosiasi di NTF-005, REST di NTF-036, audit di NTF-011. Di Fase 1 Default Notification Scheme sudah ada sebagai data seed yang dipakai pipeline (NTF-008) walaupun UI editornya baru hadir di Fase 2.

### NTF-005 Notification Scheme Association

- **Deskripsi:** Pengaitan notification scheme ke company-managed project melalui project settings.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects) melihat; Admin Situs (Administer Jira) mengganti dan mengedit scheme.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Notifications menampilkan nama scheme aktif beserta ringkasan event dan recipient; halaman dapat dilihat Admin Proyek.
  2. Actions > Use a different scheme menampilkan daftar scheme dan pilihan None; aksi ini hanya tersedia bagi Admin Situs.
  3. Pilihan None berarti project tidak mengirim email notifikasi event sama sekali (mention dan Share tetap terkirim karena tidak bergantung scheme).
  4. Actions > Edit notifications membawa Admin Situs ke halaman detail scheme (NTF-004); Admin Proyek tidak melihat aksi ini.
  5. Halaman menampilkan indikator bahwa scheme dipakai bersama oleh N project beserta daftar project tersebut.
  6. Project baru yang dibuat dengan opsi "share settings with an existing project" ikut memakai notification scheme yang sama dengan project sumber.
  7. Setiap perubahan asosiasi (scheme lama, scheme baru, project, actor) tercatat di audit log.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Share settings di [04](04-proyek.md); audit di NTF-011 dan [22](22-administrasi-sistem.md).

### NTF-006 Notification Recipient Types

- **Deskripsi:** Jenis penerima yang dapat dipasang pada event dalam notification scheme; penerima akhir dievaluasi saat event terjadi, bukan saat konfigurasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengonfigurasi; Sistem mengevaluasi saat event.
- **Aturan perilaku (acceptance criteria):**
  1. All Watchers: semua watcher work item pada saat event diproses.
  2. ⚠ Current Assignee: assignee setelah perubahan; pada event Issue Assigned, assignee sebelumnya juga dinotifikasi (sekali, dengan alasan "You were previously assigned") agar mengetahui pengalihan tugas, dan tetap dideduplikasi bila cocok dengan recipient lain (misalnya watcher).
  3. Reporter: nilai field Reporter work item.
  4. Current User: user pemicu event; tetap tunduk pada preferensi "notify own changes" (NTF-020).
  5. Project Lead: lead project work item.
  6. Component Lead: lead setiap component yang terpasang pada work item; work item tanpa component tidak menghasilkan penerima dari tipe ini.
  7. User: satu user tertentu yang dipilih admin.
  8. Group: semua anggota aktif group tertentu.
  9. Project Role: semua user dan anggota group dalam role tersebut pada project work item.
  10. User Custom Field Value: user pada custom field user picker single maupun multi.
  11. Group Custom Field Value: anggota group pada custom field group picker.
  12. ⚠ Single Email Address: email dikirim dalam format Text ke alamat tersebut hanya bila work item dapat dilihat anonim (Browse Projects untuk Anyone); bila tidak, alamat dilewati dan dicatat di log pengiriman.
  13. Deduplikasi: user yang cocok dengan beberapa recipient type hanya menerima satu email dan satu notifikasi in-app per event; alasan penerimaan di footer memakai alasan paling spesifik (assignee > reporter > watcher > lainnya).
  14. Semua penerima tetap difilter oleh Browse Projects dan issue security level sebelum dikirimi (NTF-008).
- **Team-managed vs company-managed:** Team-managed hanya mendukung subset recipient (NTF-007).
- **Catatan & ketergantungan:** Minimal MVP Fase 2: All Watchers, Current Assignee, Reporter, User, Group, Project Role; sisanya dapat menyusul di Fase 3. Group dan role dari [02](02-pengguna-grup-organisasi.md) dan [03](03-izin-peran-keamanan.md); custom field picker dari [05](05-konfigurasi-work-item-skema.md).

### NTF-007 Team-managed Project Notifications

- **Deskripsi:** Halaman notifikasi per project untuk team-managed project yang dikonfigurasi Admin Proyek tanpa scheme global.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (role Administrator pada team-managed project).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Notifications menampilkan daftar event beserta recipient per event.
  2. ⚠ Event yang tersedia: work item created, work item updated, work item assigned, work item deleted, status changed (transition), comment added, comment edited, comment deleted, worklog added/updated/deleted.
  3. ⚠ Recipient per event dapat diedit dengan pilihan All watchers, Current assignee, Reporter, User, Group, dan Project role (Administrator, Member, Viewer, serta custom role); recipient dari user picker field tidak disediakan.
  4. ⚠ Setiap event memiliki toggle aktif/nonaktif; event nonaktif tidak mengirim email untuk project tersebut tanpa menghapus daftar recipient-nya.
  5. Konfigurasi terisolasi hanya untuk project tersebut dan tidak memengaruhi project lain.
  6. Project baru mendapat default All watchers, Current assignee, dan Reporter pada sebagian besar event.
  7. ⚠ Tidak ada custom event, tidak ada event Resolved/Closed/Reopened terpisah, dan tidak ada recipient Project Lead atau Component Lead.
- **Team-managed vs company-managed:** Fitur ini khusus team-managed; padanan company-managed adalah NTF-004 dan NTF-005.
- **Catatan & ketergantungan:** Secara internal disimpan sebagai notification scheme privat milik project (tidak tampil di daftar scheme global) agar pipeline NTF-008 memakai satu jalur evaluasi. Role team-managed di [03](03-izin-peran-keamanan.md).

## Pipeline pengiriman & aturan inti

### NTF-008 Notification Delivery Rules & Permission Filtering

- **Alias inventaris:** Notification Delivery Pipeline
- **Deskripsi:** Pipeline terpadu dan aturan inti yang menentukan siapa benar-benar menerima notifikasi email, in-app, dan real-time setelah recipient dievaluasi, serta bagaimana pengiriman dilakukan secara asinkron dan andal pada volume tinggi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem; aturan memakai Browse Projects (PERM-020), issue security (PERM-044), dan visibility comment/worklog.
- **Aturan perilaku (acceptance criteria):**
  1. Penerima wajib memiliki Browse Projects pada project work item saat event diproses; user tanpa permission ini dilewati walaupun tercantum sebagai watcher atau recipient eksplisit.
  2. Work item dengan security level hanya dikirim ke user yang termasuk level tersebut.
  3. Comment atau worklog dengan visibility terbatas (role/group) hanya dikirim ke anggota role/group tersebut; penerima lain pada event yang sama tidak mendapat notifikasi untuk comment/worklog itu.
  4. Field yang disembunyikan oleh field configuration tidak ditampilkan di email maupun preview in-app.
  5. Pemicu perubahan (actor) tidak menerima notifikasi atas aksinya sendiri kecuali mengaktifkan preferensi notify own changes (NTF-020).
  6. User deactivated, user tanpa product access, dan user tanpa alamat email tidak dikirimi email; user tanpa product access juga tidak menerima notifikasi in-app.
  7. Email dikirim ke alamat email utama akun user (modul 01), bukan alamat sekunder.
  8. Preferensi personal hanya dapat mengurangi notifikasi yang dihasilkan scheme, tidak dapat menambah event atau work item yang tidak dicakup scheme (pengecualian: mention dan Share yang memang di luar scheme).
  9. Notifikasi diproses asinkron melalui antrean (transactional outbox: event ditulis dalam transaksi yang sama dengan perubahan work item, lalu dikonsumsi worker); kegagalan pengiriman tidak membatalkan atau memperlambat transaksi work item.
  10. Pengiriman email yang gagal di-retry dengan exponential backoff (misalnya 1, 5, 15, 60 menit, maksimal 5 kali) lalu dipindahkan ke dead-letter queue internal yang dapat dilihat admin lewat log pengiriman (NTF-011).
  11. Urutan pipeline baku: event domain → resolver penerima (notification scheme atau konfigurasi team-managed, watchers, mentions) → filter permission dan issue security → preferensi user per kanal → kanal (email, in-app, real-time, push, chat).
  12. Deduplikasi penerima dilakukan sebelum batching sehingga satu user mendapat paling banyak satu notifikasi per event per kanal, lalu email masuk ke jendela batching (NTF-019).
  13. Notifikasi in-app disimpan dengan status read/unread, kategori Direct/Watching, dan masa retensi (NTF-015).
  14. Preferensi user (notify own changes, toggle per jenis aktivitas, per kanal) dievaluasi per penerima pada saat pengiriman, bukan saat event dibuat.
  15. Bulk change dengan opsi send notification nonaktif tidak menghasilkan email maupun notifikasi in-app event.
  16. Pengiriman asinkron tidak memblokir request utama; target p95 penambahan latensi pada request perubahan work item akibat notifikasi < 10 ms (hanya penulisan outbox).
- **Team-managed vs company-managed:** Tidak ada perbedaan; sumber recipient berbeda (scheme global vs konfigurasi project) tetapi aturan filter dan pengiriman sama.
- **Catatan & ketergantungan:** Implementasi worker dan antrean mengikuti [27](27-nfr-arsitektur-platform.md); event domain dan outbox di [28](28-model-data-inti.md). Pengecekan permission memakai service otorisasi yang sama dengan UI ([03](03-izin-peran-keamanan.md)) agar tidak ada jalur kebocoran. Repo belum memiliki worker background; modul ini memperkenalkan worker in-process pertama (goroutine pool + tabel outbox PostgreSQL) yang dapat dipisah menjadi proses tersendiri kemudian.

### NTF-009 Notification Limits & Throttling

- **Deskripsi:** Batasan sistem terkait volume dan ukuran notifikasi untuk mencegah banjir email dan menjaga reputasi pengirim.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Sistem; Admin Situs dapat mengubah nilai batas di konfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Batching (NTF-019) dipakai sebagai mekanisme utama menekan volume email per work item per penerima.
  2. Proses import tidak mengirim notifikasi (konsisten dengan NTF-001 butir 17).
  3. ⚠ Ukuran email notifikasi dibatasi (default 1 MB HTML); konten yang melebihi batas dipotong dengan tautan "View full work item".
  4. ⚠ Jumlah recipient per event dibatasi (default 1.000 user setelah resolusi group); bila terlampaui, notifikasi dikirim ke 1.000 penerima pertama berdasarkan prioritas recipient langsung (assignee, reporter, watcher) dan admin mendapat peringatan di log.
  5. Endpoint `/notify` (NTF-035) memiliki rate limit tersendiri per user dan per site; permintaan berlebih dijawab HTTP 429 dengan header `Retry-After`.
  6. Ukuran email masuk dibatasi mengikuti batas ukuran attachment site; email yang lebih besar ditolak dan dicatat di log mail handler.
  7. Clone menerapkan rate limit pengiriman per user dan per site (misalnya maksimal 500 email per penerima per jam) serta proteksi email storm: rule automation yang memicu lebih dari ambang notifikasi per menit (default 200) dihentikan sementara (status throttled) dan pemilik rule diberi tahu lewat in-app dan email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rate limit API umum di [23](23-api-publik-webhook.md) dan [27](27-nfr-arsitektur-platform.md); loop automation di [20](20-otomasi.md). Semua angka default adalah keputusan clone dan dapat dikonfigurasi lewat environment.

### NTF-010 Notification Localization & Timezone

- **Deskripsi:** Penyesuaian bahasa dan waktu notifikasi per penerima.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Sistem; preferensi bahasa dan timezone diatur user di profil (modul 21).
- **Aturan perilaku (acceptance criteria):**
  1. Email dan notifikasi in-app dirender dalam bahasa preferensi masing-masing penerima, bukan bahasa actor; satu event dengan penerima berbahasa berbeda menghasilkan render terpisah per bahasa.
  2. Timestamp di email (waktu perubahan, due date bertipe datetime, worklog) ditampilkan dalam timezone penerima beserta label zona.
  3. Nama status, issue type, priority, dan field mengikuti terjemahan bila tersedia untuk bahasa penerima.
  4. Bila bahasa penerima tidak tersedia atau belum diatur, render memakai bahasa default site.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Repo saat ini mendukung EN/ID lewat `internal/shared/constants/error.go` dan `LanguageMiddleware()`; template email wajib memakai katalog terjemahan yang sama (EN dan ID sebagai bahasa awal). Preferensi bahasa/timezone di [21](21-navigasi-ux-preferensi.md).

### NTF-011 Notification Audit & Logging

- **Deskripsi:** Pencatatan perubahan konfigurasi notifikasi dan jejak pengiriman email untuk troubleshooting dan kepatuhan.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) melihat audit log dan log pengiriman.
- **Aturan perilaku (acceptance criteria):**
  1. Audit log mencatat create, edit, delete notification scheme, penambahan/penghapusan recipient, dan perubahan asosiasi scheme ke project, lengkap dengan actor, waktu, nilai lama, dan nilai baru.
  2. Audit log mencatat perubahan outgoing mail settings, incoming mail server, dan mail handler (tanpa menyimpan password atau token dalam teks).
  3. Log pemrosesan mail handler mencatat setiap email masuk: pengirim, subject, handler, hasil (work item dibuat, comment ditambahkan, diabaikan, gagal) dan alasan.
  4. Clone menyimpan log pengiriman per email (penerima, event, work item, status sent/failed/suppressed, jumlah retry, pesan galat SMTP) dengan retensi default 30 hari dan pencarian berdasarkan penerima atau work item.
- **Team-managed vs company-managed:** Perubahan konfigurasi notifikasi team-managed tercatat di audit log project tersebut.
- **Catatan & ketergantungan:** Audit log global di [22](22-administrasi-sistem.md); log pengiriman menggantikan Mail Queue DC (NTF-041) dan dipakai Notification Helper (NTF-022).

## Watch, mention & autowatch

### NTF-012 Watchers

- **Deskripsi:** Mekanisme mengikuti work item agar user menerima notifikasi untuk event yang memasukkan All Watchers sebagai recipient.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (watch diri sendiri dengan Browse Projects), Admin Proyek; View Voters and Watchers dan Manage Watchers (PERM-033).
- **Aturan perilaku (acceptance criteria):**
  1. Issue view menampilkan tombol Watch/Stop watching (ikon mata) beserta jumlah watcher; klik mengubah status watch user saat ini tanpa reload.
  2. Daftar nama watchers hanya terlihat bagi user dengan permission View Voters and Watchers; user lain hanya melihat jumlah.
  3. Menambah atau menghapus watcher lain membutuhkan permission Manage Watchers; tanpa permission itu kontrol tambah/hapus tidak ditampilkan dan API menolak dengan 403.
  4. Saat menambah watcher yang tidak memiliki Browse Projects, UI menampilkan peringatan bahwa user tersebut tidak akan menerima notifikasi; watcher tetap tersimpan.
  5. ⚠ Global setting "Allow users to watch issues" di General configuration dapat menonaktifkan fitur watch; bila nonaktif, tombol Watch disembunyikan dan recipient All Watchers tidak menghasilkan penerima.
  6. Keyboard shortcut `w` pada issue view men-toggle status watch user saat ini (Watch ↔ Stop watching) tanpa membuka panel, sama dengan klik tombol Watch; shortcut tidak aktif saat fokus berada di field input. Katalog shortcut dimiliki [21](21-navigasi-ux-preferensi.md) NAV-041; butir ini hanya merujuk perilaku yang ditetapkan di sana.
  7. JQL `watcher = currentUser()` dan filter bawaan "Watched work items" menampilkan work item yang di-watch user.
  8. Bulk change menyediakan operasi Watch dan Stop watching untuk diri sendiri pada banyak work item.
  9. Tautan "Stop watching" di footer email langsung menghapus user dari watcher work item tersebut setelah konfirmasi satu klik (tautan bertoken, berlaku meski sesi tidak aktif).
  10. ⚠ Mention tidak otomatis menambahkan user yang di-mention sebagai watcher.
  11. Opsi CC watchers pada mail handler menambahkan penerima CC email yang cocok dengan user sebagai watcher work item (NTF-032).
- **Team-managed vs company-managed:** Tidak ada perbedaan; di team-managed, permission watcher mengikuti role Administrator/Member.
- **Catatan & ketergantungan:** Model data dan UI dasar watcher juga dijelaskan di [06](06-work-item-inti.md) ISS-036 dan [28](28-model-data-inti.md) DM-008; modul ini memiliki dampak notifikasinya. Permission di [03](03-izin-peran-keamanan.md) PERM-033. JQL di [12](12-jql-bahasa-kueri.md). Penyelarasan fase: karena default notification scheme Fase 1 memakai recipient All Watchers, watch diri sendiri (butir 1, 6, 9), model data watcher (DM-008), dan bagian dasar PERM-033 (View Voters and Watchers, Manage Watchers dengan grant default ke Administrator proyek) harus tersedia di Fase 1; ISS-036 dan DM-008 diusulkan dipindah ke Fase 1 di modul masing-masing. Butir yang bergantung pada Fase 2 (JQL `watcher` dan filter bawaan pada butir 7, bulk change pada butir 8, mail handler pada butir 11) baru aktif ketika fitur pendukungnya hadir. Pemilik katalog keyboard shortcut adalah [21](21-navigasi-ux-preferensi.md) NAV-041.

### NTF-013 Mention Notifications

- **Deskripsi:** Notifikasi langsung kepada user yang di-@mention di description, comment, atau rich text field.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim; mention membutuhkan Browse Users and Groups (PERM-008) untuk autocomplete.
- **Aturan perilaku (acceptance criteria):**
  1. Autocomplete @mention memprioritaskan user terkait work item (assignee, reporter, watcher, commenter) lalu anggota project, lalu user lain.
  2. Mention memicu notifikasi terlepas dari notification scheme, mengikuti preferensi personal "You're mentioned" (NTF-020); project dengan scheme None tetap mengirim notifikasi mention.
  3. Saat me-mention user yang tidak memiliki akses ke work item, editor menampilkan peringatan bahwa user tersebut tidak akan dinotifikasi.
  4. Mention di comment dengan visibility terbatas hanya dinotifikasi bila user yang di-mention termasuk role/group visibility tersebut.
  5. Deduplikasi dengan email event: user yang di-mention sekaligus watcher hanya menerima satu email, dengan header "X mentioned you" ditonjolkan.
  6. Notifikasi in-app "X mentioned you" muncul di tab Direct pusat notifikasi.
  7. Edit comment atau description hanya menotifikasi mention yang baru ditambahkan, bukan mention yang sudah ada sebelumnya.
  8. ⚠ Mention team (Teams, modul 02) menotifikasi seluruh anggota team yang memiliki akses ke work item, tunduk pada batas recipient NTF-009.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Node mention, parsing ADF, dan typeahead di [07](07-editor-konten-kaya.md) EDT-014; modul ini hanya pengiriman notifikasinya.

### NTF-014 Autowatch

- **Deskripsi:** Pengaturan agar user otomatis menjadi watcher pada work item yang ia buat atau komentari.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim mengatur preferensi sendiri; Admin Situs mengatur default (NTF-021).
- **Aturan perilaku (acceptance criteria):**
  1. Personal settings menyediakan opsi Autowatch Enabled/Disabled.
  2. Nilai default untuk user baru mengikuti User Default Settings (NTF-021); default pabrik Enabled.
  3. Bila aktif, user otomatis ditambahkan sebagai watcher saat membuat work item (sebagai reporter) dan saat menambah comment.
  4. ⚠ Bila user telah melakukan Stop watching secara manual pada sebuah work item, aksi comment berikutnya tidak menambahkannya kembali sebagai watcher.
  5. ⚠ Autowatch tidak berlaku untuk work item hasil import.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman personal settings di [21](21-navigasi-ux-preferensi.md); comment di [08](08-komentar-lampiran-aktivitas.md). Butir 4 membutuhkan penanda "explicitly unwatched" per user per work item. Pembagian kepemilikan: NTF = pipeline, event, dan preferensi; NAV = UI panel/toggle. Toggle Autowatch di halaman personal settings dirinci di NAV-031 ([21](21-navigasi-ux-preferensi.md)); modul ini memiliki logika penambahan watcher otomatis dan nilai default-nya.

## Notifikasi in-app

### NTF-015 In-product Notifications Hub

- **Deskripsi:** Panel notifikasi (ikon lonceng) di top navigation yang menampilkan notifikasi in-app untuk seluruh produk di instalasi ini (work management, Service Management, Product Discovery).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Semua user yang login dengan product access; item hanya tampil bila user masih berhak melihat sumbernya.
- **Aturan perilaku (acceptance criteria):**
  1. Ikon lonceng di top navigation menampilkan badge jumlah notifikasi belum dibaca (ditampilkan "9+" bila lebih dari 9).
  2. Panel memiliki tab Direct (mention, assignment, comment/aktivitas pada work item yang dilaporkan atau ditugaskan ke user) dan tab Watching (aktivitas pada work item yang di-watch).
  3. Toggle "Only show unread" menyaring daftar hanya ke item belum dibaca; status toggle diingat per user.
  4. Setiap item dapat ditandai read/unread, dan tombol "Mark all as read" menandai semua item pada tab aktif sebagai terbaca.
  5. Notifikasi dikelompokkan per work item (mis. "3 updates on KEY-123") dengan kemampuan expand untuk melihat item individual.
  6. ⚠ Panel menyediakan filter per produk (work management, Service Management, Product Discovery) bila lebih dari satu produk aktif.
  7. Klik notifikasi membuka work item terkait (dengan anchor ke comment bila relevan) dan menandai item sebagai terbaca.
  8. Item menampilkan preview isi comment atau ringkasan perubahan (field lama → baru) tanpa harus membuka work item.
  9. Panel memiliki tautan ke halaman notification settings (NTF-016, NTF-020).
  10. ⚠ Notifikasi disimpan dengan retensi terbatas (default 30 hari untuk item terbaca, 90 hari untuk item belum dibaca) lalu dihapus otomatis oleh job terjadwal.
  11. Item dari work item yang aksesnya dicabut (permission, security level, atau work item dihapus) disembunyikan pada saat panel dimuat.
  12. Notifikasi baru muncul real-time tanpa reload halaman (WebSocket atau Server-Sent Events) dan badge diperbarui; bila koneksi real-time gagal, klien melakukan polling berkala.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan: NTF = pipeline, event, dan preferensi; NAV = UI panel/toggle. UI panel lonceng (tata letak, tab, toggle) dirinci di NAV-018 ([21](21-navigasi-ux-preferensi.md)); NTF-015 memiliki pembuatan item notifikasi, penyimpanan, retensi, filter akses, status read/unread, dan pengiriman real-time. Top navigation di [21](21-navigasi-ux-preferensi.md); transport real-time di [27](27-nfr-arsitektur-platform.md). Di Fase 1 hub memakai Default Notification Scheme tetap (watchers, assignee, reporter) plus mention. Endpoint usulan: `GET /api/v1/notifications`, `PATCH /api/v1/notifications/{id}`, `POST /api/v1/notifications/read-all`.

### NTF-016 In-app Notification Preferences

- **Deskripsi:** Pengaturan per user untuk notifikasi in-app yang terpisah dari email.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Setiap user untuk dirinya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. User dapat mengaktifkan/menonaktifkan kategori in-app: mention, assigned, comments, dan status changes pada work item yang di-watch atau dilaporkan.
  2. Halaman notification settings memisahkan kanal Email, In-app, dan Mobile push; mematikan satu kanal tidak memengaruhi kanal lain.
  3. Preferensi "notify own changes" berlaku juga untuk in-app; bila nonaktif, aksi user sendiri tidak muncul di hub.
  4. ⚠ Notifikasi in-app event mengikuti hasil resolusi notification scheme project yang sama dengan email (project dengan scheme None tidak menghasilkan notifikasi in-app event, kecuali mention dan assignment langsung).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disajikan di halaman yang sama dengan NTF-020; push di NTF-038.

### NTF-017 Notification Digest / Summary Emails

- **Deskripsi:** Email ringkasan berkala berisi notifikasi in-app yang belum dibaca.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Setiap user untuk dirinya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Sistem dapat mengirim email ringkasan berisi notifikasi in-app yang belum dibaca sejak digest terakhir; bila tidak ada item belum dibaca, email tidak dikirim.
  2. ⚠ Frekuensi digest dapat dipilih Harian atau Mingguan dengan jam kirim di timezone penerima; default nonaktif.
  3. User dapat opt-out dari digest di notification settings dan lewat tautan di footer email digest.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Ketersediaan fitur ini di produk acuan belum pasti; kebutuhan digest periodik di produk acuan umumnya dipenuhi lewat Filter Subscriptions (NTF-030). Dibangun hanya bila pemilik produk menyetujui (lihat Pertanyaan terbuka).

## Email notifikasi work item

### NTF-018 Email Notification Content

- **Deskripsi:** Struktur dan isi email notifikasi work item yang dihasilkan untuk setiap event.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Sistem menghasilkan; penerima akhir membaca. Isi tunduk pada filter NTF-008.
- **Aturan perilaku (acceptance criteria):**
  1. Subject berisi prefix site, issue key, dan summary dengan format tetap (mis. `[PRODUK] (KEY-123) Summary`) sehingga konsisten untuk semua email work item yang sama.
  2. Header email menampilkan avatar dan nama actor beserta ringkasan aksi (mis. "X changed 2 fields", "X mentioned you", "X created a Bug").
  3. Isi per event: Created menampilkan field terisi; Updated menampilkan nilai lama dicoret lalu nilai baru; Commented menampilkan isi comment; Assigned menampilkan assignee baru; Transitioned menampilkan status lama → baru; Deleted menampilkan ringkasan snapshot; Moved menampilkan key/project lama dan baru.
  4. Rich text (ADF) dirender ke HTML email yang aman (CSS inline, tanpa script): mention, emoji, code block, table, link, dan panel; elemen yang tidak didukung diganti teks fallback.
  5. Attachment ditampilkan sebagai nama file dan thumbnail (untuk gambar); file tidak dilampirkan ke email.
  6. Email memuat tombol/tautan "View work item" dan "Add comment" yang membuka work item (dan fokus ke editor comment).
  7. Footer menjelaskan alasan menerima email (watching, reporter, assignee, mentioned, recipient scheme), tautan Manage notifications, dan tautan Stop watching.
  8. Header threading `Message-ID`, `In-Reply-To`, dan `References` disusun sehingga semua email satu work item tergabung dalam satu thread di mail client.
  9. ⚠ Email menyertakan header `Auto-Submitted: auto-generated` untuk mencegah auto-reply loop serta `List-Unsubscribe` (dan `List-Unsubscribe-Post` one-click) yang mengarah ke halaman pengaturan notifikasi.
  10. Email dikirim multipart: HTML dan plain text, dengan bagian yang ditampilkan mengikuti preferensi format penerima (NTF-020).
  11. Email menampilkan nama project, ikon issue type, dan parent (mis. Epic) work item.
  12. ⚠ Konten panjang (description atau comment melebihi ±20.000 karakter, atau tabel perubahan melebihi 50 field) dipotong dengan tautan ke work item.
  13. ⚠ Alamat From/Reply-To default bersifat no-reply; membalas email notifikasi tidak menambah comment kecuali project email address diarahkan ke mailbox yang diproses mail handler (NTF-025, NTF-032).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Render ADF di [07](07-editor-konten-kaya.md); field tersembunyi di [05](05-konfigurasi-work-item-skema.md); bahasa di NTF-010. Nama merek di prefix subject menunggu nama final produk (lihat [00](00-visi-ruang-lingkup-glosarium-baseline.md)). Template disimpan di kode (Go `html/template` + `text/template`) dan tidak dapat dikustomisasi admin (NTF-042).

### NTF-019 Batched Email Notifications

- **Deskripsi:** Penggabungan beberapa perubahan pada work item yang sama dalam jendela waktu singkat menjadi satu email per work item per penerima.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Sistem; penerima akhir.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Jendela batching dimulai sejak perubahan pertama yang belum terkirim dan berlangsung 10 menit (dapat dikonfigurasi admin); email dikirim saat jendela berakhir.
  2. Satu email batch dapat menggabungkan perubahan dari beberapa actor dan menampilkan nama setiap actor.
  3. Ringkasan batch disusun berurutan per jenis: comment, perubahan field, lalu transition, masing-masing berurutan waktu.
  4. Mention dan assignment kepada penerima ditonjolkan di bagian atas email batch.
  5. ⚠ Event Issue Created dan Issue Deleted dikirim segera tanpa batching; perubahan berikutnya setelah Created memulai jendela batch baru.
  6. ⚠ User dapat memilih untuk tidak memakai batching di pengaturan personal; bila dipilih, setiap event dikirim sebagai email terpisah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada outbox NTF-008; batch dipakai juga oleh bulk change (NTF-028). Perilaku Data Center (email per event) dapat direproduksi dengan mengatur jendela 0 menit.

### NTF-020 Personal Email Notification Settings

- **Deskripsi:** Halaman pengaturan notifikasi pribadi (Personal settings > Notifications) tempat user memilih email yang ingin diterima, formatnya, dan kanalnya.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Setiap user untuk dirinya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Master toggle mematikan semua email notifikasi work item; email akun/transaksional (NTF-027) tetap dikirim.
  2. Toggle per peran: email saat saya watching, saat saya reporter, saat saya assignee, dan saat saya di-mention.
  3. ⚠ Toggle per jenis aktivitas: work item created, comments, status changes, assignment, field updates, dan worklog.
  4. Opsi "You make changes to work items" (notify own changes) dengan default off.
  5. Pilihan format email HTML (default) atau Text; pilihan ini juga berlaku untuk email filter subscription (NTF-030).
  6. ⚠ Untuk setiap kategori, user dapat memilih kanal email, in-app, serta Slack/Microsoft Teams bila akun telah terhubung (NTF-039, NTF-040).
  7. Tautan Manage notifications di footer email membuka halaman ini langsung (dengan login bila sesi tidak aktif).
  8. Perubahan disimpan segera dan berlaku untuk event berikutnya; notifikasi yang sudah berada di jendela batch dievaluasi ulang saat dikirim.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman personal settings di [21](21-navigasi-ux-preferensi.md). Preferensi disimpan dalam tabel `notification_preferences` (user, kanal, kategori, nilai) agar kanal baru dapat ditambah tanpa migrasi kolom. Mencakup item inventaris "Email Format Preference".

### NTF-021 User Default Settings (Notification defaults)

- **Deskripsi:** Pengaturan admin untuk nilai default preferensi notifikasi bagi user baru.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat menetapkan default email format (HTML/Text).
  2. Admin dapat menetapkan default notify users of their own changes (Yes/No).
  3. Admin dapat menetapkan default autowatch (Enabled/Disabled).
  4. ⚠ Admin dapat memilih "Apply to existing users" yang menimpa nilai preferensi terkait milik semua user yang sudah ada setelah konfirmasi; aksi ini tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman di [22](22-administrasi-sistem.md); dipakai NTF-014 dan NTF-020.

### NTF-022 Notification Helper

- **Deskripsi:** Alat troubleshooting admin untuk mengetahui apakah user tertentu menerima notifikasi untuk work item dan event tertentu beserta alasannya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Input berupa user, work item (key), dan event (dari daftar NTF-001/NTF-003).
  2. Output menyatakan ya/tidak beserta recipient type yang cocok (mis. "Current Assignee", "Group: dev-team").
  3. Bila tidak menerima, output menyebut alasan: tidak punya Browse Projects, tidak termasuk issue security level, scheme tidak memuat user, preferensi personal menonaktifkan, user deactivated/tanpa email, atau outgoing mail nonaktif.
  4. Output menyediakan tautan ke notification scheme dan permission scheme project terkait.
  5. ⚠ Alat diakses dari Settings > System > Admin helper dan juga dari menu "…" di issue view bagi Admin Situs (work item terisi otomatis).
- **Team-managed vs company-managed:** Mendukung kedua jenis; untuk team-managed, rujukan scheme diganti tautan ke halaman Notifications project.
- **Catatan & ketergantungan:** Memakai resolver yang sama dengan pipeline NTF-008 dalam mode dry-run; log pengiriman NTF-011 ditampilkan sebagai bukti kiriman terakhir. Padanan Permission Helper di [03](03-izin-peran-keamanan.md).

## Infrastruktur & pengaturan email keluar

### NTF-023 Email Infrastructure (SMTP Outbound / IMAP Inbound)

- **Alias inventaris:** SMTP Mail Server Configuration
- **Deskripsi:** Infrastruktur pengiriman email keluar lewat SMTP milik pelanggan dan penerimaan email masuk (IMAP polling atau webhook inbound) yang andal. Karena produk ini self-hosted, konfigurasi SMTP server (di produk acuan hanya ada di Data Center) wajib disediakan.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Admin Situs (Administer Jira / system admin); Engineering untuk konfigurasi environment.
- **Aturan perilaku (acceptance criteria):**
  1. SMTP outbound mendukung TLS implisit dan STARTTLS serta autentikasi; from address dan sender name dapat ditimpa per project (NTF-025).
  2. Custom sending domain dengan verifikasi DNS SPF/DKIM tersedia sebagai lapisan opsional (NTF-026).
  3. Email keluar melewati antrean dengan retry dan rate limit per domain tujuan (default maksimal 5 koneksi paralel dan 60 pesan per menit per domain, dapat dikonfigurasi lewat environment) agar tidak ditolak penyedia mail; pesan yang tertahan rate limit tetap di antrean, bukan dihitung gagal.
  4. Batched notifications menggabungkan perubahan beruntun dalam jendela waktu (default ±10 menit, NTF-019).
  5. Setiap email memakai template HTML + plaintext dan dirender dalam bahasa serta timezone penerima (NTF-010).
  6. Header `List-Unsubscribe`, `Message-ID`, `In-Reply-To`, dan `References` diisi untuk threading dan unsubscribe (NTF-018).
  7. ⚠ Inbound mendukung IMAP polling (NTF-031) dan endpoint webhook inbound (format JSON/MIME dari penyedia seperti SES/Mailgun/Postmark) untuk create work item/add comment dengan strip quoted reply; cakupan software project mengikuti NTF-032, sedangkan business project memakai NTF-034.
  8. Pencegahan loop: email masuk dengan header `Auto-Submitted` selain `no`, `Precedence: bulk/junk/list`, bounce (DSN), atau berasal dari alamat pengirim produk ini sendiri diabaikan.
  9. Bounce handling: hard bounce dan complaint memasukkan alamat ke suppression list sehingga email berikutnya ke alamat itu tidak dikirim (status "suppressed" di log); admin dapat menghapus alamat dari suppression list.
  10. ⚠ Mode development memakai mail catcher (Mailpit) lewat konfigurasi `SMTP_HOST`/`SMTP_PORT`, dan tersedia toggle global disable outgoing mail (NTF-024).
  11. Konfigurasi SMTP server memuat name, description, from address, dan email prefix.
  12. Konfigurasi koneksi memuat protocol SMTP atau SECURE_SMTP, host, port, timeout (detik), dan mode TLS/STARTTLS; port default 587 (STARTTLS) atau 465 (SMTPS).
  13. Kredensial berupa username/password disimpan terenkripsi dan tidak pernah ditampilkan kembali di UI; JNDI location tidak didukung dan field-nya tidak disediakan.
  14. Tombol "Send a test email" mengirim email uji ke alamat pilihan dan menampilkan log debug percakapan SMTP (tanpa kredensial); kegagalan menampilkan pesan galat server.
  15. Outgoing mail dapat dinonaktifkan lewat flag startup (env `MAIL_OUTGOING_DISABLED=true`) atau dari admin UI; saat nonaktif email tetap ditulis ke outbox berstatus "skipped" tanpa dikirim.
  16. Konfigurasi dapat diberikan lewat environment (`SMTP_*`, diutamakan dan mengunci field UI) atau lewat admin UI; perubahan konfigurasi berlaku tanpa restart dan pengiriman memakai retry dari NTF-008.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Transport keluar adalah SMTP di balik interface `Mailer` ([ADR-0003](../adr/0003-transport-email-keluar.md)), dikirim lewat queue ([ADR-0004](../adr/0004-job-latar-dan-queue.md)). Konfigurasi dimuat oleh `internal/core/config/` sesuai pola repo; secret disimpan memakai mekanisme enkripsi [27](27-nfr-arsitektur-platform.md). Inbound email lengkap untuk Service Management (email channel) dibangun di modul 30 di atas komponen ini. Butir inbound (7) aktif bersama NTF-031 di Fase 3; butir outbound wajib di Fase 1.

### NTF-024 Outgoing Mail Settings

- **Deskripsi:** Pengaturan pengiriman email keluar tingkat site.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle enable/disable outgoing email untuk seluruh site; saat nonaktif, banner peringatan tampil di halaman admin notifikasi dan Notification Helper melaporkan alasan ini.
  2. From address default site dapat diatur (default `noreply@<domain instalasi>`); alamat divalidasi formatnya.
  3. From header format dapat diatur dengan variabel (mis. `${fullname} (<nama produk>)`), menghasilkan display name berbasis actor dengan alamat tetap from address.
  4. ⚠ Email subject prefix dapat diatur (default nama produk dalam kurung siku) dan dipakai NTF-018.
  5. Setiap perubahan outgoing mail settings tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman di [22](22-administrasi-sistem.md); server SMTP di NTF-023.

### NTF-025 Project Email Address (Custom Sender)

- **Deskripsi:** Alamat pengirim kustom per project agar email notifikasi tampak berasal dari alamat project.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; ⚠ Admin Proyek (Administer Projects) juga dapat mengubah alamat project di clone, terbatas pada domain terverifikasi.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Notifications > Actions > Edit project email address menampilkan form alamat (company-managed).
  2. Alamat harus berasal dari domain terverifikasi (NTF-026) atau domain default instalasi; alamat lain ditolak dengan pesan galat.
  3. Bila tidak diatur, email project memakai from address site (NTF-024).
  4. Display name mengikuti From header format site.
  5. ⚠ Team-managed project juga menyediakan pengaturan alamat project di halaman Notifications-nya.
  6. ⚠ Pola reply-to-comment: bila project email address disamakan dengan mailbox incoming mail server yang memakai handler "Create a new issue or add a comment to an existing issue", balasan atas email notifikasi menjadi comment; UI menampilkan petunjuk konfigurasi ini dan peringatan bila alamat tidak cocok dengan mailbox mana pun.
- **Team-managed vs company-managed:** Lihat butir 5.
- **Catatan & ketergantungan:** Mail handler di NTF-032; email channel JSM di modul 30 memakai alamat serupa.

### NTF-026 Custom Email Domain & Deliverability

- **Deskripsi:** Konfigurasi organisasi agar email produk dikirim dari domain milik pelanggan dengan otentikasi pengirim yang benar.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs / Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Admin menambahkan domain dan memverifikasi kepemilikan lewat record DNS TXT.
  2. Sistem menampilkan record DNS SPF dan DKIM (kunci DKIM dibuat sistem) yang harus dipasang untuk sending.
  3. Status verifikasi (Pending, Verified, Failed) ditampilkan beserta pesan galat spesifik (record tidak ditemukan, nilai tidak cocok) dan tombol cek ulang.
  4. Bila verifikasi gagal atau dicabut, email otomatis kembali memakai domain default instalasi.
  5. ⚠ Bounce yang diterima untuk domain kustom diproses ke suppression list (NTF-023 butir 9) dan statistik bounce per domain ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pada instalasi yang memakai SMTP relay pelanggan, SPF/DKIM bisa dikelola relay; DKIM signing oleh produk bersifat opsional. Domain verifikasi juga dipakai modul 02 (domain claim) bila ada.

### NTF-027 System & Account Emails

- **Deskripsi:** Email transaksional non-work-item yang dikirim platform, misalnya undangan, akses produk, verifikasi email, dan reset password.
- **Prioritas / Fase / Tier:** P1 / Fase 1 / platform
- **Aktor & izin:** Sistem; dipicu oleh aksi user atau Admin Situs/Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Email undangan user ke site berisi pengundang, tautan bertoken (kedaluwarsa default 7 hari), dan nama produk yang diberi akses.
  2. Email permintaan akses produk dikirim ke admin yang berwenang, dan email persetujuan/penolakan dikirim ke pemohon.
  3. Email verifikasi alamat dan reset password akun berisi tautan bertoken sekali pakai dengan masa berlaku pendek (default 1 jam untuk reset).
  4. Email ini tidak dapat dimatikan lewat preferensi notifikasi personal (NTF-020) dan tetap dikirim walaupun master toggle email notifikasi nonaktif; hanya toggle global outgoing mail yang menghentikannya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alur akun di [01](01-identitas-autentikasi-sesi.md) dan [02](02-pengguna-grup-organisasi.md); modul ini menyediakan template dan pengiriman. Repo saat ini belum punya reset password berbasis email, sehingga ini menjadi konsumen pertama pipeline email.

## Email dari aksi pengguna

### NTF-028 Send Notifications Option on Bulk Change

- **Deskripsi:** Kontrol apakah bulk operation mengirim notifikasi email.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Make Bulk Changes (PERM-011); mematikan notifikasi membutuhkan Admin Situs atau Admin Proyek di semua project terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Langkah konfirmasi bulk edit, bulk transition, bulk move, dan bulk delete menampilkan checkbox "Send notification".
  2. Checkbox default aktif; checkbox hanya dapat dinonaktifkan oleh Admin Situs atau user yang menjadi Admin Proyek di semua project work item terpilih, selain itu checkbox terkunci dengan tooltip penjelasan.
  3. Bila aktif, setiap penerima mendapat paling banyak satu email per work item untuk operasi tersebut, dan email dapat digabung oleh batching (NTF-019).
  4. Operasi bulk Watch / Stop watching tidak mengirim notifikasi apa pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Wizard bulk change di [06](06-work-item-inti.md) ISS-035, [09](09-workflow-status-transisi.md) WF-013, dan [13](13-pencarian-filter-ekspor.md) SRC-031.

### NTF-029 Share Issue / Filter / Dashboard via Email

- **Deskripsi:** Fitur Share untuk mengirim tautan work item, filter, atau dashboard ke user, team, atau alamat email beserta catatan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Setiap user yang dapat melihat item yang dibagikan.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Share tersedia di issue view, halaman filter (navigator), dan dashboard.
  2. Penerima dapat berupa user, team, atau alamat email; alamat non-user dapat memicu undangan ke site hanya bila pengaturan site mengizinkan undangan oleh user, selain itu hanya email tautan yang dikirim.
  3. Pengirim dapat menambahkan catatan opsional (maksimal 1.000 karakter) yang tampil di email.
  4. Tersedia aksi Copy link sebagai alternatif tanpa email.
  5. Penerima tanpa akses tetap menerima email, tetapi membuka tautan menghasilkan halaman "tidak memiliki akses" tanpa membocorkan isi item; email Share tidak memuat isi item selain judul bila penerima tidak berhak melihatnya.
  6. Share tidak tunduk pada notification scheme dan tetap terkirim walaupun scheme project None.
  7. Penerima yang merupakan user juga mendapat notifikasi in-app di tab Direct.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** UI Share di [06](06-work-item-inti.md) ISS-038, [13](13-pencarian-filter-ekspor.md) SRC-048, dan [17](17-dashboard-gadget.md); undangan di [02](02-pengguna-grup-organisasi.md). Rate limit Share mengikuti NTF-009 untuk mencegah spam.

### NTF-030 Filter Subscriptions (Email delivery)

- **Deskripsi:** Pengiriman hasil saved filter via email terjadwal ke user atau group.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pemilik atau user yang dapat melihat filter; subscription untuk group membutuhkan Manage Group Filter Subscriptions (PERM-010); Admin Situs dapat mengelola semua subscription.
- **Aturan perilaku (acceptance criteria):**
  1. User dapat berlangganan untuk diri sendiri; berlangganan untuk group hanya tersedia bagi user dengan global permission Manage group filter subscriptions.
  2. Pilihan jadwal: Daily, Days per week (pilih hari), Days per month (tanggal atau hari ke-N), dan Advanced (cron expression yang divalidasi; ekspresi tidak valid ditolak dengan pesan galat).
  3. Interval dapat diatur once per day atau setiap N jam dalam rentang, dengan waktu kirim di timezone pembuat subscription.
  4. Opsi "Email this filter even if there are no issues found"; bila nonaktif dan hasil kosong, email tidak dikirim.
  5. ⚠ Isi email berupa tabel work item sesuai kolom filter dengan batas 200 baris; bila lebih, email menampilkan jumlah total dan tautan ke navigator.
  6. Hasil filter dievaluasi dengan permission masing-masing penerima (anggota group tidak melihat work item yang tidak boleh ia lihat).
  7. Halaman filter menampilkan daftar subscription dengan aksi edit, delete, dan run now (kirim segera).
  8. Email mengikuti preferensi format HTML/Text masing-masing penerima.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Saved filter dan UI subscription di [13](13-pencarian-filter-ekspor.md) SRC-047; scheduler job di [27](27-nfr-arsitektur-platform.md); permission di [03](03-izin-peran-keamanan.md) PERM-010.

## Email masuk

### NTF-031 Incoming Mail Servers

- **Deskripsi:** Konfigurasi server email masuk agar produk dapat membaca mailbox untuk membuat work item atau comment.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > System > Incoming mail menyediakan aksi Add incoming mail server.
  2. Pilihan provider: Google (OAuth 2.0), Microsoft (OAuth 2.0), dan Other (IMAP/POP).
  3. Untuk IMAP/POP admin mengisi host, port, protocol (IMAP, IMAPS, POP3, POP3S), username, password, dan timeout.
  4. Untuk OAuth admin melakukan authorize akun Google/Microsoft; refresh token disimpan terenkripsi dan status koneksi ditampilkan (Connected/Expired).
  5. Folder IMAP yang dipindai dapat diatur (default INBOX); POP selalu memakai inbox.
  6. Tombol Test connection memverifikasi koneksi dan autentikasi tanpa memproses email, menampilkan pesan sukses atau galat server.
  7. Basic auth untuk Microsoft 365 tidak didukung; memilih Microsoft selalu memakai OAuth.
  8. Admin dapat mengedit dan menghapus server; server yang masih dipakai mail handler tidak dapat dihapus sebelum handler dihapus atau dipindah.
  9. ⚠ Mailbox dipoll secara berkala (default setiap 1 menit, dapat diatur 1-60 menit) oleh job latar belakang.
  10. Pesan yang berhasil diproses dihapus (POP) atau ditandai terbaca/dipindah ke folder processed (IMAP) sesuai protokol, sehingga tidak diproses dua kali.
- **Team-managed vs company-managed:** Tidak ada perbedaan (konfigurasi tingkat site).
- **Catatan & ketergantungan:** Komponen inbound di NTF-023; audit di NTF-011. Email channel JSM (modul 30) memakai konektor yang sama.

### NTF-032 Mail Handlers

- **Deskripsi:** Aturan yang menentukan bagaimana email dari incoming mail server diubah menjadi work item atau comment.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Handler "Create a new issue or add a comment to an existing issue": bila subject memuat issue key yang ada, email menjadi comment; bila tidak, menjadi work item baru.
  2. Handler "Add a comment from the non quoted email body": hanya bagian body yang bukan kutipan balasan yang menjadi comment.
  3. Handler "Add a comment with the entire email body": seluruh body menjadi comment.
  4. Handler "Add a comment before a specified marker or separator in the email body": admin memasukkan regex pemisah; teks sebelum pemisah pertama menjadi comment, regex tidak valid ditolak saat simpan.
  5. Handler "Create a new issue from each email message": setiap email menjadi work item baru tanpa melihat issue key.
  6. Handler pembuat work item mewajibkan project dan issue type default.
  7. Default reporter dipakai bila pengirim tidak cocok dengan user mana pun; tanpa default reporter, email dari pengirim tak dikenal ditolak.
  8. ⚠ Opsi Create users membuat akun untuk pengirim tak dikenal (tunduk pada lisensi dan kebijakan undangan site).
  9. Opsi Notify users mengirim email undangan/akun kepada user yang dibuat oleh handler.
  10. Opsi CC assignee menjadikan penerima CC pertama yang cocok sebagai assignee, dan opsi CC watchers menambahkan semua penerima CC yang cocok sebagai watcher.
  11. Opsi Bulk menentukan perlakuan email bulk/auto-generated: Ignore, Forward, Delete, atau Accept.
  12. Opsi Forward email menentukan alamat tujuan penerusan email yang gagal diproses.
  13. Opsi Strip quotes menghapus teks kutipan balasan dari body sebelum disimpan.
  14. Opsi Catch email address membatasi handler hanya memproses email yang dialamatkan ke alamat tertentu (To/CC).
  15. Tombol Test handler menjalankan handler dalam mode dry-run terhadap email terbaru di mailbox dan menampilkan hasil yang akan terjadi tanpa membuat data.
  16. Admin dapat mengedit dan menghapus handler.
  17. Custom mail handler class tidak didukung; perluasan perilaku dilakukan lewat Automation atau webhook inbound (NTF-023).
- **Team-managed vs company-managed:** Handler dapat menargetkan project kedua jenis; issue type dan field mengikuti konfigurasi project tujuan.
- **Catatan & ketergantungan:** Skenario reply-to-notification-to-comment dicapai dengan handler comment (non quoted atau before marker) yang membaca issue key di subject balasan (NTF-025). App Marketplace untuk email lanjutan berada di luar cakupan. Layanan create work item dipakai bersama [06](06-work-item-inti.md).

### NTF-033 Incoming Email Processing Rules

- **Deskripsi:** Aturan parsing dan pemrosesan email masuk yang berlaku untuk semua mail handler dan alamat email project.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Sistem; permission pengirim dievaluasi sebagai user yang cocok.
- **Aturan perilaku (acceptance criteria):**
  1. Subject menjadi summary (dipotong ke 255 karakter) dan body menjadi description; body HTML dikonversi ke rich text (ADF), plain text dipertahankan.
  2. Lampiran email menjadi attachment work item dan tunduk pada batas ukuran attachment; lampiran melebihi batas dilewati dan dicatat sebagai peringatan.
  3. Inline image (CID) disimpan sebagai attachment dan dirujuk dari rich text.
  4. Pencocokan work item dilakukan berdasarkan issue key pada subject (mis. `KEY-123`, termasuk setelah prefix `Re:`/`Fwd:`).
  5. Balasan atas email notifikasi menjadi comment hanya bila mailbox dan handler dikonfigurasi; balasan ke alamat noreply tidak pernah diproses.
  6. Pengirim dicocokkan ke user berdasarkan alamat email (tidak peka huruf besar/kecil) termasuk alamat sekunder terverifikasi.
  7. Permission pengirim dicek: Create Issues untuk work item baru dan Add Comments untuk comment; tanpa permission email ditolak dan diteruskan sesuai konfigurasi handler.
  8. ⚠ Lampiran pada email yang menjadi comment hanya ditambahkan bila pengirim punya permission Create Attachments; bila tidak, comment tetap dibuat tanpa lampiran.
  9. Auto-reply/out-of-office terdeteksi lewat header dan diabaikan; email yang berasal dari produk ini sendiri (header penanda internal atau alamat pengirim sistem) diabaikan untuk mencegah loop.
  10. Encoding karakter (UTF-8, ISO-8859-x, quoted-printable, base64) dan multipart/alternative ditangani; bagian HTML diutamakan bila tersedia.
  11. Email yang gagal diproses dicatat di log mail handler beserta alasannya dan dapat diteruskan ke alamat forward.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi HTML ke ADF memakai komponen [07](07-editor-konten-kaya.md); attachment di [08](08-komentar-lampiran-aktivitas.md). Mencakup item inventaris "Reply-To and Email Threading" bersama NTF-018.

### NTF-034 Create Work Items via Email (Business projects)

- **Deskripsi:** Alamat email khusus per business project untuk membuat work item dari email tanpa konfigurasi mail server oleh admin.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek mengaktifkan; Pengguna Tim Bisnis mengirim email.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings menyediakan toggle untuk mengaktifkan alamat email project; alamat dibuat otomatis dengan format unik (mis. `<projectkey>-<acak>@<domain inbound>`).
  2. Email masuk menjadi work item dengan issue type default yang dipilih di pengaturan.
  3. Subject menjadi summary, body menjadi description, dan lampiran menjadi attachment (aturan NTF-033).
  4. ⚠ Hanya pengirim yang merupakan user produk dengan permission Create Issues di project tersebut yang diterima; email lain ditolak dan dicatat.
  5. Alamat dapat disalin dari halaman pengaturan dan dibagikan; Admin Proyek dapat membuat ulang alamat (alamat lama berhenti berfungsi).
- **Team-managed vs company-managed:** Tersedia untuk business project kedua jenis.
- **Catatan & ketergantungan:** ⚠ Status fitur ini di produk acuan terpadu 2025-2026 belum pasti. Membutuhkan domain inbound dengan webhook inbound atau catch-all mailbox (NTF-023). Business project di [26](26-kerja-bisnis-formulir-persetujuan.md).

## API & otomasi

### NTF-035 Custom Notification via REST API

- **Deskripsi:** Endpoint untuk mengirim notifikasi email kustom terkait work item dan untuk menekan notifikasi saat update lewat API.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Integrator/Developer API yang terautentikasi dengan Browse Projects pada work item.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /rest/api/3/issue/{issueIdOrKey}/notify` menerima `subject`, `textBody`, dan `htmlBody` (HTML disanitasi); minimal salah satu body wajib diisi, selain itu 400.
  2. Target `to` dapat berisi kombinasi reporter, assignee, watchers, voters, daftar users (accountId), dan daftar groups.
  3. Objek `restrict` membatasi pengiriman hanya ke penerima yang termasuk group atau memiliki permission tertentu.
  4. Hanya penerima yang punya akses ke work item (Browse Projects dan issue security) yang dikirimi; penerima lain dilewati tanpa galat.
  5. Query param `notifyUsers=false` pada edit work item (dan operasi API terkait) menekan notifikasi; hanya diizinkan bagi Admin Situs atau Admin Proyek, selain itu 403.
  6. Rate limit API berlaku untuk endpoint ini (NTF-009) dan respons 429 menyertakan `Retry-After`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi API dan versi path (`/rest/api/3` sebagai lapisan kompatibilitas di samping `/api/v1` internal repo) di [23](23-api-publik-webhook.md). Endpoint wajib dianotasi Swagger dengan `@Security BearerAuth` sesuai pola repo.

### NTF-036 Notification Scheme REST API

- **Deskripsi:** API untuk membaca dan mengelola notification scheme dan event.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Integrator/Developer API; operasi tulis membutuhkan Administer Jira, baca scheme project membutuhkan Administer Projects.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /notificationscheme` mengembalikan daftar terpaginasi (`startAt`, `maxResults`) dengan `expand` (all, field, group, notificationSchemeEvents, projectRole, user), dan `GET /notificationscheme/{id}` mengembalikan satu scheme.
  2. `POST /notificationscheme` membuat scheme (nama unik, 400 bila duplikat) dan `PUT /notificationscheme/{id}` mengubah nama/deskripsi.
  3. `PUT /notificationscheme/{id}/notification` menambah recipient ke satu atau lebih event; duplikat ditolak sesuai NTF-004 butir 9.
  4. `DELETE /notificationscheme/{id}` menghapus scheme (409 bila masih terasosiasi) dan `DELETE /notificationscheme/{id}/notification/{notificationId}` menghapus satu recipient.
  5. `GET /notificationscheme/project` mengembalikan pemetaan project ke scheme, dapat difilter per project atau scheme.
  6. `GET /project/{projectKeyOrId}/notificationscheme` mengembalikan scheme yang dipakai project.
  7. `GET /events` mengembalikan daftar event (ID dan nama) termasuk custom event.
- **Team-managed vs company-managed:** Scheme privat team-managed (NTF-007) tidak dikembalikan oleh endpoint daftar; endpoint project untuk team-managed mengembalikan 404 dengan pesan bahwa project memakai konfigurasi team-managed.
- **Catatan & ketergantungan:** [23](23-api-publik-webhook.md); audit perubahan via API sama dengan UI (NTF-011).

### NTF-037 Notifications from Automation

- **Deskripsi:** Interaksi notifikasi dengan Automation: aksi pengirim email/chat dan kontrol notifikasi pada aksi edit/transition.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek (rule project) dan Admin Situs (rule global).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Send email mendukung to/cc/bcc (alamat, user field, group, role), subject, body HTML atau plain text, dan smart values.
  2. Aksi edit dan transition memiliki opsi "Send notifications" (default aktif); bila dimatikan, event aksi tersebut tidak menghasilkan email/in-app.
  3. Email notifikasi yang dipicu rule menampilkan actor rule (default "Automation for <nama produk>" atau user yang dikonfigurasi sebagai actor).
  4. Aksi Send Slack message dan Send Microsoft Teams message mengirim pesan ke channel/webhook yang dikonfigurasi.
  5. ⚠ Setiap eksekusi rule dibatasi jumlah email (default 100 penerima per eksekusi dan 1.000 email per rule per jam); eksekusi yang melampaui batas gagal dengan status di audit log rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail aksi dan smart value di [20](20-otomasi.md); proteksi email storm di NTF-009.

## Kanal personal eksternal

### NTF-038 Mobile Push Notifications

- **Deskripsi:** Notifikasi push ke aplikasi mobile iOS/Android produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Setiap user yang memasang aplikasi mobile dan login.
- **Aturan perilaku (acceptance criteria):**
  1. Pengaturan push per kategori: mentions, assigned, comments, status changes, dan watching.
  2. Notification schedule / Do not disturb dengan rentang jam dan hari; push di luar jadwal ditahan dan hanya tersedia di inbox.
  3. Badge ikon aplikasi menampilkan jumlah notifikasi belum dibaca.
  4. Tap push membuka work item terkait di aplikasi.
  5. Inbox notifikasi di aplikasi sinkron dengan hub web (status read/unread sama).
  6. Pengaturan push disimpan per site bila user memakai lebih dari satu instalasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk ini belum memiliki aplikasi mobile; fitur ini bergantung keputusan aplikasi mobile/PWA ([25](25-integrasi-impor-ekspor-ekstensi.md)). Web Push di browser dapat menjadi langkah antara (lihat Pertanyaan terbuka). Pengiriman memakai APNs/FCM.

### NTF-039 Slack Personal Notifications

- **Deskripsi:** Notifikasi personal via DM dari aplikasi Slack first-party produk ini.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim (menghubungkan akun), Admin Proyek (notifikasi channel project).
- **Aturan perilaku (acceptance criteria):**
  1. User menghubungkan akun produk ke akun Slack lewat OAuth; koneksi dapat diputus dari personal settings maupun dari Slack.
  2. Toggle notifikasi DM: assigned, mentioned, serta comments/status changes pada work item yang di-watch atau dilaporkan.
  3. Pesan DM menyediakan aksi langsung: comment, assign, transition, dan watch, yang dieksekusi dengan permission user terhubung.
  4. Tautan work item yang ditempel di Slack di-unfurl menjadi kartu ringkas (key, summary, status, assignee) bila user terhubung berhak melihatnya.
  5. Pengaturan tersedia via home tab aplikasi atau perintah `/<produk> notify`.
  6. Notifikasi channel per project/filter tersedia sebagai fitur terkait yang dikonfigurasi Admin Proyek.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Koneksi app chat dan pengaturan workspace di INT-032 ([25](25-integrasi-impor-ekspor-ekstensi.md)); INT-032 adalah pemilik integrasi Slack (aplikasi, OAuth, jenis pesan DM, aksi, unfurl), sedangkan NTF-039 hanya mengatur preferensi dan routing notifikasi ke kanal Slack. Prioritas/tier diselaraskan dengan INT-032. Preferensi kanal pribadi di halaman personal settings dirinci di NAV-019 ([21](21-navigasi-ux-preferensi.md)); preferensi kanal di NTF-020. Pesan DM tunduk pada filter NTF-008.

### NTF-040 Microsoft Teams Personal Notifications

- **Deskripsi:** Notifikasi personal via aplikasi Microsoft Teams first-party produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim.
- **Aturan perilaku (acceptance criteria):**
  1. User menghubungkan akun produk ke akun Microsoft lewat OAuth.
  2. Notifikasi dikirim ke activity feed atau chat bot untuk mention, assigned, comments, dan status change.
  3. Kartu notifikasi (Adaptive Card) menyediakan aksi langsung: comment, transition, dan assign.
  4. User dapat memilih jenis notifikasi personal yang diterima di Teams.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** [25](25-integrasi-impor-ekspor-ekstensi.md); preferensi kanal di NTF-020.

## Fitur Data Center & legacy (rujukan)

### NTF-041 Mail Queue

- **Deskripsi:** UI antrian email keluar dan error queue di Data Center. Tidak dibangun sebagai UI tersendiri; kebutuhannya diserap outbox dan dead-letter queue NTF-008 serta log pengiriman NTF-011.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Jumlah email di antrean dan dead-letter queue ditampilkan sebagai ringkasan di halaman log pengiriman (NTF-011), bukan halaman Mail Queue terpisah.
  2. Flush mail queue manual tidak disediakan; worker memproses antrean terus-menerus.
  3. Item dead-letter dapat dikirim ulang (resend) atau dihapus dari halaman log pengiriman.
  4. Detail galat per email (respons SMTP, jumlah retry) tersedia di log pengiriman.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tidak tersedia di produk acuan Cloud.

### NTF-042 Email Templates Customization

- **Deskripsi:** Kustomisasi template email notifikasi (Velocity) khas Data Center. Tidak dibangun; template disimpan di kode.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Template Velocity per event (HTML dan text) tidak didukung; template produk ini adalah template Go yang dikirim bersama binary.
  2. Kustomisasi header/footer dan CSS email oleh admin tidak didukung (logo site pada header email adalah satu-satunya branding yang mengikuti pengaturan tampilan site).
  3. Template baru untuk custom event tidak didukung; custom event memakai template bawaan (NTF-003).
  4. Perilaku ini konsisten dengan produk acuan Cloud yang tidak mendukung kustomisasi template notifikasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dapat ditinjau ulang bila ada permintaan white-label.

### NTF-043 Send Email to Users (Admin)

- **Deskripsi:** Alat admin Data Center untuk mengirim email massal ke group atau anggota project role. Tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Pemilihan recipient berdasarkan project roles atau groups tidak disediakan sebagai alat admin; kebutuhan serupa dipenuhi aksi Send email di Automation (NTF-037).
  2. Penyusunan subject, body, dan format HTML/text untuk email massal admin tidak disediakan.
  3. Opsi Reply-To dan BCC untuk email massal admin tidak disediakan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengumuman site dapat memakai banner/announcement di [22](22-administrasi-sistem.md).

### NTF-044 User Email Visibility

- **Deskripsi:** Pengaturan Data Center apakah alamat email user terlihat oleh user lain. Clone memakai model profile visibility per akun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs, setiap user.
- **Aturan perilaku (acceptance criteria):**
  1. Mode site-wide Public, Hidden, Masked, dan Logged-in users only tidak dibangun.
  2. Visibilitas email dikontrol pengaturan profile visibility akun (modul 01).
  3. Email notifikasi tidak pernah menampilkan alamat email actor yang disembunyikan; hanya nama dan avatar yang ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Profile visibility di [01](01-identitas-autentikasi-sesi.md).

### NTF-045 Contact Administrators Form

- **Deskripsi:** Form agar user mengirim email ke administrator (Data Center). Tidak dibangun pada fase mana pun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Setiap user; Admin Situs sebagai penerima.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle enable/disable form tidak disediakan.
  2. Pesan kustom pada form tidak disediakan.
  3. Pengiriman email ke admin lewat form tidak disediakan; alur pengganti adalah permintaan akses produk (NTF-027) dan halaman bantuan yang mencantumkan kontak admin.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Ketersediaan di produk acuan Cloud belum pasti.

## Pertanyaan terbuka

- Berapa jendela batching default (10 menit seperti produk acuan, atau lebih pendek) dan apakah user boleh menonaktifkan batching (NTF-019)?
- Apakah notifikasi in-app event juga mengikuti notification scheme project, atau in-app selalu dikirim ke watcher/assignee/reporter terlepas dari scheme (NTF-016)?
- Apakah Digest email notifikasi belum dibaca dibangun, atau cukup mengandalkan Filter Subscriptions (NTF-017)?
- Apakah Admin Proyek boleh mengubah project email address sendiri, atau hanya Admin Situs (NTF-025)?
- Apakah Single Email Address sebagai recipient (penerima non-user) diizinkan, mengingat risiko kebocoran data (NTF-006)?
- Untuk penerimaan email masuk, apakah produk menyediakan domain inbound bawaan (webhook inbound) atau hanya IMAP mailbox milik pelanggan (NTF-023, NTF-034)?
- Apakah aplikasi mobile native dibangun, atau Web Push di browser cukup sebagai pengganti Mobile Push Notifications (NTF-038)?
- Apa nama produk final untuk prefix subject, From header, dan actor automation (NTF-018, NTF-024, NTF-037)?
- Berapa nilai default batas volume (recipient per event, email per jam per penerima, batas email automation) yang dapat diterima pelanggan self-hosted (NTF-009, NTF-037)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk email dan notifikasi lanjutan (mis. Enterprise Mail Handler/JEMH, Notification Assistant, email template editor, SMS gateway) tidak dibangun sebagai fitur native; kebutuhan serupa dipenuhi lewat Automation (modul 20), webhook (modul 23), dan titik ekstensi (modul 25).
- **Fitur Data Center-only / legacy** yang hanya dicatat sebagai rujukan: Mail Queue UI (NTF-041), Email Templates Customization berbasis Velocity (NTF-042), Send Email to Users admin tool (NTF-043), User Email Visibility site-wide (NTF-044), Contact Administrators Form (NTF-045), JNDI mail session, dan custom mail handler class (NTF-032 butir 17).
- **Layanan Atlassian lintas produk**: notifikasi lintas produk Atlassian (Confluence, Bitbucket, Trello, Loom) di hub, Atlassian Administration sebagai penyedia custom email domain, dan aplikasi mobile Atlassian tidak direplikasi; padanan clone dibangun sendiri dalam satu instalasi (NTF-015, NTF-026, NTF-038).
- **Customer notifications dan email channel Service Management** (notifikasi ke customer portal, template customer, keep-on-sending) didefinisikan di [30 JSM Help Center & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md), bukan di modul ini.
- **SMS, panggilan telepon, dan notifikasi on-call** untuk alert operasional berada di modul 34 (JSM Operations).
- **Penyedia email terkelola** (mis. layanan SMTP relay komersial) tidak dioperasikan oleh produk ini; pelanggan menyediakan SMTP/relay sendiri.
