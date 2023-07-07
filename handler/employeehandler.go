package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"p3/datamodels"
)

type EmployeeHandler struct {
	dbInstance *gorm.DB
}

func NewEmployeeHandler(router *echo.Echo, db *gorm.DB) {
	handler := &EmployeeHandler{
		dbInstance: db,
	}
	router.POST("/employee", handler.AddEmployee)
	router.GET("/employee", handler.GetEmployee)
}

func (h *EmployeeHandler) AddEmployee(c echo.Context) error {
	var e datamodels.Employee
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
	if mErr := h.dbInstance.AutoMigrate(&datamodels.Employee{}); mErr != nil {
		return c.JSON(http.StatusBadRequest, mErr)
	}
	e.ID = uuid.New()
	if err := h.dbInstance.Create(e).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}
func (h *EmployeeHandler) GetEmployee(c echo.Context) error {

	emailId := c.QueryParams().Get("email")
	if h.dbInstance == nil {
		return c.JSON(http.StatusBadRequest, "Connection Error")
	}
	var emp *datamodels.Employee
	// Get first matched record
	h.dbInstance.Where("email = ?", emailId).First(&emp)
	if emp == nil {
		return c.JSON(http.StatusForbidden, "Data not found")
	}
	return c.JSON(http.StatusOK, emp)

}
