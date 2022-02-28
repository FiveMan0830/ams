package account

import (
	"errors"
	"fmt"
	"log"

	ldap "github.com/go-ldap/ldap/v3"

	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/config"
)

// CreateOu is a function for user to create ou
func (lm *LDAPManagement) CreateOu(adminUser, adminPasswd, ouname string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	addReq := ldap.NewAddRequest(fmt.Sprintf("ou=%s,%s", ouname, config.GetDC()), []ldap.Control{})

	addReq.Attribute("objectClass", []string{"top", "organizationalUnit"})
	addReq.Attribute("ou", []string{ouname})

	if err := conn.Add(addReq); err != nil {
		return errors.New("this Organization Unit already exists")
	}

	return nil
}

// DeleteOu is a function for user to delete ou
func (lm *LDAPManagement) DeleteOu(adminUser, adminPasswd, ouname string) error {
	conn, _ := lm.getConnectionWithoutTLS()
	defer conn.Close()
	lm.bindAuth(conn, adminUser, adminPasswd)

	d := ldap.NewDelRequest(fmt.Sprintf("ou=%s,%s", ouname, lm.BaseDN), nil)
	err := conn.Del(d)

	if err != nil {
		log.Println("Organization Unit entry could not be deleted :", err)
		return err
	}

	return nil
}
