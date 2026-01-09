<script lang="ts">
	import { flip } from 'svelte/animate';
	let box1 = $state<string[]>(['item 1', 'item 2', 'item 3']);
	let DraggedItem = $state<Element | null>(null);
	let DraggedItemIndex: number = $state(-1);
	let DraggedOverItemIndex: number = $state(-1);

	function handleDragStart(event: DragEvent | null, item: string, element: Element, index: number) {
		if (!event?.dataTransfer) return;
		event.dataTransfer.setData('text/plain', item);
		DraggedItem = element;
		DraggedItemIndex = index;
		console.log(element);
	}

	function swap(indexA: number, indexB: number) {
		const temp = box1[indexA];
		box1[indexA] = box1[indexB];
		box1[indexB] = temp;
	}

	function handleDragOver(event: DragEvent | null, index: number) {
		if (!event) return;
		event.preventDefault();
		DraggedOverItemIndex = index;
	}

	function handleDrop() {
		swap(DraggedItemIndex, DraggedOverItemIndex);
	}
</script>

<!-- <div>
	<h1 class="text-white"><a href="/">home</a></h1>
</div> -->

<section class="w-full h-full flex flex-row gap-4">
	<div class="drag-grid flex-1 border-4" role="none" ondrop={handleDrop}>
		{#each box1 as item, index (item)}
			<div
				class="p-4 m-2 bg-gray-700 text-white rounded"
				draggable="true"
				ondragover={(event) => handleDragOver(event, index)}
				ondragstart={(event) => handleDragStart(event, item, event.currentTarget, index)}
				animate:flip
				role="none"
			>
				{item}
			</div>
		{/each}
	</div>
</section>
