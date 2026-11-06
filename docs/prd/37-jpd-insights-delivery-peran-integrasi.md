# 37 — Product Discovery: Insights, Delivery, Peran, Integrasi & API

Modul ini mendefinisikan lapisan **bukti, eksekusi, akses, dan konektivitas** Product Discovery di produk ini, di atas inti idea/field/view yang dimiliki modul 36: model **akses dan lisensi** (Creator berlisensi, Contributor gratis, Stakeholder melalui published view), **access level** proyek (Open/Limited/Private), role dan permission proyek, pembatasan view dan roadmap, pengaturan global, serta plan dan batasnya; **insights** sebagai entitas bukti (feedback pelanggan, riset, tiket support, peluang sales) yang dilampirkan ke idea lewat tab Insights, tempel tautan (Smart Link), data provider, ekstensi browser, dan konektor helpdesk/CRM, termasuk field jumlah insight dan skor berbasis insight; **delivery tracking** yang menautkan idea ke work item eksekusi (Epic/Story/Task) di proyek Jira lain dengan field Delivery progress, Delivery status, autofill tanggal, panel idea di sisi work item dan Plans, serta template otomasi; **kolaborasi dan berbagi** (komentar dan reaction, votes, komentar view, share, publish, akses stakeholder, notifikasi); **integrasi** chat (Slack, Microsoft Teams), dokumentasi, Service Management, dan Goals; serta **data dan platform** (JQL, REST API & webhook, ekspor, analitik/data lake, AI, template, dukungan mobile, dan nama legacy). Idea tetap work item biasa dengan `project_type_key='product_discovery'`, sehingga link, komentar, watcher, otomasi, JQL, dan REST API inti dipakai ulang; modul ini hanya mengatur perilaku khusus discovery. Modul ini menyebut wadah discovery sebagai "project discovery" atau "proyek" (label UI "space" mengikuti VISI-002), selaras dengan modul 36.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas project, work item, link, komentar, maupun peran produk. Seluruh fitur modul ini adalah pekerjaan baru dan diusulkan menjadi feature module tambahan di samping modul discovery dari modul 36: `internal/features/discoveryinsight` (insight, label, impact, copy, capture API untuk ekstensi/chat), `internal/features/discoverydelivery` (tab Delivery, konfigurasi progress, autofill dates, panel idea), dan `internal/features/discoveryaccess` (role produk Creator/Contributor, access level, role proyek, view/roadmap restriction, publish settings). Entitas bersama diusulkan di `internal/shared/domain/entity` (mis. `Insight`, `InsightLabel`, `InsightSnippet`, `DeliveryLink`, `DeliveryProgressConfig`, `ProductRoleAssignment`, `SpaceAccessLevel`, `ViewRestriction`), UI di `frontend/src/routes` (SPA SvelteKit), dan route mengikuti konvensi repo `/api/v1/discovery/...`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — tier plan, istilah project/proyek dan work item (VISI-002), batas cakupan.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — app access & product role (USR-018), billable users (USR-020), audit log organisasi (USR-045).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — product access (PERM-004), batas izin per plan (PERM-005), izin voters & watchers (PERM-033), izin Plans (PERM-052).
- [06 Work Item Inti](06-work-item-inti.md) — issue links (ISS-029), remote links (ISS-030), watchers (ISS-036).
- [07 Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) — mention (EDT-014), Smart Links inline/card/embed (EDT-021), AI di editor (EDT-030).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar dan reaction (CMT-001 s.d. CMT-006).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category (WF-002) sebagai dasar perhitungan delivery.
- [12 Bahasa Kueri (JQL)](12-jql-bahasa-kueri.md) — field klasifikasi project (JQL-017), field engagement (JQL-025).
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — saved filter (SRC-036), ekspor CSV (SRC-050), Excel/Google Sheets (SRC-055).
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Plans (PLAN-027), Goals (PLAN-061, PLAN-062).
- [19 Notifikasi & Email](19-notifikasi-email.md) — watchers (NTF-012), mention (NTF-013), notifikasi Slack/Teams personal (NTF-039, NTF-040).
- [20 Otomasi](20-otomasi.md) — library template (AUTO-010), administrasi global (AUTO-002), batas pemakaian (AUTO-015), smart values (AUTO-063).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log (ADM-014), application links (ADM-032), pengaturan AI (ADM-052).
- [23 REST API & Webhook](23-api-publik-webhook.md) — Issues API (API-016), link API (API-023), votes & watchers API (API-024), katalog event webhook (API-065).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — registry Smart Link (INT-002), pola helpdesk/CRM (INT-004), CSV importer (INT-008), analitik & data lake (INT-029), Slack (INT-030 s.d. INT-035), Microsoft Teams (INT-036), dokumentasi (INT-042 s.d. INT-045), Goals & Projects (INT-056), rujukan delivery discovery (INT-057).
- [16 Laporan & Insight](16-laporan-insight.md) — Analytics & Data Lake (RPT-038 s.d. RPT-041).
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — Smart Link platform (NAV-017), paritas app mobile (NAV-050).
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) — lisensi agent (JSM-PRJ-046) dan request sebagai sumber insight.
- [36 Product Discovery: Ideas, Fields, Formula & Views](36-jpd-ideas-field-views.md) — idea, idea view, field (termasuk Formula Fields dan Insights & Delivery Computed Fields), view, publish dasar, roadmap lintas proyek.

## Persona terkait

- **Product Manager Discovery** — Creator berlisensi: mengumpulkan insight, menilai idea, menautkan delivery, mengonfigurasi progress, membagikan dan mem-publish view.
- **Kontributor/Stakeholder Discovery** — Contributor gratis: melihat proyek, menambah insight, vote, comment, react, dan (bila diizinkan) mengajukan idea; stakeholder melihat published view.
- **Admin Situs** — menetapkan role produk Creator/Contributor, mengaktifkan publish view, mengelola pengaturan global Product Discovery, plan, dan aplikasi chat.
- **Admin Proyek/Project Lead** — pemegang role Administrator di project discovery: access level, role, permission, fitur proyek, template otomasi, notifikasi.
- **Developer/Anggota Tim** — melihat idea terkait dari work item delivery dan menjalankan pekerjaan yang dilacak progresnya.
- **Scrum Master/Team Lead** — perencana yang mengelompokkan work item per idea di Plans.
- **Product Owner/PM** — pemilik backlog delivery yang menerima work item dari idea dan memastikan tautan delivery akurat.
- **Agent Layanan** — menautkan request Service Management sebagai insight dan meneruskan permintaan fitur dari portal.
- **Stakeholder/Viewer** — pembaca roadmap melalui published view atau embed dokumentasi.
- **Integrator/Developer API** — membaca/menulis idea lewat REST API, webhook, JQL, ekspor, dan data lake.

## User story modul

- Sebagai Product Manager Discovery, saya ingin melampirkan feedback pelanggan sebagai insight beserta tautan sumbernya ke idea agar keputusan prioritas didukung bukti yang dapat ditelusuri (JPD-INS-009, JPD-INS-010).
- Sebagai Kontributor/Stakeholder Discovery, saya ingin menangkap kutipan dari halaman web mana pun lewat ekstensi browser agar feedback tidak hilang di tool lain (JPD-INS-014).
- Sebagai Product Manager Discovery, saya ingin menyalin satu insight ke beberapa idea dan melihat jumlah serta skor insight di view agar satu bukti dapat menguatkan banyak idea yang terkait (JPD-INS-011, JPD-INS-012).
- Sebagai Agent Layanan, saya ingin menautkan request Service Management dan ticket helpdesk ke idea agar suara pelanggan dari support terhitung dalam prioritisasi (JPD-INS-013, JPD-INS-015, JPD-INS-035).
- Sebagai Product Manager Discovery, saya ingin membuat atau menautkan Epic delivery dari idea dan melihat progresnya per status category agar saya tahu kapan idea benar-benar terkirim (JPD-INS-017, JPD-INS-018, JPD-INS-019).
- Sebagai Product Manager Discovery, saya ingin membuat work item delivery untuk banyak idea sekaligus dan mengisi tanggal idea otomatis dari work item terkait agar roadmap tetap sinkron tanpa entri ganda (JPD-INS-020, JPD-INS-021).
- Sebagai Developer/Anggota Tim, saya ingin melihat idea terkait di work item dan di Plans agar saya memahami konteks dan bukti di balik pekerjaan (JPD-INS-022).
- Sebagai Admin Proyek/Project Lead, saya ingin template otomasi yang menandai idea selesai saat semua Epic delivery selesai agar status discovery tidak basi (JPD-INS-023).
- Sebagai Admin Situs, saya ingin hanya Creator yang dihitung sebagai seat berbayar sementara pengguna lain menjadi Contributor gratis agar seluruh organisasi bisa berkontribusi tanpa biaya lisensi tambahan (JPD-INS-001, JPD-INS-008).
- Sebagai Admin Proyek/Project Lead, saya ingin mengatur access level, role, dan izin Contributor membuat idea agar proyek terbuka bagi feedback namun tetap terkendali (JPD-INS-002, JPD-INS-003, JPD-INS-004).
- Sebagai Product Manager Discovery, saya ingin membatasi siapa yang dapat melihat atau mengedit view dan roadmap tertentu agar draf prioritas sensitif tidak tersebar (JPD-INS-005, JPD-INS-006).
- Sebagai Stakeholder/Viewer, saya ingin membuka roadmap yang di-publish tanpa lisensi dan memberi feedback lewat komentar atau vote agar suara saya didengar tim produk (JPD-INS-026, JPD-INS-027, JPD-INS-029, JPD-INS-030).
- Sebagai Kontributor/Stakeholder Discovery, saya ingin membuat idea dan insight langsung dari pesan Slack atau Microsoft Teams agar feedback internal tertangkap di tempat percakapan terjadi (JPD-INS-032, JPD-INS-033).
- Sebagai Integrator/Developer API, saya ingin membaca idea dan field discovery lewat REST API, JQL, webhook, dan ekspor agar data discovery dapat dianalisis di sistem lain (JPD-INS-037, JPD-INS-038, JPD-INS-039, JPD-INS-040).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 7 |
| P1 | 20 |
| P2 | 10 |
| P3 | 7 |
| **Total** | **44** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 0 |
| Fase 4 | 0 |
| Fase 5 | 0 |
| Fase 6 | 44 |
| Fase 7 | 0 |
| Fase — | 0 |

Strategi urutan implementasi: seluruh modul berada di Fase 6 (Product Discovery) dan baru dimulai setelah inti idea, field, dan view modul 36 tersedia, karena insight, delivery, dan izin melekat pada idea dan view. Gelombang 6a membangun fondasi akses (JPD-INS-001, JPD-INS-002), tab Insights dengan tempel tautan dan field jumlah insight (JPD-INS-009, JPD-INS-010, JPD-INS-012), serta tab Delivery dengan Delivery progress (JPD-INS-017, JPD-INS-018); gelombang ini sudah cukup untuk siklus "kumpulkan bukti → prioritaskan → tautkan Epic → pantau progres". Gelombang 6b menambah fitur P1: role proyek, pengajuan idea oleh Contributor, copy insight, data provider, ekstensi browser, Delivery status, bulk create, panel idea di Plans, template otomasi, komentar/votes, share/publish, notifikasi, Slack, dokumentasi, ekspor, template, dan plan. Gelombang 6c menampung fitur P2/P3 yang bergantung pada plan Premium/Enterprise atau modul lain yang lebih belakang (view/roadmap restriction, autofill dates, Microsoft Teams, Service Management, Goals, JQL/REST khusus, AI, integrasi instance self-hosted, data lake, mobile, nama legacy). Inventaris menggabungkan 23 fitur ringkas dari domain lampiran (permissions, notifications, admin, API, lampiran Product Discovery) ke fitur rinci yang setara; lihat baris **Alias inventaris** di setiap fitur.

## Fitur

## Akses, lisensi, peran & plan

### JPD-INS-001 JPD App Access & Licensing (Creators vs Contributors)

- **Alias inventaris:** Jira Product Discovery Roles; Roles: Creators, Contributors, Stakeholders
- **Deskripsi:** Model lisensi dan role produk Product Discovery: Creator berlisensi dengan akses penuh untuk membuat dan mengelola idea, view, insight, dan delivery; Contributor tidak berlisensi yang dapat berkontribusi tanpa app access khusus; Stakeholder melihat published view secara read-only. Role produk menentukan batas atas hak, sedangkan role proyek (JPD-INS-003) menentukan hak di tiap proyek.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs dan org admin (menetapkan role produk), Admin Proyek/Project Lead (role Administrator; menugaskan role proyek), Creator, Contributor, Stakeholder.
- **Aturan perilaku (acceptance criteria):**
  1. Creator (berlisensi, seat berbayar) dapat membuat dan mengedit idea, view, insight, delivery, dan sharing; membuat field dan menjalankan administrasi proyek hanya bila juga memegang role Administrator di proyek tersebut.
  2. Contributor (tidak berlisensi, gratis bagi pengguna yang sudah punya akses produk lain) memiliki akses read-only ke proyek ditambah comment, add insight, vote, dan reaction; membuat idea hanya bila setting "Allow contributors to create ideas" aktif (JPD-INS-004).
  3. Contributor tidak dapat mengedit idea, membuat atau mengubah field dan view, maupun mengelola proyek; percobaan lewat API mengembalikan HTTP 403 dengan pesan "Aksi ini memerlukan peran Creator".
  4. Pengguna yang memiliki app access ke pelacak kerja (software/business) atau Service Management otomatis menjadi Contributor di Product Discovery tanpa penugasan tambahan.
  5. Org admin memberi akses Creator lewat Directory > pengguna > Grant access > Product Discovery > pilih role, atau dengan menambahkan pengguna ke grup default Creator produk.
  6. Halaman Apps > Product Discovery > Access memungkinkan admin mengaktifkan grup sehingga anggotanya dapat membuka proyek sebagai Contributor; UI memperingatkan "Grup ini akan memakai lisensi Creator" bila grup yang dipilih juga merupakan grup app access Creator.
  7. Perhitungan billing hanya menghitung pengguna dengan role produk Creator; halaman billing menampilkan jumlah Creator aktif terhadap kuota plan.
  8. Pengguna yang hanya memiliki akses Product Discovery sebagai Contributor tidak dihitung sebagai billable user pada langganan keamanan organisasi (padanan Guard Standard, USR-037).
  9. Hanya Admin Situs (atau org admin) yang dapat menetapkan role produk di tingkat site; setiap perubahan role produk dicatat di audit log organisasi (USR-045) dengan aktor, pengguna target, role lama, dan role baru.
  10. ⚠ Nama role produk yang tampil di UI admin adalah `Creator` dan `Contributor`; nilai disimpan sebagai enum stabil (`creator`, `contributor`) sehingga perubahan label tampilan tidak mengubah data atau API.
  11. Mencabut akses Creator dari pengguna yang masih memiliki app access produk lain menurunkannya otomatis menjadi Contributor tanpa menghapus idea, insight, atau view yang pernah ia buat.
  12. Rincian izin per proyek (access level, role proyek, custom role) dikelola di JPD-INS-002 dan JPD-INS-003; role proyek tidak dapat memberikan hak Creator kepada pengguna yang di tingkat produk hanya Contributor.
  13. Stakeholder melihat published view dan roadmap secara read-only tanpa lisensi dan tanpa app access (JPD-INS-029); Stakeholder bukan role produk yang dapat ditugaskan di halaman akses.
  14. Administrator proyek mengelola pengaturan proyek (fitur, access, people, field, workflow, otomasi); role produk Creator saja tidak cukup untuk membuka Project settings.
  15. Penugasan role proyek dilakukan di Project settings > Access (People): tambah pengguna atau grup, pilih role, cari dan ubah role yang sudah ada, serta keluarkan anggota.
  16. UI menonaktifkan (disabled dengan tooltip "Hanya Creator yang dapat melakukan ini") aksi yang tidak diizinkan untuk Contributor, misalnya edit nilai field, membuat view, atau menautkan delivery, alih-alih menampilkan galat setelah aksi dijalankan.
  17. Halaman Apps > Product Discovery > Access menampilkan baris hak role produk Creator: create/edit idea, field, view, insight, dan delivery, dengan label "Seat berbayar", sehingga admin melihat dampak lisensi sebelum menugaskan role.
  18. Matriks yang sama menampilkan baris hak role produk Contributor: lihat idea, comment, add insight, vote, react, dan (opsional, bila diaktifkan per proyek) create idea, dengan label "Gratis"; aksi di luar daftar ini ditolak untuk Contributor baik lewat UI maupun API.
  19. Status Contributor gratis hanya berlaku bagi pengguna yang memiliki akses ke produk lain di site; pengguna tanpa akses produk mana pun yang membuka tautan proyek melihat halaman "Minta akses" dan tidak otomatis menjadi Contributor.
- **Team-managed vs company-managed:** Tidak ada perbedaan; project discovery selalu memakai model akses discovery tersendiri.
- **Catatan & ketergantungan:** App access dan product role umum di USR-018 ([02](02-pengguna-grup-organisasi.md)) dan PERM-004 ([03](03-izin-peran-keamanan.md)); kuota plan di JPD-INS-008. Keputusan terbuka: apakah produk ini menamai ulang Creator/Contributor di glosarium VISI. Catatan migrasi: tabel `users` saat ini belum punya kolom role produk; diperlukan tabel `product_role_assignments` (user_id, product, role).

### JPD-INS-002 Space Access Levels (Open / Limited / Private)

- **Alias inventaris:** Project access and permissions (JPD)
- **Deskripsi:** Tingkat akses default project discovery bagi semua pengguna site (Open, Limited, Private) beserta penugasan pengguna/grup dan batasan hak Contributor. Access level menentukan role default yang didapat pengguna yang tidak ditambahkan secara eksplisit.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; mengubah access level dan anggota), Admin Situs (selalu dapat membuka proyek), Creator, Contributor.
- **Aturan perilaku (acceptance criteria):**
  1. Open: setiap Contributor site dapat membuka proyek sebagai Contributor dan setiap Creator site otomatis menjadi Creator di proyek.
  2. Limited: semua pengguna site dapat membuka proyek sebagai Contributor; hanya pengguna/grup yang ditambahkan secara eksplisit menjadi Creator atau Administrator.
  3. Private: hanya pengguna/grup yang ditambahkan (dan Admin Situs) yang dapat membuka proyek; pengguna lain mendapat HTTP 404 saat membuka URL proyek, idea, atau view-nya, dan proyek tidak muncul di navigasi maupun pencarian.
  4. Access level diubah lewat Project settings > Access > "Change project access" > pilih level > Change; dialog menampilkan ringkasan dampak (mis. "Semua Creator site akan kehilangan akses edit") sebelum konfirmasi.
  5. Hanya Administrator proyek yang dapat mengubah access level; percobaan oleh role lain mengembalikan HTTP 403.
  6. Pada plan Free access level tidak dapat diubah: pemetaan tetap Creator = Creator + Administrator, Contributor = Contributor; tombol "Change project access" menampilkan pesan "Perubahan access level memerlukan plan Standard atau lebih tinggi".
  7. Perubahan access level berlaku langsung (paling lambat pada request berikutnya) pada view, roadmap, dan published view; cache izin per pengguna diinvalidasi saat perubahan disimpan.
  8. ⚠ Tiga level Open/Limited/Private adalah satu-satunya model access level; proyek yang dibuat sebelum model ini diterapkan dipetakan ke level yang paling mendekati tanpa memperluas akses.
  9. Administrator proyek dapat menambah pengguna atau grup dengan role tertentu per proyek (lihat JPD-INS-003); penambahan grup berlaku dinamis untuk anggota grup yang bergabung kemudian.
  10. Setting izin Contributor membuat idea dikelola per proyek (JPD-INS-004) dan tidak bergantung pada access level.
  11. ⚠ Permission granular per proyek (create ideas, edit views, manage fields) dapat dipisah lewat custom role (JPD-INS-003) sehingga tidak semua Creator otomatis mengelola field.
  12. Contributor tidak dapat mengedit nilai field prioritas (rating impact/effort, select, number, date, formula input) pada idea; field tersebut tampil read-only bagi Contributor, kecuali Votes dan Reactions yang memang untuk Contributor.
  13. Izin proyek diterapkan juga pada published view dan embed: embed di halaman dokumentasi hanya menampilkan konten bila pembaca memiliki akses proyek, sedangkan published view mengikuti pengaturan publish (JPD-INS-029).
- **Team-managed vs company-managed:** Tidak ada perbedaan; access level discovery mirip model access team-managed (open/limited/private) dan tidak memakai permission scheme company-managed.
- **Catatan & ketergantungan:** Project Access Levels (JPD) di modul 36 mengatur sisi pembuatan proyek; modul ini mengatur perilaku izin runtime. Model access team-managed umum di modul 03. Keputusan terbuka: apakah Admin Situs dapat dikecualikan dari proyek Private.

### JPD-INS-003 Space Roles & Permissions

- **Deskripsi:** Role default (Administrator, Creator, Contributor) dan custom role di project discovery beserta permission khusus discovery seperti Manage Views, Publish a view, dan Export a view.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; mengelola role dan permission), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap proyek memiliki tiga role default yang tidak dapat dihapus maupun diganti nama: Administrator, Creator, Contributor.
  2. Administrator memiliki semua hak: mengelola settings dan people, membuat/mengedit field, workflow, dan idea type, serta seluruh hak Creator.
  3. Creator memiliki semua hak kecuali fitur administrasi seperti membuat/mengedit field project dan mengubah pengaturan proyek.
  4. Contributor memiliki akses read-only ditambah membuat idea (bila diizinkan), comment, add insight, vote, dan reaction.
  5. Custom role tidak tersedia di plan Free; tombol "Create role" menampilkan pesan upgrade ke Standard.
  6. Menghapus custom role meminta pemilihan role pengganti atau tetap mempertahankan pemegangnya di proyek dengan role default sesuai access level; pemegang role tidak dikeluarkan dari proyek.
  7. Role default bagi pengguna yang tidak ditambahkan eksplisit mengikuti access level: Open → Creator untuk pemegang role produk Creator, Limited → Contributor, Private → hanya Admin Situs dan pengguna yang ditambahkan.
  8. Permission yang dapat dipilih pada custom role minimal: Manage Views, Publish a view, Publish a view to anyone with the link, Export a view; masing-masing dapat diaktifkan/nonaktifkan per role.
  9. ⚠ Permission tambahan pada custom role mencakup merge idea, archive idea, mengelola insight milik orang lain, dan delete insight; default-nya diberikan ke Administrator dan Creator.
  10. Hanya Administrator proyek yang dapat membuat, mengubah, atau menghapus role dan menugaskan role; perubahan dicatat di audit log (ADM-014).
  11. ⚠ Hak membuat/mengedit field project secara default hanya untuk Administrator; konfigurasi field yang dipicu dari fitur lain (mis. mengaktifkan Autofill dates, JPD-INS-021) mengikuti aturan yang sama sehingga tidak ada jalur bagi Creator non-admin untuk mengubah definisi field.
- **Team-managed vs company-managed:** Tidak ada perbedaan; role discovery tidak memakai project role company-managed.
- **Catatan & ketergantungan:** Plan gating di JPD-INS-008 dan PERM-005 ([03](03-izin-peran-keamanan.md)). Dokumentasi rujukan tidak konsisten soal siapa yang boleh membuat field (Administrator vs Creator); produk ini memilih Administrator sebagai default dan mencatatnya sebagai keputusan terbuka.

### JPD-INS-004 Contributor Idea Submission

- **Deskripsi:** Mengizinkan Contributor membuat idea baru melalui form terkendali sehingga proyek dapat menjadi saluran pengajuan ide internal.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; mengaktifkan dan mengatur form), Contributor (mengajukan idea), Creator (meninjau).
- **Aturan perilaku (acceptance criteria):**
  1. Secara default Contributor tidak dapat membuat idea; tombol Create tidak tampil baginya.
  2. Administrator mengaktifkan lewat Project settings > Features > Create ideas > "Allow contributors to create ideas".
  3. Saat diaktifkan, Administrator memilih apakah Contributor boleh membuat semua idea type atau idea type tertentu saja; daftar type di form Contributor dibatasi sesuai pilihan.
  4. Idea form untuk Contributor dapat dikustomisasi Administrator dengan memilih view dan template (field yang tampil, description template); field di luar form tidak dapat diisi Contributor.
  5. Contributor membuat idea lewat tombol Create di navigasi proyek; form tervalidasi (summary wajib, field wajib form) sebelum dikirim.
  6. Idea yang diajukan Contributor masuk dengan status awal workflow (mis. Parking lot) dan Contributor tercatat sebagai reporter serta watcher otomatis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Idea Creation Settings dan Idea Templates (Description Templates) di modul 36; template di JPD-INS-042. Contributor tidak dapat mengedit idea setelah dibuat (JPD-INS-001).

### JPD-INS-005 View Permissions (View Restrictions)

- **Deskripsi:** Tingkat akses per view untuk mengatur siapa yang dapat melihat dan mengedit view tertentu di dalam proyek (Premium).
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator dengan permission Manage Views (mengatur restriction), Admin Proyek/Project Lead (role Administrator).
- **Aturan perilaku (acceptance criteria):**
  1. Open: semua anggota proyek dapat melihat view; Creator dengan Manage Views dapat mengedit.
  2. Open, with selected editors: semua anggota melihat; hanya pengguna/grup terpilih yang dapat mengedit.
  3. Restricted: hanya pengguna/grup terpilih yang dapat melihat dan/atau mengedit; view tidak muncul di sidebar bagi pengguna lain dan URL-nya mengembalikan HTTP 404.
  4. Pengaturan dibuka lewat ikon gembok di header view > pilih level > Additional access > tambah pengguna/grup dengan opsi "can view" atau "can edit".
  5. Pengguna dengan hak view saja tidak dapat rename/delete view, mengubah title/description/emoji, grouping, field, settings, atau permission; filter dan sort yang ia ubah bersifat sementara dan di-reset saat halaman di-refresh.
  6. Contributor, Creator tanpa permission Manage Views, dan pengguna tanpa app access tidak dapat mengedit view meski tercantum sebagai editor; UI menandai entri tersebut dengan peringatan "Pengguna ini tidak memiliki hak edit".
  7. Global view (view lintas proyek/roadmap) tidak dapat diedit dari dalam proyek; tautan "Edit di roadmap" diarahkan ke JPD-INS-006.
  8. Saat upgrade ke Premium semua view awal bernilai Open; saat downgrade, view Restricted menjadi terlihat semua anggota proyek dan restriction disimpan sehingga dipulihkan otomatis bila upgrade ulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** View Description & Ownership dan Views Navigation & Sections di modul 36; plan di JPD-INS-008.

### JPD-INS-006 Roadmaps Permissions (Cross-space Roadmaps)

- **Deskripsi:** Permission untuk Roadmaps, yaitu view yang menggabungkan idea dari banyak proyek (Premium), termasuk kepemilikan, level akses, dan pewarisan izin idea dari proyek asal.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator (pembuat/owner roadmap), Admin Proyek/Project Lead (role Administrator), pengguna berlisensi pelacak kerja (viewer).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna berlisensi Creator dapat membuat roadmap baru dan melihat roadmap lain yang level aksesnya mengizinkan.
  2. Pengguna berlisensi pelacak kerja (tanpa lisensi Creator) dapat melihat roadmap yang Open secara read-only.
  3. Owner roadmap dapat mengedit nama, emoji, dan description, menambah/menghapus proyek sumber, membuat view di roadmap, serta mengubah permission; owner dapat dialihkan ke Creator lain.
  4. ⚠ Level akses roadmap adalah Private, Restricted, dan Open, dengan default Restricted (hanya owner dan pengguna yang ditambahkan).
  5. Izin atas setiap idea di roadmap mengikuti proyek asalnya: pengguna dapat mengedit idea di roadmap hanya bila ia Creator di proyek asal idea.
  6. Pengguna tanpa akses ke suatu proyek tidak melihat idea dari proyek tersebut di roadmap; jumlah idea dan agregat di roadmap dihitung hanya dari idea yang terlihat olehnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Cross-project Roadmaps di modul 36 (struktur dan view roadmap); izin Plans umum di PERM-052 ([03](03-izin-peran-keamanan.md)).

### JPD-INS-007 Product Discovery global settings

- **Alias inventaris:** Jira Product Discovery settings (lampiran)
- **Deskripsi:** Halaman pengaturan tingkat site untuk Product Discovery yang dikelola Admin Situs: akses Creator/Contributor, global fields, siapa yang boleh membuat project discovery, integrasi delivery, serta pengaturan publish dan stakeholder access.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs (satu-satunya yang dapat membuka halaman ini).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Apps > Product Discovery memuat tab Global fields untuk membuat, mengedit, dan menghapus global field yang tersedia di semua project discovery (rincian perilaku field di Global Fields modul 36).
  2. Admin Situs menentukan siapa yang boleh membuat project discovery: semua Creator, grup tertentu, atau hanya admin; pengguna di luar daftar tidak melihat template discovery saat membuat proyek.
  3. ⚠ Tab Access memuat toggle "Allow publishing views" dan "Allow publishing to anyone with the link" tingkat site; keduanya default off dan menjadi prasyarat JPD-INS-029.
  4. ⚠ Admin Situs dapat menetapkan default site untuk Contributor, misalnya mengaktifkan grup Contributor dan mematikan pembuatan idea oleh Contributor secara global sehingga setting per proyek tidak dapat mengaktifkannya.
  5. Tab Access menampilkan jumlah Creator dan Contributor serta tautan ke pengelolaan product access (JPD-INS-001).
  6. Global fields tidak dapat dihapus selama masih dipakai formula atau view; dialog hapus menampilkan daftar proyek dan view yang memakainya.
  7. Pengaturan integrasi delivery memungkinkan Admin Situs mengaktifkan/menonaktifkan penautan idea ke work item di proyek lain dan melihat daftar application link lintas site (JPD-INS-017).
  8. ⚠ Admin Situs dapat mengatur default publish view dan roadmap (mis. mewajibkan login untuk semua published view) yang berlaku sebagai batas atas bagi setting proyek.
  9. Halaman pengaturan Product Discovery memuat bagian Dates (padanan rujukan Products > Product Discovery > Dates) tempat Admin Situs memilih bulan awal tahun fiskal (default Januari); pilihan ini berlaku untuk seluruh project discovery di site, menentukan batas kuartal dan tahun pada field tanggal/interval berpresisi kuartal (JPD-IDE-033, [36](36-jpd-ideas-field-views.md)) serta label kuartal di timeline dan roadmap, dan perubahan hanya menggeser tampilan/pengelompokan tanpa mengubah nilai tanggal yang tersimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman administrasi produk lain di modul 22; application link di ADM-032 ([22](22-administrasi-sistem.md)). Nama dan isi menu mengikuti struktur navigasi admin produk ini, bukan salinan menu rujukan.

### JPD-INS-008 JPD Plans & Limits (Free / Standard / Premium / Enterprise)

- **Alias inventaris:** JPD licensing and product access
- **Deskripsi:** Plan Product Discovery (Free, Standard, Premium, Enterprise) dengan perhitungan seat hanya untuk Creator, batas objek, dan fitur yang dibuka per plan, serta pengelolaan product access di admin organisasi.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Org admin dan Admin Situs (mengelola plan, billing, product access).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Free membatasi jumlah Creator hingga 3, tanpa kemampuan mengubah access level dan tanpa custom role; menambah Creator ke-4 ditolak dengan pesan "Batas Creator plan Free tercapai. Upgrade ke Standard".
  2. Standard ditagih per Creator dan membuka perubahan access level, custom role, serta publish view.
  3. Premium membuka Roadmaps lintas proyek, view restrictions, idea hierarchies, penyimpanan tanpa batas, AI (JPD-INS-041), admin insights, sandbox, dukungan 24/7, kuota otomasi lebih tinggi, SLA uptime, IP allowlisting, release tracks, dan integrasi instance self-hosted (JPD-INS-024).
  4. Enterprise menambah SLA 99,95%, sandbox, multi-site, langganan keamanan organisasi, audit log organisasi, BYOK, data residency, dan Analytics/Data Lake (lihat JPD-INS-040).
  5. Jumlah Contributor tidak dibatasi di semua plan dan tidak dihitung dalam billing.
  6. ⚠ Batas objek per plan ditegakkan dan dapat dikonfigurasi: jumlah idea per proyek, view per proyek, dan field per proyek; saat batas tercapai, aksi create ditolak dengan pesan yang menyebut batas dan plan.
  7. Downgrade menghilangkan fitur plan lebih tinggi tanpa menghapus data: restriction view dinonaktifkan (JPD-INS-005), roadmap menjadi read-only, dan konfigurasi dipulihkan saat upgrade ulang.
  8. ⚠ Free memiliki batas jumlah Creator yang kecil (default 3), dan nilai batas dapat diubah oleh operator produk melalui konfigurasi tanpa rilis ulang.
  9. Hanya Creator yang dihitung sebagai seat berbayar; Contributor dan Stakeholder tidak dihitung dalam tagihan maupun kuota seat.
  10. Product access (pemberian/pencabutan role Creator, grup default) dikelola di halaman admin organisasi (USR-018), bukan di dalam proyek.
  11. ⚠ Fitur Premium seperti Roadmaps lintas proyek dan global fields ditandai badge "Premium" di UI dan dikunci pada plan lebih rendah dengan pesan upgrade.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Billable users di USR-020 ([02](02-pengguna-grup-organisasi.md)); batas usage umum ADM-042 ([22](22-administrasi-sistem.md)); kuota otomasi AUTO-015 ([20](20-otomasi.md)). Tanggal GA plan Premium di produk rujukan tidak relevan bagi produk ini; tier yang sama dipetakan ke kebijakan harga produk ini (keputusan terbuka). Migrasi: project discovery beserta idea, insight, view, field, dan tautan delivery ikut dalam migrasi cloud-to-cloud antar-site (ADM-039 di [22](22-administrasi-sistem.md), INT-025 di [25](25-integrasi-impor-ekspor-ekstensi.md)); role produk Creator di site tujuan harus tersedia dalam kuota plan sebelum migrasi dijalankan.

## Insights: penangkapan, pengelolaan & skor bukti

### JPD-INS-009 Insights Panel (Idea Insights Tab)

- **Alias inventaris:** Insights
- **Deskripsi:** Tab/panel Insights di idea view untuk mengumpulkan bukti (feedback pelanggan, riset, tiket support, peluang sales, pesan stakeholder) yang mendukung idea. Insight adalah entitas tersendiri (bukan komentar) dan setiap insight dari jalur mana pun (editor, tempel tautan, ekstensi, chat, data provider) masuk ke tab Insights idea tujuan.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (membuat insight; mengedit/menghapus insight miliknya), Admin Proyek/Project Lead (role Administrator; mengelola insight siapa pun), viewer proyek (membaca).
- **Aturan perilaku (acceptance criteria):**
  1. Tab "Insights" tersedia di sidebar/detail idea dan juga terbuka saat pengguna mengklik field Insights di view List, Board, Matrix, maupun Timeline.
  2. Tombol "Create an insight" membuka editor insight inline di bagian atas daftar insight.
  3. Field Description wajib diisi dan berupa rich text (format dokumen editor modul 07) yang mendukung mention @user, emoji, dan gambar; insight dengan description kosong ditolak dengan pesan "Description insight wajib diisi".
  4. Field opsional "Paste a link" menerima URL web apa pun atau tautan produk yang didukung (work item, request Service Management, halaman dokumentasi, Zendesk, Salesforce) yang dirender sebagai Smart Link/embedded object (JPD-INS-010).
  5. Insight tanpa tautan (teks saja) diperbolehkan dan tampil tanpa kartu sumber.
  6. Impact rating diberikan dengan hover pada deretan titik impact lalu mengklik level (mis. 1-5) untuk memberi bobot signifikansi insight; impact dapat dikosongkan.
  7. Label opsional pada insight (mis. "VIP customer", "Key customer") dipilih dari daftar label proyek atau dibuat baru dari editor.
  8. ⚠ Insight dapat ditandai bernilai positif atau negatif terhadap idea (mendukung vs menentang), dan tanda negatif dihitung terpisah dalam skor insight (JPD-INS-012).
  9. Edit insight dilakukan dengan hover lalu klik ikon pensil; perubahan description, link, label, dan impact disimpan dengan penanda "diedit" dan waktu edit.
  10. Setiap insight memiliki checkbox untuk memilih satu atau banyak insight sebagai target aksi bulk (Copy to idea, Delete); toolbar bulk menampilkan jumlah terpilih.
  11. Delete insight dilakukan dengan memilih checkbox lalu klik Delete dan konfirmasi; penghapusan permanen, tidak dapat dipulihkan, dan dicatat di history idea.
  12. Daftar insight di tab dapat difilter (label, impact, author, sumber) dan di-sort (terbaru, terlama, impact tertinggi).
  13. ⚠ Pengguna dapat memberi reaction emoji pada insight; jumlah per emoji tampil di bawah insight.
  14. Metadata otomatis setiap insight: author, tanggal dibuat, dan sumber/URL (termasuk jalur pembuatan: editor, ekstensi, Slack, Teams, data provider).
  15. Snippet/kutipan teks hasil highlight dari ekstensi browser (JPD-INS-014) tampil sebagai blok kutipan terpisah dari description insight.
  16. Screenshot yang ditangkap via ekstensi browser tersimpan sebagai lampiran insight dan tampil sebagai thumbnail yang dapat diperbesar.
  17. Empty state tab menampilkan penjelasan singkat insight, tautan pemasangan ekstensi browser, dan petunjuk penangkapan dari Slack/Teams.
  18. Menambah, menghapus, atau menyalin insight langsung memperbarui field jumlah Insights pada idea (JPD-INS-012) tanpa reload halaman.
  19. ⚠ Saat idea di-merge ke idea target, seluruh insight idea sumber dipindahkan ke idea target dengan metadata asli dan penanda "dipindah dari <key>".
  20. ⚠ Tidak ada feed insight lintas idea dalam satu proyek pada versi awal; alternatifnya pengguna men-sort view berdasarkan field Insights.
  21. Contributor dapat menambah insight tanpa lisensi Creator; Contributor hanya dapat mengedit/menghapus insight miliknya sendiri.
  22. Pengguna dapat melampirkan screenshot atau gambar langsung (upload/drag-drop/paste dari clipboard) ke insight, dengan batas ukuran lampiran yang sama seperti lampiran work item (modul 08).
  23. Jumlah insight menjadi dasar sort dan filter idea di view (rinci di JPD-INS-012).
  24. Insight (jumlah dan impact) dapat dipakai sebagai input weighted score dan roll-up formula (Formula Fields modul 36).
  25. ⚠ Atribut per insight (label, impact) dapat dikonfigurasi per proyek: daftar label dikelola Administrator dan skala impact dapat diubah.
  26. ⚠ Pengguna dapat membuat insight dari comment idea lewat menu comment "Buat insight"; teks comment menjadi description dan author asli dicatat sebagai sumber.
  27. ⚠ Insight dapat dipindahkan ke idea lain (bukan disalin): insight hilang dari idea sumber, hitungan kedua idea diperbarui, dan perpindahan dicatat di history kedua idea.
  28. Tombol "Create an insight" di tab Insights tampil bagi setiap pengguna yang berhak menambah insight (Creator, Contributor) dan disembunyikan bagi pembaca tanpa hak, termasuk pembaca published view.
  29. Satu insight dapat memuat sekaligus description/snippet teks dan satu link sumber; link dirender sebagai Smart Link preview di atas description.
  30. Menempel URL langsung ke field Description insight otomatis mengubahnya menjadi smart card inline tanpa langkah tambahan; pengguna dapat mengembalikannya ke URL biasa lewat menu tampilan tautan.
  31. Tombol Edit dan Delete insight hanya tampil bagi penulis insight dan Administrator proyek; API menolak permintaan pengguna lain dengan HTTP 403.
  32. Metadata "dibuat oleh" dan "waktu dibuat" bersifat read-only dan tidak dapat diubah lewat UI maupun API.
  33. Field jumlah Insights tersedia otomatis di setiap project discovery sebagai field sistem tanpa perlu dibuat Administrator.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama legacy insight adalah "Data points" (JPD-INS-044). Idea View dan Merge Ideas di modul 36; editor rich text di modul 07; lampiran di modul 08. Insight disimpan di tabel tersendiri (`insights`, `insight_labels`, `insight_snippets`) berelasi ke work item idea, bukan sebagai comment.

### JPD-INS-010 Insight Capture via Paste Link (Smart Links)

- **Alias inventaris:** Smart Links and third-party previews
- **Deskripsi:** Membuat insight dengan menempelkan URL: sistem mengambil data sumber melalui resolver Smart Link dan merendernya sebagai Smart Link/embedded object yang tampil live. Preview kaya yang sama juga berlaku untuk tautan di description idea.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor (penulis insight), viewer (pembaca preview sesuai izin sumber).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna menempel URL di field "Paste a link" pada editor insight; preview muncul sebelum insight disimpan.
  2. Resolver produk ini mendukung minimal: work item (key, summary, status, type, assignee), request Service Management, halaman dokumentasi, video, board kanban ringan, dan repositori kode bila resolvernya terpasang.
  3. Resolver pihak ketiga yang didukung registry Smart Link platform (mis. Google Drive, Figma, YouTube) dipakai tanpa konfigurasi tambahan di project discovery.
  4. Bila resolver tidak tersedia atau pengguna belum terautentikasi ke sumber, tautan ditampilkan sebagai link biasa (judul dari OpenGraph bila ada) dan insight tetap tersimpan.
  5. Untuk sumber ber-OAuth (mis. Salesforce) preview menampilkan tombol "Connect account"; setelah otorisasi, preview diperbarui tanpa membuat ulang insight.
  6. Konten Smart Link menghormati izin sumber: pengguna tanpa akses ke sumber hanya melihat URL/judul terbatas, tidak melihat isi seperti status atau field.
  7. Data preview diambil live saat insight dilihat (dengan cache singkat), bukan salinan statis yang dibekukan saat insight dibuat.
  8. URL divalidasi berformat http/https; skema lain (mis. `javascript:`, `file:`) ditolak dengan pesan "Tautan harus berupa URL http atau https".
  9. Konektor preview yang tersedia mengikuti registry (mis. Figma, Miro, Google Docs/Sheets, Loom, Dovetail); daftar konektor aktif tampil di halaman pengaturan Smart Link admin.
  10. Tautan di insight dan description idea dapat ditampilkan dalam mode inline, card, atau embed; mode embed hanya ditawarkan untuk penyedia yang mendukung iframe.
  11. Autentikasi ke penyedia yang memerlukan login dilakukan sekali per pengguna per penyedia dan berlaku untuk semua insight serta description.
  12. Penyedia yang tidak dikenal selalu jatuh ke URL biasa yang dapat diklik tanpa galat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Registry resolver dan koneksi akun di INT-002 ([25](25-integrasi-impor-ekspor-ekstensi.md)); rendering mode di EDT-021 ([07](07-editor-konten-kaya.md)) dan NAV-017 ([21](21-navigasi-ux-preferensi.md)); perlindungan SSRF di INT-001. Resolver untuk layanan lintas produk vendor (dokumentasi, video, repositori milik vendor rujukan) tidak dibangun; padanannya adalah modul produk ini sendiri bila ada.

### JPD-INS-011 Copy Insights to Existing Ideas

- **Deskripsi:** Menyalin insight yang sama ke satu atau beberapa idea lain sehingga satu bukti dapat mendukung banyak idea, dengan opsi membuat idea baru sebagai target.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor (dengan akses ke idea target).
- **Aturan perilaku (acceptance criteria):**
  1. Dari idea sumber > Insights > centang insight > "Copy to idea" membuka dialog pemilihan target.
  2. Dialog menyediakan pencarian idea yang sudah ada (summary/key, lintas proyek yang dapat diakses) atau mengetik nama untuk membuat idea baru, lalu tombol Copy.
  3. Insight ditambahkan ke idea target sebagai salinan dengan metadata author dan tanggal asli; insight sumber tetap ada dan tidak berubah.
  4. Banyak insight dapat dipilih sekaligus lewat checkbox dan disalin dalam satu operasi; hasil menampilkan "N insight disalin ke <idea>".
  5. ⚠ Semua jenis insight (teks biasa maupun insight dari data provider seperti Zendesk atau Salesforce) dapat disalin; insight dari data provider mempertahankan konfigurasi properties-nya.
  6. Penyalinan menambah hitungan Insights di idea target dan memicu event "insight added" pada idea target.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mengoreksi asumsi lama bahwa insight hanya milik satu idea; salinan adalah baris insight terpisah, bukan referensi bersama. Pemindahan insight di JPD-INS-009.

### JPD-INS-012 Insights Count & Insight-based Scoring Fields

- **Deskripsi:** Field dinamis Insights yang menampilkan jumlah insight per idea serta agregasi impact insight untuk dipakai di view dan formula prioritisasi.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Admin Proyek/Project Lead (role Administrator; memakai field di view dan formula), semua viewer proyek (melihat nilai).
- **Aturan perilaku (acceptance criteria):**
  1. Field dinamis "Insights" menampilkan jumlah insight idea; mengklik nilainya membuka sidebar tab Insights untuk melihat atau membuat insight.
  2. Field bersifat read-only dan diperbarui real-time (melalui kanal realtime modul 27) saat insight ditambah, dihapus, atau disalin.
  3. View dapat di-sort, difilter (mis. "Insights ≥ 5"), dan di-group berdasarkan jumlah insight.
  4. ⚠ Tersedia field "Insight impact score" yang mengagregasi impact semua insight idea (jumlah nilai impact) untuk membandingkan idea.
  5. Jumlah insight dan impact score dapat menjadi input Custom formula: Roll-up, Weighted Score, atau expression (mis. `Impact * Insights / Effort`).
  6. Field turunan data provider (mis. jumlah tiket support terkait idea, JPD-INS-013) dapat dipakai sebagai input prioritisasi dengan cara yang sama.
  7. Field "Number of insights" tidak dapat dipakai untuk sort, filter, group, kolom board, atau sumbu matrix pada published view; publish ditolak dengan pesan yang menyebut field tersebut (JPD-INS-029).
  8. Ekspor CSV hanya menyertakan jumlah insight, bukan isi insight (JPD-INS-039).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Insights & Delivery Computed Fields dan Formula Fields di modul 36; realtime update di modul 27.

### JPD-INS-013 Insight Data Providers (Salesforce, Zendesk, Jira Service Management)

- **Alias inventaris:** Salesforce integration (JPD)
- **Deskripsi:** Integrasi berbasis tautan untuk menarik data terstruktur dari tiket support dan objek sales ke insight sehingga dampak idea dapat dikuantifikasi (mis. nilai opportunity, jumlah tiket).
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor (menambah insight), Creator/Administrator (membuat field turunan); pengguna wajib punya akun di sistem sumber.
- **Aturan perilaku (acceptance criteria):**
  1. Salesforce: pengguna menyalin tautan object (mis. Opportunity, Account), menempelkannya di insight, melakukan autentikasi bila diminta, hover embedded object, memilih properties yang ditampilkan, klik Save lalu Create.
  2. Pilihan properties Salesforce diingat per tipe object: insight berikutnya dari object bertipe sama hanya menampilkan properties yang dipilih pertama kali, dan harus dikonfigurasi ulang untuk menampilkan field lain.
  3. Tautan dari instance sandbox Salesforce tidak didukung; preview menampilkan pesan "Tautan sandbox tidak didukung" dan insight tersimpan sebagai link biasa.
  4. Zendesk: pengguna dapat menambahkan ticket maupun view/filter Zendesk ke idea sebagai insight untuk melacak masalah penting.
  5. Service Management: pengguna dapat menambahkan request sebagai insight untuk konteks masalah pengguna (JPD-INS-035).
  6. Label/properties pada insight dari Service Management dapat dikonfigurasi (hover label > pilih properties yang dipertahankan); properties yang di-uncheck lalu disimpan tidak dapat dipulihkan untuk insight tersebut.
  7. Data provider dapat dipakai untuk membuat field prioritisasi (mis. jumlah tiket support yang menyebut idea, total nilai opportunity) dan custom formula; membuat field ini memerlukan role Creator/Administrator.
  8. Pengguna memerlukan akun di sistem sumber serta app access ke produk ini; tanpa akun sumber, insight hanya menampilkan tautan.
  9. ⚠ Provider tambahan (analitik produk, feedback in-app, percakapan support, pre-sales) diperlakukan sebagai resolver opsional di registry Smart Link; provider yang hanya tersedia sebagai app pihak ketiga tidak dibangun native.
  10. ⚠ Penambahan insight dari record Salesforce mendukung object Account, Opportunity, dan Case; bentuk integrasi adalah resolver tautan, bukan app terpasang di sisi Salesforce.
  11. Record Salesforce di insight tampil sebagai Smart Link card dengan nama record dan properties terpilih.
  12. ⚠ Atribut pelanggan dari Salesforce (nama account, revenue/ARR) disimpan sebagai nilai terstruktur di insight sehingga dapat dirujuk field.
  13. ⚠ Nilai numerik dari Salesforce (mis. amount opportunity) dapat dipakai dalam roll-up atau weighted score formula.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pola integrasi helpdesk/CRM di INT-004 ([25](25-integrasi-impor-ekspor-ekstensi.md)); paste link di JPD-INS-010. Provider yang dalam rujukan berupa app Marketplace (mis. widget feedback pihak ketiga) di luar cakupan.

### JPD-INS-014 Jira Product Discovery Chrome Extension

- **Alias inventaris:** Chrome extension (JPD)
- **Deskripsi:** Ekstensi browser resmi produk ini untuk menangkap teks dan screenshot dari halaman web atau web app mana pun (termasuk halaman dokumentasi, Service Management, Zendesk, Salesforce) lalu menambahkannya sebagai insight ke idea.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor dengan akses ke proyek tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Ekstensi dipasang dari tautan "Browser extension" di tab Insights atau dari Chrome Web Store ("Add to Chrome"); setelah pasang, pengguna diarahkan untuk mem-pin ikon di toolbar.
  2. Pengguna login dengan akun produk ini (OAuth ke instance) dan memilih site/instance tujuan; token disimpan di storage ekstensi dan dapat dicabut dari Connected apps (AUTH-013).
  3. Highlight teks > klik kanan > "Add to idea" membuka panel ekstensi berisi teks terpilih.
  4. Ekstensi memiliki antrean (queue): ikon menampilkan badge jumlah item yang menunggu dikirim, dan item dalam antrean dapat di-bulk edit (target idea, label).
  5. Item dapat dihapus dari antrean dengan ikon silang di bawah tombol "Add to idea".
  6. "Select more text" memungkinkan menambah potongan teks berulang ke item yang sama sebelum dikirim.
  7. Ikon kamera di editor ekstensi menangkap screenshot area terlihat halaman dan melampirkannya ke item.
  8. Pengguna memilih project discovery dan idea tujuan (pencarian), atau membuat idea baru langsung dari ekstensi.
  9. Pengguna dapat menambahkan description dan label opsional sebelum menekan Add.
  10. Teks terpilih menjadi snippet insight dan URL halaman otomatis ditautkan sebagai sumber insight.
  11. Ekstensi dapat dipakai untuk menangkap request Service Management, catatan interview di halaman dokumentasi, dan object CRM sebagai insight.
  12. ⚠ Browser yang didukung resmi adalah Chrome; browser berbasis Chromium lain (mis. Edge) berjalan dengan paket yang sama namun tidak dijamin.
  13. Hanya pengguna dengan akses proyek (Creator/Contributor) yang dapat menambah insight; proyek tanpa akses tidak muncul di pemilih, dan API capture mengembalikan 403 bila akses dicabut.
  14. ⚠ Screenshot halaman yang ditangkap ekstensi tersimpan sebagai bagian insight (gambar lampiran) dan tampil di tab Insights (JPD-INS-009).
  15. Ekstensi dapat dipasang oleh pengguna mana pun tanpa hak admin site; pemasangan pertama menampilkan onboarding singkat (login, pilih site, cara highlight).
  16. Teks terseleksi disimpan sebagai snippet teks polos tanpa markup HTML; tanpa seleksi teks, menu "Add to idea" tetap membuka panel dengan snippet kosong agar pengguna dapat menulis description sendiri.
  17. URL dan judul halaman sumber disimpan otomatis; halaman internal browser (mis. `chrome://`) atau file lokal tidak dapat ditangkap dan ekstensi menampilkan "Halaman ini tidak dapat ditangkap".
  18. Saat membuat idea baru dari ekstensi, pengguna memilih proyek dan mengisi summary; insight langsung tertaut ke idea baru tersebut.
  19. Ekstensi mengingat site terakhir yang dipilih dan menyediakan pengalih site; logout menghapus token dan antrean lokal yang belum terkirim setelah konfirmasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Membutuhkan endpoint capture `/api/v1/discovery/insights/capture` dengan OAuth 2.0 (modul 01 dan 23); kebijakan CORS dan CSP untuk origin ekstensi (modul 27). Ekstensi asisten AI lintas produk (NAV-053) adalah fitur terpisah.

### JPD-INS-015 Zendesk integration (JPD)

- **Deskripsi:** Menautkan ticket Zendesk sebagai insight ke idea dari dalam Zendesk melalui app sidebar yang diterbitkan produk ini untuk Zendesk.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Agent support di Zendesk (dengan akun produk ini sebagai Creator/Contributor), Creator, Contributor.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Zendesk memasang app produk ini di Zendesk dan mengisi URL instance; setiap agent menghubungkan akun produk ini lewat OAuth saat pertama kali memakai app.
  2. Panel sidebar ticket Zendesk menyediakan pencarian idea (proyek yang dapat diakses agent) dan tombol "Add insight".
  3. Agent dapat memilih snippet dari percakapan ticket sebagai isi insight; bila tidak memilih, description wajib diisi manual.
  4. Insight yang dibuat memuat Smart Link ke ticket Zendesk (nomor, subjek, status) sebagai sumber.
  5. Panel sidebar menampilkan daftar idea yang sudah terhubung ke ticket beserta status idea.
  6. Agent dapat membuat idea baru dari Zendesk (proyek, idea type, summary) yang langsung berisi insight dari ticket tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App sidebar Zendesk adalah artefak resmi produk ini yang dipublikasikan di marketplace Zendesk; kebutuhan endpoint sama dengan JPD-INS-014. Pola umum di INT-004 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JPD-INS-016 Intercom integration (JPD)

- **Deskripsi:** Menangkap percakapan Intercom sebagai insight pada idea.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Agent support di Intercom, Creator, Contributor.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Integrasi dipasang sebagai app di Intercom (bukan hanya via ekstensi browser); admin Intercom menghubungkan instance produk ini dan agent mengautentikasi akunnya.
  2. Inbox Intercom menyediakan aksi "Add insight" pada percakapan.
  3. Insight memuat snippet percakapan dan tautan ke conversation Intercom sebagai sumber.
  4. Dari aksi tersebut agent dapat mencari idea yang ada atau membuat idea baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bila app Intercom tidak dibangun, jalur pengganti adalah ekstensi browser (JPD-INS-014) dan paste link (JPD-INS-010); keputusan dicatat di Pertanyaan terbuka.

## Delivery tracking: idea ke work item eksekusi

### JPD-INS-017 Delivery Tab (Link Ideas to Jira Work Items)

- **Alias inventaris:** Delivery tracking
- **Deskripsi:** Tab Delivery di idea view untuk menghubungkan idea ke work item pelacak kerja (Epic/Story/Task) sebagai jejak eksekusi, lintas proyek dan lintas site, lengkap dengan progress bar dan tabel work item tertaut.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Admin Proyek/Project Lead (role Administrator; menautkan/membuat/melepas), Admin Situs (application link lintas site), Developer/Anggota Tim (melihat dari sisi work item); permission Create work items dan Link work items di proyek target wajib dimiliki.
- **Aturan perilaku (acceptance criteria):**
  1. Tab "Delivery" di panel kanan idea menampilkan progress bar (JPD-INS-018) dan tabel linked work items.
  2. "Link a work item": pengguna memilih site target, proyek, lalu mencari work item (key/summary) > Add; work item yang sudah tertaut tidak muncul lagi di hasil pencarian.
  3. "+ Create work item": pengguna memilih site, proyek, dan work type (mis. Epic); nama terisi dari summary idea dan dapat diedit; checkbox opsional menyalin description dan nilai field idea ke work item; klik Create membuat dan langsung menautkan work item.
  4. Unlink dilakukan lewat hover baris > ikon silang dengan konfirmasi; work item yang dilepas tidak lagi dihitung dalam progress maupun status.
  5. Tabel linked items dapat dikonfigurasi hingga 10 kolom dari daftar: Type, Key, Status, Story points, Assignee, Due date, Sprint, Fix version, Team, Time tracking; konfigurasi kolom berlaku per proyek.
  6. Penautan didukung ke proyek pelacak kerja company-managed maupun team-managed dan ke site lain produk ini.
  7. Penautan lintas site memerlukan application link OAuth 2.0 yang dibuat Admin Situs; setiap pengguna harus mengautentikasi pada pemakaian pertama, dan sebelum itu tabel menampilkan "Hubungkan akun untuk melihat work item di <site>" tanpa progress terbaru.
  8. Prasyarat: minimal satu proyek pelacak kerja ada; pengguna harus Creator/Administrator di project discovery dan memiliki permission Create work items serta Link work items di proyek target; bila tidak, tombol Create/Link dinonaktifkan dengan tooltip penyebabnya.
  9. Tautan memakai link type sistem khusus delivery dengan nama tampilan "implements / is implemented by" (nama internal kompatibilitas "Polaris work item link", JPD-INS-044); link type ini tidak dapat dihapus admin.
  10. Satu idea dapat memiliki banyak work item delivery dari beberapa proyek dan site sekaligus.
  11. Idea yang tertaut tampil di dalam work item pelacak kerja melalui panel idea (JPD-INS-022).
  12. Field sistem "Linked work items" menampilkan jumlah work item tertaut; field ini tidak dapat dipakai untuk sort, filter, atau group di published view.
  13. Menghapus link tidak menghapus item di sisi mana pun; menghapus idea menghapus tautan delivery-nya, sedangkan tautan ke instance eksternal yang tidak dapat dihubungi ditandai "menggantung" dan dibersihkan saat koneksi pulih.
  14. ⚠ Nama link type delivery yang tampil di UI pelacak kerja adalah "implements" (outward) dan "is implemented by" (inward), dan dapat dipakai pada JQL `issueLinkType`.
  15. Progress bar menghitung jumlah work item To do / In progress / Done termasuk child issue sesuai konfigurasi (JPD-INS-018).
  16. Field agregat Delivery status tersedia untuk ringkasan status eksekusi idea (JPD-INS-019).
  17. ⚠ Pembuatan work item delivery secara massal untuk banyak idea tersedia dari list view (JPD-INS-020).
  18. Pencarian pada "Link a work item" mencakup semua proyek pelacak kerja yang dapat diakses pengguna di site terpilih, dengan filter per proyek dan pencarian berdasarkan key maupun teks summary.
  19. Pada "+ Create work item", summary terisi otomatis dari summary idea; bila summary dikosongkan, tombol Create dinonaktifkan dengan pesan "Summary wajib diisi".
  20. Unlink tercatat di history idea (aktor, waktu, work item) dan Delivery progress serta Delivery status dihitung ulang segera setelah unlink.
  21. Panel idea terkait tampil otomatis di work item view setiap work item delivery yang tertaut (JPD-INS-022) tanpa konfigurasi tambahan di proyek pelacak kerja.
  22. Bila idea memiliki delivery dari beberapa proyek atau site, tabel linked items dapat dikelompokkan per proyek/site dan menampilkan nama proyek pada setiap baris.
- **Team-managed vs company-managed:** Proyek target team-managed maupun company-managed didukung; perbedaan hanya pada field estimasi yang dihitung (JPD-INS-018) dan ketersediaan field Team.
- **Catatan & ketergantungan:** Issue links di ISS-029 ([06](06-work-item-inti.md)) dan CFG-056 ([05](05-konfigurasi-work-item-skema.md)); application link di ADM-032 ([22](22-administrasi-sistem.md)) dan API-014 ([23](23-api-publik-webhook.md)); rujukan INT-057 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JPD-INS-018 Delivery Progress Field

- **Alias inventaris:** Delivery settings
- **Deskripsi:** Field sistem yang merangkum progres work item tertaut sebagai progress bar per status category, dengan konfigurasi metode perhitungan per proyek yang dilakukan dari field atau tab Delivery.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator atau Admin Proyek/Project Lead (role Administrator; mengonfigurasi), Creator dan Contributor (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Progress bar tersegmentasi To do / In progress / Done (warna status category, WF-002) dengan tooltip berisi jumlah atau story points per segmen dan persentase Done.
  2. Konfigurasi dibuka dengan mengklik progress bar (di field view atau tab Delivery) > "Configure delivery progress".
  3. Hanya Creator/Administrator yang dapat mengonfigurasi; konfigurasi berlaku untuk seluruh idea di proyek.
  4. Estimation type Default: menghitung status direct children (satu level) dari setiap work item tertaut, atau total story points item tertaut ditambah direct children; item tanpa child dihitung memakai statusnya sendiri.
  5. Estimation type Multilevel: menghitung semua level keturunan, item top-level yang ditautkan dikecualikan dari hitungan, dan item duplikat (tercapai dari dua jalur) dihitung sekali; untuk item dari instance self-hosted eksternal tidak didukung dan otomatis memakai Default.
  6. Estimation type Simple: hanya item yang ditautkan langsung yang dihitung; child diabaikan.
  7. Estimation value dapat dipilih: Issue count atau Story points.
  8. Story points hanya diambil dari field "Story points" proyek company-managed; field "Story point estimate" team-managed tidak didukung dan item team-managed dihitung 0 poin dengan penanda di tooltip.
  9. Nilai progress dihitung per pengguna sesuai izin terhadap work item, sehingga dua pengguna dapat melihat nilai berbeda; item yang tidak terlihat tidak disebut di tooltip.
  10. Field read-only, dapat ditampilkan dan di-sort di view, dan tidak disertakan pada ekspor Excel/Google Sheets.
  11. ⚠ Metode perhitungan progress (Issue count atau Story points) adalah bagian dari pengaturan delivery proyek dan tersimpan bersama estimation type.
  12. ⚠ Administrator proyek dapat menonaktifkan fitur delivery untuk proyek (Project settings > Features); tab Delivery dan field delivery disembunyikan, tautan yang ada tetap tersimpan.
  13. ⚠ Administrator proyek dapat menetapkan default proyek dan work type untuk "+ Create work item" sehingga dialog terisi otomatis.
- **Team-managed vs company-managed:** Story points hanya dihitung dari field company-managed (lihat aturan 8).
- **Catatan & ketergantungan:** Konfigurasi ada di field/tab Delivery, bukan halaman pengaturan terpisah; Story Points vs Story point estimate di CFG-016 ([05](05-konfigurasi-work-item-skema.md)) dan TIME-008 ([14](14-pelacakan-waktu-estimasi.md)). Perhitungan dilakukan worker yang mendengarkan event status/hierarki work item (modul 28).

### JPD-INS-019 Delivery Status Field

- **Deskripsi:** Field sistem yang menampilkan status ringkas eksekusi idea berdasarkan jumlah work item atau story points per status work item tertaut.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor (melihat); Creator (memakai di view).
- **Aturan perilaku (acceptance criteria):**
  1. Nilai dihitung dari jumlah work item atau story points per status work item tertaut, mengikuti estimation value di JPD-INS-018.
  2. Ditampilkan sebagai lozenge berwarna sesuai status category dominan (mis. "In progress"), dengan tooltip rincian.
  3. ⚠ Field dapat dipakai di view untuk tampilan, filter, dan group (mis. kolom board per delivery status).
  4. Read-only dan diperbarui ketika status work item tertaut berubah, ditautkan, atau dilepas.
  5. Tidak disertakan pada ekspor Excel/Google Sheets.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status category di WF-002 ([09](09-workflow-status-transisi.md)); field terhitung di modul 36.

### JPD-INS-020 Bulk Create Work Items from List View

- **Deskripsi:** Membuat work item delivery untuk banyak idea sekaligus dari list view, masing-masing otomatis tertaut ke idea asalnya.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dengan permission Create work items di proyek target.
- **Aturan perilaku (acceptance criteria):**
  1. Prasyarat: proyek memiliki minimal satu list view; aksi tidak tersedia di board, matrix, atau timeline.
  2. Pengguna mencentang checkbox di samping beberapa idea; toolbar bulk menampilkan jumlah terpilih.
  3. Klik "Create work items" > pilih proyek dan work type > Create; dialog menampilkan pratinjau jumlah work item yang akan dibuat.
  4. Setiap idea mendapat satu work item baru (summary = summary idea) yang otomatis tertaut di tab Delivery; kegagalan per idea dilaporkan tanpa membatalkan idea lain.
  5. Setelah selesai, sistem menyarankan menambahkan field Delivery progress dan Delivery status ke view bila belum tampil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** List View dan Idea Bulk Edit & Bulk Actions di modul 36; bulk API di API-026 ([23](23-api-publik-webhook.md)).

### JPD-INS-021 Autofill Dates

- **Deskripsi:** Mengisi field date idea secara otomatis dari work item delivery tertaut atau dari target date project di modul Goals & Projects.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator atau Admin Proyek/Project Lead (role Administrator; membuat/mengedit field; lihat JPD-INS-003).
- **Aturan perilaku (acceptance criteria):**
  1. Prasyarat: minimal satu date field terisi di work item tertaut; hanya Creator/Administrator yang dapat membuat/mengedit field.
  2. Di daftar Fields, pengguna memilih atau membuat date field lalu mengaktifkan "Autofill dates"; field menjadi read-only bagi pengguna selama autofill aktif.
  3. Date source dapat dipilih: "Project target date" (dari Goals & Projects, JPD-INS-036) atau "Linked work items".
  4. Untuk linked work items, pengguna memilih date field sumber (mis. Due date, Target start, Target end) dan metode "Calculation by": earliest atau latest.
  5. Rule opsional menambahkan warna dan emoji saat tanggal mendekati (N hari) atau terlewat.
  6. Hover tanggal di view menampilkan keterangan apakah nilai dihitung dari tanggal paling awal atau paling akhir dan dari work item mana.
  7. Field autofill dipakai untuk posisi idea di timeline view dan diperbarui saat tanggal work item berubah.
  8. Untuk work item dari instance self-hosted eksternal (JPD-INS-024), autofill mendukung due date, actual start/end, dan target start/end; start date biasa tidak didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Date & Interval Fields dan Timeline View (Roadmap) di modul 36; due/start date di ISS-021 ([06](06-work-item-inti.md)).

### JPD-INS-022 Ideas in Jira Work Items & Jira Plans

- **Deskripsi:** Dari sisi pelacak kerja (proyek software dan Plans), work item menampilkan idea terkait sehingga tim engineering dan perencana melihat konteks discovery.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Developer/Anggota Tim (melihat panel idea), Scrum Master/Team Lead dan perencana (Plans; plan Premium/Enterprise), Creator.
- **Aturan perilaku (acceptance criteria):**
  1. Work item view pelacak kerja memiliki panel idea yang menampilkan idea tertaut (key, summary, status, ikon idea type).
  2. Mengklik idea membuka idea view di Product Discovery; pengguna tanpa akses project discovery melihat pesan "Anda tidak memiliki akses ke idea ini" dan summary disembunyikan.
  3. ⚠ JQL `issueLinkType = "is implemented by"` (atau fungsi setara) dapat dipakai untuk mencari Epic yang terkait idea.
  4. Di Plans (memerlukan plan Premium/Enterprise pelacak kerja) pengguna menambahkan kolom field "Idea" ke plan lewat ikon +.
  5. Hover nama idea di plan menampilkan key idea dan tautan ke idea.
  6. Pengguna dapat menautkan work item ke idea dari kolom Idea di plan (memilih idea dari proyek yang dapat diakses).
  7. Plan dapat difilter berdasarkan idea.
  8. View settings > Group by > Idea mengelompokkan work item plan per idea.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Plans di PLAN-027 dan modul 18; JQL di modul 12; izin Plans PERM-052 ([03](03-izin-peran-keamanan.md)).

### JPD-INS-023 Delivery Automation Templates & Rules

- **Alias inventaris:** Automation for JPD
- **Deskripsi:** Otomasi berbasis aturan (modul 20) untuk menyinkronkan idea dan work item delivery, dengan template khusus discovery di tingkat proyek dan global, serta dukungan trigger, action, dan smart values field discovery.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; rule dan template proyek), admin otomasi global/Admin Situs (template global), Creator (melihat hasil).
- **Aturan perilaku (acceptance criteria):**
  1. Template proyek (Project settings > Automation > Templates) minimal mencakup: idea transitioned to delivery → create delivery work item; idea done → archive; idea idle 3 bulan → notify assignee; idle 4 bulan → kembalikan ke Parking lot; transitioned to discovery → add watchers; added to roadmap → update Slack channel; idea created → create halaman dokumentasi.
  2. Global templates (System > Global automation > Templates) minimal mencakup: delivery work item created → copy description from idea; delivery work item done → mark idea as done.
  3. Contoh rule yang harus dapat dibuat: trigger "issue linked" dengan link type delivery + kondisi work type Epic → edit description Epic.
  4. Rule global: work item Done → branch linked issues "implements" → transition idea ke Done; untuk idea dengan banyak Epic, template menambah kondisi "semua work item tertaut Done" sebelum transition.
  5. Template yang diaktifkan dapat disesuaikan atau dinonaktifkan kapan saja tanpa memengaruhi rule lain.
  6. Beberapa field discovery belum didukung penuh oleh action edit; field date dan checkbox memerlukan format nilai khusus (mis. date sebagai ISO `YYYY-MM-DD` atau JSON fuzzy date, checkbox sebagai `true/false`), dan nilai tidak valid membuat rule gagal dengan pesan di audit log rule.
  7. ⚠ Trigger khusus "insight added" dan "vote changed" tersedia untuk project discovery.
  8. Plan Premium: rule multi-proyek dapat mencakup banyak proyek tanpa batas; jumlah eksekusi mengikuti kuota otomasi per plan (AUTO-015).
  9. ⚠ Trigger yang tersedia untuk idea: idea created, field value changed, status changed, delivery work item linked/updated.
  10. ⚠ Action yang tersedia untuk idea: edit field idea, transition idea, create delivery work item (otomatis tertaut).
  11. Template "transition idea ke Done saat semua delivery work item Done" tersedia sebagai template siap pakai.
  12. Smart values mendukung field discovery (mis. `{{issue.Impact}}`, `{{issue.insights.count}}`, `{{issue.deliveryProgress}}`) dan field yang tidak didukung mengembalikan nilai kosong.
  13. Rule dapat dibuat per project discovery dengan scope single-proyek secara default.
  14. Eksekusi rule discovery dihitung terhadap batas eksekusi otomasi plan (AUTO-015); saat kuota habis rule berhenti berjalan dan Administrator proyek menerima notifikasi "Batas eksekusi otomasi tercapai".
  15. ⚠ Rule dengan trigger manual dapat dijalankan langsung dari idea view (menu aksi idea > Automation > pilih rule); daftar hanya menampilkan rule manual yang scope-nya mencakup proyek idea dan yang boleh dijalankan pengguna tersebut, eksekusi memakai idea itu sebagai konteks trigger, dan hasilnya tercatat di audit log rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan; project discovery memakai model rule single-proyek seperti team-managed.
- **Catatan & ketergantungan:** Template library AUTO-010, administrasi global AUTO-002, smart values AUTO-063 ([20](20-otomasi.md)); template di JPD-INS-042; Slack di JPD-INS-032.

### JPD-INS-024 JPD Integration for Jira Data Center

- **Deskripsi:** Menghubungkan idea di Product Discovery (cloud) ke work item di instance pelacak kerja self-hosted eksternal untuk delivery tracking (Premium). Di produk ini ditafsirkan sebagai integrasi ke instance self-hosted yang kompatibel dengan API rujukan, melalui komponen penghubung yang dipasang gratis di sisi instance tersebut.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs di kedua sisi, Creator.
- **Aturan perilaku (acceptance criteria):**
  1. Integrasi memerlukan plan Premium Product Discovery; komponen penghubung di sisi instance self-hosted tidak berbayar.
  2. Prasyarat versi minimum instance self-hosted dicatat di halaman integrasi (setara 10.3+, dan 10.3.10+ untuk panel idea); Admin Situs dibutuhkan di kedua sisi dan issue linking harus aktif di instance self-hosted.
  3. Koneksi dibuat lewat application link atau application tunnel; tunnel hanya menghubungkan satu instance self-hosted ke satu site Product Discovery.
  4. Setiap pengguna membutuhkan akun di kedua sistem dan autentikasi OAuth silang pada pemakaian pertama.
  5. Pengguna dapat membuat atau menautkan work item self-hosted dari idea; delivery progress dan linked items tampil di Product Discovery, dan panel idea tampil di work item self-hosted.
  6. Create memerlukan permission CREATE_ISSUES di instance self-hosted; hanya Project, Issue type, Summary, dan Assignee yang boleh menjadi field wajib di layar create, dan description dikirim sebagai plain text; bila field wajib lain ada, create ditolak dengan pesan yang menyebut field tersebut.
  7. Delivery progress memerlukan link type "implements / is implemented by" di instance self-hosted; estimation type Multilevel tidak didukung.
  8. Field Team tidak tersedia untuk work item dari instance self-hosted.
  9. Komponen penghubung menyediakan audit log, app log, statistik, endpoint OpenMetrics, dan rate limiting opsional (default off).
  10. Remote link tidak ikut tersalin saat work item self-hosted di-clone atau di-split.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada application link (ADM-032, API-014). Produk ini sendiri tidak memiliki edisi Data Center vendor; pembangunan komponen penghubung untuk produk vendor bersifat opsional dan dicatat di Pertanyaan terbuka.

## Kolaborasi, berbagi & notifikasi

### JPD-INS-025 Comments & Reactions on Ideas

- **Alias inventaris:** Comments and collaboration on ideas
- **Deskripsi:** Diskusi dan reaksi (comments and reactions) pada idea memakai komentar work item standar (modul 08), terbuka untuk Contributor tanpa lisensi Creator, lengkap dengan field dinamis Comments dan Reactions untuk dipakai di view.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (menulis komentar dan reaction), penulis komentar (edit/hapus miliknya), Admin Proyek/Project Lead (role Administrator; menghapus komentar siapa pun); permission Add comments dan Delete own/all comments di role proyek (JPD-INS-003).
- **Aturan perilaku (acceptance criteria):**
  1. Field dinamis "Comments" menampilkan jumlah komentar pada idea; mengklik nilainya di view (List, Board, Matrix, Timeline) membuka sidebar idea langsung pada tab komentar.
  2. Komentar memakai editor rich text (modul 07) dan mendukung mention @user; pengguna yang di-mention menerima notifikasi (JPD-INS-031) meskipun ia belum menjadi watcher.
  3. Pengguna dapat memberi reaction emoji pada komentar; field dinamis "Reactions" menghitung total reaction pada idea dan dapat ditampilkan sebagai kolom view.
  4. Pengguna dapat mengedit dan menghapus komentarnya sendiri; komentar yang diedit diberi penanda "diedit" beserta waktu edit, dan penghapusan meminta konfirmasi.
  5. Field Reactions tidak disertakan pada ekspor CSV dan tidak dapat dipakai untuk sort, filter, atau group di published view (JPD-INS-029); view yang memakainya untuk sort/filter/group tidak dapat di-publish.
  6. Add, edit, dan delete comment pada idea mengikuti perilaku komentar work item standar (CMT-001 s.d. CMT-003), termasuk riwayat di tab aktivitas idea.
  7. Mention @user mencari pengguna yang dapat melihat proyek; pengguna tanpa akses proyek tetap dapat di-mention namun UI memperingatkan "Pengguna ini tidak dapat melihat idea ini".
  8. Reaction pada komentar memakai set emoji yang sama dengan komentar work item (CMT-006); satu pengguna hanya dapat memberi satu reaction per emoji per komentar (klik ulang membatalkan).
  9. Komentar baru, edit, dan mention mengirim notifikasi ke watchers idea sesuai skema notifikasi proyek (JPD-INS-031).
  10. Contributor dapat menambah komentar tanpa lisensi Creator; Contributor tidak dapat menghapus komentar pengguna lain.
  11. Field "Comments count" (jumlah komentar) dapat dipakai untuk sort dan filter di view internal serta sebagai input formula (modul 36).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komentar dan reaction di CMT-001 s.d. CMT-006 ([08](08-komentar-lampiran-aktivitas.md)); mention di EDT-014 ([07](07-editor-konten-kaya.md)) dan NTF-013 ([19](19-notifikasi-email.md)). Komentar dapat dijadikan insight lewat menu "Buat insight" (JPD-INS-009).

### JPD-INS-026 Votes Field

- **Deskripsi:** Field dinamis untuk voting idea dengan kuota vote per pengguna (vote budget), sehingga stakeholder internal dapat menyatakan dukungan secara terukur.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; membuat field dan menetapkan kuota), Creator dan Contributor (memberi vote).
- **Aturan perilaku (acceptance criteria):**
  1. Administrator menetapkan jumlah vote per pengguna (vote budget) saat membuat atau mengedit field Votes; nilai minimal 1, dan perubahan kuota tidak menghapus vote yang sudah diberikan tetapi pengguna yang melebihi kuota baru tidak dapat menambah vote sampai sisanya kembali positif.
  2. Creator dan Contributor dapat memberi vote; Stakeholder pada published view tidak dapat vote.
  3. Pengguna dapat menambah dan mengurangi vote-nya pada idea (lebih dari satu vote per idea diizinkan selama kuota tersisa); UI menampilkan sisa vote "Anda memiliki N vote tersisa", dan percobaan melebihi kuota ditolak dengan pesan "Kuota vote Anda telah habis".
  4. Field Votes dapat dipakai untuk sort dan filter di view internal serta sebagai input Formula Fields (modul 36).
  5. ⚠ Tooltip/sidebar field menampilkan ringkasan vote (vote summaries): total vote, jumlah pemberi vote, dan daftar pemberi vote dengan jumlah vote masing-masing; ringkasan dapat dirangkum AI pada plan Premium (JPD-INS-041).
  6. Votes tidak dapat dipakai untuk sort, filter, atau group di published view; view yang memakainya untuk hal tersebut tidak dapat di-publish (JPD-INS-029).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari vote work item standar (izin voters di PERM-033, [03](03-izin-peran-keamanan.md)); field Votes discovery bertipe khusus di katalog field modul 36. Vote disimpan per (idea, field, user, jumlah).

### JPD-INS-027 Comment on Views

- **Deskripsi:** Komentar di tingkat view (bukan idea) untuk komunikasi antara tim customer-facing dan tim produk tentang isi view atau roadmap.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Siapa pun dengan akses view (Creator, Contributor; Stakeholder hanya membaca pada published view).
- **Aturan perilaku (acceptance criteria):**
  1. Siapa pun yang memiliki akses ke view (sesuai access level proyek dan view restriction, JPD-INS-005) dapat berkomentar; pengguna tanpa akses edit view tetap dapat berkomentar.
  2. Tombol "Comment" di kanan atas view membuka panel komentar view; pengguna menulis komentar lalu klik Create, dan komentar baru tampil di urutan teratas dengan author dan waktu.
  3. Pengguna dapat mengedit dan menghapus komentarnya sendiri serta memberi reaction emoji pada komentar view.
  4. Pada plan Premium, tombol "Summarize" merangkum seluruh komentar view dengan AI (JPD-INS-041); tombol disembunyikan bila AI dinonaktifkan admin.
  5. ⚠ Komentar view tidak memiliki status resolve dan tidak mengirim notifikasi email pada versi awal; mention @user di komentar view mengirim notifikasi in-app.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `view_comments` terpisah dari komentar work item; view di modul 36.

### JPD-INS-028 Share Views with Team

- **Deskripsi:** Berbagi view kepada orang yang sudah atau akan memiliki akses ke proyek, melalui email, notifikasi, atau tautan langsung.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (membagikan), Admin Proyek/Project Lead (role Administrator; mengatur izin pengguna baru yang diundang).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Share > masukkan nama pengguna, grup, atau alamat email > Share; dapat menambah pesan opsional.
  2. Opsi "Share view in full screen" membuat tautan yang membuka view dalam mode layar penuh tanpa navigasi samping.
  3. Penerima mendapat email dan notifikasi bell in-app berisi nama view, pengirim, dan pesan.
  4. "Copy link" menyalin URL view untuk pengguna yang sudah berada di proyek; pengguna tanpa akses yang membuka tautan melihat halaman "Minta akses".
  5. Membagikan ke pengguna yang belum memiliki akses langsung mengirim undangan; pengguna tersebut masuk sebagai Contributor dan Administrator proyek mengatur role/izinnya lebih lanjut (JPD-INS-003).
  6. Dialog Share menampilkan petunjuk bahwa pihak luar organisasi sebaiknya diberikan published view (JPD-INS-029), bukan undangan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Undangan pengguna mengikuti kebijakan undangan organisasi di modul 02; email di modul 19.

### JPD-INS-029 Publish Views

- **Deskripsi:** Publikasi view sebagai halaman read-only untuk stakeholder tanpa lisensi atau akses proyek, yang ter-update otomatis mengikuti view sumber.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs (mengaktifkan fitur), Admin Proyek/Project Lead (role Administrator; default dapat publish), Creator dengan custom role "Publish a view", Stakeholder (pembaca).
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs mengaktifkan publish di Apps > Product Discovery > Access; nilai default nonaktif, dan selama nonaktif tombol Publish disembunyikan.
  2. Publish memerlukan plan Standard atau lebih tinggi; pada plan Free tombol Publish menampilkan ajakan upgrade (JPD-INS-008).
  3. Secara default hanya Administrator proyek yang dapat publish; Creator dapat publish bila diberi custom role dengan permission "Publish a view" (JPD-INS-003).
  4. Contributor tidak dapat publish; opsi tidak ditampilkan dan API mengembalikan HTTP 403.
  5. Share > Publish menampilkan toggle field mana yang ditampilkan dan apakah description idea ikut tampil, beserta tombol Preview sebelum publish.
  6. Setelah publish, "Copy link" menyalin URL published view; viewer dapat memfilter dan men-sort ulang data di published view tanpa mengubah view sumber.
  7. ⚠ Publish ke siapa pun dengan tautan (tanpa akun) hanya tersedia bila Admin Situs mengaktifkannya dan pengguna memiliki permission "Publish a view to anyone with the link"; ketersediaan per plan mengikuti JPD-INS-008.
  8. Pengguna organisasi yang memakai SSO tetap membutuhkan akun produk ini untuk membuka published view yang tidak dipublikasikan ke "siapa pun dengan tautan".
  9. View yang di-publish ditandai ikon globe hijau di daftar view; Share > Publish > Unpublish menonaktifkan tautan dan pembuka tautan lama melihat "View ini tidak lagi dipublikasikan".
  10. View tidak dapat di-publish bila memakai field Goals, Type, project Goals & Projects, Votes, Reactions, Custom formula, Number of linked work items, atau Number of insights untuk sort/filter/group/kolom/sumbu matrix; dialog menyebut field penyebabnya.
  11. Published view ter-update otomatis ketika view sumber atau datanya berubah, tanpa perlu publish ulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Publish dasar view juga dirujuk di modul 36; batas field published view disebut di JPD-INS-017, JPD-INS-025, JPD-INS-026. Tautan publik harus memakai token acak tidak dapat ditebak dan dapat dicabut (unpublish).

### JPD-INS-030 Stakeholder Access & Feedback Collection

- **Deskripsi:** Kumpulan jalur bagi stakeholder internal maupun eksternal untuk melihat roadmap dan memberi feedback ke tim produk tanpa lisensi Creator.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Stakeholder/Viewer, Creator, Contributor, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Stakeholder internal yang memiliki akses produk lain otomatis dapat masuk sebagai Contributor tidak berlisensi (JPD-INS-001) dan memberi comment, vote, serta insight.
  2. Stakeholder tanpa akses proyek melihat roadmap melalui published view (JPD-INS-029) secara read-only.
  3. Feedback dapat dikumpulkan melalui Service Management (JPD-INS-035), channel Slack/Teams (JPD-INS-032, JPD-INS-033), shared view (JPD-INS-028), serta field dan votes (JPD-INS-026); setiap jalur berakhir sebagai insight, comment, atau vote pada idea.
  4. View dapat di-embed di halaman dokumentasi (JPD-INS-034) sehingga stakeholder membaca roadmap di tempat kerjanya, dengan izin pembaca dihormati.
  5. Stakeholder yang memiliki akses (Contributor) dapat watch idea untuk menerima pembaruan (JPD-INS-031).
  6. ⚠ Tidak ada role "Stakeholder" yang dapat ditugaskan dan tidak ada email "stakeholder updates" khusus pada versi awal; pembaruan stakeholder dilakukan lewat watch, otomasi, dan published view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini adalah komposisi fitur lain; tidak membutuhkan entitas baru. Portal widget feedback pihak ketiga di luar cakupan.

### JPD-INS-031 Space Notifications & Watching

- **Alias inventaris:** JPD Notifications; Notifications for ideas
- **Deskripsi:** Notifikasi email dan in-app untuk event idea (comment, mention, perubahan status, insight baru), skema notifikasi per proyek bergaya team-managed, dan watch/unwatch idea.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; mengatur notifikasi proyek), Creator dan Contributor (watch/unwatch, menerima notifikasi).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Notifications tersedia untuk project discovery dengan model notifikasi team-managed (daftar event dan penerima per proyek).
  2. Default untuk event work item: notifikasi dikirim ke watchers, assignee, dan reporter idea.
  3. "Add notification" memungkinkan memilih event lalu penerima berupa pengguna tertentu, role proyek, atau grup.
  4. Notifikasi yang ada dapat diedit (mengubah penerima) dan dihapus; menghapus seluruh penerima suatu event mematikan notifikasi event tersebut.
  5. Pengguna dapat watch/unwatch idea dari idea view; otomasi dapat menambah watchers (template "transitioned to discovery > add watchers", JPD-INS-023).
  6. Notifikasi personal Slack dan Microsoft Teams dapat diaktifkan masing-masing pengguna (JPD-INS-032, JPD-INS-033).
  7. ⚠ Event "insight added" tersedia sebagai event notifikasi proyek dan dikirim ke watchers idea; email digest tidak tersedia pada versi awal.
  8. Mention dan comment pada idea mengirim notifikasi ke pengguna yang di-mention dan watchers (NTF-013).
  9. Watch idea menambah pengguna ke daftar watchers idea yang sama dengan watcher work item (ISS-036).
  10. ⚠ Penambahan insight baru pada idea mengirim notifikasi in-app ke watchers idea, berisi author insight dan cuplikan description.
  11. Tombol Watch/Unwatch tampil di header idea view dengan jumlah watchers; menonton idea tidak memerlukan lisensi Creator.
  12. Event yang memicu notifikasi minimal mencakup: comment baru, mention, insight baru, dan perubahan status idea.
  13. ⚠ Pengaturan notifikasi khusus project discovery hanya mengatur event idea; preferensi pribadi pengguna (modul 19) tetap dapat mematikan kanal email.
  14. Setiap notifikasi juga tampil di pusat notifikasi in-app (bell) produk ini dan dapat ditandai sudah dibaca.
- **Team-managed vs company-managed:** Project discovery memakai skema notifikasi team-managed (per proyek), bukan notification scheme company-managed.
- **Catatan & ketergantungan:** Watchers di NTF-012 dan mention di NTF-013 ([19](19-notifikasi-email.md)); notifikasi Slack/Teams personal di NTF-039, NTF-040; watchers work item di ISS-036 ([06](06-work-item-inti.md)).

## Integrasi chat, dokumentasi, layanan & goals

### JPD-INS-032 Slack Integration (Jira Cloud for Slack)

- **Alias inventaris:** Slack integration (JPD)
- **Deskripsi:** Integrasi Slack melalui app chat umum produk ini (padanan Jira Cloud for Slack, bukan app khusus discovery) untuk membuat idea dari pesan, menambahkan pesan sebagai insight, preview tautan, aksi cepat, dan notifikasi.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (membuat idea/insight), admin workspace Slack (memasang app), Admin Situs (mengaktifkan app di produk ini), Admin Proyek/Project Lead (role Administrator; menghubungkan channel).
- **Aturan perilaku (acceptance criteria):**
  1. Prasyarat: app Slack produk ini terpasang dan terhubung ke site, pengguna memiliki app access ke Product Discovery/pelacak kerja/Service Management, dan minimal satu project discovery ada; bila prasyarat tidak terpenuhi daftar site/proyek di dialog kosong dengan pesan penyebabnya.
  2. App dipasang di workspace Slack oleh admin workspace; setiap pengguna login ke akun produk ini dari Slack untuk menautkan identitas Slack dengan akunnya.
  3. Admin Situs harus mengaktifkan app di halaman konfigurasi integrasi; bila dinonaktifkan, pengguna Slack melihat pesan "Koneksi ke <site> dinonaktifkan oleh admin".
  4. Create idea dari pesan: hover pesan > ... > More message shortcuts > "Create issue from" > pilih site discovery, proyek, work type "Idea", isi detail > submit; idea dibuat dengan pesan sebagai description awal.
  5. Setelah idea dibuat, pesan konfirmasi Slack memberi peringatan bahwa idea baru mungkin tidak langsung terlihat bila view yang dibuka memiliki filter.
  6. Add insight dari pesan: ... > More message shortcuts > "Add to issue" > pilih site, ketik nama idea > Add; insight tersimpan di tab Insights idea target.
  7. Channel dapat dihubungkan ke proyek dengan mengundang bot atau perintah slash `connect` milik app produk ini (padanan rujukan `/jira connect`; tidak berlaku di DM); perintah `manage` (padanan `/jira manage`) mengubah pengaturan koneksi channel.
  8. Notifikasi personal diatur lewat perintah slash `notify` milik app produk ini (padanan rujukan `/jira notify`) dengan pilihan timing (segera atau ringkasan berkala).
  9. Tautan atau key work item/idea yang ditempel di Slack menampilkan preview; key huruf besar tanpa URL hanya di-unfurl bila channel terhubung ke proyek tersebut.
  10. Dari preview dan Home tab app, pengguna dapat watch, unwatch, assign, transition, dan comment pada idea sesuai izinnya.
  11. Semua aksi menghormati izin produk ini; work item/idea dengan issue security level tidak memicu notifikasi Slack dan tidak di-unfurl.
  12. Dokumentasi bantuan menyarankan pola dedicated channel (mis. #product-feedback) yang terhubung ke project discovery untuk mengumpulkan feedback internal.
  13. Template otomasi "When an idea is added to the roadmap > update Slack channel" tersedia (JPD-INS-023) dan memakai action kirim pesan Slack.
  14. Integrasi tersedia sebagai bagian app Slack produk ini; tidak ada instalasi app Slack terpisah khusus discovery.
  15. Message action "Add insight" (label "Add to issue") tersedia pada setiap pesan di channel publik, privat, dan DM yang dapat diakses bot.
  16. Dialog add insight memungkinkan memilih proyek dan idea target, atau membuat idea baru bila idea yang dicari belum ada.
  17. Insight yang dibuat dari Slack menyimpan snippet teks pesan dan permalink ke pesan sumber sebagai link insight.
  18. Tab Insights menampilkan smart preview pesan Slack (author, channel, waktu, cuplikan) untuk insight bersumber Slack, sesuai izin pembaca terhadap channel.
  19. Autentikasi akun produk ini dari Slack memakai OAuth; token dapat dicabut pengguna dari halaman Connected apps profilnya.
  20. ⚠ Perubahan idea (status, komentar, insight baru) dapat dikirim sebagai notifikasi ke channel Slack terhubung atau DM personal sesuai pengaturan perintah `notify` dan `manage` (aturan 7 dan 8).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App Slack umum di INT-030 s.d. INT-035 ([25](25-integrasi-impor-ekspor-ekstensi.md)) dan NTF-039 ([19](19-notifikasi-email.md)); modul ini hanya menambah work type Idea, aksi add insight, dan preview insight. App Slack adalah artefak resmi produk ini yang dipublikasikan di direktori Slack.

### JPD-INS-033 Microsoft Teams Integration (Jira Cloud for Microsoft Teams)

- **Alias inventaris:** Microsoft Teams integration (JPD)
- **Deskripsi:** Integrasi Microsoft Teams melalui app chat umum produk ini (padanan Jira Cloud for Microsoft Teams) untuk notifikasi, tab idea, membuat idea, dan menambah insight dari percakapan.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor, admin Microsoft Teams (memasang app), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Notifikasi personal diaktifkan dengan mengetik "notifications" ke bot, memilih jenis notifikasi, lalu menyimpan.
  2. Tab app di chat/channel dibuat dengan memilih app produk ini, site, dan saved filter (JPD-INS-038) untuk menampilkan daftar idea.
  3. Create idea dari percakapan: menu ... pada pesan > "Create issue in <nama produk>" > pilih project discovery dan work type Idea > isi summary dan description > simpan.
  4. Add insight dari percakapan: menu ... > "Comment in <nama produk>" > tulis insight > pilih proyek dan idea target > post; hasilnya tersimpan sebagai insight, bukan komentar.
  5. Pengguna dapat mencari dan membagikan idea dari area compose (ikon app) sebagai kartu preview.
  6. Uninstall dilakukan dengan klik kanan ikon app > Uninstall; token koneksi pengguna dicabut saat uninstall.
  7. Integrasi memerlukan akun Microsoft Teams berbayar; pada akun gratis halaman pemasangan menampilkan pesan bahwa akun tidak didukung.
  8. Notifikasi channel dikirim melalui action otomasi "Send Microsoft Teams message" (modul 20).
  9. App dipasang dari katalog app Microsoft Teams oleh pengguna atau admin Teams sesuai kebijakan tenant.
  10. Message action add insight tersedia pada pesan chat dan channel.
  11. Dialog add insight memungkinkan memilih proyek/idea target atau membuat idea baru.
  12. Insight dari Teams menyimpan tautan ke pesan sumber; pembuka tautan tanpa akses ke percakapan Teams hanya melihat URL.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App Teams umum di INT-036 ([25](25-integrasi-impor-ekspor-ekstensi.md)) dan NTF-040 ([19](19-notifikasi-email.md)).

### JPD-INS-034 Confluence Integration

- **Alias inventaris:** Confluence integration (JPD)
- **Deskripsi:** Integrasi dengan modul/aplikasi dokumentasi (padanan Confluence): insight dari catatan interview, tautan spec, embed view dan roadmap, serta Smart Link idea di halaman dokumentasi.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor, pengguna dokumentasi (pembaca embed), Stakeholder.
- **Aturan perilaku (acceptance criteria):**
  1. Catatan interview di halaman dokumentasi dapat ditangkap sebagai insight melalui ekstensi browser (JPD-INS-014) dengan highlight teks.
  2. Spec dapat ditautkan ke idea dengan menempel URL halaman di description idea (dirender sebagai Smart Link) atau di field Hyperlink.
  3. Embed view: menempel URL view discovery di halaman dokumentasi lalu memilih tampilan "Embed" menampilkan view interaktif read-only di halaman.
  4. Embed menghormati izin: pembaca tanpa akses ke proyek/view melihat kartu "Anda tidak memiliki akses" tanpa data idea.
  5. Tautan idea di halaman dokumentasi dapat ditampilkan sebagai inline, card, atau embed (EDT-021).
  6. Template otomasi "idea created > create halaman dokumentasi" tersedia (JPD-INS-023) dan membuat halaman dengan tautan balik ke idea.
  7. Embed view atau roadmap memakai mekanisme Smart Link embed yang sama untuk view List, Board, Matrix, Timeline, dan roadmap lintas proyek.
  8. Idea yang ditempel sebagai Smart Link inline/card menampilkan key, summary, status, dan ikon idea type.
  9. URL halaman dokumentasi yang ditempel di "Paste a link" insight (JPD-INS-010) dirender sebagai kartu halaman (judul, space dokumentasi, penulis).
  10. Data idea pada embed dan card selalu dihitung sesuai izin pembaca saat dibuka, bukan izin penulis halaman.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Integrasi dokumentasi di INT-042 s.d. INT-045 ([25](25-integrasi-impor-ekspor-ekstensi.md)); Smart Link di EDT-021 ([07](07-editor-konten-kaya.md)). Produk dokumentasi vendor tidak dibangun; integrasi diarahkan ke modul dokumentasi produk ini atau aplikasi dokumentasi yang mendukung oEmbed/iframe.

### JPD-INS-035 Jira Service Management Integration

- **Alias inventaris:** Jira Service Management integration (JPD)
- **Deskripsi:** Menghubungkan Service Management dengan Product Discovery sehingga request support menjadi insight pada idea dan permintaan fitur dari portal dapat menjadi idea.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Agent Layanan (Contributor di project discovery), admin Service Management (otomasi portal), Creator.
- **Aturan perilaku (acceptance criteria):**
  1. Prasyarat: service project dengan queue berisi request (modul 29/30) dan minimal satu project discovery; tanpa keduanya opsi integrasi tidak ditampilkan.
  2. Request dapat ditambahkan sebagai insight melalui ekstensi browser (JPD-INS-014) atau menempel tautan request di "Paste a link" (JPD-INS-010).
  3. Label dan properti insight yang berasal dari Service Management (mis. organisasi customer, request type, prioritas) dapat dikonfigurasi di pengaturan data provider (JPD-INS-013).
  4. Ideas dari portal: rule otomasi di service project (trigger work item created + request type "New Feature") membuat Idea di project discovery dengan pemetaan field (summary, description, reporter, tautan balik ke request).
  5. Administrator dapat mengizinkan customer luar membagikan idea lewat portal (request type khusus permintaan fitur), tanpa memberi mereka akses ke project discovery.
  6. ⚠ Panel "Linked ideas" di agent view request menampilkan idea yang memakai request tersebut sebagai insight (key, summary, status).
  7. Dari request, agent dapat menautkan request ke idea yang ada sebagai insight.
  8. Dari request, agent dapat membuat idea baru dengan request sebagai insight pertamanya.
  9. Tab Insights idea menampilkan request terkait sebagai kartu (key, summary, status, reporter) sesuai izin pembaca.
  10. Agent Service Management dapat menambah insight sebagai Contributor tanpa lisensi Creator (JPD-INS-001).
  11. ⚠ Insight bersumber request menyertakan data customer dan organisasi Service Management sebagai label/properti insight, dan dapat dipakai untuk filter insight.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada modul 29 (service project, request type) dan customers & organizations di modul JSM; otomasi lintas produk di modul 20. Karena Service Management berada di Fase 5, integrasi ini dibangun setelah Fase 5 selesai.

### JPD-INS-036 Atlassian Goals & Projects Integration (Atlas)

- **Deskripsi:** Menautkan idea ke Goals dan project di modul Goals & Projects produk ini (padanan Atlas/Home) sehingga discovery selaras dengan tujuan organisasi.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator; mengaktifkan integrasi), Creator (mengisi field).
- **Aturan perilaku (acceptance criteria):**
  1. Integrasi aktif secara default untuk proyek baru; untuk proyek lama Administrator mengaktifkannya di Project settings > Features > "Enable Goals & Projects integration".
  2. Administrator proyek memilih site sumber Goals & Projects yang dipakai proyek.
  3. Field "Goals", "Project", dan "Project status" ditambahkan ke view melalui Fields > Available fields.
  4. Tombol "Connect to Goals & Projects" di grid view tampil bila integrasi belum dihubungkan dan membuka alur koneksi.
  5. Site yang tidak memiliki akses Goals & Projects tidak muncul di daftar pilihan site sumber.
  6. Target date project dapat disinkronkan ke date field idea melalui Autofill Dates (JPD-INS-021).
  7. Field Goals & Projects tidak disertakan pada ekspor CSV dan tidak dapat dipakai di published view (JPD-INS-029).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Goals di PLAN-061, PLAN-062 ([18](18-timeline-rencana-tampilan-proyek.md)); INT-056 ([25](25-integrasi-impor-ekspor-ekstensi.md)). Bila modul Goals & Projects belum ada saat Fase 6, fitur ditunda.

## Data, API & platform

### JPD-INS-037 JPD REST API & Webhooks

- **Alias inventaris:** Jira Product Discovery APIs (Appendix); REST API for ideas
- **Deskripsi:** Akses programatik ke idea melalui REST API work item standar produk ini (karena idea adalah work item) beserta webhook event work item, dengan batasan bahwa insight, view, dan konfigurasi delivery belum memiliki REST publik penuh.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Integrator/Developer API, Admin Situs (mendaftarkan webhook), Creator (pemilik token yang menulis idea).
- **Aturan perilaku (acceptance criteria):**
  1. Endpoint work item standar (padanan `/rest/api/3/issue`, di produk ini `/api/v1/issues`) dapat membaca, membuat, mengubah, dan menghapus idea; permintaan memakai izin pemilik token, sehingga Contributor yang mencoba mengedit idea mendapat HTTP 403.
  2. Field discovery (Impact, Effort, Insights count, field kustom proyek) terekspos sebagai custom field dengan ID stabil di respons dan metadata field.
  3. ⚠ Insight, view, dan konfigurasi delivery tidak diekspos di REST publik pada versi awal; endpoint internal SPA tidak didokumentasikan di Swagger publik dan dapat berubah tanpa pemberitahuan.
  4. Webhook event work item (created, updated, deleted, comment) berlaku untuk idea dan menyertakan `project_type_key='product_discovery'` di payload sehingga penerima dapat memfilter.
  5. ⚠ Platform ekstensi (padanan Forge/Connect) dapat membaca idea melalui API work item yang sama; tidak ada extension point khusus discovery pada versi awal.
  6. Ekspor isi insight lengkap (description, link, label, impact) untuk satu proyek disediakan melalui permintaan ekspor admin (job asinkron yang menghasilkan file unduhan), sebagai pengganti API insight.
  7. Idea adalah work item dengan work type "Idea" di proyek ber-`project_type_key='product_discovery'` yang dapat dibaca/ditulis lewat endpoint work item standar.
  8. ⚠ Dukungan tulis untuk field discovery lewat API terbatas: field terhitung (Delivery progress, Delivery status, formula, Insights count) bersifat read-only dan upaya menulisnya ditolak HTTP 400 "Field ini bersifat read-only".
  9. ⚠ Permintaan ke path REST publik untuk insight atau view (mis. `/api/v1/discovery/insights`) dengan token API pihak ketiga mengembalikan HTTP 404, dan dokumentasi Swagger publik tidak mencantumkan path tersebut; endpoint internal SPA hanya menerima sesi browser.
  10. Endpoint pencarian project (padanan `/project/search`) menerima filter `typeKey=product_discovery` dan mengembalikan type key tersebut di setiap project discovery.
  11. Operasi CRUD idea (create, read, update, delete) via API work item mencatat history dan memicu otomasi serta notifikasi sama seperti aksi UI.
  12. Field discovery terekspos sebagai custom field dengan tipe skema yang terdokumentasi di endpoint metadata field.
  13. ⚠ Tidak ada endpoint publik untuk membuat/menghapus insight atau mengonfigurasi view pada versi awal.
  14. Webhook issue events dapat difilter dengan JQL `project = KEY AND issuetype = Idea` untuk hanya menerima event idea.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issues API di API-016, link API API-023, votes & watchers API API-024, katalog event webhook API-065 ([23](23-api-publik-webhook.md)). Penyediaan API publik insight/view dicatat sebagai pertanyaan terbuka.

### JPD-INS-038 JPD JQL Support

- **Deskripsi:** Idea adalah work item yang dapat dicari dengan JQL di pencarian lanjutan, saved filter, dan integrasi; filter view discovery tetap mekanisme terpisah.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator, Admin Situs, Integrator/Developer API; hasil mengikuti izin pembaca terhadap proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Kueri `project = KEY AND issuetype = Idea` mengembalikan idea di project discovery; `project.type = product_discovery` (JQL-017) mengembalikan idea di seluruh project discovery.
  2. Pencarian lanjutan (advanced search) atas idea dapat diekspor ke Excel/Google Sheets (JPD-INS-039).
  3. ⚠ Field discovery tertentu dapat dicari sebagai custom field (mis. select, number, label, date) dengan grammar dan operator per tipe mengikuti JQL-033; daftar field discovery yang diekspos ke searcher selaras dengan JPD-IDE-056.
  4. ⚠ Insight, Delivery progress, Delivery status, dan formula tidak dapat dicari dengan JQL; kueri yang memakainya mengembalikan galat "Field '<nama>' tidak dapat dicari".
  5. Saved filter berbasis JQL dapat dipakai sebagai sumber tab Microsoft Teams (JPD-INS-033) dan langganan filter.
  6. JQL tidak menggantikan filter view discovery; filter view (modul 36) tetap disimpan di view dan tidak diterjemahkan ke JQL.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Satu kemampuan dengan JPD-IDE-056 JPD Search & JQL Support ([36](36-jpd-ideas-field-views.md)), yang memuat aturan pencarian idea sebagai work type, field discovery yang didukung searcher, dan link ke delivery ticket; fitur ini hanya menambah `project.type = product_discovery`, ekspor hasil pencarian (JPD-INS-039), dan pemakaian saved filter di integrasi, dan aturan kedua heading berlaku bersama tanpa diimplementasikan dua kali. Pemilik tunggal grammar dan perilaku field discovery di JQL (referensi nama, `cf[id]`, operator per tipe, field yang tidak dapat dicari beserta pesan galatnya) adalah JQL-033 ([12](12-jql-bahasa-kueri.md)); modul 36 dan 37 hanya menyatakan field discovery mana yang diekspos ke searcher dan tidak mendefinisikan sintaks sendiri. JQL lain di modul 12 (JQL-017, JQL-025); searcher custom field CFG-038; saved filter SRC-036 ([13](13-pencarian-filter-ekspor.md)).

### JPD-INS-039 Data Export (CSV, Excel, Google Sheets)

- **Alias inventaris:** Import and export ideas
- **Deskripsi:** Ekspor data idea dari view ke CSV, ekspor ke Excel/Google Sheets melalui pencarian lanjutan, serta impor idea dari CSV dengan pemetaan kolom ke field discovery.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Admin Proyek/Project Lead (role Administrator; ekspor sesuai permission "Export a view"), Admin Situs (impor CSV global).
- **Aturan perilaku (acceptance criteria):**
  1. Menu view > Export > "Download CSV" mengunduh idea yang terlihat di view (mengikuti filter view) dengan kolom sesuai field yang tampil.
  2. Ekspor dikendalikan permission proyek "Export a view"; tanpa permission menu Export disembunyikan.
  3. Ekspor Excel dan Google Sheets dilakukan melalui pencarian lanjutan (JPD-INS-038) memakai mekanisme ekspor SRC-055.
  4. Field tambahan yang tidak tampil di view dapat ditambahkan ke ekspor melalui "+ More" sebelum mengunduh.
  5. CSV tidak menyertakan Reactions dan field Goals & Projects; kolom tersebut dihilangkan tanpa galat.
  6. Excel/Google Sheets tidak menyertakan Delivery progress, Delivery status, dan formula.
  7. Isi insight tidak dapat diekspor dari view; hanya jumlah insight yang disertakan, dan ekspor isi insight lengkap tersedia lewat permintaan ekspor admin (JPD-INS-037).
  8. ⚠ Impor CSV idea tersedia dari menu project discovery > Import (Creator/Administrator) dan dari importer CSV global Admin Situs (INT-008); idea hasil impor memakai work type Idea dan status awal workflow bila kolom status tidak dipetakan.
  9. Impor CSV: pengguna mengunggah file CSV lalu memetakan setiap kolom ke field discovery (summary, description, status, field select/number/date/label).
  10. Sebelum impor dijalankan, wizard menampilkan pratinjau baris pertama dan memvalidasi tipe nilai; baris tidak valid ditandai tanpa membatalkan baris lain.
  11. ⚠ Nilai select yang belum ada sebagai opsi dibuat otomatis sebagai opsi baru bila pengguna mencentang "Buat opsi baru"; bila tidak dicentang, nilai tersebut dilaporkan sebagai galat baris.
  12. Setelah impor selesai, laporan menampilkan jumlah idea dibuat, jumlah baris gagal, dan alasan per baris yang dapat diunduh sebagai CSV.
  13. ⚠ Ekspor CSV dari view menghasilkan file yang dapat diimpor ulang lewat wizard impor (aturan 9) tanpa mengubah pemetaan kolom untuk field dasar (summary, description, status, select, number, date, label).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ekspor CSV SRC-050 dan Excel/Google Sheets SRC-055 ([13](13-pencarian-filter-ekspor.md)); CSV importer INT-008 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JPD-INS-040 Jira Product Discovery data in Atlassian Analytics / Data Lake

- **Deskripsi:** Ketersediaan data discovery (idea, nilai field, insight, tautan delivery) di data lake produk ini sehingga dapat di-query dan divisualisasikan di modul analitik bersama data pelacak kerja dan Service Management.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Org admin dan admin analitik (koneksi data lake), Creator (pembaca dashboard).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Schema data lake memuat tabel discovery minimal: `discovery_ideas`, `discovery_insights`, `discovery_field_values`, dan `discovery_delivery_links`, dengan kolom kunci ke tabel work item dan project.
  2. Saat membuat atau mengedit koneksi data lake, admin memilih site dan project discovery yang disertakan.
  3. Data discovery dapat di-query dengan SQL maupun visual query editor dan di-join dengan data pelacak kerja (mis. idea ke work item delivery).
  4. ⚠ Tidak ada dashboard template khusus discovery pada versi awal; pengguna membuat dashboard sendiri.
  5. Izin data mengikuti pengaturan koneksi data lake dan akses analitik, bukan view restriction atau access level project discovery; UI koneksi menampilkan peringatan tentang hal ini.
  6. ⚠ Data diperbarui berkala oleh proses ETL data lake (bukan real-time), dengan waktu sinkron terakhir ditampilkan di koneksi.
  7. Data lake melengkapi Data Export (JPD-INS-039) karena menyertakan isi insight yang tidak tersedia di ekspor CSV/Excel/Sheets.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tersedia untuk plan Enterprise (JPD-INS-008); Analytics & Data Lake di RPT-038 s.d. RPT-041 ([16](16-laporan-insight.md)) dan INT-029 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JPD-INS-041 Rovo AI in JPD

- **Alias inventaris:** AI features (Atlassian Intelligence/Rovo) in JPD
- **Deskripsi:** Asisten AI opsional di editor dan panel discovery untuk brainstorm, menulis dan mengubah teks, serta merangkum insight, komentar, dan vote; tersedia pada plan Premium dan hanya bila diaktifkan admin.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (pemakai), org admin (mengaktifkan AI di pengaturan AI organisasi, ADM-052).
- **Aturan perilaku (acceptance criteria):**
  1. Fitur AI hanya tersedia pada plan Premium ke atas dan setelah org admin mengaktifkan AI; bila tidak, ikon AI disembunyikan.
  2. Brainstorm: ikon AI > Brainstorm menghasilkan daftar ide/sudut pandang dari konteks idea.
  3. Shortcut `/ai` di editor membuka menu perintah AI; produk ini tidak memakai alias bermerek vendor (mis. `/rovo`).
  4. Generate & transform: membuat draf, memperbaiki typo, menyederhanakan, dan mengubah tone teks terpilih.
  5. Summarize tersedia untuk description idea, idea template, komentar view, description view, description field, ringkasan vote, dan description insight.
  6. Hasil AI ditampilkan sebagai pratinjau yang dapat di-review, ditambahkan (Insert/Replace), atau dilanjutkan dengan prompt follow-up; tidak ada perubahan yang disimpan tanpa konfirmasi pengguna.
  7. Setiap panel hasil AI menampilkan disclaimer "Konten AI dapat tidak akurat; periksa sebelum digunakan".
  8. Ringkasan insight pada idea merangkum seluruh insight (description dan snippet) menjadi poin-poin tema utama.
  9. Generate/improve description idea dari summary dan insight yang ada.
  10. ⚠ Agen AI khusus product discovery (mis. agen yang menyarankan prioritas) tidak disediakan pada versi awal; asisten umum dipakai.
  11. ⚠ Deteksi idea serupa/duplikat tidak tersedia pada versi awal; kandidat fitur lanjutan.
  12. Ketersediaan mengikuti plan dan pengaturan AI admin; AI hanya mengakses data yang dapat dilihat pengguna yang memintanya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** AI editor di EDT-030 ([07](07-editor-konten-kaya.md)); pengaturan AI ADM-052 ([22](22-administrasi-sistem.md)). Penyedia model AI adalah keputusan terbuka; layanan AI lintas produk vendor (Rovo) tidak dibangun.

### JPD-INS-042 JPD Templates

- **Deskripsi:** Kumpulan template di Product Discovery: idea description template (juga untuk formulir Contributor), template otomasi proyek dan global, dan template proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator, Admin Proyek/Project Lead (role Administrator), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Idea description template dapat dibuat Administrator proyek dan dipilih saat membuat idea; template yang ditandai default dipakai juga pada formulir pengajuan idea Contributor (JPD-INS-004).
  2. Template otomasi tersedia di tingkat proyek dan global (JPD-INS-023).
  3. ⚠ Saat membuat project discovery, pengguna memilih template proyek (mis. Product roadmap, Customer feedback) yang menyiapkan field, view, dan idea type awal.
  4. Pada plan Premium, idea template dapat dirangkum/dihasilkan dengan AI (JPD-INS-041).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Idea Templates (Description Templates) dan pembuatan proyek di modul 36; template otomasi AUTO-010 ([20](20-otomasi.md)).

### JPD-INS-043 JPD Mobile Support

- **Deskripsi:** Product Discovery tidak didukung di app mobile native; akses mobile dilakukan melalui browser.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor.
- **Aturan perilaku (acceptance criteria):**
  1. App mobile native tidak menampilkan project discovery dan tidak ada rencana dukungan pada roadmap awal; daftar proyek di app mobile menyembunyikan proyek ber-type `product_discovery`.
  2. Bila idea dibuka di app mobile (mis. dari tautan notifikasi), app hanya menampilkan field dasar work item dan banner "Insights, delivery, view, dan field discovery tidak tersedia di app; buka di browser" dengan tombol buka di browser.
  3. SPA discovery dapat diakses dari browser mobile dengan tata letak responsif minimal (idea view dan list view dapat dibaca; editing view tidak dioptimalkan).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paritas app mobile di NAV-050 ([21](21-navigasi-ux-preferensi.md)).

### JPD-INS-044 Polaris Legacy Naming

- **Deskripsi:** Daftar nama internal legacy (codename "Polaris") yang masih terlihat di API, link type, dan field, dipertahankan untuk kompatibilitas dengan data/integrasi bergaya produk rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Integrator/Developer API, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Codename "Polaris" boleh muncul di identifier internal dan dokumentasi kompatibilitas, tetapi tidak pernah tampil di UI pengguna.
  2. Link type delivery memiliki nama internal kompatibilitas "Polaris work item link" (sebelumnya "Polaris issue link") dengan nama tampilan "implements / is implemented by" (JPD-INS-017).
  3. Istilah legacy "Data points" dipetakan ke "Insights" di UI, dokumentasi, dan impor dari sumber lama.
  4. ⚠ Key field discovery memakai prefix `jira.polaris` (mis. `jira.polaris:votes`) pada metadata custom field untuk kompatibilitas integrasi.
  5. ⚠ Project type key project discovery adalah `product_discovery`.
  6. Terminologi "project" telah diganti "space" (VISI-002); API tetap menerima dan mengembalikan istilah `project` untuk kompatibilitas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penamaan internal di modul 00 (VISI-002) dan API modul 23.

## Pertanyaan terbuka

- Apakah produk ini menamai ulang role produk Creator dan Contributor di glosarium (modul 00), atau memakai istilah rujukan apa adanya?
- Apakah Admin Situs dapat dikecualikan dari proyek Private (selalu dapat membuka), atau Private berlaku mutlak?
- Siapa default yang boleh membuat field di project discovery: hanya Administrator (pilihan saat ini) atau juga Creator?
- Bagaimana tier Free/Standard/Premium/Enterprise discovery dipetakan ke kebijakan harga produk ini, termasuk kuota Creator pada plan Free?
- Apakah app Intercom (JPD-INS-016) dibangun, atau cukup jalur ekstensi browser dan paste link?
- Apakah komponen penghubung ke instance pelacak kerja self-hosted eksternal (JPD-INS-024) dibangun, dan untuk produk self-hosted apa saja?
- Apakah produk ini menyediakan REST API publik untuk insight dan view (JPD-INS-037), dan kapan?
- Penyedia model AI mana yang dipakai untuk fitur AI discovery (JPD-INS-041), dan bagaimana data pelanggan dilindungi?
- Apakah publish view ke "siapa pun dengan tautan" (JPD-INS-029) diizinkan di semua plan berbayar atau hanya Premium ke atas?
- Apakah app Slack dan Microsoft Teams produk ini (JPD-INS-032, JPD-INS-033) dibangun sendiri pada Fase 6, atau menunggu app chat umum modul 25?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk Product Discovery (widget feedback, konektor survei, sinkronisasi roadmap pihak ketiga) tidak dibangun sebagai fitur native; produk ini hanya menyediakan API/webhook dan registry Smart Link agar ekosistem dapat terhubung.
- **Layanan lintas produk vendor rujukan**: Atlassian Home/Atlas, Confluence, Rovo, Atlassian Analytics/Data Lake, Atlassian Guard, dan customer feedback portal vendor tidak dibangun; padanannya adalah modul produk ini sendiri (Goals & Projects, dokumentasi, AI, data lake, keamanan organisasi, ekspor admin) bila ada.
- **Edisi Data Center vendor**: produk ini tidak memiliki edisi Data Center; integrasi ke instance self-hosted eksternal (JPD-INS-024) hanya opsional dan bergantung pada keputusan terbuka.
- **App mobile native untuk discovery** (JPD-INS-043) tidak dibangun; akses mobile lewat browser.
- **Agen AI khusus discovery dan deteksi duplikat AI** (JPD-INS-041) di luar cakupan versi awal.
- **API publik insight dan view** (JPD-INS-037) di luar cakupan versi awal.
- **Konektor helpdesk/CRM selain Zendesk, Salesforce, Intercom, dan Service Management** (mis. HubSpot, Freshdesk) hanya didukung lewat ekstensi browser dan paste link, bukan app khusus.
