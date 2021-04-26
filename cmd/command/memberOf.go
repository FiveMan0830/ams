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

	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")

	surname, _ := reader.ReadString('\n')
	surname = strings.ReplaceAll(surname, "\n", "")
	surname = strings.ReplaceAll(surname, "\r", "")

	givenname, _ := reader.ReadString('\n')
	givenname = strings.ReplaceAll(givenname, "\n", "")
	givenname = strings.ReplaceAll(givenname, "\r", "")

	fmt.Println(" username: " + username)
	result, err := accountManagement.CreateGroup(adminUser, adminPasswd, username, surname, givenname)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
