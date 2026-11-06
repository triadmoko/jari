# 22 — Administrasi Sistem & Pengaturan Global

Modul ini mendefinisikan konsol administrasi global produk ini dan seluruh pengaturan tingkat situs yang berlaku lintas proyek: pintu masuk Settings beserta registry halaman admin dan model akses admin, General configuration (identitas situs, mode akses, toggle perilaku global, opsi warisan Data Center), internasionalisasi, tampilan (look and feel, announcement banner, default user preferences), Advanced settings berbasis properti bertipe, pengaturan lampiran dan email keluar/masuk, audit log konfigurasi dan alat diagnosis admin (Permission helper, Notification helper, Find your field), administrasi objek bersama (filter, dashboard, system dashboard, kolom navigator default), kategori/arsip/trash proyek dan kebijakan siklus hidup data, pengaturan aplikasi kerja (parallel sprints, Plans, Teams, Goals), titik masuk integrasi (webhook, development tools, application links, issue collector, manajemen app/ekstensi), backup/restore/impor/migrasi, batas objek dan kuota plan, Site optimizer, kebijakan keamanan dan tata kelola data, sandbox/release tracks/feature preview/AI, serta langganan dan siklus hidup situs. Modul ini memiliki halaman admin dan penyimpanan pengaturan global; logika domain yang dikonfigurasi (permission, notifikasi, field, workflow, filter, dashboard, sprint) tetap dimiliki modul masing-masing dan hanya dirujuk. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users` tanpa kolom peran admin, JWT di context `sess`, middleware bahasa EN/ID, konfigurasi dari environment di `internal/core/config`), sehingga modul ini adalah pekerjaan baru yang dikemas sebagai feature module `internal/features/admin` (handler, DTO, use case) dengan entitas `SiteSetting` (registry key-value bertipe), `AdminPage` (registry halaman), `AuditRecord`, `MailServer`, `MailHandler`, `BackupJob`, `ProjectCategory`, dan `Announcement` di `internal/shared/domain/entity`, endpoint di bawah `/api/v1/admin/...`, serta halaman SPA `/settings/...` di `frontend/`.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — terminologi Work items/Spaces (VISI-002), plan tier & feature flag (VISI-003).
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — step-up auth, kebijakan sesi, API token.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — peran Admin Organisasi/Admin Situs (USR-021), global permissions (USR-022), anonymous access (USR-024), audit log organisasi dan Jira (USR-045, USR-046), IP allowlist (USR-038), data residency (USR-042), BYOK (USR-043).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — matriks peran admin (PERM-003), Administer Jira (PERM-007), Browse users and groups (PERM-008), data classification (PERM-064), data security policies (PERM-065).
- [04 Proyek](04-proyek.md) — kategori proyek (PRJ-019), arsip (PRJ-032), trash (PRJ-033), aturan nama/key (PRJ-011, PRJ-012), issue collector proyek (PRJ-061), Projects administration (PRJ-063).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — hierarki work type (CFG-006, CFG-007), Find your field (CFG-051), issue features (CFG-057), custom field optimizer (CFG-040), translations (CFG-061).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — Attachment settings and limits (CMT-016), file storage (CMT-019).
- [11 Backlog & Sprint](11-backlog-sprint.md) — Parallel sprints (SPR-023).
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — indexing (SRC-003), kolom (SRC-025), admin filters (SRC-049), ekspor (SRC-050).
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — system dashboard (DSH-009), admin dashboard management (DSH-014, DSH-015).
- [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md) — Plans, Teams, Goals.
- [19 Notifikasi & Email](19-notifikasi-email.md) — notification scheme, template email, Events.
- [20 Otomasi](20-otomasi.md) — global automation rules, kuota eksekusi.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — top navigation, Personal settings, tema.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — payload webhook, REST audit, rate limit.
- [24 Integrasi DevOps](24-integrasi-devops.md) — DVCS, build/deployment, Toolchain.
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — importer CSV/JSON, model ekstensi, application links.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — project email address.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — backup infrastruktur, job queue, observability, multi-tenancy.
- [28 Model Data Inti](28-model-data-inti.md) — entitas pengaturan, event domain `admin.*`.
- [35 JSM AI & Admin](35-jsm-ai-otomasi-api-admin-csm.md) — pengaturan admin khusus Service Management.

## Persona terkait

- **Admin Situs** — persona utama, dengan varian peran Admin Organisasi, Admin Akses Pengguna, dan Admin App sesuai definisi di [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) (bagian Persona, USR-021) dan ADM-002: mengelola General configuration, tampilan, email, lampiran, audit log, objek bersama, backup, webhook, integrasi, kebijakan keamanan, batas, dan langganan.
- **Admin Proyek/Project Lead** — memakai Permission/Notification helper dan Find your field untuk proyeknya, menetapkan kategori proyek, mengelola issue collector dan toolchain proyek.
- **Scrum Master/Team Lead** — terdampak pengaturan parallel sprints dan Teams; meminta admin mengaktifkannya.
- **Developer/Anggota Tim** — terdampak default preferences, banner, batas objek, dan pesan galat kuota; mengajukan permintaan app.
- **Product Owner/PM** — memakai hierarki di atas Epic dan Goals yang diaktifkan admin.
- **Stakeholder/Viewer** — melihat system dashboard dan banner tanpa hak ubah.
- **Integrator/Developer API** — mendaftarkan webhook, application link, OAuth credentials, serta membaca audit record lewat REST.
- **Tamu/Kolaborator Eksternal** — terdampak anonymous access, banner publik, dan system dashboard anonim.

## User story modul

- Sebagai Admin Situs, saya ingin satu menu Settings yang hanya menampilkan halaman admin yang boleh saya buka beserta pencarian halaman agar saya cepat menemukan pengaturan tanpa tersesat (ADM-001, ADM-002).
- Sebagai Admin Situs, saya ingin mengubah judul situs, akses anonim, serta toggle voting, watching, dan unassigned untuk seluruh situs agar perilaku produk sesuai kebijakan organisasi (ADM-003, ADM-004).
- Sebagai Admin Situs, saya ingin menetapkan bahasa dan time zone default serta logo dan warna situs agar pengguna baru langsung mendapat pengalaman yang sesuai identitas instansi (ADM-005, ADM-006, ADM-008).
- Sebagai Admin Situs, saya ingin memasang announcement banner untuk jadwal pemeliharaan agar semua pengguna tahu sebelum terjadi downtime (ADM-007).
- Sebagai Admin Situs, saya ingin mengonfigurasi SMTP dan mengirim email uji agar notifikasi dipastikan terkirim sebelum tim mulai bekerja (ADM-012).
- Sebagai Admin Situs, saya ingin mail handler yang membuat work item atau komentar dari kotak surat bersama agar permintaan via email tidak tercecer (ADM-013).
- Sebagai Admin Situs, saya ingin audit log yang dapat difilter dan diekspor dengan nilai lama dan baru agar setiap perubahan konfigurasi dapat ditelusuri saat audit (ADM-014).
- Sebagai Admin Proyek/Project Lead, saya ingin Permission helper dan Notification helper menjelaskan mengapa seorang pengguna tidak bisa melihat work item atau tidak menerima email agar keluhan dapat diselesaikan tanpa menebak konfigurasi (ADM-015, ADM-016, ADM-017).
- Sebagai Admin Situs, saya ingin memindahkan kepemilikan filter dan dashboard milik pengguna yang keluar agar board, gadget, dan langganan tetap berjalan (ADM-018, ADM-019).
- Sebagai Admin Situs, saya ingin memulihkan proyek yang terhapus dari Trash dalam 60 hari agar kesalahan hapus tidak berujung kehilangan data (ADM-023, ADM-025).
- Sebagai Scrum Master/Team Lead, saya ingin admin dapat mengaktifkan parallel sprints agar dua sub-tim di satu board dapat menjalankan sprint bersamaan (ADM-026).
- Sebagai Integrator/Developer API, saya ingin membuat webhook dengan filter JQL dan secret HMAC agar sistem eksternal hanya menerima event yang relevan dan dapat memverifikasi pengirimnya (ADM-030).
- Sebagai Admin Situs, saya ingin backup terjadwal database dan lampiran yang restore-nya teruji otomatis agar instance self-hosted dapat pulih dari bencana (ADM-035, ADM-036).
- Sebagai Admin Situs, saya ingin melihat pemakaian terhadap batas objek dan membersihkan field atau proyek usang lewat Site optimizer agar performa situs tetap terjaga (ADM-041, ADM-042, ADM-043).
- Sebagai Admin Situs, saya ingin membuat sandbox dari konfigurasi produksi agar perubahan workflow dan app dapat diuji tanpa mengganggu tim (ADM-049).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 7 |
| P1 | 16 |
| P2 | 14 |
| P3 | 21 |
| **Total** | **58** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 6 |
| Fase 2 | 6 |
| Fase 3 | 19 |
| Fase 4 | 23 |
| Fase — (dc-only/deprecated, rujukan) | 4 |

Strategi urutan implementasi: Fase 1 hanya membangun fondasi yang dibutuhkan semua modul lain, yaitu registry halaman admin dan menu Settings (ADM-001), pemeriksaan peran admin di server (ADM-002), registry `SiteSetting` dengan halaman General configuration minimal (ADM-003), pengaturan lampiran (ADM-011), email keluar (ADM-012, UI admin untuk NTF-024 yang dibutuhkan undangan dan reset password), dan backup database plus object storage yang restore-nya teruji (ADM-035), karena data tim harian tidak boleh berjalan tanpa cadangan. Fase 2 menambah toggle perilaku global, i18n, kategori dan trash proyek, parallel sprints, serta batas objek, seiring hadirnya JQL, email, dan board lanjutan. Fase 3 mengisi konsol admin penuh: tampilan, banner, default preferences, Advanced settings, mail handler, audit log, ketiga helper diagnosis, shared items, arsip proyek premium (ADM-024, selaras PRJ-032), webhook, backup manager, impor eksternal, dan kuota plan; audit log (ADM-014) sebaiknya dibangun paling awal di Fase 3 karena hampir semua halaman lain menulis ke sana. Fase 4 berisi fitur premium/enterprise dan platform (Plans, Teams, Goals, DevOps, ekstensi, sandbox, release tracks, kebijakan data, AI, langganan) yang sebagian besar dimodelkan sebagai entitlement/feature flag VISI-003 tanpa layanan eksternal. Fitur Data Center-only (ADM-010, ADM-040, ADM-057, ADM-058) hanya menjadi rujukan; subset yang berguna bagi instance self-hosted diserap ke modul 27.

## Fitur

## Navigasi admin & model akses

### ADM-001 Jira settings navigation (Settings menu / admin hub)

- **Deskripsi:** Pintu masuk administrasi global lewat ikon gear (Settings) di top navigation yang mengelompokkan halaman admin per kategori dan menyesuaikan isi menu dengan peran pengguna. Semua halaman admin didaftarkan di satu registry yang dapat disaring per permission.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira), Admin Organisasi, Admin Akses Pengguna, Admin Proyek (hanya tautan proyek yang ia kelola); pengguna biasa hanya melihat Personal settings.
- **Aturan perilaku (acceptance criteria):**
  1. Ikon Settings (gear) tampil di top navigation untuk semua pengguna login; isi menu dihitung dari peran: Jira admin melihat System, Apps, Projects/Spaces, dan Work items; Admin Organisasi melihat tautan User management dan Billing; Admin Akses Pengguna hanya User management; Admin App hanya Apps; Admin Proyek hanya Projects (proyek yang ia kelola); pengguna tanpa peran admin hanya melihat Personal settings.
  2. Menu terbagi dua grup berlabel "Administrasi organisasi" (User management, Billing) dan "Pengaturan produk" (System, Apps, Projects/Spaces, Work items, Apps/ekstensi terpasang); grup tanpa item yang boleh diakses tidak dirender.
  3. Tautan "Personal settings" tampil terpisah di menu yang sama, di luar kedua grup admin, dan selalu tersedia bagi pengguna login.
  4. ⚠ Setiap kategori memiliki sidebar kiri dengan grup sub-halaman; untuk System susunannya: General configuration; Security (Global permissions, Project roles, Issue security); User interface (Default user preferences, System dashboard, Look and feel, Announcement banner); Import and export (Backup manager, External system import); Mail (Global mail settings, Incoming mail); Admin helper (Permission helper, Notification helper); Shared items (Filters, Dashboards); Advanced (Advanced settings, Attachments, Events, WebHooks, Issue collectors); Auditing (Audit log); Automation (global rules); Site optimizer. Susunan disimpan sebagai data registry `admin_pages (key, category, group, order, min_role, route, keywords)`, bukan di-hardcode di SPA.
  5. Halaman yang dimiliki modul lain (Global permissions di modul 03, Events di modul 19, global automation di modul 20, fields/screens/workflows/schemes di modul 05 dan 09) hanya didaftarkan sebagai entri registry yang menunjuk rute modul pemiliknya; modul 22 tidak menduplikasi logikanya.
  6. Setiap halaman admin memiliki deep link stabil berpola `/settings/<kategori>/<halaman>` (mis. `/settings/system/general-configuration`) dan breadcrumb "Settings > <kategori> > <halaman>"; URL lama berpola `/secure/admin/*.jspa` dan `/jira/settings/...` dialihkan 301 ke padanannya bila ada, atau ke halaman kategori bila tidak.
  7. ⚠ Sidebar Settings memiliki kotak pencarian halaman admin dan shortcut keyboard `g g` dari mana pun untuk membuka dialog pencarian admin; pencarian mencocokkan judul dan kata kunci registry, hanya menampilkan halaman yang boleh diakses pengguna, dan menampilkan "Tidak ada halaman admin yang cocok" bila kosong.
  8. Non-admin yang membuka URL admin secara langsung mendapat halaman "Anda tidak memiliki izin untuk membuka halaman ini" dengan tautan kembali ke beranda; endpoint API admin yang sama mengembalikan 403 dengan kode `admin_permission_required`, dan 401 bila tanpa token.
  9. Step-up auth admin (padanan websudo Data Center) tersedia sebagai opsi yang default-nya nonaktif: bila aktif, masuk ke halaman admin meminta password ulang (atau re-auth SSO), sesi admin berakhir setelah idle default 10 menit (dapat dikonfigurasi 5-60 menit), banner "Anda sedang mengakses fungsi administratif" tampil dengan tombol "Drop access" untuk keluar dari mode admin, dan opsi dapat dimatikan lewat konfigurasi server `ADMIN_STEPUP_ENABLED=false` (lazim saat SSO/SAML).
  10. Tautan keluar untuk users, groups, product access, security, billing, dan data settings menuju konsol administrasi organisasi produk ini (modul 02, rute `/admin/...`), bukan domain eksternal.
  11. ⚠ Label menu mengikuti terminologi baru (Work items, Work types, Apps, Spaces) sesuai glosarium VISI-002; label "Projects" vs "Spaces" dipilih lewat satu pengaturan terminologi situs dan tidak memengaruhi URL atau API.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk menu global; pengaturan proyek team-managed tidak muncul di Settings karena dikelola di Project settings masing-masing (modul 04).
- **Catatan & ketergantungan:** Level minimum per halaman mengacu matriks PERM-003. Registry juga dipakai pencarian admin dan pemeriksaan otorisasi middleware (satu sumber kebenaran). Repo saat ini belum memiliki konsep peran admin, sehingga perlu migrasi kolom/relasi peran (modul 02/03) sebelum menu ini bisa difilter. Struktur menu produk acuan berubah berkala; registry berbasis data memungkinkan penyesuaian tanpa rilis SPA.

### ADM-002 Administrator roles & admin access model

- **Deskripsi:** Model peran administrasi yang membatasi siapa yang boleh mengubah konfigurasi global produk dan siapa yang mengelola organisasi/situs, dengan membedakan Admin Organisasi, Admin Akses Pengguna, Admin App (Jira admin), dan Admin Proyek. Definisi permission rinci dimiliki modul 03; fitur ini mengatur batas akses ke halaman admin.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs (Administer Jira), Admin Proyek (Administer projects / role Administrator team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Pemegang global permission Administer Jira (PERM-007) dapat membuka dan mengubah semua halaman Settings > System, Work items, Projects, dan Apps; tanpa permission ini endpoint `/api/v1/admin/system/*` mengembalikan 403.
  2. Admin Organisasi mengelola billing, user, security policy, data residency, sandbox, dan release track; Admin Organisasi dapat menetapkan dirinya sebagai Jira admin, dan aksi itu tercatat di audit log.
  3. ⚠ Konsol organisasi membedakan Admin Akses Pengguna (User access admin: memberi/mencabut akses pengguna ke app) dan Admin App (App admin: mengubah konfigurasi app); produk ini menyediakan kedua peran tersebut ditambah Billing admin (hanya halaman Billing) sebagai peran granular, sedangkan peran lain di luar daftar ini tidak dibuat.
  4. Admin Proyek (Administer projects) hanya dapat mengakses Project settings proyeknya; permintaan ke endpoint konfigurasi global ditolak 403 walaupun ia admin di banyak proyek.
  5. Pada proyek team-managed, Admin Proyek (role Administrator) dapat mengelola work type, field, workflow, dan akses lokal proyek tanpa melibatkan Jira admin.
  6. Pada proyek company-managed, Admin Proyek hanya dapat mengubah sebagian konfigurasi (detail proyek, components, versions, people/role actor, dan workflow hanya bila skema tidak dibagikan ke proyek lain serta Extended project administration aktif); perubahan skema tetap hanya oleh Jira admin.
  7. Peran Data Center "Jira System Administrators" (di atas "Jira Administrators", dengan akses SMTP, listeners, services, restore, logging) tidak dibuat sebagai peran terpisah; kewenangan setara di instance self-hosted dipegang operator server lewat konfigurasi/CLI dan Admin Situs.
  8. Setiap pemberian dan pencabutan hak admin (global permission, peran organisasi, role Administrator proyek) tercatat di audit log dengan aktor, target, peran, dan waktu.
  9. ⚠ Perlindungan lockout: sistem menolak pencabutan, deaktivasi, atau penghapusan yang membuat situs tanpa satu pun pengguna aktif pemegang Administer Jira, dengan pesan "Situs harus memiliki minimal satu administrator aktif"; pemulihan darurat tersedia lewat CLI server (`jari admin recover --email`, modul 02) dan tercatat di audit log.
  10. Hak admin tidak otomatis memberi akses isi work item: Jira admin dapat mengonfigurasi proyek, tetapi melihat work item tetap mengikuti Browse projects di permission scheme dan issue security (PERM-001, PERM-044).
- **Team-managed vs company-managed:** Team-managed: konfigurasi proyek oleh role Administrator proyek tanpa Jira admin. Company-managed: skema dikelola Jira admin; Admin Proyek terbatas pada pengaturan per proyek.
- **Catatan & ketergantungan:** Matriks kewenangan dan tabel uji akses di PERM-003; peran organisasi di USR-021. Penamaan peran mengikuti definisi persona modul [02](02-pengguna-grup-organisasi.md): persona Admin Situs dengan varian Admin Organisasi, Admin Akses Pengguna, dan Admin App; Billing admin adalah peran granular tambahan (USR-021); semua modul lain merujuk istilah ini. Pemeriksaan dilakukan di middleware otorisasi (PERM-002) yang membaca `min_role` dari registry ADM-001. Migrasi repo: tabel `users` belum memiliki peran; diperlukan tabel `global_permission_grants` dan `org_role_assignments`.

## Konfigurasi umum situs

### ADM-003 General configuration — Settings

- **Deskripsi:** Bagian Settings di Settings > System > General configuration yang memuat identitas situs (judul, base URL) dan mode akses, ditampilkan read-only dan diubah lewat tombol "Edit settings".
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira); Admin Organisasi untuk akses anonim dan mode akses.
- **Aturan perilaku (acceptance criteria):**
  1. Title (application title) wajib diisi, 1-255 karakter, dan dipakai di judul tab browser, subjek/footer email, serta header; nilai default instance baru adalah nama produk ini.
  2. Base URL di instance self-hosted dapat diedit dan wajib berupa URL absolut `http(s)://` tanpa path query; nilai ini dipakai untuk semua tautan absolut (email, webhook, OAuth redirect). Pada mode hosted multi-tenant, Base URL read-only dan dikelola lewat ADM-055.
  3. Mode akses Public/Private: Private (default) berarti hanya admin yang dapat menambah pengguna; Public mengizinkan sign-up mandiri. Bila konsol organisasi (USR-012) mengatur self-signup/approved domains, pengaturan di sana yang berlaku dan halaman ini menampilkan nilainya read-only beserta tautannya.
  4. Anonymous access: toggle situs yang mengizinkan pengguna anonim membuka proyek yang Browse projects-nya diberikan ke "Anyone on the web"; default nonaktif, dan saat nonaktif grant publik diabaikan (PERM-046, USR-024).
  5. ⚠ Contact administrators form dapat diaktifkan/dinonaktifkan dengan pesan kustom (maksimum 2.000 karakter); bila aktif, halaman galat akses dan login menampilkan tautan "Hubungi administrator" yang mengirim email ke semua pemegang Administer Jira.
  6. Introduction text: teks rich-text opsional (maksimum 10.000 karakter, disanitasi) yang tampil di gadget Introduction system dashboard dan halaman login.
  7. Halaman menampilkan nilai saat ini secara read-only; tombol "Edit settings" membuka formulir, validasi dijalankan saat simpan, dan galat ditampilkan per field tanpa menyimpan sebagian.
  8. Setiap perubahan tercatat di audit log kategori System dengan nilai lama dan baru per field.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan di registry `site_settings` (ADM-009) dengan key bertipe. Setting "Public sharing of filters and dashboards" (dirujuk SRC-045 dan DSH-013) juga ditampilkan di bagian ini, default nonaktif. Repo saat ini membaca `SERVER_HOST`/`SERVER_PORT` dari env; Base URL publik menjadi setting baru yang di-seed dari env `APP_BASE_URL` saat instalasi pertama.

### ADM-004 General configuration — Options

- **Deskripsi:** Kumpulan toggle perilaku global di bagian Options General configuration yang berlaku untuk semua proyek: voting, watching, unassigned, visibilitas email, visibilitas komentar, dan batas nama/key proyek.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Allow users to vote on work items (default ON); saat OFF, field Votes dan aksi Vote disembunyikan di UI dan API vote mengembalikan 403 `voting_disabled`, sedangkan data vote lama tetap tersimpan.
  2. Allow users to watch work items (default ON); saat OFF, panel Watchers, aksi Watch, autowatch, dan notifikasi peran Watcher dinonaktifkan, sedangkan daftar watcher lama tetap tersimpan.
  3. Allow unassigned work items (default ON); saat OFF, Assignee wajib diisi saat create dan edit, serta opsi "Unassigned" hilang dari default assignee proyek dan component.
  4. Menonaktifkan "Allow unassigned" ditolak dengan pesan yang mendaftar proyek dan component yang default assignee-nya masih Unassigned, sampai semuanya diganti.
  5. ⚠ User email visibility memiliki opsi Public, Hidden, Masked (mis. `user at example dot com`), dan Logged in users only (default); preferensi visibilitas email per profil pengguna (modul 21) dapat mempersempit tetapi tidak memperluas pengaturan situs.
  6. Comment visibility bernilai "Groups & project roles" atau "Project roles only" (default) dan menentukan isi dropdown "Restrict to" saat berkomentar dan mencatat worklog (PERM-045).
  7. Exclude email recipients from email header: saat ON, email notifikasi dikirim per penerima atau dengan BCC sehingga alamat penerima lain tidak terlihat di To/CC.
  8. ⚠ Maximum project name size (default 80, rentang 2-255) dan maximum project key size (default 10, rentang 2-255) dapat diatur; nilai baru hanya divalidasi untuk proyek baru dan perubahan nama/key berikutnya (PRJ-011, PRJ-012).
  9. Pencarian pengguna di user picker dan mention mengikuti global permission Browse users and groups (PERM-008); halaman ini menampilkan catatan beserta tautan ke Global permissions, bukan toggle tersendiri.
  10. Perubahan berlaku seketika (cache setting diinvalidasi paling lambat 60 detik) untuk proyek company-managed dan team-managed tanpa restart.
- **Team-managed vs company-managed:** Tidak ada perbedaan; berlaku sama untuk keduanya.
- **Catatan & ketergantungan:** Definisi perilaku issue features di CFG-057; dampak ke work item di modul 06 dan notifikasi di modul 19. Inline edit dan JQL autocomplete selalu aktif sehingga tidak memiliki toggle (lihat ADM-010). Priority schemes selalu aktif dan dikelola di CFG-054.

### ADM-005 Internationalization & default locale settings

- **Deskripsi:** Pengaturan bahasa default situs, zona waktu default, dan bahasa indexing yang dipakai selama pengguna belum menetapkan preferensinya sendiri.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (mengatur); semua pengguna (terdampak, dapat menimpa di Personal settings).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Default language dipilih dari daftar bahasa terpasang; rilis awal menyediakan English (en) dan Bahasa Indonesia (id) dengan default instance baru English, dan bahasa lain dapat ditambahkan sebagai berkas terjemahan tanpa perubahan kode.
  2. Default user time zone dipilih dari daftar zona IANA (default `UTC`) dan berlaku untuk pengguna baru atau pengguna tanpa preferensi.
  3. Indexing language (English, Indonesian, atau "Other" tanpa stemming) menentukan konfigurasi stemming full-text search (mis. konfigurasi `english`/`simple` PostgreSQL); mengubahnya memicu re-index latar belakang (modul 13).
  4. Halaman menampilkan daftar Installed languages beserta persentase kelengkapan terjemahan UI; bahasa dengan kelengkapan di bawah 100% memakai fallback English per string.
  5. Preferensi bahasa dan time zone di Personal settings (modul 21) selalu mengalahkan default situs; header `Accept-Language` hanya dipakai untuk pengguna anonim atau sebelum login.
  6. Email notifikasi dirender dalam bahasa dan time zone penerima, bukan pelaku aksi.
  7. ⚠ Admin dapat menerjemahkan nama dan deskripsi status, work type, priority, dan resolution per bahasa terpasang (CFG-061); nilai tanpa terjemahan memakai nama asli.
  8. Time zone efektif pengguna dipakai oleh fungsi JQL relatif (`startOfDay()`, `endOfWeek()`, `now()`) dan semua tampilan tanggal; penyimpanan di database selalu UTC.
  9. Mengubah default bahasa atau time zone situs tidak menimpa pengguna yang sudah memilih preferensi sendiri; pengguna yang belum memilih langsung mengikuti default baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Repo sudah memiliki `middleware.LanguageMiddleware()` dan `constants.GetError` EN/ID untuk pesan galat backend; fitur ini memperluasnya menjadi i18n UI penuh di SPA serta urutan resolusi bahasa: preferensi pengguna → default situs → `Accept-Language` → English.

### ADM-006 Look and feel

- **Deskripsi:** Kustomisasi tampilan situs oleh admin: logo, judul situs di header, warna atau tema merek, favicon, dan format tanggal/waktu global.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Upload logo situs berformat PNG, JPG, GIF, atau SVG (disanitasi) dengan ukuran maksimum 1 MB; logo diskalakan ke tinggi header, dan tombol "Reset" mengembalikan logo default produk ini.
  2. Toggle "Show site title next to logo" menampilkan atau menyembunyikan Title (ADM-003) di samping logo.
  3. ⚠ Warna merek dibatasi pada satu brand color (header/navigasi) dan satu warna aksen tautan; sistem otomatis menurunkan warna teks dengan kontras minimum WCAG AA 4.5:1 dan menolak kombinasi yang tidak memenuhi kontras dengan pesan galat.
  4. Tombol "Pick from logo" mengekstrak warna dominan logo dan mengisikannya sebagai usulan brand color yang masih dapat diubah sebelum disimpan.
  5. ⚠ Upload favicon (ICO/PNG, maksimum 256 KB, minimal 32x32 px) tersedia, dengan fallback ke favicon default produk ini.
  6. Pengaturan warna gadget dashboard khusus admin (fitur Data Center) tidak dibangun; gadget mengikuti tema aplikasi.
  7. Date/time formats: Time format, Day format, Complete date/time format, dan Day/Month/Year format diatur dengan pola token (kompatibel konsep SimpleDateFormat, mis. `dd/MMM/yy h:mm a`) dan pratinjau langsung; pola tidak valid ditolak dengan pesan galat.
  8. Toggle "Use ISO 8601 standard in date pickers" membuat kalender date picker dimulai hari Senin dan memakai penomoran minggu ISO.
  9. Tombol "Reset to defaults" mengembalikan seluruh pengaturan look and feel (logo, warna, favicon, format) ke default setelah konfirmasi.
  10. Tema light/dark/match system tidak diatur di halaman ini; tema dipilih per pengguna di Personal settings (modul 21), dan brand color diterapkan pada kedua tema.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format tanggal juga dapat dioverride di Advanced settings (ADM-009, `jira.date.picker.*`, `jira.lf.date.*`); halaman ini adalah UI utama dan keduanya membaca key yang sama. Logo dan favicon disimpan di file storage (CMT-019). Nama produk dan merek bawaan tidak dapat diganti admin (VISI-001).

### ADM-007 Announcement banner

- **Deskripsi:** Banner pengumuman global yang tampil di bagian atas semua halaman, dikelola di Settings > System > User interface > Announcement banner.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengelola); semua pengguna dan pengguna anonim sesuai visibilitas (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle aktif/nonaktif banner; banner nonaktif tetap menyimpan isi terakhir.
  2. Isi pesan memakai rich text terbatas (bold, italic, tautan, emoji) yang disanitasi; HTML mentah tidak diterima.
  3. ⚠ Panjang pesan maksimum 1.000 karakter teks terlihat; melebihi batas ditolak dengan pesan "Pesan banner melebihi 1.000 karakter".
  4. Visibility: Public (tampil juga di halaman login dan untuk pengguna anonim) atau Private (hanya pengguna login, default).
  5. Tombol Preview menampilkan banner seperti yang akan dilihat pengguna sebelum disimpan.
  6. ⚠ Pengguna dapat menutup banner; status tutup disimpan per pengguna beserta versi isi banner, dan banner muncul kembali bila isinya diubah atau diaktifkan ulang.
  7. Banner tampil di semua halaman aplikasi (board, backlog, dashboard, work item view, settings, navigator) di atas top navigation, tidak menutupi konten, dan dapat dibaca screen reader (`role="status"`).
  8. Setiap perubahan isi, visibilitas, dan status aktif tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Announcement di help center Service Management adalah fitur terpisah (modul 30). Butir "Announcement banner" pada Admin Operational Tools (ADM-057) dipenuhi oleh fitur ini. Endpoint publik `GET /api/v1/announcement` mengembalikan banner Public tanpa autentikasi.

### ADM-008 Default user preferences

- **Deskripsi:** Nilai default preferensi untuk pengguna baru atau pengguna yang belum mengubah preferensinya, dikelola di Settings > System > User interface > Default user preferences.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengatur); semua pengguna (terdampak, dapat menimpa).
- **Aturan perilaku (acceptance criteria):**
  1. Default outgoing email format: HTML (default) atau Text.
  2. Default "Notify users of their own changes": ON atau OFF (default OFF).
  3. Default sharing untuk filter dan dashboard baru: Private (default) atau shared dengan semua pengguna login (SRC-045, DSH-010).
  4. ⚠ Number of work items displayed per page di navigator: default 50, pilihan 25/50/100.
  5. ⚠ Default keyboard shortcuts aktif (default ON) dan default autowatch (default ON) tersedia sebagai pengaturan.
  6. Pada preferensi email format dan notify own changes tersedia opsi "Apply to all existing users" yang, setelah konfirmasi, menimpa nilai semua pengguna yang ada sebagai job latar belakang dan tercatat di audit log.
  7. Pengguna tetap dapat menimpa setiap nilai di Personal settings (modul 21); default hanya berlaku bagi pengguna yang belum pernah menyimpan preferensi tersebut.
  8. ⚠ Default halaman awal setelah login dapat dipilih admin antara "For you" (default) dan "Halaman terakhir dibuka"; pengguna dapat menimpanya di preferensi personal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan preferensi personal di modul 21; urutan resolusi: preferensi pengguna → default situs → default kode. Diimplementasikan sebagai key `default.user.*` di registry `site_settings`.

### ADM-009 Advanced settings (jira.* application properties)

- **Deskripsi:** Halaman Settings > System > Advanced settings berisi daftar properti aplikasi bertipe yang dapat diedit inline untuk menyesuaikan perilaku sistem tingkat rendah.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan tabel key, nilai, dan deskripsi; nilai diedit inline, dan tombol "Revert to default" mengembalikan nilai default registry.
  2. `jira.clone.prefix`: prefix summary saat clone work item, default `CLONE -`; nilai kosong berarti tanpa prefix.
  3. `jira.date.picker.java.format`, `jira.date.picker.javascript.format`, `jira.date.time.picker.java.format`, dan `jira.date.time.picker.javascript.format` mengatur format input date picker; pasangan format server/klien divalidasi agar menghasilkan tanggal yang sama untuk contoh uji.
  4. `jira.lf.date.*` (time, day, complete, dmy) adalah key yang sama dengan Date/time formats di Look and feel (ADM-006); perubahan di salah satu halaman langsung terlihat di halaman lain.
  5. `jira.issue.actions.order`: urutan default activity/komentar `asc` atau `desc`.
  6. `jira.view.issue.links.sort.order`: urutan link work item (mis. `type, status, priority`).
  7. `jira.comment.collapsing.minimum.hidden`: jumlah minimum komentar yang dilipat (bilangan bulat ≥ 0; 0 berarti tidak pernah dilipat).
  8. ⚠ `jira.projectkey.warning` dan `jira.projectkey.description` (teks bantuan saat membuat key) tersedia; `jira.projectkey.pattern` tidak disediakan di instance mana pun karena format key tetap mengikuti PRJ-011 dan PRJ-012, dan konfigurasi pola key adalah fitur Data Center yang tidak dibangun (PRJ-068); upaya menyetel key tersebut lewat UI, API, atau berkas konfigurasi ditolak sebagai key tidak terdaftar.
  9. `jira.text.field.character.limit`: batas karakter text field multi-baris, default dan maksimum 32.767.
  10. `jira.table.cols.subtasks`: daftar kolom yang ditampilkan di tabel sub-task (mis. `issuetype, status, assignee, progress`).
  11. ⚠ `jira.newsletter.tip.delay.days` tidak disediakan karena produk ini tidak memiliki newsletter tips; key yang tidak terdaftar ditolak.
  12. `jira.search.views.default.max` (default 1.000) dan `jira.search.views.max.limit` (default 1.000) membatasi jumlah baris ekspor navigator format non-CSV (XML, Word, cetak); ekspor CSV asinkron memakai batas 10.000 work item (SRC-006 butir 2, SRC-050); kedua nilai hanya dapat diturunkan, tidak dinaikkan melewati angka NFR-004/ADM-041.
  13. Setiap key memiliki tipe (string, integer, boolean, enum, regex, format tanggal) dan validator; nilai tidak valid ditolak dengan pesan galat spesifik, mis. "Nilai harus bilangan bulat antara 0 dan 32767".
  14. Setiap perubahan tercatat di audit log dengan key, nilai lama, dan nilai baru.
  15. Di instance self-hosted, properti tambahan dapat diset lewat berkas konfigurasi server (padanan `jira-config.properties`); key dari berkas ditampilkan read-only dengan label "Diatur oleh konfigurasi server" dan mengalahkan nilai database.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Registry `site_settings (key, type, value, default, description, source)` dipakai semua halaman pengaturan di modul ini sehingga Advanced settings hanya menampilkan subset key yang ditandai `advanced=true`. Nama key `jira.*` dipertahankan untuk kompatibilitas REST (`/rest/api/3/application-properties`, modul 23). Keputusan format key proyek: tetap dan tidak dapat dikonfigurasi (PRJ-011, PRJ-012, PRJ-068 Fase —).

### ADM-010 General configuration — Data Center-only options

- **Deskripsi:** Opsi General configuration dan keamanan yang hanya ada di produk acuan versi Data Center/Server dan tidak tersedia di versi Cloud; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Operator server / Admin Situs (bila diserap).
- **Aturan perilaku (acceptance criteria):**
  1. Maximum authentication attempts sebelum CAPTCHA tidak dibangun di halaman ini; padanannya adalah kebijakan pembatasan login dan lockout di modul 01.
  2. CAPTCHA on signup (mode Public) tidak dibangun di sini; proteksi bot sign-up ditangani modul 01.
  3. Logout confirmation (Never, Cookie, Always) tidak dibangun; logout selalu langsung tanpa konfirmasi.
  4. Use gzip compression tidak menjadi setting admin; kompresi respons selalu aktif di server (modul 27).
  5. Accept remote API calls tidak menjadi toggle; REST API selalu aktif dan dikontrol lewat autentikasi, API token policy (ADM-044), dan rate limit.
  6. External user management tidak menjadi toggle; sumber identitas eksternal ditangani SSO/SCIM (modul 01/02).
  7. Enable inline edit dan Enable JQL autocomplete tidak memiliki toggle; keduanya selalu aktif.
  8. Enable user picker / Allow users to see other users digantikan global permission Browse users and groups (PERM-008).
  9. ⚠ Allowlist outbound (domain, exact URL, wildcard, regex, opsi "Allow incoming", entri otomatis dari application links, cakupan gadget eksternal/RSS/konten remote, audit) tidak dibangun sebagai halaman admin terpisah; perlindungan SSRF untuk webhook, application link, dan unfurl diwujudkan sebagai allowlist egress di konfigurasi server (modul 27, profil regulated VISI-003).
  10. Show help link dan help path tidak menjadi setting; tautan bantuan menunjuk dokumentasi produk ini yang dapat diganti lewat konfigurasi build.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; lihat "Di luar cakupan & ekosistem". Padanan IP allowlist arah masuk ada di ADM-044.

## Lampiran & email

### ADM-011 Attachment settings

- **Deskripsi:** Pengaturan global lampiran (aktif/nonaktif, ukuran maksimum, thumbnail, dukungan ZIP) di Settings > System > Advanced > Attachments yang berlaku untuk semua proyek; izin per proyek tetap diatur permission scheme atau role.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (mengatur); pengguna dengan Create attachments (terdampak).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle Enable attachments (default ON); saat OFF tombol attach, drag-drop, dan paste gambar hilang di semua proyek, endpoint upload mengembalikan 403 `attachments_disabled`, sedangkan lampiran lama tetap dapat dilihat dan diunduh.
  2. ⚠ Maximum attachment size per file dapat diatur dari 1 MB sampai batas atas server (default 100 MB, batas atas default 2 GB yang dapat diubah operator lewat konfigurasi `ATTACHMENT_MAX_BYTES_HARD`); file melebihi batas ditolak sebelum upload selesai.
  3. Toggle Enable thumbnails (default ON) membuat thumbnail untuk gambar (PNG, JPG, GIF, WebP) secara asinkron; saat OFF hanya ikon tipe file yang ditampilkan.
  4. Toggle Enable ZIP support (default ON) mengizinkan melihat daftar isi arsip ZIP di media viewer dan aksi "Download all" yang mengemas semua lampiran work item sebagai ZIP.
  5. Lokasi penyimpanan tidak diedit dari UI; operator memilih backend storage (filesystem lokal atau S3-compatible) lewat konfigurasi server, dan halaman ini menampilkan backend aktif secara read-only.
  6. Total storage lampiran, avatar, dan logo dihitung terhadap kuota tier aktif (pembanding Free 2 GB, Standard 250 GB, Premium/Enterprise tak terbatas; dapat dikonfigurasi, ADM-042), dan halaman menampilkan pemakaian saat ini.
  7. Pesan galat dibedakan dan diterjemahkan: "File melebihi batas ukuran X MB" (413), "Kuota penyimpanan situs penuh" (422 `storage_quota_exceeded`), dan "Lampiran dinonaktifkan oleh administrator" (403).
  8. ⚠ Guardrail jumlah lampiran per work item default 2.000; upload ke-2.001 ditolak dengan pesan yang merujuk batas objek (ADM-041).
  9. Data security policy (ADM-046) dapat memblokir unduhan atau ekspor lampiran pada konten terklasifikasi; tombol unduh disembunyikan dan API mengembalikan 403 dengan nama policy.
- **Team-managed vs company-managed:** Tidak ada perbedaan untuk pengaturan global; izin Create/Delete attachments diatur permission scheme (company-managed) atau role (team-managed).
- **Catatan & ketergantungan:** Perilaku upload/unduh dan pemindaian malware di CMT-011 sampai CMT-019; halaman ini adalah sumber nilai yang dibaca CMT-016. Batas ukuran juga ditegakkan di reverse proxy dan handler Gin (`MaxMultipartMemory`).

### ADM-012 Outgoing mail configuration

- **Deskripsi:** Pengaturan email keluar: aktif/nonaktif global, server SMTP, identitas pengirim, prefix subjek, pengujian, dan antrean email, di Settings > System > Mail > Global mail settings.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira); operator server untuk kredensial SMTP berbasis env.
- **Aturan perilaku (acceptance criteria):**
  1. Toggle outgoing mail global (default ON setelah SMTP valid); saat OFF semua email notifikasi berhenti dikirim, banner peringatan "Email keluar dinonaktifkan" tampil di halaman admin, dan email baru dibuang tanpa diantrekan.
  2. From address default `noreply@<host Base URL>` dan dapat diubah ke alamat valid; alamat tidak valid ditolak.
  3. Domain pengirim kustom dapat diverifikasi: sistem menampilkan rekaman DNS SPF dan DKIM yang harus dipasang, memeriksa keberadaannya, dan menandai domain "Terverifikasi" sebelum From address di domain tersebut dapat dipakai.
  4. Format nama pengirim memakai template dengan variabel `${fullname}`, default `${fullname} (<nama situs>)`.
  5. ⚠ Email subject prefix opsional (mis. `[TIM]`) ditambahkan di depan subjek semua email notifikasi.
  6. ⚠ Tombol "Send a test email" mengirim email uji ke alamat yang diisi admin dan menampilkan hasil sukses atau galat SMTP lengkap (kode dan pesan server).
  7. ⚠ Halaman Mail queue menampilkan jumlah email antre, tombol "Flush queue" untuk mengirim segera, dan Error queue berisi email gagal beserta alasan dengan aksi "Retry" dan "Delete".
  8. Konfigurasi SMTP (host, port, mode TLS none/STARTTLS/TLS, autentikasi username/password, timeout default 10 detik) dapat diatur di UI atau lewat env `SMTP_*`; nilai env mengalahkan UI dan ditampilkan read-only, dan password tidak pernah ditampilkan kembali.
  9. Pada mode hosted multi-tenant, tenant tidak dapat mengganti SMTP server platform untuk notifikasi; hanya From address terverifikasi yang dapat diatur.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pembagian kepemilikan: halaman admin dan penyimpanan pengaturan dimiliki fitur ini, sedangkan perilaku pengiriman, antrean, dan retry didefinisikan di NTF-024 Outgoing Mail Settings (modul 19); keduanya Fase 1 karena undangan pengguna dan reset password membutuhkan email keluar. Template email dan aturan penerima di modul 19; sender per proyek (PRJ-054). Pengiriman memakai job queue (modul 27) dengan retry eksponensial. Restore backup menonaktifkan outgoing mail otomatis (ADM-036).

### ADM-013 Incoming mail & mail handlers

- **Deskripsi:** Konfigurasi mail server masuk dan mail handler yang membuat work item atau menambahkan komentar dari email, di Settings > System > Mail > Incoming mail.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Tambah incoming mail server dengan tipe Google (Gmail) atau Microsoft 365 lewat OAuth 2.0, atau custom POP/IMAP (host, port, protokol, SSL/TLS, username, password).
  2. Tombol "Test connection" wajib berhasil sebelum server dapat disimpan; kegagalan menampilkan pesan galat server.
  3. Tambah mail handler dengan memilih server, folder (default INBOX, khusus IMAP), dan interval polling (minimum 1 menit, default 5 menit).
  4. Tipe handler: Create a new work item or add a comment to an existing work item; Add a comment from the non-quoted email body; Add a comment with the entire email body; Create a new work item from each email message; Add a comment before a specified marker or separator (marker diisi regex).
  5. Opsi handler: proyek dan work type target, default reporter, strip quotes, CC watchers (penerima CC yang dikenal menjadi watcher), CC assignee, bulk email handling (ignore/forward/delete untuk email auto-reply dan bulk), dan alamat forward untuk email yang tidak dapat diproses.
  6. Bila subjek email memuat key work item yang ada (mis. `ABC-123`) dan pengirim berhak berkomentar, email ditambahkan sebagai komentar ke work item itu, bukan work item baru.
  7. Pengirim yang tidak dikenal dicatat sebagai default reporter dengan alamat pengirim ditulis di deskripsi; opsi membuat pengguna baru otomatis tidak disediakan.
  8. ⚠ Handler dapat diedit, dinonaktifkan, dan dihapus; halaman menampilkan waktu proses terakhir, jumlah email diproses, dan log galat 30 hari terakhir per handler.
  9. Untuk proyek bisnis, project email address (modul 26) tersedia sebagai alternatif tanpa konfigurasi server dan tidak memerlukan Administer Jira.
  10. Lampiran email diproses sesuai Attachment settings (ADM-011); lampiran yang melebihi batas dilewati dan dicatat sebagai komentar sistem.
- **Team-managed vs company-managed:** Tidak ada perbedaan; target proyek dapat bertipe apa pun.
- **Catatan & ketergantungan:** Pembagian kepemilikan: UI admin (Settings > System > Mail > Incoming mail) dimiliki fitur ini; perilaku koneksi server dan pemrosesan email didefinisikan di NTF-031 Incoming Mail Servers dan NTF-032 Mail Handlers (modul 19) dengan prioritas dan fase yang sama (P2 / Fase 3). Email channel Service Management adalah fitur terpisah (modul 30). Polling dijalankan job scheduler (modul 27) dengan deduplikasi `Message-ID`.

## Audit & alat diagnosis admin

### ADM-014 Audit log (Jira settings)

- **Deskripsi:** Log perubahan konfigurasi dan aktivitas administratif yang dapat difilter, dicari, dan diekspor di Settings > System > Auditing > Audit log; audit log organisasi lintas app ada terpisah di konsol organisasi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira), Admin Organisasi, petugas kepatuhan dengan akses admin baca.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap entri memuat tanggal/waktu (UTC, ditampilkan sesuai time zone pembaca), author (atau "System"/nama app), IP address, summary, kategori, dan objek terdampak.
  2. Kategori minimal: projects, users & groups, permissions, workflows, fields, notifications, system, schemes, boards, dan automation.
  3. Filter rentang tanggal, multi-pilih kategori, dan pencarian teks bebas pada summary, author, dan nama objek.
  4. Baris dapat di-expand untuk melihat "Changed values" (nilai lama dan baru per atribut) dan "Associated items" (objek lain yang terkait, mis. proyek yang memakai skema).
  5. ⚠ Hasil yang sudah difilter dapat diekspor ke CSV (maksimum 100.000 baris per ekspor, diproses asinkron dengan tautan unduh).
  6. ⚠ Audit log settings: logging selalu aktif (tidak dapat dimatikan) dan retensi otomatis dapat dipilih 30, 90, 180, atau 365 hari (default 90); entri yang melewati retensi dihapus oleh job harian.
  7. REST API `GET /rest/api/3/auditing/record` (alias `/api/v1/admin/audit-records`) mendukung `offset`, `limit`, `filter`, `from`, dan `to`, hanya untuk Administer Jira.
  8. Entri tidak dapat diedit atau dihapus manual oleh siapa pun lewat UI maupun API; penghapusan hanya oleh job retensi.
  9. Coverage area dan coverage level Data Center (Off/Base/Advanced/Full), rotasi file log, dan integrasi SIEM berbasis file tidak dibangun; audit log produk ini selalu dalam cakupan penuh.
  10. Audit log organisasi lintas app (USR-045) dengan retensi lebih panjang, API, dan stream ke SIEM dikelola di modul 02; entri konfigurasi produk juga diteruskan ke sana.
  11. ⚠ Audit log tersedia di semua tier termasuk Free; tier hanya memengaruhi pilihan retensi maksimum (Free/Standard 90 hari, Premium/Enterprise 365 hari).
- **Team-managed vs company-managed:** Perubahan konfigurasi proyek team-managed (field, workflow, akses lokal) tercatat dengan kategori projects dan menyebut proyeknya.
- **Catatan & ketergantungan:** Fitur ini adalah UI dan kebijakan; definisi kategori dan penyimpanan bersama dengan USR-046 (tabel `audit_records` dengan dimensi `scope=app|org`). Perubahan isi work item tidak masuk audit log ini (history di modul 08). Semua modul menulis lewat satu use case `AuditRecorder.Record(ctx, entry)`.

### ADM-015 Admin helper: Permission helper

- **Deskripsi:** Alat diagnosis yang menjelaskan mengapa pengguna tertentu dapat atau tidak dapat melakukan suatu aksi pada work item tertentu, berdasarkan permission scheme, role, group, issue security, dan akses produk.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek untuk proyeknya.
- **Aturan perilaku (acceptance criteria):**
  1. Input terdiri dari pengguna (user picker), work item (key), dan permission (mis. Browse projects, Edit work items, Transition work items); ketiganya wajib.
  2. Output menampilkan Yes/No dan alasan rinci per grant: diberikan lewat group, project role, single user, reporter, assignee, atau application access, termasuk nama grant yang cocok.
  3. Output menunjukkan dampak issue security level (mis. "Ditolak karena security level Internal") dan akses produk (mis. "Pengguna tidak memiliki akses ke app ini").
  4. Helper dapat dibuka dari Settings > System > Admin helper > Permission helper dan dari tombol "Permission helper" di halaman permission scheme dengan proyek terisi otomatis.
  5. Setiap alasan memiliki tautan langsung ke skema, role, atau security level terkait untuk memperbaiki konfigurasi.
  6. ⚠ Proyek team-managed didukung: alasan dijelaskan dari role proyek (Administrator/Member/Viewer) dan access level proyek, bukan permission scheme.
  7. ⚠ Admin Proyek dapat memakai helper hanya untuk work item di proyek yang ia kelola; permintaan untuk proyek lain ditolak 403.
- **Team-managed vs company-managed:** Company-managed: evaluasi permission scheme dan issue security scheme. Team-managed: evaluasi role dan access level proyek.
- **Catatan & ketergantungan:** Pembagian kepemilikan: titik masuk Settings > System > Admin helper dimiliki fitur ini; logika evaluasi dan penjelasan izin didefinisikan di PERM-055 Permission Helper (modul 03), dengan prioritas yang disamakan (P1 / Fase 3). Memakai evaluator yang sama dengan penegakan runtime (PERM-001) dalam mode "explain" agar hasil helper selalu konsisten. Pemakaian helper tidak dicatat di audit log.

### ADM-016 Admin helper: Notification helper

- **Deskripsi:** Alat diagnosis yang memeriksa apakah pengguna tertentu akan menerima notifikasi untuk event tertentu pada work item tertentu, beserta alasannya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek untuk proyeknya.
- **Aturan perilaku (acceptance criteria):**
  1. Input terdiri dari pengguna, work item, dan notification event (Work item created, updated, assigned, commented, dll.).
  2. Output menyatakan menerima atau tidak, beserta peran penerima yang cocok (watcher, assignee, reporter, group, project role, user custom field, single user).
  3. Bila pengguna tidak dapat melihat work item (Browse projects atau issue security), hasilnya "Tidak menerima" dengan alasan izin, walaupun ia cocok sebagai penerima.
  4. Helper memperhitungkan preferensi personal "notify own changes" dan pengaturan notifikasi personal serta status outgoing mail global (ADM-012).
  5. Hasil memiliki tautan ke notification scheme dan event terkait.
  6. ⚠ Proyek team-managed didukung dengan menjelaskan aturan notifikasi proyek team-managed, bukan notification scheme.
- **Team-managed vs company-managed:** Company-managed: notification scheme. Team-managed: pengaturan notifikasi proyek.
- **Catatan & ketergantungan:** Pembagian kepemilikan: titik masuk Settings > System > Admin helper dimiliki fitur ini; logika resolusi penerima dan penjelasannya didefinisikan di NTF-022 Notification Helper (modul 19), dengan prioritas dan fase yang sama (P2 / Fase 3). Memakai resolver penerima yang sama dengan pengiriman notifikasi (modul 19) dalam mode "explain" tanpa mengirim apa pun.

### ADM-017 Admin helper: Field helper ('Find your field')

- **Deskripsi:** Alat diagnosis di work item view yang menjelaskan mengapa suatu field tidak terlihat, ditinjau dari screen, field configuration, field context, dan nilai kosong.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek untuk proyeknya.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Find your field" tersedia di menu Configure/admin pada work item view, create dialog, dan edit, hanya untuk admin; pengguna biasa tidak melihat aksi ini.
  2. Admin memilih field yang dicari, lalu sistem menampilkan status: tidak ada di screen, hidden di field configuration, context tidak berlaku untuk proyek/work type ini, atau kosong sehingga disembunyikan di view.
  3. Setiap status memiliki tautan perbaikan langsung (tambah ke screen, ubah context, unhide di field configuration).
  4. Alasan ditampilkan terpisah untuk create, edit, dan view karena masing-masing dapat memakai screen berbeda.
  5. ⚠ Untuk proyek team-managed, helper menjelaskan konfigurasi field per work type proyek (CFG-042) dan menautkan ke halaman konfigurasi work type proyek.
- **Team-managed vs company-managed:** Company-managed: screen, field configuration, dan context. Team-managed: konfigurasi field per work type proyek.
- **Catatan & ketergantungan:** Perilaku inti didefinisikan di CFG-051; fitur ini menambahkan entri di menu admin helper dan registry ADM-001.

## Objek bersama & default tampilan

### ADM-018 Shared items administration — Filters

- **Deskripsi:** Halaman Settings > System > Shared items > Filters tempat admin mencari filter milik siapa pun, mengubah pemiliknya, mengubah sharing, dan menghapusnya.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian filter berdasarkan nama (substring), owner (user picker), dan sharing (project, group, role, private, logged-in).
  2. Tabel menampilkan kolom nama, owner, shared with, editors, dan popularity (jumlah favorit), dengan urutan default popularity menurun.
  3. Aksi Change owner memindahkan kepemilikan ke pengguna aktif lain agar board, gadget, dan subscription yang memakai filter tetap berjalan; owner baru yang sudah punya filter bernama sama ditolak dengan pesan galat.
  4. Admin dapat mengubah Viewers dan Editors filter milik siapa pun dengan aturan yang sama seperti SRC-045.
  5. Delete filter menampilkan peringatan dependensi berisi daftar board, gadget, dan subscription yang memakai filter; konfirmasi diperlukan, dan board yang kehilangan filter ditandai rusak.
  6. Admin dapat melihat dan mengelola filter private untuk keperluan perbaikan; setiap akses ubah atas filter private orang lain tercatat di audit log.
  7. ⚠ Bulk change owner: admin memilih banyak filter lalu memindahkannya ke satu owner baru sekaligus, dengan laporan hasil per filter (sukses/gagal beserta alasan).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku inti filter di SRC-049; fitur ini memastikan halaman terdaftar di registry admin dan berbagi komponen tabel dengan ADM-019. Filter milik pengguna nonaktif ditandai "(nonaktif)" dan dapat disaring.

### ADM-019 Shared items administration — Dashboards

- **Deskripsi:** Halaman Settings > System > Shared items > Dashboards untuk mengelola semua dashboard situs, termasuk mengubah pemilik dan sharing.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian dashboard berdasarkan nama, owner, dan sharing, dengan kolom nama, owner, shared with, editors, dan popularity.
  2. Change owner dashboard ke pengguna aktif lain; validasi nama unik per owner berlaku.
  3. Admin dapat mengubah Viewers dan Editors dashboard milik siapa pun.
  4. Delete dashboard milik siapa pun dengan konfirmasi, kecuali system dashboard yang tidak menampilkan aksi Delete.
  5. Filter cepat "Owner nonaktif" menampilkan dashboard milik pengguna yang dinonaktifkan atau dihapus agar dapat dipindahkan atau dihapus.
  6. ⚠ Bulk change owner tersedia dengan laporan hasil per dashboard.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku inti di DSH-014 dan DSH-015 (modul 17).

### ADM-020 System dashboard (default dashboard)

- **Deskripsi:** Dashboard default situs yang dilihat pengguna yang belum memiliki atau belum memilih dashboard sendiri; hanya admin yang dapat mengubahnya.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengubah); semua pengguna dan pengguna anonim (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > System > User interface > System dashboard membuka system dashboard dalam edit mode untuk mengubah layout dan gadget.
  2. System dashboard tampil untuk pengguna baru dan, bila anonymous access aktif (ADM-003), untuk pengguna anonim.
  3. ⚠ Pengguna dapat menyalin system dashboard lewat aksi "Copy dashboard" sebagai titik awal dashboard baru.
  4. System dashboard tidak dapat dihapus; API delete mengembalikan 400 `system_dashboard_protected`.
  5. Perubahan system dashboard tidak memengaruhi salinan yang sudah dibuat pengguna sebelumnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi inti di DSH-009 System default dashboard (P0 / Fase 3). Perbedaan prioritas disengaja: DSH-009 (P0) mencakup keberadaan dan tampilan system dashboard dengan layout bawaan hasil seed, sehingga pengguna baru selalu punya halaman awal; fitur ini (P2) hanya menambah titik masuk admin untuk mengedit layout tersebut, yang dapat menyusul tanpa memblokir DSH-009. Keduanya tetap Fase 3.

### ADM-021 Issue navigator default columns

- **Deskripsi:** Pengaturan kolom default list view hasil pencarian untuk pengguna yang belum mengatur kolomnya sendiri.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (mengatur); semua pengguna (terdampak).
- **Aturan perilaku (acceptance criteria):**
  1. Admin menentukan set kolom default dan urutannya; default instance baru: Type, Key, Summary, Assignee, Reporter, Priority, Status, Resolution, Created, Updated, Due date.
  2. Admin dapat menambah, menghapus, dan mengurutkan kolom (drag-and-drop) dari system field dan custom field.
  3. Urutan prioritas kolom: kolom yang disimpan pada filter (oleh filter owner) → "My defaults" pengguna → default situs.
  4. Tombol "Restore defaults" mengembalikan set kolom ke default instance baru.
  5. ⚠ Halaman berada di Settings > Work items > Issue navigator default columns (rute `/settings/work-items/navigator-columns`).
- **Team-managed vs company-managed:** Tidak ada perbedaan; custom field team-managed dapat dipilih sebagai kolom.
- **Catatan & ketergantungan:** Konfigurasi kolom per pengguna dan filter di SRC-025.

## Proyek & siklus hidup data

### ADM-022 Project categories

- **Deskripsi:** Pengelompokan proyek ke dalam kategori buatan admin untuk menyaring direktori proyek, JQL, dan laporan.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (mengelola kategori); Admin Proyek (menetapkan kategori proyeknya).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > Projects > Project categories menampilkan daftar kategori dengan nama, deskripsi, dan jumlah proyek.
  2. Add category memerlukan nama unik (case-insensitive, 1-255 karakter) dan deskripsi opsional; nama duplikat ditolak dengan pesan "Kategori dengan nama ini sudah ada".
  3. Edit dan delete kategori; menghapus kategori membuat proyek terkait menjadi tanpa kategori setelah konfirmasi yang menyebut jumlah proyek terdampak.
  4. Kategori ditetapkan di Project settings > Details untuk proyek company-managed maupun team-managed.
  5. Direktori proyek dapat disaring berdasarkan kategori.
  6. JQL `category = "<nama>"` dan `category is EMPTY` dapat dipakai di filter, gadget, dan automation.
  7. Satu proyek hanya dapat memiliki satu kategori (atau tanpa kategori).
- **Team-managed vs company-managed:** Tidak ada perbedaan; keduanya dapat diberi kategori.
- **Catatan & ketergantungan:** Definisi entitas dan perilaku proyek di PRJ-019; JQL di modul 12. Tabel `project_categories (id, name, description)` dan kolom `projects.category_id`.

### ADM-023 Project trash (deleted projects)

- **Deskripsi:** Proyek yang dihapus masuk Trash selama 60 hari dan dapat dipulihkan atau dihapus permanen oleh admin.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > Projects > Trash menampilkan proyek terhapus beserta nama, key, tanggal dihapus, penghapus, dan sisa hari sebelum penghapusan permanen.
  2. Aksi Restore mengembalikan proyek beserta seluruh work item, board, konfigurasi, dan aksesnya seperti sebelum dihapus.
  3. Aksi Delete permanently tersedia sebelum masa retensi habis dan memerlukan konfirmasi dengan mengetik key proyek.
  4. Job harian menghapus permanen proyek yang sudah 60 hari di Trash.
  5. ⚠ Proyek di Trash tidak dapat diakses (URL, JQL, board, filter mengembalikan "tidak ditemukan"), dan key-nya tetap tercadang sehingga proyek baru tidak dapat memakai key yang sama sampai penghapusan permanen.
  6. Berlaku untuk proyek company-managed dan team-managed.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Siklus hidup proyek di PRJ-033; kebijakan purge di ADM-025. Restore dan delete permanen tercatat di audit log.

### ADM-024 Project archiving (admin)

- **Deskripsi:** Mengarsipkan proyek dan work item (tier Premium/Enterprise) sehingga menjadi read-only dan keluar dari pencarian default, serta memulihkannya kembali.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / premium
- **Aktor & izin:** Admin Situs (Administer Jira); Admin Proyek untuk archive work item proyeknya bila diizinkan permission Archive work items.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi Archive project tersedia dari direktori Projects dan Project settings untuk Jira admin, dengan konfirmasi.
  2. Proyek terarsip bersifat read-only, tidak muncul di hasil JQL default dan pemilih proyek, serta board dan filter yang hanya memuat proyek itu menjadi tidak aktif.
  3. Direktori proyek memiliki filter "Archived" untuk melihat proyek terarsip.
  4. Restore archived project mengembalikan seluruh data, board, dan akses seperti sebelum diarsipkan.
  5. Berlaku untuk proyek company-managed dan team-managed.
  6. Archive work item individual dan bulk archive lewat JQL tersedia; halaman Archived work items menampilkan daftar dengan aksi restore dan ekspor CSV.
  7. ⚠ Storage proyek dan work item terarsip tetap dihitung terhadap kuota situs.
  8. ⚠ Fitur hanya aktif pada tier Premium/Enterprise; setelah downgrade, data yang sudah terarsip tetap terarsip dan dapat di-restore, tetapi arsip baru tidak dapat dibuat.
  9. Archive dan restore (proyek dan work item) tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku proyek di PRJ-032 Project Archive; fase disamakan dengan PRJ-032 (Fase 3) karena halaman admin dan perilaku arsip dirilis bersama; archive massal juga dari Site optimizer (ADM-043); gating tier lewat VISI-003.

### ADM-025 Data Lifecycle: Trash, Archive & Purge

- **Deskripsi:** Kebijakan siklus hidup terpadu untuk data yang dihapus atau diarsipkan: masa tinggal di Trash, arsip yang dapat dipulihkan, dan purge permanen yang aman bagi database.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Admin Situs; Engineering (job purge).
- **Aturan perilaku (acceptance criteria):**
  1. Proyek yang dihapus masuk Trash 60 hari, dapat di-restore admin, lalu di-purge permanen otomatis (ADM-023).
  2. Archive proyek (Premium/Enterprise) membuat proyek read-only, tidak muncul di pencarian default, dan dapat di-restore (ADM-024).
  3. Archive work item (Premium/Enterprise) membuat work item tidak muncul di board dan pencarian default, dapat di-restore, dan daftar archived dapat diekspor.
  4. Custom field yang dihapus masuk Trash 60 hari dan dapat di-restore beserta nilainya sebelum purge (modul 05).
  5. ⚠ Delete work item bersifat permanen (memerlukan permission Delete work items) dan menghapus sub-task, link, lampiran, worklog, dan changelog-nya; tidak ada recycle bin work item.
  6. Job purge memproses cascade besar secara bertahap (batch, mis. 1.000 baris per transaksi) dengan progres yang dapat dilanjutkan, sehingga tidak menahan lock tabel lebih dari beberapa detik.
  7. Halaman admin menampilkan laporan storage (per proyek, lampiran, terarsip) dan riwayat purge (objek, jumlah baris, waktu, pelaku).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dipakai bersama oleh modul 04, 05, 06; retensi otomatis berbasis umur di ADM-047. Purge berjalan di job queue (modul 27).

## Pengaturan aplikasi kerja (agile, Plans, Teams, Goals)

### ADM-026 Jira configuration (Apps/Products settings: parallel sprints)

- **Deskripsi:** Halaman Settings > Apps > konfigurasi aplikasi kerja yang memuat setting agile tingkat situs seperti Parallel sprints.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle Parallel sprints (default OFF) mengizinkan lebih dari satu sprint aktif di board Scrum company-managed.
  2. Saat aktif, board menampilkan beberapa tab sprint aktif, dialog Complete sprint meminta memilih sprint, dan laporan sprint dipisah per sprint (SPR-023).
  3. ⚠ Proyek team-managed juga dapat menjalankan lebih dari satu sprint aktif bila toggle situs aktif dan fitur Sprints proyek aktif.
  4. Field Rank dipakai global untuk urutan backlog dan board tanpa UI admin.
  5. Pengaturan Rank field, LexoRank, dan Epic Link gaya Data Center tidak dibangun di halaman ini; kesehatan rank dipantau di modul 27.
- **Team-managed vs company-managed:** Toggle situs berlaku untuk keduanya; team-managed juga memerlukan fitur Sprints proyek aktif.
- **Catatan & ketergantungan:** Logika sprint di modul 11. Menonaktifkan toggle saat ada board dengan lebih dari satu sprint aktif ditolak sampai sprint tambahan diselesaikan.

### ADM-027 Plans settings (Advanced planning admin)

- **Deskripsi:** Pengaturan global untuk Plans: hierarki work type di atas Epic, link type yang dianggap dependensi, dan izin plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > Work items > Work type hierarchy menampilkan level hierarki; admin dapat menambah level di atas Epic (mis. Initiative) dan memetakan work type ke tiap level (CFG-006, CFG-007).
  2. ⚠ Aturan hierarki: sub-task selalu di bawah level story, Epic di atas level story, setiap work type hanya di satu level, dan jumlah level kustom di atas Epic maksimum 5.
  3. Dependencies settings memilih satu atau lebih link type yang diperlakukan sebagai dependensi (default `blocks / is blocked by`).
  4. ⚠ Global permission "Create plans" mengatur siapa yang dapat membuat plan; setiap plan memiliki permission view/edit per user, group, atau role.
  5. ⚠ Plans hanya aktif pada tier Premium/Enterprise; pada tier lain halaman menampilkan penjelasan plan dan tidak ada versi terbatas.
  6. Plans memakai field Team, Target start, dan Target end; halaman menampilkan status ketiga field tersebut dan tautan ke konfigurasinya.
  7. ⚠ Level hierarki di atas Epic berlaku juga untuk proyek team-managed.
  8. ⚠ Financial year: admin memilih bulan dan tanggal awal tahun keuangan di level situs (default 1 Januari); kuartal yang ditampilkan di timeline Plans (zoom Quarters, PLAN-011) dihitung dari tanggal awal tersebut, nama tahun keuangan memakai tahun berakhirnya periode (mis. awal 1 Juli 2026 menghasilkan "FY2027"), dan pengaturan berlaku untuk semua plan tanpa override per plan (PLAN-031).
- **Team-managed vs company-managed:** Hierarki situs berlaku untuk keduanya.
- **Catatan & ketergantungan:** Fitur Plans di modul 18; zoom dan navigasi timeline di PLAN-011, pengaturan per plan di PLAN-031 (tidak memuat tahun keuangan karena nilainya global); gating tier VISI-003.

### ADM-028 Teams administration (Atlassian Teams)

- **Deskripsi:** Pengaturan tingkat organisasi untuk Teams yang dipakai di field Team, Plans, dan filter board.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi / Admin Situs (mengatur); anggota team (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Izin pembuatan team dapat diatur ke "Semua pengguna" (default) atau "Admin saja".
  2. Managed team disinkronkan dari group atau IdP (SCIM); keanggotaannya read-only di UI dan hanya berubah lewat sumber sinkronisasi.
  3. Field Team tersedia di work item dan dapat dipakai di JQL (`team = <id>`), quick filter board, dan Plans.
  4. Delete/archive team menampilkan jumlah work item yang memakai team tersebut; team terarsip tetap tampil read-only di work item, sedangkan team terhapus dikosongkan dari field Team setelah konfirmasi.
  5. Team memiliki profil (nama, deskripsi, avatar, anggota) dan visibilitas bagi semua pengguna situs.
  6. ⚠ Team types tidak dibangun pada rilis awal; semua team bertipe sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas & perilaku Team: PLAN-057 ([18](18-timeline-rencana-tampilan-proyek.md)); Operations team memakai entitas yang sama (JSM-OPS-001, [34](34-jsm-operations-alerting-on-call.md)). ADM-028 hanya memiliki halaman pengaturan admin (izin pembuatan team, managed team, arsip/hapus); prioritas disamakan dengan PLAN-057 (P1). Field Team di CFG-019; sinkronisasi SCIM di USR-035. Modul 18 diminta menambahkan rujukan balik ADM-028 di Catatan PLAN-057.

### ADM-029 Goals administration

- **Deskripsi:** Pengaturan Goals yang muncul di produk ini lewat field Goals pada work item dan proyek.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi / Admin Situs (aktivasi dan izin); pengguna (membuat/mengikuti goal).
- **Aturan perilaku (acceptance criteria):**
  1. Goals diaktifkan per situs lewat toggle (default nonaktif); saat nonaktif field Goals dan menu Goals disembunyikan.
  2. ⚠ Goal dapat memiliki sub-goal (maksimum 5 tingkat) dan satu tipe goal pada rilis awal.
  3. Field Goals tersedia di work item dan proyek untuk menautkan pekerjaan ke goal.
  4. Izin membuat dan mengedit goal diatur ke "Semua pengguna" atau group tertentu.
  5. Pengikut goal menerima notifikasi saat status atau update goal berubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas Goal: PLAN-061 ([18](18-timeline-rencana-tampilan-proyek.md)); ADM-029 hanya memiliki toggle aktivasi dan izin tingkat situs, sedangkan model goal, sub-goal, dan notifikasi pengikut mengikuti PLAN-061. Modul 18 diminta menambahkan rujukan balik ADM-029 di Catatan PLAN-061.

## Integrasi & ekstensi

### ADM-030 Webhooks administration

- **Deskripsi:** Halaman Settings > System > Advanced > WebHooks untuk membuat dan mengelola webhook tingkat admin yang mengirim payload event ke URL eksternal.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Create webhook memerlukan nama, URL `https://` (http hanya bila diizinkan konfigurasi server), status Enabled, dan secret opsional; bila secret diisi, setiap kiriman membawa header `X-Hub-Signature: sha256=<HMAC>`.
  2. Events dapat dipilih per kategori: work item created/updated/deleted, comment, worklog, attachment, issue link, project, version, user, sprint, board, filter, dan configuration.
  3. Filter JQL opsional membatasi event work item; JQL tidak valid ditolak saat simpan.
  4. Opsi "Exclude body" mengirim hanya notifikasi tanpa payload, dan URL mendukung variabel seperti `${issue.key}` dan `${project.id}`.
  5. Webhook dapat diedit, dinonaktifkan, dan dihapus; daftar menampilkan nama, URL, status, event, dan waktu update terakhir.
  6. Retry, timeout, dan retensi log mengikuti API-066 butir 5, 15, 18; halaman ini menampilkan log pengiriman dan tombol redeliver dari API-066 tanpa aturan sendiri.
  7. ⚠ Jumlah webhook admin maksimum 100 per situs; pembuatan webhook ke-101 ditolak dengan pesan batas tercapai.
  8. Webhook juga dapat didaftarkan lewat REST API (dynamic webhook) dan oleh ekstensi; halaman ini menampilkan webhook admin dan menandai sumber selain admin secara read-only.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat API-064 ([23](23-api-publik-webhook.md)). Kepemilikan: API-064/API-066 memiliki model webhook, payload, delivery, retry, dan keamanan; ADM-030 hanya memiliki UI admin di Settings > System > Advanced > WebHooks. Modul 23 diminta menambahkan 'ADM-030' di Catatan API-064. Guard SSRF menolak URL ke alamat privat kecuali diizinkan konfigurasi (modul 27).

### ADM-031 Development tools, Toolchain & DVCS admin

- **Deskripsi:** Halaman admin untuk menghubungkan alat pengembangan (source code, CI/CD, deployment, feature flag, security) ke produk ini, termasuk akun DVCS, OAuth credentials, dan Toolchain.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs; Admin Proyek untuk Toolchain proyeknya.
- **Aturan perilaku (acceptance criteria):**
  1. DVCS accounts: admin menambah akun GitHub, GitHub Enterprise, GitLab, dan Bitbucket lewat integrasi native modul 24, dengan status koneksi per akun.
  2. Sinkronisasi repositori dapat dipicu manual, repositori baru dapat di-auto-link, dan smart commits dapat diaktifkan/dinonaktifkan per repositori.
  3. OAuth credentials (Apps > Integrations) membuat client ID/secret untuk sistem CI/CD (mis. Jenkins) yang mengirim data build/deployment; secret hanya ditampilkan sekali saat dibuat.
  4. Toolchain menampilkan alat tersambung per kategori: Source code, CI/CD, Deployment, Feature flags, Security, Documentation.
  5. Development panel, deployment, dan feature flag di work item terisi dari integrasi tersebut (modul 24).
  6. Security in Jira menyediakan kerangka native untuk menyambungkan security provider; provider-nya sendiri pihak ketiga.
  7. Disconnect atau remove tool menampilkan dampaknya (data dev yang sudah tersinkron disembunyikan dari work item) dan memerlukan konfirmasi.
  8. Toolchain juga tersedia di Project settings proyek team-managed dan company-managed (PRJ-058).
- **Team-managed vs company-managed:** Tidak ada perbedaan; Toolchain proyek tersedia di keduanya.
- **Catatan & ketergantungan:** Integrasi rinci di modul 24; provider pihak ketiga (Snyk, LaunchDarkly, dll.) di luar cakupan.

### ADM-032 Application links

- **Deskripsi:** Pengelolaan application links yang menghubungkan produk ini dengan instance lain atau aplikasi eksternal lewat OAuth.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar application links menampilkan nama, URL, tipe, dan status (connected/broken).
  2. Create link: admin memasukkan URL aplikasi, tipe dideteksi otomatis bila mungkin, lalu mengonfigurasi arah incoming dan outgoing.
  3. Autentikasi memakai OAuth 2.0 (atau OAuth 1.0a untuk kompatibilitas, dengan/tanpa impersonation), dan halaman menampilkan status konfigurasi per arah.
  4. Link dapat diedit, dijadikan primary untuk tipenya, dan dihapus dengan konfirmasi.
  5. App dalam satu instance produk ini tersambung otomatis tanpa application link.
  6. Application link dipakai untuk remote link ke halaman wiki eksternal, dev info dari server Git self-hosted, dan gadget lintas aplikasi.
  7. Link rusak menampilkan pesan troubleshooting (URL tidak terjangkau, kredensial ditolak, sertifikat tidak valid) dengan tombol "Test connection".
  8. Host application link otomatis ditambahkan ke allowlist egress server (padanan perilaku Data Center).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kepemilikan: INT-005 ([25](25-integrasi-impor-ekspor-ekstensi.md)) memiliki fungsi linked instance (remote link lintas instans, makro); API-014 ([23](23-api-publik-webhook.md)) memiliki autentikasi OAuth/JWT incoming dan outgoing; ADM-032 hanya memiliki halaman admin (daftar, create/edit/hapus, status, Test connection). Butir 3 dan 6 di atas mengikuti aturan rinci API-014 dan INT-005; bila berbeda, modul pemilik yang berlaku. Modul 23 diminta menambahkan rujukan ADM-032/INT-005 di Catatan API-014. Aplikasi pihak lain yang ditautkan berada di luar cakupan.

### ADM-033 Issue collectors (global list)

- **Deskripsi:** Daftar issue collector lintas proyek company-managed; issue collector adalah widget yang ditanam di situs eksternal untuk membuat work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (daftar global); Admin Proyek (collector proyeknya).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > System > Advanced > Issue collectors menampilkan semua collector beserta proyek, status, pembuat, dan jumlah work item yang dibuat.
  2. Collector dapat diaktifkan/dinonaktifkan dan dihapus; collector nonaktif menolak kiriman dengan pesan di widget.
  3. Setiap baris menautkan ke konfigurasi collector di proyek (work type, reporter, template form) (PRJ-061).
  4. Issue collector hanya tersedia untuk proyek company-managed; proyek team-managed tidak tampil di daftar.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Konfigurasi rinci di PRJ-061. Widget publik dilindungi rate limit dan CAPTCHA (modul 27).

### ADM-034 Apps management (Manage apps / UPM)

- **Deskripsi:** Halaman Settings > Apps untuk mencari, memasang, memperbarui, mengonfigurasi, dan mencopot ekstensi, termasuk melihat akses data ekstensi dan menangani permintaan ekstensi dari pengguna.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs dan Admin Organisasi (mengelola); pengguna (mengajukan permintaan).
- **Aturan perilaku (acceptance criteria):**
  1. Explore apps menampilkan katalog ekstensi (registry ekstensi produk ini) dengan filter kategori dan badge yang menandai ekstensi berjalan di dalam instance tanpa egress.
  2. Install ekstensi menampilkan layar consent berisi scopes/permission yang diminta; instalasi gagal bila admin menolak consent.
  3. Manage apps menampilkan daftar ekstensi terpasang beserta versi, vendor, status lisensi, dan tombol Configure/Get started.
  4. Ekstensi dapat diperbarui dan admin melihat catatan perubahan versi; update yang menambah scope memerlukan consent ulang.
  5. ⚠ Uninstall menampilkan peringatan kehilangan data ekstensi; aksi Disable (sementara, data tetap) juga tersedia.
  6. Info akses data per ekstensi: scopes, data egress (host tujuan), indikator berjalan di dalam instance, dan lokasi data ekstensi.
  7. ⚠ Produk ini hanya menyediakan satu model ekstensi native (modul 25); model ganda gaya Forge/Connect tidak direplikasi.
  8. Manage apps > Settings mengaktifkan instalasi ekstensi privat (dari berkas/URL) dan development mode untuk pengembang ekstensi.
  9. App requests: pengguna mengajukan permintaan ekstensi dengan alasan, admin menyetujui atau menolak, dan pemohon menerima notifikasi keputusan.
  10. Kebijakan user-installed apps menentukan apakah pengguna non-admin boleh memasang integrasi pribadi (mis. integrasi chat) tanpa persetujuan admin; default tidak boleh.
  11. Lisensi ekstensi (trial, bulanan/tahunan) ditampilkan per ekstensi; penagihan nyata berada di luar cakupan (ADM-053).
  12. Instalasi, update, disable, dan uninstall ekstensi tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model ekstensi di modul 25. Marketplace pihak ketiga dan fungsi app pihak ketiga di luar cakupan. Safe mode UPM dicatat di ADM-057.

## Backup, impor & migrasi

### ADM-035 Backup, Restore & Data Portability

- **Deskripsi:** Mekanisme cadangan dan pemulihan seluruh data instance (database dan object storage) dengan restore yang teruji, plus export/import situs dalam format terdokumentasi.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Operator server/SRE (konfigurasi infrastruktur); Admin Situs (melihat status, memicu export).
- **Aturan perilaku (acceptance criteria):**
  1. Backup PostgreSQL berjalan terjadwal (`pg_dump` harian default) ditambah WAL archiving untuk point-in-time recovery, dengan retensi terkonfigurasi (default 14 hari).
  2. Object storage (lampiran, avatar, logo) dicadangkan lewat versioning bucket atau replikasi, dengan penanda waktu yang selaras dengan snapshot database sehingga restore tidak menghasilkan referensi lampiran yang hilang.
  3. Uji restore otomatis berkala (default mingguan) memulihkan backup terbaru ke lingkungan terisolasi, memverifikasi checksum dan jumlah baris tabel inti, dan memberi alert bila gagal.
  4. ⚠ Export situs penuh (database + lampiran + avatar) tersedia dalam format terdokumentasi (arsip ZIP berisi JSON per entitas dan berkas biner) dan dibatasi satu export per 48 jam per situs.
  5. Import dari backup ke instance baru mendukung pemetaan pengguna berdasarkan email, dan pengguna yang tidak cocok dibuat nonaktif.
  6. Backup otomatis dibuat sebelum setiap upgrade versi aplikasi yang menjalankan migrasi database; migrasi dibatalkan bila backup gagal.
  7. Restore satu proyek dari Trash (ADM-023) tidak memerlukan restore penuh.
  8. File backup dan export dienkripsi (AES-256) dengan kunci dari konfigurasi server; file tanpa kunci yang benar tidak dapat diimpor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail infrastruktur di modul 27; UI admin di ADM-036. Migrasi `golang-migrate` di repo (`make migration-up`) perlu dibungkus perintah upgrade yang memanggil backup terlebih dahulu.

### ADM-036 Backup manager (Jira Cloud)

- **Deskripsi:** Halaman Settings > System > Import and export > Backup manager untuk membuat dan mengunduh backup situs, baik untuk restore ke instance produk ini maupun migrasi, serta mengimpor backup.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira); Admin Organisasi untuk import yang menimpa data.
- **Aturan perilaku (acceptance criteria):**
  1. "Create backup" mengekspor data situs ke arsip ZIP dalam format export ADM-035.
  2. Opsi format kedua "Create backup for self-hosted" menghasilkan arsip yang dapat diimpor ke instance self-hosted produk ini dengan versi skema yang sama atau lebih baru.
  3. Opsi "Include attachments, avatars, and logos" (default OFF) menambahkan berkas biner, dengan peringatan ukuran perkiraan.
  4. ⚠ Backup dibatasi satu per 48 jam per situs; tombol Create dinonaktifkan dengan pesan waktu tersedia berikutnya.
  5. Halaman menampilkan progres pembuatan backup, dan pemicu menerima notifikasi in-app serta email saat selesai atau gagal.
  6. ⚠ File backup dapat diunduh; tautan unduh kedaluwarsa 14 hari setelah dibuat.
  7. Info backup terakhir menampilkan tanggal, pembuat, ukuran, dan apakah lampiran disertakan.
  8. ⚠ Import/restore backup ke situs menimpa seluruh data situs target, memerlukan konfirmasi berlapis (mengetik nama situs), menonaktifkan outgoing mail otomatis setelah restore, dan tersedia di halaman yang sama untuk Admin Organisasi.
  9. ⚠ Halaman menyatakan batas cakupan: pengaturan organisasi (konsol admin), data ekstensi pihak ketiga, dan sebagian data Service Management/Product Discovery yang tidak tercakup ditampilkan sebagai daftar eksplisit.
- **Team-managed vs company-managed:** Tidak ada perbedaan; keduanya tercakup.
- **Catatan & ketergantungan:** Format dan enkripsi di ADM-035; pembuatan berjalan di job queue (modul 27). Pembuatan dan unduhan backup tercatat di audit log.

### ADM-037 Atlassian Backup and restore (scheduled backups)

- **Deskripsi:** Layanan backup tingkat organisasi yang mendukung backup terjadwal, penyimpanan terkelola atau bucket S3-compatible milik pelanggan, serta restore ke situs yang sama atau situs lain.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Organisasi saja.
- **Aturan perilaku (acceptance criteria):**
  1. Backup dapat dijalankan manual atau terjadwal (harian/mingguan, jam dan time zone dipilih).
  2. Cakupan dipilih per situs/app, dengan atau tanpa lampiran.
  3. ⚠ Tujuan penyimpanan: storage terkelola instance atau bucket S3-compatible milik pelanggan (endpoint, bucket, kredensial diuji sebelum simpan).
  4. Restore ke situs yang sama (overwrite) atau ke situs lain/sandbox (ADM-049).
  5. ⚠ Restore sebagian tingkat proyek tersedia: admin memilih satu atau beberapa proyek dari backup untuk dipulihkan ke situs target tanpa menimpa proyek lain.
  6. Riwayat backup dan restore menampilkan status, durasi, ukuran, dan galat.
  7. ⚠ Retensi backup otomatis default 30 hari dan dapat diatur 7-365 hari.
  8. Restore yang menimpa data memerlukan peringatan dan konfirmasi berlapis (ringkasan dampak, mengetik nama situs target).
  9. Hanya Admin Organisasi yang dapat menjalankan backup dan restore; peran lain mendapat 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Membangun di atas ADM-035; gating tier Premium/Enterprise (VISI-003).

### ADM-038 External system import (admin entry point)

- **Deskripsi:** Halaman Settings > System > Import and export > External system import untuk mengimpor data dari CSV, JSON, dan alat lain ke proyek; mapping importer rinci di modul 25.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. CSV import admin: upload berkas, pilih delimiter, encoding, dan format tanggal, mapping kolom ke field, mapping nilai, lalu buat proyek baru atau pilih proyek yang ada.
  2. Konfigurasi mapping dapat disimpan sebagai configuration file dan dimuat ulang untuk import berikutnya.
  3. JSON import (format importer terdokumentasi) mendukung work item, link, history, worklog, dan pengguna.
  4. ⚠ Import dari alat lain tersedia untuk Trello, Asana, dan GitHub Issues pada rilis awal; sumber lain ditambahkan sebagai importer modul 25.
  5. ⚠ Importer legacy (Bugzilla, Mantis, Redmine, Trac, FogBugz, Pivotal Tracker) tidak dibangun; data tersebut diimpor lewat CSV/JSON.
  6. Validasi awal (dry run) menampilkan galat per baris sebelum import, dan laporan hasil memuat jumlah sukses/gagal serta log yang dapat diunduh.
  7. Pengguna dipetakan berdasarkan email, dan custom field baru dapat dibuat selama import.
  8. ⚠ Peringatan ditampilkan bila proyek target team-managed (field yang tidak didukung dilewati), dan berkas dibatasi 100 MB / 250.000 baris per import.
  9. Non-admin dapat memakai CSV import dari navigator untuk proyek yang ia berhak create (modul 13); halaman admin ini memiliki opsi lebih lengkap (proyek baru, pengguna, history).
  10. Import dari instance lain produk ini atau situs acuan dilakukan lewat backup/JSON (ADM-036) atau alat migrasi (ADM-039).
- **Team-managed vs company-managed:** Target team-managed memiliki dukungan field terbatas dan menampilkan peringatan.
- **Catatan & ketergantungan:** Importer pihak ketiga dari Marketplace di luar cakupan. Import berjalan di job queue dengan progres.

### ADM-039 Cloud migration tools (JCMA & cloud-to-cloud)

- **Deskripsi:** Alat migrasi untuk memindahkan proyek, pengguna, dan konfigurasi dari instance self-hosted ke hosted atau antar situs; bagi produk ini setara dengan "import dari instance lain" lewat export/import JSON.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Asisten migrasi di instance sumber memilih proyek, pengguna dan group, serta lampiran, menjalankan pre-migration checks, dan menghasilkan error report.
  2. Migrasi dapat dilakukan bertahap per proyek; konfigurasi bernama sama (workflow, field, skema) di tujuan ditangani dengan pilihan "gunakan yang ada" atau "buat salinan dengan sufiks".
  3. Migrasi antar situs (situs-ke-situs) di konsol organisasi memilih situs sumber/tujuan, proyek, dan data (work item, konfigurasi, pengguna).
  4. Pre-flight check menampilkan estimasi durasi dan volume, dan laporan hasil memuat jumlah objek per tipe yang berhasil atau gagal.
  5. App assessment menampilkan ekstensi terpasang di sumber dan ketersediaannya di tujuan; data ekstensi tidak dimigrasikan.
  6. Pengguna di-merge berdasarkan email; pengguna tanpa padanan dibuat nonaktif di tujuan.
  7. ⚠ Laporan cakupan menyatakan objek yang tidak dimigrasikan (filter dan dashboard yang dibagikan ke objek di luar cakupan, data ekstensi).
- **Team-managed vs company-managed:** Tidak ada perbedaan; konfigurasi team-managed ikut proyeknya.
- **Catatan & ketergantungan:** Format data sama dengan ADM-035; importer dari produk acuan (JSON/CSV) di modul 25. Aplikasi migrasi resmi produk acuan di luar cakupan.

### ADM-040 Backup & restore system (Data Center XML)

- **Deskripsi:** Fitur Data Center untuk membuat backup XML dan memulihkan seluruh data (atau satu proyek) dari berkas backup dengan menimpa data yang ada; dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Operator server.
- **Aturan perilaku (acceptance criteria):**
  1. Backup XML ke folder export dengan service terjadwal tidak dibangun; digantikan backup terjadwal ADM-035.
  2. Restore dari folder import server tidak dibangun; restore lewat ADM-036/ADM-037 atau CLI server.
  3. Opsi "Restore attachments" dan path lampiran digantikan opsi lampiran pada ADM-036.
  4. Penanganan license saat restore tidak relevan karena produk ini tidak memakai license key.
  5. Peringatan bahwa seluruh data akan dihapus tetap diterapkan pada restore ADM-036.
  6. Outgoing mail dinonaktifkan otomatis setelah restore diterapkan pada ADM-036 butir 8.
  7. Progres dan log restore diterapkan pada ADM-036/ADM-037.
  8. Project import dari berkas backup dipenuhi restore sebagian ADM-037 butir 5 dan PRJ-067.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; format XML Data Center tidak didukung sebagai format impor rilis awal.

## Batas, kapasitas & kesehatan situs

### ADM-041 Object limits & guardrails

- **Deskripsi:** Batas teknis yang menjaga performa: jumlah lampiran, komentar, link, dan worklog per work item, panjang teks, serta batas konfigurasi dan API, dengan perilaku penolakan yang konsisten.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Platform (menegakkan); Admin Situs (melihat dan, di self-hosted, mengubah batas lunak).
- **Aturan perilaku (acceptance criteria):**
  1. Text field multi-baris dibatasi 32.767 karakter; single-line text dan summary dibatasi 255 karakter; input lebih panjang ditolak 400 dengan nama field.
  2. ⚠ Batas per work item mengikuti NFR-004 butir 5 sebagai satu-satunya sumber angka: lampiran 2.000, komentar 5.000, worklog 10.000, dan issue link 2.000 (rincian domain di ISS-044).
  3. ⚠ Child work item/sub-task per parent dibatasi 1.000 (NFR-004 butir 5, ISS-044).
  4. ⚠ Custom field dibatasi 700 per situs dan 50 work type per proyek (hard limit), dengan peringatan rekomendasi pada 80% batas.
  5. ⚠ Proyek team-managed dibatasi 150 field, 100 work type, dan 150 status per proyek.
  6. Ekspor navigator: CSV asinkron dibatasi 10.000 work item dan format lain (XML, Word, cetak) 1.000 work item per ekspor (SRC-006 butir 2, SRC-050, ADM-009 butir 12); REST pagination `maxResults` maksimum 100 (5.000 untuk endpoint tertentu yang didokumentasikan), dan bulk operation 1.000 work item per operasi.
  7. ⚠ Board dan backlog menampilkan maksimum 5.000 work item; melebihi batas menampilkan pesan agar filter board dipersempit.
  8. Automation rule dibatasi 65 komponen per rule, dan REST API dikenai rate limit per pengguna/token dengan header `Retry-After` (modul 23).
  9. Saat batas tercapai aksi ditolak dengan pesan galat yang menyebut batas dan nilainya, error code stabil sesuai katalog galat NFR-004 butir 12 (NFR-021 butir 2), tautan dokumentasi, dan tidak ada penulisan sebagian.
- **Team-managed vs company-managed:** Batas field/work type/status khusus team-managed (butir 5).
- **Catatan & ketergantungan:** NFR-004 Limits & Guardrails (modul 27) adalah satu-satunya sumber angka guardrail; fitur ini dan ISS-044 Work Item Limits & Guardrails hanya menyalin nilainya, dan perubahan angka dilakukan di NFR-004 terlebih dahulu. Semua angka disimpan di satu tabel konfigurasi batas (dapat diubah operator self-hosted) dan dibaca validator modul terkait; pemantauan dan pembersihan di ADM-043.

### ADM-042 Usage, storage & plan limits

- **Deskripsi:** Informasi pemakaian dan batas plan yang ditegakkan sistem: jumlah pengguna, storage, eksekusi automation, dan fitur yang dikunci per tier.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs (melihat dan mengatur kuota self-hosted).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ User limit per tier dapat dikonfigurasi dengan nilai pembanding Free maksimum 10 pengguna; tier lain default tak terbatas di instance self-hosted.
  2. Storage per tier dapat dikonfigurasi (pembanding Free 2 GB, Standard 250 GB, Premium/Enterprise tak terbatas) dengan indikator pemakaian per kategori (lampiran, avatar, backup).
  3. ⚠ Eksekusi automation rule per bulan dapat dikonfigurasi per tier (pembanding Free 100, Standard 1.700, Premium/Enterprise tak terbatas), dan halaman menampilkan pemakaian bulan berjalan.
  4. Pemetaan fitur yang dikunci per tier ditampilkan (mis. Free tanpa project role kustom dan issue security; archiving, Plans, dan sandbox hanya Premium+), bersumber dari konfigurasi VISI-003.
  5. Peringatan in-app dan email ke admin pada 80% dan 100% batas; saat batas tercapai aksi terkait ditolak (mis. tidak dapat menambah pengguna) dengan pesan yang menyebut batas.
  6. Fitur terkunci menampilkan prompt penjelasan plan yang dapat dinonaktifkan oleh Admin Situs pada instance self-hosted.
  7. Admin insights (Premium/Enterprise) menampilkan pengguna aktif 30 hari, pemakaian per app, dan insight keamanan dasar (pengguna admin, token API aktif).
  8. ⚠ Analytics data lake (Enterprise) tidak dibangun pada rilis awal; sebagai gantinya tersedia ekspor data terjadwal ke database baca terpisah (keputusan terbuka).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penegakan batas pengguna di USR-020; tier dan flag di VISI-003; kuota storage di CMT-019.

### ADM-043 Site optimizer

- **Deskripsi:** Halaman admin terpusat di Settings > System > Site optimizer untuk melihat kesehatan situs dan membersihkan data lama atau tidak terpakai sebelum mencapai batas objek, mencakup work item, custom field, opsi field, skema, role, permission grant, security level, proyek, work type, component, priority, release, status, dan workflow.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs (Administer Jira), Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Dashboard kesehatan menampilkan jumlah per jenis objek dibandingkan batas rekomendasi dan batas situs (ADM-041), dengan indikator hijau/kuning (≥80%)/merah (≥100%).
  2. Work items: arsip massal work item berdasarkan umur atau last updated (mis. tidak diubah 2 tahun), dengan pratinjau jumlah sebelum eksekusi.
  3. Custom fields per proyek: menemukan field yang tidak dipakai di proyek, context, atau screen mana pun, atau tidak terisi selama 2 tahun, lalu "Review and clean up"; field buatan ekstensi dikecualikan, dan field yang dihapus masuk trash 60 hari.
  4. Field options per field: menemukan dan menghapus opsi select yang tidak dipakai work item mana pun.
  5. Optimasi field configuration scheme per proyek, termasuk "Default Field Configuration Scheme" implisit yang ditampilkan tetapi tidak dapat diedit atau dihapus.
  6. Project roles per pengguna per proyek: meninjau dan mengurangi jumlah role per pengguna.
  7. Permission grants per permission: menemukan dan menghapus grant redundan (mis. grant ke group yang juga sudah tercakup role yang sama).
  8. Project role actors per proyek: memantau dan membersihkan actor yang melewati batas rekomendasi.
  9. Security levels per skema: menghapus security level yang tidak dipakai work item mana pun.
  10. Proyek: menyaring berdasarkan proyek dan last updated atau memilih satu per satu, lalu Archive (Premium/Enterprise) atau Trash secara massal.
  11. Work types per proyek: melepaskan work type yang tidak dipakai dari skema.
  12. Components dan releases/versions per proyek: memantau jumlah, lalu menghapus atau mengarsipkan yang tidak dipakai.
  13. Priorities: meninjau jumlah priority per proyek terhadap batas.
  14. Statuses per workflow dan workflows per proyek: mendeteksi workflow kompleks atau dengan status terlalu banyak dan menautkannya ke editor workflow.
  15. Melepas field dari skema tidak menghapus data, tetapi nilainya tidak dapat diakses sampai field dikaitkan lagi; sebelum setiap cleanup ditampilkan saran membuat backup (ADM-036).
  16. ⚠ Setiap aksi cleanup tercatat di audit log dengan jumlah dan daftar objek terdampak.
  17. ⚠ Dashboard kesehatan dan pembersihan custom field tersedia di semua tier; aksi massal (archive work item/proyek, cleanup lintas jenis objek) hanya Premium/Enterprise.
- **Team-managed vs company-managed:** Pembersihan skema dan permission grant hanya untuk company-managed; pembersihan proyek dan work item berlaku untuk keduanya.
- **Catatan & ketergantungan:** Perluasan dari Custom field optimizer (CFG-040) dan ADM-041. Semua cleanup berjalan sebagai job asinkron dengan progres.

## Keamanan & tata kelola data

### ADM-044 Security & access policies (org-level, terlihat di Jira)

- **Deskripsi:** Kebijakan keamanan organisasi yang memengaruhi akses ke produk ini: IP allowlist, kebijakan aplikasi mobile, durasi sesi, kebijakan API token, enkripsi kunci milik pelanggan, dan kebijakan akses anonim. Halaman di Settings menampilkan ringkasan read-only beserta tautan ke konsol organisasi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi (mengatur); Admin Situs (melihat ringkasan).
- **Aturan perilaku (acceptance criteria):**
  1. IP allowlist (Premium/Enterprise) membatasi akses UI dan REST API berdasarkan IP/CIDR; permintaan dari luar daftar ditolak 403 `ip_not_allowed`, dan admin yang menyimpan aturan yang mengecualikan IP-nya sendiri mendapat peringatan lockout yang harus dikonfirmasi (USR-038).
  2. Mobile app policy dapat memblokir screenshot, mewajibkan passcode perangkat, dan melarang copy-paste di aplikasi mobile (USR-039).
  3. Session duration (default 30 hari) dan idle timeout (default nonaktif, dapat diatur 15 menit-30 hari) berlaku untuk sesi web; sesi yang melewati batas diminta login ulang (modul 01).
  4. User API token policy dapat memblokir pembuatan API token atau mewajibkan masa berlaku maksimum (mis. 365 hari); token yang melanggar policy ditolak 401.
  5. BYOK encryption (Enterprise) memakai kunci KMS milik pelanggan untuk enkripsi data at-rest; pencabutan kunci membuat situs tidak dapat diakses dengan halaman pemberitahuan (USR-043).
  6. Kebijakan anonymous access dan proyek publik di tingkat organisasi dapat mematikan anonymous access situs (ADM-003) sehingga toggle situs menjadi read-only.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi rinci di modul 02 (USR-038, USR-039, USR-043) dan modul 01 (sesi, token); fitur ini memastikan kebijakan terlihat di Settings dan ditegakkan di middleware. IP allowlist juga tersedia di Premium.

### ADM-045 Data classification

- **Deskripsi:** Definisi level klasifikasi data dan penerapannya pada proyek, work item, dan field sebagai dasar kebijakan keamanan (Enterprise / add-on keamanan).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi (definisi level); Admin Proyek (default proyek); pengguna berizin edit (per work item).
- **Aturan perilaku (acceptance criteria):**
  1. Classification level memiliki nama, warna, deskripsi, guideline, dan urutan sensitivitas; nama unik per organisasi.
  2. Level dapat di-publish (dapat dipakai) dan di-archive (tidak dapat dipilih baru, label lama tetap tampil).
  3. Default classification dapat ditetapkan per proyek; work item baru mewarisinya.
  4. ⚠ Klasifikasi dapat ditetapkan per work item dan per field teks/paragraf; klasifikasi field tidak boleh lebih rendah daripada klasifikasi work item-nya.
  5. Badge klasifikasi tampil di work item view (dan kartu board serta hasil pencarian sesuai PERM-064).
  6. ⚠ Menurunkan klasifikasi work item di bawah default proyek memerlukan konfirmasi dan alasan yang tercatat.
  7. Level klasifikasi dapat menjadi cakupan data security policy (ADM-046).
  8. Setiap perubahan definisi level dan klasifikasi objek tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi bersama dengan PERM-064 dan USR-041; fitur ini adalah halaman admin dan aturan pewarisan.

### ADM-046 Data security policies

- **Deskripsi:** Kebijakan keamanan data yang membatasi tindakan berisiko pada konten dengan cakupan seluruh workspace, proyek, atau level klasifikasi (Enterprise / add-on keamanan).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Rule App access memblokir ekstensi membaca konten pada cakupan policy.
  2. Rule Export memblokir ekspor work item (CSV, XML, Word, Excel), print, dan unduhan lampiran pada cakupan policy, dengan pesan yang menyebut nama policy.
  3. Rule Public links dan Anonymous access menonaktifkan tautan publik dan grant anonim pada cakupan policy.
  4. ⚠ Rule AI access memblokir fitur AI (ringkasan, generate, NL ke JQL) membaca konten pada cakupan policy (ADM-052).
  5. Coverage policy: seluruh workspace, proyek tertentu, atau level klasifikasi tertentu.
  6. Exemption dapat diberikan untuk ekstensi tertentu pada rule App access.
  7. Policy memiliki status aktif/nonaktif, dan setiap pembuatan, perubahan, aktivasi, dan pemblokiran tercatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penegakan di Authorizer (PERM-002, PERM-065) dan modul ekspor (SRC-058); definisi organisasi di USR-040.

### ADM-047 Data retention policies

- **Deskripsi:** Kebijakan tingkat organisasi untuk menghapus konten secara otomatis setelah periode tertentu demi kepatuhan. Produk acuan belum terkonfirmasi menyediakannya untuk app pelacak kerja, sehingga fitur ini diperlakukan sebagai kapabilitas tambahan, bukan paritas.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Organisasi, petugas kepatuhan.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Aturan retensi berbasis umur konten (mis. hapus work item yang resolved lebih dari N tahun, lampiran lebih dari N tahun, komentar lebih dari N tahun) dengan N minimum 1 tahun.
  2. ⚠ Cakupan aturan per app, situs, atau proyek.
  3. ⚠ Pengecualian (exempt/legal hold) untuk proyek atau work item tertentu menghentikan penghapusan sampai hold dicabut.
  4. ⚠ Setiap penghapusan otomatis menghasilkan entri audit log berisi aturan, jumlah, dan rentang objek.
  5. ⚠ Penghapusan oleh retensi melewati Trash (ADM-023) dan langsung masuk antrean purge bertahap (ADM-025), sedangkan objek terarsip (ADM-024) tetap tunduk pada aturan retensi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Gap paritas; pemilik produk memutuskan apakah dibangun (lihat Pertanyaan terbuka). Bergantung pada job purge ADM-025.

### ADM-048 Data residency

- **Deskripsi:** Kontrol untuk menyematkan data app di wilayah tertentu dan memindahkannya antar wilayah; bagi instance self-hosted, lokasi data ditentukan oleh tempat instance dijalankan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Pada mode hosted, lokasi data dipilih dari daftar wilayah yang dikonfigurasi operator (bukan daftar tetap); instance self-hosted menampilkan "Ditentukan oleh infrastruktur Anda".
  2. Halaman Data residency menampilkan lokasi saat ini per app.
  3. Request move menjadwalkan window migrasi, dan situs menjadi read-only selama perpindahan dengan banner pemberitahuan.
  4. Halaman menjelaskan data in-scope (work item, lampiran, komentar) dan out-of-scope (cache, log, analytics).
  5. Data residency ekstensi ditampilkan untuk ekstensi yang mendeklarasikan lokasi penyimpanannya.
  6. Riwayat dan status permintaan pemindahan (dijadwalkan, berjalan, selesai, gagal) ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Definisi organisasi di USR-042; arsitektur multi-region di modul 27.

## Lingkungan, rilis, preview & AI

### ADM-049 Sandbox management

- **Deskripsi:** Lingkungan sandbox (Premium/Enterprise) untuk menguji perubahan konfigurasi, ekstensi, dan rilis sebelum diterapkan ke produksi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / premium
- **Aktor & izin:** Admin Organisasi, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Create sandbox per app yang terhubung ke situs produksi, dengan URL dan database terisolasi.
  2. ⚠ Copy data produksi ke sandbox tersedia dalam mode "hanya konfigurasi", "konfigurasi + proyek terpilih", atau "seluruh data"; email keluar sandbox nonaktif default.
  3. Sandbox dapat di-reset (menyalin ulang dari produksi) dan dihapus dengan konfirmasi.
  4. Akses sandbox diberikan ke admin atau pengguna tertentu tanpa memakai kuota pengguna produksi.
  5. Ekstensi dan release track Preview (ADM-050) dapat diuji di sandbox.
  6. ⚠ Jumlah sandbox dibatasi 1 per app untuk Premium dan 5 per app untuk Enterprise.
  7. ⚠ Promote konfigurasi dari sandbox ke produksi: admin memilih objek konfigurasi (konfigurasi proyek, workflow, custom field, skema, automation rule), melihat diff sandbox vs produksi, sistem mendeteksi konflik dan dependency, eksekusi berjalan asinkron dengan progres dan log, kegagalan sebagian dilaporkan per objek, dan seluruh promote tercatat di audit log.
- **Team-managed vs company-managed:** Konfigurasi team-managed dipromosikan sebagai bagian dari proyeknya.
- **Catatan & ketergantungan:** Isolasi data dan notifikasi mengikuti VISI-003 butir 14. Alat promosi konfigurasi pihak ketiga di luar cakupan.

### ADM-050 Release tracks

- **Deskripsi:** Pengaturan (Premium/Enterprise) untuk mengendalikan kapan perubahan fitur dirilis ke situs dengan track Continuous, Bundled, atau Preview; bagi produk ini setara feature flag/rollout ring.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / premium
- **Aktor & izin:** Admin Organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Release track dipilih per app: Continuous (default, fitur baru aktif otomatis), Bundled (rilis berkala dalam batch), atau Preview (hanya untuk sandbox).
  2. Halaman menampilkan daftar perubahan mendatang beserta release notes.
  3. Aksi "Release now" merilis batch perubahan sebelum jadwal setelah konfirmasi.
  4. Admin menerima notifikasi email tentang batch berikutnya (minimal 7 hari sebelum jadwal).
  5. Perbaikan keamanan dan bug kritis selalu dirilis segera tanpa menunggu batch.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Diimplementasikan di atas layanan feature flag VISI-003 (butir 13) sebagai grup flag per batch.

### ADM-051 Feature previews & opt-in toggles

- **Deskripsi:** Mekanisme opt-in sementara untuk pengalaman baru di tingkat pengguna maupun admin sebelum menjadi default.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (toggle situs); pengguna (toggle personal).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna dapat mengaktifkan atau menonaktifkan pengalaman baru tertentu (mis. navigasi baru) lewat toggle di menu profil; pilihan disimpan per pengguna.
  2. Admin dapat mengaktifkan fitur beta tingkat situs (mis. fitur AI beta) di halaman Settings > System > Feature previews.
  3. Setiap preview memiliki masa transisi dengan tanggal wajib yang ditampilkan; setelah tanggal itu toggle hilang dan fitur menjadi default.
  4. Saat pengguna opt-out, form feedback singkat opsional ditampilkan.
  5. ⚠ Halaman Feature previews menampilkan daftar preview yang tersedia saat ini (dari registry flag bertipe `preview`); fitur tanpa entri registry tidak memiliki toggle.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai layanan flag VISI-003. Dark features tersembunyi gaya Data Center dicatat di ADM-057.

### ADM-052 Atlassian Intelligence & Rovo settings

- **Deskripsi:** Pengaturan AI untuk mengaktifkan fitur AI produk ini per situs/app, fitur beta AI, asisten dan agent, kontrol pemakaian, serta pemantauan kuota; produk ini memakai penyedia LLM yang dikonfigurasi operator, bukan layanan AI produk acuan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs (mengatur); pengguna (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle AI per situs atau per app (default nonaktif); saat nonaktif semua entry point AI disembunyikan dan endpoint AI mengembalikan 403 `ai_disabled`.
  2. Fitur AI beta diaktifkan terpisah dari fitur AI umum.
  3. ⚠ Asisten/agent dan pencarian AI dapat diaktifkan; agent kustom dan connector ke sumber data eksternal dikelola di halaman yang sama, dengan connector nonaktif default.
  4. Daftar fitur AI terdampak ditampilkan: ringkasan work item/komentar, natural language ke JQL, AI work breakdown, generate deskripsi, dan AI di automation.
  5. ⚠ Akses AI dapat dibatasi per group (default semua pengguna berlisensi).
  6. ⚠ Kuota AI (credit) dapat dikonfigurasi per situs dan per pengguna berlisensi oleh operator.
  7. ⚠ Halaman pemakaian AI menampilkan konsumsi per app, per periode, dan per fitur (chat, agent, search, AI di work item, AI di automation).
  8. Saat kuota habis fitur AI dibatasi atau dinonaktifkan, dan banner/flag peringatan tampil bagi pengguna serta admin.
  9. ⚠ Admin menerima notifikasi saat pemakaian mencapai 80% dan 100% kuota.
  10. ⚠ Konsumsi AI dari automation dicatat dan dikuota terpisah dari pemakaian interaktif.
  11. ⚠ Kuota di-reset setiap awal siklus (bulanan default), dan penambahan kuota dilakukan operator lewat konfigurasi.
  12. Kebijakan data menyatakan konten tidak dipakai untuk melatih model, dan region penyedia LLM ditampilkan untuk keperluan residency.
  13. Data security policy (ADM-046) dapat memblokir AI pada konten terklasifikasi.
  14. ⚠ Ketersediaan AI tidak digating tier; hanya bergantung pada konfigurasi penyedia LLM dan toggle admin.
  15. Halaman ini adalah kontrol pusat satu-satunya untuk semua fitur AI: setiap fitur AI di modul lain aktif hanya bila diaktifkan di ADM-052 dan tidak tersedia pada profil regulated (VISI-003); pada profil regulated semua toggle di halaman ini terkunci OFF, ditampilkan read-only dengan label "Dinonaktifkan oleh profil deployment", dan tidak ada egress ke penyedia LLM.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur AI yang digating oleh halaman ini: ISS-045, EDT-030, AUTO-073, JQL-059, SRC-014, SRC-019, NAV-015, NAV-051, NAV-052, NAV-053, PERM-050, REL-020, RPT-042, WF-043, DEV-032, INT-055, JSM-AST-047, JSM-ITSM-030, dan JSM-ITSM-044. Pasangan duplikat: NAV-052 (katalog fitur AI di sisi pengguna) merujuk halaman ini sebagai tempat pengaturan; JQL-059 dan SRC-019 adalah fitur yang sama (natural language ke JQL); NAV-015 dan SRC-014 adalah fitur yang sama (Rovo Search). Fitur AI per halaman di modul 21; profil regulated deployment (VISI-003) mematikan seluruh AI. Layanan AI dan agent produk acuan di luar cakupan.

## Langganan & siklus hidup situs

### ADM-053 Billing & subscription management

- **Deskripsi:** Pengelolaan langganan di konsol organisasi: plan, siklus tagihan, pembayaran, invoice, trial, multi-instance, dan penghapusan app. Untuk distribusi open-source dimodelkan sebagai pengelolaan tier/entitlement tanpa pemrosesan pembayaran nyata.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Billing admin.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman menampilkan plan saat ini (Free/Standard/Premium/Enterprise) dan jumlah pengguna billable per app (USR-020).
  2. Upgrade/downgrade plan dan trial Premium tersedia, dengan ringkasan dampak fitur sebelum konfirmasi (rincian di ADM-054).
  3. Siklus tagihan bulanan per pengguna atau tahunan per user tier ditampilkan sebagai atribut langganan, termasuk permintaan quote tahunan, pada mode hosted.
  4. Payment method, billing address, tax ID, dan billing contacts dapat dikelola pada mode hosted lewat penyedia pembayaran yang dikonfigurasi operator; pada self-hosted bagian ini disembunyikan.
  5. Invoice dan riwayat pembayaran dapat diunduh sebagai PDF pada mode hosted.
  6. Langganan ekstensi berbayar ditampilkan dan ditagih bersama langganan utama.
  7. ⚠ Cancel subscription atau remove app menonaktifkan app, memberi masa tenggang 60 hari, lalu menghapus data permanen; admin menerima email pada awal dan 7 hari sebelum akhir masa tenggang.
  8. ⚠ Multi-instance (Enterprise): satu langganan dapat memiliki banyak situs (batas dikonfigurasi operator, default 150), pengguna unik lintas situs dihitung sekali, dan konfigurasi, data residency, serta release track diatur independen per situs; plan lain dibatasi satu situs per langganan.
  9. Pengelolaan URL situs dirujuk ke ADM-055.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model multi-tenant di modul 27; tier di VISI-003. SLA uptime dan tingkat support bersifat kontraktual dan tidak dimodelkan.

### ADM-054 Plan trials, downgrade & site lifecycle

- **Deskripsi:** Perilaku saat situs memulai trial, upgrade, downgrade, berhenti berlangganan, atau dinonaktifkan otomatis karena tidak aktif di Free plan: fitur yang terkunci, data yang menjadi read-only atau dihapus, dan jadwal penghapusan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs, Billing admin; System (scheduler).
- **Aturan perilaku (acceptance criteria):**
  1. Free trial Standard/Premium (default 30 hari) berakhir dengan konversi ke plan berbayar atau kembali ke plan semula; admin menerima pengingat 7 dan 1 hari sebelum akhir trial.
  2. ⚠ Downgrade Premium ke Standard: Plans dan level hierarki di atas Epic menjadi read-only, sandbox dihapus setelah 30 hari, release track kembali ke Continuous, dan archiving baru tidak tersedia (arsip lama tetap dapat di-restore).
  3. Downgrade ke Free memerlukan maksimal 10 pengguna dan storage 2 GB terpenuhi lebih dulu; setelahnya permission kustom, issue security, dan project role tidak lagi ditegakkan.
  4. ⚠ Konfigurasi lama saat downgrade tetap tersimpan tetapi tidak ditegakkan, dan aktif kembali utuh bila plan dinaikkan (VISI-003 butir 6).
  5. Blocking checks sebelum downgrade menampilkan daftar syarat yang belum terpenuhi (mis. jumlah pengguna, storage) dan menonaktifkan tombol konfirmasi sampai terpenuhi.
  6. ⚠ Pembatalan langganan menonaktifkan app, memberi masa tenggang 60 hari, lalu menghapus data permanen.
  7. ⚠ Pada mode hosted, situs Free tanpa aktivitas pengguna selama 120 hari dinonaktifkan otomatis setelah email peringatan ke admin (30, 7, dan 1 hari sebelumnya) yang memuat tautan aktivasi ulang; fitur ini tidak berlaku di self-hosted.
  8. Status "deactivated": app tidak dapat diakses tetapi data disimpan selama masa tenggang, lalu dihapus permanen tanpa dapat dipulihkan.
  9. Admin dapat mengaktifkan ulang app atau situs dari konsol organisasi selama masa tenggang.
  10. ⚠ Konsol organisasi > Products menampilkan status (aktif, trial, deactivated) dan tanggal deaktivasi/penghapusan terjadwal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memerlukan job terjadwal deteksi tenant tidak aktif, notifikasi bertahap, soft-delete lalu purge (ADM-025), dan audit event (modul 27).

### ADM-055 Site URL rename & custom domains

- **Deskripsi:** Mengubah URL situs (subdomain pada mode hosted) dan memakai custom domain, termasuk redirect URL lama dan dampaknya ke integrasi; pada self-hosted, kebutuhan ini dipenuhi Base URL yang dapat dikonfigurasi (ADM-003).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Rename URL situs tersedia di konsol organisasi dan dibatasi satu kali per 30 hari.
  2. Validasi subdomain: unik, 3-63 karakter huruf kecil/angka/tanda hubung, tidak diawali atau diakhiri tanda hubung, dan bukan nama yang dicadangkan (mis. `admin`, `api`, `www`).
  3. URL lama dialihkan permanen (301) ke URL baru untuk tautan, smart link, dan bookmark.
  4. Halaman konfirmasi mendaftar integrasi yang mungkin perlu dikonfigurasi ulang: application links, webhook, OAuth app, dan ekstensi.
  5. ⚠ Custom domain: admin menambah domain, sistem menampilkan rekaman DNS (CNAME/TXT) untuk verifikasi, menerbitkan dan memperbarui sertifikat TLS otomatis, dan mengalihkan URL bawaan ke custom domain.
  6. Setelah custom domain aktif, tautan email, payload webhook, OAuth redirect, dan smart link memakai domain baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Custom domain help center Service Management adalah fitur terpisah (modul 30). Semua tautan absolut dibangun dari satu fungsi `BaseURL()` agar rename konsisten.

### ADM-056 Atlassian Collections (bundled licensing & entitlements)

- **Deskripsi:** Paket bundel lisensi yang menggabungkan beberapa app (mis. pelacak kerja, Service Management, Product Discovery, AI) dalam satu langganan dan memengaruhi entitlement fitur, akses AI, dan upgrade prompt; bagi produk ini dimodelkan sebagai entitlement bundle per tenant tanpa billing nyata.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Organisasi, Billing admin, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Bundle dapat dipilih atau di-upgrade dari halaman Billing (ADM-053).
  2. Bundle memberi entitlement gabungan per pengguna untuk semua app di dalamnya dalam satu langganan.
  3. ⚠ Tier (Standard/Premium/Enterprise) berlaku seragam untuk semua app dalam bundle.
  4. ⚠ Pengguna yang diberi akses bundle otomatis mendapat akses ke semua app di dalamnya; akses per app tetap dapat dicabut satu per satu.
  5. Kuota AI dan akses agent (ADM-052) dapat terikat pada bundle.
  6. Halaman plan/billing menampilkan bundle aktif dan app di dalamnya, dan app standalone menampilkan prompt upgrade ke bundle bila diaktifkan admin.
  7. ⚠ Keluar dari bundle mengembalikan langganan ke per app dengan entitlement masing-masing; app yang tidak dilanggan ulang mengikuti aturan pembatalan ADM-054.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitlement dievaluasi di layanan `entitlements` bersama flag VISI-003 (dirujuk modul 02). Paket komersial produk acuan di luar cakupan.

## Operasi sistem self-hosted (rujukan Data Center)

### ADM-057 System info & troubleshooting tools (Data Center)

- **Alias inventaris:** Admin Operational Tools
- **Deskripsi:** Kumpulan alat informasi sistem, operasi infrastruktur, dan troubleshooting yang di produk acuan hanya ada di versi Data Center/Server; dicatat sebagai rujukan, dengan subset yang relevan bagi instance self-hosted (System info, health check, job, sesi, logging) diserap ke modul 27.
- **Prioritas / Fase / Tier:** P2 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Operator server/SRE, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan System information (versi, build, JVM, OS, database, memori, path, plugin); padanan produk ini menampilkan versi aplikasi, commit build, versi Go runtime, OS, versi PostgreSQL, dan pemakaian memori.
  2. Versions & licenses (license key, user tier, update check) tidak dibangun karena produk ini tidak memakai license key; cek versi terbaru opsional.
  3. Rujukan troubleshooting tools (instance health checks, log analyzer, support zip, support request); padanannya adalah health check (butir 17) dan unduhan bundel diagnostik tanpa data pribadi.
  4. Rujukan logging and profiling; padanannya ada di butir 21.
  5. Services (mail fetcher POP/IMAP, export/backup service, service kustom kelas Java dengan interval/cron) tidak dibangun sebagai halaman admin; job terjadwal didefinisikan di kode dan dipantau lewat butir 18.
  6. Scheduler details (daftar job internal, run mode, jadwal berikutnya, status run terakhir) diserap sebagai bagian dari butir 18.
  7. Listeners (event listener kustom) tidak dibangun; ekstensi memakai webhook dan model ekstensi (modul 25).
  8. Integrity checker untuk memperbaiki data tidak konsisten tidak disediakan sebagai UI; pemeriksaan integritas berjalan sebagai job internal dan melaporkan temuan ke log.
  9. Database monitoring, instrumentation, dan clustering (daftar node, status, cluster lock, replikasi cache) tidak dibangun di UI; metrik diekspos untuk sistem observability (modul 27).
  10. CDN configuration tidak menjadi setting admin; aset SPA ter-embed di binary dan dapat ditaruh di belakang CDN oleh operator.
  11. Rate limiting global REST API dan exemption per pengguna diatur lewat konfigurasi server (modul 23/27), bukan halaman admin.
  12. LexoRank management (indikator kesehatan Healthy/Warning/Critical, statistik panjang rank, full rebalance manual dengan progres, rebalancing otomatis latar belakang, pemeriksaan dan perbaikan duplikat/marker hilang, board tetap dapat dipakai selama rebalancing) diwujudkan sebagai job otomatis dan metrik di modul 27/28 tanpa halaman admin.
  13. Safe mode untuk app (menonaktifkan semua ekstensi sementara) diserap sebagai aksi "Disable all extensions" di ADM-034 untuk isolasi masalah.
  14. ⚠ Dark features (halaman tersembunyi daftar key fitur) tidak dibangun; feature flag dikelola di halaman flag VISI-003.
  15. Pada mode hosted tidak ada akses runtime atau database bagi admin tenant; digantikan halaman status layanan dan dukungan operator.
  16. Halaman System info admin menampilkan versi, uptime, status koneksi database, status indeks pencarian, dan statistik job (antre, berjalan, gagal 24 jam).
  17. Health checks untuk database, storage, mail, indeks, dan job queue terlihat di UI admin dengan status OK/Warning/Error dan pesan penyebab, serta tersedia di endpoint `/healthz` dan `/readyz` untuk orkestrator.
  18. Trigger reindex dan statusnya tersedia untuk operator self-hosted (rincian di ADM-058).
  19. Admin dapat melihat job gagal beserta galatnya, menjalankan retry per job, dan menjalankan "Clear cache" untuk cache aplikasi.
  20. Admin dapat melihat sesi pengguna aktif (pengguna, IP, perangkat, terakhir aktif) dan memaksa logout satu sesi atau semua sesi pengguna.
  21. Konfigurasi level logging per komponen dapat diubah runtime, dan endpoint profiling `pprof` hanya dapat diakses Admin Situs terautentikasi (nonaktif default).
  22. ⚠ Announcement banner dipenuhi oleh ADM-007.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Butir 1-15 berasal dari "System info & troubleshooting tools (Data Center)" dan butir 16-22 dari "Admin Operational Tools"; keduanya ditandai dc-only-or-deprecated di inventaris sehingga berstatus rujukan, tetapi butir 16-21 direkomendasikan diadopsi untuk instance self-hosted di modul 27 (Background Jobs & Queue, Observability). Keputusan fasenya ada di Pertanyaan terbuka.

### ADM-058 Indexing administration

- **Deskripsi:** Pengelolaan indeks pencarian. Di versi Data Center admin dapat memicu re-index; di versi Cloud indexing berjalan otomatis tanpa UI admin. Produk ini mengikuti model otomatis.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Operator server (self-hosted).
- **Aturan perilaku (acceptance criteria):**
  1. Full re-index (dengan lock) dan background re-index dengan progres dan estimasi waktu tidak tersedia di UI admin; operator self-hosted dapat menjalankan re-index latar belakang lewat CLI server dengan progres di log.
  2. Project re-index dari halaman proyek tidak disediakan.
  3. Pesan "reindex required" setelah perubahan field atau context tidak ditampilkan karena re-index terkait berjalan otomatis.
  4. Index recovery, snapshot, dan replikasi indeks di cluster tidak dibangun; indeks berbasis PostgreSQL ikut backup database (ADM-035).
  5. Perubahan diproses otomatis tanpa re-index manual, dengan kemungkinan jeda singkat pada hasil JQL.
  6. Indexing language mengikuti pengaturan Internationalization (ADM-005).
  7. Hasil JQL bersifat eventually consistent setelah update besar atau bulk; UI menampilkan catatan "Hasil mungkin belum mencakup perubahan beberapa detik terakhir" saat job indeks masih antre.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme indexing internal di SRC-003 (modul 13).

## Pertanyaan terbuka

- Apakah subset operasional self-hosted ADM-057 (System info, health check, retry job, sesi aktif, logging/pprof) dinaikkan dari rujukan menjadi fitur Fase 3, atau cukup didefinisikan di modul 27?
- Apakah step-up auth admin (ADM-001 butir 9) diaktifkan secara default pada instalasi tanpa SSO?
- Apakah produk ini menyediakan mode hosted multi-tenant (billing, trial, deaktivasi Free, rename subdomain, data residency) atau hanya distribusi self-hosted, sehingga ADM-048, ADM-053, ADM-054, ADM-055, dan ADM-056 cukup dimodelkan sebagai entitlement tanpa UI komersial?
- Apakah kebijakan retensi data (ADM-047), yang belum terkonfirmasi ada di produk acuan, tetap dibangun sebagai nilai tambah kepatuhan?
- Berapa nilai default batas objek (ADM-041) untuk instance self-hosted: menyalin angka produk acuan atau menetapkan angka sendiri berdasarkan uji beban?
- Sumber data dan format apa yang dipakai sebagai pengganti analytics data lake Enterprise (ADM-042 butir 8)?
- Importer alat lain mana (selain Trello, Asana, GitHub Issues) yang wajib tersedia saat rilis ADM-038?
- Apakah promote konfigurasi sandbox ke produksi (ADM-049 butir 7), yang di produk acuan masih sebagian roadmap, termasuk cakupan Fase 4 atau ditunda?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga:** fungsi app pihak ketiga, katalog Marketplace produk acuan, model ganda Forge/Connect, importer pihak ketiga, alat promosi konfigurasi (mis. Configuration Manager, Salto), dan provider DevOps/security pihak ketiga (Snyk, LaunchDarkly, dll.) tidak dibangun; produk ini hanya menyediakan model ekstensi native (ADM-034, modul 25) dan kerangka integrasi (ADM-031).
- **Fitur Data Center-only (rujukan, Fase —):** ADM-010 General configuration — Data Center-only options (CAPTCHA, logout confirmation, gzip, remote API toggle, external user management, allowlist outbound UI, help path); ADM-040 Backup & restore system (Data Center XML); ADM-057 System info & troubleshooting tools (Data Center) beserta Admin Operational Tools (JVM, license key, services, listeners, integrity checker, clustering, CDN, LexoRank UI, dark features); ADM-058 Indexing administration (re-index manual dan snapshot indeks); websudo sebagai peran terpisah dan peran "Jira System Administrators".
- **Layanan lintas produk produk acuan:** konsol organisasi eksternal produk acuan, Atlassian Backup and restore sebagai layanan, Jira Cloud Migration Assistant dan migrasi cloud-to-cloud resmi, layanan AI dan agent produk acuan (Atlassian Intelligence, Rovo, Rovo credits), Atlassian Analytics & Data Lake, paket Collections komersial, pemrosesan pembayaran dan invoice nyata, SLA uptime dan tingkat support kontraktual; semuanya hanya dirujuk sebagai pembanding perilaku, dengan padanan internal pada ADM-035 sampai ADM-039 dan ADM-052 sampai ADM-056.
- **Ditangani modul lain:** logika permission dan skema (modul 03, 05), notifikasi dan template email (modul 19), automation global (modul 20), payload webhook dan REST API (modul 23), importer rinci (modul 25), serta infrastruktur backup, job queue, observability, dan multi-tenancy (modul 27).
