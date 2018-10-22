package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
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

		case "qsize":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qsize := data.qSizeHandler()
			write(conn, qsize)

		case "qfront":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qfront := data.qFrontHandler()
			write(conn, qfront)

		case "qdeq":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qdeq := data.qDeqHandler()
			write(conn, qdeq)

		case "qenq":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			qenq := data.qEnqHandler()
			write(conn, qenq)

		case "hset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hset := data.hSetHandler()
			write(conn, hset)

		case "hget":
			if len(command) < 3 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hget := data.hGetHandler()
			write(conn, hget)

		case "hgetall":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hgetall := data.hGetAllHandler()
			write(conn, hgetall)

		case "hdel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hdel := data.hDelHandler()
			write(conn, hdel)

		case "hpush":
			if len(command) < 4 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hpush := data.hPushHandler()
			write(conn, hpush)

		case "hrm", "hremove":
			if len(command) < 3 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			hrm := data.hRemoveHandler()
			write(conn, hrm)

		case "lset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lset := data.lSetHandler()
			write(conn, lset)

		case "lget":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lget := data.lGetHandler()
			write(conn, lget)

		case "ldel":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			ldel := data.lDelHandler()
			write(conn, ldel)

		case "lpush":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			if len(command) < 3 {
				write(conn, "TYPE VALUES TO PUSH IT")
				continue
			}
			lpush := data.lPushHandler()
			write(conn, lpush)

		case "lpop":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lpop := data.lPopHandler()
			write(conn, lpop)

		case "lshift":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lshift := data.lShiftHandler()
			write(conn, lshift)

		case "lunshift":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lunshift := data.lUnShiftHandler()
			write(conn, lunshift)

		case "lrm", "lremove":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lremove := data.lRemoveHandler()
			write(conn, lremove)

		case "lunlink":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lunlink := data.lUnlinkHandler()
			write(conn, lunlink)

		case "lseek":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			lseek := data.lSeekHandler()
			write(conn, lseek)

		case "set":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			set := data.setHandler()
			write(conn, set)

		case "get":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			get := data.getHandler()
			write(conn, get)

		case "del":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			del := data.delHandler()
			write(conn, del)

		case "isset":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			isset := data.issetHandler()
			write(conn, isset)

		case "dump":
			content := data.dumpHandler()
			write(conn, content)

		case "clear":
			clear := data.clearHandler()
			write(conn, clear)

		case "which":
			witch := data.witchHandler()
			write(conn, witch)

		case "use":
			if len(command) < 2 {
				write(conn, "UNEXPECTED KEY")
				continue
			}
			use := data.useHandler()
			write(conn, use)

		case "show", "ls":
			show := data.showHandler()
			write(conn, show)

		case "bye":
			data.byeHandler()

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

func (a *Actions) qSizeHandler() string {
	k := a.stringArray[1]
	if queue, err := a.client.dbpointer.getQueue(k); err == nil {
		stringVal := strconv.Itoa(queue.Size())
		return stringVal
	}
	return "Queue Does not Exist"
}

func (a *Actions) qFrontHandler() string {
	k := a.stringArray[1]
	if queue, err := a.client.dbpointer.getQueue(k); err == nil {
		return queue.Front()
	}
	return "Queue Does not Exist"
}

func (a *Actions) qDeqHandler() string {
	k := a.stringArray[1]
	if queue, err := a.client.dbpointer.getQueue(k); err == nil {
		return queue.Dequeue()
	}
	return "Queue Does not Exist"
}

func (a *Actions) qEnqHandler() string {
	k := a.stringArray[1]
	v := a.stringArray[2:]

	if queue, err := a.client.dbpointer.getQueue(k); err == nil {
		for i := range v {
			queue.Enqueue(v[i])
		}
		return "Ok"
	}
	return "Queue Does not Exist"
}

func (a *Actions) hSetHandler() string {
	k := a.stringArray[1]
	a.client.dbpointer.createHashTable(k)
	return k
}

func (a *Actions) hGetHandler() string {
	k := a.stringArray[1]
	mapKey := a.stringArray[2]
	v, err := a.client.dbpointer.getHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}
	if value, ok := v.Values[mapKey]; ok {
		return value
	}
	return "This Key Does not Exist"
}

/**
* @to-do
* test this function
 */
func (a *Actions) hGetAllHandler() string {
	k := a.stringArray[1]
	v, err := a.client.dbpointer.getHashTable(k)

	if err != nil {
		return "Hash table Does not Exist"
	}

	for k, v := range v.Values {
		write(a.client.conn, k+" "+v)
	}
	return ""
}

func (a *Actions) hDelHandler() string {
	k := a.stringArray[1]
	if _, err := a.client.dbpointer.getHashTable(k); err == nil {
		a.client.dbpointer.delHashTable(k)
		return k + " " + "has deleted"
	}
	return "Hash table Does not Exist"
}

func (a *Actions) hPushHandler() string {
	k := a.stringArray[1]
	mapKey := a.stringArray[2]

	hash, err := a.client.dbpointer.getHashTable(k)
	if err != nil {
		return "Hash table Does not Exist"
	}

	value := a.stringArray[3]
	hash.Values = hash.Push(mapKey, value)
	return "OK"
}

func (a *Actions) hRemoveHandler() string {
	k := a.stringArray[1]
	mapKey := a.stringArray[2]

	if hash, err := a.client.dbpointer.getHashTable(k); err == nil {

		hash.Values = hash.Remove(mapKey)
		return "OK"
	}
	return "Hash table Does not Exist"
}

func (a *Actions) lSetHandler() string {
	k := a.stringArray[1]
	v := a.stringArray[2:]

	list := a.client.dbpointer.createList(k)

	for i := range v {
		list.Push(v[i])
	}
	return "OK"
}

// test it
func (a *Actions) lGetHandler() string {
	k := a.stringArray[1]
	v, err := a.client.dbpointer.getList(k)

	if err != nil || v.Start == nil {
		return "empty or not exit"
	}
	write(a.client.conn, v.Start.Value)
	current := v.Start
	for current.Next != nil {
		current = current.Next
		write(a.client.conn, current.Value)
	}
	return ""
}

func (a *Actions) lDelHandler() string {
	k := a.stringArray[1]
	if _, err := a.client.dbpointer.getList(k); err == nil {
		a.client.dbpointer.delList(k)
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lPushHandler() string {
	k := a.stringArray[1]
	list, err := a.client.dbpointer.getList(k)
	if err != nil {
		return "List Does not Exist"
	}
	values := a.stringArray[2:]
	for i := range values {
		list.Push(values[i])
	}
	return "OK"
}

func (a *Actions) lPopHandler() string {
	k := a.stringArray[1]
	if list, err := a.client.dbpointer.getList(k); err == nil {
		p, err := list.Pop()
		if err != nil {
			return "list is empty"
		}
		return p.Value
	}
	return "List Does not Exist"
}

//test it
func (a *Actions) lShiftHandler() string {
	k := a.stringArray[1]
	var v string

	if len(a.stringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.stringArray[2:], "")
	}

	if list, err := a.client.dbpointer.getList(k); err == nil {
		list.Shift(v)
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lUnShiftHandler() string {
	k := a.stringArray[1]

	if list, err := a.client.dbpointer.getList(k); err == nil {
		unshifted, err := list.Unshift()
		if err != nil {
			return "list is empty"
		}
		return unshifted.Value
	}
	return "List Does not Exist"
}

func (a *Actions) lRemoveHandler() string {
	k := a.stringArray[1]
	values := a.stringArray[2:]
	if list, err := a.client.dbpointer.getList(k); err == nil {
		for i := range values {
			err := list.Remove(values[i])
			if err != nil {
				return "list is empty"
			}
		}
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lUnlinkHandler() string {
	k := a.stringArray[1]
	values := a.stringArray[2:]
	if list, err := a.client.dbpointer.getList(k); err == nil {

		for i := range values {
			intVal, _ := strconv.Atoi(values[i])
			err := list.Unlink(intVal)
			if err != nil {
				return "LinkedList is empty OR Step Not Exist"
			}
		}
		return "OK"
	}
	return "List Does not Exist"
}

func (a *Actions) lSeekHandler() string {
	k := a.stringArray[1]
	var v string
	if len(a.stringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.stringArray[2:], "")
	}
	if list, err := a.client.dbpointer.getList(k); err == nil {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return "LinkedList is empty OR Step Not Exist"
		}
		value, err := list.Seek(intVal)
		if err != nil {
			return "LinkedList is empty OR Step Not Exist"
		}
		return value
	}
	return "List Does not Exist"
}

func (a *Actions) setHandler() string {
	k := a.stringArray[1]
	var v string
	if len(a.stringArray) == 2 {
		v = "NIL"
	} else {
		v = strings.Join(a.stringArray[2:], "")
	}

	a.client.dbpointer.set(k, v)
	return "OK"
}

func (a *Actions) getHandler() string {
	k := a.stringArray[1]
	v, err := a.client.dbpointer.get(k)
	if err != nil {
		return "NIL"
	}
	return v
}

func (a *Actions) delHandler() string {
	k := a.stringArray[1]
	a.client.dbpointer.del(k)
	return "OK"
}

func (a *Actions) issetHandler() string {
	k := a.stringArray[1]
	if a.client.dbpointer.isset(k) {
		return "TRUE"
	}
	return "FALSE"
}

func (a *Actions) dumpHandler() string {
	return a.client.dbpointer.dump()
}

func (a *Actions) clearHandler() string {
	a.client.dbpointer.clear()
	return "OK"
}

func (a *Actions) witchHandler() string {
	return a.client.dbpointer.name()
}

func (a *Actions) useHandler() string {
	key := a.stringArray[1]
	if db, ok := Databases[key]; ok {
		a.client.dbpointer = db
	} else {
		Databases[key] = &database{
			key,
			true,
			make(map[string]string),

			make(map[string]*linkedlist.LinkedList),
			make(map[string]*queue.Queue),
			make(map[string]*hashtable.HashTable),
		}
		a.client.dbpointer = Databases[key]
	}
	return "OK"
}

func (a *Actions) showHandler() string {
	var content bytes.Buffer
	for name := range Databases {
		if name == a.client.dbpointer.name() {
			name = fmt.Sprintf("-> %s\n", name)
		} else {
			name = fmt.Sprintf("%s\n", name)
		}
		content.WriteString(name)
	}
	return content.String()
}

func (a *Actions) byeHandler() {
	a.client.conn.Close()
	return
}
