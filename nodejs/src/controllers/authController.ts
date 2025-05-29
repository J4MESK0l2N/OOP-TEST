import { Request, Response } from "express";
import jwt from "jsonwebtoken";

export const Login = (req: Request, res: Response) => {
  const { username, password, role } = req.body;

  try {
    if (username === "admin" && password === "@min1234") {
      const token = jwt.sign({ id: 1, username: username, role: role }, process.env.JWT_SECRET!, {
        expiresIn: "2h",
      });

      res.status(200).send({ data: { token: token }, code: 200, message: "OK" });
      return;
    }

    res.status(400).send({ message: "Server Error.", code: 400 });
  } catch (err) {
    res.status(401).send({ error: err, message: "Wrong Username or Password.", code: 401 });
  }
};

module.exports = { Login };
