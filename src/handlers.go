package main

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	//if VerifyHeader(r.Header) {
	w.Write([]byte(trimMACtoShow("5410ECA0FA12")))
	//} else {
	//	w.WriteHeader(403)
	//}
}
