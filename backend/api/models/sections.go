package models

type Sections struct {
	SectionID   int     `gorm:"column:section_id;primaryKey"`
	SectionName string  `gorm:"column:section_name;type:varchar(50)"`
	Company     Company `gorm:"foreignKey:CompanyID"`
}

func (Sections) TableName() string {
	return "sections"
}
