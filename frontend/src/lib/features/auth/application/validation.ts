/** Pragmatic check; the backend remains the source of truth. */
export function isValidEmail(value: string): boolean {
	return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value.trim());
}
