package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	// Create a new email message
	m := gomail.NewMessage()

	// Set the sender and recipient
	m.SetHeader("From", "sambacarlson@gmail.com")
	m.SetHeader("To", "cedrick@iknite.studio")
	m.SetHeader("Subject", "Test Email")
	m.SetBody("text/plain", "This is a test email sent using gomail.")

	// SMTP server configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, "sambacarlson@gmail.com", GOOGLE_APP_PASSWORD)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	println("Email sent successfully!")
}
