package main

import (
	"log"
)

func getUidByUsernameAndPassword(username string, password string) int {
	passwordMD5 := getPasswordMD5(password)
	stmt, err := db.Prepare("select uid from user where username = ? and password = ?")
	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer stmt.Close()
	var uid int
	err = stmt.QueryRow(username, passwordMD5).Scan(&uid)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return uid
}

func getRoleByUid(uid int, role *Role) *Role {
	stmt, err := db.Prepare("select rolename, isadmin from user,role where uid = ? and user.roleid = role.roleid")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()
	var isadmin bool
	var rolename string
	err = stmt.QueryRow(uid).Scan(&rolename, &isadmin)
	if err != nil {
		return nil
	}
	role.rolename = rolename
	role.isadmin = isadmin
	return role
}

func SetPasswdByUid(uid int, newPassword string) bool {
	stmt, err := db.Prepare("update user set password = ? where uid = ?")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(newPassword, uid)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
