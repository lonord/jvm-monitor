package main

import (
	"log"
	"os"
)

var debugLogger *log.Logger

func debugf(format string, v ...interface{}) {
	if debugLogger == nil {
		return
	}
	debugLogger.Printf(format, v...)
}

func debugln(v ...interface{}) {
	if debugLogger == nil {
		return
	}
	debugLogger.Println(v...)
}

func enableDebugLogger() {
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	debugLogger.Println("enable debug logger")
}

func disableDebugLogger() {
	if debugLogger != nil {
		debugLogger.Println("disable debug logger")
	}
	debugLogger = nil
}
