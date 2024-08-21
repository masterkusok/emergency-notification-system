package routings

import (
	"github.com/labstack/echo/v4"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
)

const PREFIX = "/api/v1/"

func New(handler *handlers.ContactHandler) *echo.Echo {
	e := echo.New()
	e.POST(PREFIX+"users/:id/contacts", handler.LoadContacts)
	e.GET(PREFIX+"users/:id/contacts", handler.GetUserContacts)
	return e
}
