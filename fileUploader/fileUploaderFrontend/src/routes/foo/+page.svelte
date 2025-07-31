<script lang="ts">
	import { onMount, onDestroy } from 'svelte';

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
			console.log('Received data from server:', data);
			console.log('File ID:', data.id);
			console.log('File Name:', data.name);
			const date = new Date(data.id / 1000000);
			console.log('File Date:', date.toLocaleString());

			/* console.log('Message from server:', event.data); */
		};
	});

	async function handleSubmit(event: SubmitEvent) {
		event.preventDefault();
		const file = fileInput?.[0];
		let formData = new FormData();

		formData.append('file', file);

		try {
			const response = await fetch('/api/upload', {
				method: 'POST',
				body: formData
			});

			const result = await response.json();

			if (!response.ok) {
				throw new Error('Failed to upload file');
			}

			console.log('File uploaded successfully:', result.message);
		} catch (error) {
			console.error('Error uploading file:', error);
		}
	}

	onDestroy(() => {
		if (socket) {
			socket.close();
			console.log('WebSocket connection closed');
		}
	});
</script>

<section class="h-screen flex w-screen justify-center gap-5">
	<div class="bg-gray-100 p-4 rounded-lg shadow-md">
		<form
			class="flex flex-col gap-2 justify-center items-center text-center"
			id="form"
			enctype="multipart/form-data"
			method="post"
			onsubmit={handleSubmit}
		>
			<input type="file" name="file" multiple bind:files={fileInput} />
			<button class="hover:bg-cyan-800 bg-blue-600 p-5" type="submit">Submit</button>
		</form>
	</div>
</section>
