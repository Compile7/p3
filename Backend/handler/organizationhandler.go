package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"p3/datamodels"
	"p3/logic"
)

type OrganizationHandler struct {
	oLogic *logic.OrganizationLogic
}

func NewOrganizationHandler(router *echo.Echo, logic *logic.OrganizationLogic) {
	handler := &OrganizationHandler{
		oLogic: logic,
	}
	router.POST("/organization", handler.AddOrg)
	router.GET("/organization", handler.GetOrg)
}

func (h *OrganizationHandler) AddOrg(c echo.Context) error {
	var e datamodels.Organization
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid post body")
	}
	if err = c.Validate(e); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	res, addErr := h.oLogic.AddOrg(&e)
	if addErr != "" {
		return c.JSON(http.StatusBadRequest, addErr)
	}
	return c.JSON(http.StatusOK, res)
	//return c.JSON(http.StatusOK, e)
}
func (h *OrganizationHandler) GetOrg(c echo.Context) error {

	emailId := c.QueryParams().Get("email")
	res, err := h.oLogic.GetOrg(emailId)
	if res != nil {
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusForbidden, err)
}
