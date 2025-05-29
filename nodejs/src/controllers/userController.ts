import { Request, Response } from "express";

import { User } from "../models/User";
import { CreateUserValidate, CreateUserSchema, PublicUserSchema } from "../schemas/UserSchemas";

const users: User[] = [];

export const GetUsers = (req: Request, res: Response) => {
  try {
    const user_list: PublicUserSchema[] = users.map((user) => user.PublicUserToJson());
    res.status(200).send({ data: { users: user_list }, code: 200, message: "OK." });
    return;
  } catch (err) {
    console.log("err", err);
    res.status(400).send({ error: err, message: "Something went wrong.", code: 400 });
    return;
  }
};

export const CreateUser = (req: Request, res: Response) => {
  try {
    const result = CreateUserValidate.safeParse(req.body);

    if (!result.success) {
      res.status(400).send({ message: result.error.format(), code: 400 });
      return;
    }

    const user_body: CreateUserSchema = result.data;
    const user = new User(
      users.length + 1,
      user_body.username,
      user_body.password,
      user_body.role || "user"
    );

    users.push(user);

    res.status(201).send({ message: "OK", code: 201, data: user.PublicUserToJson() });
  } catch (err) {
    res.status(400).send({ message: "Something went wrong.", code: 400 });
  }
};

module.exports = { GetUsers, CreateUser };
