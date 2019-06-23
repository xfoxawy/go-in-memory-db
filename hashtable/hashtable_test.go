package hashtable

import (
	"testing"
)

var hash *HashTable

func init() {
	hash = NewHashTable()
}

func hashTableTamplate() *HashTable {
	newHash := NewHashTable()
	newHash.Push("key1", "value1")
	newHash.Push("key2", "value2")
	newHash.Push("key3", "value3")
	newHash.Push("key4", "value4")
	newHash.Push("key5", "value5")
	return newHash
}

func TestPush(t *testing.T) {
	res1 := hash.Push("key", "value")
	if len(hash.Values) == 0 {
		t.Errorf("expected length > 0")
	}
	res2 := hash.Push("key", "value2")
	if res1["key"] != res2["key"] {
		t.Errorf("expected push one only")
	}
}

func TestRemove(t *testing.T) {
	hash.Push("key", "value")
	beforeRemove := len(hash.Values)
	hash.Remove("key")
	afterRemove := len(hash.Values)
	if afterRemove != beforeRemove-1 {
		t.Error("expected", beforeRemove-1, "got", afterRemove)
	}
}

func TestSeek(t *testing.T) {
	seekHash := hashTableTamplate()
	if seekHash.Seek("key3") != "value3" {
		t.Error("expected", "value3", "got", seekHash.Seek("key3"))
	}
}

func TestUpdate(t *testing.T) {
	updateHash := hashTableTamplate()
	updatedValue := updateHash.Update("key3", "value33")
	if updatedValue["key3"] != "value33" {
		t.Error("expected", "value33", "got", updatedValue["key3"])
	}
}

func TestFind(t *testing.T) {
	findHash := hashTableTamplate()

	if findHash.Find("value3") != "key3" {
		t.Error("expected", "key3", "got", findHash.Find("value3"))
	}
}

func TestSize(t *testing.T) {
	sizeHash := hashTableTamplate()
	if sizeHash.Size() != 5 {
		t.Error("expected", 3, "got", sizeHash.Size())
	}
}
