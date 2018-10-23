package databases

import (
	"errors"

	"github.com/go-in-memory-db/hashtable"
)

func (db *Database) GetHashTable(k string) (*hashtable.HashTable, error) {
	if _, ok := db.dataHashTable[k]; ok {
		return db.dataHashTable[k], nil
	}

	return nil, errors.New("not found")

}

func (db *Database) CreateHashTable(k string) *hashtable.HashTable {
	if _, ok := db.dataHashTable[k]; ok {
		errors.New("Hash Table Exists")
	}
	db.dataHashTable[k] = hashtable.NewHashTable()
	return db.dataHashTable[k]
}

func (db *Database) DelHashTable(k string) {
	delete(db.dataHashTable, k)
}
