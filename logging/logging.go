package logging

import (
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
