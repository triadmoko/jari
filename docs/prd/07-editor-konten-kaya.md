# 07 — Editor Teks Kaya & Rendering Konten

Modul ini mendefinisikan editor teks kaya (rich text) terpadu yang dipakai di seluruh produk ini untuk Description, Comment, field sistem Environment, custom field Paragraph, deskripsi worklog, Create dialog, dan transition screen, beserta format dokumen penyimpanannya, katalog elemen (mark format, heading, list, code block, quote, tabel, panel, expand, divider, emoji, mention, action item, status, date, layout, media, smart link), cara input cepat (markdown shortcut, slash command, keyboard shortcut, paste), perilaku field Description dan field rich text lain, kompatibilitas dengan renderer wiki markup lama, rendering aman dari XSS, serta konversi konten ke kanal lain (email, API, ekspor, indeks pencarian, otomasi). Format penyimpanan utama adalah dokumen JSON yang kompatibel dengan Atlassian Document Format (ADF), bukan HTML, sehingga konten dapat divalidasi di server, di-render lewat whitelist node/mark, dan dikonversi ke kanal lain secara deterministik. Kondisi repo saat ini (backend Go + Gin + GORM + PostgreSQL dengan fitur auth dan user; frontend SvelteKit 5 SPA) belum memiliki editor, kolom konten rich text, maupun pipeline sanitasi; seluruh isi modul ini adalah pekerjaan baru. Rekomendasi teknis awal: editor di frontend berbasis ProseMirror/TipTap dengan schema yang dipetakan ke ADF, validator dan konverter ADF di backend Go (paket bersama di `pkg/`), penyimpanan ADF di kolom JSONB plus kolom plaintext turunan untuk indeks pencarian. Penyimpanan berkas lampiran, thread komentar, dan Activity diatur di modul 08; modul ini hanya mengatur bagaimana konten tersebut ditulis dan ditampilkan.

Modul terkait:

- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission Edit Issues, Add Comments, Edit All Comments, Create Attachments, Browse users and groups yang mengontrol editor.
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — field configuration (renderer, required, default value), custom field Paragraph/Short text, layout team-managed, template description.
- [06 Work Item Inti](06-work-item-inti.md) — issue view, Create dialog, Clone, History field Description, Issue Link formal.
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — comment thread, visibilitas comment, attachment, Media Viewer, comment reactions, Activity.
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — undangan pengguna dari typeahead mention, profil dan privasi email.
- [09 Workflow, Status & Transisi](09-workflow-status-transisi.md) — status workflow (berbeda dari Status Lozenge Element), transition screen yang memakai editor.
- [18 Timeline, Rencana & Tampilan Proyek](18-timeline-rencana-tampilan-proyek.md) — entitas Teams untuk mention team.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — pengaturan global custom emoji, kebijakan preview/embed, aktivasi AI.
- [24 Integrasi DevOps](24-integrasi-devops.md) — resolver smart link repositori/PR dan pengenalan issue key di commit.
- [12 JQL](12-jql-bahasa-kueri.md) dan [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — operator `text ~`/`description ~`, indeks plaintext, ekspor CSV/Word/XML.
- [19 Notifikasi In-app & Email](19-notifikasi-email.md) — notifikasi "mentioned you", render HTML email.
- [20 Otomasi](20-otomasi.md) — smart value `{{issue.description}}`/`{{comment.body}}`, input wiki markup dari aksi otomasi.
- [21 Navigasi, UX & Preferensi](21-navigasi-ux-preferensi.md) — shortcut global issue view, preferensi UI per user (link preferences).
- [23 API Publik & Webhook](23-api-publik-webhook.md) — body ADF di REST API v3, wiki markup di REST API v2.
- [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md) — impor CSV ber-wiki markup, SSRF guard, integrasi pihak ketiga untuk smart link.
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) — batas ukuran, keamanan, aksesibilitas, AI opsional.
- [28 Model Data Inti](28-model-data-inti.md) — skema penyimpanan konten rich text, comment, attachment.
- [29 JSM Service Project](29-jsm-service-project-request-type-formulir.md) dan [30 JSM Portal](30-jsm-help-center-portal-pelanggan-kanal.md) — editor dipakai agent dan customer di portal.
- [36 JPD Ideas](36-jpd-ideas-field-views.md) — deskripsi idea memakai editor yang sama.

## Persona terkait

- **Developer/Anggota Tim** — penulis utama description dan comment; memakai code block, tabel, mention, action item, markdown shortcut, dan paste screenshot.
- **Product Owner/PM** — menulis spesifikasi panjang di Description dengan heading, panel, tabel, expand, dan smart link ke dokumen eksternal.
- **Scrum Master/Team Lead** — memakai action item, mention, dan date untuk menindaklanjuti keputusan di comment.
- **Stakeholder/Viewer** — membaca konten ter-render, mencentang action item bila berizin, membuka hover card dan Media Viewer.
- **Pengguna Tim Bisnis** — memakai editor tanpa pengetahuan markup; mengandalkan toolbar, slash command, dan paste dari Word/Google Docs.
- **Admin Situs** — mengatur renderer field, batas custom emoji, kebijakan preview/embed, dan aktivasi AI.
- **Admin Proyek/Project Lead** — mengatur Description (required, default, layout) di project.
- **Integrator/Developer API** — membaca/menulis ADF lewat REST API v3 dan wiki markup lewat API v2; bergantung pada validasi dan konversi yang deterministik.
- **Agent Layanan** dan **Customer/Pelapor** — memakai editor yang sama di request (agent view dan portal) di modul 29-30.

## User story modul

- Sebagai Developer/Anggota Tim, saya ingin menulis description dengan heading, list, dan code block ber-syntax highlighting agar langkah reproduksi bug mudah dibaca (EDT-001, EDT-005, EDT-006, EDT-007).
- Sebagai Developer/Anggota Tim, saya ingin mem-paste screenshot langsung ke comment agar bukti visual tersimpan sebagai attachment tanpa langkah unggah terpisah (EDT-019, EDT-025).
- Sebagai Scrum Master/Team Lead, saya ingin me-mention rekan dengan '@' dan menambahkan action item bertanggal agar tindak lanjut jelas dan orang yang disebut menerima notifikasi (EDT-014, EDT-016, EDT-018).
- Sebagai Product Owner/PM, saya ingin menyusun spesifikasi dengan tabel, panel, dan expand agar informasi kritis menonjol dan detail panjang bisa dilipat (EDT-010, EDT-011, EDT-012).
- Sebagai Pengguna Tim Bisnis, saya ingin menyisipkan elemen lewat menu '/' dan mem-paste dari Word tanpa kehilangan format agar saya tidak perlu belajar sintaks markup (EDT-023, EDT-025).
- Sebagai Developer/Anggota Tim, saya ingin mengetik sintaks markdown yang otomatis dikonversi agar menulis secepat di editor kode (EDT-022, EDT-024).
- Sebagai Stakeholder/Viewer, saya ingin melihat tautan work item lain sebagai kartu dengan status terkini agar saya tahu progres tanpa membuka setiap tautan (EDT-020, EDT-021).
- Sebagai Developer/Anggota Tim, saya ingin draft description yang belum disimpan tidak hilang saat saya tidak sengaja menutup issue agar pekerjaan menulis tidak terbuang (EDT-001).
- Sebagai Admin Situs, saya ingin semua konten pengguna di-render lewat whitelist dan divalidasi di server agar produk ini aman dari XSS dan dokumen berbahaya (EDT-002).
- Sebagai Admin Proyek/Project Lead, saya ingin menjadikan Description wajib dan memberi template default agar setiap work item berisi informasi minimum yang dibutuhkan tim (EDT-026, EDT-027).
- Sebagai Integrator/Developer API, saya ingin membaca dan menulis konten sebagai ADF (API v3) maupun wiki markup (API v2) agar integrasi lama dan baru tetap berjalan (EDT-003, EDT-028).
- Sebagai Stakeholder/Viewer, saya ingin email notifikasi menampilkan konten terformat dengan aman agar saya bisa membaca perubahan tanpa membuka aplikasi (EDT-003).
- Sebagai Developer/Anggota Tim, saya ingin bantuan AI opsional untuk merapikan tulisan dan merangkum comment agar komunikasi lebih ringkas (EDT-030).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 9 |
| P1 | 12 |
| P2 | 5 |
| P3 | 4 |
| **Total** | **30** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 9 |
| Fase 2 | 13 |
| Fase 3 | 5 |
| Fase 4 | 2 |
| Fase — (dc-only/deprecated, rujukan) | 1 |

Strategi urutan implementasi: Fase 1 membangun fondasi yang tidak boleh berubah belakangan, yaitu schema dokumen ADF, validator dan renderer whitelist di server (EDT-002), editor inti dengan dua mode (full dan comment), click-to-edit, draft lokal, serta elemen yang paling sering dipakai (mark format, heading, list, code block, mention, media) dan perilaku field Description. Schema ADF harus didefinisikan lengkap sejak awal (termasuk node yang baru dibangun UI-nya di fase berikut) agar dokumen yang dibuat lewat API atau impor tidak ditolak dan node yang belum didukung UI ditampilkan sebagai blok fallback tanpa hilang. Fase 2 menambah elemen struktur (tabel, panel, quote, divider, action item, emoji standar), akselerator input (markdown shortcut, slash command, keyboard shortcut, paste dari Word/Sheets), unfurl work item internal, rich text di custom field, serta konversi ke kanal email dan plaintext. Fase 3 menambah elemen sekunder (expand, status, date), smart link generik dengan resolver OpenGraph/oEmbed sisi server ber-SSRF guard, dan kompatibilitas wiki markup untuk REST API v2 dan impor. Fase 4 menyisakan layout multi-kolom (bila dikonfirmasi) dan asisten AI opsional; mode teks wiki markup tidak dibangun.

## Fitur

## Fondasi editor, format dokumen & keamanan

### EDT-001 Rich Text Editor (Atlassian Editor / ADF) - Core

- **Deskripsi:** Editor WYSIWYG terpadu yang menyimpan konten sebagai dokumen JSON kompatibel ADF (bukan HTML) dan dipakai di Description, Comment, custom field Paragraph, Environment, deskripsi worklog, Create dialog, dan transition screen. Editor dan renderer read-only berbagi schema yang sama sehingga tampilan baca konsisten dengan tampilan tulis.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan permission Edit Issues (Description dan field), Add Comments (comment), Create Issues (Create dialog); viewer dengan Browse Projects melihat renderer read-only.
- **Aturan perilaku (acceptance criteria):**
  1. Tersedia dua mode editor: full editor untuk Description/Paragraph dengan toolbar lengkap, dan comment editor dengan toolbar ringkas, tombol Save/Cancel, serta kondisi tertutup (satu baris placeholder) sampai diklik.
  2. Toolbar utama full editor memuat: dropdown text style (Normal text, Heading 1-6), Bold, Italic, menu More formatting (Underline, Strikethrough, Code, Subscript, Superscript, Clear formatting), text color, bullet list, numbered list, action item, link, mention, emoji, table, code snippet, media, dan menu insert '+'.
  3. Menu insert '+' memuat: Action item, Link, Mention, Emoji, Table, Code snippet, Info panel, Quote, Divider, Expand, Date, Status, Image/video/file, dan 'View more' yang membuka dialog Browse all elements (EDT-023). Entri toolbar maupun menu '+' untuk elemen yang belum dirilis pada fase berjalan (mis. Expand, Date, Status di Fase 3) disembunyikan, bukan ditampilkan nonaktif.
  4. Toolbar responsif: item yang tidak muat pada lebar sempit dipindahkan ke menu overflow (...) tanpa ada aksi yang hilang.
  5. Di issue view, field rich text tampil read-only dan masuk mode edit saat diklik; Save (atau Ctrl/Cmd+Enter) menyimpan, Cancel (atau Esc) membatalkan dan mengembalikan nilai terakhir tersimpan.
  6. Menutup issue atau berpindah navigasi ketika editor kotor (ada perubahan belum disimpan) menampilkan konfirmasi 'You have unsaved changes' dengan opsi tetap mengedit atau discard.
  7. Menutup Create dialog yang sudah berisi konten menampilkan konfirmasi discard (detail dialog di [06 Work Item Inti](06-work-item-inti.md)).
  8. ⚠ Menutup tab atau reload browser dengan editor kotor memicu peringatan `beforeunload` bawaan browser.
  9. ⚠ Isi editor yang belum disimpan dipertahankan sementara sebagai draft di browser dan dipulihkan otomatis saat editor yang sama dibuka lagi; durasi retensi draft ditetapkan (usulan 7 hari).
  10. ⚠ Draft disimpan dengan kunci per user + work item + field (localStorage), tidak tersinkron antar perangkat, dan dihapus setelah Save atau Cancel.
  11. Undo/Redo berbasis transaksi editor tersedia lewat Ctrl/Cmd+Z dan Ctrl/Cmd+Shift+Z atau Ctrl+Y; riwayat undo terbatas pada sesi edit yang sedang berjalan.
  12. Editor kosong menampilkan placeholder kontekstual (mis. 'Add a description...' untuk Description; 'Type @ to mention...' untuk comment; petunjuk '/ai' hanya bila AI aktif, lihat EDT-030).
  13. Description, Comment, dan field text multi-line dibatasi 32.767 karakter (dihitung dari plaintext turunan); melampaui batas menampilkan pesan galat di editor dan API menolak dengan HTTP 400 beserta nama field.
  14. Backend memvalidasi dokumen ADF setiap kali disimpan: node/mark tak dikenal ditolak atau di-fallback sesuai EDT-002, kedalaman nesting dan ukuran dokumen dibatasi; batas total ukuran konten per work item (di luar attachment) dapat dikonfigurasi dengan default 500 MB.
  15. Editor tidak ditampilkan (renderer read-only saja) bila user tidak memiliki Edit Issues untuk Description/field, atau tidak memiliki Add Comments untuk comment editor.
  16. ⚠ Tidak ada mode fullscreen untuk Description; editor tumbuh mengikuti tinggi konten.
  17. Toolbar tetap menempel (sticky) di atas area edit saat pengguna menggulir dokumen panjang.
  18. Renderer read-only memakai schema dan gaya yang sama dengan editor sehingga hasil render identik secara visual dengan tampilan saat mengedit.
  19. Node ADF yang tidak didukung klien ditampilkan sebagai blok fallback 'unsupported content' (unsupportedBlock/unsupportedInline); saat dokumen disimpan ulang, node tersebut dipertahankan utuh dan tidak hilang.
  20. Input IME (CJK) dan teks RTL dasar tidak merusak struktur dokumen; spell check memakai spell checker native browser (`spellcheck=true`).
  21. Aksesibilitas: toolbar dapat dinavigasi dengan keyboard (roving tabindex), setiap tombol memiliki ARIA label, dan penyisipan node diumumkan ke screen reader lewat live region.
  22. ⚠ Tidak ada co-editing real-time pada Description; bila user lain menyimpan perubahan saat seseorang sedang mengedit, issue view memberi tahu adanya pembaruan dan penyimpanan berikutnya mengikuti aturan konflik di EDT-026.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada editor; perbedaan hanya pada field mana yang rich text (lihat EDT-027 dan EDT-028).
- **Catatan & ketergantungan:** Penyimpanan: kolom JSONB `*_adf` plus kolom plaintext turunan untuk JQL `text ~` (modul 12/13, model di modul 28). Frontend SvelteKit memerlukan komponen editor bersama (ProseMirror/TipTap) yang dipakai ulang oleh modul 06, 08, 14, 26, 29, 30, 36. Draft lokal bergantung pada preferensi UI per user di modul 21. Keputusan terbuka: durasi retensi draft dan apakah draft perlu tersinkron di server.

### EDT-002 Safe Rich Text Rendering (ADF & Wiki Markup)

- **Deskripsi:** Persyaratan keamanan untuk menyimpan dan menampilkan konten pengguna (description, comment, field paragraph) tanpa celah XSS, injeksi, maupun DoS. Semua render HTML hanya dilakukan dari ADF tervalidasi lewat renderer whitelist.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / platform
- **Aktor & izin:** Engineering (persyaratan platform); berlaku untuk semua jalur tulis (UI, REST API, impor, otomasi).
- **Aturan perilaku (acceptance criteria):**
  1. Format penyimpanan utama adalah ADF JSON; HTML hanya dihasilkan oleh renderer yang memakai whitelist node dan mark, dan tidak pernah menyisipkan string HTML mentah dari pengguna (tidak ada `{@html}` atas input pengguna di frontend).
  2. Server memvalidasi skema ADF pada setiap penyimpanan; node tidak dikenal ditolak (HTTP 400 dengan path node) atau di-strip sesuai mode jalur masuk, dan atribut di luar skema dibuang.
  3. URL pada link, smart link, dan media disanitasi: hanya skema `http`, `https`, `mailto` (dan skema aman lain yang ditetapkan) yang diterima; `javascript:`, `data:`, dan `vbscript:` ditolak di server dan di klien.
  4. Wiki markup dari REST API v2 dan field renderer wiki (company-managed) dikonversi ke ADF lewat konverter yang juga melakukan sanitasi; konten tidak pernah di-render langsung dari wiki markup ke HTML.
  5. Markdown shortcut dan paste dikonversi ke ADF di klien, lalu hasilnya divalidasi ulang di server sebelum disimpan (klien tidak dipercaya).
  6. Node yang didukung renderer mencakup minimal: mention, emoji, status lozenge, panel, code block (highlight di sisi klien tanpa `eval`), table, media, expand, date, inline card; daftar lengkap mengikuti schema EDT-001.
  7. Ukuran dokumen dan kedalaman nesting dibatasi (usulan: maksimum 32.767 karakter teks, ukuran JSON maksimum 1 MB per field, kedalaman maksimum 20) untuk mencegah DoS; pelanggaran ditolak dengan HTTP 400/413.
  8. Unfurl smart link dan pengambilan metadata URL dilakukan di sisi server melalui SSRF guard (blokir IP privat/loopback/link-local/metadata cloud, batas redirect, timeout, batas ukuran respons).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Validator ADF ditempatkan sebagai paket bersama Go (mis. `pkg/adf`) yang dipakai handler REST, impor, dan otomasi. Content Security Policy dan header keamanan HTTP di modul 27; SSRF guard bersama di modul 25. Test wajib: korpus payload XSS (OWASP) terhadap renderer HTML email dan SPA.

### EDT-003 Rich Text Rendering Across Channels

- **Deskripsi:** Konversi dan penyajian konten rich text di luar issue view: email, ekspor, mobile, API, indeks pencarian, dan otomasi. ADF tetap menjadi sumber kebenaran; HTML, plaintext, dan wiki markup adalah turunan yang boleh di-cache.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Sistem (pipeline notifikasi, indeks, ekspor) dan Integrasi (REST API, otomasi); konten yang dikeluarkan tetap tunduk pada permission viewer/penerima.
- **Aturan perilaku (acceptance criteria):**
  1. ADF dikonversi ke HTML untuk email notifikasi dengan dukungan mention (nama tampilan), emoji (karakter Unicode atau gambar), tabel, dan gambar yang ditampilkan sebagai link aman ke produk ini (bukan lampiran tertanam tanpa autentikasi).
  2. ADF dikonversi ke wiki markup atau plain text untuk REST API v2, ekspor CSV, dan smart value otomasi (`{{issue.description}}`, `{{comment.body}}`); hasil konversi deterministik untuk input yang sama.
  3. ADF dikonversi ke plain text untuk indeks pencarian dan preview ringkas (hover card, card board); mention menjadi nama tampilan dan media menjadi nama file.
  4. Klien mobile (di luar cakupan awal) atau web responsif menampilkan subset elemen yang didukung; elemen lain memakai fallback EDT-001 aturan 19.
  5. Ekspor Word/Printable/XML menyertakan description dan comment dalam format terformat (HTML hasil renderer whitelist).
  6. Node yang tidak didukung suatu kanal di-fallback ke representasi teks (mis. status menjadi `[TEKS]`, date menjadi tanggal terformat, smart link menjadi URL).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pipeline email di modul 19; ekspor di modul 13; smart value di modul 20; API v2/v3 di modul 23. Cache turunan (plaintext/HTML) disimpan di modul 28 dan diinvalidasi saat konten berubah. Aplikasi mobile native di luar cakupan PRD ini.

## Format teks & struktur blok dasar

### EDT-004 Text Formatting Marks

- **Deskripsi:** Format inline pada teks (Text Formatting Marks): tebal, miring, garis bawah, coret, inline code, subscript/superscript, warna teks, dan link. Disimpan sebagai array `marks` pada node text ADF.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Editor konten (pengguna yang sedang mengedit field rich text sesuai EDT-001).
- **Aturan perilaku (acceptance criteria):**
  1. Bold (Ctrl/Cmd+B), Italic (Ctrl/Cmd+I), Underline (Ctrl/Cmd+U), dan Strikethrough (Ctrl/Cmd+Shift+S) dapat diterapkan dan dilepas (toggle) pada seleksi atau pada posisi kursor untuk teks yang diketik berikutnya.
  2. Inline code diterapkan lewat Ctrl/Cmd+Shift+M atau dengan mengetik teks di antara backtick (`` `code` ``).
  3. Subscript dan Superscript saling eksklusif pada satu rentang teks; menerapkan salah satunya melepas yang lain.
  4. Text color dipilih dari palet tetap (default plus warna preset); UI tidak menyediakan input hex custom. Server menerima nilai hex valid (`#RRGGBB`) dari API/impor agar dokumen lama tidak ditolak, menolak nilai yang bukan hex valid dengan HTTP 400, dan renderer menampilkan warna di luar palet apa adanya.
  5. ⚠ Background/highlight color teks tersedia sebagai mark terpisah dari palet tetap; bila diputuskan tidak dibangun, mark tersebut tetap diterima dari API dan di-render.
  6. Clear formatting (Ctrl/Cmd+\) menghapus semua mark pada seleksi, kecuali link.
  7. Mark code tidak dapat dikombinasikan dengan sebagian mark lain (mis. text color); tombol mark yang bentrok dinonaktifkan saat seleksi berada di inline code, sedangkan mark lain dapat dikombinasikan bebas.
  8. Link mark disisipkan/diubah lewat Ctrl/Cmd+K dengan field URL dan display text; mengklik link di editor menampilkan floating toolbar Edit, Unlink, Open, Copy.
  9. URL divalidasi skemanya (http, https, mailto, dan skema aman lain); skema berbahaya seperti `javascript:` ditolak dengan pesan galat, dan prefix `https://` ditambahkan otomatis bila skema tidak ada.
  10. Mark disimpan sebagai array `marks` pada node text ADF dengan tipe `strong`, `em`, `underline`, `strike`, `code`, `subsup`, `textColor`, dan `link`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sanitasi URL di EDT-002. Konversi mark ke wiki markup di EDT-028.

### EDT-005 Headings and Paragraph Styles

- **Deskripsi:** Gaya blok teks (Headings and Paragraph Styles): Normal text dan Heading 1 sampai 6, termasuk aturan penempatan heading di dalam container lain.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Dropdown text style menampilkan Normal text dan Heading 1-6, dengan gaya aktif ditandai sesuai posisi kursor.
  2. Shortcut Ctrl/Cmd+Alt+0 mengubah blok menjadi Normal text dan Ctrl/Cmd+Alt+1..6 menjadi Heading 1-6.
  3. Mengetik '# ' sampai '###### ' di awal baris mengubah baris menjadi Heading 1-6 (lihat EDT-022).
  4. Schema ADF mengizinkan heading di dalam table cell, panel, expand, dan layout column, tetapi tidak di dalam list item; perintah heading dinonaktifkan di dalam list item.
  5. ⚠ Alignment paragraf/heading (left/center/right) tersedia sebagai mark `alignment` di blok; bila tidak dibangun di UI, atribut tersebut tetap diterima dari API dan di-render.
  6. Shift+Enter menyisipkan hard break di dalam paragraf yang sama, sedangkan Enter membuat paragraf baru.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Heading dipakai juga oleh template description (modul 05, CFG-052).

### EDT-006 Bullet and Numbered Lists

- **Deskripsi:** List berpoin dan bernomor dengan nesting, indent/outdent, nomor awal custom, dan konversi antar tipe.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Bullet list dibuat lewat Ctrl/Cmd+Shift+8 dan numbered list lewat Ctrl/Cmd+Shift+7, atau lewat tombol toolbar.
  2. Mengetik '* ', '- ', atau '+ ' di awal baris membuat bullet list; '1. ' atau '1) ' membuat numbered list.
  3. ⚠ Tab melakukan indent (nested list) dan Shift+Tab melakukan outdent; kedalaman nesting maksimum 6 level, dan Tab pada level maksimum tidak berbuat apa-apa.
  4. Numbered list mendukung nomor awal custom lewat atribut `order` (mis. mengetik '5. ' memulai list dari 5).
  5. Menekan Enter pada item kosong atau Backspace di awal item keluar dari list (atau outdent satu level bila item ter-nest).
  6. Konversi bullet ke numbered dan sebaliknya mempertahankan seluruh isi dan struktur nesting.
  7. ⚠ List item dapat berisi paragraf, code block, dan nested list; media di dalam list item diizinkan oleh schema dan di-render.
  8. Paste list dari Word/Google Docs mempertahankan struktur list dan nesting (lihat EDT-025).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Action item (EDT-016) adalah tipe list terpisah (`taskList`).

### EDT-007 Code Snippet (Code Block)

- **Deskripsi:** Blok kode (Code Snippet) dengan pemilih bahasa, syntax highlighting di mode baca, nomor baris, dan tombol salin.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Editor konten (pengguna); viewer memakai tombol Copy di renderer.
- **Aturan perilaku (acceptance criteria):**
  1. Code block disisipkan lewat toolbar, '/code', atau mengetik ``` (opsional diikuti nama bahasa, mis. ```js, yang langsung mengatur bahasa).
  2. Language selector menyediakan pencarian dan daftar bahasa umum (Java, JavaScript, TypeScript, Go, Python, SQL, JSON, YAML, Bash, dan lainnya) plus 'None'.
  3. Syntax highlighting diterapkan di mode baca memakai pustaka highlight sisi klien tanpa `eval`.
  4. Renderer menampilkan nomor baris di sisi kiri code block.
  5. Renderer menyediakan tombol Copy code yang menyalin isi mentah tanpa nomor baris.
  6. ⚠ Tersedia toggle wrap baris panjang; default tanpa wrap dengan scroll horizontal.
  7. Tab di dalam code block menyisipkan indentasi dan tidak memindahkan fokus keluar editor.
  8. Whitespace dipertahankan persis saat paste kode multi-baris; pengguna keluar dari block dengan panah bawah di baris terakhir atau Enter berulang (tiga kali) pada baris kosong terakhir.
  9. Code block dapat dihapus lewat floating toolbar (ikon trash).
  10. Code block tidak mendukung mark formatting di dalamnya; mark pada teks yang di-paste ke dalam code block dibuang.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi ke `{code:lang}` wiki markup di EDT-028.

### EDT-008 Quote (Blockquote)

- **Deskripsi:** Blok kutipan (Quote) untuk mengutip teks, termasuk mengutip isi comment orang lain.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Quote disisipkan lewat menu '+', '/quote', mengetik '> ' di awal baris, atau shortcut Ctrl/Cmd+Alt+9.
  2. Blockquote dapat berisi paragraf, list, dan code block.
  3. Menekan Enter pada baris kosong di dalam quote keluar dari quote.
  4. Pengguna dapat mengutip comment orang lain dengan menyalin teks lalu menjadikannya quote; konten yang di-paste ke quote mempertahankan format yang diizinkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur "reply/quote comment" otomatis (bila ada) diatur di modul 08.

### EDT-009 Divider (Horizontal Rule)

- **Deskripsi:** Garis pemisah horizontal (Divider) antar bagian konten.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Divider disisipkan lewat menu '+', '/divider', mengetik '---', atau shortcut Ctrl/Cmd+Shift+-.
  2. Divider dapat diseleksi (node selection) dan dihapus seperti node blok lain.
  3. ⚠ Schema tidak mengizinkan divider di dalam sebagian container (termasuk table cell); perintah divider dinonaktifkan di lokasi yang tidak diizinkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditempatkan di Fase 2 walau P2 karena biaya implementasinya sangat kecil dan diperlukan oleh markdown shortcut EDT-022.

## Elemen blok lanjutan

### EDT-010 Tables

- **Deskripsi:** Tabel (Tables) dengan header, operasi baris/kolom, merge/split, warna sel, resize, dan sorting kolom.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Tabel disisipkan (default 3x3) lewat toolbar, '/table', atau shortcut Shift+Alt+T; tombol toolbar menyediakan size picker berbentuk grid.
  2. Floating toolbar 'Table options' menyediakan toggle header row (aktif secara default), header column, dan numbered column.
  3. ⚠ Baris dan kolom dapat ditambah/dihapus lewat kontrol di tepi tabel, context menu sel, atau shortcut keyboard (kombinasi tombol ditetapkan dan didokumentasikan di dialog shortcut EDT-024).
  4. Sel yang diseleksi dapat di-merge, dan sel hasil merge dapat di-split kembali.
  5. Cell background color dipilih dari palet tetap.
  6. Lebar kolom dapat di-resize dengan drag, dan perintah distribute columns menyamakan lebar kolom terpilih.
  7. ⚠ Layout lebar tabel tersedia dalam opsi normal, wide, dan full width; opsi yang tidak dibangun di UI tetap diterima dari API.
  8. Menu header kolom menyediakan sort ascending/descending; sort dinonaktifkan (dengan tooltip alasan) bila tabel berisi merged cell.
  9. Tersedia aksi delete table, clear cells, delete row, dan delete column.
  10. Tab/Shift+Tab berpindah antar sel; Tab di sel terakhir menambah baris baru.
  11. Paste tabel dari Excel, Google Sheets, atau HTML dikonversi menjadi tabel ADF (lihat EDT-025).
  12. ⚠ Header row menempel (sticky) saat menggulir tabel panjang di editor dan renderer.
  13. Sel dapat berisi paragraf, heading, list, code block, panel, media, mention, dan nested expand; tabel di dalam tabel tidak didukung dan perintah table dinonaktifkan di dalam sel.
  14. Tabel yang lebih lebar dari area konten di-scroll secara horizontal di renderer tanpa merusak layout halaman.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi ke `||header||`/`|cell|` wiki markup di EDT-028; render tabel di email di EDT-003.

### EDT-011 Panels

- **Deskripsi:** Blok panel berwarna (Panels) untuk info, note, success, warning, dan error, dengan ikon sesuai tipenya.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Tipe panel yang tersedia: Info (biru), Note (ungu), Success (hijau), Warning (kuning), Error (merah), masing-masing dengan ikon bawaan.
  2. Panel disisipkan lewat menu '+' atau slash command '/info', '/note', '/success', '/warning', '/error', '/panel'.
  3. Floating toolbar panel menyediakan penggantian tipe panel tanpa kehilangan isi.
  4. ⚠ Custom panel dengan warna latar dari palet tetap dan ikon emoji pilihan (atau tanpa ikon) tersedia; bila tidak dibangun di UI, panel custom tetap diterima dari API dan di-render.
  5. Panel dapat berisi heading, paragraf, list, code block, dan action item.
  6. Panel dapat dihapus lewat floating toolbar.
  7. Pengguna keluar dari panel dengan panah bawah atau Enter di baris kosong terakhir.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi ke `{panel}` wiki markup di EDT-028.

### EDT-012 Expand

- **Deskripsi:** Blok collapsible berjudul (Expand) untuk menyembunyikan konten panjang seperti log atau detail teknis.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Editor konten (pengguna); viewer membuka/menutup di mode baca.
- **Aturan perilaku (acceptance criteria):**
  1. Expand disisipkan lewat menu '+' atau '/expand'.
  2. Judul expand dapat diedit, dengan placeholder 'Give this expand a title...' bila kosong.
  3. State buka/tutup di mode baca berlaku per viewer dan tidak disimpan ke dokumen maupun memicu perubahan field.
  4. Expand dapat berisi tabel, panel, code block, dan media; di dalam table cell dipakai varian `nestedExpand` yang tidak boleh berisi expand lain.
  5. Expand dapat dihapus lewat floating toolbar.
  6. ⚠ Expand tertutup secara default di renderer dan terbuka saat baru dibuat di editor.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fallback di kanal email/plaintext menampilkan judul dan isi tanpa interaksi (EDT-003).

### EDT-013 Layouts (Columns)

- **Deskripsi:** Tata letak multi-kolom (Layouts) di konten rich text. Dukungannya di editor Description belum dikonfirmasi sebagai perilaku standar, sehingga fitur ini hanya dibangun setelah keputusan produk.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Layout disisipkan lewat '/layout' di field rich text bila fitur diaktifkan; bila tidak diaktifkan, node `layoutSection` dari API tetap di-render dan dipertahankan saat disimpan ulang.
  2. Pilihan layout: 2 kolom, 3 kolom, serta kolom dengan sidebar kiri atau kanan.
  3. Setiap kolom dapat berisi konten bebas (paragraf, list, media) sesuai schema.
  4. Layout menjadi satu kolom (kolom ditumpuk) pada layar sempit/mobile.
  5. Menghapus layout mempertahankan isi semua kolom sebagai blok berurutan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Jangan dibangun UI-nya sebelum dikonfirmasi (lihat Pertanyaan terbuka); schema tetap mendukung agar impor tidak gagal.

## Elemen inline & interaktif

### EDT-014 Mentions

- **Deskripsi:** Menyebut pengguna dengan '@' (Mentions) di description, comment, atau field Paragraph agar mereka menerima notifikasi. Modul ini mencakup typeahead, penyimpanan node mention, parsing mention dari ADF, dan pemicu event; pengiriman notifikasinya ada di modul 19.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna yang menulis (memerlukan global permission Browse users and groups untuk typeahead); user yang di-mention sebagai penerima notifikasi.
- **Aturan perilaku (acceptance criteria):**
  1. Mengetik '@' membuka typeahead pengguna yang mencari berdasarkan nama tampilan, atau email hanya bila pengaturan privasi profil user tersebut mengizinkan.
  2. Saran diurutkan dengan kolaborator relevan lebih dulu: assignee, reporter, commenter, watcher, lalu user yang baru saja berinteraksi dengan penulis.
  3. Mention disimpan sebagai node `mention` berisi `accountId` dan teks tampilan; renderer menampilkan nama terbaru user (bukan teks tersimpan) saat konten dibaca.
  4. Notifikasi 'mentioned you' (in-app/email) dikirim saat konten disimpan, bukan saat diketik; event domain `mention.created` dipancarkan per user yang di-mention.
  5. Bila user yang di-mention tidak memiliki akses ke work item, editor menampilkan peringatan (mis. "won't be notified, they don't have access"); mention tetap tersimpan, tetapi notifikasi tidak dikirim.
  6. Typeahead pengguna hanya tersedia bagi user dengan global permission Browse users and groups; tanpa permission tersebut '@' tidak membuka daftar user.
  7. ⚠ Bila undangan pengguna diizinkan oleh admin, typeahead menampilkan opsi mengundang alamat email yang belum menjadi user (alur undangan di modul 02).
  8. Mention ke user yang dinonaktifkan tampil dengan gaya berbeda (abu-abu) dan tidak memicu notifikasi.
  9. Mengklik mention membuka profile card berisi nama, jabatan, waktu lokal, dan tautan profil, mengikuti pengaturan privasi profil.
  10. ⚠ Mention team (entitas Teams, PLAN-057 di [18 Timeline & Rencana](18-timeline-rencana-tampilan-proyek.md)) menotifikasi seluruh anggota team bila fitur Teams aktif; bila tidak aktif, entri team tidak muncul di typeahead.
  11. Mention di comment yang visibilitasnya dibatasi (role/group) tidak menotifikasi user di luar role/group tersebut.
  12. ⚠ Saat comment atau description diedit, notifikasi hanya dikirim untuk mention baru (selisih `accountId` antara versi lama dan baru).
  13. Mention terhadap diri sendiri tidak memicu notifikasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Notifikasi di [19 Notifikasi](19-notifikasi-email.md); visibilitas comment di modul 08; permission Browse users and groups di modul 03; profil dan privasi di modul 01/02. Endpoint typeahead harus dibatasi rate dan tidak membocorkan email.

### EDT-015 Emoji

- **Deskripsi:** Penyisipan emoji lewat picker dan autocomplete shortcode, termasuk custom emoji site yang diunggah pengguna. Emoji standar (aturan 1-6 dan 13) dibangun di Fase 2; custom emoji (aturan 7-12) boleh ditunda ke Fase 3 dengan prioritas internal P3 sesuai catatan inventaris.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna); Admin Situs untuk kebijakan dan moderasi custom emoji.
- **Aturan perilaku (acceptance criteria):**
  1. Emoji picker menampilkan kategori Frequent, People, Nature, Food, Activity, Travel, Objects, Symbols, Flags, Productivity, dan Custom.
  2. Picker menyediakan pencarian emoji berdasarkan nama/kata kunci.
  3. Mengetik ':' diikuti minimal satu huruf (mis. ':smile') membuka autocomplete emoji yang dapat dinavigasi dengan keyboard.
  4. Skin tone selector tersedia untuk emoji yang mendukung variasi warna kulit, dan pilihan diingat per user.
  5. ⚠ Emoticon ASCII (mis. ':)') dikonversi otomatis menjadi emoji saat diikuti spasi, dan Ctrl/Cmd+Z mengembalikannya.
  6. Daftar 'Frequently used' disimpan per user berdasarkan pemakaian.
  7. Pengguna dapat mengunggah gambar sebagai custom emoji yang dipakai bersama di seluruh site, dan dapat menghapus custom emoji miliknya.
  8. ⚠ Tombol 'Add your own emoji' di picker membuka alur: pilih file, isi shortname unik (`:nama:`), lihat preview, lalu simpan; shortname yang sudah dipakai ditolak dengan pesan galat.
  9. Custom emoji tampil di kategori 'Custom'/'Your uploads' pada picker dan dapat dipakai di description, comment, dan comment reactions (modul 08).
  10. ⚠ Custom emoji dapat dihapus oleh pengunggah atau Admin Situs; konten lama yang memakai emoji terhapus menampilkan fallback shortname/teks.
  11. ⚠ Admin Situs dapat membatasi siapa yang boleh mengunggah custom emoji (semua user atau grup tertentu) di pengaturan global.
  12. ⚠ File custom emoji dibatasi format PNG/JPG/GIF dengan ukuran maksimum yang dapat dikonfigurasi (usulan 1 MB) dan di-resize otomatis ke ukuran standar.
  13. Emoji disimpan di ADF sebagai node `emoji` dengan atribut `shortName`, `id`, dan `text` fallback.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan file custom emoji memakai storage di modul 08 (CMT-019 File Storage); pengaturan admin di [22 Administrasi Sistem](22-administrasi-sistem.md). Layanan emoji lintas produk milik vendor di luar cakupan; produk ini memiliki katalog custom emoji sendiri per site.

### EDT-016 Action Items (Task List) and Decisions

- **Deskripsi:** Checklist (action item) dengan checkbox yang dapat dicentang langsung di konten tersimpan, plus decision item bila diputuskan didukung. Action item bukan work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna); pengguna dengan Edit Issues untuk mencentang di description; penulis comment atau pemegang Edit All Comments untuk mencentang di comment.
- **Aturan perilaku (acceptance criteria):**
  1. Action item disisipkan lewat toolbar, menu '+', '/action', atau mengetik '[] ' di awal baris.
  2. Checkbox dapat dicentang/dilepas di mode baca tanpa masuk mode edit; perubahan langsung disimpan sebagai update field melalui endpoint yang mengubah `attrs.state` (TODO/DONE) secara atomik berdasarkan `localId`, sehingga tidak menimpa edit description lain yang berjalan bersamaan.
  3. Perubahan checkbox di description dicatat di History sebagai perubahan Description (nilai lama dan baru).
  4. Action item dapat di-nest dengan Tab.
  5. Action item dapat berisi mention dan date; date lampau pada item yang belum dicentang ditampilkan merah (overdue).
  6. ⚠ Di description hanya user dengan Edit Issues yang dapat mencentang checkbox; di comment hanya penulis comment atau pemegang Edit All Comments; user lain melihat checkbox nonaktif.
  7. ⚠ Decision item (shortcut '<> ') tersedia sebagai node `decisionList`/`decisionItem`; bila tidak dibangun di UI, node tetap di-render dari API.
  8. Action item tidak menjadi sub-task atau work item, dan tidak ada agregasi 'My tasks' lintas work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** History di CMT-022/CMT-023 (modul 08); permission di modul 03. Konversi action item ke subtask (bila diinginkan) adalah keputusan terbuka di luar inventaris.

### EDT-017 Status Lozenge Element

- **Deskripsi:** Elemen inline berupa label status berwarna dengan teks bebas, murni dekoratif dan tidak terhubung dengan workflow.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. Status disisipkan lewat menu '+' atau '/status'.
  2. Teks status bebas dan di-render dalam huruf kapital.
  3. Pilihan warna: neutral (abu-abu), purple, blue, red, yellow, green.
  4. Mengklik elemen di editor membuka popup untuk mengubah teks dan warna.
  5. Elemen bersifat statis dan tidak terhubung dengan status workflow work item.
  6. Disimpan sebagai node `status` ADF dengan atribut `text`, `color`, dan `localId`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Status workflow sesungguhnya di modul 09; lozenge status work item pada smart link di EDT-020.

### EDT-018 Date Element

- **Deskripsi:** Elemen inline tanggal yang dipilih dari date picker dan tampil sebagai lozenge.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Date disisipkan lewat menu '+', '/date', atau mengetik '//'.
  2. Date picker menyediakan kalender dan input tanggal berupa teks.
  3. Format tampilan mengikuti locale viewer (bukan locale penulis).
  4. Di dalam action item yang belum dicentang, tanggal lampau ditampilkan merah (overdue).
  5. Disimpan sebagai timestamp UTC (milidetik) di `attrs.timestamp` node `date`.
  6. Elemen date tidak terhubung dengan field Due date work item.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Locale dan zona waktu viewer dari preferensi di modul 21.

## Media & tautan

### EDT-019 Media in Editor (Images and Files)

- **Deskripsi:** Menyisipkan gambar, video, dan file langsung ke description atau comment; file yang disisipkan otomatis menjadi attachment work item. Node media merujuk id attachment, dan renderer membuat URL ber-token untuk mengakses file.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan permission Create Attachments (menyisipkan); viewer dengan Browse Projects (melihat); editor konten.
- **Aturan perilaku (acceptance criteria):**
  1. Media disisipkan lewat tombol media di toolbar, '/image' atau '/file', drag-drop ke editor, atau paste dari clipboard (screenshot).
  2. Gambar tampil inline; file non-gambar tampil sebagai kartu file (nama, ikon tipe, ukuran); video tampil sebagai player inline.
  3. Media yang ditambahkan lewat editor otomatis tercatat di panel Attachments work item.
  4. Gambar dapat di-resize dengan drag handle; lebar tersimpan di atribut node.
  5. ⚠ Layout/alignment media tersedia: align left/center/right, wrap left/right, wide, dan full width; opsi yang tidak dibangun di UI tetap diterima dari API.
  6. Gambar dapat diberi alt text untuk aksesibilitas, yang dipakai renderer sebagai atribut `alt`.
  7. ⚠ Gambar dapat diberi caption di bawahnya.
  8. Gambar dapat dibungkus link (klik gambar di mode baca membuka URL tersebut, bukan Media Viewer).
  9. ⚠ Gambar dapat diberi border dengan warna dari palet tetap dan ketebalan terbatas (1-3 px).
  10. Media dapat di-toggle antara inline (`mediaInline`) dan block (`mediaSingle`).
  11. Mengklik media di mode baca membuka Media Viewer (modul 08).
  12. Selama upload ditampilkan placeholder progres; upload yang gagal menampilkan pesan galat dan opsi retry, dan node yang gagal tidak tersimpan.
  13. ⚠ Gambar eksternal lewat URL disimpan sebagai external media (tanpa attachment) dan diambil lewat proxy/SSRF guard atau ditampilkan langsung sesuai kebijakan CSP.
  14. Menghapus media dari editor tidak menghapus attachment dari work item.
  15. ⚠ Media di comment yang visibilitasnya dibatasi tetap menjadi attachment work item yang terlihat di panel Attachments bagi semua viewer work item.
  16. Tanpa permission Create Attachments, tombol media, slash command media, drag-drop, dan paste file disembunyikan atau diblokir dengan pesan galat.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Penyimpanan berkas, batas ukuran upload, dan Media Viewer di [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md). URL ber-token berumur pendek untuk menghindari kebocoran akses. Aturan 15 adalah perilaku yang berpotensi membocorkan konten dan perlu keputusan produk (lihat Pertanyaan terbuka).

### EDT-020 Smart Link Unfurl and Auto-linking of Issue Keys

- **Deskripsi:** Link ke work item lain di site yang sama di-unfurl dengan key, summary, ikon work type, dan status terkini; issue key dikenali sebagai link. Ini adalah resolver internal pertama untuk smart link (EDT-021).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna (penulis) dan viewer konten; metadata di-resolve sesuai permission viewer.
- **Aturan perilaku (acceptance criteria):**
  1. URL work item di site yang sama di-unfurl menjadi inline card berisi ikon work type, key, summary, dan lozenge status yang selalu mengikuti status terbaru saat dibaca (tidak di-snapshot saat disimpan).
  2. ⚠ Issue key (mis. ABC-123) di teks dikenali sebagai link ke work item di renderer; konversi otomatis menjadi inline card saat diketik di editor bersifat opsional dan dapat dimatikan.
  3. Work item yang tidak dapat diakses viewer ditampilkan tanpa metadata (hanya URL/key), tanpa membocorkan summary atau status.
  4. Issue key di comment/description tidak otomatis membuat Issue Link formal (Issue Link diatur di modul 06).
  5. ⚠ Unfurl URL generik (di luar work item internal) memakai resolver OpenGraph/oEmbed sisi server di EDT-021; bila resolver tidak tersedia, URL tampil sebagai link biasa.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Resolver internal memakai endpoint batch (banyak key sekaligus) agar issue view tidak melakukan N+1 request. Pengenalan issue key di pesan commit ada di modul 24.

### EDT-021 Smart Links (Inline, Card, Embed)

- **Deskripsi:** URL yang di-paste dikonversi menjadi smart link yang menampilkan metadata sumbernya (judul, ikon, status) dengan mode tampilan inline, card, dan embed, termasuk embed video dan konten eksternal. Produk ini membangun resolver internal dan resolver generik OpenGraph/oEmbed; resolver pihak ketiga ber-OAuth bersifat opsional.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Pengguna (penulis), viewer konten, Admin Situs/Org admin (kebijakan preview/embed).
- **Aturan perilaku (acceptance criteria):**
  1. Paste URL otomatis menjadi inline smart link (ikon + judul) bila ada resolver yang cocok; Ctrl/Cmd+Z tepat setelah paste mengembalikannya ke URL biasa.
  2. Floating toolbar menyediakan mode tampilan: URL, Inline, Card (block card dengan deskripsi), dan Embed (preview iframe).
  3. Hover card (preview saat hover) menampilkan ringkasan, status, pemilik, dan waktu update bila tersedia dari resolver.
  4. Link ke entitas native produk ini (work item, proyek, board, filter, dashboard, goal bila modul Goals tersedia, repositori/PR yang terhubung lewat integrasi DevOps modul 24) menampilkan metadata native lewat resolver internal; link ke produk vendor lintas produk (halaman wiki/whiteboard, rekaman video, goals, repositori vendor) tidak memiliki resolver native dan diperlakukan sebagai URL generik (OpenGraph, aturan 5 dan EDT-020 aturan 5).
  5. Link pihak ketiga (GitHub, GitLab, Figma, Google Drive, Dropbox, OneDrive, Slack, Miro, dll) dapat di-resolve lewat resolver opsional yang meminta autentikasi akun pengguna ('Connect your account'); tanpa resolver, fallback ke OpenGraph generik.
  6. Embed tersedia untuk domain yang diizinkan (YouTube, Vimeo, Loom, Figma, Google Docs/Sheets/Slides, Miro, dll) lewat oEmbed/iframe ber-sandbox; domain yang tidak didukung embed kembali ke card atau inline.
  7. Embed dapat di-resize (lebar/tinggi, aspect ratio) dan diatur layout-nya (center, wide, full width).
  8. Bila viewer tidak memiliki akses ke resource, smart link menampilkan status 'You don't have access' dan, bila resolver mendukung, tombol 'Request access'.
  9. Link yang tidak ditemukan atau gagal di-resolve (timeout, 4xx/5xx) kembali ditampilkan sebagai URL biasa.
  10. Floating toolbar smart link menyediakan Edit link (URL dan display text), Open link, Unlink, Copy link, dan Settings (link preferences).
  11. Smart link hanya ada di field rich text; field single-line seperti Summary tetap plain text.
  12. Metadata di-resolve di sisi viewer (per permintaan viewer dengan kredensial viewer), sehingga setiap viewer melihat data sesuai izinnya sendiri; cache metadata dipisah per viewer untuk resource berizin.
  13. Link preferences per user (tampilan default inline atau URL, dan opsi preview) disimpan di pengaturan profil (modul 21).
  14. ⚠ Admin Situs dapat menonaktifkan preview/embed secara global atau membatasi daftar domain embed lewat pengaturan organisasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fetch metadata wajib lewat SSRF guard (EDT-002 aturan 8, modul 25). Iframe embed memakai atribut `sandbox` dan CSP `frame-src` dari allowlist. Resolver pihak ketiga ber-OAuth bergantung pada integrasi di modul 24/25; layanan linking platform milik vendor di luar cakupan.

## Akselerator input: markdown, slash command, shortcut & paste

### EDT-022 Markdown Shortcuts and Autoformatting

- **Deskripsi:** Sintaks markdown yang diketik dikonversi otomatis menjadi format editor (Markdown Shortcuts and Autoformatting), dengan undo yang mengembalikan teks asli.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna).
- **Aturan perilaku (acceptance criteria):**
  1. '**teks**' atau '__teks__' menjadi bold; '*teks*' atau '_teks_' menjadi italic; '~~teks~~' menjadi strikethrough; '`teks`' menjadi inline code, saat penanda penutup diketik.
  2. '# ' sampai '###### ' di awal baris menjadi Heading 1-6.
  3. '* ', '- ', '+ ' di awal baris menjadi bullet list; '1. ' menjadi numbered list.
  4. '[] ' di awal baris menjadi action item.
  5. '> ' di awal baris menjadi quote.
  6. '```' menjadi code block.
  7. '---' di baris sendiri menjadi divider.
  8. ⚠ '[teks](https://alamat)' menjadi link dengan display text 'teks', setelah URL divalidasi sesuai EDT-004.
  9. ⚠ Autoformat tipografi: '->' dan '<-' menjadi panah, '--' menjadi em dash, dan tanda kutip lurus menjadi smart quotes; tidak berlaku di dalam inline code dan code block.
  10. Ctrl/Cmd+Z tepat setelah autoformat mengembalikan teks markdown asli tanpa format.
  11. Paste teks Markdown dikonversi menjadi rich text (detail di EDT-025).
  12. Paste as plain text (Ctrl/Cmd+Shift+V) mengabaikan semua format dan markdown.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Konversi terjadi di klien; hasilnya divalidasi ulang server (EDT-002 aturan 5). Pengganti mode teks wiki markup (EDT-029).

### EDT-023 Slash Commands (Quick Insert)

- **Deskripsi:** Mengetik '/' membuka menu cepat (Quick Insert) untuk menyisipkan elemen editor lewat pencarian.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna); daftar elemen disaring oleh permission pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. '/' di awal baris atau setelah spasi membuka typeahead quick insert; '/' di tengah kata (mis. URL) tidak membukanya.
  2. Pencarian kata kunci mencocokkan nama dan alias elemen (mis. /table, /code, /info, /date, /status, /expand, /divider, /action, /emoji, /mention, /image, /heading, /quote).
  3. Typeahead dapat dinavigasi dengan panah atas/bawah, Enter memilih, dan Esc menutup tanpa menyisipkan apa pun (teks '/' tetap ada).
  4. Entri 'View more' membuka dialog 'Browse all elements' yang mengelompokkan elemen per kategori dengan pencarian.
  5. Elemen AI (mis. '/ai') hanya muncul bila asisten AI diaktifkan untuk site dan user (EDT-030).
  6. Elemen yang tidak diizinkan untuk konteks atau user (mis. media tanpa Create Attachments, divider di dalam sel tabel) tidak ditampilkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Registri elemen editor bersama dipakai oleh toolbar (EDT-001), menu '+', dan slash command agar daftarnya konsisten.

### EDT-024 Editor Keyboard Shortcuts

- **Deskripsi:** Pintasan keyboard untuk operasi editor beserta dialog bantuan shortcut.
- **Prioritas / Fase / Tier:** P2 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. ⚠ Dialog bantuan 'Keyboard shortcuts' editor terbuka lewat Ctrl/Cmd+/ saat fokus di dalam editor, dan juga dapat dibuka dari menu bantuan.
  2. Tersedia shortcut untuk format teks, heading, list, link, code, table, quote, divider, undo/redo, dan clear formatting sesuai yang tercantum di EDT-001 sampai EDT-011; dialog bantuan juga mencantumkan daftar markdown shortcut EDT-022 (pengganti halaman 'Text formatting notation', lihat EDT-029).
  3. Ctrl/Cmd+Enter menyimpan dan Esc membatalkan di editor comment/description.
  4. Shift+Enter menyisipkan line break.
  5. Tab/Shift+Tab melakukan indent/outdent list dan berpindah antar sel tabel.
  6. Label shortcut ditampilkan dalam varian Mac (Cmd/Option) atau Windows/Linux (Ctrl/Alt) sesuai platform yang terdeteksi.
  7. Shortcut global issue view 'M' (saat fokus tidak berada di field input) membuka dan memfokuskan comment editor; selama fokus berada di dalam editor, shortcut global satu huruf tidak aktif agar tidak bentrok dengan pengetikan. Katalog lengkap di [21 Navigasi](21-navigasi-ux-preferensi.md) (NAV-041).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Ditempatkan di Fase 2 walau P2 karena sebagian besar shortcut sudah melekat pada elemen Fase 1-2; yang tersisa hanyalah dialog bantuan.

### EDT-025 Paste and Clipboard Handling

- **Deskripsi:** Konten yang di-paste dari berbagai sumber dikonversi menjadi ADF yang valid dan aman.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Editor konten (pengguna); paste file/gambar memerlukan Create Attachments.
- **Aturan perilaku (acceptance criteria):**
  1. Paste dari Microsoft Word, Google Docs, Excel/Sheets, dan HTML web mempertahankan heading, list, tabel, link, serta bold/italic; format di luar schema dibuang.
  2. Paste antar field atau antar work item di produk ini mempertahankan node ADF (mention, panel, status) lewat format clipboard ADF khusus di samping HTML.
  3. Paste URL menjadi smart link (EDT-021); paste URL di atas teks yang diseleksi membuat link pada teks tersebut.
  4. Paste gambar/screenshot diunggah sebagai attachment lalu disisipkan sebagai node media (EDT-019).
  5. Paste file dari file manager diunggah sebagai attachment dan disisipkan sebagai kartu file.
  6. Paste di dalam code block selalu menjadi plain text dengan whitespace dipertahankan.
  7. Paste teks markdown dikonversi menjadi rich text (heading, list, code block, link, tabel GFM).
  8. Paste as plain text (Ctrl/Cmd+Shift+V) menyisipkan teks tanpa format.
  9. HTML yang di-paste disanitasi: elemen `script`, `style`, `iframe`, dan atribut event (`on*`) dibuang sebelum konversi ke ADF.
  10. Paste konten sangat besar dibatasi limit karakter field (EDT-001 aturan 13); bila melampaui, paste ditolak dengan pesan galat dan isi editor tidak berubah.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Sanitasi klien bukan pengganti validasi server (EDT-002). Paste lintas produk vendor lain hanya didukung lewat HTML.

## Field rich text & kompatibilitas renderer

### EDT-026 Description Field Behaviors

- **Deskripsi:** Perilaku khusus field Description sebagai konten rich text utama work item: penempatan, default/template, kewajiban isi, riwayat, kloning, pengindeksan, dan penanganan konflik edit.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues (mengubah); Admin Situs (field configuration company-managed); Admin Proyek (layout dan pengaturan field team-managed).
- **Aturan perilaku (acceptance criteria):**
  1. Description tampil di kolom utama issue view di bawah Summary; dapat dihapus dari layout (team-managed) atau dari screen (company-managed), dan bila dihapus tidak tampil di create/edit/view.
  2. ⚠ Default value/template Description dapat diatur per field context (company-managed) dan per work type (team-managed); template diterapkan saat create hanya bila Description kosong.
  3. Description dapat dibuat required lewat field configuration (company-managed) atau pengaturan field per work type (team-managed); menyimpan Description kosong (termasuk dokumen ADF yang hanya berisi paragraf kosong) ditolak dengan pesan galat.
  4. Perilaku Description (required, hidden, renderer) dapat berbeda per work type lewat field configuration scheme (company-managed).
  5. Setiap perubahan Description tercatat di History dengan nilai lama dan baru (representasi plaintext/diff untuk tampilan, ADF untuk API).
  6. ⚠ Description panjang di renderer dipotong pada tinggi tertentu dengan tombol 'Show more'/'Show less'.
  7. Description ikut tersalin saat Clone work item, termasuk referensi media; attachment yang dirujuk ikut disalin bila opsi salin attachment dipilih, dan referensi ke attachment yang tidak disalin diubah menjadi fallback.
  8. Description terindeks untuk pencarian teks (JQL `description ~` dan `text ~`) melalui kolom plaintext turunan.
  9. ⚠ Tidak ada co-editing real-time; penyimpanan memakai optimistic concurrency (versi field), dan bila versi berubah sejak editor dibuka, pengguna diberi tahu dan diminta memilih antara memuat ulang atau menimpa (last-write-wins hanya dengan konfirmasi eksplisit).
- **Team-managed vs company-managed:** Company-managed mengatur default, required, dan renderer lewat field configuration/context; team-managed mengatur lewat layout dan pengaturan field per work type, dan Description selalu rich text.
- **Catatan & ketergantungan:** Konfigurasi field di [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) (CFG-043/CFG-044 field configuration, CFG-052 template description per work type); Clone di ISS-031 (modul 06) dan History di CMT-022/CMT-023 (modul 08); indeks di modul 13. Migrasi repo: tabel work item belum ada; kolom `description_adf JSONB`, `description_text TEXT`, dan `description_version` didefinisikan di modul 28.

### EDT-027 Rich Text in Custom Fields

- **Deskripsi:** Dukungan editor untuk custom field Paragraph (multi-line text) dan field sistem Environment, serta batasan field teks single-line.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna dengan Edit Issues (mengisi); Admin Situs (renderer dan konfigurasi field).
- **Aturan perilaku (acceptance criteria):**
  1. Custom field 'Paragraph' (Text Field multi-line) memakai editor rich text ADF (mode full editor).
  2. Di company-managed, renderer di field configuration menentukan apakah field tampil sebagai rich text atau plain text (EDT-028).
  3. Di team-managed, field Paragraph selalu rich text.
  4. Field Short text (single line) tidak mendukung rich text dan dibatasi 255 karakter; melampaui batas ditolak dengan pesan galat.
  5. Field sistem Environment mendukung rich text dengan aturan yang sama seperti Paragraph.
  6. ⚠ Mention di custom field Paragraph mengirim notifikasi seperti mention di Description (EDT-014).
  7. ⚠ Media di custom field Paragraph juga menjadi attachment work item (EDT-019).
- **Team-managed vs company-managed:** Company-managed dapat memilih renderer per field; team-managed selalu rich text untuk Paragraph.
- **Catatan & ketergantungan:** Tipe custom field di modul 05; penyimpanan nilai custom field (JSONB) di modul 28.

### EDT-028 Legacy Wiki Style Renderer (Field Renderers)

- **Deskripsi:** Pilihan renderer per field di company-managed project, yaitu Wiki Style Renderer (rich text) atau Default Text Renderer (plain text), beserta kompatibilitas wiki markup untuk API v2, impor, dan otomasi. Cukup disediakan konverter ADF ke/dari wiki markup minimal.
- **Prioritas / Fase / Tier:** P3 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (field configuration); Integrator/Developer API (REST API v2); sistem impor dan otomasi.
- **Aturan perilaku (acceptance criteria):**
  1. Renderer dipilih per field di Field configuration (company-managed) dengan opsi 'Wiki Style Renderer' atau 'Default Text Renderer'; default untuk field baru adalah Wiki Style Renderer.
  2. Pilihan renderer berlaku untuk field yang mendukung renderer: Description, Comment, Environment, custom field Paragraph (multi-line), dan field text lain yang ditandai mendukung renderer.
  3. Di issue view, field dengan Wiki Style Renderer memakai editor rich text ADF, sedangkan field dengan Default Text Renderer tampil dan diedit sebagai plain text tanpa formatting (textarea).
  4. Konverter mengenali sintaks wiki markup: `*bold*`, `_italic_`, `+underline+`, `-strike-`, `{{monospace}}`, `h1.`-`h6.`, list `*` dan `#`, tabel `||header||` dan `|cell|`, `[teks|url]`, `{code:java}`, `{noformat}`, `{quote}`, `bq.`, `{panel:title=..}`, `{color:red}`, `!image.png!`, `[~accountid:...]`, `----`, dan emoticon `:)` `(/)` `(x)`.
  5. Konversi dua arah wiki markup ke ADF dan ADF ke wiki markup tersedia untuk REST API v2; elemen ADF tanpa padanan wiki markup dikonversi ke representasi teks terbaik.
  6. Wiki markup tetap diterima di REST API v2, impor CSV, dan aksi otomasi yang menerima teks; semuanya dikonversi ke ADF tersanitasi sebelum disimpan (EDT-002).
  7. Team-managed project tidak memiliki field configuration/renderer; field Paragraph selalu rich text.
  8. Field text single-line tidak mendukung rich text, terlepas dari renderer.
- **Team-managed vs company-managed:** Hanya berlaku di company-managed; team-managed selalu rich text.
- **Catatan & ketergantungan:** Konfigurasi renderer di modul 05; REST API v2 di modul 23; impor CSV di modul 25; otomasi di modul 20. Mengganti renderer pada field yang sudah berisi data tidak mengubah ADF tersimpan, hanya cara tampil dan edit.

### EDT-029 Wiki Markup Text Mode (Visual/Text Editor Toggle)

- **Deskripsi:** Mode editor lama dengan tab Visual/Text untuk menulis wiki markup langsung dan melihat preview; hanya ada di edisi Data Center. Dicatat sebagai rujukan dan tidak dibangun.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Pengguna edisi Data Center (rujukan).
- **Aturan perilaku (acceptance criteria):**
  1. Tidak dibangun: toggle tab 'Visual'/'Text' di editor field yang memakai Wiki Style Renderer.
  2. Tidak dibangun: mode Text untuk input wiki markup mentah dengan tombol preview.
  3. Tidak dibangun: halaman bantuan 'Text formatting notation'; sebagai gantinya dialog shortcut (EDT-024) mencantumkan markdown shortcut.
  4. Pengguna memakai markdown shortcut (EDT-022) dan paste markdown (EDT-025) sebagai pengganti mode teks, sesuai perilaku issue view Cloud modern.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Lihat bagian Di luar cakupan. Integrator yang butuh wiki markup memakai REST API v2 (EDT-028).

## Asisten AI di editor

### EDT-030 Atlassian Intelligence / Rovo in Editor

- **Deskripsi:** Fitur AI generatif opsional di editor dan comment untuk menulis, merapikan, dan merangkum konten (padanan asisten AI vendor; di produk ini disebut "asisten AI"). Dapat dimatikan sepenuhnya dan bergantung pada penyedia model yang dikonfigurasi admin.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Pengguna (memakai); Admin Situs/Org admin (mengaktifkan/menonaktifkan, memilih penyedia, kuota).
- **Aturan perilaku (acceptance criteria):**
  1. Asisten diakses lewat '/ai' atau tombol AI di toolbar editor, hanya bila diaktifkan.
  2. Tersedia aksi Improve writing, Fix spelling and grammar, Make shorter, Make longer, dan Change tone pada seleksi atau seluruh isi.
  3. Pengguna dapat brainstorm atau membuat draft description dari prompt bebas.
  4. Summarize comments menghasilkan ringkasan thread comment di section Activity (modul 08).
  5. Asisten dapat menyarankan child work item atau subtask dari isi description; pembuatannya tetap melalui alur create di modul 06 dengan konfirmasi pengguna.
  6. ⚠ Define terms menjelaskan istilah pada teks terseleksi.
  7. ⚠ Translate menerjemahkan konten ke bahasa pilihan pengguna.
  8. Hasil AI ditampilkan sebagai preview dengan opsi Insert, Replace, Try again, dan Discard; tidak ada perubahan pada dokumen sebelum pengguna memilih Insert/Replace.
  9. Admin Situs dapat menonaktifkan AI untuk seluruh site; ketersediaan bergantung plan dan kebijakan AI yang dikonfigurasi, dan saat nonaktif semua entry point AI disembunyikan.
  10. AI hanya memakai konten yang dapat diakses user pemanggil (permission dicek sebelum konteks dikirim ke model).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Strategi AI opsional dan penyedia model di modul 27; layanan AI milik vendor di luar cakupan. Kuota dan biaya per plan adalah keputusan terbuka.

## Pertanyaan terbuka

- Berapa lama draft lokal editor dipertahankan, dan apakah draft perlu disinkronkan di server agar tersedia lintas perangkat (EDT-001)?
- Apakah produk ini membangun UI untuk mark/atribut yang belum pasti tersedia di editor work item (highlight color, alignment paragraf, custom panel, caption dan border gambar, decision item), atau cukup menerima dan me-render-nya dari API (EDT-004, EDT-005, EDT-011, EDT-016, EDT-019)?
- Apakah layout multi-kolom (EDT-013) diaktifkan di field work item, atau hanya didukung di schema untuk kompatibilitas impor?
- Apakah media yang disisipkan di comment ber-visibilitas terbatas tetap terlihat di panel Attachments bagi semua viewer, atau visibilitas attachment mengikuti comment asalnya (EDT-019 aturan 15)?
- Apakah konflik edit Description memakai optimistic concurrency dengan dialog konfirmasi, atau cukup last-write-wins dengan notifikasi pembaruan (EDT-026)?
- Apakah issue key yang diketik di editor otomatis dikonversi menjadi inline card, atau hanya dikenali di renderer (EDT-020)?
- Resolver smart link pihak ketiga mana (jika ada) yang dibangun native, dan apakah embed iframe diaktifkan secara default (EDT-021)?
- Apakah custom emoji dibangun di Fase 2 bersama emoji standar atau ditunda, dan siapa yang boleh mengunggahnya secara default (EDT-015)?
- Penyedia model AI, kuota, dan plan mana yang mendapat asisten AI, serta apakah data konten boleh dikirim ke penyedia eksternal (EDT-030)?

## Di luar cakupan & ekosistem

- **Mode teks wiki markup (Visual/Text toggle) Data Center** — EDT-029 bertier dc-only-or-deprecated; hanya dicatat sebagai rujukan, digantikan oleh markdown shortcut dan paste markdown.
- **Co-editing real-time (multi-kursor) pada Description** — tidak dibangun; penanganan konflik memakai EDT-026.
- **Layanan lintas produk milik vendor**: layanan emoji bersama lintas produk, linking platform/object resolver vendor, profil dan Teams vendor, wiki/whiteboard vendor, rekaman video vendor, Goals vendor, serta asisten AI vendor. Produk ini menyediakan padanan native minimal (katalog custom emoji per site, resolver internal + OpenGraph/oEmbed, asisten AI opsional dengan penyedia yang dapat dikonfigurasi).
- **App Marketplace pihak ketiga** untuk editor (macro tambahan, diagram, checklist lanjutan, template description berbayar, smart link resolver pihak ketiga khusus) — tidak dibangun sebagai fitur native; ekstensibilitas diatur di [25 Integrasi, Impor/Ekspor & Ekstensi](25-integrasi-impor-ekspor-ekstensi.md).
- **Aplikasi mobile native iOS/Android** — di luar cakupan PRD ini; web responsif menampilkan subset elemen (EDT-003).
- **Agregasi 'My tasks' lintas work item dan decision log** seperti di produk wiki — tidak dibangun (EDT-016).
