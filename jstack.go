package main

import (
	"fmt"
	"os/exec"
	"path"
	"time"

	"github.com/shirou/gopsutil/process"
)

const dumpCPUThreshold = 100

func checkJStackEnv() error {
	_, err := exec.LookPath("jstack")
	return err
}

func takeJStackDump(proc *process.Process) error {
	cp, err := proc.Percent(time.Second * 1)
	debugln("read jvm cpu usage:", cp)
	if err != nil {
		return err
	}
	if cp < dumpCPUThreshold {
		return nil
	}
	debugln("take stack dump")
	fileLogger.Printf("cpu %.1f, take stack dump", cp)
	dumpFile := path.Join(outputDir, fmt.Sprintf("jstack_dump_%s.log", time.Now().Format("20060102150405")))
	debugln("dump file:", dumpFile)
	_, err = execCmd(fmt.Sprintf("jstack -l %d > %s", proc.Pid, dumpFile))
	if err != nil {
		debugln("take dump failed:", err.Error())
	}
	return err
}
