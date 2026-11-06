# 36 — Product Discovery: Ideas, Fields, Formula & Views

Modul ini mendefinisikan inti **Product Discovery** di produk ini: tipe project discovery (discovery space) beserta template, pengaturan, dan level aksesnya; **Idea** sebagai work item utama (key, summary, status, idea type dan hierarki, pembuatan dari berbagai titik, form create, impor CSV, idea view, tata letak field, description template, workflow sederhana); siklus hidup dan kolaborasi idea (archive/restore, merge, move, bulk action, link dan Connection field, komentar, lampiran, watcher, history, reactions dan votes); **fields discovery** (manajemen field project, global fields, system fields, select, rating, slider, number/checkbox/hyperlink/text, date dan interval dengan fuzzy date, people/labels/team, goals, **formula** weighted score/custom/roll-up, pola prioritisasi RICE/ICE/MoSCoW/Now-Next-Later, field terhitung insights dan delivery); serta **views** sebagai navigasi utama (sidebar dan sections, List, Board, Matrix, Timeline, Tree, filter, sort, group, visibilitas field, save/autosave, template view, publish/share, stakeholder updates, ekspor, roadmap lintas project), ditambah pencarian/JQL, sandbox, pintasan keyboard, dan detail UI. Mesin umum (entitas project, work item, field storage, workflow engine, komentar, lampiran, notifikasi, JQL, CSV importer) dimiliki modul inti; modul ini mengatur **perilaku khusus discovery** di atasnya. Insights, delivery tickets, peran JPD (Creator/Contributor/Stakeholder), integrasi, dan API discovery dirinci di modul 37.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas project, work item, field, maupun view. Seluruh fitur modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/discovery` (project discovery, idea, idea type, template), `internal/features/discoveryfield` (field project/global, formula engine), dan `internal/features/discoveryview` (view, section, publish, roadmap lintas project), dengan entitas bersama di `internal/shared/domain/entity` (mis. `DiscoveryView`, `ViewSection`, `DiscoveryField`, `FieldOption`, `FormulaDefinition`, `PublishedView`) dan UI di `frontend/src/routes` (SPA SvelteKit). Idea disimpan sebagai baris work item biasa dengan `project_type_key='product_discovery'` agar pencarian, link, dan history inti dapat dipakai ulang.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah Work item/Space (VISI-002) dan tier plan sebagai feature flag (VISI-003).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — product access (PERM-004), level akses team-managed (PERM-038), role bawaan dan custom (PERM-039, PERM-040), akses anonim/publik (PERM-046).
- [04 Proyek](04-proyek.md) — entitas project (PRJ-001), template library (PRJ-005), wizard (PRJ-010), aturan key (PRJ-012), ubah key (PRJ-022), directory (PRJ-025), trash (PRJ-033), settings hub team-managed (PRJ-035).
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — field library team-managed (CFG-041), field sistem (CFG-011), Team field (CFG-019), description template (CFG-052).
- [06 Work Item Inti](06-work-item-inti.md) — key (ISS-001), create dialog (ISS-002), create another (ISS-004), issue view (ISS-007, ISS-008), navigasi (ISS-013), realtime (ISS-014), link (ISS-029), move (ISS-032), delete (ISS-033), archive (ISS-034), bulk change (ISS-035), watcher (ISS-036), votes (ISS-037).
- [07 Editor Teks Kaya](07-editor-konten-kaya.md) — editor ADF (EDT-001), tabel (EDT-010), panel (EDT-011), mention (EDT-014), smart link (EDT-021).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar (CMT-001 s.d. CMT-009), lampiran (CMT-011 s.d. CMT-018), history (CMT-022).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category (WF-002), editor workflow team-managed (WF-017).
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — referensi custom field (JQL-033), quick search (SRC-009), bulk change dari pencarian (SRC-031), ekspor CSV (SRC-050).
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Atlassian Teams (PLAN-057, PLAN-058), Atlassian Goals (PLAN-061, PLAN-062).
- [19 Notifikasi & Email](19-notifikasi-email.md) — watcher (NTF-012), mention (NTF-013), autowatch (NTF-014), notifikasi team-managed (NTF-007), hub in-app (NTF-015).
- [20 Otomasi](20-otomasi.md) — otomasi team-managed (AUTO-009), trigger field berubah (AUTO-021).
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — tema (NAV-034), pintasan keyboard (NAV-038 s.d. NAV-041), empty state (NAV-045), live update (NAV-046).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — Teams (ADM-028), Goals (ADM-029), sandbox (ADM-049).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — CSV importer (INT-008 s.d. INT-014), Slack (INT-030, INT-033), Confluence (INT-044), Goals & Projects Atlassian Home (INT-056), integrasi discovery ke delivery (INT-057).
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) dan [28 Model Data Inti](28-model-data-inti.md) — realtime (NFR-009), aksesibilitas (NFR-022), penyimpanan custom field (DM-005), event domain (DM-014).
- [37 Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) (prefix JPD-INS) — insights, delivery tickets dan progress, role Creator/Contributor/Stakeholder, integrasi dan API discovery.

## Persona terkait

- **Product Manager Discovery** — peran Creator berlisensi: membuat project discovery, mengelola idea, field, formula, views, roadmap, dan mempublikasikan view ke stakeholder.
- **Kontributor/Stakeholder Discovery** — peran Contributor (tanpa lisensi Creator) yang berkomentar, memberi reaction/vote, dan (bila diizinkan) membuat idea; serta Stakeholder yang membaca published view dan menerima stakeholder updates.
- **Admin Proyek/Project Lead** — admin project discovery: mengatur akses, field, idea type, workflow status, description template, form create, dan sections views.
- **Admin Situs** — memberi product access Product Discovery, mengelola global fields, roadmap lintas project, dan sandbox.
- **Product Owner/PM** — pemimpin produk yang memakai roadmap lintas project untuk menyelaraskan beberapa tim discovery.
- **Stakeholder/Viewer** — membaca roadmap terpublikasi, termasuk embed di Confluence, tanpa mengubah data.
- **Developer/Anggota Tim** — anggota tim delivery yang membuka idea dari work item tertaut dan memindahkan work item lama menjadi idea.
- **Integrator/Developer API** — mencari idea lewat JQL, mengimpor idea via CSV, dan memakai saved filter untuk bulk change.

## User story modul

1. Sebagai Product Manager Discovery, saya ingin membuat project discovery dari template prioritisasi atau roadmap agar fields, views, dan contoh idea langsung siap dipakai (JPD-IDE-001, JPD-IDE-002).
2. Sebagai Admin Proyek/Project Lead, saya ingin memilih level akses Open, Limited, atau Private agar idea sensitif hanya terlihat oleh orang yang tepat (JPD-IDE-004).
3. Sebagai Product Manager Discovery, saya ingin membuat idea langsung dari kolom board atau group dan mewarisi nilai kolom tersebut agar pencatatan ide cepat tanpa mengisi ulang field (JPD-IDE-007).
4. Sebagai Kontributor/Stakeholder Discovery, saya ingin mengajukan idea lewat form create yang sederhana agar masukan saya masuk ke backlog discovery tanpa butuh lisensi Creator (JPD-IDE-008).
5. Sebagai Product Manager Discovery, saya ingin membuat field Reach, Impact, Confidence, dan Effort serta formula RICE agar setiap idea punya skor prioritas objektif yang diperbarui otomatis (JPD-IDE-030, JPD-IDE-031, JPD-IDE-032, JPD-IDE-036, JPD-IDE-037).
6. Sebagai Product Manager Discovery, saya ingin memetakan idea di Matrix Impact vs Effort dan menyeret bubble untuk mengubah nilai agar trade-off nilai vs usaha terlihat jelas (JPD-IDE-044).
7. Sebagai Product Manager Discovery, saya ingin roadmap Now/Next/Later di Board dan roadmap Timeline dengan fuzzy date "Q3 2026" agar saya dapat berkomitmen pada arah tanpa tanggal palsu (JPD-IDE-029, JPD-IDE-033, JPD-IDE-043, JPD-IDE-045).
8. Sebagai Kontributor/Stakeholder Discovery, saya ingin memberi reaction atau vote dan berkomentar pada idea agar dukungan saya menjadi sinyal prioritas (JPD-IDE-020, JPD-IDE-024).
9. Sebagai Product Manager Discovery, saya ingin menggabungkan idea duplikat dan mengarsipkan idea yang tidak relevan agar backlog discovery tetap bersih tanpa kehilangan data (JPD-IDE-015, JPD-IDE-016).
10. Sebagai Admin Proyek/Project Lead, saya ingin menata views dalam sections di sidebar dan menyimpan konfigurasi filter, sort, dan group per view agar setiap audiens punya tampilan yang relevan (JPD-IDE-041, JPD-IDE-047, JPD-IDE-048, JPD-IDE-049).
11. Sebagai Stakeholder/Viewer, saya ingin membuka roadmap terpublikasi yang selalu live tanpa lisensi Creator agar saya tahu rencana produk terkini (JPD-IDE-052).
12. Sebagai Product Owner/PM, saya ingin roadmap lintas beberapa project discovery dengan global fields agar prioritas antar tim dapat dibandingkan dengan konsisten (JPD-IDE-027, JPD-IDE-055).
13. Sebagai Admin Proyek/Project Lead, saya ingin mendefinisikan idea type Opportunity dan Solution beserta hierarkinya dan melihatnya di Tree view agar discovery bertingkat termodelkan (JPD-IDE-006, JPD-IDE-046).
14. Sebagai Developer/Anggota Tim, saya ingin memindahkan work item lama ke project discovery atau mengimpornya via CSV agar permintaan fitur yang tersebar terkumpul sebagai idea (JPD-IDE-010, JPD-IDE-017).
15. Sebagai Integrator/Developer API, saya ingin mencari idea dengan JQL `project = KEY AND issuetype = Idea` agar data discovery dapat dipakai di filter, dashboard, dan bulk change (JPD-IDE-056).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 18 |
| P1 | 21 |
| P2 | 17 |
| P3 | 3 |
| **Total** | **59** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1-5 | 0 |
| Fase 6 | 59 |
| Fase 7 | 0 |
| Fase — | 0 |

Strategi urutan implementasi: seluruh modul ini adalah inti Fase 6 dan dikerjakan dalam tiga gelombang. Gelombang 6a (P0) membangun kerangka: tipe project discovery, pengaturan dan level akses (JPD-IDE-001, JPD-IDE-003, JPD-IDE-004), Idea work type, create, idea view, status (JPD-IDE-005, JPD-IDE-007, JPD-IDE-011, JPD-IDE-014), field management dengan tipe select, rating, date/interval, dan formula (JPD-IDE-025, JPD-IDE-029, JPD-IDE-030, JPD-IDE-033, JPD-IDE-036), serta views List, Board, Timeline dengan filter, sort, dan group (JPD-IDE-041 s.d. JPD-IDE-043, JPD-IDE-045, JPD-IDE-047, JPD-IDE-048) — cukup untuk roadmap Now/Next/Later dan scoring dasar. Gelombang 6b (P1) menambah template, form create, impor CSV, layout, description template, archive, bulk action, link, komentar (JPD-IDE-020, selaras JPD-INS-025), lampiran, watcher, reactions/votes, global fields, tipe field sisanya, Matrix, save/autosave, dan publish view. Gelombang 6c (P2-P3) menambah idea type dan hierarki beserta Tree view, merge, move, Slack, history lanjutan, goals, pola framework, template view, stakeholder updates, ekspor, roadmap lintas project, JQL lanjutan, sandbox, dan detail UI. Inventaris memuat banyak pasangan tumpang tindih antara lampiran JPD inti dan domain JPD rinci; 37 nama inventaris digabung ke heading yang lebih rinci melalui baris Alias inventaris sehingga 96 fitur jatah tercakup dalam 59 heading.

## Fitur

## Project discovery: pembuatan, template, pengaturan & akses

### JPD-IDE-001 Discovery Project Creation

- **Alias inventaris:** Product Discovery project type
- **Deskripsi:** Membuat project (discovery space) baru sebagai wadah ideas, fields, dan views. Project discovery memakai model konfigurasi mandiri mirip team-managed dan strukturnya berbeda dari project software/business: views menjadi navigasi utama dan tidak ada sprint/backlog.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (product access Product Discovery sebagai Creator) dengan global permission Create projects; Admin Situs; pembuat project otomatis menjadi admin project.
- **Aturan perilaku (acceptance criteria):**
  1. Wizard Create project (PRJ-010) menampilkan kategori "Product discovery" di template gallery sebagai titik masuk pembuatan project discovery.
  2. Langkah detail meminta Name (wajib) dan Key (dibentuk otomatis dari nama, dapat diedit, unik per site, hanya huruf kapital/angka, diawali huruf); tombol Create nonaktif selama ada field tidak valid dan key bentrok ditolak dengan pesan "Project key sudah dipakai".
  3. Pengguna memilih access level awal Open, Limited, atau Private (JPD-IDE-004); default Open.
  4. Pengguna memilih template (JPD-IDE-002) atau project kosong sebelum menekan Create.
  5. Avatar/ikon project dapat diubah setelah create dari Details (JPD-IDE-003).
  6. Pembuat project otomatis ditambahkan dengan role Administrator project.
  7. Validasi: hanya pengguna dengan akses produk Product Discovery sebagai Creator dan global permission Create projects (atau sesuai pengaturan admin) yang dapat membuat project; selain itu tombol Create tidak tampil dan API mengembalikan 403.
  8. Setelah create, pengguna diarahkan ke view pertama dengan empty state/onboarding berisi sample ideas dari template dan petunjuk langkah berikutnya.
  9. Project tampil di Projects directory (PRJ-025) dengan tipe "Product Discovery" dan ikon khusus, serta dapat difilter berdasarkan produk.
  10. Hapus project memindahkannya ke trash dan dapat di-restore dalam 60 hari (PRJ-033); setelah 60 hari project beserta ideas dihapus permanen.
  11. ⚠ Batas jumlah Creator pada plan Free dicek saat membuat atau mengelola project; bila kuota habis, pengguna tanpa seat Creator mendapat pesan yang menjelaskan batas plan.
  12. Konfigurasi project discovery dikelola per project oleh admin project/Creator (fields, workflow, views, idea types) dan tidak memakai scheme company-managed.
  13. Project memiliki key, name, avatar/ikon, project lead, dan description seperti project lain (PRJ-001).
  14. Sidebar project discovery berisi daftar Views dalam sections dan tautan Project settings, bukan Backlog/Board/Sprint (JPD-IDE-041).
  15. Work type default "Idea" dibuat otomatis saat project dibuat (JPD-IDE-005).
  16. ⚠ Menu Project settings memuat Details, Access/People, Fields, Idea types/Workflow, Description templates, Delivery, Features, Notifications, Automation (rinci di JPD-IDE-003).
  17. Lifecycle project (trash 60 hari, archive bila didukung) mengikuti lifecycle project inti.
  18. Project discovery muncul di Projects directory dan navigasi global dengan ikon Product Discovery.
  19. Hanya Creator berlisensi atau admin yang dapat membuat project discovery; Contributor yang mencoba mendapat pesan bahwa ia memerlukan lisensi Creator.
  20. Batasan: project discovery tidak memiliki sprint, backlog scrum, board software, estimasi story points, maupun releases/versions; endpoint dan menu terkait tidak tersedia untuk tipe project ini.
  21. Tipe project Product discovery di halaman Create project hanya tersedia bila produk Product Discovery aktif di site; bila tidak aktif, tipe ini tidak ditampilkan dan permintaan API pembuatan project bertipe discovery ditolak (403) dengan pesan bahwa Product Discovery belum diaktifkan di site.
- **Team-managed vs company-managed:** Project discovery selalu memakai model konfigurasi mandiri per project (mirip team-managed); tidak ada varian company-managed maupun share settings.
- **Catatan & ketergantungan:** Wizard inti PRJ-010, aturan key PRJ-012, trash PRJ-033, product access PERM-004. Sejak 2025 label UI "project" dapat ditampilkan sebagai "space" (VISI-002); label sebaiknya dapat dikonfigurasi. Dukungan archive project untuk discovery belum terkonfirmasi. Plan (Free/Standard/Premium/Enterprise) diterapkan sebagai feature flag (VISI-003). Migrasi: tabel `projects` belum ada; modul 04 menambahkannya dan modul ini menambah nilai `project_type_key='product_discovery'`.

### JPD-IDE-002 Discovery Project Templates

- **Alias inventaris:** JPD project templates
- **Deskripsi:** Template yang meng-seed project discovery dengan fields, views, sections, description template, dan contoh ideas sesuai use-case product management (roadmap, prioritisasi, feedback pelanggan, discovery).
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) saat create project; Admin Proyek/Project Lead untuk menambah set view/field ke project yang sudah ada.
- **Aturan perilaku (acceptance criteria):**
  1. Template bertema roadmap membuat Board Now/Next/Later berdasarkan select field Roadmap dan Timeline view.
  2. Template bertema prioritization membuat field Impact/Effort, Matrix Impact vs Effort, dan List yang diurutkan berdasarkan formula.
  3. Template bertema customer feedback/feature requests berfokus pada insights dan reactions (field Insights dan Reactions terlihat di views bawaan).
  4. ⚠ Template bertema discovery/opportunity membuat views untuk pemetaan opportunity dan solusi.
  5. Opsi project kosong membuat fields default minimal dan satu view List.
  6. Preview template menampilkan deskripsi, daftar views, dan fields yang akan dibuat sebelum dipilih; preview tidak membuat data.
  7. Sample ideas dapat dihapus setelah create, satu per satu atau sekaligus lewat bulk action.
  8. Fields default yang di-seed umumnya: Summary, Status, Assignee, Impact, Effort, Goals, Roadmap, Labels, Insights, Delivery progress, Delivery status.
  9. Semua konfigurasi hasil template (fields, views, sections, template description) dapat diubah atau dihapus setelah create; tidak ada yang terkunci kecuali system fields.
  10. Template Now/Next/Later: Board view dengan kolom dari select field Roadmap.
  11. Template Value vs Effort: Matrix view Impact (sumbu Y) vs Effort (sumbu X).
  12. Template RICE scoring: fields Reach, Impact, Confidence, Effort dan custom formula RICE.
  13. Template MoSCoW: select field dengan opsi Must, Should, Could, Won't.
  14. Template Timeline roadmap: Timeline view memakai Start date dan Target date.
  15. Template default/dasar memiliki set view minimal (satu List dan satu Board).
  16. ⚠ Galeri template dapat menambahkan set view/field ke project yang sudah ada; field yang sudah ada dengan nama sama dipakai ulang, bukan diduplikasi.
  17. Template customer feedback/insights-driven menyiapkan view List yang diurutkan menurun berdasarkan jumlah Insights sehingga idea dengan bukti pelanggan terbanyak tampil paling atas.
  18. Preview template di galeri menampilkan daftar views dan fields yang akan dibuat; menutup preview kembali ke galeri tanpa membuat project maupun mengubah pilihan template.
  19. Fields dan views hasil template tetap dapat diubah atau dihapus setelah project dibuat; menghapus view hasil template tidak menghapus field yang dipakainya, dan menghapus field hasil template menghapus kolom/kartu terkait dari views.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template disimpan sebagai data seed (JSON) agar daftar dapat diperbarui tanpa rilis kode; nama dan isi template di produk acuan berubah dari waktu ke waktu. Terkait template library PRJ-005 dan pre-built views JPD-IDE-051. Pasangan JPD-INS-042 JPD Templates ([37](37-jpd-insights-delivery-peran-integrasi.md)): modul 36 memiliki isi dan perilaku template project (fields, views, sections, sample ideas); modul 37 memiliki katalog template lintas jenis (description template untuk formulir Contributor, template otomasi, ringkasan AI) dan merujuk balik ke heading ini untuk template project.

### JPD-IDE-003 Project Settings (JPD)

- **Deskripsi:** Halaman pengaturan project discovery yang berbeda dari project software, berisi konfigurasi khusus discovery.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator project); Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Details: nama, key, avatar, deskripsi, project lead, dan kategori project dapat diubah; nama wajib dan key mengikuti aturan PRJ-012.
  2. Access: access level dan daftar people/groups beserta role (detail role di modul 37).
  3. Fields: daftar project fields dan global fields yang terasosiasi, dengan aksi create/edit/delete (JPD-IDE-025, JPD-IDE-027).
  4. Workflow/Statuses: kelola idea statuses (JPD-IDE-014).
  5. Types of ideas: kelola idea types dan hierarki (JPD-IDE-006; hanya plan Premium ke atas).
  6. Idea templates: kelola description templates (JPD-IDE-013).
  7. Create ideas settings: siapa boleh membuat idea dan field di form create (JPD-IDE-008).
  8. Halaman Delivery, Insights, Integrations, Notifications, dan Automation tersedia sebagai tautan ke modul terkait (modul 37, NTF-007, AUTO-009).
  9. ⚠ Halaman Features menyediakan toggle per project (mis. reactions, votes, contributor create) dan perubahan toggle langsung berlaku di semua views.
  10. Hanya admin project yang dapat membuka Project settings; pengguna lain tidak melihat tautannya dan URL langsung menampilkan halaman tanpa izin (403).
  11. Mengubah project key membuat key idea baru dengan awalan baru; key lama tetap me-redirect ke idea yang sama (PRJ-022).
- **Team-managed vs company-managed:** Struktur mengikuti settings hub team-managed (PRJ-035); tidak ada halaman scheme company-managed.
- **Catatan & ketergantungan:** Roles, insights, delivery, dan integrations dirinci di modul 37. Daftar toggle Features di produk acuan belum terkonfirmasi; clone menetapkan daftar sendiri dan menyimpannya sebagai `project_features` (key-value boolean).

### JPD-IDE-004 Project Access Levels (JPD)

- **Deskripsi:** Level akses project discovery yang menentukan siapa dapat melihat dan berkontribusi pada ideas dan views.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead mengubah access level dan keanggotaan; berlaku untuk semua pengguna site.
- **Aturan perilaku (acceptance criteria):**
  1. Open: semua pengguna site dengan akses Product Discovery dapat melihat dan berkontribusi sesuai lisensinya (Creator mengedit, Contributor berkomentar/react).
  2. Limited: semua pengguna site dengan akses Product Discovery dapat membuka project sebagai Contributor (melihat, berkomentar, react/vote); hanya people/groups yang ditambahkan secara eksplisit menjadi Creator atau Administrator dan dapat mengedit idea dan field (JPD-INS-002 butir 2).
  3. Private: hanya people/groups yang ditambahkan (dan Admin Situs) yang dapat membuka project; non-anggota mendapat HTTP 404 untuk URL project, idea, atau view, dan project tidak muncul di directory maupun pencarian (JPD-INS-002 butir 3). Published view dari project Private tetap mengikuti pengaturan publish (JPD-INS-029).
  4. Admin dapat menambah people/groups dengan role project Administrator, Creator, Contributor, atau custom role (JPD-INS-003); tidak ada role "Stakeholder" yang dapat ditugaskan (JPD-INS-030 butir 6).
  5. Perubahan access level berlaku langsung untuk semua views dan published views tanpa perlu menyimpan ulang view.
  6. Edge case: pengguna tanpa lisensi Creator tidak dapat mengedit field meskipun project Open; kontrol edit tidak ditampilkan dan API mengembalikan 403.
  7. Pada plan Free access level tidak dapat diubah dan tombol "Change project access" menampilkan ajakan upgrade (JPD-INS-002 butir 6).
- **Team-managed vs company-managed:** Mengikuti model level akses team-managed (PERM-038); tidak memakai permission scheme company-managed.
- **Catatan & ketergantungan:** Pasangan JPD-INS-002 Space Access Levels (Open / Limited / Private) ([37](37-jpd-insights-delivery-peran-integrasi.md)): modul 36 memiliki pemilihan access level saat create project (JPD-IDE-001) dan tampilannya di Project settings > Access (JPD-IDE-003); modul 37 memiliki perilaku izin runtime, dialog perubahan access level, dan batasan Contributor. Bila ada perbedaan aturan, JPD-INS-002 yang berlaku. Model akses team-managed PERM-038, PERM-039, PERM-040 ([03](03-izin-peran-keamanan.md)).

## Idea: work type, pembuatan, detail & workflow

### JPD-IDE-005 Idea Work Type

- **Alias inventaris:** Ideas
- **Deskripsi:** Work item utama project discovery bernama Idea yang merepresentasikan ide, opportunity, atau solusi produk, dengan key project, summary, status, dan fields discovery; tidak memakai work type software (story/bug). Fitur ini juga merangkum operasi dasar idea (create, edit, delete, watch, realtime, bulk, move) yang dirinci di heading lain.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) membuat dan mengedit; Kontributor/Stakeholder Discovery membuat bila diizinkan, berkomentar, dan react; Stakeholder/Viewer hanya membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap idea memiliki key unik dari project key dan counter berurutan per project (mis. DISC-12), memakai mekanisme key inti (ISS-001); nomor yang pernah dipakai tidak dipakai ulang.
  2. Summary wajib dan maksimal 255 karakter; summary kosong atau hanya spasi ditolak dengan pesan "Summary wajib diisi".
  3. Idea memakai ikon work type "Idea" yang khas discovery dan berbeda dari ikon work type software/business.
  4. Idea adalah work item biasa sehingga dapat dicari (JPD-IDE-056), di-link (JPD-IDE-019), dan dibuka dari project lain sesuai izin.
  5. Default satu work type "Idea" per project; tipe tambahan dan hierarki dikelola lewat JPD-IDE-006 (Premium).
  6. Idea tidak memiliki sub-tasks, sprint, story points, atau time tracking native; aksi dan field tersebut tidak ditampilkan dan API menolak penetapannya.
  7. System fields Created, Updated, Creator, Reporter, Assignee tersedia di setiap idea (JPD-IDE-028).
  8. Key tetap saat status berubah; key berubah bila idea dipindah ke project lain dan key lama me-redirect (JPD-IDE-017).
  9. Idea dapat dibuat lewat tombol Create di view, quick-create inline di List/Board/kolom (nilai field mengikuti grouping/kolom view), atau dialog Create global (rinci di JPD-IDE-007).
  10. Summary wajib; field lain opsional kecuali ditandai wajib di form create (JPD-IDE-008).
  11. Field dapat diedit inline dari view List, Board, Matrix, Timeline maupun dari idea view, sesuai tipe field.
  12. Delete idea bersifat permanen dengan dialog konfirmasi dan dibedakan dari Archive (JPD-IDE-015).
  13. Menu idea menyediakan Copy link, Open in new tab, serta membuka idea sebagai sidebar atau full page.
  14. Idea key mengikuti project key (mis. PD-12) dan ditampilkan di card, baris list, dan header idea view.
  15. Pengguna dapat watch idea dan menerima notifikasi (JPD-IDE-022).
  16. Perubahan oleh pengguna lain tampil langsung di view yang sedang terbuka tanpa reload (NAV-046, NFR-009).
  17. Contributor dapat membuat idea hanya bila fitur/permission diaktifkan di project (JPD-IDE-008).
  18. List view mendukung bulk actions: pilih banyak idea untuk edit field, archive, atau delete (JPD-IDE-018).
  19. ⚠ Idea dapat dipindah ke project discovery lain (JPD-IDE-017).
- **Team-managed vs company-managed:** Tidak ada perbedaan; Idea selalu work type milik project discovery.
- **Catatan & ketergantungan:** Memakai entitas work item inti ([06](06-work-item-inti.md)) dengan validasi khusus tipe project discovery. Multi idea type dipisah ke JPD-IDE-006.

### JPD-IDE-006 Idea Types & Hierarchies

- **Alias inventaris:** Idea types
- **Deskripsi:** Admin dapat mendefinisikan tipe idea kustom (mis. user problem, opportunity, bet, solution) dan menyusun hierarki parent-child di antaranya untuk memodelkan discovery bertingkat.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead mengelola tipe dan hierarki; Product Manager Discovery (Creator) menetapkan tipe dan parent pada idea.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman konfigurasi "Types of ideas" tersedia di Project settings bila plan mendukung.
  2. Admin dapat create, rename, dan delete tipe idea kustom dengan nama (wajib, unik per project) dan ikon.
  3. ⚠ Menghapus tipe yang sedang dipakai mewajibkan pemilihan tipe pengganti untuk semua idea terdampak sebelum penghapusan dijalankan.
  4. Admin menyusun hierarki antar tipe sebagai level parent-child (mis. Opportunity > Solution); tipe pada level bawah hanya dapat memiliki parent dari level tepat di atasnya.
  5. Creator dapat mengatur parent pada idea dan menambah child dari idea view; parent dengan tipe yang tidak sesuai hierarki ditolak.
  6. Idea type dipilih saat create idea (modal create dan inline create); default adalah tipe "Idea".
  7. ⚠ Creator dapat mengubah idea type pada idea yang sudah ada; field yang tidak berlaku pada tipe baru disembunyikan namun nilainya tidak dihapus.
  8. ⚠ Field dan layout idea view dapat diatur per tipe idea (JPD-IDE-012).
  9. Idea type dan Parent tersedia untuk filter, group, dan kolom di views.
  10. Hierarki divisualisasikan di Tree view (JPD-IDE-046) dan roadmap.
  11. ⚠ Jumlah level dan tipe per project dibatasi oleh batas yang dapat dikonfigurasi admin situs (default usulan: 5 tipe, 3 level).
  12. Default project hanya memiliki satu tipe "Idea".
  13. Admin dapat menambah, rename, mengganti ikon, dan menghapus idea type di project settings.
  14. Fields dapat dikaitkan per tipe sehingga form dan idea view hanya menampilkan field yang relevan.
  15. Views mendukung filter dan group by idea type.
  16. Idea type dipakai bersama Connection field untuk relasi opportunity ke solution (JPD-IDE-019).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hanya plan Premium/Enterprise; status GA di produk acuan belum terkonfirmasi. Mirip konsep hierarki work type CFG-007 tetapi berlaku hanya di dalam project discovery. Entitas `idea_types` (project_id, name, icon, level) dan kolom `parent_id` pada work item.

### JPD-IDE-007 Create Idea (in-product)

- **Deskripsi:** Membuat idea dari berbagai titik: tombol Create, inline di list, kolom board, group, atau Create global, dengan pewarisan nilai dari konteks view.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator); Kontributor/Stakeholder Discovery bila create ideas diizinkan (JPD-IDE-008).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Create idea" tersedia di header setiap view bagi pengguna yang berhak membuat idea.
  2. Inline create di baris bawah List view atau di dalam group: ketik summary lalu Enter membuat idea dan membuka baris kosong baru.
  3. Create di kolom Board membuat idea yang mewarisi nilai field kolom dan swimlane.
  4. Create di dalam group mewarisi nilai field group tersebut.
  5. Idea baru mewarisi nilai filter aktif bila memungkinkan; bila idea tetap tidak cocok dengan filter view, muncul pesan bahwa idea dibuat tetapi tidak terlihat di view ini beserta tautan untuk membukanya.
  6. Modal create memuat pemilihan project (dari Create global), idea type, summary, description dengan idea template, dan field tambahan sesuai create settings.
  7. Opsi "Create another" mempertahankan modal terbuka untuk membuat beberapa idea berturut-turut (ISS-004).
  8. Create global di top nav (ISS-002) dapat memilih project discovery dan menampilkan field form khusus idea.
  9. Menu + Create menyediakan entry point impor CSV (JPD-IDE-010).
  10. Permission create diperiksa per role: Contributor hanya dapat membuat bila diizinkan; bila tidak, tombol tidak tampil dan API mengembalikan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Create langsung di sel Matrix/Timeline belum terkonfirmasi di produk acuan; tidak masuk cakupan awal. Memakai ISS-002, ISS-005.

### JPD-IDE-008 Idea Creation Settings

- **Deskripsi:** Pengaturan siapa dan bagaimana idea dibuat, termasuk field yang tampil di form create untuk Creators dan Contributors.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead mengatur; Creator dan Contributor memakai form.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle "Contributors dapat membuat idea" (default off); bila off, Contributor tidak melihat tombol create.
  2. Admin memilih field yang tampil di form create idea dan urutannya (drag and drop).
  3. ⚠ Admin dapat menandai field tertentu wajib di form create; submit tanpa nilai ditolak dengan pesan per field.
  4. Admin memilih description template default untuk idea baru (JPD-IDE-013).
  5. ⚠ Admin dapat menambahkan teks bantuan/instruksi yang tampil di bagian atas form create.
  6. Status awal idea baru selalu mengikuti status awal workflow (JPD-IDE-014).
  7. ⚠ Contributor hanya dapat mengedit idea yang ia buat sendiri sesuai permission yang diberikan; idea lain read-only baginya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Form intake discovery di produk acuan berkembang cepat (nama "idea submission form" dan link form terpisah belum terkonfirmasi). Role Contributor rinci di modul 37.

### JPD-IDE-009 Create Idea from Slack

- **Deskripsi:** Membuat idea dari pesan Slack melalui app integrasi Slack first-party milik produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery dan Kontributor/Stakeholder Discovery yang akunnya sudah ter-link ke Slack dan berhak membuat idea.
- **Aturan perilaku (acceptance criteria):**
  1. Message action pada pesan Slack menyediakan aksi membuat idea di project discovery yang dipilih.
  2. Summary dan description terisi otomatis dari isi pesan dan dapat diedit sebelum submit.
  3. Pesan Slack dapat ditambahkan sebagai insight ke idea yang sudah ada (detail di modul 37).
  4. Idea yang dibuat menyimpan link balik ke pesan Slack asal dan ditampilkan di idea view.
  5. Syarat: app Slack terpasang di workspace dan akun pengguna ter-link (INT-030); bila belum, aksi menampilkan ajakan menghubungkan akun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada integrasi Slack Fase 4 (INT-030, INT-033). Perilaku create idea vs add insight di produk acuan belum terkonfirmasi; integrasi discovery rinci di modul 37.

### JPD-IDE-010 Bulk Import Ideas (CSV) & Import Existing Jira Work Items

- **Deskripsi:** Mengimpor banyak idea sekaligus dari file CSV dengan pemetaan kolom ke fields, termasuk membawa work item yang sudah ada sebagai idea.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator), Admin Proyek/Project Lead, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Entry point impor tersedia dari tombol + Create, menu project/view, dan importer CSV inti (INT-007).
  2. Pengguna mengunggah CSV dan memilih delimiter serta encoding (default koma, UTF-8).
  3. Kolom CSV dipetakan ke field discovery: Summary, Description, Status, select, number, labels, people, date.
  4. ⚠ Nilai opsi select dipetakan ke opsi yang ada; nilai yang belum ada dibuat sebagai opsi baru otomatis bila pengguna mencentang opsi tersebut.
  5. Header CSV hasil ekspor view (JPD-IDE-054) dapat dipakai sebagai template impor.
  6. Alur duplikasi work item software/service menjadi idea: ekspor CSV dari pencarian lalu impor ke project discovery (bulk move tidak memiliki opsi copy).
  7. Delivery work item dapat ditautkan saat impor lewat kolom key delivery dan key idea; link dibuat dengan link type delivery khusus (modul 37).
  8. Preview menampilkan contoh baris hasil pemetaan dan laporan akhir menampilkan jumlah berhasil/gagal beserta alasan per baris.
  9. Validasi: Summary wajib; baris tanpa Summary ditolak, nilai tidak valid pada field lain diabaikan dan dicatat di laporan.
  10. Hanya Creator/admin yang dapat menjalankan impor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai engine CSV importer inti (INT-008 s.d. INT-014) dengan profil pemetaan khusus discovery; belum terkonfirmasi apakah produk acuan memiliki importer khusus. Nama link type delivery di produk acuan ("Polaris work item link") diganti nama netral di clone.

### JPD-IDE-011 Idea View (Idea detail page)

- **Alias inventaris:** Idea view (idea detail)
- **Deskripsi:** Panel atau halaman detail idea yang dibuka dari view, berisi summary, fields, description, insights, delivery, komentar, dan history.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator mengedit; Contributor membaca, berkomentar, react; Stakeholder/Viewer read-only melalui published view/roadmap.
- **Aturan perilaku (acceptance criteria):**
  1. Idea dapat dibuka sebagai side panel dari view atau full page lewat URL berisi idea key (ISS-008).
  2. Saat panel terbuka, navigasi next/previous berpindah idea sesuai urutan view aktif.
  3. Header memuat key, idea type, summary (inline edit), status, watch, share/copy link, dan more actions (archive, merge, move, delete) sesuai izin.
  4. Bagian fields mendukung inline edit sesuai tipe field.
  5. Description memakai editor teks kaya (EDT-001) dengan table, panel, mention, emoji, code, dan media.
  6. Tab Insights, Delivery, Comments, dan History tersedia; isi tab Insights dan Delivery dirinci di modul 37.
  7. Bagian parent/child ideas tampil bila hierarki aktif (JPD-IDE-006, Premium).
  8. Panel dapat di-expand ke full screen dan ditutup dengan Esc.
  9. Field read-only (formula, insights, delivery) tampil tanpa kontrol edit.
  10. Mode sidebar dan full page memakai komponen yang sama sehingga perilaku identik.
  11. Description mendukung mention, tabel, gambar, smart links, embed, dan attachment.
  12. Tab Insights menampilkan daftar insights terlampir dan tombol Add insight.
  13. Tab Delivery menampilkan delivery tickets dan progress.
  14. Bagian Comments mendukung @mention dan reactions (JPD-IDE-020).
  15. History/Activity mencatat log perubahan field (JPD-IDE-023).
  16. Panel Fields menampilkan semua field idea dengan field pinned di atas dan opsi hide empty fields (JPD-IDE-012).
  17. Navigasi next/previous mengikuti urutan sort/rank view saat idea dibuka.
  18. Menu aksi memuat Archive, Merge, Delete, Copy link, dan Watch.
  19. Linked ideas/issues ditampilkan di idea view (JPD-IDE-019).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai kerangka issue view inti (ISS-007, ISS-009, ISS-013) dengan layout khusus discovery.

### JPD-IDE-012 Idea Layout & Field Sections

- **Alias inventaris:** Idea layout and pinned fields
- **Deskripsi:** Konfigurasi tata letak field di idea view: urutan, pin ke atas, dan field yang disembunyikan, agar field penting tampil lebih dulu.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead dan Creator dengan hak mengelola fields.
- **Aturan perilaku (acceptance criteria):**
  1. Field dapat di-pin ke bagian atas idea view (pinned fields).
  2. Urutan field diatur dengan drag and drop dan langsung tersimpan.
  3. Field dapat disembunyikan dari idea view tanpa menghapus field atau nilainya.
  4. Toggle "Show more"/sembunyikan field kosong mengatur apakah field tanpa nilai ditampilkan.
  5. Field description tampil sebagai help text/tooltip di samping label field.
  6. ⚠ Layout dapat diatur per idea type (terkait JPD-IDE-006).
  7. ⚠ Konfigurasi layout berlaku untuk semua pengguna project, bukan personal.
  8. Pin field berlaku untuk semua pengguna project.
  9. Urutan field pinned diatur dengan drag and drop.
  10. Field tertentu dapat disembunyikan dari idea view.
  11. Toggle hide empty fields tersedia di panel Fields.
  12. Pengaturan layout disimpan per project (bukan per user).
  13. Field tersembunyi tetap tersedia di views, filter, dan formula.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sections bernama (grup custom) di idea view belum terkonfirmasi. Mirip ISS-009 dan PRJ-048.

### JPD-IDE-013 Idea Templates (Description Templates)

- **Alias inventaris:** Idea templates
- **Deskripsi:** Template konten description idea untuk menstandarkan isi (problem, hypothesis, solution) saat membuat atau mengedit idea.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead mengelola; Creator dan Contributor menerapkan saat menulis description.
- **Aturan perilaku (acceptance criteria):**
  1. Template dikelola di Project settings: create, rename, edit konten rich text, delete.
  2. Template contoh bawaan tersedia (mis. problem statement, opportunity, requirements).
  3. Admin dapat menetapkan satu template default untuk idea baru.
  4. Pengguna dapat menerapkan template dari description kosong di idea view.
  5. Bila description sudah berisi, menerapkan template meminta konfirmasi replace; membatalkan tidak mengubah description.
  6. Template mendukung heading, list, table, panel, dan placeholder text.
  7. ⚠ Template dapat mempra-isi nilai field tertentu saat diterapkan pada idea baru.
  8. Project dapat memiliki beberapa description template.
  9. Template default dipakai otomatis untuk idea baru.
  10. Template dapat dipilih/diterapkan saat mengisi description.
  11. Template mendukung rich text, heading, tabel, dan panel.
  12. Template dapat diedit dan dihapus; menghapus template tidak mengubah description idea yang sudah memakainya.
  13. ⚠ Contoh bawaan meliputi Problem statement, Hypothesis, dan Success metrics.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penamaan UI di produk acuan berganti antara "description templates" dan "idea templates"; clone memakai "Idea templates". Mirip CFG-052.

### JPD-IDE-014 Idea Statuses & Workflow

- **Alias inventaris:** JPD workflow and statuses
- **Deskripsi:** Workflow sederhana khusus discovery untuk siklus hidup idea dengan editor status yang ringan.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead mengelola status; Creator mengubah status idea.
- **Aturan perilaku (acceptance criteria):**
  1. Status default contoh: Parking lot, Discovery, Ready for delivery, Delivery, Impact, Won't do (bervariasi per template).
  2. Admin dapat menambah, rename, hapus, dan reorder status; nama status wajib dan unik per project.
  3. Setiap status memiliki kategori To do/In progress/Done dengan warna lozenge (WF-002).
  4. Transisi default all-to-all: setiap status dapat berpindah ke setiap status lain.
  5. Status diubah via dropdown di idea view, drag antar kolom board (columns by Status), atau bulk edit.
  6. Status dapat dipakai sebagai field di filter, group, sort, dan kolom.
  7. Menghapus status yang dipakai mewajibkan pemilihan status pengganti untuk semua idea terdampak.
  8. ⚠ Editor workflow lanjutan (rules/validators) untuk discovery tersedia memakai editor team-managed (WF-017, WF-018).
  9. ⚠ Status default bawaan project kosong: Parking lot, Discovery, Ready for delivery, Delivery, Impact, Won't do.
  10. Status category dapat diatur saat menambah atau mengubah status.
  11. Transisi bebas antar status mengikuti model team-managed.
  12. Status dapat diubah dari board maupun idea view.
  13. Views dapat di-group dan difilter berdasarkan status.
  14. Automation dapat mentransisi idea saat delivery selesai melalui JPD-INS-023 Delivery Automation Templates & Rules ([37](37-jpd-insights-delivery-peran-integrasi.md)).
- **Team-managed vs company-managed:** Memakai engine workflow inti dalam mode team-managed sederhana (WF-017); tidak ada workflow scheme.
- **Catatan & ketergantungan:** Workflow engine inti di [09](09-workflow-status-transisi.md); discovery memakai editor sederhana.

## Siklus hidup & kolaborasi idea

### JPD-IDE-015 Archive & Restore Ideas

- **Alias inventaris:** Archive ideas
- **Deskripsi:** Mengarsipkan idea yang tidak relevan agar tersembunyi dari views tanpa menghapus datanya, dengan kemampuan restore; terpisah dari delete permanen.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) dan Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Archive tersedia dari more actions di idea view dan sebagai bulk action.
  2. Idea terarsip disembunyikan dari semua views secara default.
  3. Halaman/view "Archived ideas" di sidebar menampilkan idea terarsip.
  4. Restore mengembalikan idea ke views yang filternya cocok.
  5. ⚠ Fields Archived, Archived by, dan Archived on tercatat untuk audit dan dapat ditampilkan di view Archived ideas.
  6. ⚠ Archive tidak mengubah status workflow idea.
  7. Idea terarsip tetap dapat dicari dan di-link.
  8. Delete permanen adalah aksi terpisah dengan dialog konfirmasi yang menegaskan tidak dapat di-undo.
  9. Permission: archive dan delete hanya untuk Creator/admin; Contributor tidak melihat aksinya.
  10. Archive dapat dilakukan dari idea view atau bulk di List.
  11. Archived ideas tersembunyi dari views default.
  12. View khusus Archive dapat dipakai untuk melihat dan restore.
  13. Restore/unarchive mengosongkan field Archived, Archived by, dan Archived on.
  14. Fields Archived, Archived by, Archived on terisi otomatis saat archive.
  15. Hanya Creator/admin yang dapat archive.
  16. Archived ideas tetap dapat dicari via JQL.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme internal (flag vs status) di produk acuan belum terkonfirmasi; clone memakai flag `archived_at`/`archived_by` pada work item, selaras ISS-034.

### JPD-IDE-016 Merge Ideas

- **Alias inventaris:** Merge ideas
- **Deskripsi:** Menggabungkan idea duplikat ke satu idea target sambil membawa data penting dari idea sumber.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) dengan hak edit pada idea sumber dan target.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Merge dari idea view membuka pencarian idea target berdasarkan summary/key; idea sumber tidak dapat dipilih sebagai target.
  2. ⚠ Pengguna memilih data yang disalin: field values, insights, delivery tickets, attachments, comments.
  3. ⚠ Idea sumber diarsipkan dan di-link ke target dengan link type "merged into".
  4. Dialog konfirmasi ditampilkan sebelum merge dijalankan.
  5. ⚠ Merge lintas project discovery diperbolehkan bila pengguna berhak di kedua project.
  6. Riwayat merge tercatat di History idea sumber dan target.
  7. Pengguna memilih satu idea target.
  8. Insights, delivery tickets, dan nilai field (opsional) dipindahkan ke target.
  9. Idea sumber diarsip setelah merge.
  10. Link antar idea dicatat.
  11. Dialog konfirmasi berisi ringkasan item yang akan dipindah (jumlah insights, delivery tickets, field).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail granularitas opsi di produk acuan belum terkonfirmasi. Operasi dijalankan dalam satu transaksi; kegagalan membatalkan seluruh merge.

### JPD-IDE-017 Move Idea Between Projects

- **Deskripsi:** Memindahkan idea dari satu project discovery ke project discovery lain, atau memindahkan work item yang sudah ada ke project discovery menjadi idea.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator dengan hak di project asal dan tujuan; Admin Proyek/Project Lead; Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Move tersedia dari more actions idea view.
  2. Pengguna memilih project discovery tujuan dari daftar project yang ia berhak buat idea.
  3. Status dan nilai field yang tidak ada di project tujuan dipetakan secara eksplisit; global fields tetap dibawa tanpa pemetaan.
  4. Key berubah sesuai project tujuan dan key lama me-redirect.
  5. ⚠ Insights dan delivery links ikut dipindah bersama idea.
  6. Bulk move work item yang ada ke project discovery: pencarian > Bulk change > Move ke project discovery, dengan pemetaan work type ke Idea dan pemetaan field (ISS-035).
  7. Batasan: bulk move tidak memiliki opsi copy; untuk duplikasi gunakan ekspor/impor CSV (JPD-IDE-010).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai wizard move inti ISS-032 dan bulk operations ISS-035. Perilaku move antar project discovery di produk acuan belum terkonfirmasi.

### JPD-IDE-018 Idea Bulk Edit & Bulk Actions

- **Deskripsi:** Memilih banyak idea di view dan menerapkan aksi sekaligus.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Checkbox seleksi tersedia di List view, termasuk select all dalam view atau dalam group.
  2. ⚠ Range selection dengan Shift-click memilih semua baris di antara dua baris.
  3. Toolbar aksi tampil dengan jumlah idea terpilih.
  4. Bulk edit nilai field (select, multi-select, labels, people, number, date, rating); untuk field multi-value tersedia mode tambah, ganti, dan hapus.
  5. Bulk ubah status ke status mana pun.
  6. Bulk archive dan bulk delete meminta konfirmasi dengan jumlah idea terdampak.
  7. ⚠ Bulk move ke project lain tersedia (JPD-IDE-017).
  8. Bulk change via bulk operations inti dari hasil pencarian tersedia sebagai titik integrasi (SRC-031).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Operasi massal dijalankan sebagai job asinkron bila jumlah idea melebihi ambang (mis. 100) dengan progres dan laporan per item.

### JPD-IDE-019 Link Ideas & Connection Field

- **Alias inventaris:** Connection field; Linked ideas
- **Deskripsi:** Menghubungkan idea dengan idea lain atau work item lain melalui issue links standar dan field Connection untuk relasi terstruktur.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator membuat/menghapus link dan mengisi Connection; Contributor membuat link bila berhak edit.
- **Aturan perilaku (acceptance criteria):**
  1. Issue link (relates to, blocks, duplicates, dll.) dapat ditambahkan ke idea atau work item project lain (ISS-029).
  2. Delivery tickets memakai link type khusus delivery (detail di modul 37).
  3. ⚠ Field Connection menyimpan relasi antar idea/idea type (mis. Opportunity ke Solutions).
  4. Linked issues tampil di idea view dan dapat menjadi kolom/field di view.
  5. Link dapat dihapus (remove link).
  6. ⚠ Views dapat difilter dan di-group berdasarkan connection.
  7. Connection field dapat memilih item terhubung di project yang sama atau project discovery lain.
  8. Item terhubung ditampilkan di cards dan List.
  9. Group by item terhubung menghasilkan tampilan hierarki.
  10. Klik item terhubung menavigasi ke item tersebut.
  11. Idea dapat di-link ke idea lain dengan relates to, blocks, duplicates.
  12. Daftar linked issues tampil di idea view.
  13. Link dapat di-unlink.
  14. Link dibuat otomatis saat merge ideas (JPD-IDE-016).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue linking inti ISS-029. Hierarki parent-child formal ada di JPD-IDE-006. Connection field adalah fitur 2025 di produk acuan; cakupan dan namanya belum terkonfirmasi.

### JPD-IDE-020 Comments on Ideas

- **Deskripsi:** Diskusi berbasis komentar pada idea yang terbuka bagi Contributors maupun Creators.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor; admin project menghapus komentar orang lain.
- **Aturan perilaku (acceptance criteria):**
  1. Komentar rich text mendukung mention @user, emoji, dan attachment inline (CMT-001).
  2. Pengguna dapat edit/delete komentar sendiri; admin dapat menghapus komentar orang lain.
  3. Emoji reactions dapat diberikan pada komentar (CMT-006).
  4. ⚠ Balasan/thread komentar didukung (CMT-005).
  5. Urutan komentar dapat diubah newest first/oldest first dan pilihan diingat per pengguna.
  6. Watchers dan user yang di-mention menerima notifikasi.
  7. Jumlah komentar tersedia sebagai field yang dapat ditampilkan di view.
  8. Contributor dapat berkomentar tanpa lisensi Creator.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai engine komentar inti CMT-001 s.d. CMT-009. Tumpang tindih dengan JPD-INS-025 Comments & Reactions on Ideas ([37](37-jpd-insights-delivery-peran-integrasi.md)) yang memuat field dinamis Comments/Reactions dan aturan published view; kedua fitur diimplementasikan sebagai satu kemampuan dan prioritas disamakan menjadi P1 (inventaris domain ini mencatat P0). Pembagian kepemilikan: modul 36 memiliki perilaku komentar pada idea view (urutan, thread, field jumlah komentar di view); modul 37 memiliki izin kolaborasi, notifikasi mention, dan batasan published view; bila ada perbedaan aturan, JPD-INS-025 yang berlaku. Reactions pada idea dikonfigurasi di JPD-IDE-024.

### JPD-IDE-021 Attachments on Ideas

- **Deskripsi:** Lampiran file/gambar pada idea melalui description, komentar, atau area attachments.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Contributor (sesuai izin edit/komentar).
- **Aturan perilaku (acceptance criteria):**
  1. Upload via drag-and-drop atau paste ke description, komentar, atau area attachments.
  2. Preview gambar, PDF, dan video di media viewer (CMT-015).
  3. Attachment dapat di-download dan di-delete oleh pengunggah atau admin.
  4. Batas ukuran file mengikuti setting attachment site (CMT-016); file melebihi batas ditolak dengan pesan berisi batas maksimum.
  5. ⚠ Attachment ikut disalin saat merge bila opsi attachments dipilih (JPD-IDE-016).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan attachment inti CMT-011 s.d. CMT-019.

### JPD-IDE-022 Watchers & Notifications (Ideas)

- **Deskripsi:** Mengikuti (watch) idea untuk menerima notifikasi perubahan, komentar, dan insight baru.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor, Viewer dengan akses baca ke idea; menambah watcher lain butuh izin manage watchers.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Watch/Unwatch di idea view menampilkan jumlah watchers.
  2. Pengguna dengan permission dapat menambah watcher lain.
  3. ⚠ Pembuat idea dan pemberi komentar otomatis menjadi watcher (NTF-014).
  4. Notifikasi email/in-app dikirim untuk komentar, perubahan status, mention, insight baru, dan perubahan delivery.
  5. Pengaturan notifikasi project discovery memetakan event ke penerima (NTF-007).
  6. Notifikasi in-app tampil di bell notifikasi global (NTF-015).
- **Team-managed vs company-managed:** Memakai notifikasi team-managed (NTF-007).
- **Catatan & ketergantungan:** NTF-012, NTF-013, NTF-014. Detail halaman notifikasi discovery di produk acuan belum terkonfirmasi. Event insight dan delivery berasal dari modul 37.

### JPD-IDE-023 Idea History & Activity

- **Deskripsi:** Jejak audit perubahan pada idea.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator, Contributor, Viewer dengan akses baca.
- **Aturan perilaku (acceptance criteria):**
  1. Tab History menampilkan perubahan field, status, summary, dan description dengan nilai lama/baru (CMT-022).
  2. Create, archive/restore, merge, link, dan move dicatat sebagai entri history.
  3. Setiap entri menampilkan timestamp dan aktor.
  4. ⚠ Aktivitas dapat difilter antara Comments dan History.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model changelog inti CMT-023; nilai formula tidak dicatat di history karena terhitung.

### JPD-IDE-024 Reactions & Votes Fields

- **Alias inventaris:** Votes field; Reactions field
- **Deskripsi:** Konfigurasi field Reactions dan Votes di project discovery (menambahkan field ke project, menampilkan nilainya di views, dan memakainya di view internal serta formula) agar pengguna menyatakan dukungan pada idea lewat emoji reactions (mis. 👍) atau votes terhitung. Fitur ini adalah gabungan eksplisit dengan JPD-INS-026 Votes Field dan JPD-INS-025 Comments & Reactions on Ideas ([37](37-jpd-insights-delivery-peran-integrasi.md)): modul 36 memiliki konfigurasi field dan tampilan view; modul 37 memiliki aturan izin, vote budget, kolaborasi, dan batasan published view.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead (role Administrator) menambah field Votes dan menetapkan vote budget (JPD-INS-026); Creator dan Contributor memberi reaction/vote; Stakeholder/Viewer pada published view hanya membaca dan tidak dapat react/vote.
- **Aturan perilaku (acceptance criteria):**
  1. Field Reactions: pengguna menambah emoji pada idea dan jumlah per emoji ditampilkan.
  2. Satu pengguna satu reaksi per emoji; klik lagi membatalkan reaksi.
  3. Tooltip menampilkan daftar pengguna yang bereaksi.
  4. Field Votes menyimpan hitungan dukungan per idea sebagai nilai numerik (total vote dari semua pengguna); satu pengguna dapat memberi lebih dari satu vote per idea selama vote budget tersisa (JPD-INS-026 butir 3).
  5. Contributor dapat react/vote tanpa lisensi Creator.
  6. Nilai reactions dan votes dapat dipakai di sort, filter, dan formula pada view internal; pengecualian untuk published view diatur butir 20.
  7. Nilai tampil di card board dan kolom list.
  8. React/vote tidak dapat dilakukan dari published view: kontrol reaction dan vote disembunyikan bagi pembaca published view, dan API menolak permintaan react/vote melalui token published view dengan HTTP 403 (JPD-INS-026 butir 2).
  9. Field Votes ditambahkan ke project dari panel Fields.
  10. Vote budget (jumlah vote per pengguna per field Votes, minimal 1) ditetapkan Administrator proyek saat membuat atau mengedit field; vote melebihi budget ditolak dengan pesan "Kuota vote Anda telah habis". Aturan lengkap budget dan perubahan kuota dimiliki JPD-INS-026 butir 1 dan 3.
  11. Contributor dapat vote tanpa lisensi Creator.
  12. Total votes tampil di views dan dapat dipakai untuk sort/formula di view internal (butir 20 untuk published view).
  13. Pengguna yang berhak dapat melihat siapa saja yang memberi vote beserta jumlah vote masing-masing (ringkasan vote, JPD-INS-026 butir 5).
  14. Pengguna dapat mencabut atau mengurangi vote miliknya; vote yang dicabut mengembalikan sisa budget.
  15. Emoji reaction dapat ditambahkan dari view maupun idea view.
  16. Jumlah reactions ditampilkan per emoji.
  17. Daftar user yang bereaksi dapat dilihat.
  18. Contributor dapat bereaksi.
  19. Reactions ditampilkan di cards dan List.
  20. View yang memakai field Votes atau Reactions untuk sort, filter, group, kolom board, atau sumbu matrix tidak dapat di-publish; dialog Publish menyebut field penyebabnya (JPD-INS-029 butir 10, JPD-INS-025 butir 5, JPD-INS-026 butir 6).
  21. Field Reactions tidak disertakan pada ekspor CSV view (JPD-INS-025 butir 5, JPD-IDE-054).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Gabungan eksplisit: aturan izin vote, vote budget, dan ringkasan vote dimiliki JPD-INS-026 Votes Field ([37](37-jpd-insights-delivery-peran-integrasi.md)); field dinamis Comments/Reactions dan kolaborasi dimiliki JPD-INS-025 Comments & Reactions on Ideas ([37](37-jpd-insights-delivery-peran-integrasi.md)); batasan published view dimiliki JPD-INS-029 Publish Views ([37](37-jpd-insights-delivery-peran-integrasi.md)). Modul ini hanya memiliki konfigurasi field (katalog, panel Fields) dan tampilan di views; bila ada perbedaan aturan, aturan modul 37 yang berlaku. Model data votes/watchers DM-008; votes discovery berbeda dari votes work item inti ISS-037 karena mendukung budget dan disimpan per (idea, field, user, jumlah).

## Fields discovery: manajemen, katalog & field sistem

### JPD-IDE-025 Field Management (Project Fields)

- **Alias inventaris:** Field management
- **Deskripsi:** Membuat dan mengelola custom fields khusus discovery di project (tipe, nama, deskripsi, emoji, opsi) dari panel Fields di view maupun Project settings.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) dengan hak manage fields; Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Panel Fields tersedia di sidebar view dan di Project settings, dilengkapi search berdasarkan nama field.
  2. Create field: pilih tipe, isi nama (wajib, unik per project), deskripsi, dan emoji.
  3. Edit field: rename, ubah deskripsi, kelola options.
  4. Delete field meminta konfirmasi dan menghapus nilainya di semua idea; system fields tidak dapat dihapus.
  5. ⚠ Tipe field tidak dapat diubah setelah dibuat; untuk mengganti tipe, buat field baru lalu hapus yang lama.
  6. Sebelum delete, sistem menampilkan views dan formula yang memakai field; formula yang merujuk field terhapus ditandai invalid.
  7. ⚠ Jumlah field per project dibatasi (default usulan 100 field custom, dapat dikonfigurasi admin situs); melebihi batas ditolak dengan pesan batas.
  8. ⚠ Hanya Creator dengan hak admin project/manage fields yang dapat mengelola fields.
  9. Field dapat dibuat dari sidebar Fields di view atau dari Project settings.
  10. Nama, field description (tooltip), dan emoji/ikon field dapat diedit.
  11. Opsi select/labels dapat ditambah, rename, dihapus, diurutkan ulang, diberi warna, dan emoji.
  12. Menghapus opsi meminta konfirmasi; nilai idea yang memakai opsi ikut dikosongkan.
  13. Visibilitas field diatur per view (kolom/card) tanpa memengaruhi view lain.
  14. Delete field permanen setelah konfirmasi; data field hilang di semua idea.
  15. Field bawaan template dapat diubah seperti field custom.
  16. Daftar field menampilkan indikator field global vs field project.
- **Team-managed vs company-managed:** Mirip field library team-managed (CFG-041): field dimiliki project, tidak ada field context/configuration scheme.
- **Catatan & ketergantungan:** Penyimpanan nilai memakai model custom field DM-005. Duplikasi field tidak didukung karena belum terkonfirmasi di produk acuan.

### JPD-IDE-026 Custom field types (JPD)

- **Alias inventaris:** Jira Product Discovery fields
- **Deskripsi:** Katalog tipe field yang dapat dibuat Creator di project discovery untuk menangkap data prioritisasi dan atribut idea; setiap tipe dirinci di heading tersendiri.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator), Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Number untuk nilai seperti Reach atau revenue (JPD-IDE-032).
  2. Rating skala 1-5 ditampilkan sebagai dots/bars (JPD-IDE-030).
  3. Slider 0-100 (JPD-IDE-031).
  4. Checkbox (boolean) (JPD-IDE-032).
  5. Select (single select) dengan opsi berwarna dan emoji (JPD-IDE-029).
  6. Labels/Multi-select dengan opsi berwarna (JPD-IDE-029, JPD-IDE-034).
  7. People untuk satu atau banyak pengguna (JPD-IDE-034).
  8. Date dengan presisi fuzzy day, month, quarter (JPD-IDE-033).
  9. Short text dan Hyperlink (JPD-IDE-032).
  10. Formula: Weighted score, Custom formula, Roll-up (JPD-IDE-036).
  11. Votes dan Reactions (JPD-IDE-024).
  12. Atlassian goals, Atlassian Project, dan Team field (JPD-IDE-034, JPD-IDE-035).
  13. ⚠ Connection field (JPD-IDE-019).
  14. Tipe field tidak dapat diubah setelah dibuat; ganti dengan menghapus dan membuat ulang.
  15. ⚠ Jumlah field per project dibatasi (JPD-IDE-025).
  16. Formula fields mendukung weighted score dan roll-up.
  17. Rating, Slider, Checkbox, Votes, dan Insights count tersedia sebagai tipe/field discovery.
  18. Delivery progress dan Delivery status tersedia sebagai field terhitung (JPD-IDE-038).
  19. Field discovery dimiliki per project dan tidak muncul di halaman Custom fields global (CFG-034), kecuali global fields (JPD-IDE-027).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tipe "Atlassian Project" dan "Atlassian goals" di clone dinamai netral (Project platform, Goals platform) dan bergantung pada INT-056 dan PLAN-061.

### JPD-IDE-027 Global Fields

- **Alias inventaris:** Global fields
- **Deskripsi:** Fields yang didefinisikan sekali di level site dan dipakai lintas project discovery untuk konsistensi data dan cross-project views.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs membuat/mengedit; Admin Proyek/Project Lead mengasosiasikan ke project.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Global fields dikelola di pengaturan admin situs bagian Product Discovery.
  2. Create global field: tipe, nama (unik per site), deskripsi, options.
  3. Global field dapat diasosiasikan ke project atau dilepas; melepas field menyembunyikan nilainya dari project tanpa menghapus data selama field masih ada.
  4. Options dikelola terpusat; project tidak dapat mengubah opsi.
  5. Nilai dari project field dapat disalin ke global field (migrasi) dengan pemetaan opsi.
  6. Daftar fields menampilkan ikon pembeda global vs project field.
  7. Hanya Admin Situs yang dapat membuat/mengedit global field.
  8. Global fields menjadi dasar filter/group/sort pada cross-project roadmaps (JPD-IDE-055).
  9. Global field dibuat oleh admin di pengaturan global Product Discovery.
  10. Admin project menambahkan global field ke project tertentu.
  11. Opsi select/labels global dikelola terpusat dan terkunci di level project.
  12. Nilai konsisten untuk roadmaps/cross-project views.
  13. ⚠ Field project dapat dikonversi menjadi global field.
  14. ⚠ Tidak semua tipe field tersedia sebagai global; formula, reactions, dan votes tetap per project.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ketersediaan per plan belum terkonfirmasi (kemungkinan semua plan, sedangkan roadmap lintas project Premium). Entitas `discovery_fields` dengan `scope` (project|global).

### JPD-IDE-028 Time Stamp & System Fields

- **Alias inventaris:** System fields (JPD)
- **Deskripsi:** Field bawaan yang tersedia di setiap project discovery, sebagian dihitung otomatis dan read-only.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Semua pengguna berhak baca; nilai yang dapat diedit mengikuti izin edit idea.
- **Aturan perilaku (acceptance criteria):**
  1. Created dan Updated (time stamp) tersedia untuk sort dan filter relatif.
  2. Key, Summary, Status, dan Idea type tersedia di semua views.
  3. Parent tersedia bila hierarki idea types aktif (Premium).
  4. Creator, Reporter, dan Assignee tersedia sebagai system people fields.
  5. Comments count dan Linked issues tersedia sebagai field.
  6. ⚠ Fields arsip Archived, Archived by, dan Archived on tersedia.
  7. System fields tidak dapat dihapus, hanya disembunyikan dari view.
  8. Daftar system fields: Summary, Description, Status, Assignee, Reporter, Creator, Labels, Created, Updated, Key, Issue type.
  9. Insights (jumlah insights) bersifat read-only.
  10. Delivery progress (bar progres dari delivery tickets) bersifat read-only.
  11. Delivery status (status agregat delivery) bersifat read-only.
  12. Comments count dan Linked issues dapat ditampilkan di view.
  13. ⚠ Fields Archived, Archived by, Archived on terisi otomatis oleh aksi archive.
  14. Field Project platform menautkan idea ke project di Atlassian Home padanan clone (INT-056).
  15. System field tidak dapat dihapus; sebagian dapat disembunyikan dari views.
  16. Field computed tidak dapat diedit manual; API menolak penulisan dengan galat validasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog field sistem inti CFG-011, ISS-015.

## Fields discovery: tipe nilai

### JPD-IDE-029 Select & Multi-select Fields

- **Deskripsi:** Field pilihan tunggal dan ganda dengan opsi berwarna untuk kategorisasi idea (mis. Roadmap, Area, Team, Goals lokal).
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Admin Proyek/Project Lead mengelola opsi; Creator mengisi nilai.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi dapat ditambah, rename, reorder (drag), dan dihapus; nama opsi wajib dan unik dalam field.
  2. Setiap opsi memiliki warna.
  3. Setiap opsi dapat diberi weight numerik untuk Weighted score (JPD-IDE-036).
  4. Urutan opsi menentukan urutan kolom board/group dan urutan sort.
  5. Menghapus opsi yang dipakai mengosongkan nilai pada idea terdampak setelah konfirmasi yang menampilkan jumlah idea.
  6. ⚠ Opsi baru dapat dibuat langsung dari idea view saat mengisi nilai.
  7. Nilai tampil sebagai lozenge berwarna di card dan list.
  8. Multi-select di board/group: idea muncul di beberapa kolom; drag menambah atau memindah nilai sesuai mode (JPD-IDE-043).
  9. Field Roadmap (Now/Next/Later) default memakai tipe single select.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Emoji per opsi belum terkonfirmasi; clone mendukung emoji opsional. Mirip CFG-028, CFG-037.

### JPD-IDE-030 Rating Field (Impact/Effort)

- **Alias inventaris:** Impact and Effort fields
- **Deskripsi:** Field skala 1-5 yang diisi dengan klik ikon, dipakai untuk penilaian cepat seperti Impact atau Effort dan menjadi dasar Matrix value vs effort.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Nilai yang valid adalah kosong atau bilangan bulat 1..5; nilai lain ditolak API.
  2. Template prioritization membuat field default Impact dan Effort bertipe rating.
  3. Nilai rating dapat di-clear kembali ke kosong.
  4. Rating dapat dipakai sebagai sumbu Matrix, sort, filter, dan input formula.
  5. Rating tampil sebagai dots/bars di card dan kolom list.
  6. Impact berupa rating 1-5.
  7. Effort berupa rating 1-5.
  8. Impact dan Effort dipakai sebagai sumbu X/Y Matrix view.
  9. Impact dan Effort menjadi input weighted score/custom formula.
  10. Field dapat di-rename; mengganti tipe tampilan dilakukan dengan membuat ulang field.
  11. Nilai diedit langsung dengan klik di List atau drag di Matrix.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bentuk ikon persis belum terkonfirmasi; clone memakai dots.

### JPD-IDE-031 Slider Field

- **Deskripsi:** Field numerik 0-100 yang diisi via slider, cocok untuk confidence atau persentase.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Nilai berupa bilangan bulat 0-100; input di luar rentang ditolak dengan pesan rentang.
  2. Nilai diisi dengan drag slider atau mengetik angka.
  3. Nilai tampil sebagai bar di list/card.
  4. Slider dapat dipakai sebagai sumbu Matrix, sort, filter, dan formula.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kustom min/max kemungkinan tidak didukung di produk acuan; clone tidak menyediakannya pada fase awal.

### JPD-IDE-032 Number, Checkbox, Hyperlink & Text Fields

- **Deskripsi:** Tipe field dasar untuk angka, boolean, URL, dan teks.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Number menerima integer/desimal dan dipakai di formula, sort, dan filter range.
  2. Checkbox bernilai true/false, dapat di-toggle di list dan card, dan difilter checked/unchecked.
  3. Hyperlink menyimpan URL dengan validasi format (http/https) dan tampil sebagai link yang dapat diklik.
  4. Short text satu baris (maks 255 karakter) mendukung filter contains.
  5. ⚠ Long text/paragraph sebagai custom field tersedia; bila tidak, teks panjang hanya di Description.
  6. Summary dan Description adalah system text fields.
  7. Semua tipe ini mendukung inline edit dari sel list.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format tampilan number (percent/currency) belum terkonfirmasi; lihat JPD-IDE-039. Mirip CFG-025, CFG-026, CFG-031.

### JPD-IDE-033 Date & Interval Fields

- **Alias inventaris:** Date fields (fuzzy dates)
- **Deskripsi:** Field tanggal tunggal dan rentang (interval) yang menjadi dasar Timeline/roadmap, dengan dukungan presisi tidak pasti (hari, bulan, kuartal).
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Date diisi dengan date picker.
  2. Interval menyimpan rentang start-end dalam satu field.
  3. Granularitas fuzzy: day, month, quarter (mis. "Q3 2026").
  4. Nilai dapat diset dengan drag/resize bar di Timeline (JPD-IDE-045).
  5. Filter tanggal mendukung before/after/between dan relatif (mis. kuartal ini).
  6. ⚠ Sort tanggal: fuzzy date diurutkan berdasarkan awal periode untuk ascending dan akhir periode untuk tie-break.
  7. Nilai dapat di-clear.
  8. Field Start date dan Target date tersedia sebagai pasangan default.
  9. Pengguna memilih presisi specific day, month, atau quarter per nilai.
  10. Nilai ditampilkan sebagai teks sesuai presisi, mis. "Q3 2026" atau "Jun 2026".
  11. Field tanggal dapat dipakai sebagai sumbu Timeline view.
  12. Ideas dapat di-sort dan difilter berdasarkan tanggal.
  13. ⚠ Target date sebelum start date ditolak dengan pesan "Target date tidak boleh sebelum start date".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama field default di produk acuan belum terkonfirmasi. Disimpan sebagai `{start, end, precision}` agar fuzzy date dapat dibandingkan. Terkait CFG-022.

### JPD-IDE-034 People, Labels & Atlassian Team Fields

- **Alias inventaris:** Team field (Atlassian teams)
- **Deskripsi:** Field untuk menghubungkan idea dengan orang, label bebas, dan tim platform (padanan Atlassian Teams di clone).
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. People field memakai user picker single/multi dan menampilkan avatar.
  2. Assignee, Reporter, dan Creator adalah system people fields.
  3. Labels bersifat multi-value dengan autocomplete label yang ada dan pembuatan label baru.
  4. Team field memilih team dari direktori Teams platform (PLAN-057).
  5. Board dan list dapat di-group/kolom berdasarkan people, labels, dan team.
  6. Filter "Me" (current user) tersedia untuk people fields.
  7. Team dipilih dari directory teams.
  8. Views dapat di-group/difilter berdasarkan team.
  9. Team field konsisten dengan Team field di work item dan Plans (CFG-019, PLAN-058).
  10. Team dapat diganti atau dihapus dari idea.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Teams adalah fitur platform (ADM-028). Rename/delete label terpusat belum terkonfirmasi.

### JPD-IDE-035 Goals Field (Atlassian Goals)

- **Alias inventaris:** Goals and themes
- **Deskripsi:** Mengaitkan idea dengan goals bisnis dan tema strategis, baik field Goals select lokal maupun Goals platform, agar roadmap selaras strategi.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator); Contributor membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Field Goals bawaan template berupa select/multi-select lokal dengan weight per opsi.
  2. ⚠ Field Goals platform menautkan idea ke goal di modul Goals platform (PLAN-061).
  3. ⚠ Status/progress goal (on track/at risk/off track) tampil di idea view.
  4. Views dapat di-group/difilter berdasarkan goal.
  5. Field Goals platform menghubungkan idea ke goal di Atlassian Home padanan clone.
  6. ⚠ Status goal ditampilkan sebagai lozenge di card dan list.
  7. Tema strategis dimodelkan dengan select/labels field (mis. Theme, Strategic pillar).
  8. Board, List, dan Timeline dapat di-group by goal/theme.
  9. Views dapat difilter berdasarkan goal/theme.
  10. Weighted score dapat memberi bobot berdasarkan opsi theme.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Goals lintas produk adalah platform (ADM-029, INT-056, PLAN-062). Integrasi dan tier belum terkonfirmasi.

## Formula & prioritisasi

### JPD-IDE-036 Formula Fields

- **Alias inventaris:** Formula fields
- **Deskripsi:** Field terkomputasi read-only yang menghitung skor atau agregasi dari field lain; dasar prioritization frameworks dan fitur paling khas discovery.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) dan Admin Proyek/Project Lead membuat/mengedit formula; semua pembaca melihat hasil.
- **Aturan perilaku (acceptance criteria):**
  1. Weighted score: pilih field (rating, slider, number, select dengan option weight, insights count) dan beri bobot positif/negatif.
  2. ⚠ Roll-up mengagregasi nilai dari insights, delivery tickets, atau field lain dengan operasi sum/count/avg.
  3. Custom formula/expression mendukung aritmatika + - * / dan kurung yang mereferensikan field dan konstanta, mis. Reach*Impact*Confidence/Effort.
  4. Editor formula menyediakan autocomplete nama field dan preview hasil pada contoh idea.
  5. Warning tampil bila field yang dirujuk dihapus atau tidak valid; hasil formula menjadi kosong sampai diperbaiki.
  6. Nilai hasil dapat negatif.
  7. ⚠ Hasil dibulatkan sesuai pengaturan desimal formula (default 2 desimal).
  8. Formula dapat dipakai di sort, filter, sumbu Matrix, dan kolom list.
  9. ⚠ Pembagian dengan nol atau field kosong menghasilkan nilai kosong, bukan galat.
  10. Weighted score dinormalisasi sehingga skor sebanding antar idea.
  11. Opsi select memiliki nilai numerik untuk weighted score.
  12. Custom formula mendukung operator + - * / dan kurung, referensi field, dan konstanta (contoh RICE = Reach * Impact * Confidence / Effort).
  13. ⚠ Roll-up mendukung count insights serta sum/avg nilai dari insights atau delivery tickets.
  14. Preview hasil tampil saat konfigurasi.
  15. Formula read-only di idea dan diperbarui otomatis saat input berubah.
  16. ⚠ Nilai kosong pada input diperlakukan sebagai kosong sehingga hasil kosong, kecuali weighted score yang memperlakukannya sebagai 0.
  17. Formula dapat dipakai untuk sort, filter, sumbu, dan ukuran bubble Matrix.
  18. Referensi siklik antar formula ditolak saat menyimpan dengan pesan yang menyebut rantai field.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine formula di backend (parser ekspresi aman, tanpa eval) dengan perhitungan ulang inkremental saat event field berubah (DM-014). Roll-up insights/delivery bergantung modul 37.

### JPD-IDE-037 Prioritization Framework Patterns (RICE/ICE/Weighted)

- **Alias inventaris:** Prioritization frameworks (RICE, MoSCoW, Now/Next/Later)
- **Deskripsi:** Pola prioritisasi berbasis skor (RICE/ICE, Value vs Effort) dan kategorikal (MoSCoW, Now/Next/Later) yang dibangun dari fields, formula, dan views; tidak ada engine khusus.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator), Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. RICE: Reach (number), Impact (rating), Confidence (slider), Effort (number/rating) ditambah custom formula.
  2. ICE: Impact, Confidence, Ease ditambah formula.
  3. Value vs Effort: weighted score dengan Effort berbobot negatif.
  4. ⚠ Saat create formula, pengguna dapat memilih template formula RICE, ICE, atau Value vs Effort yang membuat field input yang belum ada.
  5. Penjelasan framework dituliskan via field description (JPD-IDE-039).
  6. RICE juga mendukung Impact bertipe number dan Confidence bertipe persen.
  7. Pola RICE dapat dimodifikasi menjadi ICE atau skor lain.
  8. List view dapat di-sort berdasarkan skor dengan visualisasi bar.
  9. Select field Roadmap memiliki opsi Now, Next, Later (opsional Won't do).
  10. Select field MoSCoW memiliki opsi Must have, Should have, Could have, Won't have.
  11. Board view berkolom per opsi; drag idea antar kolom mengubah nilai.
  12. Warna opsi dipakai untuk konsistensi visual antar view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sebagian besar berupa panduan; wizard RICE native belum terkonfirmasi. Bergantung pada JPD-IDE-029 s.d. JPD-IDE-036.

### JPD-IDE-038 Insights & Delivery Computed Fields

- **Deskripsi:** Field sistem otomatis untuk jumlah insights serta status/progress delivery tickets terhubung, ditampilkan di views.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Semua pembaca idea; nilai dihitung sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Insights menampilkan jumlah insights; klik membuka tab Insights di idea view.
  2. Delivery progress menampilkan progress bar berdasarkan status category delivery tickets.
  3. Delivery status menampilkan ringkasan status delivery.
  4. Field ini read-only namun dapat ditampilkan, di-sort, difilter, dan dipakai di roll-up.
  5. ⚠ Basis perhitungan progress dapat dipilih antara jumlah ticket dan estimasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Logika insights dan delivery dimiliki modul 37; modul ini hanya mendefinisikan perilaku sebagai field di views.

### JPD-IDE-039 Field Descriptions & Display Formatting

- **Deskripsi:** Deskripsi field dan cara nilai ditampilkan agar penilaian konsisten.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Admin Proyek/Project Lead dan Creator.
- **Aturan perilaku (acceptance criteria):**
  1. Field description tampil sebagai tooltip di idea view dan header kolom.
  2. Field dapat memiliki emoji yang tampil di samping nama.
  3. Warna opsi select dan lozenge status dipakai konsisten di semua view.
  4. Rating, slider, dan progress bar ditampilkan secara visual.
  5. Header kolom list menampilkan ikon tipe field.
  6. ⚠ Opsi format tampilan nilai (mis. compact/teks penuh, persen) dapat dipilih per field.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Terkait JPD-IDE-025 dan JPD-IDE-032.

## Views: model, navigasi & tipe view

### JPD-IDE-040 Views (general)

- **Deskripsi:** View adalah cara menampilkan subset ideas dengan konfigurasi filter, sort, grouping, dan fields; inti navigasi project discovery dan disimpan bersama untuk semua pengguna project.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator membuat dan mengubah view; Contributor dan Stakeholder membuka view sesuai akses.
- **Aturan perilaku (acceptance criteria):**
  1. Tipe view yang didukung: List, Board, Matrix, Timeline (ditambah Tree bila hierarki aktif, JPD-IDE-046).
  2. Creator dapat create, rename, duplicate, dan delete view.
  3. View memiliki emoji dan description (JPD-IDE-050).
  4. View disimpan bersama dan sama untuk semua pengguna project.
  5. ⚠ Perubahan konfigurasi memakai autosave per view atau banner unsaved changes dengan Save/Discard (JPD-IDE-049).
  6. Search ideas tersedia di dalam view (summary/key).
  7. Filter, sort, dan group dikonfigurasi per view.
  8. Visibilitas field dikonfigurasi per view.
  9. Jumlah ideas yang cocok dengan view ditampilkan di header.
  10. Real-time collaboration: perubahan idea dan konfigurasi view oleh pengguna lain tampil tanpa reload.
  11. ⚠ Pengguna dapat berkomentar pada view untuk diskusi tentang roadmap.
  12. ⚠ Jumlah view per project dibatasi (default usulan 200); melebihi batas ditolak dengan pesan batas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `discovery_views` (project_id, type, name, emoji, config JSONB, section_id, rank). Realtime memakai NFR-009.

### JPD-IDE-041 Views Navigation & Sections (Sidebar)

- **Alias inventaris:** View sections
- **Deskripsi:** Sidebar project discovery yang mengorganisasi views dalam sections (folder) untuk navigasi.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator dan Admin Proyek/Project Lead mengelola sections dan views; semua anggota melihat sidebar.
- **Aturan perilaku (acceptance criteria):**
  1. Views dikelompokkan dalam sections yang dapat di-collapse.
  2. Section dapat dibuat, rename, dihapus, dan diurutkan ulang.
  3. View dapat di-drag antar sections dan diurutkan ulang.
  4. Create view: pilih tipe List/Board/Matrix/Timeline/Tree (Tree hanya bila hierarki aktif).
  5. View dapat ditambahkan dari pre-built views (beta) atau view templates (JPD-IDE-051).
  6. View dapat diduplikasi beserta seluruh konfigurasinya.
  7. View dapat di-rename dan diberi emoji.
  8. Delete view meminta konfirmasi; ideas tidak terhapus.
  9. Search views berdasarkan nama di sidebar.
  10. Sidebar memuat tautan ke Archived ideas, Fields, dan Project settings.
  11. Section dibuat dengan nama dan emoji.
  12. View dapat di-drag ke section; views dan sections dapat diurutkan ulang.
  13. ⚠ Section dapat di-rename dan dihapus; menghapus section memindahkan views di dalamnya ke luar section (tidak menghapus view).
  14. Section dapat di-collapse/expand.
  15. Urutan views dan sections berlaku untuk semua pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Default/landing view belum terkonfirmasi; clone membuka view terakhir yang dibuka pengguna, atau view pertama. Sidebar memakai kerangka NAV-005.

### JPD-IDE-042 List View

- **Alias inventaris:** List view
- **Deskripsi:** View tabel untuk melihat, mengedit inline, dan meranking idea dengan kolom field.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator mengedit; Contributor dan Viewer membaca (Contributor dapat react/komentar).
- **Aturan perilaku (acceptance criteria):**
  1. Kolom field dapat ditambah, dihapus, diurutkan ulang, dan diubah lebarnya.
  2. Setiap sel dapat diedit inline sesuai tipe field dan izin.
  3. Kolom Summary sticky saat scroll horizontal.
  4. Manual ranking dengan drag baris tersedia saat sort manual.
  5. Group by field dengan collapse dan count per group.
  6. Sort satu/multi field ascending/descending.
  7. Inline create di bawah list atau group.
  8. Row selection untuk bulk actions (JPD-IDE-018).
  9. Klik baris membuka idea panel.
  10. Virtualized scrolling menjaga performa untuk ribuan idea.
  11. Kolom dapat ditambah/dihapus/diurutkan dan diubah lebar.
  12. Inline edit tiap sel.
  13. Kolom Summary selalu sticky.
  14. Rows dapat di-group berdasarkan field dengan collapse.
  15. Sort per field atau manual ranking drag-and-drop.
  16. Bulk select untuk edit, archive, delete.
  17. Nilai divisualisasikan (rating dots, slider bar, delivery progress bar).
  18. Quick create baris baru.
  19. Navigasi keyboard antar sel (panah, Enter untuk edit, Esc untuk batal).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ranking manual memakai LexoRank per view (DM-010). Komponen tabel dapat dipakai ulang dari SRC-023.

### JPD-IDE-043 Board View

- **Alias inventaris:** Board view
- **Deskripsi:** View kanban dengan kolom berdasarkan nilai field (mis. Roadmap atau Status) dan card yang dapat di-drag.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator memindahkan card dan mengonfigurasi; Contributor dan Viewer membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Columns by field: Status, select, multi-select, people, labels, team.
  2. Kolom "No value" menampung idea tanpa nilai.
  3. Drag card antar kolom mengubah nilai field; untuk multi-select tersedia perilaku tambah (Ctrl/Alt) atau pindah (default).
  4. Swimlanes (group by) memakai field kedua.
  5. Kolom dapat diurutkan ulang dan disembunyikan.
  6. Kolom/swimlane kosong dapat disembunyikan.
  7. Urutan card dalam kolom dapat diatur manual.
  8. Card layout: pilih fields yang tampil di card.
  9. Idea dapat dibuat dari kolom.
  10. Setiap kolom menampilkan jumlah idea.
  11. Field kolom dipilih dari select, status, people, atau labels.
  12. Drag card antar kolom mengubah nilai field.
  13. Swimlanes berdasarkan field kedua.
  14. Fields pada card dapat dikonfigurasi.
  15. Kolom dapat diurutkan ulang dan kolom kosong disembunyikan.
  16. Kolom "No value" tersedia.
  17. Urutan card manual atau berdasarkan sort.
  18. Quick create di kolom.
  19. Filter dan search berlaku pada board.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Board Now/Next/Later adalah use-case roadmap utama. Board discovery tidak memiliki WIP limit, sprint, atau column-status mapping (berbeda dari BRD di [10](10-board-kolom-kartu.md)).

### JPD-IDE-044 Matrix View

- **Alias inventaris:** Matrix view
- **Deskripsi:** Grafik 2D (mis. Impact vs Effort) yang memetakan idea sebagai bubble berdasarkan field sumbu.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator men-drag bubble dan mengonfigurasi; Contributor dan Viewer membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Field sumbu X dan Y dipilih dari rating, slider, number, formula, atau select.
  2. Ukuran bubble berdasarkan field ketiga (mis. insights, reactions).
  3. ⚠ Warna bubble berdasarkan field select.
  4. Drag bubble mengubah nilai field sumbu, hanya untuk field yang editable; sumbu formula tidak dapat di-drag.
  5. Idea bertumpuk digabung dengan count; klik menampilkan daftar idea.
  6. Hover menampilkan summary, klik membuka idea.
  7. ⚠ Arah sumbu dapat dibalik (reverse).
  8. ⚠ Label kuadran (mis. quick wins, big bets) dapat ditampilkan dan diedit.
  9. Filter view tetap berlaku.
  10. Sumbu X dan Y mendukung number, rating, slider, select, formula.
  11. Ukuran bubble berdasarkan field ketiga (mis. Insights).
  12. Warna bubble berdasarkan field select.
  13. Drag idea mengubah nilai field sumbu.
  14. Arah sumbu dapat dibalik.
  15. Hover/klik bubble membuka detail idea.
  16. Ideas tanpa nilai sumbu ditampilkan terpisah di panel "Tanpa nilai".
  17. Ideas bertumpuk digabung dalam satu bubble dengan jumlah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sumbu select memakai urutan opsi (JPD-IDE-029).

### JPD-IDE-045 Timeline View (Roadmap)

- **Alias inventaris:** Timeline view
- **Deskripsi:** View roadmap horizontal yang menempatkan idea pada waktu berdasarkan field Date/Interval.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator menjadwalkan; Contributor dan Viewer membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih field start/target date atau satu field Interval.
  2. Skala waktu weeks/months/quarters dapat di-zoom.
  3. Rentang tampilan timeline dapat diatur.
  4. Drag bar memindah, resize mengubah durasi; perubahan menyimpan nilai field tanggal.
  5. Swimlanes/group by field (mis. Team, Goal).
  6. Daftar idea tanpa tanggal tersedia untuk dijadwalkan dengan drag ke timeline.
  7. Today marker ditampilkan.
  8. Bar menampilkan field terpilih.
  9. ⚠ Markers/milestones dapat ditambahkan ke timeline.
  10. ⚠ Dependencies antar idea tidak ditampilkan pada fase awal.
  11. Field start dan target date dapat dipilih.
  12. Skala waktu months atau quarters.
  13. Drag dan resize bar mengubah tanggal.
  14. Rows dapat di-group berdasarkan field (team, goal, status).
  15. Panel ideas belum dijadwalkan tersedia.
  16. Today marker tampil.
  17. Fields pada bar dapat dikonfigurasi.
  18. Rentang waktu tampilan dapat diatur.
  19. Fuzzy date ditampilkan sesuai presisi (bar mencakup seluruh bulan/kuartal); drag pada skala quarter menyimpan presisi quarter.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari timeline work item (modul 18); tidak ada scheduling sprint/release.

### JPD-IDE-046 Tree View

- **Deskripsi:** Tipe view yang menampilkan idea dalam struktur pohon berdasarkan hierarki idea types (parent-child), lengkap dengan field dan filter.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator membuat dan mengubah; Contributor dan Viewer membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Tree view dapat dibuat dari sidebar hanya bila hierarki idea types dikonfigurasi (JPD-IDE-006).
  2. ⚠ Navigasi expand/collapse per level serta expand all/collapse all.
  3. Relasi parent-child antar tipe idea ditampilkan (mis. Opportunity > Solution).
  4. Field/kolom yang tampil di tree view dapat diatur.
  5. ⚠ Filter berlaku di tree view; parent tetap tampil (redup) bila ada child yang cocok filter.
  6. ⚠ Sort berlaku di dalam setiap level.
  7. Klik node membuka idea panel.
  8. ⚠ Parent dapat diubah dengan drag node atau menambah child dari node.
  9. Konfigurasi view tersimpan, dan share/publish mengikuti aturan view lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kemungkinan hanya plan Premium/Enterprise. Bergantung pada JPD-IDE-006.

## Views: konfigurasi & perilaku

### JPD-IDE-047 View Filters

- **Alias inventaris:** View filters
- **Deskripsi:** Penyaringan idea dalam view berdasarkan nilai field dan kata kunci.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Creator menyimpan filter view; Contributor dan Viewer dapat memfilter sementara.
- **Aturan perilaku (acceptance criteria):**
  1. Filter per field dengan operator sesuai tipe: is, is not, contains, is empty, >, <, between.
  2. Beberapa filter multi-field digabung dengan AND.
  3. Filter tanggal relatif (mis. bulan ini, kuartal depan).
  4. Quick search teks pada summary/key.
  5. Filter "Me" untuk people fields.
  6. Filter berdasarkan Idea type.
  7. Indikator jumlah filter aktif dan tombol Clear all.
  8. Filter dapat dibuat pada field apa pun (select, labels, people, number range, rating, date, status, idea type).
  9. Operator is, is not, is empty, is not empty, dan range tersedia.
  10. Filter current user tersedia untuk people fields.
  11. Kombinasi beberapa filter memakai AND.
  12. ⚠ Viewer dapat menerapkan filter sementara tanpa menyimpan view (JPD-IDE-049).
  13. Archived ideas disembunyikan secara default.
  14. Text search tersedia dalam view.
  15. Jumlah filter aktif ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Filter view bukan JQL (JPD-IDE-056). Logika OR antar filter belum terkonfirmasi; tidak masuk cakupan awal.

### JPD-IDE-048 View Sort, Group & Field Visibility

- **Alias inventaris:** View sorting and ranking; View grouping; Card configuration
- **Deskripsi:** Pengaturan urutan (sort dan manual ranking), pengelompokan, field yang tampil, dan konfigurasi card di view.
- **Prioritas / Fase / Tier:** P0 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Sort by field ascending/descending atau manual rank.
  2. Group by (list), Columns + Swimlanes (board/timeline).
  3. Group dapat di-collapse/expand dan group kosong disembunyikan.
  4. Fields dapat ditampilkan/disembunyikan per view (kolom/card).
  5. Card layout dapat dikonfigurasi.
  6. Field dapat di-drag dari panel Fields ke view.
  7. Konfigurasi tersimpan per view.
  8. Sort ascending/descending per field.
  9. Multi-level sort (mis. Score desc lalu Created asc).
  10. Manual ranking drag-and-drop disimpan per view.
  11. Pengguna dapat beralih antara sort by field dan manual; beralih ke manual mempertahankan urutan terakhir sebagai rank awal.
  12. Sort berdasarkan formula score tersedia.
  13. Group by select, labels, people, status, goal, team, connection.
  14. Sub-grouping/swimlanes tersedia.
  15. Group dapat di-collapse/expand.
  16. Urutan group mengikuti urutan opsi field.
  17. Group "No value" menampung idea tanpa nilai.
  18. Idea multi-label muncul di lebih dari satu group.
  19. Pengguna memilih fields yang tampil di card.
  20. Urutan fields di card dapat diatur.
  21. ⚠ Warna/aksen card dapat ditentukan oleh field select.
  22. Konfigurasi card berlaku per view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi disimpan di `config` JSONB view (JPD-IDE-040).

### JPD-IDE-049 View Changes: Save & Autosave

- **Deskripsi:** Perubahan konfigurasi view tidak langsung memengaruhi pengguna lain sampai disimpan, kecuali autosave aktif.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Creator menyimpan; Contributor dan Viewer hanya mengubah tampilan lokal.
- **Aturan perilaku (acceptance criteria):**
  1. Banner unsaved changes tampil saat filter/sort/group diubah.
  2. Save menyimpan untuk semua pengguna; Reset membatalkan perubahan lokal.
  3. ⚠ Toggle autosave per view menyimpan perubahan Creator secara langsung tanpa banner.
  4. Contributors/Viewers dapat mengubah tampilan lokal tanpa menyimpan; perubahan hilang saat halaman ditutup.
  5. ⚠ Konflik edit bersamaan diselesaikan dengan last write wins dan pengguna lain menerima pemberitahuan bahwa view diperbarui.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perubahan lokal disimpan di state klien (boleh `sessionStorage`), bukan server.

### JPD-IDE-050 View Description & Ownership

- **Deskripsi:** Deskripsi view untuk menjelaskan tujuan dan cara membacanya, serta informasi pemilik view.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator mengedit; semua pembaca melihat.
- **Aturan perilaku (acceptance criteria):**
  1. View description rich text tampil di header view dan dapat di-expand/collapse.
  2. Description dapat diedit oleh Creator.
  3. ⚠ Info owner/pembuat dan last updated view ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ownership eksplisit belum terkonfirmasi; clone menyimpan `created_by` dan `updated_at`.

### JPD-IDE-051 View Templates & Pre-built Views

- **Deskripsi:** Konfigurasi view siap pakai untuk use-case umum, termasuk galeri pre-built views (beta) yang dapat ditambahkan ke project.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator).
- **Aturan perilaku (acceptance criteria):**
  1. Template Roadmap Now/Next/Later (Board by Roadmap).
  2. Template Impact vs Effort (Matrix).
  3. Template Prioritization list (List dengan formula, sort desc).
  4. Template Timeline roadmap.
  5. Blank view per tipe.
  6. Galeri pre-built views (beta) untuk kasus umum seperti prioritization dan roadmap.
  7. Pre-built view dapat ditambahkan ke sidebar/section views.
  8. Setelah ditambahkan, view dapat dikustomisasi (filter, field, sort, group) tanpa terikat ke template.
  9. Label status beta ditampilkan di UI galeri.
  10. ⚠ Template membuat field yang dibutuhkan bila belum ada di project.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perbedaan persis pre-built views dengan view bawaan project template belum terkonfirmasi; clone memakai sumber data seed yang sama dengan JPD-IDE-002.

## Berbagi, publikasi & roadmap lintas project

### JPD-IDE-052 Share & Publish Views

- **Alias inventaris:** Publish and share views
- **Deskripsi:** Titik masuk Share dan Publish pada view discovery (tombol Share di header view, penanda view terpublikasi di sidebar, dan pemetaan konfigurasi view ke published view). Aturan berbagi, izin publish, audiens, dan batasan field dimiliki JPD-INS-028 Share Views with Team dan JPD-INS-029 Publish Views ([37](37-jpd-insights-delivery-peran-integrasi.md)); heading ini adalah rujukan sisi konfigurasi view.
- **Prioritas / Fase / Tier:** P1 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs mengaktifkan fitur publish (default nonaktif, JPD-INS-029 butir 1); Administrator proyek dapat publish secara default; Creator hanya dapat publish bila diberi custom role dengan permission "Publish a view" (JPD-INS-003); Contributor tidak dapat publish (HTTP 403) tetapi dapat membagikan view lewat Share (JPD-INS-028); Stakeholder/Viewer membuka published view secara read-only. Publish memerlukan plan Standard atau lebih tinggi (JPD-INS-029 butir 2).
- **Aturan perilaku (acceptance criteria):**
  1. Copy link view tersedia untuk pengguna dengan akses project; pengguna tanpa akses melihat halaman "Minta akses" (JPD-INS-028 butir 4).
  2. Publish view menghasilkan versi read-only yang selalu live mengikuti view sumber tanpa publish ulang (JPD-INS-029 butir 11).
  3. ⚠ Audiens published view: pengguna organisasi yang memiliki akun produk ini secara default; publish ke siapa pun dengan tautan hanya bila Admin Situs mengaktifkannya dan pengguna memiliki permission "Publish a view to anyone with the link" (JPD-INS-029 butir 7-8).
  4. Unpublish (Share > Publish > Unpublish) langsung menonaktifkan tautan; pembuka tautan lama melihat "View ini tidak lagi dipublikasikan" (JPD-INS-029 butir 9).
  5. Viewer membuka idea dalam mode read-only dari published view.
  6. Published view dapat disematkan ke alat dokumentasi eksternal melalui Smart Link/oEmbed generik (EDT-021, [07](07-editor-konten-kaya.md)); tidak ada integrasi native dengan produk Confluence.
  7. ⚠ Published view tidak menampilkan field/idea dari project yang tidak dapat diakses viewer.
  8. Published view dapat diakses tanpa lisensi Creator (Stakeholder tanpa akses proyek, JPD-INS-029).
  9. Konfigurasi view yang memakai field terlarang untuk published view (Goals, Type, Votes, Reactions, Custom formula, Number of linked work items, Number of insights) di sort/filter/group/kolom/sumbu matrix membuat tombol Publish nonaktif dengan penjelasan field penyebab (JPD-INS-029 butir 10).
  10. Dialog Publish menampilkan toggle field yang terlihat di published view, opsi menampilkan description idea, dan tombol Preview sebelum publish (JPD-INS-029 butir 5).
  11. View yang di-publish ditandai ikon globe di daftar view sidebar (JPD-IDE-040).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik aturan: JPD-INS-028 Share Views with Team dan JPD-INS-029 Publish Views ([37](37-jpd-insights-delivery-peran-integrasi.md)); modul 36 memiliki konfigurasi view, modul 37 memiliki izin, audiens, dan kolaborasi; bila ada perbedaan aturan, aturan modul 37 yang berlaku. Sub-fitur lampiran inventaris ("Publish and share views": link read-only, unpublish, copy link, data live, akses tanpa lisensi Creator, pembatasan field, embed) tercakup di butir 1, 2, 4, 6, 8, dan 10. Akses anonim/publik mengikuti PERM-046 ([03](03-izin-peran-keamanan.md)). Entitas `published_views` (view_id, token acak, audience, allowed_fields, show_description, revoked_at).

### JPD-IDE-053 Stakeholder Updates

- **Deskripsi:** Mengirim ringkasan update roadmap/view ke daftar stakeholder via email.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator mengirim; Stakeholder menerima.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Daftar stakeholders (pengguna atau alamat email) dikelola per project atau per view.
  2. Creator menyusun update berisi pesan dan view terlampir.
  3. ⚠ Riwayat update yang terkirim tersimpan dan dapat dilihat.
  4. Stakeholder membuka view terpublikasi dari tautan di email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rilis 2025 di produk acuan; ketersediaan dan plan belum terkonfirmasi. Pengiriman memakai infrastruktur email NTF ([19](19-notifikasi-email.md)). Peran stakeholder rinci di modul 37.

### JPD-IDE-054 Export Views

- **Deskripsi:** Mengekspor isi view untuk dibagikan di luar produk.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Pengguna dengan akses project.
- **Aturan perilaku (acceptance criteria):**
  1. View List dapat diekspor ke CSV dengan field yang terlihat sesuai view dan urutan view.
  2. ⚠ View dapat diekspor sebagai gambar/PNG.
  3. Hanya pengguna dengan akses project yang dapat mengekspor; published view tidak menyediakan ekspor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai engine ekspor CSV SRC-050. Format ekspor di produk acuan belum terkonfirmasi.

### JPD-IDE-055 Cross-project Roadmaps

- **Alias inventaris:** Roadmaps (cross-project views)
- **Deskripsi:** View yang menggabungkan ideas dari beberapa project discovery dalam satu roadmap/list untuk pemimpin produk dan stakeholder.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Product Owner/PM dan Creator membuat; Admin Situs; Stakeholder membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna memilih project sumber.
  2. Hanya global fields yang dapat dipakai untuk filter/group/sort lintas project.
  3. ⚠ Tipe view yang didukung adalah List, Board, dan Timeline.
  4. ⚠ Hierarki idea types ditampilkan lintas project bila tipe yang sama dipakai.
  5. Akses mengikuti permission tiap project sumber; idea dari project yang tidak dapat diakses tidak ditampilkan.
  6. Roadmap dapat di-publish/share ke stakeholder (JPD-IDE-052).
  7. Beberapa project discovery dapat dipilih sebagai sumber.
  8. Global fields menjadi dasar filter/group yang konsisten.
  9. ⚠ Tipe view roadmap mengikuti daftar yang didukung.
  10. Permission mengikuti akses pengguna ke tiap project.
  11. Roadmap dikelola di area Roadmaps di luar sidebar project.
  12. Roadmap dapat dipublikasikan ke stakeholder.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kemungkinan plan Premium. Bergantung pada JPD-IDE-027. Berbeda dari Plans lintas project untuk work item (modul 18).

## Pencarian, sandbox & detail UX

### JPD-IDE-056 JPD Search & JQL Support

- **Alias inventaris:** JQL and search for ideas; JPD Field Search
- **Deskripsi:** Pencarian ideas lewat pencarian global, JQL, dan pencarian di dalam fitur discovery.
- **Prioritas / Fase / Tier:** P2 / Fase 6 / jpd
- **Aktor & izin:** Creator, Viewer, Admin Situs, Integrator/Developer API; hasil selalu dibatasi izin (JQL-051).
- **Aturan perilaku (acceptance criteria):**
  1. Ideas muncul di quick search dan pencarian global (SRC-009).
  2. JQL `project = KEY` dan `issuetype = Idea` mengembalikan ideas.
  3. ⚠ Field discovery dapat di-query via JQL dengan cakupan terbatas; formula tidak dapat di-query.
  4. Pencarian ideas tersedia saat link dan merge (summary/key).
  5. Saved filters untuk ideas didukung sebagai titik integrasi, termasuk entry point bulk change/move (SRC-031, SRC-036).
  6. `project = X AND issuetype = Idea` berfungsi di navigator.
  7. Sebagian field discovery searchable sebagai custom field (JQL-033).
  8. Ideas muncul di global search.
  9. Views discovery memakai filter sendiri, bukan JQL.
  10. ⚠ Field computed (insights, delivery progress) tidak searchable lewat JQL.
  11. Idea dicari sebagai work type "Idea" dalam project discovery.
  12. ⚠ Field Impact, Effort, Goals, dan Insights tersedia di JQL hanya bila field tersebut bertipe yang didukung searcher.
  13. Link idea ke delivery ticket dicari via `issueLinkType` / `linkedIssues()`.
  14. View discovery memakai filter sendiri, bukan JQL penuh.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JQL di [12](12-jql-bahasa-kueri.md); searcher custom field CFG-038. Tumpang tindih dengan JPD-INS-038 JPD JQL Support ([37](37-jpd-insights-delivery-peran-integrasi.md)) yang menambah `project.type = product_discovery` dan ekspor hasil pencarian (JPD-INS-039); keduanya satu kemampuan dengan prioritas sama (P2) dan aturan di kedua modul berlaku bersama. Cakupan JQL untuk field khusus discovery di produk acuan belum terkonfirmasi.

### JPD-IDE-057 JPD Sandbox Support

- **Deskripsi:** Dukungan sandbox untuk Product Discovery agar konfigurasi (fields, views, workflow) dapat diuji di lingkungan terpisah sebelum diterapkan di produksi.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Admin Situs (dan admin organisasi).
- **Aturan perilaku (acceptance criteria):**
  1. Sandbox dibuat dari administrasi organisasi (ADM-049) dan dapat menyertakan produk Product Discovery.
  2. ⚠ Data/konfigurasi produksi Product Discovery dapat disalin ke sandbox (projects, fields, views, ideas).
  3. Akses sandbox terpisah dari produksi; perubahan di sandbox tidak memengaruhi produksi.
  4. Sandbox dapat dihapus atau di-reset.
  5. ⚠ Hanya plan Premium/Enterprise; jumlah sandbox per produk mengikuti plan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sandbox adalah fitur platform admin (ADM-049).

### JPD-IDE-058 JPD Keyboard Shortcuts

- **Deskripsi:** Pintasan keyboard untuk navigasi dan aksi cepat di fitur discovery.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery (Creator) dan semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Esc menutup panel/dialog.
  2. ⚠ Navigasi next/previous idea di panel memakai tombol J/K.
  3. Enter men-submit inline create.
  4. Shortcut global (NAV-038, NAV-039) tetap berlaku dan dapat dimatikan pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Set shortcut khusus discovery di produk acuan belum terkonfirmasi.

### JPD-IDE-059 UI Specifics (Dark Mode, Responsiveness, Empty States)

- **Deskripsi:** Aspek UI discovery: tema, responsivitas, dan state khusus.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Tema light/dark/system mengikuti pengaturan pengguna (NAV-034).
  2. Empty states untuk view kosong dan filter tanpa hasil menampilkan ajakan (create idea atau clear filter).
  3. Loading skeleton tampil untuk views besar.
  4. ⚠ Dukungan mobile terbatas: views dapat dibaca dan idea dibuka, tetapi Matrix/Timeline drag tidak didukung di layar kecil.
  5. Tooltips dan onboarding hints tampil untuk pengguna baru.
  6. Aksesibilitas keyboard dan ARIA memenuhi NFR-022.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dark mode adalah fitur platform (NAV-034); empty state umum NAV-045.

## Pertanyaan terbuka

1. Apakah clone menerapkan pembagian plan Product Discovery (Free/Standard/Premium/Enterprise) untuk idea types & hierarki, Tree view, roadmap lintas project, public link, dan sandbox, atau semua tersedia di instance self-hosted tanpa pembatasan?
2. Apakah label UI memakai "project" atau "space" untuk discovery, dan apakah dapat dikonfigurasi per site?
3. Berapa nilai default vote budget saat field Votes dibuat (vote budget sendiri sudah ditetapkan di JPD-INS-026 dengan nilai minimal 1)?
4. Bagaimana perilaku perubahan konfigurasi view yang dipilih sebagai default: autosave atau banner Save/Discard?
5. Apakah opsi publish ke siapa pun dengan tautan (tanpa login) diaktifkan pada instance self-hosted, mengingat JPD-INS-029 butir 7 menjadikannya opsional dan dikendalikan Admin Situs?
6. Berapa batas jumlah field custom, view, idea type, dan level hierarki per project?
7. Apakah formula boleh direferensikan oleh formula lain (bersarang) selama tidak siklik, dan bagaimana pembulatan hasil ditampilkan?
8. Apakah Merge dan Move lintas project discovery masuk cakupan Fase 6 atau ditunda?
9. Apakah field discovery (Impact, Effort, Goals, formula) dibuat searchable penuh di JQL, mengingat produk acuan membatasinya?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk roadmap/prioritisasi (mis. app roadmap, scoring, voting portal pihak ketiga) tidak dibangun; kebutuhan serupa dipenuhi oleh fields, formula, dan views native modul ini.
- **Layanan lintas produk platform asli** (Atlassian Home/Goals/Projects, Atlassian Teams directory, admin.atlassian.com sandbox) tidak direplikasi apa adanya; clone memakai padanan internal (modul 18, 22, 25) dan modul ini hanya mendefinisikan titik integrasinya.
- **Confluence** sebagai produk tidak dibangun; embed published view hanya berupa Smart Link/oEmbed generik (EDT-021) ke halaman eksternal yang kompatibel.
- **Insights, delivery tickets, peran Creator/Contributor/Stakeholder, integrasi (Slack, browser extension, Teams), dan API discovery** dirinci di modul 37, bukan di modul ini.
- **AI** (ringkasan idea, pembuatan idea otomatis, saran insight berbasis Rovo/Atlassian Intelligence) di luar cakupan modul ini.
- **Fitur Data Center-only atau deprecated**: tidak ada pada jatah modul ini (Product Discovery hanya tersedia di Cloud).
- **Dependencies antar idea di Timeline, logika OR antar filter view, dan sections bernama di idea view** tidak dibangun pada fase awal karena belum terkonfirmasi ada di produk acuan.
