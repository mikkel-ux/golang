<script lang="ts">
	import { onMount } from 'svelte';

	let socket: WebSocket;

	onMount(() => {
		socket = new WebSocket('/ws');
		socket.onopen = () => {
			console.log('WebSocket connection established');
		};

		socket.onmessage = (event) => {
			console.log('Message from server:', event.data);
		};
	});
</script>

<h1>foo</h1>
