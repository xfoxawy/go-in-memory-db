package main

import (
	"github.com/alash3al/go-color"

	"log"

	"github.com/tidwall/redcon"
	"github.com/xfoxawy/go-in-memory-db/actions"
	"github.com/xfoxawy/go-in-memory-db/clients"
)

var addr = ":6380"

func main() {

	go log.Printf("started server at %s", addr)
	err := make(chan error)

	err <- redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {

			var stringCommands []string

			for _, v := range cmd.Args {
				stringCommands = append(stringCommands, string(v))
			}

			client := clients.ResolveClinet(conn)

			action := &actions.Actions{stringCommands, client}
			actions.TakeAction(action)
		},
		func(conn redcon.Conn) bool {
			// use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// this is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err := <-err; err != nil {
		color.Red(err.Error())
	}
}
