// Package entities contains all base structures
package entities

import (
	"time"
)

const (
	TG = iota
	EMAIL
	SMS
	PUSH
)

type Contact struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Platform  int       `json:"platform"`
	Name      string    `json:"name" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Template struct {
	ID        uint      `json:"id"`
	Text      string    `validate:"required" json:"text"`
	UserID    uint      `json:"-"`
	User      User      `validate:"required" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type User struct {
	ID           uint
	Username     string `gorm:"unique"`
	Salt         string
	PasswordHash string
	Templates    []Template
	Contacts     []Contact
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
