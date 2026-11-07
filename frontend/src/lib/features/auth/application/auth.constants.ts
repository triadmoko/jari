import { PRODUCT_SLUG } from '#lib/brand.js';

export const AUTH_ROUTES = {
	login: '/auth/login',
	register: '/auth/register',
	home: '/',
	profile: '/profile'
} as const;

export const PROTECTED_PREFIXES = ['/admin', AUTH_ROUTES.profile] as const;
export const GUEST_ONLY_PATHS = [AUTH_ROUTES.login, AUTH_ROUTES.register] as const;

/** Sliding window, same semantics as the old 12h cookie. */
export const SESSION_TTL_MS = 12 * 60 * 60 * 1000;

export const SESSION_STORAGE_KEY = `${PRODUCT_SLUG}_session`;

/** Mirrors the backend rules in `dto_auth.RegisterRequest.Validate`. */
export const AUTH_LIMITS = {
	usernameMin: 3,
	usernameMax: 20,
	passwordMin: 6
} as const;

export const AUTH_MESSAGES = {
	sessionExpired: 'Sesi Anda telah berakhir. Silakan masuk kembali.',
	registered: 'Akun berhasil dibuat. Silakan masuk.',
	required: (field: string) => `${field} wajib diisi.`,
	invalidEmail: 'Format email tidak valid.',
	usernameTooShort: `Username minimal ${AUTH_LIMITS.usernameMin} karakter.`,
	usernameTooLong: `Username maksimal ${AUTH_LIMITS.usernameMax} karakter.`,
	passwordTooShort: `Kata sandi minimal ${AUTH_LIMITS.passwordMin} karakter.`
} as const;
