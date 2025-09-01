<script lang="ts">
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

<h1>Sign Up</h1>
<form onsubmit={createUser}>
	<label>
		Username:
		<input type="text" bind:value={userName} required />
	</label>
	<label>
		Password:
		<input type="password" bind:value={password} required />
	</label>
	<button type="submit">Sign Up</button>
</form>
