package logging

import (
	"testing"
)

func TestLog(t *testing.T) {
	log := log("user", "file", "content here")

	if log == false {
		t.Error("error in log func")
	}
}

func TestInArray(t *testing.T) {
	check := in_array("string", []string{"string", "alaa"})
	if check != true {
		t.Error("error in in_array func")
	}
}

func TestvalidateLogParams(t *testing.T) {
	newLog := newLogging()
	validation := newLog.validateLogParams("application", "file")
	if validation != true {
		t.Error("error in validation func")
	}
	newLog = newLogging()
	validation = newLog.validateLogParams("applicatio", "fil")
	if validation != false {
		t.Error("error in validation func")
	}
}
