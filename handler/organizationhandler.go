package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"p3/datamodels"
	cErr "p3/err"
	"p3/utils"
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
	router.PUT("/organization", handler.EditOrgName)
}

func (h *OrganizationHandler) AddOrg(c echo.Context) error {
	var e datamodels.Organization
	err := c.Bind(&e)
	email := c.Get("Email").(string)
	if email == "" {
		return utils.GetError(cErr.InvalidEmail)
	}
	customError := cErr.CustomError{}
	if err != nil {
		return utils.GetError(cErr.InvalidPostBody)
	}
	if err = c.Validate(e); err != nil {
		customError.Message = err.Error()
		//return utils.GetError(c, customError)
		return c.JSON(http.StatusBadRequest, customError)
	}

	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var o *datamodels.Organization
	h.dbInstance.Where("created_by = ?", email).First(&o)
	if o.Name != "" && o.IsActive {
		return utils.GetError(cErr.OrgAlreadyExist)
	}
	if mErr := h.dbInstance.AutoMigrate(&datamodels.Organization{}); mErr != nil {
		customError.Message = mErr.Error()
		return c.JSON(http.StatusBadRequest, customError)
	}
	e.IsActive = true
	e.IsDeleted = false
	e.CreatedBy = email
	if err = h.dbInstance.Create(&e).Error; err != nil {
		customError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, customError)
	}
	return c.JSON(http.StatusOK, e)
	//return c.JSON(http.StatusOK, e)
}
func (h *OrganizationHandler) GetOrg(c echo.Context) error {

	email := c.Get("Email").(string)
	if email == "" {
		return utils.GetError(cErr.InvalidEmail)
	}
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var o *datamodels.Organization
	// Get first matched record
	h.dbInstance.Where("created_by = ?", email).Select([]string{"name", "created_at", "updated_at", "created_by"}).First(&o)
	if o == nil || o.Name == "" {
		return utils.GetError(cErr.DataNotFound)
	}
	return c.JSON(http.StatusOK, o)
}
func (h *OrganizationHandler) EditOrgName(c echo.Context) error {

	email := c.Get("Email").(string)
	if email == "" {
		return utils.GetError(cErr.InvalidEmail)
	}
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var e datamodels.EditOrganization
	err := c.Bind(&e)
	if err != nil {
		return utils.GetError(cErr.InvalidPostBody)
	}
	customError := cErr.CustomError{}
	if err = c.Validate(e); err != nil {
		customError.Message = err.Error()
		return c.JSON(http.StatusBadRequest, customError)
	}
	var o *datamodels.Organization
	// Get first matched record
	h.dbInstance.Where("created_by = ?", email).First(&o)
	if o == nil || o.Name == "" {
		return utils.GetError(cErr.DataNotFound)
	}
	o.Name = e.Name
	h.dbInstance.Save(o)
	return c.JSON(http.StatusOK, o)
}
