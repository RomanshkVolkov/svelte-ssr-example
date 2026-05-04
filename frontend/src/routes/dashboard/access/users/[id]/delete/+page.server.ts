import { redirect } from "@sveltejs/kit";
import { message, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { fetcherAPI } from "@/api/api.main";
import { ROUTES } from "@/constants";
import { serverAction } from "@/safe-actions";
import type { PageServerLoad } from "./$types";
import { DeleteUserSchema } from "./delete-user.schema";
import type { DeleteUserResponse } from "./delete-user.types";

export const load: PageServerLoad = async ({ params }) => {
  const { id } = params;
  const userResponse = await fetcherAPI<DeleteUserResponse>(
    `/users/${id}?fields=ID,Name,Email`,
    {
      method: "GET",
    },
  );

  if (!userResponse.success) {
    return redirect(303, ROUTES.usersList.path);
  }

  return {
    user: userResponse.data,
    form: await superValidate(zod(DeleteUserSchema), {
      defaults: {
        id,
      },
    }),
  };
};

export const actions = {
  default: async (event) =>
    serverAction
      .ValidateSchema(event, DeleteUserSchema)
      .action(async ({ safeData, form }) => {
        const response = await fetcherAPI<DeleteUserResponse>(
          `/users/${safeData.id}`,
          {
            method: "DELETE",
          },
        );

        if (!response.success) {
          return message(form, response.message.en);
        }

        redirect(303, ROUTES.usersList.path);
      }),
};
