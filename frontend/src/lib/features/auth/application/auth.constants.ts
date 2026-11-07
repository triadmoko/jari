import { PRODUCT_SLUG } from '#lib/brand.js';
import { t } from '#lib/i18n/index.js';

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
	get sessionExpired() {
		return t('auth.sessionExpired');
	},
	get registered() {
		return t('auth.registered');
	},
	required: (field: string) => t('validation.required', { field }),
	get invalidEmail() {
		return t('validation.invalidEmail');
	},
	get usernameTooShort() {
		return t('validation.usernameTooShort', { min: AUTH_LIMITS.usernameMin });
	},
	get usernameTooLong() {
		return t('validation.usernameTooLong', { max: AUTH_LIMITS.usernameMax });
	},
	get passwordTooShort() {
		return t('validation.passwordTooShort', { min: AUTH_LIMITS.passwordMin });
	}
};
