<script lang="ts">
	import { onMount } from 'svelte';

	let userName = $state<string>('');
	let password = $state<string>('');
	let token = $state<string>('');

	const login = async (e: Event) => {
		e.preventDefault();

		const user = {
			userName: userName,
			password: password
		};

		const response = await fetch('/api/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(user)
		});

		if (response.ok) {
			const result = await response.json();
			document.cookie = `token=${result.token}; expires=${new Date(result.expiresAt * 1000).toUTCString()};`;
			console.log('Login successful:', result);
		} else {
			console.error('Login failed');
		}

		userName = '';
		password = '';
	};

	/* const createUser = async () => {
		const user = {
			userName: userName,
			password: password
		};
		try {
			const response = await fetch('/api/user', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(user)
			});

			if (response.ok) {
				const result = await response.json();
				console.log('User created successfully:', result);
			} else {
				console.error('Failed to create user');
			}
		} catch (error) {
			console.error('Error creating user:', error);
		}

		userName = '';
		password = '';
	}; */
</script>

<section class="h-screen p-4 m-4 flex justify-center items-center">
	<div class="bg-white p-6 rounded-lg shadow-lg">
		<form onsubmit={login}>
			<input type="text" bind:value={userName} placeholder="Username" required />
			<input type="password" bind:value={password} placeholder="Password" required />
			<button type="submit">Login</button> <br />
			<div class="grid grid-cols-3 gap-2 mt-4">
				<!-- <button type="button" onclick={createUser} class="ml-2 bg-blue-500 hover:bg-blue-700"
					>Sign Up</button
				> -->
				<a href="/signup">Sign Up</a>
				<!-- <button type="button" onclick={tokenTest} class="ml-2 bg-blue-500 hover:bg-blue-700"
					>Token Test</button
				>
				<button
					type="button"
					onclick={() => (token = '')}
					class="ml-2 bg-blue-500 hover:bg-blue-700">Clear Token</button
				> -->
			</div>
		</form>
	</div>
</section>
