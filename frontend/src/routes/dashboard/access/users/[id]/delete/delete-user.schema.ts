import z from "zod";

export const DeleteUserSchema = z.object({
  id: z.string().cuid2("Please provide a valid user ID"),
});
