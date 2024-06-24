package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	utils "github.com/abyanmajid/codemore.io/compiler/utils"
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
	// Extract the class name
	className, err := utils.ExtractClassName(code)
	if err != nil {
		return "", fmt.Errorf("could not extract class name: %v", err)
	}

	// Save the code to a file
	fileName := className + ".java"
	err = os.WriteFile(fileName, []byte(code), 0644)
	if err != nil {
		return "", fmt.Errorf("could not write java file: %v", err)
	}

	// Compile the Java code
	cmd := exec.Command("javac", fileName)
	var compileOut bytes.Buffer
	cmd.Stdout = &compileOut
	var compileErr bytes.Buffer
	cmd.Stderr = &compileErr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("compilation error: %v: %s", err, compileErr.String())
	}

	// Execute the compiled Java code
	cmd = exec.Command("java", append([]string{className}, args...)...)
	var runOut bytes.Buffer
	cmd.Stdout = &runOut
	var runErr bytes.Buffer
	cmd.Stderr = &runErr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("execution error: %v: %s", err, runErr.String())
	}

	return runOut.String(), nil
}

func executeCpp(code string, args []string) (string, error) {
	// Save the code to a file
	sourceFile := "main.cpp"
	err := os.WriteFile(sourceFile, []byte(code), 0644)
	if err != nil {
		return "", fmt.Errorf("could not write C++ file: %v", err)
	}

	// Compile the C++ code
	cmd := exec.Command("g++", "-o", "a.out", sourceFile)
	var compileOut bytes.Buffer
	cmd.Stdout = &compileOut
	var compileErr bytes.Buffer
	cmd.Stderr = &compileErr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("compilation error: %v: %s", err, compileErr.String())
	}

	// Execute the compiled C++ code
	cmd = exec.Command("./a.out", args...)
	var runOut bytes.Buffer
	cmd.Stdout = &runOut
	var runErr bytes.Buffer
	cmd.Stderr = &runErr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("execution error: %v: %s", err, runErr.String())
	}

	return runOut.String(), nil
}

func executeJavaScript(code string, args []string) (string, error) {
	// Save the code to a file
	fileName := "script.js"
	err := os.WriteFile(fileName, []byte(code), 0644)
	if err != nil {
		return "", fmt.Errorf("could not write JavaScript file: %v", err)
	}

	// Execute the JavaScript code
	cmd := exec.Command("node", append([]string{fileName}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("execution error: %v: %s", err, stderr.String())
	}

	return out.String(), nil
}
