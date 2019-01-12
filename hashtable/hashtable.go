package hashtable

// Element single for Hashtable
type Element struct {
	value interface{}
}

// Value of Single Element
func (e *Element) Value() interface{} {
	return e.value
}

// HashTable struct
type HashTable struct {
	Values map[string]*Element
}

// NewHashTable generator
func NewHashTable() *HashTable {
	return &HashTable{
		Values: make(map[string]*Element),
	}
}

// Exists check the exitance of a Key
func (h *HashTable) Exists(k string) bool {
	if _, ok := h.Values[k]; ok {
		return ok
	}
	return false
}

// Get retreives an Element from HashTable
func (h *HashTable) Get(k string) *Element {
	if el, ok := h.Values[k]; ok {
		return el
	}
	return nil
}

// Insert an Element HashTable, Does not override
// [index1:value1 ,index2: value2]
//  push value3
// [index1:value1 ,index2: value2, index3: value3]
func (h *HashTable) Insert(k string, v interface{}) *HashTable {
	if _, ok := h.Values[k]; ok {
		return h
	}
	h.Values[k] = &Element{
		value: v,
	}
	return h
}

// Update inserts an Element, overrides if Key exists
func (h *HashTable) Update(k string, v interface{}) *HashTable {
	h.Values[k] = &Element{
		value: v,
	}
	return h
}

// Remove value from HashTable
//[value1 , value2 , ...]
//delete value1
// [value2 , ...]
func (h *HashTable) Remove(key string) *HashTable {
	if _, ok := h.Values[key]; ok {
		delete(h.Values, key)
	}
	return h
}

// Length of Values in HashTable
func (h *HashTable) Length() int {
	return len(h.Values)
}
