<script lang="ts">
	import { BookCopy, CalendarCheck, House, UsersRound, Warehouse, Workflow } from 'lucide-svelte';

	import { page } from '$app/stores';
	import ctuLogo from '$lib/images/ctu.png';
	import { Settings } from 'lucide-svelte';
	import type { Snippet } from 'svelte';

	type Link = {
		to: string;
		label: string;
		icon: Snippet;
	};

	const links: Link[] = [
		{
			to: '/',
			label: 'Dashboard',
			icon: dashboardIcon
		},
		{
			to: '/schedules',
			label: 'Schedules',
			icon: schedulesIcon
		},
		{
			to: '/sections',
			label: 'Sections',
			icon: sectionsIcon
		},
		{
			to: '/subjects',
			label: 'Subjects',
			icon: subjectsIcon
		},
		{
			to: '/rooms',
			label: 'Rooms',
			icon: roomsIcon
		},
		{
			to: '/instructors',
			label: 'Instructors',
			icon: instructorsIcon
		}
	];
</script>

{#snippet dashboardIcon()}
	<House class="h-4 w-4" />
{/snippet}

{#snippet schedulesIcon()}
	<CalendarCheck class="h-4 w-4" />
{/snippet}

{#snippet sectionsIcon()}
	<Workflow class="h-4 w-4" />
{/snippet}

{#snippet subjectsIcon()}
	<BookCopy class="h-4 w-4" />
{/snippet}

{#snippet roomsIcon()}
	<Warehouse class="h-4 w-4" />
{/snippet}

{#snippet instructorsIcon()}
	<UsersRound class="h-4 w-4" />
{/snippet}

<div class="flex h-14 items-center border-b px-4">
	<a href="/" class="flex w-full items-center gap-2 font-semibold">
		<img src={ctuLogo} alt="" class="h-10 w-10" />
		<span class="flex flex-col">
			<span class="text-lg font-bold leading-4">CTU COE CSS</span>
			<span class="text-xs font-light">Class Scheduling System</span>
		</span>
	</a>
</div>
<div class="flex-1">
	<nav class="grid items-start px-2 text-sm font-medium">
		{#each links as link}
			<a
				href={link.to}
				class:bg-muted={$page.url.pathname === link.to}
				class:text-primary={$page.url.pathname === link.to}
				class="text-muted-foreground hover:text-primary flex items-center gap-3 rounded-lg px-3 py-2 transition-all"
			>
				{@render link.icon()}
				{link.label}
			</a>
		{/each}
	</nav>
</div>
<div class="mt-auto pb-2">
	<nav class="grid items-start px-2 text-sm font-medium">
		<a
			href="/settings"
			class:bg-muted={$page.url.pathname === '/settings'}
			class:text-primary={$page.url.pathname === '/settings'}
			class="text-muted-foreground hover:text-primary flex items-center gap-3 rounded-lg px-3 py-2 transition-all"
		>
			<Settings class="h-4 w-4" />
			Settings
		</a>
	</nav>
</div>
