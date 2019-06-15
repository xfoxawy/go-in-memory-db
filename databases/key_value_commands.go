package databases

import (
	"errors"
)

// Set Function
func (db *Database) Set(k string, v string) bool {
	db.data[k] = v
	return true
}

// Get Function
func (db *Database) Get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

// Del Function
func (db *Database) Del(k string) bool {
	delete(db.data, k)
	return true
}
