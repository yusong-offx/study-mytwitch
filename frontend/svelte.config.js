// import adapter from '@sveltejs/adapter-static';
import adapter from '@sveltejs/adapter-node';
// import adapter from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter()
	}
};

export default config;
