package connection

import "github.com/tidwall/redcon"

var (
	connectionChannel = make(chan redcon.Conn)
)

func ShareConnection(conn redcon.Conn) {

	go func() {
		connectionChannel <- conn
	}()
}

func GetSharedConnection() redcon.Conn {
	return <-connectionChannel
}
