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

	ouName, _ := reader.ReadString('\n')
	ouName = strings.ReplaceAll(ouName, "\n", "")
	ouName = strings.ReplaceAll(ouName, "\r", "")

	fmt.Println("adminUser: " + adminUser + " adminPasswd: " + adminPasswd + " ouName: " + ouName)
	accountManagement.DeleteOu(adminUser, adminPasswd, ouName)
}
