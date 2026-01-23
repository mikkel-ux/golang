<script lang="ts">
	type itemsType = {
		id: number;
		name: string;
	};
	let items: itemsType[] = [
		{ id: 1, name: 'Item 1' },
		{ id: 2, name: 'Item 2' }
	];
	function dndzone(node: HTMLElement, params: { items: any[] }) {
		node.addEventListener('mouseenter', () => {
			node.style.cursor = 'pointer';
		});
		let draggedEl: HTMLElement | null = null;
		let draggedItem: any = null;

		function onPointerDown(e: PointerEvent) {
			const target = e.target as HTMLElement;
			if (target.parentElement !== node) return;

			node.style.userSelect = 'none';
			draggedEl = target;

			const index = Array.from(node.children).indexOf(target);
			draggedItem = params.items[index];
			console.log('Dragging item at index:', index, draggedItem);

			const rect = target.getBoundingClientRect();

			const ghost = target.cloneNode(true) as HTMLElement;
			ghost.style.position = 'fixed';
			ghost.style.width = rect.width + 'px';
			ghost.style.left = rect.left + 'px';
			ghost.style.top = rect.top + 'px';
			ghost.style.pointerEvents = 'none';
			ghost.style.color = 'white';
			ghost.style.opacity = '0.7';

			document.body.appendChild(ghost);

			function onMove(ev: PointerEvent) {
				ghost.style.left = ev.pageX - rect.width / 2 + 'px';
				ghost.style.top = ev.pageY - rect.height / 2 + 'px';
				console.log(document.elementFromPoint(ev.clientX, ev.clientY));
				const elemBelow = document.elementFromPoint(ev.clientX, ev.clientY) as HTMLElement;
				if (!elemBelow) return;
				const dropZone = elemBelow.closest('.dropAriar') as HTMLElement;
				if (dropZone) {
					onEnterDropZone(dropZone);
				} else {
					removeAllHighlights();
				}
			}

			function removeAllHighlights() {
				const dropZones = document.querySelectorAll('.dropAriar');
				dropZones.forEach((dz) => onLeaveDropZone(dz as HTMLElement));
			}

			function onEnterDropZone(dropZone: HTMLElement) {
				dropZone.style.backgroundColor = 'rgba(255, 255, 255, 0.1)';
			}

			function onLeaveDropZone(dropZone: HTMLElement) {
				dropZone.style.backgroundColor = '';
			}

			function onUp(ev: PointerEvent) {
				window.removeEventListener('pointermove', onMove);
				window.removeEventListener('pointerup', onUp);
				ghost.remove();
				draggedEl = null;
				console.log('Pointer up at:', ev.pageX, ev.pageY);
				removeAllHighlights();
				node.style.userSelect = 'auto';
			}

			window.addEventListener('pointermove', onMove);
			window.addEventListener('pointerup', onUp);
		}

		node.addEventListener('pointerdown', onPointerDown);

		return {
			update(newParams: { items: any[] }) {
				params = newParams;
			},
			destroy() {
				node.removeEventListener('mouseenter', () => {
					node.style.cursor = 'pointer';
				});
				node.removeEventListener('pointerdown', onPointerDown);
			}
		};
	}

	function handleConsider(event: CustomEvent) {
		console.log(event);
	}

	function handleFinalize(event: CustomEvent) {
		console.log(event);
	}
</script>

<section
	use:dndzone={{ items }}
	class="dropAriar text-white border p-4 h-50 overflow-auto"
	on:consider={handleConsider}
	on:finalize={handleFinalize}
>
	{#each items as item (item.id)}
		<div class="p-4 m-2 bg-gray-700 rounded list-none text-white">
			{item.name}
		</div>
	{/each}
</section>
