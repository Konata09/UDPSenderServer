package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	uid      int
	username string
	isadmin  bool
}

type Role struct {
	rolename string
	isadmin  bool
}

type UserChangePasswordBody struct {
	OldPass string `json:"oldpass"`
	NewPass string `json:"newpass"`
}

type AdminChangePasswordBody struct {
	Uid     int    `json:"uid"`
	NewPass string `json:"newpass"`
}

func UserChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := getUserInfoFromJWT(r)
	var body UserChangePasswordBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		returnErr(w)
		return
	}
	oldPass, err := GetPasswordByUid(user.uid)
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
	ok := SetPasswordByUid(user.uid, body.NewPass)
	if ok {
		returnOk(w)
	} else {
		returnErr(w)
	}
}

func AdminChangePassword(w http.ResponseWriter, r *http.Request) {

}
