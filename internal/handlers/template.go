package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *TemplateHandler) GetUserTemplates(c echo.Context) error {
	userId := c.(*AuthContext).Id
	templates, err := h.provider.GetUserTemplates(uint(userId))
	if err != nil {
		return err
	}
	response := new(manyTemplatesResponse)
	response.Seed(templates)

	return c.JSON(http.StatusCreated, response)
}

func (h *TemplateHandler) CreateTemplate(c echo.Context) error {
	userId := c.(*AuthContext).Id
	request := new(createTemplateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := h.provider.CreateTemplate(uint(userId), request.Text)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, nil)
}

func (h *TemplateHandler) DeleteTemplate(c echo.Context) error {
	templateId, err := strconv.Atoi(c.Param("templateId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = h.provider.DeleteTemplate(uint(templateId))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
