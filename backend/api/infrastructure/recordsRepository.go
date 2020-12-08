package infrastructure

import "github.com/fullstacktf/ControlHorarios-Backend/api/models"

func CreateRecord(record models.EmployeeRecord) (error, int) {
	result := DB().Debug().Create(&record)
	return result.Error, record.RecordID
}
