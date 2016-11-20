package email

import (
	"errors"
	"net/smtp"
	"os"
)

func Send(to []string, msg []byte) error {
	pass := os.Getenv("EMAIL_PASS")
	if pass == "" {
		return errors.New("EMAIL_PASS is empty")
	}
	auth := smtp.PlainAuth("", "johnshenkthegreat@gmail.com", pass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "johnshenkthegreat@gmail.com", to, msg)
	return err
}
