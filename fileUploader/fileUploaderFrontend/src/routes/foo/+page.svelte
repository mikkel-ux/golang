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
			if (data.clientsCount !== null) {
				return;
			} else if (data.fileWasRemoved !== '') {
				const test = filesArray.find((file) => file.id === data.fileWasRemoved);
				if (!test) {
					console.log('File not found in the array:', data.fileWasRemoved);
					return;
				}

				filesArray = filesArray.filter((file) => file.id !== data.fileWasRemoved);
				return;
			} else {
				const newFile: file = {
					id: data.file.id,
					name: data.file.name
				};
				filesArray = [...filesArray, newFile];
				return;
			}
		};
	});

	onMount(async () => {
		try {
			const response = await fetch('/api/upload');
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

	const downloadFile = async (fileId: string, fileName: string) => {
		try {
			const response = await fetch(`/api/upload/${fileId}`, {
				method: 'GET'
			});
			if (!response.ok) {
				throw new Error('Failed to download file');
			} else {
				const blob = await response.blob();
				const url = window.URL.createObjectURL(blob);
				console.log(url);

				const a = document.createElement('a');
				a.href = url;
				a.download = fileName;
				document.body.appendChild(a);
				a.click();
				console.log(a);
				document.body.removeChild(a);

				window.URL.revokeObjectURL(url);
			}
		} catch (error) {
			console.error('Error downloading file:', error);
		}
	};
</script>

<section class="h-screen grid grid-rows-[auto_1fr] grid-cols-1 gap-4 p-4 m-4">
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
	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 overflow-auto h-full mb-4">
		{#if filesArray.length === 0}
			<p class="text-gray-500">No files uploaded yet.</p>
		{/if}
		{#each filesArray as file (file.id)}
			<div class="bg-gray-100 p-4 rounded-lg shadow-md w-full">
				<button
					class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
					onclick={() => downloadFile(file.id, file.name)}
				>
					Download
				</button>
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
