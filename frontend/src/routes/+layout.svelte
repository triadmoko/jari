<script lang="ts">
	import './layout.css';
	import { Toaster } from '$lib/components/ui/sonner/index.js';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<Toaster position="bottom-right" closeButton={true} />
{@render children()}
