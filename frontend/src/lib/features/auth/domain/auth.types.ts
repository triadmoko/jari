export type SessionUser = {
	id: string;
	email: string;
	username: string;
	role: string;
	first_name?: string;
	last_name?: string;
};

export type Session = {
	token: string;
	user: SessionUser;
	/** Epoch millis. Session is treated as logged out after this moment. */
	expiresAt: number;
};
