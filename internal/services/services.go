package services

import (
	"garavel/internal/repositories"

	"github.com/jinzhu/copier"
)

func ToModel[O any, M any](obj O, model M) error {
	err := copier.Copy(&model, &obj)
	if err != nil {
		return err
	}
	return nil
}

func List[M any]() []M {
	var list []M
	repositories.Find(list)
	return list
}

func Create[M any, R any](request R) (*M, error) {
	var model M
	err := ToModel(&request, &model)
	if err != nil {
		return nil, err
	}

	repositories.Create(&model)
	return &model, nil
}

func Show[M any](ID string) (*M, error) {
	var model M
	err := repositories.GetByID(&model, ID)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func Update[M any, R any](ID string, request R) (*M, error) {
	model, err := Show[M](ID)
	if err != nil {
		return nil, err
	}

	err = ToModel(&request, &model)
	if err != nil {
		return nil, err
	}

	err = repositories.Save(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func Delete[M any](ID string) error {
	model, err := Show[M](ID)
	if err != nil {
		return err
	}

	repositories.Delete(&model)
	return nil
}
