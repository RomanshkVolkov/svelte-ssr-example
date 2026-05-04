import {type Actions, redirect} from "@sveltejs/kit";
import {superValidate} from "sveltekit-superforms";
import {zod} from "sveltekit-superforms/adapters";
import type z from "zod";
import {fetcherAPI, handleFormError} from "@/api/api.main.js";
import {ROUTES} from "@/constants.js";
import {serverAction} from "@/safe-actions.js";
import type {PageServerLoad} from "./$types.js";
import {CreateUserSchema} from "./create-user.shema.js";
import type {RoleListResponse} from "../common-user.types";

export const load: PageServerLoad = async () => {
        const roleListResponse = await fetcherAPI<RoleListResponse[]>("/roles");
        return {
            roles: roleListResponse.data.map(({id: value, name: label}) => ({label, value})),
            form: await superValidate(zod(CreateUserSchema)),
        };
    }
;

export const actions = {
    default: async (event) =>
        serverAction
            .ValidateSchema(event, CreateUserSchema)
            .action(async ({safeData, form}) => {
                const response = await fetcherAPI<z.infer<typeof CreateUserSchema>>(
                    "/users",
                    {
                        method: "POST",
                        body: JSON.stringify(safeData),
                    },
                );

                if (!response.success) {
                    return handleFormError(form, response);
                }

                redirect(303, ROUTES.userList.path);
            }),
} satisfies Actions;
