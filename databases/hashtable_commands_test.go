package databases

import "testing"

func TestCreateHashTable(t *testing.T) {
	key := randString(12)
	res, err := db.CreateHashTable(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
	secondRes, secondErr := db.CreateHashTable(key)
	if secondErr == nil {
		t.Error("expected", "error message", "got", secondRes)
	}
}

func TestDelHashTable(t *testing.T) {
	key := randString(12)
	db.CreateHashTable(key)
	db.DelHashTable(key)
	res, err := db.CreateHashTable(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
}

func TestGetHashTable(t *testing.T) {
	key := randString(12)
	res, err := db.GetHashTable(key)
	if err == nil {
		t.Error("expected", "error", "got", res)
	}
	newHashtable, _ := db.CreateHashTable(key)
	secondRes, secondErr := db.GetHashTable(key)
	if secondErr != nil || secondRes != newHashtable {
		t.Error("expected", newHashtable, "got", secondErr)
	}
}
