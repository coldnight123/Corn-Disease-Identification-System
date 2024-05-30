package mysql

import "project/cornleafdisease/models"

func GetDiseaseByName(name string) (data *models.Disease, err error) {
	var disease models.Disease
	tx := DB.Where("disease_name = ?", name).First(&disease)
	err = tx.Error
	if tx.RowsAffected == 0 {
		return nil, ErrorDiseaseNotExist
	}
	return &disease, err
}
