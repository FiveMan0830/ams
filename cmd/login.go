package main

import (
	"bufio"
	"os"
	"strings"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	accountManagement := account.NewLDAPManagement()

	adminUser, _ := reader.ReadString('\n')
	adminUser = strings.ReplaceAll(adminUser, "\n", "")
	adminUser = strings.ReplaceAll(adminUser, "\r", "")
	adminPasswd, _ := reader.ReadString('\n')
	adminPasswd = strings.ReplaceAll(adminPasswd, "\n", "")
	adminPasswd = strings.ReplaceAll(adminPasswd, "\r", "")

	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")
	password, _ := reader.ReadString('\n')
	password = strings.ReplaceAll(password, "\n", "")
	password = strings.ReplaceAll(password, "\r", "")



	accountManagement.Login(adminUser, adminPasswd, username, password)
}
