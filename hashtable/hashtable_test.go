package hashtable

import "testing"

var hash *HashTable

func init() {
	hash = NewHashTable()
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
