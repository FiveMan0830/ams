package server

import (
	"github.com/gin-gonic/gin"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func RegisterV1Api(
	rg *gin.RouterGroup,
	am account.Management,
) {
	login(rg)
	teams(rg)
	auth(rg)
	role(rg)
	mySQL(rg)
	registerUserApi(rg, am)
}
