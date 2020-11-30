package models

type Employee struct {
	EmployeeID     int  `gorm:"column:employee_id;primaryKey"`
	User           User `gorm:"foreignKey:UserID"`
	UserID         int
	CompanyID      int
	FirstName      string   `gorm:"column:first_name;type:varchar(50); NOT NULL json:"first_name"`
	Company        Company  `gorm:"foreignKey:CompanyID"`
	LastName       string   `gorm:"column:last_name;type:varchar(50); NOT NULL json:"last_name"`
	ProfilePicture byte     `gorm:"column:profile_picture;type:blob"`
	Projects       Projects `gorm:"foreignKey:ProjectID"`
}

type UserEmployee struct {
	User     User     `json:"User"`
	Employee Employee `json:"Employee"`
}

func (Employee) TableName() string {
	return "employee"
}
