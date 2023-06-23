package handler

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type Employee struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

//	func (h *Handler) CreateUserByInvitation(c echo.Context, email string) error {
//		// send email, create entry in invitation table and
//
// }
func (h *Handler) CreateEmployee(c echo.Context) error {
	var e Employee
	err := c.Bind(&e)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	e.ID, _ = uuid.NewV4()
	if err := h.DB.Create(&e).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, e)
}
func (h *Handler) GetEmployee(c echo.Context) error {
	var e Employee
	h.DB.First(&e, "email=?", c.Param("email"))
	return c.JSON(http.StatusOK, e)
}
