import { API_ERROR_MESSAGES, fail, type FieldErrors, type Result } from '#lib/api/result.js';
import type { AuthRepository, RegisterInput, SessionUser } from '../domain/auth.types.js';
import { AUTH_LIMITS, AUTH_MESSAGES } from './auth.constants.js';
import { isValidEmail } from './validation.js';

export function validateRegister(input: RegisterInput): FieldErrors {
	const errors: FieldErrors = {};
	const username = input.username.trim();

	if (!input.name.trim()) errors.name = [AUTH_MESSAGES.required('Nama')];

	if (!username) errors.username = [AUTH_MESSAGES.required('Username')];
	else if (username.length < AUTH_LIMITS.usernameMin)
		errors.username = [AUTH_MESSAGES.usernameTooShort];
	else if (username.length > AUTH_LIMITS.usernameMax)
		errors.username = [AUTH_MESSAGES.usernameTooLong];

	if (!input.email.trim()) errors.email = [AUTH_MESSAGES.required('Email')];
	else if (!isValidEmail(input.email)) errors.email = [AUTH_MESSAGES.invalidEmail];

	if (!input.password) errors.password = [AUTH_MESSAGES.required('Kata sandi')];
	else if (input.password.length < AUTH_LIMITS.passwordMin) {
		errors.password = [AUTH_MESSAGES.passwordTooShort];
	}

	return errors;
}

export async function register(
	repo: AuthRepository,
	input: RegisterInput
): Promise<Result<SessionUser>> {
	const fieldErrors = validateRegister(input);
	if (Object.keys(fieldErrors).length > 0) {
		return fail('VALIDATION_ERROR', API_ERROR_MESSAGES.validation, fieldErrors);
	}
	return repo.register({
		name: input.name.trim(),
		username: input.username.trim(),
		email: input.email.trim(),
		password: input.password
	});
}
