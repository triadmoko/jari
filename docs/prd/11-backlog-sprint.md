# 11 — Backlog & Sprint (Scrum)

Modul ini mendefinisikan halaman Backlog dan seluruh siklus hidup Sprint di produk ini: tampilan backlog (bagian sprint aktif/future dan bagian Backlog), pengurutan prioritas berbasis Rank global (LexoRank-kompatibel), multi-select dan aksi massal, inline create, menu konteks, estimasi inline, panel Epic dan Version, pemecahan work item (Split), ringkasan beban kerja per assignee, serta operasi sprint (create, edit, goal, reorder, start, scope change, complete, reopen, delete, parallel sprints), semantik field Sprint, penegakan permission Manage Sprints, backlog untuk board Kanban (Kanplan), rilis dari board, dan tautan halaman dokumentasi ke sprint. Konfigurasi board (filter, kolom, card layout, estimation statistic, quick filter) ada di modul 10; definisi field Sprint/Rank/Story Points di modul 05; definisi permission di modul 03; laporan sprint (Sprint report, Burndown, Velocity) di modul 16. Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`), sehingga entitas `sprints`, relasi work item ke sprint (termasuk riwayat sprint closed), kolom rank, dan endpoint backlog adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/sprint` dan `internal/features/backlog`) mengikuti pola Clean Architecture repo, dengan halaman Backlog sebagai route baru di SPA `frontend/`.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — Manage Sprints (PERM-025), Schedule Issues/Edit Issues (PERM-024), Create/Edit/Delete Issues (PERM-021), role team-managed (PERM-039, PERM-040).
- [04 Proyek](04-proyek.md) — Features toggle Backlog/Sprints (PRJ-049, PRJ-050), company-managed vs team-managed (PRJ-003).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — Sprint field (CFG-014), Rank field (CFG-015), Story Points (CFG-016), Parent field (CFG-013), Flagged (CFG-018), link type (CFG-056), terminologi Sprint/Epic (CFG-010).
- [06 Work Item Inti](06-work-item-inti.md) — Rank (ISS-022), Flag (ISS-023), inline create (ISS-005), issue view, bulk change, delete, issue link.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status category, transisi dari lozenge status.
- [10 Board](10-board-kolom-kartu.md) — filter board, pemetaan kolom, kolom paling kanan, unmapped status, card layout, estimation statistic, Active Sprint Board, issue detail panel, Features toggle team-managed, filter & search backlog.
- [12 JQL](12-jql-bahasa-kueri.md) — fungsi `openSprints()`, `closedSprints()`, `futureSprints()`, `ORDER BY Rank`.
- [14 Pelacakan Waktu & Estimasi](14-pelacakan-waktu-estimasi.md) — estimation statistic, time tracking, rollover Remaining Estimate.
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Releases page, dialog release, Fix Version.
- [16 Laporan & Insight](16-laporan-insight.md) — Sprint report, Burndown, Velocity, Insights panel.
- [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md) — epic di Timeline team-managed, sprint marker.
- [19 Notifikasi](19-notifikasi-email.md) — notifikasi perubahan work item.
- [20 Otomasi](20-otomasi.md) — Sprint triggers, sprint management via automation.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — keyboard shortcut board & backlog.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — Agile Backlog & Rank API, Agile Sprints API, webhook sprint.
- [25 Integrasi & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — penautan site wiki/dokumentasi eksternal, impor sprint via CSV.
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — realtime, virtualisasi daftar besar, konkurensi rank.
- [28 Model Data Inti](28-model-data-inti.md) — entitas Board, Sprint, Rank dan katalog event `sprint.*`.

## Persona terkait

- **Scrum Master/Team Lead** — membuat, mengedit, memulai, menyelesaikan, membuka kembali, dan menghapus sprint; memfasilitasi sprint planning memakai beban kerja per assignee.
- **Product Owner/PM** — menjaga urutan backlog dengan Rank, memecah work item besar, mengelompokkan per Epic dan Version.
- **Developer/Anggota Tim** — membuat work item inline, memberi estimasi, menarik pekerjaan ke sprint, melihat sprint goal.
- **Admin Proyek/Project Lead** — mengatur siapa yang memegang Manage Sprints, mengaktifkan Backlog/Sprints di Features, merilis versi dari board.
- **Admin Situs** — mengaktifkan Parallel Sprints secara global dan mengatur permission scheme.
- **Stakeholder/Viewer** — membaca backlog, sprint goal, dan halaman tertaut sprint tanpa mengubahnya.
- **Integrator/Developer API** — mengelola sprint dan rank lewat REST API serta menerima event `Sprint started`/`Sprint completed`.

## User story modul

- Sebagai Product Owner/PM, saya ingin mengurutkan backlog dengan drag-and-drop dan aksi Move to top/bottom agar tim selalu mengerjakan hal terpenting lebih dulu (SPR-001, SPR-003, SPR-006).
- Sebagai Developer/Anggota Tim, saya ingin membuat work item langsung di sprint atau backlog cukup dengan mengetik summary agar pencatatan ide saat planning tidak memperlambat rapat (SPR-005).
- Sebagai Scrum Master/Team Lead, saya ingin memilih banyak work item sekaligus dan memindahkannya ke sprint dengan urutan tetap agar sprint planning cepat (SPR-004, SPR-012).
- Sebagai Developer/Anggota Tim, saya ingin mengisi estimasi inline dan melihat total estimasi per sprint agar komitmen sprint realistis (SPR-007, SPR-011).
- Sebagai Product Owner/PM, saya ingin memfilter backlog per Epic dan Version dan men-drag work item ke Epic/Version agar backlog tertata sesuai inisiatif dan rilis (SPR-008, SPR-009).
- Sebagai Product Owner/PM, saya ingin memecah work item yang terlalu besar menjadi beberapa work item tertaut agar setiap bagian muat dalam satu sprint (SPR-010).
- Sebagai Scrum Master/Team Lead, saya ingin melihat beban kerja per assignee dalam sprint agar pembagian kerja seimbang sebelum sprint dimulai (SPR-011).
- Sebagai Scrum Master/Team Lead, saya ingin membuat, menamai, mengatur tanggal, dan menetapkan goal sprint agar tim punya fokus dan batas waktu yang jelas (SPR-014, SPR-015, SPR-016, SPR-017).
- Sebagai Scrum Master/Team Lead, saya ingin memulai sprint dengan snapshot komitmen awal agar perubahan scope dan burndown dapat dilacak (SPR-018, SPR-019).
- Sebagai Scrum Master/Team Lead, saya ingin menyelesaikan sprint dan memilih tujuan work item yang belum selesai agar sisa pekerjaan tidak hilang (SPR-020).
- Sebagai Admin Proyek/Project Lead, saya ingin hanya orang tertentu yang dapat memulai dan menyelesaikan sprint agar ritme Scrum terjaga (SPR-013).
- Sebagai Admin Situs, saya ingin mengizinkan beberapa sprint aktif sekaligus agar tim besar dengan beberapa sub-tim bisa berbagi satu board (SPR-023).
- Sebagai Developer/Anggota Tim di tim Kanban, saya ingin punya backlog terpisah dari board agar board hanya berisi pekerjaan yang siap dikerjakan (SPR-024).
- Sebagai Admin Proyek/Project Lead, saya ingin merilis versi langsung dari kolom terakhir board Kanban agar rilis berkelanjutan tidak memerlukan langkah terpisah (SPR-025).
- Sebagai Integrator/Developer API, saya ingin setiap work item menyimpan sprint aktif/future dan riwayat sprint closed serta menerbitkan event sprint agar laporan dan otomasi eksternal akurat (SPR-012, SPR-018, SPR-020).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 11 |
| P1 | 7 |
| P2 | 7 |
| P3 | 1 |
| **Total** | **26** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 11 |
| Fase 2 | 10 |
| Fase 3 | 4 |
| Fase 4 | 1 |
| Fase — (dc-only/deprecated, rujukan) | 0 |

Strategi urutan implementasi: Fase 1 membangun tulang punggung Scrum, yaitu entitas Sprint dengan state future/active/closed dan riwayat sprint per work item (SPR-012), penegakan Manage Sprints serta Schedule Issues/Edit Issues di server (SPR-013), halaman Backlog dengan bagian sprint dan Backlog (SPR-001), ranking LexoRank dengan drag-and-drop (SPR-003), inline create (SPR-005), estimasi inline (SPR-007), pencatatan scope change (SPR-019), serta create/edit/start/complete sprint (SPR-014, SPR-015, SPR-018, SPR-020). Sejak Fase 1 setiap perubahan sprint dan Sprint field wajib menulis history dan menerbitkan event domain (DM-010), snapshot committed scope dicatat saat start, dan scope change dicatat per event, karena Sprint report (RPT-010) dan Burndown (RPT-009) juga berada di Fase 1 dan tidak dapat direkonstruksi belakangan; karena itu SPR-007 dan SPR-019 dinaikkan dari P1 inventaris menjadi P0 agar selaras dengan TIME-025, RPT-009, dan RPT-010. Fase 2 menambah produktivitas planning: multi-select, menu konteks, panel Epic dan Version, sprint goal, delete dan reorder sprint, release dari board, Kanplan (selaras BRD-043), dan parallel sprints (selaras toggle situs ADM-026). Fase 3 menambah kapabilitas lanjutan: reopen sprint, split issue, workload per assignee, dan preferensi tampilan backlog. Fase 4 menambah tautan halaman dokumentasi ke sprint yang bergantung pada integrasi wiki eksternal (modul 25).

## Fitur

## Backlog: tampilan, urutan & interaksi

### SPR-001 Backlog View

- **Deskripsi:** Halaman Backlog pada board Scrum menampilkan sprint aktif dan future di bagian atas serta daftar work item yang belum dijadwalkan (bagian Backlog) di bawahnya, sebagai ruang kerja utama sprint planning.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Scrum Master, Product Owner; melihat backlog memerlukan akses ke board (modul 10) dan Browse Projects pada project yang tercakup filter board; hanya work item yang boleh dilihat pengguna yang tampil.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman tersusun dari satu bagian per sprint, dimulai dari sprint aktif lalu future sprint sesuai urutan sprint (SPR-017), diikuti satu bagian "Backlog" di paling bawah.
  2. Header setiap sprint menampilkan nama sprint, rentang tanggal (bila sudah diisi), jumlah work item, total estimasi sesuai estimation statistic board, dan ringkasan estimasi per status category (To Do / In Progress / Done) dalam tiga badge berwarna; header Backlog menampilkan jumlah work item dan total estimasi.
  3. Setiap bagian sprint dan bagian Backlog dapat di-collapse/expand; status collapse diingat per pengguna per board di sisi klien.
  4. Setiap baris work item menampilkan ikon work type, key, summary, lozenge parent/Epic, Fix Version, lozenge status, badge estimasi, avatar assignee, ikon priority, ikon flag, dan field tambahan yang dikonfigurasi di card layout backlog (modul 10); key work item yang resolved ditampilkan dicoret.
  5. Company-managed: work item yang statusnya dipetakan ke kolom paling kanan board tidak tampil di bagian Backlog; work item done yang berada di dalam sprint tetap tampil di bagian sprint tersebut.
  6. Work item dengan status yang tidak dipetakan ke kolom mana pun (unmapped, modul 10) tidak tampil di backlog.
  7. Sub-task tidak tampil sebagai baris tersendiri; jumlah sub-task ditampilkan sebagai indikator pada baris parent (lihat BRD-022 Sub-tasks and hierarchy on board).
  8. Bagian Backlog yang kosong menampilkan empty state "Backlog Anda kosong" beserta tombol Create; sprint kosong menampilkan area drop dengan ajakan "Rencanakan sprint dengan men-drag work item dari Backlog ke sini".
  9. Klik baris work item membuka issue detail panel di sisi kanan tanpa meninggalkan halaman dan menambahkan `?selectedIssue=KEY-123` ke URL; Esc atau tombol tutup menutup panel.
  10. Header halaman menampilkan jumlah work item total; saat filter/quick filter/pencarian aktif, setiap bagian menampilkan "x dari y work item" agar jelas ada work item yang disembunyikan.
  11. Daftar memakai lazy loading dan virtualisasi baris sehingga backlog dengan 5.000 work item tetap dapat di-scroll tanpa memuat seluruh baris ke DOM; data dimuat bertahap per halaman dari API.
  12. Backlog hanya menampilkan work item yang cocok dengan filter board (saved filter/JQL board); work item yang tidak cocok tidak tampil di bagian mana pun, termasuk bila Sprint field-nya menunjuk sprint board ini.
  13. Menu "..." pada header sprint menyediakan akses ke ringkasan beban kerja per assignee (SPR-011) selain aksi sprint lainnya.
  14. Team-managed: Backlog dapat diaktifkan/dimatikan di Features project (PRJ-049); bila dimatikan, menu Backlog hilang dari navigasi project dan work item baru langsung tampil di kolom pertama board.
- **Team-managed vs company-managed:** Company-managed menyembunyikan work item kolom paling kanan dari bagian Backlog berdasarkan pemetaan kolom; team-managed memakai toggle Backlog di Features dan kolom Done board sebagai kolom paling kanan.
- **Catatan & ketergantungan:** Pencarian, filter assignee/Epic/type/label, dan quick filter di backlog didefinisikan di BRD-034 (Board and Backlog Filters & Search); issue detail panel di BRD-018; tampilan dan total estimasi backlog selaras dengan TIME-025; Insights panel di modul 16. Endpoint ringkas yang dibutuhkan: `GET /api/v1/boards/{id}/backlog` dan `GET /api/v1/boards/{id}/sprints?state=active,future` dengan paginasi (modul 23).

### SPR-002 Backlog View Settings

- **Deskripsi:** Pengaturan tampilan backlog per pengguna untuk menampilkan atau menyembunyikan panel samping dan mengatur kepadatan baris.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat backlog; tidak memerlukan permission tambahan karena hanya memengaruhi tampilan pengguna itu sendiri.
- **Aturan perilaku (acceptance criteria):**
  1. Menu "View settings" di toolbar backlog memiliki toggle tampil/sembunyi Epic panel (SPR-008); bila disembunyikan, filter Epic yang aktif dari panel dihapus.
  2. Company-managed: menu yang sama memiliki toggle tampil/sembunyi Version panel (SPR-009); toggle tidak tampil bila fitur Releases tidak aktif.
  3. ⚠ Tersedia opsi tampilan ringkas (compact) yang mengurangi tinggi baris dan menyembunyikan field tambahan card layout.
  4. ⚠ Preferensi tampilan disimpan per pengguna per board di server sehingga berlaku di perangkat lain; nilai default untuk pengguna baru adalah semua panel tersembunyi dan mode normal.
- **Team-managed vs company-managed:** Team-managed tidak menampilkan toggle Version panel bila Releases tidak diaktifkan di Features.
- **Catatan & ketergantungan:** Isi menu View settings berubah antar rilis produk referensi; daftar opsi final menjadi keputusan desain. Preferensi disimpan di tabel preferensi pengguna (modul 21).

### SPR-003 Backlog Ranking (Drag and Drop)

- **Deskripsi:** Mengurutkan prioritas work item dengan drag-and-drop yang menyimpan nilai Rank global per work item (LexoRank-kompatibel), sekaligus memindahkan work item antar sprint dan Backlog.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Product Owner, Scrum Master, Developer/Anggota Tim dengan Schedule Issues (ranking) serta Schedule Issues + Edit Issues (pindah sprint), PERM-024.
- **Aturan perilaku (acceptance criteria):**
  1. Drag work item naik/turun di dalam bagian Backlog atau di dalam sebuah sprint mengubah Rank work item tersebut menjadi nilai di antara dua tetangga barunya; urutan baru langsung tampil (optimistic update) dan dikembalikan bila server menolak.
  2. Drag antar bagian (Backlog ke sprint, sprint ke sprint, sprint ke Backlog) mengubah Sprint field dan Rank dalam satu operasi atomik; bila salah satu gagal, keduanya dibatalkan.
  3. Rank bersifat global per work item sehingga urutan konsisten di semua board yang filternya diakhiri `ORDER BY Rank` (ISS-022, CFG-015).
  4. Saat drag, garis indikator posisi drop tampil di antara baris dan halaman auto-scroll ketika kursor mendekati tepi atas/bawah viewport; drop di luar area valid membatalkan drag.
  5. Bila filter board tidak diakhiri `ORDER BY Rank`, drag untuk mengubah urutan dinonaktifkan (drag antar bagian tetap boleh) dan banner menampilkan pesan "Ranking dinonaktifkan karena filter board tidak diurutkan berdasarkan Rank" dengan tautan ke pengaturan filter untuk board admin.
  6. Ranking memerlukan Schedule Issues; memindahkan work item antar sprint memerlukan Schedule Issues dan Edit Issues; tanpa izin, handle drag tidak tampil dan API mengembalikan 403 dengan pesan permission yang kurang.
  7. Menu konteks menyediakan aksi Move to top of backlog, Move to bottom of backlog, dan pindah ke sprint tertentu (SPR-006) dengan hasil sama seperti drag.
  8. Shortcut keyboard `s` lalu `t` memindahkan work item terpilih ke atas bagian tempatnya berada (send to top), dan `s` lalu `b` ke bawah (send to bottom); shortcut terdaftar di katalog shortcut NAV-040 (modul 21).
  9. Sub-task di-rank relatif terhadap sub-task lain di bawah parent yang sama; sub-task tidak dapat di-drag keluar dari parent-nya di backlog.
  10. Perubahan urutan oleh pengguna lain tercermin saat refresh dan, setelah realtime tersedia (Fase 2, modul 27), tanpa refresh; bila work item tujuan drop sudah berpindah, server menghitung posisi berdasarkan tetangga terkini tanpa error.
  11. Penyisipan rank hanya menghitung nilai tengah antara dua tetangga tanpa memperbarui work item lain; rebalancing berkala berjalan di latar belakang oleh sistem tanpa mengubah urutan yang terlihat.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed selalu memakai urutan Rank pada board-nya.
- **Catatan & ketergantungan:** Algoritme rank dan rebalancing didefinisikan di ISS-022, DM-010, dan modul 27; endpoint `PUT /api/v1/backlog/rank` (rankBefore/rankAfter) di API-033 (Agile Backlog & Rank API). History work item mencatat "Ranked higher/lower".

### SPR-004 Backlog Multi-select & Bulk Actions

- **Deskripsi:** Memilih banyak work item sekaligus di backlog untuk dipindah, di-rank, di-flag, atau diedit massal.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Product Owner, Scrum Master; setiap aksi memeriksa permission per work item (Schedule Issues, Edit Issues, Make bulk changes PERM-011 untuk Bulk change).
- **Aturan perilaku (acceptance criteria):**
  1. Ctrl+klik (Windows/Linux) atau Cmd+klik (macOS) men-toggle pilihan satu baris; Shift+klik memilih rentang dari baris terpilih terakhir sampai baris yang diklik di dalam bagian yang sama.
  2. Men-drag salah satu baris terpilih memindahkan semua work item terpilih sekaligus ke posisi/bagian tujuan dengan urutan relatif di antara mereka dipertahankan.
  3. Selama drag multi-item, badge jumlah ("5 work item") ditampilkan pada kursor drag.
  4. Klik kanan atau tombol "..." pada seleksi membuka menu dengan Move to (top/bottom/sprint tertentu/Backlog), Add flag, Remove flag, dan Bulk change.
  5. "Bulk change" membuka wizard bulk edit (modul 06) dengan work item terpilih sudah terisi; bila jumlahnya melebihi batas bulk change, pengguna diberi tahu batasnya.
  6. Seleksi direset setelah aksi selesai atau saat pengguna mengklik area kosong; Esc membatalkan seleksi tanpa mengubah apa pun.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bila sebagian work item gagal dipindah karena izin, operasi diterapkan pada yang berhasil dan ringkasan kegagalan ditampilkan per key. Bergantung pada SPR-003, SPR-006, Flag (ISS-023), dan Bulk Change (ISS-035).

### SPR-005 Inline Issue Create in Backlog

- **Deskripsi:** Membuat work item langsung di bagian Backlog atau di sebuah sprint tanpa membuka Create dialog penuh.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Create Issues pada project tujuan; membuat di sprint juga memerlukan Schedule Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "+ Create" tampil di bawah setiap bagian sprint dan bagian Backlog; tombol tidak tampil bila pengguna tidak punya Create Issues di project mana pun pada board.
  2. Pemilih work type menampilkan tipe standar yang tersedia di project (default: tipe yang terakhir dipakai pengguna atau default work type project); sub-task tidak tersedia di pemilih.
  3. Pengguna mengetik summary lalu Enter untuk membuat; input tetap terbuka dan kosong untuk work item berikutnya dengan work type yang sama.
  4. Work item yang dibuat di sebuah sprint otomatis mendapat Sprint field sprint tersebut; yang dibuat di Backlog tidak punya sprint open.
  5. Work item baru ditempatkan di posisi paling bawah bagian tempat ia dibuat (Rank tepat di bawah work item terakhir bagian itu).
  6. Bila work type yang dipilih memiliki field wajib selain Summary, Create dialog lengkap (ISS-002) dibuka dengan summary, work type, dan sprint sudah terisi.
  7. Bila work item baru tidak cocok dengan filter board atau filter aktif, tampil notifikasi "Work item KEY-123 dibuat tetapi tidak terlihat karena filter" beserta tautan ke work item.
  8. ⚠ Bila sebuah Epic sedang dipilih di Epic panel/filter Epic, work item baru otomatis mendapat Parent Epic tersebut.
  9. ⚠ Pada board multi-project, input inline menyediakan pemilih project tujuan (default: project pertama pada filter board yang pengguna punya Create Issues).
  10. Esc membatalkan input tanpa membuat work item; summary dibatasi 255 karakter dengan penghitung saat mendekati batas dan pesan galat bila terlampaui; summary kosong tidak dapat dikirim.
- **Team-managed vs company-managed:** Tidak ada perbedaan perilaku; sumber daftar work type berbeda (skema company-managed vs work type team-managed, modul 05).
- **Catatan & ketergantungan:** Validasi sama dengan ISS-003/ISS-005. Bila sprint tujuan aktif, berlaku peringatan scope change (SPR-019).

### SPR-006 Backlog Issue Context Menu

- **Deskripsi:** Menu aksi per work item di backlog (klik kanan atau tombol "...") untuk operasi cepat tanpa membuka issue view.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Product Owner; setiap item menu hanya tampil bila pengguna memiliki permission yang dibutuhkan (dihitung di server).
- **Aturan perilaku (acceptance criteria):**
  1. Submenu "Move to" berisi Top of backlog, Bottom of backlog, daftar sprint active/future board ini, dan Backlog (bila work item sedang di sprint); memerlukan Schedule Issues (+ Edit Issues bila pindah sprint).
  2. Aksi Add flag, Add flag and comment (membuka input comment wajib), dan Remove flag (bila sudah di-flag) mengikuti ISS-023.
  3. Aksi "Change parent" membuka pemilih Epic/parent yang valid sesuai hierarki (CFG-013); memerlukan Edit Issues.
  4. ⚠ Assignee dan estimasi dapat diubah dari menu (submenu Assignee dengan pencarian pengguna dan input estimasi), selain lewat edit inline pada baris.
  5. Aksi "Split issue" membuka dialog SPR-010; tidak tampil untuk sub-task.
  6. Aksi "Copy issue link" menyalin URL absolut `/browse/KEY-123` dan "Copy issue key" menyalin `KEY-123`, masing-masing dengan konfirmasi "Disalin".
  7. Aksi "Delete" tampil hanya untuk pengguna dengan Delete Issues dan membuka konfirmasi penghapusan modul 06.
  8. Saat beberapa work item terpilih, menu menampilkan aksi massal termasuk "Bulk change" (SPR-004).
  9. ⚠ Klik lozenge status pada baris membuka daftar transisi yang tersedia (modul 09); transisi yang memerlukan screen membuka dialog transisi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Daftar aksi tersedia diambil dari endpoint operations work item (modul 06). Keyboard shortcut terkait didefinisikan di modul 21.

### SPR-007 Inline Estimation in Backlog

- **Deskripsi:** Mengisi dan melihat estimasi work item langsung di baris backlog sesuai estimation statistic board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim dengan Edit Issues; untuk estimasi berbasis waktu juga mengikuti aturan time tracking modul 14.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap baris menampilkan badge estimasi; klik badge mengubahnya menjadi input inline yang disimpan dengan Enter atau blur dan dibatalkan dengan Esc.
  2. Field yang diedit mengikuti estimation statistic board: Story Points (company-managed), Story point estimate (team-managed), Original Estimate, atau custom number field yang dipilih di pengaturan board (modul 10/14).
  3. Work item tanpa estimasi ditampilkan sebagai "-"; total sprint hanya menjumlahkan work item yang berestimasi, dan tooltip total menyebut jumlah work item tanpa estimasi.
  4. Header sprint menampilkan total estimasi sprint dan total per status category yang langsung diperbarui setelah estimasi berubah.
  5. ⚠ Estimasi sub-task tidak dijumlahkan ke total sprint, termasuk untuk estimation statistic berbasis waktu.
  6. Bila field estimasi tidak tersedia di screen/field context work type tersebut, edit ditolak dengan pesan "Field estimasi tidak tersedia untuk work type ini" dan badge tidak dapat diklik.
  7. Estimasi berbasis waktu menerima format durasi (mis. `1w 2d 3h 30m`) sesuai pengaturan time tracking; input tidak valid menampilkan galat format; nilai angka (story point) menerima desimal non-negatif.
- **Team-managed vs company-managed:** Company-managed memakai field Story Points; team-managed memakai Story point estimate (CFG-016).
- **Catatan & ketergantungan:** Perubahan estimasi pada work item di sprint aktif dicatat sebagai scope change (SPR-019). Aturan estimation statistic di TIME-022 (Board Estimation Statistic) dan TIME-025 (Estimation in Backlog & Sprint Planning). Prioritas dinaikkan dari P1 (inventaris) ke P0/Fase 1 karena TIME-025 (P0/Fase 1) memakai inline edit ini untuk sprint planning MVP.

## Panel perencanaan, pemecahan & beban kerja

### SPR-008 Epics Panel in Backlog

- **Deskripsi:** Panel samping di backlog untuk melihat, membuat, dan mengelola Epic serta memfilter backlog per Epic. Hubungan work item ke Epic memakai field Parent dalam hierarki terpadu (field Epic Link legacy tidak dipakai).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Product Owner, Developer/Anggota Tim; melihat panel memerlukan akses backlog; drag ke Epic dan rename/ubah warna memerlukan Edit Issues; create Epic memerlukan Create Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Panel Epic dapat ditampilkan/disembunyikan dari toolbar backlog atau View settings (SPR-002).
  2. Panel menampilkan daftar Epic yang belum selesai dengan warna, nama, progress bar (proporsi child per status category), jumlah child, dan total estimasi child.
  3. Klik Epic memfilter backlog dan sprint sehingga hanya menampilkan child Epic tersebut; opsi "All issues" menghapus filter dan "Issues without epic" menampilkan work item tanpa Parent Epic.
  4. Men-drag satu atau beberapa work item ke Epic di panel mengisi field Parent; drop pada work item yang tidak valid secara hierarki (mis. Epic ke Epic, sub-task) ditolak dengan pesan.
  5. Tombol "Create epic" di panel membuat Epic baru dengan nama (summary) dan, bila wajib, membuka Create dialog lengkap.
  6. Setiap Epic dapat di-expand untuk menampilkan detail: start date, due date, progress, dan tautan "View details" ke issue view Epic.
  7. ⚠ Menu Epic di panel menyediakan Rename, ubah warna Epic (palet tetap), dan Mark as done (company-managed) yang mentransisikan Epic ke status Done.
  8. Epic di panel dapat diurutkan dengan drag yang mengubah Rank Epic tersebut.
  9. ⚠ Epic dari project lain tampil di panel bila ada work item di board yang menjadi child Epic tersebut, dengan label project asal.
  10. Epic yang statusnya Done (status category Done) disembunyikan dari panel.
  11. Team-managed: Epic tidak ditampilkan sebagai panel tetapi sebagai filter "Epic" di toolbar backlog, dan pengelolaan Epic (urutan, tanggal) dilakukan di Timeline (modul 18).
- **Team-managed vs company-managed:** Lihat butir 7 dan 11.
- **Catatan & ketergantungan:** Parent field (CFG-013), Epic Colour legacy (CFG-024). Panel yang sama dipakai di Kanplan (SPR-024) dan didefinisikan untuk board Kanban di BRD-044 (Epics panel).

### SPR-009 Versions Panel in Backlog

- **Deskripsi:** Panel samping untuk melihat versi project dan menetapkan Fix Version langsung dari backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Product Owner, release manager (Admin Proyek); menetapkan versi memerlukan Edit Issues; membuat versi memerlukan Administer Projects atau izin pengelolaan versi.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: panel Versions dapat ditampilkan/disembunyikan dari toolbar backlog atau View settings (SPR-002).
  2. Panel menampilkan daftar versi unreleased dengan progress bar, jumlah work item, start date, dan release date; versi yang lewat release date diberi penanda terlambat.
  3. Klik versi memfilter backlog dan sprint ke work item dengan Fix Version tersebut; opsi "All issues" menghapus filter dan "Issues without version" menampilkan work item tanpa Fix Version.
  4. Men-drag satu atau beberapa work item ke versi menambahkan versi tersebut ke Fix Version (nilai lain tetap dipertahankan).
  5. Tombol "Create version" membuka form inline berisi name (wajib, unik per project), start date, release date, dan description; release date sebelum start date ditolak.
  6. Setiap versi memiliki tautan ke halaman Release/versi (modul 15) dan dari sana ke aksi Release.
  7. Panel yang sama tersedia di Kanban backlog (SPR-024).
  8. ⚠ Menetapkan versi memerlukan Edit Issues; membuat versi memerlukan Administer Projects atau izin Manage versions bila izin terpisah itu diadopsi (modul 03/15); tombol Create version disembunyikan bagi pengguna tanpa izin.
- **Team-managed vs company-managed:** Company-managed selalu menyediakan panel. ⚠ Team-managed menyediakan Releases bila diaktifkan di Features; ketersediaan panel Versions di backlog team-managed belum dipastikan dan menjadi keputusan terbuka.
- **Catatan & ketergantungan:** Bergantung pada modul 15 (Releases page, REL-021 Versions panel (backlog)); sisi board didefinisikan di BRD-045 (Versions panel).

### SPR-010 Split Issue

- **Deskripsi:** Memecah work item yang terlalu besar menjadi beberapa work item baru langsung dari backlog atau board, dengan tautan "split" ke work item asli.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Product Owner, Developer/Anggota Tim dengan Create Issues dan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Split issue" tersedia di menu "..." work item pada backlog (SPR-006) dan pada card board (modul 10).
  2. Dialog menampilkan summary dan estimasi work item asli yang dapat diubah, serta satu atau lebih baris work item baru (tombol "Add another") masing-masing dengan summary (wajib) dan estimasi.
  3. Dialog memiliki pemilih tujuan untuk work item baru: sprint asal (default), sprint active/future lain, atau Backlog.
  4. ⚠ Work item baru mewarisi field work item asli: project, work type, Parent, Priority, Labels, Components, Fix Version, dan assignee; field lain dikosongkan. Daftar final field yang diwarisi menjadi keputusan desain.
  5. Setiap work item baru di-link ke work item asli dengan link type "Issue split" (outward "split to", inward "split from"); link type ini dibuat sebagai link type sistem (CFG-056).
  6. Bila work item baru ditambahkan ke sprint aktif, penambahan tercatat sebagai scope change (SPR-019) dan peringatan scope change ditampilkan sebelum konfirmasi.
  7. Aksi tidak tersedia untuk sub-task.
  8. Pengguna tanpa Create Issues atau Edit Issues tidak melihat aksi ini; server menolak request dengan 403.
- **Team-managed vs company-managed:** ⚠ Tersedia di company-managed; ketersediaan di team-managed belum dipastikan. Produk ini mengusulkan menyediakan perilaku yang sama di keduanya (lihat Pertanyaan terbuka).
- **Catatan & ketergantungan:** Bila ada field wajib di work type tujuan yang tidak terisi, dialog meminta pengguna membuka Create dialog lengkap. Dibuat dalam satu transaksi agar tidak ada split parsial. Link type dan tautan mengikuti ISS-029 (Issue Links).

### SPR-011 Sprint Workload by Assignee

- **Deskripsi:** Ringkasan beban kerja per assignee untuk satu sprint di backlog, berisi jumlah work item dan total estimasi per orang, untuk menyeimbangkan beban saat sprint planning.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master, Admin Proyek, Developer/Anggota Tim; siapa pun yang dapat melihat backlog dapat membuka ringkasan ini.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Ringkasan dibuka dari menu "..." pada header sprint di backlog dengan nama aksi "Workload by assignee" dan ditampilkan sebagai dialog.
  2. Tabel berisi satu baris per assignee: avatar dan nama, jumlah work item, total estimasi sesuai estimation statistic board, dan total Remaining Estimate bila time tracking aktif.
  3. Work item tanpa assignee dikelompokkan di baris "Unassigned".
  4. Nilai dihitung ulang setiap kali work item ditambahkan ke atau dikeluarkan dari sprint, atau estimasi/assignee berubah; dialog yang terbuka diperbarui setelah aksi tersebut.
  5. Berlaku untuk sprint aktif dan future sprint.
  6. Hanya menghitung work item yang boleh dilihat pengguna (permission-aware, termasuk issue security); total tidak membocorkan jumlah work item tersembunyi.
  7. ⚠ Estimasi sub-task dihitung mengikuti konfigurasi estimation/time tracking board: untuk story point hanya parent yang dihitung; untuk time tracking Remaining Estimate sub-task dijumlahkan ke assignee sub-task.
  8. ⚠ Perilaku team-managed sama dengan company-managed (dialog tersedia di backlog team-managed yang memiliki Sprints aktif).
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Laporan workload lintas proyek ada di TIME-032 (User Workload Report) dan RPT-030 (User workload report). Kapasitas tim di Plans ada di modul 18.

## Model sprint & izin

### SPR-012 Sprint Field & Sprint Ownership

- **Deskripsi:** Field Sprint menghubungkan work item ke sprint; setiap sprint dimiliki board asal (origin board) tetapi dapat terlihat di board lain yang filternya mencakup work item di dalamnya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Developer/Anggota Tim, Scrum Master; mengubah Sprint field memerlukan Schedule Issues dan Edit Issues (PERM-024); sistem memelihara riwayat.
- **Aturan perilaku (acceptance criteria):**
  1. Field "Sprint" dapat diedit di issue view, Create dialog, dan bulk edit (modul 06) bila field ada di screen/layout.
  2. Sprint picker menampilkan sprint active dan future yang relevan (sprint dari board yang mencakup project work item, dikelompokkan per board, dengan pencarian nama); sprint closed tidak muncul dan tidak dapat dipilih manual; API menolak penetapan ke sprint closed dengan 400.
  3. Sebuah work item hanya dapat berada di satu sprint open (active atau future) pada satu waktu, tetapi menyimpan daftar riwayat sprint closed yang pernah memuatnya; memindahkan ke sprint open lain mengganti nilai open tanpa menghapus riwayat closed.
  4. Sprint tampil di backlog/board asal dan di board lain yang filternya cocok dengan minimal satu work item di sprint tersebut; aksi start/complete sprint dapat dilakukan dari board mana pun yang menampilkannya sesuai izin.
  5. Work item yang dimasukkan ke sprint tetapi tidak cocok dengan filter sebuah board tidak tampil di board tersebut, walaupun board itu adalah origin board sprint.
  6. Setiap perubahan Sprint field (tambah, pindah, keluarkan, carry-over saat complete) tercatat di History work item dengan nilai lama dan baru serta aktornya.
  7. JQL mendukung `sprint = <id atau nama>`, `sprint in openSprints()`, `sprint in closedSprints()`, `sprint in futureSprints()`, dan `sprint is EMPTY` (modul 12); `sprint = X` cocok untuk work item yang pernah berada di sprint X walaupun kini di sprint lain.
  8. Mengubah Sprint field memerlukan Schedule Issues dan Edit Issues; tanpa keduanya field tampil read-only dan API mengembalikan 403. Sub-task mewarisi sprint parent dan Sprint field-nya tidak dapat diubah terpisah (CFG-014).
  9. Entitas Sprint memiliki atribut: `id`, `name`, `state` (future/active/closed), `startDate`, `endDate`, `completeDate`, `goal`, `originBoardId`, dan `sequence` (urutan tampil); relasi work item–sprint disimpan di tabel penghubung dengan penanda open/closed agar riwayat tetap ada.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed menampilkan field Sprint hanya bila Sprints diaktifkan di Features (PRJ-049).
- **Catatan & ketergantungan:** Definisi field locked di CFG-014; skema tabel `sprints` dan `issue_sprints` (added_at/removed_at untuk riwayat) beserta event domain di DM-010 (Data Model: Boards, Sprints & Ranking (LexoRank)). Atribut `sequence` pada butir 9 belum tercantum di DM-010 dan perlu ditambahkan di modul 28. Endpoint Agile Sprints API di API-034.

### SPR-013 Manage Sprints Permission

- **Deskripsi:** Penegakan permission project yang mengontrol siapa yang dapat membuat, memulai, menyelesaikan, dan mengelola sprint pada backlog dan board. Definisi permission ada di PERM-025; fitur ini mengatur perilaku UI dan server di modul sprint.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs dan Admin Proyek memberikan permission lewat permission scheme (company-managed) atau role (team-managed); Scrum Master memakainya.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: permission "Manage sprints" tersedia di permission scheme (PERM-014) dan default diberikan ke role Administrators dan Members pada default scheme produk ini.
  2. Permission mencakup create, start, complete, reopen, edit (nama, tanggal, goal), reorder, dan delete sprint; setiap endpoint sprint memeriksa permission ini di server.
  3. ⚠ Pada board multi-project, pengguna memerlukan Manage Sprints di semua project yang work item-nya berada di sprint tersebut (untuk sprint kosong: semua project yang tercakup filter board); bila kurang di salah satu, aksi ditolak dengan daftar project yang kurang izin.
  4. Permission terkait yang tidak tercakup: Schedule Issues (rank, pindah sprint), Edit Issues (mengubah Sprint field), dan Transition Issues (memindah kolom di board aktif); memindahkan work item ke/dari sprint tidak memerlukan Manage Sprints.
  5. Tombol aksi sprint (Create sprint, Start sprint, Complete sprint, menu Edit/Delete/Move) disembunyikan atau dinonaktifkan dengan tooltip "Anda memerlukan izin Manage sprints di project X" bila pengguna tidak punya hak; panggilan API langsung mengembalikan 403.
  6. ⚠ Team-managed: dikontrol role Administrator dan Member (keduanya dapat mengelola sprint, Viewer tidak); custom role dengan izin manage sprints tersedia di plan Premium/Enterprise (PERM-040).
- **Team-managed vs company-managed:** Lihat butir 1 dan 6.
- **Catatan & ketergantungan:** Pemisahan izin Start/Complete vs Edit Sprints adalah fitur Data Center-only (PERM-067) dan tidak dibangun. Kondisi repo saat ini belum memiliki permission scheme; Fase 1 memakai pemetaan role sederhana yang kemudian dimigrasikan ke scheme penuh di Fase 3 tanpa mengubah semantik.

## Siklus hidup sprint

### SPR-014 Sprint Create

- **Deskripsi:** Membuat future sprint baru di backlog untuk perencanaan.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints (SPR-013).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Create sprint" tampil di header bagian Backlog (dan di bawah sprint terakhir) pada board Scrum.
  2. ⚠ Sprint baru diberi nama default berpola `<prefix> Sprint N`, dengan prefix key project board (atau nama board untuk board multi-project) dan N bilangan auto-increment per board yang dimulai dari 1.
  3. Sprint dapat dibuat kosong lalu diisi dengan men-drag work item ke dalamnya.
  4. Future sprint dapat diberi tanggal rencana (start/end) dan sprint goal melalui Edit sprint (SPR-015); tanggal tidak wajib saat create.
  5. Beberapa future sprint dapat dibuat berurutan; sprint baru selalu diletakkan di bawah future sprint terakhir.
  6. Pengguna tanpa Manage Sprints tidak melihat tombol; API `POST /api/v1/sprints` mengembalikan 403.
  7. Sprint mencatat board asal (`originBoardId`) yaitu board tempat sprint dibuat; API mewajibkan `originBoardId` dan menolak board Kanban.
- **Team-managed vs company-managed:** Team-managed hanya menampilkan tombol bila Sprints diaktifkan di Features.
- **Catatan & ketergantungan:** Menerbitkan event `sprint.created` (modul 28) dan webhook terkait (modul 23).

### SPR-015 Sprint Edit

- **Deskripsi:** Mengubah atribut future sprint atau sprint aktif.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi dibuka dari menu "..." header sprint > "Edit sprint"; team-managed juga menampilkan tautan "Add dates" pada future sprint yang belum bertanggal, yang membuka dialog yang sama.
  2. Dialog berisi Sprint name (wajib), Duration (1 week, 2 weeks, 3 weeks, 4 weeks, Custom), Start date, End date, dan Sprint goal; memilih durasi selain Custom menghitung End date dari Start date, sedangkan mengubah End date manual mengubah Duration menjadi Custom.
  3. Saat sprint aktif, nama, tanggal, dan goal tetap dapat diubah; perubahan tanggal memengaruhi garis guideline burndown dan indikator sisa hari di board aktif.
  4. ⚠ Sprint closed tidak dapat diedit dari UI (aksi tidak tersedia); API menolak edit sprint closed kecuali untuk name dan goal.
  5. ⚠ Validasi: End date harus setelah Start date dengan pesan "Tanggal selesai harus setelah tanggal mulai"; nama sprint dibatasi 30 karakter dengan penghitung dan pesan galat bila terlampaui.
  6. Pengguna tanpa Manage Sprints tidak melihat aksi Edit sprint; API mengembalikan 403.
  7. Perubahan nama sprint langsung tercermin di Sprint field semua work item terkait, di sprint picker, dan di hasil JQL berbasis nama.
- **Team-managed vs company-managed:** Lihat butir 1.
- **Catatan & ketergantungan:** Perubahan tanggal sprint aktif tidak mengubah committed scope (SPR-018). Batas 30 karakter adalah validasi UI/API; kolom `name` di DM-010 berkapasitas 255 karakter sehingga batas dapat diperlonggar tanpa migrasi. Event `sprint.updated` (DM-010).

### SPR-016 Sprint Goal

- **Deskripsi:** Pernyataan tujuan sprint yang ditampilkan di backlog, board aktif, dan laporan sprint.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master (mengisi dengan Manage Sprints), Developer/Anggota Tim dan Stakeholder (membaca).
- **Aturan perilaku (acceptance criteria):**
  1. Sprint goal diisi di dialog Edit sprint (SPR-015) atau di dialog Start sprint (SPR-018); goal bersifat opsional dan berupa teks biasa.
  2. Goal ditampilkan di header sprint pada backlog dan di header Active Sprint Board (modul 10).
  3. Teks goal yang melebihi satu baris dipotong dengan elipsis; klik atau hover menampilkan teks lengkap.
  4. Goal ditampilkan di Sprint report (modul 16) untuk sprint aktif maupun closed.
  5. Goal dapat diubah saat sprint aktif melalui Edit sprint dan perubahan langsung terlihat di backlog dan board.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Batas panjang goal belum ditentukan; usulan 1.000 karakter.

### SPR-017 Sprint Reorder

- **Deskripsi:** Mengubah urutan future sprint di backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. Menu "..." header future sprint menyediakan "Move sprint up" dan "Move sprint down"; aksi yang tidak berlaku (sprint paling atas/bawah) disembunyikan.
  2. Urutan future sprint menentukan sprint teratas yang tombol "Start sprint"-nya aktif berikutnya (SPR-018).
  3. Future sprint tidak dapat dipindah ke atas sprint aktif; sprint aktif selalu tampil paling atas.
  4. Pengguna tanpa Manage Sprints tidak melihat aksi; API mengembalikan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Urutan disimpan di atribut `sequence` sprint (SPR-012); API swap sprint di modul 23.

### SPR-018 Sprint Start

- **Deskripsi:** Memulai future sprint sehingga menjadi sprint aktif, mencatat komitmen awal, dan menampilkannya di Active Sprint Board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints di semua project yang terlibat (SPR-013).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Start sprint" aktif pada future sprint teratas; bila Parallel Sprints aktif (SPR-023), future sprint mana pun dapat di-start walaupun sudah ada sprint aktif.
  2. Dialog Start sprint menampilkan jumlah work item yang akan dimulai, Sprint name, Duration (1/2/3/4 weeks/Custom), Start date (default: sekarang, dengan jam), End date (dihitung dari durasi, dapat diubah), dan Sprint goal; nilai yang sudah diisi lewat Edit sprint dipakai sebagai default.
  3. Bila sudah ada sprint aktif dan Parallel Sprints tidak aktif, tombol nonaktif dengan tooltip "Selesaikan sprint aktif sebelum memulai sprint baru"; API mengembalikan 409.
  4. ⚠ Sprint tanpa work item tidak dapat di-start: tombol nonaktif dengan tooltip "Tambahkan work item ke sprint sebelum memulainya".
  5. Start date di masa lalu diperbolehkan (untuk mencatat sprint yang sudah berjalan); End date harus setelah Start date.
  6. Saat start, sistem mencatat snapshot committed scope (daftar work item beserta estimasinya dan statusnya) yang dipakai Sprint report dan Burndown (modul 16); snapshot tidak berubah oleh edit berikutnya.
  7. Setelah start berhasil, pengguna diarahkan ke Active Sprint Board dengan sprint tersebut terpilih.
  8. Pengguna tanpa Manage Sprints tidak melihat tombol Start sprint; API mengembalikan 403.
  9. Start sprint menerbitkan event `Sprint started` yang dapat memicu rule otomasi (modul 20) dan webhook (modul 23), berisi id sprint, board, tanggal, dan aktor.
- **Team-managed vs company-managed:** Tidak ada perbedaan; team-managed memerlukan Sprints aktif di Features.
- **Catatan & ketergantungan:** State sprint berubah `future` → `active` secara atomik bersama snapshot (`sprint_start_snapshots`, DM-010); kegagalan snapshot membatalkan start. Event sprint untuk otomasi dan webhook dirinci di BRD-053 dan AUTO-026.

### SPR-019 Sprint Scope Change

- **Deskripsi:** Penambahan atau pengeluaran work item dan perubahan estimasi setelah sprint dimulai dicatat sebagai perubahan scope agar laporan sprint akurat.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master, Developer/Anggota Tim dengan Schedule Issues dan Edit Issues; sistem mencatat event scope.
- **Aturan perilaku (acceptance criteria):**
  1. Menambah work item ke sprint aktif lewat drag dari Backlog, edit Sprint field, inline create, atau split menampilkan peringatan "Scope sprint akan berubah" dengan tombol konfirmasi; bila dibatalkan, work item tidak dipindah.
  2. Mengeluarkan work item dari sprint aktif ke Backlog atau ke sprint lain diizinkan dan dicatat sebagai pengurangan scope beserta waktu dan aktornya.
  3. Perubahan nilai estimation statistic pada work item di sprint aktif setelah start dicatat sebagai scope change (nilai lama, nilai baru, waktu).
  4. Work item yang ditambahkan setelah start ditandai dengan asterisk (`*`) di Sprint report beserta keterangan "ditambahkan setelah sprint dimulai".
  5. Burndown chart menampilkan kenaikan dan penurunan scope pada waktu terjadinya (modul 16).
  6. Menambah atau mengeluarkan work item memerlukan Schedule Issues dan Edit Issues; tanpa izin, aksi ditolak dengan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Data scope change disimpan sebagai log event per sprint (`sprint_scope_events`, DM-010) yang dibaca oleh mesin kalkulasi laporan historis (RPT-009 Burndown, RPT-010 Sprint report). Terkait TIME-029 (Sprint Report (Estimation & Scope Change)). Prioritas dinaikkan dari P1 (inventaris) ke P0/Fase 1 karena RPT-009 dan RPT-010 (P0/Fase 1) membutuhkan data scope change sejak sprint pertama.

### SPR-020 Sprint Complete

- **Deskripsi:** Menyelesaikan sprint aktif dan menentukan tujuan work item yang belum selesai.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints di semua project yang terlibat.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Complete sprint" tampil di header Active Sprint Board dan di header sprint aktif pada backlog.
  2. Dialog menampilkan jumlah work item completed dan open pada sprint tersebut.
  3. Company-managed: work item dianggap selesai bila statusnya dipetakan ke kolom paling kanan board tempat sprint diselesaikan.
  4. ⚠ Team-managed: work item dianggap selesai bila statusnya berada di kolom paling kanan board (status category Done).
  5. Untuk open work item, dialog menyediakan pilihan tujuan: Backlog, future sprint tertentu (daftar future sprint board), atau "New sprint" (future sprint baru dibuat otomatis dengan nama default SPR-014); default adalah future sprint teratas bila ada, selain itu Backlog.
  6. Sub-task ikut berpindah bersama parent-nya ke tujuan yang sama; sub-task yang selesai dari parent yang belum selesai tetap ikut parent.
  7. Sprint tidak dapat diselesaikan bila ada parent di kolom paling kanan yang masih memiliki sub-task belum selesai; dialog menampilkan galat yang menyebut key work item terkait dan tombol Complete dinonaktifkan sampai diperbaiki.
  8. Work item yang dipindah tetap mencatat sprint closed di riwayat Sprint field (SPR-012), sehingga sebuah work item dapat memiliki riwayat banyak sprint.
  9. Sprint dapat diselesaikan sebelum End date; `completeDate` dicatat sebagai waktu complete aktual.
  10. Pengguna tanpa Manage Sprints tidak melihat tombol Complete sprint; API mengembalikan 403.
  11. Setelah complete berhasil, pengguna diarahkan ke Sprint report sprint tersebut (modul 16).
  12. Complete sprint menerbitkan event `Sprint completed` untuk otomasi (modul 20) dan webhook (modul 23).
- **Team-managed vs company-managed:** Lihat butir 3 dan 4.
- **Catatan & ketergantungan:** Penanganan Remaining Estimate open work item saat sprint selesai di TIME-026 (Remaining Estimate Rollover on Sprint Completion). Event otomasi/webhook dirinci di BRD-053 dan AUTO-026. Halaman retrospective yang ditautkan ke sprint dibahas di SPR-026. Complete sprint berlangsung dalam satu transaksi; kegagalan memindahkan sebagian work item membatalkan seluruh operasi.

### SPR-021 Sprint Reopen

- **Deskripsi:** Membuka kembali sprint yang telah selesai (company-managed) untuk mengoreksi penutupan yang terlalu dini.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Reopen sprint" tersedia di Sprint report untuk sprint closed pada board asal sprint.
  2. Reopen hanya dapat dilakukan bila tidak ada sprint aktif lain pada board, kecuali Parallel Sprints aktif; bila tidak memenuhi, aksi nonaktif dengan tooltip penjelasan.
  3. Work item yang dipindahkan ke Backlog/sprint lain saat complete tidak dikembalikan otomatis ke sprint yang dibuka kembali; pengguna harus memindahkannya manual.
  4. Saat sprint yang dibuka kembali diselesaikan ulang, Sprint report dan Velocity diperbarui memakai hasil complete terakhir.
  5. Pengguna tanpa Manage Sprints tidak melihat aksi; API mengembalikan 403.
  6. ⚠ Team-managed: aksi Reopen sprint tidak disediakan pada Fase 3 sampai perilaku referensi dipastikan.
- **Team-managed vs company-managed:** Lihat butir 6.
- **Catatan & ketergantungan:** State sprint berubah `closed` → `active`; `completeDate` dikosongkan dan riwayat complete sebelumnya dicatat di log sprint.

### SPR-022 Sprint Delete

- **Deskripsi:** Menghapus sprint; work item di dalamnya kembali ke Backlog.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Scrum Master atau pengguna dengan Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi dibuka dari menu "..." header sprint > "Delete sprint".
  2. Dialog konfirmasi menyebut nama sprint dan jumlah work item; setelah konfirmasi, semua work item di sprint kembali ke Backlog dan sprint tersebut dihapus dari Sprint field mereka (tanpa menjadi riwayat closed).
  3. ⚠ Future sprint dapat dihapus; sprint aktif juga dapat dihapus dengan konfirmasi tambahan bahwa data laporan sprint tersebut ikut hilang.
  4. Sprint closed tidak dapat dihapus dari UI (aksi tidak tersedia).
  5. Pengguna tanpa Manage Sprints tidak melihat aksi; API mengembalikan 403.
  6. Penomoran default sprint berikutnya tetap berlanjut dan tidak mengisi nomor sprint yang dihapus.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** History work item mencatat pengeluaran dari sprint yang dihapus. Event `sprint.deleted` (modul 28).

### SPR-023 Parallel Sprints

- **Deskripsi:** Mengizinkan lebih dari satu sprint aktif secara bersamaan di satu board, misalnya untuk beberapa sub-tim yang berbagi backlog.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs mengaktifkan; Scrum Master memakai dengan Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. Fitur diaktifkan secara global oleh Admin Situs di pengaturan aplikasi (bagian konfigurasi Software > Parallel sprints, ADM-026); default nonaktif.
  2. Saat aktif, Start sprint dapat dilakukan pada future sprint mana pun walaupun sudah ada sprint aktif.
  3. Active Sprint Board menampilkan pemilih sprint untuk memilih satu atau beberapa sprint aktif yang ditampilkan; pilihan diingat per pengguna.
  4. Complete sprint dilakukan per sprint secara terpisah dan tidak memengaruhi sprint aktif lain.
  5. Sprint report, Burndown, dan Insights tersedia per sprint aktif.
  6. Sebuah work item tetap hanya dapat berada di satu sprint open pada satu waktu.
- **Team-managed vs company-managed:** ⚠ Berlaku untuk board company-managed; dukungan di team-managed belum dipastikan dan menjadi keputusan terbuka (ADM-026 butir 3 mengusulkan team-managed ikut mendukung bila toggle situs dan fitur Sprints proyek aktif).
- **Catatan & ketergantungan:** Selaras dengan ADM-026: menonaktifkan toggle Parallel Sprints ditolak selama masih ada board dengan lebih dari satu sprint aktif, sampai sprint tambahan diselesaikan. Flag tersimpan sebagai `parallel_sprints_enabled` (DM-010); tanpa flag, start sprint kedua ditolak 409. Fase 2 dipilih (bukan Fase 3) agar toggle ADM-026 di Fase 2 berfungsi.

## Kanban backlog, rilis & tautan sprint

### SPR-024 Kanplan (Kanban Backlog)

- **Deskripsi:** Backlog untuk board Kanban agar tim Kanban dapat merencanakan dan memprioritaskan pekerjaan tanpa sprint, terpisah dari board.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Board admin (mengaktifkan di company-managed), Admin Proyek (toggle Features di team-managed), Developer/Anggota Tim (memakai); memindah work item memerlukan Schedule Issues dan Transition Issues.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: Kanban backlog diaktifkan dengan memetakan satu atau lebih status ke kolom khusus "Kanban backlog" di Board settings > Columns (modul 10); kolom ini muncul di paling kiri pengaturan kolom.
  2. Team-managed: Kanban backlog diaktifkan lewat toggle Backlog di Features project (PRJ-049).
  3. Backlog tampil di halaman terpisah dari board; work item dengan status yang dipetakan ke kolom Kanban backlog tidak tampil di board.
  4. Epics panel (SPR-008) dan Versions panel (SPR-009) tersedia di Kanban backlog company-managed, dengan opsi di pengaturan board untuk menampilkan Epic panel alih-alih menjadikan Epic sebagai baris.
  5. Work item dipindah dari backlog ke board (ke status kolom pertama board) lewat drag ke bagian "Selected for development"/kolom pertama atau aksi menu "Move to board"; perpindahan melakukan transisi status sesuai workflow dan gagal dengan pesan bila transisi tidak tersedia.
  6. Kanban backlog tidak memiliki sprint (tidak ada bagian sprint, tombol Create sprint, maupun Sprint field picker); ranking SPR-003, multi-select, inline create, dan menu konteks tetap berlaku.
- **Team-managed vs company-managed:** Lihat butir 1 dan 2.
- **Catatan & ketergantungan:** Fitur yang sama dari domain Kanban (BRD-043 "Kanplan (Kanban backlog)", P1/Fase 2) didefinisikan di modul 10 dari sisi board; modul ini menjadi acuan perilaku halaman backlog. Prioritas/fase disamakan dengan BRD-043 (inventaris domain Scrum menulis P2) agar aktivasi kolom Kanban backlog dan halaman backlog-nya rilis bersamaan.

### SPR-025 Release from Board

- **Deskripsi:** Merilis versi langsung dari board; pada Kanban dari kolom paling kanan, pada Scrum lewat halaman Releases.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Release manager, Admin Proyek dengan Administer Projects atau izin pengelolaan versi.
- **Aturan perilaku (acceptance criteria):**
  1. Kanban company-managed: tombol "Release" tampil di header kolom paling kanan board; menekannya menetapkan Fix Version ke semua work item di kolom tersebut yang terlihat pada board, lalu menandai versi released.
  2. Dialog release berisi version name (pilih versi unreleased yang ada atau ketik nama baru yang dibuat otomatis), release date (default hari ini), dan description; nama baru yang bentrok dengan versi existing ditolak.
  3. Board Scrum tidak memiliki tombol Release di kolom; rilis dilakukan dari halaman Releases yang dapat diakses dari Versions panel (SPR-009).
  4. Penanganan unresolved work item pada versi yang dirilis (pindahkan ke versi lain atau abaikan) dilakukan di dialog Release pada halaman Releases (modul 15).
- **Team-managed vs company-managed:** Team-managed menyediakan rilis lewat halaman Releases bila diaktifkan di Features.
- **Catatan & ketergantungan:** Pemilik: REL-022 ([15](15-rilis-versi-komponen.md)); BRD-047 untuk UI board ([10](10-board-kolom-kartu.md)). Prioritas P1 / Fase 2 diselaraskan dengan BRD-047 "Release from board" dan REL-022 "Release from Kanban board"; aturan rilis versi (Fix Version, tanggal rilis, unresolved work item) mengikuti REL-022; work item yang dirilis dari kolom paling kanan disembunyikan dari board Kanban sesuai perilaku Done column modul 10.

### SPR-026 Sprint Linked Pages (Confluence)

- **Deskripsi:** Menautkan atau membuat halaman dokumentasi (mis. retrospective, catatan sprint planning) pada sistem wiki yang terhubung, ke sprint tertentu, dan menampilkannya di Sprint report. Pada produk referensi ini adalah integrasi platform dengan wiki milik vendor yang sama; produk ini menyediakannya sebagai integrasi wiki generik.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Scrum Master, Admin Proyek, Developer/Anggota Tim dengan Manage Sprints untuk menautkan/menghapus tautan; pembaca memerlukan akses ke board.
- **Aturan perilaku (acceptance criteria):**
  1. Company-managed: header Sprint report menampilkan tautan "Linked pages" dengan jumlah halaman tertaut.
  2. Pengguna dapat membuat halaman retrospective baru di wiki yang terhubung dari template sprint/retrospective; judul default memuat nama sprint dan tanggal.
  3. Pengguna dapat menautkan halaman wiki yang sudah ada ke sprint tertentu dengan mencari judul atau menempelkan URL.
  4. Dialog Linked pages menampilkan daftar halaman tertaut per sprint (judul, penulis, tanggal); setiap tautan dapat dihapus tanpa menghapus halamannya.
  5. Fitur hanya tersedia bila ada site wiki yang terhubung (modul 25) dan pengguna memiliki izin yang sesuai di wiki tersebut; bila belum terhubung, tautan menampilkan ajakan menghubungkan site untuk Admin Situs.
  6. ⚠ Pengguna tanpa akses ke wiki melihat judul tautan tetapi tidak dapat membuka halaman; klik menampilkan pesan "Anda tidak memiliki akses ke halaman ini".
- **Team-managed vs company-managed:** Butir 1 berlaku untuk company-managed; team-managed tidak menyediakan Linked pages di Sprint report.
- **Catatan & ketergantungan:** Bergantung pada penautan site wiki di INT-042 (Confluence site linking & cross-product setup, Fase 4) dan Sprint report RPT-010. Produk wiki itu sendiri di luar cakupan.

## Pertanyaan terbuka

- Apakah Split Issue (SPR-010), Parallel Sprints (SPR-023), Sprint Reopen (SPR-021), dan Versions panel (SPR-009) juga disediakan di project team-managed, atau dibatasi ke company-managed seperti produk referensi?
- Field apa saja yang diwarisi oleh work item hasil Split Issue (SPR-010), dan apakah daftarnya dapat dikonfigurasi admin?
- Apakah batas panjang nama sprint tetap 30 karakter (SPR-015) atau diperlonggar, dan berapa batas panjang sprint goal (SPR-016)?
- Apakah sprint aktif boleh dihapus (SPR-022), mengingat data Sprint report dan Velocity ikut hilang?
- Apakah sprint kosong diizinkan untuk di-start (SPR-018) agar tim dapat memulai sprint lalu mengisinya?
- Apakah preferensi View settings backlog (SPR-002) dan pilihan sprint di Parallel Sprints (SPR-023) disimpan di server lintas perangkat atau cukup di browser?
- Untuk board multi-project, apakah syarat Manage Sprints dihitung dari isi sprint atau dari seluruh project pada filter board (SPR-013)?
- Apakah integrasi Linked pages (SPR-026) cukup dengan tautan URL generik ke wiki eksternal, atau memerlukan konektor resmi ke produk wiki tertentu?

## Di luar cakupan & ekosistem

- App Marketplace pihak ketiga untuk planning poker, estimasi kolaboratif, retrospective board, dan kapasitas sprint lanjutan; produk ini hanya menyediakan estimasi inline (SPR-007) dan workload per assignee (SPR-011).
- Produk wiki/dokumentasi (Confluence) itu sendiri beserta template retrospective-nya; produk ini hanya menyimpan tautan dan memakai integrasi modul 25 (SPR-026).
- Pemisahan izin Start/Complete sprint vs Edit sprint yang hanya ada di Data Center (PERM-067).
- Halaman manajemen rank/LexoRank untuk admin (hanya Data Center); rebalancing berjalan otomatis (ISS-022, CFG-015).
- Perencanaan sprint lintas tim dan kapasitas tim di Plans (Advanced Roadmaps) didefinisikan di modul 18, bukan di backlog board.
- Laporan sprint (Sprint report, Burndown, Velocity, Insights) didefinisikan di modul 16; modul ini hanya menyediakan data snapshot dan scope change.
- Layanan lintas produk vendor referensi (Atlassian Home/Goals, Loom, Rovo) yang terkait sprint tidak dibangun di modul ini.
