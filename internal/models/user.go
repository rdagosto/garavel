package models

import (
	"garavel/internal/databases"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *User) GetDrive() string {
	return databases.MysqlDB
}
