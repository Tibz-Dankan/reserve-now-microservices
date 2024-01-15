import express from "express";
import {
  sendMail,
  sendPasswordResetMail,
} from "../controllers/emailController";

const router = express.Router();

router.post("/send-mail", sendMail);
router.post("/password-reset", sendPasswordResetMail);

export { router as emailRoutes };
