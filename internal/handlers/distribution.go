package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Distribute godoc
// @ID distribute
// @Summary Distributes template message
// @Description Sends message with template text to all contacts in users list
// @Tags distribution
// @Accept json
// @Produce json
// @Param templateId path integer true "Id of template to be distributed"
// @Success 200 {object} nil
// @Failure 400 {object} nil
// @Failure 500 {object} nil
// @Security BearerAuth
// @Router /api/v1/distribute [post]
func (d *DistributionHandler) Distribute(c echo.Context) error {
	userId := c.(*AuthContext).Id
	templateId, err := strconv.Atoi(c.Param("templateId"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user, err := d.userProvider.GetUserEager(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	targetTemplateText := ""
	for _, template := range user.Templates {
		if template.ID == uint(templateId) {
			targetTemplateText = template.Text
		}
	}

	if len(targetTemplateText) == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	for _, contact := range user.Contacts {
		err = d.distributor.Send(targetTemplateText, contact)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}
	return c.JSON(http.StatusOK, nil)
}
