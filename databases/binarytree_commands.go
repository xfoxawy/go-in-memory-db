package databases

import (
	"errors"

	"github.com/go-in-memory-db/binarytrees"
)

// Create binary tree funtion
func (db *Database) CreateBinaryTree(k string) (*binarytrees.BinaryTree, error) {
	if binarytree, ok := db.binarytree[k]; ok {
		return binarytree, errors.New("binary tree is exist")
	}
	db.binarytree[k] = binarytree.newBinaryTree()
	return db.binarytree[k], nil
}

// GetBinaryTree Function
func (db *Database) GetBinaryTree(k string) (*binarytree.BinaryTree, error) {
	if _, ok := db.binarytree[k]; ok {
		return db.binarytree[k], nil
	}
	return nil, errors.New("not found")
}

// DelBinaryTree Function
func (db *Database) DelBinaryTree(k string) {
	delete(db.binarytree, k)
}
