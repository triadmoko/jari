<script lang="ts">
	import AppHeader from '#lib/components/app-header.svelte';
	import { pageTitle } from '#lib/brand.js';
	import { t } from '#lib/i18n/index.js';
	import ProfileForm from '#lib/features/profile/components/profile-form.svelte';
</script>

<svelte:head><title>{pageTitle(t('profile.title'))}</title></svelte:head>

<AppHeader />
<main class="mx-auto flex max-w-4xl justify-center p-6">
	<ProfileForm />
</main>
