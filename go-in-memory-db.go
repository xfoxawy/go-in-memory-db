package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/go-in-memory-db/actions"
	"github.com/go-in-memory-db/clients"
)

const port string = "8080"

func main() {

	portFlag := flag.String("port", port, "connection port")

	flag.Parse()

	port := fmt.Sprintf(":%s", *portFlag)

	fmt.Println("initiating DB connection on port: " + port)

	li, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	actionsChannle := make(chan *actions.Actions)

	for {
		conn, err := li.Accept()
		client := clients.ResolveClinet(conn)
		if err != nil {
			log.Fatalln(err)
		}
		go handle(client, actionsChannle)
		go actions.TakeAction(actionsChannle)
	}
}

func handle(c *clients.Client, ch chan *actions.Actions) {
	defer c.Conn.Close()

	log.SetOutput(os.Stdout)

	scanner := bufio.NewScanner(c.Conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		ch <- &actions.Actions{fs, c}
	}

	if err := scanner.Err(); err != nil {
		delete(c.GetConnections(), c.Address)
		fmt.Fprintln(os.Stderr, c.Address+" connection is closed")
	}
}
