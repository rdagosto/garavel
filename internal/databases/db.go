package databases

import (
	"garavel/internal/configs"
	"log"

	"gorm.io/gorm"
)

const (
	MysqlDB = "mysql"
)

type Database interface {
	Connect() error
	GetDB() *gorm.DB
	RunMigration()
}

var databases map[string]Database

func init() {
	databases = make(map[string]Database)
}

func connect(dbType string) Database {
	var db Database
	switch dbType {
	case MysqlDB:
		db = &MySQLDatabase{}
	default:
		log.Fatalf("❌ Unsupported database type: %s", dbType)
	}

	err := db.Connect()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	if configs.EnvInt("IS_TEST", "0") == 0 {
		db.RunMigration()
	}

	return db
}

func GetDB(dbType string) *gorm.DB {
	if databases[dbType] == nil {
		databases[dbType] = connect(dbType)
	}
	return databases[dbType].GetDB()
}
