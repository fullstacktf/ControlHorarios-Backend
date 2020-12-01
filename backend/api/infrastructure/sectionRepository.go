package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func UpdateSectionName(id int, name string) error {
	result := DB().Debug().Model(&models.Sections{}).Where("section_id = ?", id).Update("section_name", name)
	return result.Error
}
