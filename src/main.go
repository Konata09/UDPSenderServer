package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var jwtKey = []byte("sd*ust#konata&2O20")
var db *sql.DB

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
	mux.Handle("/api/v1/welcome", VerifyHeader(http.HandlerFunc(Welcome)))
	mux.Handle("/api/v1/user/changePassword", VerifyHeader(http.HandlerFunc(UserChangePassword)))
	//mux.Handle("/api/v1/getDevice", VerifyHeader(http.HandlerFunc(UserChangePassword)))
	//mux.Handle("/api/v1/getCommand", VerifyHeader(http.HandlerFunc(UserChangePassword)))
	//mux.Handle("/api/v1/sendCommand", VerifyHeader(http.HandlerFunc(UserChangePassword)))
	mux.Handle("/api/v1/admin/changePassword", VerifyHeader(VerifyAdmin(http.HandlerFunc(AdminChangePassword))))
	mux.Handle("/api/v1/admin/setUser", VerifyHeader(VerifyAdmin(http.HandlerFunc(SetUser))))
	mux.Handle("/api/v1/admin/setDevice", VerifyHeader(VerifyAdmin(http.HandlerFunc(SetDevice))))
	mux.Handle("/api/v1/admin/setCommand", VerifyHeader(VerifyAdmin(http.HandlerFunc(SetCommand))))

	//fmt.Println(getPasswordMD5("admin"))
	fmt.Println("UDPServer listen on 9999")
	log.Panic(http.ListenAndServe(":9999", mux))
}
