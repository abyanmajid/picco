package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func executePython(code string, args []string) (string, error) {
	cmd := exec.Command("python3", append([]string{"-c", code}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}
	return out.String(), nil
}

func executeJava(code string, args []string) (string, error) {
	cmd := exec.Command("java", append([]string{"-cp", ".", code}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}
	return out.String(), nil
}

func executeCpp(code string, args []string) (string, error) {
	cmd := exec.Command("g++", append([]string{"-o", "a.out", "-xc++", "-"}, args...)...)
	cmd.Stdin = bytes.NewBufferString(code)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}

	cmd = exec.Command("./a.out")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}
	return out.String(), nil
}

func executeJavaScript(code string, args []string) (string, error) {
	cmd := exec.Command("node", append([]string{"-e", code}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}
	return out.String(), nil
}
