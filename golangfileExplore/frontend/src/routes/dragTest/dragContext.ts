import type { Item } from '$lib/types/types';

export type DragContext = {
	sourceId: string | null;
	item: Item | null;
	sourceIndex: number;
};

export const dragContext: DragContext = {
	sourceId: null,
	item: null,
	sourceIndex: -1
};
