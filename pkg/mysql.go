package pkg

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlClient() *gorm.DB {
	host := os.Getenv("AMS_MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("AMS_MYSQL_PORT"))
	if err != nil {
		panic("failed to get ams mysql port from environment variable")
	}
	user := os.Getenv("AMS_MYSQL_USER")
	password := os.Getenv("AMS_MYSQL_PASSWORD")
	database := os.Getenv("AMS_MYSQL_DATABASE")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
