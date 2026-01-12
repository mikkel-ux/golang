<script lang="ts">
	import { flip } from 'svelte/animate';
	let box1 = $state<string[]>(['item 1', 'item 2', 'item 3']);
	let previewItems = $state<string[]>([]);
	let draggedItemIndex: number = $state(-1);
	let draggedOverItemIndex: number = $state(-1);

	const preview = $derived(() => {
		if (
			draggedItemIndex === -1 ||
			draggedOverItemIndex === -1 ||
			draggedItemIndex === draggedOverItemIndex
		) {
			return box1;
		}
		const updated = [...box1];
		const [moved] = updated.splice(draggedItemIndex, 1);
		updated.splice(draggedOverItemIndex, 0, moved);
		return updated;
	});

	function handleDragStart(event: DragEvent | null, item: string, index: number) {
		if (!event?.dataTransfer) return;
		event.dataTransfer.setData('text/plain', item);
		draggedItemIndex = index;
	}

	/* function createTempDiv(event: DragEvent | null, index: number) {
		if (!event) return;
		const original = event.currentTarget as HTMLElement;

		const ghost = original.cloneNode(true) as HTMLElement;

		ghost.style.position = 'absolute';
		ghost.style.top = '-9999px';
		ghost.style.left = '-9999px';
		ghost.style.opacity = '0.5';

		document.body.appendChild(ghost);

		event.dataTransfer?.setDragImage(ghost, 30, 30);

		requestAnimationFrame(() => {
			document.body.removeChild(ghost);
		});
	} */

	function swap(indexA: number, indexB: number) {
		const temp = box1[indexA];
		box1[indexA] = box1[indexB];
		box1[indexB] = temp;
	}

	function handleDragOver(event: DragEvent | null, index: number) {
		if (!event) return;
		event.preventDefault();
		draggedOverItemIndex = index;
	}

	function handleDrop(event: DragEvent | null) {
		if (!event) return;
		if (preview() !== box1) {
			box1 = preview();
		}
		draggedItemIndex = -1;
		draggedOverItemIndex = -1;
	}
</script>

<!-- <div>
	<h1 class="text-white"><a href="/">home</a></h1>
</div> -->

<section class="w-full h-full flex flex-row gap-4">
	<div class="drag-grid flex-1 border-4" role="none" ondrop={handleDrop}>
		{#each preview() as item, index (item)}
			<div
				class="p-4 m-2 bg-gray-700 text-white rounded"
				draggable="true"
				ondragover={(event) => handleDragOver(event, index)}
				ondragstart={(event) => handleDragStart(event, item, index)}
				role="none"
				animate:flip
			>
				<!-- {#if index !== DraggedItemIndex}
					{item}
				{/if} -->
				{item}
			</div>
		{/each}
	</div>
</section>

<!-- <style>
	.hidden {
		display: none;
	}
</style> -->
