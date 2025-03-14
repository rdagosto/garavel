package services

import (
	"garavel/internal/libs"
	"garavel/internal/models"
	"garavel/internal/repositories"
	"garavel/internal/validators"
)

type User struct {
}

func (u User) Create(request *validators.UserStore) (*models.User, error) {
	var model models.User
	err := ToModel(&request, &model)
	if err != nil {
		return nil, err
	}

	model.Password, err = libs.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	repositories.Create(&model)

	return &model, nil
}

func (u User) Update(ID string, request *validators.UserUpdate) (*models.User, error) {
	model, err := Show[models.User](ID)
	if err != nil {
		return nil, err
	}

	err = ToModel(&request, &model)
	if err != nil {
		return nil, err
	}

	model.Password, err = libs.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	err = repositories.Save(&model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
