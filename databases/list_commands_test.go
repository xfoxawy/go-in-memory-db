package databases

import (
	"math/rand"
	"testing"
	"time"
)

// db created before in key value commands file
func init() {
	db = CreateMasterDB()
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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
		t.Error("expected", err, "got", res)
	}
	newList, _ := db.CreateList(key)
	secondRes, secondErr := db.GetList(key)
	if secondErr != nil || secondRes != newList {
		t.Error("expected", newList, "got", secondErr)
	}
}
