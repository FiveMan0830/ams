package database

import (
	"strconv"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func GetRole(userID, teamID string) (int, error) {
	db, err := sql.Open("mysql",config.DbURL(config.BuildDBConfig()))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT `role` FROM `role_relation` WHERE `unit_id` = ? AND `team_id` = ?")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	var role int

	err = stmt.QueryRow(userID, teamID).Scan(&role)

	fmt.Println(strconv.Itoa(role))

	if err != nil {
		panic(err.Error())
	}

	return role, nil
}

// func connectDatabase() (*DB, error) {
// 	db, err := sql.Open("mysql",config.DbURL(config.BuildDBConfig()))

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	err = db.Ping()

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return db, nil
// }