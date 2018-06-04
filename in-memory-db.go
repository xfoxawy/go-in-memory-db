package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	portFlag := flag.String("port", "8080", "connection port")

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
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	log.SetOutput(os.Stdout)

	help := "\n In Memory DB \n\n" +
		"Use: \n" +
		"SET key value \n" +
		"GET key \n" +
		"DEL key \n" +
		"Example: \n" +
		"SET fav peanutbutter \n" +
		"GET fav \n" +
		"EXIST key \n" +
		"HELP ??\n"

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {

		case "HELP":
			write(conn, help)

		case "GET":
			k := fs[1]
			if v, ok := data[k]; ok {
				write(conn, v+"\n")
			} else {
				write(conn, "NIL\n")
			}

		case "SET":
			if len(fs) != 3 {
				write(conn, "UNEXPECTED VALUE\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
			write(conn, "OK\n")

		case "DEL":
			k := fs[1]
			delete(data, k)
			write(conn, "OK\n")

		case "ISSET":
			k := fs[1]
			if _, ok := data[k]; ok {
				write(conn, "TRUE\n")
			} else {
				write(conn, "FALSE\n")
			}
		case "DUMP":
			if len(data) > 0 {
				for k, v := range data {
					s := fmt.Sprintf("%s %s\n", k, v)
					write(conn, s)
				}
			} else {
				write(conn, "NIL\n")
			}
		case "CLEAR":
			for k := range data {
				delete(data, k)
			}
			write(conn, "OK\n")
		case "BYE":
			conn.Close()

		default:
			write(conn, "UNKNOWN CMD \n\n")
		}
	}
}

func write(w io.Writer, s string) {
	io.WriteString(w, s)
}
