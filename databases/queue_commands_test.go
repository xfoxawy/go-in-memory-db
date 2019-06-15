package databases

import "testing"

func TestCreateQueue(t *testing.T) {
	key := randString(12)
	res, err := db.CreateQueue(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
	secondRes, secondErr := db.CreateQueue(key)
	if secondErr == nil {
		t.Error("expected", "error message", "got", secondRes)
	}
}

func TestDelQueue(t *testing.T) {
	key := randString(12)
	db.CreateQueue(key)
	db.DelQueue(key)
	res, err := db.CreateQueue(key)
	if err != nil {
		t.Error("expected", res, "got error", err)
	}
}

func TestGetQueue(t *testing.T) {
	key := randString(12)
	res, err := db.GetQueue(key)
	if err == nil {
		t.Error("expected", "error", "got", res)
	}
	newHashtable, _ := db.CreateQueue(key)
	secondRes, secondErr := db.GetQueue(key)
	if secondErr != nil || secondRes != newHashtable {
		t.Error("expected", newHashtable, "got", secondErr)
	}
}
