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
	// import {FormControl} from '$lib/components/ui/form'
	import Input from '$lib/components/ui/input/input.svelte';
	import { Label } from '$lib/components/ui/label';
	import { showError } from '$lib/toasts';
	import { CheckIfNeedsSetup, RegisterUserCommand } from '$lib/wailsjs/go/backend/App';

	let uname = $state('');
	let passw = $state('');
	let passw2 = $state('');

	async function onRegister(e: SubmitEvent) {
		e.preventDefault();

		if (passw != passw2) return showError('passwords-dont-match');

		await RegisterUserCommand(uname.trim(), passw.trim()).catch(showError);
		if (!(await CheckIfNeedsSetup())) goto('/');
	}
</script>

<div class="flex h-svh w-full items-center justify-center">
	<form onsubmit={onRegister} class="w-[90%] max-w-md">
		<Card>
			<CardHeader>
				<CardTitle class="flex flex-col text-2xl">
					<span>Welcome to</span>

					<span>Cebu Technological University</span>
					<span>College of Engineering</span>
					<span>Class Scheduling System</span>
				</CardTitle>
				<CardDescription>Create an admin user to start using this app.</CardDescription>
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
				<div class="grid gap-2">
					<Label for="password2">Confirm Password</Label>
					<Input id="password2" type="password" required bind:value={passw2} />
				</div>
			</CardContent>
			<CardFooter>
				<Button type="submit" class="w-full">Create</Button>
			</CardFooter>
		</Card>
	</form>
</div>
