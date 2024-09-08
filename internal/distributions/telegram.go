package distributions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/masterkusok/emergency-notification-system/internal/persistence"
	"log"
)

type TelegramDistributor struct {
	bot  *tgbotapi.BotAPI
	repo *persistence.TelegramChatRepository
}

func NewTelegramDistributor(repo *persistence.TelegramChatRepository, token string) (*TelegramDistributor, error) {
	distributor := new(TelegramDistributor)
	distributor.repo = repo

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	distributor.bot = bot
	go distributor.listen()
	return distributor, nil
}

func (t TelegramDistributor) listen() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := t.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			username := update.Message.From.UserName

			if _, err := t.repo.GetChatIdByUsername(username); err != nil {
				err := t.repo.CreateChat(username, chatID)
				if err != nil {
					log.Printf("Failed to memorize chat id, id:%d, username: %s\n", chatID, username)
				}
			}
		}
	}
}

func (t TelegramDistributor) Send(message string, contact entities.Contact) error {
	chatID, err := t.repo.GetChatIdByUsername(contact.Address)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := t.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
