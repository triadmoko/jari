<script lang="ts">
	import './layout.css';
	import { Toaster } from '#lib/components/ui/sonner/index.js';
	import logo from '#lib/assets/logo.svg';
	import { PRODUCT_DESCRIPTION, pageTitle } from '#lib/brand.js';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={logo} type="image/svg+xml" />
	<title>{pageTitle()}</title>
	<meta name="description" content={PRODUCT_DESCRIPTION} />
</svelte:head>
<Toaster position="bottom-right" closeButton={true} />
{@render children()}
