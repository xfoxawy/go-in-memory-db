package actions

import (
	"strings"

	"github.com/xfoxawy/go-in-memory-db/clients"
)

type Actions struct {
	StringArray []string
	Client      *clients.Client
}

func TakeAction(data *Actions) {
	command := data.StringArray
	conn := data.Client.Conn

	if len(command) < 1 {
		conn.WriteString("PLEASE ENTER A COMMAND")
		return
	}

	ad := NewDecisionManager(data)
	checkCommandExist := ad.CheckCommandAvailablity(strings.ToLower(command[0]))
	if !checkCommandExist {
		conn.WriteString("UNKNOWN COMMAND, USE HELP")
		return
	}
	runner := ad.RunCommand(command)
	if runner == "" {
		conn.WriteString("OK")
		return
	}
	conn.WriteString(runner)
}
