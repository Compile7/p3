package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"p3/datamodels"
	"p3/logic"
)

type EmployeeHandler struct {
	eLogic *logic.EmployeeLogic
}

func NewEmployeeHandler(router *echo.Echo, logic *logic.EmployeeLogic) {
	handler := &EmployeeHandler{
		eLogic: logic,
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
	/*err := json.NewDecoder(c.Request().Body).Decode(&e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}*/
	res, eErr := h.eLogic.AddEmployee(&e)
	if eErr != "" {
		return c.JSON(http.StatusBadRequest, eErr)
	}
	return c.JSON(http.StatusOK, res)
}
func (h *EmployeeHandler) GetEmployee(c echo.Context) error {

	emailId := c.QueryParams().Get("email")
	res, err := h.eLogic.GetEmployee(emailId)
	if res != nil {
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusForbidden, err)
}
