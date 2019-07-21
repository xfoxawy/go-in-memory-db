package timeseries

import (
	"fmt"
	"testing"
	"time"
)

var ts *Timeseries
var (
	now = time.Now().Unix()
)

func TestInsert(t *testing.T) {
	ts = New()
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

		if table.Get("testKey").Value().(string) != "testValue" {
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

	if halfwayhash.Get(halfwaykey).Value().(string) != halfwayvalue {
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
	ts = New()
	inputs := make(map[string]string)

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

func TestAfter(t *testing.T) {
	ts = New()

	beforepoint := now - int64(5)
	snapshot := ts.After(beforepoint, 100)

	if snapshot.Length() != 0 {
		t.Fatal("List should be empty")
	}

	for i := 1; i < 101; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		timestamp := now + int64(i)
		ts.Insert(timestamp, k, v)
	}

	snapshot = ts.After(beforepoint, 100)

	if snapshot.Length() != 100 {
		t.Fatalf("List should be contain 100 elements, returned length %v", snapshot.Length())
	}

	snapshot = ts.After(beforepoint, 50)

	if snapshot.Length() != 50 {
		t.Fatalf("List should be contain 50 elements, returned length %v", snapshot.Length())
	}

	midpoint := now + int64(50)

	snapshot = ts.After(midpoint, 10)

	if snapshot.Length() != 10 {
		t.Fatalf("List should be contain 10 elements, returned length %v", snapshot.Length())
	}

	snapshot = ts.After(midpoint, 88)

	if snapshot.Length() != 52 {
		t.Fatalf("List should be contain 50 elements, returned length %v", snapshot.Length())
	}

	nopoint := now + int64(200)

	snapshot = ts.After(nopoint, 100)

	if snapshot.Length() != 0 {
		t.Fatalf("List should contain 0 elements, returned length is %v", snapshot.Length())
	}

	randompoint := now + int64(23)

	snapshot = ts.After(randompoint, 12)

	for timestamp, _ := range snapshot {
		t.Log(timestamp)
	}
}

func TestBefore(t *testing.T) {
	ts = New()
	beforepoint := now - int64(5)
	snapshot := ts.Before(beforepoint, 100)

	if snapshot.Length() != 0 {
		t.Fatal("List should be empty")
	}

	for i := 1; i < 101; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		timestamp := now + int64(i)
		ts.Insert(timestamp, k, v)
	}

	onebefore := now + int64(1)

	snapshot = ts.Before(onebefore, 100)

	if snapshot.Length() != 1 {
		t.Log("Invalid length for before query")
	}

	fiftybefore := now + int64(50)

	snapshot = ts.Before(fiftybefore, 100)

	if snapshot.Length() != 49 {
		t.Log("Invalid length for before query")
	}

	hunderedbefore := now + 100

	snapshot = ts.Before(hunderedbefore, 100)

	if snapshot.Length() != 99 {
		t.Log("Invalid length for before query")
	}

}

func TestRange(t *testing.T) {
	ts = New()
	for i := 1; i < 101; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		timestamp := now + int64(i)
		ts.Insert(timestamp, k, v)
	}

	afterpoint := now + int64(100)
	snapshot := ts.Range(now, afterpoint)

	if snapshot.Length() != 99 {
		t.Fatal("Invalid Length of snapshot")
	}

	if snapshot[now+2] == nil {
		t.Fatal("Fatal first point in snapshot")
	}

	fatalfirstpoint := now + 1000
	fatallastpoint := fatalfirstpoint + 1000

	snapshot = ts.Range(fatalfirstpoint, fatallastpoint)

	if snapshot.Length() != 0 {
		t.Fatal("Invalid Length, it supposed to be zero")
	}

	wrongstartpoint := now
	wrongafterpoint := now - 100
	snapshot = ts.Range(wrongstartpoint, wrongafterpoint)

	if snapshot.Length() != 0 {
		t.Fatal("Invalid Length, it supposed to be zero")
	}

}

func TestSeekTo(t *testing.T) {
	ts = New()

	for i := 1; i < 101; {
		i++
		k := fmt.Sprintf("%s%d", "k", i)
		v := fmt.Sprintf("%d", i)
		timestamp := now + int64(i)
		ts.Insert(timestamp, k, v)
	}

	hash := ts.SeekTo(now - 100)

	if hash == nil {
		t.Fatal("Hash should not be nil")
	}
}
