package databases

import (
	"math/rand"
	"time"

	"github.com/xfoxawy/go-in-memory-db/hashtable"
	"github.com/xfoxawy/go-in-memory-db/linkedlist"
	"github.com/xfoxawy/go-in-memory-db/queue"
	"github.com/xfoxawy/go-in-memory-db/stack"
)

var db *Database

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	db = &Database{
		randString(10),
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*stack.Stack),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
	}
	rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
