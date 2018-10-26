package databases

import (
	"testing"
)

func TestIsset(t *testing.T) {
	key := randString(5)
	res := db.Isset(key)
	if res != false {
		t.Error("exprected", "fasle", "got", res)
	}
	db.Set(key, "value")
	secondRes := db.Isset(key)
	if secondRes != true {
		t.Error("expected", "true", "got", secondRes)
	}
}

func TestClear(t *testing.T) {
	key := randString(5)
	db.Set(key, "value")
	beforeClear, _ := db.Get(key)
	db.Clear()
	afterClear, _ := db.Get(key)
	if beforeClear == afterClear {
		t.Error(afterClear)
	}
}
