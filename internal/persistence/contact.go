package persistence

import (
	"github.com/go-playground/validator/v10"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/gorm"
)

type ContactRepository struct {
	baseRepository
}

func CreateContactRepository(db *gorm.DB) *ContactRepository {
	repo := ContactRepository{baseRepository{db: db, validator: validator.New()}}
	return &repo
}

func (c *ContactRepository) GetSingleContact(id uint) (entities.Contact, error) {
	contact := entities.Contact{}
	ctx := c.db.Find(&contact, id)
	return contact, ctx.Error
}

func (c *ContactRepository) GetUserContacts(userId uint) ([]entities.Contact, error) {
	contacts := make([]entities.Contact, 0)
	ctx := c.db.Where(entities.Contact{UserID: userId}).Find(&contacts)
	return contacts, ctx.Error
}

func (c *ContactRepository) CreateContacts(userId uint, contacts []entities.Contact) error {
	for _, contact := range contacts {
		contact.UserID = userId
		err := c.validator.Struct(contact)
		if err != nil {
			return err
		}
	}
	ctx := c.db.Create(contacts)
	return ctx.Error
}

func (c *ContactRepository) DeleteContacts(id []uint) error {
	ctx := c.db.Delete(&entities.Contact{}, id)
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
	err = c.validator.Struct(contact)
	if err != nil {
		return err
	}
	ctx := c.db.Save(contact)
	return ctx.Error
}
