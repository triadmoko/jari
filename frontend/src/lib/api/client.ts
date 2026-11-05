import { goto } from '$app/navigation';
import { AUTH_ROUTES } from '#lib/features/auth/application/auth.constants.js';
import { session } from '#lib/features/auth/application/session.svelte.js';
import { Api } from './api.js';

/**
 * Browser API client (SPA: there is no server-side client).
 *
 * The generated paths already contain `/api/v1`, so baseURL is only the origin.
 * Empty VITE_API_BASE_URL (inlined at build time) = same origin. The generator falls back to
 * localhost when baseURL is falsy, hence `'/'`.
 */
export const api = new Api({
	baseURL: (import.meta.env.VITE_API_BASE_URL as string | undefined) || '/'
});

api.instance.interceptors.request.use((config) => {
	const token = session.token;
	if (token) {
		config.headers.set('Authorization', `Bearer ${token}`);
		session.touch();
	}
	return config;
});

api.instance.interceptors.response.use(
	(response) => response,
	(error) => {
		const status = error?.response?.status;
		const url: string = error?.config?.url ?? '';
		// A 401 from the login call itself means wrong credentials, not an expired session.
		if (status === 401 && session.isAuthenticated && !url.endsWith('/auth/login')) {
			session.logout();
			const here = location.pathname + location.search;
			// The /auth/login page is not typed in resolve() yet and the app has no base path.
			// eslint-disable-next-line svelte/no-navigation-without-resolve
			void goto(`${AUTH_ROUTES.login}?redirectTo=${encodeURIComponent(here)}`);
		}
		return Promise.reject(error);
	}
);
