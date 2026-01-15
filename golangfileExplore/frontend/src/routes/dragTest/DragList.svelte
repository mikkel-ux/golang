<script lang="ts">
	import type { Item } from '$lib/types/types';
	import type { Snippet } from 'svelte';

	type DragListRenderProps = {
		items: Item[];
		dragState: {
			draggedIndex: number;
			drappedOverIndex: number;
		};
		drop: typeof handleDrop;
		dragstart: typeof handleDragStart;
		dragover: typeof handleDragOver;
		dragend: typeof handleDragEnd;
		dragLeave: typeof dragLeave;
		dragEnter: typeof dragEnter;
	};

	let {
		items = $bindable<Item[]>(),
		children
	}: {
		items?: Item[];
		children?: Snippet<[DragListRenderProps]>;
	} = $props();

	let drappedOverIndex: number = $state(-1);
	let draggedIndex: number = $state(-1);

	function handleDragStart(event: DragEvent | null, index: number) {
		if (!event?.dataTransfer) return;
		event.dataTransfer.setData('text/plain', index.toString());
		/* event.dataTransfer.setDragImage(new Image(), 0, 0); */
		draggedIndex = index;
	}

	function handleDragOver(event: DragEvent | null, overIndex: number) {
		if (!event) return;
		event.preventDefault();
		drappedOverIndex = overIndex;
	}

	function swap(indexA: number, indexB: number) {
		if (!items) return;
		if (indexA === indexB) return;
		if (indexA < 0 || indexB < 0) return;
		if (indexA >= items.length || indexB >= items.length) return;

		const next = items.slice();
		[next[indexA], next[indexB]] = [next[indexB], next[indexA]];
		items = next;
	}

	function handleDrop(event: DragEvent | null) {
		if (!event) return;
		if (draggedIndex === -1 || drappedOverIndex === -1) return;
		swap(draggedIndex, drappedOverIndex);
		handleDragEnd();
	}

	function handleDragEnd() {
		draggedIndex = -1;
		drappedOverIndex = -1;
	}

	function dragLeave(event: DragEvent | null) {
		if (!event) return;
		event.preventDefault();
		drappedOverIndex = -1;
		console.log('drag leave');
	}

	function dragEnter(event: DragEvent | null) {
		if (!event) return;
		event.preventDefault();
		console.log('drag enter');
	}
</script>

{@render children?.({
	items,
	drop: handleDrop,
	dragState: { draggedIndex, drappedOverIndex },
	dragstart: handleDragStart,
	dragover: handleDragOver,
	dragend: handleDragEnd,
	dragLeave: dragLeave,
	dragEnter: dragEnter
})}
