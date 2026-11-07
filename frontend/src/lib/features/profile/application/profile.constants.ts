import { t } from '#lib/i18n/index.js';

export const PROFILE_MESSAGES = {
	get nameRequired() {
		return t('validation.required', { field: t('field.name') });
	},
	get updated() {
		return t('profile.updated');
	}
};
