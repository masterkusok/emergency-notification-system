package distributions

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"net/smtp"
)

type SMTPDistributor struct {
	auth     smtp.Auth
	email    string
	password string
}

func NewSMTPDistributor(email, password string) (*SMTPDistributor, error) {
	distributor := new(SMTPDistributor)
	distributor.password = password
	distributor.email = email
	auth := smtp.PlainAuth("", email, password, "smtp.gmail.com")
	distributor.auth = auth
	return distributor, nil
}

func (s *SMTPDistributor) Send(message string, contact entities.Contact) error {
	err := smtp.SendMail("smtp.gmail.com:587", s.auth, s.email, []string{contact.Address}, []byte(message))
	return err
}
