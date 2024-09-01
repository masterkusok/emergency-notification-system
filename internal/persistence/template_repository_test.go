package persistence

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"testing"
)

var templateRepository *TemplateRepository
var templateDb *gorm.DB

func init() {
	templateDb, _ = gorm.Open(sqlite.Open("templateTest.userDb"), &gorm.Config{})
	templateDb.AutoMigrate(&entities.Template{})
	templateRepository = CreateTemplateRepository(templateDb)
}

func TestTemplateRepository_CreateTemplate(t *testing.T) {
	// clear userDb
	templateDb.Exec("DELETE FROM templates")

	err := templateRepository.CreateTemplate(0, "sample text")

	if err != nil {
		log.Println("Failed to create template:", err)
		t.Fail()
	}

	dbTemplates := []entities.Template{}
	err = templateDb.Find(&dbTemplates).Error

	if len(dbTemplates) != 1 {
		log.Println("Expected 1 template, found:", len(dbTemplates))
		t.Fail()
	}

	if dbTemplates[0].Text != "sample text" {
		log.Println("Expected text 'sample text', found:", dbTemplates[0].Text)
		t.Fail()
	}
}

func TestTemplateRepository_GetUserTemplates(t *testing.T) {
	templateDb.Exec("DELETE FROM templates")

	templates := []entities.Template{
		{Text: "template 1"},
		{Text: "template 2"},
		{Text: "template 3"},
	}

	err := templateDb.Create(templates).Error
	if err != nil {
		log.Println("Failed to create templates:", err)
		t.Fail()
	}

	templatesFromRepo, _ := templateRepository.GetUserTemplates(0)

	if len(templatesFromRepo) != len(templates) {
		log.Println("Expected", len(templates), "templates, found:", len(templatesFromRepo))
		t.Fail()
	}

	for i := 0; i < len(templates); i++ {
		if templates[i].Text != templatesFromRepo[i].Text {
			log.Println("Expected text", templates[i].Text, "found:", templatesFromRepo[i].Text)
			t.Fail()
		}
	}
}

func TestTemplateRepository_DeleteTemplate(t *testing.T) {
	templateDb.Exec("DELETE FROM templates")

	templates := []entities.Template{
		{Text: "template 1"},
		{Text: "template 2"},
		{Text: "template 3"},
	}

	err := templateDb.Create(templates).Error
	if err != nil {
		log.Println("Failed to create templates:", err)
		t.Fail()
	}

	err = templateRepository.DeleteTemplate(templates[0].ID)
	if err != nil {
		log.Println("Failed to delete template:", err)
		t.Fail()
	}

	dbTemplates := []entities.Template{}
	err = templateDb.Find(&dbTemplates).Error

	if err != nil {
		log.Println("Failed to find templates:", err)
		t.Fail()
	}
	if len(dbTemplates) != len(templates)-1 {
		log.Println("Expected", len(templates)-1, "templates, found:", len(dbTemplates))
		t.Fail()
	}
}
