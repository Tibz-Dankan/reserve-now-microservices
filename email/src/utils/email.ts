import path from "path";
import pug from "pug";
import SGmail from "@sendgrid/mail";

export class Email {
  from: string;
  recipients: string;
  subject: string;
  constructor(recipients: string, subject: string) {
    SGmail.setApiKey(process.env.SENDGRID_API_KEY!);

    this.from = "cryptopile20@gmail.com";
    this.recipients = recipients;
    this.subject = subject;
  }

  async sendHtml(html: any, subject: string) {
    const mailOptions = {
      to: this.recipients,
      from: { email: this.from, name: "Reserve Now" },
      subject: subject,
      html: html,
    };
    try {
      console.log("sending mail");
      await SGmail.send(mailOptions);
      console.log("mail sent");
    } catch (error) {
      console.log("error sending email", error);
    }
  }

  async sendWelcome(firstName: string) {
    const html = pug.renderFile(
      path.join(__dirname, "../templates/welcome.pug"),
      {
        subject: this.subject,
        firstName: firstName,
      }
    );
    await this.sendHtml(html, "Welcome to Reserve Now");
  }

  async sendPasswordReset(url: string, firstName: any) {
    const html = pug.renderFile(
      path.join(__dirname, "../templates/resetPassword.pug"),
      {
        subject: "Password Reset",
        firstName: firstName,
        resetURL: url,
      }
    );
    await this.sendHtml(html, "Reset Password");
  }

  async sendContactUs(
    username: string,
    message: string,
    email: string,
    subject: string
  ) {
    const html = pug.renderFile(
      path.join(__dirname, "../templates/contact.pug"),
      {
        subject: "Contact Us Message",
        username: username,
        message: message,
        email: email,
        contactSubject: subject,
      }
    );
    await this.sendHtml(html, "Contact Us Message");
  }
}
