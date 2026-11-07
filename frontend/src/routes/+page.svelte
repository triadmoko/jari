<script lang="ts">
	import { resolve } from '$app/paths';
	import AppHeader from '#lib/components/app-header.svelte';
	import { PRODUCT_NAME, pageTitle } from '#lib/brand.js';
	import { Button } from '#lib/components/ui/button/index.js';
	import { AUTH_ROUTES, session } from '#lib/features/auth/index.js';

	const displayName = $derived(session.user?.name || session.user?.username);
</script>

<svelte:head><title>{pageTitle('Beranda')}</title></svelte:head>

<AppHeader />
<main class="mx-auto flex max-w-4xl flex-col gap-4 p-6">
	{#if session.isAuthenticated}
		<h1 class="text-2xl font-semibold">Halo, {displayName}</h1>
		<p class="text-muted-foreground">Ruang kerja Anda akan tampil di sini.</p>
		<div>
			<Button variant="outline" href={resolve(AUTH_ROUTES.profile)}>Lihat profil</Button>
		</div>
	{:else}
		<h1 class="text-2xl font-semibold">Selamat datang di {PRODUCT_NAME}</h1>
		<p class="text-muted-foreground">Masuk atau daftar untuk mulai bekerja.</p>
	{/if}
</main>
