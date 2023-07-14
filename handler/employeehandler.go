package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"p3/datamodels"
	cErr "p3/err"
	"p3/utils"
	"strings"
)

type EmployeeHandler struct {
	dbInstance *gorm.DB
}

func NewEmployeeHandler(router *echo.Echo, db *gorm.DB) {
	handler := &EmployeeHandler{
		dbInstance: db,
	}
	router.POST("/employee", handler.AddEmployee)
	router.GET("/employee", handler.GetAllEmployees)
	router.PUT("/employee", handler.EditEmployee)
}

func (h *EmployeeHandler) AddEmployee(c echo.Context) error {
	var e datamodels.Employee
	err := c.Bind(&e)
	if err != nil {
		return utils.GetError(cErr.InvalidPostBody)
	}
	if err = c.Validate(e); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	if mErr := h.dbInstance.AutoMigrate(&datamodels.Employee{}); mErr != nil {
		return c.JSON(http.StatusBadRequest, mErr)
	}
	//e.ID = uuid.New()
	if err := h.dbInstance.Create(e).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, e)
}
func (h *EmployeeHandler) GetEmployee(c echo.Context) error {
	//emailId := c.QueryParams().Get("email")
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var emp *[]datamodels.Employee
	// Get first matched record
	h.dbInstance.Find(&emp)
	if emp == nil {
		return utils.GetError(cErr.DataNotFound)
	}
	return c.JSON(http.StatusOK, emp)

}
func (h *EmployeeHandler) GetAllEmployees(c echo.Context) error {
	email := c.Get("Email").(string)
	if email == "" {
		return utils.GetError(cErr.InvalidEmail)
	}
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var isAdmin bool
	var o *datamodels.Organization
	h.dbInstance.Where("created_by = ?", email).First(&o)
	if o.Name != "" && o.IsActive {
		isAdmin = true
	}
	var emp *[]datamodels.Employee
	// Get first matched record
	if isAdmin {
		h.dbInstance.Find(&emp, datamodels.Employee{OrgId: 123})
	} else {
		var managerData *datamodels.Employee
		h.dbInstance.Find(&emp, datamodels.Employee{Email: email})
		if managerData == nil || managerData.Email == "" {
			return c.JSON(http.StatusForbidden, "Invalid user access")
		}
		h.dbInstance.Find(&emp, datamodels.Employee{ManagedBy: email, OrgId: managerData.OrgId})
	}
	if emp == nil || len(*emp) == 0 {
		return utils.GetError(cErr.DataNotFound)
	}
	var returnEmployeeData []datamodels.GetEmployee
	for _, e := range *emp {
		returnEmp := datamodels.GetEmployee{
			Email:          e.Email,
			Name:           e.Name,
			ProfilePicture: e.ProfilePicture,
			ManagedBy:      e.ManagedBy,
		}
		returnEmployeeData = append(returnEmployeeData, returnEmp)
	}

	return c.JSON(http.StatusOK, returnEmployeeData)
}
func (h *EmployeeHandler) EditEmployee(c echo.Context) error {
	email := c.Get("Email").(string)
	if email == "" {
		return utils.GetError(cErr.InvalidEmail)
	}
	var editEmployee datamodels.EditEmployee
	err := c.Bind(&editEmployee)
	if err != nil {
		return utils.GetError(cErr.InvalidPostBody)
	}
	if err = c.Validate(editEmployee); err != nil {
		customError := cErr.CustomError{
			Status:  400,
			Message: err.Error(),
		}
		return utils.GetError(customError)
	}
	editEmployee.Email = strings.ToLower(editEmployee.Email)
	editEmployee.ManagedBy = strings.ToLower(editEmployee.ManagedBy)

	if editEmployee.Email == editEmployee.ManagedBy {
		return utils.GetError(cErr.EmailAndManageByNotSame)
	}
	if h.dbInstance == nil {
		return utils.GetError(cErr.ConnectionError)
	}
	var o *datamodels.Organization
	h.dbInstance.Where("created_by = ?", email).First(&o)
	if o == nil || o.Name == "" {
		return utils.GetError(cErr.AccessDenied)
	}
	var emp *datamodels.Employee
	var manager *datamodels.Employee
	// Get first matched record
	h.dbInstance.Find(&emp, datamodels.Employee{OrgId: 123, Email: editEmployee.Email})
	if emp == nil || emp.Email == "" {
		return utils.GetError(cErr.EmployeeNotFound)
	}
	h.dbInstance.Find(&manager, datamodels.Employee{OrgId: 123, Email: editEmployee.ManagedBy})
	if manager == nil || manager.Email == "" {
		return utils.GetError(cErr.ManagerNotFound)
	}
	if manager.ManagedBy == emp.Email {
		return utils.GetError(cErr.UpdateError)
	}
	emp.ManagedBy = strings.ToLower(editEmployee.ManagedBy)
	h.dbInstance.Save(emp)
	returnEmp := datamodels.GetEmployee{
		Email:          emp.Email,
		Name:           emp.Name,
		ProfilePicture: emp.ProfilePicture,
		ManagedBy:      emp.ManagedBy,
	}
	return c.JSON(http.StatusOK, returnEmp)

}
