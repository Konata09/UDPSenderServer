package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
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

func hexStringToByte(hexString string) ([]byte, error) {
	hex, err := hex.DecodeString(hexString)
	if err != nil {
		errMsg := fmt.Sprint(err)
		if strings.Contains(errMsg, "invalid byte") {
			return nil, errors.New("命令不能含有[0-9|A-F]以外的字符")
		} else if strings.Contains(errMsg, "odd length hex string") {
			return nil, errors.New("命令不能为奇数个字符")
		} else {
			return nil, err
		}
	}
	return hex, nil
}
