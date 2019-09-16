package main

import (
	"crypto/tls"
	"flag"
	"github.com/go-in-memory-db/actions"
	"github.com/go-in-memory-db/clients"
	"github.com/redcon"
	"log"
)

var addr = ":6380"

func main() {

	go log.Printf("started server at %s", addr)
	ssl := flag.Bool("ssl", false, "")
	privateKey := flag.String("private-key", "", "")
	publicKey := flag.String("public-key", "", "")
	flag.Parse()
	var err error
	if *ssl == true && *privateKey != "" && *publicKey != "" {
		cer, err := tls.LoadX509KeyPair(*publicKey, *privateKey)
		if err != nil {
			log.Fatal(err)
			return
		}
		config := &tls.Config{Certificates: []tls.Certificate{cer}}
		err = redcon.ListenAndServeTLS(addr, connectionHandler, acceptConnection, closeConnection, config)
	} else {
		err = redcon.ListenAndServe(addr, connectionHandler, acceptConnection, closeConnection)
	}
	if err != nil {
		log.Fatal(err)
	}
}

var connectionHandler = func(conn redcon.Conn, cmd redcon.Command) {
	var stringCommands []string

	for _, v := range cmd.Args {
		stringCommands = append(stringCommands, string(v))
	}

	client := clients.ResolveClinet(conn)

	action := handle(client, stringCommands)
	actions.TakeAction(action)
}

var acceptConnection = func(conn redcon.Conn) bool {
	// use this function to accept or deny the connection.
	// log.Printf("accept: %s", conn.RemoteAddr())
	return true
}

var closeConnection = func(conn redcon.Conn, err error) {
	// this is called when the connection has been closed
	// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
}

func handle(c *clients.Client, fs []string) *actions.Actions {
	return &actions.Actions{fs, c}
}
