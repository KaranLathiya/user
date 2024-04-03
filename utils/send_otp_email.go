package utils

import (
	"net/smtp"
	"user/config"

	error_handling "user/error"
)

func SendOTPEmail(email string, otp string, subject string) error {
	from := config.ConfigVal.SMTP.EmailFrom
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		otp
	err := smtp.SendMail(config.ConfigVal.SMTP.Host+":"+config.ConfigVal.SMTP.Port,
		smtp.PlainAuth("", from, config.ConfigVal.SMTP.EmailPassword, config.ConfigVal.SMTP.Host),
		from, []string{to}, []byte(msg))
	if err != nil {
		return error_handling.SendEmailError
	}
	return nil
}
