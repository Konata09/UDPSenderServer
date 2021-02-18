package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	if VerifyHeader(r.Header) {
		w.Write([]byte(fmt.Sprintf("Welcome %s", "asd")))
	} else {
		w.WriteHeader(403)
	}
}
