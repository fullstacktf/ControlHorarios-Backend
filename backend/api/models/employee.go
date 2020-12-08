package models

type Employee struct {
	EmployeeID     int  `gorm:"column:employee_id;primaryKey"`
	User           User `gorm:"foreignKey:UserID"`
	UserID         int
	CompanyID      int
	FirstName      string   `gorm:"column:first_name;type:varchar(50);"`
	Company        Company  `gorm:"foreignKey:CompanyID"`
	LastName       string   `gorm:"column:last_name;type:varchar(50);"`
	ProfilePicture byte     `gorm:"column:profile_picture;type:blob"`
	Projects       Projects `gorm:"foreignKey:ProjectID"`
}

func (Employee) TableName() string {
	return "employee"
}
