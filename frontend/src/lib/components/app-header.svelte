<script lang="ts">
	import { resolve } from '$app/paths';
	import { Button } from '#lib/components/ui/button/index.js';
	import { AUTH_ROUTES, session } from '#lib/features/auth/index.js';
	import LogoutButton from '#lib/features/auth/components/logout-button.svelte';
</script>

<header class="border-b">
	<div class="mx-auto flex h-14 max-w-4xl items-center justify-between px-6">
		<a href={resolve(AUTH_ROUTES.home)} class="font-semibold">Jari</a>
		<nav class="flex items-center gap-1">
			{#if session.isAuthenticated}
				<Button variant="ghost" size="sm" href={resolve(AUTH_ROUTES.profile)}>Profil</Button>
				<LogoutButton />
			{:else}
				<Button variant="ghost" size="sm" href={resolve(AUTH_ROUTES.login)}>Masuk</Button>
				<Button size="sm" href={resolve(AUTH_ROUTES.register)}>Daftar</Button>
			{/if}
		</nav>
	</div>
</header>
