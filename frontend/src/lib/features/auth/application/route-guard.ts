import type { Session } from '../domain/auth.types.js';
import { AUTH_ROUTES, GUEST_ONLY_PATHS, PROTECTED_PREFIXES } from './auth.constants.js';

function matchesPrefix(pathname: string, prefix: string): boolean {
	return pathname === prefix || pathname.startsWith(prefix + '/');
}

/** Only same-origin absolute paths are allowed as a post-login target (no open redirect). */
export function safeRedirectTarget(value: string | null | undefined): string | null {
	if (!value || !value.startsWith('/') || value.startsWith('//') || value.startsWith('/\\')) {
		return null;
	}
	return value;
}

/** Where to go after a successful login: a safe redirectTo, otherwise home. */
export function postLoginTarget(redirectTo?: string | null): string {
	return safeRedirectTarget(redirectTo) ?? AUTH_ROUTES.home;
}

/**
 * Pure route guard. Returns the URL to redirect to, or null when the
 * navigation is allowed.
 */
export function resolveGuard(
	pathname: string,
	search: string,
	session: Session | null
): string | null {
	if (!session && PROTECTED_PREFIXES.some((p) => matchesPrefix(pathname, p))) {
		const redirectTo = encodeURIComponent(pathname + search);
		return `${AUTH_ROUTES.login}?redirectTo=${redirectTo}`;
	}

	if (session && GUEST_ONLY_PATHS.some((p) => matchesPrefix(pathname, p))) {
		return AUTH_ROUTES.home;
	}

	return null;
}
