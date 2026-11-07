<script lang="ts">
	import { resolve } from '$app/paths';
	import AppHeader from '#lib/components/app-header.svelte';
	import { PRODUCT_NAME, pageTitle } from '#lib/brand.js';
	import { Button } from '#lib/components/ui/button/index.js';
	import { AUTH_ROUTES, session } from '#lib/features/auth/index.js';
	import { t } from '#lib/i18n/index.js';

	const displayName = $derived(session.user?.name || session.user?.username);
</script>

<svelte:head><title>{pageTitle(t('page.home'))}</title></svelte:head>

<AppHeader />
<main class="mx-auto flex max-w-4xl flex-col gap-4 p-6">
	{#if session.isAuthenticated}
		<h1 class="text-2xl font-semibold">{t('home.greeting', { name: displayName ?? '' })}</h1>
		<p class="text-muted-foreground">{t('home.spacesPlaceholder')}</p>
		<div>
			<Button variant="outline" href={resolve(AUTH_ROUTES.profile)}>{t('home.viewProfile')}</Button>
		</div>
	{:else}
		<h1 class="text-2xl font-semibold">{t('home.welcome', { product: PRODUCT_NAME })}</h1>
		<p class="text-muted-foreground">{t('home.welcomeHint')}</p>
	{/if}
</main>
