<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { user } from '$lib/states/session.svelte';
	import { SessionTerminate } from '$lib/wailsjs/go/backend/App';

	async function logout() {
		await SessionTerminate();
		console.log(user.session);
	}
</script>

nice
<Button onclick={logout}>Logout</Button>
