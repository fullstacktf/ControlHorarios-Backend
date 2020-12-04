package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func UpdateProjectName(id int, name string) error {
	result := DB().Debug().Model(&models.Projects{}).Where("project_id = ?", id).Update("project_name", name)
	return result.Error
}

func InsertProject(id int, name string) error {
	project := models.Projects{ProjectName: name, CompanyID: id}
	result := DB().Debug().Create(&project)
	return result.Error
}
