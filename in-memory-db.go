package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
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

	help := "\n In Memory DB \n\n" +
		"Use: \n" +
		"SET key value \n" +
		"GET key \n" +
		"DEL key \n" +
		"Example: \n" +
		"SET fav peanutbutter \n" +
		"GET fav \n" +
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
			v := data[k]
			write(conn, v+"\n")
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "UNEXPECTED VALUE")
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
		default:
			write(conn, "wait for more \n\n")
		}
	}
}

func write(w io.Writer, s string) {
	io.WriteString(w, s)
}
