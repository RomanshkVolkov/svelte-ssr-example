import z from "zod";

export const CreateUserSchema = z.object({
  name: z.string().min(3, "Enter at least 3 characters"),
  email: z.string().email("Type a valid email address"),
  password: z.string().min(6, "The password must be at least 6 characters long")
    .max(500),
  roleID: z.string().cuid2("Select a valid option"),
});
