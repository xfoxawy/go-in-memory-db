package databases

import (
	"bytes"
	"fmt"
)

// Isset function
// to - do check isset in all data types
func (db *Database) Isset(k string) bool {
	if _, ok := db.data[k]; ok {
		return true
	}
	if _, ok := db.dataList[k]; ok {
		return true
	}
	return false
}

// Dump wihout unit test fuction
func (db *Database) Dump() string {
	var content bytes.Buffer
	if len(db.data) > 0 {
		for k, v := range db.data {
			content.WriteString(fmt.Sprintf("%s %s\n", k, v))
		}
	}
	return content.String()
}

// Clear remove all from memory
// to - do remove from list and other types
func (db *Database) Clear() {
	for k := range db.data {
		delete(db.data, k)
	}
}

// Name without unit test
func (db *Database) Name() string {
	return db.Namespace
}
