<script lang="ts">
	import { goto } from '$app/navigation';
	import { Toaster } from '$lib/components/ui/sonner';
	import { user } from '$lib/states/session.svelte';
	import { CheckIfNeedsSetup } from '$lib/wailsjs/go/backend/App';
	import { EventsOn } from '$lib/wailsjs/runtime/runtime';
	import { ModeWatcher } from 'mode-watcher';
	import { onMount } from 'svelte';
	import '../app.css';

	const { children } = $props();

	onMount(async () => {
		EventsOn('session-changed', (s) => {
			console.log('emitted event: ', s);

			user.setSession(s);
		});
		if (await CheckIfNeedsSetup()) goto('/setup');
	});
</script>

<Toaster />
<ModeWatcher />
<div class="flex h-svh w-full">
	{@render children()}
</div>
