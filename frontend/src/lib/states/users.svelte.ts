import { showError } from '$lib/toasts';
import { GetUsers } from '$lib/wailsjs/go/backend/App';
import { models } from '$lib/wailsjs/go/models';
import { EventsOn } from '$lib/wailsjs/runtime/runtime';
import { appState } from './app-state.svelte';

let _all = $state<models.User[]>([]);
const _error = $state<string>();

export const users = {
	get all() {
		return _all;
	},
	get error() {
		return _error;
	},
	async load() {
		try {
			_all = await GetUsers();
			appState.setNeedsSetup(!_all?.length);
		} catch (err: unknown) {
			showError(String(err));
		}

		EventsOn('data-changed', (data) => {
			if (data.model === 'User') {
				_all = data.items;

				if (appState.needsSetup && _all.length) appState.setNeedsSetup(false);
			}
		});
	}
};
