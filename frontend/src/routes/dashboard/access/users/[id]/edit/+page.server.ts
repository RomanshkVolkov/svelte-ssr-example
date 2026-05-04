import {type Actions, redirect} from "@sveltejs/kit";
import {superValidate} from "sveltekit-superforms";
import {zod} from "sveltekit-superforms/adapters";
import type z from "zod";
import {fetcherAPI, handleFormError} from "@/api/api.main.js";
import {ROUTES} from "@/constants.js";
import {serverAction} from "@/safe-actions.js";
import type {CommonOptionTypes} from "@/types/base-model.js";
import type {PageServerLoad} from "./$types.js";
import {EditUserSchema} from "./edit-user.schema.js";
import type {GetUserByIDResponse} from "./edit-user.types.js";
import type {RoleListResponse} from "../../common-user.types";

export const load: PageServerLoad = async (event) => {
    const {id} = event.params;
    const userListResponse = await fetcherAPI<GetUserByIDResponse>(
        `/users/${id}`,
    );
    const roleListResponse = await fetcherAPI<RoleListResponse[]>("/roles");

    return {
        roles: roleListResponse.data.map(({id: value, name: label}) => ({label, value})),
        form: await superValidate(zod(EditUserSchema), {
            defaults: {
                id,
                name: userListResponse.data.name,
                roleID: userListResponse.data.role.id ?? "",
                email: userListResponse.data.email ?? "",
            },
        }),
    };
};

export const actions = {
    default: async (event) =>
        serverAction
            .ValidateSchema(event, EditUserSchema)
            .action(async ({safeData, form}) => {
                const {id, ...data} = safeData;

                const response = await fetcherAPI<z.infer<typeof EditUserSchema>>(
                    `/users/${id}`,
                    {
                        method: "PUT",
                        body: JSON.stringify(data),
                    },
                );

                if (!response.success) {
                    return handleFormError(form, response);
                }

                return redirect(303, ROUTES.usersList.path);
            }),
} satisfies Actions;
