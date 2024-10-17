let _initialLoading = $state(true);
let _needsSetup = $state(false);

export const appState = {
	get initialLoading() {
		return _initialLoading;
	},
	get needsSetup() {
		return _needsSetup;
	},
	setInitialLoading(val: boolean) {
		_initialLoading = val;
	},
	setNeedsSetup(val: boolean) {
		console.log('setNeedsSetup:', val);

		_needsSetup = val;
	}
};
