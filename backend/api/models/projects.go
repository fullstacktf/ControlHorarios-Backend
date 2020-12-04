package models

type Projects struct {
	ProjectID   int    `gorm:"column:;primaryKey"`
	ProjectName string `gorm:"column:project_name;type:varchar(50)`
	CompanyID   int    `gorm:"foreignKey:CompanyID"`
}

func (Projects) TableName() string {
	return "projects"
}
