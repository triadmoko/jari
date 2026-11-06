# 30 — Service Management: Help Center, Portal, Customers & Kanal

Modul ini mendefinisikan seluruh **permukaan yang menghadap customer** pada Service Management di produk ini: help center tingkat situs (beranda, branding, layout, topics, announcement, pencarian, bahasa, multiple help centers, custom domain), portal per service project (portal groups, help resources, request form, deflection artikel KB, halaman detail request, daftar Requests, participants & sharing, approval dan CSAT dari sisi customer), **akun customer** (portal-only vs akun pengguna situs, signup/undangan/password, login, profil, akses anonim, SSO customer, administrasi akun), **customers & organizations** di service project, **customer notifications** (event, template, variabel, look & feel, pengirim), serta **kanal masuk** selain portal: email channel (koneksi mailbox, parsing, perlindungan loop, email log), chat Slack/Microsoft Teams, live chat, dan embeddable widget, ditutup dengan pengaturan hub channel, permission administrasi, dan REST API help center/customer. Mekanik umum (entitas service project, request type, forms, queue, SLA, knowledge base, approval steps, workflow engine, notification scheme internal, otomasi) dimiliki modul masing-masing; modul ini mengatur **perilaku dari sudut pandang customer dan kanal**.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`, JWT dengan claims di context `sess`). Belum ada service project, request type, akun customer, maupun infrastruktur email masuk. Seluruh fitur modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/helpcenter` (help center, portal, layout, announcement, topics, bahasa), `internal/features/customer` (akun portal-only, customers page, organizations, customer permissions), `internal/features/customernotification` (event, template, look & feel), dan `internal/features/channel` (email channel, chat, widget, live chat), dengan entitas bersama di `internal/shared/domain/entity`. UI customer dibangun sebagai route group terpisah di SPA SvelteKit (`frontend/src/routes`) dengan layout, branding, dan bahasa sendiri, terpisah dari UI agen; path kompatibilitas `/servicedesk/customer/...` tetap dilayani oleh fallback SPA `index.html` di `webui.RegisterRoutes`.

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — glosarium Help center/Portal/Customer, tier plan sebagai feature flag (VISI-003), penamaan produk (VISI-001).
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — model akun dengan `accountType` `customer` (AUTH-001), login, verifikasi email, password, CAPTCHA, authentication policies (AUTH-003 s.d. AUTH-011).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — direktori pengguna, domain verification dan managed accounts (USR-031, USR-032), self-signup pengguna internal (USR-012).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission project, peran Service Desk Team/Customers, issue security.
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — rich text editor untuk komentar portal dan announcement.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar publik vs internal, upload attachment (CMT-001, CMT-004, CMT-011).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status, transition, validator yang dipakai customer transitions.
- [12 JQL](12-jql-bahasa-kueri.md) — field `organizations`, `request-channel-type`, `Satisfaction`.
- [19 Notifikasi & Email](19-notifikasi-email.md) — infrastruktur SMTP/IMAP (NTF-023), custom email domain (NTF-026), notification scheme internal untuk agent (NTF-004).
- [20 Otomasi](20-otomasi.md) — action "Send email", reopen otomatis saat customer membalas.
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — bahasa & zona waktu (NAV-033), aksesibilitas (NAV-036), aplikasi mobile (NAV-050).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — look and feel situs (ADM-006), attachment settings (ADM-011), audit log (ADM-014), custom domain situs (ADM-055).
- [23 REST API & Webhook](23-api-publik-webhook.md) — konvensi REST, paginasi, error, autentikasi API token/OAuth 2.0 (API-001 s.d. API-010).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — aplikasi Slack dan Microsoft Teams untuk pengguna internal (INT-030 s.d. INT-036).
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) — i18n (NFR-021) dan aksesibilitas WCAG 2.1 AA (NFR-022).
- [28 Model Data Inti](28-model-data-inti.md) — entitas dan katalog event domain.
- 29 Service Project, Request Type, Forms & Agent (`29-jsm-service-project-request-type-formulir.md`) — service project, request type, field mapping, forms, raise on behalf dari tampilan agen.
- 31 Queues, SLA, Knowledge Base, CSAT & Laporan (`31-jsm-antrian-sla-knowledge-base-laporan.md`) — knowledge base, deflection report, laporan CSAT, queue per organization.
- 32 ITSM & Approvals (`32-jsm-itsm-insiden-problem-perubahan-layanan.md`) — konfigurasi approval steps.
- 33 Assets (`33-jsm-assets-cmdb.md`) — Assets object field di portal.
- 35 Virtual Service Agent, AI, Otomasi, API, Admin & CSM (`35-jsm-ai-otomasi-api-admin-csm.md`) — virtual service agent, jawaban AI di help center, Customer Service Management.

## Persona terkait

- **Customer/Pelapor** — mencari bantuan di help center, mengajukan request lewat portal, email, chat, atau widget, melacak status, berkomentar, berbagi request, memberi rating CSAT, dan mengelola profilnya.
- **Tamu/Kolaborator Eksternal** — pengunjung anonim yang membaca artikel publik dan mengajukan request tanpa login di portal yang dibuka untuk siapa pun.
- **Agent Layanan** — menambah customer dan organization, mengajukan request atas nama customer, membalas lewat email/chat, menangani live chat, dan membuka portal sebagai customer.
- **Admin Proyek/Project Lead** — admin service project: portal settings, portal groups, help resources, customer permissions, customer notifications, email channel, chat, widget, live chat, dan bahasa portal.
- **Admin Situs** — admin produk dan organisasi: branding dan layout help center, topics, announcement, signup & akses, multiple help centers, custom domain, akun customer, SSO customer, dan domain pengirim email.
- **Manajer Layanan/ITSM** — menetapkan kebijakan akses customer, organisasi klien, nada komunikasi email, dan memantau CSAT serta kesehatan kanal.
- **Stakeholder/Viewer** — approver (berlisensi maupun customer) yang menyetujui atau menolak request dari help center atau email.
- **Integrator/Developer API** — memakai REST API help center/customer untuk membuat request, mengelola customer dan organization, serta menanamkan widget.

## User story modul

1. Sebagai Customer/Pelapor, saya ingin satu help center yang hanya menampilkan portal yang boleh saya akses beserta pencarian lintas portal agar saya cepat menemukan formulir atau artikel yang tepat (JSM-PRT-001, JSM-PRT-006).
2. Sebagai Admin Situs, saya ingin mengatur nama, logo, warna, layout, dan topics help center agar pengalaman customer konsisten dengan identitas perusahaan (JSM-PRT-002, JSM-PRT-003, JSM-PRT-004).
3. Sebagai Admin Situs, saya ingin membuat beberapa help center dengan custom domain masing-masing agar setiap brand atau audiens (internal vs eksternal) punya pintu masuk sendiri (JSM-PRT-008, JSM-PRT-009).
4. Sebagai Customer/Pelapor, saya ingin mendaftar, menerima undangan, login, dan mereset password tanpa memerlukan akun penuh produk agar saya bisa meminta bantuan dengan mudah (JSM-PRT-012, JSM-PRT-013, JSM-PRT-014).
5. Sebagai Admin Proyek/Project Lead, saya ingin memilih apakah portal hanya untuk customer terdaftar, untuk semua customer berakun, atau terbuka untuk siapa pun agar akses sesuai kebijakan layanan (JSM-PRT-015, JSM-PRT-016, JSM-PRT-017).
6. Sebagai Customer/Pelapor, saya ingin mengisi request form dengan bantuan teks, lampiran, dan saran artikel KB agar masalah saya bisa selesai sendiri atau tercatat lengkap (JSM-PRT-024, JSM-PRT-025).
7. Sebagai Customer/Pelapor, saya ingin melihat status ramah customer, aktivitas publik, dan semua request saya di satu daftar agar saya tidak perlu menanyakan progres ke agent (JSM-PRT-026, JSM-PRT-027, JSM-PRT-029).
8. Sebagai Customer/Pelapor, saya ingin membagikan request ke rekan atau seluruh organisasi saya agar pihak terkait ikut memantau dan berkomentar (JSM-PRT-030, JSM-PRT-034).
9. Sebagai Stakeholder/Viewer yang menjadi approver, saya ingin menyetujui atau menolak request dari help center atau langsung dari email agar keputusan tidak tertunda (JSM-PRT-031).
10. Sebagai Agent Layanan, saya ingin menambah customer dan organization ke project serta mengajukan request atas nama customer agar permintaan lewat telepon atau tatap muka tetap tercatat (JSM-PRT-033, JSM-PRT-024).
11. Sebagai Admin Proyek/Project Lead, saya ingin mengatur event, template, variabel, dan identitas pengirim customer notification agar email ke customer jelas dan aman (JSM-PRT-036 s.d. JSM-PRT-039).
12. Sebagai Customer/Pelapor, saya ingin mengirim email ke alamat dukungan dan membalas email notifikasi agar percakapan otomatis tercatat di request tanpa membuka portal (JSM-PRT-040, JSM-PRT-041).
13. Sebagai Admin Proyek/Project Lead, saya ingin email log, perlindungan loop, dan status koneksi mailbox agar kegagalan kanal email cepat terdeteksi (JSM-PRT-042, JSM-PRT-043).
14. Sebagai Customer/Pelapor, saya ingin meminta bantuan dari Slack, Microsoft Teams, live chat, atau widget di situs web agar saya bisa dibantu di tempat saya bekerja (JSM-PRT-044 s.d. JSM-PRT-047).
15. Sebagai Integrator/Developer API, saya ingin REST API untuk request customer, customer, dan organization agar sistem lain dapat membuat request dan menyinkronkan data klien (JSM-PRT-050).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 19 |
| P1 | 19 |
| P2 | 11 |
| P3 | 1 |
| **Total** | **50** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 0 |
| Fase 4 | 0 |
| Fase 5 | 46 |
| Fase 6 | 0 |
| Fase 7 | 4 |
| Fase — | 0 |

Strategi urutan implementasi: seluruh modul berada di Fase 5 karena bergantung pada service project dan request type (modul 29) serta infrastruktur email (modul 19); di dalam Fase 5 urutannya dibagi menjadi tiga gelombang. Gelombang pertama (semua P0) membangun akun customer portal-only, login/signup, customer permissions, help center beranda dan portal, portal groups, request form, halaman detail, daftar Requests, participants & sharing, customers page, organizations, customer notifications dasar, saran artikel KB (deflection) di portal, dan email channel beserta parsing, sehingga satu service project sudah dapat menerima dan menjawab request lewat portal dan email. Gelombang kedua (P1) menambahkan branding, announcement, bahasa, SSO customer, customer transitions, status ramah customer, approval di portal, CSAT, profil, administrasi akun, variabel template, identitas pengirim, perlindungan loop, email log, aksesibilitas, REST API, dan permission administrasi. Gelombang ketiga (P2/P3) melengkapi layout editor dan custom pages, topics, multiple help centers, custom domain, akses anonim, look & feel email, widget, hub channel, dan akses mobile. Fase 7 menampung kanal chat Slack/Microsoft Teams, live chat, dan detail profil customer/organization karena berkaitan erat dengan virtual service agent dan Customer Service Management (modul 35). Inventaris memuat lima belas fitur ganda antara domain lampiran JSM dan domain JSM rinci; semuanya digabung ke heading yang lebih rinci dengan baris "Alias inventaris".

## Fitur

## Help center: beranda, tampilan & jangkauan

### JSM-PRT-001 Help Center (home page)

- **Alias inventaris:** Help center
- **Deskripsi:** Halaman beranda tingkat situs yang menjadi pintu masuk tunggal customer ke semua portal service project yang boleh diakses, berisi pencarian global, announcement, topics, artikel KB, serta navigasi ke Requests dan Approvals.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor dan pengunjung anonim (melihat sesuai akses), Agent Layanan (melihat sebagai customer), Admin Situs (mengonfigurasi; global permission Administer), Admin Proyek (konten portal).
- **Aturan perilaku (acceptance criteria):**
  1. URL default help center adalah `<base-url>/servicedesk/customer/portals` (path kompatibilitas) dengan alias pendek `/help`; bila custom domain aktif (JSM-PRT-009), help center juga dilayani di root domain tersebut.
  2. Header memuat logo, nama help center, menu Requests, menu Approvals, dan avatar profil dengan menu Profile, pilihan bahasa, dan Log out; menu Requests/Approvals hanya tampil untuk pengguna yang login.
  3. Daftar portal hanya berisi portal yang boleh diakses pengguna saat ini menurut customer permissions (JSM-PRT-016); portal tanpa akses tidak dirender dan tidak bocor lewat API.
  4. Search bar global mencari lintas semua portal yang dapat diakses dan mengembalikan request type serta artikel KB (perilaku rinci di JSM-PRT-006).
  5. Help center announcement (bila berisi) tampil di bagian atas beranda sebelum section lain (JSM-PRT-005).
  6. Section Topics dan featured portals dirender sesuai layout yang dipublikasikan (JSM-PRT-003, JSM-PRT-004); tanpa kustomisasi, layout default berisi Search, Announcement, dan Portals.
  7. Footer menampilkan teks "Powered by <nama produk>" yang berasal dari konfigurasi produk (VISI-001), bukan merek pihak lain.
  8. Pengunjung anonim hanya melihat portal yang mengizinkan akses tanpa login; bila tidak ada portal semacam itu, ia diarahkan ke login page (JSM-PRT-012) dengan parameter tujuan.
  9. ⚠ Bila customer yang login hanya punya akses ke tepat satu portal, help center langsung mengarahkan ke portal tersebut; perilaku ini dapat dimatikan lewat pengaturan help center.
  10. Breadcrumb "Help Center > Portal > Request" tampil di halaman portal dan detail request, dan setiap segmen dapat diklik.
  11. Bila customer yang login belum punya akses ke portal mana pun, beranda menampilkan empty state dengan pesan "Anda belum memiliki akses ke portal mana pun" dan tautan kontak admin.
  12. ⚠ Produk ini hanya membangun satu generasi pengalaman setara "Modern help center"; perilaku khusus Legacy help center (mis. landing pages) dicatat hanya sebagai rujukan impor (JSM-PRT-003).
  13. Beranda responsif di browser mobile mulai lebar 320 px tanpa scroll horizontal (JSM-PRT-011).
  14. Beranda yang sudah dikustomisasi dapat memuat link Requests dan Approvals sebagai kartu pintasan selain di header.
  15. Branding (nama, logo, banner, warna top bar dan link, judul halaman) diterapkan di beranda sesuai JSM-PRT-002.
  16. ⚠ Layout editor menyusun section beranda (hero/search, portals, topics, announcements, artikel, rich text) sesuai JSM-PRT-003; daftar jenis section final mengikuti JSM-PRT-003.
  17. Topics mengurasi portal dan artikel KB per tema dan ditampilkan sebagai kartu di beranda (JSM-PRT-004).
  18. Global search di beranda mencakup request type dan artikel KB lintas portal yang dapat diakses (JSM-PRT-006).
  19. ⚠ Custom pages tambahan dapat ditautkan dari beranda (perilaku dan tier di JSM-PRT-003).
  20. Seluruh teks konfigurasi beranda ditampilkan dalam bahasa customer bila terjemahan tersedia, dengan fallback ke bahasa default (JSM-PRT-007).
  21. Menu profil customer berisi tautan Profile, pemilih bahasa, dan Log out (JSM-PRT-019).
- **Team-managed vs company-managed:** Tidak ada perbedaan; portal dari service project team-managed dan company-managed tampil dengan cara yang sama.
- **Catatan & ketergantungan:** Bergantung pada service project dan portal di modul 29 dan JSM-PRT-021. Route customer di SPA harus terpisah dari layout agen (sidebar dan top navigation [21](21-navigasi-ux-preferensi.md) tidak tampil). Endpoint beranda mengembalikan data portal sudah difilter izin agar tidak ada kebocoran nama portal.

### JSM-PRT-002 Help center branding

- **Deskripsi:** Pengaturan identitas visual help center (nama, logo, banner, warna) yang berlaku di beranda, semua portal, dan login page.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (global permission Administer); Admin Proyek hanya untuk logo portal.
- **Aturan perilaku (acceptance criteria):**
  1. Nama help center dapat diubah (1-100 karakter, wajib) dan tampil di header serta judul tab browser (`<nama halaman> - <nama help center>`).
  2. Logo help center dapat diunggah dalam format PNG, JPG, GIF, atau SVG maksimal 2 MB; UI merekomendasikan rasio horizontal dan tinggi tampil 40 px, dan file di luar format/ukuran ditolak dengan pesan galat.
  3. Banner/hero image beranda dapat diunggah (PNG/JPG maksimal 5 MB, rekomendasi 1920×400 px) dan dapat dihapus untuk kembali ke latar warna polos.
  4. Warna top bar/header, latar banner, link, dan tombol/primary color diatur dengan color picker atau input hex `#RRGGBB`; input tidak valid ditolak.
  5. Warna teks banner dapat dipilih terang atau gelap; sistem menampilkan peringatan bila rasio kontras teks terhadap latar di bawah 4.5:1 (JSM-PRT-010).
  6. Perubahan dapat dilihat di preview sebelum disimpan; menutup editor tanpa menyimpan membuang perubahan.
  7. Aksi "Reset to default" mengembalikan semua nilai branding ke default produk setelah konfirmasi.
  8. Logo portal per project (JSM-PRT-021) menimpa logo help center hanya di halaman portal tersebut dan halaman request-nya.
  9. Pada multiple help centers (JSM-PRT-008), branding disimpan terpisah per help center dan perubahan satu help center tidak memengaruhi yang lain.
  10. Tidak ada editor custom HTML/CSS/JavaScript; field yang menerima teks disanitasi sehingga skrip tidak pernah dieksekusi.
  11. Hanya Admin Situs yang dapat mengubah branding; pengguna lain mendapat 403 dari API dan tidak melihat menu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan file memakai layanan attachment ([08](08-komentar-lampiran-aktivitas.md), [22](22-administrasi-sistem.md) ADM-011). Look and feel aplikasi agen (ADM-006) terpisah dari branding help center. Kustomisasi CSS penuh hanya lewat ekstensi (di luar cakupan).

### JSM-PRT-003 Help center layout customization

- **Deskripsi:** Editor tata letak help center untuk menyusun section (search, announcements, topics, portals, konten kustom) di beranda maupun di custom pages tambahan.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengedit dan mempublikasikan), Customer/Pelapor (melihat hasil).
- **Aturan perilaku (acceptance criteria):**
  1. Mode "Customize" dibuka dari Help centers > Customize help center atau ikon pensil di beranda, dan hanya tampil untuk Admin Situs.
  2. Section dapat ditambah, dihapus, dan diurutkan dengan drag-and-drop maupun tombol naik/turun (alternatif keyboard).
  3. Section Search berupa hero dengan banner (JSM-PRT-002) dan search bar; maksimal satu per halaman.
  4. Section Announcements menampilkan help center announcement (JSM-PRT-005).
  5. Section Topics menampilkan topik yang dipilih beserta urutannya (JSM-PRT-004).
  6. Section Portals dapat dikonfigurasi menampilkan semua portal atau hanya featured portals.
  7. ⚠ Section konten kustom tersedia dalam jenis rich text, gambar, dan quick links (judul + URL); konten disanitasi.
  8. ⚠ Section artikel KB menampilkan artikel populer atau terbaru dari space KB yang dapat diakses viewer.
  9. Jumlah portal yang tampil di section Portals dapat diatur (mis. 6, 9, 12) dan bila lebih banyak tampil tautan "View all".
  10. Perubahan disimpan sebagai draft; tersedia aksi Publish (berlaku untuk customer) dan Discard (kembali ke versi terpublikasi).
  11. Preview menampilkan halaman sebagai customer (dengan opsi meniru customer login vs anonim) sebelum publish.
  12. Tab Pages di editor memungkinkan membuat, menduplikasi, mengganti nama, mengedit, menata, dan menghapus custom pages (mis. halaman new-starter atau kebijakan).
  13. Custom pages memakai komponen section yang sama dengan editor beranda (hero, rich text, topics, portals, artikel KB, gambar).
  14. Setiap custom page punya slug unik per help center (huruf kecil, angka, tanda hubung) yang membentuk URL `/help/pages/{slug}`, dan dapat ditautkan dari navigasi atau beranda.
  15. Setiap custom page memiliki status draft/published dan preview sendiri; page draft mengembalikan 404 bagi customer.
  16. ⚠ Visibilitas custom page dapat dipilih per audiens: semua pengunjung, hanya customer login, atau audiens tertentu (organization/group).
  17. ⚠ Konten custom page dapat diterjemahkan per bahasa help center (JSM-PRT-007) dengan fallback ke bahasa default.
  18. Halaman topic tersedia di tier Standard, sedangkan custom pages hanya aktif bila feature flag tier Premium/Enterprise menyala (VISI-003).
  19. Landing pages gaya Legacy help center tidak dibangun sebagai fitur terpisah; data landing pages hasil impor dikonversi menjadi custom pages.
  20. ⚠ Batas default 100 custom pages per help center; percobaan membuat page ke-101 ditolak dengan pesan galat.
  21. Custom pages adalah satu-satunya mekanisme halaman tambahan dan menggantikan landing pages.
  22. Pada multiple help centers, layout dan custom pages disimpan per help center.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rich text memakai editor [07](07-editor-konten-kaya.md). Artikel KB dari modul 31. Model penyimpanan layout disarankan berupa dokumen JSON berversi (draft dan published) per help center.

### JSM-PRT-004 Help center topics & featured portals

- **Deskripsi:** Pengelompokan portal dan request type ke dalam topik, serta penonjolan portal tertentu di beranda help center.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengelola), Customer/Pelapor (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Topik dibuat dengan nama (wajib, maksimal 100 karakter), deskripsi opsional, dan ikon/gambar opsional dari pustaka ikon atau unggahan.
  2. Portal dan/atau request type spesifik dapat ditambahkan ke topik lewat pemilih yang hanya menampilkan item dari help center yang sama.
  3. Satu portal atau request type dapat masuk ke lebih dari satu topik.
  4. Urutan topik dan urutan item di dalam topik dapat diatur dengan drag-and-drop.
  5. Topik dapat diubah dan dihapus; menghapus topik tidak menghapus portal atau request type di dalamnya.
  6. Setiap topik punya halaman topik (`/help/topics/{id}`) yang menampilkan semua item topik itu; tersedia di tier Standard.
  7. Admin dapat memilih featured portals beserta urutannya untuk ditampilkan paling atas di section Portals.
  8. Item yang tidak boleh diakses viewer otomatis disembunyikan dari topik saat dirender.
  9. Topik yang tidak memiliki item yang dapat diakses viewer tidak tampil sama sekali.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pengecekan akses item memakai aturan yang sama dengan JSM-PRT-016 dan JSM-PRT-023.

### JSM-PRT-005 Announcements (help center, login, portal)

- **Alias inventaris:** Portal & help center announcements
- **Deskripsi:** Pesan pengumuman rich text untuk customer di beranda help center, login page, dan portal setiap project, misalnya untuk menginformasikan gangguan atau info layanan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (help center dan login), Admin Proyek (portal), Customer/Pelapor dan pengunjung anonim (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Help center announcement terdiri dari judul dan isi rich text dan tampil di beranda.
  2. Login announcement tampil di login page (mis. instruksi SSO atau akun).
  3. Portal announcement diatur per service project oleh Admin Proyek dan tampil di halaman portal beserta halaman create request.
  4. Format rich text yang didukung: bold, italic, link, dan list (bullet/numbered); elemen lain dibuang saat sanitasi.
  5. Mengosongkan judul dan isi announcement menyembunyikannya tanpa menghapus konfigurasi lain.
  6. Announcement dapat diterjemahkan untuk setiap bahasa tambahan yang aktif (JSM-PRT-007), dengan fallback ke bahasa default.
  7. Announcement tampil untuk pengunjung anonim maupun customer yang login.
  8. ⚠ Panjang judul maksimal 100 karakter dan isi maksimal 1.000 karakter teks; penyimpanan di atas batas ditolak dengan pesan galat.
  9. Help center dan login announcement hanya dapat diedit Admin Situs; portal announcement oleh Admin Proyek project tersebut (dan Admin Situs).
  10. Announcement global disimpan per help center pada multiple help centers.
  11. Perubahan announcement berlaku segera tanpa perlu publish layout dan tercatat di audit log (ADM-014).
- **Team-managed vs company-managed:** Tidak ada perbedaan; lokasi menu portal announcement berada di Portal settings untuk keduanya.
- **Catatan & ketergantungan:** Editor teks dari [07](07-editor-konten-kaya.md). Section Announcements pada layout di JSM-PRT-003.

### JSM-PRT-006 Help center search

- **Deskripsi:** Pencarian terpadu di help center dan portal yang menampilkan request type serta artikel knowledge base yang relevan dan boleh diakses customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, pengunjung anonim (hasil sesuai akses).
- **Aturan perilaku (acceptance criteria):**
  1. Search dari beranda mencari lintas semua portal yang dapat diakses viewer.
  2. Search di dalam portal hanya mengembalikan request type dan artikel KB milik portal tersebut.
  3. Hasil dikelompokkan menjadi dua bagian: Request types dan Knowledge base articles, masing-masing dengan jumlah hasil.
  4. Request type tanpa portal group atau yang dibatasi (JSM-PRT-023) tidak muncul di hasil.
  5. Artikel KB hanya tampil bila viewer berhak melihat space KB tersebut (anonim, semua customer, atau hanya pengguna berlisensi) sesuai pengaturan modul 31.
  6. Hasil dapat berasal dari banyak space KB yang ditautkan ke portal-portal yang dapat diakses.
  7. Saran instan (typeahead) muncul setelah 2 karakter dengan debounce sekitar 300 ms dan maksimal 5 saran per kelompok.
  8. Hasil artikel menampilkan cuplikan teks dengan kata kunci yang di-highlight.
  9. Bila tidak ada hasil, halaman menampilkan pesan "Tidak ada hasil" dan ajakan untuk raise request (tautan ke portal atau request type default).
  10. ⚠ Help resources (link eksternal, JSM-PRT-022) ikut muncul di kelompok Request types dan membuka URL eksternal saat diklik.
  11. ⚠ Jawaban AI/semantic search di help center hanya tersedia bila modul virtual service agent/AI (modul 35) diaktifkan; tanpa itu pencarian berbasis kata kunci.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Indeks pencarian KB dan request type dibangun di modul 31/29. Pemfilteran izin dilakukan di server sebelum hasil dikembalikan.

### JSM-PRT-007 Help center languages & translations

- **Deskripsi:** Dukungan multi-bahasa agar teks help center, portal, request form, dan notifikasi tampil sesuai bahasa customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (teks help center), Admin Proyek (Language support project), Customer/Pelapor (memilih bahasa).
- **Aturan perilaku (acceptance criteria):**
  1. Bahasa default help center mengikuti bahasa default situs (minimal EN dan ID sesuai NFR-021).
  2. Admin Situs dapat menerjemahkan nama help center, help center announcement, dan login message per bahasa.
  3. Admin Proyek mengaktifkan bahasa tambahan di halaman Language support project; hanya bahasa yang didukung produk yang dapat dipilih.
  4. Nama dan deskripsi portal, portal announcement, dan nama portal group dapat diterjemahkan per bahasa aktif.
  5. Nama, deskripsi, nama field, dan help text request type dapat diterjemahkan.
  6. Judul dan deskripsi help resources dapat diterjemahkan.
  7. Customer-visible status names (JSM-PRT-027) dapat diterjemahkan.
  8. Template customer notification dapat diterjemahkan per bahasa (JSM-PRT-036).
  9. Teks yang belum diterjemahkan fallback ke bahasa default tanpa menampilkan kunci kosong.
  10. Customer login melihat bahasa yang dipilih di profilnya; pengunjung anonim mengikuti header `Accept-Language` browser (sejalan dengan `middleware.LanguageMiddleware()`), lalu fallback ke default.
  11. Teks UI sistem (tombol, label, pesan galat) disediakan produk dalam semua bahasa yang didukung tanpa konfigurasi admin.
  12. ⚠ Terjemahan opsi field (select/radio) dan isi forms terbatas; versi awal hanya menerjemahkan nama/help text field, bukan opsi.
  13. Halaman Language support menampilkan indikator kelengkapan terjemahan per bahasa (jumlah teks yang belum diterjemahkan).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Infrastruktur i18n di NFR-021 ([27](27-nfr-arsitektur-platform.md)) dan preferensi bahasa pengguna di NAV-033 ([21](21-navigasi-ux-preferensi.md)). Pesan galat backend sudah mendukung EN/ID lewat `internal/shared/constants/error.go`.

### JSM-PRT-008 Multiple help centers

- **Deskripsi:** Membuat lebih dari satu help center per situs, misalnya per brand, region, atau internal vs eksternal, masing-masing dengan branding, URL, portal, dan KB sendiri.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs saja.
- **Aturan perilaku (acceptance criteria):**
  1. Help center baru dibuat dengan nama dan slug/URL unik (`/help/{slug}`); slug duplikat ditolak, dan slug yang bentrok dengan path sistem (`pages`, `topics`, `portal`, `requests`, `approvals`, `login`, `signup`, `profile`) ditolak dengan pesan galat.
  2. Halaman daftar help center menampilkan semua help center dengan penanda help center default.
  3. Service project (portal) dapat ditautkan atau dilepaskan dari help center lewat pemilih project.
  4. ⚠ Satu project dapat ditautkan ke lebih dari satu help center sekaligus.
  5. Branding, layout, custom pages, topics, announcement, dan bahasa diatur per help center.
  6. Konten KB per help center mengikuti project yang ditautkan.
  7. Setiap help center dapat dipetakan ke custom domain sendiri (JSM-PRT-009).
  8. Help center non-default dapat dihapus setelah konfirmasi; project yang tertaut tetap ada dan tetap tampil di help center default.
  9. ⚠ Jumlah help center per situs dibatasi konfigurasi tier (default 20 untuk Premium/Enterprise); fitur nonaktif di tier Free/Standard.
  10. Hanya Admin Situs yang dapat membuat, mengubah, dan menghapus help center.
  11. Help center default tidak dapat dihapus; tombol hapus tidak tampil dan API mengembalikan 400.
  12. Pengaturan akses (signup, anonim) dapat berbeda per help center.
  13. Pengaturan KB yang ditampilkan dapat dikonfigurasi per help center.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Semua entitas help center (branding, layout, topic, announcement) wajib punya `help_center_id` sejak awal walau Fase 5 gelombang pertama hanya mengaktifkan help center default, agar tidak perlu migrasi skema belakangan.

### JSM-PRT-009 Custom domain for help center

- **Deskripsi:** Menyajikan help center di domain milik organisasi (mis. `support.company.com`) dengan sertifikat TLS yang dikelola otomatis oleh produk.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (organisasi dan produk).
- **Aturan perilaku (acceptance criteria):**
  1. Custom domain ditambahkan di konsol administrasi (Products > Custom domains) dengan input nama host.
  2. Kepemilikan domain diverifikasi lewat DNS record CNAME ke host produk (dan TXT bila diperlukan); UI menampilkan nilai record yang harus dibuat.
  3. Sertifikat TLS diterbitkan dan diperpanjang otomatis (mis. ACME/Let's Encrypt) tanpa input admin.
  4. Status domain ditampilkan sebagai pending, active, atau failed, dengan pesan galat konfigurasi DNS yang spesifik (record tidak ditemukan, nilai salah, sertifikat gagal).
  5. Domain ditetapkan ke help center tertentu, termasuk saat multiple help centers aktif; satu domain hanya untuk satu help center.
  6. Link di customer notification memakai custom domain bila aktif.
  7. ⚠ URL bawaan produk tetap berfungsi dan mengalihkan (HTTP 301) ke custom domain untuk halaman customer.
  8. Domain dapat dihapus atau diganti; penghapusan mengembalikan link ke URL bawaan.
  9. ⚠ Hanya subdomain yang didukung (bukan apex domain) dan jumlah domain per situs dibatasi konfigurasi tier (default 10).
  10. Fitur hanya aktif bila feature flag tier berbayar menyala (VISI-003).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Custom domain situs untuk aplikasi agen di ADM-055 ([22](22-administrasi-sistem.md)). Untuk deployment self-hosted, verifikasi DNS dan TLS bisa didelegasikan ke reverse proxy; keputusan di Pertanyaan terbuka.

### JSM-PRT-010 Portal accessibility

- **Deskripsi:** Standar aksesibilitas help center dan portal agar dapat dipakai semua customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor; tidak ada konfigurasi.
- **Aturan perilaku (acceptance criteria):**
  1. Semua halaman customer memenuhi target WCAG 2.1 AA, diverifikasi dengan audit otomatis (mis. axe) di CI dan audit manual sebelum rilis.
  2. Seluruh fungsi dapat dijalankan dengan keyboard saja dan indikator fokus selalu terlihat.
  3. Tersedia skip link "Langsung ke konten" dan landmark ARIA (header, nav, main, footer).
  4. Setiap field form punya label terhubung, dan pesan galat validasi diumumkan ke screen reader (`aria-live`) serta ditautkan ke field-nya.
  5. Logo dan banner memiliki alt text (default dari nama help center, dapat diubah admin).
  6. Kombinasi warna branding divalidasi kontrasnya (JSM-PRT-002) dan tombol/link default memenuhi rasio 4.5:1.
  7. Zoom hingga 200% dan perbesaran teks tidak merusak layout atau memotong konten.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Standar umum di NFR-022 ([27](27-nfr-arsitektur-platform.md)) dan NAV-036 ([21](21-navigasi-ux-preferensi.md)).

### JSM-PRT-011 Portal access via mobile

- **Deskripsi:** Akses customer ke help center dan portal dari perangkat mobile.
- **Prioritas / Fase / Tier:** P3 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Help center dan portal responsif di browser mobile (lebar 320-768 px) dengan menu header yang dilipat.
  2. Raise request, komentar, dan unggah attachment (termasuk dari kamera/galeri) dapat dilakukan dari web mobile.
  3. Approve/decline dari email approval berfungsi di klien email mobile (tautan membuka halaman konfirmasi yang responsif).
  4. ⚠ Aplikasi mobile native (NAV-050) hanya untuk agent dan pengguna berlisensi; customer portal-only tidak dapat login ke aplikasi native.
  5. Push notification hanya dikirim ke agent; customer menerima update lewat email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aplikasi mobile di NAV-050 ([21](21-navigasi-ux-preferensi.md)); push notification di NTF-038 ([19](19-notifikasi-email.md)).

## Akun customer, login & akses

### JSM-PRT-012 Help center login page

- **Deskripsi:** Halaman login customer yang mendeteksi tipe akun dari email, lalu menangani login dengan password portal-only, akun pengguna situs, atau SSO.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, pengunjung anonim; tidak memerlukan izin.
- **Aturan perilaku (acceptance criteria):**
  1. Langkah pertama hanya meminta email; setelah submit, server menentukan alur login (portal-only, akun pengguna situs, atau SSO) tanpa mengungkap apakah email terdaftar bila akses tidak diizinkan.
  2. Akun portal-only login dengan email dan password di halaman help center (bukan halaman login aplikasi agen).
  3. Akun pengguna situs (setara akun global, AUTH-001) diarahkan ke alur login utama produk (AUTH-003), termasuk login sosial bila diaktifkan, lalu dikembalikan ke help center.
  4. Bila authentication policy mewajibkan SSO (AUTH-011, JSM-PRT-018), pengguna diarahkan ke identity provider dan tidak ditawari password.
  5. Link "Forgot password" mengirim email reset untuk akun portal-only (JSM-PRT-013); respons selalu sama ("Jika email terdaftar, kami telah mengirim tautan") untuk mencegah enumerasi akun.
  6. Link "Sign up" hanya tampil bila public signup diizinkan (JSM-PRT-015).
  7. Login announcement (JSM-PRT-005) dan branding help center (JSM-PRT-002) tampil di halaman login.
  8. ⚠ CAPTCHA ditampilkan otomatis setelah beberapa percobaan login gagal berturut-turut sesuai perlindungan platform (AUTH-006), bukan setting yang diatur admin.
  9. Setelah login berhasil, pengguna dikembalikan ke halaman tujuan (deep link, parameter `destination` yang divalidasi hanya untuk path internal).
  10. Kredensial salah atau akun tidak ditemukan selalu menghasilkan pesan generik "Email atau password salah"; pesan "Akun Anda tidak memiliki akses ke help center ini" dan "Akun dinonaktifkan, hubungi admin" hanya ditampilkan setelah kredensial terbukti valid, sehingga konsisten dengan butir 1.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sesi customer memakai mekanisme JWT yang sama ([01](01-identitas-autentikasi-sesi.md) AUTH-007) dengan claim `accountType=customer` dan audience help center agar token customer tidak diterima endpoint agen. Repo saat ini hanya punya `/api/v1/auth/login` untuk tabel `users`; perlu endpoint login customer terpisah.

### JSM-PRT-013 Customer signup, invitation & password flow

- **Deskripsi:** Alur pembuatan akun portal-only lewat signup mandiri atau undangan, termasuk verifikasi email dan pengelolaan password.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, Agent Layanan (mengundang bila diizinkan), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Form signup meminta email, nama lengkap, dan password; semua wajib dan email divalidasi formatnya.
  2. Setelah signup, sistem mengirim email verifikasi berisi tautan aktivasi; akun berstatus belum terverifikasi dan tidak bisa login sampai tautan diklik.
  3. Agent/admin yang menambah customer (JSM-PRT-033) memicu email undangan berisi tautan set password.
  4. ⚠ Undangan dapat dikirim ulang dari halaman Customers per customer yang belum aktif; pengiriman ulang membatalkan tautan lama.
  5. ⚠ Tautan undangan berlaku 7 hari dan tautan reset password 24 jam (dapat dikonfigurasi); tautan kedaluwarsa menampilkan pesan dengan tombol minta tautan baru.
  6. ⚠ Akun yang dibuat otomatis dari email channel atau widget tidak menerima undangan saat dibuat; saat pemiliknya pertama kali membuka portal, ia diminta menetapkan password lewat email verifikasi.
  7. Reset password lewat "Forgot password" mengirim tautan sekali pakai; setelah dipakai, semua sesi lama akun tersebut dicabut.
  8. ⚠ Password portal-only mengikuti kebijakan kekuatan password produk (AUTH-005): minimal 8 karakter dan ditolak bila terlalu lemah atau termasuk daftar password bocor.
  9. CAPTCHA wajib pada signup publik (JSM-PRT-015).
  10. Bila public signup dimatikan, endpoint signup mengembalikan 403 dan form tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai ulang verifikasi email AUTH-004 dan manajemen password AUTH-005 ([01](01-identitas-autentikasi-sesi.md)). Email undangan dan reset adalah email akun sistem (NTF-027), bukan customer notification.

### JSM-PRT-014 Customer account types (portal-only vs Atlassian account)

- **Deskripsi:** Dua jenis akun customer: portal-only yang terikat ke satu situs dan hanya bisa mengakses help center, dan akun pengguna situs (padanan Atlassian account) yang juga dipakai pengguna berlisensi saat bertindak sebagai customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, Agent Layanan, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Akun portal-only (`accountType=customer`, AUTH-001) dapat dibuat lewat signup, undangan, email channel, widget, chat, atau REST API (JSM-PRT-050).
  2. Akun portal-only hanya berlaku di satu situs dan ditolak (403) saat mengakses UI atau API agen, project, board, maupun modul lain.
  3. Akun pengguna situs (termasuk pengguna internal) dapat menjadi customer tanpa memakai lisensi agent.
  4. Customer tidak memakai lisensi dan jumlahnya tidak dibatasi tier; jumlah customer tidak dihitung di halaman lisensi.
  5. ⚠ Satu email hanya dapat memiliki satu akun per situs: bila email sudah terdaftar sebagai akun pengguna situs, pembuatan portal-only dengan email yang sama ditolak dan akun yang ada dipakai sebagai customer.
  6. Email dengan domain terverifikasi/diklaim organisasi (USR-031, USR-032) dibuat atau dikonversi sebagai managed account, bukan portal-only, sesuai kebijakan domain.
  7. Status akun customer: active, deactivated, atau unverified; akun deactivated tidak bisa login namun request lamanya tetap ada.
  8. Halaman Customers menampilkan penanda tipe akun (Portal-only / Pengguna) untuk setiap customer.
  9. Agent dapat membuka help center/portal dan raise request sebagai dirinya sendiri.
  10. ⚠ Di project tempat ia menjadi agent, request yang ia buat untuk dirinya tetap mencatatnya sebagai reporter dan halaman detail portal menampilkan tautan "Buka di tampilan agen".
  11. Di project lain, pengguna berlisensi diperlakukan sebagai customer dan memerlukan customer permission project atau akses terbuka.
  12. Agent yang bertindak sebagai customer di portal hanya melihat request miliknya, tempat ia menjadi participant, atau yang dibagikan ke organization-nya, bukan seluruh request project.
  13. ⚠ Pengguna berlisensi yang menjadi reporter menerima notification scheme internal (NTF-004), bukan customer notification; keputusan akhir dicatat di Pertanyaan terbuka.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keputusan skema data: akun customer disarankan disimpan di tabel `users` yang sama dengan kolom `account_type` dan `site_scope` agar relasi reporter/participant seragam (lihat Pertanyaan terbuka). Migrasi tabel `users` saat ini perlu menambah kolom `account_type` dan status verifikasi.

### JSM-PRT-015 Customer access & signup settings (site/help center level)

- **Alias inventaris:** Customer access & authentication
- **Deskripsi:** Pengaturan global yang menentukan siapa yang boleh membuat akun dan mengakses help center, termasuk akses tanpa login, serta ringkasan cara customer login.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs saja; Customer/Pelapor terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi signup: "Anyone can sign up" (public signup) atau "Only agents/admins can add customers" (invite only); default invite only.
  2. Verifikasi email wajib untuk setiap akun baru yang dibuat lewat signup dan tidak dapat dimatikan.
  3. CAPTCHA aktif pada form signup publik dan halaman publik yang menerima input (request anonim, widget).
  4. Toggle "Allow access without login" mengizinkan akses help center dan raise request tanpa login; toggle ini prasyarat opsi "Anyone" di customer permissions project (JSM-PRT-016).
  5. Pengaturan signup berlaku untuk semua project yang mengizinkan "customer dengan akun di help center".
  6. Pembuatan customer otomatis dari email channel dan widget mengikuti setting ini: bila invite only dan project restricted, pengirim baru ditolak.
  7. ⚠ Setting siapa yang boleh membuat dan mengelola organization: hanya Admin Situs, atau juga Admin Proyek/agent.
  8. ⚠ Signup dapat dibatasi ke daftar email domain tertentu; signup dari domain lain ditolak dengan pesan "Domain email Anda tidak diizinkan".
  9. Hanya Admin Situs yang dapat mengubah pengaturan ini; setiap perubahan tercatat di audit log.
  10. Opsi akses project (restricted / customer berakun / anonim) dikonfigurasi per project di JSM-PRT-016.
  11. Public signup global mengaktifkan alur JSM-PRT-013.
  12. ⚠ Akses anonim memakai channel `anonymousportal` pada request yang dibuat tanpa login (JSM-PRT-017).
  13. CAPTCHA juga dipakai untuk halaman publik lain yang menerima input (lihat butir 3).
  14. Dua tipe akun customer (portal-only vs akun pengguna situs) didefinisikan di JSM-PRT-014.
  15. ⚠ Login sosial (Google/Microsoft) tersedia untuk akun pengguna situs lewat alur AUTH-003; untuk akun portal-only tidak tersedia pada versi awal.
  16. SSO/SAML untuk customer portal-only dikonfigurasi lewat authentication policy (JSM-PRT-018).
  17. Verifikasi email dipakai saat signup dan reset password (JSM-PRT-013).
  18. Profil customer (nama tampilan, bahasa, zona waktu, ubah password) diatur di JSM-PRT-019.
  19. Halaman Customer accounts di pengaturan produk (lihat, ubah nama/email, nonaktifkan, hapus akun portal-only) dirinci di JSM-PRT-020.
  20. ⚠ Migrasi akun portal-only ke akun pengguna situs dirinci di JSM-PRT-020.
- **Team-managed vs company-managed:** Tidak ada perbedaan; pengaturan bersifat situs/help center.
- **Catatan & ketergantungan:** Self-signup pengguna internal (USR-012, [02](02-pengguna-grup-organisasi.md)) terpisah dan tidak memengaruhi signup customer. Pada multiple help centers, pengaturan dapat ditimpa per help center (JSM-PRT-008).

### JSM-PRT-016 Customer permissions (project level)

- **Alias inventaris:** Customer permissions & sharing settings
- **Deskripsi:** Pengaturan per service project tentang siapa yang bisa melihat portal, raise request, berbagi request, melihat customer lain, dan mengakses knowledge base.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer projects), Admin Situs; Customer/Pelapor terdampak.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi akses "Only customers and organizations added to the project" (restricted): hanya customer/organization yang ditambahkan yang dapat melihat portal dan raise request.
  2. Opsi akses "Any customer who has an account on the help center": semua akun customer situs dapat melihat portal dan raise request.
  3. Opsi akses "Anyone on the web" (anonim) hanya dapat dipilih bila "Allow access without login" aktif di level situs (JSM-PRT-015); bila tidak, opsi tampil nonaktif dengan penjelasan.
  4. Pengaturan sharing: customer dapat berbagi request dengan "no one", "anggota organization yang sama", atau "semua customer/organization di project".
  5. ⚠ Toggle "Agents can add customers and organizations" menentukan apakah agent (bukan hanya admin) boleh menambah customer/organization ke project.
  6. Email dari pengirim non-customer ditolak bila project restricted (JSM-PRT-041) dan tercatat di email log.
  7. Portal tidak terlihat di help center dan pencarian bagi pengguna yang tidak berhak; akses URL langsung mengembalikan halaman "Anda tidak memiliki akses".
  8. Beralih ke opsi publik menampilkan dialog peringatan bahwa siapa pun dapat membuat request dan memerlukan konfirmasi.
  9. Opsi pencarian/autocomplete customer lain dan organization saat sharing mengikuti setting butir 4 sehingga daftar customer tidak bocor ke customer yang tidak berhak.
  10. Approver yang bukan customer project tetap dapat melihat dan memutuskan approval lewat portal atau email (JSM-PRT-031).
  11. Akses customer ke knowledge base dapat diatur: semua orang (termasuk anonim) atau hanya customer yang login; pengaturan disinkronkan dengan izin space KB di modul 31.
  12. Perubahan customer permissions berlaku segera untuk sesi customer yang sedang aktif pada request berikutnya.
  13. Konfigurasi customer permissions dan sharing dari project team-managed hasil impor dipetakan ke opsi company-managed yang setara (akses, sharing, penambahan customer oleh agent) tanpa mengubah perilaku efektif; pemetaan impor dirinci di [25](25-integrasi-impor-ekspor-ekstensi.md).
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Implementasi sebagai lapisan otorisasi customer di atas permission scheme [03](03-izin-peran-keamanan.md); peran "Customers" pada permission scheme company-managed. Sharing rinci di JSM-PRT-030. Pembagian kepemilikan: modul ini memiliki perilaku customer-facing (efek pada portal, sharing, email, KB); halaman pengaturan Project settings > Customer permissions dimiliki JSM-PRJ-051 ([29](29-jsm-service-project-request-type-formulir.md)); pengaturan dan default tingkat admin global dimiliki JSM-AI-002 Customer Permissions (Space) ([35](35-jsm-ai-otomasi-api-admin-csm.md)). Ketiganya memakai satu model data pengaturan yang sama.

### JSM-PRT-017 Anonymous requests

- **Deskripsi:** Pengunjung tanpa login dapat membuat request di portal yang dibuka untuk akses anonim.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Tamu/Kolaborator Eksternal (pengunjung anonim), Admin Proyek (mengaktifkan).
- **Aturan perilaku (acceptance criteria):**
  1. Fitur aktif hanya bila situs mengizinkan akses tanpa login (JSM-PRT-015) dan project memilih opsi "Anyone" (JSM-PRT-016).
  2. ⚠ Submit request anonim wajib melewati CAPTCHA; kegagalan CAPTCHA menolak submit tanpa menghapus isi form.
  3. ⚠ Form anonim meminta alamat email agar pengirim dapat menerima update; request dicatat dengan reporter berupa customer yang dicocokkan/dibuat dari email tersebut, atau reporter "Anonymous" bila email tidak diisi.
  4. Pengunjung anonim tidak dapat membuka menu Requests atau halaman detail request; setelah submit ia melihat pesan konfirmasi tanpa tautan detail.
  5. Field "Share with" tidak ditampilkan pada form anonim.
  6. Pengunjung anonim dapat membaca artikel KB yang dibuka untuk publik.
  7. Risiko spam dimitigasi dengan CAPTCHA serta rate limit per IP (mis. maksimal 10 request per jam), dan request anonim ditandai channel `anonymousportal`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rate limiting mengikuti API-006 ([23](23-api-publik-webhook.md)) dan CAPTCHA AUTH-006 ([01](01-identitas-autentikasi-sesi.md)).

### JSM-PRT-018 SSO & authentication policies for customers

- **Deskripsi:** Single sign-on dan authentication policy untuk akun customer menggunakan identity provider organisasi.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (organisasi), Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat mengonfigurasi SAML SSO khusus customer portal-only di konsol administrasi (entity ID, SSO URL, sertifikat IdP).
  2. Identity provider standar SAML 2.0 (mis. Okta, Microsoft Entra ID, Google Workspace) dapat dihubungkan ke customer accounts.
  3. Admin memilih help center mana yang memakai SSO customer; help center lain tetap memakai password.
  4. Login page menampilkan tombol "Log in with SSO" untuk help center yang mengaktifkan SSO customer.
  5. ⚠ Akun customer dibuat otomatis saat login SSO pertama (just-in-time provisioning) dengan email dan nama dari assertion SAML.
  6. Akun pengguna situs yang menjadi customer mengikuti authentication policy organisasi (SSO, two-step verification; AUTH-010, AUTH-011).
  7. ⚠ Session timeout customer dapat diatur per policy (default 30 hari idle); sesi kedaluwarsa meminta login ulang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada SSO enterprise modul 01/02 (Fase 4); SSO customer dapat dirilis di akhir Fase 5 setelah fondasi SAML tersedia. Prioritas diselaraskan dengan JSM-AI-009 Customer Authentication ([35](35-jsm-ai-otomasi-api-admin-csm.md)), yang memiliki metode login customer di konsol admin global (password, SSO SAML, login sosial); fitur ini memiliki perilaku customer-facing di help center (tombol login SSO, JIT provisioning, session timeout). Langganan keamanan identitas pihak lain (mis. Atlassian Guard) tidak dibangun; kemampuan setara disediakan native.

### JSM-PRT-019 Customer profile

- **Deskripsi:** Halaman profil customer di help center untuk mengatur identitas, bahasa, password, dan preferensi.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (profilnya sendiri).
- **Aturan perilaku (acceptance criteria):**
  1. Customer portal-only dapat mengubah nama tampilan (1-255 karakter).
  2. Email ditampilkan read-only; perubahan email hanya oleh admin (JSM-PRT-020).
  3. Customer memilih bahasa help center dari bahasa yang didukung; pilihan langsung diterapkan.
  4. ⚠ Customer memilih zona waktu yang dipakai untuk menampilkan tanggal di portal dan email; default dari browser.
  5. Customer portal-only dapat mengubah password dengan memasukkan password lama dan baru.
  6. ⚠ Preferensi global "Email notifications" (on/off) untuk semua request; mematikan global tidak menghapus langganan per request.
  7. Akun pengguna situs melihat tautan "Manage account" ke halaman profil akun produk (AUTH-002) alih-alih form nama/password.
  8. Avatar tampil untuk akun pengguna situs; akun portal-only memakai avatar inisial berwarna.
  9. Tombol Log out mengakhiri sesi help center dan kembali ke login page atau beranda publik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Profil akun pengguna di AUTH-002 ([01](01-identitas-autentikasi-sesi.md)).

### JSM-PRT-020 Customer accounts administration (admin.atlassian.com)

- **Deskripsi:** Pengelolaan akun customer portal-only di konsol administrasi produk ini (padanan admin.atlassian.com), termasuk migrasi ke akun pengguna situs.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (organisasi/situs).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Customer accounts menampilkan daftar akun portal-only per situs dengan pencarian nama/email dan paginasi.
  2. Detail akun menampilkan nama, email, tanggal dibuat, dan last active.
  3. Admin dapat mengubah nama dan email customer; email baru wajib unik di situs dan memicu email verifikasi ke alamat baru.
  4. Admin dapat mengirim email reset password ke customer.
  5. Admin dapat menonaktifkan dan mengaktifkan kembali akun; akun nonaktif tidak bisa login dan tidak menerima notifikasi.
  6. Delete account bersifat permanen setelah konfirmasi ketik email; request yang pernah dibuat tetap ada dengan reporter ditampilkan sebagai "Former customer".
  7. Admin dapat memigrasi akun portal-only ke akun pengguna situs.
  8. Saat migrasi, customer menerima email untuk menyelesaikan proses (set password/login akun baru).
  9. Riwayat request, keanggotaan participant, dan organization dipertahankan setelah migrasi.
  10. ⚠ Migrasi tidak dapat dibatalkan; dialog konfirmasi menyatakan hal ini.
  11. Migrasi ditolak bila email sudah dipakai akun pengguna situs lain.
  12. Daftar dapat difilter berdasarkan status (active, deactivated, unverified).
  13. ⚠ Impor customer massal native hanya lewat REST API (JSM-PRT-050); tidak ada impor CSV di halaman ini pada versi awal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konsol administrasi pengguna di [02](02-pengguna-grup-organisasi.md); semua aksi tercatat di audit log (ADM-014). Tumpang tindih dengan JSM-AI-008 Customer Accounts Management ([35](35-jsm-ai-otomasi-api-admin-csm.md)): fitur ini adalah versi rinci perilaku akun customer (daftar, edit, nonaktif, hapus, migrasi), sedangkan JSM-AI-008 memiliki penempatan di konsol admin global; keduanya harus memakai entitas dan endpoint yang sama. Penghapusan akun harus selaras dengan kebijakan penghapusan data pribadi (AUTH-014).

## Portal service project & pengajuan request

### JSM-PRT-021 Customer portal (per service project)

- **Alias inventaris:** Customer portal
- **Deskripsi:** Portal milik setiap service project yang menampilkan request type, help resources, KB, dan announcement project tersebut, tempat customer memilih request type, mengirim request, dan melacak statusnya.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Portal settings), Customer/Pelapor (sesuai customer permissions), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. URL portal adalah `/servicedesk/customer/portal/{portalId}` dengan `portalId` numerik stabil yang tidak berubah walau nama/key project berubah.
  2. Portal settings memuat nama portal (wajib, maksimal 100 karakter), introduction text (rich text, opsional), dan logo portal (format dan batas sama dengan JSM-PRT-002).
  3. Portal announcement tampil di bagian atas portal (JSM-PRT-005).
  4. Halaman portal menampilkan request type (ikon, nama, deskripsi) dan help resources yang dikelompokkan per portal group (JSM-PRT-022).
  5. Search bar portal "What can we help you with?" mencari request type dan artikel KB portal itu (JSM-PRT-006).
  6. Portal menyediakan tautan kembali ke help center lewat breadcrumb dan logo.
  7. Portal hanya terlihat oleh pengguna yang diizinkan customer permissions project (JSM-PRT-016).
  8. Customer portal-only yang mencoba membuka URL tampilan agen project mendapat 403 dan diarahkan ke portal.
  9. Agent dapat membuka "View customer portal" dari sidebar project, yang membuka portal di tab baru sebagai dirinya.
  10. Setiap service project memiliki tepat satu portal yang dibuat otomatis bersama project.
  11. ⚠ Admin Proyek dapat menonaktifkan portal sebagai channel (JSM-PRT-048): portal tidak tampil di help center dan URL-nya menampilkan pesan "Portal tidak tersedia", sementara email dan kanal lain tetap berjalan.
  12. Halaman create request memuat form, attachment drag & drop, dan saran artikel KB (rinci di JSM-PRT-024 dan JSM-PRT-025).
  13. Agent dapat memakai "Raise this request on behalf of" dengan pencarian customer (JSM-PRT-024).
  14. Field "Share with" organization/customer lain saat create mengikuti customer permissions (JSM-PRT-030).
  15. Setelah submit, portal menampilkan konfirmasi beserta tautan ke request.
  16. Detail request menampilkan summary, field yang diisi, status customer-facing, activity (komentar publik saja), dan attachment (JSM-PRT-026).
  17. Customer dapat menambah komentar, attachment, dan participant dari detail request.
  18. Panel approval untuk approver menampilkan tombol Approve/Decline (JSM-PRT-031).
  19. Toggle "Don't notify me"/"Notify me" tersedia per request (JSM-PRT-036).
  20. Portal tidak pernah menampilkan komentar internal, SLA, atau field yang tidak dipetakan ke request type.
- **Team-managed vs company-managed:** Tidak ada perbedaan; Portal settings berada di Project settings > Channels & self-service > Portal untuk keduanya.
- **Catatan & ketergantungan:** Service project dan request type di modul 29. Penyajian customer harus memakai DTO khusus customer sehingga field internal tidak pernah dikirim ke klien portal.

### JSM-PRT-022 Portal groups

- **Deskripsi:** Pengelompokan request type dan help resources di portal agar customer mudah menemukan formulir atau sumber yang tepat.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengelola), Customer/Pelapor (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Portal group dapat dibuat, diubah namanya (wajib, maksimal 100 karakter, unik per portal), dan dihapus; menghapus grup tidak menghapus request type di dalamnya.
  2. Urutan grup dan urutan item di dalam grup diatur dengan drag-and-drop dan disimpan per grup.
  3. Satu request type dapat berada di banyak grup.
  4. Request type tanpa grup tidak tampil di portal, tetapi tetap dapat dipakai lewat email, API, atau agent.
  5. Grup tanpa item yang terlihat oleh viewer tidak ditampilkan.
  6. Bila portal hanya memiliki satu grup yang terlihat, portal menampilkan daftar item tanpa navigasi grup.
  7. Nama grup dapat diterjemahkan (JSM-PRT-007).
  8. Help resources dapat ditambahkan sebagai item berupa link ke resource eksternal dengan URL (wajib, http/https), judul (wajib), deskripsi, dan ikon.
  9. Help resources tampil di portal group bersama request type, dapat diurutkan, dan dipindah antar grup.
  10. Mengklik help resource membuka URL eksternal di tab baru dan tidak membuat request.
  11. Help resource dapat diedit dan dihapus.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Titik integrasi dengan request type di modul 29. Penyimpanan urutan memakai kolom `position` per pasangan grup-item. Fitur ini adalah pemilik heading inventaris "Portal groups" (model data, aturan tampil di portal, help resources); JSM-PRJ-024 Request type portal groups ([29](29-jsm-service-project-request-type-formulir.md)) hanya halaman pengaturan Project settings > Request types. Keputusan: JSM-PRJ-024 sebaiknya melepas baris alias "Portal groups" agar inventaris tidak tercakup dua kali (perubahan di modul 29).

### JSM-PRT-023 Request type visibility & restrictions on portal

- **Deskripsi:** Membatasi siapa yang dapat melihat dan memakai request type tertentu di portal.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengatur), Customer/Pelapor (terdampak).
- **Aturan perilaku (acceptance criteria):**
  1. Request type tanpa portal group tersembunyi dari portal tetapi tetap dapat dipakai lewat email atau agent.
  2. ⚠ Request type dapat dibatasi ke pengguna, group, atau organization tertentu; customer di luar daftar tidak melihatnya.
  3. Request type yang dibatasi tidak muncul di search (JSM-PRT-006) maupun topics (JSM-PRT-004) bagi viewer yang tidak berhak.
  4. Akses langsung lewat URL form request type yang tidak berhak mengembalikan halaman "Anda tidak memiliki akses ke request type ini" dan API mengembalikan 403.
  5. Agent tetap dapat membuat request dengan request type terbatas atas nama customer mana pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi utama request type di modul 29.

### JSM-PRT-024 Raise a request (portal request form)

- **Deskripsi:** Pengalaman customer mengisi dan mengirim request melalui request form di portal.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, pengunjung anonim (bila diizinkan), Agent Layanan (raise on behalf).
- **Aturan perilaku (acceptance criteria):**
  1. Customer memilih request type dari portal, search, topik, atau deep link, lalu form request type tersebut terbuka.
  2. Form memuat summary, description, custom field yang dipetakan ke request type, dan forms bila terlampir (modul 29).
  3. Setiap field menampilkan help text dan penanda required (`*`).
  4. Validasi field required dan format (email, angka, tanggal, URL) dilakukan sebelum submit di klien dan diulang di server; galat ditampilkan per field.
  5. File dapat dilampirkan dengan drag-and-drop atau pemilih file, mengikuti batas ukuran dan tipe dari attachment settings (ADM-011); file di atas batas ditolak dengan pesan galat.
  6. Agent melihat field "Raise this request on behalf of" dengan pencarian customer; bila diizinkan customer permissions, agent dapat mengetik email baru untuk membuat customer sekaligus.
  7. Field "Share with" menawarkan opsi private atau organization milik customer sesuai customer permissions (JSM-PRT-030).
  8. Saran artikel KB muncul saat customer mengetik summary (JSM-PRT-025).
  9. Hidden fields dengan nilai preset tetap terisi pada request yang dibuat walau tidak tampil di form.
  10. Setelah submit, portal menampilkan pesan konfirmasi lalu mengarahkan ke halaman detail request (JSM-PRT-026).
  11. Deep link `/servicedesk/customer/portal/{portalId}/create/{requestTypeId}` membuka form request type tertentu secara langsung.
  12. Field objek Assets dapat ditampilkan di form bila modul Assets aktif (modul 33).
  13. ⚠ Request anonim wajib lolos CAPTCHA sebelum submit (JSM-PRT-017).
  14. Request yang dibuat dari portal ditandai channel `portal` dan memicu event "Request created" untuk customer notification.
  15. Isi form yang belum dikirim disimpan sementara di browser sehingga tidak hilang saat customer membuka artikel KB atau me-refresh halaman.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Field mapping dan forms di modul 29; attachment di [08](08-komentar-lampiran-aktivitas.md) CMT-011. Pembuatan work item memakai use case create work item ([06](06-work-item-inti.md)).

### JSM-PRT-025 KB article suggestions on portal (deflection)

- **Deskripsi:** Saran artikel knowledge base saat customer mengetik request agar masalah selesai tanpa membuat tiket.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, pengunjung anonim.
- **Aturan perilaku (acceptance criteria):**
  1. Saran muncul berdasarkan teks summary yang diketik: minimal 3 karakter, debounce 300 ms, maksimal 5 artikel (JSM-QSK-035 butir 2 dan 13).
  2. Saran dibatasi ke space KB yang ditautkan ke project.
  3. Artikel dibuka di panel samping atau halaman portal tanpa kehilangan isi form.
  4. Artikel menampilkan tombol "Was this helpful?" (Yes/No).
  5. Bila customer memilih Yes, ia ditawari membatalkan pembuatan request; pembatalan tercatat sebagai deflection.
  6. Hanya artikel yang boleh dilihat viewer yang muncul sebagai saran.
  7. Halaman artikel di portal merender isi artikel KB beserta daftar artikel terkait.
  8. Data deflection (artikel, request type, hasil helpful) dikirim sebagai event untuk laporan KB di modul 31.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Logika saran dan pencatatan deflection dimiliki JSM-QSK-035 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)); fitur ini hanya sisi render portal. Konfigurasi KB dan laporan deflection di modul 31.

## Pelacakan request oleh customer

### JSM-PRT-026 Request details page (portal)

- **Deskripsi:** Halaman yang menampilkan status, detail, dan aktivitas request dari sudut pandang customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (reporter, participant, anggota organization yang dibagikan), approver, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Header menampilkan key, summary, request type (ikon + nama), dan customer-visible status (JSM-PRT-027).
  2. Hanya field yang dipetakan ke request type yang ditampilkan, dengan nama field versi customer.
  3. Activity menampilkan komentar publik, perubahan status customer-visible, dan lampiran secara kronologis.
  4. Komentar internal agent tidak pernah dikirim ke klien portal (diverifikasi di level API, bukan hanya UI).
  5. Customer dapat menambah komentar dengan rich text editor dan melampirkan file; komentar customer selalu publik.
  6. Panel "Shared with" menampilkan reporter, participants, dan organizations yang dibagikan.
  7. Toggle notifikasi per request ("Turn on/off notifications") tersedia (JSM-PRT-036).
  8. Tombol customer transitions tampil sesuai konfigurasi dan status saat ini (JSM-PRT-028).
  9. Panel approval tampil bila request menunggu persetujuan (JSM-PRT-031).
  10. SLA, linked work items, assignee internal, dan field internal tidak ditampilkan.
  11. ⚠ Customer tidak dapat mengedit field setelah submit; jawaban forms dapat diedit hanya bila form diatur dapat diedit (modul 29).
  12. Akses ditolak (halaman "Anda tidak memiliki akses", API 403/404) bila pengguna bukan reporter, participant, anggota organization yang dibagikan, atau approver.
  13. ⚠ Agent dengan akses project melihat tautan "Buka di tampilan agen" ke halaman work item.
  14. Halaman detail menampilkan rating CSAT bila request sudah resolved dan feedback aktif (JSM-PRT-032).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komentar publik/internal di CMT-001 dan CMT-004 ([08](08-komentar-lampiran-aktivitas.md)); editor di [07](07-editor-konten-kaya.md).

### JSM-PRT-027 Customer-visible status names

- **Deskripsi:** Pemetaan nama status workflow internal ke nama status yang ramah customer di portal dan email.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Nama tampilan per status diatur di tab Workflow statuses pada konfigurasi request type (maksimal 60 karakter).
  2. Satu status internal dapat memiliki nama tampilan berbeda per request type.
  3. Nama tampilan dipakai di portal, filter status pada daftar Requests, dan variabel notifikasi `${request.status}`.
  4. Event "Customer-visible status changed" hanya terpicu bila nama tampilan sebelum dan sesudah transition berbeda.
  5. Nama tampilan dapat diterjemahkan (JSM-PRT-007).
  6. Default nama tampilan sama dengan nama status workflow; mengosongkan field mengembalikan ke default.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status dan workflow di [09](09-workflow-status-transisi.md); request type di modul 29.

### JSM-PRT-028 Customer transitions

- **Deskripsi:** Workflow transition yang boleh dijalankan customer langsung dari portal, misalnya membatalkan atau membuka kembali request.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengonfigurasi), Customer/Pelapor (reporter dan participant yang berhak).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin memilih transition yang tersedia bagi customer per request type di tab Workflow statuses request type.
  2. Admin menetapkan nama transition yang tampil ke customer (default nama transition workflow).
  3. Menjalankan transition membuka dialog konfirmasi dengan komentar opsional yang disimpan sebagai komentar publik.
  4. Hanya transition yang berasal dari status saat ini yang tampil.
  5. ⚠ Transition yang memerlukan screen dengan field wajib atau validator yang tidak dapat dipenuhi customer tidak dapat dipilih sebagai customer transition; UI konfigurasi menampilkan alasannya.
  6. Transition yang mengubah resolution memicu event notifikasi Request resolved/reopened.
  7. Contoh transition yang disediakan template: "Cancel request" dan "Reopen".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mesin transition dan validator di [09](09-workflow-status-transisi.md). Eksekusi transition oleh customer tetap melewati post function dan otomasi.

### JSM-PRT-029 Requests list (My requests / All requests)

- **Alias inventaris:** Requests list (portal)
- **Deskripsi:** Halaman "Requests" di help center yang menampilkan request customer lintas portal, dengan filter, pencarian, dan halaman Approvals.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor, approver, Agent Layanan (sebagai customer).
- **Aturan perilaku (acceptance criteria):**
  1. Menu "Requests" menawarkan filter kepemilikan: Created by me, Where I'm a participant, Shared with organization <nama> (satu per organization), dan All.
  2. Filter status: Open requests, Closed requests, Any status; status memakai nama customer-visible.
  3. Filter request type dan service project/portal tersedia sebagai dropdown.
  4. Pencarian teks mencocokkan summary atau key.
  5. Kolom yang tampil: Type, Reference (key), Summary, Status, Service project, Requester.
  6. Kolom dapat diurutkan dan daftar dipaginasi (default 50 per halaman).
  7. Request yang dibagikan ke organization customer ikut tampil pada filter All dan filter organization tersebut.
  8. Daftar menggabungkan request dari semua portal, termasuk request yang dibuat pengguna berlisensi/agent sebagai dirinya sendiri.
  9. Menu Requests tidak tersedia untuk pengunjung anonim; URL langsung mengarahkan ke login page.
  10. Halaman Approvals menampilkan request yang menunggu keputusan pengguna sebagai approver (JSM-PRT-031).
  11. Pengguna hanya dapat melihat request miliknya, tempat ia menjadi participant, atau yang dibagikan ke organization-nya; query dibatasi di server.
  12. Filter yang dipilih tercermin di URL sehingga dapat dibagikan dan di-bookmark.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Query daftar memakai indeks pada reporter, participant, dan organization; padanan API `GET /rest/servicedeskapi/request` (JSM-PRT-050).

### JSM-PRT-030 Request participants & sharing

- **Deskripsi:** Berbagi request dengan customer lain atau organization agar mereka dapat melihat, berkomentar, dan menerima notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (sesuai customer permissions), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Field sistem "Request participants" bertipe multi-user tersedia di setiap service project.
  2. Customer menambah participant dari portal lewat tombol "Share" sesuai pengaturan sharing (JSM-PRT-016).
  3. Agent dapat menambah participant dari tampilan agen tanpa dibatasi pengaturan sharing customer.
  4. Share ke organization membuat semua anggota organization dapat melihat request.
  5. Saat create, "Share with" menampilkan organization milik customer; customer yang tergabung di banyak organization memilih salah satu.
  6. Customer tanpa organization hanya melihat opsi private.
  7. ⚠ Nilai default "Share with" mengikuti setting project (private atau organization pertama).
  8. Penerima CC pada email masuk otomatis menjadi participant (JSM-PRT-041).
  9. Participant menerima customer notification dan dapat berkomentar.
  10. Participant dapat dihapus dan organization dapat di-unshare hanya oleh agent atau reporter (selaras dengan JSM-PRJ-054 butir 18); participant lain, termasuk yang menambahkannya, tidak dapat menghapus participant, tetapi setiap participant dapat menghapus dirinya sendiri dari request.
  11. Menambah participant memicu event "Participant added"; menambah organization memicu "Organization added".
  12. Cakupan pencarian customer saat share mengikuti setting: no one, anggota organization yang sama, atau semua customer project; hasil pencarian tidak pernah menampilkan customer di luar cakupan.
  13. ⚠ Jumlah participant per request mengikuti batas JSM-PRJ-054 butir 21 (default 100); share dari portal yang melebihi batas ditolak dengan pesan galat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemilik entitas Request participants (field, batas, aturan hapus) adalah JSM-PRJ-054 Request details and request participants ([29](29-jsm-service-project-request-type-formulir.md)), yang memegang alias inventaris "Request participants"; fitur ini hanya memiliki perilaku sharing sisi portal. Pengelolaan participant dan organization dari issue view agent di JSM-QSK-013 ([31](31-jsm-antrian-sla-knowledge-base-laporan.md)). Event notifikasi di JSM-PRT-036; JQL `"Request participants"` dan `organizations` di [12](12-jql-bahasa-kueri.md).

### JSM-PRT-031 Approvals on portal & email

- **Deskripsi:** Pengalaman approver menyetujui atau menolak request dari help center atau langsung dari email.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Stakeholder/Viewer sebagai approver (pengguna berlisensi atau customer), Customer/Pelapor (melihat status approval).
- **Aturan perilaku (acceptance criteria):**
  1. Menu "Approvals" di help center berisi daftar approval yang menunggu keputusan pengguna, dengan jumlah di badge.
  2. Panel detail approval menampilkan daftar approver, status per approver (pending/approved/declined), dan keputusan sebelumnya.
  3. Tombol Approve/Decline tampil di halaman request bagi approver yang belum memutuskan.
  4. Email "Approval required" memuat tombol Approve/Decline yang membuka halaman konfirmasi (bukan langsung memutuskan lewat GET) sehingga pemindai tautan email tidak memicu keputusan.
  5. ⚠ Approver dapat menambahkan komentar opsional saat memutuskan; komentar disimpan sebagai komentar publik.
  6. Approver mendapat akses melihat request walaupun bukan reporter atau participant.
  7. ⚠ Approver dapat berupa pengguna berlisensi atau customer portal-only.
  8. Riwayat approval tetap terlihat di halaman request setelah approval selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konfigurasi approval steps di modul 32; approval bisnis non-JSM di [26](26-kerja-bisnis-formulir-persetujuan.md).

### JSM-PRT-032 Request feedback (CSAT)

- **Deskripsi:** Survei kepuasan customer yang dikirim setelah request diselesaikan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengaktifkan), Customer/Pelapor (memberi rating).
- **Aturan perilaku (acceptance criteria):**
  1. Toggle "Request feedback" tersedia per project (default nonaktif untuk project baru dari template kosong).
  2. Survei dikirim saat request mendapat resolution (bukan sekadar berpindah status).
  3. Rating berupa 1-5 bintang dengan komentar opsional maksimal 1.000 karakter.
  4. Rating dapat diberikan dari email (klik bintang membuka halaman portal yang sudah terisi) atau dari halaman detail request.
  5. ⚠ Survei dikirim ke reporter saja; participant tidak menerima survei.
  6. Field "Satisfaction" dan "Satisfaction date" dapat dipakai di JQL dan queue.
  7. ⚠ Hanya satu rating per request; rating dapat diubah selama request belum di-reopen, dan reopen lalu resolve ulang memicu survei baru.
  8. Hasil rating masuk ke laporan customer satisfaction (modul 31).
  9. ⚠ Admin dapat mengubah teks pengantar survei.
  10. Email survei diterjemahkan sesuai bahasa customer.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Laporan CSAT dan agregasi di modul 31; JQL `Satisfaction` di [12](12-jql-bahasa-kueri.md).

## Customers & organizations

### JSM-PRT-033 Customers page (project)

- **Alias inventaris:** Customers management
- **Deskripsi:** Halaman Customers di service project untuk menambah, mengundang, dan mengelola customer serta organization yang punya akses ke portal.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Agent Layanan (bila "Agents can add customers and organizations" aktif).
- **Aturan perilaku (acceptance criteria):**
  1. Daftar menampilkan customer dan organization project dengan kolom nama, email, dan organization.
  2. Pencarian berdasarkan nama atau email; daftar dapat difilter per organization.
  3. "Add customers" menerima beberapa email sekaligus (dipisah koma atau baris baru, maksimal 100 per aksi), membuat akun portal-only bila belum ada, lalu mengirim undangan.
  4. "Add organization" menambah organization yang sudah ada atau membuat organization baru (sesuai setting JSM-PRT-015).
  5. Aksi per customer: add to organization, remove from organization, dan remove from project.
  6. ⚠ Nama customer portal-only dapat diedit dari halaman ini; perubahan email hanya lewat administrasi akun (JSM-PRT-020).
  7. Detail customer menampilkan email, organization, dan daftar request terkait di project.
  8. ⚠ Customer yang undangannya belum diterima atau akunnya belum terverifikasi ditandai label "Invited"/"Unverified".
  9. Daftar dipaginasi (default 50 per halaman) untuk project dengan ribuan customer.
  10. Menghapus customer dari project tidak menghapus akun di situs maupun request lamanya.
  11. Akses halaman: Admin Proyek, atau agent bila diizinkan; pengguna lain mendapat 403.
  12. Aksi "Resend invite" tersedia untuk customer berstatus Invited.
  13. ⚠ Impor customer dari file CSV (kolom email, nama, organization) dengan laporan baris yang gagal.
  14. Customer otomatis dibuat saat mengirim email ke project yang terbuka (JSM-PRT-041) dan muncul di daftar ini.
  15. Customer tidak memakai lisensi dan jumlahnya tidak dibatasi (JSM-PRT-014).
  16. ⚠ Detail customer dapat menampilkan customer detail fields (tier, region, account manager) bila JSM-PRT-035 aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Akun dan tipe akun di JSM-PRT-014; undangan di JSM-PRT-013. Butir 13 (impor CSV) bersifat kandidat dan bertentangan dengan default JSM-PRT-020 butir 13 (impor massal hanya lewat REST API); keputusan akhir di Pertanyaan terbuka nomor 11, dan bila CSV tidak dipilih butir 13 dihapus.

### JSM-PRT-034 Organizations

- **Deskripsi:** Kelompok customer (misalnya per perusahaan klien) yang ditambahkan ke project agar request bisa dibagikan, dicari, dan dilaporkan per organization.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (halaman site-wide), Admin Proyek dan Agent Layanan (sesuai setting), Customer/Pelapor (anggota).
- **Aturan perilaku (acceptance criteria):**
  1. Organization dibuat dengan nama wajib yang unik di situs (case-insensitive); nama duplikat ditolak.
  2. Organization dapat diubah namanya dan dihapus; penghapusan melepas organization dari semua request dan project setelah konfirmasi.
  3. Anggota dapat ditambah (berdasarkan email; customer dibuat bila belum ada) dan dihapus.
  4. Satu customer dapat menjadi anggota banyak organization.
  5. Organization bersifat site-wide dan dapat ditambahkan ke banyak service project.
  6. Halaman Organizations site-wide untuk Admin Situs menampilkan semua organization, jumlah anggota, dan project tertaut.
  7. Field "Organizations" (multi-value) tersedia pada request dan dapat dicari lewat JQL `organizations = "Nama"`.
  8. Request yang dibagikan ke organization terlihat oleh semua anggotanya.
  9. Menambah organization ke request memicu event notifikasi "Organization added".
  10. ⚠ Setting siapa yang boleh membuat/mengelola organization (hanya admin atau juga agent) mengikuti JSM-PRT-015.
  11. Organization properties (key-value JSON) dapat dibaca dan ditulis lewat REST API (JSM-PRT-050).
  12. Queue dan filter dapat memakai kondisi organization (modul 31).
  13. ⚠ Saat request dibuat oleh customer yang tergabung di satu organization project, field Organizations otomatis terisi organization tersebut bila "Share with" memilih organization; customer multi-organization memilih salah satu.
  14. Laporan dapat dikelompokkan per organization (modul 31).
  15. REST API `/rest/servicedeskapi/organization` mendukung CRUD organization dan anggota.
  16. Organization detail fields (domain, plan, region) tersedia bila JSM-PRT-035 aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Istilah "organization" di modul ini berbeda dari organisasi administrasi situs di [02](02-pengguna-grup-organisasi.md); nama tabel disarankan `customer_organizations` untuk mencegah ambiguitas. JQL di [12](12-jql-bahasa-kueri.md).

### JSM-PRT-035 Customer and organization details (profiles)

- **Deskripsi:** Field detail kustom pada customer dan organization (misalnya tier atau account manager) yang dilihat agent saat menangani request.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (mendefinisikan field), Agent Layanan (mengisi dan melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Admin mendefinisikan field detail bertipe text, select (single/multi), number, date, URL, atau user.
  2. Agent mengisi nilai detail di profil customer atau organization.
  3. Detail tampil di panel request untuk agent (panel "Customer details" dan "Organization details").
  4. Detail tidak pernah terlihat oleh customer di portal atau email.
  5. Halaman profil customer/organization menampilkan detail beserta riwayat request.
  6. ⚠ Nilai detail dapat dipakai di JQL dan kondisi otomasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditempatkan di Fase 7 karena menjadi dasar Customer Service Management (modul 35). Otomasi di [20](20-otomasi.md).

## Customer notifications

### JSM-PRT-036 Customer notifications

- **Alias inventaris:** JSM Customer Notifications
- **Deskripsi:** Email otomatis ke customer, participant, organization, dan approver untuk event request, dengan template yang dapat dikustomisasi per project dan terpisah dari notification scheme internal.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengonfigurasi), Customer/Pelapor dan participant (penerima), Agent Layanan (terpengaruh lewat notification scheme).
- **Aturan perilaku (acceptance criteria):**
  1. Event bawaan: Request created, Public comment added, Customer-visible status changed, Request resolved, Request reopened, Participant added, Organization added, dan Approval required.
  2. ⚠ Event "Public comment edited" tersedia dan default nonaktif.
  3. Setiap event dapat diaktifkan atau dinonaktifkan per project.
  4. Subject dan isi template setiap event dapat diedit (rich text atau HTML terbatas yang disanitasi).
  5. Variabel dapat disisipkan lewat pemilih variabel (JSM-PRT-037); variabel tidak dikenal ditolak saat menyimpan.
  6. Template dapat di-preview dengan data contoh dan di-reset ke default.
  7. Penerima mengikuti event: reporter, participants, anggota organization yang dibagikan, dan approver (untuk Approval required); pemicu aksi tidak menerima email atas aksinya sendiri.
  8. Template tersedia per bahasa aktif dan dipilih sesuai bahasa profil penerima (fallback default).
  9. Komentar publik dikirim ke customer; komentar internal tidak pernah memicu customer notification.
  10. Customer/participant dapat mematikan atau menyalakan notifikasi satu request dari halaman request atau dari tautan di email.
  11. Status langganan disimpan per pengguna per request dan tidak mengubah preferensi global.
  12. Participant yang mematikan notifikasi tetap dapat melihat dan mengomentari request.
  13. ⚠ Customer yang mematikan notifikasi otomatis berlangganan kembali saat ia berkomentar di request tersebut.
  14. ⚠ Mode "safe customer notifications" per service project membuat email hanya berisi tautan ke portal tanpa isi komentar atau field sensitif.
  15. ⚠ Saat safe customer notifications aktif, attachment dikirim sebagai tautan (bukan lampiran) dan balasan email tetap diproses menjadi komentar.
  16. ⚠ Pengguna berlisensi (agent) yang menjadi reporter atau participant menerima notifikasi internal (notification scheme), bukan customer notification, untuk menghindari email ganda.
  17. Alternatif pengiriman email kustom tersedia lewat action "Send email" di otomasi ([20](20-otomasi.md)).
  18. ⚠ Log customer notification per request (event, penerima, waktu, status kirim) tersedia untuk troubleshooting.
  19. Email channel per project (alamat cloud atau email kustom) menjadi alamat pengirim dan balasan email menjadi komentar (JSM-PRT-040, JSM-PRT-041).
  20. Allowlist/blocklist pengirim email request dikelola di JSM-PRT-042.
  21. ⚠ Daftar notifikasi default di atas menjadi baseline; event tambahan (mis. Public comment edited) ditandai sebagai opsional.
  22. Custom notification rules dapat dibuat dengan trigger event, kondisi JQL, penerima, dan konten.
  23. Email look & feel global (header, footer, logo, warna) dikonfigurasi di JSM-PRT-038.
  24. Sender name dan alamat From per project dikonfigurasi di JSM-PRT-039.
  25. Agent menerima notifikasi lewat notification scheme internal ([19](19-notifikasi-email.md) NTF-004), bukan lewat konfigurasi ini.
  26. ⚠ Sintaks variabel template memakai bentuk `${...}` (mis. `${request.summary}`, `${request.status}`); daftar dan nama final dirinci di JSM-PRT-037.
  27. ⚠ Konfigurasi customer notification dari project team-managed hasil impor (event bawaan aktif/nonaktif dan template) dipetakan ke konfigurasi company-managed yang setara; clone tidak menyediakan mode subset team-managed ([25](25-integrasi-impor-ekspor-ekstensi.md)).
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Infrastruktur pengiriman (antrean, retry, SMTP) di NTF-023/NTF-041 ([19](19-notifikasi-email.md)). Event domain dari katalog [28](28-model-data-inti.md). Modul ini memiliki perilaku customer notification (event, penerima, template, langganan); halaman pengaturan Project settings > Customer notifications dimiliki JSM-PRJ-067 ([29](29-jsm-service-project-request-type-formulir.md)). Email harus memiliki header `Message-ID`/`In-Reply-To` dan token key request untuk threading balasan (JSM-PRT-041).

### JSM-PRT-037 Notification template variables

- **Deskripsi:** Variabel dinamis yang dapat disisipkan di subject dan isi template customer notification.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Variabel `${request.key}` dan `${request.summary}` tersedia di semua event.
  2. Variabel `${request.status}` menghasilkan customer-visible status (JSM-PRT-027).
  3. Variabel `${request.url}` menghasilkan tautan portal request (memakai custom domain bila aktif).
  4. Variabel `${request.portal.name}` menghasilkan nama portal dalam bahasa penerima.
  5. Variabel `${request.description}` menghasilkan description yang dirender ke HTML.
  6. Variabel `${comment}` menghasilkan komentar publik terbaru (hanya pada event komentar).
  7. ⚠ Variabel nama penerima dan reporter (mis. `${recipient.name}`, `${request.reporter.name}`) tersedia.
  8. Event Approval required menyediakan variabel tautan approve/decline yang mengarah ke halaman konfirmasi.
  9. ⚠ Nama variabel final dipublikasikan di dokumentasi produk dan pemilih variabel; variabel tidak dikenal ditolak saat menyimpan dengan pesan galat yang menyebut nama variabel.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine template harus meng-escape nilai variabel (anti-injeksi HTML) kecuali variabel yang memang berupa HTML tersanitasi.

### JSM-PRT-038 Email look and feel (notification template wrapper)

- **Deskripsi:** Template pembungkus (header/footer) yang dipakai semua customer notification di project.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Header dan footer dapat diedit dengan editor HTML terbatas (tag aman dan inline CSS) yang disanitasi.
  2. Wrapper wajib memuat placeholder konten utama notifikasi; penyimpanan tanpa placeholder ditolak.
  3. Teks pemisah "Reply above this line" disisipkan otomatis di bagian atas email untuk parsing balasan dan tidak dapat dihapus.
  4. Footer memuat tautan untuk mematikan notifikasi request tersebut.
  5. ⚠ Logo dan warna default mengikuti branding help center (JSM-PRT-002).
  6. Wrapper dapat di-preview dan di-reset ke default.
  7. Versi plain text dibangkitkan otomatis dari HTML dan dikirim sebagai bagian `multipart/alternative`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template email internal di NTF-042 ([19](19-notifikasi-email.md)).

### JSM-PRT-039 Sender name, from address & reply-to

- **Deskripsi:** Pengaturan identitas pengirim email ke customer dan alamat tujuan balasan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (alamat dan nama), Admin Situs (domain pengirim).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Email dikirim dari alamat email default project, atau dari alamat tempat request masuk bila request berasal dari email.
  2. ⚠ Sender display name dapat dikustomisasi per project (default nama project).
  3. Header Reply-To mengarah ke alamat project sehingga balasan menjadi komentar.
  4. Custom email sending domain dengan verifikasi DNS (SPF, DKIM) dikonfigurasi di konsol administrasi (NTF-026).
  5. Tanpa domain terverifikasi, email dikirim lewat alamat sistem produk dengan nama tampilan project dan Reply-To ke alamat project.
  6. Bila project punya beberapa alamat email, admin memilih alamat default pengirim.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Alamat pengirim project di NTF-025 dan deliverability di NTF-026 ([19](19-notifikasi-email.md)).

## Email channel

### JSM-PRT-040 Email channel – connect email accounts

- **Alias inventaris:** Email channel
- **Deskripsi:** Menghubungkan alamat email ke service project agar email masuk menjadi request dan balasan email tersinkron sebagai komentar.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengelola alamat), Admin Situs (pengaturan global), Customer/Pelapor (pengirim).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap service project otomatis mendapat alamat cloud bawaan di domain produk (mis. `<project-key>@<situs>.<domain-produk>`) yang aktif tanpa konfigurasi.
  2. Admin dapat menghubungkan Gmail/Google Workspace lewat OAuth 2.0 (scope baca/tulis mailbox) tanpa menyimpan password.
  3. Admin dapat menghubungkan Microsoft 365/Outlook lewat OAuth 2.0.
  4. Akun Microsoft/Google lama yang masih memakai basic auth ditandai "perlu reauthorize" dan dapat dialihkan ke OAuth tanpa kehilangan pemetaan request type.
  5. Admin dapat menghubungkan mail server lain via IMAP atau POP3 dengan host, port, SSL/TLS, username, dan password (disimpan terenkripsi).
  6. Tombol "Test connection" memverifikasi kredensial dan akses mailbox sebelum disimpan, dengan pesan galat spesifik (autentikasi gagal, host tidak terjangkau, folder tidak ada).
  7. Admin dapat memakai forwarding dari domain sendiri ke alamat cloud; UI menampilkan instruksi dan alamat tujuan forwarding.
  8. Satu project dapat memiliki banyak alamat email, masing-masing dipetakan ke satu request type.
  9. Satu alamat email hanya dapat dihubungkan ke satu project; percobaan menghubungkan alamat yang sudah dipakai ditolak.
  10. Request type yang dipilih untuk email tidak boleh memiliki field required tanpa nilai default selain summary/description; pemilihan request type yang melanggar ditolak dengan daftar field penyebabnya.
  11. Satu alamat dapat ditandai sebagai default email address project.
  12. Setiap alamat dapat dinonaktifkan, diaktifkan kembali, dan dihapus.
  13. Toggle project "Email requests" mematikan pemrosesan semua alamat email project tanpa menghapus konfigurasi.
  14. Status koneksi tiap alamat ditampilkan: connected, error (dengan pesan terakhir), atau needs reauthorization.
  15. ⚠ Mailbox kustom di-poll sekitar setiap 1 menit; alamat cloud memproses email secara push.
  16. Summary request berasal dari subject dan description dari body email (rinci di JSM-PRT-041).
  17. Balasan email menjadi komentar publik pada request yang sama.
  18. Attachment dan gambar inline menjadi attachment request.
  19. Penerima CC dapat ditambahkan sebagai participant (opsional, default aktif).
  20. Pengirim baru dibuat sebagai customer bila project terbuka dan ditolak bila project restricted.
  21. Email log mencatat email diterima, diproses, dan gagal beserta alasannya (JSM-PRT-043).
  22. ⚠ Allowlist/blocklist pengirim dan pengabaian auto-reply dikelola di JSM-PRT-042.
  23. Pembersihan quoted reply dan signature dilakukan saat parsing (JSM-PRT-041).
  24. Admin Situs dapat mematikan email requests secara global untuk semua service project.
  25. Email ke project yang diarsipkan atau dinonaktifkan tidak diproses dan tercatat sebagai rejected di email log.
  26. ⚠ Jumlah alamat email per project dibatasi konfigurasi (default 10).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Infrastruktur IMAP/POP dan mail handler di NTF-023, NTF-031, NTF-032 ([19](19-notifikasi-email.md)); email masuk project bisnis (NTF-034) memakai pipeline yang sama tetapi aturan customer di modul ini. Alamat cloud memerlukan layanan inbound mail (mis. MX ke penerima SMTP internal); keputusan di Pertanyaan terbuka. Pengaturan global dan batas email request (lintas project) dimiliki JSM-AI-007 Email Requests Global Settings & Limits ([35](35-jsm-ai-otomasi-api-admin-csm.md)); fitur ini memiliki koneksi email per project.

### JSM-PRT-041 Email request processing & parsing

- **Deskripsi:** Aturan yang mengubah email masuk menjadi request baru atau komentar pada request yang sudah ada.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (pengirim), Agent Layanan (membalas lewat email), sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Email baru membuat request dengan subject sebagai summary (dipotong 255 karakter) dan body sebagai description; subject kosong menjadi "(no subject)".
  2. Balasan dengan key request di subject (mis. `[ABC-123]`) atau header `In-Reply-To`/`References` yang cocok dengan thread menjadi komentar pada request tersebut.
  3. Teks kutipan dan semua isi di bawah penanda "Reply above this line" dibuang dari komentar.
  4. ⚠ Opsi project untuk menyembunyikan email signature dari tampilan request dan portal.
  5. Body HTML dikonversi ke rich text dan inline image dipertahankan sebagai attachment yang dirender di posisinya.
  6. Attachment email menjadi attachment request sesuai batas ukuran (ADM-011); attachment di atas batas dibuang dan dicatat di email log.
  7. Penerima CC ditambahkan sebagai participant bila opsi aktif.
  8. Penerima BCC tidak pernah ditambahkan sebagai participant.
  9. Pengirim baru dibuat sebagai customer bila customer permissions dan signup mengizinkan (toggle "Create customer accounts from email").
  10. Email dari non-customer ditolak di project restricted dan tercatat dengan alasan "Pengirim bukan customer".
  11. ⚠ Opsi mengizinkan email dari pengirim eksternal (non-customer) pada request yang sudah ada ditambahkan sebagai komentar.
  12. Balasan email dari agent ke request menjadi komentar publik atas nama agent tersebut.
  13. Balasan ke request yang sudah resolved menjadi komentar; reopen otomatis dilakukan lewat rule otomasi ([20](20-otomasi.md)), bukan perilaku bawaan.
  14. Email auto-reply/out-of-office (header `Auto-Submitted`, `X-Autoreply`, `Precedence`) diabaikan dan tercatat di email log.
  15. ⚠ Satu email yang dikirim ke beberapa alamat project yang sama tidak membuat request duplikat (deduplikasi via `Message-ID`).
  16. Request dari email ditandai channel `email`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pemrosesan dijalankan worker asinkron idempoten berbasis `Message-ID`; aturan pemrosesan umum di NTF-033 ([19](19-notifikasi-email.md)). Batas global (ukuran pesan, jumlah per periode) dan pengaturan email request tingkat situs mengikuti JSM-AI-007 ([35](35-jsm-ai-otomasi-api-admin-csm.md)).

### JSM-PRT-042 Email loop prevention & limits

- **Deskripsi:** Pengaturan keamanan dan perlindungan email channel dari loop, spam, spoofing, dan volume berlebihan.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs, sistem.
- **Aturan perilaku (acceptance criteria):**
  1. Email yang berasal dari alamat produk atau alamat project sendiri diabaikan.
  2. ⚠ Throttling: maksimal N email per pengirim per periode (default 20 per 10 menit); email berlebih ditolak dan dicatat.
  3. ⚠ Batas ukuran email total dan attachment (default 25 MB per email) diterapkan; email di atas batas ditolak dengan alasan di log.
  4. ⚠ Mailbox dinonaktifkan otomatis setelah error koneksi berulang (mis. 10 kali berturut-turut) dan statusnya menjadi error.
  5. Admin project menerima peringatan (email dan notifikasi in-app) saat channel bermasalah atau email request gagal diproses.
  6. ⚠ Allowlist/blocklist pengirim atau domain per project; blocklist menolak email, allowlist (bila diisi) hanya menerima pengirim terdaftar.
  7. ⚠ Pengaturan DMARC untuk email masuk: tolak atau terima email yang gagal DMARC; default terima dengan penanda di log.
  8. Email bulk/auto-generated (header `Precedence: bulk/list`, `List-Unsubscribe` dengan pola mailing list) diabaikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Batas umum di NTF-009 ([19](19-notifikasi-email.md)) dan API-006 ([23](23-api-publik-webhook.md)).

### JSM-PRT-043 Email logs

- **Deskripsi:** Log email masuk yang diproses email channel beserta hasil dan alasannya.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Daftar email masuk per project dan per alamat menampilkan waktu, pengirim, subject, dan alamat tujuan.
  2. Status per email: request created, comment added, rejected, atau failed.
  3. Alasan penolakan/kegagalan ditampilkan (bukan customer, loop, field required, gagal DMARC, ukuran, auto-reply, project nonaktif).
  4. Email yang berhasil memiliki tautan ke request hasil.
  5. Log dapat difilter berdasarkan rentang tanggal dan status.
  6. ⚠ Log disimpan 30 hari lalu dihapus otomatis; isi body tidak disimpan, hanya metadata.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Log ini juga menjadi sumber alert JSM-PRT-042 butir 5.

## Widget, chat & live chat

### JSM-PRT-044 Embeddable widget

- **Deskripsi:** Widget bantuan JavaScript yang ditanam di situs web atau aplikasi agar pengunjung dapat mencari artikel KB dan membuat request tanpa membuka portal.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengonfigurasi), Customer/Pelapor dan pengunjung anonim (memakai).
- **Aturan perilaku (acceptance criteria):**
  1. Widget diaktifkan di pengaturan channel Widget project (default nonaktif).
  2. Admin memilih request type yang tampil di widget.
  3. ⚠ Hanya request type yang seluruh field-nya bertipe didukung widget (text, paragraph, select, date, attachment) yang dapat dipilih; request type lain tampil nonaktif dengan alasan.
  4. Pencarian artikel KB di widget dapat diaktifkan atau dimatikan.
  5. Kode embed berupa tag `<script>` dengan atribut `data-key` (kunci widget) dan base URL yang dapat disalin satu klik.
  6. Widget menampilkan tombol help mengambang di situs host dan membuka panel di atas halaman tanpa mengganggu skrip host (dirender dalam iframe terisolasi).
  7. Field email wajib karena pengguna widget tidak login.
  8. Setelah request dibuat, widget menampilkan pesan konfirmasi berisi key request.
  9. Request dari widget mengikuti customer permissions project: pada project restricted, email yang bukan customer ditolak dengan pesan galat.
  10. ⚠ Widget mendukung attachment sesuai batas attachment project.
  11. Setiap project memiliki snippet embed sendiri; kunci widget dapat diregenerasi untuk mencabut snippet lama.
  12. ⚠ Warna tombol dan posisi (kanan bawah/kiri bawah) dapat dikustomisasi.
  13. Request dari widget ditandai channel `widget`.
  14. Widget hanya dapat dimuat dari domain yang ada di daftar allowed origins project (bila diisi).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Widget dibangun sebagai bundle JS terpisah yang disajikan dari binary (embed) dengan endpoint publik ber-CORS terbatas dan rate limit.

### JSM-PRT-045 Chat channel – Slack (JSM Chat / Assist)

- **Alias inventaris:** Chat (Slack & Microsoft Teams)
- **Deskripsi:** Membuat dan menangani request dari percakapan Slack (dan Microsoft Teams, lihat JSM-PRT-046) melalui aplikasi chat native produk ini.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek (menghubungkan), Agent Layanan, Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Slack workspace dihubungkan dari pengaturan Chat project lewat OAuth instalasi aplikasi Slack milik produk.
  2. Admin memilih satu atau lebih request channel tempat customer meminta bantuan.
  3. Admin memilih agent triage channel (channel privat) yang menerima kartu request baru.
  4. Request dapat dibuat dari pesan lewat emoji reaction yang dikonfigurasi (mis. `:ticket:`) atau message action "Create request".
  5. Opsi membuat request otomatis untuk setiap pesan baru (bukan balasan thread) di request channel.
  6. Saat membuat request, modal Slack meminta request type dan field wajibnya.
  7. Balasan di thread Slack dan komentar request tersinkron dua arah.
  8. Requester menerima update status dan komentar lewat DM dari bot.
  9. Agent dapat assign, mengubah status, dan membalas langsung dari kartu di triage channel.
  10. Balasan agent di triage channel tersimpan sebagai komentar internal (internal note).
  11. Pengguna Slack dicocokkan ke customer berdasarkan email; akun baru mengikuti customer permissions dan signup.
  12. Virtual service agent dapat menjawab di request channel bila diaktifkan (modul 35).
  13. ⚠ Slack Connect untuk customer di workspace eksternal didukung bila workspace mengizinkan.
  14. Workspace Microsoft Teams dihubungkan lewat JSM-PRT-046 dengan konsep channel yang sama.
  15. Pesan di request channel tertentu dapat menjadi request sesuai butir 4-5.
  16. Request dapat dibuat dengan emoji reaction, shortcut, atau form modal.
  17. Triage channel menampilkan notifikasi request baru beserta aksi assign/transition.
  18. Customer dapat mengirim DM ke bot untuk membuat dan melacak request miliknya.
  19. Admin menetapkan request type default untuk request dari chat.
  20. Approver dapat menyetujui/menolak dari pesan chat (tombol Approve/Decline).
  21. Request dari chat ditandai channel `chat` dan dapat dicari lewat JQL `request-channel-type = chat`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Aplikasi Slack untuk notifikasi pengguna internal ada di INT-030 s.d. INT-035 ([25](25-integrasi-impor-ekspor-ekstensi.md)); disarankan satu aplikasi Slack dengan dua fungsi. Aplikasi pihak ketiga (mis. Halp lama) tidak dibangun.

### JSM-PRT-046 Chat channel – Microsoft Teams

- **Deskripsi:** Integrasi chat native di Microsoft Teams agar request dibuat dan ditangani di Teams.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Agent Layanan, Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Tenant Microsoft Teams dihubungkan lewat persetujuan admin tenant (Microsoft Entra ID) dan aplikasi bot produk.
  2. Admin memilih request channel dan agent triage channel.
  3. Request dibuat dari pesan lewat message action "Create request".
  4. Komentar tersinkron dua arah antara thread Teams dan request.
  5. Requester menerima notifikasi personal dari bot.
  6. Agent dapat assign, mengubah status, dan membalas dari Teams.
  7. ⚠ Kemampuan yang tidak tersedia di Teams dibanding Slack (mis. emoji reaction, request otomatis per pesan) ditampilkan sebagai nonaktif di pengaturan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Integrasi Teams untuk pengguna internal di INT-036 ([25](25-integrasi-impor-ekspor-ekstensi.md)).

### JSM-PRT-047 Live chat

- **Deskripsi:** Live chat native di portal/help center: ketika customer yang sedang mengobrol dengan virtual service agent meminta bicara dengan manusia, agent dapat masuk ke jendela chat dan membantu secara real time.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Customer/Pelapor, Agent Layanan, Admin Proyek (space admin).
- **Aturan perilaku (acceptance criteria):**
  1. Pengaturan berada di Project settings > Channels & self-service > Live chat.
  2. Live chat dapat diaktifkan atau dinonaktifkan per service project.
  3. Customer yang sedang mengobrol dengan virtual service agent dapat meminta eskalasi ke agent manusia.
  4. Chat yang dieskalasi masuk ke inbox Live chat khusus di service project dengan indikator chat menunggu.
  5. Agent bergabung ke chat, membalas real time (WebSocket), lalu menutup chat.
  6. ⚠ Bila tidak ada agent tersedia dalam batas waktu (mis. 2 menit) atau di luar jam layanan, customer ditawari membuat request.
  7. ⚠ Setiap sesi live chat dapat dibuat atau ditautkan ke work item beserta transkrip percakapan sebagai komentar.
  8. ⚠ Kanal yang didukung versi awal hanya portal/help center; widget dan kanal lain menyusul.
  9. Live chat berbeda dari Chat channel Slack/Teams (JSM-PRT-045/046) dan dari hand-off virtual service agent yang langsung membuat work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada virtual service agent (modul 35) dan infrastruktur realtime ([27](27-nfr-arsitektur-platform.md)). Aplikasi live chat Marketplace (mis. "Chat for JSM") di luar cakupan.

## Administrasi, pengaturan kanal & API

### JSM-PRT-048 Channels & self-service settings hub

- **Deskripsi:** Area pengaturan project yang mengumpulkan semua channel dan pengaturan self-service.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Menu Channels & self-service memuat Portal, Email, Chat, Live chat, Widget, Virtual service agent, dan Knowledge base; channel yang modulnya belum tersedia tidak ditampilkan.
  2. Setiap channel menampilkan status aktif/nonaktif dan dapat diaktifkan dari halaman ini.
  3. Hub menyediakan akses ke Customer permissions, Customer notifications, dan Language support.
  4. Hanya Admin Proyek (dan Admin Situs) yang dapat membuka hub; pengguna lain mendapat 403.
  5. Hub menampilkan tautan ke pengaturan help center (hanya aktif untuk Admin Situs).
- **Team-managed vs company-managed:** Tidak ada perbedaan; struktur menu sama untuk keduanya.
- **Catatan & ketergantungan:** ⚠ Label "project" vs "space" di menu mengikuti keputusan terminologi VISI-002 ([00](00-visi-ruang-lingkup-glosarium-baseline.md)).

### JSM-PRT-049 Help center administration permissions

- **Deskripsi:** Peran dan permission yang menentukan siapa boleh mengonfigurasi help center, portal, dan customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs, Admin Proyek, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Admin Situs (Administer) mengelola branding, layout, custom pages, topics, announcement help center dan login, signup/akses, serta multiple help centers.
  2. Admin Proyek (Administer projects) mengelola portal settings, portal groups, help resources, customer permissions, customer notifications, email channel, chat, widget, dan live chat project.
  3. Agent Layanan dapat menambah customer/organization bila diizinkan (JSM-PRT-016) dan raise request on behalf of customer.
  4. ⚠ Peran "Help center admin" khusus dapat diberikan ke pengguna non-admin situs untuk mengelola branding, layout, topics, dan announcement saja.
  5. Admin Situs (organisasi) mengelola customer accounts, SSO customer, custom domain, dan email sending domain.
  6. Setiap perubahan konfigurasi help center, portal, customer permissions, dan channel tercatat di audit log (ADM-014) dengan aktor, nilai lama, dan nilai baru.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Model permission di [03](03-izin-peran-keamanan.md); audit log di [22](22-administrasi-sistem.md).

### JSM-PRT-050 Help center & customer REST API

- **Deskripsi:** REST API Service Management (kompatibel `servicedeskapi`) untuk portal, request customer, customer, dan organization.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Integrator/Developer API (izin mengikuti pengguna pemilik token).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /rest/servicedeskapi/servicedesk`, `/servicedesk/{id}`, dan `/servicedesk/{id}/requesttype` (beserta `requesttypegroup` dan `/requesttype/{id}/field`) mengembalikan data yang dapat diakses pemanggil.
  2. `POST /rest/servicedeskapi/request` membuat request, termasuk `raiseOnBehalfOf` (hanya agent) dan `requestParticipants`.
  3. `GET /rest/servicedeskapi/request` mengembalikan daftar request pemanggil dengan filter `requestOwnership`, `requestStatus`, dan `searchTerm`.
  4. `GET/POST /request/{key}/comment` mendukung komentar `public` true/false; customer hanya dapat membuat dan melihat komentar publik.
  5. `GET /request/{key}/status` mengembalikan riwayat status customer-visible dan `GET/POST /request/{key}/transition` menjalankan customer transition.
  6. Attachment diunggah lewat `POST /servicedesk/{id}/attachTemporaryFile` lalu dilampirkan dengan `POST /request/{key}/attachment`.
  7. `GET/POST/DELETE /request/{key}/participant` mengelola participant.
  8. `GET /request/{key}/approval` dan `POST /request/{key}/approval/{id}` membaca dan memutuskan approval.
  9. `GET/POST/DELETE /request/{key}/feedback` mengelola CSAT.
  10. `GET/PUT/DELETE /request/{key}/notification` membaca, berlangganan, dan berhenti berlangganan notifikasi request.
  11. `POST /rest/servicedeskapi/customer` membuat customer portal-only (hanya admin).
  12. `GET/POST/DELETE /servicedesk/{id}/customer` mengelola customer project.
  13. `/organization`, `/organization/{id}/user`, dan `/servicedesk/{id}/organization` mengelola organization dan keanggotaannya.
  14. `/organization/{id}/property/{key}` mengelola organization properties (JSON maksimal 32 KB per properti).
  15. `GET /knowledgebase/article?query=` dan `/servicedesk/{id}/knowledgebase/article` mencari artikel KB sesuai izin.
  16. `GET /rest/servicedeskapi/info` mengembalikan versi dan informasi build.
  17. Endpoint eksperimental hanya dapat dipanggil dengan header `X-ExperimentalApi: opt-in`; tanpa header mengembalikan 403 dengan pesan penjelasan.
  18. Paginasi memakai `start`/`limit` dengan respons `size`, `isLastPage`, dan `_links`; semua hasil difilter berdasarkan izin pengguna pemanggil.
  19. Autentikasi memakai API token (Basic) atau OAuth 2.0 dengan scope Service Management (API-008, API-009).
  20. Endpoint internal SPA versi `/api/v1/...` memakai use case yang sama dengan endpoint kompatibilitas agar perilaku izin identik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konvensi umum, format error, dan rate limit di API-001, API-004, API-005, API-006 ([23](23-api-publik-webhook.md)); endpoint agen/queue/SLA di modul 35. Anotasi Swagger wajib diperbarui (`make swag`) dan endpoint diberi `@Security BearerAuth`.

## Pertanyaan terbuka

1. Apakah akun customer portal-only disimpan di tabel `users` yang sama (dengan kolom `account_type`) atau di tabel `customers` terpisah, mengingat dampaknya pada relasi reporter/participant dan pada login JWT yang ada (JSM-PRT-014)?
2. Apakah satu email boleh sekaligus memiliki akun portal-only dan akun pengguna situs di satu situs, atau selalu satu akun per email (JSM-PRT-014)?
3. Apakah pengguna berlisensi yang menjadi reporter request menerima customer notification, notifikasi internal, atau keduanya dengan deduplikasi (JSM-PRT-014, JSM-PRT-036)?
4. Bagaimana infrastruktur email masuk untuk alamat cloud bawaan: penerima SMTP/MX sendiri di dalam binary, layanan inbound pihak ketiga, atau hanya IMAP/POP dan forwarding (JSM-PRT-040)?
5. Untuk deployment self-hosted, apakah custom domain dan TLS dikelola produk (ACME) atau diserahkan ke reverse proxy (JSM-PRT-009)?
6. Apakah multiple help centers, custom pages, dan custom domain dibatasi feature flag tier Premium/Enterprise atau tersedia di semua tier (JSM-PRT-003, JSM-PRT-008, JSM-PRT-009)?
7. Apakah request anonim wajib mencantumkan email, dan bila ya apakah customer otomatis dibuat dari email tersebut (JSM-PRT-017)?
8. Apakah mode "safe customer notifications" dibangun sejak gelombang pertama Fase 5 sebagai kebutuhan kepatuhan, atau ditunda (JSM-PRT-036)?
9. Apakah kanal chat (Slack, Microsoft Teams) dan live chat tetap di Fase 7, atau Slack dimajukan ke Fase 5 karena prioritas P1 (JSM-PRT-045, JSM-PRT-047)?
10. Apakah peran khusus "Help center admin" diperlukan, atau cukup Admin Situs (JSM-PRT-049)?
11. Apakah impor customer via CSV disediakan di halaman Customers, atau impor massal hanya lewat REST API (JSM-PRT-020, JSM-PRT-033)?
12. Berapa batas default (alamat email per project, help center per situs, custom pages, throttling email, retensi email log) yang dipakai produk ini (JSM-PRT-003, JSM-PRT-008, JSM-PRT-040, JSM-PRT-042, JSM-PRT-043)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk portal dan help center (tema/CSS kustom penuh, editor halaman lanjutan, live chat pihak ketiga seperti "Chat for JSM", portal untuk customer dengan login sosial khusus, survei CSAT lanjutan, pembuat formulir eksternal) tidak dibangun sebagai fitur native; kebutuhan tersebut dipenuhi lewat REST API dan webhook (modul [23](23-api-publik-webhook.md)) atau ekstensi (modul [25](25-integrasi-impor-ekspor-ekstensi.md)).
- **Layanan Atlassian lintas produk** tidak dibangun: Atlassian Administration (admin.atlassian.com), id.atlassian.com, Atlassian Guard, Confluence sebagai penyimpan knowledge base, Atlassian Intelligence/Rovo, dan Halp/Assist. Produk ini menyediakan padanan native: konsol administrasi sendiri ([02](02-pengguna-grup-organisasi.md), [22](22-administrasi-sistem.md)), akun dan SSO sendiri ([01](01-identitas-autentikasi-sesi.md)), knowledge base internal (modul 31), dan AI opsional (modul 35).
- **Legacy help center dan landing pages** tidak dibangun sebagai pengalaman terpisah; hanya satu generasi help center (setara Modern help center) yang dibuat, dan landing pages hasil impor dikonversi menjadi custom pages (JSM-PRT-003).
- **Pengalaman customer di aplikasi mobile native** tidak dibangun; customer memakai web responsif (JSM-PRT-011).
- **Custom HTML/CSS/JavaScript** di help center tidak didukung secara native (JSM-PRT-002).
- **Fitur Data Center-only** (mis. portal customer Server/DC, plugin email handler lama, JSM Data Center customer portal themes) tidak ada di jatah modul ini; hanya dirujuk sebagai sumber impor di modul 25.
- **Billing dan lisensi customer** tidak dibangun; customer tidak memakai lisensi dan tier diperlakukan sebagai feature flag (VISI-003).
