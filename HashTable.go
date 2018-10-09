package main

type HashTable struct {
	values map[string]string
}

func NewHashTable() *HashTable {
	return &HashTable{values: make(map[string]string)}
}

/**
* push value in map
* [index1:value1 ,index2: value2]
* push value3
* [index1:value1 ,index2: value2, index3: value3]
 */
func (h *HashTable) push(k string, v string) map[string]string {
	if _, ok := h.values[k]; ok {
		return h.values
	} else {
		h.values[k] = v
		return h.values
	}
}

/**
* delete value from slice
* [value1 , value2 , ...]
* delete value1
* [value2 , ...]
 */
func (h *HashTable) remove(key string) map[string]string {
	if _, ok := h.values[key]; ok {
		delete(h.values, key)
	}
	return h.values
}

/**
* delete value using index
* [index1 , index2 , ...]
* unlink 0
* [index2 , ...]
* we can't use it in loop becuase we will have indexing issue
 */
// func (h *HashTable) unlink(index int) []string {
// 	if index >= 0 && index < len(h.values) {
// 		result := append(h.values[:index], h.values[index+1:]...)
// 		return result
// 	}
// 	return h.values
// }

// /**
// * get element value using index
// * [value1 , value2 , ...]
// * seek 0
// * value1
//  */
// func (h *HashTable) seek(index int) string {
// 	if index >= 0 && index < len(h.values) {
// 		return h.values[index]
// 	}
// 	return ""
// }

// func getElementIndex(element string, array []string) int {
// 	for i := range array {
// 		if element == array[i] {
// 			return i
// 			break
// 		}
// 	}
// 	return -1
// }
