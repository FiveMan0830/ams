package server

import (
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
)

type roleRelationRequest struct {
	teamID string
	userID string
}

func mySQL(rg *gin.RouterGroup) {
	db := rg

	db.POST("/role", getRole)
}

func getRole(c *gin.Context) {
	reqbody := &roleRelationRequest{}

	c.Bind(reqbody)

	result, err := database.GetRole(reqbody.userID, reqbody.teamID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, result)
}