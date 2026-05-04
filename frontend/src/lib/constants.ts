import { env } from "$env/dynamic/public";

export const USER_ROLES = {
  ADMIN: "qqovic77d7bz181zv3l00kq6",
  EMPLOYE: "rwfcnio5w3mnxspqzrlrs9ae",
};

export const ROUTES = {
  userList: {
    label: "Users",
    path: "/dashboard/access/users",
  },
  createUser: {
    label: "Create user",
    path: "/dashboard/access/users/create",
  },
  editUser: {
    label: "Edit user",
    path: "/dashboard/access/users/[id]/edit",
  },
  deleteUser: {
    label: "Delete user",
    path: "/dashboard/access/users/[id]/delete",
  },
  roleList: {
    label: "Roles",
    path: "/dashboard/access/roles",
  },
  createRole: {
    label: "Create role",
    path: "/dashboard/access/roles/create",
  },
  editProfile: {
    label: "Edit profile",
    path: "/dashboard/access/roles/[id]/edit",
  },
  deleteProfile: {
    label: "Delete profile",
    path: "/dashboard/access/roles/[id]/delete",
  },
};

export const URL_API = env.PUBLIC_API_URL || "http://localhost:8080/api/v1";
