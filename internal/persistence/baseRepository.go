package persistence

import (
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}
