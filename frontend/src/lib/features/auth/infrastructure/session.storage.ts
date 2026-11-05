import type { Session } from '../domain/auth.types.js';
import { SESSION_STORAGE_KEY } from '../application/auth.constants.js';

function storage(): Storage | null {
	try {
		return typeof localStorage === 'undefined' ? null : localStorage;
	} catch {
		// Access can throw (blocked site data, private mode)
		return null;
	}
}

function isSession(value: unknown): value is Session {
	if (typeof value !== 'object' || value === null) return false;
	const v = value as Partial<Session>;
	return (
		typeof v.token === 'string' &&
		v.token !== '' &&
		typeof v.expiresAt === 'number' &&
		typeof v.user === 'object' &&
		v.user !== null
	);
}

/** Returns the stored session, or null when missing, malformed or expired. */
export function readSession(now: number = Date.now()): Session | null {
	const store = storage();
	if (!store) return null;

	try {
		const raw = store.getItem(SESSION_STORAGE_KEY);
		if (!raw) return null;

		const parsed: unknown = JSON.parse(raw);
		if (!isSession(parsed) || parsed.expiresAt <= now) {
			store.removeItem(SESSION_STORAGE_KEY);
			return null;
		}
		return parsed;
	} catch {
		return null;
	}
}

export function writeSession(session: Session): void {
	try {
		storage()?.setItem(SESSION_STORAGE_KEY, JSON.stringify(session));
	} catch {
		// Quota / blocked storage: session just won't persist
	}
}

export function clearSession(): void {
	try {
		storage()?.removeItem(SESSION_STORAGE_KEY);
	} catch {
		// ignore
	}
}
