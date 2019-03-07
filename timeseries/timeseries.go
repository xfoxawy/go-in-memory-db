package timeseries

import (
	"fmt"
	"log"
	"time"

	"github.com/go-in-memory-db/hashtable"
	"github.com/ryszard/goskiplist/skiplist"
)

// Snapshot is a copied slice from the timeseries, Holding values as a return from queries
// treat Snapshot as immutable instance
type Snapshot map[int64]*hashtable.HashTable

// Push a Key Value in Snapshot
func (s Snapshot) Push(key int64, table *hashtable.HashTable) Snapshot {
	s[key] = table
	return s
}

// Length of how many keys of a Snapshot
func (s Snapshot) Length() int {
	return len(s)
}

// IsEmpty Snapshot
func (s Snapshot) IsEmpty() bool {
	return len(s) == 0
}

// Timeseries data set, consistes of a skiplist and a hastable
// is only moving forward, past is immutable only inserting moving forward in time
type Timeseries struct {
	timemap  *skiplist.SkipList
	table    map[int64]*hashtable.HashTable
	skiplist *skiplist.SkipList
	current  int64 // Last inserted timestamp
	span     time.Duration
	Ticking  time.Duration
}

// NewTimeseries data structure
func NewTimeseries() *Timeseries {
	return &Timeseries{
		timemap: skiplist.NewCustomMap(func(l, r interface{}) bool {
			return l.(int64) < r.(int64)
		}),
		table: make(map[int64]*hashtable.HashTable),
		skiplist: skiplist.NewCustomMap(func(l, r interface{}) bool {
			return l.(int64) < r.(int64)
		}),
		current: 0,
		span:    time.Second,
		Ticking: 500 * time.Millisecond,
	}
}

func Ticker(t *Timeseries) chan bool {
	done := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(t.Ticking)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				t.Clean()
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()
	return done
}

func (t *Timeseries) Clean() {

	now := time.Now()

	tensecondsago := now.Add(-t.span)

	x := t.timemap.Seek(tensecondsago.Unix())

	if x != nil {
		// for x.Next() {
		// 	ts := x.Key().(int64)
		// 	fmt.Printf("%v", ts)
		// }
		log.Printf("value %v", x.Value())
		log.Printf("value %v", x.Previous())
		// for x.Previous() {
		// ts := x.Key().(int64)
		// log.Printf("%v", ts)
		// }
	}

	// range ??za
	// fmt.Println(tensecondsago)

}

// Length of Timeseries
func (t *Timeseries) Length() int {
	return t.skiplist.Len()
}

// BulkInsert keys, values
// Insert is only forward in time
func (t *Timeseries) BulkInsert(timestamp int64, kvs map[string]string) {
	if timestamp > t.current {
		t.timemap.Set(time.Now().Unix(), timestamp)
		t.skiplist.Set(timestamp, time.Now().Unix())
		t.table[timestamp] = hashtable.NewHashTable()
		t.current = timestamp
		for k, v := range kvs {
			t.table[timestamp].Insert(k, v)
		}
	}
}

// Insert a key, value at pin point in time
// Inserting is only moving forward in time
func (t *Timeseries) Insert(timestamp int64, key, value string) {
	if timestamp > t.current {
		// insert in timemap for faster cleaning job
		t.timemap.Set(time.Now().Unix(), timestamp)
		// insert in skiplist (real values) for faster search job
		t.skiplist.Set(timestamp, time.Now().Unix())
		// create and save key-value in hashtable indexed by input timestamp
		t.table[timestamp] = hashtable.NewHashTable()
		t.table[timestamp].Insert(key, value)
		// keep the last timestamp handy for quick comparsion , this timeseries is forward only
		t.current = timestamp
	}
}

// Retrieve a Hashtable from exact timestamp
func (t *Timeseries) Retrieve(timestamp int64) *hashtable.HashTable {
	if table, ok := t.table[timestamp]; ok {
		return table
	}
	return nil
}

// SeekTo an equal or greater than a timestamp
func (t *Timeseries) SeekTo(timestamp int64) *hashtable.HashTable {
	bound := t.skiplist.Seek(timestamp)

	if bound != nil {
		ts := bound.Key().(int64)
		return t.table[ts]
	}

	return nil
}

// Get return an extact timestamp and key from Hashtable associated with this timestamp
func (t *Timeseries) Get(timestamp int64, key string) string {
	if table, ok := t.table[timestamp]; ok {
		if ok := table.Exists(key); ok {
			return table.Get(key).Stringify()
		}
	}
	return ""
}

// Range seeks a snapshot from an equal or greater to start timestamp
// to less than or equal end timestamp
func (t *Timeseries) Range(start, end int64) Snapshot {
	snapshot := make(Snapshot)

	if start > end {
		return snapshot
	}

	bound := t.skiplist.Seek(start)

	if bound != nil {
		ts := bound.Key().(int64)
		table := t.table[ts]
		snapshot.Push(ts, table)

		for bound.Next() && bound.Key().(int64) <= end {
			ts := bound.Key().(int64)
			table := t.table[ts]
			snapshot = snapshot.Push(ts, table)
		}
	}

	return snapshot
}

// First timestamp inserted in timeseries
func (t *Timeseries) First() Snapshot {
	bound := t.skiplist.SeekToFirst()
	snapshot := make(Snapshot)

	defer bound.Close()

	if bound.Key() != nil {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}

	return snapshot
}

// Last timestamp inserted in timeseries
func (t *Timeseries) Last() Snapshot {
	bound := t.skiplist.SeekToLast()
	snapshot := make(Snapshot)

	defer bound.Close()

	if bound.Key() != nil {
		timestamp := bound.Key().(int64)
		table := t.table[timestamp]
		snapshot = snapshot.Push(timestamp, table)
	}

	return snapshot
}

// Before seeks equal or less than timestamp til a span number of elements
func (t *Timeseries) Before(timestamp int64, span int) Snapshot {
	bound := t.skiplist.Seek(timestamp)
	snapshot := make(Snapshot)

	//@todo : running tests on it, it throws nil pointer exceptions
	// defer bound.Close()

	if bound != nil {
		ts := bound.Key().(int64)
		table := t.table[ts]
		snapshot = snapshot.Push(ts, table)

		for bound.Previous() && snapshot.Length() != span {
			ts := bound.Key().(int64)
			table := t.table[ts]
			snapshot = snapshot.Push(ts, table)
		}
	}
	return snapshot
}

// After seeks equal or greater than timestamp til a span number of elements
func (t *Timeseries) After(timestamp int64, span int) Snapshot {
	bound := t.skiplist.Seek(timestamp)
	snapshot := make(Snapshot)

	//@todo : running tests on it, it throws nil pointer exceptions
	// defer bound.Close()

	if bound != nil {
		ts := bound.Key().(int64)
		table := t.table[ts]
		snapshot = snapshot.Push(ts, table)

		for bound.Next() && snapshot.Length() != span {
			ts := bound.Key().(int64)
			table := t.table[ts]
			snapshot = snapshot.Push(ts, table)
		}
	}

	return snapshot
}
