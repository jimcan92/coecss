<script lang="ts">
	import { page } from '$app/stores';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import * as Select from '$lib/components/ui/select';
	import { users } from '$lib/states/users.svelte';
	import { DeleteUser, RegisterUserCommand } from '$lib/wailsjs/go/backend/App';
	import type { Selected } from 'bits-ui';
	import { Plus } from 'lucide-svelte';

	let uname = $state('');
	let pass = $state('');
	let role = $state<Selected<string>>({ value: '' });
	let updating = $state(false);

	async function onSave(e: SubmitEvent) {
		e.preventDefault();

		RegisterUserCommand(uname, pass, role.value).catch((err) => console.log(err));
	}

	const currentUser = $derived.by(() => {
		const username = $page.url.searchParams.get('username');
		return users.all.find((u) => u.username === username);
	});

	$effect(() => {
		if (currentUser) {
			uname = currentUser.username;
			pass = currentUser.password;
			role = { value: currentUser.role, label: currentUser.role };
			updating = false;
		} else {
			uname = '';
			pass = '';
			role = { value: '' };
			updating = true;
		}
	});

	async function toggleUpdating() {
		if (updating) {
			updating = false;
		} else {
			updating = true;
		}
	}

	function onUpdate() {}

	function deleteUser(id?: string) {
		if (id) DeleteUser(id).catch((err) => console.log(err));
	}
</script>

<div class="flex h-16 items-center p-4">
	<h4 class="text-3xl font-bold">Users</h4>
</div>
<div class="flex h-[calc(100%-4rem)] w-full">
	<div class="flex h-full w-full">
		<div class="flex h-full w-52 flex-col p-2">
			<h6 class="h-10 pb-2 text-xl">Registered Users</h6>
			<div class="flex h-[calc(100%-2.5rem)] rounded-lg border p-2">
				<ScrollArea class="h-full w-full">
					<ul class="flex flex-col gap-2">
						{#each users.all as user}
							<a
								href="/users?username={user.username}"
								class:bg-muted={$page.url.searchParams.get('username') === user.username}
								class:text-primary={$page.url.searchParams.get('username') === user.username}
								class="text-muted-foreground hover:bg-muted hover:text-primary flex w-full flex-col items-start rounded-lg border px-4 py-1"
							>
								<span class="leading-3">{user.username}</span>
								<span class="text-foreground/50 text-xs font-light">{user.role}</span>
							</a>
						{/each}
						<a
							href="/users"
							class:bg-muted={!$page.url.searchParams.get('username')}
							class:text-primary={!$page.url.searchParams.get('username')}
							class="text-muted-foreground hover:bg-muted hover:text-primary sticky inset-x-0 bottom-0 flex w-full items-center justify-center gap-2 rounded-lg border px-2 py-1"
						>
							<span>New User</span>
							<Plus class="h-4 w-4" />
						</a>
					</ul>
				</ScrollArea>
			</div>
		</div>
		<div class="bg-muted/50 h-full w-[calc(100%-13rem)] rounded-tl-xl p-4">
			<div class="flex h-full flex-col">
				<h6 class="h-14 pb-2 text-xl">
					{$page.url.searchParams.get('username') ?? 'New User'}
				</h6>
				<form onsubmit={onSave} class="flex h-[calc(100%-3.5rem)] flex-col gap-4">
					<div class="grid max-w-sm gap-2">
						<Label for="username" class={!updating ? 'text-muted-foreground' : ''}>Username</Label>
						<Input
							id="username"
							type="text"
							autofocus
							autocomplete="off"
							placeholder="Username"
							required
							disabled={!updating}
							bind:value={uname}
						/>
					</div>
					<div class="grid max-w-sm gap-2">
						<Label for="password" class={!updating ? 'text-muted-foreground' : ''}>Password</Label>
						<Input
							id="password"
							type="password"
							placeholder="Password"
							required
							disabled={!updating}
							bind:value={pass}
						/>
					</div>
					<div class="grid max-w-sm gap-2">
						<Label for="role" class={!updating ? 'text-muted-foreground' : ''}>Role</Label>
						<Select.Root
							onSelectedChange={(e) => {
								console.log(e);
							}}
							disabled={!updating}
							bind:selected={role}
						>
							<Select.Trigger id="role">
								<Select.Value placeholder="Select a role" />
							</Select.Trigger>
							<Select.Content>
								<Select.Group>
									<!-- <Select.Label>Fruits</Select.Label> -->
									<Select.Item value={'Student'} label={'Student'}>{'Student'}</Select.Item>
									<Select.Item value={'Instructor'} label={'Instructor'}>{'Instructor'}</Select.Item
									>
									<Select.Item value={'Admin'} label={'Admin'}>{'Admin'}</Select.Item>
								</Select.Group>
							</Select.Content>
							<Select.Input
								name="favoriteFruit"
								bind:value={role}
								onselect={(e) => {
									console.log(e.currentTarget.value);
								}}
							/>
						</Select.Root>
					</div>
					<div class="mt-auto flex items-center justify-end gap-2">
						{#if currentUser}
							<Button onclick={() => deleteUser(currentUser?.username)}>Delete</Button>
							<Button onclick={toggleUpdating}>Edit</Button>
							<Button onclick={onUpdate}>Update</Button>
						{:else}
							<Button type="submit">Save</Button>
						{/if}
					</div>
				</form>
			</div>
		</div>
	</div>
</div>
