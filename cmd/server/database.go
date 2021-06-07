package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
)

func mySQL(rg *gin.RouterGroup) {
	db := rg

	db.POST("/role", getRole)
}

func getRole(c *gin.Context) {
	result, err := database.GetRole("00002", "00001")

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	fmt.Println(strconv.Itoa(result))

	c.JSON(200, result)

}