import { API_ERROR_MESSAGES, fail, type Result } from '#lib/api/result.js';
import type { Profile, ProfileRepository } from '../domain/profile.types.js';
import { PROFILE_MESSAGES } from './profile.constants.js';

export async function updateProfileName(
	repo: ProfileRepository,
	name: string
): Promise<Result<Profile>> {
	const trimmed = name.trim();
	if (!trimmed) {
		return fail('VALIDATION_ERROR', API_ERROR_MESSAGES.validation, {
			name: [PROFILE_MESSAGES.nameRequired]
		});
	}
	return repo.updateName(trimmed);
}
