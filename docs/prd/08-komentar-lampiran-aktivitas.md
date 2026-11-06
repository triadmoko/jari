# 08 — Komentar, Lampiran & Aktivitas

Modul ini mendefinisikan seluruh lapisan kolaborasi dan jejak aktivitas pada work item di produk ini: comment (tambah, edit, hapus, visibility, thread reply, reaction, pin, permalink, urutan dan paginasi, event serta representasi API), attachment (upload, download, hapus, panel strip/list, Media Viewer, pengaturan dan batas, keamanan, siklus hidup saat work item di-clone/dipindah/dihapus, serta abstraksi penyimpanan berkas untuk attachment dan avatar), section Activity beserta tab-tabnya (All, Comments, History, Work log, Approvals), model changelog yang menjadi sumber History, laporan, dan predikat JQL historis, panel Automation di issue view, serta fitur berbagi dan konten tertaut (Share, web link/remote link, halaman wiki tertaut, embed video). Kondisi repo saat ini baru memiliki fitur auth dan user (tabel `users`, JWT di context `sess`, pesan galat EN/ID via `constants.GetError`); belum ada tabel comment, attachment, changelog, maupun abstraksi object storage, sehingga seluruh isi modul ini adalah pekerjaan baru yang dikemas sebagai feature module baru (mis. `internal/features/comment`, `internal/features/attachment`, dan `internal/features/activity`) dengan entity dan repository interface di `internal/shared/domain/` serta driver storage di `internal/shared/infrastructure/`.

Modul terkait:

- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — Former user (USR-008), undangan dari Share (USR-009, USR-012), swap group saat group dihapus (USR-015).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission comment (PERM-030), attachment (PERM-031), time tracking (PERM-032), workflow property (PERM-037), issue security (PERM-044), visibility comment/worklog (PERM-045), akses anonim (PERM-046), actor otomasi (PERM-049), data security policy (PERM-065), redaksi data (PERM-066).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field Attachment dan Comment di screen (CFG-046), field configuration (CFG-043), time tracking settings (CFG-058).
- [06 Work Item Inti](06-work-item-inti.md) — issue view, clone, move, delete, archive, issue link, web/remote link, Share Work Item, redaksi data.
- [07 Editor Teks Kaya](07-editor-konten-kaya.md) — comment editor ADF, mention, emoji, media inline, smart link (EDT-001 s.d. EDT-003).
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — transition screen dengan field Comment/Attachment, workflow property `jira.permission.comment*`.
- [12 JQL](12-jql-bahasa-kueri.md) — `comment ~`, `attachments IS EMPTY`, operator `WAS`/`CHANGED`.
- [13 Pencarian, Filter & Ekspor](13-pencarian-filter-ekspor.md) — export Word/XML/print yang memuat comment dan attachment.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — aturan worklog dan estimasi yang ditampilkan di tab Work log.
- [16 Laporan & Insight](16-laporan-insight.md) — Control chart, CFD, cycle time yang membaca tabel transisi status.
- [19 Notifikasi](19-notifikasi-email.md) — event Issue commented/comment edited/comment deleted, email Share, autowatch.
- [20 Otomasi](20-otomasi.md) — trigger Work item commented, aksi Add comment, audit log rule, manual trigger.
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — shortcut `m`, preferensi urutan Activity dan autowatch.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — General configuration (Comment visibility), Settings > System > Attachments.
- [23 API Publik & Webhook](23-api-publik-webhook.md) — REST comment, attachment, changelog, webhook `comment_*` dan `attachment_*`.
- [24 Integrasi DevOps](24-integrasi-devops.md) — panel Development untuk commit/PR yang menggantikan tab Source legacy (CMT-027).
- [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — import CSV dengan attachment, integrasi wiki/Confluence, Slack/Teams, panel app.
- [26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md) — approval pada project bisnis (BIZ-042, BIZ-046) yang ditampilkan di tab Approvals (CMT-025).
- [27 NFR & Arsitektur](27-nfr-arsitektur-platform.md) — job asinkron, partisi tabel, kuota, enkripsi.
- [28 Model Data Inti](28-model-data-inti.md) — tabel comment, attachment, change_groups/change_items, katalog event domain.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) — Internal note vs Reply to customer.
- [30 JSM Help Center, Portal & Kanal](30-jsm-help-center-portal-pelanggan-kanal.md) — email channel dan mail handler yang membuat comment dari email (JSM-PRT-041).
- [32 JSM ITSM & Approvals](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — approval step (JSM-ITSM-011) yang ditampilkan di tab Approvals (CMT-025).

## Persona terkait

- **Developer/Anggota Tim** — menulis, membalas, dan memberi reaction pada comment, melampirkan screenshot/log, membaca History untuk memahami perubahan.
- **Scrum Master/Team Lead** — memakai pinned comment untuk keputusan penting, memantau History dan Work log saat review sprint.
- **Product Owner/PM** — berdiskusi dengan stakeholder lewat comment dan Share, menautkan dokumen spesifikasi dan halaman wiki.
- **Admin Proyek/Project Lead** — mengatur siapa yang dapat menambah/mengedit/menghapus comment dan attachment lewat permission scheme dan role, melihat panel Automation.
- **Admin Situs** — mengaktifkan attachment secara global, menetapkan batas ukuran dan kuota, memilih mode Comment visibility.
- **Stakeholder/Viewer** — membaca comment yang boleh dilihat, mengunduh attachment, membuka permalink dari email.
- **Pengguna Tim Bisnis** — memakai quick reply chip, melampirkan dokumen Office dan melihatnya lewat Media Viewer.
- **Integrator/Developer API** — membuat comment/attachment lewat REST, menyimpan comment properties, mengonsumsi webhook dan endpoint changelog.
- **Agent Layanan** — melihat tab Approvals pada service project dan memahami perbedaan comment internal pada service project.
- **Tamu/Kolaborator Eksternal** — menerima tautan Share dan berkomentar sesuai akses yang diberikan.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin menambahkan comment rich text dengan mention dan screenshot langsung dari issue view agar diskusi teknis terdokumentasi di tempat pekerjaan berada (CMT-001, CMT-011).
- Sebagai Developer/Anggota Tim, saya ingin membalas comment tertentu dalam thread agar percakapan paralel di satu work item tetap mudah diikuti (CMT-005).
- Sebagai Admin Proyek/Project Lead, saya ingin membatasi comment sensitif hanya untuk role atau group tertentu agar informasi internal tidak terlihat oleh viewer lain di UI, API, notifikasi, maupun export (CMT-004).
- Sebagai Scrum Master/Team Lead, saya ingin menyematkan comment keputusan di atas daftar comment agar anggota tim tidak perlu menggulir panjang untuk menemukan konteks penting (CMT-007).
- Sebagai Stakeholder/Viewer, saya ingin membuka tautan langsung ke sebuah comment dari email agar saya langsung melihat bagian diskusi yang relevan (CMT-008).
- Sebagai Pengguna Tim Bisnis, saya ingin melampirkan file dengan drag-and-drop atau paste lalu melihatnya tanpa mengunduh agar review dokumen dan gambar cepat dilakukan (CMT-011, CMT-014, CMT-015).
- Sebagai Admin Situs, saya ingin mengatur batas ukuran file, jumlah attachment, dan kuota storage agar biaya penyimpanan terkendali dan instance tetap stabil (CMT-016, CMT-019).
- Sebagai Admin Situs, saya ingin file yang diunggah disajikan secara aman dan dapat dipindai antivirus agar attachment tidak menjadi vektor XSS atau malware (CMT-017, CMT-019).
- Sebagai Product Owner/PM, saya ingin melihat siapa mengubah apa dan kapan pada work item beserta nilai lama dan barunya agar saya dapat menelusuri keputusan dan perubahan scope (CMT-021, CMT-022).
- Sebagai Integrator/Developer API, saya ingin changelog tersimpan append-only dengan nilai teks historis dan dapat diambil per halaman agar laporan cycle time dan audit eksternal akurat (CMT-023).
- Sebagai Admin Proyek/Project Lead, saya ingin melihat eksekusi rule otomasi terbaru pada sebuah work item agar saya cepat mendiagnosis mengapa field berubah atau rule gagal (CMT-026).
- Sebagai Product Owner/PM, saya ingin membagikan work item ke rekan dengan pesan singkat dan menautkan halaman dokumentasi agar stakeholder memiliki konteks lengkap tanpa mencari sendiri (CMT-028, CMT-029, CMT-030).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 10 |
| P1 | 11 |
| P2 | 7 |
| P3 | 4 |
| **Total** | **32** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 10 |
| Fase 2 | 10 |
| Fase 3 | 7 |
| Fase 4 | 4 |
| Fase 5 | 0 (perluasan JSM pada CMT-025) |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 membangun tiga fondasi yang dipakai hampir semua modul lain, yaitu model changelog append-only (CMT-023) yang ditulis dalam transaksi yang sama dengan setiap perubahan work item, abstraksi file storage dengan driver local dan S3-compatible (CMT-019), serta CRUD comment dan attachment dasar dengan penegakan permission di server (CMT-001 s.d. CMT-003, CMT-011 s.d. CMT-013). Section Activity dengan tab Comments dan History (CMT-021, CMT-022) dirilis di Fase 1 agar pengguna langsung melihat jejak perubahan. Fase 2 menambah kemampuan yang bergantung pada konfigurasi dan role: visibility restriction, thread reply, permalink, urutan dan paginasi, panel attachment, Media Viewer, pengaturan batas, tab Work log, dan Share. Fase 3 menambah fitur kolaborasi sekunder (reaction, pinned comment), pengerasan keamanan file, siklus hidup attachment pada operasi work item, panel Automation, web link bergaya smart link, serta tab Approvals (CMT-025) yang dirilis bersamaan dengan approval project bisnis (BIZ-042, BIZ-046) lalu diperluas untuk approval Service Management di Fase 5. Fase 4 berisi integrasi wiki, embed video, dan anotasi gambar.

## Fitur

## Komentar

### CMT-001 Add Comment

- **Deskripsi:** Menambahkan comment rich text pada work item lewat comment editor di section Activity, transition screen, bulk change, REST API, otomasi, dan kanal lain.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan project permission Add Comments (PERM-030); Integrasi API dan Automation memakai permission actor-nya (PERM-049).
- **Aturan perilaku (acceptance criteria):**
  1. Section Activity menampilkan kotak placeholder "Add a comment..." di atas (urutan Newest first) atau di bawah (Oldest first) daftar comment, disertai pro tip "Pro tip: press M to comment"; shortcut `m` memfokuskan editor kecuali saat fokus berada di input teks.
  2. ⚠ Di bawah kotak placeholder tampil chip saran comment cepat dengan teks default "Looks good!", "Need help?", "This is blocked...", "Can you clarify...?", dan "This is on track"; daftar chip disimpan sebagai konfigurasi aplikasi agar teks dapat diganti tanpa rilis kode.
  3. ⚠ Chip quick reply hanya tampil selama editor belum aktif; klik chip membuka editor berisi teks template (boleh mengandung emoji) yang masih dapat diedit sebelum Save; chip disembunyikan bagi pengguna tanpa Add Comments; teks chip dilokalisasi mengikuti bahasa pengguna (EN/ID).
  4. Comment editor mendukung semua elemen ADF yang diizinkan untuk comment (formatting, mention, emoji, media inline, table, code block, smart link) sesuai EDT-001; elemen yang tidak diizinkan dibuang di server saat validasi.
  5. Comment disimpan dengan tombol Save atau Ctrl+Enter (Windows/Linux) / Cmd+Enter (macOS); tombol Cancel menutup editor dan membuang isinya setelah konfirmasi bila ada isi; tombol Save dinonaktifkan selama request berjalan untuk mencegah duplikasi.
  6. ⚠ Draft comment yang belum disimpan dipertahankan di penyimpanan lokal browser per pengguna dan per work item; saat editor dibuka kembali draft dipulihkan, dan draft dihapus setelah Save atau Cancel berhasil.
  7. Pemilih visibility (CMT-004) tersedia di editor sebelum Save; nilai default adalah tanpa restriksi.
  8. Comment yang kosong atau hanya berisi whitespace tidak dapat disimpan: tombol Save nonaktif di UI dan API mengembalikan HTTP 400 dengan pesan "Comment body tidak boleh kosong".
  9. Server menolak comment dari pengguna tanpa Add Comments dengan HTTP 403; editor dan placeholder tidak dirender bagi pengguna tersebut.
  10. Bila preferensi personal Autowatch aktif (modul 21), penulis comment otomatis ditambahkan sebagai watcher bila belum menjadi watcher.
  11. Setiap comment baru menerbitkan event domain "Issue commented" (modul 28) yang dikonsumsi notifikasi (modul 19), webhook `comment_created` (modul 23), dan trigger Automation (modul 20).
  12. Bila transition screen memuat field Comment (modul 09), comment yang diisi disimpan dalam transaksi yang sama dengan transisi; bila transisi gagal validasi, comment tidak tersimpan.
  13. ⚠ Wizard Bulk change (modul 06) menyediakan input comment opsional pada operasi bulk edit dan bulk transition; comment ditambahkan ke setiap work item yang berhasil diproses dan menerbitkan event per work item.
  14. Comment dapat dibuat lewat REST API (modul 23), aksi Automation "Add comment", dan aplikasi mobile; semua kanal memakai layanan create comment yang sama sehingga validasi, permission, dan event identik.
  15. ⚠ Untuk project non-service, membalas email notifikasi tidak membuat comment pada Fase 1; pembuatan comment dari email reply hanya tersedia lewat mail handler service project (modul 30, JSM-PRT-041) dan menjadi keputusan terbuka untuk project lain.
  16. Setiap comment ditampilkan dengan avatar dan nama penulis, waktu relatif (mis. "2 hours ago" / "2 jam lalu") dengan tooltip waktu absolut sesuai zona waktu dan format tanggal pengguna.
  17. Comment pada work item berstatus akhir (kategori Done) tetap diizinkan, kecuali company-managed workflow property `jira.permission.comment` (mis. `jira.permission.comment.denied`, `jira.permission.comment.group`) pada status tersebut membatasinya (PERM-037); pembatasan diterapkan di server.
  18. Satu work item menampung maksimal 5.000 comment; comment ke-5.001 ditolak dengan HTTP 400 dan pesan "Batas 5.000 comment per work item telah tercapai", dan editor menampilkan pesan yang sama.
- **Team-managed vs company-managed:** Company-managed mengatur Add Comments lewat permission scheme dan dapat dibatasi workflow property per status; team-managed mengatur lewat role bawaan/kustom (PERM-039, PERM-040) tanpa workflow property.
- **Catatan & ketergantungan:** Bergantung pada EDT-001 (editor), PERM-030, PERM-037, modul 19, 20, 23, 28. Tabel baru `comments` (id, issue_id, author_id, update_author_id, body_adf JSONB, created_at, updated_at, visibility_type, visibility_value, parent_comment_id, pinned_at, pinned_by); penghapusan comment bersifat permanen; satu-satunya pengecualian adalah comment induk yang masih memiliki reply (CMT-003 butir 8), yang isinya dikosongkan dan ditandai kolom `deleted_at` agar placeholder thread dapat dirender. Pesan galat memakai `constants.GetError` EN/ID yang sudah ada di repo.

### CMT-002 Edit Comment

- **Deskripsi:** Mengubah isi dan visibility comment yang sudah dipublikasikan sesuai permission.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Penulis comment dengan Edit Own Comments; pengguna dengan Edit All Comments untuk comment siapa pun (PERM-030).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Edit" di bawah comment membuka comment editor inline berisi konten lama; Cancel mengembalikan tampilan semula tanpa perubahan.
  2. Aksi Edit hanya dirender bila pengguna memiliki Edit Own Comments (untuk comment miliknya) atau Edit All Comments; server menolak edit tanpa permission dengan HTTP 403.
  3. Comment yang pernah diedit menampilkan label "Edited" di samping waktu, dengan tooltip waktu edit terakhir dan nama pengedit bila berbeda dari penulis.
  4. Pemilih visibility dapat diubah saat edit; pengguna hanya dapat memilih role/group yang ia ikuti (CMT-004) dan perubahan berlaku segera di semua jalur baca.
  5. Edit menerbitkan event "Issue comment edited" yang dikonsumsi notifikasi, webhook `comment_updated`, dan trigger Automation "Comment edited".
  6. UI tidak menyediakan riwayat versi isi comment, dan edit comment tidak menghasilkan entri di tab History.
  7. Company-managed workflow property `jira.permission.commentEditOwn` / `jira.permission.commentEditAll` (dengan suffix `.denied`, `.user`, `.group`, atau `.projectrole`) dapat membatasi edit comment pada status tertentu; pembatasan diterapkan di server dan menyembunyikan aksi Edit (PERM-037).
  8. Setiap edit memperbarui field `updateAuthor` (pengguna pengedit) dan `updated` (timestamp) pada entity comment; `author` dan `created` tidak berubah.
- **Team-managed vs company-managed:** Workflow property hanya berlaku di company-managed; team-managed memakai permission role.
- **Catatan & ketergantungan:** Bergantung pada CMT-001, CMT-004, PERM-030, PERM-037. Edit tidak mengubah urutan comment (urutan berdasarkan `created`).

### CMT-003 Delete Comment

- **Deskripsi:** Menghapus comment secara permanen sesuai permission.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Penulis comment dengan Delete Own Comments; pengguna dengan Delete All Comments (PERM-030).
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Delete" pada comment membuka dialog konfirmasi "Hapus comment ini? Tindakan ini tidak dapat dibatalkan." dengan tombol Delete dan Cancel; comment hanya dihapus setelah konfirmasi.
  2. Aksi Delete hanya dirender bila pengguna memiliki Delete Own Comments (comment miliknya) atau Delete All Comments; server menolak tanpa permission dengan HTTP 403.
  3. Penghapusan menerbitkan event "Issue comment deleted" serta webhook `comment_deleted` yang memuat id comment dan id work item.
  4. ⚠ Attachment yang disisipkan di comment tetap tersimpan sebagai attachment work item setelah comment dihapus dan tetap tampil di panel Attachments (CMT-014).
  5. ⚠ Penghapusan comment tidak menghasilkan entri di tab History; jejak penghapusan hanya tercatat di audit log aplikasi (modul 22) untuk keperluan admin.
  6. Company-managed workflow property `jira.permission.commentDeleteOwn` / `jira.permission.commentDeleteAll` dapat membatasi delete comment per status (PERM-037); aksi Delete disembunyikan dan server menolak dengan HTTP 403.
  7. Notifikasi email yang sudah terkirim tidak ditarik kembali; notifikasi in-app yang merujuk comment terhapus menampilkan teks "Comment ini telah dihapus" saat dibuka.
  8. ⚠ Saat comment induk yang memiliki reply dihapus, reply dipertahankan dan comment induk diganti placeholder "Comment ini telah dihapus" agar thread tetap terbaca (isi `body`, reaction, dan properties induk dihapus permanen, baris induk hanya menyimpan `deleted_at`); bila tidak ada reply, baris comment dihapus sepenuhnya. Placeholder tidak dihitung dalam batas 5.000 comment dan otomatis dihapus saat reply terakhirnya dihapus.
- **Team-managed vs company-managed:** Workflow property hanya berlaku di company-managed.
- **Catatan & ketergantungan:** Bergantung pada CMT-005 (thread), PERM-030, modul 22 (audit log). Placeholder induk terhapus adalah keputusan produk yang perlu dikonfirmasi (lihat Pertanyaan terbuka).

### CMT-004 Comment Visibility Restrictions

- **Deskripsi:** Membatasi visibilitas comment hanya untuk project role atau group tertentu, berbeda dari issue security level yang membatasi seluruh work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Add Comments memilih restriksi; Admin Situs mengatur mode global Comment visibility di General configuration (modul 22).
- **Aturan perilaku (acceptance criteria):**
  1. Comment editor memiliki dropdown visibility dengan ikon gembok berisi opsi "Tanpa restriksi" dan daftar role/group yang dapat dipilih.
  2. Daftar pilihan berisi project role tempat pengguna menjadi anggota pada project work item tersebut (mis. Administrators, Developers, Member).
  3. Pilihan group hanya muncul bila setting global "Comment visibility" di General configuration bernilai "Groups & Project roles"; nilai alternatif "Project roles only" (default) hanya menampilkan role.
  4. Pengguna hanya dapat memilih role/group yang ia ikuti; API yang mengirim role/group di luar keanggotaan penulis ditolak dengan HTTP 400 "Anda bukan anggota role/group yang dipilih".
  5. Comment yang dibatasi tampil dengan ikon gembok dan label "Restricted to <nama role/group>" bagi pengguna yang boleh melihatnya.
  6. Viewer di luar role/group tidak melihat comment sama sekali, termasuk di Activity, hasil pencarian, notifikasi, REST API, webhook payload untuk penerima yang tidak berhak, dan export; filter visibility diterapkan di repository layer untuk semua jalur baca.
  7. Notifikasi in-app dan email untuk comment yang dibatasi hanya dikirim ke penerima yang memenuhi restriksi.
  8. Pencarian JQL `comment ~ "teks"` hanya mencocokkan comment yang boleh dilihat pengguna yang menjalankan kueri (modul 12).
  9. ⚠ Team-managed project menawarkan restriksi ke role project (Administrator, Member, Viewer, dan role kustom); pilihan group tidak ditampilkan di team-managed walaupun mode global mengizinkan group, selaras dengan PERM-045.
  10. Worklog memakai mekanisme visibility yang sama (role/group) dengan aturan pemilihan dan penyaringan yang identik (CMT-024, PERM-045).
  11. Service project memakai konsep berbeda, yaitu "Internal note" vs "Reply to customer" (modul 29); dropdown role/group tidak dipakai untuk membedakan comment internal dari publik di service project.
- **Team-managed vs company-managed:** Company-managed memakai project role dari scheme dan group global (bila mode global mengizinkan); team-managed hanya memakai role project bawaan dan kustom.
- **Catatan & ketergantungan:** Bergantung pada PERM-045, PERM-018/PERM-019 (role), modul 22. Saat group yang dipakai restriksi dihapus, admin memilih group pengganti (swap group) atau visibilitas menjadi terbatas untuk admin saja (USR-015 butir 3); user/group nonaktif dalam grant mengikuti PERM-054.

### CMT-005 Reply to Comment (Threaded Replies) and Quoting

- **Deskripsi:** Membalas comment tertentu sehingga balasan tampil bersarang (thread) di bawah comment induk, serta mengutip isi comment memakai blockquote.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Add Comments; reply tunduk pada permission edit/delete yang sama dengan comment biasa.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Reply" pada comment membuka editor balasan tepat di bawah comment induk dengan mention penulis comment induk sudah terisi (dapat dihapus).
  2. Balasan tampil sebagai thread satu tingkat di bawah comment induk; reply terhadap reply disimpan ke induk tingkat atas yang sama (tidak ada nesting lebih dari satu tingkat). Thread berlaku di project software maupun bisnis, team-managed maupun company-managed.
  3. Thread berlaku untuk semua project dan tidak memiliki toggle untuk mematikannya per site atau per project.
  4. ⚠ Thread reply tidak tersedia di service project pada fase ini; service project tetap memakai daftar comment datar (modul 29) sampai diputuskan lain.
  5. ⚠ REST API publik v1 produk ini menyediakan field opsional `parentId` pada create comment agar integrasi dapat membuat reply; perilaku ini melampaui paritas dan didokumentasikan sebagai ekstensi di modul 23.
  6. Reply menotifikasi penulis comment induk dan pengguna yang di-mention, selain penerima notifikasi "Issue commented" standar sesuai notification scheme.
  7. ⚠ Reply pada comment induk yang dibatasi mewarisi restriksi visibility induk secara default; penulis tidak dapat memilih visibility yang lebih longgar daripada induk.
  8. ⚠ Thread dengan lebih dari 3 balasan ditampilkan collapsed dengan tautan "Show N more replies"; pengguna dapat expand/collapse, dan jumlah balasan tampil di comment induk.
  9. ⚠ Tidak ada aksi "Quote comment" khusus; pengguna menyalin teks lalu memakai elemen blockquote di editor (EDT-001).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model data: kolom `parent_comment_id` (nullable, FK ke `comments.id`) pada tabel `comments`. Perilaku reply saat induk dihapus diatur di CMT-003 butir 8. Batas 5.000 comment per work item menghitung reply.

### CMT-006 Comment Reactions

- **Deskripsi:** Memberi reaksi emoji pada comment sebagai umpan balik ringan tanpa menulis comment baru.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Pengguna yang dapat melihat comment (Browse Projects dan lolos visibility); tidak memerlukan Add Comments.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol "Add reaction" (ikon emoji) tampil di bawah setiap comment.
  2. ⚠ Picker menampilkan set quick reaction default (jempol, tepuk tangan, api, hati, kaget, berpikir) dan tautan ke emoji picker lengkap (modul 07).
  3. Setiap reaksi menampilkan penghitung; klik reaksi yang sudah ada menambah reaksi pengguna, klik lagi mencabutnya; satu pengguna hanya dapat memberi satu reaksi per emoji per comment.
  4. Hover pada reaksi menampilkan daftar nama pengguna yang memberi reaksi tersebut (maksimal 10 nama lalu "dan N lainnya").
  5. ⚠ Reaksi tidak mengirim notifikasi email maupun in-app kepada penulis comment.
  6. Reaksi pada comment yang dibatasi hanya terlihat oleh viewer yang memenuhi restriksi comment tersebut, termasuk lewat API.
  7. Menambah atau mencabut reaksi tidak tercatat di tab History.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model data: tabel `comment_reactions` (comment_id, user_id, emoji_id, created_at) dengan unique (comment_id, user_id, emoji_id). Reaksi terhapus bersama comment.

### CMT-007 Pinned Comments

- **Deskripsi:** Menyematkan comment penting agar tampil di area khusus di atas daftar comment pada section Activity.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Add Comments dan Admin Proyek; aturan siapa boleh pin ditentukan di butir 4.
- **Aturan perilaku (acceptance criteria):**
  1. Menu "..." (more actions) pada comment memuat aksi "Pin" untuk comment yang belum di-pin dan "Unpin" untuk comment yang sudah di-pin.
  2. Comment yang di-pin tampil di area "Pinned" di atas daftar comment dengan tautan "Lihat di timeline" yang menggulir ke posisi asli comment.
  3. ⚠ Maksimal 3 comment dapat di-pin per work item; percobaan pin ke-4 ditolak dengan pesan "Maksimal 3 comment dapat disematkan; lepaskan salah satu terlebih dulu".
  4. ⚠ Pin/unpin dapat dilakukan oleh pengguna dengan Add Comments pada work item tersebut; pengguna tanpa Add Comments tidak melihat aksi Pin.
  5. Comment yang dibatasi visibilitasnya tetap mengikuti restriksi saat di-pin; viewer di luar restriksi tidak melihatnya di area Pinned maupun di jumlah pinned.
  6. Edit comment yang di-pin langsung memperbarui tampilan di area Pinned; delete comment otomatis menghapusnya dari area Pinned.
  7. Area Pinned selalu tampil di atas dan tidak terpengaruh toggle urutan Newest/Oldest first; urutan di area Pinned mengikuti waktu pin terbaru di atas.
  8. ⚠ Pin/unpin tidak tercatat di tab History; kolom `pinned_by` dan `pinned_at` cukup sebagai jejak.
  9. ⚠ Fitur tersedia di team-managed dan company-managed project; di service project pin tersedia hanya untuk agent dan hanya pada comment internal.
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission (scheme vs role).
- **Catatan & ketergantungan:** Model data: kolom `pinned_at` dan `pinned_by` pada `comments`. Batas jumlah pin adalah keputusan terbuka.

### CMT-008 Comment Permalink

- **Deskripsi:** Tautan langsung ke comment tertentu yang memfokuskan dan menyorot comment itu saat dibuka.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat comment.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap comment memiliki aksi "Copy link to comment" (ikon link) yang menyalin URL absolut ke clipboard dan menampilkan flag "Tautan disalin".
  2. Format URL adalah `/browse/KEY-1?focusedCommentId=<id>`; untuk kompatibilitas, parameter `focusedId` juga diterima.
  3. Saat URL dibuka, issue view memaksa tab Comments (atau All) aktif, menggulir ke comment tersebut, dan menyorotnya dengan latar berwarna selama beberapa detik.
  4. Bila comment target belum dimuat (berada di balik "Load more"), klien memuat halaman yang memuat comment tersebut lalu memfokuskannya; reply di dalam thread collapsed ikut membuka thread-nya.
  5. Permalink ke comment yang dihapus atau tidak boleh dilihat membuka work item tanpa sorotan dan tanpa pesan yang membocorkan keberadaan comment.
  6. Email notifikasi comment (modul 19) memakai format tautan yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Rute SPA `/browse/:key` di-fallback ke `index.html` oleh `webui`; parameter dibaca di klien. Bergantung pada CMT-009 untuk pemuatan bertahap.

### CMT-009 Comment Sorting and Pagination

- **Deskripsi:** Pengurutan dan pemuatan bertahap comment serta entri activity lain.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Semua pengguna yang dapat melihat work item.
- **Aturan perilaku (acceptance criteria):**
  1. Section Activity memiliki toggle urutan "Newest first" / "Oldest first"; default Newest first.
  2. Pilihan urutan disimpan sebagai preferensi per pengguna di server (modul 21) dan berlaku di semua work item dan perangkat.
  3. Comment dimuat bertahap: halaman awal memuat 20 comment pertama sesuai arah urutan aktif (Newest first: 20 terbaru; Oldest first: 20 terlama) lalu tombol "Load more" / "Show N more" memuat halaman berikutnya; API memakai paginasi `startAt`/`maxResults` (maksimal 100 per halaman).
  4. ⚠ Comment yang tingginya melebihi ambang tampilan (sekitar 20 baris) di-collapse dengan tautan "Show more" dan dapat di-expand atau di-collapse lagi dengan "Show less".
  5. Toggle urutan yang sama berlaku di tab All dan tab History, sehingga berganti tab tidak mengubah arah urutan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CMT-021. Index database `(issue_id, created_at)` pada `comments`.

### CMT-010 Comment Events, Properties and API Surface

- **Deskripsi:** Aspek integrasi comment yang memengaruhi model data: event, comment properties, batasan, pencarian, dan representasi entity di API.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Integrator/Developer API, Automation, Admin Situs; akses baca/tulis mengikuti permission comment dan visibility.
- **Aturan perilaku (acceptance criteria):**
  1. Sistem menerbitkan event domain `comment_created`, `comment_updated`, dan `comment_deleted` untuk webhook, serta trigger Automation "Work item commented" dan "Comment edited" (modul 20, 23).
  2. Entity comment di API memuat `id`, `author`, `updateAuthor`, `body` (ADF pada API v3-compatible, wiki markup pada v2-compatible), `renderedBody` (opsional via `expand=renderedBody`), `created`, `updated`, `visibility` (`type` = `role`/`group`, `value`/`identifier`), `jsdPublic` (hanya service project), dan `properties`.
  3. Comment properties berupa pasangan key-value JSON (key maksimal 255 karakter, value maksimal 32 KB) yang dapat dibuat, dibaca, dan dihapus oleh app dan integrasi dengan permission yang sama seperti edit comment.
  4. Batas 5.000 comment per work item ditegakkan di layanan comment sehingga berlaku untuk semua kanal (UI, API, Automation, email, bulk).
  5. JQL `comment ~ "teks"` mencari teks comment melalui index full-text dan menghormati visibility (modul 12).
  6. Comment dari Automation ditulis atas nama actor rule yang dikonfigurasi; default actor adalah akun sistem "Automation" yang ditampilkan dengan avatar khusus (PERM-049).
  7. Comment ikut di export Word/XML/print (modul 13) hanya bila pengguna pengekspor memenuhi visibility comment tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail endpoint REST di modul 23; katalog event di modul 28. Tabel `comment_properties` (comment_id, key, value JSONB).

## Lampiran

### CMT-011 Upload Attachments

- **Deskripsi:** Melampirkan file ke work item lewat berbagai cara upload (tombol, drag-and-drop, paste, create/transition screen, editor, API, email).
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan permission Create Attachments (PERM-031); Integrator lewat API dengan permission yang sama.
- **Aturan perilaku (acceptance criteria):**
  1. Tombol attach (ikon paperclip) di toolbar issue view dan item "Attachment" di menu "+ Add" membuka pemilih file sistem operasi.
  2. Drag-and-drop file ke area mana pun di issue view menampilkan overlay "Drop files to attach"; melepas file memulai upload, dan melepas di luar overlay membatalkan.
  3. Paste file atau screenshot dari clipboard ke issue view atau editor mengunggah file; screenshot tanpa nama diberi nama `image-<yyyyMMdd-HHmmss>.png`.
  4. Beberapa file dapat diunggah sekaligus; setiap file menampilkan progress bar sendiri dan tombol cancel yang membatalkan upload file itu saja tanpa memengaruhi file lain.
  5. Create dialog menampilkan field Attachment bila field ada di create screen/layout; file diunggah sementara dan dikaitkan ke work item setelah create berhasil, lalu file sementara yang tidak terpakai dibersihkan.
  6. Transition screen yang memuat field Attachment (modul 09) menerima upload; attachment dikaitkan dalam transaksi yang sama dengan transisi.
  7. File yang disisipkan lewat editor description/comment otomatis menjadi attachment work item dan dirujuk oleh node media ADF.
  8. Upload tersedia lewat REST API multipart (header `X-Atlassian-Token: no-check` diterima untuk kompatibilitas), aplikasi mobile, dan email request service project (modul 30).
  9. Upload memerlukan Create Attachments dan setting global attachment aktif (CMT-016); bila salah satu tidak terpenuhi, kontrol upload disembunyikan dan API mengembalikan HTTP 403.
  10. Nama file duplikat diizinkan dan disimpan sebagai attachment terpisah dengan id berbeda.
  11. Upload ditolak dengan pesan galat spesifik bila ukuran melebihi batas ("File melebihi batas <N> MB"), kuota storage habis ("Kuota penyimpanan instance telah habis"), atau batas 2.000 attachment per work item tercapai ("Batas 2.000 attachment per work item telah tercapai").
  12. Upload berhasil menerbitkan webhook `attachment_created` dan mencatat entri History field "Attachment" dengan nilai baru berupa nama file.
- **Team-managed vs company-managed:** Company-managed dapat menyembunyikan field Attachment lewat field configuration; team-managed mengatur lewat layout work type.
- **Catatan & ketergantungan:** Bergantung pada CMT-019 (storage), CMT-016 (batas), CMT-023 (history), PERM-031. Tabel baru `attachments` (id, issue_id, author_id, filename, mime_type, size_bytes, storage_key, sha256, scan_status, thumbnail_key, created_at).

### CMT-012 Download Attachments

- **Deskripsi:** Mengunduh attachment satu per satu atau semuanya sekaligus sebagai ZIP.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects yang lolos issue security level (PERM-044).
- **Aturan perilaku (acceptance criteria):**
  1. Download per file tersedia dari hover thumbnail panel, baris list view, dan header Media Viewer.
  2. Aksi "Download all" di menu "..." panel Attachments mengunduh semua attachment work item sebagai satu file ZIP bernama `<KEY>-attachments.zip`; aksi ini boleh dirilis pada Fase 2.
  3. Nama file asli dipertahankan pada download; nama yang bentrok di dalam ZIP diberi suffix ` (1)`, ` (2)`, dan seterusnya sebelum ekstensi.
  4. Download memerlukan Browse Projects dan lolos issue security level; permintaan tanpa akses mengembalikan HTTP 404 agar keberadaan file tidak bocor.
  5. URL download memakai token sementara (signed URL berumur pendek) atau redirect ke endpoint media terpisah; URL kedaluwarsa mengembalikan HTTP 403 dan klien meminta URL baru.
  6. Data security policy (PERM-065) yang memblokir download menyembunyikan aksi download dan menolak permintaan di server dengan pesan "Download diblokir oleh kebijakan keamanan data".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CMT-019. ZIP dibuat secara streaming agar tidak menahan seluruh file di memori.

### CMT-013 Delete Attachments

- **Deskripsi:** Menghapus attachment dari work item secara permanen sesuai permission.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pemilik attachment dengan Delete Own Attachments; pengguna dengan Delete All Attachments (PERM-031).
- **Aturan perilaku (acceptance criteria):**
  1. Tombol Delete (ikon trash) tampil saat hover thumbnail atau baris list dan membuka dialog konfirmasi "Hapus <nama file>? Tindakan ini tidak dapat dibatalkan."
  2. Aksi Delete hanya dirender bila pengguna memiliki Delete Own Attachments (attachment yang ia unggah) atau Delete All Attachments; server menolak tanpa permission dengan HTTP 403.
  3. Penghapusan bersifat permanen tanpa trash atau restore; objek di storage dihapus oleh job asinkron setelah record database dihapus.
  4. Node media di description/comment yang merujuk file terhapus dirender sebagai placeholder "File tidak tersedia" tanpa merusak konten lain.
  5. Penghapusan mencatat entri History field "Attachment" (nama file -> kosong) dan menerbitkan webhook `attachment_deleted`.
  6. ⚠ UI tidak menyediakan bulk delete attachment; penghapusan massal hanya lewat REST API satu per satu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CMT-019 (cleanup), CMT-023 (history), PERM-031.

### CMT-014 Attachment Panel Views (Strip and List)

- **Deskripsi:** Panel Attachments di issue view dengan tampilan strip (thumbnail) dan list (tabel).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Panel "Attachments (N)" hanya muncul bila work item memiliki minimal satu attachment, dengan badge jumlah attachment yang dapat dilihat pengguna.
  2. Strip view menampilkan thumbnail horizontal dengan scroll serta tombol prev/next bila thumbnail melebihi lebar panel.
  3. List view menampilkan tabel dengan kolom Name, Size, dan Date added; klik header kolom mengurutkan berdasarkan kolom tersebut.
  4. ⚠ Toggle strip/list tersedia di menu "..." panel dan pilihan disimpan sebagai preferensi per pengguna yang berlaku di semua work item.
  5. Menu "..." menyediakan sort berdasarkan Name atau Date added, ascending atau descending.
  6. Thumbnail otomatis dibuat untuk gambar, PDF (halaman pertama), dan video (frame awal); tipe lain memakai ikon sesuai kategori tipe file.
  7. Hover attachment menampilkan nama file, ukuran, tanggal ditambahkan, serta tombol Download dan Delete (sesuai permission).
  8. Attachment yang diunggah lewat comment atau description ikut tampil di panel.
  9. ⚠ Section Attachments terpadu (rilis Juli 2026) menampilkan dalam satu section berkas yang diunggah (CMT-011), dokumen tertaut (web/remote link dan halaman wiki, CMT-029, CMT-030, INT-043), dan video tertaut (embed video, CMT-032, INT-047); section menyediakan filter per tipe sumber (Uploaded, Linked documents, Linked videos, All) dan grid view di samping strip view dan list view, dengan badge jumlah yang menghitung item sesuai filter aktif, item tertaut yang tidak dapat diakses viewer ditampilkan terbatas tanpa judul, dan aksi Delete pada item tertaut berarti unlink (mengikuti izin Link Issues), bukan menghapus berkas sumber.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Thumbnail berasal dari job CMT-019 butir 8; selama job belum selesai panel menampilkan ikon tipe. Section terpadu (butir 9) bergantung pada CMT-029, CMT-030, dan CMT-032 di modul ini serta INT-043 dan INT-047 di [25 Integrasi, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md); bila integrasi wiki/video belum aktif, filter tipe tersebut disembunyikan.

### CMT-015 Media Viewer (Attachment Preview)

- **Deskripsi:** Viewer fullscreen untuk melihat attachment tanpa mengunduhnya; produk ini memakai preview native browser ditambah konversi sisi server bila tersedia.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Pengguna dengan Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. Klik attachment di panel atau media inline di editor membuka overlay Media Viewer.
  2. Preview gambar mendukung PNG, JPG, GIF animasi, WebP, SVG (dirender sebagai gambar tanpa eksekusi script), dan HEIC (dikonversi ke JPEG oleh job), dengan kontrol zoom in, zoom out, dan fit to screen.
  3. ⚠ Preview PDF menampilkan navigasi halaman dan pencarian teks memakai viewer PDF berbasis browser.
  4. Preview video (MP4, MOV, WebM) dan audio (MP3, WAV, OGG) memakai player bawaan; format yang tidak dapat diputar browser ditranskode oleh job asinkron bila transcoder diaktifkan.
  5. ⚠ Preview dokumen Office (DOCX, XLSX, PPTX) ditampilkan dari hasil konversi ke PDF oleh layanan konversi opsional; bila layanan tidak aktif, tampil "Preview unavailable".
  6. ⚠ Preview teks/kode (TXT, JSON, CSV, LOG, MD, dan ekstensi kode umum) dirender sebagai teks monospace dengan syntax highlight untuk bahasa yang dikenali, dibatasi 1 MB pertama.
  7. Tombol panah kiri/kanan dan tombol keyboard panah berpindah antar attachment work item sesuai urutan panel.
  8. Header viewer menampilkan nama file, tombol Download, dan Close; tombol Esc menutup viewer.
  9. File yang tidak dapat di-preview menampilkan "Preview unavailable" dengan tombol Download.
  10. File baru atau besar yang sedang diproses menampilkan status "Processing..." dan viewer memperbarui tampilan otomatis setelah pemrosesan selesai.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CMT-019 (thumbnail, MIME, scan). File berstatus scan infected tidak dapat dibuka (CMT-017).

### CMT-016 Attachment Settings and Limits

- **Deskripsi:** Konfigurasi global attachment serta batas ukuran file, jumlah attachment per work item, dan kuota storage.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (Administer Jira, PERM-007); org admin untuk kebijakan keamanan data.
- **Aturan perilaku (acceptance criteria):**
  1. Halaman Settings > System > Attachments menyediakan toggle untuk mengaktifkan/menonaktifkan attachment di seluruh site; saat nonaktif, attachment lama tetap dapat dilihat dan diunduh tetapi upload baru ditolak.
  2. ⚠ Admin mengatur batas ukuran maksimum per file; default produk ini 100 MB dengan batas atas konfigurasi 2 GB, dan perubahan batas tidak memengaruhi file yang sudah ada.
  3. Setiap work item menampung maksimal 2.000 attachment; batas ini dapat diturunkan admin tetapi tidak dapat dinaikkan melebihi 2.000.
  4. ⚠ Kuota storage per instance dapat dikonfigurasi (default tanpa batas pada self-hosted); halaman pengaturan menampilkan pemakaian saat ini dan pembanding paket (2 GB, 250 GB, tanpa batas) hanya sebagai referensi.
  5. ⚠ Tidak ada allowlist/blocklist tipe file secara default; tipe berisiko selalu disajikan sebagai download (CMT-017), dan admin dapat mengaktifkan blocklist ekstensi opsional.
  6. Data security policy (PERM-065) dapat memblokir download attachment atau akses publik ke attachment terlepas dari permission project.
  7. Company-managed field configuration (CFG-043) dapat menyembunyikan field Attachment sehingga panel dan kontrol upload tidak tampil di project terkait.
  8. Pengaturan thumbnail dan ZIP preview gaya Server/DC tidak disediakan sebagai opsi admin; thumbnail selalu aktif.
- **Team-managed vs company-managed:** Company-managed memakai field configuration untuk menyembunyikan field; team-managed memakai layout work type.
- **Catatan & ketergantungan:** Bergantung pada modul 22 (halaman admin) dan CMT-019 (penegakan kuota). Nilai default disimpan di tabel pengaturan global.

### CMT-017 Attachment Security and Malware Scanning

- **Deskripsi:** Perlindungan keamanan untuk file yang diunggah dan disajikan, mencakup isolasi domain, sanitasi tipe berisiko, pemindaian, enkripsi, dan kontrol akses.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Sistem; Admin Situs/org admin mengatur opsi pemindaian dan kebijakan.
- **Aturan perilaku (acceptance criteria):**
  1. File attachment disajikan dari origin/domain media yang terpisah dari domain aplikasi bila dikonfigurasi (mis. `media.<host>`), sehingga script dalam file tidak dapat mengakses cookie atau token sesi aplikasi.
  2. Content-Type berisiko (HTML, XHTML, SVG yang memuat script, JavaScript, XML) selalu disajikan dengan `Content-Disposition: attachment` atau disanitasi sebelum dirender inline.
  3. ⚠ Pemindaian malware bukan bagian dari paritas Jira; produk ini menyediakan pemindaian antivirus opsional (CMT-019 butir 7) yang dapat diaktifkan admin, dan file berstatus infected tidak dapat diunduh atau di-preview.
  4. File dienkripsi at-rest (enkripsi sisi server storage) dan in-transit (HTTPS/TLS untuk semua upload dan download).
  5. Akses attachment selalu mengikuti permission work item dan issue security level di setiap permintaan, termasuk thumbnail dan preview.
  6. Akses anonim ke attachment hanya diizinkan bila Browse Projects diberikan ke "Anyone on the web" (PERM-046) dan attachment tidak diblokir data security policy.
  7. Penambahan dan penghapusan attachment selalu tercatat di History work item dengan actor dan waktu.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Implementasi teknis di CMT-019; kebijakan akses di PERM-044, PERM-046, PERM-065. Integrasi ClamAV adalah nilai tambah clone.

### CMT-018 Attachment Lifecycle on Issue Operations

- **Deskripsi:** Perilaku attachment saat work item di-clone, dipindah, diarsip, diimpor, dicari, diekspor, atau dihapus.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna yang menjalankan operasi work item (modul 06); Admin Situs untuk import.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Dialog Clone work item (modul 06) memuat checkbox "Include attachments"; bila dicentang, attachment disalin sebagai record baru yang merujuk objek storage yang sama (copy-on-reference) dan tercatat sebagai attachment work item hasil clone.
  2. Move work item ke project lain memindahkan seluruh attachment tanpa mengubah id dan storage key, dan akses mengikuti permission project tujuan.
  3. Delete work item menghapus seluruh attachment-nya secara permanen; objek storage dihapus oleh job cleanup bila tidak dirujuk record lain.
  4. Archive work item atau project mempertahankan attachment, dan attachment kembali dapat diakses saat work item/project di-restore.
  5. Import CSV (modul 25) dapat menyertakan attachment lewat kolom URL; file diunduh oleh job import dan gagal per file dicatat di laporan import tanpa menggagalkan baris.
  6. JQL mendukung `attachments IS EMPTY` dan `attachments IS NOT EMPTY` (modul 12).
  7. Export work item (XML/Word/print) menyertakan daftar attachment (nama, ukuran, tanggal, pengunggah) tanpa menyematkan isi file.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada operasi Clone (ISS-031), Move (ISS-032), Delete (ISS-033), dan Archive (ISS-034) di modul 06 serta import di modul 25. Copy-on-reference memerlukan reference counting pada CMT-019.

### CMT-019 File Storage for Attachments & Avatars

- **Deskripsi:** Penyimpanan berkas lewat abstraksi object storage dengan keamanan, thumbnail, kuota, dan pengelolaan avatar.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; Admin Situs mengatur driver, batas, dan pemindaian; pengguna akhir mengunggah dan mengunduh.
- **Aturan perilaku (acceptance criteria):**
  1. Interface `FileStorage` di domain layer memiliki dua driver: local filesystem (dev/single node) dan S3-compatible (AWS S3, MinIO, Cloudflare R2) untuk produksi; driver dipilih lewat variabel lingkungan (mis. `STORAGE_DRIVER`, `STORAGE_BUCKET`).
  2. Upload dilakukan secara streaming lewat server dengan batas ukuran yang ditegakkan saat stream, atau lewat presigned URL langsung ke object storage yang dikonfirmasi oleh endpoint finalize.
  3. ⚠ Batas ukuran per file dapat dikonfigurasi dengan default 100 MB; batas atas yang dapat diatur admin adalah 2 GB.
  4. Download disajikan lewat signed URL berumur pendek (default 5 menit) atau proxy yang memeriksa Browse Projects dan issue security pada setiap permintaan.
  5. Respons file menyertakan `Content-Disposition: attachment` untuk tipe berisiko, inline hanya untuk gambar raster dan PDF yang aman, header `X-Content-Type-Options: nosniff`, dan disajikan dari origin terpisah untuk konten pengguna bila dikonfigurasi.
  6. Tipe MIME dideteksi dari magic bytes, bukan dari ekstensi atau header klien; ketidaksesuaian dicatat dan tipe hasil deteksi yang dipakai.
  7. Antivirus scan asinkron (mis. ClamAV) bersifat opsional dengan status `pending`, `clean`, dan `infected`; file `infected` diblokir untuk download/preview dan pengunggah serta admin menerima notifikasi.
  8. Thumbnail dan preview gambar (resize, strip metadata EXIF/GPS) dibuat sebagai background job; kegagalan job tidak menggagalkan upload.
  9. ⚠ Kuota storage per instance dapat dikonfigurasi dengan peringatan pada 80% dan 100% pemakaian; upload ditolak saat kuota penuh. Angka paket Jira Cloud (Free 2 GB, Standard 250 GB, Premium/Enterprise tanpa batas) hanya pembanding.
  10. Deduplikasi berbasis sha256 bersifat opsional; job cleanup berkala menghapus objek orphan (tidak dirujuk record mana pun) yang berumur lebih dari 24 jam.
  11. Avatar pengguna, project, dan work type diunggah dengan crop persegi dan disimpan dalam ukuran standar 16, 24, 32, dan 48 piksel (ditambah ukuran asli maksimum 512 piksel).
  12. Setting admin enable/disable attachments secara global (CMT-016) dan permission Create Attachments, Delete Own Attachments, Delete All Attachments (PERM-031) ditegakkan di layanan storage, bukan hanya di UI.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fondasi untuk CMT-011 s.d. CMT-018 dan avatar di modul 01/04/05. Konfigurasi dibaca lewat `internal/core/config/` sesuai pola repo; driver ditempatkan di `internal/shared/infrastructure/storage`. Job asinkron memakai mekanisme antrean di modul 27.

### CMT-020 Image Annotation (Media Editor)

- **Deskripsi:** Menganotasi gambar (panah, teks, bentuk, coretan) sebelum atau sesudah dilampirkan.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Create Attachments.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Media Viewer menampilkan tombol "Annotate" untuk attachment gambar bagi pengguna dengan Create Attachments.
  2. ⚠ Editor anotasi menyediakan alat panah, kotak, lingkaran, teks, brush, pemilih warna, dan pemilih ketebalan garis, dengan undo/redo.
  3. ⚠ Hasil anotasi disimpan sebagai attachment gambar baru (nama asli dengan suffix `-annotated`); gambar asli tidak diubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keberadaan fitur ini di Jira Cloud saat ini belum dapat dipastikan; jangan diimplementasikan sebelum dikonfirmasi pemilik produk. Bergantung pada CMT-015.

## Aktivitas & riwayat perubahan

### CMT-021 Activity Section and Tabs

- **Deskripsi:** Section Activity di issue view yang berisi tab untuk melihat comment, riwayat perubahan, worklog, dan approval.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects; isi tiap tab difilter sesuai visibility dan permission.
- **Aturan perilaku (acceptance criteria):**
  1. Tab "All" menampilkan gabungan comment, entri History, dan worklog dalam satu timeline kronologis dengan penanda jenis entri.
  2. Tab "Comments" adalah tab default saat issue view dibuka pertama kali.
  3. Tab "History" menampilkan riwayat perubahan field (CMT-022).
  4. Tab "Work log" menampilkan daftar worklog (CMT-024) dan hanya muncul bila time tracking aktif (CFG-058).
  5. Tab "Approvals" muncul pada work item di project yang memakai approval step (CMT-025).
  6. Tab tambahan dari app pihak ketiga (issue activity extension) adalah ekosistem dan tidak dibangun native; produk ini menyediakan titik ekstensi sesuai modul 25.
  7. ⚠ Tab yang terakhir dipilih disimpan sebagai preferensi per pengguna dan dipulihkan saat membuka work item lain.
  8. Toggle Newest/Oldest first (CMT-009) berlaku di semua tab dan dipertahankan saat berganti tab.
  9. Tab hanya dirender bila fiturnya aktif pada project/site (mis. Work log tersembunyi bila time tracking nonaktif, Approvals tersembunyi bila tidak ada approval); preferensi tab yang tidak tersedia jatuh kembali ke Comments.
  10. Bila fitur AI opsional aktif (ISS-045, modul 06), tombol "Summarize" di atas daftar comment menghasilkan ringkasan comment yang hanya memakai comment yang boleh dilihat pengguna; tombol tidak tampil bila AI nonaktif.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Automation bukan tab Activity melainkan panel terpisah (CMT-026). Bergantung pada CMT-009, CMT-022, CMT-024, CMT-025.

### CMT-022 History Tab and Change Diff

- **Alias inventaris:** History / Changelog
- **Deskripsi:** Rekam jejak lengkap semua perubahan pada work item yang ditampilkan di tab History: siapa yang mengubah, kapan, field apa, serta nilai lama dan barunya.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects; Sistem menulis entri. Tidak ada peran yang dapat mengedit atau menghapus entri.
- **Aturan perilaku (acceptance criteria):**
  1. Setiap entri menampilkan avatar dan nama actor, kalimat aksi (mis. "changed the Status", "updated the Description", "created the Work item"), dan waktu relatif dengan tooltip waktu absolut.
  2. Setiap entri perubahan field memuat nama field serta nilai lama dan nilai baru; nilai kosong ditampilkan sebagai "None".
  3. Nilai From -> To dirender sesuai tipe field: status sebagai lozenge berwarna kategori, user sebagai avatar dan nama, tanggal sesuai format pengguna, select/label/sprint sebagai teks/chip.
  4. Perubahan yang dicatat minimal: created, transition/status, edit field sistem dan custom, assignee, sprint, rank, issue link, attachment ditambah/dihapus, parent, security level, move (key dan project lama), serta worklog ditambah/diubah/dihapus.
  5. Perubahan attachment (tambah/hapus), issue link (tambah/hapus dengan jenis relasi), parent, sprint, dan rank (ditampilkan "Ranked higher" / "Ranked lower") memiliki kalimat aksi khusus.
  6. Entri pertama setiap work item adalah "created the Work item" dengan actor pembuat dan waktu create.
  7. Beberapa field yang berubah dalam satu operasi (mis. transisi dengan screen, bulk edit) dikelompokkan menjadi satu changelog group dengan satu header actor dan waktu.
  8. Perubahan watcher, edit comment, dan reaction tidak dicatat di History.
  9. Perubahan field rich text (Description, paragraph custom field) disimpan sebagai teks lama dan teks baru (plain text hasil render ADF).
  10. ⚠ Nilai rich text panjang dipotong dengan tautan "Show more"; tampilan diff berwarna (penanda tambah/hapus per kata) disediakan sebagai mode opsional "Show differences".
  11. Perubahan oleh app, integrasi, atau AI agent ditandai dengan nama dan avatar actor terkait (mis. akun app), bukan pengguna yang memicu secara tidak langsung.
  12. Perubahan oleh Automation memakai actor rule yang dikonfigurasi; default akun sistem "Automation" (PERM-049).
  13. Perubahan oleh pengguna yang sudah dihapus ditampilkan dengan nama "Former user" dan avatar default (USR-008).
  14. ⚠ Saat teks diredaksi (PERM-066, ISS-043), entri History yang memuat teks tersebut ikut diganti placeholder redaksi pada nilai lama maupun baru.
  15. Perubahan pada field yang tidak dapat dilihat pengguna (field tersembunyi, field dengan restriksi, security level yang tidak dapat diakses) disembunyikan dari tab History pengguna tersebut.
  16. UI tidak menyediakan aksi edit atau delete pada entri History untuk peran mana pun, termasuk Admin Situs.
  17. API tidak menyediakan endpoint untuk mengubah atau menghapus entri changelog; entri hanya bertambah (CMT-023).
  18. Tab History mengikuti urutan Newest first atau Oldest first sesuai preferensi pengguna (CMT-009).
  19. Entri History juga tampil di tab All bersama comment dan worklog dengan urutan yang sama.
  20. History panjang dimuat bertahap dengan tombol "Load more" (default 50 changelog group per halaman).
  21. REST API menyediakan endpoint changelog terpisah per work item (mis. `GET /api/v1/issues/{key}/changelog`) dengan paginasi `startAt`/`maxResults`, selain `expand=changelog` pada get work item.
  22. Representasi REST changelog berisi `id` group, `author`, `created`, dan `items` (masing-masing `field`, `fieldtype`, `fieldId`, `from`, `fromString`, `to`, `toString`).
  23. Operator JQL `WAS` dan `CHANGED` dievaluasi dari data changelog yang sama dengan tab History, bukan dari snapshot terpisah (modul 12).
  24. Riwayat perubahan dapat dicari lewat JQL `WAS`/`CHANGED` dengan predikat FROM/TO/BY/DURING/BEFORE/AFTER/ON untuk field yang didukung modul 12 (CMT-023 butir 3).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model data dan penulisan transaksional di CMT-023. Format tanggal dan zona waktu dari preferensi modul 21. Field yang disembunyikan mengikuti field configuration (CFG-043) dan layout.

### CMT-023 Changelog & Issue History Model

- **Deskripsi:** Model data dan aturan penulisan changelog yang menjadi sumber tab History, timeline activity, laporan agile, dan predikat JQL historis.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering; ditulis oleh layanan domain, dibaca oleh History, laporan, JQL, dan API.
- **Aturan perilaku (acceptance criteria):**
  1. Skema terdiri atas tabel `change_groups` (id, issue_id, author_id, created_at) dan `change_items` (group_id, field, field_type `jira`/`custom`, field_id, from_value, from_string, to_value, to_string).
  2. Setiap edit, transisi, link, attachment, worklog, move, rank, dan perubahan sprint menulis changelog dalam transaksi database yang sama dengan perubahan data; bila salah satu gagal, keduanya di-rollback.
  3. Index history dibuat pada (field, to_value, created_at), (field, from_value, created_at), dan (issue_id, created_at) untuk mendukung `WAS`, `WAS IN`, `WAS NOT`, `CHANGED` dengan FROM/TO/BY/DURING/BEFORE/AFTER/ON.
  4. Tabel turunan `status_transitions` (issue_id, from_status, to_status, transitioned_at, actor_id) diisi pada setiap transisi untuk Control chart, Cumulative flow diagram, cycle time, dan time in status (modul 16).
  5. Changelog bersifat append-only tanpa API edit; delete work item menghapus seluruh changelog-nya dalam transaksi yang sama.
  6. API menyediakan paginasi changelog (`startAt`/`maxResults`) per work item dan bulk fetch changelog untuk banyak work item sekaligus (maksimal 1.000 work item per permintaan) bagi integrasi dan laporan.
  7. Retensi changelog tidak dibatasi secara default; tabel besar dipartisi per bulan berdasarkan `created_at`.
  8. Representasi `from_string`/`to_string` disimpan saat perubahan terjadi agar nilai historis tetap terbaca walaupun opsi, user, atau version kemudian dihapus atau di-rename.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Translator JQL membatasi `WAS` ke Assignee, Fix Version, Priority, Reporter, Resolution, Status (dan `CHANGED` untuk field yang sama) walaupun penyimpanan bersifat umum (modul 12). Skema lengkap di modul 28; partisi dan retensi di modul 27. Repo saat ini belum memiliki tabel ini; migrasi dibuat via `make migration-create`.

### CMT-024 Work Log Tab (Activity View)

- **Deskripsi:** Tampilan worklog di section Activity; aturan time tracking dan estimasi berada di modul 14.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Work On Issues untuk mencatat; Edit/Delete Own/All Worklogs untuk mengubah (PERM-032).
- **Aturan perilaku (acceptance criteria):**
  1. Tab Work log menampilkan daftar worklog berisi avatar dan nama pengguna, time spent (format `1w 2d 3h 30m`), tanggal mulai, dan deskripsi kerja dalam rich text.
  2. Aksi Edit dan Delete pada worklog tampil sesuai permission Edit Own Worklogs/Edit All Worklogs dan Delete Own Worklogs/Delete All Worklogs; server menolak tanpa permission dengan HTTP 403.
  3. Visibility worklog dapat dibatasi ke role/group dengan mekanisme yang sama seperti comment (CMT-004, PERM-045).
  4. Aksi "Log time" tersedia dari tab Work log maupun dari field Time tracking di context panel, membuka dialog yang sama (modul 14).
  5. Deskripsi worklog memakai varian editor rich text ringkas (tanpa table dan media besar) dari modul 07.
  6. ⚠ Satu work item menampung maksimal 5.000 worklog; worklog berikutnya ditolak dengan pesan "Batas worklog per work item telah tercapai".
- **Team-managed vs company-managed:** Tidak ada perbedaan selain sumber permission.
- **Catatan & ketergantungan:** Bergantung pada modul 14 (aturan estimasi dan rollup) dan CMT-023 (worklog dicatat di History).

### CMT-025 Approvals Tab (Activity View)

- **Deskripsi:** Tab Approvals di section Activity yang menampilkan riwayat dan status approval pada work item yang memakai langkah approval di workflow, baik approval project bisnis (BIZ-042) maupun approval step Service Management (JSM-ITSM-011). Subset bisnis (aturan `one|all`, approver dari field user picker) dirilis di Fase 3 bersama BIZ-042 dan BIZ-046; perluasan JSM (aturan `percent|count`, approver group, numberPerPrincipal, keputusan dari portal/email) ditambahkan di Fase 5.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Approver, Agent Layanan, dan anggota project; dilihat oleh pengguna dengan Browse Projects. Tidak ada peran yang dapat mengedit atau menghapus entri approval.
- **Aturan perilaku (acceptance criteria):**
  1. Tab Approvals menampilkan approver, keputusan (Approved/Declined), waktu keputusan, dan comment keputusan bila ada.
  2. Tab hanya tampil pada project yang workflow-nya memiliki minimal satu status dengan langkah approval (project bisnis lewat BIZ-042 di Fase 3, service project lewat JSM-ITSM-011 di Fase 5); pada project lain tab disembunyikan dan preferensi tab jatuh kembali ke Comments (CMT-021 butir 9).
  3. Approval yang masih pending menampilkan daftar approver yang belum memutuskan beserta aturan jumlah persetujuan yang berlaku; untuk project bisnis aturan ditampilkan sebagai "Satu approver" (`one`) atau "Semua approver" (`all`), dan untuk service project juga "Persentase" (`percent`) atau "N approver" (`count`), memakai enum `rule` yang sama dengan model data JSM-ITSM-011.
  4. Hasil approval juga tercatat sebagai transisi status di tab History (CMT-022).
  5. Riwayat langkah approval sebelumnya tetap ditampilkan dan dikelompokkan per langkah (status dan waktu mulai) bila work item kembali ke status approval, sesuai BIZ-046 butir 2.
  6. Entri approval bersifat read-only: UI tidak menyediakan aksi edit/hapus dan API menolak perubahan entri riwayat approval (BIZ-046 butir 3).
  7. ⚠ Tab menampilkan peringatan "Belum ada approver" bila field approver kosong saat work item berada di status approval, selaras dengan BIZ-043 butir 4 dan JSM-ITSM-011 butir 10.
  8. Tombol Approve/Decline tidak berada di tab ini; keputusan diambil dari panel approval di issue view (BIZ-044 untuk project bisnis, JSM-ITSM-013 untuk service project), dan tab hanya menampilkan hasilnya.
- **Team-managed vs company-managed:** Isi tab tidak berbeda; ketersediaan langkah approval mengikuti BIZ-042 (company-managed lebih dulu, team-managed di balik feature flag) dan JSM-ITSM-011 (team-managed lewat workflow request type).
- **Catatan & ketergantungan:** Konfigurasi dan keputusan approval project bisnis di BIZ-042 s.d. BIZ-046 ([26 Proyek Bisnis, Formulir & Persetujuan](26-kerja-bisnis-formulir-persetujuan.md)); approval Service Management di JSM-ITSM-011 s.d. JSM-ITSM-015 ([32 JSM ITSM & Approvals](32-jsm-itsm-insiden-problem-perubahan-layanan.md)). Fase dipindah dari Fase 5 ke Fase 3 agar tab sudah tersedia saat approval bisnis dirilis. Enum aturan jumlah memakai satu kolom `rule` bersama (`one|all|percent|count`); project bisnis hanya memakai `one|all`, sehingga penamaan `any|all` di BIZ-042/BIZ-043 perlu diselaraskan di modul 26 menjadi `one|all`.

### CMT-026 Automation Panel (Recent Rule Runs)

- **Deskripsi:** Panel Automation di issue view yang menampilkan eksekusi rule terbaru terkait work item dan tempat menjalankan rule manual.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects melihat panel; Admin Proyek melihat audit log rule; manual trigger mengikuti pengaturan rule.
- **Aturan perilaku (acceptance criteria):**
  1. Issue view memiliki panel/glance "Automation" dengan bagian "Recent rule runs" dan tombol refresh.
  2. Panel menampilkan hingga 8 eksekusi rule terakhir yang menyentuh work item: nama rule, status (Success, No actions performed, Some errors, Failed), dan waktu eksekusi.
  3. Daftar eksekusi tidak memiliki filter atau pencarian.
  4. Setiap entri memiliki tautan ke audit log rule bagi pengguna yang berhak melihat rule; audit log menyimpan riwayat 90 hari (modul 20).
  5. Panel dan menu Automation (ikon petir) menyediakan daftar rule manual trigger yang berlaku untuk work item; menjalankannya mengeksekusi rule dengan konteks work item tersebut (modul 20).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ini panel, bukan tab Activity. Data eksekusi dibaca dari audit log Automation (modul 20).

### CMT-027 Activity Stream Tab (Legacy)

- **Deskripsi:** Tab "Activity" bergaya activity stream dan tab source/commit integrasi lama di issue view Server/DC; dicatat sebagai rujukan dan tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna Data Center (tidak relevan untuk produk ini).
- **Aturan perilaku (acceptance criteria):**
  1. Tab activity stream bergaya feed RSS tidak dibangun; tab All (CMT-021) menggantikannya.
  2. Tab Source/Reviews dari integrasi Fisheye/Crucible tidak dibangun.
  3. Informasi commit dan pull request ditampilkan lewat panel Development (modul 24), bukan tab Activity.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan.

## Berbagi & konten tertaut

### CMT-028 Share Work Item

- **Deskripsi:** Membagikan work item ke pengguna lain lewat notifikasi/email dengan pesan opsional, atau menyalin tautannya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Browse Projects pada work item; undangan pengguna baru tunduk pada pengaturan invite (USR-009).
- **Aturan perilaku (acceptance criteria):**
  1. Header issue view memiliki tombol "Share".
  2. Dialog share menyediakan input penerima dengan typeahead untuk user, group, atau alamat email, serta kolom pesan opsional (maksimal 1.000 karakter).
  3. Penerima mendapat email dan notifikasi in-app berisi tautan work item, summary, nama pengirim, dan pesan (modul 19).
  4. ⚠ Share ke alamat email di luar site mengirim undangan bergabung bila pengguna boleh mengundang (USR-009, USR-012); bila tidak boleh, alamat tersebut ditolak dengan pesan "Anda tidak dapat mengundang pengguna baru".
  5. Dialog menampilkan peringatan per penerima yang tidak memiliki akses ke work item ("<nama> tidak dapat melihat work item ini"); pengirim tetap dapat mengirim, tetapi penerima tanpa akses hanya mendapat tautan tanpa isi summary.
  6. Tombol "Copy link" di dialog share dan ikon link di header menyalin URL absolut `/browse/KEY`.
  7. ⚠ Share ke Slack/Microsoft Teams tersedia di dialog yang sama hanya bila integrasi chat aktif (modul 25); bila tidak aktif, opsi tidak ditampilkan.
  8. Share tidak menjadikan penerima sebagai watcher work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur ini juga muncul sebagai ISS-038 Share Work Item di modul 06 (aksi di header issue view); modul ini menjadi acuan perilaku dialog dan pengiriman. Pengiriman email via modul 19.

### CMT-029 Web Links and Remote Links Display

- **Deskripsi:** Menampilkan tautan web eksternal dan remote link dari integrasi dengan tampilan smart link; issue link antar work item berada di modul 06.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna dengan Link Issues (PERM-024) untuk menambah/mengubah; Browse Projects untuk melihat.
- **Aturan perilaku (acceptance criteria):**
  1. Aksi "Add web link" membuka form dengan field URL (wajib, harus http/https) dan Link text (opsional; default URL).
  2. Web link dirender sebagai smart link dengan ikon, judul, dan status bila resolver mendukung domain tersebut (EDT-003); bila tidak, tampil sebagai tautan teks dengan favicon.
  3. Web link dapat diedit (URL dan teks) dan dihapus oleh pengguna dengan Link Issues; penghapusan meminta konfirmasi.
  4. Remote link dari integrasi (repositori kode, wiki, app) tampil dengan ikon aplikasi sumber, judul, dan status yang dilaporkan integrasi.
  5. Satu work item menampung maksimal 2.000 remote link; penambahan berikutnya ditolak dengan pesan "Batas 2.000 remote link per work item telah tercapai".
  6. ⚠ URL yang sudah ditautkan pada work item yang sama ditolak dengan pesan "Tautan ini sudah ada di work item"; pemeriksaan dilakukan setelah normalisasi URL (skema dan host huruf kecil, tanpa trailing slash).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Model remote link (globalId, application, object) di ISS-030 (modul 06) dan modul 23; issue link antar work item di ISS-029. Penambahan/penghapusan web link dicatat di History (CMT-022).

### CMT-030 Linked Confluence Pages

- **Deskripsi:** Menautkan dan menampilkan halaman wiki/dokumentasi yang terkait dengan work item; di produk ini ditujukan untuk integrasi wiki eksternal atau modul halaman internal di masa depan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Link Issues; akses ke halaman mengikuti izin di sistem wiki.
- **Aturan perilaku (acceptance criteria):**
  1. Issue view memiliki section "Pages" (setara "Confluence pages") yang tampil bila integrasi wiki aktif.
  2. Aksi "Link page" memungkinkan mencari halaman berdasarkan judul atau menempel URL halaman.
  3. Halaman yang menyebut key work item (lewat macro atau smart link) otomatis tampil dengan label "Mentioned in" berdasarkan callback dari integrasi.
  4. Halaman yang ditautkan secara eksplisit tampil dengan label "Linked".
  5. Klik halaman membukanya di tab baru; hover menampilkan smart link card berisi judul, space, dan penulis.
  6. Aksi "Unlink" menghapus remote link halaman setelah konfirmasi; halaman "Mentioned in" tidak dapat di-unlink dari sisi work item.
  7. Section hanya aktif bila sistem wiki terhubung ke site (konfigurasi integrasi di modul 25); tanpa integrasi, section disembunyikan.
  8. Halaman yang tidak dapat diakses viewer ditampilkan terbatas sebagai "Halaman terbatas" tanpa judul.
  9. Tautan disimpan sebagai remote issue link dengan `globalId` berformat `appId=<id>&pageId=<id>` sehingga tidak terduplikasi.
  10. ⚠ Jumlah halaman terkait tampil sebagai badge pada glance/section.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Produk Confluence adalah layanan Atlassian lintas produk dan di luar cakupan; fitur ini dibangun sebagai integrasi generik (modul 25) atau web link biasa (CMT-029).

### CMT-031 Create Confluence Page from Issue

- **Deskripsi:** Membuat halaman wiki baru langsung dari work item dan otomatis menautkannya, bila integrasi wiki mendukung pembuatan halaman.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Link Issues dan akses tulis di sistem wiki.
- **Aturan perilaku (acceptance criteria):**
  1. Opsi "Create page" tersedia di menu "+ Add" dan di section Pages bila integrasi wiki aktif.
  2. ⚠ Dialog meminta pemilihan space dan template; daftar template diambil dari sistem wiki, dengan opsi "Blank" selalu tersedia.
  3. Editor halaman dibuka di tab baru (atau side panel bila integrasi mendukung) dengan smart link ke work item sudah terisi.
  4. Halaman baru otomatis ditautkan ke work item sebagai "Linked" (CMT-030).
  5. ⚠ Pembuatan whiteboard atau database dari work item tidak disediakan.
  6. Prompt cross-sell untuk mencoba produk wiki tidak ditampilkan; bila integrasi tidak aktif, opsi "Create page" disembunyikan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada CMT-030 dan integrasi di modul 25.

### CMT-032 Loom Video Recording and Embed

- **Deskripsi:** Merekam atau menyematkan video di description/comment; produk ini mendukung embed URL video lewat oEmbed dan upload video sebagai attachment, sedangkan perekaman native Loom adalah layanan pihak lain.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna dengan Edit Issues/Add Comments; Admin Situs mengaktifkan integrasi video.
- **Aturan perilaku (acceptance criteria):**
  1. Editor comment/description menampilkan tombol "Record a video" hanya bila integrasi perekaman video aktif; tanpa integrasi tombol disembunyikan.
  2. Perekaman layar/kamera dilakukan lewat SDK penyedia video; hasilnya langsung disisipkan sebagai smart link embed di editor.
  3. Menempel URL video dari penyedia yang didukung oEmbed (Loom, YouTube, Vimeo) menghasilkan player inline.
  4. Embed video dapat ditampilkan dalam mode inline (teks), card, atau embed (player), sesuai pilihan tampilan smart link (EDT-003).
  5. Viewer dapat menonton sesuai pengaturan berbagi di penyedia video; bila tidak berhak, player menampilkan pesan dari penyedia.
  6. ⚠ Ketersediaan perekaman bergantung pada integrasi penyedia yang aktif di organisasi dan dikendalikan admin per site.
  7. Alternatif tanpa integrasi: upload file video (MP4) sebagai attachment (CMT-011) yang diputar inline lewat Media Viewer (CMT-015).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Loom adalah layanan Atlassian lintas produk (di luar cakupan sebagai fitur native); oEmbed dan embed smart link di modul 07.

## Pertanyaan terbuka

- Apakah produk ini perlu mendukung pembuatan comment dari balasan email notifikasi pada project non-service, atau cukup lewat mail handler service project (CMT-001 butir 15)?
- Saat comment induk dengan reply dihapus, apakah reply dipertahankan dengan placeholder induk (usulan saat ini) atau ikut terhapus (CMT-003 butir 8)?
- Berapa batas jumlah pinned comment per work item dan siapa yang boleh pin/unpin (usulan: 3 comment, pengguna dengan Add Comments) (CMT-007)?
- Apakah REST API publik perlu langsung menyediakan pembuatan reply (`parentId`) walau melampaui paritas Jira (CMT-005)?
- Apakah antivirus scan (ClamAV) diaktifkan secara default di distribusi produksi atau hanya opsional (CMT-017, CMT-019)?
- Berapa batas ukuran file default dan kuota storage default untuk instance hosted produk ini (usulan 100 MB per file, kuota tanpa batas di self-hosted) (CMT-016, CMT-019)?
- Apakah layanan konversi dokumen Office dan transcoder video disertakan sebagai komponen bawaan atau sebagai dependency opsional (CMT-015)?
- Apakah fitur Image Annotation tetap dibangun mengingat statusnya di Jira Cloud tidak pasti (CMT-020)?
- Integrasi wiki mana yang menjadi target pertama (modul halaman internal, Confluence via API, atau wiki open source lain) untuk CMT-030 dan CMT-031?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga:** tab activity tambahan dari app Forge/Connect, app pengelola attachment (mis. preview format khusus, anotasi lanjutan), app template comment, dan app pelacak email-to-comment tidak dibangun sebagai fitur native; produk ini hanya menyediakan titik ekstensi di modul 25.
- **Fitur Data Center-only / deprecated:** Activity Stream Tab (Legacy) beserta tab Source/Reviews Fisheye/Crucible (CMT-027), serta pengaturan thumbnail dan ZIP preview gaya Server/DC (CMT-016 butir 8) hanya dicatat sebagai rujukan.
- **Layanan Atlassian lintas produk:** Confluence (CMT-030, CMT-031 hanya sebagai integrasi generik), Loom (CMT-032 hanya sebagai embed/integrasi), Atlassian Media Platform (digantikan pipeline storage dan preview sendiri di CMT-015, CMT-019), layanan Reactions platform (dibangun sendiri di CMT-006), Atlassian Guard/data security policies (dirujuk ke PERM-065), serta Rovo/Atlassian Intelligence (ringkasan comment AI opsional dirujuk ke modul 06).
- **Integrasi chat:** pengiriman Share ke Slack/Microsoft Teams bergantung pada modul 25 dan tidak menjadi bagian dari Fase 2 modul ini.
