package main

import (
	"fmt"
	"github.com/gobwas/glob"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	. "search/src/opts"
	"search/src/searcher"
)

func main() {
	// Parse arguments
	var opts Opts
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}
	// Unsupported command
	if len(args) == 1 {
		fmt.Printf("USAGE : %s [--base <root_directory>] [--max <maximum matches>] <search> \n", os.Args[0])
		fmt.Printf("\t --base or -b : Specifiy directory in which to search for files")
		fmt.Printf("\t --stop or -s : Stops search on first match")
		fmt.Printf("\t --max or -m : Maximum amount of matches")
		os.Exit(0)
	}
	// Add search query as Glob and root directory to struct
	opts.Search = glob.MustCompile(filepath.FromSlash(args[1]))
	opts.Base = filepath.FromSlash(opts.Base)
	// Initialize Stop flag
	opts.StopFlag = false
	// Initialize Counter
	opts.MaxCounter = 0

	// Find files!
	searcher.FindFiles(&opts)
}
