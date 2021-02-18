package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"regexp"
	"time"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     int    `json:"role"` // 0:admin 1:user
	jwt.StandardClaims
}

type ResponseToken struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	role := CheckUserByPass(creds.Username, creds.Password)
	if role == -1 {
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: -1,
			Message: "Wrong username or password",
		})
		return
	}

	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%s Login success: %s\n", time.Now().Format(time.UnixDate), creds.Username)
	json.NewEncoder(w).Encode(&ApiReturn{
		Retcode: 0,
		Message: "OK",
		Data: &ResponseToken{
			Token: tokenString,
		},
	})
}

func VerifyHeader(header http.Header) bool {
	re := regexp.MustCompile(`Bearer\s(.*)$`)

	headerAuth := header.Get("Authorization")
	if len(headerAuth) == 0 {
		return false
	}
	tknStr := re.FindStringSubmatch(headerAuth)

	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Printf("%d", claims.ExpiresAt)
	if err != nil || !tkn.Valid {
		return false
	}
	return true
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`Bearer\s(.*)$`)
	headerAuth := r.Header.Get("Authorization")
	if len(headerAuth) == 0 {
		return
	}
	tknStr := re.FindStringSubmatch(headerAuth)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 24*time.Hour {
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: -1,
			Message: "Token not expires in one day",
		})
		return
	}
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&ApiReturn{
		Retcode: 0,
		Message: "OK",
		Data: &ResponseToken{
			Token: tokenString,
		},
	})
}
