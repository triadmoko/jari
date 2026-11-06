# 23 — REST API Publik & Webhook

Modul ini mendefinisikan permukaan integrasi publik produk ini: kontrak REST API (konvensi request, ADF, pagination, format error, versioning, rate limit, async task), seluruh metode autentikasi integrasi (API token, OAuth 2.0 3LO, scope classic/granular, service account, kebijakan admin, model app ala Forge, application links), katalog resource REST yang dipakai SPA maupun pihak ketiga (work item, komentar, lampiran, worklog, link, properti, bulk, arsip, redaksi, pencarian JQL, Agile, proyek, konfigurasi skema, workflow, pengguna, grup, izin, filter, dashboard, konfigurasi global, audit, klasifikasi, Plans, Teams), API data DevOps, webhook (admin sistem, katalog event, payload, pengiriman andal, dynamic webhook untuk app, web request otomasi), serta pengalaman developer (dokumentasi OpenAPI, kebijakan deprecation, developer console, CLI/GraphQL/MCP). Modul ini memiliki kontrak API dan mekanisme webhook; logika domain di balik setiap resource tetap dimiliki modul masing-masing dan hanya dirujuk. Kondisi repo saat ini: semua route berada di bawah `/api/v1` dengan format respons seragam `{error, status, message, data, errors}` dari package `internal/shared/delivery/http/response`, autentikasi JWT lewat `middleware.AuthMiddleware()` (claims di context `sess`), bahasa dari `Accept-Language` lewat `middleware.LanguageMiddleware()`, dan dokumentasi OpenAPI dihasilkan `swag` lalu ditampilkan Scalar di `/swagger/index.html`. Pekerjaan modul ini menambahkan: lapisan kompatibel bergaya Jira (`/rest/api/3`, `/rest/agile/1.0`) di atas use case yang sama, middleware autentikasi API token/OAuth, middleware rate limit, feature module `internal/features/webhook` (entitas `Webhook`, `WebhookDelivery`), `internal/features/apitoken`, `internal/features/oauthapp`, serta `internal/features/task` untuk async task.

Konvensi penulisan modul ini: path endpoint ditulis relatif terhadap base `/rest/api/3` (lapisan kompatibel) kecuali disebut lain; path Agile relatif terhadap `/rest/agile/1.0`; padanan native ada di `/api/v1/...` dengan nama resource yang sama.

Modul terkait:

- [00 Visi & Ruang Lingkup](00-visi-ruang-lingkup-glosarium-baseline.md) — penamaan produk (VISI-001), istilah Work item (VISI-002), plan tier & feature flag (VISI-003).
- [01 Identitas, Autentikasi & Sesi](01-identitas-autentikasi-sesi.md) — proteksi login gagal & CAPTCHA (AUTH-006), authentication policies (AUTH-011), API tokens (AUTH-012), connected apps (AUTH-013).
- [02 Pengguna, Grup & Organisasi](02-pengguna-grup-organisasi.md) — user search REST (USR-025), service account (USR-028), organization API key (USR-030), SCIM (USR-035), IP allowlist (USR-038), data security policies (USR-040), GDPR API (USR-044).
- [03 Izin, Peran & Keamanan](03-izin-peran-keamanan.md) — permission check REST (PERM-058 s.d. PERM-063), klasifikasi (PERM-064), data security policy (PERM-065), redaksi (PERM-066).
- [04 Proyek](04-proyek.md) — Project REST API & validasi (PRJ-064), project properties (PRJ-062), trash (PRJ-033).
- [05 Konfigurasi Work Item](05-konfigurasi-work-item-skema.md) — issue type, custom field, context, screen, field configuration.
- [06 Work Item Inti](06-work-item-inti.md) — create/edit, arsip (ISS-034), bulk change (ISS-035), batas work item (ISS-044).
- [07 Editor Konten Kaya](07-editor-konten-kaya.md) — ADF & rendering aman (EDT-001, EDT-002).
- [08 Komentar, Lampiran & Aktivitas](08-komentar-lampiran-aktivitas.md) — komentar & API surface (CMT-010), lampiran (CMT-011 s.d. CMT-016).
- [09 Workflow](09-workflow-status-transisi.md) — workflow REST API (WF-041), properti workflow (WF-031).
- [10 Board](10-board-kolom-kartu.md) — Agile REST API (BRD-052), event board & sprint (BRD-053).
- [11 Backlog & Sprint](11-backlog-sprint.md) — sprint, rank, backlog.
- [12 JQL](12-jql-bahasa-kueri.md) — grammar, evaluasi berizin (JQL-051), batas eksekusi (JQL-052), REST JQL (JQL-060).
- [13 Pencarian & Filter](13-pencarian-filter-ekspor.md) — filter tersimpan, sharing, subscription.
- [14 Pelacakan Waktu](14-pelacakan-waktu-estimasi.md) — Worklog REST API (TIME-040), worklog webhooks (TIME-042).
- [15 Rilis & Komponen](15-rilis-versi-komponen.md) — Version REST API (REL-027), Component REST API (REL-034).
- [17 Dashboard](17-dashboard-gadget.md) — Dashboard REST API (DSH-056), gadgets REST API (DSH-057).
- [18 Timeline & Plans](18-timeline-rencana-tampilan-proyek.md) — Plans, Teams.
- [19 Notifikasi & Email](19-notifikasi-email.md) — event sistem (NTF-001), custom notification via REST (NTF-035), notification scheme REST (NTF-036).
- [20 Otomasi](20-otomasi.md) — incoming webhook (AUTO-025), send web request (AUTO-054), koneksi & secret (AUTO-056).
- [21 Navigasi & Preferensi](21-navigasi-ux-preferensi.md) — preferensi user yang diekspos lewat API.
- [22 Administrasi Sistem](22-administrasi-sistem.md) — halaman admin webhook, audit log (ADM-014), advanced settings (ADM-009), banner (ADM-007).
- [24 Integrasi Pengembangan & DevOps](24-integrasi-devops.md) — domain data development, build, deployment (sisi produk API-061 s.d. API-063).
- [25 Integrasi Pihak Ketiga, Impor/Ekspor & Ekstensibilitas](25-integrasi-impor-ekspor-ekstensi.md) — model ekstensi/runtime app, application links, Marketplace sebagai ekosistem.
- [27 NFR & Arsitektur Platform](27-nfr-arsitektur-platform.md) — job queue, storage, limits & guardrails, outbox event.
- [28 Model Data Inti & Katalog Event Domain](28-model-data-inti.md) — skema data dan katalog event yang menjadi sumber webhook.
- [29 Service Management: Service Project, Request Type, Forms & Agent](29-jsm-service-project-request-type-formulir.md) — komentar internal/publik di service project.
- [32 Service Management: ITSM](32-jsm-itsm-insiden-problem-perubahan-layanan.md) — change management & deployment gating.
- [35 Service Management: AI, Otomasi, API, Admin & CSM](35-jsm-ai-otomasi-api-admin-csm.md) — REST API Service Management, scope dan event JSM.
- [37 Product Discovery: Insights, Delivery, Peran, Integrasi & API](37-jpd-insights-delivery-peran-integrasi.md) — API dan event Product Discovery.

## Persona terkait

- **Integrator/Developer API** — persona utama: membangun skrip, integrasi server-side, dan app di atas REST API, OAuth, dan webhook.
- **Admin Situs** — mengelola webhook sistem, OAuth credentials DevOps, kebijakan API token dan akses app, application links, audit, serta membaca konfigurasi global lewat API.
- **Admin Proyek/Project Lead** — mengotomasi konfigurasi proyek (komponen, versi, role, properti) dan memakai web request otomasi ber-scope proyek.
- **Scrum Master/Team Lead** — mengotomasi sprint, backlog, dan rank lewat Agile API.
- **Developer/Anggota Tim** — membuat API token pribadi, memakai CLI, dan menerima data development dari tool SCM di work item.
- **Product Owner/PM** — mengekspor dan menyinkronkan work item serta Plans ke tool lain lewat API.
- **Stakeholder/Viewer** — membaca data lewat integrasi read-only dengan permission terbatas.
- **Tamu/Kolaborator Eksternal** — akses anonim ke REST API bila proyek publik.
- **Agent Layanan** — terdampak visibilitas komentar internal/publik saat integrasi menulis komentar ke service project.

## User story modul

- Sebagai Integrator/Developer API, saya ingin kontrak REST yang konsisten (expand, fields, accountId, ISO 8601) dengan format error per field agar klien saya dapat dibangun sekali dan dipakai di semua resource (API-001, API-005).
- Sebagai Integrator/Developer API, saya ingin pola pagination yang terdokumentasi dan penanda halaman terakhir yang andal agar sinkronisasi data besar tidak kehilangan item (API-004).
- Sebagai Integrator/Developer API, saya ingin header rate limit dan `Retry-After` yang jelas agar klien saya bisa melakukan backoff tanpa diblokir (API-006).
- Sebagai Developer/Anggota Tim, saya ingin membuat API token berlabel dengan masa berlaku dan mencabutnya kapan saja agar skrip pribadi saya aman (API-008).
- Sebagai Integrator/Developer API, saya ingin alur OAuth 2.0 dengan consent dan scope granular agar app saya hanya mendapat akses minimum yang disetujui pengguna (API-009, API-010).
- Sebagai Admin Situs, saya ingin memblokir pembuatan API token dan app buatan pengguna serta membatasi akses lewat IP allowlist agar data organisasi tidak bocor lewat integrasi liar (API-012).
- Sebagai Integrator/Developer API, saya ingin membuat, mengubah, mentransisikan, dan mengomentari work item dengan validasi layar yang sama seperti UI agar integrasi tidak menghasilkan data yang melanggar konfigurasi (API-016, API-017, API-018, API-020).
- Sebagai Integrator/Developer API, saya ingin mencari work item dengan JQL bertoken halaman dan parameter `reconcileIssues` agar hasil langsung konsisten setelah saya menulis data (API-029).
- Sebagai Scrum Master/Team Lead, saya ingin memulai dan menutup sprint serta memindahkan work item ke sprint lewat API agar ritual sprint bisa diotomasi dari tool tim (API-034, API-033).
- Sebagai Admin Proyek/Project Lead, saya ingin membuat versi, komponen, dan mengatur role proyek lewat API agar penyiapan proyek baru dapat diskripkan (API-038, API-039, API-040).
- Sebagai Admin Situs, saya ingin mendaftarkan webhook dengan filter JQL dan secret HMAC agar sistem eksternal hanya menerima event yang relevan dan bisa memverifikasi pengirimnya (API-064, API-066).
- Sebagai Integrator/Developer API, saya ingin header identitas pengiriman dan retry otomatis pada webhook agar penerima saya dapat idempoten dan tidak kehilangan event (API-066).
- Sebagai Developer/Anggota Tim, saya ingin commit, branch, dan PR yang menyebut key work item muncul di panel Development agar status kode terlihat tanpa membuka tool SCM (API-062).
- Sebagai Integrator/Developer API, saya ingin spesifikasi OpenAPI yang bisa diunduh beserta changelog dan masa deprecation minimal enam bulan agar client yang saya generate tidak rusak mendadak (API-069, API-070, API-002).
- Sebagai Admin Situs, saya ingin membaca audit record dan status async task lewat API agar pemantauan konfigurasi dan operasi berat bisa dimasukkan ke tool observabilitas (API-057, API-007).

## Ringkasan prioritas & fase

| Prioritas | Jumlah fitur |
|---|---|
| P0 | 16 |
| P1 | 27 |
| P2 | 23 |
| P3 | 8 |
| **Total** | **74** |

| Fase | Jumlah fitur |
|---|---|
| Fase 1 | 14 |
| Fase 2 | 17 |
| Fase 3 | 22 |
| Fase 4 | 19 |
| Fase — (dc-only/deprecated, rujukan) | 2 |

Catatan: jatah inventaris berjumlah 76 fitur; dua pasang fitur tumpang tindih digabung menjadi satu heading (API-006 dan API-066) dengan baris alias.

Strategi urutan implementasi: Fase 1 menetapkan kontrak dasar (konvensi request, ADF, pagination, format error) bersamaan dengan resource yang dipakai SPA sendiri, yaitu work item, createmeta/editmeta, transisi, komentar, lampiran, proyek, pengguna, serta board, backlog, dan sprint, sehingga SPA menjadi klien pertama API publik dan kontraknya teruji sejak awal. Fase 2 membuka API untuk pihak ketiga lewat API token, rate limit, kebijakan versioning dengan pemeriksaan breaking change di CI, pencarian JQL, serta resource konfigurasi yang muncul bersama JQL, field kustom, versi, komponen, dan time tracking. Fase 3 menambah webhook (admin, katalog event, pengiriman andal), bulk API, async task, entity properties, resource skema dan workflow, audit, serta dokumentasi publik dan kebijakan deprecation; webhook delivery (API-066) sebaiknya dibangun sebelum katalog event lengkap karena semua event memakai antrean yang sama. Fase 4 menghadirkan ekosistem app: OAuth 2.0 3LO dan scope, service account, kebijakan akses app, model app ala Forge, dynamic webhook, developer console, API DevOps, Plans, Teams, arsip, redaksi, dan klasifikasi. Fitur Data Center-only (API-015, API-074) hanya menjadi rujukan kompatibilitas.

## Fitur

## Fondasi kontrak API

### API-001 REST API Platform & Request Conventions

- **Deskripsi:** Kontrak dasar REST API publik produk ini: versi, base URL, format body, dan konvensi parameter yang berlaku di semua resource (padanan konvensi Jira Cloud REST v2/v3). Kontrak ini menjadi fondasi agar integrasi pihak ketiga dan SPA memakai bentuk API yang sama.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Integrator/Developer API, Admin Situs, app pihak ketiga; setiap panggilan tunduk pada global permission, project permission, dan issue security pemanggil.
- **Aturan perilaku (acceptance criteria):**
  1. Base URL per site melayani `/rest/api/3/...` (rich text ADF) dan `/rest/api/2/...` (rich text wiki markup/string) dengan resource dan perilaku identik kecuali format rich text; native `/api/v1/...` memakai use case yang sama.
  2. Panggilan dengan OAuth 2.0, scoped token, atau service account diterima melalui path gateway `/ex/{siteId}/rest/api/3/...`; pada instalasi single-site path gateway adalah alias base URL langsung, dan `siteId` yang tidak dikenal menghasilkan 404.
  3. Di v3, field `description`, `environment`, body komentar, komentar worklog, dan custom field textarea (multi-line) diterima dan dikembalikan sebagai ADF JSON; di v2 field yang sama berupa string wiki markup; mengirim string ke v3 untuk field ADF ditolak 400.
  4. Parameter `expand` menerima daftar dipisah koma: `renderedFields`, `names`, `schema`, `transitions`, `operations`, `editmeta`, `changelog`, `versionedRepresentations`, beserta sub-expand seperti `transitions.fields`; nilai expand yang tidak dikenal diabaikan tanpa error.
  5. Parameter `fields` menerima daftar id/key field, `*all`, `*navigable`, dan prefix minus untuk mengecualikan (contoh `*all,-comment`); `fieldsByKeys=true` membuat daftar dan respons memakai key field, bukan id.
  6. Parameter `properties` menyertakan entity properties pada respons (`*all`, key tertentu, atau prefix minus untuk mengecualikan).
  7. Semua path yang menerima `{issueIdOrKey}`/`{projectIdOrKey}` menerima id numerik maupun key; setiap objek di respons menyertakan link `self` absolut.
  8. Body tulis mendukung `fields` (set langsung) dan `update` dengan operasi `set`/`add`/`remove`/`edit`/`copy` per field; operasi yang tidak didukung field tersebut (lihat `editmeta`) ditolak 400 dengan pesan di `errors[fieldId]`.
  9. Datetime memakai ISO 8601 dengan offset zona waktu (contoh `2025-01-01T10:00:00.000+0700`); field bertipe date memakai `yyyy-MM-dd`; format lain ditolak 400.
  10. Pengguna hanya diidentifikasi dengan `accountId`; parameter `username`/`userKey` tidak didukung dan diabaikan atau ditolak 400 dengan pesan deprecation.
  11. Request dan respons memakai `Content-Type`/`Accept: application/json`; endpoint multipart mewajibkan header `X-Atlassian-Token: no-check` (proteksi XSRF), tanpanya request ditolak 403.
  12. Header `Accept-Language` menentukan bahasa teks terlokalisasi (nama field, status, pesan error) sesuai `LanguageMiddleware`; `X-Force-Accept-Language: true` memaksa bahasa header walau preferensi profil berbeda.
  13. ⚠ Endpoint berlabel Experimental tidak memiliki jaminan stabilitas; sebagian endpoint eksperimental mewajibkan header opt-in `X-ExperimentalApi: opt-in` dan tanpa header menolak 400 dengan pesan yang menjelaskan opt-in.
  14. Setiap respons menyertakan header debugging: id request unik (`X-AREQUESTID`, sama dengan id di log server), trace id (`Atl-Traceid`/`traceparent`), dan `X-AACCOUNTID` pemanggil terautentikasi.
  15. Hasil pencarian bersifat eventually consistent terhadap indeks dan bisa tertinggal setelah write; enhanced search menyediakan `reconcileIssues` untuk konsistensi read-after-write (lihat API-029).
  16. Setiap endpoint menegakkan project permission, issue security, dan global permission pemanggil; item yang tidak boleh dilihat disaring diam-diam dari daftar, sedangkan akses langsung dijawab 403 atau 404 sesuai API-005.
  17. Bila Browse Projects diberikan ke "Anyone on the web", request tanpa autentikasi dapat membaca resource proyek tersebut; selain itu request anonim mendapat 401 atau hasil kosong.
  18. CORS tidak diaktifkan untuk panggilan langsung dari origin browser pihak ketiga dengan API token; preflight dari origin asing ditolak, dan integrasi harus berjalan server-side atau lewat model app (API-013).
  19. Penamaan UI "work item" tidak mengubah kontrak API: path tetap `/issue`, nama field tetap `issuetype`, dan key respons tetap `issues`.
- **Team-managed vs company-managed:** Objek project menyertakan `simplified` (true untuk team-managed) dan `style` (`classic`/`next-gen`); resource lain memakai kontrak yang sama, perbedaan perilaku dijelaskan per fitur.
- **Catatan & ketergantungan:** Rekomendasi inventaris: bentuk v3 (ADF) sebagai kanonik, v2 sebagai alias opsional. Lapisan kompatibel dibangun sebagai router group terpisah yang memanggil use case yang sama dengan `/api/v1`; keputusan ruang lingkup lapisan kompatibel ada di API-002. Penegakan izin bergantung pada modul 03; ADF bergantung pada modul 07.

### API-002 API Versioning, Compatibility & OpenAPI

- **Deskripsi:** Kebijakan versi REST API produk ini, lapisan kompatibel bentuk API Jira, dan kontrak OpenAPI yang dihasilkan otomatis serta dijaga di CI.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / platform
- **Aktor & izin:** Engineering, Integrator/Developer API; tidak ada permission pengguna khusus.
- **Aturan perilaku (acceptance criteria):**
  1. Semua route native berada di `/api/v1` sesuai konvensi repo; lapisan kompatibel `/rest/api/3` dan `/rest/agile/1.0` disediakan untuk integrasi yang mengharapkan bentuk API Jira, dengan daftar endpoint kompatibel terdokumentasi (endpoint di luar daftar mengembalikan 404).
  2. Kontrak OpenAPI dihasilkan `swag` (`make swag`) dan CI membandingkan spesifikasi terhadap cabang utama; perubahan breaking (hapus field/endpoint, ubah tipe, field wajib baru) menggagalkan build kecuali diberi label persetujuan.
  3. Endpoint yang dideprecate mengirim header `Deprecation` dan `Sunset` (tanggal), dengan masa minimal 6 bulan sebelum dihapus.
  4. Endpoint native memakai format respons seragam package `response` dengan kode error stabil; lapisan kompatibel memakai bentuk `{errorMessages, errors}` (API-005).
  5. Pagination konsisten: `startAt`/`maxResults`/`total`/`isLast` atau `nextPageToken` sesuai pola API-004.
  6. Parameter `expand` (`renderedFields`, `names`, `schema`, `transitions`, `editmeta`, `changelog`) dan seleksi `fields` didukung di endpoint native dan kompatibel dengan semantik yang sama.
  7. POST create menerima header `Idempotency-Key`; request ulang dengan key dan body yang sama dalam 24 jam mengembalikan respons pertama tanpa membuat objek ganda, sedangkan key sama dengan body berbeda ditolak 409.
  8. Dokumentasi Scalar di `/swagger/index.html` memuat contoh request dan skema autentikasi (`BearerAuth`, Basic API token, OAuth 2.0) untuk setiap operasi.
  9. SDK klien TypeScript dihasilkan dari OpenAPI dan dipakai SPA di `frontend/`; build frontend gagal bila SDK tidak sinkron dengan spesifikasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Fitur berasal dari domain NFR; detail endpoint ada di fitur API lain modul ini. Berkaitan dengan API-069 (dokumentasi) dan API-070 (deprecation). Keputusan cakupan lapisan kompatibel dicatat di Pertanyaan terbuka.

### API-003 Atlassian Document Format (ADF) in API

- **Deskripsi:** Format JSON rich text (ADF) yang dipakai REST v3 untuk description, komentar, dan field teks panjang; server memvalidasi dan me-render struktur node ini.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Integrator/Developer API, pengembang frontend; izin mengikuti field yang ditulis.
- **Aturan perilaku (acceptance criteria):**
  1. Dokumen root wajib berbentuk `{"version": 1, "type": "doc", "content": [...]}`; versi atau type lain ditolak 400.
  2. ⚠ Block node yang didukung: `paragraph`, `heading` (level 1-6), `bulletList`, `orderedList`, `listItem`, `codeBlock` (attr `language`), `blockquote`, `panel` (info/note/warning/success/error), `rule`, `table`/`tableRow`/`tableHeader`/`tableCell`, `mediaSingle`/`mediaGroup`, `expand`, `nestedExpand`, `taskList`/`taskItem`, dan `decisionList`.
  3. ⚠ Inline node yang didukung: `text`, `hardBreak`, `mention` (attr `id` = accountId), `emoji`, `inlineCard` (smart link), `date`, `status`, dan media inline.
  4. ⚠ Mark yang didukung: `strong`, `em`, `underline`, `strike`, `code`, `link` (`href`), `textColor`, `subsup`, dan `backgroundColor`.
  5. Node `mention` memicu notifikasi ke pengguna yang disebut hanya bila pengguna tersebut memiliki Browse pada work item; mention ke pengguna tanpa akses tetap tersimpan tanpa notifikasi.
  6. Node media mereferensikan id file lampiran work item; id yang bukan milik work item yang sama ditolak 400.
  7. ADF yang tidak valid (node tak dikenal, struktur salah, attr wajib hilang) ditolak 400 dengan pesan di `errors[fieldId]`.
  8. `expand=renderedFields` mengembalikan HTML hasil render yang sudah disanitasi untuk tampilan read-only.
  9. Konversi otomatis wiki markup (v2) ⇄ ADF (v3) terjadi saat membaca/menulis lewat versi API berbeda; node tanpa padanan dikonversi lossy (misalnya menjadi teks biasa) dan perilaku ini terdokumentasi.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tumpang tindih dengan modul 07 (EDT-001 editor, EDT-002 rendering aman); di sini dicatat dari sisi kontrak API. Penyimpanan internal disarankan ADF agar v3 tanpa konversi.

### API-004 Pagination Models

- **Deskripsi:** Pola pagination yang didukung API (offset-based, token-based, berbasis waktu) dengan batas `maxResults` yang konsisten dan terdokumentasi per endpoint.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Endpoint list baru memakai PageBean offset-based: `startAt`, `maxResults`, `total`, `isLast`, `values`, ditambah link `self` dan `nextPage` (tidak ada `nextPage` di halaman terakhir).
  2. Pola lama tanpa `isLast` (`startAt`/`maxResults`/`total` dengan nama koleksi seperti `comments`/`worklogs`) dipertahankan untuk list komentar dan worklog.
  3. Pola token-based (`nextPageToken`) dipakai `/search/jql`, `/changelog/bulkfetch`, dan evaluasi expression multi-issue; respons tidak memuat `total`, dan tidak adanya `nextPageToken` menandakan halaman terakhir.
  4. Pola berbasis waktu dipakai `/worklog/updated` dan `/worklog/deleted` dengan `since`/`until` dan `lastPage`.
  5. Agile API memakai `startAt`/`maxResults`/`isLast`/`total` untuk boards, sprints, dan issues.
  6. ⚠ Setiap endpoint mendokumentasikan default dan maksimum `maxResults` (umumnya default 50, maksimum 50-100; user search s.d. 1000; `search/jql` lebih besar bila hanya id/key diminta); nilai di atas maksimum dipangkas ke maksimum tanpa error.
  7. Server boleh mengembalikan item lebih sedikit dari `maxResults` (misalnya setelah penyaringan izin); klien wajib berpatokan pada `isLast`/`nextPageToken`, dan dokumentasi menyatakannya.
  8. Enhanced search tidak menghitung total; jumlah perkiraan tersedia di `/search/approximate-count`.
  9. Parameter `orderBy` (contoh `+name`, `-created`, `sequence`) didukung per endpoint sesuai daftar yang terdokumentasi; nilai tidak dikenal ditolak 400.
  10. Endpoint bulk-by-id (POST `.../list`, `bulkfetch`) tersedia sebagai alternatif paging bila id item sudah diketahui.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Selaras dengan `pkg/pagination.go` di repo; utilitas ini diperluas untuk PageBean dan token (token opaque, ditandatangani, berumur terbatas).

### API-005 Error Format & HTTP Status Codes

- **Deskripsi:** Format error standar dan pemetaan kode status HTTP agar klien dapat menampilkan error validasi per field secara konsisten.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Body error lapisan kompatibel berbentuk `{"errorMessages": [string], "errors": {"fieldId": "pesan"}}` dengan `warningMessages` opsional pada endpoint tertentu; endpoint native membungkus struktur yang sama di `errors` format `response`.
  2. Error validasi per field memakai id field sebagai key (contoh `summary`, `customfield_10010`) dan pesan dilokalisasi sesuai `Accept-Language`.
  3. 200/201 untuk sukses; 204 tanpa body untuk transisi, assign, dan delete.
  4. 400 untuk input atau JQL tidak valid; 401 bila kredensial hilang atau tidak valid; 403 bila permission kurang, fitur dimatikan, atau scope token tidak mencukupi.
  5. 404 dikembalikan baik untuk resource yang tidak ada maupun yang tidak boleh diakses pemanggil, tanpa membedakan keduanya.
  6. 409 untuk konflik (contoh versi workflow usang), 413 untuk payload terlalu besar, 415 untuk media type salah.
  7. 429 saat terkena rate limit, selalu disertai `Retry-After` (API-006).
  8. 500 untuk error server tak terduga (tanpa stack trace di body, dengan id request) dan 503 untuk layanan sementara tidak tersedia.
  9. 207 Multi-Status untuk operasi parsial seperti rank (status per item); bulk create mengembalikan 201 dengan array `errors` per `elementNumber`.
  10. Operasi async mengembalikan 303 See Other dengan `Location` ke `/task/{id}` (API-007).
  11. Saat login/basic auth gagal atau CAPTCHA terpicu, respons menyertakan header `X-Seraph-LoginReason` (`AUTHENTICATED_FAILED`, `AUTHENTICATION_DENIED`).
  12. Agile API memakai format `errorMessages`/`errors` yang sama.
  13. ⚠ API generasi baru (workflows, bulk operations, plans, redact) boleh memakai format `{code, message, status}`; skemanya didokumentasikan per endpoint di OpenAPI.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pesan error memakai `internal/shared/constants/error.go` (EN/ID). Lapisan kompatibel memetakan error domain ke bentuk Jira lewat satu adaptor.

### API-006 Rate Limiting & Service Limits

- **Alias inventaris:** Rate Limiting & Abuse Protection
- **Deskripsi:** Pembatasan laju dan kuota API per pengguna/token/app/IP, proteksi penyalahgunaan (brute-force), serta batas data per work item dan per request, dengan header kompatibel agar klien dapat melakukan backoff.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Integrator/Developer API, operator platform/SRE, Admin Situs (melihat dan mengatur batas yang dapat dikonfigurasi).
- **Aturan perilaku (acceptance criteria):**
  1. Saat batas terlampaui server mengembalikan 429 dengan header `Retry-After` (detik) dan `Beta-Retry-After`.
  2. Setiap respons ber-rate-limit menyertakan `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`, dan `X-RateLimit-NearLimit: true` saat sisa budget di bawah 20%.
  3. Respons 429 menyertakan `RateLimit-Reason` bernilai salah satu dari `jira-burst-based`, `jira-quota-global-based`, `jira-quota-tenant-based`, `jira-per-issue-on-write`.
  4. ⚠ Model berbasis biaya/poin per app per tenant per jam: endpoint mahal (search, bulk, expand changelog) mengonsumsi lebih banyak poin, dengan tabel biaya per endpoint terdokumentasi.
  5. Burst limit per detik dan quota per jam diterapkan terpisah; batas untuk app (OAuth/model app) dikonfigurasi berbeda dari batas API token pengguna.
  6. ⚠ Per-issue write limit membatasi jumlah write ke satu work item dalam jendela waktu pendek; pelanggaran dijawab 429 dengan reason `jira-per-issue-on-write`.
  7. Dokumentasi merekomendasikan exponential backoff dengan jitter, kepatuhan pada `Retry-After`, dan pengurangan konkurensi; SDK resmi mengimplementasikannya.
  8. ⚠ Batas per work item: komentar (~5.000), lampiran (~2.000), worklog (~10.000), issue link (~2.000), remote link (~2.000), watcher, dan ukuran changelog; write yang melampaui ditolak 400/413 dengan pesan batas.
  9. `GET /issue/limit/report` (eksperimental) mengembalikan work item yang mendekati atau melampaui batas, dengan `isReturningKeys` dan threshold per field.
  10. Entity property: value maksimum 32 KB dan key maksimum 255 karakter.
  11. Batas ukuran lampiran per file mengikuti pengaturan admin dan terbaca di `/attachment/meta`.
  12. Batas item per bulk request: bulk create 50, bulkfetch 100, rank 50, move to sprint/backlog 50, Bulk Operations s.d. 1000 work item; kelebihan ditolak 400.
  13. JQL memiliki batas panjang/kompleksitas dan Jira expression memiliki batas kompleksitas; pelanggaran ditolak 400 dengan pesan yang menyebut batasnya.
  14. ⚠ Rate limit dihitung berbasis cost/points per pengguna, per token, dan per app, dengan burst limit (angka default dikonfigurasi lewat environment).
  15. Header 429 dan `X-RateLimit-*` berlaku juga untuk route native `/api/v1`.
  16. Batas khusus untuk login, reset password, signup, dan pembuatan API token (anti brute-force) aktif sejak MVP, per IP dan per akun.
  17. Setelah N login gagal berturut-turut diterapkan lockout atau progressive delay; CAPTCHA opsional dapat diaktifkan admin.
  18. Body request JSON dibatasi (default 10 MB, dapat dikonfigurasi); kelebihan ditolak 413 sebelum diparsing.
  19. Concurrency limit per pengguna untuk operasi berat (export, bulk, reindex); permintaan di atas batas ditolak 429 atau diantrekan.
  20. IP allowlisting tersedia sebagai fitur opsional per plan (pembanding: Premium/Enterprise), lihat API-012.
  21. Eksekusi otomasi di-throttle sesuai plan dan loop rule terdeteksi lalu dihentikan (lihat modul 20).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Digabung dari domain api-webhooks (P1, free-standard) dan nfr-architecture (P1, platform). Pengecualian fase: butir 16-17 (batas khusus login/reset password/signup/pembuatan token dan lockout) wajib aktif sejak Fase 1 bersama AUTH-006. Implementasi sebagai middleware Gin (token bucket per kunci, penyimpanan Redis/Postgres). Proteksi login dirinci di AUTH-006; batas item bulk dipusatkan di [modul 27](27-nfr-arsitektur-platform.md) (Limits & Guardrails) dan ISS-044; throttling otomasi di modul 20.

### API-007 Async Tasks API

- **Deskripsi:** Mekanisme long-running task untuk operasi berat (hapus field, bulk properties, arsip, publish workflow scheme, hapus proyek) yang statusnya dipantau lewat polling.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Integrator/Developer API, Admin Situs; task hanya dapat dibaca/dibatalkan oleh pengirimnya atau Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Operasi async mengembalikan 303 (atau 202) dengan header `Location: /rest/api/3/task/{taskId}`.
  2. `GET /task/{taskId}` mengembalikan `status` (`ENQUEUED`, `RUNNING`, `COMPLETE`, `FAILED`, `CANCEL_REQUESTED`, `CANCELLED`, `DEAD`), `progress` (0-100), `message`, `result`, `submittedBy`, `submitted`, `started`, `finished`, `lastUpdate`.
  3. `POST /task/{taskId}/cancel` mengubah status menjadi `CANCEL_REQUESTED` lalu `CANCELLED` bila operasi mendukung pembatalan; task yang sudah selesai menolak cancel 400.
  4. ⚠ Riwayat task disimpan sementara (default 14 hari, dapat dikonfigurasi); task kedaluwarsa menghasilkan 404.
  5. Bulk Operations memakai antrean terpisah `/bulk/queue/{taskId}` (API-026) dan redaksi memakai `/redact/status/{jobId}` (API-028), dengan semantik status yang setara.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Bergantung pada job queue di [modul 27](27-nfr-arsitektur-platform.md). Task worker berjalan di proses yang sama dengan graceful shutdown; task `RUNNING` saat shutdown dilanjutkan atau ditandai `DEAD`.

## Autentikasi & otorisasi integrasi

### API-008 API Tokens & Basic Authentication

- **Deskripsi:** Autentikasi skrip dan integrasi sederhana dengan email + API token lewat HTTP Basic auth, beserta siklus hidup token dan scoped API token.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / platform
- **Aktor & izin:** Developer/Anggota Tim dan Integrator/Developer API (membuat token sendiri), Admin Situs/org admin (melihat, mencabut, dan memblokir token akun terkelola).
- **Aturan perilaku (acceptance criteria):**
  1. Header `Authorization: Basic base64(email:api_token)` diterima di semua endpoint REST; Basic auth dengan password akun dan cookie session lewat `/rest/auth/1/session` ditolak 401.
  2. Token dibuat dari halaman keamanan profil (Personal settings > Security > API tokens) dengan label wajib (1-100 karakter).
  3. Masa berlaku token wajib 1-365 hari; token tanpa expiry tidak dapat dibuat, dan token hasil migrasi tanpa expiry diberi tanggal kedaluwarsa default.
  4. Scoped API token memilih produk dan scope granular saat pembuatan dan hanya dapat dipakai lewat path gateway `/ex/{siteId}/...`; panggilan di luar scope ditolak 403.
  5. Nilai token hanya ditampilkan sekali saat dibuat; server hanya menyimpan hash, dan tidak ada endpoint yang mengembalikan nilai token lagi.
  6. Pengguna dapat mencabut token satu per satu atau semua sekaligus; daftar token menampilkan label, tanggal dibuat, tanggal kedaluwarsa, dan terakhir dipakai.
  7. Token unscoped mewarisi seluruh permission pengguna tanpa pembatasan per proyek.
  8. Org admin dapat melihat dan mencabut token akun terkelola serta memblokir pembuatan API token lewat policy; pembuatan yang diblokir ditolak 403 dengan pesan policy.
  9. Login gagal berulang dapat memicu CAPTCHA yang memblokir Basic auth (respons 401 dengan `X-Seraph-LoginReason: AUTHENTICATION_DENIED`) sampai pengguna login lewat browser.
  10. Akun SSO/terkelola tetap dapat memakai API token kecuali diblokir policy.
  11. Request tanpa header `Authorization` diperlakukan anonim dan hanya mendapat permission yang diberikan ke "Anyone".
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Untuk MVP clone: personal API token per pengguna dengan expiry, label, revoke, dan last-used; scoped token menyusul bersama API-010. Personal Access Token (Bearer) hanya ada di Data Center (API-074). UI token dirinci di AUTH-012; brute-force di AUTH-006. Migrasi repo: `AuthMiddleware` saat ini hanya menerima JWT `Bearer`; perlu cabang Basic yang memvalidasi hash token di tabel baru `api_tokens` lalu menyusun `*jwt.JWTClaims` yang sama ke context `sess`.

### API-009 OAuth 2.0 (3LO) Authorization Code Flow

- **Deskripsi:** Alur OAuth 2.0 three-legged agar aplikasi eksternal bertindak atas nama pengguna dengan consent eksplisit, termasuk token exchange, refresh, dan penemuan site.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pemilik app), pengguna yang memberi consent, Admin Situs/org admin (mengelola app terhubung).
- **Aturan perilaku (acceptance criteria):**
  1. Endpoint otorisasi `/oauth/authorize` menerima `audience`, `client_id`, `scope`, `redirect_uri`, `state`, `response_type=code`, dan `prompt=consent`; `redirect_uri` yang tidak terdaftar ditolak tanpa redirect.
  2. Layar consent menampilkan nama app, daftar scope yang diminta dalam bahasa manusia, dan pilihan site yang diberi akses.
  3. `POST /oauth/token` dengan `grant_type=authorization_code`, `client_id`, `client_secret`, `code`, `redirect_uri` menukar code sekali pakai (berlaku maksimal 10 menit) menjadi token.
  4. Access token Bearer berumur pendek (sekitar 1 jam) dan ditolak 401 setelah kedaluwarsa.
  5. ⚠ Refresh token hanya diterbitkan bila scope `offline_access` diminta; refresh token bersifat rotating (token lama invalid setelah dipakai), kedaluwarsa setelah sekitar 90 hari tidak aktif, dan memiliki umur absolut terbatas.
  6. `GET /oauth/token/accessible-resources` mengembalikan `id` (siteId), `name`, `url`, `scopes`, dan `avatarUrl` per site yang diberi akses.
  7. `GET /me` (scope `read:me`) mengembalikan profil pengguna yang mengotorisasi.
  8. Semua panggilan API dengan token OAuth wajib lewat `/ex/{siteId}/...`; panggilan ke base URL langsung dengan token OAuth ditolak 401.
  9. Parameter `state` dikembalikan apa adanya untuk mitigasi CSRF, dan app dapat mendaftarkan beberapa callback URL di developer console (API-071).
  10. Pengguna dapat mencabut akses app dari profil (Connected apps); admin dapat melihat dan mengelola app pihak ketiga yang terhubung, dan pencabutan langsung membatalkan token aktif.
  11. Permission efektif adalah irisan scope app dan permission pengguna di site tersebut.
  12. Admin dapat membatasi atau memblokir koneksi app buatan pengguna (user-installed apps/app access policy); upaya consent yang diblokir menampilkan pesan policy.
  13. Distribusi app dapat private (hanya developer pemilik) atau sharing aktif yang mewajibkan vendor info, privacy policy, dan personal data declaration.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Connected apps UI ada di AUTH-013. Implementasi otorisasi server sebagai feature `internal/features/oauthapp`. Scope dijelaskan di API-010.

### API-010 OAuth 2.0 Scopes (Classic & Granular)

- **Deskripsi:** Katalog scope yang membatasi akses app OAuth, model app, dan scoped token ke resource, terdiri dari scope classic yang luas dan scope granular per resource dan aksi.
- **Prioritas / Fase / Tier:** P1 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, pengguna yang memberi consent.
- **Aturan perilaku (acceptance criteria):**
  1. Scope classic: `read:jira-work`, `write:jira-work`, `read:jira-user`, `manage:jira-project`, `manage:jira-configuration`, `manage:jira-webhook`, `manage:jira-data-provider`.
  2. Scope platform: `read:me`, `offline_access`, `read:account`, `report:personal-data`.
  3. Scope granular berformat `aksi:resource:jira`, contoh `read:issue:jira`, `write:issue:jira`, `delete:issue:jira`, `read:comment:jira`, `write:comment:jira`, `read:attachment:jira`, `read:field:jira`, `read:project:jira`, `read:user:jira`, `read:group:jira`, `read:workflow:jira`, `write:webhook:jira`.
  4. Scope Agile: `read:board-scope:jira-software`, `write:board-scope:jira-software`, `read:sprint:jira-software`, `write:sprint:jira-software`, `read:epic:jira-software`, `read:issue-details:jira-software`, plus classic read/write untuk agile.
  5. ⚠ Scope DevOps untuk data provider mencakup build, deployment, dev-info, feature-flag, dan remote-link.
  6. Scope Service Management terpisah: `read:servicedesk-request`, `write:servicedesk-request`, `manage:servicedesk-customer`.
  7. Dokumentasi setiap operasi mencantumkan set scope classic dan granular yang dibutuhkan (API-069).
  8. ⚠ Jumlah scope per app dibatasi (sekitar 50); pendaftaran melebihi batas ditolak di developer console.
  9. Satu endpoint dapat dipenuhi oleh set classic ATAU set granular; token yang mencampur keduanya dievaluasi per set dan dokumentasi menyarankan tidak mencampur.
  10. Perubahan scope app memerlukan re-consent; token lama tetap terbatas pada scope lama sampai consent baru diberikan.
  11. Scope adalah plafon, bukan pemberian hak: scope tidak pernah menambah hak di atas permission pengguna, dan panggilan tanpa scope cukup ditolak 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Nama scope dipertahankan agar app kompatibel; prefix produk dapat dialiaskan ke nama netral di endpoint native. Scope Service Management dirinci di [modul 35](35-jsm-ai-otomasi-api-admin-csm.md).

### API-011 Service Accounts & Organization Admin APIs

- **Deskripsi:** Akun non-manusia untuk otomasi dengan kredensial OAuth client credentials atau scoped API token, plus API tingkat organisasi untuk administrasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Org admin (membuat service account dan Admin API key), Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Org admin membuat service account di administrasi organisasi dan memberinya akses produk/grup seperti pengguna biasa.
  2. Kredensial service account berupa OAuth 2.0 client credentials (`client_id`/`client_secret`) dan/atau scoped API token dengan expiry.
  3. Grant `client_credentials` ke `/oauth/token` menghasilkan access token yang dipakai lewat `/ex/{siteId}/...`.
  4. ⚠ Service account tidak dapat login ke UI (upaya login ditolak) dan kebijakan lisensi/billing-nya terdokumentasi.
  5. Organization REST API (`/admin/v1/orgs/{orgId}`) dengan Admin API key menyediakan users, domains, events (audit log org), policies, dan directories.
  6. Admin API key memiliki masa berlaku dan dapat dicabut; key kedaluwarsa atau dicabut ditolak 401.
  7. User provisioning lewat SCIM 2.0 dari identity provider tersedia (lihat USR-035).
  8. User management API (`/users/{accountId}/manage`) mengelola profil dan lifecycle akun terkelola (deactivate/delete).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat tier dan syarat plan service account belum pasti. Tumpang tindih dengan modul 02 (USR-028, USR-030, USR-035); modul ini hanya memiliki kontrak autentikasinya.

### API-012 Admin Controls for API & App Access

- **Deskripsi:** Kebijakan admin organisasi yang membatasi akses API dan app ke data, seperti IP allowlist, kebijakan API token, dan data security policy.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / enterprise
- **Aktor & izin:** Org admin, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. IP allowlisting berlaku untuk browser dan API (plan Premium/Enterprise); akses API dari IP di luar allowlist ditolak 403.
  2. Kebijakan pembuatan/penggunaan API token untuk akun terkelola dapat diatur blokir atau izinkan.
  3. Data security policy "App access" memblokir app (OAuth/model app) mengakses konten dalam cakupan (proyek/klasifikasi) dan tercermin di respons app data policy (API-072).
  4. Kontrol user-installed apps memblokir atau mengizinkan pengguna memasang app OAuth 2.0/model app sendiri.
  5. Kontrol instalasi ekstensi menentukan siapa yang boleh memasang dan alur permintaan app.
  6. Authentication policies (SSO, 2FA) tidak mencabut API token yang ada kecuali policy menyatakannya.
  7. Respons API saat policy memblokir akses adalah 403 dengan pesan yang menyebut policy (tanpa membocorkan detail konfigurasi).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat tier per kontrol belum pasti (IP allowlist: Premium/Enterprise; data security policy dan API token policy: add-on keamanan). Rujukan: USR-038, USR-040, PERM-065, AUTH-011. Kontrol Marketplace hanya relevan bila model ekstensi [modul 25](25-integrasi-impor-ekspor-ekstensi.md) ada.

### API-013 Forge App Authentication & Event Integration

- **Deskripsi:** Model autentikasi app native (padanan Forge): app memanggil API sebagai app atau sebagai pengguna, dengan trigger event, web trigger, dan scheduled trigger sebagai pengganti webhook.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pengembang app), Admin Situs (menyetujui instalasi dan upgrade scope).
- **Aturan perilaku (acceptance criteria):**
  1. Helper `requestJira` pada SDK app memakai route template literal yang meng-encode setiap segmen path untuk mencegah path injection; path mentah tanpa template ditolak SDK.
  2. Mode `asApp()` memanggil API sebagai app user (`accountType: app`) dengan permission yang diberikan ke app.
  3. Mode `asUser()` memanggil atas nama pengguna yang sedang berinteraksi; pemakaian pertama memerlukan consent pengguna ("Allow access"), dan tanpa consent panggilan ditolak.
  4. Scope dideklarasikan di manifest; penambahan scope memicu major version yang baru aktif setelah admin menyetujui upgrade.
  5. Product event trigger tersedia untuk created/updated/deleted/assigned/commented issue, worklog, attachment, sprint, version, project, field, dan filter (contoh `avi:jira:created:issue`).
  6. Trigger mendukung `ignoreSelf` (abaikan event dari app itu sendiri) dan filter berbasis expression (API-031).
  7. Web trigger menyediakan URL HTTP masuk (static atau dynamic) yang memanggil fungsi app.
  8. Scheduled trigger tersedia dengan interval `fiveMinute`, `hour`, `day`, `week`.
  9. Mode remote: backend eksternal menerima invocation token bertanda tangan (JWT) beserta token app/pengguna untuk memanggil API.
  10. Egress ke domain eksternal hanya diizinkan untuk domain yang dideklarasikan di manifest.
  11. App user account dibuat otomatis saat instalasi dan dapat ditambahkan ke project role/grup.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Clone tidak wajib meniru Forge; fitur ini menjadi inspirasi model `asApp`/`asUser` dan event subscription internal. Runtime app (hosting fungsi) adalah keputusan [modul 25](25-integrasi-impor-ekspor-ekstensi.md). Platform Forge sendiri adalah layanan vendor dan di luar cakupan.

### API-014 Application Links (Incoming/Outgoing)

- **Deskripsi:** Koneksi produk ini ke produk lain (wiki, SCM on-prem, instance lain) lewat Application Links dengan autentikasi OAuth.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Admin Situs (Administer).
- **Aturan perilaku (acceptance criteria):**
  1. Halaman admin Products > Application links menyediakan buat, edit, dan hapus link.
  2. Autentikasi incoming/outgoing memakai OAuth 2.0 atau OAuth 1.0a (impersonation), termasuk ke instance self-hosted.
  3. Setiap link menampilkan status (connected, error) dan URL yang dikonfigurasi; status error menyertakan pesan penyebab.
  4. Application link dipakai oleh remote issue link (API-023) dan makro sematan di produk wiki.
  5. Link antar produk cloud di organisasi yang sama dikelola otomatis dan tidak dapat dihapus manual.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris belum memastikan apakah OAuth 1.0a masih dapat dipilih untuk link baru; usulan clone: OAuth 2.0 saja untuk link baru. Rujukan [modul 25](25-integrasi-impor-ekspor-ekstensi.md).

### API-015 Atlassian Connect JWT Authentication

- **Deskripsi:** Model app legacy (Connect) yang memakai JWT bertanda tangan shared secret dan lifecycle callback; menuju end-of-support sehingga hanya dicatat sebagai rujukan.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Integrator/Developer API (app legacy).
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan: lifecycle callback installed/uninstalled/enabled/disabled mengirim `clientKey` dan `sharedSecret` (atau asymmetric install JWT).
  2. Rujukan: JWT HS256 dengan claim `iss`, `iat`, `exp`, `qsh` (query string hash), dan `sub`.
  3. Rujukan: context JWT untuk modul iframe dan validasi request masuk.
  4. Rujukan: impersonation pengguna lewat OAuth 2.0 JWT bearer grant (scope `ACT_AS_USER`).
  5. Rujukan: static webhooks di app descriptor dengan event, url, filter JQL, `excludeBody`, `propertyKeys`.
  6. Rujukan: scope descriptor `READ`, `WRITE`, `DELETE`, `PROJECT_ADMIN`, `ADMIN`, `ACT_AS_USER`.
  7. Rujukan: akses email pengguna lewat `/user/email` hanya untuk app yang disetujui.
  8. ⚠ Rujukan: listing app baru dihentikan dan end-of-support direncanakan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Tidak dibangun karena model sudah deprecated; konsep yang berguna (lifecycle callback, shared secret per instalasi, `qsh`) dapat dipakai sebagai referensi desain model app [modul 25](25-integrasi-impor-ekspor-ekstensi.md). Lihat Di luar cakupan.

## API work item

### API-016 Issues API (CRUD)

- **Deskripsi:** Endpoint inti untuk membuat, membaca, mengubah, dan menghapus work item; fondasi semua integrasi sekaligus API yang dipakai SPA.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API, app; Browse Projects (baca), Create Issues (buat), Edit Issues (ubah), Delete Issues (hapus).
- **Aturan perilaku (acceptance criteria):**
  1. `POST /issue` menerima `fields`, `update`, `properties`, `transition` (status awal), dan `historyMetadata`, dengan parameter `updateHistory`.
  2. Respons create 201 berisi `id`, `key`, `self`, dan hasil `transition` bila diminta.
  3. `GET /issue/{issueIdOrKey}` mendukung `fields`, `fieldsByKeys`, `expand`, `properties`, `updateHistory=true` (mencatat ke daftar recent pengguna), dan `failFast` (gagal cepat bila field mahal error).
  4. `PUT /issue/{issueIdOrKey}` mengubah `fields`/`update` dengan `notifyUsers` (nilai false hanya untuk admin/app), `overrideScreenSecurity`, `overrideEditableFlag` (admin/app), `returnIssue`, dan `expand`; respons 204, atau 200 dengan issue bila `returnIssue=true`.
  5. `DELETE /issue/{key}?deleteSubtasks=true` menghapus beserta subtask; bila work item punya subtask dan flag false, ditolak 400; tanpa Delete Issues ditolak 403.
  6. Validasi field mengikuti screen dan field configuration (company-managed) atau layout issue type (team-managed); field yang tidak ada di screen ditolak 400 dengan pesan "Field cannot be set".
  7. Parent diset lewat `fields.parent` (key/id) untuk subtask dan hierarki termasuk Epic dan level di atasnya; parent dengan level hierarki tidak valid ditolak 400.
  8. Field sistem yang didukung: `summary`, `issuetype`, `project`, `description`, `priority`, `labels`, `components`, `fixVersions`, `versions`, `duedate`, `assignee`, `reporter`, `environment`, `timetracking`, `security`, `parent`.
  9. Custom field `customfield_NNNNN` menerima format per tipe: option `{value}` atau `{id}`, cascading `{value, child}`, user `{accountId}`, array, number, date, datetime; format salah ditolak 400 di `errors[customfield_NNNNN]`.
  10. Setelah work item dipindah atau key proyek diganti, key lama tetap resolve ke work item baru (GET mengembalikan key baru).
  11. Representasi issue memuat `id`, `key`, `self`, `fields`, dan sesuai expand `renderedFields`, `names`, `schema`, `transitions`, `operations`, `editmeta`, `changelog`, `versionedRepresentations`.
  12. ⚠ Create yang disertai `transition`/`properties` diproses berurutan; bila langkah lanjutan gagal setelah work item tercipta, respons menyatakan kegagalan sebagian dan work item tidak di-rollback.
  13. Pindah proyek/issue type tidak dapat dilakukan lewat PUT biasa (ditolak 400) dan harus memakai Bulk Move (API-026).
- **Team-managed vs company-managed:** Validasi field memakai layout issue type pada team-managed dan screen/field configuration pada company-managed; custom field team-managed tetap ber-id `customfield_xxx`.
- **Catatan & ketergantungan:** Logika domain dimiliki modul 06; event `jira:issue_*` dikirim ke webhook (API-065). Migrasi repo: feature baru `internal/features/issue` yang endpoint native-nya dipakai SPA dan diekspos ulang di lapisan kompatibel.

### API-017 Issue Create & Edit Metadata (createmeta/editmeta)

- **Deskripsi:** Endpoint metadata yang mengembalikan field tersedia, field wajib, dan allowed values untuk create atau edit, agar klien membangun form dinamis.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Integrator/Developer API, SPA; Create Issues (createmeta), Edit Issues (editmeta).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/createmeta/{projectIdOrKey}/issuetypes` (paginated) mengembalikan issue type yang boleh dibuat pemanggil di proyek tersebut.
  2. `GET /issue/createmeta/{projectIdOrKey}/issuetypes/{issueTypeId}` (paginated) mengembalikan field dengan `required`, `schema`, `name`, `key`, `hasDefaultValue`, `defaultValue`, `operations`, `allowedValues`, `autoCompleteUrl`.
  3. `GET /issue/createmeta` bentuk lama (`expand=projects.issuetypes.fields`) tidak disediakan dan mengembalikan 404 dengan pesan pengganti.
  4. `GET /issue/{key}/editmeta` mengembalikan field yang dapat diedit beserta `operations` (set/add/remove/edit) dan `allowedValues`.
  5. Parameter `overrideScreenSecurity` dan `overrideEditableFlag` pada editmeta hanya berlaku untuk admin/app; pengguna biasa yang mengirimnya ditolak 403.
  6. Hasil bergantung pada permission pengguna, screen create/edit, field configuration, dan status workflow; status dengan properti `jira.issue.editable=false` menghasilkan editmeta kosong.
  7. Pada team-managed field ditentukan dari layout per issue type tanpa screen scheme.
  8. Field user/label/option bernilai banyak menyertakan `autoCompleteUrl` alih-alih daftar `allowedValues` lengkap.
- **Team-managed vs company-managed:** Team-managed memakai layout per issue type; company-managed memakai screen scheme dan field configuration.
- **Catatan & ketergantungan:** Bergantung pada modul 05 (screen, field configuration) dan 09 (properti status).

### API-018 Issue Transitions API

- **Deskripsi:** Endpoint untuk melihat transisi workflow yang tersedia dan menjalankannya, termasuk mengisi field di transition screen.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Transition Issues plus condition workflow.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/{key}/transitions` mengembalikan `id`, `name`, `to` (status + statusCategory), `hasScreen`, `isGlobal`, `isInitial`, `isAvailable`, `isConditional`, `isLooped`, dan `fields` (dengan expand).
  2. Parameter yang didukung: `transitionId`, `expand=transitions.fields`, `skipRemoteOnlyCondition`, `includeUnavailableTransitions`, `sortByOpsBarAndStatus`.
  3. `POST /issue/{key}/transitions` menerima `transition.id`, `fields`, `update` (contoh menambah komentar), `historyMetadata`, dan `properties`.
  4. Validator workflow (field wajib, permission, expression) dijalankan; kegagalan dikembalikan 400 di `errors`/`errorMessages`.
  5. Condition menyembunyikan transisi yang syaratnya tidak terpenuhi, kecuali `includeUnavailableTransitions=true` yang menampilkannya dengan `isAvailable=false`.
  6. Post function dijalankan setelah transisi (set resolution, assign, update field, fire event, trigger webhook) dalam urutan yang dikonfigurasi.
  7. Resolution diset lewat field di transition screen atau post function pada company-managed; mengirim resolution tanpa field di screen ditolak 400.
  8. Pada team-managed default transisi any-to-any dengan rule transisi sederhana dan tanpa transition screen.
  9. Respons sukses 204 tanpa body.
- **Team-managed vs company-managed:** Team-managed: any-to-any, tanpa transition screen; company-managed: condition, validator, post function, dan screen penuh.
- **Catatan & ketergantungan:** Logika workflow di modul 09 (WF-041 untuk surface workflow). Event transisi diteruskan ke webhook dan otomasi.

### API-019 Issue Assign, Notify & Changelog API

- **Deskripsi:** Endpoint khusus untuk assign, mengirim notifikasi email ad-hoc, dan membaca riwayat perubahan work item.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Assign Issues, Browse Projects.
- **Aturan perilaku (acceptance criteria):**
  1. `PUT /issue/{key}/assignee` menerima `accountId`; `null` berarti unassigned (ditolak bila unassigned dimatikan global) dan `"-1"` berarti assignee default (component lead/project default).
  2. Assign memerlukan Assign Issues pada pemanggil dan Assignable User pada assignee; pelanggaran ditolak 403/400.
  3. `POST /issue/{key}/notify` menerima `subject`, `textBody`, `htmlBody`, `to` (reporter, assignee, watchers, voters, users, groups via `groupId`), dan `restrict` (permissions, groups).
  4. Notify ditolak 400 bila outgoing mail dimatikan; email ke diri sendiri mengikuti preferensi pengguna.
  5. `GET /issue/{key}/changelog` (paginated) mengembalikan id history, `author`, `created`, dan `items` (`field`, `fieldtype`, `fieldId`, `from`, `fromString`, `to`, `toString`).
  6. `POST /issue/{key}/changelog/list` mengambil changelog berdasarkan daftar id.
  7. `POST /changelog/bulkfetch` mengambil changelog banyak work item sekaligus dengan filter `fieldIds` dan `nextPageToken`.
  8. `expand=changelog` pada GET issue dibatasi ke history terbaru (dokumentasi menyebut batasnya; sisanya lewat endpoint changelog).
  9. `historyMetadata` (`type`, `description`, `actor`, `generator`, `cause`, `extraData`) yang dikirim saat write tersimpan dan ditampilkan sebagai konteks perubahan di activity.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Notifikasi kustom dirinci di NTF-035 (modul 19); history di modul 08.

### API-020 Comments API

- **Deskripsi:** CRUD komentar work item lewat API, termasuk pembatasan visibilitas dan properti komentar.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Add Comments, Edit All/Own Comments, Delete All/Own Comments.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/{key}/comment` (paginated) mendukung `orderBy=created|-created` dan `expand=renderedBody`.
  2. `POST /issue/{key}/comment` menerima `body` (ADF di v3), `visibility` `{type: role|group, value/identifier}`, dan `properties`; visibility ke role/grup yang tidak dimiliki pemanggil ditolak 400.
  3. `GET`/`PUT`/`DELETE /issue/{key}/comment/{id}` tersedia; PUT mendukung `notifyUsers` dan `overrideEditableFlag`.
  4. `POST /comment/list` mengambil komentar berdasarkan daftar id lintas work item (maksimum 1000), menyaring yang tidak boleh dilihat.
  5. Comment properties: `GET` keys dan `GET`/`PUT`/`DELETE /comment/{commentId}/properties/{key}`.
  6. Permission ditegakkan: Add Comments untuk membuat, Edit Own/All untuk mengubah, Delete Own/All untuk menghapus; pelanggaran ditolak 403.
  7. Representasi komentar memuat `author`, `updateAuthor`, `created`, `updated`, `visibility`, `jsdPublic`, dan `renderedBody` (bila di-expand).
  8. Node mention di ADF komentar memicu notifikasi sesuai API-003.
  9. Pada service project, komentar internal/publik ditentukan lewat properti `sd.public.comment` (`{"internal": true}`); tanpa properti komentar agent bersifat publik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain komentar di modul 08 (CMT-010); komentar internal/publik Service Management di [modul 29](29-jsm-service-project-request-type-formulir.md)/[35](35-jsm-ai-otomasi-api-admin-csm.md).

### API-021 Attachments API

- **Deskripsi:** Upload, unduh, metadata, thumbnail, dan penghapusan lampiran work item lewat API.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Create Attachments, Delete Own/All Attachments.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /issue/{key}/attachments` menerima `multipart/form-data` dengan field `file` (boleh banyak) dan wajib header `X-Atlassian-Token: no-check`.
  2. `GET /attachment/meta` mengembalikan `enabled` dan `uploadLimit` (byte).
  3. `GET /attachment/{id}` mengembalikan metadata `filename`, `author`, `created`, `size`, `mimeType`, URL `content`, dan `thumbnail`.
  4. `GET /attachment/content/{id}` mengunduh file (redirect ke storage), mendukung `redirect=false` untuk stream langsung dan header `Range` (206 Partial Content).
  5. `GET /attachment/thumbnail/{id}` mendukung `width`, `height`, `redirect`, dan `fallbackToDefault`; tanpa thumbnail dan `fallbackToDefault=false` menghasilkan 404.
  6. `GET /attachment/{id}/expand/human` dan `/expand/raw` mengembalikan daftar isi arsip ZIP; file non-arsip mengembalikan 409 atau daftar kosong.
  7. `DELETE /attachment/{id}` memerlukan Delete Own Attachments (milik sendiri) atau Delete All Attachments.
  8. Upload ditolak bila lampiran dimatikan global (403), Create Attachments tidak dimiliki (403), atau ukuran melebihi batas (413).
  9. File disimpan di object storage; ADF mereferensikan id node media yang terhubung ke lampiran (API-003).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain lampiran di modul 08 (CMT-011 s.d. CMT-016); storage di [modul 27](27-nfr-arsitektur-platform.md).

### API-022 Worklogs API

- **Deskripsi:** CRUD worklog pada work item beserta penyesuaian remaining estimate dan endpoint sinkronisasi worklog yang berubah.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Work On Issues, Edit/Delete Own/All Worklogs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/{key}/worklog` (paginated) mendukung `startedAfter`, `startedBefore`, dan `expand=properties`.
  2. `POST /issue/{key}/worklog` menerima `timeSpent` atau `timeSpentSeconds`, `started`, `comment` (ADF), `visibility`, dan `properties`; tanpa durasi ditolak 400.
  3. `adjustEstimate` menerima `new` (dengan `newEstimate`), `leave`, `manual` (dengan `reduceBy`), dan `auto`; DELETE memakai `increaseBy` untuk mode manual.
  4. `PUT`/`DELETE /issue/{key}/worklog/{id}` mendukung `notifyUsers` dan `overrideEditableFlag`.
  5. `POST /worklog/list` mengambil worklog berdasarkan id (maksimum 1000).
  6. `GET /worklog/updated?since=` dan `/worklog/deleted?since=` mengembalikan id worklog yang berubah untuk sinkronisasi (`values`, `since`, `until`, `lastPage`).
  7. Bulk delete worklog (`DELETE /issue/{key}/worklog` dengan `ids`) dan bulk move ke work item lain (`POST /issue/{key}/worklog/move`) tersedia.
  8. Worklog properties tersedia di `/issue/{key}/worklog/{id}/properties`.
  9. Semua endpoint ditolak 400/403 bila time tracking dimatikan; permission Work On Issues dan Edit/Delete Own/All Worklogs ditegakkan.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 14 (TIME-040, TIME-042).

### API-023 Issue Links, Link Types & Remote Links API

- **Deskripsi:** API untuk menghubungkan work item (link internal), mengelola tipe link, dan membuat remote link ke objek eksternal.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Link Issues, Browse Projects; tipe link dikelola Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /issueLink` menerima `type` `{name|id}`, `inwardIssue`, `outwardIssue`, dan `comment` opsional; respons 201.
  2. `GET`/`DELETE /issueLink/{linkId}` tersedia.
  3. `GET /issueLinkType`, `POST` (`name`, `inward`, `outward`), dan `GET`/`PUT`/`DELETE /issueLinkType/{id}` tersedia; nama duplikat ditolak 400.
  4. Link ditolak bila issue linking dimatikan global; memerlukan Link Issues pada work item sumber dan Browse pada kedua work item.
  5. `GET /issue/{key}/remotelink` mendukung filter `globalId`.
  6. `POST /issue/{key}/remotelink` bersifat create-or-update berdasarkan `globalId`, dengan `object` `{url, title, summary, icon, status {resolved, icon}}`, `application` `{type, name}`, dan `relationship`.
  7. `GET`/`PUT`/`DELETE /issue/{key}/remotelink/{linkId}` dan `DELETE` berdasarkan `globalId` tersedia.
  8. Remote link ke halaman wiki dan produk terkait memakai Application Links (API-014).
  9. GET issue menyertakan field `issuelinks` dengan representasi `inwardIssue`/`outwardIssue` dan tipe link.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain link di modul 06; batas link per work item di API-006.

### API-024 Votes & Watchers API

- **Deskripsi:** API untuk voting dan watching work item serta pengecekan status watch secara massal.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; View Voters and Watchers, Manage Watchers.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issue/{key}/votes` mengembalikan `votes`, `hasVoted`, dan `voters`; daftar voter hanya terisi bila pemanggil punya View Voters and Watchers.
  2. ⚠ `POST`/`DELETE /issue/{key}/votes` berlaku untuk pengguna saat ini; vote pada work item yang dilaporkan sendiri atau yang sudah resolved ditolak 400.
  3. Endpoint vote ditolak 404/400 bila voting dimatikan global.
  4. `GET /issue/{key}/watchers` mengembalikan `watchCount`, `isWatching`, dan `watchers`.
  5. `POST /issue/{key}/watchers` menerima body string `accountId` (kosong berarti pengguna saat ini); menambah pengguna lain memerlukan Manage Watchers.
  6. `DELETE /issue/{key}/watchers?accountId=` menghapus watcher; menghapus orang lain memerlukan Manage Watchers.
  7. `POST /issue/watching` mengecek `isWatching` secara bulk untuk daftar id work item.
  8. Endpoint watcher ditolak bila watching dimatikan global.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Toggle global di modul 22; bulk watch lewat API-026.

### API-025 Entity Properties API

- **Deskripsi:** Penyimpanan key-value JSON yang melekat pada entitas untuk kebutuhan integrasi, yang dapat diindeks agar bisa dicari lewat JQL.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Integrator/Developer API, Admin Situs; permission mengikuti entitasnya.
- **Aturan perilaku (acceptance criteria):**
  1. Issue properties: `GET` keys dan `GET`/`PUT`/`DELETE /issue/{key}/properties/{propertyKey}`.
  2. Bulk issue properties: `PUT`/`DELETE /issue/properties/{propertyKey}` dengan filter (`entityIds`, `currentValue`, `hasProperty`) berjalan sebagai async task (API-007).
  3. `POST /issue/properties/multi` men-set banyak properti ke banyak work item dalam satu request.
  4. Properti tersedia juga untuk project, issue type, comment, worklog, user, dashboard item, board, sprint, dan workflow transition dengan pola endpoint serupa.
  5. App properties tersedia di `/addons/{key}/properties` dan `/forge/app/properties` (API-072).
  6. Value maksimum 32 KB dan harus JSON valid; key maksimum 255 karakter; pelanggaran ditolak 400.
  7. Properti dapat diindeks lewat deklarasi modul app agar bisa dicari di JQL (`issue.property[key].path`).
  8. Permission mengikuti entitas: Edit Issues untuk issue property, Administer Projects untuk project property, dan seterusnya.
  9. Properti dapat diset saat create work item, komentar, dan worklog lewat field `properties`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Pencarian properti di JQL-036; project properties di PRJ-062; event `issue_property_set/deleted` di API-065.

### API-026 Bulk Issue APIs

- **Deskripsi:** Operasi massal: bulk create, bulk fetch, dan Bulk Operations API untuk edit, move, transition, delete, serta watch/unwatch secara async.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna, Admin Situs, Integrator/Developer API; global permission Make bulk changes plus permission per work item.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /issue/bulk` membuat s.d. 50 work item (termasuk subtask); respons memuat `issues` yang sukses dan `errors` per `elementNumber`.
  2. `POST /issue/bulkfetch` mengambil s.d. 100 work item berdasarkan id/key dengan `fields`, `expand`, `properties`, `fieldsByKeys`; yang tidak ditemukan dilaporkan di `issueErrors`.
  3. `GET /bulk/issues/fields` mengembalikan field yang dapat di-bulk edit untuk work item terpilih (dengan `searchText` dan paging).
  4. `POST /bulk/issues/fields` melakukan bulk edit (`selectedIssueIdsOrKeys`, `selectedActions`, `editedFieldsInput`, `sendBulkNotification`) untuk s.d. 1000 work item.
  5. `POST /bulk/issues/move` memindah proyek/issue type dengan `targetToSourcesMapping`, `inferFieldDefaults`, `inferStatusDefaults`, `inferSubtaskTypeDefault`, `targetMandatoryFields`, `targetStatus`, `targetClassification`.
  6. `GET /bulk/issues/transition` mengembalikan transisi yang tersedia per workflow dan `POST /bulk/issues/transition` menjalankannya.
  7. `POST /bulk/issues/delete`, `/bulk/issues/watch`, dan `/bulk/issues/unwatch` tersedia.
  8. `GET /bulk/queue/{taskId}` mengembalikan progres (`status`, `progressPercent`, `processedAccessibleIssues`, `failedAccessibleIssues`, `invalidOrInaccessibleIssueCount`, `totalIssueCount`).
  9. Operasi Bulk Operations memerlukan Make bulk changes beserta permission per work item; work item tanpa izin dihitung sebagai inaccessible, bukan menggagalkan seluruh job.
  10. ⚠ Batas jumlah work item, jumlah field per bulk, dan jumlah antrean aktif per pengguna terdokumentasi dan ditegakkan (kelebihan ditolak 400/429).
  11. `sendBulkNotification=false` menekan email massal untuk operasi tersebut.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** UI bulk change di ISS-035 dan SRC-031; notifikasi bulk di NTF-028; antrean job di [modul 27](27-nfr-arsitektur-platform.md).

### API-027 Issue Archiving API

- **Deskripsi:** API untuk mengarsipkan dan mengembalikan work item pada plan Premium/Enterprise, beserta ekspor daftar work item terarsip.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin Situs, Admin Proyek/Project Lead (sesuai konfigurasi).
- **Aturan perilaku (acceptance criteria):**
  1. `PUT /issue/archive` mengarsipkan berdasarkan daftar id/key (s.d. 1000) dengan error per work item.
  2. `POST /issue/archive` mengarsipkan berdasarkan JQL secara async (API-007).
  3. `PUT /issue/unarchive` mengembalikan work item terarsip (s.d. 1000).
  4. `PUT /issues/archive/export` mengekspor daftar work item terarsip (filter proyek, issue type, `archivedBy`, rentang tanggal) sebagai CSV yang dikirim lewat email.
  5. Work item terarsip bersifat read-only (write ditolak 400), tidak muncul di pencarian/board, dan key-nya tetap direservasi.
  6. Subtask ikut diarsip bersama parent; mengarsip work item yang sudah terarsip dilaporkan sebagai error per item.
  7. Memerlukan Admin Situs atau Admin Proyek sesuai konfigurasi; plan yang tidak memiliki entitlement menerima 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain arsip di ISS-034; pencarian work item terarsip di SRC-035; entitlement plan di VISI-003.

### API-028 Issue Redaction API

- **Deskripsi:** API untuk menyunting (redact) data sensitif dari field dan komentar work item secara permanen.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Admin Situs, org admin dengan permission khusus redaksi.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /redact` menjalankan redaksi async dengan daftar redaction (`externalId`, `reason`, `contentItem` `{entityType, entityId, id}`, `redactionPosition` `{expectedText, from, to, adfPointer}`); `expectedText` yang tidak cocok membuat item tersebut gagal.
  2. `GET /redact/status/{jobId}` mengembalikan status job redaksi.
  3. Teks yang diredaksi diganti penanda redaksi dan tidak dapat dipulihkan, termasuk dari history.
  4. Setiap redaksi tercatat di audit log dan history work item (tanpa isi asli).
  5. Hanya pengguna dengan permission redaksi yang diberikan admin yang dapat memanggil; lainnya ditolak 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat syarat plan belum pasti (pembanding: add-on keamanan tingkat atas). Domain di ISS-043 dan PERM-066.

## API pencarian & ekspresi

### API-029 Search API (JQL Enhanced Search)

- **Deskripsi:** Endpoint pencarian work item berbasis JQL dengan pagination token, hitungan perkiraan, dan konsistensi read-after-write.
- **Prioritas / Fase / Tier:** P0 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Browse Projects dan issue security per work item.
- **Aturan perilaku (acceptance criteria):**
  1. `GET`/`POST /search/jql` menerima `jql`, `nextPageToken`, `maxResults`, `fields` (default hanya `id`), `expand`, `properties`, `fieldsByKeys`, `failFast`, dan `reconcileIssues`.
  2. JQL wajib bounded (memiliki restriksi, contoh `project = X`); query unbounded ditolak 400 dengan pesan yang meminta restriksi.
  3. `reconcileIssues` menerima daftar id work item yang baru ditulis sehingga hasil memuat nilai terbaru walau indeks belum mutakhir.
  4. `POST /search/approximate-count` mengembalikan jumlah perkiraan hasil JQL.
  5. Endpoint lama `/rest/api/3/search` dan `/rest/api/2/search` (`startAt`/`total`, `validateQuery`) tidak disediakan dan mengembalikan 410/404 dengan pesan pengganti.
  6. Hasil otomatis disaring oleh permission dan issue security tanpa memberi tahu jumlah yang disaring.
  7. Error JQL dikembalikan di `errorMessages` dengan status 400.
  8. ⚠ `maxResults` maksimum lebih besar bila hanya `id`/`key` yang diminta (nilai terdokumentasi di OpenAPI).
  9. Expand per work item sama dengan GET issue (`renderedFields`, `names`, `schema`, `changelog`, `operations`).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Grammar dan evaluasi JQL di modul 12 (JQL-050 compiler, JQL-051 evaluasi berizin, JQL-052 batas); arsitektur indeks di SRC-001/SRC-002.

### API-030 JQL Utility & Picker APIs

- **Deskripsi:** Endpoint pendukung JQL dan picker: parsing, autocomplete, sanitasi, pencocokan, konversi pengguna, fungsi JQL app, dan issue picker.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Integrator/Developer API, SPA; Browse Projects, Browse users and groups untuk picker pengguna.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /jql/parse?validation=strict|warn|none` mengembalikan AST (`where`, `orderBy`, `clauses`, `operand`) beserta error per query.
  2. `GET`/`POST /jql/autocompletedata` mengembalikan `visibleFieldNames` (`value`, `displayName`, `cfid`, `operators`, `types`, `orderable`, `searchable`, `auto`), `visibleFunctionNames`, dan `jqlReservedWords`; versi POST dapat dibatasi `projectIds`.
  3. `GET /jql/autocompletedata/suggestions?fieldName=&fieldValue=&predicateName=&predicateValue=` mengembalikan saran nilai.
  4. `POST /jql/match` mengecek apakah daftar id work item cocok dengan beberapa JQL (`matchedIssues`, `errors`).
  5. `POST /jql/sanitize` menyanitasi query untuk `accountId` tertentu dengan menyamarkan referensi yang tidak boleh dilihatnya (contoh proyek privat).
  6. `POST /jql/pdcleaner` mengonversi `username`/`userKey` di JQL menjadi `accountId`.
  7. `GET`/`POST /jql/function/computation` menyediakan precomputation fungsi JQL milik app.
  8. `GET /issue/picker?query=&currentJQL=&currentIssueKey=&currentProjectId=&showSubTasks=&showSubTaskParent=` mengembalikan section History Search dan Current Search.
  9. `GET /groupuserpicker` mengembalikan saran pengguna dan grup sekaligus untuk field picker.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Detail perilaku di JQL-060, autocomplete di JQL-055, fungsi app di JQL-049.

### API-031 Jira Expressions API

- **Deskripsi:** Bahasa ekspresi mirip JavaScript untuk mengevaluasi data di server, dipakai app dan rule/kondisi workflow.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pengembang app); evaluasi mengikuti permission pemanggil.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /expression/evaluate` menerima `expression` dan `context`; endpoint lama `/expression/eval` tidak disediakan.
  2. Context yang didukung: `issue` (id/key), `issues` via JQL (dengan `nextPageToken`), `project`, `sprint`, `board`, `serviceDesk`, `customerRequest`, dan `custom`.
  3. `POST /expression/analyse?check=syntax|type|complexity` memvalidasi tanpa eksekusi dan mengembalikan error per ekspresi.
  4. Batas kompleksitas (`steps`, `expensiveOperations`, `beans`, `primitiveValues`) ditegakkan dan `meta.complexity` disertakan di respons; pelampauan ditolak 400.
  5. Ekspresi dapat dipakai di workflow condition/validator app, UI modifications, filter trigger app, dan display conditions.
  6. Evaluasi mengikuti permission pengguna pemanggil; data tanpa akses dievaluasi sebagai `null`.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Keputusan terbuka: membangun bahasa ekspresi sendiri atau mengadopsi engine yang ada (mis. CEL/expr); pembanding expression di modul 09 dan 20.

## Agile API

### API-032 Agile Boards API & Board Configuration

- **Deskripsi:** Agile REST API (`/rest/agile/1.0`) untuk board Scrum/Kanban: daftar, buat, hapus, konfigurasi, dan isi board.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, admin board, Integrator/Developer API; Browse Projects pada filter board, admin board untuk ubah.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /board` (paginated) mendukung `type` (`scrum`, `kanban`, `simple`), `name`, `projectKeyOrId`, `accountIdLocation`, `projectLocation`, `includePrivate`, `negateLocationFiltering`, `orderBy`, `expand`, `filterId`.
  2. `POST /board` menerima `name`, `type`, `filterId`, dan `location` `{type: project|user, projectKeyOrId}`; filter yang tidak boleh dilihat ditolak 400.
  3. `GET`/`DELETE /board/{boardId}` dan `GET /board/filter/{filterId}` tersedia.
  4. `GET /board/{id}/configuration` mengembalikan `filter`, `subQuery` (Kanban), `columnConfig` (columns, statuses, min/max, `constraintType`), `estimation` (field), `ranking` (`rankCustomFieldId`), dan `location`.
  5. `GET /board/{id}/issue` (JQL tambahan, `validateQuery`, `fields`, `expand`), `/board/{id}/epic`, `/board/{id}/epic/{epicId}/issue`, dan `/board/{id}/epic/none/issue` tersedia.
  6. `GET /board/{id}/project`, `/board/{id}/project/full`, dan `/board/{id}/version` (filter `released`) tersedia.
  7. `GET`/`PUT /board/{id}/features` men-toggle fitur board (backlog, sprints, estimation, reports, releases, code, deployments).
  8. `GET /board/{id}/quickfilter` dan `/quickfilter/{id}` bersifat read-only.
  9. `GET /board/{id}/reports` mengembalikan daftar report yang tersedia untuk board.
  10. Board properties: `GET`/`PUT`/`DELETE /board/{id}/properties/{key}`.
  11. `GET`/`PUT /board/{id}/settings/refined-velocity` tersedia.
  12. ⚠ Board team-managed (`type: simple`) berjumlah satu per proyek; endpoint create dan configure untuk board ini dibatasi (ditolak 400 dengan pesan yang jelas).
  13. Endpoint internal `/rest/greenhopper/1.0` tidak disediakan sebagai kontrak publik.
- **Team-managed vs company-managed:** Team-managed memakai satu board `simple` per proyek dengan konfigurasi terbatas; company-managed mendukung banyak board berbasis filter.
- **Catatan & ketergantungan:** Domain board di modul 10 (BRD-052, BRD-053).

### API-033 Agile Backlog & Rank API

- **Deskripsi:** API untuk memindahkan work item ke backlog dan mengubah urutan rank.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Schedule Issues dan Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /board/{id}/backlog` (paginated) mendukung `jql`, `fields`, `expand`, `validateQuery`.
  2. `POST /backlog/issue` memindahkan s.d. 50 work item ke backlog (keluar dari sprint).
  3. `POST /backlog/{boardId}/issue` memindahkan ke backlog board dengan `rankBeforeIssue`/`rankAfterIssue`.
  4. `PUT /issue/rank` menerima `issues` (maks 50), `rankBeforeIssue` atau `rankAfterIssue`, dan `rankCustomFieldId`; mengirim keduanya sekaligus ditolak 400.
  5. Respons rank adalah 207 Multi-Status dengan status per work item.
  6. Rank bersifat global (LexoRank) dan berlaku lintas board yang memakai field Rank yang sama.
  7. Memerlukan Schedule Issues dan Edit Issues; tanpa salah satunya work item tersebut gagal dengan 403 di entri 207.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain rank dan backlog di modul 11.

### API-034 Agile Sprints API

- **Deskripsi:** API sprint: buat, ubah, mulai, tutup, hapus, pindahkan work item, tukar urutan, dan properti sprint.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Scrum Master/Team Lead, pengguna, Integrator/Developer API; Manage Sprints.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /board/{id}/sprint` (paginated) mendukung filter `state` (`future`, `active`, `closed`).
  2. `POST /sprint` menerima `name` (maks 30 karakter), `startDate`, `endDate`, `originBoardId`, `goal`; nama lebih panjang ditolak 400.
  3. `GET`, `PUT` (full update), dan `POST` (partial update) `/sprint/{id}` tersedia.
  4. ⚠ Transisi state lewat update: `future` → `active` wajib `startDate` dan `endDate` (sprint aktif kedua hanya bila parallel sprints aktif), lalu `active` → `closed`; work item yang belum selesai saat close lewat API dipindah ke backlog kecuali target sprint ditentukan.
  5. ⚠ `DELETE /sprint/{id}` memindahkan work item terbuka ke backlog; sprint `closed` tidak dapat dihapus (400).
  6. `GET /sprint/{id}/issue` mendukung `jql`, `fields`, `expand`, `validateQuery`.
  7. `POST /sprint/{id}/issue` memindahkan s.d. 50 work item dengan `rankBeforeIssue`/`rankAfterIssue`; sprint `closed` menolak 400.
  8. `POST /sprint/{id}/swap` menukar urutan dengan sprint lain (`sprintToSwapWith`).
  9. Sprint properties: `GET`/`PUT`/`DELETE /sprint/{id}/properties/{key}`.
  10. Field Sprint (custom field) pada work item berisi id sprint dan dapat diset lewat Agile API maupun update field.
  11. Respons sprint memuat `id`, `self`, `state`, `name`, `startDate`, `endDate`, `completeDate`, `createdDate`, `originBoardId`, `goal`.
  12. Semua operasi tulis memerlukan Manage Sprints pada proyek board; tanpa itu ditolak 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan pada kontrak; team-managed memakai board `simple` tunggal.
- **Catatan & ketergantungan:** Domain sprint di modul 11 (SPR-023 parallel sprints); event sprint di API-065.

### API-035 Agile Epics & Issue API

- **Deskripsi:** Endpoint Agile untuk epic (legacy Epic Link) dan representasi work item agile, termasuk estimasi board.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Browse/Edit Issues.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /epic/{epicIdOrKey}` mengembalikan `id`, `key`, `name`, `summary`, `color`, `done`.
  2. `POST /epic/{epicIdOrKey}` melakukan partial update (`name`, `summary`, `color`, `done`).
  3. `GET /epic/{id}/issue` dan `POST /epic/{id}/issue` (memindahkan s.d. 50 work item ke epic) tersedia.
  4. `POST /epic/none/issue` melepaskan work item dari epic; `GET /epic/none/issue` mengembalikan work item tanpa epic.
  5. `PUT /epic/{id}/rank` me-rank epic relatif terhadap epic lain.
  6. Endpoint epic ditandai legacy (label Deprecated di dokumentasi) karena relasi epic memakai field `parent`; operasi di atas dipetakan ke `parent`.
  7. `GET /issue/{key}` (Agile) mengembalikan work item dengan field agile (`sprint`, `closedSprints`, `epic`, `flagged`).
  8. `GET`/`PUT /issue/{key}/estimation?boardId=` membaca/mengubah nilai estimasi sesuai konfigurasi board.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hierarki parent di modul 05/06; estimasi board di modul 14.

## API proyek

### API-036 Projects API

- **Deskripsi:** API inti untuk mencari, membuat, mengubah, mengarsipkan, dan menghapus proyek.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Admin Situs (buat/hapus), Admin Proyek/Project Lead (ubah), Integrator/Developer API; Browse Projects untuk membaca.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /project/search` (paginated) mendukung `query`, `typeKey` (`software`/`business`/`service_desk`/`product_discovery`), `categoryId`, `action` (`view`/`browse`/`edit`/`create`), `status` (`live`/`archived`/`deleted`), `expand` (`description`, `lead`, `issueTypes`, `url`, `projectKeys`, `permissions`, `insight`), `properties`, dan `orderBy`.
  2. `GET /project/{key}` mendukung expand; `GET /project/recent` mengembalikan proyek yang baru dilihat pemanggil.
  3. `POST /project` menerima `key`, `name`, `projectTypeKey`, `projectTemplateKey`, `leadAccountId`, `assigneeType` (`PROJECT_LEAD`/`UNASSIGNED`), `description`, `url`, `avatarId`, `categoryId`, serta id scheme (permission, notification, issueSecurity, workflow, issueTypeScheme, issueTypeScreenScheme, fieldConfiguration); hanya Admin Situs yang boleh, lainnya 403.
  4. ⚠ Template team-managed (contoh `...gh-simplified-agility-scrum`/`-kanban`) dibedakan dari template company-managed (`...gh-simplified-scrum-classic`/`-kanban-classic`), dan daftar template didokumentasikan.
  5. ⚠ Project templates API (`POST /project-template`) membuat proyek dari template custom/shared sebagai endpoint eksperimental.
  6. `PUT /project/{key}` mengubah atribut termasuk key; key lama tetap me-redirect ke proyek.
  7. `DELETE /project/{key}?enableUndo=true` memindahkan ke trash (60 hari); `POST /project/{key}/delete` menghapus permanen secara async.
  8. `POST /project/{key}/archive` (Premium) mengarsipkan dan `/restore` memulihkan dari trash atau arsip.
  9. `GET /projectvalidate/key`, `/validProjectKey`, dan `/validProjectName` memvalidasi serta menyarankan key/nama yang tersedia.
  10. Respons menyertakan `style` (`classic`/`next-gen`) dan `simplified` untuk membedakan company-managed dan team-managed.
- **Team-managed vs company-managed:** Dibedakan lewat template saat create dan field `style`/`simplified`; scheme id hanya berlaku untuk company-managed.
- **Catatan & ketergantungan:** Domain di modul 04 (PRJ-064, PRJ-033). Migrasi repo: feature baru `internal/features/project`.

### API-037 Project Configuration Sub-resources API

- **Deskripsi:** Sub-resource proyek: kategori, tipe, fitur, avatar, properti, email pengirim, hierarki, dan asosiasi scheme.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. Project categories: `GET`/`POST`/`PUT`/`DELETE /projectCategory` (hanya Admin Situs untuk tulis).
  2. Project types: `GET /project/type`, `/project/type/{key}`, `/project/type/{key}/accessible`, dan `/project/type/accessible`.
  3. `GET`/`PUT /project/{key}/features` men-toggle fitur proyek (backlog, sprints, releases, reports, code, deployments, pages, dsb.), terutama untuk team-managed.
  4. Project avatars: `GET /project/{key}/avatars`, `POST /project/{key}/avatar2` (crop `x`, `y`, `size`), `PUT` memilih avatar, `DELETE` avatar custom.
  5. Project properties: `GET`/`PUT`/`DELETE /project/{key}/properties/{propertyKey}`.
  6. `GET`/`PUT /project/{projectId}/email` membaca/mengubah alamat email pengirim notifikasi; alamat tidak terverifikasi ditolak 400.
  7. `GET /project/{projectId}/hierarchy` mengembalikan level hierarki issue type di proyek.
  8. `GET`/`PUT /project/{key}/permissionscheme`, `GET /project/{key}/notificationscheme`, dan `GET /project/{key}/issuesecuritylevelscheme` tersedia.
  9. Universal avatars: `GET /universal_avatar/type/{type}/owner/{id}`, `POST`, `DELETE`; `GET /avatar/{type}/system`; `GET /universal_avatar/view/type/{type}/avatar/{id}` (format png/svg, size xsmall s.d. xlarge).
  10. Default classification level proyek: `GET`/`PUT`/`DELETE /project/{key}/classification-default` (memerlukan entitlement klasifikasi).
- **Team-managed vs company-managed:** Toggle fitur terutama untuk team-managed; asosiasi scheme hanya bermakna untuk company-managed.
- **Catatan & ketergantungan:** Kategori di PRJ-019, properti di PRJ-062; klasifikasi di API-058.

### API-038 Project Components API

- **Deskripsi:** CRUD komponen proyek beserta lead, default assignee, dan jumlah work item terkait.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead; Browse Projects untuk membaca.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /project/{key}/component` (paginated: `query`, `orderBy`, `componentSource` `jira`/`compass`/`auto`) dan `/project/{key}/components` (semua) tersedia.
  2. `POST /component` menerima `name`, `description`, `project`, `leadAccountId`, `assigneeType` (`PROJECT_DEFAULT`, `COMPONENT_LEAD`, `PROJECT_LEAD`, `UNASSIGNED`); nama duplikat di proyek ditolak 400.
  3. `GET`/`PUT /component/{id}` mengembalikan `realAssignee` dan `isAssigneeTypeValid`.
  4. `DELETE /component/{id}?moveIssuesTo=` menghapus dan memindahkan work item ke komponen lain bila diberikan.
  5. `GET /component/{id}/relatedIssueCounts` mengembalikan jumlah work item terkait.
  6. `GET /component` (paginated, lintas proyek) tersedia.
  7. Komponen dari katalog layanan eksternal dapat dipakai sebagai komponen (`componentSource=compass`) bila integrasi tersebut ada.
  8. ⚠ Komponen di team-managed memiliki kesetaraan API dengan company-managed, dan perbedaan yang ada didokumentasikan.
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Domain di modul 15 (REL-034). Katalog layanan eksternal (Compass) adalah produk lain dan di luar cakupan.

### API-039 Project Versions API

- **Deskripsi:** CRUD versi/rilis, urutan, merge, release, arsip, dan related work.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead, release manager; Browse Projects untuk membaca.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /project/{key}/version` (paginated: `query`, `status` released/unreleased/archived, `orderBy` sequence/name/startDate/releaseDate/description, `expand` issuesstatus/operations/approvers/driver) dan `/project/{key}/versions` tersedia.
  2. `POST /version` menerima `name`, `description`, `projectId`, `startDate`, `releaseDate`, `released`, `archived`, `driver`; nama duplikat di proyek ditolak 400.
  3. `GET`/`PUT /version/{id}` tersedia, termasuk release dengan `moveUnfixedIssuesTo`.
  4. `POST /version/{id}/move` menerima `position` (`Earlier`/`Later`/`First`/`Last`) atau `after`.
  5. `PUT /version/{id}/mergeto/{moveIssuesTo}` menggabungkan versi.
  6. `POST /version/{id}/removeAndSwap` menghapus sambil memindahkan fixVersion/affectedVersion dan custom field version.
  7. `GET /version/{id}/relatedIssueCounts` dan `/unresolvedIssueCount` tersedia.
  8. Related work: `GET`/`POST`/`PUT`/`DELETE /version/{id}/relatedwork` (link desain, dokumen, komunikasi).
  9. Status `overdue` dihitung dari `releaseDate` yang sudah lewat pada versi unreleased.
  10. ⚠ Operasi tulis memerlukan Administer Projects atau Manage versions.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 15 (REL-027); event versi di API-065.

### API-040 Project Roles API

- **Deskripsi:** API untuk role global, default actors, dan actors per proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead (actors proyek), Admin Situs (role global).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /role`, `POST` (`name`, `description`), dan `GET`/`PUT`/`POST` (partial)/`DELETE /role/{id}` dengan `swap` tersedia; hapus role yang dipakai scheme tanpa `swap` ditolak 400.
  2. `GET`/`POST`/`DELETE /role/{id}/actors` mengelola default actors untuk proyek baru.
  3. `GET /project/{key}/role` mengembalikan map nama role ke URL.
  4. `GET /project/{key}/role/{id}` mengembalikan actors (`atlassian-user-role-actor`, `atlassian-group-role-actor`).
  5. `POST`/`PUT`/`DELETE /project/{key}/role/{id}` menambah, men-set, atau menghapus actor (pengguna `accountId`, grup `groupId`).
  6. `GET /project/{key}/roledetails?currentMember=&excludeConnectAddons=` tersedia.
  7. Pada team-managed role ber-scope proyek (Administrator, Member, Viewer, plus custom role di plan berbayar).
- **Team-managed vs company-managed:** Team-managed memakai role per proyek; company-managed memakai role global dengan actor per proyek.
- **Catatan & ketergantungan:** Domain di modul 03 (PERM-061).

## API konfigurasi work item, skema & workflow

### API-041 Fields & Custom Fields API

- **Deskripsi:** API untuk membaca semua field dan mengelola custom field global, termasuk trash dan restore.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs (kelola), Integrator/Developer API (baca).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /field` mengembalikan semua field sistem dan custom (`id`, `key`, `name`, `custom`, `orderable`, `navigable`, `searchable`, `clauseNames`, `schema`, dan `scope` untuk team-managed).
  2. `GET /field/search` (paginated) mendukung filter `type` (custom/system), `id`, `query`, `orderBy` (`contextsCount`, `lastUsed`, `name`, `screensCount`), dan `expand` (`key`, `lastUsed`, `screensCount`, `contextsCount`, `isLocked`, `searcherKey`).
  3. `POST /field` membuat custom field (`name`, `description`, `type`, `searcherKey`); tipe tidak dikenal ditolak 400.
  4. `PUT /field/{fieldId}` mengubah nama, deskripsi, dan `searcherKey`.
  5. `POST /field/{id}/trash` dan `/field/{id}/restore` tersedia; `GET /field/search/trashed` mendaftar field di trash, yang dihapus permanen setelah 60 hari.
  6. `DELETE /field/{id}` menghapus permanen sebagai async task (API-007).
  7. Locked field (dikelola sistem/app) tidak dapat diubah atau dihapus; upaya tersebut ditolak 403.
  8. Tipe custom field bawaan: short text, paragraph, number, date, datetime, select single/multi, radio, checkbox, cascading select, user picker single/multi, group picker, labels, URL, version picker, project picker.
  9. ⚠ Field associations mengaitkan field ke proyek/issue type secara bulk.
  10. App custom field value: `PUT /app/field/value` dan `/app/field/{fieldIdOrKey}/value` memperbarui nilai oleh app pemilik tanpa notifikasi; app bukan pemilik ditolak 403.
  11. App custom field configuration: `GET`/`PUT /app/field/{fieldIdOrKey}/context/configuration` (configuration dan schema untuk custom field app).
  12. Custom field team-managed bersifat lokal proyek namun tetap ber-id `customfield_xxx`.
- **Team-managed vs company-managed:** Custom field team-managed ber-scope proyek; company-managed global dengan context.
- **Catatan & ketergantungan:** Domain di modul 05.

### API-042 Custom Field Contexts & Options API

- **Deskripsi:** Pengelolaan context custom field (scope proyek/issue type, default value) dan opsi untuk field select/cascading.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET`/`POST /field/{fieldId}/context` dan `PUT`/`DELETE /field/{fieldId}/context/{contextId}` tersedia.
  2. Context dapat global atau dibatasi ke proyek dan issue type tertentu; satu pasangan proyek/issue type hanya boleh tercakup satu context (konflik ditolak 400).
  3. `PUT` dan `POST .../remove` pada `/field/{fieldId}/context/{contextId}/issuetype` dan `/project` menambah/melepas cakupan.
  4. `GET /field/{fieldId}/context/issuetypemapping` dan `/projectmapping` tersedia.
  5. `POST /field/{fieldId}/context/mapping` mengembalikan context untuk pasangan proyek/issue type.
  6. `GET`/`PUT /field/{fieldId}/context/defaultValue` mendukung tipe option, cascading, user, group, date, datetime, text, number, URL, version, project, labels.
  7. `GET /field/{fieldId}/context/{contextId}/option` (paginated), `POST` create, `PUT` update (`value`, `disabled`), dan `DELETE` tersedia.
  8. `PUT .../option/move` mengurutkan ulang (`position` First/Last atau `after`).
  9. `DELETE .../option/{optionId}/issue?replaceWith=` mengganti nilai pada work item lalu menghapus opsi (async).
  10. Cascading select mendukung opsi parent dan child (`optionId` parent).
  11. App-owned field options API (`/field/{fieldKey}/option`) untuk field app legacy tersedia sebagai rujukan kompatibilitas.
- **Team-managed vs company-managed:** Team-managed tidak memakai context; opsi diatur langsung per field proyek.
- **Catatan & ketergantungan:** Domain di modul 05.

### API-043 Issue Types & Schemes API

- **Deskripsi:** API untuk issue type, issue type scheme, dan issue type screen scheme beserta asosiasinya ke proyek.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issuetype`, `GET /issuetype/project?projectId=&level=`, `POST` (`name`, `description`, `type`/`hierarchyLevel`), dan `GET`/`PUT`/`DELETE /issuetype/{id}?alternativeIssueTypeId` tersedia; hapus tanpa alternatif saat masih dipakai ditolak 400.
  2. `GET /issuetype/{id}/alternatives` mengembalikan pengganti yang valid saat issue type dihapus.
  3. `POST /issuetype/{id}/avatar2` mengunggah avatar issue type.
  4. `hierarchyLevel`: -1 subtask, 0 standard, 1 epic; level di atas epic dikonfigurasi lewat issue type hierarchy (Premium).
  5. Issue type properties tersedia di `/issuetype/{id}/properties`.
  6. Issue type schemes: `GET`/`POST`/`PUT`/`DELETE /issuetypescheme`, mapping, tambah/hapus/reorder issue type, dan `defaultIssueTypeId`.
  7. `PUT /issuetypescheme/project` mengasosiasikan scheme ke proyek; gagal 400 bila ada work item yang memakai type yang tidak ada di scheme baru.
  8. Issue type screen schemes: CRUD `/issuetypescreenscheme`, mapping issue type ke screen scheme, default mapping, dan asosiasi proyek.
  9. Pada team-managed issue type milik proyek (`scope: PROJECT`) tanpa scheme.
- **Team-managed vs company-managed:** Lihat butir 9.
- **Catatan & ketergantungan:** Domain di modul 05 (CFG-006, CFG-007).

### API-044 Priorities, Priority Schemes & Resolutions API

- **Deskripsi:** API untuk mengelola prioritas global, priority scheme per proyek, dan resolution.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /priority` dan `GET /priority/search` (paginated: `id`, `projectId`, `priorityName`, `onlyDefault`, `expand`) tersedia.
  2. `POST`/`PUT /priority` menerima `name`, `description`, `statusColor`, `iconUrl`/`avatarId`; `DELETE` berjalan async.
  3. `PUT /priority/move` mengurutkan ulang dan `PUT /priority/default` menetapkan default.
  4. Priority schemes: `GET`/`POST /priorityscheme`, `PUT`/`DELETE /priorityscheme/{id}`, `/priorityscheme/{id}/priorities`, dan `/priorityscheme/{id}/projects`.
  5. Menghapus prioritas dari scheme yang masih dipakai work item mewajibkan remap (`mappings` in/out); tanpa mapping ditolak 400.
  6. `GET /priorityscheme/priorities/available` dan `POST /priorityscheme/mappings` mendukung dialog migrasi.
  7. Resolutions: `GET /resolution`, `/resolution/search`, `POST`, `PUT`, `DELETE` async dengan `replaceWith`, `/resolution/move`, dan `/resolution/default`.
  8. ⚠ Team-managed memakai prioritas global yang sama; resolution di team-managed diset otomatis saat work item masuk status kategori Done.
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Domain di modul 05 dan 09.

### API-045 Statuses & Status Categories API

- **Deskripsi:** API untuk status workflow (bulk) dan kategori status (To Do, In Progress, Done).
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Admin Situs, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /statuses?id=` (bulk) dan `GET /statuses/search` (paginated: `projectId`, `searchString`, `statusCategory`) tersedia.
  2. `POST /statuses` (bulk create dengan scope `GLOBAL` atau `PROJECT` + `projectId`), `PUT /statuses` (bulk update), dan `DELETE /statuses?id=` tersedia.
  3. `GET /statuses/{statusId}/project/{projectId}/issueTypeUsages`, `/statuses/{statusId}/projectUsages`, dan `/workflowUsages` tersedia.
  4. `GET /status` dan `/status/{idOrName}` (legacy) tetap tersedia untuk kompatibilitas.
  5. `GET /statuscategory` dan `/statuscategory/{idOrKey}` mengembalikan `undefined`, `new`, `indeterminate`, `done` beserta `colorName`.
  6. `GET /project/{key}/statuses` mengembalikan status per issue type di proyek.
  7. Status yang masih dipakai workflow tidak dapat dihapus (ditolak 400 dengan daftar workflow pemakai).
- **Team-managed vs company-managed:** Status team-managed ber-scope `PROJECT`; company-managed `GLOBAL`.
- **Catatan & ketergantungan:** Domain di modul 09.

### API-046 Screens, Screen Schemes & Field Configurations API

- **Deskripsi:** API untuk screen, tab, field pada screen, screen scheme, field configuration, dan field configuration scheme (company-managed).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /screens` (paginated: `id`, `queryString`, `scope`, `orderBy`), `POST`, `PUT`, dan `DELETE` tersedia; screen yang dipakai scheme tidak dapat dihapus.
  2. `GET /screens/tabs` (bulk), CRUD `/screens/{id}/tabs`, dan `POST /screens/{id}/tabs/{tabId}/move/{pos}` tersedia.
  3. `GET`/`POST /screens/{id}/tabs/{tabId}/fields`, `DELETE` field, dan `POST` field move (`after`/`position`) tersedia.
  4. `GET /screens/{id}/availableFields` dan `POST /screens/addToDefault/{fieldId}` tersedia.
  5. Screen schemes: CRUD `/screenscheme` dengan screen per operasi default/create/edit/view.
  6. Field configurations: CRUD `/fieldconfiguration` dan `GET`/`PUT /fieldconfiguration/{id}/fields` (`description`, `isHidden`, `isRequired`, `renderer`).
  7. Field configuration schemes: CRUD, mapping issue type ke field configuration, dan `PUT /fieldconfigurationscheme/project`.
  8. Default Field Configuration tidak dapat dihapus (ditolak 400).
  9. Team-managed tidak memakai screen/field configuration; layout field diatur per issue type.
  10. ⚠ Field schemes (penyederhanaan konfigurasi field) beserta API-nya disediakan bila model field scheme diadopsi modul 05.
- **Team-managed vs company-managed:** Lihat butir 9.
- **Catatan & ketergantungan:** Domain di modul 05.

### API-047 Workflows API

- **Deskripsi:** API workflow generasi baru (bulk) dan legacy untuk membaca, membuat, memvalidasi, dan memperbarui workflow beserta status, transisi, dan rule.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Admin Proyek (team-managed), Integrator/Developer API (pengembang app).
- **Aturan perilaku (acceptance criteria):**
  1. `POST /workflows` mengambil workflow berdasarkan id/nama dengan expand (`workflows.usages`, `statuses.usages`).
  2. `GET /workflows/search` (paginated: `queryString`, `orderBy`, scope `GLOBAL`/`PROJECT`, `isActive`, `expand`) tersedia.
  3. `POST /workflows/create` dan `/workflows/create/validation` menerima workflow dengan `statuses` (`statusReference`), `transitions` (from/to, type initial/global/directed), `rules` (conditions, validators, actions), dan `triggers`.
  4. `POST /workflows/update` dan `/workflows/update/validation` mendukung `statusMappings` untuk work item yang statusnya dihapus; tanpa mapping yang diperlukan ditolak 400.
  5. `GET /workflows/capabilities` mengembalikan rule dan trigger yang tersedia untuk workflow/proyek.
  6. `GET /workflows/defaultEditor` mengembalikan editor `new` atau `legacy`.
  7. `GET /workflow/{id}/projectUsages`, `/issueTypeUsages`, dan `/workflowSchemes` tersedia.
  8. `DELETE /workflow/{entityId}` hanya untuk workflow inactive; workflow aktif ditolak 400.
  9. Legacy: `GET /workflow/search`, `POST /workflow` (deprecated), dan transition properties `/workflow/transitions/{id}/properties` tersedia untuk kompatibilitas.
  10. Rule transisi milik app: `GET`/`PUT /workflow/rule/config` dan `PUT /workflow/rule/config/delete`.
  11. Team-managed: workflow per issue type dengan scope `PROJECT` + `projectId`.
  12. Properti transisi dan status (contoh `jira.issue.editable`, `opsbar-sequence`, `jira.permission.*`) dapat dibaca dan ditulis.
  13. Workflow memiliki `version` (`id`, `versionNumber`) untuk optimistic concurrency; update dengan versi usang ditolak 409.
- **Team-managed vs company-managed:** Lihat butir 11.
- **Catatan & ketergantungan:** Domain di modul 09 (WF-041, WF-031, WF-025).

### API-048 Workflow Schemes & Drafts API

- **Deskripsi:** API workflow scheme: mapping issue type ke workflow, draft untuk scheme aktif, publish dengan migrasi status, dan asosiasi proyek.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /workflowscheme` (paginated), `POST`, dan `GET`/`PUT`/`DELETE /workflowscheme/{id}` tersedia; scheme yang dipakai proyek tidak dapat dihapus.
  2. `GET`/`PUT`/`DELETE /workflowscheme/{id}/default` dan `/workflowscheme/{id}/issuetype/{issueType}` tersedia.
  3. `GET`/`PUT`/`DELETE /workflowscheme/{id}/workflow?workflowName=` tersedia.
  4. `POST /workflowscheme/{id}/createdraft` dan `GET`/`PUT`/`DELETE /workflowscheme/{id}/draft` beserta sub-resource default/issuetype/workflow tersedia.
  5. `POST /workflowscheme/{id}/draft/publish` dengan `statusMappings` berjalan sebagai async task.
  6. ⚠ `GET`/`PUT /workflowscheme/project` membaca/menetapkan scheme per proyek; mengganti scheme proyek yang memerlukan migrasi status berjalan async dengan status mapping.
  7. `POST /workflowscheme/update/mappings` mempratinjau status mapping yang dibutuhkan.
  8. ⚠ Bulk read/update workflow scheme generasi baru (`POST /workflowscheme/read`, `/workflowscheme/update`) tersedia.
  9. Pada company-managed, scheme aktif hanya dapat diubah lewat draft; `PUT` langsung ke scheme aktif ditolak 400.
- **Team-managed vs company-managed:** Team-managed tidak memakai workflow scheme.
- **Catatan & ketergantungan:** Domain di modul 09.

### API-049 Issue Security Schemes API

- **Deskripsi:** API untuk security scheme, level keamanan work item, anggota level, dan asosiasi proyek (company-managed).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs; Set Issue Security untuk men-set level pada work item.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /issuesecurityschemes`, `/issuesecurityschemes/search`, `POST` (dengan levels dan members), dan `GET`/`PUT`/`DELETE /issuesecurityschemes/{id}` tersedia.
  2. `GET /issuesecurityschemes/level`, `PUT` tambah level, `PUT` update level, dan `DELETE` level dengan `replaceWith` (async) tersedia.
  3. `GET /issuesecurityschemes/level/member`, `PUT` tambah member, dan `DELETE` member tersedia.
  4. Member type: `group`, `user`, `projectRole`, `reporter`, `assignee`, `projectLead`, `applicationRole`, `userCustomField`, `groupCustomField`.
  5. `PUT /issuesecurityschemes/project` mengasosiasikan ke proyek dengan `oldToNewSecurityLevelMappings` (async).
  6. Setiap scheme dapat memiliki default security level yang diterapkan saat create.
  7. `GET /project/{key}/securitylevel` (level yang dapat dipilih pengguna) dan `/project/{key}/issuesecuritylevelscheme` tersedia.
  8. Field `security` pada work item diset lewat `fields.security.id` dan memerlukan Set Issue Security; tanpa itu ditolak 400.
  9. ⚠ Team-managed tidak memakai security scheme; pembatasan work item di team-managed memakai mekanisme proyek sendiri yang didokumentasikan terpisah.
- **Team-managed vs company-managed:** Lihat butir 9.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat ketersediaan issue security per plan (mulai Standard) belum pasti. Domain di modul 03 (PERM-062) dan ISS-042.

## API pengguna, grup, izin & notifikasi

### API-050 Users, Myself & Preferences API

- **Deskripsi:** API pencarian dan detail pengguna berbasis accountId, pengguna saat ini, preferensi, properti, dan kolom navigator.
- **Prioritas / Fase / Tier:** P0 / Fase 1 / free-standard
- **Aktor & izin:** Pengguna, Integrator/Developer API; Browse users and groups untuk pencarian.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /user?accountId=&expand=groups,applicationRoles` dan `GET /user/bulk` (s.d. 200 accountId) tersedia.
  2. `GET /user/search?query=&accountId=&property=` dan `GET /users/search` (semua pengguna, paginated) tersedia; tanpa Browse users and groups hasil kosong.
  3. `GET /user/search/query` dan `/user/search/query/key` mendukung pencarian terstruktur (`is assignee of`, `is watcher of`, `[propertyKey].entity.property.path is value`).
  4. `GET /user/picker?query=&showAvatar=&exclude=` tersedia.
  5. `GET /user/assignable/search` (project/issueKey) dan `/user/assignable/multiProjectSearch` hanya mengembalikan pengguna dengan Assignable User.
  6. `GET /user/viewissue/search` dan `/user/permission/search?permissions=` tersedia.
  7. `GET /user/groups?accountId=` mengembalikan grup pengguna.
  8. `GET /user/email` dan `/user/email/bulk` hanya untuk app yang diizinkan; email yang disembunyikan pengaturan privasi profil tidak dikembalikan.
  9. `POST /user` dan `DELETE /user` bersifat eksperimental/terbatas; manajemen akun diarahkan ke Admin API (API-011).
  10. `GET`/`PUT`/`DELETE /user/properties/{key}` tersedia (milik sendiri, atau admin untuk pengguna lain).
  11. `GET`/`PUT`/`DELETE /user/columns` mengelola kolom issue navigator milik pengguna.
  12. ⚠ User nav properties (`GET`/`PUT /user/nav4-opt-property/{key}`) menyimpan preferensi navigasi.
  13. `GET /myself?expand=groups,applicationRoles` mengembalikan pengguna terautentikasi.
  14. `GET`/`PUT`/`DELETE /mypreferences?key=` tersedia; `GET`/`PUT /mypreferences/locale` disediakan sebagai deprecated.
  15. Atribut pengguna: `accountType` (`atlassian`, `app`, `customer`), `active`, `timeZone`, `locale`, `avatarUrls`, `displayName`; field yang tampil mengikuti profile visibility.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 02 (USR-025) dan 21. Migrasi repo: feature `internal/features/user` sudah ada; perlu kolom `account_id` (UUID publik) terpisah dari id internal, dan `/myself` dipetakan ke handler profil yang ada.

### API-051 Groups & Application Roles API

- **Deskripsi:** API grup berbasis groupId, keanggotaan, picker, serta application role dan lisensi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, org admin; Browse users and groups untuk picker.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /group` (`name`) dan `DELETE /group?groupId=&swapGroupId=` (memindahkan grant ke grup lain) tersedia.
  2. `GET /group/bulk` (paginated: `groupId`, `groupName`, `accessType`, `applicationKey`) tersedia.
  3. `GET /group/member?groupId=&includeInactiveUsers=` tersedia.
  4. `POST /group/user` dan `DELETE /group/user` (accountId) tersedia.
  5. `GET /groups/picker?query=&exclude=&caseInsensitive=` tersedia.
  6. Parameter `groupname` deprecated; `groupId` dipakai di semua endpoint baru.
  7. Grup dikelola di level organisasi; grup yang disinkron dari IdP tidak dapat diubah lewat API ini (ditolak 400).
  8. `GET /applicationrole` dan `/applicationrole/{key}` mengembalikan `key` (software, service desk, product discovery), `groups`, `defaultGroups`, `numberOfSeats`, `remainingSeats`, `userCount`, `hasUnlimitedSeats`, `platform`.
  9. `GET /instance/license` mengembalikan produk dan plan (`FREE`, `STANDARD`, `PREMIUM`, `ENTERPRISE`).
  10. `GET /license/approximateLicenseCount` dan `/license/approximateLicenseCount/product/{key}` tersedia.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 02 dan 03 (PERM-063).

### API-052 Permissions & Permission Schemes API

- **Deskripsi:** API untuk mengecek permission pengguna dan mengelola permission scheme beserta grant-nya.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (scheme), Integrator/Developer API (cek izin).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /mypermissions?permissions=&projectKey=&issueKey=&projectId=&issueId=` dengan parameter `permissions` wajib (tanpa parameter ditolak 400).
  2. `GET /permissions` mengembalikan semua permission (global, project, dan yang didefinisikan app).
  3. `POST /permissions/check` mengecek secara bulk global dan project permission untuk `accountId` atau pengguna saat ini.
  4. `POST /permissions/project` mengembalikan proyek yang diizinkan untuk permission tertentu.
  5. `GET`/`POST`/`PUT`/`DELETE /permissionscheme` dengan expand (`permissions`, `user`, `group`, `projectRole`, `field`, `all`) tersedia.
  6. `GET`/`POST /permissionscheme/{id}/permission` dan `GET`/`DELETE` grant tertentu tersedia.
  7. Holder type: `anyone`, `group`, `user`, `projectRole`, `applicationRole`, `reporter`, `projectLead`, `assignee`, `userCustomField`, `groupCustomField`, `sd.customer.portal.only`.
  8. Permission key mencakup `BROWSE_PROJECTS`, `CREATE_ISSUES`, `EDIT_ISSUES`, `TRANSITION_ISSUES`, `ASSIGN_ISSUES`, `ASSIGNABLE_USER`, `SCHEDULE_ISSUES`, `MANAGE_SPRINTS_PERMISSION`, `ADMINISTER_PROJECTS`, dll.
  9. Global permission (`ADMINISTER`, `SYSTEM_ADMIN`, `USER_PICKER`, `CREATE_SHARED_OBJECTS`, `BULK_CHANGE`, `MANAGE_GROUP_FILTER_SUBSCRIPTIONS`) hanya dapat dicek lewat API; grant-nya dikelola lewat UI.
  10. Pada team-managed permission berasal dari role dan access level proyek (open/limited/private), bukan scheme.
- **Team-managed vs company-managed:** Lihat butir 10.
- **Catatan & ketergantungan:** Domain di modul 03 (PERM-058 s.d. PERM-060).

### API-053 Notification Schemes & Issue Events API

- **Deskripsi:** API notification scheme (company-managed) dan daftar issue event yang dipakai notifikasi dan post function.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /notificationscheme` (paginated, expand: `all`, `field`, `group`, `notificationSchemeEvents`, `projectRole`, `user`) tersedia.
  2. `POST /notificationscheme` menerima `name`, `description`, `notificationSchemeEvents`.
  3. `GET`/`PUT /notificationscheme/{id}` dan `PUT /notificationscheme/{id}/notification` (menambah notifikasi ke event) tersedia.
  4. `DELETE` scheme dan `DELETE` notifikasi tertentu tersedia; scheme yang dipakai proyek tidak dapat dihapus.
  5. `GET /notificationscheme/project` mengembalikan mapping scheme ke proyek.
  6. Recipient type: `CurrentAssignee`, `Reporter`, `CurrentUser`, `ProjectLead`, `ComponentLead`, `User`, `Group`, `ProjectRole`, `EmailAddress`, `AllWatchers`, `UserCustomField`, `GroupCustomField`.
  7. `GET /events` mengembalikan daftar issue event (Issue Created, Updated, Assigned, Resolved, Closed, Commented, Comment Edited, Comment Deleted, Reopened, Work Started/Stopped, Work Logged, Worklog Updated, Worklog Deleted, Issue Moved, Issue Deleted, Generic Event, dan custom event).
  8. Pada team-managed notifikasi diatur per proyek tanpa scheme.
- **Team-managed vs company-managed:** Lihat butir 8.
- **Catatan & ketergantungan:** Domain di modul 19 (NTF-001, NTF-003, NTF-036).

## API objek bersama, konfigurasi global & perencanaan

### API-054 Filters API

- **Deskripsi:** API filter tersimpan: CRUD, pencarian, favorit, sharing, edit permission, kolom, dan pergantian pemilik.
- **Prioritas / Fase / Tier:** P1 / Fase 2 / free-standard
- **Aktor & izin:** Pengguna, Admin Situs; Create Shared Objects untuk berbagi.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /filter` menerima `name`, `description`, `jql`, `favourite`, `sharePermissions`, `editPermissions`, dengan parameter `overrideSharePermissions` (admin saja).
  2. `GET /filter/search` (paginated) mendukung `filterName`, `accountId`, `groupname`/`groupId`, `projectId`, `id`, `orderBy` (description, favourite_count, is_favourite, id, is_shared, name, owner), `expand` (description, favourite, favouritedCount, jql, owner, searchUrl, sharePermissions, editPermissions, isWritable, approximateLastUsed, subscriptions, viewUrl), dan `overrideSharePermissions`.
  3. `GET`/`PUT`/`DELETE /filter/{id}` tersedia; ubah/hapus oleh bukan pemilik/editor ditolak 403.
  4. `GET /filter/my?includeFavourites=` dan `GET /filter/favourite` tersedia.
  5. `PUT`/`DELETE /filter/{id}/favourite` tersedia.
  6. `GET`/`POST /filter/{id}/permission` dan `GET`/`DELETE /filter/{id}/permission/{permissionId}` tersedia.
  7. Share type: global/public (dapat dibatasi admin), project, project + role, group, user, authenticated (loggedin).
  8. `GET`/`PUT /filter/defaultShareScope` (`GLOBAL`, `AUTHENTICATED`, `PRIVATE`) tersedia.
  9. `GET`/`PUT`/`DELETE /filter/{id}/columns` tersedia.
  10. `PUT /filter/{id}/owner` mengganti pemilik (admin saja).
  11. Berbagi filter memerlukan global permission Create Shared Objects; tanpa itu `sharePermissions` non-privat ditolak 400.
  12. ⚠ Subscription hanya terlihat lewat expand `subscriptions`; CRUD subscription tidak tersedia di API publik.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 13 (SRC-041 s.d. SRC-049).

### API-055 Dashboards & Gadgets API

- **Deskripsi:** API dashboard dan gadget: CRUD, salin, bulk edit, sharing, dan properti item.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Pengguna, Admin Situs; Create Shared Objects untuk berbagi.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /dashboard?filter=my|favourite` (paginated) dan `/dashboard/search` (`dashboardName`, `accountId`, `groupId`, `projectId`, `orderBy`, `status` active/archived/deleted, `expand`) tersedia.
  2. `POST /dashboard` menerima `name`, `description`, `sharePermissions`, `editPermissions`, dengan `extendAdminPermissions`.
  3. `GET`/`PUT`/`DELETE /dashboard/{id}` dan `POST /dashboard/{id}/copy` tersedia.
  4. `PUT /dashboard/bulk/edit` mendukung action `changeOwner`, `changePermission`, `addPermission`, `removePermission`.
  5. `GET /dashboard/gadgets` mengembalikan gadget yang tersedia (`moduleKey`, `uri`, `title`).
  6. `GET`/`POST /dashboard/{id}/gadget` dan `PUT`/`DELETE /dashboard/{id}/gadget/{gadgetId}` (`color`, `position` row/column, `title`, `ignoreUriAndModuleKeyValidation`) tersedia; posisi bentrok ditolak 400.
  7. Dashboard item properties: `GET`/`PUT`/`DELETE /dashboard/{id}/items/{itemId}/properties/{key}` menyimpan konfigurasi gadget.
  8. ⚠ Layout kolom dashboard diatur lewat UI; bila API menyediakannya, layout disertakan di representasi dashboard.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 17 (DSH-056, DSH-057).

### API-056 Server Info & Global Configuration API

- **Deskripsi:** API informasi instance dan pengaturan global: server info, konfigurasi fitur, application properties, time tracking, banner, label, dan kolom default.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (tulis), Integrator/Developer API (baca).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /serverInfo` mengembalikan `baseUrl`, `displayUrl`, `version`, `versionNumbers`, `deploymentType`, `buildNumber`, `buildDate`, `serverTime`, `scmInfo`, `serverTitle`, `defaultLocale`, `serverTimeZone`; dapat dipanggil anonim dengan subset field.
  2. `GET /configuration` mengembalikan `votingEnabled`, `watchingEnabled`, `unassignedIssuesAllowed`, `subTaskingEnabled`, `issueLinkingEnabled`, `timeTrackingEnabled`, `attachmentsEnabled`, `timeTrackingConfiguration`.
  3. `GET /application-properties` (`key`, `permissionLevel`, `keyFilter`), `/application-properties/advanced-settings`, dan `PUT /application-properties/{id}` (contoh `jira.title`, `jira.option.globalsharing`, format tanggal) tersedia; nilai tidak valid untuk tipe properti ditolak 400.
  4. `GET /configuration/timetracking` (provider aktif), `GET /configuration/timetracking/list`, dan `PUT` untuk memilih provider (bawaan atau app) tersedia.
  5. `GET`/`PUT /configuration/timetracking/options` mengelola `workingHoursPerDay`, `workingDaysPerWeek`, `timeFormat` (pretty/days/hours), `defaultUnit`.
  6. `GET`/`PUT /announcementBanner` mengelola `message`, `isDismissible`, `isEnabled`, `visibility` public/private, dan `hashId`.
  7. `GET /label` (paginated) mengembalikan semua label yang dipakai; label dibuat implisit saat work item di-update.
  8. `GET`/`PUT /settings/columns` mengelola kolom default issue navigator.
  9. `GET /locale` mengembalikan locale pengguna saat ini.
  10. Endpoint tulis memerlukan Admin Situs; properti read-only menolak PUT dengan 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 22 (ADM-007, ADM-009) dan 14.

### API-057 Audit Records API

- **Deskripsi:** API untuk membaca catatan audit log (perubahan konfigurasi dan administrasi).
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /auditing/record?offset=&limit=&filter=&from=&to=` dengan `limit` maksimum 1000.
  2. Record memuat `id`, `summary`, `remoteAddress`, `authorAccountId`, `created`, `category`, `eventSource`, `description`, `objectItem` (`id`, `name`, `typeName`, `parentId`, `parentName`), `changedValues` (`fieldName`, `changedFrom`, `changedTo`), dan `associatedItems`.
  3. Kategori mencakup user management, group management, permissions, workflows, fields, projects, schemes, system, dll.
  4. ⚠ Hanya Admin Situs yang dapat membaca (lainnya 403), dan durasi retensi record terdokumentasi serta dapat dikonfigurasi.
  5. API bersifat read-only: tidak ada endpoint untuk mengubah atau menghapus record (PUT/DELETE mengembalikan 405).
  6. Audit log organisasi tersedia terpisah lewat Organization API events (API-011).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat ketersediaan audit log per plan (Free vs Standard) belum pasti. Domain di ADM-014; audit org di USR-045/USR-046.

### API-058 Data Classification API

- **Deskripsi:** API untuk menetapkan classification level pada work item dan default per proyek sesuai kebijakan data organisasi.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / enterprise
- **Aktor & izin:** Org admin (definisi level), Admin Proyek (default proyek), pengguna dengan Edit Issues (klasifikasi work item).
- **Aturan perilaku (acceptance criteria):**
  1. `GET /classification-levels` mengembalikan level (`id`, `name`, `rank`, `color`, `status` PUBLISHED/ARCHIVED, `guideline`, `description`).
  2. `GET`/`PUT`/`DELETE /issue/{key}/classification-level` tersedia; level ARCHIVED tidak dapat dipilih (400).
  3. `GET`/`PUT`/`DELETE /project/{key}/classification-level/default` tersedia.
  4. Level didefinisikan org admin di administrasi organisasi, bukan lewat API ini.
  5. Perubahan klasifikasi tercatat di history work item dan audit log.
  6. Klasifikasi dapat dipakai oleh data security policies (API-012).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat syarat plan belum pasti (pembanding: add-on keamanan). Domain di PERM-064 dan USR-041.

### API-059 Plans API

- **Deskripsi:** API Plans untuk membuat dan mengelola plan, sumber work item, pengaturan penjadwalan, dan tim dalam plan.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / premium
- **Aktor & izin:** Admin plan, Admin Situs; permission Edit/View plan.
- **Aturan perilaku (acceptance criteria):**
  1. `GET /plans/plan` (paginated: `includeTrashed`, `includeArchived`) tersedia.
  2. `POST /plans/plan` menerima `name`, `leadAccountId`, `issueSources` (Board, Project, Filter), `scheduling` (estimation StoryPoints/Days/Hours, field startDate/endDate, inferredDates, dependencies Sequential/Concurrent), `exclusionRules`, `crossProjectReleases`, `customFields`, dan `permissions` (holder Edit/View).
  3. `GET /plans/plan/{id}` dan `PUT` dengan JSON Patch (`application/json-patch+json`) tersedia; patch tidak valid ditolak 400.
  4. `PUT /plans/plan/{id}/archive` dan `/trash` serta `POST /plans/plan/{id}/duplicate` tersedia.
  5. Teams dalam plan: `GET /plans/plan/{id}/team` dan `POST`/`PUT`/`DELETE` plan-only team maupun team platform (`planningStyle` Scrum/Kanban, `capacity`, `sprintLength`, `issueSourceId`, `memberAccountIds`).
  6. Hanya tersedia di plan Premium/Enterprise; plan lain menerima 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di modul 18 (PLAN-027, PLAN-050).

### API-060 Atlassian Teams API

- **Deskripsi:** API tim tingkat platform (dipakai field Team) untuk mengelola tim dan keanggotaannya.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Pengguna, org admin, Integrator/Developer API; mengikuti kebijakan pembuatan tim organisasi.
- **Aturan perilaku (acceptance criteria):**
  1. Teams public API (`/public/teams/v1/org/{orgId}/teams`) menyediakan create, get, update, dan delete tim.
  2. Keanggotaan dikelola lewat tambah/hapus anggota dan list anggota (paginated).
  3. ⚠ Tipe tim dibedakan antara tim open/member-invite dan tim yang disinkron dari IdP/terkelola; tim tersinkron tidak dapat diubah keanggotaannya lewat API.
  4. Field Team pada work item berisi id tim (UUID) dan dapat dicari lewat JQL.
  5. Permission mengikuti kebijakan pembuatan tim di organisasi; pembuatan yang dilarang policy ditolak 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di CFG-019, JQL-032, dan modul 18.

## API data DevOps

### API-061 OAuth Credentials for DevOps Integrations

- **Deskripsi:** Halaman admin untuk membuat OAuth client credentials bagi tool DevOps on-prem/custom yang mengirim data build, deployment, dan development tanpa membuat app.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Admin Situs; tool DevOps sebagai klien.
- **Aturan perilaku (acceptance criteria):**
  1. Settings > Apps > OAuth credentials menyediakan pembuatan credential dengan nama, server base URL, dan logo; nama dan base URL wajib.
  2. ⚠ Admin memilih permission per credential: Builds, Deployments, Development information, Feature flags, dan Remote links.
  3. Pembuatan menghasilkan client ID dan secret untuk grant `client_credentials`; secret hanya ditampilkan sekali.
  4. Credential dapat diedit, dihapus, dan dicabut; pencabutan langsung membatalkan token aktif.
  5. Data yang dikirim dengan credential tampil di panel Development/Releases/Deployments work item.
  6. Hanya Admin Situs yang dapat mengelola halaman ini (lainnya 403).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain integrasi DevOps di [modul 24](24-integrasi-devops.md); halaman admin di modul 22.

### API-062 Development Information API (DevOps)

- **Deskripsi:** API bagi penyedia SCM untuk mengirim repository, commit, branch, dan pull request yang terhubung ke work item lewat key.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / free-standard
- **Aktor & izin:** Integrator/Developer API, tool DevOps; scope data provider atau OAuth credentials DevOps.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /rest/devinfo/0.10/bulk` menerima `repositories` dengan `commits`, `branches`, `pullRequests`, `updateSequenceId`, `properties`, dan `providerMetadata`; respons 202 dengan entitas yang diterima/ditolak.
  2. Asosiasi ke work item memakai `issueKeys` yang diambil dari pesan commit, nama branch, atau judul PR; key yang tidak ada diabaikan dan dilaporkan.
  3. `GET`/`DELETE /repository/{repositoryId}`, `DELETE /bulkByProperties`, dan `DELETE` entitas per tipe dan id tersedia.
  4. `GET /existsByProperties` mengecek keberadaan data berdasarkan properti.
  5. `updateSequenceId` dipakai untuk menolak update usang (out-of-order): entitas dengan sequence lebih kecil dari yang tersimpan diabaikan.
  6. Status PR: `OPEN`, `MERGED`, `DECLINED`, `DRAFT`, beserta reviewer, `commentCount`, dan `url`.
  7. Data muncul di panel Development work item dan dapat dicari lewat JQL `development[...]`.
  8. Autentikasi memakai app dengan modul devops atau OAuth credentials dari Admin Situs (API-061).
  9. ⚠ Batas ukuran payload dan jumlah entitas per request terdokumentasi; pelampauan ditolak 413/400.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di [modul 24](24-integrasi-devops.md); JQL di JQL-035.

### API-063 Builds, Deployments & Other DevOps APIs

- **Deskripsi:** API DevOps untuk build, deployment, feature flag, remote link generik, security, operations, dan DevOps components yang terhubung ke work item.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / free-standard
- **Aktor & izin:** Tool DevOps, Integrator/Developer API; scope data provider atau OAuth credentials DevOps.
- **Aturan perilaku (acceptance criteria):**
  1. Builds: `POST /rest/builds/0.1/bulk` menerima `pipelineId`, `buildNumber`, `updateSequenceNumber`, `displayName`, `url`, `state` (pending/in_progress/successful/failed/cancelled/unknown), `lastUpdated`, `issueKeys`, `testInfo`, `references`.
  2. Builds: `GET`/`DELETE` build per `pipelineId` + `buildNumber` dan `DELETE bulkByProperties` tersedia.
  3. Deployments: `POST /rest/deployments/0.1/bulk` menerima `deploymentSequenceNumber`, `environment` `{id, displayName, type development/testing/staging/production/unmapped}`, `pipeline`, `state`, dan `associations` (issueKeys/issueIdOrKeys/serviceIdOrKeys/commit/repository).
  4. Deployment gating: `GET /pipelines/{p}/environments/{e}/deployments/{d}/gating-status` terintegrasi dengan change management Service Management.
  5. Feature flags: `POST /rest/featureflags/0.1/bulk` menerima `key`, `displayName`, `summary` (status enabled/defaultValue/rollout), dan `details` per environment.
  6. Remote links generik: `/rest/remotelinks/1.0/bulk` untuk alert, incident, dan dokumen terkait work item.
  7. Security: `/rest/security/1.0` untuk linked workspaces dan vulnerabilities (severity, status, identifiers).
  8. Operations: `/rest/operations/1.0` untuk incidents dan post-incident reviews.
  9. DevOps components: `/rest/devopscomponents/1.0/bulk`.
  10. Data tampil di panel Releases/Deployments/Code dan dapat dicari lewat JQL (`deployment`, `build`).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Domain di [modul 24](24-integrasi-devops.md); deployment gating di [modul 32](32-jsm-itsm-insiden-problem-perubahan-layanan.md).

## Webhook

### API-064 System WebHooks Administration

- **Deskripsi:** Halaman admin global (Settings > System > WebHooks) untuk mendaftarkan webhook yang mengirim HTTP POST ke URL eksternal saat event terjadi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs (Administer).
- **Aturan perilaku (acceptance criteria):**
  1. Hanya Admin Situs yang dapat membuat, mengubah, atau menghapus webhook; lainnya tidak melihat menu dan mendapat 403 di API.
  2. Field webhook: Name (wajib), Status (Enabled/Disabled), URL (wajib, https), Description, dan Secret (opsional, untuk signature HMAC; ditampilkan tersamar setelah disimpan).
  3. Variabel URL didukung dan disubstitusi saat kirim: `${issue.id}`, `${issue.key}`, `${project.id}`, `${project.key}`, `${sprint.id}`, `${board.id}`, `${version.id}`, `${mergedVersion.id}`, `${modifiedUser.accountId}`, `${modifiedUser.key}` (legacy).
  4. Filter JQL opsional untuk event issue/comment/worklog; hanya work item yang cocok yang memicu webhook, dan JQL tidak valid ditolak saat simpan.
  5. Event dipilih per kategori lewat checkbox: Issue, Comment, Worklog, Attachment, Issue link, Entity property, Project, Version, User, Jira configuration (perubahan status fitur), Sprint, Board, Filter.
  6. Opsi Exclude body mengirim notifikasi POST tanpa payload JSON.
  7. Daftar webhook menampilkan status dan last updated, dengan aksi edit, disable, dan delete.
  8. Webhook tanpa event terpilih dapat dipicu dari post function "Trigger a webhook" di workflow company-managed.
  9. ⚠ REST admin `GET`/`POST`/`PUT`/`DELETE /rest/webhooks/1.0/webhook` (auth admin) tersedia sebagai lapisan kompatibel.
  10. ⚠ Jumlah webhook per site dibatasi (default 100, dapat dikonfigurasi); pembuatan di atas batas ditolak dengan pesan batas.
  11. Event dari proyek team-managed juga terkirim dan filter JQL berlaku sama.
- **Team-managed vs company-managed:** Lihat butir 11; post function "Trigger a webhook" hanya di company-managed.
- **Catatan & ketergantungan:** Halaman admin dirujuk modul 22; pengiriman di API-066; skema data di [modul 28](28-model-data-inti.md). Migrasi repo: feature baru `internal/features/webhook` dengan tabel `webhooks` dan `webhook_deliveries`.

### API-065 Webhook Events Catalogue

- **Deskripsi:** Daftar lengkap event webhook beserta identifier `webhookEvent` dan sub-tipe `issue_event_type_name`.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Admin Situs, Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Issue: `jira:issue_created`, `jira:issue_updated`, `jira:issue_deleted`, dengan `issue_event_type_name` antara lain `issue_created`, `issue_updated`, `issue_assigned`, `issue_generic`, `issue_resolved`, `issue_closed`, `issue_reopened`, `issue_moved`, `issue_commented`, `issue_comment_edited`, `issue_comment_deleted`, `issue_work_logged`, `issue_worklog_updated`, `issue_worklog_deleted`, `issue_work_started`, `issue_work_stopped`.
  2. Comment: `comment_created`, `comment_updated`, `comment_deleted`.
  3. Worklog: `worklog_created`, `worklog_updated`, `worklog_deleted`.
  4. Attachment: `attachment_created`, `attachment_deleted`.
  5. Issue link: `issuelink_created`, `issuelink_deleted`.
  6. Entity property: `issue_property_set`, `issue_property_deleted`.
  7. Project: `project_created`, `project_updated`, `project_deleted`, `project_soft_deleted`, `project_restored_deleted`, `project_archived`, `project_restored_archived`.
  8. Version: `jira:version_created`, `jira:version_updated`, `jira:version_released`, `jira:version_unreleased`, `jira:version_moved`, `jira:version_merged`, `jira:version_deleted`.
  9. User: `user_created`, `user_updated`, `user_deleted`.
  10. Jira configuration: `option_voting_changed`, `option_watching_changed`, `option_unassigned_issues_changed`, `option_subtasks_changed`, `option_attachments_changed`, `option_issuelinks_changed`, `option_timetracking_changed`.
  11. Sprint: `sprint_created`, `sprint_updated`, `sprint_deleted`, `sprint_started`, `sprint_closed`.
  12. Board: `board_created`, `board_updated`, `board_deleted`, `board_configuration_changed`.
  13. Filter: `filter_created`, `filter_updated`, `filter_deleted`.
  14. ⚠ Event issue type dan component (`component_created`/`updated`/`deleted`, `issuetype_*`) tersedia di katalog.
  15. Event Service Management dan Product Discovery ditambahkan oleh modul produk masing-masing ke katalog yang sama.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Katalog event domain di [modul 28](28-model-data-inti.md) menjadi sumber; event JSM di [modul 35](35-jsm-ai-otomasi-api-admin-csm.md), event JPD di [modul 37](37-jpd-insights-delivery-peran-integrasi.md). Nama identifier dipertahankan untuk kompatibilitas penerima yang ada.

### API-066 Webhook Payload, Delivery & Security

- **Alias inventaris:** Webhook Delivery Reliability
- **Deskripsi:** Format payload, header, semantik pengiriman, retry, keandalan, observabilitas, dan verifikasi tanda tangan webhook.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / free-standard
- **Aktor & izin:** Integrator/Developer API (penerima), Admin Situs (memantau dan redeliver), Engineering.
- **Aturan perilaku (acceptance criteria):**
  1. Payload JSON memuat `timestamp`, `webhookEvent`, `issue_event_type_name`, `user`, `issue` (representasi REST), `changelog` (`id`, `items` from/to), `comment`, `worklog`, `project`, `sprint`, `board`, `version`, dan `matchedWebhookIds` sesuai jenis event.
  2. ⚠ Representasi work item di payload mengikuti bentuk REST v2 (description berupa string); varian ADF tersedia lewat opsi per webhook.
  3. Setiap pengiriman membawa header `X-Atlassian-Webhook-Identifier` (id unik per event, tetap sama saat retry, untuk dedup) dan `X-Atlassian-Webhook-Retry` (nomor retry).
  4. Bila secret dikonfigurasi, header `X-Hub-Signature: sha256=<hex>` berisi HMAC SHA-256 dari body mentah dengan secret.
  5. Retry otomatis s.d. 5 kali dengan backoff untuk respons 408, 409, 425, 429, 5xx, atau timeout/kegagalan koneksi; respons 4xx lain tidak di-retry.
  6. ⚠ Respons 2xx dianggap sukses; batas waktu koneksi/respons sekitar 10-20 detik.
  7. Urutan tidak dijamin dan pengiriman bersifat at-least-once, sehingga dokumentasi mewajibkan penerima idempoten.
  8. Payload disaring sesuai permission pemilik webhook (admin) atau app/pengguna yang mendaftarkan.
  9. Pengiriman di-throttle bila site menghasilkan event berlebih, dengan batas konkurensi; pada beban ekstrem event dapat ditunda atau di-drop dan pengabaian tersebut tercatat di log.
  10. Hanya URL publik yang dikirimi; alamat privat, loopback, link-local, dan localhost ditolak saat simpan dan saat kirim (SSRF guard, termasuk setelah resolusi DNS).
  11. Mode Exclude body mengirim POST tanpa body.
  12. Skema `webhooks` menyimpan name, url, events, JQL filter, secret (terenkripsi), enabled, exclude body, created by; `webhook_deliveries` menyimpan payload, status code, latency, attempt, dan next retry.
  13. ⚠ Pengiriman at-least-once memakai retry exponential dengan jumlah percobaan dan interval yang terdokumentasi.
  14. Signature HMAC-SHA256 di `X-Hub-Signature` dihitung identik dengan butir 4 dan diuji dengan vektor uji di dokumentasi.
  15. Timeout per pengiriman ≤ 10 detik, circuit breaker per endpoint, dan auto-disable setelah kegagalan beruntun (ambang dapat dikonfigurasi) disertai notifikasi ke admin.
  16. Filter JQL dievaluasi saat event terjadi (bukan saat pengiriman), sehingga retry memakai hasil evaluasi awal.
  17. Payload menghormati izin pemilik webhook/app pada saat event dibuat.
  18. Log pengiriman disimpan dengan retensi terbatas (default 30 hari) dan tombol redeliver yang mengirim ulang payload asli dengan identifier sama.
  19. Ordering per work item best-effort dengan `timestamp` dan sequence; header identitas unik dipakai untuk dedupe.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Digabung dari domain api-webhooks dan nfr-architecture. Antrean pengiriman memakai job queue [modul 27](27-nfr-arsitektur-platform.md) (outbox pattern di transaksi domain agar event tidak hilang). Penamaan header kompatibel dipertahankan; alias header netral dapat ditambahkan.

### API-067 Dynamic Webhooks API (Apps)

- **Deskripsi:** API bagi app OAuth 2.0 dan app legacy untuk mendaftarkan webhook secara terprogram, dengan filter JQL terbatas dan masa berlaku.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pengembang app); scope `manage:jira-webhook`.
- **Aturan perilaku (acceptance criteria):**
  1. `POST /webhook` menerima `url` dan `webhooks` `[{events, jqlFilter, fieldIdsFilter, issuePropertyKeysFilter}]`; respons memuat `webhookRegistrationResult` per item (`createdWebhookId` atau `errors`).
  2. `GET /webhook` (paginated) mengembalikan webhook milik app pemanggil saja.
  3. `DELETE /webhook` dengan body `webhookIds` menghapus webhook milik app.
  4. `PUT /webhook/refresh` memperpanjang masa berlaku; webhook OAuth 2.0 kedaluwarsa 30 hari setelah dibuat/di-refresh dan berhenti dikirim setelahnya.
  5. `GET /webhook/failed` mengembalikan pengiriman gagal 72 jam terakhir untuk app.
  6. ⚠ `jqlFilter` dibatasi subset JQL: field `project`, `issueKey`, `issuetype`, `status`, `assignee`, `reporter`, `issue.property` dengan operator `=`, `!=`, `IN`, `NOT IN`; di luar subset ditolak di `errors`.
  7. `fieldIdsFilter` membuat `jira:issue_updated` hanya terpicu bila field tertentu berubah.
  8. `issuePropertyKeysFilter` berlaku untuk event `issue_property_set`/`issue_property_deleted`.
  9. Event yang didukung terbatas pada issue, comment, worklog, attachment, issuelink, dan issue property; event lain ditolak.
  10. ⚠ Jumlah webhook dibatasi: OAuth 5 per pengguna per app per tenant dan app legacy 100 per app per tenant.
  11. Memerlukan scope `manage:jira-webhook`; data dikirim sesuai permission pengguna yang mendaftarkan.
  12. ⚠ Webhook app OAuth 2.0 membawa header `Authorization` berisi JWT yang ditandatangani client secret app; app legacy memakai JWT shared secret.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Memakai mesin pengiriman API-066. Bergantung pada API-009/API-010.

### API-068 Automation Web Requests & Incoming Webhooks

- **Deskripsi:** Webhook lewat otomasi: action "Send web request" untuk mengirim keluar dan trigger "Incoming webhook" untuk menerima.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / free-standard
- **Aktor & izin:** Admin Proyek/Project Lead (rule proyek), Admin Situs (rule global).
- **Aturan perilaku (acceptance criteria):**
  1. Action Send web request mengonfigurasi URL, method (GET/POST/PUT/PATCH/DELETE), headers (dapat ditandai hidden/secret), dan body (Empty, Issue data format otomasi, Issue data format REST, Custom data dengan smart values).
  2. Opsi "Delay execution of subsequent rule actions until we've received a response" membuat `{{webResponse.status}}`, `{{webResponse.body}}`, dan `{{webResponse.headers}}` dapat diakses action berikutnya.
  3. Trigger Incoming webhook menghasilkan URL unik per rule dengan secret yang diverifikasi lewat header `X-Automation-Webhook-Token`; token salah ditolak 401.
  4. Incoming webhook dapat memproses work item dari payload (`issues: [keys]`), dari JQL, atau tanpa work item; data payload diakses lewat `{{webhookData}}`.
  5. Batas timeout respons, ukuran body, dan kuota eksekusi rule per bulan mengikuti plan.
  6. Audit log rule menampilkan request/response untuk debugging (header secret tersamar).
  7. URL harus domain publik (SSRF guard sama dengan API-066); rule di team-managed ber-scope proyek.
- **Team-managed vs company-managed:** Rule team-managed ber-scope proyek.
- **Catatan & ketergantungan:** Detail rule di modul 20 (AUTO-025, AUTO-054, AUTO-056, AUTO-015). Dicatat di sini sebagai alternatif webhook.

## Pengalaman developer & ekosistem

### API-069 API Documentation, OpenAPI Spec & SDKs

- **Deskripsi:** Dokumentasi referensi API publik, spesifikasi OpenAPI yang dapat diunduh untuk generate client, dan framework/SDK resmi.
- **Prioritas / Fase / Tier:** P1 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API; dokumentasi dapat dibaca tanpa login (dapat dibatasi admin).
- **Aturan perilaku (acceptance criteria):**
  1. Referensi tersedia untuk REST platform v3, v2, Agile, DevOps, dan Service Management, dikelompokkan per area.
  2. Spesifikasi OpenAPI 3 dapat diunduh (platform dan Agile terpisah) dari URL stabil.
  3. Setiap operasi mencantumkan permission yang dibutuhkan, scope classic dan granular, scope app, dampak data security policy, contoh request (curl, Node.js, Java, Python, PHP, SDK app), dan contoh respons.
  4. Operasi diberi label Experimental, Deprecated, atau Beta yang tampil di dokumentasi dan di spesifikasi (`x-` extension).
  5. ⚠ Koleksi Postman resmi dihasilkan dari OpenAPI dan dapat diunduh.
  6. SDK/framework resmi: SDK app dan CLI untuk model app; tidak ada SDK REST untuk semua bahasa, dan dokumentasi menjelaskan generate client dengan openapi-generator.
  7. Resource dikelompokkan per tag (Issues, Issue search, Projects, Users, Workflows, dll.) yang sama dengan tag swag.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Clone sudah memakai swag + Scalar (`internal/shared/delivery/http/apidocs`); setiap handler perlu anotasi tag, scope, dan permission (`@Security BearerAuth`). Library komunitas (jira.js, go-jira) hanya ekosistem. Berkaitan dengan API-002.

### API-070 API Deprecation Policy & Changelog

- **Deskripsi:** Kebijakan deprecation dan changelog publik yang mengumumkan perubahan, endpoint baru, dan penghapusan API.
- **Prioritas / Fase / Tier:** P2 / Fase 3 / platform
- **Aktor & izin:** Integrator/Developer API, tim platform.
- **Aturan perilaku (acceptance criteria):**
  1. Changelog API platform dan Agile dipublikasikan dengan tanggal dan tipe entri (added, deprecated, removed), dapat dilanggan lewat RSS/feed.
  2. Masa notice deprecation minimal 6 bulan untuk API stabil; endpoint experimental dapat berubah tanpa jaminan masa notice.
  3. Daftar deprecation besar yang tidak diadopsi clone sejak awal didokumentasikan: username/userKey, `/search` lama ke `/search/jql`, createmeta lama, Epic Link/Parent Link ke `parent`, legacy workflow API, `/expression/eval`, dan model app legacy.
  4. ⚠ Header opt-in/opt-out masa transisi tersedia untuk sebagian perubahan, sehingga klien dapat menguji perilaku baru lebih awal.
  5. Perubahan besar didahului RFC publik di forum komunitas developer dengan periode umpan balik.
  6. Status endpoint di dokumentasi: GA, Beta, Experimental, Deprecated.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Header `Deprecation`/`Sunset` di API-002.

### API-071 Developer Console & App Management

- **Deskripsi:** Konsol developer untuk membuat app OAuth 2.0/app native, mengatur scope, kredensial, distribusi, dan kewajiban data pribadi, plus site developer untuk menguji app dan integrasi.
- **Prioritas / Fase / Tier:** P2 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (pemilik app).
- **Aturan perilaku (acceptance criteria):**
  1. Developer dapat membuat app bertipe OAuth 2.0 integration atau app native.
  2. Halaman Permissions menambahkan API (platform, Agile, Service Management, user identity) lalu memilih scope classic/granular.
  3. Halaman Authorization mengatur satu atau lebih callback URL dan menampilkan contoh authorization URL.
  4. Halaman Settings menampilkan client ID, rotate client secret (secret lama invalid setelah rotasi), nama, avatar, dan deskripsi.
  5. Halaman Distribution mengatur sharing on/off, vendor name, privacy policy, terms, customer support, dan personal data declaration (wajib bila menyimpan data pribadi; tanpa itu sharing tidak dapat diaktifkan).
  6. App native memiliki environment (development, staging, production), deployment, logs, metrics (invocation, error), installations, dan distribution link.
  7. Personal data reporting API `POST /app/report-accounts/` menerima `accountId` + `updatedAt` (siklus maksimal 7 hari) dan merespons `closed`/`updated` untuk menandai data pengguna yang harus dihapus atau diperbarui.
  8. ⚠ Konsol menampilkan monitoring penggunaan dan rate limit per app.
  9. ⚠ Developer site gratis tersedia untuk membangun dan menguji app/integrasi REST dengan batas pengguna kecil (sekitar 5) dan akses fitur edisi lebih tinggi untuk pengujian.
  10. Instalasi app dalam development mode di developer site didukung (padanan install/tunnel untuk app native dan sideload untuk app legacy).
  11. Developer site tidak untuk penggunaan produksi dan berbeda dari Sandbox Premium/Enterprise yang ditautkan ke site produksi (Sandbox dicatat di modul 22).
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** ⚠ Inventaris mencatat batas pengguna dan fitur developer site belum pasti; padanan clone: tenant uji atau mode developer pada instance self-hosted. GDPR API terkait USR-044.

### API-072 App-specific Platform APIs

- **Deskripsi:** Kumpulan endpoint khusus app: UI modifications, app properties, dynamic modules, data policy, migrasi, dan service registry.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API (app terpasang); hanya app pemilik yang dapat mengakses data app-nya.
- **Aturan perilaku (acceptance criteria):**
  1. UI modifications: `GET`/`POST`/`PUT`/`DELETE /uiModifications` dengan `contexts` (`projectId`, `issueTypeId`, `viewType` GIC/IssueView/IssueTransition) dan `data` string.
  2. App properties: `GET`/`PUT`/`DELETE /addons/{addonKey}/properties/{key}` dan `/forge/app/properties/{key}`; app lain ditolak 403.
  3. Dynamic modules (app legacy): `GET`/`POST`/`DELETE /atlassian-connect/1/app/module/dynamic`.
  4. App data policies: `GET /data-policy` dan `/data-policy/project` (`anyContentBlocked` per proyek) mencerminkan data security policy (API-012).
  5. App migration (self-hosted ke cloud): `/atlassian-connect/1/migration/field`, `properties`, `workflow/rule/search`.
  6. Service registry: `GET /atlassian-connect/1/service-registry`.
  7. Jira expressions dan fungsi JQL untuk app memakai API-031 dan API-030.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Hanya relevan bila model ekstensi [modul 25](25-integrasi-impor-ekspor-ekstensi.md) dibangun; path legacy disediakan hanya bila kompatibilitas app legacy diputuskan.

### API-073 Atlassian CLI, GraphQL API & Remote MCP Server

- **Deskripsi:** Antarmuka resmi tambahan di atas API: CLI, GraphQL gateway, dan Remote MCP server untuk tool AI.
- **Prioritas / Fase / Tier:** P3 / Fase 4 / platform
- **Aktor & izin:** Integrator/Developer API, pengguna, org admin (mengaktifkan/menonaktifkan).
- **Aturan perilaku (acceptance criteria):**
  1. CLI menyediakan `auth login` (API token atau OAuth via browser), `status`, dan `switch` site.
  2. CLI work item: `workitem create/edit/view/search (JQL)/transition/assign/comment/clone/delete` dengan output JSON/CSV.
  3. ⚠ CLI mencakup perintah untuk project, board, sprint, filter, dan admin (`admin user activate/deactivate`).
  4. GraphQL API (`/graphql`) dengan explorer tersedia; bagian skema yang belum stabil ditandai beta.
  5. ⚠ Remote MCP server memakai OAuth 2.1 dan menyediakan tool search, read, create, dan update work item untuk klien AI.
  6. Semua akses tunduk pada permission pengguna dan kebijakan admin; admin dapat menonaktifkan CLI/GraphQL/MCP per organisasi, setelah itu panggilan ditolak 403.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** CLI clone dapat dibangun dalam Go memakai SDK hasil OpenAPI; MCP server dapat memakai tool yang memetakan langsung ke endpoint API-016/API-029. Layanan AI lintas produk vendor (Rovo) di luar cakupan.

### API-074 Jira Data Center REST API & Personal Access Tokens (Reference)

- **Deskripsi:** Perbedaan API versi Data Center sebagai referensi kompatibilitas; bukan fitur cloud dan tidak dibangun sebagai fitur tersendiri.
- **Prioritas / Fase / Tier:** P3 / Fase — / dc-only-or-deprecated
- **Aktor & izin:** Integrator/Developer API.
- **Aturan perilaku (acceptance criteria):**
  1. Rujukan: base `/rest/api/2` dan `/rest/api/latest` tanpa v3/ADF.
  2. Rujukan: Personal Access Token (`Authorization: Bearer`) dengan expiry dan revoke oleh admin.
  3. Rujukan: identifikasi pengguna dengan username/key, bukan accountId.
  4. Rujukan: search `/rest/api/2/search` dengan `startAt`/`total` masih didukung.
  5. Rujukan: cookie session `/rest/auth/1/session` masih didukung.
  6. Rujukan: webhook dengan konfigurasi serupa plus REST `/rest/webhooks/1.0`.
  7. Rujukan: rate limiting dikonfigurasi admin (token bucket per pengguna).
  8. Rujukan: OAuth 2.0 incoming/outgoing lewat Application Links.
- **Team-managed vs company-managed:** Tidak ada perbedaan.
- **Catatan & ketergantungan:** Untuk instance self-hosted clone, rate limit yang dapat dikonfigurasi admin diserap ke API-006 dan Application Links ke API-014; PAT Bearer dapat menjadi opsi bila diputuskan (lihat Pertanyaan terbuka). Lihat Di luar cakupan.

## Pertanyaan terbuka

- Apakah lapisan kompatibel `/rest/api/3` dan `/rest/agile/1.0` dibangun penuh, sebagian (subset populer: issue, search, comment, transition, project, user, board, sprint), atau ditunda, mengingat route native `/api/v1` sudah menjadi konvensi repo?
- Apakah `/rest/api/2` (wiki markup) disediakan sebagai alias, atau clone hanya mendukung ADF (v3) untuk menghindari biaya konversi lossy?
- Apakah format error native `/api/v1` (`{error, status, message, data, errors}`) dipertahankan berbeda dari bentuk `{errorMessages, errors}` di lapisan kompatibel, atau diseragamkan?
- Apakah selain Basic auth dengan API token, clone juga menyediakan Personal Access Token bergaya Bearer (pola Data Center) karena lebih sederhana bagi instance self-hosted?
- Model app mana yang diadopsi: hanya OAuth 2.0 3LO + webhook dinamis, atau juga runtime app ala Forge (event trigger, web trigger, UI modifications) yang memerlukan hosting fungsi?
- Apakah nama header dan scope bermerek vendor (`X-Atlassian-*`, `*:jira-*`) dipertahankan apa adanya demi kompatibilitas, atau dialiaskan ke nama netral dengan nama lama sebagai alias?
- Berapa angka default rate limit (burst per detik, quota per jam, biaya per endpoint) dan batas per work item yang dipakai clone?
- Apakah bahasa ekspresi (API-031) dibangun kompatibel dengan Jira expressions atau memakai engine standar (CEL/expr) dengan sintaks berbeda?
- Berapa retensi log pengiriman webhook, task async, dan audit record default, serta apakah dapat dikonfigurasi per plan?
- Apakah GraphQL API dan Remote MCP server masuk roadmap produk ini, atau cukup REST + CLI?

## Di luar cakupan & ekosistem

- **App Marketplace pihak ketiga** dan library komunitas (jira.js, go-jira, klien Python, dll.) — hanya ekosistem; produk ini menyediakan OpenAPI agar klien dapat di-generate.
- **Platform Forge sebagai layanan vendor** (hosting fungsi, Forge CLI resmi, @forge/api) — API-013 hanya menjadi acuan model `asApp`/`asUser` dan event subscription; runtime app clone diputuskan di [modul 25](25-integrasi-impor-ekspor-ekstensi.md).
- **Atlassian Connect (deprecated)** — API-015 hanya rujukan; JWT shared secret, lifecycle callback, dan static webhook descriptor tidak dibangun sebagai fitur.
- **Jira Data Center REST API & Personal Access Tokens** — API-074 hanya rujukan kompatibilitas (dc-only-or-deprecated).
- **Endpoint privat `/rest/greenhopper/1.0`** — tidak menjadi kontrak publik (API-032).
- **Layanan Atlassian lintas produk** — identitas `id.atlassian.com`, gateway `api.atlassian.com`, Atlassian Administration, Atlassian Guard, Compass, Confluence, Bitbucket, Rovo/Atlassian Intelligence, Atlassian GraphQL Gateway, dan Rovo MCP sebagai layanan vendor; produk ini menyediakan padanan internalnya sendiri sebagaimana dijelaskan di fitur terkait.
- **Koleksi Postman resmi vendor, Atlassian Developer Community forum, dan developer.atlassian.com** — produk ini menerbitkan dokumentasi, changelog, dan forum RFC-nya sendiri (API-069, API-070).
- **Layanan Media Platform vendor** — lampiran disimpan di object storage milik clone (API-021).
