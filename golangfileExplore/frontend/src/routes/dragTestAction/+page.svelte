<script lang="ts">
	type itemsType = {
		id: number;
		name: string;
	};
	let items: itemsType[] = [{ id: 1, name: 'Item 1' }];
	function dndzone(node: HTMLElement, params: { items: any[] }) {
		node.addEventListener('mouseenter', () => {
			node.style.cursor = 'pointer';
		});
		let draggedEl: HTMLElement | null = null;

		function onPointerDown(e: PointerEvent) {
			const target = e.target as HTMLElement;
			const index = Number(target.dataset.index);

			if (target.parentElement !== node) return;

			draggedEl = target;
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
			}

			function onUp(ev: PointerEvent) {
				window.removeEventListener('pointermove', onMove);
				window.removeEventListener('pointerup', onUp);
				ghost.remove();
				draggedEl = null;
			}

			/* window.onpointermove = (ev: PointerEvent) => {
				node.style.userSelect = 'none';
				target.style.position = 'absolute';
				target.style.left = ev.pageX + 'px';
				target.style.top = ev.pageY + 'px';
			};
			window.onpointerup = (ev: PointerEvent) => {
				node.style.userSelect = 'auto';
				window.onpointermove = null;
				window.onpointerup = null;
			}; */

			window.addEventListener('pointermove', onMove);
			window.addEventListener('pointerup', onUp);
		}

		node.addEventListener('pointerdown', onPointerDown);

		return {
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

<!-- <ul use:dndzone={{ items }} class="dropAriar text-white border p-4 h-30 overflow-auto">
	{#each items as item (item.id)}
		<li class="p-4 m-2 bg-gray-700 rounded list-none" data-index={item.id}>
			{item.name}
		</li>
	{/each}
</ul> -->

<section
	use:dndzone={{ items }}
	class="dropAriar text-white border p-4 h-30 overflow-auto"
	on:consider={handleConsider}
	on:finalize={handleFinalize}
>
	{#each items as item (item.id)}
		<div class="p-4 m-2 bg-gray-700 rounded list-none text-white" data-index={item.id}>
			{item.name}
		</div>
	{/each}
</section>
