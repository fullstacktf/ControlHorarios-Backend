package dto

type CreateEmployeeRequestDto struct {
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Rol       string `json:"Rol"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName "`
}
