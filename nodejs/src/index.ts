import express from "express";
import userRouter from "./routes/userRoutes";
import dotenv from "dotenv";

dotenv.config();

const app = express();
app.use(express.json());
app.use("/api/v1", userRouter);

const PORT = 8000;

app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
  