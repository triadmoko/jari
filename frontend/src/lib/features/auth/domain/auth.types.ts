import type { Result } from '#lib/api/result.js';

export type SessionUser = {
	id: string;
	email: string;
	username: string;
	role: string;
	/** Display name; mirrors `name` in the backend DTOs. Optional so older stored sessions still load. */
	name?: string;
};

export type Session = {
	token: string;
	user: SessionUser;
	/** Epoch millis. Session is treated as logged out after this moment. */
	expiresAt: number;
};

export type LoginInput = { email: string; password: string };
export type RegisterInput = { name: string; username: string; email: string; password: string };
export type LoginOutput = { token: string; user: SessionUser };

export interface AuthRepository {
	login(input: LoginInput): Promise<Result<LoginOutput>>;
	register(input: RegisterInput): Promise<Result<SessionUser>>;
}
