package routings

import (
	"github.com/labstack/echo/v4"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
	"github.com/masterkusok/emergency-notification-system/internal/middleware"
)

const PREFIX = "/api/v1"

func New(contactHandler *handlers.ContactHandler, templateHandler *handlers.TemplateHandler, authHandler *handlers.AuthHandler) *echo.Echo {
	e := echo.New()

	guestGroup := e.Group(PREFIX + "/auth")
	authGroup := e.Group(PREFIX)

	authGroup.Use(middleware.AuthJWT)

	// contact routes
	authGroup.POST("/contacts", contactHandler.LoadContacts)
	authGroup.GET("/contacts", contactHandler.GetUserContacts)
	authGroup.DELETE("/contacts", contactHandler.DeleteContacts)
	authGroup.PUT("/contacts/:contactId", contactHandler.UpdateContact)

	// template routes
	authGroup.POST("/templates", templateHandler.CreateTemplate)
	authGroup.GET("/templates", templateHandler.GetUserTemplates)
	authGroup.DELETE("/templates/:templateId", templateHandler.DeleteTemplate)

	authGroup.GET("/current", authHandler.CurrentUser)

	// auth routes
	guestGroup.POST("/register", authHandler.SignUp)
	guestGroup.POST("/login", authHandler.SignIn)
	return e
}
