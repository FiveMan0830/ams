package server

import (
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type GetRoleRequest struct {
	TeamName string
	InputUserID string
}

func role(rg *gin.RouterGroup) {
	role := rg

	role.POST("/role/from/ams", getRoleFromAMS)
}

func getRoleFromAMS(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetRoleRequest{}

	c.Bind(reqbody)

	teamUUID, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), reqbody.TeamName)

	result, err := database.GetRole(reqbody.InputUserID, teamUUID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, result)
}