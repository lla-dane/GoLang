package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	emailFrom := "abhinavagarwalla6@gmail.com"
	emailPassword := "doublemint_00"

	emailTo := []string{
		"abhinavkr.agarwalla.mat22@itbhu.ac.in",
	}

	message := []byte("To: abhinavkr.agarwalla.mat22@itbhu.ac.in\r\n" +
		"Subject: Hello!\r\n" +
		"\r\n" +
		"Test sms!\r\n")

	auth := smtp.PlainAuth("", emailFrom, emailPassword, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, emailTo, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent successfully!")

}
