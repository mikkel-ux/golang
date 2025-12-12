<script lang="ts">
	import { onMount } from 'svelte';
	import { NewTab } from '$lib/wailsjs/go/tabs/Manager';
	import type { tabs } from '$lib/wailsjs/go/models';

	let tabsManager = $state<tabs.Tab[]>([]);

	onMount(async () => {
		let newTab = await NewTab();
		console.log('New Tab Created:', newTab);
		tabsManager = [...tabsManager, newTab];
	});

	const closeTab = (tabId: string) => {
		tabsManager = tabsManager.filter((tab) => tab.ID !== tabId);
	};
</script>

<h1 class="text-white">Test Page</h1>
<h1 class="text-white"><a href="/">home</a></h1>
<button
	class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
	onclick={async () => {
		const tab = await NewTab();
		console.log('New Tab Created:', tab);
		tabsManager = [...tabsManager, tab];
	}}
>
	New Tab
</button>

<button
	onclick={() => console.log(tabsManager)}
	class="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-4 rounded">log</button
>
{#if tabsManager}
	{#each tabsManager as tab}
		<div class="mt-4 p-4 border border-gray-700 rounded">
			<h2 class="text-white font-bold">Tab ID: {tab.ID}</h2>
			<p class="text-white">Title: {tab.Title}</p>
			<button
				onclick={() => closeTab(tab.ID)}
				class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded">x</button
			>
		</div>
	{/each}
{/if}
