package main

import (
	"database/sql"
	"fmt"
	"log"
)

func initDBConn() {
	var err error
	db, err = sql.Open("sqlite3", "db.db?cache=shared&mode=wrc")
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		fmt.Printf("Error open database.\n")
	}
}
// 0:admin 1:user -1:error
func CheckUserByPass(username string, password string) int {
	stmt, err := db.Prepare("select uid, password, isadmin from user,role where username = ? and user.roleid = role.roleid")
	if err != nil {
		log.Fatal(err)
		return -1
	}
	defer stmt.Close()
	var uid int
	var PasswordMD5 string
	var isadmin bool
	err = stmt.QueryRow(username).Scan(&uid, &PasswordMD5, &isadmin)
	if err != nil {
		return -1
	}

	if !checkMD5(password,PasswordMD5) {
		return -1
	} else {
		if isadmin {
			return 0
		} else {
			return 1
		}
	}

}
