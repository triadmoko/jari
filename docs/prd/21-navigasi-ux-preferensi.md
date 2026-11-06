# 21 — Navigasi Global, Halaman Personal & Preferensi

Modul ini mendefinisikan "kerangka" aplikasi web produk ini yang dipakai oleh semua modul lain: top bar global, sidebar dan navigasi di dalam proyek (label UI "Spaces", kunci terminologi VISI-002), breadcrumbs, app/site switcher, halaman personal "For you", recent dan starred items, dialog Create global, quick search, hub notifikasi in-app, menu Settings dan Profile, halaman profil dan hover card, pengaturan akun dan preferensi personal (bahasa, zona waktu, theme, aksesibilitas, autowatch, notifikasi personal), keyboard shortcut dan Commands dialog, skema URL dan deep link, komponen umpan balik (flag, banner, error/empty state), pembaruan real-time di UI, persistensi state UI per pengguna, onboarding, prompt edisi/plan, serta kebijakan responsif, browser, aksesibilitas (WCAG 2.1 AA), cetak, dan aplikasi mobile. Modul ini hanya mengatur entry point dan perilaku UX lintas halaman; logika domain (isi work item, JQL, notification scheme, board, dashboard, autentikasi) tetap milik modul masing-masing dan dirujuk silang. Kondisi repo saat ini: SPA SvelteKit 5 di `frontend/` sudah memiliki pustaka komponen UI (sidebar, navigation-menu, drawer, tooltip, skeleton, empty, alert, resizable) dan hook `is-mobile`, tetapi belum memiliki shell navigasi, halaman personal, maupun penyimpanan preferensi; backend baru menyediakan auth + user dengan deteksi bahasa EN/ID dari header `Accept-Language`. Fitur asisten AI (setara Rovo/Atlassian Intelligence) dicatat sebagai opsional.

Modul terkait:

- [00 Visi, Ruang Lingkup, Glosarium & Baseline Repo](00-visi-ruang-lingkup-glosarium-baseline.md) (terminologi Work item/Space, baseline repo)
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) (akun, profil & visibilitas, sesi, API token, logout)
- [02 Manajemen Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) (managed account, tim, batas user)
- [03 Izin, Peran Proyek & Keamanan Work Item](03-izin-peran-keamanan.md) (permission yang menyaring item navigasi)
- [04 Proyek: Tipe, Template, Pengaturan & Siklus Hidup](04-proyek.md) (header proyek, fitur proyek, arsip)
- [05 Konfigurasi Work Item: Tipe, Field, Layar & Skema](05-konfigurasi-work-item-skema.md) (create screen)
- [06 Work Item Inti: Buat, Lihat, Ubah & Operasi](06-work-item-inti.md) (work item view, create, operasi)
- [07 Editor Teks Kaya & Rendering Konten](07-editor-konten-kaya.md) (smart link di editor, shortcut editor)
- [10 Board: Kolom, Kartu, Swimlane & Kanban](10-board-kolom-kartu.md) dan [11 Backlog & Sprint (Scrum)](11-backlog-sprint.md) (aksi yang dipicu shortcut board/backlog)
- [12 Bahasa Kueri (JQL-kompatibel)](12-jql-bahasa-kueri.md) dan [13 Pencarian, Navigator, Filter Tersimpan & Ekspor](13-pencarian-filter-ekspor.md) (search work items, filter, export printable)
- [17 Dashboard & Gadget](17-dashboard-gadget.md)
- [18 Timeline, Rencana Lintas Proyek & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) (Plans, Teams, Goals)
- [19 Notifikasi In-app & Email](19-notifikasi-email.md) (event, notification scheme, email)
- [22 Administrasi Sistem & Pengaturan Global](22-administrasi-sistem.md) (look and feel, announcement banner, default bahasa/zona waktu)
- [20 Otomasi Berbasis Aturan](20-otomasi.md) (agent AI di aturan otomasi, pembuatan aturan dengan bahasa natural)
- [23 REST API Publik & Webhook](23-api-publik-webhook.md) (API recent projects, user properties)
- [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) (extension point UI, Slack/Teams)
- [27 Kebutuhan Non-fungsional & Arsitektur Platform](27-nfr-arsitektur-platform.md) (realtime, kinerja, browser)
- [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) (entitas preferensi, user properties, aktivitas)
- Service Management: [29](29-jsm-service-project-request-type-formulir.md), [30](30-jsm-help-center-portal-pelanggan-kanal.md), [31](31-jsm-antrian-sla-knowledge-base-laporan.md), [32](32-jsm-itsm-insiden-problem-perubahan-layanan.md), [33 Assets](33-jsm-assets-cmdb.md), [34 Operations](34-jsm-operations-alerting-on-call.md), [35](35-jsm-ai-otomasi-api-admin-csm.md); Product Discovery: [36](36-jpd-ideas-field-views.md), [37](37-jpd-insights-delivery-peran-integrasi.md) (entry point Queues, Assets, Operations, Ideas)

## Persona terkait

- **Developer/Anggota Tim** — memakai top bar, sidebar, For you, quick search, Create, notifikasi, shortcut keyboard, dan preferensi personal setiap hari.
- **Pengguna Tim Bisnis** — membutuhkan navigasi yang mudah dipahami, onboarding, theme, dan bahasa UI lokal.
- **Scrum Master/Team Lead** — bergantung pada shortcut board/backlog, starred boards, dan pembaruan real-time board.
- **Product Owner/PM** — memakai starred filters/dashboards/plans, share link, dan entry point Plans/Goals.
- **Stakeholder/Viewer** — membuka tautan yang dibagikan (deep link), melihat halaman profil dan hover card, menerima state "tidak punya akses" yang jelas.
- **Admin Situs** — mengatur branding top bar, announcement banner, default bahasa/zona waktu, onboarding kustom, dan melihat prompt plan/limit.
- **Admin Proyek/Project Lead** — mengelola navigasi proyek (views, shortcuts, link Project settings).
- **Integrator/Developer API** — membuat ekstensi UI pada extension point navigasi dan memakai API recent/user properties.
- **Tamu/Kolaborator Eksternal** — membuka halaman lewat deep link dan anonymous access terbatas.
- **Agent Layanan** dan **Product Manager Discovery** — memakai entry point Queues/Ideas di sidebar dan For you.

## User story modul

1. Sebagai Developer/Anggota Tim, saya ingin top bar dan sidebar yang konsisten di setiap halaman agar saya selalu bisa berpindah ke proyek, filter, dan dashboard tanpa kehilangan konteks (NAV-001, NAV-002, NAV-003).
2. Sebagai Developer/Anggota Tim, saya ingin halaman For you yang menampilkan pekerjaan yang saya kerjakan, lihat, dan di-assign ke saya agar saya bisa melanjutkan pekerjaan terakhir dengan cepat (NAV-010, NAV-011).
3. Sebagai Product Owner/PM, saya ingin menandai board, filter, dan dashboard dengan bintang agar item yang sering saya pakai selalu ada di sidebar (NAV-012).
4. Sebagai Developer/Anggota Tim, saya ingin membuka dialog Create dari mana saja dengan tombol `c` dan proyek/work type ter-prefill agar mencatat pekerjaan baru tidak memutus alur kerja saya (NAV-013).
5. Sebagai Developer/Anggota Tim, saya ingin mengetik key work item di quick search lalu Enter agar langsung terbuka tanpa melalui navigator (NAV-014).
6. Sebagai Developer/Anggota Tim, saya ingin panel notifikasi dengan tab Direct dan Watching serta penanda read/unread agar saya tidak melewatkan mention dan assignment (NAV-018, NAV-019).
7. Sebagai Pengguna Tim Bisnis, saya ingin memilih bahasa Indonesia, zona waktu, dan theme gelap agar produk nyaman dipakai sesuai lokasi dan kebiasaan saya (NAV-033, NAV-034).
8. Sebagai Scrum Master/Team Lead, saya ingin shortcut keyboard di board dan backlog (j/k, n/p, s+t) agar mengelola sprint lebih cepat tanpa mouse (NAV-038, NAV-040, NAV-042).
9. Sebagai Stakeholder/Viewer, saya ingin URL yang stabil dan bisa dibagikan untuk setiap work item, board, dan filter agar tautan di chat atau email selalu membuka halaman yang benar (NAV-043, NAV-016).
10. Sebagai Developer/Anggota Tim, saya ingin board dan work item view diperbarui real-time saat rekan mengubah data agar saya tidak bekerja di atas informasi basi (NAV-046).
11. Sebagai Stakeholder/Viewer, saya ingin pesan yang jelas saat halaman tidak ditemukan atau saya tidak punya akses, lengkap dengan tombol "Request access", agar saya tahu langkah selanjutnya (NAV-045).
12. Sebagai Admin Situs, saya ingin memasang announcement banner dan logo kustom agar seluruh pengguna melihat informasi penting dan identitas organisasi (NAV-001, NAV-044).
13. Sebagai Developer/Anggota Tim yang memakai keyboard dan screen reader, saya ingin semua fungsi bisa dioperasikan dengan keyboard dan screen reader agar saya bisa bekerja setara dengan rekan lain (NAV-035, NAV-036).
14. Sebagai Pengguna Tim Bisnis baru, saya ingin checklist Quickstart dan tur singkat agar saya cepat memahami cara memakai proyek pertama saya (NAV-023).
15. Sebagai Developer/Anggota Tim, saya ingin lebar sidebar, view terakhir, dan kolom list saya tersimpan agar UI kembali seperti terakhir saya tinggalkan di perangkat mana pun (NAV-037).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur | Fitur |
|---|---|---|
| P0 | 18 | NAV-001, NAV-002, NAV-003, NAV-004, NAV-010, NAV-011, NAV-012, NAV-013, NAV-014, NAV-018, NAV-019, NAV-025, NAV-026, NAV-032, NAV-033, NAV-043, NAV-044, NAV-045 |
| P1 | 18 | NAV-016, NAV-017, NAV-020, NAV-024, NAV-027, NAV-028, NAV-029, NAV-030, NAV-031, NAV-034, NAV-036, NAV-037, NAV-038, NAV-039, NAV-041, NAV-046, NAV-047, NAV-048 |
| P2 | 9 | NAV-005, NAV-006, NAV-007, NAV-015, NAV-021, NAV-023, NAV-035, NAV-040, NAV-042 |
| P3 | 11 | NAV-008, NAV-009, NAV-022, NAV-049, NAV-050, NAV-051, NAV-052, NAV-053, NAV-054, NAV-055, NAV-056 |

| Fase | Jumlah fitur | Fitur |
|---|---|---|
| Fase 1 | 17 | NAV-001, NAV-002, NAV-003, NAV-004, NAV-010, NAV-011, NAV-012, NAV-013, NAV-014, NAV-018, NAV-025, NAV-026, NAV-032, NAV-033, NAV-043, NAV-044, NAV-045 |
| Fase 2 | 17 | NAV-016, NAV-019, NAV-020, NAV-027, NAV-028, NAV-029, NAV-030, NAV-031, NAV-034, NAV-036, NAV-037, NAV-038, NAV-039, NAV-041, NAV-046, NAV-047, NAV-048 |
| Fase 3 | 11 | NAV-005, NAV-007, NAV-017, NAV-021, NAV-022, NAV-023, NAV-024, NAV-035, NAV-040, NAV-042, NAV-049 |
| Fase 4 | 8 | NAV-006, NAV-008, NAV-009, NAV-015, NAV-050 (padanan PWA + Web Push; app native Fase —), NAV-051, NAV-052, NAV-053 |
| Fase — | 3 | NAV-054, NAV-055, NAV-056 (referensi legacy/Data Center, tidak dibangun) |

Strategi urutan implementasi: Fase 1 membangun shell SPA lebih dulu (top bar, sidebar, navigasi proyek, breadcrumbs) beserta skema URL stabil dan komponen umpan balik (flag, error/empty/permission state), karena semua modul inti lain dirender di dalam shell ini dan bergantung pada fallback `index.html` yang sudah ada di `webui`. Bersamaan dengan itu dibangun pelacakan recent/starred dan halaman For you, dialog Create global, quick search dasar, hub notifikasi in-app, serta menu Settings/Profile dengan logout dan preferensi bahasa/zona waktu (EN/ID). Fase 2 menambah halaman profil, hover card, pengaturan personal (notifikasi, autowatch, theme), persistensi state UI server-side, katalog shortcut global dan work item, pembaruan real-time, serta baseline aksesibilitas, responsif, dan dukungan browser yang diuji di CI. Fase 3 menambah kustomisasi sidebar, app switcher, smart links, what's new/feature flag, onboarding, Commands dialog, shortcut board/backlog lengkap, aksesibilitas personal, dan prompt edisi plan; Fase 4 menyusul untuk multi-workspace, extension point UI, entry point Plans/Teams/Goals/Assets, padanan mobile berbasis PWA + Web Push (app native di luar cakupan, NAV-050), dan fitur asisten AI opsional.

## Fitur

## Kerangka navigasi global

### NAV-001 Top Navigation Bar (new navigation 2025)

**Deskripsi:** Bar atas global yang selalu tampil di seluruh halaman produk ini. Isinya kontrol tingkat aplikasi: toggle sidebar, app switcher, logo produk, search, tombol Create, asisten AI (opsional), Notifications, Help, Settings, dan Profile.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna terautentikasi; pengguna anonim (bila anonymous access diaktifkan) melihat versi terbatas; Admin Situs mengatur branding. Visibilitas item mengikuti global permission dan lisensi produk ([modul 03](03-izin-peran-keamanan.md)).

**Aturan perilaku (acceptance criteria):**

1. Urutan dari kiri ke kanan: tombol expand/collapse sidebar, app switcher (ikon 9 titik), lalu logo + nama aplikasi; klik logo membuka landing default (For you, NAV-010).
2. Area tengah berisi search bar dengan placeholder "Search" (NAV-014), diikuti tombol "+ Create" berwarna primer yang membuka dialog Create global (NAV-013).
3. Area kanan berisi, berurutan: tombol asisten AI (hanya bila fitur AI aktif untuk workspace, NAV-051), Notifications (ikon bell + badge unread), Help (?), Settings (cog), dan avatar Profile.
4. ⚠ Badge unread menampilkan angka bila jumlah unread 1-9 dan "9+" bila lebih; badge disembunyikan saat jumlah unread 0.
5. Pada lebar layar sempit (breakpoint tablet/mobile, NAV-047) search bar menyusut menjadi ikon, tombol Create menjadi ikon "+", dan item kanan yang tidak muat masuk ke menu overflow tanpa ada fungsi yang hilang.
6. Logo kustom dan warna branding yang diatur Admin Situs (look and feel, [modul 22](22-administrasi-sistem.md)) diterapkan di top bar dan tetap kontras di theme light maupun dark (NAV-034).
7. Setiap ikon memiliki tooltip dan `aria-label`; semua kontrol dapat dicapai lewat Tab secara berurutan dengan focus ring yang terlihat.
8. Skip links "Skip to: Main content / Sidebar / Top bar" muncul saat fokus keyboard pertama pada halaman dan memindahkan fokus ke landmark yang dipilih.
9. Top bar tetap sticky saat halaman di-scroll; z-index-nya berada di atas konten halaman tetapi di bawah modal/dialog dan flag.
10. Item disembunyikan sesuai permission: bagian administrasi di menu Settings hanya tampil untuk admin, sedangkan item personal (profil, preferensi) selalu tampil untuk pengguna terautentikasi.
11. ⚠ Nama workspace/site aktif tidak ditampilkan di bar utama, melainkan di app switcher (NAV-007) dan site switcher (NAV-008).
12. ⚠ Pada workspace Free atau trial ditampilkan lozenge edisi/trial atau CTA "Upgrade"/"Try Premium" (NAV-024), dan pada workspace sandbox ditampilkan indikator "Sandbox" yang selalu terlihat.
13. Pengguna anonim (bila anonymous access aktif) melihat top bar terbatas tanpa Create dan Notifications, dengan tombol "Log in" di posisi avatar.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Implementasi pada SPA `frontend/` memakai komponen `navigation-menu`, `tooltip`, dan `drawer` yang sudah ada; shell dipasang di `+layout.svelte` route terautentikasi.
- Badge notifikasi bergantung pada hub notifikasi (NAV-018) dan kanal realtime (NAV-046).
- Penamaan "Work items"/"Spaces" mengikuti glosarium di [00](00-visi-ruang-lingkup-glosarium-baseline.md). Prosa modul ini memakai "proyek" (sama dengan modul lain); "Spaces"/"Space" hanya muncul sebagai label UI EN yang dikutip dan diambil dari kunci terminologi VISI-002, sehingga set terminologi "klasik" menampilkan "Projects"/"Project settings".

### NAV-002 Sidebar Navigation (new navigation 2025)

**Deskripsi:** Sidebar kiri persisten yang berisi semua entry point utama beserta isi bersarang (proyek, filters, dashboards, plans). Ukurannya bisa diubah, bisa di-collapse, dan dipersonalisasi per pengguna.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna terautentikasi. Tombol create per section dikontrol global permission terkait (Create project, Create shared objects/filters & dashboards, plan permission); item produk tambahan (Service Management, Product Discovery, Assets) mengikuti lisensi produk.

**Aturan perilaku (acceptance criteria):**

1. Item top-level default: For you, Recent, Starred, Apps, Plans, Spaces, Filters, Dashboards, Teams, Goals; ditambah Assets, Operations, dan Customer service hanya bila produk terkait aktif di workspace.
2. Recent dan Starred membuka flyout berisi daftar item yang dikelompokkan per tipe (work items, proyek, boards, filters, dashboards, plans) dengan tautan "View all" (NAV-011, NAV-012).
3. Section Spaces dapat di-expand dan menampilkan proyek ber-star, proyek terbaru (recent), "View all spaces", serta tombol "+" Create project yang hanya tampil bila pengguna memiliki permission Create project.
4. Setiap proyek dapat di-expand untuk menampilkan views-nya (Summary, Board, Backlog, List, dll., NAV-003) dan menu "..." (Star, Project settings, Archive, Copy link) yang itemnya disaring sesuai permission pengguna di proyek tersebut.
5. Section Filters menampilkan starred filters, filter default (My open work items, Reported by me, All work items, dll.), "View all filters", tautan "Search work items", dan tombol "+" Create filter ([modul 13](13-pencarian-filter-ekspor.md)).
6. Section Dashboards menampilkan starred/recent dashboards, "View all dashboards", dan "+" Create dashboard ([modul 17](17-dashboard-gadget.md)).
7. Section Plans (tier premium) menampilkan starred/recent plans, "View all plans", dan "+" Create plan; pada edisi Free/Standard section ini tampil sebagai CTA upgrade atau disembunyikan sesuai pengaturan (NAV-006, NAV-024).
8. Item yang sesuai dengan route saat ini di-highlight dengan selected state dan atribut `aria-current="page"`.
9. Lebar sidebar dapat diubah dengan drag pada tepi kanan dalam batas minimum 240 px dan maksimum 480 px (nilai awal 280 px) dan tersimpan per pengguna (NAV-037).
10. Sidebar dapat di-collapse/expand lewat tombol di top bar atau shortcut `[`; saat collapsed, hover pada tepi kiri memunculkan flyout sementara (peek) yang tertutup saat kursor keluar.
11. Pada layar kecil sidebar menjadi overlay/drawer yang menutupi konten dan tertutup otomatis setelah pengguna memilih tujuan navigasi.
12. Menu "More" berisi item top-level yang disembunyikan pengguna (NAV-005).
13. State expanded/collapsed setiap section bersarang tersimpan per pengguna dan dipulihkan saat halaman dibuka kembali.
14. Setiap item dirender sebagai anchor link asli sehingga klik tengah atau Cmd/Ctrl+klik membuka tujuan di tab baru.
15. Tombol "+" inline per section hanya tampil bila pengguna memiliki permission create yang relevan; bila tidak, tombol tidak dirender (bukan sekadar dinonaktifkan).
16. Sidebar dapat dinavigasi dengan keyboard: Tab berpindah antar section, Arrow atas/bawah antar item, Arrow kanan/kiri expand/collapse, Enter membuka, dan Esc menutup flyout.

**Team-managed vs company-managed:** Tidak ada perbedaan pada sidebar global; perbedaan isi per proyek dibahas di NAV-003.

**Catatan & ketergantungan:**

- Komponen `sidebar` dan `resizable` di `frontend/src/lib/components/ui` dapat dipakai sebagai dasar.
- Teams dan Goals mengarah ke halaman yang didefinisikan [modul 18](18-timeline-rencana-tampilan-proyek.md); Assets ke [modul 33](33-jsm-assets-cmdb.md); Operations ke [modul 34](34-jsm-operations-alerting-on-call.md); Customer service ke [modul 35](35-jsm-ai-otomasi-api-admin-csm.md).
- Data recent/starred berasal dari NAV-011/NAV-012; keputusan terbuka: batas min/max lebar sidebar di atas adalah usulan awal.

### NAV-003 Project/Space Navigation within Sidebar

**Deskripsi:** Perilaku navigasi di dalam sebuah proyek dari sidebar dan header, termasuk perbedaan team-managed dan company-managed.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna dengan permission Browse projects pada proyek; Admin Proyek/Project Lead untuk Project settings, views, dan shortcuts.

**Aturan perilaku (acceptance criteria):**

1. Header proyek menampilkan avatar, nama, tipe (team-managed/company-managed), tombol star (NAV-012), dan menu "..." yang itemnya disaring permission.
2. Views/tabs proyek mencakup Summary, Timeline, Backlog, Board, Calendar, List, Forms, Goals, Reports, Development/Code, Releases, Components, Pages, Shortcuts, dan Apps; hanya view yang sesuai template dan fitur aktif proyek yang ditampilkan.
3. Pada proyek team-managed, fitur (Backlog, Sprints, Timeline, Reports, Development, dll.) dapat dinyalakan/dimatikan per proyek oleh admin proyek, dan item navigasi terkait langsung muncul/hilang tanpa reload.
4. Pada proyek company-managed tersedia board switcher yang menampilkan semua board yang terkait proyek, termasuk board berbasis filter lintas proyek, dan mengingat board terakhir yang dibuka (NAV-037).
5. ⚠ View tabs horizontal memiliki tombol "+" untuk menambah view; admin proyek dapat mengurutkan ulang dan menyembunyikan tab, dan urutan tersebut berlaku untuk semua anggota proyek.
6. Shortcuts proyek (tautan URL kustom yang ditambahkan admin proyek) tampil di navigasi proyek dan membuka URL eksternal di tab baru.
7. Tautan "Project settings" hanya tampil bagi pengguna dengan permission Administer projects pada proyek tersebut.
8. Proyek Service Management menampilkan Queues, Service requests, dan Knowledge base; proyek Product Discovery menampilkan daftar Ideas views (rinci di [modul 29](29-jsm-service-project-request-type-formulir.md) sampai [31](31-jsm-antrian-sla-knowledge-base-laporan.md) dan [modul 36](36-jpd-ideas-field-views.md)).

**Team-managed vs company-managed:** Team-managed: fitur dan view diaktifkan per proyek oleh admin proyek, tidak ada board switcher (satu board per proyek). Company-managed: board switcher multi-board, board bisa berbasis filter lintas proyek, konfigurasi melalui skema.

**Catatan & ketergantungan:**

- Konfigurasi fitur proyek dan tipe proyek di [04 Proyek](04-proyek.md); board di [10](10-board-kolom-kartu.md); backlog di [11](11-backlog-sprint.md).
- Status aktif tab mengikuti URL (NAV-043) agar dapat dibagikan.

### NAV-004 Breadcrumbs

**Deskripsi:** Jejak navigasi hierarkis di header halaman yang menunjukkan lokasi pengguna dan memudahkan naik level.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna; segmen yang tidak dapat diakses pengguna (mis. parent di proyek lain tanpa izin) ditampilkan tanpa tautan.

**Aturan perilaku (acceptance criteria):**

1. Pada work item view breadcrumb berbentuk Spaces / Nama proyek / [Parent] / KEY-123, dan setiap segmen dapat diklik untuk menuju halaman terkait.
2. Hierarki parent multi-level (mis. Initiative > Epic > Story > Subtask) ditampilkan lengkap dengan ikon work type per segmen.
3. Pada board/backlog breadcrumb berbentuk Spaces / Nama proyek / Nama board, dan pada proyek company-managed segmen board berfungsi sebagai board switcher.
4. Pada halaman pengaturan proyek breadcrumb berbentuk Spaces / Nama proyek / Project settings / sub-halaman.
5. ⚠ Pada halaman administrasi global breadcrumb menampilkan kategori admin (mis. Settings / Work items / Custom fields).
6. Pada lebar sempit segmen tengah dipotong dengan ellipsis dan tooltip berisi teks lengkap; segmen yang disembunyikan dapat dibuka lewat dropdown "...".
7. Work item key di segmen terakhir memiliki aksi copy link yang menyalin URL kanonik `/browse/KEY-123` dan menampilkan flag "Link copied".
8. Label breadcrumb mengikuti terminologi baru "Spaces" dan "Work items" sesuai glosarium.
9. Breadcrumb dirender sebagai `nav` dengan `aria-label="Breadcrumbs"` dan item terakhir memiliki `aria-current="page"`.

**Team-managed vs company-managed:** Company-managed menampilkan board switcher di segmen board; team-managed menampilkan nama board tunggal tanpa switcher.

**Catatan & ketergantungan:**

- Hierarki parent di atas Epic bergantung pada konfigurasi hierarki ([modul 05](05-konfigurasi-work-item-skema.md) dan [modul 18](18-timeline-rencana-tampilan-proyek.md)).

### NAV-005 Customize Sidebar

**Deskripsi:** Kemampuan pengguna mengatur item top-level yang tampil di sidebar dan yang dipindah ke menu "More". Preferensi disimpan per pengguna lintas sesi dan perangkat.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Semua pengguna terautentikasi untuk sidebar miliknya sendiri; tidak ada pengaturan default sidebar tingkat workspace pada fase ini.

**Aturan perilaku (acceptance criteria):**

1. Opsi "Customize sidebar" tersedia dari menu "More" dan dari menu "..." di sidebar, membuka panel pengaturan sidebar.
2. Panel menampilkan toggle show/hide per item top-level (Recent, Starred, Apps, Plans, Filters, Dashboards, Teams, Goals, dll.); For you dan Spaces tidak dapat disembunyikan.
3. Item yang disembunyikan tetap dapat diakses lewat menu "More" dan route-nya tetap berfungsi bila dibuka langsung lewat URL.
4. Menu "..." pada setiap item top-level memiliki aksi cepat "Hide from sidebar" yang langsung memindahkan item ke "More" dan menampilkan flag dengan aksi "Undo".
5. ⚠ Item top-level dapat diurutkan ulang dengan drag-and-drop, dengan alternatif keyboard/menu "Move up"/"Move down" untuk aksesibilitas.
6. Preferensi tersimpan server-side per pengguna (user properties, NAV-037) dan berlaku di semua browser/perangkat setelah login.
7. Item yang tidak diizinkan atau tidak berlisensi untuk pengguna tidak muncul di daftar kustomisasi dan tidak dapat dimunculkan.
8. ⚠ Tombol "Reset to default" mengembalikan visibilitas dan urutan item ke default produk setelah konfirmasi.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Bergantung pada NAV-002 dan penyimpanan preferensi NAV-037.

### NAV-006 Plans, Teams, Goals & Assets entry points

**Deskripsi:** Entry point sidebar menuju area lintas proyek: Plans (perencanaan lintas proyek), Teams, Goals, dan Assets.

**Prioritas / Fase / Tier:** P2 / Fase 4 / premium

**Aktor & izin:** Semua pengguna untuk Teams/Goals; pengguna di workspace Premium/Enterprise untuk Plans; pengguna dengan akses Assets untuk Assets. Create plan/team mengikuti permission di [modul 18](18-timeline-rencana-tampilan-proyek.md).

**Aturan perilaku (acceptance criteria):**

1. Plans menampilkan daftar, recent, dan starred plans beserta "Create plan" hanya pada edisi Premium/Enterprise; pada Free/Standard ditampilkan CTA upgrade (NAV-024).
2. Teams menampilkan tim pengguna, "View all teams", "Create team", dan tautan ke team profile.
3. Goals menampilkan goals yang diikuti pengguna, tautan ke halaman Goals, dan jumlah work item yang tertaut ke setiap goal.
4. ⚠ Assets tampil hanya bila modul Assets aktif untuk workspace (Service Management atau edisi Premium/Enterprise) dan pengguna memiliki akses ke minimal satu skema objek.
5. Item hilang dari sidebar bila produk/tier tidak tersedia; item yang tersedia dapat disembunyikan pengguna lewat Customize Sidebar (NAV-005).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Fungsi Plans/Teams/Goals dirinci di [modul 18](18-timeline-rencana-tampilan-proyek.md); Assets di [modul 33](33-jsm-assets-cmdb.md). Modul ini hanya mendefinisikan entry point.

### NAV-007 App Switcher (Product Switcher)

**Deskripsi:** Menu ikon 9 titik untuk berpindah antar aplikasi produk ini (pelacak kerja, Service Management, Product Discovery) dan workspace yang dapat diakses pengguna, serta menemukan aplikasi yang belum diaktifkan.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Semua pengguna; Admin Situs/organisasi melihat tautan Administration dan mengatur product discovery/request.

**Aturan perilaku (acceptance criteria):**

1. Menampilkan daftar aplikasi pada workspace saat ini yang dapat diakses pengguna (setara Jira, Service Management, Product Discovery, dan aplikasi tambahan yang diaktifkan); aplikasi tanpa lisensi untuk pengguna tidak ditampilkan di bagian ini.
2. Menampilkan bagian "Recent" berisi aplikasi/workspace yang terakhir dibuka dan bagian per workspace.
3. Pengguna dengan akses ke banyak workspace melihat daftar workspace (nama + URL) beserta aplikasinya (NAV-008).
4. Tautan "Administration" ke konsol admin hanya tampil untuk admin organisasi/workspace.
5. Bagian "Discover more"/"Try" menampilkan aplikasi yang belum diaktifkan; admin dapat mematikan product discovery dan product request sehingga bagian ini disembunyikan untuk semua pengguna.
6. Non-admin dapat mengirim "Request access" atau permintaan trial untuk aplikasi yang belum berlisensi; permintaan dikirim ke admin sebagai notifikasi dan email, dan pengguna melihat flag konfirmasi.
7. Klik aplikasi membuka di tab yang sama; Cmd/Ctrl+klik atau klik tengah membuka di tab baru.
8. Setiap aplikasi ditampilkan dengan logo dan warna identitasnya sendiri yang konsisten dengan theme light/dark.
9. Menu dapat diakses dengan keyboard dan memakai pola ARIA menu (role `menu`/`menuitem`, Arrow untuk berpindah, Esc menutup dan mengembalikan fokus ke tombol).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Untuk instalasi single-workspace cukup stub minimal berisi aplikasi aktif; struktur data tetap mendukung multi-workspace (keputusan terbuka di [modul 01](01-identitas-autentikasi-sesi.md)).
- Produk Atlassian lain (Confluence, Compass, Loom, Bitbucket, Trello) tidak dibangun; lihat Di luar cakupan.

### NAV-008 Site Switcher & Multi-account Switching

**Deskripsi:** Perpindahan antar workspace (instance/site) dan antar beberapa akun yang login di browser yang sama.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Semua pengguna; admin organisasi Enterprise mengelola banyak workspace.

**Aturan perilaku (acceptance criteria):**

1. Daftar semua workspace yang dapat diakses akun ditampilkan di app switcher (NAV-007).
2. Workspace aktif ditandai dengan indikator (nama + URL) dan centang.
3. "Switch account" di profile menu (NAV-026) menampilkan akun lain yang sudah login di browser dan berpindah tanpa memasukkan ulang kredensial selama sesinya masih valid.
4. "Add another account" membuka alur login untuk akun tambahan tanpa mengakhiri sesi akun lain.
5. Deep link ke workspace lain dibuka memakai akun yang berwenang; bila tidak ada akun yang berwenang ditampilkan halaman "You don't have access" dengan opsi ganti akun (NAV-045).
6. ⚠ Edisi Enterprise mendukung banyak workspace dalam satu organisasi (target hingga sekitar 150 workspace) dengan app/site switcher yang tetap responsif.
7. Pengguna dapat logout dari satu akun atau dari semua akun sekaligus (NAV-032).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Dalam produk ini dipetakan ke workspace/tenant switcher; bergantung pada model multi-workspace dan sesi multi-akun di [01](01-identitas-autentikasi-sesi.md).

### NAV-009 Apps Navigation & Forge/Connect UI Extension Points

**Deskripsi:** Entry point navigasi untuk ekstensi terpasang (global page, project page, admin page, personal settings) serta modul UI ekstensi yang dirender di dalam produk ini.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Semua pengguna melihat ekstensi yang lolos display conditions; Admin Situs memasang/mengelola ekstensi; Integrator/Developer API membangunnya.

**Aturan perilaku (acceptance criteria):**

1. Section "Apps" di sidebar menampilkan daftar ekstensi yang memiliki global page, tautan "Explore more apps" ke katalog ekstensi, dan "Manage your apps" (khusus admin).
2. Non-admin dapat mengirim "Request app" ke admin bila fitur app request diaktifkan admin.
3. Extension point global page (setara `jira:globalPage`) tampil di Apps dengan route `/jira/apps/{appId}/{envId}` dan dapat memiliki sub-pages.
4. Extension point project page (setara `jira:projectPage`) tampil di navigasi proyek (NAV-003).
5. Extension point admin page dan project settings page (setara `jira:adminPage`, `jira:projectSettingsPage`) tampil di area administrasi dan Project settings.
6. Extension point personal settings page (setara `jira:personalSettingsPage`) tampil sebagai tab di Personal settings (NAV-030).
7. Work item view menyediakan extension point issuePanel, issueContext, issueAction, issueActivity, dan issueGlance.
8. Extension point lain (dashboard gadget, custom field/custom field type, background script work item view, workflow validator/condition/post function) didefinisikan di modul domainnya masing-masing.
9. Ekstensi dapat memakai komponen UI bawaan (UI Kit) atau UI kustom dalam iframe ber-sandbox dengan bridge API; design tokens theme aktif diteruskan ke ekstensi.
10. ⚠ Model ekstensi lama setara Connect (generalPages, webPanels, webItems) tidak dibangun; bila dibutuhkan kompatibilitas, ditandai deprecated sejak awal.
11. Ekstensi hanya tampil bila terpasang/berlisensi dan pengguna lolos display conditions (permission, tipe proyek, produk).
12. Halaman ekstensi ditandai ikon ekstensi dan memiliki tautan ke halaman detail ekstensi di katalog.

**Team-managed vs company-managed:** Tidak ada perbedaan; display conditions dapat menyaring berdasarkan tipe proyek.

**Catatan & ketergantungan:**

- Kerangka ekstensibilitas dirinci di [modul 25](25-integrasi-impor-ekspor-ekstensi.md). App Marketplace pihak ketiga berada di luar cakupan; produk ini hanya menyediakan extension point.

## Halaman personal, akses cepat & pencarian

### NAV-010 For you page (Your work)

**Deskripsi:** Halaman personal tempat pengguna mendarat setelah login, menampilkan pekerjaan yang relevan bagi dirinya: worked on, viewed, assigned, starred, dan boards. Di navigasi baru disebut "For you", di navigasi lama "Your work".

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi; setiap daftar disaring dengan permission Browse projects dan issue security ([modul 03](03-izin-peran-keamanan.md)).

**Aturan perilaku (acceptance criteria):**

1. Bagian atas berisi carousel "Recent spaces" berupa kartu (nama, avatar, tipe, jumlah work item open/done, quick links ke board dan work items) serta tautan "View all spaces".
2. Tab "Worked on" menampilkan item yang pengguna buat, edit, komentari, atau transisikan, dikelompokkan per waktu (Today, Yesterday, In the last week, In the last month) beserta aksi terakhir pengguna pada item tersebut.
3. Tab "Viewed" menampilkan item yang baru dilihat (work items, proyek, boards, filters, dashboards, plans), dikelompokkan per waktu yang sama.
4. Tab "Assigned to me" menampilkan work item yang di-assign ke pengguna dan belum berstatus kategori Done, dikelompokkan per status category (In progress, To do), dengan tautan "View all" ke navigator berisi JQL `assignee = currentUser() AND statusCategory != Done`.
5. Tab "Starred" menampilkan semua item yang di-star (proyek, boards, filters, dashboards, plans, queues) (NAV-012).
6. Tab "Boards" menampilkan board yang baru dikunjungi atau terkait pengguna (board pada proyek tempat ia aktif).
7. Setiap baris menampilkan ikon work type, key, summary, nama proyek, dan status atau waktu aktivitas relatif (NAV-033).
8. ⚠ Setiap tab memuat 20 item pertama dengan tombol "Load more"; tersedia filter teks lokal untuk menyaring item yang sudah dimuat.
9. ⚠ Item dari aplikasi lain dalam produk ini (mis. halaman knowledge base atau ideas) dapat tampil di tab Viewed/Worked on dengan ikon aplikasinya.
10. Setiap tab yang kosong menampilkan empty state berilustrasi dengan CTA yang sesuai (Create work item, Browse spaces) (NAV-045).
11. ⚠ Riwayat worked on/viewed disimpan dengan retensi 90 hari; item yang lebih tua tidak ditampilkan.
12. Hanya item yang masih dapat diakses pengguna saat ini yang ditampilkan (permission-trimmed saat query, bukan saat pencatatan).
13. Halaman tersedia di `/jira/for-you` dengan alias `/jira/your-work` yang me-redirect; tab aktif ter-encode di URL (mis. `?tab=assigned`) dan tab terakhir diingat per pengguna (NAV-037).
14. ⚠ Tab "Recommended" (rilis April 2026) menampilkan item yang butuh tindakan pengguna: task overdue yang di-assign kepadanya, approval yang menunggu keputusannya (BIZ-044 di [modul 26](26-kerja-bisnis-formulir-persetujuan.md), JSM-ITSM-011 di [modul 32](32-jsm-itsm-insiden-problem-perubahan-layanan.md)), pull request terbuka yang terkait work item miliknya (DEV-013 di [modul 24](24-integrasi-devops.md)), dan komentar yang menunggu balasan (mention terhadap pengguna yang belum dibalas). Setiap item disaring permission saat query dan memiliki aksi Dismiss; item yang di-dismiss tidak tampil lagi untuk pengguna tersebut.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Sumber data: log aktivitas pengguna (event domain di [modul 28](28-model-data-inti.md)) dan recent items (NAV-011).
- Produk ini tidak menyediakan pilihan homepage seperti "My Jira Home" di Data Center (NAV-056); For you selalu menjadi landing default.
- Tab Recommended (butir 14) bergantung pada sumber approval (BIZ-044, JSM-ITSM-011) dan development info (DEV-013); tiap kategori hanya tampil bila modul sumbernya aktif. Status ⚠ karena cakupan dan urutan kategori pada rilis referensi masih dapat berubah.

### NAV-011 Recent Items

**Deskripsi:** Daftar item yang baru diakses pengguna, dipakai di sidebar, quick search, dan For you untuk kembali cepat ke pekerjaan terakhir.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi; hanya pemilik yang dapat melihat daftar recent-nya sendiri.

**Aturan perilaku (acceptance criteria):**

1. Sistem mencatat kunjungan ke work items, proyek, boards, backlogs, filters, dashboards, plans, dan queues (Service Management) dengan timestamp akses terakhir per pengguna.
2. Flyout Recent di sidebar menampilkan item yang dikelompokkan per tipe beserta tautan "View all" ke tab Viewed di For you.
3. Saat query quick search kosong, dropdown menampilkan recent items (NAV-014).
4. ⚠ Daftar diurutkan berdasarkan waktu akses terakhir dan dibatasi 20 item per tipe di flyout serta 200 item per tipe di penyimpanan; item tertua dibuang saat batas terlampaui.
5. Proyek terbaru (recent) ditampilkan di section Spaces (daftar proyek) sidebar dan di bagian atas project picker (mis. dialog Create, NAV-013).
6. Item yang dihapus, diarsipkan, atau tidak dapat diakses lagi dikeluarkan otomatis dari daftar saat ditampilkan.
7. Daftar tersimpan server-side sehingga tersinkron lintas perangkat.
8. Tersedia REST API recent projects (setara `GET /rest/api/3/project/recent`) yang mengembalikan proyek terbaru pengguna pemanggil, dengan parameter `expand` dan batas maksimum 20.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- API dirinci di [modul 23](23-api-publik-webhook.md); event kunjungan tidak memicu notifikasi atau audit log.
- Satu sumber data dengan SRC-011 Recently viewed items ([modul 13](13-pencarian-filter-ekspor.md)): fitur ini memiliki pencatatan, retensi, sinkronisasi lintas perangkat, dan REST API recent projects; SRC-011 hanya membaca penyimpanan yang sama untuk pencarian dan JQL.

### NAV-012 Starred Items (Favorites)

**Deskripsi:** Menandai item penting dengan bintang agar mudah dibuka dari sidebar dan For you. Berlaku untuk proyek, boards, filters, dashboards, dan plans.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi pada item yang dapat ia lihat.

**Aturan perilaku (acceptance criteria):**

1. Ikon star tersedia di header proyek, board, filter, dashboard, dan plan, serta sebagai kolom star di setiap halaman daftar "View all".
2. Klik ikon men-toggle star/unstar secara optimistik dan menampilkan flag konfirmasi ("Added to starred"/"Removed from starred").
3. Item starred tampil di section Starred sidebar dan tab Starred di For you.
4. Proyek ber-star tampil di atas proyek terbaru pada section Spaces (daftar proyek) sidebar.
5. Starred filters setara "favourite filters": setiap filter menyimpan favourite count (jumlah pengguna yang menandainya) yang ditampilkan di halaman View all filters ([modul 13](13-pencarian-filter-ekspor.md)).
6. Item yang tidak lagi dapat diakses pengguna disembunyikan dari daftar starred tetapi star-nya tetap tersimpan dan muncul kembali bila akses dipulihkan.
7. ⚠ Pengguna dapat mengurutkan manual starred items per tipe (drag-and-drop atau menu Move up/Move down); urutan default adalah waktu star terbaru.
8. ⚠ Jumlah starred dibatasi 100 item per tipe; percobaan menambah melebihi batas ditolak dengan pesan galat yang menyebut batasnya.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Favourite count filter dipakai pada pencarian filter populer di [13](13-pencarian-filter-ekspor.md).

### NAV-013 Global Create Dialog

**Deskripsi:** Dialog Create work item yang dibuka dari tombol "+ Create" global atau shortcut `c`. Field yang tampil mengikuti create screen proyek dan work type yang dipilih.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Pengguna dengan permission Create issues pada proyek terpilih; Modify reporter untuk mengubah Reporter; Assign issues untuk mengisi Assignee; Create attachments untuk lampiran; Link issues untuk tautan.

**Aturan perilaku (acceptance criteria):**

1. Dialog dibuka dari tombol Create di top bar, shortcut `c`, dan entry point kontekstual; proyek dan work type ter-prefill sesuai konteks (proyek/board saat ini, atau pilihan terakhir pengguna).
2. Picker proyek menampilkan proyek terbaru di atas dan dapat dicari; picker Work type hanya berisi tipe dari issue type scheme proyek terpilih ([modul 05](05-konfigurasi-work-item-skema.md)).
3. Field yang tampil mengikuti create screen: company-managed memakai screen scheme, team-managed memakai layout field per work type; mengganti proyek/work type memuat ulang field dengan mempertahankan nilai field yang sama.
4. Field wajib ditandai "*" dan divalidasi sebelum submit; galat tampil inline di bawah field terkait dan fokus berpindah ke field galat pertama.
5. ⚠ Dialog menyediakan opsi "Configure fields" untuk memilih field opsional yang ditampilkan (pilihan "Semua field" atau "Field pilihan saya", disimpan per pengguna per proyek dan work type; field wajib selalu tampil), di proyek company-managed maupun team-managed; tautan ke konfigurasi screen/layout hanya tampil bagi Admin Proyek/Admin Situs (sama dengan ISS-002 butir 8).
6. Summary wajib, maksimum 255 karakter (galat "Summary must be less than 255 characters"); Description memakai rich editor ([modul 07](07-editor-konten-kaya.md)).
7. Checkbox "Create another" membuat dialog tetap terbuka setelah submit, mengosongkan Summary/Description, dan mempertahankan proyek, work type, serta field lain.
8. Dialog dapat di-minimize menjadi bar draft di pojok bawah layar dan di-restore dengan seluruh nilai yang sudah diisi; draft bertahan selama navigasi di dalam SPA, tetapi tidak dijamin bertahan saat reload atau menutup tab (sama dengan ISS-002 butir 5).
9. Dialog dapat di-expand ke mode full screen dan dikembalikan ke ukuran modal.
10. ⚠ Work item yang dibuat selalu ditempatkan di initial status workflow ([modul 09](09-workflow-status-transisi.md)); picker status awal tidak disediakan pada Fase 1 dan menjadi keputusan terbuka (sama dengan ISS-002 butir 4).
11. Dialog mendukung upload attachment, menambah link work item, dan memilih parent sebelum submit.
12. Setelah berhasil tampil flag "KEY-123 has been created" dengan aksi "View" dan "Copy link".
13. Menutup dialog yang berisi perubahan belum tersimpan memunculkan konfirmasi discard; Esc tanpa perubahan langsung menutup.
14. Field Reporter hanya dapat diubah bila pengguna memiliki Modify reporter (default pengguna saat ini) dan Assignee hanya dapat diisi dengan Assign issues; tanpa izin, field dirender read-only.
15. Create subtask/child dari konteks parent membuka dialog dengan parent dan proyek ter-prefill dan work type terbatas pada tipe child yang valid.
16. ⚠ Bila asisten AI aktif, tersedia aksi untuk menyusun summary/description atau usulan work breakdown dari teks bebas; hasil selalu ditinjau pengguna sebelum disimpan.

**Team-managed vs company-managed:** Company-managed: field dari create screen pada screen scheme. Team-managed: field dari layout per work type yang diatur admin proyek. Opsi Configure fields tersedia di keduanya.

**Catatan & ketergantungan:**

- Fitur ini sama dengan ISS-002 Global Create Dialog (Create work item) di [modul 06](06-work-item-inti.md); ISS-002 adalah pemilik perilaku (field, validasi, initial status, Configure fields, minimize/draft, konfirmasi discard). Bila ada perbedaan aturan, ISS-002 yang berlaku; modul ini mencatat entry point global (tombol Create di top bar, shortcut `c`) dan perilaku shell dialog. Validasi create di ISS-003, konfigurasi field di [05](05-konfigurasi-work-item-skema.md).

### NAV-014 Quick Search (top bar)

**Deskripsi:** Search box global dengan dropdown hasil instan untuk work items, proyek, boards, filters, people, dan konten lain dalam produk ini.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi; hasil disaring permission dan issue security.

**Aturan perilaku (acceptance criteria):**

1. Search box mendapat fokus lewat shortcut `/` atau klik; saat query kosong dropdown menampilkan recent items (NAV-011).
2. Hasil instan muncul saat mengetik dengan debounce 250 ms dan minimal 1 karakter, dikelompokkan: Work items, Spaces, Boards, Filters, Dashboards, Plans, People, dan konten aplikasi lain; maksimum 5 hasil per kelompok, dengan urutan grup mengikuti SRC-009.
3. Mengetik key lengkap (mis. ABC-123, tidak peka huruf besar/kecil) lalu Enter langsung membuka work item tersebut bila ada dan dapat diakses.
4. ⚠ Dropdown menyediakan filter cepat per aplikasi, proyek, dan assignee yang menyaring hasil instan.
5. "View all results" membuka halaman pencarian lengkap (NAV-015) atau Search work items dengan JQL `text ~ "<query>"`.
6. Tautan "Advanced search"/"Search work items" membuka issue navigator JQL ([modul 13](13-pencarian-filter-ekspor.md)).
7. Navigasi keyboard: Arrow up/down berpindah hasil, Enter membuka, Esc menutup dropdown, Cmd/Ctrl+Enter membuka hasil di tab baru.
8. Term yang cocok di-highlight pada hasil dan semua hasil dibatasi permission pengguna.
9. Hasil people membuka halaman profil (NAV-027) dan hasil board membuka board.
10. Pencarian teks mencocokkan summary, description, dan comment work item.
11. Bila tidak ada hasil ditampilkan state "No results" beserta saran (periksa ejaan, coba Advanced search).
12. ⚠ Dropdown menampilkan hingga 5 recent searches pengguna yang dapat dihapus satu per satu.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Indeks teks dan JQL di [12](12-jql-bahasa-kueri.md) dan [13](13-pencarian-filter-ekspor.md); sintaks smart querying lama tidak didukung (NAV-055).
- Pembagian kepemilikan dengan SRC-009 Quick Search (global search bar) di [modul 13](13-pencarian-filter-ekspor.md): SRC-009 memiliki logika hasil (sumber data, urutan grup, perilaku Enter, permission, ranking) dan SRC-010 memiliki pengenalan key (butir 3); fitur ini memiliki shell UI (posisi search box, shortcut, highlight, filter cepat, recent searches, tampilan "No results"). Bila ada perbedaan aturan logika hasil, SRC-009 yang berlaku.

### NAV-015 Rovo Search (cross-product search page)

**Deskripsi:** Halaman pencarian lengkap lintas aplikasi produk ini (work items, halaman knowledge base, ideas, people, teams) dengan filter dan jawaban AI opsional; setara Rovo Search.

**Prioritas / Fase / Tier:** P2 / Fase 4 / platform

**Aktor & izin:** Pengguna terlisensi; Admin Situs mengonfigurasi konektor dan fitur AI. Hasil selalu dibatasi permission sumber.

**Aturan perilaku (acceptance criteria):**

1. Halaman hasil full-page menyediakan filter tipe: All, Work items, Pages, People, Teams, dan tipe lain yang tersedia.
2. Filter tambahan: aplikasi, proyek, contributor, last updated, dan type; filter ter-encode di URL.
3. Bila asisten AI aktif dan query berbentuk pertanyaan, jawaban AI ringkas dengan sitasi ditampilkan di atas hasil.
4. Konektor ke sumber eksternal (mis. Google Drive, Slack, GitHub, SharePoint) dikonfigurasi admin; hasil dari konektor dibatasi permission di sumber aslinya.
5. Hasil People & Teams menampilkan hover card (NAV-028).
6. Hasil dapat diurutkan berdasarkan relevance atau recency.
7. ⚠ Admin dapat mengelola bookmarks (hasil yang dipromosikan untuk kata kunci tertentu) dan definitions (penjelasan istilah internal).
8. ⚠ Halaman pencarian tersedia untuk edisi berbayar; edisi Free mendapat versi terbatas tanpa jawaban AI dan tanpa konektor.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Tahap awal cukup pencarian lintas entitas internal; jawaban AI dan konektor opsional (NAV-051). Konektor first-party ke layanan eksternal dirinci di [modul 25](25-integrasi-impor-ekspor-ekstensi.md).
- Tumpang tindih dengan SRC-012 Search results page (full-page search) dan SRC-014 Rovo Search (cross-product search) di [modul 13](13-pencarian-filter-ekspor.md): SRC-012 memiliki halaman hasil full-page untuk work items, SRC-014 memiliki logika pencarian lintas produk dan AI; fitur ini memiliki entry point navigasi dan tampilan halaman di shell. Prioritas SRC-014 (P3) dan fitur ini (P2) masih berbeda dan perlu diselaraskan pemilik produk (Pertanyaan terbuka nomor 11).

### NAV-016 Share Dialog & Copy Link

**Deskripsi:** Pola UI global untuk membagikan item (work item, board, filter, dashboard, plan) ke pengguna/tim/email atau menyalin URL kanonik.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna yang dapat melihat item; penerima tetap tunduk pada permission item.

**Aturan perilaku (acceptance criteria):**

1. Tombol Share di header item membuka dialog berisi picker pengguna/tim/alamat email dan pesan opsional (maksimum 1.000 karakter).
2. Penerima mendapat email berisi tautan dan pesan; bila penerima tidak memiliki akses ke item, dialog menampilkan peringatan sebelum dikirim.
3. "Copy link" menyalin URL kanonik item (mis. `/browse/KEY-123`) ke clipboard dan menampilkan flag "Link copied".
4. ⚠ Bila integrasi chat first-party (Slack/Microsoft Teams) terhubung, dialog menyediakan opsi share ke channel tersebut.
5. ⚠ Share via email dibatasi rate limit per pengguna (usulan 50 penerima per jam); melebihi batas menampilkan galat "You've shared too many times. Try again later."

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Pengiriman email memakai infrastruktur [modul 19](19-notifikasi-email.md); integrasi chat di [modul 25](25-integrasi-impor-ekspor-ekstensi.md).

### NAV-017 Smart Links & Link Previews (platform)

**Deskripsi:** Rendering URL menjadi tautan kaya (inline/card/embed) beserta hover preview, termasuk unfurl tautan produk ini di luar aplikasi.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Semua pengguna; preview hanya menampilkan data yang boleh dilihat pengguna yang melihat.

**Aturan perilaku (acceptance criteria):**

1. Menempelkan URL internal atau URL pihak ketiga yang didukung di editor otomatis mengubahnya menjadi smart link inline (ikon, judul, status lozenge).
2. Tampilan yang tersedia: URL, Inline, Card, dan Embed; tampilan default mengikuti "Link preferences" pengguna (NAV-029).
3. ⚠ Hover pada smart link menampilkan preview (judul, status, assignee) dengan aksi cepat seperti buka di tab baru dan salin tautan.
4. Bila pengguna tidak berhak, smart link menampilkan "Request access" atau "Connect account" tanpa membocorkan judul/isi item.
5. Tautan produk ini dapat di-unfurl di Slack, Microsoft Teams, dan aplikasi wiki melalui integrasi terkait.
6. Untuk sumber pihak ketiga, pengguna menghubungkan akun lewat OAuth sebelum preview dapat diambil.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Pembagian kepemilikan: EDT-021 Smart Links (Inline, Card, Embed) di [07](07-editor-konten-kaya.md) memiliki rendering di editor (mode tampilan, toolbar, paste); INT-002 Smart Links & link previews (third-party accounts) di [modul 25](25-integrasi-impor-ekspor-ekstensi.md) memiliki registry resolver, koneksi akun pihak ketiga (OAuth, butir 6 di sini), dan unfurl di aplikasi chat; NAV-017 hanya memiliki hover preview tingkat platform, default tampilan dari "Link preferences" (NAV-029), dan state tanpa akses di luar editor (mis. hover card di daftar, notifikasi in-app, dashboard).
- Prioritas disamakan dengan EDT-021 dan INT-002 (P1 / Fase 3) karena ketiganya dirilis sebagai satu kemampuan; hover preview tidak berguna tanpa resolver, dan resolver tidak terlihat tanpa rendering.

## Notifikasi, bantuan, onboarding & edisi

### NAV-018 Notifications Hub (in-app notifications panel)

**Deskripsi:** Panel notifikasi in-app dari ikon bell yang menampilkan aktivitas relevan lintas aplikasi produk ini, dengan pengelompokan, status read/unread, dan aksi cepat.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna terautentikasi untuk notifikasinya sendiri; isi notifikasi disaring permission saat ditampilkan.

**Aturan perilaku (acceptance criteria):**

1. Ikon bell menampilkan badge jumlah unread yang ter-update real-time tanpa reload (NAV-046).
2. Klik bell membuka panel/drawer dengan tab "Direct" (mention, assign, reply yang ditujukan ke pengguna) dan "Watching" (aktivitas pada item yang di-watch).
3. Toggle "Only show unread" menyaring daftar dan tersimpan per pengguna.
4. Tersedia "Mark all as read" untuk tab aktif serta mark read/unread per notifikasi.
5. Beberapa event pada work item yang sama dalam rentang waktu berdekatan dikelompokkan menjadi satu entri yang dapat di-expand untuk melihat detail (isi comment, perubahan field).
6. ⚠ Panel memuat notifikasi dari semua aplikasi produk ini (pelacak kerja, Service Management, Product Discovery) dengan filter per aplikasi.
7. Klik notifikasi membuka item dan menandainya read; Cmd/Ctrl+klik membuka item di tab baru dan tetap menandainya read.
8. ⚠ Notifikasi comment dan mention mendukung aksi inline reply dan reaction langsung dari panel tanpa membuka work item.
9. Ikon cog di header panel menuju "Notification settings" personal (NAV-019).
10. ⚠ Daftar dimuat dengan infinite scroll (30 entri per halaman) dan notifikasi disimpan dengan retensi 30 hari.
11. Bila tidak ada notifikasi (atau tidak ada unread saat filter aktif) ditampilkan empty state "You're all caught up".
12. Konten notifikasi disembunyikan dan diganti teks netral bila pengguna kehilangan akses ke item terkait.
13. Event yang menghasilkan notifikasi in-app minimal: mentioned, assigned, comment/reply, perubahan status/field pada item yang di-watch, dan approval request (Service Management).
14. ⚠ Status read tersinkron dengan push notification mobile: membaca di satu kanal menandai read di kanal lain.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Pembangkitan event, notification scheme, dan email dirinci di [19 Notifikasi In-app & Email](19-notifikasi-email.md); modul ini hanya UI hub in-app.

### NAV-019 Personal Notification Settings

**Deskripsi:** Halaman preferensi notifikasi per pengguna yang mengontrol kanal email, in-app, dan mobile per kategori event.

**Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi untuk preferensinya sendiri.

**Aturan perilaku (acceptance criteria):**

1. ⚠ Halaman `/jira/settings/personal/notifications` menyediakan pengaturan per kanal: Email, Slack/Microsoft Teams (bila integrasi first-party terhubung), dan mobile push (Web Push untuk PWA terpasang, NAV-050).
2. Toggle per peran/kategori: You're watching, You're assigned, You're the reporter, You're mentioned, plus per event (work item created, updated, commented, transitioned, dll.).
3. Toggle "You make changes" menentukan apakah pengguna menerima notifikasi atas perubahan yang ia buat sendiri; default off.
4. ⚠ Email batching menggabungkan beberapa perubahan pada work item yang sama dalam satu email dengan jendela waktu 10 menit.
5. Bahasa email mengikuti preferensi bahasa pengguna penerima (NAV-033).
6. Pengaturan notifikasi disimpan terpisah per aplikasi (pelacak kerja, Service Management, Product Discovery).
7. ⚠ Push mobile (Web Push pada PWA, NAV-050) diatur per tipe notifikasi di halaman ini, termasuk Do not disturb/quiet hours.
8. Setiap email notifikasi memuat tautan unsubscribe/manage yang menuju pengaturan terkait tanpa perlu menavigasi manual.
9. Notification scheme proyek tetap menjadi batas atas: pengguna tidak dapat menerima notifikasi yang tidak dikirim scheme, dan halaman menampilkan keterangan ini.
10. Pilihan format email HTML vs Text tidak disediakan (hanya ada di Data Center); email selalu dikirim multipart HTML + teks.

**Team-managed vs company-managed:** Team-managed mengatur notifikasi per proyek lewat pengaturan proyek; company-managed lewat notification scheme. Preferensi personal berlaku sama di keduanya.

**Catatan & ketergantungan:**

- Notification scheme dan email server di [modul 19](19-notifikasi-email.md). Integrasi Slack/Teams di [modul 25](25-integrasi-impor-ekspor-ekstensi.md).
- Tumpang tindih dengan NTF-020 Personal Email Notification Settings ([modul 19](19-notifikasi-email.md)): NTF-020 memiliki aturan email (master toggle, toggle per peran, format); fitur ini memiliki halaman Personal settings dan kanal lain. Prioritas disamakan dengan NTF-020 (P0 / Fase 2).

### NAV-020 Help Menu & In-product Help Panel

**Deskripsi:** Menu bantuan "?" yang membuka panel help kontekstual, daftar shortcut, what's new, feedback, dan sumber belajar.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna; opsi dukungan menyesuaikan edisi workspace.

**Aturan perilaku (acceptance criteria):**

1. Panel help terbuka di sisi kanan tanpa meninggalkan halaman dan menyediakan pencarian artikel dokumentasi.
2. Panel menampilkan artikel rekomendasi kontekstual sesuai halaman aktif (board, backlog, settings, dll.).
3. Tersedia tautan "Browse help", komunitas pengguna, dan materi pelatihan (dapat dikonfigurasi URL-nya).
4. "Keyboard shortcuts" membuka dialog shortcut yang sama dengan tombol `?` (NAV-038).
5. "What's new" menampilkan changelog fitur terbaru (NAV-021).
6. "Give feedback" membuka dialog feedback (NAV-022).
7. ⚠ "Get mobile app" menampilkan QR code dan petunjuk memasang PWA (Add to Home Screen); tautan app store hanya tampil bila app native kelak disetujui (NAV-050).
8. "Contact support"/"Ask the community" menyesuaikan edisi: Free hanya komunitas, Standard dukungan jam kerja, Premium/Enterprise dukungan 24/7 untuk work item support berprioritas critical.
9. Tautan "Terms of use", "Privacy policy", dan "Notice at collection" selalu tersedia.
10. ⚠ Admin Situs dapat menambahkan tautan help kustom (mis. wiki internal) yang tampil di menu help.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Konten dokumentasi dikelola di luar aplikasi; URL dan tautan kustom dikonfigurasi di [modul 22](22-administrasi-sistem.md).

### NAV-021 What's New, Changeboarding & Feature Opt-ins

**Deskripsi:** Mekanisme memperkenalkan dan mencoba fitur baru: daftar what's new, lozenge "New", spotlight, banner opt-in/opt-out, dan fitur eksperimental (labs).

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Semua pengguna; Admin Situs mengatur release track dan feature flag tingkat workspace.

**Aturan perilaku (acceptance criteria):**

1. Daftar "What's new" di help panel diurutkan per tanggal rilis terbaru dan menandai entri yang belum dibaca.
2. Lozenge "NEW" atau pulsing dot ditampilkan pada item navigasi/tombol fitur baru dan hilang setelah pengguna membukanya atau setelah 30 hari.
3. Spotlight/coachmark berisi judul, deskripsi, tombol "Next" dan "Dismiss"; setiap spotlight tampil sekali per pengguna.
4. Banner "Try the new X" dengan toggle opt-in/opt-out tersedia selama periode rollout fitur (mis. tampilan work item baru).
5. Saat pengguna opt-out, ditampilkan prompt feedback opsional berisi alasan.
6. ⚠ Personal settings menyediakan toggle "Labs" untuk mengaktifkan fitur eksperimental per pengguna.
7. Status dismissed (spotlight, banner, lozenge) disimpan server-side per pengguna.
8. Admin edisi Premium/Enterprise dapat memilih release track (Continuous, Bundled, Preview via sandbox) untuk menunda perubahan fitur ([modul 22](22-administrasi-sistem.md)).
9. Feature flag server-side dapat menargetkan per pengguna, per grup, atau per workspace, dan dievaluasi di backend sebelum UI dirender.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Mekanisme feature flag juga dipakai untuk gating edisi (NAV-024) dan rollout bertahap ([modul 27](27-nfr-arsitektur-platform.md)).

### NAV-022 Give Feedback Dialog

**Deskripsi:** Dialog untuk mengirim feedback produk ke tim pengembang produk ini dari berbagai titik UI.

**Prioritas / Fase / Tier:** P3 / Fase 3 / platform

**Aktor & izin:** Semua pengguna terautentikasi; admin dapat menonaktifkan pengiriman feedback.

**Aturan perilaku (acceptance criteria):**

1. Dialog menyediakan pilihan tipe: "Ask a question", "Leave a comment", "Report a bug", "Suggest an improvement"; tipe wajib dipilih.
2. Tersedia textarea feedback (wajib, maksimum 2.000 karakter) serta checkbox "You can contact me about this feedback" dan "I'd like to participate in product research" (default off).
3. Konteks halaman/URL dan metadata browser (user agent, ukuran viewport, versi aplikasi) dilampirkan otomatis dan ditampilkan ke pengguna sebelum kirim.
4. Dialog dapat dibuka dari help menu dan dari entry point fitur tertentu (mis. "Give feedback" di view baru), dengan konteks fitur ikut terlampir.
5. Setelah berhasil terkirim ditampilkan flag "Thanks for your feedback".

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Tujuan pengiriman feedback (endpoint internal atau pelacak publik) adalah keputusan terbuka untuk instalasi self-hosted.

### NAV-023 Onboarding, Tours & Quickstart

**Deskripsi:** Pengalaman pengguna baru: sign-up, kuesioner, rekomendasi template, undangan tim, tur produk, checklist Quickstart, dan konten onboarding kustom yang dikelola admin.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Pengguna baru; Admin Situs dan Jira admin; pengguna dengan global permission "Manage custom onboarding".

**Aturan perilaku (acceptance criteria):**

1. Sign-up membuat workspace (subdomain unik, divalidasi format dan ketersediaannya), memilih aplikasi, dan membuat akun atau login dengan Google/Microsoft/Apple/Slack ([modul 01](01-identitas-autentikasi-sesi.md)).
2. Kuesioner opsional menanyakan jenis pekerjaan/tim (software, marketing, HR, dll.), pengalaman dengan pelacak kerja, dan ukuran tim; jawaban dapat dilewati.
3. Berdasarkan jawaban ditampilkan rekomendasi template dan dibuat proyek pertama dengan nama dan key yang dapat diedit sebelum disimpan.
4. Langkah "Invite your team" menerima email atau approved domain dan menolak undangan yang melebihi batas pengguna edisi (NAV-024).
5. Checklist "Quickstart" di panel kanan menampilkan progres (create work item, invite teammates, set up board, connect tools) dan dapat disembunyikan atau dibuka kembali dari help menu.
6. Spotlight tour untuk board, backlog, dan timeline memiliki tombol Next/Back/Skip dan tampil sekali per pengguna.
7. ⚠ Template tertentu (mis. Scrum dan Kanban) dapat dibuat dengan sample data berisi contoh work items yang dapat dihapus sekaligus.
8. Contextual onboarding tampil saat pertama kali membuka fitur (mis. empty state "Start your first sprint").
9. Pengguna yang diundang melihat welcome screen dan pilihan peran saat login pertama.
10. Status onboarding (langkah selesai, tour dilihat, checklist disembunyikan) disimpan per pengguna server-side.
11. ⚠ Admin atau pengguna dengan "Manage custom onboarding" dapat menyusun konten onboarding khusus organisasi (checklist/panduan) untuk pengguna baru.
12. ⚠ Halaman pengelolaan custom onboarding di settings memungkinkan membuat, mengedit, mengaktifkan, dan menonaktifkan konten onboarding.
13. ⚠ Konten custom onboarding dapat ditargetkan ke pengguna baru atau grup tertentu dan dipratinjau sebelum dipublikasikan.
14. ⚠ Custom onboarding melengkapi onboarding bawaan secara default, dengan opsi admin untuk menggantikan checklist bawaan.
15. Akses ke pengelolaan custom onboarding dibatasi global permission "Manage custom onboarding" selain Jira admin; pengguna lain mendapat 403.

**Team-managed vs company-managed:** Tidak ada perbedaan; template proyek pertama dapat team-managed atau company-managed sesuai pilihan.

**Catatan & ketergantungan:**

- Global permission didefinisikan di [03](03-izin-peran-keamanan.md); template proyek di [04](04-proyek.md); pendaftaran akun di [01](01-identitas-autentikasi-sesi.md).

### NAV-024 Plan Editions, User Limits & Upgrade Prompts

**Deskripsi:** Batasan per edisi (plan) yang terlihat di UX (jumlah pengguna, storage, ketersediaan fitur) serta mekanisme prompt upgrade/trial.

**Prioritas / Fase / Tier:** P1 / Fase 3 / platform

**Aktor & izin:** Admin Situs dan Billing admin mengelola edisi; semua pengguna melihat prompt dan lozenge fitur terkunci.

**Aturan perilaku (acceptance criteria):**

1. ⚠ Edisi Free: maksimal 10 pengguna, storage 2 GB, dukungan komunitas, tanpa user roles/permissions lanjutan, tanpa anonymous access, tanpa audit log.
2. ⚠ Edisi Standard: hingga 100.000 pengguna, storage 250 GB, user roles & permissions, data residency, dukungan jam kerja.
3. Edisi Premium: hingga 100.000 pengguna, storage tidak terbatas, Plans, sandbox, release tracks, target SLA uptime 99,9%, dukungan 24/7 untuk work item support berprioritas critical.
4. Edisi Enterprise: multi-workspace (hingga sekitar 150), target SLA uptime 99,95%, fitur keamanan organisasi lanjutan, analitik lintas workspace.
5. Undangan yang membuat jumlah pengguna melebihi batas edisi diblokir dengan pesan yang menyebut batas dan CTA upgrade.
6. Saat storage mencapai 80% ditampilkan banner peringatan untuk admin; saat 100% upload attachment diblokir dengan galat "Storage limit reached".
7. Fitur terkunci menampilkan lozenge edisi ("PREMIUM"/"STANDARD") dan klik membuka dialog "Upgrade" atau "Try it free".
8. Trial Premium (30 hari) dapat dimulai dari prompt in-product oleh admin, dengan banner sisa hari yang tampil untuk admin.
9. Non-admin yang mengklik fitur terkunci melihat "Ask your admin" dan dapat mengirim permintaan yang diteruskan ke admin.
10. ⚠ Saat downgrade, data fitur edisi lebih tinggi dipertahankan tetapi fiturnya menjadi read-only atau tersembunyi sesuai fitur, dan admin melihat ringkasan dampak sebelum konfirmasi.
11. Lisensi agent Service Management dan creator Product Discovery dihitung terpisah dari pengguna pelacak kerja ([modul 29](29-jsm-service-project-request-type-formulir.md) dan 37).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Dimodelkan sebagai feature flag per edisi workspace (NAV-021); angka batas adalah konfigurasi, bukan konstanta kode. Untuk instalasi self-hosted gating edisi dapat dimatikan seluruhnya (keputusan terbuka).

## Menu Settings, profil, akun & sesi

### NAV-025 Settings Menu (cog)

**Deskripsi:** Menu Settings di top bar yang mengelompokkan pengaturan personal dan administrasi sesuai permission pengguna.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna (bagian personal); Jira admin (global permission Administer Jira) untuk bagian administrasi produk; Admin Situs/organisasi untuk administrasi pengguna dan billing.

**Aturan perilaku (acceptance criteria):**

1. Bagian personal berisi "Personal settings" (NAV-030) dan "Account settings" (NAV-029).
2. Bagian administrasi produk (khusus Jira admin) berisi System, Apps, Spaces, Work items, dan Products.
3. Bagian administrasi organisasi (Admin Situs/organisasi) berisi User management dan Billing yang membuka konsol admin.
4. Item tanpa permission tidak dirender; non-admin hanya melihat bagian personal.
5. ⚠ Admin dapat membuka admin search (shortcut `g` lalu `g`) untuk mencari halaman administrasi berdasarkan nama.
6. Halaman administrasi memakai sidebar khusus admin dengan tombol "Back to app" yang kembali ke halaman terakhir di luar area admin.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Isi halaman administrasi di [modul 22](22-administrasi-sistem.md) dan [02](02-pengguna-grup-organisasi.md); route admin `frontend/src/lib/features/admin` yang sudah ada menjadi titik awal.

### NAV-026 Profile Menu (avatar dropdown)

**Deskripsi:** Dropdown dari avatar di top bar untuk akses profil, pengaturan akun, theme, ganti akun, dan logout.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna terautentikasi.

**Aturan perilaku (acceptance criteria):**

1. Header dropdown berisi avatar, nama tampilan, dan email pengguna.
2. "Profile" membuka halaman profil pengguna sendiri (NAV-027).
3. "Account settings" membuka pengaturan akun tingkat platform (NAV-029).
4. "Personal settings" membuka pengaturan personal aplikasi (NAV-030).
5. ⚠ Submenu "Theme" berisi Light, Dark, Match browser, dan opsi increased contrast bila tersedia; pilihan langsung diterapkan tanpa reload (NAV-034).
6. "Switch account" membuka daftar akun yang login di browser (NAV-008).
7. "Log out" mengakhiri sesi di browser tersebut (NAV-032).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Data header dari endpoint profil pengguna saat ini ([modul 01](01-identitas-autentikasi-sesi.md)).

### NAV-027 Profile Page (Atlassian account profile)

**Deskripsi:** Halaman profil pengguna yang menampilkan identitas, info kontak sesuai visibility, tim, tempat bekerja, dan aktivitas.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna yang memiliki global permission Browse users and groups; pemilik profil untuk mengedit; admin organisasi untuk akun terkelola.

**Aturan perilaku (acceptance criteria):**

1. Halaman menampilkan header image, avatar besar, nama, public name, job title, department, organization, based in, dan local time.
2. Email dan info kontak ditampilkan sesuai visibility masing-masing field (Anyone / Organization / Only you and admins) ([modul 01](01-identitas-autentikasi-sesi.md)).
3. Section "Worked on" menampilkan aktivitas terbaru pengguna hanya pada item yang juga dapat diakses viewer.
4. Section "Places you work in" menampilkan proyek tempat pengguna aktif.
5. Section Teams menampilkan tim tempat pengguna menjadi anggota dengan tautan ke team profile.
6. ⚠ Reporting lines (manager dan direct reports) ditampilkan bila datanya tersedia dari organisasi/IdP.
7. Tautan "Assigned work items" (JQL `assignee = <accountId>`) dan "Reported work items" membuka navigator.
8. Pemilik profil dapat mengedit inline nama, job title, department, dan field lain yang tidak dikelola IdP.
9. Field yang dikelola admin (managed account/SCIM) tampil terkunci dengan keterangan "Managed by your organization".
10. Profil pengguna yang dinonaktifkan/dihapus ditampilkan sebagai "Former user" atau "Account deactivated" tanpa data kontak.
11. URL profil `/jira/people/{accountId}` memakai accountId; URL berbasis username tidak didukung dan mengembalikan 404.
12. ⚠ Halaman menampilkan kudos dan aktivitas lintas aplikasi pengguna bila fitur tersebut tersedia.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Model identitas `accountId` di [01](01-identitas-autentikasi-sesi.md); tim di [modul 18](18-timeline-rencana-tampilan-proyek.md) dan [02](02-pengguna-grup-organisasi.md).

### NAV-028 Profile Hover Card

**Deskripsi:** Kartu ringkas yang muncul saat hover/focus pada avatar atau nama pengguna di mana pun dalam produk ini.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna; isi kartu mengikuti visibility profil.

**Aturan perilaku (acceptance criteria):**

1. Kartu muncul setelah hover 500 ms atau saat focus keyboard, dan tertutup saat mouse leave (dengan toleransi 300 ms) atau Esc.
2. Isi kartu: avatar, nama, job title, department, local time, dan email sesuai visibility.
3. ⚠ Aksi di kartu: "View profile" dan "Assigned work items"; aksi lain dapat ditambahkan per konteks.
4. Pengguna nonaktif ditampilkan sebagai "Account deactivated"/former user tanpa aksi selain View profile.
5. Data kartu dimuat lazy saat pertama dibutuhkan dan di-cache di sisi client selama 5 menit.
6. Team hover card menampilkan nama tim, anggota (maksimum 5 avatar + jumlah sisa), dan aksi "View team".
7. Kartu dapat diakses keyboard dan dibacakan screen reader (role dialog non-modal dengan label nama pengguna).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Local time memakai zona waktu profil (NAV-033).

### NAV-029 Atlassian Account Settings

**Deskripsi:** Pengaturan akun tingkat platform yang berlaku lintas aplikasi produk ini: profil, email, keamanan, preferensi, privasi, dan connected apps (setara id.atlassian.com).

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Pemilik akun; admin organisasi untuk akun terkelola.

**Aturan perilaku (acceptance criteria):**

1. Bagian Profile and visibility: avatar (upload dengan crop, atau inisial default), header image, full name, public name, job title, department, organization, based in, beserta visibility per field.
2. Bagian Email: perubahan email memerlukan verifikasi ke alamat baru sebelum berlaku; visibility email dapat diatur.
3. ⚠ Bagian Security: ubah password, two-step verification (authenticator app, recovery code, passkeys/security keys), daftar sesi aktif per perangkat dengan aksi logout, dan API tokens (create/revoke dengan expiry).
4. Bagian Account preferences: Language (atau auto-detect dari browser), Time zone, dan "Link preferences" (default tampilan smart link: inline/card/embed, NAV-017).
5. Bagian Connected apps menampilkan aplikasi OAuth yang diberi akses dan memungkinkan mencabutnya.
6. ⚠ Bagian Privacy: cookie preferences, ekspor data pribadi, dan hapus/deactivate akun dengan grace period.
7. Pada akun terkelola, field tertentu dikunci oleh admin organisasi/IdP (SAML/SCIM) dan tampil nonaktif.
8. Perubahan profil ter-propagate ke semua aplikasi dan workspace tempat akun menjadi anggota.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- NAV-029 hanya halaman UI self-service (tata letak, navigasi antar bagian, state terkunci). Perilaku akun dimiliki [01](01-identitas-autentikasi-sesi.md): model identitas AUTH-001 Atlassian Account (Identity Model), profil & visibility AUTH-002 (butir 1), perubahan email AUTH-004 (butir 2), password AUTH-005 dan 2SV AUTH-010 (butir 3), sesi & perangkat AUTH-009, API token AUTH-012, connected apps AUTH-013 (butir 5), dan penutupan akun AUTH-014 (butir 6). Bila ada perbedaan aturan, modul 01 yang berlaku.

### NAV-030 Personal Jira Settings

**Deskripsi:** Halaman pengaturan khusus aplikasi per pengguna yang mengatur notifikasi, autowatch, keyboard shortcuts, labs, theme, dan aksesibilitas.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna terautentikasi.

**Aturan perilaku (acceptance criteria):**

1. Halaman menampilkan tautan ke pengaturan language dan time zone yang dikelola di Account settings (NAV-029), beserta nilai saat ini.
2. Tersedia tautan/menu ke "Notification settings" personal (NAV-019).
3. Toggle Autowatch "Watch work items you create or comment on" (NAV-031).
4. Toggle Keyboard shortcuts untuk mengaktifkan/menonaktifkan semua shortcut (NAV-038).
5. ⚠ Toggle Labs untuk fitur eksperimental (NAV-021).
6. Pilihan Theme/Appearance (Light/Dark/Match browser) (NAV-034).
7. Bagian Accessibility settings (NAV-035).
8. Ekstensi terpasang dapat menambah tab lewat extension point personal settings page (NAV-009).
9. Route berada di bawah `/jira/settings/personal/...` (mis. `/general`, `/notifications`) dan setiap tab dapat dibuka langsung lewat URL.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Produk ini tidak menyediakan pilihan homepage, jumlah hasil per halaman, atau format email (khusus Data Center, NAV-056).

### NAV-031 Autowatch

**Deskripsi:** Preferensi agar pengguna otomatis menjadi watcher pada work item yang ia buat atau komentari.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna; watcher hanya dapat ditambahkan bila pengguna memiliki permission Browse projects pada work item.

**Aturan perilaku (acceptance criteria):**

1. Toggle on/off tersedia di Personal settings; nilai default on untuk pengguna baru.
2. ⚠ Autowatch dipicu saat pengguna membuat work item dan saat menambah comment; mengedit field tidak memicu autowatch.
3. Watcher tidak ditambahkan bila pengguna tidak memiliki permission view pada work item tersebut.
4. ⚠ Pengguna tetap dapat unwatch manual setelah di-autowatch, dan aksi comment berikutnya tidak menambahkannya kembali.
5. Reporter dan assignee tetap menerima notifikasi sesuai notification scheme terlepas dari status watch.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Fitur ini sama dengan NTF-014 Autowatch di [modul 19](19-notifikasi-email.md); NTF-014 memiliki aturan default (NTF-021) dan pengecualian import, fitur ini memiliki toggle di Personal settings. Status ⚠ butir 4 disamakan dengan NTF-014 butir 4.
- Fitur Watchers di [06](06-work-item-inti.md); notifikasi di [modul 19](19-notifikasi-email.md). Keputusan terbuka pada butir 4 (apakah unwatch manual mencegah autowatch berikutnya) diusulkan "ya"; lihat Pertanyaan terbuka nomor 10.

### NAV-032 Logout & Session UX

**Deskripsi:** Perilaku logout dan sesi dari sisi pengalaman pengguna.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna; admin organisasi menetapkan kebijakan sesi.

**Aturan perilaku (acceptance criteria):**

1. "Log out" dari profile menu mengakhiri sesi di browser tersebut (token dicabut di server, bukan hanya dihapus di client).
2. Setelah logout ditampilkan konfirmasi singkat lalu redirect ke halaman login; tombol Back browser tidak menampilkan data terproteksi.
3. Session idle timeout mengikuti kebijakan autentikasi yang ditetapkan admin ([modul 01](01-identitas-autentikasi-sesi.md)).
4. ⚠ Lima menit sebelum idle timeout ditampilkan dialog peringatan dengan tombol "Stay logged in".
5. ⚠ Draft yang belum tersimpan (dialog Create, comment) disimpan lokal dan dipulihkan setelah login ulang pada perangkat yang sama.
6. Pengguna dapat melihat dan logout sesi aktif per perangkat dari Account settings (NAV-029).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Kondisi repo saat ini: JWT HS256 24 jam tanpa revocation; logout server-side bergantung pada model sesi yang dapat dicabut di [01](01-identitas-autentikasi-sesi.md).

## Bahasa, tampilan, aksesibilitas & persistensi UI

### NAV-033 Language, Region & Date/Time Display

**Deskripsi:** Preferensi bahasa UI dan zona waktu per pengguna beserta pola tampilan tanggal dan waktu yang konsisten di seluruh UI.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna untuk preferensinya sendiri; Admin Situs untuk default bahasa dan zona waktu workspace.

**Aturan perilaku (acceptance criteria):**

1. ⚠ Produk ini mendukung bahasa UI English dan Bahasa Indonesia pada Fase 1; arsitektur i18n (katalog pesan per locale) memungkinkan penambahan bahasa lain (target jangka panjang setara ~25 bahasa: Deutsch, Español, Français, 日本語, dll.) tanpa perubahan kode.
2. Pengguna dapat memilih "Auto-detect" sehingga bahasa mengikuti header `Accept-Language` browser dengan fallback ke default workspace.
3. Setiap pengguna memiliki time zone (IANA, mis. `Asia/Jakarta`); bila kosong dipakai default workspace yang diatur admin ([modul 22](22-administrasi-sistem.md)).
4. Admin mengatur default language workspace; preferensi pengguna selalu meng-override default tersebut.
5. Aktivitas terbaru ditampilkan sebagai waktu relatif ("just now", "5 minutes ago", "Yesterday") hingga 7 hari, dan sebagai tanggal absolut untuk yang lebih lama.
6. Hover/focus pada waktu relatif menampilkan tooltip tanggal/waktu lengkap sesuai zona waktu pengguna.
7. Due date dan field date-only ditampilkan tanpa jam dan tanpa konversi zona waktu (nilai disimpan sebagai tanggal kalender).
8. ⚠ Format tanggal/angka mengikuti locale pengguna (12h/24h, urutan hari-bulan, pemisah desimal); admin dapat mengatur format default workspace.
9. Due date yang lewat dan work item belum Done ditandai indikator overdue (warna dan ikon, tidak hanya warna).
10. Local time pengguna lain tampil di hover card dan halaman profil (NAV-027, NAV-028).
11. Email notifikasi dikirim dalam bahasa penerima, bukan bahasa pengirim aksi.
12. ⚠ Konten buatan pengguna (nama field, status, workflow) tidak diterjemahkan otomatis; admin dapat menyimpan terjemahan nama status dan work type per bahasa.
13. ⚠ Layout right-to-left (RTL) tidak didukung pada rilis awal.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Kondisi repo saat ini: `middleware.LanguageMiddleware()` mendeteksi EN/ID dari `Accept-Language` dan `internal/shared/constants/error.go` menyediakan pesan galat EN/ID; preferensi bahasa tersimpan di akun ([modul 01](01-identitas-autentikasi-sesi.md)) harus mengalahkan header bila ada. Kolom `locale` dan `timezone` belum ada di tabel `users` dan perlu migrasi.

### NAV-034 Theme & Appearance (light/dark/system)

**Deskripsi:** Pemilihan tema tampilan per pengguna dengan dukungan dark mode berbasis design tokens.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna.

**Aturan perilaku (acceptance criteria):**

1. Pilihan theme: Light, Dark, dan Match browser; Match browser mengikuti `prefers-color-scheme` dan berubah otomatis saat preferensi OS berubah tanpa reload.
2. Theme dapat diubah dari profile menu (NAV-026) dan Personal settings (NAV-030).
3. Preferensi tersimpan per pengguna lintas perangkat; bila belum diset dipakai Match browser.
4. ⚠ Tersedia theme increased contrast sebagai opsi tambahan untuk Light dan Dark.
5. Seluruh komponen (kartu board, status lozenge, chart, editor) memakai design tokens sehingga tidak ada warna hard-coded yang tidak berganti antar theme.
6. Konten kustom (warna epic/label, avatar proyek, logo workspace) tetap terbaca di dark mode dengan penyesuaian kontras otomatis atau latar netral.
7. Email notifikasi selalu memakai tampilan light terlepas dari theme pengguna.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Theme diterapkan sebelum render pertama (atribut pada `<html>`) untuk mencegah kilatan theme salah; tokens diteruskan ke ekstensi (NAV-009).

### NAV-035 Accessibility Settings (personal)

**Deskripsi:** Pengaturan aksesibilitas per pengguna untuk meningkatkan keterbacaan dan kemudahan interaksi.

**Prioritas / Fase / Tier:** P2 / Fase 3 / platform

**Aktor & izin:** Semua pengguna.

**Aturan perilaku (acceptance criteria):**

1. Opsi "Always underline links" menampilkan garis bawah pada semua tautan teks.
2. ⚠ Opsi pola/ikon pada elemen berbasis warna (lozenge, chart, label) agar informasi tidak bergantung pada warna saja.
3. Animasi dan transisi dikurangi atau dimatikan bila `prefers-reduced-motion` aktif, tanpa perlu pengaturan tambahan.
4. ⚠ Opsi theme increased contrast tersedia dari pengaturan aksesibilitas (sama dengan NAV-034 butir 4).
5. Toggle keyboard shortcuts tersedia di sini untuk mencegah konflik dengan screen reader (NAV-038).
6. Ukuran teks mengikuti zoom browser; tidak ada pengaturan ukuran font khusus dan layout tetap berfungsi pada zoom hingga 200%.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Persyaratan kepatuhan menyeluruh di NAV-036 dan NFR-022 Accessibility (WCAG 2.1 AA) di [modul 27](27-nfr-arsitektur-platform.md) (kriteria uji dan gate CI); pengaturan personal di sini tidak boleh menurunkan baseline NFR-022 (mis. zoom 200% pada butir 6 sama dengan NFR-023 butir 6).

### NAV-036 Accessibility Compliance (WCAG 2.1 AA)

**Deskripsi:** Persyaratan aksesibilitas produk yang mengikuti WCAG 2.1 level AA beserta dokumentasi VPAT/ACR.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna, termasuk pengguna dengan disabilitas.

**Aturan perilaku (acceptance criteria):**

1. Semua fungsi dapat dioperasikan dengan keyboard saja; urutan fokus logis mengikuti urutan visual dan focus ring selalu terlihat.
2. Setiap halaman memiliki skip links dan landmark ARIA (banner, navigation, main, complementary).
3. Ikon/tombol memiliki label aksesibel dan setiap halaman memiliki struktur heading berurutan (satu H1).
4. Kontras warna minimal 4,5:1 untuk teks normal dan 3:1 untuk komponen UI serta teks besar; informasi tidak disampaikan hanya lewat warna.
5. Setiap interaksi drag-and-drop (board, backlog, ranking, resize sidebar) memiliki alternatif keyboard/menu.
6. Flag, pembaruan real-time, dan validasi form diumumkan lewat live region (`aria-live`).
7. Modal memiliki focus trap, mengembalikan fokus ke pemicu saat ditutup, dan dapat ditutup dengan Esc.
8. Produk diuji dengan screen reader NVDA, JAWS, VoiceOver, dan TalkBack untuk alur inti.
9. Layout mendukung reflow pada zoom 200-400%, penyesuaian text spacing, dan reduced motion.
10. Gambar memiliki alt text dan chart memiliki alternatif data (tabel atau ringkasan teks).
11. Galat form terasosiasi ke field melalui `aria-describedby` dan field invalid ditandai `aria-invalid`.
12. ⚠ Dokumen VPAT/ACR dipublikasikan per rilis mayor dengan daftar known issues.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Pengujian otomatis aksesibilitas (axe atau sejenis) dijalankan di CI frontend; target NFR di [modul 27](27-nfr-arsitektur-platform.md).
- Fitur ini sama cakupannya dengan NFR-022 Accessibility (WCAG 2.1 AA) di [modul 27](27-nfr-arsitektur-platform.md): NFR-022 memiliki target dan gerbang kualitas lintas produk, fitur ini memiliki aturan perilaku UI shell; prioritas keduanya P1 / Fase 2.

### NAV-037 Per-user UI State Persistence

**Deskripsi:** Penyimpanan otomatis preferensi tampilan kecil per pengguna agar UI kembali seperti terakhir digunakan.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna; hanya pemilik yang dapat membaca/menulis state UI-nya, kecuali admin lewat API user properties.

**Aturan perilaku (acceptance criteria):**

1. Lebar dan state collapse sidebar tersimpan dan dipulihkan (NAV-002).
2. Lebar panel detail di board/backlog/list dan posisi pemisah tersimpan per tipe halaman.
3. View terakhir per proyek (board/list/timeline) dan board terakhir per proyek tersimpan dan dipakai saat membuka proyek tanpa view eksplisit di URL.
4. ⚠ Kolom list/navigator, sort, dan group by tersimpan per pengguna per view.
5. ⚠ Quick filter dan view settings board yang dipilih tersimpan per pengguna; quick filter aktif hanya bertahan dalam sesi browser.
6. State expand/collapse section di work item view (Details, More fields) dan pinned fields tersimpan per pengguna.
7. ⚠ Preferensi cara membuka work item (modal vs sidebar/panel) tersimpan per pengguna.
8. Tab terakhir di For you tersimpan (NAV-010).
9. State disimpan server-side sebagai user properties (setara REST `/user/properties`, key-value JSON maksimum 32 KB per key) agar lintas perangkat; hal kecil dan sementara memakai localStorage sebagai fallback yang dibungkus penanganan galat.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Entitas user property di [modul 28](28-model-data-inti.md); API di [modul 23](23-api-publik-webhook.md). State di URL (NAV-043) selalu mengalahkan state tersimpan.

## Keyboard shortcut & perintah cepat

### NAV-038 Keyboard Shortcuts Toggle & Help Dialog

**Deskripsi:** Dialog daftar keyboard shortcut yang dibuka dengan `?` serta toggle untuk mengaktifkan atau menonaktifkan shortcut.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna.

**Aturan perilaku (acceptance criteria):**

1. `?` (Shift+/) membuka dialog shortcut dari halaman mana pun, kecuali saat fokus berada di input/textarea/editor.
2. Dialog dikelompokkan per kategori: Global, Navigating, Board, Backlog, Work item actions, Search/Navigator, Editor/Formatting; kategori halaman aktif ditampilkan pertama.
3. Toggle "Keyboard shortcuts" on/off tersedia di dalam dialog dan di Personal settings; nilai tersimpan per pengguna dan berlaku segera.
4. Shortcut single-key tidak aktif saat fokus berada di input/textarea/editor; Esc keluar dari field dan mengembalikan mode shortcut.
5. Shortcut berurutan (sequence) seperti `g` lalu `d` didukung dengan timeout 1 detik antar tombol; melewati timeout membatalkan sequence.
6. Label tombol menyesuaikan OS: Cmd di macOS, Ctrl di Windows/Linux.
7. ⚠ Dialog menyediakan kotak pencarian yang menyaring shortcut berdasarkan nama aksi atau tombol.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Registry shortcut terpusat di frontend agar katalog NAV-039 sampai NAV-042 dan dialog ini selalu sinkron (satu sumber data).

### NAV-039 Keyboard Shortcuts Catalog - Global & Navigation

**Deskripsi:** Daftar shortcut global untuk navigasi dan aksi umum di seluruh produk ini.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna; shortcut admin hanya untuk Jira admin.

**Aturan perilaku (acceptance criteria):**

1. `/` memindahkan fokus ke quick search (NAV-014).
2. `c` membuka dialog Create (NAV-013).
3. `?` membuka dialog bantuan shortcut (NAV-038).
4. `[` men-toggle sidebar (NAV-002).
5. `g` lalu `d` membuka Dashboards.
6. `g` lalu `p` membuka daftar proyek (Spaces).
7. ⚠ `g` lalu `a` membuka daftar Boards.
8. `g` lalu `i` membuka Search work items (issue navigator).
9. ⚠ `g` lalu `g` membuka admin search, hanya untuk Jira admin; pengguna lain tidak mendapat respons apa pun.
10. `Esc` menutup dialog, panel, atau dropdown paling atas.
11. `Ctrl/Cmd+Enter` men-submit form atau comment yang sedang difokuskan.
12. Semua shortcut single-key dapat dimatikan agar tidak konflik dengan screen reader; shortcut berbasis modifier (Ctrl/Cmd+Enter, Esc) tetap aktif.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Shortcut editor (markdown autoformat, `/`, `@`, `:`) dimiliki [07](07-editor-konten-kaya.md).

### NAV-040 Keyboard Shortcuts Catalog - Board & Backlog

- **Alias inventaris:** Board & Backlog Keyboard Shortcuts; Board keyboard shortcuts

**Deskripsi:** Shortcut untuk navigasi dan aksi cepat di board Scrum/Kanban dan backlog. Fitur ini menggabungkan jatah dari domain board Scrum/backlog dan board Kanban.

**Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard

**Aktor & izin:** Semua pengguna; aksi yang dipicu shortcut tetap tunduk pada permission (Edit issues, Assign issues, Add comments, Schedule issues untuk rank).

**Aturan perilaku (acceptance criteria):**

1. `j` / `k` memilih work item/kartu berikutnya/sebelumnya; bila belum ada yang terpilih, `j` memilih item pertama di kolom/daftar paling kiri/atas.
2. `n` / `p` memindahkan seleksi ke kolom berikutnya/sebelumnya di board (memilih kartu teratas kolom tujuan).
3. `o` atau `Enter` membuka detail work item terpilih (panel detail atau work item view sesuai preferensi NAV-037).
4. ⚠ `e` membuka mode edit work item terpilih.
5. `i` meng-assign work item terpilih ke pengguna saat ini.
6. `a` membuka picker assignee untuk work item terpilih.
7. `m` membuka input comment untuk work item terpilih.
8. `l` membuka editor labels untuk work item terpilih.
9. `.` membuka Commands dialog (NAV-042) untuk work item terpilih, termasuk aksi add/remove flag.
10. ⚠ `s` lalu `t` mengirim work item terpilih ke urutan teratas dan `s` lalu `b` ke urutan terbawah di backlog (ranking); di board shortcut ini berlaku di dalam kolom bila board diurutkan berdasarkan rank.
11. ⚠ Pada board company-managed, `1` membuka Backlog, `2` Active sprint/Board, dan `3` Reports.
12. ⚠ `-` men-collapse/expand semua swimlane di board.
13. ⚠ `t` men-toggle panel detail work item terpilih.
14. Setiap aksi drag-and-drop memiliki alternatif keyboard/menu "Move to"/"Move work item" (ke kolom, sprint, atau posisi rank tertentu) untuk aksesibilitas.
15. `/` memindahkan fokus ke kotak pencarian board/backlog dan `c` membuka dialog Create dengan proyek dan sprint/kolom ter-prefill.
16. `?` menampilkan dialog shortcut dengan kategori Board/Backlog di urutan pertama.
17. ⚠ `z` (presentation/projector mode layar penuh) dari board lama tidak dibangun; tekan `z` tidak melakukan apa pun.
18. Ctrl/Cmd+klik menambah/menghapus item dari multi-select dan Shift+klik memilih rentang item; aksi bulk berikutnya (drag, `.`) berlaku pada semua item terpilih.

**Team-managed vs company-managed:** Shortcut `1`/`2`/`3` hanya pada board company-managed; team-managed menavigasi antar view lewat tab proyek. Shortcut lain sama.

**Catatan & ketergantungan:**

- Perilaku aksi (rank, assign, flag, swimlane) dimiliki [10](10-board-kolom-kartu.md) dan [11](11-backlog-sprint.md).
- Tiga sumber inventaris (katalog navigasi, board Scrum/backlog, board Kanban) memuat banyak shortcut identik (`j`/`k`, `n`/`p`, `s`+`t`, `-`, `z`, `1`/`2`/`3`, `e`, `?`); masing-masing ditulis sekali dan tanda ⚠ dipertahankan bila salah satu sumber meragukannya.

### NAV-041 Keyboard Shortcuts Catalog - Work Item View & Navigator

- **Alias inventaris:** Keyboard Shortcuts on Work Items

**Deskripsi:** Shortcut di work item view dan daftar hasil search/list view untuk navigasi dan aksi pada work item. Fitur ini menggabungkan jatah dari domain work item inti.

**Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard

**Aktor & izin:** Semua pengguna; aksi tunduk pada permission work item (Edit issues, Assign issues, Add comments, Transition issues).

**Aturan perilaku (acceptance criteria):**

1. ⚠ `e` membuka mode edit work item (fokus ke field pertama yang dapat diedit).
2. `a` membuka picker assignee.
3. `i` meng-assign work item ke pengguna saat ini.
4. `m` membuka input comment dan memindahkan fokus ke editor comment.
5. `l` membuka editor labels.
6. `c` membuka dialog Create dengan proyek ter-prefill dari work item yang sedang dibuka.
7. `.` membuka Commands dialog (NAV-042) berisi operasi work item (transition, clone, move, log work, dll.).
8. ⚠ `s` membuka Share dialog (NAV-016).
9. ⚠ `w` men-toggle watch/unwatch work item dan menampilkan flag hasilnya.
10. ⚠ `,` membuka edit field cepat (field picker untuk mengedit satu field).
11. `j` / `k` berpindah ke work item berikutnya/sebelumnya di navigator, list view, atau detail view yang dibuka dari hasil search.
12. `u` kembali ke daftar hasil search dari detail view navigator.
13. `o` atau `Enter` membuka work item terpilih di daftar.
14. `Ctrl/Cmd+Enter` menyimpan comment atau description yang sedang diedit dan men-submit form.
15. Transisi cepat: tekan `.` lalu ketik nama status untuk menampilkan transisi yang valid, lalu Enter untuk menjalankannya.
16. `?` menampilkan dialog shortcut dengan kategori Work item actions di urutan pertama.
17. Shortcut dapat dimatikan lewat preferensi personal (NAV-038), dan saat dimatikan tidak ada satu pun shortcut single-key di halaman ini yang aktif.
18. Shortcut tidak aktif saat fokus berada di input, textarea, atau editor; teks yang diketik tidak pernah memicu aksi.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Operasi work item dimiliki [06](06-work-item-inti.md); transisi di [09](09-workflow-status-transisi.md); navigator di [13](13-pencarian-filter-ekspor.md).
- Shortcut navigator (`/`, `j`/`k`, `o`/Enter, `u`, `?`) tumpang tindih dengan SRC-021 Navigator keyboard shortcuts di [modul 13](13-pencarian-filter-ekspor.md); katalog dan definisi tombol dimiliki fitur ini, SRC-021 hanya mendaftarkan perilaku di halaman hasil pencarian. Prioritas yang dipakai adalah P1 / Fase 2 dan SRC-021 (saat ini P2 / Fase 3) perlu disamakan ke nilai ini.
- Prioritas dinaikkan dari P2 (katalog navigasi) ke P1/Fase 2 mengikuti alias dari domain work item inti, karena shortcut dasar work item (`c`, `a`, `i`, `m`, `?`, `j`/`k`) dipakai harian. Sub-fitur yang identik antar dua sumber (Ctrl/Cmd+Enter, `j`/`k`, `e`/`w`/`s`) ditulis sebagai satu butir.

### NAV-042 Commands Dialog (quick operations '.')

**Deskripsi:** Dialog perintah cepat berbasis keyboard untuk menjalankan aksi pada work item terpilih/terbuka tanpa mouse.

**Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard

**Aktor & izin:** Semua pengguna; daftar aksi disaring permission dan workflow.

**Aturan perilaku (acceptance criteria):**

1. Dialog dibuka dengan `.` pada work item terpilih (board, backlog, navigator) atau di work item view; tanpa work item terpilih, `.` tidak melakukan apa pun.
2. Daftar aksi dapat dicari: transition ke status tertentu, assign, add label, log work, clone, move, delete, link, attach, add flag, convert to subtask/work item.
3. Pencarian memakai fuzzy match dengan highlight karakter yang cocok; Arrow memilih dan Enter menjalankan aksi.
4. Aksi dibatasi permission dan workflow: hanya transisi yang valid dari status saat ini yang tampil, dan aksi tanpa izin tidak ditampilkan.
5. ⚠ Command palette global `Ctrl/Cmd+K` untuk navigasi lintas halaman (mencari halaman, proyek, dan aksi global) disediakan sebagai pelengkap dialog `.`.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Butir 5 tidak ada di referensi Cloud saat ini; diusulkan sebagai keputusan terbuka (lihat Pertanyaan terbuka).

## URL, umpan balik UI & pembaruan real-time

### NAV-043 URL Schemes & Deep Links

**Deskripsi:** Struktur URL yang stabil dan dapat dibagikan untuk semua halaman, termasuk redirect URL legacy dan deep link mobile.

**Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard

**Aktor & izin:** Semua pengguna; resolusi URL selalu melewati pemeriksaan permission sebelum konten dikirim.

**Aturan perilaku (acceptance criteria):**

1. `/browse/KEY-123` membuka work item dan `/browse/KEY` membuka proyek; key tidak peka huruf besar/kecil dan di-redirect ke bentuk kanonik huruf besar.
2. Key lama tetap resolve (redirect 301 ke key baru) setelah project key diganti atau work item dipindah ke proyek lain.
3. Board: `/jira/software/projects/KEY/boards/{id}` (team-managed) dan `/jira/software/c/projects/KEY/boards/{id}` (company-managed); backlog di sub-path `/backlog`.
4. Query `?selectedIssue=KEY-1` membuka panel detail work item di board, backlog, atau list tanpa meninggalkan halaman.
5. Proyek bisnis memakai `/jira/core/projects/KEY/...` dan proyek Service Management memakai `/jira/servicedesk/projects/KEY/...`.
6. Halaman global: `/jira/for-you`, `/jira/your-work`, `/jira/projects`, `/jira/filters`, `/jira/dashboards`, `/jira/plans`, `/jira/people/{accountId}`.
7. Issue navigator: `/issues/?jql=...` dan `/issues/?filter={id}`.
8. Dashboard: `/jira/dashboards/{id}`; URL legacy `/secure/Dashboard.jspa?selectPageId={id}` di-redirect ke URL baru.
9. URL legacy `/secure/RapidBoard.jspa?rapidView={id}` di-redirect ke URL board baru sesuai tipe proyek.
10. Ekstensi memakai `/jira/apps/{appId}/{envId}` dan halaman pengaturan memakai `/jira/settings/...`.
11. Permalink comment (`?focusedCommentId={id}`) dan permalink worklog membuka work item, men-scroll, dan menyorot entri terkait.
12. State UI (tab, filter, group by, view) ter-encode di query param agar dapat dibagikan, dan state dari URL mengalahkan state tersimpan (NAV-037).
13. ⚠ Deep link mobile membuka PWA terpasang bila browser mendukung penanganan tautan oleh PWA, dan tab web bila tidak; universal link/App Link baru berlaku bila app native kelak disetujui (NAV-050).
14. Redirect login mempertahankan URL tujuan (parameter `continue`, hanya menerima path relatif same-origin untuk mencegah open redirect).
15. URL yang tidak ditemukan atau tanpa akses menampilkan halaman error yang sesuai (NAV-045) dengan status HTTP 404/403 yang benar untuk request API.

**Team-managed vs company-managed:** Board team-managed tanpa segmen `/c/`, company-managed dengan segmen `/c/`; keduanya di-redirect bila tipe salah.

**Catatan & ketergantungan:**

- Kondisi repo saat ini: `webui.RegisterRoutes` sebagai `NoRoute` mengembalikan `index.html` untuk GET yang tidak dikenal, sedangkan `/api/*` dan `/swagger/*` selalu JSON 404; skema di atas harus tidak bertabrakan dengan prefix tersebut, dan redirect legacy (`/browse`, `/secure/*`) ditangani router SPA atau handler backend khusus sebelum fallback.
- Riwayat key untuk redirect disimpan di [modul 04](04-proyek.md)/06.

### NAV-044 Flags, Banners & In-app Messaging

**Deskripsi:** Komponen umpan balik global: flag/toast untuk hasil aksi, banner tingkat workspace, dan dialog konfirmasi.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna; Admin Situs (Administer Jira) mengonfigurasi announcement banner.

**Aturan perilaku (acceptance criteria):**

1. Flag bertipe sukses/info/warning/error tampil di pojok kiri bawah dan ditumpuk hingga 3 flag terlihat; flag sukses/info hilang otomatis setelah 8 detik (jeda saat di-hover/focus), flag error tetap sampai ditutup.
2. Flag dapat memuat aksi (mis. "View work item", "Undo", "Copy link", "Try again").
3. Aksi destruktif (delete, archive, discard) memakai modal konfirmasi dengan teks peringatan dan tombol berlabel aksi spesifik (mis. "Delete"), bukan "OK".
4. Admin dapat mengonfigurasi announcement banner (teks dengan format terbatas dan tautan; tanpa HTML bebas untuk mencegah XSS) dengan visibility public atau hanya pengguna login.
5. Banner sistem tampil untuk kondisi: trial akan habis, masalah billing, batas pengguna/storage hampir tercapai, dan maintenance/mode read-only.
6. Banner koneksi "You're offline"/"Reconnecting" tampil saat koneksi jaringan atau kanal realtime terputus dan hilang otomatis saat pulih.
7. ⚠ Cookie consent banner tampil sesuai wilayah pengguna dan pilihan tersimpan per browser.
8. Banner yang dapat di-dismiss menyimpan status dismiss per pengguna; banner kritis (read-only, billing) tidak dapat di-dismiss.
9. Isi flag dan banner diumumkan ke screen reader melalui `aria-live` (polite untuk info/sukses, assertive untuk error).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Komponen `alert` di `frontend/src/lib/components/ui` dapat dipakai sebagai dasar; konfigurasi announcement banner di [modul 22](22-administrasi-sistem.md).

### NAV-045 Error, Empty & Permission States

**Deskripsi:** Halaman dan state standar untuk error, data kosong, tidak ada akses, dan loading.

**Prioritas / Fase / Tier:** P0 / Fase 1 / platform

**Aktor & izin:** Semua pengguna, termasuk pengguna anonim.

**Aturan perilaku (acceptance criteria):**

1. Halaman 404 "We can't find this page" (atau "this work item") menampilkan tautan kembali ke For you; respons tidak membedakan "tidak ada" dan "tidak berhak" bila membedakannya membocorkan keberadaan item.
2. ⚠ Halaman "You don't have access" menampilkan tombol "Request access" bila diaktifkan admin, baik untuk akses aplikasi maupun akses ke proyek; permintaan diteruskan ke admin terkait.
3. Work item yang dihapus menampilkan pesan khusus "This work item was deleted"; work item yang dipindah ke proyek tanpa akses menampilkan pesan tidak ada akses.
4. Sesi kedaluwarsa memicu redirect ke login lalu kembali ke URL semula setelah login berhasil (NAV-043 butir 14).
5. Galat umum "Something went wrong" menampilkan tombol reload dan ID error/trace yang dapat disalin untuk support.
6. Empty state berilustrasi dengan CTA tersedia untuk daftar kosong (no work items, no filters, no dashboards, dll.).
7. Skeleton loading dipakai untuk list, board, dan work item view; spinner hanya untuk aksi singkat; skeleton tidak tampil untuk muatan di bawah 300 ms.
8. Halaman rate-limit (429) dan maintenance (503) menampilkan pesan dan waktu coba lagi bila diketahui.
9. ⚠ Anonymous access menampilkan halaman publik bila permission diberikan ke "Anyone on the web"; fitur ini tidak tersedia di edisi Free.
10. Proyek yang diarsipkan menampilkan banner "This space is archived" dan seluruh konten dalam mode read-only.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Komponen `empty` dan `skeleton` sudah tersedia di frontend; pesan galat backend memakai `constants.GetError` EN/ID. Kondisi repo saat ini: halaman `+error.svelte` lama sudah dihapus dan perlu dibangun ulang sesuai fitur ini.
- State kosong/galat/izin khusus hasil pencarian dimiliki SRC-007 Search result empty, error, and permission states ([modul 13](13-pencarian-filter-ekspor.md)) dan memakai komponen standar fitur ini.

### NAV-046 Real-time Refresh & Live Updates

**Deskripsi:** Pembaruan UI tanpa reload ketika data diubah pengguna lain atau proses otomatis.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna; event realtime hanya dikirim ke klien yang berhak melihat data terkait.

**Aturan perilaku (acceptance criteria):**

1. ⚠ Board dan backlog memperbarui kartu secara live (pindah kolom, ganti assignee, ubah rank) saat pengguna lain mengubahnya, untuk board team-managed maupun company-managed.
2. Bila perubahan tidak dapat diterapkan otomatis (mis. konfigurasi board berubah), ditampilkan banner "This board has been updated" dengan tombol Refresh.
3. Work item view menampilkan comment baru dan perubahan field secara real-time.
4. Bila field yang sedang diedit pengguna diubah orang lain, ditampilkan peringatan konflik dengan pilihan memuat nilai terbaru atau menimpa.
5. Badge dan panel notifikasi diperbarui real-time (NAV-018).
6. Aksi pengguna memakai optimistic UI; bila server menolak, perubahan di-rollback dan flag error ditampilkan.
7. ⚠ Description mendukung indikator presence (siapa yang sedang melihat/mengedit) tanpa co-editing penuh pada rilis awal.
8. Bila WebSocket tidak tersedia klien memakai fallback polling (interval 30 detik) atau SSE, dan reconnect otomatis dengan exponential backoff.
9. Perubahan massal (bulk change) diringkas menjadi satu prompt refresh, bukan banyak pembaruan beruntun.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Kanal WebSocket/SSE per proyek/board dan katalog event di [modul 27](27-nfr-arsitektur-platform.md) dan 28.
- Tumpang tindih dengan ISS-014 Real-time Updates & Concurrent Editing ([modul 06](06-work-item-inti.md), pemilik aturan konflik edit di work item view) dan NFR-009 Real-time Updates (SSE / WebSocket) ([modul 27](27-nfr-arsitektur-platform.md), pemilik transport, fallback, dan reconnect); fitur ini memiliki perilaku UI lintas halaman (banner refresh, optimistic UI, badge).

## Perangkat, browser & cetak

### NAV-047 Responsive Web Experience

**Deskripsi:** Perilaku UI web pada berbagai ukuran layar dan perangkat sentuh.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna.

**Aturan perilaku (acceptance criteria):**

1. Breakpoint yang didukung: desktop (≥ 1280 px), laptop kecil (1024-1279 px), tablet (768-1023 px), dan mobile web (< 768 px).
2. Di layar sempit sidebar menjadi overlay/drawer dan top bar memadatkan item (NAV-001 butir 5, NAV-002 butir 11).
3. Di lebar sempit kolom kanan work item view (details) pindah ke bawah konten utama.
4. Board memakai horizontal scroll kolom dengan lebar minimum kartu 240 px; kartu tidak pernah terpotong.
5. Tabel/list memakai horizontal scroll dengan kolom key/summary sticky.
6. Modal menjadi full-screen di mobile web.
7. ⚠ Drag-and-drop dasar (pindah kartu, rank) mendukung input sentuh dengan long-press; bila tidak memungkinkan tersedia menu "Move to".
8. ⚠ Di mobile web tidak tampil smart app banner "Open in app" selama app native di luar cakupan; sebagai gantinya browser yang mendukung dapat menampilkan prompt "Install app" PWA yang dapat di-dismiss dan tidak muncul lagi setelah di-dismiss (NAV-050).
9. Konten reflow tanpa kehilangan fungsi pada zoom 200% (WCAG 1.4.10) dan tanpa horizontal scroll halaman di lebar 320 px CSS.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Hook `is-mobile` dan komponen `drawer` di frontend sudah tersedia.

### NAV-048 Browser Support

**Deskripsi:** Kebijakan dukungan browser untuk aplikasi web produk ini.

**Prioritas / Fase / Tier:** P1 / Fase 2 / platform

**Aktor & izin:** Semua pengguna.

**Aturan perilaku (acceptance criteria):**

1. Didukung: dua versi mayor terbaru Chrome, Firefox, Microsoft Edge (Chromium), dan Safari, masing-masing di desktop dan mobile (termasuk Safari iOS), sama dengan NFR-023 butir 1.
2. Tidak didukung: Internet Explorer 11 dan Edge Legacy.
3. Browser mobile didukung penuh melalui layout responsif (NAV-047, NFR-023 butir 2); pengguna tidak dialihkan ke aplikasi mobile native, karena app native di luar cakupan (NAV-050).
4. Aplikasi mewajibkan JavaScript dan cookies first-party; beberapa integrasi/smart link pihak ketiga memerlukan third-party cookies dan menampilkan pesan bila diblokir.
5. Browser yang tidak didukung menampilkan banner "Your browser isn't supported" yang dapat di-dismiss tanpa memblokir akses.
6. Kebijakan rolling: daftar versi yang didukung bergeser otomatis mengikuti dua versi mayor terbaru; CI menjalankan test E2E pada engine chromium, firefox, dan webkit untuk alur kritis (NFR-023 butir 5).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Sumber kebenaran kebijakan browser adalah NFR-023 Browser & Device Support di [modul 27](27-nfr-arsitektur-platform.md); NAV-048 hanya mengatur perilaku UI (banner, pesan cookies) dan harus mengikuti NFR-023 bila ada perubahan. Target build Vite/SvelteKit (browserslist) diselaraskan dengan NFR-023.

### NAV-049 Print Styles & Printable Views

**Deskripsi:** Dukungan cetak halaman work item dan hasil search.

**Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard

**Aktor & izin:** Semua pengguna yang dapat melihat konten; export tunduk pada global permission export bila ada.

**Aturan perilaku (acceptance criteria):**

1. Print CSS untuk work item view menyembunyikan navigasi/sidebar/top bar dan menampilkan field utama, description, dan comments dengan warna yang ramah cetak.
2. ⚠ Menu "..." work item menyediakan aksi "Print", "Export Word", dan "Export XML".
3. Navigator menyediakan export "Printable" (HTML) untuk hasil search ([modul 13](13-pencarian-filter-ekspor.md)).
4. ⚠ Dashboard dan board mengandalkan print browser tanpa print view khusus, dengan print CSS dasar yang menyembunyikan chrome navigasi.
5. Export PDF native tidak disediakan; pengguna dapat memakai "Save as PDF" dari dialog print browser.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Export PDF lewat app Marketplace pihak ketiga berada di luar cakupan.

### NAV-050 Jira Mobile Apps (iOS & Android) Parity

**Deskripsi:** Kebutuhan mobile setara aplikasi Jira untuk iOS dan Android. Sikap produk: aplikasi native iOS/Android tidak dibangun (sejalan dengan README Ruang lingkup dan NFR-023); padanannya adalah PWA di atas web responsif (NAV-047) dengan Web Push. Butir di bawah menjadi acceptance criteria untuk padanan PWA; butir yang hanya mungkin di app native ditandai ⚠ dan diperlakukan sebagai rujukan.

**Prioritas / Fase / Tier:** P3 / Fase — untuk app native; padanan PWA + Web Push: Fase 4 / platform

**Aktor & izin:** Semua pengguna; admin organisasi menetapkan kebijakan mobile.

**Aturan perilaku (acceptance criteria):**

1. Login dengan akun produk ini/SSO, mendukung multi-workspace dan multi-account switching (NAV-008).
2. Home menampilkan assigned, recent, dan starred.
3. ⚠ Notifications feed dengan push notification lewat Web Push (izin browser, service worker); pengaturan push per tipe dan Do not disturb di tingkat produk. Critical alerts yang menembus mode senyap OS (iOS Critical alerts, bypass DND Android) tidak dapat dipenuhi PWA dan hanya berlaku bila app native kelak disetujui.
4. ⚠ Search work items dan filters, termasuk JQL penuh.
5. ⚠ Work item view mendukung edit field, transition, assign, comment dengan mention, attachment dari kamera/galeri, dan log work.
6. Create work item dengan validasi field wajib.
7. ⚠ Board: lihat dan drag kartu antar kolom serta filter; backlog dan sprint dengan aksi terbatas.
8. ⚠ Dashboards: menampilkan gadget yang mendukung render mobile; gadget lain menampilkan tautan "Open in browser".
9. Approvals dan queues Service Management tersedia untuk agent dan approver ([modul 31](31-jsm-antrian-sla-knowledge-base-laporan.md) dan 32).
10. Dark mode mengikuti OS; aksesibilitas VoiceOver/TalkBack dan dynamic type didukung.
11. Deep link dari email/notifikasi membuka PWA terpasang (atau tab browser) pada item terkait (NAV-043 butir 13).
12. ⚠ App lock/biometric dan kebijakan mobile app management (MAM) dari admin organisasi hanya berlaku untuk app native; pada PWA kebijakan setara terbatas pada sesi dan perangkat (AUTH-009) serta timeout sesi.
13. ⚠ Tidak tersedia di mobile: admin/Project settings, workflow editor, automation editor, dan Plans.
14. ⚠ Aplikasi desktop native (setara Jira Cloud for Mac yang sudah dihentikan) tidak dibangun.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Keputusan: app native = Fase — (di luar cakupan, lihat "Di luar cakupan & ekosistem"); padanan PWA + Web Push = Fase 4. Keputusan final masih Pertanyaan terbuka nomor 8 dan README (keputusan terbuka nomor 1).
- Fitur turunan yang bergantung pada keputusan ini: USR-039 (kontrol akses klien mobile, [modul 02](02-pengguna-grup-organisasi.md)), TIME-049 (pencatatan waktu mobile, [modul 14](14-pelacakan-waktu-estimasi.md)), DSH-060 (gadget mobile, [modul 17](17-dashboard-gadget.md)), NTF-038 (mobile push, [modul 19](19-notifikasi-email.md)), JSM-AI-065 ([modul 35](35-jsm-ai-otomasi-api-admin-csm.md)), dan JSM-OPS-029 Mobile App for On-call ([modul 34](34-jsm-operations-alerting-on-call.md)), yang butir Critical alerts iOS/DND Android-nya hanya dapat dipenuhi app native. Modul pemilik fitur tersebut menyesuaikan prioritas/fasenya dengan keputusan di sini.
- Push notification bergantung pada [modul 19](19-notifikasi-email.md); prompt instalasi PWA di NAV-047 butir 8.

## Asisten AI (opsional)

### NAV-051 Rovo Chat & Rovo Agents in Jira

**Deskripsi:** Asisten AI percakapan yang dibuka dari top bar dan konteks halaman, serta agent yang dapat menjalankan tugas di produk ini (setara Rovo Chat & Rovo Agents).

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Pengguna terlisensi pada edisi berbayar; admin organisasi mengaktifkan/menonaktifkan AI dan membuat agent kustom.

**Aturan perilaku (acceptance criteria):**

1. Tombol asisten AI di top bar membuka panel chat di sisi kanan yang sadar konteks halaman (work item atau board yang sedang dibuka).
2. ⚠ Entry point asisten tersedia juga di work item view dan menu editor.
3. Pengguna dapat bertanya tentang work items, meminta ringkasan, dan mencari informasi lintas aplikasi produk ini.
4. ⚠ Tersedia agent bawaan (mis. work breakdown/planner, readiness checker) dan agent kustom yang dibuat admin dengan instruksi dan sumber pengetahuan.
5. ⚠ Agent dapat dipanggil dari aturan otomasi atau transisi workflow, dan work item dapat di-assign ke agent ([modul 20](20-otomasi.md)).
6. Panel menyimpan riwayat percakapan per pengguna, menyediakan mulai chat baru, salin jawaban, dan feedback thumbs up/down.
7. Jawaban menyertakan sumber/sitasi yang dapat diklik dan hanya merujuk data yang boleh diakses pengguna.
8. Admin dapat menonaktifkan AI per organisasi/aplikasi; saat nonaktif tombol dan entry point tidak dirender.
9. ⚠ Jawaban diberikan dalam bahasa pengguna (minimal EN/ID).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Opsional; dapat diintegrasikan dengan penyedia LLM yang dapat dikonfigurasi admin. Data handling dan kuota diatur di [modul 22](22-administrasi-sistem.md) dan [modul 27](27-nfr-arsitektur-platform.md).

### NAV-052 Atlassian Intelligence features in Jira (catalog)

**Deskripsi:** Katalog fitur AI yang terlihat di produk ini (setara Atlassian Intelligence/Rovo). Dicatat sebagai referensi cakupan, bukan kewajiban MVP; implementasi detail ada di modul domain terkait.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Pengguna terlisensi edisi berbayar; admin organisasi mengatur toggle AI.

**Aturan perilaku (acceptance criteria):**

1. Tombol "Summarize" di work item view merangkum details dan comments.
2. Menu AI di editor (termasuk perintah `/ai`) menyediakan generate/rewrite description, ubah tone, perbaiki ejaan, dan buat lebih ringkas ([modul 07](07-editor-konten-kaya.md)).
3. Saran child work items/work breakdown dari epic atau deskripsi.
4. Natural language to JQL di Search work items ("Search with AI") ([modul 12](12-jql-bahasa-kueri.md)).
5. Pembuatan aturan otomasi dengan bahasa natural ([modul 20](20-otomasi.md)).
6. ⚠ Saran work item serupa/terkait dan deteksi duplikat saat create.
7. ⚠ Pembuatan work items dari halaman/teks yang disorot di aplikasi wiki.
8. ⚠ Ringkasan smart link/halaman wiki yang ditempel di work item.
9. Definitions untuk istilah internal organisasi yang tampil saat istilah di-hover.
10. ⚠ Rekomendasi template berbasis AI saat membuat proyek.
11. ⚠ Insight berbasis AI pada halaman Summary dan laporan ([modul 16](16-laporan-insight.md)).
12. ⚠ Terjemahan konten work item ke bahasa pengguna.
13. ⚠ Admin dapat men-toggle AI per aplikasi; UI menampilkan disclaimer penanganan data dan kuota AI per edisi.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- AI Service Management dicatat di [modul 35](35-jsm-ai-otomasi-api-admin-csm.md). Semua fitur di sini bergantung pada infrastruktur AI opsional NAV-051.

### NAV-053 Atlassian Browser Extension (Rovo)

**Deskripsi:** Ekstensi browser resmi (Chrome/Edge) yang membawa pencarian, preview tautan, dan asisten AI produk ini ke halaman web mana pun tanpa membuka aplikasi.

**Prioritas / Fase / Tier:** P3 / Fase 4 / platform

**Aktor & izin:** Pengguna terlisensi edisi berbayar; admin organisasi mengontrol ketersediaan.

**Aturan perilaku (acceptance criteria):**

1. Pencarian lintas aplikasi (termasuk work item) tersedia dari toolbar/side panel browser.
2. ⚠ URL produk ini yang muncul di halaman web pihak ketiga ditampilkan dengan preview smart link saat di-hover.
3. ⚠ Asisten AI memakai konteks halaman yang sedang dibuka (mis. merangkum halaman, membuat work item dari isi halaman).
4. Login memakai akun produk ini dan semua hasil dibatasi permission pengguna.
5. ⚠ Ekstensi memerlukan edisi berbayar dengan fitur AI aktif.
6. ⚠ Admin dapat mengontrol ketersediaan lewat pengaturan AI organisasi dan kebijakan browser terkelola.
7. Ekstensi ini terpisah dari ekstensi browser Product Discovery untuk menangkap insight ([modul 37](37-jpd-insights-delivery-peran-integrasi.md)).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Opsional/ekosistem; bergantung pada REST API publik ([modul 23](23-api-publik-webhook.md)) dan OAuth 2.0.

## Referensi legacy & Data Center

### NAV-054 Legacy Navigation (horizontal top nav 2020-2025)

**Deskripsi:** Navigasi lama dengan top nav horizontal berisi dropdown dan sidebar khusus project. Dicatat hanya sebagai referensi perilaku; produk ini langsung memakai navigasi baru (NAV-001, NAV-002, NAV-003).

**Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated

**Aktor & izin:** Semua pengguna (referensi).

**Aturan perilaku (acceptance criteria):**

1. Top nav lama (logo, Your work, Projects, Filters, Dashboards, Teams/People, Plans, Apps, Create, search, notifications, help, settings, profile) tidak dibangun; fungsinya dipetakan ke NAV-001 dan NAV-002.
2. Dropdown Projects (recent, starred, View all, Create) dipetakan ke section Spaces (daftar proyek) sidebar.
3. Dropdown Filters (starred/recent, View all, Advanced issue search) dipetakan ke section Filters sidebar.
4. Dropdown Dashboards (starred/recent, View all, Create) dipetakan ke section Dashboards sidebar.
5. Dropdown Teams/People (search people/teams, Start a team) dipetakan ke item Teams (NAV-006).
6. Dropdown Apps (global pages, Explore more apps, Manage your apps) dipetakan ke section Apps (NAV-009).
7. Sidebar project lama (grup Planning, Development, Operations, Project pages, Add shortcut, Project settings) dipetakan ke navigasi proyek (NAV-003).
8. Footer "You're in a team-managed/company-managed project" diganti informasi tipe di header proyek (NAV-003 butir 1).
9. Board switcher di header sidebar dipetakan ke board switcher NAV-003 butir 4.
10. "Add shortcut" dipetakan ke Shortcuts proyek (NAV-003 butir 6).
11. ⚠ Toggle opt-in/opt-out navigasi lama/baru tidak disediakan karena produk ini hanya memiliki satu navigasi.
12. Navigasi era 2018 (sidebar global biru gelap) dan header ala Server tidak dibangun.

**Team-managed vs company-managed:** Tidak ada perbedaan (referensi saja).

**Catatan & ketergantungan:**

- Header dropdown gaya Server/DC masih dipakai di Data Center; berada di luar cakupan.

### NAV-055 Quick Search Smart Querying (legacy)

**Deskripsi:** Sintaks shortcut di quick search lama (mis. `my`, `overdue`, `r:me`) yang otomatis diubah menjadi JQL. Hanya ada di Data Center; tidak dibangun di produk ini.

**Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated

**Aktor & izin:** Semua pengguna (referensi).

**Aturan perilaku (acceptance criteria):**

1. Kata kunci `my`, `r:me`, `overdue`, `created:today`, `updated:-1w` tidak dikonversi menjadi JQL; quick search memperlakukannya sebagai teks bebas.
2. Prefix project key, issue type, status, `c:` (component), dan `v:`/`ff:` (version) tidak didukung.
3. Kombinasi kata kunci dengan teks bebas tidak diproses khusus.
4. Pengguna yang butuh kueri terstruktur diarahkan ke Search work items/JQL ([modul 12](12-jql-bahasa-kueri.md)) melalui tautan "Advanced search" (NAV-014 butir 6).

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Alasan: perilaku sudah digantikan quick search dan pencarian lintas entitas di Cloud.
- Rujukan legacy yang sama tercatat di JQL-061 Quick Search Smart Querying (legacy) ([modul 12](12-jql-bahasa-kueri.md)) dan SRC-013 Quick search smart querying ([modul 13](13-pencarian-filter-ekspor.md)); ketiganya berstatus P3 / Fase — / dc-only-or-deprecated.

### NAV-056 User Profile & Preferences (Data Center)

**Deskripsi:** Halaman profil dan preferensi gaya lama yang hanya ada di Data Center, berisi pengaturan yang tidak tersedia di Cloud. Dicatat agar jelas pengaturan mana yang sengaja tidak dibangun.

**Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated

**Aktor & izin:** Semua pengguna (referensi).

**Aturan perilaku (acceptance criteria):**

1. Pilihan "My Jira Home" (halaman setelah login) tidak disediakan; landing selalu For you (NAV-010).
2. Preferences DC (Number of issues per page, Email type HTML/Text, Default sharing filter/dashboard) tidak disediakan; Language dan Time zone dipetakan ke NAV-029 (dan NAV-033), My changes ke NAV-019 butir 3, Keyboard shortcuts ke NAV-038, dan Autowatch ke NAV-031.
3. Profile summary gaya DC (username, groups, project roles, activity stream) diganti halaman profil berbasis accountId (NAV-027); username tidak ditampilkan.
4. Personal Access Tokens (PAT) tidak disediakan; akses REST API memakai API token akun ([modul 01](01-identitas-autentikasi-sesi.md)).
5. Ubah password lokal tersedia lewat Account settings (NAV-029) hanya untuk akun yang tidak memakai SSO.

**Team-managed vs company-managed:** Tidak ada perbedaan.

**Catatan & ketergantungan:**

- Referensi Data Center saja; lihat Di luar cakupan.

## Pertanyaan terbuka

1. Apakah produk ini mendukung multi-workspace dan multi-account dalam satu instalasi sejak awal, atau app/site switcher (NAV-007, NAV-008) cukup stub untuk satu workspace?
2. Apakah skema URL kompatibel Jira (`/browse`, `/jira/software/c/projects/...`, redirect `/secure/*.jspa`) dipertahankan persis demi kemudahan migrasi, atau disederhanakan dengan tetap menyediakan redirect dari URL Jira?
3. Bahasa UI apa saja selain EN/ID yang menjadi target, dan apakah terjemahan nama status/work type per bahasa (NAV-033 butir 12) masuk rilis awal?
4. Apakah gating edisi (Free/Standard/Premium/Enterprise, NAV-024) diberlakukan pada instalasi self-hosted, atau semua fitur terbuka dan gating hanya untuk layanan terkelola?
5. Apakah command palette global `Ctrl/Cmd+K` (NAV-042 butir 5) dibangun walaupun tidak ada di referensi Cloud saat ini?
6. Penyedia LLM mana yang didukung untuk asisten AI (NAV-051 sampai NAV-053), dan apakah fitur AI dapat dijalankan sepenuhnya on-premise?
7. Ke mana feedback produk (NAV-022) dan permintaan akses/upgrade (NAV-007, NAV-024) dikirim pada instalasi self-hosted?
8. Sikap usulan untuk NAV-050: app native iOS/Android = Fase — (di luar cakupan), padanan PWA + Web Push = Fase 4. Apakah pemilik produk menyetujuinya? Jika app native kelak masuk peta jalan, fitur turunan (USR-039, TIME-049, DSH-060, NTF-038, JSM-AI-065, JSM-OPS-029 butir Critical alerts) perlu dijadwal ulang; jika tidak, butir tersebut harus ditandai sebagai rujukan di modul masing-masing.
9. Berapa nilai final batas yang diusulkan di modul ini (retensi For you 90 hari, recent 20/200 per tipe, starred 100 per tipe, retensi notifikasi 30 hari, rate limit share 50 per jam)?
10. Apakah unwatch manual mencegah autowatch berikutnya pada work item yang sama (NAV-031 butir 4, diusulkan "ya"), dan apakah edit field ikut memicu autowatch (NAV-031 butir 2)?
11. Apakah halaman pencarian lintas produk (NAV-015) tetap P2 / Fase 4 atau disamakan dengan SRC-014 (P3 / Fase 4)?
12. Apakah Create dialog menyediakan picker status awal setelah Fase 1 (NAV-013 butir 10, ISS-002 butir 4)?

## Di luar cakupan & ekosistem

- **Produk Atlassian lain lintas produk:** Confluence, Compass, Loom, Bitbucket, Trello, Atlassian Home sebagai produk terpisah, dan konsol admin.atlassian.com tidak dibangun; app switcher (NAV-007) hanya memuat aplikasi produk ini.
- **App Marketplace pihak ketiga:** export PDF/Word lanjutan, tema/sidebar kustom, dan ekstensi navigasi pihak ketiga tidak dibangun sebagai fitur native; produk ini hanya menyediakan extension point (NAV-009) dan katalog ekstensi di [modul 25](25-integrasi-impor-ekspor-ekstensi.md).
- **Model ekstensi lama setara Connect** (generalPages, webPanels, webItems) tidak dibangun (NAV-009 butir 10).
- **Fitur Data Center-only / deprecated (rujukan saja):** Legacy Navigation (NAV-054), Quick Search Smart Querying (NAV-055), User Profile & Preferences Data Center termasuk My Jira Home, issues per page, format email HTML/Text, dan Personal Access Tokens (NAV-056), shortcut `z` presentation mode (NAV-040 butir 17), toggle opt-in navigasi lama/baru (NAV-054 butir 11).
- **Aplikasi mobile native iOS/Android** tidak dibangun; kebutuhan mobile dipenuhi web responsif (NAV-047) dan PWA + Web Push (NAV-050, Fase 4), sejalan dengan README dan NFR-023. Critical alerts iOS/bypass DND Android dan MAM khusus app native ikut di luar cakupan.
- **Aplikasi desktop native** (setara Jira Cloud for Mac yang dihentikan) tidak dibangun (NAV-050 butir 14).
- **Layanan AI Atlassian (Rovo, Rovo Studio, Atlassian Intelligence) dan konektor Rovo** tidak dipakai; padanannya di produk ini bersifat opsional dan memakai penyedia LLM yang dikonfigurasi (NAV-051 sampai NAV-053).
- **Layanan dukungan, komunitas, dan pelatihan vendor** (Atlassian Community, Atlassian University, support berjenjang) hanya berupa tautan yang dapat dikonfigurasi (NAV-020).
- **Atlassian Guard** (kebijakan autentikasi dan mobile app management) tidak dibangun sebagai produk terpisah; kebijakan setara dirinci di [modul 01](01-identitas-autentikasi-sesi.md) dan 02.
