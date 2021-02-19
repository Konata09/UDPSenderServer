package main

import (
	"crypto/md5"
	"encoding/hex"
)

var salt = "ustsdnicm3002"

func getPasswordMD5(password string) string {
	data := []byte(password + salt)
	encodePass := md5.Sum(data)
	return hex.EncodeToString(encodePass[:])
}

func checkPassword(passIn string, passMD5 string) bool {
	data := []byte(passIn + salt)
	encodedPass := md5.Sum(data)
	return hex.EncodeToString(encodedPass[:]) == passMD5
}
