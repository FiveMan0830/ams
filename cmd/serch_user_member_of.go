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

	memberName, _ := reader.ReadString('\n')
	memberName = strings.ReplaceAll(groupName, "\n", "")
	memberName = strings.ReplaceAll(groupName, "\r", "")

	fmt.Println("adminUser: " + adminUser + " adminPasswd: " + adminPasswd + " memberName: " + memberName)
	result, err := accountManagement.SearchUserMemberOf(adminUser, adminPasswd, memberName)
	if (err == nil) {
		fmt.Println(result)
	}
}