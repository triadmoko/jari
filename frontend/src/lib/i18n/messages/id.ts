import { GLOSSARY, capitalize } from '../glossary.js';

const term = GLOSSARY.id;

/**
 * Indonesian catalog, the reference locale: its keys define `MessageKey`, and `en.ts` must
 * provide every one of them (checked by the compiler). Placeholders use `{name}`.
 *
 * Wording of product concepts comes from `glossary.ts`, never typed inline.
 */
export const id = {
	'brand.description': 'Platform manajemen kerja dan proyek open-source.',

	'action.login': 'Masuk',
	'action.register': 'Daftar',
	'action.logout': 'Keluar',
	'action.save': 'Simpan',

	'nav.profile': 'Profil',
	'nav.language': 'Bahasa',

	'field.name': 'Nama',
	'field.username': 'Username',
	'field.email': 'Email',
	'field.password': 'Kata sandi',

	'validation.required': '{field} wajib diisi.',
	'validation.invalidEmail': 'Format email tidak valid.',
	'validation.usernameTooShort': 'Username minimal {min} karakter.',
	'validation.usernameTooLong': 'Username maksimal {max} karakter.',
	'validation.passwordTooShort': 'Kata sandi minimal {min} karakter.',

	'api.validation': 'Periksa kembali isian Anda.',
	'api.unauthorized': 'Sesi Anda tidak valid. Silakan masuk kembali.',
	'api.invalidCredentials': 'Email atau kata sandi salah.',
	'api.forbidden': 'Anda tidak memiliki akses untuk tindakan ini.',
	'api.network': 'Tidak dapat terhubung ke server. Coba lagi beberapa saat.',
	'api.unknown': 'Terjadi kesalahan. Coba lagi beberapa saat.',

	'page.home': 'Beranda',
	'home.greeting': 'Halo, {name}',
	'home.spacesPlaceholder': `${capitalize(term.space.many)} Anda akan tampil di sini.`,
	'home.viewProfile': 'Lihat profil',
	'home.welcome': 'Selamat datang di {product}',
	'home.welcomeHint': 'Masuk atau daftar untuk mulai bekerja.',

	'auth.login.title': 'Masuk',
	'auth.login.description': 'Masuk dengan email dan kata sandi Anda.',
	'auth.login.noAccount': 'Belum punya akun?',
	'auth.register.title': 'Daftar',
	'auth.register.description': 'Buat akun baru untuk mulai memakai {product}.',
	'auth.register.haveAccount': 'Sudah punya akun?',
	'auth.registered': 'Akun berhasil dibuat. Silakan masuk.',
	'auth.sessionExpired': 'Sesi Anda telah berakhir. Silakan masuk kembali.',

	'profile.title': 'Profil',
	'profile.description': 'Lihat dan ubah nama tampilan Anda.',
	'profile.updated': 'Profil berhasil diperbarui.',

	'error.notFound': 'Halaman tidak ditemukan.',
	'error.generic': 'Terjadi kesalahan.',
	'error.backHome': 'Kembali ke beranda'
} as const;

export type MessageKey = keyof typeof id;
