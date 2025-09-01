import { redirect } from "@sveltejs/kit";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ fetch, url }) => {
    const token = document.cookie.split('; ').find(row => row.startsWith('token='))?.split('=')[1];

    const response = await fetch("/api/validate", {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`
		}
    });


    if (response.ok) {
        if (url.pathname === "/login" || url.pathname === "/signup"){
            throw redirect(302, "/foo");
        }
    }
    }