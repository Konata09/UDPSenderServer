package main

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(trimMACtoShow("5410ECA0FA12")))
	//} else {
	//	w.WriteHeader(403)
	//}
}
