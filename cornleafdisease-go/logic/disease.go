package logic

import (
	"project/cornleafdisease/dao/mysql"
	"project/cornleafdisease/models"
)

func GetSearch(name string) (data *models.Disease, err error) {
	data, err = mysql.GetDiseaseByName(name)
	return
}
