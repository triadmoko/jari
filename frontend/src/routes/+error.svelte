<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
</script>

<main class="flex min-h-svh flex-col items-center justify-center gap-2 p-6 text-center">
	<h1 class="text-4xl font-semibold">{page.status}</h1>
	<p class="text-muted-foreground">
		{page.status === 404
			? 'Halaman tidak ditemukan.'
			: (page.error?.message ?? 'Terjadi kesalahan.')}
	</p>
	<a href={resolve('/(guest)')} class="mt-2 underline underline-offset-4">Kembali ke beranda</a>
</main>
