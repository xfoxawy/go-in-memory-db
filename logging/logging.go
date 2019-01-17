package logging

import (
	"log"
	"os"
	"sync"
)

// logging struct
// loggingType user || application
// loggingInOptions file || other storage
// loggingContent string of log
type logging struct {
	loggingTypes     []string
	loggingInOptions []string
	finalType        string
	finalOption      string
	content          string
}

var (
	mutex = &sync.Mutex{}
)

// log func
// this is the first function will called from out side
// first validation depend our sets in newLogging
// second run log in newLogging struct
func loggingLog(loggingType string, loggingInOption string, loggingContent string) bool {

	newLogging := newLogging()
	validation := newLogging.validateLogParams(loggingType, loggingInOption)

	if validation == false {
		return false
	}

	newLogging.finalType = loggingType
	newLogging.finalOption = loggingInOption
	newLogging.content = loggingContent
	newLogging.startLogProcess()
	return true
}

// newLogging func
// create newLogging
func newLogging() *logging {
	return &logging{
		[]string{"user", "application"},
		[]string{"file"},
		"",
		"",
		"",
	}
}

// validateLogParams
func (nl logging) validateLogParams(loggingType string, loggingInOption string) bool {
	if in_array(loggingType, nl.loggingTypes) && in_array(loggingInOption, nl.loggingInOptions) {
		return true
	}
	return false
}

func (nl logging) startLogProcess() {

	switch nl.finalOption {
	case "file":
		nl.logInFile()
	default:
		nl.logInFile()
	}

}

func (nl logging) logInFile() {

	logFilesPath := "../logs/"
	fileName := nl.finalType + nl.finalOption + ".log"

	createdFile := createFile(logFilesPath + fileName)
	writeInFile(createdFile, nl.content)
}

func writeInFile(filePath string, content string) bool {

	mutex.Lock()
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(content)
	mutex.Unlock()
	return true
}

func createFile(filePath string) string {
	var _, err = os.Stat(filePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, _ = os.Create(filePath)
		defer file.Close()
	}
	return filePath
}

// in_array func
// helper function in valiadtion
func in_array(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
