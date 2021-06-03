package server

import (

	"github.com/gin-gonic/gin"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	_ "github.com/go-sql-driver/mysql"
)

func database(rg *gin.RouterGroup) {
	db := rg

	db.POST("", )
}