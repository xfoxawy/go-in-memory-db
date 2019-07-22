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
			color.Blue("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			log.Println(conn.RemoteAddr())
			// this is called when the connection has been closed
			errlog := ""
			if err != nil {
				errlog = err.Error()
			}
			color.Red("closed: %s, err: %s", conn.RemoteAddr(), errlog)
		},
	)
	if err := <-err; err != nil {
		color.Red(err.Error())
	}
}
