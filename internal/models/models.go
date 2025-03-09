package models

import (
	"garavel/internal/configs"
	"garavel/internal/databases"
	"log"

	"gorm.io/gorm"
)

type Model interface {
	GetDrive() string
}

func Factory(model string) Model {
	switch model {
	case "customer":
		return &Customer{}
	case "user":
		return &User{}
	default:
		log.Fatalf("❌ Unsupported model type: %s", model)
		return nil
	}
}

func GetDB(model ...Model) *gorm.DB {
	if len(model) == 0 || model[0] == nil {
		return databases.GetDB(configs.Config("database.default", ""))
	}
	return databases.GetDB(model[0].GetDrive())
}
