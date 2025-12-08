<script lang="ts">
	import { onMount } from 'svelte';
	import { DistanceTests } from '$lib/wailsjs/go/files/DistanceTest';
	import { Close, Maximize, Minimize, Unmaximize } from '$lib/wailsjs/go/main/App';
	let inputText: string = $state('');
	let results: string[] = $state([]);
	let isMaximized: boolean = $state(false);

	onMount(() => {
		DistanceTests('example').then((res) => {
			results = res;
		});
	});

	$effect(() => {
		DistanceTests(inputText).then((res) => {
			results = res;
		});
	});

	const maximize = () => {
		if (isMaximized) {
			Unmaximize();
		} else {
			Maximize();
		}
		isMaximized = !isMaximized;
	};
</script>

<div
	class="h-8 w-full bg-black text-white flex items-center justify-between px-4"
	style="--wails-draggable:drag"
>
	<span>golangfileExplore</span>
	<div class="flex space-x-2">
		<button onclick={() => Close()} class="w-4 h-4 bg-red-500 rounded-full" title="Close"></button>
		<button onclick={() => Minimize()} class="w-4 h-4 bg-yellow-500 rounded-full" title="Minimize"
		></button>
		<button
			onclick={() => maximize()}
			class="w-4 h-4 bg-green-500 rounded-full"
			title="Maximize/Unmaximize"
		></button>
	</div>
</div>

<input type="text" bind:value={inputText} />

{#each results as result}
	<p class="text-white">{result}</p>
{/each}
