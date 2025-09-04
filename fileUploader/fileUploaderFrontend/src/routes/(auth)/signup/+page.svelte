<script lang="ts">
	import { goto } from '$app/navigation';
	let userName = $state<string>('');
	let password = $state<string>('');

	const createUser = async () => {
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
				goto('/login');
			} else {
				console.error('Failed to create user');
			}
		} catch (error) {
			console.error('Error creating user:', error);
		}

		userName = '';
		password = '';
	};
</script>

<section class="h-screen p-4 m-4 flex justify-center items-center">
	<div class="bg-white p-6 rounded-lg shadow-lg">
		<form onsubmit={createUser}>
			<input type="text" bind:value={userName} placeholder="Username" required />
			<input type="password" bind:value={password} placeholder="Password" required />
			<button type="submit">Sign Up</button> <br />
			<div class="grid grid-cols-3 gap-2 mt-4">
				<a href="/login">login</a>
			</div>
		</form>
	</div>
</section>
