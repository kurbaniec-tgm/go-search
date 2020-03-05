package main

import (
	"fmt"
	"os"
	"search/src/searcher"
)

func main() {
	// Unsupported command
	if len(os.Args) == 1 {
		// TODO document modes
		fmt.Printf("USAGE : %s <target_directory> <target_filename or part of filename> \n", os.Args[0])
		os.Exit(0)
	}
	// Declare variables
	var targetDirectory string
	var fileName string
	// Init target directory and filename
	if len(os.Args) == 2 {
		targetDirectory = "/"
		fileName = os.Args[1]
	} else {
		targetDirectory = os.Args[1]
		fileName = os.Args[2]
	}

	searcher.FindFile(targetDirectory, fileName)
}
