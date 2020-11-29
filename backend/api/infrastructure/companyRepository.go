package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func InsertProject(id int, name string) error {
	project := models.Projects{ProjectName: name, CompanyID: id}
	result := DB().Debug().Create(&project)
	return result.Error
}