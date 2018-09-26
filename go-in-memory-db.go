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

type Database interface {
	set(k string, v string) bool
	get(k string) (string, error)
	del(k string) bool
	isset(k string) bool
	dump() string
	name() string
	getDataList() map[string]*LinkedList
	clear()
}

type database struct {
	namespace string
	public    bool
	data      map[string]string
	dataList  map[string]*LinkedList
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

func indexOf(list []string, value string) int {
	for k, v := range list {
		if v == value {
			return k
		}
	}
	return -1
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

func (db *database) getDataList() map[string]*LinkedList {
	return db.dataList
}

func createMasterDB() *database {
	db := database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*LinkedList),
	}
	return &db
}

// MasterDb placeholder
var (
	MasterDb = createMasterDB()
)

// All Databases
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

	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) >= 1 {

			switch strings.ToLower(fs[0]) {

			case "help":
				write(c.conn, help())

			case "slist":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				c.dbpointer.getDataList()[k] = NewList()
				write(c.conn, "OK")

			case "glist":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v := c.dbpointer.getDataList()[k]
				
				if v==nil || v.length==0 {
					write(c.conn , "empty or not exit")
					continue
				}
				write(c.conn , v.start.value)
				current := v.start
				for current.next != nil {
					current = current.next
					write(c.conn , current.value)
				}

			case "dlist":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				if _, ok := c.dbpointer.getDataList()[k]; ok {
					delete(c.dbpointer.getDataList(), k)
					write(c.conn , "This List Has Deleted")
					write(c.conn , k)
					continue
				}
				write(c.conn, "This List Not Exist")
				write(c.conn, k)

			case "push":
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
				if list, ok := c.dbpointer.getDataList()[k]; ok {
					list.push(v)
					write(c.conn , "OK")
					continue
				}

				write(c.conn, "This List Not Exist")
				write(c.conn, k)

			case "pop":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				if list, ok := c.dbpointer.getDataList()[k]; ok {
					p,err := list.pop()
					if err != nil {
						write(c.conn , "list is empty")
					}
					write(c.conn, p.value)
					continue
				}
				write(c.conn, "This List Not Exist")
				write(c.conn, k)

			case "shift":
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

				if list, ok := c.dbpointer.getDataList()[k]; ok {
					list.shift(v)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "This List Not Exist")
				write(c.conn, k)

			case "unshift":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				if list, ok := c.dbpointer.getDataList()[k]; ok {
					unshifted,err := list.unshift()
					if err != nil {
						write(c.conn , "list is empty")
					}
					write(c.conn, unshifted.value)
					continue
				}
				write(c.conn, "This List Not Exist")
				write(c.conn, k)

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

						make(map[string]*LinkedList),
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
		"USE name\n" +
		"WHICH \n" +
		"SHOW \n" +
		"CLEAR \n" +
		"DUMP \n" +
		"BYE \n" +
		"HELP ??\n"
}
