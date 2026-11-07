import { isAxiosError } from 'axios';

export type ApiErrorCode =
	'UNAUTHORIZED' | 'FORBIDDEN' | 'INVALID_CREDENTIALS' | 'VALIDATION_ERROR' | 'UNKNOWN_ERROR';

/** Per-field messages, keyed by the JSON field name used by the backend. */
export type FieldErrors = Record<string, string[]>;

export type Failure = {
	ok: false;
	code: ApiErrorCode;
	message: string;
	fieldErrors?: FieldErrors;
};

export type Result<T> = { ok: true; value: T } | Failure;

export const API_ERROR_MESSAGES = {
	validation: 'Periksa kembali isian Anda.',
	unauthorized: 'Sesi Anda tidak valid. Silakan masuk kembali.',
	invalidCredentials: 'Email atau kata sandi salah.',
	forbidden: 'Anda tidak memiliki akses untuk tindakan ini.',
	network: 'Tidak dapat terhubung ke server. Coba lagi beberapa saat.',
	unknown: 'Terjadi kesalahan. Coba lagi beberapa saat.'
} as const;

export function ok<T>(value: T): Result<T> {
	return { ok: true, value };
}

export function fail(code: ApiErrorCode, message: string, fieldErrors?: FieldErrors): Failure {
	return fieldErrors ? { ok: false, code, message, fieldErrors } : { ok: false, code, message };
}

function parseFieldErrors(value: unknown): FieldErrors | undefined {
	if (typeof value !== 'object' || value === null || Array.isArray(value)) return undefined;

	const out: FieldErrors = {};
	for (const [field, messages] of Object.entries(value)) {
		if (Array.isArray(messages)) {
			const strings = messages.filter((m): m is string => typeof m === 'string');
			if (strings.length > 0) out[field] = strings;
		}
	}
	return Object.keys(out).length > 0 ? out : undefined;
}

/**
 * Maps an axios error to a domain failure. The backend envelope is
 * `{ error, status, message, data, errors }`; `errors` carries per-field validation messages.
 *
 * `unauthorizedCode` lets login report a 401 as INVALID_CREDENTIALS instead of an expired session.
 */
export function toFailure(
	error: unknown,
	options: { unauthorizedCode?: 'UNAUTHORIZED' | 'INVALID_CREDENTIALS' } = {}
): Failure {
	if (!isAxiosError(error) || !error.response) {
		return fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.network);
	}

	const { status, data } = error.response;
	const body = (typeof data === 'object' && data !== null ? data : {}) as {
		message?: unknown;
		errors?: unknown;
	};
	const serverMessage = typeof body.message === 'string' && body.message ? body.message : null;

	if (status === 400 || status === 422) {
		const fieldErrors = parseFieldErrors(body.errors);
		// With field errors the UI shows them inline; otherwise the server message is the reason
		// (e.g. email already registered).
		return fail(
			'VALIDATION_ERROR',
			fieldErrors
				? API_ERROR_MESSAGES.validation
				: (serverMessage ?? API_ERROR_MESSAGES.validation),
			fieldErrors
		);
	}

	if (status === 401) {
		return options.unauthorizedCode === 'INVALID_CREDENTIALS'
			? fail('INVALID_CREDENTIALS', API_ERROR_MESSAGES.invalidCredentials)
			: fail('UNAUTHORIZED', API_ERROR_MESSAGES.unauthorized);
	}

	if (status === 403) return fail('FORBIDDEN', API_ERROR_MESSAGES.forbidden);

	return fail('UNKNOWN_ERROR', API_ERROR_MESSAGES.unknown);
}
