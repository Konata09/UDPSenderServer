package main

func getUidByUsernameAndPassword(username string, password string) int {
	passwordMD5 := getPasswordMD5(password)
	stmt, err := db.Prepare("select uid from user where username = ? and password = ?")
	if err != nil {
		return -1
	}
	defer stmt.Close()
	var uid int
	err = stmt.QueryRow(username, passwordMD5).Scan(&uid)
	if err != nil {
		return -1
	}
	return uid
}

func getUidByUsername(username int) int {
	stmt, err := db.Prepare("select uid from user where username = ?")
	if err != nil {
		return -1
	}
	defer stmt.Close()
	var uid int
	err = stmt.QueryRow(username).Scan(&uid)
	if err != nil {
		return -1
	}
	return uid
}

func getRoleByUid(uid int, role *Role) *Role {
	stmt, err := db.Prepare("select rolename, isadmin from user,role where uid = ? and user.roleid = role.roleid")
	if err != nil {
		return nil
	}
	defer stmt.Close()
	var isadmin bool
	var rolename string
	err = stmt.QueryRow(uid).Scan(&rolename, &isadmin)
	if err != nil {
		return nil
	}
	role.Rolename = rolename
	role.Isadmin = isadmin
	return role
}

func GetPasswordByUid(uid int) (res string, err error) {
	stmt, err := db.Prepare("select password from user where uid = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var passMD5 string
	err = stmt.QueryRow(uid).Scan(&passMD5)
	if err != nil {
		return "", err
	}
	return passMD5, nil
}

func SetPasswordByUid(uid int, newPassword string) bool {
	stmt, err := db.Prepare("update user set password = ? where uid = ?")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(newPassword, uid)
	if err != nil {
		return false
	}
	return true
}

func GetUsers() []User {
	stmt, err := db.Prepare("select uid, username, rolename, isadmin from user,role where user.roleid = role.roleid")
	if err != nil {
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Uid, &user.Username, &user.Rolename, &user.Isadmin)
		users = append(users, user)
	}
	return users
}

func AddUser(username string, password string, roleid int) bool {
	stmt, err := db.Prepare("insert into user (username, password, roleid) values (?, ?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, roleid)
	if err != nil {
		return false
	}
	return true
}

func deleteUser(uid int) bool {
	stmt, err := db.Prepare("delete from user where uid = ?")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid)
	if err != nil {
		return false
	}
	return true
}
