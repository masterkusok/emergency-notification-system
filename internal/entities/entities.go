package entities

import (
	"gorm.io/gorm"
	"time"
)

const (
	TG = iota
	EMAIL
	SMS
	PUSH
)

type Contact struct {
	gorm.Model
	ID        uint      `json:"-"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Platform  int       `json:"platform"`
	Name      string    `json:"name" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Template struct {
	gorm.Model
	ID        uint      `json:"id"`
	Text      string    `validate:"required" json:"text"`
	UserID    uint      `json:"-"`
	User      User      `validate:"required" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type User struct {
	gorm.Model
	ID           uint
	Username     string `gorm:"unique"`
	Salt         string
	PasswordHash string
	templates    []string
	contacts     []Contact
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
