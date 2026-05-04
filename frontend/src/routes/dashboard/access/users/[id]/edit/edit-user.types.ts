import type { BaseModel } from "@/types/base-model";

export type GetUserByIDResponse = BaseModel & {
  email: string;
  name: string;
  otp: string;
  expiresOTP: string;
  role: BaseModel & {
    name: string;
  };
};
