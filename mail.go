// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/smtp"
)

func codeGenerator() (code string) {
	bytes := make([]byte, 5)
	rand.Read(bytes)
	code = hex.EncodeToString(bytes)
	return code
}

func sendMail(password, dest, code string) {
	auth := smtp.PlainAuth("", "safecallnoreply@gmail.com", password, "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{dest}

	msg := []byte("To: " + dest + "\r\n" +

		"Subject: Forgot my Password SafeCall\r\n" +

		"\r\n" +

		"This is the code you'll need to change your password " + code + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "safecallnoreply@gmail.com", to, msg)

	if err != nil {
		log.Fatal(err)
	}

}
