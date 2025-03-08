package models

import (
	"time"
)

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *Customer) GetDrive() string {
	return "mysql"
}
