package main

import (
	"errors"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

var procNotExistError error

func init() {
	procNotExistError = errors.New("ProcNotExist")
}

func getPidInfo(p string) (*process.Process, error) {
	pid, _ := strconv.Atoi(p)
	ext, err := process.PidExists(int32(pid))
	if err != nil {
		return nil, err
	}
	if !ext {
		return nil, procNotExistError
	}
	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	return proc, nil
}
