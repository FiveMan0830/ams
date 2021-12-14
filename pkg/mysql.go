package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysqlClient() *gorm.DB {
	host := os.Getenv("AMS_MYSQL_HOST")
	port, err := strconv.Atoi(os.Getenv("AMS_MYSQL_PORT"))
	if err != nil {
		fmt.Println(err)
		panic("failed to get ams mysql port from environment variable")
	}
	user := os.Getenv("AMS_MYSQL_USER")
	password := os.Getenv("AMS_MYSQL_PASSWORD")
	database := os.Getenv("AMS_MYSQL_DATABASE_V2")

	fmt.Println("host", host)
	fmt.Println("port", port)
	fmt.Println("user", user)
	fmt.Println("password", password)
	fmt.Println("database", database)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	return db
}
