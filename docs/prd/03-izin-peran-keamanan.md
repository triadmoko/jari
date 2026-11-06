# 03 — Izin, Peran Proyek & Keamanan Work Item

Modul ini mendefinisikan seluruh model otorisasi produk ini: product access (lisensi/peran aplikasi), global permission tingkat site, permission scheme dan project role untuk company-managed project, access level dan role untuk team-managed project, issue security (keamanan per work item), visibilitas komentar/worklog, akses khusus (publik, guest, app, automation actor, AI agent), izin objek (board, filter, plan), diagnostik dan audit izin, REST API terkait izin, serta kontrol keamanan data tingkat enterprise (klasifikasi, data security policy, redaksi). Inti modul adalah satu mesin evaluasi izin terpusat di backend yang dipakai oleh SEMUA permukaan (UI, JQL, REST, notifikasi, board, laporan, ekspor, automation, webhook) sehingga tidak ada jalur yang membocorkan data. Modul ini tidak mendefinisikan siapa user/grup itu (lihat modul 02) maupun cara login (modul 01); modul ini menjawab pertanyaan "boleh atau tidak".

Modul terkait:

- [00 Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md) — plan tier & feature flag.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — sesi, token, status akun aktif.
- [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — user, group, product access di tingkat organisasi.
- [04 Proyek](04-proyek.md) — tipe project (company-/team-managed), pengaturan, arsip project.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field, screen/layout, scheme lain.
- [06 Work Item Inti](06-work-item-inti.md) — operasi create/edit/move/delete/archive/bulk.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar, attachment, history.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — conditions, workflow properties.
- [10 Board](10-board-kolom-kartu.md) dan [11 Backlog & Sprint](11-backlog-sprint.md) — board admin, sprint.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — permission filter pada hasil pencarian, share filter.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md), [15 Rilis & Komponen](15-rilis-versi-komponen.md), [17 Dashboard](17-dashboard-gadget.md), [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md).
- [19 Notifikasi](19-notifikasi-email.md), [20 Otomasi](20-otomasi.md), [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log, General configuration.
- [23 REST API Publik & Webhook](23-api-publik-webhook.md), [24 Integrasi DevOps](24-integrasi-devops.md), [25 Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md).
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md), [28 Model Data Inti](28-model-data-inti.md) — tabel scheme/grant/role.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md), [30 JSM Portal & Customers](30-jsm-help-center-portal-pelanggan-kanal.md), [37 JPD Peran](37-jpd-insights-delivery-peran-integrasi.md).

## Persona terkait

- **Admin Situs** — mengelola product access, group, guest, kebijakan keamanan data organisasi.
- **Admin Proyek/Project Lead** — mengelola People/role project, access level team-managed, issue restriction, AI agent di project.
- **Scrum Master/Team Lead** — memerlukan Manage Sprints, board admin.
- **Developer/Anggota Tim** — pemakai izin harian (create, edit, transition, comment, log work).
- **Product Owner/PM** — Manage Versions, share filter/dashboard, Plans.
- **Stakeholder/Viewer** — akses baca (Browse, View Voters and Watchers, View Aggregated Data).
- **Integrator/Developer API** — memeriksa izin lewat REST, mengelola scheme secara terprogram, app permission.
- **Tamu/Kolaborator Eksternal** — guest yang dibatasi ke satu project.
- **Agent Layanan** dan **Customer/Pelapor** — izin service project (Service Project Agent, akses portal).

## User story modul

- Sebagai Admin Situs, saya ingin mengatur product access dan group default agar hanya orang berlisensi yang bisa memakai produk ini (PERM-004, PERM-006).
- Sebagai Admin Situs, saya ingin memberi global permission Administer Jira hanya ke group admin dan terlindung dari lockout agar konfigurasi site aman (PERM-006, PERM-007).
- Sebagai Admin Situs, saya ingin membuat permission scheme yang dipakai bersama banyak project agar kebijakan akses konsisten dan mudah diubah sekali untuk semua (PERM-014, PERM-015, PERM-016).
- Sebagai Admin Proyek/Project Lead, saya ingin menambah orang/group ke project role di halaman People agar akses tim diatur tanpa meminta Admin Situs (PERM-019).
- Sebagai Developer/Anggota Tim, saya ingin tombol dan aksi yang tidak boleh saya lakukan disembunyikan atau dinonaktifkan dengan penjelasan agar saya tidak bingung saat ditolak (PERM-056, PERM-002).
- Sebagai Scrum Master/Team Lead, saya ingin memiliki Manage Sprints tanpa menjadi Admin Proyek agar dapat memulai dan menutup sprint sendiri (PERM-025).
- Sebagai Admin Proyek/Project Lead, saya ingin mengatur team-managed project menjadi Open, Limited, atau Private agar tingkat keterbukaan sesuai kebutuhan tim (PERM-038, PERM-039).
- Sebagai Admin Situs, saya ingin mendefinisikan issue security level agar work item sensitif (mis. kerentanan keamanan) hanya terlihat oleh kelompok tertentu di semua permukaan termasuk JQL dan notifikasi (PERM-042, PERM-043, PERM-044).
- Sebagai Developer/Anggota Tim, saya ingin membatasi komentar atau worklog ke role tertentu agar diskusi internal tidak terlihat semua orang (PERM-045).
- Sebagai Admin Proyek/Project Lead, saya ingin Permission Helper menjelaskan mengapa seseorang tidak bisa melihat atau mengedit work item agar saya cepat memperbaiki konfigurasi (PERM-055).
- Sebagai Tamu/Kolaborator Eksternal, saya ingin bekerja di satu project yang mengundang saya tanpa melihat project lain agar kolaborasi lintas organisasi tetap aman (PERM-047).
- Sebagai Integrator/Developer API, saya ingin memeriksa izin user secara bulk lewat REST agar app saya menampilkan aksi yang tepat tanpa N+1 request (PERM-058, PERM-059).
- Sebagai Stakeholder/Viewer, saya ingin dapat menelusuri project publik tanpa login bila admin mengizinkan agar transparansi ke komunitas terjaga (PERM-046).
- Sebagai Admin Situs, saya ingin setiap perubahan scheme, grant, dan role tercatat di audit log agar perubahan akses dapat ditelusuri (PERM-057).
- Sebagai Admin Situs, saya ingin memblokir ekspor dan akses anonim pada project berklasifikasi tinggi agar data sensitif tidak bocor (PERM-064, PERM-065).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 26 |
| P1 | 16 |
| P2 | 17 |
| P3 | 8 |
| **Total** | **67** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 22 |
| Fase 2 | 11 |
| Fase 3 | 23 |
| Fase 4 | 9 |
| Fase 5 | 1 |
| Fase — (rujukan DC) | 1 |

Strategi urutan implementasi: Fase 1 membangun mesin evaluasi izin terpusat (Authorizer) beserta data model scheme/grant/role yang lengkap sejak awal, tetapi hanya mengisinya dengan default permission scheme, role bawaan, dan access level team-managed sehingga UI cukup berupa halaman People dan pengaturan akses; izin yang sudah dibutuhkan fitur Fase 1 modul lain (Link/Schedule/Move/Modify Reporter, Voters & Watchers, Browse users and groups, Create team-managed projects) ikut ditegakkan dengan grant default. Kunci permission (permission key) untuk seluruh katalog didaftarkan sejak Fase 1 agar kode use case di modul lain langsung memanggil `authorizer.Can(...)` dengan key yang benar walau grant-nya masih default. Fase 2 menambah izin yang mengikuti fitur konfigurabilitas (versi, time tracking, share filter, bulk change, visibilitas komentar/worklog) serta endpoint mypermissions untuk SPA. Fase 3 membuka editor permission scheme penuh, issue security, Permission Helper, audit, akses publik, automation actor, izin arsip work item/project, dan REST API administrasi izin. Fase 4 ke atas menangani guest, app/AI agent, fitur Premium (custom role, Plans), dan kontrol keamanan data enterprise; izin JSM menyusul di Fase 5 bersama service project.

## Fitur

## Model & mesin otorisasi

### PERM-001 Permission Evaluation Model

- **Deskripsi:** Mesin evaluasi izin inti yang menentukan apakah user atau agent boleh melakukan suatu aksi pada site, project, atau work item berdasarkan product access, global permission, permission scheme, project role, issue security, dan workflow permission properties. Model ini adalah kontrak perilaku yang diimplementasikan oleh PERM-002.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Sistem (evaluator), Admin Situs (permission Administer Jira) dan Admin Proyek (pengonfigurasi), user dan AI agent (subjek evaluasi).
- **Aturan perilaku (acceptance criteria):**
  1. Model bersifat additive: izin efektif adalah union semua grant; permission scheme tidak mendukung explicit deny. User mendapat izin bila cocok dengan SALAH SATU grantee pada permission tersebut.
  2. Urutan evaluasi tetap: (a) akun aktif, (b) product access (application role/lisensi atau guest), (c) global permission atau project permission dari scheme/access level, (d) issue security level bila work item memilikinya, (e) workflow permission properties pada status saat ini (company-managed). Kegagalan di tahap mana pun menghentikan evaluasi dengan hasil "tolak".
  3. Browse Projects adalah prasyarat hampir semua project/issue permission; evaluasi izin issue lain (Edit, Comment, dst.) menolak bila user tidak memiliki Browse Projects pada project tersebut.
  4. Pengecualian: Create Issues dievaluasi tanpa mensyaratkan Browse Projects sehingga skenario create-only (mis. publik membuat work item tanpa bisa melihatnya) didukung; setelah create, user tanpa Browse diarahkan ke halaman konfirmasi tanpa menampilkan isi work item.
  5. Issue security bersifat restriktif: user dengan Browse Projects tetap tidak bisa melihat work item bila bukan anggota security level work item tersebut.
  6. Grantee dinamis (Reporter, Current assignee, Project lead, User custom field value, Group custom field value) dievaluasi per work item saat runtime terhadap nilai field terkini.
  7. Izin konteks project (Create Issues, Manage Sprints, Administer Projects) dievaluasi tanpa work item; izin konteks issue dievaluasi terhadap work item tertentu. Pemanggil wajib menyertakan konteks yang sesuai; konteks yang salah menghasilkan galat pemrograman, bukan "izinkan".
  8. Admin Situs mewarisi Administer Projects di semua project, tetapi TIDAK otomatis mendapat Browse Projects atas work item; tanpa grant eksplisit Admin Situs dapat membuka project settings namun tidak melihat work item.
  9. Semua permukaan (UI, JQL, REST API, notifikasi, board, backlog, laporan, gadget, export, automation, @mention, webhook payload) wajib memakai permission filter yang sama; uji kontrak memverifikasi setiap permukaan terhadap matriks yang sama.
  10. Hasil JQL, count, dan agregasi hanya menghitung work item yang boleh dilihat user, kecuali data agregat yang sengaja dibuka lewat View Aggregated Data (PERM-028).
  11. Hasil evaluasi di-cache per user/project; cache diinvalidasi saat scheme, grant, role membership, group membership, security level, atau product access berubah, dan perubahan berlaku paling lambat dalam hitungan detik (target awal ≤ 5 detik; batas final mengikuti keputusan di Pertanyaan terbuka). Uji: setelah grant dicabut, request berikutnya setelah batas waktu tersebut ditolak.
  12. User tanpa product access tidak dapat memakai produk walau tercantum di grant; pengecualian: customer JSM yang mengakses lewat portal; guest dibatasi ke project yang mengundangnya.
  13. AI agent dengan akun sendiri dievaluasi seperti user biasa (butuh role/grant eksplisit); agent yang bertindak atas nama user mewarisi izin user tersebut dan tidak pernah melebihinya.
  14. Edge case: user yang dinonaktifkan tetap tercatat di grant/role (untuk riwayat) tetapi tidak pernah dievaluasi sebagai aktif.
  15. Edge case: work item di project archived tidak dapat diakses user biasa sampai project di-restore (lihat PERM-053).
  16. Edge case: work item archived (Premium) tersembunyi dari pencarian biasa dan bersifat read-only bagi semua user termasuk admin (lihat PERM-035).
- **Team-managed vs company-managed:** Company-managed memakai permission scheme + project role + issue security scheme + workflow properties. Team-managed memakai access level (Open/Limited/Private), role bawaan/kustom, dan issue restriction; tahap (e) workflow properties tidak berlaku.
- **Catatan & ketergantungan:** Implementasi teknis di PERM-002; data model tabel grant `(scheme_id, permission_key, holder_type, holder_parameter)` di [28 Model Data Inti](28-model-data-inti.md). Produk ini tidak menyediakan field-level security native; menyembunyikan field dilakukan lewat screen/field configuration/layout di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md). UI memakai istilah "space"/"work item" sesuai glosarium modul 00, sementara permission key internal tetap memakai nama klasik (BROWSE_PROJECTS, CREATE_ISSUES, dst.) demi kompatibilitas API.

### PERM-002 Server-side Authorization (RBAC Enforcement)

- **Deskripsi:** Mesin otorisasi terpusat di backend yang mengevaluasi product access, global permission, project permission, dan issue security untuk setiap akses. Ini adalah realisasi teknis model PERM-001 di arsitektur Clean Architecture repo.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering (pemelihara), dipakai oleh seluruh use case modul lain.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia satu interface `Authorizer` di `internal/shared/domain/` yang dipanggil dari use case (bukan dari handler) sehingga perilaku konsisten untuk REST, SSE/realtime, export, email, dan payload webhook; review kode menolak pengecekan izin ad hoc di handler.
  2. Evaluasi berlapis berurutan: product access → global permission → project permission (scheme/role atau access level team-managed) → issue security level → visibilitas comment/worklog; setiap lapisan dapat diuji terpisah.
  3. Tersedia permission check batch untuk daftar work item (mis. `FilterVisible(user, issueIDs)` atau predikat SQL yang disuntikkan ke kueri) sehingga list/board/JQL tidak melakukan N+1 query; uji performa memastikan jumlah query konstan terhadap jumlah work item.
  4. Endpoint mypermissions dan permissions/check (PERM-058) disediakan agar SPA dapat menyembunyikan tombol, tetapi server tetap menolak aksi tanpa izin dengan HTTP 403 (atau 404 untuk objek yang tidak boleh diketahui keberadaannya).
  5. Anonymous access untuk project publik didukung hanya bila admin memberi grant ke Public (PERM-046); request tanpa token dievaluasi sebagai subjek "anonymous", bukan ditolak di middleware.
  6. Penolakan akses berulang dari satu subjek dalam jendela waktu tertentu (indikasi enumerasi key work item) dicatat sebagai event keamanan dan dapat memicu rate limit.
  7. Tersedia test matrix otomatis per permission key × holder type × tipe project yang dijalankan di CI; penambahan permission key baru gagal di CI bila belum masuk matriks.
- **Team-managed vs company-managed:** Authorizer menyembunyikan perbedaan di balik satu API; resolver grant dipilih berdasarkan tipe project.
- **Catatan & ketergantungan:** Kondisi repo saat ini: `middleware.AuthMiddleware()` hanya memvalidasi JWT dan menyimpan klaim di `"sess"`; entity `User` memiliki field `Role` (varchar, default `user`) yang belum ada di migrasi `001_create_users_table`. Migrasi: field `Role` global digantikan oleh product access + global permission (PERM-004, PERM-006); nilai `admin` dipetakan ke group admin default saat migrasi. Middleware tetap hanya mengautentikasi; otorisasi pindah ke use case. Rincian NFR (latensi, cache) di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

### PERM-003 Administrative Role Matrix

- **Deskripsi:** Hierarki peran administrasi produk ini beserta kewenangan masing-masing, dari Admin Organisasi sampai admin objek (board/filter/plan). Menjadi acuan tunggal untuk menentukan siapa boleh mengakses halaman administrasi mana.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Admin Organisasi, User access admin (eks site admin), Admin Situs (permission Administer Jira; product admin role), Admin Proyek, pemilik objek.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Organisasi: mengelola organisasi, domain, kebijakan (policy), kontrol keamanan data, billing, semua user dan app; merupakan peran tertinggi.
  2. Site admin model lama (dipertahankan sebagai alias User access admin untuk seluruh produk di site; berbeda dari persona Admin Situs): mengelola user dan product access di site; tidak otomatis memegang Administer Jira.
  3. User access admin: mengelola akses user untuk produk tertentu (memberi/mencabut product access) tanpa dapat mengubah konfigurasi produk.
  4. ⚠ Trusted user (model lama): dapat mengundang user dan memberi product access; di model manajemen user baru peran ini dipetakan ke User access admin dan tidak ditawarkan untuk site baru.
  5. Admin Situs (product admin role): setara pemegang global permission Administer Jira (PERM-007).
  6. Admin Proyek: pemegang Administer Projects (company-managed) atau role Administrator (team-managed) pada project tertentu saja.
  7. Admin objek: board admin, pemilik/editor filter dan dashboard, pemilik plan; kewenangannya terbatas pada objek itu.
  8. PRD dan dokumentasi admin memuat tabel kemampuan per level (lihat tabel di bawah) yang menjadi acuan uji akses halaman administrasi; setiap halaman admin mendeklarasikan level minimum.
- **Team-managed vs company-managed:** Admin Proyek team-managed = role Administrator; company-managed = Administer Projects via scheme.
- **Catatan & ketergantungan:** Tabel ringkas kewenangan:

  | Level | Kelola user/akses | Global permission & scheme | Pengaturan project | Objek (board/filter/plan) |
  |---|---|---|---|---|
  | Admin Organisasi | Ya | Hanya bila juga Admin Situs | Hanya bila juga Admin Situs | — |
  | User access admin (eks site admin) | Ya (akses produk) | Tidak | Tidak | — |
  | Admin Situs | Tidak | Ya | Ya (semua project) | Ya |
  | Admin Proyek | Undang via People (bila diizinkan) | Tidak | Ya (project-nya) | Board project-nya |
  | Admin objek | Tidak | Tidak | Tidak | Objek miliknya |

  Detail user/group di [02 Manajemen Pengguna](02-pengguna-grup-organisasi.md).

### PERM-004 Product Access & Application Roles

- **Deskripsi:** Mengatur siapa yang memiliki akses (lisensi) ke produk beserta product role-nya; product access adalah gerbang pertama evaluasi izin.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Admin Organisasi, User access admin.
- **Aturan perilaku (acceptance criteria):**
  1. Product role yang didukung: User, Product admin, User access admin; Guest untuk kolaborator eksternal (PERM-047). Satu user dapat memiliki role berbeda per produk (Jira, Service Management, Product Discovery).
  2. Setiap produk memiliki default access group; user baru yang diberi akses produk otomatis dimasukkan ke group tersebut.
  3. Group default `<product>-users-<site>` dan `<product>-admins-<site>` (selaras seed DM-001 di [28](28-model-data-inti.md); alias impor: `jira-users-<site>`, `jira-admins-<site>`) dibuat saat site dibuat dan tidak dapat dihapus selama menjadi default access group.
  4. User tanpa akses dapat mengajukan access request; admin menyetujui atau menolak, dan pemohon diberi tahu hasilnya.
  5. Approved domain / self sign-up: user dengan email domain terverifikasi dapat bergabung sendiri dan langsung atau setelah persetujuan mendapat akses sesuai konfigurasi.
  6. Batas jumlah user per plan ditegakkan (Free terbatas); menambah akses melebihi batas ditolak dengan pesan yang menyebut batas plan.
  7. Pencabutan product access segera menghapus kemampuan memakai produk (sesi aktif kehilangan akses pada request berikutnya) tanpa menghapus grant/role yang mereferensikan user.
  8. Grantee "Application access" di permission scheme (PERM-015) merujuk daftar user dengan product access ini dan dievaluasi dinamis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengelolaan UI akses produk, undangan, dan access request ada di [02 Manajemen Pengguna](02-pengguna-grup-organisasi.md); modul ini hanya menetapkan dampaknya pada evaluasi izin. REST baca di PERM-063.

### PERM-005 Plan-tier Permission Limitations

- **Deskripsi:** Batasan kemampuan konfigurasi izin per plan produk (Free, Standard, Premium, Enterprise). Diimplementasikan lewat feature flag plan dari modul 00.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Sistem (penegak), Admin Situs (permission Administer Jira; terdampak).
- **Aturan perilaku (acceptance criteria):**
  1. Free: tidak dapat assign project role, tidak dapat membatasi izin via permission scheme, dan tidak dapat mengonfigurasi issue security; UI menampilkan ajakan upgrade, bukan galat.
  2. Free: semua user ber-product access diperlakukan memiliki Administer Projects di semua project.
  3. Free: role team-managed tidak dapat diubah dan tidak ada issue-level security (issue restriction).
  4. Downgrade ke Free atau migrasi dari self-managed ke Free: scheme, role, dan security scheme yang ada dipertahankan tetapi read-only; data tidak dihapus dan kembali dapat diedit setelah upgrade.
  5. Standard: permission scheme, project role, dan issue security tersedia lengkap.
  6. Premium: menambah archive work item (PERM-035) dan Plans (PERM-052); di plan lebih rendah aksi terkait tidak tampil dan API mengembalikan galat plan.
  7. ⚠ Premium: menambah custom role team-managed (PERM-040); tier minimum dapat bergeser ke Standard sesuai keputusan produk.
  8. Enterprise / add-on keamanan data: data classification, data security policy, dan data redaction.
- **Team-managed vs company-managed:** Batasan berlaku untuk keduanya sesuai butir di atas.
- **Catatan & ketergantungan:** Feature flag plan didefinisikan di [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md). Untuk instalasi self-hosted open-source, pemilik produk dapat memilih semua fitur terbuka (lihat Pertanyaan terbuka).

## Global permission

### PERM-006 Global Permissions

- **Deskripsi:** Halaman administrasi (Settings > System > Global permissions) untuk memberikan izin tingkat site yang tidak terikat project tertentu kepada group.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira), Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar global permission yang didukung: Administer Jira, Browse users and groups, Share dashboards and filters, Manage group filter subscriptions, Manage custom onboarding, Make bulk changes, Create team-managed projects (label UI "Create team-managed spaces"). Setiap permission menampilkan deskripsi singkat.
  2. Ekstensi/app terpasang dapat mendaftarkan global permission miliknya sendiri yang tampil di daftar yang sama dengan penanda sumber app (lihat PERM-048).
  3. Global permission hanya dapat diberikan ke Group, tidak ke individual user; API yang mencoba memberi ke user ditolak dengan galat validasi.
  4. Administer Jira dikelola lewat group admin produk (setara product admin role); Admin Organisasi/User access admin yang ingin mengonfigurasi produk harus menambahkan dirinya ke group tersebut.
  5. Dialog "Grant access": pilih permission lalu pilih satu group; grant duplikat ditolak.
  6. Menghapus group dari sebuah global permission meminta konfirmasi yang menyebut nama permission dan group.
  7. Tampilan per permission: daftar group pemegang izin dengan tautan ke daftar anggota group.
  8. Guest tidak pernah mendapat global permission apa pun walau menjadi anggota group pemegang izin; evaluasi mengabaikan grant global untuk guest.
  9. Group default site (`<product>-users-<site>`, `<product>-admins-<site>`, `site-admins`, `org-admins`; alias impor: `jira-administrators`, `jira-users-<site>`, `jira-software-users-<site>`, `jira-admins-<site>`) memiliki global permission bawaan saat site dibuat.
  10. Setiap penambahan/penghapusan grant global tercatat di audit log (PERM-057).
  11. ⚠ Proteksi lockout: penghapusan grant Administer Jira terakhir (atau yang membuat admin yang sedang login kehilangan akses) ditolak atau memerlukan konfirmasi eksplisit dengan peringatan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga USR-022 Jira global permissions ([02](02-pengguna-grup-organisasi.md)); semantik izin dimiliki modul 03, modul 02 hanya mengelola group pemegangnya. Global permission khusus Plans dan "Jira System Administrators" adalah fitur Data Center (PERM-067). Tidak ada global permission khusus JSM/JPD.

### PERM-007 Administer Jira Global Permission

- **Deskripsi:** Global permission tertinggi di level produk yang memberi akses ke seluruh konfigurasi (scheme, workflow, field, work type, pembuatan project). Memberi product admin role identik dengan menambahkan user ke group pemegang izin ini.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira), Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Pemegang izin dapat membuat company-managed maupun team-managed project.
  2. Pemegang izin mengelola semua scheme: permission, notification, issue security, workflow, screen, field configuration, work type.
  3. Pemegang izin mengelola project role global, work type, custom field, status, resolution, dan priority.
  4. Pemegang izin mewarisi Administer Projects di semua project, termasuk mengganti scheme yang dipakai project.
  5. Pemegang izin tetap dapat login walau ada pembatasan akses tertentu yang memblokir user biasa (mis. mode maintenance produk), agar dapat memulihkan konfigurasi.
  6. Pemegang izin TIDAK dapat mengelola user, impor/ekspor data site, atau mengubah pengaturan email sistem bila tidak juga User access admin/Admin Organisasi; endpoint terkait menolak dengan 403.
  7. Pemegang izin tidak otomatis mendapat Browse Projects atas work item; untuk melihat work item ia harus memberi dirinya grant (UI project settings menampilkan petunjuk ini).
  8. Pemegang izin mengelola global permission dan General configuration (comment visibility, unassigned issues, voting, watching).
  9. Pemegang izin mengakses System settings: audit log, archived projects/work items, dan trash project.
- **Team-managed vs company-managed:** Untuk team-managed, Admin Situs selalu dapat membuka project settings termasuk project Private, tetapi tidak melihat isi work item project Private tanpa menambahkan dirinya sebagai anggota (PERM-038, PRJ-039).
- **Catatan & ketergantungan:** Pembagian kewenangan dengan User access admin/Admin Organisasi di PERM-003; halaman System settings di [22 Administrasi Sistem](22-administrasi-sistem.md).

### PERM-008 Browse Users and Groups Global Permission

- **Deskripsi:** Mengatur apakah user dapat melihat seluruh user dan group di site, yang berdampak pada user picker, group picker, dan @mention.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user (terdampak).
- **Aturan perilaku (acceptance criteria):**
  1. Dengan izin ini, user picker dan @mention menampilkan nama semua user aktif di site.
  2. Tanpa izin ini, user picker/@mention hanya menampilkan user yang memiliki Browse Projects di project saat ini; pencarian nama lain mengembalikan hasil kosong, bukan galat.
  3. Izin ini mengendalikan kemampuan melihat daftar anggota group dan memakai group picker.
  4. Izin ini dipakai saat share filter/dashboard ke user/group dan saat menambah watcher; tanpa izin, target share/watcher terbatas pada user yang dapat browse project terkait.
  5. Hasil tetap tunduk pada pengaturan privasi profil akun (email, jabatan, dll.) yang diatur pemilik akun; field privat tidak pernah tampil walau user memiliki izin ini.
  6. Default Fase 1: global permission ini diberikan ke group default akses produk (`<product>-users-<site>`) sehingga semua user berlisensi produk memegang Browse users and groups; guest dan anonim tidak pernah memegangnya (PERM-006 butir 8). Typeahead `@` dan user picker Fase 1 memakai aturan butir 1-2.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dinaikkan ke Fase 1 karena EDT-014 Mentions butir 6 ([07 Editor Konten Kaya](07-editor-konten-kaya.md)) dan NTF-013 Mention Notifications ([19 Notifikasi](19-notifikasi-email.md)) sudah bergantung pada izin ini; semantik izin dimiliki modul 03. Pengaturan privasi profil di [01 Identitas](01-identitas-autentikasi-sesi.md); endpoint pencarian user di PERM-059.

### PERM-009 Share Dashboards and Filters Global Permission

- **Deskripsi:** Mengatur siapa yang boleh membagikan saved filter dan dashboard ke user, group, project, role, organisasi, atau publik.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user (pemilik filter/dashboard).
- **Aturan perilaku (acceptance criteria):**
  1. Tanpa izin ini user hanya dapat membuat filter/dashboard Private; opsi share lain tidak tampil.
  2. Target share yang didukung: user, group, project (semua yang dapat browse project), project role tertentu dalam project, my organization/any logged-in user, dan Public (bila diizinkan admin).
  3. Pengaturan Viewers dan Editors terpisah untuk setiap filter/dashboard.
  4. Share ke Public dapat diblokir oleh pengaturan admin atau data security policy (PERM-065); opsi Public tidak tampil saat diblokir.
  5. Penerima share tetap hanya melihat work item yang boleh ia Browse; share tidak pernah memberi izin work item tambahan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga DSH-012 Share dashboards and filters global permission ([17](17-dashboard-gadget.md)); semantik izin dimiliki modul 03. UI filter di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) dan dashboard di [17 Dashboard](17-dashboard-gadget.md).

### PERM-010 Manage Group Filter Subscriptions Global Permission

- **Deskripsi:** Memberi izin membuat dan menghapus filter subscription (email terjadwal hasil filter) yang dikirim ke seluruh anggota sebuah group.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user (pembuat subscription).
- **Aturan perilaku (acceptance criteria):**
  1. Tanpa izin ini user hanya dapat membuat subscription untuk dirinya sendiri.
  2. Dengan izin ini user dapat memilih group sebagai penerima subscription.
  3. Isi email dirender per penerima dan disaring sesuai izin Browse masing-masing penerima; penerima yang tidak dapat melihat work item tidak menerima baris tersebut.
  4. ⚠ Pemegang izin dapat menghapus subscription group; apakah ia juga dapat menghapus subscription group buatan user lain ditetapkan sebelum implementasi (usulan default: hanya pembuat dan Admin Situs).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penjadwalan dan pengiriman subscription di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) dan [19 Notifikasi](19-notifikasi-email.md).

### PERM-011 Make Bulk Changes Global Permission

- **Deskripsi:** Mengatur siapa yang boleh memakai operasi bulk change (edit, move, transition, delete, watch/unwatch, archive) atas banyak work item sekaligus.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user.
- **Aturan perilaku (acceptance criteria):**
  1. Tanpa izin ini aksi "Bulk change" di issue navigator tidak tersedia dan endpoint bulk menolak dengan 403.
  2. Setiap operasi bulk tetap memerlukan izin project per work item (Edit Issues; Move Issues + Create Issues di tujuan; Transition Issues; Delete Issues; Archive; Manage Watchers; Set Issue Security).
  3. Work item yang tidak memenuhi izin dilewati dan dilaporkan dalam ringkasan hasil operasi beserta alasannya; operasi tidak dibatalkan seluruhnya.
  4. ⚠ Opsi "Send email notification" tersedia pada bulk operation; mematikannya mensyaratkan izin tambahan (usulan: Administer Projects pada semua project terdampak atau Admin Situs).
  5. Batas jumlah work item per bulk operation (mis. 1.000) berlaku terpisah dari izin dan ditampilkan di wizard.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Wizard bulk di [06 Work Item Inti](06-work-item-inti.md).

### PERM-012 Create Team-managed Projects Global Permission

- **Deskripsi:** Mengatur siapa yang boleh membuat team-managed project tanpa menjadi Admin Situs.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user.
- **Aturan perilaku (acceptance criteria):**
  1. User dengan izin ini otomatis menjadi Administrator di team-managed project yang ia buat.
  2. Company-managed project tetap hanya dapat dibuat oleh pemegang Administer Jira.
  3. Admin dapat mencabut izin ini dari group default untuk mengendalikan proliferasi project; pencabutan tidak memengaruhi project yang sudah dibuat.
  4. Dialog Create project menyembunyikan atau menonaktifkan template company-managed bagi non-admin, dengan keterangan bahwa diperlukan Admin Situs.
- **Team-managed vs company-managed:** Fitur ini khusus pembuatan team-managed project.
- **Catatan & ketergantungan:** Perilaku wizard dan entry point di PRJ-013 Create Project Permissions ([04](04-proyek.md)); definisi izin di sini adalah sumber kebenaran. Wizard Fase 1 (PRJ-010) sudah membutuhkan izin ini sehingga fase disamakan ke Fase 1; default Fase 1 memberi izin ini ke group default akses produk.

### PERM-013 Manage Custom Onboarding Global Permission

- **Deskripsi:** Global permission untuk membuat dan mengedit konten custom onboarding bagi user site serta mengakses analitik onboarding tersebut.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), user pemegang izin (penulis onboarding).
- **Aturan perilaku (acceptance criteria):**
  1. Pemegang izin dapat membuat dan mengedit konten onboarding kustom yang ditampilkan kepada user site.
  2. Pemegang izin dapat melihat analitik custom onboarding (mis. jumlah tayangan dan penyelesaian).
  3. Izin diberikan ke group seperti global permission lain.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Tier dan cakupan fitur onboarding-nya sendiri belum dipastikan; fitur onboarding didefinisikan di [21 Navigasi](21-navigasi-ux-preferensi.md) bila dibangun.

## Permission scheme & project role (company-managed)

### PERM-014 Permission Schemes

- **Deskripsi:** Kumpulan grant project permission yang dapat dipakai bersama oleh banyak company-managed project (Settings > Work items > Permission schemes).
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman daftar permission scheme menampilkan kolom nama, deskripsi, dan daftar project yang memakai.
  2. Add permission scheme: nama wajib dan unik (case-insensitive) serta deskripsi opsional; nama duplikat ditolak dengan pesan galat.
  3. Copy scheme membuat salinan lengkap seluruh grant dengan nama "Copy of <nama>" yang dapat diubah.
  4. Edit nama dan deskripsi scheme.
  5. Delete scheme hanya diizinkan bila tidak dipakai project mana pun; tombol delete dinonaktifkan dengan keterangan jumlah project pemakai.
  6. Halaman detail scheme mengelompokkan permission: Project permissions, Work item permissions, Voters & watchers, Comments, Attachments, Time tracking, dan kategori Service Management/app bila ada.
  7. Setiap baris permission menampilkan daftar grantee beserta tombol "Update" (tambah grantee) dan aksi remove per grantee.
  8. Satu project hanya memakai tepat satu permission scheme; satu scheme dapat dipakai banyak project.
  9. Permission scheme hanya berlaku untuk company-managed project; team-managed memakai access level dan role sendiri (PERM-038 s.d. PERM-040).
  10. Perubahan scheme berlaku seketika (setelah invalidasi cache, PERM-001 butir 11) untuk semua project terkait.
  11. Seluruh operasi pengelolaan scheme mensyaratkan Administer Jira; Admin Proyek hanya dapat melihat (PERM-017).
  12. Plan Free: scheme tidak dapat diedit; scheme hasil downgrade/migrasi tetap ada tetapi read-only (PERM-005).
  13. Setiap perubahan scheme dan grant tercatat di audit log (PERM-057).
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Pada Fase 1 tabel scheme/grant sudah ada dan terisi oleh default scheme (PERM-016); halaman editor lengkap dibuka di Fase 3. Data model di [28 Model Data Inti](28-model-data-inti.md); REST di PERM-060.

### PERM-015 Grant Permission Dialog & Grantee Types

- **Deskripsi:** Dialog untuk menambahkan satu atau lebih grantee pada sebuah permission di permission scheme, beserta semua tipe grantee yang didukung.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Admin memilih permission lalu "Update" untuk memilih satu atau beberapa tipe grantee sekaligus dalam satu dialog; penyimpanan bersifat atomik.
  2. Grantee Project role: semua anggota role tersebut di project yang memakai scheme.
  3. Grantee Application access: semua user yang memiliki akses ke produk tertentu (PERM-004).
  4. Grantee Group: grant menyimpan `groupId` (bukan nama) sehingga rename group tidak memutus grant.
  5. Grantee Public (anyone on the web / anonymous) menampilkan peringatan risiko yang harus dikonfirmasi; ditolak bila data security policy memblokir anonymous access.
  6. Grantee Any logged in user: semua user yang login dan memiliki product access.
  7. Grantee Current assignee: assignee work item saat evaluasi.
  8. Grantee Reporter: reporter work item saat evaluasi.
  9. Grantee Group custom field value: anggota group yang dipilih di custom field group picker (single/multi) pada work item.
  10. Grantee Project lead: project lead project terkait (terpisah dari Administer Projects).
  11. Grantee Single user: hanya user yang sudah diundang/aktif di site dapat dipilih.
  12. Grantee User custom field value: user yang dipilih di custom field user picker (single/multi) pada work item.
  13. Grantee JSM "Service project customer - portal access" hanya tersedia di scheme yang dipakai service project (PERM-036).
  14. Grant duplikat (permission + tipe + parameter sama) ditolak dengan pesan "Grant sudah ada".
  15. Grantee dinamis (Reporter, Current assignee, custom field value) tidak bermakna untuk izin konteks project; dialog menampilkan peringatan, dan evaluator tidak boleh memakai grant tersebut untuk memunculkan nama project di direktori/pencarian project (menghindari kebocoran seperti yang dikenal di produk acuan).
  16. Grant ke Any logged in user juga berlaku untuk AI agent yang memiliki akun sendiri; dialog menampilkan keterangan ini.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Grantee "Reporter (show only projects with create permission)" adalah fitur Data Center dan tidak dibangun (PERM-067).

### PERM-016 Default Permission Schemes

- **Deskripsi:** Permission scheme bawaan yang dibuat sistem dan dipakai company-managed project baru.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Sistem (seed), Admin Situs (permission Administer Jira; pengubah).
- **Aturan perilaku (acceptance criteria):**
  1. Default Permission Scheme dibuat saat inisialisasi site dengan grant ke role Administrators, role anggota (Member/Viewer), dan role `app-project-access` (role akses app di produk ini; alias impor: `atlassian-addons-project-access`).
  2. Default software scheme untuk template Scrum/Kanban mencakup Manage Sprints dan izin agile lain.
  3. Setiap service project mendapat scheme JSM bawaan dengan grant ke Service Desk Team dan customer portal access (PERM-036).
  4. Default scheme tidak dapat dihapus; tombol delete tidak tersedia.
  5. Opsi "Share settings with an existing project" saat create project membuat project baru memakai permission scheme project acuan, bukan salinannya.
  6. Isi grant default dapat berbeda tergantung versi seed site (site lama vs baru); migrasi tidak mengubah grant default site yang sudah ada.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai role bawaan (PERM-039).
- **Catatan & ketergantungan:** Isi grant default persis (matriks permission × role) harus didokumentasikan di [28 Model Data Inti](28-model-data-inti.md) sebagai seed migrasi. Migrasi dari repo saat ini: user ber-`Role=admin` dimasukkan ke group admin; user lain ke group user default.

### PERM-017 Project Settings Permissions Page

- **Deskripsi:** Halaman Project settings > Permissions (company-managed) yang menampilkan permission scheme aktif dan grant-nya.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek (lihat), Admin Situs (permission Administer Jira; ubah).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan nama scheme aktif dan daftar grant per permission; bagi Admin Proyek tampilan read-only.
  2. Menu Actions berisi "Use a different scheme" dan "Edit permissions", hanya tampil bagi Admin Situs.
  3. Wizard ganti scheme: pilih scheme dari daftar lalu "Associate"; perubahan berlaku seketika dan tercatat di audit log.
  4. Indikator "dipakai bersama oleh N project" tampil bila scheme shared, dengan peringatan bahwa edit akan memengaruhi semua project tersebut.
  5. Tautan ke Permission Helper (PERM-055) dengan konteks project terisi otomatis.
  6. Halaman Issue security di project settings menampilkan scheme security aktif (PERM-042).
- **Team-managed vs company-managed:** Team-managed tidak memiliki halaman ini; diganti halaman Access (PERM-038).
- **Catatan & ketergantungan:** Navigasi project settings di [04 Proyek](04-proyek.md).

### PERM-018 Project Roles (Global Definition)

- **Deskripsi:** Daftar role global (Settings > System > Project roles) yang anggotanya diisi per project dan dirujuk oleh scheme serta workflow.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Membuat role baru dengan nama unik (case-insensitive) dan deskripsi; nama duplikat ditolak.
  2. Mengedit nama/deskripsi role; perubahan nama tidak memutus referensi karena referensi memakai ID.
  3. Menghapus role menampilkan peringatan dampak (daftar scheme/workflow yang memakainya) dan meminta konfirmasi; grant terkait ikut dihapus.
  4. View usage: menampilkan permission scheme, notification scheme, issue security scheme, dan workflow yang merujuk role.
  5. Menampilkan anggota role per project (daftar project dan user/group anggotanya).
  6. Role bawaan: Administrators, `app-project-access` (alias impor: `atlassian-addons-project-access`), ditambah Member/Viewer (site baru) atau Developers/Users (site lama).
  7. Mengaktifkan Service Management menambahkan role Service Desk Team dan Service Desk Customers.
  8. Plan Free: role tidak dapat di-assign atau diubah.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai role lokal per project.
- **Catatan & ketergantungan:** Default members per role (perilaku DC) tidak dibangun (PERM-067). REST di PERM-061.

### PERM-019 Project Role Membership (People Page)

- **Deskripsi:** Halaman Project settings > People (company-managed) untuk menambah user/group ke project role per project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects), Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Add people": pilih satu atau beberapa user atau group, lalu pilih role; minimal satu role wajib dipilih.
  2. Satu user/group dapat memiliki banyak role di project yang sama.
  3. Role dapat diubah via dropdown per baris, dan user/group dapat dihapus dari project dengan konfirmasi.
  4. Tersedia pencarian nama dan filter berdasarkan role.
  5. Group tampil sebagai satu entri; jumlah anggotanya dihitung dinamis dari keanggotaan group.
  6. Menambah user tanpa product access memicu undangan (bila pengundang berwenang) atau access request ke admin.
  7. Admin Proyek dapat mengelola membership tanpa menjadi Admin Situs.
  8. Perubahan membership langsung berlaku di semua scheme yang merujuk role (invalidasi cache izin).
- **Team-managed vs company-managed:** Team-managed memiliki halaman People serupa dengan role lokal (PERM-039).
- **Catatan & ketergantungan:** Undangan dan access request di [02 Manajemen Pengguna](02-pengguna-grup-organisasi.md).

## Katalog project permission: akses project & work item

### PERM-020 Project Permissions: Administer Projects & Browse Projects

- **Deskripsi:** Dua permission inti kategori akses project: kemampuan mengelola pengaturan project (Administer Projects) dan kemampuan melihat project beserta work item-nya (Browse Projects).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs (permission Administer Jira), user.
- **Aturan perilaku (acceptance criteria):**
  1. Administer Projects: mengedit nama, key, avatar, deskripsi, URL, kategori, project lead, dan default assignee project.
  2. Administer Projects: mengelola People (assign project role), components, dan versions.
  3. Administer Projects: menambah, mengonfigurasi, dan menghapus AI agent di project (PERM-050).
  4. Administer Projects tidak mencakup mengganti scheme yang dipakai project (hanya Admin Situs); aksi tersebut disembunyikan bagi Admin Proyek.
  5. Administer Projects diwarisi Admin Situs dan hanya dapat diberikan ke user yang sudah diundang (bukan Public/anonymous); dialog grant menolak kombinasi tersebut.
  6. Plan Free: semua user memiliki Administer Projects.
  7. Browse Projects: melihat project, work item, hasil pencarian, board, laporan, dan timeline, kecuali work item yang disembunyikan issue security.
  8. Browse Projects tidak memerlukan product access sehingga dapat diberikan ke siapa pun termasuk Public (cara umum membuat project publik, PERM-046).
  9. Browse Projects adalah syarat menerima notifikasi email tentang work item project dan syarat muncul di user picker bagi user tanpa Browse users and groups.
  10. Kehilangan Browse Projects membuat project hilang dari pencarian, daftar recent, starred, dan gadget; tautan langsung menampilkan halaman tanpa akses (PERM-056).
- **Team-managed vs company-managed:** Team-managed: Administer Projects = role Administrator; Browse Projects ditentukan access level dan keanggotaan role.
- **Catatan & ketergantungan:** Pengaturan project di [04 Proyek](04-proyek.md).

### PERM-021 Issue Permissions: Create, Edit, Delete Issues

- **Deskripsi:** Permission dasar untuk membuat, mengubah, dan menghapus work item dalam project.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Create Issues: project muncul di project picker dialog Create hanya bila user memiliki izin ini; izin mencakup membuat sub-task dan clone.
  2. Create Issues diperlukan di project tujuan saat Move dan Clone; tanpa izin, project tujuan tidak tampil di wizard.
  3. Create Issues dapat berlaku tanpa Browse Projects dan tanpa product access (dapat diberikan ke Public); dialog grant menampilkan peringatan risiko spam.
  4. Edit Issues: mengedit summary, description, dan field pada edit screen, termasuk inline edit di issue view, board, dan list.
  5. Edit Issues tidak mencakup field yang diatur permission khusus: Assignee (Assign Issues), Reporter (Modify Reporter), Due date/rank (Schedule Issues), Security level (Set Issue Security), Resolution/Fix version (Resolve Issues); field tersebut read-only bagi yang tidak memiliki izin khususnya.
  6. Delete Issues: menghapus work item beserta data terkait (komentar, attachment, worklog, sub-task) secara permanen setelah konfirmasi yang menyebut jumlah sub-task terdampak.
  7. Delete Issues berlaku walau user tidak memiliki Delete All Comments/Delete All Attachments, tetapi tidak memberi hak menghapus komentar/attachment satuan.
  8. Default scheme membatasi Delete Issues ke role Administrators; UI menawarkan Archive (PERM-035) sebagai alternatif bila tersedia.
  9. Workflow property `jira.issue.editable=false` membuat work item read-only di status tertentu walau user memiliki Edit Issues (PERM-037).
- **Team-managed vs company-managed:** Team-managed: setara izin role "Create work items", "Edit any work item", "Delete any work item".
- **Catatan & ketergantungan:** Operasi work item di [06 Work Item Inti](06-work-item-inti.md).

### PERM-022 Assign Issues & Assignable User Permissions

- **Deskripsi:** Pasangan permission yang mengatur siapa yang boleh mengubah assignee (Assign Issues) dan siapa yang boleh dijadikan assignee (Assignable User).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Assign Issues: mengubah field Assignee dan memakai tombol "Assign to me".
  2. Assign Issues tidak membuat user lain menjadi assignable; kedua izin dievaluasi terpisah.
  3. Assignable User: hanya user dengan izin ini muncul di assignee picker; menetapkan user tanpa izin ini ditolak.
  4. Assignable User tidak memberi hak mengubah field Assignee.
  5. Project lead dan default assignee seharusnya Assignable User (UI memperingatkan bila tidak); component lead dapat menjadi default assignee per component.
  6. Opsi global "Allow unassigned issues" menentukan apakah nilai Unassigned valid; bila nonaktif, create/edit tanpa assignee ditolak dan default assignee wajib diisi.
  7. Default assignee diterapkan saat create bila user tidak memiliki Assign Issues atau tidak mengisi assignee.
  8. Assign via REST/automation memvalidasi bahwa target adalah Assignable User; bila tidak, respons 400 dengan pesan "User tidak dapat di-assign pada work item ini".
  9. User yang kehilangan Assignable User tetap tercatat sebagai assignee pada work item lama sampai diubah.
  10. Team-managed: role yang memiliki izin dari set "Work on" otomatis assignable.
- **Team-managed vs company-managed:** Lihat butir 10; team-managed tidak memiliki izin Assignable User terpisah.
- **Catatan & ketergantungan:** Default assignee project/component di [04 Proyek](04-proyek.md) dan [15 Rilis & Komponen](15-rilis-versi-komponen.md); pencarian user assignable di PERM-059.

### PERM-023 Transition, Resolve & Close Issues Permissions

- **Deskripsi:** Permission untuk memindahkan work item antar status serta mengatur resolution dan fix version.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Transition Issues: menjalankan transisi workflow (tombol status, drag kartu di board).
  2. Di company-managed, Transition Issues dikombinasikan dengan workflow conditions (mis. "User is in project role", "Permission condition"); transisi hanya tampil bila keduanya terpenuhi.
  3. Resolve Issues: mengisi/mengubah field Resolution dan Fix version.
  4. Close Issues: izin legacy untuk menetapkan resolution ke status tertutup; dipakai sebagai syarat pada condition transisi tertentu.
  5. Drag kartu ke kolom lain ditolak dengan pesan "Anda tidak memiliki izin untuk memindahkan work item ini" dan kartu kembali ke posisi semula bila user tidak memiliki Transition Issues.
  6. Team-managed: setara izin role "Transition any work item".
- **Team-managed vs company-managed:** Lihat butir 6; team-managed tidak memakai izin Resolve/Close terpisah.
- **Catatan & ketergantungan:** Conditions di [09 Workflow](09-workflow-status-transisi.md); drag di board di [10 Board](10-board-kolom-kartu.md).

### PERM-024 Link, Move, Schedule & Modify Reporter Permissions

- **Deskripsi:** Permission work item tambahan untuk linking, memindahkan antar project/work type, menjadwalkan, dan mengubah reporter.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Link Issues: membuat dan menghapus issue link, hanya bila issue linking aktif di site.
  2. Link antar project hanya dapat dibuat dan tampil lengkap bila user memiliki Link Issues di kedua project (sumber dan target); bila user hanya memiliki Browse Projects di project target, link tampil read-only tanpa aksi hapus, dan link ke work item yang tidak boleh dilihat ditampilkan tanpa detail (key saja atau disembunyikan).
  3. Move Issues: memindahkan work item ke project lain atau mengganti work type; mensyaratkan Create Issues di project tujuan.
  4. Wizard move menangani mapping status, field, dan security level; work item tidak dipindah sampai semua mapping wajib terisi.
  5. Schedule Issues: mengubah Due date dan ranking di board/backlog.
  6. Menambah/memindahkan work item ke sprint mensyaratkan Schedule Issues + Edit Issues.
  7. Modify Reporter: mengubah field Reporter dan memilih reporter saat create.
  8. Tanpa Modify Reporter, reporter otomatis diisi user pembuat dan field Reporter tidak tampil di create screen.
  9. Default Fase 1 (sebelum editor permission scheme dibuka di Fase 3): default permission scheme (PERM-016) memberi Link Issues, Schedule Issues, dan Move Issues ke role anggota (Member) dan Administrators, sedangkan Modify Reporter hanya ke Administrators; di team-managed padanannya mengikuti role bawaan PERM-039. Dengan demikian setiap pemegang Edit Issues pada role Member di Fase 1 juga dapat membuat link, me-rank backlog, mengubah due date, dan memindah work item ke sprint.
- **Team-managed vs company-managed:** Team-managed: "Link any work item", "Move any work item", "Edit any due date", "Edit reporters".
- **Catatan & ketergantungan:** Keempat permission key didaftarkan dan ditegakkan sejak Fase 1 karena fitur P0 Fase 1 bergantung padanya: ISS-029 Issue Links butir 6 (Link Issues) di [06 Work Item Inti](06-work-item-inti.md), SPR-003 Backlog Ranking butir 6 (Schedule Issues) dan SPR-012 Sprint Field (Schedule + Edit Issues) di [11 Backlog & Sprint](11-backlog-sprint.md), serta PERM-025 butir 4. Penegakan Move Issues baru teruji end-to-end saat wizard move di modul 06 tersedia; Modify Reporter dipakai create/edit screen modul 06. Aturan default Fase 1 ada di butir 9; semantik izin dimiliki modul 03.

## Katalog project permission: perencanaan & pengembangan

### PERM-025 Manage Sprints Permission

- **Deskripsi:** Project permission untuk membuat, memulai, menyelesaikan sprint, serta mengubah durasi dan goal sprint di board Scrum.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek, Scrum master, user.
- **Aturan perilaku (acceptance criteria):**
  1. Membuat sprint, mengedit nama/tanggal/goal, dan mengurutkan ulang sprint di backlog.
  2. Start sprint dan Complete sprint.
  3. Delete sprint.
  4. Tidak mencakup memindah work item ke/dari sprint; itu mensyaratkan Schedule Issues + Edit Issues (PERM-024).
  5. Untuk sprint yang mencakup banyak board/project, Manage Sprints dibutuhkan di semua project yang work item-nya ada di sprint; bila kurang di salah satu, aksi ditolak dengan daftar project yang kurang izin.
  6. Board dengan filter JQL lintas project yang kompleks dapat mensyaratkan izin ini di semua project yang tercakup filter; UI menjelaskan syarat ini saat tombol dinonaktifkan.
  7. Team-managed: setara izin role "Manage sprints".
- **Team-managed vs company-managed:** Lihat butir 7.
- **Catatan & ketergantungan:** Lihat juga SPR-013 Manage Sprints Permission ([11](11-backlog-sprint.md)); definisi dan semantik izin dimiliki modul 03 (fitur ini), penegakan pada aksi sprint ada di SPR-013. Pemisahan Start/Complete vs Edit Sprints adalah fitur Data Center (PERM-067). Sprint di [11 Backlog & Sprint](11-backlog-sprint.md).

### PERM-026 Manage Versions Permission

- **Deskripsi:** Project permission untuk mengelola versions/releases tanpa harus menjadi Admin Proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, release manager, user.
- **Aturan perilaku (acceptance criteria):**
  1. Create, edit, release, archive, merge, dan delete version di halaman Releases.
  2. Izin ini terpisah dari Administer Projects; pemegang Administer Projects juga dapat mengelola version tanpa izin ini.
  3. Mengatur release date dan start date version.
- **Team-managed vs company-managed:** Team-managed: pengelolaan release mengikuti role Administrator (atau izin setara bila tersedia).
- **Catatan & ketergantungan:** Detail release di [15 Rilis & Komponen](15-rilis-versi-komponen.md).

### PERM-027 View Development Tools & View Read-only Workflow Permissions

- **Deskripsi:** Permission untuk melihat informasi development (branch, commit, PR, build, deployment) dan diagram workflow read-only dari work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** User, developer.
- **Aturan perilaku (acceptance criteria):**
  1. View Development Tools: menampilkan panel Development di issue view.
  2. View Development Tools: menampilkan indikator dev info di board, halaman release, dan deployments.
  3. View Development Tools hanya dapat diberikan ke user yang diundang (bukan Public).
  4. View Read-Only Workflow: tautan "View workflow" di dekat status work item menampilkan diagram workflow.
  5. View Read-Only Workflow tidak memerlukan product access.
  6. Tanpa View Read-Only Workflow user tetap dapat melakukan transisi (bila berizin) tetapi tautan diagram tidak tampil.
- **Team-managed vs company-managed:** Team-managed: "Access development tools"; diagram workflow mengikuti akses project.
- **Catatan & ketergantungan:** Permission key didaftarkan di Fase 2; panel development baru berisi data setelah integrasi [24 Integrasi DevOps](24-integrasi-devops.md) (Fase 4).

### PERM-028 View Aggregated Data Permission

- **Deskripsi:** Project permission untuk melihat insight agregat board, backlog, dan deployment.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** User, Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Menampilkan panel insights pada board dan backlog.
  2. Menampilkan data agregat deployment (mis. frekuensi deployment).
  3. Tidak memerlukan product access sehingga dapat diberikan ke siapa saja.
  4. Dialog grant memperingatkan risiko: data agregat dapat mengekspos rata-rata, jumlah work item, work type, dan nama epic kepada pemegang izin tanpa Browse Projects.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Panel insights di [16 Laporan & Insight](16-laporan-insight.md). Pengecualian terhadap PERM-001 butir 10.

### PERM-029 Edit Workflows & Manage Work Item Layouts Permissions

- **Deskripsi:** Project permission di company-managed project yang memungkinkan non-Admin Situs mengubah workflow dan layout work item secara terbatas (pengganti model "Extended project administration").
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Edit workflows: pemegang izin dapat mengubah status dan transisi di workflow settings memakai workflow editor baru.
  2. Edit workflows hanya berlaku untuk workflow yang tidak dipakai project lain; workflow shared tampil read-only dengan keterangan "Dipakai oleh N project".
  3. Manage work item layouts: pemegang izin dapat menambah field yang sudah ada ke work item layout, tetapi tidak dapat membuat field baru.
  4. Manage work item layouts hanya berlaku bila screen/layout tidak dipakai bersama project lain.
  5. Kedua izin mensyaratkan Browse Projects.
  6. Kedua izin tidak dapat diberikan di scheme yang dipakai service project (JSM); dialog grant menyembunyikannya.
  7. Kedua izin hanya dapat diberikan ke user yang diundang (bukan Public/anonymous).
- **Team-managed vs company-managed:** Hanya company-managed; di team-managed role Administrator sudah mengelola workflow dan field lokal.
- **Catatan & ketergantungan:** Batasan persis (mis. boleh/tidak mengedit conditions, validators, post functions) ditetapkan bersama [09 Workflow](09-workflow-status-transisi.md); layout di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md). Model DC "Extended project administration" tidak dibangun (PERM-067).

## Katalog project permission: kolaborasi, waktu & arsip

### PERM-030 Comment Permissions

- **Deskripsi:** Permission untuk menambah, mengedit, dan menghapus komentar milik sendiri atau milik semua orang.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Add Comments: menambah komentar pada work item yang dapat dilihat; tanpa izin, editor komentar tidak tampil.
  2. Edit All Comments: mengedit komentar siapa pun.
  3. Edit Own Comments: mengedit komentar milik sendiri saja.
  4. Delete All Comments: menghapus komentar siapa pun.
  5. Delete Own Comments: menghapus komentar milik sendiri saja.
  6. Komentar dengan visibility terbatas hanya terlihat oleh anggota role/group target (PERM-045); izin edit/delete tidak membuat komentar terbatas terlihat oleh non-anggota.
  7. Komentar via email reply tetap diperiksa izin Add Comments pengirim; email dari pengirim tanpa izin ditolak dan pengirim menerima balasan penolakan.
  8. Komentar internal JSM mengikuti aturan agent (hanya Service Project Agent yang dapat melihat/menulis catatan internal, lihat [29 JSM](29-jsm-service-project-request-type-formulir.md)).
- **Team-managed vs company-managed:** Team-managed: "Add comments", "Edit any comment", "Edit their own comments", "Delete any comment", "Delete their own comments".
- **Catatan & ketergantungan:** Fungsi komentar di [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md); email masuk di [19 Notifikasi](19-notifikasi-email.md).

### PERM-031 Attachment Permissions

- **Deskripsi:** Permission untuk mengunggah dan menghapus attachment work item.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Create Attachments: upload file, paste gambar, drag-drop, dan attachment via email; tanpa izin, kontrol upload tidak tampil dan upload via email diabaikan dengan catatan.
  2. Delete All Attachments: menghapus attachment siapa pun.
  3. Delete Own Attachments: menghapus attachment milik sendiri saja.
  4. Melihat/mengunduh attachment cukup dengan Browse Projects (dan keanggotaan security level bila ada); URL unduhan selalu diperiksa izin, tidak bergantung URL rahasia.
  5. Izin hanya relevan bila attachment aktif di site; batas ukuran file diatur global dan upload melebihi batas ditolak dengan pesan yang menyebut batasnya.
  6. Data security policy dapat memblokir unduhan/ekspor attachment (PERM-065).
- **Team-managed vs company-managed:** Team-managed: "Add attachments", "Delete any attachment", "Delete their own attachments".
- **Catatan & ketergantungan:** Penyimpanan file di [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md).

### PERM-032 Time Tracking Permissions

- **Deskripsi:** Permission untuk mencatat kerja dan mengelola worklog milik sendiri atau milik orang lain.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Work On Issues: log work dan mengubah estimasi waktu (original/remaining).
  2. Edit Own Worklogs: mengedit worklog milik sendiri.
  3. Edit All Worklogs: mengedit worklog siapa pun.
  4. Delete Own Worklogs: menghapus worklog milik sendiri.
  5. Delete All Worklogs: menghapus worklog siapa pun.
  6. Izin hanya relevan bila time tracking aktif di site; bila nonaktif, kontrol log work tidak tampil.
  7. Worklog dengan visibility terbatas hanya terlihat oleh role/group target (PERM-045).
  8. Team-managed: izin "Log work on any work item", "Edit their own/any work log entry", "Delete their own/any work log entry".
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Time tracking di [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md).

### PERM-033 Voters & Watchers Permissions

- **Deskripsi:** Permission untuk melihat daftar voter/watcher dan mengelola watcher work item.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. View Voters and Watchers: melihat nama voter dan watcher; tanpa izin hanya jumlahnya yang tampil.
  2. Manage Watchers: menambah dan menghapus user lain sebagai watcher.
  3. Watch/unwatch diri sendiri dan vote cukup dengan Browse Projects.
  4. Watcher yang kehilangan Browse tidak menerima notifikasi tentang work item tersebut (tetap tercatat sebagai watcher).
  5. Voting dan watching dapat dimatikan global di General configuration; bila dimatikan, kontrol terkait hilang dan API menolak.
  6. Default Fase 1: default permission scheme (PERM-016) memberi View Voters and Watchers ke role anggota (Member, Viewer) dan Administrators, sedangkan Manage Watchers hanya ke Administrators; di team-managed, "View watchers" dimiliki semua role bawaan dan "Add or remove work item watchers" hanya Administrator (PERM-039 butir 2).
  7. ⚠ Clause JQL `watcher` dan `voter` hanya mengembalikan atau mencocokkan nama watcher/voter bila user memiliki View Voters and Watchers pada project terkait; tanpa izin, nilai selain `currentUser()` diabaikan untuk project tersebut.
- **Team-managed vs company-managed:** Team-managed: "View watchers" dan "Add or remove work item watchers".
- **Catatan & ketergantungan:** Dinaikkan ke Fase 1 karena ISS-036 Watchers butir 2-3 ([06 Work Item Inti](06-work-item-inti.md)), NTF-012 Watchers ([19 Notifikasi](19-notifikasi-email.md)), serta JQL-021 dan JQL-051 ([12 JQL](12-jql-bahasa-kueri.md)) sudah bergantung pada izin ini; semantik izin dimiliki modul 03. General configuration di [22 Administrasi Sistem](22-administrasi-sistem.md).

### PERM-034 Set Issue Security Permission

- **Deskripsi:** Permission untuk mengubah security level work item pada project yang memakai issue security scheme.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** User, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Field Security Level tampil dan dapat diedit di create/edit hanya bagi pemegang izin ini.
  2. Izin hanya relevan bila project memiliki issue security scheme (PERM-042).
  3. Field Security Level harus ada di layout/screen project agar dapat diedit; bila tidak ada, Permission Helper menjelaskan penyebabnya.
  4. Izin ini dapat diberikan ke guest.
  5. Tidak tersedia di plan Free.
- **Team-managed vs company-managed:** Team-managed memakai issue restriction (PERM-041).
- **Catatan & ketergantungan:** Enforcement di PERM-044.

### PERM-035 Archive & Restore Archived Issues Permissions

- **Deskripsi:** Permission untuk mengarsipkan work item agar tersembunyi dari pencarian dan mengembalikannya (Premium/Enterprise).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Proyek, Admin Situs (permission Administer Jira), user.
- **Aturan perilaku (acceptance criteria):**
  1. Archive work items: mengarsipkan work item tunggal atau bulk; sering diberikan sebagai pengganti Delete.
  2. Restore archived work items: mengembalikan work item arsip ke kondisi aktif.
  3. Work item archived tidak muncul di JQL biasa, board, dan laporan; tersedia di halaman Archived work items.
  4. Work item archived bersifat read-only untuk semua user.
  5. Sub-task ikut terarsip (dan ter-restore) bersama parent.
  6. Bulk archive via JQL dibatasi jumlah maksimum per operasi; permintaan di atas batas ditolak dengan pesan batasnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fase disamakan ke Fase 3 dengan fitur yang membutuhkannya: ISS-034 Archive / Unarchive Work Items ([06 Work Item Inti](06-work-item-inti.md)) dan SRC-035 Archived work items search ([13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md)); semantik izin dimiliki modul 03. Endpoint publik API-027 Issue Archiving API ([23 REST API Publik & Webhook](23-api-publik-webhook.md)) dapat menyusul di Fase 4; sebelum itu operasi arsip tersedia lewat UI dan API internal SPA dengan pemeriksaan izin yang sama. Halaman arsip di [22 Administrasi Sistem](22-administrasi-sistem.md).

## Izin per tipe project & per status

### PERM-036 Service Project Agent & Customer Permissions

- **Deskripsi:** Permission dan grantee khusus Service Management di permission scheme service project.
- **Prioritas / Fase / Tier:** P3 / Fase 5 / jsm
- **Aktor & izin:** Agent, customer, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Service Project Agent: memberi akses ke queues, SLA, laporan, knowledge base, dan kemampuan menjawab customer; mensyaratkan lisensi agent.
  2. Grantee "Service project customer - portal access" memberi izin customer lewat portal tanpa product access.
  3. Scheme bawaan JSM memakai role Service Desk Team dan Service Desk Customers.
  4. Pengaturan siapa yang boleh mengirim request (customer permissions) dikelola di modul JSM, bukan di permission scheme.
  5. Edit workflows dan Manage work item layouts tidak dapat diberikan di service project.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail di [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) dan [30 JSM Portal & Customers](30-jsm-help-center-portal-pelanggan-kanal.md).

### PERM-037 Workflow Permission Properties

- **Deskripsi:** Property workflow yang mengubah izin per status di company-managed project.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Property `jira.permission.*` pada status membatasi (atau menambah) izin tertentu pada status tersebut; daftar key yang didukung dipublikasikan di dokumentasi admin.
  2. Property `jira.issue.editable=false` membuat work item read-only di status tertentu (edit field ditolak walau memiliki Edit Issues).
  3. Contoh yang harus didukung dan diuji: menolak edit dan komentar pada status Done.
  4. Condition berbasis permission dan project role pada transisi dievaluasi oleh Authorizer yang sama.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Editor property di [09 Workflow](09-workflow-status-transisi.md).

## Akses team-managed project

### PERM-038 Team-managed Project Access Levels

- **Deskripsi:** Pengaturan akses team-managed project: Open, Limited, atau Private, yang menggantikan permission scheme pada tipe project ini.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek (role Administrator), Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Open: semua user ber-product access di site dapat melihat, membuat, dan mengedit work item; user login yang belum ditambahkan diperlakukan sebagai role Member.
  2. Limited: semua user di site dapat melihat dan berkomentar, tetapi tidak membuat/mengedit; user login yang belum ditambahkan diperlakukan sebagai role Viewer.
  3. Private: hanya orang/agent yang ditambahkan sebagai anggota yang dapat melihat project dan work item-nya, termasuk di hasil pencarian; project tidak muncul di direktori bagi orang lain. Mengubah ke Private langsung mencabut akses non-member, dan filter/board/dashboard milik non-member tidak lagi menampilkan work item project tersebut.
  4. Access level diatur di Project settings > Access oleh Administrator project; perubahan berlaku seketika.
  5. Anonymous access tidak tersedia untuk team-managed project; tidak ada opsi Public.
  6. ⚠ Admin Situs selalu dapat melihat project Private di Settings > Projects dan membuka project settings-nya (termasuk Access) untuk mengelola akses, tetapi tidak dapat melihat isi work item tanpa menambahkan dirinya sebagai anggota; aksi penambahan diri tersebut tercatat di audit log (selaras PERM-001 butir 8 dan PRJ-039 butir 7).
  7. Kolaborator eksternal hanya dapat diundang sebagai guest (PERM-047).
  8. ⚠ Plan Free: permission dan role project tidak dapat disunting dan opsi Private tidak tersedia (opsi disembunyikan/nonaktif dengan keterangan plan minimum), kecuali pemilik produk memutuskan sebaliknya (selaras PRJ-039 butir 8).
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah grant Browse Projects di permission scheme.
- **Catatan & ketergantungan:** Aturan access level ini identik dengan PRJ-039 Project Access Levels (team-managed) di [04 Proyek](04-proyek.md); modul 04 memiliki sisi UI/wizard (pemilihan saat create project, indikator gembok), modul ini memiliki enforcement di Authorizer (PERM-002). Keputusan visibilitas work item Private bagi Admin Situs dan ketersediaan Private di Free mengikuti Pertanyaan terbuka PRJ-039.

### PERM-039 Team-managed Built-in Roles

- **Deskripsi:** Role bawaan team-managed project: Administrator, Member, dan Viewer.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Proyek, user.
- **Aturan perilaku (acceptance criteria):**
  1. Administrator: semua izin termasuk administer project dan seluruh izin work item.
  2. Member: create, edit, assign, link, log work, move, transition work item, manage sprints, edit due date, akses development tools; TIDAK dapat edit reporter, menghapus konten orang lain, atau manage watchers.
  3. Viewer: melihat work item, menambah komentar dan attachment, serta mengedit/menghapus komentar dan attachment miliknya.
  4. Semua user yang ditambahkan ke project otomatis mendapat izin melihat project dan melihat work item.
  5. Menambah orang/group dengan role lewat dialog "Add people".
  6. Role anggota dapat diubah dan anggota dapat dihapus; Administrator terakhir tidak dapat dihapus atau diturunkan (pesan galat menjelaskan).
  7. Pembuat project otomatis menjadi Administrator.
  8. Plan Free: role tidak dapat diubah.
- **Team-managed vs company-managed:** Khusus team-managed.
- **Catatan & ketergantungan:** Matriks izin role bawaan didokumentasikan sebagai seed di [28 Model Data Inti](28-model-data-inti.md).

### PERM-040 Team-managed Custom Roles & Granular Permissions

- **Deskripsi:** Membuat role khusus di team-managed project dengan checklist izin granular.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Proyek (role Administrator).
- **Aturan perilaku (acceptance criteria):**
  1. Create role: nama (unik dalam project), deskripsi, dan checklist izin; minimal satu izin dipilih.
  2. Izin yang tersedia: Access development tools, Add attachments, Add comments, Add or remove work item watchers, Administer the space, Assign any work item, Create work items, Delete any attachment, Delete any comment, Delete any work item, Delete any work log entry, Delete their own attachments, Delete their own comments, Delete their own work log entries, Edit any comment, Edit any due date, Edit any work item, Edit any work log entry, Edit reporters, Edit their own comments, Edit their own work log entries, Link any work item, Log work on any work item, Manage sprints, Move any work item, Transition any work item, View watchers.
  3. Izin dikelompokkan dalam set: Collaborate on work items, Manage space work items, Work on space work items.
  4. Memberi izin dari set "Work on" membuat anggota role tersebut assignable.
  5. Custom role dapat diedit dan dihapus; menghapus role yang masih memiliki anggota meminta pemilihan role pengganti untuk anggota tersebut.
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah project role + permission scheme.
- **Catatan & ketergantungan:** ⚠ Tier minimum (Premium atau Standard ke atas) belum dipastikan; yang pasti tidak tersedia di Free.

### PERM-041 Team-managed Issue Restrictions

- **Deskripsi:** Membatasi siapa yang dapat melihat work item tertentu di team-managed project berdasarkan role (padanan issue security).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek, user berizin.
- **Aturan perilaku (acceptance criteria):**
  1. Kontrol restriction (ikon gembok) di issue view memungkinkan memilih satu atau beberapa role yang boleh melihat work item; kontrol hanya tampil bagi user yang berwenang menyetel restriction.
  2. Hanya anggota role terpilih (plus Admin Situs dengan akses project) yang dapat melihat work item restricted.
  3. ⚠ Restriction berlaku juga pada child/sub-task work item tersebut.
  4. Work item restricted disaring dari JQL, board, dan notifikasi bagi non-anggota.
  5. Tidak tersedia di plan Free.
- **Team-managed vs company-managed:** Khusus team-managed; padanan company-managed adalah issue security (PERM-042 s.d. PERM-044).
- **Catatan & ketergantungan:** Tersedia mulai Standard. ⚠ Izin role yang berwenang menyetel restriction belum dipastikan; usulan: Administrator dan pemegang "Edit any work item".

## Issue security & visibilitas konten

### PERM-042 Issue Security Schemes

- **Deskripsi:** Scheme berisi security level untuk membatasi visibilitas work item di company-managed project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman daftar scheme menampilkan nama scheme dan project yang memakai.
  2. Add scheme: nama unik dan deskripsi.
  3. Copy, edit, dan delete scheme; delete hanya diizinkan bila scheme tidak dipakai project.
  4. Asosiasi scheme ke project dilakukan lewat Project settings > Issue security.
  5. Wizard ganti/lepas scheme meminta mapping setiap level lama ke level baru (atau None) untuk work item yang sudah ada; proses berjalan sebagai background job dengan indikator progres.
  6. Satu project memakai maksimal satu issue security scheme (opsional).
  7. Dokumentasi admin menyarankan memakai Browse Projects/data classification, bukan security scheme, untuk membatasi seluruh project.
  8. Tidak tersedia di plan Free; scheme hasil downgrade tetap ada tetapi read-only.
  9. Hanya untuk company-managed.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed memakai PERM-041.
- **Catatan & ketergantungan:** Data model di [28 Model Data Inti](28-model-data-inti.md); REST di PERM-062.

### PERM-043 Issue Security Levels & Members

- **Deskripsi:** Pengelolaan security level di dalam issue security scheme beserta anggota tiap level.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Add level: nama (unik dalam scheme) dan deskripsi.
  2. Edit nama/deskripsi level.
  3. Set default level atau "tanpa default"; hanya satu default per scheme.
  4. Hapus level: work item yang memakai level tersebut wajib dipetakan ke level lain atau None sebelum penghapusan selesai.
  5. ⚠ Add member dengan tipe: Reporter, Current assignee, Project lead, Group, Project role, Single user, User custom field value, Group custom field value, Application access/Any logged in user.
  6. Hapus member dari level; perubahan berlaku seketika pada visibilitas.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Tipe member memakai resolver grantee yang sama dengan PERM-015.

### PERM-044 Issue Security Enforcement

- **Deskripsi:** Perilaku runtime security level pada work item: siapa yang dapat menyetel, efek visibilitas, dan pewarisan.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** User, sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Default level diterapkan saat create bila field tidak diisi (termasuk create via REST, email, automation).
  2. Dropdown Security Level hanya menampilkan level yang user sendiri menjadi anggotanya, agar user tidak mengunci dirinya sendiri.
  3. Sub-task mewarisi level parent dan tidak dapat diubah sendiri; mengubah level parent ikut mengubah sub-task.
  4. Work item dengan level disaring dari JQL, board, laporan, notifikasi, gadget, dan REST bagi non-anggota.
  5. Move ke project lain menerapkan default level project tujuan; bila project tujuan tanpa scheme, work item terlihat oleh semua yang memiliki akses project tujuan (wizard memperingatkan).
  6. Bulk change yang memindah/mengubah work item ber-level meminta mapping level.
  7. JQL mendukung field `level` (mis. `level = "Internal"`), hanya dengan nilai level yang dapat dilihat user.
  8. Akses ke work item tanpa izin ditampilkan sebagai "tidak ditemukan/tidak punya akses" (HTTP 404 di API) tanpa membocorkan summary.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Field `level` di [12 JQL](12-jql-bahasa-kueri.md).

### PERM-045 Comment & Worklog Visibility Restriction

- **Deskripsi:** Membatasi visibilitas komentar dan worklog ke project role atau group tertentu.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** User (penulis), Admin Situs (permission Administer Jira; pengatur opsi global).
- **Aturan perilaku (acceptance criteria):**
  1. Pilihan visibility (ikon gembok) tersedia di editor komentar dan dialog log work.
  2. Target visibility: project role atau group (bila diizinkan setting global); dropdown hanya menampilkan role/group yang user sendiri menjadi anggotanya.
  3. Setting global "Comment visibility" di General configuration: "role saja" atau "group dan role".
  4. Komentar restricted tidak dikirim lewat notifikasi (email/in-app) ke non-anggota target.
  5. Hanya anggota target yang dapat melihat; mengedit tetap memerlukan izin Edit Own/All Comments (atau Worklogs).
- **Team-managed vs company-managed:** Team-managed membatasi ke role project.
- **Catatan & ketergantungan:** Visibilitas juga berlaku pada history, REST, export, dan webhook (PERM-002 butir 2).

## Akses khusus: publik, guest, app, automation & AI agent

### PERM-046 Anonymous & Public Access

- **Deskripsi:** Memberikan akses company-managed project ke pengguna anonim lewat grantee Public.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira; memberi), pengunjung anonim.
- **Aturan perilaku (acceptance criteria):**
  1. Grant Browse Projects ke Public membuat project dan work item-nya dapat dilihat tanpa login (kecuali yang dilindungi issue security).
  2. Create Issues dan Add Comments dapat diberikan ke Public; UI menandai risiko spam dan sistem menerapkan rate limit/CAPTCHA untuk subjek anonim.
  3. Admin melihat peringatan eksplisit yang harus dikonfirmasi saat memberi grant Public, dan daftar project berakses publik tampil di halaman administrasi.
  4. Data security policy dapat memblokir anonymous access (PERM-065); saat aktif, grant Public diabaikan oleh evaluator dan ditandai di UI.
  5. Filter/dashboard dapat di-share ke Public (PERM-009) dan hanya menampilkan work item yang boleh dilihat anonim.
  6. Tidak berlaku untuk team-managed project.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Lihat juga USR-024 Anonymous & public access ([02](02-pengguna-grup-organisasi.md)); semantik izin dimiliki modul 03. Subjek anonim dievaluasi oleh Authorizer (PERM-002 butir 5); proteksi abuse di [27 NFR](27-nfr-arsitektur-platform.md).

### PERM-047 Guest Access

- **Deskripsi:** Akses terbatas bagi kolaborator eksternal yang hanya bekerja di project yang mengundangnya.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** User access admin, Admin Proyek (pengundang), guest.
- **Aturan perilaku (acceptance criteria):**
  1. Guest gratis hingga 5 per paid user, tetapi tetap dihitung dalam batas user site; melebihi kuota ditolak dengan pesan.
  2. Guest tidak pernah mendapat global permission apa pun.
  3. Company-managed: Browse guest terbatas pada satu project; guest dapat diberi izin konfigurasi dan izin work item (administer, edit workflow, manage sprint, create/edit/transition/delete, komentar, attachment, worklog).
  4. Company-managed: guest hanya dapat link dan move work item dalam project yang sama.
  5. Team-managed: guest mendapat set izin nested yang lebih luas, termasuk administer project dan menghapus konten siapa pun, tetap terbatas pada project tersebut.
- **Team-managed vs company-managed:** Lihat butir 3-5.
- **Catatan & ketergantungan:** ⚠ Plan yang mendukung guest belum dipastikan. Lihat juga USR-019 Guest access (Jira guests) ([02](02-pengguna-grup-organisasi.md)) untuk undangan, kuota, dan siklus hidup akun guest; semantik izin dimiliki modul 03.

### PERM-048 App Permissions

- **Deskripsi:** Izin untuk app/ekstensi yang terpasang di produk ini (padanan app Connect/Forge).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (permission Administer Jira), app.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap app terpasang memiliki app user dan dimasukkan ke role `app-project-access` (alias impor: `atlassian-addons-project-access`) di setiap project.
  2. Tersedia group `app-users-admin` (alias impor: `atlassian-addons-admin`; selaras USR-017 butir 5 di [02](02-pengguna-grup-organisasi.md)) untuk app yang memerlukan akses admin.
  3. Scope app (read/write/manage) ditampilkan dan disetujui admin saat instal; app tidak dapat memakai API di luar scope (403).
  4. App dapat bertindak asUser (dievaluasi dengan izin user) atau asApp (dievaluasi dengan izin app user).
  5. App dapat mendefinisikan global permission dan project permission sendiri yang muncul di UI global permissions dan permission scheme.
  6. Data security policy dapat memuat app access rule yang memblokir app dari data tertentu (PERM-065).
  7. Akses app ke project tetap dibatasi permission scheme; menghapus role app dari scheme memutus akses app ke project tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kerangka ekstensi di [25 Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md). App Marketplace pihak ketiga adalah ekosistem, di luar cakupan.

### PERM-049 Automation Actor Permissions

- **Deskripsi:** Izin yang dipakai saat automation rule dijalankan.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / platform
- **Aktor & izin:** Admin Proyek, Admin Situs (permission Administer Jira), sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Actor default adalah akun sistem "Automation" (padanan "Automation for Jira").
  2. Actor dapat diganti ke user tertentu oleh pemilik rule yang berwenang.
  3. Actor harus memiliki izin project untuk setiap aksi rule (create, edit, transition, comment, dst.); evaluasi memakai Authorizer yang sama.
  4. Rule gagal pada aksi yang izinnya kurang, dan kegagalan tercatat di audit log rule dengan nama permission yang kurang.
  5. Pembuat rule: Admin Proyek untuk project rule; Admin Situs untuk global/multi-project rule.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Prasyarat AUTO-008 Automation permissions ([20 Otomasi](20-otomasi.md)) sehingga prioritas disamakan ke P0; semantik izin actor dimiliki modul 03, sedangkan pengaturan actor pada rule dimiliki modul 20. CMT-001 Add Comment ([08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md)) juga merujuk fitur ini untuk komentar yang dibuat actor automation/API.

### PERM-050 AI Agent Permissions

- **Deskripsi:** Pengaturan akses AI agent di project: agent dengan akun sendiri membutuhkan role/grant eksplisit, agent atas nama user mewarisi izin user.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek, Admin Situs (permission Administer Jira), AI agent.
- **Aturan perilaku (acceptance criteria):**
  1. Administer Projects diperlukan untuk menambah, mengonfigurasi, dan menghapus agent di project.
  2. Agent dengan akun sendiri harus diberi project role atau grant di scheme sebelum dapat membaca/mengubah work item.
  3. Agent yang bertindak atas nama user mewarisi izin user tersebut dan tidak dapat melebihinya.
  4. Grant ke Any logged in user otomatis berlaku untuk semua agent berakun sendiri.
  5. Agent dapat ditambahkan ke team-managed project Private.
- **Team-managed vs company-managed:** Team-managed: agent ditambahkan lewat People dengan role.
- **Catatan & ketergantungan:** ⚠ Fitur baru di produk acuan; detail perilaku perlu dipastikan sebelum implementasi. Fitur AI bersifat opsional (Fase 4).

## Izin objek & siklus hidup

### PERM-051 Board & Filter Permissions

- **Deskripsi:** Izin terkait board (board admin, filter sumber) dan aturan visibilitas isi board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek, board admin, user.
- **Aturan perilaku (acceptance criteria):**
  1. Board admin (user atau group) dapat mengubah konfigurasi board; non-admin melihat pengaturan read-only.
  2. Visibilitas board mengikuti share permission saved filter-nya.
  3. Board location: project atau user; location menentukan tempat board tampil di navigasi.
  4. Isi board hanya menampilkan work item yang boleh dilihat user.
  5. Mengubah filter board memerlukan hak editor filter; tanpa itu tombol edit filter dinonaktifkan dengan penjelasan.
  6. Team-managed: board mengikuti akses project.
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** Board admin dan visibilitas di [10 Board](10-board-kolom-kartu.md); share filter di [13 Pencarian & Filter](13-pencarian-filter-ekspor.md).

### PERM-052 Plans Permissions

- **Deskripsi:** Izin pada Plans (perencanaan lintas project) di Premium/Enterprise.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** User, Admin Situs (permission Administer Jira), pemilik plan.
- **Aturan perilaku (acceptance criteria):**
  1. Plan access: Open (semua user berlisensi dapat melihat) atau Private (hanya yang ditambahkan).
  2. Peran plan: Viewer (lihat) dan Editor (ubah dan simpan perubahan).
  3. Work item di plan dibatasi Browse Projects masing-masing user; plan tidak pernah menampilkan work item yang tidak boleh dilihat.
  4. Menyimpan perubahan plan ke work item memerlukan izin work item terkait (Edit, Schedule, Create); perubahan yang tidak berizin gagal per item dan dilaporkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Detail model akses Plans belum dipastikan terhadap produk acuan. Plans di [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md). Global permission Plans user/administrator hanya DC (PERM-067).

### PERM-053 Project Archiving Permission Impact

- **Deskripsi:** Dampak arsip project terhadap akses dan izin.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Hanya Admin Situs yang dapat archive dan restore project.
  2. Project archived disembunyikan dari user dan pencarian.
  3. Work item di project archived read-only dan tidak dapat diakses user biasa; tautan langsung menampilkan pesan bahwa project diarsipkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga PRJ-032 Project Archive ([04 Proyek](04-proyek.md)) dan ADM-024 Project archiving (admin) ([22 Administrasi Sistem](22-administrasi-sistem.md)), keduanya Fase 3; semantik izin dimiliki modul 03, sedangkan alur UI arsip/restore dimiliki modul 04 dan 22.

### PERM-054 Deactivated Users & Deleted Groups in Grants

- **Deskripsi:** Perilaku grant ketika user dinonaktifkan atau group dihapus.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs (permission Administer Jira), sistem.
- **Aturan perilaku (acceptance criteria):**
  1. User nonaktif tetap terlihat di role/grant dengan penanda "Nonaktif" tetapi tidak dievaluasi sebagai aktif.
  2. Saat group dihapus, grant dan role membership yang merujuk group tersebut ikut dihapus (dengan peringatan jumlah grant terdampak sebelum penghapusan).
  3. Rename group aman karena semua referensi memakai `groupId`.
  4. Tersedia pembersihan grant yatim (grant yang merujuk entitas yang tidak ada) sebagai tugas admin atau job terjadwal, dengan laporan hasil.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penonaktifan user di [02 Manajemen Pengguna](02-pengguna-grup-organisasi.md).

## Diagnostik, UX sadar-akses & audit

### PERM-055 Permission Helper

- **Deskripsi:** Alat diagnostik admin yang menjelaskan mengapa user dapat atau tidak dapat melakukan aksi tertentu pada work item/project.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (permission Administer Jira), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Input: user, work item (atau project), dan permission yang ditanyakan.
  2. Output: hasil dapat/tidak beserta grant yang menyebabkannya (role, group, scheme, tipe grantee); bila tidak, daftar grant yang dapat ditambahkan untuk memberi akses.
  3. Penjelasan mencakup faktor product access, Browse Projects, issue security level, dan workflow property status saat ini, masing-masing dengan status lolos/gagal.
  4. Dapat diakses dari halaman permission scheme, Project settings > Permissions, dan menu admin helper.
  5. Hanya Admin Situs dan Admin Proyek (untuk project-nya) yang dapat memakai helper.
- **Team-managed vs company-managed:** Untuk team-managed, helper menjelaskan access level dan role.
- **Catatan & ketergantungan:** Memakai fungsi "explain" dari Authorizer (PERM-002) agar penjelasan selalu konsisten dengan evaluasi nyata.

### PERM-056 Access-aware UX Behaviors

- **Deskripsi:** Perilaku UI yang menyesuaikan dengan izin user agar penolakan dapat dipahami dan tidak membocorkan data.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** User.
- **Aturan perilaku (acceptance criteria):**
  1. Saat @mention user yang tidak memiliki akses ke work item, editor menampilkan peringatan bahwa orang tersebut tidak akan dapat melihatnya.
  2. Share work item ke user tanpa akses memicu undangan atau access request sesuai wewenang pengirim.
  3. Halaman "You don't have access" / work item tidak ditemukan ditampilkan secara seragam tanpa membocorkan summary atau keberadaan work item yang dilindungi.
  4. ⚠ User dapat mengajukan request access ke project/site dari halaman tanpa akses; cakupan (project vs site) ditetapkan bersama modul 02.
  5. Aksi tanpa izin disembunyikan atau dinonaktifkan (dengan tooltip alasan) berdasarkan hasil mypermissions; server tetap menolak bila dipanggil langsung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** SPA memakai endpoint PERM-058; access request di [02 Manajemen Pengguna](02-pengguna-grup-organisasi.md).

### PERM-057 Permission Change Auditing

- **Deskripsi:** Pencatatan perubahan izin di audit log.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Audit log mencatat perubahan scheme, grant, role, role membership, dan security level dengan aktor, waktu, objek, nilai sebelum dan sesudah.
  2. Audit log mencatat perubahan global permission.
  3. Audit log dapat difilter berdasarkan kategori (mis. "Permissions") dan rentang waktu.
- **Team-managed vs company-managed:** Perubahan access level dan role team-managed juga dicatat.
- **Catatan & ketergantungan:** UI audit log di [22 Administrasi Sistem](22-administrasi-sistem.md).

## REST API izin

### PERM-058 My Permissions & Permission Check REST API

- **Deskripsi:** Endpoint REST untuk memeriksa izin user saat ini atau secara bulk.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Integrator/developer API, app.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /rest/api/3/mypermissions` dengan parameter `permissions` wajib (daftar key); permintaan tanpa parameter ditolak 400.
  2. Konteks opsional `projectKey`, `projectId`, `issueKey`, atau `issueId`; hasil per key berisi `havePermission`.
  3. `GET /rest/api/3/permissions` mengembalikan daftar semua permission (global dan project, termasuk yang didefinisikan app).
  4. `POST /rest/api/3/permissions/check` melakukan bulk check global/project permission untuk user atau app (hanya admin yang dapat memeriksa user lain).
  5. `POST /rest/api/3/permissions/project` mengembalikan daftar project tempat user memiliki izin tertentu.
- **Team-managed vs company-managed:** Tidak ada perbedaan; API mengembalikan hasil setara untuk kedua tipe.
- **Catatan & ketergantungan:** Dipakai SPA sejak Fase 1 melalui versi internal; versi publik terdokumentasi di [23 REST API Publik](23-api-publik-webhook.md) dan route berada di bawah `/api/v1` dengan alias kompatibilitas.

### PERM-059 User Permission Search REST API

- **Deskripsi:** REST API untuk mencari user berdasarkan izin.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Integrator/developer API, app.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /user/assignable/search` mencari user assignable untuk project atau work item.
  2. `GET /user/assignable/multiProjectSearch` mencari user assignable di beberapa project sekaligus.
  3. `GET /user/viewissue/search` mencari user yang dapat melihat work item tertentu.
  4. `GET /user/permission/search` mencari user yang memiliki permission tertentu pada konteks project/work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hasil tunduk pada PERM-008 (Browse users and groups) dan privasi profil. Lihat juga USR-025 User search & picker REST APIs ([02](02-pengguna-grup-organisasi.md)) dan API-050 Users, Myself & Preferences API ([23](23-api-publik-webhook.md)); semantik filter izin pada hasil pencarian dimiliki modul 03, sedangkan kontrak endpoint dan konvensi respons dimiliki modul 23.

### PERM-060 Permission Schemes REST API

- **Deskripsi:** REST API untuk pengelolaan permission scheme dan grant.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/developer API, Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. `GET/POST/PUT/DELETE /permissionscheme` untuk daftar, buat, ubah, dan hapus scheme (DELETE ditolak 400 bila scheme dipakai).
  2. `GET/POST/DELETE /permissionscheme/{id}/permission` dan `/permission/{permissionId}` untuk mengelola grant.
  3. `GET/PUT /project/{projectKeyOrId}/permissionscheme` untuk membaca dan mengganti scheme project.
  4. Parameter `expand` (`permissions`, `user`, `group`, `projectRole`, `field`, `all`) mengatur detail respons.
- **Team-managed vs company-managed:** Hanya company-managed; memanggil untuk team-managed project mengembalikan 400.
- **Catatan & ketergantungan:** Konvensi API di [23 REST API Publik](23-api-publik-webhook.md).

### PERM-061 Project Roles REST API

- **Deskripsi:** REST API untuk role global dan anggota role per project.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/developer API, Admin Situs (permission Administer Jira), Admin Proyek (untuk actor per project).
- **Aturan perilaku (acceptance criteria):**
  1. `GET/POST/PUT/DELETE /role` untuk role global.
  2. `GET /project/{projectIdOrKey}/role` dan `/role/{id}` untuk role dan anggotanya di project.
  3. `POST/DELETE` actor role per project (user, group, groupId).
  4. `GET /project/{projectIdOrKey}/roledetails` untuk ringkasan role project.
  5. ⚠ Endpoint default actors `/role/{id}/actors` disediakan hanya bila keputusan produk mempertahankan default members (lihat PERM-067); bila tidak, mengembalikan 410/404 terdokumentasi.
- **Team-managed vs company-managed:** Untuk team-managed, endpoint mengembalikan role lokal project.
- **Catatan & ketergantungan:** Konvensi API di [23 REST API Publik](23-api-publik-webhook.md).

### PERM-062 Issue Security REST API

- **Deskripsi:** REST API untuk issue security scheme dan level.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/developer API, Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. `GET/POST/PUT/DELETE /issuesecurityschemes`.
  2. Kelola level dan member lewat `/issuesecurityschemes/level` dan `/level/member`.
  3. `GET /project/{projectKeyOrId}/securitylevel` mengembalikan level yang dapat dipakai user pemanggil.
  4. `GET /project/{projectKeyOrId}/issuesecuritylevelscheme` dan endpoint asosiasi scheme ke project.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Konvensi API di [23 REST API Publik](23-api-publik-webhook.md).

### PERM-063 Application Roles REST API

- **Deskripsi:** REST API untuk membaca application role (product access).
- **Prioritas / Fase / Tier:** P3 / Fase 3 / platform
- **Aktor & izin:** Integrator/developer API (dengan hak admin).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /applicationrole` mengembalikan semua application role.
  2. `GET /applicationrole/{key}` mengembalikan satu role; key tidak dikenal menghasilkan 404.
  3. Respons memuat default group serta jumlah user dan seat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sumber data PERM-004.

## Keamanan data enterprise

### PERM-064 Data Classification

- **Deskripsi:** Label klasifikasi data pada project dan work item (Enterprise / add-on keamanan data) yang menjadi dasar data security policy.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi (definisi), Admin Proyek (default project), user (per work item).
- **Aturan perilaku (acceptance criteria):**
  1. Level klasifikasi (nama, warna, deskripsi, urutan sensitivitas) didefinisikan di tingkat organisasi.
  2. Admin Proyek dapat menetapkan default klasifikasi per project; work item baru mewarisinya.
  3. Klasifikasi dapat diubah per work item oleh user berizin edit; menurunkan klasifikasi di bawah default project memerlukan konfirmasi.
  4. Klasifikasi dapat menjadi cakupan data security policy (PERM-065).
  5. Indikator klasifikasi (lozenge) tampil di issue view, kartu board, dan hasil pencarian.
  6. ⚠ Dokumentasi admin menyarankan klasifikasi untuk membatasi seluruh project; mekanisme pembatasan aksesnya (via policy, bukan permission scheme) ditetapkan sebelum implementasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga USR-041 Data classification ([02](02-pengguna-grup-organisasi.md)) dan ADM-045 Data classification ([22](22-administrasi-sistem.md)); semantik izin dan penegakan dimiliki modul 03, sedangkan UI definisi level dan administrasi ada di modul 02/22. Bagian dari kontrol keamanan data organisasi di [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md).

### PERM-065 Data Security Policies

- **Deskripsi:** Policy tingkat organisasi untuk membatasi ekspor, public link, anonymous access, dan app access terhadap data.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Rule "blokir anonymous access": grant Public dan share Public diabaikan pada cakupan policy.
  2. Rule "blokir public links": tautan publik (filter/dashboard/board publik) dinonaktifkan pada cakupan policy.
  3. Rule "blokir ekspor data dan unduhan attachment": ekspor CSV/Excel/PDF/XML dan unduhan attachment ditolak dengan pesan yang menyebut policy.
  4. Rule "blokir akses app": app tidak dapat membaca data pada cakupan policy (kecuali app yang dikecualikan).
  5. Cakupan policy ditetapkan berdasarkan klasifikasi, project, atau seluruh produk.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga USR-040 Data security policies ([02](02-pengguna-grup-organisasi.md)) dan ADM-046 Data security policies ([22](22-administrasi-sistem.md)); semantik izin dan penegakan dimiliki modul 03. Penegakan oleh Authorizer (PERM-002) dan modul ekspor ([13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md)).

### PERM-066 Data Redaction (Atlassian Guard)

- **Deskripsi:** Kemampuan meredaksi (menghapus permanen) data sensitif seperti credential atau PII di field dan history work item, secara manual, dari temuan deteksi data sensitif, atau otomatis lewat policy. Padanan kemampuan add-on keamanan produk acuan; melengkapi Issue Redaction API di modul API.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, admin deteksi data sensitif (padanan Guard Detect admin), User access admin.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin memilih potongan teks di field yang didukung (Summary, Description, Environment, Comment, custom field teks single line/paragraph) lalu menjalankan aksi "Redact"; dukungan untuk komentar worklog ditetapkan kemudian.
  2. Teks yang diredaksi diganti placeholder bar solid berlabel "redacted" di work item view.
  3. Redaksi bersifat permanen: nilai asli dihapus dari penyimpanan, bukan sekadar di-mask, termasuk dari History/changelog.
  4. Bila nilai sensitif muncul lebih dari sekali, admin dapat memilih instance mana yang diredaksi.
  5. Redaksi dapat dijalankan dari temuan/alert deteksi data sensitif, termasuk bulk redaction atas banyak temuan.
  6. Redaksi otomatis berbasis policy tersedia untuk tipe data sensitif yang terdeteksi.
  7. Alasan redaksi wajib diisi; aktivitas redaksi (siapa, kapan, work item, field) tercatat di audit log organisasi.
  8. ⚠ Permission: Organization admin atau admin deteksi data sensitif; kewenangan redaksi dapat didelegasikan ke peran lain.
  9. Admin yang meredaksi harus memiliki akses ke instance produk; bila project permission atau issue security menyembunyikan work item, excerpt tidak terlihat dan redaksi tidak dapat dilakukan.
  10. Proses berjalan asinkron dengan status `IN_PROGRESS`, `SUCCEEDED`, `FAILED` yang dapat dipantau.
  11. Edge case: redaksi gagal bila jumlah instance terlalu banyak; UI menyarankan fallback edit field manual atau clone lalu hapus work item asli.
  12. Edge case: email notifikasi yang sudah terkirim tidak ikut berubah; UI menampilkan keterangan ini sebelum konfirmasi.
  13. ⚠ Edge case: teks yang diredaksi tidak lagi ditemukan lewat JQL text search, pencarian, dan export (indeks pencarian diperbarui).
  14. Edge case: attachment tidak dapat diredaksi sebagian (hanya dihapus utuh).
  15. ⚠ Restore data redaksi hanya mungkin lewat API deteksi dalam jendela waktu terbatas (nilai 10 atau 30 hari ditetapkan kemudian); setelah jendela lewat data tidak dapat dipulihkan.
  16. Ketersediaan: memerlukan langganan add-on keamanan tingkat tertinggi, terlepas dari tier produk; tidak tersedia di lingkungan pemerintahan (government cloud).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat juga ISS-043 Data Redaction (Atlassian Guard) ([06](06-work-item-inti.md)); tier disamakan dengan ISS-043 (platform, karena ketersediaan mengikuti add-on keamanan, bukan tier produk; butir 16), semantik izin redaksi dimiliki modul 03 dan perilaku pada work item dimiliki modul 06. Sisi API di "Issue Redaction API" modul [23 REST API Publik & Webhook](23-api-publik-webhook.md). Nama fitur mengikuti inventaris; di UI produk ini tidak memakai merek pihak acuan.

## Rujukan Data Center (tidak dibangun)

### PERM-067 Data Center-only Permissions

- **Deskripsi:** Permission dan opsi yang hanya ada di edisi Data Center produk acuan, dicatat agar tidak dikira fitur Cloud. Tidak dibangun; tercantum sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs (permission Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Global permission "Jira System Administrators" terpisah dari "Jira Administrators" tidak dibangun; kewenangan sistem dibagi lewat PERM-003.
  2. Global permission "Jira Users" dan "Create shared objects" (legacy) tidak dibangun; digantikan product access (PERM-004) dan PERM-009.
  3. Global permission "Plans user" / "Plans administrator" tidak dibangun; akses Plans diatur PERM-052.
  4. Project permission "Start/Complete Sprints" dan "Edit Sprints" terpisah tidak dibangun; digantikan Manage Sprints (PERM-025).
  5. Project permission "Browse archive" tidak dibangun; arsip diatur PERM-035 dan PERM-053.
  6. Grantee "Reporter (show only projects with create permission)" tidak dibangun.
  7. Default members pada project role tidak dibangun.
  8. Model "Extended project administration" tidak dibangun; digantikan PERM-029.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bila importer dari Data Center menemukan konfigurasi ini, importer memetakannya ke padanan Cloud atau melaporkannya sebagai tidak didukung ([25 Impor/Ekspor](25-integrasi-impor-ekspor-ekstensi.md)).

## Pertanyaan terbuka

- Apakah instalasi self-hosted open-source produk ini menerapkan pembatasan plan-tier (Free/Standard/Premium/Enterprise) untuk fitur izin, atau semua fitur (issue security, custom role, archive) terbuka secara default?
- Permission key internal (mis. `BROWSE_PROJECTS`) dipertahankan demi kompatibilitas API, sedangkan nama role/group sistem memakai nama netral (`app-project-access`, `app-users-admin`, `<product>-users-<site>`) dengan alias impor ke nama produk acuan; apakah alias tersebut juga diterima di REST API publik atau hanya oleh importer? Keputusan diselaraskan dengan Pertanyaan terbuka modul [02](02-pengguna-grup-organisasi.md) tentang nama grup sistem.
- Apakah default permission scheme memakai role Member/Viewer (site baru) saja, atau juga menyediakan Developers/Users untuk kompatibilitas impor?
- Bagaimana perilaku persis proteksi lockout saat menghapus grant Administer Jira terakhir: ditolak keras atau cukup konfirmasi berlapis?
- Siapa yang boleh menghapus group filter subscription buatan user lain: hanya pembuat dan Admin Situs, atau semua pemegang Manage group filter subscriptions?
- Izin apa yang mensyaratkan pemadaman notifikasi email pada bulk operation?
- Role team-managed mana yang berwenang menyetel issue restriction, dan apakah restriction diwariskan ke child/sub-task?
- Apakah Guest Access, Data Classification, Data Security Policies, dan Data Redaction masuk roadmap produk ini atau hanya dicatat sebagai rujukan enterprise?
- Berapa target maksimal keterlambatan propagasi perubahan izin (invalidasi cache) yang dapat diterima: real time (≤ 5 detik) atau eventual (≤ 1 menit)?
- Bagaimana memetakan field `Role` pada entity `User` di repo saat ini (default `user`) ke group dan global permission saat migrasi pertama?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk izin lanjutan (mis. field-level security, permission per field/komentar berbasis app, pembatasan berbasis atribut, "permission reporting" pihak ketiga) — tidak dibangun sebagai fitur native; produk ini hanya menyediakan titik ekstensi app permission (PERM-048).
- **Field-level security native** — tidak ada di produk acuan Cloud; menyembunyikan field dicapai lewat screen/field configuration/layout ([05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md)).
- **Fitur Data Center-only (dc-only-or-deprecated)** — dicatat sebagai rujukan di PERM-067: Jira System Administrators terpisah, Jira Users/Create shared objects, Plans user/administrator, Start/Complete Sprints dan Edit Sprints terpisah, Browse archive, grantee "Reporter (show only projects with create permission)", default members project role, Extended project administration.
- **Layanan lintas produk milik vendor acuan** — administrasi organisasi terpusat lintas produk, add-on keamanan berbayar (deteksi data sensitif, kebijakan organisasi lintas produk), dan identitas akun global di luar produk ini; produk ini hanya membangun padanan minimal yang tercantum di PERM-003, PERM-004, dan PERM-064 s.d. PERM-066.
- **Peran khusus Product Discovery** (Creator/Contributor dan sejenisnya) — dibahas di [37 JPD Peran](37-jpd-insights-delivery-peran-integrasi.md), bukan di modul ini.
- **Izin Assets/CMDB** — dibahas di modul [33 JSM Assets](33-jsm-assets-cmdb.md).
