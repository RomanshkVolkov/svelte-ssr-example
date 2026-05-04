import {
  message,
  superValidate,
  type SuperValidated,
} from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { fail, type RequestEvent } from "@sveltejs/kit";
import type { AnyZodObject, z } from "zod";

type ServerActionHandler<T extends AnyZodObject> = (options: {
  safeData: z.infer<T>;
  form: SuperValidated<z.infer<T>, any>;
  event: RequestEvent;
}) => Promise<any>;

interface ServerAction<T extends AnyZodObject> {
  action: (handler: ServerActionHandler<T>) => Promise<any>;
}

function ValidateSchema<T extends AnyZodObject>(
  event: RequestEvent,
  schema: T,
) {
  return {
    action: async (handler: ServerActionHandler<T>) => {
      const form = await superValidate(event.request, zod(schema)) as any;

      if (!form.valid) {
        return message(form, "Please fix the errors in the form.", {
          status: 400,
        });
      }

      return await handler({
        safeData: form.data,
        form,
        event,
      });
    },
  };
}

export const serverAction = { ValidateSchema };
