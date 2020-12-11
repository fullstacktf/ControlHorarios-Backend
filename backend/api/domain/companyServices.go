package domain

import (
	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
)

func CreateCompany(companyDto dto.CreateCompanyRequestDto) (error, int) {
	user := models.User{Username: companyDto.Username, Password: companyDto.Password, Email: companyDto.Email, Rol: companyDto.Rol}
	id := infrastructure.CreateUser(user)
	company := models.Company{UserID: id, CompanyName: companyDto.CompanyName, Location: companyDto.Location}
	return infrastructure.CreateCompany(company)
}

func GetCompany(id int) models.Company {
	return infrastructure.GetCompany(id)
}

func CreateProject(id int, projectDto dto.ProjectDto) error {
	return infrastructure.InsertProject(id, projectDto.ProjectName)
}

func CreateHoliday(holidayDto dto.CreateHolidaysRequestDto, companyId int) error {
	holidays := models.Holidays{HolidayTitle: holidayDto.Title, HolidayDate: holidayDto.Date, CompanyID: companyId}
	return infrastructure.CreateHolidays(holidays)
}

func CreateSection(sectionDto dto.CreateSectionDto, companyId int) error {
	return infrastructure.CreateSection(models.Sections{SectionName: sectionDto.SectionName, CompanyID: companyId})
}

func GetHolidays(id int) []models.Holidays {
	return infrastructure.GetHolidaysByCompanyID(id)
}

func GetEmployees(id int) []models.Employee {
	return infrastructure.GetEmployeesByCompanyID(id)
}

func GetProjects(id int) []models.Projects {
	return infrastructure.GetProjectsByCompanyID(id)
}

func GetSections(id int) []models.Sections {
	return infrastructure.GetSectionsByCompanyID(id)
}

func UpdateProjectName(id int, projectDto dto.ProjectDto) error {
	return infrastructure.UpdateProjectName(id, projectDto.ProjectName)
}

func UpdateSectionName(id int, sectionDto dto.UpdateSectionsRequestDto) error {
	return infrastructure.UpdateSectionName(id, sectionDto.SectionName)
}

func UpdateHolidaysName(id int, holidaysDto dto.UpdateHolidaysRequestDto) error {
	return infrastructure.UpdateHolidaysName(id, holidaysDto.HolidaysName, holidaysDto.NewDate)
}

func DeleteHolidays(id int) error {
	return infrastructure.DeleteHolidays(id)
}

func UpdateEmployee(employeeStatusDto dto.UpdateEmployeeStatusDto) error {
    return  infrastructure.UpdateEmployee(employeeStatusDto.UserID)
}
