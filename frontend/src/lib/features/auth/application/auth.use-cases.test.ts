import { describe, expect, it, vi } from 'vitest';
import { ok } from '#lib/api/result.js';
import type { AuthRepository } from '../domain/auth.types.js';
import { login } from './login.use-case.js';
import { register, validateRegister } from './register.use-case.js';

const user = { id: '1', email: 'a@b.co', username: 'abc', role: 'user', name: 'A' };

function repo(): AuthRepository {
	return {
		login: vi.fn(async () => ok({ token: 't', user })),
		register: vi.fn(async () => ok(user))
	};
}

describe('login use case', () => {
	it('rejects empty and malformed input without calling the repository', async () => {
		const r = repo();
		const result = await login(r, { email: 'nope', password: '' });
		expect(result).toMatchObject({ ok: false, code: 'VALIDATION_ERROR' });
		expect(!result.ok && result.fieldErrors).toHaveProperty('email');
		expect(!result.ok && result.fieldErrors).toHaveProperty('password');
		expect(r.login).not.toHaveBeenCalled();
	});

	it('trims the email and delegates to the repository', async () => {
		const r = repo();
		const result = await login(r, { email: ' a@b.co ', password: 'secret' });
		expect(result.ok).toBe(true);
		expect(r.login).toHaveBeenCalledWith({ email: 'a@b.co', password: 'secret' });
	});
});

describe('register use case', () => {
	const valid = { name: 'Ana', username: 'ana', email: 'ana@b.co', password: 'secret' };

	it('mirrors the backend limits', () => {
		expect(validateRegister(valid)).toEqual({});
		expect(validateRegister({ ...valid, username: 'ab' })).toHaveProperty('username');
		expect(validateRegister({ ...valid, username: 'x'.repeat(21) })).toHaveProperty('username');
		expect(validateRegister({ ...valid, password: '12345' })).toHaveProperty('password');
		expect(validateRegister({ ...valid, name: '  ' })).toHaveProperty('name');
	});

	it('does not call the repository when invalid', async () => {
		const r = repo();
		const result = await register(r, { ...valid, email: 'bad' });
		expect(result.ok).toBe(false);
		expect(r.register).not.toHaveBeenCalled();
	});

	it('delegates trimmed input when valid', async () => {
		const r = repo();
		await register(r, { ...valid, name: ' Ana ' });
		expect(r.register).toHaveBeenCalledWith({ ...valid, name: 'Ana' });
	});
});
