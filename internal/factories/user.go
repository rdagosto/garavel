package factories

import (
	"garavel/internal/libs"
	"garavel/internal/models"

	"syreclabs.com/go/faker"
)

type User struct {
}

func (f User) Make() models.User {
	password, _ := libs.Hash(faker.Internet().Password(5, 5))
	return models.User{
		Email:    faker.Internet().Email(),
		Password: password,
	}
}
