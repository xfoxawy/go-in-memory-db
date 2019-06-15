package databases

import (
	"testing"
)

func TestSet(t *testing.T) {
	key := randString(12)
	value := randString(12)
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
	_, err := db.Get(randString(12))
	if err == nil {
		t.Error("expected", "error", "got nothing")
		return
	}
}

func TestDel(t *testing.T) {
	key := randString(12)
	value := randString(12)
	db.Set(key, value)
	getValue, _ := db.Get(key)
	delValue := db.Del(key)
	if delValue != true {
		t.Error("expected", "true", "got", delValue)
	}
	getValueAfterDel, _ := db.Get(key)
	if getValue == getValueAfterDel {
		t.Error("expected", "", "got", getValueAfterDel)
	}
}
