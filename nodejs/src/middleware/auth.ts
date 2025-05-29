/// <reference path="../types/express/express.d.ts" />

import { NextFunction, Request, Response } from "express";
import jwt from "jsonwebtoken";
import { User } from "../models/User";
import { JwtPayloadSchema, PublicUserSchema } from "../schemas/UserSchemas";

export function authMiddleWare(req: Request, res: Response, next: NextFunction) {
  const token = req.headers.authorization?.split(" ")[1];

  if (!token) {
    res.status(401).json({ message: "Unauthorized", code: 401 });
    return;
  }
  try {
    const jwt_data = jwt.verify(token, process.env.JWT_SECRET!);
    const result = JwtPayloadSchema.safeParse(jwt_data);

    if (!result.success) {
      res.status(400).send({ message: result.error.format(), code: 400 });
      return;
    }

    const payload: PublicUserSchema = result.data;
    req.user = new User(payload.id, payload.username, "", payload.role);
    next();
  } catch (err) {
    res.status(403).json({ message: "Forbidden" });
  }
}
