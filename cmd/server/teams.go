package server

import (
	"github.com/google/uuid"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"

	_ "ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

type GetGroupRequest struct {
	GroupName string
	Username  string
}

type GetGroupsListRequest struct {
	GroupList []string
}

type GetGroupsRequest struct {
	GroupName string
}

type GetUsersRequest struct {
	Username string
}

type AddMemberRequest struct {
	GroupName string
	Username  string
}

func teams(rg *gin.RouterGroup) {
	team := rg.Group("/")

	team.POST("/team/create", createTeam)
	team.POST("/team/get", getTeam)
	team.POST("/team/get/leader", getTeamLeader)
	team.POST("/team/get/memberOf", getTeamMember)
	team.POST("/team/get/uuid/user", getUUIDOfUser)
	team.POST("/team/get/uuid/team", getUUIDOfTeam)
	team.POST("/team/get/Name", getName)
	team.POST("/team/delete", deleteTeam)
	team.POST("/team/add/member", addMember)
	team.POST("/team/leader/handover", handoverLeader)
}

func createTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)
	teamID := uuid.New().String()
	info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username, teamID)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, info)
}

func getTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupsRequest{}
	c.Bind(reqbody)
	GroupList, err := accountManagement.GetGroups(config.GetAdminUser(), config.GetAdminPassword())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, GroupList)
}

func getTeamLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupsRequest{}
	c.Bind(reqbody)
	leaderList, err := accountManagement.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, leaderList)
}

func getTeamMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetUsersRequest{}
	c.Bind(reqbody)
	GroupList, err := accountManagement.SearchUserMemberOf(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, GroupList)
}

func getUUIDOfUser(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	uuid, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, uuid)
}

func getUUIDOfTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	uuid, err := accountManagement.SearchGroupUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, uuid)
}

func getName(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody, err := ioutil.ReadAll(c.Request.Body)
	name, err := accountManagement.SearchNameByUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, name)
}

func deleteTeam(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupsRequest{}
	c.Bind(reqbody)
	err := accountManagement.DeleteGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

	if err != nil {
		c.JSON(500, err)
		return
	}
}

func addMember(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &AddMemberRequest{}
	c.Bind(reqbody)
	memberList, err := accountManagement.AddMemberToGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, memberList)
}

func handoverLeader(c *gin.Context) {
	accountManagement := account.NewLDAPManagement()
	reqbody := &GetGroupRequest{}
	c.Bind(reqbody)
	err := accountManagement.UpdateGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}
}

// func main() {
// 	router := gin.Default()
// 	accountManagement := account.NewLDAPManagement()
// 	router.Static("/", "./web")

// 	router.POST("/create/team", func(c *gin.Context) {
// 		reqbody := &CreateGroupRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		teamID := uuid.New().String()
// 		info, err := accountManagement.CreateGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username, teamID)

// 		if err != nil {
// 			c.JSON(500, err.Error())
// 			return
// 		}
// 		c.JSON(200, info)
// 	})

// 	router.POST("/get/groups", func(c *gin.Context) {
// 		GroupList, err := accountManagement.GetGroups(config.GetAdminUser(), config.GetAdminPassword())
// 		reqbody := &GetGroupsRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, GroupList)
// 	})

// 	router.POST("/get/leader", func(c *gin.Context) {
// 		reqbody := &GetLeaderRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		leaderList, err := accountManagement.SearchGroupLeader(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, leaderList)
// 	})

// 	router.POST("/get/groups/byuser", func(c *gin.Context) {
// 		reqbody := &GetGroupOfMemberRequest{}
// 		c.Bind(reqbody)

// 		log.Println(reqbody)
// 		log.Println(reqbody.Username)

// 		GroupList, err := accountManagement.SearchUserMemberOf(config.GetAdminUser(), config.GetAdminPassword(), reqbody.Username)

// 		log.Println(GroupList)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, GroupList)
// 	})

// 	router.POST("/delete/team", func(c *gin.Context) {
// 		reqbody := &DeleteGroupRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		err := accountManagement.DeleteGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 	})

// 	router.POST("/add/member", func(c *gin.Context) {
// 		reqbody := &AddMemberRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		memberList, err := accountManagement.AddMemberToGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

// 		if err != nil {
// 			c.JSON(500, err.Error())
// 			return
// 		}
// 		c.JSON(200, memberList)
// 	})

// 	router.POST("/get/members", func(c *gin.Context) {
// 		reqbody := &GetMemberRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		memberList, err := accountManagement.GetGroupMembers(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, memberList)
// 	})

// 	router.POST("/remove/member", func(c *gin.Context) {
// 		reqbody := &AddMemberRequest{}
// 		c.Bind(reqbody)
// 		log.Println(reqbody)
// 		memberList, err := accountManagement.RemoveMemberFromGroup(config.GetAdminUser(), config.GetAdminPassword(), reqbody.GroupName, reqbody.Username)

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, memberList)
// 	})

// 	router.POST("/get/uuid", func(c *gin.Context) {
// 		// reqbody := &GetUUIDByUsernameRequest{}
// 		// c.Bind(reqbody)
// 		reqbody, err := ioutil.ReadAll(c.Request.Body)
// 		log.Println(reqbody)
// 		uuid, err := accountManagement.GetUUIDByUsername(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, uuid)
// 	})

// 	router.POST("/get/teamUid", func(c *gin.Context) {
// 		// reqbody := &GetUUIDByUsernameRequest{}
// 		// c.Bind(reqbody)
// 		reqbody, err := ioutil.ReadAll(c.Request.Body)
// 		log.Println(reqbody)
// 		uuid, err := accountManagement.SearchGroupUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

// 		if err != nil {
// 			c.JSON(500, err)
// 			return
// 		}
// 		c.JSON(200, uuid)
// 	})

// 	router.POST("/get/unitdto", func(c *gin.Context) {
// 		reqbody, err := ioutil.ReadAll(c.Request.Body)
// 		name, err := accountManagement.SearchNameByUUID(config.GetAdminUser(), config.GetAdminPassword(), string(reqbody))

// 		if err != nil {
// 			c.JSON(500, err)
// 			log.Println(err)
// 			return
// 		}
// 		c.JSON(200, name)
// 	})

// 	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }