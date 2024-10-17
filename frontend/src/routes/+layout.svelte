<script lang="ts">
	import { goto } from '$app/navigation';
	import { Skeleton } from '$lib/components/ui/skeleton';
	import { Toaster } from '$lib/components/ui/sonner';
	import { appState } from '$lib/states/app-state.svelte';
	import { user } from '$lib/states/session.svelte';
	import { users } from '$lib/states/users.svelte';
	import { CheckIfNeedsSetup } from '$lib/wailsjs/go/backend/App';
	import { EventsOn } from '$lib/wailsjs/runtime/runtime';
	import { ModeWatcher } from 'mode-watcher';
	import { onMount } from 'svelte';
	import '../app.css';

	const { children } = $props();

	onMount(async () => {
		EventsOn('session-changed', (s) => {
			user.setSession(s);
		});

		EventsOn('data-changed', (e) => {
			console.log(e);
		});

		if (await CheckIfNeedsSetup()) goto('/setup');
		appState.setInitialLoading(false);

		users.load();
	});
</script>

<Toaster />
<ModeWatcher />
{#if appState.initialLoading}
	<div class="grid h-svh place-items-center">
		<div class="flex items-center space-x-4">
			<Skeleton class="h-12 w-12 rounded-full" />
			<div class="space-y-2">
				<Skeleton class="h-4 w-[250px]" />
				<Skeleton class="h-4 w-[200px]" />
			</div>
		</div>
	</div>
{:else}
	<div class="flex h-svh w-full">
		{@render children()}
	</div>
{/if}
