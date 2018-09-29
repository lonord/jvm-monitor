package main

import (
	"errors"
	"os/exec"
	"strings"
)

var jpsFound = false

func init() {
	if _, err := exec.LookPath("jps"); err == nil {
		jpsFound = true
		debugln("found jps")
	}
}

func checkJavaPID(pid string) error {
	debugln("check java pid", pid)
	if jpsFound {
		out, err := execCmd("jps")
		if err == nil {
			debugf("jps output >>>\n%s\n<<<\n", out)
			found := false
			for _, line := range strings.Split(out, "\n") {
				for idx, fg := range strings.Split(line, " ") {
					if idx == 0 && fg != "" {
						if fg == pid {
							debugf("pid %s is jvm\n", pid)
							found = true
							break
						}
					}
				}
				if found {
					break
				}
			}
			if !found {
				return errors.New("invalid pid")
			}
		}
	}
	return nil
}
