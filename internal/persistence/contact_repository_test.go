package persistence

import (
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"testing"
)

var contactRepo *ContactRepository
var contactDb *gorm.DB

func init() {
	contactDb, _ = gorm.Open(sqlite.Open("test.userDb"), &gorm.Config{})
	contactDb.AutoMigrate(&entities.Contact{})
	contactRepo = CreateContactRepository(contactDb)
}

func TestContactRepository_CreateContacts(t *testing.T) {
	// clear userDb
	contactDb.Exec("DELETE FROM contacts")
	contactList := []entities.Contact{
		{
			Address: "addr1", Name: "name1", Platform: entities.TG,
		},
		{
			Address: "addr2", Name: "name2", Platform: entities.EMAIL,
		},
		{
			Address: "addr3", Name: "name3", Platform: entities.SMS,
		},
	}

	err := contactRepo.CreateContacts(0, contactList)
	if err != nil {
		log.Println("Failed to create contacts:", err)
		t.Fail()
	}

	dbContacts := []entities.Contact{}
	err = contactDb.Find(&dbContacts).Error
	if err != nil {
		log.Println("Failed to find contacts in DB:", err)
		t.Fail()
	}

	if len(dbContacts) != len(contactList) {
		log.Println("Expected", len(contactList), "contacts, found:", len(dbContacts))
		t.Fail()
	}

	for i := 0; i < len(contactList); i++ {
		if dbContacts[i].Name != contactList[i].Name {
			log.Println("Expected name", contactList[i].Name, "found:", dbContacts[i].Name)
			t.Fail()
		}
		if dbContacts[i].Address != contactList[i].Address {
			log.Println("Expected address", contactList[i].Address, "found:", dbContacts[i].Address)
			t.Fail()
		}
		if dbContacts[i].Platform != contactList[i].Platform {
			log.Println("Expected platform", contactList[i].Platform, "found:", dbContacts[i].Platform)
			t.Fail()
		}
	}
}

func TestContactRepository_GetSingleContact(t *testing.T) {
	contactDb.Exec("DELETE FROM contacts")
	expectedContact := entities.Contact{Platform: entities.TG, Name: "name", Address: "address"}
	r := contactDb.Create(&expectedContact)
	if r.Error != nil {
		log.Println("Failed to create contact:", r.Error)
		t.Fail()
	}

	contact, err := contactRepo.GetSingleContact(expectedContact.ID)
	if err != nil {
		log.Println("Failed to get single contact by ID:", err)
		t.Fail()
	}

	if contact.Name != expectedContact.Name {
		log.Println("Expected name", expectedContact.Name, "found:", contact.Name)
		t.Fail()
	}

	if contact.Address != expectedContact.Address {
		log.Println("Expected address", expectedContact.Address, "found:", contact.Address)
		t.Fail()
	}

	if contact.Platform != expectedContact.Platform {
		log.Println("Expected platform", expectedContact.Platform, "found:", contact.Platform)
		t.Fail()
	}

}

func TestContactRepository_DeleteContact(t *testing.T) {
	contactDb.Exec("DELETE FROM contacts")
	contactList := []entities.Contact{
		{
			Address: "addr1", Name: "name1", Platform: entities.TG,
		},
		{
			Address: "addr2", Name: "name2", Platform: entities.EMAIL,
		},
		{
			Address: "addr3", Name: "name3", Platform: entities.SMS,
		},
	}

	err := contactRepo.CreateContacts(0, contactList)
	if err != nil {
		log.Println("Failed to create contacts:", err)
		t.Fail()
	}

	repoContacts, err := contactRepo.GetUserContacts(0)
	if err != nil {
		log.Println("Failed to get user contacts:", err)
		t.Fail()
	}

	randomId := repoContacts[rand.Intn(len(repoContacts))].ID

	err = contactRepo.DeleteContact(randomId)
	if err != nil {
		log.Println("Failed to delete contacts:", err)
		t.Fail()
	}

	dbContacts := []entities.Contact{}
	contactDb.Find(&dbContacts)
	for _, contact := range dbContacts {
		if contact.ID == randomId {
			log.Println("Failed to delete random id")
			t.Fail()
		}
	}
}

func TestContactRepository_UpdateContact(t *testing.T) {
	contactDb.Exec("DELETE FROM contacts")
	contact := entities.Contact{
		Platform: entities.TG,
		Name:     "name",
		Address:  "Address",
	}
	contactDb.Create(&contact)

	err := contactRepo.UpdateContact(contact.ID, "", "zzzz")
	if err != nil {
		log.Println("Failed to update contact address:", err)
		t.Fail()
	}

	repoContact, _ := contactRepo.GetSingleContact(contact.ID)
	if repoContact.Address != "zzzz" {
		log.Println("Expected address 'zzzz', found:", repoContact.Address)
		t.Fail()
	}

	err = contactRepo.UpdateContact(contact.ID, "zycv", "zzz")
	if err != nil {
		log.Println("Failed to update contact name and address")
		t.Fail()
	}

	repoContact, _ = contactRepo.GetSingleContact(contact.ID)
	if repoContact.Address != "zzz" || repoContact.Name != "zycv" {
		log.Println("Invalid name or address after updating contact")
		t.Fail()
	}
}
