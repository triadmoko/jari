/** Locales the UI is translated into. Keep in sync with `constants.Lang*` in the backend. */
export const LOCALES = ['id', 'en'] as const;

export type Locale = (typeof LOCALES)[number];

export const DEFAULT_LOCALE: Locale = 'id';

/** Language names in their own language, so they stay readable whatever the UI language is. */
export const LOCALE_NAMES: Record<Locale, string> = {
	id: 'Indonesia',
	en: 'English'
};

export function isLocale(value: unknown): value is Locale {
	return typeof value === 'string' && (LOCALES as readonly string[]).includes(value);
}
