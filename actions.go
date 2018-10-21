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

		case "qget":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qget := data.qGetHandler()
			write(conn, qget)

		case "qdel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qdel := data.qDelHandler()
			write(conn, qdel)

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

func (a *Actions) qGetHandler() string {

	k := a.stringArray[1]
	if q, err := a.client.dbpointer.getQueue(k); err == nil {
		write(a.client.conn, q.Queue.Start.Value)
		current := q.Queue.Start
		for current.Next != nil {
			current = current.Next
			write(a.client.conn, current.Value)
		}
		return ""
	}
	return "Queue Does not Exist"
}

func (a *Actions) qDelHandler() string {
	k := a.stringArray[1]
	if _, err := a.client.dbpointer.getQueue(k); err == nil {
		a.client.dbpointer.delQueue(k)
		return "OK"
	}
	return "Queue Does not Exist"
}
