<script lang="ts">
	import * as Field from '#lib/components/ui/field/index.js';
	import { Input } from '#lib/components/ui/input/index.js';
	import type { HTMLInputAttributes } from 'svelte/elements';

	type Props = {
		id: string;
		label: string;
		value: string;
		errors?: string[];
		type?: 'text' | 'email' | 'password';
		autocomplete?: HTMLInputAttributes['autocomplete'];
		disabled?: boolean;
		readonly?: boolean;
	};

	let {
		id,
		label,
		value = $bindable(),
		errors = [],
		type = 'text',
		autocomplete,
		disabled = false,
		readonly = false
	}: Props = $props();

	const invalid = $derived(errors.length > 0);
</script>

<Field.Field data-invalid={invalid}>
	<Field.Label for={id}>{label}</Field.Label>
	<Input
		{id}
		name={id}
		{type}
		{autocomplete}
		{disabled}
		{readonly}
		aria-invalid={invalid}
		aria-describedby={invalid ? `${id}-error` : undefined}
		bind:value
	/>
	{#if invalid}
		<Field.Error id="{id}-error" errors={errors.map((message) => ({ message }))} />
	{/if}
</Field.Field>
