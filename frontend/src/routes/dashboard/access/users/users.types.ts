export type UserRow = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  email: string;
  role: { id: string; name: string };
};
