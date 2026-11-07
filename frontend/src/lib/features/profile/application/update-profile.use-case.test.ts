import { describe, expect, it, vi } from 'vitest';
import { ok } from '#lib/api/result.js';
import type { Profile, ProfileRepository } from '../domain/profile.types.js';
import { updateProfileName } from './update-profile.use-case.js';

const profile: Profile = { id: '1', email: 'a@b.co', username: 'abc', name: 'Baru', role: 'user' };

function repo(): ProfileRepository {
	return { get: vi.fn(async () => ok(profile)), updateName: vi.fn(async () => ok(profile)) };
}

describe('updateProfileName', () => {
	it('rejects a blank name locally', async () => {
		const r = repo();
		const result = await updateProfileName(r, '   ');
		expect(result).toMatchObject({ ok: false, code: 'VALIDATION_ERROR' });
		expect(r.updateName).not.toHaveBeenCalled();
	});

	it('sends the trimmed name', async () => {
		const r = repo();
		const result = await updateProfileName(r, ' Baru ');
		expect(result.ok).toBe(true);
		expect(r.updateName).toHaveBeenCalledWith('Baru');
	});
});
