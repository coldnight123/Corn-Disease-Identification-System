package models

import "time"

type Record struct {
	ID         int64     `json:"ID"`
	UserID     int64     `json:"UserID"`
	Confidence float64   `json:"Confidence"`
	DiseaseId  int64     `json:"DiseaseId"`
	CreateTime time.Time `json:"CreateTime"`
	Disease    Disease   `gorm:"foreignKey:DiseaseId"`
}

func (Record) TableName() string {
	return "record"
}

type ApiRecord struct {
	ID          int64     `json:"id"`
	DiseaseName string    `json:"DiseaseName"`
	CreateTime  time.Time `json:"CreateTime"`
	Alias       string
}
