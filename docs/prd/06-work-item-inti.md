# 06 — Work Item Inti: Buat, Lihat, Ubah & Operasi

Modul ini mendefinisikan perilaku runtime work item di produk ini: identitas (key dan ID), cara membuat work item (Create dialog, inline create, URL prefill), tampilan issue view beserta mode buka dan pengeditan inline, semantik field sistem (Assignee, Reporter, Priority, Resolution, Labels, tanggal, Rank, Flag), hierarki dan relasi (Subtask, Parent, child, issue link, web/remote link), operasi siklus hidup (clone, move, change type, convert, delete, archive, bulk change), kolaborasi (watch, vote, share, print/export, issue collector), serta penegakan izin, restriction, guardrail, dan fitur AI opsional pada work item. Konfigurasi yang membentuk work item (work type, field, screen, scheme, nilai global priority/resolution/link type) ada di modul 05; modul ini hanya mengatur cara konfigurasi itu dipakai saat runtime. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga seluruh entitas work item (tabel work item, counter key per project, link, watcher, vote, rank, history) adalah pekerjaan baru yang dirancang bersama model data di modul 28 dan dikemas sebagai feature module baru `internal/features/workitem` mengikuti pola Clean Architecture repo.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission Browse/Create/Edit/Delete/Move/Assign/Link/Schedule, issue security, Make bulk changes.
- [04 Proyek](04-proyek.md) — project key, default assignee, company-managed vs team-managed, archive project.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — work type, hierarki, field sistem/custom, screen, field configuration, issue layout, priority, resolution, link type.
- [07 Editor Teks Kaya](07-editor-konten-kaya.md) — Description/rich text, smart link, mention, deep link comment.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — comment, attachment, tab Activity/History.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — initial status, transition, validator, post function, workflow property.
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — card, swimlane, rank di board/backlog, sprint.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — klausa JQL field work item, navigator, ekspor CSV/Excel.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — Log work, estimasi, rollup.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Fix/Affects versions, Components, component lead.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — bar timeline, Target start/end, Calendar.
- [19 Notifikasi](19-notifikasi-email.md) — event Issue created/updated/assigned/moved/deleted, auto-watch, Share email.
- [20 Otomasi](20-otomasi.md) — Recurring work items, aksi otomasi pada work item.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — global nav, shortcut, recently viewed, preferensi personal.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — General configuration (voting, watching, unassigned), audit log.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST create/edit/link/remote link, webhook, redaction API.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — Slack/Teams, wiki/Confluence, panel apps.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — Forms sebagai titik masuk create.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — realtime, job asinkron, guardrail kinerja.
- [28 Model Data Inti](28-model-data-inti.md) — skema tabel work item dan katalog event domain.

## Persona terkait

- **Developer/Anggota Tim** — membuat, mengedit, menugaskan, menautkan, dan memindahkan work item sehari-hari; memakai shortcut dan inline create.
- **Scrum Master/Team Lead** — mengatur Rank, Flag, Subtask, Parent, dan bulk change untuk menjaga backlog tetap rapi.
- **Product Owner/PM** — menyusun Epic dan child, memprioritaskan dengan Priority/Rank/Votes, membagikan work item ke stakeholder.
- **Admin Proyek/Project Lead** — mengatur default assignee, issue collector, restriction, dan menjalankan move/delete/archive.
- **Admin Situs** — mengatur toggle global (voting, watching, unassigned), priority/resolution, permission bulk change, redaksi data, dan AI.
- **Stakeholder/Viewer** — membuka work item lewat link, watch, vote, dan mencetak/mengekspor.
- **Pengguna Tim Bisnis** — membuat work item sederhana lewat Create dialog dan list view.
- **Integrator/Developer API** — membuat work item lewat REST/URL prefill, mengelola remote link, bergantung pada key dan ID stabil.
- **Tamu/Kolaborator Eksternal** — mengirim feedback lewat issue collector atau melihat work item yang dibagikan sesuai izin.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin membuat work item dari mana saja dengan shortcut `c` dan konteks yang terisi otomatis agar pencatatan pekerjaan tidak memutus fokus saya (ISS-002, ISS-004, ISS-005).
- Sebagai Integrator/Developer API, saya ingin setiap work item punya key yang mudah dibaca dan ID numerik yang stabil agar referensi di commit, branch, dan sistem lain tidak rusak saat work item dipindah (ISS-001, ISS-032).
- Sebagai Admin Proyek/Project Lead, saya ingin field wajib divalidasi dengan cara yang sama di UI, inline create, dan REST API agar kualitas data terjaga di semua kanal (ISS-003, ISS-010).
- Sebagai Developer/Anggota Tim, saya ingin melihat dan mengedit semua informasi work item di satu halaman tanpa membuka dialog terpisah agar pembaruan cepat dilakukan (ISS-007, ISS-008, ISS-009, ISS-010).
- Sebagai Product Owner/PM, saya ingin memecah Epic menjadi child dan Subtask serta melihat progresnya agar saya tahu seberapa dekat sebuah inisiatif dengan selesai (ISS-024, ISS-025, ISS-026).
- Sebagai Scrum Master/Team Lead, saya ingin mengurutkan backlog dengan drag-drop dan menandai item yang terhambat agar tim mengerjakan hal terpenting lebih dulu (ISS-022, ISS-023).
- Sebagai Developer/Anggota Tim, saya ingin menautkan work item dengan relasi blocks/duplicates dan menautkan resource eksternal agar ketergantungan terlihat jelas (ISS-029, ISS-030).
- Sebagai Admin Proyek/Project Lead, saya ingin memindahkan work item ke project atau tipe lain dengan mapping status dan field yang aman agar tidak ada data hilang tanpa peringatan (ISS-027, ISS-028, ISS-032).
- Sebagai Scrum Master/Team Lead, saya ingin mengubah hingga 1.000 work item sekaligus agar pembersihan backlog tidak dilakukan satu per satu (ISS-035).
- Sebagai Admin Situs, saya ingin mengarsipkan work item lama tanpa menghapusnya agar pencarian dan board tetap ringan tetapi data masih bisa dipulihkan (ISS-033, ISS-034).
- Sebagai Stakeholder/Viewer, saya ingin watch dan vote work item serta menerima tautan yang dibagikan agar saya mengikuti perkembangan yang relevan bagi saya (ISS-036, ISS-037, ISS-038).
- Sebagai Admin Proyek/Project Lead, saya ingin membatasi siapa yang bisa melihat work item sensitif agar informasi keamanan atau HR tidak bocor (ISS-041, ISS-042).
- Sebagai Admin Situs, saya ingin meredaksi data sensitif secara permanen dengan jejak audit agar organisasi patuh terhadap kebijakan privasi (ISS-043).
- Sebagai Developer/Anggota Tim, saya ingin bantuan AI opsional untuk meringkas aktivitas dan menyarankan child work item agar saya lebih cepat memahami konteks (ISS-045, ISS-046).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 19 |
| P1 | 16 |
| P2 | 7 |
| P3 | 5 |
| **Total** | **47** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 19 |
| Fase 2 | 15 |
| Fase 3 | 8 |
| Fase 4 | 4 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 membangun tulang punggung work item, yaitu tabel work item dengan ID numerik dan counter key per project, Create dialog dan inline create dengan validasi Summary/required field, issue view full page dan modal dengan inline edit, field sistem inti (Assignee, Reporter default, Priority, Resolution, Labels, Rank), Subtask, Parent, issue link, watchers (entitas, tombol Watch, dan autowatch yang dibutuhkan notifikasi Fase 1), delete, dan penegakan permission di server untuk setiap endpoint. Seluruh perubahan sejak Fase 1 sudah harus menulis entri history dan menerbitkan event domain (modul 28) agar notifikasi, webhook, dan otomasi di fase berikutnya tidak memerlukan migrasi data. Fase 2 menambah kenyamanan dan visibilitas: Create another, context fields panel, quick-add, actions menu, child panel, convert/change type, web/remote link, clone, share, due/start date, flag, realtime, dan guardrail. Fase 3 menambah operasi skala organisasi yang bergantung pada mesin skema lengkap (Move wizard, Bulk Change, archive Premium, issue security, votes, URL prefill, navigasi antar work item, print/export). Fase 4 berisi kemampuan opsional (issue collector legacy, redaksi data, AI dan agent) yang dibungkus feature flag per plan.

## Fitur

## Identitas & pembuatan work item

### ISS-001 Work Item Key & Numbering

- **Deskripsi:** Setiap work item memiliki key unik berformat `PROJECTKEY-N` yang mudah dibaca dan ID numerik internal yang tidak pernah berubah. Key dipakai di URL, referensi teks, branch/commit, dan pencarian.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects melihat dan menyalin key; Admin Proyek (Administer projects) dapat mengganti project key (PRJ-022); sistem membuat nomor urut.
- **Aturan perilaku (acceptance criteria):**
  1. Project key mengikuti aturan PRJ-012: diawali huruf kapital Latin `A-Z`, diikuti huruf kapital atau angka, panjang 2-10 karakter; key work item adalah project key ditambah `-` dan nomor urut desimal mulai dari 1. Pola key kustom tidak didukung (pola konfigurabel hanya ada di Data Center, lihat ISS-047).
  2. Counter nomor urut disimpan per project dan dinaikkan secara atomik di dalam transaksi create; nomor yang sudah pernah dipakai tidak pernah dipakai ulang, termasuk setelah work item dihapus atau dipindah keluar project.
  3. Setiap work item punya ID numerik internal (issue id) yang tetap sama walaupun key berubah karena move atau ganti project key; REST API menerima key maupun ID sebagai pengenal.
  4. URL kanonik adalah `/browse/KEY-123`; board, backlog, dan navigator menerima parameter `?selectedIssue=KEY-123` yang membuka work item tersebut di panel/modal.
  5. Work item yang dipindah ke project lain (ISS-032) mendapat key baru; key lama disimpan sebagai alias sehingga `/browse/OLD-1` me-redirect (HTTP 301 untuk API, redirect klien untuk SPA) ke key baru dan JQL `key = OLD-1` tetap menemukan work item tersebut.
  6. Saat project key diganti (PRJ-022), seluruh key work item di project itu berubah memakai prefix baru dengan nomor yang sama, dan semua key lama tetap me-redirect serta dapat dicari.
  7. Breadcrumb/header issue view menyediakan aksi "Copy key" (menyalin `KEY-123`) dan "Copy link" (menyalin URL absolut `/browse/KEY-123`) dengan flag konfirmasi "Disalin".
  8. Key work item dengan resolution terisi (ISS-019) ditampilkan dengan gaya coret (strikethrough) di link, mention, dan daftar linked work items.
  9. Teks yang cocok dengan pola `[A-Z][A-Z0-9]{1,9}-\d+` dan merujuk work item yang ada serta dapat di-browse pengguna dirender sebagai smart link di editor, comment, commit, branch, dan PR (modul 07, 24); key yang tidak ada atau tidak boleh dilihat tetap tampil sebagai teks biasa.
  10. Quick search dan pencarian key bersifat case-insensitive: `abc-12` menemukan `ABC-12`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada PRJ-012 dan PRJ-022. Model data (modul 28) memerlukan tabel `project_issue_counters` dan tabel alias `work_item_key_history`. Kondisi repo saat ini belum memiliki entitas project maupun work item. Rujukan silang ID: EDT-020 (auto-link key), DM-014 (event domain).

### ISS-002 Global Create Dialog (Create work item)

- **Deskripsi:** Modal create lengkap yang dibuka dari tombol Create di global nav atau shortcut `c`, berisi pemilih project dan work type serta field sesuai create screen (company-managed) atau layout work type (team-managed).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan permission Create Issues pada project tujuan; Admin Proyek mengatur field yang tampil lewat konfigurasi modul 05.
- **Aturan perilaku (acceptance criteria):**
  1. Modal berisi pemilih Project (default: project konteks halaman saat ini; bila tidak ada, project terakhir yang dipakai pengguna) dan pemilih Work type (default: default work type project, CFG-001).
  2. Field dirender dari create screen pada issue type screen scheme (company-managed) atau dari layout field work type (team-managed, CFG-042); mengganti project atau work type memuat ulang daftar field tanpa menutup modal.
  3. Field yang didukung minimal: Summary, Description (editor rich text, modul 07), Assignee (dengan tautan "Assign to me"), Reporter, Priority, Labels, Parent, Sprint, Team, Components, Fix versions, Affects versions, Due date, Start date, Story points, Attachment, Linked work items, Environment, Original estimate, dan Security level; setiap field hanya tampil bila ada di screen/layout.
  4. ⚠ Work item yang dibuat selalu ditempatkan di initial status workflow (modul 09); pemilihan status awal di Create dialog tidak disediakan pada Fase 1 dan menjadi keputusan terbuka.
  5. Tombol Minimize mengecilkan modal menjadi bar di pojok layar; pengguna dapat bernavigasi ke halaman lain lalu memulihkan draft dengan seluruh nilai yang sudah diisi.
  6. Menutup modal (tombol X, Esc, atau klik di luar) saat ada nilai yang berubah menampilkan konfirmasi "Buang work item ini?" dengan pilihan Buang dan Lanjutkan mengedit; tanpa perubahan modal langsung tertutup.
  7. Saat project atau work type diganti, nilai field yang juga ada di project/tipe baru dipertahankan; nilai yang tidak berlaku (mis. opsi select di luar context baru) dikosongkan dan ditandai.
  8. ⚠ Modal menyediakan opsi "Configure fields" yang memungkinkan pengguna memilih field opsional yang ditampilkan di modal (pilihan "Semua field" atau "Field pilihan saya", disimpan per pengguna per project dan work type; field wajib selalu tampil), ditambah tautan ke konfigurasi screen/layout yang hanya tampil bagi Admin Proyek/Admin Situs.
  9. Bila dibuka dari board, sprint, atau Epic, nilai Project, Sprint, dan Parent terisi otomatis sesuai konteks dan tetap dapat diubah.
  10. Ctrl+Enter (Windows/Linux) atau Cmd+Enter (macOS) mengirim form; tombol Create dinonaktifkan selama request berjalan untuk mencegah duplikasi.
  11. Pemilih project hanya menampilkan project tempat pengguna memiliki Create Issues; bila tidak ada satu pun, tombol Create di global nav tetap tampil tetapi modal menampilkan pesan "Anda tidak memiliki izin membuat work item di project mana pun".
  12. ⚠ Create form sederhana (rilis Juli 2026) secara default hanya menampilkan field kunci (minimal Project, Work type, Summary, dan seluruh field required) dengan tautan "Show all fields" yang membuka form penuh tanpa kehilangan nilai yang sudah diisi; field required yang tidak termasuk field kunci tetap ditampilkan dan divalidasi sesuai ISS-003 sehingga form sederhana tidak pernah dapat melewati validasi required.
- **Team-managed vs company-managed:** Company-managed memakai create screen dari screen scheme; team-managed memakai satu layout per work type yang sama untuk create, view, dan edit.
- **Catatan & ketergantungan:** Bergantung pada CFG-008, CFG-042, CFG-046 s.d. CFG-048 (modul 05) dan initial status di modul 09. Endpoint REST `POST /api/v1/issues` di modul 23 memakai validasi yang sama (ISS-003). Shortcut `c` didaftarkan di modul 21.
  - **Pasangan duplikat (pemilik):** [NAV-013 Global Create Dialog](21-navigasi-ux-preferensi.md) di modul 21 membahas fitur yang sama. Pembagian kepemilikan: ISS-002 adalah acuan untuk isi form, field yang dirender, validasi, nilai default/prefill, dan izin pembuatan; NAV-013 adalah acuan untuk entry point global (tombol Create di top bar, shortcut `c`), shell dialog (minimize, full screen), dan flag sukses. Bila kedua spesifikasi berbeda, ISS-002 yang berlaku untuk aturan pembuatan dan NAV-013 untuk perilaku shell; perubahan pada salah satu harus disinkronkan ke yang lain.

### ISS-003 Required Fields & Create Validation

- **Deskripsi:** Aturan validasi saat create dan edit, mencakup field wajib, format, batas panjang, dan validator pada create transition workflow, yang berlaku identik di UI dan API.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna yang membuat/mengedit work item; Admin Situs/Admin Proyek mengatur required field lewat field configuration atau layout.
- **Aturan perilaku (acceptance criteria):**
  1. Summary selalu wajib dan tidak bisa dibuat opsional; field lain wajib bila ditandai required di field configuration (company-managed, CFG-043) atau di layout work type (team-managed, CFG-042).
  2. Field wajib ditandai asterisk (`*`); kegagalan validasi menampilkan pesan galat inline di bawah field terkait (mis. "Summary wajib diisi") dan fokus berpindah ke field pertama yang gagal.
  3. Validator dan condition pada create transition workflow (modul 09) dijalankan saat submit; kegagalan validator menampilkan pesan dari konfigurasi validator dan work item tidak dibuat.
  4. Summary maksimal 255 karakter; field teks multi-baris (Description, Environment, paragraph custom field) maksimal 32.767 karakter; melewati batas menampilkan galat dengan jumlah karakter saat ini dan batasnya.
  5. Validasi tipe: tanggal harus valid sesuai format, angka harus numerik dalam rentang yang didukung, user harus memiliki Assignable User untuk Assignee, dan opsi select harus termasuk opsi aktif di field context yang berlaku.
  6. Bila field required tidak ada di create screen, create gagal dengan pesan "Field <nama> wajib tetapi tidak ada di layar create; hubungi admin untuk menambahkannya ke screen", dan admin melihat tautan ke konfigurasi screen.
  7. Default value dari field context (CFG-036) atau layout terisi otomatis saat form dibuka dan dapat diubah pengguna.
  8. Inline/quick create (ISS-005) yang tidak memenuhi required field selain Summary dialihkan ke Create dialog lengkap dengan nilai yang sudah diketik.
  9. Validasi yang sama dijalankan di server untuk REST API create/edit; respons galat HTTP 400 berisi peta `errors` per field id dan `errorMessages` umum.
- **Team-managed vs company-managed:** Company-managed: required diatur di field configuration dan bisa berbeda per work type lewat scheme; team-managed: required diatur per field di layout work type.
- **Catatan & ketergantungan:** Bergantung pada CFG-036, CFG-042, CFG-043 (modul 05) dan validator workflow (modul 09). Pesan galat mengikuti mekanisme multi-bahasa `constants.GetError` yang sudah ada di repo (EN/ID).

### ISS-004 Create Another & Post-create Feedback

- **Deskripsi:** Opsi untuk membuat beberapa work item berturut-turut dari Create dialog, ditambah umpan balik setelah create berhasil.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Checkbox "Create another" membuat modal tetap terbuka setelah submit; Project, Work type, dan field konteks (Sprint, Parent, Components, Labels, Fix versions) dipertahankan, sedangkan Summary, Description, dan Attachment dikosongkan.
  2. Status checkbox "Create another" disimpan per pengguna dan dipulihkan saat modal dibuka berikutnya.
  3. Setelah create berhasil, flag sukses menampilkan key work item, tautan "View work item", dan aksi "Copy link".
  4. Work item baru langsung tampil di board, backlog, atau list yang sedang terbuka bila memenuhi filter tampilan tersebut, tanpa reload halaman.
  5. Bila work item baru tidak memenuhi filter board/list yang sedang terbuka, flag sukses menambahkan peringatan "Work item ini tidak terlihat di tampilan saat ini" beserta alasannya bila diketahui (mis. tidak masuk filter board).
  6. Pembuat otomatis menjadi watcher bila preferensi personal "Automatically watch work items you create" aktif (default aktif, modul 21).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada ISS-002 dan ISS-036. Penentuan visibilitas di board memakai evaluasi filter board (modul 10, 12).

### ISS-005 Quick Create / Inline Create

- **Deskripsi:** Pembuatan work item langsung di tempat (board, backlog, list, timeline, panel child) cukup dengan mengetik summary.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues pada project terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "+ Create" tersedia di bawah kolom board, di bawah setiap sprint dan backlog, di akhir list view, di timeline, dan di panel child work items/subtasks.
  2. Input inline memiliki pemilih work type dengan ikon; tipe terakhir yang dipakai pengguna di konteks tersebut diingat sebagai default.
  3. Enter membuat work item lalu mengosongkan input yang tetap terbuka untuk item berikutnya; Esc atau klik di luar dengan input kosong membatalkan.
  4. Konteks diwariskan ke work item baru: status sesuai kolom board (status pertama yang dipetakan ke kolom), sprint, parent Epic, dan nilai swimlane (mis. assignee atau Epic) bila relevan.
  5. ⚠ Menempel teks multi-baris menampilkan konfirmasi "Buat N work item?" dan bila disetujui membuat satu work item per baris tidak kosong; dukungan per konteks (board, backlog, list) ditetapkan saat desain.
  6. Bila work type memiliki required field selain Summary, sistem membuka Create dialog lengkap (ISS-002) dengan Summary dan konteks yang sudah terisi.
  7. Rank work item baru mengikuti posisi tempat dibuat (mis. di bawah item terakhir sprint atau di antara dua card), lihat ISS-022.
  8. Tombol "+ Create" disembunyikan bila pengguna tidak memiliki Create Issues pada project tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada ISS-003 dan ISS-022; titik masuk di board/backlog/timeline didefinisikan di modul 10, 11, dan 18.

### ISS-006 Prefilled Create URL & Create Entry Points

- **Deskripsi:** Pembuatan work item lewat URL dengan parameter prefill dan dari titik masuk lain di luar Create dialog utama.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues; Integrator yang membangun tautan create.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ URL `/secure/CreateIssueDetails!init.jspa` dengan parameter `pid`, `issuetype`, `summary`, `description`, `priority`, `assignee`, `reporter`, `labels`, `components`, `fixVersions`, `duedate`, dan `customfield_<id>` membuka Create dialog dengan nilai terisi; produk ini menyediakan juga rute ringkas `/create?pid=...` dengan parameter yang sama.
  2. URL `/secure/CreateIssue.jspa?pid=<id>&issuetype=<id>` membuka Create dialog dengan project dan work type terpilih tanpa field lain.
  3. Parameter memakai ID (pid, issuetype id, priority id, component id, version id), bukan nama; ID yang tidak valid diabaikan dan field dibiarkan default dengan peringatan di modal.
  4. Pembuatan work item dari teks yang di-highlight di halaman wiki/Confluence tersedia lewat integrasi (modul 25) dan membuka Create dialog dengan Summary terisi dan remote link ke halaman sumber.
  5. Create dari integrasi Slack/Microsoft Teams, Forms (modul 26), dan email/mail handler (modul 19, 30) memakai layanan create yang sama sehingga validasi ISS-003 berlaku.
  6. Aksi "Create linked work item" dari issue view membuka Create dialog dengan link terisi (detail di ISS-029).
  7. Shortcut `c` berlaku di semua halaman aplikasi kecuali saat fokus berada di input teks.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rute legacy `.jspa` dipertahankan demi kompatibilitas tautan lama dan di-redirect SPA ke Create dialog (fallback `index.html` di `webui`). Integrasi Slack/Teams/wiki ada di modul 25. Rujukan silang ID: INT-045 (create dari halaman wiki), NAV-039 (shortcut global).

## Issue view & pengeditan

### ISS-007 Issue View (Work item view) Layout

- **Deskripsi:** Halaman detail work item dengan dua kolom: konten utama di kiri dan context fields di kanan; di layar sempit tata letak menjadi satu kolom.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects; elemen edit hanya tampil sesuai permission (ISS-041).
- **Aturan perilaku (acceptance criteria):**
  1. Header berisi breadcrumb (project > parent > key), ikon work type (klik untuk ganti tipe, ISS-028), tombol Watch dengan jumlah watcher, Vote, Share, menu Actions (•••), serta tombol Close dan Expand saat dibuka sebagai modal.
  2. Area utama berurutan: Summary, quick-add buttons (ISS-011), Description, field konten multi-baris, Child work items/Subtasks, Linked work items, Web links, konten wiki/Confluence, Attachments, dan Activity dengan tab All, Comments, History, dan Work log (modul 08, 14).
  3. Panel kanan berisi tombol Status/transition (modul 09), tombol Automation/Actions (modul 20), group Details (Assignee, Reporter, Labels, Parent, Sprint, Priority, Team, dll.), More fields, Development (modul 24), Releases (modul 15), dan panel apps (modul 25).
  4. Footer metadata menampilkan Created, Updated, dan Resolved dalam waktu relatif (mis. "3 jam lalu") dengan tooltip waktu absolut sesuai zona waktu pengguna, serta tautan Configure bagi pengguna yang berwenang mengatur layout.
  5. Setiap section dan group dapat di-collapse/expand, dan statusnya disimpan per pengguna (per section) lalu dipulihkan saat work item lain dibuka.
  6. Field kosong tampil sebagai "None" (field pilihan/user) atau placeholder "Add ..." (field teks) dan dapat diklik untuk mengisi bila pengguna punya Edit Issues.
  7. Banner kontekstual tampil di atas Summary untuk kondisi flagged (ISS-023), archived (ISS-034), dan restricted (ISS-042); beberapa banner dapat tampil bersamaan dengan urutan archived > restricted > flagged.
  8. Status lozenge berwarna sesuai status category: To Do abu-abu, In Progress biru, Done hijau.
  9. ⚠ Lebar kolom kanan dapat di-resize dengan drag dalam batas minimum/maksimum, dan lebar terakhir diingat per pengguna.
  10. Pada lebar layar di bawah breakpoint tablet, panel kanan dipindah ke bawah Summary sehingga tata letak menjadi satu kolom tanpa scroll horizontal.
- **Team-managed vs company-managed:** Company-managed mengikuti issue layout per screen (CFG-049); team-managed mengikuti layout work type (CFG-042). Struktur halaman sama.
- **Catatan & ketergantungan:** Implementasi frontend di `frontend/` (SvelteKit SPA). Data issue view sebaiknya disajikan satu endpoint agregat (`GET /api/v1/issues/{key}?expand=...`) agar render pertama cepat (modul 27).

### ISS-008 Issue View Modes (full page, modal, sidebar)

- **Deskripsi:** Work item dapat dibuka sebagai halaman penuh, modal, atau detail panel/sidebar dari board, backlog, list, dan navigator.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Full page tersedia di `/browse/KEY` dan menampilkan navigasi project di samping.
  2. Modal tampil di atas board/list dengan tombol Expand yang membuka full page untuk work item yang sama.
  3. Detail panel tersedia di backlog dan di navigator mode "Detail view" (daftar di kiri, detail di kanan); memilih item lain di daftar mengganti isi panel tanpa reload.
  4. ⚠ Preferensi pengguna "Buka work item dari board di: Sidebar / Modal" tersedia di menu tampilan board dan disimpan per pengguna.
  5. "Open in new tab" di menu konteks card/baris dan middle-click/Ctrl+klik pada key membuka `/browse/KEY` di tab baru.
  6. Tombol Close dan tombol Esc menutup modal/sidebar dan mengembalikan fokus ke card/baris asal; Esc tidak menutup modal bila sedang ada field dalam mode edit (Esc membatalkan edit lebih dulu).
  7. Membuka modal/sidebar memperbarui URL dengan `?selectedIssue=KEY` sehingga URL dapat dibagikan dan tombol Back browser menutup modal.
  8. ⚠ Lebar sidebar dapat diubah dengan drag, dan lebar terakhir disimpan per pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada ISS-001 (URL kanonik) dan ISS-007. Mode panel di navigator didefinisikan di modul 13.

### ISS-009 Context Fields Panel (Details, Pinned fields, More fields)

- **Deskripsi:** Pengelompokan field di panel kanan issue view; pengguna dapat mem-pin field, dan field tertentu disembunyikan saat kosong.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects (pin bersifat personal); Admin Proyek mengatur susunan lewat issue layout.
- **Aturan perilaku (acceptance criteria):**
  1. Group "Details" berisi field utama sesuai issue layout (context fields).
  2. ⚠ "Pinned fields": pengguna mem-pin field lewat ikon pin sehingga field muncul di bagian atas panel; pin disimpan per pengguna per project, dan unpin mengembalikan field ke posisi semula.
  3. Group "More fields" berisi field berstatus "Hide when empty" yang hanya tampil di Details saat bernilai, atau saat group di-expand.
  4. Panel "Find your field" menerima nama field dan menjelaskan alasan field tidak tampil: tidak ada di screen/layout, context tidak berlaku untuk project/tipe ini, atau field disembunyikan di field configuration, beserta tautan perbaikan untuk admin (CFG-051).
  5. Urutan field di setiap group mengikuti konfigurasi issue layout (CFG-049) atau layout work type (CFG-042).
  6. State collapse group Details dan More fields diingat per pengguna.
  7. Field read-only (mis. Creator, Created, field yang dikunci workflow) tampil tanpa affordance edit (tanpa hover highlight dan ikon pensil).
- **Team-managed vs company-managed:** Company-managed memakai issue layout per screen; team-managed memakai section Context fields dan Hide when empty di layout work type.
- **Catatan & ketergantungan:** Bergantung pada CFG-042, CFG-049, CFG-051 (modul 05). Preferensi pin disimpan di tabel preferensi pengguna (modul 21).

### ISS-010 Inline Edit

- **Deskripsi:** Field di issue view diedit langsung di tempat tanpa dialog Edit terpisah, dengan validasi dan permission per field.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues; field tertentu membutuhkan permission tambahan (lihat butir 6).
- **Aturan perilaku (acceptance criteria):**
  1. Klik field masuk ke mode edit; Enter (field satu baris) atau ikon centang menyimpan, Esc atau ikon X membatalkan dan mengembalikan nilai lama.
  2. ⚠ Klik di luar field menyimpan perubahan untuk field teks satu baris dan field pilihan, sedangkan editor rich text (Description) mempertahankan draft dan meminta pengguna menekan Save atau Cancel.
  3. Editor sesuai tipe field: text, rich text, select/multi-select, user picker, date/datetime picker, number, labels, cascading select, checkbox/radio, URL, serta version/component picker.
  4. Validasi inline (required, format, panjang) dijalankan sebelum request; bila gagal, pesan galat tampil di bawah field dan nilai tidak disimpan; bila gagal di server, nilai dikembalikan ke nilai terakhir yang tersimpan.
  5. Pembaruan bersifat optimistic: nilai baru tampil segera, lalu bila server menolak, nilai dikembalikan dan error flag "Perubahan tidak tersimpan: <alasan>" ditampilkan.
  6. Permission per aksi: Edit Issues untuk field umum, ditambah Assign Issues untuk Assignee, Modify Reporter untuk Reporter, dan Schedule Issues untuk Due date dan Rank; field yang tidak diizinkan tampil read-only.
  7. Status dan Resolution tidak dapat diedit inline; keduanya hanya berubah lewat transition (modul 09, ISS-019).
  8. Summary tidak boleh kosong atau hanya spasi dan maksimal 255 karakter; pelanggaran menampilkan galat dan menolak penyimpanan.
  9. Work item pada status dengan workflow property `jira.issue.editable=false` tidak dapat diedit; semua field tampil read-only dengan tooltip "Work item tidak dapat diedit pada status ini" (PERM-037).
  10. ⚠ Shortcut di issue view: `a` membuka picker Assignee, `i` assign ke diri sendiri, `l` membuka Labels, `m` membuka komentar baru; shortcut `e` untuk masuk mode edit Summary disediakan sebagai tambahan produk ini.
  11. Setiap perubahan yang tersimpan dicatat di History (nilai lama, nilai baru, aktor, waktu) dan menerbitkan event domain `issue.updated` untuk notifikasi, webhook, dan otomasi.
  12. Bila pengguna lain mengubah field yang sama, nilai terbaru dari server menggantikan tampilan saat event realtime diterima, kecuali field tersebut sedang dalam mode edit oleh pengguna ini (lihat ISS-014).
- **Team-managed vs company-managed:** Tidak ada perbedaan dalam interaksi; sumber daftar field berbeda (screen edit vs layout work type).
- **Catatan & ketergantungan:** Endpoint `PATCH /api/v1/issues/{key}` memakai validasi ISS-003 dan permission ISS-041. History ada di modul 08, event di modul 28. Rujukan silang ID: NAV-041 (katalog shortcut issue view), DM-014.

### ISS-011 Quick-add Buttons

- **Deskripsi:** Baris tombol di bawah Summary untuk menambahkan konten terkait dengan cepat.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan permission sesuai aksi (Create Attachments, Create Issues, Link Issues).
- **Aturan perilaku (acceptance criteria):**
  1. "Add attachment" membuka pemilih file dan juga menerima drag-drop file ke issue view (modul 08).
  2. "Add child work item" atau "Create subtask" tampil sesuai level hierarki work item (Epic menampilkan child, Standard menampilkan subtask, Subtask tidak menampilkan keduanya).
  3. "Link work item" membuka form link (ISS-029).
  4. "Add web link" dan "Link Confluence page" membuka form web link/halaman wiki (ISS-030).
  5. Tombol tambahan dari ekstensi (modul 25) dapat didaftarkan ke baris ini tanpa mengubah tombol bawaan.
  6. Tombol disembunyikan bila fiturnya nonaktif secara global (mis. linking dimatikan di CFG-056 atau subtask dimatikan di CFG-005) atau pengguna tidak memiliki permission terkait.
  7. Tombol yang tidak muat di lebar area utama dipindahkan ke menu overflow (•••) dengan urutan yang sama.
- **Team-managed vs company-managed:** Di team-managed, ketersediaan subtask/child mengikuti Features toggle project (PRJ-049).
- **Catatan & ketergantungan:** Bergantung pada ISS-024, ISS-026, ISS-029, ISS-030.

### ISS-012 Work Item Actions Menu (•••)

- **Deskripsi:** Menu aksi di header issue view untuk operasi tingkat work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan permission sesuai aksi; Admin Proyek untuk Configure.
- **Aturan perilaku (acceptance criteria):**
  1. Menu berisi "Add flag" atau "Remove flag" sesuai status flag (ISS-023).
  2. Menu berisi Log work (modul 14), Clone (ISS-031), Move (ISS-032), Archive (ISS-034, hanya plan Premium/Enterprise), dan Delete (ISS-033).
  3. Menu berisi "Convert to subtask" atau "Convert to work item" (ISS-027) dan "Change parent" (ISS-025).
  4. Menu berisi Print, Export XML, dan Export Word (ISS-039).
  5. Menu berisi "Manage watchers" dan "View voters" (ISS-036, ISS-037).
  6. Menu berisi "Configure" (membuka issue layout/layout work type) dan "Find your field" (ISS-009).
  7. ⚠ Menu berisi "Set up recurrence" untuk membuat jadwal Recurring Work Items (modul 20); bila fitur recurrence tidak aktif, item tidak tampil.
  8. Item menu hanya muncul sesuai permission dan konteks: mis. Convert to subtask tidak tampil pada Subtask atau Epic yang memiliki child, Archive tidak tampil pada plan Free/Standard, Delete tidak tampil tanpa Delete Issues.
  9. Shortcut `.` membuka command palette berisi seluruh operasi work item yang tersedia untuk pengguna, dengan pencarian teks.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Daftar operasi tersedia dihitung di server (`GET /api/v1/issues/{key}/operations`) agar konsisten dengan ISS-041. Rujukan silang ID: AUTO-072 (Recurring work items), NAV-041.

### ISS-013 Navigation between Work Items

- **Deskripsi:** Navigasi antar work item dari konteks pencarian atau board.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Saat work item dibuka dari hasil pencarian, issue view menampilkan tombol Next/Previous beserta posisi ("3 dari 42"); shortcut `j`/`k` berpindah ke item berikutnya/sebelumnya.
  2. Tautan "Back to search" atau "Back to board" mengembalikan pengguna ke daftar asal dengan filter dan posisi scroll yang sama.
  3. Work item yang dibuka tercatat di Recently viewed dan tampil di quick search serta halaman Your work (modul 21).
  4. Breadcrumb memuat tautan ke project dan parent (bila ada).
  5. Klik key work item di mana pun (card, list, link, mention) membuka work item tersebut sesuai mode buka konteksnya (ISS-008).
  6. Deep link ke comment tertentu memakai `?focusedCommentId=<id>` yang men-scroll dan menyorot comment tersebut (modul 07, 08).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konteks daftar asal disimpan di state klien (session) dan tidak membutuhkan endpoint baru.

### ISS-014 Real-time Updates & Concurrent Editing

- **Deskripsi:** Perilaku issue view saat work item diubah pengguna lain atau proses lain secara bersamaan.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects; sistem (otomasi, integrasi).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Issue view berlangganan event perubahan work item lewat kanal realtime (modul 27); perubahan field oleh pihak lain diterapkan otomatis, dan bila penerapan otomatis tidak memungkinkan ditampilkan penanda "Work item ini telah diperbarui" dengan tombol Refresh.
  2. ⚠ Comment baru dari pengguna lain muncul di tab Activity tanpa reload halaman.
  3. Konflik edit diselesaikan dengan aturan "simpanan terakhir menang" per field; perubahan lokal yang belum disimpan tetap ada di editor, dan pengguna diberi tahu bila nilai server berubah selama ia mengedit.
  4. Perubahan dari otomasi tercermin di issue view setelah eksekusi rule selesai, dengan aktor otomasi tercatat di History.
  5. Board dan list yang sedang terbuka ikut ter-update (posisi card, status, assignee) saat menerima event yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butuh infrastruktur push (WebSocket/SSE) di modul 27. Fase 1 cukup memuat ulang data setelah setiap aksi pengguna sendiri; realtime ditambahkan di Fase 2. Rujukan silang ID: NFR-009 (SSE/WebSocket), NAV-046 (live updates).

## Field sistem & atribut work item

### ISS-015 Core System Fields on Work Item

- **Deskripsi:** Field sistem yang dimiliki setiap work item beserta semantik dan batasannya saat runtime; katalog konfigurasinya ada di CFG-011.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects melihat; pengubahan mengikuti ISS-010 dan ISS-041; sistem mengisi field otomatis.
- **Aturan perilaku (acceptance criteria):**
  1. Summary wajib dengan maksimal 255 karakter; Description bertipe rich text (dokumen terstruktur setara ADF, modul 07) dengan maksimal 32.767 karakter dalam representasi teks.
  2. Setiap work item memiliki Work type, Status, Resolution, Priority, Assignee, Reporter, dan Creator; Creator diisi otomatis dengan pembuat, read-only, dan dapat berbeda dari Reporter.
  3. Field klasifikasi tersedia: Labels, Components, Affects versions, Fix versions, dan Environment.
  4. Field tanggal sistem: Due date, Created, Updated, Resolved, Last viewed (per pengguna, tidak global), dan Status category changed (diisi saat status category berubah).
  5. Field bawaan agile/perencanaan: Parent, Sprint, Story points / Story point estimate, Team, dan Start date.
  6. Field relasi dan keamanan: Security level, Watchers, Votes, Attachment, dan Linked work items.
  7. Field time tracking Original estimate, Remaining estimate, dan Time spent dikelola modul 14, dengan nilai agregat (Σ) yang menggulung nilai dari subtask.
  8. ⚠ Locked field yang dikelola sistem (minimal Sprint, Rank, Flagged, Story point estimate, Parent, Team) tidak dapat dihapus atau di-rename, dan konfigurasinya dibatasi sesuai CFG-012.
  9. Field kosong tampil sebagai "None" dan dapat dicari dengan `EMPTY` atau `IS NULL` di JQL (modul 12).
  10. Field Updated diperbarui pada setiap edit field, comment baru/diubah, transition, perubahan link, dan perubahan attachment; Last viewed tidak memengaruhi Updated.
- **Team-managed vs company-managed:** Field sistem sama; di team-managed beberapa field (Sprint, Story point estimate, Components) hanya tampil bila fitur project terkait aktif (PRJ-049).
- **Catatan & ketergantungan:** Definisi konfigurasi di CFG-011 s.d. CFG-018; tabel `work_items` di modul 28 menyimpan field sistem sebagai kolom, custom field di tabel nilai terpisah.

### ISS-016 Assignee Rules

- **Deskripsi:** Aturan penugasan work item: siapa yang dapat ditugaskan, default assignee, dan perilaku unassigned.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Assign Issues mengubah assignee; hanya pengguna dengan Assignable User yang dapat dipilih; Admin Proyek mengatur default assignee (PRJ-021); Admin Situs mengatur toggle unassigned.
- **Aturan perilaku (acceptance criteria):**
  1. Picker Assignee hanya menampilkan pengguna yang memiliki Assignable User pada project; mengubah assignee memerlukan Assign Issues (PERM-022), dan permintaan API dengan pengguna non-assignable ditolak HTTP 400 "User tidak dapat ditugaskan".
  2. Tautan "Assign to me" tersedia di field Assignee dan shortcut `i`; shortcut `a` membuka picker Assignee.
  3. Opsi "Automatic" saat create (company-managed) memakai component lead bila work item memiliki component dengan default assignee lead, selain itu default assignee project (Project lead atau Unassigned, PRJ-021).
  4. Setting global "Allow unassigned work items" (CFG-057); bila OFF, Assignee wajib diisi saat create dan tidak bisa dikosongkan, dan default assignee project tidak boleh Unassigned.
  5. Picker menampilkan saran (pengguna terakhir dipilih, pengguna yang sering ditugaskan di project), pencarian berdasarkan nama atau email (sesuai visibilitas profil AUTH-002), dan avatar.
  6. Hanya satu Assignee per work item; kebutuhan multi-assignee dipenuhi lewat custom field multi-user picker (CFG-030).
  7. Pengguna yang dinonaktifkan atau dihapus tetap tersimpan sebagai assignee dan tampil sebagai "Former user" (dihapus) atau nama dengan label nonaktif (USR-008), serta tidak muncul di picker.
  8. Di team-managed, semua pengguna dengan akses ke project (PRJ-038) dapat ditugaskan.
  9. Perubahan assignee menerbitkan event `issue.assigned` dan tercatat di History.
  10. Post function workflow (modul 09) dan aksi otomasi (modul 20) dapat mengubah assignee otomatis, dengan aktor sistem tercatat di History.
  11. ⚠ Agent AI (ISS-046) dapat muncul di picker Assignee dengan penanda AI bila fitur agent aktif.
  12. JQL mendukung `assignee = currentUser()`, `assignee IS EMPTY`, dan `assignee WAS <user>` (modul 12).
- **Team-managed vs company-managed:** Company-managed: Assignable User dari permission scheme dan opsi Automatic; team-managed: siapa pun yang punya akses project, tanpa opsi Automatic berbasis component.
- **Catatan & ketergantungan:** Bergantung pada PERM-022, PRJ-021, PRJ-038, USR-008, CFG-057.

### ISS-017 Reporter Rules

- **Deskripsi:** Aturan field Reporter: default-nya pembuat work item, hanya dapat diubah dengan permission khusus, dan berbeda dari Creator.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Modify Reporter mengubah reporter; Admin Proyek mengatur screen.
- **Aturan perilaku (acceptance criteria):**
  1. Default Reporter adalah pengguna yang membuat work item (termasuk lewat API dengan kredensial pengguna tersebut).
  2. Mengubah Reporter saat create maupun edit memerlukan permission Modify Reporter (PERM-024); tanpa izin, field tampil read-only dan nilai Reporter dari API diabaikan atau ditolak HTTP 403.
  3. Creator adalah field read-only yang tidak berubah walaupun Reporter diganti.
  4. Reporter dari issue collector (ISS-040) atau email/mail handler adalah default user yang dikonfigurasi, atau pengguna yang cocok dengan alamat email pengirim bila ditemukan.
  5. Reporter tidak dapat memberi vote pada work item yang ia laporkan (ISS-037).
  6. Reporter menerima notifikasi sesuai notification scheme (peran "Reporter", modul 19).
  7. Field Reporter dapat dibuat wajib atau disembunyikan dari screen (company-managed); bila disembunyikan, nilai default tetap diisi otomatis.
- **Team-managed vs company-managed:** Di team-managed, Reporter dapat diubah oleh peran Administrator dan Member; konfigurasi screen tidak tersedia.
- **Catatan & ketergantungan:** Fase 1 sudah mengisi Reporter = pembuat secara otomatis; pengubahan Reporter dan aturan permission ditambahkan di Fase 2.

### ISS-018 Priority

- **Deskripsi:** Field sistem untuk tingkat kepentingan work item dengan nilai global, ikon, warna, dan urutan; administrasi nilainya ada di CFG-053 dan CFG-054.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs mengelola nilai priority; pengguna dengan Edit Issues mengubah priority work item.
- **Aturan perilaku (acceptance criteria):**
  1. Nilai default untuk site baru: Highest, High, Medium, Low, Lowest; data hasil impor dapat membawa skema lama Blocker, Critical, Major, Minor, Trivial dan dipetakan saat impor (modul 25).
  2. Setiap priority memiliki nama (unik), deskripsi, ikon, dan warna.
  3. Default priority (default: Medium) dipakai saat create bila pengguna tidak memilih atau field tidak ada di screen.
  4. Admin dapat menambah, mengedit, mengurutkan ulang, dan menghapus priority; penghapusan priority yang dipakai mewajibkan pemilihan priority pengganti dan memindahkan semua work item secara asinkron sebelum nilai dihapus.
  5. Urutan priority menentukan hasil sort `ORDER BY priority` (urutan admin, bukan alfabet).
  6. Ikon priority tampil di card board, list/backlog, hasil pencarian, dan issue view.
  7. Team-managed memakai priority global dari default priority scheme dan tidak dapat mengkustomnya per project.
  8. Field Priority dapat disembunyikan atau dibuat tidak wajib lewat field configuration; bila disembunyikan, default priority tetap diisi.
- **Team-managed vs company-managed:** Company-managed dapat memakai priority scheme berbeda per project (CFG-054); team-managed selalu memakai default priority scheme.
- **Catatan & ketergantungan:** Seed priority default dibuat di migrasi awal modul 28.

### ISS-019 Resolution

- **Deskripsi:** Field yang menandai apakah work item dianggap selesai; nilai kosong berarti Unresolved, dan field ini menjadi dasar semantik "done" di laporan dan JQL. Administrasi nilainya ada di CFG-055.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs mengelola nilai; pengguna dengan Transition Issues/Resolve Issues mengisinya lewat transition; sistem (post function) mengisi/mengosongkan.
- **Aturan perilaku (acceptance criteria):**
  1. Nilai default untuk site baru: Done, Won't Do, Duplicate, Cannot Reproduce; Admin Situs dapat menambah, mengedit, menghapus (dengan pengganti), mengurutkan, dan menentukan nilai default.
  2. Unresolved berarti field kosong (NULL); produk ini menolak pembuatan nilai resolution bernama "Unresolved" (case-insensitive) dengan galat yang menjelaskan bahwa nilai tersebut akan membuat work item terhitung resolved.
  3. Company-managed: resolution diisi lewat transition screen atau post function, dan transition yang membuka kembali work item (reopen) harus mengosongkannya lewat post function "Clear resolution" (modul 09).
  4. ⚠ Team-managed: resolution otomatis di-set ke nilai default (Done) saat work item masuk status berkategori Done dan dikosongkan saat keluar dari kategori Done.
  5. Resolved date diisi saat resolution di-set dan dikosongkan saat resolution dihapus.
  6. Gaya coret key (ISS-001), filter `resolution = Unresolved`, dan perhitungan gadget/report (modul 16, 17) bergantung pada field ini, bukan pada status.
  7. Nilai yang boleh dipilih pada transition dapat dibatasi lewat workflow property `jira.field.resolution.include` / `jira.field.resolution.exclude` (PERM-037, modul 09).
  8. Resolution tidak dapat diedit inline atau lewat Bulk Edit; perubahan hanya lewat transition atau Bulk Transition (ISS-035).
  9. Vote tidak dapat diberikan pada work item yang sudah resolved (ISS-037).
- **Team-managed vs company-managed:** Company-managed dikendalikan transition screen/post function; team-managed otomatis berdasarkan status category.
- **Catatan & ketergantungan:** Bergantung pada modul 09 (transition) dan CFG-055.

### ISS-020 Labels

- **Deskripsi:** Tag bebas yang dipakai bersama secara global untuk mengkategorikan work item.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Field Labels menampilkan autocomplete dari label yang sudah ada; mengetik label baru lalu Enter membuat label tersebut.
  2. Label tidak boleh mengandung spasi (spasi saat mengetik memisahkan label) dan panjang maksimal 255 karakter; pelanggaran ditolak dengan pesan galat.
  3. ⚠ Penyimpanan dan autocomplete mempertahankan huruf asli label, sedangkan pencarian JQL `labels = X` tidak membedakan huruf besar/kecil.
  4. Label bersifat global lintas project: label yang dibuat di satu project muncul di autocomplete project lain.
  5. Klik label membuka pencarian `labels = "<label>"` di navigator.
  6. Tidak ada UI admin untuk rename/delete label global; label hilang dari autocomplete saat tidak lagi dipakai work item mana pun, dan rename dilakukan lewat Bulk Edit (ISS-035).
  7. Custom field bertipe Labels (CFG-031) memiliki namespace label sendiri yang terpisah dari field Labels sistem.
  8. Bulk add, remove, dan replace labels tersedia lewat Bulk Edit.
  9. Label tampil di card board bila dikonfigurasi di card layout (modul 10).
  10. JQL mendukung `labels`, `labels IN (...)`, dan `labels IS EMPTY`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan sebagai relasi work item ke tabel label global dengan kolom nilai asli dan nilai ternormalisasi (lowercase) untuk pencarian.

### ISS-021 Due Date & Start Date

- **Deskripsi:** Field tanggal untuk penjadwalan: Due date (sistem) dan Start date (field bawaan), dipakai Timeline, Calendar, dan indikator overdue.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues; Due date juga memerlukan Schedule Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Date picker menampilkan dan menerima format tanggal sesuai setting site dan locale pengguna (modul 21, 22).
  2. Mengubah Due date memerlukan Schedule Issues (PERM-024); tanpa izin, field tampil read-only.
  3. Card dan issue view menampilkan indikator overdue (merah, due date sebelum hari ini dan belum resolved) dan due soon (kuning, dalam ambang default 3 hari).
  4. Timeline dan Calendar (modul 18) menggambar bar dari Start date ke Due date; drag atau resize bar mengubah kedua tanggal dan tercatat di History.
  5. ⚠ Start date setelah Due date ditampilkan sebagai peringatan pada field dan bar timeline, tetapi tidak memblokir penyimpanan.
  6. JQL mendukung `due`, `duedate < now()`, `duedate <= endOfWeek()`, dan `"Start date[Date]"` (modul 12).
  7. Plans memakai Target start/Target end sebagai field terpisah (CFG-022, modul 18) yang tidak tersinkron otomatis dengan Start/Due date.
  8. Due date disimpan sebagai tanggal tanpa zona waktu (DATE), sehingga nilainya sama bagi semua pengguna di zona waktu mana pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan; di team-managed Start date tersedia di field library.
- **Catatan & ketergantungan:** Bergantung pada modul 18 dan CFG-022.

### ISS-022 Rank (LexoRank)

- **Deskripsi:** Field Rank global berbasis LexoRank yang menentukan urutan work item di backlog, board, dan list.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Schedule Issues dan Edit Issues; sistem menjalankan rebalancing.
- **Aturan perilaku (acceptance criteria):**
  1. Rank disimpan sebagai string leksikografis berformat `bucket|value` (mis. `0|i0000:`); menyisipkan item di antara dua item hanya menghitung nilai tengah tanpa memperbarui item lain.
  2. Terdapat satu urutan Rank global lintas project dan board; board memakai urutan ini bila filternya diakhiri `ORDER BY Rank`.
  3. Drag-drop di backlog, board, dan list mengubah rank; tersedia juga aksi "Move to top", "Move to bottom", "Move up", dan "Move down" di menu konteks.
  4. Memilih banyak item (Shift/Ctrl+klik) lalu men-drag mengubah rank semua item terpilih dan mempertahankan urutan relatifnya.
  5. Mengubah rank memerlukan Schedule Issues dan Edit Issues; tanpa izin, drag tidak tersedia dan aksi Move to top/bottom disembunyikan.
  6. Rebalancing otomatis dijalankan di latar belakang saat panjang nilai melewati ambang, dengan rotasi tiga bucket (0, 1, 2) sehingga urutan tidak berubah dan pengguna tidak terganggu.
  7. Work item baru ditempatkan di bawah urutan global, atau di posisi tempat dibuat bila dibuat inline (ISS-005).
  8. Subtask diurutkan di dalam parent-nya; drag subtask hanya mengubah urutan di antara sibling.
  9. History mencatat entri "Ranked higher" atau "Ranked lower".
  10. Board yang filternya tidak memakai `ORDER BY Rank` tidak mengizinkan drag reorder dan menampilkan peringatan yang menjelaskan alasannya.
  11. Rank adalah locked field yang tidak dapat dihapus atau dikonfigurasi ulang (CFG-012, CFG-015).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Implementasi LexoRank perlu operasi atomik dan indeks B-tree pada kolom rank (collation `C`) di PostgreSQL (modul 27, 28). Rujukan silang ID: BRD-017, SPR-003, PLAN-006, API-033, DM-010.

### ISS-023 Flag (Impediment)

- **Deskripsi:** Penanda bahwa work item terhambat, disimpan di field Flagged dengan nilai Impediment dan disorot di board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. "Add flag" tersedia dari menu ••• di issue view, menu klik kanan card di board/backlog, dan command palette `.`.
  2. Dialog flag menawarkan opsi menambahkan comment saat flag dipasang atau dilepas; comment diawali penanda "Flag added" / "Flag removed".
  3. Card yang di-flag diberi latar sorotan dan ikon bendera, dan issue view menampilkan banner "Flagged".
  4. "Remove flag" melepas flag dari lokasi yang sama.
  5. JQL mendukung `flagged = Impediment` dan `flagged IS NOT EMPTY`.
  6. Memasang dan melepas flag memerlukan Edit Issues.
  7. Flagged adalah locked custom field bertipe checkbox dengan satu opsi "Impediment" (CFG-018).
  8. Pemasangan dan pelepasan flag tercatat di History.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CFG-018 dan tampilan card di modul 10.
  - **Pasangan duplikat (pemilik):** [BRD-021 Flagging Issues](10-board-kolom-kartu.md) di modul 10 membahas fitur yang sama. Pembagian kepemilikan: ISS-023 adalah acuan untuk field Flagged (model data, nilai Impediment), izin Edit Issues, operator JQL, comment penanda, dan pencatatan History; BRD-021 adalah acuan untuk tampilan dan aksi di board/backlog (sorotan kartu/baris, bulk flag/unflag, insight "Issues for attention", pemakaian di quick filter/swimlane) serta integrasi otomasi. Bila kedua spesifikasi berbeda, pemilik bagian yang bersangkutan yang berlaku.

## Hierarki & relasi

### ISS-024 Subtasks

- **Deskripsi:** Work item level -1 untuk memecah pekerjaan di bawah parent; subtask berbagi konteks project, keamanan, dan sprint dengan parent-nya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues dan Browse Projects; Admin Situs mengatur toggle subtask global.
- **Aturan perilaku (acceptance criteria):**
  1. Subtask dapat dibuat dari panel subtasks di issue view (inline), dari Create dialog (field Parent wajib dipilih saat tipe Subtask dipilih), atau dari board.
  2. Company-managed dapat memiliki beberapa work type bertipe Sub-task (CFG-005); team-managed memakai satu tipe "Subtask" bawaan.
  3. Subtask mewarisi project, Security level, dan Sprint dari parent; nilai tersebut tidak dapat diubah langsung pada subtask dan ikut berubah saat parent diubah.
  4. Subtask tidak dapat memiliki subtask dan tidak dapat dipindah ke project lain tanpa parent-nya; percobaan lewat API ditolak HTTP 400.
  5. Panel subtasks menampilkan progress bar, status, assignee, priority, urutan yang bisa di-drag (rank dalam parent), toggle "Hide done", dan tautan "Open in search" (`parent = KEY`).
  6. Di board, subtask tampil sebagai card terpisah atau dikelompokkan lewat swimlane per parent ("Stories", modul 10).
  7. Subtask memiliki workflow sendiri sesuai work type-nya, terpisah dari workflow parent (modul 09).
  8. Menghapus parent ikut menghapus semua subtask, dengan peringatan yang menyebutkan jumlah subtask (ISS-033).
  9. Subtask dapat dinonaktifkan secara global di company-managed (CFG-005); saat nonaktif, tipe Sub-task tidak dapat dipilih tetapi subtask yang sudah ada tetap dapat dibuka.
  10. ⚠ Jumlah subtask per parent dibatasi oleh guardrail yang dapat dikonfigurasi (default 2.000, ISS-044); saat batas tercapai create subtask ditolak dengan pesan yang menyebutkan batasnya.
- **Team-managed vs company-managed:** Company-managed: banyak tipe subtask dan toggle global; team-managed: satu tipe Subtask yang diaktifkan lewat Features toggle project (PRJ-049).
- **Catatan & ketergantungan:** Bergantung pada CFG-005, CFG-006, dan ISS-025. Kolom `parent_id` pada `work_items` dipakai bersama untuk subtask dan child Epic.

### ISS-025 Parent Field & Epic Children

- **Deskripsi:** Field Parent terpadu yang menggantikan Epic Link dan Parent Link untuk relasi child ke parent di semua level hierarki.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues pada child dan Browse Projects pada parent.
- **Aturan perilaku (acceptance criteria):**
  1. Field Parent tersedia di issue view, Create dialog, Bulk Edit, dan JQL (`parent = KEY`, `parent IN (...)`).
  2. Picker parent mendukung pencarian key/summary, menampilkan item terbaru, dan hanya menampilkan work item dari level hierarki yang valid.
  3. Breadcrumb menyediakan aksi "Add parent", "Change parent", dan "Remove parent".
  4. Epic memiliki panel child work items, progress rollup berdasarkan status category child, dan Issue color (CFG-020) yang tampil di card dan timeline.
  5. Company-managed: Epic boleh berisi child dari project lain selama pengguna memiliki Browse di kedua project.
  6. ⚠ Team-managed: parent harus berasal dari project yang sama; relasi lintas project untuk team-managed menjadi keputusan terbuka.
  7. Data lama Epic Link, Epic Name, Epic Status, dan Epic Color dipetakan ke Parent dan Summary saat impor; produk ini tidak menyimpan field Epic legacy sebagai field terpisah (ISS-047, CFG-024).
  8. Perubahan parent tercatat di History dan tidak mengubah sprint child (kecuali subtask, yang mengikuti sprint parent).
  9. Validasi hierarki: parent harus berada tepat satu level di atas child (CFG-006); pilihan yang melanggar ditolak dengan pesan "Parent harus berupa <level> untuk work type ini".
  10. Child di project yang tidak dapat di-browse pengguna tidak ditampilkan di panel child dan tidak dihitung di progress yang dilihat pengguna tersebut.
- **Team-managed vs company-managed:** Company-managed mengizinkan parent lintas project; team-managed dibatasi ke project yang sama.
- **Catatan & ketergantungan:** Bergantung pada CFG-006, CFG-013. Level di atas Epic (Premium) didefinisikan di CFG-007 dan modul 18.

### ISS-026 Child Work Items Panel

- **Deskripsi:** Panel di issue view parent untuk menampilkan dan mengelola child work item atau subtask.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects melihat; Create Issues/Edit Issues/Transition Issues untuk aksi.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar child menampilkan ikon tipe, key, summary, priority, assignee, dan status; status dapat ditransition langsung dari baris bila pengguna memiliki Transition Issues.
  2. Progress bar dihitung dari proporsi child per status category (To Do, In Progress, Done).
  3. Create child inline menyediakan pemilih work type yang hanya memuat tipe di level tepat di bawah parent.
  4. ⚠ Aksi "Add existing" menautkan work item yang sudah ada sebagai child dengan picker yang memfilter level valid.
  5. Child dapat diurutkan dengan drag-drop; urutan memakai Rank (ISS-022).
  6. Menu per baris berisi Remove from parent, Delete, dan Open.
  7. Panel mendukung sort/filter, toggle "Hide done", dan tautan "View in search" (`parent = KEY`).
  8. ⚠ Pengguna dapat memilih kolom yang tampil di panel (mis. Story points, Due date), dan pilihan disimpan per pengguna.
  9. Saran child work item dengan AI tersedia bila AI aktif (ISS-045).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Panel yang sama dipakai untuk subtask (ISS-024) dan child Epic (ISS-025).

### ISS-027 Convert to Subtask / Convert to Work Item

- **Deskripsi:** Wizard untuk mengubah work item standard menjadi subtask atau sebaliknya di dalam project yang sama.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues (dan Move Issues bila alur memakai Move).
- **Aturan perilaku (acceptance criteria):**
  1. Convert to subtask: langkah memilih parent dan subtask type, mapping status bila workflow berbeda, mengisi field wajib target, lalu konfirmasi ringkasan.
  2. Convert to work item: langkah memilih standard work type, mapping status bila workflow berbeda, mengisi field wajib, lalu konfirmasi.
  3. Work item yang masih memiliki subtask tidak dapat dijadikan subtask, dan Epic yang memiliki child tidak dapat dikonversi langsung; aksi tidak tampil dan API menolak dengan pesan alasannya.
  4. Subtask hasil konversi mewarisi sprint dan security level parent (ISS-024).
  5. Konversi hanya dalam project yang sama; untuk berpindah project gunakan Move (ISS-032).
  6. Key tidak berubah, dan konversi tercatat di History sebagai perubahan Work type dan Parent.
  7. Permission yang dicek: Edit Issues, ditambah Move Issues untuk alur yang memakai wizard Move.
  8. ⚠ Di team-managed, konversi dilakukan lewat ganti tipe di issue view (ISS-028) dengan wizard yang sama bila status perlu dipetakan.
- **Team-managed vs company-managed:** Team-managed memakai ganti tipe di issue view; company-managed memakai menu Convert.
- **Catatan & ketergantungan:** Berbagi logika mapping status/field dengan ISS-028 dan ISS-032.

### ISS-028 Change Work Type

- **Deskripsi:** Mengganti work type sebuah work item (mis. Task menjadi Bug) dari ikon tipe di issue view atau lewat Move.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues (Move Issues bila memakai wizard Move).
- **Aturan perilaku (acceptance criteria):**
  1. Klik ikon work type di breadcrumb menampilkan daftar tipe lain di level hierarki yang sama dalam scheme project.
  2. Bila workflow, field configuration, atau screen tipe target berbeda, pengguna dialihkan ke wizard (mapping status dan pengisian field wajib); bila sama, perubahan langsung diterapkan.
  3. Perpindahan level (mis. Story ke Epic) dibatasi agar relasi hierarki tetap valid: tidak diizinkan bila work item punya parent atau child yang menjadi tidak valid, dengan pesan alasannya.
  4. ⚠ Field yang tidak ada di tipe target disembunyikan dari tampilan, tetapi nilainya tetap tersimpan dan muncul kembali bila work item dikembalikan ke tipe yang memakainya; ringkasan wizard menyebutkan field yang tidak lagi terlihat.
  5. Ganti tipe massal dilakukan lewat Bulk Move (ISS-035).
  6. Perubahan tipe tercatat di History.
- **Team-managed vs company-managed:** Tidak ada perbedaan dalam alur; daftar tipe berasal dari scheme (company-managed) atau tipe lokal project (team-managed).
- **Catatan & ketergantungan:** Bergantung pada CFG-008, CFG-009 dan modul 09 untuk mapping status.

### ISS-029 Issue Links (Linked work items)

- **Deskripsi:** Relasi berarah antar work item berdasarkan link type (blocks, relates to, duplicates, clones), dalam project yang sama maupun lintas project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Link Issues di project asal dan Browse di project target.
- **Aturan perilaku (acceptance criteria):**
  1. Form "Link work item" memuat pilihan link type beserta arah (mis. "blocks" / "is blocked by"), pencarian satu atau beberapa work item berdasarkan key/summary, lalu tombol Link.
  2. Aksi "Create linked work item" membuka Create dialog (ISS-002) dengan link type dan work item sumber sudah terisi.
  3. Panel Linked work items dikelompokkan per link type dan arah; setiap item menampilkan key, summary, status, priority, dan assignee, dapat dibuka, dan memiliki aksi Unlink.
  4. Link otomatis tampil dengan arah kebalikannya (inward/outward) di work item lawannya.
  5. Link lintas project diperbolehkan bila pengguna memiliki Browse di kedua project.
  6. Permission yang dicek: Link Issues di project asal dan Browse Projects di project target; pelanggaran ditolak HTTP 403.
  7. Item yang sudah resolved di panel tampil dengan key dicoret.
  8. JQL mendukung `issueLinkType = "blocks"` dan `issue in linkedIssues(KEY, "blocks")` (modul 12).
  9. ⚠ Jumlah link per work item dibatasi guardrail yang dapat dikonfigurasi (default 2.000, ISS-044).
  10. Pembuatan dan penghapusan link tercatat di History kedua work item.
  11. ⚠ Link ke work item di instance lain produk ini disimpan sebagai remote link (ISS-030), bukan issue link, dan tidak ikut dihitung di JQL `linkedIssues`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Link type dikelola di CFG-056; linking dapat dimatikan global (CFG-056). Link type "Cloners" dipakai otomatis oleh ISS-031. Rujukan silang ID: API-023, JQL-029, AUTO-045.

### ISS-030 Web Links, Confluence Links & Remote Links

- **Deskripsi:** Tautan dari work item ke resource eksternal: URL web, halaman wiki/Confluence lewat integrasi, dan remote link dari aplikasi atau integrasi.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Link Issues; Integrator lewat REST API.
- **Aturan perilaku (acceptance criteria):**
  1. "Add web link" meminta URL (wajib, skema http/https) dan link text (opsional); tautan dirender sebagai smart link dengan judul dan ikon yang diambil otomatis dari metadata halaman.
  2. "Link Confluence page" mencari halaman wiki lewat integrasi (modul 25) atau menerima URL; halaman tampil di section konten wiki dengan pemisahan "Linked" dan "Mentioned on".
  3. Entri "Mentioned on" dibuat otomatis saat halaman wiki yang terintegrasi menyebut key work item.
  4. REST API remote link mendukung `globalId`, `title`, `url`, `icon`, dan `status.resolved`; POST dengan `globalId` yang sama memperbarui link yang ada (idempotent), bukan membuat duplikat.
  5. Web link dapat diedit dan dihapus; perubahannya tercatat di History.
  6. Menambah, mengedit, dan menghapus web link memerlukan Link Issues.
  7. Smart link mengikuti izin pengguna atas resource; bila pengguna tidak punya akses, tautan tampil sebagai URL biasa atau dengan ajakan "Connect account".
  8. Smart link dapat ditampilkan sebagai inline, card, atau embed sesuai pilihan pengguna dan dukungan resource.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Resolusi metadata smart link dilakukan di server dengan proteksi SSRF (modul 27). Integrasi wiki pihak ketiga ada di modul 25; produk wiki tidak dibangun. Rujukan silang ID: CMT-029 (tampilan web/remote link), EDT-021, INT-002, API-023.

## Operasi siklus hidup

### ISS-031 Clone Work Item

- **Deskripsi:** Menyalin work item menjadi work item baru, dengan opsi menyertakan data terkait.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues di project tujuan dan Browse Projects pada work item asli.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Dialog Clone menampilkan Summary yang dapat diedit dengan default prefix "CLONE - " diikuti summary asli; dialog juga menyediakan opsi mengganti Assignee dan Reporter (Reporter memerlukan Modify Reporter).
  2. ⚠ Checkbox opsi: Attachments, Links, Subtasks/child work items, Sprint value, dan Comments; semua default tidak dicentang kecuali Subtasks.
  3. Link "clones / is cloned by" dibuat otomatis antara work item clone dan aslinya.
  4. ⚠ Data yang tidak pernah disalin: worklog, history, votes, watchers, dan info development; comments hanya ikut bila opsi Comments dicentang.
  5. Clone dimulai di initial status workflow, bukan status work item asli, dan Resolution dikosongkan.
  6. Clone yang membawa banyak subtask atau attachment dijalankan asinkron dengan indikator progres; pengguna menerima flag dengan tautan ke clone saat selesai.
  7. ⚠ Clone ke project lain langsung dari dialog tidak disediakan di Fase 2; pengguna melakukan clone lalu Move (ISS-032).
  8. Permission yang dicek: Create Issues di project dan Browse Projects pada work item asli; tanpa izin, aksi Clone tidak tampil.
  9. Nilai field yang tidak valid di context baru (mis. opsi select tidak aktif, versi diarsipkan) dikosongkan, dan ringkasan setelah clone menyebutkan field yang dikosongkan.
  10. Bulk clone tidak disediakan secara native; kebutuhan ini dipenuhi lewat aksi otomasi "Clone work item" (modul 20).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai job asinkron (modul 27) dan link type Cloners (CFG-056). Rujukan silang ID: AUTO-044 (aksi Clone issue untuk kebutuhan bulk clone).

### ISS-032 Move Work Item Wizard

- **Deskripsi:** Wizard bertahap untuk memindahkan work item ke project lain dan/atau work type lain, dengan mapping status dan field.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Move Issues di project asal dan Create Issues di project tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Langkah 1 memilih project dan work type tujuan, termasuk work type tujuan untuk setiap subtask type yang dimiliki work item.
  2. Langkah 2 memetakan status saat ini ke status di workflow tujuan bila status tersebut tidak ada; langkah dilewati bila status ada.
  3. Langkah 3 meminta pengisian field wajib di target dan pemetaan versions, components, serta sprint yang tidak ada di project tujuan (pilih nilai pengganti atau kosongkan).
  4. Langkah 4 menampilkan ringkasan perubahan, termasuk daftar field yang nilainya akan hilang, sebelum pengguna menekan Move.
  5. Work item mendapat key baru dari counter project tujuan; key lama tetap me-redirect dan dapat dicari (ISS-001).
  6. Subtask ikut berpindah bersama parent dan tidak dapat dipindah sendirian ke project lain.
  7. Perpindahan antara team-managed dan company-managed menampilkan peringatan bahwa custom field lokal team-managed tidak terbawa dan data sprint/board dapat hilang.
  8. Permission yang dicek: Move Issues di project asal dan Create Issues di project tujuan; project tujuan tanpa Create Issues tidak tampil di pilihan.
  9. Attachments, comments, history, links, dan watchers ikut berpindah bersama work item.
  10. Security level dipetakan ke level dengan nama sama di issue security scheme project tujuan, atau dihapus bila tidak ada padanan, dengan peringatan di ringkasan.
  11. Event `issue.moved` diterbitkan ke notifikasi dan webhook dengan key lama dan key baru.
  12. Wizard yang sama dipakai untuk mengganti work type dengan workflow berbeda di project yang sama (ISS-028).
  13. ⚠ Wizard ditampilkan sebagai modal bertahap di issue view (bukan halaman terpisah), dengan tombol Back di setiap langkah.
- **Team-managed vs company-managed:** Perpindahan lintas tipe project memicu peringatan kehilangan data (butir 7); mapping field di team-managed memakai field library project.
- **Catatan & ketergantungan:** Bergantung pada skema lengkap modul 05 dan 09, issue security (PRJ-042), dan job asinkron untuk work item dengan banyak subtask.

### ISS-033 Delete Work Item

- **Deskripsi:** Penghapusan permanen work item beserta datanya, dengan konfirmasi dan permission khusus.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Delete Issues (default hanya peran Administrator project).
- **Aturan perilaku (acceptance criteria):**
  1. Dialog konfirmasi menyatakan bahwa penghapusan bersifat permanen dan menyarankan resolve atau archive sebagai alternatif.
  2. Menghapus parent ikut menghapus semua subtask; dialog menampilkan jumlah subtask yang akan terhapus.
  3. Menghapus Epic (atau level di atasnya) tidak menghapus child; child hanya kehilangan nilai Parent dan perubahan tersebut tercatat di History child.
  4. Semua issue link di work item lain yang menunjuk ke work item terhapus ikut dihapus.
  5. Penghapusan memerlukan Delete Issues (PERM-021); tanpa izin, aksi tidak tampil dan API menolak HTTP 403.
  6. Tidak ada recycle bin atau undo untuk work item, dan nomor key yang dihapus tidak pernah dipakai ulang.
  7. Event `issue.deleted` diterbitkan ke notifikasi dan webhook dengan snapshot field utama.
  8. Bulk delete dilakukan lewat Bulk Change (ISS-035).
  9. URL work item yang sudah dihapus menampilkan pesan "Work item tidak ada atau Anda tidak memiliki izin untuk melihatnya" (tanpa membedakan dua kondisi tersebut).
- **Team-managed vs company-managed:** Tidak ada perbedaan; di team-managed Delete Issues dimiliki peran Administrator.
- **Catatan & ketergantungan:** Attachment dihapus dari storage secara asinkron (modul 08). Audit log admin mencatat penghapusan (modul 22).

### ISS-034 Archive / Unarchive Work Items

- **Deskripsi:** Fitur Premium/Enterprise untuk mengarsipkan work item agar keluar dari pencarian dan board tanpa dihapus, dan dapat dipulihkan kapan saja.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Pengguna dengan permission Archive/Unarchive work items (default Admin Situs dan Admin Proyek, PERM-035).
- **Aturan perilaku (acceptance criteria):**
  1. Satu work item dapat diarsipkan dari menu ••• setelah konfirmasi.
  2. Bulk archive dari hasil pencarian tersedia dengan maksimal 1.000 work item per operasi.
  3. Archive dan Unarchive (Restore) dikendalikan permission terpisah yang default dimiliki admin.
  4. Subtask ikut diarsipkan bersama parent dan tidak dapat diarsipkan sendiri.
  5. Work item arsip bersifat read-only dan hilang dari JQL, board, report, dan dashboard; work item tetap dapat dibuka lewat link langsung dengan banner "Archived".
  6. Halaman "Archived work items" menampilkan daftar dengan filter project, tanggal archive, dan pengarsip, serta aksi Restore per item.
  7. Restore mengembalikan work item ke kondisi semula (status, field, sprint bila sprint masih ada, link).
  8. ⚠ Daftar work item arsip dapat diekspor ke CSV dari halaman Archived work items.
  9. Key tetap dicadangkan selama diarsipkan, dan attachment serta history tetap disimpan.
  10. Archive work item berbeda dari archive project/space (PRJ-032); work item di project yang diarsipkan tidak tercantum di halaman ini.
  11. Fitur tidak tersedia di plan Free/Standard; aksi tidak tampil dan API mengembalikan galat plan (VISI-003).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada feature flag plan (VISI-003) dan PERM-035. Work item arsip dikecualikan dari indeks pencarian (modul 13). Rujukan silang ID: SRC-035 (pencarian work item arsip), API-027, ADM-025.

### ISS-035 Bulk Change (Bulk operations)

- **Deskripsi:** Operasi pada banyak work item sekaligus, lewat wizard klasik dari hasil pencarian maupun bulk edit langsung di list view/navigator.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan global permission Make bulk changes ditambah permission per operasi di setiap project.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih work item dari hasil JQL/filter dengan checkbox, Select all, atau Shift+klik; maksimal 1.000 work item per operasi.
  2. Operasi yang tersedia: Edit, Move, Transition, Delete, Watch, Stop watching, dan Archive (hanya Premium/Enterprise).
  3. ⚠ Bulk Edit field multi-value (labels, components, versions, multi-select) menawarkan opsi Add to existing, Replace all, Remove, dan Clear; opsi "Find and remove" untuk label tertentu menjadi keputusan terbuka.
  4. Bulk Edit juga dapat mengubah Parent, Sprint, Assignee, Priority, dan custom field; Resolution hanya diubah lewat Bulk Transition.
  5. Bulk Transition mengelompokkan item per workflow dan transition yang tersedia; transition screen diisi sekali per grup.
  6. Bulk Move memetakan project/tipe, status, dan field sekali per kombinasi asal (project + work type).
  7. Langkah konfirmasi menampilkan ringkasan perubahan dan jumlah item yang akan diproses.
  8. Checkbox "Send email notification" default aktif; hanya Admin Situs atau Admin Proyek dari semua project terkait yang dapat mematikannya.
  9. Global permission Make bulk changes (PERM-011) wajib dimiliki, ditambah permission per operasi (Edit/Move/Transition/Delete/Archive) di setiap project; item tanpa izin dikeluarkan dari operasi dan dilaporkan.
  10. Eksekusi berjalan di background dengan progress, notifikasi saat selesai, dan laporan daftar item yang gagal beserta alasannya.
  11. Field yang tidak tersedia di semua project/tipe terpilih tidak dapat diedit, dan alasannya ditampilkan di samping field.
  12. ⚠ List view/navigator menyediakan panel aksi bulk pada baris terpilih untuk edit cepat field umum tanpa membuka wizard.
- **Team-managed vs company-managed:** Tidak ada perbedaan; field lokal team-managed hanya dapat diedit bila semua item terpilih berasal dari project yang sama.
- **Catatan & ketergantungan:** Bergantung pada ISS-032, modul 13 (navigator), PERM-011, dan job asinkron modul 27. Batas 1.000 item mengikuti ISS-044. Rujukan silang ID: SRC-031, WF-013 (Bulk transition), NTF-028, API-026.

## Kolaborasi, berbagi & ekspor

### ISS-036 Watchers

- **Deskripsi:** Pengguna yang mengikuti work item dan menerima notifikasi perubahan sesuai notification scheme. Pemisahan cakupan: ISS-036 (modul ini) memiliki entitas watcher, UI Watch/daftar watchers, izin, toggle global, bulk watch, dan JQL; NTF-012 ([19](19-notifikasi-email.md)) memiliki perilaku penerima notifikasi All Watchers. Keduanya P0 / Fase 1 karena notifikasi Fase 1 dan autowatch comment (CMT-001) membutuhkan watcher.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects untuk watch diri sendiri; View Voters and Watchers untuk melihat daftar; Manage Watchers untuk mengelola watcher lain (PERM-033); Admin Situs untuk toggle global.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tombol Watch/Stop watching berikon mata menampilkan jumlah watcher; shortcut `w` men-toggle watch untuk pengguna saat ini.
  2. Daftar watchers dapat dibuka oleh pengguna dengan View Voters and Watchers; tanpa izin, hanya jumlah yang tampil.
  3. Menambah atau menghapus watcher lain memerlukan Manage Watchers.
  4. Menambahkan watcher yang tidak memiliki Browse pada work item menampilkan peringatan bahwa pengguna tersebut tidak akan menerima notifikasi.
  5. Auto-watch saat create dan saat comment mengikuti preferensi personal (modul 21).
  6. Setting global "Watching" ON/OFF di General configuration, yaitu CFG-057 ([05](05-konfigurasi-work-item-skema.md)) di halaman ADM-003 ([22](22-administrasi-sistem.md)); saat OFF, tombol Watch disembunyikan dan notifikasi peran Watcher tidak dikirim.
  7. Bulk watch dan stop watching tersedia lewat Bulk Change (ISS-035).
  8. JQL mendukung `watcher = currentUser()` dan `watchers > N`.
  9. Pengguna nonaktif tetap tercatat sebagai watcher tetapi tidak menerima email.
- **Team-managed vs company-managed:** Tidak ada perbedaan; di team-managed Manage Watchers dimiliki peran Administrator.
- **Catatan & ketergantungan:** Pengiriman notifikasi dan penentuan penerima All Watchers dimiliki NTF-012 di [19 Notifikasi & Email](19-notifikasi-email.md); modul ini tidak mendefinisikan ulang aturan pengirimannya. Rujukan silang ID: NTF-012, NTF-014, NAV-031 (autowatch), API-024, DM-008.

### ISS-037 Votes

- **Deskripsi:** Mekanisme voting untuk menunjukkan dukungan terhadap sebuah work item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects; View Voters and Watchers untuk daftar voter; Admin Situs untuk toggle global.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Vote/Remove vote menampilkan jumlah vote; satu pengguna maksimal satu vote per work item.
  2. Daftar voter dapat dibuka oleh pengguna dengan View Voters and Watchers.
  3. Reporter tidak dapat memberi vote pada work item yang ia laporkan; tombol nonaktif dengan tooltip alasannya.
  4. Vote tidak dapat diberikan pada work item yang sudah resolved; vote yang sudah ada tetap tersimpan.
  5. Setting global "Voting" ON/OFF di General configuration (CFG-057); saat OFF, tombol Vote dan kolom Votes disembunyikan.
  6. JQL mendukung `votes > N` dan `voter = currentUser()`.
  7. Kolom Votes tersedia di navigator dan dapat dipakai untuk sort.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Vote pada idea Product Discovery adalah mekanisme berbeda (modul 36). Rujukan silang ID: API-024, JQL-025, DM-008.

### ISS-038 Share Work Item

- **Deskripsi:** Membagikan work item ke pengguna, team, atau alamat email lewat notifikasi berisi pesan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Share di header membuka form dengan input pengguna/team/email (beberapa penerima) dan pesan opsional.
  2. Penerima mendapat email (dan notifikasi in-app bila pengguna terdaftar) berisi tautan work item, summary, dan pesan pengirim.
  3. Form menampilkan peringatan untuk penerima yang tidak memiliki Browse pada work item.
  4. Tombol "Copy link" menyalin URL `/browse/KEY`.
  5. Share ke Slack/Microsoft Teams tersedia bila integrasinya terpasang (modul 25).
  6. Share hanya mengirim notifikasi dan tidak pernah mengubah permission atau restriction work item.
  7. ⚠ Shortcut `s` membuka form Share di issue view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada pengiriman email modul 19 dan rate limit anti-spam (modul 27).
  - **Pasangan duplikat (pemilik):** [CMT-028 Share Work Item](08-komentar-lampiran-aktivitas.md) di modul 08 adalah spesifikasi pemilik untuk perilaku dialog share dan pengiriman (penerima, batas pesan, undangan pengguna baru, peringatan akses, Copy link, Slack/Teams). ISS-038 hanya mendefinisikan titik masuk di header issue view dan shortcut `s`; bila ada perbedaan, CMT-028 yang berlaku.

### ISS-039 Print & Single Work Item Export

- **Deskripsi:** Tampilan cetak dan ekspor satu work item ke XML atau Word.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Print membuka tampilan printable di `/si/jira.issueviews:issue-html/KEY/KEY.html` (rute kompatibilitas) yang memuat field, description, comment, dan link tanpa navigasi aplikasi.
  2. Export XML di `/si/jira.issueviews:issue-xml/KEY/KEY.xml` berisi field, comment, dan link dalam struktur RSS-kompatibel.
  3. Export Word menghasilkan file `.doc` berformat HTML.
  4. Konten ekspor mengikuti izin pengguna: field tersembunyi, comment terbatas, dan work item restricted tidak ikut.
  5. Ekspor banyak work item dari navigator (CSV/Excel/PDF) ada di modul 13.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rute `/si/...` dilayani backend Go (bukan SPA) dan harus dikecualikan dari fallback `index.html` di `webui`.

### ISS-040 Issue Collectors

- **Deskripsi:** Widget JavaScript yang dipasang di website eksternal untuk mengumpulkan feedback atau bug langsung menjadi work item; hanya untuk company-managed (fitur legacy).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek (Administer projects) mengelola collector; Tamu/Kolaborator Eksternal dan pengguna anonim mengirim feedback.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Issue collectors menyediakan aksi create, edit, disable, dan delete collector.
  2. Konfigurasi memilih work type, reporter default, dan opsi mencocokkan pengguna yang login atau alamat email pengirim.
  3. Template form: "Got feedback?", "Raise a bug", atau Custom dengan field yang didukung (teks, select, tanggal, dan field sistem yang aman untuk publik).
  4. Gaya trigger: tombol bawaan (posisi dan teks dapat diatur) atau trigger JavaScript kustom.
  5. Halaman collector menampilkan kode embed (`<script>` tag) dan opsi "Collect browser info" (user agent, URL halaman, resolusi layar).
  6. Pengiriman anonim diizinkan dengan nama/email opsional; pengiriman dilindungi rate limit dan CAPTCHA (AUTH-006).
  7. ⚠ Pengirim dapat melampirkan screenshot atau file sesuai batas attachment (modul 08).
  8. Halaman collector menampilkan daftar dan statistik work item yang terkumpul per collector.
  9. Issue collector tidak tersedia di team-managed project; menu tidak tampil.
  10. Reporter default harus memiliki Create Issues; penyimpanan konfigurasi ditolak bila tidak, dan collector aktif dinonaktifkan otomatis bila izin dicabut.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Alternatif modern adalah Forms (modul 26) dan portal Service Management (modul 30). Endpoint publik memerlukan CORS terbatas per collector (modul 27). Rujukan silang ID: PRJ-061 (pengaturan project), ADM-033 (daftar global), INT-006 (sudut integrasi); fitur ini dijabarkan bersama ketiganya agar tidak diimplementasikan ganda.

## Izin, keamanan & batas

### ISS-041 Work Item Permission Enforcement

- **Deskripsi:** Penerapan matriks permission project pada setiap aksi work item, di UI maupun di server; definisi scheme dan grant ada di modul 03.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs dan Admin Proyek mengatur permission scheme/role; semua pengguna dikenai pemeriksaan.
- **Aturan perilaku (acceptance criteria):**
  1. Browse Projects diperlukan untuk melihat work item, termasuk muncul di pencarian, board, link, mention, dan notifikasi (PERM-020).
  2. Create Issues, Edit Issues, Transition Issues, Delete Issues, dan Move Issues mengontrol aksi create, edit, transition, delete, dan move (PERM-021, PERM-023, PERM-024).
  3. Assign Issues, Assignable User, Modify Reporter, dan Schedule Issues (Due date dan Rank) mengontrol field terkait (PERM-022, PERM-024).
  4. Link Issues, Set Issue Security, Manage Watchers, dan View Voters and Watchers mengontrol link, restriction, dan daftar watcher/voter (PERM-024, PERM-033, PERM-034).
  5. Archive dan Unarchive work items (Premium/Enterprise) mengontrol ISS-034 (PERM-035).
  6. Global permission Make bulk changes diperlukan untuk Bulk Change (PERM-011).
  7. Edit/Delete own vs Edit/Delete all untuk comments, attachments, dan worklogs ditegakkan di modul 08 dan 14 (PERM-030 s.d. PERM-032).
  8. Aksi yang tidak diizinkan disembunyikan (menu, tombol) atau dinonaktifkan dengan tooltip yang menyebutkan permission yang dibutuhkan bila keberadaan aksinya penting untuk dipahami (mis. field read-only).
  9. Setiap endpoint REST memeriksa permission di server terlepas dari UI; pelanggaran mengembalikan HTTP 403 untuk work item yang terlihat, dan HTTP 404 untuk work item yang tidak boleh di-browse agar keberadaannya tidak bocor (PERM-002).
  10. ⚠ Team-managed memakai role bawaan Administrator, Member, dan Viewer; custom role team-managed hanya tersedia di plan tertentu sesuai VISI-003 dan PERM-005.
- **Team-managed vs company-managed:** Company-managed memakai permission scheme; team-managed memakai role bawaan project (PRJ-038, PRJ-039).
- **Catatan & ketergantungan:** Pemeriksaan permission diimplementasikan sebagai layanan bersama di use case (bukan di handler) agar dipakai ulang oleh REST, otomasi, dan job bulk. Middleware `AuthMiddleware` yang sudah ada di repo hanya memvalidasi identitas; otorisasi per work item adalah pekerjaan baru.

### ISS-042 Restrict Work Item (Issue Security Level)

- **Deskripsi:** Membatasi siapa yang dapat melihat work item tertentu lewat security level (company-managed) atau restriction berbasis role (team-managed); tidak tersedia di plan Free.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Set Issue Security (PERM-034); Admin Proyek mengatur issue security scheme (PRJ-042).
- **Aturan perilaku (acceptance criteria):**
  1. Ikon gembok di header issue view membuka pilihan restriction untuk memilih atau menghapus level; ikon menunjukkan status terkunci bila level terpasang.
  2. Company-managed memakai Issue Security Scheme yang setiap level-nya berisi user/group/project role/field pengguna; mengubah level memerlukan Set Issue Security.
  3. Team-managed menyediakan "Restrict to roles" di issue view yang membatasi work item ke satu atau beberapa role project.
  4. Default security level dari scheme diterapkan saat create bila pengguna tidak memilih.
  5. Subtask mewarisi security level parent dan tidak dapat diubah sendiri (ISS-024).
  6. Pengguna tanpa akses melihat pesan "Work item tidak ada atau Anda tidak memiliki izin untuk melihatnya", dan work item tidak muncul di pencarian, board, dashboard, maupun link/mention bagi pengguna tersebut.
  7. Perubahan level tercatat di History.
  8. Fitur tidak tersedia di plan Free; ikon gembok tidak tampil dan API menolak field Security level dengan galat plan (VISI-003).
- **Team-managed vs company-managed:** Company-managed: security level berbasis scheme; team-managed: restriction berbasis role project.
- **Catatan & ketergantungan:** Konfigurasi scheme di PRJ-042 dan PERM-034. Filter keamanan harus diterapkan di level query (modul 12, 13) agar hitungan agregat tidak membocorkan work item.

### ISS-043 Data Redaction (Atlassian Guard)

- **Deskripsi:** Penghapusan permanen (redaksi) teks sensitif di work item, mis. password atau data pribadi di summary, description, comment, atau field teks, secara manual atau dari temuan deteksi otomatis, dengan jejak audit yang tetap tersimpan. Di produk ini kemampuan ini dibangun sebagai fitur keamanan opsional berbayar, bukan produk terpisah.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs/admin organisasi dan peran redaksi yang ditunjuk; tim keamanan.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Pengguna berwenang memilih teks di Summary, Description, comment, atau field teks lalu menjalankan aksi "Redact" dari menu konteks seleksi di issue view.
  2. Dialog meminta alasan redaksi (wajib) dan konfirmasi eksplisit bahwa aksi permanen dan tidak dapat dibatalkan.
  3. Konten yang diredaksi diganti placeholder "redacted" untuk semua pengguna, termasuk di REST API, ekspor, dan notifikasi berikutnya.
  4. ⚠ Konten yang diredaksi juga dihapus dari History/changelog, versi lama comment/description, dan indeks pencarian.
  5. ⚠ Hanya admin organisasi atau pengguna dengan peran redaksi yang ditunjuk yang dapat meredaksi; pengguna lain tidak melihat aksi Redact.
  6. Temuan dari modul deteksi data sensitif (bila diaktifkan) ditampilkan sebagai saran redaksi yang dapat diterima atau diabaikan per temuan.
  7. Setiap redaksi tercatat di audit log organisasi (modul 22) dengan aktor, waktu, work item, field, dan alasan, tanpa menyimpan teks asli.
  8. Proses redaksi berjalan asinkron; REST Redaction API memakai async task dengan status yang dapat dipantau (modul 23).
  9. Fitur hanya tersedia bila add-on keamanan diaktifkan untuk organisasi (feature flag plan, VISI-003); tanpa add-on, aksi dan menu tidak tampil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada audit log modul 22, API async modul 23, dan reindeks modul 13. Mesin deteksi data sensitif otomatis berada di luar cakupan awal (lihat Di luar cakupan). Rujukan silang ID: PERM-066 (izin redaksi), API-028 (Issue Redaction API).

### ISS-044 Work Item Limits & Guardrails

- **Deskripsi:** Batas teknis pada data work item untuk menjaga performa, yang dapat dikonfigurasi admin sistem dalam batas aman.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Sistem menegakkan; Admin Situs melihat dan (dalam rentang tertentu) mengatur nilai.
- **Aturan perilaku (acceptance criteria):**
  1. Summary maksimal 255 karakter dan field teks maksimal 32.767 karakter, berlaku identik di UI, API, impor, dan otomasi.
  2. ⚠ Batas default per work item: comments 5.000, attachments 2.000, issue links 2.000, remote links 2.000, worklogs 10.000, dan subtasks 2.000; nilai disimpan sebagai konfigurasi sistem.
  3. Bulk operation maksimal 1.000 work item per operasi.
  4. Ukuran maksimal attachment per file mengikuti setting admin (modul 08).
  5. Saat batas terlampaui, UI menampilkan pesan yang menyebutkan batas dan jenis data, dan REST API mengembalikan HTTP 400 (atau 413 untuk ukuran payload) dengan kode galat yang stabil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nilai guardrail didokumentasikan bersama NFR di modul 27; batas field teks juga dipakai ISS-003 dan ISS-015.
  - **Pasangan duplikat (pemilik):** [NFR-004 Limits & Guardrails](27-nfr-arsitektur-platform.md) di modul 27 adalah katalog pemilik untuk seluruh batas sistem, format galat dan error code, serta publikasi daftar batas. ISS-044 adalah rincian domain work item yang dirujuk NFR-004 dan menjadi acuan nilai guardrail per work item. ⚠ Saat ini nilai default subtask berbeda (ISS-044 butir 2: 2.000; NFR-004 butir 5: 1.000) dan remote links hanya tercantum di ISS-044; nilai final ditetapkan satu kali di NFR-004 lalu disalin ke sini (lihat Pertanyaan terbuka).

## AI & agent (opsional)

### ISS-045 Atlassian Intelligence / Rovo on Work Items

- **Deskripsi:** Fitur AI opsional di issue view untuk meringkas, menulis, memecah pekerjaan, dan menemukan konten terkait. Di produk ini fitur ini dibangun di atas penyedia LLM yang dapat dikonfigurasi dan dapat dimatikan sepenuhnya.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Browse Projects (dan Edit Issues untuk menerapkan hasil); Admin Situs/admin organisasi mengaktifkan AI.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ "Summarize" menghasilkan ringkasan comment dan aktivitas panjang di tab Activity.
  2. ⚠ AI di editor menyediakan aksi improve writing, ubah tone, ringkas, dan buat Description dari Summary; hasil selalu ditampilkan sebagai pratinjau yang harus diterima pengguna.
  3. ⚠ AI work breakdown menyarankan daftar child work item/subtask; pengguna memilih item yang dibuat dan item dibuat lewat layanan create biasa (ISS-003).
  4. ⚠ Deteksi work item serupa/duplikat menampilkan kandidat saat create dan di issue view.
  5. ⚠ Definitions menjelaskan istilah/akronim di teks berdasarkan konten yang dapat diakses pengguna.
  6. ⚠ Natural language ke JQL disediakan di modul 13.
  7. ⚠ Asisten chat membaca konteks work item untuk menjawab pertanyaan; aksi yang mengubah data didefinisikan di ISS-046.
  8. ⚠ Admin dapat mengaktifkan/menonaktifkan AI per site dan per produk; AI hanya memakai data yang dapat di-browse pengguna dan menghormati pengaturan residensi data.
  9. ⚠ Konten hasil AI diberi label "Dibuat dengan AI" dan tombol feedback (thumbs up/down).
  10. ⚠ AI tersedia di plan Standard, Premium, dan Enterprise sesuai feature flag plan (VISI-003), tidak di Free.
  11. ⚠ Field suggestions on create (rilis Oktober 2026): saat membuat work item, sistem menyarankan nilai Labels, Assignee, Parent, Priority, Components, dan custom field select berdasarkan work item serupa yang dapat di-browse pengguna; saran tampil sebagai chip/pratinjau per field, tidak ada nilai yang diisi ke field sebelum pengguna menerimanya, dan nilai yang diterima tetap melewati validasi ISS-003.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia LLM, kuota, dan logging prompt didefinisikan di modul 27. Rancangan awal mengasumsikan organisasi menyediakan kredensial penyedia LLM sendiri (BYO-key); apakah produk juga menyediakan penyedia bawaan masih menjadi pertanyaan terbuka. Rujukan silang ID: JQL-059 dan SRC-019 (natural language ke JQL).
  - **Pasangan duplikat (pemilik):** [ADM-052 Atlassian Intelligence & Rovo settings](22-administrasi-sistem.md) di modul 22 adalah kontrol pusat dan spesifikasi pemilik untuk aktivasi AI, pembatasan per group, kuota, kebijakan data, dan profil regulated. ISS-045 hanya mendefinisikan perilaku fitur AI di work item; butir 8 dan 10 di atas tunduk pada ADM-052, dan bila berbeda (mis. ADM-052 butir 14 menyatakan AI tidak digating tier, sedangkan butir 10 di atas menyebut tidak tersedia di Free), ADM-052 yang berlaku.

### ISS-046 Agents in Jira (Rovo agents as assignee & @mention)

- **Deskripsi:** Agent AI diperlakukan sebagai rekan kerja di work item: dapat dijadikan Assignee atau di-@mention di comment untuk meminta bantuan, lalu mengerjakan tugas dan melaporkan hasilnya ke work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Assign Issues/Add Comments; Admin Situs mengatur agent yang tersedia; agent bertindak sebagai aktor sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat dipilih di field Assignee; user picker menampilkan agent dengan penanda AI yang membedakannya dari pengguna manusia.
  2. @mention agent di comment atau editor menjalankan permintaan dalam konteks work item tersebut.
  3. ⚠ Agent menulis hasilnya sebagai comment, saran perubahan field, atau usulan child work item; perubahan baru diterapkan setelah pengguna meninjau dan menerima.
  4. ⚠ Issue view menampilkan indikator status kerja agent: sedang bekerja, selesai, atau gagal (dengan alasan).
  5. Agent menghormati permission pengguna pemicu dan permission agent sekaligus (irisan keduanya), sehingga tidak dapat mengakses atau mengubah work item yang tidak boleh diakses.
  6. Aksi agent tercatat di History/Activity dengan agent sebagai aktor dan pengguna pemicu sebagai konteks.
  7. Admin dapat membatasi agent yang tersedia per site atau per project lewat pengaturan AI.
  8. ⚠ Pemakaian agent dicatat terhadap kuota AI organisasi; model penagihan menjadi keputusan terbuka.
  9. ⚠ Agent yang dinonaktifkan atau dihapus padahal masih menjadi assignee ditampilkan sebagai "Agent tidak tersedia", dan work item dapat dialihkan ke Unassigned lewat aksi admin atau otomatis sesuai konfigurasi.
  10. Agent sebagai assignee atau watcher tidak menerima notifikasi email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada ISS-016, ISS-045, dan modul 20 (agent dalam workflow/otomasi). Status fitur di produk acuan masih beta; dibangun setelah ISS-045 stabil.

## Warisan & kompatibilitas

### ISS-047 Legacy Issue View & Epic Fields

- **Deskripsi:** Elemen lama yang sudah deprecated di Cloud atau hanya ada di Data Center; dicatat sebagai rujukan kompatibilitas data dan tidak dibangun sebagai fitur.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (impor dan kompatibilitas).
- **Aturan perilaku (acceptance criteria):**
  1. Old issue view dan dialog "Edit issue" klasik tidak dibangun; seluruh pengeditan memakai issue view baru (ISS-007, ISS-010).
  2. Epic Link, Epic Name, Epic Status, dan Epic Color tidak dibangun sebagai field; data impor dipetakan ke Parent, Summary, status, dan Issue color (ISS-025, CFG-024).
  3. Parent Link dari perencanaan lanjutan tidak dibangun; nilainya dipetakan ke Parent.
  4. Pola project key yang dapat dikonfigurasi (`jira.projectkey.pattern`) tidak didukung; aturan key tetap mengikuti ISS-001.
  5. Tab "Transitions" di Activity tidak dibangun; riwayat transition tersedia di tab History.
  6. JQL dan REST menerima nama field lama (mis. `"Epic Link"`) sebagai alias yang dipetakan ke `parent` agar filter dan integrasi lama tetap berjalan (modul 12, 23).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; lihat Di luar cakupan.

## Pertanyaan terbuka

- Apakah Create dialog perlu mengizinkan pemilihan status awal selain initial status workflow (ISS-002), dan bila ya, apakah dikontrol per transition atau per permission?
- Apakah team-managed project mengizinkan Parent lintas project (ISS-025), atau tetap dibatasi ke project yang sama demi kesederhanaan model izin?
- Apakah Bulk Edit labels perlu opsi "Find and remove" untuk menghapus satu label tertentu dari banyak work item tanpa mengganti label lain (ISS-035)?
- Apakah Clone perlu mendukung pemilihan project tujuan langsung di dialog, atau cukup alur Clone lalu Move (ISS-031)?
- Bagaimana model penagihan kuota AI dan agent (per pengguna, per organisasi, atau berbasis kredit), dan apakah produk menyediakan penyedia LLM bawaan atau hanya BYO-key (ISS-045, ISS-046)?
- Apakah fitur Data Redaction diposisikan sebagai add-on berbayar terpisah atau bagian plan Enterprise, dan apakah mesin deteksi data sensitif otomatis dibangun sendiri (ISS-043)?
- Apakah nilai guardrail per work item (ISS-044) dapat diubah Admin Situs pada deployment self-hosted, dan berapa batas atas yang aman?
- Berapa nilai default final guardrail subtask per work item (ISS-044 menyebut 2.000, NFR-004 menyebut 1.000), dan apakah remote links ikut dibatasi di katalog NFR-004? Nilai yang diputuskan ditetapkan di NFR-004 lalu disinkronkan ke ISS-044.
- Apakah fitur AI di work item (ISS-045) digating plan (tidak tersedia di Free) atau, sesuai ADM-052, hanya bergantung pada konfigurasi penyedia LLM dan toggle admin?
- Apakah Issue Collectors (ISS-040) tetap dibangun mengingat statusnya legacy, atau langsung digantikan Forms publik dan portal Service Management?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga:** bulk clone, multi-assignee native, template work item, checklist, dan ekstensi issue view dari Marketplace tidak dibangun sebagai fitur native; titik ekstensi (panel apps, quick-add dari ekstensi, custom field dari ekstensi) disediakan lewat modul 25.
- **Fitur Data Center-only/deprecated (ISS-047):** old issue view dan dialog Edit klasik, field Epic Link/Epic Name/Epic Status/Epic Color dan Parent Link sebagai field terpisah, pola project key konfigurabel `jira.projectkey.pattern`, dan tab Transitions di Activity; hanya didukung sebagai pemetaan data impor dan alias JQL/REST.
- **Layanan lintas produk vendor acuan:** produk wiki (Confluence) tidak dibangun, hanya integrasi tautan (ISS-030, ISS-006); layanan keamanan organisasi terpisah (Guard) dan mesin deteksi data sensitifnya tidak dibangun sebagai produk, hanya kemampuan redaksi opsional di ISS-043; asisten AI dan agent vendor acuan (Atlassian Intelligence/Rovo) tidak dipakai, digantikan lapisan AI opsional dengan penyedia LLM yang dapat dikonfigurasi (ISS-045, ISS-046).
- **Link lintas instance lewat application link:** sinkronisasi dua arah dengan instance produk lain tidak dibangun; hanya remote link satu arah (ISS-029, ISS-030).
- **Integrasi chat (Slack/Microsoft Teams):** perilaku create/share dari chat didefinisikan di modul 25, bukan di modul ini.
