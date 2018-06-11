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
)

const port string = "8080"

type database interface {
	set(k string, v string) bool
	get(k string) (string, error)
	del(k string) bool
	isset(k string) bool
	dump() string
	name() string
	clear()
	switchDb(name string) database
}

type masterInstance struct {
	namespace string
	public    bool
	data      map[string]string
	sub       map[string]*subInstance
}

type subInstance struct {
	namespace string
	public    bool
	data      map[string]string
	parent    *masterInstance
}

type client struct {
	address   string
	conn      net.Conn
	dbpointer database
}

/**
MasterDb methods
**/
func (db *masterInstance) set(k string, v string) bool {
	db.data[k] = v
	return true
}

func (db *masterInstance) get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

func (db *masterInstance) del(k string) bool {
	delete(db.data, k)
	return true
}

func (db *masterInstance) isset(k string) bool {
	if _, ok := db.data[k]; ok {
		return true
	}
	return false
}

func (db *masterInstance) dump() string {
	var content bytes.Buffer
	if len(db.data) > 0 {
		for k, v := range db.data {
			content.WriteString(fmt.Sprintf("%s %s\n", k, v))
		}
	}
	return content.String()
}

func (db *masterInstance) clear() {
	for k := range db.data {
		delete(db.data, k)
	}
}

func (db *masterInstance) name() string {
	return db.namespace
}

func (db *masterInstance) switchDb(name string) database {
	if name == "main" || name == "master" {
		return db
	}

	if _, ok := db.sub[name]; ok == false {
		db.sub[name] = createSubDB(db, name, true)
	}
	return db.sub[name]
}

/**
SubDb Methods
**/
func (db *subInstance) set(k string, v string) bool {
	db.data[k] = v
	return true
}

func (db *subInstance) get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

func (db *subInstance) del(k string) bool {
	delete(db.data, k)
	return true
}

func (db *subInstance) isset(k string) bool {
	if _, ok := db.data[k]; ok {
		return true
	}
	return false
}

func (db *subInstance) dump() string {
	var content bytes.Buffer
	if len(db.data) > 0 {
		for k, v := range db.data {
			content.WriteString(fmt.Sprintf("%s %s\n", k, v))
		}
	}
	return content.String()
}

func (db *subInstance) clear() {
	for k := range db.data {
		delete(db.data, k)
	}
}

func (db *subInstance) name() string {
	return db.namespace
}

func (db *subInstance) switchDb(name string) database {
	masterDb := db.parent

	if name == "main" || name == "master" {
		return masterDb
	}

	if _, ok := masterDb.sub[name]; ok == false {
		masterDb.sub[name] = createSubDB(masterDb, name, true)
	}
	return masterDb.sub[name]

}

func createMasterDB() *masterInstance {
	db := masterInstance{
		"master",
		true,
		make(map[string]string),
		make(map[string]*subInstance),
	}
	return &db
}

func createSubDB(master *masterInstance, name string, visibility bool) *subInstance {
	db := subInstance{
		name,
		visibility,
		make(map[string]string),
		master,
	}
	return &db
}

var (
	MasterDb = createMasterDB()
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

	for {
		conn, err := li.Accept()
		client := resolveClinet(conn)
		if err != nil {
			log.Fatalln(err)
		}
		go handle(client)
	}
}

func handle(c *client) {
	defer c.conn.Close()

	log.SetOutput(os.Stdout)

	// create new scanner
	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) >= 1 {

			switch strings.ToLower(fs[0]) {

			case "help":
				write(c.conn, help())

			case "set":
				if len(fs) != 3 {
					write(c.conn, "UNEXPECTED VALUE")
					continue
				}
				k := fs[1]
				v := fs[2]
				c.dbpointer.set(k, v)
				write(c.conn, "OK")

			case "get":
				k := fs[1]
				v, err := c.dbpointer.get(k)
				if err != nil {
					write(c.conn, "NIL")
					break
				}
				write(c.conn, v)

			case "del":
				k := fs[1]
				c.dbpointer.del(k)
				write(c.conn, "OK")

			case "isset":
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
				dbName := fs[1]
				c.dbpointer = c.dbpointer.switchDb(dbName)

			case "show":
				var content bytes.Buffer
				for name, _ := range MasterDb.sub {
					if name == c.dbpointer.name() {
						name = fmt.Sprintf("*%s\n", name)
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
		"CLEAR \n" +
		"DUMP \n" +
		"BYE \n" +
		"HELP ??\n"
}
