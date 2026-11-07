<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { Button } from '#lib/components/ui/button/index.js';
	import { AUTH_ROUTES } from '../application/auth.constants.js';
	import { session } from '../application/session.svelte.js';

	async function logout() {
		session.logout();
		await goto(resolve(AUTH_ROUTES.login));
	}
</script>

<Button variant="ghost" size="sm" onclick={logout}>Keluar</Button>
