# 13 — Pencarian, Navigator, Filter Tersimpan & Ekspor

Modul ini mendefinisikan seluruh pengalaman menemukan work item di produk ini: mesin pencarian dan strategi indeks (PostgreSQL FTS sebagai default, OpenSearch sebagai opsi skala besar), perilaku pencarian teks, quick search global dan lompatan langsung lewat key, riwayat item yang baru dilihat, halaman advanced search (navigator) dengan mode Basic dan JQL, tampilan hasil (List dan Detail, kolom, sort, pagination, hierarki, inline edit/create, bulk change), navigator dalam konteks project (filter pill, grouping, work item terarsip), siklus hidup filter tersimpan (simpan, ubah, star, system filter, direktori, hapus, pemakaian sebagai sumber data), berbagi dan langganan filter, serta ekspor hasil (CSV, XML, Word, cetak, XLSX form data, spreadsheet) beserta batasan keamanan dan performanya. Grammar, operator, dan fungsi JQL didefinisikan di modul 12; modul ini hanya mengatur cara JQL dipakai di UI. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, middleware bahasa EN/ID), sehingga indeks pencarian, tabel `filters` dan turunannya (sharing, favourite, subscription, kolom), riwayat item yang dilihat, job ekspor, dan halaman navigator adalah pekerjaan baru. Pekerjaan itu dikemas sebagai feature module baru (mis. `internal/features/search`, `internal/features/filter`, `internal/features/export`) mengikuti pola Clean Architecture repo, dengan route SPA baru `/issues` dan `/filters` di `frontend/`.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi "work item", plan tier dan feature flag.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — global permission (USR-022), akses anonim (USR-024), user picker API (USR-025), data security policies (USR-040), audit log (USR-046).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects (PERM-020), Share Dashboards and Filters (PERM-009), Manage Group Filter Subscriptions (PERM-010), Make Bulk Changes (PERM-011), Issue Security Enforcement (PERM-044), Anonymous & Public Access (PERM-046), Board & Filter Permissions (PERM-051), Data Security Policies (PERM-065).
- [04 Proyek](04-proyek.md) — Change Project Key (PRJ-022), navigasi project (PRJ-027), arsip project (PRJ-032).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — Parent field (CFG-013), custom field contexts (CFG-035), searcher dan indexing custom field (CFG-038).
- [06 Work Item Inti](06-work-item-inti.md) — key work item (ISS-001), inline create (ISS-005), issue view (ISS-007, ISS-008), navigasi antar work item (ISS-013), Move (ISS-032), Archive (ISS-034), Bulk Change (ISS-035), Share work item (ISS-038).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar yang ikut diindeks, history.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — transisi dari dropdown status (WF-011), bulk transition (WF-013).
- [10 Board](10-board-kolom-kartu.md) — board dari saved filter (BRD-003), board filter (BRD-010).
- [11 Backlog & Sprint](11-backlog-sprint.md) — multi-select backlog (SPR-004).
- [12 JQL](12-jql-bahasa-kueri.md) — grammar, operator, fungsi (`currentUser()`, `issueHistory()`), klausa `filter =`, `ORDER BY`.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — kolom Σ time tracking.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — "View in search" dari version dan component.
- [16 Laporan & Insight](16-laporan-insight.md) — drill-down laporan ke navigator.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget berbasis filter.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — tab List project, Plans dengan filter sebagai issue source.
- [19 Notifikasi & Email](19-notifikasi-email.md) — email share dan email subscription.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — top nav/sidebar, halaman For you, katalog keyboard shortcut, preferensi user.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — General configuration (public sharing), Default user preferences, Shared items.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — `POST /api/v1/search/jql`, Filters API, rate limit.
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — CSV importer, titik ekstensi app.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — Forms (sumber ekspor XLSX).
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — target latensi, outbox, background job, rate limiting.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Filter, FilterShare, Subscription dan event domain `issue.*`.

## Persona terkait

- **Developer/Anggota Tim** — memakai quick search dan key untuk berpindah cepat, mencari dengan Basic/JQL, mengedit inline, men-star filter tim.
- **Scrum Master/Team Lead** — menyusun filter untuk board dan rapat triase, membagikan filter ke project/role, berlangganan email filter.
- **Product Owner/PM** — membuat filter backlog lintas project, mengelompokkan dan mengekspor hasil untuk perencanaan.
- **Stakeholder/Viewer** — membuka permalink dan filter yang dibagikan, menerima email subscription, mengekspor CSV untuk laporan.
- **Pengguna Tim Bisnis** — memakai halaman List project dengan filter pill dan grouping tanpa JQL.
- **Admin Proyek/Project Lead** — memakai bulk change, melihat work item terarsip, menjaga filter yang dipakai board project.
- **Admin Situs** — mengatur global permission berbagi/subscription/bulk, public sharing, default sharing, mengelola filter milik pengguna lain dan filter yatim.
- **Integrator/Developer API** — mengambil hasil lewat XML/CSV berbasis URL dan REST API, memakai filter sebagai sumber data.
- **Tamu/Kolaborator Eksternal** — mencari di project yang dibuka publik/anonim sesuai izin.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin menekan `/` lalu mengetik kata atau key agar saya langsung menemukan dan membuka work item tanpa meninggalkan halaman (SRC-009, SRC-010).
- Sebagai Developer/Anggota Tim, saya ingin melihat item yang baru saya buka di dropdown search agar saya bisa kembali ke pekerjaan sebelumnya dengan satu klik (SRC-011).
- Sebagai Pengguna Tim Bisnis, saya ingin mencari dengan dropdown Project, Type, Status, dan Assignee tanpa menulis JQL agar saya tetap bisa menyaring pekerjaan sendiri (SRC-016, SRC-033).
- Sebagai Scrum Master/Team Lead, saya ingin menulis JQL dengan autocomplete dan pesan galat yang jelas agar query kompleks dapat saya susun tanpa coba-coba (SRC-017, SRC-018).
- Sebagai Product Owner/PM, saya ingin mengatur kolom, sort, dan grouping hasil lalu menyimpannya sebagai filter agar tampilan triase yang sama dapat dipakai ulang tim (SRC-025, SRC-026, SRC-034, SRC-036).
- Sebagai Developer/Anggota Tim, saya ingin mengubah assignee, prioritas, dan status langsung di tabel hasil agar triase puluhan work item tidak perlu membuka satu per satu (SRC-029, SRC-024).
- Sebagai Admin Proyek/Project Lead, saya ingin memilih banyak hasil dan menjalankan bulk edit/transition/move agar pembersihan data berlangsung sekali jalan (SRC-031).
- Sebagai Scrum Master/Team Lead, saya ingin membagikan filter ke project role dan memberi hak edit ke rekan agar filter board dapat dirawat bersama (SRC-045, SRC-046, SRC-044).
- Sebagai Stakeholder/Viewer, saya ingin menerima hasil filter lewat email setiap Senin pagi agar saya memantau progres tanpa login (SRC-047).
- Sebagai Stakeholder/Viewer, saya ingin mengekspor hasil ke CSV dengan kolom filter agar data bisa diolah di spreadsheet (SRC-050).
- Sebagai Admin Situs, saya ingin mencari semua filter di site, termasuk filter publik dan milik pengguna nonaktif, lalu mengganti owner-nya agar board dan dashboard tidak rusak (SRC-049, SRC-043).
- Sebagai Admin Situs, saya ingin mengendalikan siapa yang boleh berbagi filter, berlangganan untuk group, dan bulk change agar data tidak bocor dan operasi massal terkendali (SRC-005, SRC-058).
- Sebagai Integrator/Developer API, saya ingin hasil pencarian selalu memfilter permission di dalam query dan menampilkan perubahan saya sendiri seketika agar integrasi tidak membocorkan data dan tidak melihat data basi (SRC-001, SRC-006).
- Sebagai Tamu/Kolaborator Eksternal, saya ingin mencari work item di project publik tanpa akun agar saya bisa menindaklanjuti laporan yang terbuka (SRC-009, SRC-015, SRC-005).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 17 |
| P1 | 20 |
| P2 | 11 |
| P3 | 10 |
| **Total** | **58** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 12 |
| Fase 2 | 23 |
| Fase 3 | 11 |
| Fase 4 | 8 |
| Fase — (dc-only/deprecated, rujukan) | 4 |

Strategi urutan implementasi: Fase 1 membangun fondasi yang tidak bisa diganti belakangan, yaitu abstraksi `SearchRepository` di atas PostgreSQL FTS dengan permission filter di dalam query dan read-your-writes (SRC-001, SRC-004), lalu di atasnya quick search, lompatan key, riwayat item, halaman navigator dengan mode Basic, tabel List dengan sort dan pagination, halaman Work items project, dan system filter. Mode Basic di Fase 1 sudah mengompilasi kriteria ke AST query yang sama dengan yang nanti dihasilkan parser JQL (modul 12), sehingga perpindahan ke JQL di Fase 2 tidak mengubah lapisan eksekusi. Fase 2 menambah JQL editor dan toggle Basic/JQL, filter tersimpan beserta star, direktori, sharing, permalink, dan pemakaiannya oleh board/gadget, kolom yang dapat dikonfigurasi, Detail view, inline edit, filter pill project, serta batasan performa dan pesan kondisi khusus, ditambah filter subscription email (SRC-047) yang dijadwalkan bersama kanal pengirimannya di modul 19 (NTF-030). Fase 3 menambah operasi skala organisasi: bulk change, hak edit filter, administrasi filter oleh admin, ekspor CSV/XML asinkron, grouping, hierarki, inline create, pencarian work item terarsip, dan shortcut navigator. Fase 4 berisi fitur opsional dan ekosistem (natural language ke JQL, pencarian AI lintas modul, Word/cetak/XLSX form, konektor spreadsheet, titik ekstensi app, kebijakan keamanan data), sedangkan fitur Data Center-only hanya dicatat sebagai rujukan.

## Fitur

## Fondasi mesin pencarian & keamanan hasil

### SRC-001 Search Engine Architecture (Postgres FTS vs OpenSearch)

- **Deskripsi:** Keputusan arsitektur mesin pencarian yang melayani JQL dan pencarian teks: PostgreSQL FTS sebagai mode default yang mengikuti stack repo, dan OpenSearch/Elasticsearch sebagai mode opsional untuk instance besar, di balik satu abstraksi repository.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering dan Architect yang membangun dan mengoperasikan; tidak ada permission pengguna akhir, tetapi setiap query yang dihasilkan wajib menegakkan Browse Projects (PERM-020) dan issue security (PERM-044).
- **Aturan perilaku (acceptance criteria):**
  1. Mode default memakai kolom `tsvector` generated pada work item yang menggabungkan summary, description, environment, dan teks komentar (komentar diagregasi lewat tabel/indeks pendamping), ditambah ekstensi `pg_trgm` untuk pencarian parsial pada summary; pencarian `summary ~ "logi"` mengembalikan work item dengan summary "Login gagal".
  2. Konfigurasi FTS memakai dua konfigurasi: `simple` (tanpa stemming, untuk key, kode, nama) dan `english` (stemming); stemming Bahasa Indonesia dipakai melalui dictionary kustom bila tersedia di server, dan bila tidak tersedia sistem jatuh ke `simple` tanpa error.
  3. Mode opsional OpenSearch/Elasticsearch menyimpan dokumen work item terdenormalisasi (field sistem, custom field, nama project, status category, ACL) dan direkomendasikan untuk instance di atas 1 juta work item; mode dipilih lewat konfigurasi environment, bukan UI.
  4. Use case pencarian hanya bergantung pada interface `SearchRepository`; JQL compiler (modul 12) menghasilkan AST yang diterjemahkan oleh implementasi SQL atau OpenSearch DSL, dan kedua implementasi lulus test suite kontrak yang sama untuk hasil dan urutan.
  5. Pada mode Postgres, indeks pencarian diperbarui sinkron di dalam transaksi yang sama dengan perubahan work item; pada mode OpenSearch, perubahan ditulis ke tabel outbox lalu diproses indexer asinkron dengan target lag di bawah 2 detik pada p95.
  6. Reindex penuh dan per project berjalan sebagai background job dengan progress persen yang dapat dipantau admin, tanpa downtime: pencarian tetap melayani indeks lama sampai indeks baru siap lalu dialihkan secara atomik.
  7. ⚠ Operator teks `~` mendukung kata, frasa `"..."`, dan wildcard akhiran `*`; sintaks Lucene lain (fuzzy `~`, proximity `"a b"~5`, boost `^`) ditolak dengan pesan galat yang menyebut sintaks tidak didukung, mengikuti cakupan produk referensi Cloud.
  8. Pencocokan teks membuang stop words, tidak membedakan huruf besar-kecil, dan menormalkan aksen (`unaccent`) sehingga "cafe" cocok dengan "café".
  9. Filter permission (Browse Projects, issue security level, dan pembatasan anonim) disusun sebagai predikat di dalam query yang sama, bukan post-filter di aplikasi, sehingga total hasil dan pagination tidak pernah menghitung work item yang tidak boleh dilihat.
  10. Read-your-writes: work item yang baru dibuat atau diubah oleh pengguna langsung muncul pada pencarian berikutnya oleh pengguna itu, termasuk pada mode OpenSearch (mis. dengan menggabungkan hasil indeks dengan perubahan milik pengguna yang belum terindeks, setara parameter `reconcileIssues` produk referensi).
- **Team-managed vs company-managed:** Tidak ada perbedaan pada mesin; perbedaan semantik field dibahas di SRC-008.
- **Catatan & ketergantungan:** Target latensi dan kapasitas di modul 27; pola outbox dan background job di modul 27/28. Keputusan terbuka: kapan mode OpenSearch dibangun (diusulkan Fase 4 bila ada kebutuhan pelanggan besar).

### SRC-002 Indexing Strategy

- **Deskripsi:** Rancangan index database PostgreSQL yang mendukung pola akses utama navigator, board, dan backlog tanpa membebani operasi tulis.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering dan DBA; tidak ada interaksi pengguna akhir.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel work item memiliki composite index `(project_id, status_id)`, `(assignee_id, status_category)`, `(project_id, updated DESC)`, `(project_id, created DESC)`, dan index pada kolom `rank`; query navigator default per project memakai index scan (diverifikasi dengan `EXPLAIN` di test integrasi).
  2. Tabel relasi `issue_sprints` memiliki index `(sprint_id, issue_id)` dan `(issue_id)` untuk query sprint aktif dan backlog.
  3. Partial index `WHERE resolution_id IS NULL` tersedia untuk kolom yang sering dikombinasikan dengan `resolution = Unresolved`, karena mayoritas filter (termasuk system filter) memakai klausa itu.
  4. GIN index dipakai untuk kolom array labels, kolom JSONB nilai custom field, dan kolom `tsvector`; index trigram GIN dipakai untuk `summary` agar pencarian `contains` tidak melakukan sequential scan.
  5. BRIN index berbasis waktu dipakai untuk tabel append-only besar (audit log, changelog/history).
  6. Setiap foreign key memiliki index pendamping agar cascade delete dan join tidak melakukan sequential scan; pemeriksaan otomatis di CI menggagalkan migrasi yang menambah FK tanpa index.
  7. Migrasi produksi yang menambah index memakai `CREATE INDEX CONCURRENTLY` (di file migrasi terpisah tanpa transaksi) sehingga tabel tidak terkunci untuk tulis.
  8. Ekstensi `pg_stat_statements` diaktifkan; laporan query paling lambat dan index yang tidak terpakai ditinjau secara rutin (diusulkan per rilis) dan index tak terpakai dihapus lewat migrasi.
  9. Covering index (`INCLUDE`) dipakai untuk kolom populer list view navigator (key, summary, status, assignee, priority, updated) sehingga halaman pertama hasil dapat dilayani index-only scan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Migrasi memakai golang-migrate (`make migration-create`). Index dasar FK sudah dibuat sejak Fase 1 bersama tabel terkait; fitur ini meresmikan seluruh strategi di Fase 2 seiring JQL. Detail skema di modul 28.

### SRC-003 Search index administration (Indexing)

- **Deskripsi:** Halaman admin untuk reindex manual di produk referensi Data Center. Produk ini tidak menyediakan halaman setara karena indeks dikelola otomatis oleh sistem; fitur dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (di produk referensi); di produk ini hanya operator platform.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak ada halaman System > Indexing dengan pilihan full/background/project reindex untuk admin; reindex hanya dipicu otomatis oleh sistem atau oleh operator lewat job internal (SRC-001 butir 6).
  2. Perubahan konfigurasi field atau scheme tidak pernah memunculkan pesan "reindex required"; sistem menjadwalkan reindex parsial sendiri dan hasil pencarian tetap konsisten selama proses berjalan.
  3. Index recovery dan replikasi indeks antar node tidak dibangun; pada mode Postgres indeks ikut replikasi database, pada mode OpenSearch memakai replikasi bawaan cluster.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur Data Center-only; lihat bagian Di luar cakupan. Job reindex internal di SRC-001.

### SRC-004 Full-text search behavior

- **Deskripsi:** Aturan pencarian teks ("Contains text", `text ~`, `summary ~`) yang dipakai seragam oleh quick search, Basic search, dan JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna yang mencari; hasil tunduk pada Browse Projects dan issue security.
- **Aturan perilaku (acceptance criteria):**
  1. Field virtual `text` mencakup summary, description, environment, semua komentar yang boleh dilihat pengguna, dan semua custom field bertipe teks yang searchable; komentar ber-visibility terbatas (role/group) hanya mencocokkan bagi pengguna yang boleh melihatnya.
  2. Kata dicocokkan setelah stemming dan penghapusan stop words sesuai konfigurasi bahasa indeks (SRC-001 butir 2), sehingga "running" cocok dengan "run" pada konfigurasi english.
  3. Teks dalam tanda kutip ganda dicari sebagai frasa persis berurutan; di JQL ditulis `text ~ "\"login gagal\""`.
  4. Wildcard `*` (banyak karakter) dan `?` (satu karakter) didukung di tengah atau akhir kata; wildcard di awal kata ditolak dengan pesan "Wildcard tidak boleh berada di awal kata".
  5. ⚠ Fuzzy search (`kata~`) dan proximity search (`"a b"~5`) tidak didukung; query yang memakainya mengembalikan galat validasi yang menjelaskan sintaks tidak didukung.
  6. Operator boolean di dalam string (`AND`, `OR`, `NOT`, `+wajib`, `-kecuali`) dievaluasi; tanpa operator, beberapa kata digabung dengan AND.
  7. Pencocokan selalu case-insensitive.
  8. Karakter reserved (`+ - & | ! ( ) { } [ ] ^ ~ * ? \ :`) yang ingin dicari literal harus di-escape dengan `\\`; karakter reserved tanpa escape pada posisi tidak valid menghasilkan pesan galat yang menunjuk karakter tersebut.
  9. ⚠ Bila query memuat kriteria teks dan tidak memiliki `ORDER BY`, hasil diurutkan berdasarkan skor relevansi (`ts_rank`) menurun, lalu `key` menurun sebagai tie-breaker.
  10. ⚠ Isi file attachment (PDF, dokumen Office, teks) tidak diindeks dan tidak pernah dicocokkan oleh `text ~`; nama file attachment juga tidak termasuk field `text`.
- **Team-managed vs company-managed:** Custom field teks team-managed ikut tercakup `text` hanya untuk work item dari project pemiliknya.
- **Catatan & ketergantungan:** Grammar operator `~` di modul 12; implementasi di SRC-001.

### SRC-005 Global permissions and settings for search & sharing

- **Deskripsi:** Kumpulan global permission dan setting site yang menentukan apa yang boleh dicari, dibagikan, dilanggankan, dan diubah massal dari navigator.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) mengatur; semua pengguna terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. Global permission "Share dashboards and filters" (PERM-009) menjadi syarat untuk menyimpan entri sharing selain Private dan entri Editors; tanpa permission ini dialog permission hanya menampilkan Private dan API menolak dengan 403.
  2. Global permission "Manage group filter subscriptions" (PERM-010) menjadi syarat memilih Recipient = Group pada subscription; tanpa permission ini opsi Group disembunyikan.
  3. Global permission "Make bulk changes" (PERM-011) menjadi syarat entri Bulk change di navigator; tanpa permission ini menu Bulk change disembunyikan dan checkbox seleksi tetap dapat dipakai untuk aksi non-bulk.
  4. Global permission "Browse users and groups" (PERM-008) menentukan apakah user picker di Basic search dan dialog sharing dapat mencari pengguna/group; tanpa permission ini picker hanya menerima email/nama persis dan tidak menampilkan saran.
  5. Isi hasil pencarian selalu dibatasi project permission Browse Projects (PERM-020) dan issue security level (PERM-044) pengguna yang menjalankan query.
  6. Pengguna anonim hanya dapat mencari work item di project yang memberikan Browse Projects kepada Public/anyone (PERM-046); bila tidak ada project publik, halaman pencarian mengarahkan anonim ke login.
  7. Setting "Public sharing of filters and dashboards" di General configuration (modul 22) bernilai aktif/nonaktif, default nonaktif; saat nonaktif, opsi Public tidak tampil di dialog sharing dan filter yang sudah publik tetap tersimpan tetapi tidak dapat dibuka anonim.
  8. Default sharing untuk filter dan dashboard baru diatur di Default user preferences (modul 22) dengan nilai default Private; pengguna dapat menimpa default ini di preferensi personal (modul 21).
- **Team-managed vs company-managed:** Team-managed memakai role project (Administrator/Member/Viewer) untuk Browse; global permission berlaku sama.
- **Catatan & ketergantungan:** Definisi permission ada di modul 03 dan 02 (USR-022); fitur ini hanya mengatur dampaknya pada pencarian dan berbagi.

### SRC-006 Search limits and performance safeguards

- **Deskripsi:** Batasan sistem yang melindungi stabilitas pada pencarian, ukuran hasil, operasi massal, dan ekspor, beserta pesan yang jelas saat batas tercapai.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Semua pengguna dan Admin Situs; batas tidak dapat dilewati oleh permission apa pun kecuali konfigurasi operator.
- **Aturan perilaku (acceptance criteria):**
  1. Query yang melewati batas waktu eksekusi (diusulkan 10 detik di UI, dapat dikonfigurasi operator) dibatalkan di database (`statement_timeout`) dan pengguna menerima pesan "Pencarian terlalu lambat atau terlalu kompleks. Persempit kriteria lalu coba lagi."
  2. Bulk change dibatasi 1.000 work item per operasi; ekspor CSV asinkron dibatasi 10.000 work item; ekspor format lain (XML, Word, cetak) dibatasi 1.000 work item; permintaan di atas batas ditolak sebelum diproses dengan pesan yang menyebut batasnya.
  3. Fungsi JQL berbasis riwayat atau daftar besar memiliki batas hasil (mis. `issueHistory()` maksimal 50 work item, SRC-011); batas setiap fungsi didokumentasikan di modul 12.
  4. ⚠ Panjang JQL dibatasi (diusulkan 10.000 karakter) dan URL navigator dibatasi (diusulkan 8.000 karakter); bila JQL terlalu panjang untuk URL, UI menyarankan menyimpannya sebagai filter dan memakai `?filter=ID`.
  5. Pada mode indeks asinkron (OpenSearch), perubahan pengguna lain tampil di hasil setelah jeda singkat (target di bawah 2 detik); perubahan milik pengguna sendiri tampil seketika (SRC-001 butir 10).
  6. Pencarian dari UI dan API tunduk pada rate limiting per pengguna dan per token (modul 23/27); permintaan yang melebihi batas mendapat HTTP 429 dengan header `Retry-After`.
  7. Bila hasil atau ekspor dipotong karena batas, UI menampilkan pesan eksplisit, mis. "Menampilkan 10.000 dari 23.512 work item. Persempit query untuk mengekspor sisanya."
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Angka final NFR di modul 27; rate limit API di modul 23.

### SRC-007 Search result empty, error, and permission states

- **Deskripsi:** Penanganan seragam untuk kondisi tanpa hasil, filter hilang, akses ditolak, referensi kedaluwarsa, dan kegagalan jaringan pada hasil pencarian dan filter.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna termasuk anonim.
- **Aturan perilaku (acceptance criteria):**
  1. Query valid tanpa hasil menampilkan ilustrasi dan pesan "Tidak ada work item yang cocok dengan pencarian Anda" disertai saran menghapus kriteria, memeriksa ejaan, atau memperluas project; tombol "Clear filters" tampil bila ada kriteria aktif.
  2. Membuka `?filter=ID` yang tidak ada atau sudah dihapus menampilkan pesan "Filter tidak ditemukan atau telah dihapus" dengan tautan ke direktori filter (SRC-041).
  3. Membuka filter yang tidak dibagikan kepada pengguna menampilkan pesan "Anda tidak memiliki izin untuk melihat filter ini" tanpa menampilkan nama, owner, atau JQL filter.
  4. JQL yang merujuk project, field, atau nilai yang tidak dapat diakses pengguna menghasilkan galat yang sama dengan entitas yang tidak ada (mis. "Nilai 'X' tidak ada untuk field 'project'"), sehingga keberadaan entitas tersembunyi tidak bocor.
  5. Filter lama yang merujuk field atau nilai yang sudah dihapus menampilkan galat JQL yang menyebut referensi tersebut beserta tombol "Edit filter" untuk memperbaikinya; board/gadget pemakai filter menampilkan galat serupa.
  6. Pengguna anonim yang membuka navigator atau filter yang memerlukan login diarahkan ke halaman login dengan parameter kembali ke URL semula.
  7. Kegagalan jaringan atau server (5xx/timeout klien) menampilkan pesan "Gagal memuat hasil" dengan tombol Retry, tanpa menghapus hasil terakhir yang berhasil dimuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pesan galat tersedia EN/ID melalui `internal/shared/constants/error.go` dan `LanguageMiddleware`. Empty state minimal (butir 1) sudah dibutuhkan List view Fase 1 (SRC-023 butir 11).

### SRC-008 Team-managed vs company-managed search differences

- **Deskripsi:** Aturan perbedaan perilaku pencarian, kriteria, dan kolom antara project team-managed (konfigurasi lokal) dan company-managed (konfigurasi berbasis scheme).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang mencari lintas project; Admin Proyek yang mengonfigurasi field lokal.
- **Aturan perilaku (acceptance criteria):**
  1. Field team-managed bersifat lokal per project dan tampil di Basic search, picker kolom, dan autocomplete dengan label nama project, mis. "Story size (Proyek Mobile)".
  2. ⚠ Custom field team-managed dapat dicari via JQL dengan nama field; bila beberapa field berbagi nama, autocomplete menawarkan bentuk yang tidak ambigu (`cf[ID]` atau nama dengan label project) dan query berbasis nama mencakup semua field bernama sama.
  3. Status dan issue type team-managed yang bernama sama dengan entitas di project lain tetap entitas berbeda; kriteria berbasis nama (`status = "Review"`) mencakup semuanya, sedangkan pemilihan di Basic search dapat memilih entitas per project.
  4. Board team-managed tidak berbasis saved filter (langsung pada project), sedangkan board company-managed selalu berbasis saved filter (BRD-003); konsekuensinya menghapus/mengubah filter tidak pernah memengaruhi board team-managed.
  5. Kolom field team-managed di list view hanya terisi untuk work item dari project pemilik field; baris lain kosong.
  6. Hierarki di hasil memakai field Parent terpadu (CFG-013) untuk kedua jenis project; field Epic Link lama tidak tersedia sebagai kriteria baru dan diterjemahkan ke `parent`.
  7. Custom field context company-managed (CFG-035) menentukan field mana yang tersedia di Basic mode berdasarkan project dan issue type yang dipilih (SRC-016 butir 8).
- **Team-managed vs company-managed:** Seluruh fitur ini adalah daftar perbedaannya.
- **Catatan & ketergantungan:** Model field lokal di modul 05; tipe project di modul 04.

## Pencarian cepat global

### SRC-009 Quick Search (global search bar)

- **Deskripsi:** Kotak pencarian di navigasi global yang menampilkan hasil seketika (work item, project, board, filter, plan, dashboard, orang/tim) selagi pengguna mengetik, tanpa perlu JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna login dan pengguna anonim; setiap tipe hasil difilter oleh permission objeknya (Browse Projects, izin board/filter/dashboard).
- **Aturan perilaku (acceptance criteria):**
  1. Menekan `/` dari halaman mana pun (kecuali saat fokus di input/editor teks) memindahkan fokus ke kotak search.
  2. Saat kotak difokuskan dan masih kosong, dropdown menampilkan item yang terakhir dilihat (work item, board, project/space, filter, dashboard, plan) dari penyimpanan recent items NAV-011 (lewat SRC-011), maksimal 10 item.
  3. Hasil muncul selagi mengetik dengan debounce 250 ms dan minimal 1 karakter, dikelompokkan per tipe objek dalam urutan kanonis Work items, Spaces (project), Boards, Filters, Plans, Dashboards, People/Teams, lalu konten modul/aplikasi lain; setiap grup menampilkan maksimal 5 hasil dan grup kosong disembunyikan. Urutan ini juga berlaku untuk dropdown NAV-014.
  4. Setiap hasil work item menampilkan ikon issue type, key, summary, dan nama project; klik membuka issue view.
  5. Menekan Enter tanpa memilih hasil dievaluasi berurutan: (a) bila input cocok persis dengan pola key work item yang ada dan dapat diakses, work item itu langsung dibuka (SRC-010 butir 2); (b) selain itu, bila halaman hasil penuh (SRC-012) aktif, Enter membuka halaman itu dengan query yang sama; (c) bila tidak, Enter membuka advanced search (SRC-015) dengan JQL `text ~ "kata"`. Tautan "View all work items" selalu membuka advanced search dengan JQL tersebut, dan tautan "View all results" membuka SRC-012.
  6. Tab cakupan produk (mis. Work management, Service, Discovery) hanya tampil bila site memiliki lebih dari satu produk/modul aktif; tanpa itu tab tidak dirender.
  7. Panah atas/bawah memindahkan sorotan antar hasil melintasi grup, Enter membuka hasil yang disorot, dan Esc menutup dropdown lalu mengembalikan fokus ke elemen sebelumnya.
  8. Hasil dibatasi Browse Projects dan issue security level; work item, board, atau filter tanpa akses tidak pernah muncul, termasuk dalam jumlah hasil.
  9. ⚠ Saat pengguna berada dalam konteks project, hasil dari project tersebut diberi bobot lebih tinggi (ditampilkan lebih dulu) tanpa menyembunyikan hasil project lain.
  10. Pencocokan teks memakai indeks full-text (summary, description, comment, environment, custom field teks) dengan stemming sesuai SRC-004.
  11. Bila tidak ada hasil, dropdown menampilkan "Tidak ada hasil untuk 'kata'" dan tautan untuk membuka advanced search dengan query tersebut.
  12. Pengguna anonim hanya melihat hasil dari project yang memberi Browse Projects kepada Public; tipe objek lain hanya bila dibagikan publik.
  13. Work item terarsip (SRC-035) dan work item terhapus tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan dengan [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md): SRC-009 adalah pemilik logika hasil (sumber data, urutan grup, perilaku Enter, permission, ranking), sedangkan NAV-014 Quick Search (top bar) adalah pemilik shell UI (posisi search box, shortcut, highlight, filter cepat, recent searches, tampilan "No results"). Bila ada perbedaan aturan, urutan grup dan perilaku Enter di fitur ini yang berlaku; NAV-014 perlu menambahkan grup Dashboards agar selaras. Endpoint usulan `GET /api/v1/search/quick?q=&types=` dengan respons per grup. Pencarian AI lintas modul di SRC-014.

### SRC-010 Search by issue key (jump to work item)

- **Deskripsi:** Mengetik key work item yang valid (mis. `PROJ-123`) langsung membuka work item itu atau menampilkannya sebagai hasil teratas.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna; memerlukan Browse Projects pada project work item dan lolos issue security.
- **Aturan perilaku (acceptance criteria):**
  1. Input yang cocok dengan pola key project (huruf/angka diawali huruf, `-`, angka) dikenali secara case-insensitive, sehingga `proj-123` diperlakukan sebagai `PROJ-123`.
  2. Menekan Enter dengan key yang cocok persis membuka `/browse/PROJ-123` tanpa menampilkan dropdown hasil.
  3. Key lama dari work item yang sudah dipindah ke project lain (ISS-032) tetap me-redirect ke key baru dengan HTTP 301 di level route.
  4. Key lama dari project yang key-nya diganti (PRJ-022) tetap dikenali dan me-redirect ke key baru.
  5. Key yang tidak ada dan key yang tidak boleh diakses menampilkan pesan generik yang identik, "Work item tidak ditemukan atau Anda tidak memiliki akses", sehingga keberadaannya tidak bocor.
  6. ⚠ Mengetik angka saja (mis. `123`) saat berada di konteks project menampilkan `PROJ-123` sebagai hasil teratas bila ada; Enter membukanya.
  7. Di advanced search, teks berbentuk key yang dimasukkan ke kotak teks Basic diterjemahkan menjadi `key = PROJ-123` (bukan pencarian teks).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format key dan riwayat key lama di ISS-001 dan modul 28 (tabel alias key). Aturan Enter dengan key lengkap ini adalah langkah pertama perilaku Enter di SRC-009 butir 5 dan sama dengan NAV-014 butir 3 ([modul 21](21-navigasi-ux-preferensi.md)); NAV-014 hanya merender shell UI, logika pengenalan key dan redirect dimiliki fitur ini.

### SRC-011 Recently viewed items

- **Deskripsi:** Pemakaian riwayat item yang terakhir dilihat oleh fitur pencarian: dropdown quick search, menu Filters, system filter "Viewed recently", fungsi JQL `issueHistory()`, dan field sort `lastViewed`. Penyimpanan riwayat dan tampilan navigasinya (flyout Recent, For you) dimiliki NAV-011 Recent Items.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Setiap pengguna login atas riwayatnya sendiri; tidak dapat dilihat pengguna lain atau admin di UI.
- **Aturan perilaku (acceptance criteria):**
  1. Riwayat disimpan per pengguna dan diurutkan berdasarkan waktu akses terakhir menurun; membuka ulang item memindahkannya ke posisi teratas tanpa duplikasi.
  2. Riwayat mencakup work item, board, backlog, project/space, filter, dashboard, plan, dan queue (Service Management), masing-masing dengan tipe objek agar dapat ditampilkan per kategori; daftar tipe ini sama dengan NAV-011 butir 1.
  3. System filter "Viewed recently" menjalankan `issuekey in issueHistory() ORDER BY lastViewed DESC` (SRC-040).
  4. ⚠ `issueHistory()` mengembalikan maksimal 50 work item terakhir per pengguna yang diambil dari penyimpanan NAV-011 (retensi 200 item per tipe); entri di luar 50 teratas tidak ikut dalam hasil fungsi.
  5. Item yang terhapus atau tidak lagi dapat diakses pengguna otomatis disembunyikan dari semua tampilan riwayat saat dibaca (difilter permission), tanpa menunggu pembersihan data.
  6. Field `lastViewed` tersedia sebagai kunci sort di JQL dan kolom list view, berisi waktu pengguna terakhir membuka work item.
  7. Riwayat yang sama dipakai halaman For you / Your work (modul 21).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Satu sumber data dengan NAV-011 Recent Items ([modul 21](21-navigasi-ux-preferensi.md)): NAV-011 memiliki pencatatan, retensi, sinkronisasi lintas perangkat, dan REST API recent projects; fitur ini hanya membaca penyimpanan yang sama untuk kebutuhan pencarian dan JQL, sehingga tidak ada tabel riwayat kedua. Prioritas disamakan dengan NAV-011 (P0 / Fase 1) karena quick search (SRC-009 butir 2) membutuhkannya. Tabel usulan bersama `user_item_history(user_id, object_type, object_id, viewed_at)`; pencatatan dilakukan asinkron agar membuka work item tidak melambat.

### SRC-012 Search results page (full-page search)

- **Deskripsi:** Halaman hasil pencarian penuh yang dibuka dari quick search, dengan tab per tipe konten dan facet untuk mempersempit hasil tanpa JQL.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Semua pengguna login; hasil permission-aware.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman memiliki tab per produk/tipe konten (Work items, Projects, Boards, Filters, Dashboards, dan tipe lain dari modul aktif); tab aktif tercermin di URL.
  2. ⚠ Tab Work items menyediakan facet Project, Assignee, Status, dan Last updated (rentang waktu); memilih facet mempersempit hasil seketika dan kombinasi facet digabung dengan AND.
  3. Hasil diurutkan berdasarkan skor relevansi teks, lalu waktu update terbaru sebagai tie-breaker.
  4. Tautan "Buka di advanced search" membuka navigator (SRC-015) dengan JQL setara query teks dan facet aktif.
  5. Hasil selalu difilter permission yang sama seperti SRC-009 butir 8.
  6. Hasil dimuat bertahap (infinite scroll 25 item per batch) dengan indikator memuat dan pesan akhir daftar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sebelum fitur ini dibangun, Enter di quick search langsung membuka advanced search (SRC-009 butir 5); keputusan apakah halaman ini diperlukan tercatat di Pertanyaan terbuka.

### SRC-013 Quick search smart querying

- **Deskripsi:** Kata kunci khusus pada quick search klasik produk referensi (mis. `my`, `overdue`, `r:me`) yang diterjemahkan menjadi kriteria JQL. Hanya didukung resmi di Server/Data Center; produk ini tidak membangunnya dan dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna akhir di produk referensi.
- **Aturan perilaku (acceptance criteria):**
  1. Kata `my` tidak diterjemahkan menjadi `assignee = currentUser()`; diperlakukan sebagai teks biasa (pengguna memakai system filter "My open work items").
  2. Awalan `r:me` / `r:username` tidak dikenali sebagai reporter; diperlakukan sebagai teks.
  3. Kata `overdue` tidak diterjemahkan menjadi `due < now() AND resolution is EMPTY`.
  4. Awalan `created:`, `updated:`, `due:` dengan nilai relatif tidak didukung di quick search; kriteria tanggal dibuat di Basic search (SRC-016).
  5. Nama/key project dalam teks tidak otomatis menjadi filter project, kecuali bias konteks project (SRC-009 butir 9).
  6. Nama issue type, status, atau resolution dalam teks tidak otomatis menjadi kriteria.
  7. Awalan `c:` (component), `v:` (affects version), dan `ff:` (fix version) tidak didukung.
  8. Semua kata diperlakukan sebagai text search sesuai SRC-004.
  9. Hasil "View all" tetap diteruskan ke navigator sebagai JQL yang bisa diedit (`text ~ "..."`), sesuai SRC-009 butir 5.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Quick search Cloud produk referensi yang baru tidak mendukung fitur ini secara resmi; lihat Di luar cakupan. Kebutuhan serupa dilayani natural language ke JQL (SRC-019).

### SRC-014 Rovo Search (cross-product search)

- **Deskripsi:** Mesin indeks dan query pencarian lintas entitas/modul berbantuan AI opsional (work management, service, discovery, knowledge base) yang melayani halaman hasil penuh dan jawaban ringkas bersitasi. Pembagian kepemilikan: modul ini memiliki mesin indeks lintas entitas, ranking, penegakan izin, dan grounding jawaban AI; halaman/UI pencarian lintas produk (filter, kartu hasil, sort, hover card) dimiliki NAV-015 ([21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md)). Konektor ke produk lain tidak dibangun.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna login; Admin Situs mengaktifkan/menonaktifkan fitur AI.
- **Aturan perilaku (acceptance criteria):**
  1. Satu kotak search mencakup semua modul produk ini yang aktif (work item, artikel knowledge base, ideas, project, board, filter, dashboard); konektor ke produk pihak ketiga tidak disediakan (Di luar cakupan).
  2. API query mesin indeks menerima filter tipe konten/modul, kontributor, waktu update, dan project/space; kontrol UI filter tersebut dirender oleh halaman NAV-015 (butir 1-2).
  3. Bila AI diaktifkan, jawaban ringkas tampil di atas hasil disertai sitasi berupa tautan ke item sumber; setiap kalimat jawaban hanya boleh bersumber dari item yang boleh dilihat pengguna.
  4. Dokumen indeks work item menyimpan status dan assignee sehingga kartu hasil di halaman NAV-015 dapat menampilkan pratinjau keduanya tanpa query tambahan.
  5. Hasil dan jawaban AI hanya memakai konten yang boleh diakses pengguna; prompt AI tidak pernah memuat item di luar izin pengguna.
  6. Admin Situs dapat mengaktifkan/menonaktifkan pencarian AI di administrasi (modul 22); saat nonaktif kotak search berperilaku seperti SRC-009/SRC-012.
  7. Ranking memakai relevansi teks ditambah sinyal aktivitas personal (item yang sering dilihat/diubah pengguna, SRC-011).
  8. Hasil work item menyediakan tautan ke advanced search (SRC-015) untuk query terstruktur.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat NAV-015 ([21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md)) untuk halaman/UI pencarian lintas produk; prioritas kedua fitur diselaraskan P2 / Fase 4 dan dirilis bersama (mesin di sini, halaman di NAV-015). Mesin ini dibangun di atas abstraksi `SearchRepository` (SRC-001) dan indeks teks SRC-004. Teamwork Graph dan konektor pihak ketiga produk referensi di luar cakupan. Penyedia model AI menjadi keputusan terbuka (modul 27).

## Navigator: halaman pencarian & penyusunan query

### SRC-015 Advanced search page (Issue navigator / Search work items)

- **Deskripsi:** Halaman pencarian lintas project (`/issues/?jql=…` atau `/issues/?filter=ID`) untuk mencari, memfilter, melihat, menyimpan, membagikan, dan mengekspor work item.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna login dan anonim (bila ada project publik); hasil dibatasi Browse Projects dan issue security.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman dapat diakses dari menu/sidebar Filters > "Search work items" / "View all work items" (SRC-042) dan dari quick search (SRC-009 butir 5).
  2. State halaman tersimpan di URL: `jql`, `filter`, `selectedIssue`, posisi halaman (`startAt`), dan mode tampilan (`view=list|detail`); memuat ulang URL menghasilkan tampilan yang sama.
  3. Header menampilkan nama filter saat membuka filter tersimpan, atau "Search work items" untuk query baru.
  4. Toolbar memuat toggle Basic/JQL, tombol Search, Save filter / Save as, Filter details, Share, menu ••• (Export, Bulk change), tombol spreadsheet (bila aktif), Apps (bila ada app), dan switcher List/Detail; tombol yang belum tersedia di fase berjalan tidak dirender.
  5. ⚠ Membuka halaman tanpa parameter memakai query default `ORDER BY created DESC` dengan kriteria kosong (seluruh work item yang boleh dilihat); filter terakhir pengguna tidak dipulihkan otomatis.
  6. Saat filter tersimpan diubah (query, sort, atau kolom), header menampilkan penanda "Edited" serta tombol Save changes dan Discard (SRC-037).
  7. JQL yang tidak valid menampilkan galat di bawah editor tanpa mengosongkan hasil sebelumnya.
  8. Hasil selalu dibatasi Browse Projects dan issue security pengguna (SRC-001 butir 9).
  9. Pengguna anonim dapat membuka halaman bila ada project publik; toolbar untuk anonim hanya memuat Search, switcher, dan Export yang diizinkan (tanpa Save, Share email, Bulk change).
  10. Tombol Back/Forward browser mengikuti setiap perubahan query yang dieksekusi (setiap eksekusi menambah entri history, perubahan `selectedIssue` mengganti entri).
  11. ⚠ Tombol Refresh menjalankan ulang query yang sama dan memperbarui hasil; saat ada perubahan dari pengguna lain (realtime, modul 27), indikator "Hasil mungkin sudah berubah" muncul tanpa me-refresh otomatis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase 1 hanya mendukung mode Basic dan parameter kriteria Basic; parameter `jql` didukung penuh sejak Fase 2 (SRC-018). Endpoint backend `POST /api/v1/search` (modul 23).

### SRC-016 Basic search mode

- **Deskripsi:** Antarmuka pencarian berbasis dropdown kriteria yang membentuk query di belakang layar sehingga pengguna tidak perlu menulis JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna; daftar nilai di dropdown dibatasi yang boleh dilihat pengguna (project yang bisa di-browse, user sesuai Browse users and groups).
- **Aturan perilaku (acceptance criteria):**
  1. Kriteria default yang selalu tampil adalah Project, Type, Status, Assignee, dan kotak teks "Search" (contains text sesuai SRC-004).
  2. Tombol "+ More" menampilkan semua system field dan custom field yang searchable (Reporter, Priority, Resolution, Labels, Sprint, Fix versions, Components, Created, Updated, Due, Parent, dan lainnya) untuk ditambahkan sebagai kriteria.
  3. Dropdown "More" memiliki kotak pencarian nama field yang memfilter daftar seketika.
  4. Kriteria pilihan memakai dropdown multi-select dengan checkbox dan pencarian nilai; nilai khusus "Current user", "Unassigned", dan "EMPTY" tersedia pada field yang relevan.
  5. Kriteria tanggal menawarkan opsi "Within the last N (menit/jam/hari/minggu)", "More than N ago", "Between" (dua tanggal), "In range" (relatif dari-sampai), dan relatif terhadap hari ini.
  6. Setiap field teks yang searchable dapat ditambahkan sebagai kriteria "contains" tersendiri.
  7. User picker mendukung pemilihan pengguna individual dan group ("Members of <group>").
  8. Custom field company-managed hanya muncul di "+ More" bila context field relevan dengan project/issue type yang sedang dipilih; bila belum ada project dipilih, semua field global tampil.
  9. Status, issue type, dan field team-managed yang namanya sama ditampilkan terpisah dengan label nama project.
  10. Kriteria aktif tampil sebagai tombol ringkas, mis. "Status: To Do +2", dengan tombol × untuk menghapus kriteria itu.
  11. Tombol "Clear" mengembalikan semua kriteria ke kosong dan menghapus kriteria tambahan.
  12. Pencarian berjalan otomatis setiap kali kriteria berubah (dengan debounce untuk teks) dan juga bisa dipicu dengan tombol Search atau Enter.
  13. Query yang memuat `OR`, kurung bersarang, `NOT` kompleks, atau fungsi di luar daftar yang didukung Basic tidak dapat ditampilkan di Basic mode (lihat SRC-017 butir 3).
- **Team-managed vs company-managed:** Lihat butir 8 dan 9; detail di SRC-008.
- **Catatan & ketergantungan:** Kriteria dikompilasi ke AST query yang sama dengan JQL (modul 12) agar konversi dua arah di SRC-017 deterministik.

### SRC-017 Switch between Basic and JQL mode

- **Deskripsi:** Toggle dua arah antara Basic search dan JQL dengan konversi query otomatis.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Toolbar menampilkan toggle "Basic" / "JQL"; mode aktif disorot.
  2. Perpindahan Basic → JQL selalu berhasil, dan JQL hasil konversi kriteria langsung tampil di editor dengan format yang konsisten (urutan klausa sesuai urutan kriteria).
  3. Perpindahan JQL → Basic hanya diizinkan bila query dapat direpresentasikan di Basic; bila tidak, toggle Basic dinonaktifkan dengan tooltip "Query terlalu kompleks untuk ditampilkan di mode Basic".
  4. Mode terakhir yang dipakai diingat per pengguna di server dan dipakai saat membuka navigator berikutnya.
  5. Klausa `ORDER BY` dipertahankan saat berpindah mode ke arah mana pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi memakai AST bersama (SRC-016, modul 12). Preferensi disimpan di tabel preferensi pengguna (modul 21).

### SRC-018 JQL editor in navigator

- **Deskripsi:** Editor JQL di halaman search dengan autocomplete, validasi sintaks, dan eksekusi query. Grammar lengkap didefinisikan di modul 12.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna; saran nilai autocomplete dibatasi entitas yang boleh dilihat pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Autocomplete menyarankan nama field, operator yang valid untuk field tersebut, nilai (status, user, project, version, sprint, filter), fungsi, dan keyword (`AND`, `OR`, `NOT`, `ORDER BY`, `ASC`, `DESC`) sesuai posisi kursor.
  2. Editor memberi highlight sintaks (field, operator, nilai, keyword, fungsi) dan indikator valid/tidak valid yang diperbarui selagi mengetik.
  3. Galat ditampilkan rinci beserta posisinya: field tidak dikenal, nilai tidak valid untuk field, dan operator tidak didukung untuk field (mis. "Operator '>' tidak didukung untuk field 'assignee'").
  4. ⚠ Masalah non-fatal (mis. nilai user atau label yang tidak ditemukan pada operator `in`) ditampilkan sebagai warning kuning; query tetap dijalankan dengan nilai yang valid.
  5. Enter menjalankan query; Shift+Enter menyisipkan baris baru.
  6. Editor multi-line dan dapat diperbesar (resize/expand) untuk query panjang.
  7. Tautan "Bantuan sintaks" membuka dokumentasi JQL produk ini.
  8. Field dapat dirujuk dengan nama atau `cf[ID]`; bila beberapa custom field bernama sama, autocomplete menampilkan semuanya dengan pembeda (project/context) dan menyisipkan `cf[ID]` saat dipilih.
  9. ⚠ Editor menyimpan 10 query terakhir yang dijalankan pengguna dan menampilkannya sebagai saran saat editor kosong.
- **Team-managed vs company-managed:** Field team-managed disarankan dengan label nama project (SRC-008).
- **Catatan & ketergantungan:** Endpoint autocomplete `GET /api/v1/jql/autocomplete` dan validasi `POST /api/v1/jql/parse` didefinisikan di modul 12/23.

### SRC-019 Natural language to JQL (Rovo / Atlassian Intelligence)

- **Deskripsi:** Pengguna menulis pertanyaan dalam bahasa alami, lalu AI mengubahnya menjadi JQL yang dapat ditinjau, diedit, dan dijalankan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna login di site yang mengaktifkan fitur AI; Admin Situs mengaktifkan.
- **Aturan perilaku (acceptance criteria):**
  1. Kotak search navigator memiliki toggle/ikon AI untuk beralih ke mode bahasa alami; toggle hanya tampil bila fitur AI aktif.
  2. Pengguna dapat mengetik kalimat seperti "bug prioritas tinggi yang saya buat minggu lalu" lalu menekan Enter.
  3. JQL hasil generate ditampilkan di editor JQL sebelum atau bersamaan dengan eksekusi, dan dapat diedit manual seperti JQL biasa.
  4. Bila AI gagal menghasilkan JQL yang valid, ditampilkan pesan "Tidak dapat mengubah pertanyaan ini menjadi query. Coba kalimat lain atau gunakan mode JQL." tanpa menjalankan query.
  5. Setiap hasil generate menyediakan tombol umpan balik jempol atas/bawah yang tercatat untuk evaluasi kualitas.
  6. Fitur bergantung pada aktivasi AI oleh Admin Situs dan plan yang mengizinkannya (feature flag, modul 00); bila nonaktif, toggle tidak tampil dan API menolak dengan 403.
  7. AI hanya membentuk query; hasil eksekusi tetap difilter permission pengguna, dan metadata yang dikirim ke model (nama field, project) dibatasi yang boleh dilihat pengguna.
  8. ⚠ Bahasa input yang didukung minimal Bahasa Indonesia dan Inggris.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyedia model dan kuota menjadi keputusan terbuka (modul 27). Pengganti smart querying (SRC-013).

### SRC-020 Open search from other contexts

- **Deskripsi:** Entry point untuk membuka navigator dengan JQL yang sesuai dari board, laporan, dashboard, release, component, dan issue view.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat konteks asal; hasil tetap difilter permission.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman release/version, component, dan Epic menyediakan tautan "View in search" / "View all work items" yang membuka navigator dengan JQL setara (mis. `fixVersion = 10001`, `component = 200`, `parent = PROJ-10`).
  2. Mengklik angka di gadget Filter results, Pie chart, dan Two dimensional stats membuka navigator dengan JQL filter gadget ditambah kriteria segmen yang diklik.
  3. Board menyediakan aksi membuka work item kolom, sprint, atau backlog di navigator dengan JQL filter board ditambah kriteria status/sprint terkait.
  4. Laporan (modul 16) menyediakan drill-down dari titik/segmen ke daftar work item di navigator.
  5. Mengklik label atau component di issue view membuka navigator dengan `labels = "x"` atau `component = "y"`.
  6. Issue view menyediakan tautan ke pencarian child (`parent = KEY`) dan linked items (`issue in linkedIssues(KEY)`).
- **Team-managed vs company-managed:** Pada board team-managed, JQL dasar adalah `project = KEY` alih-alih filter board.
- **Catatan & ketergantungan:** Setiap konteks membangun URL `/issues/?jql=` lewat helper bersama di frontend. Bergantung modul 10, 15, 16, 17.

### SRC-021 Navigator keyboard shortcuts

- **Deskripsi:** Pintasan keyboard untuk navigasi dan aksi cepat pada hasil pencarian.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna; aksi yang dipicu shortcut tetap memeriksa permission terkait.
- **Aturan perilaku (acceptance criteria):**
  1. `/` memfokuskan kotak search (quick search atau editor navigator sesuai konteks).
  2. `j` / `k` memindahkan seleksi ke work item berikutnya/sebelumnya di daftar dan panel Detail.
  3. `o` atau Enter membuka work item terpilih.
  4. Pada work item terpilih: `e` membuka edit, `a` assign, `i` assign to me, `m` menambah komentar, dan `.` membuka palet operasi; shortcut tanpa permission terkait tidak melakukan apa pun dan menampilkan toast "Anda tidak memiliki izin".
  5. ⚠ `u` kembali dari halaman work item ke daftar hasil dengan posisi dan seleksi sebelumnya.
  6. `?` menampilkan dialog daftar shortcut yang berlaku di halaman ini.
  7. Shortcut dapat dimatikan di preferensi personal (modul 21); saat dimatikan, hanya `?` dan Esc yang tetap aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Shortcut tidak aktif saat fokus berada di input teks. Didaftarkan ke katalog shortcut modul 21.

### SRC-022 Legacy issue navigator

- **Deskripsi:** Navigator lama produk referensi dengan panel filter kiri dan menu Tools; sudah dihapus dari Cloud dan hanya ada di Data Center. Produk ini tidak meniru UI lama; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna akhir di produk referensi.
- **Aturan perilaku (acceptance criteria):**
  1. Panel kiri berisi Favorite filters, System filters, dan "Find filters" tidak dibangun; fungsinya disediakan sidebar Filters (SRC-042) dan direktori filter (SRC-041).
  2. Menu Tools (Configure columns, Bulk change halaman ini / semua N, Export) tidak dibangun; fungsinya ada di tombol Columns (SRC-025), Bulk change (SRC-031), dan menu Export (SRC-050).
  3. List view dan Detail view gaya lama tidak dibangun; digantikan SRC-023 dan SRC-024.
  4. Toggle "switch to old/new search" tidak disediakan karena produk ini hanya memiliki satu navigator.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat Di luar cakupan.

## Tampilan & interaksi hasil

### SRC-023 List view (search results)

- **Deskripsi:** Tampilan tabel hasil pencarian dengan satu baris per work item dan kolom yang dapat dikonfigurasi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna; nilai field yang tidak boleh dilihat pengguna (mis. field dengan field-level restriction) tampil kosong.
- **Aturan perilaku (acceptance criteria):**
  1. Hasil tampil sebagai tabel dengan header kolom dan nilai field per baris.
  2. Kolom default adalah Type, Key, Summary, Assignee, Reporter, Priority, Status, Resolution, Created, Updated, dan Due date (dapat diubah lewat SRC-025).
  3. Klik key atau summary membuka work item; di List view membuka halaman penuh, dan di halaman Work items project membuka panel samping (ISS-008).
  4. Teks panjang dipotong dengan ellipsis pada satu baris, dan tooltip menampilkan teks lengkap saat hover.
  5. Lebar kolom dapat diubah dengan drag tepi header, dan urutan kolom dapat diubah dengan drag-and-drop header; perubahan disimpan sesuai sumber kolom aktif (SRC-025).
  6. Header tabel tetap terlihat (sticky) saat scroll vertikal.
  7. Setiap baris memiliki checkbox untuk seleksi bulk, dan header memiliki checkbox "pilih semua di halaman ini".
  8. Status tampil sebagai lozenge berwarna sesuai status category (To Do abu-abu, In Progress biru, Done hijau).
  9. Field user tampil dengan avatar dan nama tampilan.
  10. Klik tengah atau Ctrl/Cmd+klik pada key/summary membuka work item di tab baru (tautan memakai elemen `<a href>` asli).
  11. Bila tidak ada hasil, tabel diganti empty state "No work items were found matching your search" (versi Indonesia: "Tidak ada work item yang cocok dengan pencarian Anda").
- **Team-managed vs company-managed:** Tidak ada perbedaan; nilai field lokal team-managed mengikuti SRC-008 butir 5.
- **Catatan & ketergantungan:** Tabel memakai virtualisasi baris untuk 100+ baris (modul 27).

### SRC-024 Detail view (split view)

- **Deskripsi:** Tampilan dua panel untuk triase cepat: daftar ringkas hasil di kiri dan issue view lengkap di kanan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna; aksi di panel kanan memeriksa permission seperti issue view biasa.
- **Aturan perilaku (acceptance criteria):**
  1. Switcher List ⇄ Detail di toolbar mengganti tampilan tanpa menjalankan ulang query; pilihan terakhir diingat per pengguna di server.
  2. Panel kiri menampilkan ikon issue type, key, summary, dan lozenge status ringkas per baris; baris terpilih disorot.
  3. Panel kanan memuat issue view penuh (ISS-007) yang mendukung edit field, komentar, dan transisi tanpa meninggalkan halaman.
  4. URL menyertakan `selectedIssue=KEY` sehingga tampilan dengan work item terpilih dapat di-deep link; bila `selectedIssue` tidak ada di hasil, work item pertama dipilih.
  5. Tombol `j`/`k` atau panah atas/bawah di panel kiri memindahkan seleksi ke work item berikutnya/sebelumnya dan memuat panel kanan.
  6. Lebar panel dapat diubah dengan drag pemisah; lebar diingat per pengguna di sisi klien.
  7. Work item yang setelah diedit tidak lagi cocok dengan query tetap tampil di panel kiri sampai pengguna me-refresh atau menjalankan ulang query.
  8. Sort dan pagination yang sama berlaku pada panel kiri (SRC-026, SRC-027).
  9. Tombol "Buka di halaman penuh" membuka `/browse/KEY`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada komponen issue view yang dapat di-embed (ISS-008).

### SRC-025 Column configuration

- **Deskripsi:** Pengaturan kolom list view pada tiga level: My defaults (per pengguna), Filter (tersimpan bersama filter), dan System default (oleh admin).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna (My defaults); owner/Editors filter (kolom filter); Admin Situs (System default).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Columns" membuka picker dengan kotak pencarian field, checkbox per field, dan daftar kolom terpilih yang dapat diurutkan.
  2. Semua system field dan custom field yang dapat dilihat pengguna dapat dijadikan kolom, termasuk Σ Time Spent/Σ Original Estimate/Σ Remaining Estimate, Parent, Sprint, Labels, Development, dan Images (thumbnail attachment).
  3. Picker memiliki pilihan sumber kolom: "My defaults", "Filter" (hanya saat membuka filter tersimpan yang punya kolom), dan "System"; sumber aktif ditandai.
  4. Hanya owner dan Editors filter yang dapat menyimpan kolom ke filter; Viewer lain melihat kolom filter tetapi perubahan mereka hanya berlaku sementara atau disimpan ke My defaults.
  5. Kolom filter dipakai oleh ekspor "Excel CSV (filter fields)" (SRC-050) dan email subscription (SRC-047).
  6. Tombol "Restore defaults" mengembalikan kolom ke sumber di level berikutnya (Filter → My defaults → System).
  7. ⚠ Admin Situs mengatur System default columns di administrasi (modul 22, Shared items/navigator); perubahan berlaku bagi pengguna yang tidak memiliki My defaults.
  8. Urutan kolom disimpan bersama pilihan kolom di setiap level.
  9. Kolom untuk field yang tidak relevan bagi suatu work item (tidak ada di context atau project) tampil kosong, bukan galat.
- **Team-managed vs company-managed:** Field team-managed tampil di picker dengan label project (SRC-008).
- **Catatan & ketergantungan:** Tabel usulan `column_layouts(scope_type user|filter|system, scope_id, columns jsonb)`.

### SRC-026 Sorting results

- **Deskripsi:** Mengurutkan hasil dengan klik header kolom atau klausa `ORDER BY` di JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Klik header kolom mengurutkan ascending; klik kedua descending; header menampilkan panah arah sort aktif.
  2. Sort lewat header menulis ulang atau menambahkan klausa `ORDER BY` pada JQL (terlihat di mode JQL sejak Fase 2).
  3. Multi-level sort dapat ditulis di JQL (`ORDER BY priority DESC, updated DESC`); kolom yang diklik menjadi kunci utama dan kunci lain bergeser ke belakang tanpa duplikasi.
  4. ⚠ Tanpa `ORDER BY` dan tanpa kriteria teks, hasil diurutkan `created DESC`, lalu `key DESC` sebagai tie-breaker; dengan kriteria teks berlaku SRC-004 butir 9.
  5. Header kolom yang field-nya tidak dapat di-sort (mis. Description, Images) tidak dapat diklik dan tidak menampilkan kursor pointer.
  6. Sort khusus: Priority mengikuti urutan priority scheme, Status mengikuti urutan status di administrasi, dan Rank mengikuti urutan rank global/board.
  7. Nilai kosong ditempatkan konsisten: di akhir untuk ASC dan di awal untuk DESC (setara `NULLS LAST` pada ASC).
  8. Sort ikut tersimpan ke filter saat Save filter atau Save changes (SRC-036, SRC-037).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semantik `ORDER BY` di modul 12. Fase 1: sort header disimpan sebagai parameter `orderBy` di URL.

### SRC-027 Pagination and result count

- **Deskripsi:** Pembagian hasil per halaman dan tampilan jumlah total hasil.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Ukuran halaman default adalah 50 work item.
  2. ⚠ Navigator memakai kontrol Previous/Next dan nomor halaman; halaman Work items project dan panel kiri Detail view memakai load bertahap saat scroll (batch 50).
  3. Indikator "1–50 dari 1.234" menampilkan rentang dan total hasil.
  4. ⚠ Bila total hasil melebihi 10.000, indikator menampilkan "10.000+" dan pagination dibatasi sampai batas tersebut, dengan saran mempersempit query.
  5. Posisi halaman (`startAt`) tersimpan di URL sehingga dapat di-bookmark dan dipulihkan dengan Back.
  6. Perpindahan halaman memakai urutan stabil (tie-breaker key) sehingga item tidak terduplikasi atau terlewat selama tidak ada perubahan data; perubahan data di antara halaman ditangani best-effort.
  7. Preferensi "Number of work items per page" di profil pengguna (perilaku Data Center) tidak disediakan; ukuran halaman tetap 50 di UI dan dapat diatur lewat parameter API `maxResults`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Respons API memakai `startAt`, `maxResults`, `total` (modul 23); pagination berbasis token dipertimbangkan untuk mode OpenSearch.

### SRC-028 Hierarchy display in results

- **Deskripsi:** Menampilkan hubungan parent-child (level di atas Epic, Epic, work item, subtask) sebagai tree di daftar hasil.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna; parent atau child yang tidak boleh dilihat tidak ditampilkan.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle "Hierarchy" di toolbar mengubah daftar datar menjadi struktur tree; toggle disimpan per pengguna per tampilan.
  2. Baris parent memiliki tombol expand/collapse; status expand diingat selama sesi.
  3. Kolom Parent tersedia di daftar datar maupun tree dan menampilkan key serta summary parent.
  4. ⚠ Child yang tidak cocok query tidak ditampilkan secara default; opsi "Tampilkan child di luar hasil sebagai konteks" menampilkannya dengan gaya redup dan tidak dihitung dalam total.
  5. Subtask tampil di bawah parent-nya pada level terdalam.
  6. Setiap level memiliki indentasi; level kustom di atas Epic (Premium, modul 18) didukung tanpa batas jumlah level yang dikonfigurasi.
  7. Sort berlaku di dalam setiap level secara terpisah; urutan parent tidak diubah oleh nilai child.
- **Team-managed vs company-managed:** Kedua tipe memakai field Parent terpadu (SRC-008 butir 6).
- **Catatan & ketergantungan:** Hierarki work type di modul 05; level di atas Epic di modul 18.

### SRC-029 Inline edit in search results

- **Deskripsi:** Mengubah nilai field langsung di sel tabel hasil tanpa membuka work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues, Assign Issues, atau Transition Issues sesuai field yang diubah.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Klik sel field editable (summary, assignee, priority, labels, status, due date, story points, dan custom field bertipe teks, angka, tanggal, pilihan, user) memunculkan editor inline sesuai tipe field.
  2. Sel Status menampilkan dropdown transisi yang tersedia (WF-011); transisi yang memiliki transition screen membuka dialog transisi.
  3. Validasi field berjalan seperti di issue view dan galat ditampilkan inline di bawah sel tanpa menutup editor.
  4. Editor inline hanya tersedia bila pengguna memiliki permission yang sesuai (Edit Issues, Assign Issues, Transition Issues); sel tanpa izin tidak menampilkan affordance edit.
  5. Company-managed: field yang tidak ada di edit screen work item tersebut tidak dapat diedit inline.
  6. Perubahan diterapkan optimistis di UI dan dikembalikan ke nilai semula dengan toast galat bila server menolak.
  7. Enter menyimpan dan Esc membatalkan; untuk field multi-line, Ctrl/Cmd+Enter menyimpan.
  8. Work item yang setelah diedit tidak lagi cocok query tetap tampil sampai refresh.
- **Team-managed vs company-managed:** Team-managed tidak memiliki edit screen; semua field yang ada di layout work type dapat diedit.
- **Catatan & ketergantungan:** Memakai endpoint edit work item yang sama dengan issue view (modul 06).

### SRC-030 Inline create in work item list

- **Deskripsi:** Membuat work item baru langsung dari baris daftar navigator project tanpa membuka dialog Create.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Create Issues di project terkait.
- **Aturan perilaku (acceptance criteria):**
  1. Baris "+ Create" tersedia di bawah daftar dan di bawah setiap grup bila grouping aktif (SRC-034).
  2. Pengguna memilih issue type, mengetik summary, lalu menekan Enter untuk membuat; baris tetap terbuka untuk entri berikutnya.
  3. ⚠ Nilai grup tempat baris berada (mis. assignee atau status grup) dan nilai filter pill sederhana yang aktif dipakai sebagai nilai awal field.
  4. Bila work type memiliki field wajib selain summary, dialog Create penuh terbuka dengan nilai yang sudah diketik.
  5. Baris "+ Create" hanya tampil bagi pengguna dengan Create Issues.
  6. Work item baru langsung tampil di daftar walaupun tidak cocok dengan query, dengan penanda "Baru dibuat".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Logika create memakai ISS-005. Ketersediaan di navigator global tidak disediakan pada fase ini (hanya navigator project).

### SRC-031 Bulk change from search results

- **Deskripsi:** Memilih banyak work item dari hasil pencarian lalu menjalankan operasi massal (edit, transition, move, delete, watch, archive, change parent).
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan global permission Make bulk changes (PERM-011) ditambah permission project per operasi (Edit, Transition, Move, Delete, Archive).
- **Aturan perilaku (acceptance criteria):**
  1. Seleksi memakai checkbox per baris dan "select all" untuk halaman ini, ditambah opsi "Pilih semua N hasil" yang dibatasi 1.000 work item per operasi; melebihi batas menampilkan pesan batas (SRC-006 butir 2).
  2. Entry point tersedia di menu ••• > Bulk change dan di toolbar aksi yang muncul saat ada work item terpilih.
  3. Wizard berlangsung dalam empat langkah: Choose work items → Choose operation → Operation details → Confirmation, dengan tombol Back di setiap langkah.
  4. Operasi yang tersedia: Edit, Move, Transition, Delete, Watch/Stop watching, Archive (hanya plan Premium/Enterprise), dan Change parent; operasi yang tidak diizinkan bagi pengguna dinonaktifkan dengan alasan.
  5. Bulk edit field multi-value (labels, components, versions, multi-select) menawarkan opsi Add, Replace, Remove, dan Clear.
  6. Opsi "Send notification" aktif secara default; mematikannya hanya bisa dilakukan pengguna dengan Administer Projects pada semua project terkait atau Admin Situs.
  7. Bulk transition mengelompokkan work item per workflow dan transisi yang tersedia; work item yang tidak eligible dilewati dan dicantumkan dengan alasannya di ringkasan.
  8. Bulk move memetakan project tujuan, issue type, status, dan field wajib tujuan, termasuk perpindahan antara team-managed dan company-managed (ISS-032).
  9. Server memeriksa global permission Make bulk changes dan permission project per work item per operasi; work item tanpa izin dilewati dengan alasan, bukan menggagalkan seluruh operasi.
  10. Operasi berjalan sebagai background job dengan progress bar, ringkasan sukses/gagal per work item, dan notifikasi in-app saat selesai (modul 19).
  11. ⚠ Untuk edit field sederhana (assignee, priority, labels, status), tersedia panel samping bulk edit tanpa wizard yang menerapkan perubahan langsung ke item terpilih.
- **Team-managed vs company-managed:** Bulk move lintas tipe project memetakan field lokal team-managed (butir 8).
- **Catatan & ketergantungan:** Logika operasi massal didefinisikan di ISS-035 dan WF-013; fitur ini hanya entry point dari hasil pencarian. Fase mengikuti ISS-035 (Fase 3).

## Navigator dalam konteks project

### SRC-032 Project Work items page (Issues / All work / List)

- **Deskripsi:** Navigator dalam konteks project yang menampilkan semua work item project dengan filter cepat, grouping, dan aksi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects pada project.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman diakses dari sidebar project: project software company-managed memakai item "Work items" (`/projects/KEY/issues`), sedangkan project bisnis dan team-managed memakai tab "List" (gabungan All work dan List).
  2. Query selalu dibatasi `project = KEY`; pembatas ini tidak dapat dihapus di halaman project.
  3. Panel atau dropdown menampilkan system filter (SRC-040) dan starred filter (SRC-039) yang dapat diterapkan dalam konteks project.
  4. Halaman mendukung mode list dan detail seperti navigator global (SRC-023, SRC-024).
  5. Tombol "Open in advanced search" membuka navigator global dengan JQL setara, termasuk `project = KEY`.
  6. Kotak search teks mencari dalam project sesuai SRC-004.
  7. Query aktif dapat disimpan sebagai filter (SRC-036).
  8. Inline create (SRC-030), inline edit (SRC-029), dan bulk actions (SRC-031) tersedia dari daftar sesuai fase dan izinnya.
  9. ⚠ Tampilan default menampilkan semua work item project (resolved dan unresolved) diurutkan `created DESC`; pengguna dapat menerapkan filter "Unresolved" dengan satu klik.
- **Team-managed vs company-managed:** Lihat butir 1 untuk perbedaan entry point.
- **Catatan & ketergantungan:** Navigasi project di PRJ-027; tab List bisnis juga dibahas di modul 18 dan 26.

### SRC-033 Filter pills and quick filters (project-level)

- **Deskripsi:** Kontrol filter cepat berbentuk avatar dan dropdown di halaman Work items/List project.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Deretan avatar assignee memfilter per orang; memilih beberapa avatar menggabungkan dengan OR, dan avatar "Unassigned" tersedia.
  2. Dropdown pill tersedia untuk Type, Status, Assignee, Labels, Priority, Sprint, Parent/Epic, Version, dan Component (pill yang fiturnya tidak aktif di project disembunyikan).
  3. Kotak search teks memfilter daftar sesuai SRC-004.
  4. Tombol "Filter" membuka panel filter lanjutan dengan field tambahan dan opsi tanggal (Created, Updated, Due) seperti Basic search.
  5. Kriteria antar-field digabung dengan AND, sedangkan nilai di dalam satu field digabung dengan OR.
  6. Badge menampilkan jumlah filter aktif, dan tombol "Clear filters" menghapus semuanya.
  7. State filter tercermin di URL sehingga dapat dibagikan dan dipulihkan.
  8. Aksi "Buka di advanced search" mengonversi state pill menjadi JQL setara.
- **Team-managed vs company-managed:** Pill Sprint/Version hanya tampil bila fitur terkait aktif di Features project team-managed.
- **Catatan & ketergantungan:** Berbeda dari quick filter board (modul 10), yang berbasis JQL dan dikonfigurasi admin board.

### SRC-034 Group by in project work items view

- **Deskripsi:** Mengelompokkan work item di tampilan Work items/List project berdasarkan field.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna; drag antar grup memerlukan Edit Issues/Transition Issues.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Tombol "Group" menawarkan opsi None, Status, Assignee, Priority, Sprint, Story point estimate, Category, Parent, dan Work type.
  2. Setiap grup memiliki header yang dapat di-collapse dan menampilkan jumlah item.
  3. Work item dengan nilai kosong dikumpulkan di grup "None" / "Unassigned" yang tampil paling akhir.
  4. Sort aktif berlaku di dalam setiap grup.
  5. ⚠ Men-drag work item ke grup lain mengubah nilai field grup (atau menjalankan transisi untuk Status); bila tidak diizinkan atau transisi tidak tersedia, drop ditolak dengan pesan.
  6. Pilihan grouping disimpan per pengguna per project di server.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Grouping di navigator global tidak disediakan pada fase ini.

### SRC-035 Archived work items search

- **Deskripsi:** Halaman pencarian khusus work item terarsip, karena item terarsip dikecualikan dari pencarian normal.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Situs dan Admin Proyek; Restore memerlukan permission Archive/Restore (PERM-035).
- **Aturan perilaku (acceptance criteria):**
  1. Work item terarsip dikecualikan dari quick search, advanced search, board, filter, dan JQL secara default.
  2. Halaman "Archived work items" tersedia secara global (Admin Situs) dan per project (Admin Proyek), dengan filter project, pengarsip, tanggal arsip, dan teks.
  3. Aksi Restore per baris dan massal tersedia bagi pengguna dengan permission Archive/Restore; work item yang di-restore kembali muncul di pencarian normal.
  4. ⚠ Daftar work item terarsip dapat diekspor ke CSV dengan kolom key, summary, project, pengarsip, dan tanggal arsip.
  5. Membuka work item terarsip menampilkan issue view read-only dengan banner "Work item ini diarsipkan".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Operasi archive di ISS-034 (Fase 3, premium). Gating plan lewat feature flag (modul 00).

## Filter tersimpan: siklus hidup & pemakaian ulang

### SRC-036 Save filter (create saved filter)

- **Deskripsi:** Menyimpan query (JQL, sort, dan kolom) sebagai filter bernama yang dapat dipakai ulang dan dirujuk.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna login; sharing selain Private memerlukan Share dashboards and filters (PERM-009).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Save filter" (query baru) atau "Save as" (filter yang sedang dibuka) membuka dialog berisi Name (wajib), Description (opsional), Viewers & Editors, dan toggle star.
  2. Nama filter harus unik per owner (case-insensitive); nama duplikat ditolak dengan pesan "Anda sudah memiliki filter dengan nama ini".
  3. Nama maksimal 255 karakter dan tidak boleh hanya spasi; deskripsi maksimal 2.000 karakter (usulan).
  4. Filter baru mendapat ID numerik unik dan URL `/issues/?filter=ID`.
  5. Hanya pengguna login yang dapat menyimpan filter; tombol Save tidak tampil bagi anonim dan API mengembalikan 401.
  6. Filter baru otomatis di-star oleh pembuatnya, kecuali toggle star dimatikan di dialog.
  7. Entri sharing awal mengikuti preferensi default sharing pengguna, atau default site bila pengguna tidak mengaturnya (SRC-005 butir 8).
  8. Menyimpan entri sharing selain Private tanpa Share dashboards and filters ditolak; dialog hanya menampilkan opsi Private bagi pengguna tersebut.
  9. "Save as" atau aksi Copy pada filter yang ada membuat salinan baru milik pengguna saat ini dengan nama default "Copy of <nama>", tanpa menyalin subscription.
  10. Filter dapat dirujuk di JQL dengan `filter = ID` atau `filter = "Nama"` (SRC-044 butir 4).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `filters(id, owner_id, name, description, jql, order_by, created_at, updated_at)` di modul 28; event `filter.created`. API Filters di modul 23.

### SRC-037 Edit and update saved filter

- **Deskripsi:** Mengubah query, nama, deskripsi, kolom, permission, dan owner filter yang sudah ada.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Owner filter, Editors filter (SRC-046), dan Admin Situs (change owner).
- **Aturan perilaku (acceptance criteria):**
  1. Mengubah query, sort, atau kolom pada filter yang sedang dibuka memunculkan penanda "Edited" dengan tombol Save changes dan Discard; Discard mengembalikan query tersimpan.
  2. Save changes hanya tersedia bagi owner dan Editors; pengguna lain hanya melihat tombol Save as.
  3. Menu filter (•••) memuat Edit name and description, Edit permissions, Copy, dan Delete sesuai hak pengguna.
  4. Change owner dapat dilakukan owner atau Admin Situs; owner baru harus memiliki akses ke filter (dapat melihatnya setelah perpindahan) dan nama tidak boleh bentrok dengan filter milik owner baru, dengan pesan galat yang menyebut alasan penolakan.
  5. Perubahan query langsung berdampak pada semua konsumen yang merujuk filter: board, gadget, plan, subscription, dan filter lain yang memakai `filter =`.
  6. Bila filter dipakai oleh satu atau lebih board, dialog konfirmasi Save changes menampilkan peringatan berisi daftar board terdampak.
  7. Saat dua editor menyimpan bersamaan, simpanan terakhir yang berlaku (last write wins) dan history filter mencatat kedua perubahan beserta pengubahnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan; board team-managed tidak terpengaruh (SRC-008 butir 4).
- **Catatan & ketergantungan:** Perubahan owner dan permission tercatat di audit log (USR-046). Event `filter.updated`.

### SRC-038 Filter details panel

- **Deskripsi:** Panel metadata filter yang menampilkan owner, deskripsi, permission, dan subscription.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat filter; aksi tertentu hanya bagi owner/Editors/Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Filter details" di header filter membuka panel samping.
  2. Panel menampilkan Owner (avatar dan nama), Description, daftar Viewers, dan daftar Editors.
  3. Panel menampilkan subscription milik pengguna saat ini pada filter ini dengan tombol New subscription, edit, dan delete (SRC-047); owner dan Admin Situs juga melihat subscription group.
  4. Tautan "Edit permissions" tampil bagi owner dan Editors yang berhak (SRC-046 butir 2).
  5. Panel menampilkan jumlah pengguna yang men-star filter.
  6. Aksi Copy, Delete, dan Change owner tampil sesuai hak pengguna (SRC-037 butir 3-4, SRC-043).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Subscription tersedia sejak Fase 3; sebelum itu bagian subscription tidak dirender.

### SRC-039 Star filter (favourite filters)

- **Deskripsi:** Menandai filter sebagai favorit agar muncul di menu/sidebar Filters dan sebagai pilihan cepat di gadget dan board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna login yang dapat melihat filter.
- **Aturan perilaku (acceptance criteria):**
  1. Ikon bintang tersedia di header filter, setiap baris direktori filter, dan menu/sidebar Filters; klik men-toggle star.
  2. Hanya filter yang dapat dilihat pengguna yang dapat di-star; API menolak star pada filter tanpa akses dengan galat generik.
  3. Starred filters tampil di menu/sidebar Filters (SRC-042) sesuai urutan yang dapat diubah pengguna.
  4. Jumlah "Starred by" tampil di direktori filter (SRC-041).
  5. Bila akses pengguna ke filter dicabut, star tetap tersimpan tetapi filter disembunyikan dari semua daftar pengguna itu sampai akses kembali.
  6. Starred filter menjadi daftar pilihan teratas saat memilih filter di gadget dashboard dan saat membuat board dari filter.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tabel `filter_favourites(user_id, filter_id, sequence)`.

### SRC-040 System filters (default filters)

- **Deskripsi:** Filter bawaan yang tersedia bagi semua pengguna tanpa perlu dibuat.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna login (sebagian juga anonim pada project publik); hasil mengikuti permission pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. "My open work items" menjalankan `assignee = currentUser() AND resolution = Unresolved ORDER BY priority DESC, updated DESC`.
  2. "Reported by me" menjalankan `reporter = currentUser() ORDER BY created DESC`.
  3. "All work items" menjalankan `ORDER BY created DESC` (seluruh work item yang boleh dilihat).
  4. "Open work items" menjalankan `resolution = Unresolved ORDER BY priority DESC, updated DESC`.
  5. "Done work items" menjalankan `statusCategory = Done ORDER BY updated DESC`.
  6. "Viewed recently" menjalankan `issuekey in issueHistory() ORDER BY lastViewed DESC` (SRC-011).
  7. "Created recently" menjalankan `created >= -1w ORDER BY created DESC`.
  8. "Resolved recently" menjalankan `resolutiondate >= -1w ORDER BY updated DESC`.
  9. "Updated recently" menjalankan `updated >= -1w ORDER BY updated DESC`.
  10. System filter tidak dapat diubah atau dihapus oleh siapa pun; mengubah query-nya menampilkan tombol Save as untuk menyimpannya sebagai filter pribadi (tersedia sejak Fase 2).
  11. System filter tampil di menu/sidebar Filters dan di halaman Work items project (SRC-032).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase 1 mengeksekusi system filter sebagai AST query bawaan (tidak perlu tabel `filters`); fungsi JQL terkait di modul 12. System filter memakai ID negatif atau slug tetap (mis. `?filter=myopen`) agar tidak bentrok dengan ID filter pengguna.

### SRC-041 Filters directory (View all filters)

- **Deskripsi:** Halaman daftar semua filter yang dapat diakses pengguna, dengan pencarian dan aksi manajemen.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna login; aksi manajemen bagi owner/Editors/Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman diakses dari menu/sidebar Filters > "View all filters" (`/filters`).
  2. Pencarian berdasarkan nama filter, owner, project (sharing ke project), dan group (sharing ke group).
  3. ⚠ Tab tampilan tersedia: "My filters" (milik pengguna), "Starred", dan "All filters" (semua filter yang dapat diakses).
  4. Tabel memuat kolom Star, Name, Owner, Viewers, Editors, Starred by, dan aksi •••.
  5. Setiap kolom kecuali aksi dapat di-sort, dan daftar memakai pagination 25 baris per halaman.
  6. Aksi per baris: Edit name & description, Edit permissions, Copy, Delete, dan Change owner (khusus owner/Admin Situs), masing-masing hanya tampil bila pengguna berhak.
  7. Tombol "Create filter" membuka navigator dengan query kosong.
  8. Bila pengguna belum memiliki filter, tab My filters menampilkan empty state dengan ajakan membuat filter pertama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Direktori admin site-wide terpisah di SRC-049.

### SRC-042 Filters navigation (top nav menu & sidebar)

- **Deskripsi:** Akses filter dari navigasi global dan sidebar.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna login.
- **Aturan perilaku (acceptance criteria):**
  1. Item "Filters" di sidebar navigasi global berisi bagian Recent dan Starred filters, masing-masing maksimal 10 item (usulan).
  2. Item tersebut menyediakan tautan "View all filters" (SRC-041) dan "Search work items" (SRC-015).
  3. Recent filters diambil dari riwayat per pengguna (SRC-011) dan hanya memuat filter yang masih dapat diakses.
  4. Filter yang sedang dibuka ditandai aktif di sidebar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Struktur navigasi global di modul 21. Pada Fase 1, item Filters hanya berisi system filter dan Search work items.

### SRC-043 Delete filter

- **Deskripsi:** Menghapus filter tersimpan dengan menampilkan dampaknya pada objek yang bergantung padanya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Owner filter dan Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog konfirmasi menampilkan dampak: jumlah subscription, jumlah pengguna yang men-star, daftar board, dan daftar gadget yang memakai filter (nama board/dashboard yang tidak boleh dilihat penghapus disamarkan sebagai "1 board pribadi").
  2. Hanya owner dan Admin Situs yang dapat menghapus; Editors tidak dapat menghapus dan API mengembalikan 403.
  3. Board yang kehilangan filter menampilkan galat "Filter board tidak ditemukan" dan meminta board admin memilih filter baru (BRD-003).
  4. Filter lain yang merujuk filter terhapus lewat JQL (`filter = ID`) menjadi galat saat dijalankan (SRC-007 butir 5).
  5. Penghapusan bersifat permanen tanpa trash; event `filter.deleted` diterbitkan dan tercatat di audit log.
  6. Semua subscription terkait ikut terhapus dalam transaksi yang sama.
- **Team-managed vs company-managed:** Board team-managed tidak terdampak.
- **Catatan & ketergantungan:** Daftar "used by" disediakan SRC-044 butir 7.

### SRC-044 Saved filters as data source

- **Deskripsi:** Filter tersimpan dipakai ulang sebagai sumber data oleh board, dashboard, plan, subscription, dan JQL lain.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dan board admin yang memilih filter; konsumen hanya dapat memakai filter yang dapat ia lihat.
- **Aturan perilaku (acceptance criteria):**
  1. Board Scrum/Kanban company-managed selalu berbasis satu saved filter (BRD-003, BRD-010).
  2. Gadget dashboard (Filter results, Pie chart, Two dimensional stats, dan lainnya) memakai filter sebagai sumber (modul 17).
  3. Plans (Premium) dapat memakai filter sebagai issue source (modul 18).
  4. JQL `filter = ID`, `filter = "Nama"`, dan `filter in (…)` menggabungkan kriteria filter lain ke dalam query.
  5. Referensi melingkar (filter A merujuk B yang merujuk A, langsung atau tidak langsung) dideteksi saat menyimpan dan saat eksekusi, dengan pesan "Filter memuat referensi melingkar: A → B → A".
  6. Konsumen hanya dapat memilih filter yang dapat dilihatnya; gadget menampilkan galat "Filter ini tidak dibagikan kepada Anda" bila filter tidak dibagikan ke viewer dashboard.
  7. Saat mengubah atau menghapus filter, UI menampilkan daftar "Used by" berisi board, gadget, plan, subscription, dan filter lain yang merujuknya.
- **Team-managed vs company-managed:** Board team-managed tidak memakai saved filter (langsung berbasis project).
- **Catatan & ketergantungan:** Tabel referensi usulan `filter_usages(filter_id, consumer_type, consumer_id)` dipelihara oleh modul konsumen.

## Berbagi, izin & langganan filter

### SRC-045 Filter sharing (Viewers)

- **Deskripsi:** Mengatur siapa yang dapat melihat dan memakai filter: private, user, group, project, project role, semua pengguna login, atau publik.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Owner filter dan Admin Situs; memerlukan Share dashboards and filters (PERM-009).
- **Aturan perilaku (acceptance criteria):**
  1. Sharing default adalah Private (hanya owner), atau mengikuti "Default sharing for filters and dashboards" di Default user preferences (SRC-005 butir 8).
  2. Filter dapat dibagikan ke User tertentu.
  3. Filter dapat dibagikan ke Group.
  4. Filter dapat dibagikan ke Project (semua pengguna yang dapat browse project) atau Project + Role tertentu.
  5. Filter dapat dibagikan ke semua pengguna login di site/organisasi.
  6. Filter dapat dibagikan ke Public (siapa pun di web) hanya bila Admin Situs mengaktifkan public sharing (SRC-005 butir 7); memilih opsi ini menampilkan peringatan risiko yang harus dikonfirmasi.
  7. Beberapa entri sharing dapat digabung pada satu filter (mis. Group A dan Project X role Developers).
  8. Pengguna hanya dapat membagikan ke group tempat ia menjadi anggota dan ke project/role yang dapat ia akses; Admin Situs dikecualikan.
  9. Menyimpan entri sharing memerlukan Share dashboards and filters; tanpa permission ini hanya Private yang tersedia.
  10. Viewer tetap hanya melihat work item yang diizinkan permission-nya sendiri; berbagi filter tidak pernah memperluas akses ke work item.
  11. Entri sharing otomatis dihapus bila group, project, atau role yang dirujuk dihapus.
- **Team-managed vs company-managed:** Project role team-managed (Administrator/Member/Viewer) tersedia sebagai target Project + Role.
- **Catatan & ketergantungan:** Model sharing dipakai bersama dashboard (modul 17) dan didefinisikan di PERM-051. Tabel `filter_shares(filter_id, share_type, target_id, role_id, rights view|edit)`.

### SRC-046 Filter edit permissions (Editors)

- **Deskripsi:** Memberi hak ubah filter kepada pengguna, group, project, atau role selain owner.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Owner filter menetapkan; Editors memakai hak; memerlukan Share dashboards and filters.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog permission memiliki bagian "Editors" yang terpisah dari "Viewers".
  2. ⚠ Editors dapat menyimpan perubahan query, nama, deskripsi, dan kolom filter; Editors juga dapat mengubah daftar Viewers dan Editors, tetapi tidak dapat menghapus entri yang memberi hak edit kepada owner.
  3. Hanya owner dan Admin Situs yang dapat menghapus filter atau mengganti owner.
  4. Target Editors adalah User, Group, Project, atau Project role; opsi Public dan semua pengguna anonim tidak tersedia.
  5. Setiap Editor otomatis juga Viewer tanpa perlu entri Viewer terpisah.
  6. Menambah Editors memerlukan Share dashboards and filters.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan sebagai entri `filter_shares` dengan `rights = edit` (SRC-045).

### SRC-047 Filter subscriptions

- **Deskripsi:** Pengiriman hasil filter terjadwal lewat email ke diri sendiri atau ke group.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna yang dapat melihat filter (subscription personal); Manage group filter subscriptions (PERM-010) untuk group; owner filter dan Admin Situs untuk mengelola.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog New subscription memuat Recipient (Personal atau Group), Schedule, dan opsi "Email this filter even if there are no work items" (default nonaktif).
  2. Jadwal mendukung Daily (sekali sehari pada jam tertentu, atau setiap N jam/menit dalam rentang jam), Days per week (hari tertentu), Days per month (tanggal tertentu atau pola seperti "first Monday"), dan Advanced (cron expression yang divalidasi dengan pratinjau jadwal berikutnya).
  3. Jadwal dievaluasi dalam zona waktu pembuat subscription (fallback zona waktu site).
  4. Recipient = Group hanya tersedia bagi pengguna dengan Manage group filter subscriptions.
  5. ⚠ Email berisi daftar work item dengan kolom filter (SRC-025 butir 5), maksimal 200 work item per email dengan keterangan bila dipotong, serta tautan ke filter.
  6. Hasil dievaluasi terpisah dengan permission setiap penerima sehingga penerima hanya melihat work item yang boleh ia lihat; penerima yang tidak dapat melihat filter tidak dikirimi.
  7. Subscription dikelola di Filter details (SRC-038): edit, delete, dan "Run now" yang mengirim segera tanpa mengubah jadwal.
  8. Penerima nonaktif atau terhapus dilewati.
  9. Format email HTML atau teks mengikuti preferensi setiap penerima (modul 19/21).
  10. Subscription ikut terhapus saat filter dihapus (SRC-043 butir 6).
  11. Subscription milik pengguna yang dinonaktifkan tetap terlihat dan dapat dihapus atau dialihkan oleh owner filter dan Admin Situs di UI (perbaikan atas batasan produk referensi Cloud yang menyembunyikannya), dan pengirimannya dihentikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase dan prioritas disamakan dengan NTF-030 Filter Subscriptions (Email delivery) di [19 Notifikasi & Email](19-notifikasi-email.md) (P1 / Fase 2): fitur ini memiliki model subscription, jadwal, dan evaluasi hasil per penerima, sedangkan NTF-030 memiliki template dan pengiriman email. Penjadwalan memakai background scheduler (modul 27). Tabel `filter_subscriptions(id, filter_id, owner_id, group_id, schedule, cron, email_on_empty, last_run_at, next_run_at)`.

### SRC-048 Share search (email share & permalink)

- **Deskripsi:** Membagikan hasil pencarian atau filter lewat URL permanen atau email langsung.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna login (email share); semua pengguna dapat menyalin tautan.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Share" membuka dialog untuk mengirim email berisi tautan filter/query ke pengguna atau alamat email, dengan catatan opsional maksimal 1.000 karakter.
  2. Aksi "Copy link" menyalin URL berisi `?jql=` (query belum disimpan) atau `?filter=ID` (filter tersimpan) ke clipboard dan menampilkan konfirmasi.
  3. Permalink query yang belum disimpan dapat dibuka siapa pun yang dapat mengakses navigator; hasilnya mengikuti permission penerima.
  4. Permalink filter tersimpan mewajibkan penerima menjadi Viewer; bila tidak, ditampilkan pesan tidak punya akses (SRC-007 butir 3).
  5. JQL di URL di-encode; bila panjang URL melampaui batas (SRC-006 butir 4), Copy link menampilkan saran menyimpan sebagai filter.
  6. Tautan dapat menyertakan `selectedIssue` untuk deep link ke work item terpilih.
  7. Teks JQL dapat disalin dari editor dengan tombol salin.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Email dikirim via modul 19; mencerminkan pola Share work item (ISS-038).

### SRC-049 Admin filters management (Shared items > Filters)

- **Deskripsi:** Halaman administrasi untuk mencari dan mengelola filter di seluruh site, termasuk mengganti owner dan menghapus.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > System > Shared items > Filters hanya dapat diakses Admin Situs; pengguna lain mendapat 403.
  2. Pencarian berdasarkan nama filter dan owner.
  3. Daftar mencakup semua filter di site, termasuk filter private yang tidak dibagikan ke admin; JQL filter private hanya terlihat setelah admin mengambil alih ownership (butir 9).
  4. Kolom "Shared with" menampilkan ringkasan sharing dan dapat difilter untuk menemukan filter "Shared with the public".
  5. Aksi ••• > Change owner membuka dialog pemilihan pengguna baru dan konfirmasi; validasi mengikuti SRC-037 butir 4.
  6. Aksi ••• > Delete filter mengikuti dialog dampak SRC-043.
  7. Filter milik pengguna nonaktif/terhapus ditandai "Owner nonaktif" dan dapat dialihkan ke owner aktif agar board/gadget pemakainya tetap berjalan; filter seperti itu tetap dieksekusi dengan permission owner terakhir sampai dialihkan.
  8. Tidak ada aksi bulk change owner di UI; perubahan owner massal dilakukan lewat REST API (modul 23).
  9. Untuk mengedit isi filter orang lain, admin mengambil alih ownership terlebih dahulu lalu mengembalikannya; kedua perubahan owner tercatat.
  10. Setiap perubahan owner dan penghapusan oleh admin tercatat di audit log (USR-046).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bagian Shared items juga memuat dashboard (modul 17, 22).

## Ekspor, cetak & titik ekstensi hasil

### SRC-050 Export CSV (Excel CSV: all / current / filter fields)

- **Deskripsi:** Mengekspor hasil pencarian ke file CSV (comma-separated) dengan seluruh field, kolom yang sedang tampil, atau kolom filter.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat hasil, termasuk anonim pada project publik; dapat dibatasi kebijakan keamanan data (SRC-058).
- **Aturan perilaku (acceptance criteria):**
  1. Menu ••• > Export menyediakan "Excel CSV (all fields)", "Excel CSV (current fields)", dan "Excel CSV (filter fields)"; opsi terakhir hanya tampil saat membuka filter tersimpan yang memiliki kolom.
  2. All fields mencakup semua field termasuk description, comments, attachments (URL), worklogs, links, sprint, dan custom field; field multi-value dipecah menjadi beberapa kolom dengan header yang sama (mis. tiga kolom "Labels").
  3. Current fields dan filter fields mengikuti kolom aktif atau kolom filter beserta urutannya.
  4. ⚠ Hasil hingga 1.000 work item diekspor sinkron dan langsung diunduh; hasil di atas 1.000 dialihkan ke background job asinkron (INT-015) dengan progress bar yang mendukung hingga 10.000 work item, dan file siap diunduh lewat tautan yang berlaku 24 jam (usulan) disertai notifikasi in-app.
  5. Untuk hasil di atas 10.000, UI menyarankan memecah per JQL (mis. per rentang `created`); endpoint URL ekspor menerima parameter `tempMax` (maksimal 1.000 per permintaan), `pager/start`, dan `startAt` untuk pengambilan batch.
  6. ⚠ Delimiter dapat dipilih: koma (default), titik koma (untuk locale yang memakai koma desimal), pipe (`|`), atau caret (`^`).
  7. Tanggal ditulis mengikuti format dan zona waktu preferensi pengguna yang mengekspor.
  8. Hanya field dan work item yang boleh dilihat pengguna yang diekspor; komentar ber-visibility terbatas hanya disertakan bila pengguna berhak.
  9. Nilai yang diawali `=`, `+`, `-`, atau `@` di-escape (diawali tanda kutip tunggal) untuk mencegah CSV/formula injection, dan dialog ekspor menampilkan peringatan keamanan saat file dibuka di spreadsheet.
  10. Format CSV kompatibel untuk diimpor ulang lewat CSV importer (modul 25), termasuk header field multi-value berulang.
- **Team-managed vs company-managed:** Field lokal team-managed diekspor dengan header bernama field plus label project bila bentrok nama.
- **Catatan & ketergantungan:** Lihat INT-021 ([25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md)). SRC-050 adalah pemilik perilaku ekspor CSV (menu, opsi kolom, delimiter, batas sinkron/asinkron, izin); INT-021 merinci format file dan kompatibilitas round-trip dengan CSV importer (INT-009, INT-012) serta job asinkron dan BOM/formula injection (INT-015). Prioritas diselaraskan P0 karena ekspor adalah jalur migrasi data, dengan fase tetap Fase 3 agar dirilis serentak dengan importer CSV. Job ekspor memakai tabel `export_jobs` dan storage file yang sama dengan attachment (modul 08/27). Batas di SRC-006.

### SRC-051 Export XML

- **Deskripsi:** Mengekspor hasil pencarian dalam format XML berbasis RSS lengkap dengan field work item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dan Integrator/Developer API dengan akses ke hasil.
- **Aturan perilaku (acceptance criteria):**
  1. Menu ••• > Export menyediakan "Export XML".
  2. Setiap work item menjadi satu elemen berisi key, summary, description (HTML ter-render), status, priority, assignee, reporter, tanggal (created, updated, due, resolved), custom field, comments, metadata attachment, dan issue links.
  3. Parameter URL `field` (dapat diulang) membatasi field yang disertakan; tanpa parameter semua field disertakan.
  4. Parameter `tempMax` membatasi jumlah work item, maksimal 1.000; nilai lebih besar dipangkas ke 1.000.
  5. Ekspor dapat diakses lewat URL stabil setara `sr/jira.issueviews:searchrequest-xml/...` dengan autentikasi sesi atau token API; tanpa autentikasi hanya data publik yang dikembalikan.
  6. Ekspor XML juga tersedia dari issue view untuk satu work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Untuk integrasi baru, REST API (modul 23) adalah jalur yang direkomendasikan; XML dipertahankan demi kompatibilitas alat yang ada.

### SRC-052 Export Word

- **Deskripsi:** Mengekspor hasil pencarian sebagai dokumen Word (`.doc` berbasis HTML).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat hasil.
- **Aturan perilaku (acceptance criteria):**
  1. Menu ••• > Export menyediakan "Export Word".
  2. Setiap work item tampil sebagai blok lengkap berisi field, description, dan komentar yang boleh dilihat pengguna.
  3. Ekspor di atas 150 work item menampilkan peringatan bahwa dokumen mungkin berat (rekomendasi, bukan batas keras); batas keras mengikuti SRC-006 butir 2.
  4. Ekspor Word juga tersedia dari issue view untuk satu work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dapat memakai template HTML yang sama dengan Print details (SRC-053).

### SRC-053 Printable view (Print list / Print details)

- **Deskripsi:** Tampilan hasil pencarian yang ramah printer.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat hasil.
- **Aturan perilaku (acceptance criteria):**
  1. Menu ••• > Export menyediakan "Print list" (tabel sesuai kolom aktif) dan "Print details" (detail lengkap per work item).
  2. Halaman cetak dibuka di tab baru tanpa chrome aplikasi (navigasi, sidebar) dan siap dicetak dengan Ctrl/Cmd+P memakai stylesheet `@media print`.
  3. Header halaman berisi nama filter atau JQL serta waktu generate dan nama pengguna.
  4. Jumlah work item dibatasi maksimal 1.000 (SRC-006 butir 2); bila dipotong, header menyebut jumlah yang ditampilkan dari total.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tidak ada.

### SRC-054 Export XLSX (form data)

- **Deskripsi:** Mengekspor jawaban form yang terpasang pada work item hasil pencarian ke file XLSX.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Pengguna yang dapat melihat work item dan form-nya.
- **Aturan perilaku (acceptance criteria):**
  1. Menu Export menyediakan "XLSX (form data)" yang meminta pengguna memilih satu form dari daftar form yang ada pada hasil.
  2. Hanya work item hasil pencarian yang memiliki form terpilih yang disertakan; jumlah yang dilewati ditampilkan.
  3. File berisi satu baris per kombinasi work item/form, dengan kolom key, summary, dan satu kolom per pertanyaan form.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada Forms (modul 26; form JSM di modul 29). Menghasilkan XLSX memerlukan library Go seperti excelize (keputusan implementasi).

### SRC-055 Export to Google Sheets / Microsoft Excel

- **Deskripsi:** Konektor untuk menarik hasil JQL/filter langsung ke spreadsheet. Di produk referensi berupa add-on resmi di Google Sheets/Excel; untuk produk ini bersifat opsional, dengan CSV sebagai jalur utama.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna login; data mengikuti permission akun yang diotorisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Ikon Google Sheets dan Microsoft Excel tampil di atas hasil pencarian (bukan di menu Export) hanya bila Admin Situs mengaktifkan konektor.
  2. Pada pemakaian pertama, pengguna login dan mengotorisasi akses konektor ke data produk ini (OAuth 2.0, modul 01/23).
  3. Hasil terbuka di spreadsheet baru dengan kolom sesuai kolom aktif.
  4. Data dapat di-refresh manual atau terjadwal dari add-on di sisi spreadsheet.
  5. Data yang ditarik selalu mengikuti permission akun yang terhubung pada saat refresh.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Add-on di marketplace Google/Microsoft adalah ekosistem pihak ketiga (Di luar cakupan); produk ini cukup menyediakan REST API dan OAuth yang memungkinkannya.

### SRC-056 RSS feeds (RSS and RSS with comments)

- **Deskripsi:** Feed RSS hasil pencarian/filter untuk RSS reader. Masih ada di menu Export Data Center dan tidak lagi tercantum di dokumentasi Cloud; produk ini tidak membangunnya.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna akhir di produk referensi.
- **Aturan perilaku (acceptance criteria):**
  1. Menu Export tidak menyediakan "RSS" maupun "RSS (with comments)".
  2. Feed RSS 2.0 dengan satu item per work item tidak disediakan; kebutuhan pemantauan dilayani subscription email (SRC-047) dan webhook (modul 23).
  3. Varian with comments tidak disediakan.
  4. Tidak ada endpoint feed yang perlu diautentikasi; endpoint XML (SRC-051) tetap memakai autentikasi sesi/token.
  5. Parameter `tempMax` untuk feed tidak berlaku; hanya berlaku pada ekspor XML/CSV.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Endpoint RSS produk referensi Cloud mungkin masih berfungsi walau tidak terdokumentasi; lihat Di luar cakupan.

### SRC-057 Apps menu / app extension points on search results

- **Deskripsi:** Titik ekstensi di toolbar navigator bagi app pihak ketiga untuk menambah aksi atau format ekspor terhadap hasil pencarian.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna akhir (memakai aksi app); developer app (mendaftarkan modul); Admin Situs (memasang app).
- **Aturan perilaku (acceptance criteria):**
  1. Toolbar navigator menampilkan menu "Apps" atau item tambahan di menu Export/••• yang berasal dari app terpasang.
  2. ⚠ App dapat mendaftarkan format ekspor baru atau aksi khusus pada hasil melalui modul ekstensi `search.results.action` dan `search.results.export`.
  3. Saat dipanggil, aksi app menerima konteks berupa JQL aktif, ID filter (bila ada), dan daftar key terpilih; app mengambil data lewat REST API dengan permission pengguna.
  4. Menu Apps hanya tampil bila ada app terpasang yang mendaftarkan modul tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kerangka ekstensi di modul 25. App Marketplace yang memakai titik ini (time tracking, test management, dan sejenisnya) di luar cakupan.

### SRC-058 Data security policies affecting export

- **Deskripsi:** Kebijakan keamanan data tingkat organisasi yang dapat membatasi ekspor dan akses anonim konten work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin organisasi (Admin Situs) yang mengelola kebijakan; semua pengguna terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Rule ekspor membatasi siapa yang boleh mengekspor (semua ekspor SRC-050 s.d. SRC-055 dan endpoint URL ekspor) pada cakupan yang ditentukan; pengecualian dapat diberikan ke group tertentu.
  2. Rule akses anonim memblokir akses publik ke work item pada cakupan tertentu meskipun project memberi Browse Projects ke Public; pencarian anonim pada cakupan tersebut tidak mengembalikan hasil.
  3. Kebijakan diterapkan per cakupan: classification level, project, atau seluruh site.
  4. Saat diblokir kebijakan, item ekspor disembunyikan dari menu, dan permintaan API/URL ekspor ditolak dengan 403 "Ekspor diblokir oleh kebijakan keamanan data organisasi".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi kebijakan di USR-040 dan PERM-065; di produk referensi memerlukan langganan keamanan terpisah, di produk ini dikendalikan feature flag plan (modul 00).

## Pertanyaan terbuka

- Apakah halaman hasil pencarian penuh (SRC-012) perlu dibangun, atau Enter di quick search cukup selalu membuka advanced search seperti usulan default?
- Apakah stemming Bahasa Indonesia (dictionary kustom PostgreSQL) menjadi syarat rilis Fase 1, atau konfigurasi `simple` + `english` dianggap cukup pada awalnya (SRC-001)?
- Pada ambang berapa (jumlah work item atau throughput) mode OpenSearch wajib disediakan, dan apakah mode ini masuk Fase 4 atau ditunda sampai ada pelanggan yang membutuhkannya?
- Berapa batas final timeout query, panjang JQL, dan panjang URL navigator (SRC-006), dan apakah batas dapat dikonfigurasi per instance self-hosted?
- Apakah Editors filter boleh mengubah daftar Viewers/Editors, atau hak itu khusus owner (SRC-046 butir 2)?
- Berapa batas jumlah work item per email subscription, dan apakah lampiran CSV ditawarkan untuk hasil besar (SRC-047)?
- Filter milik pengguna nonaktif: dieksekusi dengan permission owner terakhir, permission viewer, atau dinonaktifkan sampai dialihkan (SRC-049 butir 7)?
- Apakah konektor Google Sheets/Excel (SRC-055) dan natural language ke JQL (SRC-019) termasuk peta jalan produk, dan penyedia model AI mana yang dipakai?
- Apakah ekspor CSV diizinkan bagi pengguna anonim pada project publik, atau hanya bagi pengguna login?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** yang memakai titik ekstensi navigator (time tracking, test management, reporting, ekspor PDF/Excel lanjutan): tidak dibangun sebagai fitur native; produk ini hanya menyediakan titik ekstensi (SRC-057) dan REST API.
- **Add-on spreadsheet resmi** di marketplace Google Workspace dan Microsoft AppSource: dibangun pihak terpisah bila ada; produk ini hanya menyediakan API dan OAuth (SRC-055).
- **Layanan pencarian lintas produk Atlassian** (Rovo Search dengan Teamwork Graph, konektor ke Confluence, Loom, dan aplikasi SaaS pihak ketiga): tidak dibangun; diganti pencarian global lintas modul produk ini (SRC-014).
- **Fitur Data Center-only atau usang (dc-only-or-deprecated)**, dicatat sebagai rujukan saja: halaman admin Indexing dan reindex manual (SRC-003), quick search smart querying (SRC-013), legacy issue navigator dan toggle old/new search (SRC-022), RSS feeds (SRC-056), serta preferensi "Number of work items per page" di profil (SRC-027 butir 7).
- **Langganan keamanan data terpisah** (setara Atlassian Guard) sebagai produk berbayar terpisah: kebijakan setara dibangun native di balik feature flag plan (SRC-058), bukan sebagai produk terpisah.
- **Pengindeksan isi file attachment** (PDF, dokumen Office) untuk pencarian teks: tidak dibangun pada fase mana pun dalam PRD ini (SRC-004 butir 10).
