import { api } from '#lib/api/client.js';
import type { DtoAuthRegisterResponse } from '#lib/api/api.js';
import { API_ERROR_MESSAGES, fail, ok, toFailure } from '#lib/api/result.js';
import type { AuthRepository, SessionUser } from '../domain/auth.types.js';

function toSessionUser(dto: DtoAuthRegisterResponse | undefined): SessionUser | null {
	if (!dto?.id) return null;
	return {
		id: dto.id,
		email: dto.email ?? '',
		username: dto.username ?? '',
		role: dto.role ?? 'user',
		name: dto.name
	};
}

export const apiAuthRepository: AuthRepository = {
	async login(input) {
		try {
			const res = await api.api.v1AuthLoginCreate(input);
			const token = res.data.data?.token;
			const user = toSessionUser(res.data.data?.user);
			if (!token || !user) return fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.unknown);
			return ok({ token, user });
		} catch (error) {
			return toFailure(error, { unauthorizedCode: 'INVALID_CREDENTIALS' });
		}
	},

	async register(input) {
		try {
			const res = await api.api.v1AuthRegisterCreate(input);
			const user = toSessionUser(res.data.data);
			if (!user) return fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.unknown);
			return ok(user);
		} catch (error) {
			return toFailure(error);
		}
	}
};
