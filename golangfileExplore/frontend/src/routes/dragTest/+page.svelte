<script lang="ts">
	type Item = {
		id: number;
		name: string;
	};
	import { flip } from 'svelte/animate';
	let box1 = $state<Item[]>([
		{ id: 1, name: 'item 1' },
		{ id: 2, name: 'item 2' },
		{ id: 3, name: 'item 3' }
	]);
	// svelte-ignore state_referenced_locally
	let preview = $state<Item[]>(box1);
	let draggedId: number = $state(0);

	function handleDragStart(event: DragEvent | null, id: number) {
		if (!event?.dataTransfer) return;
		event.dataTransfer.setData('text/plain', '');
		draggedId = id;
		preview = box1;
	}

	function handleDragOver(event: DragEvent | null, overId: number) {
		if (!event) return;
		event.preventDefault();
		if (draggedId === 0 || draggedId === overId) return;

		const from = preview.findIndex((item) => item.id === draggedId);
		const to = preview.findIndex((item) => item.id === overId);

		if (from === -1 || to === -1) return;

		const updated = [...preview];
		const [moved] = updated.splice(from, 1);
		updated.splice(to, 0, moved);
		preview = updated;
	}

	function handleDrop(event: DragEvent | null) {
		if (!event) return;
		box1 = preview;
		draggedId = 0;
	}
</script>

<!-- <div>
	<h1 class="text-white"><a href="/">home</a></h1>
</div> -->

<section class="w-full h-full flex flex-row gap-4">
	<div class="drag-grid flex-1 border-4" role="none" ondrop={handleDrop}>
		{#each preview as item (item)}
			<div
				class="p-4 m-2 bg-gray-700 text-white rounded"
				draggable="true"
				ondragover={(event) => handleDragOver(event, item.id)}
				ondragstart={(event) => handleDragStart(event, item.id)}
				role="none"
				animate:flip={{ duration: 150 }}
			>
				{item.name}
			</div>
		{/each}
	</div>
</section>
