package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var envLoaded bool = false
var configLoaded bool = false

func getConfigPath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	isTest := EnvInt("IS_TEST", "0") == 1
	var configPath string
	if isTest {
		configPath = wd + "/../internal/configs"
	} else {
		configPath = wd + "/internal/configs"
	}
	return configPath
}

func loadConfig() {
	configLoaded = true

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getConfigPath())
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func Config(key string, fallback string) string {
	if !configLoaded {
		loadConfig()
	}
	value := viper.GetString(key)
	if value == "" {
		return fallback
	}
	return value

}

func loadEnv() {
	envLoaded = true
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Env(key string, fallback string) string {
	if !envLoaded {
		loadEnv()
	}
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func EnvInt(key string, fallback string) int {
	value, _ := strconv.Atoi(Env(key, fallback))
	return value
}
