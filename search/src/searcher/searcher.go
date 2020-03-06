package searcher

import (
	"fmt"
	"github.com/gobwas/glob"
	"io/ioutil"
	"path/filepath"
	. "search/src/opts"
	"sync"
)

// Searches for files and directories.
func FindFiles(opts *Opts) {
	// Create top-level WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Start search process
	find(opts, opts.Base, opts.Search, wg)
	// Wait for finish
	wg.Wait()
}

// Searches for files and directories asynchronously with goroutines.
//
// WaitGroup needs to be a pointer, so that Go does not make a unnecessary
// copy that calls done.
// Source: https://stackoverflow.com/a/25234899
func find(opts *Opts, path string, pattern glob.Glob, wg *sync.WaitGroup) {
	// WaitGroup that counts all goroutines this goroutine starts
	nextWg := &sync.WaitGroup{}
	// Flag that checks if the element is the last element in the file tree
	end := true
	// Get all files/directories in the current path
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, entry := range files {
			// Check if a match was found on the stop flag or
			// the requested matches exceeded limit
			if opts.Check() {
				// Build new path from file/directory
				filePath := filepath.Join(path, entry.Name())
				// Check if filename or filepath matches pattern
				check := pattern.Match(entry.Name()) || pattern.Match(filePath)
				// If so print it, check stop status and increase counter
				if check {
					fmt.Println(filePath)
					opts.MatchFound()
				}
				// If entry is a directory spawn a new goroutine that searches for files
				// in that directory
				if entry.IsDir() {
					end = false
					nextWg.Add(1)
					go find(opts, filePath, pattern, nextWg)
				}
			} else {
				break
			}
		}
		// If this function spawned goroutines, wait for them to finish
		if !end {
			nextWg.Wait()
		}
	}
	wg.Done()
}
