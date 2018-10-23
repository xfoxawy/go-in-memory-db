package databases

import (
	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
)

type DatabaseInterface interface {
	Set(k string, v string) bool
	Get(k string) (string, error)
	Del(k string) bool
	Isset(k string) bool
	Dump() string
	Name() string
	GetList(k string) (*linkedlist.LinkedList, error)
	CreateList(k string) *linkedlist.LinkedList
	DelList(k string)
	GetQueue(k string) (*queue.Queue, error)
	CreateQueue(k string) *queue.Queue
	DelQueue(k string)
	GetHashTable(k string) (*hashtable.HashTable, error)
	CreateHashTable(k string) *hashtable.HashTable
	DelHashTable(k string)
	Clear()
}

type Database struct {
	Namespace     string
	Public        bool
	data          map[string]string
	dataList      map[string]*linkedlist.LinkedList
	queue         map[string]*queue.Queue
	dataHashTable map[string]*hashtable.HashTable
}

func CreateMasterDB() *Database {
	db := Database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
	}
	return &db
}

func GetActiveDatabase(key string) *Database {
	db := Database{
		key,
		true,
		make(map[string]string),

		make(map[string]*linkedlist.LinkedList),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
	}
	return &db
}
