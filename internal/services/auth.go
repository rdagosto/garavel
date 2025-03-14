package services

import (
	"garavel/internal/libs"
	"garavel/internal/models"
	"garavel/internal/validators"
)

type Auth struct {
}

func (u Auth) Login(request *validators.UserStore) (*string, error) {
	var dbUser models.User
	err := models.GetDB().Where("email = ?", request.Email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	if !libs.CheckHash(dbUser.Password, request.Password) {
		return nil, err
	}

	token, err := libs.GenerateJWT(dbUser.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
