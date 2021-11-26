import adapter from '@sveltejs/adapter-auto';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: preprocess(),

	kit: {
		adapter: adapter(),

		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',

		// disable for now, since I am still not quite sure how this ideally works
		ssr: false,

		vite: {
			server: {
				proxy: {
					'/api': 'http://localhost:4300'
				}
			}
		}
	}
};

export default config;
