import z from "zod";

export const EditUserSchema = z.object({
  id: z.string().cuid2(),
  name: z.string().min(3, "Enter at least 3 characters"),
  email: z.string().email("Type a valid email address"),
  roleID: z.string().cuid2("Select a valid option"),
});
