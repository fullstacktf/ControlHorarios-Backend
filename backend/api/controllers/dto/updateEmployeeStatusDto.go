package dto

type UpdateEmployeeStatusDto struct {
	UserID int  `json:"UserID"`
	Status bool `json:"Status"`
}
