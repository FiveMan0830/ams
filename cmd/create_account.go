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
	username, _ := reader.ReadString('\n')
	username = strings.ReplaceAll(username, "\n", "")
	username = strings.ReplaceAll(username, "\r", "")
	surname, _ := reader.ReadString('\n')
	surname = strings.ReplaceAll(surname, "\n", "")
	surname = strings.ReplaceAll(surname, "\r", "")
	password, _ := reader.ReadString('\n')
	password = strings.ReplaceAll(password, "\n", "")
	password = strings.ReplaceAll(password, "\r", "")

	accountManagement.AddUser(username, surname, password)
}
