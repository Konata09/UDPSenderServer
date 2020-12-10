package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var jwtKey = []byte("sd*ust#konata&2O20")
var db *sql.DB
type ApiReturn struct{
	Retcode int `json:"retcode"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func main() {
	initDBConn()

	http.HandleFunc("/api/v1/login", Login)
	http.HandleFunc("/api/v1/welcome", Welcome)
	http.HandleFunc("/api/v1/refresh", RefreshToken)

	fmt.Println("UDPServer listen on 9999")
	log.Panic(http.ListenAndServe(":9999",nil))
}