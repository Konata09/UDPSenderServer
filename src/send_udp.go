package main

import (
	"fmt"
	"net"
)

type SendUdpPacket struct {
	TargetId      []Target `json:"targetid"`
	Port          int      `json:"port"`
	CommandId     int      `json:"commandid"`
	UseCustom     bool     `json:"usecustom"`
	CustomPayload string   `json:"custompayload"`
	Repeat        int      `json:"repeat"`
}

type Target struct {
	id int `json:"id"`
}

func sendUdp(ip string, port int, payload []byte) {
	pc, err := net.ListenPacket("udp4", "")
	if err != nil {
		panic(err)
	}
	defer pc.Close()
	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}

	_, err = pc.WriteTo(payload, addr)
	if err != nil {
		panic(err)
	}
}
