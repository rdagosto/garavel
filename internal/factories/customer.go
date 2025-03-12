package factories

import (
	"garavel/internal/models"

	"syreclabs.com/go/faker"
)

type Customer struct {
}

func (f Customer) Make() models.Customer {
	return models.Customer{
		Name:  faker.Name().Name(),
		Email: faker.Internet().Email(),
	}
}
