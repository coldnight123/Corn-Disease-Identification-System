package mysql

import (
	"fmt"
	"project/cornleafdisease/models"
)

func GetRecordList(page, size, userID int64) (records []*models.ApiRecord, err error) {
	offset := (page - 1) * size
	limit := size
	tx := DB.Table("record").
		Select("record.id as ID,record.confidence,disease.disease_name, disease.alias, record.create_time").
		Joins("JOIN disease ON record.disease_id = disease.id").
		Where("record.user_id = ?", userID).
		Order("record.create_time DESC").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&records)
	err = tx.Error
	return
}

func SaveRecord(record *models.Record) (err error) {
	create := DB.Create(&record)
	err = create.Error
	return
}

func GetNumber(userid int64) (count int64, err error) {
	tx := DB.Model(&models.Record{}).Where("user_id=?", userid).Count(&count)
	err = tx.Error
	fmt.Println(count)
	return
}
