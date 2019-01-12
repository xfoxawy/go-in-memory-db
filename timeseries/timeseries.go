package timeseries

import (
	"time"

	"github.com/xfoxawy/go-in-memory-db/hashtable"
	"github.com/xfoxawy/go-in-memory-db/skiplist"
)

type SnapShot struct {
}

type TimeSeries struct {
	ID       string
	table    map[int64]*hashtable.HashTable
	skiplist *skiplist.SkipList
	expire   time.Duration
	tikcer   *time.Ticker
}

func (t *TimeSeries) Seek() {

}

func (t *TimeSeries) Retrieve(timestamp int64) *hashtable.HashTable {
	if table, ok := t.table[timestamp]; ok {
		return table
	}
	return nil
}

func (t *TimeSeries) Get(timestamp int64, key string) *hashtable.Element {
	table := t.Retrieve(timestamp)

	if table != nil {
		return table.Get(key)
	}

	return nil
}

func (t *TimeSeries) Insert(timestamp int64, key string, value string) *TimeSeries {
	t.table[timestamp].Update(key, value)
	t.skiplist.Set(float64(timestamp), time.Now().Unix())
	return t
}

func (t *TimeSeries) Remove(timestamp int64) *TimeSeries {
	if _, ok := t.table[timestamp]; ok {
		t.table[timestamp].Remove(string(timestamp))
		t.skiplist.Remove(float64(timestamp))
		delete(t.table, timestamp)
	}
	return t
}

func New(id string, expire time.Duration) *TimeSeries {
	return &TimeSeries{
		id,
		make(map[int64]*hashtable.HashTable),
		skiplist.New(),
		expire,
		time.NewTicker(expire),
	}
}
