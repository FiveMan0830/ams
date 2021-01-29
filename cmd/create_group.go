package main

import (
	"bufio"
	"fmt"
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

	groupName, _ := reader.ReadString('\n')
	groupName = strings.ReplaceAll(groupName, "\n", "")
	groupName = strings.ReplaceAll(groupName, "\r", "")

	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")

	fmt.Println("adminUser: " + adminUser + " adminPasswd: " + adminPasswd + " groupName: " + groupName + " groupID: " + "username: " + username)
	groupname,err:=accountManagement.CreateGroup(adminUser, adminPasswd, groupName, username)
	if (err == nil) {
		fmt.Println(groupname)
	}
}
