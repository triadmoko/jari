import { AxiosError, type InternalAxiosRequestConfig } from 'axios';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

const goto = vi.hoisted(() => vi.fn());
vi.mock('$app/navigation', () => ({ goto }));

import { AUTH_ROUTES } from '#lib/features/auth/application/auth.constants.js';
import { session } from '#lib/features/auth/application/session.svelte.js';
import { api } from './client.js';

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

/** Replaces the network: every request is recorded and answered with `status`. */
function stubServer(status: number) {
	const seen: InternalAxiosRequestConfig[] = [];
	api.instance.defaults.adapter = async (config) => {
		seen.push(config);
		const response = { data: {}, status, statusText: '', headers: {}, config };
		if (status >= 400) {
			throw new AxiosError('failed', AxiosError.ERR_BAD_REQUEST, config, null, response);
		}
		return response;
	};
	return seen;
}

const user = { id: '1', email: 'a@b.c', username: 'a', role: 'admin', name: 'Admin' };

describe('api client', () => {
	beforeEach(() => {
		vi.stubGlobal('localStorage', fakeStorage());
		vi.stubGlobal('location', { pathname: '/admin/users', search: '?page=2' });
		goto.mockReset();
		session.logout();
	});
	afterEach(() => {
		session.logout();
		vi.unstubAllGlobals();
	});

	it('attaches the Bearer token from the session store', async () => {
		const seen = stubServer(200);
		session.login('jwt-123', user);

		await api.instance.get('/api/v1/users');

		expect(seen[0].headers.get('Authorization')).toBe('Bearer jwt-123');
	});

	it('sends no Authorization header without a session', async () => {
		const seen = stubServer(200);

		await api.instance.get('/api/v1/users');

		expect(seen[0].headers.get('Authorization')).toBeUndefined();
	});

	it('clears the session and redirects to login on 401', async () => {
		stubServer(401);
		session.login('expired', user);

		await expect(api.instance.get('/api/v1/users')).rejects.toBeInstanceOf(AxiosError);

		expect(session.isAuthenticated).toBe(false);
		expect(goto).toHaveBeenCalledTimes(1);
		expect(goto).toHaveBeenCalledWith(
			`${AUTH_ROUTES.login}?redirectTo=${encodeURIComponent('/admin/users?page=2')}`
		);
	});

	it('keeps the session when the login call itself returns 401', async () => {
		stubServer(401);
		session.login('still-valid', user);

		await expect(api.instance.post('/api/v1/auth/login', {})).rejects.toBeInstanceOf(AxiosError);

		expect(session.isAuthenticated).toBe(true);
		expect(goto).not.toHaveBeenCalled();
	});

	it('does not redirect on 401 when nobody is logged in', async () => {
		stubServer(401);

		await expect(api.instance.get('/api/v1/users')).rejects.toBeInstanceOf(AxiosError);

		expect(goto).not.toHaveBeenCalled();
	});

	it('leaves the session alone on other errors such as 403', async () => {
		stubServer(403);
		session.login('jwt', user);

		await expect(api.instance.get('/api/v1/users')).rejects.toBeInstanceOf(AxiosError);

		expect(session.isAuthenticated).toBe(true);
		expect(goto).not.toHaveBeenCalled();
	});
});
