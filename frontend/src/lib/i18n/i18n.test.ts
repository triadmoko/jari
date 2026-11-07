import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { GLOSSARY, TERM_KEYS } from './glossary.js';
import { LOCALE_STORAGE_KEY, i18n, interpolate, t } from './i18n.svelte.js';
import { DEFAULT_LOCALE, LOCALES } from './locales.js';
import { en } from './messages/en.js';
import { id } from './messages/id.js';

function fakeStorage(initial: Record<string, string> = {}): Storage {
	const data = new Map(Object.entries(initial));
	return {
		get length() {
			return data.size;
		},
		clear: () => data.clear(),
		getItem: (k) => data.get(k) ?? null,
		key: (i) => [...data.keys()][i] ?? null,
		removeItem: (k) => void data.delete(k),
		setItem: (k, v) => void data.set(k, v)
	};
}

const placeholders = (text: string) => [...text.matchAll(/\{(\w+)\}/g)].map((m) => m[1]).sort();

describe('catalogs', () => {
	it('have the same keys in every locale', () => {
		expect(Object.keys(en).sort()).toEqual(Object.keys(id).sort());
	});

	it('use the same placeholders in every locale', () => {
		for (const key of Object.keys(id) as (keyof typeof id)[]) {
			expect(placeholders(en[key]), key).toEqual(placeholders(id[key]));
		}
	});

	it('have no empty messages', () => {
		for (const catalog of [id, en]) {
			for (const [key, value] of Object.entries(catalog)) expect(value.trim(), key).not.toBe('');
		}
	});
});

describe('glossary', () => {
	it('defines every term in every locale', () => {
		for (const locale of LOCALES) {
			for (const key of TERM_KEYS) {
				const term = GLOSSARY[locale][key];
				expect(term.one.trim(), `${locale}.${key}`).not.toBe('');
				expect(term.many.trim(), `${locale}.${key}`).not.toBe('');
			}
		}
	});

	it('words work items and spaces as the PRD terminology requires (EN)', () => {
		expect(GLOSSARY.en.workItem.one).toBe('work item');
		expect(GLOSSARY.en.space.many).toBe('spaces');
	});
});

describe('interpolate', () => {
	it('replaces placeholders and leaves unknown ones visible', () => {
		expect(interpolate('Hi {name}, {n} new', { name: 'Ana', n: 3 })).toBe('Hi Ana, 3 new');
		expect(interpolate('Hi {name}', {})).toBe('Hi {name}');
		expect(interpolate('Hi {name}')).toBe('Hi {name}');
	});
});

describe('i18n store', () => {
	beforeEach(() => {
		vi.stubGlobal('localStorage', fakeStorage());
		i18n.setLocale(DEFAULT_LOCALE);
	});
	afterEach(() => {
		i18n.setLocale(DEFAULT_LOCALE);
		vi.unstubAllGlobals();
	});

	it('translates in the active locale and switches without a reload', () => {
		expect(t('action.login')).toBe(id['action.login']);
		i18n.setLocale('en');
		expect(i18n.locale).toBe('en');
		expect(t('action.login')).toBe(en['action.login']);
	});

	it('fills params into the message', () => {
		i18n.setLocale('en');
		expect(t('validation.usernameTooShort', { min: 3 })).toBe(
			'Username must be at least 3 characters.'
		);
	});

	it('persists the preference and ignores unsupported locales', () => {
		i18n.setLocale('en');
		expect(localStorage.getItem(LOCALE_STORAGE_KEY)).toBe('en');

		// @ts-expect-error deliberately outside the Locale union
		i18n.setLocale('fr');
		expect(i18n.locale).toBe('en');
	});

	it('survives storage that throws', () => {
		vi.stubGlobal('localStorage', {
			getItem: () => {
				throw new Error('blocked');
			},
			setItem: () => {
				throw new Error('blocked');
			}
		});
		expect(() => i18n.setLocale('en')).not.toThrow();
		expect(i18n.locale).toBe('en');
	});
});
