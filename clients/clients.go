package clients

import (
	"net"

	"github.com/go-in-memory-db/databases"
)

type Client struct {
	Address   string
	Conn      net.Conn
	Dbpointer databases.DatabaseInterface
}

// MasterDb placeholder
// All Databases
var (
	MasterDb = databases.CreateMasterDB()
)

var (
	Databases = map[string]*databases.Database{"master": databases.CreateMasterDB()}
)

var (
	Connections = make(map[string]*Client)
)

func ResolveClinet(conn net.Conn) *Client {
	addr := conn.RemoteAddr().String()
	if _, ok := Connections[addr]; ok == false {
		Connections[addr] = &Client{
			addr,
			conn,
			MasterDb,
		}
	}
	return Connections[addr]
}

func (c *Client) GetConnections() map[string]*Client {
	return Connections
}

func (c *Client) UseNewDatabase(key string) {
	if db, ok := Databases[key]; ok {
		c.Dbpointer = db
	} else {
		Databases[key] = databases.GetActiveDatabase(key)
		c.Dbpointer = Databases[key]
	}
}

func (c *Client) GetAllDatabases() map[string]*databases.Database {
	return Databases
}
