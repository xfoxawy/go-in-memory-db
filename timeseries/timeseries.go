package timeseries

import (
	"time"

	"github.com/go-in-memory-db/hashtable"
	"github.com/ryszard/goskiplist/skiplist"
)

type Snapshot map[int64]*hashtable.HashTable

func (s Snapshot) Push(key int64, table *hashtable.HashTable) Snapshot {
	s[key] = table
	return s
}

func (s Snapshot) Length() int {
	return len(s)
}

func (s Snapshot) IsEmpty() bool {
	return len(s) == 0
}

type Timeseries struct {
	table    map[int64]*hashtable.HashTable
	skiplist *skiplist.SkipList
	current  int64 // Last inserted timestamp
}

// New Timeseries data structure
func New() *Timeseries {
	return &Timeseries{
		table: make(map[int64]*hashtable.HashTable),
		skiplist: skiplist.NewCustomMap(func(l, r interface{}) bool {
			return l.(int64) < r.(int64)
		}),
		current: 0,
	}
}
func (t *Timeseries) Length() int {
	return t.skiplist.Len()
}

func (t *Timeseries) BulkInsert(timestamp int64, kvs map[string]string) {
	if timestamp > t.current {
		t.skiplist.Set(timestamp, time.Now().Unix())
		t.table[timestamp] = hashtable.New()
		t.current = timestamp
		for k, v := range kvs {
			t.table[timestamp].Push(k, v)
		}
	}
}
func (t *Timeseries) Insert(timestamp int64, key, value string) {
	if timestamp > t.current {
		t.skiplist.Set(timestamp, time.Now().Unix())
		t.table[timestamp] = hashtable.New()
		t.table[timestamp].Insert(key, value)
		t.current = timestamp
	}
}

func (t *Timeseries) Retrieve(timestamp int64) *hashtable.HashTable {
	if table, ok := t.table[timestamp]; ok {
		return table
	}
	return nil
}

func (t *Timeseries) Get(timestamp int64, key string) string {
	if table, ok := t.table[timestamp]; ok {
		if ok := table.Exist(key); ok {
			return table.Get(key).(string)
		}
	}
	return ""
}

func (t *Timeseries) Range(start, end int64) Snapshot {
	bound := t.skiplist.Range(start, end)
	snapshot := make(Snapshot)

	for bound.Next() {
		timestamp := bound.Key().(int64)
		content := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, content)
	}

	return snapshot
}

func (t *Timeseries) First() Snapshot {
	bound := t.skiplist.SeekToFirst()
	snapshot := make(Snapshot)

	if bound.Key() != nil {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}

	return snapshot
}

func (t *Timeseries) Last() Snapshot {
	bound := t.skiplist.SeekToLast()
	snapshot := make(Snapshot)

	if bound.Key() != nil {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}

	return snapshot
}

func (t *Timeseries) Before(timestamp int64, span int) Snapshot {
	bound := t.skiplist.Seek(timestamp)
	snapshot := make(Snapshot)

	for bound.Previous() && snapshot.Length() != span {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}
	return snapshot
}

func (t *Timeseries) After(timestamp int64, span int) Snapshot {
	bound := t.skiplist.Seek(timestamp)
	snapshot := make(Snapshot)

	for bound.Next() && snapshot.Length() != span {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}
	return snapshot
}
