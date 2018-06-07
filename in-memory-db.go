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
	clear()
}

type db struct {
	namespace string
	public    bool
	data      map[string]string
}

func (db *db) set(k string, v string) bool {
	db.data[k] = v
	return true
}

func (db *db) get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

func (db *db) del(k string) bool {
	delete(db.data, k)
	return true
}

func (db *db) isset(k string) bool {
	if _, ok := db.data[k]; ok {
		return true
	}
	return false
}

func (db *db) dump() string {
	var content bytes.Buffer
	if len(db.data) > 0 {
		for k, v := range db.data {
			content.WriteString(fmt.Sprintf("%s %s\n", k, v))
		}
	}
	return content.String()
}

func (db *db) clear() {
	for k := range db.data {
		delete(db.data, k)
	}
}

func createNewDB(name string) db {
	db := db{
		name,
		true,
		make(map[string]string),
	}
	return db
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

	publicDb := createNewDB("test")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn, &publicDb)
	}
}

func handle(conn net.Conn, db *db) {
	defer conn.Close()

	log.SetOutput(os.Stdout)

	// create new scanner
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) >= 1 {

			switch strings.ToLower(fs[0]) {
			case "help":
				write(conn, help())
			case "set":
				if len(fs) != 3 {
					write(conn, "UNEXPECTED VALUE")
					continue
				}
				k := fs[1]
				v := fs[2]
				db.set(k, v)
				write(conn, "OK")

			case "get":
				k := fs[1]
				v, err := db.get(k)
				if err != nil {
					write(conn, "NIL")
					break
				}
				write(conn, v)

			case "del":
				k := fs[1]
				db.del(k)
				write(conn, "OK")

			case "isset":
				k := fs[1]
				if db.isset(k) {
					write(conn, "TRUE")
					break
				}
				write(conn, "FALSE")

			case "dump":
				content := db.dump()
				write(conn, content)

			case "clear":
				db.clear()
				write(conn, "OK")

			case "bye":
				conn.Close()

			default:
				write(conn, "UNKNOWN CMD")
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)

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
