<script lang="ts">
	import { Button } from '#lib/components/ui/button/index.js';
	import { LOCALES, LOCALE_NAMES, i18n, t } from '#lib/i18n/index.js';
</script>

<div role="group" aria-label={t('nav.language')} class="flex items-center gap-1">
	{#each LOCALES as locale (locale)}
		<Button
			variant={i18n.locale === locale ? 'secondary' : 'ghost'}
			size="sm"
			aria-pressed={i18n.locale === locale}
			title={LOCALE_NAMES[locale]}
			lang={locale}
			onclick={() => i18n.setLocale(locale)}
		>
			{locale.toUpperCase()}
		</Button>
	{/each}
</div>
