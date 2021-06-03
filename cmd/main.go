package main

import (
	"fmt"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/cmd/server"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	defer config.DB.Close()
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	server.Run()
}