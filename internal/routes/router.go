// Package routes provides routing for all api handlers
package routes

import (
	"github.com/labstack/echo/v4"
	_ "github.com/masterkusok/emergency-notification-system/internal/docs"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
	"github.com/masterkusok/emergency-notification-system/internal/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const PREFIX = "/api/v1"

// New godoc
// This method is used to create new echo object with all necessary routes, middlewares etc.
func New(contactHandler *handlers.ContactHandler, templateHandler *handlers.TemplateHandler,
	authHandler *handlers.AuthHandler, distributionHandler *handlers.DistributionHandler) *echo.Echo {
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

	authGroup.POST("/distribute/:templateId", distributionHandler.Distribute)

	// auth routes
	guestGroup.POST("/register", authHandler.SignUp)
	guestGroup.POST("/login", authHandler.SignIn)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
