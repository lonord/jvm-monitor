package main

import (
	"log"
	"os"
	"os/user"
	"path"
)

var (
	appName    = "__unknow__"
	appVersion = ""
)

var outputDir string

var fileLogger *log.Logger

func initConfig() {
	u, err := user.Current()
	if err != nil {
		outputDir = "/var/log/" + appName
	} else {
		outputDir = path.Join(u.HomeDir, "."+appName)
	}
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalln("create output directory failed:", err)
	}
	debugln("create output directory:", outputDir)
	logfile := path.Join(outputDir, "output.log")
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("open log file error:", err)
	}
	fileLogger = log.New(f, "", log.LstdFlags)
}
