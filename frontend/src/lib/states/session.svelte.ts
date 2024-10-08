import { models } from '$lib/wailsjs/go/models';
import { EventsOn } from '$lib/wailsjs/runtime/runtime';

let _session = $state<models.Session | null>(null);

export const user = {
	get session() {
		return _session;
	},
	setSession(val: models.Session | null) {
		_session = val;
	},
	onSessionChanged() {
		EventsOn('session-changed', (s) => {
			_session = s;
			console.log('current session', s);
		});
	}
};
