/** biome-ignore-all lint/suspicious/noExplicitAny: this funciton is used on genenic interfaces */

import {fail} from "@sveltejs/kit";
import {type SuperValidated, setError} from "sveltekit-superforms";
import type {APIResponse} from "@/api/api.types";
import {URL_API} from "@/constants";

export async function fetcherAPI<
    T extends Record<string, unknown> | Record<string, unknown>[],
>(path: string, options?: RequestInit): Promise<APIResponse<T>> {
    const response = await fetch(`${URL_API}${path}`, options);
    return (await response.json()).data as APIResponse<T>;
}

export function handleFormError<T extends Record<string, unknown>>(
    form: SuperValidated<T>,
    response: APIResponse<T>,
) {
    const schemaErrors = response.schema;
    if (schemaErrors) {
        const keys = Object.keys(schemaErrors);

        for (const key of keys) {
            setError(form as any, key, schemaErrors[key]?.[0] ?? response.message.en);
        }
    }

    return fail(400, {
        form,
        message: response.message.en,
    });
}
