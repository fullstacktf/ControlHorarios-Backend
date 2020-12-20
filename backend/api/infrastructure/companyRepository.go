package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

type CompanyRepositoryInterface interface {
	GetCompanyId(id int) (error, int)
	CreateCompany(company models.Company) (error, int)
	GetCompany(id int) models.Company
}

type CompanyRepository struct{}

func (c CompanyRepository) GetCompanyId(id int) (error, int) {
	var company models.Company
	result := DB().Debug().Where("user_id = ?", id).First(&company)
	return result.Error, company.CompanyID
}

func (c CompanyRepository) CreateCompany(company models.Company) (error, int) {
	result := DB().Debug().Create(&company)
	return result.Error, company.CompanyID
}

func (c CompanyRepository) GetCompany(id int) models.Company {
	var company models.Company
	DB().Debug().Joins("User").First(&company, id)
	return company
}
