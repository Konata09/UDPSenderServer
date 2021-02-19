package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var jwtKey = []byte("sd*ust#konata&2O20")
var db *sql.DB

type ApiReturn struct {
	Retcode int         `json:"retcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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

func main() {
	initDBConn()
	mux := http.NewServeMux()
	mux.Handle("/api/v1/login", http.HandlerFunc(Login))
	mux.Handle("/api/v1/refresh", http.HandlerFunc(RefreshToken))
	mux.Handle("/api/v1/user/changepassword", VerifyHeader(http.HandlerFunc(Welcome)))
	mux.Handle("/api/v1/welcome", VerifyHeader(http.HandlerFunc(Welcome)))

	fmt.Println("UDPServer listen on 9999")
	log.Panic(http.ListenAndServe(":9999", mux))
}
