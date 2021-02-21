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
		break
	case "PUT":
	case "POST":
	case "DELETE":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
