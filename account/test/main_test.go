package test

import (
	"fmt"
	"os"
	"testing"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

type user struct {
	ID        string
	name      string
	givenName string
	surname   string
	password  string
	email     string
}

type group struct {
	name                string
	groupLeaderUsername string
	ID                  string
}

type member struct {
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func setup() {

	user1 := user{"c61965be-8176-4419-b289-4d52617728fb", "fiveman123", "Richard", "Zheng", "richard123", "richardEmail@gmail.com"}
	user2 := user{"c56515be-1654-7895-1564-1d52513528cf", "audi98", "Audi", "Wu", "audipm984", "audiEmail@gmail.com"}

	leader1 := user{"a34531da-8563-9517-3578-3e38754896dg", "george88", "George", "Lim", "george189", "georgeEmail@gmail.com"}
	leader2 := user{"b96875kl-6842-7539-8549-2c56482648fa", "david93", "David", "Wang", "david632", "davidEmail@gmail.com"}
	leader3 := user{"c16475kl-6762-7489-8359-2c56846648fm", "sherry99", "Sherry", "Ye", "sherry753", "sherryEmail@gmail.com"}

	group1 := group{"OIS", "david93", "d23475kl-4862-7456-8473-2c53916648fn"}
	group2 := group{"SSL", "george88", "e16987kl-9512-7424-9629-2c56884248lm"}

	accountManagement := account.NewLDAPManagement()

	accountManagement.CreateUser(adminUser, adminPassword, user1.ID, user1.name, user1.givenName, user1.surname, user1.password, user1.email)
	accountManagement.CreateUser(adminUser, adminPassword, user2.ID, user2.name, user2.givenName, user2.surname, user2.password, user2.email)
	accountManagement.CreateUser(adminUser, adminPassword, userID3, username3, givenName3, surname3, userPassword3, userEmail3)

	accountManagement.CreateUser(adminUser, adminPassword, leader1.ID, leader1.name, leader1.givenName, leader1.surname, leader1.password, leader1.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader2.ID, leader2.name, leader2.givenName, leader2.surname, leader2.password, leader2.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader3.ID, leader3.name, leader3.givenName, leader3.surname, leader3.password, leader3.email)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)

	accountManagement.CreateGroup(adminUser, adminPassword, group1.name, group1.groupLeaderUsername, group1.ID)
	accountManagement.CreateGroup(adminUser, adminPassword, group2.name, group2.groupLeaderUsername, group2.ID)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	user1 := user{"c61965be-8176-4419-b289-4d52617728fb", "fiveman123", "Richard", "Zheng", "richard123", "richardEmail@gmail.com"}
	user2 := user{"c56515be-1654-7895-1564-1d52513528cf", "audi98", "Audi", "Wu", "audipm984", "audiEmail@gmail.com"}

	leader1 := user{"a34531da-8563-9517-3578-3e38754896dg", "george88", "George", "Lim", "george189", "georgeEmail@gmail.com"}
	leader2 := user{"b96875kl-6842-7539-8549-2c56482648fa", "david93", "David", "Wang", "david632", "davidEmail@gmail.com"}
	leader3 := user{"c16475kl-6762-7489-8359-2c56846648fm", "sherry99", "Sherry", "Ye", "sherry753", "sherryEmail@gmail.com"}

	group1 := group{"OIS", "david93", "d23475kl-4862-7456-8473-2c53916648fn"}
	group2 := group{"SSL", "george88", "e16987kl-9512-7424-9629-2c56884248lm"}

	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUser(adminUser, adminPassword, user1.name)
	accountManagement.DeleteUser(adminUser, adminPassword, user2.name)
	accountManagement.DeleteUser(adminUser, adminPassword, username3)

	accountManagement.DeleteUser(adminUser, adminPassword, leader1.name)
	accountManagement.DeleteUser(adminUser, adminPassword, leader2.name)
	accountManagement.DeleteUser(adminUser, adminPassword, leader3.name)

	accountManagement.DeleteGroup(adminUser, adminPassword, group1.name)
	accountManagement.DeleteGroup(adminUser, adminPassword, group2.name)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}
