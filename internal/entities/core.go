// Package entities contains all base structures
package entities

import (
	"time"
)

const (
	TG = iota
	EMAIL
	SMS
)

type Contact struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	Platform  int       `json:"platform"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Template struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	UserID    uint      `json:"-"`
	User      User      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
