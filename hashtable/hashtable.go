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

// Remove value from slice using key
func (h *HashTable) Remove(key string) map[string]string {
	if _, ok := h.Values[key]; ok {
		delete(h.Values, key)
	}
	return h.Values
}

/**
* Seek element value using key
* [value1 , value2 , ...]
* seek 0
* value1
 */
func (h *HashTable) Seek(key string) string {
	if value, ok := h.Values[key]; ok {
		return value
	}
	return ""
}

// Update value in HashTable
func (h *HashTable) Update(key, newVal string) map[string]string {
	if _, ok := h.Values[key]; ok {
		h.Values[key] = newVal
	}
	return h.Values
}

// Find value in HashTable
func (h *HashTable) Find(value string) string {
	for i, v := range h.Values {
		if value == v {
			return i
			break
		}
	}
	return ""
}

// Size of HashTable
func (h *HashTable) Size() int {
	return len(h.Values)
}
