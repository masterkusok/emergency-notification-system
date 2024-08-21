package handlers

import "github.com/masterkusok/emergency-notification-system/internal/entities"

type userResponse struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
}

func (r *userResponse) Seed(user *entities.User) {
	r.Username = user.Username
	r.Id = user.ID
}

type manyContactsResponse struct {
	Contacts []entities.Contact `json:"contacts"`
}

func (r *manyContactsResponse) Seed(contacts []entities.Contact) {
	r.Contacts = contacts
}

type manyTemplatesResponse struct {
	Templates []entities.Template `json:"templates"`
}

func (r *manyTemplatesResponse) Seed(templates []entities.Template) {
	r.Templates = templates
}
