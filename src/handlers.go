package main

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	//if VerifyHeader(r.Header) {
	//hex, err := hexStringToByte("4C696768746F6EFE051562017DFFr")
	//if err != nil {
	//	w.Write([]byte(fmt.Sprint(err)))
	//	return
	//}
	//sendUdp("192.168.13.4", 888, hex)
	w.Write([]byte(trimMACtoShow("5410ECA0FA12")))
	//} else {
	//	w.WriteHeader(403)
	//}
}
