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

func (t *TemplateRepository) GetUserTemplates(userId uint) ([]entities.Template, error) {
	templates := make([]entities.Template, 1)
	ctx := t.db.Where(entities.Template{UserID: userId}).Find(&templates)
	return templates, ctx.Error
}

func (t *TemplateRepository) CreateTemplate(userId uint, text string) error {
	template := &entities.Template{UserID: userId, Text: text}
	err := t.validator.Struct(template)
	if err != nil {
		return err
	}
	ctx := t.db.Create(template)
	return ctx.Error
}

func (t *TemplateRepository) DeleteTemplate(templateId uint) error {
	ctx := t.db.Delete(&entities.Template{}, templateId)
	return ctx.Error
}
