package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"p3/datamodels"
	"p3/errors"
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
	customError := errors.CustomError{}
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid post body")
	}
	if err = c.Validate(e); err != nil {
		customError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, customError)
	}

	if h.dbInstance == nil {
		customError.Message = "Connection Error"
		return c.JSON(http.StatusBadRequest, customError)
	}
	var o *datamodels.Organization
	h.dbInstance.Where("created_by = ?", "hemant.manwani@loginradius.com").First(&o)
	if o != nil {
		customError.Message = "You have already created the organization."
		return c.JSON(http.StatusBadRequest, customError)
	}
	if mErr := h.dbInstance.AutoMigrate(&datamodels.Organization{}); mErr != nil {
		customError.Message = mErr.Error()
		return c.JSON(http.StatusBadRequest, customError)
	}
	e.IsActive = true
	e.IsDeleted = false
	e.CreatedBy = "hemant.manwani@loginradius.com"
	if err = h.dbInstance.Create(e).Error; err != nil {
		customError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, customError)
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
