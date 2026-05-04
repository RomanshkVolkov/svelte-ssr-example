export type APIResponse<T> = {
  success: boolean;
  message: Record<"es" | "en", string>;
  data: T;
  schema: Record<string, string[]> | null;
  error: string;
};
