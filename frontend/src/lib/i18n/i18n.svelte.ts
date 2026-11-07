import { PRODUCT_SLUG } from '#lib/brand.js';
import { DEFAULT_LOCALE, isLocale, type Locale } from './locales.js';
import { en } from './messages/en.js';
import { id, type MessageKey } from './messages/id.js';

export type MessageParams = Record<string, string | number>;

const CATALOGS: Record<Locale, Record<MessageKey, string>> = { id, en };

/** Namespaced like the session key (`jari_session`). */
export const LOCALE_STORAGE_KEY = `${PRODUCT_SLUG}_locale`;

function readStoredLocale(): Locale | null {
	try {
		const value = localStorage.getItem(LOCALE_STORAGE_KEY);
		return isLocale(value) ? value : null;
	} catch {
		return null;
	}
}

function writeStoredLocale(locale: Locale): void {
	try {
		localStorage.setItem(LOCALE_STORAGE_KEY, locale);
	} catch {}
}

/** Replaces `{name}` placeholders; an unknown placeholder is left visible rather than dropped. */
export function interpolate(template: string, params?: MessageParams): string {
	if (!params) return template;
	return template.replace(/\{(\w+)\}/g, (match, name: string) =>
		name in params ? String(params[name]) : match
	);
}

class I18n {
	#locale = $state<Locale>(readStoredLocale() ?? DEFAULT_LOCALE);

	constructor() {
		this.#applyDocumentLang();
	}

	get locale(): Locale {
		return this.#locale;
	}

	setLocale(locale: Locale): void {
		if (!isLocale(locale)) return;
		this.#locale = locale;
		writeStoredLocale(locale);
		this.#applyDocumentLang();
	}

	t(key: MessageKey, params?: MessageParams): string {
		const template = CATALOGS[this.#locale][key] ?? CATALOGS[DEFAULT_LOCALE][key] ?? key;
		return interpolate(template, params);
	}

	#applyDocumentLang(): void {
		if (typeof document !== 'undefined') document.documentElement.lang = this.#locale;
	}
}

export const i18n = new I18n();

export function t(key: MessageKey, params?: MessageParams): string {
	return i18n.t(key, params);
}
