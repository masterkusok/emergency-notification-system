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

const (
	TG_TOKEN       = "_"
	EXOLVE_API_KEY = "_"
	EXOLVE_NUMBER  = "_"
	GMAIL_ADDRESS  = "_"
	GMAIL_PASSWORD = "_"
)

// @title ENS API
// @version 1.0
// @description This is API for an Emergency Notification System app
// @host localhost:1323
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{}, &entities.Contact{}, &entities.Template{}, &entities.TelegramChat{})

	// repositories init
	tgRepo := persistence.CreateTelegramChatRepository(db)
	contactRepo := persistence.CreateContactRepository(db)
	templateRepo := persistence.CreateTemplateRepository(db)
	userRepo := persistence.CreateUserRepository(db)

	// loader init
	loader := loaders.CreateContactLoader()

	// distributors init
	tg, err := distributions.NewTelegramDistributor(tgRepo, TG_TOKEN)
	if err != nil {
		panic(err)
	}

	sms, err := distributions.NewSMSDistributor(EXOLVE_API_KEY, EXOLVE_NUMBER)
	if err != nil {
		panic(err)
	}

	email, err := distributions.NewSMTPDistributor(GMAIL_ADDRESS, GMAIL_PASSWORD)
	if err != nil {
		panic(err)
	}
	distributor := distributions.CreateDistributor(tg, sms, email)

	router := routes.New(handlers.NewContactHandler(contactRepo, loader), handlers.NewTemplateHandler(templateRepo),
		handlers.NewAuthHandler(userRepo), handlers.NewDistributionHandler(distributor, userRepo))

	router.Logger.Fatal(router.Start(":1323"))
}
