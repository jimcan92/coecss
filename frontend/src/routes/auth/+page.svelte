<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/button/button.svelte';
	import {
		Card,
		CardContent,
		CardDescription,
		CardFooter,
		CardHeader,
		CardTitle
	} from '$lib/components/ui/card';
	import Input from '$lib/components/ui/input/input.svelte';
	import { Label } from '$lib/components/ui/label';
	import { user } from '$lib/states/session.svelte';
	import { showError, showSuccess } from '$lib/toasts';
	import { LoginUserCommand } from '$lib/wailsjs/go/backend/App';

	let uname = $state('');
	let passw = $state('');

	async function onLogin(e: SubmitEvent) {
		e.preventDefault();

		await LoginUserCommand(uname.trim(), passw.trim())
			.catch(showError)
			.then((user) => {
				if (user) {
					showSuccess('Login Success', 'You have successfully logged in');
				}
			});
	}

	$effect(() => {
		if (user.session) goto('/');
	});
</script>

<div class="flex h-svh w-full items-center justify-center">
	<form onsubmit={onLogin} class="w-full max-w-sm">
		<Card>
			<CardHeader>
				<CardTitle class="text-2xl">Welcome back!</CardTitle>
				<CardDescription>Enter your username below to login to your account.</CardDescription>
			</CardHeader>
			<CardContent class="grid gap-4">
				<div class="grid gap-2">
					<Label for="text">Username</Label>
					<Input
						id="text"
						type="text"
						autocomplete="off"
						placeholder="Username"
						required
						autofocus
						bind:value={uname}
					/>
				</div>
				<div class="grid gap-2">
					<Label for="password">Password</Label>
					<Input id="password" type="password" required bind:value={passw} />
				</div>
			</CardContent>
			<CardFooter>
				<Button type="submit" class="w-full">Sign in</Button>
			</CardFooter>
		</Card>
	</form>
</div>
