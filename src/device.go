package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Device struct {
	DeviceId   int    `json:"id"`
	DeviceName string `json:"name"`
	DeviceIp   string `json:"ip"`
	DeviceMac  string `json:"mac"`
	DeviceUdp  bool   `json:"udp"`
	DeviceWol  bool   `json:"wol"`
}

type AllDevices struct {
	Count   int      `json:"count"`
	Devices []Device `json:"devices"`
}

func SetDevice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		devices := getDevices()
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: 0,
			Message: "OK",
			Data: &AllDevices{
				Count:   len(devices),
				Devices: devices,
			},
		})
	case "PUT":
		var body AllDevices
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		var devices []Device
		var msg string
		for _, dev := range body.Devices {
			if strings.TrimSpace(dev.DeviceName) == "" || (strings.TrimSpace(dev.DeviceIp) == "" && strings.TrimSpace(dev.DeviceMac) == "") {
				msg = msg + "Incomplete item found "
			} else {
				devices = append(devices, dev)
			}
		}
		if len(devices) == 0 {
			ApiErrMsg(w, msg+"No item to add")
			return
		}
		ok := addDevice(devices)
		if ok {
			ApiOkMsg(w, msg+"OK")
		} else {
			ApiErrMsg(w, msg+"请求错误")
		}
	case "POST":
		var body Device
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		if getDeviceById(body.DeviceId) == nil {
			ApiErrMsg(w, "设备不存在")
			return
		}
		ok := setDevice(body.DeviceId, body.DeviceName, body.DeviceIp, body.DeviceMac, body.DeviceUdp, body.DeviceWol)
		if ok {
			ApiOk(w)
		} else {
			ApiErr(w)
		}
	case "DELETE":
		var body Device
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		if getDeviceById(body.DeviceId) == nil {
			ApiErrMsg(w, "设备不存在")
			return
		}
		ok := deleteDevice(body.DeviceId)
		if ok {
			ApiOk(w)
		} else {
			ApiErr(w)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
