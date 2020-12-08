package infrastructure

import (
	"time"

	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func CreateRecord(record models.EmployeeRecord) (error, int) {
	result := DB().Debug().Create(&record)
	return result.Error, record.RecordID
}

func UpdateRecordTimeOut(recordID int, timeOut time.Time) error {
	result := DB().Debug().
		Model(&models.EmployeeRecord{}).
		Where("record_id = ?", recordID).
		Update("end_time", timeOut)
	return result.Error
}

func GetRecordsByID(id int) []models.EmployeeRecord {
	var records []models.EmployeeRecord
	DB().
		Select("record_id", "description", "start_time", "end_time", "employee_id").
		Where("employee_id = ?", id).
		Find(&records)
	return records
}
