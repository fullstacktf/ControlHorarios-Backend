package domain

import "github.com/fullstacktf/ControlHorarios-Backend/api/infrastructure"

var userRepository infrastructure.UserRepositoryInterface
var employeeRepository infrastructure.EmployeeRepositoryInterface
var companyRepository infrastructure.CompanyRepositoryInterface

func SetRepositories(userRep infrastructure.UserRepositoryInterface,
	employeeRep infrastructure.EmployeeRepositoryInterface,
	companyRep infrastructure.CompanyRepositoryInterface) {
	userRepository = userRep
	employeeRepository = employeeRep
	companyRepository = companyRep
}
