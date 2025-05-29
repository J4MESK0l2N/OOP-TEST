import { z } from "zod";

const USER_ROLES = z.enum(["user", "admin"]);

export interface PublicUserSchema {
  id: number;
  username: string;
  role?: string;
}

export interface CreateUserSchema {
  username: string;
  password: string;
  role?: string;
}

export const JwtPayloadSchema = z.object({
  id: z.number(),
  username: z.string(),
  role: USER_ROLES,
});

export const CreateUserValidate = z.object({
  username: z.string().min(8),
  password: z.string().min(8),
  role: USER_ROLES.optional(),
});

export type CreateUserInput = z.infer<typeof CreateUserValidate>;
