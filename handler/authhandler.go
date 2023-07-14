package handlers

import (
	"net/http"
	"p3/datamodels"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHandler struct {
	dbInstance *gorm.DB
}

func NewAuthHandler(router *echo.Echo, db *gorm.DB) {
	handler := &AuthHandler{
		dbInstance: db,
	}
	router.GET("/login", handler.Login)
}

func (h *AuthHandler) Login(c echo.Context) error {

	emailId := c.Get("Email")
	if h.dbInstance == nil {
		return c.JSON(http.StatusBadRequest, "Connection Error")
	}

	var emp *datamodels.Employee
	// Get first matched record
	// Check if employee already exist in db
	h.dbInstance.Where("email = ?", emailId).First(&emp)
	// If employee does not exist, create a new entry for employee
	if emp == nil {

		if mErr := h.dbInstance.AutoMigrate(&datamodels.Employee{}); mErr != nil {
			return c.JSON(http.StatusBadRequest, mErr)
		}
		//emp.ID = uuid.New()                                   // new UUID
		emp.Email = c.Get("Email").(string)                   // Email from token validation
		emp.Name = c.Get("Name").(string)                     // Name from token validation
		emp.ProfilePicture = c.Get("ProfilePicture").(string) // Profile Pic from token validation
		emp.Role = "Admin"                                    // update role to admin to show "add organization" component
		if err := h.dbInstance.Create(emp).Error; err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusOK, emp)
}
