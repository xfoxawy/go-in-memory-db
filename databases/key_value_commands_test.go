package databases

import (
	"testing"

	"github.com/go-in-memory-db/hashtable"
	"github.com/go-in-memory-db/linkedlist"
	"github.com/go-in-memory-db/queue"
)

var db *Database

func init() {
	db = &Database{
		"master",
		true,
		make(map[string]string),
		make(map[string]*linkedlist.LinkedList),
		make(map[string]*queue.Queue),
		make(map[string]*hashtable.HashTable),
	}
}

func TestSet(t *testing.T) {
	key := "key"
	value := "value"
	db.Set(key, value)
	v, err := db.Get(key)
	if err != nil {
		t.Error("expected", value, "got error", err)
		return
	}
	if v != value {
		t.Error("expected", value, "got", v)
	}

}

func TestGet(t *testing.T) {
	_, err := db.Get("not exist value")
	if err == nil {
		t.Error("expected", "error", "got nothing")
		return
	}
}

func TestDel(t *testing.T) {
	key := "key"
	value := "value"
	db.Set(key, value)
	getValue, _ := db.Get(key)
	delValue := db.Del("key")
	if delValue != true {
		t.Error("expected", "true", "got", delValue)
	}
	getValueAfterDel, _ := db.Get(key)
	if getValue == getValueAfterDel {
		t.Error("expected", "", "got", getValueAfterDel)
	}
}
