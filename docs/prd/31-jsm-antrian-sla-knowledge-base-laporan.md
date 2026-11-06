# 31 — Service Management: Queues, SLA, Knowledge Base, CSAT & Laporan

Modul ini mendefinisikan sisi kerja agent di service project produk ini: Queues (daftar kerja berbasis JQL untuk triase, termasuk queue default, queue kustom, kolom, urutan, sidebar, count, bulk action, dan Multi-space work), issue view agent (Reply to customer vs internal note, request details, participants & organizations, customer context, triage, change request type, link/move/clone/raise on behalf, canned responses, AI assist, Workforce Management), mesin SLA (metric, kondisi start/pause/stop, goal berbasis JQL, kalender jam kerja, tampilan timer, fungsi JQL, trigger automation, recalculation; SLA team-managed hanya sebagai rujukan karena service project team-managed tidak dibangun), Knowledge base (knowledge space terhubung, visibilitas internal/eksternal, Articles view dan editor native, lifecycle artikel, versi & administrasi space, kategori portal, saran artikel & deflection, panel artikel agent, create article from request, feedback artikel, sumber pengetahuan eksternal), umpan balik pelanggan (CSAT dan Surveys native), serta laporan service project (default report, custom report builder, SLA report, export, dashboard & gadget JSM, analitik lanjutan). Definisi service project, request type, forms, tipe pengguna, dan permission default JSM ada di modul 29; portal, help center, customers, organizations, kanal email, dan customer notifications ada di modul 30; approvals dan proses ITSM ada di modul 32; Assets di modul 33; alert/on-call di modul 34; virtual service agent, AI JSM, trigger/action automation JSM, dan REST API JSM di modul 35. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga seluruh entitas modul ini (`queues`, `queue_columns`, `priority_groups`, `sla_metrics`, `sla_goals`, `sla_calendars`, `sla_cycles`, `kb_spaces`, `kb_articles`, `kb_article_versions`, `kb_categories`, `article_feedback`, `canned_responses`, `csat_ratings`, `service_reports`) adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/servicedesk/queue`, `.../sla`, `.../knowledge`, `.../csat`, `.../servicereport`) mengikuti pola Clean Architecture repo, dengan halaman Queues, Knowledge base, dan Reports sebagai route baru di SPA `frontend/`.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Administer/Browse Projects (PERM-020), Service Project Agent & Customer Permissions (PERM-036), issue security (PERM-042, PERM-043, PERM-044), Make Bulk Changes (PERM-011).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — custom field bertipe SLA dan Satisfaction, priority scheme, field configuration.
- [06 Work Item Inti](06-work-item-inti.md) — Issue Links (ISS-029), Clone (ISS-031), Move wizard (ISS-032), Bulk Change (ISS-035), Watchers (ISS-036), Restrict Work Item (ISS-042).
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — editor rich text untuk reply, canned response, dan artikel KB.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar, Upload Attachments (CMT-011), Attachment Settings and Limits (CMT-016).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — transition, validator, screen, Bulk transition (WF-013).
- [12 JQL](12-jql-bahasa-kueri.md) — ORDER BY (JQL-006), User & Group Functions (JQL-039), engine JQL tempat fungsi SLA didaftarkan.
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — List view (SRC-023), Bulk change from search results (SRC-031), Export CSV (SRC-050).
- [16 Laporan & Insight](16-laporan-insight.md) — Created vs resolved (RPT-022), Pie chart (RPT-023), Report sharing/export (RPT-036), Atlassian Analytics (RPT-040).
- [17 Dashboard & Gadget](17-dashboard-gadget.md) — Filter Results (DSH-031), Pie Chart (DSH-035), Created vs Resolved (DSH-039).
- [18 Timeline & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — Project List view (PLAN-016), Project Calendar view (PLAN-022), Board view (PLAN-026).
- [19 Notifikasi & Email](19-notifikasi-email.md) — Mention Notifications (NTF-013), Send Notifications Option on Bulk Change (NTF-028).
- [20 Otomasi](20-otomasi.md) — Scheduled trigger (AUTO-023), Issue smart values (AUTO-063), engine rule.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — sidebar proyek, Your work/Starred, keyboard shortcut.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST API inti dan webhook.
- [26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) — Forms builder (dasar Surveys native).
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — job latar, realtime, indexing, multi-tenant time zone.
- [28 Model Data Inti](28-model-data-inti.md) — katalog event domain (`issue.*`, `comment.*`) yang dikonsumsi mesin SLA.
- [29 Service Project, Request Type & Forms](29-jsm-service-project-request-type-formulir.md) — service project, template, work categories, request types, agent/collaborator/customer, Agent request view.
- [30 Help Center, Portal, Customers & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — help center, portal, customers, organizations, email channel, widget, customer notifications.
- [32 ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — approvals, incident/problem/change, metric SLA ITSM.
- [33 Assets (CMDB)](33-jsm-assets-cmdb.md) — field objek Assets di kolom queue dan goal SLA.
- [34 Operations](34-jsm-operations-alerting-on-call.md) — laporan Ops (MTTA/MTTR), jadwal on-call.
- [35 Virtual Service Agent, AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md) — JSM JQL Fields (JSM-AI-013), JSM JQL SLA Functions (JSM-AI-014), JSM Triggers (JSM-AI-016), JSM Smart Values (JSM-AI-020), REST API Knowledge Base (JSM-AI-029), AI Writing Assistance (JSM-AI-038), AI Agent Assist (JSM-AI-039), AI Knowledge Article Drafting (JSM-AI-040), Virtual Service Agent - AI Answers (JSM-AI-052).

## Persona terkait

- **Agent Layanan** — bekerja dari queue, membalas customer atau menulis internal note, memakai canned response, memantau timer SLA, mencari dan membagikan artikel KB, membuat artikel dari request.
- **Manajer Layanan/ITSM** — merancang metric dan goal SLA, memantau report SLA/CSAT/workload, menyusun priority group, mengatur Workforce Management.
- **Admin Proyek/Project Lead** — (sebagai service project admin) membuat dan mengatur queue, kolom, SLA, kalender, knowledge base, kategori, CSAT, dan custom report.
- **Customer/Pelapor** — mencari artikel di portal, terdeflect oleh saran artikel, menilai artikel, mengisi survei CSAT, melihat support timeline.
- **Developer/Anggota Tim** — (sebagai collaborator non-agent) membuka request, menulis internal note, dan menerima @mention dari agent.
- **Admin Situs** — mengatur kebijakan global (lisensi agent, AI opsional, sumber pengetahuan eksternal, akses analitik lanjutan).
- **Stakeholder/Viewer** — membaca dashboard dan report layanan tanpa mengubah konfigurasi.
- **Integrator/Developer API** — membaca data SLA, CSAT, dan artikel lewat REST API/ekspor untuk analisis di luar produk.

## User story modul

- Sebagai Agent Layanan, saya ingin melihat queue default yang sudah terurut berdasarkan sisa waktu SLA agar request paling mendesak selalu saya kerjakan lebih dulu (JSM-QSK-001, JSM-QSK-004).
- Sebagai Admin Proyek/Project Lead, saya ingin membuat queue kustom dengan builder Basic atau JQL lengkap dengan kolom pilihan agar setiap tim punya daftar kerja yang relevan (JSM-QSK-002, JSM-QSK-003).
- Sebagai Manajer Layanan/ITSM, saya ingin mengelompokkan queue ke dalam priority group agar agent bisa berganti fokus antar-kelompok kerja dengan satu klik (JSM-QSK-006).
- Sebagai Agent Layanan, saya ingin berpindah ke request berikutnya di queue tanpa kembali ke daftar dan melakukan aksi massal saat perlu agar triase cepat (JSM-QSK-005, JSM-QSK-008).
- Sebagai Agent Layanan, saya ingin membedakan dengan jelas antara Reply to customer dan internal note agar informasi internal tidak pernah bocor ke customer (JSM-QSK-012).
- Sebagai Agent Layanan, saya ingin menyisipkan canned response dengan variabel yang terisi otomatis agar balasan rutin konsisten dan cepat (JSM-QSK-019).
- Sebagai Agent Layanan, saya ingin melihat konteks customer dan organisasinya di issue view agar saya memahami riwayat dan kebutuhannya tanpa membuka halaman lain (JSM-QSK-015, JSM-QSK-011).
- Sebagai Manajer Layanan/ITSM, saya ingin mendefinisikan metric SLA dengan kondisi start/pause/stop dan goal berbasis JQL agar target layanan sesuai perjanjian dengan pelanggan (JSM-QSK-022, JSM-QSK-023, JSM-QSK-024).
- Sebagai Manajer Layanan/ITSM, saya ingin SLA hanya berjalan di jam kerja dan tidak berjalan di hari libur agar pengukuran adil bagi tim (JSM-QSK-025).
- Sebagai Agent Layanan, saya ingin timer SLA berwarna di issue view dan queue agar saya tahu kapan request mendekati breach (JSM-QSK-026).
- Sebagai Manajer Layanan/ITSM, saya ingin aturan automation yang bereaksi saat SLA akan breach agar eskalasi berjalan otomatis (JSM-QSK-028).
- Sebagai Customer/Pelapor, saya ingin melihat artikel yang relevan saat mengetik summary request agar saya bisa menyelesaikan masalah sendiri tanpa menunggu agent (JSM-QSK-035).
- Sebagai Agent Layanan, saya ingin mencari, membagikan, dan membuat artikel KB langsung dari request agar solusi terdokumentasi dan bisa dipakai ulang (JSM-QSK-036, JSM-QSK-037).
- Sebagai Customer/Pelapor, saya ingin memberi rating kepuasan setelah request selesai agar tim layanan tahu kualitas penanganannya (JSM-QSK-040).
- Sebagai Manajer Layanan/ITSM, saya ingin report Workload, Satisfaction, deflection, dan capaian SLA serta custom report berbasis series agar kinerja layanan bisa diukur dan diperbaiki (JSM-QSK-042, JSM-QSK-043, JSM-QSK-044).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 18 |
| P1 | 21 |
| P2 | 6 |
| P3 | 3 |
| **Total** | **48** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 0 |
| Fase 2 | 0 |
| Fase 3 | 0 |
| Fase 4 | 0 |
| Fase 5 | 42 |
| Fase 6 | 0 |
| Fase 7 | 5 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Catatan hitungan: 56 fitur jatah inventaris dipetakan ke 47 heading karena 9 fitur ringkas dari lampiran JSM digabung ke fitur rinci yang setara (lihat baris **Alias inventaris** dan Catatan masing-masing fitur), ditambah 1 fitur di luar inventaris bertanda ⚠ (JSM-QSK-048), sehingga total 48 heading. JSM-QSK-030 menjadi rujukan Fase — karena clone hanya membangun service project company-managed (JSM-PRJ-012).

Strategi urutan implementasi: Fase 5 dimulai dari mesin SLA (JSM-QSK-022 s.d. JSM-QSK-027) karena queue default, kolom, dan sort bergantung pada field SLA; mesin ini dibangun sebagai konsumen event domain work item (modul 28) yang menulis siklus SLA (`sla_cycles`) secara transaksional dan dijalankan ulang oleh job latar untuk kalender dan threshold. Paralel dengan itu dibangun Queues (JSM-QSK-001 s.d. JSM-QSK-005, JSM-QSK-009) di atas engine JQL modul 12 dengan scope project dipaksa di server, lalu issue view agent (JSM-QSK-011 s.d. JSM-QSK-016) yang memisahkan komentar publik dan internal sejak skema data pertama agar tidak ada kebocoran ke portal. Knowledge base dibangun sebagai knowledge space native produk ini (bukan bergantung pada produk wiki eksternal) dengan portal suggestion dan pencatatan deflection (JSM-QSK-031 s.d. JSM-QSK-038) serta lifecycle artikel, versi, dan administrasi space (JSM-QSK-048) yang sekaligus menjadi document store bawaan untuk Pages di modul lain, diikuti CSAT dan report default/custom karena keduanya membaca data yang sudah tercatat (SLA, rating, deflection). Fase 7 menambah kapabilitas lanjut yang bergantung pada AI opsional, data operasi, atau analitik: AI assist, Workforce Management, sumber pengetahuan eksternal, Surveys native, dan analitik lanjutan.

## Fitur

## Queues: definisi, kolom & urutan

### JSM-QSK-001 Queues - Default Queues

- **Deskripsi:** Queue bawaan yang dibuat otomatis saat service project dibuat; isinya bergantung pada template project (IT service management, General service, Customer service, HR, Facilities, dan lain-lain). Queue default adalah titik masuk pertama agent untuk triase.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer projects) mengubah/menghapus; Agent Layanan (lisensi agent + Browse projects) melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Saat service project dibuat, sistem membuat queue default sesuai template: minimal All open, Assigned to me, Unassigned, Waiting on support, Waiting on customer, Recently resolved; template ITSM menambah queue per work category (Service requests, Incidents, Problems, Changes, Post-incident reviews).
  2. JQL setiap queue default disimpan tanpa klausa project, dan server selalu menambahkan `project = <key>` secara implisit saat mengeksekusi sehingga queue tidak pernah memuat work item project lain.
  3. Sort default queue berbasis SLA, mis. `ORDER BY "Time to resolution" ASC`, sehingga request dengan sisa waktu paling sedikit atau sudah breach tampil paling atas.
  4. Admin proyek dapat mengubah nama, JQL, kolom, dan sort queue default, mengurutkannya ulang, atau menghapusnya dengan perilaku identik queue kustom; queue default tidak punya status terkunci.
  5. Queue 'Assigned to me' memakai `assignee = currentUser()` sehingga isinya dievaluasi per agent yang sedang membuka dan count-nya berbeda per agent.
  6. Di project ITSM, queue default tampil dikelompokkan di bawah work category masing-masing di sidebar (lihat JSM-QSK-006).
  7. Kolom default queue: Type, Key, Summary, Reporter, Assignee, Status, Created, Time to first response, Time to resolution.
  8. ⚠ Bila project tidak memiliki queue sama sekali, halaman Queues menampilkan empty state dengan tombol "Create queue" untuk admin dan pesan "Belum ada queue; hubungi admin proyek" untuk agent.
  9. Template menyertakan queue 'Breached SLAs' bawaan dengan JQL `"Time to resolution" = breached() OR "Time to first response" = breached()` (dari jatah SLA JQL & breach handling).
  10. Template general/customer service selalu membuat service project company-managed beserta queue default yang sama; opsi team-managed tidak ditawarkan (JSM-PRJ-012), sehingga pengelolaan queue team-managed dari sidebar tidak dibangun dan queue project team-managed hasil impor dipetakan ke Queue settings company-managed (dari jatah Queues ringkas).
  11. Setiap template menyimpan daftar queue default (nama, JQL, kolom, sort) sebagai data seed, dan daftar itu tampil sebagai pratinjau di pemilih template project.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Template project dan work category didefinisikan di modul 29. Daftar persis queue default per template harus disimpan sebagai data seed template (bukan hard-code) agar bisa diubah tanpa rilis. Field SLA untuk sort berasal dari JSM-QSK-022.

### JSM-QSK-002 Queues - Create/Edit Custom Queue (Basic & JQL)

- **Alias inventaris:** Queues
- **Deskripsi:** Admin proyek membuat dan mengubah queue memakai filter mode Basic (pemilih field) atau JQL lanjutan untuk menyaring request di service project. Fitur ini juga menjadi rumah fitur ringkas "Queues" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer projects); agent tidak dapat membuat queue shared.
- **Aturan perilaku (acceptance criteria):**
  1. Dialog Create queue memuat nama queue (wajib, maksimal 255 karakter, unik case-insensitive di dalam project; galat "Nama queue sudah dipakai"), work category (khusus ITSM), filter, kolom, dan urutan.
  2. Mode Basic menyediakan dropdown Request type, Status, Assignee, Priority, Resolution, Labels, Organizations, dan field lain lewat '+ More'; pilihan diterjemahkan ke JQL setara.
  3. Mode JQL menyediakan editor dengan autocomplete field/fungsi/nilai dan validasi sintaks; galat ditampilkan inline di bawah editor dan tombol Save nonaktif selama JQL tidak valid.
  4. Pengguna bisa pindah Basic ke JQL kapan saja; pindah JQL ke Basic hanya diizinkan bila JQL dapat direpresentasikan builder, selain itu tampil pesan "Kueri terlalu kompleks untuk mode Basic".
  5. Scope project dipaksa di server: klausa project lain di JQL diabaikan/di-AND-kan dengan `project = <key>` sehingga hasil hanya work item project ini.
  6. JQL queue mendukung field/fungsi khas JSM: "Customer Request Type"/"Request Type", Organizations, "Request participants", Satisfaction, request-channel-type, `approver()`, `pending()`, fungsi SLA (JSM-QSK-027), `currentUser()`, `membersOf()`.
  7. Selama queue disusun, panel preview menampilkan daftar hasil (halaman pertama) dan jumlah total hasil, diperbarui setiap kali filter berubah.
  8. Edit queue tersedia lewat menu '...' > Edit queue; hapus lewat menu yang sama dengan dialog konfirmasi yang menyebut nama queue; menghapus queue tidak menghapus work item.
  9. Batas total 300 queue per service project per work category, dihitung lintas semua section; saat batas tercapai tombol Create nonaktif dengan pesan batas.
  10. Queue tidak bisa lintas project; untuk kebutuhan lintas project pengguna diarahkan ke filter/dashboard inti atau Multi-space work (JSM-QSK-010).
  11. ⚠ Dialog pembuatan queue versi sederhana khas team-managed (nama, filter Basic/JQL, kolom, tanpa pemilihan work category) tidak dibangun; semua service project memakai dialog Queue settings company-managed, dan pemilihan work category bersifat opsional di dialog tersebut (JSM-PRJ-012).
  12. ⚠ Aksi "Duplicate queue" menyalin nama (diberi akhiran "(copy)"), JQL, kolom, dan sort ke queue baru; bila keputusan produk menyatakan tidak native, menu ini disembunyikan.
  13. ⚠ Queue bersifat shared untuk semua agent di project; konsep queue personal tidak disediakan di versi pertama (dari jatah Queues ringkas, butir kepemilikan) dan dicatat sebagai pertanyaan terbuka.
  14. Queue disimpan dengan empat elemen: nama, filter (Basic atau JQL), kolom, dan urutan; queue tanpa kolom ditolak dengan galat "Pilih minimal satu kolom".
  15. ⚠ Queue settings menampilkan penghitung kuota "n / 300" per work category agar admin mengetahui sisa batas queue project.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Bergantung pada engine JQL modul 12 dan field JSM modul 29/35. Penempatan sub-fitur fitur ringkas "Queues" (lampiran JSM): membuat queue (nama, filter, kolom, urutan) → butir 14 (rincian di butir 1-8); queue default per template → JSM-QSK-001 butir 11; kolom → JSM-QSK-003 butir 9; order by termasuk sisa waktu SLA → JSM-QSK-004 butir 7; jumlah request di sidebar dengan batas tampilan → JSM-QSK-007 butir 8; starred queues → JSM-QSK-006 butir 12; pengelompokan per work category → JSM-QSK-006 butir 13; hide queue per user → JSM-QSK-006 butir 11; kepemilikan shared/personal → butir 13; navigasi next/previous → JSM-QSK-005 butir 11; bulk change → JSM-QSK-008 butir 8; indikator data baru → JSM-QSK-007 butir 9; akses agent/admin/collaborator → JSM-QSK-009 butir 7; batas jumlah queue → butir 15; team-managed → JSM-QSK-001 butir 10. Alert native untuk request baru di queue tidak ada; pola alert lewat automation dijelaskan di JSM-QSK-028.

### JSM-QSK-003 Queues - Columns Configuration

- **Deskripsi:** Konfigurasi kolom yang tampil di daftar request sebuah queue, termasuk field khas JSM seperti SLA, Request type, Organizations, dan Satisfaction.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek mengatur kolom; Agent Layanan hanya melihat (dan boleh resize lokal).
- **Aturan perilaku (acceptance criteria):**
  1. Editor queue memungkinkan tambah, hapus, dan urutkan ulang kolom dengan drag and drop; urutan disimpan saat queue disimpan.
  2. Kolom field sistem yang tersedia: Key, Summary, Status, Assignee, Reporter, Priority, Created, Updated, Due date, Resolution, Labels, Components.
  3. Kolom field JSM yang tersedia: Request type (ikon + nama), Request participants, Organizations, Satisfaction, Request channel, Approvals, Pending reason (ITSM).
  4. Setiap metric SLA (Time to first response, Time to resolution, SLA kustom) bisa dijadikan kolom yang menampilkan sisa waktu atau waktu terlampaui dengan format ringkas (lihat JSM-QSK-026).
  5. Custom field apa pun yang ada di field configuration project, termasuk field objek Assets (modul 33), bisa dijadikan kolom.
  6. Konfigurasi kolom disimpan per queue dan berlaku untuk semua agent; tidak ada preferensi kolom personal.
  7. ⚠ Satu queue dibatasi maksimal 30 kolom; saat batas tercapai daftar field menampilkan pesan batas.
  8. ⚠ Agent dapat mengubah lebar kolom dengan menarik batas header; lebar disimpan per pengguna di browser dan tidak mengubah konfigurasi queue.
  9. Kolom Satisfaction, Organizations, priority, dan custom field termasuk daftar kolom minimum yang harus tersedia sejak Fase 5 (dari jatah Queues ringkas, butir kolom).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kolom yang field-nya dihapus dari project disembunyikan otomatis dan dicatat di log konfigurasi. Rendering kolom memakai komponen list view yang sama dengan SRC-023 (modul 13).

### JSM-QSK-004 Queues - Sorting & Ordering

- **Deskripsi:** Urutan request di queue ditentukan oleh ORDER BY di JQL atau pilihan sort mode Basic, dan agent dapat mengurutkan sementara lewat header kolom.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek menetapkan urutan default; Agent Layanan memakai sort sementara.
- **Aturan perilaku (acceptance criteria):**
  1. Urutan default queue mengikuti ORDER BY di JQL dan mendukung multi-level, mis. `ORDER BY "Time to resolution" ASC, priority DESC, created ASC`.
  2. Klik header kolom mengurutkan sementara ascending, klik kedua descending; urutan sementara hanya berlaku untuk tampilan agent tersebut dan tidak mengubah JQL queue.
  3. Sort berdasarkan field SLA memakai sisa waktu siklus berjalan; nilai breach (negatif) tampil paling atas pada ASC, diikuti sisa waktu terkecil.
  4. ⚠ Work item tanpa siklus SLA aktif (nilai null) ditempatkan di bawah semua work item bernilai pada ASC maupun DESC.
  5. Tersedia aksi "Reset to default order" yang mengembalikan ke urutan JQL queue dan menghapus sort sementara.
  6. ⚠ Daftar queue dipaginasi dengan ukuran halaman 50 baris dan navigasi halaman; urutan tetap konsisten antar-halaman dengan tie-breaker `key ASC`.
  7. Order by berdasarkan sisa waktu SLA tersedia di builder Basic sebagai pilihan sort tanpa menulis JQL (dari jatah Queues ringkas).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ORDER BY mengikuti JQL-006 (modul 12); sort SLA memerlukan kolom terindeks `remaining_ms` pada tabel siklus aktif (lihat JSM-QSK-027).

### JSM-QSK-005 Queues - Queue List View, Inline Edit & Navigation

- **Deskripsi:** Tampilan daftar request di queue, edit cepat di baris, serta cara agent membuka request dan berpindah antar-request tanpa kembali ke daftar.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan; edit inline mengikuti Edit issues/Assign issues/Transition issues.
- **Aturan perilaku (acceptance criteria):**
  1. Klik baris request membuka issue view agent dengan sidebar queue tetap tampil dan breadcrumb kembali ke queue asal.
  2. Issue view yang dibuka dari queue menyediakan tombol Next/Previous yang mengikuti urutan queue saat dibuka, termasuk sort sementara.
  3. ⚠ Field Assignee, Status, dan Priority dapat diedit inline langsung di baris queue sesuai permission; perubahan Status menjalankan transition dan membuka screen bila workflow mewajibkannya.
  4. ⚠ Kotak "Filter this queue" menyaring sementara isi queue berdasarkan teks (key/summary) tanpa mengubah JQL queue.
  5. Aksi 'View in search' membuka JQL queue (beserta scope project) di navigator inti untuk export CSV, bulk change, atau simpan filter.
  6. Export memakai mekanisme ekspor inti (CSV all fields/current fields, SRC-050) dan hanya memuat work item yang bisa dilihat agent, dengan batas 1000 baris per ekspor.
  7. ⚠ Tombol "Export" langsung di halaman queue tersedia sebagai jalan pintas yang memanggil ekspor inti dengan kolom queue sebagai current fields.
  8. Queue kosong menampilkan empty state "Tidak ada request di queue ini" beserta ringkasan JQL queue.
  9. Kolom SLA memakai penanda visual: kuning saat sisa waktu di bawah ambang "mendekati breach", merah saat breached.
  10. ⚠ Keyboard shortcut: `j`/`k` untuk baris berikut/sebelumnya, `o`/Enter membuka request, `u` kembali ke queue; daftar shortcut tampil di dialog bantuan.
  11. Tombol Next/Previous nonaktif pada request pertama/terakhir queue dan tetap berfungsi melewati batas halaman paginasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue view agent didefinisikan di modul 29 (Agent request view) dan diperkaya JSM-QSK-011 s.d. JSM-QSK-018. Shortcut mengikuti konvensi modul 21.

## Queues: sidebar, count, aksi massal & visibilitas

### JSM-QSK-006 Queues - Sidebar Sections, Starred Queues & Priority Groups

- **Deskripsi:** Pengorganisasian queue di sidebar service project: queue berbintang yang personal, priority group yang dikelola admin dan dipilih agent, section 'View all queues', serta pengelompokan per work category.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (star, memilih group aktif); Admin Proyek (membuat/mengatur priority group dan urutan queue).
- **Aturan perilaku (acceptance criteria):**
  1. Agent memberi bintang pada queue lewat ikon bintang yang muncul saat hover; bintang bersifat personal dan tidak mengubah tampilan agent lain.
  2. Queue berbintang juga muncul di halaman 'Your work' pada tab Starred (modul 21).
  3. Admin proyek membuat priority group lewat Queues > Queue settings > tab Priority group > Create group dengan nama wajib dan unik per project.
  4. Admin menetapkan queue ke group dari tab Queues; satu queue boleh menjadi anggota lebih dari satu priority group.
  5. Agent memilih satu priority group aktif dari pemilih di sidebar; sidebar lalu menampilkan queue anggota group itu; pilihan disimpan per agent.
  6. ⚠ Admin dapat mengganti nama, mengurutkan, dan menghapus priority group; menghapus group tidak menghapus queue anggotanya.
  7. Section 'View all queues' berisi semua queue, default collapsed, dan count-nya tidak di-refresh otomatis.
  8. Di project ITSM, queue dikelompokkan per work category (Service requests, Incidents, Problems, Changes, Post-incident reviews).
  9. Urutan queue diatur admin lewat drag and drop di pengaturan queue dan berlaku untuk semua agent.
  10. ⚠ Status expand/collapse setiap section sidebar disimpan per pengguna dan dipulihkan saat kembali.
  11. ⚠ Agent dapat menyembunyikan queue dari sidebar miliknya sendiri (hide per user) tanpa memengaruhi agent lain; queue tersembunyi tetap ada di 'View all queues' (dari jatah Queues ringkas).
  12. ⚠ Penyematan queue favorit per pengguna dinamai "Starred queues" secara konsisten di UI dan API.
  13. Pengelompokan per work category hanya muncul di project dengan work category aktif; project non-ITSM menampilkan daftar queue datar.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Priority group adalah pengelompokan tampilan, bukan restriksi akses (lihat JSM-QSK-009).

### JSM-QSK-007 Queues - Work Item Counts & Refresh

- **Deskripsi:** Menampilkan jumlah request di setiap queue dan memperbarui isinya secara berkala.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan; count dihitung dengan identitas agent yang melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap queue di sidebar menampilkan badge jumlah work item yang cocok.
  2. Count queue di section Starred dan priority group aktif diperbarui lewat sync berkala (target interval 60 detik) tanpa reload halaman.
  3. Queue di 'View all queues' tidak di-refresh otomatis; count-nya diperbarui saat queue dibuka atau section diperluas.
  4. ⚠ Count di atas 999 ditampilkan sebagai '999+'.
  5. Count dihitung asinkron (job/cache) sehingga boleh tertinggal dari kondisi real-time maksimal satu interval sync; UI tidak memblokir render menunggu count.
  6. ⚠ Saat queue sedang dibuka dan isinya berubah, tampil indikator "Ada pembaruan" dengan tombol Refresh; klik memuat ulang daftar tanpa kehilangan sort sementara.
  7. Count menghormati permission agent (Browse projects dan issue security level), sehingga dua agent bisa melihat angka berbeda untuk queue yang sama.
  8. Badge jumlah request di sidebar hanya dihitung untuk queue di section yang sedang terlihat (Starred dan priority group aktif) dan memakai batas tampilan '999+' dari butir 4 (dari jatah Queues ringkas).
  9. ⚠ Indikator data baru menampilkan jumlah request baru sejak queue terakhir dimuat (mis. "3 request baru").
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mekanisme realtime/polling mengikuti modul 27. Count per agent perlu cache berkunci (project, queue, user-security-hash) agar biaya JQL terkendali.

### JSM-QSK-008 Queues - Bulk Actions

- **Deskripsi:** Aksi massal pada request di queue, baik lewat seleksi di daftar queue maupun lewat navigator inti; mesin bulk change memakai modul inti.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan dengan Make Bulk Changes (PERM-011) plus permission per aksi.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Daftar queue menyediakan checkbox per baris dan "select all di halaman ini" untuk memilih banyak request.
  2. Aksi massal yang tersedia: Assign, Transition, Edit field, Comment, Move, Delete, Watch/Unwatch, semuanya diteruskan ke wizard Bulk change inti (ISS-035).
  3. Jalan pintas "Bulk change all" membuka JQL queue di navigator lalu menjalankan Bulk change untuk semua hasil.
  4. Work item yang tidak memenuhi permission dilewati dan dilaporkan di ringkasan hasil (jumlah berhasil, dilewati, beserta alasannya).
  5. Bulk transition tetap menjalankan validator dan screen workflow; work item yang gagal validator dilaporkan per item.
  6. Satu operasi bulk dibatasi maksimal 1000 work item sesuai batas inti.
  7. Bulk comment yang bersifat publik (reply to customer) wajib menampilkan opsi "Kirim notifikasi ke customer" yang default tidak dicentang, dengan peringatan jumlah customer penerima.
  8. Aksi assign, transition, dan edit massal dapat dimulai dari konteks queue tanpa membuka navigator terlebih dahulu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** NTF-028 (modul 19) untuk opsi notifikasi; SRC-031 (modul 13).

### JSM-QSK-009 Queues - Permissions & Visibility

- **Deskripsi:** Aturan siapa yang boleh melihat, membuat, mengubah, dan menghapus queue, serta melihat request di dalamnya.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs, Admin Proyek, Agent Layanan, collaborator.
- **Aturan perilaku (acceptance criteria):**
  1. Hanya agent berlisensi dan admin yang dapat mengakses halaman Queues; customer tidak pernah melihat queue di portal maupun API customer (HTTP 403).
  2. ⚠ Collaborator (pengguna tanpa lisensi agent) dapat membuka request lewat tautan langsung, tetapi halaman Queues menampilkan pesan "Queue hanya tersedia untuk agent".
  3. Membuat, mengubah, menghapus, dan mengelompokkan queue hanya untuk pemegang Administer projects; endpoint menolak pengguna lain dengan HTTP 403.
  4. Isi dan count queue difilter otomatis oleh Browse projects dan issue security level di server.
  5. Star dan pilihan priority group aktif bersifat personal per agent dan tidak bisa diatur admin untuk agent lain.
  6. ⚠ Tidak ada restriksi visibilitas queue per tim/grup; priority group hanya pengelompokan, sehingga agent mana pun dapat membuka queue mana pun di project.
  7. ⚠ Akses queue untuk agent/admin dan pembatasan bagi collaborator dikonfigurasi lewat permission default service project (modul 29), bukan pengaturan per queue (dari jatah Queues ringkas).
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** PERM-020, PERM-036, PERM-044 (modul 03).

### JSM-QSK-010 Multi-space work (Views: List, Board, Calendar)

- **Deskripsi:** View list, board, dan calendar untuk work item service, baik per queue maupun lintas beberapa service project dalam satu daftar terpadu (sebelumnya disebut Views).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan; Admin Proyek mengaktifkan view; Admin Situs untuk ketersediaan lintas project.
- **Aturan perilaku (acceptance criteria):**
  1. Admin proyek mengaktifkan view yang tersedia lewat "Select views" (list, board, calendar); mematikan view berlaku untuk seluruh project.
  2. Board/calendar view dapat ditambahkan atau dihapus per queue; list view selalu ada dan menjadi default setiap queue.
  3. Board view menampilkan kartu per kolom status workflow; drag kartu antar-kolom menjalankan transition dengan validator dan screen tetap berlaku.
  4. ⚠ Calendar view menampilkan dan menjadwalkan work item berdasarkan field tanggal Due date atau Start date; drag kartu ke tanggal lain mengubah field tersebut.
  5. List view multi-space menampilkan daftar terpadu work item dari banyak service project dengan filter "Space" (multi-pilih).
  6. Pengguna dapat berpindah antar-view atas data (filter) yang sama tanpa kehilangan filter.
  7. List di Queues selalu ber-scope satu project, sedangkan list di Multi-space work lintas project; UI menandai scope aktif dengan jelas.
  8. Isi semua view menghormati Browse projects dan issue security level di setiap project sumber.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Komponen view dipakai ulang dari modul 18 (PLAN-016, PLAN-022, PLAN-026). Batas jumlah project dalam satu daftar multi-space menjadi pertanyaan terbuka.

## Issue view agent & kolaborasi

### JSM-QSK-011 Agent Issue View - Request Details Panel

- **Deskripsi:** Panel detail khas service management di issue view agent yang memuat atribut request (request type, reporter, channel, satisfaction, SLA, approvals, dan field layout request type).
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Browse projects); mengganti reporter memerlukan Modify reporter (PERM-024).
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan Request type dengan ikon dan nama portal tempat request type itu berada.
  2. Bila work item tidak punya request type, panel menampilkan peringatan "Customer tidak dapat melihat request ini di portal" dengan tautan ke aksi Change request type (JSM-QSK-017).
  3. Field Reporter ("Raised by") dapat diganti ke customer lain; perubahan tercatat di history dan reporter baru menerima akses portal ke request.
  4. Channel/sumber request (portal, email, chat, API, agent, widget) ditampilkan sebagai label baca-saja yang diisi saat request dibuat.
  5. Setelah request resolved dan customer menilai, panel menampilkan field Satisfaction (bintang) dan komentar CSAT.
  6. Panel memuat sub-panel SLA (JSM-QSK-026) dan Approvals (modul 32) bila relevan bagi request.
  7. Field lain ditampilkan sesuai layout request type (field yang terlihat customer dan field internal dibedakan dengan label).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Layout issue view agent didefinisikan di modul 29 (JSM-PRJ-052 Agent issue view, JSM-PRJ-054 Request details and request participants, JSM-PRJ-057 SLA panel in issue view); channel dikirim oleh modul 30 (email channel, widget) dan modul 23/35 (API).

### JSM-QSK-012 Agent Issue View - Reply to Customer vs Internal Note

- **Deskripsi:** Dua jenis komentar di issue view agent: reply yang dilihat dan dinotifikasikan ke customer, serta internal note yang hanya terlihat tim, termasuk aturan visibilitas lampiran.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Add comments + lisensi agent untuk reply); collaborator hanya internal note; Customer membaca reply.
- **Aturan perilaku (acceptance criteria):**
  1. Editor komentar menyediakan dua tab/tombol 'Reply to customer' dan 'Add internal note' dengan warna latar dan label berbeda; mode terakhir tidak diingat otomatis untuk mencegah salah kirim (default: internal note).
  2. Reply dikirim ke reporter dan request participants lewat email (modul 30) dan tampil di halaman request portal.
  3. Internal note diberi label 'Internal' dan tidak pernah tampil di portal, email customer, maupun API customer; pemeriksaan visibilitas dilakukan di server.
  4. Saat menulis reply, editor menampilkan peringatan "Komentar ini akan terlihat oleh customer".
  5. ⚠ Visibilitas komentar tidak dapat diubah setelah dikirim; untuk mengoreksi, agent menghapus lalu membuat komentar baru dengan jenis yang benar.
  6. Balasan customer lewat email ditambahkan otomatis sebagai komentar publik atas nama customer tersebut.
  7. Filter aktivitas menyediakan All, Comments, History, Work log, Approvals.
  8. ⚠ Edit/hapus komentar mengikuti permission Edit/Delete own/all comments; mengedit komentar publik tidak mengirim ulang email ke customer.
  9. Lampiran yang dikirim customer saat raise request (portal/email) tampil di request details portal.
  10. Lampiran di 'Reply to customer' terlihat customer dan ikut notifikasi email (sebagai tautan atau attachment sesuai ukuran); lampiran di internal note hanya terlihat agent/collaborator.
  11. ⚠ Lampiran yang diunggah agent langsung ke panel Attachments (bukan lewat komentar) diperlakukan internal dan tidak tampil di portal sampai dibagikan lewat reply.
  12. ⚠ Setiap lampiran menampilkan indikator "Internal" atau "Dibagikan ke customer" di panel Attachments.
  13. Lampiran mengikuti batas ukuran global (CMT-016); gambar inline dan lampiran dari email masuk (termasuk dari participant/CC) disimpan, sedangkan file dari pengirim yang diblokir ditolak dan dicatat di log kanal email.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Skema komentar wajib punya kolom `visibility` (`public`/`internal`) sejak migrasi pertama; modul 08 (komentar inti) dan JSM-PRJ-053 Reply to customer vs Add internal note di modul 29 harus memakai model yang sama.

### JSM-QSK-013 Agent Issue View - Request Participants & Organizations Sharing

- **Deskripsi:** Menentukan siapa saja di sisi customer, selain reporter, yang dapat melihat request dan menerima notifikasi.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Edit issues); Customer (berbagi dari portal sesuai customer permissions).
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat menambah dan menghapus request participants (customer atau agent) dari issue view; perubahan tercatat di history.
  2. Agent dapat menambah dan menghapus Organizations pada request sehingga seluruh anggota organisasi dapat melihat request di portal.
  3. Participants menerima notifikasi setiap reply publik dan perubahan status customer-facing.
  4. ⚠ Customer dapat membagikan request dari portal ke customer lain atau organisasi miliknya sesuai pengaturan sharing di modul 30.
  5. Pencarian customer saat menambah participant dibatasi oleh customer permissions project (mis. hanya customer di organisasi yang sama bila pengaturan itu aktif).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request participants sebagai entitas didefinisikan di JSM-PRJ-054 Request details and request participants ([29](29-jsm-service-project-request-type-formulir.md)); sisi customer (berbagi request dengan customer lain atau organization dari portal) di JSM-PRT-030 Request participants & sharing ([30](30-jsm-help-center-portal-pelanggan-kanal.md)), beserta customer permissions dan organizations di modul 30. Fitur ini hanya mendefinisikan perilaku panel di agent issue view; aturan siapa yang boleh ditambahkan dan notifikasi participant mengikuti JSM-PRJ-054 dan JSM-PRT-030 bila ada perbedaan.

### JSM-QSK-014 Agent Issue View - @Mentions & Collaborators

- **Deskripsi:** Kolaborasi internal lewat mention dan keterlibatan pengguna non-agent (collaborator).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, collaborator (Developer/Anggota Tim).
- **Aturan perilaku (acceptance criteria):**
  1. @mention agent atau pengguna internal di internal note mengirim notifikasi mention (NTF-013).
  2. Collaborator (pengguna berlisensi produk inti tanpa lisensi agent) dapat melihat request dan menambah internal note tanpa memakai kursi agent.
  3. ⚠ Collaborator tidak dapat membuat reply ke customer; kemampuan transition dan edit field mengikuti permission scheme project.
  4. Watchers internal ditampilkan terpisah dari request participants yang customer-facing.
  5. ⚠ Mention customer di internal note tidak mengirim notifikasi apa pun ke customer.
  6. Bila orang yang di-mention tidak punya akses ke request, editor menampilkan peringatan "<nama> tidak dapat melihat request ini" sebelum disimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tipe pengguna agent/collaborator/customer didefinisikan di modul 29; mention di editor di modul 07.

### JSM-QSK-015 Agent Issue View - Customer Context Panel

- **Deskripsi:** Informasi konteks tentang customer pelapor agar agent cepat memahami kebutuhannya.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Browse projects; detail customer mengikuti hak akses halaman Customers).
- **Aturan perilaku (acceptance criteria):**
  1. Panel menampilkan profile card customer: nama, email, avatar.
  2. Panel menampilkan organisasi customer beserta detail organisasi.
  3. Customer details fields yang dikonfigurasi admin (mis. tier, lokasi) tampil sebagai pasangan label-nilai.
  4. Panel menampilkan daftar recent requests dari customer yang sama (maksimal 5 terbaru yang bisa dilihat agent) dengan tautan.
  5. Klik nama customer membuka profilnya di halaman Customers (modul 30).
  6. ⚠ Panel menampilkan local time dan time zone customer bila profil customer memilikinya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data customer/organization dan customer details fields dikelola di modul 30; panel ini hanya tampilan dan memakai kerangka JSM-PRJ-056 Customer context panel (modul 29).

## Produktivitas agent

### JSM-QSK-016 Agent Productivity - Assign, Escalate & Triage

- **Deskripsi:** Aksi triase yang sering dipakai agent: ambil request, assign, eskalasi, dan atur prioritas.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Assign issues, Assignable user, Transition issues, Edit issues).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol 'Assign to me' satu klik di issue view dan baris queue; disembunyikan bila agent tidak punya Assignable user.
  2. Pemilih assignee hanya menampilkan pengguna dengan permission Assignable user di project.
  3. Eskalasi dilakukan lewat transition workflow (mis. 'Escalate' di template ITSM) atau mengubah field Team/Group.
  4. Agent dapat mengubah Priority serta Impact dan Urgency (ITSM); perubahan memicu evaluasi ulang goal SLA (JSM-QSK-024 butir 7).
  5. Saat request masuk status Pending (template ITSM), screen meminta Pending reason yang wajib diisi.
  6. Transition dapat menampilkan screen dengan kolom komentar ke customer (reply) atau internal note.
  7. Auto-assignment/round-robin dilakukan lewat automation (modul 20/35); routing berbasis jadwal/kapasitas dilakukan lewat Workforce Management (JSM-QSK-021).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Workflow dan screen transition di modul 09.

### JSM-QSK-017 Agent Productivity - Change Request Type

- **Deskripsi:** Mengubah atau menetapkan request type pada work item, termasuk work item yang dibuat di luar portal tanpa request type.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Edit issues); Admin Proyek untuk bulk edit.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi 'Change request type' tersedia di panel request details dan menampilkan request type yang cocok dengan work type saat ini.
  2. ⚠ Bila workflow request type baru berbeda, dialog meminta pemetaan status lama ke status baru sebelum disimpan.
  3. Menetapkan request type pada work item tanpa request type membuat work item tampil di portal customer.
  4. Perubahan request type dapat mengubah goal SLA yang berlaku (evaluasi ulang JQL goal) dan field yang tampil di layout.
  5. ⚠ Request type dapat diubah massal lewat bulk edit dengan aturan yang sama seperti perubahan tunggal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Request type dan pemetaan work type di modul 29.

### JSM-QSK-018 Agent Productivity - Link, Move, Clone & Raise on Behalf

- **Deskripsi:** Operasi untuk menghubungkan request ke pekerjaan lain, memindahkan, menyalin, atau membuat request atas nama customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (Link issues, Move issues, Create issues di project tujuan).
- **Aturan perilaku (acceptance criteria):**
  1. Agent dapat menautkan request ke work item project software/lain (mis. 'is caused by' bug) dan status item tertaut terlihat di issue view (ISS-029).
  2. Aksi 'Create linked work item' membuat work item di project developer langsung dari request dengan link otomatis.
  3. Move request ke service project lain memakai wizard inti (ISS-032) dengan langkah tambahan pemetaan request type dan status.
  4. Clone request menyalin field dan request type sesuai aturan ISS-031; komentar tidak ikut disalin.
  5. Agent dapat raise request on behalf of customer lewat Create dialog (memilih "Raise this request on behalf of") atau dari portal.
  6. ⚠ Merge request tidak disediakan; pola yang didukung adalah link 'duplicates' lalu menutup request duplikat dengan resolution Duplicate.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Raise on behalf dan agent-created requests dirinci di modul 29.

### JSM-QSK-019 Canned Responses

- **Deskripsi:** Template balasan siap pakai yang disisipkan agent dari editor komentar, tersedia sebagai personal atau shared, dan mendukung variabel dinamis. Heading ini menggabungkan fitur ringkas "Canned responses" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (membuat personal, memakai semua); Admin Proyek (membuat/mengelola shared).
- **Aturan perilaku (acceptance criteria):**
  1. Editor komentar menyediakan tombol canned responses pada mode Reply to customer maupun internal note.
  2. Create canned response memuat judul (wajib, unik per pemilik/scope) dan konten rich text dengan formatting dan link.
  3. Availability: Personal (hanya pembuat) atau Shared (semua agent di project).
  4. Hanya admin proyek yang dapat membuat shared atau mengubah availability; agent hanya dapat membuat personal (endpoint menolak dengan HTTP 403).
  5. Batas 100 personal canned response per agent per project dan 250 shared per project; melewati batas menampilkan galat batas.
  6. Variabel yang didukung: current user, assignee, work item key, reporter, resolution, summary, portal URL, request status, request URL; variabel terisi otomatis saat disisipkan.
  7. Saat menyisipkan, pengguna dapat mencari canned response berdasarkan judul (pencocokan sebagian, case-insensitive).
  8. Edit dan hapus tersedia; menghapus shared response menghapusnya untuk semua agent setelah konfirmasi.
  9. Konten yang disisipkan menjadi teks biasa di editor dan bisa diedit sebelum dikirim.
  10. Personal signature memakai mekanisme canned response; variabel tidak didukung di signature (ditampilkan apa adanya).
  11. Membuat, mengedit, dan menghapus canned response dikelola per project; canned response tidak dibagikan lintas project.
  12. ⚠ Sintaks variabel di konten memakai placeholder bertanda kurung kurawal ganda (mis. `{{reporter.name}}`, `{{issue.key}}`, `{{portal.url}}`) yang dipilih dari menu "Insert variable"; placeholder tak dikenal dibiarkan apa adanya.
  13. ⚠ Jumlah canned response per project dibatasi sesuai butir 5 dan batas ini sama untuk semua plan.
  14. Canned response berperilaku sama di semua service project; karena clone hanya membangun service project company-managed (JSM-PRJ-012), tidak ada varian perilaku team-managed, dan canned response project team-managed hasil impor dipertahankan setelah dipetakan ke company-managed.
  15. Formatting rich text dan tautan pada konten canned response tetap utuh saat disisipkan ke reply maupun internal note.
  16. Daftar insert memperbarui hasil pencarian nama per karakter yang diketik dan mendukung navigasi keyboard (panah dan Enter).
  17. Canned response yang disisipkan ke internal note tidak mengirim notifikasi ke customer walaupun berisi variabel portal URL atau request URL.
  18. ⚠ Canned response yang dibuat admin default berstatus shared dan yang dibuat agent default personal; agent tidak dapat melihat canned response personal agent lain.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Nama fitur ringkas lampiran ("Canned responses") identik setelah normalisasi sehingga tidak perlu baris alias; penempatan sub-fiturnya: kelola per project → butir 11; rich text & link → butir 15; variabel dinamis → butir 12; insert dengan pencarian nama → butir 16; dipakai di reply dan internal note → butir 17; kepemilikan shared/personal → butir 18; batas per project → butir 13; team-managed & company-managed → butir 14. Lihat JSM-AI-035 Canned Responses ([35](35-jsm-ai-otomasi-api-admin-csm.md)) untuk API/admin serta aturan variabel lanjutan dan saran AI (hingga 3 saran berlabel "Disarankan"); modul ini pemilik perilaku agent. Kedua fitur memakai satu entitas `CannedResponse` bersama; aturan visibility (Personal/Shared, hanya admin service project yang dapat membuat atau mengubah ke Shared), batas jumlah (butir 5), scope per project (butir 11), dan sintaks variabel template (butir 12) di sini identik dengan JSM-AI-035 dan menjadi acuan bila ada perbedaan. App canned response di Marketplace tidak dibangun (lihat Di luar cakupan).

### JSM-QSK-020 Agent Productivity - AI Assist for Agents

- **Deskripsi:** Fitur AI opsional di issue view agent untuk mempercepat penanganan request (ringkasan, perbaikan balasan, saran balasan, request serupa).
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan; Admin Situs (mengaktifkan/menonaktifkan AI organisasi).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Summarize" meringkas percakapan/komentar panjang menjadi poin-poin yang ditampilkan di panel samping; ringkasan tidak disimpan sebagai komentar kecuali agent memilih "Insert".
  2. Aksi "Improve reply" memperbaiki nada, panjang, dan ejaan draf reply; hasil ditampilkan sebagai saran yang harus diterima agent.
  3. ⚠ Saran balasan berbasis artikel KB yang relevan ditampilkan dengan tautan artikel sumber.
  4. ⚠ Panel "Similar requests" menampilkan request serupa yang bisa dilihat agent sebagai referensi.
  5. Admin organisasi dapat menonaktifkan seluruh fitur AI; saat nonaktif semua tombol AI disembunyikan dan endpoint menolak permintaan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Provider AI bersifat pluggable dan opsional (modul 27, 35); tumpang tindih dengan JSM-AI-038 AI Writing Assistance for Agents dan JSM-AI-039 AI Agent Assist & Natural Language Search di modul 35, yang menjadi sumber kebenaran perilaku AI, sedangkan fitur ini mendefinisikan titik tampilnya di issue view agent. Data yang dikirim ke provider hanya konten yang bisa dilihat agent. Layanan AI vendor tertentu di luar cakupan.

### JSM-QSK-021 Workforce Management

- **Deskripsi:** Kapabilitas untuk mengatur jadwal agent, merutekan work item ke agent yang tepat berdasarkan ketersediaan/kapasitas, dan memperkirakan kebutuhan staf.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Workforce Management admin (role terpisah), Manajer Layanan/ITSM (team lead), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Workforce Management diaktifkan per service project dengan memilih tim/agent yang dikelola.
  2. Admin WFM mengatur jadwal/shift agent (hari, jam mulai-selesai, time zone) dan status ketersediaan.
  3. Jadwal dapat diimpor massal via CSV; validasi format per baris dengan laporan galat (nomor baris, kolom, pesan) dan baris valid tetap diimpor.
  4. Work item dirutekan ke agent berdasarkan jadwal, ketersediaan, dan beban kerja (jumlah work item terbuka yang di-assign).
  5. Agent di luar shift atau berstatus tidak tersedia tidak menerima routing.
  6. Work item yang dirutekan tampil ter-assign ke agent di queue dan history mencatat "Assigned by Workforce Management".
  7. Capacity view menampilkan kapasitas tim per periode/shift dibanding volume work item.
  8. Staffing forecast memperkirakan kebutuhan staf dari volume request historis per jam/hari.
  9. Admin proyek dapat memberi role "Workforce Management admin" ke anggota tim tanpa memberikan Administer projects.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bersinggungan dengan jadwal on-call di modul 34 (jadwal shift harus memakai model kalender yang sama bila memungkinkan). Algoritma routing dan forecast menjadi pertanyaan terbuka.

## SLA: definisi metric, kondisi, goal & kalender

### JSM-QSK-022 SLAs - SLA Metrics Management

- **Alias inventaris:** SLAs (Service Level Agreements)
- **Deskripsi:** Halaman pengaturan SLA di project untuk membuat, mengubah, dan menghapus metric SLA beserta kondisi dan goal-nya. Heading ini juga menjadi rumah fitur ringkas "SLAs (Service Level Agreements)" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer projects); Manajer Layanan/ITSM bila diberi peran admin proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > SLAs (di navigasi baru: Request management > SLAs) menampilkan daftar metric SLA berisi nama dan ringkasan goal (jumlah goal, target terpendek-terpanjang).
  2. Create SLA dapat dimulai dari nol atau dari contoh template: Time to first response, Time to resolution, Time to close after resolution, Time to approve normal change, dan contoh lain sesuai template project.
  3. Nama SLA menjadi custom field bertipe SLA; metric dengan nama sama di beberapa project memakai satu field yang sama sehingga JQL lintas project konsisten; nama unik per project.
  4. Edit SLA (nama, kondisi, goal) disimpan sebagai versi baru konfigurasi dan memicu perhitungan ulang siklus yang masih berjalan (JSM-QSK-029).
  5. Delete SLA meminta konfirmasi yang menyebut jumlah work item terdampak; setelah dihapus, data SLA pada work item tidak lagi ditampilkan.
  6. Validasi simpan: wajib ada minimal satu kondisi Start dan satu kondisi Stop; galat "SLA memerlukan kondisi Start dan Stop".
  7. SLA tersedia di semua plan termasuk Free (tidak dikunci feature flag plan).
  8. Secara default SLA tidak ditampilkan ke customer di portal maupun email; tampilan ke customer hanya lewat Support timelines (JSM-QSK-026 butir 9).
  9. ⚠ Satu project dibatasi maksimal 30 metric SLA; tombol Create SLA nonaktif dengan pesan batas setelah batas tercapai.
  10. ⚠ Template ITSM menyertakan metric bawaan Time to first response, Time to resolution, Time to close after resolution, Time to approve normal change, dan Time to review normal change; template lain minimal Time to first response dan Time to resolution.
  11. Admin dapat membuat metric SLA kustom dengan nama bebas di luar daftar template.
  12. ⚠ Jumlah goal per metric dibatasi 90 (JSM-QSK-024 butir 6); batas metric dan goal ditampilkan di halaman pengaturan.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Penempatan sub-fitur fitur ringkas "SLAs (Service Level Agreements)": metric bawaan → butir 10; metric kustom → butir 11; start/pause/stop → JSM-QSK-023 butir 9, 11, 10; goals berurutan → JSM-QSK-024 butir 11; 'All remaining work items' dan 'No target' → JSM-QSK-024 butir 12; multi-cycle dan riwayat siklus → JSM-QSK-026 butir 13 (perilaku dasar di JSM-QSK-023 butir 5); timer di issue view/queue dengan warna status → JSM-QSK-026 butir 12; recalculation → JSM-QSK-029 butir 9; SLA history per request → JSM-QSK-026 butir 11; batas metric/goal → butir 9 dan 12; team-managed → JSM-QSK-030 butir 5; error log konfigurasi → JSM-QSK-029 butir 8. Model data: `sla_metrics` (project_id, field_id, name, version), `sla_conditions`, `sla_goals`, `sla_cycles` (work_item_id, metric_id, started_at, paused_ms, stopped_at, goal_id, breached, remaining_ms).

### JSM-QSK-023 SLAs - Time Metric Conditions (Start / Pause / Stop)

- **Deskripsi:** Aturan kapan waktu SLA mulai dihitung, dijeda, dan berhenti berdasarkan event pada request.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. "Start counting time when" menerima satu atau lebih kondisi berlogika OR, mis. Work item created, Entered status, Resolution cleared, Assignee from unassigned, Comment by customer.
  2. "Pause counting time when" bersifat opsional dan menerima kondisi berbasis status (mis. Waiting for customer, Pending) atau Resolution set.
  3. "Finish counting time when" menerima satu atau lebih kondisi dan siklus berhenti begitu salah satu terpenuhi, mis. Entered status, Resolution set, Comment for customers, Assignee set, Approval completed.
  4. ⚠ Daftar kondisi dapat diperluas oleh modul lain (mis. Assets, approvals) melalui registry kondisi; kondisi dari modul nonaktif ditandai "tidak tersedia" dan tidak dievaluasi.
  5. Multiple cycles: bila kondisi Start terpenuhi lagi setelah Stop (mis. request dibuka ulang), siklus SLA baru dimulai dan siklus lama tersimpan sebagai completed cycle.
  6. Waktu selama pause tidak dihitung; SLA berlanjut otomatis saat kondisi pause tidak berlaku lagi.
  7. Bila kondisi Stop terpenuhi sebelum Start, tidak ada siklus yang terbentuk dan SLA tidak tampil pada work item tersebut; editor menjelaskan perilaku ini di teks bantuan.
  8. ⚠ Editor menampilkan peringatan saat kombinasi kondisi berisiko membuat SLA tidak pernah berhenti (mis. Stop hanya "Entered status X" padahal X tidak ada di workflow project).
  9. Kondisi start yang tersedia minimal mencakup Issue created, Entered status, Resolution cleared, Comment: by customer, Comment: for customers, Assignee changed, dan Assignee from unassigned (dari jatah SLA ringkas).
  10. Kondisi stop yang tersedia minimal mencakup Entered status, Resolution set, Comment: for customers, dan Assignee set (dari jatah SLA ringkas).
  11. Kondisi pause berbasis status dapat memilih status workflow project mana pun (mis. Waiting for customer); masuk ke status pause menjeda timer seketika dan keluar dari status tersebut melanjutkan timer pada detik yang sama, tercatat di SLA history (dari jatah SLA ringkas).
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Mesin SLA mengonsumsi event domain `issue.created`, `issue.transitioned`, `issue.updated(assignee|resolution)`, `comment.created(visibility)` dari modul 28; evaluasi harus idempoten dan berurutan per work item.

### JSM-QSK-024 SLAs - Goals (JQL-based Targets)

- **Deskripsi:** Goal SLA menetapkan target waktu berbeda untuk subset request yang dipilih dengan JQL, masing-masing dengan kalendernya sendiri.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Tambah goal memuat Work items (JQL), Goal (target waktu, mis. `4h`, `30m`, `2h 30m`), dan Calendar.
  2. Goal 'All remaining work items' selalu ada di posisi paling bawah sebagai fallback dan tidak bisa dipindah atau dihapus.
  3. Opsi 'No target' mencatat waktu tanpa pernah breach.
  4. Goal dievaluasi dari atas ke bawah dan goal pertama yang cocok dipakai; urutan diatur dengan drag.
  5. Beberapa priority dapat dikelompokkan dalam satu goal ("Use priority to group SLA goals"); grup tetap dihitung sebagai satu goal.
  6. Satu SLA dibatasi maksimal 90 goal; galat batas tampil saat menambah goal ke-91.
  7. Goal dievaluasi ulang saat field yang dirujuk berubah (mis. priority naik) sehingga target siklus berjalan bisa berganti; elapsed time tetap, target dan remaining dihitung ulang.
  8. JQL goal divalidasi saat simpan; galat sintaks atau field tidak dikenal ditampilkan per baris goal.
  9. Goal dapat memakai kriteria request type, priority, organization, atau field Assets di JQL.
  10. ⚠ Target waktu menerima satuan `w`, `d`, `h`, `m`; `1d` dan `1w` dihitung sebagai 24 jam dan 7 hari jam kalender, bukan hari kerja.
  11. Setiap goal membawa target waktu dan kalender sendiri; untuk satu siklus hanya goal pertama yang cocok yang dipakai dan goal lain diabaikan.
  12. Goal 'All remaining work items' dibuat otomatis untuk setiap SLA baru dan dapat diset ke target tertentu atau 'No target'.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** JQL goal dieksekusi terhadap satu work item (mode matcher), sehingga engine JQL modul 12 harus menyediakan evaluasi in-memory per item selain kompilasi SQL.

### JSM-QSK-025 SLAs - Calendars (Working Hours & Holidays)

- **Alias inventaris:** SLA calendars
- **Deskripsi:** Kalender kerja agar SLA hanya dihitung di jam kerja, lengkap dengan time zone dan hari libur. Heading ini juga mencakup fitur ringkas "SLA calendars" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap project memiliki kalender default '24/7 Calendar' (tidak bisa dihapus) dan contoh 'Sample 9-5 Calendar'.
  2. Create calendar memuat nama (wajib, unik per project), deskripsi, dan time zone (IANA).
  3. Working hours diatur per hari dan boleh berisi beberapa interval per hari (mis. 09:00-12:00 dan 13:00-17:00); interval tumpang tindih ditolak dengan galat.
  4. Holidays memuat tanggal, nama, dan opsi berulang tiap tahun.
  5. ⚠ Kalender yang masih dipakai goal tidak bisa dihapus; dialog menampilkan daftar SLA/goal pemakainya dan meminta admin menggantinya lebih dulu.
  6. Satu kalender dapat dipakai banyak goal dan metric di project yang sama.
  7. Perubahan kalender memicu perhitungan ulang siklus yang masih berjalan (JSM-QSK-029).
  8. Fungsi JQL `withinCalendarHours()` memakai kalender goal yang berlaku (JSM-QSK-027).
  9. Daylight saving time ditangani lewat time zone kalender; jam kerja mengikuti waktu lokal kalender.
  10. Kalender ditetapkan per goal SLA (bukan per metric), sehingga goal berbeda dalam satu metric boleh memakai kalender berbeda.
  11. SLA tidak berjalan di luar jam kerja dan pada hari libur kalender; timer di UI menampilkan status "Di luar jam kerja" saat itu.
  12. Kalender bersifat per project; kalender tidak dibagikan lintas project di versi pertama.
  13. Time zone kalender wajib diisi dengan default time zone site; perubahan time zone berlaku untuk perhitungan berikutnya.
  14. Working hours boleh kosong untuk hari tertentu (hari tidak bekerja), tetapi kalender tanpa jam kerja sama sekali ditolak dengan galat.
  15. Holiday berulang tahunan berlaku pada tanggal yang sama setiap tahun; holiday non-berulang hanya berlaku pada tanggal dan tahun yang ditetapkan.
  16. Goal baru memakai kalender 24/7 secara otomatis bila admin tidak memilih kalender.
  17. `withinCalendarHours()` bernilai salah untuk work item yang saat ini berada pada holiday kalender goal-nya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penempatan sub-fitur fitur ringkas "SLA calendars": nama/deskripsi/time zone → butir 13; working hours multi-rentang → butir 14; holidays berulang → butir 15; kalender 24/7 default → butir 16; kalender per goal → butir 10; tidak berjalan di luar jam kerja → butir 11; kalender per project → butir 12; `withinCalendarHours()` → butir 17. Perhitungan durasi kerja wajib memakai pustaka time zone server (bukan offset tetap).

## SLA: tampilan, kueri, otomasi & operasi

### JSM-QSK-026 SLAs - Display on Issue View & Queues

- **Deskripsi:** Tampilan status SLA untuk agent di issue view, kolom queue, dan navigator, serta opsi menampilkan support timeline ke customer di portal.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (melihat); Customer (support timeline); Admin Proyek (mengatur support timeline).
- **Aturan perilaku (acceptance criteria):**
  1. Panel SLA di issue view menampilkan setiap metric: sisa waktu, target, dan status (running, paused, met, breached).
  2. Kode warna: normal (netral), mendekati breach (kuning/oranye, default sisa < 30 menit atau < 10% target), breached (merah, waktu negatif).
  3. SLA yang sedang dijeda menampilkan ikon pause.
  4. Riwayat siklus selesai (completed cycles) ditampilkan dengan status met/breached dan durasi.
  5. Tooltip per metric menampilkan goal yang berlaku, kalender, dan waktu mulai siklus.
  6. Kolom SLA di queue dan navigator memakai format ringkas (mis. `2h 15m`, `-30m`).
  7. Metric tanpa siklus aktif atau selesai tidak ditampilkan di panel dan tampil kosong di kolom.
  8. Secara default SLA disembunyikan dari customer di portal dan email.
  9. ⚠ Support timelines in the portal: admin memilih metric SLA yang ditampilkan ke customer sebagai target dukungan di halaman request portal; portal menampilkan tanggal/jam target dalam time zone customer dan, saat breach, teks netral "Kami sedang menindaklanjuti" tanpa waktu negatif.
  10. ⚠ Nilai SLA ikut di export CSV navigator sebagai sisa waktu dalam format ringkas dan kolom terpisah status siklus.
  11. Tab "SLA history" per request mencatat setiap siklus (mulai, jeda, selesai, goal, hasil) secara kronologis (dari jatah SLA ringkas).
  12. Timer SLA di kolom queue memakai warna status yang sama dengan panel issue view agar konsisten (dari jatah SLA ringkas).
  13. Request yang dibuka ulang menampilkan siklus baru sebagai timer aktif dan siklus sebelumnya di daftar completed cycles.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Timer di UI dihitung di klien dari snapshot server (started_at, paused_ms, target, kalender) agar tidak perlu polling per detik. Support timeline tampil di portal modul 30.

### JSM-QSK-027 SLAs - JQL Functions

- **Alias inventaris:** SLA JQL & breach handling
- **Deskripsi:** Fungsi JQL khusus SLA untuk mencari dan menyaring request berdasarkan status SLA, dipakai di queue, filter, goal, automation, dan report. Heading ini juga mencakup fitur ringkas "SLA JQL & breach handling" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek; hasil mengikuti Browse projects dan issue security.
- **Aturan perilaku (acceptance criteria):**
  1. `breached()` mengembalikan work item yang siklus terakhirnya breach.
  2. `everBreached()` mengembalikan work item yang pernah breach di siklus mana pun.
  3. `completed()` mengembalikan work item yang siklus terakhirnya sudah selesai.
  4. `running()` mengembalikan work item yang siklusnya sedang berjalan.
  5. `paused()` mengembalikan work item yang siklusnya sedang dijeda.
  6. `elapsed("duration")` membandingkan waktu yang sudah berjalan, mis. `"Time to resolution" > elapsed("4h")`.
  7. `remaining("duration")` membandingkan sisa waktu, mis. `"Time to first response" < remaining("1h")`; nilai breach bernilai negatif.
  8. `withinCalendarHours()` mengembalikan work item yang saat ini berada dalam jam kerja kalender goal-nya.
  9. Sintaks `"<SLA name>" = breached()` dan `!=` untuk negasi didukung; nama SLA tak dikenal menghasilkan galat "Field '<nama>' tidak ada".
  10. ORDER BY berdasarkan field SLA mengurutkan menurut sisa waktu siklus berjalan (JSM-QSK-004).
  11. Field SLA dapat dipakai lintas project karena satu custom field per nama SLA.
  12. Contoh kueri breach handling `"Time to resolution" < remaining("2h")` didokumentasikan di bantuan JQL dan di template queue 'SLA breaching soon'.
  13. Field SLA dan fungsi di atas dapat dipakai di filter tersimpan dan gadget dashboard (JSM-QSK-046).
  14. Kedelapan fungsi SLA (breached, completed, elapsed, everBreached, paused, remaining, running, withinCalendarHours) muncul di autocomplete JQL setelah operator pada field SLA.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fungsi didaftarkan ke engine JQL modul 12 (pola registry fungsi); daftar field/fungsi JSM lengkap ada di modul 35 (JSM-AI-013 JSM JQL Fields, JSM-AI-014 JSM JQL SLA Functions); semantik fungsi SLA didefinisikan di sini dan JSM-AI-014 merujuk ke fitur ini. Penempatan sub-fitur fitur ringkas "SLA JQL & breach handling": fungsi JQL → butir 14; contoh → butir 12; trigger 'SLA threshold breached' → JSM-QSK-028 butir 12; notifikasi email/Slack/Teams → JSM-QSK-028 butir 10; eskalasi priority/assignee → JSM-QSK-028 butir 11; queue 'Breached SLAs' → JSM-QSK-001 butir 9; field SLA di filter/gadget → butir 13.

### JSM-QSK-028 SLAs - Automation Triggers & Time-based Alerts

- **Deskripsi:** Trigger, condition, dan smart value automation khas SLA untuk eskalasi dan peringatan berbasis waktu; ini juga cara native untuk alert berbasis waktu di queue karena tidak ada setting alert khusus per queue.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengelola rule project); Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Trigger 'SLA threshold breached' memilih metric SLA dan ambang 'breached', 'will breach in X', atau 'breached X ago'; trigger dievaluasi oleh job latar minimal setiap menit.
  2. Satu rule dapat memantau beberapa metric SLA sekaligus.
  3. Smart values SLA tersedia, mis. `{{issue."Time to resolution".ongoingCycle.remainingTime.friendly}}`, `breached`, `goalDuration`, `startTime`.
  4. Condition JQL di dalam rule dapat memakai fungsi SLA (JSM-QSK-027).
  5. Action eskalasi yang relevan: notifikasi email/chat, naikkan priority, re-assign, transition.
  6. Tersedia template rule bawaan "SLA hampir breach → beri tahu assignee dan naikkan priority".
  7. Pola alert queue terdokumentasi: queue 'SLA breaching soon' (JQL `remaining("30m")` atau `breached()`) ditambah warna kolom SLA.
  8. Pola alert berbasis umur request terdokumentasi: scheduled rule dengan JQL, mis. `created <= -2d AND status = "Waiting for support"`.
  9. ⚠ Trigger tidak menembak selama SLA dijeda; bila goal berubah, ambang dihitung ulang dan setiap ambang hanya menembak sekali per siklus.
  10. Notifikasi ke agent atau lead lewat email, Slack, atau Teams dikirim melalui action automation; action Slack/Teams aktif hanya bila integrasi chat (modul 25/30) terpasang.
  11. Eskalasi otomatis dengan mengubah priority dan/atau assignee tercatat di history dengan aktor "Automation".
  12. Ambang 'will breach in X' menerima X minimal 1 menit dan maksimal 30 hari; nilai di luar rentang ditolak saat rule disimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine automation di modul 20; trigger/action JSM di modul 35 (JSM-AI-016 JSM Triggers, JSM-AI-020 JSM Smart Values). Alert per queue native tidak dibangun.

### JSM-QSK-029 SLAs - Recalculation, Errors & Troubleshooting

- **Deskripsi:** Perilaku sistem saat konfigurasi SLA berubah, alat perhitungan ulang, dan penanganan SLA yang salah atau hilang.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Mengubah SLA atau kalender menghitung ulang siklus yang masih berjalan memakai konfigurasi baru lewat job latar dengan indikator progres di halaman SLA.
  2. Siklus yang sudah selesai tidak dihitung ulang; perubahan hanya berlaku untuk siklus yang selesai setelah konfigurasi diubah.
  3. ⚠ Menu panel SLA menyediakan aksi "Recalculate SLAs" per work item yang menyusun ulang siklus dari history work item.
  4. ⚠ Perhitungan ulang massal tersedia lewat REST API (batch JQL, maksimal 1000 work item per job) dan action automation.
  5. Bila SLA gagal dihitung (mis. JQL goal tidak valid atau field terhapus), panel SLA menampilkan pesan galat dan halaman SLA menandai metric bermasalah.
  6. ⚠ Saat work item dipindah antarproject, siklus SLA project asal dihentikan dan SLA project tujuan dihitung dari history; data impor tanpa history memulai siklus dari waktu impor.
  7. ⚠ Setiap perubahan konfigurasi SLA dan kalender tercatat di audit log (aktor, waktu, perbedaan).
  8. Halaman "SLA error log" per project mencantumkan peringatan konfigurasi tidak valid (mis. JQL error) dengan waktu kejadian dan work item terdampak (dari jatah SLA ringkas).
  9. ⚠ Cakupan perhitungan ulang saat konfigurasi diubah ditampilkan di dialog konfirmasi simpan (jumlah siklus berjalan yang akan dihitung ulang) (dari jatah SLA ringkas, butir recalculation).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mode bulk recalculation non-destructive/graceful/forceful khas Data Center tidak dibangun (lihat Di luar cakupan). Audit log mengikuti modul 22.

### JSM-QSK-030 SLAs - Team-managed Service Project SLAs

- **Deskripsi:** Konfigurasi SLA versi sederhana untuk team-managed service project di produk acuan. Dicatat sebagai rujukan saja karena clone hanya membangun service project company-managed (JSM-PRJ-012); seluruh kebutuhan SLA dipenuhi JSM-QSK-022 s.d. JSM-QSK-029.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Di produk acuan: Admin Proyek team-managed (role Administrator). Di clone tidak ada aktor karena varian ini tidak dibangun.
- **Aturan perilaku (acceptance criteria):**
  1. Pengaturan SLA di project team-managed tidak dibangun; semua service project mengatur SLA lewat halaman SLA company-managed yang memakai mesin SLA tunggal (JSM-QSK-022).
  2. ⚠ Subset kondisi (status dan event dasar: created, resolution set/cleared, comment for customers) dan goal berbasis pemilih request type/priority tanpa JQL bebas tidak dibangun sebagai mode terpisah; kebutuhan yang sama dipenuhi editor kondisi JSM-QSK-023 dan editor goal JSM-QSK-024 (Basic dan JQL).
  3. Pemilihan kalender dan target waktu per goal sudah tercakup di JSM-QSK-024 dan JSM-QSK-025.
  4. ⚠ SLA project team-managed hasil impor (modul 25) dipetakan ke metric SLA company-managed sehingga tetap tampil di issue view, kolom queue, dan JQL seperti SLA lain.
  5. Tidak ada UI pengaturan SLA team-managed terpisah; struktur daftar metric, editor kondisi, dan editor goal hanya ada satu.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Heading dipertahankan agar inventaris tercakup. Keputusan dukungan service project team-managed dicatat di modul 29 (JSM-PRJ-012) dan README (pertanyaan terbuka "Service project team-managed"). Bila pemilik produk kelak membangun varian ini, gunakan satu mesin SLA dengan "mode UI" berbeda agar tidak ada dua jalur perhitungan. Lihat Di luar cakupan.

## Knowledge base: sumber, akses & pengelolaan artikel

### JSM-QSK-031 Knowledge Base - Link Confluence Space

- **Alias inventaris:** Knowledge base
- **Deskripsi:** Menghubungkan service project ke satu atau beberapa knowledge space (padanan Confluence space di produk ini), termasuk dari site lain dalam organisasi, sebagai sumber artikel untuk portal dan agent. Heading ini juga mencakup fitur ringkas "Knowledge base" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (Administer projects) menautkan; admin knowledge space (padanan Confluence admin) mengatur permission space.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Knowledge base memungkinkan admin menautkan knowledge space yang sudah ada atau membuat space baru langsung dari service project.
  2. Satu project dapat menautkan beberapa space; urutan space menentukan urutan hasil di panel agent bila skor relevansi sama.
  3. Unlink space hanya memutus tautan dan tidak menghapus konten space; dialog konfirmasi menjelaskan hal ini.
  4. Space KB baru dibuat dengan blueprint Knowledge base berisi template artikel How-to dan Troubleshooting.
  5. Bila modul knowledge space belum aktif di site, halaman menampilkan peringatan beserta tombol aktivasi; versi gratis knowledge space cukup untuk dipakai sebagai KB.
  6. ⚠ KB tersedia tanpa lisensi produk terpisah; aktivasi hanya dikontrol feature flag plan site (Fase 5: semua plan berbayar dan Free dengan batas jumlah artikel).
  7. Agent memerlukan permission tulis di space untuk membuat/mengedit artikel; customer hanya membaca.
  8. ⚠ Admin dapat memetakan space tertentu ke request type/portal sehingga suggestion portal untuk request type itu memprioritaskan artikel dari space terpetakan.
  9. Admin dapat memilih site lain dalam organisasi yang sama sebagai sumber KB lewat aksi "Add knowledge from a different site".
  10. Untuk space dari site lain, sistem meminta otorisasi lintas site dan menghormati permission site sumber; artikel site lain tetap muncul di pencarian/suggestion portal dan panel agent.
  11. ⚠ Space dari site lain tidak mendukung create article from request dan akses customer tanpa lisensi; UI menonaktifkan aksi tersebut dengan penjelasan.
  12. Artikel dari space yang ditautkan tampil di help center, portal, dan widget (modul 30) (dari jatah KB ringkas).
  13. Daftar space tertaut di Project settings > Knowledge base menampilkan nama space, jumlah artikel, dan status akses customer.
  14. Space baru yang dibuat dari pengaturan KB otomatis ditautkan ke project dan admin proyek pembuatnya menjadi admin space.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk ini tidak memiliki produk wiki terpisah; "Confluence space" diimplementasikan sebagai entitas native `kb_spaces` + `kb_articles` (konten ADF, versi, status draft/published, restriksi halaman). Keputusan: knowledge space native ini menjadi document store bawaan produk ini sejak Fase 5. Project Pages (PRJ-030, [04](04-proyek.md)), Linked Confluence Pages dan Create Confluence Page from Issue (CMT-030, CMT-031, [08](08-komentar-lampiran-aktivitas.md)), Sprint Linked Pages (SPR-026, [11](11-backlog-sprint.md)), Project pages/release notes (INT-046, [25](25-integrasi-impor-ekspor-ekstensi.md)), dan tab Pages (BIZ-026, [26](26-kerja-bisnis-formulir-persetujuan.md)) dapat menautkan atau membuat halaman di knowledge space native ini sejak Fase 5. Konektor ke wiki eksternal (INT-042, INT-043, modul 25) tetap opsional, dan pertanyaan terbukanya dipusatkan di modul 25. Lifecycle, versi, dan administrasi space dirumuskan di JSM-QSK-048. Penempatan sub-fitur fitur ringkas "Knowledge base": menautkan satu atau lebih space → butir 2; membuat space baru dari pengaturan KB → butir 14; KB native tanpa produk wiki terpisah → JSM-QSK-033 butir 8; opsi akses portal vs pengguna berlisensi → JSM-QSK-032 butir 1 dan 9; template How-to/Troubleshooting → butir 4 dan JSM-QSK-033 butir 9; create article from request dengan prefill → JSM-QSK-037 butir 7; cari & bagikan dari issue view → JSM-QSK-036 butir 7; kategori per portal → JSM-QSK-034 butir 9; artikel di help center/portal/widget → butir 12; feedback 'Was this helpful?' → JSM-QSK-038 butir 7; sumber eksternal untuk AI answers → JSM-QSK-039 butir 6; restriksi halaman → JSM-QSK-032 butir 10.

### JSM-QSK-032 Knowledge Base - Visibility, Permissions & Internal vs External KB

- **Deskripsi:** Mengatur siapa yang dapat melihat artikel (customer, pengguna anonim, atau hanya pengguna berlisensi) dan membedakan KB internal untuk agent dari KB eksternal untuk customer.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, admin knowledge space, Customer, Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi akses KB per project: "Semua yang bisa raise request di project ini" (customer tanpa lisensi) atau "Hanya pengguna berlisensi".
  2. Bila portal publik ('anyone can access'), artikel dapat dibaca anonim sesuai setting akses; bila tidak, pengguna anonim diarahkan ke login portal.
  3. Akses baca customer tanpa lisensi dikonfigurasi di permission knowledge space (padanan space permissions) dan disinkronkan dengan opsi butir 1.
  4. Restriksi halaman di knowledge space tetap berlaku di portal, pencarian, suggestion, dan panel agent (diperiksa di server).
  5. Space/artikel ditandai internal (agent-facing, mis. runbook/SOP) atau eksternal (customer-facing); konten internal tidak pernah tampil di portal, widget, maupun API customer.
  6. Customer membaca artikel dalam tampilan portal (tema portal), bukan UI pengelolaan knowledge space.
  7. Draft dan halaman yang belum dipublikasikan tidak tampil ke customer maupun di suggestion portal.
  8. ⚠ Agent tanpa permission tulis di space dapat melihat dan membagikan artikel tetapi tombol create/edit artikel disembunyikan.
  9. Perubahan opsi akses KB berlaku seketika untuk pencarian dan suggestion portal (cache hasil dibersihkan).
  10. Artikel dengan restriksi halaman yang tidak dapat diakses customer tidak muncul di daftar kategori maupun hitungan jumlah artikel portal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Akses portal dan customer di modul 30; model permission knowledge space mengikuti pola permission scheme modul 03.

### JSM-QSK-033 Knowledge Base - Knowledge Base Section & Articles View

- **Deskripsi:** Halaman Knowledge base di sidebar service project tempat agent dan admin menjelajah, membuat, mengedit, dan mengkategorikan artikel langsung di produk ini.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (dengan permission tulis space), Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Sidebar service project memiliki menu Knowledge base bagi agent dan admin.
  2. ⚠ Articles view menampilkan daftar semua artikel dari space terhubung dengan pencarian teks dan filter space, kategori, status (draft/published), visibilitas (internal/eksternal), dan pemilik.
  3. Native article editor menyediakan create, read, edit, dan delete artikel langsung di service project tanpa berpindah aplikasi.
  4. ⚠ Saat membuat artikel, pengguna memilih template How-to, Troubleshooting, FAQ, atau blank.
  5. Artikel dapat diedit dari Articles view, termasuk menambahkannya ke beberapa kategori.
  6. Halaman ini menyediakan tombol "Create article".
  7. ⚠ Editor artikel menyediakan toggle "Internal / Eksternal" yang menentukan visibilitas artikel di portal (JSM-QSK-032 butir 5).
  8. ⚠ KB berjalan sepenuhnya di produk ini tanpa produk wiki terpisah (knowledge space native), sehingga seluruh fungsi KB tersedia sejak Fase 5 (dari jatah KB ringkas).
  9. Template How-to dan Troubleshooting memiliki struktur bagian bawaan (mis. Langkah; Gejala, Penyebab, Solusi) yang dapat diubah admin space.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Editor memakai editor rich text modul 07; draft/publish, riwayat versi, hierarki, label, restriksi, trash, dan administrasi space dirumuskan sebagai aturan di JSM-QSK-048. Tampilan artikel tertaut di issue view agent didefinisikan di JSM-PRJ-059 Linked knowledge base articles in issue view ([29](29-jsm-service-project-request-type-formulir.md)) dan saran artikel di portal di JSM-PRT-025 KB article suggestions on portal (deflection) ([30](30-jsm-help-center-portal-pelanggan-kanal.md)); modul ini pemilik model artikel KB dan section Knowledge base.

### JSM-QSK-034 Knowledge Base - Categories & Featured Articles on Portal

- **Deskripsi:** Mengelompokkan artikel KB ke kategori yang tampil di portal/help center agar customer mudah menjelajah.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengelola kategori); Customer (menjelajah).
- **Aturan perilaku (acceptance criteria):**
  1. Admin membuat kategori dari section Knowledge base dengan nama (wajib) dan deskripsi, lalu menambahkan artikel.
  2. Batas 30 kategori per service project dan 100 artikel per kategori; galat batas ditampilkan saat dilampaui.
  3. Satu artikel dapat masuk ke beberapa kategori.
  4. Jelajah kategori di portal hanya menampilkan artikel yang sudah dikategorikan; artikel lain tetap bisa ditemukan lewat pencarian dan suggested articles.
  5. Artikel dapat ditandai featured sehingga tampil paling atas di kategorinya.
  6. Admin mengurutkan kategori di portal dengan drag and drop.
  7. Di help center modern (modul 30), kategori KB dapat ditampilkan sebagai kartu lewat komponen halaman help center.
  8. ⚠ Kategori dapat diedit dan dihapus; menghapus kategori hanya melepas artikel dari kategori, tidak menghapus artikel.
  9. Kategori berlaku per portal (per service project), menggantikan mekanisme kategori berbasis label (dari jatah KB ringkas).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Layout help center dan portal di modul 30.

## Knowledge base: self-service, deflection & panel agent

### JSM-QSK-035 Knowledge Base - Portal Article Suggestions & Deflection

- **Alias inventaris:** Request deflection & self-service
- **Deskripsi:** Saran artikel di help center dan portal agar customer menyelesaikan masalah sendiri sebelum membuat request, serta pencatatan deflection. Heading ini juga mencakup fitur ringkas "Request deflection & self-service" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor; Admin Proyek (mengatur KB).
- **Aturan perilaku (acceptance criteria):**
  1. Pencarian help center/portal menampilkan artikel KB bersama request type yang cocok.
  2. Suggested articles muncul real-time (debounce ±300 ms, minimal 3 karakter) saat customer mengetik summary di form request.
  3. Customer dapat membuka artikel di portal lalu menjawab 'Was this helpful?'.
  4. Deflection tercatat saat customer memilih suggested article di portal lalu menilainya helpful.
  5. Artikel di portal menampilkan breadcrumb, judul, dan konten hasil render.
  6. Setelah membaca artikel, tautan untuk tetap membuat request selalu tersedia dan form tidak kehilangan isian.
  7. ⚠ Pencarian di help center mencakup artikel dari semua portal yang dapat diakses customer, dengan label portal asal.
  8. AI answers di help center memakai KB sebagai sumber (dirinci di JSM-AI-052, modul 35).
  9. Hasil pencarian portal menampilkan bagian artikel di atas daftar request type.
  10. Sinyal deflection kedua tercatat saat customer membuka artikel dari suggestion lalu meninggalkan form tanpa membuat request dalam sesi yang sama; laporan membedakan kedua sinyal (butir 4 dan butir ini).
  11. Di halaman artikel, customer dapat menekan "Masalah saya sudah teratasi" yang dicatat sebagai deflection terkonfirmasi.
  12. Data deflection memasok report Requests deflected (JSM-QSK-042 butir 3).
  13. Suggested articles dibatasi maksimal 5 artikel dan diperbarui setiap kali summary berubah; artikel yang sudah dibuka ditandai.
  14. Bila AI answers aktif, jawaban AI tampil di atas daftar artikel dan konfirmasi "jawaban membantu" dicatat sebagai deflection AI terpisah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Form request di modul 29, help center di modul 30, AI answers di modul 35. Sisi portal fitur ini juga tercantum sebagai JSM-PRT-025 KB article suggestions on portal (deflection) ([30](30-jsm-help-center-portal-pelanggan-kanal.md)) dan artikel tertaut di issue view sebagai JSM-PRJ-059 Linked knowledge base articles in issue view ([29](29-jsm-service-project-request-type-formulir.md)); modul ini pemilik aturan deflection dan pencatatannya, sedangkan JSM-PRT-025 mengatur titik tampil di portal. Penempatan sub-fitur fitur ringkas "Request deflection & self-service": saran real-time saat mengetik summary → butir 13; artikel di atas request type → butir 9; pencatatan deflected tanpa membuat request → butir 10; customer menandai terselesaikan → butir 11; report Requests deflected → butir 12; AI answers sebagai deflection lanjutan → butir 14.

### JSM-QSK-036 Knowledge Base - Agent Article Panel (Search & Share)

- **Deskripsi:** Panel KB di issue view agent untuk melihat saran artikel, mencari, dan membagikan artikel ke customer.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Panel Knowledge base di issue view menampilkan saran otomatis berdasarkan summary/description request.
  2. Agent dapat mencari artikel secara manual dari panel.
  3. Agent dapat melihat preview artikel di panel tanpa meninggalkan request.
  4. Aksi "Share with customer" menyisipkan tautan artikel ke draf reply customer.
  5. Artikel yang dibagikan tercatat pada request dan dipakai report Requests resolved (resolved with article).
  6. Tombol "Open in knowledge base" membuka artikel di editor artikel untuk diedit (bila agent punya izin tulis).
  7. Pencarian di panel agent mencakup artikel internal dan eksternal yang bisa dilihat agent, masing-masing dengan label visibilitas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Artikel internal hanya boleh dibagikan sebagai referensi internal; aksi share disembunyikan untuk artikel internal.

### JSM-QSK-037 Knowledge Base - Create Article from Request

- **Deskripsi:** Agent membuat artikel KB langsung dari request untuk mendokumentasikan solusi.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan dengan permission Add page di space tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Panel KB di issue view menyediakan tombol 'Create article'.
  2. Bila ada beberapa space terhubung, agent memilih space tujuan (default space pertama yang bisa ditulisi).
  3. Agent memilih template How-to, Troubleshooting, atau blank.
  4. ⚠ Judul diisi dari summary dan konten awal dari description request; agent wajib meninjau sebelum publish.
  5. Editor artikel terbuka dalam modal; agent dapat publish atau menyimpan sebagai draft.
  6. Tombol 'Create article' hanya aktif bila agent punya permission Add page di space; selain itu tampil tooltip alasan.
  7. Artikel yang dibuat ditautkan balik ke request sumber (tercatat di aktivitas) (dari jatah KB ringkas, butir prefill konten).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Prefill tidak boleh menyalin internal note atau data pribadi customer secara otomatis. Drafting artikel berbantuan AI dirinci di JSM-AI-040 (modul 35).

### JSM-QSK-038 Knowledge Base - Article Feedback

- **Deskripsi:** Penilaian customer terhadap artikel untuk mengukur kegunaannya dan menghitung deflection.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (memberi vote); Admin Proyek (melihat hasil).
- **Aturan perilaku (acceptance criteria):**
  1. Artikel portal menampilkan prompt 'Was this helpful?' dengan Yes/No.
  2. Vote helpful dipakai report Requests deflected.
  3. ⚠ Saat menjawab No, customer dapat menambahkan komentar opsional (maksimal 1000 karakter) yang hanya terlihat agent/admin.
  4. ⚠ Satu customer hanya memiliki satu vote per artikel; vote berikutnya menggantikan vote sebelumnya.
  5. Feedback tidak menjadi komentar publik pada artikel.
  6. Articles view menampilkan metrik Article usage (views) dan Article effectiveness (rasio helpful) per artikel (dari jatah Service project reports ringkas).
  7. Prompt feedback berperilaku sama di portal, help center, dan widget, dan vote dari ketiganya dihitung dalam metrik yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Vote anonim (portal publik) dicatat per sesi.

### JSM-QSK-039 Knowledge Base - External Knowledge Sources

- **Deskripsi:** Menghubungkan sumber pengetahuan di luar knowledge space produk ini agar bisa dipakai pencarian/AI answers di help center.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Admin menambahkan sumber eksternal dari pengaturan knowledge/AI.
  2. ⚠ Konektor awal yang dipertimbangkan: SharePoint dan folder Google Drive; konektor lain lewat antarmuka konektor generik.
  3. Konten sumber disinkronkan dan diindeks secara berkala dan bisa dipicu manual.
  4. Visibilitas hasil mengikuti permission di sumber; konten yang tidak dapat diakses customer tidak ditampilkan.
  5. Halaman sumber menampilkan status sinkronisasi (terakhir berhasil, jumlah dokumen) dan pesan galat.
  6. ⚠ Sumber eksternal dipakai untuk AI answers dan pencarian help center saja, tidak muncul di kategori portal (dari jatah KB ringkas).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sangat bergantung pada AI answers (modul 35) dan SSRF guard (modul 25). Konektor Rovo vendor tidak dibangun.

## Umpan balik pelanggan: CSAT & Surveys

### JSM-QSK-040 Customer Satisfaction (CSAT) Surveys

- **Alias inventaris:** Customer satisfaction (CSAT)
- **Deskripsi:** Survei kepuasan yang dikirim ke customer setelah request diselesaikan, beserta pengaturan, tampilan bagi agent, dan pelaporannya. Heading ini juga mencakup fitur ringkas "Customer satisfaction (CSAT)" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (mengatur); Customer/Pelapor (menilai); Agent Layanan (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Customer satisfaction menyediakan toggle aktif/nonaktif per project (default aktif untuk template baru).
  2. Saat request resolved, survei dikirim lewat email notifikasi customer dan juga tampil di halaman request portal.
  3. Rating berskala 1-5 (bintang) dengan komentar opsional (maksimal 2000 karakter).
  4. ⚠ Admin dapat mengubah teks pertanyaan dan pesan terima kasih; teks default dipulihkan bila dikosongkan.
  5. ⚠ Satu rating per request; customer dapat memperbarui rating selama request belum ditutup final, dan setiap perubahan memperbarui Satisfaction date.
  6. Survei hanya dikirim ke reporter, bukan participants.
  7. Field Satisfaction dan Satisfaction date dapat dipakai di JQL, kolom queue, dan report.
  8. Agent melihat rating dan komentar di panel request details (JSM-QSK-011 butir 5).
  9. Report Satisfaction tetap dapat diakses walau CSAT dinonaktifkan; data lama tidak dihapus.
  10. Custom report satisfaction per request type atau per agent dapat dibuat lewat series dan JQL (JSM-QSK-043).
  11. ⚠ Survei juga dapat dikirim lewat kanal chat/Slack bila request berasal dari kanal tersebut (modul 30).
  12. Contoh JQL yang didukung: `Satisfaction = 5`, `"Satisfaction date" >= -7d`.
  13. ⚠ Saat request dibuka ulang lalu di-resolve lagi, survei baru dikirim hanya bila belum ada rating; rating lama dipertahankan dan tidak dikirim ulang.
  14. Teks survei (pertanyaan, pesan terima kasih) dapat diterjemahkan per bahasa portal dan dipilih sesuai bahasa customer.
  15. ⚠ Survei multi-pertanyaan atau per request type tidak tersedia di CSAT; kebutuhan tersebut dipenuhi Surveys native (JSM-QSK-041).
  16. Menonaktifkan CSAT menghentikan pengiriman survei baru seketika; tautan survei yang sudah terkirim kedaluwarsa setelah 30 hari.
  17. Email survei memuat tautan rating satu-klik (1-5) yang membuka halaman portal untuk konfirmasi dan komentar.
  18. Komentar survei hanya dapat dikirim bersama rating; komentar tanpa rating ditolak.
  19. Field Satisfaction bertipe angka 1-5 dan Satisfaction date bertipe tanggal-waktu; keduanya baca-saja bagi agent.
  20. Rating baru tampil di issue view agent paling lambat 1 menit setelah customer mengirim, tanpa reload halaman.
  21. Report Satisfaction dapat difilter rentang tanggal dan menampilkan rata-rata serta tren rating.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template email survei di modul 30 (Customer notifications); terjemahan portal di modul 29. Penempatan sub-fitur fitur ringkas "Customer satisfaction (CSAT)": toggle → butir 16; email & portal → butir 17; rating 1-5 dengan komentar → butir 18; kustomisasi teks dan survei multi-pertanyaan → butir 15; field Satisfaction/Satisfaction date → butir 19; rating di issue view → butir 20; report → butir 21; JQL → butir 12; reopen → butir 13; terjemahan → butir 14. Pengaturan CSAT di project settings juga dirujuk di modul 29 (JSM-PRJ-069 Project settings - Satisfaction).

### JSM-QSK-041 Surveys (Native Surveys)

- **Deskripsi:** Survei native berbasis Forms untuk mengumpulkan feedback organisasi atau karyawan di luar CSAT bintang, misalnya onboarding, exit, dan layanan, beserta review hasilnya.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek, Agent Layanan, Customer/Pelapor (termasuk karyawan), tim HR.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Surveys diaktifkan lewat Project settings > Features > Surveys; default aktif di plan berbayar.
  2. Builder survey memakai Forms (modul 26/29) dengan tipe pertanyaan kustom (pilihan, skala, teks, tanggal).
  3. Tersedia template siap pakai: onboarding, exit, service, dan lainnya.
  4. Pengaturan akses menentukan siapa yang dapat mengelola survey dan siapa yang dapat merespons.
  5. ⚠ Survey dikirim manual ke daftar penerima atau dipicu event lifecycle/layanan lewat action automation "Send survey".
  6. Halaman hasil menampilkan agregasi respons per pertanyaan dan daftar respons individu (sesuai izin).
  7. Pembuatan survey berbantuan AI opsional (draf pertanyaan dari deskripsi tujuan) ditandai eksperimental.
  8. ⚠ Surveys hanya tersedia di plan berbayar dan disembunyikan di Free.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Terpisah dari CSAT (JSM-QSK-040). App survei NPS/CES Marketplace tidak dibangun.

## Laporan, dashboard & analitik layanan

### JSM-QSK-042 Reports - Default Reports

- **Alias inventaris:** Service project reports
- **Deskripsi:** Report bawaan service project yang selalu ada dan tidak bisa diubah: Workload, Satisfaction, Requests deflected, dan Requests resolved. Heading ini juga mencakup fitur ringkas "Service project reports" dari lampiran JSM.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan (melihat), Admin Proyek, Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Workload menampilkan jumlah request terbuka yang di-assign ke tiap agent.
  2. Satisfaction menampilkan rata-rata rating CSAT tim dari waktu ke waktu dan daftar feedback (rating, komentar, request).
  3. Requests deflected menampilkan seberapa sering customer menilai artikel helpful, termasuk jumlah deflection dan views artikel di portal.
  4. Requests resolved menampilkan request yang diselesaikan dengan dan tanpa artikel KB.
  5. Report diakses lewat menu Reports di sidebar; di company-managed default report tidak dapat diedit atau dihapus.
  6. ⚠ Perilaku khas team-managed (report bawaan dapat diedit dan report baru dibuat admin project) tidak dibangun karena clone hanya menyediakan service project company-managed (JSM-PRJ-012); kebutuhan menyunting report dipenuhi custom report (JSM-QSK-043).
  7. Klik titik data membuka daftar work item yang dihitung (dibuka sebagai JQL di navigator).
  8. ⚠ Laporan Ops (MTTA, MTTR, volume alert, beban on-call) tidak termasuk default report dan disediakan oleh modul 34 untuk project yang mengaktifkan Operations.
  9. ⚠ Subset report khusus team-managed (Workload, Satisfaction, Created vs resolved, Time to resolution) tidak dibangun; setiap service project menerima empat default report beserta custom report template (JSM-QSK-043 butir 6).
  10. Report Workload dan Satisfaction termasuk report service desk minimum di Fase 5 (dari jatah Service project reports ringkas).
  11. Report Satisfaction menampilkan distribusi rating 1-5 serta rasio respons terhadap survei terkirim.
  12. Requests deflected dan Requests resolved memakai pemilih rentang tanggal yang sama dan dapat dibuka sebagai daftar work item/artikel.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Penempatan sub-fitur fitur ringkas "Service project reports": Workload → butir 10; SLA goals → JSM-QSK-044 butir 7; SLA success rate dari waktu ke waktu → JSM-QSK-044 butir 8; Created vs Resolved → JSM-QSK-043 butir 12; Resolution by component → JSM-QSK-043 butir 13; Time to first response dan Time to resolution → JSM-QSK-043 butir 10; Satisfaction → butir 11; Requests deflected/resolved by KB → butir 12; Article usage/effectiveness → JSM-QSK-038 butir 6; custom report series/metric/rentang tanggal → JSM-QSK-043 butir 14; report per work category → JSM-QSK-043 butir 11; Ops reports → butir 8; analitik lanjutan → JSM-QSK-047 butir 4; team-managed subset → butir 9.

### JSM-QSK-043 Reports - Custom Report Builder

- **Deskripsi:** Membuat report kustom berbasis series yang masing-masing didefinisikan dengan metric dan filter Basic/JQL; template project juga menyediakan custom report siap pakai.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek (membuat/mengedit); Agent Layanan (hanya melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Di company-managed hanya admin proyek yang dapat membuat custom report; agent melihat tombol New report nonaktif.
  2. Reports > New report meminta nama report lalu "Add series".
  3. Series memilih metric dari dropdown (Created, Resolved, SLA met/SLA breached per SLA, Time to resolution, Satisfaction), label, dan warna.
  4. Filter series memakai Basic (work type, status, priority, component, '+ More') atau 'Switch to JQL'.
  5. Satu report maksimal 20 series; tombol Add series nonaktif setelah batas.
  6. Template project membuat custom report bawaan: Created vs resolved, Time to resolution, SLA goals, SLA met vs breached, dan untuk ITSM: Incidents reported by priority, Mean time to resolve, Resolution by component.
  7. ⚠ Rentang waktu dapat dipilih: past week, past month, past 3/6/12 months, atau custom (tanggal mulai-akhir); default past month.
  8. Grafik time series menampilkan legend yang dapat di-toggle per series.
  9. Custom report dapat diedit dan dihapus dengan konfirmasi.
  10. Metric Time to first response tersedia sebagai series selain Time to resolution, ditampilkan sebagai rata-rata/median durasi per periode (dari jatah Service project reports ringkas).
  11. Report dapat difokuskan per work category (mis. incident, change) dengan filter work category di series atau template report per work category (dari jatah Service project reports ringkas).
  12. Template Created vs resolved menampilkan dua series (Created, Resolved) dalam satu grafik beserta selisihnya.
  13. Template Resolution by component menampilkan jumlah request resolved dikelompokkan per component.
  14. Definisi custom report (series JQL, metric, rentang tanggal) disimpan dan setiap kali dibuka menghitung ulang dengan data terkini.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** Agregasi report memakai data siklus SLA dan rating; disarankan tabel agregat harian per project untuk performa (modul 27).

### JSM-QSK-044 SLAs - SLA Reports

- **Deskripsi:** Report untuk memantau capaian SLA tim, baik dari report bawaan template maupun dari series SLA di custom report.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek, Agent Layanan (melihat), Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Report 'SLA goals' menampilkan persentase met per goal SLA dalam rentang waktu.
  2. Report 'SLA met vs breached' menampilkan tren jumlah met dan breached dari waktu ke waktu.
  3. ⚠ Report 'SLA success rate' menampilkan persentase siklus selesai yang met terhadap seluruh siklus selesai per periode (berbeda dari SLA goals yang dikelompokkan per goal).
  4. Series SLA di custom report menyediakan SLA met, SLA breached, dan persentase per metric SLA.
  5. Filter series SLA memakai JQL (mis. per request type atau priority).
  6. Klik titik data membuka daftar work item yang dihitung.
  7. Report 'SLA goals' menampilkan tabel jumlah met dan breached per goal di bawah grafik.
  8. SLA success rate dapat ditampilkan dari waktu ke waktu dengan granularitas mingguan atau bulanan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Report ini dibuat template project sebagai custom report (JSM-QSK-043 butir 6), bukan bagian dari empat default report.

### JSM-QSK-045 Reports - Export & Sharing

- **Deskripsi:** Cara mengekspor dan membagikan data report untuk analisis di luar produk.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Report service project dapat diekspor sebagai CSV (data series per periode) dan PNG grafik; di produk acuan kemampuan ini khas team-managed, sedangkan clone menyediakannya di service project company-managed (JSM-PRJ-012).
  2. ⚠ Selain ekspor report, data work item tetap dapat diekspor CSV dari navigator (SRC-050) sebagai alternatif.
  3. Setiap report memiliki permalink yang dapat dibagikan ke agent lain; penerima tanpa akses project melihat pesan tidak berizin.
  4. Ringkasan lintas project dilakukan lewat dashboard inti (JSM-QSK-046).
  5. Analitik lanjutan dirujuk ke JSM-QSK-047.
- **Team-managed vs company-managed:** Tidak berlaku (clone hanya membangun service project company-managed, JSM-PRJ-012).
- **Catatan & ketergantungan:** RPT-036 (modul 16) untuk pola export/sharing report inti.

### JSM-QSK-046 JSM Dashboards & Gadgets

- **Deskripsi:** Menampilkan data service project di dashboard lewat gadget inti yang mendukung field service management.
- **Prioritas / Fase / Tier:** P2 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, Admin Proyek, Stakeholder/Viewer (sesuai sharing dashboard).
- **Aturan perilaku (acceptance criteria):**
  1. Gadget Filter Results (DSH-031) mendukung kolom SLA, Request type, dan Satisfaction.
  2. Gadget Pie chart (DSH-035) dapat dikelompokkan berdasarkan Request type, Organizations, atau Satisfaction.
  3. Gadget Created vs resolved (DSH-039) berfungsi untuk filter berbasis service project.
  4. Gadget Two-dimensional filter statistics mendukung field JSM (mis. assignee x status, request type x priority).
  5. Customer tidak dapat mengakses dashboard dalam kondisi apa pun (endpoint dashboard menolak akun customer).
  6. ⚠ Gadget khusus JSM tidak dibangun di Fase 5; kebutuhan SLA/CSAT dipenuhi gadget inti dengan field JSM.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Mesin dashboard di modul 17.

### JSM-QSK-047 Insights & Advanced Analytics

- **Deskripsi:** Analitik lanjutan untuk service project melalui dashboard analitik dan dataset terkurasi yang bisa di-query (padanan Atlassian Analytics/Data Lake).
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs, Admin Proyek.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia template dashboard analitik untuk layanan (SLA, CSAT, beban kerja) di atas dataset terkurasi.
  2. Dataset layanan (work item, siklus SLA, rating, deflection) tersedia untuk di-query lewat ekspor terjadwal atau read-only SQL view.
  3. Akses dibatasi permission data (row-level sesuai Browse projects) dan feature flag plan.
  4. Analitik lanjutan lintas produk tidak dibangun di Fase 5; kebutuhan dasar dipenuhi report dan dashboard (dari jatah Service project reports ringkas, butir Analytics/Data Lake).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Istilah "Insights" ambigu (nama lama Assets dan insight operasi di modul 34); fitur ini hanya analitik layanan. RPT-040 (modul 16) untuk analitik inti.

## Knowledge base: lifecycle artikel, versi & administrasi space

### JSM-QSK-048 ⚠ Knowledge Base - Article Lifecycle, Versions & Space Administration

- **Deskripsi:** Siklus hidup artikel di knowledge space native (draft, publish, unpublish, trash), riwayat versi, struktur hierarki dan label, lampiran, restriksi per artikel, serta pengaturan dan permission space. Fitur ini menjadikan `kb_spaces` dan `kb_articles` (JSM-QSK-031) document store bawaan produk ini yang juga dipakai Pages di modul lain. Fitur ditambahkan di luar inventaris.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin space KB (permission Administer space) mengelola pengaturan, permission, dan trash. Agent Layanan dengan permission tulis space membuat, mengedit, dan mem-publish artikel. Customer/Pelapor hanya membaca artikel eksternal yang sudah dipublikasikan. Admin Proyek (Administer projects) menjadi admin space yang dibuat dari pengaturan KB (JSM-QSK-031 butir 14).
- **Aturan perilaku (acceptance criteria):**
  1. Artikel baru selalu tersimpan sebagai draft dan hanya terlihat oleh pengguna dengan permission tulis. Aksi Publish membuatnya tampil di portal, suggestion, dan panel agent sesuai visibilitas (JSM-QSK-032). Aksi Unpublish mengembalikannya ke draft dan langsung menghapusnya dari hasil portal, suggestion, dan pencarian, dan tautan yang pernah dibagikan ke customer menampilkan pesan "Artikel tidak tersedia".
  2. Mengedit artikel yang sudah dipublikasikan membuat draft perubahan terpisah, sehingga versi published tetap tampil ke customer sampai draft itu di-publish.
  3. Setiap publish membuat versi bernomor berurutan yang menyimpan penulis, waktu, dan catatan perubahan opsional. Halaman History menampilkan daftar versi beserta diff antar dua versi yang dipilih (tambahan dan hapusan disorot). Restore versi lama membuat versi baru dengan isi versi tersebut tanpa menghapus history.
  4. Artikel boleh memiliki satu parent di space yang sama. Urutan antar-sibling diatur dengan drag and drop. Memindahkan artikel ikut memindahkan seluruh subtree-nya, dan sistem menolak pemindahan yang membuat siklus dengan galat "Artikel tidak dapat dipindahkan ke turunannya sendiri". Hierarki ini terpisah dari kategori portal (JSM-QSK-034).
  5. ⚠ Kedalaman hierarki dibatasi 10 level; pemindahan yang melampaui batas ditolak dengan pesan galat.
  6. Artikel dapat diberi label (huruf kecil, tanpa spasi, ⚠ maksimal 50 karakter per label). Articles view (JSM-QSK-033) menyediakan filter label, dan JQL/REST KB dapat mencari artikel per label.
  7. Gambar dan file dapat diunggah ke artikel memakai storage lampiran (CMT-019) dan batas ukuran lampiran site (CMT-016). Media inline dirender di portal. Setiap unduhan lampiran artikel memeriksa izin baca artikel di server, sehingga lampiran artikel internal atau terestriksi tidak dapat diunduh customer lewat URL langsung.
  8. Admin space atau penulis artikel dapat memasang restriksi baca (pengguna/grup tertentu) dan restriksi edit per artikel. ⚠ Restriksi baca diwariskan ke child artikel. Restriksi diperiksa di server untuk portal, pencarian, suggestion, panel agent, ekspor, dan REST API (JSM-QSK-032 butir 4 dan 10).
  9. Halaman Space settings memuat nama, key, deskripsi, ikon, dan visibilitas default artikel baru (internal/eksternal). Permission space terdiri dari admin space (kelola pengaturan, permission, dan trash), penulis/agent (create, edit, publish, unpublish), dan pembaca/customer (view artikel eksternal yang sudah dipublikasikan). Sistem menolak penghapusan admin space terakhir dengan galat "Space harus memiliki minimal satu admin".
  10. Setiap tampilan artikel dicatat dengan tipe viewer (agent, customer, anonim), sumber (portal, suggestion, panel agent, tautan langsung), dan waktu. ⚠ Tampilan berulang oleh pengguna yang sama dalam 30 menit dihitung satu kali. Agregat view per artikel menjadi masukan Article usage/effectiveness (JSM-QSK-038) dan report Requests deflected/resolved (JSM-QSK-042).
  11. Menghapus artikel memindahkannya beserta seluruh child ke trash space. Artikel di trash tidak tampil di portal, suggestion, pencarian, maupun API. Admin space dapat me-restore artikel ke parent asal (atau ke root bila parent sudah tidak ada) atau menghapusnya permanen. ⚠ Artikel di trash dihapus permanen otomatis setelah 30 hari.
  12. Satu artikel atau seluruh space dapat diekspor ke arsip ZIP berisi HTML atau Markdown beserta lampiran.
  13. Impor arsip ZIP berisi Markdown/HTML membuat artikel berstatus draft dengan hierarki mengikuti struktur folder, mengunggah gambar/lampiran yang dirujuk secara relatif, dan menampilkan laporan impor (jumlah berhasil serta daftar berkas gagal beserta alasannya). Fitur ini menjadi jalur migrasi dari wiki eksternal.
  14. ⚠ Setiap publish, unpublish, perubahan restriksi atau permission, hapus, restore, dan impor tercatat di audit log (aktor, waktu, objek).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini tidak ada di inventaris dan ditambahkan karena KB dibangun sebagai entitas native (JSM-QSK-031), sehingga fungsi yang di produk acuan disediakan produk wiki terpisah harus dirumuskan di sini. Model data baru: `kb_article_versions` (article_id, version, body_adf, author_id, created_at, change_note), `kb_article_labels`, `kb_article_restrictions`, `kb_space_permissions`, `kb_article_views`, serta kolom `status` (draft/published), `parent_id`, `position`, dan `trashed_at` di `kb_articles`. Editor mengikuti modul 07, lampiran modul 08 (CMT-016, CMT-019), audit log modul 22, dan REST API KB JSM-AI-029 (modul 35). Pages di modul 04, 08, 11, 25, dan 26 memakai entitas yang sama sebagai document store bawaan (lihat Catatan JSM-QSK-031). Repo saat ini belum memiliki storage berkas atau tabel apa pun untuk konten, sehingga seluruh fitur ini adalah pekerjaan baru.

## Pertanyaan terbuka

- Apakah perlu batas jumlah artikel, kedalaman hierarki, dan masa simpan trash yang berbeda per plan untuk knowledge space (JSM-QSK-048)? Pertanyaan konektor wiki eksternal dipusatkan di [25](25-integrasi-impor-ekspor-ekstensi.md).
- Apakah queue personal (dibuat agent untuk dirinya sendiri) perlu didukung, atau semua queue tetap shared per project?
- Apakah "Duplicate queue" disediakan sebagai aksi native atau cukup lewat ekspor/impor konfigurasi?
- Berapa ambang default "mendekati breach" pada warna SLA (persentase target atau menit tetap) dan apakah dapat dikonfigurasi per metric?
- Bagaimana definisi resmi deflection untuk report: hanya vote helpful, atau termasuk meninggalkan form tanpa membuat request?
- Apakah CSAT boleh diperbarui customer setelah dikirim, dan apakah survei dikirim ulang saat request dibuka ulang?
- Algoritma routing Workforce Management (round-robin, least-loaded, skill-based) mana yang menjadi default, dan apakah jadwal shift berbagi model dengan jadwal on-call modul 34?
- Berapa batas jumlah project dalam satu daftar Multi-space work dan apakah fitur ini dikunci per plan?
- Apakah Surveys native dan Insights & Advanced Analytics termasuk plan Free, berbayar, atau hanya tier tertinggi?

## Di luar cakupan & ekosistem

- App Marketplace pihak ketiga untuk canned responses, survei NPS/CES, SLA tambahan (mis. SLA lintas project/time-in-status), knowledge base pihak ketiga, dan dashboard gadget khusus JSM; produk ini hanya menyediakan fitur native yang dijelaskan di atas dan titik ekstensi (modul 25).
- Produk Confluence sebagai layanan terpisah; knowledge space native (JSM-QSK-031, JSM-QSK-048) menjadi document store bawaan produk ini, sementara konektor wiki eksternal bersifat opsional dan diatur di modul 25.
- Service project team-managed beserta SLA, queue, dan report khas team-managed (JSM-QSK-030, rujukan Fase —); clone hanya membangun service project company-managed (JSM-PRJ-012).
- Atlassian Analytics dan Atlassian Data Lake sebagai layanan lintas produk; padanannya terbatas pada JSM-QSK-047.
- Atlassian Intelligence/Rovo dan Rovo connectors sebagai layanan vendor; AI di modul ini bersifat opsional dan provider-agnostic (JSM-QSK-020, JSM-QSK-039, JSM-QSK-041 butir 7).
- Fitur Data Center-only terkait SLA, seperti mode bulk recalculation non-destructive/graceful/forceful dan pengaturan SLA tingkat server; produk ini hanya menyediakan recalculation di JSM-QSK-029.
- Integrasi Slack dan Microsoft Teams sebagai produk; modul ini hanya memakai action automation/kanal yang disediakan modul 25 dan 30.
- Alert native per queue (tidak ada di produk acuan); pola alert dipenuhi automation (JSM-QSK-028).
- Merge request native (dipenuhi pola link 'duplicates', JSM-QSK-018 butir 6).
