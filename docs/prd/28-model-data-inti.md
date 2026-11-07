# 28 — Model Data Inti & Katalog Event Domain

Modul ini menetapkan **model data kanonik** produk ini (clone) beserta **katalog event domain** yang menjadi kontrak bersama seluruh modul fungsional. Cakupannya adalah tabel, kunci, constraint, index, aturan konsistensi transaksi, dan job latar yang wajib ada agar perilaku fungsional di modul lain (project, work item, workflow, board, sprint, JQL, notifikasi, otomasi, webhook, API) dapat dibangun tanpa saling bertabrakan: identitas dan kredensial, scheme otorisasi/notifikasi/keamanan, project dan issue key, konfigurasi issue type/screen/field, penyimpanan custom field, workflow, konten sekunder (comment, worklog, attachment, link), watcher/vote, versi/komponen/label, board/sprint/rank (LexoRank), filter/dashboard, entity properties, automation rule, graf entitas lintas produk, dan event issue. Modul ini **tidak** mendefinisikan UI; setiap fitur di sini adalah fitur platform (tier `platform`) yang dikonsumsi modul lain. Istilah skema dan API tetap memakai `issue` walaupun UI memakai "work item" (lihat DM-003 butir 13).

Kondisi repo saat ini: hanya ada tabel `users` (migrasi `migrations/001_create_users_table.up.sql`: `id VARCHAR(36)` UUID, `email` unik, `username` unik, `password` bcrypt, `name`, `is_active`, `created_at`/`updated_at`/`deleted_at` soft delete GORM) dan entity `internal/shared/domain/entity/user.go` (kolom tambahan `phone`, `status`, `role`, `provider`, dll). Token JWT dari `pkg/jwt` berlaku 24 jam tanpa refresh maupun revocation. Seluruh tabel lain di modul ini adalah pekerjaan baru yang ditambahkan lewat migrasi golang-migrate bernomor (`make migration-create name=...`), entity GORM di `internal/shared/domain/entity/`, interface repository di `internal/shared/domain/repository/`, dan implementasi di `internal/shared/infrastructure/repository/`, mengikuti Clean Architecture repo. Konvensi umum: id numerik `BIGINT` (identity) untuk entitas yang ID-nya terekspos di API kompatibel Jira (project, issue, comment, worklog, version, component, board, sprint, filter, dashboard, custom field), timestamp `timestamptz` UTC, dan JSON terstruktur sebagai `JSONB`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah work item/issue, tier plan & feature flag.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — perilaku login, sesi, API token, OAuth consent (AUTH-001, AUTH-007, AUTH-008, AUTH-009, AUTH-012).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — deaktivasi user, grup, product access (USR-007, USR-008, USR-015, USR-016, USR-018).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission scheme, issue security, evaluasi server-side (PERM-002, PERM-014, PERM-042, PERM-044).
- [04 Proyek](04-proyek.md) — aturan project key, ganti key, trash/arsip (PRJ-012, PRJ-022, PRJ-033).
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — issue type, custom field, context, screen, field configuration (CFG-001 s.d. CFG-048).
- [06 Work Item Inti](06-work-item-inti.md) — key & numbering, hierarki, move, watcher, vote, label (ISS-001, ISS-025, ISS-032, ISS-036, ISS-037).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — perilaku comment/attachment, changelog (CMT-019, CMT-023).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — editor workflow, scheme, migrasi status (WF-021, WF-033, WF-036).
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — konfigurasi board, sprint, rank (BRD-017, SPR-019, SPR-023).
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — compiler JQL→SQL, indexing, saved filter (JQL-050, SRC-002, SRC-036).
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md), [15 Rilis & Komponen](15-rilis-versi-komponen.md) — worklog, version, component.
- [17 Dashboard & Gadget](17-dashboard-gadget.md), [19 Notifikasi & Email](19-notifikasi-email.md), [20 Otomasi](20-otomasi.md) — konsumen data dashboard, event, dan automation rule.
- [22 Administrasi Sistem](22-administrasi-sistem.md), [23 REST API & Webhook](23-api-publik-webhook.md), [24 Integrasi DevOps](24-integrasi-devops.md), [25 Integrasi & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — data lifecycle, API, dev info, ekstensibilitas.
- [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) — caching, guardrail, migrasi database, konkurensi, background job & outbox, real-time, audit log (NFR-003 s.d. NFR-009, NFR-017).
- [29 Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) s.d. [37 Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) — modul Service Management & Product Discovery yang menambah tabel di atas fondasi ini.

## Persona terkait

- **Admin Situs** — mengandalkan konsistensi scheme, deaktivasi user, trash custom field, dan job migrasi yang terlihat progresnya.
- **Admin Proyek/Project Lead** — mengganti project key, scheme, dan konfigurasi team-managed tanpa merusak data historis.
- **Developer/Anggota Tim** — merasakan hasil model data lewat key issue yang stabil, rank drag-and-drop cepat, dan riwayat yang akurat.
- **Scrum Master/Team Lead** — bergantung pada histori sprint dan scope change yang tercatat benar untuk laporan.
- **Integrator/Developer API** — memakai id numerik kompatibel, entity properties, event/webhook yang idempoten, dan custom field `customfield_10xxx`.
- **Agent Layanan** — memakai flag comment internal yang dimodelkan sejak awal agar Service Management (Fase 5) tidak memerlukan migrasi besar.

## User story modul

1. Sebagai Developer/Anggota Tim, saya ingin issue key per project yang unik dan tidak pernah dipakai ulang agar tautan dan referensi di commit tetap menunjuk work item yang benar (DM-003).
2. Sebagai Admin Proyek/Project Lead, saya ingin mengganti project key tanpa memutus URL lama agar bookmark dan referensi eksternal tetap berfungsi (DM-003).
3. Sebagai Admin Situs, saya ingin penggantian workflow scheme atau issue type scheme berjalan sebagai job dengan progres agar perubahan konfigurasi besar tidak membuat situs macet (DM-004, DM-006).
4. Sebagai Integrator/Developer API, saya ingin custom field memiliki id `customfield_10xxx` dan dapat dicari dengan JQL secara cepat agar integrasi yang sudah ada untuk Jira dapat dipakai ulang (DM-005).
5. Sebagai Admin Situs, saya ingin custom field yang terhapus masuk Trash 60 hari agar kesalahan hapus bisa dipulihkan beserta nilainya (DM-005).
6. Sebagai Developer/Anggota Tim, saya ingin transisi workflow yang atomik agar work item tidak pernah berhenti di status setengah jadi saat validator atau post function gagal (DM-006).
7. Sebagai Scrum Master/Team Lead, saya ingin histori sprint dan scope change disimpan per event agar Sprint report, Burndown, dan Velocity akurat (DM-010).
8. Sebagai Developer/Anggota Tim, saya ingin drag-and-drop ranking yang hanya mengubah satu baris agar backlog besar tetap responsif (DM-010).
9. Sebagai Admin Situs, saya ingin cache izin yang diinvalidasi saat grup atau role berubah agar pengguna yang dicabut aksesnya langsung tidak melihat work item rahasia (DM-002).
10. Sebagai Admin Situs, saya ingin menonaktifkan user tanpa menghapus riwayatnya dan memindahkan kepemilikan filter/dashboard-nya agar audit dan board tim tetap utuh (DM-001, DM-011).
11. Sebagai Integrator/Developer API, saya ingin menyimpan entity property JSON pada work item agar app saya bisa menyimpan metadata tanpa custom field (DM-013, DM-012).
12. Sebagai Integrator/Developer API, saya ingin event dengan id unik dan sequence per issue agar webhook consumer dapat menolak duplikat dan mengurutkan perubahan (DM-014).
13. Sebagai Developer/Anggota Tim, saya ingin melihat konten terkait work item (PR, page, goal) tanpa bocornya objek yang tidak boleh saya lihat (DM-015).
14. Sebagai Admin Proyek/Project Lead, saya ingin watcher dan vote tercatat per user agar notifikasi dan JQL `watcher`/`voter` akurat (DM-008).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 10 |
| P1 | 3 |
| P2 | 0 |
| P3 | 2 |
| **Total** | **15** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 10 |
| Fase 2 | 2 |
| Fase 3 | 1 |
| Fase 4 | 2 |
| Fase — | 0 |

Strategi urutan implementasi: model data adalah prasyarat semua modul, sehingga seluruh fitur P0 masuk Fase 1, tetapi dibangun **bertahap per tabel** mengikuti modul konsumen: identitas (DM-001) lebih dulu karena memigrasi tabel `users` yang sudah ada, lalu project & issue key (DM-003), konfigurasi issue type/status default (DM-004), workflow (DM-006), event outbox (DM-014), konten sekunder (DM-007), versi/komponen/label (DM-009), dan board/sprint/rank (DM-010). Bagian P0 yang perilakunya baru aktif di fase berikut (custom field context lengkap, field configuration scheme, permission scheme penuh, issue security) tetap dibuat tabelnya di Fase 1 dengan satu scheme default agar Fase 2-3 tidak memerlukan migrasi data yang merusak; butir yang demikian diberi keterangan fase di aturannya. Pengecualiannya, tabel `filters` dan share filter dari DM-011 dibangun di Fase 1 bersama board/sprint (DM-010) karena board company-managed selalu berbasis saved filter (BRD-010). Watcher/vote (DM-008) serta sisa DM-011 (subscription, favourites, dashboard & gadget) menyusul di Fase 2 bersama JQL dan navigator; entity properties & automation (DM-012) di Fase 3 bersama mesin otomasi dan REST API publik; entity properties work item via REST (DM-013) dan graf entitas (DM-015) di Fase 4 bersama ekstensi dan AI opsional. Setiap migrasi wajib memiliki pasangan `down.sql` dan uji repository dengan go-sqlmock.

## Fitur

## Identitas & otorisasi

### DM-001 Data Model: Users, Groups, Roles, Sessions & Tokens

- **Deskripsi:** Model identitas pengguna, keanggotaan grup, product access, serta kredensial (sesi dengan refresh token, API token, OAuth 2.0) yang menggantikan sekaligus kompatibel dengan auth JWT repo saat ini.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering (implementasi), Admin Situs (deaktivasi, product access, default groups); pengguna akhir mengelola sesi, API token, dan preferensinya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `users` memiliki `id` internal (tetap UUID `VARCHAR(36)` repo saat ini) dan `account_id` opaque yang unik, immutable, dan dipakai di semua API/payload publik (bukan `id` internal, bukan username); `email` bertipe `citext` unik **hanya untuk akun aktif** (partial unique index `WHERE status = 'active'`), serta `display_name`, `public_name`, `avatar_key`, `timezone` (IANA), `locale` (mis. `en`, `id`), `status` (`active`/`deactivated`), dan `account_type` (`atlassian`/`app`/`customer`); payload API tidak pernah memuat kolom `password`.
  2. Tabel `groups` (id, `name` unik case-insensitive, description, `managed_by` system/admin) dan `group_members` (group_id, user_id unik); relasi grup bersarang (group di dalam group) **tidak** didukung: API menolak penambahan grup sebagai anggota grup dengan 400.
  3. Saat situs dibuat, seed membuat grup default (`<product>-users-<site>`, `<product>-admins-<site>`, `site-admins`, `org-admins`) dan tabel `product_access` (group_id, product, role `user`/`admin`); user baru yang ditambahkan ke grup default otomatis mendapat product role User, dan grup yang sedang memberi product access tidak dapat dihapus (409).
  4. Tabel `user_sessions` menyimpan `refresh_token_hash` (SHA-256, token mentah tidak pernah disimpan), `jti` access token, device, IP, user agent, `created_at`, `last_used_at`, `expires_at`, `revoked_at`; refresh token dirotasi setiap dipakai dan pemakaian ulang refresh token lama mencabut seluruh rangkaian sesi tersebut.
  5. Tabel `api_tokens` menyimpan `token_hash`, `label`, `created_at`, `last_used_at`, `expires_at` (wajib, 1-365 hari untuk token baru; nilai di luar rentang ditolak 400), dan `scopes` (array; kosong berarti classic token dengan izin penuh user); token kedaluwarsa atau dicabut ditolak 401.
  6. Tabel `oauth_clients` (client_id, secret hash, redirect URI terdaftar, scopes, owner) dan `oauth_grants` (client, user, scopes, authorization code hash, refresh token hash, expires, revoked) mendukung OAuth 2.0 (3LO); pencabutan consent oleh user menghapus semua grant aktif client tersebut untuk user itu.
  7. Deaktivasi user mengubah `status` menjadi `deactivated`, mencabut semua `user_sessions`, `api_tokens`, dan `oauth_grants`, tidak menghitung lisensi, dan mempertahankan seluruh data historis (reporter, comment, changelog tetap menunjuk user); hapus permanen meng-anonimkan profil sehingga UI/API menampilkan "Former user" dengan `account_id` tetap sama.
  8. Tabel `user_properties` (user_id, key, value JSONB) dan kolom `preferences` JSONB menyimpan preferensi notifikasi, bahasa, timezone, theme, dan keyboard shortcuts on/off; nilai yang tidak lolos JSON schema preferensi ditolak 400.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku fungsional di AUTH-001, AUTH-007, AUTH-008, AUTH-009, AUTH-012, AUTH-013 ([01](01-identitas-autentikasi-sesi.md)) dan USR-007, USR-008, USR-015, USR-016, USR-018 ([02](02-pengguna-grup-organisasi.md)). Migrasi repo: tabel `users` sekarang memakai `username` unik dan `name`; migrasi menambah `account_id`, `display_name` (diisi dari `name`), `status`, `timezone`, `locale`, `account_type`, mengubah `email` menjadi `citext` dengan partial unique index, dan mempertahankan `username` hanya sebagai alias login legacy (lihat Pertanyaan terbuka). Kolom `role` di entity `User` digantikan oleh grup + product access + global permission (DM-002). `pkg/jwt` perlu klaim `jti` dan pemeriksaan revocation untuk mendukung butir 4.

### DM-002 Data Model: Permission, Notification & Issue Security Schemes

- **Deskripsi:** Skema data otorisasi tingkat project (permission scheme), global permission, project role, notification scheme, dan issue security scheme, beserta cache predikat yang membuat evaluasi izin efisien di setiap query.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs (mengelola scheme, global permission, role global); Admin Proyek (role actors per project, akses & role team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `permission_schemes` dan `permission_grants` (scheme_id, `permission_key`, `holder_type`, `holder_param`) mendukung holder type: `user`, `group`, `project_role`, `reporter`, `assignee`, `project_lead`, `any_logged_in`, `anonymous`, `user_custom_field`, `group_custom_field`, `application_access`; kombinasi (scheme, permission, holder_type, holder_param) unik dan holder type di luar daftar ditolak 400. Fase 1 cukup satu Default Permission Scheme hasil seed; editor penuh di Fase 3.
  2. Tabel `global_permission_grants` (permission_key, group_id) mencakup minimal Administer, Browse users and groups, Share dashboards and filters, Manage group filter subscriptions, Make bulk changes, dan Create team-managed projects; grant Administer terakhir tidak dapat dihapus (409) agar situs tidak terkunci.
  3. Tabel `project_roles` global (id, name unik, description; seed Administrators, Member, Viewer) dan `project_role_actors` (project_id, role_id, actor_type `user`/`group`, actor_id) unik per kombinasi; menghapus role yang masih dipakai grant scheme memerlukan pengganti atau ditolak 409.
  4. Tabel `notification_schemes` dan `notification_scheme_entries` (scheme_id, event_type_id, recipient_type, recipient_param) mendukung recipient: `current_assignee`, `reporter`, `current_user`, `project_lead`, `component_lead`, `watchers`, `user`, `group`, `project_role`, `user_custom_field`, `group_custom_field`; event_type_id mengacu katalog event DM-014.
  5. Tabel `issue_security_schemes`, `security_levels` (scheme_id, name unik per scheme, description, `is_default`), dan `security_level_members` (level_id, holder_type, holder_param); kolom `issues.security_level_id` nullable dengan foreign key; maksimal satu level default per scheme, dan level yang masih dipakai issue tidak bisa dihapus tanpa memilih level pengganti.
  6. Sistem mem-precompute dan meng-cache per user dua himpunan: project yang dapat di-Browse dan security level yang dimiliki; himpunan ini disisipkan sebagai predikat SQL (`project_id = ANY(...)` dan `security_level_id IS NULL OR security_level_id = ANY(...)`) di setiap query pencarian, board, dan API list sehingga issue tanpa izin tidak pernah ikut terhitung (termasuk di count dan agregat).
  7. Perubahan `group_members`, `project_role_actors`, `permission_grants`, `security_level_members`, atau asosiasi scheme ke project menginvalidasi cache izin user terdampak dalam transaksi yang sama (atau via event outbox yang diproses < 5 detik); pengujian memastikan user yang dikeluarkan dari grup langsung mendapat 404 pada issue rahasia di request berikutnya.
  8. Project team-managed menyimpan `access_level` (`open`/`limited`/`private`) di tabel `projects` dan role kustom per project di tabel `project_scoped_roles` + `project_scoped_role_permissions`, terpisah dari permission scheme; evaluator izin memilih jalur scheme atau jalur team-managed berdasarkan `projects.style`.
- **Team-managed vs company-managed:** Company-managed memakai permission/notification/issue security scheme bersama (butir 1-5); team-managed memakai access level dan role kustom per project (butir 8) serta restriksi issue sederhana (PERM-041) yang dipetakan ke tabel security level ber-scope project.
- **Catatan & ketergantungan:** Perilaku fungsional dan evaluasi server-side di PERM-002, PERM-014, PERM-018, PERM-038, PERM-040, PERM-042, PERM-043, PERM-044 ([03](03-izin-peran-keamanan.md)); notification scheme di NTF-004 ([19](19-notifikasi-email.md)). Satu project company-managed menunjuk tepat satu scheme per jenis (DM-004 butir 9). Kegagalan izin pada issue yang tidak boleh dilihat menghasilkan 404, bukan 403, agar keberadaan issue tidak bocor.

## Project, work item & konfigurasi

### DM-003 Core Data Model: Projects & Issue Keys

- **Deskripsi:** Model data project dan issue (work item), termasuk sequence issue key per project yang unik, monoton, aman terhadap konkurensi, dan tetap resolve setelah rename key atau move issue.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs/Admin Proyek (membuat project, rename key); semua pengguna dengan Browse projects (resolve key).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `projects` berisi `id` BIGINT, `key`, `name`, `type` (`software`/`business`; `service_desk` dan `product_discovery` hanya tersedia bila entitlement Service Management/Product Discovery aktif, selain itu ditolak 400), `style` (`team_managed`/`company_managed`, immutable setelah dibuat), `lead_account_id`, `default_assignee` (`project_lead`/`unassigned`), `avatar_key`, `category_id`, `url`, `description`, `access_level` (hanya untuk team-managed: `open`/`limited`/`private`), dan `state` (`active`/`archived`/`trashed`).
  2. Validasi project key: diawali huruf besar A-Z, hanya berisi huruf besar, angka, dan underscore, panjang 2-10 karakter, unik case-insensitive terhadap key aktif **dan** key lama di tabel riwayat; pelanggaran menghasilkan 400 dengan pesan galat EN/ID yang menyebut aturan yang dilanggar (mis. "Project key must start with an uppercase letter").
  3. Tabel `project_key_history` (project_id, old_key, changed_at, changed_by) diisi saat rename key; lookup `OLDKEY-123` me-resolve ke issue yang sama (UI redirect 301 ke key baru, API mengembalikan issue dengan key terkini) dan pencarian teks dengan key lama tetap menemukan issue.
  4. Kolom `projects.next_issue_number` dialokasikan dengan `UPDATE projects SET next_issue_number = next_issue_number + 1 WHERE id = $1 RETURNING next_issue_number - 1` di dalam transaksi pembuatan issue; 100 create paralel pada project yang sama menghasilkan 100 nomor berbeda tanpa celah dari proses yang sukses, dan nomor tidak pernah dipakai ulang walaupun issue dihapus atau dipindah.
  5. Move issue antar project mengalokasikan nomor baru di project tujuan, menyimpan key lama di `moved_issue_keys` (old_key unik, issue_id), mencatat perubahan key di changelog (field `Key` dan `project`), dan request ke key lama di-redirect/resolve ke issue di project baru.
  6. Lookup issue key bersifat case-insensitive: `abc-1`, `Abc-1`, dan `ABC-1` me-resolve issue yang sama di URL, API (`/rest/api/3/issue/{issueIdOrKey}`), dan JQL `key = abc-1`; key yang tidak ditemukan atau tidak boleh dilihat mengembalikan 404.
  7. Tabel `issues` berisi `id` BIGINT global, `project_id`, `issue_number`, `key` (kolom terkomputasi/terdenormalisasi dari key project aktif + nomor), `issuetype_id`, `status_id`, `resolution_id`, `priority_id`, `summary` (wajib, maks 255 karakter, ditolak 400 bila kosong/whitespace atau lebih panjang), `description` (ADF JSON), `environment`, `reporter_id`, `assignee_id`, `creator_id`, `created_at`, `updated_at`, `resolution_date`, `due_date`, `start_date`, `security_level_id`, `parent_id`, `rank`, `votes_count`, `watches_count`, `time_original_estimate`, `time_estimate`, `time_spent` (detik), dan `archived` (+`archived_at`, `archived_by`).
  8. Unique constraint `(project_id, issue_number)` ditegakkan di database; pelanggaran (mis. akibat bug alokasi) menggagalkan transaksi create, bukan membuat duplikat.
  9. Hierarki memakai `parent_id` generik dan `issue_types.hierarchy_level` (Epic = 1, standard = 0, sub-task = -1; level > 1 hanya bila fitur Plans Premium aktif): parent harus tepat satu level di atas child, sub-task wajib punya parent standard, dan penetapan parent yang membentuk siklus atau level tidak cocok ditolak 400.
  10. Tabel konfigurasi (issue types, custom fields, workflows, statuses, boards) memiliki kolom `project_id` nullable: `NULL` berarti global/company-managed, terisi berarti lokal milik satu project team-managed; konfigurasi lokal tidak pernah muncul di project lain maupun di halaman admin global.
  11. Project company-managed berbagi konfigurasi lewat scheme global (DM-004 butir 9); mengubah scheme bersama berdampak ke semua project yang memakainya dan UI admin menampilkan jumlah project terdampak.
  12. Kolom denormalisasi untuk list (mis. `status_category`, `assignee_display_name`, `project_key`) diperbarui dalam transaksi aplikasi yang sama dengan perubahan sumbernya; job rekonsiliasi harian melaporkan (log + metrik) baris yang tidak sinkron.
  13. Terminologi: UI menampilkan "work item"/"work type", sedangkan nama tabel, kolom, endpoint API, JQL, dan nama event tetap memakai `issue`/`issuetype`.
- **Team-managed vs company-managed:** Team-managed menyimpan konfigurasi ber-scope `project_id` dan `access_level` (butir 1, 10); company-managed memakai konfigurasi global via scheme (butir 11).
- **Catatan & ketergantungan:** Aturan key fungsional di PRJ-012 dan PRJ-022 ([04](04-proyek.md)); numbering, hierarki, dan move di ISS-001, ISS-024, ISS-025, ISS-032 ([06](06-work-item-inti.md)). Pola project key kustom (pengaturan Data Center) tidak dibangun di Fase 1 (lihat PRJ-068 dan Pertanyaan terbuka). Hierarki di atas Epic di modul Plans ([18](18-timeline-rencana-tampilan-proyek.md)). Changelog di CMT-023 ([08](08-komentar-lampiran-aktivitas.md)).

### DM-004 Core Data Model: Issue Types, Screens, Field Configurations & Schemes

- **Deskripsi:** Struktur data issue type, priority, resolution, status, screen, field configuration, issue layout, serta scheme yang memetakan konfigurasi ke project, termasuk cache konfigurasi efektif per (project, issue type).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs (konfigurasi global & scheme); Admin Proyek (konfigurasi team-managed, issue layout).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `issue_types`: name, description, `kind` (`standard`/`subtask`), `hierarchy_level`, `avatar_id`, dan `project_id` (scope team-managed); nama unik per scope (global atau per project), dan issue type yang masih dipakai issue tidak dapat dihapus tanpa migrasi ke issue type lain.
  2. Tabel `issue_type_schemes` + `issue_type_scheme_items` (urutan) + `default_issue_type_id` (wajib anggota scheme) dan asosiasi `projects.issue_type_scheme_id` untuk company-managed; Default Issue Type Scheme hasil seed selalu memuat semua issue type global dan tidak dapat dihapus.
  3. Tabel `screens`, `screen_tabs` (urutan), `screen_tab_fields` (field_id, urutan; satu field paling banyak sekali per screen), `screen_schemes` (operasi `create`/`edit`/`view`/`default` → screen; `default` wajib), dan `issue_type_screen_schemes` (issue type atau default → screen scheme).
  4. Tabel `field_configurations` + `field_configuration_items` (field_id, `is_required`, `is_hidden`, `renderer` `wiki`/`rich_text`, description per field) dan `field_configuration_schemes` (issue type atau default → field configuration); field sistem `summary` dan `issuetype` tidak dapat di-hide atau dijadikan opsional (400).
  5. Issue layout disimpan di `issue_layouts`: company-managed menyimpan susunan field per screen di issue view (section, urutan); team-managed menyimpan field per issue type dengan section `context`/`description` dan flag `hide_when_empty`.
  6. Tabel `priorities` (name, icon, color, urutan) dan `priority_schemes` + item + default priority untuk project company-managed; project tanpa asosiasi memakai Default Priority Scheme.
  7. Tabel `resolutions` bersifat global (name unik, description, satu `is_default`); pengisian `issues.resolution_id` diatur workflow (DM-006), bukan oleh scope project.
  8. Tabel `statuses`: name, description, `status_category` (`todo`/`in_progress`/`done`; terpetakan ke key API `new`/`indeterminate`/`done`), dan `project_id` (NULL = global, terisi = status team-managed); status yang masih dipakai workflow aktif tidak dapat dihapus.
  9. Setiap project company-managed menunjuk tepat satu scheme per jenis: permission, notification, issue security (boleh kosong = tanpa keamanan), workflow, issue type, issue type screen, field configuration, dan priority; kolom foreign key di `projects` NOT NULL (kecuali issue security) dengan seed default saat project dibuat.
  10. Penggantian scheme yang memengaruhi data (workflow scheme → pemetaan status, issue type scheme → migrasi issue type) dijalankan sebagai background job berstatus (`queued`/`running`/`succeeded`/`failed`), menampilkan progres persen dan dapat dilanjutkan bila gagal; selama job berjalan, perubahan scheme lain pada project yang sama ditolak 409.
  11. Konfigurasi efektif per (project, issue type) — workflow, screen per operasi, field configuration, context custom field, priority — di-cache dan diinvalidasi saat scheme, screen, field configuration, atau context berubah; pengujian memastikan perubahan field menjadi required langsung berlaku pada create berikutnya.
- **Team-managed vs company-managed:** Company-managed memakai seluruh scheme (butir 2-4, 6, 9); team-managed memakai issue type, field, status, dan layout ber-scope project tanpa scheme (butir 1, 5, 8).
- **Catatan & ketergantungan:** Perilaku fungsional di CFG-001, CFG-003, CFG-008, CFG-009, CFG-042 s.d. CFG-048 ([05](05-konfigurasi-work-item-skema.md)), priority & resolution di ISS-018, ISS-019 ([06](06-work-item-inti.md)), project settings di PRJ-043 s.d. PRJ-052 ([04](04-proyek.md)). Background job dan cache didefinisikan di NFR-003 dan NFR-008 ([27](27-nfr-arsitektur-platform.md)).

### DM-005 Custom Fields Storage (EAV vs JSONB)

- **Deskripsi:** Keputusan dan aturan penyimpanan nilai custom field yang mendukung banyak tipe, context per project/issue type, dan pencarian JQL cepat; clone memakai model hybrid JSONB + tabel bertipe.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering, Architect; Admin Situs (membuat/menghapus custom field global); Admin Proyek (field team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `custom_fields`: `id` numerik dari sequence yang dimulai 10000 sehingga id API berbentuk `customfield_10000`, `customfield_10001`, dst.; `name`, `type_key`, `searcher_key`, `description`, dan `project_id` (NULL = global, terisi = field team-managed); id tidak pernah dipakai ulang setelah field dihapus.
  2. Tabel `custom_field_contexts` (field_id, nama, daftar project dan issue type yang berlaku atau global, `default_value` JSONB) dan options per context; satu (project, issue type) paling banyak dicakup satu context per field, dan context yang lebih spesifik menang atas context global.
  3. Tabel `custom_field_options` (context_id, `parent_option_id` untuk cascading select, `value`, `position`, `disabled`) untuk single/multi select, checkbox, radio, dan cascading select; option child cascading wajib menunjuk parent di context yang sama dan nilai option unik per (context, parent).
  4. Penyimpanan hybrid: kolom `issues.fields` JSONB (kunci `customfield_<id>`) untuk baca cepat issue view/API, ditambah tabel `issue_field_values` bertipe (`issue_id`, `field_id`, `string_value`, `number_value`, `date_value`, `user_id`, `option_id`, `position`) untuk indexing dan JQL; field multi-value disimpan sebagai beberapa baris.
  5. Index GIN (`jsonb_path_ops`) pada `issues.fields` dipakai untuk operator `=`/`IN`; index B-tree pada `issue_field_values (field_id, number_value)` dan `(field_id, date_value)` dipakai untuk operator range `>`, `>=`, `<`, `<=`; pengujian `EXPLAIN` memastikan query JQL range tidak melakukan sequential scan pada data uji 100 ribu issue.
  6. Tipe yang didukung: short text (maks 255), paragraph (rich text ADF, maks ±32.767 karakter), number, date, date time, URL, select single, select multi, checkbox, radio, cascading select, user picker single/multi, group picker single/multi, labels, version picker single/multi, project picker, read-only text, story points/estimate, team, sprint, rank, dan parent; tipe di luar daftar ditolak 400 saat pembuatan field.
  7. ⚠ Batas jumlah custom field per situs dikonfigurasi (default yang diusulkan 700 field aktif, mengikuti guardrail Cloud); di atas ambang peringatan (mis. 80%) halaman admin menampilkan peringatan performa, dan pembuatan field di atas batas ditolak 409 dengan pesan galat yang menyebut batas.
  8. Menghapus custom field memindahkannya ke Trash selama 60 hari (field hilang dari screen, JQL, dan API tetapi nilai tetap tersimpan); restore dalam periode itu mengembalikan field beserta seluruh nilai dan context, dan setelah 60 hari job purge menghapus permanen definisi beserta nilainya.
  9. Perubahan tipe custom field tidak didukung: API update yang mengubah `type_key` ditolak 400; dokumentasi admin mengarahkan membuat field baru lalu memindahkan nilai lewat bulk edit atau otomasi.
  10. Penulisan nilai memperbarui `issues.fields` dan `issue_field_values` dalam transaksi yang sama; kegagalan salah satu me-rollback keduanya, dan job verifikasi dapat membangun ulang `issue_field_values` dari JSONB untuk satu field.
  11. Validasi nilai dilakukan server-side sesuai tipe (format angka, tanggal ISO-8601, URL, user aktif) dan option context yang berlaku; option `disabled` tetap ditampilkan pada issue lama yang sudah memakainya tetapi pemilihan baru ditolak 400.
- **Team-managed vs company-managed:** Field team-managed ber-scope project (butir 1) dan tidak memiliki context terpisah; field company-managed memakai context (butir 2).
- **Catatan & ketergantungan:** Hybrid JSONB + EAV adalah keputusan arsitektur clone (Jira Server/Data Center memakai EAV murni) dan dicatat sebagai keputusan arsitektur di NFR-005 ([27](27-nfr-arsitektur-platform.md)); guardrail batas field di NFR-004. Perilaku fungsional di CFG-025 s.d. CFG-038 dan CFG-040 ([05](05-konfigurasi-work-item-skema.md)); compiler JQL di JQL-050 dan JQL-033 ([12](12-jql-bahasa-kueri.md)); trash/purge di ADM-025 ([22](22-administrasi-sistem.md)).

### DM-006 Data Model: Workflows, Transitions & Rules

- **Deskripsi:** Penyimpanan workflow sebagai graf status-transisi dengan rules (condition, validator, post function) dan workflow scheme, termasuk eksekusi transisi yang atomik.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs (workflow & scheme global); Admin Proyek (workflow team-managed); pengguna dengan Transition issues (eksekusi).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `workflows`: name (unik per scope), description, `project_id` (scope team-managed), `version` (integer naik setiap publish), `is_draft`, `is_active` (dipakai scheme/project), `published_at`; workflow aktif tidak dapat dihapus (409).
  2. Tabel `workflow_statuses` (workflow_id, status_id, `layout_x`, `layout_y`) menyimpan node beserta posisi diagram; satu status paling banyak sekali per workflow.
  3. Tabel `workflow_transitions`: `from_status_id` (NULL = global transition "any status"), `to_status_id`, name, `screen_id` (opsional), `properties` JSONB, dan `is_initial` (transisi create); setiap workflow wajib memiliki tepat satu initial transition dan setiap status harus dapat dicapai dari status awal (validasi saat publish, galat 400 menyebut status yang tidak terjangkau).
  4. Tabel `transition_rules`: `kind` (`condition`/`validator`/`post_function`), `type_key`, `config` JSONB tervalidasi skema per type, `position`, dan untuk condition `group_id` + `group_operator` (`AND`/`OR`) yang memungkinkan kelompok bersarang satu tingkat.
  5. Tabel `workflow_schemes` (default workflow wajib) dan `workflow_scheme_mappings` (issue_type_id → workflow_id); issue type tanpa mapping memakai default workflow.
  6. Publish workflow aktif yang menghapus status yang sedang dipakai issue memerlukan status mapping (status lama → status baru per issue type); tanpa mapping publish ditolak 400, dan dengan mapping sistem membuat job migrasi issue yang mencatat changelog setiap issue terpindah.
  7. Post function sistem wajib — set/clear resolution bila dikonfigurasi, create changelog, update index, fire event — selalu ada pada setiap transisi, tidak dapat dihapus lewat UI maupun API (400), dan dieksekusi dalam urutan tetap.
  8. Eksekusi transisi atomik dalam satu transaksi: evaluasi condition → validator → update issue → post functions → tulis changelog → tulis event outbox (DM-014); kegagalan apa pun sebelum commit me-rollback seluruh perubahan dan mengembalikan galat yang menyebut rule yang gagal; dua transisi bersamaan pada issue yang sama diserialisasi dengan row lock (`SELECT ... FOR UPDATE`).
  9. Team-managed: tepat satu workflow per issue type per project (`project_id` terisi, tanpa scheme), dengan subset rule yang disederhanakan (mis. restrict transition, validate field, auto-assign, update field); type rule di luar subset ditolak 400.
  10. ⚠ Setiap publish menyimpan snapshot versi workflow (`workflow_versions`: version, definisi JSONB, published_by, published_at) untuk audit; rollback ke versi sebelumnya dibuat sebagai draft baru yang harus dipublish ulang (dengan status mapping bila diperlukan).
- **Team-managed vs company-managed:** Company-managed memakai workflow global + workflow scheme + rule lengkap (butir 1-8); team-managed memakai satu workflow per issue type dengan rule sederhana (butir 9).
- **Catatan & ketergantungan:** Perilaku editor, draft, dan scheme di WF-017, WF-018, WF-021, WF-022, WF-033, WF-036 ([09](09-workflow-status-transisi.md)); event transisi di NTF-002 ([19](19-notifikasi-email.md)). Repo saat ini belum memiliki workflow; Fase 1 cukup workflow default sistem (WF-020) yang di-seed ke tabel ini.

## Konten & relasi work item

### DM-007 Data Model: Comments, Worklogs, Attachments & Links

- **Deskripsi:** Skema konten sekunder issue — comment, worklog, attachment, issue link, remote link, dan mention — beserta visibilitas, batasan ukuran, cascade, dan index pagination.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; pengguna dengan Add comments, Work on issues, Create attachments, Link issues; Agent Layanan (comment internal, Fase 5).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `comments`: `id` BIGINT, `issue_id`, `author_id`, `body` ADF JSONB, `created_at`, `updated_at`, `update_author_id`, `visibility_type` (`role`/`group`/NULL) + `visibility_value`, `properties` (via entity properties DM-012), dan `is_public BOOLEAN NOT NULL DEFAULT false` (skema visibilitas komentar tunggal JSM-PRJ-053: di service project comment bersifat internal secara default dan hanya tampil ke customer bila `is_public = true`; `is_internal = NOT is_public` hanya nilai turunan di API/UI, bukan kolom tersimpan; di proyek non-service kolom ini diabaikan oleh aturan tampilan customer); comment dengan visibility hanya dikembalikan ke user anggota role/grup tersebut.
  2. Kombinasi `is_public = true` dengan `visibility_type` non-NULL ditolak 400 dengan pesan galat yang menjelaskan bahwa comment publik untuk customer tidak dapat dibatasi ke role/grup; pemeriksaan ditegakkan di use case dan CHECK constraint `NOT (is_public AND visibility_type IS NOT NULL)`.
  3. Tabel `worklogs`: `id` BIGINT, `issue_id`, `author_id`, `started_at` (timestamptz), `time_spent_seconds` (> 0; 0 atau negatif ditolak 400), `comment`, visibility (sama seperti comment); mode penyesuaian estimate (`auto`/`leave`/`new`/`manual`) adalah parameter use case yang memperbarui `issues.time_estimate` dan `issues.time_spent` dalam transaksi yang sama, bukan kolom tersimpan.
  4. Tabel `attachments`: `id`, `issue_id`, `filename`, `mime_type`, `size_bytes`, `storage_key` (lokasi di object storage, tidak pernah terekspos langsung), `sha256`, `author_id`, `created_at`, `thumbnail_key`, `scan_status` (`pending`/`clean`/`infected`/`error`); file berstatus `infected` tidak dapat diunduh (403) dan `pending` hanya dapat diunduh oleh uploader.
  5. Tabel `issue_link_types` (name unik, `inward_description`, `outward_description`, `is_system` untuk tipe bawaan seperti parent/sub-task yang tidak bisa dihapus) dan `issue_links` (`link_type_id`, `outward_issue_id`, `inward_issue_id`) dengan unique `(link_type_id, outward_issue_id, inward_issue_id)`; link issue ke dirinya sendiri ditolak 400.
  6. Tabel `remote_links`: issue_id, `global_id` (unik per issue, dipakai untuk upsert), `url`, `title`, `summary`, `icon`, `application_type` + `application_name`, dan `relationship` (mis. "mentioned in", "Wiki Page"); POST dengan `global_id` yang sama memperbarui link, bukan menduplikasi.
  7. Batas panjang body comment, description, dan custom field paragraph adalah 32.767 karakter teks (dihitung dari teks ADF yang dirender), ditegakkan di server untuk UI dan API; melebihi batas menghasilkan 400 dengan pesan galat yang menyebut batas.
  8. Menghapus issue men-cascade delete comment, worklog, attachment metadata, link, remote link, dan mention di database dalam satu transaksi; penghapusan file fisik attachment (dan thumbnail) dijalankan asinkron oleh job yang dapat diulang dan tercatat bila gagal.
  9. Index `(issue_id, created_at, id)` pada `comments` dan `worklogs` mendukung pagination keyset dan urutan stabil (terlama/terbaru) tanpa duplikat antar halaman.
  10. Tabel `comment_mentions` (comment_id, mentioned_account_id; juga dipakai untuk mention di description lewat `source_type`) diisi ulang setiap body berubah, dan menjadi sumber notifikasi mention serta pencarian "mentioned me".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku fungsional comment dan attachment di CMT-001 s.d. CMT-019 ([08](08-komentar-lampiran-aktivitas.md)), worklog di TIME-016 s.d. TIME-021 ([14](14-pelacakan-waktu-estimasi.md)), link di ISS-029, ISS-030 ([06](06-work-item-inti.md)) dan CFG-056 ([05](05-konfigurasi-work-item-skema.md)), visibility di PERM-045 ([03](03-izin-peran-keamanan.md)), mention di EDT-014 ([07](07-editor-konten-kaya.md)). Penyimpanan file di CMT-019. Kolom `is_public` (dan nilai turunan `is_internal`) mengikuti skema visibilitas komentar tunggal JSM-PRJ-053 ([29](29-jsm-service-project-request-type-formulir.md)); nilai `public`/`internal` yang dipakai queue/UI (JSM-QSK-012, [31](31-jsm-antrian-sla-knowledge-base-laporan.md)) adalah turunan dari `is_public`.

### DM-008 Data Model: Watchers & Votes

- **Deskripsi:** Penyimpanan watcher dan vote per issue yang dipakai notifikasi, JQL (`watcher`, `voter`, `watchedIssues()`, `votedIssues()`), dan counter di issue view.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering; semua pengguna dengan Browse (watch/vote diri sendiri); pengguna dengan Manage watchers dan View voters and watchers.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `issue_watchers` (issue_id, user_id) dengan unique constraint; saat issue dibuat, reporter dan creator otomatis ditambahkan sebagai watcher bila preferensi `auto_watch` user tersebut aktif (default aktif), dan menambah watcher yang sudah ada bersifat idempoten (204, bukan galat).
  2. Tabel `issue_votes` (issue_id, user_id, created_at) dengan unique constraint; reporter tidak dapat vote issue miliknya sendiri (400), dan vote baru maupun pembatalan vote ditolak (400) pada issue yang sudah memiliki resolution.
  3. Kolom `issues.votes_count` dan `issues.watches_count` diperbarui dalam transaksi yang sama dengan insert/delete pada tabel watcher/vote, sehingga counter selalu sama dengan jumlah baris (diverifikasi job rekonsiliasi).
  4. Opsi global `voting_enabled` dan `watching_enabled` (General configuration) menonaktifkan endpoint dan UI terkait; saat dinonaktifkan, data lama tetap tersimpan tetapi tidak ditampilkan, endpoint vote/watcher mengembalikan 404 dengan pesan galat yang menyebut fitur dinonaktifkan, dan auto-watch saat create tidak dijalankan.
  5. Permission "View voters and watchers" ditegakkan server-side untuk daftar nama (tanpa izin hanya counter yang dikembalikan) dan "Manage watchers" untuk menambah/menghapus watcher selain diri sendiri (403 tanpa izin).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku fungsional di ISS-036, ISS-037 ([06](06-work-item-inti.md)), NTF-012, NTF-014 ([19](19-notifikasi-email.md)), PERM-033 ([03](03-izin-peran-keamanan.md)), toggle di CFG-057 ([05](05-konfigurasi-work-item-skema.md)); JQL di JQL-021, JQL-025, JQL-041 ([12](12-jql-bahasa-kueri.md)); API di API-024 ([23](23-api-publik-webhook.md)).

### DM-009 Data Model: Versions, Components, Labels

- **Deskripsi:** Entitas rilis (version) dan komponen per project serta label global yang dipakai field Fix versions, Affects versions, Components, dan Labels.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Proyek atau pengguna dengan Manage versions (version & component); pengguna dengan Edit issues (mengisi field).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `versions`: `id` BIGINT, `project_id`, `name` (unik case-insensitive per project, maks 255), `description`, `start_date`, `release_date` (tidak boleh sebelum `start_date`; 400), `released`, `archived`, `sequence` (urutan tampilan), dan `driver_account_id`.
  2. Tabel `issue_fix_versions` dan `issue_affects_versions` (issue_id, version_id) many-to-many dengan unique per pasangan; version hanya boleh dari project issue tersebut (400 bila lintas project).
  3. Tabel `components`: `id` BIGINT, `project_id`, `name` unik per project, `lead_account_id`, `default_assignee_type` (`project_default`/`component_lead`/`project_lead`/`unassigned`), `description`, dan `archived`.
  4. Tabel `issue_components` (issue_id, component_id) many-to-many dengan unique per pasangan; component harus milik project issue (400 bila tidak).
  5. Tabel `issue_labels` (issue_id, label) dengan label tanpa spasi (spasi ditolak 400), maks 255 karakter, case-sensitive sesuai perilaku Jira, unik per issue, dan index `label text_pattern_ops` untuk autocomplete prefix lintas situs yang hanya mengembalikan label dari issue yang boleh dilihat user.
  6. Merge/delete version dan delete component menawarkan pemindahan issue ke version/component lain atau pengosongan nilai; pemindahan dijalankan sebagai background job dengan progres dan changelog per issue, dan entitas sumber baru dihapus setelah job sukses.
  7. ⚠ Component dari layanan katalog eksternal (padanan Compass di Cloud) dicatat hanya sebagai integrasi: model menyediakan kolom `components.source` (`native`/`external`) dan `external_id` tanpa sinkronisasi pada fase awal.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada model; ketersediaan menu Releases/Components di team-managed diatur oleh toggle fitur project ([04](04-proyek.md)).
- **Catatan & ketergantungan:** Perilaku fungsional di REL-001 s.d. REL-010, REL-028 s.d. REL-036 ([15](15-rilis-versi-komponen.md)), label di ISS-020 ([06](06-work-item-inti.md)), JQL di JQL-022, JQL-027 ([12](12-jql-bahasa-kueri.md)). Integrasi katalog component eksternal di REL-035 (di luar cakupan awal).

## Perencanaan agile

### DM-010 Data Model: Boards, Sprints & Ranking (LexoRank)

- **Deskripsi:** Skema board (berbasis saved filter untuk company-managed), konfigurasi kolom, sprint, histori dan scope change sprint, serta ranking global berbasis LexoRank yang efisien untuk drag-and-drop.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; board admin (konfigurasi board); pengguna dengan Manage sprints (sprint); pengguna dengan Schedule issues/Edit issues (rank).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `boards`: `id` BIGINT, name, `type` (`scrum`/`kanban`), `filter_id` (wajib untuk company-managed), `location_type` + `location_id` (project atau user), dan konfigurasi JSONB tervalidasi skema untuk swimlane strategy, quick filters, card colors, card layout, estimation statistic, time tracking statistic, working days & non-working days, sub-filter kanban, serta toggle Kanban backlog; board admin disimpan di `board_admins` (user/group) dan minimal satu admin wajib ada.
  2. Tabel `board_columns` (board_id, name, position, `min_issues`, `max_issues` dengan `min <= max`, `is_kanban_backlog`) dan `board_column_statuses` (column_id, status_id); satu status paling banyak dipetakan ke satu kolom per board, status tak terpetakan tampil sebagai "Unmapped", dan kolom Backlog hanya berlaku untuk board kanban dengan Kanban backlog aktif.
  3. Tabel `sprints`: `id` BIGINT, `origin_board_id`, name (maks 255, wajib), `goal`, `state` (`future`/`active`/`closed`), `start_date`, `end_date` (wajib saat start, tidak boleh sebelum start), `complete_date`, dan `sequence` (urutan tampilan sprint di backlog, dipakai SPR-017 Sprint Reorder dan SPR-023 Parallel Sprints; default = `id`); transisi state hanya future → active → closed (reopen closed → active diatur SPR-021).
  4. Tabel `issue_sprints` (issue_id, sprint_id, added_at, removed_at) menyimpan histori: issue boleh tercatat di banyak sprint closed tetapi maksimal satu sprint `active`/`future` sekaligus (ditegakkan di use case + constraint); saat complete sprint, issue yang belum resolved dipindah ke backlog atau sprint berikut sesuai pilihan pengguna dan tetap tercatat di sprint lama.
  5. Tabel `sprint_scope_events` (sprint_id, issue_id, event `added`/`removed`/`estimate_changed`, nilai lama/baru, occurred_at) dicatat untuk setiap perubahan setelah sprint dimulai, dan `sprint_start_snapshots` menyimpan estimate setiap issue saat start; Sprint report, Burndown, dan Velocity dihitung dari kedua tabel ini, bukan dari nilai issue terkini.
  6. Flag global `parallel_sprints_enabled` (company-managed) mengizinkan lebih dari satu sprint `active` per board; tanpa flag, start sprint kedua ditolak 409.
  7. Rank bersifat global per instance (satu nilai `issues.rank` per issue, bukan per board) disimpan sebagai string LexoRank (`bucket|value:suffix`, bucket `0`/`1`/`2` berotasi saat rebalance) dengan index B-tree berkolasi `"C"` sehingga urutan byte sama dengan urutan rank.
  8. Operasi rank before/after menghitung nilai di antara dua tetangga dan hanya memperbarui baris issue yang dipindah; pengujian memastikan satu drag-and-drop menghasilkan tepat satu UPDATE pada tabel `issues`.
  9. Rebalancing LexoRank dijalankan sebagai background job saat panjang rank melewati ambang (mis. 128 karakter) atau tidak ada ruang di antara tetangga, memakai advisory lock Postgres agar hanya satu rebalance berjalan; operasi rank selama rebalance tetap berhasil (menulis ke bucket baru) dan urutan relatif tidak berubah.
  10. Batch rank menerima maksimal 50 issue per request (kompatibel batas Agile REST API); permintaan berisi lebih dari 50 issue ditolak 400.
  11. Board team-managed dibangun dari konfigurasi project (issue project tersebut + status workflow lokal), tidak memiliki `filter_id`, dan tidak dapat diarahkan ke saved filter (400).
  12. ⚠ Batas tampilan board: bila hasil query board melebihi batas yang dikonfigurasi (default yang diusulkan 5.000 issue), board menampilkan peringatan dan memuat hanya batas tersebut dengan saran mempersempit filter; angka batas merupakan setting situs.
- **Team-managed vs company-managed:** Company-managed board berbasis saved filter dan bisa lintas project, dengan parallel sprints opsional (butir 1, 6); team-managed board dibangun dari konfigurasi project (butir 11).
- **Catatan & ketergantungan:** LexoRank cukup kompatibel perilaku, bukan format byte-per-byte. Perilaku fungsional di BRD-003, BRD-010, BRD-017, BRD-046, BRD-052 ([10](10-board-kolom-kartu.md)), SPR-003, SPR-014 s.d. SPR-023 ([11](11-backlog-sprint.md)), ISS-022 ([06](06-work-item-inti.md)), laporan sprint di TIME-029 ([14](14-pelacakan-waktu-estimasi.md)) dan [16](16-laporan-insight.md). API rank di API-033 ([23](23-api-publik-webhook.md)).

## Filter, dashboard & data ekstensi

### DM-011 Data Model: Filters, Dashboards & Sharing

- **Deskripsi:** Penyimpanan saved filter (JQL), subscription, dashboard, gadget, dan favourites beserta share permission dan edit permission.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering; pemilik filter/dashboard; pengguna dengan global permission Share dashboards and filters dan Manage group filter subscriptions; Admin Situs (change owner).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `filters`: `id` BIGINT, `owner_id`, `name` (unik case-insensitive per owner, maks 255), `jql` (teks), `description`, `favourite_count` (terdenormalisasi), dan `column_layout` JSONB; JQL disimpan apa adanya dan divalidasi saat disimpan (JQL tidak valid ditolak 400).
  2. Tabel `share_permissions` (entity_type `filter`/`dashboard`, entity_id, `access` `view`/`edit`, `share_type` `user`/`group`/`project`/`project_role`/`organization` (any logged-in user)/`public`, param); share `public` hanya diterima bila setting situs mengizinkan akses anonim (400 bila tidak), dan pengguna tanpa Share dashboards and filters hanya dapat menyimpan entitas privat.
  3. Tabel `filter_subscriptions`: filter_id, `schedule` (interval sederhana atau ekspresi cron) dievaluasi dalam timezone pemilik subscription, `recipient_type` (`user`/`group`), recipient_id, `email_when_empty` (default false), `last_run_at`, `next_run_at`; subscription ke group memerlukan permission Manage group filter subscriptions (403).
  4. Tabel `favourites` (user_id, entity_type `filter`/`dashboard`/`board`/`project`/`plan`, entity_id, position) unik per (user, entity_type, entity_id); favourite pada entitas yang kemudian tidak boleh dilihat user disembunyikan dari daftar, bukan dihapus.
  5. Tabel `dashboards`: `id` BIGINT, `owner_id`, name, description, `layout` (`A`/`AA`/`BA`/`AB`/`AAA` atau representasi jumlah kolom setara), share & edit permission via `share_permissions`, dan `is_system`; tepat satu System dashboard yang tidak dapat dihapus dan menjadi default bagi user tanpa dashboard favorit.
  6. Tabel `dashboard_gadgets`: dashboard_id, `type_key`, `column_index`, `row_index` (unik per dashboard), `color`, `config` JSONB tervalidasi skema per type, dan `refresh_interval_minutes` (NULL = tanpa auto-refresh, atau 15/30/60/120).
  7. Filter yang mereferensikan field, project, atau nilai yang telah dihapus diberi status `invalid` dengan pesan galat JQL yang disimpan; filter tidak dihapus, board/gadget yang memakainya menampilkan galat yang menjelaskan, dan subscription-nya dilewati serta dicatat di log.
  8. Admin Situs dapat mengganti `owner_id` filter dan dashboard (mis. saat pemilik dinonaktifkan); pergantian owner mempertahankan id, share permission, subscription, dan referensi board/gadget, serta tercatat di audit log.
  9. Subset Fase 1 (prasyarat board): tabel `filters` (butir 1) dan `share_permissions` khusus `entity_type = filter` (butir 2, berfungsi sebagai filter shares) dibuat dan dipakai sejak Fase 1 karena setiap board company-managed wajib memiliki `filter_id` (DM-010 butir 1, BRD-010); di Fase 1 filter dibuat otomatis saat board dibuat (owner = pembuat board, JQL `project = X ORDER BY Rank ASC`) dan dapat diubah query, owner, serta share-nya lewat konfigurasi board, sedangkan `filter_subscriptions`, `favourites`, `dashboards`, `dashboard_gadgets`, serta UI saved filter dan filter sharing umum di navigator (SRC-036, SRC-045) menyusul di Fase 2 tanpa migrasi data yang merusak.
  10. ⚠ Di Fase 1, sebelum mesin JQL penuh tersedia (Fase 2), validasi JQL filter board dibatasi pada subset klausa yang didukung pencarian dasar (mis. `project`, `issuetype`, `status`, `assignee`, `labels`, `component`, `fixVersion`, dan `ORDER BY Rank`); JQL di luar subset ditolak 400 dengan pesan yang menyebut klausa yang belum didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada model; board team-managed tidak memakai `filters` (DM-010 butir 11), sehingga subset Fase 1 (butir 9) hanya diperlukan oleh board company-managed.
- **Catatan & ketergantungan:** Prioritas fitur secara keseluruhan tetap P1 / Fase 2 (subscription, favourites, dashboard), tetapi tabel `filters` dan share filter adalah subset Fase 1 (butir 9) agar konsisten dengan BRD-010 ([10](10-board-kolom-kartu.md)) yang berstatus P0 / Fase 1 dan board DM-010. Perilaku fungsional filter di SRC-036 s.d. SRC-049 ([13](13-pencarian-filter-ekspor.md)), subscription email di NTF-030 ([19](19-notifikasi-email.md)), dashboard di DSH-001 s.d. DSH-018 ([17](17-dashboard-gadget.md)), shared items admin di ADM-018, ADM-019 ([22](22-administrasi-sistem.md)), permission di PERM-009, PERM-010, PERM-051 ([03](03-izin-peran-keamanan.md)).

### DM-012 Data Model: Entity Properties & Automation Rules

- **Deskripsi:** Penyimpanan data ekstensi (entity properties) untuk berbagai entitas serta automation rule beserta riwayat eksekusi dan penghitungan pemakaian.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Engineering; Integrator/Developer API dan app (entity properties); Admin Situs dan Admin Proyek (automation rule sesuai scope).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `entity_properties`: `entity_type` (`issue`, `project`, `user`, `comment`, `worklog`, `issue_type`, `dashboard_item`, `board`, `sprint`), `entity_id`, `key` (maks 255 karakter), `value` JSONB (maks 32 KB per property setelah serialisasi), `created_at`, `updated_at`, dengan unique `(entity_type, entity_id, key)`; nilai melebihi 32 KB atau key melebihi 255 karakter ditolak 400, dan property ikut terhapus saat entitasnya dihapus.
  2. Indexed entity properties (P3, opsional bagi clone): definisi ekstraksi `path` JSON → alias JQL bertipe (`string`, `number`, `date`, `user`, `text`) disimpan di `entity_property_indexes` dan nilainya diproyeksikan ke tabel index agar dapat dicari dengan `issue.property[key].path` di JQL.
  3. Tabel `automation_rules`: name, `definition` JSONB berbentuk tree (trigger, conditions, branches, actions), `scope` (`global`/`multi_project`/`project_type`/`single_project` + daftar project), `actor` (akun sistem "Automation" atau user tertentu), `owner_id`, `state` (`enabled`/`disabled`), `version` (naik setiap perubahan), dan `allow_rule_trigger` (rule ini boleh dipicu oleh aksi rule lain; default false).
  4. Tabel `automation_executions` (rule_id, rule_version, trigger event id, status `success`/`some_errors`/`no_actions_performed`/`throttled`/`failed`/`loop_detected`, started_at, duration, issue terdampak, log per komponen) menjadi audit log per rule dengan retensi terbatas (default yang diusulkan 90 hari) yang dibersihkan job purge.
  5. Tabel `automation_usage` (bulan, scope/project, jumlah rule run) dihitung untuk usage limit per plan: hanya eksekusi yang menjalankan aksi yang dihitung, dan saat batas plan tercapai eksekusi berikutnya berstatus `throttled` hingga periode berikut.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada model; Admin Proyek team-managed hanya dapat membuat rule ber-scope single project.
- **Catatan & ketergantungan:** Entity properties fungsional untuk work item di DM-013 dan API-025 ([23](23-api-publik-webhook.md)); JQL property di JQL-036 ([12](12-jql-bahasa-kueri.md)); automation di AUTO-003 s.d. AUTO-016 ([20](20-otomasi.md)); actor di PERM-049 ([03](03-izin-peran-keamanan.md)). Audit log keamanan didefinisikan di NFR-017 ([27](27-nfr-arsitektur-platform.md)) dan keandalan webhook di API-066 ([23](23-api-publik-webhook.md)), bukan di sini.

### DM-013 Work Item Entity Properties

- **Deskripsi:** Penyimpanan key-value JSON per work item lewat REST API untuk integrasi dan app; tidak tampil sebagai field di UI.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, app, sistem; pengguna dengan Browse projects (baca) dan Edit issues (tulis/hapus) pada work item terkait.
- **Aturan perilaku (acceptance criteria):**
  1. REST menyediakan `GET /rest/api/3/issue/{issueIdOrKey}/properties` (daftar key), `GET`/`PUT`/`DELETE .../properties/{propertyKey}` untuk satu property; `PUT` membuat (201) atau memperbarui (200), `DELETE` property yang tidak ada mengembalikan 404, dan body bukan JSON valid ditolak 400; data disimpan di `entity_properties` dengan `entity_type = issue` (DM-012).
  2. Property dapat diindeks untuk JQL bila didaftarkan lewat descriptor app/ekstensi (DM-012 butir 2); property yang tidak terdaftar tetap tersimpan tetapi tidak dapat dicari, dan JQL `issue.property[...]` atas key tak terindeks mengembalikan hasil kosong, bukan galat.
  3. Bulk set property lewat REST (mis. `PUT /rest/api/3/issue/properties/{propertyKey}` dengan filter daftar issue id atau JQL) berjalan sebagai task asinkron yang mengembalikan URL status; issue yang tidak dapat diedit user dilewati dan dilaporkan di hasil task.
  4. Akses mengikuti permission work item: baca memerlukan Browse projects (dan lolos issue security), tulis/hapus memerlukan Edit issues; tanpa izin baca respons 404, dengan izin baca tetapi tanpa izin edit respons 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model penyimpanan di DM-012; API di API-025 dan API-026 ([23](23-api-publik-webhook.md)); ekstensi/app di modul 25 ([25](25-integrasi-impor-ekspor-ekstensi.md)). Tidak tumpang tindih penuh dengan DM-012 (DM-012 mencakup model semua entitas + automation), sehingga tidak digabung.

## Event domain & graf entitas

### DM-014 Issue Event Catalog & Domain Events

- **Deskripsi:** Katalog event issue standar beserta event non-issue yang menjadi kontrak bersama untuk notification scheme, webhook, automation trigger, indexer pencarian, dan pembaruan real-time, ditulis lewat transactional outbox.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs (custom event, company-managed); konsumen internal (notifikasi, otomasi, indexer, realtime) dan eksternal (webhook).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `event_types` di-seed dengan system event (id stabil, tidak dapat dihapus/diubah namanya): Issue Created, Issue Updated, Issue Assigned, Issue Resolved, Issue Closed, Issue Commented, Issue Comment Edited, Issue Comment Deleted, Issue Reopened, Issue Deleted, Issue Moved, Work Logged on Issue, Work Started on Issue, Work Stopped on Issue, Issue Worklog Updated, Issue Worklog Deleted, dan Generic Event; setiap operasi use case menembakkan tepat satu event utama sesuai katalog (mis. transisi ke status Done dengan resolution → Issue Resolved).
  2. Custom event (company-managed) dapat dibuat Admin Situs (`event_types.is_system = false`, name unik), ditembakkan dari post function transisi sebagai pengganti event default, dan dipetakan di notification scheme (DM-002 butir 4); custom event yang masih dirujuk transisi atau notification scheme tidak dapat dihapus (409).
  3. Event non-issue untuk webhook dan otomasi tersedia minimal untuk: project (created/updated/deleted/archived/restored), version (created/updated/released/unreleased/moved/merged/deleted), component, sprint (created/started/updated/closed/deleted), board (created/updated/configuration changed/deleted), user (created/updated/deleted), filter, attachment (created/deleted), dan issue link (created/deleted).
  4. Setiap perubahan data menulis baris ke tabel `outbox_events` dalam transaksi database yang sama (exactly-once write); dispatcher terpisah mengirim ke konsumen dengan jaminan at-least-once, menandai `dispatched_at`, mengulang dengan backoff bila gagal, dan tidak ada event yang terkirim untuk transaksi yang di-rollback.
  5. Setiap event memiliki `event_id` unik (UUID) untuk idempotensi konsumen dan `sequence` monoton per issue (atau per entitas untuk event non-issue) untuk ordering best-effort; konsumen internal (indexer, realtime) mengabaikan event dengan sequence lebih kecil dari yang sudah diproses untuk entitas yang sama.
  6. Payload event membawa flag `suppress_notifications` (berasal dari bulk change dengan opsi "send notification" dimatikan atau API `notifyUsers=false`); flag ini menghentikan pengiriman notifikasi email/in-app tetapi tidak menghentikan webhook, automation, indexer, maupun changelog.
- **Team-managed vs company-managed:** Custom event hanya tersedia di company-managed (butir 2); team-managed hanya memakai system event.
- **Catatan & ketergantungan:** Konsumen event: notifikasi NTF-001 s.d. NTF-003 ([19](19-notifikasi-email.md)), workflow WF-016 ([09](09-workflow-status-transisi.md)), board & sprint BRD-053 ([10](10-board-kolom-kartu.md)), automation trigger AUTO-020, AUTO-022 ([20](20-otomasi.md)), worklog TIME-039 ([14](14-pelacakan-waktu-estimasi.md)), real-time ISS-014 ([06](06-work-item-inti.md)) dan NAV-046 ([21](21-navigasi-ux-preferensi.md)), indexer SRC-002 ([13](13-pencarian-filter-ekspor.md)), webhook di [23](23-api-publik-webhook.md). Implementasi dispatcher outbox dan background job di NFR-006 dan NFR-008, real-time di NFR-009 ([27](27-nfr-arsitektur-platform.md)); katalog webhook di API-065. Repo saat ini belum memiliki event bus; outbox berbasis Postgres cukup untuk Fase 1.

### DM-015 Cross-product Entity Graph (Teamwork Graph)

- **Deskripsi:** Lapisan relasi entitas terpadu yang menghubungkan work item dengan page, pull request, commit, deployment, goal, team, user, dan objek aplikasi terhubung; menjadi sumber panel related content dan fitur AI opsional. Bagi clone, implikasinya adalah model edge entitas generik yang memperhatikan permission, setara konsep Teamwork Graph.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Engineering, Architect; Developer/Anggota Tim dan pengguna akhir (melihat related content sesuai izin); Integrator/Developer API (konektor).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `entity_edges`: `source_type`, `source_id`, `relation_type` (mis. `links_to`, `mentions`, `implements`, `deployed_in`, `contributes_to`, `owned_by`, `watched_by`), `target_type`, `target_id` atau `target_global_id` (untuk objek eksternal), `provider`, `created_at`, dengan unique `(source_type, source_id, relation_type, target_type, target_ref)`; mendukung relasi work item ↔ page, PR, commit, deployment, goal, team, user, dan objek eksternal.
  2. Data yang sudah ada (`issue_links`, `remote_links`, dev info, `comment_mentions`, `issue_watchers`) diekspos sebagai edge lewat view/proyeksi baca, bukan disalin sebagai sumber kebenaran kedua; perubahan di tabel sumber langsung tercermin di graph.
  3. Penelusuran graph (mis. "konten terkait work item ini", "work item yang dikerjakan tim X") dibatasi kedalaman maksimal (default 2 hop) dan jumlah hasil maksimal per request (default 50, dengan cursor); permintaan melebihi batas ditolak 400.
  4. Setiap node hasil difilter ulang oleh authorizer sumbernya (Browse + issue security untuk issue; izin provider untuk objek eksternal) sebelum dikembalikan; node yang tidak boleh dilihat dihapus dari hasil tanpa jejak (tidak ada placeholder, count, atau judul yang bocor).
  5. Konektor pihak ketiga/SCM memperkaya graph lewat ingest webhook atau API dengan kunci idempotensi (provider + id eksternal + relasi); pengiriman ulang payload yang sama tidak membuat edge ganda.
  6. Graph menjadi sumber data untuk panel related content di issue view dan untuk pencarian/AI opsional; perilaku fungsional panel dan AI didefinisikan di NAV-015 dan NAV-051 ([21](21-navigasi-ux-preferensi.md)), bukan di sini.
  7. Saat entitas dihapus atau integrasi dicabut, edge terkait dihapus oleh job cascade asinkron; sampai job selesai, edge yatim tetap disaring oleh butir 4 sehingga tidak pernah tampil.
  8. ⚠ API graph publik untuk developer (padanan Teamwork Graph API) hanya disediakan sebagai endpoint baca terbatas setelah model internal stabil; cakupan dan bentuknya diputuskan bersama modul API publik ([23](23-api-publik-webhook.md)).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur platform yang muncul lewat related content dan AI (Rovo) di platform asal, bukan fitur mandiri. Sumber edge: DEV-001 ([24](24-integrasi-devops.md)), ISS-029, ISS-030 ([06](06-work-item-inti.md)), Goals & Teams di PLAN-057, PLAN-061 ([18](18-timeline-rencana-tampilan-proyek.md)); pencarian AI di NAV-015, NAV-051 ([21](21-navigasi-ux-preferensi.md)); outbound request konektor dijaga INT-001 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

## Keputusan arsitektur

Keputusan model data Fase 1 dicatat sebagai ADR di [`docs/adr/`](../adr/README.md) (S0-15, KAN-12):

| ADR | Keputusan | Fitur terkait |
| --- | --- | --- |
| [ADR-0005](../adr/0005-tipe-identitas.md) | `projects`/`issues` memakai `BIGINT`; `users.id` tetap UUID `VARCHAR(36)`; `account_id` kolom terpisah untuk API publik. | DM-001, DM-003, DM-005 |
| [ADR-0006](../adr/0006-penyimpanan-custom-field.md) | Custom field disimpan hybrid: `issues.fields` JSONB + tabel bertipe `issue_field_values`. | DM-005 |
| [ADR-0007](../adr/0007-format-rich-text.md) | Rich text disimpan sebagai JSON ala ADF (editor TipTap/ProseMirror); konversi di tepi. | DM-003, DM-005, EDT-001, EDT-002 |

## Pertanyaan terbuka

1. Apakah id internal `users` tetap UUID `VARCHAR(36)` (kondisi repo) dengan `account_id` terpisah, atau `account_id` langsung memakai nilai UUID yang ada agar migrasi lebih sederhana? **Terjawab ([ADR-0005](../adr/0005-tipe-identitas.md)):** `users.id` tetap UUID `VARCHAR(36)` dan `account_id` kolom terpisah.
2. Apakah login dengan `username` (kolom repo saat ini) tetap didukung sebagai alias legacy, atau dihapus sepenuhnya dan diganti email saja mengikuti model Cloud?
3. Apakah model hybrid JSONB + tabel bertipe untuk custom field disetujui sebagai keputusan final, atau cukup EAV murni pada Fase 1 lalu JSONB ditambahkan saat dibutuhkan performa? **Terjawab ([ADR-0006](../adr/0006-penyimpanan-custom-field.md)):** hybrid disetujui sebagai keputusan final.
4. Berapa angka default batas jumlah custom field per situs dan batas issue yang ditampilkan board yang akan dipakai produk ini?
5. Apakah riwayat versi workflow dengan rollback dibangun sejak Fase 2 bersama editor workflow, atau hanya snapshot audit tanpa fungsi rollback?
6. Apakah pola project key kustom (pengaturan khas Data Center) disediakan sebagai opsi untuk instance self-hosted?
7. Apakah outbox cukup memakai Postgres (polling/`LISTEN`/`NOTIFY`) untuk jangka panjang, atau perlu broker pesan terpisah mulai Fase 3 ketika webhook dan otomasi aktif?
8. Berapa lama retensi log eksekusi automation dan berapa batas rule run per plan?
9. Apakah graf entitas (DM-015) dibangun di Fase 4 meskipun fitur AI belum dipastikan, atau ditunda hingga ada konsumen nyata selain panel related content?

## Di luar cakupan & ekosistem

- **Format internal milik Atlassian** — format byte-per-byte LexoRank, skema database Jira Server/Data Center (EAV `customfieldvalue`, OSWorkflow XML), dan id internal Atlassian tidak direplikasi; clone hanya kompatibel perilaku dan bentuk API.
- **Fitur Data Center-only** — pola project key kustom (`jira.projectkey.pattern`), backup XML Data Center, dan pengaturan server lain hanya menjadi rujukan (lihat PRJ-068 di [04](04-proyek.md) dan ADM-040 di [22](22-administrasi-sistem.md)).
- **App Marketplace pihak ketiga** — app yang menyimpan data lewat entity properties atau custom field type Forge/Connect tidak dibangun sebagai fitur native; modul ini hanya menyediakan tabel dan API penyimpanannya.
- **Layanan Atlassian lintas produk** — Teamwork Graph milik Atlassian, Rovo Search/Chat, Compass, Confluence, dan Atlassian Goals/Teams sebagai layanan eksternal tidak diintegrasikan secara native; clone hanya menyediakan model edge sendiri (DM-015) dan kolom referensi eksternal (DM-009 butir 7).
- **Nested group** — grup bersarang tidak didukung (sama seperti platform asal) dan tidak akan dibangun.
- **Perubahan tipe custom field** — sengaja tidak didukung (DM-005 butir 9).
- **Audit log keamanan dan keandalan pengiriman webhook** — didefinisikan di NFR-017 ([27](27-nfr-arsitektur-platform.md)) dan API-066 ([23](23-api-publik-webhook.md)), bukan di modul ini.
