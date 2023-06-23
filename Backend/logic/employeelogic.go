package logic

import (
	"github.com/google/uuid"
	"p3/datamodels"
	"p3/db"
)

type EmployeeLogic struct {
}

func NewEmployeeLogic() *EmployeeLogic {
	return &EmployeeLogic{}
}

func (logic *EmployeeLogic) AddEmployee(e *datamodels.Employee) (*datamodels.Employee, string) {
	dbInstance := db.GetInstance()
	if dbInstance == nil {
		return nil, "Connection Error"
	}
	dbInstance.AutoMigrate(&datamodels.Employee{})
	e.ID = uuid.New()
	if err := dbInstance.Create(e).Error; err != nil {
		return nil, err.Error()
	}
	return e, ""
}

func (logic *EmployeeLogic) GetEmployee(email string) (*datamodels.Employee, string) {
	dbInstance := db.GetInstance()
	if dbInstance == nil {
		return nil, "Connection Error"
	}
	var emp *datamodels.Employee
	// Get first matched record
	dbInstance.Where("email = ?", email).First(&emp)
	if emp == nil {
		return nil, "Data not found"
	}
	return emp, ""
}
