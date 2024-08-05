package main

import (
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entities.User{}, &entities.Contact{}, &entities.Template{})
	if err != nil {
		panic(err)
	}
	fmt.Println("abc")
}
