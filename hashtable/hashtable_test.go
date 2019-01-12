package hashtable

import (
	"testing"
)

var hash *HashTable

func init() {
	hash = NewHashTable()
}

func TestGet(t *testing.T) {
	el := hash.Get("fakekey")

	if el != nil {
		t.Error("This Fake Key has a value")
	}

	hash.Insert("key", 5675)

	el = hash.Get("key")

	if el.Value() != 5675 {
		t.Error("Invalid Key value on Get")
	}

	if hash.Length() != 1 {
		t.Error("Invalid Hash Keys length")
	}
}

func TestInsert(t *testing.T) {
	res1 := hash.Insert("key", "value")
	if hash.Length() == 0 {
		t.Errorf("expected length > 0")
	}
	res2 := hash.Insert("key", "value2")
	if res1.Get("key") != res2.Get("key") {
		t.Errorf("expected push one only")
	}
}

func TestRemove(t *testing.T) {
	hash.Insert("key", "value")
	beforeRemove := hash.Length()
	hash.Remove("key")
	afterRemove := hash.Length()
	if afterRemove != beforeRemove-1 {
		t.Error("expected", beforeRemove-1, "got", afterRemove)
	}
}

func TestUpdate(t *testing.T) {
	el := hash.Get("key")
	if el != nil {
		t.Error("Key has a value when its supposed to be nil")
	}

	hash.Insert("key", "firstValue")

	hash.Update("key", "secondValue")

	el = hash.Get("key")

	if el.Value() != "secondValue" {
		t.Error("Key is not updated")
	}
}

func TestLength(t *testing.T) {
	if hash.Length() != 0 {
		t.Error("Hash length hould be zero")
	}

	hash.Insert("key", 5.43)

	if hash.Length() != 1 {
		t.Error("Hash length hould be one")
	}

	hash.Update("key", "differentValue")

	if hash.Length() != 1 {
		t.Error("Hash length hould be one")
	}

	hash.Insert("differentKey", 543)

	if hash.Length() != 2 {
		t.Error("Hash length hould be two")
	}
}

func TestExists(t *testing.T) {
	if hash.Exists("FakeKey") != false {
		t.Error("Exists failed to locate a key")
	}

	hash.Insert("key", "value")

	if hash.Exists("key") != true {
		t.Error("Exists failed to locate a key")
	}
}
