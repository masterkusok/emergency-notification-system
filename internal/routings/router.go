package routings

import (
	"github.com/labstack/echo/v4"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
)

const PREFIX = "/api/v1/"

func New(contactHandler *handlers.ContactHandler, templateHandler *handlers.TemplateHandler) *echo.Echo {
	e := echo.New()

	// contact routes
	e.POST(PREFIX+"users/:userId/contacts", contactHandler.LoadContacts)
	e.GET(PREFIX+"users/:userId/contacts", contactHandler.GetUserContacts)
	e.DELETE(PREFIX+"users/:userId/contacts", contactHandler.DeleteContacts)
	e.PUT(PREFIX+"users/:userId/contacts/:contactId", contactHandler.UpdateContact)

	// template routes
	e.POST(PREFIX+"users/:userId/templates", templateHandler.CreateTemplate)
	e.GET(PREFIX+"users/:userId/templates", templateHandler.GetUserTemplates)
	e.DELETE(PREFIX+"users/:userId/templates/:templateId", templateHandler.DeleteTemplate)

	return e
}
