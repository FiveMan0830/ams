package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB * gorm.DB

type dbConfig struct {
	Host string
	Port int
	User string
	DBname string
	Password string
}

func BuildDBConfig() *dbConfig {
	db := dbConfig{
		Host: "",
		Port: 0,
		User: "",
		DBname: "",
		Password: "",
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


