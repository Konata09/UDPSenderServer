package main

type SendWolPacket struct {
	TargetId          []Target `json:"targetid"`
	DevAddress        bool     `json:"devaddress"`
	LocalNetBroadcast bool     `json:"localnetbroadcast"`
	SubNetBroadcast   bool     `json:"subnetbroadcast"`
	Repeat            int      `json:"repeat"`
}
