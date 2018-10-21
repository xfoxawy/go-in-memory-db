package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
)

const port string = "8080"

type Database interface {
	set(k string, v string) bool
	get(k string) (string, error)
	del(k string) bool
	isset(k string) bool
	dump() string
	name() string
	getList(k string) (*linkedlist.LinkedList, error)
	createList(k string) *linkedlist.LinkedList
	delList(k string)
	getQueue(k string) (*queue.Queue, error)
	createQueue(k string) *queue.Queue
	delQueue(k string)
	getHashTable(k string) (*hashtable.HashTable, error)
	createHashTable(k string) *hashtable.HashTable
	delHashTable(k string)
	clear()
}

type database struct {
	namespace     string
	public        bool
	data          map[string]string
	dataList      map[string]*linkedlist.LinkedList
	queue         map[string]*queue.Queue
	dataHashTable map[string]*hashtable.HashTable
}

type client struct {
	address   string
	conn      net.Conn
	dbpointer Database
}

func (db *database) set(k string, v string) bool {
	db.data[k] = v
	return true
}

func (db *database) get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

func (db *database) del(k string) bool {
	delete(db.data, k)
	return true
}

func (db *database) isset(k string) bool {
	if _, ok := db.data[k]; ok {
		return true
	}
	if _, ok := db.dataList[k]; ok {
		return true
	}
	return false
}

func (db *database) dump() string {
	var content bytes.Buffer
	if len(db.data) > 0 {
		for k, v := range db.data {
			content.WriteString(fmt.Sprintf("%s %s\n", k, v))
		}
	}
	return content.String()
}

func (db *database) clear() {
	go func() {
		for k := range db.data {
			delete(db.data, k)
		}
	}()
}

func (db *database) name() string {
	return db.namespace
}

func (db *database) getList(k string) (*linkedlist.LinkedList, error) {
	if _, ok := db.dataList[k]; ok {
		return db.dataList[k], nil
	}

	return nil, errors.New("not found")

}

func (db *database) createList(k string) *linkedlist.LinkedList {
	if _, ok := db.dataList[k]; ok {
		errors.New("List Exists")
	}
	db.dataList[k] = linkedlist.NewList()
	return db.dataList[k]
}

func (db *database) delList(k string) {
	delete(db.dataList, k)
}

func (db *database) createQueue(k string) *queue.Queue {
	if queue, ok := db.queue[k]; ok {
		return queue
	}
	db.queue[k] = queue.NewQueue()
	return db.queue[k]
}

func (db *database) getQueue(k string) (*queue.Queue, error) {
	if _, ok := db.queue[k]; ok {
		return db.queue[k], nil
	}
	return nil, errors.New("not found")
}

func (db *database) delQueue(k string) {
	delete(db.queue, k)
}

func (db *database) getHashTable(k string) (*hashtable.HashTable, error) {
	if _, ok := db.dataHashTable[k]; ok {
		return db.dataHashTable[k], nil
	}

	return nil, errors.New("not found")

}

func (db *database) createHashTable(k string) *hashtable.HashTable {
	if _, ok := db.dataHashTable[k]; ok {
		errors.New("Hash Table Exists")
	}
	db.dataHashTable[k] = hashtable.NewHashTable()
	return db.dataHashTable[k]
}

func (db *database) delHashTable(k string) {
	delete(db.dataHashTable, k)
}

func createMasterDB() *database {
	db := database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
	}
	return &db
}

// MasterDb placeholder
// All Databases
var (
	MasterDb = createMasterDB()
)

var (
	Databases = map[string]*database{"master": createMasterDB()}
)

var (
	Connections = make(map[string]*client)
)

func resolveClinet(conn net.Conn) *client {
	addr := conn.RemoteAddr().String()
	if _, ok := Connections[addr]; ok == false {
		Connections[addr] = &client{
			addr,
			conn,
			MasterDb,
		}
	}
	return Connections[addr]

}

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

	actionsChannle := make(chan *Actions)

	for {
		conn, err := li.Accept()
		client := resolveClinet(conn)
		if err != nil {
			log.Fatalln(err)
		}
		go handle(client, actionsChannle)
		go takeAction(actionsChannle)
	}
}

func handle(c *client, ch chan *Actions) {
	defer c.conn.Close()

	log.SetOutput(os.Stdout)

	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		ch <- &Actions{fs, c}
		if len(fs) >= 1 {

			switch strings.ToLower(fs[0]) {

			case "help":
				write(c.conn, help())

			case "qset":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v := fs[2:]

				queue := c.dbpointer.createQueue(k)

				for i := range v {
					queue.Enqueue(v[i])
				}

				write(c.conn, "OK")

			case "qget":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if q, err := c.dbpointer.getQueue(k); err == nil {
					write(c.conn, q.Queue.Start.Value)
					current := q.Queue.Start
					for current.Next != nil {
						current = current.Next
						write(c.conn, current.Value)
					}
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "qdel":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if _, err := c.dbpointer.getQueue(k); err == nil {
					c.dbpointer.delQueue(k)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "qsize":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if queue, err := c.dbpointer.getQueue(k); err == nil {
					stringVal := strconv.Itoa(queue.Size())
					write(c.conn, stringVal)
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "qfront":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if queue, err := c.dbpointer.getQueue(k); err == nil {
					write(c.conn, queue.Front())
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "qdeq":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if queue, err := c.dbpointer.getQueue(k); err == nil {
					write(c.conn, queue.Dequeue())
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "qenq":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v := fs[2:]

				if queue, err := c.dbpointer.getQueue(k); err == nil {
					for i := range v {
						queue.Enqueue(v[i])
					}
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "Queue Does not Exist")

			case "hset":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				c.dbpointer.createHashTable(k)
				write(c.conn, k)

			case "hget":
				if len(fs) < 3 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				mapKey := fs[2]
				v, err := c.dbpointer.getHashTable(k)

				if err != nil {
					write(c.conn, "Hash table Does not Exist")
					continue
				}
				if value, ok := v.Values[mapKey]; ok {
					write(c.conn, value)
					continue
				}
				write(c.conn, "This Key Does not Exist")
				continue

			case "hgetall":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v, err := c.dbpointer.getHashTable(k)

				if err != nil {
					write(c.conn, "Hash table Does not Exist")
					continue
				}

				for k, v := range v.Values {
					write(c.conn, k+" "+v)
				}

			case "hdel":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if _, err := c.dbpointer.getHashTable(k); err == nil {
					c.dbpointer.delHashTable(k)
					write(c.conn, "OK")
					write(c.conn, k)
					continue
				}
				write(c.conn, "Hash table Does not Exist")
				write(c.conn, k)

			case "hpush":
				if len(fs) < 4 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}

				k := fs[1]
				mapKey := fs[2]

				hash, err := c.dbpointer.getHashTable(k)
				if err != nil {
					write(c.conn, "Hash table Does not Exist")
					continue
				}

				value := fs[3]
				hash.Values = hash.Push(mapKey, value)
				write(c.conn, "OK")
				continue

			case "hrm", "hremove":
				if len(fs) < 3 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				mapKey := fs[2]

				if hash, err := c.dbpointer.getHashTable(k); err == nil {

					hash.Values = hash.Remove(mapKey)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "Hash table Does not Exist")
				write(c.conn, k)

			case "lset":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v := fs[2:]

				list := c.dbpointer.createList(k)

				for i := range v {
					list.Push(v[i])
				}

				write(c.conn, "OK")

			case "lget":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v, err := c.dbpointer.getList(k)

				if err != nil || v.Start == nil {
					write(c.conn, "empty or not exit")
					continue
				}
				write(c.conn, v.Start.Value)
				current := v.Start
				for current.Next != nil {
					current = current.Next
					write(c.conn, current.Value)
				}

			case "ldel":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if _, err := c.dbpointer.getList(k); err == nil {
					c.dbpointer.delList(k)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")

			case "lpush":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}

				if len(fs) < 3 {
					continue
				}

				k := fs[1]

				list, err := c.dbpointer.getList(k)

				if err != nil {
					write(c.conn, "List Does not Exist")
					continue
				}

				values := fs[2:]

				for i := range values {
					list.Push(values[i])
				}

				write(c.conn, "OK")
				continue

			case "lpop":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				if list, err := c.dbpointer.getList(k); err == nil {
					p, err := list.Pop()
					if err != nil {
						write(c.conn, "list is empty")
						continue
					}
					write(c.conn, p.Value)
					continue
				}
				write(c.conn, "List Does not Exist")

			case "lshift":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				var v string

				if len(fs) == 2 {
					v = "NIL"
				} else {
					v = strings.Join(fs[2:], "")
				}

				if list, err := c.dbpointer.getList(k); err == nil {
					list.Shift(v)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")

			case "lunshift":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				if list, err := c.dbpointer.getList(k); err == nil {
					unshifted, err := list.Unshift()
					if err != nil {
						write(c.conn, "list is empty")
						continue
					}
					write(c.conn, unshifted.Value)
					continue
				}
				write(c.conn, "List Does not Exist")

			// test this method by removing non existance key, or keep removing til its empty
			// in case empty it will show list is empty , OK
			// in non existance key it will return OK only
			case "lrm", "lremove":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				values := fs[2:]

				if list, err := c.dbpointer.getList(k); err == nil {

					for i := range values {
						err := list.Remove(values[i])
						if err != nil {
							write(c.conn, "list is empty")
							break
						}
					}
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")

			case "lunlink":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				values := fs[2:]

				if list, err := c.dbpointer.getList(k); err == nil {

					for i := range values {
						intVal, _ := strconv.Atoi(values[i])
						err := list.Unlink(intVal)
						if err != nil {
							write(c.conn, "LinkedList is empty OR Step Not Exist")
							break
						}
					}
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")

			case "lseek":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				var v string

				if len(fs) == 2 {
					v = "NIL"
				} else {
					v = strings.Join(fs[2:], "")
				}
				if list, err := c.dbpointer.getList(k); err == nil {
					intVal, err := strconv.Atoi(v)
					if err != nil {
						write(c.conn, "LinkedList is empty OR Step Not Exist")
						continue
					}
					value, err := list.Seek(intVal)
					if err != nil {
						write(c.conn, "LinkedList is empty OR Step Not Exist")
						continue
					}
					write(c.conn, value)
					continue
				}
				write(c.conn, "List Does not Exist")

			case "set":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				var v string

				if len(fs) == 2 {
					v = "NIL"
				} else {
					v = strings.Join(fs[2:], "")
				}

				c.dbpointer.set(k, v)
				write(c.conn, "OK")

			case "get":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v, err := c.dbpointer.get(k)
				if err != nil {
					write(c.conn, "NIL")
					break
				}
				write(c.conn, v)

			case "del":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				c.dbpointer.del(k)
				write(c.conn, "OK")

			case "isset":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if c.dbpointer.isset(k) {
					write(c.conn, "TRUE")
					break
				}
				write(c.conn, "FALSE")

			case "dump":
				content := c.dbpointer.dump()
				write(c.conn, content)

			case "clear":
				c.dbpointer.clear()
				write(c.conn, "OK")

			case "which":
				write(c.conn, c.dbpointer.name())

			case "use":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				key := fs[1]
				if db, ok := Databases[key]; ok {
					c.dbpointer = db
				} else {
					Databases[key] = &database{
						key,
						true,
						make(map[string]string),

						make(map[string]*linkedlist.LinkedList),
						make(map[string]*queue.Queue),
						make(map[string]*hashtable.HashTable),
					}
					c.dbpointer = Databases[key]
				}

			case "show", "ls":
				var content bytes.Buffer
				for name := range Databases {
					if name == c.dbpointer.name() {
						name = fmt.Sprintf("-> %s\n", name)
					} else {
						name = fmt.Sprintf("%s\n", name)
					}
					content.WriteString(name)
				}
				write(c.conn, content.String())

			case "bye":
				c.conn.Close()

			default:
				write(c.conn, "UNKNOWN CMD")
			}
		}

	}

	if err := scanner.Err(); err != nil {
		delete(Connections, c.address)
		fmt.Fprintln(os.Stderr, c.address+" connection is closed")
	}
}

func write(w io.Writer, s string) {
	io.WriteString(w, s+"\n")
}

func help() string {
	return "\n In Memory DB \n\n" +
		"Use: \n" +
		"SET key value \n" +
		"GET key \n" +
		"DEL key \n" +
		"ISSET key \n" +
		"\nLINKED LIST COMMANDS : \n\n" +
		"LSET key \n" +
		"LGET key \n" +
		"LDEL key \n" +
		"LPUSH key value1 value2 etc \n" +
		"LPOP key \n" +
		"LSHIFT key value \n" +
		"LUNSHIFT key \n" +
		"LRM key value1 value2 etc \n" +
		"LREMOVE key value1 value2 etc \n" +
		"LUNLINK key index1 index2 etc \n" +
		"LSEEK key index \n" +
		"\nEND OF LINKED LIST COMMANDS : \n\n" +
		"\nQUEUE COMMANDS : \n\n" +
		"QSET key \n" +
		"QDEL key \n" +
		"QSIZE key \n" +
		"QFRONT key \n" +
		"QDEQ key \n" +
		"QENQ key value1 value2 etc \n" +
		"\nEND OF QUEUE COMMANDS : \n\n" +
		"\nEND OF LINKED LIST COMMANDS : \n\n" +
		"\nHASH TABLE COMMANDS : \n\n" +
		"HSET key \n" +
		"HGET key \n" +
		"HDEL key \n" +
		"HPUSH key value1 value2 etc \n" +
		"HRM key value1 value2 etc \n" +
		"HREMOVE key value1 value2 etc \n" +
		"HUNLINK key index \n" +
		"HSEEK key index1 index2 etc... \n" +
		"\nEND OF HASH TABLE COMMANDS : \n\n" +
		"USE name\n" +
		"WHICH \n" +
		"SHOW \n" +
		"CLEAR \n" +
		"DUMP \n" +
		"BYE \n" +
		"HELP ??\n"
}
