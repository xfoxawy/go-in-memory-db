package databases

import (
	"errors"

	"github.com/xfoxawy/go-in-memory-db/timeseries"
)

func (db *Database) CreateTimeseries(k string) (*timeseries.Timeseries, error) {
	if v, ok := db.timeseries[k]; ok {
		return v, errors.New("Timeseries Exists")
	}
	db.timeseries[k] = timeseries.New()
	return db.timeseries[k], nil
}

func (db *Database) GetTimeseries(k string) (*timeseries.Timeseries, error) {
	if _, ok := db.timeseries[k]; ok {
		return db.timeseries[k], nil
	}
	return nil, errors.New("Timeseries Does not Exist")
}

func (db *Database) DelTimeseries(k string) {
	delete(db.timeseries, k)
}
