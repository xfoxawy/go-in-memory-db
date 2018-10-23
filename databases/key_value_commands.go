package databases

import (
	"errors"
)

func (db *Database) Set(k string, v string) bool {
	db.data[k] = v
	return true
}

func (db *Database) Get(k string) (string, error) {
	if v, ok := db.data[k]; ok {
		return v, nil
	}
	return "", errors.New("not found")
}

func (db *Database) Del(k string) bool {
	delete(db.data, k)
	return true
}
