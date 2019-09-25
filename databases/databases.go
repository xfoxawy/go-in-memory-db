package databases

import (
	"github.com/go-in-memory-db/binarytrees"
	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
	"github.com/go-in-memory-db/stack"
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
	GetBinaryTree(k string) (*binarytrees.BinaryTree, error)
	CreateBinaryTree(k string) (*binarytrees.BinaryTree, error)
	DelBinaryTree(k string)
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
	binarytree    map[string]*binarytrees.BinaryTree
}

// CreateMasterDB fucntion
func CreateMasterDB() *Database {
	db := Database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*stack.Stack),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
		make(map[string]*binarytrees.BinaryTree),
	}
	return &db
}

// GetActiveDatabase function
func GetActiveDatabase(key string) *Database {
	db := Database{
		key,
		true,
		make(map[string]string),

		make(map[string]*linkedlist.LinkedList),
		make(map[string]*stack.Stack),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
		make(map[string]*binarytrees.BinaryTree),
	}
	return &db
}
