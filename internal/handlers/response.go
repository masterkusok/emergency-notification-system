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

type signInResponse struct {
	IsSuccessful bool   `json:"is_successful"`
	Token        string `json:"token"`
	Message      string `json:"message"`
}

func (r *signInResponse) Seed(isSuccessful bool, token, message string) {
	r.IsSuccessful = isSuccessful
	r.Message = message
	r.Token = token
}

type ForbiddenResponse struct {
	Message string `json:"message"`
}

func GetForbiddenResponse() *ForbiddenResponse {
	return &ForbiddenResponse{Message: "authorization is required for this request"}
}
