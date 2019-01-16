package logging

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

// log func
// this is the first function will called from out side
// first validation depend our sets in newLogging
// second run log in newLogging struct
func log(loggingType string, loggingInOption string, loggingContent string) bool {

	newLogging := newLogging()
	validation := newLogging.validateLogParams(loggingType, loggingInOption)

	if validation == false {
		return false
	}

	newLogging.finalType = loggingType
	newLogging.finalOption = loggingInOption
	newLogging.content = loggingContent
	newLogging.log()
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

func (nl logging) log() {

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
