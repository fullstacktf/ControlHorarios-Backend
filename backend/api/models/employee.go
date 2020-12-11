package models

type Employee struct {
	EmployeeID     int `gorm:"column:employee_id;primaryKey;"`
	UserID         int
	User           User     `gorm:"foreignKey:UserID;"`
	FirstName      string   `gorm:"column:first_name;type:varchar(50);"`
	LastName       string   `gorm:"column:last_name;type:varchar(50);"`
	ProfilePicture byte     `gorm:"column:profile_picture;type:blob;"`
	Projects       Projects `gorm:"foreignKey:ProjectID;"`
	CompanyID      int
	Company        Company `gorm:"foreignKey:CompanyID;"`
}

func (Employee) TableName() string {
	return "employee"
}
