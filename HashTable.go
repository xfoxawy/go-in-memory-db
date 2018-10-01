package main

type HashTable struct {
	values []string
}

func NewHashTable() *HashTable {
	var values []string
	return &HashTable{values: values}
}

/**
* push value in slice
* [value1 , value2]
* push value3
* [value1 , value2 , value3]
 */
func (h *HashTable) push(v string) []string {
	result := append(h.values, v)
	return result
}

/**
* delete value from slice
* [value1 , value2 , ...]
* delete value1
* [value2 , ...]
 */
func (h *HashTable) remove(value string) []string {
	i := getElementIndex(value, h.values)
	if i != -1 {
		result := append(h.values[:i], h.values[i+1:]...)
		return result
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
func (h *HashTable) unlink(index int) []string {
	if index >= 0 && index < len(h.values) {
		result := append(h.values[:index], h.values[index+1:]...)
		return result
	}
	return h.values
}

func getElementIndex(element string, array []string) int {
	for i := range array {
		if element == array[i] {
			return i
			break
		}
	}
	return -1
}
