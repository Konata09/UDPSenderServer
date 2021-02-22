package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SendWolPacket struct {
	TargetId          []int `json:"targetid"`
	DevAddress        bool  `json:"devaddress"`
	LocalNetBroadcast bool  `json:"localnetbroadcast"`
	SubNetBroadcast   bool  `json:"subnetbroadcast"`
	Repeat            int   `json:"repeat"`
}

func SendWOL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var body SendWolPacket
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		if !body.DevAddress && !body.LocalNetBroadcast && !body.SubNetBroadcast {
			ApiErrMsg(w, "至少选择一个发送地址")
			return
		}

		var repeat int

		if body.Repeat == 0 {
			repeat = 1
		} else if body.Repeat > 5 {
			repeat = 5
		} else {
			repeat = body.Repeat
		}

		var errMsg string
		for i := 0; i < repeat; i++ {
			for _, tar := range body.TargetId {
				var address []string
				dev := getDeviceById(tar)
				if dev == nil {
					if errMsg != "" {
						errMsg = fmt.Sprintf("%s\n%s", errMsg, "设备不存在")
					} else {
						errMsg = "设备不存在"
					}
					continue
				}
				if body.DevAddress {
					address = append(address, dev.DeviceIp)
				}
				if body.SubNetBroadcast {
					address = append(address, "")
				}
				if body.LocalNetBroadcast {
					address = append(address, "255.255.255.255")
				}
			}
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
