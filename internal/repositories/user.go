package repositories

import (
	"garavel/internal/models"
)

type User struct {
}

func (f User) GetModel() models.User {
	return models.User{}
}

func (f User) GetList() []models.User {
	return []models.User{}
}
