package repositories

import (
	"garavel/internal/models"
)

func Find[M any](list []M) {
	models.GetDB().Find(&list)
}

func Create[M any](model M) {
	models.GetDB().Create(&model)
}

func GetByID[M any](model M, ID interface{}) error {
	return models.GetDB().First(&model, ID).Error
}

func Save[M any](model M) error {
	return models.GetDB().Save(&model).Error
}

func Delete[M any](model M) {
	models.GetDB().Delete(&model)
}
