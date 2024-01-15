import { Request, Response, NextFunction } from "express";
import { AppError } from "../utils/error";
import { asyncHandler } from "../utils/asyncHandler";
import { Email } from "../utils/email";

export const sendMail = asyncHandler(
  async (req: Request, res: Response, next: NextFunction) => {
    // Send mail here
  }
);

export const sendPasswordResetMail = asyncHandler(
  async (req: Request, res: Response, next: NextFunction) => {
    console.log("req.body", req.body);
    console.log("req.query", req.body);
    const recipient = req.body.recipient;
    const name = req.body.name;
    const resetURL = req.body.resetURL;

    if (!recipient) return next(new AppError("Please provide recipient", 400));
    if (!name) return next(new AppError("Please provide username", 400));
    if (!resetURL) return next(new AppError("Please provide resetURL", 400));

    await new Email(recipient, "Password Reset").sendPasswordReset(
      resetURL,
      name
    );

    res
      .status(200)
      .json({ message: "Mail sent successfully", status: "success" });
  }
);
