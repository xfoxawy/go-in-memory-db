package databases

import (
	"errors"

	"github.com/go-in-memory-db/stack"
)

// CreateStack Function
func (db *Database) CreateStack(k string) (*stack.Stack, error) {
	if stack, ok := db.stack[k]; ok {
		return stack, errors.New("Stack Exists")
	}
	db.stack[k] = stack.NewStack()
	return db.stack[k], nil
}

// GetStack Function
func (db *Database) GetStack(k string) (*stack.Stack, error) {
	if _, ok := db.stack[k]; ok {
		return db.stack[k], nil
	}
	return nil, errors.New("not found")
}

// DelStack Function
func (db *Database) DelStack(k string) {
	delete(db.stack, k)
}
