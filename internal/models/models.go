package models

import (
	"garavel/internal/configs"
	"garavel/internal/databases"
	"log"

	"gorm.io/gorm"
)

const (
	UserClass     = "user"
	CustomerClass = "customer"
)

type Model interface {
	GetDrive() string
}

func Make(class string) Model {
	switch class {
	case UserClass:
		return &User{}
	case CustomerClass:
		return &Customer{}
	default:
		log.Fatalf("❌ Unsupported model type: %s", class)
		return nil
	}
}

func GetDB(model ...Model) *gorm.DB {
	if len(model) == 0 || model[0] == nil {
		return databases.GetDB(configs.Config("database.default", ""))
	}
	return databases.GetDB(model[0].GetDrive())
}
