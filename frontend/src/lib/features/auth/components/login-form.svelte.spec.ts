import { page, userEvent } from 'vitest/browser';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { render } from 'vitest-browser-svelte';
import { fail, ok } from '#lib/api/result.js';
import LoginForm from './login-form.svelte';

const { goto, login } = vi.hoisted(() => ({ goto: vi.fn(), login: vi.fn() }));

vi.mock('$app/navigation', () => ({ goto }));
vi.mock('../infrastructure/api-auth.client.js', () => ({
	apiAuthRepository: { login, register: vi.fn() }
}));

const user = { id: '1', email: 'a@b.co', username: 'abc', role: 'user', name: 'A' };

async function submit(email: string, password: string) {
	await userEvent.fill(page.getByLabelText('Email'), email);
	await userEvent.fill(page.getByLabelText('Kata sandi'), password);
	await userEvent.click(page.getByRole('button', { name: 'Masuk' }));
}

describe('LoginForm', () => {
	beforeEach(() => {
		goto.mockReset();
		login.mockReset();
		localStorage.clear();
	});

	it('shows inline validation errors and does not call the API', async () => {
		render(LoginForm);
		await submit('bukan-email', '');

		await expect.element(page.getByText('Format email tidak valid.')).toBeInTheDocument();
		await expect.element(page.getByText('Kata sandi wajib diisi.')).toBeInTheDocument();
		expect(login).not.toHaveBeenCalled();
		expect(goto).not.toHaveBeenCalled();
	});

	it('shows the API error for wrong credentials', async () => {
		login.mockResolvedValue(fail('INVALID_CREDENTIALS', 'Email atau kata sandi salah.'));
		render(LoginForm);
		await submit('a@b.co', 'salah-banget');

		await expect.element(page.getByText('Email atau kata sandi salah.')).toBeInTheDocument();
		expect(goto).not.toHaveBeenCalled();
		expect(localStorage.getItem('jari_session')).toBeNull();
	});

	it('stores the session and navigates to the safe redirect target', async () => {
		login.mockResolvedValue(ok({ token: 'tok', user }));
		render(LoginForm, { redirectTo: '/profile' });
		await submit('a@b.co', 'secret');

		await vi.waitFor(() => expect(goto).toHaveBeenCalledWith('/profile'));
		expect(JSON.parse(localStorage.getItem('jari_session') ?? '{}')).toMatchObject({
			token: 'tok'
		});
	});

	it('announces a fresh registration', async () => {
		render(LoginForm, { registered: true });
		await expect
			.element(page.getByText('Akun berhasil dibuat. Silakan masuk.'))
			.toBeInTheDocument();
	});
});
