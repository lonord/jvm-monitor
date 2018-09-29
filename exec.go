package main

import "os/exec"

func execCmd(cmdStr string) (string, error) {
	command := exec.Command("/bin/bash", "-c", cmdStr)
	content, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(content), nil
}
