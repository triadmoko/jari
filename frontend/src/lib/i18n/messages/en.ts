import { GLOSSARY } from '../glossary.js';
import type { MessageKey } from './id.js';

const term = GLOSSARY.en;

/** English catalog. `Record<MessageKey, string>` makes a missing translation a compile error. */
export const en: Record<MessageKey, string> = {
	'brand.description': 'An open-source work and project management platform.',

	'action.login': 'Sign in',
	'action.register': 'Sign up',
	'action.logout': 'Sign out',
	'action.save': 'Save',

	'nav.profile': 'Profile',
	'nav.language': 'Language',

	'field.name': 'Name',
	'field.username': 'Username',
	'field.email': 'Email',
	'field.password': 'Password',

	'validation.required': '{field} is required.',
	'validation.invalidEmail': 'The email format is invalid.',
	'validation.usernameTooShort': 'Username must be at least {min} characters.',
	'validation.usernameTooLong': 'Username must be at most {max} characters.',
	'validation.passwordTooShort': 'Password must be at least {min} characters.',

	'api.validation': 'Please check your input.',
	'api.unauthorized': 'Your session is not valid. Please sign in again.',
	'api.invalidCredentials': 'Incorrect email or password.',
	'api.forbidden': 'You do not have permission to perform this action.',
	'api.network': 'Unable to reach the server. Please try again shortly.',
	'api.unknown': 'Something went wrong. Please try again shortly.',

	'page.home': 'Home',
	'home.greeting': 'Hello, {name}',
	'home.spacesPlaceholder': `Your ${term.space.many} will appear here.`,
	'home.viewProfile': 'View profile',
	'home.welcome': 'Welcome to {product}',
	'home.welcomeHint': 'Sign in or sign up to get started.',

	'auth.login.title': 'Sign in',
	'auth.login.description': 'Sign in with your email and password.',
	'auth.login.noAccount': "Don't have an account?",
	'auth.register.title': 'Sign up',
	'auth.register.description': 'Create a new account to start using {product}.',
	'auth.register.haveAccount': 'Already have an account?',
	'auth.registered': 'Your account has been created. Please sign in.',
	'auth.sessionExpired': 'Your session has expired. Please sign in again.',

	'profile.title': 'Profile',
	'profile.description': 'View and change your display name.',
	'profile.updated': 'Profile updated successfully.',

	'error.notFound': 'Page not found.',
	'error.generic': 'Something went wrong.',
	'error.backHome': 'Back to home'
};
