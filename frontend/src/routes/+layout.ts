import { redirect } from '@sveltejs/kit';
import { resolveGuard } from '#lib/features/auth/application/route-guard.js';
import { readSession } from '#lib/features/auth/infrastructure/session.storage.js';

// SPA: no server rendering, everything runs in the browser.
export const ssr = false;
export const prerender = false;

export const load = ({ url }) => {
	const target = resolveGuard(url.pathname, url.search, readSession());
	if (target) redirect(307, target);
};
