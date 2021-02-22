package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(getSubNetBroadcast("172.31.161.200", 24))
	w.Write([]byte(trimMACtoShow("5410ECA0FA12")))
	//} else {
	//	w.WriteHeader(403)
	//}
}
