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
		"QGET key \n" +
		"QDEL key \n" +
		"QSIZE key \n" +
		"QFRONT key \n" +
		"QDEQ key \n" +
		"QENQ key value1 value2 etc \n" +
		"\nEND OF QUEUE COMMANDS : \n\n" +
		"\nHASH TABLE COMMANDS : \n\n" +
		"HSET key \n" +
		"HGET key HashKey \n" +
		"HDEL key \n" +
		"HPUSH key HashKey value \n" +
		"HRM key HashKey c \n" +
		"HREMOVE HashKey c \n" +
		"\nEND OF HASH TABLE COMMANDS : \n\n" +
		"USE name\n" +
		"WHICH \n" +
		"SHOW \n" +
		"CLEAR \n" +
		"DUMP \n" +
		"BYE \n" +
		"HELP ??\n"
}
