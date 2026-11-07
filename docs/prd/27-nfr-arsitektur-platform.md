# 27 — Kebutuhan Non-fungsional & Arsitektur Platform

Modul ini menetapkan kebutuhan non-fungsional (NFR) dan batasan arsitektur yang berlaku lintas seluruh modul produk ini: target performa dan skala, strategi caching dan batas sistem (guardrail), aturan platform database dan migrasi, kontrol konkurensi, semantik waktu, job asinkron, pembaruan real-time, ketersediaan dan pemulihan bencana, observability, pengemasan dan deployment single binary, model tenancy, baseline keamanan aplikasi, pengelolaan secret dan enkripsi, audit log keamanan, privasi dan kepatuhan, arsitektur SPA, internasionalisasi, aksesibilitas, dukungan browser, serta strategi pengujian. Setiap butir di modul ini adalah acceptance criteria yang wajib dipenuhi oleh fitur fungsional di modul lain; modul fungsional cukup merujuk ID NFR alih-alih mengulang angkanya. Angka latency, kapasitas, dan batas di modul ini adalah target internal clone, bukan SLA resmi vendor pembanding. Kondisi repo saat ini (backend Go Clean Architecture dengan Gin, GORM, PostgreSQL; SPA SvelteKit ter-embed lewat `go:embed`; baru fitur auth dan user) menjadi titik awal: beberapa temuan baseline yang wajib dikoreksi dicatat di bagian "Catatan & ketergantungan" fitur terkait (mis. migrasi `001_create_users_table` dengan `DEFAULT VARCHAR(36)` yang tidak valid dan `TIMESTAMP` tanpa zona waktu, CORS `Access-Control-Allow-Origin: *`, docker-compose bernama `yopatungan` yang me-mount `./migrations` ke `/docker-entrypoint-initdb.d`, `server.Shutdown(context.Background())` tanpa batas waktu, Makefile yang menunjuk `./migration/` padahal folder bernama `migrations/`, serta image final `alpine:latest` yang berjalan sebagai root).

Modul terkait:

- [00 Visi, Ruang Lingkup & Baseline](00-visi-ruang-lingkup-glosarium-baseline.md) — penamaan produk (VISI-001), plan tier & feature flag (VISI-003) yang dipakai NFR-014 dan NFR-019.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — sesi, 2FA, dan jalur migrasi JWT lama (AUTH-007, AUTH-008), visibilitas profil (AUTH-002), penutupan akun (AUTH-014).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — data residency (USR-042), BYOK (USR-043), API data pribadi GDPR (USR-044), audit log organisasi dan Jira (USR-045, USR-046; retensi mengikuti NFR-017 butir 5).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — otorisasi server-side (PERM-002) yang menjadi dasar proteksi IDOR dan penyaringan event real-time.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field guardrail (CFG-039), translations field dan issue type (CFG-061).
- [06 Work Item Inti](06-work-item-inti.md) — real-time & concurrent editing (ISS-014), Rank/LexoRank (ISS-022), bulk change (ISS-035), batas work item (ISS-044).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — file storage (CMT-019), keamanan attachment (CMT-017), model changelog (CMT-023).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — batas workflow (WF-025), audit workflow (WF-039), translations status (WF-040).
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — budget performa board/backlog, working days (BRD-040).
- [12 JQL](12-jql-bahasa-kueri.md) — batas eksekusi JQL (JQL-052) dan compiler JQL → SQL.
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — arsitektur search engine (SRC-001), strategi indexing (SRC-002), batas pencarian (SRC-006).
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — time tracking settings (TIME-002).
- [19 Notifikasi & Email](19-notifikasi-email.md) — pipeline notifikasi (NTF-008), infrastruktur email (NTF-023), lokalisasi notifikasi (NTF-010).
- [20 Otomasi](20-otomasi.md) — batas penggunaan dan layanan otomasi (AUTO-015, AUTO-016), audit log rule otomasi (AUTO-014, scope `automation` di NFR-017), pencegahan loop.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — bahasa & zona waktu (NAV-033), tema (NAV-034), aksesibilitas personal (NAV-035, NAV-036), keyboard shortcut (NAV-038 s.d. NAV-041), live update (NAV-046).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log admin (ADM-014, retensi mengikuti NFR-017 butir 5), halaman admin Background jobs (dimiliki NFR-008), data lifecycle (ADM-025), backup & restore (ADM-035), batas objek (ADM-041), alat operasional DC (ADM-057, Fase —, hanya rujukan), indexing (ADM-058).
- [23 API Publik & Webhook](23-api-publik-webhook.md) — versioning & OpenAPI (API-002), rate limiting (API-006), async task API (API-007), reliabilitas webhook.
- [24 Integrasi DevOps](24-integrasi-devops.md) — model data development information (DEV-001).
- [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — SSRF guard (INT-001), performa impor/ekspor, strategi plugin.
- [28 Model Data Inti](28-model-data-inti.md) — skema tabel inti, alokasi issue key, katalog event domain yang dipakai outbox (NFR-008) dan real-time (NFR-009).

## Persona terkait

- **Admin Situs** — dalam konteks self-hosted juga berperan sebagai operator/SRE: memasang, mengonfigurasi, memantau, mencadangkan, dan meng-upgrade instance; membaca audit log dan mengatur retensi.
- **Developer/Anggota Tim** — pengguna harian yang merasakan langsung budget performa board, backlog, issue view, dan pencarian, serta pembaruan real-time.
- **Scrum Master/Team Lead** — bergantung pada konsistensi data sprint, burndown, dan working days lintas zona waktu.
- **Stakeholder/Viewer** — membuka tautan dalam (deep link) dari email/chat dengan berbagai browser dan perangkat, termasuk mobile web.
- **Pengguna Tim Bisnis** — memakai UI berbahasa Indonesia/Inggris dan fitur aksesibilitas (keyboard, kontras, reduced motion).
- **Integrator/Developer API** — mengandalkan batas yang terdokumentasi, format waktu ISO 8601, error code stabil, idempotensi webhook, dan perilaku konkurensi (409 Conflict).
- **Customer/Pelapor** dan **Agent Layanan** — menerima email dan portal dalam bahasa masing-masing serta menuntut ketersediaan tinggi untuk layanan (modul 29-35).
- Peran internal non-persona yang menjadi aktor NFR: tim Engineering, QA, SRE, Security, Architect/DBA, dan Legal pemilik produk; mereka dirujuk di baris "Aktor & izin".

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin issue view dan board terbuka dalam hitungan ratusan milidetik bahkan pada instance berisi jutaan work item agar alur kerja harian tidak tersendat (NFR-001, NFR-002).
- Sebagai Developer/Anggota Tim, saya ingin melihat perubahan rekan di board dan issue view tanpa reload, dan diberi tahu bila work item yang saya edit telah diubah orang lain, agar saya tidak menimpa pekerjaan mereka (NFR-009, NFR-006).
- Sebagai Admin Situs, saya ingin men-deploy satu binary atau satu image Docker dengan konfigurasi lewat environment variable yang divalidasi saat start agar pemasangan dan upgrade mudah dan aman (NFR-013, NFR-005).
- Sebagai Admin Situs, saya ingin endpoint health, graceful shutdown, dan rolling update tanpa downtime agar instance tetap tersedia saat upgrade (NFR-010).
- Sebagai Admin Situs, saya ingin memantau log terstruktur, metrik, trace, dan antrean job dari dashboard serta menerima alert agar gangguan terdeteksi sebelum pengguna melapor (NFR-012, NFR-008).
- Sebagai Admin Situs, saya ingin RPO dan RTO terdokumentasi beserta runbook failover yang diuji berkala agar data tim tidak hilang saat terjadi bencana (NFR-011).
- Sebagai Admin Situs, saya ingin audit log append-only yang dapat dicari dan diekspor agar setiap perubahan konfigurasi dan keamanan dapat dipertanggungjawabkan (NFR-017).
- Sebagai Admin Situs, saya ingin mengekspor data pribadi pengguna dan menganonimkan akun yang dihapus agar instance mematuhi GDPR dan UU PDP (NFR-018).
- Sebagai Admin Situs di organisasi teregulasi, saya ingin matriks kontrol dan ekspor bukti otomatis agar dapat menyiapkan audit SOC 2/ISO 27001 secara mandiri (NFR-019).
- Sebagai Integrator/Developer API, saya ingin batas sistem, error code, dan format tanggal yang stabil serta terdokumentasi agar integrasi saya tidak rusak secara tak terduga (NFR-004, NFR-007, NFR-021).
- Sebagai Pengguna Tim Bisnis, saya ingin memakai seluruh UI dengan keyboard, screen reader, tema gelap, dan bahasa Indonesia agar produk ini dapat dipakai semua anggota tim (NFR-021, NFR-022).
- Sebagai Stakeholder/Viewer, saya ingin membuka tautan work item di browser modern apa pun termasuk ponsel agar saya dapat meninjau pekerjaan di mana saja (NFR-020, NFR-023).
- Sebagai Scrum Master/Team Lead, saya ingin tanggal jatuh tempo, sprint end date, dan fungsi tanggal JQL konsisten dengan zona waktu saya agar laporan dan filter tidak meleset sehari (NFR-007).
- Sebagai Product Owner/PM, saya ingin setiap perubahan diuji otomatis untuk regresi fungsional, performa, keamanan, aksesibilitas, dan migrasi agar kualitas rilis terjaga (NFR-024, NFR-015).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 11 |
| P1 | 10 |
| P2 | 2 |
| P3 | 1 |
| **Total** | **24** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 11 |
| Fase 2 | 8 |
| Fase 3 | 3 |
| Fase 4 | 2 |
| **Total** | **24** |

Strategi urutan implementasi: NFR bukan fitur yang dirilis sekali, melainkan pagar yang dipasang sejak commit pertama dan diperketat per fase. Fase 1 wajib menuntaskan keputusan yang mahal diubah belakangan, yaitu platform database dan disiplin migrasi (NFR-005), kontrol konkurensi dan transaksi outbox (NFR-006, NFR-008), model single-tenant (keputusan awal NFR-014), baseline keamanan (NFR-015), jalur tulis audit `AuditRecorder.Record` beserta tabel `audit_records` (NFR-017; UI, pencarian, dan streaming baru di Fase 3), arsitektur SPA (NFR-020), i18n EN/ID (NFR-021), pengemasan single binary beserta health check dan graceful shutdown (NFR-010, NFR-013), budget performa dasar (NFR-001), dan strategi pengujian berlapis (NFR-024). Fase 2 menambah kemampuan yang dibutuhkan saat konfigurabilitas dan visibilitas tim bertambah: caching, guardrail, semantik waktu lengkap, real-time SSE, observability, enkripsi tingkat aplikasi, aksesibilitas, dan dukungan browser. Fase 3 menangani skala organisasi (target skala dan read replica, DR/HA, UI/pencarian/streaming audit log keamanan, privasi/GDPR), sedangkan Fase 4 membuka mode SaaS multi-tenant dan kesiapan compliance enterprise. Regression gate performa dan dataset referensi 1M work item mulai berjalan di CI paling lambat akhir Fase 2.

## Fitur

## Performa, skala & batas sistem

### NFR-001 Performance Budgets (Issue View, Board, Search)

- **Deskripsi:** Target latency dan throughput terukur untuk jalur kritis harian: membuka work item, render board/backlog, eksekusi JQL, transisi, dan create. Budget ini menjadi acceptance criteria NFR yang diuji di CI dan load test untuk setiap modul fungsional.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering, QA, dan SRE sebagai pemilik budget; diukur dari sisi pengguna mana pun yang memiliki Browse Projects. Tidak ada permission khusus.
- **Aturan perilaku (acceptance criteria):**
  1. Endpoint issue view (GET work item beserta field dan halaman pertama comment) memenuhi p50 < 150 ms dan p95 < 400 ms server-side pada dataset referensi 1M work item (NFR-002 butir 1).
  2. Panel sekunder issue view (Activity: All/Comments/History/Work log, Development, linked work items, child work items) di-load lazy lewat request terpisah setelah konten utama tampil; kegagalan atau lambatnya panel sekunder tidak menahan render ringkasan, deskripsi, dan field.
  3. Pemuatan board dan backlog hingga batas board (NFR-004 butir 2) memenuhi p95 < 1,5 s end-to-end; kolom board dan daftar backlog dirender dengan virtualized list sehingga jumlah node DOM tidak tumbuh linier dengan jumlah kartu.
  4. Pencarian JQL halaman pertama (default 50 hasil) memenuhi p95 < 800 ms untuk query terindeks; query berat (predikat history seperti `WAS`/`CHANGED`, pencarian teks `~`) memenuhi p95 < 3 s dan dihentikan oleh timeout keras dengan pesan galat yang menyarankan mempersempit query.
  5. Transisi work item dan edit field inline memenuhi p95 < 300 ms termasuk evaluasi condition, validator, post function sinkron, penulisan changelog, dan penulisan outbox event.
  6. Create work item (dialog Create maupun inline create di board dan backlog) memenuhi p95 < 400 ms termasuk alokasi issue key.
  7. Quick search, issue picker, user picker, label/version picker, dan JQL autocomplete memenuhi p95 < 150 ms; klien menerapkan debounce dan membatalkan request lama saat pengguna terus mengetik.
  8. Drag-and-drop rank di board dan backlog memenuhi p95 < 250 ms server-side; klien menerapkan optimistic update sehingga kartu langsung berpindah, dan rollback dengan pesan galat bila server menolak.
  9. SPA memenuhi Core Web Vitals LCP < 2,5 s, INP < 200 ms, CLS < 0,1 pada profil perangkat referensi; bundle JavaScript awal < 300 KB gzip dengan code-splitting per route (NFR-020).
  10. Endpoint list ber-pagination (mis. daftar project, filter, dashboard, user) memiliki TTFB < 200 ms untuk halaman pertama.
  11. Setiap request memiliki budget jumlah query DB yang dinyatakan di test; pola N+1 terdeteksi dan menggagalkan test, dan jalur kritis tidak boleh memicu sequential scan pada tabel > 100k baris (diverifikasi dengan `EXPLAIN` di integration test).
  12. CI menjalankan performance regression gate (Go benchmark untuk jalur kritis dan skenario k6) dan menggagalkan build bila p95 memburuk lebih dari 10% dibanding baseline cabang utama.
- **Team-managed vs company-managed:** Tidak ada perbedaan; budget berlaku sama untuk kedua tipe project.
- **Catatan & ketergantungan:** Angka adalah target internal clone, bukan SLA resmi vendor pembanding (vendor tidak mempublikasikan SLA latency per halaman). Budget dasar issue view, board, dan search adalah P0 Fase 1; regression gate dan dataset 1M boleh dimulai di Fase 2. Bergantung pada NFR-002 (dataset), NFR-003 (cache), NFR-020 (SPA), modul 12 dan 13 (eksekusi JQL), modul 10 dan 11 (board/backlog), CMT-023 (changelog).

### NFR-002 Scalability Targets & Capacity Planning

- **Deskripsi:** Batas skala yang harus didukung arsitektur single binary + PostgreSQL, beserta jalur scale-out horizontal ketika beban bertambah.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** SRE dan Architect menyusun target; Admin Situs memakai panduan kapasitas saat memasang instance.
- **Aturan perilaku (acceptance criteria):**
  1. Dataset referensi untuk pengujian dan sizing adalah 1M work item, 10M changelog item, 5M comment, 2M worklog, 500k attachment, 10k user, 2k project, dan 500+ custom field per instance; generator dataset sintetis tersedia di repo (NFR-024 butir 6).
  2. Instance besar mendukung 1.000 pengguna aktif bersamaan, 200 req/s berkelanjutan, dan burst 1.000 req/s tanpa melanggar budget NFR-001.
  3. Aplikasi stateless: tidak ada session atau state lokal di memori yang wajib bertahan antar-request, dan upload sementara tidak ditulis ke disk lokal replika, sehingga N replika dapat berjalan di belakang load balancer tanpa sticky session.
  4. PostgreSQL read replica dapat dipakai untuk search, report, dan gadget dashboard; routing read/write dilakukan di layer repository, dan operasi yang membutuhkan read-your-writes (mis. membaca work item tepat setelah membuatnya) diarahkan ke primary.
  5. Connection pooling (pgxpool atau PgBouncer mode transaction) memakai batas koneksi yang dapat dikonfigurasi; aplikasi tidak bergantung pada prepared statement level sesi yang tidak kompatibel dengan PgBouncer mode transaction.
  6. Tabel volume tinggi (changelog, audit log, webhook delivery, notifikasi, eksekusi otomasi) dipartisi berdasarkan waktu sehingga purge retensi dilakukan dengan drop partisi, bukan DELETE massal.
  7. Project dan work item yang diarsipkan dikeluarkan dari jalur query aktif (index dan view default) agar tidak memperlambat board, backlog, dan pencarian default (ADM-025, ISS-034).
  8. Profil load test k6 terdokumentasi dan disimpan di repo untuk skenario board, navigator, bulk edit, dan create burst.
  9. Capacity model terdokumentasi memberikan estimasi ukuran storage DB per 1M work item dan laju pertumbuhan storage attachment.
  10. Dokumentasi deployment memuat panduan vertical scaling (CPU, RAM, ukuran pool koneksi) per rentang jumlah pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Vendor pembanding hanya mempublikasikan panduan "large instance" untuk edisi Data Center; angka di sini disesuaikan untuk clone self-hosted. Stateless (butir 3) wajib dipatuhi sejak Fase 1 meski target skala penuh baru divalidasi di Fase 3. Bergantung pada NFR-005, NFR-009 (pub/sub antar replika), NFR-011, CMT-019 (object storage).

### NFR-003 Caching Strategy

- **Deskripsi:** Lapisan cache untuk konfigurasi, izin, dan data panas tanpa mengorbankan konsistensi maupun keamanan data.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering. Tidak ada permission pengguna; cache wajib menghormati hasil otorisasi PERM-002.
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi statis (issue type, status, priority, resolution, scheme) di-cache in-process (LRU dengan batas ukuran); setiap perubahan konfigurasi menerbitkan pesan invalidasi lewat pub/sub (NFR-009 butir 4) sehingga semua replika membuang entri usang dalam ≤ 5 detik.
  2. Redis tersedia sejak Fase 1 sebagai dependensi wajib untuk queue (NFR-008, [ADR-0004](../adr/0004-job-latar-dan-queue.md)) dan dapat dipakai untuk cache hasil izin terkomputasi serta counter rate limit. Cache tetap hanya optimasi: hasil harus benar walau entri cache kosong atau hilang, dengan fallback ke cache in-process dan counter berbasis PostgreSQL.
  3. Setiap project memiliki generation number yang dinaikkan saat scheme (permission, workflow, field configuration, issue security) yang terkait berubah; kunci cache menyertakan generation number sehingga seluruh entri project tersebut tidak valid secara massal tanpa scan.
  4. Respons work item dan konfigurasi mendukung ETag/If-None-Match (304 Not Modified); aset SPA ber-hash nama dikirim dengan `Cache-Control: public, max-age=31536000, immutable`, sedangkan `index.html` dikirim dengan `Cache-Control: no-cache`.
  5. Respons yang berisi data work item atau data pengguna dikirim dengan `Cache-Control: private` (atau `no-store` untuk data sensitif) dan tidak boleh di-cache oleh shared cache/CDN.
  6. Saat start, konfigurasi global (status category, issue type, priority, site settings) dimuat ke cache (cache warming) sebelum readiness (NFR-010 butir 2) dilaporkan siap.
  7. Setiap cache mengekspos metrik hit, miss, eviction, dan ukuran per nama cache (NFR-012 butir 3).
- **Team-managed vs company-managed:** Konfigurasi team-managed bersifat per project sehingga cache-nya dikunci per project; company-managed dikunci per scheme bersama dan diinvalidasi lewat generation number semua project yang memakai scheme tersebut.
- **Catatan & ketergantungan:** Compose repo sudah menyertakan Redis 7 dan kini menjadi dependensi wajib untuk queue ([ADR-0004](../adr/0004-job-latar-dan-queue.md)); pemakaian untuk cache bersifat opsional. Bergantung pada PERM-002, NFR-009, NFR-012, VISI-003 (feature flag yang juga di-cache).

### NFR-004 Limits & Guardrails

- **Deskripsi:** Batas sistem eksplisit untuk melindungi performa dan stabilitas, terinspirasi guardrail yang didokumentasikan vendor pembanding; setiap batas memiliki pesan galat yang jelas dan dokumentasi publik.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering menetapkan default; Admin Situs (Administer Jira global permission) dapat menurunkan batas tertentu lewat site settings bila diizinkan; Integrator/Developer API menerima galat batas.
- **Aturan perilaku (acceptance criteria):**
  1. Summary work item maksimal 255 karakter; field teks multi-baris maksimal 32.767 karakter; pelanggaran ditolak dengan HTTP 400 dan pesan yang menyebut nama field dan batasnya.
  2. ⚠ Board menampilkan work item hingga batas jumlah tertentu (default usulan 5.000 per board); bila terlampaui, board menampilkan pesan batas beserta saran mempersempit filter board atau menyembunyikan work item Done lama.
  3. Bulk change via UI maksimal 1.000 work item per operasi; bulk create API maksimal 50 work item per request; rank API maksimal 50 work item per request; permintaan yang melebihi ditolak dengan HTTP 400 sebelum diproses.
  4. ⚠ Parameter `maxResults` pada search API dibatasi 100 per halaman untuk respons dengan field penuh; nilai lebih besar diturunkan diam-diam ke 100 dan nilai efektif dikembalikan di respons; endpoint yang hanya mengembalikan ID dapat memakai batas lebih tinggi yang didokumentasikan terpisah.
  5. ⚠ Setiap work item memiliki guardrail jumlah attachment, comment, worklog, issue link, dan sub-task (default usulan masing-masing 2.000, 5.000, 10.000, 2.000, 1.000); operasi yang melampaui ditolak dengan pesan batas (ISS-044, CMT-016).
  6. Jumlah custom field per instance dan per project, serta jumlah opsi per select field, dibatasi dan ditampilkan sebagai indikator penggunaan di halaman admin field (CFG-039).
  7. Nilai entity property maksimal 32 KB per property, dan jumlah property per entity dibatasi; pelanggaran ditolak dengan HTTP 400.
  8. Otomasi dibatasi per jumlah rule run per bulan per plan, durasi maksimal eksekusi rule, jumlah aksi per rule, serta jumlah work item per lookup/branch (AUTO-015, AUTO-016).
  9. Jumlah webhook per instance dan per app serta ukuran payload webhook dibatasi; payload yang melebihi dipangkas dengan menyertakan tautan ke resource lengkap.
  10. Ukuran maksimal file upload dan jumlah attachment per request dibatasi dan dapat dikonfigurasi admin dalam batas atas sistem (CMT-016).
  11. Kompleksitas JQL dibatasi (jumlah klausa, panjang teks query), dan setiap query dijalankan dengan `statement_timeout` PostgreSQL; query yang melampaui dihentikan dengan pesan galat yang dapat dipahami (JQL-052).
  12. Setiap batas yang tercapai menghasilkan pesan jelas dengan error code stabil (NFR-021 butir 2), dan daftar lengkap batas beserta nilainya dipublikasikan di dokumentasi publik produk dan endpoint baca-saja untuk admin.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada batas sistem; batas jumlah field dan opsi dihitung per project untuk team-managed dan per instance untuk company-managed.
- **Catatan & ketergantungan:** Angka default bertanda ⚠ adalah usulan clone; nilai guardrail resmi vendor pembanding berubah dari waktu ke waktu. Detail per domain ditetapkan di ISS-044, CMT-016, WF-025, JQL-052, SRC-006, AUTO-015, AUTO-016, ADM-041, API-006. Rate limiting per pengguna/token ada di API-006.

## Data, konsistensi & pemrosesan asinkron

### NFR-005 Database Platform & Migration Strategy

- **Deskripsi:** Aturan penggunaan PostgreSQL, GORM, dan golang-migrate agar skema yang besar tetap konsisten, aman di-upgrade, dan dapat dijalankan bersamaan oleh beberapa replika.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering dan DBA; Admin Situs menjalankan migrasi lewat subcommand CLI (NFR-013 butir 7).
- **Aturan perilaku (acceptance criteria):**
  1. PostgreSQL 15 atau lebih baru adalah satu-satunya RDBMS yang didukung; ekstensi wajib `pgcrypto`, `pg_trgm`, `unaccent`, `btree_gin`, dan `citext` dibuat oleh migrasi awal, dan aplikasi menolak start dengan pesan jelas bila versi server lebih rendah atau ekstensi tidak tersedia.
  2. Primary key memakai `uuid` native dengan default `gen_random_uuid()`, atau `BIGINT GENERATED ALWAYS AS IDENTITY` untuk tabel volume tinggi; entitas yang diekspos API ala vendor pembanding memiliki ID numerik publik terpisah (ID work item numerik global) selain issue key string (`PROJ-123`) (modul 28).
  3. Semua kolom waktu bertipe `timestamptz` dan disimpan dalam UTC; tipe `TIMESTAMP` tanpa zona dilarang untuk kolom baru dan kolom lama dimigrasikan.
  4. Perubahan skema hanya dilakukan lewat file migrasi golang-migrate berpasangan up/down; GORM `AutoMigrate` dilarang di produksi dan build gagal bila pemanggilannya ditemukan di luar test.
  5. Perubahan kolom yang tidak kompatibel memakai pola expand/contract (tambah kolom baru → tulis ganda → backfill → pindah baca → hapus kolom lama di rilis berikutnya); migrasi data berat dijalankan sebagai background job (NFR-008), bukan di dalam transaksi migrasi.
  6. Soft delete (`deleted_at`) hanya dipakai pada entitas yang membutuhkan restore (project di Trash, custom field di Trash, opsional filter/dashboard); soft delete global GORM (`gorm.DeletedAt` di model dasar) tidak dipakai untuk entitas lain.
  7. Keunikan pada tabel ber-soft-delete memakai partial unique index yang memperhitungkan `deleted_at` (mis. `CREATE UNIQUE INDEX ... ON users (lower(email)) WHERE deleted_at IS NULL`) sehingga pengguna yang dihapus tidak memblokir pendaftaran ulang email yang sama.
  8. Isolation level default adalah READ COMMITTED; penguncian dan pengecualian isolation diatur di NFR-006.
  9. Konvensi penamaan tabel (snake_case jamak), index (`idx_<tabel>_<kolom>`), unique (`uq_`), foreign key (`fk_`), dan check (`ck_`) didokumentasikan; setiap FK menyatakan `ON DELETE` secara eksplisit (CASCADE, RESTRICT, atau SET NULL) dan pilihannya dicatat per relasi di modul 28.
  10. Seed data default (issue type, priority, resolution, status, status category, permission scheme dan notification scheme default, link type default Blocks, Cloners, Duplicate, Relates) dimasukkan lewat migrasi idempotent yang aman dijalankan ulang dan tidak menimpa data yang telah diubah admin.
  11. CI menjalankan lint migrasi (squawk) yang menggagalkan build untuk lock tabel panjang, `ADD COLUMN ... DEFAULT` dengan ekspresi volatile, dan `CREATE INDEX` tanpa `CONCURRENTLY` pada tabel yang sudah ada.
  12. Migrasi dijalankan di bawah PostgreSQL advisory lock sehingga ketika beberapa replika start bersamaan hanya satu yang mengeksekusi, sementara yang lain menunggu lalu melanjutkan setelah versi skema sesuai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Temuan repo terverifikasi yang wajib dikoreksi di Fase 1: `migrations/001_create_users_table.up.sql` memakai `id VARCHAR(36) PRIMARY KEY DEFAULT VARCHAR(36)` (DEFAULT tidak valid; ganti `uuid DEFAULT gen_random_uuid()`), `TIMESTAMP` tanpa zona, dan `UNIQUE` pada email/username bersama `deleted_at` (bentrok dengan soft delete; ganti partial unique index). docker-compose me-mount `./migrations` ke `/docker-entrypoint-initdb.d` sehingga file `.down.sql` ikut dieksekusi oleh entrypoint Postgres dan bentrok dengan golang-migrate (dihapus di NFR-013 butir 3). Target Makefile `migration-*` menunjuk `./migration/` padahal folder bernama `migrations/`. Skema rinci ada di modul 28.

### NFR-006 Concurrency Control & Data Consistency

- **Deskripsi:** Aturan menangani edit bersamaan, transaksi lintas entitas, dan konsistensi eventual komponen asinkron.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Integrator/Developer API menerima respons 409 Conflict.
- **Aturan perilaku (acceptance criteria):**
  1. Work item, konfigurasi board, dan draft workflow memiliki kolom `version`; update yang membawa versi usang ditolak dengan HTTP 409 Conflict berisi representasi terbaru dan daftar field yang berubah, dan UI menampilkan pilihan "Muat ulang" atau "Timpa" sesuai izin.
  2. Edit inline pada field yang berbeda tidak saling konflik: server menerapkan field-level merge sehingga dua pengguna yang mengubah Assignee dan Priority secara bersamaan sama-sama berhasil; konflik 409 hanya terjadi bila field yang sama diubah dari versi yang usang.
  3. Perubahan entitas, entri changelog (CMT-023), dan outbox event (NFR-008 butir 2) ditulis dalam satu transaksi database; bila salah satu gagal, seluruh perubahan dibatalkan dan tidak ada event yang terkirim.
  4. Alokasi issue key dan counter per project memakai `SELECT ... FOR UPDATE` atau `UPDATE ... RETURNING` pada baris counter sehingga dua create bersamaan tidak pernah menghasilkan key yang sama dan tidak ada celah akibat race.
  5. Operasi global (rebalance rank LexoRank, migrasi scheme, reindex, migrasi DB) memakai advisory lock; percobaan kedua saat lock dipegang ditolak dengan pesan "Operasi sedang berjalan" atau dimasukkan antrean, bukan dijalankan ganda.
  6. Webhook dan otomasi membawa event ID unik; konsumen internal memproses event secara idempoten sehingga pengiriman ulang tidak menghasilkan efek ganda (mis. comment otomasi tidak terduplikasi).
  7. Komponen eventual consistency (indeks pencarian eksternal, notifikasi, otomasi, real-time) didokumentasikan beserta target jeda maksimal (usulan p95 < 5 s); UI tidak mengklaim konsistensi langsung untuk hasil pencarian setelah edit.
  8. Otomasi mencegah loop: rule tidak dipicu oleh event dari rule lain kecuali opsi "Allow rule trigger" diaktifkan, dan loop detection menghentikan rantai yang melampaui batas kedalaman dengan entri audit log rule (modul 20).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Optimistic locking dengan 409 adalah keputusan clone; vendor pembanding umumnya berperilaku last-write-wins per field. Bergantung pada CMT-023, NFR-008, ISS-014, ISS-022, modul 20 dan 28.

### NFR-007 Timezone, Date & Duration Semantics

- **Deskripsi:** Aturan teknis tipe waktu (date, datetime, duration, working time) dan penanganan zona waktu yang benar di penyimpanan, UI, JQL, penjadwalan, dan API.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering; pengguna akhir mengatur zona waktu personal (NAV-033); Admin Situs mengatur zona waktu default sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Sistem memiliki zona waktu default yang diatur admin; setiap pengguna dapat meng-override di preferensi; seluruh datetime disimpan sebagai UTC (`timestamptz`) dan dikonversi ke zona pengguna hanya saat ditampilkan.
  2. Field Date (Due date, Start date, custom date picker) disimpan sebagai tanggal murni (`date`) tanpa zona waktu sehingga tidak bergeser saat dilihat dari zona lain; field DateTime disimpan dengan zona; Duration disimpan dalam detik (integer).
  3. Durasi time tracking dalam notasi `1w 2d 3h 4m` dikonversi memakai konfigurasi hours per day (default 8) dan days per week (default 5) dari TIME-002, dan ditampilkan dalam format pretty, days, atau hours sesuai pengaturan.
  4. Fungsi tanggal JQL (`startOfDay()`, `endOfWeek()`, `now()`, dll.) dan nilai relatif (`-1w`, `2d`) dievaluasi pada zona waktu pengguna yang menjalankan query secara konsisten di navigator, filter, board, dan gadget.
  5. Filter subscription dan scheduled trigger otomasi dijadwalkan pada zona waktu pemiliknya dan aman terhadap DST: jadwal tidak terlewat atau berjalan ganda pada jam yang hilang/berulang.
  6. Sprint end date, burndown, dan perhitungan working days board memperhitungkan non-working days dan zona waktu board (BRD-040, RPT-005).
  7. API menerima dan mengembalikan datetime dalam ISO 8601 dengan offset (mis. `2026-11-05T10:15:00.000+0700`) dan tanggal murni sebagai `YYYY-MM-DD`; input datetime tanpa offset diinterpretasi pada zona waktu pengguna pemanggil (atau zona default sistem untuk pemanggil tanpa preferensi), perilaku ini didokumentasikan di OpenAPI, dan string tanggal yang tidak valid ditolak dengan HTTP 400 yang menyebut nama field.
- **Team-managed vs company-managed:** Tidak ada perbedaan; working days diatur per board di kedua tipe.
- **Catatan & ketergantungan:** Penyimpanan UTC dan tanggal murni (butir 1-2) bersifat P0 karena memengaruhi skema dan wajib sejak Fase 1 bersama NFR-005 butir 3; sisa butir menyusul di Fase 2. Bergantung pada TIME-002, NAV-033, NTF-010, JQL (modul 12), BRD-040.

### NFR-008 Background Jobs & Queue

- **Deskripsi:** Infrastruktur job asinkron untuk pekerjaan berat, terjadwal, dan berbasis event, termasuk transactional outbox yang menjadi tulang punggung notifikasi, webhook, otomasi, indexing, dan real-time.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering dan SRE; Admin Situs (Administer Jira) melihat dan mengulang job gagal; pengguna pemicu operasi bulk melihat progres job miliknya.
- **Aturan perilaku (acceptance criteria):**
  1. Queue berbasis Redis (library `asynq`) dan Redis adalah dependensi wajib sejak Fase 1 ([ADR-0004](../adr/0004-job-latar-dan-queue.md)); aplikasi gagal start bila Redis tidak dikonfigurasi atau tidak terjangkau. Persistensi Redis (AOF) diaktifkan di compose referensi dan dianjurkan Sentinel atau Cluster di produksi.
  2. Setiap perubahan domain menulis baris outbox di PostgreSQL dalam transaksi yang sama (NFR-006 butir 3); relay/dispatcher membaca outbox, meng-enqueue ke Redis, lalu menandai baris terkirim. Pengiriman bersifat at-least-once sehingga handler notifikasi, webhook, otomasi, indexer, dan real-time wajib idempoten.
  3. Jenis job yang didukung minimal: bulk change/move/transition/delete, reindex, import CSV, export besar, migrasi scheme, pengiriman email, polling inbound mail, filter subscription, scheduled trigger otomasi, webhook delivery, scan/thumbnail attachment, rebalance LexoRank, purge trash, penghapusan GDPR, dan backup.
  4. Job gagal di-retry dengan exponential backoff hingga batas attempt per jenis job; setelah batas habis job masuk dead-letter queue dan dapat dilihat, diulang, atau dibuang. Pemilik kemampuan ini adalah NFR-008 sendiri: sejak Fase 1 tersedia endpoint admin terlindungi `/api/v1/admin/jobs` (daftar job per status/jenis, detail galat terakhir, retry, discard) dan halaman admin "Background jobs" sederhana di area Administrasi Sistem (modul 22), keduanya hanya untuk Admin Situs (Administer Jira); retry dan discard tercatat lewat `AuditRecorder.Record` (NFR-017).
  5. Setiap job memiliki idempotency key; enqueue dengan key yang sama selama job sebelumnya belum selesai tidak membuat job ganda.
  6. Antrean dipisah per kelas prioritas (interaktif, misalnya email dan real-time, vs batch, misalnya bulk change dan reindex) dengan kuota worker terpisah sehingga operasi bulk tidak menunda notifikasi.
  7. Job jangka panjang melaporkan progres (persentase, pesan status, hasil per item termasuk item gagal beserta alasan) yang dapat dipantau dari UI, setara halaman progres bulk change vendor pembanding; data progres yang sama kelak diekspos oleh Async Tasks API (API-007, Fase 3).
  8. Scheduler cron terdistribusi menjamin job terjadwal hanya berjalan satu kali per jadwal meski ada banyak replika (mekanisme, misalnya peran scheduler tunggal atau lock terdistribusi, ditetapkan pada tiket implementasi).
  9. Worker berjalan di binary yang sama dalam mode all-in-one (default) atau sebagai proses terpisah dengan flag `--role=worker` (NFR-013 butir 1).
  10. Saat shutdown worker berhenti mengambil job baru, menyelesaikan job aktif hingga batas waktu, lalu mengembalikan job yang belum selesai ke antrean; job yang tertinggal karena crash diambil ulang setelah visibility timeout habis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keputusan teknis: [ADR-0004](../adr/0004-job-latar-dan-queue.md) (Redis/`asynq`, outbox tetap di PostgreSQL); Redis sudah ada di compose repo. Bergantung pada NFR-006, NFR-010 butir 1, NFR-017 (jalur tulis audit), dan modul 28 (katalog event). Dipakai oleh NTF-008, modul 20 (otomasi), modul 23 (webhook), dan API-007 Async Tasks API (Fase 3) yang membaca status job dari tabel yang sama. UI job/DLQ dimiliki fitur ini (butir 4), bukan ADM-057 System info & troubleshooting tools yang berstatus Fase — (Data Center); halaman admin Background jobs Fase 1 dapat diperkaya di Fase 3 (filter lanjut, grafik throughput) tanpa memindahkan kepemilikan.

### NFR-009 Real-time Updates (SSE / WebSocket)

- **Deskripsi:** Penyebaran perubahan langsung ke klien (board, backlog, issue view, notifikasi in-app) agar beberapa pengguna melihat state yang sama tanpa reload.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering; pengguna akhir yang membuka board, backlog, atau issue view menerima event hanya untuk work item yang boleh ia lihat (Browse Projects + issue security, PERM-002).
- **Aturan perilaku (acceptance criteria):**
  1. Server-Sent Events (SSE) adalah transport default karena kompatibel dengan Gin dan proxy HTTP biasa; WebSocket hanya dipakai opsional untuk kolaborasi editor bila dibutuhkan.
  2. Route SSE dikecualikan dari `WriteTimeout` `http.Server` (memakai `http.ResponseController` untuk memperpanjang deadline atau route khusus), dan respons SSE mengirim header `X-Accel-Buffering: no` serta `Cache-Control: no-cache` agar proxy tidak menahan stream.
  3. Channel berbasis resource: `issue:{id}`, `board:{id}`, `project:{id}`, dan `user:{id}` (notifikasi in-app); klien berlangganan sesuai halaman yang terbuka dan berhenti saat berpindah halaman.
  4. Event disebarkan antar-replika lewat Redis Pub/Sub atau PostgreSQL `LISTEN/NOTIFY`; pengguna yang terhubung ke replika mana pun menerima event yang sama.
  5. Event hanya dikirim ke subscriber yang memiliki Browse Projects dan lolos issue security atas work item terkait; perubahan izin di tengah sesi berlaku untuk event berikutnya.
  6. Payload event ringan (jenis event, ID resource, `version`); klien melakukan refetch lewat API biasa sehingga isi field tidak pernah bocor lewat stream.
  7. Klien reconnect otomatis dengan header `Last-Event-ID` dan backoff eksponensial; server mengirim heartbeat setiap ±25 detik agar koneksi tidak diputus idle timeout proxy.
  8. ⚠ Issue view menampilkan indikator "Work item ini telah diperbarui" dengan tombol muat ulang ketika event versi baru diterima sementara pengguna sedang mengedit, dan menyerahkan deteksi konflik ke NFR-006 butir 1.
  9. Jumlah koneksi stream dibatasi per pengguna/tab (usulan 6 tab; tab berlebih berbagi koneksi lewat BroadcastChannel atau jatuh ke polling) dan per instance; koneksi yang melampaui ditolak dengan HTTP 429.
  10. Autentikasi stream memakai cookie sesi atau token sekali pakai berumur pendek (usulan 60 detik) di query string, karena `EventSource` tidak mendukung header `Authorization`; token tidak dicatat di access log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Vendor pembanding memperbarui sebagian board dan issue view secara real-time; detail perilaku paritasnya belum dipastikan sehingga butir 8 ditandai ⚠. Presence indicator (siapa yang sedang melihat) sengaja tidak dimasukkan karena tidak terkonfirmasi sebagai fitur native. Bergantung pada NFR-008 (outbox), NFR-010 butir 1 (shutdown stream), ISS-014, NAV-046, PERM-002.

## Keandalan, operasi & deployment

### NFR-010 Availability, Graceful Shutdown & Health Checks

- **Deskripsi:** Ketersediaan layanan, shutdown yang aman untuk HTTP, stream real-time, dan worker, serta endpoint health untuk orchestrator (Docker/Kubernetes).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** SRE dan Admin Situs; endpoint health tidak memerlukan autentikasi; mode maintenance diaktifkan oleh Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Saat menerima SIGTERM/SIGINT, server berhenti menerima koneksi baru dan menyelesaikan request in-flight dalam batas waktu yang dapat dikonfigurasi (default 30 detik).
  2. Karena `http.Server.Shutdown` tidak menutup koneksi long-lived/hijacked, server membatalkan context setiap stream SSE/WebSocket dan mengirim event `reconnect` sebelum menutupnya agar klien pindah ke replika lain.
  3. Worker berhenti setelah job aktif selesai atau, bila melewati batas waktu, mengembalikan job ke antrean (NFR-008 butir 10); proses keluar dengan kode 0 setelah semua komponen berhenti.
  4. `/healthz` (liveness) hanya memastikan proses hidup tanpa memeriksa dependensi; `/readyz` (readiness) memeriksa DB, cache, object storage, dan queue serta mengembalikan 503 bila salah satu dependensi wajib gagal.
  5. Kedua endpoint health tidak memerlukan auth, tidak memakai prefix `/api/v1`, tidak dialihkan ke fallback SPA, dan hanya mengembalikan status ringkas tanpa membocorkan host, versi dependensi, atau pesan galat internal.
  6. Healthcheck image Docker memakai subcommand binary (`app healthcheck`) karena image distroless tidak memuat `curl`/`wget`.
  7. Deploy tanpa downtime dilakukan lewat rolling update; migrasi DB wajib backward compatible (expand/contract, NFR-005 butir 5) sehingga versi N dan N+1 dapat berjalan bersamaan selama rollout.
  8. Setiap dependensi eksternal (SMTP, object storage S3, target webhook, OpenSearch) dipanggil dengan timeout eksplisit dan circuit breaker; circuit terbuka tercatat di log dan metrik.
  9. Mode degraded: bila search engine eksternal tidak tersedia, pencarian jatuh ke PostgreSQL FTS atau menampilkan pesan galat yang jelas (SRC-001); bila SMTP tidak tersedia, email tetap berada di antrean dan dikirim ulang setelah pulih.
  10. Admin dapat mengaktifkan mode read-only/maintenance yang menolak operasi tulis dengan HTTP 503 dan pesan maintenance, sementara banner ditampilkan ke seluruh pengguna; mode ini dipakai untuk upgrade dan restore (konsep dari edisi Data Center vendor pembanding).
  11. Target availability bulanan untuk deployment HA (NFR-011) adalah 99,9% (pembanding: SLA vendor pada plan Premium 99,9%, Enterprise 99,95%); deployment single-node tidak dijanjikan target ini.
  12. Status page per komponen (web, API, email, automation, webhook) bersifat opsional dan dapat dihubungkan ke layanan status pihak ketiga.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Repo sudah memiliki graceful shutdown HTTP di `cmd/api/main.go`, tetapi memanggil `server.Shutdown(context.Background())` tanpa batas waktu dan belum menangani stream maupun worker; Fase 1 menambahkan timeout terkonfigurasi. Target 99,9% dan status page adalah P1/P2; health check dan graceful shutdown P0. RTO/RPO diatur di NFR-011.

### NFR-011 Disaster Recovery & High Availability Topology

- **Deskripsi:** Topologi infrastruktur dan target pemulihan agar instance tahan terhadap kegagalan node, zona, atau region.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** SRE dan Admin Situs (operator deployment).
- **Aturan perilaku (acceptance criteria):**
  1. Target RPO ≤ 15 menit dan RTO ≤ 4 jam terdokumentasi untuk topologi HA referensi, beserta asumsi yang membuat target tersebut berlaku.
  2. Topologi referensi memakai PostgreSQL primary + standby dengan streaming replication dan failover otomatis (Patroni atau managed database); aplikasi pulih otomatis setelah failover tanpa restart manual (koneksi pool di-reset).
  3. Aplikasi berjalan multi-replika tanpa sticky session; klien SSE yang terputus reconnect ke node mana pun dan melanjutkan dari `Last-Event-ID` (NFR-009 butir 7).
  4. Object storage memakai replikasi lintas zona, dan backup database serta attachment disalin ke lokasi offsite/lintas region dengan enkripsi (ADM-035).
  5. Runbook insiden dan failover tertulis tersedia di dokumentasi operator, mencakup langkah deteksi, failover, verifikasi, dan failback.
  6. Uji DR (restore backup ke lingkungan terpisah dan failover database) dijalankan minimal sekali per kuartal, dan hasilnya (waktu pemulihan aktual vs target) dicatat sebagai bukti (NFR-019 butir 3).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada NFR-002 butir 3 (stateless), NFR-009, NFR-010, ADM-035 (backup & restore), CMT-019 (object storage).

### NFR-012 Observability (Logs, Metrics, Traces)

- **Deskripsi:** Kemampuan memantau kesehatan dan performa sistem di produksi lewat log terstruktur, metrik, trace, dan error tracking.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** SRE dan Engineering; endpoint metrik dilindungi (jaringan internal atau token) dan tidak tersedia untuk pengguna biasa.
- **Aturan perilaku (acceptance criteria):**
  1. Log aplikasi berformat JSON terstruktur dengan field minimal `request_id`, `user_id` (di-hash), `route` (template, bukan URL mentah), `latency_ms`, dan `status`; isi work item, token, dan password tidak pernah dicatat (NFR-018 butir 9).
  2. Correlation ID (header `X-Request-Id`, dibuat bila tidak ada) dipropagasi ke job, webhook keluar, dan email (header kustom) sehingga satu aksi pengguna dapat ditelusuri ujung ke ujung.
  3. Endpoint metrik Prometheus mengekspos RED per endpoint (rate, errors, duration), durasi query DB, penggunaan pool koneksi, kedalaman antrean, kegagalan job, jumlah koneksi SSE, laju pengiriman email/webhook, serta hit/miss cache.
  4. Distributed tracing OpenTelemetry mencakup Gin, GORM, HTTP client keluar, dan eksekusi job, dengan exporter OTLP yang dapat dikonfigurasi dan sampling rate yang dapat diatur.
  5. Repo menyediakan dashboard Grafana dan aturan alert referensi untuk error rate, p95 latency, backlog antrean, ruang disk, dan replication lag.
  6. Query DB yang melewati ambang (usulan 500 ms) dan JQL lambat dicatat di slow query log dengan teks literal ter-redaksi (nilai string/angka diganti placeholder).
  7. Error backend dan SPA dikirim ke error tracker yang kompatibel dengan protokol Sentry bila DSN dikonfigurasi; data pribadi di-scrub sebelum dikirim.
  8. Log level dapat diubah saat runtime (lewat endpoint admin terlindungi atau sinyal) tanpa restart, dan perubahan level tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Repo saat ini memakai logrus dengan formatter JSON kustom berwarna (`pkg/logger`); kode warna ANSI harus dimatikan di produksi atau logger dimigrasikan ke `log/slog`. Halaman System info versi Data Center (ADM-057) berstatus Fase — dan hanya rujukan; informasi runtime yang dibangun disediakan lewat endpoint metrik dan log modul ini, sedangkan UI job/DLQ dimiliki NFR-008 butir 4.

### NFR-013 Deployment, Packaging & Configuration

- **Deskripsi:** Cara aplikasi dibangun, dikemas, dan dikonfigurasi sebagai single binary Go dengan SPA ter-embed, image container minimal, dan konfigurasi 12-factor.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** SRE dan Engineering; Admin Situs mengubah site settings runtime lewat UI admin (modul 22).
- **Aturan perilaku (acceptance criteria):**
  1. Satu binary Go memuat API, SPA (`go:embed` dari `internal/shared/delivery/http/webui/dist`), dan worker; flag `--role=api|worker|all` (default `all`) memilih komponen yang dijalankan; handler `NoRoute` SPA tetap didaftarkan terakhir di `app.go`.
  2. Image Docker dibangun multi-stage dengan image final minimal (distroless), berjalan sebagai user non-root, kompatibel dengan filesystem read-only, dan mendefinisikan healthcheck (NFR-010 butir 6).
  3. docker-compose referensi berisi layanan app, postgres, redis, minio, dan mailpit; mount `./migrations` ke `/docker-entrypoint-initdb.d` dihapus (migrasi lewat golang-migrate), dan sisa nama `yopatungan` diganti nama netral produk.
  4. Helm chart atau manifest Kubernetes opsional disediakan dengan HorizontalPodAutoscaler, probe liveness/readiness, dan pemisahan deployment api/worker.
  5. Konfigurasi infrastruktur dibaca dari environment variable bergaya 12-factor (`DB_*`, `JWT_SECRET`, `SERVER_*`, `STORAGE_*`, `SMTP_*`, `SEARCH_*`, `REDIS_*`) dan divalidasi saat start; nilai wajib yang kosong atau tidak valid (mis. `JWT_SECRET` kosong atau < 32 byte di produksi) menghentikan proses dengan pesan yang menyebut nama variabel (fail-fast).
  6. Site settings runtime (judul situs, mode visibility, batas attachment, dsb.) disimpan di database dan diubah admin lewat UI, bukan lewat env; perubahan berlaku tanpa restart.
  7. Binary menyediakan subcommand CLI `migrate`, `reindex`, `create-admin`, `backup`, dan `healthcheck` dengan kode keluar non-nol saat gagal.
  8. Rilis memakai versi semantik dan changelog rilis; endpoint `serverInfo` mengembalikan versi, build number, commit, waktu build, dan waktu server (API-002).
  9. Migrasi otomatis saat start bersifat opsional (flag/env, default nonaktif di produksi) dan selalu berjalan di bawah advisory lock (NFR-005 butir 12).
  10. Pipeline rilis menghasilkan SBOM dan menandatangani image dengan cosign; signature dan SBOM dipublikasikan bersama image.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Temuan repo terverifikasi: docker-compose memakai nama `yopatungan` dan me-mount migrations ke initdb; compose memakai `DB_PASSWORD` sementara config membaca `DB_PASS`; Dockerfile final memakai `alpine:latest` dan berjalan sebagai root di `/root/`; `internal/core/config` memberi default untuk semua variabel (termasuk password DB) sehingga belum fail-fast. Penamaan produk mengikuti VISI-001.

### NFR-014 Multi-tenancy Model

- **Deskripsi:** Keputusan arsitektur apakah satu deployment melayani satu organisasi (single-tenant) atau banyak tenant (site), beserta aturan isolasi untuk mode SaaS dan profil deployment terisolasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Architect dan SRE; operator SaaS (super-admin di luar Admin Situs tenant) mengelola tenant.
- **Aturan perilaku (acceptance criteria):**
  1. Mode default adalah single-tenant per deployment: satu site = satu database; tidak ada kolom tenant pada skema Fase 1-3.
  2. Mode SaaS opsional menambahkan `tenant_id` pada semua tabel bisnis, dan PostgreSQL Row Level Security diaktifkan sebagai pertahanan lapis kedua di samping filter tenant di layer repository.
  3. Tenant diresolusi dari subdomain (pembanding pola `{site}.atlassian.net` di vendor), dengan fallback header untuk lingkungan pengujian; request tanpa tenant valid mendapat HTTP 404.
  4. Storage diisolasi per tenant (prefix bucket), kunci cache diawali tenant, dan rate limit dihitung per tenant.
  5. Setiap tenant memiliki konfigurasi dan plan tier sendiri (VISI-003).
  6. Batas per tenant (jumlah user, storage, jumlah run otomasi) dapat dikonfigurasi dan ditegakkan dengan pesan batas (NFR-004).
  7. Satu tenant dapat diekspor dan dimigrasikan ke deployment lain tanpa menyertakan data tenant lain.
  8. Test kebocoran lintas tenant otomatis dijalankan di CI: setiap endpoint dipanggil dengan kredensial tenant A terhadap ID resource tenant B dan wajib menghasilkan 404.
  9. ⚠ Varian deployment vendor (Government Cloud ber-otorisasi FedRAMP Moderate dengan cakupan fitur terbatas, dan Isolated Cloud single-tenant terisolasi per pelanggan) dipenuhi clone lewat mode single-tenant self-hosted/air-gapped tanpa dependensi layanan eksternal wajib.
  10. Profil deployment terisolasi: semua dependensi (DB, object storage, SMTP, search) dapat di-host di jaringan pelanggan, tidak ada telemetry/phone-home secara default, dan seluruh egress dapat dimatikan tanpa merusak fungsi inti.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Vendor pembanding multi-tenant dengan hierarki organization → site → product. Keputusan single-tenant untuk MVP diambil di Fase 1 karena memengaruhi skema; mode SaaS (butir 2-8) adalah pekerjaan Fase 4. Government Cloud dan Isolated Cloud adalah penawaran deployment vendor, bukan fitur fungsional, dan status ketersediaannya belum dipastikan. Bergantung pada VISI-003, USR-042 (data residency), NFR-003, NFR-017.

## Keamanan, privasi & kepatuhan

### NFR-015 Application Security Baseline (OWASP ASVS)

- **Deskripsi:** Standar keamanan aplikasi berdasarkan OWASP ASVS Level 2 untuk seluruh API dan SPA, mencakup validasi input, header keamanan, CSRF, CORS, proteksi IDOR, dan pemindaian otomatis.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering dan Security; Admin Situs mengatur allowlist CORS dan frame-ancestors lewat konfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Target verifikasi adalah OWASP ASVS Level 2; checklist ASVS dipetakan ke test otomatis atau bukti manual, dan butir yang belum terpenuhi tercatat sebagai issue keamanan terbuka.
  2. Semua DTO request divalidasi di server dengan go-playground/validator (panjang, enum, format, rentang); validasi di SPA tidak pernah menggantikan validasi server, dan galat validasi dikembalikan sebagai HTTP 400 dengan field yang bermasalah.
  3. Content-Security-Policy ketat dikirim untuk SPA: `script-src 'self'` (tanpa inline script, atau memakai hash bila tidak terhindarkan), `object-src 'none'`, `base-uri 'self'`, dan `frame-ancestors` terbatas pada allowlist.
  4. Header keamanan wajib: `Strict-Transport-Security` (saat HTTPS), `X-Frame-Options`/`frame-ancestors`, `Referrer-Policy: strict-origin-when-cross-origin`, `Permissions-Policy` minimal, dan `X-Content-Type-Options: nosniff`.
  5. Bila sesi memakai cookie, setiap request yang mengubah state dilindungi CSRF (cookie `SameSite=Lax/Strict` plus token double-submit atau header kustom wajib, setara header `X-Atlassian-Token: no-check` vendor untuk endpoint multipart); request dengan Bearer token di header `Authorization` tidak memerlukan token CSRF.
  6. CORS memakai allowlist origin eksplisit dari konfigurasi; wildcard `*` tidak diizinkan bersama kredensial, dan origin di luar allowlist tidak menerima header `Access-Control-Allow-Origin`.
  7. Setiap akses resource (work item, attachment, comment, worklog, filter, dashboard, board) melewati Authorizer terpusat (PERM-002); test otomatis memastikan akses dengan ID milik project lain ditolak (proteksi IDOR).
  8. Respons untuk resource yang tidak boleh dilihat identik dengan resource yang tidak ada (HTTP 404 seragam dengan pesan "Work item tidak ada atau Anda tidak memiliki izin untuk melihatnya" / "Issue does not exist or you do not have permission to see it"), sehingga keberadaan work item terlarang tidak bocor.
  9. CI menjalankan dependency scanning (govulncheck, npm/bun audit), SAST (gosec, semgrep), dan container scan (trivy); temuan severity High/Critical menggagalkan build kecuali dikecualikan dengan justifikasi tertulis.
  10. Penetration test berkala (minimal sebelum rilis mayor) dilakukan, dan `/.well-known/security.txt` serta kebijakan pelaporan kerentanan dipublikasikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** CSP berbasis nonce tidak cocok untuk SPA statis ter-embed sehingga dipakai `'self'`/hash. Temuan repo: `middleware/cors.go` mengirim `Access-Control-Allow-Origin: *` untuk semua request dan wajib diganti allowlist di Fase 1. Rendering rich text aman ada di EDT-002, keamanan file di CMT-017, SSRF guard di INT-001, rate limiting di API-006, sesi dan 2FA di AUTH-007.

### NFR-016 Secrets Management & Encryption

- **Deskripsi:** Pengelolaan rahasia aplikasi serta enkripsi data saat transit dan saat disimpan, termasuk enkripsi tingkat aplikasi untuk kolom sensitif dan rotasi kunci.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** SRE dan Security; Admin Situs tidak dapat membaca kembali secret yang sudah disimpan lewat UI (hanya mengganti).
- **Aturan perilaku (acceptance criteria):**
  1. Secret dibaca dari environment variable atau file mount (Docker secrets, Kubernetes Secret, Vault agent); secret tidak pernah di-commit ke repo, `.env.example` hanya berisi placeholder, dan CI menjalankan secret scanning.
  2. Koneksi eksternal memakai TLS 1.2 atau lebih baru; koneksi ke PostgreSQL di produksi memakai `sslmode=verify-full`, dan aplikasi memperingatkan di log start bila mode produksi memakai `sslmode=disable`.
  3. Enkripsi at-rest dipenuhi lewat enkripsi volume database dan server-side encryption pada object storage S3; persyaratan ini tercantum di panduan deployment.
  4. Kolom sensitif (password SMTP/IMAP, secret webhook, OAuth client secret, TOTP secret) dienkripsi di tingkat aplikasi dengan AES-GCM memakai envelope key; nilai tidak pernah dikembalikan utuh lewat API (hanya penanda "tersimpan").
  5. Kunci enkripsi memiliki versi; rotasi kunci menambahkan versi baru, data lama tetap dapat didekripsi dengan versi lamanya, dan job re-enkripsi (NFR-008) memindahkan data ke versi terbaru.
  6. Secret, token, password, dan header `Authorization` di-mask di log, trace, error tracking, dan audit log (NFR-012, NFR-017).
  7. ⚠ BYOK (bring your own key) lewat KMS eksternal disediakan sebagai opsi enterprise untuk envelope key, setara penawaran vendor pada plan Enterprise (USR-043).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Secret via env dan TLS dasar (butir 1-2) adalah P0 dan wajib di Fase 1; enkripsi kolom (butir 4-5) Fase 2; BYOK P2 di Fase 4. Repo saat ini mengirim `DB_SSLMODE=disable` di compose dan `JWT_SECRET` contoh di compose; nilai tersebut hanya untuk pengembangan lokal. Bergantung pada USR-043, NFR-008.

### NFR-017 Security Audit Logging

- **Deskripsi:** Log audit append-only untuk tindakan administratif dan keamanan, termasuk skema penyimpanan, kategori, deteksi manipulasi, pencarian, retensi, dan streaming ke SIEM. Fitur dibagi dua jalur: jalur tulis (`AuditRecorder.Record` dan tabel `audit_records`) wajib sejak Fase 1, sedangkan jalur baca (UI, pencarian, export, streaming) menyusul di Fase 3.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform (jalur tulis; UI, pencarian, dan streaming P1 / Fase 3)
- **Aktor & izin:** Admin Situs (Administer Jira global permission) membaca dan mengekspor audit log; Security mengonsumsi stream SIEM; tidak ada pengguna yang dapat mengubah atau menghapus entri.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel `audit_records` menyimpan id, `scope` (`app`, `org`, atau `automation`), waktu UTC, actor (user atau sistem/app), alamat IP, user agent, category, summary, object (type, id, name), associated items, changed values (nilai lama dan baru), serta hasil (berhasil/gagal); tabel bersifat append-only (tanpa UPDATE/DELETE dari aplikasi kecuali job retensi pada butir 5) dan dipartisi per bulan. Semua modul menulis hanya lewat use case `AuditRecorder.Record(ctx, entry)` yang dipanggil di dalam transaksi domain yang sama (NFR-006), sehingga aksi yang gagal di-rollback tidak meninggalkan entri dan aksi yang berhasil selalu memiliki entri; jalur tulis ini tersedia sejak Fase 1.
  2. Kategori yang dicatat minimal: manajemen user, group, permission & scheme, workflow, field, screen, create/delete/archive/restore project, login/logout dan kegagalan login, pembuatan/pencabutan token, pengaturan global, serta ekspor data.
  3. Hash chain opsional (setiap entri menyimpan hash entri sebelumnya) dapat diaktifkan, dan alat verifikasi melaporkan entri pertama yang rantainya putus.
  4. UI admin menyediakan pencarian teks dan filter (rentang waktu, kategori, actor, object) dengan paginasi, serta export CSV dari hasil filter; export itu sendiri tercatat di audit log.
  5. ⚠ Retensi dapat dikonfigurasi per scope dan menjadi satu-satunya acuan angka retensi bagi ADM-014, USR-045, USR-046, dan AUTO-014:

     | Scope | Fitur UI | Default | Pilihan | Maksimum per tier |
     |---|---|---|---|---|
     | `app` (audit log Jira/admin produk) | ADM-014, USR-046 | 90 hari | 30, 90, 180, 365 hari | Free/Standard 90 hari; Premium/Enterprise 365 hari |
     | `org` (audit log organisasi lintas app) | USR-045 | 180 hari | 180 hari hingga 7 tahun | 180 hari tanpa paket keamanan lanjutan; hingga 7 tahun dengan paket Premium/keamanan lanjutan |
     | `automation` (audit log rule otomasi) | AUTO-014 | 90 hari | 30, 90, 180, 365 hari | Sama dengan scope `app` |

     Job retensi harian (NFR-008) men-drop partisi bulanan yang seluruh isinya sudah melewati retensi dan menghapus berbatch entri kedaluwarsa di partisi yang masih bercampur; entri di bawah legal hold (NFR-019 butir 6) dikecualikan. Sejak Fase 1 job retensi berjalan dengan nilai default; pengaturan retensi lewat UI baru tersedia di Fase 3. Pada deployment self-hosted, operator dapat menaikkan batas maksimum lewat konfigurasi (NFR-013).
  6. Streaming ke SIEM (syslog RFC 5424 atau HTTP JSON) bersifat opsional dengan retry dan buffer saat target tidak tersedia.
  7. Akses data massal (ekspor, pemanggilan search API dengan hasil besar, ekspor CSV navigator di atas ambang) dicatat sebagai entri audit kategori data export.
- **Team-managed vs company-managed:** Perubahan konfigurasi team-managed (field, workflow, akses project) tercatat dengan object project terkait; company-managed tercatat pada object scheme global.
- **Catatan & ketergantungan:** Skema `audit_records` digabung dari item data model audit (`audit_log`) di inventaris; nama tabel diseragamkan dengan ADM-014 dan USR-046. **Pemisahan fase:** Fase 1 (P0) mencakup butir 1 (skema `audit_records` dan `AuditRecorder.Record`), butir 2 (kategori), butir 7 (pencatatan akses data massal), dan job retensi dengan nilai default dari butir 5, sehingga fitur Fase 1 yang mensyaratkan aksi "tercatat di audit log" (mis. ISS-033, CMT-003, PERM-006, PERM-007, PRJ-016, PRJ-039, USR-004, USR-008, ADM-002, ADM-003, NTF-024) dapat dipenuhi sejak awal walau belum ada UI. Fase 3 (P1) mencakup butir 3 (hash chain dan alat verifikasi), butir 4 (UI pencarian, filter, export CSV), pengaturan retensi lewat UI di butir 5, butir 6 (streaming SIEM), serta UI dan API di ADM-014, USR-045, USR-046, dan AUTO-014. **Rujukan silang audit log:** UI audit log app di ADM-014 (Jira settings) dan USR-046 (Jira audit log) memakai scope `app`; audit log organisasi USR-045 memakai scope `org`; audit log rule otomasi AUTO-014 memakai scope `automation` dan menautkan entri perubahan rule ke ADM-014; audit workflow di WF-039 ditulis lewat `AuditRecorder.Record`. Keempat fitur UI tersebut wajib memakai tabel retensi butir 5 dan tidak menetapkan angka sendiri. Retensi dan legal hold dikoordinasikan dengan NFR-019 butir 6 dan ADM-025.

### NFR-018 Privacy & GDPR Compliance

- **Deskripsi:** Kepatuhan perlindungan data pribadi (GDPR dan UU PDP Indonesia No. 27/2022) termasuk visibilitas profil, hak akses dan penghapusan, retensi, consent, dan prosedur pelaporan pelanggaran.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Admin Situs dan Legal; pengguna akhir menjalankan hak subjek data atas akunnya sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna mengatur visibilitas profil (siapa yang dapat melihat email, job title, departemen, zona waktu, dsb.) dengan opsi minimal "Siapa saja", "Organisasi", dan "Hanya saya", dan API serta UI menghormati pengaturan ini (AUTH-002).
  2. Pengguna dapat mengekspor data pribadinya dalam format JSON secara self-service, dan Admin Situs dapat melakukannya atas nama pengguna; ekspor dijalankan sebagai job (NFR-008) dan tercatat di audit log.
  3. Penghapusan akun menganonimkan pengguna menjadi "Former user" sambil mempertahankan account ID agar referensi historis tetap utuh, dengan grace period 14 hari sebelum anonimisasi final yang dapat dibatalkan selama periode tersebut (AUTH-014, USR-044).
  4. PII di konten bebas (mention, comment, deskripsi) menjadi tanggung jawab admin; produk menyediakan alat pencarian konten berdasarkan nama/email pengguna untuk membantu penghapusan manual.
  5. Retensi dan purge otomatis berjalan untuk trash (60 hari), log, delivery webhook, dan notifikasi sesuai periode yang dapat dikonfigurasi (ADM-025).
  6. Analytics produk hanya aktif setelah consent lewat cookie notice, dan Admin Situs dapat mematikan analytics sepenuhnya; tanpa consent tidak ada cookie non-esensial yang ditulis.
  7. Data residency dipenuhi oleh lokasi deployment clone (pembanding: pilihan region data residency vendor pada plan Standard ke atas, USR-042) dan didokumentasikan di panduan deployment.
  8. Template DPA, daftar subprocessor (untuk deployment hosted), dan kebijakan privasi produk tersedia di dokumentasi.
  9. Log menerapkan minimasi data: tidak mencatat isi work item, token, atau password, dan identitas pengguna di log memakai hash (NFR-012 butir 1).
  10. Prosedur pelaporan breach terdokumentasi untuk memberi notifikasi kepada otoritas dan subjek data dalam 72 jam setelah insiden diketahui.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** UU PDP No. 27/2022 relevan untuk pengguna Indonesia. Bergantung pada AUTH-002, AUTH-014, USR-042, USR-044, ADM-025, NFR-017.

### NFR-019 Compliance Controls & Certifications Readiness

- **Deskripsi:** Pemetaan kontrol teknis clone terhadap standar compliance yang dipenuhi vendor pembanding (SOC 2 Type II, ISO/IEC 27001, ISO/IEC 27018, CSA STAR, HIPAA pada Enterprise, FedRAMP Moderate di Government Cloud) agar operator dapat menyiapkan bukti audit sendiri.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Security, Admin Situs, Legal, dan SRE; laporan dan ekspor bukti hanya untuk pemegang Administer Jira.
- **Aturan perilaku (acceptance criteria):**
  1. Dokumentasi memuat matriks kontrol yang memetakan kontrol SOC 2 / ISO 27001 / ISO 27018 ke fitur clone (audit log, enkripsi at-rest dan in-transit, access review, retensi log, backup teruji, change management, vulnerability management) beserta ID NFR terkait.
  2. Laporan access review dapat diekspor per periode dan memuat user, group, akses produk, global permission, dan project role untuk peninjauan berkala.
  3. Ekspor bukti otomatis menghasilkan paket berisi konfigurasi keamanan aktif, riwayat perubahan konfigurasi admin, status backup dan hasil restore test, serta hasil scan dependensi terakhir.
  4. ⚠ Profil HIPAA-ready dapat diaktifkan dan mewajibkan enforcement SSO/2FA, enkripsi kolom sensitif, audit log selalu aktif, retensi minimum, serta larangan PHI di notifikasi email (memaksa mode email tanpa konten work item); syarat pembanding vendor (BAA, data residency, langganan Guard) dicatat sebagai rujukan.
  5. Admin dapat mengaktifkan opsi notifikasi email tanpa detail konten (hanya judul generik dan tautan) untuk lingkungan teregulasi (NTF-023).
  6. Kebijakan retensi dan legal hold dapat dikonfigurasi per kategori data (log, audit, attachment); data di bawah legal hold dikecualikan dari purge otomatis (selaras ADM-025).
  7. Dokumentasi shared responsibility model membedakan kontrol yang menjadi tanggung jawab aplikasi dan tanggung jawab operator deployment self-hosted.
  8. PCI DSS dinyatakan di luar cakupan karena clone tidak memproses pembayaran; produk tidak menyimpan data kartu dalam bentuk apa pun.
  9. DPA dan daftar subprocessor dirujuk ke NFR-018 butir 8 dan tidak diduplikasi di modul ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sertifikasi adalah atestasi korporat vendor pembanding, bukan fitur UI; clone open-source tidak otomatis bersertifikat. Fitur ini hanya menyediakan kontrol dan bukti agar operator dapat mengejar sertifikasi sendiri. Daftar sertifikasi per plan vendor berubah dan harus dicek ulang di trust center vendor. Bergantung pada NFR-011, NFR-016, NFR-017, NFR-018, USR-037 (Guard).

## Klien web, bahasa & aksesibilitas

### NFR-020 Frontend Architecture Constraints (SvelteKit SPA)

- **Deskripsi:** Batasan arsitektur SPA SvelteKit 5 yang dibangun statis dan di-embed ke binary Go, mencakup routing, pemuatan kode, cache data klien, penyimpanan token, deep link, dan tema.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering frontend; tidak ada permission pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. SPA dibangun dengan static adapter; route tidak dikenal untuk GET di-fallback ke `index.html`, sedangkan `/api/*` dan `/swagger/*` selalu mengembalikan JSON 404 dan tidak pernah dilayani oleh fallback SPA.
  2. Kode dipecah per route, dan komponen berat (rich text editor, diagram workflow, chart laporan) di-lazy-load hanya saat dibutuhkan agar memenuhi budget bundle NFR-001 butir 9.
  3. Data server dikelola lewat query cache stale-while-revalidate; aksi interaktif (drag kartu, transisi, edit inline) memakai optimistic update dengan rollback dan pesan galat bila server menolak.
  4. Daftar besar (backlog, navigator, kolom board) dirender dengan virtualisasi.
  5. Token akses berumur panjang tidak disimpan di `localStorage`; SPA memakai cookie `httpOnly` `Secure` atau token di memori dengan mekanisme refresh, selaras AUTH-007 dan AUTH-008.
  6. State penting dapat dibagikan lewat URL deep link setara pola vendor pembanding: `/browse/KEY`, `?selectedIssue=KEY` di board/backlog, `?jql=` di navigator, dan quick filter board aktif; membuka URL tersebut memulihkan tampilan yang sama.
  7. Setiap route memiliki error boundary yang menampilkan pesan dan opsi coba lagi tanpa merusak shell aplikasi, dan indikator offline tampil saat koneksi jaringan terputus.
  8. Design token (warna, spasi, tipografi) dipakai konsisten di seluruh komponen dan mendukung tema light, dark, dan mengikuti sistem (NAV-034).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Drag-and-drop dibangun sendiri sebagai modul reusable `frontend/src/lib/dnd/` ([ADR-0001](../adr/0001-library-drag-and-drop-board.md)). Lihat `frontend/CLAUDE.md`; mekanisme embed dan `NoRoute` sudah ada di repo (`webui.RegisterRoutes`), dan tanpa build frontend server menjawab 404 "frontend is not built". Bergantung pada NFR-001, NFR-003 butir 4, NFR-015 butir 3, AUTH-007, NAV-034.

### NFR-021 Internationalization (i18n ID/EN)

- **Deskripsi:** Dukungan bahasa Indonesia dan Inggris di UI, email, dan pesan galat API, termasuk format locale, pluralisasi, dan terjemahan nama objek konfigurasi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Pengguna akhir memilih bahasa di preferensi (NAV-033); Engineering mengelola katalog terjemahan; Admin Situs mengelola terjemahan objek konfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Semua string UI berada di file terjemahan (EN sebagai default dan fallback, ID); tidak ada string pengguna yang di-hardcode di komponen, dan CI menggagalkan build bila ada kunci terjemahan yang hilang di salah satu bahasa.
  2. Bahasa yang dipakai ditentukan berurutan: preferensi bahasa pengguna yang tersimpan, lalu header `Accept-Language`, lalu default sistem (memperluas `middleware.LanguageMiddleware` repo).
  3. Pesan galat API multi-bahasa memakai `constants.GetError`/`GetErrorMessage` dan selalu menyertakan error code stabil berbahasa Inggris (mis. `ISSUE_NOT_FOUND`) yang tidak pernah diterjemahkan agar integrasi dapat bergantung padanya.
  4. Email notifikasi dirender dalam bahasa penerima, bukan bahasa pelaku aksi (NTF-010).
  5. Tanggal, angka, dan durasi diformat sesuai locale pengguna (mis. `5 Nov 2026` / `Nov 5, 2026`, pemisah desimal koma vs titik).
  6. Pluralisasi dan interpolasi memakai ICU message format sehingga kalimat seperti "1 work item" / "3 work item" benar di kedua bahasa.
  7. ⚠ Admin dapat menyediakan terjemahan untuk nama status, issue type, resolution, dan custom field per bahasa, dan UI menampilkan terjemahan sesuai bahasa pengguna dengan fallback ke nama asli (CFG-061, WF-040).
  8. Dukungan bahasa RTL tidak diwajibkan; layout boleh mengasumsikan LTR.
- **Team-managed vs company-managed:** Terjemahan nama objek (butir 7) berlaku untuk objek global company-managed; untuk team-managed berlaku pada status dan issue type milik project bila fitur tersedia.
- **Catatan & ketergantungan:** Repo sudah memiliki deteksi bahasa dari `Accept-Language` dan pesan galat EN/ID di `internal/shared/constants/error.go`, tetapi belum ada preferensi bahasa tersimpan per pengguna maupun error code stabil di respons. Zona waktu dipindah ke NFR-007. Portal JSM multibahasa dirinci di [29 JSM Service Project, Request Type & Formulir](29-jsm-service-project-request-type-formulir.md).

### NFR-022 Accessibility (WCAG 2.1 AA)

- **Deskripsi:** Kepatuhan aksesibilitas SPA terhadap WCAG 2.1 level AA, setara komitmen aksesibilitas vendor pembanding.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Pengguna akhir dan QA; preferensi aksesibilitas personal diatur pengguna sendiri (NAV-035).
- **Aturan perilaku (acceptance criteria):**
  1. Semua fungsi dapat dioperasikan dengan keyboard penuh, termasuk alternatif drag-and-drop lewat menu (Move to top, Move to bottom, Send to sprint, pindah kolom) di board dan backlog.
  2. Keyboard shortcut ala vendor pembanding tersedia (`c` create, `/` search, `.` command/actions, `j`/`k` navigasi, `a` assign, `i` assign to me, `e` edit, `m` comment, `g` `d` dashboard, `?` daftar shortcut) dan dapat dinonaktifkan di preferensi; katalog lengkap ada di NAV-038 s.d. NAV-041.
  3. Dialog dan modal menjebak fokus selama terbuka dan mengembalikan fokus ke elemen pemicu saat ditutup; setiap halaman memiliki tautan skip-to-content.
  4. Halaman memakai ARIA landmark, setiap kontrol form memiliki label yang terbaca screen reader, dan flag/toast diumumkan lewat live region.
  5. Rasio kontras teks minimal 4,5:1 (3:1 untuk teks besar dan komponen UI) di tema light dan dark.
  6. Tema light, dark, dan match browser tersedia (NAV-034).
  7. ⚠ Pengaturan aksesibilitas tambahan tersedia di preferensi personal, minimal opsi menggarisbawahi tautan dan mode kontras tinggi (NAV-035).
  8. Informasi tidak disampaikan dengan warna saja: lozenge status selalu berisi teks, dan prioritas memiliki ikon serta label.
  9. Preferensi `prefers-reduced-motion` dihormati: animasi drag, transisi panel, dan animasi dekoratif lain dinonaktifkan atau dikurangi.
  10. axe-core dijalankan di test E2E CI dan menggagalkan build untuk pelanggaran serius/kritis, dan audit manual dengan screen reader NVDA dan VoiceOver dilakukan sebelum rilis mayor.
  11. Target sentuh di mobile web berukuran minimal 24×24 CSS px (usulan 44×44 untuk aksi utama).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kepatuhan per halaman dirinci di NAV-036; chart laporan aksesibel di RPT-008; shortcut editor di EDT-024. Bergantung pada NFR-020 butir 8 dan NFR-024 butir 8.

### NFR-023 Browser & Device Support

- **Deskripsi:** Daftar browser dan perangkat yang didukung SPA beserta perilaku pada browser yang tidak didukung.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Pengguna akhir dan QA.
- **Aturan perilaku (acceptance criteria):**
  1. SPA mendukung dua versi mayor terbaru Chrome, Edge, Firefox, dan Safari (desktop dan mobile), sejalan dengan kebijakan supported browsers vendor pembanding.
  2. Layout responsif untuk tablet dan mobile web: navigasi samping berubah menjadi drawer, dan board dapat digulir horizontal per kolom pada lebar < 768 px.
  3. Internet Explorer tidak didukung dalam versi apa pun.
  4. Bila JavaScript nonaktif, `index.html` menampilkan pesan `noscript` yang menjelaskan bahwa JavaScript diperlukan; bila browser tidak didukung terdeteksi, banner peringatan tampil dengan daftar browser yang didukung.
  5. Test E2E Playwright dijalankan di CI pada engine chromium, firefox, dan webkit untuk alur kritis (NFR-024 butir 5).
  6. Zoom browser hingga 200% tidak menghilangkan fungsi atau konten, dan reflow tanpa scroll horizontal berlaku untuk lebar 320 CSS px di halaman non-board.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aplikasi mobile native (iOS/Android) di luar cakupan modul ini (P3, lihat NAV-050). Bergantung pada NFR-020, NFR-022.

## Kualitas rekayasa

### NFR-024 Testing Strategy

- **Deskripsi:** Strategi pengujian berlapis untuk menjamin kebenaran perilaku fungsional dan NFR, dari unit test hingga failure test infrastruktur.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering dan QA.
- **Aturan perilaku (acceptance criteria):**
  1. Use case diuji dengan unit test memakai mock hasil mockery (pola repo, `.mockery.yaml`, `make mock-gen`).
  2. Repository diuji dengan go-sqlmock untuk unit test, dan query kompleks (JQL, rank, partisi, RLS) diuji tambahan dengan integration test terhadap PostgreSQL nyata via testcontainers.
  3. Golden test JQL memetakan setiap query contoh ke SQL yang dihasilkan, hasil yang diharapkan, dan pesan galat untuk query tidak valid; perubahan output golden harus di-review eksplisit.
  4. Contract test API memverifikasi setiap endpoint terhadap spesifikasi OpenAPI yang dihasilkan (`make swag`); respons yang tidak sesuai skema menggagalkan build.
  5. Test E2E Playwright mencakup alur kritis: login, create work item, drag kartu board, start/complete sprint, transisi, dan pencarian.
  6. Load test k6 dijalankan terhadap dataset sintetis 1M work item yang dibuat oleh generator di repo, dan hasilnya dibandingkan dengan budget NFR-001.
  7. Security test mencakup DAST (OWASP ZAP baseline) terhadap build staging, test matriks izin (setiap peran × setiap endpoint), dan fuzzing parser JQL serta ADF dengan Go native fuzzing.
  8. Test aksesibilitas axe dijalankan di dalam suite E2E (NFR-022 butir 10).
  9. Setiap migrasi diuji up → down → up pada skema kosong dan pada skema berisi data contoh.
  10. Coverage minimal 80% untuk paket usecase dan domain, diukur dengan `make test-coverage` dan ditegakkan di CI.
  11. Failure test mensimulasikan failover DB, object storage S3 tidak tersedia, dan SMTP tidak tersedia, serta memverifikasi mode degraded NFR-010 butir 9.
- **Team-managed vs company-managed:** Test matriks izin wajib mencakup kedua tipe project.
- **Catatan & ketergantungan:** Pola test repo saat ini: mockery dan go-sqlmock (`internal/shared/infrastructure/repository/user_repository_impl_test.go`). Bergantung pada NFR-001, NFR-005, NFR-015, NFR-022, modul 12 (JQL), modul 23 (OpenAPI).

## Keputusan arsitektur (ADR)

Keputusan teknis Fase 1 dicatat sebagai ADR di [`docs/adr/`](../adr/README.md) (S0-14, KAN-11):

| ADR | Keputusan | Fitur terkait |
| --- | --- | --- |
| [ADR-0001](../adr/0001-library-drag-and-drop-board.md) | Drag-and-drop dibangun sendiri sebagai modul frontend reusable (`lib/dnd`), tanpa library eksternal. | NFR-020, NFR-022, BRD-016, BRD-017 |
| [ADR-0002](../adr/0002-penyimpanan-lampiran.md) | Penyimpanan lampiran S3-kompatibel saja (MinIO untuk dev/self-hosted); metadata di PostgreSQL. | CMT-016, CMT-017 |
| [ADR-0003](../adr/0003-transport-email-keluar.md) | Email keluar lewat SMTP di balik interface `Mailer`, dikirim lewat queue. | NTF-023 |
| [ADR-0004](../adr/0004-job-latar-dan-queue.md) | Queue berbasis Redis (`asynq`), Redis wajib; outbox tetap di PostgreSQL dengan relay. | NFR-003, NFR-008, NFR-009 |

## Pertanyaan terbuka

- Apakah konflik edit bersamaan pada field yang sama memakai optimistic locking dengan HTTP 409 (usulan clone) atau last-write-wins per field seperti vendor pembanding demi kompatibilitas API (NFR-006 butir 1)?
- ~~Apakah queue default memakai library River (PostgreSQL) dan Redis cukup opsional, atau Redis dijadikan dependensi wajib sejak Fase 2?~~ **Terjawab:** Redis wajib sejak Fase 1 dengan `asynq`; River tidak dipakai ([ADR-0004](../adr/0004-job-latar-dan-queue.md)).
- Apakah search engine eksternal (OpenSearch) disertakan sebagai komponen resmi atau PostgreSQL FTS menjadi satu-satunya backend sampai Fase 3 (NFR-010 butir 9, SRC-001)?
- Berapa angka default guardrail yang dipakai untuk batas board, jumlah attachment/comment/worklog/link/sub-task per work item, dan `maxResults` endpoint ID-only (NFR-004 butir 2, 4, 5)?
- Apakah mode SaaS multi-tenant akan benar-benar dibangun, atau produk ini tetap single-tenant per deployment secara permanen (NFR-014)?
- Apakah angka retensi audit log per scope di NFR-017 butir 5 (app 90 hari, org 180 hari, automation 90 hari, dengan maksimum per tier) disetujui, dan berapa retensi default log aplikasi, delivery webhook, serta notifikasi (NFR-018 butir 5)?
- Apakah terjemahan nama status, issue type, resolution, dan custom field (NFR-021 butir 7) masuk Fase 1 atau ditunda karena ketersediaannya di vendor Cloud tidak pasti?
- Apakah SPA memakai cookie `httpOnly` (memerlukan CSRF) atau token di memori dengan refresh token (NFR-015 butir 5, NFR-020 butir 5), mengingat repo saat ini memakai JWT Bearer?
- Apakah profil HIPAA-ready dan BYOK (NFR-016 butir 7, NFR-019 butir 4) layak dibangun untuk clone open-source, atau cukup didokumentasikan sebagai panduan konfigurasi operator?
- Apakah logger dimigrasikan dari logrus ke `log/slog` standar Go, atau logrus dipertahankan dengan formatter JSON tanpa warna (NFR-012 butir 1)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga:** app monitoring/APM, app audit log lanjutan, app backup pihak ketiga, app compliance (mis. pengelola retensi data), dan app terjemahan konten tidak dibangun sebagai fitur native; produk ini hanya menyediakan endpoint metrik, streaming SIEM, dan titik ekstensi di modul 25.
- **Fitur Data Center-only / deprecated:** Tidak ada fitur modul ini yang bertanda dc-only-or-deprecated. Konsep yang dipinjam dari edisi Data Center vendor (mode read-only/maintenance di NFR-010 butir 10, panduan large instance di NFR-002) diadopsi sebagai kebutuhan clone self-hosted, sedangkan halaman System info DC dirujuk ke ADM-057 (Fase —); subset yang dibangun (UI job/dead-letter queue) dimiliki NFR-008 butir 4, dan metrik/log runtime oleh NFR-012.
- **Layanan Atlassian lintas produk:** atestasi sertifikasi korporat vendor (SOC 2, ISO, FedRAMP, CSA STAR) tidak diwarisi clone (NFR-019 hanya menyediakan kontrol dan bukti); Atlassian Guard (dirujuk ke USR-037), Atlassian Statuspage (status page opsional NFR-010 butir 12 hanya lewat integrasi), Government Cloud dan Isolated Cloud sebagai penawaran hosting vendor (NFR-014 butir 9 dipenuhi lewat self-hosted), layanan KMS vendor untuk BYOK, serta platform analytics vendor.
- **Kepatuhan pembayaran:** PCI DSS di luar cakupan karena produk ini tidak memproses pembayaran (NFR-019 butir 8).
- **Aplikasi mobile native:** aplikasi iOS/Android (NAV-050) di luar cakupan NFR modul ini; dukungan perangkat terbatas pada mobile web (NFR-023).
- **Dukungan RTL dan Internet Explorer:** sengaja tidak didukung (NFR-021 butir 8, NFR-023 butir 3).
