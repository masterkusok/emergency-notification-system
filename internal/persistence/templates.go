package persistence

import (
	"github.com/go-playground/validator/v10"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/gorm"
)

type TemplateRepository struct {
	baseRepository
}

func CreateTemplateRepository(db *gorm.DB) *TemplateRepository {
	repo := &TemplateRepository{baseRepository{db: db, validator: validator.New()}}
	return repo
}

func (t *TemplateRepository) CreateTemplate(userId uint, text string) (*entities.Template, error) {
	template := &entities.Template{UserID: userId, Text: text}
	err := t.validator.Struct(template)
	if err != nil {
		return nil, err
	}
	ctx := t.db.Create(template)
	return template, ctx.Error
}

func (t *TemplateRepository) DeleteTemplate(templateId uint) error {
	ctx := t.db.Delete(templateId)
	return ctx.Error
}
