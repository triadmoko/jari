<script lang="ts">
	import { page } from '$app/state';
	import { pageTitle } from '#lib/brand.js';
	import LoginForm from '#lib/features/auth/components/login-form.svelte';
	import { safeRedirectTarget } from '#lib/features/auth/index.js';

	const redirectTo = $derived(safeRedirectTarget(page.url.searchParams.get('redirectTo')));
	const registered = $derived(page.url.searchParams.get('registered') === '1');
</script>

<svelte:head><title>{pageTitle('Masuk')}</title></svelte:head>

<main class="flex min-h-svh items-center justify-center p-6">
	<LoginForm {redirectTo} {registered} />
</main>
