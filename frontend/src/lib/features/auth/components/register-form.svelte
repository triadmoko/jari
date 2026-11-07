<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import type { FieldErrors } from '#lib/api/result.js';
	import FormField from '#lib/components/form-field.svelte';
	import * as Alert from '#lib/components/ui/alert/index.js';
	import { Button } from '#lib/components/ui/button/index.js';
	import * as Card from '#lib/components/ui/card/index.js';
	import * as Field from '#lib/components/ui/field/index.js';
	import { Spinner } from '#lib/components/ui/spinner/index.js';
	import { AUTH_ROUTES } from '../application/auth.constants.js';
	import { register } from '../application/register.use-case.js';
	import { apiAuthRepository } from '../infrastructure/api-auth.client.js';

	let name = $state('');
	let username = $state('');
	let email = $state('');
	let password = $state('');
	let submitting = $state(false);
	let formError = $state<string | null>(null);
	let fieldErrors = $state<FieldErrors>({});

	async function onsubmit(event: SubmitEvent) {
		event.preventDefault();
		if (submitting) return;

		submitting = true;
		formError = null;
		fieldErrors = {};

		const result = await register(apiAuthRepository, { name, username, email, password });
		submitting = false;

		if (!result.ok) {
			fieldErrors = result.fieldErrors ?? {};
			formError = result.fieldErrors ? null : result.message;
			return;
		}

		await goto(`${resolve(AUTH_ROUTES.login)}?registered=1`);
	}
</script>

<Card.Root class="w-full max-w-sm">
	<Card.Header>
		<Card.Title>Daftar</Card.Title>
		<Card.Description>Buat akun baru untuk mulai memakai Jari.</Card.Description>
	</Card.Header>
	<Card.Content>
		<form {onsubmit} novalidate class="flex flex-col gap-4">
			{#if formError}
				<Alert.Root variant="destructive">
					<Alert.Description>{formError}</Alert.Description>
				</Alert.Root>
			{/if}

			<Field.Group>
				<FormField
					id="name"
					label="Nama"
					autocomplete="name"
					bind:value={name}
					errors={fieldErrors.name}
					disabled={submitting}
				/>
				<FormField
					id="username"
					label="Username"
					autocomplete="username"
					bind:value={username}
					errors={fieldErrors.username}
					disabled={submitting}
				/>
				<FormField
					id="email"
					label="Email"
					type="email"
					autocomplete="email"
					bind:value={email}
					errors={fieldErrors.email}
					disabled={submitting}
				/>
				<FormField
					id="password"
					label="Kata sandi"
					type="password"
					autocomplete="new-password"
					bind:value={password}
					errors={fieldErrors.password}
					disabled={submitting}
				/>
			</Field.Group>

			<Button type="submit" disabled={submitting}>
				{#if submitting}<Spinner />{/if}
				Daftar
			</Button>
		</form>
	</Card.Content>
	<Card.Footer class="justify-center text-sm">
		<span class="text-muted-foreground">Sudah punya akun?&nbsp;</span>
		<a href={resolve(AUTH_ROUTES.login)} class="underline underline-offset-4">Masuk</a>
	</Card.Footer>
</Card.Root>
