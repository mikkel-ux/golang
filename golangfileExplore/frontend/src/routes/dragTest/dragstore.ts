import type { Item } from '$lib/types/types';

export const dragstate = $state({
	lists: {
		list1: [] as Item[],
		list2: [] as Item[]
	}
});

type ListId = keyof typeof dragstate.lists;

export function move(sourceId: ListId, sourceIndex: number, targetId: ListId, targetIndex: number) {
	const source = dragstate.lists[sourceId];
	const target = dragstate.lists[targetId];

	const [item] = source.splice(sourceIndex, 1);

	if (targetIndex === -1) target.push(item);
	else target.splice(targetIndex, 0, item);
}
