<script lang="ts">
	import { preventDefault } from 'svelte/legacy';

	let number = $state(0);
	let fileInput = $state<any>();

	/* async function increment() {
		try {
			const response = await fetch('/api/hello');

			if (!response.ok) {
				throw new Error('Server responded with an error.');
			}
			const data = await response.json();
			console.log('Response from server:', data.message);
		} catch (error) {
			console.error('You failed to get the message:', error);
		}
	} */

	async function handleSubmit() {
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
</script>

<p>
	Go to <a href="/foo" class="bg-blue-600">/foo</a> to see the foo page. test
</p>
<br />
<section class="h-screen flex justify-center items-center">
	<div class="bg-gray-100 p-4 rounded-lg shadow-md">
		<form
			class="flex flex-col gap-2 justify-center items-center text-center"
			id="form"
			enctype="multipart/form-data"
			method="post"
			onsubmit={preventDefault(handleSubmit)}
		>
			<input type="file" name="file" multiple bind:files={fileInput} />
			<button class="hover:bg-cyan-800 bg-blue-600 p-5" type="submit">Submit</button>
		</form>
	</div>
</section>
