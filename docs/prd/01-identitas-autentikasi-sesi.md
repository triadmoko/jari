# 01 — Identitas, Autentikasi & Sesi

Modul ini mendefinisikan identitas pengguna produk ini (satu akun global per orang yang dirujuk lewat `accountId` opaque), alur pendaftaran, login, verifikasi email, pengelolaan password, verifikasi dua langkah (2SV), arsitektur sesi dan token yang dapat dicabut, kebijakan autentikasi tingkat organisasi, kredensial integrasi (API token dan persetujuan OAuth), serta penutupan akun mandiri. Modul ini menjadi fondasi semua modul lain: setiap referensi pengguna (assignee, reporter, watcher, mention, user picker, pemilik filter) menyimpan `accountId` yang didefinisikan di sini. Cakupannya juga mencakup jalur migrasi dari kondisi repo saat ini (JWT HS256 24 jam tanpa revocation, endpoint `POST /api/v1/auth/register` dan `POST /api/v1/auth/login`) ke model sesi yang dapat di-revoke tanpa memutus klien lama. Pengelolaan anggota organisasi, grup, domain terverifikasi, identity provider (SAML/OIDC) dan SCIM dibahas di modul 02; modul ini hanya mengatur perilaku autentikasi yang ditimbulkannya.

Modul terkait:

- [00 Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md)
- [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) (managed account, domain terverifikasi, identity provider, SCIM, deaktivasi)
- [03 Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) (izin "Browse users and groups", visibilitas profil)
- [19 Notifikasi In-app & Email](19-notifikasi-email.md) (email keamanan akun)
- [21 Navigasi Global, Halaman Personal & Preferensi](21-navigasi-ux-preferensi.md) (preferensi bahasa, tema, notifikasi personal)
- [22 Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) (audit log, pengaturan keamanan global)
- [23 REST API Publik & Webhook](23-api-publik-webhook.md) (Basic auth API token, OAuth 2.0 3LO, rate limit)
- [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) (app terhubung, akun bot)
- [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) (keamanan, kriptografi, observabilitas)
- [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) (entitas User/Session/Token, event akun)
- [30 Service Management: Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) (akun `customer` portal-only)

## Persona terkait

- **Admin Situs** — mengatur kebijakan autentikasi, mewajibkan 2SV, memaksa reset password, mencabut sesi dan token pengguna terkelola.
- **Developer/Anggota Tim** — mendaftar, login, mengelola password, 2SV, sesi, profil, dan API token pribadi.
- **Pengguna Tim Bisnis** — memakai alur login dan profil yang sama dengan bahasa yang lebih sederhana.
- **Integrator/Developer API** — memakai API token (Basic auth `email:token`) dan alur persetujuan OAuth untuk app.
- **Tamu/Kolaborator Eksternal** — masuk lewat undangan, akun unmanaged, dapat dikenai enforcement 2SV untuk pengguna eksternal.
- **Customer/Pelapor** — akun bertipe `customer` (portal-only) yang memakai alur login portal; detail portal di modul 30.
- **Stakeholder/Viewer** — melihat profile card pengguna lain sesuai pengaturan visibilitas.

## User story modul

1. Sebagai Developer/Anggota Tim, saya ingin mendaftar dengan email, nama, dan password lalu memverifikasi email agar saya bisa mulai bekerja di proyek tim saya (AUTH-003, AUTH-004).
2. Sebagai Developer/Anggota Tim, saya ingin identitas saya tetap sama walaupun email atau nama saya berubah agar assignment, komentar, dan riwayat saya tidak terputus (AUTH-001).
3. Sebagai Developer/Anggota Tim, saya ingin mereset password lewat email bila lupa agar saya bisa kembali mengakses akun tanpa bantuan admin (AUTH-005).
4. Sebagai Admin Situs, saya ingin percobaan login berulang yang gagal diperlambat dan ditantang CAPTCHA agar akun tim terlindung dari brute-force dan credential stuffing (AUTH-006).
5. Sebagai Developer/Anggota Tim, saya ingin mengaktifkan verifikasi dua langkah dengan authenticator app atau security key agar akun saya tetap aman walaupun password bocor (AUTH-010).
6. Sebagai Developer/Anggota Tim, saya ingin melihat perangkat yang sedang login dan mengeluarkan perangkat tertentu agar saya bisa menutup sesi di perangkat yang hilang (AUTH-009).
7. Sebagai Admin Situs, saya ingin logout, ganti password, dan deaktivasi langsung mengakhiri semua sesi agar akses tidak tersisa setelah insiden atau offboarding (AUTH-007, AUTH-009).
8. Sebagai Integrator/Developer API, saya ingin token lama tetap diterima selama masa transisi agar klien yang sudah ada tidak rusak saat sistem sesi baru diaktifkan (AUTH-008).
9. Sebagai Admin Situs, saya ingin membuat kebijakan autentikasi berisi aturan SSO, 2SV, password, dan durasi sesi untuk kelompok akun terkelola agar standar keamanan organisasi dapat ditegakkan (AUTH-011).
10. Sebagai Integrator/Developer API, saya ingin membuat API token berlabel dengan masa berlaku dan scope agar skrip saya bisa mengakses REST API tanpa memakai password (AUTH-012).
11. Sebagai Developer/Anggota Tim, saya ingin melihat dan mencabut app yang saya beri akses OAuth agar saya mengontrol siapa yang membaca data saya (AUTH-013).
12. Sebagai Stakeholder/Viewer, saya ingin melihat profile card rekan kerja dengan nama, jabatan, dan zona waktu agar saya tahu siapa yang harus dihubungi, tanpa melihat data yang mereka sembunyikan (AUTH-002).
13. Sebagai Tamu/Kolaborator Eksternal, saya ingin menutup akun saya sendiri dengan masa tenggang agar data personal saya dihapus sesuai hak privasi tanpa merusak konten yang pernah saya buat (AUTH-014).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur | Fitur |
|---|---|---|
| P0 | 6 | AUTH-001, AUTH-003, AUTH-004, AUTH-005, AUTH-007, AUTH-008 |
| P1 | 5 | AUTH-002, AUTH-006, AUTH-009, AUTH-010, AUTH-012 |
| P2 | 3 | AUTH-011, AUTH-013, AUTH-014 |
| P3 | 0 | — |

| Fase | Jumlah fitur | Fitur |
|---|---|---|
| Fase 1 | 6 | AUTH-001, AUTH-003, AUTH-004, AUTH-005, AUTH-007, AUTH-008 |
| Fase 2 | 4 | AUTH-002, AUTH-006, AUTH-009, AUTH-012 |
| Fase 3 | 2 | AUTH-010, AUTH-014 |
| Fase 4 | 2 | AUTH-011, AUTH-013 |

Strategi urutan implementasi: Fase 1 dimulai dengan model identitas (`accountId` opaque menggantikan ketergantungan pada `username`) dan arsitektur sesi yang dapat di-revoke, karena semua modul lain menyimpan referensi pengguna dan bergantung pada middleware autentikasi. Bersamaan dengan itu jalur kompatibilitas JWT lama dibangun di balik feature flag agar klien yang ada tidak putus. Alur login email + password, verifikasi email, dan reset password melengkapi MVP; login sosial, passkey, dan kode sekali pakai sengaja ditunda ke fase berikutnya walaupun berada dalam fitur yang sama. Fase 2 menambah perlindungan brute-force, daftar sesi/perangkat, profil dengan visibilitas, dan API token pengguna (AUTH-012) bersamaan dengan dibukanya REST API publik berbasis API token (API-008 di modul 23), lalu Fase 3 menambah 2SV pribadi dan penutupan akun. Kebijakan autentikasi organisasi dan OAuth consent menunggu modul SSO/IdP (modul 02) dan OAuth 2.0 (modul 23) di Fase 4.

## Fitur

## Model identitas & profil

### AUTH-001 Atlassian Account (Identity Model)

- **Deskripsi:** Satu identitas global per orang (disebut "akun" di produk ini) yang dipakai di semua app dan site dalam satu instalasi, dikenali lewat `accountId` opaque, bukan username atau email. Semua referensi pengguna di produk ini (assignee, reporter, mention, watcher, field user picker, pemilik filter) menyimpan `accountId`.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Semua pengguna (pemilik akun); Admin Situs/organisasi untuk akun terkelola; Integrator/Developer API sebagai konsumen `accountId`. Visibilitas atribut di luar field publik minimum dikontrol oleh pengaturan visibilitas profil (AUTH-002) dan izin global "Browse users and groups" (modul 03).

- **Aturan perilaku (acceptance criteria):**

1. `accountId` dibuat sekali saat akun dibuat, berupa string opaque (UUID v4 atau v7 dalam format teks) dengan panjang maksimum 128 karakter, dan tidak pernah berubah ketika email, nama tampilan, atau atribut lain berubah; API yang mencoba mengubahnya ditolak dengan 400.
2. Setiap akun memiliki `accountType` bernilai tepat salah satu dari `atlassian` (manusia; nama nilai sementara dipertahankan demi kompatibilitas API, lihat Pertanyaan terbuka 9; label UI "Pengguna"), `app` (bot integrasi/ekstensi), `customer` (akun portal-only Service Management), atau `unknown`; akun `app` tidak dapat login interaktif dan akun `customer` tidak dapat mengakses UI agen/proyek.
3. Satu akun dapat menjadi anggota lebih dari satu site/organisasi; hak akses dievaluasi per site dan per app (produk), sehingga kehilangan akses di satu site tidak memengaruhi akses di site lain.
4. Email adalah login identifier yang unik secara global tanpa memperhatikan huruf besar/kecil (disimpan ter-normalisasi lowercase); tampilnya email kepada pengguna lain mengikuti pengaturan visibilitas profil (AUTH-002).
5. Status akun bernilai `active`, `deactivated` (oleh admin akun terkelola; tidak dapat login, data tetap utuh), atau `closed` (ditutup pengguna sendiri atau dihapus admin); akun `deactivated` dan `closed` gagal login dengan pesan netral dan semua sesinya dicabut (AUTH-007), dengan satu pengecualian: akun `closed` yang masih dalam masa tenggang penghapusan mandiri hanya dapat masuk ke halaman pembatalan penghapusan (AUTH-014 butir 3).
6. Field publik minimum `accountId`, public name (nama tampilan), dan avatar selalu dapat dibaca oleh pengguna terautentikasi lain di site yang sama, terlepas dari pengaturan visibilitas.
7. REST API publik dan internal hanya menerima `accountId` untuk mengidentifikasi pengguna; parameter `username` atau `key` diabaikan atau ditolak dengan 400 beserta pesan yang menyarankan `accountId`.
8. Akun dengan email pada domain terverifikasi milik organisasi ditandai `managed` (atribut tertentu dikelola admin, lihat modul 02), sedangkan akun lain `unmanaged` dan dikelola sendiri oleh pemiliknya; status ini dapat berubah saat domain diverifikasi atau dilepas.
9. Setiap akun menyimpan `locale` dan `timezone` (IANA, mis. `Asia/Jakarta`); keduanya dipakai untuk format tanggal/waktu di UI dan email notifikasi, dengan fallback ke default site bila kosong.
10. Avatar default berupa inisial nama (maksimum 2 huruf) dengan warna latar yang diturunkan secara deterministik dari `accountId`; avatar unggahan pengguna menggantikannya (AUTH-002).
11. `accountId` menjadi primary key dan foreign key referensi pengguna di semua modul; tidak ada tabel lain yang memakai email atau username sebagai kunci referensi, dan tidak tersedia fitur menggabungkan (merge) dua akun.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Migrasi dari repo saat ini: tabel `users` memakai `id VARCHAR(36)` dan kolom `username` UNIQUE NOT NULL; `id` dipertahankan sebagai `accountId`, sedangkan `username` diturunkan menjadi atribut opsional non-unik (atau dihapus) dan dihilangkan dari klaim JWT baru (lihat AUTH-008). Kolom `is_active` dipetakan ke `status`, kolom `name` menjadi Full name, dan Public name diisi awal dari `name`.
- Migrasi email: constraint `UNIQUE` pada `users.email` saat ini peka huruf besar/kecil; sebelum memasang indeks unik pada `lower(email)` (atau tipe `citext`), migrasi wajib mendeteksi pasangan email yang hanya berbeda huruf besar/kecil dan melaporkannya untuk diselesaikan admin, bukan menggabungkannya otomatis (butir 4, butir 11).
- Entitas dan event akun (`account.created`, `account.closed`, dsb.) didefinisikan di [28 Model Data Inti](28-model-data-inti.md); akun `customer` dirinci di [30](30-jsm-help-center-portal-pelanggan-kanal.md); akun `app` di [25](25-integrasi-impor-ekspor-ekstensi.md) dan [USR-029 App users (bot accounts)](02-pengguna-grup-organisasi.md); managed account di USR-032 dan perlakuan pengguna nonaktif/terhapus di USR-007/USR-008 ([02](02-pengguna-grup-organisasi.md)).
- Keputusan terbuka: apakah produk ini multi-site dalam satu instalasi atau satu site per instalasi (memengaruhi butir 3, AUTH-003 butir 13-14).

### AUTH-002 Profile & visibility settings

- **Deskripsi:** Profil akun (nama, foto, jabatan, departemen, lokasi, dan lainnya) beserta pengaturan visibilitas per field yang menentukan siapa yang dapat melihat setiap data tersebut.

- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform

- **Aktor & izin:** Pemilik akun mengedit profil dan visibilitasnya; Admin Situs/organisasi dapat mengunci field tertentu untuk akun terkelola; pengguna lain melihat sesuai visibilitas dan izin "Browse users and groups".

- **Aturan perilaku (acceptance criteria):**

1. Profil memiliki field Full name, Public name, Job title, Department, Organization, Based in, Local time/timezone, Email, Profile photo, dan Header image; Full name wajib (1-200 karakter), field teks lain opsional dengan batas 255 karakter.
2. Setiap field (kecuali field publik minimum AUTH-001 butir 6) memiliki visibilitas `Anyone`, `Organisasi/site yang sama`, atau `Only you`; untuk akun terkelola, admin dapat mengunci nilai dan/atau visibilitas field tertentu sehingga kontrolnya tampil nonaktif dengan keterangan "Dikelola oleh organisasi Anda".
3. Public name wajib dan dipakai sebagai nama tampilan bagi siapa pun di luar cakupan visibilitas Full name; jika Public name kosong saat migrasi, nilainya diisi dari Full name.
4. Visibilitas email default `Only you`; bila email tidak terlihat oleh pemanggil, pencarian pengguna berdasarkan email tidak mengembalikan akun tersebut dan field `emailAddress` di respons REST API dikosongkan (bukan dihapus dari skema).
5. Pengguna dapat mengunggah foto profil (PNG/JPEG/GIF, maksimum 5 MB), memotong (crop) ke rasio 1:1, dan menghapusnya; setelah dihapus avatar kembali ke inisial (AUTH-001 butir 10).
6. Pengguna dapat mengunggah (maksimum 5 MB) dan menghapus header image profil; tanpa header image ditampilkan latar warna default.
7. Untuk akun terkelola dengan sinkronisasi IdP/SCIM, atribut yang disinkronkan tampil read-only bagi pengguna dengan keterangan sumbernya; perubahan hanya dapat dilakukan di IdP (modul 02).
8. Profile card (hover card) pada avatar/nama di seluruh UI menampilkan nama tampilan, avatar, dan field lain hanya sesuai visibilitas terhadap pengguna yang melihat, ditambah waktu lokal bila timezone terlihat.
9. Halaman profil pengguna menampilkan informasi profil, work item terbaru yang terlibat, dan tim; work item yang ditampilkan difilter oleh izin Browse projects dan security level pengamat, sehingga tidak ada work item yang bocor.
10. ⚠ Pengguna dapat membuka pratinjau profilnya sebagaimana dilihat oleh "Anyone" dan oleh "anggota organisasi" untuk memeriksa efek pengaturan visibilitas.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Preferensi personal (bahasa, tema, notifikasi) serta UI NAV-027 "Profile Page (Atlassian account profile)" dan NAV-028 "Profile Hover Card" berada di [21](21-navigasi-ux-preferensi.md); modul ini menetapkan data dan aturan visibilitas yang dipakai UI tersebut. Direktori orang (USR-027) di [02](02-pengguna-grup-organisasi.md); Teams (PLAN-057) dan halaman profil tim (PLAN-059) di [18](18-timeline-rencana-tampilan-proyek.md).
- Penyimpanan berkas avatar memakai layanan lampiran ([08](08-komentar-lampiran-aktivitas.md)).

## Pendaftaran, login & kredensial dasar

### AUTH-003 Sign up & Log in

- **Deskripsi:** Alur pendaftaran dan login akun di halaman identitas produk ini, termasuk metode login alternatif, home realm discovery ke SSO, dan alur dari undangan.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Calon pengguna dan pengguna terdaftar. Pendaftaran mandiri dapat dibatasi oleh pengaturan akses site (undangan saja, domain tertentu; modul 02/22). Metode login dibatasi oleh kebijakan autentikasi (AUTH-011).

- **Aturan perilaku (acceptance criteria):**

1. Sign up dengan email memerlukan email, full name, dan password yang lolos aturan AUTH-005; akun dibuat berstatus belum terverifikasi lalu email verifikasi (link dan/atau kode) dikirim (AUTH-004); bila email sudah terdaftar, respons tetap sama secara visual dan email berisi tautan login/reset dikirim ke pemilik email.
2. Login memakai dua layar: layar pertama meminta email, layar kedua meminta password; bila domain email memakai SSO yang di-enforce, layar password dilewati dan pengguna diarahkan ke IdP.
3. Login sosial "Continue with Google / Microsoft / Apple / Slack" (OIDC/OAuth) menautkan identitas eksternal ke akun dengan email terverifikasi yang sama; bila belum ada akun, akun baru dibuat. Ditargetkan Fase 3, bukan MVP.
4. Login dengan passkey (WebAuthn) tersedia di halaman login; pengguna dapat mendaftarkan lebih dari satu passkey dan menghapusnya; passkey tidak tersedia untuk akun terkelola. Ditargetkan Fase 4.
5. ⚠ Login dengan kode sekali pakai (6 digit, berlaku 10 menit, sekali pakai) yang dikirim ke email sebagai alternatif password.
6. Home realm discovery: email dengan domain terverifikasi yang memiliki SSO otomatis diarahkan ke IdP yang terkait dengan kebijakan autentikasi akun tersebut (AUTH-011, modul 02).
7. Opsi "Tetap login" (remember me) mempertahankan sesi di perangkat hingga batas durasi sesi kebijakan autentikasi; tanpa opsi ini sesi berakhir saat browser ditutup (cookie sesi).
8. Tautan "Can't log in?" membuka halaman bantuan berisi reset password (AUTH-005), kirim ulang verifikasi (AUTH-004), dan pemulihan 2SV (AUTH-010).
9. Membuka tautan undangan yang valid untuk email yang belum terdaftar langsung membuka formulir pembuatan akun dengan email terkunci, lalu menautkan akun ke site setelah akun dibuat; undangan kedaluwarsa menampilkan pesan dan opsi meminta undangan baru.
10. Kebijakan autentikasi dapat memblokir login pihak ketiga (Google/Microsoft/Apple/Slack) untuk akun terkelola; tombol tetap tampil tetapi percobaan login ditolak dengan pesan "Organisasi Anda tidak mengizinkan metode login ini".
11. Setelah login berhasil, pengguna diarahkan ke continue URL (parameter `continue`) yang wajib berupa path relatif atau origin yang terdaftar; URL lain diabaikan dan pengguna diarahkan ke halaman beranda (mencegah open redirect).
12. Semua pesan galat login, pendaftaran, dan reset bersifat netral ("Email atau password salah") dan memiliki waktu respons yang serupa, sehingga keberadaan akun tidak dapat ditebak (anti user enumeration).
13. Account switcher memungkinkan beberapa akun login di browser yang sama dan berpindah akun tanpa logout; setiap akun memiliki sesi terpisah.
14. Bila pengguna memiliki akses ke lebih dari satu site, setelah login ditampilkan halaman "Select site" berisi daftar site yang dapat diakses; bila hanya satu, langsung diarahkan.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Kondisi repo saat ini: `POST /api/v1/auth/register` menerima username dan `POST /api/v1/auth/login` mengembalikan JWT 24 jam. Kontrak login dipertahankan (AUTH-008), username di register menjadi opsional.
- MVP (Fase 1) cukup butir 1, 2, 7, 8, 9, 11, 12; butir 13-14 bergantung keputusan multi-site; butir 6 dan 10 bergantung SAML SSO (USR-034), domain terverifikasi (USR-031), dan AUTH-011. Pembatasan pendaftaran mandiri mengikuti USR-012; alur undangan mengikuti USR-009 dan USR-011 ([02](02-pengguna-grup-organisasi.md)).

### AUTH-004 Email verification & email change

- **Deskripsi:** Verifikasi alamat email akun dan alur penggantian email, dengan aturan berbeda untuk akun terkelola dan tidak terkelola.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Pemilik akun (unmanaged) mengganti email sendiri; Admin organisasi mengganti email akun terkelola.

- **Aturan perilaku (acceptance criteria):**

1. Akun dengan email belum terverifikasi hanya dapat mengakses halaman verifikasi dan profil dasar; UI lain menampilkan banner "Verifikasi email Anda" dan API mengembalikan 403 dengan kode galat `EMAIL_NOT_VERIFIED`.
2. Pengguna dapat mengirim ulang email verifikasi; token sebelumnya dibatalkan dan pengiriman dibatasi maksimum 5 kali per jam per akun.
3. Pengguna unmanaged mengganti email dengan memasukkan email baru dan password saat ini (re-autentikasi, AUTH-009 butir 9); email lama tetap berlaku sampai tautan verifikasi ke email baru diklik.
4. Untuk akun terkelola, field email tampil read-only bagi pengguna; hanya admin organisasi yang dapat mengubahnya, dan nilai baru harus berada di domain terverifikasi milik organisasi tersebut.
5. Email yang masih terikat ke akun lain (aktif, dinonaktifkan, atau dalam masa tenggang penghapusan) ditolak saat penggantian dengan pesan netral yang tidak mengungkap pemilik.
6. Setelah email berubah, notifikasi dikirim ke email lama berisi waktu perubahan dan tautan untuk melaporkan bila perubahan tidak dilakukan pemilik.
7. Akun yang dibuat dari tautan undangan dianggap terverifikasi setelah tautan undangan diklik, sehingga email verifikasi terpisah tidak dikirim.
8. Email akun yang disinkronkan dari SCIM/IdP read-only di produk ini dan hanya dapat diubah dari IdP.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Token verifikasi mengikuti aturan AUTH-007 butir 10 (sekali pakai, berumur pendek, disimpan sebagai hash).
- Template email di [19](19-notifikasi-email.md); managed account (USR-032) dan SCIM (USR-035) di [02](02-pengguna-grup-organisasi.md).
- Migrasi: akun yang sudah ada di repo dianggap terverifikasi (`email_verified_at` diisi waktu migrasi) agar pengguna lama tidak terkunci.

### AUTH-005 Password management

- **Deskripsi:** Pengelolaan password akun: ganti, reset, validasi kekuatan, penolakan password bocor, dan password policy dari admin untuk akun terkelola.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Pemilik akun; Admin organisasi (paksa reset, password policy pada kebijakan autentikasi).

- **Aturan perilaku (acceptance criteria):**

1. Ganti password di Account settings > Security mewajibkan password lama yang benar; password baru tidak boleh sama dengan password lama.
2. Forgot password mengirim tautan reset ke email; tautan berlaku 1 jam, sekali pakai, dan dibatalkan bila tautan baru diminta; respons halaman selalu sama terlepas dari email terdaftar atau tidak.
3. Indikator kekuatan password tampil real-time (skor 0-4 ala zxcvbn); password dengan skor di bawah ambang (default 2), kurang dari 8 karakter, atau lebih dari 128 karakter ditolak dengan pesan alasan.
4. ⚠ Password yang diketahui bocor ditolak saat daftar, ganti, dan reset dengan memeriksa basis data password bocor memakai k-anonymity (hanya 5 karakter pertama hash SHA-1 yang dikirim) atau dataset lokal.
5. Setelah ganti password, pengguna dapat memilih "Log out dari semua perangkat lain" (default aktif); setelah reset password, semua sesi lain selalu dicabut (AUTH-007 butir 3).
6. Akun terkelola dengan SSO enforced tidak memiliki password produk ini; menu ganti/reset password disembunyikan dan permintaan reset tidak mengirim email.
7. Admin dapat memaksa reset password akun terkelola per pengguna atau untuk semua anggota suatu kebijakan autentikasi; password lama langsung tidak berlaku dan pengguna menerima email reset.
8. Setiap perubahan password (ganti, reset, paksa reset) mengirim email notifikasi ke pemilik akun.
9. Password policy pada kebijakan autentikasi (AUTH-011) dapat menetapkan kekuatan minimum dan masa berlaku password (dalam hari, atau never).
10. Saat password policy diperketat atau password kedaluwarsa, pengguna yang terdampak wajib mengganti password pada login berikutnya sebelum dapat mengakses halaman lain.
11. Permintaan reset password dibatasi (default 3 per jam per email dan 20 per jam per IP); permintaan berlebih mengembalikan 429 tanpa mengungkap keberadaan akun.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Kondisi repo: `pkg/crypto` memakai bcrypt `DefaultCost` (10); target bcrypt cost ≥ 12 atau argon2id dengan rehash transparan saat login (AUTH-007 butir 5).
- bcrypt hanya memproses 72 byte pertama input; karena batas panjang password 128 karakter (butir 3), implementasi wajib memakai argon2id atau pre-hash (mis. HMAC-SHA-256 lalu bcrypt) agar dua password yang berbeda setelah byte ke-72 tidak dianggap sama.
- Password policy granular untuk akun terkelola setara add-on keamanan organisasi berbayar di produk acuan; tier final ditentukan pemilik produk.

### AUTH-006 Failed login protection & CAPTCHA

- **Deskripsi:** Perlindungan brute-force pada login lewat CAPTCHA setelah beberapa kali gagal, ditambah rate limiting dan deteksi bot.

- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform

- **Aktor & izin:** Sistem (otomatis); pengguna yang login. Ambang tidak dapat dikonfigurasi oleh admin site (mengikuti perilaku produk acuan versi Cloud); operator instalasi dapat mengubahnya lewat konfigurasi server.

- **Aturan perilaku (acceptance criteria):**

1. Setelah 3 percobaan login gagal berturut-turut untuk email yang sama (dalam 15 menit), CAPTCHA wajib diselesaikan sebelum percobaan berikutnya diproses.
2. Rate limiting diterapkan per IP (default 30 percobaan/menit) dan per akun (default 10 percobaan/15 menit); permintaan di atas batas mendapat 429 dengan header `Retry-After`.
3. Perlindungan credential stuffing: pola banyak email berbeda dari satu IP/ASN atau user-agent bot memicu CAPTCHA untuk IP tersebut dan tercatat di log keamanan.
4. Pesan galat login selalu generik ("Email atau password salah") baik untuk email tak terdaftar, password salah, maupun akun nonaktif.
5. Counter percobaan gagal suatu akun di-reset ke 0 setelah login berhasil.
6. REST API tidak menerima Basic auth dengan password akun (hanya API token, AUTH-012), sehingga CAPTCHA dan lockout tidak mengganggu integrasi.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Penyedia CAPTCHA harus dapat diganti (self-host friendly, mis. hCaptcha/Turnstile/CAPTCHA lokal); keputusan di Pertanyaan terbuka.
- Setting "Maximum Authentication Attempts" dan reset failed login count per pengguna hanya ada di edisi Data Center produk acuan dan berada di luar cakupan.
- Rate limit umum API di [23](23-api-publik-webhook.md) dan [27](27-nfr-arsitektur-platform.md).

## Arsitektur sesi & token

### AUTH-007 Authentication, Sessions & 2FA

- **Deskripsi:** Mekanisme login dan pengelolaan sesi tingkat platform yang dibangun di atas JWT yang sudah ada di repo, dilengkapi revocation, timeout, kebijakan password kriptografis, 2FA, SSO opsional, dan rotasi kunci penandatangan. Fitur ini adalah kontrak teknis yang diimplementasikan oleh fitur pengguna AUTH-003, AUTH-005, AUTH-009, AUTH-010, dan AUTH-012.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Pengguna akhir (pemegang sesi); Admin Situs (konfigurasi timeout, enforcement); tim Security/Engineering (rotasi kunci, konfigurasi server).

- **Aturan perilaku (acceptance criteria):**

1. Model A (opsi rekomendasi): access token JWT berumur ±15 menit dikirim sebagai Bearer token, ditambah refresh token opaque yang dirotasi setiap kali dipakai dan disimpan di cookie `httpOnly`, `Secure`, `SameSite=Lax` (atau `Strict`); pemakaian ulang refresh token yang sudah dirotasi (reuse) mencabut seluruh keluarga token sesi tersebut dan mencatat event keamanan.
2. Model B (alternatif kompatibel "JWT 12h sliding"): token berumur 12 jam yang diperpanjang otomatis bila ada aktivitas dalam jendela perpanjangan yang dapat dikonfigurasi; setiap token membawa klaim `jti`/`sid` yang dicek ke tabel sesi pada setiap permintaan (boleh melalui cache status sesi dengan TTL sesuai batas butir 3) sehingga dapat di-revoke. Tim memilih tepat satu model (A atau B) sebelum implementasi.
3. Revocation: logout, ganti/reset password, deaktivasi pengguna, penutupan akun, dan reset 2FA langsung mengakhiri semua sesi terkait; permintaan berikutnya dengan token sesi tersebut ditolak 401 paling lambat 5 detik setelah pencabutan (batas atas TTL cache status sesi), tanpa menunggu expiry alami; dalam Model A hal ini berarti status sesi tetap dicek walaupun access token belum kedaluwarsa.
4. Idle timeout dan absolute session timeout dapat dikonfigurasi admin (global, dan per kebijakan autentikasi pada AUTH-011); sesi yang melewati salah satu batas wajib login ulang, dan nilai efektif adalah yang paling ketat.
5. Password disimpan dengan bcrypt cost ≥ 12 atau argon2id; panjang minimal 8 karakter dan pengecekan password bocor (HIBP k-anonymity atau dataset lokal) diterapkan konsisten dengan AUTH-005; hash lama dengan parameter lebih lemah di-rehash transparan saat login berhasil.
6. 2FA TOTP (RFC 6238, 6 digit, periode 30 detik, toleransi ±1 langkah) dengan recovery codes sekali pakai; enforcement 2-step verification per organisasi/kebijakan diterapkan sebelum sesi penuh diterbitkan (detail perilaku pengguna di AUTH-010).
7. SSO SAML 2.0/OIDC dan SCIM provisioning bersifat opsional dan dapat dinonaktifkan seluruhnya; ketika aktif, sesi hasil SSO diterbitkan oleh mekanisme sesi yang sama dan tunduk pada revocation yang sama (konfigurasi IdP di modul 02).
8. API token untuk Basic auth REST (`email:token`) memiliki expiry dan scope, disimpan sebagai hash, dan divalidasi oleh middleware yang sama dengan Bearer token (detail di AUTH-012).
9. Kunci penandatangan JWT mendukung rotasi: header `kid` wajib pada token baru, beberapa kunci verifikasi dapat aktif bersamaan, dan kunci yang dipensiunkan berhenti menerima token setelah masa tenggang; `JWT_SECRET` wajib minimal 256-bit (32 byte) dan server gagal start (fail-fast) bila kosong, terlalu pendek, atau masih bernilai default.
10. Token verifikasi email dan lupa password bersifat sekali pakai, berumur pendek (verifikasi ≤ 24 jam, reset ≤ 1 jam), disimpan sebagai hash, dan respons endpoint tidak pernah mengungkap apakah email terdaftar.
11. Profil pengguna menyediakan daftar sesi/perangkat aktif dan aksi "Log out of all devices" (detail UI di AUTH-009).

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Kondisi repo terverifikasi: `pkg/jwt.GenerateToken` menerbitkan HS256 dengan expiry 24 jam tetap, tanpa refresh dan revoke; `AuthMiddleware` hanya memvalidasi signature dan expiry; `JWT_SECRET` default `your-secret-key` (tidak fail-fast). Asumsi "12h sliding" pada brief tidak sesuai kode, sehingga keputusan model A/B menjadi pertanyaan terbuka.
- Transisi ke model baru diatur di AUTH-008; kebutuhan kriptografi dan latensi pengecekan sesi di [27](27-nfr-arsitektur-platform.md); tabel `sessions`, `refresh_tokens`, `signing_keys` di [28](28-model-data-inti.md).
- 2FA P1 (Fase 3) dan SSO/SCIM P2 (Fase 4) walaupun fitur induk ini P0; butir 1-5, 9, 10 adalah lingkup Fase 1.

### AUTH-008 Legacy JWT Compatibility & Migration Path

- **Deskripsi:** Rencana transisi dari JWT HS256 24 jam yang ada di repo ke model sesi yang dapat di-revoke (AUTH-007) tanpa memutus klien yang sudah ada.

- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform

- **Aktor & izin:** Engineering (implementasi dan rollout); operator instalasi (mengaktifkan feature flag); klien API yang sudah ada.

- **Aturan perilaku (acceptance criteria):**

1. Selama masa transisi, token lama (tanpa klaim `jti` dan header `kid`) tetap diterima sampai expiry alaminya (maksimum 24 jam sejak terbit); setiap token yang diterbitkan oleh versi baru wajib membawa `jti` dan `kid`. Setelah tanggal cutover yang dikonfigurasi, setiap token tanpa `jti` atau tanpa `kid` ditolak 401 dengan kode galat `TOKEN_FORMAT_UNSUPPORTED` terlepas dari nilai `exp`-nya, sehingga klien diarahkan untuk login ulang.
2. Klaim yang saat ini disimpan middleware di context `sess` (`*jwt.Claims` dengan `user_id`, `email`, `username`) tetap tersedia dengan nama dan tipe yang sama sehingga handler lama berjalan tanpa perubahan; field baru (`sid`, `account_type`) ditambahkan tanpa menghapus field lama.
3. Endpoint refresh/sliding baru (mis. `POST /api/v1/auth/refresh` dan `POST /api/v1/auth/logout`) ditambahkan tanpa mengubah kontrak request dan response `POST /api/v1/auth/login`; field tambahan pada respons login bersifat aditif.
4. Feature flag (mis. `AUTH_REVOCATION_ENFORCED`) mengaktifkan pengecekan sesi ke tabel sesi; saat flag mati, sesi tetap dicatat (mode shadow) tetapi token tidak ditolak karena revocation, dan setiap penolakan yang akan terjadi dicatat di log.
5. Test kompatibilitas otomatis wajib mencakup token lama yang valid, token baru yang valid, token yang sudah direvoke, token kedaluwarsa, dan token dengan `kid` tidak dikenal, masing-masing dengan hasil yang diharapkan (diterima/401) pada kondisi flag hidup dan mati.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Bergantung pada keputusan model A/B di AUTH-007. Catatan: dokumentasi repo menyebut tipe klaim `*jwt.JWTClaims`, sedangkan kode memakai `jwt.Claims` di `pkg/jwt`; kontrak yang dipertahankan adalah yang ada di kode.
- Pengurangan klaim `username` dari token baru menunggu migrasi AUTH-001; konsumen klaim tersebut harus dipetakan ke `accountId` sebelum cutover.
- Dokumentasi perubahan API di Swagger (`make swag`) dan catatan rilis di [23](23-api-publik-webhook.md).

### AUTH-009 Account sessions & devices

- **Deskripsi:** Pengelolaan sesi login aktif oleh pengguna dan admin, termasuk daftar perangkat, durasi sesi, idle timeout, logout paksa, dan re-autentikasi untuk aksi sensitif.

- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform

- **Aktor & izin:** Pemilik akun (sesi sendiri); Admin organisasi (sesi akun terkelola; konfigurasi durasi lewat AUTH-011).

- **Aturan perilaku (acceptance criteria):**

1. Account settings > Security menampilkan "Recent devices": browser, sistem operasi, perkiraan lokasi (negara/kota dari IP), alamat IP tersamar sebagian, waktu login, dan aktivitas terakhir; sesi saat ini ditandai "Perangkat ini".
2. Pengguna dapat log out dari satu perangkat tertentu atau dari "semua perangkat lain"; sesi yang dicabut langsung gagal pada permintaan berikutnya (AUTH-007 butir 3).
3. Admin organisasi dapat log out semua sesi akun terkelola tertentu (mis. saat insiden atau offboarding); aksi tercatat di audit log dengan aktor dan target.
4. ⚠ Session duration maksimum diatur di kebijakan autentikasi dengan pilihan rentang dari beberapa jam hingga 90 hari (default 30 hari), dan sesi lebih lama dari durasi itu wajib login ulang.
5. ⚠ Idle session timeout diatur per kebijakan autentikasi dengan pilihan Never (default, nonaktif), 15 menit, 30 menit, 1 jam, 4 jam, 12 jam, 1 hari, 7 hari, dan 30 hari (rentang 15 menit-30 hari, sama dengan ADM-044 butir 3); sesi tanpa aktivitas melewati batas tersebut berakhir.
6. Sesi dicabut otomatis saat pengguna dinonaktifkan, password direset, atau akun ditutup/dihapus.
7. Sesi aplikasi mobile/desktop (refresh token jangka panjang) ikut dicabut saat deaktivasi atau logout paksa, dan aplikasi menampilkan layar login pada sinkronisasi berikutnya.
8. ⚠ Email notifikasi dikirim saat login berhasil dari perangkat (browser + OS) atau lokasi (negara) yang belum pernah dipakai akun tersebut, berisi detail perangkat dan tautan untuk mengamankan akun.
9. Aksi sensitif (ganti email, ganti password, membuat API token, menonaktifkan 2SV, menutup akun) memerlukan re-autentikasi bila login terakhir lebih dari 15 menit yang lalu.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Bergantung pada tabel sesi AUTH-007/AUTH-008; audit log di ADM-014 ([22](22-administrasi-sistem.md)) dan USR-045 ([02](02-pengguna-grup-organisasi.md)); deaktivasi pengguna di USR-007 ([02](02-pengguna-grup-organisasi.md)); kebijakan aplikasi mobile di USR-039.
- Ringkasan kebijakan durasi sesi dan idle timeout yang tampil di Settings serta penegakannya di middleware dibahas di ADM-044 Security & access policies ([22](22-administrasi-sistem.md)); daftar opsi di butir 4-5 adalah sumber kebenaran dan harus identik dengan ADM-044 butir 3.
- Perkiraan lokasi memerlukan basis data GeoIP lokal (mis. MaxMind GeoLite2) agar tidak bergantung layanan eksternal.

### AUTH-010 Two-step verification (2SV)

- **Deskripsi:** Autentikasi dua langkah yang dapat diaktifkan sendiri oleh pengguna atau diwajibkan (enforced) oleh admin lewat kebijakan autentikasi.

- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform

- **Aktor & izin:** Pemilik akun; Admin organisasi (enforcement, reset 2SV akun terkelola); operator instalasi (pemulihan akun unmanaged).

- **Aturan perilaku (acceptance criteria):**

1. Aktivasi dari Account settings > Security memakai aplikasi authenticator TOTP: pengguna memindai QR code atau memasukkan kunci manual, lalu wajib memasukkan satu kode valid untuk menyelesaikan setup.
2. Pengguna dapat mendaftarkan beberapa faktor sekaligus, termasuk security key (WebAuthn/FIDO2) selain authenticator app; setiap faktor memiliki nama dan tanggal ditambahkan.
3. Recovery key ditampilkan sekali saat setup (dengan aksi salin/unduh), dapat di-regenerate (key lama langsung tidak berlaku), dan tersedia pemulihan lewat email bila perangkat hilang dengan konfirmasi tambahan.
4. Saat login, setelah password benar pengguna diminta kode 6 digit atau sentuhan security key; kode salah 5 kali berturut-turut memicu jeda dan CAPTCHA (AUTH-006).
5. Pengguna dapat menonaktifkan 2SV (dengan re-autentikasi), kecuali kebijakan autentikasinya mewajibkan 2SV; dalam kondisi itu tombol nonaktif dengan keterangan alasan.
6. Pengguna dapat mengganti/mengonfigurasi ulang perangkat authenticator dan menghapus security key, selama minimal satu faktor tetap ada saat 2SV diwajibkan.
7. Enforcement 2SV per kebijakan autentikasi untuk akun terkelola: pengguna tanpa 2SV dipaksa menyelesaikan setup pada login berikutnya sebelum dapat mengakses halaman lain.
8. Saat enforcement diaktifkan, email notifikasi dikirim ke setiap pengguna terdampak berisi tenggat dan instruksi setup.
9. ⚠ Admin dapat mereset 2SV akun terkelola yang kehilangan perangkat (aksi "Reset two-step verification"); semua faktor dan recovery key dihapus, sesi dicabut, dan pengguna wajib setup ulang bila enforcement aktif.
10. Pengguna unmanaged yang kehilangan semua faktor dan recovery key tidak dapat dipulihkan oleh admin site; pemulihan hanya lewat prosedur dukungan operator instalasi dengan verifikasi identitas.
11. 2SV tidak diminta untuk login via SSO karena MFA diurus oleh IdP; faktor 2SV yang tersimpan tetap ada tetapi tidak dipakai selama SSO enforced.
12. Status 2SV (aktif/nonaktif, jenis faktor) per pengguna tampil di daftar akun terkelola dan dapat difilter.
13. Enforcement 2SV dapat diterapkan juga pada pengguna eksternal (akun non-managed yang mengakses site organisasi); pengguna tersebut wajib setup 2SV sebelum mengakses site itu.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Kontrak TOTP dan recovery code di AUTH-007 butir 6; enforcement memerlukan AUTH-011 (Fase 4), sehingga butir 7, 8, 12, 13 baru aktif setelah kebijakan autentikasi tersedia.
- Daftar akun terkelola di [02](02-pengguna-grup-organisasi.md).

## Kebijakan autentikasi organisasi

### AUTH-011 Authentication policies

- **Deskripsi:** Kumpulan aturan autentikasi bernama yang di-assign ke sekelompok akun terkelola: SSO, 2SV, password, sesi, dan login pihak ketiga. Setara dengan kebijakan autentikasi pada add-on keamanan organisasi di produk acuan.

- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform

- **Aktor & izin:** Admin organisasi (membuat, mengedit, menghapus policy dan meng-assign anggota). Hanya berlaku bagi akun terkelola pada domain terverifikasi (modul 02).

- **Aturan perilaku (acceptance criteria):**

1. Organisasi selalu memiliki tepat satu default policy (tidak dapat dihapus) dan dapat membuat policy tambahan dengan nama unik (1-100 karakter); akun terkelola baru otomatis masuk default policy.
2. Anggota di-assign ke policy secara manual (pilih pengguna), lewat unggah CSV (kolom email; baris tidak valid dilaporkan per baris), atau berdasarkan IdP asal; satu akun hanya berada di satu policy pada satu waktu.
3. Setting "Enforce single sign-on" mewajibkan login lewat identity provider yang dipilih dari daftar IdP organisasi; anggota policy tidak lagi dapat login dengan password.
4. Setting "Enforce two-step verification" hanya tersedia untuk policy non-SSO dan memicu perilaku AUTH-010 butir 7-8.
5. Setting password requirements menetapkan kekuatan minimum (skor) dan masa berlaku password (hari atau never) sesuai AUTH-005 butir 9-10; tidak tersedia pada policy dengan SSO enforced.
6. Setting session duration dan idle session timeout sesuai pilihan AUTH-009 butir 4-5 berlaku bagi semua anggota policy; perubahan berlaku pada permintaan berikutnya tanpa mencabut sesi yang masih dalam batas baru.
7. ⚠ Setting izinkan/blokir login pihak ketiga (Google/Microsoft/Apple/Slack) per policy, dengan efek seperti AUTH-003 butir 10.
8. Admin dapat membuat policy non-SSO untuk akun admin/break-glass agar tetap dapat login saat IdP bermasalah; UI memperingatkan bila tidak ada satu pun admin di policy non-SSO ketika SSO di-enforce pada default policy.
9. Policy dapat diedit dan dihapus; saat policy dihapus, semua anggotanya pindah ke default policy dan aturan default langsung berlaku.
10. Daftar policy menampilkan ringkasan jumlah anggota per policy beserta setting utama (SSO, 2SV, durasi sesi).
11. Aksi massal "Reset passwords" memaksa reset password semua anggota suatu policy (AUTH-005 butir 7) dengan dialog konfirmasi yang menyebut jumlah pengguna terdampak.
12. Setiap pembuatan, perubahan setting, perpindahan anggota, dan penghapusan policy tercatat di audit log dengan nilai sebelum dan sesudah.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Bergantung pada domain terverifikasi (USR-031), managed account (USR-032), pengguna eksternal (USR-033), SAML SSO (USR-034), dan langganan keamanan organisasi (USR-037) di [02](02-pengguna-grup-organisasi.md); audit log organisasi di USR-045 dan [22](22-administrasi-sistem.md).
- Di produk acuan fitur ini memerlukan add-on keamanan berbayar (termasuk dalam plan Enterprise); apakah fitur ini disertakan tanpa biaya pada edisi self-host adalah keputusan terbuka.

## Kredensial integrasi

### AUTH-012 API tokens (user)

- **Deskripsi:** Token yang dibuat pengguna untuk autentikasi REST API (Basic auth `email:token`) dan tool CLI, sebagai pengganti password.

- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform

- **Aktor & izin:** Pemilik akun (membuat, melihat, mencabut token sendiri); Admin organisasi (melihat/mencabut token akun terkelola, memblokir pembuatan token). Token mewarisi izin pemiliknya, dibatasi oleh scope bila ada.

- **Aturan perilaku (acceptance criteria):**

1. Halaman Account settings > Security > API tokens menampilkan daftar token dan tombol "Create API token"; aksi membuat token memerlukan re-autentikasi (AUTH-009 butir 9).
2. Create API token mewajibkan label (1-100 karakter, unik per akun) dan tanggal kedaluwarsa antara 1 dan 365 hari (default 365).
3. Opsi "API token with scopes" memungkinkan memilih app/produk dan scope tertentu (mis. `read:work-item`, `write:work-item`); permintaan di luar scope ditolak 403 meskipun pemilik memiliki izin.
4. Nilai token hanya ditampilkan sekali saat dibuat beserta tombol copy; setelah dialog ditutup nilai tidak dapat ditampilkan kembali karena hanya hash yang disimpan.
5. Daftar token menampilkan label, tanggal dibuat, terakhir dipakai (tanggal), tanggal kedaluwarsa, dan scope; token kedaluwarsa ditandai.
6. Pengguna dapat me-revoke satu token atau "Revoke all"; token yang dicabut langsung ditolak 401.
7. Email notifikasi dikirim 7 hari dan 1 hari sebelum token kedaluwarsa.
8. Admin organisasi dapat melihat daftar token akun terkelola (tanpa nilai token) dan mencabutnya.
9. Admin organisasi dapat memblokir pembuatan API token oleh akun terkelola; saat diblokir tombol create nonaktif dengan keterangan dan token baru tidak dapat dibuat lewat API.
10. Token tidak berlaku (401) selama akun pemiliknya dinonaktifkan atau ditutup.
11. Basic auth dengan password akun ditolak 401 dengan pesan yang menyarankan API token; hanya `email:api_token` atau Bearer token yang diterima.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Dijadwalkan Fase 2 agar pembuatan token oleh pengguna tersedia bersamaan dengan API-008 (P0 / Fase 2), yang membuka REST API publik bagi pihak ketiga lewat API token; scope token (butir 3) boleh menyusul bila API-010 belum siap, dengan token tanpa scope sebagai default.
- Kontrak validasi di AUTH-007 butir 8; sisi REST API (API-008 API Tokens & Basic Authentication, API-010 scope) dan rate limit di [23](23-api-publik-webhook.md); API key organisasi untuk admin di USR-030 ([02](02-pengguna-grup-organisasi.md)), berbeda dari token pribadi ini.
- Autentikasi integrasi (OAuth 2.0 3LO, akun app/bot) dibahas di [23](23-api-publik-webhook.md) dan [25](25-integrasi-impor-ekspor-ekstensi.md).

### AUTH-013 Connected apps & OAuth consents

- **Deskripsi:** Pengelolaan app pihak ketiga yang diberi akses ke data pengguna lewat OAuth 2.0 (3LO), termasuk layar persetujuan (consent) dan pencabutan akses.

- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform

- **Aktor & izin:** Pemilik akun (memberi/mencabut persetujuan); Admin organisasi (melihat app yang disetujui, mengatur izin instalasi oleh pengguna).

- **Aturan perilaku (acceptance criteria):**

1. Layar consent OAuth menampilkan nama dan pengembang app, daftar scope yang diminta dalam bahasa yang mudah dipahami, dan pilihan site yang diizinkan; pengguna dapat menyetujui atau menolak, dan penolakan mengembalikan `error=access_denied` ke redirect URI.
2. Profil pengguna memiliki daftar "Connected apps" berisi nama app, scope yang diberikan, site, dan tanggal persetujuan.
3. Pengguna dapat me-revoke akses satu app; semua access dan refresh token app tersebut untuk pengguna itu langsung dicabut.
4. Admin organisasi dapat melihat daftar app yang dipasang/disetujui oleh pengguna di organisasinya beserta jumlah pengguna per app.
5. Admin organisasi dapat memblokir atau mengizinkan instalasi app oleh pengguna (user-installed apps), secara global atau per app; saat diblokir layar consent menampilkan pesan bahwa organisasi tidak mengizinkan app tersebut.
6. Token OAuth milik pengguna dicabut otomatis saat pengguna dinonaktifkan atau akun ditutup.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Server otorisasi OAuth 2.0 (pendaftaran app, endpoint authorize/token, PKCE) didefinisikan di API-009 dan scope di API-010 ([23](23-api-publik-webhook.md)); fitur ini hanya mengatur persetujuan dan pencabutan sisi pengguna/admin.
- App Marketplace pihak ketiga berada di luar cakupan; fitur ini hanya menyediakan mekanisme generik.

## Siklus hidup akun

### AUTH-014 Account closure & deletion (self-service)

- **Deskripsi:** Pengguna menutup dan menghapus akun unmanaged miliknya sendiri dengan masa tenggang, sesuai hak penghapusan data pribadi (GDPR/UU PDP).

- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform

- **Aktor & izin:** Pemilik akun unmanaged. Akun terkelola hanya dapat dihapus oleh admin organisasi (modul 02).

- **Aturan perilaku (acceptance criteria):**

1. Menu "Delete account" tersedia di Account settings > Account preferences dan memerlukan re-autentikasi (AUTH-009 butir 9).
2. Dialog konfirmasi menjelaskan dampak (semua akses app dan site hilang; konten yang dibuat seperti work item dan komentar tetap ada) dan mewajibkan pengguna mengetik ulang email untuk konfirmasi.
3. Setelah dikonfirmasi, akun berstatus `closed` dengan tanggal penghapusan terjadwal selama masa tenggang 14 hari (semua sesi, API token, dan token OAuth dicabut, dan akun tidak dapat mengakses app atau site mana pun); selama periode ini pemilik dapat membatalkan penghapusan lewat tautan di email konfirmasi atau dengan login, yang hanya membuka halaman "Batalkan penghapusan akun"; pembatalan mengembalikan status ke `active` tanpa memulihkan token yang sudah dicabut.
4. Setelah masa tenggang, data personal (nama, email, foto, atribut profil) dihapus permanen dan semua konten diatribusikan ke "Former user" dengan avatar netral.
5. `accountId` tetap tersimpan sebagai referensi anonim di work item, history, komentar, dan worklog sehingga integritas data dan laporan tetap terjaga.
6. Akun terkelola tidak menampilkan menu "Delete account" bagi pemiliknya; dialog menjelaskan bahwa hanya admin organisasi yang dapat menghapusnya.
7. Email akun yang telah selesai dihapus dapat dipakai lagi untuk mendaftar akun baru dengan `accountId` baru; selama masa tenggang email tersebut tetap terkunci.
8. Sistem menerbitkan event `account.closed` yang dapat dikonsumsi app terhubung (setara Personal Data Reporting API) agar app menghapus data personal yang mereka simpan.

- **Team-managed vs company-managed:** Tidak ada perbedaan.

- **Catatan & ketergantungan:**

- Event akun dan perilaku "Former user" pada referensi data di [28](28-model-data-inti.md); webhook/event untuk app di [23](23-api-publik-webhook.md); API data personal (USR-044) dan penghapusan oleh admin (USR-007/USR-008) di [02](02-pengguna-grup-organisasi.md).
- Repo saat ini memakai soft delete (`deleted_at`); penghapusan data personal memerlukan anonimisasi kolom, bukan hanya soft delete.

## Pertanyaan terbuka

1. Model sesi mana yang dipilih: Model A (access token ±15 menit + refresh token rotasi) atau Model B (JWT 12 jam sliding dengan `jti` yang dicek ke tabel sesi)? (AUTH-007, AUTH-008)
2. Apakah produk ini mendukung banyak site/organisasi dalam satu instalasi (account switcher, halaman "Select site") atau satu site per instalasi? (AUTH-001, AUTH-003)
3. Apakah `username` dari skema repo saat ini dihapus sepenuhnya atau dipertahankan sebagai handle tampilan non-unik? (AUTH-001, AUTH-008)
4. Pengecekan password bocor memakai layanan online k-anonymity (butuh akses internet) atau dataset lokal yang dapat diperbarui untuk instalasi tertutup? (AUTH-005, AUTH-007)
5. Penyedia CAPTCHA apa yang dipakai agar tetap ramah self-host dan privasi? (AUTH-006)
6. Apakah kebijakan autentikasi, enforcement 2SV, dan kontrol API token admin disertakan pada semua edisi atau hanya tier enterprise? (AUTH-010, AUTH-011, AUTH-012)
7. Penyedia login sosial mana yang diprioritaskan, dan pada fase berapa passkey serta login kode email diaktifkan? (AUTH-003)
8. Berapa lama masa transisi token lama sebelum cutover penolakan token tanpa `jti`/`kid`? (AUTH-008)
9. Apakah nilai `accountType` untuk manusia tetap `atlassian` demi kompatibilitas klien REST yang meniru produk acuan, atau memakai nilai netral (mis. `user`) dengan alias kompatibilitas? Pilihan ini juga menyangkut penghindaran merek pihak ketiga pada kontrak API produk ini. (AUTH-001)

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk keamanan identitas (mis. SSO connector, password manager, audit tambahan): tidak dibangun sebagai fitur native; produk ini hanya menyediakan OAuth 2.0 dan API token generik (AUTH-012, AUTH-013).
- **Fitur Data Center-only** (dc-only-or-deprecated): setting "Maximum Authentication Attempts", reset failed login count per pengguna oleh admin, direktori pengguna LDAP/Crowd internal, dan pengelolaan username/userKey legacy; disebut hanya sebagai rujukan (AUTH-006, AUTH-001).
- **Layanan identitas lintas produk produk acuan**: akun global lintas seluruh produk vendor acuan, admin hub organisasi lintas produk, dan dukungan pemulihan akun oleh vendor; produk ini hanya mengelola identitas di instalasinya sendiri.
- **Merge akun**: penggabungan dua akun tidak disediakan (sejalan dengan produk acuan).
- **Konfigurasi identity provider (SAML/OIDC), SCIM provisioning, domain terverifikasi, dan deaktivasi pengguna**: dibahas di modul [02](02-pengguna-grup-organisasi.md), bukan di modul ini.
- **Server otorisasi OAuth 2.0 dan akun app/bot**: dibahas di modul [23](23-api-publik-webhook.md) dan [25](25-integrasi-impor-ekspor-ekstensi.md).
