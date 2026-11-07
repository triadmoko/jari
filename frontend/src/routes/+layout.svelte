<script lang="ts">
	import './layout.css';
	import { Toaster } from '#lib/components/ui/sonner/index.js';
	import logo from '#lib/assets/logo.svg';
	import { pageTitle } from '#lib/brand.js';
	import { t } from '#lib/i18n/index.js';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={logo} type="image/svg+xml" />
	<title>{pageTitle()}</title>
	<meta name="description" content={t('brand.description')} />
</svelte:head>
<Toaster position="bottom-right" closeButton={true} />
{@render children()}
