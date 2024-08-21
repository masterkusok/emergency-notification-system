package main

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/masterkusok/emergency-notification-system/internal/handlers"
	"github.com/masterkusok/emergency-notification-system/internal/loaders"
	"github.com/masterkusok/emergency-notification-system/internal/persistence"
	"github.com/masterkusok/emergency-notification-system/internal/routings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{}, &entities.Contact{}, &entities.Template{})

	loader := loaders.CreateContactLoader()
	contactRepo := persistence.CreateContactRepository(db)
	templateRepo := persistence.CreateTemplateRepository(db)
	router := routings.New(handlers.NewContactHandler(contactRepo, loader), handlers.NewTemplateHandler(templateRepo))

	router.Logger.Fatal(router.Start(":1323"))
}
