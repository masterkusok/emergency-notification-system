// Package distributions provides functionality to send messages on different platforms.
// TODO: realize concrete distributors
package distributions

import (
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"log"
)

type Sender interface {
	Send(string, entities.Contact) error
}

type mockSender struct {
}

func (m mockSender) Send(template string, contact entities.Contact) error {
	log.Printf(`Mock send:\nMessage:"%s"\nTo address: %s\nTo name: %sName\nplatform: %d`,
		template, contact.Address, contact.Name, contact.Platform)
	return nil
}

type Distributor struct {
	senders []Sender
}

func CreateDistributor(tg *TelegramDistributor, sms *SMSDistributor, email *SMTPDistributor) *Distributor {
	d := Distributor{}
	d.senders = make([]Sender, 3)
	d.senders[entities.TG] = tg
	d.senders[entities.SMS] = sms
	d.senders[entities.EMAIL] = email
	return &d
}

func (d *Distributor) Send(template string, contact entities.Contact) error {
	if contact.Platform >= len(d.senders) || contact.Platform < 0 {
		return fmt.Errorf("invalid contact platform")
	}
	if len(template) == 0 {
		return fmt.Errorf("empty template")
	}

	return d.senders[contact.Platform].Send(template, contact)
}
