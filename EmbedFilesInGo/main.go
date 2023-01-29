package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed empty_executable_file_example
var myEmbedFile []byte

func WriteExecutableFile() {
	err := os.WriteFile("my_executable_file", myEmbedFile, 0700)
	if err != nil {
		panic("could not write file")
	}
}

func ExecCommand() []byte {
	output, err := exec.Command("my_executable_file", "arg1", "arg2").Output()
	if err != nil {
		panic("could not run executable")
	}
	return output
}

func RemoveFile() {
	err := os.Remove("my_executable_file")
	if err != nil {
		panic("could not remove exec file")
	}
}

func main() {
	WriteExecutableFile()
	output := ExecCommand()
	RemoveFile()
	fmt.Print(output)
}
