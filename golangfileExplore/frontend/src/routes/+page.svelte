<script lang="ts">
	import { onMount } from 'svelte';
	import { DistanceTests } from '$lib/wailsjs/go/files/DistanceTest';
	let inputText: string = $state('');
	let results: string[] = $state([]);

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
</script>

<a class="text-white" href="/test">test page</a> <br />
<a class="text-white" href="/dragTest">drag test page</a>
<h1 class="text-white"><a href="/dragTestAction">action test</a></h1>
<br />
<input type="text" bind:value={inputText} />

{#each results as result}
	<p class="text-white">{result}</p>
{/each}
