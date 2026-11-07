<script lang="ts">
	import { pageTitle } from '#lib/brand.js';
	import LocaleSwitcher from '#lib/components/locale-switcher.svelte';
	import { t } from '#lib/i18n/index.js';
	import RegisterForm from '#lib/features/auth/components/register-form.svelte';
</script>

<svelte:head><title>{pageTitle(t('auth.register.title'))}</title></svelte:head>

<main class="relative flex min-h-svh items-center justify-center p-6">
	<div class="absolute top-4 right-4"><LocaleSwitcher /></div>
	<RegisterForm />
</main>
