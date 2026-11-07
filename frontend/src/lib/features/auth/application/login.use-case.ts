import { fail, type FieldErrors, type Result } from '#lib/api/result.js';
import { API_ERROR_MESSAGES } from '#lib/api/result.js';
import type { AuthRepository, LoginInput, LoginOutput } from '../domain/auth.types.js';
import { AUTH_MESSAGES } from './auth.constants.js';
import { isValidEmail } from './validation.js';

export function validateLogin(input: LoginInput): FieldErrors {
	const errors: FieldErrors = {};

	if (!input.email.trim()) errors.email = [AUTH_MESSAGES.required('Email')];
	else if (!isValidEmail(input.email)) errors.email = [AUTH_MESSAGES.invalidEmail];

	if (!input.password) errors.password = [AUTH_MESSAGES.required('Kata sandi')];

	return errors;
}

/** Validates locally, then delegates to the repository. Persisting the session is the caller's job. */
export async function login(repo: AuthRepository, input: LoginInput): Promise<Result<LoginOutput>> {
	const fieldErrors = validateLogin(input);
	if (Object.keys(fieldErrors).length > 0) {
		return fail('VALIDATION_ERROR', API_ERROR_MESSAGES.validation, fieldErrors);
	}
	return repo.login({ email: input.email.trim(), password: input.password });
}
