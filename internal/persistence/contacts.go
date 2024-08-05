package persistence

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func (c *ContactRepository) GetSingleContact(id uint) (entities.Contact, error) {
	contact := entities.Contact{}
	ctx := c.db.Find(&contact, id)
	return contact, ctx.Error
}

func (c *ContactRepository) GetUserContacts(userId uint) ([]entities.Contact, error) {
	contacts := make([]entities.Contact, 0)
	ctx := c.db.Where("user_id = ?", userId).Find(&contacts)
	return contacts, ctx.Error
}

func (c *ContactRepository) CreateContacts(userId uint, contacts []entities.Contact) error {
	for _, contact := range contacts {
		contact.UserID = userId
	}
	ctx := c.db.Create(contacts)
	return ctx.Error
}

func (c *ContactRepository) DeleteContacts(contacts []entities.Contact) error {
	ctx := c.db.Delete(contacts)
	return ctx.Error
}

func (c *ContactRepository) UpdateContact(id uint, name, address string) error {
	contact, err := c.GetSingleContact(id)
	if err != nil {
		return err
	}
	if name != "" {
		contact.Name = name
	}
	if address != "" {
		contact.Address = address
	}
	ctx := c.db.Save(contact)
	return ctx.Error
}
