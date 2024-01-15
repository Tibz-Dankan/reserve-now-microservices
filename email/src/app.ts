import express, { Request, Response } from "express";
import cors from "cors";
import dotenv from "dotenv";
import logger from "morgan";
import http from "http";
import { errorController } from "./controllers/errorController";
import { emailRoutes } from "./routes/emailRoutes";

dotenv.config();

const app = express();

app.use(cors({ origin: "*" }));

const server = http.createServer(app);

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(logger("dev"));

app.use("/api/v1/email", emailRoutes);

app.use(errorController);

app.use("*", (req: Request, res: Response) => {
  res.status(404).json({
    status: "fail",
    message: "Endpoint not found!",
  });
});

const PORT = 5000 || process.env.PORT;

server.listen(PORT, () => {
  console.log(`email server running on port ${PORT}`);
});
