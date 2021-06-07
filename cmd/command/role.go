package main

import (
	"fmt"
	"strconv"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/database"
)

func main() {
	role, err := database.GetRole("00002", "00001")

	if err != nil {
		fmt.Println("Error!")
	} else {
		fmt.Println(strconv.Itoa(role))
	}
}