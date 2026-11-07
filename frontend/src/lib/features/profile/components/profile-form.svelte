<script lang="ts">
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';
	import type { FieldErrors } from '#lib/api/result.js';
	import FormField from '#lib/components/form-field.svelte';
	import * as Alert from '#lib/components/ui/alert/index.js';
	import { Button } from '#lib/components/ui/button/index.js';
	import * as Card from '#lib/components/ui/card/index.js';
	import * as Field from '#lib/components/ui/field/index.js';
	import { Skeleton } from '#lib/components/ui/skeleton/index.js';
	import { Spinner } from '#lib/components/ui/spinner/index.js';
	import { session } from '#lib/features/auth/index.js';
	import { PROFILE_MESSAGES } from '../application/profile.constants.js';
	import { updateProfileName } from '../application/update-profile.use-case.js';
	import type { Profile } from '../domain/profile.types.js';
	import { apiProfileRepository } from '../infrastructure/api-profile.client.js';

	let profile = $state<Profile | null>(null);
	let name = $state('');
	let loading = $state(true);
	let submitting = $state(false);
	let loadError = $state<string | null>(null);
	let formError = $state<string | null>(null);
	let fieldErrors = $state<FieldErrors>({});

	onMount(async () => {
		const result = await apiProfileRepository.get();
		loading = false;
		if (!result.ok) {
			loadError = result.message;
			return;
		}
		profile = result.value;
		name = result.value.name;
	});

	async function onsubmit(event: SubmitEvent) {
		event.preventDefault();
		if (submitting) return;

		submitting = true;
		formError = null;
		fieldErrors = {};

		const result = await updateProfileName(apiProfileRepository, name);
		submitting = false;

		if (!result.ok) {
			fieldErrors = result.fieldErrors ?? {};
			formError = result.fieldErrors ? null : result.message;
			return;
		}

		profile = result.value;
		name = result.value.name;
		if (session.user) session.updateUser({ ...session.user, name: result.value.name });
		toast.success(PROFILE_MESSAGES.updated);
	}
</script>

<Card.Root class="w-full max-w-md">
	<Card.Header>
		<Card.Title>Profil</Card.Title>
		<Card.Description>Lihat dan ubah nama tampilan Anda.</Card.Description>
	</Card.Header>
	<Card.Content>
		{#if loading}
			<div class="flex flex-col gap-4" aria-busy="true">
				<Skeleton class="h-9 w-full" />
				<Skeleton class="h-9 w-full" />
				<Skeleton class="h-9 w-full" />
			</div>
		{:else if loadError}
			<Alert.Root variant="destructive">
				<Alert.Description>{loadError}</Alert.Description>
			</Alert.Root>
		{:else if profile}
			<form {onsubmit} novalidate class="flex flex-col gap-4">
				{#if formError}
					<Alert.Root variant="destructive">
						<Alert.Description>{formError}</Alert.Description>
					</Alert.Root>
				{/if}

				<Field.Group>
					<FormField id="email" label="Email" value={profile.email} readonly />
					<FormField id="username" label="Username" value={profile.username} readonly />
					<FormField
						id="name"
						label="Nama"
						autocomplete="name"
						bind:value={name}
						errors={fieldErrors.name}
						disabled={submitting}
					/>
				</Field.Group>

				<Button type="submit" disabled={submitting}>
					{#if submitting}<Spinner />{/if}
					Simpan
				</Button>
			</form>
		{/if}
	</Card.Content>
</Card.Root>
