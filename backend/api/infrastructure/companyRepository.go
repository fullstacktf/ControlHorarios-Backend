package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func GetCompanyId(id int) int {
	var company models.Company
	DB().Debug().Where("user_id = ?", id).First(&company)
	return company.CompanyID
}
