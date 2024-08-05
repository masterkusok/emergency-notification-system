package entities

import "time"

const (
	TG = iota
	EMAIL
	SMS
	PUSH
)

type Contact struct {
	ID        uint      `json:"-"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Platform  int       `json:"platform"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Template struct {
	ID        uint
	Text      string
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID           uint
	Username     string `gorm:"unique"`
	Salt         string
	PasswordHash string
	templates    []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
