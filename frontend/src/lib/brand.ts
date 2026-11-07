/**
 * Single source of the product identity on the frontend. Change it here when the final
 * product name is decided. The backend counterpart is `pkg/brand`.
 */
export const PRODUCT_NAME = 'Jari';

/** Lowercase identifier for storage keys and the like. */
export const PRODUCT_SLUG = 'jari';

export const PRODUCT_DESCRIPTION = 'Platform manajemen kerja dan proyek open-source.';

/** `Masuk · Jari`; the bare product name when there is no section. */
export function pageTitle(section?: string): string {
	return section ? `${section} · ${PRODUCT_NAME}` : PRODUCT_NAME;
}
