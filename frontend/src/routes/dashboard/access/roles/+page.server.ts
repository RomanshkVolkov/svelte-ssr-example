import type {ServerLoad} from "@sveltejs/kit";
import {fetcherAPI} from "@/api/api.main";
import type {ProfileRow} from "./roles.types";

export const load: ServerLoad = async ({params}) => {
    const response = await fetcherAPI<ProfileRow[]>("/roles");

    return {
        profiles: response.data,
    };
};
