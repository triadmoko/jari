# 12 — Bahasa Kueri (JQL-kompatibel)

Modul ini mendefinisikan bahasa kueri terstruktur produk ini yang kompatibel dengan JQL: grammar dan struktur clause, keyword logika, aturan value/quoting/escaping, reserved word, operator (equality, perbandingan, list, IS/EMPTY, teks, history WAS/CHANGED beserta predicate), seluruh field sistem dan custom yang dapat dicari, katalog fungsi bawaan (tanggal, user/grup, sprint, versi, hierarki, aktivitas personal), aturan evaluasi berbasis permission, batas eksekusi, katalog pesan galat/peringatan, autocomplete dan perilaku editor JQL, REST API khusus JQL, serta komponen JQL Compiler (JQL → SQL terparameter di PostgreSQL). Modul ini adalah "mesin" kueri yang dikonsumsi banyak modul lain: issue navigator dan saved filter (modul 13), board filter/quick filter/swimlane (modul 10), backlog (modul 11), dashboard gadget (modul 17), otomasi (modul 20), Plans (modul 18), webhook dan REST API search (modul 23), serta queue Service Management (modul 31). Rincian UI navigator, Basic search, dan ekspor tetap di modul 13; di sini hanya aturan bahasa dan evaluasinya. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT claims di context `sess`, pesan galat EN/ID di `internal/shared/constants/error.go`), sehingga seluruh modul ini adalah pekerjaan baru: parser dan AST sebaiknya ditempatkan sebagai paket murni tanpa dependensi infrastruktur (mis. `pkg/jql` untuk lexer/parser/AST), sedangkan resolver field, evaluator fungsi, dan translator SQL berbasis GORM/PostgreSQL ditempatkan di feature module baru (mis. `internal/features/search`) yang mengikuti pola Clean Architecture repo (use case bergantung pada interface repository, bukan implementasi).

Modul terkait:

- [00 Visi & Baseline](00-visi-ruang-lingkup-glosarium-baseline.md) — tier plan & feature flag (VISI-003), terminologi work item.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — grup untuk `membersOf()` (USR-015), akses anonim (USR-024), user picker API (USR-025).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Browse Projects (PERM-020), View voters and watchers (PERM-033), issue security (PERM-042, PERM-043, PERM-044), visibilitas komentar/worklog (PERM-045), anonim (PERM-046), izin filter (PERM-051), arsip (PERM-035, PERM-053).
- [04 Proyek](04-proyek.md) — kategori proyek (PRJ-019), Project Lead (PRJ-020), ganti key (PRJ-022), arsip proyek (PRJ-032), fungsi & field JQL proyek (PRJ-065), team-managed vs company-managed (PRJ-003).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — hierarki (CFG-006, CFG-007), Parent (CFG-013), Sprint (CFG-014), Rank (CFG-015), Story Points (CFG-016), Flagged (CFG-018), Team (CFG-019), Category (CFG-021), field legacy Epic (CFG-024), tipe custom field (CFG-025 s.d. CFG-033), searcher & indexing (CFG-038), resolusi (CFG-055), link type (CFG-056), voting/watching (CFG-057).
- [06 Work Item Inti](06-work-item-inti.md) — key & numbering (ISS-001), priority (ISS-018), resolution (ISS-019), labels (ISS-020), Rank (ISS-022), Flag (ISS-023), parent (ISS-025), issue link (ISS-029), remote link (ISS-030), arsip work item (ISS-034), watcher (ISS-036), vote (ISS-037), security level (ISS-042).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar dan visibilitasnya (CMT-001, CMT-004), lampiran.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status (WF-001), status category (WF-002), semantik resolution (WF-009), riwayat status (WF-014).
- [10 Board](10-board-kolom-kartu.md) — board filter, sub-filter Kanban, quick filter, swimlane, card color.
- [11 Backlog & Sprint](11-backlog-sprint.md) — sprint state dan riwayat sprint (SPR-012), `ORDER BY Rank` (SPR-003).
- [13 Pencarian, Navigator & Filter](13-pencarian-filter-ekspor.md) — issue navigator, Basic search, saved filter, filter subscription, quick search.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — estimasi, worklog, pengaturan jam kerja.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — version sequence, released/archived, component lead.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget berbasis filter/JQL.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — Plans, Start date, Target start/end, sumber issue.
- [20 Otomasi](20-otomasi.md) — JQL condition, lookup issues, scheduled trigger, smart values.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — endpoint `/search/jql`, pagination, webhook berfilter JQL.
- [24 Integrasi DevOps](24-integrasi-devops.md) — data development (PR, commit, build, deployment).
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — entity property, custom JQL function dari ekstensi.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — index pencarian, performa, telemetry, AI opsional.
- [28 Model Data Inti](28-model-data-inti.md) — tabel changelog, status transition, custom field value.
- [31 JSM Queues, SLA & KB](31-jsm-antrian-sla-knowledge-base-laporan.md) — field & fungsi JQL khusus Service Management (SLA, approval, organization, request channel) sebagai konsumen dan perluasan bahasa ini.

## Persona terkait

- **Developer/Anggota Tim** — menulis kueri sehari-hari (`assignee = currentUser() AND statusCategory != Done`), memakai autocomplete dan pesan galat untuk memperbaiki kueri.
- **Scrum Master/Team Lead** — menyusun board filter, quick filter, dan kueri sprint (`sprint in openSprints()`), memakai operator history untuk retrospektif.
- **Product Owner/PM** — menyaring backlog per Epic, versi, label, dan prioritas; memakai fungsi versi dan hierarki.
- **Admin Proyek/Project Lead** — membuat kueri lintas proyek untuk audit konfigurasi; memahami perbedaan team-managed vs company-managed.
- **Admin Situs** — memastikan kueri tidak membocorkan data (permission-aware), mengatur fitur AI, memantau kueri lambat.
- **Stakeholder/Viewer** — menjalankan saved filter bersama dan melihat hasil sesuai izinnya sendiri.
- **Pengguna Tim Bisnis** — memakai Basic search yang dikonversi ke JQL dan kueri bahasa alami.
- **Integrator/Developer API** — memanggil API parse/autocomplete/match/sanitize, mengandalkan format galat yang stabil dan batas eksekusi yang terdokumentasi.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin menulis kueri berbentuk `field operator value` yang digabung AND/OR/NOT dengan tanda kurung agar saya dapat menemukan work item yang tepat tanpa bantuan admin (JQL-001, JQL-002, JQL-007, JQL-009).
- Sebagai Developer/Anggota Tim, saya ingin autocomplete menyarankan field, operator, value, dan fungsi yang valid saat saya mengetik agar saya tidak perlu menghafal sintaks (JQL-013, JQL-055).
- Sebagai Developer/Anggota Tim, saya ingin pesan galat yang menunjuk posisi dan penyebab kesalahan dalam bahasa saya agar kueri dapat segera diperbaiki (JQL-057, JQL-056).
- Sebagai Scrum Master/Team Lead, saya ingin mencari work item di sprint aktif dan yang terbawa dari sprint sebelumnya agar sprint planning dan review akurat (JQL-026, JQL-042).
- Sebagai Scrum Master/Team Lead, saya ingin mencari work item yang statusnya pernah berubah dari "In Progress" ke "Done" oleh orang tertentu dalam rentang waktu agar retrospektif berbasis data (JQL-014, JQL-015, JQL-016).
- Sebagai Product Owner/PM, saya ingin menyaring berdasarkan versi rilis berikutnya dan seluruh turunan sebuah Epic agar saya dapat memantau cakupan rilis (JQL-043, JQL-046, JQL-047).
- Sebagai Developer/Anggota Tim, saya ingin mencari teks dengan stemming, frasa persis, dan wildcard di summary, description, dan komentar agar pencarian teks bebas tetap relevan (JQL-011, JQL-012, JQL-019).
- Sebagai Product Owner/PM, saya ingin memakai tanggal relatif dan fungsi seperti `startOfWeek()` agar saved filter tetap benar setiap minggu tanpa diedit (JQL-023, JQL-024, JQL-038).
- Sebagai Admin Situs, saya ingin setiap kueri hanya mengembalikan work item, komentar, dan value yang boleh dilihat penjalannya agar JQL tidak menjadi celah kebocoran data (JQL-051, JQL-031).
- Sebagai Stakeholder/Viewer, saya ingin saved filter yang dibagikan kepada saya dievaluasi dengan izin dan `currentUser()` saya sendiri agar hasilnya relevan dan aman (JQL-039, JQL-051, JQL-054).
- Sebagai Admin Proyek/Project Lead, saya ingin mereferensi custom field dengan nama atau `cf[id]` dan membedakan field bernama sama antar proyek team-managed agar kueri lintas proyek tidak ambigu (JQL-033, JQL-053).
- Sebagai Pengguna Tim Bisnis, saya ingin mengetik kalimat bahasa alami lalu melihat JQL yang dihasilkan agar saya bisa mencari tanpa mempelajari sintaks (JQL-059, JQL-058).
- Sebagai Integrator/Developer API, saya ingin endpoint untuk mem-parse, memvalidasi, dan mencocokkan JQL dengan daftar work item agar integrasi saya dapat memvalidasi kueri sebelum menyimpannya (JQL-060, JQL-052).
- Sebagai Admin Situs, saya ingin kueri berat dibatasi waktu dan kompleksitasnya serta kueri lambat tercatat agar satu pengguna tidak menurunkan performa seluruh situs (JQL-050, JQL-052).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 27 |
| P1 | 20 |
| P2 | 10 |
| P3 | 4 |
| **Total** | **61** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 21 |
| Fase 2 | 19 |
| Fase 3 | 13 |
| Fase 4 | 7 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: walau JQL secara umum termasuk Fase 2, board filter, backlog, dan pencarian dasar di Fase 1 sudah bergantung pada mesin kueri, sehingga Fase 1 membangun inti mesin: JQL Compiler (JQL-050) dengan grammar, keyword logika, value/quoting, ORDER BY, operator `=`, `!=`, `IN`, `IS EMPTY`, `~`, matrix operator-field, field sistem inti (project, key, parent, teks, status, people, priority, labels, sprint/Rank), `currentUser()`/`membersOf()`, fungsi sprint, evaluasi permission-aware, dan katalog galat. Keputusan arsitektur yang tidak boleh ditunda adalah AST yang stabil dengan posisi galat, translator SQL terparameter, dan penerapan filter Browse Projects + issue security di level SQL sejak kueri pertama, karena semua konsumen (modul 10, 11, 13, 17, 20, 23) akan bergantung padanya. Fase 2 menambah konfigurabilitas pencarian: operator perbandingan, field dan fungsi tanggal, custom field, semantik full-text (PostgreSQL `tsvector`), versi/komponen, time tracking, link, alias, reserved word, autocomplete dan editor, serta round-trip Basic search. Fase 3 menambah operator history WAS/CHANGED (membutuhkan tabel changelog terindeks yang harus sudah ditulis sejak Fase 1 oleh modul 28), REST API JQL publik dan batas eksekusi, konsumen otomasi/dashboard, serta field/fungsi P2. Fase 4 menambah Team, data development, `childIssuesOf()` di atas Epic, entity property, custom JQL function dari ekstensi, dan JQL bahasa alami berbasis AI yang opsional.

## Fitur

## Grammar, struktur & penamaan

### JQL-001 JQL Grammar & Clause Structure

- **Deskripsi:** Struktur dasar bahasa: kueri terdiri dari satu atau lebih clause berbentuk `field operator value`, digabung keyword logika, dengan `ORDER BY` opsional di akhir. Ini fondasi parser JQL di produk ini.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat mencari (Developer/Anggota Tim, Admin, Integrator/Developer API); tidak ada permission khusus untuk mem-parse, hasil eksekusi tunduk pada JQL-051.
- **Aturan perilaku (acceptance criteria):**
  1. Clause sederhana `<field> <operator> <value>` diterima, mis. `project = ABC`, dan menghasilkan node clause dengan field, operator, dan operand value.
  2. Clause fungsi `<field> <operator> <function>(args)` diterima, mis. `assignee = currentUser()`; operand bertipe function menyimpan nama dan daftar argumen.
  3. Clause `issue`/`issuekey IN <function>(...)` diterima untuk fungsi yang mengembalikan daftar work item, mis. `issue in linkedIssues(ABC-1)`.
  4. Clause history `<field> WAS|CHANGED [predicate ...]` diterima dan diparse menjadi node dengan daftar predicate (aturan detail di JQL-014 s.d. JQL-016).
  5. Kueri kosong valid dan mengembalikan semua work item yang boleh dilihat pengguna; pada API `/search/jql` kueri kosong tunduk pada aturan penolakan unbounded query (JQL-052).
  6. Kueri yang hanya berisi `ORDER BY ...` tanpa clause (mis. `ORDER BY created DESC`) valid di navigator dan setara kueri kosong yang diurutkan.
  7. Nama field, operator, keyword, dan nama fungsi case-insensitive: `Project = abc` dan `project = ABC` menghasilkan AST dan hasil yang sama.
  8. Whitespace dan newline tidak signifikan; kueri multi-baris diparse sama dengan versi satu baris.
  9. Komentar di dalam kueri tidak didukung: tidak ada sintaks komentar (`//`, `/* */`, `#`) yang dikenali, sehingga teks tersebut diperlakukan sebagai bagian kueri; karena `/`, `*`, dan `#` adalah reserved character (JQL-004), kueri semacam itu ditolak dengan galat reserved character atau galat sintaks.
  10. Daftar value di dalam kurung dipisah koma, mis. `(A, B, "C D")`; koma ganda atau koma di akhir daftar menghasilkan galat "Expecting either a value, list or function".
  11. Argumen fungsi dipisah koma dan boleh tanpa quote selama tidak mengandung spasi, reserved character, atau reserved word.
  12. Hasil parse memiliki struktur tetap: where-clause (compound dengan operator AND/OR/NOT, field clause berisi field, operator, operand), operand (value, list, function, keyword EMPTY), predicate history, dan order-by (field + arah); struktur ini juga dikembalikan API parse (JQL-060).
  13. Parser deterministik berbasis grammar formal dan setiap galat sintaks menyertakan posisi baris dan kolom (rincian implementasi di JQL-050).
  14. ⚠ Penggantian istilah UI "issue" menjadi "work item" tidak mengubah nama field/fungsi JQL (`issuetype`, `issuekey`, `issueHistory` tetap berlaku); alias baru berbasis "work item" tidak ditambahkan sampai ada keputusan produk.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail tiap operator, field, dan fungsi ada di fitur lain modul ini. Grammar menjadi kontrak bagi modul 10, 11, 13, 17, 20, 23. Lihat juga Pertanyaan terbuka tentang alias "work item".

### JQL-002 Logical Keywords (AND, OR, NOT) & Precedence

- **Deskripsi:** Keyword logika untuk menggabungkan dan menegasi clause, beserta aturan precedence dan pengelompokan dengan tanda kurung.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; tidak ada permission khusus.
- **Aturan perilaku (acceptance criteria):**
  1. `AND` mengembalikan work item yang memenuhi semua clause yang digabung.
  2. `OR` mengembalikan work item yang memenuhi minimal satu clause.
  3. `NOT` menegasi satu clause atau satu grup dalam kurung, mis. `NOT assignee = currentUser()` dan `NOT (status = Done OR resolution is not EMPTY)`.
  4. Precedence evaluasi adalah `NOT` > `AND` > `OR`; `a OR b AND c` dievaluasi sebagai `a OR (b AND c)`, dan tanda kurung mengubah urutan.
  5. Kurung bersarang diizinkan tanpa batas kedalaman eksplisit, hanya dibatasi batas kompleksitas kueri (JQL-052).
  6. Kurung tidak seimbang menghasilkan galat berposisi, mis. "Expecting ')' before the end of the query." atau "Expecting either 'OR' or 'AND' but got ')'".
  7. `AND`, `OR`, `NOT` adalah reserved word; bila dipakai sebagai value harus di-quote (`labels = "and"`), tanpa quote menghasilkan galat reserved word (JQL-004).
  8. `NOT field = value` tidak mengembalikan work item yang field-nya EMPTY (semantik sama dengan `!=`); untuk menyertakannya pengguna harus menambah `OR field IS EMPTY`.
  9. ⚠ Simbol alternatif `&`, `&&`, `|`, `||`, `!` diterima sebagai sinonim `AND`/`OR`/`NOT` demi kompatibilitas kueri lama; autocomplete dan dokumentasi hanya menampilkan bentuk kata kunci.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Translator SQL (JQL-050) harus mempertahankan semantik three-valued (EMPTY tidak cocok dengan negasi) secara eksplisit, bukan mengandalkan `NOT` SQL atas kolom NULL.

### JQL-003 Values, Quoting & Escaping Rules

- **Deskripsi:** Aturan penulisan value literal: quoting, escaping, angka, ID vs nama, dan karakter cadangan.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian dan Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Value boleh di-quote dengan double quote (`"In Progress"`) atau single quote (`'In Progress'`) dengan arti identik.
  2. Value wajib di-quote bila mengandung spasi, reserved character, atau reserved word; tanpa quote parser mengembalikan galat yang menyebut token penyebab.
  3. ⚠ Escape di dalam string memakai backslash: `\"`, `\'`, `\\`, `\n`, `\t`, dan `\uXXXX`; escape lain yang tidak dikenal menghasilkan galat sintaks.
  4. Untuk text search (`~`), karakter khusus mesin pencari di dalam string di-escape dengan double backslash, mis. `summary ~ "\\[ERROR\\]"`.
  5. ⚠ Value numerik pada field berbasis entitas (project, status, issuetype, priority, resolution, version, component, sprint, filter) ditafsirkan sebagai ID terlebih dahulu; bila tidak ada entitas dengan ID itu, dicocokkan sebagai nama.
  6. Nama entitas dicocokkan case-insensitive: `status = done` setara `status = Done`.
  7. Issue key valid berformat `PROJECT-123`; key lama setelah work item dipindah atau proyek ganti key tetap menemukan work item yang sama (ISS-001, PRJ-022).
  8. Karakter Unicode (mis. huruf beraksen, CJK, emoji) diizinkan di dalam value yang di-quote dan dicocokkan apa adanya.
  9. String yang tidak ditutup quote menghasilkan galat berposisi, mis. "Error in the JQL Query: The quoted string 'In Progress' has not been completed. (line 1, character 10)".
  10. Tidak ada tipe boolean; `true`/`false` adalah reserved word dan harus di-quote bila dipakai sebagai value.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada aturan literal; resolusi nama ke banyak ID di team-managed diatur di JQL-053.
- **Catatan & ketergantungan:** Resolusi value ke ID dilakukan resolver JQL-050 sebelum translasi SQL.

### JQL-004 Reserved Words & Reserved Characters

- **Deskripsi:** Daftar kata dan karakter yang dicadangkan bahasa kueri dan harus di-quote atau di-escape bila dipakai sebagai value atau nama field.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian dan Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Reserved character: spasi, `+ . , ; ? | * / % ^ $ # @ [ ]`; bila muncul di value tanpa quote, kueri ditolak dan pengguna harus memakai quote atau escape `\uXXXX`.
  2. Daftar reserved word mengikuti daftar kompatibilitas JQL (a, an, abort, access, add, after, alias, all, alter, and, any, are, as, asc, audit, avg, before, begin, between, boolean, break, by, byte, catch, cf, char, character, check, checkpoint, collate, collation, column, commit, connect, continue, count, create, current, date, decimal, declare, decrement, default, defaults, define, delete, delimiter, desc, difference, distinct, divide, do, double, drop, else, empty, encoding, end, equals, escape, exclusive, exec, execute, exists, explain, false, fetch, file, field, first, float, for, from, function, go, goto, grant, greater, group, having, identified, if, immediate, in, increment, index, initial, inner, inout, input, insert, int, integer, intersect, intersection, into, is, isempty, isnull, join, last, left, less, like, limit, lock, long, max, min, minus, mode, modify, modulo, more, multiply, next, noaudit, not, notin, nowait, null, number, object, of, on, option, or, order, outer, output, power, previous, prior, privileges, public, raise, raw, remainder, rename, resource, return, returns, revoke, right, row, rowid, rownum, rows, select, session, set, share, size, sqrt, start, strict, string, subtract, sum, synonym, table, then, to, trans, transaction, trigger, true, uid, union, unique, update, user, validate, values, view, when, whenever, where, while, with) dan disimpan sebagai satu daftar konstanta yang dipakai parser dan API.
  3. Reserved word tanpa quote menghasilkan galat "'X' is a reserved JQL word. You must surround it in quotation marks to use it in a query." (dilokalisasi, JQL-057).
  4. Reserved character tanpa quote menghasilkan galat "The character 'X' is a reserved JQL character. You must enclose it in a string or use the escape '\uXXXX' instead."
  5. Daftar reserved word dikembalikan API autocomplete data pada properti `jqlReservedWords` (JQL-060).
  6. Nama custom field yang mengandung spasi atau reserved word harus di-quote, mis. `"Story Points" > 3`; tanpa quote menghasilkan galat sintaks.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kelengkapan daftar reserved word divalidasi terhadap dokumentasi JQL saat implementasi; penambahan kata baru adalah perubahan breaking bagi saved filter dan harus disertai migrasi (lihat Pertanyaan terbuka).

### JQL-005 Field Aliases & Naming

- **Deskripsi:** Daftar alias nama field sistem yang diterima JQL, termasuk nama tampilan dan bentuk referensi custom field.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian dan Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. `issuekey`, `key`, `id`, dan `issue` adalah alias satu field identitas.
  2. `issuetype` setara `type`.
  3. `fixVersion` setara nama tampilan "Fix versions", dan `affectedVersion` setara "Affects versions".
  4. `created` = `createdDate`, `updated` = `updatedDate`, `resolved` = `resolutionDate`, `due` = `duedate`.
  5. `originalEstimate` = `timeOriginalEstimate`, `remainingEstimate` = `timeEstimate`, `timeSpent` = `timespent`.
  6. `filter` = `request` = `savedFilter` = `searchRequest`.
  7. Nama tampilan field sistem yang di-quote (mis. `"Due date"`, `"Fix versions"`) diterima dan diresolusi ke field yang sama.
  8. Custom field dapat direferensi dengan nama, `cf[id]`, atau `"Nama[Tipe]"` (JQL-033).
  9. ⚠ Nama tampilan field yang terlokalisasi (mis. bahasa Indonesia) diterima sebagai alias tambahan hanya bila diaktifkan; nama kanonik Inggris selalu diterima dan dipakai saat menyimpan filter.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk field sistem; custom field team-managed lihat JQL-053.
- **Catatan & ketergantungan:** Saat filter disimpan, produk ini mempertahankan teks asli pengguna; normalisasi ke nama kanonik hanya dipakai di AST.

### JQL-006 ORDER BY Clause

- **Deskripsi:** Pengurutan hasil dengan satu atau lebih field beserta arah ASC/DESC.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; board memakai `ORDER BY Rank` untuk ranking.
- **Aturan perilaku (acceptance criteria):**
  1. Sintaks `ORDER BY field1 [ASC|DESC], field2 [ASC|DESC], ...` diterima dan hanya boleh muncul satu kali di akhir kueri.
  2. ⚠ Bila arah tidak ditulis, dipakai arah default per field: ascending untuk sebagian besar field dan descending untuk field tanggal (`created`, `updated`, `resolved`, `lastViewed`); default per field disimpan di metadata field dan diekspos ke autocomplete.
  3. Sort multi-level menerapkan field berikutnya sebagai tie-breaker secara berurutan, lalu `id` sebagai tie-breaker stabil terakhir (JQL-050).
  4. Field yang tidak sortable (`description`, `comment`, `text`, `attachments`, `environment`) menghasilkan galat "Field 'X' does not support sorting."
  5. `ORDER BY Rank ASC` mengurutkan sesuai rank backlog/board (LexoRank, ISS-022); board mensyaratkan filter berakhir dengan `ORDER BY Rank` agar work item dapat di-rank (modul 10).
  6. Sort `status` mengikuti urutan status global, `priority` mengikuti urutan priority global, dan `resolution` mengikuti urutan resolution global, bukan alfabet.
  7. Sort `fixVersion`/`affectedVersion` mengikuti sequence version di proyek, bukan alfabet.
  8. Sort custom field select list mengikuti urutan option, bukan alfabet.
  9. Sort field user (assignee, reporter, creator, custom user picker) mengikuti display name.
  10. ⚠ Nilai EMPTY ditempatkan di akhir untuk ASC dan di awal untuk DESC secara konsisten di semua field.
  11. Sort `key` mengurutkan per key proyek lalu nomor secara numerik (`ABC-2` sebelum `ABC-10`).
  12. Field `sprint`, `lastViewed`, `votes`, `watchers`, `workratio`, `statusCategoryChangedDate`, serta custom field number dan date dapat diurutkan.
  13. Klik header kolom di issue navigator menulis ulang klausa ORDER BY (UI di modul 13); field yang diklik menjadi sort utama.
  14. ⚠ Navigator tanpa ORDER BY memakai urutan default `ORDER BY created DESC`.
  15. Kueri quick filter, swimlane, dan card color menolak klausa ORDER BY dengan galat validasi (modul 10).
- **Team-managed vs company-managed:** Tidak ada perbedaan; urutan status team-managed per proyek diatur di JQL-053.
- **Catatan & ketergantungan:** `CHANGED` dan `WAS` tidak dapat dipakai di ORDER BY (JQL-015). Rank memakai index rank di modul 28.

## Operator

### JQL-007 Equality Operators (=, !=)

- **Deskripsi:** Operator persamaan dan ketidaksamaan untuk mencocokkan nilai field secara tepat.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `=` cocok bila nilai field sama persis dengan value (bukan pencarian teks).
  2. `!=` cocok bila nilai field tidak sama dengan value.
  3. `!=` tidak mengembalikan work item dengan field EMPTY: `assignee != X` tidak menyertakan work item unassigned; pengguna perlu `OR assignee is EMPTY`.
  4. Pada field multi-value (labels, component, fixVersion, multi-select), `=` cocok bila salah satu nilai sama dengan value.
  5. ⚠ Pada field multi-value, `!=` mengecualikan work item bila salah satu nilainya sama dengan value (dan tetap mengecualikan yang EMPTY).
  6. `=` tidak didukung pada text field (summary, description); galat "The operator '=' is not supported by the 'summary' field." dan pengguna diarahkan memakai `~`.
  7. `= EMPTY` dan `!= EMPTY` diterima sebagai alternatif `IS EMPTY` / `IS NOT EMPTY`.
  8. `=` menerima fungsi skalar, mis. `assignee = currentUser()`, `created = now()`.
  9. `=` pada field datetime dengan tanggal tanpa jam hanya cocok tepat pukul 00:00 di timezone pengguna; editor dapat menampilkan peringatan dan menyarankan rentang.
  10. Operator `<>` tidak didukung dan menghasilkan galat sintaks yang menyarankan `!=`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Matrix operator per field di JQL-013.

### JQL-008 Comparison Operators (>, >=, <, <=)

- **Deskripsi:** Operator perbandingan untuk field yang memiliki urutan: tanggal, angka, durasi, version, priority, resolution, dan issue key.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. Didukung pada field date/datetime: `created`, `updated`, `due`, `resolved`, `lastViewed`, `statusCategoryChangedDate`, `worklogDate`, dan custom date/datetime.
  2. Didukung pada field number: `votes`, `watchers`, `workratio`, Story Points, dan custom number.
  3. Didukung pada durasi `originalEstimate`, `remainingEstimate`, `timeSpent` dengan format `2h`, `1d`, `"1w 2d"`.
  4. Didukung pada field version berdasarkan sequence version, mis. `fixVersion >= "2.0"`.
  5. Didukung pada `issuekey`/`id` dalam proyek yang sama, mis. `key > ABC-100`.
  6. Didukung pada `priority` dan `resolution` berdasarkan urutan global, mis. `priority >= High`.
  7. Didukung pada `hierarchyLevel` dan counter `development[...]`.
  8. Tidak didukung pada text, select, labels, user, dan status; galat "The operator '>' is not supported by the 'X' field."
  9. Rentang dapat dibentuk dengan kombinasi, mis. `created >= -7d AND created < startOfDay()`.
  10. Tanggal relatif dan fungsi tanggal diterima di semua operator perbandingan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format tanggal di JQL-024; fungsi tanggal di JQL-038; durasi mengikuti pengaturan time tracking modul 14.

### JQL-009 List Operators (IN, NOT IN)

- **Deskripsi:** Operator untuk mencocokkan salah satu dari beberapa nilai, termasuk hasil fungsi yang mengembalikan daftar.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `IN (v1, v2, ...)` setara rangkaian `OR` dari `=`.
  2. `NOT IN (v1, v2, ...)` setara rangkaian `AND` dari `!=`.
  3. `NOT IN` tidak mengembalikan work item dengan field EMPTY; pengguna perlu `OR field is EMPTY`.
  4. `IN` menerima fungsi list: `membersOf()`, `openSprints()`, `releasedVersions()`, `standardIssueTypes()`, `projectsLeadByUser()`, `linkedIssues()`, `watchedIssues()`, dan fungsi list lain di modul ini.
  5. List boleh memuat EMPTY, mis. `component in (EMPTY, UI)` mengembalikan work item tanpa component atau dengan component UI.
  6. List boleh mencampur value dan fungsi, mis. `assignee in (currentUser(), "<accountId>")`.
  7. `issue in (ABC-1, ABC-2)` / `key in (...)` mengembalikan daftar work item spesifik.
  8. List berisi satu item valid, mis. `status in (Done)`.
  9. ⚠ Jumlah item dalam satu list dibatasi (default konfigurasi 1.000 item) dan kelebihan menghasilkan galat kompleksitas kueri (JQL-052).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fungsi list dievaluasi sebagai pre-resolution ke set ID (JQL-050).

### JQL-010 EMPTY / NULL Keywords & IS / IS NOT Operators

- **Deskripsi:** Keyword EMPTY dan NULL (sinonim) beserta operator IS / IS NOT untuk mencari work item yang field-nya kosong atau berisi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `field IS EMPTY` cocok bila field tidak berisi nilai; `field IS NOT EMPTY` cocok bila berisi minimal satu nilai.
  2. `NULL` adalah sinonim `EMPTY`: `IS NULL` / `IS NOT NULL` menghasilkan hasil yang identik.
  3. `IS` / `IS NOT` hanya menerima EMPTY/NULL; `status IS Done` menghasilkan galat yang menyarankan `=`.
  4. `field = EMPTY` dan `field != EMPTY` diterima sebagai alternatif.
  5. EMPTY boleh muncul di dalam list `IN`, mis. `fixVersion in (EMPTY, "1.0")`.
  6. EMPTY dapat dipakai dengan WAS, mis. `assignee WAS EMPTY`.
  7. ⚠ EMPTY dapat dipakai di predicate `CHANGED FROM`/`TO`, mis. `assignee CHANGED FROM EMPTY`.
  8. Field yang selalu terisi (project, issuetype, status, created; reporter pada umumnya) menerima `IS EMPTY` secara sintaks tetapi selalu bernilai false (`IS NOT EMPTY` selalu true); validator tidak menolak.
  9. Text field: `description is EMPTY` didukung; `summary is EMPTY` valid tetapi praktis tidak pernah cocok karena summary wajib.
  10. `attachments IS EMPTY` / `IS NOT EMPTY` adalah satu-satunya bentuk kueri lampiran (JQL-029).
  11. Clause IS dapat dinegasi dan dikelompokkan dengan NOT dan kurung seperti clause lain.
  12. Daftar field yang mendukung IS ditentukan oleh matrix operator per field (JQL-013); field yang tidak mendukung menghasilkan galat operator tidak didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Translator SQL memetakan EMPTY ke `IS NULL` pada kolom skalar dan `NOT EXISTS` pada relasi multi-value (JQL-050).

### JQL-011 Text Operators (~, !~)

- **Deskripsi:** Operator CONTAINS (`~`) dan DOES NOT CONTAIN (`!~`) untuk pencarian full-text pada text field dengan semantik mesin pencari (berbasis token, bukan substring).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; hasil komentar tunduk pada visibilitas komentar (JQL-051).
- **Aturan perilaku (acceptance criteria):**
  1. `~` didukung untuk `summary`, `description`, `environment`, `comment`, `text`, `worklogComment`, dan custom text field (single-line, paragraph).
  2. `!~` mengembalikan work item yang text field-nya tidak mengandung term; work item dengan field EMPTY tidak disertakan.
  3. `~` tidak didukung untuk select, user, labels, URL, dan version; pengguna diarahkan memakai `=`.
  4. Pencocokan berbasis kata/token: `summary ~ "log"` cocok dengan "log in" tetapi tidak dengan "catalog".
  5. `~` pada field non-text menghasilkan galat "The operator '~' is not supported by the 'X' field."
  6. Pengecekan text field kosong memakai `IS EMPTY`, bukan `~ ""`; `~ ""` menghasilkan galat value kosong.
  7. ⚠ Hasil kueri text diurutkan sesuai ORDER BY (atau urutan default navigator), bukan skor relevansi.
  8. Semantik stemming, frasa, wildcard, dan escaping mengikuti JQL-012.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase 1 dapat memakai PostgreSQL full-text dasar; semantik penuh dilengkapi di Fase 2 (JQL-012).

### JQL-012 Text Search Semantics (stemming, wildcard, phrase, fuzzy, escaping)

- **Deskripsi:** Aturan interpretasi value pada operator `~` oleh index pencarian produk ini.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; visibilitas komentar ditegakkan (PERM-045).
- **Aturan perilaku (acceptance criteria):**
  1. Pencocokan teks case-insensitive.
  2. Stemming bahasa Inggris diterapkan: `customize` juga cocok dengan "customized" dan "customizing".
  3. Stop word (a, an, and, the, is, of, ...) diabaikan; kueri yang seluruhnya stop word tidak memberi hasil dan editor menampilkan peringatan.
  4. ⚠ Beberapa kata tanpa tanda kutip frasa diperlakukan sebagai AND implisit: work item harus mengandung semua kata dalam urutan apa pun.
  5. Frasa persis ditulis dengan quote yang di-escape di dalam string, mis. `summary ~ "\"full screen\""`, dan hanya cocok bila kata muncul berurutan.
  6. ⚠ Wildcard `*` di akhir kata mencocokkan banyak karakter (`win*` cocok "windows", "winter"); wildcard di tengah kata tidak didukung pada rilis awal.
  7. ⚠ Wildcard `?` untuk satu karakter (`te?t`) tidak didukung pada rilis awal dan diperlakukan sebagai karakter biasa yang dihapus tokenizer.
  8. Wildcard di awal kata (`*foo`, `?foo`) ditolak dengan galat "The text query '*foo' for field 'summary' is not allowed to start with a wildcard."
  9. ⚠ Sintaks Lucene lama (fuzzy `roam~`, proximity `"a b"~10`, boosting `^`, operator Boolean/`+`/`-` di dalam string, regex) tidak didukung; karakter tersebut diperlakukan sebagai pemisah token dan editor menampilkan peringatan.
  10. Karakter khusus `+ - & | ! ( ) { } [ ] ^ ~ * ? \ :` harus di-escape dengan `\\`; karakter yang tidak dapat diindeks tidak dapat dicari dan tidak menimbulkan galat.
  11. ⚠ Token bertanda baca (mis. `v1.2.3`, `ERR-500`) diindeks sebagai token utuh dan juga dipecah per bagian, sehingga `~ "ERR-500"` dan `~ "500"` sama-sama cocok; perilaku ini dicakup tes regresi tokenizer.
  12. `text ~` mencari serentak summary, description, environment, comment, dan semua custom text field yang dapat dilihat pengguna.
  13. Komentar dengan visibilitas terbatas (role/group) hanya ikut dicocokkan untuk pengguna yang berhak melihat komentar itu.
  14. Rich text (description, comment) diindeks sebagai plain text hasil ekstraksi dokumen editor (modul 07).
  15. Index eventually consistent: perubahan terbaru boleh muncul dengan jeda singkat (target p95 di bawah 5 detik, modul 27); API menyediakan opsi rekonsiliasi (JQL-052).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Implementasi awal memakai PostgreSQL `tsvector` dengan konfigurasi `english`; mesin pencari terpisah dipertimbangkan di modul 27. Konfigurasi bahasa Indonesia untuk stemming adalah keputusan terbuka.

### JQL-013 Operator-Field Compatibility Matrix

- **Deskripsi:** Aturan operator mana yang valid untuk tiap field atau tipe field, dipakai bersama oleh validator dan autocomplete.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian, Integrator/Developer API; admin custom field menentukan tipe field (modul 05).
- **Aturan perilaku (acceptance criteria):**
  1. Text field menerima `~`, `!~`, `IS`, `IS NOT`.
  2. Select/radio/multi-select/checkbox menerima `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  3. User/group picker menerima `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`; `WAS`/`CHANGED` hanya untuk `assignee` dan `reporter`.
  4. Date/datetime menerima `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  5. Number/durasi menerima `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  6. Version (fixVersion, affectedVersion, version picker) menerima `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, `IS NOT`; `WAS`/`CHANGED` hanya untuk `fixVersion`.
  7. Labels menerima `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  8. Issue key/id menerima `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`; `parent` menerima `=`, `!=`, `IN`, `NOT IN` (plus `IS` per JQL-018).
  9. Status menerima `=`, `!=`, `IN`, `NOT IN`, `WAS`, `WAS IN`, `WAS NOT`, `WAS NOT IN`, `CHANGED`.
  10. Priority dan resolution menerima `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, `IS NOT`, seluruh varian `WAS`, dan `CHANGED`.
  11. ⚠ URL field menerima `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT` dan menolak `~`.
  12. Kombinasi yang tidak valid menghasilkan galat terstandar "The operator 'X' is not supported by the 'Y' field."
  13. Matrix diekspos ke autocomplete sebagai daftar `operators` per field di autocomplete data (JQL-055, JQL-060).
- **Team-managed vs company-managed:** Tidak ada perbedaan; matrix ditentukan oleh tipe field, bukan tipe proyek.
- **Catatan & ketergantungan:** Matrix disimpan sebagai satu tabel deklaratif di kode (tipe field → operator) agar validator, autocomplete, dan dokumentasi tidak menyimpang.

## Operator history

### JQL-014 History Operators (WAS, WAS IN, WAS NOT, WAS NOT IN)

- **Deskripsi:** Operator untuk mencari work item berdasarkan nilai yang pernah atau sedang dimiliki field menurut change history.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Product Owner/PM, Developer; hasil tetap tunduk pada JQL-051.
- **Aturan perilaku (acceptance criteria):**
  1. `field WAS value` cocok bila field pernah atau sedang bernilai value.
  2. `field WAS IN (v1, v2)` cocok bila pernah bernilai salah satu value dalam daftar.
  3. `field WAS NOT value` cocok bila field tidak pernah bernilai value (termasuk saat ini).
  4. `field WAS NOT IN (v1, v2)` cocok bila tidak pernah bernilai satu pun value dalam daftar.
  5. Field yang didukung hanya `assignee`, `fixVersion`, `priority`, `reporter`, `resolution`, dan `status`.
  6. Field lain (custom field, sprint, labels, component, affectedVersion) dengan WAS menghasilkan galat "History searches are not supported for the 'X' field."
  7. EMPTY didukung, mis. `assignee WAS EMPTY` cocok untuk work item yang pernah unassigned.
  8. Fungsi diterima sebagai value, mis. `assignee WAS currentUser()`, `assignee WAS IN membersOf("dev")`.
  9. Predicate `AFTER`, `BEFORE`, `BY`, `DURING`, `ON` diterima; `FROM`/`TO` pada WAS ditolak (JQL-016).
  10. Nilai saat ini ikut dihitung: work item yang sekarang berstatus Done cocok dengan `status WAS Done` meski belum pernah berubah dari Done.
  11. ⚠ Pencocokan nilai historis berbasis ID entitas, sehingga rename status atau version tetap cocok dengan nama barunya; nama lama juga diterima selama masih tercatat di changelog.
  12. Evaluasi memakai tabel change history terindeks per field (modul 28) agar kueri history tidak memindai seluruh changelog.
- **Team-managed vs company-managed:** Tidak ada perbedaan sintaks; status team-managed bernama sama diresolusi ke semua ID (JQL-053).
- **Catatan & ketergantungan:** Changelog wajib ditulis sejak Fase 1 oleh modul 06/09/28 agar operator ini benar untuk data lama. Riwayat status di WF-014.

### JQL-015 History Operator CHANGED

- **Deskripsi:** Operator untuk mencari work item yang nilai field-nya pernah berubah, opsional dengan kondisi waktu, pelaku, nilai asal, dan nilai tujuan.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Product Owner/PM, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. `field CHANGED` cocok bila field pernah berubah kapan pun.
  2. Field yang didukung hanya `assignee`, `fixVersion`, `priority`, `reporter`, `resolution`, dan `status`; field lain menghasilkan galat history tidak didukung.
  3. Predicate `AFTER`, `BEFORE`, `BY`, `DURING`, `ON`, `FROM`, `TO` dapat dikombinasikan bebas dan semuanya harus terpenuhi oleh satu peristiwa perubahan yang sama.
  4. Contoh `status CHANGED FROM "In Progress" TO "Done" BY currentUser() DURING (startOfWeek(), now())` mengembalikan work item yang dipindah pengguna saat ini dari In Progress ke Done minggu ini.
  5. `NOT field CHANGED` mengembalikan work item yang field-nya tidak pernah berubah.
  6. ⚠ `FROM EMPTY` dan `TO EMPTY` diterima untuk perubahan dari/ke nilai kosong.
  7. CHANGED tidak dapat dipakai di ORDER BY; percobaan menghasilkan galat sintaks.
  8. ⚠ Nilai awal saat work item dibuat tidak dihitung sebagai perubahan; `CHANGED` hanya mencatat perubahan setelah create.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada changelog modul 28; kombinasi predicate diterjemahkan menjadi satu subquery `EXISTS` atas baris changelog yang sama (JQL-050).

### JQL-016 History Predicates (AFTER, BEFORE, BY, DURING, ON, FROM, TO)

- **Deskripsi:** Predicate yang mempersempit operator WAS dan CHANGED berdasarkan waktu, pelaku, dan nilai.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna pencarian lanjutan (Scrum Master, PM, Admin Proyek).
- **Aturan perilaku (acceptance criteria):**
  1. `AFTER "date"` membatasi pada perubahan/kondisi setelah tanggal.
  2. `BEFORE "date"` membatasi pada perubahan/kondisi sebelum tanggal.
  3. `ON "date"` membatasi pada satu hari kalender di timezone pengguna.
  4. `DURING ("date1", "date2")` membatasi pada rentang inklusif kedua ujung.
  5. ⚠ `BY user` membatasi pada perubahan oleh pengguna tertentu (accountId atau `currentUser()`); `BY (u1, u2)` dan `BY membersOf("grup")` juga diterima.
  6. `FROM value` / `TO value` khusus untuk CHANGED dan berarti nilai sebelum/sesudah perubahan.
  7. Predicate tanggal menerima tanggal absolut, tanggal relatif (`-1w`), dan fungsi tanggal (`startOfMonth()`).
  8. Beberapa predicate boleh digabung berurutan dalam satu clause, mis. `status WAS "In Progress" BY currentUser() DURING (-2w, now())`.
  9. `FROM`/`TO` pada WAS menghasilkan galat validasi "The predicate 'FROM' is not supported by the 'WAS' operator."
  10. Autocomplete menyarankan predicate yang valid setelah operator, lalu saran value predicate (parameter `predicateName`/`predicateValue`, JQL-060).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aturan tanggal di JQL-024 dan JQL-038.

## Field sistem inti

### JQL-017 Project & Classification Fields (project, category, projectType, issuetype, hierarchyLevel)

- **Deskripsi:** Field untuk memfilter berdasarkan proyek, kategori proyek, tipe proyek, work type (issue type), dan level hierarki.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; proyek yang tidak dapat di-browse tidak dapat dipakai (JQL-051).
- **Aturan perilaku (acceptance criteria):**
  1. `project` menerima key, nama, atau ID dengan operator `=`, `!=`, `IN`, `NOT IN`, `IS`; karena proyek selalu terisi, `project IS EMPTY` selalu false.
  2. `project` menerima fungsi `projectsLeadByUser()`, `projectsWhereUserHasPermission()`, dan `projectsWhereUserHasRole()` (JQL-045).
  3. `category` mencocokkan project category (nama atau ID, PRJ-019) dengan `=`, `!=`, `IN`, `NOT IN`, `IS EMPTY`.
  4. ⚠ `projectType` menerima `software`, `business`, `service_desk`, `product_discovery` dengan `=`, `!=`, `IN`, `NOT IN`.
  5. `issuetype` (alias `type`) menerima nama atau ID serta fungsi `standardIssueTypes()` dan `subtaskIssueTypes()`.
  6. Di team-managed, work type bernama sama di banyak proyek (ID berbeda) semuanya cocok bila dicari dengan nama.
  7. `hierarchyLevel` bernilai -1 (sub-task), 0 (standard), 1 (Epic), dan 2 ke atas untuk level kustom di atas Epic (CFG-007).
  8. `hierarchyLevel` mendukung `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`.
  9. ⚠ Work item di proyek yang diarsip tidak muncul di hasil sampai proyek di-restore (PRJ-032).
  10. Proyek yang tidak dapat di-browse pengguna menghasilkan galat/peringatan "The value 'X' does not exist for the field 'project'." yang sama dengan proyek yang benar-benar tidak ada.
- **Team-managed vs company-managed:** Work type team-managed project-scoped (lihat aturan 6 dan JQL-053).
- **Catatan & ketergantungan:** Field & fungsi JQL proyek juga dirujuk di PRJ-065; level hierarki di CFG-006/CFG-007.

### JQL-018 Issue Identity Fields (issuekey/key/id/issue, parent, Epic Link legacy)

- **Deskripsi:** Field untuk mereferensi work item spesifik dan relasi parent-child.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `issuekey`, `key`, `id`, `issue` adalah alias dan menerima key (`ABC-1`) atau ID numerik.
  2. Operator `=`, `!=`, `IN`, `NOT IN`, `>`, `>=`, `<`, `<=` didukung; perbandingan berlaku dalam proyek yang sama.
  3. Key lama setelah work item dipindah ke proyek lain tetap mencocokkan work item tersebut (riwayat key, ISS-001).
  4. Key dengan prefix proyek huruf kecil (`abc-1`) diterima dan dinormalisasi.
  5. `parent` mencocokkan child langsung (Epic → story, story → sub-task), menerima key/ID dan list `IN`.
  6. `parent IS EMPTY` / `parent IS NOT EMPTY` didukung.
  7. Setelah unifikasi parent, `"Epic Link"` dan `"Parent Link"` masih diterima sebagai alias deprecated dari `parent` dan editor menampilkan peringatan deprecation.
  8. Key yang tidak ada atau tidak dapat dilihat menghasilkan galat "An issue with key 'X' does not exist for field 'key'." tanpa membedakan kedua kasus.
  9. `issue in (...)` menerima fungsi yang mengembalikan work item (mis. `linkedIssues()`, `parentEpic()`, `childIssuesOf()`, `watchedIssues()`).
- **Team-managed vs company-managed:** Tidak ada perbedaan setelah unifikasi parent.
- **Catatan & ketergantungan:** Parent field di CFG-013/ISS-025; field legacy Epic di CFG-024.

### JQL-019 Text Content Fields (summary, description, environment, comment, text)

- **Deskripsi:** Field berbasis teks yang dicari dengan operator `~` / `!~`.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; komentar dibatasi visibilitas.
- **Aturan perilaku (acceptance criteria):**
  1. `summary` mencari judul work item dengan `~` / `!~`.
  2. `description` mencari isi rich text yang diindeks sebagai plain text, dengan `~`, `!~`, `IS EMPTY`, `IS NOT EMPTY`.
  3. `environment` mencari field teks environment (umumnya company-managed).
  4. `comment` mencari isi semua komentar yang dapat dilihat pengguna yang menjalankan kueri.
  5. `text` adalah master-field yang mencari summary, description, environment, comment, dan semua custom text field sekaligus.
  6. Tidak ada field native untuk penulis atau tanggal komentar: clause seperti `commentAuthor = X` atau `commentDate > -1d` ditolak dengan galat field tidak ada. Pendekatan terdekat yang tersedia adalah `issue in updatedBy(user, dateFrom, dateTo)` (JQL-041), yang mencakup aktivitas komentar tetapi juga perubahan lain; field komentar khusus (biasanya disediakan app Marketplace) di luar cakupan.
  7. ⚠ Mention (nama tampilan), emoji (shortcode), dan isi code block diindeks sebagai teks biasa.
  8. Text field tidak sortable; ORDER BY atas field ini menghasilkan galat (JQL-006).
- **Team-managed vs company-managed:** Environment umumnya hanya tersedia di company-managed.
- **Catatan & ketergantungan:** Semantik teks di JQL-012; komentar di CMT-001/CMT-004.

### JQL-020 Status & Resolution Fields (status, statusCategory, resolution)

- **Deskripsi:** Field untuk memfilter berdasarkan status workflow, status category, dan resolution.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `status` menerima nama atau ID dengan `=`, `!=`, `IN`, `NOT IN`, `WAS`, `WAS IN`, `WAS NOT`, `WAS NOT IN`, `CHANGED`.
  2. Status bernama sama di banyak workflow atau proyek team-managed semuanya cocok bila dicari dengan nama.
  3. Status yang tidak ada menghasilkan galat "The value 'X' does not exist for the field 'status'."
  4. `statusCategory` menerima `"To Do"`, `"In Progress"`, `"Done"` (nama atau ID; key `new`/`indeterminate`/`done`) dengan `=`, `!=`, `IN`, `NOT IN`.
  5. ⚠ `statusCategory` tidak mendukung `WAS`/`CHANGED`; percobaan menghasilkan galat history tidak didukung.
  6. `resolution` menerima nama atau ID; nilai khusus `Unresolved` setara `resolution IS EMPTY`.
  7. `resolution` mendukung perbandingan berurutan (urutan global resolution) serta `WAS`/`CHANGED`.
  8. Di team-managed, resolution diset otomatis saat work item masuk status category Done (WF-009), sehingga `resolution = Unresolved` praktis setara `statusCategory != Done` di proyek tersebut.
  9. Sort `status` mengikuti urutan status global.
- **Team-managed vs company-managed:** Lihat aturan 2 dan 8.
- **Catatan & ketergantungan:** Status di WF-001, status category di WF-002, resolusi di CFG-055/WF-009.

### JQL-021 People Fields (assignee, reporter, creator, watcher, voter)

- **Deskripsi:** Field berbasis pengguna untuk mencari work item berdasarkan orang yang terlibat.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; `watcher`/`voter` untuk pengguna lain memerlukan View voters and watchers (PERM-033).
- **Aturan perilaku (acceptance criteria):**
  1. `assignee` mendukung `=`, `!=`, `IN`, `NOT IN`, `IS`, seluruh varian `WAS`, dan `CHANGED`, dengan value accountId, `currentUser()`, atau `membersOf()`.
  2. `reporter` mendukung operator yang sama dengan `assignee`, termasuk `WAS`/`CHANGED`.
  3. `creator` mencocokkan pembuat work item (tidak pernah berubah) dan tidak mendukung `WAS`/`CHANGED`.
  4. `watcher` mencocokkan work item yang di-watch pengguna tertentu; mencari watcher selain diri sendiri memerlukan View voters and watchers, tanpa izin itu clause tidak mencocokkan apa pun untuk proyek tersebut.
  5. `voter` mencocokkan work item yang di-vote pengguna tertentu dengan aturan izin yang sama.
  6. Value pengguna disimpan sebagai accountId; autocomplete menampilkan display name dan avatar lalu menulis accountId ke kueri.
  7. ⚠ Username/userkey legacy tidak diterima (konversi lewat API pdcleaner, JQL-060); alamat email diterima sebagai value hanya bila visibilitas profil pemilik email mengizinkan pengguna yang mencari melihatnya (AUTH-002).
  8. Pengguna nonaktif atau terhapus tetap dapat dicari dengan accountId.
  9. Work item tanpa assignee dicari dengan `assignee IS EMPTY`.
  10. `currentUser()` untuk pengguna anonim tidak mencocokkan apa pun (tidak galat).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model akun di AUTH-001; watcher/vote di ISS-036/ISS-037. Repo saat ini memakai ID integer/UUID tabel `users`; accountId kompatibel adalah string opaque yang dipetakan ke ID internal (modul 01/28).

### JQL-022 Priority & Labels Fields

- **Deskripsi:** Field priority (berurutan) dan labels (multi-value bebas).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `priority` menerima nama atau ID dengan `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, seluruh varian `WAS`, dan `CHANGED`.
  2. Urutan priority untuk perbandingan dan sort mengikuti urutan global priority (ISS-018), bukan urutan di priority scheme.
  3. Priority yang tidak ada di priority scheme suatu proyek tetap valid secara global dan hanya tidak mencocokkan work item di proyek itu (CFG-054).
  4. `labels` mendukung `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`; `~` dan wildcard tidak didukung.
  5. ⚠ Label harus diketik persis tanpa spasi; pencocokan label case-sensitive mengikuti penyimpanan label (ISS-020).
  6. `labels NOT IN (...)` dan `labels != X` tidak mengembalikan work item tanpa label.
  7. Custom field bertipe Labels memakai aturan yang sama dengan `labels`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat Pertanyaan terbuka tentang case-sensitivity label.

### JQL-023 Date & Time Fields

- **Deskripsi:** Field tanggal sistem dan custom yang mendukung tanggal absolut, relatif, dan fungsi.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `created` (alias `createdDate`) adalah waktu pembuatan work item.
  2. `updated` (alias `updatedDate`) adalah waktu perubahan terakhir.
  3. `resolved` (alias `resolutionDate`) adalah waktu resolution diset; kosong bila belum resolved.
  4. `due` (alias `duedate`, `"Due date"`) adalah tanggal tanpa jam.
  5. `lastViewed` adalah waktu terakhir pengguna yang menjalankan kueri melihat work item (per pengguna).
  6. `statusCategoryChangedDate` adalah waktu terakhir status category berubah.
  7. `"Start date"` (timeline/kalender) serta `"Target start"` / `"Target end"` (Plans, modul 18) dapat dicari sebagai field tanggal.
  8. Custom Date Picker (tanggal saja) dan Date Time Picker dapat dicari (CFG-027).
  9. Operator yang didukung: `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  10. Pada field datetime, equality dengan tanggal tanpa jam hanya cocok pukul 00:00 tepat; pada field date-only, equality mencocokkan seluruh hari.
  11. ⚠ Evaluasi tanggal memakai timezone profil pengguna yang menjalankan kueri; untuk filter subscription dipakai timezone penerima.
  12. Semua field tanggal sortable.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format literal di JQL-024; fungsi di JQL-038.

### JQL-024 Date Formats & Relative Date Syntax

- **Deskripsi:** Format literal tanggal yang diterima JQL dan sintaks tanggal relatif.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. Format absolut yang diterima: `"yyyy/MM/dd HH:mm"`, `"yyyy-MM-dd HH:mm"`, `"yyyy/MM/dd"`, `"yyyy-MM-dd"`.
  2. Format yang memuat jam wajib di-quote karena mengandung spasi.
  3. Format relatif `-1d`, `-2w`, `-4h`, `-30m`, `+3d` diterima dengan unit `w` (minggu), `d` (hari), `h` (jam), `m` (menit).
  4. Kombinasi unit seperti `"-1w 2d"` diterima bila di-quote.
  5. ⚠ Format tanpa tanda (mis. `4d`) diperlakukan sebagai positif (masa depan).
  6. Tanggal relatif dihitung dari waktu eksekusi (now), bukan dari awal hari.
  7. ⚠ Unit tahun `y` dan bulan `M` hanya diterima pada increment fungsi `startOf*`/`endOf*`; pada literal relatif ditolak dengan galat format tanggal.
  8. ⚠ Konversi `w`/`d` pada literal tanggal mengikuti kalender nyata (1w = 7 hari, 1d = 24 jam), bukan jam kerja time tracking.
  9. Format tidak valid menghasilkan galat "Date value 'X' for field 'created' is invalid. Valid formats include: 'yyyy/MM/dd HH:mm', 'yyyy-MM-dd HH:mm', 'yyyy/MM/dd', 'yyyy-MM-dd', or a period format e.g. '-5d', '4w 2d'."
  10. Tanggal relatif juga berlaku di predicate WAS/CHANGED dan argumen fungsi (`updatedBy`, `startOfDay`).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format durasi time tracking (yang memakai jam kerja) berbeda dan diatur di JQL-028/modul 14.

### JQL-025 Engagement Count Fields (votes, watchers)

- **Deskripsi:** Field numerik untuk jumlah vote dan jumlah watcher.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `votes` adalah jumlah vote dengan operator `=`, `!=`, `>`, `>=`, `<`, `<=`.
  2. `watchers` adalah jumlah watcher dengan operator numerik yang sama.
  3. Kedua field sortable, mis. `ORDER BY votes DESC`.
  4. ⚠ Bila voting atau watching dinonaktifkan global (CFG-057), field terkait tetap valid secara sintaks tetapi dihitung bernilai 0 dan autocomplete menyembunyikannya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penghitung disimpan denormalisasi pada work item agar sort murah (modul 28).

## Field agile, rilis, waktu & relasi

### JQL-026 Agile Fields (sprint, Rank, Flagged, Story Points)

- **Deskripsi:** Field yang dipakai board, backlog, dan sprint.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer, Product Owner.
- **Aturan perilaku (acceptance criteria):**
  1. `sprint` menerima nama atau ID sprint dengan `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  2. `sprint in openSprints()` / `closedSprints()` / `futureSprints()` didukung (JQL-042).
  3. `sprint IS EMPTY` mengembalikan work item di backlog (tanpa sprint).
  4. Work item yang terbawa (carry-over) tercatat di beberapa sprint (sprint closed dan sprint aktif), sehingga cocok dengan clause untuk masing-masing sprint (SPR-012).
  5. Nama sprint yang duplikat antar board semuanya cocok bila dicari dengan nama; autocomplete menampilkan board dan menulis ID untuk presisi.
  6. `Rank` hanya dapat dipakai di `ORDER BY Rank ASC|DESC`; dipakai sebagai clause filter menghasilkan galat.
  7. `Flagged` adalah checkbox dengan option `Impediment`; `Flagged = Impediment` dan `Flagged is not EMPTY` mengembalikan work item yang di-flag (ISS-023).
  8. `"Story Points"` (company-managed) dan `"Story point estimate"` (team-managed) adalah field number dengan operator numerik (CFG-016).
  9. Field legacy company-managed `"Epic Name"`, `"Epic Status"`, `"Epic Colour"` diterima sebagai deprecated dan editor menampilkan peringatan (CFG-024).
  10. `sprint` tidak mendukung `WAS`/`CHANGED`.
- **Team-managed vs company-managed:** Story point berbeda nama (aturan 8); field Epic legacy hanya di company-managed.
- **Catatan & ketergantungan:** Sprint di CFG-014/SPR-012, Rank di CFG-015/ISS-022.

### JQL-027 Version & Component Fields (fixVersion, affectedVersion, component)

- **Deskripsi:** Field untuk version rilis dan component proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian, Product Owner/PM sebagai pengelola rilis.
- **Aturan perilaku (acceptance criteria):**
  1. `fixVersion` menerima nama atau ID dengan `=`, `!=`, `>`, `>=`, `<`, `<=`, `IN`, `NOT IN`, `IS`, seluruh varian `WAS`, dan `CHANGED`.
  2. `affectedVersion` ("Affects versions") mendukung operator yang sama kecuali `WAS`/`CHANGED`.
  3. Perbandingan version memakai sequence version di proyek (modul 15).
  4. Nama version yang sama di beberapa proyek semuanya cocok; untuk presisi pengguna memakai ID atau menggabung dengan clause `project`.
  5. `component` menerima nama atau ID dengan `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  6. `component in componentsLeadByUser()` didukung (JQL-045).
  7. ⚠ Version yang diarsip tetap dapat dicari; autocomplete menampilkannya di bagian terpisah "Archived".
  8. Custom field Version Picker mengikuti aturan yang sama dengan `fixVersion` (tanpa WAS/CHANGED); component yang bersumber dari katalog komponen eksternal mengikuti operator `component`.
- **Team-managed vs company-managed:** Affects versions umumnya hanya di company-managed (JQL-053).
- **Catatan & ketergantungan:** Version dan component di modul 15; fungsi version di JQL-043.

### JQL-028 Time Tracking & Worklog Fields

- **Deskripsi:** Field estimasi, waktu terpakai, rasio kerja, dan atribut worklog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; worklog dengan visibilitas terbatas mengikuti PERM-045.
- **Aturan perilaku (acceptance criteria):**
  1. `originalEstimate` (alias `timeOriginalEstimate`) adalah durasi dengan operator numerik.
  2. `remainingEstimate` (alias `timeEstimate`) adalah durasi dengan operator numerik.
  3. `timeSpent` (alias `timespent`) adalah durasi dengan operator numerik.
  4. `workratio` adalah persentase `timeSpent / originalEstimate * 100`; work item tanpa original estimate tidak memiliki workratio dan tidak cocok dengan perbandingan apa pun.
  5. Format durasi `1h`, `2d`, `"1w 2d"` dikonversi memakai pengaturan time tracking (jam per hari, hari per minggu, modul 14).
  6. `worklogAuthor` mencocokkan pengguna yang mencatat worklog dengan `=`, `!=`, `IN`, `NOT IN`, dan `membersOf()`.
  7. `worklogDate` adalah tanggal worklog dengan operator tanggal.
  8. `worklogComment` adalah teks komentar worklog dengan `~` / `!~`.
  9. ⚠ `worklogAuthor` dan `worklogDate` dalam satu kueri dievaluasi independen dan tidak dijamin mengacu ke worklog yang sama; dokumentasi bantuan menyebut keterbatasan ini.
  10. Bila time tracking dinonaktifkan, field-field ini tetap valid secara sintaks tetapi tidak muncul di autocomplete.
  11. Field agregat (Σ Time Spent, Σ Remaining Estimate) tidak dapat difilter; percobaan menghasilkan galat "Field does not exist".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Time tracking dan worklog di modul 14.

### JQL-029 Link, Attachment & Remote Link Search

- **Deskripsi:** Field dan fungsi terkait issue link, lampiran, dan remote link.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. `issueLinkType` menerima nama arah link (`blocks`, `is blocked by`, `relates to`, `clones`, `is cloned by`, dan link type kustom) dengan `=`, `!=`, `IN`, `NOT IN`.
  2. `issue in linkedIssues(issueKey[, linkType])` mengembalikan work item yang ter-link ke work item tertentu, opsional dibatasi arah link.
  3. `attachments IS EMPTY` / `IS NOT EMPTY` didukung; pencarian nama file atau isi lampiran tidak didukung.
  4. ⚠ `issue in issuesWithRemoteLinksByGlobalId(id1, ...)` mengembalikan work item dengan remote link globalId tertentu, maksimal 100 argumen.
  5. Link ke work item yang tidak dapat dilihat pengguna tidak dihitung di `linkedIssues()` maupun `issueLinkType`.
  6. Bila issue linking dinonaktifkan (CFG-056), `issueLinkType` dan `linkedIssues()` tetap valid tetapi tidak mengembalikan hasil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Link di ISS-029, remote link di ISS-030, link type di CFG-056.

### JQL-030 Saved Filter Reference Field (filter)

- **Deskripsi:** Field untuk menyisipkan hasil saved filter lain ke dalam kueri.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna pencarian lanjutan; akses ke filter yang direferensi diatur PERM-051.
- **Aturan perilaku (acceptance criteria):**
  1. `filter` (alias `request`, `savedFilter`, `searchRequest`) menerima nama filter atau ID.
  2. Operator yang didukung: `=`, `!=`, `IN`, `NOT IN`.
  3. Pengguna harus memiliki akses ke filter yang direferensi; tanpa akses dihasilkan galat "The value 'X' does not exist for the field 'filter'."
  4. Rekursi (filter yang mereferensi dirinya langsung atau melalui rantai) dideteksi dan ditolak dengan galat "Filter 'X' references itself."
  5. Klausa ORDER BY dari filter yang direferensi diabaikan.
  6. ⚠ Referensi berbasis ID tetap valid setelah filter di-rename; referensi berbasis nama gagal dengan galat value tidak ada, dan editor menyarankan memakai ID saat menyimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Saved filter di modul 13; batas kedalaman referensi mengikuti JQL-052.

### JQL-031 Security Level Field (level)

- **Deskripsi:** Field untuk mencari work item berdasarkan issue security level.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin, Admin Proyek, pengguna yang menjadi anggota security level.
- **Aturan perilaku (acceptance criteria):**
  1. `level` menerima nama atau ID security level dengan `=`, `!=`, `IN`, `NOT IN`, `IS`, `IS NOT`.
  2. Hanya level yang dapat dilihat pengguna yang muncul di saran dan lolos validasi; level lain menghasilkan galat value tidak ada.
  3. Hasil selalu difilter Browse Projects dan issue security; tidak ada bentuk kueri yang dapat mem-bypass keamanan (PERM-044).
  4. ⚠ Field `level` hanya berlaku di company-managed; restriksi akses team-managed (PERM-041) tidak dapat dicari lewat `level` pada rilis awal.
- **Team-managed vs company-managed:** Lihat aturan 4.
- **Catatan & ketergantungan:** Skema keamanan di PERM-042/PERM-043; restrict work item di ISS-042.

### JQL-032 Team Field (Atlassian Teams)

- **Deskripsi:** Field Team tingkat platform untuk memfilter work item berdasarkan tim.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Product Owner/PM, Scrum Master, perencana lintas tim.
- **Aturan perilaku (acceptance criteria):**
  1. `team` atau `"Team[Team]"` menerima ID tim (UUID), mis. `team = "a1b2c3d4-..."`.
  2. Autocomplete menampilkan nama tim lalu menulis ID ke kueri.
  3. Operator yang didukung: `=`, `!=`, `IN`, `NOT IN`, `IS EMPTY`, `IS NOT EMPTY`.
  4. Tim adalah entitas platform (direktori People & Teams, USR-027), bukan per proyek; satu tim cocok di semua proyek.
  5. Field Team dipakai di Plans, board filter, dan timeline; semuanya memakai resolver yang sama.
  6. ⚠ Pencarian dengan nama tim (bukan ID) diterima bila nama unik; nama ganda menghasilkan galat ambigu yang meminta ID.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan: definisi field Team (locked field, screen/layout) di [05 CFG-019](05-konfigurasi-work-item-skema.md); field Team di work item dan picker-nya di [18 PLAN-058](18-timeline-rencana-tampilan-proyek.md); entitas tim platform di [18 PLAN-057](18-timeline-rencana-tampilan-proyek.md) dan direktori People & Teams di modul 02 (USR-027). JQL-032 hanya mendefinisikan semantik kueri `team`. Tier mengikuti inventaris dan PLAN-057/PLAN-058, yaitu `platform` (tersedia lintas plan sebagai kemampuan platform); CFG-019 saat ini tertulis `free-standard` dan perlu diselaraskan oleh pemilik modul 05. Produk ini memakai nama "Team" generik tanpa merek.

## Custom field, field platform & data ekstensi

### JQL-033 Custom Field Referencing & Type-specific Behavior

- **Deskripsi:** Cara mereferensi custom field dan perilaku operator berdasarkan tipe field.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin (pengelola custom field), pengguna pencarian lanjutan.
- **Aturan perilaku (acceptance criteria):**
  1. Custom field dapat direferensi dengan nama, mis. `"Story Points" > 3` (di-quote bila mengandung spasi).
  2. Custom field dapat direferensi dengan ID `cf[10010]` yang stabil terhadap rename.
  3. Nama ganda didisambiguasi dengan `"Nama[Tipe]"`, mis. `"Team[Team]"`, `"Story Points[Number]"`, `"Severity[Dropdown]"`; nama ganda tanpa disambiguasi yang tipenya berbeda menghasilkan galat ambigu yang menyarankan bentuk `Nama[Tipe]` atau `cf[id]`.
  4. Di team-managed, field project-scoped bernama sama di banyak proyek ditampilkan autocomplete sebagai varian per tipe; kueri dengan nama mencakup semua field bernama sama yang setipe.
  5. Custom field company-managed dengan searcher "None" (CFG-038) tidak dapat dicari dan menghasilkan galat "Field 'X' is not searchable."
  6. Single-select/radio mendukung `=`, `!=`, `IN`, `NOT IN`, `IS EMPTY` dengan value nama option atau ID option.
  7. Multi-select/checkbox cocok bila salah satu option terpilih; option yang dinonaktifkan tetap dapat dicari.
  8. Cascading select dicari dengan `cascadeOption()` (JQL-048).
  9. User picker (single/multi), group picker, project picker, dan version picker dapat dicari dengan operator sesuai tipe (JQL-013).
  10. Number, date, datetime, URL, text single-line/paragraph, dan labels dapat dicari dengan operator sesuai tipe.
  11. Field dari ekstensi (CFG-033) dapat dicari bila ekstensi mendeklarasikan searcher/index; tanpa itu field tidak muncul di autocomplete.
  12. ⚠ Field yang disembunyikan di field configuration tetap dapat dicari selama pengguna dapat melihat proyeknya; yang disembunyikan hanya tampilannya di layar.
  13. Field yang tidak terlihat pengguna (tidak ada di konteks proyek yang dapat di-browse) menghasilkan galat "Field 'X' does not exist or you do not have permission to view it."
- **Team-managed vs company-managed:** Lihat aturan 4 dan JQL-053.
- **Catatan & ketergantungan:** Tipe custom field di CFG-025 s.d. CFG-033, konteks di CFG-035, searcher di CFG-038. Translator memakai join `EXISTS` ke tabel nilai custom field bertipe (JQL-050, modul 28).

### JQL-034 Platform & Business Fields Visible in JQL (Start date, Category, Goals)

- **Deskripsi:** Field tingkat platform atau proyek bisnis yang dapat dipakai di JQL.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Pengguna Tim Bisnis, Product Owner/PM.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ `"Start date"` dapat dicari dengan operator tanggal untuk timeline/kalender proyek bisnis dan software.
  2. ⚠ Field Category (kategori berwarna di proyek bisnis, CFG-021) dapat dicari sebagai field select dengan nama `category` khusus work item; konflik nama dengan project category diselesaikan dengan bentuk `"Category[Category]"`.
  3. ⚠ Field Goals (tujuan yang ditautkan ke work item, modul 18) dapat dicari dengan `=`, `IN`, `IS EMPTY` memakai ID goal.
  4. ⚠ Field proyek lintas-produk (eks layanan pembaruan status proyek) tidak dibangun sebagai field native; tercatat sebagai rujukan di Di luar cakupan.
  5. Component dari katalog komponen eksternal mengikuti operator `component` (JQL-027).
- **Team-managed vs company-managed:** Category umumnya hanya pada proyek bisnis.
- **Catatan & ketergantungan:** Seluruh perilaku perlu divalidasi terhadap modul 18 dan 26 saat implementasi.

### JQL-035 Development Information Fields (development[...])

- **Deskripsi:** Field untuk mencari work item berdasarkan data development dari integrasi SCM/CI/CD.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Scrum Master; data hanya tersedia untuk pengguna dengan View Development Tools (modul 03/24).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ `development[pullrequests].all > 0` dan `development[pullrequests].open > 0` didukung; status lain (merged, declined) ditambahkan bila modul 24 menyimpan status tersebut.
  2. `development[commits].all > 0` didukung.
  3. `development[builds].all`, `.failing`, dan `.passing` didukung.
  4. `development[reviews].all` dan `.open` didukung.
  5. ⚠ `development[branches].all` dan `development[deployments].all` didukung sesuai data yang disimpan modul 24.
  6. Sintaks legacy `issue.property[development].openprs > 0`, `.prs`, `.commits`, `.reviews` diterima sebagai alias deprecated dengan peringatan.
  7. Field hanya terisi bila integrasi (GitHub, GitLab, Bitbucket, CI/CD) terhubung; tanpa integrasi semua counter bernilai 0.
  8. Operator yang didukung: `=`, `!=`, `>`, `>=`, `<`, `<=`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Counter disimpan sebagai agregat per work item yang diperbarui event integrasi (modul 24, 28).

### JQL-036 Entity Property Search (issue.property[...])

- **Deskripsi:** Pencarian berdasarkan entity property work item yang diindeks oleh ekstensi atau integrasi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, pengembang ekstensi.
- **Aturan perilaku (acceptance criteria):**
  1. Sintaks `issue.property[propertyKey].path.to.value = nilai` diterima.
  2. Hanya property yang dideklarasikan untuk diindeks oleh ekstensi yang dapat dicari; property lain menghasilkan galat field tidak ada.
  3. Tipe index `number`, `string`, `text`, `date`, `user` menentukan operator yang valid.
  4. Operator yang didukung: `=`, `!=`, `>`, `<` (dan `>=`, `<=`), `~` (tipe text), `IN`, `IS EMPTY`.
  5. Alias pencarian (`searchAlias`) yang dideklarasikan ekstensi muncul sebagai nama field di autocomplete.
  6. Property di-set lewat REST API issue properties (modul 23); perubahan diindeks ulang secara asinkron.
  7. ⚠ Entity property selain work item (proyek, komentar, pengguna) tidak dapat dicari lewat JQL pada rilis awal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model ekstensi di modul 25. Indeks disimpan di tabel property ber-JSONB dengan index ekspresi per alias.

## Fungsi

### JQL-037 Function Arguments & Validation

- **Deskripsi:** Aturan umum argumen fungsi JQL dan validasinya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna pencarian lanjutan, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Jumlah argumen minimum/maksimum per fungsi divalidasi dengan galat "Function 'X' expected between N and M arguments but got K."
  2. Argumen harus di-quote bila mengandung spasi atau reserved character.
  3. ⚠ Argumen fungsi tidak boleh berupa pemanggilan fungsi lain; `membersOf(currentUser())` ditolak dengan galat sintaks.
  4. Fungsi list hanya valid dengan `IN`/`NOT IN` (dan `WAS IN`/`WAS NOT IN`); fungsi skalar hanya valid dengan `=`, `!=`, `>`, `>=`, `<`, `<=`, dan predicate; kombinasi lain menghasilkan galat operator.
  5. Fungsi tidak dikenal menghasilkan galat "Unable to find JQL function 'X'."
  6. Fungsi milik ekstensi yang telah di-uninstall menyebabkan saved filter yang memakainya gagal dengan galat fungsi tidak dikenal; filter tetap tersimpan dan dapat diedit.
  7. Nama fungsi case-insensitive (`CURRENTUSER()` setara `currentUser()`).
  8. Fungsi hanya valid untuk field tertentu (mis. `openSprints()` hanya untuk `sprint`); penggunaan pada field lain menghasilkan galat "Function 'openSprints' is not supported for field 'X'."
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Registri fungsi menyimpan signature (argumen, tipe kembalian, field yang kompatibel) dan dipakai validator, autocomplete, dan API.

### JQL-038 Date Functions (now, startOf*/endOf*)

- **Deskripsi:** Fungsi tanggal relatif terhadap waktu sekarang dan batas periode kalender.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `now()` mengembalikan waktu saat eksekusi.
  2. `startOfDay([inc])` / `endOfDay([inc])` mengembalikan awal/akhir hari.
  3. `startOfWeek([inc])` / `endOfWeek([inc])` mengembalikan awal/akhir minggu.
  4. `startOfMonth([inc])` / `endOfMonth([inc])` mengembalikan awal/akhir bulan.
  5. `startOfYear([inc])` / `endOfYear([inc])` mengembalikan awal/akhir tahun.
  6. Increment berformat `(+/-)nn(y|M|w|d|h|m)`, mis. `startOfMonth(-1)`, `startOfDay("+1d")`, `endOfWeek("-1w")`, `startOfYear("-1y")`.
  7. Increment tanpa unit memakai unit periode fungsi: `startOfMonth(-1)` adalah awal bulan lalu, `startOfWeek(1)` awal minggu depan.
  8. ⚠ Awal minggu mengikuti locale/preferensi pengguna (Senin atau Minggu, modul 21).
  9. Semua fungsi tanggal dievaluasi di timezone pengguna yang menjalankan kueri.
  10. Fungsi tanggal dapat dipakai dengan operator perbandingan dan di predicate history.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fungsi dievaluasi sekali per eksekusi menjadi literal timestamp sebelum translasi SQL (JQL-050).

### JQL-039 User & Group Functions (currentUser, membersOf)

- **Deskripsi:** Fungsi yang menghasilkan pengguna secara dinamis berdasarkan sesi atau keanggotaan grup.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; melihat anggota grup tidak memerlukan Browse Users and Groups untuk keperluan JQL (hasil tetap berupa work item).
- **Aturan perilaku (acceptance criteria):**
  1. `currentUser()` mengembalikan pengguna yang menjalankan kueri: untuk filter bersama yaitu penonton, untuk filter subscription yaitu penerima.
  2. ⚠ `membersOf("group")` mengembalikan semua anggota grup berdasarkan nama grup; ID grup juga diterima sebagai argumen.
  3. `membersOf()` dapat dipakai dengan `IN` pada assignee, reporter, creator, watcher, voter, worklogAuthor, dan custom user field.
  4. ⚠ `membersOf()` dapat dipakai di predicate `BY` pada history (JQL-016).
  5. Tidak ada grup bersarang; `membersOf()` hanya mengembalikan anggota langsung.
  6. Grup yang tidak ada menghasilkan galat "The value 'X' does not exist for the field 'membersOf'."
  7. Untuk pengguna anonim, `currentUser()` tidak mencocokkan apa pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Grup di USR-015; claims JWT `sess` di repo menyediakan identitas untuk `currentUser()`.

### JQL-040 Login Functions (lastLogin, currentLogin)

- **Deskripsi:** Fungsi tanggal berdasarkan sesi login pengguna.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna yang login.
- **Aturan perilaku (acceptance criteria):**
  1. `currentLogin()` mengembalikan waktu mulai sesi pengguna saat ini.
  2. `lastLogin()` mengembalikan waktu login sebelum sesi saat ini.
  3. Contoh `created > lastLogin()` mengembalikan work item baru sejak login terakhir.
  4. ⚠ Pada SSO atau sesi panjang, bantuan menjelaskan bahwa nilai bisa jauh di masa lalu; pengguna anonim mendapat galat fungsi memerlukan login.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data sesi dari AUTH-009.

### JQL-041 Personal Issue-set Functions (watchedIssues, votedIssues, issueHistory, updatedBy)

- **Deskripsi:** Fungsi yang mengembalikan work item berdasarkan aktivitas pengguna.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang login; `updatedBy` untuk pengguna lain tunduk pada visibilitas hasil biasa.
- **Aturan perilaku (acceptance criteria):**
  1. `issue in watchedIssues()` mengembalikan work item yang di-watch pengguna saat ini.
  2. `issue in votedIssues()` mengembalikan work item yang di-vote pengguna saat ini.
  3. ⚠ `issue in issueHistory()` mengembalikan work item yang baru dilihat pengguna saat ini, maksimal 50 terakhir.
  4. `issue in updatedBy(user[, dateFrom[, dateTo]])` mengembalikan work item yang diubah pengguna (perubahan field, komentar, worklog, transisi) dalam rentang waktu.
  5. Argumen tanggal `updatedBy` menerima tanggal absolut maupun relatif.
  6. `watchedIssues()`, `votedIssues()`, `issueHistory()` hanya untuk pengguna saat ini (tanpa argumen); `updatedBy()` dapat menerima pengguna lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Riwayat tampilan dari modul 21; aktivitas dari changelog modul 28.

### JQL-042 Sprint Functions (openSprints, closedSprints, futureSprints)

- **Deskripsi:** Fungsi yang mengembalikan daftar sprint berdasarkan state.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, Developer.
- **Aturan perilaku (acceptance criteria):**
  1. `openSprints()` mengembalikan sprint aktif (sudah dimulai, belum selesai) di semua board.
  2. `closedSprints()` mengembalikan sprint yang sudah diselesaikan.
  3. `futureSprints()` mengembalikan sprint yang dibuat tetapi belum dimulai.
  4. Fungsi dipakai dengan `sprint IN` / `sprint NOT IN`.
  5. Fungsi tidak menerima argumen; mempersempit ke board/proyek dilakukan dengan menggabung clause `project` atau `filter`.
  6. Fungsi berlaku untuk sprint di board company-managed maupun team-managed.
  7. Dengan parallel sprints, semua sprint aktif dikembalikan oleh `openSprints()`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** State sprint di SPR-012; parallel sprints di modul 11.

### JQL-043 Version Functions

- **Deskripsi:** Fungsi yang mengembalikan version berdasarkan status rilis.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Product Owner/PM, Admin Proyek; hasil terbatas pada proyek yang dapat di-browse.
- **Aturan perilaku (acceptance criteria):**
  1. `releasedVersions([project])` mengembalikan semua version yang sudah released.
  2. `unreleasedVersions([project])` mengembalikan semua version yang belum released.
  3. `latestReleasedVersion(project)` mengembalikan version released terakhir menurut sequence.
  4. `earliestUnreleasedVersion(project)` mengembalikan version unreleased paling awal menurut sequence.
  5. ⚠ `latestReleasedVersionByReleaseDate(project)` dan `earliestUnreleasedVersionByReleaseDate(project)` memakai release date sebagai urutan; version tanpa release date diabaikan.
  6. Argumen proyek berupa key, nama, atau ID; tanpa argumen berarti semua proyek yang dapat dilihat pengguna.
  7. ⚠ Version yang diarsip tidak ikut dihitung oleh fungsi-fungsi ini.
  8. Fungsi dipakai dengan `fixVersion`, `affectedVersion`, atau custom version field melalui `IN` / `NOT IN`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Release dan sequence version di modul 15.

### JQL-044 Issue Type Functions (standardIssueTypes, subtaskIssueTypes)

- **Deskripsi:** Fungsi yang mengembalikan kelompok work type.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `standardIssueTypes()` mengembalikan semua work type non-sub-task, termasuk Epic dan level di atasnya.
  2. `subtaskIssueTypes()` mengembalikan semua work type sub-task.
  3. Fungsi dipakai dengan `issuetype IN` / `NOT IN`.
  4. Hasil mencakup work type dari proyek team-managed dan company-managed.
  5. Bantuan editor menyebut alternatif modern `hierarchyLevel = -1` / `hierarchyLevel = 0` (JQL-017).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Work type di CFG-002/CFG-006.

### JQL-045 Project & Component Lead/Permission Functions

- **Deskripsi:** Fungsi yang mengembalikan proyek atau component berdasarkan peran atau permission pengguna.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, pengguna pencarian lanjutan.
- **Aturan perilaku (acceptance criteria):**
  1. `projectsLeadByUser([user])` mengembalikan proyek yang dipimpin pengguna (default pengguna saat ini, PRJ-020).
  2. `componentsLeadByUser([user])` mengembalikan component yang dipimpin pengguna (default pengguna saat ini).
  3. `projectsWhereUserHasPermission("Permission Name")` mengembalikan proyek tempat pengguna saat ini memiliki permission tertentu, mis. `"Resolve Issues"`.
  4. `projectsWhereUserHasRole("Role Name")` mengembalikan proyek tempat pengguna saat ini memiliki project role tertentu.
  5. Nama permission atau role yang tidak dikenal menghasilkan galat "The permission 'X' does not exist." / "The project role 'X' does not exist."
  6. ⚠ Di team-managed, role bawaan (Administrator, Member, Viewer) dan role kustom (PERM-039, PERM-040) dievaluasi oleh `projectsWhereUserHasRole()` dengan nama role tersebut.
- **Team-managed vs company-managed:** Lihat aturan 6.
- **Catatan & ketergantungan:** Evaluasi permission memakai layanan otorisasi modul 03 yang sama dengan penegakan RBAC, bukan logika terpisah.

### JQL-046 parentEpic() Function

- **Deskripsi:** Fungsi untuk mencari Epic beserta seluruh child dan sub-task-nya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Product Owner/PM, Scrum Master.
- **Aturan perilaku (acceptance criteria):**
  1. `issue in parentEpic("KEY-1")` mengembalikan Epic tersebut, child-nya (story/task/bug), dan sub-task dari child tersebut.
  2. ⚠ Bila argumen adalah story atau sub-task, Epic induknya dicari terlebih dahulu lalu dikembalikan Epic beserta seluruh keturunannya.
  3. Argumen berupa issue key atau ID.
  4. Bantuan editor menyebut alternatif satu level `parent in (...)`.
- **Team-managed vs company-managed:** Tidak ada perbedaan setelah unifikasi parent.
- **Catatan & ketergantungan:** Hierarki di CFG-006; untuk level di atas Epic gunakan `childIssuesOf()` (JQL-047).

### JQL-047 childIssuesOf() Function

- **Deskripsi:** Fungsi untuk menelusuri seluruh turunan sebuah work item secara rekursif di semua level hierarki, termasuk di atas Epic.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Perencana (Product Owner/PM), Scrum Master; fitur plan premium (VISI-003).
- **Aturan perilaku (acceptance criteria):**
  1. `issue in childIssuesOf("KEY-1")` mengembalikan semua turunan rekursif.
  2. Work item argumen itu sendiri tidak disertakan.
  3. Hierarki level kustom di atas Epic (Initiative, dsb., CFG-007) didukung.
  4. `portfolioChildIssuesOf()` diterima sebagai alias deprecated dari `childIssuesOf()` dengan peringatan.
  5. Fungsi dapat dipakai sebagai sumber work item di Plans dan board (modul 18, 10).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Traversal memakai CTE rekursif PostgreSQL dengan batas kedalaman dan ukuran hasil (JQL-052). Ketersediaan pada plan non-premium adalah keputusan terbuka.

### JQL-048 cascadeOption() Function

- **Deskripsi:** Fungsi untuk mencari custom field cascading select berdasarkan option parent dan child.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna pencarian lanjutan.
- **Aturan perilaku (acceptance criteria):**
  1. `cascadeOption(parentOption)` mencocokkan semua work item dengan option parent tersebut (child apa pun atau tanpa child).
  2. `cascadeOption(parentOption, childOption)` mencocokkan pasangan parent-child spesifik.
  3. `cascadeOption(parentOption, "none")` mencocokkan work item dengan parent tersebut tanpa child.
  4. Argumen menerima nama atau ID option.
  5. Fungsi dipakai dengan `IN` / `NOT IN` pada field cascading select; field lain menghasilkan galat fungsi tidak didukung untuk field.
- **Team-managed vs company-managed:** Cascading select umumnya hanya di company-managed.
- **Catatan & ketergantungan:** Field cascading di CFG-029.

### JQL-049 Custom JQL Functions (Forge jira:jqlFunction)

- **Deskripsi:** Titik ekstensi platform agar ekstensi mendefinisikan fungsi JQL baru.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengembang ekstensi, Admin Situs (memasang ekstensi).
- **Aturan perilaku (acceptance criteria):**
  1. Ekstensi mendaftarkan fungsi dengan nama, daftar argumen, tipe, dan operator yang didukung lewat manifest titik ekstensi `jqlFunction` produk ini.
  2. Saat eksekusi, fungsi diekspansi menjadi sub-JQL yang kemudian diparse dan dievaluasi dengan permission pengguna.
  3. Hasil fungsi dapat di-precompute; ekstensi menyimpan dan me-refresh hasil lewat REST `jql/function/computation` (GET/POST, JQL-060).
  4. Fungsi ekstensi muncul di autocomplete bersama fungsi bawaan dengan penanda sumber ekstensi.
  5. ⚠ Eksekusi fungsi dibatasi waktu (default 5 detik) dan ukuran sub-JQL; pelanggaran menghasilkan galat yang menyebut nama fungsi.
  6. ⚠ Mekanisme fungsi JQL gaya Connect (legacy) tidak didukung; hanya satu model ekstensi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model ekstensi di modul 25. Fungsi pihak ketiga seperti `issueFunction` adalah app Marketplace dan di luar cakupan.

## Evaluasi, keamanan & eksekusi

### JQL-050 JQL Compiler (JQL → SQL)

- **Deskripsi:** Komponen yang mem-parse JQL menjadi AST, memvalidasi terhadap field dan permission, lalu menerjemahkannya ke SQL terparameter yang aman dan memanfaatkan index.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Tim engineering; dipakai seluruh konsumen JQL.
- **Aturan perilaku (acceptance criteria):**
  1. Parser memakai grammar formal (PEG atau hand-written recursive descent) dan menghasilkan AST dengan posisi galat, mis. "Error in the JQL Query: ... (line x, character y)".
  2. Resolver field memetakan nama sistem, alias, nama custom field, dan `cf[10010]`; nama ambigu menghasilkan galat yang menawarkan pilihan ID.
  3. Translator SQL memakai parameter binding untuk setiap value (tanpa interpolasi string) sehingga tidak ada SQL injection; tes fuzz menyertakan payload injeksi.
  4. Field multi-value dan custom field bertipe diterjemahkan dengan subquery `EXISTS`/`NOT EXISTS` ke tabel relasi/nilai.
  5. Predicate history diterjemahkan ke subquery pada tabel changelog/status transition terindeks.
  6. Fungsi (`currentUser()`, `membersOf()`, `openSprints()`, `startOfDay()`, `issueHistory()`, `linkedIssues()`, dll.) dievaluasi sebagai pre-resolution menjadi set ID atau literal tanggal pada timezone pengguna sebelum SQL disusun.
  7. Value tidak dikenal (pengguna/version/proyek yang tidak ada atau tidak boleh dilihat) menghasilkan galat atau peringatan tanpa membocorkan keberadaannya.
  8. ORDER BY multi-kolom memakai tie-breaker stabil `id`; `ORDER BY Rank` memakai index rank.
  9. Pagination mendukung `startAt`/`maxResults` (legacy) dan `nextPageToken` (endpoint `/search/jql`).
  10. Approximate count disediakan terpisah dari pencarian; total eksak opsional dan dapat dimatikan untuk kueri besar.
  11. Query cost guard menerapkan `statement_timeout`, batas jumlah clause, dan batas ukuran set ID hasil fungsi (JQL-052).
  12. Telemetry kueri lambat mencatat teks kueri ter-redaksi (value diganti placeholder), durasi, dan rencana eksekusi untuk tuning index.
- **Team-managed vs company-managed:** Tidak ada perbedaan; resolver menangani field/status project-scoped (JQL-053).
- **Catatan & ketergantungan:** Usulan penempatan di repo: `pkg/jql` (lexer, parser, AST, tanpa dependensi DB) dan `internal/features/search` (resolver, registri fungsi, translator GORM/SQL, use case search). Filter permission (JQL-051) disisipkan oleh translator, bukan oleh pemanggil. NFR performa di modul 27, skema tabel di modul 28.

### JQL-051 Permission-aware JQL Evaluation

- **Deskripsi:** Aturan bagaimana permission dan status arsip memengaruhi evaluasi JQL.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna; Admin Situs mengatur permission scheme dan issue security.
- **Aturan perilaku (acceptance criteria):**
  1. Hasil hanya memuat work item di proyek tempat pengguna memiliki Browse Projects (PERM-020) dan yang lolos issue security level (PERM-044).
  2. Field atau value yang tidak dapat dilihat diperlakukan seolah tidak ada, dengan galat/peringatan generik yang sama dengan entitas yang benar-benar tidak ada.
  3. Clause `watcher`/`voter` untuk pengguna lain memerlukan View voters and watchers (PERM-033).
  4. Komentar dengan visibilitas terbatas hanya ikut dicocokkan `comment ~` / `text ~` untuk pengguna yang berhak (PERM-045).
  5. Shared filter dievaluasi dengan permission penonton, bukan pemilik filter.
  6. ⚠ Filter subscription dievaluasi per penerima, sehingga setiap penerima email hanya melihat work item yang boleh ia lihat.
  7. Akses anonim (proyek publik, PERM-046) memakai permission grant anonymous.
  8. ⚠ Work item yang diarsip (ISS-034) dan work item di proyek yang diarsip (PRJ-032) tidak muncul di hasil JQL.
- **Team-managed vs company-managed:** Team-managed memakai access level dan restriksi work item (PERM-038, PERM-041) yang dievaluasi oleh layanan otorisasi yang sama.
- **Catatan & ketergantungan:** Filter permission diterapkan di SQL (join ke set proyek yang dapat di-browse dan level keamanan pengguna), tidak dengan pasca-filter di memori, agar pagination dan count benar.

### JQL-052 JQL Search Execution Limits & Constraints

- **Deskripsi:** Batasan eksekusi kueri JQL yang ditiru produk ini demi stabilitas dan kompatibilitas API.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, Admin Situs (mengatur batas).
- **Aturan perilaku (acceptance criteria):**
  1. Endpoint `/rest/api/3/search/jql` menolak unbounded query (tanpa clause pembatas) dengan galat 400 yang meminta minimal satu restriksi.
  2. ⚠ Pagination berbasis `nextPageToken`; `maxResults` default 50, maksimum 100 bila field lengkap diminta dan maksimum 5.000 bila hanya `id`/`key` diminta.
  3. Endpoint `/rest/api/3/search/approximate-count` mengembalikan perkiraan jumlah hasil.
  4. Endpoint lama `/rest/api/3/search` (dengan `startAt`/`total`) disediakan sebagai kompatibilitas tetapi ditandai deprecated di dokumentasi dan header respons.
  5. Parameter `reconcileIssues` (daftar ID) memaksa hasil membaca perubahan terbaru untuk work item tersebut meski index belum sinkron.
  6. ⚠ Batas default dapat dikonfigurasi admin: panjang kueri 20.000 karakter, 300 clause, dan 1.000 item per list `IN`; pelanggaran menghasilkan galat yang menyebut batasnya.
  7. ⚠ Hasil fungsi dibatasi (mis. `issueHistory()` 50 item, `issuesWithRemoteLinksByGlobalId()` 100 argumen, fungsi set lain 10.000 ID).
  8. Kueri berat (history operator, text search luas) dibatasi `statement_timeout` (default 30 detik) dan gagal dengan galat "The JQL query took too long to execute" tanpa membebani koneksi lain.
  9. Hasil dan jumlah selalu difilter permission pengguna, sehingga jumlah total untuk kueri yang sama dapat berbeda antar pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kontrak endpoint search penuh di modul 23; angka batas final ditetapkan bersama modul 27.

### JQL-053 Team-managed vs Company-managed JQL Differences

- **Deskripsi:** Perbedaan perilaku JQL antara proyek team-managed dan company-managed.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, pengguna pencarian lanjutan.
- **Aturan perilaku (acceptance criteria):**
  1. Work type, status, dan custom field team-managed bersifat project-scoped dengan ID unik per proyek.
  2. Kueri dengan nama mencocokkan semua entitas bernama sama; autocomplete menampilkan varian `"Nama[Tipe]"` beserta nama proyek.
  3. `cf[id]` berbeda di tiap proyek team-managed untuk field bernama sama; kueri `cf[id]` hanya mencocokkan proyek pemilik field.
  4. Story point bernama `"Story point estimate"` di team-managed dan `"Story Points"` di company-managed; keduanya terpisah di JQL.
  5. ⚠ Field team-managed tidak memiliki konfigurasi searcher; semua tipe yang didukung otomatis dapat dicari.
  6. Environment, Affects versions, dan security level umumnya hanya ada di company-managed; clause atas field itu tidak mencocokkan work item team-managed.
  7. Resolution diset otomatis di team-managed saat masuk status category Done (JQL-020).
  8. Parent dan hierarki berperilaku sama di keduanya setelah unifikasi parent.
- **Team-managed vs company-managed:** Seluruh fitur ini membahas perbedaan tersebut.
- **Catatan & ketergantungan:** Konsep tipe proyek di PRJ-003; field team-managed di CFG-041/PRJ-051.

### JQL-054 JQL Consumers & Context Variables

- **Deskripsi:** Tempat JQL dipakai di produk ini dan perbedaan konteks evaluasinya (rincian UI di modul masing-masing).
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna, Admin, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Issue navigator dan saved filter dievaluasi dengan konteks pengguna penonton (modul 13).
  2. Board filter, sub-filter Kanban, quick filter, swimlane query, dan card color query memakai JQL; quick filter, swimlane, dan card color menolak ORDER BY (modul 10).
  3. Dashboard gadget (filter results, pie chart, created vs resolved, two-dimensional) memakai filter/JQL dengan konteks penonton dashboard (modul 17).
  4. ⚠ Filter subscription email dievaluasi dengan konteks penerima (modul 13, 19).
  5. Otomasi memakai JQL di condition, lookup issues, branch, dan scheduled trigger; smart value `{{...}}` disubstitusi sebelum parsing dan hasil substitusi di-escape (modul 20).
  6. Plans memakai filter, board, atau proyek sebagai sumber work item (modul 18).
  7. Webhook dapat dibatasi dengan filter JQL untuk event work item (modul 23).
  8. ⚠ `currentUser()` di otomasi mengacu ke actor rule (pengguna yang dikonfigurasi sebagai pelaku aturan).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semua konsumen wajib memanggil use case search yang sama (JQL-050) dengan konteks evaluasi eksplisit (pengguna, timezone, mode validasi).

## Editor, bantuan & konversi kueri

### JQL-055 JQL Autocomplete

- **Deskripsi:** Saran kontekstual saat mengetik JQL: field, operator, value, fungsi, keyword, dan predicate.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian; saran tunduk pada izin melihat entitas.
- **Aturan perilaku (acceptance criteria):**
  1. Saran nama field mencakup field sistem dan custom yang dapat dilihat pengguna, beserta tipenya untuk disambiguasi.
  2. Setelah field dipilih, hanya operator yang valid untuk field itu yang disarankan (JQL-013).
  3. Saran value mencakup proyek, status, pengguna (avatar + display name), version, component, sprint, label, option, dan tim.
  4. Saran fungsi ditampilkan lengkap dengan tanda kurung dan hanya fungsi yang kompatibel dengan field/operator saat ini.
  5. Keyword `AND`, `OR`, `NOT`, `ORDER BY`, `ASC`, `DESC`, `EMPTY` disarankan pada posisi yang valid.
  6. Setelah `WAS`/`CHANGED` disarankan predicate yang valid dan kemudian value predicate.
  7. ⚠ Permintaan saran di-debounce 150 ms dan dibatasi maksimal 15 saran per daftar.
  8. Navigasi keyboard: panah atas/bawah memindah sorotan, Enter/Tab memilih, Esc menutup daftar.
  9. Saran value memakai pencocokan prefix dan fuzzy pada nama.
  10. Saran hanya menampilkan entitas yang boleh dilihat pengguna.
  11. Saran field team-managed menampilkan nama + tipe/proyek.
- **Team-managed vs company-managed:** Lihat aturan 11.
- **Catatan & ketergantungan:** Data saran dari API autocomplete (JQL-060). Komponen frontend SvelteKit dibagikan (JQL-056).

### JQL-056 JQL Editor UX (syntax highlighting, validation, linting)

- **Deskripsi:** Perilaku editor JQL di aplikasi: highlight sintaks, validasi realtime, serta pesan galat dan peringatan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. Syntax highlighting membedakan field, operator, value, fungsi, dan keyword.
  2. Indikator validasi (valid/invalid) tampil dan posisi galat diberi garis bawah beserta tooltip pesan.
  3. Editor mendukung mode multi-baris dan tombol expand untuk kueri panjang.
  4. ⚠ Enter menjalankan kueri dan Shift+Enter menyisipkan baris baru.
  5. Tombol "Switch to basic" dinonaktifkan dengan tooltip penjelasan bila kueri terlalu kompleks untuk Basic search (JQL-058).
  6. Editor menyediakan tautan bantuan sintaks JQL yang terbuka di panel/tab baru.
  7. Komponen editor yang sama dipakai di navigator, board filter, otomasi, dan dashboard gadget (konsumen di modul lain).
  8. ⚠ Lint warning (mis. value tidak ada di sebagian proyek, field deprecated, equality pada datetime) ditampilkan tanpa memblokir eksekusi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validasi realtime memanggil API parse dengan mode `warn` (JQL-060); shortcut global di modul 21.

### JQL-057 JQL Error & Warning Messages

- **Deskripsi:** Katalog pesan galat dan peringatan standar dari parser dan validator.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Semua pengguna pencarian, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Galat sintaks operator: "Error in the JQL Query: Expecting operator but got 'X'. The valid operators are '=', '!=', '<', '>', '<=', '>=', '~', '!~', 'IN', 'NOT IN', 'IS', 'IS NOT', 'WAS', 'WAS IN', 'WAS NOT', 'WAS NOT IN' and 'CHANGED'."
  2. Galat sintaks operand dan kurung: "Expecting either a value, list or function but got 'X'." dan "Expecting ')' before the end of the query."
  3. Galat field: "Field 'X' does not exist or you do not have permission to view it."
  4. Galat value: "The value 'X' does not exist for the field 'Y'."
  5. Galat operator: "The operator 'X' is not supported by the 'Y' field."
  6. Galat fungsi: "Unable to find JQL function 'X'." dan galat jumlah argumen (JQL-037).
  7. Galat tanggal menampilkan daftar format valid (JQL-024).
  8. Galat key: "An issue with key 'X' does not exist for field 'key'."
  9. Galat sort: "Field 'X' does not support sorting."
  10. Galat reserved word/character menjelaskan perlunya quote (JQL-004).
  11. Mode validasi `strict`/`warn`/`none` (API parse dan parameter validasi di search) menentukan apakah value tidak dikenal menjadi galat, peringatan, atau diabaikan.
  12. Pesan dilokalisasi sesuai bahasa pengguna (EN/ID) memakai mekanisme `constants.GetError`/`GetErrorMessage` dan `LanguageMiddleware` yang sudah ada di repo; kode galat stabil disertakan untuk integrasi.
  13. Pesan tidak membocorkan keberadaan entitas yang tidak dapat diakses: teks untuk "tidak ada" dan "tidak boleh dilihat" identik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format respons API `errorMessages`/`warnings` di JQL-060. Pesan Inggris dipertahankan identik dengan JQL demi kompatibilitas skrip integrasi.

### JQL-058 Basic Search to JQL Round-trip

- **Deskripsi:** Konversi dua arah antara Basic search (dropdown) dan JQL; detail UI di modul 13.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna pencarian, terutama Pengguna Tim Bisnis.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap pilihan di Basic search menghasilkan clause JQL ekuivalen, mis. `project in (...) AND status in (...)`.
  2. JQL hanya dapat dibuka di Basic bila berupa rangkaian `AND` dari clause field `IN`/`=`/`~` sederhana yang masing-masing dapat diwakili kontrol Basic.
  3. `OR`, `NOT`, kurung, sebagian fungsi, dan operator history membuat kueri ditandai "too complex for basic" dan editor tetap di mode JQL.
  4. Kotak teks Basic menghasilkan clause `text ~ "..."`.
  5. Kontrol "More" di Basic menambah clause untuk field lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi dilakukan atas AST (JQL-001), bukan manipulasi string.

### JQL-059 Natural Language to JQL (AI search)

- **Deskripsi:** Fitur AI opsional yang mengubah kalimat bahasa alami menjadi JQL di pencarian work item.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna pencarian; Admin Situs mengaktifkan/menonaktifkan fitur AI.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat mengetik kalimat bahasa alami (mis. "bug saya yang belum selesai minggu ini") di pencarian work item.
  2. Sistem menghasilkan JQL yang ditampilkan dan dapat diedit sebelum maupun sesudah dijalankan.
  3. Generasi memakai konteks field, proyek, dan pengguna yang dapat diakses penanya; nama entitas yang tidak boleh dilihat tidak dikirim ke model.
  4. ⚠ Input bahasa Indonesia dan Inggris didukung pada rilis awal.
  5. Pengguna dapat memberi feedback (jempol naik/turun) atas hasil, tersimpan untuk evaluasi kualitas.
  6. Admin dapat menonaktifkan fitur AI di tingkat organisasi; saat nonaktif, input bahasa alami tidak ditampilkan.
  7. ⚠ Ketersediaan per plan diatur feature flag tier (VISI-003).
  8. Bila generasi gagal atau JQL hasil tidak lolos parser (JQL-050), pengguna diarahkan ke JQL manual dengan pesan penjelasan.
  9. ⚠ Penggunaan di konteks lain (otomasi, board filter) tidak tersedia pada rilis awal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan: JQL-059 adalah pemilik terjemahan bahasa alami ke JQL (generasi, konteks izin, validasi parser, feedback, kebijakan admin); UX input bahasa alami di issue navigator (penempatan kotak, peralihan ke mode JQL, tampilan hasil) dimiliki [13 SRC-019](13-pencarian-filter-ekspor.md), yang memanggil layanan terjemahan ini. Implementasi dengan LLM yang dapat diganti penyedianya + validasi wajib via parser sendiri; kebijakan AI di modul 27. Layanan AI merek lain tidak dipakai sebagai nama fitur.

## API JQL

### JQL-060 JQL REST APIs (autocomplete, parse, sanitize, pdcleaner, match, function computation)

- **Deskripsi:** Endpoint REST untuk parsing, validasi, autocomplete, dan transformasi JQL.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, pengembang ekstensi; autentikasi JWT/API token, hasil tunduk pada izin pemanggil.
- **Aturan perilaku (acceptance criteria):**
  1. `GET`/`POST /rest/api/3/jql/autocompletedata` mengembalikan `visibleFieldNames` (value, displayName, operators, types, cfid, searchable, orderable, auto), `visibleFunctionNames`, dan `jqlReservedWords`; versi POST dapat difilter `projectIds`.
  2. `GET /rest/api/3/jql/autocompletedata/suggestions?fieldName&fieldValue&predicateName&predicateValue` mengembalikan saran value.
  3. `POST /rest/api/3/jql/parse?validation=strict|warn|none` mengembalikan AST (`structure.where`, `orderBy`) dan galat per kueri, untuk beberapa kueri sekaligus.
  4. `POST /rest/api/3/jql/sanitize` mengganti referensi yang tidak boleh dilihat pengguna tertentu dengan ID.
  5. `POST /rest/api/3/jql/pdcleaner` mengonversi username/userkey legacy menjadi accountId.
  6. `POST /rest/api/3/jql/match` memeriksa apakah daftar ID work item cocok dengan beberapa JQL dan mengembalikan pasangan yang cocok.
  7. `GET`/`POST /rest/api/3/jql/function/computation` menyediakan precomputation untuk custom function ekstensi (JQL-049).
  8. Evaluasi ekspresi (`/rest/api/3/expression/eval`) dengan konteks `jql` mengembalikan hasil JQL untuk dihitung ekspresinya.
  9. Respons galat mengikuti format `errorMessages` (array) dan `warnings` (array) dengan status HTTP 400 untuk kueri tidak valid.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Endpoint `/search/jql` dan kontrak umum API di modul 23; anotasi Swagger wajib (`make swag`) dan `@Security BearerAuth` sesuai konvensi repo. Path dipertahankan kompatibel; versi internal repo tetap di bawah `/api/v1` dengan lapisan kompatibilitas di modul 23.

## Fitur lama (rujukan)

### JQL-061 Quick Search Smart Querying (legacy)

- **Deskripsi:** Kata kunci pintar di quick search lama (mis. `my`, `overdue`, `r:me`, `created:today`) yang diterjemahkan menjadi JQL; fitur Data Center/lama yang tidak dibangun, dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Tidak ada (tidak dibangun).
- **Aturan perilaku (acceptance criteria):**
  1. Keyword `my` (`assignee = currentUser()`), `r:me` / `r:<user>`, `overdue`, `open`, `resolved`, serta `created:`/`updated:`/`due:` dengan rentang (`today`, `yesterday`, `-1w`) tidak didukung; quick search produk ini (modul 13) memperlakukannya sebagai teks biasa.
  2. Pengenalan otomatis nama proyek/work type/status/component di teks quick search tidak dibangun.
  3. Teks sisa tidak diterjemahkan otomatis ke `text ~`; pengguna dapat membuka navigator dengan JQL manual.
  4. Tidak ada pembukaan hasil quick search sebagai JQL di navigator berbasis smart query; kebutuhan serupa dipenuhi JQL-059 (bahasa alami) bila diaktifkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur yang sama juga dicatat sebagai rujukan di [13 SRC-013](13-pencarian-filter-ekspor.md) (sisi pencarian/navigator) dan [21 NAV-055](21-navigasi-ux-preferensi.md) (sisi kotak quick search); ketiganya Fase — dan tidak dibangun, JQL-061 hanya mencatat dampaknya pada bahasa kueri. Alasan tidak dibangun: smart querying tidak lagi didukung di quick search versi cloud saat ini dan digantikan pencarian berbasis teks serta AI (JQL-059, SRC-019). Lihat Di luar cakupan.

## Pertanyaan terbuka

- Apakah produk ini menambahkan alias berbasis istilah "work item" (mis. `workitemkey`, `workItemType`) di samping nama JQL klasik, atau tetap hanya memakai nama klasik demi kompatibilitas penuh (JQL-001)?
- Apakah stemming dan stop word bahasa Indonesia ditambahkan di samping bahasa Inggris untuk operator `~`, dan bagaimana memilih konfigurasi per situs atau per dokumen (JQL-012)?
- Apakah pencocokan `labels` case-sensitive (seperti penyimpanan apa adanya) atau dinormalisasi case-insensitive (JQL-022)?
- Berapa angka final batas eksekusi (panjang kueri, jumlah clause, item `IN`, timeout, `maxResults` untuk id/key) yang ditetapkan untuk situs self-hosted vs cloud produk ini (JQL-009, JQL-052)?
- Apakah `childIssuesOf()` tersedia di semua plan atau hanya plan premium (JQL-047)?
- Apakah nama field terlokalisasi (bahasa Indonesia) diterima sebagai alias kueri, dan bagaimana dampaknya pada saved filter yang dibagikan lintas bahasa (JQL-005)?
- Apakah penambahan reserved word baru di masa depan diperbolehkan, mengingat dapat merusak saved filter yang ada (JQL-004)?
- Apakah fitur AI bahasa-alami ke JQL dikirim sebagai bagian inti atau modul opsional dengan penyedia LLM yang dapat dikonfigurasi admin (JQL-059)?

## Di luar cakupan & ekosistem

- **Fungsi JQL app Marketplace pihak ketiga** (mis. `issueFunction in ...` dari ScriptRunner, fungsi JQL dari app pelaporan/komentar lanjutan, pencarian penulis/tanggal komentar). Produk ini hanya menyediakan titik ekstensi JQL-049; fungsi pihak ketiga tidak dibangun sebagai fitur native.
- **Quick Search Smart Querying (legacy)** (JQL-061) — fitur Data Center/lama, ditandai dc-only-or-deprecated; hanya dicatat sebagai rujukan.
- **Mekanisme fungsi JQL gaya Connect (legacy)** — hanya satu model ekstensi yang didukung (JQL-049).
- **Sintaks Lucene lama di text search** (fuzzy, proximity, boosting, regex) — tidak didukung (JQL-012).
- **Layanan lintas produk Atlassian**: field proyek eks layanan pembaruan status proyek, Goals/Teams sebagai layanan terpusat lintas produk, dan AI bermerek Atlassian Intelligence/Rovo. Produk ini menyediakan padanan internal (JQL-032, JQL-034, JQL-059) tanpa integrasi ke layanan tersebut.
- **Field & fungsi JQL khusus Service Management** (SLA `breached()`, `withinCalendarHours()`, approval `approver()`/`pendingBy()`, organization, request channel) dan **AQL/`aqlFunction()` Assets** — tidak didefinisikan di modul ini; dibahas sebagai perluasan bahasa di modul 31 dan 33 dengan memakai registri field/fungsi JQL-037 dan JQL-050.
- **Rincian UI issue navigator, Basic search, saved filter, subscription, dan ekspor** — di modul 13; modul ini hanya mendefinisikan bahasa dan evaluasinya.
