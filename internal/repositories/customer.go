package repositories

import (
	"garavel/internal/models"
)

type Customer struct {
}

func (f Customer) GetModel() models.Customer {
	return models.Customer{}
}

func (f Customer) GetList() []models.Customer {
	return []models.Customer{}
}
