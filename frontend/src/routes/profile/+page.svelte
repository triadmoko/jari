<script lang="ts">
	import AppHeader from '#lib/components/app-header.svelte';
	import ProfileForm from '#lib/features/profile/components/profile-form.svelte';
</script>

<svelte:head><title>Profil · Jari</title></svelte:head>

<AppHeader />
<main class="mx-auto flex max-w-4xl justify-center p-6">
	<ProfileForm />
</main>
