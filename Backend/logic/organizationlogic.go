package logic

import (
	"p3/datamodels"
	"p3/db"
)

type OrganizationLogic struct {
}

func NewOrganizationLogic() *OrganizationLogic {
	return &OrganizationLogic{}
}

func (logic *OrganizationLogic) AddOrg(o *datamodels.Organization) (*datamodels.Organization, string) {
	dbInstance := db.GetInstance()
	if dbInstance == nil {
		return nil, "Connection Error"
	}
	dbInstance.AutoMigrate(&datamodels.Organization{})
	o.IsActive = true
	o.IsDeleted = false
	if o.CreatedBy == "" {
		o.CreatedBy = "hemant.manwani@loginradius.com"
	}
	if err := dbInstance.Create(o).Error; err != nil {
		return nil, err.Error()
	}
	return o, ""
}

func (logic *OrganizationLogic) GetOrg(email string) (*datamodels.Organization, string) {
	dbInstance := db.GetInstance()
	if dbInstance == nil {
		return nil, "Connection Error"
	}
	var o *datamodels.Organization
	// Get first matched record
	dbInstance.Where("created_by = ?", email).First(&o)
	if o == nil {
		return nil, "Data not found"
	}
	return o, ""
}
