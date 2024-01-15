import express from "express";
import { sendMail } from "../controllers/emailController";

const router = express.Router();

router.post("/send-mail", sendMail);

export { router as emailRoutes };
