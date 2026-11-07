import { api } from '#lib/api/client.js';
import type { DtoUserUserResponse } from '#lib/api/api.js';
import { API_ERROR_MESSAGES, fail, ok, toFailure } from '#lib/api/result.js';
import type { Profile, ProfileRepository } from '../domain/profile.types.js';

function toProfile(dto: DtoUserUserResponse | undefined): Profile | null {
	if (!dto?.id) return null;
	return {
		id: dto.id,
		email: dto.email ?? '',
		username: dto.username ?? '',
		name: dto.name ?? '',
		role: dto.role ?? 'user'
	};
}

export const apiProfileRepository: ProfileRepository = {
	async get() {
		try {
			const res = await api.api.v1UsersProfileList();
			const profile = toProfile(res.data.data);
			return profile ? ok(profile) : fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.unknown);
		} catch (error) {
			return toFailure(error);
		}
	},

	async updateName(name) {
		try {
			const res = await api.api.v1UsersProfileUpdate({ name });
			const profile = toProfile(res.data.data);
			return profile ? ok(profile) : fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.unknown);
		} catch (error) {
			return toFailure(error);
		}
	}
};
