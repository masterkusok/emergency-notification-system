package handlers

import "github.com/masterkusok/emergency-notification-system/internal/entities"

type deleteContactsRequest struct {
	IdList []uint `json:"id_list"`
}

type updateContactRequest struct {
	Contact entities.Contact `json:"contact"`
}
