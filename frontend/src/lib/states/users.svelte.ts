import { GetUsers } from '$lib/wailsjs/go/backend/App';
import { models } from '$lib/wailsjs/go/models';
import { EventsOn } from '$lib/wailsjs/runtime/runtime';

let _all = $state<models.User[]>([]);
let _error = $state<string>();

export const users = {
	get all() {
		return _all;
	},
	get error() {
		return _error;
	},
	load() {
		GetUsers()
			.catch((err) => (_error = err))
			.then((users) => (_all = users));

		EventsOn('data-changed', (data) => {
			console.log(data);

			if (data.model === 'User') _all = data.items;
		});
	}
};
