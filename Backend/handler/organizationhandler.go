package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"p3/datamodels"
)

type OrganizationHandler struct {
	dbInstance *gorm.DB
}

func NewOrganizationHandler(router *echo.Echo, db *gorm.DB) {
	handler := &OrganizationHandler{
		dbInstance: db,
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

	if h.dbInstance == nil {
		return c.JSON(http.StatusBadRequest, "Connection Error")
	}
	if mErr := h.dbInstance.AutoMigrate(&datamodels.Organization{}); mErr != nil {
		return c.JSON(http.StatusBadRequest, mErr)
	}
	e.IsActive = true
	e.IsDeleted = false
	if err = h.dbInstance.Create(e).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, e)
	//return c.JSON(http.StatusOK, e)
}
func (h *OrganizationHandler) GetOrg(c echo.Context) error {

	emailId := c.QueryParams().Get("email")
	if h.dbInstance == nil {
		return c.JSON(http.StatusBadRequest, "Connection Error")
	}
	var o *datamodels.Organization
	// Get first matched record
	h.dbInstance.Where("created_by = ?", emailId).First(&o)
	if o == nil {
		return c.JSON(http.StatusForbidden, "Data not found")
	}
	return c.JSON(http.StatusForbidden, o)
}
