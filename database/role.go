package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

func InsertRole(userID string, teamID string, role int) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("error :", err)
	}

	defer db.Close()

	db.Exec("INSERT INTO `role_relation` VALUE(?, ?, ?)", teamID, userID, role)
}

func GetTeamLeader(teamID string) (string, error) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("error :", err)
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT `unit_id` FROM `role_relation` WHERE `team_id` = ? AND `role` = 1")

	if err != nil {
		log.Println("error :", err)
	}

	defer stmt.Close()

	var leaderID string

	err = stmt.QueryRow(teamID).Scan(&leaderID)

	if err != nil {
		log.Println("error :", err)
	}

	return leaderID, nil
}

func UpdateLeader(oldLeaderID, newLeaderID, teamID string) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("error :", err)
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT `unit_id` FROM `role_relation` WHERE `team_id` = ?")

	if err != nil {
		log.Println("error :", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(teamID)

	defer rows.Close()

	oldLeader := false
	newLeader := false

	for rows.Next() {
		var id string

		if err := rows.Scan(&id); err != nil {
			log.Println("error :", err)
		}

		if id == oldLeaderID {
			oldLeader = true
		}

		if id == newLeaderID {
			newLeader = true
		}
	}

	if oldLeader == true && newLeader == true {
		db.Exec("UPDATE `role_relation` SET `role` = 1 WHERE `unit_id` = ? AND `team_id` = ?", newLeaderID, teamID)
		db.Exec("UPDATE `role_relation` SET `role` = 0 WHERE `unit_id` = ? AND `team_id` = ?", oldLeaderID, teamID)
	}
}

func DeleteTeam(teamID string) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	log.Println(teamID)
	if err != nil {
		log.Println("error :", err)
	}

	defer db.Close()

	db.Exec("DELETE FROM `role_relation` WHERE `team_id` = ?", teamID)
}

func DeleteRole(userID, teamID string) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("error :", err)
	}

	defer db.Close()

	db.Exec("DELETE FROM `role_relation` WHERE `team_id` = ? AND `unit_id` = ?", teamID, userID)
}

func GetRole(userID, teamID string) (int, error) {
	db, err := sql.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		log.Println("error :", err)
		return 5, err
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT `role` FROM `role_relation` WHERE `unit_id` = ? AND `team_id` = ?")

	if err != nil {
		log.Println("error :", err)
		return 5, err
	}

	defer stmt.Close()

	var role int

	err = stmt.QueryRow(userID, teamID).Scan(&role)

	fmt.Printf("in team %s, user %s is a %s\n", teamID, userID, strconv.Itoa(role))

	if err != nil {
		log.Println("error :", err)
		return 5, err
	}

	return role, nil
}
