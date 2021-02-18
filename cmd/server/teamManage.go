package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"io/ioutil"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type CreateGroupRequest struct {
	GroupName string
	Username  string
}

type GetGroupsRequest struct {
	GroupList []string
}

type DeleteGroupRequest struct {
	GroupName string
}

type AddMemberRequest struct {
	GroupName string
	Username  string
}

type GetMemberRequest struct {
	GroupName string
}

type GetLeaderRequest struct {
	GroupName string
}

type GetGroupOfMemberRequest struct {
	Username string
}

type GetUUIDByUsernameRequest struct {
	Username string
}

func main() {
	router := gin.Default()
	accountManagement := account.NewLDAPManagement()
	router.Static("/", "./web")

	router.POST("/create/team", func(c *gin.Context) {
		reqbody := &CreateGroupRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, info)
	})

	router.POST("/get/groups", func(c *gin.Context) {
		GroupList, err := accountManagement.GetGroups(config.GetAdminUser(), config.GetAdminPassword())
		reqbody := &GetGroupsRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, GroupList)
	})

	router.POST("/get/leader", func(c *gin.Context) {
		reqbody := &GetLeaderRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		leaderList, err := accountManagement.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, leaderList)
	})

	router.POST("/get/groups/byuser", func(c *gin.Context) {
		reqbody := &GetGroupOfMemberRequest{}
		c.Bind(reqbody)

		log.Println(reqbody)
		log.Println(reqbody.Username)

		GroupList, err := accountManagement.SearchUserMemberOf(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)

		log.Println(GroupList)

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, GroupList)
	})

	router.POST("/delete/team", func(c *gin.Context) {
		reqbody := &DeleteGroupRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		err := accountManagement.DeleteGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(500, err)
			return
		}
	})

	router.POST("/add/member", func(c *gin.Context) {
		reqbody := &AddMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.AddMemberToGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, memberList)
	})

	router.POST("/get/members", func(c *gin.Context) {
		reqbody := &GetMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, memberList)
	})

	router.POST("/remove/member", func(c *gin.Context) {
		reqbody := &AddMemberRequest{}
		c.Bind(reqbody)
		log.Println(reqbody)
		memberList, err := accountManagement.RemoveMemberFromGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, memberList)
	})

	router.POST("/get/uuid", func(c *gin.Context) {
		// reqbody := &GetUUIDByUsernameRequest{}
		// c.Bind(reqbody)
		x, err := ioutil.ReadAll(c.Request.Body)
		log.Println(x)
		uuid, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), string(x))

		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, uuid)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
