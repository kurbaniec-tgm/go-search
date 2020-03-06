package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	. "search/src/opts"
	"search/src/searcher"
)

func main() {
	var opts Opts

	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	// fmt.Println(opts.Base)
	// fmt.Printf("Wow %d", opts.Max)
	// fmt.Printf(strings.Join(args, " "))

	// Unsupported command
	if len(args) == 1 {
		// TODO document modes
		fmt.Printf("USAGE : %s [--base <root_directory>] [--max <maximum matches>] <search> \n", os.Args[0])
		fmt.Printf("\t --base or -b : Specifiy directory in which to search for files")
		fmt.Printf("\t --max or -m : Maximum amount of matches")
		os.Exit(0)
	}

	// Add search query to struct
	opts.Search = filepath.FromSlash(args[1])
	opts.Base = filepath.FromSlash(opts.Base)

	/**
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
	}*/

	searcher.FindFiles(opts)
}
