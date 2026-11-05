import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { SESSION_STORAGE_KEY } from '../application/auth.constants.js';
import type { Session } from '../domain/auth.types.js';
import { clearSession, readSession, writeSession } from './session.storage.js';

function fakeStorage(): Storage {
	const data = new Map<string, string>();
	return {
		get length() {
			return data.size;
		},
		clear: () => data.clear(),
		getItem: (k) => data.get(k) ?? null,
		key: (i) => [...data.keys()][i] ?? null,
		removeItem: (k) => void data.delete(k),
		setItem: (k, v) => void data.set(k, v)
	};
}

const make = (expiresAt: number): Session => ({
	token: 'jwt',
	expiresAt,
	user: { id: '1', email: 'a@b.c', username: 'a', role: 'user' }
});

describe('session.storage', () => {
	beforeEach(() => vi.stubGlobal('localStorage', fakeStorage()));
	afterEach(() => vi.unstubAllGlobals());

	it('round-trips a valid session', () => {
		const s = make(Date.now() + 60_000);
		writeSession(s);
		expect(readSession()).toEqual(s);
	});

	it('drops and removes an expired session', () => {
		writeSession(make(Date.now() - 1));
		expect(readSession()).toBeNull();
		expect(localStorage.getItem(SESSION_STORAGE_KEY)).toBeNull();
	});

	it('ignores malformed data', () => {
		localStorage.setItem(SESSION_STORAGE_KEY, '{not json');
		expect(readSession()).toBeNull();
		localStorage.setItem(SESSION_STORAGE_KEY, JSON.stringify({ token: 1 }));
		expect(readSession()).toBeNull();
	});

	it('clears the session', () => {
		writeSession(make(Date.now() + 60_000));
		clearSession();
		expect(readSession()).toBeNull();
	});

	it('survives storage that throws', () => {
		vi.stubGlobal('localStorage', {
			getItem: () => {
				throw new Error('blocked');
			},
			setItem: () => {
				throw new Error('blocked');
			},
			removeItem: () => {
				throw new Error('blocked');
			}
		});
		expect(readSession()).toBeNull();
		expect(() => writeSession(make(Date.now() + 1000))).not.toThrow();
		expect(() => clearSession()).not.toThrow();
	});
});
