package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
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

func trimMACtoShow(in string) string {
	if in == "" {
		return in
	}
	s := strings.ToUpper(strings.Replace(strings.Replace(in, ":", "", -1), "-", "", -1))
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", s[0:2], s[2:4], s[4:6], s[6:8], s[8:10], s[10:12])
}
func trimMACtoStor(in string) string {
	return strings.ToUpper(strings.Replace(strings.Replace(in, ":", "", -1), "-", "", -1))
}
func trimCommandToStor(in string) string {
	return strings.ToUpper(in)
}
