package models

type Disease struct {
	ID                 int64
	Alias              string
	Introduction       string `gorm:"type:longtext"`
	DiseaseName        string
	OnsetConditions    string `gorm:"type:longtext"`
	PathogenicFeatures string `gorm:"type:longtext"`
	SymptomsOfHarm     string `gorm:"type:longtext"`
	Treatment          string `gorm:"type:longtext"`
}

func (Disease) TableName() string {
	return "disease"
}
