package persistence

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type baseRepository struct {
	db        *gorm.DB
	validator *validator.Validate
}
