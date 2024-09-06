package main

import (
	"github.com/masterkusok/emergency-notification-system/internal/distributions"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
	"github.com/masterkusok/emergency-notification-system/internal/loaders"
	"github.com/masterkusok/emergency-notification-system/internal/persistence"
	"github.com/masterkusok/emergency-notification-system/internal/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title ENS API
// @version 1.0
// @description This is API for an Emergency Notification System app
// @host localhost:1323
// @BasePath /

// @securityDefinitions.jwt JwtAuth
// @in header
// @name Authorization

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{}, &entities.Contact{}, &entities.Template{})

	loader := loaders.CreateContactLoader()
	distributor := distributions.CreateDistributor()
	contactRepo := persistence.CreateContactRepository(db)
	templateRepo := persistence.CreateTemplateRepository(db)
	userRepo := persistence.CreateUserRepository(db)

	router := routes.New(handlers.NewContactHandler(contactRepo, loader), handlers.NewTemplateHandler(templateRepo),
		handlers.NewAuthHandler(userRepo), handlers.NewDistributionHandler(distributor, userRepo))

	router.Logger.Fatal(router.Start(":1323"))
}
