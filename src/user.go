package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Rolename string `json:"rolename"`
	Isadmin  bool   `json:"isadmin"`
}

type Role struct {
	Rolename string
	Isadmin  bool
}

type UserChangePasswordBody struct {
	OldPass string `json:"oldpass"`
	NewPass string `json:"newpass"`
}

type AdminChangePasswordBody struct {
	Uid     int    `json:"uid"`
	NewPass string `json:"newpass"`
}

type AllUsers struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}

func UserChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := getUserInfoFromJWT(r)
	var body UserChangePasswordBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		returnErr(w)
		return
	}
	oldPass, err := GetPasswordByUid(user.Uid)
	if err != nil {
		returnErr(w)
		return
	}
	if oldPass != getPasswordMD5(body.OldPass) {
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: -1,
			Message: "Wrong password",
		})
		return
	}
	ok := SetPasswordByUid(user.Uid, getPasswordMD5(body.NewPass))
	if ok {
		returnOk(w)
	} else {
		returnErrMsg(w, "修改失败")
	}
}

func AdminChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var body AdminChangePasswordBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		returnErr(w)
		return
	}
	ok := SetPasswordByUid(body.Uid, getPasswordMD5(body.NewPass))
	if ok {
		returnOk(w)
	} else {
		returnErrMsg(w, "修改失败")
	}
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users := GetUsers()
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: 0,
			Message: "OK",
			Data: &AllUsers{
				Count: len(users),
				Users: users,
			},
		})
	case "PUT":

	case "DELETE":

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
