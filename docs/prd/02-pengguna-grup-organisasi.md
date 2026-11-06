# 02 — Manajemen Pengguna, Grup & Organisasi

Modul ini mendefinisikan bagaimana produk ini mengelola **siapa** yang boleh memakai sistem: entitas Organisasi, site dan app (Jira terpadu, Service Management, Product Discovery), direktori pengguna beserta siklus hidupnya (undang, aktif, suspend, nonaktif, hapus), grup dan grup akses default, peran per app dan peran administrator, global permission, akun non-manusia (service account dan app user), identitas terkelola (verifikasi domain, managed account, SSO SAML, provisioning SCIM), kebijakan keamanan dan perlindungan data level organisasi, serta audit log dan insight pemakaian. Autentikasi (login, sandi, 2SV, sesi) dibahas di modul 01, sedangkan izin per proyek (permission scheme, project role, issue security) dibahas di modul 03. Modul ini menjadi sumber kebenaran untuk model "akun → akses app → grup → izin" yang dipakai semua modul lain. Konsol admin di produk ini disebut **Konsol Admin Organisasi** (padanan Atlassian Administration), dan paket fitur keamanan lanjutan disebut **paket Keamanan Lanjutan** (padanan Atlassian Guard); merek Atlassian hanya dipakai sebagai rujukan perilaku. Modul ini adalah **pemilik kanonik definisi tingkat organisasi**: akses app & role (USR-018), guest (USR-019), global permission (USR-022), data security policies (USR-040), data classification (USR-041), dan data residency (USR-042). Modul [03 Izin](03-izin-peran-keamanan.md) memegang penegakan izinnya (PERM-004, PERM-006, PERM-047, PERM-064, PERM-065), modul [22 Administrasi Sistem](22-administrasi-sistem.md) menjadi entry point admin di Settings (ADM-045, ADM-046, ADM-048), dan pasangan fitur tersebut wajib memakai prioritas, fase, dan tier yang sama dengan heading di modul ini.

**Modul terkait:**

- [00 Visi, Ruang Lingkup, Glosarium & Baseline](00-visi-ruang-lingkup-glosarium-baseline.md) — glosarium Organisasi, site, app, akun.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — login, sandi, 2SV, authentication policy, sesi, API token pribadi.
- [03 Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) — permission scheme, project role, issue security, permission helper.
- [04 Proyek](04-proyek.md) — akses proyek team-managed (Open/Limited/Private), project lead.
- [05 Konfigurasi Work Item: Tipe, Field, Layar & Skema](05-konfigurasi-work-item-skema.md) — field user/group picker, field classification.
- [06 Work Item Inti](06-work-item-inti.md) — assignee, reporter, perilaku user nonaktif di work item.
- [07 Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) — node @mention.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — visibilitas komentar ke grup, history work item, "Former user" di activity.
- [12 Bahasa Kueri (JQL-kompatibel)](12-jql-bahasa-kueri.md) — `membersOf()`, pencarian berdasarkan accountId.
- [13 Pencarian, Navigator, Filter Tersimpan & Ekspor](13-pencarian-filter-ekspor.md) dan [17 Dashboard & Gadget](17-dashboard-gadget.md) — kepemilikan dan sharing filter/dashboard.
- [18 Timeline, Rencana Lintas Proyek & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — entitas Teams dan profil tim (PLAN-057, PLAN-059).
- [19 Notifikasi In-app & Email](19-notifikasi-email.md) — email undangan, notifikasi access request.
- [20 Otomasi Berbasis Aturan](20-otomasi.md) — actor otomasi sebagai app user.
- [21 Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) — profil pengguna, visibilitas profil, entry point navigasi.
- [22 Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) — billing, sandbox, release tracks, pengaturan global.
- [23 REST API Publik & Webhook](23-api-publik-webhook.md) — konvensi REST API, paginasi, OAuth.
- [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — impor user/grup dari Jira lain, ekstensi dan app user, integrasi Slack/Teams.
- [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) — multi-tenant, enkripsi, retensi data.
- [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) — entitas User, Group, Organization dan event audit.
- [29 Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) — role Agent pada service project.
- [30 Service Management: Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — customer portal dan self-signup customer.
- [35 Service Management: Virtual Service Agent, AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md) — peran agent/customer dan admin JSM.
- [37 Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) — peran Creator/Contributor JPD.

## Persona terkait

- **Admin Situs** — dalam modul ini mencakup tiga varian peran: Admin Organisasi (kontrol penuh org), Admin Akses Pengguna (mengelola user dan grup untuk app tertentu), dan Admin Situs/App Admin (konfigurasi app). Persona utama modul ini.
- **Admin Proyek/Project Lead** — menambahkan orang dan grup ke proyek, mengundang guest ke satu space, menerima peringatan saat project lead nonaktif.
- **Developer/Anggota Tim** — memakai user picker, @mention, "Assign to me", mengundang rekan bila diizinkan.
- **Stakeholder/Viewer** — melihat profil orang dan direktori tim; mengakses proyek publik tanpa login.
- **Pengguna Tim Bisnis** — bergabung lewat approved domain atau invite link tanpa melibatkan admin.
- **Integrator/Developer API** — memakai REST API pencarian user/grup, API admin organisasi, service account, dan API data personal (GDPR).
- **Tamu/Kolaborator Eksternal** — role Guest yang dibatasi ke satu space, serta external user (akun di luar domain terverifikasi).
- **Agent Layanan** dan **Customer/Pelapor** — dirujuk sebagai role app Service Management (detail di modul 29, 30, 35).
- **Product Manager Discovery** dan **Kontributor/Stakeholder Discovery** — dirujuk sebagai role app Product Discovery (detail di modul 37).

## User story modul

- Sebagai Admin Situs, saya ingin mengundang banyak orang sekaligus lewat email dan langsung memilih role app serta grupnya agar tim baru bisa bekerja di hari pertama tanpa konfigurasi manual satu per satu (USR-009, USR-018).
- Sebagai Admin Situs, saya ingin melihat daftar semua pengguna dengan status, last active, dan akses app, lalu memfilter dan mengekspornya agar saya bisa mengaudit lisensi dan membersihkan akun yang tidak terpakai (USR-004, USR-006).
- Sebagai Admin Situs, saya ingin men-suspend atau menonaktifkan akun orang yang keluar dari perusahaan tanpa menghapus work item, komentar, dan riwayatnya agar jejak kerja tetap utuh dan seat berbayar berkurang (USR-007, USR-008).
- Sebagai Admin Situs, saya ingin membuat grup dan menandai grup akses default per app agar pemberian izin di permission scheme, project role, dan sharing filter cukup dikelola lewat keanggotaan grup (USR-015, USR-016, USR-017).
- Sebagai Admin Situs, saya ingin mendefinisikan approved domain dan aturan siapa yang boleh mengundang agar karyawan bisa bergabung sendiri dengan aman sementara orang luar tetap butuh persetujuan (USR-012, USR-013).
- Sebagai Admin Proyek/Project Lead, saya ingin menambahkan kolaborator eksternal sebagai Guest yang hanya melihat satu space agar vendor bisa berkolaborasi tanpa memakai seat berbayar dan tanpa melihat proyek lain (USR-019).
- Sebagai Developer/Anggota Tim, saya ingin user picker dan @mention menyarankan rekan satu proyek terlebih dahulu dan menyembunyikan akun nonaktif agar saya cepat memilih orang yang benar (USR-026, USR-008).
- Sebagai Developer/Anggota Tim, saya ingin saat me-mention orang yang belum punya akses muncul tawaran untuk mengundang atau meminta akses agar kolaborasi tidak terhenti (USR-009, USR-026).
- Sebagai Integrator/Developer API, saya ingin REST API pencarian user dan grup yang kompatibel dengan pola Jira Cloud agar integrasi dan app yang ada mudah diporting (USR-025).
- Sebagai Integrator/Developer API, saya ingin membuat service account dengan API token ber-scope dan masa berlaku agar integrasi otomatis tidak bergantung pada akun manusia (USR-028, USR-030).
- Sebagai Admin Situs, saya ingin memverifikasi domain perusahaan lalu menghubungkan IdP lewat SAML dan SCIM agar login terpusat dan akun karyawan yang keluar otomatis dinonaktifkan (USR-031, USR-032, USR-034, USR-035).
- Sebagai Admin Situs, saya ingin membatasi akses hanya dari rentang IP kantor dan memblokir ekspor atau akses anonim pada data rahasia agar memenuhi kebijakan keamanan perusahaan (USR-038, USR-040, USR-041).
- Sebagai Admin Situs, saya ingin audit log organisasi dan audit log konfigurasi yang bisa difilter dan diekspor agar setiap perubahan admin, grup, dan izin bisa ditelusuri saat audit (USR-045, USR-046).
- Sebagai Stakeholder/Viewer, saya ingin mencari orang dan tim dari navigasi lalu melihat work item yang di-assign ke mereka agar saya tahu siapa mengerjakan apa (USR-027).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
| --- | --- |
| P0 | 12 |
| P1 | 10 |
| P2 | 15 |
| P3 | 12 |
| **Total** | **49** |

| Fase | Jumlah fitur | Fitur |
| --- | --- | --- |
| Fase 1 | 12 | USR-004, 005, 007, 008, 009, 015, 016, 018, 021, 022, 025, 026 |
| Fase 2 | 2 | USR-017, 023 |
| Fase 3 | 16 | USR-001, 002, 003, 006, 010, 011, 012, 013, 020, 024, 027, 029, 030, 044, 045, 046 |
| Fase 4 | 18 | USR-014, 019, 028, 031–043, 047, 048 |
| Fase — | 1 | USR-049 (referensi Data Center) |

Catatan: inventaris berisi 50 fitur; "Invite People & Request Access (navigation entry)" digabung ke USR-009 sebagai alias, sehingga ada 49 heading.

Strategi urutan implementasi: Fase 1 membangun model data inti yang tidak boleh berubah belakangan, yaitu satu Organisasi dan satu site implisit per instance, akun dengan `accountId` stabil dan status siklus hidup (invited/active/suspended/deactivated/deleted), grup dengan `groupId` UUID, role per app, grup akses default, peran admin, dan global permission; tanpa ini modul 03 (permission scheme) dan semua picker tidak bisa dibangun. Bersamaan, user picker, REST API pencarian user/grup, dan perilaku "Former user" dibangun di Fase 1 karena dipakai oleh work item, komentar, board, dan JQL. Fase 2 menambah grup sistem terlindungi dan entry point dari pengaturan proyek. Fase 3 menambahkan alur onboarding mandiri (approved domain, invite link, access request, impor CSV), batas user/billing, akses anonim, API admin, app user untuk otomasi, serta audit log. Fase 4 menambahkan guest (bersama penegakannya di PERM-047) dan fitur enterprise yang bergantung pada verifikasi domain (managed account, SAML, SCIM, kebijakan keamanan, data residency, BYOK) dan di-feature-flag di balik paket Keamanan Lanjutan.

## Fitur

### Organisasi, site & model manajemen

### USR-001 Organization (Atlassian Administration)

**Deskripsi:** Organisasi adalah container administratif tertinggi di Konsol Admin Organisasi; di bawahnya ada site, app, user, domain, dan kebijakan keamanan. Produk ini memakai satu organisasi dan satu site per instance pada MVP, dengan model data yang sudah siap multi-site.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi (peran Organization admin). Pengguna lain tidak melihat Konsol Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Organisasi dibuat otomatis saat site pertama dibuat (pada instalasi pertama clone); admin organisasi dapat memberi nama dan me-rename organisasi; nama wajib diisi (1-255 karakter) dan perubahan nama tercatat di audit log.
2. Halaman Org overview menampilkan ringkasan app terpasang, jumlah user per status, status keamanan (2SV/SSO, domain terverifikasi), dan daftar rekomendasi tindakan.
3. Satu organisasi boleh memiliki banyak site, dan satu site boleh memiliki banyak app; satu site hanya dimiliki satu organisasi (constraint unik `site.org_id` non-null).
4. Pemindahan site/app antar organisasi hanya bisa dimulai oleh admin yang merupakan Organization admin di organisasi asal **dan** tujuan; selain itu tombol "Move" disembunyikan dan API mengembalikan 403.
5. ⚠ Sebelum pemindahan dijalankan, sistem menjalankan pemeriksaan validasi dan menampilkan daftar pemblokir: konflik authentication policy, perbedaan data residency, status paket Keamanan Lanjutan/langganan, dan status billing; pemindahan tidak bisa dilanjutkan selama ada pemblokir.
6. ⚠ Saat sebuah site dipindahkan, semua app di site tersebut ikut dipindahkan bersama; tidak ada pemindahan parsial sebagian app.
7. ⚠ Setelah pemindahan, user, grup, app access, dan managed account dipetakan ke organisasi tujuan; user tetap mempertahankan akses app yang dimilikinya sebelum pindah.
8. ⚠ Event pemindahan site/app tercatat di audit log organisasi asal dan organisasi tujuan dengan actor, waktu, dan objek yang dipindah.
9. ⚠ Opsi pemindahan tidak tersedia pada kondisi atau plan tertentu (mis. migrasi data residency sedang berjalan); UI menampilkan alasan ketidaktersediaan.
10. Admin yang mengelola lebih dari satu organisasi dapat berpindah organisasi lewat org switcher di header konsol admin; daftar hanya berisi organisasi tempat ia menjadi admin.
11. Navigasi Konsol Admin Organisasi terdiri dari: Directory (Users, Groups, Managed accounts, Service accounts), Apps, Security, Billing, Settings; menu yang tidak relevan dengan peran admin disembunyikan.
12. UI admin memakai istilah "apps" (bukan "products") secara konsisten di label, URL, dan respons API.
13. Setiap organisasi memiliki Org ID (UUID) yang tidak berubah dan dipakai sebagai path parameter di Organizations REST API (USR-030).
14. ⚠ Organisasi hanya bisa dihapus setelah semua site dan app di dalamnya dihapus atau dipindahkan; jika masih ada, tombol hapus non-aktif dengan pesan penjelasan.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Fase 1 sudah wajib memiliki baris `organizations` tunggal dan `sites` tunggal yang dibuat saat bootstrap (migrasi baru setelah `001_create_users_table`), agar seluruh entitas lain bisa diberi `org_id`/`site_id`; UI konsol organisasi penuh dan pemindahan site datang di Fase 3-4 (pemindahan site realistisnya P3). Rujukan: [28 Model Data Inti](28-model-data-inti.md), [27 NFR](27-nfr-arsitektur-platform.md) untuk multi-tenant.

### USR-002 User management experience (original vs centralized)

**Deskripsi:** Dua model manajemen user di Jira Cloud: original (user dan grup per site) dan centralized (user dan grup di level organisasi). Produk ini langsung memakai model centralized; model original hanya didokumentasikan sebagai rujukan perilaku legacy.

**Prioritas / Fase / Tier:** P3 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Model original (rujukan saja, tidak dibangun): user dan grup dikelola per site dengan role Site admin dan Trusted user; clone tidak menyediakan toggle ke model ini.
2. Model centralized adalah satu-satunya model clone: user dan grup dikelola di level organisasi, dengan peran User access admin dan App admin (USR-021).
3. Tidak ada migrasi model untuk clone; untuk data yang diimpor dari site Jira ber-model original (modul 25), role Site admin dipetakan ke Organization admin dan Trusted user dipetakan ke User access admin, dan pemetaan ini dilaporkan di laporan impor.
4. Kemampuan yang di Jira hanya ada pada model centralized (suspend access, OAuth 2.0 credential untuk service account, akses app per user) tersedia penuh di clone tanpa syarat model.
5. Menu manajemen user berada di Directory level organisasi; tautan lama gaya "Site > Users" (mis. dari dokumentasi atau bookmark) diarahkan (redirect 301) ke Directory > Users.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Keputusan arsitektur ini menyederhanakan USR-017 (grup `site-admins` dan `trusted-users` hanya relevan untuk impor). Rujukan impor: [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md).

### USR-003 Sites & apps administration

**Deskripsi:** Pengelolaan site (alamat instance) dan app yang terpasang (Jira terpadu, Service Management, Product Discovery), termasuk plan dan akses user per app.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi (semua site), Admin Situs/App Admin (detail app yang dikelolanya, tanpa billing).

**Aturan perilaku (acceptance criteria):**

1. Halaman Apps menampilkan daftar site dan app per site beserta plan masing-masing (Free, Standard, Premium, Enterprise); plan menentukan feature flag yang aktif (mis. IP allowlisting hanya Premium/Enterprise).
2. ⚠ Setiap site memiliki URL (subdomain/host); admin organisasi dapat me-rename URL site dengan batasan (format DNS valid, unik, batas frekuensi rename) dan URL lama di-redirect; dukungan custom domain penuh ditentukan di [27 NFR](27-nfr-arsitektur-platform.md).
3. Halaman detail app menampilkan jumlah user per role, daftar default group (USR-016), dan daftar app admin.
4. Admin organisasi dapat menambahkan app baru ke site (mis. mengaktifkan Service Management) dan memilih default akses untuk user baru (otomatis diberi akses atau tidak).
5. Setiap app memiliki tautan ke halaman billing/langganan app tersebut (detail di modul 22).
6. Sandbox dan release tracks untuk plan Premium/Enterprise ditampilkan sebagai tautan dari halaman app; perilakunya didefinisikan di [22 Administrasi Sistem](22-administrasi-sistem.md).
7. Aksi pemindahan site/app ke organisasi lain tersedia dari halaman site dan mengikuti aturan USR-001 butir 4-9.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Di MVP hanya ada satu site dengan app Jira terpadu; app Service Management (Fase 5) dan Product Discovery (Fase 6) diaktifkan sebagai "app" pada site yang sama, sehingga model app-per-site harus ada sejak Fase 1 walaupun halamannya Fase 3.

### Direktori & siklus hidup pengguna

### USR-004 User directory (Users list)

**Deskripsi:** Daftar semua user yang punya akses ke site/organisasi di Directory > Users, lengkap dengan pencarian, filter, ekspor, dan aksi massal.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna (hanya user dari app yang dikelolanya). Pengguna biasa tidak dapat membuka halaman ini (403).

**Aturan perilaku (acceptance criteria):**

1. Tabel menampilkan kolom: nama (dengan avatar), email, status (`active`, `invited`, `suspended`, `deactivated`), last active, app access (role per app), dan tipe akun (`managed`, `external`, `app`, `service`).
2. Kotak pencarian mencari berdasarkan nama atau email (case-insensitive, substring, minimal 1 karakter, debounce ±300 ms); hasil kosong menampilkan "Tidak ada pengguna yang cocok".
3. Filter yang bisa dikombinasikan: app access, role, status, last active (rentang), grup, managed vs external, app users, service accounts; filter aktif tampil sebagai chip yang bisa dihapus dan tersimpan di query string URL.
4. Tabel bisa diurutkan berdasarkan nama atau last active (naik/turun) dan dipaginasi server-side (default 50 per halaman, maks 100) agar tetap responsif untuk organisasi dengan ≥ 50.000 user.
5. Tombol Export menghasilkan CSV dengan pilihan kolom (groups, app access, last active, status); ekspor berjalan asinkron untuk > 10.000 baris dan hasilnya dikirim sebagai tautan unduhan; aksi ekspor tercatat di audit log.
6. Menu aksi per baris: Lihat detail (USR-005), Ubah akses, Resend invite (hanya status `invited`), Suspend/Remove (USR-007); aksi yang tidak berlaku untuk status/tipe akun disembunyikan.
7. ⚠ Pengguna dapat memilih beberapa baris (checkbox, termasuk "pilih semua hasil filter") lalu menjalankan bulk action: tambah ke grup, ubah app access, suspend/remove; hasilnya dilaporkan per user (berhasil/gagal dengan alasan).
8. Managed accounts ditampilkan di tampilan terpisah (Directory > Managed accounts, USR-032), tidak hanya sebagai filter.
9. User yang sudah diundang tetapi belum menerima undangan ditandai badge "Invited" beserta tanggal undangan.
10. Header halaman menampilkan jumlah total user dan jumlah billable user per app (aturan billable di USR-020).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Kondisi repo saat ini: tabel `users` hanya punya `is_active`, `status` (string bebas), `role` (string `user`), dan soft-delete `deleted_at`; perlu migrasi untuk enum status, `account_type`, `last_active_at`, relasi `user_app_access`, dan grup. Endpoint list user yang ada (`internal/features/user`) berubah menjadi endpoint admin yang dilindungi peran admin. Rujukan: [28 Model Data Inti](28-model-data-inti.md). Penahapan butir: butir 1-7, 9, 10 wajib di Fase 1; tampilan Managed accounts (butir 8) dan filter managed vs external/service accounts aktif setelah USR-031/USR-032 dan USR-028 tersedia (Fase 4).

### USR-005 User details page

**Deskripsi:** Halaman detail per user di konsol admin yang menampilkan profil, akses app, grup, status, peran admin, dan aksi administratif.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna (aksi terbatas pada app yang dikelolanya; peran admin hanya bisa diubah Admin Organisasi).

**Aturan perilaku (acceptance criteria):**

1. Header menampilkan avatar, nama, email, status, tipe akun (managed/external/app/service), dan tanggal user ditambahkan ke organisasi.
2. Seksi App access menampilkan role per app; admin dapat mengubah role atau mencabut akses app, dengan konfirmasi bila pencabutan menghapus satu-satunya sumber akses user.
3. Seksi Groups menampilkan daftar grup user; admin dapat menambahkan ke grup (picker grup) atau mengeluarkan dari grup; grup hasil sinkronisasi IdP ditampilkan read-only (USR-015 butir 9).
4. Menampilkan last active per app dan aktivitas terakhir keseluruhan; nilai "Never accessed" bila user belum pernah masuk.
5. Seksi Admin roles menampilkan peran admin user; Admin Organisasi dapat grant/revoke peran admin (USR-021) dengan dialog konfirmasi.
6. Untuk managed account, menampilkan authentication policy yang berlaku dan status 2SV (aktif/tidak); untuk akun lain seksi ini disembunyikan (detail kebijakan di modul 01).
7. Menu aksi: Resend invite, Suspend access, Restore access, Remove user, Deactivate account, Delete account, Reset password, Log out sessions; setiap aksi hanya muncul bila valid untuk status dan tipe akun (mis. Deactivate/Delete hanya untuk managed account), dan setiap aksi destruktif memerlukan konfirmasi.
8. Profil (nama, job title, email) hanya bisa diedit admin untuk managed account yang tidak disinkronkan dari IdP; untuk akun tersinkron muncul pesan "Atribut ini dikelola oleh identity provider".
9. Untuk managed account dengan paket Keamanan Lanjutan, menampilkan daftar API token milik user (nama, dibuat, kedaluwarsa, terakhir dipakai) dengan aksi revoke.
10. Tautan "Lihat aktivitas di audit log" membuka audit log organisasi (USR-045) yang sudah terfilter ke user ini sebagai actor atau target.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Reset password dan log out sessions memakai layanan di [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md). Endpoint `GET /api/v1/admin/users/{accountId}` menggantikan endpoint detail user yang sekarang. Penahapan butir: butir 6 (authentication policy & 2SV per managed account) dan butir 9 (API token user dengan paket Keamanan Lanjutan) baru tampil setelah USR-032 dan USR-037 tersedia (Fase 4); butir 10 aktif bersama USR-045 (Fase 3); seksi lain wajib di Fase 1. Selama Fase 1-3, aksi Deactivate/Delete pada butir 7 tersedia untuk semua akun lokal mengikuti catatan penahapan USR-007.

### USR-006 Last active & inactive user management

**Deskripsi:** Pencatatan waktu terakhir aktif per user dan per app untuk audit dan untuk mengurangi seat berbayar.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs.

**Aturan perilaku (acceptance criteria):**

1. Kolom Last active di daftar user menampilkan tanggal terakhir aktif, atau "Never accessed" bila user belum pernah memakai app mana pun.
2. Halaman detail user menampilkan last active terpisah untuk setiap app (Jira, Service Management, Product Discovery).
3. Filter "Tidak aktif lebih dari N hari" dengan pilihan cepat 30/60/90/180 hari dan nilai kustom; user "Never accessed" ikut terfilter bila tanggal undangan lebih lama dari N hari.
4. Konsol admin menampilkan rekomendasi "Cabut akses N user tidak aktif untuk menghemat N seat", dengan tombol yang membuka bulk action cabut akses pada daftar terfilter.
5. ⚠ Pada plan Enterprise atau dengan paket Keamanan Lanjutan, admin dapat mengaktifkan kebijakan pencabutan akses app otomatis untuk user yang tidak aktif melebihi ambang N hari; user menerima email peringatan sebelum pencabutan, dan setiap pencabutan tercatat di audit log.
6. Data last active tidak real-time: diperbarui berkala (agregasi dari event sesi/request, maks tertinggal 24 jam); UI menampilkan keterangan waktu pembaruan terakhir.
7. Kolom last active ikut disertakan saat ekspor CSV daftar user (USR-004 butir 5).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Implementasi: middleware auth mencatat aktivitas ke buffer dan job berkala menulis `last_active_at` per `(account_id, app)` agar tidak menulis DB setiap request. Rujukan: [27 NFR](27-nfr-arsitektur-platform.md).

### USR-007 Suspend / remove / deactivate users

**Deskripsi:** Berbagai cara mencabut akses user: suspend akses, remove dari organisasi, deactivate managed account, dan delete account, masing-masing dengan dampak berbeda.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna (suspend/restore/remove untuk app yang dikelolanya); deactivate/delete managed account hanya Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Suspend access: user tidak bisa mengakses app organisasi (login ke app ditolak dengan pesan "Akses Anda ke situs ini telah ditangguhkan"), tetapi tetap tercatat di directory, keanggotaan grup dipertahankan, dan tidak dihitung billable.
2. Restore access mengembalikan user yang di-suspend ke status aktif dengan app access dan grup seperti sebelum suspend; validasi batas user plan berlaku (USR-020).
3. Remove user dari organisasi: semua app access dan keanggotaan grup dihapus; akun tetap ada sebagai entitas global dan konten yang dibuatnya tetap ada.
4. Deactivate managed account: user tidak bisa login ke app mana pun di instance, semua sesi aktif dan refresh token dicabut segera (maks 60 detik untuk propagasi), dan API token tidak lagi diterima.
5. Reactivate managed account mengembalikan kemampuan login; akses app dikembalikan sesuai keadaan sebelum deactivate.
6. Delete managed account: akun masuk masa tenggang 14 hari (bisa dibatalkan, status "Pending deletion"); setelah 14 hari data personal (nama, email, avatar) dihapus permanen dan akun tampil sebagai "Former user" (USR-008, USR-044).
7. Setiap aksi menampilkan dialog konfirmasi yang menjelaskan dampaknya, minimal: konten tetap ada, user tidak bisa di-assign, filter/dashboard miliknya tetap ada, dan perubahan pada billing.
8. Proteksi lockout: admin tidak bisa men-suspend/deactivate/remove dirinya sendiri, dan Organization admin terakhir tidak bisa dicabut aksesnya; API mengembalikan 409 dengan kode galat `LAST_ORG_ADMIN` atau `SELF_ACTION_FORBIDDEN`.
9. Deprovision via SCIM (USR-035) menonaktifkan user secara otomatis dengan efek sama seperti butir 4 dan dicatat dengan actor "SCIM".
10. User yang dicabut aksesnya tidak dihitung pada siklus tagihan berikutnya (USR-020 butir 7).
11. Setiap aksi suspend, restore, remove, deactivate, reactivate, delete, dan pembatalan delete dicatat sebagai audit event dengan actor, target, dan waktu.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Kondisi repo: `is_active=false` dan `deleted_at` (soft delete GORM) saat ini dipakai sebagai satu-satunya mekanisme; perlu dipisah menjadi status organisasi (`suspended`/removed) dan status akun global (`deactivated`/`pending_deletion`/`deleted`). Soft delete GORM tidak boleh dipakai untuk "delete account" karena referensi `accountId` di work item harus tetap valid. Middleware `AuthMiddleware` wajib memeriksa status akun pada setiap request (bukan hanya saat login). Rujukan: [01 Identitas](01-identitas-autentikasi-sesi.md). Penahapan butir: suspend/restore/remove dan proteksi lockout wajib di Fase 1; deactivate/reactivate/delete untuk managed account (butir 4-6) di Fase 1 berlaku untuk semua akun lokal yang dibuat di instance (lewat undangan/registrasi), yang diperlakukan sebagai akun dikelola organisasi tunggal karena verifikasi domain belum ada; setelah USR-031/USR-032 (Fase 4) aksi ini dibatasi ke managed account dan akun external hanya bisa di-suspend/remove; deprovision SCIM (butir 9) mengikuti USR-035 (Fase 4).

### USR-008 Deactivated & deleted user handling in Jira

**Deskripsi:** Perilaku produk terhadap user yang dinonaktifkan, dicabut aksesnya, atau dihapus di work item, field, filter, board, dan JQL.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Admin Situs/App Admin (Administer Jira) untuk perbaikan kepemilikan; semua pengguna sebagai pengamat.

**Aturan perilaku (acceptance criteria):**

1. User nonaktif tetap tampil sebagai assignee/reporter dengan penanda "(Nonaktif)" dan avatar abu-abu; akun yang dihapus tampil sebagai "Former user".
2. User nonaktif atau tanpa akses app tidak muncul di user picker maupun assignee suggestions, dan API menolak assign ke work item baru dengan galat validasi "User tidak aktif atau tidak memiliki akses".
3. JQL tetap bisa mencari berdasarkan `accountId` user nonaktif (mis. `assignee = 5b10ac8d82e05b22cc7d4ef5`), dan autocomplete JQL menampilkan user nonaktif dengan penanda.
4. Field user picker yang berisi user nonaktif tetap menyimpan nilainya; saat field itu diedit, pengguna harus memilih user aktif (nilai lama boleh dihapus tetapi tidak bisa dipilih ulang).
5. Bila project lead atau default assignee proyek nonaktif, halaman pengaturan proyek menampilkan peringatan dan admin harus menggantinya; pembuatan work item dengan default assignee nonaktif menjadi Unassigned.
6. Filter dan dashboard milik user nonaktif tetap ada dan bisa diakses sesuai sharing-nya; Admin dapat menjalankan "Change owner" di Settings > System > Filters/Dashboards.
7. Filter subscription milik user nonaktif berhenti mengirim email sampai pemiliknya diganti atau diaktifkan kembali.
8. Automation rule dengan actor/owner nonaktif bisa gagal dijalankan; audit log rule mencatat galat "Actor tidak aktif" (detail di [20 Otomasi](20-otomasi.md)).
9. Watcher dan mention user nonaktif tetap tercatat pada work item, tetapi user tersebut tidak menerima notifikasi in-app maupun email.
10. Akun yang dihapus menampilkan "Former user" dan avatar default di history, komentar, dan activity stream; `accountId` tetap ada dan tetap konsisten di semua referensi.

**Team-managed vs company-managed:** Tidak ada perbedaan, kecuali bahwa di proyek team-managed peringatan project lead nonaktif ditampilkan di Project settings > Access.

**Catatan & ketergantungan:** Semua tabel yang menyimpan user wajib mereferensikan `account_id` (bukan email/username). Kondisi repo memakai `id` VARCHAR(36) sebagai kunci; ini dipertahankan sebagai `accountId`. Rujukan: [06 Work Item Inti](06-work-item-inti.md), [12 JQL](12-jql-bahasa-kueri.md), [13 Pencarian & Filter](13-pencarian-filter-ekspor.md), [19 Notifikasi](19-notifikasi-email.md).

### Undangan, bergabung mandiri & permintaan akses

### USR-009 Invite users

- **Alias inventaris:** Invite People & Request Access (navigation entry)

**Deskripsi:** Mengundang orang ke site/app lewat email sekaligus memilih role per app dan grup, baik dari konsol admin maupun dari entry point navigasi di dalam app (tombol "Invite people", prompt saat mention/assign, halaman permission denied).

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna, Admin Proyek (lewat "Add people"), dan pengguna biasa bila diizinkan User access settings (USR-012).

**Aturan perilaku (acceptance criteria):**

1. Dialog "Invite users" di konsol admin menerima satu atau banyak email yang dipisah koma, spasi, atau baris baru (maks 50 email per pengiriman); email tidak valid ditandai merah dan tombol kirim non-aktif sampai diperbaiki.
2. Pengundang memilih app dan role per app (Jira: User, Admin, Guest, atau tanpa akses; app lain sesuai modul 35/37); default diambil dari pengaturan default akses app (USR-003 butir 4).
3. Pengundang dapat langsung memasukkan user baru ke satu atau beberapa grup (picker grup multi-pilih).
4. Pesan personal opsional (maks 1.000 karakter, teks polos) disisipkan ke email undangan.
5. ⚠ Admin dapat mencentang opsi "Tambahkan tanpa mengirim email undangan"; user tetap berstatus `invited` sampai login pertama kali, dan opsi ini tidak tersedia bagi pengundang non-admin.
6. Dari dalam app, tombol "Invite people/Add people" tersedia di navigasi utama dan di halaman People/Access proyek; undangan dari proyek sekaligus menambahkan user ke role proyek yang dipilih.
7. Saat pengguna me-@mention atau meng-assign orang yang belum punya akses, sistem menampilkan prompt "Undang <email>?" (bila pengguna boleh mengundang) atau "Minta admin mengundang" (bila butuh approval).
8. User yang diundang berstatus "Invited" sampai menerima undangan; admin bisa Resend invite (maks 1 kali per 10 menit per user) dan Cancel invite (menghapus undangan beserta token-nya); token undangan kedaluwarsa setelah 30 hari.
9. Sebelum undangan dikirim, sistem memvalidasi batas user plan (mis. Free maks 10 user); bila jumlah undangan melampaui sisa kuota, pengiriman ditolak dengan pesan "Batas pengguna plan tercapai (X/10). Upgrade plan atau cabut akses pengguna lain."
10. Siapa yang boleh mengundang mengikuti User access settings (USR-012): hanya admin, pengguna dengan persetujuan admin, atau pengguna tanpa persetujuan; undangan non-admin yang butuh persetujuan masuk antrian Access requests (USR-013).
11. Undangan ke email yang sudah memiliki akun di instance langsung menautkan akun tersebut ke site/app (tidak membuat akun baru); user menerima email "Anda telah ditambahkan" alih-alih tautan pembuatan akun.
12. (Entry point navigasi) Tombol "Invite people/teammates" tersedia di sidebar, halaman People, dan alur onboarding pengguna baru.
13. (Entry point navigasi) Dialog invite dari navigasi menerima banyak email sekaligus, pilihan produk/role, dan pesan opsional dengan aturan validasi yang sama seperti butir 1-4.
14. (Entry point navigasi) Pengundang non-admin mengirim undangan langsung atau mengajukannya untuk persetujuan admin sesuai User access settings; UI menjelaskan mode yang berlaku sebelum tombol kirim ditekan.
15. ⚠ (Entry point navigasi) Bila approved domain dengan auto-join aktif atau invite link aktif, dialog invite menampilkan opsi tersebut (mis. "Siapa pun dengan email @perusahaan.com bisa bergabung" dan tombol salin invite link).
16. (Entry point navigasi) Dialog menampilkan peringatan batas user per plan saat sisa kuota ≤ 2 user atau habis.
17. (Entry point navigasi) Halaman permission denied (403 pada app atau tautan work item) menampilkan tombol "Request access" yang membuat access request ke admin (USR-013 butir 1).

**Team-managed vs company-managed:** Di proyek company-managed, "Add people" menambahkan user ke project role yang dipilih; di proyek team-managed, user ditambahkan ke role Administrator/Member/Viewer proyek (lihat [04 Proyek](04-proyek.md)).

**Catatan & ketergantungan:** Kondisi repo: hanya ada registrasi mandiri (`/api/v1/auth/register`); di produk ini pendaftaran terbuka harus dikontrol User access settings, dan undangan menjadi jalur utama. Email undangan memakai [19 Notifikasi In-app & Email](19-notifikasi-email.md); pembuatan sandi saat menerima undangan di [01 Identitas](01-identitas-autentikasi-sesi.md). Penahapan butir: butir 15 (approved domain/invite link di dialog) aktif setelah USR-011 dan USR-012 (Fase 3); butir 9 dan 16 memakai aturan batas USR-020 yang di Fase 1 cukup berupa batas konfigurasi sederhana.

### USR-010 Bulk import users (CSV)

**Deskripsi:** Menambahkan banyak user sekaligus dengan mengunggah file CSV di konsol admin.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs.

**Aturan perilaku (acceptance criteria):**

1. Konsol menyediakan template CSV yang bisa diunduh dengan kolom `email` (wajib), `display name`, `group` (beberapa grup dipisah titik koma); encoding UTF-8 dengan atau tanpa BOM.
2. Setelah upload, sistem menampilkan preview yang memisahkan baris valid dan tidak valid (email tidak valid, kolom wajib kosong, duplikat dalam file) sebelum impor dijalankan.
3. Admin memilih app access default yang diberikan ke semua user yang diimpor.
4. Admin memilih mengirim atau tidak mengirim email undangan ke user yang diimpor.
5. Setelah impor, laporan hasil menampilkan jumlah berhasil, duplikat (user sudah ada), dan gagal beserta alasan per baris; laporan bisa diunduh sebagai CSV.
6. ⚠ Satu impor dibatasi maksimum jumlah baris (usulan default 1.000 baris); file yang melebihi batas ditolak dengan pesan yang menyebut batasnya.
7. ⚠ Grup yang disebut di CSV tetapi belum ada dibuat otomatis bila admin mencentang "Buat grup yang belum ada"; bila tidak dicentang, baris tersebut tetap diimpor tanpa grup itu dan dilaporkan sebagai peringatan.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Impor berjalan sebagai job asinkron dan tunduk pada batas user plan (USR-020). Impor user dari Jira lain (bersama data proyek) ada di [25 Integrasi, Impor/Ekspor](25-integrasi-impor-ekspor-ekstensi.md).

### USR-011 Invite links

**Deskripsi:** Link undangan yang bisa dibagikan sehingga siapa pun yang memiliki link bisa bergabung ke app tanpa diundang satu per satu.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs (mengelola); pengguna akhir (menyalin/membagikan link bila fitur aktif).

**Aturan perilaku (acceptance criteria):**

1. Invite link bisa diaktifkan/dinonaktifkan per app dari konsol admin atau dari dialog Add people di app; menonaktifkan link langsung membuat link tidak berlaku.
2. ⚠ Setiap link memiliki masa berlaku dengan tanggal kedaluwarsa yang terlihat di UI (usulan pilihan 7, 30, 60, 90 hari; default 30 hari); link kedaluwarsa menampilkan halaman "Link undangan sudah tidak berlaku".
3. Admin dapat me-reset/regenerate link; link lama langsung tidak berlaku dan link baru memiliki token acak baru (≥ 128 bit).
4. Admin menentukan role default yang diberikan ke user yang bergabung lewat link (mis. User di Jira); role Admin tidak boleh dipilih sebagai role default.
5. Batas user plan tetap berlaku: saat kuota penuh, halaman join menampilkan "Situs ini sudah mencapai batas pengguna; hubungi admin" dan akun tidak ditambahkan.
6. ⚠ Admin dapat membatasi link hanya untuk email dari domain tertentu; email di luar domain ditolak dengan pesan yang jelas.
7. Audit event dicatat saat link dibuat, direset, dinonaktifkan, dan setiap kali dipakai untuk bergabung (dengan akun yang bergabung).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Halaman join memakai alur pendaftaran di [01 Identitas](01-identitas-autentikasi-sesi.md). Rate limit pemakaian link per IP didefinisikan di [27 NFR](27-nfr-arsitektur-platform.md).

### USR-012 User access settings (approved domains & self-signup)

**Deskripsi:** Pengaturan cara orang mendapatkan akses ke app: approved domains, self-signup, dan aturan undangan oleh pengguna non-admin.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs.

**Aturan perilaku (acceptance criteria):**

1. Admin mengelola daftar approved domains (domain email yang boleh bergabung); domain divalidasi formatnya dan tidak boleh duplikat.
2. Per domain, admin memilih "Grant access" (otomatis diberi akses) atau "Require admin approval" (masuk antrian USR-013).
3. Per domain, admin memilih app yang diberikan (mis. hanya Jira, atau Jira dan Product Discovery); role yang diberikan selalu role user standar app, bukan role admin.
4. User dengan email dari domain yang disetujui dapat langsung bergabung dari halaman login/landing site atau dari tautan app tanpa undangan.
5. ⚠ Aturan undangan oleh pengguna non-admin memiliki tiga opsi: "Siapa saja bisa mengundang", "Siapa saja bisa mengundang dengan persetujuan admin", "Hanya admin yang bisa mengundang"; default instalasi baru "dengan persetujuan admin".
6. Domain email publik (gmail.com, outlook.com, yahoo.com, dan daftar domain publik yang dipelihara sistem) tidak bisa ditambahkan sebagai approved domain; galat "Domain email publik tidak bisa dijadikan approved domain".
7. Pengaturan self-signup customer Service Management tidak ada di halaman ini; halaman menautkan ke pengaturan portal di [30 Help Center & Portal](30-jsm-help-center-portal-pelanggan-kanal.md).
8. Setiap perubahan pengaturan (tambah/hapus domain, ubah mode, ubah aturan undangan) tercatat di audit log dengan nilai lama dan baru.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Kondisi repo: endpoint register terbuka untuk siapa saja; setelah fitur ini ada, register hanya diterima bila email cocok approved domain, membawa token undangan/invite link, atau pada bootstrap admin pertama.

### USR-013 Access requests

**Deskripsi:** Permintaan akses ke app dari user yang belum punya akses, termasuk antrian persetujuan admin.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Pengguna akhir (mengajukan), Admin Organisasi, Admin Situs, Admin Akses Pengguna untuk app terkait (memutuskan).

**Aturan perilaku (acceptance criteria):**

1. User yang login tetapi tanpa akses app melihat halaman "Request access" saat membuka site atau tautan work item; ia bisa menambahkan alasan opsional (maks 500 karakter); request ganda untuk app yang sama ditolak dengan pesan "Permintaan Anda sedang menunggu persetujuan".
2. Undangan dari pengguna non-admin yang membutuhkan persetujuan (USR-012 butir 5) masuk ke antrian Access requests dengan pengundang tercatat.
3. Daftar Access requests di konsol admin menampilkan pemohon, app, waktu, dan pengundang atau alasan; bisa difilter per app.
4. Admin dapat Approve (memberi akses app dan memasukkan ke default group app, USR-016) atau Deny dengan alasan opsional.
5. Email notifikasi dikirim ke admin terkait saat ada request baru, dan ke pemohon (serta pengundang bila ada) saat keputusan dibuat.
6. ⚠ Email notifikasi ke admin memuat tombol Approve/Deny yang bekerja langsung dari email melalui tautan bertanda tangan sekali-pakai (kedaluwarsa 7 hari) dan tetap mewajibkan sesi admin yang valid.
7. Saat Approve, batas user plan divalidasi; bila penuh, approve ditolak dengan pesan batas plan dan request tetap di antrian.
8. Request yang sudah diputuskan hilang dari antrian; riwayat keputusan tersedia di audit log (USR-045).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Halaman "Request access" juga dipicu dari entry point navigasi (USR-009 butir 17). Notifikasi via [19 Notifikasi In-app & Email](19-notifikasi-email.md).

### USR-014 Product requests & discovered products (shadow IT controls)

**Deskripsi:** Kontrol di konsol admin yang menentukan apakah managed account boleh membuat site/app baru sendiri atau harus mengajukan request ke admin, ditambah daftar site/app yang dibuat managed account di luar organisasi (discovered products).

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi; pengguna dengan managed account (mengajukan).

**Aturan perilaku (acceptance criteria):**

1. ⚠ Pengaturan product requests memiliki tiga opsi: izinkan managed user membuat site/app sendiri, wajibkan request ke admin, atau blokir pembuatan.
2. Saat managed user mencoba membuat site baru dan pembuatan dibatasi, ia diarahkan ke formulir request ke Organization admin alih-alih alur pembuatan site.
3. Daftar product request tertunda tampil di konsol; admin dapat approve (site dibuat di dalam organisasi) atau decline dengan alasan.
4. Halaman Discovered products menampilkan daftar site yang dibuat managed account di luar organisasi (nama site, pembuat, tanggal dibuat, jumlah user).
5. ⚠ Admin dapat mengklaim site yang ditemukan ke dalam organisasi; klaim memerlukan konfirmasi dan mengikuti aturan pemindahan USR-001.
6. ⚠ Email notifikasi dikirim ke admin saat ada product request baru atau site baru terdeteksi.
7. ⚠ Fitur ini mensyaratkan domain terverifikasi (USR-031) dan managed account; halaman Discovered products hanya aktif dengan paket Keamanan Lanjutan.
8. Fitur ini terpisah dari Access requests (USR-013, akses ke app yang sudah ada) dan User access settings (USR-012, approved domains); UI menautkan ketiganya agar admin tidak bingung.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Untuk clone self-hosted satu instance, "discovered products" hanya bermakna bila produk dijalankan sebagai layanan multi-tenant (SaaS). Keputusan terbuka: apakah clone akan dioperasikan multi-tenant (lihat Pertanyaan terbuka).

### Grup

### USR-015 Groups management

**Deskripsi:** Membuat dan mengelola grup user di Directory > Groups. Grup dipakai untuk app access, permission scheme, project role, issue security, notifikasi, dan sharing filter/dashboard.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna (membuat dan mengelola anggota); Admin Situs/App Admin dengan Administer Jira hanya bisa memakai grup di konfigurasi, tidak mengelola anggota (USR-023 butir 6).

**Aturan perilaku (acceptance criteria):**

1. Create group: nama wajib, unik secara case-insensitive di organisasi (maks 255 karakter, tanpa karakter kontrol), deskripsi opsional; nama duplikat ditolak dengan galat "Grup dengan nama ini sudah ada".
2. ⚠ Rename group diperbolehkan untuk grup non-sistem; karena semua konfigurasi mereferensikan `groupId`, scheme, role, dan sharing tetap valid setelah rename, dan JQL yang memakai nama lama perlu diperbarui (sistem menampilkan peringatan jumlah filter terdampak).
3. Hapus grup memerlukan dialog konfirmasi yang menampilkan peringatan bila grup dipakai di scheme, role, filter, atau dashboard; bila ada komentar/worklog yang visibilitasnya dibatasi ke grup itu, admin wajib memilih grup pengganti (swap group) atau membiarkan visibilitas berubah menjadi terbatas untuk admin saja.
4. Admin dapat menambah banyak anggota sekaligus (picker multi-user, maks 100 per aksi) dan menghapus anggota.
5. Daftar grup menampilkan nama, jumlah anggota, dan penanda grup pemberi akses app; tersedia pencarian nama dan filter "grup yang memberi akses app".
6. Halaman detail grup menampilkan daftar anggota (paginasi), app access yang diberikan grup, dan admin access yang diberikan grup (bila ada).
7. Setiap grup memiliki `groupId` (UUID) yang tidak berubah dan nama; REST API menerima `groupId` sebagai identitas utama dan `groupname` untuk kompatibilitas (USR-025).
8. Grup bersarang (nested groups) tidak didukung: grup tidak bisa menjadi anggota grup lain.
9. Grup hasil sinkronisasi IdP (SCIM/Google Workspace) bersifat read-only: anggota tidak bisa ditambah/dihapus dan nama tidak bisa diubah dari UI; tombol terkait non-aktif dengan keterangan "Dikelola oleh identity provider".
10. Grup dapat dipakai di: global permissions, permission scheme, project roles, issue security levels, notification scheme, visibilitas komentar/worklog, sharing filter/dashboard, custom field group picker, dan fungsi JQL `membersOf()`; setiap pemakaian tercatat agar dialog hapus/rename bisa menghitung dampaknya.
11. ⚠ Admin dapat mengekspor anggota grup ke CSV (nama, email bila terlihat, status, tipe akun).
12. ⚠ Batas sistem: usulan maks 10.000 grup per organisasi dan maks 100.000 anggota per grup; melewati batas ditolak dengan galat yang menyebut batasnya.

**Team-managed vs company-managed:** Proyek company-managed memakai grup di permission scheme dan project role; proyek team-managed hanya bisa menambahkan grup ke role proyek (Administrator/Member/Viewer), bukan scheme.

**Catatan & ketergantungan:** Kondisi repo: belum ada tabel grup; perlu migrasi `groups` (`id` UUID, `org_id`, `name`, `name_lower` unik, `description`, `source` local/scim/google, `is_system`) dan `group_members`. Rujukan: [03 Izin, Peran Proyek & Keamanan](03-izin-peran-keamanan.md), [12 JQL](12-jql-bahasa-kueri.md), [28 Model Data Inti](28-model-data-inti.md). Penahapan butir: butir 9 (grup hasil sinkronisasi IdP) aktif bersama USR-035/USR-036 (Fase 4). Pembatasan visibilitas komentar/worklog ke grup dan penggantian grup saat hapus (butir 3) dirujuk di [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) (CMT-004).

### USR-016 Default access groups

**Deskripsi:** Grup default per app/role yang otomatis diisi user baru ketika user tersebut diberi akses ke app.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs.

**Aturan perilaku (acceptance criteria):**

1. Setiap app memiliki minimal satu default group untuk role user (mis. `jira-users-<site>`), dibuat otomatis saat app ditambahkan.
2. Setiap app memiliki default group untuk role admin (mis. `jira-admins-<site>`); anggotanya mendapat role Admin app tersebut.
3. Di halaman app access, admin dapat menandai atau menghapus tanda "default" pada grup pemberi akses.
4. User baru yang diberi akses app (lewat undangan, approved domain, invite link, access request, atau impor) otomatis dimasukkan ke semua default group app tersebut.
5. Mengeluarkan user dari grup pemberi akses mencabut akses app bila grup itu satu-satunya sumber akses user ke app tersebut; UI menampilkan peringatan sebelum aksi.
6. Grup pemberi akses app (product access group) ditandai berbeda dari grup biasa di daftar grup dan picker (badge "Akses app").
7. Bila sebuah app tidak memiliki default group untuk role user, konsol menampilkan peringatan "User baru tidak akan mendapat akses ke <app>".
8. Default group role user dipakai sebagai pemegang permission standar (mis. Browse projects, Create work items) di default permission scheme yang dibuat saat instalasi (lihat modul 03).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Akses app dihitung sebagai gabungan akses langsung per user (USR-018 butir 2) dan akses via grup. Bootstrap instalasi pertama membuat grup default dan menempatkan user pertama ke grup admin.

### USR-017 System & protected groups

**Deskripsi:** Grup bawaan yang memiliki perilaku khusus dan dilindungi agar tidak terhapus sembarangan.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs.

**Aturan perilaku (acceptance criteria):**

1. `site-admins`: di Jira model original memberi role site admin; di clone (model centralized, USR-002) grup ini hanya dibuat saat impor data dari site ber-model original dan dipetakan ke peran admin.
2. `jira-admins-<site>` (dan `jira-administrators` pada site lama) memberi global permission Administer Jira; grup ini terdaftar di halaman global permissions secara default.
3. `jira-users-<site>` memberi akses Jira; `jira-software-users-<site>` diperlakukan sebagai grup legacy yang tetap dikenali saat impor dan memberi akses Jira yang sama.
4. Grup akses app lain dibuat oleh modul masing-masing: `jira-servicemanagement-users-<site>` (Service Management, modul 35) dan `jira-product-discovery-users-<site>` (Product Discovery, modul 37).
5. `atlassian-addons-admin` (nama di clone dapat berupa `app-users-admin`, dengan alias impor) menampung app user dari integrasi/ekstensi (USR-029).
6. `trusted-users-<id>` untuk role trusted user (model original) hanya dikenali saat impor; tidak dibuat pada instalasi baru.
7. `administrators`/`system-administrators` (legacy) dari site lama dikenali saat impor dan dipetakan ke `jira-admins-<site>`.
8. Grup sistem tidak bisa dihapus atau di-rename; tombol hapus/rename non-aktif dan API mengembalikan 409 `PROTECTED_GROUP`; perubahan keanggotaan grup admin memunculkan peringatan keras dengan konfirmasi ketik nama grup.
9. App user yang menjadi anggota grup sistem tidak dihitung billable (USR-020 butir 3).
10. Grup sistem untuk sebuah app dibuat otomatis saat app tersebut ditambahkan ke site.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Penamaan sufiks `<site>` mengikuti nama site; untuk instance tunggal clone sufiks bisa berupa slug site. Tidak ada grup `org-admins`: Organization admin dan User access admin adalah peran (USR-021), bukan grup.

### Akses app, guest & lisensi

### USR-018 App access & roles (product access)

**Deskripsi:** Pemberian akses user ke app (Jira terpadu, Service Management, Product Discovery) dengan role per app. Role menentukan hak akses dan status billable user.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs, Admin Akses Pengguna.

**Aturan perilaku (acceptance criteria):**

1. Role Jira per app: User (billable), Admin (billable, mengelola konfigurasi app), Guest (tidak berbayar selama dalam kuota, hanya Standard ke atas, USR-019), atau tanpa akses; setiap user memiliki tepat satu role per app.
2. Akses bisa diberikan lewat grup pemberi akses (USR-016) atau langsung per user; akses efektif = role tertinggi dari semua sumber, dan UI detail user menampilkan sumber akses ("langsung" atau "via grup X").
3. Admin dapat mengubah role/akses untuk banyak user sekaligus dari daftar user (USR-004 butir 7).
4. Admin dapat mencabut akses app tanpa menghapus user dari organisasi; user tetap ada di directory dengan app access "Tidak ada".
5. App access menentukan siapa yang dihitung dalam batas user plan; perubahan role yang membuat user menjadi billable divalidasi terhadap batas (USR-020).
6. Pekerjaan bisnis (eks Jira Work Management) sudah terpadu dalam app Jira; tidak ada akses app terpisah untuk proyek bisnis.
7. Halaman Apps menampilkan jumlah user per role untuk setiap app.
8. Role app lain (Service Management: Agent/Customer; Product Discovery: Creator/Contributor) mengikuti model yang sama dan didefinisikan di [35 JSM Admin](35-jsm-ai-otomasi-api-admin-csm.md) dan [37 JPD Peran](37-jpd-insights-delivery-peran-integrasi.md).

**Team-managed vs company-managed:** Tidak ada perbedaan; akses app adalah prasyarat sebelum izin proyek dievaluasi.

**Catatan & ketergantungan:** Kondisi repo: kolom `users.role` (string, default `user`) dipakai sebagai role global; harus dimigrasi ke tabel `app_access(account_id, app_key, role, source)` dan kolom `role` lama dihapus setelah backfill (user dengan `role='admin'` menjadi Admin Jira + Organization admin pada bootstrap). Middleware wajib memeriksa akses app sebelum handler proyek dijalankan. Rujukan: [03 Izin](03-izin-peran-keamanan.md). Pasangan fitur: PERM-004 Product Access & Application Roles (modul 03) adalah sisi penegakan (evaluasi akses app sebelum permission proyek) dari definisi di sini; keduanya P0 / Fase 1 / platform dan aturan role serta sumber akses mengikuti fitur ini.

### USR-019 Guest access (Jira guests)

**Deskripsi:** Role Guest untuk kolaborator eksternal. Guest hanya bisa mengakses satu space/proyek dan tidak memakai seat berbayar selama masih dalam kuota.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi, Admin Situs (mengundang guest ke site), Admin Proyek (menambahkan guest ke space), Tamu/Kolaborator Eksternal.

**Aturan perilaku (acceptance criteria):**

1. Alur dua langkah: admin situs/organisasi mengundang guest ke site dengan role Guest, lalu Admin Proyek menambahkan guest ke satu space (proyek); guest tanpa space melihat halaman "Anda belum ditambahkan ke space mana pun".
2. Guest dibatasi ke satu space: pencarian, JQL, board, dashboard, user picker, dan REST API tidak pernah mengembalikan konten di luar space tersebut; percobaan akses langsung via URL menghasilkan 404.
3. Guest permissions adalah kumpulan permission terbatas yang bisa diberikan ke guest di space (mis. browse, create, comment, edit work item sendiri); permission administratif proyek dan global (Administer, Bulk change, Share filter publik) tidak pernah bisa diberikan ke guest.
4. Kuota guest gratis dihitung dari jumlah seat Jira berbayar di site (rasio dikonfigurasi per plan); seat agent Service Management, user app lain, dan customer portal tidak menambah kuota; guest di luar kuota ditolak atau dihitung billable sesuai kebijakan plan.
5. Role Guest hanya tersedia untuk plan Standard, Premium, dan Enterprise; pada plan Free opsi Guest disembunyikan di dialog invite.
6. Ekstensi terpasang (mekanisme ekstensibilitas modul 25; app Marketplace pihak ketiga di luar cakupan) yang mendeklarasikan dukungan guest dapat dipakai guest, terbatas pada data di space yang ditugaskan; ekstensi tanpa deklarasi tidak muncul untuk guest.
7. Mengonversi licensed user menjadi Guest tidak bisa dilakukan langsung: admin harus mencabut akses User terlebih dahulu (atau deactivate), lalu memberi role Guest; UI menjelaskan langkah ini.
8. Guest bisa diubah menjadi user penuh dengan memberinya role User; batas user plan divalidasi dan akses ke space lain mengikuti permission normal.
9. ⚠ Di picker dan mention, guest hanya terlihat oleh pengguna di space tempat guest ditugaskan; guest sendiri hanya melihat anggota space tersebut.

**Team-managed vs company-managed:** Guest dapat ditambahkan ke kedua jenis proyek; di company-managed, guest permissions dikonfigurasi lewat permission scheme khusus guest, di team-managed lewat role proyek yang dibatasi (detail di [03 Izin](03-izin-peran-keamanan.md)).

**Catatan & ketergantungan:** Istilah "space" untuk proyek mengikuti Jira 2025+; di produk ini UI memakai istilah proyek dengan alias space (lihat glosarium [00](00-visi-ruang-lingkup-glosarium-baseline.md)). Kepatuhan pembatasan guest wajib diuji di semua endpoint daftar/pencarian. Perilaku izin guest di tingkat proyek dirinci di PERM-047 Guest Access ([03 Izin](03-izin-peran-keamanan.md)); modul ini memegang alur undangan, kuota, dan role app. Tier dan fase diselaraskan dengan PERM-047 (P2 / Fase 4 / platform) karena penegakan pembatasan satu space bergantung pada permission scheme penuh (Fase 3) yang harus stabil lebih dulu; tier platform berarti role Guest adalah kapabilitas lintas app, dengan pembatasan plan Standard ke atas tetap berlaku sesuai butir 5.

### USR-020 User limits & billable users

**Deskripsi:** Aturan jumlah user maksimum per plan dan cara menghitung user yang ditagih.

**Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard

**Aktor & izin:** Admin Organisasi, Billing admin.

**Aturan perilaku (acceptance criteria):**

1. Plan Free: maksimum 10 user billable per app Jira; penambahan user ke-11 ditolak di semua jalur dengan pesan yang menyarankan upgrade.
2. ⚠ Plan Standard/Premium/Enterprise memiliki batas user per site yang dapat dikonfigurasi (usulan default 50.000, Enterprise hingga 100.000); nilai batas disimpan sebagai konfigurasi plan, bukan konstanta kode.
3. Billable = user dengan akses app aktif berrole User/Admin; tidak termasuk user suspended, deactivated, app user, service account (lihat USR-028 butir 7), guest dalam kuota, dan customer Service Management.
4. Konsol admin menampilkan peringatan dan banner saat pemakaian ≥ 90% batas dan saat batas tercapai.
5. Saat batas tercapai, invite, invite link, approve access request, restore access, impor CSV, dan SCIM provisioning yang menambah user billable diblokir dengan galat `USER_LIMIT_REACHED`; jalur lain (mis. menambah guest dalam kuota) tetap berjalan.
6. Langganan annual memakai user tier (mis. 1-10, 11-100, dst.); langganan monthly menyesuaikan jumlah tagihan otomatis sesuai jumlah billable user.
7. Mencabut akses user mengurangi tagihan pada siklus berikutnya (bukan prorata mundur).
8. Halaman billing menampilkan jumlah user billable per app; detail billing ada di [22 Administrasi Sistem](22-administrasi-sistem.md).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Untuk deployment self-hosted tanpa billing, batas user bisa dimatikan (unlimited) lewat konfigurasi lisensi; aturan penghitungan billable tetap dipakai untuk laporan. Pengecekan batas harus atomik (transaksi/lock) agar undangan paralel tidak melampaui batas.

### Peran admin & izin global

### USR-021 Organization & site admin roles

**Deskripsi:** Peran administratif di konsol admin: Organization admin, User access admin, Site admin, App admin, Trusted user, dan Billing admin.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Admin Organisasi (grant/revoke semua peran admin).

**Aturan perilaku (acceptance criteria):**

1. Organization admin memiliki kontrol penuh organisasi: keamanan, domain, billing, semua site/app, user, dan grup.
2. User access admin mengelola user dan grup untuk app tertentu (dipilih saat grant) tanpa akses ke halaman Security, Domains, dan Billing.
3. Site admin (rujukan model original) mengelola user, grup, dan semua app di satu site; di clone peran ini dipetakan ke User access admin untuk semua app di site tersebut.
4. App admin (mis. Admin Jira) mengelola konfigurasi app (scheme, workflow, field) tanpa mengelola user; setara dengan role Admin pada app access dan global permission Administer Jira.
5. Trusted user (rujukan model original) boleh mengundang user dan menambah app tanpa menjadi admin penuh; di clone hanya dikenali saat impor dan dipetakan ke izin undangan tanpa approval.
6. Billing admin mengelola tagihan dan langganan tanpa akses manajemen user.
7. Dialog Grant admin role dan Revoke admin role tersedia di detail user dan daftar administrator; grant memerlukan pemilihan cakupan (organisasi, site, atau app).
8. Halaman daftar administrator organisasi menampilkan semua pemegang peran admin dengan peran dan cakupannya.
9. Organisasi wajib memiliki minimal satu Organization admin aktif; revoke/suspend/deactivate admin terakhir ditolak dengan galat `LAST_ORG_ADMIN`.
10. Setiap perubahan peran admin tercatat di audit log organisasi dengan actor, target, peran, dan cakupan.
11. Untuk organisasi Enterprise multi-site, admin dapat didelegasikan per site atau per app (mis. User access admin hanya untuk site A).
12. ⚠ Penamaan peran mengikuti terminologi Jira terbaru (2025) dengan label yang dapat diubah di satu tempat (i18n) agar penyesuaian nama tidak memengaruhi model data.

**Team-managed vs company-managed:** Tidak ada perbedaan; peran admin proyek team-managed (Administrator) dibahas di [04 Proyek](04-proyek.md).

**Catatan & ketergantungan:** Kondisi repo: hanya ada `users.role`; peran admin harus menjadi tabel `admin_role_assignments(account_id, role, scope_type, scope_id)`. Bootstrap: user pertama yang mendaftar pada instalasi kosong menjadi Organization admin. Klaim JWT `sess` tidak boleh menyimpan peran admin secara permanen; peran dicek dari DB/cache agar revoke berlaku segera. Penahapan butir: Fase 1 cukup Organization admin, User access admin, App admin, dan Billing admin dengan cakupan organisasi/app; delegasi per site untuk organisasi multi-site (butir 11) di Fase 4.

### USR-022 Jira global permissions

**Deskripsi:** Permission level app (Settings > System > Global permissions) yang diberikan ke grup untuk kemampuan lintas proyek/space.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Admin Situs/App Admin dengan Administer Jira.

**Aturan perilaku (acceptance criteria):**

1. Administer Jira: mengonfigurasi app (scheme, workflow, field, pengaturan sistem); pemegangnya adalah App admin Jira.
2. Browse users and groups: melihat user/grup lain di picker, mention, profil, dan autocomplete JQL.
3. Share dashboards and filters: berbagi filter/dashboard ke grup, proyek, atau publik; tanpa permission ini filter hanya bisa privat.
4. Manage group filter subscriptions: membuat filter subscription yang dikirim ke grup.
5. Make bulk changes: menjalankan bulk edit/move/transition/delete; tanpa permission ini menu bulk change disembunyikan dan API mengembalikan 403.
6. Create team-managed projects (space): membuat proyek team-managed; tanpa permission ini opsi team-managed disembunyikan di wizard proyek.
7. Global permission tambahan untuk Service Management dan Product Discovery didefinisikan di modul 35 dan 37 dan tampil di halaman yang sama dengan pengelompokan per app.
8. Halaman daftar global permission menampilkan setiap permission dengan grup pemegangnya, dan admin dapat menambah/menghapus grup per permission.
9. Administer Jira tidak boleh kosong: menghapus grup terakhir dari permission ini ditolak dengan galat "Minimal satu grup harus memiliki Administer Jira".
10. Tanpa Browse users and groups, user picker/mention hanya menerima input accountId/email persis (tanpa daftar saran) dan REST user search mengembalikan 403.
11. ⚠ Pada plan Free, kustomisasi permission dibatasi (global permission tetap bisa diatur, tetapi permission scheme proyek tidak bisa dikustomisasi; lihat modul 03).

**Team-managed vs company-managed:** Global permission berlaku sama; khusus "Create team-managed projects" hanya memengaruhi pembuatan proyek team-managed.

**Catatan & ketergantungan:** Tabel `global_permission_grants(permission_key, group_id)`; evaluasi permission di-cache per user dengan invalidasi saat keanggotaan grup berubah. 'Create shared objects' dan 'Jira System Administrators' hanya di Data Center (lihat USR-049). Permission proyek di [03 Izin](03-izin-peran-keamanan.md). Pasangan fitur: PERM-006 Global Permissions (modul 03) adalah sisi penegakan (Authorizer) dari daftar global permission yang didefinisikan di sini; keduanya P0 / Fase 1 / free-standard, dan daftar permission serta aturan butir 9-10 mengikuti fitur ini.

### USR-023 Jira-side user management entry points

**Deskripsi:** Titik masuk manajemen user dari dalam app (Settings > User management, halaman People/Access proyek) yang terhubung ke Konsol Admin Organisasi.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Admin Situs/App Admin, Admin Proyek, pengguna akhir.

**Aturan perilaku (acceptance criteria):**

1. Menu Settings > User management di app mengarahkan ke Directory > Users di konsol admin (tab yang sama bila konsol berada dalam SPA yang sama).
2. Project settings > People (company-managed) memungkinkan Admin Proyek menambah user/grup ke project role dan menghapusnya.
3. Project settings > Access (team-managed) menampilkan access level Open/Limited/Private dan role Administrator/Member/Viewer (detail di modul 03/04).
4. Tombol "Add people" di proyek dapat sekaligus mengundang user baru ke site bila pengguna diizinkan mengundang (USR-009, USR-012).
5. App menyediakan halaman direktori people dan halaman profil user (rujukan USR-027 dan modul 21).
6. Admin Jira yang bukan Organization admin/User access admin tidak bisa mengelola user/grup; ia hanya bisa mengelola global permissions dan project roles, dan menu manajemen user menampilkan pesan "Minta admin organisasi Anda".
7. Permission helper untuk mendiagnosis mengapa user bisa/tidak bisa melakukan sesuatu tersedia dari halaman People proyek (detail di [03 Izin](03-izin-peran-keamanan.md)).
8. Halaman Settings > System > Filters/Dashboards menyediakan aksi Change owner untuk keperluan offboarding (USR-008 butir 6).

**Team-managed vs company-managed:** Company-managed memakai halaman People (project role); team-managed memakai halaman Access (access level + role sederhana).

**Catatan & ketergantungan:** Rute frontend: `/settings/user-management` (redirect), `/projects/{key}/settings/people`, `/projects/{key}/settings/access`. Permission helper (butir 7) dirinci di PERM-055 (modul 03).

### USR-024 Anonymous & public access

**Deskripsi:** Pengaturan akses tanpa login ke proyek lewat permission "Anyone on the web", beserta kontrol pemblokiran di level organisasi.

**Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard

**Aktor & izin:** Admin Situs/App Admin, Admin Organisasi, pengguna anonim.

**Aturan perilaku (acceptance criteria):**

1. Admin dapat memberikan "Browse projects" ke "Anyone on the web" di permission scheme company-managed; proyek dengan scheme ini bisa dilihat tanpa login (work item, board, komentar publik).
2. ⚠ Proyek team-managed tidak mendukung akses anonim; access level team-managed tidak memiliki opsi publik.
3. Saat memberikan akses anonim, UI menampilkan peringatan bahwa proyek publik bisa diindeks mesin pencari dan semua konten (termasuk lampiran) terlihat publik; admin wajib mencentang konfirmasi.
4. Organization admin dapat memblokir akses anonim untuk seluruh organisasi atau app tertentu lewat data security policy (USR-040); policy ini mengalahkan permission scheme.
5. Pengguna anonim tidak bisa berkomentar, membuat, atau mengubah work item kecuali permission terkait juga diberikan ke "Anyone on the web"; aksi yang tidak diizinkan mengarahkan ke halaman login.
6. Filter/dashboard bisa dibagikan publik hanya bila pemilik memiliki Share dashboards and filters dan proyek sumbernya mengizinkan akses anonim; konten dari proyek non-publik tidak tampil bagi pengguna anonim.
7. Mode Public/Private instance dengan self-signup hanya ada di Data Center dan tidak dibangun (lihat USR-049).

**Team-managed vs company-managed:** Hanya company-managed yang mendukung akses anonim (butir 2).

**Catatan & ketergantungan:** Middleware perlu mode "optional auth" untuk rute yang mengizinkan anonim; rate limit anonim di [27 NFR](27-nfr-arsitektur-platform.md). Permission scheme di [03 Izin](03-izin-peran-keamanan.md). Perilaku permission "Anyone on the web" di tingkat scheme dirinci di PERM-046 (modul 03); modul ini memegang kontrol level organisasi dan peringatan admin.

### Pencarian, pemilihan & direktori orang

### USR-025 User search & picker REST APIs

**Deskripsi:** REST API untuk pencarian dan pengelolaan user/grup yang dipakai user picker, assignee, mention, dan integrasi. Path dan bentuk respons kompatibel dengan pola Jira Cloud REST API v3 agar integrasi mudah diporting.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Integrator/Developer API, frontend produk ini (atas nama pengguna akhir). Semua endpoint pencarian mewajibkan Browse users and groups; endpoint tulis grup mewajibkan peran pengelola user/grup.

**Aturan perilaku (acceptance criteria):**

1. `GET /rest/api/3/user/search?query=` mengembalikan user yang cocok dengan nama tampilan atau email (sesuai butir 15).
2. `GET /rest/api/3/users/search` mengembalikan semua user (aktif dan nonaktif, ditandai `active`) secara terpaginasi.
3. `GET /rest/api/3/user/assignable/search` (parameter `issueKey` atau `project`) dan `/user/assignable/multiProjectSearch` (parameter `projectKeys`) hanya mengembalikan user aktif yang memiliki permission Assignable user.
4. `GET /rest/api/3/user/viewissue/search` mengembalikan user yang bisa melihat work item tertentu, dan `/user/permission/search` mengembalikan user yang memiliki permission tertentu (parameter `permissions`).
5. `GET /rest/api/3/user/picker` mengembalikan hasil untuk picker dengan `html` ber-highlight bagian yang cocok (`<strong>`) dan total hasil.
6. `GET /rest/api/3/user?accountId=` mengembalikan satu user (404 bila tidak ada), dan `GET /rest/api/3/user/bulk?accountId=...` mengembalikan banyak user (maks 200 accountId).
7. `GET /rest/api/3/user/groups?accountId=` mengembalikan grup milik user.
8. `GET /rest/api/3/groups/picker` (pencarian grup), `/group/bulk` (daftar grup terpaginasi, filter `groupId`/`groupName`), dan `/group/member` (anggota grup terpaginasi, opsi `includeInactiveUsers`).
9. `GET /rest/api/3/groupuserpicker` mengembalikan gabungan hasil user dan grup dalam satu respons (`users`, `groups`).
10. `POST /rest/api/3/group` membuat grup dan `DELETE /rest/api/3/group` menghapus grup (parameter `groupId`/`groupname`, opsional `swapGroup`/`swapGroupId` untuk memindahkan visibilitas komentar/worklog, USR-015 butir 3).
11. `POST /rest/api/3/group/user` menambah anggota dan `DELETE /rest/api/3/group/user` menghapus anggota; grup hasil sinkronisasi IdP mengembalikan 400 "Group is managed by an identity provider".
12. User properties (`/rest/api/3/user/properties/{propertyKey}`) memungkinkan app/integrasi menyimpan data JSON per user (maks 32 KB per properti); user hanya bisa menulis propertinya sendiri kecuali admin/app.
13. Semua endpoint pencarian mewajibkan Browse users and groups (403 bila tidak), dan atribut seperti email difilter sesuai profile visibility (modul 21).
14. Paginasi memakai `startAt`/`maxResults` dengan batas maksimum hasil (default 50, maks 1.000 untuk `/users/search`, maks 100 untuk endpoint lain); nilai di atas batas dipotong ke batas.
15. Pencarian mencocokkan prefix kata nama tampilan (case-insensitive, tanpa memperhatikan aksen), serta email hanya jika email user tersebut terlihat oleh pemanggil.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Konvensi umum REST API publik (autentikasi, galat, versi) di [23 REST API Publik & Webhook](23-api-publik-webhook.md); API internal SPA tetap di `/api/v1`, sedangkan path kompatibel `/rest/api/3` disediakan sebagai lapisan publik. Kondisi repo: hanya ada `/api/v1/users` (list/detail) tanpa kontrol permission granular. Swagger wajib diperbarui (`make swag`).

### USR-026 User picker & mention UI behavior

**Deskripsi:** Perilaku pemilihan user di UI: field user picker, assignee, @mention, dan sharing.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna berlisensi; hasil tergantung Browse users and groups dan profile visibility.

**Aturan perilaku (acceptance criteria):**

1. Autocomplete muncul setelah 1 karakter (debounce ±200 ms), mencocokkan nama atau email, dan menampilkan avatar, nama, serta email (bila terlihat); navigasi keyboard (↑/↓/Enter/Esc) didukung.
2. Hasil diprioritaskan: pengguna yang baru dipilih (recent users) dan anggota proyek saat ini lebih dulu, kemudian pengguna lain.
3. Picker assignee menyediakan opsi "Assign to me" (pintasan keyboard `i` di tampilan work item) dan "Unassigned" (bila proyek mengizinkan unassigned).
4. Mengetik `@` di editor membuka saran mention; memilih user menyisipkan node mention (berisi accountId) dan memicu notifikasi ke user tersebut saat konten disimpan (modul 07, 19).
5. Mention atau assign user yang tidak punya akses ke proyek/app memunculkan prompt untuk mengundang atau meminta akses (USR-009 butir 7).
6. User nonaktif, suspended, dan akun yang dihapus disembunyikan dari picker (USR-008 butir 2).
7. Bila tidak ada hasil, picker menampilkan "No matching users" (dilokalisasi: "Tidak ada pengguna yang cocok"); hasil dan atribut yang tampil mengikuti profile visibility.
8. Group picker (untuk custom field group, sharing filter/dashboard, dan konfigurasi permission) memakai pola autocomplete yang sama berdasarkan nama grup.
9. Saran assignee hanya berisi user dengan permission Assignable user di proyek work item tersebut.

**Team-managed vs company-managed:** Tidak ada perbedaan; sumber "anggota proyek" adalah project role (company-managed) atau role Access (team-managed).

**Catatan & ketergantungan:** Komponen Svelte tunggal `UserPicker`/`GroupPicker` dipakai ulang di seluruh frontend. Rujukan: [07 Editor Teks Kaya](07-editor-konten-kaya.md), [06 Work Item Inti](06-work-item-inti.md), [19 Notifikasi](19-notifikasi-email.md).

### USR-027 People & Teams Directory (navigation entry)

**Deskripsi:** Entry point untuk mencari orang dan tim dari navigasi, menuju halaman profil orang dan profil tim.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Semua pengguna berlisensi (Browse users and groups); guest dibatasi sesuai USR-019.

**Aturan perilaku (acceptance criteria):**

1. Direktori bisa diakses dari section Teams/People di sidebar dan dari hasil kategori People/Teams pada pencarian global.
2. ⚠ Halaman direktori menyediakan pencarian orang dan tim berdasarkan nama/email, serta daftar "Your teams" dan "People you work with" (orang yang sering berinteraksi di work item yang sama dalam 90 hari terakhir).
3. Tombol "Create team"/"Start a team" membuat tim baru dan mengundang anggota (model Teams di modul 18).
4. Hasil dibatasi visibilitas profil dan pengaturan admin tentang daftar user (Browse users and groups); tanpa izin, direktori menampilkan pesan bahwa daftar orang tidak tersedia.
5. Profil orang/tim menampilkan tautan ke work item yang di-assign ke orang/tim tersebut (menjalankan JQL `assignee = <accountId>` atau `team = <teamId>`).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Profil personal dan preferensi di [21 Navigasi Global](21-navigasi-ux-preferensi.md); entitas Teams dan Team field di [18 Timeline, Rencana & Tampilan](18-timeline-rencana-tampilan-proyek.md). Pencarian orang memakai endpoint USR-025. Penahapan butir: di Fase 3 direktori hanya berisi orang (butir 1, 2 bagian orang, 4, 5 bagian orang); bagian tim (butir 3, daftar "Your teams", `team = <teamId>`) aktif setelah entitas Teams PLAN-057 dan profil tim PLAN-059 tersedia (Fase 4). Halaman profil orang dirinci di NAV-027.

### Akun non-manusia & API admin

### USR-028 Service accounts

**Deskripsi:** Akun non-manusia yang dikelola organisasi di Directory > Service accounts untuk integrasi otomatis, dengan kredensial API sendiri.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi (membuat dan mengelola); Integrator/Developer API (memakai kredensial).

**Aturan perilaku (acceptance criteria):**

1. Admin membuat service account di Directory > Service accounts dengan nama (wajib, unik) dan deskripsi; sistem membuat akun dengan `accountType = service` dan email sintetis non-routable.
2. Service account bisa diberi app access, keanggotaan grup, dan project role seperti user biasa; semua permission tetap dievaluasi normal (tidak ada bypass).
3. Create credentials > API token: nama, masa berlaku 1-365 hari, dan scopes wajib dipilih (mis. `read:jira-work`, `write:jira-work`); token hanya ditampilkan sekali saat dibuat.
4. Create credentials > OAuth 2.0 (client credentials / 2LO) dengan scopes: menghasilkan client ID dan client secret (secret ditampilkan sekali); access token berlaku ±3600 detik dan tidak ada refresh token.
5. Seksi Credentials menampilkan semua kredensial (nama, tipe, dibuat, kedaluwarsa, terakhir dipakai) dan aksi Revoke yang langsung membuat kredensial ditolak (401).
6. Service account tidak bisa login interaktif ke UI; percobaan login menghasilkan galat dan tidak ada alur reset sandi.
7. ⚠ Status billable service account dikonfigurasi per plan (usulan: tidak billable hingga kuota tertentu per organisasi, misal 5, di atasnya billable); batas jumlah service account per organisasi dikonfigurasi.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** OAuth 2.0 server dan validasi scopes di [01 Identitas](01-identitas-autentikasi-sesi.md) dan [23 REST API Publik](23-api-publik-webhook.md). Aksi oleh service account tercatat di audit log dengan tipe actor `service`.

### USR-029 App users (bot accounts)

**Deskripsi:** Akun bertipe "app" yang dibuat otomatis saat ekstensi atau integrasi dipasang, lalu dipakai sebagai actor aksi app tersebut.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Sistem (membuat otomatis), Admin Organisasi (melihat, membatasi izin).

**Aturan perilaku (acceptance criteria):**

1. App user memiliki `accountType = app` dan tampil dengan ikon bot di directory, history, komentar, dan picker.
2. App user tidak dihitung billable.
3. App user otomatis dimasukkan ke grup `atlassian-addons-admin` (padanan clone, USR-017 butir 5) atau grup akses app yang dideklarasikan ekstensi.
4. Directory menyediakan filter "App users" (USR-004 butir 3).
5. App user tidak bisa login; tidak bisa dihapus manual selama app/ekstensi masih terpasang (aksi hapus non-aktif dengan keterangan); saat ekstensi dicopot, app user dinonaktifkan otomatis.
6. Otomasi memakai actor app user bawaan (padanan "Automation for Jira") untuk aksi yang dijalankan rule, sehingga history menampilkan actor otomasi.
7. Permission app user bisa dibatasi lewat project role/permission scheme seperti user lain.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Actor otomasi di [20 Otomasi](20-otomasi.md); ekstensibilitas di [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md). Marketplace pihak ketiga di luar cakupan; app user tetap dibutuhkan untuk integrasi native (GitHub/GitLab, Slack) dan otomasi.

### USR-030 Organization API keys & admin REST APIs

**Deskripsi:** API key level organisasi untuk memanggil REST API konsol admin: Organizations, User management, dan Audit log/events.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi (membuat/mencabut), Integrator/Developer API (memakai).

**Aturan perilaku (acceptance criteria):**

1. Settings > API keys: admin membuat API key dengan nama (wajib, unik di organisasi).
2. ⚠ Setiap key wajib memiliki tanggal kedaluwarsa (usulan maks 1 tahun dari pembuatan); key kedaluwarsa ditolak dengan 401 dan email pengingat dikirim 7 hari sebelumnya.
3. ⚠ Key memiliki scope (mis. baca saja vs baca-tulis; per kelompok endpoint users/domains/policies/events); panggilan di luar scope ditolak 403.
4. Organization ID dan API key ditampilkan sekali saat pembuatan; setelah dialog ditutup, key tidak bisa dilihat lagi (disimpan sebagai hash).
5. Admin dapat me-revoke key; revoke berlaku segera.
6. Daftar key menampilkan nama, pembuat, tanggal dibuat, dan tanggal kedaluwarsa (serta terakhir dipakai).
7. Organizations API menyediakan endpoint users, domains, policies, dan events (audit log) dengan path `/admin/v1/orgs/{orgId}/...`.
8. User management API menyediakan lifecycle akun (disable/enable/delete), profil managed account, dan pengelolaan API token user (`/users/{accountId}/manage/...`).
9. ⚠ Endpoint directory menyediakan pengelolaan grup dan anggota di level organisasi (`/admin/v1/orgs/{orgId}/directory/...`), terpisah dari API grup app di USR-025.
10. Audit event dicatat saat key dibuat dan dicabut.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Autentikasi `Authorization: Bearer <api-key>`; rate limit per key di [27 NFR](27-nfr-arsitektur-platform.md). Konvensi API di [23 REST API Publik](23-api-publik-webhook.md).

### Identitas terkelola, SSO & provisioning

### USR-031 Domain verification

**Deskripsi:** Verifikasi kepemilikan domain email oleh organisasi agar akun dengan domain tersebut bisa diklaim menjadi managed account.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Admin menambahkan domain dan memilih metode verifikasi: DNS TXT record (nilai token unik per organisasi) atau HTTPS (file berisi token di `https://<domain>/.well-known/...` atau path yang ditentukan).
2. Domain bisa terverifikasi otomatis lewat integrasi Google Workspace (USR-036) tanpa langkah DNS/HTTPS.
3. Status domain: `pending`, `verified`, `failed`; tombol "Verify" bisa dijalankan ulang, dan sistem juga memeriksa ulang secara berkala (mis. tiap 24 jam selama pending, hingga 14 hari).
4. Organisasi dapat memverifikasi beberapa domain dan subdomain; subdomain dari domain terverifikasi tetap perlu ditambahkan eksplisit.
5. Claim accounts adalah langkah terpisah setelah verifikasi: admin memilih "Claim accounts" dan semua akun dengan email domain tersebut menjadi managed account.
6. Menghapus domain membuat akun dengan domain itu kembali unmanaged; dialog konfirmasi menyebut jumlah akun terdampak dan bahwa SSO/kebijakan tidak berlaku lagi.
7. Satu domain hanya bisa diverifikasi oleh satu organisasi; percobaan kedua ditolak dengan pesan "Domain sudah diverifikasi oleh organisasi lain".
8. Pengguna menerima email notifikasi saat akunnya menjadi managed account (berisi nama organisasi dan dampaknya).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Prasyarat untuk USR-032, USR-034, USR-035. Untuk instance self-hosted tunggal, semua akun dengan domain terverifikasi otomatis menjadi milik satu organisasi.

### USR-032 Managed accounts

**Deskripsi:** Pengelolaan akun yang email-nya berada di domain terverifikasi organisasi, termasuk aksi pemulihan akses oleh admin.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Directory > Managed accounts menampilkan daftar managed account dengan filter status, status 2SV, authentication policy, dan last active.
2. Admin dapat mengedit profil (full name, public name, job title, email) bila atribut tidak disinkronkan dari IdP.
3. Admin dapat deactivate, reactivate, dan delete account (masa tenggang 14 hari) dengan aturan USR-007 butir 4-6.
4. Admin dapat mengubah email akun ke domain terverifikasi lain milik organisasi (mis. saat user kehilangan akses email lama); email ke domain tidak terverifikasi ditolak.
5. Aksi recovery: reset password (kirim tautan reset), reset 2SV (user wajib mendaftar ulang 2SV saat login berikutnya), log out semua sesi, dan resend invite.
6. Admin dapat meng-assign authentication policy per user (detail policy di [01 Identitas](01-identitas-autentikasi-sesi.md)).
7. Managed account tetap bisa mengakses site di luar organisasi sebagai external user di organisasi lain (relevan bila produk dijalankan multi-tenant).
8. Passkey tidak tersedia untuk managed account (login memakai SSO/sandi + 2SV sesuai policy).
9. Dengan paket Keamanan Lanjutan, admin dapat melihat dan mencabut API token milik managed account.
10. Admin yang terkunci karena IdP bermasalah dapat masuk lewat authentication policy non-SSO (bypass); bila Organization admin terakhir terkunci total, pemulihan dilakukan lewat prosedur operator/CLI instance (padanan Atlassian Support).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Untuk self-hosted, prosedur pemulihan admin terakhir disediakan sebagai perintah CLI biner (mis. `jari admin recover --email`) yang tercatat di audit log; detail di [22 Administrasi Sistem](22-administrasi-sistem.md).

### USR-033 External users management

**Deskripsi:** Pengelolaan user dengan akun non-managed (email domain luar) yang mengakses site organisasi, termasuk kebijakan keamanan khusus.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Daftar user menyediakan filter "External users" (akun yang email-nya tidak berada di domain terverifikasi).
2. External user security policy (paket Keamanan Lanjutan) dapat mewajibkan 2SV dan mengatur durasi sesi maksimum untuk external users.
3. Admin dapat mencabut semua sesi external users sekaligus atau per user.
4. Admin tidak bisa mengedit profil external users (nama, email); field profil read-only dengan keterangan "Akun ini tidak dikelola organisasi Anda".
5. ⚠ Profil dan picker menampilkan penanda "External" untuk external users agar pengguna internal sadar sedang berbagi dengan pihak luar.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Bergantung pada USR-031. Kebijakan sesi di [01 Identitas](01-identitas-autentikasi-sesi.md).

### USR-034 SAML single sign-on

**Deskripsi:** Integrasi SSO berbasis SAML 2.0 dengan identity provider eksternal untuk managed account.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi; tersedia dengan paket Keamanan Lanjutan (tingkat Standard).

**Aturan perilaku (acceptance criteria):**

1. Admin menambah identity provider dengan memilih tipe: Okta, Microsoft Entra ID, Google Workspace, OneLogin, Ping, ADFS, JumpCloud, Auth0, atau "Other provider" (generik); tipe menentukan panduan konfigurasi yang ditampilkan.
2. Konfigurasi IdP: Identity provider Entity ID, SSO URL, dan public x509 certificate (PEM); sertifikat tidak valid ditolak dengan pesan galat parsing.
3. Konsol menampilkan informasi Service Provider untuk disalin ke IdP: SP Entity ID dan ACS URL (serta metadata XML yang bisa diunduh).
4. Admin menghubungkan satu atau lebih domain terverifikasi ke IdP; satu domain hanya terhubung ke satu IdP.
5. ⚠ Organisasi dapat memiliki beberapa identity provider (usulan batas 10 per organisasi), masing-masing untuk domain berbeda.
6. Admin dapat menguji SSO (login uji dengan akun admin) sebelum SSO ditegakkan lewat authentication policy; policy enforce tidak bisa diaktifkan sebelum uji berhasil minimal sekali.
7. Mendukung login SP-initiated (dari halaman login produk, berdasarkan domain email) dan IdP-initiated (dari portal IdP).
8. Saat login SSO pertama, akun dengan domain terverifikasi ditautkan atau dibuat otomatis (JIT) tanpa undangan; app access mengikuti pengaturan default/grup.
9. NameID = email; atribut nama (mis. `firstName`/`lastName`/`displayName`) dipetakan ke profil; email di assertion yang tidak cocok domain terhubung ditolak.
10. Admin dapat merotasi sertifikat (menyimpan dua sertifikat sementara), dan sistem mengirim peringatan 30, 7, dan 1 hari sebelum sertifikat kedaluwarsa.
11. Admin dapat login bypass saat IdP bermasalah lewat authentication policy non-SSO (USR-032 butir 10).
12. ⚠ SAML Single Logout tidak didukung; logout dari produk tidak me-logout sesi IdP dan sebaliknya.
13. ⚠ OIDC generik sebagai protokol SSO user tidak disediakan pada rilis awal; login Google memakai integrasi Google Workspace (USR-036). Keputusan dukungan OIDC dicatat sebagai pertanyaan terbuka.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Alur login dan authentication policy di [01 Identitas](01-identitas-autentikasi-sesi.md). Validasi assertion (signature, audience, NotOnOrAfter, replay) wajib mengikuti spesifikasi SAML 2.0.

### USR-035 User provisioning (SCIM)

**Deskripsi:** Sinkronisasi otomatis user dan grup dari identity provider lewat SCIM 2.0.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi (mengonfigurasi), IdP (klien SCIM); tersedia dengan paket Keamanan Lanjutan.

**Aturan perilaku (acceptance criteria):**

1. Admin membuat identity provider directory yang menghasilkan SCIM base URL dan API key (ditampilkan sekali).
2. Provision user: IdP dapat membuat user, memperbarui atribut (nama, email, job title, department, organization, manager), dan menonaktifkan user (`active=false`) lewat endpoint `/Users` sesuai RFC 7643/7644.
3. IdP dapat mem-push grup dan keanggotaannya lewat endpoint `/Groups`; grup hasil sinkronisasi read-only di produk (USR-015 butir 9).
4. Deprovision (`active=false` atau DELETE) menonaktifkan managed account atau mencabut aksesnya sesuai pengaturan directory.
5. Admin memetakan grup hasil sinkronisasi ke app access (mis. grup IdP "eng" memberi akses Jira role User).
6. ⚠ Bila grup yang di-push memiliki nama sama dengan grup lokal, admin memilih: gabungkan (grup lokal menjadi dikelola IdP) atau tolak (sinkronisasi grup gagal dengan galat konflik di log).
7. Konsol menampilkan log provisioning (operasi, objek, hasil) dan galat per operasi selama minimal 30 hari.
8. Provisioning tidak otomatis memberi akses app; akses hanya diberikan lewat pemetaan grup (butir 5) atau default akses app.
9. Admin dapat merotasi/regenerasi SCIM API key; key lama langsung tidak berlaku.
10. Batas user plan berlaku saat provisioning; user yang melebihi batas dibuat tanpa akses app dan dicatat sebagai galat `USER_LIMIT_REACHED`.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Endpoint SCIM terpisah dari REST API publik (`/scim/v2/directory/{directoryId}/...`). Rujukan: USR-007 butir 9, USR-020.

### USR-036 Google Workspace & Microsoft Entra integration

**Deskripsi:** Integrasi khusus dengan Google Workspace (sinkronisasi user/grup dan login Google) dan Microsoft Entra ID.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Admin menghubungkan Google Workspace dengan akun admin Google (OAuth consent dengan scope directory read-only); domain Google otomatis terverifikasi.
2. Admin memilih menyinkronkan semua user atau hanya grup tertentu.
3. Sinkronisasi berjalan berkala (usulan tiap 1 jam) dan bisa dipicu manual dengan tombol "Sync now"; status dan waktu sinkronisasi terakhir ditampilkan.
4. Login Google tersedia sebagai metode SSO untuk managed account dari domain Google yang terhubung.
5. Microsoft Entra ID didukung lewat aplikasi gallery/enterprise app yang memakai SAML (USR-034) dan SCIM (USR-035); konsol menyediakan panduan konfigurasi Entra.
6. Disconnect integrasi menampilkan dampaknya: user hasil sinkronisasi tetap ada tetapi menjadi tidak tersinkron, grup hasil sinkronisasi menjadi grup lokal yang bisa diedit (atau dihapus sesuai pilihan admin).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Login sosial Google untuk akun non-managed (bila ada) dibahas di [01 Identitas](01-identitas-autentikasi-sesi.md).

### Keamanan & perlindungan data organisasi

### USR-037 Atlassian Guard subscription

**Deskripsi:** Paket keamanan berlangganan (padanan Atlassian Guard, eks Atlassian Access) yang membuka fitur identitas dan keamanan lanjutan untuk organisasi. Di produk ini disebut paket Keamanan Lanjutan dan diimplementasikan sebagai kumpulan feature flag organisasi.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi, Billing admin.

**Aturan perilaku (acceptance criteria):**

1. Tingkat Standard membuka: SSO (USR-034), SCIM (USR-035), authentication policies, external user policy (USR-033), data security policies (USR-040), kontrol API token user, dan audit log lanjutan (USR-045).
2. ⚠ Tingkat Premium menambahkan: Detect (USR-048), content scanning, dan retensi audit log lebih panjang.
3. Penagihan dihitung per managed account unik lintas app (satu akun dengan akses ke tiga app dihitung satu kali).
4. Plan Enterprise sudah termasuk tingkat Standard tanpa langganan terpisah.
5. Admin dapat memulai trial dan mengaktifkan langganan dari konsol admin; saat trial/langganan berakhir, fitur terkunci menjadi read-only (konfigurasi disimpan tetapi tidak ditegakkan) dan admin diberi peringatan 14 hari sebelumnya.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Packaging komersial; untuk distribusi open source keputusan apakah fitur ini dibuka gratis dicatat di Pertanyaan terbuka. Feature flag dievaluasi di satu layanan `entitlements` (lihat [22 Administrasi Sistem](22-administrasi-sistem.md)).

### USR-038 IP allowlisting

**Deskripsi:** Membatasi akses ke app hanya dari alamat atau rentang IP tertentu.

**Prioritas / Fase / Tier:** P2 / Fase 4 / premium

**Aktor & izin:** Admin Organisasi; tersedia pada plan Premium dan Enterprise.

**Aturan perilaku (acceptance criteria):**

1. Admin membuat allowlist berisi IP tunggal atau rentang CIDR (IPv4 dan IPv6); entri tidak valid ditolak dengan pesan galat format.
2. Allowlist diterapkan ke app/site tertentu yang dipilih.
3. Allowlist berlaku untuk akses web, mobile, dan REST API (termasuk API token dan OAuth).
4. ⚠ Admin dapat mengecualikan ekstensi dan integrasi terpasang (app user) dari allowlist, per ekstensi.
5. Pengguna yang mengakses dari IP di luar allowlist melihat halaman/respons 403 "Akses dari alamat IP Anda (x.x.x.x) tidak diizinkan oleh organisasi".
6. Saat menyimpan, bila IP admin saat ini tidak termasuk allowlist, sistem menampilkan peringatan keras agar admin tidak mengunci dirinya dan meminta konfirmasi eksplisit.
7. Setiap allowlist bisa diaktifkan/dinonaktifkan tanpa dihapus.
8. Setiap perubahan allowlist dicatat sebagai audit event dengan nilai lama dan baru.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** IP klien diambil dari header proxy tepercaya yang dikonfigurasi (`X-Forwarded-For` hanya dari proxy terdaftar). Portal customer Service Management dikecualikan secara default (modul 30).

### USR-039 Mobile app policy

**Deskripsi:** Kebijakan keamanan app mobile (mobile app management) untuk managed account.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi; bagian dari paket Keamanan Lanjutan.

**Aturan perilaku (acceptance criteria):**

1. Policy dibuat per app mobile dan per platform (iOS/Android).
2. Policy dapat mewajibkan PIN atau biometrik untuk membuka app.
3. Policy dapat memblokir screenshot dan screen recording.
4. Policy dapat memblokir copy/paste dan berbagi data ke app lain.
5. Policy dapat memblokir download/ekspor lampiran.
6. Policy dapat memblokir perangkat jailbroken/rooted.
7. ⚠ Policy dapat mensyaratkan versi OS minimum dan enkripsi perangkat.
8. Organisasi yang memakai MDM (Intune, Workspace ONE) dapat memilih mode "kelola lewat MDM" sebagai alternatif policy bawaan.
9. Policy dapat memblokir akses mobile sepenuhnya (app mobile menampilkan pesan "Organisasi Anda tidak mengizinkan akses dari perangkat mobile").

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Hanya relevan jika produk ini memiliki app mobile native; saat ini repo hanya memiliki SPA web. Butir 9 (blokir akses mobile, berdasarkan klien/user agent) dapat ditegakkan server-side lebih dulu; butir lain menunggu keputusan app mobile (Pertanyaan terbuka).

### USR-040 Data security policies

**Deskripsi:** Kebijakan perlindungan data untuk membatasi akses ekstensi, public link, akses anonim, ekspor, dan akses AI pada cakupan tertentu.

**Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise

**Aktor & izin:** Admin Organisasi; tersedia di plan Enterprise atau sebagai add-on paket Keamanan Lanjutan tingkat Standard (USR-037).

**Aturan perilaku (acceptance criteria):**

1. Cakupan policy dapat berupa seluruh organisasi, app/site tertentu, atau konten dengan classification level tertentu (USR-041).
2. Rule "blokir akses app": ekstensi/integrasi terpasang tidak bisa membaca konten dalam cakupan (API mengembalikan 403 untuk app user).
3. Rule "blokir public links": tautan publik ke filter/dashboard/konten dalam cakupan dinonaktifkan.
4. Rule "blokir anonymous access": permission "Anyone on the web" tidak berlaku untuk konten dalam cakupan (USR-024 butir 4).
5. ⚠ Rule "blokir export": ekspor CSV/Excel/PDF/Word dari pencarian, work item, dan dashboard dalam cakupan ditolak.
6. ⚠ Rule "blokir akses AI": fitur AI opsional tidak boleh memakai konten dalam cakupan sebagai konteks.
7. Policy bisa diaktifkan/dinonaktifkan; setiap perubahan dicatat sebagai audit event.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Evaluasi policy harus terpusat (policy decision point) dan dipanggil dari modul ekspor ([13](13-pencarian-filter-ekspor.md)), izin ([03](03-izin-peran-keamanan.md)), dan ekstensi ([25](25-integrasi-impor-ekspor-ekstensi.md)). Pasangan fitur: PERM-065 Data Security Policies ([03 Izin](03-izin-peran-keamanan.md), penegakan di Authorizer) dan ADM-046 Data security policies ([22 Administrasi Sistem](22-administrasi-sistem.md), entry point admin); ketiganya P3 / Fase 4 / enterprise dan definisi rule serta cakupan mengikuti fitur ini. Fitur terkait Data Redaction (ISS-043 di [06 Work Item Inti](06-work-item-inti.md), PERM-066 di modul 03) memakai klasifikasi dan policy yang sama.

### USR-041 Data classification

**Deskripsi:** Label klasifikasi data (mis. Public, Internal, Confidential) yang didefinisikan organisasi dan diterapkan ke proyek/space dan work item untuk menandai sensitivitas data.

**Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise

**Aktor & izin:** Admin Organisasi (mendefinisikan level), Admin Situs/App Admin dan Admin Proyek (menerapkan default proyek), pengguna dengan Edit work items (mengubah klasifikasi work item).

**Aturan perilaku (acceptance criteria):**

1. Admin organisasi mendefinisikan classification levels dengan nama, warna, deskripsi, dan guideline di konsol admin.
2. Level bisa dipublish (tersedia untuk dipakai) atau diarsipkan (tidak bisa dipilih baru, nilai lama tetap tampil); satu level dapat ditandai sebagai default organisasi.
3. Admin proyek menetapkan classification default di level proyek/space.
4. Classification per work item mengikuti default proyek atau meng-override-nya (lebih tinggi atau lebih rendah); perubahan tercatat di history work item.
5. Classification level dapat dipakai sebagai cakupan data security policy (USR-040 butir 1).
6. ⚠ Classification tampil sebagai badge di tampilan work item dan hasil pencarian, dan dapat dicari via JQL (mis. `classification = Confidential`).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Field classification diperlakukan sebagai system field (modul 05); JQL di [12 JQL](12-jql-bahasa-kueri.md). Tersedia di plan Enterprise atau sebagai add-on paket Keamanan Lanjutan (USR-037). Pasangan fitur: PERM-064 Data Classification ([03 Izin](03-izin-peran-keamanan.md), penegakan dan indikator) dan ADM-045 Data classification ([22 Administrasi Sistem](22-administrasi-sistem.md), halaman admin dan aturan pewarisan); ketiganya P3 / Fase 4 / enterprise dan definisi level mengikuti fitur ini.

### USR-042 Data residency

**Deskripsi:** Memilih dan memindahkan lokasi penyimpanan data in-scope app ke region tertentu.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. ⚠ Admin memilih realm penyimpanan dari daftar region yang dioperasikan (rujukan Jira: US, EU, Germany, Australia, Singapore, Canada, UK, Japan, India, South Korea); daftar realm adalah konfigurasi operator.
2. Setiap app menampilkan status residency: `pinned`, `not pinned`, atau `moving`.
3. Admin menjadwalkan migrasi realm dan memilih jendela migrasi.
4. Pengguna diberi notifikasi sebelum migrasi, dan app dapat mengalami downtime/read-only selama migrasi dengan banner pemberitahuan.
5. Konsol menampilkan daftar data in-scope (work item, lampiran, komentar) vs out-of-scope (mis. log analitik).
6. Ekstensi terpasang yang mendeklarasikan dukungan data residency ikut dipindahkan bersama data app; ekstensi tanpa dukungan ditampilkan dengan status "tidak mendukung residency" (padanan app Marketplace pihak ketiga di luar cakupan).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Hanya bermakna bila produk dioperasikan sebagai SaaS multi-region; deployment self-hosted secara alami memilih lokasinya sendiri. Rujukan: [27 NFR](27-nfr-arsitektur-platform.md). Pasangan fitur: ADM-048 Data residency ([22 Administrasi Sistem](22-administrasi-sistem.md)) adalah entry point admin; keduanya P3 / Fase 4 / platform dan definisi realm serta status mengikuti fitur ini. ⚠ Ketersediaan per plan (rujukan: plan berbayar, bukan Free) disimpan sebagai entitlement plan, bukan pembeda tier. Perluasan cakupan data Service Management (JSM-AI-011 JSM Data Residency di [35 JSM Admin](35-jsm-ai-otomasi-api-admin-csm.md), P2 / Fase 7 / jsm) memakai mekanisme realm dan status yang sama dan hanya menambah daftar data in-scope app Service Management.

### USR-043 Bring your own key (BYOK) encryption

**Deskripsi:** Enkripsi data memakai kunci milik pelanggan di AWS KMS.

**Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise

**Aktor & izin:** Admin Organisasi; hanya plan Enterprise.

**Aturan perilaku (acceptance criteria):**

1. Admin menghubungkan AWS KMS key milik pelanggan (ARN key dan role akses lintas akun); koneksi divalidasi dengan operasi uji encrypt/decrypt.
2. Rotasi kunci didukung tanpa downtime (envelope encryption: data key dienkripsi ulang dengan kunci baru).
3. Mencabut akses kunci di KMS menghentikan akses ke data terenkripsi; app menampilkan status "Data tidak dapat diakses: kunci enkripsi dicabut".
4. Penggunaan kunci tercatat di AWS CloudTrail milik pelanggan.
5. ⚠ BYOK hanya tersedia untuk plan Enterprise dan site tertentu/baru (aktivasi pada site yang sudah berisi data memerlukan migrasi).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Arsitektur enkripsi di [27 NFR](27-nfr-arsitektur-platform.md). Untuk self-hosted, padanannya adalah enkripsi at-rest dengan kunci yang dikelola operator (di luar cakupan modul ini).

### USR-044 GDPR & personal data APIs

**Deskripsi:** Kepatuhan privasi untuk data personal user, termasuk API bagi ekstensi dan integrasi.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Integrator/Developer API (ekstensi), Sistem.

**Aturan perilaku (acceptance criteria):**

1. Personal Data Reporting API: ekstensi melaporkan daftar `accountId` yang disimpannya beserta waktu terakhir diperbarui.
2. Respons API berisi daftar akun yang ditutup (ekstensi wajib menghapus data personal) dan akun yang diperbarui (ekstensi wajib me-refresh data cache).
3. Atribut `emailAddress` pada respons user hanya dikembalikan bila profile visibility user atau izin ekstensi mengizinkan; selain itu field dihilangkan.
4. `GET /rest/api/3/user/email?accountId=` mengembalikan email hanya untuk ekstensi yang diberi izin khusus; pemanggil lain mendapat 403.
5. Data personal (nama, email, avatar, atribut profil) dihapus setelah akun ditutup (USR-007 butir 6); konten buatan user tetap ada dengan referensi `accountId`.
6. Display name akun yang dihapus menjadi "Former user" di seluruh UI dan API.
7. ⚠ User dapat mengajukan permintaan export data personalnya (data subject access request) yang diproses admin/operator dan menghasilkan arsip unduhan.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Penghapusan harus menjangkau indeks pencarian, cache, dan payload event/webhook historis yang masih tersimpan. Rujukan: [23 REST API Publik](23-api-publik-webhook.md), [27 NFR](27-nfr-arsitektur-platform.md).

### Audit & insight

### USR-045 Organization audit log

**Deskripsi:** Log aktivitas administratif dan keamanan level organisasi di konsol admin.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Organisasi (User access admin hanya melihat event user/grup untuk app yang dikelolanya).

**Aturan perilaku (acceptance criteria):**

1. Event yang dicatat minimal: perubahan peran admin, user ditambah/dihapus/di-suspend/dinonaktifkan, perubahan grup dan keanggotaan, perubahan policy, domain, API key, invite link, dan User access settings; dengan paket Keamanan Lanjutan juga event login dan aktivitas app.
2. Filter: rentang tanggal, actor, jenis aktivitas, dan lokasi/IP.
3. Pencarian teks bebas pada ringkasan event, nama actor, dan target.
4. Detail event menampilkan waktu (UTC dan zona waktu admin), actor (termasuk tipe actor: user/service/app/SCIM/sistem), target, IP, dan user agent.
5. Hasil filter dapat diekspor ke CSV (asinkron untuk hasil besar).
6. ⚠ Retensi berbeda per plan dan paket Keamanan Lanjutan dan angkanya mengikuti NFR-017 butir 5 untuk scope `org` (default 180 hari, hingga 7 tahun dengan paket Premium/keamanan lanjutan); event di luar retensi dihapus oleh job retensi NFR-017.
7. Audit log API (events endpoint, USR-030 butir 7) mendukung paginasi kursor dan filter waktu untuk dikonsumsi SIEM.
8. ⚠ Dengan paket Keamanan Lanjutan, admin dapat mengatur activity tracking aktivitas app (view/export konten) per app.
9. Dengan paket Keamanan Lanjutan, event dapat di-stream ke SIEM (Splunk, dll.) lewat konektor/webhook bertanda tangan.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Audit log bersifat append-only (tidak bisa diedit/dihapus oleh admin); penyimpanan, tamper-evidence, dan retensi dimiliki [NFR-017 Security Audit Logging](27-nfr-arsitektur-platform.md) (`scope = org`), sedangkan audit log app produk ada di USR-046 dan UI-nya di ADM-014; skema event di [28 Model Data Inti & Katalog Event](28-model-data-inti.md). Semua fitur modul ini yang menyebut "tercatat di audit log" menulis ke sini.

### USR-046 Jira audit log

**Deskripsi:** Log audit di Settings > System > Audit log yang mencatat perubahan konfigurasi app. Di modul ini USR-046 hanya menetapkan kategori dan isi event users & groups; UI admin, filter, ekspor, dan API dispesifikasikan lengkap di [ADM-014 Audit log (Jira settings)](22-administrasi-sistem.md), sedangkan penyimpanan append-only, tamper-evidence, dan retensi di [NFR-017 Security Audit Logging](27-nfr-arsitektur-platform.md).

**Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard

**Aktor & izin:** Admin Situs/App Admin (Administer Jira).

**Pembagian kepemilikan audit log:**

| Aspek | Pemilik | Isi |
|---|---|---|
| Penyimpanan | NFR-017 | Tabel `audit_records` append-only (scope `app`/`org`/`automation`), `AuditRecorder.Record`, hash chain dan alat verifikasi (tamper-evidence), job retensi, streaming SIEM. |
| Retensi | NFR-017 butir 5 | Satu-satunya acuan angka retensi; ADM-014, USR-045, dan USR-046 tidak menetapkan angka sendiri. |
| UI admin produk | ADM-014 | Halaman Settings > System > Audit log: daftar entri, filter, detail "Changed values"/"Associated items", ekspor CSV, REST API `auditing/record`. |
| Kategori event user/grup | USR-046 (modul ini) | Daftar event users & groups beserta isi entrinya (butir 1 di bawah). |
| Audit log organisasi | USR-045 | Scope `org` di konsol admin organisasi. |

**Aturan perilaku (acceptance criteria):**

1. Kategori event mengikuti daftar kategori ADM-014 butir 2 (users & groups, permissions, workflows, fields, projects, schemes, system, dan seterusnya). Modul ini menetapkan isi kategori users & groups: user dibuat/diundang, diaktifkan/dinonaktifkan, dihapus, diberi/dicabut akses app (USR-008), grup dibuat/diganti nama/dihapus, anggota grup ditambah/dikeluarkan, dan perubahan default group/akses produk; setiap event menyimpan object (user atau grup) dan associated items (grup, app, atau proyek terdampak).
2. Filter rentang tanggal, kategori, dan teks bebas identik dengan ADM-014 butir 3; tidak ada filter tambahan khusus USR-046.
3. Detail entri menampilkan summary, author, IP, objek yang berubah, serta nilai lama dan baru (diff per atribut) sesuai ADM-014 butir 4; untuk event keanggotaan grup, "Changed values" memuat daftar anggota yang ditambah/dikeluarkan.
4. ⚠ Retensi tidak ditetapkan di sini dan identik dengan NFR-017 butir 5 untuk scope `app` (default 90 hari; pilihan 30, 90, 180, 365 hari; maksimum per tier sesuai tabel NFR-017), sama seperti ADM-014 butir 6 dan 11.
5. ⚠ Admin dapat menampilkan atau menyembunyikan perubahan yang dilakukan oleh sistem/ekstensi (toggle "Tampilkan perubahan oleh app") pada UI ADM-014; event yang disembunyikan tetap tersimpan.
6. ⚠ Hasil filter dapat diekspor ke CSV dengan batas dan mekanisme asinkron yang sama dengan ADM-014 butir 5.
7. REST API `GET /rest/api/3/auditing/record` mendukung parameter `offset`, `limit`, `filter`, `from`, `to` sebagaimana dispesifikasikan di ADM-014 butir 7; event users & groups dikembalikan dengan kategori yang sama.
8. Perubahan work item tidak masuk audit log ini; perubahan tersebut ada di history work item (modul 08).

**Team-managed vs company-managed:** Perubahan konfigurasi proyek team-managed (field, workflow per proyek) tercatat dengan kategori projects dan menyebut proyeknya (sama dengan ADM-014).

**Catatan & ketergantungan:** Entri ditulis lewat `AuditRecorder.Record` ke tabel `audit_records` dengan `scope = app` (NFR-017 butir 1), sehingga sifat append-only dan tamper-evidence (hash chain NFR-017 butir 3) berlaku tanpa implementasi terpisah; penyimpanan sama dengan USR-045 yang memakai `scope = org`. Coverage level dan coverage area hanya ada di Data Center (USR-049). Rujukan: [22 Administrasi Sistem](22-administrasi-sistem.md) (ADM-014), [27 NFR](27-nfr-arsitektur-platform.md) (NFR-017), [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md). Tab History work item dirinci di CMT-022 dan CMT-023.

### USR-047 Organization insights (product usage & security insights)

**Deskripsi:** Halaman Insights di konsol admin yang menampilkan statistik pemakaian app (user aktif, adopsi) dan insight keamanan untuk keputusan lisensi, adopsi, dan postur keamanan.

**Prioritas / Fase / Tier:** P3 / Fase 4 / premium

**Aktor & izin:** Admin Organisasi.

**Aturan perilaku (acceptance criteria):**

1. Grafik user aktif harian/bulanan per app dan per site.
2. ⚠ Grafik tren penambahan user dan pemakaian lisensi/seat dari waktu ke waktu.
3. Filter rentang waktu (7, 30, 90, 365 hari) serta filter per site/app.
4. ⚠ Insight keamanan: persentase cakupan 2SV/SSO, jumlah managed vs unmanaged/external accounts, dan rekomendasi tindakan; sebagian insight hanya tampil dengan paket Keamanan Lanjutan.
5. ⚠ Statistik pemakaian fitur AI opsional per app (bila fitur AI diaktifkan).
6. ⚠ Data insight dapat diekspor ke CSV.
7. ⚠ Halaman hanya bisa diakses Organization admin; ketersediaan metrik bergantung plan (Premium/Enterprise) atau paket Keamanan Lanjutan, dan metrik yang terkunci ditampilkan dengan penjelasan upgrade.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Berbeda dari USR-048 (deteksi ancaman) dan USR-006 (data per user). Sumber data: agregasi last active (USR-006) dan event audit. Laporan proyek ada di [16 Laporan & Insight](16-laporan-insight.md).

### USR-048 Atlassian Guard Detect & insights

**Deskripsi:** Fitur keamanan lanjutan (padanan Guard Detect) untuk mendeteksi aktivitas mencurigakan, memberi alert, dan memindai konten sensitif.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Admin Organisasi, tim keamanan; bagian dari paket Keamanan Lanjutan tingkat Premium.

**Aturan perilaku (acceptance criteria):**

1. Detections bawaan untuk aktivitas tidak biasa: bulk export, akses massal ke banyak work item dalam waktu singkat, perubahan peran admin, dan eskalasi privilege (ambang dapat dikonfigurasi).
2. Alert dikirim lewat email, webhook ke Slack/Teams, dan integrasi SIEM; setiap alert memiliki status (open, investigating, resolved).
3. ⚠ Content scanning mendeteksi data sensitif (mis. kredensial, nomor kartu) di deskripsi dan komentar work item, dengan aturan pola yang dapat dikonfigurasi.
4. Insight keamanan tambahan dari Detect ditampilkan di halaman Organization insights (USR-047).
5. Dashboard keamanan menampilkan rekomendasi (mis. aktifkan 2SV, cabut API token tidak terpakai).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Product requests/discovered products dipisah ke USR-014. Sumber sinyal adalah audit log (USR-045). Integrasi Slack/Teams di [25 Integrasi](25-integrasi-impor-ekspor-ekstensi.md).

### Referensi Data Center (tidak dibangun)

### USR-049 Jira Data Center user management (legacy)

**Deskripsi:** Fitur manajemen user khusus Jira Data Center yang tidak ada di Cloud. Dicatat hanya sebagai referensi; tidak diimplementasikan.

**Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated

**Aktor & izin:** Admin Jira Data Center (rujukan).

**Aturan perilaku (acceptance criteria):**

1. User directories (Internal, LDAP, Active Directory, Crowd, delegated LDAP) dengan urutan directory dan sinkronisasi: tidak dibangun; padanan clone adalah SCIM (USR-035) dan SAML (USR-034).
2. Secure administrator sessions (websudo): tidak dibangun sebagai fitur DC; kebutuhan re-autentikasi untuk aksi sensitif dibahas di modul 01.
3. Maximum authentication attempts dan reset failed login count: tidak dibangun sebagai pengaturan admin DC; proteksi brute force ada di modul 01.
4. Mode Public/Private dengan self-signup dan CAPTCHA signup: tidak dibangun; padanan clone adalah User access settings (USR-012).
5. Anonymize users (DC 8.7+): tidak dibangun; padanan clone adalah delete account dan "Former user" (USR-007, USR-044).
6. Pengaturan email visibility global (public/hidden/masked/logged-in only): tidak dibangun; padanan clone adalah profile visibility per user (modul 21).
7. Pemisahan Jira System Administrators vs Jira Administrators: tidak dibangun; clone memakai Organization admin vs App admin (USR-021).
8. Global permission "Create shared objects": tidak dibangun; sharing dikontrol "Share dashboards and filters" (USR-022).
9. Kirim email ke grup/project role dari admin: tidak dibangun di modul ini.
10. Session timeout lewat web.xml dan remember-me cookie: tidak dibangun; kebijakan sesi di modul 01.
11. Coverage level audit log: tidak dibangun; audit log clone selalu penuh (USR-046).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:** Lihat bagian Di luar cakupan & ekosistem.

## Pertanyaan terbuka

- Apakah produk ini akan dioperasikan sebagai SaaS multi-tenant (banyak organisasi per deployment) atau hanya self-hosted satu organisasi per instance? Jawaban menentukan cakupan USR-001 (pemindahan site, org switcher), USR-014 (discovered products), dan USR-042 (data residency).
- Untuk distribusi open source, apakah fitur paket Keamanan Lanjutan (SSO, SCIM, data security policy, audit log lanjutan) dibuka gratis atau tetap di balik lisensi/entitlement berbayar?
- Apakah batas plan Free (10 user) dan pemisahan plan Free/Standard/Premium/Enterprise berlaku juga untuk deployment self-hosted, atau batas user dinonaktifkan secara default?
- Apakah OIDC generik didukung sebagai protokol SSO (selain SAML dan Google Workspace), mengingat banyak IdP self-hosted (Keycloak, Authentik) lebih mudah memakai OIDC?
- Apakah produk ini akan memiliki app mobile native? Bila tidak, USR-039 dibatasi pada pemblokiran akses klien mobile/web-mobile saja.
- Berapa kuota guest gratis per seat berbayar dan berapa kuota service account non-billable per organisasi?
- Apakah nama grup sistem memakai nama asli Jira (`jira-users-<site>`, `atlassian-addons-admin`) demi kompatibilitas impor, atau nama netral clone dengan alias impor?
- Apakah endpoint kompatibel `/rest/api/3/...` disediakan sejak Fase 1, atau Fase 1 cukup memakai `/api/v1` internal dan lapisan kompatibel ditambahkan di Fase 3 bersama REST API publik?

## Di luar cakupan & ekosistem

- **Fitur Data Center-only (USR-049):** user directories LDAP/AD/Crowd, websudo, maximum authentication attempts, mode Public/Private dengan CAPTCHA, anonymize users, email visibility global, Jira System Administrators, "Create shared objects", kirim email ke grup dari admin, session timeout web.xml, coverage level audit log. Dicatat hanya sebagai rujukan; padanan Cloud disebut per butir.
- **Model original user management (USR-002):** tidak dibangun; clone hanya memakai model centralized.
- **App Marketplace pihak ketiga:** app manajemen user/akses tambahan (mis. bulk user management, laporan lisensi, provisioning LDAP pihak ketiga) tidak dibangun sebagai fitur native; ekstensi pihak ketiga hanya dirujuk lewat mekanisme ekstensibilitas di [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md).
- **Layanan Atlassian lintas produk:** Confluence, Atlassian Home sebagai produk terpisah (entitas Goals dan Teams native tetap dibangun di [18 Timeline, Rencana & Tampilan](18-timeline-rencana-tampilan-proyek.md)), Atlassian Support sebagai jalur pemulihan, Atlassian Intelligence/Rovo, dan billing Atlassian (my.atlassian.com) tidak direplikasi; padanannya di clone adalah fitur lokal (CLI pemulihan admin, AI opsional, billing internal di modul 22).
- **Infrastruktur cloud vendor:** realm data residency dan BYOK berbasis AWS KMS hanya relevan untuk operasi SaaS; untuk self-hosted, lokasi data dan manajemen kunci menjadi tanggung jawab operator.
- **MDM pihak ketiga (Intune, Workspace ONE):** hanya dirujuk sebagai alternatif kebijakan mobile (USR-039), tidak diintegrasikan secara native.
