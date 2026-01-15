<script lang="ts">
	import type { Item } from '$lib/types/types';
	import { flip } from 'svelte/animate';
	import DragList from './DragList.svelte';
	let box1 = $state<Item[]>([
		{ id: 1, name: 'item 1' },
		{ id: 2, name: 'item 2' },
		{ id: 3, name: 'item 3' }
	]);
	let box2 = $state<Item[]>([
		{ id: 4, name: 'item 4' },
		{ id: 5, name: 'item 5' },
		{ id: 6, name: 'item 6' }
	]);
</script>

<!-- <div>
	<h1 class="text-white"><a href="/">home</a></h1>
</div> -->

<section class="w-full h-full flex flex-row gap-4">
	<DragList bind:items={box1}>
		{#snippet children({
			items,
			dragState,
			dragstart,
			dragover,
			dragend,
			drop,
			dragLeave,
			dragEnter
		})}
			<div
				ondrop={drop}
				ondragleave={dragLeave}
				ondragenter={dragEnter}
				role="none"
				class="drag-grid flex-1 border-4"
			>
				{#each items as item, index (item)}
					<div
						class="test p-4 m-2 bg-gray-700 text-white rounded"
						class:opacity-50={index === dragState.draggedIndex}
						class:border-4={index === dragState.drappedOverIndex}
						draggable="true"
						ondragover={(event) => dragover(event, index)}
						ondragstart={(event) => dragstart(event, index)}
						ondragend={dragend}
						role="none"
						animate:flip={{ duration: 350 }}
					>
						{item.name}
					</div>
				{/each}
			</div>
		{/snippet}
	</DragList>
	<DragList bind:items={box2}>
		{#snippet children({
			items,
			dragState,
			dragstart,
			dragover,
			dragend,
			drop,
			dragLeave,
			dragEnter
		})}
			<div ondrop={drop} role="none" class="drag-grid flex-1 border-4">
				{#each items as item, index (item)}
					<div
						class="p-4 m-2 bg-gray-700 text-white rounded"
						class:opacity-50={index === dragState.draggedIndex}
						class:border-4={index === dragState.drappedOverIndex}
						draggable="true"
						ondragover={(event) => dragover(event, index)}
						ondragstart={(event) => dragstart(event, index)}
						ondragend={dragend}
						role="none"
						animate:flip={{ duration: 350 }}
					>
						{item.name}
					</div>
				{/each}
			</div>
		{/snippet}
	</DragList>
</section>
