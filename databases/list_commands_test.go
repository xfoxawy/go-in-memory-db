package databases

import (
	"testing"
)

func TestCreateList(t *testing.T) {
	key := randString(12)
	res, err := db.CreateList(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
	secondRes, secondErr := db.CreateList(key)
	if secondErr == nil {
		t.Error("expected", "error message", "got", secondRes)
	}
}

func TestDelList(t *testing.T) {
	key := randString(12)
	db.CreateList(key)
	db.DelList(key)
	res, err := db.CreateList(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
}

func TestGetList(t *testing.T) {
	key := randString(12)
	res, err := db.GetList(key)
	if err == nil {
		t.Error("expected", "error", "got", res)
	}
	newList, _ := db.CreateList(key)
	secondRes, secondErr := db.GetList(key)
	if secondErr != nil || secondRes != newList {
		t.Error("expected", newList, "got", secondErr)
	}
}
