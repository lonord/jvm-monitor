package main

import (
	"errors"
	"os/exec"
	"strings"
)

func checkJavaPID(pid string) error {
	if _, err := exec.LookPath("jps"); err == nil {
		out, err := execCmd("jps")
		if err == nil {
			found := false
			for _, line := range strings.Split(out, "\n") {
				for idx, fg := range strings.Split(line, " ") {
					if idx == 0 && fg != "" {
						if fg == pid {
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
