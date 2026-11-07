<script lang="ts">
	import { pageTitle } from '#lib/brand.js';
	import RegisterForm from '#lib/features/auth/components/register-form.svelte';
</script>

<svelte:head><title>{pageTitle('Daftar')}</title></svelte:head>

<main class="flex min-h-svh items-center justify-center p-6">
	<RegisterForm />
</main>
