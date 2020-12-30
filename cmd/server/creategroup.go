package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type CreateGroupRequest struct {
	GroupName string
}

type GetGroupsRequest struct {
	GroupList []string
}

func main() {
	router := gin.Default()
	accountManagement := account.NewLDAPManagement()
	router.Static("/", "./web")

	router.POST("/create/team", func(c *gin.Context) {
        // c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &CreateGroupRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(401,err)
			return
		}
		c.JSON(200, info)
	})

	router.POST("/get/groups", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		GroupList, err := accountManagement.GetGroups(config.GetAdminUser(), config.GetAdminPassword())
		reqbody := &GetGroupsRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, GroupList)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}