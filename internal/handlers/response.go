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

type singleTemplateResponse struct {
	Id     uint   `json:"template_id"`
	UserId uint   `json:"user_id"`
	Text   string `json:"text"`
}

type manyTemplatesResponse struct {
	Templates []singleTemplateResponse `json:"templates"`
}
