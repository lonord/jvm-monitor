package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const checkInterval = time.Second * 10

func init() {
	if appName == "__unknow__" {
		log.Fatalln("invalid exectuable: build flags error")
	}
}

func showHelpAndExit() {
	fmt.Printf("\n%s version %s by Loy B. <lonord@gmail.com>\n\n", appName, appVersion)
	fmt.Printf("    Usage: %s <jvm_pid>\n\n", appName)
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		showHelpAndExit()
	}
	pid := os.Args[1]
	if debugEnv, found := os.LookupEnv("DEBUG"); found && debugEnv != "" {
		enableDebugLogger()
	}
	initConfig()
	_, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatalln("invalid pid:", pid)
	}
	err = checkJavaPID(pid)
	if err != nil {
		log.Fatalln("not a jvm pid:", pid)
	}
	err = checkEnv()
	if err != nil {
		log.Fatalln("check env error:", err)
	}
	fileLogger.Println("start monitor on pid", pid)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	tickTimer := time.NewTimer(checkInterval)
	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGHUP:
				fallthrough
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGQUIT:
				fallthrough
			case syscall.SIGTERM:
				fileLogger.Println("got signal [", sig.String(), "] stop")
				os.Exit(0)
			}
		case <-tickTimer.C:
			err := doAction(pid)
			if err != nil {
				fileLogger.Println("tick action error:", err)
			}
			tickTimer.Reset(checkInterval)
		}
	}
}

func checkEnv() error {
	if err := checkJStackEnv(); err != nil {
		return err
	}
	return nil
}

func doAction(pid string) error {
	proc, err := getPidInfo(pid)
	if err != nil {
		if err == procNotExistError {
			fileLogger.Printf("proc with pid %s is exited\n\n\n", pid)
			debugln("jvm exited, stop monitor")
			os.Exit(0)
		}
		return err
	}
	return takeJStackDump(proc)
}
