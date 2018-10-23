package databases

import (
	"errors"

	"github.com/go-in-memory-db/queue"
)

func (db *Database) CreateQueue(k string) *queue.Queue {
	if queue, ok := db.queue[k]; ok {
		return queue
	}
	db.queue[k] = queue.NewQueue()
	return db.queue[k]
}

func (db *Database) GetQueue(k string) (*queue.Queue, error) {
	if _, ok := db.queue[k]; ok {
		return db.queue[k], nil
	}
	return nil, errors.New("not found")
}

func (db *Database) DelQueue(k string) {
	delete(db.queue, k)
}
