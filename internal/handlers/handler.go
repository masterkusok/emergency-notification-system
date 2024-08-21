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

type templateProvider interface {
	CreateTemplate(userId uint, text string) error
	DeleteTemplate(templateId uint) error
	GetUserTemplates(userId uint) ([]entities.Template, error)
}

type TemplateHandler struct {
	provider templateProvider
}

func NewTemplateHandler(provider templateProvider) *TemplateHandler {
	return &TemplateHandler{provider: provider}
}
