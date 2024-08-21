package handlers

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/masterkusok/emergency-notification-system/internal/loaders"
)

type contactProvider interface {
	CreateContacts(uint, []entities.Contact) error
	GetUserContacts(uint) ([]entities.Contact, error)
	DeleteContacts([]uint) error
	UpdateContact(id uint, name, address string) error
}

type ContactHandler struct {
	loader   *loaders.ContactLoader
	provider contactProvider
}

func NewContactHandler(provider contactProvider, loader *loaders.ContactLoader) *ContactHandler {
	return &ContactHandler{provider: provider, loader: loader}
}
