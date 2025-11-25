<script lang="ts">
	import { onMount } from 'svelte';
	import { DoSomething } from '$lib/wailsjs/go/test/SampleTest';

	function doGreeting(name: string) {
		DoSomething(name).then((result) => {
			alert(result);
		});
	}

	onMount(() => {
		doGreeting('World');
	});
</script>
