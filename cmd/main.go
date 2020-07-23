package main

import (
	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/ldapconn"
)


func main()  {
	
	l := ldapconn.ConnectLDAP()
	defer l.Close()
	ldapconn.BindLDAP(l)
	// ldapconn.AddGroup(l)
	// ldapconn.AddUser(l)

	
}