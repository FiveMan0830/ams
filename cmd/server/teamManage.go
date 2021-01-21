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
	Username string
}

type GetGroupsRequest struct {
	GroupList []string
}

type DeleteGroupRequest struct {
	GroupName string
}

type AddMemberRequest struct {
	GroupName string
	Username string
}

type GetMemberRequest struct {
	GroupName string
}

type GetLeaderRequest struct {
	GroupName string
}

type GetGroupOfMemberRequest struct {
	Username string
	GroupList []string
}


func main() {
	router := gin.Default()
	accountManagement := account.NewLDAPManagement()
	router.Static("/", "./web")

	router.POST("/create/team", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		reqbody := &CreateGroupRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(401, err)
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

	router.POST("/get/leader", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &GetLeaderRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		leaderList, err := accountManagement.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, leaderList)
	})

	router.POST("/get/groups/byuser", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        reqbody := &GetGroupOfMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		GroupList, err := accountManagement.SearchUserMemberOf(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, GroupList)
	})

	router.POST("/delete/team", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &DeleteGroupRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		err := accountManagement.DeleteGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(401, err)
			return
		}
	})

	router.POST("/add/member", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &AddMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.AddMemberToGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, memberList)
	})

	router.POST("/get/members", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &GetMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, memberList)
	})

	router.POST("/remove/member", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		reqbody := &AddMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.RemoveMemberFromGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(401, err)
			return
		}
		c.JSON(200, memberList)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}