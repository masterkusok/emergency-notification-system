package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUserTemplates godoc
// @ID get-templates
// @Summary Returns list of user templates
// @Description Returns list of user templates
// @Tags templates
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/templates [get]
func (h *TemplateHandler) GetUserTemplates(c echo.Context) error {
	userId := c.(*AuthContext).Id
	templates, err := h.provider.GetUserTemplates(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	response := new(manyTemplatesResponse)
	response.Seed(templates)

	return c.JSON(http.StatusCreated, response)
}

// CreateTemplate godoc
// @ID create-template
// @Summary Creates new template
// @Description Creates new template
// @Tags templates
// @Accept json
// @Produce json
// @Param data body createTemplateRequest true "New template data"
// @Success 201 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/templates [post]

func (h *TemplateHandler) CreateTemplate(c echo.Context) error {
	userId := c.(*AuthContext).Id
	request := new(createTemplateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := c.Validate(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.provider.CreateTemplate(uint(userId), request.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusCreated, nil)
}

// DeleteTemplate godoc
// @ID delete-template
// @Summary Deletes template
// @Description Deletes template by its id
// @Tags templates
// @Accept json
// @Produce json
// @Param data body createTemplateRequest true "New template data"
// @Param templateID path integer true "Template id"
// @Success 201 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security JwtAuth
// @Router /ap1/v1/templates [post]
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
