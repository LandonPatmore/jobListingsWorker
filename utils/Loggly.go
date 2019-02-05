package utils

import (
	"fmt"
	"os"
)

const logglyURL string = "http://logs-01.loggly.com/inputs/"

// TODO: DO NOT UPLOAD THIS
var logglyAPIKey = os.Getenv("LOGGLY_API_KEY")

// Error log.
func Error(message interface{}) {
	sendOutLogMessage("Error", message)
}

// Warn log.
func Warn(message interface{}) {
	sendOutLogMessage("Warn", message)
}

// Debug log.
func Debug(message interface{}) {
	sendOutLogMessage("Debug", message)
}

// Info log.
func Info(message interface{}) {
	sendOutLogMessage("Info", message)
}

// Trace log.
func Trace(message interface{}) {
	sendOutLogMessage("Trace", message)
}

// Will print out to the console what the log message is, as well as send it to loggly.
func EchoLog(tag string, message interface{}) {
	sendOutLogMessage(tag, message)
	fmt.Printf("Tag: %s\t Message: %s", tag, message)
}

// Abstracts the message sending to loggly.
func sendOutLogMessage(tag string, message interface{}) {
	var url = buildURL(tag)

	switch m := message.(type) {
	case string:
		Post(url, string(m))
	default:
		PostJson(url, message)
	}
}

// Builds a URL for the log messages to be sent to.
func buildURL(tag string) string {
	return logglyURL + logglyAPIKey + "/tag/" + tag
}
