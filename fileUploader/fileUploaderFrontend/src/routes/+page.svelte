<script lang="ts">
	import { preventDefault } from 'svelte/legacy';
	import { onDestroy, onMount } from 'svelte';

	let socket: WebSocket;

	let fileInput = $state<any>();

	type file = {
		id: string;
		name: string;
	};
	let filesArray = $state<file[]>([]);

	onMount(() => {
		socket = new WebSocket('ws://localhost:8080/ws');
		socket.onopen = () => {
			console.log('WebSocket connection established');
		};

		socket.onmessage = (event) => {
			const data = JSON.parse(event.data);
			const newFile: file = {
				id: data.id,
				name: data.name
			};
			filesArray = [...filesArray, newFile];
		};
	});

	onMount(async () => {
		try {
			const response = await fetch('/api/files');
			if (!response.ok) {
				throw new Error('Failed to fetch files');
			}
			const files: file[] = await response.json();
			if (files.length === 0) {
				console.log('No files found');
			} else {
				filesArray = files;
			}
		} catch (error) {
			console.error('Error fetching files:', error);
		}
	});

	onDestroy(() => {
		if (socket) {
			socket.close();
			console.log('WebSocket connection closed');
		}
	});
</script>

<section class="h-screen flex w-screen justify-center gap-5">
	<p>
		Go to <a href="/foo" class="bg-blue-600">/foo</a> to see the foo page. test
	</p>

	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 overflow-auto h-full mb-4">
		{#if filesArray.length === 0}
			<p class="text-gray-500">No files uploaded yet.</p>
		{/if}
		{#each filesArray as file (file.id)}
			<div class="bg-gray-100 p-4 rounded-lg shadow-md w-full">
				<p>File ID: {file.id}</p>
				<p>File Name: {file.name}</p>
				<p>
					File Date: {new Date(Number(file.id) / 1000000).toLocaleString('en-US', {
						hour12: false
					})}
				</p>
			</div>
		{/each}
	</div>
</section>
