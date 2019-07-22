package databases

import (
	"github.com/xfoxawy/go-in-memory-db/hashtable"
	"github.com/xfoxawy/go-in-memory-db/linkedlist"
	"github.com/xfoxawy/go-in-memory-db/queue"
	"github.com/xfoxawy/go-in-memory-db/stack"
	"github.com/xfoxawy/go-in-memory-db/timeseries"
)

// DatabaseInterface Inteface
type DatabaseInterface interface {
	Set(k string, v string) bool
	Get(k string) (string, error)
	Del(k string) bool
	Isset(k string) bool
	Dump() string
	Name() string
	GetList(k string) (*linkedlist.LinkedList, error)
	CreateList(k string) (*linkedlist.LinkedList, error)
	DelList(k string)
	GetQueue(k string) (*queue.Queue, error)
	CreateQueue(k string) (*queue.Queue, error)
	DelQueue(k string)
	GetStack(k string) (*stack.Stack, error)
	CreateStack(k string) (*stack.Stack, error)
	DelStack(k string)
	GetHashTable(k string) (*hashtable.HashTable, error)
	CreateHashTable(k string) (*hashtable.HashTable, error)
	DelHashTable(k string)
	CreateTimeseries(k string) (*timeseries.Timeseries, error)
	GetTimeseries(k string) (*timeseries.Timeseries, error)
	DelTimeseries(k string)
	Clear()
}

// Database struct
type Database struct {
	Namespace     string
	Public        bool
	data          map[string]string
	dataList      map[string]*linkedlist.LinkedList
	stack         map[string]*stack.Stack
	queue         map[string]*queue.Queue
	dataHashTable map[string]*hashtable.HashTable
	timeseries    map[string]*timeseries.Timeseries
}

// CreateMasterDB fucntion
func CreateMasterDB() *Database {
	return CreateNewDatabase("master")
}

// CreateNewDatabase function
func CreateNewDatabase(key string) *Database {
	db := Database{
		key,
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*stack.Stack),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
		make(map[string]*timeseries.Timeseries),
	}
	return &db
}
