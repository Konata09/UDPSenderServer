package main

import (
	"encoding/json"
	"net/http"
)

type Command struct {
	CommandId    int    `json:"id"`
	CommandName  string `json:"name"`
	CommandValue string `json:"value"`
	CommandPort  int    `json:"port"`
}
type AllCommands struct {
	Count    int       `json:"count"`
	Commands []Command `json:"commands"`
}

func SetCommand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		commands := getCommands()
		json.NewEncoder(w).Encode(&ApiReturn{
			Retcode: 0,
			Message: "OK",
			Data: &AllCommands{
				Count:    len(commands),
				Commands: commands,
			},
		})
	case "PUT":
		var body AllCommands
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		ok := addCommand(body.Commands)
		if ok {
			ApiOk(w)
		} else {
			ApiErr(w)
		}
	case "POST":
		var body Command
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		if getCommandById(body.CommandId) == nil {
			ApiErrMsg(w, "命令不存在")
			return
		}
		ok := setCommand(body.CommandId, body.CommandName, body.CommandValue, body.CommandPort)
		if ok {
			ApiOk(w)
		} else {
			ApiErr(w)
		}
	case "DELETE":
		var body Command
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			ApiErr(w)
			return
		}
		if getCommandById(body.CommandId) == nil {
			ApiErrMsg(w, "命令不存在")
			return
		}
		ok := deleteCommand(body.CommandId)
		if ok {
			ApiOk(w)
		} else {
			ApiErr(w)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
