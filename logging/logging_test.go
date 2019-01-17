package logging

import (
	"os"
	"testing"
)

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

func TestCreatefile(t *testing.T) {

	filePath := "../logs/alaa.log"
	createFile(filePath)
	var _, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		t.Error("error in create file func")
	}
	os.Remove(filePath)
}

func TestWriteInFile(t *testing.T) {
	for i := 0; i < 10; i++ {
		check := writeInFile("../logs/fortest.log", "content here of error")

		if check == false {
			t.Error("error in write in file function")

		}
	}
}
