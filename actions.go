package main

import (
	"strings"
)

type Actions struct {
	stringArray []string
	client      *client
}

func takeAction(ch chan *Actions) {
	for {
		data := <-ch
		command := data.stringArray
		conn := data.client.conn
		if len(command) < 1 {
			write(conn, "please type somthing :D")
			continue
		}
		switch strings.ToLower(command[0]) {
		case "help":
			write(conn, help())
		case "qset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			data.qSetHanlder()
			write(conn, "OK")
		default:
			write(conn, "please use help command to know the commands you can use")
		}
	}
}

func (a *Actions) qSetHanlder() {
	k := a.stringArray[1]
	v := a.stringArray[2:]

	queue := a.client.dbpointer.createQueue(k)

	for i := range v {
		queue.Enqueue(v[i])
	}

}
