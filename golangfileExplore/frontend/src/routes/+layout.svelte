<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	import { Close, Minimize, MaximizeUnmaximize } from '$lib/wailsjs/go/main/App';
	let { children } = $props();
	let activeTab = $state(1);
	let tabs = [{ title: 'thing 1', id: 1 }];
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<nav
	class="h-11 w-full bg-black text-white flex items-center justify-between p-4"
	style="--wails-draggable:drag"
>
	<div
		class="flex items-center max-w-full overflow-x-auto space-x-2 min-w-[245px] mr-5"
		style="--wails-draggable:no-drag"
	>
		{#each tabs as tab}
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div
				class="flex items-cente rounded-md px-3 py-1 grow shrink min-w-[120px] max-w-60 hover:bg-gray-600"
				class:bg-gray-800={tab.id === activeTab}
				class:bg-gray-700={tab.id !== activeTab}
				class:hover:bg-gray-900={tab.id === activeTab}
				onclick={() => (activeTab = tab.id)}
			>
				<span class="truncate flex-1">
					{tab.title}
				</span>
				<span class="ml-2 cursor-pointer shrink-0">x</span>
			</div>
		{/each}
	</div>

	<div class="flex-1" style="--wails-draggable:drag"></div>

	<div class="flex space-x-2 justify-self-end-safe" style="--wails-draggable:no-drag">
		<button
			onclick={() => Minimize()}
			class="w-4 h-4 bg-yellow-500 rounded-full hover:bg-yellow-600"
			title="Minimize"
		></button>
		<button
			onclick={() => MaximizeUnmaximize()}
			class="w-4 h-4 bg-green-500 rounded-full hover:bg-green-600"
			title="Maximize/Unmaximize"
		></button>
		<button
			onclick={() => Close()}
			class="w-4 h-4 bg-red-500 rounded-full hover:bg-red-600"
			title="Close"
		></button>
	</div>
</nav>

{@render children()}

<style>
	nav ::-webkit-scrollbar {
		height: 6px;
	}
	nav ::-webkit-scrollbar-thumb {
		background: rgba(255, 255, 255, 0.2);
		border-radius: 3px;
	}
	nav ::-webkit-scrollbar-thumb:hover {
		background: rgba(255, 255, 255, 0.4);
	}

	nav ::-webkit-scrollbar-track {
		background: rgba(255, 255, 255, 0.1);
		border-radius: 3px;
	}
</style>
