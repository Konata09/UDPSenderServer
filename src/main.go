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

func main() {
	initDBConn()
	mux := http.NewServeMux()
	mux.Handle("/api/v1/login", http.HandlerFunc(Login))
	mux.Handle("/api/v1/refresh", http.HandlerFunc(RefreshToken))
	mux.Handle("/api/v1/welcome", VerifyHeader(http.HandlerFunc(Welcome)))

	fmt.Println("UDPServer listen on 9999")
	log.Panic(http.ListenAndServe(":9999", mux))
}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}
