let _initialLoading = $state(true);

export const appState = {
	get initialLoading() {
		return _initialLoading;
	},
	setInitialLoading(val: boolean) {
		_initialLoading = val;
	}
};
