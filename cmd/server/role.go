package server

import (

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type GetRoleRequest struct {
	TeamName string
	InputUserID string
}

func role(rg *gin.RouterGroup) {
	role := rg

	role.POST("/roleFromAMS", getRoleFromAMS)
}

func getRoleFromAMS(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetRoleRequest{}

	c.Bind(reqbody)

	result, err := accountManagement.SearchUserRole(reqbody.TeamName, reqbody.InputUserID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, result)
}