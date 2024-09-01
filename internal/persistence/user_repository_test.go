package persistence

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

var userRepository *UserRepository
var userDb *gorm.DB

func init() {
	userDb, _ = gorm.Open(sqlite.Open("templateTest.userDb"), &gorm.Config{})
	userDb.AutoMigrate(&entities.Template{})
	userRepository = CreateUserRepository(userDb)
}

func TestUserRepository_CreateUser(t *testing.T) {
	// clear userDb
	userDb.Exec("DELETE FROM users")

	_, err := userRepository.CreateUser("name", "salt", "hash")

	if err != nil {
		log.Println("Failed to create user:", err)
		t.Fail()
	}

	dbUsers := []entities.User{}
	err = userDb.Find(&dbUsers).Error

	if len(dbUsers) != 1 {
		log.Println("Expected 1 user, found:", len(dbUsers))
		t.Fail()
	}

	if dbUsers[0].Username != "name" {
		log.Println("Expected username 'name', found:", dbUsers[0].Username)
		t.Fail()
	}

	if dbUsers[0].Salt != "salt" {
		log.Println("Expected salt 'salt', found:", dbUsers[0].Salt)
		t.Fail()
	}

	if dbUsers[0].PasswordHash != "hash" {
		log.Println("Expected password hash 'hash', found:", dbUsers[0].PasswordHash)
		t.Fail()
	}
}

func TestUserRepository_GetUserById(t *testing.T) {
	// clear userDb
	userDb.Exec("DELETE FROM users")

	user := &entities.User{Username: "name", PasswordHash: "hash", Salt: "salt"}

	err := userDb.Create(user).Error

	if err != nil {
		log.Println("Failed to create user:", err)
		t.Fail()
	}

	dbUser, err := userRepository.GetUserById(user.ID)
	if err != nil {
		log.Println("Failed to get user by ID:", err)
		t.Fail()
	}

	if dbUser.Username != user.Username {
		log.Println("Expected username 'name', found:", dbUser.Username)
		t.Fail()
	}
	if dbUser.Salt != user.Salt {
		log.Println("Expected salt 'salt', found:", dbUser.Salt)
		t.Fail()
	}
	if dbUser.PasswordHash != user.PasswordHash {
		log.Println("Expected password hash 'hash', found:", dbUser.PasswordHash)
		t.Fail()
	}
}

func TestUserRepository_GetUserByName(t *testing.T) {
	// clear userDb
	userDb.Exec("DELETE FROM users")

	user := &entities.User{Username: "name", PasswordHash: "hash", Salt: "salt"}

	err := userDb.Create(user).Error

	if err != nil {
		log.Println("Failed to create user:", err)
		t.Fail()
	}

	dbUser, err := userRepository.GetUserByName(user.Username)
	if err != nil {
		log.Println("Failed to get user by name:", err)
		t.Fail()
	}

	if dbUser.Username != user.Username {
		log.Println("Expected username 'name', found:", dbUser.Username)
		t.Fail()
	}
	if dbUser.Salt != user.Salt {
		log.Println("Expected salt 'salt', found:", dbUser.Salt)
		t.Fail()
	}
	if dbUser.PasswordHash != user.PasswordHash {
		log.Println("Expected password hash 'hash', found:", dbUser.PasswordHash)
		t.Fail()
	}
}
