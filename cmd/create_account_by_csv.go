package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	filename = strings.ReplaceAll(filename, "\n", "")
	filename = strings.ReplaceAll(filename, "\n", "")
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	adminUser := config.GetAdminUser()
	adminPassword := config.GetAdminPassword()

	accountManagement := account.NewLDAPManagement()

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		userID := record[0]
		username := record[1]
		givenname := record[2]
		surname := record[3]
		password := record[4]
		email := record[5]

		accountManagement.AddUser(adminUser, adminPassword, userID, username, givenname, surname, password, email)
	}
}
