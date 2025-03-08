package databases

import (
	"fmt"
	"garavel/internal/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	dbInstance *gorm.DB
}

func (m *MySQLDatabase) Connect() error {
	charset := configs.Config("database.mysql.charset", "")
	parseTime := configs.Config("database.mysql.parseTime", "")
	loc := configs.Config("database.mysql.loc", "")
	host := configs.Env("DB_HOST", "localhost")
	port := configs.Env("DB_PORT", "3306")
	user := configs.Env("DB_USER", "")
	password := configs.Env("DB_PASSWORD", "")
	database := configs.Env("DB_DATABASE", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", user, password, host, port, database, charset, parseTime, loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	m.dbInstance = db
	return nil
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	if m.dbInstance == nil {
		log.Fatal("Database not initialized. Call Connect() first.")
	}
	return m.dbInstance
}
