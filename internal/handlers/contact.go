package handlers

import (
	"fmt"
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

func (h *ContactHandler) LoadContacts(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	extension := allowedExtensions[filepath.Ext(file.Filename)] - 1
	if extension == -1 {
		return fmt.Errorf("Extension is not supported\n")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	contacts, err := h.loader.ParseFrom(src, extension)
	if err := h.provider.CreateContacts(uint(userId), contacts); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, nil)
}

func (h *ContactHandler) GetUserContacts(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}

	contacts, err := h.provider.GetUserContacts(uint(userId))
	if err != nil {
		return err
	}
	response := new(manyContactsResponse)
	response.Seed(contacts)

	return c.JSON(http.StatusCreated, response)
}

func (h *ContactHandler) DeleteContacts(c echo.Context) error {
	request := new(deleteContactsRequest)
	if err := c.Bind(&request); err != nil {
		return err
	}

	err := h.provider.DeleteContacts(request.IdList)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *ContactHandler) UpdateContact(c echo.Context) error {
	contactId, err := strconv.Atoi(c.Param(":contactId"))
	if err != nil {
		return err
	}

	request := new(updateContactRequest)
	if err = c.Bind(&request); err != nil {
		return err
	}

	if err = h.provider.UpdateContact(uint(contactId), request.Contact.Name, request.Contact.Address); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
