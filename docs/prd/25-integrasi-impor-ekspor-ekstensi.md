# 25 — Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas

Modul ini mendefinisikan bagaimana produk ini bertukar data dan berinteraksi dengan dunia luar: fondasi keamanan request keluar (SSRF guard) dan resolver Smart Link, pola integrasi generik (iPaaS, helpdesk/CRM, application links, issue collector), seluruh jalur impor (CSV, JSON, Trello, Asana/monday.com, importer legacy), format ekspor hasil pencarian dan rencana, backup/restore/migrasi antarsitus beserta pipeline analitik, konektor chat (Slack, Microsoft Teams, Google Chat) dan alat produktivitas (Google Sheets, Excel, Outlook, Gmail), integrasi dokumentasi/desain/video (padanan Confluence, Figma, Loom, Trello, IDE), rujukan silang ke produk lintas suite (Opsgenie, Statuspage, Compass, Align, Rovo, Goals, Product Discovery, Assets, data provider DevOps), serta model ekstensi app: katalog dan manajemen app, izin/scope/consent, runtime app, dan seluruh extension point (panel issue, custom field type, workflow rule, fungsi JQL, gadget, halaman app, UI modifications, aksi otomasi, entity properties, menu aksi, permission app, event dan lifecycle). Modul ini memiliki job impor/ekspor, registry integrasi dan konektor, serta kontrak extension point; logika domain yang disentuh (work item, field, workflow, JQL, notifikasi, otomasi, admin) tetap dimiliki modul masing-masing dan hanya dirujuk. Kondisi repo saat ini baru memiliki fitur auth dan user (module path `app`, GORM + PostgreSQL, JWT di context `sess`), sehingga seluruh isi modul ini adalah pekerjaan baru yang diusulkan dikemas sebagai feature module `internal/features/importexport` (job impor/ekspor), `internal/features/integrations` (konektor chat/office/dokumen, Smart Link resolver), dan `internal/features/apps` (registry app, consent, extension point), dengan pengaman jaringan bersama di `pkg/netguard` dan entitas `ImportJob`, `ExportJob`, `IntegrationConnection`, `ChannelSubscription`, `RemoteLink`, `App`, `AppInstallation`, `AppGrant`, `EntityProperty` di `internal/shared/domain/entity`, endpoint di bawah `/api/v1/import`, `/api/v1/export`, `/api/v1/integrations`, `/api/v1/apps`, dan halaman SPA terkait di `frontend/`.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — plan tier, feature flag/entitlement, batas cakupan Marketplace.
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — OAuth consent dan connected apps (AUTH-013), API token.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — App users/bot (USR-029), bulk import user (USR-010), product access (USR-018).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — App Permissions (PERM-048), permission scheme, issue security.
- [04 Proyek](04-proyek.md) — Apps di project settings (PRJ-060), Issue Collectors (PRJ-061), Project Import from Backup DC (PRJ-067).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — custom field, context, screen, link type, Epic legacy field.
- [06 Work Item Inti](06-work-item-inti.md) — Web/Confluence/Remote Links (ISS-030), ekspor work item tunggal (ISS-039), Issue Collectors (ISS-040).
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — Smart Links inline/card/embed (EDT-020, EDT-021).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — remote links display (CMT-029), halaman Confluence terkait (CMT-030, CMT-031), Loom (CMT-032).
- [09 Workflow](09-workflow-status-transisi.md) — condition/validator/post function, impor/ekspor workflow (WF-045).
- [12 JQL](12-jql-bahasa-kueri.md) — pencarian entity property (JQL-036), remote link (JQL-029).
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — UI ekspor (SRC-050 s.d. SRC-058).
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — impor estimasi/worklog (TIME-045), time tracking provider.
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — gadget dari app (DSH-058).
- [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md) — ekspor timeline/list/plan (PLAN-013, PLAN-020, PLAN-054), Goals (PLAN-062).
- [19 Notifikasi & Email](19-notifikasi-email.md) — notifikasi personal Slack/Teams (NTF-039, NTF-040), infrastruktur email.
- [20 Otomasi](20-otomasi.md) — Incoming webhook (AUTO-025), Set entity property (AUTO-051), Send web request (AUTO-054).
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — menu Apps & extension point UI (NAV-009), Smart Links platform (NAV-017).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — halaman admin webhook, application links, issue collector, Manage apps, backup, External system import, migrasi (ADM-030 s.d. ADM-039).
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST API, OAuth 2.0 (3LO), webhook, rate limit, pagination.
- [24 Integrasi DevOps](24-integrasi-devops.md) — GitHub/GitLab/Bitbucket, data provider DevOps, IDE.
- [26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md) — proyek bisnis, formulir, list view.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — job queue, object storage, observability, multi-tenancy.
- [28 Model Data Inti](28-model-data-inti.md) — entitas dan katalog event domain yang menjadi kontrak impor/ekspor dan trigger app.
- [32 JSM ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — status page dan integrasi Statuspage untuk incident (JSM-ITSM-026, JSM-ITSM-027).
- [33 JSM Assets](33-jsm-assets-cmdb.md) — Assets objects custom field (JSM-AST-026) dan impor aset (JSM-AST-031 s.d. JSM-AST-036).
- [34 JSM Operations](34-jsm-operations-alerting-on-call.md) — integrasi alert dan sinkronisasi work item (JSM-OPS-031, JSM-OPS-038), migrasi Opsgenie (JSM-OPS-047).
- [37 JPD Insights & Delivery](37-jpd-insights-delivery-peran-integrasi.md) — delivery tab dan progress idea (JPD-INS-017 s.d. JPD-INS-022).

## Persona terkait

- **Admin Situs** — persona utama untuk impor site-level, backup/restore, migrasi, manajemen app, consent dan scope, kontrol Smart Link, application links, serta allowlist request keluar.
- **Admin Proyek/Project Lead** — mengimpor data ke proyeknya, menghubungkan channel chat dan halaman dokumentasi proyek, mengelola issue collector, dan mencabut akses app per proyek.
- **Scrum Master/Team Lead** — mengimpor backlog dan sprint dari alat lama, berlangganan notifikasi channel tim, mengekspor laporan sprint.
- **Developer/Anggota Tim** — memakai Smart Links, Slack/Teams, IDE, Figma, dan panel app di issue; mengekspor hasil pencarian.
- **Product Owner/PM** — memakai ekspor CSV/Sheets dan ekspor Plans, menautkan dokumen requirements dan desain.
- **Stakeholder/Viewer** — menerima unfurl dan notifikasi channel, membaca ekspor/printable, melihat makro data di dokumen.
- **Pengguna Tim Bisnis** — memindahkan proyek dari Trello, Asana, atau monday.com; memakai Outlook/Gmail untuk membuat work item dari email.
- **Integrator/Developer API** — membangun konektor iPaaS, app dengan extension point, entity properties, event trigger, serta importer JSON.
- **Tamu/Kolaborator Eksternal** — mengirim feedback lewat issue collector; terdampak batas data di channel Slack Connect.
- **Agent Layanan** / **Manajer Layanan/ITSM** — rujukan silang integrasi helpdesk/CRM, Opsgenie, Statuspage, dan Assets.
- **Product Manager Discovery** — rujukan silang tautan idea ke delivery.

## User story modul

- Sebagai Admin Situs, saya ingin semua request keluar (webhook, Smart Link, importer URL) melewati pengaman SSRF dengan allowlist yang bisa saya atur agar fitur integrasi tidak bisa dipakai untuk menyerang jaringan internal (INT-001).
- Sebagai Developer/Anggota Tim, saya ingin tautan yang saya tempel berubah menjadi Smart Link dengan judul dan status live sesuai izin pembaca agar konteks terlihat tanpa membuka tab baru (INT-002).
- Sebagai Admin Proyek/Project Lead, saya ingin mengimpor CSV dari alat lama dengan mapping kolom, nilai, hierarki, dan validasi dry-run agar migrasi tim berlangsung tanpa data rusak (INT-007 s.d. INT-014).
- Sebagai Pengguna Tim Bisnis, saya ingin memindahkan board Trello atau proyek Asana/monday.com ke proyek baru beserta status, assignee, dan komentar agar tim saya cepat beralih (INT-017, INT-018).
- Sebagai Integrator/Developer API, saya ingin importer JSON yang memuat history, worklog, dan link agar migrasi dari instans lain lengkap dan dapat diulang (INT-016).
- Sebagai Product Owner/PM, saya ingin mengekspor hasil pencarian ke CSV yang bisa diimpor ulang dan aman dari formula injection agar analisis di spreadsheet tidak merusak data (INT-021, INT-015).
- Sebagai Admin Situs, saya ingin membuat backup situs dengan atau tanpa lampiran dan merestorenya di situs lain agar data dapat dipulihkan saat bencana atau pindah instans (INT-024, INT-025).
- Sebagai Scrum Master/Team Lead, saya ingin channel Slack tim menerima update work item proyek beserta tombol transition dan assign agar diskusi dan tindakan terjadi di satu tempat (INT-030, INT-031, INT-034).
- Sebagai Developer/Anggota Tim, saya ingin membuat work item dari pesan Slack atau Teams dengan isi yang terisi otomatis agar laporan bug dari chat tidak hilang (INT-033, INT-036).
- Sebagai Stakeholder/Viewer, saya ingin menarik data work item ke Google Sheets dengan refresh terjadwal agar laporan saya selalu mutakhir (INT-038).
- Sebagai Developer/Anggota Tim, saya ingin menautkan desain Figma dan dokumen requirements ke work item agar implementasi merujuk sumber yang benar (INT-043, INT-046, INT-048).
- Sebagai Admin Situs, saya ingin meninjau scope dan egress sebuah app sebelum memasangnya serta mencabut aksesnya per proyek agar data organisasi tetap terkendali (INT-062, INT-065).
- Sebagai Integrator/Developer API, saya ingin extension point panel issue, custom field type, dan workflow rule dengan kontrak versi agar app internal kami dapat memperluas produk tanpa fork (INT-069, INT-070, INT-071).
- Sebagai Integrator/Developer API, saya ingin entity properties yang dapat diindeks ke JQL dan event trigger app agar data integrasi dapat dicari dan bereaksi terhadap perubahan (INT-077, INT-080).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 8 |
| P1 | 9 |
| P2 | 26 |
| P3 | 37 |
| **Total** | **80** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 1 |
| Fase 3 | 19 |
| Fase 4 | 51 |
| Fase 6 | 1 |
| Fase 7 | 3 |
| Fase — (dc-only/deprecated, rujukan) | 5 |

Strategi urutan implementasi: modul ini tidak punya fitur Fase 1 karena MVP belum membutuhkan pertukaran data eksternal, tetapi Fase 2 wajib menghadirkan SSRF guard (INT-001) lebih dulu karena semua jalur keluar berikutnya (webhook, Smart Link, importer URL lampiran) bergantung padanya. Fase 3 dibuka dengan importer CSV inti (INT-007 s.d. INT-011, INT-014) dan ekspor CSV round-trip (INT-021), selaras dengan pemetaan fase (impor/ekspor di Fase 3), UI ekspor SRC-050, dan entry point admin ADM-038, agar tim dapat bermigrasi masuk dan keluar dengan satu format; Fase 3 lalu melengkapi impor (komentar, worklog, sprint, JSON, Trello, Asana/monday.com), NFR operasi massal, ekspor legacy, backup situs, Smart Links, entity properties, dan pola iPaaS yang hanya menuntut REST API dan webhook publik modul 23. Fase 4 berisi ekosistem: konektor chat dan office, padanan Confluence/Figma/Loom, manajemen app, serta extension point yang dibangun mengikuti strategi INT-060 (mulai dari REST + webhook + entity properties, baru kemudian UI extension point ber-iframe sandbox). Integrasi lintas produk JSM/JPD mengikuti fase produknya (Fase 6-7), sementara importer legacy, JCMA, Project import DC, dan Data Pipeline hanya rujukan.

## Fitur

## Fondasi integrasi & tautan eksternal

### INT-001 SSRF & Outbound Request Guard

- **Deskripsi:** Pengaman terpusat untuk semua request HTTP keluar dari server (webhook, Automation "Send web request", resolver Smart Link, importer URL lampiran, avatar URL, konektor) agar tidak dapat diarahkan ke jaringan internal atau metadata cloud.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / platform
- **Aktor & izin:** Engineering dan Security (implementasi), Admin Situs (Administer Jira) untuk allowlist/denylist dan proxy; semua fitur yang memanggil URL eksternal wajib memakai klien ini.
- **Aturan perilaku (acceptance criteria):**
  1. Setelah resolusi DNS, request ke alamat privat, loopback, link-local, atau metadata cloud (127.0.0.0/8, 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, 169.254.0.0/16, ::1, fc00::/7, fe80::/10, termasuk 0.0.0.0 dan IPv4-mapped IPv6) ditolak dengan galat `outbound_blocked_private_address` dan tidak pernah membuka koneksi TCP.
  2. Proteksi DNS rebinding: hostname di-resolve satu kali, IP hasil validasi dipakai langsung oleh custom dialer, sehingga perubahan record DNS di antara validasi dan koneksi tidak berpengaruh.
  3. Hanya skema `http` dan `https` yang diterima (skema lain seperti `file`, `gopher`, `ftp` ditolak); port dibatasi ke 80/443 kecuali port tercantum di allowlist admin.
  4. Setiap request memiliki timeout koneksi dan total (default usulan 5 detik/30 detik), batas ukuran response (default usulan 10 MB, body dipotong dan request gagal dengan `outbound_response_too_large`), dan batas redirect (default 5) dengan validasi ulang alamat di setiap hop.
  5. Admin dapat mengelola allowlist/denylist domain (wildcard subdomain didukung); instans self-hosted dapat mengizinkan rentang jaringan internal secara eksplisit lewat konfigurasi environment, dan pengaturan ini dicatat di audit log.
  6. Egress proxy HTTP(S) opsional dapat dikonfigurasi; bila aktif semua request keluar melewati proxy dan validasi alamat tetap dijalankan terhadap tujuan akhir.
  7. Setiap request keluar dicatat (waktu, fitur pemanggil, host tujuan, IP hasil resolusi, status, durasi, alasan penolakan) tanpa menyimpan header otorisasi atau body sensitif, dan log dapat difilter admin.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Diimplementasikan sebagai paket bersama `pkg/netguard` yang dipakai webhook ([23 API Publik & Webhook](23-api-publik-webhook.md)), AUTO-054 di [20 Otomasi](20-otomasi.md), INT-002, importer (INT-012, INT-016, INT-017), dan konektor. Rincian NFR keamanan lain di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

### INT-002 Smart Links & link previews (third-party accounts)

- **Deskripsi:** Tautan yang ditempel di description, comment, atau field rich text otomatis menjadi Smart Link (inline, card, atau embed) dengan metadata dari produk ini maupun layanan pihak ketiga yang terhubung. Modul ini memiliki registry resolver dan koneksi akun pihak ketiga; rendering editor ada di modul 07.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Semua pengguna (penulis dan pembaca konten), Admin Situs/Org admin (kebijakan preview pihak ketiga), Integrator (resolver tambahan dari app).
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia empat mode tampilan: URL biasa, inline (ikon + judul + status lozenge), card (judul, deskripsi, metadata, avatar), dan embed (iframe interaktif untuk penyedia yang mendukung seperti Figma, Loom, Google Docs, YouTube); penyedia tanpa dukungan embed hanya menawarkan URL/inline/card.
  2. Saat link dipilih di editor, toolbar menampilkan pilihan mode; untuk embed, ukuran (lebar/tinggi) dan perataan dapat diatur dan disimpan di dokumen.
  3. Registry resolver per penyedia: work item produk ini (key, summary, status, assignee), halaman/whiteboard dokumentasi, Loom, Trello, Bitbucket/GitHub/GitLab, Google Drive, OneDrive/SharePoint, Dropbox, Box, Figma, Slack; URL lain jatuh ke resolver generik oEmbed/OpenGraph.
  4. Untuk sumber privat, pengguna menekan "Connect account" dan mengotorisasi akun pihak ketiga lewat OAuth; sebelum terhubung link tampil berstatus "Connect to preview", dan bila sudah terhubung tetapi tak berhak tampil "Restricted".
  5. Hover di atas Smart Link menampilkan hover card berisi preview dan aksi cepat (Open, Copy link); untuk work item ditampilkan juga status terkini.
  6. Metadata di-resolve dengan kredensial dan izin pengguna yang melihat, bukan pengguna yang menempelkan link; cache resolusi dipartisi per pengguna atau hanya menyimpan metadata publik.
  7. Bila resolver gagal, timeout, atau diblokir INT-001, link dirender sebagai URL biasa tanpa galat yang mengganggu editor.
  8. Menempel URL work item produk ini otomatis menjadi inline lozenge dengan status live yang diperbarui saat status berubah.
  9. Web link (remote link) di work item menampilkan judul hasil resolve bila pengguna tidak mengisi judul manual.
  10. ⚠ Admin organisasi dapat menonaktifkan preview pihak ketiga per penyedia atau seluruhnya; saat nonaktif semua link penyedia itu dirender sebagai URL biasa.
  11. ⚠ App dapat mendaftarkan resolver Smart Link tambahan untuk pola URL tertentu lewat manifest, dengan prioritas di bawah resolver bawaan.
  12. Di email notifikasi Smart Link dirender sebagai link biasa berjudul; di aplikasi mobile/web responsif mode embed turun menjadi card.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rendering dan perilaku editor di EDT-020/EDT-021 ([07 Editor Konten Kaya](07-editor-konten-kaya.md)) dan NAV-017 ([21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md)); token OAuth pihak ketiga disimpan terenkripsi di `IntegrationConnection`. Semua fetch lewat INT-001.

### INT-003 Generic iPaaS integrations (Zapier, Make, Power Automate, Workato, n8n)

- **Deskripsi:** Pola integrasi ke ribuan aplikasi lewat platform iPaaS yang memakai REST API, webhook, dan OAuth produk ini. Konektor iPaaS dibuat pihak iPaaS; tanggung jawab produk ini adalah menyediakan API yang ramah konektor.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, Admin Situs (registrasi OAuth client, kebijakan token); aksi berjalan dengan izin akun yang mengotorisasi.
- **Aturan perilaku (acceptance criteria):**
  1. API mendukung trigger polling: work item baru berbasis JQL (`created >= <cursor>`), work item ter-update (`updated >= <cursor>` dengan urutan stabil), comment baru, dan versi baru, dengan cursor yang tidak melewatkan item saat timestamp sama.
  2. Instant trigger didukung lewat registrasi webhook dinamis via REST (buat, perpanjang, hapus) dengan filter JQL dan masa berlaku; webhook kedaluwarsa dihapus otomatis.
  3. Tersedia aksi API: create/update work item, add comment, transition, add attachment, find work item (JQL/key), dan find user (email/nama sesuai privasi).
  4. Autentikasi menerima OAuth 2.0 (3LO) dengan scope dan API token + email (basic auth); token yang dicabut langsung menghasilkan 401.
  5. Konektor milik vendor iPaaS (mis. Power Automate dari Microsoft) berada di luar cakupan; produk ini menerbitkan spesifikasi OpenAPI agar konektor kustom dapat dibuat.
  6. Respons rate limit memakai 429 dengan header `Retry-After`, dan semua endpoint daftar memakai pagination konsisten (startAt/maxResults atau cursor); create mendukung idempotency key agar retry iPaaS tidak membuat duplikat.
  7. Alternatif tanpa iPaaS tersedia lewat Automation "Send web request" (AUTO-054) dan trigger "Incoming webhook" (AUTO-025).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kontrak API, OAuth, webhook, dan rate limit dimiliki [23 API Publik & Webhook](23-api-publik-webhook.md); modul ini hanya menetapkan kebutuhan agar ramah iPaaS.

### INT-004 Support & CRM tool integrations (Zendesk, Salesforce) – integration pattern

- **Deskripsi:** Pola integrasi helpdesk/CRM untuk menautkan ticket/case ke work item, membuat work item dari ticket, dan menyinkronkan comment/status. Konektor spesifik dibuat vendor; produk ini menyediakan primitif generik "external ticket link".
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Agent Layanan/support agent (di sisi helpdesk), Developer/Anggota Tim, Integrator; aksi dari helpdesk memakai akun integrasi dengan izin Browse/Create/Add comments.
- **Aturan perilaku (acceptance criteria):**
  1. Work item menampilkan panel "External tickets" berisi daftar ticket/case yang ter-link (sistem, ID, judul, status, pelanggan) beserta jumlahnya sebagai indikator dampak ke customer; jumlah dapat dicari via JQL.
  2. Dari sisi helpdesk/CRM, integrasi dapat membuat work item (project, type, summary, description) atau menautkan ke work item yang sudah ada lewat key/pencarian.
  3. Sinkronisasi comment dua arah dapat diaktifkan: comment produk ini dikirim sebagai internal note atau sebaliknya, dengan kontrol per comment (tanda "sync" atau visibility) agar tidak semua diskusi internal terkirim.
  4. Saat status work item berubah atau resolved, integrasi menerima pemberitahuan (webhook) sehingga ticket dapat diperbarui otomatis.
  5. Admin integrasi dapat mengatur mapping field ticket ke field work item (mis. prioritas, produk, organisasi) di konfigurasi konektor.
  6. Mekanisme dasar yang wajib tersedia: REST API, webhook, remote links, entity properties, dan OAuth 2.0 (3LO); tidak ada konektor Zendesk/Salesforce bawaan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konektor Zendesk, Salesforce, HubSpot, Intercom, Freshdesk adalah ekosistem (di luar cakupan). Remote link di ISS-030 ([06 Work Item Inti](06-work-item-inti.md)), entity properties INT-077.

### INT-005 Application links

- **Deskripsi:** Konfigurasi "linked instances" untuk menghubungkan produk ini dengan instans lain (produk ini di situs berbeda, wiki/dokumentasi self-hosted, CI lama) agar remote link lintas instans dan makro dapat dipakai.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman admin Products > Application links menerima URL aplikasi; sistem memeriksa ketersediaan dan mendeteksi tipe aplikasi secara otomatis bila mungkin sebelum link dibuat.
  2. Autentikasi link memakai OAuth 2.0 (atau OAuth 1.0a untuk kompatibilitas, dengan impersonation atau 2-legged tanpa impersonation) dengan konfigurasi incoming dan outgoing yang terpisah, dan halaman menampilkan status konfigurasi per arah.
  3. Satu link per jenis aplikasi dapat ditandai primary dan dipakai sebagai default untuk remote link dan makro.
  4. Remote issue link dapat menunjuk work item di instans lain dan halaman dokumentasi self-hosted, dengan judul/status di-resolve lewat link tersebut.
  5. Daftar link menampilkan status koneksi (online, offline, autentikasi gagal) dan mendukung edit serta hapus dengan konfirmasi.
  6. Tipe aplikasi CI/code review lama (Bamboo, Fisheye/Crucible) hanya didukung sebagai link generik tanpa integrasi khusus.
  7. Aplikasi dan modul di dalam satu instans produk ini tersambung otomatis tanpa application link.
  8. Application link dipakai untuk remote link ke halaman wiki/dokumentasi eksternal, dev info dari server Git self-hosted ([24 Integrasi DevOps](24-integrasi-devops.md)), dan gadget lintas aplikasi.
  9. Link rusak menampilkan pesan troubleshooting (URL tidak terjangkau, kredensial ditolak, sertifikat tidak valid) dengan tombol "Test connection".
  10. Host application link otomatis ditambahkan ke allowlist egress INT-001 (padanan perilaku Data Center); penghapusan link mencabut entri allowlist tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** INT-005 adalah spesifikasi pemilik (source of truth) application links; ADM-032 ([22 Administrasi Sistem](22-administrasi-sistem.md)) hanya titik masuk halaman admin dan merujuk ke sini, dengan tier yang disamakan (free-standard). Semua koneksi keluar lewat INT-001.

### INT-006 Issue collector

- **Deskripsi:** Widget feedback yang disematkan di website eksternal untuk mengumpulkan bug/feedback yang langsung menjadi work item, hanya untuk proyek company-managed.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Proyek (Administer Projects) membuat collector; Tamu/Kolaborator Eksternal (pengunjung website) mengirim feedback; default reporter harus punya izin Create issues.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Issue collectors (hanya company-managed) menyediakan daftar dan tombol buat collector; di team-managed menu tidak ditampilkan.
  2. Admin memilih work type dan reporter default, serta opsi mencocokkan email pengirim dengan pengguna yang ada sebagai reporter.
  3. Template form tersedia: "Got feedback?", "Raise a bug", atau custom dengan pemilihan field dari layar create.
  4. Trigger dapat berupa tombol bawaan (posisi dan teks dapat diatur) atau fungsi JavaScript custom yang dipanggil situs pemilik.
  5. Info browser/environment (user agent, resolusi layar) dan URL halaman dikumpulkan otomatis ke field Environment.
  6. Sistem menghasilkan snippet embed script dengan ID collector; snippet dimuat dari domain produk ini dan menghormati CSP/CORS yang dikonfigurasi.
  7. Collector dapat di-enable/disable; daftar menampilkan jumlah work item yang terkumpul dan tren mingguan.
  8. Pengirim anonim memakai default reporter; bila default reporter kehilangan izin Create issues, pengiriman ditolak dengan pesan ramah dan admin proyek diberi peringatan.
  9. ⚠ Pengirim dapat melampirkan screenshot/file (batas ukuran mengikuti pengaturan lampiran) yang menjadi lampiran work item.
- **Team-managed vs company-managed:** Hanya company-managed; team-managed diarahkan ke formulir (modul 26) atau widget layanan JSM.
- **Catatan & ketergantungan:** Fitur legacy; padanan di PRJ-061 ([04 Proyek](04-proyek.md)), ISS-040 ([06 Work Item Inti](06-work-item-inti.md)), dan daftar global ADM-033. Endpoint submit wajib rate limit dan CAPTCHA opsional.

## Impor data

### INT-007 Import entry points & permissions (site-level vs project-level)

- **Deskripsi:** Jalur impor dengan aturan izin berbeda: External System Import untuk admin di level situs, importer CSV untuk pengguna non-admin ke proyek yang ada, dan impor saat membuat proyek. Semua jalur memakai mesin job impor yang sama.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira) untuk impor site-level; Admin Proyek dan pengguna dengan global permission Make bulk changes + Create issues di proyek tujuan untuk impor user-level.
- **Aturan perilaku (acceptance criteria):**
  1. Settings > System > Import and export > External system import hanya dapat diakses Administer Jira dan menampilkan sumber yang tersedia: CSV, JSON, Trello, Asana, monday.com; pengguna lain yang membuka URL-nya menerima 403.
  2. ⚠ Impor CSV user-level tersedia dari navigator/Filters lewat "Import work items from CSV" dan hanya ditampilkan bila pengguna punya global permission Make bulk changes dan Create issues di minimal satu proyek; proyek tujuan dibatasi ke proyek dengan izin Create issues.
  3. Importer user-level hanya mengimpor ke proyek yang sudah ada; opsi "Create new project" dan "Import as new custom field" tidak ditampilkan dan permintaan API yang memuatnya ditolak dengan 403.
  4. Importer site-level dapat membuat proyek baru, membuat custom field, dan mengatur mapping pengguna.
  5. ⚠ Alur create project menawarkan langkah opsional "Import data" dari Trello, Asana, monday.com, atau spreadsheet/CSV, yang meneruskan proyek baru sebagai tujuan.
  6. ⚠ Untuk proyek team-managed, impor CSV hanya dapat memetakan ke field yang ada di proyek; impor site-level ke team-managed diperlakukan sama (tanpa membuat custom field global).
  7. Impor berjalan asinkron sebagai `ImportJob` dengan progres realtime (baris diproses/total), ringkasan hasil, dan log per baris yang dapat diunduh; pengguna dapat menutup halaman dan kembali ke halaman riwayat impor.
  8. Tidak ada rollback otomatis; UI menampilkan peringatan untuk menguji di proyek uji, dan sebagai nilai tambah produk ini menyediakan aksi "Delete imported batch" yang menghapus work item yang dibuat job tersebut (butuh izin Delete issues) dalam 30 hari.
- **Team-managed vs company-managed:** Team-managed hanya ke field proyek yang ada; company-managed dapat memakai field global dan membuat custom field (site-level).
- **Catatan & ketergantungan:** Halaman admin di ADM-038 ([22 Administrasi Sistem](22-administrasi-sistem.md)); permission di [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md); job queue di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md). Repo saat ini belum memiliki job queue maupun proyek; tabel `import_jobs` dan `import_job_rows` adalah migrasi baru.

### INT-008 CSV importer – file upload & import settings

- **Deskripsi:** Langkah awal importer CSV: unggah file, muat konfigurasi tersimpan, encoding, delimiter, proyek tujuan, format tanggal, dan pengaturan pengguna.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (site-level), pengguna dengan izin impor user-level (INT-007).
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna mengunggah tepat satu file CSV; baris pertama wajib berupa header kolom, dan file tanpa header atau kosong ditolak dengan galat `csv_missing_header`.
  2. Opsi "Use an existing configuration file" menerima file JSON konfigurasi dari impor sebelumnya dan mengisi otomatis semua mapping; kolom yang tidak ditemukan di file baru ditandai peringatan.
  3. Bagian Advanced menyediakan file encoding (default UTF-8; BOM UTF-8 dibuang otomatis) dan delimiter (koma, titik koma, tab, atau satu karakter lain).
  4. Proyek tujuan dipilih dari proyek yang ada atau, di importer site-level, "Create new" dengan nama, key, dan tipe/template; validasi key mengikuti aturan proyek.
  5. ⚠ Opsi email suffix untuk pengguna baru (legacy) tidak dibangun; nilai pengguna yang tidak cocok ditangani mapping pengguna INT-010.
  6. Date format memakai pola gaya Java SimpleDateFormat (mis. `dd/MMM/yy h:mm a`) yang berlaku untuk semua kolom tanggal, dengan pratinjau hasil parse baris pertama.
  7. Parsing mengikuti RFC 4180: nilai berkutip, baris baru di dalam sel, dan tanda kutip ganda yang di-escape (`""`) diproses benar.
  8. ⚠ Batas default usulan 50 MB atau 250.000 baris per file; file melebihi batas ditolak dengan pesan untuk memecah file.
  9. Header duplikat diperbolehkan dan diperlakukan sebagai kolom multi-value untuk field yang sama.
- **Team-managed vs company-managed:** Opsi "Create new" membuat proyek company-managed atau team-managed sesuai template yang dipilih.
- **Catatan & ketergantungan:** Parsing streaming wajib (INT-015). File unggahan disimpan sementara di object storage dan dihapus setelah job selesai + 7 hari.

### INT-009 CSV importer – field mapping & custom field creation

- **Deskripsi:** Memetakan tiap kolom CSV ke field (system atau custom), membuat custom field baru saat impor, dan menandai kolom yang nilainya perlu dipetakan.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (termasuk pembuatan custom field), pengguna impor user-level (mapping ke field yang ada saja).
- **Aturan perilaku (acceptance criteria):**
  1. Setiap kolom memiliki dropdown field tujuan atau "Don't map this field"; kolom yang tidak dipetakan diabaikan.
  2. Checkbox "Map field value" per kolom mengaktifkan langkah mapping nilai (INT-010) untuk kolom tersebut.
  3. "Import as new custom field" hanya tersedia di importer site-level dengan pilihan tipe: text, paragraph, number, date, datetime, select single/multi, checkboxes, radio, cascading, user picker single/multi, labels, URL.
  4. Kolom yang dipetakan ke Summary wajib ada; bila kolom Issue Type tidak ada, work type default proyek dipakai dan hal ini ditampilkan di ringkasan.
  5. Field system yang dapat dipetakan: Summary, Description, Issue Type, Status, Priority, Resolution, Assignee, Reporter, Labels, Components, Affects Version/s, Fix Version/s, Due date, Created, Updated, Resolved, Environment, Security Level, Watchers, Original Estimate, Remaining Estimate, Time Spent, Parent, Issue Id, Parent Id, Sprint, Story Points/Story point estimate, Start date, Team.
  6. Beberapa kolom dengan nama sama dipetakan ke satu field multi-value (Labels, Components, Versions, Watchers, Comment, Attachment, Log Work, multi-select).
  7. Bila field tujuan tidak ada di context atau screen proyek/work type, UI menampilkan peringatan "field tidak akan terlihat" tetapi tetap mengizinkan impor.
  8. Untuk proyek team-managed, dropdown hanya berisi field yang ada di proyek tersebut.
  9. ⚠ Cascading select diimpor dengan format `parent->child`; nilai tanpa separator dianggap hanya parent.
  10. ⚠ Version dan Component yang belum ada dibuat otomatis di proyek tujuan (company-managed dan team-managed dengan fitur Releases/Components aktif) dan dicantumkan di ringkasan.
- **Team-managed vs company-managed:** Team-managed terbatas pada field lokal proyek; company-managed dapat memakai field global dan membuat custom field baru.
- **Catatan & ketergantungan:** Definisi field dan context di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md).

### INT-010 CSV importer – value mapping (status, type, priority, users)

- **Deskripsi:** Memetakan nilai CSV ke nilai yang ada untuk status workflow, work type, priority, resolution, dan pengguna.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pengguna impor user-level.
- **Aturan perilaku (acceptance criteria):**
  1. Tabel mapping menampilkan setiap nilai unik kolom beserta jumlah kemunculannya; tiap nilai dapat dipetakan ke nilai yang ada atau dibiarkan apa adanya (dicocokkan berdasarkan nama tanpa membedakan huruf besar/kecil).
  2. Status: work item langsung ditempatkan di status yang dipetakan tanpa menjalankan transition, sehingga condition, validator, dan post function tidak dieksekusi; status yang tidak ada di workflow proyek menjadi galat baris.
  3. Resolution dipetakan ke resolution yang ada; work item di kategori status Done tanpa resolution diberi peringatan di log, dan resolution diisi otomatis bila proyek memakai resolusi default.
  4. Work type dipetakan ke work type scheme proyek; baris sub-task tanpa parent valid ditolak.
  5. Priority dipetakan ke priority scheme proyek; nilai tidak dipetakan memakai priority default.
  6. ⚠ Pengguna (Assignee, Reporter, Watchers, user picker) dicocokkan lewat email atau account ID; pengguna yang tidak ditemukan dikosongkan, atau Reporter diganti pengguna pengimpor, dan setiap penggantian dicatat di log.
  7. ⚠ Importer tidak membuat akun pengguna baru; admin diarahkan ke undangan/bulk import pengguna (USR-010) sebelum impor.
  8. ⚠ Opsi select/checkbox/radio yang belum ada dibuat otomatis bila pengimpor punya izin admin field; bila tidak, baris diimpor tanpa nilai tersebut dengan peringatan.
  9. Sel kosong diperlakukan sebagai nilai kosong, kecuali field punya default value yang kemudian diterapkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber status (workflow proyek team-managed vs scheme company-managed).
- **Catatan & ketergantungan:** Workflow di [09 Workflow](09-workflow-status-transisi.md); pengguna di [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md).

### INT-011 CSV importer – hierarchy, sub-tasks & issue links

- **Deskripsi:** Mengimpor hierarki (epic > story > sub-task) dan relasi antar work item dari kolom Issue Id, Parent Id, Parent, dan kolom link.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pengguna impor user-level (butuh Link issues untuk kolom link).
- **Aturan perilaku (acceptance criteria):**
  1. Kolom "Issue Id" adalah ID sementara unik per baris di file (bukan key); duplikat Issue Id menjadi galat validasi.
  2. Kolom "Parent Id" merujuk Issue Id baris lain sebagai parent untuk sub-task maupun level hierarki lain.
  3. ⚠ Kolom "Parent" menerima key work item yang sudah ada di produk ini; key yang tidak ada atau tidak dapat dilihat pengimpor menjadi galat baris.
  4. ⚠ Kolom legacy Epic Link/Epic Name diterima dan dikonversi menjadi relasi parent untuk kompatibilitas file lama.
  5. Sub-task tanpa parent yang valid ditolak dan dicatat di log.
  6. ⚠ Referensi parent diselesaikan dalam dua tahap (buat semua baris, lalu tautkan) sehingga urutan baris acak tetap menghasilkan hierarki benar.
  7. Kolom `Link "<link type>"` (mis. `Link "Blocks"`, `Link "Relates"`) berisi Issue Id baris lain dan menghasilkan issue link dengan arah outward dari baris tersebut; link type yang tidak ada menjadi peringatan.
  8. ⚠ Kolom link juga menerima key work item yang sudah ada untuk menautkan ke data lama.
  9. Kedalaman hierarki mengikuti hierarki work type situs; level di atas Epic hanya dapat diimpor bila hierarki tersebut sudah dikonfigurasi (fitur premium).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hierarki di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) (CFG-006, CFG-007, CFG-024); link di ISS-029 ([06 Work Item Inti](06-work-item-inti.md)).

### INT-012 CSV importer – comments, worklogs, attachments & time tracking columns

- **Deskripsi:** Format kolom khusus untuk mengimpor comment, worklog, lampiran lewat URL, dan estimasi waktu.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pengguna impor user-level.
- **Aturan perilaku (acceptance criteria):**
  1. Kolom "Comment" dapat berulang; format opsional `tanggal;user;isi` mempertahankan tanggal dan author asli, sedangkan isi tanpa titik koma dianggap ditulis pengimpor saat impor.
  2. ⚠ Kolom "Log Work" berulang berformat `comment;tanggal;user;durasi_detik`; durasi tidak valid menjadi galat baris.
  3. Kolom "Attachment" berisi URL yang dapat diakses importer, dengan format opsional `tanggal;user;nama file;URL`; unduhan melewati INT-001 dan batas ukuran lampiran.
  4. Lampiran dari path lokal (mis. `C:\` atau `/home/...`) ditolak dengan pesan bahwa hanya URL yang didukung.
  5. Original Estimate, Remaining Estimate, dan Time Spent dibaca dalam detik.
  6. ⚠ Teks berformat wiki markup dikonversi ke rich text dokumen; Markdown sederhana juga dikenali.
  7. Tanggal di comment, worklog, dan lampiran mengikuti date format yang dipilih di INT-008.
  8. ⚠ Pengguna yang tidak cocok di comment/worklog digantikan pengguna pengimpor dan nama asli disisipkan sebagai prefix teks.
  9. Change history tidak dapat diimpor lewat CSV; UI mengarahkan ke JSON importer (INT-016).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Format diselaraskan dengan ekspor CSV INT-021 agar round-trip; estimasi dan worklog di TIME-045 ([14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md)).

### INT-013 CSV importer – sprints, dates & special fields

- **Deskripsi:** Aturan khusus untuk mengimpor sprint, tanggal sistem, story points, label, security level, dan field agile.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pengguna impor user-level (sprint baru butuh izin Manage sprints di board).
- **Aturan perilaku (acceptance criteria):**
  1. Kolom Sprint berisi nama sprint; kolom berulang mencatat riwayat sprint dan sprint terakhir dianggap sprint aktif/terkini work item.
  2. ⚠ Sprint yang belum ada dibuat otomatis dan dikaitkan ke board proyek tujuan dengan status future (atau closed bila semua work item di sprint itu sudah Done).
  3. Created, Updated, Resolved, dan Due date mengikuti date format; Created boleh berisi tanggal historis dan tidak diganti waktu impor.
  4. Labels tidak boleh berisi spasi (spasi diganti garis bawah dengan peringatan); banyak label lewat kolom berulang.
  5. Story Points (company-managed) dan Story point estimate (team-managed) diperlakukan sebagai field berbeda dan dropdown hanya menawarkan yang berlaku di proyek tujuan.
  6. Start date, Due date, dan Team dipetakan hanya bila field ada di context proyek.
  7. Security Level dipetakan ke issue security scheme proyek; level yang tidak ada menjadi galat baris.
  8. Rank tidak dapat diimpor; urutan backlog mengikuti urutan baris di file.
  9. ⚠ Key dari sistem lain tidak dipakai sebagai key; key baru dibuat berurutan, dan key asal disimpan di field "External key"/entity property untuk penelusuran.
- **Team-managed vs company-managed:** Story Points vs Story point estimate berbeda per tipe proyek (butir 5).
- **Catatan & ketergantungan:** Sprint dan board di [11 Backlog & Sprint](11-backlog-sprint.md); issue security di [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md).

### INT-014 CSV importer – validation, saved configuration, logs & updating existing issues

- **Deskripsi:** Dry-run validasi sebelum impor, menyimpan konfigurasi mapping, log hasil, dan pembaruan work item yang sudah ada lewat CSV.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pengguna impor user-level (update butuh Edit issues pada work item tujuan).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Validate" menjalankan dry-run penuh tanpa menulis data dan menampilkan daftar galat dan peringatan per baris; impor tetap dapat dimulai meski ada peringatan.
  2. "Save configuration" mengunduh JSON berisi mapping kolom, mapping nilai, proyek, date format, encoding, dan delimiter yang dapat dimuat di INT-008.
  3. "Begin import" berjalan asinkron dengan progres; ringkasan menampilkan jumlah work item dibuat, diperbarui, gagal, dan peringatan.
  4. Log detail per baris (nomor baris, status, key hasil, pesan) dapat diunduh sebagai CSV/teks selama minimal 30 hari.
  5. ⚠ Mode update: bila kolom Issue Key dipetakan ke key yang sudah ada, baris tersebut memperbarui field yang dipetakan alih-alih membuat baru; status tidak diubah lewat update, dan key milik proyek lain atau tak berizin Edit menjadi galat baris.
  6. Impor bersifat parsial: baris gagal dilewati, baris lain tetap dibuat (tidak transaksional), dan ringkasan menyebut baris gagal.
  7. Tidak ada undo bawaan selain "Delete imported batch" (INT-007); pembersihan lain lewat bulk delete.
  8. ⚠ Notifikasi email dan event otomasi "created" untuk work item hasil impor ditekan secara default; admin dapat mengaktifkan opsi "Kirim notifikasi" sebelum mulai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Event impor tetap ditulis ke audit log ([22 Administrasi Sistem](22-administrasi-sistem.md)) dan event domain `import.completed` ([28 Model Data Inti](28-model-data-inti.md)).

### INT-015 Import/Export Performance & Data Formats

- **Deskripsi:** Kebutuhan non-fungsional untuk operasi data massal: impor CSV, ekspor navigator, dan migrasi dari instans lain, agar operasi besar tidak mengganggu pengguna aktif dan file aman dibuka di spreadsheet.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Situs, Engineering.
- **Aturan perilaku (acceptance criteria):**
  1. Impor CSV diproses streaming (baris dibaca bertahap, memori proses tidak tumbuh linear dengan ukuran file), dengan validasi per baris, dry-run, dan laporan galat; uji beban 250.000 baris tidak melebihi batas memori worker yang dikonfigurasi.
  2. ⚠ Ekspor navigator (CSV all/current fields, XML, Word, HTML) dijalankan streaming untuk hasil kecil dan sebagai job asinkron untuk hasil besar; batas default sinkron 1.000 baris dan batas job 10.000 baris dapat diatur admin.
  3. Output CSV memakai UTF-8; opsi "BOM untuk Excel" menambahkan BOM UTF-8 agar karakter non-ASCII terbaca benar di Excel.
  4. Impor dari produk lain (CSV/JSON/API) menyediakan langkah pemetaan pengguna dan field yang konsisten antarformat, dan konfigurasi mapping dapat dipakai ulang.
  5. Job impor/ekspor memakai job queue dengan rate limit per tenant dan konkurensi maksimum (usulan 2 job impor per situs); job tambahan masuk antrean dengan posisi yang terlihat.
  6. Ekspor CSV mencegah formula injection: nilai yang diawali `=`, `+`, `-`, `@`, tab, atau carriage return diberi prefix tanda kutip tunggal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Infrastruktur job, object storage, dan observability di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md); UI ekspor di [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md).

### INT-016 JSON importer

- **Deskripsi:** Importer site-level dengan file JSON terstruktur berisi users, projects, versions, components, dan work item lengkap dengan history, comments, worklogs, attachments, custom fields, dan links. Dijadikan format migrasi kanonik produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira), Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Root object memuat `users`, `links`, dan `projects`; key tidak dikenal diabaikan dengan peringatan dan struktur tidak valid ditolak sebelum job dimulai.
  2. Setiap project memuat name, key, type, template, description, lead, versions (name, released, releaseDate), dan components; proyek dengan key yang sudah ada dapat digabung bila admin memilih "Merge into existing".
  3. Setiap work item memuat externalId, key (opsional), summary, description, issueType, status, priority, resolution, reporter, assignee, created, updated, resolutionDate, duedate, labels, watchers, components, affectedVersions, fixedVersions, environment.
  4. `customFieldValues` berisi fieldName, fieldType (key tipe field), dan value; field yang belum ada dibuat sesuai fieldType.
  5. `history` berisi perubahan field dengan author dan timestamp, dan diimpor sebagai riwayat aktivitas asli (bukan sebagai comment).
  6. `comments` (body, author, created) dan `worklogs` (author, comment, startDate, timeSpent dalam durasi ISO 8601 seperti `PT2H30M`) diimpor dengan author dan tanggal asli.
  7. `attachments` berisi name, attacher, created, uri, description; uri diunduh lewat INT-001.
  8. `links` berisi name (link type atau `sub-task-link`), sourceId, dan destinationId berdasarkan externalId.
  9. ⚠ `users` berisi name, fullname, email, groups, active; pengguna yang belum ada tidak dibuat otomatis melainkan dipetakan ke pengguna yang ada atau ditandai "inactive placeholder" sesuai pilihan admin.
  10. Semua tanggal memakai ISO 8601; galat validasi dan impor dicatat di log job per entitas.
  11. ⚠ Batas default usulan 500 MB per file JSON; file lebih besar ditolak dengan anjuran memecah per proyek.
- **Team-managed vs company-managed:** Proyek baru dibuat sesuai `type`/`template`; field lokal team-managed dibuat di proyek, bukan global.
- **Catatan & ketergantungan:** Dirancang round-trip dengan exporter JSON migrasi (lihat Pertanyaan terbuka). Skema mengikuti entitas [28 Model Data Inti](28-model-data-inti.md).

### INT-017 Trello importer

- **Deskripsi:** Mengimpor board Trello ke proyek: kartu menjadi work item dan list menjadi status, termasuk label, member, checklist, dan comment.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, pembuat proyek (Create projects) saat alur create project.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna mengotorisasi akun Trello (OAuth) lalu memilih satu board dari daftar board yang dapat diaksesnya.
  2. Proyek tujuan dipilih dari yang ada atau dibuat baru dalam alur yang sama.
  3. Setiap list Trello dipetakan ke status proyek; list tanpa mapping memakai status awal workflow.
  4. Opsi "Sertakan kartu archived" default nonaktif; bila aktif kartu archived diimpor dengan status Done.
  5. Label Trello menjadi label; member kartu menjadi assignee (member pertama) dan watcher (lainnya), dicocokkan lewat email.
  6. ⚠ Checklist menjadi sub-task per item (default) atau daftar centang di description sesuai pilihan pengguna.
  7. ⚠ Comment dan lampiran kartu ikut diimpor; lampiran berupa tautan eksternal menjadi web link.
  8. Due date kartu menjadi Due date.
  9. ⚠ Importer tersedia di External System Import dan di langkah "Import data" alur create project.
- **Team-managed vs company-managed:** Tidak ada perbedaan; proyek baru default team-managed bisnis.
- **Catatan & ketergantungan:** Berbeda dari integrasi Trello Power-Up (INT-049). Token Trello dibuang setelah job selesai.

### INT-018 Asana & monday.com importers (business projects)

- **Deskripsi:** Importer untuk memindahkan proyek Asana atau board monday.com ke proyek, terutama template bisnis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pembuat proyek (Create projects), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Pengguna menghubungkan akun Asana atau monday.com lewat OAuth atau API token; token disimpan terenkripsi hanya selama job.
  2. Pengguna memilih workspace dan project (Asana) atau board (monday.com).
  3. Section (Asana) atau group/kolom status (monday.com) dipetakan ke status proyek.
  4. Task dan subtask menjadi work item dan sub-task dengan relasi parent terjaga.
  5. Assignee dicocokkan lewat email; due date, description, dan comment ikut diimpor.
  6. ⚠ Custom field sederhana (teks, angka, tanggal, dropdown) dipetakan ke custom field; tipe kompleks dilewati dengan peringatan.
  7. Importer diakses dari alur create project ("Import data") dan External System Import.
  8. Sumber lain (ClickUp, Smartsheet, Excel) diarahkan ke impor CSV (INT-008) dengan panduan format.
- **Team-managed vs company-managed:** Tidak ada perbedaan; target utama proyek bisnis ([26 Kerja Bisnis & Formulir](26-kerja-bisnis-formulir-persetujuan.md)).
- **Catatan & ketergantungan:** ⚠ Daftar sumber importer SaaS harus ditinjau ulang menjelang implementasi karena API vendor berubah.

### INT-019 GitHub Issues importer

- **Deskripsi:** Importer legacy untuk memindahkan GitHub Issues (label, milestone, comment). Tidak dibangun sebagai importer khusus; jalurnya ekspor dari GitHub ke CSV lalu INT-008.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tidak ada koneksi token GitHub untuk impor; dokumentasi menjelaskan cara mengekspor repository issues ke CSV.
  2. Panduan mapping: label menjadi work type atau label.
  3. Panduan mapping: milestone menjadi Fix Version.
  4. Panduan mapping: open/closed menjadi status/resolution.
  5. Panduan mapping: comment dan author dicocokkan ke pengguna lewat email.
  6. ⚠ Importer GitHub khusus berada di luar cakupan; integrasi GitHub untuk development info ada di modul 24.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; lihat [24 Integrasi DevOps](24-integrasi-devops.md).

### INT-020 Legacy tracker importers (Bugzilla, FogBugz, Mantis, Redmine, Trac, Pivotal Tracker)

- **Deskripsi:** Importer legacy untuk issue tracker lama lewat koneksi database atau API. Tidak dibangun; dicatat sebagai rujukan dan jalur alternatifnya CSV/JSON.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Bugzilla (koneksi database, mapping product, component, versi, status/resolution, pengguna, comment, lampiran) tidak didukung langsung; dokumentasi memberi skrip konversi ke JSON INT-016 sebagai saran komunitas.
  2. FogBugz (On Demand atau database, mapping project/area/milestone) tidak didukung langsung.
  3. Mantis (koneksi database, category ke component) tidak didukung langsung.
  4. Redmine (REST API, tracker ke work type) tidak didukung langsung.
  5. Trac (environment Trac) tidak didukung langsung.
  6. Pivotal Tracker tidak relevan karena layanannya ditutup pada 2025.
  7. Alur mapping project, field, nilai, dan pengguna dengan konfigurasi tersimpan disediakan oleh importer CSV/JSON generik.
  8. ⚠ Importer ini dianggap hanya tersisa di edisi Data Center sehingga tidak termasuk paritas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; lihat Di luar cakupan.

## Ekspor data

### INT-021 Issue search export – CSV / Excel CSV

- **Deskripsi:** Format dan aturan ekspor hasil pencarian (JQL/filter/list) ke CSV yang kompatibel dengan Excel, dengan pilihan semua field atau field yang sedang tampil. Modul ini memiliki format file dan kompatibilitas round-trip; UI menu ekspor dimiliki SRC-050.
- **Prioritas / Fase / Tier:** P0 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat hasil pencarian, termasuk anonim di proyek publik; dapat dibatasi kebijakan keamanan data (SRC-058).
- **Aturan perilaku (acceptance criteria):**
  1. Menu Export menyediakan "Export Excel CSV (all fields)" dan "Export Excel CSV (current fields)"; current fields mengikuti kolom dan urutan yang tampil.
  2. ⚠ Delimiter dapat dipilih: koma (default), titik koma, pipe, atau caret.
  3. ⚠ Batas default 1.000 work item per ekspor sinkron; hasil lebih besar dialihkan ke job asinkron (INT-015) atau UI menyarankan memecah lewat JQL (mis. per rentang `created`).
  4. Field multi-value (Labels, Components, Fix versions, Watchers, Sprint) diekspor sebagai kolom berulang dengan header sama sehingga dapat diimpor ulang lewat INT-009.
  5. Comment diekspor `tanggal;accountId;isi`, attachment `tanggal;accountId;nama file;URL`, dan worklog `comment;tanggal;accountId;durasi_detik` sehingga kompatibel dengan INT-012.
  6. Header custom field berformat `Custom field (Nama)` dan link berformat `Outward issue link (Type)`/`Inward issue link (Type)`.
  7. Kolom ID tambahan disertakan pada all fields: Issue id, Parent id, Parent summary, Project key, Project type, Project lead.
  8. Tanggal ditulis mengikuti format dan zona waktu preferensi pengguna yang mengekspor.
  9. Hanya work item dan field yang boleh dilihat pengguna yang diekspor; comment ber-visibility terbatas hanya disertakan bila pengguna berhak.
  10. ⚠ List view proyek software dan bisnis menyediakan ekspor CSV yang sama dengan kolom view aktif.
  11. Menu Export menyertakan "Export to Google Sheets" yang membuka add-on INT-038 dengan query terkait.
- **Team-managed vs company-managed:** Field lokal team-managed diekspor dengan nama field; bila bentrok nama ditambah label proyek.
- **Catatan & ketergantungan:** UI di SRC-050 dan SRC-055 ([13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md)); formula injection dan BOM di INT-015. Keputusan fase: walau P0 di inventaris, ekspor CSV ditempatkan di Fase 3 bersama importer CSV (INT-008 s.d. INT-014) dan UI SRC-050 agar format round-trip dirilis serentak.

### INT-022 Issue & search exports: XML, Word, Printable & RSS

- **Deskripsi:** Format ekspor legacy untuk hasil pencarian dan work item tunggal: XML (berbasis RSS), Word (HTML .doc), tampilan cetak, dan feed RSS/Atom.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat work item; feed memerlukan autentikasi (API token/sesi).
- **Aturan perilaku (acceptance criteria):**
  1. XML hasil pencarian berformat RSS 2.0 dengan elemen tambahan: key, summary, type, status, priority, resolution, assignee, reporter, created, updated, comments, custom fields, attachments, subtasks, issuelinks.
  2. ⚠ URL ekspor XML menerima parameter `tempMax` (maksimal 1.000) dan `field=` (berulang) untuk membatasi jumlah dan field.
  3. ⚠ Ekspor Word hasil pencarian menghasilkan dokumen HTML ber-ekstensi .doc berisi detail setiap work item.
  4. Printable menyediakan "Print list" (tabel kolom aktif) dan "Print details" (detail per work item) dengan stylesheet cetak.
  5. ⚠ Filter menyediakan feed RSS work item dan RSS comments; feed tanpa autentikasi valid mengembalikan 401 kecuali proyek publik.
  6. ⚠ Activity stream proyek/pengguna tersedia sebagai feed Atom dengan autentikasi yang sama.
  7. Work item tunggal menyediakan menu "..." > Export XML, Export Word, Print.
  8. Bagian Attachments menyediakan "Download all" yang menghasilkan file zip berisi lampiran yang dapat dilihat pengguna.
  9. PDF tidak dibangun sebagai format bawaan; pengguna memakai Print > Save as PDF browser.
  10. Semua format mengikuti izin viewer (work item, field, comment visibility).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** UI di SRC-051 s.d. SRC-053, SRC-056 ([13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md)) dan ISS-039 ([06 Work Item Inti](06-work-item-inti.md)). Ekspor PDF bermerek adalah ekosistem app.

### INT-023 Plans, Timeline & list exports

- **Deskripsi:** Ekspor dan berbagi data Plans, Timeline, dan list view dalam bentuk CSV, gambar, embed, atau link.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Pengguna dengan akses view plan (Plans premium); pengguna proyek untuk timeline/list (free-standard).
- **Aturan perilaku (acceptance criteria):**
  1. Plans > Share/Export > CSV menghasilkan file berisi hierarki (level dan parent), tanggal, tim, estimasi, dan field yang tampil di view.
  2. Plans > Share menyediakan embed ke halaman dokumentasi (iframe ber-autentikasi) dan link view yang membuka view yang sama.
  3. ⚠ Timeline plan dapat diekspor sebagai gambar PNG sesuai rentang waktu dan baris yang tampil.
  4. Timeline proyek dapat dibagikan/di-embed lewat Smart Link (INT-002).
  5. ⚠ List view menyediakan ekspor CSV sesuai kolom dan filter aktif (lihat INT-021 butir 10).
  6. Semua ekspor mengikuti filter, pengelompokan, dan view aktif; work item yang tidak dapat dilihat pengekspor tidak disertakan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur sumber di PLAN-013, PLAN-020, PLAN-054 ([18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md)); tier premium hanya untuk Plans.

## Backup, restore, migrasi & analitik

### INT-024 Site backup (Backup manager) & Atlassian Backup and Restore

- **Deskripsi:** Membuat backup seluruh data situs (dengan atau tanpa lampiran) untuk diunduh, plus backup otomatis terjadwal dan restore granular.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (hanya site admin yang dapat membuat dan mengunduh backup).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > System > Backup manager menyediakan tombol "Create backup" yang memulai job backup asinkron.
  2. Opsi "Include attachments, avatars, and logos" menentukan apakah file biner ikut; default nonaktif.
  3. ⚠ Backup manual dibatasi satu per 48 jam per situs; tombol dinonaktifkan dengan sisa waktu tertera selama masa tunggu.
  4. Halaman menampilkan progres dan, setelah selesai, link unduh .zip yang berlaku terbatas (usulan 7 hari) dan hanya dapat diakses site admin.
  5. Isi backup: work item, proyek, konfigurasi, pengguna/grup tanpa password hash maupun token, dan lampiran opsional; data app pihak ketiga tidak termasuk, dan isi ini didokumentasikan di halaman.
  6. ⚠ Backup otomatis terjadwal (harian/mingguan), retensi, dan restore granular per proyek tersedia di halaman organisasi sebagai fitur tier lebih tinggi.
  7. Pengguna tanpa peran site admin menerima 403 untuk halaman maupun URL unduhan, dan setiap pembuatan/pengunduhan dicatat di audit log.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman admin di ADM-035 s.d. ADM-037 ([22 Administrasi Sistem](22-administrasi-sistem.md)); backup infrastruktur (dump DB + object storage terenkripsi) di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

### INT-025 Jira Cloud backup restore & cloud-to-cloud migration

- **Deskripsi:** Mengimpor data dari instans lain: restore backup situs (menimpa seluruh data) atau memindahkan proyek tertentu antarsitus.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs, Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. Backup manager > Import cloud data menerima backup .zip produk ini dan menimpa SEMUA data situs tujuan.
  2. Sebelum restore, UI menampilkan peringatan bahwa data lama akan terhapus dan mewajibkan konfirmasi eksplisit (mengetik nama situs) oleh site admin.
  3. Pengguna dari backup digabungkan dengan pengguna organisasi berdasarkan email; pengguna baru dibuat nonaktif tanpa akses produk sampai admin memberi akses.
  4. Lampiran ikut direstore bila backup memuatnya; bila tidak, metadata lampiran ditandai "file tidak tersedia".
  5. Backup dari format pihak lain yang tidak dikenali ditolak; migrasi dari edisi self-hosted lama memakai importer JSON/REST (lihat INT-026).
  6. Migrasi antarsitus memilih proyek, pengguna, dan grup tertentu ke situs lain, dengan pengecekan konflik (key proyek, nama field, nama grup) sebelum mulai.
  7. ⚠ Data yang tidak ikut (data app pihak ketiga, sebagian Plans dan Product Discovery) ditampilkan di layar ringkasan sebelum konfirmasi.
  8. ⚠ Backup dan restore granular otomatis mengikuti INT-024 butir 6.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Ketersediaan migrasi antarsitus dan restore granular per tier menjadi keputusan entitlement ([00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md)). Admin di ADM-039.

### INT-026 Jira Cloud Migration Assistant (Server/DC → Cloud)

- **Deskripsi:** App migrasi dari edisi Server/Data Center ke Cloud untuk menilai dan memindahkan proyek, pengguna, grup, lampiran, dan data app. Tidak dibangun; padanan strategisnya adalah importer "dari Jira" berbasis REST/JSON.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs, Org admin.
- **Aturan perilaku (acceptance criteria):**
  1. Penilaian status migrasi per app terpasang tidak dibangun.
  2. Migration plan (situs tujuan, proyek, pengguna dan grup, lampiran saja, konfigurasi) digantikan konfigurasi job importer JSON INT-016.
  3. Pre-migration checks (konflik key proyek, email duplikat, versi, custom field ganda) disediakan oleh validasi dry-run importer.
  4. Migrasi bertahap dan merge ke situs yang ada didukung lewat opsi "Merge into existing" importer JSON.
  5. Migrasi data app (App Migration Platform) tidak dibangun.
  6. Laporan dan log migrasi memakai log job impor.
  7. Pusat perencanaan migrasi di konsol organisasi tidak dibangun.
  8. Penilaian app rinci (ketersediaan jalur migrasi per app dan keputusan admin per app) tidak dibangun.
  9. Consent admin untuk akses vendor ke data migrasi app tidak relevan karena tidak ada migrasi data app.
  10. API pemetaan ID entitas lama ke baru untuk vendor app tidak dibangun; importer menyimpan pemetaan externalId → ID internal untuk keperluan audit.
  11. ⚠ Pemetaan custom field type dan workflow rule milik app ke padanannya berada di luar cakupan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja. Prioritas importer "dari Jira" (REST API instans sumber) dapat dinaikkan karena sangat membantu adopsi; lihat Pertanyaan terbuka.

### INT-027 Project import from backup (Data Center)

- **Deskripsi:** Fitur edisi Data Center untuk merestore satu proyek dari backup XML tanpa menimpa seluruh instans. Tidak dibangun; padanannya restore granular INT-024/INT-025.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Pemilihan file backup XML dan direktori lampiran tidak dibangun.
  2. Pemilihan satu proyek dari backup tersedia lewat restore granular (INT-024 butir 6).
  3. Validasi kecocokan konfigurasi (custom field, workflow, scheme, versi) dilakukan oleh dry-run restore granular.
  4. Restore work item, comment, worklog, history, versions, dan components dicakup restore granular.
  5. Laporan galat dan pengguna yang hilang memakai log job.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; PRJ-067 ([04 Proyek](04-proyek.md)) dan ADM-040.

### INT-028 Data Pipeline (Data Center)

- **Deskripsi:** Fitur edisi Data Center untuk mengekspor data massal ke file CSV bagi platform BI tanpa membebani REST API. Tidak dibangun sebagai paritas; subset "bulk export job" dicatat sebagai opsi untuk instans self-hosted.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Admin Situs, analis data.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman admin Data pipeline tidak dibangun.
  2. ⚠ Full export entitas (work item, field, history, sprint, pengguna, proyek) hanya dicatat sebagai kebutuhan opsional untuk instans self-hosted.
  3. ⚠ Penjadwalan ekspor berkala dan retensi file tidak dibangun.
  4. Penulisan output ke direktori shared home tidak relevan; opsi self-hosted menulis ke object storage.
  5. ⚠ Pembatasan cakupan (updated-since, pengecualian custom field) menjadi bagian opsi ekspor inkremental self-hosted.
  6. Pembatalan ekspor dan riwayat status mengikuti job queue umum.
  7. REST API pemicu/status/batal tidak dibangun terpisah.
  8. Dampak performa diatasi dengan replika baca (lihat modul 27).
  9. Kompatibilitas template dashboard BI vendor berada di luar cakupan.
  10. Edisi Data Center berlisensi tidak ada; padanan Cloud adalah INT-029.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rujukan saja; lihat [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

### INT-029 Atlassian Analytics & Atlassian Data Lake

- **Deskripsi:** Padanan layanan analitik lintas produk dan data lake untuk membuat chart/dashboard berbasis SQL dari data produk ini (work item, Service Management, Assets) dan sumber eksternal. Untuk produk ini diwujudkan sebagai read replica/CDC ke warehouse dengan skema terdokumentasi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Org admin (koneksi dan izin), analis, manajer.
- **Aturan perilaku (acceptance criteria):**
  1. Admin membuat koneksi data lake dengan memilih produk/situs (work item, Service Management, Assets) dan cakupan proyek/service space yang disertakan.
  2. ⚠ Skema terdokumentasi memuat tabel work item, field, changelog, sprint, worklog, project, user, status, version, component dengan kolom ID stabil; dokumentasi schema beserta deskripsi kolom tersedia di dalam produk.
  3. Data source eksternal (Snowflake, Redshift, BigQuery, PostgreSQL, MySQL, SQL Server, Databricks) dapat dihubungkan lewat kredensial terenkripsi dan INT-001.
  4. Tersedia Visual SQL (no-code) dan mode SQL penuh; SQL hanya-baca dan dibatasi timeout.
  5. Tipe chart: bar, line, area, pie, table, pivot, KPI, scatter, heatmap.
  6. Dashboard mendukung controls/filter interaktif dan template bawaan (mis. siklus waktu, throughput).
  7. Dashboard dapat dibagikan dengan izin per dashboard, di-embed ke halaman dokumentasi, dan dikirim lewat email terjadwal.
  8. Opsi koneksi "Enforce permissions" membuat data yang dibaca pengguna mengikuti Browse Projects dan issue security; bila dimatikan, akses mengikuti izin koneksi yang diatur admin, koneksi hanya boleh diberikan ke group admin, UI memperingatkan bahwa issue security tidak diterapkan per baris, dan aksi tercatat di audit log.
  9. ⚠ Data share ke platform eksternal (Snowflake/Databricks atau ekspor terjadwal Parquet/CSV ke object storage) tersedia sebagai opsi lanjutan.
  10. ⚠ Data lake diperbarui berkala; latensi sinkronisasi (default ≤ 3 jam) dan retensi data (mengikuti retensi produk) ditampilkan di halaman koneksi, dan waktu "terakhir diperbarui" tampil di setiap dashboard.
  11. ⚠ Untuk Service Management, data lake menyediakan tabel requests, SLA, CSAT, Assets, alerts, dan percakapan virtual agent dengan skema terdokumentasi.
  12. Tersedia template dashboard Service Management (volume, SLA, CSAT, beban agent).
  13. Data yang dikecualikan admin dari koneksi (service space/proyek atau jenis data) dihapus dari data lake pada sinkronisasi berikutnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** INT-029 adalah spesifikasi pemilik (source of truth) Analytics & Data Lake; RPT-041 Atlassian Data Lake & data shares ([16 Laporan & Insight](16-laporan-insight.md)) dan JSM-AI-066 Atlassian Analytics & Data Lake for JSM ([35 JSM AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md)) menjadi rujukan singkat ke sini; RPT-038 Atlassian Analytics ([16 Laporan & Insight](16-laporan-insight.md)) adalah sisi analitik yang sama dan harus merujuk ke INT-029 sebagai pemilik spesifikasi (bila ada perbedaan butir, INT-029 yang berlaku). Tier tunggal: enterprise (P3 / Fase 4) untuk inti; butir 11-13 khusus Service Management baru aktif setelah modul Service Management tersedia (Fase 5-7) dan mengikuti entitlement paket Service Management. ⚠ Tier paket (enterprise vs premium untuk Service Management) ditetapkan di entitlement. Laporan bawaan di [16 Laporan & Insight](16-laporan-insight.md) dan dashboard di [17 Dashboard & Gadget](17-dashboard-gadget.md). Arsitektur read replica PostgreSQL dengan view tersanitasi diputuskan di [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md).

## Integrasi chat (Slack, Microsoft Teams, Google Chat)

### INT-030 Jira Cloud for Slack – installation & account connection

- **Deskripsi:** Konektor Slack resmi produk ini yang menghubungkan workspace Slack dengan satu atau beberapa situs serta akun Slack tiap pengguna dengan akun produk ini. Semua konektor chat memakai abstraksi `ChatProvider` yang sama.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs (memasang dan mengatur), admin Slack (persetujuan app), semua pengguna (menghubungkan akun pribadi).
- **Aturan perilaku (acceptance criteria):**
  1. Konektor dapat dipasang dari direktori app Slack (OAuth v2 install) atau dari Settings > Apps > Integrations > Slack di produk ini; keduanya menghasilkan `IntegrationConnection` yang sama.
  2. Satu workspace Slack dapat terhubung ke beberapa situs, dan setiap perintah/modal meminta pilihan situs bila lebih dari satu.
  3. "Connect personal account" menjalankan OAuth per pengguna sehingga setiap aksi dari Slack dijalankan dengan izin pengguna tersebut; aksi tanpa akun terhubung memunculkan ajakan connect.
  4. ⚠ Halaman konfigurasi menampilkan daftar workspace terhubung, tombol putuskan koneksi, dan pengaturan siapa yang boleh membuat notifikasi channel (semua pengguna, Admin Proyek, atau Admin Situs saja).
  5. Slack Enterprise Grid didukung (instalasi tingkat org dengan pilihan workspace) dan alur app approval admin Slack dihormati.
  6. App Home Slack menampilkan ringkasan work item yang di-assign ke pengguna, notifikasi terbaru, dan pengaturan.
  7. Pengguna dapat memutus akunnya (token dicabut dan dihapus); uninstall workspace menghapus semua token dan subscription channel terkait.
  8. ⚠ Saat subscription atau unfurl dibuat di channel Slack Connect yang memuat organisasi eksternal, pengguna diberi peringatan bahwa data terlihat oleh pihak luar dan harus mengonfirmasi.
  9. Konektor ini terpisah dari kanal chat Service Management (modul 30/35) dan tidak menangani percakapan customer.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Token bot dan user disimpan terenkripsi; tabel `slack_user_link`. Fase 4 mengikuti pemetaan fase (Slack/Teams); sebelum itu tim dapat memakai Automation "Send web request" ke incoming webhook Slack (AUTO-054).

### INT-031 Slack channel notifications & filters

- **Deskripsi:** Subscription channel Slack ke event proyek atau filter sehingga update work item otomatis diposting ke channel tim.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Anggota proyek dan Admin Proyek (sesuai pengaturan INT-030 butir 4), anggota channel Slack.
- **Aturan perilaku (acceptance criteria):**
  1. Subscription dibuat lewat `/jira connect` atau UI "Manage notifications", dengan memilih situs dan proyek; hasilnya `ChannelSubscription {provider, channelId, projectId|filterJql, eventTypes[], filters, createdBy}`.
  2. ⚠ Subscription dapat berbasis saved filter atau JQL bebas untuk cakupan lintas proyek; JQL divalidasi saat disimpan.
  3. ⚠ Jenis event yang dapat dipilih: work item created, transitioned, commented, assigned, dan update field tertentu (priority, summary, due date, field pilihan).
  4. ⚠ Filter tambahan: work type, status/status tujuan, priority, label, dan component.
  5. Satu channel dapat memiliki beberapa subscription; setiap subscription dapat diedit, di-pause, dan dihapus.
  6. Pesan memuat card work item (key, summary, status, assignee, priority, actor) dengan tombol aksi Transition, Assign, Comment, Watch yang dijalankan dengan izin pengguna yang menekan.
  7. ⚠ Konten pesan dibatasi izin pembuat subscription: work item atau field yang tak dapat dilihat pembuat (termasuk issue security level) tidak diposting, dan bila pembuat kehilangan akses subscription dinonaktifkan otomatis dengan pemberitahuan.
  8. Private channel mensyaratkan bot diundang terlebih dulu; bila belum, UI menampilkan instruksi `/invite`.
  9. Halaman konfigurasi menampilkan daftar semua subscription (channel, cakupan, pembuat, status) untuk audit admin.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Event berasal dari katalog event [28 Model Data Inti](28-model-data-inti.md); evaluasi JQL di [12 JQL](12-jql-bahasa-kueri.md).

### INT-032 Slack personal notifications

- **Deskripsi:** Notifikasi DM dari konektor Slack ke tiap pengguna tentang work item yang relevan (mention, assignment, comment, perubahan status).
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna dengan akun Slack terhubung.
- **Aturan perilaku (acceptance criteria):**
  1. Jenis notifikasi: di-mention, di-assign, comment pada work item yang di-watch/di-assign/di-report, perubahan status, dan perubahan assignee.
  2. Setiap jenis dapat di-toggle lewat `/jira notify` atau App Home > Settings, dan pengaturan disimpan per pengguna.
  3. Dari DM pengguna dapat membalas comment, melakukan transition, dan assign tanpa membuka web.
  4. Pengguna tidak menerima notifikasi atas aksinya sendiri.
  5. Pengiriman menghormati status Do Not Disturb Slack (Slack menahan notifikasi; konektor tidak mengirim ulang).
  6. Opsi "Matikan semua" menonaktifkan seluruh notifikasi personal sekaligus.
  7. Notifikasi email berjalan terpisah dan tidak terpengaruh pengaturan Slack.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kanal "chat DM" di mesin notifikasi; preferensi di NTF-039 ([19 Notifikasi & Email](19-notifikasi-email.md)).

### INT-033 Slack create issue & comment from message

- **Deskripsi:** Membuat work item atau menambahkan comment ke work item dari pesan Slack lewat message shortcut atau slash command.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Pengguna dengan akun terhubung; butuh Create issues/Add comments di proyek tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Message shortcut "Create work item" membuka modal dengan summary/description terisi dari pesan, serta pilihan situs, proyek, work type, assignee, dan priority.
  2. `/jira create` membuka modal create yang sama dari channel atau DM mana pun.
  3. ⚠ Field wajib bertipe sederhana (teks, select, user, tanggal) dirender di modal; bila ada field wajib kompleks, modal menampilkan tautan "Lanjutkan di web" dengan isian terbawa.
  4. Setelah dibuat, bot membalas di thread dengan card work item.
  5. Message shortcut "Add comment to work item" memilih work item lewat pencarian atau key lalu menambahkan isi pesan sebagai comment.
  6. Permalink pesan Slack disertakan sebagai web link/referensi di work item atau comment.
  7. ⚠ File lampiran pesan ikut disalin bila pengguna mencentang "Sertakan lampiran" dan ukurannya di bawah batas lampiran; bila gagal, permalink file disertakan.
  8. ⚠ Langkah Slack Workflow Builder "Create work item" disediakan sebagai custom step.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Metadata create memakai createmeta [23 API Publik & Webhook](23-api-publik-webhook.md).

### INT-034 Slack link unfurl & interactive actions

- **Deskripsi:** Saat URL work item diposting di Slack, konektor menampilkan preview kaya beserta tombol aksi.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / free-standard
- **Aktor & izin:** Pengguna yang memposting (harus terhubung dan berhak melihat), pengguna yang menekan tombol (izin masing-masing).
- **Aturan perilaku (acceptance criteria):**
  1. Unfurl URL work item menampilkan key, summary, status, assignee, priority, work type, dan potongan description.
  2. Tombol aksi: Transition, Assign (ke saya atau pengguna lain), Comment, Watch, Open in browser.
  3. ⚠ URL filter, board, dan proyek di-unfurl menjadi card ringkas (nama, jumlah hasil/pemilik).
  4. Preview hanya tampil bila pemosting terhubung dan punya akses; bila tidak, pemosting menerima pesan ephemeral ajakan connect dan tidak ada data yang diposting.
  5. ⚠ Admin workspace atau pengguna dapat menonaktifkan unfurl per channel.
  6. ⚠ Key work item mentah (mis. `PROJ-123`) dalam pesan dideteksi dan ditawarkan unfurl secara ephemeral kepada pemosting.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mengikuti batas data INT-030 butir 8.

### INT-035 Slack slash commands & share to Slack

- **Deskripsi:** Set perintah `/jira` untuk mengelola integrasi dari Slack dan berbagi work item dari produk ini ke Slack. Nama perintah final mengikuti nama produk.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna terhubung.
- **Aturan perilaku (acceptance criteria):**
  1. `/jira connect` menghubungkan channel ke proyek/filter (INT-031).
  2. `/jira manage` menampilkan dan mengelola subscription channel.
  3. `/jira notify` membuka pengaturan notifikasi personal (INT-032).
  4. `/jira create` membuat work item (INT-033).
  5. `/jira help` menampilkan daftar perintah beserta contoh.
  6. ⚠ `/jira logout` memutus akun pribadi pengguna.
  7. ⚠ Dialog Share di work item menyediakan tujuan channel/pengguna Slack dan memposting card dengan pesan opsional.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dialog Share di NAV-016 ([21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md)).

### INT-036 Jira Cloud for Microsoft Teams

- **Deskripsi:** Konektor Microsoft Teams dengan bot, notifikasi channel/personal, tab, message extension, dan link unfurl, memakai abstraksi `ChatProvider` yang sama dengan Slack.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna, admin tenant Microsoft 365, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Konektor dipasang dari Teams app store; pengguna menghubungkan situs dan akun produk ini lewat OAuth per pengguna.
  2. Personal app menampilkan work item saya, notifikasi, dan pencarian.
  3. ⚠ Bot command yang didukung: `create`, `find`, `connect`, `help`, dan `logout`.
  4. Notifikasi channel per proyek/filter dengan pemilihan jenis event (mengikuti INT-031).
  5. Notifikasi personal untuk mention, assignment, dan comment pada work item yang di-watch.
  6. Channel tab dapat menyematkan board, filter, atau daftar work item.
  7. Message extension memungkinkan mencari dan membagikan work item sebagai card.
  8. Message action "Create work item" dari pesan Teams dengan prefill summary/description.
  9. ⚠ URL work item di-unfurl menjadi Adaptive Card dengan aksi Transition, Assign, Comment.
  10. Instalasi dapat mensyaratkan persetujuan admin tenant Microsoft 365; status menunggu persetujuan ditampilkan.
  11. Berjalan di desktop, web, dan mobile; batasan Adaptive Card (mis. jumlah aksi) ditangani dengan tautan "Buka di browser".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Preferensi personal di NTF-040 ([19 Notifikasi & Email](19-notifikasi-email.md)).

### INT-037 Jira Cloud for Google Chat

- **Deskripsi:** Konektor Google Chat untuk menerima notifikasi dan berinteraksi dengan work item dari space Google Chat.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna, admin Google Workspace.
- **Aturan perilaku (acceptance criteria):**
  1. Konektor dipasang dari Google Workspace Marketplace dan pengguna menghubungkan akun produk ini.
  2. Space dapat berlangganan notifikasi per proyek untuk event created, transitioned, dan commented.
  3. ⚠ Notifikasi personal dan aksi dasar (assign, comment, transition) tersedia di DM bot.
  4. ⚠ URL work item di-unfurl menjadi card.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Status app referensi tidak aktif dikembangkan; konektor ini prioritas terendah di antara chat provider.

## Integrasi alat produktivitas (spreadsheet & email)

### INT-038 Jira Cloud for Google Sheets

- **Deskripsi:** Add-on Google Sheets resmi produk ini untuk menarik data work item lewat JQL/filter ke spreadsheet, dengan refresh terjadwal.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna, analis; data mengikuti izin pengguna yang mengotorisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Add-on dipasang dari Google Workspace Marketplace; pengguna login OAuth dan memilih situs (lebih dari satu situs didukung).
  2. Sidebar "Get work items" menerima basic search, JQL, atau saved filter.
  3. Pengguna memilih field/kolom termasuk custom field dan mengatur urutannya.
  4. Hasil ditulis ke sheet/range tertentu; satu spreadsheet dapat memuat beberapa query dengan range berbeda.
  5. ⚠ Refresh manual dan scheduled refresh (pilihan per jam, harian, mingguan) tersedia per query.
  6. ⚠ Custom function di sel, mis. `=WORKITEMS("jql", "kolom")`, mengembalikan hasil sebagai array.
  7. Opsi "Export > Google Sheets" dari navigator membuka Sheets dengan query terkait (INT-021 butir 11).
  8. Data mengikuti izin pengguna yang mengotorisasi; sidebar menampilkan peringatan bahwa berbagi sheet dapat membocorkan data ke orang yang tidak berhak.
  9. ⚠ Batas default 10.000 baris per tarikan; hasil terpotong diberi catatan di sel status.
  10. ⚠ Pembaruan work item dari Sheets tidak didukung pada rilis pertama (read-only).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai REST search [23 API Publik & Webhook](23-api-publik-webhook.md); UI pintasan di SRC-055.

### INT-039 Jira Cloud for Microsoft Excel

- **Deskripsi:** Add-in Excel untuk mengimpor work item ke workbook lalu menyegarkannya.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna, analis; mengikuti izin pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Add-in dipasang dari Microsoft AppSource; pengguna login OAuth.
  2. Work item diambil lewat JQL, filter, atau basic search.
  3. Pengguna memilih kolom/field termasuk custom field.
  4. ⚠ Refresh on-demand tersedia; refresh terjadwal tidak didukung pada rilis pertama.
  5. Mendukung Excel desktop dan web.
  6. Data mengikuti izin pengguna yang login.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Alternatif minimum: ekspor CSV/XLSX (INT-021, SRC-054) dan feed JSON dengan API token; add-in dibangun bila ada permintaan.

### INT-040 Jira Cloud for Outlook

- **Deskripsi:** Add-in Outlook untuk membuat work item dari email, mendeteksi key work item di email, dan menindaklanjuti tanpa meninggalkan inbox.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna; mengikuti izin pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Add-in dipasang dari AppSource atau di-deploy admin Microsoft 365 untuk seluruh organisasi.
  2. Create work item dari email: subject menjadi summary, body menjadi description, dan lampiran email dapat dipilih untuk ikut.
  3. Key/URL work item di email dideteksi dan detailnya ditampilkan di panel.
  4. Panel menyediakan aksi comment, assign, transition, dan watch.
  5. Isi email dapat ditambahkan sebagai comment pada work item yang ada.
  6. Pengguna memilih situs, proyek, dan work type saat membuat.
  7. Semua aksi mengikuti izin pengguna.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari mail handler email-to-work-item ([19 Notifikasi & Email](19-notifikasi-email.md), ADM-013).

### INT-041 Jira Cloud for Gmail

- **Deskripsi:** Add-on Gmail untuk melihat dan menindaklanjuti work item yang disebut dalam email serta membuat work item dari email.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Semua pengguna; mengikuti izin pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Add-on dipasang dari Google Workspace Marketplace per pengguna atau oleh admin untuk seluruh domain.
  2. Key/URL work item di email dideteksi dan ditampilkan sebagai card.
  3. ⚠ Card menyediakan aksi cepat comment, assign, dan transition.
  4. Create work item dengan prefill summary dari subject dan description dari body.
  5. Autentikasi memakai OAuth per pengguna.
  6. Mendukung Gmail web dan mobile.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbagi komponen deteksi key dengan INT-040.

## Integrasi dokumentasi, desain & media

### INT-042 Confluence site linking & cross-product setup

- **Deskripsi:** Menghubungkan produk ini dengan wiki/dokumentasi (Confluence asli lewat REST API atau modul Pages internal bila kelak ada) agar link halaman, makro, mention, dan Project pages dapat dipakai, memakai konsep "linked documents" generik.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs, Org admin, semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Bila penyedia dokumentasi berada di organisasi/situs yang sama (identitas bersama), koneksi aktif otomatis tanpa konfigurasi.
  2. Penyedia di situs lain atau self-hosted dihubungkan lewat Application links (INT-005).
  3. App switcher menampilkan penyedia dokumentasi terhubung; ajakan trial lintas produk tidak dibangun.
  4. Izin dokumentasi tetap terpisah: pengguna butuh akses ke penyedia untuk melihat/membuat halaman, dan bila tidak berhak tampil pesan "Anda tidak memiliki akses ke halaman ini".
  5. Pengguna dapat mengirim permintaan akses produk dokumentasi yang diteruskan ke admin untuk disetujui (USR-018).
  6. Halaman yang menyebut work item otomatis membuat remote link berelasi "mentioned in" di work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk ini tidak memiliki wiki; keputusan modul Pages internal vs konektor eksternal ada di Pertanyaan terbuka. Product access di [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md).

### INT-043 Confluence pages in issue view (link, create, mentioned in)

- **Deskripsi:** Dari work item, pengguna dapat mengaitkan halaman dokumentasi yang ada, membuat halaman baru dari template, dan melihat halaman yang menyebut work item tersebut.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Anggota proyek dengan izin Link issues/Edit issues; akses halaman mengikuti penyedia dokumentasi.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Add > Page" (halaman di penyedia dokumentasi terhubung, mis. Confluence) membuka pencarian halaman berdasarkan judul atau menerima tempelan URL.
  2. "Create page" dari work item meminta space, parent page, dan template (Product requirements, Decision, Retrospective, Meeting notes, How-to); halaman yang dibuat otomatis ter-link balik ke work item.
  3. Bagian "Pages" memisahkan "Linked pages" dan "Mentioned in".
  4. Jumlah halaman terkait tampil di area linked work items/glance.
  5. Unlink menghapus remote link tanpa menghapus halaman di penyedia.
  6. Halaman draft atau restricted tidak ditampilkan kepada pengguna yang tidak berhak (jumlah pun tidak menghitungnya).
  7. Setiap halaman dirender sebagai Smart Link card (judul, space, terakhir diperbarui).
  8. ⚠ Whiteboard yang memuat kartu work item juga tercatat sebagai "mentioned in".
  9. ⚠ Saran halaman terkait berbasis AI tidak termasuk rilis pertama dan mengikuti modul AI bila diaktifkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model `RemoteLink {issueId, url, title, relationship, application{type,name}, globalId, icon, status}`; tampilan di CMT-030/CMT-031 ([08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md)) dan ISS-030.

### INT-044 Confluence Jira macros (Jira Issues, Jira Chart, Jira Timeline) & Team Calendars

- **Deskripsi:** Makro di penyedia dokumentasi untuk menampilkan data live produk ini (work item tunggal, tabel JQL/filter, jumlah, chart, timeline) serta feed kalender. Makro berjalan di sisi penyedia; produk ini menyediakan API dan feed yang dibutuhkan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Penulis dokumentasi, pembaca; data dirender dengan izin pembaca.
- **Aturan perilaku (acceptance criteria):**
  1. Makro mode single menampilkan lozenge key + summary + status work item.
  2. Mode tabel menerima JQL, saved filter, atau URL pencarian; kolom dapat dipilih dan diurutkan, ada batas maksimum work item, dan tombol refresh.
  3. Mode count menampilkan jumlah hasil query sebagai link ke navigator.
  4. Makro chart: Pie (statistic type per field), Created vs Resolved (periode, kumulatif, tampilkan versi), dan Two-dimensional (sumbu X/Y, jumlah baris), dengan data dari endpoint chart.
  5. Opsi chart: border, show info, dan lebar.
  6. ⚠ Makro timeline menyematkan timeline proyek atau plan dalam mode baca.
  7. URL filter/board/work item yang ditempel dapat di-embed lewat Smart Link (inline/card/embed).
  8. Data dirender dengan izin pembaca; pembaca tanpa akses melihat pesan "Work item tidak dapat ditampilkan".
  9. Makro dapat disisipkan lewat slash command di editor penyedia.
  10. Feed kalender (iCal) dari JQL/filter berdasarkan field tanggal, event sprint, dan event versi/rilis tersedia untuk kalender tim.
  11. ⚠ Blueprint laporan legacy (Change log/Status report) tidak dibangun; cukup makro tabel/chart.
  12. ⚠ Respons makro di-cache singkat (usulan 5 menit) dan hasil query dipotong sesuai batas dengan catatan "menampilkan N dari M".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Membutuhkan JQL search, field metadata, dan chart data API ([12 JQL](12-jql-bahasa-kueri.md), [16 Laporan & Insight](16-laporan-insight.md)); kalender proyek di PLAN-022.

### INT-045 Create Jira issues from Confluence (inline, table bulk create, whiteboards)

- **Deskripsi:** Membuat work item langsung dari teks yang di-highlight, baris tabel, atau sticky note whiteboard di penyedia dokumentasi; work item yang terbentuk otomatis ter-link ke halaman.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Penulis dokumentasi dengan Create issues di proyek tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Highlight teks lalu "Create work item": summary terisi dari teks, lalu pengguna memilih situs, proyek, work type, dan description.
  2. Bulk create dari tabel: header kolom dipetakan ke field (Summary, Description, dan lainnya), lalu banyak work item dibuat sekaligus dengan ringkasan hasil per baris.
  3. Teks yang di-highlight diganti makro single work item.
  4. ⚠ Field wajib divalidasi; bila ada field wajib yang tidak didukung dialog, pengguna diarahkan ke form create penuh dengan isian terbawa.
  5. Sticky note whiteboard dapat dikonversi menjadi work item, dan kartu tersinkron dua arah untuk status dan assignee.
  6. Work item yang dibuat mendapat remote link ke halaman asal.
  7. ⚠ Pengguna dapat memilih Epic/parent saat membuat dari halaman requirements.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Membutuhkan endpoint create, createmeta, dan remote link ([23 API Publik & Webhook](23-api-publik-webhook.md)).

### INT-046 Project pages, release Related work & release notes

- **Deskripsi:** Menghubungkan proyek dan rilis ke dokumentasi atau link lain agar dokumen tim (requirements, retrospective, release notes) dapat diakses dari produk ini. Minimum yang dibangun lebih awal: Related work generik pada versi dan Project shortcuts.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Admin Proyek (mengatur koneksi), anggota proyek (membaca/menambah Related work dengan Manage versions/Edit).
- **Aturan perilaku (acceptance criteria):**
  1. Sidebar proyek memiliki item "Pages" untuk menghubungkan space atau page tree dokumentasi yang ada, atau membuat space baru bila penyedia mendukung.
  2. Page tree tertanam ditampilkan dengan pencarian dan tombol "Create page" dari template.
  3. Hanya Admin Proyek yang dapat mengubah atau memutus koneksi space; pengguna lain melihat tombol dinonaktifkan.
  4. Saran template mengikuti tipe proyek (scrum: sprint planning dan retrospective; kanban: dokumen proses).
  5. ⚠ Proyek company-managed dan team-managed sama-sama mendukung Pages; lokasi item sidebar sama di keduanya.
  6. Halaman versi/rilis memiliki "Related work" dengan kategori Communication, Design, Documentation, Software, Other, berisi URL apa pun atau halaman dokumentasi beserta judul.
  7. ⚠ Halaman rilis menyediakan aksi "Create release notes" yang membuat halaman dokumentasi berisi daftar work item rilis, atau mengunduh release notes Markdown bila tidak ada penyedia.
  8. Project shortcuts (link ke URL eksternal di sidebar) tersedia sebagai alternatif ringan tanpa penyedia dokumentasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan (butir 5).
- **Catatan & ketergantungan:** Versi di [15 Rilis, Versi & Komponen](15-rilis-versi-komponen.md); shortcuts di [04 Proyek](04-proyek.md); sprint pages di SPR-026.

### INT-047 Loom integration

- **Deskripsi:** Merekam dan menyematkan video penjelasan (Loom atau padanannya) dari work item untuk bug, demo, atau update async; produk ini menyediakan perekam layar berbasis MediaRecorder sebagai lampiran dan embed penyedia video via oEmbed.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna dengan Create attachments/Add comments; Org admin untuk akses penyedia video.
- **Aturan perilaku (acceptance criteria):**
  1. Editor description dan comment memiliki tombol rekam video; hasil rekaman otomatis disisipkan sebagai embed (lampiran atau video penyedia).
  2. Smart Link embed video menampilkan player inline, judul, durasi, dan thumbnail.
  3. Untuk Loom, pengguna harus terhubung ke workspace Loom; tanpa koneksi tombol menampilkan ajakan connect.
  4. Izin menonton mengikuti pengaturan sharing penyedia video (workspace, link, privat); video internal mengikuti izin lampiran work item.
  5. ⚠ Pembuatan work item/bug report dari video dengan bantuan AI tidak termasuk rilis pertama.
  6. Admin dapat mengaktifkan/menonaktifkan perekaman dan penyedia video di pengaturan organisasi.
  7. ⚠ Ketersediaan perekam bawaan dapat dibatasi per tier lewat entitlement.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tampilan di CMT-032 ([08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md)); penyimpanan video memakai storage lampiran.

### INT-048 Figma for Jira & design integrations

- **Deskripsi:** Integrasi desain yang menautkan file atau frame Figma ke work item, menampilkan preview live, dan status desain.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Desainer, Developer/Anggota Tim; desain privat hanya terlihat oleh pengguna yang berhak di Figma.
- **Aturan perilaku (acceptance criteria):**
  1. Desain ditautkan lewat tempel URL (Smart Link embed) atau aksi "Add design".
  2. ⚠ Bagian "Designs" di work item menampilkan daftar desain, thumbnail, dan status (mis. "Ready for development").
  3. Status desain disinkronkan dari Figma Dev Mode, dan perubahan desain sejak ditautkan ditandai "Updated".
  4. Plugin di Figma menampilkan work item di canvas dengan status live, serta membuat atau menautkan work item dari Figma.
  5. Preview interaktif (pan/zoom) tersedia lewat embed.
  6. Setiap pengguna menghubungkan akun Figma lewat OAuth; desain privat tampil "Restricted" bagi yang tidak berhak.
  7. ⚠ Data desain diterima lewat modul design provider di API DevOps (`designInfoProvider`) agar penyedia lain dapat memakai kontrak yang sama.
  8. Zeplin, Adobe XD, Miro, Lucid, dan Sketch dihubungkan lewat Smart Links atau app vendor masing-masing (bukan konektor bawaan).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Entitas `Design {issueId, url, provider, title, status, thumbnail, lastUpdated}`; data provider di INT-059 dan [24 Integrasi DevOps](24-integrasi-devops.md).

### INT-049 Trello integration (Jira Power-Up)

- **Deskripsi:** Power-Up di Trello untuk menautkan atau membuat work item dari kartu Trello dan menampilkan status work item di kartu.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna Trello, semua pengguna; izin mengikuti akun produk ini.
- **Aturan perilaku (acceptance criteria):**
  1. Power-Up diaktifkan di board Trello lalu dihubungkan ke situs produk ini.
  2. Work item yang ada dapat dilampirkan ke kartu lewat pencarian key atau summary.
  3. Work item dapat dibuat dari kartu (judul menjadi summary, isi menjadi description).
  4. Badge status tampil di depan kartu dan detail work item di card back.
  5. Izin mengikuti akun produk ini milik pengguna Trello.
  6. Kartu/board Trello yang ditempel di produk ini menjadi Smart Link.
  7. Integrasi ini berbeda dari Trello importer (INT-017).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Status Power-Up referensi perlu ditinjau menjelang implementasi; Power-Up dibangun hanya bila ada permintaan.

### INT-050 Atlassian for VS Code (IDE integration)

- **Deskripsi:** Extension IDE (VS Code) untuk menampilkan dan mengelola work item dari editor kode, cukup dibangun di atas REST API publik.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer/Anggota Tim; mengikuti izin pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Extension login ke instans produk ini (OAuth atau API token) dan pengguna memilih situs.
  2. Daftar work item berdasarkan JQL, lihat detail, create, transition, dan comment tersedia di panel IDE.
  3. "Start work" dari work item membuat branch Git yang memuat key work item dan opsional melakukan transition ke In Progress.
  4. Key work item di kode/commit dideteksi dan menjadi tautan ke work item.
  5. Integrasi PR dan pipeline ditampilkan untuk penyedia yang terhubung (GitHub, GitLab, Bitbucket) melalui data development DEV-001 ([24 Integrasi DevOps](24-integrasi-devops.md)).
  6. Login memakai akun produk ini melalui OAuth 2.0 dengan PKCE (bukan akun Atlassian); API token hanya sebagai alternatif, dan token dapat dicabut dari Connected apps (AUTH-013, [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md)).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** INT-050 adalah spesifikasi pemilik (source of truth) extension IDE; DEV-031 Atlassian for VS Code (IDE integration) di [24 Integrasi DevOps](24-integrasi-devops.md) menjadi rujukan singkat ke sini. Ekstensi resmi Atlassian hanya bekerja dengan produk referensi; membangun ekstensi sendiri adalah keputusan terbuka. Memakai REST API publik [23 API Publik & Webhook](23-api-publik-webhook.md). Extension JetBrains adalah ekosistem. Integrasi branch/commit di [24 Integrasi DevOps](24-integrasi-devops.md).

## Integrasi lintas produk (rujukan silang)

### INT-051 Opsgenie integration

- **Deskripsi:** Integrasi alert dengan work item untuk membuat work item dari alert dan menyinkronkan status dua arah. Untuk produk ini kemampuan tersebut menjadi bagian Service Management Operations; fitur ini hanya kontrak konektor ke alat alerting eksternal.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** On-call engineer, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Konektor alerting dapat membuat work item otomatis dari alert dengan filter priority, tag, dan source.
  2. Field alert dipetakan ke field work item oleh admin konektor.
  3. Sinkronisasi dua arah lewat webhook: transition work item menutup atau acknowledge alert, dan sebaliknya.
  4. Comment disinkronkan dua arah sebagai catatan alert.
  5. Work item menampilkan remote link ke alert asal.
  6. ⚠ Opsgenie sudah end-of-sale (Juni 2025); produk ini tidak membangun konektor khusus Opsgenie dan mengarahkan ke modul Operations bawaan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rincian alert/on-call dan sinkronisasi work item di [34 JSM Operations](34-jsm-operations-alerting-on-call.md) (JSM-OPS-031, JSM-OPS-038); migrasi dari Opsgenie di JSM-OPS-047.

### INT-052 Statuspage integration

- **Deskripsi:** Integrasi halaman status publik dengan incident Service Management untuk mengaitkan incident dengan status komponen yang dilihat publik.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (incident manager), Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Admin menghubungkan halaman status eksternal (API key) ke situs/proyek layanan.
  2. Dari work item incident, pengguna dapat membuat atau memperbarui incident di halaman status.
  3. Pengguna memilih komponen terdampak dan status komponen (operational, degraded, partial outage, major outage).
  4. Template pesan incident dapat dipilih dan diedit sebelum dipublikasikan.
  5. ⚠ Otomasi menyediakan aksi "Update status page" untuk publikasi otomatis.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tanpa modul layanan, integrasi hanya lewat Smart Link/webhook. Incident dan status page di [32 JSM ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) (JSM-ITSM-026 status page bawaan, JSM-ITSM-027 integrasi Statuspage).

### INT-053 Compass integration (Compass components in Jira)

- **Deskripsi:** Menghubungkan katalog software component eksternal ke proyek sehingga field Components dapat memakai komponen katalog beserta metadatanya.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Proyek, Developer/Anggota Tim.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Project settings > Components menyediakan pilihan sumber "Komponen proyek" atau "Komponen katalog"; pilihan katalog hanya untuk company-managed pada rilis pertama.
  2. Saat memilih komponen katalog di work item, owner team, scorecard, dan health ditampilkan.
  3. Work item menautkan ke halaman komponen di katalog.
  4. ⚠ Beralih ke komponen katalog tidak menghapus komponen proyek lama; nilai lama tetap tersimpan read-only dan dapat dipetakan.
  5. Pilihan katalog hanya tersedia bila katalog komponen terhubung di situs yang sama.
- **Team-managed vs company-managed:** Rilis pertama hanya company-managed (butir 1).
- **Catatan & ketergantungan:** Komponen di [15 Rilis, Versi & Komponen](15-rilis-versi-komponen.md); service catalog di modul 32.

### INT-054 Jira Align connector

- **Deskripsi:** Konektor enterprise untuk menyinkronkan data (epic, story, sprint, release, team) dengan alat perencanaan portofolio eksternal. Produk ini menyediakan API sinkronisasi dua arah generik (webhook, query updated-since, external id).
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Portfolio admin, Admin Situs; akun integrasi khusus.
- **Aturan perilaku (acceptance criteria):**
  1. Koneksi dikonfigurasi di alat portofolio dengan URL situs, akun integrasi, dan autentikasi API (OAuth/API token).
  2. Proyek/board dipetakan ke program/team di alat portofolio.
  3. Sinkronisasi dua arah epic/feature/story, sprint/iteration, release/fix version, dan dependency berjalan lewat webhook + polling updated-since.
  4. Custom field dan status dipetakan per tipe item.
  5. Jadwal sinkronisasi dan log galat tersedia di konektor; API mengembalikan galat terstruktur per item.
  6. Aturan konflik/source of truth ditetapkan per field dan dihormati lewat versi optimistik (ETag/updated) di API.
  7. Banyak instans produk ini dapat terhubung ke satu alat portofolio, dibedakan lewat external id per instans.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk portofolio eksternal di luar cakupan; Plans internal di [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md).

### INT-055 Rovo connectors & third-party knowledge in Jira

- **Deskripsi:** Konektor pengetahuan yang mengindeks konten pihak ketiga (Google Drive, SharePoint, Slack, GitHub) agar dapat dicari dan dipakai agent AI di produk ini. Seluruhnya opsional dan bergantung pada fitur AI.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Org admin (mengaktifkan connector), semua pengguna, Integrator (agent/action dari app).
- **Aturan perilaku (acceptance criteria):**
  1. Admin mengaktifkan connector per sumber (Google Drive, SharePoint, Slack, Teams, GitHub, Figma) dan memilih cakupan yang diindeks.
  2. Pencarian lintas sumber hanya mengembalikan item yang boleh dilihat pengguna menurut izin sistem asal.
  3. App dapat mendaftarkan agent dan action AI lewat manifest.
  4. Chat AI di work item memakai konteks work item (summary, description, comment) dan sumber terindeks.
  5. Admin dapat mengaktifkan/menonaktifkan fitur AI dan setiap connector; menonaktifkan connector menghapus indeksnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Tier/paket AI ditetapkan di entitlement. Aspek AI lain di modul 35 dan [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md).

### INT-056 Atlassian Goals & Projects (Atlassian Home) integration

- **Deskripsi:** Rujukan singkat. Padanan native fitur ini adalah entitas Goal milik produk ini di PLAN-061 Atlassian Goals beserta penautan work item di PLAN-062 Link work items to Goals ([18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md)); penautan idea Product Discovery ke goal/project ada di JPD-INS-036 Atlassian Goals & Projects Integration (Atlas) ([37 JPD Insights & Delivery](37-jpd-insights-delivery-peran-integrasi.md)). Integrasi ke layanan Atlassian Home/Goals & Projects milik Atlassian berada di luar cakupan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Mengikuti PLAN-061/PLAN-062 (pembuat goal, goal owner, Edit Issues untuk menautkan) dan JPD-INS-036 (role Administrator dan Creator Product Discovery).
- **Aturan perilaku (acceptance criteria):**
  1. Field/panel Goals di work item dan progres linked work pada goal mengikuti PLAN-062 butir 1-3; tidak ada spesifikasi terpisah di modul ini.
  2. Project lintas tim yang menautkan epic sebagai sumber progres dipenuhi oleh goal native (PLAN-061) yang ditautkan ke epic lewat PLAN-062; entitas "Project" ala Atlassian Home tidak dibangun terpisah.
  3. Status update berkala dengan progres work item terkait (selesai/total) mengikuti PLAN-061 butir 3 dan PLAN-062 butir 3.
  4. Visibilitas goal dan work item terkait mengikuti permission di PLAN-062 butir 2 (hanya yang boleh diakses pengguna).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini dipertahankan sebagai heading agar inventaris tercakup; spesifikasi pemilik adalah PLAN-061/PLAN-062 (P2 / Fase 4) untuk goal native dan JPD-INS-036 (P2 / Fase 6) untuk idea Product Discovery. Prioritas P3 di sini hanya berlaku untuk konektor opsional ke layanan Goals & Projects eksternal, yang tidak dibangun (lihat "Di luar cakupan & ekosistem"); perbedaan prioritas dengan JPD-INS-036 disengaja karena kebutuhan fungsionalnya sudah dipenuhi padanan native. Administrasi Goals di ADM-029 Goals administration ([22 Administrasi Sistem](22-administrasi-sistem.md)).

### INT-057 Jira Product Discovery ↔ Jira delivery integration (cross-reference)

- **Deskripsi:** Menghubungkan idea Product Discovery dengan epic/work item delivery dan menampilkan progres delivery di idea. Rincian di modul 37.
- **Prioritas / Fase / Tier:** P3 / Fase 6 / jpd
- **Aktor & izin:** Product Manager Discovery, Product Owner/PM.
- **Aturan perilaku (acceptance criteria):**
  1. Dari idea, pengguna dapat membuat delivery ticket baru atau menautkan work item yang ada.
  2. Delivery progress dihitung dari kategori status work item terkait beserta child-nya.
  3. ⚠ Work item delivery menampilkan idea terkait di panel tersendiri.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rincian di [37 JPD Insights & Delivery](37-jpd-insights-delivery-peran-integrasi.md): JPD-INS-017 (Delivery tab), JPD-INS-018/JPD-INS-019 (delivery progress/status), JPD-INS-022 (idea di work item).

### INT-058 Assets integration (cross-reference)

- **Deskripsi:** Objek Assets (CMDB) dipakai sebagai custom field "Assets object" di proyek untuk menautkan aset/CI ke work item. Rincian di modul 33.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (konfigurasi field), semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Custom field Assets object mendukung filter AQL untuk membatasi objek yang dapat dipilih.
  2. Atribut aset terpilih ditampilkan di work item.
  3. Aset dapat diimpor dari CSV/JSON/database/cloud (dirinci modul 33).
  4. ⚠ Field Assets di proyek non-layanan hanya tersedia bila entitlement Assets aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rincian di [33 JSM Assets](33-jsm-assets-cmdb.md): JSM-AST-026 (Assets objects custom field), JSM-AST-031 s.d. JSM-AST-036 (impor aset).

### INT-059 Extension point: DevOps data providers (cross-reference)

- **Deskripsi:** Modul untuk mengirim data development (branch, commit, PR, build, deployment, feature flag, remote link, security, design, operations) ke work item. Rincian di modul 24.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (app dengan scope data provider).
- **Aturan perilaku (acceptance criteria):**
  1. Tipe provider: development info, builds, deployments, feature flags, remote links, security info, designs, operations.
  2. Data tampil di panel Development, Releases, Deployments, Security, dan Designs.
  3. Asosiasi ke work item lewat key di commit/branch/PR atau ID eksplisit.
  4. Pengiriman data membutuhkan scope `manage:data-provider` atau modul devops di manifest app.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat [24 Integrasi DevOps](24-integrasi-devops.md).

## Model app: strategi, katalog & manajemen

### INT-060 Extensibility & Plugin Strategy

- **Deskripsi:** Strategi ekstensi produk ini tanpa meniru marketplace pihak ketiga: memulai dari primitif API, lalu model app (OAuth client + scope + webhook), dan opsional UI extension point ber-sandbox. Menjadi acuan semua fitur INT-061 s.d. INT-080.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Arsitek, Integrator/Developer API, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Tahap awal ekstensi hanya melalui REST API, webhook, entity properties (INT-077), dan Automation "Send web request"; tidak ada kode pihak ketiga yang berjalan di produk.
  2. Model app terdiri dari OAuth client + daftar scope + registrasi webhook + akun app user, menggantikan konsep Connect/Forge.
  3. UI extension point (issue panel, glance, project page) bersifat opsional dan hanya diaktifkan setelah desain keamanan iframe sandbox + postMessage API disetujui.
  4. Custom field type dari app tidak diprioritaskan; INT-070 dijadwalkan paling akhir di antara extension point.
  5. Server tidak pernah mengeksekusi kode pihak ketiga di proses yang sama (tidak ada plugin Go `.so` atau eval skrip); kode app berjalan di browser sandbox atau di server milik vendor.
  6. Setiap extension point memiliki versi kontrak (mis. `issuePanel@1`) dan setiap app mendeklarasikan versi yang dipakai; perubahan breaking menaikkan versi mayor dengan masa dukung versi lama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** App ekosistem pihak ketiga (Tempo, Xray, ScriptRunner, BigPicture) hanya referensi, bukan fitur. NFR di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

### INT-061 Apps menu & in-product Marketplace discovery

- **Deskripsi:** Menu "Apps" di navigasi dan halaman "Explore apps" untuk mencari dan memasang app dari registry app milik produk ini (bukan Atlassian Marketplace).
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna (melihat menu dan meminta app), Admin Situs (memasang).
- **Aturan perilaku (acceptance criteria):**
  1. Menu "Apps" menampilkan global page dari app terpasang yang boleh dilihat pengguna, tautan "Explore more apps", dan "Manage your apps" (khusus admin).
  2. ⚠ Explore apps menyediakan pencarian, kategori, filter (gratis, berjalan di platform produk, dibuat tim inti, pilihan kurator, terverifikasi keamanan), dan sortir (popularitas, terbaru, rating).
  3. Listing app memuat deskripsi, screenshot, rating/review, jumlah instalasi, harga, dokumentasi, bagian Privacy & Security, data residency, dan info vendor.
  4. Tombol aksi menyesuaikan izin: admin melihat "Try it free"/"Install"; pengguna non-admin melihat "Request app" (INT-063).
  5. Rekomendasi app kontekstual dapat tampil di work item dan project settings, dan dapat dimatikan admin.
  6. Project page app tampil di sidebar proyek bila app terpasang dan display condition terpenuhi.
  7. Kategori katalog mencakup: time tracking, test management, reporting/BI, workflow extension & scripting, diagramming, portfolio/roadmapping, issue sync, document generation/PDF, checklist, connector CRM/support, admin & audit, security & compliance.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Navigasi di NAV-009 ([21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md)). Registry dapat berupa katalog lokal (self-hosted) atau katalog publik komunitas; lihat Pertanyaan terbuka.

### INT-062 App installation, updates & management (Manage apps)

- **Deskripsi:** Halaman admin untuk memasang, memperbarui, mengonfigurasi, dan uninstall app.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (Administer Jira atau peran app admin), developer app (development mode).
- **Aturan perilaku (acceptance criteria):**
  1. Settings > Apps > Manage apps menampilkan daftar app dengan filter User-installed, All, dan System.
  2. Detail app menampilkan versi, vendor, scope, lisensi, serta tombol Configure, Get started, dan Uninstall.
  3. App dipasang dari registry atau, dalam development mode, dengan mengunggah/menunjuk URL manifest.
  4. Update minor dipasang otomatis; update yang menambah scope atau egress ditahan sampai admin menyetujuinya dan app tetap berjalan di versi lama.
  5. Uninstall meminta konfirmasi dan menjelaskan bahwa entity properties dan custom field milik app dapat tertinggal; admin dapat memilih "Hapus data app".
  6. Pengaturan "Enable development mode" dan "Enable private listings" tersedia dan default nonaktif.
  7. App dapat dipasang lewat link instalasi bertanda tangan atau CLI developer.
  8. Setiap instalasi, update, dan uninstall dicatat di audit log (siapa, kapan, versi, scope).
  9. Instalasi gagal menampilkan status galat dan alasan; app lintas produk menampilkan produk mana saja yang dipakainya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Halaman admin ADM-034 ([22 Administrasi Sistem](22-administrasi-sistem.md)); apps per proyek di PRJ-060.

### INT-063 App requests

- **Deskripsi:** Alur bagi pengguna non-admin untuk meminta app ke admin, yang meninjau lalu menyetujui atau menolak.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Semua pengguna (meminta), Admin Situs (meninjau).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Request app" meminta pesan alasan (wajib, maksimal 500 karakter).
  2. Admin menerima notifikasi email dan in-app untuk setiap permintaan baru.
  3. Halaman App requests menampilkan daftar peminta dan jumlah peminta per app, dengan aksi Approve (memasang) atau Dismiss.
  4. Peminta diberi tahu saat app dipasang atau permintaan ditolak.
  5. ⚠ Admin dapat menonaktifkan fitur request; tombol lalu tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Notifikasi lewat [19 Notifikasi & Email](19-notifikasi-email.md).

### INT-064 App licensing, trials & billing

- **Deskripsi:** Model lisensi app berbayar yang terikat jumlah pengguna, trial gratis, dan tagihan. Untuk produk open-source cukup flag "enabled per tenant" dan kunci lisensi opsional dari vendor.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Billing admin, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. App dapat menawarkan trial (default 30 hari) dengan tanggal berakhir terlihat di Manage apps.
  2. Tier jumlah pengguna lisensi app disamakan dengan jumlah pengguna produk host.
  3. App dapat gratis sepenuhnya atau gratis sampai jumlah pengguna tertentu.
  4. Tagihan bulanan/tahunan diproses oleh vendor atau penyelenggara layanan; produk ini hanya menyimpan status lisensi.
  5. Status lisensi (active, trial, expired) diteruskan ke app lewat konteks; saat expired extension point app disembunyikan dan admin diberi tahu.
  6. ⚠ Lisensi di sandbox dan situs Free mengikuti kebijakan vendor dan tidak dibatasi produk.
  7. Pembatalan langganan menghentikan app di akhir periode; kebijakan refund milik vendor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penagihan produk di ADM (langganan situs) [22 Administrasi Sistem](22-administrasi-sistem.md).

### INT-065 App permissions, scopes, consent & app users

- **Deskripsi:** Kontrol akses data app: scope di manifest, persetujuan admin, akun app user, dan project role khusus untuk membatasi akses. Inti model keamanan app produk ini.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Admin Situs, Org admin, semua pengguna (consent asUser dan OAuth 3LO), developer app.
- **Aturan perilaku (acceptance criteria):**
  1. Manifest mendeklarasikan scope classic (`read:work`, `write:work`, `read:user`, `manage:project`, `manage:configuration`, `manage:webhook`, `manage:data-provider`) dan/atau scope granular; token app hanya dapat memanggil endpoint yang sesuai scope, selebihnya 403 `insufficient_scope`.
  2. Layar consent saat instalasi merinci kemampuan app dalam bahasa awam per scope dan egress.
  3. Setiap app memiliki akun app user (bot) yang tampil di daftar pengguna dengan tanda "App" dan tidak memakan lisensi.
  4. Grup `app-addons-admin` dan project role `app-addons-project-access` ditambahkan ke permission scheme; admin dapat mencabut role tersebut per proyek sehingga app tidak dapat mengakses proyek itu.
  5. Pemanggilan asApp memakai izin app user; asUser memerlukan consent per pengguna ("Allow access") sebelum pertama kali dipakai.
  6. Issue security dan permission scheme berlaku untuk app user sama seperti pengguna biasa.
  7. Domain egress/remote yang dideklarasikan tampil di listing dan layar consent.
  8. Update yang menambah scope atau egress memerlukan persetujuan ulang admin (INT-062 butir 4).
  9. OAuth 2.0 (3LO) app yang diotorisasi pengguna dapat dilihat dan dicabut di profil pengguna; admin dapat memblokir lewat kebijakan "user-installed apps".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama scope final mengikuti [23 API Publik & Webhook](23-api-publik-webhook.md); App users di USR-029; App Permissions di PERM-048; consent di AUTH-013.

### INT-066 App data residency, security & trust signals

- **Deskripsi:** Kontrol admin atas lokasi data app dan sinyal keamanan untuk menilai risiko app.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Org admin, security admin.
- **Aturan perilaku (acceptance criteria):**
  1. App yang mendukung residency dapat di-pin/dipindahkan ke region yang sama dengan data produk.
  2. Badge "Berjalan di platform produk" diberikan pada app tanpa egress/remote sehingga data tidak keluar dan mengikuti residency produk.
  3. ⚠ Badge program keamanan dan reliabilitas ditampilkan bila vendor lolos peninjauan kurator registry.
  4. Tab Privacy & Security menampilkan data yang disimpan, lokasi, sub-processor, dan compliance yang dideklarasikan vendor.
  5. Kebijakan keamanan data dapat memuat app access rule yang memblokir app dari konten yang dicakup kebijakan.
  6. IP allowlisting memiliki opsi mengizinkan traffic dari app terpasang.
  7. ⚠ Data residency app tersedia di tier Standard ke atas sesuai entitlement.
  8. Aktivitas app (instalasi, akses data sensitif, pemblokiran) tercatat di audit log organisasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data residency di USR-042, IP allowlist USR-038 ([02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md)); data security policies di PERM-065.

### INT-067 Development mode, private apps & app distribution

- **Deskripsi:** Memasang app internal/privat yang tidak terdaftar di katalog untuk kustomisasi organisasi, beserta perangkat developer.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. CLI developer mendukung create, deploy, install, environment development/staging/production, dan tunnel lokal untuk menguji tanpa deploy.
  2. Link instalasi bertanda tangan dapat dibagikan untuk distribusi app privat ke situs tertentu.
  3. Mode development memungkinkan memasang manifest via URL/unggah, hanya bila "Enable development mode" aktif.
  4. Private listing di registry hanya terlihat oleh situs yang diundang vendor.
  5. Developer console menampilkan log, invocation, galat, dan metrik per app dan environment.
  6. Environment variable/secret dapat disimpan per environment dan tidak pernah ditampilkan kembali setelah disimpan.
  7. ⚠ Format manifest lama (padanan Connect) tidak didukung; hanya satu model manifest yang berlaku sejak awal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Runtime di INT-068.

### INT-068 App runtime & UI framework (Forge-like)

- **Deskripsi:** Platform runtime yang menjalankan app: UI deklaratif atau iframe sandbox, bridge API, backend function, dan pemanggilan API produk dengan autentikasi otomatis.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app.
- **Aturan perilaku (acceptance criteria):**
  1. Manifest YAML mendeklarasikan modul, resources, functions, scopes, egress, dan CSP; manifest divalidasi terhadap skema saat deploy dan galat dilaporkan per baris.
  2. UI Kit menyediakan komponen deklaratif dengan desain sistem produk ini, dirender oleh host.
  3. Custom UI berupa bundle statis di iframe sandbox (`sandbox="allow-scripts allow-forms"`, tanpa same-origin) dengan bridge postMessage: invoke, requestProduct, router, view.getContext, Modal, events.
  4. Backend function/resolver berjalan di luar proses server produk (layanan vendor atau runtime terisolasi), sesuai INT-060 butir 5.
  5. Mode remote memanggil backend vendor dengan token terkelola berumur pendek yang ditandatangani produk ini.
  6. `requestProduct` dapat berjalan asApp atau asUser sesuai consent (INT-065).
  7. Theming (light/dark) dan i18n diteruskan dari preferensi pengguna ke app lewat konteks.
  8. Batas ukuran bundle, timeout invocation, dan jumlah invocation per menit ditegakkan dan dilaporkan di developer console.
  9. App bergaya Connect (di-host vendor dengan JWT) tidak didukung terpisah; mode remote menjadi penggantinya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keamanan iframe dan CSP di [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md).

## Extension point app

### INT-069 Extension point: Issue view panels, glances, context & activity

- **Deskripsi:** Titik ekstensi di halaman work item tempat app menambahkan UI: panel konten, glance, section konteks, tab activity, dan background script.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, semua pengguna (melihat); ditampilkan hanya bila app punya akses ke proyek.
- **Aturan perilaku (acceptance criteria):**
  1. `issuePanel`: panel di bawah description yang ditambahkan lewat tombol Apps/quick-add, dapat di-collapse atau dihapus per work item oleh pengguna dengan Edit issues.
  2. `issueGlance`: baris ringkas di sidebar detail berisi label/badge yang membuka panel samping saat diklik.
  3. `issueContext`: section collapsible di area details dengan status collapse tersimpan per pengguna.
  4. `issueActivity`: tab tambahan di bagian Activity di samping Comments/History.
  5. `issueViewBackgroundScript`: logika tanpa UI yang berjalan saat work item dibuka (mis. memuat data untuk glance).
  6. Konteks yang diterima app: issueKey, issueId, projectKey, issueType, user (accountId, locale, timezone).
  7. Display conditions (proyek, work type, izin, ekspresi) dievaluasi server sebelum modul dirender.
  8. App menerima event refresh saat work item berubah sehingga panel dapat memuat ulang data.
  9. ⚠ Modul juga tampil di modal work item dan side panel board; di tampilan mobile hanya glance yang dirender.
  10. Tidak ada padanan modul Connect lama (webPanel, issue tab panel); hanya kontrak di atas yang didukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Slot UI berbasis manifest, iframe sandbox (INT-068). Layout work item di [06 Work Item Inti](06-work-item-inti.md).

### INT-070 Extension point: Custom field types from apps

- **Deskripsi:** App mendefinisikan tipe custom field sendiri dengan penyimpanan, rendering view/edit, validasi, nilai terhitung, dan pencarian JQL. Dijadwalkan paling akhir sesuai INT-060.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, Admin Situs (membuat field dari tipe app), semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia `customField` (satu instance) dan `customFieldType` (tipe yang dapat dibuat admin berkali-kali di Custom fields).
  2. Tipe data: string, number, date, datetime, user, group, object (JSON), dan collection dari tipe-tipe tersebut.
  3. Renderer view/edit app dipakai di create dialog, work item view, dan transition screen.
  4. Value function menghasilkan computed field read-only yang dihitung ulang saat work item berubah.
  5. Validation expression dievaluasi server saat simpan dan menolak nilai dengan pesan dari app.
  6. Formatter expression menghasilkan teks untuk list, ekspor, dan notifikasi tempat renderer iframe tidak dapat dipakai.
  7. Admin mengatur context proyek/work type dan konfigurasi per context seperti custom field bawaan.
  8. ⚠ Nilai dapat dicari via JQL; property object dapat diindeks per path yang dideklarasikan.
  9. Mendukung mode read-only, default value, dan schema object yang divalidasi.
  10. ⚠ Batas ukuran nilai (usulan 32 KB), dukungan team-managed, dan tampilan di kartu board terbatas pada formatter teks.
  11. Tidak ada padanan modul Connect lama (issue fields).
- **Team-managed vs company-managed:** Rilis pertama company-managed; team-managed menyusul (butir 10).
- **Catatan & ketergantungan:** Custom field di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md); JQL di [12 JQL](12-jql-bahasa-kueri.md).

### INT-071 Extension point: Workflow rules (validators, conditions, post functions)

- **Deskripsi:** App menambahkan aturan workflow kustom pada transition: condition, validator, dan post function.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, Admin Situs (mengonfigurasi workflow).
- **Aturan perilaku (acceptance criteria):**
  1. `workflowCondition` berbasis ekspresi menyembunyikan transition bila hasilnya false.
  2. `workflowValidator` berbasis ekspresi atau fungsi app menolak transition dengan pesan galat kustom yang tampil di layar transition.
  3. `workflowPostFunction` memanggil fungsi app secara asinkron setelah transition berhasil; kegagalan dicatat dan tidak membatalkan transition.
  4. UI konfigurasi rule (create/edit/view) dirender app di dalam workflow editor.
  5. Konfigurasi tersimpan per transition sebagai JSON dan ikut dalam ekspor/impor workflow.
  6. Hanya untuk workflow company-managed; team-managed memakai rules bawaan.
  7. Rule app didukung di editor workflow baru maupun tampilan lama.
  8. Tidak ada padanan modul Connect lama; hanya kontrak berbasis ekspresi/fungsi ini.
- **Team-managed vs company-managed:** Hanya company-managed.
- **Catatan & ketergantungan:** Rules bawaan di [09 Workflow](09-workflow-status-transisi.md) (termasuk WF-045 impor/ekspor).

### INT-072 Extension point: Custom JQL functions

- **Deskripsi:** App mendaftarkan fungsi JQL baru yang dievaluasi menjadi klausa JQL standar.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app, semua pengguna (memakai di pencarian).
- **Aturan perilaku (acceptance criteria):**
  1. `jqlFunction` dideklarasikan dengan nama, operator yang didukung (`in`, `not in`, `=`, `!=`), dan argumen.
  2. Fungsi menerima argumen dan konteks pengguna, lalu mengembalikan klausa pengganti (mis. `issue in (…)`) atau galat.
  3. ⚠ Hasil di-precompute dengan cache dan refresh berkala (usulan 15 menit), dan UI menandai bahwa hasil dapat usang.
  4. Fungsi app muncul di autocomplete editor JQL dengan deskripsi argumen.
  5. Batas waktu eksekusi (usulan 10 detik) dan ukuran klausa (usulan 10.000 ID) ditegakkan; melebihi batas menghasilkan galat.
  6. Galat fungsi ditampilkan dengan pesan ramah di UI pencarian tanpa menggagalkan seluruh halaman.
  7. Hanya model app baru yang mendukung modul ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fungsi JQL bawaan di [12 JQL](12-jql-bahasa-kueri.md).

### INT-073 Extension point: Dashboard gadgets from apps

- **Deskripsi:** App menyediakan gadget dashboard dengan mode view dan edit yang dapat ditambahkan pengguna.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, pemilik/editor dashboard.
- **Aturan perilaku (acceptance criteria):**
  1. `dashboardGadget` dideklarasikan dengan title, description, dan thumbnail, lalu tampil di galeri "Add gadget".
  2. Gadget memiliki mode edit (form konfigurasi) dan view; konfigurasi disimpan per instance gadget.
  3. Gadget mendukung refresh (manual/terjadwal) dan tinggi dinamis.
  4. ⚠ `dashboardBackgroundScript` menjalankan logika tingkat dashboard tanpa UI.
  5. Konteks yang diterima: dashboardId, gadgetId, user.
  6. Data yang dimuat gadget mengikuti izin viewer (asUser), bukan pemilik dashboard.
  7. Tidak ada padanan modul Connect lama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** DSH-058 ([17 Dashboard & Gadget](17-dashboard-gadget.md)).

### INT-074 Extension point: App pages (global, project, project settings, admin, personal settings)

- **Deskripsi:** Halaman penuh yang disediakan app di berbagai lokasi navigasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, semua pengguna, Admin Proyek, Admin Situs (sesuai lokasi).
- **Aturan perilaku (acceptance criteria):**
  1. `globalPage` tampil di menu Apps, dengan subpage opsional.
  2. `projectPage` tampil di sidebar proyek company-managed dan team-managed.
  3. `projectSettingsPage` tampil di Project settings dan hanya dapat dibuka Admin Proyek.
  4. `adminPage` tampil di Settings > Apps, termasuk configure page dan get started page yang ditautkan dari Manage apps.
  5. `personalSettingsPage` tampil di Personal settings.
  6. ⚠ Halaman app memiliki routing internal dengan deep link stabil (`/apps/<appId>/<page>/<subpath>`).
  7. Display conditions dievaluasi per halaman; bila tidak terpenuhi menu disembunyikan dan URL langsung menghasilkan 404.
  8. Tidak ada padanan modul Connect lama (general pages, admin pages, post-install page).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Routing SPA di [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md); webui `NoRoute` fallback tetap menangani deep link.

### INT-075 Extension point: UI modifications (UIM)

- **Deskripsi:** App mengubah perilaku field di form secara dinamis (hide, required, read-only, opsi, nilai) sesuai konteks dan input pengguna. Bukan kontrol keamanan; validasi server tetap berlaku.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app, Admin Situs (mengatur konteks), semua pengguna (terdampak di form).
- **Aturan perilaku (acceptance criteria):**
  1. `uiModifications` berlaku di Global create, work item view, dan layar transition.
  2. Konteks UIM (proyek, work type, view type) diatur lewat REST API UIM oleh app atau admin.
  3. App menerima callback `onInit` saat form dibuka dan `onChange` saat nilai field berubah.
  4. Field API: setName, setDescription, setVisible, setRequired, setReadOnly, setValue, setOptionsVisibility, getValue.
  5. Hanya field system dan custom field yang didukung; field dari app lain tidak dapat dimodifikasi.
  6. Bila dua app memodifikasi properti yang sama pada field yang sama, perubahan terakhir menang dan form menampilkan peringatan konflik kepada admin.
  7. ⚠ UIM berlaku di proyek company-managed dan team-managed.
  8. Server tetap memvalidasi semua input (required, read-only) berdasarkan konfigurasi field, terlepas dari UIM; menyembunyikan field tidak mencegah pengisian lewat API.
  9. ⚠ Batas jumlah konteks per app (usulan 100) ditegakkan dengan galat API.
- **Team-managed vs company-managed:** Tidak ada perbedaan (butir 7).
- **Catatan & ketergantungan:** Padanan "field behaviors" bawaan dapat dipertimbangkan di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md).

### INT-076 Extension point: Automation actions & triggers from apps

- **Deskripsi:** App memperluas Otomasi dengan aksi kustom, atau Otomasi memanggil app lewat web request/incoming webhook.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app, pembuat aturan otomasi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Modul `automationAction` dari app muncul sebagai action di pembuat aturan otomasi.
  2. ⚠ Action mendeklarasikan input schema dan output yang dapat dipakai sebagai smart value di langkah berikutnya.
  3. Pola umum tanpa modul: "Send web request" ke endpoint app dan trigger "Incoming webhook".
  4. Perubahan entity property/field milik app dapat menjadi trigger aturan.
  5. ⚠ Action AI dari app dapat dipakai agent AI dan otomasi dengan kontrak yang sama.
  6. Audit log otomasi mencatat eksekusi action app beserta durasi dan hasilnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** [20 Otomasi](20-otomasi.md) (AUTO-025, AUTO-051, AUTO-054).

### INT-077 Extension point: Entity properties & app storage

- **Deskripsi:** Penyimpanan key-value JSON yang dilekatkan ke entitas, dengan indexing untuk JQL, plus storage milik app. Ini primitif ekstensi tahap awal (INT-060 butir 1).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, developer app; baca mengikuti izin melihat entitas, tulis mengikuti izin edit entitas.
- **Aturan perilaku (acceptance criteria):**
  1. Entity properties tersedia pada work item, proyek, pengguna, comment, work type, workflow transition, dashboard item, board, dan sprint.
  2. ⚠ Setiap property maksimal 32 KB JSON valid; jumlah property per entitas dibatasi (usulan 100) dengan galat 400 saat terlampaui.
  3. App/admin dapat mendeklarasikan index path JSON menjadi alias JQL bertipe number, string, text, date, atau user.
  4. Property terlihat oleh siapa pun yang dapat melihat entitas; dokumentasi dan UI developer memperingatkan agar tidak menyimpan data rahasia.
  5. App properties global dan storage app (key-value, custom entities, secret storage terenkripsi) tersedia per app dan terisolasi antar app.
  6. REST CRUD property tersedia (list keys, get, set, delete) per entitas.
  7. ⚠ Saat work item di-clone, property ikut disalin; saat dipindah proyek, property dipertahankan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Disimpan sebagai kolom JSONB `properties` atau tabel `entity_properties` dengan index ekspresi; pencarian di JQL-036 ([12 JQL](12-jql-bahasa-kueri.md)); REST di [23 API Publik & Webhook](23-api-publik-webhook.md).

### INT-078 Extension point: Action buttons, menus & navigator export formats

- **Deskripsi:** App menambahkan item aksi di menu kontekstual dan format ekspor tambahan di navigator.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app, semua pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. `issueAction` menambahkan item di menu "..." work item yang membuka modal app.
  2. ⚠ `boardAction`, `backlogAction`, dan `sprintAction` menambahkan item di menu board, backlog, dan sprint.
  3. ⚠ `command` menambahkan perintah ke command palette.
  4. `searchExportFormat` menambahkan format ekspor tambahan di menu Export navigator; app menerima JQL dan mengembalikan file.
  5. `projectReport` menambahkan laporan app di halaman Reports proyek.
  6. ⚠ App dapat mendaftarkan keyboard shortcut yang tidak bentrok dengan shortcut bawaan (bentrok ditolak saat instalasi).
  7. Setiap item mendukung display conditions dan menerima konteks seleksi (work item terpilih, board, sprint).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** SRC-057 ([13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md)), RPT-037 ([16 Laporan & Insight](16-laporan-insight.md)).

### INT-079 Extension point: App-defined permissions & time tracking provider

- **Deskripsi:** App mendefinisikan permission proyek/global sendiri lewat permission scheme dan dapat menggantikan provider time tracking.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Developer app, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `projectPermission` menambahkan permission baru yang muncul di permission scheme dan dapat diberikan ke peran/grup.
  2. `globalPermission` menambahkan permission baru di Global permissions.
  3. App mengecek izin pengguna lewat REST `mypermissions` dengan key permission app.
  4. `timeTrackingProvider` memungkinkan admin memilih provider di Time tracking settings; bila provider app aktif, UI worklog bawaan diganti UI app.
  5. Hanya satu provider time tracking aktif per situs; uninstall app provider mengembalikan provider bawaan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) (PERM-048), [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md).

### INT-080 Extension point: Events, triggers & app lifecycle

- **Deskripsi:** Mekanisme app untuk bereaksi terhadap event domain dan siklus hidup app, termasuk trigger terjadwal dan web trigger.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Developer app; event hanya dikirim untuk entitas yang dapat diakses app user.
- **Aturan perilaku (acceptance criteria):**
  1. Product event trigger tersedia untuk event katalog domain (work item created, updated, deleted, commented, version created, sprint created, dan lainnya).
  2. ⚠ Manifest dapat memfilter event (proyek, work type, ekspresi); pengiriman gagal di-retry dengan backoff eksponensial hingga batas usulan 24 jam.
  3. Scheduled trigger dapat dijadwalkan per jam, hari, atau minggu.
  4. Web trigger menyediakan URL HTTP unik per app untuk dipanggil sistem eksternal.
  5. Event lifecycle installed, upgraded, dan uninstalled dikirim ke app.
  6. Webhook app dapat difilter JQL sama seperti webhook admin.
  7. Antrean event asinkron tersedia untuk pekerjaan berat agar trigger tidak timeout.
  8. Kuota timeout, memori, dan rate limit per app ditegakkan; pelanggaran tercatat di developer console.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog event di [28 Model Data Inti](28-model-data-inti.md); webhook di [23 API Publik & Webhook](23-api-publik-webhook.md).

## Pertanyaan terbuka

- Apakah produk ini akan membangun modul dokumentasi/Pages internal, atau cukup konektor ke penyedia wiki eksternal untuk INT-042 s.d. INT-046?
- Apakah importer "dari Jira" (via REST API instans sumber, termasuk history dan lampiran) perlu dinaikkan prioritasnya ke Fase 3 sebagai jalur adopsi utama, menggantikan peran JCMA (INT-026)?
- Apakah dibutuhkan exporter JSON migrasi per proyek yang round-trip dengan INT-016, di luar backup situs penuh?
- Apakah konektor Slack (INT-030 s.d. INT-034) yang berprioritas P1 perlu ditarik dari Fase 4 ke Fase 3 agar hadir bersama otomasi dan webhook?
- Registry app (INT-061) berbentuk katalog lokal per instans, katalog publik komunitas, atau keduanya, dan siapa yang mengkurasi badge keamanan?
- Apakah UI extension point berbasis iframe (INT-069 s.d. INT-075) termasuk peta jalan, atau ekstensi cukup dibatasi pada REST + webhook + entity properties?
- Nama perintah slash dan nama konektor chat mengikuti nama produk "Jari"; apakah prefiks `/jira` diganti sejak awal?

## Di luar cakupan & ekosistem

- App Marketplace pihak ketiga dan konektor buatan vendor: Zendesk Support for Jira, konektor Salesforce (ServiceRocket, Peeklogic, zAgile, Appfire), HubSpot, Intercom, Freshdesk, Zeplin, Adobe XD, Miro, Lucid, Sketch, extension JetBrains, serta app populer seperti Tempo, Xray, ScriptRunner (Behaviours), BigPicture, dan generator PDF; produk ini hanya menyediakan primitif integrasi (REST, webhook, remote link, entity properties, OAuth).
- Konektor milik vendor iPaaS (Zapier, Make, Power Automate, Workato, n8n) dibangun dan dirawat vendor tersebut (INT-003).
- Atlassian Marketplace, penagihan app terpusat Atlassian, App Migration Platform, dan model app Atlassian Connect/Forge asli; produk ini memakai model app sendiri (INT-060, INT-068).
- Fitur Data Center-only atau deprecated yang hanya dicatat sebagai rujukan: GitHub Issues importer (INT-019), importer tracker legacy Bugzilla/FogBugz/Mantis/Redmine/Trac/Pivotal Tracker (INT-020), Jira Cloud Migration Assistant (INT-026), Project import from backup (INT-027), dan Data Pipeline (INT-028).
- Layanan Atlassian lintas produk sebagai produk tersendiri: Confluence, Loom, Trello, Figma, Opsgenie (end-of-sale), Statuspage, Compass, Jira Align, Atlassian Analytics/Data Lake, Rovo, dan Atlassian Home; produk ini hanya membangun titik integrasi atau padanan minimal yang dijelaskan di INT-042 s.d. INT-059 (padanan Analytics/Data Lake di INT-029, dirujuk RPT-038 dan JSM-AI-066). Integrasi ke layanan Atlassian Home/Goals & Projects tidak dibangun; padanan native-nya adalah entitas Goal di PLAN-061/PLAN-062 dan JPD-INS-036 (INT-056).
- Ekspor PDF bawaan bermerek (INT-022 butir 9) dan importer khusus untuk ClickUp, Smartsheet, atau Excel (dialihkan ke CSV, INT-018 butir 8).
