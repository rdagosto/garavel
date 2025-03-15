package databases

import (
	"fmt"
	"garavel/internal/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	mmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func getMigrationPath() string {
	if configs.EnvInt("IS_TEST", "0") == 1 {
		return "file://../internal/migrations"
	}
	return "file://internal/migrations"
}

func (m *MySQLDatabase) RunMigration() {
	sqlDB, err := m.dbInstance.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB from Gorm: %v", err)
	}

	driver, err := mmysql.WithInstance(sqlDB, &mmysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create MySQL migration driver: %v", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(getMigrationPath(), "mysql", driver)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	err = mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully!")
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	if m.dbInstance == nil {
		log.Fatal("Database not initialized. Call Connect() first.")
	}
	return m.dbInstance
}
