package service

import (
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
	"strconv"
)

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(from, to, subject, body string) error {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/plain", body)

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	username := os.Getenv("EMAIL_LOGIN")
	password := os.Getenv("EMAIL_PASSWORD")

	dialer := gomail.NewDialer(smtpHost, smtpPort, username, password)

	if err := dialer.DialAndSend(message); err != nil {
		log.Println("Error:", err)
		return err
	} else {
		log.Println("Email sent successfully!")
		return nil
	}
}
