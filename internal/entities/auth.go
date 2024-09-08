package entities

import "time"

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
