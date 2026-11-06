# 33 — Service Management: Assets (CMDB)

Modul ini mendefinisikan **Assets**, yaitu CMDB (configuration management database) native produk ini untuk mencatat aset, configuration item (CI), dan entitas bisnis lain (karyawan, lokasi, vendor, lisensi, service) sebagai **object** terstruktur yang saling berelasi. Cakupannya meliputi workspace dan entitlement Assets, model data (object schema, object type, attribute, status, reference type), pengelolaan object (CRUD, list, pencarian, bulk, history, attachment, comment, QR, graph), **Assets Query Language (AQL)**, integrasi dengan work item (field Assets objects, JQL `aqlFunction()`, request form dan portal, praktik ITSM, Services), import dan sinkronisasi data (import structure, mapping, CSV, JSON, jadwal, external import, Entra ID, Discovery, Data Manager, ekspor/impor schema), model data standar dan Hardware Asset Management (HAM), serta otomasi, laporan, REST API, embed di halaman dokumen, dan bantuan AI. Mekanik umum (custom field, JQL engine, automation engine, request type, portal, ITSM) dimiliki modul masing-masing; modul ini mengatur **perilaku khusus Assets** dan titik pemasangannya.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas project, work item, service project, maupun apa pun yang terkait Assets. Seluruh fitur di modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/assets` (schema, object type, attribute, object, konfigurasi global, report), `internal/features/assetsimport` (import structure, mapping, eksekusi, token), dan paket parser/evaluator AQL di `pkg/aql` yang dipakai bersama oleh UI, field, JQL, automation, dan API. Entitas bersama (`ObjectSchema`, `ObjectType`, `ObjectTypeAttribute`, `Object`, `ObjectAttributeValue`, `ObjectReference`, `StatusType`, `ReferenceType`, `ImportSource`) ditempatkan di `internal/shared/domain/entity` dengan repository di `internal/shared/infrastructure/repository`; UI di `frontend/src/routes` (SPA SvelteKit) di bawah route `/assets/...`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah, tier plan (Free/Standard/Premium/Enterprise) sebagai feature flag.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — akun, group, product access/lisensi agent, user yang dinonaktifkan.
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — global permission Administer, browse permission work item yang membatasi connected tickets.
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — custom field, field context, screen, tempat field Assets objects dipasang.
- [06 Work Item Inti](06-work-item-inti.md) — issue view, bulk edit work item.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — penyimpanan attachment, mention, pola comment.
- [12 Bahasa Kueri (JQL)](12-jql-bahasa-kueri.md) — fungsi JQL `aqlFunction()` dan autocomplete.
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — kolom navigator, ekspor CSV.
- [16 Laporan & Insight](16-laporan-insight.md) dan [17 Dashboard & Gadget](17-dashboard-gadget.md) — chart dan gadget berbasis filter.
- [20 Otomasi](20-otomasi.md) — engine rule, rule actor (AUTO-006), scheduled trigger (AUTO-023), kuota eksekusi.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log, batas attachment site, pengaturan global.
- [23 REST API Publik & Webhook](23-api-publik-webhook.md) — konvensi `/api/v1`, autentikasi API token/OAuth, rate limit, format error.
- [24 Integrasi DevOps](24-integrasi-devops.md) — repository yang terhubung (tipe attribute repository).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — migrasi dari server/DC, ekosistem app.
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) — job asinkron, kuota, skala data.
- [28 Model Data Inti](28-model-data-inti.md) — katalog entitas dan event domain.
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) — request type, Forms, field di request form.
- [30 Help Center, Portal & Customers](30-jsm-help-center-portal-pelanggan-kanal.md) — portal customer, detail request portal.
- [31 Queues, SLA, Knowledge Base & Laporan](31-jsm-antrian-sla-knowledge-base-laporan.md) — kolom dan filter queue, SLA berbasis JQL.
- [32 ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — incident, problem, change, halaman Services, approvals.
- [34 Operations](34-jsm-operations-alerting-on-call.md) — tim operasi (tipe attribute team).
- [35 Virtual Service Agent, AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md) — toggle AI, CSM yang memakai Common Data Model, trigger/kondisi/aksi Assets di automation JSM (JSM-AI-016, JSM-AI-017, JSM-AI-018, JSM-AI-020), dan entri keluarga API JSM-AI-031 yang merujuk ke JSM-AST-045 sebagai pemilik spesifikasi Assets REST API.

## Persona terkait

- **Admin Situs** — bertindak sebagai Assets admin: mengaktifkan Assets, membuat object schema, mengelola konfigurasi global (status global, reference type, icon, kuota), menyetel HAM, dan memasang field Assets objects.
- **Manajer Layanan/ITSM** — menjadi object schema manager: merancang model CMDB, mengatur role schema, menyusun import, dan memakai data CI untuk analisis dampak incident/change.
- **Agent Layanan** — mencari dan melihat object, memperbarui CI sebagai object schema developer, menautkan object ke request, menelusuri graph dependency dan connected tickets.
- **Customer/Pelapor** — memilih asetnya sendiri (mis. laptop miliknya) di request form portal tanpa melihat CMDB penuh.
- **Admin Proyek/Project Lead** — menambahkan field Assets ke request type, form, dan queue project service; membuat rule automation berbasis object.
- **Integrator/Developer API** — memakai Assets REST API, Imports REST API dengan import token, dan pencarian AQL dari sistem eksternal.
- **Developer/Anggota Tim** — memakai field Assets objects di project non-service (Assets di produk inti) untuk menautkan komponen atau aset ke work item.
- **Stakeholder/Viewer** — sebagai object schema user, melihat object, laporan Assets, dan tabel Assets yang disematkan di halaman dokumen secara read-only.

## User story modul

1. Sebagai Admin Situs, saya ingin Assets otomatis tersedia sebagai satu workspace per site pada plan yang memenuhi syarat agar tim bisa langsung membangun CMDB tanpa instalasi terpisah (JSM-AST-001, JSM-AST-004).
2. Sebagai Manajer Layanan/ITSM, saya ingin membuat object schema dari template ITAM lalu menyesuaikan object type dan attribute-nya agar model CMDB sesuai organisasi kami (JSM-AST-005, JSM-AST-006, JSM-AST-007, JSM-AST-008).
3. Sebagai Manajer Layanan/ITSM, saya ingin mengatur role Manager, Developer, dan User per schema serta pembatasan per object type agar data sensitif seperti karyawan hanya terlihat oleh pihak berwenang (JSM-AST-002).
4. Sebagai Agent Layanan, saya ingin mencari object dengan teks bebas atau AQL di semua schema yang saya akses agar CI yang relevan ditemukan dalam hitungan detik (JSM-AST-015, JSM-AST-016, JSM-AST-024).
5. Sebagai Agent Layanan, saya ingin melihat graph inbound/outbound reference dan daftar connected tickets sebuah server agar saya bisa menilai dampak incident dan riwayat masalahnya (JSM-AST-022, JSM-AST-023).
6. Sebagai Customer/Pelapor, saya ingin memilih laptop milik saya di portal saat melapor kerusakan agar agent langsung tahu perangkat yang bermasalah (JSM-AST-027, JSM-AST-025).
7. Sebagai Admin Proyek/Project Lead, saya ingin field Assets objects dengan filter AQL bertingkat (laptop berdasarkan karyawan terpilih) agar pilihan object di request selalu relevan (JSM-AST-026, JSM-AST-025).
8. Sebagai Agent Layanan, saya ingin mencari semua incident yang menyentuh object berstatus tertentu lewat JQL `aqlFunction()` agar antrian dan dashboard bisa difokuskan per CI atau service (JSM-AST-028, JSM-AST-029).
9. Sebagai Manajer Layanan/ITSM, saya ingin mengimpor data dari CSV, JSON, dan sumber eksternal secara terjadwal dengan identifier yang mencegah duplikasi agar CMDB selalu sinkron dengan sumber kebenaran (JSM-AST-031 s.d. JSM-AST-036).
10. Sebagai Admin Situs, saya ingin memindai jaringan dan merekonsiliasi data multi-sumber sebelum masuk ke Assets agar inventaris hardware dan software akurat (JSM-AST-038, JSM-AST-039).
11. Sebagai Manajer Layanan/ITSM, saya ingin mengelola katalog model, stockroom, dan stock rule hardware agar pengadaan dan penggantian perangkat terencana (JSM-AST-041, JSM-AST-042).
12. Sebagai Admin Proyek/Project Lead, saya ingin rule automation yang mencari, membuat, dan mengubah object serta bereaksi terhadap perubahan object agar status aset diperbarui otomatis saat request selesai (JSM-AST-043, JSM-AST-029).
13. Sebagai Integrator/Developer API, saya ingin REST API Assets dan Imports REST API dengan import token agar sistem inventaris internal kami dapat menulis ke CMDB secara aman (JSM-AST-045, JSM-AST-036).
14. Sebagai Stakeholder/Viewer, saya ingin laporan jumlah object per type dan status serta tabel Assets yang disematkan di halaman dokumen agar saya memantau inventaris tanpa membuka CMDB (JSM-AST-044, JSM-AST-046).
15. Sebagai Agent Layanan, saya ingin riwayat perubahan object yang lengkap dengan sumber perubahannya agar audit CI dapat dipertanggungjawabkan (JSM-AST-018).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 16 |
| P1 | 17 |
| P2 | 11 |
| P3 | 4 |
| **Total** | **48** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1-6 | 0 |
| Fase 7 | 47 |
| Fase — | 1 |

Strategi urutan implementasi: seluruh modul berada di Fase 7 (ITSM lanjut & Assets/CMDB) karena bergantung pada service project, request type, portal, custom field, JQL, dan automation dari fase sebelumnya; di dalam Fase 7 pekerjaan dibagi menjadi tiga gelombang. Gelombang 7a membangun fondasi P0: workspace dan entitlement (JSM-AST-001), role (JSM-AST-002), model data schema/object type/attribute/reference (JSM-AST-005, JSM-AST-007 s.d. JSM-AST-010), CRUD dan list object (JSM-AST-014, JSM-AST-015), parser AQL (JSM-AST-024, JSM-AST-025), field Assets objects dan portal (JSM-AST-026, JSM-AST-027), serta import CSV dengan mapping (JSM-AST-031 s.d. JSM-AST-033); parser AQL dibangun paling awal karena dipakai oleh list, field, import, JQL, automation, dan API. Gelombang 7b menambahkan fitur P1 yang menaikkan nilai operasional: konfigurasi global dan kuota, template, status, history, graph, connected tickets, JQL, ITSM, Services, JSON dan scheduled import, automation, serta REST API. Gelombang 7c menampung P2/P3 yang bergantung pada agen on-premise atau ekosistem luar (Discovery, Data Manager, Entra ID, external import, Common Data Model, HAM, laporan, QR, embed dokumen, AI); fitur warisan Insight (JSM-AST-048) tidak dibangun dan hanya menjadi rujukan migrasi. Inventaris menggabungkan sepuluh entri lampiran JSM (`jsm-appendix`) ke fitur rinci domain Assets yang setara, dicatat lewat baris **Alias inventaris** atau catatan penggabungan.

## Fitur

## Fondasi workspace, akses & konfigurasi global

### JSM-AST-001 Assets Workspace, Access & Product Entitlement

- **Deskripsi:** Assets adalah CMDB native produk ini (dahulu dikenal sebagai Insight di Jira). Setiap site memiliki tepat satu Assets workspace (`workspaceId`) yang menampung semua object schema, konfigurasi global, import, dan API; aksesnya bergantung pada plan dan lisensi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (global permission Administer, otomatis Assets admin); Agent Layanan berlisensi service; pengguna dengan role schema; Customer hanya lewat portal.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap site memiliki tepat satu Assets workspace dengan `workspaceId` (UUID) yang stabil; `GET /api/v1/servicedesk/assets/workspace` (padanan `GET /rest/servicedeskapi/assets/workspace`) mengembalikan `{workspaceId}` untuk pengguna terautentikasi berlisensi, dan 404 bila workspace belum disediakan.
  2. Workspace disediakan otomatis saat site memiliki plan Service Management Premium/Enterprise; menu "Assets" muncul di navigasi global (sidebar) hanya bila workspace aktif dan pengguna punya akses.
  3. Halaman landing Assets menampilkan daftar object schema yang boleh diakses pengguna, kotak pencarian object, serta shortcut ke Import, Reports, dan Configure (Configure hanya untuk Assets admin).
  4. ⚠ Entitlement per plan: Assets penuh hanya aktif pada tier Premium/Enterprise; pada Free/Standard seluruh endpoint Assets mengembalikan 403 `ASSETS_NOT_ENTITLED` dan menu disembunyikan. Pemetaan tier dikonfigurasi lewat feature flag plan (VISI) agar perubahan kebijakan paket tidak memerlukan perubahan kode.
  5. ⚠ Field Assets objects dapat dipakai di project non-service pada site yang sama bila entitlement "Assets di produk inti" aktif; bila nonaktif, tipe field tidak muncul di daftar tipe custom field project non-service.
  6. ⚠ "Assets di produk inti": Assets juga dapat diaktifkan untuk plan produk inti Premium/Enterprise tanpa Service Management; workspace yang sama dipakai, bukan workspace kedua.
  7. ⚠ Sistem menyimpan matriks kapabilitas per entitlement: fitur yang memerlukan Service Management (portal customer, Services, HAM, connected tickets dari request) dinonaktifkan dan disembunyikan bila site hanya punya entitlement produk inti.
  8. Object Assets dapat ditautkan ke work item project non-service lewat field Assets objects (JSM-AST-026) dengan perilaku picker dan tampilan yang sama seperti di project service.
  9. ⚠ Keterbatasan Assets di produk inti didokumentasikan sebagai daftar flag kapabilitas yang dapat dibaca lewat API konfigurasi (`GET /api/v1/assets/config/capabilities`), sehingga UI dapat menampilkan alasan fitur tidak tersedia.
  10. Bila site memiliki entitlement produk inti dan Service Management sekaligus, keduanya memakai satu workspace, satu kuota object (JSM-AST-004), dan satu set konfigurasi global.
  11. ⚠ Saat plan di-downgrade, data Assets tetap tersimpan utuh tetapi UI dan API menjadi tidak dapat diakses (403 `ASSETS_NOT_ENTITLED`) sampai plan dinaikkan kembali; nilai field Assets objects di work item tetap tersimpan dan ditampilkan sebagai teks label read-only.
  12. UI Assets hanya dapat dibuka oleh pengguna berlisensi agent service atau yang diberi role schema; Customer yang membuka URL `/assets/...` diarahkan ke portal dan API mengembalikan 403.
  13. Pengguna dengan global permission Administer otomatis menjadi Assets admin: dapat membuat schema, mengelola konfigurasi global, dan mengakses semua schema meskipun tidak diberi role eksplisit.
  14. Pengguna pada plan yang tidak mendukung yang membuka menu atau URL Assets melihat halaman penjelasan ("Assets tersedia pada plan Premium dan Enterprise") dengan tautan ke admin, bukan halaman galat kosong.
- **Team-managed vs company-managed:** Workspace bersifat global per site; field Assets objects dapat dipasang di project company-managed lewat field context dan di project team-managed sebagai field project-scoped bila entitlement mengizinkan.
- **Catatan & ketergantungan:** Tier dan feature flag di [00](00-visi-ruang-lingkup-glosarium-baseline.md); lisensi agent di [02](02-pengguna-grup-organisasi.md); navigasi di [21](21-navigasi-ux-preferensi.md). Billing dan pengelolaan paket bundel vendor asal tidak dibangun. Entitas `AssetsWorkspace` disimpan per site; karena repo saat ini single-tenant, workspace dibuat sebagai baris tunggal saat Assets diaktifkan.

### JSM-AST-002 Object Schema & Object Type Roles (Permissions)

- **Alias inventaris:** Assets permissions & roles
- **Deskripsi:** Kontrol akses berbasis role per object schema (Object schema managers, Object schema developers, Object schema users), akses khusus customer portal, serta override atau pembatasan per object type untuk type sensitif.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Assets admin (Admin Situs); Object schema manager mengelola role schema-nya; Agent Layanan dan Customer sebagai penerima role.
- **Aturan perilaku (acceptance criteria):**
  1. Role Object schema managers memberi kontrol penuh atas schema: mengubah konfigurasi, role, status schema, object type, attribute, import, dan semua object di schema itu.
  2. ⚠ Role Object schema developers dapat membuat, mengubah, dan menghapus object serta menjalankan import yang sudah ada, tetapi tidak dapat mengubah object type, attribute, role, atau konfigurasi schema; tombol Configure tidak tampil untuk role ini.
  3. Role Object schema users hanya dapat melihat object (read-only); upaya create/edit/delete lewat UI atau API mengembalikan 403 `ASSETS_PERMISSION_DENIED`.
  4. ⚠ Role dapat ditetapkan ke pengguna individual dan group; penetapan ke project role (mis. semua agent project service tertentu) didukung sebagai opsi tambahan bila diputuskan (lihat Pertanyaan terbuka).
  5. ⚠ Schema memiliki toggle "Akses customer portal" yang, bila aktif, mengizinkan Customer memilih object dari schema itu di request form portal tanpa memberi akses UI Assets.
  6. ⚠ Agent Layanan tidak otomatis memperoleh akses ke schema baru; schema baru hanya dapat diakses pembuatnya, Assets admin, dan pihak yang diberi role setelahnya.
  7. ⚠ Object type roles: object schema manager dapat menetapkan role per object type yang menimpa role schema, baik untuk memperluas (mis. developer hanya pada type "Laptop") maupun membatasi (mis. type "Employee" hanya untuk group HR); role object type berlaku turun ke child type yang mewarisi.
  8. UI menyembunyikan atau menonaktifkan tombol Create object, Edit, Delete, Import, dan Configure sesuai role efektif pengguna pada schema/object type tersebut.
  9. Akses REST API memakai role pemilik API token/OAuth, dan automation memakai role rule actor; tidak ada jalur yang melewati pemeriksaan role.
  10. Konfigurasi global (status global, reference type, icon global, kuota) hanya dapat diubah oleh Assets admin; object schema manager hanya melihatnya.
  11. ⚠ Role schema minimal yang tersedia adalah Object schema managers, Object schema users, dan Object viewers (read-only; dalam produk ini disamakan dengan Object schema users pada butir 3), ditambah Object schema developers dari butir 2; daftar role tambahan dapat diperluas tanpa mengubah model data (role disimpan sebagai tabel `assets_role` dengan daftar permission).
  12. Role object type yang diterapkan mengoverride role schema untuk object type tersebut, dan hasil override dapat dilihat di tab Roles sebagai "role efektif" per pengguna.
  13. ⚠ Customer project service hanya dapat mengakses object tertentu melalui picker field Assets di portal sesuai filter AQL field (mis. `Owner = currentReporter()`), bukan melalui role schema.
  14. Pengaturan global menentukan siapa yang dapat membuat object schema (default: hanya Assets admin).
  15. Akses penuh UI Assets mensyaratkan lisensi agent service (atau entitlement produk inti); pengguna tanpa lisensi yang diberi role schema hanya dapat melihat object lewat field Assets objects di work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model permission umum di [03](03-izin-peran-keamanan.md); rule actor di AUTO-006 ([20](20-otomasi.md)); visibilitas field untuk pengguna tanpa role diatur oleh setting schema di JSM-AST-005 dan JSM-AST-026. Butir 11-15 berasal dari entri lampiran "Assets permissions & roles".

### JSM-AST-003 Assets Global Configuration

- **Deskripsi:** Pengaturan tingkat workspace yang berlaku lintas schema: status global, reference type, library icon, siapa yang boleh membuat schema, info kuota, audit konfigurasi, dan toggle AI.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Assets admin (Admin Situs); object schema manager hanya membaca.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Configure > Global menyediakan tab Statuses yang mengelola status global sesuai JSM-AST-012.
  2. Tab Reference types mengelola reference type sesuai JSM-AST-013.
  3. Tab Icons menampilkan library icon bawaan dan kustom, dengan penanda lingkup global (semua schema) atau khusus schema tertentu.
  4. ⚠ Create custom icon menerima unggahan PNG atau SVG maksimal 1 MB dengan dimensi persegi; SVG disanitasi (skrip dan event handler dibuang) sebelum disimpan, dan file di luar format ditolak dengan pesan "Format icon tidak didukung".
  5. Icon kustom dapat di-rename; nama wajib unik dalam lingkupnya.
  6. ⚠ Delete icon kustom meminta konfirmasi yang menyebut jumlah object type pemakai; setelah dihapus, object type tersebut kembali memakai icon default.
  7. ⚠ Setting "Siapa yang dapat membuat object schema" berisi pilihan Assets admin saja (default) atau daftar group tambahan.
  8. Panel kuota menampilkan jumlah object terpakai, batas plan, dan persentase (sumber data JSM-AST-004).
  9. ⚠ Setiap perubahan konfigurasi Assets (schema, object type, attribute, role, status, reference type, icon, import structure) dicatat ke audit log site dengan aktor, waktu, kategori "Assets", dan ringkasan perubahan.
  10. ⚠ Toggle fitur AI untuk Assets (JSM-AST-047) tersedia di halaman ini dan default nonaktif; bila toggle AI tingkat site nonaktif, toggle ini terkunci.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Audit log di [22](22-administrasi-sistem.md); penyimpanan file icon memakai storage attachment ([08](08-komentar-lampiran-aktivitas.md)).

### JSM-AST-004 Assets Object Limits & Quotas

- **Deskripsi:** Batas kapasitas Assets per site/plan dan batas teknis (attribute, hasil AQL, ukuran file import, field multi-object, bulk) yang harus dapat dikonfigurasi dan ditegakkan konsisten.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Assets admin dan Admin Situs melihat dan (untuk instalasi self-host) mengubah batas; semua pengguna menerima pesan galat saat batas tercapai.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Batas jumlah object per workspace adalah parameter konfigurasi dengan default 3.000.000 untuk tier Premium/Enterprise; object yang dihapus tidak dihitung.
  2. Indikator pemakaian kuota (terpakai/batas, persentase, bar) tampil di halaman konfigurasi global dan memberi peringatan visual pada 80% dan 95%.
  3. Saat kuota tercapai, create object (UI, API, automation) dan import yang akan membuat object baru ditolak dengan galat 409 `ASSETS_OBJECT_LIMIT_REACHED` dan pesan yang menyebut batasnya; update object existing tetap diizinkan.
  4. ⚠ Batas jumlah attribute per object type default 150 (termasuk attribute warisan); penambahan melewati batas ditolak saat validasi.
  5. ⚠ `maxResults` pencarian AQL per halaman default 25 dan maksimum 1.000 untuk API; nilai di atas maksimum dipotong ke maksimum dan respons menyertakan nilai efektif.
  6. ⚠ Batas ukuran file import CSV/JSON default 100 MB per file; file lebih besar ditolak sebelum diproses dengan pesan yang menyebut batas.
  7. ⚠ Batas jumlah object dalam satu nilai field Assets objects multiple default 100; pemilihan melewati batas ditolak di picker dan API.
  8. ⚠ Batas jumlah object per operasi bulk default 1.000; pilihan melebihi batas memunculkan pesan untuk mempersempit filter.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semua angka default perlu dikonfirmasi pemilik produk dan didokumentasikan sebagai parameter di [27](27-nfr-arsitektur-platform.md); kuota dipakai bersama oleh produk inti dan Service Management (JSM-AST-001 butir 10).

## Model data: schema, object type & attribute

### JSM-AST-005 Object Schemas

- **Alias inventaris:** Assets: object schemas
- **Deskripsi:** Object schema adalah container tingkat atas yang mengelompokkan object type, object, role, status khusus schema, dan import (mis. IT Assets, HR, Facilities). Admin dapat membuat, mengubah, dan menghapus schema.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Assets admin (atau group yang diizinkan di JSM-AST-003) membuat schema; Object schema manager mengubah dan menghapus schema-nya.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Dialog Create object schema meminta pilihan template (JSM-AST-006) atau "Empty schema", lalu Name (wajib, maks. 255 karakter, unik di workspace), Key (wajib, 2-10 karakter huruf besar A-Z dan angka, diawali huruf, unik di workspace), dan Description (opsional); pelanggaran menampilkan galat per field.
  2. ⚠ Key schema menjadi prefix object key (mis. `ITSM-123`); key schema tidak dapat diubah setelah object pertama dibuat, dan UI menampilkan key sebagai read-only dengan penjelasannya.
  3. Daftar object schema menampilkan Name, Key, jumlah object type, jumlah object, dan icon untuk setiap schema yang boleh diakses pengguna.
  4. Panel tree kiri menampilkan hierarki object type schema beserta jumlah object per type; memilih node membuka list object type tersebut.
  5. Schema configuration memiliki tab General/Details, Roles, Statuses (khusus schema), Import, dan Object types; tab hanya terlihat oleh object schema manager dan Assets admin.
  6. Detail schema (Name, Description, icon) dapat diubah; perubahan Name memvalidasi keunikan dan tercatat di audit log.
  7. Setting "Izinkan pengguna lain memilih object dari schema ini" (default nonaktif) membuat field Assets objects dapat menampilkan dan memilih object kepada pengguna berlisensi tanpa role schema.
  8. Setting akses customer portal (JSM-AST-002 butir 5) tersedia di tab General dan wajib aktif agar picker portal menampilkan object dari schema ini.
  9. Delete object schema meminta konfirmasi dengan mengetik key schema; tindakan menghapus semua object type, object, status khusus schema, dan import structure, serta menampilkan peringatan berisi daftar field Assets objects yang memakai schema ini; penghapusan berjalan asinkron dengan progress.
  10. Daftar schema dapat dicari/difilter berdasarkan nama (case-insensitive, substring).
  11. Schema "Services" dibuat oleh sistem untuk menyimpan service (JSM-AST-030), ditandai sebagai schema sistem, dan tidak dapat dihapus.
  12. Schema sistem Common Data Model (JSM-AST-041) ditampilkan dengan badge "Sistem" dan dipisahkan dari schema kustom di daftar.
  13. Pembuatan schema dari template tersedia untuk kategori IT assets, HR, Facilities, dan Blank (rinci di JSM-AST-006); template terpilih ditampilkan di langkah ringkasan sebelum schema dibuat.
  14. Name, Key, dan Description schema tampil di header halaman schema dan di breadcrumb object detail.
  15. Schema configuration juga memuat pengelolaan object status, reference type yang tersedia bagi schema, dan role; perubahan di tab tersebut berlaku segera tanpa reindex manual.
  16. ⚠ Struktur schema dapat diekspor dan diimpor ke workspace lain (rinci di JSM-AST-040).
  17. Penghapusan schema ikut menghapus semua object di dalamnya secara permanen; tidak ada trash untuk schema yang dihapus.
  18. Halaman schema menampilkan tampilan pohon object type di kiri dan daftar object type terpilih di kanan dalam satu layout.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 13-18 berasal dari entri lampiran "Assets: object schemas". Entitas `ObjectSchema` (id, workspace_id, name, key, description, icon_id, is_system, allow_others_select, portal_access, object_sequence). Prefix key dipakai juga oleh JSM-AST-014.

### JSM-AST-006 Object Schema Templates

- **Deskripsi:** Template siap pakai untuk membuat object schema dengan object type dan attribute yang sudah terdefinisi untuk kasus umum (ITAM, HR, Facilities, struktur Discovery), berbeda dari schema sistem Common Data Model.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Assets admin atau group yang diizinkan membuat schema.
- **Aturan perilaku (acceptance criteria):**
  1. Template IT Asset Management membuat object type minimal Hardware (dengan child Laptop, Desktop, Server, Mobile device), Software, Licenses, Locations, dan Vendors beserta attribute dan reference dasarnya.
  2. Template Human Resources membuat object type Employees, Departments, dan Positions dengan reference Employee → Department dan Employee → Position.
  3. Template Facilities Management membuat object type Buildings, Floors, Rooms, dan Desks dengan hierarki reference "Located in".
  4. Template struktur Discovery membuat object type Hosts, Network Devices, Software, dan Operating Systems yang cocok dengan mapping bawaan Assets Discovery (JSM-AST-038).
  5. Empty schema membuat schema tanpa object type.
  6. Langkah preview menampilkan daftar object type (tree) dan attribute utama template sebelum pengguna menekan Create.
  7. Object type dan attribute hasil template sepenuhnya dapat diedit, di-rename, dan dihapus setelah dibuat; tidak ada keterikatan ke template.
  8. ⚠ Template hanya membuat struktur (object type, attribute, status, reference type) tanpa contoh object; keputusan menyediakan opsi "sertakan data contoh" dicatat di Pertanyaan terbuka.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template disimpan sebagai berkas definisi (JSON) versi-an di binary sehingga dapat diperbarui antar rilis; formatnya sama dengan format ekspor schema (JSM-AST-040). Berbeda dari schema sistem di JSM-AST-041.

### JSM-AST-007 Object Types

- **Alias inventaris:** Assets: object types & attributes
- **Deskripsi:** Object type mendefinisikan kelas object dalam schema (mis. Laptop, Server, Employee) beserta attribute, icon, dan posisi di hierarki, termasuk inheritance dan abstract type.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager dan Assets admin.
- **Aturan perilaku (acceptance criteria):**
  1. Create object type meminta Name (wajib, unik dalam schema), Description, Icon (wajib, default icon generik), dan Parent object type (opsional).
  2. ⚠ Object type membentuk hierarki parent-child di tree dengan kedalaman maksimum 10 level; upaya membuat level ke-11 ditolak.
  3. Urutan dan parent object type dapat diubah dengan drag-and-drop di tree; perpindahan parent memvalidasi aturan inheritance (butir 6).
  4. Abstract object type tidak dapat memiliki object sendiri dan hanya menjadi parent/grup; tombol Create object dinonaktifkan pada type abstract.
  5. Opsi "Inherit attributes" membuat child mewarisi semua attribute parent; attribute warisan tampil read-only di child dan hanya dapat diubah di parent.
  6. ⚠ Inheritance dan status abstract tidak dapat diubah bila object type sudah memiliki child type (untuk inheritance) atau object (untuk abstract); UI menjelaskan alasannya.
  7. Opsi "Tampilkan object dari child type" pada parent membuat list parent memuat object semua turunan.
  8. Object type dapat di-rename, diganti icon, dan deskripsinya diubah.
  9. Delete object type menghapus semua object-nya setelah konfirmasi; dialog menampilkan peringatan bila type dipakai sebagai target reference, oleh field Assets objects, atau oleh import mapping, dan menolak penghapusan bila type masih memiliki child type.
  10. ⚠ Copy object type menduplikasi definisi (attribute, icon, pengaturan) tanpa object, dengan nama "Copy of <nama>".
  11. Jumlah object per object type tampil di tree dan diperbarui setelah create/delete/import.
  12. Pengaturan per object type mencakup pemilihan label attribute, kolom default di list, dan object type roles (JSM-AST-002).
  13. ⚠ Icon object type dipilih dari library bawaan atau diunggah (PNG/SVG maks. 1 MB, persegi) langsung dari dialog object type.
  14. Icon dapat berlingkup global (semua schema) atau khusus schema; icon object type tampil di tree, list, object view, field Assets objects, picker portal, dan graph viewer.
  15. Hierarki object type menampilkan icon tiap type di setiap node tree.
  16. Abstract object type dipakai bersama inheritance agar attribute bersama didefinisikan sekali di parent.
  17. Attribute kategori Default yang tersedia saat mendefinisikan object type mencakup Text, Textarea, Integer, Float, Boolean, Date, DateTime, URL, Email, Select, dan IP address (rinci di JSM-AST-009).
  18. ⚠ Attribute non-default yang tersedia mencakup Object (reference), User, Group, Status, dan Project (rinci di JSM-AST-010 dan JSM-AST-011).
  19. Setiap attribute dapat dikonfigurasi unique, cardinality min/max, regex, hidden, label, read-only, dan description (rinci di JSM-AST-008).
  20. Attribute Label menjadi nama tampilan object, sedangkan key object dibangkitkan otomatis dari key schema (JSM-AST-014).
  21. Reference type (Depends on, Installed, Located in, Owned by, dan kustom) dengan warna dapat dipilih pada attribute Object (JSM-AST-013).
  22. Object status memiliki kategori yang menentukan warna lozenge (JSM-AST-012).
  23. Setiap object type otomatis memiliki attribute sistem Key, Created, dan Updated yang tidak dapat dihapus.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 15-23 berasal dari entri lampiran "Assets: object types & attributes" dan merujuk ke fitur rinci terkait. Entitas `ObjectType` (id, schema_id, parent_id, name, icon_id, abstract, inherit, position, label_attribute_id).

### JSM-AST-008 Attributes (Object Type Attributes)

- **Deskripsi:** Attribute mendefinisikan field data pada object type. Setiap object type memiliki attribute sistem (Key, Name, Created, Updated) dan attribute kustom dengan aturan cardinality, unique, label, dan validasi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager.
- **Aturan perilaku (acceptance criteria):**
  1. Object type baru otomatis memiliki attribute sistem Key, Name (label default), Created, dan Updated; Key/Created/Updated read-only dan tidak dapat dihapus.
  2. Tambah attribute meminta Name (wajib, unik dalam object type termasuk attribute warisan), Type, sub-tipe/type value (mis. target object type untuk Object), dan Description.
  3. Urutan attribute diatur dengan drag-and-drop dan menentukan urutan tampil di object view dan dialog create/edit.
  4. ⚠ Edit attribute mengizinkan perubahan Name, Description, dan opsi; perubahan Type ditolak bila attribute sudah memiliki nilai pada object mana pun, kecuali perubahan Text ↔ Textarea.
  5. Delete attribute menghapus semua nilainya setelah konfirmasi; dialog memperingatkan bila nama attribute dipakai di filter AQL field Assets objects, import mapping, atau label attribute.
  6. Tepat satu attribute per object type ditandai sebagai label (default "Name"); label tidak boleh bertipe Object multi, dan mengganti label langsung mengubah nama tampilan semua object.
  7. Opsi Unique mewajibkan nilai unik dalam object type (case-insensitive untuk teks) dan divalidasi saat create, edit, import, dan API dengan galat "Nilai sudah dipakai oleh <key>".
  8. Cardinality minimum (0 = opsional, ≥1 = wajib) dan maksimum (1 atau tak terbatas) divalidasi di semua jalur penulisan.
  9. Validasi Regex dengan pesan galat kustom tersedia untuk tipe Text, Textarea, Email, dan URL; regex yang tidak valid ditolak saat disimpan.
  10. ⚠ Opsi tampilan "Hidden" menyembunyikan attribute dari object view dan list (tetap dapat dicari dengan AQL dan diakses API).
  11. ⚠ Opsi Summable untuk Integer/Float menampilkan total kolom di list object.
  12. Attribute warisan ditandai icon "inherited" dan tombol edit/delete-nya dinonaktifkan di child.
  13. Setiap attribute memiliki ID numerik unik yang stabil dan dapat dipakai di AQL, import mapping, dan REST API sebagai alternatif nama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Opsi read-only dan description dari entri lampiran tercakup di JSM-AST-007 butir 19; nilai disimpan di `ObjectAttributeValue` dengan kolom bertipe (text, number, bool, date, ref_id) agar AQL dapat dieksekusi efisien di PostgreSQL.

### JSM-AST-009 Default Attribute Types

- **Deskripsi:** Tipe attribute dasar (kategori Default) untuk nilai primitif dengan validasi khusus per tipe.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager mendefinisikan; pengguna yang boleh mengedit object mengisi nilai.
- **Aturan perilaku (acceptance criteria):**
  1. Text: string satu baris maks. 255 karakter, opsi validasi Regex.
  2. ⚠ Textarea: teks multi-baris maks. 32.000 karakter; disimpan sebagai teks polos dan dirender dengan pemutusan baris (dukungan rich text menjadi keputusan terbuka).
  3. Integer: bilangan bulat 64-bit, opsi suffix/satuan (mis. "GB") yang ditampilkan setelah nilai.
  4. Float: bilangan desimal dengan presisi tampilan yang dapat diatur, opsi suffix/satuan.
  5. Boolean: true/false, ditampilkan sebagai toggle; nilai kosong dibedakan dari false.
  6. Date: tanggal tanpa waktu, ditampilkan mengikuti format tanggal site dan disimpan sebagai ISO `YYYY-MM-DD`.
  7. Date Time: tanggal dan waktu disimpan dalam UTC dan ditampilkan sesuai timezone pengguna.
  8. URL: validasi format URL absolut (http/https); nilai tampil sebagai tautan yang dibuka di tab baru.
  9. ⚠ Email: validasi format email; opsi pembatasan domain yang diizinkan (daftar domain) per attribute.
  10. Select: daftar opsi statis yang dikelola pada attribute; multi-value diaktifkan dengan cardinality maksimum > 1, dan menghapus opsi yang sedang dipakai memerlukan konfirmasi.
  11. ⚠ IP Address: validasi format IPv4 dan IPv6.
  12. ⚠ Tipe Time (durasi tanpa tanggal) tidak disediakan; definisi yang diimpor dari server/DC dengan tipe ini dipetakan ke Text dengan peringatan migrasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format tanggal dan timezone mengikuti preferensi di [21](21-navigasi-ux-preferensi.md).

### JSM-AST-010 Object Reference Attribute

- **Deskripsi:** Attribute bertipe Object yang membuat relasi (reference) antar object dan membentuk dependency graph CMDB.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager mendefinisikan; Agent Layanan dengan role developer mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Attribute Object wajib memilih target object type, yang boleh berada di object schema lain dalam workspace yang sama.
  2. Attribute Object wajib memilih Reference type (JSM-AST-013) yang menentukan label dan warna relasi.
  3. Opsi "Sertakan child type" mengizinkan pemilihan object dari turunan target object type.
  4. ⚠ Filter AQL opsional pada attribute membatasi object yang boleh dipilih (mis. `Status = "In Stock"`); nilai yang tidak memenuhi filter ditolak saat disimpan.
  5. Cardinality min/max menentukan reference tunggal atau multi.
  6. Reference disimpan sebagai outbound di object sumber dan tampil sebagai inbound di object target tanpa duplikasi data.
  7. Object terkait di object view dapat diklik untuk membuka object tersebut.
  8. AQL mendukung dot notation untuk attribute object terkait (mis. `"Owner"."Email" LIKE "@contoh.id"`).
  9. Menghapus object target menghapus reference-nya di semua object sumber; dialog delete menampilkan jumlah inbound reference yang akan hilang.
  10. Attribute Object dapat menunjuk ke object type di schema sistem Common Data Model (mis. Geography, People) bila pengguna punya akses ke schema tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Reference disimpan di tabel `object_reference` (source_object_id, target_object_id, attribute_id, reference_type_id) dengan indeks dua arah untuk graph (JSM-AST-022) dan fungsi `inR()/outR()` (JSM-AST-025).

### JSM-AST-011 User, Group, Status & Integration Attribute Types

- **Deskripsi:** Tipe attribute non-default yang merujuk ke entitas platform atau konfigurasi Assets: User, Group, Status, halaman dokumen, repository, tim operasi, serta Project/Version.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager mendefinisikan; pengguna berhak edit mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. User: memilih akun pengguna site; opsi membatasi ke group tertentu; nilai dapat dicocokkan dengan `currentUser()`/`currentReporter()` di AQL.
  2. Group: memilih group site.
  3. Status: memilih dari status Assets (global dan khusus schema); opsi membatasi status yang diizinkan per attribute.
  4. ⚠ Tipe halaman dokumen: menyimpan referensi ke halaman di sistem dokumentasi yang terhubung (URL + judul); integrasi wiki eksternal hanya berupa tautan dan opsi filter space tidak disediakan di fase awal.
  5. Repository: merujuk repository yang terhubung melalui integrasi DevOps (modul 24).
  6. ⚠ Tim operasi: merujuk team di modul Operations ([34](34-jsm-operations-alerting-on-call.md)) dengan nama tipe "Operations team".
  7. ⚠ Project dan Version: merujuk project dan version produk ini; ketersediaannya dirilis bersama atau setelah tipe User/Group.
  8. ⚠ Tidak ada tipe attribute "Work item"; keterkaitan work item ↔ object hanya lewat field Assets objects dan connected tickets (JSM-AST-023).
  9. ⚠ User/group yang dihapus dari site tetap tersimpan sebagai referensi dan ditampilkan sebagai "Pengguna tidak dikenal"/"Group dihapus" tanpa error di object view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Akun dan group di [02](02-pengguna-grup-organisasi.md); repository di [24](24-integrasi-devops.md). Tipe yang merujuk produk eksternal vendor asal (wiki, Git hosting, alerting) diganti padanan modul produk ini.

### JSM-AST-012 Statuses (Assets Status Types)

- **Deskripsi:** Konfigurasi status object (mis. In Use, In Stock, Retired) dengan kategori, global maupun per schema, untuk attribute bertipe Status.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Assets admin (status global); Object schema manager (status khusus schema).
- **Aturan perilaku (acceptance criteria):**
  1. Status global berlaku di semua schema dan dikelola di konfigurasi global.
  2. Status khusus schema hanya tersedia di satu schema dan dikelola di tab Statuses schema configuration.
  3. Status memiliki Name (wajib, unik dalam lingkupnya), Description, dan Category (Active / Inactive / Pending).
  4. Warna lozenge status mengikuti kategori (Active hijau, Inactive abu-abu, Pending biru/kuning) secara konsisten di list, object view, dan picker.
  5. Status dapat dibuat, diubah, dan dihapus; delete status yang sedang dipakai object ditolak dengan pesan yang menyebut jumlah object pemakai.
  6. Attribute Status dapat dibatasi ke subset status yang diizinkan.
  7. AQL mendukung filter berdasarkan nama status (mis. `Status = "In Use"`) dan kategori.
  8. REST API menyediakan resource `/config/statustype` (list, get, create, update, delete).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status Assets terpisah dari status workflow work item ([09](09-workflow-status-transisi.md)).

### JSM-AST-013 Reference Types

- **Deskripsi:** Definisi jenis relasi antar object (nama, deskripsi, warna) yang dipakai di attribute Object dan fungsi AQL references.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Workspace baru berisi reference type bawaan: Depends on, Installed, Link, Location, Owner, dan Reference.
  2. Create reference type meminta Name (wajib, unik), Description, dan Color.
  3. Reference type dapat diubah; delete hanya diizinkan bila tidak dipakai attribute mana pun, selain itu ditolak dengan daftar attribute pemakai.
  4. ⚠ Reference type berlingkup global untuk workspace; lingkup per schema menjadi keputusan terbuka.
  5. Reference type dapat dipakai sebagai filter di `inboundReferences()`/`outboundReferences()`.
  6. Warna reference type tampil pada edge graph viewer dan di daftar references object.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai oleh JSM-AST-010, JSM-AST-022, JSM-AST-025.

## Object: pengelolaan, tampilan & penelusuran

### JSM-AST-014 Objects (Create, View, Edit, Delete)

- **Alias inventaris:** Assets: objects
- **Deskripsi:** Object adalah record/CI individual (mis. satu laptop tertentu) dengan nilai attribute, key unik, reference, dan aktivitas.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema developer dan manager (create/edit/delete); Object schema user (view); Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Create object meminta object type (atau terisi dari konteks list), menampilkan attribute sesuai urutan, memvalidasi required/unique/regex/format sebelum simpan, dan menyediakan opsi "Create another" yang mempertahankan dialog terbuka.
  2. Object key dibangkitkan otomatis `<SCHEMA_KEY>-<nomor>` dari sequence per schema yang berurutan dan tidak pernah dipakai ulang; key tidak dapat diubah manual.
  3. Object memiliki ID numerik internal (dipakai API) yang berbeda dari object key yang mudah dibaca; keduanya dapat dipakai untuk membuka object.
  4. Object detail view memuat header (icon, label, key, object type, breadcrumb schema), panel attribute, references/graph, history, attachments, comments (bila aktif), dan connected tickets.
  5. Attribute dapat diedit inline satu per satu (simpan saat blur/Enter, batal dengan Esc) atau melalui mode edit penuh.
  6. Penyimpanan yang melanggar cardinality minimum (attribute wajib kosong) ditolak dengan galat per attribute; tidak ada penyimpanan parsial.
  7. ⚠ Clone object membuat object baru dengan nilai attribute yang sama kecuali attribute unique (dikosongkan) dan Key/Created/Updated (dibangkitkan ulang).
  8. Delete object meminta konfirmasi dan menampilkan peringatan bila object direferensikan object lain atau tertaut di work item (jumlah inbound reference dan connected tickets).
  9. Setiap object memiliki deep link URL stabil `/assets/object/<id>` dan `/assets/object/<KEY>` yang menghormati permission.
  10. ⚠ Watch object: pengguna dapat mengikuti object dan menerima notifikasi in-app saat attribute berubah; dirilis setelah fitur inti bila diputuskan (lihat Pertanyaan terbuka).
  11. ⚠ Move object ke object type lain dalam schema yang sama dengan langkah mapping attribute; attribute tanpa pasangan dibuang setelah konfirmasi.
  12. ⚠ Object detail dapat dilihat dari aplikasi mobile/tampilan responsif SPA dalam mode baca.
  13. Object detail menampilkan bagian inbound dan outbound references terpisah, connected work items, attachments, comments, dan history sebagai tab.
  14. Object view menyediakan tombol untuk membuka graph viewer relasi object (JSM-AST-022).
  15. Bulk edit dan bulk delete berdasarkan filter AQL tersedia dari list object (rinci di JSM-AST-017).
  16. ⚠ Print label/QR code object tersedia dari object view (rinci di JSM-AST-021).
  17. Object dapat ditampilkan dalam mode list maupun tree object type dengan kolom yang dapat dipilih (rinci di JSM-AST-015).
  18. Object hasil filter dapat diekspor ke CSV dengan nilai attribute terformat (rinci di JSM-AST-017 butir 4).
  19. Setiap perubahan nilai attribute object tercatat di history sebagai audit (rinci di JSM-AST-018).
  20. Siklus dasar object mencakup create, edit, delete, dan clone; keempat aksi tersedia di menu object view sesuai role dan masing-masing menghasilkan entri history.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 13-20 berasal dari entri lampiran "Assets: objects". Event domain `assets.object.created|updated|deleted` dipublikasikan untuk automation (JSM-AST-043), history, dan webhook; daftarkan di katalog event [28](28-model-data-inti.md).

### JSM-AST-015 Object List View & Filtering

- **Deskripsi:** Tampilan daftar object per object type dengan filter dasar, AQL, pilihan kolom, sorting, pagination, dan panel detail.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema user ke atas; Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. List menampilkan tabel object dengan kolom attribute yang dapat dipilih (column picker) dan diurutkan ulang dengan drag; pilihan kolom tersimpan per pengguna per object type.
  2. Basic filter (attribute, operator, nilai) dapat ditambah lebih dari satu dan dikonversi menjadi AQL yang dapat dilihat.
  3. ⚠ Advanced mode menerima input AQL langsung dengan validasi syntax saat diketik dan autocomplete nama attribute, operator, dan fungsi.
  4. Klik header kolom mengurutkan ASC/DESC; sorting diterjemahkan ke `ORDER BY` AQL.
  5. List memakai pagination (default 25 per halaman, opsi 50/100) dan menampilkan total jumlah object.
  6. Memilih baris membuka panel detail object di samping tanpa meninggalkan list.
  7. Tree object type menyediakan pencarian cepat berdasarkan label/key.
  8. Toggle "Tampilkan object dari child type" memuat object turunan di list parent.
  9. Object schema manager dapat menetapkan kolom default per object type untuk pengguna yang belum menyimpan pilihan.
  10. ⚠ Filter AQL dapat disimpan sebagai filter favorit pribadi per schema dan dibuka dari daftar "Filter tersimpan".
  11. ⚠ Nilai sel dapat diedit inline dari list oleh pengguna dengan hak edit, dengan validasi yang sama seperti object view.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Parser AQL dari JSM-AST-024; pola tabel dan column picker selaras dengan navigator di [13](13-pencarian-filter-ekspor.md).

### JSM-AST-016 Global Object Search

- **Deskripsi:** Pencarian object lintas schema dengan teks bebas atau AQL, terbatas pada schema yang boleh diakses pengguna.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dan setiap pengguna dengan role schema.
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian teks bebas mencocokkan label, key, dan nilai attribute teks (prefix dan substring, case-insensitive).
  2. Mode AQL menjalankan query lintas semua schema yang dapat diakses pengguna.
  3. Hasil dapat difilter per schema dan per object type dengan facet berisi jumlah hasil.
  4. Hasil hanya memuat object dari schema/object type tempat pengguna punya role efektif; tidak ada kebocoran jumlah total dari schema lain.
  5. Hasil pencarian menyediakan aksi bulk dan ekspor CSV (JSM-AST-017).
  6. Pencarian berbasis AI tersedia sebagai mode tambahan bila toggle AI aktif (JSM-AST-047).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pencarian teks memakai indeks full-text PostgreSQL (`tsvector`/`pg_trgm`) di tabel nilai attribute; persyaratan latensi mengikuti [27](27-nfr-arsitektur-platform.md).

### JSM-AST-017 Bulk Object Actions

- **Deskripsi:** Operasi massal pada kumpulan object hasil filter atau AQL: edit, delete, ekspor, dan cetak QR.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Object schema developer dan manager (edit/delete); Object schema user (ekspor).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat memilih semua object hasil filter (lintas halaman) atau sebagian lewat checkbox; jumlah terpilih selalu ditampilkan.
  2. ⚠ Bulk edit mengubah satu atau lebih attribute dengan mode Set/Replace untuk nilai tunggal dan Add/Remove untuk attribute multi-value; validasi unique dan cardinality dijalankan per object dan kegagalan dilaporkan per object.
  3. Bulk delete meminta konfirmasi dengan jumlah object serta jumlah inbound reference yang akan hilang.
  4. Ekspor CSV menawarkan pilihan "kolom yang tampil" atau "semua attribute", memakai UTF-8 dengan header nama attribute.
  5. Bulk print QR code menghasilkan PDF berisi label QR object terpilih (JSM-AST-021).
  6. ⚠ Satu operasi bulk dibatasi oleh batas bulk di JSM-AST-004 (default 1.000 object).
  7. Operasi bulk berjalan asinkron sebagai job dengan progress indicator, hasil akhir (sukses/gagal per object), dan dapat ditinggalkan tanpa membatalkan job.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Infrastruktur job asinkron dan endpoint progress bersama ([27](27-nfr-arsitektur-platform.md), JSM-AST-045 butir 12).

### JSM-AST-018 Object History

- **Deskripsi:** Log audit perubahan object: siapa, kapan, attribute apa, nilai lama dan baru, serta sumber perubahan.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Semua pengguna yang dapat melihat object; Assets admin untuk audit.
- **Aturan perilaku (acceptance criteria):**
  1. History mencatat pembuatan object, setiap perubahan nilai attribute (nilai lama → baru), serta penambahan/penghapusan reference.
  2. Setiap entri mencatat sumber perubahan: pengguna (UI), import (nama import structure), automation (nama rule), atau API (nama token/aplikasi).
  3. Perubahan attachment dan comment tercatat di history.
  4. ⚠ History dapat difilter berdasarkan attribute dan pelaku serta diurutkan terbaru/terlama.
  5. REST API `GET /object/{id}/history` mengembalikan entri history berurutan dengan pagination.
  6. Entri history tetap utuh walau pengguna pelakunya dinonaktifkan atau dihapus (nama tampil sebagai "Pengguna tidak aktif").
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** History disimpan append-only (`object_history`), tidak dapat diedit atau dihapus kecuali ikut terhapus saat object dihapus; retensi mengikuti kebijakan data di [22](22-administrasi-sistem.md).

### JSM-AST-019 Object Attachments

- **Deskripsi:** Lampiran file pada object (mis. invoice, kontrak, foto aset).
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Object schema developer ke atas (unggah/hapus); Object schema user (lihat/unduh).
- **Aturan perilaku (acceptance criteria):**
  1. Attachment dapat diunggah dengan drag-and-drop atau file picker di tab Attachments object view; beberapa file sekaligus didukung.
  2. Daftar attachment menampilkan nama, ukuran, pengunggah, dan tanggal.
  3. Gambar dan PDF dapat di-preview; semua tipe dapat diunduh.
  4. Attachment dapat dihapus oleh pengunggahnya (bila masih memiliki hak edit) dan oleh object schema manager.
  5. ⚠ Batas ukuran file mengikuti setting batas attachment site; unggahan yang melampaui ditolak dengan pesan yang menyebut batas.
  6. Penambahan dan penghapusan attachment tercatat di history object.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai layanan penyimpanan attachment yang sama dengan work item ([08](08-komentar-lampiran-aktivitas.md)); setting batas di [22](22-administrasi-sistem.md).

### JSM-AST-020 Object Comments

- **Deskripsi:** Komentar pada object untuk catatan kolaboratif terkait CI/aset.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Object schema user ke atas (menambah komentar); penulis dan object schema manager (edit/hapus).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna yang dapat melihat object dapat menambah komentar teks pada object.
  2. ⚠ Penulis dapat mengedit/menghapus komentarnya sendiri; object schema manager dapat menghapus komentar siapa pun.
  3. ⚠ Komentar mendukung mention pengguna yang memicu notifikasi in-app ke pengguna yang disebut bila ia punya akses ke object.
  4. Komentar ditampilkan kronologis dengan nama penulis dan waktu.
  5. Penambahan, perubahan, dan penghapusan komentar tercatat di history object.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen editor dan mention memakai pola [07](07-editor-konten-kaya.md) dan [08](08-komentar-lampiran-aktivitas.md); notifikasi in-app di [19](19-notifikasi-email.md).

### JSM-AST-021 Object QR Codes

- **Deskripsi:** Pembuatan dan pencetakan QR code object untuk identifikasi aset fisik dan pemindaian dari perangkat mobile.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dan Object schema developer (cetak); siapa pun yang memindai tetap tunduk permission.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap object memiliki QR code yang mengenkode deep link object detail (`/assets/object/<KEY>`).
  2. QR code tunggal dapat dicetak dari object view (dialog print/PDF).
  3. Bulk print menghasilkan lembar label berisi banyak QR sekaligus dari list atau hasil pencarian.
  4. ⚠ Label di bawah QR memuat key dan label object; ukuran label dapat dipilih dari preset (kecil, sedang, besar).
  5. ⚠ Pemindaian QR dengan kamera perangkat mobile membuka object di tampilan responsif SPA.
  6. Membuka object lewat QR tetap memerlukan login dan role schema; pengguna tanpa akses melihat halaman "Tidak punya akses", bukan data object.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** QR dibangkitkan di server (pustaka Go) agar konsisten di cetak tunggal dan bulk.

### JSM-AST-022 References & Dependency Graph Viewer

- **Deskripsi:** Visualisasi relasi object (inbound/outbound) dalam bentuk graph untuk analisis dampak dan dependency.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Object schema user ke atas; Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Object view menampilkan daftar inbound dan outbound reference, dikelompokkan per reference type dengan jumlah per kelompok.
  2. Graph viewer interaktif menampilkan node object dengan icon object type dan edge berwarna sesuai reference type, dengan arah panah.
  3. Node dapat di-expand untuk menelusuri level relasi berikutnya; node yang sudah tampil tidak diduplikasi.
  4. Graph dapat difilter berdasarkan reference type dan arah (inbound/outbound/keduanya).
  5. Klik node membuka object detail (panel atau halaman).
  6. ⚠ Opsi menampilkan connected work items sebagai node khusus di graph.
  7. ⚠ Graph tingkat schema menampilkan relasi antar object type (berdasarkan attribute Object) untuk dokumentasi model.
  8. ⚠ Jumlah node yang dirender dibatasi (default 200); saat batas tercapai, graph menampilkan pesan untuk mempersempit filter atau expand bertahap.
  9. Object yang tidak boleh dilihat pengguna ditampilkan sebagai node anonim "Object terbatas" tanpa label dan tidak dapat di-expand.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 9 adalah penegasan aturan permission (JSM-AST-002) pada graph. Data graph diambil per level lewat endpoint `referenceinfo` agar tidak memuat seluruh graph sekaligus.

### JSM-AST-023 Connected Jira Issues on Objects

- **Deskripsi:** Object menampilkan daftar work item (issue/request) yang merujuknya lewat field Assets objects, sehingga riwayat incident/permintaan per CI dapat dilihat.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dan pengguna dengan role schema; visibilitas tiap work item mengikuti browse permission project.
- **Aturan perilaku (acceptance criteria):**
  1. Object view memiliki tab "Connected tickets" (Work item terkait).
  2. Daftar menampilkan kolom key, summary, status, work type, dan project, diurutkan dari pembaruan terbaru.
  3. Klik baris membuka work item di tab yang sama atau baru.
  4. Hanya work item yang boleh di-browse pengguna yang tampil; jumlah total tidak menyertakan work item tersembunyi.
  5. Fungsi AQL `connectedTickets(JQL)` mencari object berdasarkan work item terkait (rinci di JSM-AST-025).
  6. REST API `GET /objectconnectedtickets/{objectId}/tickets` mengembalikan daftar work item terkait dengan field yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Relasi dibangun dari tabel indeks `issue_asset_link` (issue_id, field_id, object_id) yang diperbarui saat nilai field Assets objects berubah; browse permission di [03](03-izin-peran-keamanan.md).

## Assets Query Language (AQL)

### JSM-AST-024 Assets Query Language (AQL) – Grammar, Operators & Sorting

- **Alias inventaris:** Assets Query Language (AQL)
- **Deskripsi:** Bahasa query untuk mencari object berdasarkan attribute, tipe, schema, dan relasi; dipakai di UI, field Assets objects, automation, import mapping, JQL, dan API.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Object schema manager, Integrator/Developer API; hasil selalu dibatasi role pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Struktur dasar `<attribute> <operator> <value>` dapat digabung dengan `AND`/`OR` dan dikelompokkan dengan tanda kurung; `AND` berprioritas lebih tinggi dari `OR`.
  2. Operator yang didukung: `=`, `!=`, `<`, `>`, `<=`, `>=`, `LIKE`, `NOT LIKE`, `IN()`, `NOT IN()`, `IS EMPTY`, `IS NOT EMPTY`, `STARTSWITH`, `ENDSWITH`, `HAVING`, `NOT HAVING`; operator perbandingan numerik/tanggal ditolak untuk attribute teks dengan galat tipe.
  3. ⚠ Keyword yang didukung: `objectType`, `objectTypeId`, `objectSchema`, `objectSchemaId`, `Key`, `Label`, `objectId`, `Created`, `Updated`, dan `anyAttribute` (mencari nilai di attribute mana pun).
  4. Nama atau nilai yang mengandung spasi wajib diapit tanda kutip ganda; karakter khusus (kutip, backslash) di-escape dengan backslash.
  5. ⚠ Nama attribute, keyword, dan operator bersifat case-insensitive; perbandingan nilai teks dengan `=` juga case-insensitive, sedangkan `LIKE` mencocokkan substring case-insensitive.
  6. Dot notation menelusuri attribute object terkait (mis. `"Owner"."Department" = HR`), maksimum 3 tingkat.
  7. Attribute dapat dirujuk dengan nama maupun ID numerik (mis. `"123" = "abc"`); ID menang bila nama ambigu.
  8. `ORDER BY <attribute> ASC|DESC` didukung, termasuk pengurutan berdasarkan attribute object terkait via dot notation; tanpa ORDER BY, hasil diurutkan berdasarkan label.
  9. Perbandingan tanggal relatif didukung (mis. `Created > now(-7d)`, unit `m`, `h`, `d`, `w`, `M`, `y`).
  10. Query tidak valid menghasilkan galat 400 dengan pesan yang menyebut posisi token dan penyebabnya (mis. "Attribute 'Ownr' tidak ditemukan pada object type Laptop").
  11. Hasil hanya berisi object yang boleh diakses pemanggil; schema tanpa akses diperlakukan seolah tidak ada.
  12. IQL diterima sebagai nama lama/alias AQL di semua titik masuk (endpoint `/iql` dan `iqlFunction()` sebagai alias, lihat JSM-AST-048).
  13. Contoh query `objectType = "Laptop" AND Owner = currentUser()` dievaluasi dengan benar dan menjadi kasus uji acuan.
  14. Kombinasi operator lengkap (`=`, `!=`, `<`, `>`, `<=`, `>=`, `LIKE`, `NOT LIKE`, `IN`, `NOT IN`, `IS EMPTY`, `IS NOT EMPTY`, `STARTSWITH`, `ENDSWITH`) diuji untuk setiap tipe attribute yang relevan.
  15. Keyword `objectSchema`, `objectType`, `objectTypeId`, `Key`, `Label`, `Created`, `Updated`, `ORDER BY`, dan `AND`/`OR` tersedia di autocomplete editor AQL.
  16. Fungsi referensi `inR()`, `outR()`, `inboundReferences()`, `outboundReferences()`, dan `connectedTickets()` tersedia di grammar (rinci di JSM-AST-025).
  17. Fungsi pengguna/waktu `currentUser()`, `user()`, `group()`, `now()`, `startOfDay()`, dan sejenisnya tersedia di grammar (rinci di JSM-AST-025).
  18. `HAVING` dipakai untuk query relasi terhadap hasil fungsi references/connected tickets (mis. `object HAVING inboundReferences(objectType = Incident)`).
  19. ⚠ Placeholder konteks work item (mis. `${customfield_10010}`, `${Reporter}`) diterima di AQL yang dievaluasi dalam konteks field (rinci di JSM-AST-025).
  20. JQL menerima AQL lewat `"Nama field Assets" in aqlFunction("AQL")` (rinci di JSM-AST-028).
  21. Konversi bahasa alami ke AQL tersedia sebagai bantuan AI opsional (rinci di JSM-AST-047).
  22. Validasi AQL berjalan di editor (sebelum eksekusi) dan di server; pesan galat ditampilkan dalam bahasa sesuai `Accept-Language` (EN/ID).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 13-22 berasal dari entri lampiran "Assets Query Language (AQL)". Parser (lexer + recursive descent) dan compiler AQL → SQL ditempatkan di `pkg/aql` dan diuji dengan tabel kasus; pesan galat EN/ID memakai `internal/shared/constants/error.go`. Pola bahasa sejalan dengan JQL di [12](12-jql-bahasa-kueri.md) tetapi grammar-nya terpisah.

### JSM-AST-025 AQL Functions & Placeholders

- **Deskripsi:** Fungsi bawaan AQL untuk pengguna, group, tanggal, reference, dan keterkaitan work item, serta placeholder konteks work item.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Object schema manager, Admin Situs (konfigurasi field).
- **Aturan perilaku (acceptance criteria):**
  1. Fungsi pengguna: `currentUser()` (pengguna yang mengeksekusi), `currentReporter()` (reporter work item/request dalam konteks field atau portal; kosong di luar konteks), dan `user("accountId")`.
  2. `group("nama")` mencocokkan attribute User yang anggotanya termasuk group tersebut atau attribute Group dengan group itu.
  3. Fungsi tanggal `now()`, `startOfDay()`, `endOfDay()`, `startOfWeek()`, `endOfWeek()`, `startOfMonth()`, `endOfMonth()`, `startOfYear()`, `endOfYear()` menerima offset (mis. `now(-2h)`, `startOfMonth(1M)`) dan dihitung dalam timezone pengguna.
  4. `inboundReferences(AQL, referenceTypes)` dan alias `inR()` mengembalikan object yang direferensikan oleh object hasil AQL dalam argumen, dengan filter reference type opsional.
  5. `outboundReferences(AQL, referenceTypes)` dan alias `outR()` mengembalikan object yang mereferensikan object hasil AQL, dengan filter reference type opsional.
  6. `connectedTickets(JQL)` mengembalikan object yang tertaut ke work item hasil JQL; JQL dievaluasi dengan permission pemanggil, dan JQL kosong berarti "punya work item terkait apa pun".
  7. `objectTypeAndChildren("nama")` menyertakan object dari type tersebut dan semua turunannya.
  8. `HAVING`/`NOT HAVING` dapat dikombinasikan dengan `inboundReferences`, `outboundReferences`, dan `connectedTickets` (mis. `object HAVING connectedTickets(status != Done)`).
  9. ⚠ Placeholder konteks work item tersedia saat AQL dievaluasi dalam konteks field: `${customfield_NNNNN}` (nilai field; untuk field Assets memakai key object), `${customfield_NNNNN.label}`, `.key`, `.id`, `${Reporter}`, `${Assignee}`, dan `${Project.key}`; placeholder tak dikenal menghasilkan galat validasi saat konfigurasi field disimpan.
  10. ⚠ Placeholder yang nilainya kosong menghasilkan klausa yang tidak cocok dengan object mana pun (hasil kosong), bukan mengabaikan filter; perilaku ini didokumentasikan di help text konfigurasi field.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konteks field dan reporter dari [05](05-konfigurasi-work-item-skema.md) dan [06](06-work-item-inti.md); evaluasi `connectedTickets` memanggil engine JQL ([12](12-jql-bahasa-kueri.md)) dan indeks `issue_asset_link` (JSM-AST-023).

## Integrasi dengan work item, JQL, portal & ITSM

### JSM-AST-026 Assets Objects Custom Field

- **Alias inventaris:** Assets custom field & portal usage
- **Deskripsi:** Custom field "Assets objects" untuk mengaitkan object ke work item/request, dikonfigurasi per context dengan schema, filter AQL, cardinality, dan attribute yang ditampilkan.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (membuat field dan context); Agent Layanan dan Customer (mengisi); visibilitas object mengikuti role schema atau setting "izinkan pengguna lain".
- **Aturan perilaku (acceptance criteria):**
  1. Tipe custom field "Assets objects" tersedia di halaman custom field; field diatur context/project-nya dan dipasang ke screen seperti custom field lain.
  2. Setiap context wajib memilih satu Object schema.
  3. Filter scope (AQL) membatasi object yang dapat dipilih (mis. `objectType = Laptop AND Status = "In Use"`).
  4. Filter issue scope (AQL) mendukung placeholder nilai field lain untuk cascading (mis. Laptop berdasarkan Employee terpilih: `objectType = Laptop AND Owner = ${customfield_10020.label}`); picker diperbarui saat field sumber berubah tanpa reload.
  5. Admin memilih attribute yang ditampilkan di issue view dan di dropdown pencarian picker.
  6. Opsi "Field dapat menyimpan banyak object" menentukan single vs multi; multi tunduk batas JSM-AST-004 butir 7.
  7. ⚠ Default object dapat ditetapkan per context dan juga dipakai saat field tampil di portal customer.
  8. Picker mencari berdasarkan label, key, dan attribute yang ditampilkan (minimal 1 karakter, debounce), dengan hasil terbatas filter scope.
  9. Klik object di issue view membuka panel detail object (bila pengguna punya akses) atau tooltip label (bila tidak).
  10. Field dapat dipakai di create/edit/transition screen, request type form, dan sebagai kolom queue.
  11. Nilai field dapat diubah lewat bulk edit work item dan action automation "Edit issue".
  12. ⚠ Pengguna tanpa role schema hanya dapat melihat dan memilih object bila schema mengaktifkan "Izinkan pengguna lain memilih object"; bila tidak, picker kosong dengan pesan "Anda tidak punya akses ke schema ini".
  13. Field yang konfigurasinya belum lengkap (schema belum dipilih) menampilkan pesan validasi kepada admin di halaman konfigurasi dan disembunyikan dari screen sampai lengkap.
  14. Pilihan single vs multiple object diuji pada create, edit, dan API (`PUT` nilai array berisi lebih dari satu ditolak untuk field single).
  15. Attribute object yang dipilih tampil di issue view di bawah label object dalam format ringkas.
  16. Field dapat ditampilkan di portal agar customer memilih asetnya (rinci di JSM-AST-027).
  17. Field mendukung default value dan dapat diatur wajib/opsional per screen atau request type.
  18. Field Assets dapat dipakai di Forms (rinci di JSM-AST-027 butir 7).
  19. Work item yang mengisi field ini muncul di tab Connected tickets object (JSM-AST-023).
  20. Approver request dapat diambil dari attribute User object yang dipilih di field (mis. Manager dari object Department) lewat konfigurasi approval di modul ITSM.
  21. Konfigurasi field selalu menyimpan tiga unsur per context: schema, filter scope (AQL), dan filter issue scope dinamis (opsional); ketiganya dapat dilihat dan diubah dalam satu layar konfigurasi context.
- **Team-managed vs company-managed:** Di company-managed, field diatur lewat field context dan screen; di team-managed, field ditambahkan per project dengan konfigurasi schema/filter yang sama dan tidak berbagi context.
- **Catatan & ketergantungan:** Butir 14-21 berasal dari entri lampiran "Assets custom field & portal usage". Mekanik custom field di [05](05-konfigurasi-work-item-skema.md); pemasangan field di request type form di JSM-PRJ-025 butir 9 dan di form builder di JSM-PRJ-035 ([29](29-jsm-service-project-request-type-formulir.md)); action "Edit issue" (butir 11) di AUTO-039; approver dari attribute object (butir 20) di JSM-ITSM-011 butir 13 dan JSM-ITSM-012 butir 4 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)); field Affected services (JSM-ITSM-005) adalah varian field ini yang terkunci ke schema Services. Nilai disimpan sebagai daftar object ID + indeks `issue_asset_link`.

### JSM-AST-027 Assets in Request Forms & Customer Portal

- **Deskripsi:** Customer dapat memilih object (mis. laptop miliknya) saat mengajukan request di portal lewat field Assets objects di request type.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (service project admin) mengonfigurasi request type; Customer/Pelapor mengisi.
- **Aturan perilaku (acceptance criteria):**
  1. Field Assets objects dapat ditambahkan ke request type dengan nama tampilan dan help text khusus customer.
  2. Picker portal menampilkan label object (dan attribute yang dikonfigurasi) serta kotak pencarian.
  3. Filter scope AQL dengan `currentReporter()` membatasi object yang tampil hanya milik pelapor (mis. `objectType = Laptop AND Owner = currentReporter()`); saat agent membuat request atas nama customer, `currentReporter()` mengacu ke customer tersebut.
  4. Picker portal hanya berfungsi bila object schema mengaktifkan akses customer portal; bila tidak, field tidak ditampilkan di portal dan admin melihat peringatan di konfigurasi request type.
  5. Validasi required request type berlaku untuk field Assets; submit tanpa pilihan ditolak dengan pesan field.
  6. Object terpilih tampil di detail request portal sebagai label saja (tanpa tautan ke object).
  7. ⚠ Field Assets dapat dipakai di Forms request type bila pertanyaan form dipetakan ke field Assets objects; pilihan form tersinkron ke nilai field.
  8. Customer tidak dapat membuka object detail, list, graph, maupun API Assets; permintaan langsung ke endpoint Assets mengembalikan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi field request type di JSM-PRJ-025 (butir 9 memuat field Assets objects); Forms di JSM-PRJ-034, tipe pertanyaan Assets di form builder JSM-PRJ-035, dan pemetaan pertanyaan ke field di JSM-PRJ-038 — semuanya di modul [29](29-jsm-service-project-request-type-formulir.md); request fulfillment perangkat di JSM-ITSM-003 ([32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)); portal di modul [30](30-jsm-help-center-portal-pelanggan-kanal.md).

### JSM-AST-028 Assets in JQL & Issue Search

- **Deskripsi:** Pencarian work item berdasarkan object yang terkait lewat field Assets objects, termasuk via AQL.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Admin Situs; hasil mengikuti browse permission work item dan role Assets pemanggil.
- **Aturan perilaku (acceptance criteria):**
  1. Fungsi JQL `aqlFunction("AQL")` didukung pada field Assets objects (mis. `"Affected hardware" in aqlFunction("Status = Broken")`), dievaluasi dengan permission Assets pemanggil.
  2. ⚠ `iqlFunction()` diterima sebagai alias legacy dengan perilaku identik.
  3. ⚠ Field Assets dapat dicari langsung dengan object key atau label (mis. `"Affected hardware" = "ITSM-123"`, `"Affected hardware" ~ "MacBook"`).
  4. Field Assets dapat dijadikan kolom di navigator work item dan queue, menampilkan label object.
  5. Queue service dapat difilter berdasarkan object (mis. semua incident untuk service tertentu) memakai JQL di atas.
  6. ⚠ Hasil `aqlFunction` dibatasi (default 10.000 object); bila terlampaui, JQL gagal dengan pesan yang meminta AQL dipersempit.
  7. ⚠ Editor JQL menyediakan autocomplete nilai field Assets (label/key) dan sintaks `aqlFunction`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Registrasi fungsi JQL kustom di [12](12-jql-bahasa-kueri.md); field JQL Service Management (termasuk `aqlFunction()`/`iqlFunction()`) dicatat juga di JSM-AI-013 ([35](35-jsm-ai-otomasi-api-admin-csm.md)) dengan perilaku rinci dimiliki fitur ini; kolom navigator di [13](13-pencarian-filter-ekspor.md); queue di modul [31](31-jsm-antrian-sla-knowledge-base-laporan.md).

### JSM-AST-029 Assets in ITSM Practices (Integration Points)

- **Deskripsi:** Titik integrasi Assets dengan proses ITSM inti (incident, problem, change, request) dan fitur inti produk (queue, SLA, dashboard).
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek (service project admin), Manajer Layanan/ITSM (change manager).
- **Aturan perilaku (acceptance criteria):**
  1. Template project ITSM menyertakan field Assets "Affected CI/hardware" pada work type Incident; CI terkait tampil di issue view untuk analisis dampak.
  2. Problem dapat dikaitkan dengan CI yang sama dengan banyak incident; tab Connected tickets object memperlihatkan semua incident dan problem terkait.
  3. ⚠ Change dapat menautkan CI/service terdampak; penilaian risiko change dapat memakai attribute object (mis. tier service) sebagai input rule.
  4. Request fulfilment: aset yang dipilih customer tersimpan di request, dan template automation memperbarui status aset (mis. In Use) setelah request diselesaikan.
  5. Onboarding/offboarding HR: template automation membuat object Employee atau menugaskan/melepas aset ke employee.
  6. Queue dan SLA dapat difilter berdasarkan field Assets lewat JQL `aqlFunction`.
  7. Gadget dashboard berbasis filter mendukung filter yang memakai `aqlFunction`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Proses ITSM di modul [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md): field incident (termasuk Affected services) di JSM-ITSM-016, field Affected services di JSM-ITSM-005, problem di JSM-ITSM-033, tautan incident/problem/change di JSM-ITSM-034, field Affected CI pada change di JSM-ITSM-036 butir 21, risk assessment berbasis attribute di JSM-ITSM-042, request fulfillment di JSM-ITSM-003, template automation ITSM di JSM-ITSM-048; SLA/queue di modul [31](31-jsm-antrian-sla-knowledge-base-laporan.md); gadget di [17](17-dashboard-gadget.md); template automation di [20](20-otomasi.md) dan JSM-AST-043.

### JSM-AST-030 Services in Assets

- **Deskripsi:** Service (service registry) disimpan sebagai object di schema "Services" dalam Assets, sehingga dapat direlasikan dengan CI lain.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek (service project admin), Assets admin.
- **Aturan perilaku (acceptance criteria):**
  1. Saat Services diaktifkan, sistem membuat schema "Services" dengan object type Service dan attribute bawaan (Name, Description, Service tier, Owner team, Service type).
  2. Service yang dibuat/diubah/dihapus dari halaman Services tersinkron dua arah ke object Assets dalam transaksi yang sama (atau via event dengan retry yang idempoten).
  3. ⚠ Object schema manager dapat menambah attribute kustom ke object type Service; attribute bawaan tidak dapat dihapus atau diubah tipenya.
  4. Service dapat direlasikan dengan CI (server, aplikasi) lewat attribute Object.
  5. Field "Affected services" di incident/change memakai object dari schema ini.
  6. Struktur sistem schema Services (object type Service, attribute bawaan) read-only dan schema tidak dapat dihapus.
  7. Dependency service-to-service tampil di graph Assets dan di halaman detail service.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman Services dan registry dirinci di JSM-ITSM-004, field Affected services di JSM-ITSM-005, peran service di JSM-ITSM-006, service tier di JSM-ITSM-007, dependency service di JSM-ITSM-008, dan permission service via role schema di JSM-ITSM-009 (modul [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)); ⚠ JSM-ITSM-004 butir 4 menyebut dua schema bawaan (Business Portfolio dan Software Registry) sedangkan butir 1 di sini menyebut satu schema "Services" — nama schema bawaan diputuskan bersama modul 32 (lihat Pertanyaan terbuka). Schema Services ditandai sistem di JSM-AST-005 butir 11.

## Import, sinkronisasi & kualitas data

### JSM-AST-031 Assets Import – Import Structures (Framework)

- **Alias inventaris:** Assets import & integrations
- **Deskripsi:** Kerangka import Assets berupa import structure per schema yang mendefinisikan tipe sumber, mapping, eksekusi, dan log.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager dan Assets admin; Object schema developer hanya dapat menjalankan "Sync now".
- **Aturan perilaku (acceptance criteria):**
  1. Tab Import di schema configuration menampilkan daftar import structure dengan nama, tipe, status (enabled/disabled), waktu sync terakhir, dan hasil terakhir.
  2. ⚠ Create import meminta pilihan tipe: CSV, JSON, Assets Discovery, External import/app, dan Data Manager.
  3. ⚠ Tipe import pengguna/group site (membuat object dari akun dan group) tidak disediakan di fase awal; kebutuhan ini dipenuhi lewat CSV/JSON atau integrasi Entra ID.
  4. Konfigurasi umum memuat Name, Description, Concatenator (pemisah saat beberapa locator digabung), kebijakan Empty values (Ignore/Remove), dan kebijakan Unknown values (Ignore/Add) untuk attribute Select dan Status.
  5. Opsi "Buat struktur otomatis" membuat object type dan attribute dari struktur sumber sekaligus mengisi mapping awal.
  6. Import structure dapat di-enable/disable; import disabled tidak dapat dijalankan manual maupun terjadwal.
  7. Tombol "Sync now" menjalankan import secara manual dan menampilkan status progress (persentase, baris diproses).
  8. Import yang sedang berjalan dapat dibatalkan; object yang sudah diproses tetap tersimpan dan log mencatat status "Dibatalkan".
  9. Import history menyimpan setiap eksekusi dengan jumlah created, updated, removed, unchanged, dan error per baris beserta pesan galatnya; log dapat diunduh.
  10. Import structure dapat diubah dan dihapus; menghapus import structure tidak menghapus object yang pernah diimpor.
  11. Import token untuk sumber eksternal dapat dibangkitkan, disalin (sekali tampil), dan dicabut; token terikat ke satu import structure.
  12. ⚠ Hanya satu eksekusi import per schema yang berjalan bersamaan; permintaan berikutnya ditolak dengan pesan "Import lain sedang berjalan".
  13. Import CSV dengan mapping kolom didukung sebagai tipe dasar (rinci di JSM-AST-033).
  14. Import JSON didukung (rinci di JSM-AST-034).
  15. Hasil Assets Discovery dapat diimpor lewat tipe import Discovery (rinci di JSM-AST-038).
  16. ⚠ Integrasi sumber cloud (AWS, Azure, Google Cloud) dan MDM (Jamf, Intune) tidak dibangun native; sumber tersebut masuk lewat External import (JSM-AST-036) atau adapter Data Manager (JSM-AST-039).
  17. Mapping object type/attribute dan identifier menentukan update vs create (rinci di JSM-AST-032).
  18. Penanganan object yang hilang dari sumber mendukung Ignore, Update, dan Remove (rinci di JSM-AST-032 butir 7).
  19. Scheduled import dan log hasilnya didukung (rinci di JSM-AST-035).
  20. External imports REST API tersedia untuk integrasi kustom (rinci di JSM-AST-036).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 13-20 berasal dari entri lampiran "Assets import & integrations". Eksekusi import berjalan sebagai job latar belakang dengan antrean per schema ([27](27-nfr-arsitektur-platform.md)); object hasil import tetap tunduk kuota JSM-AST-004.

### JSM-AST-032 Import Object Type & Attribute Mapping

- **Deskripsi:** Pemetaan data sumber ke object type dan attribute Assets, termasuk identifier, reference, dan penanganan object yang hilang dari sumber.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager.
- **Aturan perilaku (acceptance criteria):**
  1. Object type mapping memilih object type target dan Selector opsional (subset data, mis. path JSON atau filter baris).
  2. Attribute mapping memetakan data locator ke attribute; beberapa locator dapat digabung memakai Concatenator.
  3. Minimal satu attribute per object type mapping wajib ditandai Identifier; object existing dengan nilai identifier sama di-update, selain itu dibuat baru.
  4. Attribute Object diisi lewat AQL dengan placeholder locator (mis. `Name = ${Manufacturer}`); bila AQL tidak menemukan object, nilai dikosongkan dan baris dicatat sebagai peringatan.
  5. ⚠ Attribute User/Group dipetakan berdasarkan email atau accountId untuk User dan nama group untuk Group; nilai yang tidak cocok dicatat sebagai peringatan.
  6. ⚠ Transformasi nilai dengan regex (capture group) dapat diterapkan per attribute mapping sebelum disimpan.
  7. Missing objects handling per object type mapping: Remove (hapus object), Update (set attribute/status tertentu, mis. Status = Retired), atau Ignore.
  8. Empty values: Ignore mempertahankan nilai lama, Remove mengosongkan nilai attribute.
  9. Mapping child object type mewarisi mapping dan attribute parent bila inheritance aktif.
  10. ⚠ Filter AQL pada mapping membatasi object existing yang dapat di-update atau di-remove oleh import (mis. hanya object dengan `Source = "Intune"`).
  11. Mapping per object type dapat di-enable/disable tanpa menghapus konfigurasinya.
  12. Urutan eksekusi diatur otomatis agar object type yang direferensikan diimpor lebih dulu; dependency melingkar dilaporkan sebagai galat validasi.
  13. Validasi mapping (identifier ada, tipe kompatibel, AQL valid) dijalankan sebelum sync; mapping tidak valid mencegah "Sync now" dengan daftar masalah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Placeholder dan AQL dari JSM-AST-024/JSM-AST-025.

### JSM-AST-033 CSV Import

- **Deskripsi:** Import object dari file CSV lewat import structure.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ File CSV diunggah di konfigurasi import, tunduk batas ukuran JSM-AST-004 butir 6 (default 100 MB).
  2. Konfigurasi delimiter (koma, titik koma, tab, kustom), encoding (UTF-8 default), header row, serta format date dan datetime.
  3. Data locator adalah nama kolom header.
  4. Object type dan attribute dapat dibangkitkan otomatis dari header (tipe default Text, dapat diubah sebelum dibuat).
  5. Sinkronisasi berikutnya dilakukan dengan mengunggah ulang file; file terakhir disimpan beserta waktu unggahnya.
  6. ⚠ Nilai multi-value dalam satu sel dipisah dengan pemisah yang dapat dikonfigurasi (default `|`).
  7. Laporan baris gagal mencantumkan nomor baris, kolom, dan alasan (unique, required, format) serta dapat diunduh sebagai CSV.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pola parser CSV dapat dipakai bersama impor CSV work item ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JSM-AST-034 JSON Import

- **Deskripsi:** Import object dari file JSON dengan selector path untuk struktur bersarang.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager.
- **Aturan perilaku (acceptance criteria):**
  1. File JSON (UTF-8) diunggah di konfigurasi import; JSON tidak valid ditolak dengan posisi galat.
  2. Setiap object type mapping memiliki Selector (path, mis. `data.devices`) untuk memilih array object.
  3. Data locator berbasis key atau path JSON relatif terhadap elemen selector (mis. `hardware.serial`).
  4. ⚠ Nested object dan array didukung hingga kedalaman 10; array nilai primitif dipetakan ke attribute multi-value.
  5. Mapping dapat dibangkitkan otomatis dari contoh data (elemen pertama array).
  6. Tipe data JSON (number, boolean, string tanggal ISO 8601) dikonversi ke tipe attribute; konversi gagal dicatat sebagai galat baris.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format path selaras dengan JSONPath sederhana (dot notation); keputusan dukungan JSONPath penuh dicatat di Pertanyaan terbuka.

### JSM-AST-035 Scheduled Imports

- **Deskripsi:** Penjadwalan import structure agar berjalan otomatis secara berkala.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Object schema manager.
- **Aturan perilaku (acceptance criteria):**
  1. Jadwal dapat diaktifkan per import structure.
  2. ⚠ Jadwal memakai interval preset (setiap jam, harian, mingguan) atau ekspresi cron 5 field dengan interval minimum 1 jam.
  3. Jadwal memiliki timezone eksplisit (default timezone site).
  4. Setiap eksekusi terjadwal tercatat di import history dengan penanda "Terjadwal".
  5. ⚠ Eksekusi terjadwal dilewati (dicatat "Dilewati") bila eksekusi sebelumnya masih berjalan.
  6. Untuk External import dan Discovery, jadwal dikendalikan dari sisi klien yang memakai token; UI menampilkan informasi ini dan tidak menawarkan jadwal server.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Untuk CSV/JSON, scheduled import memproses file terakhir yang diunggah atau URL sumber bila fitur URL diputuskan (Pertanyaan terbuka). Scheduler bersama dengan scheduled trigger automation ([20](20-otomasi.md)).

### JSM-AST-036 External Imports (Imports REST API)

- **Deskripsi:** Framework import eksternal yang memungkinkan aplikasi atau skrip mengirim data ke Assets lewat Imports REST API dengan import token.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Integrator/Developer API (pemegang token); Assets admin/object schema manager membuat import structure dan token.
- **Aturan perilaku (acceptance criteria):**
  1. Tipe import "External import" terikat ke aplikasi/skrip lewat import token yang dibangkitkan di import structure.
  2. Alur API: ambil info import dengan token → kirim schema dan mapping (JSON) → mulai execution → kirim data per chunk → tandai selesai (submit), dengan proses diakhiri pemrosesan asinkron.
  3. Mapping didefinisikan klien dalam JSON dan dapat dilihat (read-only) di UI import structure.
  4. Progress dan status execution tersedia via API (queued, ingesting, processing, done, failed) beserta ringkasan hasil.
  5. Mencabut token langsung memutus integrasi; panggilan berikutnya mengembalikan 401.
  6. ⚠ Ukuran chunk maksimum 5 MB dan rate limit per token mengikuti kebijakan API ([23](23-api-publik-webhook.md)); pelanggaran mengembalikan 413/429.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konektor siap pakai ke penyedia cloud, MDM, dan sistem ITAM lain termasuk ekosistem app dan tidak dibangun native (lihat Di luar cakupan).

### JSM-AST-037 Assets - Microsoft Entra ID Integration

- **Deskripsi:** Integrasi first-party yang mengimpor user, group, dan device dari Microsoft Entra ID ke Assets memakai framework external import.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin bersama admin tenant Entra ID (memberi consent).
- **Aturan perilaku (acceptance criteria):**
  1. Admin menghubungkan integrasi Entra ID melalui alur OAuth admin consent tenant; status koneksi dan tenant ID tampil di halaman integrasi.
  2. Integrasi membuat import structure bertipe app/external pada object schema target yang dipilih.
  3. ⚠ Admin memilih data yang diimpor: users, groups, dan devices (checkbox per jenis).
  4. ⚠ Mapping bawaan ke object type/attribute Assets (mis. User → Employee, Device → Laptop/Computer) dapat disesuaikan sebelum sync pertama.
  5. Sinkronisasi berjalan terjadwal (default harian) dan setiap eksekusi tercatat di riwayat/log import.
  6. Object yang dihapus atau dinonaktifkan di Entra ID diperlakukan sesuai kebijakan missing objects (JSM-AST-032 butir 7).
  7. Integrasi diputus dengan mencabut token atau menghapus koneksi; object yang sudah diimpor tetap ada.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dibangun di atas JSM-AST-036 sebagai konektor bawaan produk ini (bukan app pihak ketiga); berbeda dari adapter Entra ID/AD di Data Manager (JSM-AST-039).

### JSM-AST-038 Assets Discovery

- **Deskripsi:** Tool first-party yang memindai jaringan (agentless maupun berbasis agent), menemukan perangkat dan software, lalu mengirim hasilnya ke Assets; diunduh dan dijalankan di infrastruktur pelanggan.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin dan admin TI/operasi (menjalankan tool); object schema manager (import structure tujuan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Paket Assets Discovery (Discovery tool/scanner, Collector, Agent) dapat diunduh dari halaman Import untuk Linux, Windows, dan macOS.
  2. ⚠ Agentless scan mendukung SSH (Linux/Unix/macOS), WMI (Windows), SNMP (perangkat jaringan), dan API VMware.
  3. Discovery Agent dipasang di endpoint dan mengirim data ke Collector.
  4. Collector menggabungkan hasil dari agent dan scanner sebelum dikirim.
  5. Konfigurasi scan memuat IP range, exclusions, credentials per protokol, timeout, dan jumlah thread.
  6. Patterns berisi skrip bawaan per OS/perangkat (hardware, OS, software, patch, network interface) dan dapat ditambah pattern kustom.
  7. Scan dapat dijadwalkan dari Discovery tool.
  8. Hasil diunggah ke Assets via import token (tipe import Assets Discovery) atau diekspor ke file untuk diunggah manual.
  9. Template mapping Discovery bawaan memetakan hasil ke object type (Host, Software, Network Device, dll.), selaras dengan template JSM-AST-006 butir 4.
  10. Object diidentifikasi unik berdasarkan serial number/hostname sehingga scan berikutnya meng-update object existing.
  11. Log scan dan diagnostik koneksi (host tidak terjangkau, kredensial gagal) tersedia di tool.
  12. Komponen Discovery terdiri dari Discovery tool, Agent, dan Collector yang dapat dipasang terpisah.
  13. Scan pattern mencakup WMI, SSH, dan SNMP sebagai protokol minimum.
  14. Rentang IP dan kredensial dikonfigurasi per profil scan.
  15. Jadwal scan dapat diatur per profil scan.
  16. File hasil scan diimpor ke Assets melalui import configuration bertipe Discovery.
  17. Hasil mendeteksi hardware, software terpasang, dan perangkat jaringan.
  18. Tool diunduh dan dijalankan on-premise pelanggan; tidak ada pemindaian yang dijalankan dari server produk ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini menggabungkan dua entri inventaris bernama sama, "Assets Discovery" (lampiran JSM, butir 12-18) dan "Assets Discovery" (domain Assets, butir 1-11). Discovery adalah binary terpisah (Go) yang memakai Imports REST API (JSM-AST-036); pengembangannya dapat ditunda tanpa memblokir fitur lain.

### JSM-AST-039 Assets Data Manager

- **Deskripsi:** Modul yang mengumpulkan data aset dari banyak sumber lewat adapter, lalu membersihkan, merekonsiliasi, dan menganalisis kualitas data sebelum masuk ke Assets.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin dan pengelola aset TI (Manajer Layanan/ITSM).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Adapter tersedia untuk sumber umum (Active Directory, Intune, SCCM, Jamf, AWS, Azure, Qualys, Tenable, flat file); daftar adapter bawaan fase awal minimal flat file (CSV) dan Active Directory/Entra ID.
  2. Adapter client dapat dipasang on-prem untuk mengambil data dari sumber internal dan mengirimkannya ke Data Manager.
  3. Data source didefinisikan beserta jadwal pengambilan data (jobs).
  4. Object class (mis. Compute, People, Network, Software) menjadi model data terpadu tempat data sumber dipetakan.
  5. Cleansing rule mentransformasi nilai (replace, default value, format, regex, merge) dan dapat diuji pada contoh data sebelum diterapkan.
  6. Reconciliation mencocokkan record antar sumber dengan key dan prioritas sumber per attribute (sumber prioritas tertinggi menang).
  7. Analisis kualitas data menampilkan gap antar sumber (mis. perangkat ada di MDM tetapi tidak di AD), duplikat, dan coverage.
  8. Dashboard/report kualitas data menampilkan metrik per object class dan per sumber.
  9. Hasil rekonsiliasi diimpor ke Assets lewat import structure bertipe Data Manager dengan mapping ke object type.
  10. Jobs history dan log tersimpan per data source.
  11. Banyak sumber data dapat dihubungkan sekaligus melalui adapter.
  12. Cleansing rule mendukung normalisasi nilai (mis. nama vendor "Dell Inc." → "Dell").
  13. Rekonsiliasi menggabungkan record duplikat lintas sumber menjadi satu record emas (golden record).
  14. Analisis gap dan kualitas data dapat dijalankan sebelum publikasi.
  15. Hasil dipublikasikan ke Assets hanya setelah pengguna menekan "Publish" atau sesuai jadwal publikasi.
  16. ⚠ Data Manager hanya tersedia pada tier Enterprise (default) dengan opsi pengaktifan untuk Premium lewat feature flag.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini menggabungkan dua entri inventaris bernama sama, "Assets Data Manager" (lampiran JSM, butir 11-16) dan "Assets Data Manager" (domain Assets, butir 1-10). Data Manager adalah subsistem besar (staging store terpisah); adapter ke produk vendor tertentu di luar daftar minimal masuk ekosistem.

### JSM-AST-040 Object Schema Export & Import (Copy Schema)

- **Deskripsi:** Ekspor object schema ke file untuk diimpor ke site/workspace lain, plus dukungan migrasi.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin.
- **Aturan perilaku (acceptance criteria):**
  1. Export object schema menghasilkan file (ZIP berisi JSON) berisi object type, attribute, status khusus schema, reference type yang dipakai, dan role (opsional).
  2. ⚠ Opsi "Sertakan object" menambahkan data object dan reference ke file ekspor.
  3. ⚠ Ekspor menyertakan icon kustom yang dipakai schema.
  4. Import schema dari file membuat schema baru di workspace tujuan; schema existing tidak pernah ditimpa.
  5. Konflik key/nama schema wajib diselesaikan saat import dengan memasukkan key/nama baru sebelum lanjut.
  6. ⚠ Ekspor/impor dibatasi (default 500.000 object per file); di atas batas pengguna diarahkan ke import CSV/JSON.
  7. ⚠ Migrasi dari Assets/Insight server/DC dan antar-site memakai format file ini atau importer di modul 25; tool migrasi vendor asal tidak didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format file sama dengan definisi template JSM-AST-006 agar template dapat dibuat dari ekspor; migrasi di [25](25-integrasi-impor-ekspor-ekstensi.md).

## Model data standar & Hardware Asset Management

### JSM-AST-041 Assets Common Data Model (System Schemas)

- **Deskripsi:** Model data standar Assets berupa schema sistem bawaan yang dikelola produk ini dan dipakai bersama fitur lain (HAM, Customer Service Management, Services), berbeda dari template karena strukturnya dikelola sistem.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Assets admin dan Admin Situs; akses object mengikuti role schema.
- **Aturan perilaku (acceptance criteria):**
  1. Schema sistem External organizations menyimpan organisasi eksternal (customer, vendor) untuk dipakai fitur CSM.
  2. Schema sistem Geography menyimpan lokasi (negara, kota, site, gedung) dan dipakai antara lain oleh stockroom HAM.
  3. Schema sistem Hardware assets menyimpan model, hardware, stockroom, dan stock rule yang dipakai HAM (JSM-AST-042).
  4. Schema sistem People menyimpan orang (karyawan, kontak) yang dapat dihubungkan ke akun pengguna.
  5. Schema sistem dibuat dan dikelola oleh produk ini (versi definisi disimpan), ditandai badge "Sistem" dan dipisahkan dari schema kustom di daftar schema.
  6. ⚠ Object type dan attribute inti schema sistem tidak dapat dihapus, di-rename, atau diubah tipenya; admin dapat menambah attribute dan object type kustom di dalamnya.
  7. Schema kustom dapat mereferensikan object di schema sistem lewat attribute Object.
  8. ⚠ Workspace yang sudah memiliki data lama harus menjalankan wizard migrasi ke Common Data Model (pemetaan object type lama → object type sistem) sebelum HAM console dapat dibuka.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi schema sistem dibundel di binary dan di-upgrade lewat migrasi data versi-an; CSM di modul [35](35-jsm-ai-otomasi-api-admin-csm.md).

### JSM-AST-042 Hardware Asset Management (HAM)

- **Deskripsi:** Modul pengelolaan siklus hidup hardware di atas Assets: katalog model, stockroom, stock rule, register hardware, hardware work, console, dan dashboard; semua data disimpan sebagai object di schema sistem Hardware assets.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Pengelola aset TI (Manajer Layanan/ITSM), Assets admin, Admin Situs, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Wizard "Set up hardware asset management" hanya dapat dijalankan oleh pengguna yang sekaligus Admin Situs dan Assets admin; wizard membuat schema Hardware assets dan konfigurasi awal.
  2. Data HAM tersimpan di schema Hardware assets (Common Data Model) dan aksesnya mengikuti role schema Assets.
  3. Model catalogue memungkinkan menambah model satu per satu (nama model, vendor, spesifikasi) atau mengimpor dari CSV.
  4. Setiap model memiliki status Approved, Under review, Exception, Deprecated, atau Prohibited; model Prohibited tidak dapat dipilih saat mendaftarkan hardware baru.
  5. Register hardware mencatat hardware individual (serial, aset tag, model, status, pemegang, lokasi) yang mereferensikan model.
  6. Hardware overview dan inventory management menampilkan ringkasan jumlah hardware per status, model, dan lokasi.
  7. Stockroom memiliki nama, tipe, Location (Geography), parent stockroom, dan manager; register dapat difilter berdasarkan lokasi stockroom.
  8. Stok dihitung per stockroom dan per lokasi (agregasi stockroom anak ke parent).
  9. Stock rule menetapkan kuantitas min/max per model per stockroom; saat stok turun di bawah minimum, rule ditandai "Perlu replenishment" dan memicu event yang dapat dipakai automation.
  10. ⚠ Hardware work adalah work item terkait hardware dengan jenis Assign, Return, Retire, dan Replace; penyelesaian work memperbarui status dan pemegang hardware.
  11. HAM console menjadi halaman kerja terpadu (tab Overview, Models, Hardware, Stockrooms, Stock rules, Work).
  12. Dashboard HAM bawaan tersedia, dan pengguna dapat membuat custom hardware dashboard dari widget laporan Assets.
  13. Glossary HAM (definisi model, stockroom, stock rule, status) tersedia sebagai halaman bantuan di console.
  14. ⚠ Console tidak dapat dibuka sampai prasyarat Common Data Model dan migrasi data lama (JSM-AST-041 butir 8) selesai; console menampilkan langkah yang belum dipenuhi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada JSM-AST-041, laporan JSM-AST-044, dan automation JSM-AST-043.

## Otomasi, laporan, API & AI

### JSM-AST-043 Assets in Jira Automation (Triggers, Conditions & Actions)

- **Alias inventaris:** Assets automation & reporting
- **Deskripsi:** Integrasi Assets dengan engine automation produk ini: trigger berbasis object, action lookup/create/edit object, kondisi, dan smart value Assets.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs, Admin Proyek (service project admin), Object schema manager; rule actor wajib punya role schema.
- **Aturan perilaku (acceptance criteria):**
  1. Action "Lookup objects" menjalankan AQL dan menyimpan hasil di `{{lookupObjects}}` untuk branch, kondisi, dan smart value (maks. 100 object per lookup).
  2. Action "Create object" memilih schema dan object type lalu mengisi attribute (boleh dengan smart value); object baru tersedia sebagai `{{createdObject}}`.
  3. Action "Edit object" mengubah attribute pada object dari field Assets work item atau hasil lookup.
  4. ⚠ Trigger berbasis object: Object created, Object updated (opsional per attribute), dan Object deleted, dengan pilihan schema/object type.
  5. ⚠ Scheduled trigger dapat memproses object hasil AQL (mis. lisensi yang kedaluwarsa dalam 30 hari) tanpa memerlukan work item.
  6. ⚠ Rule berlingkup Assets/object schema dapat dibuat dan dikelola dari halaman schema (tab Automation) dan muncul juga di daftar rule global.
  7. Smart value field Assets: `{{issue.<field>.Name}}`, `{{issue.<field>.objectKey}}`, dan attribute via dot notation (mis. `{{issue.Laptop.Owner.Email}}`).
  8. ⚠ Smart value object pada trigger object: `{{object.<attribute>}}`, `{{object.objectKey}}`, `{{object.objectType}}`.
  9. Action "Edit issue" dapat mengisi field Assets via AQL atau smart value (mis. `Key = {{lookupObjects.first.objectKey}}`).
  10. Kondisi berdasarkan attribute object dapat dibuat dengan smart value compare atau JQL `aqlFunction`.
  11. Rule actor wajib memiliki role schema yang sesuai; action yang melanggar gagal dengan entri audit rule "Permission denied pada schema <key>".
  12. Setiap eksekusi rule Assets dihitung terhadap kuota eksekusi automation plan.
  13. Action Lookup objects, Create object, dan Edit object tersedia di daftar action kategori "Assets".
  14. ⚠ Trigger berbasis perubahan object tersedia di daftar trigger kategori "Assets".
  15. Smart value object Assets dapat dipakai di action email, komentar, dan webhook.
  16. Laporan Assets menampilkan jumlah object per type/attribute dengan chart (rinci di JSM-AST-044).
  17. Data Assets dapat dianalisis di alat analitik tingkat Enterprise melalui ekspor/API (rinci di JSM-AST-044 butir 7).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 13-17 berasal dari entri lampiran "Assets automation & reporting". Engine, rule actor (AUTO-006), dan scheduled trigger (AUTO-023) di [20](20-otomasi.md); action "Edit issue" (butir 9) di AUTO-039; kondisi Assets/object generik di AUTO-038 butir 3; smart value work item di AUTO-063 dan sintaks smart value di AUTO-065; action email dan web request (butir 15) di AUTO-052 dan AUTO-054. Katalog automation Service Management di [35](35-jsm-ai-otomasi-api-admin-csm.md) merujuk fitur ini: trigger Assets di JSM-AI-016 butir 6, kondisi AQL di JSM-AI-017 butir 7, aksi Assets di JSM-AI-018 butir 11, smart value `{{lookupObjects}}`/`{{attributeChangelog}}`/`{{containerObject}}` di JSM-AI-020 butir 7; perilaku rinci action Assets dimiliki fitur ini. Automation schema lama di JSM-AI-023 dan JSM-AST-048. Event `assets.object.*` dari JSM-AST-014.

### JSM-AST-044 Assets Reports

- **Deskripsi:** Pelaporan bawaan Assets untuk menganalisis jumlah dan distribusi object serta keterkaitannya dengan work item.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, pengelola aset TI, Assets admin; data mengikuti role schema pembaca.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Report object berdasarkan object type/attribute dengan visualisasi bar, pie, dan tabel.
  2. Report object dari waktu ke waktu (jumlah created/updated per hari/minggu/bulan).
  3. Report work item per object (mis. jumlah incident per service) memakai connected tickets.
  4. Report dapat difilter dengan AQL dan pilihan schema.
  5. ⚠ Report dapat disimpan dengan nama dan dibagikan ke pengguna/group yang juga punya akses schema.
  6. Data report dapat diekspor ke CSV.
  7. Data Assets tersedia untuk analitik lanjutan lewat ekspor terjadwal atau API baca massal (padanan sumber data analitik Enterprise).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen chart dipakai bersama [16](16-laporan-insight.md) dan [17](17-dashboard-gadget.md); API baca massal butir 7 memakai JSM-AST-045; padanan analitik/data lake Service Management di JSM-AI-066 ([35](35-jsm-ai-otomasi-api-admin-csm.md)); layanan analitik/data lake vendor asal tidak dibangun.

### JSM-AST-045 Assets REST API

- **Deskripsi:** REST API publik Assets untuk CRUD schema, object type, attribute, object, pencarian AQL, import, status, dan icon.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Integrator/Developer API; setiap panggilan memakai role Assets pemilik kredensial.
- **Aturan perilaku (acceptance criteria):**
  1. Base path native tunggal adalah `/api/v1/assets/workspace/{workspaceId}/v1` (semua path endpoint di fitur ini relatif terhadap base tersebut); satu-satunya path kompatibel adalah `/jsm/assets/workspace/{workspaceId}/v1` yang melayani endpoint dan payload identik dengan base native. Path lain (mis. tanpa segmen `/v1`) mengembalikan HTTP 404.
  2. ⚠ Autentikasi memakai Basic (email + API token) atau OAuth 2.0 (3LO) dengan scope `read:assets`, `write:assets`, dan `delete:assets`.
  3. Object schema: list, get, create, update, delete, serta list object type dan attribute per schema.
  4. Object type: create, update, delete, ubah posisi (pindah parent/urutan), dan list attribute.
  5. Object type attribute: create, update, delete.
  6. ⚠ Object: get, create, update, delete, history, referenceinfo, dan attachments.
  7. Pencarian AQL `POST /object/aql` dengan pagination (`startAt`, `maxResults`), `includeAttributes`, dan total hasil.
  8. Endpoint IQL lama (`/iql`) dilayani sebagai alias deprecated dengan header `Deprecation`.
  9. Config: statustype, reference type, dan icon (list/get/create/update/delete sesuai izin).
  10. Endpoint import source dan Imports REST API external (JSM-AST-036) tersedia di bawah base yang sama.
  11. Connected tickets `GET /objectconnectedtickets/{id}/tickets`.
  12. Endpoint progress `GET /progress/{taskId}` untuk operasi asinkron (bulk, delete schema, import).
  13. ⚠ Rate limiting dan kode galat standar (400, 401, 403, 404, 409, 413, 429) dengan header `Retry-After` mengikuti kebijakan API site.
  14. Penemuan workspace: `GET /api/v1/servicedeskapi/assets/workspace` (di bawah base REST API Service Management modul 35) mengembalikan `workspaceId` situs yang dipakai pada base path butir 1; pemanggil tanpa entitlement Assets mendapat daftar kosong.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JSM-AST-045 adalah **pemilik** spesifikasi Assets REST API; JSM-AI-031 di [35](35-jsm-ai-otomasi-api-admin-csm.md) hanya mendaftarkan API Assets dalam keluarga API Service Management dan wajib mengikuti base path native dan kompatibel pada butir 1 (keduanya sudah selaras: `/api/v1/assets/workspace/{workspaceId}/v1` dan `/jsm/assets/workspace/{workspaceId}/v1`). Selisih yang masih harus diselaraskan di modul 35: JSM-AI-031 butir 9 menyatakan endpoint IQL tidak disediakan, sedangkan modul ini menetapkan `/iql` sebagai alias deprecated (butir 8, JSM-AST-048 butir 4); keputusan pemilik di sini yang berlaku. Konvensi route, autentikasi (API token, OAuth 2.0), rate limit, dan format error di [23](23-api-publik-webhook.md); webhook/extension point Assets import untuk app di JSM-AI-033; dokumentasi Swagger wajib diregenerasi (`make swag`) dan endpoint memakai `@Security BearerAuth`.

### JSM-AST-046 Assets in Confluence (Assets Macro / Live Doc)

- **Deskripsi:** Menampilkan data object Assets (hasil AQL) di halaman dokumen; di produk ini diwujudkan sebagai blok/embed tabel Assets yang dapat disisipkan di halaman dokumen internal dan di wiki eksternal yang mendukung embed.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Penulis halaman dokumen (memasang embed); Agent Layanan dan pembaca halaman (melihat data sesuai role Assets).
- **Aturan perilaku (acceptance criteria):**
  1. Penulis dapat menyisipkan tabel/daftar Assets berbasis AQL di halaman atau dokumen hidup.
  2. Penulis memilih object schema dan kolom/attribute yang ditampilkan.
  3. Data yang tampil mengikuti role Assets milik pembaca; pembaca tanpa akses melihat pesan "Tidak ada akses ke data Assets" tanpa data object.
  4. ⚠ Data di-refresh setiap kali halaman dibuka (tanpa cache bersama antar-pembaca).
  5. ⚠ Klik object membuka object detail di Assets.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Integrasi langsung dengan produk wiki vendor asal di luar cakupan; embed memakai endpoint baca AQL (JSM-AST-045) dan komponen tabel SPA. Penempatan di Pages project ([04](04-proyek.md)) bergantung pada keputusan modul dokumen.

### JSM-AST-047 Atlassian Intelligence / Rovo for Assets

- **Deskripsi:** Fitur AI opsional produk ini untuk membantu menulis AQL dan menelusuri data Assets (padanan fitur AI vendor asal); tidak diimplementasikan sebelum dikonfirmasi pemilik produk.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Assets admin (mengaktifkan).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Natural language ke AQL: pengguna mengetik kalimat (mis. "laptop yang garansinya habis bulan ini") dan memperoleh AQL yang dapat ditinjau dan diedit sebelum dijalankan; AQL yang dihasilkan selalu divalidasi parser.
  2. ⚠ Pencarian object Assets lewat asisten/pencarian AI produk ini, dengan hasil tetap dibatasi role pengguna.
  3. ⚠ Ringkasan CI terkait work item untuk agent (attribute penting, relasi, incident terbaru) di issue view.
  4. Fitur AI Assets dapat diaktifkan/nonaktifkan oleh admin (JSM-AST-003 butir 10) dan default nonaktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Toggle AI tingkat situs, penyedia LLM, dan kuota pemakaian diatur di ADM-052 ([22](22-administrasi-sistem.md)); fitur Assets AI hanya aktif bila AI situs (ADM-052) dan toggle Assets (JSM-AST-003 butir 10) sama-sama aktif. Pencarian bahasa alami agent di JSM-AI-039 dan lapisan AI opsional Service Management di modul [35](35-jsm-ai-otomasi-api-admin-csm.md); data Assets tidak dikirim ke penyedia AI eksternal tanpa persetujuan admin.

## Fitur warisan (rujukan migrasi)

### JSM-AST-048 Legacy Insight Features (IQL, Insight Fields, Schema Automation)

- **Deskripsi:** Fitur warisan Insight (field Insight, IQL, automation schema lama) yang di Cloud digantikan field Assets objects, AQL, dan automation; hanya relevan untuk migrasi dan kompatibilitas, tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs dan Object schema manager (saat migrasi).
- **Aturan perilaku (acceptance criteria):**
  1. Field Insight Object / Insight Object (Multiple) dengan konfigurasi IQL tidak dibangun; saat impor dari server/DC field tersebut dipetakan ke field Assets objects (single/multi) dengan filter AQL setara.
  2. Field Insight Referenced Object dan Insight Readonly Object tidak dibangun; importer memetakannya ke field Assets objects read-only atau melaporkannya sebagai item yang tidak dimigrasi.
  3. Dokumentasi migrasi menjelaskan deprecation field Insight dan langkah beralih ke field Assets objects.
  4. IQL hanya didukung sebagai alias AQL (endpoint `/iql`, `iqlFunction()`), sesuai JSM-AST-024 butir 12 dan JSM-AST-028 butir 2.
  5. Automation schema lama (trigger object added/updated/deleted/schedule, condition AQL, action email/set attribute/create issue/webhook/Groovy) tidak dibangun; skrip Groovy tidak didukung sama sekali.
  6. Importer menghasilkan laporan rule automation schema lama yang perlu dibuat ulang manual di automation (JSM-AST-043).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dicantumkan agar inventaris tercakup; lihat Di luar cakupan. Importer di [25](25-integrasi-impor-ekspor-ekstensi.md); rujukan automation Assets lama juga dicatat di JSM-AI-023 ([35](35-jsm-ai-otomasi-api-admin-csm.md)) dengan padanan JSM-AI-016/JSM-AI-017/JSM-AI-018 dan JSM-AST-043.

## Pertanyaan terbuka

1. Apakah Assets hanya dibuka untuk tier Premium/Enterprise seperti produk acuan, atau versi dasar (mis. satu schema, batas object kecil) disediakan juga di tier Standard sebagai pembeda kompetitif?
2. Apakah "Assets di produk inti" (field Assets di project non-service tanpa Service Management) diaktifkan sejak rilis pertama Assets, atau menunggu Assets service stabil?
3. Berapa batas default resmi (jumlah object per workspace, attribute per object type, hasil AQL, ukuran file import, object per field multi, bulk) untuk deployment SaaS maupun self-host?
4. Apakah role schema boleh ditetapkan ke project role selain pengguna dan group, dan apakah Agent Layanan otomatis mendapat role User pada schema baru?
5. Apakah template schema menyertakan opsi data contoh, dan template mana yang menjadi katalog resmi rilis pertama?
6. Apakah Watch object, Move object, inline edit di list, filter AQL tersimpan, dan comment object (fitur yang jelas ada di DC) dibangun di Fase 7a, 7b, atau ditunda?
7. Apakah reference type perlu lingkup per schema, atau cukup global per workspace?
8. Apakah scheduled import untuk CSV/JSON boleh mengambil file dari URL (HTTP/S3) agar sinkronisasi terjadwal benar-benar otomatis?
9. Apakah Assets Discovery dan Data Manager dibangun sendiri (binary on-prem dan staging store), dikerjakan sebagai proyek terpisah, atau diserahkan ke ekosistem melalui Imports REST API?
10. Apakah integrasi Microsoft Entra ID menjadi konektor bawaan pertama, atau konektor lain (mis. Google Workspace) diprioritaskan lebih dulu untuk pasar sasaran?
11. Base path sudah ditetapkan di JSM-AST-045 butir 1 (native `/api/v1/assets/workspace/{workspaceId}/v1`, kompatibel `/jsm/assets/workspace/{workspaceId}/v1`); yang masih terbuka: seberapa dalam kompatibilitas payload (nama field, bentuk error, paginasi) dengan Assets Cloud yang wajib dijaga, dan apakah alias `/iql` dipertahankan (JSM-AST-045 butir 8) atau dihapus seperti JSM-AI-031 butir 9?
12. Apakah fitur AI untuk Assets (NL→AQL, ringkasan CI) dimasukkan ke peta jalan, dan penyedia model mana yang dipakai bila ya?
13. Bagaimana embed tabel Assets di halaman dokumen diwujudkan bila produk ini belum memiliki modul dokumen/wiki sendiri?
14. Apakah schema bawaan untuk Services adalah satu schema "Services" (JSM-AST-030) atau dua schema Business Portfolio dan Software Registry (JSM-ITSM-004 butir 4)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga dan konektor ekosistem** untuk Assets (mis. konektor AWS/Azure/Google Cloud, Jamf, Intune, Snipe-IT, Lansweeper, ServiceNow, Snow, Tempo, pembuat label/QR lanjutan) tidak dibangun sebagai fitur native; integrasi tersebut memakai External Imports REST API (JSM-AST-036) atau Assets REST API (JSM-AST-045). Pengecualian: konektor Microsoft Entra ID (JSM-AST-037) dibangun sebagai konektor bawaan.
- **Fitur Data Center-only / deprecated**: field Insight (Insight Object, Referenced Object, Readonly Object), automation schema Insight dengan Groovy, dan tipe attribute Time tidak dibangun; hanya dirujuk lewat JSM-AST-048 dan JSM-AST-009 butir 12 untuk keperluan migrasi.
- **Tool migrasi vendor asal** (Jira Cloud Migration Assistant, migrasi Cloud-to-Cloud) tidak didukung; migrasi memakai format ekspor schema (JSM-AST-040) dan importer modul 25.
- **Layanan lintas produk vendor asal**: produk wiki (Confluence) sebagai layanan, Bitbucket, Opsgenie, Atlassian Analytics/Data Lake, Atlassian Intelligence/Rovo, dan paket bundel Service Collection tidak dibangun; produk ini menyediakan padanan internal (repository modul 24, tim Operations modul 34, ekspor/API analitik, lapisan AI opsional modul 35) atau titik integrasi saja.
- **Billing dan pengelolaan lisensi** Premium/Enterprise vendor asal tidak dibangun; entitlement Assets diperlakukan sebagai feature flag konfigurasi plan.
- **Aplikasi mobile native** untuk pemindaian QR dan melihat object tidak dibangun; kebutuhan dipenuhi tampilan responsif SPA dan kamera perangkat.
- **Pemindaian jaringan dari server produk ini** tidak dilakukan; Assets Discovery selalu berjalan di infrastruktur pelanggan.
