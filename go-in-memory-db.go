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
	getList(k string) (*LinkedList, error)
	createList(k string) *LinkedList
	delList(k string)
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

func (db *database) getList(k string) (*LinkedList, error) {
	if _, ok := db.dataList[k]; ok {
		return db.dataList[k], nil
	}

	return nil, errors.New("not found")

}

func (db *database) createList(k string) *LinkedList {
	if _, ok := db.dataList[k]; ok {
		errors.New("List Exists")
	}
	db.dataList[k] = NewList()
	return db.dataList[k]
}

func (db *database) delList(k string) {
	delete(db.dataList, k)
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

			case "lset":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v := fs[2:]

				list := c.dbpointer.createList(k)

				for i := range v {
					list.push(v[i])
				}

				write(c.conn, "OK")

			case "lget":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				v, err := c.dbpointer.getList(k)

				if err != nil {
					write(c.conn, "empty or not exit")
					continue
				}
				write(c.conn, v.start.value)
				current := v.start
				for current.next != nil {
					current = current.next
					write(c.conn, current.value)
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
					write(c.conn, k)
					continue
				}
				write(c.conn, "List Does not Exist")
				write(c.conn, k)

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
					list.push(values[i])
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
					p, err := list.pop()
					if err != nil {
						write(c.conn, "list is empty")
						continue
					}
					write(c.conn, p.value)
					continue
				}
				write(c.conn, "List Does not Exist")
				write(c.conn, k)

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
					list.shift(v)
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")
				write(c.conn, k)

			case "lunshift":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]

				if list, err := c.dbpointer.getList(k); err == nil {
					unshifted, err := list.unshift()
					if err != nil {
						write(c.conn, "list is empty")
						continue
					}
					write(c.conn, unshifted.value)
					continue
				}
				write(c.conn, "List Does not Exist")
				write(c.conn, k)

			// test this method by removing non existance key, or keep removing til its empty
			case "lrm", "lremove":
				if len(fs) < 2 {
					write(c.conn, "UNEXPECTED KEY")
					continue
				}
				k := fs[1]
				
				values := fs[2:]

				if list, err := c.dbpointer.getList(k); err == nil {

					for i := range values {
						err := list.remove(values[i])
						if err != nil {
							write(c.conn, "list is empty")
							break
						}
					}
					write(c.conn, "OK")
					continue
				}
				write(c.conn, "List Does not Exist")
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
