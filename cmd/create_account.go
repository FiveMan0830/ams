package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"

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

	userID, _ := reader.ReadString('\n')
	userID = strings.ReplaceAll(userID, "\n", "")
	userID = strings.ReplaceAll(userID, "\r", "")

	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")

	givenname, _ := reader.ReadString('\n')
	givenname = strings.ReplaceAll(givenname, "\n", "")
	givenname = strings.ReplaceAll(givenname, "\r", "")

	surname, _ := reader.ReadString('\n')
	surname = strings.ReplaceAll(surname, "\n", "")
	surname = strings.ReplaceAll(surname, "\r", "")

	password, _ := reader.ReadString('\n')
	password = strings.ReplaceAll(password, "\n", "")
	password = strings.ReplaceAll(password, "\r", "")

	email, _ := reader.ReadString('\n')
	email = strings.ReplaceAll(email, "\n", "")
	email = strings.ReplaceAll(email, "\r", "")

	fmt.Println("admin: " + adminUser + ", adminPasswd: " + adminPasswd + ", userName: " + username + ", surname: " + surname + ", password: " + password)

	accountManagement.AddUser(adminUser, adminPasswd, userID, username, givenname, surname, password, email)
}
