package dto

type CreateCompanyRequestDto struct {
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	Rol         string `json:"Rol"`
	CompanyName string `json:"CompanyName"`
	Location    string `json:"Location"`
}
