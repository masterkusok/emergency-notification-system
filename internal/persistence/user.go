package persistence

import (
	"github.com/go-playground/validator/v10"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	baseRepository
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	repo := UserRepository{baseRepository{db: db, validator: validator.New()}}
	return &repo
}

func (u *UserRepository) CreateUser(username, salt, hash string) (*entities.User, error) {
	user := &entities.User{Username: username, Salt: salt, PasswordHash: hash}
	ctx := u.db.Create(user)
	return user, ctx.Error
}

func (u *UserRepository) GetUserById(id uint) (*entities.User, error) {
	user := new(entities.User)
	ctx := u.db.Find(user, id)
	return user, ctx.Error
}

func (u *UserRepository) GetUserByName(username string) (*entities.User, error) {
	user := new(entities.User)
	ctx := u.db.Where(&entities.User{Username: username}).First(user)
	return user, ctx.Error
}

func (u *UserRepository) GetUserEager(id uint) (*entities.User, error) {
	user := new(entities.User)
	err := u.db.Model(&entities.User{}).Preload("Contacts").Preload("Templates").Find(user, id).Error
	return user, err
}
