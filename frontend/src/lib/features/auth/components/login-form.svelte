<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import type { FieldErrors } from '#lib/api/result.js';
	import FormField from '#lib/components/form-field.svelte';
	import * as Alert from '#lib/components/ui/alert/index.js';
	import { Button } from '#lib/components/ui/button/index.js';
	import * as Card from '#lib/components/ui/card/index.js';
	import { Spinner } from '#lib/components/ui/spinner/index.js';
	import * as Field from '#lib/components/ui/field/index.js';
	import { i18n, t } from '#lib/i18n/index.js';
	import { AUTH_ROUTES } from '../application/auth.constants.js';
	import { login } from '../application/login.use-case.js';
	import { postLoginTarget } from '../application/route-guard.js';
	import { session } from '../application/session.svelte.js';
	import { apiAuthRepository } from '../infrastructure/api-auth.client.js';

	let {
		redirectTo = null,
		registered = false
	}: { redirectTo?: string | null; registered?: boolean } = $props();

	let email = $state('');
	let password = $state('');
	let submitting = $state(false);
	let formError = $state<string | null>(null);
	let fieldErrors = $state<FieldErrors>({});

	$effect(() => {
		void i18n.locale;
		formError = null;
		fieldErrors = {};
	});

	async function onsubmit(event: SubmitEvent) {
		event.preventDefault();
		if (submitting) return;

		submitting = true;
		formError = null;
		fieldErrors = {};

		const result = await login(apiAuthRepository, { email, password });
		submitting = false;

		if (!result.ok) {
			fieldErrors = result.fieldErrors ?? {};
			formError = result.fieldErrors ? null : result.message;
			return;
		}

		session.login(result.value.token, result.value.user);
		// The target is a runtime path (redirectTo), which resolve() cannot type.
		// eslint-disable-next-line svelte/no-navigation-without-resolve
		await goto(postLoginTarget(redirectTo));
	}
</script>

<Card.Root class="w-full max-w-sm">
	<Card.Header>
		<Card.Title>{t('auth.login.title')}</Card.Title>
		<Card.Description>{t('auth.login.description')}</Card.Description>
	</Card.Header>
	<Card.Content>
		<form {onsubmit} novalidate class="flex flex-col gap-4">
			{#if registered && !formError}
				<Alert.Root role="status">
					<Alert.Description>{t('auth.registered')}</Alert.Description>
				</Alert.Root>
			{/if}
			{#if formError}
				<Alert.Root variant="destructive">
					<Alert.Description>{formError}</Alert.Description>
				</Alert.Root>
			{/if}

			<Field.Group>
				<FormField
					id="email"
					label={t('field.email')}
					type="email"
					autocomplete="email"
					bind:value={email}
					errors={fieldErrors.email}
					disabled={submitting}
				/>
				<FormField
					id="password"
					label={t('field.password')}
					type="password"
					autocomplete="current-password"
					bind:value={password}
					errors={fieldErrors.password}
					disabled={submitting}
				/>
			</Field.Group>

			<Button type="submit" disabled={submitting}>
				{#if submitting}<Spinner />{/if}
				{t('action.login')}
			</Button>
		</form>
	</Card.Content>
	<Card.Footer class="justify-center text-sm">
		<span class="text-muted-foreground">{t('auth.login.noAccount')}&nbsp;</span>
		<a href={resolve(AUTH_ROUTES.register)} class="underline underline-offset-4">
			{t('action.register')}
		</a>
	</Card.Footer>
</Card.Root>
