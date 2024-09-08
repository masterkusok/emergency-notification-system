package persistence

import (
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/gorm"
)

type TelegramChatRepository struct {
	baseRepository
}

func CreateTelegramChatRepository(db *gorm.DB) *TelegramChatRepository {
	repo := TelegramChatRepository{baseRepository{db: db}}
	return &repo
}

func (t *TelegramChatRepository) GetChatIdByUsername(username string) (int64, error) {
	chat := new(entities.TelegramChat)
	err := t.db.Where(&entities.TelegramChat{Username: username}).First(chat).Error
	if chat == nil {
		return 0, fmt.Errorf("chat not found")
	}
	return chat.ChatId, err
}

func (t *TelegramChatRepository) CreateChat(username string, id int64) error {
	chat := new(entities.TelegramChat)
	chat.ChatId = id
	chat.Username = username
	err := t.db.Create(chat).Error
	return err
}
