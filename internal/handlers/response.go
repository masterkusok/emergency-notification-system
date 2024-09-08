package handlers

import "github.com/masterkusok/emergency-notification-system/internal/entities"

type distributionResponse struct {
	SuccessfulCount int    `json:"successful_count"`
	FailedCount     int    `json:"failed_count"`
	FailedIdList    []uint `json:"failed_id_list"`
}

type singleTemplateResponse struct {
	IsSuccessful bool               `json:"is_successful"`
	Template     *entities.Template `json:"template"`
}

func (r *singleTemplateResponse) Seed(isSuccessful bool, template *entities.Template) {
	r.IsSuccessful = isSuccessful
	r.Template = template
}

type singleContactResponse struct {
	IsSuccessful bool              `json:"is_successful"`
	Contact      *entities.Contact `json:"contact"`
}

func (r *singleContactResponse) Seed(isSuccessful bool, contact *entities.Contact) {
	r.IsSuccessful = isSuccessful
	r.Contact = contact
}

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
