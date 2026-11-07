import { describe, expect, it } from 'vitest';
import type { Session } from '../domain/auth.types.js';
import { postLoginTarget, resolveGuard, safeRedirectTarget } from './route-guard.js';

const session = (role = 'user'): Session => ({
	token: 't',
	expiresAt: Date.now() + 1000,
	user: { id: '1', email: 'a@b.c', username: 'a', role }
});

describe('resolveGuard', () => {
	it('redirects protected routes to login with redirectTo when logged out', () => {
		expect(resolveGuard('/admin/users', '?page=2', null)).toBe(
			'/auth/login?redirectTo=' + encodeURIComponent('/admin/users?page=2')
		);
		expect(resolveGuard('/profile', '', null)).toContain('/auth/login?redirectTo=');
	});

	it('does not treat lookalike prefixes as protected', () => {
		expect(resolveGuard('/administrator', '', null)).toBeNull();
	});

	it('allows protected routes when logged in', () => {
		expect(resolveGuard('/admin', '', session())).toBeNull();
	});

	it('sends logged-in users away from the guest pages', () => {
		expect(resolveGuard('/auth/login', '', session())).toBe('/');
		expect(resolveGuard('/auth/register', '', session())).toBe('/');
	});

	it('allows public routes for everyone', () => {
		expect(resolveGuard('/', '', null)).toBeNull();
		expect(resolveGuard('/auth/login', '', null)).toBeNull();
		expect(resolveGuard('/auth/register', '', null)).toBeNull();
	});
});

describe('safeRedirectTarget', () => {
	it('accepts same-origin paths only', () => {
		expect(safeRedirectTarget('/admin?x=1')).toBe('/admin?x=1');
		expect(safeRedirectTarget('//evil.com')).toBeNull();
		expect(safeRedirectTarget('https://evil.com')).toBeNull();
		expect(safeRedirectTarget('/\\evil.com')).toBeNull();
		expect(safeRedirectTarget('')).toBeNull();
		expect(safeRedirectTarget(null)).toBeNull();
	});
});

describe('postLoginTarget', () => {
	it('prefers a safe redirectTo', () => {
		expect(postLoginTarget('/profile')).toBe('/profile');
	});
	it('falls back to home for missing or unsafe targets', () => {
		expect(postLoginTarget('//evil.com')).toBe('/');
		expect(postLoginTarget()).toBe('/');
	});
});
