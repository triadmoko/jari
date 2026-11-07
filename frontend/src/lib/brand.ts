/**
 * Single source of the product identity on the frontend. Change it here when the final
 * product name is decided. The backend counterpart is `pkg/brand`.
 */
export const PRODUCT_NAME = 'Jari';

export const PRODUCT_SLUG = 'jari';

/** `Masuk · Jari`; the bare product name when there is no section. */
export function pageTitle(section?: string): string {
	return section ? `${section} · ${PRODUCT_NAME}` : PRODUCT_NAME;
}
