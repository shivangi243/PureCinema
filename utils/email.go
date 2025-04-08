package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

func SendTicket(toEmail, filePath string) error {
	e := email.NewEmail()

	e.From = "Cinema Ticketing <rudrakkho.work@gmail.com>"
	e.To = []string{toEmail}
	e.Subject = "🎟️ Your Cinema Booking Confirmation"
	e.Text = []byte("Thank you for booking your ticket. Find your receipt attached.")

	// Check if file exists
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file not found: %v", err)
	}

	// Attach the file
	_, err = e.AttachFile(filePath)
	if err != nil {
		return fmt.Errorf("could not attach file: %v", err)
	}

	// Gmail SMTP (TLS on port 465)
	smtpServer := "smtp.gmail.com"
	smtpPort := 465
	auth := smtp.PlainAuth("", "rudrakkho.work@gmail.com", "jgnmxjpsmkuuuswr", smtpServer)

	// Use TLS connection (important!)
	err = e.SendWithTLS(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, &tls.Config{
		ServerName: smtpServer,
	})
	if err != nil {
		return fmt.Errorf("email sending failed: %v", err)
	}

	return nil
}
