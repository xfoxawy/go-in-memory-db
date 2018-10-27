package hashtable

// HashTable struct
type HashTable struct {
	Values map[string]string
}

// NewHashTable generator
func NewHashTable() *HashTable {
	return &HashTable{Values: make(map[string]string)}
}

// Push value in map
// [index1:value1 ,index2: value2]
//  push value3
// [index1:value1 ,index2: value2, index3: value3]
func (h *HashTable) Push(k string, v string) map[string]string {
	if _, ok := h.Values[k]; ok {
		return h.Values
	}
	h.Values[k] = v
	return h.Values
}

// Remove value from slice
//[value1 , value2 , ...]
//delete value1
// [value2 , ...]
func (h *HashTable) Remove(key string) map[string]string {
	if _, ok := h.Values[key]; ok {
		delete(h.Values, key)
	}
	return h.Values
}
