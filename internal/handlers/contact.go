package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"path/filepath"
	"strconv"
)

var allowedExtensions = map[string]int{
	".json": 1,
	".csv":  2,
	".xlsx": 3,
}

// LoadContacts godoc
// @ID load-contacts
// @Summary Load users contacts
// @Description Load contacts from file types, described in allowed extensions, runs parser and saves them to db.
// @Tags contacts
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Contact file, allowed extensions are .json, .csv, .xlsx"
// @Success 201 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/contacts [post]
func (h *ContactHandler) LoadContacts(c echo.Context) error {
	userId := c.(*AuthContext).Id

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	extension := allowedExtensions[filepath.Ext(file.Filename)] - 1
	if extension == -1 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer src.Close()

	contacts, err := h.loader.ParseFrom(src, extension)
	if err := h.provider.CreateContacts(userId, contacts); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, nil)
}

// GetUserContacts godoc
// @ID get-user-contacts
// @Summary Returns list of all user contacts
// @Description Returns list of all contacts, registered for specific user
// @Tags contacts
// @Accept json
// @Produce json
// @Success 200 {object} manyContactsResponse
// @Failure 400 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/contacts [get]
func (h *ContactHandler) GetUserContacts(c echo.Context) error {
	userId := c.(*AuthContext).Id

	contacts, err := h.provider.GetUserContacts(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	response := new(manyContactsResponse)
	response.Seed(contacts)

	return c.JSON(http.StatusOK, response)
}

// DeleteContacts godoc
// @ID delete-contacts
// @Summary Deletes list of contacts
// @Description Deletes list of contacts, by their ids
// @Tags contacts
// @Accept json
// @Produce json
// @Param list body deleteContactsRequest true "List of contacts to be deleted"
// @Success 200 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/contacts [delete]
func (h *ContactHandler) DeleteContacts(c echo.Context) error {
	request := new(deleteContactsRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := h.provider.DeleteContacts(request.IdList)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// UpdateContact godoc
// @ID update-contact
// @Summary Updates specific contact data
// @Description Can be used to change specific contact field values based on its id
// @Tags contacts
// @Accept json
// @Produce json
// @Param data body updateContactRequest true "New contact data"
// @Param contactID path integer true "Contact Id"
// @Success 200 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/contacts [put]
func (h *ContactHandler) UpdateContact(c echo.Context) error {
	contactId, err := strconv.Atoi(c.Param(":contactId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	request := new(updateContactRequest)
	if err = c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err = h.provider.UpdateContact(uint(contactId), request.Contact.Name, request.Contact.Address); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
