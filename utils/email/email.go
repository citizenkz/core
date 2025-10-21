package email

import (
	"fmt"
	"net/smtp"

	"github.com/citizenkz/core/config"
)

type EmailService struct {
	cfg *config.Config
}

func New(cfg *config.Config) *EmailService {
	return &EmailService{
		cfg: cfg,
	}
}

func (e *EmailService) SendOTP(to, otpCode string) error {
	subject := "Password Reset OTP Code"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .otp-code { font-size: 24px; font-weight: bold; color: #4CAF50; padding: 10px; background: #f5f5f5; border-radius: 5px; text-align: center; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Password Reset Request</h2>
        <p>You have requested to reset your password. Use the OTP code below:</p>
        <div class="otp-code">%s</div>
        <p>This code will expire in 10 minutes.</p>
        <p>If you didn't request this, please ignore this email.</p>
    </div>
</body>
</html>
`, otpCode)

	return e.send(to, subject, body)
}

func (e *EmailService) SendPasswordChanged(to string) error {
	subject := "Password Successfully Changed"
	body := `
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Password Changed</h2>
        <p>Your password has been successfully changed.</p>
        <p>If you didn't make this change, please contact support immediately.</p>
    </div>
</body>
</html>
`

	return e.send(to, subject, body)
}

func (e *EmailService) SendEmailChanged(to, newEmail string) error {
	subject := "Email Address Changed"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Email Address Changed</h2>
        <p>Your email address has been successfully changed to: <strong>%s</strong></p>
        <p>If you didn't make this change, please contact support immediately.</p>
    </div>
</body>
</html>
`, newEmail)

	return e.send(to, subject, body)
}

func (e *EmailService) SendAccountDeleted(to string) error {
	subject := "Account Deleted"
	body := `
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <h2>Account Deleted</h2>
        <p>Your account has been successfully deleted.</p>
        <p>We're sorry to see you go. If you change your mind, you can create a new account anytime.</p>
    </div>
</body>
</html>
`

	return e.send(to, subject, body)
}

func (e *EmailService) send(to, subject, body string) error {
	from := e.cfg.SMTP.From
	if from == "" {
		from = e.cfg.SMTP.Username
	}

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", from, to, subject, body)

	auth := smtp.PlainAuth("", e.cfg.SMTP.Username, e.cfg.SMTP.Password, e.cfg.SMTP.Host)

	addr := fmt.Sprintf("%s:%d", e.cfg.SMTP.Host, e.cfg.SMTP.Port)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
