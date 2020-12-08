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
	var holidays []models.Holidays
	infrastructure.DB().Debug().Where("company_id = ?", id).Find(&holidays)
	return holidays
}

func GetEmployees(id int) []models.Employee {
	var employees []models.Employee
	infrastructure.DB().Debug().Where("company_id").Find(&employees)
	return employees
}

func GetProjects(id int) []models.Projects {
	var projects []models.Projects
	infrastructure.DB().Debug().Where("company_id").Find(&projects)
	return projects
}

func GetSections(id int) []models.Sections {
	var sections []models.Sections
	infrastructure.DB().Debug().Where("company_id").Find(&sections)
	return sections
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
