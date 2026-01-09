import { writable } from 'svelte/store';

export const tabs = writable<string[]>([]);
export const activeTab = writable<string | null>(null);
