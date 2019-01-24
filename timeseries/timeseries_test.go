package timeseries

import (
	"fmt"
	"testing"
	"time"
)

var ts *Timeseries

func init() {
	ts = New()
}

func TestInsert(t *testing.T) {
	now := time.Now().Unix()
	ts.Insert(now, "testKey", "testValue")

	if ts.Length() != 1 {
		t.Fatal("Invalid length")
	}

	for i := 1; i < 100; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		timestamp := now + int64(i)
		ts.Insert(timestamp, k, v)
	}

	snapshot := ts.First()

	if snapshot.Length() > 1 {
		t.Fatal("Invalid Snapshot length")
	}

	if snapshot[now] == nil {
		t.Fatal("Invalid pointer for exiting point in time")
	}

	for timestamp, table := range snapshot {
		if timestamp != now {
			t.Fatal("Invalid first timestamp")
		}

		if table.Get("testKey") != "testValue" {
			t.Fatal("Invalid value in the hash")
		}
	}

	halfwaypoint := now + int64(50)

	halfwayhash := ts.Retrieve(halfwaypoint)

	if halfwayhash == nil {
		t.Fatal("Missing halfway point")
	}

	halfwaykey := fmt.Sprintf("%s%d", "k", 50)
	halfwayvalue := fmt.Sprintf("%d", 50)

	if halfwayhash.Get(halfwaykey) != halfwayvalue {
		t.Fatal("Missing Key Value for halfway point")
	}

	faltpoint := now + int64(200)

	if snapshot[faltpoint] != nil {
		t.Fatal("Invalid pointer for exiting point in time")
	}

	pastpoint := now - int64(100)

	ts.Insert(pastpoint, "invalidKey", "invalidValue")

	pastsnapshot := ts.Retrieve(pastpoint)

	if pastsnapshot != nil {
		t.Fatal("Invalid Pointer in the past, should not insert older that first point")
	}

	pastvalue := ts.Get(pastpoint, "invalidKey")

	if pastvalue != "" {
		t.Fatal("Invalid Pointer in the past, should not insert older that first point")
	}
}

func TestBulkInsert(t *testing.T) {
	inputs := make(map[string]string)
	now := time.Now().Unix()

	for i := 1; i < 100; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		inputs[k] = v
	}

	ts.BulkInsert(now, inputs)

	hash := ts.Retrieve(now)
	v := hash.Get("k23")
	if v == nil {
		t.Fatal("Non value found")
	}

}
