package logging

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func writeInFile(filePath string, content string) bool {

	mutex.Lock()
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.Warnln(content)
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
