# 32 — Service Management: ITSM (Incident, Problem, Change, Services, Approvals)

Modul ini mendefinisikan praktik **IT Service Management (ITSM)** di atas service project (label UI: "service space") produk ini: template ITSM dengan work category bawaan (Service requests, Incidents, Problems, Changes, Post-incident reviews), workflow dan perilaku workflow khusus Service Management, **service request fulfillment**, **katalog Services** (service registry berbasis Assets) beserta peran, tier, dan dependensi, **Approvals** Service Management (approval step, approver dan approver group, kanal keputusan, riwayat dan JQL, auto-approval), **incident management** (field, queue, major incident, responders, linked alerts, timeline, chat channel, conference call, stakeholder communication, status page native, investigasi, Incident Command Center, AI, Post-incident review), **problem management** dan eskalasi ke tim development, **change management** (tipe change, workflow, CAB, change calendar, change windows, risk assessment, risk insights, deployment tracking, deployment gating), serta panduan agent (Playbooks), orkestrasi lintas tim (Journeys), case management, template otomasi ITSM, dan penyelarasan ITIL 4. Mekanik umum (entitas service project, request type, portal, queue, SLA, knowledge base, Assets, alert/on-call, automation engine) dimiliki modul lain; modul ini mengatur **perilaku khusus proses ITSM** dan cara mekanik tersebut dirangkai.

Kondisi repo saat ini: baru ada fitur `auth` dan `user` (tabel `users`, migrasi `001_create_users_table`). Belum ada entitas project, work item, workflow, service project, Assets, maupun approval. Seluruh fitur di modul ini adalah pekerjaan baru dan diusulkan menjadi feature module `internal/features/itsm` (work category, incident, problem, change, PIR, change calendar & windows, risk assessment), `internal/features/servicecatalog` (services, peran, tier, dependensi; disimpan sebagai object Assets), dan perluasan `internal/features/approval` (yang diusulkan modul 26) untuk approval Service Management, dengan entitas bersama di `internal/shared/domain/entity` dan UI di `frontend/src/routes` (SPA SvelteKit).

**Modul terkait:**

- [00 Visi, Ruang Lingkup & Glosarium](00-visi-ruang-lingkup-glosarium-baseline.md) — istilah Work item/Work type/Space dan tier plan (paket Standard/Premium/Enterprise diperlakukan sebagai feature flag).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission scheme, Link work items (PERM-024), View Development Tools (PERM-027).
- [05 Konfigurasi Work Item & Skema](05-konfigurasi-work-item-skema.md) — work type, custom field, issue link types (CFG-056), field development (CFG-023).
- [06 Work Item Inti](06-work-item-inti.md) — linked work items (ISS-029), clone (ISS-031), remote links (ISS-030).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — activity tabs (CMT-021), tab Approvals (CMT-025), halaman Confluence tertaut (CMT-030, CMT-031).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — workflow editor (WF-021), condition (WF-027), validator (WF-028), post function (WF-029).
- [12 JQL](12-jql-bahasa-kueri.md) — fungsi JQL approval, field development (JQL-035), batas eksekusi (JQL-052).
- [19 Notifikasi & Email](19-notifikasi-email.md) — notifikasi approval, push mobile (NTF-038), Slack/Teams personal (NTF-039, NTF-040).
- [20 Otomasi](20-otomasi.md) — rule builder (AUTO-003), template library (AUTO-010), trigger build/deployment (AUTO-031), smart values (AUTO-063), batas eksekusi (AUTO-015).
- [22 Administrasi Sistem](22-administrasi-sistem.md) — audit log (ADM-014), Teams (ADM-028), development tools admin (ADM-031).
- [23 REST API & Webhook](23-api-publik-webhook.md) — Development Information/Deployments API (API-062, API-063), issue links API (API-023).
- [24 Integrasi DevOps](24-integrasi-devops.md) — data model deployment (DEV-001), development panel (DEV-013), feature flags (DEV-021), deployment gating sisi DevOps (DEV-027).
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — Slack (INT-030 s.d. INT-035), Microsoft Teams (INT-036), Confluence (INT-042 s.d. INT-044).
- [26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) — approval non-JSM (BIZ-042 s.d. BIZ-046) yang menjadi fondasi approval Service Management.
- 29 Service Project, Request Type, Forms & Agent (`29-jsm-service-project-request-type-formulir.md`) — entitas service project, request type, JSM Forms, agent issue view, lisensi agent.
- 30 Help Center, Portal, Customers & Kanal (`30-jsm-help-center-portal-pelanggan-kanal.md`) — portal customer, halaman Approvals di portal, email channel, customer notifications.
- 31 Queues, SLA, Knowledge Base, CSAT & Laporan (`31-jsm-antrian-sla-knowledge-base-laporan.md`) — engine queue, SLA (Time to approve), knowledge base, laporan.
- 33 Assets (CMDB) (`33-jsm-assets-cmdb.md`) — object schema, object type, atribut, permission schema, AQL; services disimpan sebagai object Assets.
- 34 Operations (`34-jsm-operations-alerting-on-call.md`) — alert, on-call schedule, escalation policy, notification rules, operations team, integrasi monitoring.
- 35 Virtual Service Agent, AI, Otomasi, API & CSM (`35-jsm-ai-otomasi-api-admin-csm.md`) — fitur AI/Rovo, automation khusus JSM, REST API JSM.

## Persona terkait

- **Admin Situs** — menyalakan paket/fitur ITSM, menghubungkan integrasi tingkat site (Zoom, Slack, Microsoft Teams, Confluence, CI/CD), mengelola workflow company-managed dan schema Assets untuk services.
- **Admin Proyek/Project Lead** — sebagai service project admin: mengatur work category, request type ITSM, queue, approval step, change management settings, risk assessment, change windows, playbook, journey.
- **Agent Layanan** — menangani service request, incident, problem, dan change; menambah responder dan stakeholder; mengeskalasi ke development; membuat PIR.
- **Manajer Layanan/ITSM** — incident manager, problem manager, change manager, dan anggota CAB: mengoordinasikan major incident, memutuskan approval, memantau change calendar dan risiko.
- **Customer/Pelapor** — melaporkan gangguan, mengajukan service request atau change lewat portal, menjadi approver tanpa lisensi agent.
- **Developer/Anggota Tim** — menerima eskalasi dari service project, menjadi responder on-call, memicu deployment yang di-gate.
- **Stakeholder/Viewer** — menerima update incident dan status page tanpa menjadi responder.
- **Integrator/Developer API** — menghubungkan pipeline CI/CD, tool monitoring, dan status page lewat API.

## User story modul

1. Sebagai Admin Proyek/Project Lead, saya ingin membuat service project dari template IT service management agar work category Incidents, Problems, Changes, dan Post-incident reviews beserta queue dan field ITSM langsung tersedia (JSM-ITSM-001, JSM-ITSM-002).
2. Sebagai Customer/Pelapor, saya ingin mengajukan permintaan hardware lewat portal dan memilih perangkat dari Assets agar permintaan saya diproses dan disetujui tanpa email bolak-balik (JSM-ITSM-003, JSM-ITSM-011).
3. Sebagai Manajer Layanan/ITSM, saya ingin katalog services dengan owner team, tier, responders, dan change approvers agar incident dan change otomatis melibatkan orang yang tepat (JSM-ITSM-004, JSM-ITSM-005, JSM-ITSM-006).
4. Sebagai Customer/Pelapor yang ditunjuk sebagai approver, saya ingin menyetujui atau menolak dari email, portal, atau Slack tanpa lisensi agent agar keputusan cepat dan tercatat (JSM-ITSM-012, JSM-ITSM-013, JSM-ITSM-014).
5. Sebagai Agent Layanan, saya ingin incident dari alert monitoring otomatis terbuat dengan prioritas yang dihitung dari Impact x Urgency agar triase konsisten (JSM-ITSM-016, JSM-ITSM-021, JSM-ITSM-048).
6. Sebagai Manajer Layanan/ITSM, saya ingin menandai major incident lalu otomatis membuat chat channel, conference call, dan memberi tahu stakeholder agar koordinasi respons berjalan dalam hitungan menit (JSM-ITSM-018, JSM-ITSM-023, JSM-ITSM-024, JSM-ITSM-025).
7. Sebagai Agent Layanan, saya ingin menambah team responder sehingga on-call mereka langsung dipanggil dan status acknowledge terlihat di incident (JSM-ITSM-019, JSM-ITSM-020).
8. Sebagai Stakeholder/Viewer, saya ingin menerima update incident dan melihat status page agar saya tahu dampak gangguan tanpa bertanya ke tim responder (JSM-ITSM-025, JSM-ITSM-026).
9. Sebagai Developer/Anggota Tim, saya ingin melihat deployment, commit, dan perubahan feature flag terbaru pada affected services agar penyebab incident cepat ditemukan (JSM-ITSM-028, JSM-ITSM-029).
10. Sebagai Manajer Layanan/ITSM, saya ingin membuat Post-incident review dari incident dengan timeline terisi agar pelajaran dan action items terdokumentasi (JSM-ITSM-022, JSM-ITSM-032).
11. Sebagai Manajer Layanan/ITSM, saya ingin membuat problem dari kumpulan incident serupa dan menerbitkan workaround sebagai artikel known error agar incident berikutnya cepat diselesaikan (JSM-ITSM-033, JSM-ITSM-034).
12. Sebagai Agent Layanan, saya ingin mengeskalasi incident ke tim development dengan membuat linked work item dan memantau statusnya dari service project agar customer otomatis diberi tahu saat bug diperbaiki (JSM-ITSM-035).
13. Sebagai Manajer Layanan/ITSM, saya ingin change normal melalui risk assessment dan approval CAB sementara standard change disetujui otomatis agar beban CAB hanya untuk change berisiko (JSM-ITSM-036, JSM-ITSM-037, JSM-ITSM-039, JSM-ITSM-042).
14. Sebagai Manajer Layanan/ITSM, saya ingin change calendar yang menampilkan change, maintenance window, dan freeze window beserta konfliknya agar jadwal perubahan tidak bentrok (JSM-ITSM-040, JSM-ITSM-041, JSM-ITSM-043).
15. Sebagai Developer/Anggota Tim, saya ingin deployment ke production otomatis membuat change request dan menunggu approval agar kepatuhan change management tidak memperlambat pipeline secara manual (JSM-ITSM-046, JSM-ITSM-047).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 10 |
| P1 | 21 |
| P2 | 14 |
| P3 | 7 |
| **Total** | **52** |

| Fase | Jumlah fitur |
|---|---|
| Fase 5 | 8 |
| Fase 7 | 44 |
| Fase — | 0 |

Strategi urutan implementasi: Fase 5 (Service Management inti) hanya mengambil bagian modul ini yang dibutuhkan service desk umum, yaitu template ITSM dengan work category (JSM-ITSM-001), workflow dan perilaku workflow JSM (JSM-ITSM-002), service request fulfillment (JSM-ITSM-003), dan seluruh rangkaian Approvals (JSM-ITSM-011 s.d. JSM-ITSM-015) yang dibangun di atas fondasi approval modul 26. Fase 7 dimulai dari katalog Services dan field Affected services (JSM-ITSM-004 s.d. JSM-ITSM-009) karena hampir semua fitur incident dan change bergantung padanya, lalu incident inti (field, queue, responders, major incident, linked alerts, timeline) yang membutuhkan modul 34 Operations. Setelah itu problem management, linking, dan developer escalation dibangun, disusul change management (work type, tipe, workflow, CAB, calendar, windows, risk assessment) dan integrasi CI/CD (deployment tracking dan gating) yang bergantung pada modul 24. Fitur kolaborasi eksternal (chat channel, conference call, status page, Statuspage), AI (ICC, AI assistance, AI risk assessment, Incident Prevention Center), Journeys, dan case management ditempatkan di akhir Fase 7 karena bergantung pada integrasi pihak luar atau masih berstatus EAP. Inventaris memuat 11 pasangan fitur tumpang tindih antara lampiran JSM dan domain proses ITSM rinci; keduanya digabung menjadi satu heading dengan baris Alias inventaris sehingga 63 fitur inventaris tercakup dalam 52 heading.

## Fitur

## Fondasi ITSM: template, work category & workflow

### JSM-ITSM-001 ITSM Service Space Template & Work Categories

- **Deskripsi:** Template service project "IT service management" yang menyediakan work category bawaan (Service requests, Incidents, Problems, Changes, Post-incident reviews) untuk mengelompokkan work type, request type, queue, dan halaman khusus ITSM.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (membuat service project company-managed), Admin Proyek/Project Lead sebagai service project admin (mengatur pemetaan kategori), Agent Layanan (memakai halaman per kategori; butuh lisensi agent dan permission Browse projects).
- **Aturan perilaku (acceptance criteria):**
  1. Galeri template service project menyediakan template "IT service management" yang terpisah dari template General service management dan Customer service management; hanya template ITSM yang membuat work category ITSM, work type ITSM, dan field ITSM.
  2. Service project dari template ITSM memiliki tepat lima work category bawaan: Service requests, Incidents, Problems, Changes, Post-incident reviews.
  3. Setiap work type dipetakan ke paling banyak satu work category; request type otomatis mengikuti kategori work type yang dipakainya dan tidak dapat diberi kategori berbeda.
  4. Admin service project dapat memindahkan work type ke kategori lain atau mengosongkan kategorinya dari Project settings > Work types; perubahan langsung memindahkan work item terkait ke halaman kategori baru tanpa mengubah data work item.
  5. Sidebar service project menampilkan entri navigasi per work category (Service requests, Incidents, Problems, Changes, Post-incident reviews), masing-masing membuka list view dan queue khusus kategori tersebut.
  6. Template membuat work type bawaan: Service request, Service request with approvals, Incident, Problem, Change, Post-incident review, masing-masing sudah dipetakan ke kategorinya.
  7. Template membuat field ITSM bawaan: Impact, Urgency, Affected services, Pending reason, Change type, Change risk, Change reason, Change start date, Change completion date, Root cause, Workaround, Approvers, CAB; field ini terpasang pada screen work type yang relevan.
  8. ⚠ Template menyiapkan request type contoh yang dapat diedit atau dihapus: Get IT help, Request new hardware, Request new software, Set up VPN, Report a system problem, Request a change, Investigate a problem, Onboard employee, Offboard employee.
  9. Template membuat queue bawaan per kategori, minimal: All open, Assigned to me, Major incidents, Changes awaiting approval, Open problems.
  10. Kategori Incidents dan Service requests tersedia di semua paket; kategori Changes dan Problems (beserta queue-nya) hanya aktif bila feature flag paket Premium/Enterprise menyala untuk site.
  11. ⚠ Saat paket turun dari Premium ke Standard, sidebar menyembunyikan kategori Changes dan Problems, tetapi work item Change/Problem tetap ada, tetap dapat dicari dengan JQL, dan tampil di queue umum sebagai work item biasa.
  12. ⚠ Kategori tetap tampil di sidebar walaupun tidak ada work type yang dipetakan ke kategori itu, dengan empty state yang menautkan ke Project settings > Work types.
- **Team-managed vs company-managed:** Template ITSM hanya menghasilkan service project company-managed; service project team-managed tidak dibangun (JSM-PRJ-012), sehingga pemetaan work category hanya berlaku untuk company-managed.
- **Catatan & ketergantungan:** Entitas service project dan galeri template di modul 29 (`29-jsm-service-project-request-type-formulir.md`); work type dan field di [05](05-konfigurasi-work-item-skema.md). Istilah UI mengikuti rebranding "work item", "work type", dan label "service space"; PRD tetap menulis "service project". Model data disarankan: kolom `work_category` (`service_request|incident|problem|change|pir|null`) pada tabel `work_types` per service project.

### JSM-ITSM-002 ITSM Workflows & JSM Workflow Behaviors

- **Deskripsi:** Workflow bawaan untuk setiap work type ITSM dan perilaku workflow khusus Service Management (approval step, status customer-facing, Pending reason, kaitan SLA).
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengedit workflow company-managed, permission Administer Jira), Admin Proyek/Project Lead (memilih workflow scheme service project).
- **Aturan perilaku (acceptance criteria):**
  1. Template ITSM membuat enam workflow bawaan yang terpasang di workflow scheme service project: Service request fulfilment, Service request fulfilment with approvals, Incident management, Problem management, Change management, Post-incident review.
  2. Status mana pun pada workflow service project dapat dijadikan approval step (dirinci di JSM-ITSM-011); workflow "with approvals" dan Change sudah memuat approval step bawaan.
  3. Setiap status memiliki nama customer-facing yang dapat berbeda dari nama internal; portal dan customer notification hanya menampilkan nama customer-facing.
  4. Saat transisi ke status Pending, layar transisi meminta field Pending reason; nilai Pending reason ditampilkan di issue view dan dikosongkan otomatis saat work item keluar dari status Pending.
  5. SLA dapat dikonfigurasi untuk start, stop, dan pause berdasarkan masuk/keluar status workflow ITSM (mis. pause saat Waiting for customer atau Pending).
  6. Workflow bawaan memuat post function standar (set resolution saat ke status Done-category, kosongkan resolution saat dibuka ulang, update change history, fire event) dan tidak dapat dihapus selama status terkait masih ada.
  7. Transisi hasil approval (Approve/Decline) dijalankan oleh sistem dan tidak tampil sebagai tombol transisi manual di agent issue view.
  8. Saat admin mengubah workflow (menghapus/mengganti nama status) yang direferensikan oleh SLA atau approval step, editor menampilkan peringatan berisi daftar SLA dan approval yang terdampak sebelum disimpan; SLA yang berjalan dihitung ulang mengikuti konfigurasi baru.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); workflow ITSM hanya memakai workflow scheme dan workflow editor global company-managed.
- **Catatan & ketergantungan:** Engine workflow di [09](09-workflow-status-transisi.md) (WF-021, WF-029); SLA di modul 31 (`31-jsm-antrian-sla-knowledge-base-laporan.md`); detail workflow Change di JSM-ITSM-038 dan workflow Problem di JSM-ITSM-033.

### JSM-ITSM-003 Service Request Fulfillment

- **Deskripsi:** Penanganan service request standar (permintaan akses, hardware, software) melalui request type, form, approval, dan automation.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (membuat request via portal), Agent Layanan (memproses; lisensi agent), approver (tanpa lisensi agent).
- **Aturan perilaku (acceptance criteria):**
  1. Work type Service request dan Service request with approvals tersedia di kategori Service requests; keduanya memakai workflow yang berbeda (JSM-ITSM-002).
  2. Template menyediakan request type contoh: Get IT help, Request new hardware, Request new software, Set up VPN, Request admin access, Onboard employee, Offboard employee.
  3. ⚠ Workflow Service request fulfilment memiliki status Waiting for support, In progress, Waiting for customer, Pending, Resolved, Canceled, Closed, dengan transisi Resolve, Cancel, dan Reopen yang tersedia bagi agent.
  4. Workflow Service request fulfilment with approvals menambahkan status Waiting for approval sebagai approval step sebelum In progress.
  5. Setiap request type dapat memiliki form/field sendiri yang berfungsi sebagai template permintaan; field wajib pada form harus diisi sebelum request dapat dikirim.
  6. Request type dapat memuat field Assets object (mis. perangkat atau lisensi) yang dibatasi dengan AQL filter sehingga customer hanya dapat memilih object yang relevan.
  7. Automation untuk service request (auto-assign, auto-transition, notifikasi) dapat dipasang per request type dan tersedia sebagai template di JSM-ITSM-048.
  8. Fulfillment multi-langkah dalam satu proyek memakai subtask atau checklist; proses lintas proyek diarahkan ke Journeys (JSM-ITSM-050) dan panduan langkah agent ke Playbooks (JSM-ITSM-049).
  9. Bila approver menolak, request bertransisi ke status Declined (atau Canceled pada workflow tanpa status Declined) dan customer menerima notifikasi penolakan.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); hanya service project company-managed.
- **Catatan & ketergantungan:** Request type dan JSM Forms di modul 29; portal di modul 30; field Assets object di modul 33 (`33-jsm-assets-cmdb.md`); approval di JSM-ITSM-011.

## Services: katalog layanan & peran

### JSM-ITSM-004 Services (Service Registry in Assets)

- **Alias inventaris:** Services & service registry
- **Deskripsi:** Katalog service terpusat yang disimpan sebagai object pada schema Assets bawaan, dipakai sebagai affected services pada incident, problem, dan change serta sebagai sumber pemilik, tier, responders, dan dependensi.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM sebagai service owner, Admin Situs/Assets admin (mengelola schema), Agent Layanan (melihat dan memilih service). Akses diatur oleh permission object schema Assets (JSM-ITSM-009).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Services dapat diakses dari navigasi service project dan dari menu Operations; halaman menampilkan daftar service dengan kolom nama, tipe, tier, dan owner team.
  2. Pengguna berizin dapat membuat, mengedit, dan menghapus service dengan atribut name (wajib, unik per site), description, service type, tier, dan owner team; nama kosong atau duplikat ditolak dengan pesan galat.
  3. Service type bernilai salah satu dari Business service, Application, Capability, Software service.
  4. Saat fitur services diaktifkan, sistem memprovisi dua schema Assets bawaan, Business Portfolio dan Software Registry, yang memuat object type untuk tiap service type.
  5. Admin Assets dapat menambah custom attribute pada object type service dan menambah custom object type di schema services; atribut tambahan tampil di detail service.
  6. Halaman detail service menampilkan incident, change, dan alert terkait (yang mereferensikan service tersebut di Affected services) beserta statusnya.
  7. Daftar service dapat dicari berdasarkan nama dan difilter berdasarkan service type, tier, dan owner team.
  8. ⚠ Detail service dapat menautkan proyek dan repository terkait; tautan ini dipakai oleh Incident Investigation (JSM-ITSM-028) dan deployment tracking (JSM-ITSM-046).
  9. ⚠ Admin dapat mengimpor services dari sumber yang sudah dikenal graph data internal (mis. komponen atau repository terhubung) dengan wizard pemetaan sumber ke atribut service object.
  10. ⚠ Service dari registry lama (sebelum berbasis Assets) dimigrasikan sekali jalan ke service object Assets dengan mempertahankan nama, tipe, tier, owner team, dan relasi; referensi Affected services pada work item lama dipetakan ke object baru.
  11. ⚠ Relasi dependensi melingkar (A bergantung pada B dan B bergantung pada A) disimpan tanpa error; visualisasi dan analisis dampak berhenti pada node yang sudah dikunjungi.
  12. Formulir create service menerima nama, deskripsi, tipe (Applications, Business services, Capabilities, Software services), dan tier 1-4 dalam satu langkah.
  13. Owner team, responders, dan change approvers dapat diatur per service saat create atau edit (dirinci di JSM-ITSM-006).
  14. Dependensi antar service dapat ditambahkan dari detail service (dirinci di JSM-ITSM-008).
  15. Service yang dibuat langsung tersedia sebagai pilihan pada field Affected services (JSM-ITSM-005).
  16. ⚠ Sinkronisasi dengan komponen dari katalog developer eksternal ditangani di JSM-ITSM-010 dan dapat dimatikan per site.
  17. ⚠ Setelah migrasi ke Assets, halaman registry lama hanya-baca dan menampilkan tautan ke service object baru.
  18. Pencarian service mendukung pencarian teks pada nama dan deskripsi dengan hasil maksimal 50 per halaman dan paginasi.
  19. Saat service ditetapkan sebagai affected service pada incident, responders service tersebut dinotifikasi (JSM-ITSM-019).
- **Team-managed vs company-managed:** Tidak ada perbedaan; services berlaku tingkat site dan dapat dipakai oleh semua service project.
- **Catatan & ketergantungan:** Engine Assets, object schema, dan atribut di modul 33 (`33-jsm-assets-cmdb.md`). Fitur services tidak tersedia di paket Free (feature flag). Import berbasis graph data lintas produk Atlassian tidak dibangun; padanannya adalah import dari sumber internal produk ini (repository terhubung di [24](24-integrasi-devops.md), komponen di modul 15).

### JSM-ITSM-005 Affected Services Field

- **Deskripsi:** Field multi-value yang menghubungkan incident, problem, dan change dengan service object di Assets.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dan Manajer Layanan/ITSM (mengisi dari issue view), Customer/Pelapor (mengisi dari portal bila ditampilkan), Admin Proyek/Project Lead (menampilkan field pada request type).
- **Aturan perilaku (acceptance criteria):**
  1. Field "Affected services" bersifat multi-value dan terpasang pada work type Incident, Problem, dan Change dari template ITSM.
  2. Nilai field hanya dapat dipilih dari service object di schema services Assets; nilai teks bebas ditolak.
  3. Admin dapat menampilkan field pada request type portal sehingga customer dapat memilih service; customer hanya melihat service yang diizinkan permission schema (JSM-ITSM-009).
  4. JQL mendukung pencarian `"Affected services" = "<nama service>"`, `in (...)`, `is EMPTY`, dan `is not EMPTY`.
  5. Saat service ditambahkan ke field, sistem otomatis menambahkan responders (incident), stakeholders (incident, Premium/Enterprise), dan change approvers (change) dari peran service tersebut; peran kosong tidak menambah apa pun.
  6. Automation menyediakan smart value untuk atribut service terdampak (mis. nama, tier, owner team) agar dapat dipakai di condition dan action.
  7. ⚠ Field dikunci dari kustomisasi tertentu: tidak dapat dihapus, diganti tipenya, atau diganti sumber datanya, tetapi dapat diganti nama tampilan dan konteksnya.
  8. Bila service object dihapus, work item yang mereferensikannya menampilkan entri "Service dihapus" pada field tanpa error, dan referensi kosong tersebut dapat dibersihkan oleh agent.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** JQL di [12](12-jql-bahasa-kueri.md); smart value di [20](20-otomasi.md) (AUTO-063); field Assets object di modul 33.

### JSM-ITSM-006 Service Roles: Owner Team, Change Approvers, Responders, Stakeholders

- **Deskripsi:** Peran per service yang menentukan pemilik, approver change, responder, dan stakeholder yang otomatis dilibatkan saat service terdampak.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM sebagai service owner (mengatur peran; butuh izin edit object service), Agent Layanan (melihat).
- **Aturan perilaku (acceptance criteria):**
  1. Owner team adalah satu team (Teams) yang bertanggung jawab atas service dan ditampilkan di detail service serta di panel incident/change.
  2. Change approvers (user dan/atau group) wajib menyetujui change yang menyertakan service tersebut di Affected services; mereka ditambahkan ke field Approvers/Approver groups change secara otomatis.
  3. Responders (user dan/atau team) diberi tahu saat service menjadi affected service pada incident dan ditambahkan ke field Responders.
  4. Stakeholders (user dan/atau stakeholder group) ditambahkan ke field Stakeholders incident saat service menjadi affected service; perilaku ini hanya aktif pada paket Premium/Enterprise.
  5. Semua peran opsional saat pembuatan service dan dapat ditambah atau diubah kemudian sebagai atribut object Assets.
  6. Peran yang kosong tidak memicu penambahan otomatis apa pun dan tidak menimbulkan peringatan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Teams di [22](22-administrasi-sistem.md) (ADM-028); responders di JSM-ITSM-019; stakeholders di JSM-ITSM-025; CAB di JSM-ITSM-039.

### JSM-ITSM-007 Service Tiers

- **Deskripsi:** Klasifikasi kritikalitas service dari Tier 1 sampai Tier 4 yang memengaruhi prioritas incident dan risiko change.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM sebagai service owner (menetapkan tier), Admin Proyek/Project Lead (memakai tier di risk assessment dan automation).
- **Aturan perilaku (acceptance criteria):**
  1. Tier bernilai Tier 1 (paling kritikal), Tier 2, Tier 3, atau Tier 4; nilai default saat create adalah Tier 4 dan setiap tier memiliki deskripsi bawaan yang dapat diedit admin.
  2. Tier service terdampak tersedia sebagai input kriteria pada Automated Change Risk Assessment (JSM-ITSM-042).
  3. Template automation tersedia untuk menaikkan prioritas incident (mis. Highest) bila salah satu affected service bertier 1.
  4. Daftar services, list view incident/change, dan laporan dapat difilter dan dikelompokkan per tier service.
  5. Mengubah tier service tidak mengubah prioritas atau risiko work item yang sudah ada; hanya evaluasi berikutnya yang memakai tier baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Deskripsi baku tiap tier menjadi keputusan produk. Laporan di modul 31.

### JSM-ITSM-008 Service Dependencies

- **Deskripsi:** Pemetaan relasi antar service (depends on / used by) untuk analisis dampak.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM sebagai service owner (mengelola relasi), Agent Layanan dan change manager (melihat dampak).
- **Aturan perilaku (acceptance criteria):**
  1. Detail service memungkinkan menambah dependency ("depends on") dan dependent ("used by"); relasi disimpan satu kali dan tampil dua arah.
  2. ⚠ Detail service menampilkan visualisasi relasi (graph) satu tingkat secara default dengan opsi memperluas ke tingkat berikutnya; graph object Assets juga menampilkan relasi yang sama.
  3. Analisis dampak menampilkan daftar service downstream (yang bergantung, langsung maupun tidak langsung) dari service yang sedang terganggu; daftar ini tampil di incident dan change yang memuat service tersebut.
  4. Menghapus service ikut menghapus seluruh relasi dependency/dependent yang melibatkan service tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Relasi object di modul 33; dependensi melingkar di JSM-ITSM-004 butir 11.

### JSM-ITSM-009 Service Permissions

- **Deskripsi:** Kontrol siapa yang dapat membuat, mengubah, dan melihat service melalui permission object schema Assets.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs/Assets admin, Manajer Layanan/ITSM, Agent Layanan, Customer/Pelapor.
- **Aturan perilaku (acceptance criteria):**
  1. Hak akses ke service object diatur oleh role object schema Assets (Object schema manager, Object schema user, Object viewer) pada schema services.
  2. Hanya admin JSM/Assets dan pengguna dengan role manager/user pada schema services yang dapat membuat, mengedit, atau menghapus service.
  3. Agent dengan akses lihat schema dapat melihat service dan memilihnya di field Affected services.
  4. Customer dapat memilih service di portal hanya bila field ditampilkan pada request type dan schema services mengizinkan akses customer (mis. Object viewer untuk customer).
  5. Pengguna tanpa akses schema tidak dapat membuka detail service; field Affected services pada work item tetap menampilkan nama service tanpa tautan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Role dan permission Assets di modul 33; permission service project di [03](03-izin-peran-keamanan.md).

### JSM-ITSM-010 Compass Integration for Services

- **Deskripsi:** Keterkaitan komponen dari katalog komponen developer (padanan produk katalog komponen Atlassian) dengan service Service Management.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Admin Situs (menghubungkan katalog), Manajer Layanan/ITSM sebagai service owner, Developer/Anggota Tim.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat menghubungkan katalog komponen developer ke service management dari pengaturan services; koneksi dapat diputus kapan saja.
  2. ⚠ Komponen dapat ditautkan ke service yang sudah ada atau diimpor sebagai service baru; arah sinkronisasi default satu arah dari katalog komponen ke service.
  3. ⚠ Owner team dan dependencies komponen dipetakan ke owner team dan dependency service.
  4. ⚠ Bila komponen dihapus di katalog, service yang tertaut tidak dihapus; service diberi penanda "Sumber terputus" dan tautannya dilepas.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk Compass Atlassian berada di luar cakupan; fitur ini hanya diimplementasikan bila produk ini membangun katalog komponen sendiri atau lewat integrasi pihak ketiga via API ([23](23-api-publik-webhook.md)). Lihat bagian Di luar cakupan.

## Approvals Service Management

### JSM-ITSM-011 Approval Step in Workflow

- **Alias inventaris:** Approvals
- **Deskripsi:** Konfigurasi status workflow service project sebagai approval step lengkap dengan sumber approver, aturan jumlah persetujuan, transisi Approve/Decline, dan pengecualian approver; mencakup ringkasan perilaku approval Service Management end-to-end.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Admin Situs (mengonfigurasi approval step pada workflow company-managed, permission Administer Jira), Admin Proyek/Project Lead (memilih workflow scheme yang memuat approval step), approver (tanpa lisensi agent), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Di workflow editor, admin memilih status lalu mengaktifkan opsi "Include approval step"; opsi hanya tersedia pada workflow yang dipakai service project.
  2. Layar konfigurasi pertama meminta sumber approver berupa custom field user picker atau group picker yang sudah ada; bila belum ada field yang sesuai, layar menampilkan tautan untuk membuat field terlebih dahulu dan tombol simpan dinonaktifkan.
  3. Aturan jumlah persetujuan dapat dipilih: satu approver, semua approver, persentase approver, atau jumlah tertentu (N); nilai N lebih besar dari jumlah approver saat runtime diperlakukan sebagai "semua approver".
  4. Untuk sumber multi-group picker, admin dapat menetapkan jumlah persetujuan tertentu per grup (numberPerPrincipal); approval selesai hanya bila setiap grup memenuhi jumlahnya.
  5. Layar konfigurasi kedua mewajibkan pemilihan transisi untuk hasil Approve dan hasil Decline; kedua transisi harus keluar dari status approval tersebut.
  6. Layar konfigurasi kedua menyediakan opsi "Exclude approvers from" dengan pilihan Assignee dan/atau Reporter; user yang dikecualikan tidak dapat memberi keputusan walau tercantum di field approver.
  7. Secara default satu Decline langsung menolak keseluruhan approval dan menjalankan transisi Decline, terlepas dari aturan jumlah.
  8. Workflow dapat memiliki beberapa approval step berurutan (mis. manager lalu CAB); setiap step dievaluasi independen dengan approver dan aturannya sendiri.
  9. Opsi exclusion hanya mendukung Assignee, Reporter, dan field user picker standar; atribut object Assets tidak dapat dipakai sebagai sumber exclusion.
  10. Bila field approver kosong saat work item masuk approval step, panel approval menampilkan peringatan "Belum ada approver"; approval tidak dapat selesai sampai field diisi, dan queue "Approval tanpa approver" menampilkan work item tersebut.
  11. Bila Reporter termasuk dalam exclusion dan reporter adalah satu-satunya approver, panel approval menampilkan peringatan "Tidak ada approver yang memenuhi syarat".
  12. Approval diaktifkan melalui Configure approval pada status workflow company-managed; mode konfigurasi approval di workflow request type (team-managed) tidak dibangun karena service project team-managed tidak didukung (JSM-PRJ-012).
  13. Sumber approver dapat berupa field user picker (mis. Approvers), group picker (Approver groups), atau atribut object Assets bertipe user (mis. manager pada object karyawan).
  14. Aturan jumlah dapat berupa semua approver atau sejumlah N approver/anggota grup.
  15. Hasil Approve dan Decline dipetakan ke transisi workflow dan dijalankan otomatis setelah aturan terpenuhi.
  16. Approver dapat memutuskan via portal, email, issue view, Slack/Teams, dan aplikasi mobile (dirinci di JSM-ITSM-013).
  17. Saat work item masuk approval step, setiap approver menerima notifikasi "Approval required".
  18. Panel approval di issue view menampilkan status pending/approved/declined per approver beserta waktu keputusan.
  19. Approver dapat diganti selama approval berjalan oleh pengguna yang berhak mengedit field approver; keputusan approver yang tetap tercantum dipertahankan.
  20. Approver tidak memerlukan lisensi agent; customer portal dapat menjadi approver.
  21. JQL mendukung fungsi `approval = approved()`, `approver(user)`, `decidedBy(user)`, `myApproval()`, `myPending()`, `pending()`, `pendingBy(user)`, `rejected()` (dirinci di JSM-ITSM-014).
  22. Automation menyediakan trigger dan condition terkait approval (approval required, approval completed, approval decision) (dirinci di JSM-ITSM-015).
  23. ⚠ Reporter dapat dikecualikan dari approver sehingga pembuat request tidak dapat menyetujui permintaannya sendiri.
  24. SLA "Time to approve" dapat dikonfigurasi untuk change dengan start saat masuk approval step dan stop saat approval selesai.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); approval step hanya diatur di workflow editor company-managed per status dengan opsi lengkap (persentase, numberPerPrincipal, exclusion).
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Approvals" dan domain proses ITSM "Approval Step in Workflow"). Fondasi approval non-JSM di BIZ-042 s.d. BIZ-046 ([26](26-kerja-bisnis-formulir-persetujuan.md)); model data `workflow_status_approvals` diperluas dengan kolom `rule` (`one|all|percent|count`), `threshold`, `per_principal`, `exclude_assignee`, `exclude_reporter`. Workflow editor di [09](09-workflow-status-transisi.md) (WF-021); SLA di modul 31. Catatan impor: approval dari service project team-managed pada data sumber (dikonfigurasi di workflow request type) dipetakan ke approval step company-managed dengan aturan jumlah setara saat impor.

### JSM-ITSM-012 Approvers Field & Approver Groups

- **Deskripsi:** Field standar untuk menyimpan approver per request dan cara pengisiannya (portal, agent, automation, Assets, services).
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Customer/Pelapor (mengisi di portal), Agent Layanan (mengedit), Admin Proyek/Project Lead (menampilkan field pada request type), approver.
- **Aturan perilaku (acceptance criteria):**
  1. Field "Approvers" bertipe multi user picker tersedia di service project ITSM.
  2. Field "Approver groups" bertipe multi group picker tersedia di service project ITSM.
  3. Field dapat diisi customer di portal (bila ditampilkan pada request type), oleh agent di issue view, atau oleh automation action Edit work item.
  4. Field dapat diisi dari atribut object Assets (mis. manager karyawan, service owner) melalui automation atau konfigurasi approval step berbasis atribut Assets.
  5. Saat Affected services diisi pada change, change approvers dari service tersebut ditambahkan ke field Approvers/Approver groups (JSM-ITSM-006).
  6. User yang ditambahkan sebagai approver tidak memerlukan lisensi agent dan tidak dihitung sebagai pengguna berlisensi.
  7. User yang dinonaktifkan tetap tercantum di field tetapi ditandai "Nonaktif", tidak menerima notifikasi, tidak dapat memberi keputusan, dan tidak dihitung dalam aturan "semua approver".
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); field approver hanya tersedia sebagai field global company-managed.
- **Catatan & ketergantungan:** Field user/group picker di [05](05-konfigurasi-work-item-skema.md); lisensi agent di modul 29; Assets di modul 33.

### JSM-ITSM-013 Approval Channels: Portal, Email, Slack, Teams, Mobile

- **Deskripsi:** Approver dapat menyetujui atau menolak permintaan dari berbagai kanal dengan hasil yang identik.
- **Prioritas / Fase / Tier:** P0 / Fase 5 / jsm
- **Aktor & izin:** Approver (customer atau agent), Agent Layanan.
- **Aturan perilaku (acceptance criteria):**
  1. Portal menyediakan halaman Approvals yang mendaftar semua approval pending milik pengguna, diurutkan dari yang paling lama menunggu.
  2. Detail request di portal menampilkan tombol Approve dan Decline bagi approver yang memenuhi syarat.
  3. Email notifikasi approval memuat ringkasan request dan tautan ke halaman approval; tautan mewajibkan login bila sesi tidak aktif.
  4. ⚠ Approver dapat memutuskan dengan membalas email notifikasi berisi kata kunci "approve" atau "decline" pada baris pertama; balasan dari alamat yang bukan approver diabaikan dan dicatat.
  5. Notifikasi Slack dan Microsoft Teams memuat tombol Approve dan Decline yang menjalankan keputusan tanpa membuka browser, bila akun chat pengguna sudah terhubung.
  6. Aplikasi mobile menampilkan daftar approval pending dan tombol Approve/Decline dengan push notification saat approval baru.
  7. Semua kanal menyediakan kolom komentar opsional saat memutuskan; komentar disimpan bersama keputusan.
  8. Bila approval sudah selesai (oleh approver lain atau karena aturan terpenuhi), membuka kanal mana pun menampilkan status final dan tombol keputusan tidak lagi tersedia; keputusan susulan ditolak dengan pesan "Approval sudah selesai".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Portal di modul 30; email channel inbound di modul 30; Slack/Teams di [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-030, INT-036) dan [19](19-notifikasi-email.md) (NTF-038 s.d. NTF-040).

### JSM-ITSM-014 Approval History & Approval JQL

- **Deskripsi:** Riwayat keputusan approval untuk audit dan pencarian approval melalui fungsi JQL.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Agent Layanan, approver, Manajer Layanan/ITSM (audit); Customer/Pelapor melihat riwayat di portal.
- **Aturan perilaku (acceptance criteria):**
  1. Panel approval di issue view menampilkan setiap approver, keputusannya (pending/approved/declined), komentar, dan waktu keputusan.
  2. Tab Approvals di bagian Activity menampilkan kronologi semua approval step yang pernah dijalankan pada work item.
  3. Riwayat approval tetap tersimpan dan tidak dapat diedit setelah work item berpindah status atau approval step dihapus dari workflow.
  4. Portal menampilkan riwayat approval (siapa menyetujui/menolak dan kapan) kepada reporter dan request participants.
  5. ⚠ JQL mendukung fungsi `approval = pending()`, `approval = approved()`, `approval = rejected()`, `approval = myApproval()`, `approval = myPending()`, `approval = approver(user)`, `approval = pendingBy(user)`, `approval = decidedBy(user)`; argumen user tidak valid menghasilkan galat validasi JQL.
  6. Template ITSM menyediakan queue "Waiting for approval" berbasis `approval = pending()`.
  7. Bila work item masuk approval step yang sama untuk kedua kalinya, sistem membuat entri approval baru sehingga riwayat memuat beberapa entri terpisah dengan nomor urut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tab Approvals CMT-025 ([08](08-komentar-lampiran-aktivitas.md)); fungsi JQL didaftarkan di [12](12-jql-bahasa-kueri.md); queue di modul 31; audit log di [22](22-administrasi-sistem.md) (ADM-014).

### JSM-ITSM-015 Auto-approval & Approval Automation

- **Deskripsi:** Penyetujuan otomatis dan pengaturan approver lewat automation rule.
- **Prioritas / Fase / Tier:** P1 / Fase 5 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead (membuat rule; permission Administer projects), actor automation.
- **Aturan perilaku (acceptance criteria):**
  1. Automation menyediakan action "Approve/Decline request" yang memberi keputusan pada approval step aktif atas nama actor rule.
  2. Automation menyediakan trigger "Approval required" yang berjalan saat work item masuk approval step, serta trigger "Approval completed" dengan smart value keputusan.
  3. Template "Auto-approve standard change / risiko rendah" menyetujui change dengan Change type = Standard atau Change risk = Low secara otomatis.
  4. Template "Set approver from reporter's manager" mengisi field Approvers dari atribut manager (Assets atau direktori pengguna) milik reporter.
  5. Template pengingat mengirim notifikasi ulang ke approver yang belum memutuskan setelah durasi tertentu (mis. 24 jam).
  6. Template eskalasi menambah approver cadangan atau mengubah prioritas bila approval melewati batas waktu yang ditetapkan.
  7. Action approval gagal dengan entri audit log "Actor tidak berizin" bila actor rule bukan approver yang memenuhi syarat dan tidak memiliki permission yang dibutuhkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine automation di [20](20-otomasi.md) (AUTO-003, AUTO-010, AUTO-014); template ITSM lain di JSM-ITSM-048.

## Incident: pencatatan, queue & respons

### JSM-ITSM-016 Incident Work Type & Incident Fields

- **Alias inventaris:** Incident management
- **Deskripsi:** Work type Incident beserta request type portal dan field ITSM untuk mengklasifikasi, memprioritaskan, dan melacak gangguan layanan dari deteksi hingga pemulihan; heading ini juga merangkum alur incident management end-to-end.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Customer/Pelapor (melapor via portal/email), Agent Layanan (membuat dan menangani; lisensi agent), responder dan on-call (modul 34), Manajer Layanan/ITSM sebagai incident manager, Admin Proyek/Project Lead (konfigurasi field dan request type).
- **Aturan perilaku (acceptance criteria):**
  1. Template menyediakan request type portal default "Report a system problem" yang membuat work item bertipe Incident; agent juga dapat membuat incident dari tombol "Create incident" di halaman Incidents.
  2. Incident dapat dibuat dari portal, email channel, chat, REST API, alert dari Operations/integrasi monitoring, dan automation; kanal asal tercatat di field Request channel.
  3. ⚠ Field Impact memiliki opsi bawaan Extensive/Widespread, Significant/Large, Moderate/Limited, Minor/Localized yang dapat diganti nama oleh admin.
  4. Field Urgency memiliki opsi Critical, High, Medium, Low.
  5. Field Priority dapat diisi manual atau dihitung otomatis dari matriks Impact x Urgency melalui template automation (JSM-ITSM-048).
  6. Field Affected services berupa service picker dari Assets services (JSM-ITSM-005).
  7. Field Pending reason wajib diisi saat incident masuk status Pending.
  8. Field Major incident tersedia hanya pada paket Premium/Enterprise (JSM-ITSM-018).
  9. Field Responders menerima user dan team (JSM-ITSM-019).
  10. Field Stakeholders tersedia hanya pada paket Premium/Enterprise (JSM-ITSM-025).
  11. Field Conference call tersedia hanya pada paket Premium/Enterprise (JSM-ITSM-024).
  12. Panel Alerts menampilkan linked alerts bila Operations aktif di site (JSM-ITSM-021).
  13. Incident juga memakai field inti Service Management: Request type, Organizations, Request participants, Components, Labels, Satisfaction.
  14. Incident yang dibuat dari alert mewarisi summary, deskripsi, dan priority dari alert; perubahan alert setelahnya tidak menimpa field incident kecuali diatur sinkronisasi (JSM-ITSM-020).
  15. Work category Incidents memakai workflow Incident management bawaan (JSM-ITSM-002).
  16. Fitur "create incident from alert" tersedia dari halaman alert dan dari panel alert di issue view.
  17. Field Priority, Impact, Urgency, Affected services, dan Responders tampil pada create screen agent untuk work type Incident.
  18. ⚠ Toggle Major incident menampilkan incident pada queue/filter Major incidents; ketersediaan per paket mengikuti feature flag Premium/Enterprise.
  19. Menambah team/individu on-call sebagai responder memicu notifikasi lewat Operations.
  20. Status alert dan incident tersinkron sesuai konfigurasi (mis. alert ditutup saat incident resolved).
  21. Issue view incident menampilkan activity timeline (JSM-ITSM-022).
  22. Video conference bridge (Zoom, Microsoft Teams, Google Meet) dapat dimulai dari issue view sesuai integrasi yang terhubung (JSM-ITSM-024).
  23. War room chat: agent dapat membuat channel Slack/Teams khusus incident dan mengundang responders (JSM-ITSM-023).
  24. Investigasi menampilkan deployment dan commit terkait (Bitbucket, GitHub, GitLab) di sekitar waktu incident (JSM-ITSM-028).
  25. ⚠ Stakeholder updates dapat dikirim dari incident; ketersediaan mengikuti paket Premium/Enterprise (JSM-ITSM-025).
  26. ⚠ Peran incident seperti incident commander dapat ditetapkan melalui custom field user picker "Incident commander" yang ditampilkan di panel incident; tidak ada field native bawaan.
  27. Fitur AI similar incidents dan incident summary tersedia bila AI aktif (JSM-ITSM-030).
  28. Incident dapat ditautkan ke problem dari issue view (JSM-ITSM-034).
  29. Queue Incidents menyediakan filter cepat all/open/major.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); incident management hanya tersedia pada service project company-managed dari template ITSM.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Incident management" dan domain proses ITSM "Incident Work Type & Incident Fields"). Incident work category, queue, alert, dan on-call tersedia di semua paket; major incident, stakeholders, dan conference call di Premium/Enterprise. Alert di modul 34 (`34-jsm-operations-alerting-on-call.md`); email channel dan request channel di modul 30.

### JSM-ITSM-017 Incident Queues & List Views

- **Deskripsi:** Halaman Incidents dengan list view, queue, dan filter khusus untuk memantau incident aktif.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Manajer Layanan/ITSM, Admin Proyek/Project Lead (mengelola queue; permission Administer projects).
- **Aturan perilaku (acceptance criteria):**
  1. List view incident menampilkan kolom status, priority, assignee, affected services, responders, major incident, dan created; kolom dapat diatur ulang dan disortir.
  2. Filter cepat tersedia untuk status, priority, affected service, assignee, dan major incident, dan dapat dikombinasikan.
  3. Template menyediakan queue incident terbuka dan queue "Major incidents"; queue Major incidents hanya tampil pada paket Premium/Enterprise.
  4. Admin dapat membuat dan mengedit queue incident dengan JQL memakai engine queue inti Service Management; JQL tidak valid ditolak saat simpan.
  5. Tombol "Create incident" tersedia di halaman Incidents dan membuka create screen dengan work type Incident terpilih.
  6. Queue Major incidents menampilkan aksi untuk membuka atau memulai conference call per baris (Premium/Enterprise).
  7. Incident yang dipindah ke work type di luar kategori Incidents otomatis tidak tampil lagi di halaman Incidents.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Engine queue di modul 31; conference call di JSM-ITSM-024.

### JSM-ITSM-018 Major Incident Management

- **Deskripsi:** Menandai incident sebagai major incident untuk eskalasi, koordinasi respons lebih luas, dan queue khusus.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan dan Manajer Layanan/ITSM dengan permission Edit work items; Stakeholder/Viewer menerima dampaknya.
- **Aturan perilaku (acceptance criteria):**
  1. Issue view incident menampilkan toggle "Major incident" untuk mark/unmark.
  2. Queue "Major incidents" menampilkan semua incident dengan flag aktif.
  3. Incident ber-flag major menampilkan badge visual di issue view, list view, dan queue.
  4. Automation menyediakan trigger/condition berdasarkan field Major incident sehingga rule dapat membuat chat channel, memberi tahu stakeholder, atau membuat PIR.
  5. Template automation dapat menandai major secara otomatis berdasarkan kriteria (mis. Priority = Highest dan Impact = Extensive/Widespread).
  6. Saat major incident di-resolve, issue view menampilkan saran "Buat Post-incident review" dengan tautan ke aksi pembuatan PIR (JSM-ITSM-032).
  7. Setiap mark/unmark tercatat di history dan timeline incident dengan pelaku dan waktu.
  8. Hanya user dengan permission Edit work items pada incident yang dapat mengubah flag; user lain melihat toggle dalam keadaan nonaktif.
  9. Unmark major incident tidak menghapus chat channel, conference call, atau PIR yang sudah dibuat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise (feature flag). Diberi P1 karena praktik standar ITSM.

### JSM-ITSM-019 Incident Responders

- **Deskripsi:** Menambah, melihat, dan menghapus responder (user atau team) pada incident sehingga mereka dipanggil melalui alert dan on-call Operations.
- **Prioritas / Fase / Tier:** P0 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder dan on-call engineer (Developer/Anggota Tim), team admin Operations.
- **Aturan perilaku (acceptance criteria):**
  1. Field/panel Responders memungkinkan menambah banyak user dan team sekaligus dalam satu aksi.
  2. Menambah team membuat alert yang diteruskan ke on-call schedule dan escalation policy team tersebut.
  3. Menambah user mengirim notifikasi sesuai notification rules user tersebut (push, SMS, voice, email).
  4. Panel menampilkan status acknowledge per responder berdasarkan alert terkait (Not acknowledged, Acknowledged, Closed).
  5. Responder dapat dihapus dari incident; penghapusan tidak menutup alert yang sudah dibuat.
  6. Responder otomatis ditambahkan dari peran Responders pada affected service (JSM-ITSM-006).
  7. Automation menyediakan action untuk menambah responder dan membuat alert.
  8. Hanya team yang terdaftar sebagai operations team dengan on-call yang menerima alert; team lain ditolak dengan pesan "Team belum dikonfigurasi untuk Operations".
  9. ⚠ Team tanpa on-call schedule tidak menerima alert; sistem menampilkan peringatan di panel Responders dan tidak melakukan fallback ke anggota team.
  10. Alert yang dibuat saat responder ditambahkan memakai message dari summary incident dan alias berbasis key incident, sehingga penambahan responder berikutnya pada incident yang sama tidak menduplikasi alert.
  11. Status acknowledge di panel diperbarui otomatis setiap kali state alert terkait berubah, tanpa perlu memuat ulang halaman.
  12. ⚠ Saran responder berbasis AI tersedia pada paket Premium (rincian di JSM-OPS-042, [34](34-jsm-operations-alerting-on-call.md)).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Versi kanonis fitur ini; inventaris juga mencatat "Incident Responders" di domain Operations (JSM-OPS-040, [34](34-jsm-operations-alerting-on-call.md)) dengan nama yang sama. Pembagian cakupan: modul 32 memiliki field Responders, panel dan UI di incident, aksi automation, serta aturan penambahan/penghapusan; modul 34 hanya memiliki mekanisme notifikasi on-call (routing rule, escalation policy, notification rules, state alert). Prioritas kanonis P0 (mengikuti inventaris domain ITSM); JSM-OPS-040 diselaraskan ke prioritas ini. Responders dan linked alerts tersedia di semua paket. Conference call dirinci di JSM-ITSM-024.

### JSM-ITSM-020 Operations Integration & On-call Link

- **Deskripsi:** Titik integrasi proses incident dengan fitur Operations (alerts, on-call, escalations) dari issue view dan pengaturan service project.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, on-call engineer, Admin Proyek/Project Lead (pengaturan Operations).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings service project memiliki menu Operations yang memuat pengaturan incident management, change management, dan Incident Command Center.
  2. Panel Responders menampilkan tautan "Who's on-call" dari setiap team responder ke on-call schedule team tersebut.
  3. Panel Responders menampilkan tautan ke escalation policy team.
  4. Alert ke responder dapat dikirim otomatis saat incident ditandai major atau berprioritas tinggi, melalui automation atau pengaturan incident management.
  5. ⚠ Priority incident dan alert tertaut tersinkron dua arah secara default (perubahan pada salah satu memperbarui yang lain), dan dapat diubah menjadi satu arah di pengaturan.
  6. Bila fitur Operations dinonaktifkan untuk site atau proyek, panel Alerts dan tautan on-call disembunyikan tanpa menghapus data alert yang tersimpan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Seluruh mekanik Operations di modul 34.

### JSM-ITSM-021 Linked Alerts on Incidents

- **Deskripsi:** Keterhubungan alert Operations dengan incident sehingga agent dapat melihat, membuat, dan menautkan alert dari issue view.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, on-call engineer, responder.
- **Aturan perilaku (acceptance criteria):**
  1. Panel Alerts menampilkan alert terkait dengan status, priority, dan acknowledged by.
  2. Alert dapat dieskalasi menjadi incident; incident baru otomatis tertaut ke alert tersebut.
  3. Alert yang sudah ada dapat ditautkan ke incident dan tautannya dapat dilepas.
  4. Agent dapat membuat alert baru dari incident untuk memanggil on-call.
  5. Penutupan alert saat incident resolved dapat dikonfigurasi lewat automation atau pengaturan integrasi; default-nya tidak menutup otomatis.
  6. Integrasi monitoring (mis. Datadog, Prometheus, CloudWatch) dapat membuat incident otomatis melalui alert atau incident rule.
  7. ⚠ Banyak alert dapat ditautkan ke satu incident; satu alert hanya dapat ditautkan ke satu incident pada satu waktu, dan upaya menautkan alert yang sudah tertaut meminta konfirmasi pemindahan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dedup, grouping, dan integrasi monitoring di modul 34.

### JSM-ITSM-022 Incident Timeline

- **Deskripsi:** Catatan kronologis incident (status, alert, aktivitas ICC, catatan internal) yang menjadi rekam jejak dan dasar PIR.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder, Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Timeline otomatis memuat perubahan status, priority, alert terkait, penambahan/penghapusan responder, dan aktivitas ICC dengan stempel waktu.
  2. Agent dapat menambah internal note untuk waktu sekarang atau waktu lampau (backdated) melalui Activity > Comments; entri backdated ditempatkan sesuai waktu yang dipilih dan ditandai "dicatat pada <waktu aktual>".
  3. Komentar internal dan balasan ke customer tampil di activity dengan pembeda visual.
  4. Activity menyediakan filter All, Comments, History, Work log, Approvals.
  5. Setelah incident resolved, aplikasi chat yang terhubung mengirim prompt ke assignee di Slack untuk memilih pesan/update penting dari channel incident, lalu mengirimnya ke timeline incident.
  6. Bila timeline dari Slack sudah pernah diposting ke incident, sistem menampilkan peringatan sebelum memposting ulang untuk mencegah duplikat.
  7. Timeline dapat dipakai sebagai input PIR, baik disalin manual maupun dipakai oleh AI draft (JSM-ITSM-032).
  8. Entri timeline yang dimasukkan ke PIR wajib diurutkan kronologis berdasarkan waktu kejadian (bukan waktu pencatatan); tes regresi memverifikasi urutan ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Activity tabs CMT-021 ([08](08-komentar-lampiran-aktivitas.md)); Slack di [25](25-integrasi-impor-ekspor-ekstensi.md). Inventaris mencatat bug urutan timeline di produk acuan; produk ini menjadikannya aturan eksplisit.

## Incident: kolaborasi & komunikasi

### JSM-ITSM-023 Incident Chat Channels (Slack & Microsoft Teams)

- **Deskripsi:** Membuat chat channel khusus incident (war room) di Slack atau Microsoft Teams langsung dari incident untuk koordinasi respons.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder, Manajer Layanan/ITSM; admin workspace Slack/Teams (otorisasi app).
- **Aturan perilaku (acceptance criteria):**
  1. Issue view incident menampilkan aksi "Create channel" bila aplikasi chat Slack atau Microsoft Teams produk ini terhubung ke site; bila belum, aksi menampilkan petunjuk koneksi.
  2. Nama channel dibuat otomatis dari key dan summary incident (mis. `inc-123-database-down`), dapat diubah sebelum dibuat, dan divalidasi terhadap aturan nama platform chat.
  3. Semua responders yang memiliki akun chat terhubung diundang ke channel saat dibuat; responder yang ditambah kemudian ikut diundang.
  4. Channel menerima pesan pembuka berisi ringkasan, priority, status, dan tautan incident.
  5. Perubahan status dan priority incident diposting otomatis ke channel.
  6. ⚠ Dari channel, pengguna dapat mengubah priority/status, menambah responder, dan acknowledge alert lewat tombol atau perintah; setiap aksi mengikuti permission akun produk ini yang terhubung.
  7. Untuk Microsoft Teams, aksi tambahan "Start chat" dan "Schedule meeting" tersedia dari incident.
  8. Perintah "Summarize incident" di Slack menghasilkan ringkasan AI dari incident dan percakapan channel bila AI aktif.
  9. Tautan channel tampil di panel incident setelah dibuat.
  10. Automation menyediakan action "Create chat channel" agar channel dibuat otomatis saat incident ditandai major.
  11. Pembuatan channel gagal dengan pesan yang jelas bila admin workspace Slack/Teams belum mengotorisasi app.
  12. Channel yang diarsipkan di platform chat tidak menghapus tautan di incident; tautan ditandai "Diarsipkan".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise (chat tools). Integrasi Slack/Teams di [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-030 s.d. INT-036).

### JSM-ITSM-024 Incident Conference Calls

- **Deskripsi:** Memulai conference call bawaan atau meeting Zoom/Microsoft Teams dari incident untuk koordinasi responder.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder, Admin Situs (menghubungkan Zoom), Admin Proyek/Project Lead (mengaktifkan per proyek).
- **Aturan perilaku (acceptance criteria):**
  1. Field Conference call pada issue view menyediakan aksi untuk memulai call baru atau menautkan URL call yang sudah ada.
  2. Conference call dapat dimulai langsung dari queue Major incidents.
  3. Integrasi Zoom memerlukan dua langkah: Admin Situs menghubungkan site ke akun Zoom, lalu admin proyek mengaktifkan Zoom untuk proyek; tanpa kedua langkah opsi Zoom tidak tampil.
  4. Meeting Microsoft Teams dapat dibuat dari incident melalui aplikasi Teams yang terhubung.
  5. Tautan meeting dibagikan ke responder (notifikasi) dan ke chat channel incident bila ada.
  6. Menghapus meeting Zoom dari incident hanya melepas tautan dan tidak mengakhiri meeting yang berlangsung.
  7. Bila integrasi belum dikonfigurasi, aksi menampilkan prompt untuk menghubungi admin beserta tautan ke halaman pengaturan (untuk admin).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Conference call bawaan hanya Premium/Enterprise. Google Meet tidak terkonfirmasi native; dapat dipakai lewat tautan URL manual. Layanan video conference bawaan merupakan keputusan terbuka (bangun sendiri atau hanya integrasi).

### JSM-ITSM-025 Incident Stakeholders & Stakeholder Communications

- **Deskripsi:** Menambah stakeholder ke incident dan mengirim update progres terstruktur kepada mereka tanpa menjadikan mereka responder.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM, Agent Layanan, Stakeholder/Viewer, Admin Situs (mengelola daftar stakeholder).
- **Aturan perilaku (acceptance criteria):**
  1. Field Stakeholders pada incident menerima individu dan stakeholder group.
  2. Stakeholder otomatis ditambahkan dari peran Stakeholders pada affected service (JSM-ITSM-006).
  3. Automation menyediakan action untuk menambah stakeholder; stakeholder juga dapat ditambah manual dari issue view.
  4. Aksi "Send update" di issue view mengirim update/notifikasi ke semua stakeholder incident melalui email dan notifikasi in-app.
  5. Update dapat dibuat dari incident email templates yang dapat dikelola admin, dengan placeholder status, impact, dan ETA.
  6. Formulir update memiliki opsi untuk memposting isi yang sama ke status page native (JSM-ITSM-026).
  7. Riwayat update stakeholder (isi, pengirim, waktu, penerima) tersimpan dan tampil di incident.
  8. Siapa pun yang memiliki akses Browse ke incident dapat membuat dan mengirim update.
  9. Penambahan, perubahan, dan penghapusan entitas stakeholder (daftar dan stakeholder group tingkat site) hanya dapat dilakukan Admin Situs dan admin aplikasi.
  10. Sistem menolak penambahan stakeholder baru dengan pesan galat bila jumlah stakeholder site mencapai 5.000.
  11. Stakeholder tidak memerlukan lisensi agent.
  12. Akun yang memiliki data stakeholder warisan dari sistem alerting lama yang belum digabung menampilkan stakeholder tersebut sebagai hanya-baca sampai migrasi data selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise. Email template di [19](19-notifikasi-email.md) (NTF-042). Butir 12 merujuk perilaku akun Opsgenie yang belum di-merge di produk acuan; produk ini hanya relevan bila ada impor data dari sistem alerting eksternal (modul 34).

### JSM-ITSM-026 Status Pages (native in Jira Service Management)

- **Deskripsi:** Status page bawaan Service Management untuk mempublikasikan status layanan dan update incident ke publik atau ke grup privat.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead (membuat dan mengonfigurasi), Manajer Layanan/ITSM dan Agent Layanan (memposting update), Tamu/Kolaborator Eksternal sebagai pengunjung/subscriber, Stakeholder/Viewer.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Status pages dapat diakses dari navigasi Service Management/service project.
  2. Admin dapat membuat dan mengustomisasi status page: nama, branding (logo, warna), dan daftar komponen/layanan yang ditampilkan (dipilih dari services).
  3. Status page dapat di-publish, di-unpublish, dihapus, dan dipulihkan (restore) dalam masa retensi; penghapusan meminta konfirmasi.
  4. Pengunjung dapat membuka tampilan publik status page tanpa login (kecuali private), yang menampilkan status komponen dan incident aktif/terbaru.
  5. Pengguna berizin dapat mengedit status page dan mengirim update status komponen atau update incident.
  6. Update stakeholder dari incident dapat diposting langsung ke status page (JSM-ITSM-025).
  7. Admin dapat melihat daftar subscriber, menambah subscriber manual, dan menghapus subscriber.
  8. Pengunjung dapat berlangganan status page publik via email atau webhook; payload webhook memuat komponen, status, dan isi update.
  9. ⚠ Status page dapat memakai custom domain setelah verifikasi DNS (CNAME) dan sertifikat TLS diterbitkan otomatis.
  10. ⚠ Status page private hanya dapat diakses oleh anggota group yang dipilih setelah login.
  11. Setiap perubahan status page (konfigurasi, publish, update) tercatat di audit log.
  12. ⚠ Status page yang di-unpublish mengembalikan HTTP 404 untuk pengunjung publik, sementara data dan subscriber tetap tersimpan dan tidak dinotifikasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Berbeda dari integrasi produk status page terpisah (JSM-ITSM-027). Batas per paket (jumlah page, private page, custom domain) adalah keputusan terbuka. Audit log di [22](22-administrasi-sistem.md) (ADM-014); webhook di [23](23-api-publik-webhook.md). Diberi P2 karena bukan inti agar produk dapat dipakai.

### JSM-ITSM-027 Statuspage Integration for Incidents

- **Alias inventaris:** Statuspage integration
- **Deskripsi:** Integrasi dengan produk status page eksternal (padanan Atlassian Statuspage) untuk mempublikasikan dan memperbarui status layanan dari incident.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Manajer Layanan/ITSM (communications lead), subscriber eksternal; memerlukan akun dan permission di produk status page eksternal.
- **Aturan perilaku (acceptance criteria):**
  1. Admin dapat menghubungkan satu atau lebih halaman status page eksternal ke service project dengan API key.
  2. Dari incident, agent dapat membuat incident status page dengan judul, status (Investigating, Identified, Monitoring, Resolved), dan pesan.
  3. Agent memilih komponen terdampak dan status komponennya (Degraded performance, Partial outage, Major outage, Under maintenance).
  4. Agent dapat mengirim update lanjutan ke incident status page dari incident.
  5. Setiap update memiliki opsi untuk memberi tahu atau tidak memberi tahu subscriber.
  6. Agent dapat memilih template incident dari status page eksternal untuk mengisi judul dan pesan.
  7. Tautan incident status page ditampilkan di panel incident.
  8. Aksi integrasi tidak tersedia dan menampilkan pesan bila langganan atau permission di produk status page eksternal tidak memadai.
  9. Menghapus incident di status page eksternal tidak menghapus incident di produk ini; tautannya ditandai rusak.
  10. Penautan dapat dilakukan tingkat site sehingga beberapa service project memakai halaman yang sama.
  11. Admin dapat memetakan affected services ke komponen status page sehingga komponen terdampak terpilih otomatis.
  12. Status komponen yang tersedia mencakup degraded performance, partial outage, dan major outage; pilihan default mengikuti priority incident.
  13. Agent dapat memakai incident templates status page saat membuat incident.
  14. Saat incident di-resolve, sistem menawarkan untuk me-resolve incident status page dan mengembalikan komponen ke Operational.
  15. Produk status page eksternal tetap produk terpisah dengan subscriber sendiri; produk ini tidak menyimpan atau mengelola subscriber-nya.
  16. Membuat dan memperbarui incident status page dari incident dilakukan dari satu panel yang sama; kegagalan API eksternal menampilkan galat dan tidak mengubah incident di produk ini.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Statuspage integration" dan domain proses ITSM). Produk Atlassian Statuspage berada di luar cakupan; integrasi hanya dibangun bila diperlukan dan dapat digantikan oleh status page native (JSM-ITSM-026).

## Incident: investigasi, AI & post-incident review

### JSM-ITSM-028 Incident Investigation (Recent Deployments, Commits & Feature Flags)

- **Deskripsi:** Tampilan investigasi pada incident untuk melihat deployment, commit, dan perubahan feature flag terbaru yang terkait dengan affected services, sehingga responder dapat menemukan perubahan yang dicurigai sebagai penyebab.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder, Manajer Layanan/ITSM; membutuhkan permission View Development Tools pada proyek software terkait.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Agent issue view untuk work type Incident memiliki tab "Investigate" yang hanya tampil bila prasyarat butir 10 terpenuhi.
  2. Tab menampilkan daftar deployment terbaru dari tool CI/CD terhubung (Bitbucket Pipelines, GitHub, GitLab, Jenkins, dll.) beserta waktu relatif terhadap waktu incident dibuat (mis. "2 jam sebelum incident").
  3. Tab menampilkan commit dan pull request terkait dari repository/proyek software yang dipetakan ke affected services.
  4. ⚠ Tab menampilkan perubahan feature flag dari provider feature flag terhubung dalam rentang waktu incident.
  5. ⚠ Daftar dapat difilter berdasarkan affected service, environment, dan rentang waktu (default 24 jam sebelum incident dibuat).
  6. ⚠ Responder dapat menandai deployment/commit sebagai "penyebab yang dicurigai"; penandaan tercatat di timeline incident dan dapat dibatalkan.
  7. Setiap deployment yang memiliki change request (dari deployment tracking) menampilkan tautan ke change tersebut untuk korelasi incident-change.
  8. ⚠ Setiap deployment, commit, dan pull request memiliki tautan untuk dibuka di tool sumber; tidak ada aksi rollback native, rollback dilakukan di tool CI/CD.
  9. Empty state menjelaskan langkah yang diperlukan bila belum ada development tool terhubung atau affected services belum diisi.
  10. Tab hanya muncul bila service project memiliki integrasi development tools dan services yang dipetakan ke repository, pipeline, atau proyek software.
  11. Saat Incident Command Center aktif, alat investigasi dipindah ke tab Investigation di ICC dan tab di issue view disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada JSM-ITSM-046 dan JSM-ITSM-005; data development di [24](24-integrasi-devops.md) (DEV-001, DEV-013, DEV-021). Ketersediaan per paket kemungkinan Premium/Enterprise (keputusan terbuka). Berbeda dari JSM-ITSM-030 yang berbasis AI.

### JSM-ITSM-029 Incident Command Center (ICC)

- **Deskripsi:** Pengalaman incident terpusat berbasis AI yang menyatukan konteks, investigasi, mitigasi, komunikasi, dan PIR dalam satu tampilan.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM, responder, Admin Proyek/Project Lead (aktivasi), Stakeholder/Viewer (menerima komunikasi). Permission ICC: Create/Update ICC Room, Delete ICC Room, Join ICC Session.
- **Aturan perilaku (acceptance criteria):**
  1. Admin service project mengaktifkan ICC melalui Project settings > Operations > Incident management dengan toggle ICC.
  2. ⚠ Saat ICC aktif, incident baru terbuka di tampilan ICC dengan tab Overview, Investigation, Mitigation, Communications, dan PIR.
  3. Bagian "Get up to speed" menampilkan ringkasan kondisi incident (status, dampak, responder, aktivitas terakhir) yang diperbarui saat incident berubah.
  4. Tab Investigation memuat alat investigasi dari JSM-ITSM-028 secara terintegrasi.
  5. Tab Mitigation menampilkan rekomendasi mitigasi dan memungkinkan menjalankan langkah mitigasi yang memerlukan autentikasi ke GitHub; aksi gagal autentikasi menampilkan prompt login.
  6. ICC menampilkan data dari telemetry connector (metrik, log, sinyal kesehatan dari tool monitoring) dan aktivitas source code terbaru.
  7. Tab Communications memungkinkan membuat update komunikasi ke stakeholder (JSM-ITSM-025).
  8. Tab PIR membuat draft PIR berbasis AI yang wajib ditinjau dan disunting sebelum disimpan sebagai work item PIR.
  9. Permission ICC (Create/Update ICC Room, Delete ICC Room, Join ICC Session) diatur di pengaturan Operations; user tanpa permission tidak dapat membuka atau mengubah room.
  10. ⚠ Incident yang dibuat sebelum ICC diaktifkan tetap memakai tampilan issue view lama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise; status rollout di produk acuan masih berubah. ICC lama berbasis Teams (dihentikan 2024) tidak dibangun. AI di modul 35 (`35-jsm-ai-otomasi-api-admin-csm.md`); konsep incident commander lihat JSM-ITSM-016 butir 26.

### JSM-ITSM-030 AI Assistance for Incidents

- **Deskripsi:** Bantuan AI untuk memahami dan menyelidiki incident (resource terkait, ringkasan, investigasi, perintah chat).
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, responder, Manajer Layanan/ITSM; memerlukan fitur AI aktif di site.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Find related resources" menampilkan incident, alert, halaman pengetahuan, dan work item terkait yang hanya dapat diakses sesuai permission pengguna.
  2. Aksi ringkasan menghasilkan ringkasan incident dan aktivitasnya dalam bahasa pengguna.
  3. Aksi "Investigate with AI" menganalisis perubahan/deployment terkait dan menampilkan kandidat penyebab beserta sumbernya.
  4. Asisten chat AI menerima perintah pengelolaan incident (mis. "buat PIR untuk incident ini") dan meminta konfirmasi sebelum menjalankan aksi yang mengubah data.
  5. Semua aksi AI disembunyikan bila fitur AI dinonaktifkan oleh admin atau paket tidak mendukung.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail AI, model, dan kebijakan data di modul 35. Layanan AI Atlassian (Atlassian Intelligence/Rovo) di luar cakupan; produk ini memakai penyedia AI yang dapat dikonfigurasi.

### JSM-ITSM-031 Incident Prevention Center (Service Collection)

- **Deskripsi:** Pusat pencegahan incident yang menyatukan sinyal risiko dari changes, incidents, dan services untuk mencegah incident berulang.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (incident manager, change manager), service owner.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Halaman menampilkan risiko per service berdasarkan jumlah incident terbaru, change terjadwal, dan change yang gagal dalam rentang waktu yang dapat dipilih.
  2. ⚠ Halaman menampilkan rekomendasi tindakan pencegahan berbasis AI (mis. buat problem untuk incident berulang) yang dapat diterima atau diabaikan.
  3. Setiap sinyal risiko menautkan ke change dan incident terkait.
  4. ⚠ Halaman hanya tersedia bila paket Premium/Enterprise dan AI aktif; selain itu tidak tampil di navigasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Kapabilitas di produk acuan belum terdokumentasi rinci; spesifikasi ini adalah usulan minimum dan perlu keputusan produk. Packaging "Service Collection" Atlassian tidak dibangun.

### JSM-ITSM-032 Post-incident Review (PIR)

- **Alias inventaris:** Post-incident reviews (PIR)
- **Deskripsi:** Work type Post-incident review untuk menganalisis incident setelah selesai, tertaut ke incident, dan dapat diekspor/diterbitkan ke dokumen eksternal.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (incident manager), Agent Layanan, tim engineering (Developer/Anggota Tim), Stakeholder/Viewer (pembaca).
- **Aturan perilaku (acceptance criteria):**
  1. Work category "Post-incident reviews" tersedia dengan work type Post-incident review.
  2. Aksi "Create PIR" dari incident membuka create screen dengan work type PIR terpilih dan incident sumber otomatis ditautkan sebagai primary incident.
  3. ⚠ Link type khusus PIR-incident bernama "reviews / is reviewed by" dibuat oleh template dan tidak dapat dihapus selama work type PIR ada.
  4. Satu PIR dapat ditautkan ke beberapa incident.
  5. ⚠ PIR memiliki field bawaan: deskripsi, timeline, root cause, dampak, contributing factors, dan action items.
  6. Action items dibuat sebagai work item tertaut, termasuk di proyek software, dengan link type "relates to".
  7. Aksi "Suggest description" menyusun draft dari incident dan channel Slack tertaut bila AI aktif (Premium/Enterprise).
  8. Aksi "Create and publish" mengekspor PIR ke XML, Word, PDF, atau ke halaman di wiki/dokumentasi yang terhubung; opsi tujuan dokumen hanya tampil bila integrasi terhubung.
  9. PIR memakai workflow PIR bawaan dan memiliki list view serta queue sendiri.
  10. Template automation membuat PIR otomatis saat major incident di-resolve.
  11. Draft AI selalu ditandai "Draft AI — tinjau sebelum dipublikasikan" dan tidak dapat dipublikasikan tanpa disunting atau dikonfirmasi pengguna.
  12. PIR tersedia sebagai work category tersendiri yang dapat dinonaktifkan per proyek; saat dinonaktifkan, entri sidebar Post-incident reviews disembunyikan tanpa menghapus work item PIR.
  13. Membuat PIR dari incident menautkan keduanya secara otomatis tanpa langkah manual.
  14. ⚠ Field PIR bawaan dapat ditambah dengan custom field melalui screen work type PIR.
  15. Template PIR tersedia untuk dokumen di wiki/dokumentasi yang terhubung, dengan placeholder dari field PIR.
  16. Action items ditampilkan di panel linked work items PIR beserta status dan assignee.
  17. AI dapat menyusun draft PIR dari timeline incident (JSM-ITSM-022).
  18. Template menyediakan queue PIR dan laporan jumlah PIR per periode serta action items yang belum selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Post-incident reviews (PIR)" dan domain proses ITSM). Paket Premium/Enterprise; diberi P1 karena praktik standar ITSM. Integrasi dokumen di [25](25-integrasi-impor-ekspor-ekstensi.md) (INT-042, INT-043) dan [08](08-komentar-lampiran-aktivitas.md) (CMT-031); link type di [05](05-konfigurasi-work-item-skema.md) (CFG-056).

## Problem management & eskalasi ke development

### JSM-ITSM-033 Problem Management

- **Alias inventaris:** Problem management
- **Deskripsi:** Work type Problem untuk investigasi akar penyebab incident berulang, dokumentasi workaround, dan pengelolaan known error.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (problem manager), Agent Layanan, tim engineering, knowledge manager, Customer/Pelapor (membaca artikel known error di portal).
- **Aturan perilaku (acceptance criteria):**
  1. Template menyediakan request type "Investigate a problem" yang secara default agent-only (tersembunyi dari portal).
  2. Work category Problems memiliki list view dan queue sendiri, aktif hanya pada paket Premium/Enterprise.
  3. Work type Problem memiliki field Root cause dan Workaround (rich text).
  4. Work type Problem memiliki field Impact, Urgency, Priority, dan Affected services.
  5. ⚠ Workflow Problem management bawaan memiliki status Open, Under investigation, Under review, Completed, dan Canceled dengan transisi maju dan Cancel dari setiap status aktif.
  6. Problem dapat dibuat dari satu incident atau dari kumpulan incident yang dipilih di list view (bulk "Create problem"); semua incident terpilih otomatis tertaut.
  7. Known error dimodelkan sebagai problem yang memiliki Root cause dan Workaround terisi; list view menyediakan filter "Known errors" berdasarkan kondisi tersebut, tanpa modul atau work type terpisah.
  8. Workaround/known error dapat dipublikasikan sebagai artikel knowledge base agar agent dan customer dapat menemukannya.
  9. Aksi "Create article" dari problem membuat artikel KB dari template known error atau troubleshooting; artikel tertaut balik ke problem di panel knowledge base.
  10. Visibilitas artikel known error dapat dibatasi untuk agent internal saja atau dibagikan ke portal customer, mengikuti permission knowledge base proyek.
  11. Saat menangani incident terkait, agent dapat mencari dan membagikan artikel known error dari panel knowledge base di issue view.
  12. Request type/forms dan template dokumen (mis. RCA) dapat dipakai untuk menstandarkan pencatatan problem.
  13. Template automation memberi tahu (komentar dan notifikasi) seluruh incident terkait saat field Workaround diisi.
  14. Problem proaktif tanpa incident terkait dapat dibuat dan diproses tanpa peringatan.
  15. Artikel KB tetap ada dan tetap dapat diakses setelah problem ditutup.
  16. ⚠ Work category Problems memakai workflow Problem management bawaan (JSM-ITSM-002); work item Problem baru selalu dimulai di status awal workflow tersebut (Open pada workflow bawaan).
  17. Banyak incident dapat ditautkan ke satu problem, dan problem menampilkan jumlah incident terkait.
  18. ⚠ Field Root cause dan Workaround merupakan field bawaan template dan tidak perlu dibuat manual.
  19. Problem dapat dieskalasi ke tim development melalui linked work item (JSM-ITSM-035).
  20. Template menyediakan queue "Open problems".
  21. Fitur AI menampilkan incident dan problem serupa pada problem bila AI aktif (JSM-ITSM-030).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM dan domain proses ITSM, nama sama). Paket Premium/Enterprise. Tidak ada work type "Known error" atau modul KEDB native. Knowledge base di modul 31 (`31-jsm-antrian-sla-knowledge-base-laporan.md`).

### JSM-ITSM-034 Linking Incidents, Problems & Changes

- **Deskripsi:** Penautan incident dengan problem dan change untuk analisis akar masalah dan pelacakan perbaikan.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Manajer Layanan/ITSM (problem manager, change manager); membutuhkan permission Link work items.
- **Aturan perilaku (acceptance criteria):**
  1. Issue view incident menyediakan aksi "Create problem" yang membuat problem baru dan otomatis menautkannya ke incident.
  2. Issue view incident menyediakan aksi "Link problem" untuk menautkan ke problem yang sudah ada melalui pencarian.
  3. Issue view incident dan problem menyediakan aksi "Create change" yang membuat change tertaut.
  4. ⚠ Template membuat link type bawaan "Problem/Incident" dengan arah "causes / is caused by".
  5. Panel linked work items mengelompokkan tautan per work category (Incidents, Problems, Changes, lainnya).
  6. Problem menampilkan seluruh incident terkait dalam satu daftar dengan status masing-masing.
  7. Change penyebab incident dapat ditautkan (change-induced incident) dan tautan ini ikut tampil di Risk insights change tersebut (JSM-ITSM-043).
  8. Template automation: saat problem mencapai Completed, semua incident terkait menerima komentar dan opsional ditransisikan.
  9. Tautan lintas service project dan proyek software diperbolehkan; alur eskalasi ke development dirinci di JSM-ITSM-035.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Issue links di [06](06-work-item-inti.md) (ISS-029) dan [05](05-konfigurasi-work-item-skema.md) (CFG-056); permission di [03](03-izin-peran-keamanan.md) (PERM-024).

### JSM-ITSM-035 Developer Escalations

- **Deskripsi:** Agent mengeskalasi request, incident, atau problem ke tim development dengan membuat atau menautkan work item di proyek software dari issue view, lalu memantau progresnya tanpa keluar dari service project.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan, Developer/Anggota Tim, Admin Proyek/Project Lead, problem manager; butuh permission Link work items di kedua proyek dan Browse projects di proyek tujuan.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Create linked work item" di agent issue view meminta proyek tujuan, work type, dan link type (mis. causes / is caused by, relates to).
  2. ⚠ Formulir eskalasi menyediakan opsi menyalin summary, deskripsi, dan attachment ke work item development (default: summary dan deskripsi disalin, attachment tidak).
  3. Agent dapat menautkan ke work item development yang sudah ada.
  4. Panel linked work items di request menampilkan status dan assignee work item development secara langsung.
  5. ⚠ Panel development info (commit, branch, PR) dan informasi versi/release dari work item development tertaut tampil di request bagi agent yang memiliki View Development Tools.
  6. ⚠ Template automation tersedia untuk: transisi request saat work item development selesai, komentar ke customer saat bug di-fix, dan sinkronisasi status/komentar.
  7. Reporter dan request participants dapat diberi tahu saat linked development work selesai melalui automation atau customer notification.
  8. Untuk problem management, problem ditautkan ke bug/story di proyek software sementara workaround tetap dicatat di problem.
  9. Aksi eskalasi hanya tersedia bila user memiliki Link work items di kedua proyek dan Browse projects di proyek tujuan.
  10. ⚠ Agent tanpa akses ke proyek software melihat tautan dengan key dan status saja (tanpa summary) dan menerima galat permission bila membukanya.
  11. Bila work item development dihapus, tautan hilang dari panel; bila dipindah ke proyek lain, tautan mengikuti key baru; tautan lintas site tidak didukung sebagai issue link native dan hanya dapat berupa remote link.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Titik integrasi utama Service Management dengan kerja software. Development panel di [24](24-integrasi-devops.md) (DEV-013); remote links di [06](06-work-item-inti.md) (ISS-030); template di JSM-ITSM-048.

## Change management: work type, tipe, workflow & CAB

### JSM-ITSM-036 Change Work Type & Change Fields

- **Alias inventaris:** Change management
- **Deskripsi:** Work type Change dengan request type "Request a change" dan field ITSM untuk perencanaan, penilaian risiko, dan otorisasi perubahan; heading ini juga merangkum alur change management end-to-end.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Customer/Pelapor atau Agent Layanan sebagai change requester, Manajer Layanan/ITSM (change manager, anggota CAB), approver.
- **Aturan perilaku (acceptance criteria):**
  1. Template menyediakan request type "Request a change" yang dapat dibuat tersedia di portal atau agent-only.
  2. Field Change type memiliki opsi Standard, Normal, Emergency.
  3. Field Change risk memiliki opsi Low, Medium, High, Critical, diisi manual atau oleh risk assessment (JSM-ITSM-042).
  4. Field Change reason memiliki opsi Repair, Upgrade, Maintenance, New functionality, Other.
  5. Work type Change memiliki field Impact dan Urgency.
  6. Field Change start date dan Change completion date (date-time) menentukan posisi change di change calendar; completion date sebelum start date ditolak dengan pesan galat.
  7. ⚠ Template menyediakan field Implementation plan, Backout plan, dan Test plan (rich text) yang dapat dihapus atau diganti custom field/form.
  8. Request type dapat memuat form penilaian risiko (JSM Forms) sebagai input manual risiko.
  9. Work type Change memiliki field Affected services.
  10. Work type Change memiliki field Approvers dan CAB.
  11. Issue view change menampilkan informasi deployment pada panel Development/Deployments.
  12. Halaman Changes menyediakan list view, filter, dan queue bawaan (Awaiting approval, Scheduled, In progress).
  13. Field wajib dapat ditegakkan lewat request type (required field) atau workflow validator; pelanggaran menampilkan nama field yang belum diisi.
  14. Change dengan Change type = Emergency dapat melewati sebagian langkah perencanaan (validator rencana tidak dijalankan) sesuai konfigurasi workflow.
  15. Work category Change memakai workflow change ITSM (JSM-ITSM-038).
  16. Tipe change Standard, Normal, dan Emergency menentukan jalur approval (JSM-ITSM-037).
  17. ⚠ Field bawaan change mencakup Change type, Change risk, Change reason, Impact, Urgency, Affected services, Planned start/end, serta rencana implementasi, backout, dan test.
  18. Approval CAB berjalan via field approvers/grup pada status approval (JSM-ITSM-039).
  19. Standard change dapat disetujui otomatis dengan automation (JSM-ITSM-015).
  20. SLA "Time to approve" dan "Time to review" change dapat dikonfigurasi.
  21. Field Affected CI menerima object Assets (configuration item) selain services.
  22. Change dapat ditautkan ke incident, problem, dan deployment.
  23. Template menyediakan queue Changes.
  24. ⚠ Change freeze/blackout ditegakkan melalui change windows (JSM-ITSM-041) dan, bila perlu blokir keras, melalui automation atau workflow validator yang menolak transisi ke Implementing selama freeze window.
- **Team-managed vs company-managed:** Tidak berlaku (JSM-PRJ-012); hanya service project company-managed dari template ITSM.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Change management" dan domain proses ITSM). Paket Premium/Enterprise. Validator di [09](09-workflow-status-transisi.md) (WF-028); Forms di modul 29; object Assets di modul 33.

### JSM-ITSM-037 Change Types: Standard, Normal, Emergency

- **Deskripsi:** Penanganan berbeda per tipe change sesuai ITIL, termasuk pre-approval standard change dan jalur cepat emergency.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change manager, CAB), change requester.
- **Aturan perilaku (acceptance criteria):**
  1. Standard change diperlakukan pre-approved dan berisiko rendah; template workflow/automation memindahkannya langsung ke Awaiting implementation tanpa approval manual.
  2. Katalog standard change dibuat sebagai request type terpisah dengan form dan field preset (pengganti objek "change template").
  3. Change sebelumnya dapat dikloning untuk perubahan berulang memakai fitur clone inti.
  4. Normal change wajib melalui risk assessment dan approval manager/CAB sebelum implementasi.
  5. Emergency change memakai jalur approval dipercepat (ECAB) dan mewajibkan dokumentasi pasca-implementasi sebelum Completed.
  6. Workflow mendukung transisi kondisional berdasarkan Change type (condition pada transisi).
  7. Template automation: bila Change type = Standard, change otomatis disetujui dan ditransisikan.
  8. Template automation memicu approval ulang bila Change type diubah setelah approval selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Perilaku dikonfigurasi lewat workflow dan automation, bukan hardcode. Clone ISS-031 ([06](06-work-item-inti.md)); condition WF-027 ([09](09-workflow-status-transisi.md)).

### JSM-ITSM-038 Change Management Workflow

- **Deskripsi:** Workflow bawaan untuk work type Change yang mencakup perencanaan, otorisasi, implementasi, dan review.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM, Agent Layanan, anggota CAB, Admin Situs (mengedit workflow).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Workflow bawaan memiliki status Planning, Awaiting approval, Awaiting implementation, Implementing, Peer review, Completed, Failed, Canceled, dan Declined.
  2. Status Awaiting approval adalah approval step (JSM-ITSM-011).
  3. Keputusan approval menjalankan transisi Approve (ke Awaiting implementation) atau Decline (ke Declined) secara otomatis.
  4. Workflow memiliki transisi Implement (ke Implementing), Mark as failed (ke Failed), dan Complete (ke Completed/Peer review).
  5. Setiap status memiliki nama customer-facing yang tampil di portal.
  6. Validator pada transisi menuju Awaiting approval mewajibkan field rencana (Implementation plan, Backout plan) terisi.
  7. Status change diperbarui otomatis dari event deployment: deployment sukses memindahkan ke Completed, deployment gagal ke Failed (JSM-ITSM-046).
  8. Backout dicatat dengan transisi ke status Failed disertai komentar wajib.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Workflow engine di [09](09-workflow-status-transisi.md); ITSM workflows di JSM-ITSM-002.

### JSM-ITSM-039 Change Advisory Board (CAB)

- **Deskripsi:** Approval oleh CAB untuk change berisiko melalui field CAB/approver dan grup.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (anggota CAB, change manager), change requester.
- **Aturan perilaku (acceptance criteria):**
  1. Work type Change memiliki field CAB (multi user) dan Approver groups (multi group).
  2. Status approval CAB dapat memakai aturan all/any/N approver (JSM-ITSM-011).
  3. Template automation menetapkan anggota CAB otomatis berdasarkan Change risk atau affected services.
  4. Change approvers dari service di Affected services ditambahkan ke approval (JSM-ITSM-006).
  5. Agenda CAB disusun dari queue/filter change yang menunggu approval dan dari change calendar.
  6. Emergency change memakai field atau grup ECAB terpisah dengan aturan jumlah persetujuan yang lebih kecil.
  7. Keputusan dan alasannya dicatat sebagai komentar approval dan tampil di riwayat approval.
  8. Anggota CAB yang juga reporter change dikecualikan dari approval bila opsi "Exclude approvers from: Reporter" aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** CAB bukan objek terpisah; dimodelkan dengan field approver dan grup.

### JSM-ITSM-040 Change Calendar

- **Alias inventaris:** Change calendar
- **Deskripsi:** Kalender visual untuk melihat change yang direncanakan/berjalan beserta change windows guna menghindari bentrok.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM, Agent Layanan, anggota CAB, Stakeholder/Viewer dengan akses Browse ke service project.
- **Aturan perilaku (acceptance criteria):**
  1. Calendar diakses dari halaman Changes > Calendar.
  2. Calendar menyediakan tampilan Month, Week, dan Day dengan navigasi Today, previous, dan next.
  3. Change ditempatkan berdasarkan Change start date dan Change completion date dalam zona waktu pengguna.
  4. Pengaturan "Show change requests on the change calendar" menentukan change mana yang tampil (mis. berdasarkan JQL atau status).
  5. Calendar dapat difilter berdasarkan change type, risk, status, affected services, dan assignee.
  6. Klik event membuka detail change dalam panel samping.
  7. Calendar menampilkan maintenance window dan freeze window (JSM-ITSM-041).
  8. ⚠ Calendar dapat dilanggani dari aplikasi kalender eksternal melalui feed iCal/ICS beralamat rahasia yang memuat change sesuai filter aktif; feed dapat dibuat ulang untuk membatalkan alamat lama.
  9. ⚠ Change dapat ditampilkan di kalender tim pada produk wiki/dokumentasi yang terhubung melalui feed yang sama.
  10. Change tanpa Change start date atau Change completion date tidak tampil di calendar.
  11. Change multi-hari tampil sebagai bar yang membentang sepanjang rentang tanggalnya.
  12. Tampilan daftar (list) tersedia selain Month dan Week, menampilkan change terurut menurut Change start date untuk rentang yang sedang dilihat.
  13. Change diposisikan berdasarkan Planned start/end (Change start/completion date).
  14. Filter mencakup change type, risk, status, dan affected services.
  15. Warna event ditentukan oleh risk atau change type sesuai pilihan pengguna, dengan legenda.
  16. Detail change dapat dibuka dari calendar.
  17. ⚠ Calendar dapat menampilkan deployment dari tool CI/CD terhubung sebagai event terpisah yang dapat dimatikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM dan domain proses ITSM, nama sama). Paket Premium/Enterprise. Komponen calendar view dapat berbagi implementasi dengan PLAN-022 ([18](18-timeline-rencana-tampilan-proyek.md)). Team Calendars di produk dokumentasi Atlassian di luar cakupan.

## Change management: windows, risiko & pengaturan

### JSM-ITSM-041 Change Windows (Maintenance & Freeze Windows)

- **Deskripsi:** Penanda waktu di change calendar untuk kapan change sebaiknya dijadwalkan (maintenance window) atau tidak boleh dijadwalkan (freeze/blackout window).
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change manager), Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Tipe Maintenance window menandai waktu yang direkomendasikan untuk change (di luar jam sibuk, staf siaga, customer sudah diberi tahu).
  2. Tipe Freeze window (deployment blackout window) menandai periode larangan change (libur, event bisnis kritis, setelah change besar).
  3. Pengguna berizin dapat membuat, mengedit, dan menghapus change window dari change calendar dengan atribut nama, tipe, rentang waktu (start < end), dan deskripsi.
  4. Window tampil sebagai area berwarna berbeda per tipe di semua tampilan calendar.
  5. Freeze window selalu dihitung sebagai konflik di Risk insights; maintenance window dapat dihitung atau tidak sesuai pengaturan proyek (JSM-ITSM-043).
  6. Benturan change dengan window menjadi input AI risk assessment (JSM-ITSM-044).
  7. ⚠ Change yang sudah approved lalu tumpang tindih dengan freeze window yang dibuat kemudian tidak diblokir otomatis; change tersebut ditandai konflik di Risk insights dan calendar.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise. Recurrence dan cakupan per service adalah keputusan terbuka.

### JSM-ITSM-042 Automated Change Risk Assessment

- **Alias inventaris:** Change risk assessment engine
- **Deskripsi:** Mesin penilaian risiko otomatis yang menghitung Change risk berdasarkan aturan pada atribut change dan memakainya untuk routing approval.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change manager), Admin Proyek/Project Lead (konfigurasi; permission Administer projects).
- **Aturan perilaku (acceptance criteria):**
  1. Konfigurasi berada di Project settings > Operations > Change management > Risk assessment.
  2. Admin membuat kriteria risiko berbasis field (Change type, tier affected services, Impact, Urgency, jadwal, dan field lain) dengan operator perbandingan.
  3. Hasil penilaian dipetakan ke Low, Medium, High, atau Critical dan disimpan ke field Change risk.
  4. Admin memilih request type yang dikenai penilaian; request type lain tidak dievaluasi.
  5. Hasil penilaian terintegrasi dengan approval: low risk dapat disetujui otomatis, high/critical diarahkan ke CAB.
  6. Penilaian dijalankan ulang otomatis saat field input berubah selama change belum disetujui.
  7. Change risk dapat diubah manual setelah approval; perubahan tercatat di history dan dapat memicu approval ulang melalui automation.
  8. Setiap aturan berbentuk kondisi field yang dapat diuji (mis. "Tier affected service = 1").
  9. Setiap kondisi memiliki skor; skor kondisi yang terpenuhi dijumlahkan lalu dipetakan ke level risiko dengan ambang yang dapat diatur.
  10. Field Change risk diisi otomatis dengan level hasil penilaian dan ditandai "dihitung otomatis".
  11. Level risiko dapat dipakai dalam condition automation untuk routing approval (mis. high risk butuh CAB).
  12. ⚠ Risk assessment hanya tersedia di service project company-managed.
  13. ⚠ Risk assessment hanya tersedia pada paket Premium ke atas.
- **Team-managed vs company-managed:** Hanya company-managed (JSM-PRJ-012).
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM "Change risk assessment engine" dan domain proses ITSM). Mekanisme skor berbobot diusulkan sebagai desain default produk ini, menunggu keputusan pada Pertanyaan terbuka nomor 7.

### JSM-ITSM-043 Risk Insights Panel

- **Deskripsi:** Panel pada change request yang mengumpulkan konteks risiko: konflik jadwal, window, dan incident terbuka.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change manager, CAB), approver.
- **Aturan perilaku (acceptance criteria):**
  1. Tab risk summary menampilkan change lain yang rentang waktunya bersinggungan dan berbagi affected service.
  2. Panel menampilkan freeze dan maintenance window yang bersinggungan dengan rentang change.
  3. Panel menampilkan incident terbuka pada service yang sama selama rentang change.
  4. Admin mengatur per proyek di Project settings > Operations > Change management apa yang dihitung sebagai konflik (change lain, freeze window, maintenance window, incident terbuka).
  5. Approver melihat ringkasan jumlah konflik di panel approval saat memutuskan.
  6. Change tanpa tanggal menampilkan pesan "Isi Change start date dan Change completion date untuk mengevaluasi konflik".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise.

### JSM-ITSM-044 AI Risk Assessment for Changes

- **Deskripsi:** Penilaian risiko change berbasis AI dengan uraian kategori risiko dan saran mitigasi.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (change manager), approver.
- **Aturan perilaku (acceptance criteria):**
  1. AI mengevaluasi risiko operasional change termasuk konflik jadwal dan benturan freeze/maintenance window.
  2. Hasil menampilkan breakdown risiko per kategori teknis dan operasional dengan alasannya.
  3. Hasil menampilkan saran mitigasi yang dapat disalin ke field rencana.
  4. Fitur tersedia pada work item bertipe Change saja.
  5. Fitur hanya tampil pada paket Premium/Enterprise dengan AI aktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hasil AI bersifat saran dan tidak mengisi Change risk otomatis tanpa konfirmasi. AI di modul 35.

### JSM-ITSM-045 Change Management Settings & System Change Automation

- **Deskripsi:** Halaman pengaturan change management per proyek dan rule automation sistem untuk change.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead (permission Administer projects), Manajer Layanan/ITSM.
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Operations > Change management menjadi pusat konfigurasi change management proyek.
  2. Halaman memuat pengaturan koneksi CI/CD dan pemilihan request type untuk change otomatis.
  3. Halaman memuat pengaturan risk assessment dan risk insights.
  4. Sistem menyediakan rule automation tingkat sistem untuk change (mis. tambah change approvers dari affected services, transisi dari event deployment) yang dapat dinyalakan/dimatikan tetapi tidak diedit; hanya pada paket Premium.
  5. Menonaktifkan change management menghentikan pembuatan change otomatis dari deployment; change yang sudah ada tidak berubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama menu mengikuti navigasi settings di [21](21-navigasi-ux-preferensi.md).

## Deployment & CI/CD

### JSM-ITSM-046 Deployment Tracking & Auto-created Change Requests

- **Deskripsi:** Pelacakan deployment dari tool CI/CD yang terhubung dan pembuatan change request otomatis dari deployment.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Developer/Anggota Tim (DevOps engineer), Manajer Layanan/ITSM (change manager), Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Admin dapat menghubungkan tool CI/CD (Bitbucket Pipelines, Jenkins, GitHub Actions) melalui Development Information/Deployments API.
  2. Deployment dikaitkan ke work item melalui issue key di commit, branch, atau PR.
  3. Panel Deployments/Development pada change menampilkan environment, status, dan waktu deployment.
  4. Change request dibuat otomatis untuk deployment ke environment yang dipilih (mis. production).
  5. Admin memilih request type yang dipakai untuk change otomatis.
  6. Summary dan deskripsi change otomatis diisi dari info deployment dan work item terkait.
  7. Status change diperbarui saat deployment sukses atau gagal.
  8. Deployment tracking dipakai bila tidak perlu approval (alternatif dari gating).
  9. Data deployment juga dipakai Incident Investigation untuk korelasi incident-change (JSM-ITSM-028).
  10. Deployment tanpa issue key tidak terhubung ke work item dan tidak membuat change.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise. API di [23](23-api-publik-webhook.md) (API-062, API-063); data model di [24](24-integrasi-devops.md) (DEV-001, DEV-003).

### JSM-ITSM-047 Deployment Gating

- **Alias inventaris:** Deployment gating (CI/CD change management)
- **Deskripsi:** Membuat change request otomatis dari pipeline CI/CD dan menahan deployment ke environment terproteksi sampai change disetujui.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Developer/Anggota Tim (DevOps engineer), Manajer Layanan/ITSM (change manager), approver, Admin Proyek/Project Lead.
- **Aturan perilaku (acceptance criteria):**
  1. Admin proyek mengonfigurasi gating di Project settings > Change management: memilih tool dan environment yang di-gate (mis. production).
  2. Pipeline ditahan (status pending) sampai change disetujui.
  3. Change request dibuat otomatis untuk setiap deployment yang di-gate.
  4. Approve melanjutkan deployment; Decline menggagalkan/membatalkan deployment.
  5. Untuk Bitbucket Pipelines, gating diatur per deployment environment.
  6. Untuk Jenkins, plugin mengirim info deployment dan memeriksa status gating sebelum melanjutkan.
  7. Untuk GitHub, gating berjalan melalui integrasi GitHub (Premium/Enterprise).
  8. Approver menerima notifikasi via email/Slack saat change dari gating dibuat.
  9. Change menampilkan riwayat deployment terkait.
  10. ⚠ Gating memiliki batas waktu tunggu (default 24 jam, dapat diatur); saat timeout pipeline menerima status gagal.
  11. Bila change dibatalkan saat pipeline menunggu, pipeline menerima status gagal dengan alasan "Change dibatalkan".
  12. ⚠ Re-run deployment yang sama memakai change yang sama bila change masih terbuka; bila sudah ditutup, change baru dibuat.
  13. ⚠ Tool yang didukung: Bitbucket Pipelines, GitHub Actions, GitLab, Jenkins, CircleCI, Octopus Deploy, Azure DevOps.
  14. Change management dapat diaktifkan per pipeline/environment.
  15. Admin memilih request type untuk change dari deployment.
  16. Gating: pipeline menunggu keputusan approval lalu lanjut atau batal.
  17. Admin menetapkan perilaku default saat timeout (gagal atau lanjut).
  18. Change memuat info deployment (commit, build, environment).
  19. Request channel change dari pipeline tercatat sebagai "deployment".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Dua entri inventaris digabung (lampiran JSM dan domain proses ITSM). Paket Premium/Enterprise. Sisi DevOps di DEV-027 ([24](24-integrasi-devops.md)).

## Otomasi ITSM, panduan agent & orkestrasi lintas tim

### JSM-ITSM-048 ITSM Automation Templates

- **Deskripsi:** Kumpulan template automation khusus ITSM untuk prioritas, eskalasi, major incident, problem, change, dan stakeholder.
- **Prioritas / Fase / Tier:** P1 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead (mengaktifkan template; permission Administer projects), actor automation.
- **Aturan perilaku (acceptance criteria):**
  1. Template "Priority dari Impact x Urgency" mengisi Priority saat incident dibuat atau saat Impact/Urgency berubah, berdasarkan matriks yang dapat diedit; bila salah satu field kosong, Priority tidak diubah.
  2. Template "Notify on-call" membuat alert ke on-call team responder untuk incident dengan Priority Highest atau High (ambang dapat diubah).
  3. Template "Mark major incident" menyalakan flag Major incident bila kriteria field terpenuhi (default: Priority = Highest dan Impact = Extensive/Widespread).
  4. Template "Create chat channel" membuat chat channel incident (JSM-ITSM-023) saat flag Major incident dinyalakan; bila channel sudah ada, rule tidak membuat channel kedua.
  5. Template "Transisi incident terkait" berjalan saat problem mencapai status Completed: incident tertaut yang belum resolved menerima komentar dan ditransisikan ke Resolved.
  6. Template "Auto-approve standard change" menyetujui otomatis approval step aktif pada change dengan Change type = Standard (perilaku action di JSM-ITSM-015).
  7. Template "Set approvers dari affected services" menambahkan change approvers setiap service di field Affected services ke field Approvers/Approver groups saat change dibuat atau field Affected services berubah, tanpa menghapus approver yang diisi manual.
  8. Template "Tambah stakeholder ke incident" menambahkan individu atau stakeholder group yang ditentukan ke field Stakeholders saat kondisi terpenuhi (mis. major incident atau Priority = Highest); hanya aktif pada paket Premium/Enterprise.
  9. Template "Transisi/komentar request saat linked development work item selesai" menambahkan komentar ke customer dan mentransisikan request saat semua work item development tertaut masuk status category Done (JSM-ITSM-035).
  10. Eksekusi rule template dihitung terhadap batas eksekusi per paket; saat batas tercapai rule berhenti dan admin menerima peringatan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Template library AUTO-010 dan batas AUTO-015 ([20](20-otomasi.md)).

### JSM-ITSM-049 Playbooks (Agent Playbooks)

- **Deskripsi:** Panduan langkah demi langkah untuk agent di work item; tiap langkah berupa instruksi manual atau langkah automated yang menjalankan automation dari issue view, lengkap dengan output dan execution log.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead (membuat playbook), Agent Layanan dan responder (menjalankan).
- **Aturan perilaku (acceptance criteria):**
  1. Project settings > Playbooks menampilkan daftar semua playbook dengan filter (mis. sudah pernah dijalankan).
  2. ⚠ Admin dapat membuat dan mengedit playbook dengan nama dan scope (request type, kondisi, atau JQL) yang menentukan di work item mana playbook muncul.
  3. Setiap langkah bertipe instruksi manual (checklist) atau langkah automated yang menjalankan automation rule.
  4. Playbook berstatus enabled atau inactive; playbook inactive tidak tampil dan tidak dapat dijalankan.
  5. Panel Playbooks di agent issue view memungkinkan menjalankan langkah dan menandai langkah selesai.
  6. Automation menyediakan action untuk mengatur output playbook: pesan kustom yang dilihat agent dan hasil yang disorot.
  7. Output ditampilkan per langkah dan per baris di daftar playbook.
  8. Execution log per langkah mencatat waktu trigger, hasil akhir, aksi yang dijalankan, dan flag kegagalan.
  9. Langkah automated dapat dijalankan ulang (re-run).
  10. Playbook dapat dipakai untuk runbook respons incident dan untuk service request.
  11. Satu playbook berlaku hanya untuk satu proyek.
  12. ⚠ Langkah automated yang gagal ditandai di log tanpa memblokir langkah manual berikutnya.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Paket Premium/Enterprise. Dibangun di atas engine automation [20](20-otomasi.md). Berbeda dari JSM-ITSM-048 (tanpa interaksi agent) dan JSM-ITSM-050 (lintas proyek).

### JSM-ITSM-050 Journeys

- **Alias inventaris:** Journeys
- **Deskripsi:** Orkestrasi proses lifecycle end-to-end (mis. onboarding/offboarding karyawan) yang membuat dan mengurutkan banyak work item lintas service project/tim dari satu trigger sampai selesai.
- **Prioritas / Fase / Tier:** P2 / Fase 7 / jsm
- **Aktor & izin:** Admin Proyek/Project Lead dan Admin Situs (membuat journey type), Agent Layanan, tim HR (Pengguna Tim Bisnis).
- **Aturan perilaku (acceptance criteria):**
  1. Journeys diaktifkan via Project settings > Features (label UI: Space settings) dengan toggle on/off.
  2. Admin membuat journey type dengan nama dan request type/work type journey parent pemicu.
  3. Journey parent adalah work item induk yang memulai dan merangkum journey.
  4. ⚠ Trigger default berjalan saat journey parent dibuat; trigger alternatif dapat dipasang via automation (mis. karyawan baru dari sistem HR melalui connection).
  5. Trigger dapat diberi conditions sehingga journey hanya berjalan bila kondisi terpenuhi.
  6. Setiap langkah membuat child/linked work item di proyek tujuan (IT, Facilities, Finance, dll.) dengan field yang sudah diisi.
  7. Langkah dapat bergantung pada langkah lain: langkah berikutnya ditahan sampai work item terkait mencapai status category atau status tertentu.
  8. Journey berjalan dapat dilacak lewat filter di Queues, queue kustom dengan JQL khusus journey, dan chart journey di halaman Summary.
  9. ⚠ Tersedia visualisasi progres per journey dan laporan per langkah.
  10. ⚠ Journey onboarding dapat memakai AI resolver request untuk menyelesaikan langkah sederhana otomatis.
  11. Journeys hanya tersedia di service project company-managed pada paket Premium/Enterprise.
  12. ⚠ Bila work item langkah dihapus, langkah ditandai "Gagal" dan dependency berikutnya tetap tertahan sampai admin melewatinya; bila dipindah, tautan mengikuti key baru.
  13. Journeys awalnya diaktifkan dari Project settings > Features > Features lab.
  14. Definisi journey memuat rangkaian langkah yang membuat work item di beberapa project (mis. HR, IT, Facilities).
  15. Journey dipicu saat work item/request type tertentu dibuat di service project.
  16. ⚠ Journey dapat dipicu dari data worker baru pada sistem HRIS bila koneksi dikonfigurasi.
  17. Work item turunan untuk tiap tim dibuat otomatis sesuai langkah journey.
  18. Progres keseluruhan journey tampil dalam satu tampilan.
  19. ⚠ Template journey tersedia untuk HR onboarding dan offboarding.
  20. ⚠ Fitur dirilis bertahap di balik feature flag "lab"; dukungan team-managed tidak tersedia (JSM-PRJ-012).
- **Team-managed vs company-managed:** Hanya company-managed (JSM-PRJ-012).
- **Catatan & ketergantungan:** Dua entri inventaris dengan nama sama digabung (lampiran JSM P3 dan domain proses ITSM P2; prioritas memakai P2). Koneksi HRIS (mis. Workday) adalah integrasi pihak ketiga di [25](25-integrasi-impor-ekspor-ekstensi.md). Dapat direlokasi bila ada modul ESM/HR terpisah.

### JSM-ITSM-051 Case Management (HR Cases)

- **Deskripsi:** Manajemen kasus kompleks, sensitif, dan berjangka panjang (mis. investigasi employee relations) yang terpisah dari service request standar, dengan kontrol akses ketat.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Agent Layanan (HR agent, employee relations), Admin Proyek/Project Lead, kolaborator lintas tim (Pengguna Tim Bisnis).
- **Aturan perilaku (acceptance criteria):**
  1. Work type Case tersedia terpisah dari Service request dan tidak tampil di portal secara default.
  2. ⚠ Case memiliki field Allegations, Findings, dan Outcome decision.
  3. Akses case sensitif dibatasi ke pihak berwenang melalui issue security level bawaan; pengguna lain mendapat galat "tidak ditemukan".
  4. Kolaborasi lintas tim (Legal, IT, Facilities, Finance, manager) dilakukan melalui work item yang tertaut ke case asal tanpa memberi akses ke isi case.
  5. Saat agent membuka case, AI menyarankan next steps dan case serupa yang dapat diakses agent bila AI aktif.
  6. Case management tersedia sebagai bagian dari HR service management tanpa produk tambahan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status EAP di produk acuan; bukan praktik ITSM murni dan kandidat relokasi ke modul ESM/HR. Issue security di [03](03-izin-peran-keamanan.md).

## Kepatuhan & penyelarasan praktik

### JSM-ITSM-052 ITIL 4 Alignment (PinkVERIFY)

- **Deskripsi:** Penyelarasan fitur ITSM produk ini dengan praktik ITIL 4 beserta dokumentasi pemetaan praktik ke fitur.
- **Prioritas / Fase / Tier:** P3 / Fase 7 / jsm
- **Aktor & izin:** Manajer Layanan/ITSM (IT manager), auditor (Admin Situs).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Produk ini menargetkan cakupan praktik ITIL 4 (Incident, Problem, Change enablement, Service request, Service configuration, Knowledge, Service level) yang memenuhi kriteria sertifikasi PinkVERIFY; sertifikasi formal adalah keputusan bisnis.
  2. Dokumentasi pemetaan praktik ITIL 4 ke fitur (ID JSM-ITSM) dan panduan konfigurasi template ITSM diterbitkan bersama rilis Fase 7.
  3. Dokumentasi menyatakan bahwa konfigurasi custom (workflow/field yang diubah) dapat menyimpang dari praktik ITIL dan menyediakan daftar periksa konfigurasi minimal.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sertifikasi pihak ketiga berada di luar cakupan pengembangan.

## Pertanyaan terbuka

1. Apakah paket Premium/Enterprise (Changes, Problems, major incident, PIR, chat & conference call) ditegakkan sebagai feature flag per site sejak awal, atau semua fitur ITSM dibuka untuk semua instalasi self-hosted produk ini?
2. (Terjawab oleh JSM-PRJ-012) Service project team-managed tidak dibangun; work category ITSM, risk assessment, dan Journeys hanya untuk company-managed. Yang masih terbuka: apakah impor data team-managed dari sumber eksternal perlu mendukung konversi otomatis workflow request type ke workflow scheme company-managed?
3. Apakah produk ini membangun layanan conference call bawaan sendiri, atau hanya mengintegrasikan Zoom/Microsoft Teams/Google Meet lewat tautan (JSM-ITSM-024)?
4. Apakah status page native (JSM-ITSM-026) dibangun lengkap dengan custom domain dan private page, dan apa batas jumlah status page per paket?
5. Apakah approve via balasan email (JSM-ITSM-013) diaktifkan secara default mengingat risiko spoofing, dan verifikasi apa yang diwajibkan?
6. Apakah freeze window (JSM-ITSM-041) cukup sebagai peringatan, atau perlu opsi blokir keras transisi ke Implementing secara native?
7. Mekanisme risk assessment mana yang dipilih sebagai default: rule builder kondisi-ke-level atau skor berbobot (JSM-ITSM-042)?
8. Penyedia AI mana yang dipakai untuk ICC, AI Assistance, AI risk assessment, dan PIR draft, dan apakah fitur AI dapat dimatikan total per site?
9. Apakah integrasi produk status page eksternal (JSM-ITSM-027) dan katalog komponen developer (JSM-ITSM-010) dibangun sama sekali, atau cukup lewat API/webhook?
10. Apakah Case Management (JSM-ITSM-051) dan Journeys (JSM-ITSM-050) tetap di modul ini atau dipindah ke modul ESM/HR tersendiri?
11. Berapa batas waktu tunggu default deployment gating dan perilaku default saat timeout (gagal atau lanjut) (JSM-ITSM-047)?
12. Apakah produk ini akan mengejar sertifikasi PinkVERIFY formal (JSM-ITSM-052)?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** untuk ITSM lanjutan (mis. checklist onboarding/offboarding, CMDB eksternal, change management tambahan, approval multi-level kompleks) tidak dibangun sebagai fitur native; kebutuhan tersebut dipenuhi lewat REST API dan webhook ([23](23-api-publik-webhook.md)) atau ekstensi ([25](25-integrasi-impor-ekspor-ekstensi.md)).
- **Layanan Atlassian lintas produk** tidak dibangun: Atlassian Statuspage (hanya titik integrasi di JSM-ITSM-027), Compass (JSM-ITSM-010), Confluence dan Team Calendars (hanya ekspor/feed ke wiki terhubung), Atlassian Intelligence/Rovo (diganti penyedia AI yang dapat dikonfigurasi), Atlassian Teamwork Graph (diganti impor dari sumber internal), dan packaging "Service Collection".
- **Opsgenie** sebagai produk terpisah dan Incident Command Center lama berbasis Microsoft Teams (dihentikan 2024) tidak dibangun; kemampuan alerting dan on-call ada di modul 34.
- **Integrasi pihak ketiga** seperti Zoom, Slack, Microsoft Teams, GitHub, GitLab, Jenkins, Bitbucket, CircleCI, Octopus Deploy, Azure DevOps, Datadog, Prometheus, CloudWatch, dan Workday hanya dibangun sebagai konektor; produk eksternal tersebut berada di luar cakupan.
- **Sertifikasi PinkVERIFY** dan audit pihak ketiga lainnya adalah keputusan bisnis, bukan pekerjaan pengembangan.
- **Fitur Data Center-only** tidak ada di jatah modul ini (tidak ada fitur bertier dc-only-or-deprecated); objek "change template" dan modul KEDB terpisah tidak dibangun karena tidak native di produk acuan.
