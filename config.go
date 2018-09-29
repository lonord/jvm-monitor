package main

import (
	"log"
	"os"
	"os/user"
	"path"
)

var outputDir string

var fileLogger *log.Logger

func init() {
	u, err := user.Current()
	if err != nil {
		outputDir = "/var/log/fr-monitor"
	} else {
		outputDir = path.Join(u.HomeDir, ".fr-monitor")
	}
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalln("create output directory failed:", err)
	}
	logfile := path.Join(outputDir, "output.log")
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("open log file error:", err)
	}
	fileLogger = log.New(f, "", log.LstdFlags)
}
