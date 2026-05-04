import prisma from "$lib/prisma";
import { USER_ROLES } from "@/constants";
import type { Prisma } from "@prisma/client";

export async function seedUsers() {
  await createIfNotExistUserRoles();

  const users: Prisma.UserCreateManyArgs["data"] = [
    {
      email: "jose@guz-studio.dev",
      password: "",
      name: "Root user",
      roleID: USER_ROLES.ADMIN,
    },
  ];

  await prisma.user.createMany({ data: users });
}

async function createIfNotExistUserRoles() {
  const roles = [
    { id: USER_ROLES.ADMIN, name: "Admin" },
    { id: USER_ROLES.EMPLOYE, name: "Employe" },
  ];

  for (const role of roles) {
    const isExist = await prisma.role.findUnique({ where: { id: role.id } });

    if (!isExist) {
      const createdRole = await prisma.role.create({
        data: role,
      });

      console.info("Created role: ", createdRole.name);
    }
  }
}
