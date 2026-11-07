import type { Result } from '#lib/api/result.js';

export type Profile = {
	id: string;
	email: string;
	username: string;
	name: string;
	role: string;
};

export interface ProfileRepository {
	get(): Promise<Result<Profile>>;
	updateName(name: string): Promise<Result<Profile>>;
}
