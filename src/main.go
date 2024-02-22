package main

import (
	"nexusutils/cmd"
	"os"
	"path/filepath"
)

func main() {
	var currentDir = ""
	var err error
	if currentDir, err = os.Getwd(); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "nexusutils"), os.ModePerm); err != nil {
		panic(err)
	}
	cmd.Execute()

	// Restore pre-execution working directory
	os.Chdir(currentDir)
}
