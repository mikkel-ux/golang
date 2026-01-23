// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}

	// Custom DOM events (e.g. from drag-and-drop actions)
	namespace svelteHTML {
		interface HTMLAttributes<T> {
			'on:consider'?: (event: CustomEvent<T>) => void;
			'on:finalize'?: (event: CustomEvent<T>) => void;
		}
	}
}

export {};
