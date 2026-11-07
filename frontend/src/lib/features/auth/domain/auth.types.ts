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
