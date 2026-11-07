import { readFileSync, readdirSync, statSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { parse } from 'svelte/compiler';
import { describe, expect, it } from 'vitest';

/**
 * Guards S0-13: user-visible text in the S0-09 pages must come from the i18n catalog, never be
 * typed into markup. Walks every Svelte file outside the shadcn primitives and flags literal
 * text nodes and literal text attributes that carry words.
 */

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const SKIP_DIRS = ['lib/components/ui', 'lib/hooks'];
/** Attributes whose value is read by people (screen readers included). */
const TEXT_ATTRIBUTES = new Set(['label', 'placeholder', 'title', 'alt', 'aria-label', 'content']);

function svelteFiles(dir: string): string[] {
	return readdirSync(dir).flatMap((entry) => {
		const full = path.join(dir, entry);
		const rel = path.relative(SRC, full).split(path.sep).join('/');
		if (SKIP_DIRS.some((skip) => rel === skip || rel.startsWith(`${skip}/`))) return [];
		if (statSync(full).isDirectory()) return svelteFiles(full);
		return full.endsWith('.svelte') ? [full] : [];
	});
}

/** A "word" is anything with a letter; punctuation, whitespace and `&nbsp;` do not count. */
const hasWords = (text: string) => /\p{L}/u.test(text.replace(/&nbsp;|&[a-z]+;/g, ' '));

type Node = {
	type: string;
	name?: string;
	data?: string;
	start: number;
	attributes?: {
		type: string;
		name: string;
		value: true | { type: string; data?: string } | { type: string; data?: string }[];
	}[];
	fragment?: { nodes: Node[] };
	children?: Node[];
	[key: string]: unknown;
};

function walk(node: unknown, visit: (n: Node) => void): void {
	if (Array.isArray(node)) return node.forEach((child) => walk(child, visit));
	if (typeof node !== 'object' || node === null) return;
	const n = node as Node;
	if (typeof n.type === 'string') visit(n);
	for (const [key, value] of Object.entries(n)) {
		if (key === 'attributes') continue;
		if (typeof value === 'object' && value !== null) walk(value, visit);
	}
}

function findings(file: string): string[] {
	const source = readFileSync(file, 'utf8');
	const ast = parse(source, { modern: true });
	const rel = path.relative(SRC, file);
	const found: string[] = [];
	const line = (offset: number) => source.slice(0, offset).split('\n').length;

	walk(ast.fragment, (node) => {
		if (node.type === 'Text' && hasWords(node.data ?? '')) {
			found.push(`${rel}:${line(node.start)} text "${(node.data ?? '').trim()}"`);
		}
		for (const attr of node.attributes ?? []) {
			if (attr.type !== 'Attribute' || !TEXT_ATTRIBUTES.has(attr.name)) continue;
			// `true` is a bare attribute and a single object is `attr={expression}`; both are fine.
			if (!Array.isArray(attr.value)) continue;
			for (const part of attr.value) {
				if (part.type === 'Text' && hasWords(part.data ?? '')) {
					found.push(`${rel}:${line(node.start)} ${attr.name}="${part.data}"`);
				}
			}
		}
	});
	return found;
}

describe('hard-coded UI strings', () => {
	const files = svelteFiles(SRC);

	it('finds the Svelte files to check', () => {
		expect(files.length).toBeGreaterThan(5);
	});

	it('leaves no literal text in markup outside the i18n catalog', () => {
		expect(files.flatMap(findings)).toEqual([]);
	});
});
