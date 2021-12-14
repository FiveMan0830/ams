package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type dbConfig struct {
	Host     string
	Port     int
	User     string
	DBname   string
	Password string
}

func BuildDBConfig() *dbConfig {
	host := os.Getenv("AMS_MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("AMS_MYSQL_PORT"))
	if err != nil {
		panic(err)
	}
	user := os.Getenv("AMS_MYSQL_USER")
	password := os.Getenv("AMS_MYSQL_PASSWORD")
	database := os.Getenv("AMS_MYSQL_DATABASE")

	db := dbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBname:   database,
		Password: password,
	}
	return &db
}

func DbURL(dbConfig *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBname,
	)
}
