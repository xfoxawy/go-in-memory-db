package main

import (
	"crypto/tls"
	"log"

	"github.com/go-in-memory-db/actions"
	"github.com/go-in-memory-db/clients"
	"github.com/redcon"
)

var addr = ":6380"

func main() {

	go log.Printf("started server at %s", addr)
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	err = redcon.ListenAndServeTLS(addr,
		func(conn redcon.Conn, cmd redcon.Command) {

			var stringCommands []string

			for _, v := range cmd.Args {
				stringCommands = append(stringCommands, string(v))
			}

			client := clients.ResolveClinet(conn)

			action := handle(client, stringCommands)
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
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(c *clients.Client, fs []string) *actions.Actions {
	return &actions.Actions{fs, c}
}
