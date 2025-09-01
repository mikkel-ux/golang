import { redirect } from "@sveltejs/kit";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ fetch, url }) => {
    const token = document.cookie.split('; ').find(row => row.startsWith('token='))?.split('=')[1];
    if (!token) {
        throw redirect(302, "/login");
    }

    const response = await fetch("/api/validate", {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`
		}
    });

    if (!response.ok) {
        throw redirect(302, "/login");
    }
}