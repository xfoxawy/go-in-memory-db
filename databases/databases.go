package databases

import (
	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
	"github.com/xfoxawy/go-in-memory-db/timeseries"
)

// DatabaseInterface Inteface
type DatabaseInterface interface {
	// DB Commands
	Dump() string
	Name() string
	Clear()
	// K,V commands
	Set(k string, v string) bool
	Get(k string) (string, error)
	Del(k string) bool
	Isset(k string) bool
	// List commands
	GetList(k string) (*linkedlist.LinkedList, error)
	CreateList(k string) (*linkedlist.LinkedList, error)
	DelList(k string)
	// Queue Commands
	GetQueue(k string) (*queue.Queue, error)
	CreateQueue(k string) (*queue.Queue, error)
	DelQueue(k string)
	// HashTable commands
	GetHashTable(k string) (*hashtable.HashTable, error)
	CreateHashTable(k string) (*hashtable.HashTable, error)
	DelHashTable(k string)
	// Timeseries Commands
	GetTimeSeries(k string) (*timeseries.TimeSeries, error)
	CreateTimeseries(k string) (*timeseries.TimeSeries, error)
	DelTimeseries(k string)
	ExpireTimeseries(k string, duration string) (*timeseries.TimeSeries, error)
}

// Database struct
type Database struct {
	Namespace      string
	Public         bool
	data           map[string]string
	dataList       map[string]*linkedlist.LinkedList
	queue          map[string]*queue.Queue
	dataHashTable  map[string]*hashtable.HashTable
	dataTimeseries map[string]*timeseries.TimeSeries
}

func NewPublicDatabase() *Database {
	return &Database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
		make(map[string]*timeseries.TimeSeries),
	}
}

// CreateMasterDB fucntion
func CreateMasterDB() *Database {
	return NewPublicDatabase()
}

// GetActiveDatabase function
func GetActiveDatabase(key string) *Database {
	return NewPublicDatabase()
}
