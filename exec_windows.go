package main

import "os/exec"

func execCmd(cmdStr string) (string, error) {
	debugln("exec windows cmd:", cmdStr)
	command := exec.Command("cmd.exe", "/c", cmdStr)
	content, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	str := string(content)
	debugf("cmd[%s] output >>>\n%s\n<<<\n", cmdStr, str)
	return str, nil
}
