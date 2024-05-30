package logic

import (
	"math"
	"project/cornleafdisease/dao/mysql"
	"project/cornleafdisease/models"
	"time"
)

func GetRecordList(page, size, userID int64) (data []*models.ApiRecord, err error) {
	data, err = mysql.GetRecordList(page, size, userID)
	return
}

func GetCount(userid int64) (num int64, err error) {
	num, err = mysql.GetNumber(userid)
	return
}

func SaveToRecord(result map[float64]string, userid int64) (err error) {
	var maxKey float64 = math.Inf(-1)
	name := ""
	// 遍历map
	for key, value := range result {
		// 如果当前键比最大键大，则更新最大键和对应的字符串
		if key > maxKey {
			maxKey = key
			name = value
		}
	}
	disease, err := mysql.GetDiseaseByName(name)
	if err != nil {
		return
	}
	record := &models.Record{
		UserID:     userid,
		Confidence: maxKey,
		DiseaseId:  disease.ID,
		CreateTime: time.Now(),
	}

	err = mysql.SaveRecord(record)
	return
}
