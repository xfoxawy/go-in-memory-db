package clients

import (
	"github.com/tidwall/redcon"
	"github.com/xfoxawy/go-in-memory-db/databases"
)

type Client struct {
	Address   string
	Conn      redcon.Conn
	Dbpointer databases.DatabaseInterface
}

var (
	Databases = map[string]*databases.Database{"master": databases.CreateMasterDB()}
)

var (
	Connections = make(map[string]*Client)
)

func ResolveClinet(conn redcon.Conn) *Client {
	addr := conn.RemoteAddr()
	if _, ok := Connections[addr]; ok == false {
		Connections[addr] = &Client{
			addr,
			conn,
			Databases["master"],
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
		Databases[key] = databases.CreateNewDatabase(key)
		c.Dbpointer = Databases[key]
	}
}

func (c *Client) GetAllDatabases() map[string]*databases.Database {
	return Databases
}
