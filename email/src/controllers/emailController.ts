import { Request, Response, NextFunction } from "express";
import { AppError } from "../utils/error";
import { asyncHandler } from "../utils/asyncHandler";

export const sendMail = asyncHandler(
  async (req: Request, res: Response, next: NextFunction) => {
    // Send mail here
  }
);
