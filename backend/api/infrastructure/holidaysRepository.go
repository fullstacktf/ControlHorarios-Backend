package infrastructure

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func UpdateHolidaysName(id int, name string, date string) error {
	result := DB().Debug().Model(&models.Holidays{}).Where("holiday_id = ?", id).Update("holiday_title", name).Update("holiday_date", date)
	return result.Error
}

func DeleteHolidays(id int) error {
	result := DB().Debug().Where("holiday_id = ?", id).Delete(models.Holidays{})
	return result.Error
}

func CreateHolidays(holidays models.Holidays) error {
	result := DB().Debug().Create(&holidays)
	return result.Error
}

func GetHolidaysByCompanyID(id int) []models.Holidays {
	var holidays []models.Holidays
	DB().Debug().Where("company_id = ?", id).Find(&holidays)
	return holidays
}
