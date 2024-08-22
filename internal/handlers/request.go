package handlers

import "github.com/masterkusok/emergency-notification-system/internal/entities"

type deleteContactsRequest struct {
	IdList []uint `json:"id_list"`
}

type updateContactRequest struct {
	Contact entities.Contact `json:"contact"`
}

type createTemplateRequest struct {
	Text string `json:"text"`
}

type signUpRequest struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type signInRequest struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
