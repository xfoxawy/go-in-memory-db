package actions

import (
	"strings"

	"github.com/go-in-memory-db/clients"
	"github.com/redcon"
)

type Actions struct {
	StringArray []string
	Client      *clients.Client
}

func TakeAction(data *Actions) {
	command := data.StringArray
	conn := data.Client.Conn

	if len(command) < 1 {
		write(conn, "PLEASE TYPE SOMTHING")
		return
	}

	ad := NewDecisionManager(data)
	checkCommandExist := ad.CheckCommandAvailablity(strings.ToLower(command[0]))
	if !checkCommandExist {
		write(conn, "UNKNOW COMMAND, USE HELP")
		return
	}
	runner := ad.RunCommand(command)
	if runner == "" {
		write(conn, "OK")
	} else {
		write(conn, runner)
	}
}

func write(c redcon.Conn, s string) {
	if s != "" {
		c.WriteString(s)
	}
}
