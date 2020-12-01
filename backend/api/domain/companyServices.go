package domain

import (
	"net/http"
	"strconv"

	"github.com/fullstacktf/ControlHorarios-Backend/api/controllers/dto"
	"github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"
	"github.com/fullstacktf/ControlHorarios-Backend/api/models"
	"github.com/gin-gonic/gin"
)

func CreateCompany(company models.UserCompany, c *gin.Context, id int) error {

	company.Company.UserID = id

	result := infrastructure.DB().Debug().
		Select(`Company.location, Company.company_name`).Create(&company.Company)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New company created successfully"})
	return nil
}

func GetCompany(id int) models.Company {
	var company models.Company
	infrastructure.DB().First(&company, id)

	return company
}

func CreateProject(id int, projectDto dto.ProjectDto) error {
	return infrastructure.InsertProject(id, projectDto.ProjectName)
}

func CreateHoliday(holidayCompany models.Holidays, c *gin.Context) error {

	holidayCompany.CompanyID, _ = strconv.Atoi(c.Params.ByName("id"))

	result := infrastructure.DB().Debug().Create(&holidayCompany)

	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New Holiday created successfully"})

	return nil
}

func CreateSection(section models.Sections, c *gin.Context) error {

	section.CompanyID, _ = strconv.Atoi(c.Params.ByName("id"))

	result := infrastructure.DB().Debug().Create(&section)
	if result.Error != nil {
		return result.Error
	}

	c.JSON(http.StatusOK, gin.H{"message": "New section created successfully"})
	return nil
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
