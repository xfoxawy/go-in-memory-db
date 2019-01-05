package main

import (
	"log"

	"github.com/go-in-memory-db/actions"
	"github.com/go-in-memory-db/clients"
	"github.com/redcon"
)

var addr = ":6380"

func main() {

	actionsChannle := make(chan *actions.Actions)

	go log.Printf("started server at %s", addr)
	err := redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {

			var stringCommands []string

			for _, v := range cmd.Args {
				stringCommands = append(stringCommands, string(v))
			}

			client := clients.ResolveClinet(conn)

			go handle(client, actionsChannle, stringCommands)
			go actions.TakeAction(actionsChannle)
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
	if err != nil {
		log.Fatal(err)
	}
}

func handle(c *clients.Client, ch chan *actions.Actions, fs []string) {
	ch <- &actions.Actions{fs, c}
}
