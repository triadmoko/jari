export const AUTH_ROUTES = {
	login: '/auth/login',
	home: '/',
	admin: '/admin'
} as const;

export const PROTECTED_PREFIXES = ['/admin', '/profile'] as const;
export const GUEST_ONLY_PATHS = [AUTH_ROUTES.login] as const;

/** Sliding window, same semantics as the old 12h cookie. */
export const SESSION_TTL_MS = 12 * 60 * 60 * 1000;

export const SESSION_STORAGE_KEY = 'jari_session';

export const AUTH_MESSAGES = {
	sessionExpired: 'Sesi Anda telah berakhir. Silakan masuk kembali.'
} as const;
