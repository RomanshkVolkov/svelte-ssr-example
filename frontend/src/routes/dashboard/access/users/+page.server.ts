import type { ServerLoad } from "@sveltejs/kit";
import { fetcherAPI } from "@/api/api.main";
import type { UserRow } from "./users.types";

// ssr fetch data
export const load: ServerLoad = async () => {
  const response = await fetcherAPI<UserRow[]>("/users");

  return {
    users: response.data,
  };
};
