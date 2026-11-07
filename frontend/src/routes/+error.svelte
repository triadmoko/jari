<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { pageTitle } from '#lib/brand.js';
	import { AUTH_ROUTES } from '#lib/features/auth/index.js';
	import { t } from '#lib/i18n/index.js';
</script>

<svelte:head><title>{pageTitle(String(page.status))}</title></svelte:head>

<main class="flex min-h-svh flex-col items-center justify-center gap-2 p-6 text-center">
	<h1 class="text-4xl font-semibold">{page.status}</h1>
	<p class="text-muted-foreground">
		{page.status === 404 ? t('error.notFound') : (page.error?.message ?? t('error.generic'))}
	</p>
	<a href={resolve(AUTH_ROUTES.home)} class="mt-2 underline underline-offset-4"
		>{t('error.backHome')}</a
	>
</main>
