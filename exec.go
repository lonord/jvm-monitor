package main

import "os/exec"

func execCmd(cmdStr string) (string, error) {
	debugln("exec cmd:", cmdStr)
	command := exec.Command("/bin/bash", "-c", cmdStr)
	content, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	str := string(content)
	debugf("cmd[%s] output >>>\n%s\n<<<\n", cmdStr, str)
	return str, nil
}
