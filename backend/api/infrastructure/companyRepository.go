package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func GetCompanyId(id int) int {
	var company models.Company
	DB().Debug().Where("user_id = ?", id).First(&company)
	return company.CompanyID
}

func CreateCompany(company models.Company) error {
	result := DB().Debug().Create(&company)
	return result.Error
}

func GetCompany(id int) models.Company {
	var company models.Company
	DB().Debug().First(&company, id)
	return company
}
