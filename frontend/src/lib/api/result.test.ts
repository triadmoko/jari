import { AxiosError, type AxiosResponse } from 'axios';
import { describe, expect, it } from 'vitest';
import { API_ERROR_MESSAGES, toFailure } from './result.js';

function httpError(status: number, data: unknown): AxiosError {
	const error = new AxiosError('failed');
	error.response = { status, data } as AxiosResponse;
	return error;
}

describe('toFailure', () => {
	it('maps 400 with field errors to VALIDATION_ERROR and keeps the fields', () => {
		const failure = toFailure(
			httpError(400, { message: 'Validation failed', errors: { email: ['invalid'] } })
		);
		expect(failure).toMatchObject({
			code: 'VALIDATION_ERROR',
			message: API_ERROR_MESSAGES.validation,
			fieldErrors: { email: ['invalid'] }
		});
	});

	it('uses the server message for 400 without field errors (e.g. duplicate email)', () => {
		const failure = toFailure(httpError(400, { message: 'user already exists' }));
		expect(failure).toMatchObject({ code: 'VALIDATION_ERROR', message: 'user already exists' });
		expect(failure.fieldErrors).toBeUndefined();
	});

	it('maps 401 to INVALID_CREDENTIALS only when asked to', () => {
		expect(toFailure(httpError(401, {}), { unauthorizedCode: 'INVALID_CREDENTIALS' }).code).toBe(
			'INVALID_CREDENTIALS'
		);
		expect(toFailure(httpError(401, {})).code).toBe('UNAUTHORIZED');
	});

	it('maps 403, 5xx and network errors', () => {
		expect(toFailure(httpError(403, {})).code).toBe('FORBIDDEN');
		expect(toFailure(httpError(500, {})).code).toBe('UNKNOWN_ERROR');
		expect(toFailure(new AxiosError('Network Error'))).toMatchObject({
			code: 'UNKNOWN_ERROR',
			message: API_ERROR_MESSAGES.network
		});
	});
});
