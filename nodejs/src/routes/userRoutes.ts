import express from "express";
import { GetUsers, CreateUser } from "../controllers/userController";
import { Login } from "../controllers/authController";
import { authMiddleWare } from "../middleware/auth";

const router = express.Router();

router.get("/users", authMiddleWare, GetUsers);
router.post("/user", CreateUser);
router.post("/login", Login);

export default router;
