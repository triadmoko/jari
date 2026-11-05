import type { Session, SessionUser } from '../domain/auth.types.js';
import { clearSession, readSession, writeSession } from '../infrastructure/session.storage.js';
import { SESSION_TTL_MS } from './auth.constants.js';

class SessionStore {
	#session = $state<Session | null>(readSession());

	get current(): Session | null {
		return this.#session;
	}
	get token(): string | null {
		return this.#session?.token ?? null;
	}
	get user(): SessionUser | null {
		return this.#session?.user ?? null;
	}
	get isAuthenticated(): boolean {
		return this.#session !== null;
	}

	login(token: string, user: SessionUser): Session {
		const session: Session = { token, user, expiresAt: Date.now() + SESSION_TTL_MS };
		writeSession(session);
		this.#session = session;
		return session;
	}

	/** Slide the expiry window forward (called on authenticated API activity). */
	touch(): void {
		if (!this.#session) return;
		const session = { ...this.#session, expiresAt: Date.now() + SESSION_TTL_MS };
		writeSession(session);
		this.#session = session;
	}

	/** Re-read storage (e.g. another tab logged in/out). */
	sync(): void {
		this.#session = readSession();
	}

	logout(): void {
		clearSession();
		this.#session = null;
	}
}

export const session = new SessionStore();
