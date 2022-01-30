package config

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	appEnv     string
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string

	apiPort string
}

func Get() *Config {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Env not found")
	}

	config := &Config{}

	flag.StringVar(&config.appEnv, "appenv", getEnv("APP_ENV", "productionTest"), "Application Environment")

	flag.StringVar(&config.dbUser, "dbuser", getEnv("DB_USERNAME", "root"), "DB user name")
	flag.StringVar(&config.dbPassword, "dbpassword", getEnv("DB_PASSWORD", "password"), "DB pass")
	flag.StringVar(&config.dbPort, "dbport", getEnv("DB_PORT", "3306"), "DB port")
	flag.StringVar(&config.dbHost, "dbhost", getEnv("DB_HOST", "localhost"), "DB host")
	flag.StringVar(&config.dbName, "dbname", getEnv("DB_DATABASE", "microservice_sekolah"), "DB name")

	flag.StringVar(&config.apiPort, "apiPort", getEnv("API_PORT", "8080"), "API Port")
	flag.Parse()

	return config
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return fallback
}

func (conf *Config) GetAppEnv() string {
	return conf.appEnv
}

func (conf *Config) GetDBConnStr() string {
	return conf.getDBConnStr(conf.dbHost, conf.dbName)
}

func (conf *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local",
		conf.dbUser,
		conf.dbPassword,
		dbhost,
		conf.dbPort,
		dbname,
	)
}

func (conf *Config) GetAPIPort() string {
	return ":" + conf.apiPort
}
