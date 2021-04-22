package test

import (
	"fmt"
	"os"
	"testing"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/account"
)

type user struct {
	ID string
	name string
	givenName string
	surname string
	password string
	email string
}

type group struct {
	name string
	groupLeaderUsername string
	ID string
}

type member struct {
	Username string `json:"username"`
	Displayname string `json:"displayname"`
}

const adminUser string = "admin"
const adminPassword string = "admin"

// User test data
const null string = ""
const usernameNotExists string = "Eric"

const userID string = "c61965be-8176-4419-b289-4d52617728fb"
const username string = "fiveman123"
const givenName string = "Richard"
const surname string = "Zheng"
const userPassword string = "richard123"
const userEmail string = "richardEmail@gmail.com"

const userID2 string = "c56515be-1654-7895-1564-1d52513528cf"
const username2 string = "audi98"
const givenName2 string = "Audi"
const surname2 string = "Wu"
const userPassword2 string = "audipm984"
const userEmail2 string = "audiEmail@gmail.com"

const userID3 string = "a56515be-5783-8738-1564-1d52513528cz"
const username3 string = "stella83"
const givenName3 string = "Stella"
const surname3 string = "Chen"
const userPassword3 string = "stella423"
const userEmail3 string = "stellaEmail@gmail.com"

const leaderID string = "a34531da-8563-9517-3578-3e38754896dg"
const leaderUsername string = "george88"
const leaderGivenName string = "George"
const leaderSurname string = "Lim"
const leaderPassword string = "george189"
const leaderEmail string = "georgeEmail@gmail.com"

const leaderID2 string = "b96875kl-6842-7539-8549-2c56482648fa"
const leaderUsername2 string = "david93"
const leaderGivenName2 string = "David"
const leaderSurname2 string = "Wang"
const leaderPassword2 string = "david632"
const leaderEmail2 string = "davidEmail@gmail.com"

const leaderID3 string = "c16475kl-6762-7489-8359-2c56846648fm"
const leaderUsername3 string = "sherry99"
const leaderGivenName3 string = "Sherry"
const leaderSurname3 string = "Ye"
const leaderPassword3 string = "sherry753"
const leaderEmail3 string = "sherryEmail@gmail.com"

// Group test data
const groupNameNotExists = "Sunbird"
const groupLeaderNotExists string = "Ron"

const groupName string = "OIS"
const groupLeaderUsername string = "david93"
const groupID string = "d23475kl-4862-7456-8473-2c53916648fn"

const groupName2 string = "SSL"
const groupLeaderUsername2 string = "george88"
const groupID2 string = "e16987kl-9512-7424-9629-2c56884248lm"

const groupName3 string = "KanBan"
const groupLeaderUsername3 string = "sherry88"
const groupID3 string = "f16846kl-7862-4539-1684-3a56884248lo"

// Ou test data
const ouName = "LabGroup"
const ouName2 = "NTUT"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
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
	// group3 := group{"KanBan", "sherry88", "f16846kl-7862-4539-1684-3a56884248lo"}

	accountManagement := account.NewLDAPManagement()
	accountManagement.CreateUser(adminUser, adminPassword, user1.ID, user1.name, user1.givenName, user1.surname, user1.password, user1.email)
	accountManagement.CreateUser(adminUser, adminPassword, user2.ID, user2.name, user2.givenName, user2.surname, user2.password, user2.email)
	accountManagement.CreateUser(adminUser, adminPassword, userID3,  username3, givenName3, surname3, userPassword3, userEmail3)

	accountManagement.CreateUser(adminUser, adminPassword, leader1.ID, leader1.name, leader1.givenName, leader1.surname, leader1.password, leader1.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader2.ID, leader2.name, leader2.givenName, leader2.surname, leader2.password, leader2.email)
	accountManagement.CreateUser(adminUser, adminPassword, leader3.ID, leader3.name, leader3.givenName, leader3.surname, leader3.password, leader3.email)

	accountManagement.CreateOu(adminUser, adminPassword, ouName)
	accountManagement.CreateOu(adminUser, adminPassword, "OISGroup")
	// accountManagement.CreateOu(adminUser, adminPassword, ouName2)

	accountManagement.CreateGroup(adminUser, adminPassword, group1.name, group1.groupLeaderUsername, group1.ID)
	accountManagement.CreateGroup(adminUser, adminPassword, group2.name, group2.groupLeaderUsername, group2.ID)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	accountManagement := account.NewLDAPManagement()

	accountManagement.DeleteUser(adminUser, adminPassword, username)
	accountManagement.DeleteUser(adminUser, adminPassword, username2)
	accountManagement.DeleteUser(adminUser, adminPassword, username3)

	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername)
	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername2)
	accountManagement.DeleteUser(adminUser, adminPassword, leaderUsername3)

	accountManagement.DeleteGroup(adminUser, adminPassword, groupName)
	accountManagement.DeleteGroup(adminUser, adminPassword, groupName2)

	accountManagement.DeleteOu(adminUser, adminPassword, ouName)
	// accountManagement.DeleteOu(adminUser, adminPassword, ouName2)

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
}