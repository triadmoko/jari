/**
 * Glossary of product terms, the single place that decides how a concept is worded in the UI
 * (PRD 00 › VISI-002). Message catalogs build sentences from these entries instead of typing
 * "issue", "project" and friends inline, so switching a term later is a one-line change.
 *
 * The API, JQL and URLs keep the legacy names (`issue`, `project`); only UI wording lives here.
 * Indonesian has no plural, so `one` and `many` are equal there.
 *
 * Open decision (PRD 00): the Indonesian wording of "space" — "ruang" vs keeping "proyek".
 */
import type { Locale } from './locales.js';

export type Term = { one: string; many: string };

export const TERM_KEYS = [
	'workItem',
	'workType',
	'space',
	'board',
	'backlog',
	'sprint',
	'epic',
	'workflow',
	'status'
] as const;

export type TermKey = (typeof TERM_KEYS)[number];

export const GLOSSARY: Record<Locale, Record<TermKey, Term>> = {
	id: {
		workItem: { one: 'item kerja', many: 'item kerja' },
		workType: { one: 'tipe kerja', many: 'tipe kerja' },
		space: { one: 'ruang', many: 'ruang' },
		board: { one: 'papan', many: 'papan' },
		backlog: { one: 'backlog', many: 'backlog' },
		sprint: { one: 'sprint', many: 'sprint' },
		epic: { one: 'epic', many: 'epic' },
		workflow: { one: 'alur kerja', many: 'alur kerja' },
		status: { one: 'status', many: 'status' }
	},
	en: {
		workItem: { one: 'work item', many: 'work items' },
		workType: { one: 'work type', many: 'work types' },
		space: { one: 'space', many: 'spaces' },
		board: { one: 'board', many: 'boards' },
		backlog: { one: 'backlog', many: 'backlogs' },
		sprint: { one: 'sprint', many: 'sprints' },
		epic: { one: 'epic', many: 'epics' },
		workflow: { one: 'workflow', many: 'workflows' },
		status: { one: 'status', many: 'statuses' }
	}
};

/** Upper-cases the first letter, for terms that start a sentence or a label. */
export function capitalize(text: string): string {
	return text.charAt(0).toLocaleUpperCase() + text.slice(1);
}
